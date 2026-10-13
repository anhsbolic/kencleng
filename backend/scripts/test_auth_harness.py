"""Mocked command-boundary tests for ambiguous disposable fixture cleanup."""
import json
import subprocess
import unittest
from unittest.mock import patch

import auth_harness


RUN = "RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-CLEANUP-PATCH-001"
NAME = "c1-auth-harness-a1b2c3d4e5f6"
SHORT_ID = "abcdef012345"
FULL_ID = SHORT_ID + "0" * 52


def result(args, output="", returncode=0):
    return subprocess.CompletedProcess(args, returncode, output, "")


def inspected(*, run=RUN, owner=NAME, name=NAME, container_id=FULL_ID):
    return json.dumps([{
        "Id": container_id,
        "Name": "/" + name,
        "Config": {"Labels": {"c1.auth.run": run, "c1.auth.owner": owner}},
        "Mounts": [],
    }])


class CleanupBoundaryTests(unittest.TestCase):
    def test_timeout_after_container_creation_cleans_verified_full_id(self):
        calls = []

        def podman(args, **kwargs):
            calls.append(args)
            if args[1] == "run":
                raise subprocess.TimeoutExpired(args, 30)
            if args[1] == "ps":
                return result(args, SHORT_ID + "\n" if len(calls) == 2 else "")
            if args[1] == "inspect":
                return result(args, inspected())
            if args[1] == "rm":
                return result(args)
            self.fail("unexpected command: " + " ".join(args))

        with patch.object(auth_harness.subprocess, "run", side_effect=podman):
            with self.assertRaises(subprocess.TimeoutExpired):
                try:
                    auth_harness.command(["podman", "run", "-d", "--name", NAME])
                finally:
                    auth_harness.cleanup_fixture(NAME, RUN)

        self.assertIn(["podman", "rm", "-f", "-v", "--time", "5", FULL_ID], calls)

    def test_nonzero_create_after_container_creation_cleans_verified_full_id(self):
        calls = []

        def podman(args, **kwargs):
            calls.append(args)
            if args[1] == "run":
                return result(args, returncode=125)
            if args[1] == "ps":
                return result(args, SHORT_ID + "\n" if len(calls) == 2 else "")
            if args[1] == "inspect":
                return result(args, inspected())
            if args[1] == "rm":
                return result(args)
            self.fail("unexpected command: " + " ".join(args))

        with patch.object(auth_harness.subprocess, "run", side_effect=podman):
            with self.assertRaisesRegex(RuntimeError, "exit 125"):
                try:
                    auth_harness.command(["podman", "run", "-d", "--name", NAME])
                finally:
                    auth_harness.cleanup_fixture(NAME, RUN)

        self.assertIn(["podman", "rm", "-f", "-v", "--time", "5", FULL_ID], calls)

    def test_absent_resource_after_ambiguous_failure_is_safe(self):
        with patch.object(auth_harness.subprocess, "run", return_value=result([], "")) as run:
            self.assertFalse(auth_harness.cleanup_fixture(NAME, RUN))
        run.assert_called_once()
        self.assertEqual(run.call_args.args[0][1], "ps")

    def test_mismatched_identity_is_never_deleted(self):
        calls = []

        def podman(args, **kwargs):
            calls.append(args)
            if args[1] == "ps":
                return result(args, SHORT_ID + "\n")
            if args[1] == "inspect":
                return result(args, inspected(owner="another-fixture"))
            self.fail("unverified resource must not be acted on: " + " ".join(args))

        with patch.object(auth_harness.subprocess, "run", side_effect=podman):
            with self.assertRaisesRegex(RuntimeError, "unresolved fixture resource"):
                auth_harness.cleanup_fixture(NAME, RUN)
        self.assertFalse(any(args[1] == "rm" for args in calls))

    def test_short_or_malformed_inspect_id_is_never_deleted(self):
        for invalid_id in (SHORT_ID, "g" * 64):
            with self.subTest(container_id=invalid_id):
                calls = []

                def podman(args, **kwargs):
                    calls.append(args)
                    if args[1] == "ps":
                        return result(args, SHORT_ID + "\n")
                    if args[1] == "inspect":
                        return result(args, inspected(container_id=invalid_id))
                    self.fail("unverified resource must not be acted on: " + " ".join(args))

                with patch.object(auth_harness.subprocess, "run", side_effect=podman):
                    with self.assertRaisesRegex(RuntimeError, "unresolved fixture resource"):
                        auth_harness.cleanup_fixture(NAME, RUN)
                self.assertFalse(any(args[1] == "rm" for args in calls))


class MetadataEnvironmentTests(unittest.TestCase):
    def test_fixture_credentials_are_removed_without_mutating_test_environment(self):
        test_env = {
            "DATABASE_URL": "postgres://runtime:secret@127.0.0.1/db",
            "C1_HARNESS_PROVISIONER_DSN": "postgres://admin:secret@127.0.0.1/db",
            "C1_HARNESS_OWNED_RUN": RUN,
            "GOTOOLCHAIN": "go1.26.9",
            "PATH": "/usr/bin",
        }

        metadata_env = auth_harness.metadata_environment(test_env)

        self.assertEqual(metadata_env, {"GOTOOLCHAIN": "go1.26.9", "PATH": "/usr/bin"})
        self.assertEqual(test_env["DATABASE_URL"], "postgres://runtime:secret@127.0.0.1/db")
        self.assertEqual(test_env["C1_HARNESS_PROVISIONER_DSN"], "postgres://admin:secret@127.0.0.1/db")
        self.assertEqual(test_env["C1_HARNESS_OWNED_RUN"], RUN)
        self.assertEqual(test_env["GOTOOLCHAIN"], "go1.26.9")


if __name__ == "__main__":
    unittest.main()
