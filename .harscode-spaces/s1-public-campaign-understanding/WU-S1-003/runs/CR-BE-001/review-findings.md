# CR-BE-001 — Independent Code Review Findings

> Phase: Code Review
> Work Unit: `WU-S1-003`
> Run: `CR-BE-001`
> Author: Codex CLI agent
> Role / Specialization: Reviewer / Independent Code Review — Backend Campaign public delivery
> Session: Fresh independent Code Review session
> Created: 2026-09-23
> Model / Reasoning: `gpt-5.6-terra` / `high`
> Target revision: `a053b48ef3fef35a073c5fe57e5ee4581f848e07`
> Workflow revision: `d46358563942c7e015b97aa7c5767c880ef1bc63`

## Scope reviewed

Current backend Campaign delivery diff relative to the approved-Techplan target:
`backend/migrations/000011_create_public_campaigns.*.sql`, `backend/internal/domain/campaign/**`,
`backend/internal/platform/storage/private_reader.go`,
`backend/internal/transport/http/campaign_public*.go`, `backend/cmd/server/main.go`,
`backend/cmd/campaign-seed/**`, and the direct decimal dependency update.

The review inspected live code and the approved `TP-BE-001` Techplan, rather
than relying on the Build report. No production file was edited in this Run.

## 1. Safety

No findings.

The public repository applies the `published` predicate before projection and
the content lookup applies parent eligibility plus media membership before the
private object read. The service collapses absent/ineligible resources to the
public not-found sentinel, maps eligible dependency failures to unavailable,
propagates request context into database and object calls, and closes both
pre-stream invalid content and post-stream readers. No new shared mutable
state, goroutine lifecycle, or money-concurrency path was introduced.

The relevant specialized security/test focus (anti-enumeration, controlled
media/retraction, decimal truth, and isolated dependency verification) is
already explicitly present in Techplan §12/Test Focus Pointer. No Techplan
drift was found in this pass.

## 2. Quality

No findings.

The change keeps Campaign read semantics in the domain service, concrete
goqu/pgx and MinIO mechanics at their adapters, and HTTP wire/error handling
at transport. The public response uses dedicated closed wire structures;
object keys cannot enter the projection. Names and comments make the public
boundary and operator-only seed path clear. The duplicate-looking unavailable
mapping around object reads does not alter behaviour and is not material
enough to require a patch.

## 3. Stack-Specific Best Practices

No findings.

Routed and applied sources:

- `best-practices/go/context-propagation.md` and `go/defer-pitfalls.md`: the
  request context reaches pgx/MinIO; resource cleanup is scoped to the
  relevant operation and has no defer-in-loop accumulation.
- `best-practices/go/error-handling.md` and `go/error-wrapping.md`: internal
  causes retain `%w` chains while the transport emits only generic public
  Problem Details.
- `best-practices/go/authorization-and-idor.md` and
  `restapi/anti-enumeration.md`: public eligibility and membership are checked
  server-side and failure classes use the same not-found response without an
  Authorization-dependent branch.
- `best-practices/go/file-upload-handling.md`: the operator-only local media
  path derives accepted JPEG/PNG type from decoded image content, regenerates
  an opaque key, and streams to private storage rather than accepting a
  supplied object path.
- `best-practices/go/decimal-and-money.md`: funding uses decimal parsing and
  explicit two-place rounding; no `float64` is introduced.
- `best-practices/postgresql/migrations-safety.md` and
  `restapi/openapi-spec-first-drift.md`: the additive paired migration is
  child-before-parent reversible in shape, and the handler implements the
  already reconciled split OpenAPI operation/error schemas without changing
  the contract.

The still-required disposable PostgreSQL/MinIO migration and adapter evidence
is intentionally assigned to Testing by the approved Techplan; it is not a
Code Review finding.

## 4. Consistency

No findings.

The implementation follows `AGENTS.md` and `backend/AGENTS.md`: goqu prepared
queries, decimal money handling, `%w` error chains, explicit public
authorization/anti-enumeration checks, no sensitive client errors, standard
`net/http` routing, and no write to fenced Tier-0 paths. It also matches the
approved Techplan and `INV-campaign-14`/Campaign feature 02–03: explicit
allowlist projection, private controlled media route, public no-store header,
and no direct object URL or redirect.

## Verification executed during Review

- `GOCACHE=/tmp/kencleng-review-go go test -count=1 -timeout=30s ./internal/domain/campaign`
  — domain projection/funding/media classification review question — PASS.
- `GOCACHE=/tmp/kencleng-review-go go test -count=1 -timeout=30s -run '^TestPublicCampaign' ./internal/transport/http`
  — public handler closed-wire, no-store, shared 404, safe 503, stream/close
  review question — PASS.
- `GOCACHE=/tmp/kencleng-review-go go test -count=1 -timeout=30s ./internal/transport/http ./cmd/campaign-seed ./internal/platform/storage`
  — attempted broader touched-package check — NOT USABLE in this sandbox:
  an existing Account wire test in `internal/transport/http` panicked while
  `httptest.NewServer` attempted an IPv6 loopback listener (`operation not
  permitted`). This is unrelated to the Campaign-targeted test above and is
  not treated as a Campaign finding.
- `gofmt -d` on all changed Go files — no output.
- `git diff -- backend` after review — no working-tree backend change.

No container/runtime or topology command was run. PostgreSQL/MinIO integration,
migration up/down/re-up, race, broad regression, timing-parity, and proxy/policy
evidence remain Testing or WU-S1-005/downstream responsibilities as defined by
the Techplan.

## Verdict

Approve.

No blocking or non-blocking finding requires a Build patch. Therefore no
`patch-plan.md` is created for this Run.

## Phase handoff

- Completed: independent four-pass Code Review with an approve verdict.
- Artifacts: `runs/CR-BE-001/review-findings.md`; no patch plan required.
- Human decision: none.
- Open / deferred: independent Testing must execute the approved isolated
  PostgreSQL/MinIO, migration reversibility, anti-enumeration/timing,
  cancellation, broad regression, and race evidence. WU-S1-005/downstream
  retains private-bucket policy, Caddy/header preservation, and integrated
  browser/topology proof.
- Recommended next step: start a fresh independent Testing session.
- Session transition: fresh Testing session for independence.
- Context pointers: `TP-BE-001/techplan.md`, this review artifact, and the
  backend diff anchors named in Scope reviewed.
