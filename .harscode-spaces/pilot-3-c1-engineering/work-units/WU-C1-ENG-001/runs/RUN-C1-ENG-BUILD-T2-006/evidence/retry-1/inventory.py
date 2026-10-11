import sys,json,pathlib,hashlib,re,zipfile,base64
sys.path.insert(0,str(pathlib.Path(__file__).parent)); from run import cap,base,env

def stream(s):
 out=[];dec=json.JSONDecoder()
 while s.strip():o,n=dec.raw_decode(s.lstrip());out.append(o);s=s.lstrip()[n:]
 return out

def h1(files):
 txt=''.join(hashlib.sha256(data).hexdigest()+'  '+name+'\n' for name,data in sorted(files))
 return 'h1:'+base64.b64encode(hashlib.sha256(txt.encode()).digest()).decode()
mods=stream((cap/'graph.stdout').read_text()); sums={tuple(l.split()[:2]):l.split()[2] for l in (base/'probe/go.sum').read_text().splitlines()};out=[]
for m in mods:
 if m.get('Main'):continue
 p=m['Path'];v=m['Version'];c=pathlib.Path(env['GOMODCACHE'])/'cache/download'/p/'@v';files={str(c/(v+x)):hashlib.sha256((c/(v+x)).read_bytes()).hexdigest() for x in ['.info','.mod','.zip','.ziphash']};info=json.loads((c/(v+'.info')).read_text());mod=(c/(v+'.mod')).read_bytes(); g=re.search(rb'^go\s+(\S+)',mod,re.M)
 z=zipfile.ZipFile(c/(v+'.zip')); archive=h1([(n,z.read(n)) for n in z.namelist() if not n.endswith('/')]); ziphash=(c/(v+'.ziphash')).read_text().strip(); assert archive==ziphash,(p,'archive')
 for kind,val in [(v,archive),(v+'/go.mod',h1([('go.mod',mod)]))]:
  if (p,kind) in sums: assert val==sums[p,kind],(p,kind,'go.sum mismatch')
 source=pathlib.Path(env['GOMODCACHE'])/(p+'@'+v);assert source.is_dir(),p
 out.append({'path':p,'version':v,'go':g[1].decode() if g else 'no go directive (legacy)','origin':info.get('Origin'),'time':info.get('Time'),'files':files,'source':str(source),'sum':archive,'go_mod_sum':h1([('go.mod',mod)]),'ziphash':ziphash,'archive_go_sum_present':(p,v) in sums,'mod_go_sum_present':(p,v+'/go.mod') in sums})
(cap/'inventory.json').write_text(json.dumps(out,indent=2)); baseline={m['Path']:m.get('Version') for m in stream((cap/'baseline.stdout').read_text())};delta=[{'module':m['path'],'baseline':baseline.get(m['path']),'selected':m['version']} for m in out if baseline.get(m['path'])!=m['version']];(cap/'delta.json').write_text(json.dumps(delta,indent=2));print('20 cache/source inputs inventoried; ziphash and available go.sum checks match.'); print('Missing go.sum archive entries:',[m['path'] for m in out if not m['archive_go_sum_present']]); print(json.dumps(delta))
