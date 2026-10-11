import sys,pathlib
sys.path.insert(0,str(pathlib.Path(__file__).parent));import run
run.cap=run.base/'actual';run.cap.mkdir(exist_ok=True)
def actual(name,args):return run.run(name,args,pathlib.Path.cwd()/'backend')
if __name__=='__main__':
 actual('pins',[run.go,'mod','edit','-require=github.com/jackc/pgx/v5@v5.9.2','-require=golang.org/x/text@v0.41.0']);actual('tidy',[run.go,'mod','tidy']);actual('graph',[run.go,'list','-m','-json','all']);actual('edges',[run.go,'mod','graph']);actual('integrity',[run.go,'mod','verify']);actual('focused-compile',[run.go,'test','-run','^$','./internal/platform/auth','./internal/platform/db','./cmd/server'])
