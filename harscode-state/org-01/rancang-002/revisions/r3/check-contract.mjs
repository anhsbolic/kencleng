// Documentation verification only. All writes are confined to this artifact directory.
import fs from 'node:fs';
import path from 'node:path';
import { createRequire } from 'node:module';
import { pathToFileURL, fileURLToPath } from 'node:url';
import crypto from 'node:crypto';
const here = path.dirname(fileURLToPath(import.meta.url));
const project = path.resolve(here, '../../../../../');
const require = createRequire(path.join(project, 'frontend/package.json'));
const core = require('@redocly/openapi-core');
const ts = require('typescript');
const { default: openapiTS, astToString } = await import(pathToFileURL(path.join(project,'frontend/node_modules/openapi-typescript/dist/index.mjs')));
const source = path.join(here, 'openapi-handoff.yaml');
const raw = fs.readFileSync(source, 'utf8');
const parsed = core.parseYaml(raw);
const refs = [];
function scan(value) {
  if (!value || typeof value !== 'object') return;
  if ('$ref' in value) refs.push(value.$ref);
  for (const v of Object.values(value)) scan(v);
}
scan(parsed);
if (refs.some(ref => !ref.startsWith('#/'))) throw new Error('Self-contained contract required; external refs prohibited by check harness');
const config = await core.createConfig({}, { extends: ['recommended'] });
const lint = await core.lint({ ref: source, config });
const bundled = await core.bundle({ ref: source, config, dereference: true });
const clean = p => ({ ruleId: p.ruleId, severity: p.severity, message: p.message, locations: p.location?.map(x => ({ pointer: x.pointer, source: x.source?.absoluteRef })) });
fs.writeFileSync(path.join(here, 'contract-parsed.json'), JSON.stringify(parsed, null, 2) + '\n');
fs.writeFileSync(path.join(here, 'contract-dereferenced.json'), JSON.stringify(bundled.bundle.parsed, null, 2) + '\n');
const methods = new Set(['get','post','put','patch','delete','options','head','trace']);
const operations = Object.entries(parsed.paths).flatMap(([p,item]) => Object.entries(item).filter(([m])=>methods.has(m)).map(([m,op])=>({method:m.toUpperCase(),path:p,operationId:op.operationId})));
const expected = ['POST /auth/google/attempts','GET /auth/google/callback','GET /auth/session','POST /auth/refresh','POST /auth/logout','POST /organizations','GET /organization-creation-requests/{requestId}','GET /organizations/{organizationId}','GET /organizations','POST /account-deletion-previews','POST /account-deletions'].sort();
const actual = operations.map(op=>`${op.method} ${op.path}`).sort();
if (JSON.stringify(actual)!==JSON.stringify(expected)) throw new Error('11-operation inventory mismatch');
const report = { contractSha256: crypto.createHash('sha256').update(raw).digest('hex'), node: process.version, versions: {redocly: require('@redocly/openapi-core/package.json').version, openapiTypescript: JSON.parse(fs.readFileSync(path.join(project,'frontend/node_modules/openapi-typescript/package.json'))).version, typescript:ts.version}, oasVersion:parsed.openapi, refCount: refs.length, externalRefCount:0, operations, lint:lint.map(clean), bundleProblems:bundled.problems.map(clean), limitations:['Static documentation/tooling checks only; no backend, browser, service, OAuth, DB, journal or restore conformance executed.','Installed Redocly core1.34.19 used directly; canonical api tooling2.47.0 is not installed and was not run.'] };
if (parsed.openapi !== '3.0.3') throw new Error('Expected OpenAPI3.0.3');
fs.writeFileSync(path.join(here,'tooling-checks.json'),JSON.stringify(report,null,2)+'\n');
if ([...lint,...bundled.problems].some(p=>p.severity==='error')) throw new Error('Lint or ref-resolution errors; see tooling-checks.json');
const ast = await openapiTS(parsed, { redocly: config, silent:true, cwd:pathToFileURL(here+'/') });
fs.writeFileSync(path.join(here,'generated-api.d.ts'),astToString(ast));
let smoke = 'import type { paths, operations } from "./generated-api";\n' + operations.map((op,i)=>`type Operation${i} = operations[${JSON.stringify(op.operationId)}];\ntype Path${i} = paths[${JSON.stringify(op.path)}][${JSON.stringify(op.method.toLowerCase())}];`).join('\n') + '\n';
let typedExamples=0;
const examplesOf = media => [...('example' in media ? [media.example] : []), ...Object.values(media.examples??{}).filter(ex=>'value' in ex).map(ex=>ex.value)];
for (const op of operations) {
  const definition=bundled.bundle.parsed.paths[op.path][op.method.toLowerCase()];
  for (const [contentType,media] of Object.entries(definition.requestBody?.content??{})) {
    for (const example of examplesOf(media)) smoke+=`const requestExample${typedExamples++}: NonNullable<operations[${JSON.stringify(op.operationId)}]["requestBody"]>["content"][${JSON.stringify(contentType)}] = ${JSON.stringify(example)};\n`;
  }
  for (const [status,response] of Object.entries(definition.responses)) {
    for (const [contentType,media] of Object.entries(response.content??{})) {
      for (const example of examplesOf(media)) smoke+=`const responseExample${typedExamples++}: operations[${JSON.stringify(op.operationId)}]["responses"][${/^\d+$/.test(status)?status:JSON.stringify(status)}]["content"][${JSON.stringify(contentType)}] = ${JSON.stringify(example)};\n`;
    }
  }
}
smoke+='export {};\n';
fs.writeFileSync(path.join(here,'generation-smoke.ts'),smoke);
const program = ts.createProgram([path.join(here,'generation-smoke.ts')],{strict:true,noEmit:true,skipLibCheck:false,types:[],lib:['lib.es2022.d.ts'],module:ts.ModuleKind.NodeNext,moduleResolution:ts.ModuleResolutionKind.NodeNext,incremental:false});
const diagnostics = ts.getPreEmitDiagnostics(program);
report.typescriptGeneration = { operationCount:operations.length, typedExamples, diagnostics: diagnostics.map(d=>ts.flattenDiagnosticMessageText(d.messageText,'\n')), status:diagnostics.length?'FAIL':'PASS' };
fs.writeFileSync(path.join(here,'tooling-checks.json'),JSON.stringify(report,null,2)+'\n');
if (typedExamples < 7) throw new Error('Insufficient embedded examples for meaningful typed payload smoke');
if (diagnostics.length) throw new Error('TypeScript smoke failed');
console.log(JSON.stringify({contractSha256:report.contractSha256,operations:operations.length,refs:refs.length,lintWarnings:lint.filter(p=>p.severity!=='error').length,typescript:'PASS'}));
