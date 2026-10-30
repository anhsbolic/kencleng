"""Static OAS3.0 schema/example checks, not runtime conformance."""
from pathlib import Path
import copy, hashlib, json
from jsonschema import Draft4Validator, FormatChecker

here = Path(__file__).resolve().parent
doc = json.loads((here / 'contract-dereferenced.json').read_text())
checks, failures, request_examples = [], [], []

def convert(schema):
    if isinstance(schema, list):
        return [convert(x) for x in schema]
    if not isinstance(schema, dict):
        return schema
    result = {k:convert(v) for k,v in schema.items() if k not in ('nullable','example','xml','discriminator','externalDocs')}
    if schema.get('nullable') and 'type' in result:
        result['type'] = [result['type'], 'null']
        if 'enum' in result and None not in result['enum']:
            result['enum'].append(None)
    return result

def validate(label, schema, value, expected=True):
    validator = Draft4Validator(convert(schema), format_checker=FormatChecker())
    errors = list(validator.iter_errors(value))
    passed = (not errors) == expected
    checks.append({'case':label,'expected':'valid' if expected else 'invalid','passed':passed})
    if not passed:
        failures.append({'case':label,'errors':[e.message for e in errors] or ['Unexpectedly accepted']})

def walk(value, pointer=''):
    if not isinstance(value, (dict,list)):
        return
    if isinstance(value,list):
        for i,v in enumerate(value): walk(v,f'{pointer}/{i}')
        return
    if 'schema' in value:
        schema=value['schema']
        samples=[]
        if 'example' in value: samples.append(('example',value['example']))
        for name, ex in value.get('examples',{}).items():
            if isinstance(ex,dict) and 'value' in ex: samples.append((name,ex['value']))
        for name, ex in samples:
            validate(f'{pointer}/{name}',schema,ex)
            if '/requestBody/' in pointer: request_examples.append((pointer,schema,ex))
    # Schema-level examples are validated against their owning schema too.
    if 'example' in value and any(k in value for k in ('type','allOf','oneOf','anyOf','$ref')):
        validate(pointer+'/schema-example',value,value['example'])
    for k,v in value.items(): walk(v,pointer+'/'+k)

walk(doc)
operation_coverage=[]
cookie_required={('post','/auth/google/attempts','201'),('get','/auth/google/callback','303'),('post','/auth/refresh','204'),('post','/auth/logout','204'),('post','/account-deletions','200')}
for path,item in doc['paths'].items():
    for method,op in item.items():
        if method not in ('get','post','put','patch','delete','head','options','trace'): continue
        for status,response in op['responses'].items():
            label=method.upper()+' '+path+' '+status
            headers=response.get('headers',{})
            if headers.get('Cache-Control',{}).get('required') is not True:
                failures.append({'case':label,'errors':['Cache-Control required flag missing']})
            if (method,path,status) in cookie_required and headers.get('Set-Cookie',{}).get('required') is not True:
                failures.append({'case':label,'errors':['Mandatory Set-Cookie required flag missing']})
            if status in ('204','303') and response.get('content'):
                failures.append({'case':label,'errors':['No-body branch has content']})
            if status.startswith(('4','5')):
                media=response.get('content',{}).get('application/problem+json')
                if not media:
                    failures.append({'case':label,'errors':['Missing problem+json']})
                    continue
                schema=media['schema']
                for variant in schema.get('oneOf',[schema]):
                    props=variant.get('properties',{})
                    if props.get('status',{}).get('enum') != [int(status)]:
                        failures.append({'case':label,'errors':['Problem variant status does not equal HTTP status']})
                    if not set(('type','title','status')).issubset(variant.get('required',[])):
                        failures.append({'case':label,'errors':['Problem lacks required discriminator/status/title']})
        operation_coverage.append({'operation':method.upper()+' '+path,'operationId':op['operationId'],'security':op.get('security'),'statuses':list(op['responses']),'bodyRequired':op.get('requestBody',{}).get('required',False)})
if len(request_examples) < 7:
    failures.append({'case':'request-example-coverage','errors':['Require examples for all six POST bodies, including both attempt-purpose branches (at least7).']})
for pointer,schema,example in request_examples:
    if not isinstance(example,dict): continue
    if schema.get('additionalProperties') is False:
        bad=copy.deepcopy(example); bad['unexpectedPrivilege']=True
        validate(pointer+'/reject-unknown-field',schema,bad,False)
    for key in schema.get('required',[]):
        if key in example:
            bad=copy.deepcopy(example); del bad[key]
            validate(pointer+'/reject-missing-'+key,schema,bad,False)
    for key,prop in schema.get('properties',{}).items():
        if key in example and not prop.get('nullable'):
            bad=copy.deepcopy(example); bad[key]=None
            validate(pointer+'/reject-null-'+key,schema,bad,False)

report={'contractSha256':hashlib.sha256((here/'openapi-handoff.yaml').read_bytes()).hexdigest(),'validator':'jsonschema Draft4 + OAS3.0 nullable conversion; dereferenced schemas; format checking enabled','checks':checks,'failures':failures,'requestExampleCount':len(request_examples),'operationCoverage':operation_coverage,'additionalChecks':'Required no-store headers on every branch; mandatory cookie headers; no-body204/303; every error variant required discriminator/title/status and HTTP/status equality.','limits':['No runtime behavior asserted. Conditional auth, canonical Unicode normalization, transport/header/cookie attributes, timing, atomicity, journal/restore require normative conformance cases.','OAS3.0 is not full JSON Schema; Redocly separately checks OpenAPI document structure. These checks validate instances under the explicitly documented schema conversion.']}
(here/'example-checks.json').write_text(json.dumps(report,indent=2)+'\n')
print(json.dumps({'checks':len(checks),'failures':len(failures),'requestExamples':len(request_examples)}))
if failures: raise SystemExit(1)
