# Backend runtime test harness Build report

Phase: Build. Author: Codex / Implementer. Created/Updated: 2026-10-13 Asia/Jakarta (execution completed before 01:08:19 UTC).
WORK_UNIT_ID: WU-C1-ENG-001. RUN_ID: RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001.
Participant / Session binding: PARTICIPANT-C1-ENG-T2-RUNTIME-HARNESS-IMPLEMENTER-001 / SESSION-C1-ENG-T2-RUNTIME-HARNESS-IMPLEMENTER-001; fresh assignment context, not independently attested runtime thread IDs.
Selected model / reasoning: gpt-6-sol / medium per Invocation and exact Human approval; exact runtime model/effort is not independently exposed. Approval is not runtime attestation.
Observed target HEAD / branch: `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering`. Prepared target: `dae55ac44beb0e967b94c2f791d8a0b8d8e72372`; 34/34 assignment-defining content bindings matched despite HEAD difference.
Observed workflow revision: `651ce57f044bff23fba99082322cee6c066d25a6`, matching Invocation.

## What changed

**PASS for scoped harness Build executability.** Only the three authorized backend files changed:

- `scripts/auth_harness.py`: required Run ID validated against `RUN-C1-ENG-[A-Z0-9]+(?:-[A-Z0-9]+)*-[0-9]{3}` with 128-character maximum; required allowlisted `t2`/`logout` selector. Validation precedes resource commands. Exact Run ID flows through labels, child environment, and ownership checks. Exact unique owner label and returned immutable container ID are checked before provisioning, after provisioning, and before deletion; deletion uses that ID. `t2` selects exactly the six existing cases; `logout` retains the two-case subset. Assertions, failfast, return status and bounded timeouts remain.
- `internal/platform/auth/harness_test.go`: validates supplied Run identity using the same grammar/length rather than a closed Run's ID. Reversing only the new import/marker delta reproduces the bound preimage hash: all test assertions are byte-preserved.
- `README.md`: documents exact Run invocation, suite boundaries, ownership, timeout and evidence limits. A later Run must use its own ID and fixture approval.

Before/after SHA-256 identities: `file-identities.json`. Initial 34/34 bindings and scope hash `d86d6474409252dddd6db2a38683b74ba6c9ee1c2df256ed306d4ed9fb15ffda` matched. Postflight: all other 31/31 bound inputs unchanged. Existing dirty WU, coordination, events, Invocation and Product control-tower work was preserved; none was authored here. Production/auth/store/db/server, migration, API, dependencies, scanner policy, Makefile, frontend and upstream authority were untouched. No commit/PR or next Run was created.

Read Invocation/Scope first, then current canonical Build guidance/overlay/run contract/runtime policy, Approved Techplan/T2, root/backend instructions, verification environment guide, live runner/tests/README, Handler/Store/migration and logout OpenAPI object contract, and bound terminal Build/Review/Testing reports. No material assignment contradiction was found. Invocation records both exact approvals; this Human task message dispatches the prepared occurrence. Scope's stale pending-approval wording is reconciled by the current Invocation approval record; no lifecycle file was edited.

## Tests run

Command from `backend/`:

```text
python3 scripts/auth_harness.py --run-id RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001 --suite t2
```

First sandbox attempt: exit 1 at initial `podman ps`, before any fixture creation. Podman could not set the sticky bit on `/run/user/1000/libpod` (read-only filesystem). Same exact command rerun through scoped harness escalation: exit **0**. This was environment recovery before provisioning, not an isolation/cleanup failure or discarded test assertion. No substantive failure was suppressed.

Child command: `go test -count=1 -failfast -tags=authharness -run ^(TestHarnessRuntimePrivileges|TestHarnessHTTPMeMutationLogout|TestHarnessHTTPSimulatedCallback|TestHarnessConcurrentPerson|TestHarnessSessionLockContention|TestHarnessLogoutRejectsNull)$ -v -timeout=90s ./internal/platform/auth`. Child process timeout 120s; setup operations bounded at 30s, readiness 30s with 5s probes, cleanup 20s with 5s volume checks.

Actual toolchain: **go1.26.9 linux/amd64**. Actual PostgreSQL: **16.14 (Debian 16.14-1.pgdg13+1)**. Fixture `c1-auth-harness-18a1e2703f98`; local `postgres:16` image ID `88a36c64c1003dad93f56daa12d1f8916ec66d1fa3e5fb1fb0ae7cb77efd56d1`, `--pull=never`. This identifies the local image, not registry provenance.

| Existing top-level case | Actual outcome / evidence category |
|---|---|
| TestHarnessRuntimePrivileges | PASS, 0.02s. Real PostgreSQL TCP/SCRAM success, wrong-password 28P01, restricted c1_runtime role flags, zero memberships, exact nine non-grantable table privileges; forbidden DDL/role/owner/DML operations denied. |
| TestHarnessHTTPMeMutationLogout | PASS, 0.03s. Registered handler through httptest and real restricted Store; session, Origin/CSRF/body, me/no-store/safe errors, valid object logout and revocation assertions. |
| TestHarnessHTTPSimulatedCallback | PASS, 0.06s. Eight local handler scenarios with simulated Protocol and real Store; no external Google exchange. |
| TestHarnessConcurrentPerson | PASS, 0.04s. Eight real runtime upserts; actual PostgreSQL first-login blocker edge witnessed; stable single person and distinct issuer/case boundaries. |
| TestHarnessSessionLockContention | PASS, 1.18s. Five real PostgreSQL blocker-edge scenarios: expiry, revocation, cancellation, logout and rotation. |
| TestHarnessLogoutRejectsNull | PASS, 0.01s. httptest null returns 400 with session still valid in real restricted Store. |

Package result: `ok github.com/anhsbolic/kencleng/backend/internal/platform/auth 1.350s`; all six selected cases actually ran, including all emitted subtests. Build execution proves reused tests and new ownership/suite plumbing are executable; it is not independent Testing acceptance.

Focused authored-file checks: Python AST parse PASS; three pre-resource rejection checks PASS (missing arguments, invalid ID, arbitrary suite regex), with a resource-command tripwire that was never reached. `gofmt` completed; `git diff --check` PASS. No extra fixture was used for these checks.

## Fixture isolation and cleanup

Exactly one new Run-owned ephemeral container, fresh generated passwords, private child environment/stdin, separate postgres provisioner and c1_runtime login. No `.env` or application DATABASE_URL access, host bind mounts, shared named volumes or Compose network. Loopback ephemeral TCP mapping was checked. Runtime remained NOSUPERUSER/NOCREATEDB/NOCREATEROLE/NOREPLICATION/NOBYPASSRLS; only SELECT/INSERT/UPDATE on persons and sessions plus SELECT/INSERT/DELETE on pending_logins, without grant option or owner membership, were verified.

Observed runner output: `ownership verified before and after provisioning`. Finally, exact Run/owner labels and container ID were rechecked, the exact container removed with anonymous-volume removal, name lookup returned empty and every recorded volume existence check returned 1. Final output: `cleanup verified: owned container and anonymous data volumes removed`. Cleanup established for this successful execution; forced interruption, provisioning timeout and ownership mismatch paths were not exercised. No unrelated resource cleanup occurred. DSNs, passwords, tokens, provider credentials or private environment were not persisted.

## Verification scope confirmation

The explicitly assigned six-case suite includes real DB concurrency/locking and restricted-role/SCRAM checks. These specialized checks ran only to establish the reusable harness's Build executability, as authorized by this Invocation/Scope; they are not a scope deviation. No race detector, performance/load sweep, scanner, broad `make verify`, real provider/browser/HTTPS or application startup ran.

## Contract check

- [x] Current bounded Build target satisfied; exactly three authorized code/documentation files changed plus this Run's evidence.
- [x] Live-code re-grounding did not invalidate a material contract assumption; logout remains object-only and existing assertions were preserved.
- [x] Full six-case suite ran under this Run's disposable ownership and successful cleanup.

## Deferred / not tested here

Independent runtime acceptance, fresh post-delta command gates, real Google/Human browser/HTTPS, deployed runtime credentials, integrated journey, full provenance, and existing broader T2/T3 obligations remain unestablished. Review R1's optional negative Set-Cookie assertion was deliberately not added. The historical AUTH-HARNESS-001 FAIL, LOGOUT-PATCH-001 Build PASS, logout Review 0 blocking/1 nonblocking, and BACKEND-GATE-002 command PASS are preserved at their original boundaries. Their terminal Runs were not reopened. BACKEND-GATE-002 is not automatically current for this changed backend tree. No blocker closure, T2→T3 dependency satisfaction, readiness or acceptance claim follows.

## Flagged for Techplan / Testing

No new material Product/Techplan contradiction or test failure. Orchestrator owns exact-delta Review applicability and fresh independent Testing/command-gate routing; future fixture authorization and ownership are Run-specific.

## Phase handoff

- **Outcome:** COMPLETED — scoped harness Build executability PASS.
- **Result refs:** `evidence/build-report.md`; `evidence/file-identities.json` for exact changed-file identities.
- **Findings:** no new substantive failure; initial sandbox Podman restriction recovered before resource creation.
- **Decision requests:** none within this completed Build scope; downstream applicability/routing remains Orchestrator-owned.
- **Blockers:** none for this harness Build; existing wider blockers unchanged.
- **Open / unverified:** independent Testing acceptance and fresh command gates; provider/browser/HTTPS/deployed runtime/integrated/provenance obligations; nonblocking Review R1; abnormal cleanup paths.
- **Recommended continuation:** Orchestrator reconcile exact delta, decide independent Review applicability and prepare fresh independent Testing with its own approved fixture owner and applicable post-Build command gates.
- **Context refs:** this Run's Invocation/Scope/input-bindings; Approved Techplan/T2; three changed backend files; preserved bound logout Build/Review and BACKEND-GATE-002 reports.
