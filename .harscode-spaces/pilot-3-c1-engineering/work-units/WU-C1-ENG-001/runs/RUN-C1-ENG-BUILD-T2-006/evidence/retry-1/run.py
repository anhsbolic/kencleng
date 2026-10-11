import pathlib,os,subprocess,json,datetime,sys
base=pathlib.Path(__file__).parent; cap=base/'capture'; go='/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go'; scanner='/home/anhar-solehudin/.gvm/pkgsets/go1.26.9/global/bin/govulncheck'
env=os.environ.copy(); env.update(PATH=str(pathlib.Path(go).parent)+':'+env['PATH'],GOTOOLCHAIN='local',GOCACHE=str(base/'build-cache'),GOMODCACHE='/home/anhar-solehudin/.gvm/pkgsets/go1.24.8/global/pkg/mod',GOPROXY='off')
def run(name,args,cwd=None):
 cwd=str(cwd or base/'probe'); r=subprocess.run(args,cwd=cwd,env=env,capture_output=True,text=True); (cap/(name+'.stdout')).write_text(r.stdout); (cap/(name+'.stderr')).write_text(r.stderr)
 with (cap/'commands.jsonl').open('a') as f: f.write(json.dumps({'name':name,'argv':args,'cwd':cwd,'env':{k:env[k] for k in ['PATH','GOTOOLCHAIN','GOCACHE','GOMODCACHE','GOPROXY']},'exit':r.returncode,'time':datetime.datetime.now(datetime.timezone.utc).isoformat()})+'\n')
 print(name,'exit',r.returncode,flush=True)
 if r.returncode: print(r.stderr); sys.exit(r.returncode)
 return r.stdout
if __name__=='__main__':
 for name,args in [('go-version',[go,'version']),('go-env',[go,'env','GOTOOLCHAIN','GOROOT','GOPATH','GOMOD','GOCACHE','GOMODCACHE']),('scanner-version',[go,'version','-m',scanner]),('baseline',[go,'list','-m','-json','all']),('pins',[go,'mod','edit','-require=github.com/jackc/pgx/v5@v5.9.2','-require=golang.org/x/text@v0.41.0']),('tidy',[go,'mod','tidy']),('graph',[go,'list','-m','-json','all']),('edges',[go,'mod','graph'])]: run(name,args)
