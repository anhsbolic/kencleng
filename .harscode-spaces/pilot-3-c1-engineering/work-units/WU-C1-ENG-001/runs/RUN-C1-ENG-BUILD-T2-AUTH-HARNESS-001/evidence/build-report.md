# T2 auth HTTP/PostgreSQL harness Build report

Phase: Build. Author: Codex / Implementer, backend auth-test/DB-fixture.
Created/Updated: 2026-10-12 (terminal evidence assembled after 15:07 UTC).
WORK_UNIT_ID: WU-C1-ENG-001. RUN_ID: RUN-C1-ENG-BUILD-T2-AUTH-HARNESS-001.
Participant: PARTICIPANT-C1-ENG-T2-AUTH-HARNESS-IMPLEMENTER-001.
Session binding: SESSION-C1-ENG-T2-AUTH-HARNESS-IMPLEMENTER-001 / FRESH; logical invocation binding, not an independently observed runtime thread identifier.
Model/reasoning: invocation selects and approves gpt-6-sol / medium; exact runtime model/reasoning not independently exposed to this execution. Approval is not runtime identity evidence.
Target revision: dae55ac44beb0e967b94c2f791d8a0b8d8e72372, pilot/3-c1-engineering (observed).
Workflow revision: 651ce57f044bff23fba99082322cee6c066d25a6 (observed).

## What changed

Auth tests and owned fixture were authored and executed. Overall result **FAIL**, because the retained object-only logout oracle observes HTTP 204 for JSON `null`, rather than required HTTP 400. No production fix was made. Five preceding top-level checks passed; this is Build-owned executability evidence, not independent runtime acceptance.

Exact Run delta:

| File | Before SHA-256 | After SHA-256 |
|---|---|---|
| backend/internal/platform/auth/harness_test.go | absent | bdbfe0b59267125de9572ff534eb6839650b7ace1a3ecacd5965d22219c28552 |
| backend/scripts/auth_harness.py | absent | 01a2e72822f990528a9d3831430c42aa4685a9581b95ac1b5374bd5116d7a6f6 |
| backend/README.md | 5bcea466e7eaf011d3395d2393a177f3c75e61f2c03a3198e081dd4df31d9f99 | 38b5623e84b9b2fef704978b8f3100fd68e51d2b35f14fe939e02893ba85d99f |

Tests use the separate `authharness` opt-in tag. Runner provisions exact current migration in its own schema, a provisioner and restricted runtime; bounded first-login/session contention oracles use PostgreSQL blocker edges. README explains execution, known failure, evidence limits, and new-Run rebinding after terminality. No Makefile change in this Run.

Existing dirty/untracked work was present at kickoff, including Makefile/README/scanner and orchestration artifacts; it was preserved. Production auth/store/db/server, migration semantics, dependency pins/lockfiles, API, frontend, Compose, scanner policy, Product/Techplan/task and coordination projections were not modified by this Run. No commit or messages to other participants.

## Preflight / approvals / bindings

- `input-bindings.json` SHA-256: `3d41697b6b4ffa90355c11a98572e790f7477841574e4585c9f1b648f83370d1`. All 20 assignment-defining inputs matched before writes/provisioning, including scope `bab8e336693e4e5bb325b3694548294f301bbe42a7b333af0f41f16032588bbe`.
- Both exact approvals verified in invocation and WU Events entry “Auth harness exact model and disposable fixture authorization approved” (2026-10-12, line 137): Anhar's combined “approve bro” applies to this Run's gpt-6-sol / medium and exact disposable provisioning scope. It excludes persistent/shared DB, host `.env`, wider privilege and production auth writes. Current Human instruction dispatches this occurrence.
- Read root/backend instructions, canonical Build prompt/guidelines/checklist, orchestrated overlay/run-contract, Approved parent/T2, current API and live Handler/Store/Protocol/db/migration/tests; routed integration/concurrency/PostgreSQL lock, token/privilege, cookie/CSRF/CORS and secret logging guidance. Existing command-gate report is reused for its original bounded claim only.
- End check: 19 bound inputs byte-unchanged; only README changed within allowed scope. Full before/after identities are in `file-identities.json`; no approval or authority artifact edited.

## Tests run

Focused authored checks were invoked through `python3 backend/scripts/auth_harness.py` from repository root. Child command in backend:

```text
GOTOOLCHAIN=go1.26.9 go test -count=1 -failfast -tags=authharness -run ^TestHarness -v -timeout=90s ./internal/platform/auth
```

Observed actual toolchain: `go version go1.26.9 linux/amd64`. Podman 5.7.0, API 5.7.0, built with Go1.25.0. PostgreSQL 16.14 (Debian 16.14-1.pgdg13+1), x86_64. Actual local image ID `88a36c64c1003dad93f56daa12d1f8916ec66d1fa3e5fb1fb0ae7cb77efd56d1`, route `postgres:16`, `--pull=never`; this is an observed local image identity, not a registry digest/provenance attestation. No dependency/image installation or version change.

| Authored test | Boundary / result |
|---|---|
| TestHarnessRuntimePrivileges | Real PG TCP: PASS. Runtime identity and five restricted role flags; no membership; SCRAM stored-verifier predicate and all host HBA SCRAM rules; successful pgx TCP pool plus wrong-password 28P01; exact nine non-grantable table grants; DDL/schema CREATE/role CREATE/SET ROLE/ALTER and forbidden DML denied 42501; attempted self-GRANT leaves exact grants unchanged. |
| TestHarnessHTTPMeMutationLogout | httptest + restricted real Store: PASS. Anonymous/malformed/expired/revoked 401, valid local-only me and no-store; exact Origin/JSON/session/CSRF negatives and valid charset control; invalid/extra/multiple/oversize/array/string logout bodies 400 without revocation; object logout 204/deletion cookie/DB revocation/repeat 401; closed-pool storage error 503 safe envelope; no implicit credentialed CORS. |
| TestHarnessHTTPSimulatedCallback | Simulated Protocol + restricted real Store: PASS. Start browser cookie, nonce/verifier bound across simulated exchange; fixed local callback destination, no attacker return URL reflection; success maps local identity and revokes previous session; wrong browser/provider-error/protocol-failure/duplicate-state/duplicate-code/malformed-state/missing-code mint no session. Consumed-state replay does not reach Protocol again. Provider-error scenario uses error+state without code. |
| TestHarnessConcurrentPerson | Real PG runtime contention: PASS. First uncommitted runtime INSERT holds unique identity; eight upserts start behind a channel barrier; actual blocker edge witnessed before release; all return same person, one stored row; case/issuer identities remain distinct. |
| TestHarnessSessionLockContention | Real PG runtime row contention: PASS for expiry, revocation, cancellation, logout and rotation. Actual `pg_blocking_pids` edge before release, current public LockSession, zero authority on invalid/cancelled wait; old token invalid after serialized logout/rotation; new rotation token maps same person. Expiry release follows stored deadline after the blocker edge; no fake clock. |
| TestHarnessLogoutRejectsNull | httptest + restricted real Store: FAIL. Valid Origin/JSON/session-bound CSRF, body `null`; observed HTTP 204, expected 400. Failfast stopped further behavior verification. |

Test process and wrapper exit 1; test execution duration reported 1.456s. Failed assertion was retained unchanged, not excluded or weakened.

Additional local authoring checks: `gofmt -w` on authored Go test, Python `py_compile` (generated bytecode removed), `git diff --check` passed. These do not constitute the final command gate; no broad gate/scanner rerun after substantive finding.

### Environment recovery and fixture lifecycle

1. Initial sandbox `podman version` failed: `Failed to obtain podman configuration: set sticky bit on: chmod /run/user/1000/libpod: read-only file system`. One same-command scoped harness escalation succeeded. No system permission or cache configuration changed.
2. First fixture `c1-auth-harness-c75a347fd081`: loopback disposable launch succeeded; `podman exec ... psql` setup failed exit 2, before any Go test. Cleanup ran and container removal was subsequently independently confirmed. Original wrapper did not preserve psql stderr; exact underlying cause is **not verified**. Temporary-server/socket readiness race was a hypothesis from runner inspection. This attempt does not establish schema/runtime evidence. Anonymous-volume absence was not individually queried for this first attempt; removal used ownership-checked `podman rm -f -v --time 5`.
3. Authored runner changed readiness from Unix-socket default to `pg_isready -h 127.0.0.1`, selecting the final TCP server; added sanitized operation-error detail and individual volume-existence postcheck. Bounded active fixture iteration, no substantive source finding at that point. Second scoped invocation of the wrapper succeeded in setup; captures remain distinct.
4. Second fixture `c1-auth-harness-880b836435e1`: separate own container/anonymous storage, random ephemeral loopback host port, no bind mounts/host `.env`/shared named volume/existing database. Labels `c1.auth.run=RUN-C1-ENG-BUILD-T2-AUTH-HARNESS-001` and `c1.auth.owner=<exact container name>` checked before removal. No custom network created. Provisioner `postgres` applies exact current migration in `c1_identity`, creates role and grants. Generated test-only passwords sent through private process environment/stdin, never command/evidence DSNs or password values. No credential files created.
5. Runtime `c1_runtime`: NOSUPERUSER/NOCREATEDB/NOCREATEROLE/NOREPLICATION/NOBYPASSRLS. Database CONNECT and schema USAGE only; no PUBLIC database/schema privileges, no schema CREATE/owner membership. persons SELECT/INSERT/UPDATE; pending_logins SELECT/INSERT/DELETE; sessions SELECT/INSERT/UPDATE. Every actual Handler/Store/competing lock operation uses runtime connections. Provisioner queries role/grant/HBA metadata, observes blocker edges and seeds the already-expired HTTP fixture only; those operations are not runtime authorization evidence. Session-lock holder is also runtime; its expiry/revocation update is explicitly fixture state orchestration.
6. After assertion failure: ownership-checked `podman rm -f -v --time 5` returned success; no exact-name container remained, individual second-attempt anonymous volumes returned `podman volume exists` exit 1. Subsequent read-only exact-name checks for both containers returned no IDs, exit 0. No shared resources selected or removed. Go child had already exited before cleanup; wrapper session exited 1. Failure cleanup was exercised; forced interruption/90s timeout cleanup was implemented but **not fault-injection tested**. No temporary credentials/data files remain; private child process environments ended.

## Verification scope confirmation

Real DB concurrency and focused security-class runtime privilege/SCRAM/HTTP/CSRF checks were executed as the invocation's explicit Build exception: newly authored oracles needed execution to establish they work against current T2. They remain Build-owned evidence. No race detector, performance/load suite, broad final T2 matrix, production server, real Google login/network, browser or HTTPS deployment check was executed. Existing cookie/config/protocol tests were read and reused as current narrower evidence rather than replayed.

Real category: fixture PostgreSQL TCP/query/grants/row blocking/current migration. Simulated category: Protocol returns a fixture identity; no real provider verification/code replay interoperability. Local category: registered Handler httptest HTTP responses and serialized development cookie headers. Reported/reused category: pre-existing final command-gate evidence. Not tested category: independent final acceptance, deployed application's actual credential configuration, actual SCRAM wire transcript capture, real Google/browser/production HTTPS/T3 aggregate/guard/operator/rendered/cache/Human meaning.

## Contract check

- [ ] Current build target satisfied in full — not green; null-body finding blocks claiming complete harness/T2 HTTP contract compliance.
- [ ] Live-code re-grounding did not invalidate a material contract assumption — actual invalid-input execution demonstrates object-only contract violation.
- [x] Production source and upstream contract preserved; expected result derived from active API, not adjusted to source behavior.
- [x] Bindings and both approvals verified; owned disposable fixture only; secrets not emitted.

## Deferred / not tested here

Independent Review/Testing applicability is for Orchestrator reconciliation of this actual delta/results. Whole-T2 command gate rerun, dependency commit, T2→T3 satisfaction, T3 dispatch, Google/browser/HTTPS and other integrated verification remain unclaimed. B-T2-002/B-T2-003 remain open. Runtime deployment configuration cannot be inferred from fixture SCRAM success. README/runner remain bound to this closed occurrence; future execution requires fresh Run provenance, not reopening this terminal Run.

Other unexecuted specific extensions: repeated successful callback stable-person oracle across two callback sessions, duplicate-error/missing-browser/expired-pending callback cases, production Handler-cookie serialization beyond existing reviewed cookie oracle, mutation storage-failure and revoked/expired mutation variants beyond me checks, forced abort cleanup injection. No claim that the authored cases exhaust all T2 obligations. No further tests after material failure.

## Flagged for Techplan / Testing

**F-AUTH-HARNESS-001 — logout accepts JSON null contrary to required object.**

Authority: `api/openapi/index.yaml:74`, required application/json object, additionalProperties false, InvalidRequest 400. Oracle: `backend/internal/platform/auth/harness_test.go:596`. Raw result: `harness-execution.txt` records `HTTP status: got 204 want 400`. Source: `backend/internal/platform/auth/http.go:259` decodes into map; JSON null yields nil map, len zero, decoder next yields EOF, and path proceeds to RevokeSession. This explanation is source-based. The failing test stops at status mismatch, so its later explicit DB non-revocation oracle was not reached; it does not separately record null-case revoked_at/cookie evidence. Successful `{}` logout and its DB revocation were independently asserted in the preceding passing test.

This is input-contract failure under an otherwise authenticated and CSRF-valid request; the evidence does not establish an Origin/CSRF bypass or privilege escalation. Smallest proposed patch: reject null/non-object before invoking revocation, preserving valid empty-object behavior and existing safe error contract; retain/regress with the authored null oracle. Owner: Orchestrator routes separately authorized production Build/Patch and then determines focused independent Review/Testing. No Product/API semantic change is needed or authorized here. No source patch, acceptance or existing blocker closure in this Run.

Evidence content identities:

- `harness-setup-attempt-1.txt`: `0a41ccb9ad0331a51d34553d05910e58e73b2b0d33fdbd9d27f4d0829c2f6e63`.
- `harness-execution.txt`: `95c090b8a795a8726f3eaa1ebc4ed1eac9967e1016e71ad7af5587b1fb298ee7`.
- `file-identities.json`: `8df25ecdcb1ef02baae7d216fccb0ae72a26d2eceeff0cf8a650cd10b349d10f`.

## Phase handoff

- Outcome: FAILED — authored focused harness executed; object-only logout regression fails; fixture cleaned; occurrence closes here.
- Result refs: this `evidence/build-report.md`; `harness-execution.txt`; `file-identities.json`; authored test/runner/README identities above.
- Findings: F-AUTH-HARNESS-001, null logout HTTP 204 versus 400; no suppressed assertion.
- Decision requests: Orchestrator route bounded production patch authority and exact independent phase applicability; no wider fixture request.
- Blockers: null-body defect blocks green authored harness/HTTP contract compliance claim. Safe unaffected work: inspect passing evidence and review test-only delta; no shared DB/provider action needed.
- Open / unverified: B-T2-002/B-T2-003, whole-T2 acceptance/dependency/T3 conditions and deferred checks above; first setup root cause and first anonymous-volume absence not independently established.
- Recommended continuation: fresh authorized Build/Patch for null handling, retaining this failure evidence, then claim-scoped independent Review/Testing selected by Orchestrator. Never resume this terminal occurrence or mask null assertion for PASS.
- Context refs: invocation.md + scope.md + Approved techplan/T2 + active API logout contract + changed tests/runner + current http.go logout only.
