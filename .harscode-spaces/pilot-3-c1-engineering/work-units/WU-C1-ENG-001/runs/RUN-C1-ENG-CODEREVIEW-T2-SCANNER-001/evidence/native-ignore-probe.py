import contextlib, io, json, pathlib, subprocess, sys, tempfile
from unittest.mock import patch
sys.dont_write_bytecode = True
sys.path.insert(0, '/home/anhar-solehudin/kencleng-workspace/kencleng/backend/scripts')
import scanner_gate as gate

with tempfile.TemporaryDirectory(prefix='c1-review-ignore-', dir='/tmp') as folder:
    root = pathlib.Path(folder); backend = root / 'backend'; backend.mkdir()
    def git(*args):
        r = subprocess.run(['git', '-C', folder, *args], capture_output=True, timeout=30)
        if r.returncode: raise RuntimeError('fixture Git failure')
    git('init', '-q')
    (backend / 'tracked.txt').write_text('clean fixture\n')
    git('add', 'backend')
    git('-c', 'user.name=Scanner fixture', '-c', 'user.email=scanner@example.invalid', 'commit', '-qm', 'baseline')
    (backend / 'tracked.txt').write_text('api_key = "c1SYNTHETIC7uP4qL9wE2rT8yA5sD3fG6hJ0kZ"\n')
    git('add', 'backend/tracked.txt')
    prefix = ['podman', 'run', '--rm', '--network=none', '--volume', folder + ':/repo:ro', '--workdir', '/repo', 'ghcr.io/gitleaks/gitleaks:v8.30.1']
    version = gate.run(prefix + ['version'])
    if version.returncode:
        print('PREFLIGHT_FAILURE', version.returncode)
        # Podman preflight has no scanner payload or credentials.
        print(version.stderr.decode(errors='replace'))
        sys.exit(2)
    if version.stdout.strip() != b'v8.30.1': raise RuntimeError('unsupported version')
    command = gate.leaks_command(prefix, 'git', '/repo/backend', '--staged')
    before = gate.run(command); findings = json.loads(before.stdout)
    assert before.returncode == 10 and len(findings) == 1
    (backend / '.gitleaksignore').write_text(findings[0]['Fingerprint'] + '\n')
    after = gate.run(command)
    assert after.returncode == 0 and json.loads(after.stdout) == []
    captured = io.StringIO()
    with patch.object(gate, 'ROOT', root), patch.object(gate, 'BACKEND', backend), contextlib.redirect_stdout(captured), contextlib.redirect_stderr(captured):
        gate.scan('gitleaks-worktree', prefix)
    assert 'effective=FAIL' not in captured.getvalue()
    print(json.dumps({'tool':'Gitleaks v8.30.1', 'image':'ghcr.io/gitleaks/gitleaks:v8.30.1', 'network':'none', 'synthetic_only':True, 'before_ignore':{'raw_exit':before.returncode,'findings':len(findings),'effective':'FAIL'}, 'after_ignore':{'raw_exit':after.returncode,'findings':0,'effective':'PASS'}, 'full_worktree_gate':'PASS despite staged synthetic finding', 'scan_attempts':captured.getvalue().count('gitleaks-worktree['), 'cleanup':'TemporaryDirectory'}, sort_keys=True))
