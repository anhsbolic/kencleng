"""Independent adaptation of bound original staged defect probe; safe output only."""
import ast
import contextlib
import hashlib
import io
import json
from pathlib import Path
import subprocess
import sys
import tempfile
from unittest.mock import patch

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[7]
sys.path.insert(0, str(ROOT / "backend/scripts"))
import scanner_gate as gate

ORIGINAL = ROOT / ".harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-SCANNER-001/evidence/native-ignore-probe.py"
ORIGINAL_HASH = "8ef6e4c78b28b3d91c1d728b2582f0363b33aab702194d84d5df3f523512bc97"


def require(condition, category):
    if not condition:
        raise RuntimeError(category)


def main():
    original = ORIGINAL.read_bytes()
    require(hashlib.sha256(original).hexdigest() == ORIGINAL_HASH, "bound original probe drift")
    # Read the same synthetic staged text privately from immutable source.
    tree = ast.parse(original)
    writes = sorted((n for n in ast.walk(tree) if isinstance(n, ast.Call)
                     and isinstance(n.func, ast.Attribute) and n.func.attr == "write_text"),
                    key=lambda n: n.lineno)
    fixture_text = ast.literal_eval(writes[1].args[0])
    with tempfile.TemporaryDirectory(prefix="c1-confirm-ignore-", dir="/tmp") as folder:
        root = Path(folder); backend = root / "backend"; backend.mkdir()
        def git(*args):
            result = subprocess.run(["git", "-C", folder, *args], capture_output=True,
                                    timeout=30, check=False)
            require(result.returncode == 0, "fixture Git failure")
            return result.stdout
        git("init", "-q")
        source = backend / "tracked.txt"
        source.write_text("clean fixture\n"); git("add", "backend")
        git("-c", "user.name=Fixture", "-c", "user.email=fixture@example.invalid",
            "commit", "-qm", "synthetic fixture")
        source.write_text(fixture_text); git("add", "backend/tracked.txt")
        source_state = (hashlib.sha256(source.read_bytes()).digest(),
                        hashlib.sha256(git("show", ":backend/tracked.txt")).digest())
        def unchanged():
            require(source_state == (hashlib.sha256(source.read_bytes()).digest(),
                                     hashlib.sha256(git("show", ":backend/tracked.txt")).digest()),
                    "synthetic source/index changed")
        image = "ghcr.io/gitleaks/gitleaks:v8.30.1"
        prefix = ["podman", "run", "--rm", "--network=none", "--volume", folder + ":/repo:ro",
                  "--workdir", "/repo", image]
        version = gate.run(prefix + ["version"], cwd=backend)
        if version.returncode:
            category = ("podman runtime filesystem restriction" if b"read-only file system" in version.stderr
                        else "configured tool preflight failure")
            print(json.dumps({"stage": "tool-preflight", "raw_exit": version.returncode,
                              "category": category, "scans_executed": False}))
            return 2
        require(version.stdout.strip() == b"v8.30.1", "unsupported tool version")
        command = gate.leaks_command(prefix, "git", "/repo/backend", "--staged")
        before = gate.run(command, cwd=backend); findings = gate.report(before.stdout)
        require(before.returncode == 10 and len(findings) == 1, "original unknown finding absent")
        entry = backend / ".gitleaksignore"
        original_run = gate.run
        calls = []
        def observed_run(command, **kwargs):
            calls.append(command)
            kwargs.setdefault("cwd", backend)
            return original_run(command, **kwargs)
        def full_gate(configuration=False):
            calls.clear(); output = io.StringIO(); failed = False; category = None
            with patch.object(gate, "ROOT", root), patch.object(gate, "BACKEND", backend), \
                 patch.object(gate, "run", observed_run), \
                 patch.object(gate.sys, "argv", ["gate", "gitleaks-worktree", "--", *prefix]), \
                 contextlib.redirect_stdout(output), contextlib.redirect_stderr(output):
                status = gate.main()
            text = output.getvalue()
            require(status == 1 and "effective=FAIL" in text, "unknown finding produced false PASS")
            require(all(value not in text for value in (fixture_text.strip(), findings[0]["Fingerprint"],
                                                        "Fingerprint", "Secret", "Match")), "private output escaped")
            if configuration:
                require(not calls and "scans not executed" in text and "PASS" not in text,
                        "configuration rejection not before tool work")
                category = "unreviewed native Gitleaks policy entry; scans not executed"
            return {"cli_exit": status, "effective": "REJECTED_CONFIGURATION" if configuration else "FAIL",
                    "process_calls": len(calls), "scanner_scans": sum("--report-format" in c for c in calls),
                    "safe_category": category, "payload_output": False}
        no_ignore = full_gate(); unchanged()
        entry.write_text(findings[0]["Fingerprint"] + "\n")
        entry_bytes = entry.read_bytes()
        after = gate.run(command, cwd=backend)
        require(after.returncode == 0 and gate.report(after.stdout) == [], "native suppression oracle changed")
        unchanged()
        rejection = full_gate(configuration=True)
        require(entry.read_bytes() == entry_bytes, "test-owned ignore artifact changed")
        unchanged()
        entry.unlink()  # Only disposable test-owned state.
        restored = full_gate(); unchanged()
        result = {"result": "PASS", "tool": "Gitleaks v8.30.1", "image": image,
                  "network": "none", "mount": "read-only disposable fixture", "synthetic_only": True,
                  "original_probe_sha256": ORIGINAL_HASH,
                  "before_ignore": {"raw_exit": before.returncode, "findings": len(findings), "gate": no_ignore},
                  "raw_native_suppression": {"raw_exit": after.returncode, "findings": 0},
                  "fixed_gate": rejection, "artifact_bytes_preserved": True,
                  "after_removal": restored, "source_and_index_unchanged": True,
                  "cleanup": "TemporaryDirectory"}
    print(json.dumps(result, sort_keys=True))
    return 0


if __name__ == "__main__":
    try:
        sys.exit(main())
    except (RuntimeError, gate.GateError, OSError, ValueError, TypeError, KeyError,
            IndexError, subprocess.SubprocessError):
        print("confirmation: FAIL (safe fixture/tool/assertion category)", file=sys.stderr)
        sys.exit(2)
