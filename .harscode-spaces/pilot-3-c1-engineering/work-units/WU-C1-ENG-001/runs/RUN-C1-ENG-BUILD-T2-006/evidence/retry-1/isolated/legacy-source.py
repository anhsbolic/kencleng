import pathlib,json,subprocess,tarfile,zipfile,hashlib,sys,datetime
sys.path.insert(0,str(pathlib.Path(__file__).parent)); from run import cap
mods={m['path']:m for m in json.loads((cap/'inventory.json').read_text())}; results=[]
for p in json.loads((cap/'legacy-provenance.json').read_text()):
 refs=[l.split() for l in p['refs'].splitlines()]; sha=next((h for h,r in refs if r.endswith('^{}')),refs[0][0])
 if p['module']=='gopkg.in/check.v1': sha=next(h for h,r in refs if h.startswith('10cb98267c6c'))
 repo=p['repo'].removeprefix('https://github.com/'); dest=cap/('upstream-'+p['module'].replace('/','_')+'.tar.gz'); url='https://codeload.github.com/'+repo+'/tar.gz/'+sha; args=['curl','--fail','--silent','--show-error','--max-time','30','--output',str(dest),url]; r=subprocess.run(args,capture_output=True,text=True)
 (cap/(dest.stem+'.stderr')).write_text(r.stderr)
 with (cap/'commands.jsonl').open('a') as f:f.write(json.dumps({'name':'upstream-source','argv':args,'cwd':str(pathlib.Path.cwd()),'exit':r.returncode,'network':'harness escalation','time':datetime.datetime.now(datetime.timezone.utc).isoformat()})+'\n')
 if r.returncode: print('FAIL',p['module'],r.stderr,flush=True); sys.exit(1)
 tar=tarfile.open(dest); entries={x.name.split('/',1)[1]:x for x in tar.getmembers() if x.isfile() and '/' in x.name}; m=mods[p['module']]; zip_path=next(x for x in m['files'] if x.endswith('.zip')); z=zipfile.ZipFile(zip_path); count=0; mismatch=[]
 for name in z.namelist():
  rel=name.split('@'+m['version']+'/',1)[1]
  if name.endswith('/'): continue
  if rel not in entries or z.read(name)!=tar.extractfile(entries[rel]).read(): mismatch.append(rel)
  count+=1
 result={'module':p['module'],'version':m['version'],'repo':p['repo'],'commit':sha,'archive_sha256':hashlib.sha256(dest.read_bytes()).hexdigest(),'compared_files':count,'mismatch':mismatch}; results.append(result); (cap/'legacy-source-match.json').write_text(json.dumps(results,indent=2)); print(json.dumps(result),flush=True)
 if mismatch: sys.exit(2)
