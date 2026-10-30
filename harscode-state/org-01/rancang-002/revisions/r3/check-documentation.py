"""Coordinator source/link/scope checks. No application or runtime checks."""
from pathlib import Path
import hashlib, json, re, subprocess

here=Path(__file__).resolve().parent
project=here.parents[4]
workspace=project.parent
manifest=json.loads((here/'source-revisions.json').read_text())
sha=lambda b:hashlib.sha256(b).hexdigest()
sources=[]
for name,revision in manifest['sources'].items():
    actual=sha((project/name).read_bytes())
    historical=False
    if actual!=revision['sha256'] and name=='harscode-state/org-01/rancang-002/current.md':
        original=subprocess.check_output(['git','-C',str(project),'show',manifest['projectCommit']+':'+name])
        historical=sha(original)==revision['sha256']
    sources.append({'path':name,'match':actual==revision['sha256'],'historicalNavigationMatch':historical})
for name,expected in manifest['harscodeFiles'].items():
    sources.append({'path':'harscode-workspace/'+name,'match':sha((workspace/'harscode-workspace'/name).read_bytes())==expected})
for name,expected in manifest['toolingReadOnlyFiles'].items():
    sources.append({'path':name,'match':sha((project/name).read_bytes())==expected})
history=[]
assignment=here.parent.parent
for revision,root,commit in [('r1',assignment,'35f8c0fc31f1f2f21bdac36d4c24d521541263f0'),('r2',here.parent/'r2',manifest['projectCommit'])]:
    previous=json.loads((root/'output-revisions.json').read_text())
    for name,expected in previous['files'].items():
        same=sha((project/name).read_bytes())==expected
        historical=False
        if not same and name in ('docs/engineering/README.md','harscode-state/org-01/rancang-002/current.md'):
            original=subprocess.check_output(['git','-C',str(project),'show',commit+':'+name])
            historical=sha(original)==expected
        history.append({'revision':revision,'path':name,'match':same,'historicalNavigationMatch':historical})
links=[]
for file in sorted(here.glob('*.md'))+[assignment/'current.md']:
    for target in re.findall(r'\[[^\]]*\]\(([^)]+)\)',file.read_text()):
        target=target.strip().strip('<>')
        if target.startswith(('https:','http:','mailto:','#')): continue
        path_part=target.split('#',1)[0]
        if not path_part: continue
        resolved=(file.parent/path_part).resolve()
        links.append({'source':file.name,'target':target,'exists':resolved.exists()})
status=subprocess.check_output(['git','-C',str(project),'status','--porcelain','--untracked-files=all'],text=True)
changes=[line[3:] for line in status.splitlines()]
prefix='harscode-state/org-01/rancang-002/revisions/r3/'
allowed=lambda p:p.startswith(prefix) or p=='harscode-state/org-01/rancang-002/current.md'
scope_ok=all(allowed(p) for p in changes)
report={'sourceChecks':sources,'sourceMatchCount':sum(x['match'] for x in sources),'historicalNavigationChecks':sum(x.get('historicalNavigationMatch',False) for x in sources),'historicalBundleChecks':history,'localLinks':links,'localLinkCount':len(links),'localLinkFailures':[x for x in links if not x['exists']],'changedPaths':changes,'scopeCheck':scope_ok,'limits':['Link existence only, no automatic semantic or section-anchor assurance.','Git/source checks document observed writes; sandbox is broader than authorized r3 scope.','No application/runtime/provider/legal/email/recovery checks.']}
(here/'documentation-checks.json').write_text(json.dumps(report,indent=2)+'\n')
fail=any(not x['match'] and not x.get('historicalNavigationMatch') for x in sources+history) or any(not x['exists'] for x in links) or not scope_ok
print(json.dumps({'sources':len(sources),'sourceFailures':sum(not x['match'] and not x.get('historicalNavigationMatch',False) for x in sources),'links':len(links),'linkFailures':len(report['localLinkFailures']),'scope':scope_ok}))
if fail: raise SystemExit(1)
