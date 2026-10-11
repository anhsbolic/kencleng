"""Bounded, secret-free command captures for independent T2 Testing."""
import datetime
import hashlib
import json
import os
from pathlib import Path
import subprocess
import sys
import tempfile

EVIDENCE = Path(__file__).resolve().parent
ROOT = EVIDENCE.parents[6]
GO = '/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go'
SCANNER = '/home/anhar-solehudin/.gvm/pkgsets/go1.26.9/global/bin/govulncheck'
CONFIG = EVIDENCE / 'capture-context.json'
if not CONFIG.exists():
    cache = tempfile.mkdtemp(prefix='c1-test-t2-001-', dir='/tmp')
    CONFIG.write_text(json.dumps({'build_cache': cache}, indent=2) + '\n')
cache = json.loads(CONFIG.read_text())['build_cache']
env = {key: os.environ[key] for key in ['HOME', 'USER', 'LANG', 'TMPDIR'] if key in os.environ}
env.update(PATH=os.environ['PATH'], GOTOOLCHAIN='local', GOCACHE=cache,
           GOMODCACHE='/home/anhar-solehudin/.gvm/pkgsets/go1.24.8/global/pkg/mod',
           GOPROXY='off', GOSUMDB='sum.golang.org', GOFLAGS='-mod=readonly')
commands = {
    'go-version': [GO, 'version'],
    'go-env': [GO, 'env', 'GOMOD', 'GOTOOLCHAIN', 'GOROOT', 'GOPATH', 'GOCACHE', 'GOMODCACHE', 'GOPROXY', 'GOSUMDB'],
    'scanner-version': [GO, 'version', '-m', SCANNER],
    'staticcheck-version': ['staticcheck', '-version'],
    'gosec-version': ['gosec', '-version'],
    'podman-version': ['podman', '--version'],
    'podman-version-recovery': ['podman', '--version'],
    'mod-verify': [GO, 'mod', 'verify'],
    'make-verify': ['make', 'verify'],
    'make-verify-recovery': ['make', 'verify'],
    'unit': [GO, 'test', '-count=1', '-timeout=60s', '-v', './...'],
    'race': [GO, 'test', '-count=1', '-timeout=60s', '-race', './...'],
    'contract': [GO, 'test', '-count=1', '-timeout=60s', '-tags=contract', './...'],
    'source-scan': [SCANNER, '-json', './...'],
}
name = sys.argv[1]
if name == 'make-verify-recovery':
    env['XDG_CACHE_HOME'] = cache + '/tool-cache'
    env['STATICCHECK_CACHE'] = cache + '/staticcheck'
if (EVIDENCE / (name + '.stdout')).exists():
    raise SystemExit('Capture exists; use a separately named recovery attempt.')
start = datetime.datetime.now(datetime.timezone.utc).isoformat()
with (EVIDENCE / (name + '.stdout')).open('wb') as out, (EVIDENCE / (name + '.stderr')).open('wb') as err:
    try:
        result = subprocess.run(commands[name], cwd=ROOT / 'backend', env=env,
                                stdout=out, stderr=err, timeout=180)
        code = result.returncode
    except subprocess.TimeoutExpired:
        code = 'TIMEOUT'
record = dict(name=name, argv=commands[name], cwd=str(ROOT / 'backend'),
              env={k: env[k] for k in ['GOTOOLCHAIN', 'GOCACHE', 'GOMODCACHE', 'GOPROXY', 'GOSUMDB', 'GOFLAGS']},
              started=start, ended=datetime.datetime.now(datetime.timezone.utc).isoformat(), exit=code)
record['tool_cache_overrides'] = {k: env[k] for k in ['XDG_CACHE_HOME', 'STATICCHECK_CACHE'] if k in env}
for stream in ['stdout', 'stderr']:
    record[stream + '_sha256'] = hashlib.sha256((EVIDENCE / (name + '.' + stream)).read_bytes()).hexdigest()
with (EVIDENCE / 'commands.jsonl').open('a') as log:
    log.write(json.dumps(record) + '\n')
print(json.dumps(record))
print((EVIDENCE / (name + '.stdout')).read_text()[-12000:])
print((EVIDENCE / (name + '.stderr')).read_text()[-12000:])
