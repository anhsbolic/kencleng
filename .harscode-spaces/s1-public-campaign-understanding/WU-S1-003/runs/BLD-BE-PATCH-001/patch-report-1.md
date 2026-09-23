# BLD-BE-PATCH-001 — Patch Report 1

> Phase: Build/Patch  
> Work Unit / Run: `WU-S1-003` / `BLD-BE-PATCH-001`  
> Author: Codex CLI agent  
> Created: 2026-09-23  
> Model / Reasoning / Session: gpt-5.6-terra / high / Fresh Build/Patch session  
> Target revision: `baecb78099dd1ceb0d00b9c0215023a7d9e704fc`  
> Workflow revision: `d46358563942c7e015b97aa7c5767c880ef1bc63`

## What changed

- `backend/go.mod`, `backend/go.sum`, `internal/domain/campaign/repository_integration_test.go` → test-only `testcontainers-go` and `golang-migrate` harness; starts disposable Postgres without reading `DATABASE_URL`, applies all migrations, asserts Campaign PK/FK/funding constraints, steps `000011` down, and re-applies it.
- `internal/platform/storage/private_reader_integration_test.go` → disposable MinIO private-bucket test for JPEG read, missing-object/stat failure, and cancelled context. All failures assert Campaign's safe `ErrObjectUnavailable` class.
- `internal/transport/http/campaign_public_test.go` → absent/non-published/non-member × no/garbage/valid-shaped Authorization matrix checking status, body, `Content-Type`, and `Cache-Control`. Added 128-sample logical (fake-latency) representative parity proof: every eligible-path 404 consumes exactly the same controlled 2 ms dependency budget; no wall-clock threshold is used.
- `cmd/server/main.go`, `cmd/server/main_test.go` → `ReadHeaderTimeout` is explicitly 5 s, with a unit test, closing Campaign-scope G112.
- `cmd/campaign-seed/main.go` → three narrowly scoped, documented G304 acknowledgements only for explicit operator CLI paths: manifest decoding is capped/strict and media is subsequently decoded as JPEG/PNG. No global/static-analysis policy change or blanket suppression was added.

## Tests run

- `GOCACHE=/tmp/kencleng-bld-go go test -count=1 -timeout=60s -run '^TestPublicCampaign' ./internal/transport/http` → Campaign HTTP matrix/timing and existing focused wire checks: PASS.
- `GOCACHE=/tmp/kencleng-bld-go go test -count=1 -timeout=60s ./internal/domain/campaign` → Campaign unit regression: PASS.
- `GOMODCACHE=/tmp/kencleng-bld-mod GOCACHE=/tmp/kencleng-bld-go go test -count=1 -timeout=60s ./cmd/server ./cmd/campaign-seed` → server timeout regression and seed unit regression: PASS.
- `DOCKER_HOST=unix:///run/user/1000/podman/podman.sock TESTCONTAINERS_RYUK_DISABLED=true GOMODCACHE=/tmp/kencleng-bld-mod GOCACHE=/tmp/kencleng-bld-go go test -count=1 -tags=integration -run '^TestCampaignMigrations' ./internal/domain/campaign` → isolated Postgres apply/constraints/down/re-up: PASS.
- `DOCKER_HOST=unix:///run/user/1000/podman/podman.sock TESTCONTAINERS_RYUK_DISABLED=true GOMODCACHE=/tmp/kencleng-bld-mod GOCACHE=/tmp/kencleng-bld-go go test -count=1 -tags=integration -run '^TestPrivateReader' ./internal/platform/storage` → isolated MinIO read/missing/cancelled behavior: PASS.
- `GOMODCACHE=/tmp/kencleng-bld-mod GOCACHE=/tmp/kencleng-bld-go gosec ./cmd/campaign-seed ./cmd/server ./internal/domain/campaign ./internal/platform/storage` → Campaign scope has no G112/G304 report; remaining `cmd/server` G706 is pre-existing dev-outbox logging outside this Work Unit's Campaign change: FAIL only for that external finding.
- `GOMODCACHE=/tmp/kencleng-bld-mod GOCACHE=/tmp/kencleng-bld-go make verify` → `staticcheck` passed, then `gosec ./...` failed on 12 unrelated pre-existing Account/auth/OAuth/breachcheck/cookie/server findings. Campaign G112/G304 are absent. It therefore did not reach later Makefile targets.
- `git diff --check` → PASS.

## Verification scope confirmation

No race/concurrency or load/performance suite was run. The logical timing test is deterministic handler/service-seam evidence, not a wall-clock performance claim. Security-class `gosec` was run because this exact patch plan required Campaign-scope triage; `make verify` was run once to preserve the unrelated-blocker separation.

## Contract check

- [x] Current patch target satisfied in full.
- [x] Live-code re-grounding did not invalidate a material contract assumption.

## Deferred / not tested here

- Independent Testing must run the combined integration-tag command and independently assess this Build evidence before milestone promotion.
- Caddy prefix/policy/header preservation, browser integration, and MinIO anonymous policy remain WU-S1-005/Human owned.

## Flagged for Techplan / Testing

- `make verify` remains blocked only by 12 non-Campaign findings: Google OAuth URL/redirect/logging, cookie settings, breachcheck SHA-1 behavior, auth key file reads, and the pre-existing server dev-outbox log. This Run did not modify those areas.
- Testcontainers requires a configured Docker-compatible runtime. The successful local evidence used the explicit Podman socket above; the harness fails loudly if the runtime is unavailable and never uses `DATABASE_URL`.

## Phase handoff

- Completed: Testing re-entry patch for R1/R5/R6/R7 evidence and Campaign-scope static-analysis triage.
- Artifacts: this report and changed Campaign tests/harnesses.
- Human decision: none now; Tier-1 human review remains required for milestone promotion.
- Open / deferred: unrelated repository-wide `make verify` findings require their owning work units.
- Recommended next step: return to fresh independent Testing, not Code Review.
- Session transition: start fresh Testing for independent execution of the integration/timing/security evidence.
- Context pointers: `TP-BE-001/techplan.md`, `TST-BE-001/testing-report-1.md`, `TST-BE-001/patch-plan-1.md`, and the changed files above.
