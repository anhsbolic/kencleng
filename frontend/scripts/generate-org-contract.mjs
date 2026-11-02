import { readFileSync, writeFileSync } from 'node:fs';
import { createHash } from 'node:crypto';
import { execFileSync } from 'node:child_process';
const source = '../harscode-state/org-01/rancang-002/openapi-handoff.yaml';
const bytes = readFileSync(source);
if(createHash('sha256').update(bytes).digest('hex') !== 'a2d21ee3ca4ea93c7e691a53255ab825b784da00f45baeeb88115e76564df9aa') throw Error('Accepted contract drift');
execFileSync('node', ['node_modules/openapi-typescript/bin/cli.js',source,'-o','features/org-01/generated/api.ts'],{stdio:'inherit'});
const spec=JSON.parse(bytes);
function schema(s){
 if(s.oneOf) return `z.union([${s.oneOf.map(schema).join(',')}])`;
 if(s.$ref) return s.$ref.split('/').at(-1);
 if(s.enum) return s.enum.length===1?`z.literal(${JSON.stringify(s.enum[0])})`:`z.enum(${JSON.stringify(s.enum)})`;
 if(s.type==='object') return `z.strictObject({${Object.entries(s.properties??{}).map(([k,v])=>`${JSON.stringify(k)}:${schema(v)}${s.required?.includes(k)?'':'.optional()'}`).join(',')}})`;
 if(s.type==='array') return `z.array(${schema(s.items)})${s.minItems?`.min(${s.minItems})`:''}${s.maxItems?`.max(${s.maxItems})`:''}`;
 if(s.type==='boolean') return 'z.boolean()';
 if(s.type==='integer') return 'z.number().int()';
 let x='z.string()';
 if(s.format==='uuid')x+='.uuid()';
 if(s.format==='date-time')x+='.datetime({offset:false})';
 if(s.minLength)x+=`.refine(value => [...value].length >= ${s.minLength})`;
 if(s.maxLength)x+=`.refine(value => [...value].length <= ${s.maxLength})`;
 if(s.pattern)x+=`.regex(new RegExp(${JSON.stringify(s.pattern)}))`;
 return x;
}
let out='// Generated from accepted r4. Run npm run generate:org; do not edit.\nimport { z } from "zod";\n';
for(const [k,v] of Object.entries(spec.components.schemas))out+=`export const ${k} = ${schema(v)};\n`;
out+='export const operations = {\n';
for(const [path,methods] of Object.entries(spec.paths)) for(const [method,op] of Object.entries(methods)){
 out+=`${op.operationId}:{path:${JSON.stringify(path)},method:${JSON.stringify(method.toUpperCase())},responses:{`;
 for(const [status,r]of Object.entries(op.responses)){
 const responseSchema=Object.values(r.content??{})[0]?.schema;
 out+=`${status}:${responseSchema?schema(responseSchema):'z.undefined()'},`;
 }out+='}},\n';
}out+='} as const;\n';writeFileSync('features/org-01/generated/validation.ts',out);
