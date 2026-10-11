# T2 backend — session retest

Phase: Testing diagnostics; Author: Codex; Created: 2026-10-11 Asia/Jakarta.
Target HEAD: `059d145d5df64fcda427e5a625486a8e95ce1fd8`, with existing uncommitted changes. Runtime model/session identity is not independently exposed. Authorization: Human requested backend retest, updated report and handoff in this session. This is supplemental session evidence, **not a dispatched independent orchestrated Testing Run or T2 acceptance**.

## 0. Sweep Summary

The terminal `RUN-C1-ENG-TEST-T2-001` report is preserved. Rechecked its 33 input hashes: 31 still match; backend Makefile and README changed during authorized tooling/configuration follow-ups. Consequently the old Invocation fails its exact-input preflight and cannot be reused. No new Run/model approval/Participant identity was invented. This report supplies bounded session diagnostics for the next owning route.

The current Approved Techplan was read, including runtime contract, Rules R1–R14, Testing Checklist and Test Focus Pointer. Exact Build retry-1 and Review T2-002 evidence were reopened; their results are prior claims, not new executions. Configuration-presence checks found both Google fields and both DB URLs populated; values are excluded from evidence. Production code and dependency pins were not changed.

Closed practical gaps: real serial PostgreSQL integration execution now passes using the configured fixture DB; unit/race checks, contract-tag command, secret scan and fresh dependency scan actually executed. Successful fixture cleanup was checked. Full runtime acceptance remains incomplete.

## 0a. Test Focus Pointer Execution

| Area | Pointer/context | Verification | Result |
|---|---|---|---|
| Person/auth attribution | Techplan §12 person/auth pointer; exact rationale previously opened in TEST-T2-001 | Existing serial PostgreSQL upsert and binding/replay tests; existing provider negatives | DB tests PASS; provider negatives PASS with **simulated provider**. No fresh real Google identity evidence. Exploration anchor was not reopened this diagnostic round; no complete formal pointer execution claimed. |
| Session lock timing/contention | Current Techplan §§9/12; Build retry-1; Go concurrency/integration and PostgreSQL locking guidance | Unit lock-wait/time oracle and race detector | PASS for simulated row lock/time regression and process race coverage; real row-lock contention **NOT RUN**. |

Organization/Owner and guard anchors remain outside this bounded backend T2 diagnostic scope.

## 1. Test Coverage

| Rule / scenario | Category | Actual verification | Result |
|---|---|---|---|
| R1/R9 immutable person mapping | Serial real PostgreSQL | `TestPersonUpsertPostgreSQL`: repeat identity, case-sensitive subject, distinct issuer, parameter safety and empty identity | **PASS**, real fixture DB. Fixture login owns disposable schemas; does not prove distinct restricted application runtime privileges or concurrency. |
| R9 pending login binding/replay/expiry | Serial real PostgreSQL | `TestPendingLoginBindingAndReplayPostgreSQL` | **PASS**, real fixture DB; local state replay, not provider authorization-code replay. |
| R9 forged issuer/audience/azp/signature/algorithm, expiry/nonce/key/provider failure, PKCE | Negative / unit protocol | Existing six top-level auth tests rerun | **PASS — simulated provider**, not real configured login. |
| R10 session validity after lock wait, expiry equality/revocation/missing/error, stale caller time | Unit regression | Existing fake-row tests | **PASS — simulated lock**, no real lock-order proof. |
| R9/R10 malformed authority input | Unit | Existing rejection-before-storage test | **PASS**. |
| R1/R9 actual Google login twice with stable person; callback/code replay | Real provider/browser | Config fields populated, but client registration, consent/test-user and interactive login not independently established | **NOT RUN/BLOCKED**. No provider/browser launched; credentials being present is insufficient. |
| R10 cookie/CSRF/Origin/logout/rotation at real HTTP/browser boundary | Real interface | No runtime/browser test launched after failed final gate | **NOT RUN**; production HTTPS prerequisite not established. Cookie source inspection is not runtime acceptance. |
| R1/R9/R10/R11 real concurrent first login; logout/rotation/revocation/expiry while row-lock blocked | Real concurrency | Existing integration file provides serial tests only; additional restricted-runtime isolated fixture not provisioned | **NOT RUN**. No fixture-owner result relabeled restricted-runtime contention evidence. T3 confirmation separately deferred. |
| R14 runtime pool/SCRAM/privileges | Runtime compatibility | Integration uses actual pgx pool and configured fixture credentials. Prior setup evidence documents TCP/SCRAM and distinct roles | **PARTIAL**: fixture pool works now; restricted runtime pool/session-query/under-lock correctness not independently tested this round. Prior setup evidence remains separate. |
| R14 verification gates | Required final command plus constituent diagnostics | Table below | **FAIL** overall. |
| R2–R8/R11–R13 beyond T2 | Integrated T3/T4/T5 and Human | Outside diagnostic scope | **DEFERRED**, no C1 acceptance. |

## 2. Error Verification

| Error | Expected | Actual / disposition |
|---|---|---|
| Simulated invalid provider/session/malformed input | Reject identity/authority; preserve relevant error | Existing negatives pass at unit/protocol boundary; real HTTP safe-error envelope untested. |
| Initial integration cannot reach DB in sandbox | Explicit failure, no pass | Two DB tests fail at schema creation; same command succeeds on one authorized escalation retry. Underlying test suppresses driver detail; recovery isolates environment cause, not a code regression. |
| Initial scanner cannot access DNS | Record tooling failure | Socket permission error fetching vuln.go.dev retained; one authorized retry completes. |
| Final lint fails | Fail required gate | ST1005 at oidc.go:38 and :42 remains; no suppression/fix applied. |
| Gosec cookie findings | Record and triage, not automatic exploit or false-positive claim | Three G124 candidates at http.go:49, :158 and :277; helper sets HttpOnly/Lax and conditional Secure, and development exception is intended. Runtime validation and accepted disposition remain required. |
| Gitleaks findings | No clean clearance until disposition | Four historical backend candidates and one current directory candidate; ignored local env presence is not proof of committed exposure. |

## 3. Final Verification

All argv/cwd/time/exit and sanitized output hashes are in `commands.jsonl`. Explicit Go1.26.9, `GOTOOLCHAIN=local`, writable isolated caches, existing module cache, `GOPROXY=off`, `GOSUMDB=sum.golang.org`, `GOFLAGS=-mod=readonly`. No checksum bypass/module write. Source identities and old Invocation drift are in `source-hashes.json` and `input-recheck.json`.

| Command | Result |
|---|---|
| `go mod verify` | **PASS**, exit 0, all modules verified; cache integrity only. |
| backend `make verify` | **FAIL**, exit 2 at Staticcheck ST1005. Make did not execute later constituents. |
| independent `go test -count=1 -timeout=60s -v ./...` | **PASS**, six auth tests; DB/server no test files. |
| independent `go test -race -count=1 -timeout=90s ./...` | **PASS**, no detected process data races; no proof of DB contention invariants. |
| independent `go test -tags=contract -count=1 -timeout=60s -v ./...` | **PASS command**, same unit tests; no contract-specific test files found. **Contract behavior not established.** |
| `go test -tags=integration -count=1 -timeout=60s -v ./internal/platform/auth` | Initial sandbox exit 1; escalation retry **PASS**, two real DB tests, no skips, plus six unit/protocol tests. |
| independent `gosec ./...` | **FAIL**, exit 1; three G124 candidates. Tool reports version `dev`. |
| backend `make gitleaks` through cached Compose image v8.30.1 | **FAIL**, make exit 2. History: 36 backend-path commits, four candidates; worktree: one candidate. Source scope is backend/, despite repo mount needed for Git history. No allowances introduced. |
| independent `govulncheck -json ./...` | Initial sandbox exit 1 DNS/socket restriction; escalation retry exit 0 with full parsed stream: config 1, SBOM 1, progress 2, OSV metadata 200, **finding objects 0**, empty stderr. **PASS fresh source scan**, not whole dependency acceptance. Scanner v1.7.0 built Go1.26.9. |
| Read-only fixture cleanup query | Initial sandbox exit 2; escalation retry **PASS**: zero `c1_auth_%` schemas and zero other fixture-login connections. |

These isolated constituent diagnostics after failed make do not turn `make verify` into PASS. New Go/SQL production files were not added. Migration execution succeeded in each disposable schema and cleanup dropped only generated test-owned schemas. Historical DB and dev database were not mutated by this round. Existing dev migration tracking remains uninitialized (prior direct-psql setup); no `make migrate-up/down` was run.

Gitleaks history/current candidate identities from the preceding sanitized `../gitleaks-backend-20261011/summary.md` remain useful triage context; the current rerun independently confirms counts, not a new per-candidate disposition. The current directory candidate is local ignored `.env`. No secret or provider-token content is reproduced here.

Broad provider/browser/contention work was not attempted after final lint/security failures and failed formal Invocation preflight. The next formal independent Testing needs new Invocation/Participant/Session bindings for current source/tool/config inputs and a concrete isolated restricted-runtime test plan. No blocker closure, dependency commit, T3 dispatch or lifecycle projection was performed.

Recorder limitation: concurrent initialization initially caused one JSONDecodeError before make was launched; the capture wrapper was then invoked after context creation and recorded make normally. No test result was fabricated from that recorder failure. Logs are local ignored evidence under backend/tmp; promote required safe evidence through the owning new Run before removing this directory.

## 4. New Recurring Bug Patterns

No new reusable pattern promoted. Contract-tag exit 0 without contract-specific assertions and fixture-owner serial tests without runtime contention remain explicit evidence limitations.

## Verdict

**Fail — send back to Build for gate failures/triage.** Lint failure persists; Gosec candidates are newly executed findings this round; Gitleaks candidates remain unresolved. Real provider/browser/restricted-runtime concurrency gaps remain open. No previously established real-runtime PASS regressed; those checks were never accepted.

## Phase handoff

- **Outcome:** Session diagnostics COMPLETED; final verification FAIL; formal T2 independent acceptance BLOCKED by terminal/stale Invocation and remaining gaps.
- **Result refs:** This report, `patch-plan.md`, `commands.jsonl`, `vuln-summary.json`, `source-hashes.json`, `input-recheck.json`; context and sanitized streams in this directory.
- **Findings:** ST1005 persists; three untriaged G124 cookie candidates; Gitleaks four history/one local-env candidates. Details and actions in patch-plan.md.
- **Decision requests:** Orchestrator must reconcile this supplemental evidence and define fresh Run/Participant/Session/Invocation/model-approval posture. No acceptance decision inferred from user request to rerun.
- **Blockers:** Required make gate failed. Real Google registration/consent/test-user/interactive access and HTTPS remain unestablished for their tests. Backend-only read/fixture diagnostics are unaffected.
- **Open / unverified:** Real Google identity/code replay; cookie/browser HTTP behavior; distinct restricted-runtime concurrency/session lock/cancellation; contract assertions; migration version tracking; candidate dispositions.
- **Recommended continuation:** Fresh Build/Patch context grounded on patch-plan.md for lint and security triage, then independent Testing under a new current-input Invocation. Preserve TEST-T2-001 terminal evidence; do not reopen it or advance T3 from this report.
- **Context refs:** Current Approved Techplan §§9–12; T2 task; TEST-T2-001 Invocation/report; BUILD-T2-006 retry-1 report; CODEREVIEW-T2-002 report; local DB setup/verify evidence in backend/tmp.
