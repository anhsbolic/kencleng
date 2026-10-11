import pathlib,shutil,sys,json
sys.path.insert(0,str(pathlib.Path(__file__).parent));from run import base,go
from actual import actual
import run
reg=base/'regression-control';shutil.copytree(base/'probe',reg)
text=(pathlib.Path.cwd()/'backend/internal/platform/auth/store_test.go').read_text();types=text[text.index('type sessionRow'):text.index('// The Scan barrier')];test=text[text.index('func TestLockSessionIgnoresStaleCallerTime'):text.index('func TestMalformedBrowserAuthorityRejectedBeforeStorage')]
(reg/'internal/platform/auth/regression_test.go').write_text('package auth\nimport ("bytes";"context";"encoding/base64";"errors";"testing";"time";"github.com/jackc/pgx/v5")\n'+types+test)
try:run.run('regression-original',[go,'test','-count=1','-run','^TestLockSessionIgnoresStaleCallerTime$','-v','./internal/platform/auth'],reg)
except SystemExit as e:assert e.code==1
else:raise AssertionError('Original implementation did not fail regression')
stdout=(run.cap/'regression-original.stdout').read_text();assert 'public lock API accepted expiry after stale caller time' in stdout
shutil.copy2(pathlib.Path.cwd()/'backend/internal/platform/auth/store.go',reg/'internal/platform/auth/store.go')
run.run('regression-patched',[go,'test','-count=1','-run','^TestLockSessionIgnoresStaleCallerTime$','-v','./internal/platform/auth'],reg)
(run.cap/'regression-control.json').write_text(json.dumps({'original_exit':1,'original_failure_expected':True,'patched_exit':0,'purpose':'same public-API regression fails on original source and passes on patched source','real_postgres_contention':False},indent=2));print('Expected original failure and patched PASS confirmed.')
