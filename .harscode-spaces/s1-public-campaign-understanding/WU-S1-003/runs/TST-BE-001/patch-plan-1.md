# TST-BE-001 — Patch Plan 1

> Phase: Testing
> Work Unit: `WU-S1-003`
> Run: `TST-BE-001`
> Author: Codex CLI agent
> Created: 2026-09-23
> Target revision assessed: `baecb78099dd1ceb0d00b9c0215023a7d9e704fc`

## Finding 1 — Required isolated integration evidence is absent

**Evidence:** `backend/internal/domain/campaign/` has no `*_integration_test.go`; `backend/internal/platform/storage/` has no test file; `go list -m github.com/testcontainers/testcontainers-go` reports no dependency. The approved Techplan assigns R1/R6/R7 disposable Postgres/MinIO proof to Testing.

**Build/Patch work:**

1. Add the minimal test-only `testcontainers-go` dependencies and a Campaign-local isolated fixture. It must start disposable Postgres and MinIO through the Docker-compatible Podman socket, apply all migrations, and never consult or mutate `DATABASE_URL`.
2. Add a migration test that proves fresh apply, PK/FK/check constraints (including non-negative amounts and nullable funding pair), down child-before-parent, and re-up.
3. Add `PrivateReader` integration tests for a private seeded JPEG/PNG, missing object / stat failure classification, and cancelled request context. Assert errors map safely at the Campaign boundary.
4. Ensure cleanup is reliable and test failure explicitly reports unavailable container runtime rather than silently skipping milestone evidence.

**Acceptance evidence:** `go test -count=1 -tags=integration ./internal/domain/campaign ./internal/platform/storage` against isolated containers; include commands/results in the next Build report.

## Finding 2 — R5 anti-enumeration evidence is incomplete

**Evidence:** `campaign_public_test.go` covers malformed IDs and a single Authorization-bearing success request, but lacks the approved absent/non-published/non-member × no/garbage/valid Authorization matrix and no repeated timing-parity measurement.

**Build/Patch work:**

1. Extend direct handler/real-HTTP tests to compare full status/body/`Content-Type`/`Cache-Control` for all stated 404 classes and all three Authorization variants.
2. Add a deterministic representative timing-parity test/benchmark at the handler/service seam. Document sample count and a justified threshold; avoid flaky wall-clock assertions by controlling fake dependency latency.

**Acceptance evidence:** targeted Campaign transport test command with the full matrix and documented timing result.

## Finding 3 — Official `make verify` fails

**Evidence:** `make verify` stops at `gosec ./...` with 16 reports. Scope-relevant reports include G304 in `cmd/campaign-seed/main.go` and G112 on the `cmd/server` HTTP server. The remaining findings are outside this Work Unit and must be triaged by their owning work rather than hidden.

**Build/Patch work:**

1. Address/safely justify scope-relevant findings without weakening operator input validation or Tier-0 boundaries; add tests for any behavior changed.
2. Coordinate the pre-existing findings with their owners, or update the repository's agreed static-analysis baseline through the authority that owns it. Do not add blanket `#nosec` suppression.
3. Re-run `make verify` only once its known blockers are resolved; preserve its full result.

**Acceptance evidence:** `make verify` exits zero, or an owning-authority decision explicitly separates unrelated pre-existing repository blockers while retaining a passing Campaign-relevant security check.
