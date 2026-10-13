# Scope — RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001

Status: PREPARED — exact Run-owned disposable PostgreSQL fixture authorization is recorded in the Invocation; Human mechanical dispatch is pending. No source, test, fixture, or runtime operation has occurred for this Run.

## Bounded outcome

Adapt the existing opt-in auth harness so its restricted runtime pool is created through the production `DATABASE_URL` → `db.Open` path, while `DATABASE_URL` is populated only inside this Run's disposable-fixture child process. Preserve the six existing `t2` test cases and their assertions, including real PostgreSQL lock-wait expiry behavior. This supplies an executable oracle for a later independent Testing Run; Build results are not independent acceptance.

## Allowed writes

- `backend/internal/platform/auth/harness_test.go` — use production `db.Open` for the restricted runtime pool, sourced from the private test-child `DATABASE_URL`; keep provisioning/witness operations on their separate fixture-only identity.
- `backend/scripts/auth_harness.py` — pass the generated restricted-runtime DSN as `DATABASE_URL` only to the Go test child; continue removing any inherited application/provider/test DSNs before fixture setup and never read `.env`.
- `backend/README.md` — describe the new configured-pool path and its evidence limits.

Do not change production auth, Store, `db.Open`, server, API/OpenAPI, migration, frontend, dependency pins, Makefile, scanner policy, Product/Techplan/task authority, other Run evidence, shared DB, or `.env`. Preserve all pre-existing dirty and untracked workspace content.

## Fixture and exact execution

After this exact Run's fixture authorization and Human dispatch, use only the existing locally available `postgres:16` image through the configured Podman path. Create one unique Run-owned disposable container with loopback ephemeral TCP, anonymous storage, a fixture provisioner, and the existing restricted `c1_runtime` role/grants. Generate credentials for this fixture and pass them privately to child processes. Set the child `DATABASE_URL` to the restricted role's disposable fixture DSN; never inherit the host application's `DATABASE_URL` or load `.env`. Do not print DSNs, passwords, tokens, environment dumps, or sensitive payloads.

Execute only:

```text
python3 scripts/auth_harness.py --run-id RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001 --suite t2
```

from `backend/`, with the existing Run-owned runner cleanup path and bounded timeouts. The six selected cases must exercise a pool created through `db.Open` using the child `DATABASE_URL`, including actual session queries and the real PostgreSQL lock-wait expiry assertion. Report PostgreSQL/TCP/SCRAM evidence separately from httptest and simulated Provider evidence. Verify this Run's container/anonymous-volume cleanup through the runner. If the configured pool or required assertion cannot be established within these files, stop and report the exact gap; do not change production semantics or expand fixture privileges.

No `make verify`, scanners, dependency verification, real Google login/provider exchange, browser, deployed listener, production HTTPS, shared/persistent DB, sourceguard allowance, commit, or PR. No fixture is created before authorization and dispatch.

## Re-entry and limits

This is a test-infrastructure Build/Patch Run, not a Testing Run. A later fresh independent Testing Run must own a new fixture and independently exercise the resulting configured-pool path. Orchestrator will determine Review applicability from the exact delta and inspect whether the independent run can cover both R14 and B-T2-003. This Run cannot close `B-T2-002` or `B-T2-003`, authorize dependency acceptance/commit, establish T2/T3 acceptance/readiness, release either Human hold, or authorize T3 dispatch.

## Output

Write only this Run's `evidence/patch-report-1.md` as Run evidence. Record exact bindings, changed-file hashes/delta, the actual selected command and result, the pool configuration path, actual DB versus simulated evidence, fixture cleanup evidence/limits, not-run checks, and exactly one `## Phase handoff`. Preserve historical terminal artifacts unchanged.
