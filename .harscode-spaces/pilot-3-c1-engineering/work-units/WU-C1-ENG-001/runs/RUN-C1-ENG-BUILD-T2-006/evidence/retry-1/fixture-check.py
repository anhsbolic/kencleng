import sys,pathlib,json,subprocess,time
sys.path.insert(0,str(pathlib.Path(__file__).parent));from actual import actual;from run import go,env,cap
port=actual('fixture-port',['podman','port','c1-t2-006-retry1-pg','5432/tcp']).strip().rsplit(':',1)[1]
for i in range(15):
 r=subprocess.run(['podman','exec','c1-t2-006-retry1-pg','pg_isready','-U','c1_fixture','-d','c1_fixture'],capture_output=True,text=True)
 if r.returncode==0:break
 time.sleep(1)
else:raise SystemExit('Disposable fixture never became ready')
actual('fixture-version',['podman','exec','c1-t2-006-retry1-pg','psql','-U','c1_fixture','-d','c1_fixture','-Atc','SELECT version()'])
env['C1_AUTH_TEST_DATABASE_URL']='postgres://c1_fixture@127.0.0.1:'+port+'/c1_fixture?sslmode=disable'
actual('focused-postgres',[go,'test','-count=1','-tags=integration','-run','^(TestPersonUpsertPostgreSQL|TestPendingLoginBindingAndReplayPostgreSQL)$','-v','./internal/platform/auth'])
(cap/'fixture-context.json').write_text(json.dumps({'container':'c1-t2-006-retry1-pg','image':'docker.io/library/postgres:16-alpine','host':'127.0.0.1','port':int(port),'disposable':True,'authentication':'fixture trust only','production_env_read':False,'privilege_evidence':False,'contention_evidence':False},indent=2))
