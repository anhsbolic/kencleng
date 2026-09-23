# TST-BE-001 — Independent Testing Report

> Phase: Testing
> Work Unit: `WU-S1-003`
> Run: `TST-BE-001`
> Author: Codex CLI agent
> Role / Specialization: Verifier / Independent Testing — Backend Campaign public delivery
> Session: Fresh independent Testing session
> Created: 2026-09-23
> Model / Reasoning: `gpt-5.6-terra` / `high`
> Target revision checked: `baecb78099dd1ceb0d00b9c0215023a7d9e704fc`
> Workflow revision: `d46358563942c7e015b97aa7c5767c880ef1bc63`

## 0. Sweep Summary

- Confirmed: R2/R4/R5 handler and closed-wire coverage via `go test -count=1 -timeout=30s -run '^TestPublicCampaign' ./internal/transport/http` — PASS. Existing tests exercise exact top-level shape, `Cache-Control`, a malformed-ID 404 parity case, a supplied Authorization header on success, safe 503, JPEG streaming, and reader close.
- Confirmed: R1 seed-manifest unit coverage via `go test -count=1 -timeout=30s ./cmd/campaign-seed` — PASS.
- Confirmed: Campaign-domain decimal/public semantic coverage via `go test -count=1 -timeout=60s ./internal/domain/campaign` — PASS.
- Closed from Build deferred list: broad unit regression ran outside the sandbox as `GOCACHE=/tmp/kencleng-tst-go go test -count=1 -timeout=90s ./...` — PASS. The sandbox attempt itself was unusable because unrelated `httptest.NewServer` calls could not bind IPv6 loopback; the same suite passed outside the sandbox.
- Closed from Build deferred list: targeted race evidence ran outside the sandbox as `GOCACHE=/tmp/kencleng-tst-go go test -count=1 -race -timeout=90s ./internal/domain/campaign ./internal/platform/storage ./internal/transport/http ./cmd/campaign-seed ./cmd/server` — PASS. `storage` and `cmd/server` report `[no test files]`.
- Still requires fresh Testing: isolated Postgres migration apply/down/re-up and constraint/repository behavior; isolated MinIO adapter/cancellation/missing-object behavior; representative anti-enumeration timing parity. These cannot be claimed because no Campaign integration test or `testcontainers-go` dependency exists.

## 0a. Test Focus Pointer Execution

| Area | Evidence anchor opened | Specialized verification | Result |
|---|---|---|---|
| Public projection / anti-enumeration | `EXP-BE-001/evidence/gap-analysis.md` §Area 3 | Existing direct-handler tests and source inspection for malformed ID, shared response and ignored Authorization; looked for representative repeated timing cases. | Partial: wire checks pass, but no repeated timing-parity test exists; only one success test carries one Authorization value. |
| Decimal funding truth | `EXP-BE-001/evidence/solutioning.md` §Decision 2 | Independent Campaign-domain unit execution. | PASS for existing table tests; no `float64` funding path found in Campaign package. |
| Controlled media / retraction / cache | `EXP-BE-001/evidence/gap-analysis.md` §Area 4 | Existing direct-handler tests, race test, and `PrivateReader` test inventory. | Partial: fake-backed HTTP stream behavior passes; `PrivateReader` has no tests, so cancellation and real MinIO unavailable classification are unverified. WU-S1-005 still owns policy/proxy proof. |
| Migration / integration isolation | `EXP-BE-001/evidence/gap-analysis.md` §Area 5 | Probed Podman and integration-tag inventory/dependency. Podman Engine `5.7.0` and socket `/run/user/1000/podman/podman.sock` are available outside the sandbox; `go list -m github.com/testcontainers/testcontainers-go` reports it is not a known dependency. | FAIL: no Campaign integration-tag tests or testcontainers harness exist. |
| Runtime concurrency | `EXP-BE-001/evidence/solutioning.md` §Carry-forward risks and non-goals | Targeted `-race` over Campaign-related packages. | PASS; no new balance/state-transition concurrency concern identified. |

## 1. Test Coverage

| Rule / scenario | Category | Observable verification | Result |
|---|---|---|---|
| R1 seed validation | Negative | `./cmd/campaign-seed` unit tests. | PASS |
| R1 persisted schema / reversible migration | Integration | Required isolated Postgres apply, constraints, down, and re-up. | NOT VERIFIED — missing testcontainers harness/tests. |
| R2 exact public projection | Happy/negative | Existing `TestPublicCampaignDetailHandler_ExactClosedWireShape`. | PASS |
| R3 decimal truth | Boundary | Existing Campaign-domain tests executed. | PASS for authored cases |
| R4 metadata and controlled content wire | Happy/negative | Existing `TestPublicCampaignMediaHandler_StreamsValidatedBytesAndUnavailable`. | PASS with fake object reader only |
| R5 indistinguishable 404/no-store/auth | Security | Existing malformed-ID parity test and one Authorization success request. | PARTIAL — absent/non-published/non-member plus no/garbage/valid Authorization matrix and timing samples are not independently covered. |
| R6 MinIO cancellation/pre-header unavailable | Integration/error | `platform/storage` inventory and integration-tag test command. | NOT VERIFIED — no test files/harness. |
| R7 broad regression / operator safety | Regression | Full `go test ./...` passed; required `make verify` failed in lint. | FAIL — see Final Verification. |
| R8 topology ownership | Human | File fence and Techplan read. | N/A — remains WU-S1-005/Human-owned. |

## 2. Error Verification

| Error case | Expected behavior/category | Actual | Actionable/propagated correctly? |
|---|---|---|---|
| Malformed Campaign/media ID | Uniform `404`, RFC 9457 body, `private, no-store` | Existing handler test passed. | Yes, for malformed path. |
| Fake storage failure | Generic `503`, no key/provider text, no-store | Existing handler test passed. | Yes, at fake boundary. |
| Real missing/cancelled MinIO object | `503` before headers, request cancellation reaches client | No executable adapter test. | Not verified. |
| DB constraints and repository scan failure | Constraint rejection / safe propagated service error | No isolated real Postgres test. | Not verified. |

## 3. Final Verification

- Target repo required final build/lint/test commands: `make verify` — FAIL at `gosec ./...` (16 findings). It therefore did not reach `test-unit`, `test-race`, `test-contract`, `gitleaks`, or `govulncheck`. Findings include new-scope `G304` reports in `cmd/campaign-seed/main.go` and `G112` for `cmd/server/main.go`; other reports are pre-existing Account/auth/OAuth/breachcheck paths and require ownership triage rather than silent suppression.
- Broad regression: `GOCACHE=/tmp/kencleng-tst-go go test -count=1 -timeout=90s ./...` — PASS outside sandbox. The sandbox-only loopback restriction is recorded, not a product failure.
- Broad race: attempted `go test -count=1 -race ./...` did not return a complete package result through this runner (only the first two package results before interruption). It is not claimed as evidence. The relevant targeted race suite passed.
- Migration/schema collision: FAIL / not executable — migration files exist, but required disposable Postgres apply/down/re-up and constraint checks have no Campaign testcontainers implementation.
- Backward compatibility: additive migration files `000011_create_public_campaigns.{up,down}.sql` exist and `git diff --check` passed; runtime migration compatibility is not verified.
- Broader-suite requirement for cross-cutting change: unit regression passed; official `make verify` failed as above.
- Fresh Techplan consistency read: no authority contradiction found. The planned D5/testcontainers evidence and R5 timing evidence were omitted from Build implementation, which is a delivery gap rather than Techplan drift.

## 4. New Recurring Bug Patterns

None. The missing integration harness is a Work Unit-specific delivery omission; it should not be promoted as a generic pattern yet.

## Verdict

**Fail — send back to Build.**

Blocking failures: missing Testing-owned isolated Postgres/MinIO evidence and a failing target-repo `make verify`. Non-blocking but required follow-up: complete the R5 authorization matrix and timing-parity evidence. These are new coverage gaps relative to the approved Techplan, not regressions of a previously passing Campaign test.

## Phase handoff

- Completed: independent targeted/domain/wire/race verification, full unit regression, Podman compatibility probe, integration evidence inventory, and target-repo verification.
- Artifacts: this report; `patch-plan-1.md`.
- Human decision: Human review remains required for Tier-1 milestone and for any shared/manual migration application; no human decision can replace missing executable integration evidence.
- Open / deferred: the patch plan items; WU-S1-005 topology/policy/proxy/browser checks remain out of scope.
- Recommended next step: Build/Patch using `patch-plan-1.md`, then a fresh Testing Run to execute the isolated containers, timing evidence, targeted checks, and `make verify`.
- Session transition: start a fresh Build/Patch session. The prior Build session is no longer context-optimal because the patch spans test infrastructure, security-tool findings, and a required final verification failure.
- Context pointers: `TP-BE-001/techplan.md`, this report, `patch-plan-1.md`, and `CR-BE-001/review-findings.md`.
