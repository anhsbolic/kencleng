"""Bounded C1 scanner dispositions; never emit scanner payloads or stderr."""
import hashlib
import json
import os
from pathlib import Path
import subprocess
import sys


ROOT = Path(__file__).resolve().parents[2]
BACKEND = ROOT / "backend"
# Human option 1, 2026-10-12. Blob hashes are independently reviewed source
# identities, not secret values. No path/rule/commit-wide exception is applied.
HISTORY = {
    ("6f036c6384868644989134386fc378ed1c19efeb", "backend/internal/transport/http/account_security_test.go", 169): "b0c6ec64d5122fb1b845e7e4ac266a762e7c8d727bc4000bc1e59b6ed68061cf",
    ("6f036c6384868644989134386fc378ed1c19efeb", "backend/internal/transport/http/account_security_test.go", 245): "b0c6ec64d5122fb1b845e7e4ac266a762e7c8d727bc4000bc1e59b6ed68061cf",
    ("ef428be3742eddfc59b86fc205bb03624e669e82", "backend/internal/domain/account/google_oauth_test.go", 487): "d5f11425602daa52b90bb831eba581d5be8aef382224507274b3c727715d5e3e",
    ("d68b40ce7d8d7f3b7334fdf45d5d4a1899e61c0f", "backend/internal/domain/account/service_test.go", 471): "0ad4135b8fa4db140cc8086e8014614c1b54d5e3bf29b591b3ca54906dbe7a69",
}
# Guard the reviewed constructor/composition and all current production callers.
# New/changed production Go source requires re-triage before G124 disposition.
SOURCE = {
    "cmd/server/main.go": "d217dfd21898ce89d3eee96683b7512fd7010ec0e55da04c6d211bb5686e2246",
    "internal/platform/auth/http.go": "83b575b0e43d17d85303ae7dfeb0273fa2516bb18952a9660da7b80a2d039d63",
    "internal/platform/auth/oidc.go": "78424fa09c46554101d9e8a1c8d177bdcd1421d7415262aa4fc704b8357a0ac2",
    "internal/platform/auth/store.go": "eb70175ebe756698a3a83b627e0fd6606928d8ac4c355f9dd8a05dab705552b3",
    "internal/platform/db/db.go": "56492320609c7bc0acbbf0761bd6669a9f79fa31986e92a2ea88138798a63660",
    "internal/platform/db/doc.go": "8dbdc9d74ef1d1f42f04ea216a57448b1ca1ca5d7a2aa15f5ac79a9eec619180",
}


class GateError(Exception):
    """Only fixed safe messages may cross the gate's output boundary."""


def run(command, cwd=BACKEND, env=None):
    try:
        return subprocess.run(command, cwd=cwd, env=env, capture_output=True,
                              timeout=240, check=False)
    except (OSError, subprocess.TimeoutExpired) as exc:
        raise GateError("scanner/process execution failed or timed out") from exc


def report(data):
    def unique(pairs):
        result = {}
        for key, value in pairs:
            if key in result:
                raise GateError("duplicate report field")
            result[key] = value
        return result
    try:
        return json.loads(data, object_pairs_hook=unique)
    except (ValueError, UnicodeError) as exc:
        raise GateError("missing/invalid scanner JSON report") from exc


def raw_status(exit_code, count, findings_exit=1):
    if exit_code != (findings_exit if count else 0):
        raise GateError("scanner execution/report status inconsistent")


def evaluate(result, validator, **kwargs):
    """Keep raw and effective outcomes separate, including rejected reports."""
    try:
        return validator(result.stdout, result.returncode, **kwargs)
    except (GateError, TypeError, KeyError, AttributeError):
        print(f"scanner: raw_exit={result.returncode} report_bytes={len(result.stdout)} effective=FAIL",
              file=sys.stderr)
        raise


def guard_source():
    actual = {str(p.relative_to(BACKEND)): hashlib.sha256(p.read_bytes()).hexdigest()
              for p in BACKEND.rglob("*.go") if not p.name.endswith("_test.go")}
    if actual != SOURCE:
        raise GateError("reviewed production source drift; re-triage required")


def guard_history():
    for (commit, path, _), expected in HISTORY.items():
        result = run(["git", "show", f"{commit}:{path}"], cwd=ROOT)
        if result.returncode or hashlib.sha256(result.stdout).hexdigest() != expected:
            raise GateError("reviewed history source unavailable/drifted")


def gosec_result(data, exit_code):
    d = report(data)
    if not isinstance(d, dict) or set(d) != {"Golang errors", "Issues", "Stats", "GosecVersion"}:
        raise GateError("incomplete/changed Gosec report schema")
    issues, stats = d["Issues"], d["Stats"]
    if d["Golang errors"] != {} or not isinstance(issues, list) or not isinstance(stats, dict):
        raise GateError("Gosec loading error or incomplete report")
    if (stats.get("files") != len(SOURCE) or stats.get("found") != len(issues)
            or stats.get("nosec") != 0 or type(stats.get("lines")) is not int
            or stats["lines"] <= 0 or d["GosecVersion"] != "dev"):
        raise GateError("Gosec coverage/version/statistics changed")
    raw_status(exit_code, len(issues))
    seen = set()
    for issue in issues:
        if not isinstance(issue, dict):
            raise GateError("invalid Gosec finding")
        candidate = (issue.get("rule_id"), issue.get("line"), issue.get("column"))
        if (candidate not in {("G124", "49", "21"), ("G124", "158", "23"), ("G124", "277", "23")}
                or issue.get("file") != str(BACKEND / "internal/platform/auth/http.go")
                or issue.get("nosec") is not False or issue.get("suppressions")
                or candidate in seen):
            raise GateError("unexpected/drifted Gosec finding")
        seen.add(candidate)
    # Disappearing candidates indicate changed analysis/applicability too.
    if len(seen) != 3:
        raise GateError("G124 analysis/applicability changed; re-triage required")
    return len(issues)


def gitleaks_result(data, exit_code, history=False):
    findings = report(data)
    if not isinstance(findings, list):
        raise GateError("incomplete Gitleaks report")
    # Gitleaks defaults both findings AND partial-scan errors to exit 1.
    # Reserve 10 for findings so a valid partial JSON report cannot pass.
    raw_status(exit_code, len(findings), findings_exit=10)
    seen = set()
    for finding in findings:
        if not isinstance(finding, dict):
            raise GateError("invalid Gitleaks finding")
        commit, path, line = (finding.get(k) for k in ("Commit", "File", "StartLine"))
        if not history:
            raise GateError("unexpected worktree finding")
        key = (commit, path, line)
        if (type(line) is not int or key not in HISTORY or key in seen
                or finding.get("EndLine") != line or finding.get("RuleID") != "generic-api-key"
                or finding.get("Fingerprint") != f"{commit}:{path}:generic-api-key:{line}"):
            raise GateError("unexpected/drifted history finding")
        seen.add(key)
    if history and seen != set(HISTORY):
        raise GateError("history candidate coverage/applicability changed")
    return len(findings)


def leaks_command(prefix, mode, target, *extra):
    return prefix + [mode, target, *extra, "--redact", "--no-banner", "--no-color",
                     "--timeout", "180", "--exit-code", "10", "--ignore-gitleaks-allow",
                     "--gitleaks-ignore-path", "/dev/null", "--report-format", "json",
                     "--report-path", "-"]


def guard_native_policy():
    # v8.30.1 auto-loads target .gitleaksignore even with an explicit /dev/null
    # ignore path. lstat rejects every entry, including dangling symlinks,
    # without reading user content. Only confirmed absence permits scanning.
    for directory in (ROOT, BACKEND):
        for name in (".gitleaks.toml", ".gitleaksignore"):
            try:
                (directory / name).lstat()
            except FileNotFoundError:
                continue
            except OSError as exc:
                raise GateError("native Gitleaks policy state unavailable; scans not executed") from exc
            raise GateError("unreviewed native Gitleaks policy entry; scans not executed")


def scan(mode, prefix):
    if mode == "gosec":
        guard_source()
        metadata = run(["go", "version", "-m", "-json", prefix[0]])
        info = report(metadata.stdout)
        if (metadata.returncode or info.get("Main", {}).get("Path") != "github.com/securego/gosec/v2"
                or info.get("Main", {}).get("Version") != "v2.28.0"
                or info.get("Main", {}).get("Sum") != "h1:ZsSdiDb0AtTpLFVol5z91gbMei9ZiLEPG/pZjZujp7c="
                or info.get("Main", {}).get("Replace")):
            raise GateError("Gosec build identity unsupported")
        env = os.environ.copy()
        env["GOFLAGS"] = "-mod=readonly"
        result = run(prefix + ["-fmt=json", "-nosec", "./..."], env=env)
        count = evaluate(result, gosec_result)
        guard_source()
        print(f"gosec: raw_exit={result.returncode} raw_findings={count} accepted={count} effective=PASS")
        return
    guard_native_policy()
    version = run(prefix + ["version"])
    if version.returncode or version.stdout.strip() != b"v8.30.1":
        raise GateError("Gitleaks version unsupported")
    if mode == "gitleaks-history":
        guard_history()
        result = run(leaks_command(prefix, "git", "/repo", "--log-opts=--all --full-history -- backend/"))
        count = evaluate(result, gitleaks_result, history=True)
        guard_history()
        print(f"gitleaks-history: raw_exit={result.returncode} raw_findings={count} accepted={count} effective=PASS")
        return
    files = run(["git", "ls-files", "-z", "--others", "--exclude-standard", "--", "backend/"], cwd=ROOT)
    if files.returncode:
        raise GateError("Git-visible file inventory failed")
    commands = [leaks_command(prefix, "git", "/repo/backend", flag)
                for flag in ("--pre-commit", "--staged")]
    commands += [leaks_command(prefix, "dir", "/repo/" + os.fsdecode(path))
                 for path in files.stdout.split(b"\0") if path]
    failed = False
    for index, command in enumerate(commands):
        try:
            result = run(command)
            count = evaluate(result, gitleaks_result)
            print(f"gitleaks-worktree[{index}]: raw_exit={result.returncode} raw_findings={count} accepted=0 effective=PASS")
        except GateError as exc:
            failed = True
            print(f"gitleaks-worktree[{index}]: effective=FAIL ({exc})", file=sys.stderr)
    if failed:
        raise GateError("Git-visible worktree gate failed; all inventoried scans attempted")


def main():
    try:
        if len(sys.argv) < 4 or sys.argv[1] not in {"gosec", "gitleaks-history", "gitleaks-worktree"} or sys.argv[2] != "--":
            raise GateError("invalid scanner gate invocation")
        scan(sys.argv[1], sys.argv[3:])
    except (GateError, OSError, TypeError, KeyError, AttributeError) as exc:
        # Even malformed report values never reach stdout/stderr verbatim.
        message = str(exc) if isinstance(exc, GateError) else "invalid scanner/report/source state"
        print(f"scanner gate: effective=FAIL ({message})", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    sys.exit(main())
