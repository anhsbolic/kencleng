#!/usr/bin/env python3
"""Owned disposable T2 fixture. Never reads .env or inherits application DATABASE_URL."""
import json
import argparse
import os
from pathlib import Path
import secrets
import re
import signal
import subprocess
import sys
import time

RUN_PATTERN = r"RUN-C1-ENG-[A-Z0-9]+(?:-[A-Z0-9]+)*-[0-9]{3}"
SUITES = {
    "t2": ("TestHarnessRuntimePrivileges", "TestHarnessHTTPMeMutationLogout",
           "TestHarnessHTTPSimulatedCallback", "TestHarnessConcurrentPerson",
           "TestHarnessSessionLockContention", "TestHarnessLogoutRejectsNull"),
    "logout": ("TestHarnessHTTPMeMutationLogout", "TestHarnessLogoutRejectsNull"),
}
ROOT = Path(__file__).resolve().parents[1]
FIXTURE_ENV_KEYS = ("DATABASE_URL", "C1_HARNESS_PROVISIONER_DSN", "C1_HARNESS_OWNED_RUN")


def metadata_environment(test_env):
    """Keep fixture credentials and Run identity out of tool metadata commands."""
    env = test_env.copy()
    for key in FIXTURE_ENV_KEYS:
        env.pop(key, None)
    return env


def command(args, *, env=None, data=None, timeout=30):
    result = subprocess.run(args, input=data, env=env, text=True,
                            stdout=subprocess.PIPE, stderr=subprocess.PIPE, timeout=timeout)
    if result.returncode:
        # Tool errors can contain private connection arguments. Do not dump them.
        safe = re.sub(r"[0-9a-f]{64}", "[private value]", result.stderr.strip())
        safe = re.sub(r"postgres(?:ql)?://\S+", "[private DSN]", safe)
        raise RuntimeError(f"{args[0]} {args[1]} operation failed (exit {result.returncode}): {safe}")
    return result.stdout.strip()


def cleanup_fixture(name, run, expected_id=None):
    """Remove only the unique, labelled fixture found by its exact generated name."""
    try:
        listed = command(["podman", "ps", "-a", "--filter", "name=^" + name + "$",
                          "--format", "{{.ID}}"])
    except (RuntimeError, subprocess.TimeoutExpired) as error:
        raise RuntimeError("unresolved fixture resource: exact-name lookup failed") from error
    ids = listed.splitlines() if listed else []
    if not ids:
        if expected_id is not None:
            raise RuntimeError("unresolved fixture resource: previously identified container disappeared")
        return False
    if len(ids) != 1 or not ids[0]:
        raise RuntimeError("unresolved fixture resource: exact-name lookup was ambiguous")
    discovered_id = ids[0]
    if re.fullmatch(r"[0-9a-fA-F]{12,64}", discovered_id) is None:
        raise RuntimeError("unresolved fixture resource: container ID is malformed")
    if expected_id is not None and re.fullmatch(r"[0-9a-fA-F]{64}", expected_id) is None:
        raise RuntimeError("unresolved fixture resource: expected container ID is malformed")
    if expected_id is not None and not expected_id.startswith(discovered_id):
        raise RuntimeError("unresolved fixture resource: container ID changed")
    try:
        current = json.loads(command(["podman", "inspect", discovered_id]))
        if len(current) != 1:
            raise ValueError("inspect result was not unique")
        metadata = current[0]
        labels = metadata["Config"].get("Labels", {})
        actual_name = metadata.get("Name", "").lstrip("/")
        container_id = metadata.get("Id")
        if (not isinstance(container_id, str)
                or re.fullmatch(r"[0-9a-fA-F]{64}", container_id) is None
                or not container_id.startswith(discovered_id)
                or (expected_id is not None and container_id != expected_id) or actual_name != name
                or labels.get("c1.auth.run") != run
                or labels.get("c1.auth.owner") != name):
            raise ValueError("ownership or identity did not match")
    except (RuntimeError, subprocess.TimeoutExpired, ValueError, KeyError, TypeError) as error:
        raise RuntimeError("unresolved fixture resource: ownership or identity could not be verified") from error
    try:
        command(["podman", "rm", "-f", "-v", "--time", "5", container_id], timeout=20)
        if command(["podman", "ps", "-a", "--filter", "name=^" + name + "$",
                    "--format", "{{.ID}}"]):
            raise RuntimeError("resource remained after removal")
        for mount in metadata.get("Mounts", []):
            if mount.get("Type") == "volume":
                remaining = subprocess.run(["podman", "volume", "exists", mount["Name"]],
                                           stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL, timeout=5)
                if remaining.returncode != 1:
                    raise RuntimeError("anonymous volume cleanup not established")
    except (RuntimeError, subprocess.TimeoutExpired, KeyError, TypeError) as error:
        raise RuntimeError("unresolved fixture resource: verified cleanup could not be established") from error
    return True


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--run-id", required=True)
    parser.add_argument("--suite", required=True, choices=SUITES)
    options = parser.parse_args()
    if len(options.run_id) > 128 or re.fullmatch(RUN_PATTERN, options.run_id) is None:
        parser.error("invalid Run ID: expected RUN-C1-ENG-<UPPERCASE-SEGMENTS>-<NNN>, max 128 characters")
    run = options.run_id
    tests = SUITES[options.suite]
    name = "c1-auth-harness-" + secrets.token_hex(6)
    env = os.environ.copy()
    # Do not propagate application/provider credentials into fixture/test children.
    for key in ("DATABASE_URL", "C1_AUTH_TEST_DATABASE_URL", "GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET",
                "C1_HARNESS_RUNTIME_DSN", "C1_HARNESS_PROVISIONER_DSN"):
        env.pop(key, None)
    env["POSTGRES_PASSWORD"] = secrets.token_hex(32)
    runtime_password = secrets.token_hex(32)
    create_attempted = False
    container_id = None
    metadata = None

    def owned_metadata():
        current = json.loads(command(["podman", "inspect", container_id]))[0]
        labels = current["Config"].get("Labels", {})
        if (labels.get("c1.auth.run") != run or labels.get("c1.auth.owner") != name
                or current["Id"] != container_id or current.get("Name", "").lstrip("/") != name):
            raise RuntimeError("resource operation refused: ownership mismatch")
        return current

    try:
        if command(["podman", "ps", "-a", "--filter", "name=^" + name + "$", "--format", "{{.ID}}"]):
            raise RuntimeError("resource collision")
        command(["podman", "image", "exists", "postgres:16"])
        # No volumes, existing network, host .env, or persistent data mounts.
        create_attempted = True
        container_id = command(["podman", "run", "-d", "--name", name, "--label", "c1.auth.run=" + run,
                 "--label", "c1.auth.owner=" + name, "--pull=never", "--publish", "127.0.0.1::5432",
                 "--env", "POSTGRES_PASSWORD", "--env", "POSTGRES_HOST_AUTH_METHOD=scram-sha-256",
                 "--env", "POSTGRES_INITDB_ARGS=--auth-host=scram-sha-256",
                 "--env", "POSTGRES_DB=c1_fixture", "postgres:16"], env=env)
        metadata = owned_metadata()
        print("run", run, "suite", options.suite, "tests", ",".join(tests), flush=True)
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
        env["DATABASE_URL"] = "postgres://c1_runtime:" + runtime_password + base
        env["C1_HARNESS_PROVISIONER_DSN"] = "postgres://postgres:" + env.pop("POSTGRES_PASSWORD") + base
        owned_metadata()
        print("ownership verified before and after provisioning", flush=True)
        env["C1_HARNESS_OWNED_RUN"] = run
        env["GOTOOLCHAIN"] = "go1.26.9"
        print("toolchain", command(["go", "version"], env=metadata_environment(env)), flush=True)
        test = subprocess.run(["go", "test", "-count=1", "-failfast", "-tags=authharness", "-run", "^(" + "|".join(tests) + ")$", "-v",
                               "-timeout=90s", "./internal/platform/auth"], cwd=ROOT, env=env, timeout=120)
        return test.returncode
    finally:
        if create_attempted:
            if cleanup_fixture(name, run, container_id):
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
