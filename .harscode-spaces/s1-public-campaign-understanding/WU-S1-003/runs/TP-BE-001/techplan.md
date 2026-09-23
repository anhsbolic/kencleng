# Tech Plan: Slice 1 Backend Public Campaign Delivery

> Phase             : Techplan
> Ticket            : WU-S1-003
> Work Unit         : `WU-S1-003`
> Run               : `TP-BE-001`
> Author            : Codex CLI agent (Planner)
> Model             : `gpt-5.6-terra`
> Reasoning         : `high`
> Session           : Continue existing healthy EXP-BE-001 session
> Created           : 2026-09-23
> Target revision   : `a053b48ef3fef35a073c5fe57e5ee4581f848e07`
> Workflow revision : `d46358563942c7e015b97aa7c5767c880ef1bc63`
> Status            : Draft
> Approach          : Dua public read berbasis data persisted, explicit projection, controlled private-object delivery, dan opt-in operator seed tanpa mengambil kepemilikan topology.
> Refs              : `EXP-BE-001` gap analysis/solutioning; `WU-S1-002/TP-001`; `WU-S1-002/TST-001`; `WU-S1-005/EXP-TOP-001`; Product/MVP Slice 1; `INV-campaign-14`; Campaign feature 02/03; `api/openapi/campaign.yaml`; `docs/project/kencleng-integration-map.md`

---

## 1. Background

Slice 1 membutuhkan public visitor yang dapat membaca satu Campaign eligible
yang benar-benar persisted, beserta steward, funding, provenance, media, dan
next action yang jujur. `WU-S1-002` sudah memperoleh `CONTRACT_READY`, tetapi
backend saat ini hanya memiliki Account domain: belum ada migration Campaign,
storage read client, route, atau runtime evidence untuk kedua public operation.

Work Unit ini membangun minimum backend-owned capability untuk
`BACKEND_VERIFIED`. Ia tidak mengubah Product/API contract dan tidak mengklaim
frontend, Caddy, policy MinIO, atau real same-origin integration. Public
eligibility, closed projection, decimal funding truth, media membership, dan
public Problem Details tetap menjadi tanggung jawab backend; Caddy `/api`
forwarding dan effective anonymous policy bucket tetap milik `WU-S1-005`.

## 2. Scope

**In scope:**

- Additive/reversible PostgreSQL state untuk public steward, Campaign,
  funding/lifecycle, dan media minimum Slice 1.
- `campaign` domain read repository/service, exact public projection, decimal
  progress, public eligibility, dan sentinel error boundary.
- Concrete MinIO private-object reader yang di-inject ke Campaign domain,
  controlled media byte delivery, dan server composition/route wiring.
- Opt-in operator command untuk membuat persisted seed Campaign (dan optional
  private JPEG/PNG) dari supplied manifest; tidak ada startup seed.
- Unit, wire, real Postgres, dan isolated MinIO adapter evidence yang dibutuhkan
  untuk backend-owned behavior, plus migration reversibility evidence.

**Out of scope (explicit):**

- Perubahan `api/**`, Product/MVP, Campaign spec/threat model, frontend, atau
  generated frontend types; kontrak `WU-S1-002` dikonsumsi apa adanya.
- Public/privileged listing, Campaign/Organization self-service, curation,
  upload API/UI, scheduler, closure/result, donation flow, dan Account work.
- `Caddyfile`, `docker-compose.yml`, MinIO bucket policy, proxy cache/header
  configuration, dan bukti topology end-to-end (`WU-S1-005`).
- Direct/public/signed object URL, redirect media, storage-provider framework
  multi-provider, dan global/public-read rate limiting.
- Perubahan Tier-0 `crypto`, `auth` core, donation ledger, atau disbursement
  state machine; tidak ada protected write yang direncanakan.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | Slice 1 public Campaign state harus real persisted state; operator-assisted setup boleh hanya bila tidak berpura-pura menjadi self-service/verified capability. | `docs/product/mvp-scope.md` §5–6; `docs/product/mvp-delivery-slices.md` §4 |
| Q2 | Detail public hanya mengizinkan Campaign `published` fundraising dan exact closed `PublicCampaignDetail`; optional Authorization tidak boleh mengubah hasil. | `INV-campaign-14`; feature 02; `getPublicCampaignDetail` |
| Q3 | Funding menggunakan decimal IDR, zero tetap factual, progress uncapped dan backend-authored, tanpa `float64` atau `donor_count`. | `INV-campaign-14`; feature 02; `go/decimal-and-money.md` |
| Q4 | Metadata/bytes media mematuhi parent public eligibility dan member association; private object tidak boleh menjadi public URL/redirect. | `INV-campaign-14`; feature 03; `getPublicCampaignMediaContent` |
| Q5 | Absent, malformed, dan ineligible/non-member public resource indistinguishable sebagai `404`; eligible dependency/object failure adalah `503`; semua `200`/`404`/`503` membawa no-store. | `INV-campaign-14`; feature 02/03; threat model; OpenAPI responses |
| Q6 | Contract/topology boundary tetap jujur: backend memakai `MINIO_BUCKET_PRIVATE`, tetapi policy private dan `/api` proxy/header preservation dibuktikan oleh `WU-S1-005`, bukan Work Unit ini. | WU-S1-003 manifest; WU-S1-005 solutioning; integration map |
| Q7 | Karena public-data/authorization boundary ini threat-model-raised Tier 1, milestone hanya dapat diklaim dari executable backend evidence, independent review/testing, dan human review; bukan dari files semata. | `docs/kencleng-agentic-workflow.md` §4, §15; WU-S1-003 Completion |

## 4. Rules & Validation

- **R1 — Minimal persisted state and seed truth.** Migration baru membuat hanya
  `organizations`, `campaigns`, dan `campaign_media` yang dibutuhkan Slice 1.
  `campaigns` menyimpan steward reference, title, purpose, story, status,
  published/deadline times, nullable-pair funding IDR, dan declared media
  state; `campaign_media` menyimpan parent, opaque private `object_key`, stable
  display order, MIME, alt text, dan nullable caption. Tidak ada direct URL,
  PII Organization, Account coupling, creation workflow, atau sample row di
  migration. Operator seed command harus fail closed untuk manifest/media tak
  valid; data seed hanya ditulis melalui explicit invocation, bukan startup.
- **R2 — Public eligibility and projection.** `GetPublicDetail` hanya
  mengembalikan row `status = 'published'` dengan lifecycle fundraising.
  Repository memilih explicit columns; service membuat dedicated public view;
  transport serializes dedicated wire struct. Hasil success berisi tepat sembilan
  top-level property contract (`id`, `title`, `purpose`, `story`, `steward`,
  `lifecycle`, `funding`, `media`, `donation_action`) dan tidak dapat
  memuat status/raw reasons, contact/legal data, actor/audit IDs, donor count,
  atau metadata operasional.
- **R3 — Funding and public semantic truth.** Pair target/collected yang keduanya
  `NULL` menjadi funding `unavailable`; pair parsial melanggar database
  constraint. Pair present memakai exact decimal end-to-end: zero tetap
  `"0.00"`; target zero menghasilkan `not_computable` dan percentage `null`;
  target positif menghitung `collected / target * 100`, half-up dua digit,
  tanpa cap, dengan relationship below/equal/above target. `donation_action`
  selalu unavailable `donation_flow_not_available` tanpa activation target.
- **R4 — Truthful media metadata.** Declared `absent`/`unavailable` media selalu
  memiliki empty public items; `available` hanya dapat diproyeksikan ketika
  terdapat ordered member metadata. Item public hanya memuat contract fields,
  termasuk opaque same-origin `/api/campaigns/{campaignId}/media/{mediaId}/content`,
  persisted JPEG/PNG type, non-empty alt text, nullable caption, dan
  `source: organizer`. `unavailable` tidak dipakai untuk menyamarkan object
  read error dari content operation.
- **R5 — Uniform public boundary.** Malformed UUID, absent Campaign,
  non-published Campaign, absent media, dan media yang bukan anggota parent
  memakai satu public not-found class, identical RFC 9457 body/header per
  operation; optional/malformed Authorization is not parsed or consulted.
  Eligible detail dependency failure dan eligible metadata yang object/dependency
  read-nya gagal memakai public unavailable `503`, tanpa provider/DB/object-key
  detail. Handler menetapkan `Cache-Control: private, no-store` sebelum menulis
  setiap public `200`, `404`, dan `503`.
- **R6 — Controlled content delivery.** Content service melakukan one filtered
  lookup yang mencakup Campaign ID, media ID, `published` eligibility, dan
  membership sebelum memanggil storage. Storage read selalu memakai request
  context dan `MINIO_BUCKET_PRIVATE`; object availability/MIME is established
  before headers commit. Hanya JPEG/PNG yang persisted dan validated dapat
  di-stream; reader ditutup. Failure setelah response body mulai dikirim tidak
  dapat ditulis ulang menjadi Problem/404: stream dihentikan dan error dicatat
  tanpa key/credential/PII.
- **R7 — Migration and operator safety.** New migration uses new empty tables
  (so required columns are safe), non-negative `NUMERIC(19,2)` amounts,
  nullable-pair check, foreign keys, media membership/order constraints, and a
  genuinely reversible down order. It never edits 000001–000010. The seed
  command requires an explicit Campaign UUID and rejects collision by default;
  a clearly named `--replace` is the only overwrite path. For available media,
  it validates supplied local JPEG/PNG before writing a generated opaque key to
  the private bucket; no media file is accepted for `absent`/`unavailable`.
- **R8 — Ownership/milestone honesty.** Backend test evidence may establish
  domain, persistence, storage-client, route, and direct-backend HTTP behavior.
  It cannot establish Caddy prefix behavior, persisted MinIO anonymous policy,
  proxy header preservation, browser/frontend behavior, or
  `INTEGRATED_VERIFIED`. Those remain explicit handoff evidence to WU-S1-005
  and downstream integration work.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Add only Slice-1 `organizations`, `campaigns`, `campaign_media` persistence instead of hard-coded response or historical full schema. | Chosen | Satisfies real state with no self-service/PII/curation breadth. Historical ERD public-bucket assertion is `ADAPT`/rejected. |
| D1-alt | Handler fixture/hard-coded public Campaign. | Rejected | Violates product persisted-state requirement and creates false implementation evidence. |
| D2 | Repository returns narrow internal read data; Campaign service owns eligibility/public semantics; handler owns exact wire mapping. | Chosen | Makes allowlist and stable error classification auditable; avoids persistence-to-OpenAPI coupling. |
| D2-alt | Serialize persistence entity in handler or return OpenAPI structs from repository. | Rejected | Risks field leakage and cross-layer drift. |
| D3 | Use `shopspring/decimal` (pinned through `go.mod`) for funding mapping/progress. | Chosen | Existing module has no decimal package; this preserves Q3 and explicit half-up semantics without `float64`. |
| D4 | Concrete MinIO client with a Campaign-facing narrow object-read seam, injected from server wiring. | Chosen | One configured provider is sufficient; permits deterministic domain tests while ensuring production reads use private bucket. |
| D4-alt | Public/presigned URL, redirect, handler-direct MinIO, or generic multi-provider storage framework. | Rejected | Bypasses recheck or creates unjustified cross-cutting abstraction. |
| D5 | Add `integration`-tag isolated Postgres and MinIO testcontainers for new Campaign persistence/storage evidence. | Chosen | Resolves the observed discrepancy between `backend/AGENTS.md` testcontainers direction and older Account tests that use shared `DATABASE_URL`; Campaign tests must not mutate a developer DB. Runtime/container unavailability is reported as unverified, never silently skipped for milestone evidence. |
| D6 | Opt-in `cmd/campaign-seed` manifest command; no migration/startup seed. | Chosen | Provides repeatable real state without inventing creation/upload product surfaces. |
| D7 | Continue direct backend verification independently of topology, but record WU-S1-005 as a gate for topology/integration evidence. | Chosen | Preserves WU ownership and allows truthful `BACKEND_VERIFIED` without claiming proxy/policy proof. |

## 6. Backward Compatibility

- **Existing data:** production migrations currently end at `000010`; no
  Organization/Campaign/media rows exist. Migration `000011` is additive. Its
  down migration removes only tables/indexes it created, in child-before-parent
  order. Human retains authority to apply a shared/manual database migration.
- **API/contracts/clients:** no OpenAPI modification is planned. Server routes
  implement existing unprefixed backend paths; the contract-visible media URL
  retains `/api` because Caddy owns prefix stripping. No old live Campaign
  handler/client must be preserved.
- **Future slices:** later full Organization/Campaign lifecycle and Donation
  work may extend this minimum state additively. They must not reinterpret
  `published` public eligibility, overwrite private media behavior, or reuse
  the seed command as a public/admin workflow without reconciliation.
- **Seed replacement:** default collision failure prevents accidental mutable
  content overwrite. `--replace` is deliberately operator-explicit; it may
  leave an old private object only if post-commit cleanup fails, which is a
  non-public orphan requiring operator cleanup/reporting rather than a broken
  public mapping.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| RISK-1 | Internal fields leak as schema grows. | Medium | High | Dedicated read model/wire structs, explicit select list, forbidden-field negative HTTP tests (R2). |
| RISK-2 | Visibility existence leaks via status/body/header/auth/timing. | Medium | High | Single filtered query/sentinel, identical problem writer/header, ignored auth tests; timing measurement is independent Testing evidence (R5). |
| RISK-3 | Known media URL survives retraction or storage is reachable directly. | Medium | High | Per-request parent/member query, private-bucket-only client, no redirect/no-store; real policy/proxy/retraction handoff to WU-S1-005 (R6/R8). Already downloaded bytes are accepted residual exposure. |
| RISK-4 | Object failure is misrepresented as absent/no-media. | Medium | High | Metadata state distinct from read availability; storage boundary maps eligible read failure to 503 before headers (R4/R6). |
| RISK-5 | Decimal rounding/scale loses funding truth. | Medium | High | Decimal dependency, `NUMERIC(19,2)`, explicit half-up rule, zero/over-target/not-computable table tests (R3). |
| RISK-6 | Seed creates public-looking but invalid or public-bucket media. | Medium | High | Manifest/state validation, MIME/content inspection, private client only, generated opaque key, no migration seed (R1/R7). |
| RISK-7 | Cross-resource replacement leaves stale private object. | Low | Medium | DB visibility changes only after successful seed transaction; best-effort old-object cleanup is reported, never exposed. Operators retain cleanup responsibility. |
| RISK-8 | A stream/read failure after headers commits produces misleading response. | Low | Medium | Establish object availability before headers; terminate post-commit stream without false Problem/404 and log safely (R6). |
| RISK-9 | Backend evidence is overstated as topology/integration evidence. | Medium | High | Explicit file fence and separate WU-S1-005 handoff; reports label those checks deferred/not tested (R8). |

## 8. Interface Contract

**Persistence/data shape:**

- `organizations`: `id` and non-empty public `name` only.
- `campaigns`: UUID; `organization_id`; non-empty title/purpose/story;
  constrained lifecycle status including `published`; `published_at` and
  `deadline`; non-negative `NUMERIC(19,2)` `target_amount`/`collected_amount`
  with `(target_amount IS NULL) = (collected_amount IS NULL)`; declared
  `media_state` (`absent`, `unavailable`, `available`). Only `published` is
  selectable by public reads.
- `campaign_media`: UUID, FK Campaign membership, generated opaque
  `object_key`, contiguous/unique display order per Campaign, allowed MIME
  (`image/jpeg`/`image/png`), non-empty alt text, nullable caption. Object keys
  are never returned by a public model or response.
- Build reopens live migration conventions before selecting exact constraint/
  index names. `000011` must be safe on a fresh empty schema and reversible;
  it must not retrofit historical operational columns.

**API/external interface:** Existing split OpenAPI remains authoritative:

```text
GET /campaigns/{campaignId}
  → 200 PublicCampaignDetail | 404 PublicCampaignNotFound | 503 PublicCampaignUnavailable

GET /campaigns/{campaignId}/media/{mediaId}/content
  → 200 image/jpeg|image/png | 404 PublicCampaignNotFound | 503 PublicCampaignUnavailable
```

Both routes use `security: []`, ignore all optional Authorization variants,
and write `Cache-Control: private, no-store` on every contract response. The
transport uses reusable Campaign-specific response constants/functions so all
not-found paths are byte/header-identical and safe; it must not route Campaign
errors through Account's `MapServiceError`. It sends generic RFC 9457 Problem
Details and never emits raw SQL, storage errors, filesystem paths, object keys,
or credentials.

`content_url` is derived only from public Campaign/media UUIDs in the exact
contract path rooted at `/api`; it is not read from the database and never
becomes a MinIO URL. Backend's direct route stays `/campaigns/...`; WU-S1-005
maps the browser-visible prefix later.

**Cross-layer/business boundary:**

- Campaign service owns public eligibility, read classification, decimal
  progress/action/media semantics, and member recheck.
- `platform/storage` owns concrete S3/MinIO mechanics and stable unavailable
  classification, accepting the caller context and opaque key only.
- HTTP transport owns UUID parsing collapse, headers, Problem Details, JSON,
  MIME streaming, and the un-authenticated route boundary.
- `cmd/server` only composes the pool, retained MinIO client/private bucket,
  repository/service/handler, and routes. It must not gain Campaign business
  logic.
- WU-S1-005 owns Caddy and effective private-bucket policy; frontend owns
  consuming opaque `content_url`. Neither is proof supplied by this Work Unit.

## 9. Architecture / Plan

1. Add decimal/test dependencies and reversible `000011` schema first. Use
   disposable integration infrastructure to prove migrations; do not apply a
   human/shared DB migration from the Build agent.
2. Add `internal/domain/campaign` entities/read interfaces, goqu/pgx adapter,
   and service. The adapter issues prepared parameterized queries with a
   `published` predicate before projection; the service maps narrow data to
   closed public views and sentinel errors.
3. Add the concrete `platform/storage` private MinIO reader. Server startup
   constructs it once, verifies required buckets as today, retains it, and
   injects it into Campaign composition. Only Campaign asks for the private
   bucket; no signed/public URL method exists.
4. Add public handlers and unprotected `net/http` route registrations. Parse
   malformed path UUIDs into Campaign not-found, set no-store before response
   commit, and stream only validated bytes. Do not attach `RequireSession` or
   consult Authorization.
5. Add operator seed command and its tests. Its manifest explicitly supplies
   public-safe stewardship/Campaign/funding/media-state facts; a local media
   option is valid only for available media. Command output is operational
   (persisted Campaign UUID/result), never a fabricated public URL.
6. Produce layered evidence: pure domain and handler/wire tests first,
   isolated Postgres+MinIO integration tests next, then independent testing/
   review. Handoff actual Caddy/policy/header/retraction-through-proxy evidence
   to WU-S1-005; do not update tracker/manifest or self-promote milestone.

```text
anonymous request
      ↓
Campaign public handler (no auth middleware) ── malformed ID ──→ same 404 + no-store
      ↓
Campaign service ── filtered public/membership query ──→ absent/ineligible → same 404
      ↓ eligible narrow read
explicit public mapper ──→ detail JSON + no-store
      │
      └── media content → private storage reader → JPEG/PNG stream + no-store
                              └── eligible object/dependency failure → 503 + no-store
```

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `backend/migrations/000010_widen_auth_tokens_purpose.*.sql` | Current migration ceiling and paired up/down convention. | Add paired `000011` files only; create minimum empty tables/constraints/indexes and reverse them safely. |
| `backend/go.mod` | Current Go/pgx/goqu/MinIO dependencies; no decimal/testcontainers dependency. | Pin decimal plus test-only integration dependencies needed by D3/D5; keep dependency surface narrow. |
| `backend/internal/domain/account/repository.go` — Repository port | Existing domain port/comment style. | Add separate Campaign read/seed repository interfaces; no Account-domain import or cross-domain internals. |
| `backend/internal/domain/account/repository_db.go` — `pgDialect`, prepared query methods | Proven goqu + pgx convention. | Campaign `RepositoryDB` uses package-level Postgres dialect, explicit selected columns, `Prepared(true)`, `db` tags/nullable types, `%w`; never interpolate user input. |
| `backend/internal/domain/account/profile.go` — service read seam | Minimal read service precedent. | Keep Campaign eligibility/projection/media semantic mapping pure/testable with injected repository and object reader. |
| `backend/internal/platform/storage/doc.go` | Empty platform boundary currently. | Replace with concrete MinIO private-object reader and narrow Campaign-facing reader interface; request context, reader closure, stable unavailable classification, no URL signing. |
| `backend/cmd/server/main.go` — `initMinIO`, `run` router composition | Client currently discarded after bucket existence checks; exact public routes absent. | Return/retain one MinIO client/storage adapter, inject `MINIO_BUCKET_PRIVATE` into Campaign service, and register both GET paths directly on main mux without `RequireSession`. |
| `backend/internal/transport/http/errors.go` — `WriteProblem`; `auth_login.go` — `writeJSON` | Existing safe Problem/JSON primitives do not set Campaign no-store or map Campaign errors. | Add Campaign-local responder/wire types that set cache header before calling shared primitive; do not mutate Account mapping semantics. |
| `backend/internal/transport/http/account_profile_wire_test.go` — `newWireServer` | Existing real-HTTP mux/middleware verification pattern. | Add Campaign wire server covering exact registration, no authentication middleware, response shape/header/body parity, and controlled MIME stream. |
| `backend/internal/domain/account/repository_db_integration_test.go` — `integrationEnv` | Existing integration tag and older shared-DB behavior. | Do not reuse shared-DB helper; add isolated Campaign fixture bootstrap using testcontainers, migrations, cleanup, and explicit non-skip evidence. |
| `backend/cmd/campaign-seed` (new) | No current operator seed command. | Add a narrow standard-library-flag command with JSON manifest, explicit UUID, optional `--media-file`, and explicit `--replace`; connect it to same repository/storage boundary, never server startup. |
| `docs/project/kencleng-integration-map.md` — Public Campaign Detail row | Coordination contract already recorded. | Read-only handoff: respect opaque `content_url`, private storage, parent/member recheck, `/api`, and no-store; no map mutation in this Build. |
| `.harscode-spaces/.../WU-S1-005/.../solutioning.md` — Consequences and runtime verification | Explicit topology boundary. | Reopen before Build/Test handoff; do not touch its root files or claim its proxy/policy checks. |

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `backend/go.mod`, `backend/go.sum` | Modify | Decimal and isolated integration-test dependencies only. |
| `backend/migrations/000011_*` | Add | Minimum Campaign/Organization/media state and genuine reversible down migration. |
| `backend/internal/domain/campaign/*` | Add | Entities/read views, repository port/adapter, service, unit and integration tests. |
| `backend/internal/platform/storage/*` | Modify/add | Concrete private MinIO read implementation and adapter tests. |
| `backend/internal/transport/http/campaign_public*.go` and tests | Add | Detail/media handlers, safe response/wire structs, handler/wire tests. |
| `backend/cmd/server/main.go` | Modify | Composition and direct public GET registrations only. |
| `backend/cmd/campaign-seed/*` and tests | Add | Explicit operator seed command and manifest/media validation. |

| File / area intentionally untouched | Why |
|---|---|
| `api/**`, `docs/spec/**`, Product/MVP, `frontend/**` | Settled contract and cross-stack sources are consumed, not changed, by this backend-only Work Unit. |
| `Caddyfile`, `docker-compose.yml`, `.env.example`, `docs/project/kencleng-repo-setup.md` | WU-S1-005 exclusively owns topology/policy/proxy correction. Existing `MINIO_BUCKET_PRIVATE` is sufficient configuration handoff. |
| `backend/internal/platform/auth/**`, `backend/internal/platform/crypto/**`, donation ledger, disbursement state machine | Tier-0/file-path-fenced or outside slice scope. |
| Historical Campaign listing/attachment/upload/curation routes and Account domain | Explicitly `DEFER`; no compatibility implementation exists to extend. |
| `.harscode-spaces/**/manifest.md`, tracker/control surface | Orchestration/human status ownership; Build evidence must not self-promote. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | Migration-focused integration test applies `000011` in an isolated disposable Postgres, asserts tables/PK-FK/check constraints and fresh-schema creation, then verifies child-before-parent down and repeatable re-up. Seed unit tests reject invalid manifest/state/media combinations and never write on validation failure. | Testing | Schema-only review misses partial funding/media integrity and a non-reversible migration. Shared/manual DB application remains Human-owned. |
| R2 | Table-driven service/handler tests cover eligible published detail and inspect decoded JSON key sets recursively for exact public projection; include forbidden internal fields in repository fake input and prove they never serialize. | Build | The most likely confidentiality regression is accidental entity serialization; authored tests must execute in edit loop. |
| R3 | Decimal tests cover unavailable pair, factual zero, target zero/not-computable, below/equal/above target, large values, and half-up two-digit output; static review/tests verify no `float64` funding path. | Testing | Money truth can look correct on ordinary inputs while losing precision or capping over-target facts. |
| R4 | Public detail tests cover available ordered JPEG/PNG items and absent/unavailable empty unions, content URL generation, nullable caption, alt/source fields, and never expose opaque key. | Build | Prevents false no-media state and storage metadata leakage at the public boundary. |
| R5 | Real-HTTP wire tests compare status, complete body bytes, `Content-Type`, and `Cache-Control` for malformed/absent/non-public parent and absent/non-member media; repeat with no, garbage, and syntactically valid Authorization. Independent Testing measures representative timing parity with repeated samples and documents method/results. | Testing | Status-only tests miss enumeration through body/header/auth/timing. Timing is a security-sensitive observable, not a Build-loop claim. |
| R5 | Eligible repository/detail dependency failure and storage error fakes map to 503 with generic Problem Details/no-store; assert raw error strings/object keys/filesystem paths are absent. | Build | Prevents false 404/absence and client-internal leakage. |
| R6 | Storage adapter tests use request cancellation and verify pre-header object stat/read failure classification; wire tests assert JPEG/PNG headers/bytes, no redirect/Location/object URL, and reader close. Integration-tag MinIO test fetches a private seeded object through adapter and tests missing/unavailable object classification. | Testing | The controlled-byte boundary otherwise has no proof beyond a fake and can bypass retraction/error semantics. |
| R7 | `go test ./...` after authored tests; integration-tag isolated Postgres/MinIO testcontainers run with no `DATABASE_URL`; migration down/up test is included. Review seed collision/default failure, explicit `--replace`, private bucket selection, and orphan-cleanup report behavior. | Testing | Verifies actual dependencies and operator safety rather than just Go compilation. Lack of container runtime/network must be reported as not tested and prevents the associated evidence claim. |
| R8 | Inspect diff/file fence and WU-S1-005 handoff; run WU-S1-005's prescribed real topology tests separately before any integrated statement. Do not mutate tracker/manifest in Build. | Human | Prevents ownership drift and false milestone promotion; direct backend tests cannot prove Caddy/policy/frontend integration. |
| R1–R8 | `git diff --check`; targeted `go test ./internal/domain/campaign/... ./internal/transport/http/... ./internal/platform/storage/...`; then independent `go test ./...`, `go test -race ./...`, and applicable `make verify` after tool/container prerequisites are available. | Testing | Cheap integrity plus broad project regression/security evidence; report precisely which executable commands actually ran. |

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Public projection / anti-enumeration | Public data and existence can leak through field, auth, body/header, or timing variance. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/EXP-BE-001/evidence/gap-analysis.md#area-3--campaign-domain-service-public-projection-and-http-transport` | Yes — R2/R5; independent timing/parity evidence belongs to Testing. |
| Decimal funding truth | Precision/rounding errors change public money meaning. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/EXP-BE-001/evidence/solutioning.md#decision-2--campaign-service-owns-public-semantics` | Yes — R3 and decimal-specific independent cases. |
| Controlled media/retraction/cache | Parent/member recheck, private objects, and error/caching behavior determine whether withdrawal works. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/EXP-BE-001/evidence/gap-analysis.md#area-4--controlled-media-storage-boundary-backend-owned-interface` | Yes — R5/R6; direct backend evidence here, topology/policy/proxy evidence delegated to WU-S1-005. |
| Migration/integration isolation | New public persistence schema and real dependency adapters can fail despite pure fakes. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/EXP-BE-001/evidence/gap-analysis.md#area-5--backend-test-and-verification-baseline` | Yes — R1/R7; testcontainers selection D5 resolves the old shared-DB discrepancy. |
| Runtime concurrency | Slice adds read/seed behavior but no donation/disbursement locking or shared mutable hot path. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/EXP-BE-001/evidence/solutioning.md#carry-forward-risks-and-non-goals` | N/A — no new concurrent balance/state transition; `go test -race ./...` remains broad regression evidence, not a substitute for future money-concurrency testing. |

## 13. Open Items

### Active — needs external input or verification

1. **Topology proof remains external and non-blocking for Build.** `WU-S1-005`
   must execute its own Caddy/private-policy/proxy-header/retraction checks on
   a Docker-capable environment. This Work Unit must reference, not duplicate,
   that evidence before any topology or integrated claim. It does not prevent
   direct backend capability/testing evidence.
2. **Operator-owned seed facts remain non-blocking for implementation.** The
   eventual operator must supply truthful Campaign/steward/provenance copy and
   optional owned JPEG/PNG. Build may use clearly test-only fixtures but must
   not invent production claims or represent the command as user-facing setup.

### Resolved — retained as decision history

1. ~~**Public contract shape and security semantics**~~ **RESOLVED — consume
   current `CONTRACT_READY` `WU-S1-002` split OpenAPI/spec/invariant unchanged;
   this plan only implements it.**
2. ~~**Storage ownership boundary**~~ **RESOLVED — backend uses the existing
   `MINIO_BUCKET_PRIVATE` through a narrow client seam; policy/Caddy/proxy
   evidence remains WU-S1-005 (D4/D7).**
3. ~~**Campaign integration-test mechanism**~~ **RESOLVED — use new isolated
   `integration`-tag testcontainers rather than the older shared
   `DATABASE_URL` Account helper (D5).**
