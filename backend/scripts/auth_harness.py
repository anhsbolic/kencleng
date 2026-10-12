#!/usr/bin/env python3
"""Owned disposable T2 fixture. Never reads .env or application DATABASE_URL."""
import json
import os
from pathlib import Path
import secrets
import re
import signal
import subprocess
import sys
import time

RUN = "RUN-C1-ENG-BUILD-T2-LOGOUT-PATCH-001"
ROOT = Path(__file__).resolve().parents[1]


def command(args, *, env=None, data=None, timeout=30):
    result = subprocess.run(args, input=data, env=env, text=True,
                            stdout=subprocess.PIPE, stderr=subprocess.PIPE, timeout=timeout)
    if result.returncode:
        # Tool errors can contain private connection arguments. Do not dump them.
        safe = re.sub(r"[0-9a-f]{64}", "[private value]", result.stderr.strip())
        safe = re.sub(r"postgres(?:ql)?://\S+", "[private DSN]", safe)
        raise RuntimeError(f"{args[0]} {args[1]} operation failed (exit {result.returncode}): {safe}")
    return result.stdout.strip()


def main():
    name = "c1-auth-harness-" + secrets.token_hex(6)
    env = os.environ.copy()
    # Do not propagate application/provider credentials into fixture/test children.
    for key in ("DATABASE_URL", "C1_AUTH_TEST_DATABASE_URL", "GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET",
                "C1_HARNESS_RUNTIME_DSN", "C1_HARNESS_PROVISIONER_DSN"):
        env.pop(key, None)
    env["POSTGRES_PASSWORD"] = secrets.token_hex(32)
    runtime_password = secrets.token_hex(32)
    created = False
    try:
        if command(["podman", "ps", "-a", "--filter", "name=^" + name + "$", "--format", "{{.ID}}"]):
            raise RuntimeError("resource collision")
        command(["podman", "image", "exists", "postgres:16"])
        # No volumes, existing network, host .env, or persistent data mounts.
        command(["podman", "run", "-d", "--name", name, "--label", "c1.auth.run=" + RUN,
                 "--label", "c1.auth.owner=" + name, "--pull=never", "--publish", "127.0.0.1::5432",
                 "--env", "POSTGRES_PASSWORD", "--env", "POSTGRES_HOST_AUTH_METHOD=scram-sha-256",
                 "--env", "POSTGRES_INITDB_ARGS=--auth-host=scram-sha-256",
                 "--env", "POSTGRES_DB=c1_fixture", "postgres:16"], env=env)
        created = True
        metadata = json.loads(command(["podman", "inspect", name]))[0]
        print("fixture", name, "image", metadata["Image"], flush=True)
        if any(m.get("Type") == "bind" for m in metadata.get("Mounts", [])):
            raise RuntimeError("unexpected bind mount")
        port = command(["podman", "port", name, "5432/tcp"])
        if not port.startswith("127.0.0.1:") or "\n" in port:
            raise RuntimeError("fixture port is not loopback-only")
        deadline = time.monotonic() + 30
        while True:
            # Initdb's temporary server accepts Unix sockets before shutdown.
            # TCP readiness selects the final server, avoiding that restart race.
            ready = subprocess.run(["podman", "exec", name, "pg_isready", "-h", "127.0.0.1", "-U", "postgres", "-d", "c1_fixture"],
                                   stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL, timeout=5)
            if ready.returncode == 0:
                break
            if time.monotonic() >= deadline:
                raise RuntimeError("fixture readiness timeout")
            time.sleep(0.2)
        migration = (ROOT / "migrations/000001_identity_session.up.sql").read_text()
        # Password is generated hex, never user-controlled SQL; sent privately on stdin.
        sql = """REVOKE ALL ON DATABASE c1_fixture FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM PUBLIC;
CREATE SCHEMA c1_identity AUTHORIZATION postgres;
SET search_path=c1_identity;
""" + migration + f"""
CREATE ROLE c1_runtime LOGIN PASSWORD '{runtime_password}' NOSUPERUSER NOCREATEDB NOCREATEROLE NOREPLICATION NOBYPASSRLS;
GRANT CONNECT ON DATABASE c1_fixture TO c1_runtime;
GRANT USAGE ON SCHEMA c1_identity TO c1_runtime;
GRANT SELECT,INSERT,UPDATE ON persons TO c1_runtime;
GRANT SELECT,INSERT,DELETE ON pending_logins TO c1_runtime;
GRANT SELECT,INSERT,UPDATE ON sessions TO c1_runtime;
"""
        command(["podman", "exec", "-i", name, "psql", "-X", "-v", "ON_ERROR_STOP=1", "-U", "postgres", "-d", "c1_fixture"], data=sql)
        hostport = port.split(":")[1]
        base = f"@127.0.0.1:{hostport}/c1_fixture?sslmode=disable&search_path=c1_identity"
        env["C1_HARNESS_RUNTIME_DSN"] = "postgres://c1_runtime:" + runtime_password + base
        env["C1_HARNESS_PROVISIONER_DSN"] = "postgres://postgres:" + env.pop("POSTGRES_PASSWORD") + base
        env["C1_HARNESS_OWNED_RUN"] = RUN
        env["GOTOOLCHAIN"] = "go1.26.9"
        print("toolchain", command(["go", "version"], env=env), flush=True)
        test = subprocess.run(["go", "test", "-count=1", "-failfast", "-tags=authharness", "-run", "^(TestHarnessHTTPMeMutationLogout|TestHarnessLogoutRejectsNull)$", "-v",
                               "-timeout=90s", "./internal/platform/auth"], cwd=ROOT, env=env, timeout=120)
        return test.returncode
    finally:
        if created:
            labels = json.loads(command(["podman", "inspect", "--format", "{{json .Config.Labels}}", name]))
            if labels.get("c1.auth.run") != RUN or labels.get("c1.auth.owner") != name:
                raise RuntimeError("cleanup refused: resource ownership mismatch")
            command(["podman", "rm", "-f", "-v", "--time", "5", name], timeout=20)
            if command(["podman", "ps", "-a", "--filter", "name=^" + name + "$", "--format", "{{.ID}}"]):
                raise RuntimeError("owned fixture remained after cleanup")
            for mount in metadata.get("Mounts", []):
                if mount.get("Type") == "volume":
                    remaining = subprocess.run(["podman", "volume", "exists", mount["Name"]], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL, timeout=5)
                    if remaining.returncode != 1:
                        raise RuntimeError("owned anonymous volume cleanup not established")
            print("cleanup verified: owned container and anonymous data volumes removed", flush=True)


if __name__ == "__main__":
    def interrupted(signum, frame):
        raise RuntimeError("fixture interrupted")
    signal.signal(signal.SIGTERM, interrupted)
    signal.signal(signal.SIGINT, interrupted)
    try:
        sys.exit(main())
    except RuntimeError as error:
        print(str(error), file=sys.stderr)
        sys.exit(1)
    except subprocess.TimeoutExpired:
        print("harness operation timed out; no credentials emitted", file=sys.stderr)
        sys.exit(1)
