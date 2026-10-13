# Independent Testing — T2 auth harness runtime and cleanup

Phase: Independent Testing. Author/role: `PARTICIPANT-C1-ENG-T2-AUTH-HARNESS-RUNTIME-TESTER-001` / Tester. Created/updated: 2026-10-13 (Asia/Jakarta; terminal observation at 06:33:02 UTC). Work Unit/Run: `WU-C1-ENG-001` / `RUN-C1-ENG-TEST-T2-AUTH-HARNESS-RUNTIME-001`. Session: `SESSION-C1-ENG-T2-AUTH-HARNESS-RUNTIME-TESTER-001` (assignment identity). Selected model/reasoning: `gpt-6-sol` / `medium` by Human-approved Invocation and dispatch configuration; exact runtime identity/effort is not independently attested. Target HEAD/branch: `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering`; workflow HEAD: `651ce57f044bff23fba99082322cee6c066d25a6`.

## 0. Sweep Summary

**Claim:** independent execution of the final full-ID cleanup harness's successful real Podman/PostgreSQL fixture path, all six selected `t2` cases, five focused command-boundary mocks, and the current backend `make verify` chain. This is bounded T2 authentication/session evidence, not T2 acceptance.

- Exact inputs: 48/48 assignment-defining SHA-256 bindings matched before and after execution. Scope SHA-256 `b316d1d3e5faba3b1016045457e0a6b47c9f8c026411e995282b838a4610008c`; bindings manifest SHA-256 `8f4857a145ec565f60083b6e8c929e33a80581d6962f5c009c34e89ec45f2c33`. HEAD/branch/workflow matched. Pre-existing dirty files, including untracked `backend/scripts/__pycache__/auth_harness.cpython-314.pyc`, were preserved. No source, test, dependency, scanner policy, Techplan, Product, `.env`, or shared database write was made.
- Current evidence reused: full-ID Build's two-file hash/delta and five Builder-run mocks; exact-delta Review's static no-delete ordering; original runtime-harness Build's earlier real-fixture six-case execution; BACKEND-GATE-002's `go mod verify` on byte-identical `go.mod`/`go.sum`. Those are separate evidence categories and do not substitute for this Run's fresh results.
- Fresh evidence: mock suite 5/5 PASS; one Run-owned local `postgres:16` fixture, six-case `t2` suite PASS, successful-path owned container/anonymous-volume cleanup reported and verified by the runner; post-patch `make verify` exit 0. The fixture exercised actual PostgreSQL/TCP/SCRAM and DB contention. Handler tests used `httptest`; callback Provider was simulated. No application server, Google/provider login, browser, or production HTTPS was exercised.
- Dispatch/readiness: Invocation records exact Human model and one-fixture approvals; the Human task message supplied mechanical dispatch. `go version` returned `go1.26.9 linux/amd64` from the selected Go command. Scanner v1.7.0, Staticcheck, Gosec, Python, Podman and Podman Compose were available. Local `postgres:16` image existence passed under the escalated Podman environment. The sandbox Podman preflight failed before fixture creation because `/run/user/1000/libpod` was read-only; same image-existence check succeeded through scoped harness escalation. The substantive fixture and `make verify` commands ran escalated because that Podman restriction was known. This was environment recovery, not a test assertion failure. Go checksum verification remained enabled; no module input drift required a new `go mod verify`.
- Open boundaries: `B-T2-002`/`B-T2-003`, full R1–R14, configured application-runtime restricted credential/SCRAM, post-lock expiry finding, real Google/provider/browser/HTTPS, integrated frontend/UI/cache isolation, Human semantic walkthrough, T2→T3 readiness and C1 completion remain unestablished. Existing Google/integrated Human holds remain in force.

## 0a. Test Focus Pointer Execution

| Area | Exact Exploration anchor opened | Specialized verification here | Result |
|---|---|---|---|
| Initial Owner/person/auth | `stage-3-solutioning.md#material-questions-and-constraints` items 1 and 3; attribution/authentication timing | Six-case fixture checked restricted runtime identity, real person upsert and session behavior; mocks checked cleanup ownership, while no Owner grant or real provider was exercised | PASS for T2 fixture boundary; R1/R8 and real provider remain open |
| Organization + Owner atomicity | `stage-3-solutioning.md#organization--initial-owner-consistency` | No Organization/Owner transaction exists in this T2 harness scope | Deferred to T3 R3/R6/R11; no atomicity PASS claimed |
| Known conflict | `stage-3-solutioning.md#known-unresolved-representation-conflict` | No guard/operator or representation-conflict path in this T2 harness | Deferred to T3 R7/R12; no guard PASS claimed |

These anchors preserve the concrete rationale: durable person attribution before successful privileged effect, one DB transaction for coherent Organization/Owner success, and no invented conflict detection. The Techplan's refined session-lock and provider risks were applied through the current T2 checklist; no missing pointer affecting this narrow assignment was found.

## 1. Test Coverage

Commands ran from `backend/`, with `PYTHONDONTWRITEBYTECODE=1`:

1. `python3 scripts/test_auth_harness.py`: exit 0; `Ran 5 tests in 0.002s`, `OK`. These mocked subprocess boundaries cover timeout-after-create, nonzero-after-create, absent resource, mismatched owner, and short/malformed inspect IDs with no removal on invalid identity. They are mock evidence only.
2. `python3 scripts/auth_harness.py --run-id RUN-C1-ENG-TEST-T2-AUTH-HARNESS-RUNTIME-001 --suite t2`: exit 0. Exact selected cases, all PASS:

| Case | Observable result and evidence category |
|---|---|
| `TestHarnessRuntimePrivileges` | PostgreSQL 16.14 TCP/SCRAM success and wrong-password SQLSTATE 28P01; `c1_runtime` no memberships, exact nine non-grantable table privileges, forbidden DDL/role/owner/delete denied. Real disposable fixture, not configured application runtime. |
| `TestHarnessHTTPMeMutationLogout` | `httptest` handler and real restricted Store; anonymous/malformed/expired/revoked sessions, Origin and CSRF negatives, body/type and logout/revocation behavior; 13 emitted subtests PASS. |
| `TestHarnessHTTPSimulatedCallback` | Eight emitted simulated-Provider callback subtests PASS: success, browser mismatch, provider/protocol errors, duplicate state/code, malformed state and missing code. No Google exchange. |
| `TestHarnessConcurrentPerson` | Eight real PostgreSQL upserts, observed first-login blocking edge, one stable person and distinct issuer/case boundaries. |
| `TestHarnessSessionLockContention` | Five real PostgreSQL blocker-edge subtests PASS: expiry, revocation, cancellation, logout and rotation. |
| `TestHarnessLogoutRejectsNull` | `httptest` null body rejected with 400 while the real restricted Store session remained valid. |

Package result: `ok github.com/anhsbolic/kencleng/backend/internal/platform/auth 1.484s`. This selected suite does not exercise Organization/Owner formation, guard, real provider, browser, HTTPS, or configured application runtime.

**Fixture/cleanup:** One generated unique fixture `c1-auth-harness-8ac10dd4dba4` used pre-existing local `postgres:16` image ID `88a36c64c1003dad93f56daa12d1f8916ec66d1fa3e5fb1fb0ae7cb77efd56d1`, `--pull=never`, ephemeral loopback TCP port, separate provisioner and restricted runtime credentials passed privately to the child. No host bind mount, shared named volume, Compose network, `.env`, application `DATABASE_URL`, or provider secret was used. The runner reported `ownership verified before and after provisioning`. In its current source, `cleanup_fixture` checks the exact Run label, owner label, exact generated name, discovered ID prefix and full 64-character inspected/expected ID before calling `podman rm -f -v` with that full ID; it then checks exact-name container absence and `podman volume exists` exit 1 for each inspected volume mount. This execution ended with `cleanup verified: owned container and anonymous data volumes removed`. The raw full container/volume IDs were not printed or persisted, so this report relies on the runner's observed successful checks plus the independently passing negative mocks; it does not independently attest every abnormal cleanup branch or image registry provenance.

## 2. Error Verification

| Error case | Expected | Actual and boundary |
|---|---|---|
| Short/malformed inspected ID, mismatched owner, absent resource, ambiguous create outcome | Fail closed; no unintended removal | Five mock suite PASS; no-removal assertions for invalid identity; abnormal real Podman branches not exercised |
| Wrong DB password or restricted DDL/role/owner/delete | Deny authorization | Real fixture observed 28P01 and permission denials |
| Expired/revoked/anonymous session; bad Origin/CSRF/body | Safe HTTP rejection without valid session effect | `httptest` subtests PASS against restricted Store |
| Provider/browser mismatch, duplicate state/code, malformed callback | No trusted session from invalid callback | Simulated Provider subtests PASS; real provider unverified |
| Podman sandbox setup | Report tooling restriction and recover only within envelope | Preflight failed before creation; scoped escalated route succeeded, then substantive commands PASS |

## 3. Final Verification

`make verify` ran once from `backend/`, exit **0**, with `PYTHONDONTWRITEBYTECODE=1`, `GOTOOLCHAIN=go1.26.9`, `GOFLAGS=-mod=readonly`, Run-scoped writable `GOCACHE`/`STATICCHECK_CACHE` under `/tmp`, and `COMPOSE='podman-compose --env-file /dev/null'`. The ordered chain completed:

| Constituent | Actual result |
|---|---|
| `staticcheck ./...` | Completed without diagnostics |
| Gosec policy gate | Raw exit **1**, **3 raw findings**, 3 accepted, effective PASS |
| `go test ./...` | Auth package PASS 0.234s; server/db report no test files |
| `go test -race ./...` | Auth package PASS 1.206s; no race detector failure; server/db no test files |
| `go test -tags=contract ./...` | Auth package PASS 0.246s; server/db no test files; no separate live API conformance claim |
| Gitleaks history | Raw exit **10**, **4 raw findings**, 4 accepted, effective PASS |
| Gitleaks worktree | Four scans indexed 0–3, all raw exit 0 / 0 findings / effective PASS. Indices 0–1 were unstaged/staged; indices 2–3 covered Git-visible untracked `backend/scripts/__pycache__/auth_harness.cpython-314.pyc` and `backend/scripts/test_auth_harness.py` per current inventory/source ordering. The pre-existing pycache was included and preserved. |
| `govulncheck ./...` | Completed plain source scan: `No vulnerabilities found.` No scanner loading/network error observed. |

Accepted Gosec/Gitleaks raw findings remain findings; effective policy PASS does not mean zero raw findings. The current Makefile's sequential continuation and parent exit 0 establish constituent completion; no separate raw scanner captures were retained because the policy gate printed safe result-bearing counts, and this Run did not change their established dispositions. `go.mod`/`go.sum` matched the prior `go mod verify` inputs exactly; dependency integrity was reused, not rerun or expanded to full origin/tag provenance.

- Broad checks intentionally not rerun: scanner helper negative/error matrix, `go mod verify`, API/frontend gates, and real provider/browser checks; their evidence ownership or unchanged inputs are stated above. No unrelated suite was run.
- Migration/schema collision: identity fixture migration executed successfully in its fresh schema; broader C1 migration/Organization schema collision unverified.
- Backward compatibility: current Go unit/race/contract commands and six-case T2 harness PASS. Deployed application/health/browser compatibility unverified.
- Fresh entire Approved Techplan and T2 task read: no contradiction invalidated this bounded assignment. The plan still requires later Organization/Owner, guard, browser/rendered/Human and provider evidence; historical plan prose did not become current acceptance. Existing B-T2 blockers and holds were not reclassified.

## 4. New Recurring Bug Patterns

None found. Successful fixture cleanup is a path-specific observation; the negative paths remain mocked.

## Verdict

**Pass for this bounded independent runtime/cleanup and current backend command-gate claim.** All three declared commands completed successfully, the owned successful fixture cleanup was reported and checked by the runner, and the five mocks independently exercised the guarded no-delete cases. No production fix or patch plan is required by this Run. This verdict does not close `B-T2-002`/`B-T2-003`, establish T2/T3 readiness or acceptance, change either Human hold, or attest exact runtime model identity.

## Phase handoff

- **Outcome:** COMPLETED — bounded Testing Pass.
- **Result refs:** This `evidence/testing-report-1.md`; exact current scope/manifest and two harness-file identities in this Run's Invocation/input bindings.
- **Findings:** No new substantive defect; sandbox Podman preflight restriction recovered before fixture creation. Raw scanner findings remain accepted under existing gate policy, not erased.
- **Decision requests:** None for executed scope.
- **Blockers:** None for this bounded claim; `B-T2-002`/`B-T2-003` retain their owning wider scope and status.
- **Open / unverified:** Abnormal real cleanup branches; configured application runtime, real provider/browser/HTTPS, Organization/Owner/guard/integrated/Human and full dependency provenance obligations; no T2/T3 acceptance.
- **Recommended continuation:** Orchestrator reconcile this exact Testing result against the remaining T2 obligations and holds; route any subsequent scoped verification or Build/Patch only under its own Run and authority.
- **Context refs:** Current Invocation/scope/bindings; Approved Techplan/T2 task; terminal full-ID Build and Review reports; original runtime-harness Build report; BACKEND-GATE-002 report; current backend Makefile/README.
