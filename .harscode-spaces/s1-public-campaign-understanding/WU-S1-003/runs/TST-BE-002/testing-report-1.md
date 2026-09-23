# TST-BE-002 — Independent Testing Re-entry Report

> Phase: Testing  
> Work Unit / Run: `WU-S1-003` / `TST-BE-002`  
> Author: Codex CLI agent  
> Role / Specialization: Verifier / Independent Testing re-entry — backend Campaign verification closure  
> Session: Fresh independent Testing session  
> Created: 2026-09-23  
> Model / Reasoning: `gpt-5.6-terra` / `high`  
> Target revision checked: `6b7a1034e80b0c74a6bd6bc7e31a0c10424a9010`  
> Workflow revision: `d46358563942c7e015b97aa7c5767c880ef1bc63`

## 0. Sweep Summary

- Confirmed: the prior Patch's affected integration gap through one combined independent command:
  `DOCKER_HOST=unix:///run/user/1000/podman/podman.sock TESTCONTAINERS_RYUK_DISABLED=true GOMODCACHE=/tmp/kencleng-tst2-mod GOCACHE=/tmp/kencleng-tst2-go go test -v -count=1 -tags=integration -timeout=180s ./internal/domain/campaign ./internal/platform/storage` — **PASS**. Disposable `postgres:16-alpine` and private-bucket MinIO containers ran through Podman API `1.41`; the test applied all migrations, checked Campaign constraints, stepped `000011` down and re-applied it, then read a private JPEG and classified missing/cancelled object reads safely. The integration fixtures do not reference `DATABASE_URL`.
- Closed from prior gap/deferred list: R5 absent/non-published/non-member × no/garbage/valid-shaped Authorization matrix and representative timing-parity evidence through `go test -v -count=1 -timeout=60s -run '^TestPublicCampaignHandlers_(NotFoundAuthorizationMatrix|RepresentativeNotFoundTimingParity)$' ./internal/transport/http` — **PASS**. All 12 response-parity cases passed. The timing test executes 128 samples per representative class against a fixed 2 ms logical dependency budget; it proves deterministic handler/service-seam parity, not a wall-clock database latency guarantee.
- Closed from prior finding: Campaign-scope static-analysis findings are cleared. `gosec` over Campaign, seed/server composition, storage, and HTTP reported no finding in Campaign files, no `G304` in `cmd/campaign-seed`, and no `G112` in `cmd/server`.
- Still requires fresh Testing: no Campaign-owned gap. Proxy/policy/header/browser topology evidence remains deliberately outside this Work Unit under WU-S1-005.

## 0a. Test Focus Pointer Execution

| Area | Evidence anchor opened | Specialized verification | Result |
|---|---|---|---|
| Public projection / anti-enumeration | `EXP-BE-001/evidence/gap-analysis.md` §Area 3 | Direct real-HTTP handler matrix for four 404 resource classes and three Authorization forms, comparing status, body, `Content-Type`, and `Cache-Control`; 128-sample controlled-latency parity test. | PASS — all 12 parity cases match; each representative path consumes exactly 256 ms logical budget. |
| Decimal funding truth | `EXP-BE-001/evidence/solutioning.md` §Decision 2 | Current Campaign domain tests in unit, integration combined run, and targeted race run. | PASS — exact decimal unit cases and closed projection execute; no Campaign `float64` funding path was found during current source/test inspection. |
| Controlled media / retraction / cache | `EXP-BE-001/evidence/gap-analysis.md` §Area 4 | Disposable MinIO private-bucket adapter: read JPEG, missing-object stat failure, and cancelled request classification; direct handler tests/race verify no-store and controlled stream behavior. | PASS for backend-owned direct boundary. WU-S1-005 remains owner of effective policy, Caddy, and proxy preservation. |
| Migration / integration isolation | `EXP-BE-001/evidence/gap-analysis.md` §Area 5 | Disposable Postgres migration fresh apply, constraints, `000011` down, and re-up via configured Podman Docker-compatible socket. | PASS — isolated harness, never shared `DATABASE_URL`. |
| Runtime concurrency | `EXP-BE-001/evidence/solutioning.md` §Carry-forward risks and non-goals | Race detector over all changed/affected Campaign packages. | PASS — no new shared mutable balance/state-transition concern. |

## 1. Test Coverage

| Rule / scenario | Category | Observable verification | Result |
|---|---|---|---|
| R1 schema, constraints, down/re-up | Integration | Combined integration-tag Campaign/Postgres+MinIO command above; `TestCampaignMigrations_ApplyConstraintsDownAndReUp`. | PASS |
| R1 seed validation / no implicit setup | Negative/regression | `go test -count=1 -timeout=120s ./...`. | PASS |
| R2 closed public projection | Happy/negative | Full unit suite and targeted race suite include `TestPublicCampaignDetailHandler_ExactClosedWireShape`. | PASS |
| R3 exact IDR/funding semantics | Boundary | Full unit, integration, and targeted race runs include Campaign service decimal tests. | PASS |
| R4 controlled media metadata/content | Happy/negative | Handler stream/reader-close coverage plus isolated MinIO JPEG read. | PASS |
| R5 indistinguishable 404, no-store, ignored Authorization | Security | 12-case direct HTTP matrix and deterministic representative timing test. | PASS |
| R5 eligible unavailable behavior | Error/security | Existing handler and Campaign service tests, plus actual missing-object MinIO adapter test. | PASS |
| R6 private adapter/cancellation/pre-header availability | Integration/error | `TestPrivateReader_ReadsPrivateMediaAndClassifiesUnavailable` in disposable private bucket. | PASS |
| R7 regression, migration/operator safety | Regression | Full unit suite, contract-tag suite, integration suite, targeted race, `git diff --check`, and static analysis described below. | PASS with repository-wide static-analysis follow-up outside Campaign |
| R8 topology ownership | Human | Techplan/fence re-read. | N/A — WU-S1-005/Human-owned; no integrated claim made. |

## 2. Error Verification

| Error case | Expected behavior/category | Actual | Actionable/propagated correctly? |
|---|---|---|---|
| Malformed, absent, non-published Campaign; absent/non-member media under all Authorization variants | Identical public `404` Problem Details with `private, no-store`; Authorization ignored | 12 matrix cases pass with matching status, complete body, content type, and cache header. | Yes |
| Eligible missing object / object stat failure | Generic public unavailable (`503` at Campaign boundary), no storage detail | Isolated MinIO adapter maps it to `campaign.ErrObjectUnavailable`; service/handler mappings pass. | Yes |
| Cancelled object read | Safe unavailable class; request context retained | Isolated MinIO test passes a cancelled context and receives `campaign.ErrObjectUnavailable`. | Yes |
| Migration FK, non-negative amount, and nullable funding-pair violations | Database rejects invalid persisted state | Disposable Postgres integration test passes. | Yes |

## 3. Final Verification

- Target repo required final build/lint/test commands: `GOMODCACHE=/tmp/kencleng-tst2-mod GOCACHE=/tmp/kencleng-tst2-go make verify` — **not clean**. `staticcheck ./...` passed; `gosec ./...` stopped with 12 existing non-Campaign reports: Google OAuth URL/redirect/logging, Account cookie/security-key reads, breachcheck SHA-1 behavior, and the pre-existing dev-outbox log in `cmd/server`. It did not reach subsequent Makefile targets.
- Campaign security re-check: `gosec ./cmd/campaign-seed ./cmd/server ./internal/domain/campaign ./internal/platform/storage ./internal/transport/http` found seven reports, all in pre-existing OAuth/cookie/dev-outbox files; none is in Campaign code, seed command, or storage. The earlier scope-specific `G304` seed reports and `G112` server report are absent.
- Full unit regression: `go test -count=1 -timeout=120s ./...` — **PASS**.
- Contract-tag regression (run independently because `make verify` stopped at unrelated `gosec`): `go test -count=1 -tags=contract -timeout=120s ./...` — **PASS**.
- Race: targeted affected scope `go test -v -count=1 -race -timeout=90s ./cmd/campaign-seed ./cmd/server ./internal/domain/campaign ./internal/platform/storage ./internal/transport/http` — **PASS**. Two attempted full `go test -race ./...` executions were terminated by this runner around the unrelated Account test sequence before a package-complete result; this is recorded as an environment/tooling limitation, not a passing broad-race claim.
- Broad checks intentionally not rerun: none for the affected Campaign gap. The only incomplete broad check is full-repository race as described above; targeted Campaign race passed.
- Migration/schema collision: **PASS** via isolated fresh apply, constraints, down, re-up. `git diff --check` also passed. Backend worktree had no local changes; unrelated frontend/WU-S1-004 worktree changes were preserved.
- Backward compatibility: **PASS for executable backend evidence** — additive migration passes down/re-up and full unit/contract-tag suites pass. No legacy Campaign API exists to exercise beyond the settled routes.
- Broader-suite requirement for cross-cutting change: unit and contract-tag suites pass; official `make verify` is blocked exclusively by the stated pre-existing reports outside WU-S1-003.
- Fresh Techplan consistency read: no contradiction or new Techplan drift. The implemented evidence now closes its Testing-owned R1/R5/R6/R7 checks. Timing evidence remains honestly scoped to deterministic seam parity; topology remains out of scope.

## 4. New Recurring Bug Patterns

None. The repository-wide `gosec` baseline and runner's full-race execution limit require owner/tooling follow-up, but neither is a Campaign defect pattern introduced by this Work Unit.

## Verdict

**Pass with flagged follow-ups.** Campaign backend correctness and the Testing-owned closure evidence are sufficient for `BACKEND_VERIFIED`; no production patch plan is needed. The non-blocking flags are the pre-existing repository-wide `gosec` reports that prevent a green aggregate `make verify`, and the runner-limited full-repository race command. Neither supports a Campaign failure or an integrated/topology claim.

## Phase handoff

- Completed: independent Postgres+MinIO integration closure; R5 response/Auth/timing parity; Campaign security re-check; full unit and contract-tag regression; targeted affected-scope race; final-tooling assessment.
- Artifacts: this testing report; no patch plan.
- Human decision: Tier-1 human review remains required before milestone/merge; WU-S1-005/Human must provide topology/policy/proxy/browser evidence before any integrated statement.
- Open / deferred: repository owners must triage the 12 pre-existing `gosec` reports; runner/tooling owner may address full-suite race execution limit. Neither blocks Campaign `BACKEND_VERIFIED`.
- Recommended next step: PR/review when required human gates are ready; retain this report as backend evidence and coordinate WU-S1-005 only for topology/integration proof.
- Session transition: a fresh PR session is useful for an independent final-diff/evidence read; no Build/Patch return is warranted.
- Context pointers: `TP-BE-001/techplan.md`, `TST-BE-001/testing-report-1.md`, `BLD-BE-PATCH-001/patch-report-1.md`, and this report.
