# Solutioning — WU-S1-003 / EXP-BE-001

> Phase/Stage: Exploration / Stage 3 — Solutioning
> Work Unit: `WU-S1-003`
> Run: `EXP-BE-001`
> Role: Explorer
> Specialization: backend Campaign domain, persistence, HTTP transport, controlled media delivery, seeded/operator-assisted data
> Participant: Codex CLI agent
> Session: Fresh session
> Author: Codex
> Created: 2026-09-23
> Target revision: `70ef5c6af1a9befce8492512ee11fec4b22f779f`
> Workflow revision: `d46358563942c7e015b97aa7c5767c880ef1bc63`

## Decision frame

No new evidence invalidates the reconciled Product, security, or API contract.
This solutioning selects the narrowest runtime direction that can earn
`BACKEND_VERIFIED`: two public reads over real persisted data, an explicit
public projection, controlled private-object delivery, and an operator-only
seed path. It does not add self-service Campaign/Organization workflow,
listing, donations, closure, uploads, or topology mutation.

`WU-S1-005` retains ownership of root Caddy, Docker Compose, and MinIO
policy. This Work Unit owns only the backend storage-client integration and
the promise that Campaign code requests private objects through it.

## Decision 1 — Minimal Campaign-owned relational state

| Option | Assessment |
|---|---|
| Hard-code a public response or handler fixture | Rejected: it cannot meet the persisted-state requirement and would create a fake implementation signal. |
| Recreate the full historical Organization/Campaign/attachment model | Rejected: it imports explicitly `DEFER` PII, self-service, curation, creation, upload, listing, and lifecycle behavior. |
| Persist only facts needed by Slice 1, retaining historical names only when useful | **Chosen:** genuine data with a forward-compatible domain anchor, without historical breadth. |

Create additive/reversible migrations for narrow `organizations`, `campaigns`,
and `campaign_media` tables. The rows are operator-seeded and require neither
Account users nor a creation workflow.

- `organizations` persists only public steward `id` and `name`; it does not
  add Organization contact/legal/PII fields.
- `campaigns` persists identity, steward reference, title, separate plain-text
  `purpose`/`story`, `status`, `published_at`, historical `deadline` mapped to
  `fundraising_ends_at`, and IDR `NUMERIC(19,2)` target/collected facts. Known
  lifecycle labels may be constrained for valid seed data, but no transition
  API, scheduler, or closure behavior is introduced.
- Funding is an all-or-nothing nullable pair: both absent maps to the settled
  `funding.availability = unavailable`; both present maps to `available`.
  A database check prevents a partial pair. A factual zero remains present;
  target `<= 0` maps to `not_computable`, never to float or fabricated data.
- `campaign_media` persists only opaque private object key, parent membership,
  ordered public metadata (`content_type`, `alt_text`, nullable `caption`),
  and no public object URL. A Campaign-level persisted media state preserves
  the contract distinction between `absent`, `unavailable`, and `available`.
  An object/dependency failure on a content request stays `503`, not `absent`.

This is a persistence implementation of the public contract, never an
internal-record serialization. Later Organization/Campaign slices can extend
the model additively.

## Decision 2 — Campaign service owns public semantics

| Option | Assessment |
|---|---|
| Query Postgres and serialize directly in each handler | Rejected: transport becomes coupled to persistence, and projection/error parity can drift. |
| Have repository adapters return OpenAPI/transport structs | Rejected: persistence would own public HTTP shape and encourage leakage when fields grow. |
| Repository returns narrow internal public-read data; service applies semantics; handler maps a dedicated view to exact wire structs | **Chosen:** keeps the allowlist and error boundary auditable. |

Add `backend/internal/domain/campaign` with entities/read views, a repository
interface, goqu/pgx `RepositoryDB`, domain sentinel errors, and a service.
Repository execution uses `.Prepared(true)` and `%w` wrapping. The service
must enforce these settled rules:

1. Public detail lookup admits only `status = 'published'`; absent and every
   ineligible state have one domain not-found result. The public handler maps
   malformed UUID input to that same result.
2. Mapping uses explicitly selected columns only. Raw status/reasons, audit
   values, PII, `donor_count`, internal IDs, and operational data never enter
   a public response struct.
3. Percentage uses an exact decimal type (the local guidance names
   `shopspring/decimal` as suitable), explicit half-up rounding to two
   fractional places, no cap at 100, and no `float64` anywhere in the path.
4. `donation_action` is the single Slice-1 unavailable state/reason and has
   no URL.
5. Media-content lookup combines parent ID, member ID, and `published`
   eligibility before any storage read. Optional Authorization is neither
   parsed nor consulted.

The public routes use existing Go `net/http` patterns but not
`RequireSession`. Transport-owned exact JSON and Problem Details mapping set
`Cache-Control: private, no-store` before every public `200`, `404`, or `503`.
Campaign error handling uses stable sentinel classes, never provider/driver
error strings.

## Decision 3 — Concrete private MinIO client with a narrow read seam

| Option | Assessment |
|---|---|
| Public bucket URL, presigned URL, or redirect after metadata check | Rejected by the settled revocation contract: later parent/member checks are bypassed. |
| Call MinIO directly in the handler | Rejected: mixes public authorization, persistence, streaming, and provider errors. |
| Broad multi-provider storage framework | Rejected: backend architecture does not justify interchangeable provider abstraction. |
| Concrete platform MinIO client, injected through only Campaign's required object-read behavior | **Chosen:** one provider/configuration with deterministic service tests and no broad adapter framework. |

Implement `internal/platform/storage` around the configured MinIO dependency
and use `MINIO_BUCKET_PRIVATE` exclusively for Campaign media. Server wiring
retains/injects this client rather than discarding it after bucket checks. The
Campaign-facing read seam opens an opaque private object with request context;
it offers neither URL signing nor public URL construction.

The concrete client establishes object availability before HTTP headers are
committed. An absent/unreadable object or provider outage for eligible metadata
becomes the unavailable class and public `503`. The handler streams only
persisted/validated JPEG-or-PNG media and closes the reader. A failure after
bytes are committed cannot become a valid Problem Details body; it terminates
the stream, is logged without object key/credentials, and must never become
false `404`.

This Work Unit does not alter policy, Compose, Caddy, or the existing public
bucket's use by other future concerns. `WU-S1-005` must prove actual private
policy, `/api` forwarding, and no-store preservation; backend proof is limited
to private-object use, opaque metadata, and per-request database recheck.

## Decision 4 — Opt-in operator seed command, not migration or startup data

| Option | Assessment |
|---|---|
| Insert sample Campaign content inside migrations | Rejected: migrations own schema evolution and must not silently create mutable product content. |
| Ad-hoc SQL plus manual object upload | Rejected: not repeatable, weak evidence that metadata/private object agree, and invites public-bucket mistakes. |
| Self-service creation/upload or generic admin workflow | Rejected as out of Slice 1. |
| One operator command validates a public seed manifest and optionally writes a local JPEG/PNG to private storage | **Chosen:** real data without inventing campaign copy or user-facing operations. |

Add a backend seed command separate from server startup and `golang-migrate`.
It consumes operator-supplied public-safe Campaign/steward facts, an explicit
media state, and an optional local JPEG/PNG. `available` requires valid media
metadata and private-object write; `absent`/`unavailable` never fabricate an
item. The command reports the persisted Campaign UUID and is repeatable only
with an explicit operator-selected ID/replace policy, never a startup upsert.

It creates no verification, curation, account, donation, closure, public
object, or active Donate action. Manual application of migrations remains
human-controlled; automated tests use isolated state.

## Decision 5 — Layer backend evidence; hand topology evidence to WU-S1-005

| Option | Assessment |
|---|---|
| Generated OpenAPI/types plus a happy-path handler test | Rejected: no runtime proof of allowlisting, public parity, database filtering, or controlled bytes. |
| Test only against root Compose stack | Rejected: backend correctness would be obscured by topology work owned by `WU-S1-005`. |
| Focused unit/wire tests plus isolated real-Postgres/private-object adapter tests, with topology handoff | **Chosen:** earns backend evidence while preserving ownership. |

The downstream Techplan must retain these layers:

1. **Domain tests:** decimal progress (zero/target/over-target/not-computable),
   exact mapping, provenance, no-action state, media state, and storage-error
   classification.
2. **Persistence tests:** real Postgres validates public-status/member
   predicates, nullable funding-pair integrity, ordering, and migration
   up/down behavior. New integration tests use the integration-tag/
   `testcontainers-go` direction in `backend/AGENTS.md` where applicable;
   Account's `DATABASE_URL`-only tests are a pre-existing discrepancy, not a
   reason to omit isolated persistence proof.
3. **HTTP/wire tests:** exact allowlist, Problem Details media type,
   no-store on `200`/`404`/`503`, identical public body/header outcomes for
   malformed/absent/non-public/non-member cases, ignored optional auth, and
   actual mux registration. Timing parity remains a measured Testing concern.
4. **Storage adapter tests:** isolated private-object success and missing/
   provider-failure classification occur before headers; fakes exercise
   service branches without topology dependency.
5. **Coordination:** `WU-S1-005` proves root bucket policy, Caddy `/api`
   forwarding, and proxy header preservation. Neither Work Unit claims the
   other's evidence.

## Carry-forward risks and non-goals

- Retraction is tested as a fresh request after persisted visibility changes;
  it does not promise retroactive recall of an in-flight/already-downloaded
  response.
- The single filtered lookup reduces not-found branches, but timing parity
  still needs dedicated Testing evidence.
- Public-read rate limiting is deferred pending backend/proxy attribution
  facts; this direction does not add it.
- Closed/result state remains non-public until later Slice reconciliation.
- No frontend, Product/UI copy, Caddyfile, Compose, or MinIO policy change is
  authorized by this direction.

## Recommended execution order

1. Add migrations, exact decimal dependency/data model, and reversible
   migration evidence; do not apply a shared/manual DB migration in-agent.
2. Build Campaign repository/service and pure projection/error tests.
3. Build concrete private storage client, narrow read seam, and adapter tests.
4. Add routes/router wiring and transport/wire tests.
5. Add isolated persistence/storage integration evidence and seed command.
6. Run scoped backend checks selected by Techplan; hand proxy/policy proof to
   `WU-S1-005` before any integrated claim.

## Human decision

None. This direction follows current Product/MVP, reconciled contract,
backend architecture, and Work Unit ownership. Human gates remain manual
database/index application and normal Tier-1 review/merge, not this Run.

## Open / deferred

- Operator supplies actual seed content/media; this Run does not invent it.
- Exact `testcontainers-go` bootstrap mechanics are selected during Techplan
  against the backend AGENT rule and local execution environment.
- `WU-S1-005` owns effective MinIO policy, Caddy prefix routing, and no-store
  preservation through the real topology.

## Phase handoff

- **Completed:** authority/contract review, persistence/seed/domain/HTTP/
  storage-boundary gap analysis, and selected backend delivery direction.
- **Artifacts:** `evidence/gap-analysis.md`; `evidence/solutioning.md`.
- **Human decision:** none.
- **Open / deferred:** operator-owned seed content/media; integration-test
  bootstrap mechanics; root MinIO policy, Caddy routing, and proxy-header
  proof owned by `WU-S1-005`.
- **Recommended next step:** Techplan synthesis for `WU-S1-003`.
- **Session transition:** Techplan may continue in this session because the
  current authorities, code anchors, and selected direction remain explicit
  and there were no material dead ends or redirections.
- **Context pointers:** `docs/product/mvp-delivery-slices.md` §4;
  `docs/spec/4-campaign/invariants.md` `INV-campaign-14`;
  `docs/spec/4-campaign/features/02-campaign-detail-listing.md`;
  `docs/spec/4-campaign/features/03-campaign-media.md`;
  `api/openapi/campaign.yaml` public operations/schemas;
  `backend/cmd/server/main.go`; `backend/internal/platform/storage/doc.go`;
  `WU-S1-003/manifest.md`; `WU-S1-005` ownership boundary.
