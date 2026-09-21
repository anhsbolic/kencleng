# EXP-001 — Stage 2 Gap Analysis

- Phase/Stage: Exploration / Stage 2 — Gap Analysis
- Author: Explorer
- Created: 2026-09-21
- Model: `gpt-5.6-sol`
- Reasoning: `high`
- Target revision: `c78cae58cedf9cac34dedcd3e068f15cd55d3b6d` (`validation-03-orchestrator-slice-1`)
- Workflow revision: `c8b3eca93e8d265438c2dba8bc37b6b90b251724` (`pilot/orchestrator-v0.1`)
- Work Unit: `WU-S1-001`
- Run: `EXP-001`
- Role: `Explorer`
- Specialization: none

## Scope and method

Dokumen ini menyimpan evidence Stage 2 untuk MVP Slice 1 — Public Campaign Understanding. Setiap area mencatat current state, governing requirement, concrete gap, lima sniffing lenses, dan anchors. Existing specs, contracts, migrations, tests, code, serta Probe 01 diperlakukan sebagai evidence; Product/MVP Authority dan Product Design / Brand Authority tetap upstream.

Stage ini tidak memilih solusi, tidak menetapkan FE/BE decomposition, tidak mengedit authority/contract/implementation, dan tidak membentuk downstream Work Unit graph.

## Area 1 — Product / MVP Authority

### Current state

- `docs/product/product-overview.md` menetapkan Kencleng sebagai evidence-led, bukan persuasion-led. Untuk public campaign, visitor harus dapat membedakan platform/system facts, organizer-provided information, lifecycle state, dan unknown/pending information. Funding progress tidak boleh dibaca sebagai operational progress atau impact.
- `docs/product/mvp-scope.md` menetapkan Stage A sebagai Public Campaign Detail yang deliberate public-safe dan memuat campaign identity/purpose, steward context, truthful media state, lifecycle meaning, funding facts, story/provenance, meaningful unknown/pending states, dan truthful next action.
- `docs/product/mvp-delivery-slices.md` menetapkan Slice 1 sebagai public understanding terhadap persisted eligible campaign. Seeded/operator-assisted setup dibolehkan; account, donation flow, broad discovery, full Organization profile, self-service operations, dan full post-campaign accountability UI tidak termasuk Slice 1.
- Product Authority telah mengadopsi arah keselamatan Probe 01: explicit public-safe Campaign projection, narrow public Organization projection, media visibility mengikuti parent Campaign, media public harus benar-benar revocable, dan donation eligibility backend-authoritative.
- Approved whole-product direction mempertahankan satu public Campaign identity setelah eligible closure. Namun urutan delivery yang lebih baru menempatkan Campaign closure, persistent public result, dan pergeseran ke accountability pada Slice 3, bukan Slice 1.

### Governing requirement

Slice 1 harus membuat skeptical-but-open visitor mampu memahami real persisted public Campaign serta steward, funding, source/provenance, lifecycle, media/absence, dan truthful next-action context tanpa internal-data leakage. Product-visible state harus berasal dari real application state. Tidak boleh ada active Donate action sebelum real Donation Flow tersedia.

Anchors:

- `docs/product/product-overview.md` — §3 Durable product principles, §4 Audience posture, §9 Public campaign experience, §10 Trust and accountability model, §11 Strong current product truths, §13 Public Campaign / MVP slices.
- `docs/product/mvp-scope.md` — §2 MVP inclusion rule, §5 Stage A, §6 Operational breadth, §7 MVP security floor, §8 capability classification.
- `docs/product/mvp-delivery-slices.md` — §2 Cross-slice guardrails, §4 Slice 1.
- `docs/product/probes/01-public-campaign-detail-contract-reconciliation.md` — §1–9, approved product decision and candidate contract evidence; explicitly probe evidence, not canonical contract.

### Concrete gaps

1. Authority menetapkan categories of truth yang wajib dipahami, tetapi exact Slice-1 public facts belum direkonsiliasi menjadi delivery-level acceptance criteria. Contohnya: apakah category, location, beneficiary description, selected public dates, atau organization-review meaning diperlukan pada Slice 1 belum diputuskan oleh active delivery spec.
2. Exact public wording/semantics untuk provenance, unknown/pending, public lifecycle, dan any Organization review signal sengaja belum ditetapkan secara global. Ini adalah design/delivery-resolution gap, bukan izin bagi frontend untuk mengarang semantics.
3. Whole-product approved closure continuity dapat terlihat seperti requirement Slice 1 bila Probe 01 dibaca sendiri. Current sequencing menempatkan closure/result delivery pada Slice 3. Slice 1 harus menjaga compatibility dengan direction tersebut tanpa menyerap full closed/result/accountability scope sekarang.
4. "Truthful next action" sudah wajib, tetapi pada Slice 1 tanpa Donation Flow bentuk user-observable state/action yang tepat belum direkonsiliasi. Satu hal sudah tegas: active/fake Donate control tidak boleh ada.

### Sniffing

- **Risk:** Melebarkan Slice 1 ke closure/accountability atau Donation Flow akan merusak vertical sequencing dan menyamarkan real dependency. Sebaliknya, mengabaikan persistent-public direction dapat menghasilkan public identity/contract yang harus dirombak pada Slice 3.
- **Edge cases:** Campaign persisted tetapi tidak eligible untuk public view; eligible Campaign tanpa media; unavailable/pending provenance information; target/current funding facts belum tersedia; Campaign pernah public lalu retracted berbeda dari normally closed.
- **Miscontext:** Probe 01 membahas published dan closed public states untuk memvalidasi whole lifecycle. Itu bukan bukti bahwa full closed/result UI termasuk Slice 1.
- **Misleading signals:** Field atau historical contract yang factual tidak otomatis perlu dipublikasikan. Donor count, internal Organization status, atau raw Campaign status dapat terlihat berguna tetapi berpotensi menjadi social proof, ambiguous verification, atau internal-state leakage.
- **Inconsistency:** Tidak ada contradiction pada authority tier: Product Authority menginginkan persistent public identity setelah closure, sementara delivery sequencing sengaja menunda closure/result ke Slice 3. Inconsistency muncul hanya jika probe-level lifecycle breadth diperlakukan sebagai current Slice-1 scope.

### Carry-forward anchors

- `docs/product/product-overview.md` — canonical meanings and truth classes; dipakai untuk menilai apakah lower-level artifact mengarang trust/impact/verification semantics.
- `docs/product/mvp-scope.md` — Stage A and MVP inclusion forcing question; dipakai untuk menolak historical breadth yang tidak product/security/enabling-critical.
- `docs/product/mvp-delivery-slices.md` — Slice 1 completion evidence and explicit exclusions; source utama current delivery boundary.
- `docs/product/probes/01-public-campaign-detail.md` — prior validation evidence untuk FE/BE capability hypotheses; bukan source yang boleh memperluas scope.
- `docs/product/probes/01-public-campaign-detail-contract-reconciliation.md` — approved direction untuk future reconciliation, khususnya explicit projections dan media revocation; masih bukan canonical OpenAPI.

## Area 2 — Product Design / Brand Authority

### Current state

- Canonical direction sudah dipilih dan cukup konkret: **Sunlit Editorial / Evidence-Led Optimism**, dengan warm-neutral foundation, restrained Sun/Berry, Newsreader untuk public editorial moments, Instrument Sans untuk body/UI/money/status/provenance, Phosphor untuk utility icons, border/spacing-first grouping, serta restrained radius/elevation.
- `product-design-principles.md` dan `patterns.md` memberi interaction/information rules yang langsung berlaku pada Public Campaign Detail: confidence before conversion, explicit money meaning, ordered transparency, truth-class distinction, unknown as valid state, non-gamified progress, one confident next action, structure-preserving loading, safe error/recovery, dan responsive priority preservation.
- `design-guidelines.md` mendefinisikan visual grammar terpisah untuk Platform Fact, Organizer-Provided Information, Pending/Not Yet Available, System State, funding progress, operational progress, dan reported outcome. Campaign imagery harus real; placeholder harus terlihat sebagai missing media dan tidak boleh menyerupai documentary evidence.
- `page-map.md` mengenali Public Campaign Detail sebagai `Detail` surface untuk Guest/Public Visitor. Exact route, layout, component tree, dan business rules sengaja tidak dimiliki dokumen ini.
- Dua selected-direction images tersedia sebagai approved direction evidence. Public reference terutama merupakan landing/discovery composition, bukan production Public Campaign Detail. Evidence Journal reference menggambarkan post-donation/operational concepts yang bukan Slice-1 contract evidence. Keduanya bukan production assets atau pixel specifications.
- Material design decisions yang masih `OPEN` dan relevan: final provenance terminology, exact campaign-placeholder asset/system, detailed photography governance, exact motion tokens, dan potential project-specific icon tuning.

### Governing requirement

Public Campaign Detail harus terasa editorial, human, structured, candid, dan optimistic tanpa menaikkan certainty dari evidence. Hierarki harus membuat identity/purpose, steward, funding meaning, provenance/uncertainty, serta truthful action mudah dipahami; mobile/desktop transformation tidak boleh menghilangkan trust-critical information atau action clarity.

Anchors:

- `docs/ui-ux/product-design-principles.md` — §1–8, §10–15.
- `docs/ui-ux/brand-product-ui-brief.md` — §3–8, §10, §13–15.
- `docs/ui-ux/patterns.md` — §3 Detail, §9–15, §18–19.
- `docs/ui-ux/design-guidelines.md` — §11–14, §16–18, §20–27.
- `docs/ui-ux/asset-governance.md` — §2–5, §8–12, §16.
- `docs/ui-ux/page-map.md` — §1 Guest/Public Visitor and §7–9.
- `docs/ui-ux/visual-references/selected-direction/README.md` — authority and interpretation rule.

### Concrete gaps

1. Tidak ada approved route-level Public Campaign Detail composition. Authority menetapkan hierarchy intent dan reusable visual/UX rules, tetapi exact section order, disclosure behavior untuk long story, media composition, dan responsive transformation masih harus dibuktikan pada real content/states.
2. Required generic states sudah punya pattern, tetapi campaign-specific presentation/copy untuk loading, missing media, not-public/not-found, request failure, unavailable funding/provenance, dan truthful pre-Slice-2 next-action area belum direkonsiliasi.
3. Exact campaign-placeholder implementation masih `OPEN`; tidak ada canonical production placeholder asset yang dapat dianggap siap pakai.
4. Provenance grammar ada, tetapi final production terminology sengaja `OPEN`. Conceptual labels seperti `Dari pengelola` tidak boleh dianggap final contract semantics.
5. Accessibility requirements ada pada authority level, namun contrast, keyboard/focus behavior, image alternative text, readable long content, and responsive DOM/reading order hanya dapat dibuktikan oleh implementation/rendered evidence.
6. Design readiness terlihat **PARTIAL**: goals, priority, truth grammar, core visual system, and state classes sudah established; route-specific composition dan beberapa asset/terminology details belum resolved. Tidak ditemukan bukti bahwa high-fidelity mockup wajib sebelum implementation.

### Sniffing

- **Risk:** Menyalin selected public reference secara literal akan memasukkan illustrative donor counts, impact language, campaign-discovery breadth, testimonial claims, dan active donation CTAs yang tidak dibuktikan oleh Slice-1 product/contract state.
- **Edge cases:** Long organizer story; absent/failed media; very large currency values; zero/unknown target or collected amount; missing steward subfields; source label yang panjang; narrow viewport; image crop yang menghilangkan dignified context; reduced motion/high zoom.
- **Miscontext:** `sunlit-editorial-public.png` membuktikan direction/relationship, bukan Public Campaign Detail route contract. `sunlit-editorial-evidence-journal.png` membuktikan cross-surface design viability, bukan availability of post-donation data pada Slice 1.
- **Misleading signals:** Yellow, green, checkmarks, badges, photos, progress bars, dan polished cards dapat tampak seperti trust/verification/success evidence padahal authority secara eksplisit melarang visual implication tanpa product truth.
- **Inconsistency:** `brand-product-ui-brief.md` masih menyebut exact typefaces/colors intentionally open pada tahap brief, sedangkan downstream canonical `design-guidelines.md` sudah menutup keputusan tersebut dengan Newsreader/Instrument Sans dan exact core palette. Ini bukan unresolved conflict; concern yang lebih konkret dimiliki `design-guidelines.md`.

### Carry-forward anchors

- `docs/ui-ux/product-design-principles.md` — design-readiness and agent decision boundary.
- `docs/ui-ux/patterns.md` — Detail, loading/error, status, money, progressive disclosure, responsive contracts.
- `docs/ui-ux/design-guidelines.md` — exact current visual system and truth/progress grammar.
- `docs/ui-ux/asset-governance.md` — campaign media/placeholder truthfulness and approval lifecycle.
- `docs/ui-ux/visual-references/selected-direction/sunlit-editorial-public.png` — relationship/character evidence only.
- `docs/ui-ux/visual-references/selected-direction/README.md` — prevents screenshot content from becoming product truth.

## Area 3 — Existing Delivery / Domain Specifications

### Current state

- Relevant Organization and Campaign specs are predominantly `draft`, dated 2026-08-20, and authored against the historical OpenAPI. Their headers/references still use pre-renumbering paths such as `docs/spec/campaign/...` and several call OpenAPI the “ground truth”; current root/spec routing has superseded that precedence.
- `docs/spec/4-campaign/features/02-campaign-detail-listing.md` already describes a useful composite read concept: Campaign + minimal Organization summary + funding progress. Historical behavior is public only for raw `status = published`; other states return relationship-dependent data or `403`.
- The same spec exposes `OrganizationSummary.status`, `CampaignProgress.donor_count`, and a response inheriting all `Campaign` fields. Those fields/shape are not justified automatically by current Slice-1 authority.
- `INV-campaign-14`, the Campaign threat model, and media feature spec align list visibility with parent Campaign visibility, which is a useful invariant concept. However they define non-public probing as `403` and mark existence disclosure as accepted, while approved public-safe direction calls for not-found-equivalent behavior.
- Campaign media spec assumes direct URLs in a publicly readable bucket. Endpoint list gating exists, but a previously learned object URL remains anonymously reachable after unpublish/retraction. The historical threat model incorrectly records no residual information-disclosure risk after endpoint gating.
- Historical Campaign task ordering requires Campaign creation before detail/media and includes full curation/publish/self-service operations. Current Slice 1 explicitly permits seeded/operator-assisted Organization/Campaign/media state, so that dependency order is not current delivery authority.
- Organization detail spec is authenticated and broadly returns `description`, `contact`, raw `status`, `has_overdue_report`, timestamps, and relationship-derived fields, with special gating for `npwp`. It is not a public steward projection and cannot be embedded or exposed wholesale for Slice 1.
- `docs/spec/0-foundations/features/01-frontend-experience-foundation.md` is `delivered`. It establishes the new frontend generation, design authority precedence, rendered desktop/mobile validation, and the rule not to invent Campaign semantics. It is relevant reusable project evidence, not Slice-1 business behavior.
- Publish/unpublish/closure specs contain future lifecycle and correctness evidence, but self-service publish/curation and closure are outside Slice 1 unless exploration proves a narrower enabling need. Retraction/revocation semantics remain relevant to the public-media safety boundary even when operator-assisted.

### Governing requirement

Active delivery specs must be reconciled downward from Product/MVP + Design authority. Slice 1 needs only domain behavior, invariants, and threats necessary to render a persisted eligible campaign truthfully and safely; historical domain breadth must not become scope because it is already documented.

Anchors:

- `docs/spec/README.md` — §1 Authority relationship and §2 document types.
- `docs/spec/4-campaign/features/02-campaign-detail-listing.md` — Detail behavior and response composition.
- `docs/spec/4-campaign/invariants.md` — INV-campaign-14.
- `docs/spec/4-campaign/threat-model.md` — Public detail and Campaign media.
- `docs/spec/4-campaign/features/03-campaign-media.md` — list visibility and public-bucket assumption.
- `docs/spec/3-organization/features/02-organization-detail-view.md` and `invariants.md` INV-organization-10.
- `docs/spec/0-foundations/features/01-frontend-experience-foundation.md` — delivered foundation evidence.

### Concrete gaps

1. Belum ada reconciled Slice-1 feature spec/acceptance criteria yang menggabungkan explicit `PublicCampaignDetail`, public Organization projection, funding semantics, media/missing-media state, backend-authoritative action eligibility, public-safe non-visibility behavior, dan required UI states.
2. `INV-campaign-14` bertentangan dengan approved public direction pada dua titik: raw-status-based `published`-only visibility dan existence-confirming `403` untuk public caller. Approved future closed continuity juga belum direkonsiliasi ke invariant, tetapi delivery-nya tetap bukan Slice 1.
3. Campaign detail spec mengandalkan full internal `Campaign` field set dan raw `OrganizationSummary.status`; current authority requires explicit allowlists and deliberately designed public meaning.
4. Campaign media spec/threat model tidak memenuhi revocation requirement karena direct public object URL tetap hidup setelah endpoint access dicabut.
5. Historical `CampaignProgress` memasukkan `donor_count` dan percentage capped at 100. Product need hanya menetapkan target/current funding facts yang relevan; donor count berisiko menjadi unsupported social proof dan capping dapat menyembunyikan over-target truth bila kondisi itu mungkin.
6. Organization detail specs tidak menyediakan anonymous public Organization projection. Reusing the authenticated detail would overexpose contact/internal status/overdue-report/relationship semantics.
7. Threat model belum mencakup explicit public-projection regression risk, internal-field addition leakage, direct-object revocation, hostile organizer content rendering, cache/stale visibility after retraction, atau public response variance caused by optional Authorization headers.
8. Draft task documents masih menganggap historical domain order sebagai dependency order dan menyebut OpenAPI sebagai ground truth, sehingga berisiko men-steer planning keluar dari current Product-first precedence.

### Sniffing

- **Risk:** Full-schema inheritance membuat future internal field otomatis public; direct bucket URLs menggagalkan legal/privacy/safety withdrawal; authenticated Organization detail reuse membocorkan internal/PII-adjacent context; stale task ordering dapat menarik full self-service/curation scope ke Slice 1.
- **Edge cases:** Campaign absent vs existing-non-public; published then retracted; published media URL cached/previously shared; no media; zero/over-target funding; Organization loses eligibility; optional Authorization header on otherwise public request; organizer story containing hostile markup.
- **Miscontext:** Historical task dependency “creation before detail” mengasumsikan self-service operational workflow. Slice 1 explicitly allows seeded/operator-assisted setup.
- **Misleading signals:** Media endpoint marked gated and threat “resolved” terlihat aman, tetapi direct public object access remains revocable only in documentation, not in the described architecture. `OrganizationSummary` bernama “minimal” tetapi raw status can still imply a public verification claim.
- **Inconsistency:** Feature/invariant says Campaign attachments follow detail gating, while current OpenAPI still declares list `security: []`; specs accept `403` disclosure while approved probe direction says public-safe not-found; old references call OpenAPI ground truth while current authority map places reconciled specs above contract.

### Carry-forward anchors

- `docs/spec/4-campaign/features/02-campaign-detail-listing.md` — historical composite-read behavior and unsafe/full response assumptions.
- `docs/spec/4-campaign/invariants.md` — INV-campaign-14 conflict coordinates.
- `docs/spec/4-campaign/threat-model.md` — currently accepted existence disclosure and incomplete media analysis.
- `docs/spec/4-campaign/features/03-campaign-media.md` — direct public URL/public-bucket assumption.
- `docs/spec/3-organization/features/02-organization-detail-view.md` — evidence that internal/authenticated Organization detail is not the public steward contract.
- `docs/spec/0-foundations/features/01-frontend-experience-foundation.md` — delivered frontend baseline and rendered-acceptance precedent.

## Area 4 — Existing Shared API Contract

### Current state

- Split contract sources expose `GET /campaigns/{campaignId}` as the relevant composite read, `GET /campaigns/{campaignId}/attachments` as a separate media list, and `GET /organizations/{organizationId}` as an authenticated full Organization read.
- `GET /campaigns/{campaignId}` is explicitly public only for `published`; the same operation also describes richer access for representatives/Kurator/Admin. Its `200` response is `CampaignDetail`.
- `CampaignDetail` uses `allOf` with the full internal-looking `Campaign` schema. That schema includes `organization_id`, raw `CampaignStatus`, `max_amount`, `publish_at`, `unpublish_reason`, `closed_reason`, `decision_note`, `created_by`, and general audit timestamps in addition to public content/funding fields.
- Embedded `OrganizationSummary` contains `id`, `name`, and raw `OrganizationStatus`. `CampaignProgress` contains a float percentage capped at 100, `donor_count`, and `days_remaining`. Target and collected amounts are decimal strings, which preserves transport precision.
- No `PublicCampaignDetail`, `PublicOrganization`, `PublicCampaignMedia`, backend-authoritative donation/action eligibility field, provenance/source field, or explicit accountability availability state exists.
- Attachment list is unconditionally `security: []`, has only a `200` response, and returns raw `CampaignAttachment` metadata including direct public `url`, `original_name`, `size_bytes`, `uploaded_by`, and `created_at`.
- `GET /campaigns/{campaignId}` advertises distinct `403` for an existing non-public Campaign. The common `NotFound` response exists and follows Problem Details, but current Campaign contract does not use it to conceal non-public existence.
- Most Campaign/Organization response schemas do not declare `required` properties, so generated clients cannot rely on fields being present even where prose implies them.
- Contract lint was executed from `api/` with `npm run validate`: **failed** with 1 error and 130 warnings. Blocking error: root `security` references `bearerAuth`, but `api/openapi/index.yaml` does not expose the security scheme defined in `common.yaml`. Relevant warnings include missing `operationId`, `$ref` siblings being ignored, and no `4XX` response on the media-list operation. This is repository-wide pre-existing contract health evidence, not a Slice-1-only failure.

### Governing requirement

The reconciled Slice-1 contract must expose only the public-safe facts required by the surface, preserve backend authority over visibility/action eligibility, align media visibility with parent Campaign lifecycle and real revocability, use public-safe not-found behavior, and provide stable enough schemas for frontend behavior without inheriting internal records.

Anchors:

- `api/openapi/campaign.yaml` — `/campaigns/{campaignId}`, `/campaigns/{campaignId}/attachments`, `Campaign`, `OrganizationSummary`, `CampaignProgress`, `CampaignDetail`, `CampaignAttachment`.
- `api/openapi/organization.yaml` — `/organizations/{organizationId}` and `Organization`; internal/authenticated evidence only.
- `api/openapi/common.yaml` — shared Problem Details and `NotFound` response.
- `api/openapi/index.yaml` — aggregate path registration and unresolved security-scheme wiring.
- `api/README.md` — split-source workflow and known `$ref` sibling issue.

### Concrete gaps

1. Public Campaign Detail inherits the full `Campaign` schema rather than an explicit allowlist. This creates present leakage (`decision_note`, `created_by`, raw internal statuses/reasons) and future leakage whenever internal fields are added.
2. One path mixes unauthenticated public and privileged operational detail semantics. The response contract cannot guarantee that optional Authorization never changes a public payload into a richer internal one.
3. There is no narrow public Organization/steward projection with deliberately public semantics. Raw `OrganizationStatus` risks being presented as broad verification/trust meaning.
4. No backend-authoritative `can_donate`/public action capability exists. Reconstructing eligibility from raw status/deadline/funding in React would duplicate business rules.
5. Contract does not explicitly distinguish organizer-provided story from platform facts, and it does not encode unknown/pending information where absence would be ambiguous.
6. Media contract exposes storage/uploader metadata unnecessary for the public surface, has no alt/caption/provenance fields, does not express parent visibility/not-found errors, and assumes permanently public direct URLs.
7. Funding contract includes unapproved social-proof data (`donor_count`) and a capped percentage that can hide actual over-target relationship. Exact Slice-1 funding projection is unreconciled.
8. Weak schema requiredness and currently failing split-source validation make generated-client correspondence unreliable until contract health is addressed within the narrow active-slice change.

### Sniffing

- **Risk:** Public data leakage expands silently through `allOf`; optional auth can cause cache/user-dependent response variance; immutable public media URLs defeat withdrawal; optional generated fields encourage frontend fallbacks/inference; contract lint failure can hide broken generation/bundling.
- **Edge cases:** Non-public vs absent UUID; Authorization header present on public request; campaign transitions/retraction while cached; empty attachment array vs media unavailable; `target_amount = 0` or collected above target; missing required properties; amount strings larger than JS safe integer.
- **Miscontext:** `CampaignDetail` is described as the platform’s primary “conversion surface”, but Slice 1 explicitly prioritizes understanding and cannot expose an active conversion action yet.
- **Misleading signals:** `OrganizationSummary` is labeled minimal and `CampaignAttachment` is called public, but minimality/public naming does not prove safe product semantics or revocability. A successful bundled historical file does not negate current split-source lint failure.
- **Inconsistency:** OpenAPI media list remains unconditional public while feature spec requires parent gating; OpenAPI mixes `security: []` public detail with privileged response behavior; `api/README.md` lists `bearerAuth` in shared components but the root source does not reference/import it, producing the actual lint error.

### Carry-forward anchors

- `api/openapi/campaign.yaml` `/campaigns/{campaignId}` — current composite path and unsafe response coupling.
- `api/openapi/campaign.yaml` `CampaignDetail` + `Campaign` — public inheritance leak surface.
- `api/openapi/campaign.yaml` `OrganizationSummary` / `CampaignProgress` — public meaning and unnecessary-field reconciliation points.
- `api/openapi/campaign.yaml` `/campaigns/{campaignId}/attachments` + `CampaignAttachment` — visibility, projection, and revocation gaps.
- `api/openapi/common.yaml` `NotFound` / `Problem` — reusable public-safe error contract evidence.
- Executed check: `cd api && npm run validate` — failed, 1 error / 130 warnings on target revision.

## Area 5 — Backend Current State

### Current state

- Live backend has exactly one implemented domain package: `internal/domain/account`. There are no `organization`, `campaign`, or `donation` domain packages.
- Migrations `000001`–`000010` create only Account/auth/session/security tables. No Organization, Campaign, Campaign media, funding, or seed-data persistence exists.
- `cmd/server/main.go` wires health, Swagger/OpenAPI, Account auth/security/profile routes only. No `/campaigns`, `/organizations`, or media route is registered.
- MinIO initialization verifies that configured public/private buckets exist, but the client is not retained/injected into a storage service. `internal/platform/storage` is an empty package skeleton.
- `internal/platform/scheduler` is also an empty skeleton; no Campaign scheduled publication/deadline behavior exists.
- The dev server exposes `api/openapi.yaml` and Swagger for endpoints that are not implemented. Contract documentation therefore looks substantially more complete than live backend behavior.
- Reusable foundations do exist: pgx pool setup, established parameterized `goqu` repository style, RFC 9457 `WriteProblem`, non-leaking unknown-error mapping, per-IP rate-limit middleware, MinIO dependency/config bootstrap, and a domain-driven monolith layout convention.
- The backend architecture document still records a public bucket for Campaign media and calls the old aggregate OpenAPI the single source of truth. Current Product/media safety and root API routing supersede those historical details where they conflict.
- Executed `cd backend && go test ./...`: **PASS**. Coverage includes current Account/platform/HTTP packages; storage and scheduler have no tests, and there are no Campaign/Organization packages to test. No real-Postgres integration test was run in this Exploration.

### Governing requirement

Slice 1 needs real persisted Organization/Campaign/media state and an anonymous public-safe read capability whose visibility, funding facts, and action eligibility are backend-authoritative. Operational creation may be seeded/operator-assisted, but the user-visible response cannot be hard-coded or frontend-only.

Anchors:

- `backend/cmd/server/main.go` — `run`, current dependency wiring and route registration.
- `backend/migrations/` — persistence inventory ending at Account migration `000010`.
- `backend/internal/platform/storage/doc.go` — empty future storage skeleton.
- `backend/internal/platform/scheduler/doc.go` — empty future scheduler skeleton.
- `backend/internal/transport/http/errors.go` — reusable safe Problem Details behavior.
- `backend/internal/transport/http/middleware.go` — current per-IP rate-limit behavior and proxy caveat.
- `backend/internal/domain/account/repository_db.go` — established `goqu` + pgx prepared-query pattern evidence.
- `docs/project/kencleng-backend-tech-stack.md` — architecture direction; historical public-bucket/aggregate-contract details require reconciliation.

### Concrete gaps

1. Entire Slice-1 backend capability is absent: schema/migrations, persisted Organization/Campaign/media data, repository/service/handler, router wiring, and tests.
2. No minimum operator/seed mechanism exists to create the genuine persisted eligible Campaign required by Slice 1. Existing Account capability is not an enabling dependency under current MVP scope.
3. No public visibility rule, explicit projection, anti-enumerating not-found behavior, funding calculation/read model, or backend-authoritative next-action eligibility exists in executable code.
4. Object-storage bootstrap exists, but there is no storage abstraction or access path capable of enforcing Campaign-parent visibility and later revocation. The configured “public” bucket direction conflicts with the required ability to withdraw previously public media.
5. Served Swagger exposes historical Campaign/Organization operations despite zero backend implementation, creating false capability evidence during manual exploration/integration.
6. There is no Campaign-focused threat/contract test evidence, no integration evidence against Postgres/MinIO, and no test fixture proving missing-media, not-public, failure, or funding states.
7. Current rate limiter is reusable but explicitly notes incorrect client attribution behind a reverse proxy unless proxy address handling is added. Whether it is required for this read capability remains a delivery/security decision, not an assumed Slice-1 task.

### Sniffing

- **Risk:** Largest risk is false completeness: OpenAPI/Swagger and detailed specs exist while no live route/table does. Public media architecture can create irreversible leakage. A broad implementation attempt could also pull full Organization/Campaign lifecycle into Slice 1 unnecessarily.
- **Edge cases:** Seed rerun/idempotency; Campaign references absent Organization; invalid monetary strings/zero target; media object missing while DB row exists; lifecycle changes between read and media access; stale/cached public responses after retraction; dependency failure from Postgres/MinIO.
- **Miscontext:** Presence of MinIO bucket checks, storage/scheduler packages, and OpenAPI routes may imply capabilities are scaffolded. They are only bootstrapping/placeholders, with no Campaign behavior.
- **Misleading signals:** `go test ./...` is green, but it proves only currently implemented Account/platform behavior; it provides no evidence for Slice 1. Swagger successfully serving a contract does not mean handlers exist.
- **Inconsistency:** Architecture says public bucket is appropriate for Campaign media, while current approved safety requires actual revocation. Backend serves the bundled aggregate even though split-source validation currently fails and current root routing treats the bundle as generated rather than authoring authority.

### Carry-forward code anchors

- `backend/cmd/server/main.go:run` — new domain/storage dependencies and route wiring would ultimately meet here; currently Account-only.
- `backend/internal/platform/storage/doc.go` — explicit empty storage boundary for Campaign media.
- `backend/internal/platform/scheduler/doc.go` — explicit empty lifecycle scheduler boundary; likely not Slice-1 delivery unless a narrower enabling need is proven.
- `backend/internal/transport/http/WriteProblem` — established safe Problem Details writer.
- `backend/internal/transport/http/RateLimit` — reusable anonymous-endpoint middleware with documented reverse-proxy limitation.
- `backend/internal/domain/account/repository_db.go` — established parameterized persistence convention, not domain code to copy blindly.
- Executed check: `cd backend && go test ./...` — PASS; no Campaign/Organization coverage exists.

## Area 6 — Frontend Current State

### Current state

- Frontend reboot is complete and the Frontend Experience Foundation was delivered in PR #24 / `71093b6`, independently reviewed/tested, and received human rendered acceptance PASS.
- Live `/` route is a static, text-led platform/trust calibration surface. It establishes real production precedent for Newsreader/Instrument Sans roles, warm-neutral/Sun/Berry variables, editorial hierarchy, border/spacing-led composition, semantic landmarks, focus-visible styling, responsive recomposition, and reduced-motion handling.
- The route intentionally contains no Campaign, Organization, media, funding, donation, verification, or API data. Its same-page CTA only explains the platform’s clarity posture.
- No Campaign route exists. There are no `lib/api`, `lib/hooks`, `mocks`, `public`, or `tests/browser` implementations; no generated OpenAPI TypeScript file; no MSW handlers; no Campaign feature components; and no production `ui`/`shared` component contracts.
- `next.config.ts` is empty, so no remote campaign-image policy/loader configuration exists.
- The only automated frontend behavior test covers `/` landmarks and its fragment link. No loading, media, money, source/provenance, not-found, failure, or responsive Campaign behavior is covered.
- Current architecture supports Server Components, focused typed API functions, TanStack Query when client-owned server state is needed, and contract-faithful MSW at the network boundary. It explicitly forbids production mock-mode branches and frontend recreation of eligibility/business rules.
- `docs/project/kencleng-integration-map.md` has no active mappings. It is intentionally waiting for the first real cross-stack surface.
- Executed `cd frontend && npm run verify`: **PASS** (ESLint + 1 Vitest file / 1 test), with an existing Vite future-config warning.
- Executed `npm run build`: sandbox run failed because Google Fonts were unreachable; approved network-enabled rerun **PASS**, producing only static `/` and `/_not-found` routes. No fresh rendered/browser inspection was performed in EXP-001; prior foundation human acceptance remains historical reusable evidence.

### Governing requirement

Frontend must render the reconciled public contract truthfully across loading, real-media, missing-media, success, not-public/not-found, recoverable failure, and representative responsive states. It must preserve provenance and funding meaning, safely render organizer-controlled content, and must not show an active Donate control before the real Slice-2 flow exists.

Anchors:

- `frontend/app/page.tsx` — delivered static public calibration surface.
- `frontend/app/layout.tsx` — approved font loading and Indonesian root metadata/language.
- `frontend/app/globals.css` — current consumed visual foundation and responsive/focus rules.
- `frontend/app/page.test.tsx` — sole live frontend behavior test.
- `frontend/components/README.md` — empty new-generation reusable-contract registry.
- `docs/project/kencleng-frontend-tech-stack.md` — data/API/Server-Client/testing/rendered-verification boundaries.
- `docs/project/kencleng-integration-map.md` — empty current mapping and network-boundary MSW rule.
- `frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/validation-closeout.md` — prior foundation delivery/human acceptance evidence.

### Concrete gaps

1. Entire Public Campaign Detail surface is absent: route, data boundary, generated contract types, state handling, domain presentation, and verification.
2. Current OpenAPI is not stable/safe enough to generate the intended public shape, and the frontend has no generated type artifact. Handwriting a temporary Campaign model would violate architecture.
3. No contract-faithful mocks exist for parallel frontend work, and no integration-map row records the selected surface/capability/backend owner/contract gap.
4. There is no safe organizer-content rendering decision. `Campaign.description` is just a string; whether it is plain text or supported markup is not reconciled. Unsafe HTML rendering is prohibited and any markup path requires hostile-content coverage.
5. No production Campaign image/placeholder asset exists, and no Next image-host/access policy is configured. The exact placeholder system remains a design `OPEN` item.
6. No amount/progress formatter or semantics exist. Frontend must treat decimal amount strings without lossy numeric assumptions and must not own financial/eligibility calculations.
7. Required state-specific tests and rendered evidence are absent. Existing foundation test/build/human acceptance cannot prove Campaign state behavior.
8. The existing root surface offers a healthy visual precedent, but it has no reusable component contracts. Extracting a broad design system before Campaign usage would be speculative; conversely, duplicating its raw CSS without considering real Campaign state/content would fail to establish necessary feature behavior.

### Sniffing

- **Risk:** Frontend can easily invent eligibility/status/provenance semantics while backend is absent; raw story rendering can introduce XSS; arbitrary image configuration can bypass media revocation; converting decimal strings through JS number can lose money precision; fake CTA/mocks can appear production-real.
- **Edge cases:** Very long title/story/Organization name; zero/unknown/over-target funding; broken/slow/revoked image; no media; missing optional public fields; 404 indistinguishable from non-public; request retry; mobile/high zoom; screen-reader source/context order.
- **Miscontext:** Dependencies such as TanStack Query, MSW, Zustand, and Playwright are capabilities, not required architecture. The static `/` page is a foundation precedent, not an existing Campaign surface to extend blindly.
- **Misleading signals:** Current build/verify PASS and polished `/` can suggest frontend readiness equals feature readiness. It proves scaffold/foundation health only; all Campaign data/state/integration behavior remains absent.
- **Inconsistency:** Frontend README says OpenAPI-generated types are a selected capability, but no generated artifact/script exists yet. This is intentional clean-start posture, not a defect by itself; it becomes an active gap for Slice 1.

### Carry-forward code anchors

- `frontend/app/page.tsx:Home` — current public composition/copy precedent and existing `/` ownership.
- `frontend/app/layout.tsx:RootLayout` — font, language, and metadata boundary.
- `frontend/app/globals.css` — current concrete consumed tokens/layout/focus/responsive rules.
- `frontend/app/page.test.tsx` — observable-testing precedent and current narrow coverage.
- `frontend/components/README.md` §5 Current registry — confirms no broad production contracts yet.
- `frontend/next.config.ts` — currently empty image/runtime configuration.
- `docs/project/kencleng-integration-map.md` §6 — confirms first active mapping has not been recorded.
- Executed checks: `npm run verify` PASS; network-enabled `npm run build` PASS; no fresh browser/rendered check in EXP-001.

## Area 7 — Cross-Stack Readiness and Completion Evidence

### Current state

- Development tracker records Organization backend/frontend as `NOT_STARTED`, Campaign as `NEEDS_RECONCILIATION` / `NOT_STARTED`, and Slice 1 as the next selected work. Frontend foundation is `DELIVERED`; backend Account implementation is unrelated historical evidence.
- Integration map intentionally contains no active mapping, so Public Campaign Detail is not yet connected to a product capability, reconciled operation, backend owner, or explicit contract gap in the project rendezvous layer.
- No milestone is earned: contract is not `CONTRACT_READY`, backend is not `BACKEND_VERIFIED`, frontend is not `FRONTEND_MOCK_VERIFIED`, and real integration is absent.
- Local topology can run PostgreSQL, MinIO, Caddy, backend, and frontend, but two executable configuration facts block truthful completion:
  - Caddy forwards `/api/*` without stripping `/api`, while backend routes are registered without that prefix. Intended same-origin API calls through `localhost:8080` currently 404 unless a root-scoped proxy/backend-path reconciliation occurs.
  - `minio-init` sets anonymous download on the entire `kencleng-public` bucket, which cannot satisfy per-Campaign withdrawal/revocation after a URL is known.
- There is no integrated persisted Campaign fixture/seed, no Campaign handler, no frontend request, and no real-data browser evidence. Therefore none of the Slice-1 completion states can currently be exercised end-to-end.
- Project workflow requires real backend/frontend integration, representative real-data loading/error/empty/success behavior, preservation of public-data/security boundaries, responsive rendered checks, and human browser acceptance before `INTEGRATED_VERIFIED`/finalization.
- Development tracker contains branch-era stale wording: Product Authority is marked `PROMOTION_READY` and a PR #27 merge gate remains, while the current target branch already presents Product Authority as canonical/promoted. This does not change source precedence for EXP-001, but project-status text needs reconciliation before it is used as current delivery state.

### Governing requirement

Slice 1 is complete only when a public visitor can load a persisted eligible Campaign and correctly understand Campaign/steward/funding/source context across loading, missing-media, success, not-public/not-found, failure, and responsive states without internal leakage. Completion must be based on the real contract and real backend, not only mocks, static pages, Swagger, or independent stack checks.

Anchors:

- `docs/project/kencleng-development-tracker.md` — §§3–9 current statuses and next selection.
- `docs/project/kencleng-integration-map.md` — §§4–6 lazy mapping and contract-parallel rules.
- `docs/kencleng-agentic-workflow.md` — §§3–12 preflight, sequencing, milestones, integration, finalization.
- `Caddyfile` — current unstripped `/api/*` proxy.
- `docker-compose.yml` — anonymous `kencleng-public` bucket policy.
- `docs/project/kencleng-repo-setup.md` §8 — documented root-scoped proxy caveat.

### Concrete gaps

1. No reconciled Slice-1 delivery spec, threat model, or shared contract is ready; downstream FE/BE sequencing cannot truthfully be marked decided yet.
2. No active integration-map relationship exists for Public Campaign Detail.
3. Same-origin API topology is known-broken because of the `/api` prefix mismatch.
4. Current object-storage policy is incompatible with lifecycle-driven public media revocation.
5. No genuine persisted Campaign/Organization/media fixture exists for integration or human acceptance.
6. No end-to-end evidence exists for required loading, missing-media, success, not-public/not-found, recoverable failure, or responsive states.
7. Existing green checks are isolated: API lint fails, backend unit tests pass unrelated code, frontend verify/build pass only the static foundation. None combine into Slice-1 completion evidence.
8. Project tracker is not fully current relative to the promoted authority state visible on the target branch.

### Sniffing

- **Risk:** Independent FE/BE green states can be misreported as integrated completion; proxy mismatch can be hidden by direct-backend testing; public bucket configuration can make UI retraction look successful while media remains reachable; stale tracker status can misroute authority/workflow decisions.
- **Edge cases:** Direct backend URL works while proxy URL fails; media list becomes private while cached/direct object stays public; frontend mock differs from final contract requiredness/errors; seeded record is not public-eligible; backend unavailable after static shell loads; real content breaks mobile hierarchy.
- **Miscontext:** A cross-stack commit, Swagger operation, MSW demo, or static rendered Campaign page would not establish `INTEGRATED_VERIFIED`. `FRONTEND_MOCK_VERIFIED` and `BACKEND_VERIFIED` are explicitly narrower milestones.
- **Misleading signals:** Infrastructure containers and buckets exist but no Campaign behavior uses them; the bundled OpenAPI exposes routes but backend does not; current `/` has human acceptance but Public Campaign Detail does not.
- **Inconsistency:** Intended API base `/api` conflicts with backend route paths and Caddy behavior; “public bucket” configuration conflicts with revocability; tracker promotion status conflicts with current branch authority headers.

### Carry-forward anchors

- `docs/project/kencleng-integration-map.md` §6 — first active surface mapping remains absent.
- `docs/project/kencleng-development-tracker.md` §§4, 7, 9 — current domain/slice state plus stale promotion gate.
- `Caddyfile` `handle /api/*` — same-origin integration blocker.
- `docker-compose.yml` `minio-init` — anonymous public-bucket policy and revocation blocker.
- `docs/kencleng-agentic-workflow.md` §§9–12 — evidence milestone and finalization definitions.

## Stage 2 synthesis — findings only

### Material findings

1. **Authority is sufficient to define the Slice-1 outcome, but no reconciled delivery contract exists.** Exact public fields, source semantics, action capability, media projection, and public-safe error behavior remain below Product Authority and need narrow reconciliation.
2. **Historical `CampaignDetail` is unsafe as a public contract.** It inherits internal Campaign fields, mixes public/privileged semantics, exposes raw Organization status and unnecessary progress/social-proof data, and lacks backend-authoritative action eligibility.
3. **Media revocation is a cross-layer security gap.** Spec, architecture, OpenAPI, Docker MinIO policy, and empty backend storage implementation collectively cannot make retracted media actually private after URL disclosure.
4. **Implementation is mostly absent, not partially complete.** Backend has no Campaign/Organization persistence or routes; frontend has no Campaign surface/data boundary; the only reusable implementation is the delivered public visual foundation and generic backend infrastructure patterns.
5. **Integration topology is not ready.** Current Caddy prefix behavior breaks intended same-origin API calls, and no integration mapping/fixture/evidence exists.
6. **Required experience states are well named but unimplemented.** Loading, real/missing media, success, not-public/not-found, recoverable failure, and responsive behavior have neither contract-level detail nor executable proof.
7. **Story/content safety is unresolved.** Contract exposes a string but does not define plain text versus supported markup; frontend has no safe renderer, and hostile-content coverage is mandatory if markup is allowed.
8. **Project evidence contains stale/misleading signals.** Swagger advertises unimplemented endpoints; old specs call OpenAPI ground truth; the tracker retains pre-promotion gate wording; isolated green tests do not prove Slice 1.

### Authority questions versus delivery questions

- No contradiction requiring immediate Product Authority revision was found. Current Product/MVP sequencing is sufficient to keep Donation Flow in Slice 2 and closure/result delivery in Slice 3 while preserving future compatibility.
- Material unresolved items currently appear to be delivery/design/contract questions: exact allowlisted facts, public steward meaning, pre-Slice-2 next-action presentation, content format, placeholder implementation, public media access mechanism, response/error semantics, and FE/BE sequencing.
- If Stage 3 concludes that Slice 1 must publicly communicate an Organization review/verification claim whose meaning cannot be derived from current Product Authority, that specific item must be escalated rather than invented downstream.

### Verification performed during Stage 2

| Check | Result | Meaning |
|---|---|---|
| `cd api && npm run validate` | Failed: 1 error, 130 warnings | Split OpenAPI source is not currently validation-clean; not Slice-1 readiness. |
| `cd backend && go test ./...` | Passed | Current Account/platform/HTTP unit packages are green; no Campaign/Organization evidence. |
| `cd frontend && npm run verify` | Passed: lint + 1 test | Static frontend foundation remains mechanically healthy. |
| `cd frontend && npm run build` | Passed on network-enabled rerun | Current `/` and `/_not-found` compile/prerender; first sandbox run failed only on Google Fonts network access. |
| `git diff --check` | Passed | Stage-2 evidence patch has no whitespace errors at time checked. |

### Not tested / deferred in Stage 2

- No Postgres/MinIO integration tests or migrations were run.
- No local infrastructure was started or mutated.
- No backend/frontend server was started and no real API/browser integration was exercised.
- No fresh desktop/mobile rendering was inspected; prior foundation acceptance is historical evidence only.
- No protected Tier-0 path was modified.
- No contract, product/design authority, spec, application code, proxy, or storage configuration was changed.
