import sys,json,pathlib,subprocess,datetime
sys.path.insert(0,str(pathlib.Path(__file__).parent)); from run import cap
mods=json.loads((cap/'inventory.json').read_text()); out=[]
for m in mods:
 if m['origin']: continue
 p=m['path']; v=m['version']; repo='https://github.com/'+p[len('github.com/'):] if p.startswith('github.com/') else {'gopkg.in/check.v1':'https://github.com/go-check/check','gopkg.in/yaml.v3':'https://github.com/go-yaml/yaml'}[p]
 name='origin-'+p.replace('/','_'); args=['git','ls-remote',repo,'refs/tags/'+v,'refs/tags/'+v+'^{}'] if not p=='gopkg.in/check.v1' else ['git','ls-remote',repo,'refs/heads/*','refs/tags/*']
 r=subprocess.run(args,capture_output=True,text=True,timeout=40); (cap/(name+'.stdout')).write_text(r.stdout); (cap/(name+'.stderr')).write_text(r.stderr)
 with (cap/'commands.jsonl').open('a') as f: f.write(json.dumps({'name':name,'argv':args,'cwd':str(pathlib.Path.cwd()),'exit':r.returncode,'time':datetime.datetime.now(datetime.timezone.utc).isoformat(),'network':'upstream git; harness escalation'})+'\n')
 print(p,r.returncode,r.stdout.strip(),r.stderr.strip(),flush=True)
 out.append({'module':p,'version':v,'repo':repo,'exit':r.returncode,'refs':r.stdout})
 (cap/'legacy-provenance.json').write_text(json.dumps(out,indent=2))
 if r.returncode or not r.stdout.strip(): sys.exit(1)
