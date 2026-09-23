# Gap Analysis — WU-S1-003 / EXP-BE-001

> Phase/Stage: Exploration / Stage 2 — Gap Analysis
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

## Area 1 — Active Slice-1 authority and contract baseline

### Current state

- Slice 1 is `CONTRACT_READY`; `BACKEND_VERIFIED`, frontend mock, and
  integration milestones remain unearned. The tracker explicitly permits
  separate downstream backend execution.
  (`docs/project/kencleng-development-tracker.md` §7)
- The active backend work unit is scoped to a minimum persisted public
  Campaign capability, public projection/read behavior, controlled media
  byte delivery, private-storage integration boundary, and narrow
  seeded/operator-assisted setup. Root Caddy, Docker Compose, and MinIO
  policy mutation are explicitly owned by `WU-S1-005`.
  (`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/manifest.md`
  §§Outcome, Scope, Out of Scope, Dependencies)
- `GET /campaigns/{campaignId}` (`getPublicCampaignDetail`) and
  `GET /campaigns/{campaignId}/media/{mediaId}/content`
  (`getPublicCampaignMediaContent`) are settled public operations in the
  split Campaign OpenAPI source. Both explicitly use `security: []`, define
  `200`/public `404`/`503` responses, and require the public cache-control
  header. (`api/openapi/campaign.yaml` paths at lines 110–141 and 254–291)
- The public response graph is a standalone closed allowlist, including
  exact nested objects and tagged funding/media states. It is not an
  inheritance of historical internal `Campaign` or Organization schemas.
  (`api/openapi/campaign.yaml` `PublicCampaignDetail` through
  `PublicCampaignDonationAction`, lines 629–854)
- Reconciled Campaign feature records, `INV-campaign-14`, and the threat
  model all match the contract boundary: only internally `published`
  fundraising Campaigns are public in Slice 1; non-public/absent/malformed
  detail requests are indistinguishable; media requires an origin parent and
  member recheck; and storage failure of an eligible resource is `503`, not
  false absence. (`docs/spec/4-campaign/features/02-campaign-detail-listing.md`
  §§Auth–Test checklist; `features/03-campaign-media.md` §§Auth–Test
  checklist; `invariants.md` `INV-campaign-14`; `threat-model.md`
  Slice-1 operation sections)

### Requirement

The Product/MVP authority requires a genuine persisted, public-safe Campaign
experience with truthful media, lifecycle, funding, organizer provenance,
and a non-fake next action; initial Campaign/Organization/media setup may be
seeded or operator-assisted only when it is real persisted application data
and retains truthful claims. (`docs/product/mvp-scope.md` §5 Stage A and §6;
`docs/product/mvp-delivery-slices.md` §4, especially Minimum enabling
operations and Slice 1 completion evidence)

The active delivery invariant requires backend enforcement of public
eligibility before mapping, exact allowlist projection, decimal funding
meaning, uniform public `404`, `503` dependency distinction, controlled
private-media bytes, and `Cache-Control: private, no-store`.
(`docs/spec/4-campaign/invariants.md` `INV-campaign-14`)

### Gap

Contract and delivery authority are reconciled, but no runtime backend
capability has yet been evidenced. The tracker correctly leaves
`BACKEND_VERIFIED` unearned. Subsequent areas must establish whether live
Campaign persistence, domain/HTTP wiring, controlled storage access, and
seeded state exist; this area found no authority-level reason to reopen the
settled public contract.

### Sniffing

- **Risk:** A backend implementation that maps from internal models or lets
  optional authentication select a richer branch would violate both the
  public allowlist and anti-enumeration boundary. This has public data-leak
  reach, not merely a handler-local consequence.
- **Edge cases:** The contract distinguishes zero from unavailable funding,
  absent from unavailable media, malformed/non-resolvable from non-public
  identifiers, and an eligible dependency outage from a missing resource.
  These distinctions require runtime evidence rather than schema-only
  assertions.
- **Miscontext:** Product permits seeded/operator-assisted setup; it does
  not permit hard-coded frontend facts or turn self-service Campaign/media
  creation into Slice-1 scope.
- **Misleading signals:** `CONTRACT_READY` and generated types validate the
  contract artifact only. They do not establish a route, persistence,
  storage privacy, cache-header preservation, or runtime response parity.
- **Inconsistency:** No conflict was found between the active Product/MVP
  sources, reconciled Campaign delivery records, OpenAPI, tracker, and this
  Work Unit manifest. Historical listing, attachment-list/upload, and
  privileged detail are explicitly `DEFER`, rather than competing Slice-1
  acceptance criteria.

### Code / evidence anchors

- `docs/project/kencleng-development-tracker.md` §7 — milestone honesty and
  downstream execution boundary.
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/manifest.md`
  §§Outcome–Dependencies — active backend Work Unit scope and `WU-S1-005`
  topology ownership boundary.
- `docs/spec/4-campaign/invariants.md` `INV-campaign-14` — runtime
  correctness/security properties requiring later verification.
- `api/openapi/campaign.yaml` `getPublicCampaignDetail`,
  `getPublicCampaignMediaContent`, and `PublicCampaignDetail` public graph
  — exact HTTP and projection correspondence targets.

## Area 2 — Persisted Campaign/steward/media state and seeded setup

### Current state

- Live migrations stop at `000010_widen_auth_tokens_purpose`; they create
  account/auth tables only. There is no migration for `organizations`,
  `campaigns`, or `campaign_attachments`, and no production Campaign or
  Organization domain package. (`backend/migrations/`; `backend/internal/domain/`)
- There is no application seed command, fixture, or operator setup artifact
  for Campaign data. Existing `seed*` helpers occur inside Account tests and
  do not establish a runtime public Campaign. (`backend/Makefile`;
  `backend/internal/domain/account/*_test.go`)
- The historical ERD is useful evidence of intended relational concepts:
  `organizations.name`, Campaign-to-Organization membership, `published`
  status, `published_at`, `deadline`, and `NUMERIC(19,2)` target/collected
  amounts. It is not a live migration or Slice-1 schema authority.
  (`docs/project/kencleng-erd.md` `organizations`, `campaigns`, and
  `campaign_attachments` sections)
- The historical Campaign table has a single nullable `description` and the
  attachment table contains stored-object metadata, but neither captures the
  settled separate required `purpose.content` and `story.content`, nor the
  required media `alt_text`/nullable `caption` public metadata. The ERD also
  calls Campaign media a public bucket, which conflicts with the active
  Slice-1 private-storage contract.
- Existing database access is a `pgxpool` opened at startup; backend
  convention requires `goqu` for repository SQL and forbids `float64` for
  money. (`backend/internal/platform/db/db.go`; `backend/AGENTS.md` §2;
  root `AGENTS.md` Golden rules)

### Requirement

Slice 1 requires enough persisted Organization/steward, Campaign, funding,
lifecycle, and Campaign-media data to render a genuine eligible public
Campaign. Seeded/operator-assisted setup is allowed only when it creates real
persisted data and does not misrepresent manual work as automated/reviewed
capability. (`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/manifest.md`
 Scope; `docs/product/mvp-scope.md` §6; `docs/product/mvp-delivery-slices.md`
 §4 Minimum enabling operations)

The public contract requires distinct required organizer text objects for
purpose/story, a public steward `id`/`name`, published fundraising lifecycle
times, decimal-string funding truth, and truthful media state/item metadata.
(`api/openapi/campaign.yaml` `PublicCampaignDetail` through
`PublicCampaignMediaItem`; `docs/spec/4-campaign/invariants.md`
 `INV-campaign-14`)

### Gap

The repository has no live persistence schema or seed path from which either
public operation can read. Historical ERD material is insufficient as-is:
it omits settled public content/accessibility fields and explicitly carries a
now-invalid public-bucket assertion. A future Build must introduce the
minimum schema/data access and repeatable operator/seed evidence while
treating historical self-service/curation/listing breadth as `DEFER`.

### Sniffing

- **Risk:** Seeding only a response fixture or encoding a Campaign in a
  handler would falsely appear to satisfy the page while violating the
  persisted-state requirement. Reusing the ERD's public-bucket assumption
  would directly defeat media withdrawal.
- **Edge cases:** A public Campaign requires a real `published` lifecycle,
  factual zero/over-target decimal values, no media, and media metadata whose
  alt text is non-empty. A seed must also avoid turning absent vs unavailable
  media into the same state.
- **Miscontext:** Historical ERD tables document a broader operational model,
  not authorization to implement Campaign creation, curation, attachment
  upload, public listing, or a direct bucket in Slice 1.
- **Misleading signals:** `docs/project/kencleng-erd.md` contains credible
  table definitions and `NUMERIC(19,2)`, but no matching migration exists;
  its `campaign_attachments` definition cannot by itself satisfy the active
  public metadata or privacy contract.
- **Inconsistency:** The ERD statement that Campaign media is public-bucket
  content conflicts with `INV-campaign-14`, Task 03, and the current OpenAPI
  controlled same-origin content path. Product/spec/contract direction takes
  precedence; the historical ERD needs adaptation if used as implementation
  evidence.

### Code / evidence anchors

- `backend/migrations/000001_create_users.up.sql` through
  `000010_widen_auth_tokens_purpose.up.sql` — live migration baseline;
  no Campaign/Organization/media schema exists.
- `docs/project/kencleng-erd.md` `organizations`, `campaigns`, and
  `campaign_attachments` — historical relational vocabulary only; contains
  the public-bucket contradiction and omits active public fields.
- `backend/internal/platform/db/db.go` `Open` — current shared pool boundary.
- `backend/Makefile` migration targets — manual migration execution path;
  no seed target exists.
- `../harscode-workspace/best-practices/go/decimal-and-money.md` — confirmed
  carry-forward rule: no `float64` across amount parsing, persistence,
  calculation, or serialization; explicit rounding and decimal equality.
- `../harscode-workspace/best-practices/postgresql/migrations-safety.md` —
  migration additions need genuine reversible up/down evidence and must not
  edit already-applied migrations.
- `../harscode-workspace/best-practices/go/goqu-query-builder.md` — use
  prepared, parameterized goqu paths for runtime repository execution; do
  not interpolate values into SQL.

## Area 3 — Campaign domain service, public projection, and HTTP transport

### Current state

- `backend/internal/domain/` contains only the Account domain. There is no
  `campaign` package, Campaign repository port/adapter, public read model,
  service, or public eligibility sentinel vocabulary.
  (`backend/internal/domain/` file inventory)
- `cmd/server/main.go` constructs only Account dependencies and registers
  health, documentation, `/auth/`, and authenticated `/account/` routes.
  Neither settled Campaign operation is registered.
  (`backend/cmd/server/main.go` `run`, router setup)
- The established Account shape has a domain repository port, a goqu/pgx
  `RepositoryDB` adapter, a service with injected dependency seams, and
  transport handlers that depend on a narrow interface for unit testing.
  (`backend/internal/domain/account/repository.go`,
  `repository_db.go`, `service.go`; `backend/internal/transport/http/account_profile.go`)
- `net/http` Go 1.22 method patterns and `http.ServeMux` are already the
  routing convention. Public success serialization is handled by
  `writeJSON`; Problem Details by `WriteProblem`. The latter intentionally
  hides unknown error internals but does not apply public Campaign
  `Cache-Control` headers. (`backend/cmd/server/main.go` router;
  `backend/internal/transport/http/auth_login.go` `writeJSON`;
  `errors.go` `WriteProblem`)
- Existing `RequireSession` rejects missing/invalid credentials before a
  handler runs. It is appropriate for account routes but cannot be a
  middleware boundary for public Campaign operations, whose contract
  explicitly says optional Authorization cannot change visibility or output.
  (`backend/internal/transport/http/account_security.go` `RequireSession`;
  `docs/spec/4-campaign/invariants.md` `INV-campaign-14`)

### Requirement

The active operations must resolve a Campaign and evaluate public eligibility
before a closed public projection is mapped; absent, malformed/non-resolvable,
and ineligible Campaigns must produce the same public `404` response. The
projection must have only its exact allowed public fields and backend-authored
funding/action/media meaning. (`docs/spec/4-campaign/features/02-campaign-detail-listing.md`
 §§Auth–Behavior; `INV-campaign-14`; `api/openapi/campaign.yaml`
 `getPublicCampaignDetail` and `PublicCampaignDetail`)

The media operation must be an unauthenticated controlled-byte endpoint with
the same parent eligibility and media-member lookup boundary, same public
`404`, eligible-dependency `503`, JPEG/PNG-only success, and no-store headers.
(`docs/spec/4-campaign/features/03-campaign-media.md` §§Auth–Behavior;
`api/openapi/campaign.yaml` `getPublicCampaignMediaContent`)

### Gap

There is no Campaign domain or HTTP implementation to enforce or prove any
runtime public boundary. Existing Account abstractions are implementation
evidence, not a ready Campaign component: they do not carry the public
projection, public error classes, public cache header, optional-auth
invariance, or controlled media behavior. Router wiring and tests for both
operations are absent.

### Sniffing

- **Risk:** Attaching these routes to `RequireSession`, checking an optional
  token to select fields, or serializing a persistence entity directly would
  disclose lifecycle/Organization/operational data and break the uniform
  public boundary.
- **Edge cases:** An invalid UUID is contractually collapsed into public
  `404`, not a generic validation `400`; the same applies to missing parent,
  non-public parent, and non-member media. A success/error must receive
  no-store before headers are written.
- **Miscontext:** The Account transport's authentication middleware and
  error-map semantics solve Account cases. They are not a reason to add
  authentication, privileged variants, or account role behavior to Slice-1
  Campaign reads.
- **Misleading signals:** `WriteProblem` provides RFC 9457-shaped generic
  errors and `writeJSON` serializes JSON, but neither alone ensures exact
  `PublicCampaignNotFound`/`PublicCampaignUnavailable` vocabulary, header
  parity, or closed JSON projection. A handler that merely returns a 404 is
  insufficient.
- **Inconsistency:** `backend/internal/transport/http/doc.go` still calls
  the transport an empty skeleton/health-only despite present Account routes;
  it is stale descriptive text, not an executable or active Slice-1
  authority. No active contract conflict was found.

### Code / evidence anchors

- `backend/cmd/server/main.go` `run` router setup — composition point with
  no Campaign wiring today.
- `backend/internal/domain/account/repository.go` and `repository_db.go`
  `Repository` / `RepositoryDB` — current port/adapter and prepared-goqu
  convention to reopen during Build, without copying Account semantics.
- `backend/internal/transport/http/errors.go` `WriteProblem` and
  `auth_login.go` `writeJSON` — existing response primitives that lack the
  Campaign-specific no-store/error boundary by themselves.
- `backend/internal/transport/http/account_security.go` `RequireSession` —
  explicit negative boundary: it must not gate public Campaign reads.
- `backend/internal/transport/http/account_profile_wire_test.go`
  `newWireServer` — current real-HTTP route/middleware verification pattern.
- `../harscode-workspace/best-practices/restapi/anti-enumeration.md` —
  carry-forward requirement for indistinguishable status/body/header and
  timing parity at the public existence boundary.

## Area 4 — Controlled media storage boundary (backend-owned interface)

### Current state

- The only `internal/platform/storage` source is a package comment; no
  MinIO/S3 client is exposed to a domain, no object-read operation exists,
  and there are no storage tests. (`backend/internal/platform/storage/doc.go`)
- Startup creates a MinIO client only to check that both configured buckets
  exist, then discards it. No Campaign component receives storage access.
  (`backend/cmd/server/main.go` `initMinIO` and `run` step 4)
- The repository's established configuration already supplies both public and
  private bucket names. Current Compose initialization creates both and makes
  `kencleng-public` anonymously downloadable. (`.env.example` MinIO values;
  `docker-compose.yml` `minio-init`)
- Active architecture/spec authority is explicit that Slice-1 Campaign media
  must remain private, be delivered by the same-origin backend operation,
  never redirect/reveal an object URL, and distinguish an eligible storage
  failure as `503`. (`docs/project/kencleng-backend-tech-stack.md` File
  Storage row; `docs/spec/4-campaign/features/03-campaign-media.md`;
  `INV-campaign-14`)
- Topology/policy work is an explicit separate concern: this Work Unit may
  explore the backend integration boundary only, while root Caddy, Compose,
  and MinIO policy mutation are owned by `WU-S1-005`.
  (`WU-S1-003/manifest.md` Out of Scope and Dependencies;
  `docs/project/kencleng-repo-setup.md` §8)

### Requirement

For every request to the media-content operation, backend-origin behavior
must recheck public parent eligibility and member association before bytes
are served; it must return only JPEG/PNG bytes for an eligible member,
uniform public `404` otherwise, and public `503` when eligible metadata
exists but storage/object service cannot serve it. It must not issue a direct
object-storage URL or redirect and must set `Cache-Control: private,
no-store` on `200`, `404`, and `503`.
(`docs/spec/4-campaign/features/03-campaign-media.md` §§Auth–Test checklist;
`docs/spec/4-campaign/invariants.md` `INV-campaign-14`;
`api/openapi/campaign.yaml` `getPublicCampaignMediaContent`)

### Gap

No backend-owned storage integration exists, so the controlled content
operation cannot stream bytes, classify object/dependency failure, or prove
that it never exposes object URLs. Existing topology presently retains an
anonymous public bucket and the known `/api` prefix caveat; these are
material integration constraints but are not mutable in this Work Unit.
The backend work must keep private-object access and origin recheck as its
interface boundary, while `WU-S1-005` owns actual bucket-policy and proxy
conformance evidence.

### Sniffing

- **Risk:** A direct/public/signed object URL would make a known media URL
  fetchable after campaign retraction, bypassing parent/member checks. The
  risk reaches any withdrawn public Campaign media, not just one handler.
- **Edge cases:** Parent ineligible, unknown/non-member media, missing object,
  unavailable storage dependency, cancellation while reading bytes, and an
  unsupported content type must remain distinguishable only where the
  contract allows (`404` versus eligible `503`), without exposing object
  identity or storage errors.
- **Miscontext:** The existence of `MINIO_BUCKET_PUBLIC`, the Compose public
  bucket, and historical attachment metadata do not authorize Campaign media
  to use it. Upload functionality and policy/topology mutation remain
  deferred/out of scope here.
- **Misleading signals:** `initMinIO` succeeding proves only bucket existence;
  it proves neither the effective access policy, private object placement,
  every-request authorization, no redirect, header preservation, nor a
  usable storage client in the Campaign path.
- **Inconsistency:** `docker-compose.yml` currently sets anonymous download on
  `kencleng-public`, while active Slice-1 authority prohibits public-bucket
  Campaign media. This is a real pre-existing topology contradiction,
  explicitly routed to `WU-S1-005`; it must not be silently worked around by
  backend code or changed in this Run.

### Code / evidence anchors

- `backend/internal/platform/storage/doc.go` — empty shared-infrastructure
  placeholder; no live storage interface.
- `backend/cmd/server/main.go` `initMinIO` — startup only validates both
  bucket names and does not retain/pass a client.
- `.env.example` MinIO variables — existing private/public configuration
  inputs; not proof of effective access policy.
- `docker-compose.yml` `minio-init` and `Caddyfile` `/api/*` — topology
  evidence only; `WU-S1-005` owns any write and end-to-end validation.
- `docs/project/kencleng-repo-setup.md` §8 — known `/api` prefix handling
  caveat and root-scoped ownership.
- `../harscode-workspace/best-practices/go/context-propagation.md` —
  storage reads must retain request context cancellation/deadline rather
  than introduce a mid-chain background context.
- `../harscode-workspace/best-practices/go/error-wrapping.md` — dependency
  errors need stable classification/wrapping at the storage boundary rather
  than string matching, so eligible `503` does not degrade into false `404`.

## Area 5 — Backend test and verification baseline

### Current state

- `go test ./...` completed successfully on 2026-09-23. It covers current
  Account/platform/transport packages only; `cmd/server`, `db`, `ratelimit`,
  `scheduler`, and `storage` have no tests, and there is no Campaign package
  to cover.
- Existing Account tests demonstrate useful patterns: table-driven domain
  tests with fakes, handler tests with `httptest.ResponseRecorder`, and
  real-HTTP wire tests that rebuild the exact mux/middleware composition.
  (`backend/internal/domain/account/*_test.go`;
  `backend/internal/transport/http/account_profile_wire_test.go`)
- Existing real-Postgres tests are behind an `integration` build tag and use
  `DATABASE_URL`, manually-applied migrations, random UUID fixtures, and
  cleanup. They skip when `DATABASE_URL` is unavailable. There are no
  Campaign persistence/integration fixtures.
  (`backend/internal/domain/account/*_integration_test.go`, especially
  `repository_db_integration_test.go` `integrationEnv`)
- `backend/Makefile` defines unit, race, contract-tag, and security targets,
  but no source currently has a `contract` build tag and no test covers
  OpenAPI-to-runtime correspondence, Campaign routes, MinIO behavior, or
  Caddy/header preservation. (`backend/Makefile`; backend test-tag inventory)

### Requirement

The reconciled Campaign feature specs and `INV-campaign-14` require runtime
proof of exact allowlist mapping; public absent/non-public/auth parity,
including timing; decimal progress without float conversion; truthful media
states; no-store headers; private storage and retraction behavior; and
eligible storage-failure `503` classification.
(`docs/spec/4-campaign/features/02-campaign-detail-listing.md` Test
checklist; `features/03-campaign-media.md` Test checklist;
`docs/spec/4-campaign/invariants.md` `INV-campaign-14` Verification;
`threat-model.md` Slice-1 residual risks)

`WU-S1-003` must supply backend runtime evidence sufficient for
`BACKEND_VERIFIED`, without claiming frontend, root-proxy, or real topology
integration evidence owned elsewhere. (`WU-S1-003/manifest.md` Completion;
`docs/project/kencleng-development-tracker.md` §7)

### Gap

There are no tests or fixtures for either public Campaign operation, its
persistence projection, funding calculation, public error/header parity,
controlled media read, or route wiring. The current green unit suite proves
only pre-existing Account/platform behavior. No backend test presently
exercises object storage; topology-level private policy and `/api` header
preservation remain deliberately outside this Work Unit and require
`WU-S1-005` evidence.

### Sniffing

- **Risk:** Schema/handler code can look correct while leaking a forbidden
  internal field, returning different cache headers/bodies for non-public
  state, or converting an eligible storage outage into `404`. These failures
  are public confidentiality/truth failures and need negative runtime tests,
  not a happy-path-only check.
- **Edge cases:** Test data must cover published versus each ineligible state,
  malformed IDs, absent parent, non-member media, no media, unavailable
  media, object failure, zero and over-target money, optional/malformed auth
  headers, and a previously-known content URL after retraction.
- **Miscontext:** A passing `go test ./...` and contract-generated TypeScript
  evidence do not validate Go runtime projection, Postgres lookup semantics,
  or storage behavior. Conversely, backend unit/integration evidence cannot
  claim Caddy/MinIO policy conformance.
- **Misleading signals:** `make test-contract` exists, but the current source
  has no `contract`-tagged tests. The test target therefore is not evidence
  that OpenAPI operations are implemented or contract-faithful.
- **Inconsistency:** `backend/AGENTS.md` describes an established
  `testcontainers-go`/integration-tag pattern where applicable, while live
  Account integration tests use the `integration` tag plus an externally
  provisioned `DATABASE_URL` and `go.mod` has no `testcontainers-go`
  dependency. This is a verification-infrastructure discrepancy to resolve
  during Techplan/Build selection, not a reason to weaken required
  persistence evidence.

### Code / evidence anchors

- `backend/Makefile` `verify`, `test-unit`, `test-race`, `test-contract` —
  available commands; only the unit suite was run in this Exploration.
- `backend/internal/domain/account/repository_db_integration_test.go`
  `integrationEnv` — live manual Postgres integration fixture pattern and
  its environment dependency.
- `backend/internal/transport/http/account_profile_wire_test.go`
  `newWireServer` — executable real-HTTP route/middleware test precedent.
- `backend/internal/platform/storage/doc.go` — no storage test seam exists.
- `docs/spec/4-campaign/invariants.md` `INV-campaign-14` Verification —
  required backend test evidence categories.
- `WU-S1-005` remains the owner of proxy/policy/retraction-through-real-
  topology evidence; this Run records the dependency but makes no topology
  claim.
