"""Failure/cleanup regressions with a simulated S3 service; no runtime evidence."""

import contextlib
import importlib.util
import io
from pathlib import Path
import subprocess
import unittest
from unittest.mock import patch

SPEC = importlib.util.spec_from_file_location("local_storage", Path(__file__).with_name("local-storage.py"))
storage = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(storage)


class FakeStorage(storage.Storage):
    def __init__(self, failure=None):
        super().__init__("podman-compose")
        self.failure = failure
        self.objects = {("kencleng-private", "existing-user-object"): b"preserve me"}
        self.deleted = []

    def ready(self):
        pass

    def cli(self, *args, endpoint=storage.INTERNAL):
        operation = args[1]
        bucket = args[args.index("--bucket") + 1]
        key = args[args.index("--key") + 1]
        if operation == "put-object":
            self.objects[bucket, key] = storage.PAYLOAD
            if self.failure == "ambiguous-upload" and bucket == "kencleng-private":
                raise RuntimeError("connection lost after server stored the upload")
        elif operation == "delete-object":
            if self.failure == "cleanup":
                raise RuntimeError("delete unavailable")
            self.deleted.append((bucket, key))
            self.objects.pop((bucket, key), None)
        else:
            raise AssertionError(operation)
        return ""

    def presign(self, bucket, key, expires=60):
        return f"{storage.PUBLIC}/{bucket}/{key}?ttl={expires}"

    def request(self, url, expected, data=None, method="GET", headers=None):
        path = url.removeprefix(storage.PUBLIC + "/").split("?")[0]
        bucket, key = path.split("/", 1)
        response_headers = {}
        status = 200
        if method == "OPTIONS":
            response_headers = {"Access-Control-Allow-Origin": "http://localhost:3000", "Access-Control-Allow-Methods": "GET, PUT, HEAD"}
            if self.failure == "cors":
                response_headers = {}
        elif method == "PUT":
            status = 403
        elif "ttl=2" in url:
            status = 200 if self.failure == "expiry" else 403
        elif "?" not in url and bucket == "kencleng-private":
            status = 200 if self.failure == "private-public" else 403
        if expected != status:
            raise RuntimeError(f"expected {expected}, received {status}")
        return self.objects[bucket, key], response_headers


class StorageChecksTests(unittest.TestCase):
    def check(self, fake, restart=False):
        with patch.object(storage.time, "sleep"), contextlib.redirect_stdout(io.StringIO()) as output:
            fake.check(restart=restart)
        return output.getvalue()

    def assert_preserved(self, fake):
        self.assertEqual(fake.objects, {("kencleng-private", "existing-user-object"): b"preserve me"})
        self.assertTrue(all(key.startswith("kencleng-setup-probe/") for _, key in fake.deleted))

    def test_success_removes_only_this_runs_objects(self):
        fake = FakeStorage()
        self.assertIn("PASS", self.check(fake))
        self.assert_preserved(fake)

    def test_security_failures_are_fatal_and_cleanup_runs(self):
        for failure in ("private-public", "expiry", "cors", "ambiguous-upload"):
            with self.subTest(failure=failure):
                fake = FakeStorage(failure)
                with self.assertRaises(RuntimeError):
                    self.check(fake)
                self.assert_preserved(fake)

    def test_cleanup_failure_does_not_report_pass(self):
        fake = FakeStorage("cleanup")
        with patch.object(storage.time, "sleep"), contextlib.redirect_stdout(io.StringIO()) as output:
            with self.assertRaisesRegex(RuntimeError, "Probe cleanup failed"):
                fake.check()
        self.assertNotIn("PASS", output.getvalue())

    def test_restart_failure_is_fatal_and_cleanup_runs(self):
        fake = FakeStorage()
        with patch.object(storage.subprocess, "run", side_effect=subprocess.CalledProcessError(1, "restart")) as run:
            with self.assertRaises(subprocess.CalledProcessError):
                self.check(fake, restart=True)
        self.assertEqual(run.call_args.args[0][-2:], ["restart", "seaweedfs"])
        self.assert_preserved(fake)

    def test_cli_preserves_provider_arguments_and_does_not_shell_interpolate(self):
        instance = storage.Storage("docker --context local compose")
        result = subprocess.CompletedProcess([], 0, stdout="{}", stderr="")
        with patch.object(storage.subprocess, "run", return_value=result) as run:
            instance.cli("s3api", "list-buckets")
        args = run.call_args.args[0]
        self.assertEqual(args[:4], ["docker", "--context", "local", "compose"])
        self.assertIn("--rm", args)
        self.assertNotIn("shell", run.call_args.kwargs)

    def test_bearer_url_not_printed_when_cli_fails(self):
        instance = storage.Storage("podman-compose")
        result = subprocess.CompletedProcess([], 1, stdout="sensitive-presigned-url", stderr="network unavailable")
        with patch.object(storage.subprocess, "run", return_value=result):
            with self.assertRaises(RuntimeError) as error:
                instance.cli("s3", "presign", "s3://bucket/key")
        self.assertNotIn("sensitive-presigned-url", str(error.exception))


if __name__ == "__main__":
    unittest.main()
