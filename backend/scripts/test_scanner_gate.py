"""Behavioral oracles for scoped dispositions and fail-closed execution."""
import contextlib
import copy
import io
import json
from pathlib import Path
import subprocess
import tempfile
import unittest
from unittest.mock import patch

import scanner_gate as gate


def history_report():
    return [{"Commit": commit, "File": path, "RuleID": "generic-api-key",
             "StartLine": line, "EndLine": line,
             "Fingerprint": f"{commit}:{path}:generic-api-key:{line}"}
            for commit, path, line in gate.HISTORY]


def gosec_report():
    return {"Golang errors": {}, "GosecVersion": "dev",
            "Stats": {"files": 6, "lines": 750, "nosec": 0, "found": 3},
            "Issues": [{"rule_id": "G124", "file": str(gate.BACKEND / "internal/platform/auth/http.go"),
                        "line": line, "column": column, "nosec": False}
                       for line, column in (("49", "21"), ("158", "23"), ("277", "23"))]}


class DispositionTests(unittest.TestCase):
    def test_native_entries_rejected_before_any_scan(self):
        for location in ("root", "backend"):
            for name in (".gitleaksignore", ".gitleaks.toml"):
                for state in ("file", "directory", "symlink", "dangling", "loop"):
                    with self.subTest(location=location, name=name, state=state), tempfile.TemporaryDirectory() as folder:
                        root = Path(folder); backend = root / "backend"; backend.mkdir()
                        entry = (root if location == "root" else backend) / name
                        target = root / "fixture-policy"; target.write_text("PRIVATE_POLICY")
                        if state == "file":
                            entry.write_text("PRIVATE_POLICY")
                        elif state == "directory":
                            entry.mkdir()
                        else:
                            entry.symlink_to(target if state == "symlink" else
                                             root / "missing" if state == "dangling" else entry)
                        for mode in ("gitleaks-history", "gitleaks-worktree"):
                            with patch.object(gate, "ROOT", root), patch.object(gate, "BACKEND", backend), patch.object(gate, "run") as run:
                                with self.assertRaisesRegex(gate.GateError, "policy entry; scans not executed"):
                                    gate.scan(mode, ["scanner"])
                                run.assert_not_called()
                        self.assertTrue(entry.is_symlink() or entry.exists())

    def test_native_file_state_errors_fail_closed_and_output_is_safe(self):
        for location in ("root", "backend"):
            for name in (".gitleaksignore", ".gitleaks.toml"):
                for error in (PermissionError("PRIVATE_STATE"), NotADirectoryError("PRIVATE_STATE"), OSError("PRIVATE_STATE")):
                    with self.subTest(location=location, name=name, error=type(error).__name__), tempfile.TemporaryDirectory() as folder:
                        root = Path(folder); backend = root / "backend"; backend.mkdir()
                        selected = (root if location == "root" else backend) / name
                        original = Path.lstat
                        def state(path):
                            if path == selected:
                                raise error
                            return original(path)
                        output = io.StringIO()
                        for mode in ("gitleaks-history", "gitleaks-worktree"):
                            with patch.object(gate, "ROOT", root), patch.object(gate, "BACKEND", backend), patch.object(Path, "lstat", state), patch.object(gate, "run") as run, patch.object(gate.sys, "argv", ["gate", mode, "--", "scanner"]), contextlib.redirect_stderr(output):
                                self.assertEqual(gate.main(), 1)
                                run.assert_not_called()
                        self.assertIn("state unavailable; scans not executed", output.getvalue())
                        self.assertNotIn("PRIVATE_STATE", output.getvalue())
                        self.assertNotIn("PASS", output.getvalue())

    def test_absent_native_entries_allow_normal_history_and_worktree_routes(self):
        with tempfile.TemporaryDirectory() as folder:
            root = Path(folder); backend = root / "backend"; backend.mkdir()
            for mode in ("gitleaks-history", "gitleaks-worktree"):
                calls = []
                def run(command, **kwargs):
                    calls.append(command)
                    if command[-1] == "version":
                        return subprocess.CompletedProcess(command, 0, b"v8.30.1", b"")
                    if command[0] == "git":
                        return subprocess.CompletedProcess(command, 0, b"backend/new.txt\0", b"")
                    return subprocess.CompletedProcess(command, 10 if mode == "gitleaks-history" else 0,
                                                       json.dumps(history_report()).encode() if mode == "gitleaks-history" else b"[]", b"")
                output = io.StringIO()
                with patch.object(gate, "ROOT", root), patch.object(gate, "BACKEND", backend), patch.object(gate, "run", run), patch.object(gate, "guard_history"), contextlib.redirect_stdout(output):
                    gate.scan(mode, ["scanner"])
                scans = [c for c in calls if "--report-format" in c]
                self.assertEqual(len(scans), 1 if mode == "gitleaks-history" else 3)
                if mode == "gitleaks-history":
                    self.assertIn("--log-opts=--all --full-history -- backend/", scans[0])
                self.assertNotIn("FAIL", output.getvalue())

    def test_exact_candidates_and_real_source_guards(self):
        gate.guard_source()
        gate.guard_history()
        self.assertEqual(gate.gosec_result(json.dumps(gosec_report()), 1), 3)
        self.assertEqual(gate.gitleaks_result(json.dumps(history_report()), 10, True), 4)
        self.assertEqual(gate.gitleaks_result("[]", 0), 0)

    def test_history_drift_unrelated_new_and_duplicates_fail(self):
        for field, value in (("Commit", "0" * 40), ("File", "backend/new.go"),
                             ("RuleID", "other"), ("StartLine", 170),
                             ("EndLine", 170), ("Fingerprint", "different")):
            d = history_report()
            d[0][field] = value
            with self.subTest(field=field), self.assertRaises(gate.GateError):
                gate.gitleaks_result(json.dumps(d), 10, True)
        for d in (history_report()[:-1], history_report() + [history_report()[0]],
                  history_report() + [{"RuleID": "new-secret"}]):
            with self.assertRaises(gate.GateError):
                gate.gitleaks_result(json.dumps(d), 10, True)
        with self.assertRaises(gate.GateError):
            gate.gitleaks_result(json.dumps(history_report()), 10)  # no worktree allowances

    def test_gosec_rule_location_path_and_analysis_drift_fail(self):
        for field, value in (("rule_id", "G101"), ("line", "50"), ("column", "22"),
                             ("file", "/different/http.go"), ("nosec", True)):
            d = gosec_report()
            d["Issues"][0][field] = value
            with self.subTest(field=field), self.assertRaises(gate.GateError):
                gate.gosec_result(json.dumps(d), 1)
        for change in ("missing", "duplicate", "loading", "coverage", "version"):
            d = gosec_report()
            if change == "missing":
                d["Issues"].pop(); d["Stats"]["found"] -= 1
            elif change == "duplicate":
                d["Issues"][1] = copy.deepcopy(d["Issues"][0])
            elif change == "loading":
                d["Golang errors"] = {"package": [{"error": "synthetic failure"}]}
            elif change == "coverage":
                d["Stats"]["files"] = 5
            else:
                d["GosecVersion"] = "new"
            with self.subTest(change=change), self.assertRaises(gate.GateError):
                gate.gosec_result(json.dumps(d), 1)

    def test_source_mutation_new_caller_and_history_source_error_fail(self):
        with tempfile.TemporaryDirectory() as folder:
            root = Path(folder)
            for path in gate.SOURCE:
                p = root / path; p.parent.mkdir(parents=True, exist_ok=True)
                p.write_bytes((gate.BACKEND / path).read_bytes())
            with patch.object(gate, "BACKEND", root):
                gate.guard_source()
                (root / "new_caller.go").write_text("package backend\n")
                with self.assertRaises(gate.GateError):
                    gate.guard_source()
                (root / "new_caller.go").unlink()
                with (root / "cmd/server/main.go").open("a") as f:
                    f.write("// changed composition\n")
                with self.assertRaises(gate.GateError):
                    gate.guard_source()
        for rc, blob in ((0, b"changed blob"), (128, b"")):
            with patch.object(gate, "run", return_value=subprocess.CompletedProcess([], rc, blob, b"")):
                with self.assertRaises(gate.GateError):
                    gate.guard_history()

    def test_report_and_execution_failures_cannot_be_filtered(self):
        for data in (b"", b"[", b"null", b"{}", b'{"Issues":[],"Issues":[]}'):
            with self.subTest(data=data), self.assertRaises(gate.GateError):
                gate.gosec_result(data, 0)
            with self.assertRaises(gate.GateError):
                gate.gitleaks_result(data, 0)
        for rc in (0, 2, -9):
            with self.assertRaises(gate.GateError):
                gate.gosec_result(json.dumps(gosec_report()), rc)
            with self.assertRaises(gate.GateError):
                gate.gitleaks_result(json.dumps(history_report()), rc, True)
        with self.assertRaises(gate.GateError):
            gate.gitleaks_result("[]", 1)
        # Partial scans use exit 1 even when all approved findings are present.
        with self.assertRaises(gate.GateError):
            gate.gitleaks_result(json.dumps(history_report()), 1, True)
        for error in (FileNotFoundError(), subprocess.TimeoutExpired("scanner", 240)):
            with patch.object(gate.subprocess, "run", side_effect=error):
                with self.assertRaises(gate.GateError):
                    gate.run(["scanner"])

    def test_worktree_all_scopes_attempted_after_failure_and_payload_is_private(self):
        calls = []
        def fake_run(command, **kwargs):
            calls.append(command)
            if command[-1] == "version":
                return subprocess.CompletedProcess(command, 0, b"v8.30.1\n", b"")
            if command[0] == "git":
                return subprocess.CompletedProcess(command, 0, b"backend/new file.txt\0", b"")
            return subprocess.CompletedProcess(command, 10, b'[{"Secret":"SYNTHETIC_PRIVATE_PAYLOAD"}]',
                                               b"SYNTHETIC_PRIVATE_PAYLOAD")
        output = io.StringIO()
        with patch.object(gate, "run", side_effect=fake_run), contextlib.redirect_stderr(output):
            with self.assertRaises(gate.GateError):
                gate.scan("gitleaks-worktree", ["scanner"])
        scans = [c for c in calls if "--report-format" in c]
        self.assertEqual(len(scans), 3)
        self.assertIn("--pre-commit", scans[0]); self.assertIn("--staged", scans[1])
        self.assertEqual(scans[2][1:3], ["dir", "/repo/backend/new file.txt"])
        self.assertNotIn("SYNTHETIC_PRIVATE_PAYLOAD", output.getvalue())
        self.assertEqual(output.getvalue().count("gitleaks-worktree["), 3)

    def test_gosec_executed_integration_checks_metadata_and_propagates_failure(self):
        metadata = {"Main": {"Path": "github.com/securego/gosec/v2", "Version": "v2.28.0",
                             "Sum": "h1:ZsSdiDb0AtTpLFVol5z91gbMei9ZiLEPG/pZjZujp7c="}}
        for rc, data, version in ((1, json.dumps(gosec_report()), "v2.28.0"),
                                  (2, json.dumps(gosec_report()), "v2.28.0"),
                                  (0, "{}", "v2.28.0"),
                                  (1, json.dumps(gosec_report()), "v9.0.0")):
            meta = copy.deepcopy(metadata); meta["Main"]["Version"] = version
            calls = []
            def fake_run(command, **kwargs):
                calls.append((command, kwargs))
                if command[0] == "go":
                    return subprocess.CompletedProcess(command, 0, json.dumps(meta).encode(), b"")
                return subprocess.CompletedProcess(command, rc, data.encode(), b"private scanner stderr")
            output = io.StringIO()
            with patch.object(gate, "run", side_effect=fake_run), contextlib.redirect_stdout(output), contextlib.redirect_stderr(output):
                if rc == 1 and version == "v2.28.0":
                    gate.scan("gosec", ["scanner"])
                    self.assertIn("raw_exit=1 raw_findings=3 accepted=3 effective=PASS", output.getvalue())
                else:
                    with self.assertRaises(gate.GateError):
                        gate.scan("gosec", ["scanner"])
            if version == "v2.28.0":
                self.assertEqual(calls[1][0], ["scanner", "-fmt=json", "-nosec", "./..."])
                self.assertEqual(calls[1][1]["env"]["GOFLAGS"], "-mod=readonly")
            else:
                self.assertEqual(len(calls), 1)
            self.assertNotIn("private scanner stderr", output.getvalue())


if __name__ == "__main__":
    unittest.main()
