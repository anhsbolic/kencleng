import pathlib,json,sys,collections,re
sys.path.insert(0,str(pathlib.Path(__file__).parent));from run import cap
from inventory import stream
raw=stream((cap/'source-scan.stdout').read_text()); count=collections.Counter(k for o in raw for k in o); assert count['config']==1 and count['SBOM']==1,(count,'scan incomplete'); assert not (cap/'source-scan.stderr').read_text().strip(); findings=[o['finding'] for o in raw if 'finding' in o]; assert not findings,findings
mods=json.loads((cap/'inventory.json').read_text());versions={m['path']:m['version'] for m in mods};versions['stdlib']='v1.26.9';versions['toolchain']='v1.26.9'
osv={o['osv']['id']:o['osv'] for o in raw if 'osv' in o}
root=pathlib.Path.cwd();prior=root/'.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-001/evidence/dependency-vuln.json';before=stream(prior.read_text());required={o['finding']['osv'] for o in before if 'finding' in o};required|={'GO-2026-4771','GO-2026-4772','GO-2026-5004','GO-2026-5970','GO-2026-6629'};assert required<=osv.keys(),('missing fresh advisory metadata',required-osv.keys())
def ver(v):
 m=re.fullmatch(r'v?(\d+)\.(\d+)\.(\d+)(?:-([^+]+))?(?:\+.*)?',v);assert m,v
 return tuple(int(m[i]) for i in range(1,4))+(0 if m[4] else 1,m[4] or '')
rows=[]
for id in sorted(required):
 affects=[]
 for a in osv[id]['affected']:
  p=a['package']['name'];sel=versions.get(p);ranges=[];bad=False
  for r in a.get('ranges',[]):
   if r['type']!='SEMVER':continue
   start=None
   for e in r['events']:
    if 'introduced' in e:start=e['introduced']
    elif 'fixed' in e:
     end=e['fixed'];ranges.append({'introduced':start,'fixed':end});bad|=bool(sel and (start=='0' or ver(sel)>=ver(start)) and ver(sel)<ver(end));start=None
    elif 'last_affected' in e:
     end=e['last_affected'];ranges.append({'introduced':start,'last_affected':end});bad|=bool(sel and (start=='0' or ver(sel)>=ver(start)) and ver(sel)<=ver(end));start=None
   if start is not None:ranges.append({'introduced':start});bad|=bool(sel and (start=='0' or ver(sel)>=ver(start)))
  assert not bad,(id,p,sel,ranges)
  affects.append({'module':p,'selected':sel,'ranges':ranges,'disposition':'selected outside affected ranges' if sel else 'module not selected'})
 rows.append({'id':id,'fresh_metadata':True,'affected':affects,'finding_records':0,'traces':[],'disposition':'fixed selection or absent module; no source finding; no waiver'})
(cap/'advisory-dispositions.json').write_text(json.dumps(rows,indent=2)); summary={'counts':dict(count),'finding_count':len(findings),'required_advisory_ids':len(required),'config':raw[0].get('config'),'SBOM':next(o['SBOM'] for o in raw if 'SBOM' in o)};(cap/'scan-summary.json').write_text(json.dumps(summary,indent=2));print(json.dumps({'counts':dict(count),'required_ids_reconciled':len(required),'findings':len(findings)}))
# Authenticated targeted download checks, independent of go.sum entries for unused packages.
for d in stream((cap/'missing-sums.stdout').read_text()):
 assert 'Error' not in d,d
 m=next(x for x in mods if x['path']==d['Path']);assert d['Version']==m['version'];assert d['Sum']==m['sum'];assert d['GoModSum']==m['go_mod_sum']
print('Targeted download metadata checksums match cache hashes for six modules.')
