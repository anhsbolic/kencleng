#!/usr/bin/env python3
"""Initialize/check the explicitly selected local Kencleng S3 Compose service.

Uses the pinned AWS CLI container for signing; no SDK or signing code is added to
the application. Checks create UUID-scoped objects and remove only those objects.
"""

import argparse
import json
from pathlib import Path
import shlex
import subprocess
import sys
import time
import urllib.error
import urllib.request
import uuid

ROOT = Path(__file__).resolve().parents[1]
INTERNAL = "http://seaweedfs:8333"
PUBLIC = "http://127.0.0.1:9087"
BUCKETS = ("kencleng-public", "kencleng-private")
PAYLOAD = (ROOT / "infra/storage/probe.txt").read_bytes()


class Storage:
    def __init__(self, compose):
        self.compose = [*shlex.split(compose), "-p", "kencleng-storage", "-f", "docker-compose.storage.yml"]
        if not shlex.split(compose):
            raise ValueError("Compose provider must be explicit and nonempty")
        # A local probe must bypass user-configured HTTP proxies.
        self.http = urllib.request.build_opener(urllib.request.ProxyHandler({}))

    def cli(self, *args, endpoint=INTERNAL):
        result = subprocess.run(
            [*self.compose, "run", "--rm", "-T", "storage-cli", "--endpoint-url", endpoint,
             "--region", "us-east-1", "--no-cli-pager", "--cli-connect-timeout", "3",
             "--cli-read-timeout", "5", *args],
            cwd=ROOT, text=True, capture_output=True, timeout=90, check=False,
        )
        if result.returncode:
            # Do not print captured stdout: presign returns bearer URLs there.
            raise RuntimeError(f"Storage CLI failed ({result.returncode}): {result.stderr.strip()}")
        return result.stdout.strip()

    def ready(self):
        deadline = time.monotonic() + 120
        while True:
            try:
                self.cli("s3api", "list-buckets", "--output", "json")
                return
            except RuntimeError:
                if time.monotonic() >= deadline:
                    raise
                time.sleep(2)

    def init(self):
        self.ready()
        for bucket in BUCKETS:
            try:
                self.cli("s3api", "head-bucket", "--bucket", bucket)
            except RuntimeError as error:
                if "404" not in str(error) and "NotFound" not in str(error):
                    raise
                self.cli("s3api", "create-bucket", "--bucket", bucket)
        policy = {"Version": "2012-10-17", "Statement": [{
            "Effect": "Allow", "Principal": "*", "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::kencleng-public/*",
        }]}
        self.cli("s3api", "put-bucket-policy", "--bucket", BUCKETS[0], "--policy", json.dumps(policy))
        cors = {"CORSRules": [{
            "AllowedOrigins": ["http://localhost:3000", "http://127.0.0.1:3000"],
            "AllowedMethods": ["GET", "PUT", "HEAD"], "AllowedHeaders": ["*"],
            "ExposeHeaders": ["ETag"], "MaxAgeSeconds": 300,
        }]}
        for bucket in BUCKETS:
            self.cli("s3api", "put-bucket-cors", "--bucket", bucket, "--cors-configuration", json.dumps(cors))
        print("Storage initialized: public object reads, private bucket, localhost CORS.")

    def request(self, url, expected, data=None, method="GET", headers=None):
        request = urllib.request.Request(url, data=data, method=method, headers=headers or {})
        try:
            response = self.http.open(request, timeout=10)
        except urllib.error.HTTPError as error:
            response = error
        with response:
            body = response.read()
            status = response.status
            response_headers = response.headers
        if status != expected:
            # URL may contain a presigned bearer token; never include it in errors.
            raise RuntimeError(f"Storage HTTP {method}: expected {expected}, received {status}")
        return body, response_headers

    def presign(self, bucket, key, expires=60):
        url = self.cli("s3", "presign", f"s3://{bucket}/{key}", "--expires-in", str(expires), endpoint=PUBLIC)
        if not url.startswith(f"{PUBLIC}/{bucket}/") or "X-Amz-Signature=" not in url:
            raise RuntimeError("Expected path-style SigV4 URL using the browser-accessible endpoint")
        return url

    def check(self, restart=False):
        self.ready()
        key = f"kencleng-setup-probe/{uuid.uuid4()}.txt"
        created = []
        try:
            for bucket in BUCKETS:
                # UUID keys cannot predate this run; include ambiguous upload failures in cleanup.
                created.append(bucket)
                self.cli("s3api", "put-object", "--bucket", bucket, "--key", key, "--body", "/config/probe.txt")
            public_url = f"{PUBLIC}/{BUCKETS[0]}/{key}"
            private_url = f"{PUBLIC}/{BUCKETS[1]}/{key}"
            body, _ = self.request(public_url, 200)
            if body != PAYLOAD:
                raise RuntimeError("Public object content did not match the uploaded probe")
            self.request(private_url, 403)
            for url in (public_url, private_url):
                self.request(url, 403, method="PUT", data=PAYLOAD)
            for bucket in BUCKETS:
                body, _ = self.request(self.presign(bucket, key), 200)
                if body != PAYLOAD:
                    raise RuntimeError("Presigned object content did not match the uploaded probe")
            signed = self.presign(BUCKETS[1], key, expires=2)
            time.sleep(4)
            self.request(signed, 403)
            for url in (public_url, private_url):
                _, headers = self.request(url, 200, method="OPTIONS", headers={
                    "Origin": "http://localhost:3000", "Access-Control-Request-Method": "PUT",
                    "Access-Control-Request-Headers": "content-type",
                })
                if headers.get("Access-Control-Allow-Origin") != "http://localhost:3000":
                    raise RuntimeError("Local browser CORS origin was not allowed")
                if "PUT" not in headers.get("Access-Control-Allow-Methods", ""):
                    raise RuntimeError("Local browser CORS did not allow PUT")
            if restart:
                subprocess.run([*self.compose, "restart", "seaweedfs"], cwd=ROOT, check=True, timeout=90)
                self.ready()
                body, _ = self.request(self.presign(BUCKETS[1], key), 200)
                if body != PAYLOAD:
                    raise RuntimeError("Probe content did not survive the storage restart")
        finally:
            failures = []
            for bucket in created:
                try:
                    self.cli("s3api", "delete-object", "--bucket", bucket, "--key", key)
                except Exception as error:
                    failures.append(f"{bucket}/{key}: {error}")
            if failures:
                raise RuntimeError("Probe cleanup failed; remove only these keys:\n" + "\n".join(failures))
        print("PASS: authenticated upload, public/private reads, anonymous write denial, presigned GET/expiry, CORS" +
              (", restart persistence." if restart else "."))


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--compose", required=True, help="docker compose, podman-compose, or podman compose")
    parser.add_argument("action", choices=("init", "check", "persistence-check"))
    args = parser.parse_args()
    storage = Storage(args.compose)
    if args.action == "init":
        storage.init()
    else:
        storage.check(restart=args.action == "persistence-check")


if __name__ == "__main__":
    try:
        main()
    except (RuntimeError, OSError, ValueError, subprocess.SubprocessError, KeyboardInterrupt) as error:
        print(f"Storage check stopped: {error}", file=sys.stderr)
        sys.exit(1)
