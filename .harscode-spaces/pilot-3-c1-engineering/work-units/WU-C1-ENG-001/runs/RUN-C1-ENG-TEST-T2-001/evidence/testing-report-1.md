# Independent Testing — RUN-C1-ENG-TEST-T2-001

- Phase / Role / Author: Independent Testing / Tester / Codex.
- Created: 2026-10-11 Asia/Jakarta.
- Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-TEST-T2-001`.
- Participant / Session: `PARTICIPANT-C1-ENG-T2-TESTER-001` / `SESSION-C1-ENG-T2-TESTER-001`, Invocation bindings for this fresh Testing context; runtime thread ID is not established.
- Model / Reasoning: Invocation selects and records approval for `gpt-6-sol / medium`; actual runtime identity/effort is not independently exposed. No model approval is inferred from these bindings.
- Target HEAD / branch: `059d145d5df64fcda427e5a625486a8e95ce1fd8` / `pilot/3-c1-engineering`.
- Workflow HEAD: `91a97af8ba64f98b0a692a83fbff866a210c45c0`.
- Dispatch authority: current Human request to execute the exact Invocation.
- Outcome: **STALLED — real-runtime prerequisites blocked; required final gate failed. No T2 acceptance.**

## 0. Sweep Summary

Authority preflight matched all 33 input identities checked against Invocation (`bound-inputs.json`), including current Approved Techplan, T2 task, manifest, exact retry-1 Build report, independent Review T2-002 report, six bound code files, root/backend instructions, lifecycle authorities, runtime command authority, Test Focus source and targeted workflow/guidance. Target and workflow HEAD/branch match. Existing working-tree changes were retained. WU, Events, Control Tower and Invocation were not changed.

Fresh end-to-end Approved Techplan read completed, including §§9–10/12 and §13. Exact latest retry-1 Build and independent Review reports were read. Build claims were not treated as executed Testing evidence. Source hashes for the current backend Go/SQL/module inputs are in `tested-source-hashes.json`; final boundary recheck is in `final-boundary.json`.

Confirmed by fresh execution:

- `go mod verify`: exit 0, `all modules verified`; cache integrity only, not provenance, vulnerability clearance or runtime acceptance.
- `go test -count=1 -timeout=60s -v ./...`: exit 0, six auth top-level tests pass without skips. This spot-check confirms the existing protocol-negative and unit regression claims. Provider transport/signing fixtures and fake session-row scans remain **simulations**. Server and DB packages report no test files; their output establishes compilation only.
- Podman executable exists. Sandbox `podman --version` failed on `/run/user/1000/libpod` read-only filesystem; harness-approved diagnostic retry returned Podman 5.7.0. This verifies version availability only, not PostgreSQL/container/SCRAM operation.

Closed from prior deferred list: none of the real provider/browser/DB/concurrency gaps. Testing-owned final verification remains incomplete. Config presence checks found no inherited APP_ENV/APP_ORIGIN/Google client/secret/DATABASE_URL/C1_AUTH_TEST_DATABASE_URL and neither root nor backend `.env`. No credential values, tokens, codes or payloads were printed or persisted. No application DB connection was attempted.

## 0a. Test Focus Pointer Execution

| Area | Evidence anchor opened | Specialized verification | Result |
|---|---|---|---|
| Initial person/auth attribution | `RUN-C1-ENG-EXPLORATION-001/evidence/stage-3-solutioning.md#material-questions-and-constraints`, items 1 and 3; `### Attribution and authentication timing` | Require real configured Google identity, stable issuer/subject mapping and distinct restricted DB runtime; spot-check protocol and post-lock expiry oracles | Rationale recovered: durable person attribution at success, implementation authorization separate from product approval. Simulated unit/protocol checks pass; real provider and DB tests NOT RUN/BLOCKED. |
| Session lock/revocation/expiry | Current Techplan §§9–10/12 and exact retry-1 report | Existing post-Scan expiry/revocation negatives; planned real blocked-row expiry/logout/rotation contention via current T2 Store/LockSession boundary | Unit simulations pass. No real contention execution; absent disposable authorized DSN/runtime role. |

Organization + Owner consistency and known conflict anchors were not opened: no observed T2 test crossed those T3 boundaries. No new Test Focus drift was established. Matching Go concurrency, integration, token lifecycle, authorization, dependency/supply-chain and PostgreSQL locking guidance was read through current routing; REST cookie/CSRF guidance also informed the boundary assessment. Local JWT signing/refresh rules do not apply to opaque local sessions.

## 1. Test Coverage

| Rule / scenario | Category | Observable verification | Result |
|---|---|---|---|
| R1/R9 real Google login across sessions and stable durable person | Happy / repeat / real interface | Config presence preflight; actual callback/login would need registered client/callback, consent/test-user and authorized isolated DB | **NOT RUN/BLOCKED**. Missing configuration and external prerequisite evidence. No real identity claim. |
| R9 forged issuer/audience/azp/signature/algorithm, nonce/expiry/key/provider failure; PKCE/input | Negative / edge | Existing `TestGoogleProtocolNegatives` and `TestGoogleProtocolPKCEAndMissingExchangeInputs` via production constructor/verifier with simulated discovery/JWKS/token transport | **PASS — SIMULATED provider**, independently rerun. Valid signed control plus 14 negatives pass. Not real Google callback/code-replay evidence. |
| R9 state/browser binding and local pending-login replay/expiry | Negative / edge / DB | Existing integration tests require explicit disposable URL | **NOT RUN/BLOCKED**. No DSN; suite not invoked to avoid misreporting its skip as acceptance. Serial Build fixture evidence remains Build-owned only. |
| R1/R9 concurrent first-login unique upsert | Concurrent / real DB | Real restricted PostgreSQL runtime required | **NOT RUN/BLOCKED**. No authorized disposable connection/identity. No replacement trust/superuser fixture used. |
| R10 session expiry after row acquisition, expiry equality, revocation, missing/storage error, stale caller time | Edge / regression | `TestLockSessionExpiredDuringLockWait`, `TestLockSessionRevalidation`, `TestLockSessionIgnoresStaleCallerTime` | **PASS — SIMULATED row/lock wait**, independently rerun; zero authority on failure. No real DB lock-order proof. |
| R9/R10 malformed browser/session input | Negative | `TestMalformedBrowserAuthorityRejectedBeforeStorage` | **PASS — unit boundary**; malformed input rejected without reaching storage. |
| R10 cookie flags/dev-only exception, production HTTP rejection, rotation/logout/CSRF/Origin/no-store/fixed redirect at actual HTTP/browser boundary | Happy / negative / edge | Current handler/server source inspected to identify surface; configured HTTP/browser execution needs provider/DB and HTTPS for production cookie evidence | **NOT RUN/BLOCKED** for real runtime/browser. No server/browser launched; no cookie or HTTP-response pass claimed from source inspection. |
| R10/R11/R14 logout/revocation/rotation contention and expiry while waiting for actual session row lock | Concurrent / real DB | Current Store/LockSession boundary; T3 confirmation handler separately deferred | **NOT RUN/BLOCKED** for T2 contention; **DEFERRED** for integrated T3 confirmation. |
| R14 configured pool/SCRAM/query mode/restricted credentials and session queries | Compatibility / real DB | Config preflight and exact current source identity | **NOT RUN/BLOCKED**. Credential privileges, SCRAM negotiation, query mode, connection/pool behavior and under-lock decisions unverified. Podman version is insufficient. |
| R14 integrity / final verification | Final gate | `go mod verify`; backend `make verify` and diagnostic recovery | Integrity **PASS**; final gate **FAIL**, detail in §3. |
| R2–R8/R11–R13 organization, guard, scoped inspection, replay/unknown outcomes, UI/private cache and operator behavior | Integrated / Human-owned | Outside current bounded T2 Invocation | **DEFERRED** to T3/T4/T5 and later integrated Testing/Human owners. R1 overall establishment success also unverified. No whole-C1 verdict. |

All rule IDs are accounted for; a pass on T2 protocol/unit coverage does not establish the broader product rule. No new test harness changed the production tree or invented identity/authorization semantics.

## 2. Error Verification

| Error case | Expected behavior/category | Actual | Actionable/propagated correctly? |
|---|---|---|---|
| Simulated invalid provider claims, key/token/provider failure | Reject identity; no trusted principal | Existing protocol negatives pass | Confirmed at constructor/verifier boundary only; real HTTP/callback behavior unverified. |
| Simulated missing/revoked/expired session | `ErrUnauthenticated`, zero Session | Unit matrix and stale-time regression pass | Sentinel and zero result confirmed; external safe 401 envelope not exercised. |
| Simulated query failure | Preserve error identity, zero Session | `query_error` subtest passes | Error-chain behavior confirmed; actual DB failure propagation not exercised. |
| Malformed browser/state/session input | Reject before storage | Existing unit test passes | Confirmed at unit boundary; no real browser claim. |
| Missing runtime prerequisites | Explicit NOT RUN/BLOCKED | Provider/DB/browser tests not invoked | Reported to Orchestrator/Human environment owner; no fake session/login evidence. |
| Lint/tooling failures | Record exact failed command; no false final pass | Initial cache error and recovery ST1005 errors retained | Environment and substantive lint failure separated. |

## 3. Final Verification

Exact argv/cwd/time/exit and output SHA-256 are in `commands.jsonl`; full captures remain alongside this report. Go1.26.9 linux/amd64 was selected explicitly, `GOTOOLCHAIN=local`, `GOFLAGS=-mod=readonly`, writable unique `/tmp` Go cache, existing exact module cache, `GOPROXY=off`, checksum verification `GOSUMDB=sum.golang.org` retained. These commands did not write go.mod/go.sum. Scanner metadata is v1.7.0 built with Go1.26.9; metadata inspection is not a fresh scan.

| Command / step | Actual result |
|---|---|
| `go mod verify` | **PASS**, exit 0; `mod-verify.stdout` / `.stderr`. |
| backend `make verify`, initial sandbox attempt | **FAILED tooling**, exit 2 at Staticcheck cache write; `make-verify.stdout` / `.stderr`. GOCACHE was writable, but initial Staticcheck cache remained in read-only home. |
| backend `make verify`, one harness-approved diagnostic recovery | **FAIL**, exit 2 at Staticcheck ST1005 in `oidc.go:38` and `:42`; `make-verify-recovery.stdout` / `.stderr`. Recovery used `XDG_CACHE_HOME=<capture cache>/tool-cache`, `STATICCHECK_CACHE=<capture cache>/staticcheck`; those overrides were not included by the initial recorder schema, so they are explicitly preserved here. No gate continuation/acceptance inferred from recovery. |
| `gosec ./...`, Makefile unit/race/contract/security steps | **NOT RUN by make**: first lint constituent failed. |
| `gitleaks detect` | **NOT RUN/BLOCKED**: executable absent from configured PATH; make did not reach security. No secret-scanning clearance. |
| `govulncheck ./...` | **NOT RUN in this Testing Run**: make did not reach security; no fresh vulnerability clearance. Prior Build source scans remain prior evidence, not a Testing execution. |
| independent `go test -count=1 -timeout=60s -v ./...` spot-check | **PASS**, exit 0; six auth tests, no skip, `unit.stdout` / `.stderr`. This isolated spot-check after failed gate does not reopen/clear final verification. |
| `go test -race ./...`, `go test -tags=contract ./...` | **NOT RUN** after final lint failure; these required final subgates remain open. |
| Real PostgreSQL integration | **NOT RUN/BLOCKED** before DB connection/mutation. |

F-TEST-T2-001: required lint fails on two capitalized constructor validation error messages. Source is current T2 `oidc.go`, outside the prior retry's reviewed six-file diff. This is a newly observed final-gate gap; no functional/security regression by that retry is established. See `patch-plan-1.md`; Testing made no production fix or suppression.

Migration/schema collision: **NOT TESTED**, no disposable DB/migration execution. Backward compatibility: current backend packages compile in unit execution, but no runtime health/auth/browser/deployed DB compatibility claim. Broader final suite: **INCOMPLETE** as above; neither prior Build captures nor unit simulations replace it.

Fresh Techplan consistency read: no new material product/authority contradiction established. Historical baseline/planning prose is not current execution proof. Full-C1 rows remain assigned to later integrated owners; configured environment obligations remain open. No protected authority repair was made.

Cleanup/resource boundary: no database pool/session, schema, migration, application server, browser context or container was created by this Run; no production/unknown DSN was accessed. Every launched capture subprocess returned without timeout. `final-boundary.json` records the terminal read-only source/authority recheck and matching process-name inspection; no process was killed or container removed. The disposable Go/tool cache is retained at the path recorded in `capture-context.json`, for diagnostic reproducibility; it contains no application credentials or DB state.

## 4. New Recurring Bug Patterns

None proposed. The two lint errors are a bounded current-source finding; environment/cache/tool prerequisites are reported as execution limitations. No Harscode guidance/examples were edited.

## Verdict

**Fail — send the bounded lint finding back to Build; real runtime testing remains NOT RUN/BLOCKED.**

Blocking for T2 acceptance: F-TEST-T2-001 (failed required final gate), missing real Google/browser configuration and authorized disposable restricted-credential DB, and unavailable secret-scanner tooling. Successful unit simulations and integrity do not establish real provider/DB/security acceptance. No vulnerability, credential privilege or production auth failure is claimed without execution. No blocker closure, dependency commit, T3 dispatch or C1 completion is authorized; B-T2-002/B-T2-003 remain untouched.

## Phase handoff

- **Outcome:** STALLED — bounded spot-check evidence recorded; real-runtime prerequisites blocked and final verification failed.
- **Result refs:** `evidence/testing-report-1.md`; `patch-plan-1.md`; `preflight.json`, `bound-inputs.json`, `tested-source-hashes.json`, `commands.jsonl`, command captures, `final-boundary.json` and `evidence-hashes.json`.
- **Findings:** F-TEST-T2-001, §3; bounded lint patch owned by Build, not a demonstrated auth/security regression.
- **Decision requests:** Orchestrator reconcile failed gate and prerequisite gaps; Human/environment owner provide safe configured access through approved channels if real runtime Testing is to continue. No secrets requested in chat and no new product/security meaning selected.
- **Blockers:** Real provider/browser and restricted disposable DB evidence unavailable; final verification blocked by lint and missing gitleaks. Unit/integrity evidence is unaffected but insufficient for T2 acceptance.
- **Open / unverified:** Real callback/code replay, cookies/HTTPS/CSRF/logout/rotation at HTTP/browser boundary, SCRAM/privileges/query mode/concurrent first login/session lock contention, race/contract/security gates; integrated T3/T4/T5/Human evidence deferred.
- **Recommended continuation:** Reconcile with Orchestrator; fresh Build/Patch for `patch-plan-1.md`, provision required tools and authorized disposable/provider prerequisites, then route fresh independent Testing with explicit preflight/recovery terms. Preserve this terminal report. No automatic commit/downstream route.
- **Context refs:** This Run's `invocation.md`; current Approved `techplan/techplan.md` and T2 task; exact retry-1 Build and Review T2-002 reports bound by Invocation.
