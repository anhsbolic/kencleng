import json,pathlib,sys,subprocess,datetime,re
sys.path.insert(0,str(pathlib.Path(__file__).parent));from run import cap,run
mods=json.loads((cap/'inventory.json').read_text()); results=[]
for m in mods:
 origin=m['origin']
 if not origin:continue
 p=m['path'];name='ref-'+p.replace('/','_')
 if origin.get('Ref'):
  txt=run(name,['git','ls-remote',origin['URL'],origin['Ref'],origin['Ref']+'^{}']);refs=[l.split() for l in txt.splitlines()];assert refs,(p,'missing upstream ref');commit=next((h for h,r in refs if r.endswith('^{}')),refs[0][0]);assert commit==origin['Hash'],(p,'Origin does not match upstream',commit,origin['Hash']); results.append({'module':p,'commit':commit,'upstream_ref':origin['Ref'],'matches_origin':True})
for p,repo,sha in [('github.com/jackc/pgservicefile','jackc/pgservicefile','5a60cdf6a76120dc3d5152b95f3b5fd8aa7cc9eb'),('gopkg.in/check.v1','go-check/check','10cb98267c6cb43ea9cd6793f29ff4089c306974')]:
 txt=run('commit-'+p.replace('/','_'),['curl','--fail','--silent','--show-error','--max-time','30','https://api.github.com/repos/'+repo+'/git/commits/'+sha]);j=json.loads(txt);assert j['sha']==sha;date=j['committer']['date'];m=next(x for x in mods if x['path']==p);v=m['version'];stamp=v.split('-')[-2];assert date.replace('-','').replace(':','').replace('T','').replace('Z','')==stamp,(p,date,stamp);results.append({'module':p,'commit':sha,'committer_date':date,'pseudo_version_timestamp_matches':True})
 (cap/'upstream-validation.json').write_text(json.dumps(results,indent=2))
print('Upstream refs and pseudo-version commit timestamps match.',flush=True)
