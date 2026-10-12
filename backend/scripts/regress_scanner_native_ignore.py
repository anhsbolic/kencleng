"""Actual v8.30.1 regression on private disposable synthetic Git fixtures.

Run from the repository root: python3 backend/scripts/regress_scanner_native_ignore.py
No report payloads, fingerprints, scanner stderr, or fixture identities are emitted.
"""
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
import scanner_gate as gate

IMAGE = "ghcr.io/gitleaks/gitleaks:v8.30.1"


def require(condition, category):
    if not condition:
        raise gate.GateError(category)


def regression(mode):
    with tempfile.TemporaryDirectory(prefix="c1-patch-ignore-", dir="/tmp") as folder:
        root = Path(folder); backend = root / "backend"; backend.mkdir()
        def git(*args):
            result = gate.run(["git", "-C", folder, *args], cwd=root)
            require(result.returncode == 0, "fixture Git failure")
        def commit():
            git("add", "backend")
            git("-c", "user.name=Scanner fixture", "-c", "user.email=scanner@example.invalid",
                "commit", "-qm", "synthetic fixture")
        git("init", "-q")
        prefix = ["podman", "run", "--rm", "--network=none", "--volume", folder + ":/repo:ro",
                  "--workdir", "/repo", IMAGE]
        # Tool discovery stays read-only, offline, and bounded by gate.run timeout.
        version = gate.run(prefix + ["version"], cwd=root)
        if version.returncode:
            category = ("podman runtime filesystem restriction" if b"read-only file system" in version.stderr
                        else "configured tool preflight failure")
            print(json.dumps({"stage": "tool-preflight", "raw_exit": version.returncode,
                              "category": category, "scans_executed": False}))
            raise gate.GateError("configured tool unavailable")
        require(version.stdout.strip() == b"v8.30.1", "unsupported tool version")
        source = backend / "tracked.txt"
        source.write_text("clean fixture\n")
        commit()
        # Only synthetic values; no user repository/configuration/credential reads.
        def value(index):
            return 'api_key = "c1SYNTHETIC' + str(index) + 'uP4qL9wE2rT8yA5sD3fG6hJ0kZ"\n'
        accepted = {}
        accepted_pass = None
        if mode == "history":
            source.write_text("".join(value(i) for i in range(4)))
            commit()
            command = gate.leaks_command(prefix, "git", "/repo", "--log-opts=--all --full-history -- backend/")
            initial = gate.run(command, cwd=root)
            findings = gate.report(initial.stdout)
            require(initial.returncode == 10 and len(findings) == 4, "fixture history candidate setup failure")
            for f in findings:
                blob = gate.run(["git", "show", f'{f["Commit"]}:{f["File"]}'], cwd=root)
                require(blob.returncode == 0, "fixture history blob failure")
                accepted[(f["Commit"], f["File"], f["StartLine"])] = hashlib.sha256(blob.stdout).hexdigest()
            # --all scans every ref. Verify the candidate-only fixture BEFORE
            # committing the extra finding; checking out its parent cannot
            # remove that later commit from full-history coverage.
            output = io.StringIO()
            with patch.object(gate, "ROOT", root), patch.object(gate, "BACKEND", backend), patch.object(gate, "HISTORY", accepted), contextlib.redirect_stdout(output), contextlib.redirect_stderr(output):
                gate.scan("gitleaks-history", prefix)
            require("accepted=4 effective=PASS" in output.getvalue(), "fixture candidate-only gate failure")
            require("SYNTHETIC" not in output.getvalue(), "unsafe output boundary")
            accepted_pass = {"effective": "PASS", "scanner_scans": 1, "payload_output": False}
            source.write_text(source.read_text() + value(4)); commit()
        elif mode == "untracked":
            (backend / "new.txt").write_text(value(4))
            command = gate.leaks_command(prefix, "dir", "/repo/backend/new.txt")
        else:
            source.write_text(value(4))
            if mode == "staged":
                git("add", "backend/tracked.txt")
            command = gate.leaks_command(prefix, "git", "/repo/backend",
                                         "--staged" if mode == "staged" else "--pre-commit")
        scan_mode = "gitleaks-history" if mode == "history" else "gitleaks-worktree"
        original_run = gate.run
        calls = []
        def run(command, **kwargs):
            calls.append(command)
            kwargs.setdefault("cwd", backend)
            return original_run(command, **kwargs)
        def scan(expected, rejected=False):
            calls.clear(); output = io.StringIO(); failed = False
            with contextlib.redirect_stdout(output), contextlib.redirect_stderr(output):
                try:
                    gate.scan(scan_mode, prefix)
                except gate.GateError as exc:
                    failed = True
                    # GateError messages are fixed safe categories; upstream stderr stays private.
                    print(str(exc))
            text = output.getvalue()
            require(failed == (expected == "FAIL"), "unexpected effective gate outcome")
            require("SYNTHETIC" not in text and "Fingerprint" not in text and "Secret" not in text
                    and "Match" not in text, "unsafe output boundary")
            if rejected:
                require(not calls and "scans not executed" in text and "PASS" not in text,
                        "native rejection did not precede scanner execution")
            return {"effective": "REJECTED_CONFIGURATION" if rejected else expected,
                    "scanner_scans": sum("--report-format" in c for c in calls),
                    "payload_output": False}
        with patch.object(gate, "ROOT", root), patch.object(gate, "BACKEND", backend), patch.object(gate, "HISTORY", accepted), patch.object(gate, "run", run):
            before = gate.run(command)
            findings = gate.report(before.stdout)
            require(before.returncode == 10 and len(findings) == (5 if mode == "history" else 1),
                    "synthetic finding absent before ignore")
            extra = [f for f in findings if (f.get("Commit"), f.get("File"), f.get("StartLine")) not in accepted]
            require(len(extra) == 1, "additional unapproved fixture finding ambiguous")
            no_ignore = scan("FAIL")
            locations = {}
            for label, directory in (("root", root), ("backend", backend)):
                entry = directory / ".gitleaksignore"
                entry.write_text(extra[0]["Fingerprint"] + "\n")
                direct = gate.run(command)
                after = gate.report(direct.stdout)
                gate.raw_status(direct.returncode, len(after), findings_exit=10)
                # Actual omitted-tool oracle on each native-discovery target.
                if (mode == "history" and label == "root") or (mode in ("staged", "unstaged") and label == "backend"):
                    require(len(after) == (4 if mode == "history" else 0), "native discovery oracle changed")
                result = scan("FAIL", rejected=True)
                result.update({"direct_raw_exit": direct.returncode, "direct_findings": len(after),
                               "artifact_preserved": entry.is_file()})
                require(entry.is_file(), "native artifact removed")
                locations[label] = result
                entry.unlink()  # Only this test-owned fixture entry.
            restored = scan("FAIL")
            if mode != "history":
                if mode == "untracked":
                    (backend / "new.txt").unlink()
                else:
                    source.write_text("clean fixture\n")
                    if mode == "staged":
                        git("add", "backend/tracked.txt")
                clean = scan("PASS")
            else:
                clean = accepted_pass
        return {"mode": mode, "before_ignore": {"raw_exit": before.returncode, "findings": len(findings)},
                "no_ignore": no_ignore, "native_entries": locations, "after_removal": restored,
                "normal_allowed": clean, "cleanup": "TemporaryDirectory"}


def main():
    try:
        for mode in ("history", "unstaged", "staged", "untracked"):
            print(json.dumps(regression(mode), sort_keys=True))
        print(json.dumps({"result": "PASS", "tool": "Gitleaks v8.30.1", "image": IMAGE,
                          "network": "none", "mount": "read-only disposable fixture", "synthetic_only": True}))
        return 0
    except gate.GateError as exc:
        # GateError values come only from fixed local categories, never tool payloads.
        print(f"native-ignore regression: FAIL ({exc})", file=sys.stderr)
        return 2
    except (OSError, ValueError, TypeError, KeyError, subprocess.SubprocessError):
        print("native-ignore regression: FAIL (safe fixture/tool/assertion category)", file=sys.stderr)
        return 2


if __name__ == "__main__":
    sys.exit(main())
