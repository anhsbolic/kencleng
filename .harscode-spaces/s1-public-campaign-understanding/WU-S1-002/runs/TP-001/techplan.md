# Tech Plan: Slice 1 Public Contract & Delivery Reconciliation

> Phase             : Techplan
> Ticket            : WU-S1-002
> Author            : Planner
> Model             : `gpt-5.6-sol`
> Reasoning         : `high`
> Created           : 2026-09-21
> Target revision   : `e31e23b60ef2a8f994db623f8943a4eac3595e7e`
> Workflow revision : `396b9ba664aaab9cb959786a97fd2594346c700e`
> Status            : Approved by @anhsbolic
> Approach          : Reconcile the smallest explicit public Campaign contract and its owning delivery/security records before any backend or frontend production Build.
> Refs              : `WU-S1-001 / EXP-001` gap analysis and solutioning; `docs/product/product-overview.md`; `docs/product/mvp-scope.md`; `docs/product/mvp-delivery-slices.md`; applicable `docs/ui-ux/`; current Campaign specs and split OpenAPI sources

> Amendment — `TPR-RES-001` (2026-09-21): The reviewed exact-allowlist intent is made executable by requiring closed-object semantics; this does not add, remove, or reinterpret a public field or public behavior.

---

## 1. Background

Slice 1 harus memungkinkan public visitor memahami satu persisted eligible Campaign secara aman dan jujur sebelum Donation Flow tersedia. Product/MVP dan Design Authority sudah cukup jelas, tetapi delivery artifacts saat ini masih membawa historical contract: `GET /campaigns/{campaignId}` mencampur public dan privileged detail, `CampaignDetail` mewarisi internal `Campaign`, non-public Campaign menghasilkan existence-confirming `403`, media memakai direct public-bucket URL, dan requiredness schema lemah.

Live contract validation pada target revision juga masih gagal dengan satu error karena root `bearerAuth` tidak terhubung ke scheme di `common.yaml`; 130 warning lain adalah backlog historis. Tidak ada live Campaign/Organization backend route atau frontend Campaign consumer yang perlu dipertahankan. Reconciliation ini menghasilkan authoritative Slice-1 delivery/spec/threat/API spine dan generated TypeScript types yang cukup stabil untuk contract-parallel backend/frontend planning. Reconciliation ini tidak mengimplementasikan runtime capability dan tidak dengan sendirinya membuktikan `BACKEND_VERIFIED`, `FRONTEND_MOCK_VERIFIED`, atau `INTEGRATED_VERIFIED`.

## 2. Scope

**In scope:**

- Reconcile Campaign Task 02/03 delivery records untuk active Slice 1, termasuk acceptance criteria dan klasifikasi historical listing/upload/self-service breadth.
- Replace `INV-campaign-14` dengan durable public-eligibility boundary yang memakai public-safe `404`, public-only projection, dan parent-governed media visibility.
- Reconcile Campaign threat model untuk allowlist regression, optional-auth variance, enumeration, plain-text organizer content, media revocation, cache staleness, dan dependency failure.
- Define exact `GET /campaigns/{campaignId}` public contract and a controlled `GET /campaigns/{campaignId}/media/{mediaId}/content` byte-delivery contract.
- Remove the historical anonymous attachment-list operation from the active public contract; keep unreconciled upload/operational behavior explicitly deferred.
- Fix the root `bearerAuth` reference error, regenerate the aggregate OpenAPI bundle, and establish one committed OpenAPI-generated TypeScript type artifact.
- Narrowly correct backend architecture text that still prescribes a public Campaign-media bucket or treats the generated bundle as the hand-authored source.
- Add the first Public Campaign Detail integration mapping and reconcile project tracker state only after contract evidence passes.
- Preserve future extension to Slice 2 donation availability and Slice 3 closed-result continuity without exposing those behaviors now.

**Out of scope (explicit):**

- Backend migrations, persistence, seed command, repositories, services, handlers, storage implementation, route wiring, or runtime tests.
- Frontend Campaign route, components, fetch functions, MSW handlers, presentation copy, assets, or rendered verification.
- `Caddyfile`, `docker-compose.yml`, MinIO policy mutation, or any topology implementation; these remain a downstream explicitly owned Work Unit/batch.
- Campaign/Organization creation, broad listing/discovery, full authenticated detail, upload, curation, publish UI/API implementation, scheduler, closure/result/accountability, Donation Flow, or Account work.
- Repository-wide OpenAPI cleanup; historical warnings outside touched coordinates remain backlog.
- Editing Product/MVP or Product Design Authority.
- Editing protected Tier-0 implementation paths.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | Public detail exposes only deliberately public Campaign, steward, lifecycle, funding, provenance, media, and action semantics needed by Slice 1. | `docs/product/mvp-scope.md#stage-a--public-understanding`; `docs/product/mvp-delivery-slices.md#slice-1--public-campaign-understanding` |
| Q2 | Public contract is an explicit allowlist and never inherits internal Campaign or authenticated Organization records. | Product Probe 01; `EXP-001/evidence/solutioning.md#decision-2--make-the-existing-public-detail-path-public-only`; Decision 3 |
| Q3 | Absent and non-public Campaigns are indistinguishable at the public boundary; optional Authorization cannot enrich or otherwise vary the payload. | Probe 01 §6; Exploration Decisions 2–3; `best-practices/restapi/anti-enumeration.md` |
| Q4 | Organizer purpose/story is plain text with machine-readable organizer provenance; final Indonesian label remains presentation-owned. | Exploration Decision 4; UI/UX truth grammar; `best-practices/pwa/xss-and-content-sanitization.md` |
| Q5 | Funding uses decimal strings and backend-authored availability/progress meaning; zero is not treated as absence, over-target truth is preserved, and `donor_count` is excluded. | Root money rule; Exploration Decision 5; `best-practices/go/decimal-and-money.md` |
| Q6 | Slice 1 represents donation action as unavailable and exposes no URL/control target that could become a fake Donate flow. | `docs/product/mvp-delivery-slices.md#frontend-boundary-before-slice-2`; Exploration Decision 5 |
| Q7 | Public media metadata follows the parent public-eligibility predicate; bytes remain in private object storage and are served only through a parent/member-authorizing operation. | Product Probe 01 §7–8; Exploration Decision 6 |
| Q8 | Retraction prevents new anonymous origin fetches through a previously known media URL; system-controlled caching must not defeat withdrawal. | Exploration Decision 6 and carry-forward risks 3–4 |
| Q9 | Split OpenAPI is the authoring source, aggregate OpenAPI and frontend types are generated artifacts, and touched success/error shapes remain synchronized. | Root `AGENTS.md`; `api/README.md`; `best-practices/restapi/openapi-spec-first-drift.md` |
| Q10 | Reconciliation reaches only `CONTRACT_READY`; backend, frontend, topology, and integration milestones remain unearned. | `docs/kencleng-agentic-workflow.md` project development states; WU-S1-002 manifest |

## 4. Rules & Validation

- **R1 — Public-only operation.** Given any request to `GET /campaigns/{campaignId}`, when the Campaign satisfies the Slice-1 public-eligibility predicate, then the response is `200 PublicCampaignDetail`; the presence, absence, validity, or role content of an `Authorization` header does not change response fields or visibility. For the currently deliverable Slice-1 state, the predicate admits only the internal `published` fundraising state; draft, pending, approved-but-not-published, scheduled, rejected, unpublished/retracted, and historical closed states are non-public. Closed continuity is added only by Slice 3 reconciliation.
- **R2 — Anti-enumerating visibility.** Given an absent Campaign, a malformed/non-resolvable identifier, or a Campaign that does not satisfy public eligibility, when the public detail or media-content operation is requested, then each returns the same `PublicCampaignNotFound` (`404` Problem Details) response without confirming which condition occurred.
- **R3 — Explicit projection.** `PublicCampaignDetail` is a standalone required-field schema and contains exactly `id`, `title`, `purpose`, `story`, `steward`, `lifecycle`, `funding`, `media`, and `donation_action`. It does not use `allOf` with `Campaign` or `Organization` and cannot expose raw status, contact/legal data, internal IDs/reasons, actor IDs, audit timestamps, `donor_count`, or operational metadata. Every object schema in the public response graph that represents an exact projection MUST set `additionalProperties: false`; required fields alone are not sufficient to close the projection.
- **R4 — Plain-text provenance.** `purpose.content` and `story.content` are plain strings rendered with framework escaping downstream. Both use `PublicCampaignOrganizerText` with `source = organizer`; the API does not prescribe final Indonesian display wording and does not imply platform verification.
- **R5 — Public lifecycle.** `lifecycle.public_state` is `fundraising` for Slice 1 and includes required `published_at` and `fundraising_ends_at` date-times. It never exposes internal `CampaignStatus`, curation/scheduling states, or unpublish/closure reasons. Slice 3 may extend the public enum only through a later reconciled contract change.
- **R6 — Funding truth.** `funding` is an explicit tagged availability shape. The `available` branch requires `currency = IDR`, `target_amount`, `collected_amount`, and backend-authored `progress`; amounts and progress percentage use decimal strings, never OpenAPI `number`/`float`. Money uses the established `NUMERIC(19,2)` transport pattern `^(0|[1-9][0-9]{0,16})\.[0-9]{2}$`. A computed percentage is the exact decimal ratio rounded half-up to at most two fractional digits, serialized as a non-negative decimal string, and never capped at 100. `progress.relationship` is one of `below_target`, `target_reached`, `above_target`, or `not_computable`; percentage is explicit `null` when not computable. The `unavailable` branch has no fabricated amounts. A factual zero remains the string `"0.00"`, not the unavailable branch.
- **R7 — No fake action.** `donation_action` is required and, for Slice 1, has `availability = unavailable` and `reason = donation_flow_not_available`. It has no `href` or other activation target. Downstream frontend may explain the state but may not infer or enable donation eligibility.
- **R8 — Truthful media states.** `media` is a discriminated union: `available` requires at least one `PublicCampaignMediaItem`; `absent` and `unavailable` require an empty item list and remain semantically distinct. Each public item exposes only `id`, same-origin `content_url`, `content_type`, non-empty `alt_text`, nullable `caption`, and `source = organizer`. Each object branch/item is closed under R3.
- **R9 — Controlled media delivery.** `GET /campaigns/{campaignId}/media/{mediaId}/content` is anonymous/public-only but rechecks parent eligibility and media membership on every origin request. It returns only documented JPEG/PNG bytes on success, the R2 `404` for absent/non-public/non-member resources, and `503` Problem Details for an eligible metadata record whose storage dependency/object cannot currently serve bytes. It never redirects to or returns an anonymously readable object-storage URL.
- **R10 — Cache/revocation contract.** Successful public detail and media responses, plus public `404`/`503` responses at these operations, specify `Cache-Control: private, no-store`. Retraction means metadata becomes unavailable and a previously known controlled content URL can no longer obtain bytes from the origin; already downloaded client-held bytes are explicitly outside the guarantee.
- **R11 — Contract executability.** `getPublicCampaignDetail` and `getPublicCampaignMediaContent` have stable `operationId`s, required/nullable semantics, valid Problem Details responses, and no `$ref` siblings at touched coordinates. The dereferenced bundle preserves `additionalProperties: false` for every R3 public object schema and contains no permissive object in that exact-projection graph. Root `bearerAuth` resolves through `common.yaml`. Split-source validation has zero errors and no new warning fingerprint from this delta; unrelated historical warnings need not be removed.
- **R12 — Generated correspondence.** `api/openapi.yaml` is regenerated from split sources. `frontend/lib/api/generated/openapi.ts` is regenerated via a committed `generate:api-types` script from that bundle, compiles under the current frontend TypeScript setup, and contains the two touched operations/public schemas without handwritten duplicate response models.
- **R13 — Delivery authority consistency.** Campaign Task 02/03, `INV-campaign-14`, the Campaign threat model, backend architecture wording, OpenAPI, integration map, and tracker do not contradict the rules above. Historical listing/upload/lifecycle breadth is explicitly `DEFER`, not silently deleted as product history or pulled into Slice 1.
- **R14 — Milestone honesty.** Tracker promotion to `CONTRACT_READY` occurs only after R1–R13 reconciliation and executable contract checks pass. Backend/frontend/topology remain not started/unverified, and no orchestration Control Surface/Run state is rewritten by the Build artifact.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Keep `GET /campaigns/{campaignId}` as the stable public-only composite read. | Chosen | No live consumer requires historical mixed semantics; one auditable public projection avoids optional-auth/cache leakage. |
| D1-alt | Keep mixed public/privileged response or add a duplicate public path. | Rejected | Mixed payload is unsafe; a duplicate path creates migration/semantic cost without a live consumer. |
| D2 | Use the exact standalone `PublicCampaignDetail` fields in R3 and nested public schemas described in R4–R8, with `additionalProperties: false` on each exact-projection object. | Chosen | Resolves schema naming/shape before Build and keeps persistence/internal models non-public by default. The closed-object requirement executes the already-settled allowlist; it does not alter the public field set or runtime behavior. |
| D2-alt | Generic key/value facts or nullable-everything payload. | Rejected | Weakens generated-type guarantees and forces frontend inference of absence/provenance semantics. |
| D3 | Keep category, location, beneficiary description, Organization-review status, `donor_count`, and broad Organization profile data out of Slice 1. | Deferred | No active authority demonstrates they are required for the minimum public-understanding outcome. |
| D4 | Plain text only for purpose/story; `story.source = organizer`. | Chosen | Meets product need without opening an HTML/Markdown sanitization contract. |
| D4-alt | HTML/Markdown or frontend-inferred provenance. | Rejected | Adds unnecessary stored-XSS/sanitization risk or lets presentation invent source semantics. |
| D5 | Use tagged funding availability plus decimal-string amount/percentage semantics and an uncapped relationship enum. | Chosen | Distinguishes unknown from zero, prevents `float64` money, and preserves over-target truth without React owning authoritative calculation. |
| D6 | Encode `donation_action` as an unavailable, non-activatable capability. | Chosen | Makes current next-action truth explicit while preventing a fake Slice-2 control. |
| D7 | Carry ordered media metadata in detail and add `/campaigns/{campaignId}/media/{mediaId}/content`; remove the anonymous attachment-list GET from the active contract. | Chosen | One detail request serves the page; controlled bytes enforce parent visibility and revocation. Historical upload remains deferred. |
| D7-alt | Public bucket/direct or long-lived signed URLs. | Rejected | Both permit an unacceptable post-retraction fetch window; endpoint-gated metadata alone is insufficient. |
| D8 | Use `Cache-Control: private, no-store` on the touched public boundary. | Chosen | Avoids system-controlled intermediary/browser reuse defeating origin rechecks; performance caching is deferred until it can preserve revocation. |
| D8-alt | Short public caching TTL. | Rejected | Even a short known stale window contradicts the selected legal/privacy/safety withdrawal guarantee. |
| D9 | Fix only the root `bearerAuth` error and warnings introduced/exposed at touched operations. | Chosen | Makes generation executable without converting Slice 1 into a 130-warning repository cleanup. |
| D10 | Generate one types-only aggregate at `frontend/lib/api/generated/openapi.ts`; keep fetch functions and SDK behavior out of this WU. | Chosen | Fits current frontend architecture, prevents handwritten response models, and avoids speculative client generation. |
| D11 | Reconcile existing Campaign Task 02/03 files and `INV-campaign-14` rather than add a parallel Slice-1 spec hierarchy. | Chosen | These files already own the exact domain concerns; rewriting the relevant sections avoids competing delivery authority. |
| D12 | Do not modify Organization detail/invariants in this WU. | Chosen | The public steward is a dedicated projection composed by Campaign; historical authenticated Organization detail is not a dependency. Campaign spec records that boundary explicitly. |
| D13 | No application-level rate-limit addition in reconciliation. | Chosen | This Run changes no runtime; availability/abuse controls can be decided with the backend implementation and proxy attribution facts. Public-data confidentiality controls remain mandatory now. |
| D14 | A route-local non-documentary missing-media treatment may be used later; no reusable custom placeholder asset is required by the contract. | Chosen | Avoids blocking `CONTRACT_READY` on an intentionally open asset decision while preserving Design review for precedent-setting assets. |

## 6. Backward Compatibility

- **Existing data:** none for Organization/Campaign/media in live migrations. This WU performs no migration or data transformation.
- **API/contracts/clients:** changing `GET /campaigns/{campaignId}` from historical `CampaignDetail` to `PublicCampaignDetail`, removing its `403`, and removing anonymous `GET .../attachments` is contract-breaking on paper. It is accepted because no backend route, generated frontend artifact, or live Campaign consumer exists. PATCH/DELETE on the shared path and deferred upload POST remain untouched.
- **Historical schemas:** remove `CampaignDetail` and `OrganizationSummary` if they become unreferenced. Keep `Campaign`, `CampaignProgress`, `CampaignListItem`, and `CampaignAttachment` only for untouched historical operations; add explicit comments that they are not the public Slice-1 detail projection and remain unreconciled outside the touched boundary.
- **Future Slice 2/3:** donation availability and closed-result continuity require additive/reconciled enum/shape changes later. Slice 1 must not pre-claim their runtime behavior, but it must keep public identity and nested boundaries extensible.
- **Migration/deprecation compatibility:** no deprecation window is needed because the replaced public operations are unimplemented. Record this fact in the feature spec rather than implying a deployed migration.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| RISK-1 | Future internal field addition leaks through public JSON. | Medium | High | Standalone allowlist schema/mapping; forbidden-field negative tests are required in downstream backend work. |
| RISK-2 | Absent vs non-public response reveals existence through status/body/header/auth variance. | Medium | High | Same `404` contract; optional Authorization ignored; threat model requires parity evidence. Timing parity remains a downstream implementation/test concern. |
| RISK-3 | Previously known media URL remains fetchable after retraction. | Medium | High | Private storage, controlled content route, parent/member recheck per request, no-store caching, downstream retraction test. |
| RISK-4 | Storage outage or missing object is presented as intentional no-media. | Medium | Medium | `media.unavailable` and content `503` are distinct from `media.absent`/`404`; downstream tests cover both. |
| RISK-5 | Cache serves stale public detail/media after visibility withdrawal. | Medium | High | Exact no-store contract; topology implementation must not override it. Already downloaded bytes remain accepted/unavoidable exposure. |
| RISK-6 | Money loses precision or over-target truth is capped/hidden. | Medium | High | Decimal strings only, backend-authored uncapped percentage/relationship, no donor count, schema/type verification. |
| RISK-7 | Plain text is later rendered as HTML or source labels imply verification. | Low | High | Spec fixes plain-text/source semantics; downstream frontend tests prohibit unsafe renderer and require visible provenance. Final wording remains Human Design-reviewed. |
| RISK-8 | Generated aggregate/types drift from split sources. | Medium | Medium | Deterministic bundle/type generation commands, committed generated diffs, compile check, independent correspondence inspection. |
| RISK-9 | Existing warning backlog hides new touched warnings. | Medium | Medium | Compare warning coordinates/fingerprints; require zero errors and zero new warnings from the delta, not repository-wide zero warnings. |
| RISK-10 | Tracker says `CONTRACT_READY` while one authority surface still contradicts the contract. | Low | High | Update tracker last; R13 consistency sweep and Human gate precede milestone claim. |
| RISK-11 | Scope expands into topology or production implementation while reconciling docs. | Medium | Medium | Files NOT Changed fence in §11; route runtime mechanics remain anchors for downstream Work Units only. |

## 8. Interface Contract

**Persistence/data shape:** Not changed in this WU. The reconciled spec states the minimum future persisted facts—public identifier, title, purpose, plain-text story, steward identity/name, public lifecycle dates/eligibility, IDR target/current amounts, and ordered media metadata—but does not freeze table/migration/seed mechanics. Public schemas are mapping projections, never persistence serialization.

**API/external interface:** Base URL remains same-origin `/api`.

```text
GET /campaigns/{campaignId}
operationId: getPublicCampaignDetail
security: []
200 application/json: PublicCampaignDetail
404 application/problem+json: PublicCampaignNotFound (absent/non-public/invalid, identical)
503 application/problem+json: PublicCampaignUnavailable (dependency unavailable)
Cache-Control: private, no-store
```

`PublicCampaignDetail` is an exact closed object (`additionalProperties: false`) with this required top-level contract:

| Property | Type / contract |
|---|---|
| `id` | UUID; stable public Campaign identity |
| `title` | non-empty string |
| `purpose` | `PublicCampaignOrganizerText { content: non-empty plain string, source: organizer }` |
| `story` | `PublicCampaignOrganizerText { content: plain string, source: organizer }` |
| `steward` | `PublicCampaignSteward { id: uuid, name: string }` |
| `lifecycle` | `PublicCampaignLifecycle { public_state: fundraising, published_at: date-time, fundraising_ends_at: date-time }` |
| `funding` | tagged `PublicCampaignFundingAvailable` / `PublicCampaignFundingUnavailable` union per R6 |
| `media` | tagged available/absent/unavailable union per R8 |
| `donation_action` | `PublicCampaignDonationAction { availability: unavailable, reason: donation_flow_not_available }` |

Nested schema contract is exact as follows:

| Schema | Required shape |
|---|---|
| `PublicCampaignOrganizerText` | closed object; non-empty plain `content`, `source = organizer` |
| `PublicCampaignSteward` | closed object; `id: uuid`, `name: string` |
| `PublicCampaignLifecycle` | closed object; `public_state: fundraising`, `published_at: date-time`, `fundraising_ends_at: date-time` |
| `PublicCampaignFundingAvailable` | closed object; `availability = available`, `currency = IDR`, decimal-string `target_amount`, decimal-string `collected_amount`, and `progress` |
| `PublicCampaignFundingUnavailable` | closed object; `availability = unavailable`, `reason = not_available`; no amount properties |
| `PublicCampaignFundingProgress` | closed object; `state = computed \| not_computable`, nullable decimal-string `percentage`, and `relationship = below_target \| target_reached \| above_target \| not_computable` |
| `PublicCampaignMediaAvailable` | closed object; `state = available`, `items` with `minItems: 1` |
| `PublicCampaignMediaAbsent` | closed object; `state = absent`, `items` with `maxItems: 0` |
| `PublicCampaignMediaUnavailable` | closed object; `state = unavailable`, `reason = temporarily_unavailable`, `items` with `maxItems: 0` |
| `PublicCampaignMediaItem` | closed object; `id`, same-origin `content_url`, `content_type`, non-empty `alt_text`, nullable `caption`, and `source = organizer` only |
| `PublicCampaignDonationAction` | closed object; `availability = unavailable`, `reason = donation_flow_not_available`; no URL/action target |

`PublicCampaignFunding` uses `oneOf` with discriminator `availability`; `PublicCampaignMedia` uses `oneOf` with discriminator `state`. These union wrapper schemas are composition-only; every object branch they expose is closed. For available funding, `state = not_computable` pairs with `percentage = null` and `relationship = not_computable`; `state = computed` requires a non-negative, uncapped percentage matching `^(0|[1-9][0-9]*)(\.[0-9]{1,2})?$` and one of the other three relationships. All public schema properties are listed in `required`, including deliberately nullable `caption`/`percentage`, and every listed exact-projection object sets `additionalProperties: false`, so omission or an undeclared property does not acquire accidental meaning. OpenAPI 3.0 limitations that cannot express every cross-field condition do not weaken the feature-spec rule or its downstream tests.

Funding state mapping is also fixed: missing product-owned funding facts use `PublicCampaignFundingUnavailable`; target/current dependency failure fails the whole detail request as `503` rather than manufacturing unavailable product truth; a present target `<= 0` yields `not_computable`; otherwise the backend compares exact decimal amounts to select below/equal/above relationship. Media `absent` means no public media is intentionally associated; `unavailable` means media is expected/known but cannot currently be presented. A storage/object failure must never be converted to `absent`.

```text
GET /campaigns/{campaignId}/media/{mediaId}/content
operationId: getPublicCampaignMediaContent
security: []
200 image/jpeg | image/png: binary
404 application/problem+json: PublicCampaignNotFound (absent/non-public/non-member, identical)
503 application/problem+json: PublicCampaignUnavailable (eligible metadata exists but bytes/dependency unavailable)
Cache-Control: private, no-store
```

`PublicCampaignNotFound` and `PublicCampaignUnavailable` are reusable response components local to `campaign.yaml`. Each contains the applicable `Problem` schema plus an exact `Cache-Control: private, no-store` response header. This avoids invalid `$ref` siblings and avoids changing every unrelated shared response in `common.yaml`.

`content_url` is an opaque same-origin URI-reference rooted at `/api/campaigns/{campaignId}/media/{mediaId}/content`; frontend consumes it but does not derive authorization from IDs or translate it into an object-storage URL.

**Cross-layer/business boundary:**

- Backend owns public eligibility, mapping allowlist, funding availability/progress, media membership, and action availability.
- Frontend owns safe text presentation, number/date formatting that does not recalculate business meaning, final reviewed labels, and observable UI state mapping.
- Public detail embeds the narrow steward projection; there is no frontend call to authenticated `GET /organizations/{organizationId}`.
- Root topology owns `/api` prefix stripping and private-bucket policy later; this WU records but does not implement those obligations.
- The public response must not vary by optional authentication, so shared/public caches and consumer types have one auditable shape even though this WU selects no caching.

## 9. Architecture / Plan

Execution order is intentionally linear because later generated/status artifacts derive from earlier authority:

1. Reconcile Campaign tasks/features first: mark only public detail and controlled media delivery active for Slice 1; classify listing, upload, authenticated detail, lifecycle/self-service breadth as `DEFER`.
2. Rewrite `INV-campaign-14` and the matching threat-model rows so eligibility, anti-enumeration, projection allowlisting, plain text, revocation, cache, and dependency-failure semantics are authoritative before schema editing.
3. Narrowly update backend architecture: split sources are authored, the aggregate is generated, and Campaign public media uses private storage plus controlled delivery. Do not design storage interfaces or topology commands here.
4. Edit split OpenAPI: wire root `bearerAuth`, replace the public detail response, add controlled media content path/schemas, remove the anonymous attachment-list GET, and leave unrelated operations untouched.
5. Validate split sources, regenerate `api/openapi.yaml`, add the types-only generation script, and generate `frontend/lib/api/generated/openapi.ts`.
6. Update the integration map with one structural row. Update tracker authority wording/current selection and claim `CONTRACT_READY` only after all prior checks and consistency review pass.
7. Perform an allowlist/closed-object/error/cache/generated correspondence sweep across every touched authority. Verify the dereferenced bundle retains `additionalProperties: false` on every §8 exact-projection object and the regenerated TypeScript declarations expose only the same named properties (with no permissive index signature). Do not start Build for backend/frontend/topology from this Run.

The topology dependency remains explicit for downstream execution:

```text
reconciled public contract
        ↓
backend public projection + controlled media     frontend against generated types/MSW
        ↓                                         ↓
private MinIO policy + /api proxy correction → real same-origin integration
```

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `docs/spec/4-campaign/tasks.md` — Task 02/03 | Current local index still assumes creation-first historical domain order and public attachments. | Mark Slice-1 active subsets and seeded/operator-assisted prerequisite; classify listing/upload/operational breadth `DEFER`; retain historical task identities. |
| `docs/spec/4-campaign/features/02-campaign-detail-listing.md` — Detail | Owns historical mixed detail/listing behavior. | Reframe as active Slice-1 public detail acceptance criteria; make listings explicit deferred historical evidence; record risk tier 1 for downstream implementation. |
| `docs/spec/4-campaign/features/03-campaign-media.md` — List/Upload | Owns historical public-bucket media assumption. | Reconcile controlled content delivery and media-state acceptance; defer list/upload; record exact retraction and `503` behavior. |
| `docs/spec/4-campaign/invariants.md` — `INV-campaign-14` | Current statement is published-only + privileged `403`. | Replace with public-eligibility predicate, identical `404`, optional-auth invariance, projection boundary, and media parent/member visibility. |
| `docs/spec/4-campaign/threat-model.md` — public detail/media | Current model accepts existence disclosure and incorrectly claims no media residual risk. | Replace STRIDE rows/residuals with allowlist, enumeration, auth variance, plain-text/XSS boundary, private storage, revocation/cache, and dependency-failure evidence obligations. |
| `docs/project/kencleng-backend-tech-stack.md` — File Storage; API Contract & Codegen | Historical architecture says public Campaign bucket and hand-authored aggregate source. | Narrowly correct to private Campaign storage + controlled delivery; describe split sources as authored and bundle as generated. Preserve no-Go-codegen decision. |
| `api/openapi/index.yaml` — root security/components and paths | Root security currently cannot resolve `bearerAuth`; new media path needs registration. | Add external `components.securitySchemes.bearerAuth` ref to `common.yaml`; register controlled content path. |
| `api/openapi/campaign.yaml` — `/campaigns/{campaignId}` GET | Current public operation returns inherited internal schema and `403`. | Add `getPublicCampaignDetail`, exact R1–R7 responses/headers/schema; leave PATCH/DELETE intact. |
| `api/openapi/campaign.yaml` — `/campaigns/{campaignId}/attachments` GET | Current anonymous list returns direct public metadata and has no 4XX. | Remove GET only; leave POST explicitly outside active reconciliation and ensure no public detail references `CampaignAttachment`. |
| `api/openapi/campaign.yaml` — new media content path | Required controlled presentation reference. | Add `getPublicCampaignMediaContent` with UUID parameters, JPEG/PNG binary `200`, identical `404`, dependency `503`, and no-store header. |
| `api/openapi/campaign.yaml` — public schemas/responses | No explicit public schemas or cache-bearing public error responses exist. | Add exact required, closed-object (`additionalProperties: false`) schemas/unions and local `PublicCampaignNotFound`/`PublicCampaignUnavailable` responses in §8; remove unreferenced `CampaignDetail`/`OrganizationSummary`; keep historical schemas isolated. |
| `api/README.md` — structure/editing workflow | Owns split/bundle mechanics and known warning context. | Clarify authored/generated authority, document frontend type command/path, and state touched-warning policy without hiding backlog. |
| `api/openapi.yaml` | Generated aggregate used for Swagger/type generation. | Regenerate only; never hand-edit. Review diff for unexpected cross-domain changes. |
| `frontend/package.json` — scripts | `openapi-typescript` exists but no generation command. | Add `generate:api-types` invoking `openapi-typescript ../api/openapi.yaml -o lib/api/generated/openapi.ts`; no SDK/client generation. |
| `frontend/lib/api/generated/openapi.ts` | First generated frontend contract artifact. | Generate and commit; never hand-edit. Downstream frontend imports contract types from it. |
| `docs/project/kencleng-integration-map.md` — Active mappings | No cross-stack mapping exists. | Add Public Campaign Detail row referencing both operationIds, Campaign backend owner, embedded steward projection, and controlled media coordination note; add no status column. |
| `docs/project/kencleng-development-tracker.md` — Product Authority, Campaign, next selection, merge gate | Contains stale pre-promotion/PR #27 wording and `NEEDS_RECONCILIATION`. | Record Product Authority as canonical, replace obsolete merge gate, and set Campaign/Slice 1 contract to `CONTRACT_READY` only after evidence; keep backend/frontend/integration unverified. |

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `docs/spec/4-campaign/tasks.md` | Modify | Active Task 02/03 scope and dependency/classification reconciliation. |
| `docs/spec/4-campaign/features/02-campaign-detail-listing.md` | Modify | Reconciled Slice-1 public detail spec. |
| `docs/spec/4-campaign/features/03-campaign-media.md` | Modify | Reconciled public metadata/content delivery spec; upload/list deferred. |
| `docs/spec/4-campaign/invariants.md` | Modify | Replace `INV-campaign-14` only. |
| `docs/spec/4-campaign/threat-model.md` | Modify | Replace public detail/media rows and residual-risk record only. |
| `docs/project/kencleng-backend-tech-stack.md` | Modify | Narrow authority corrections for Campaign media storage and OpenAPI source topology. |
| `api/openapi/index.yaml` | Modify | Resolve shared security scheme and register content path. |
| `api/openapi/campaign.yaml` | Modify | Public operation/media path/schemas and touched response cleanup. |
| `api/README.md` | Modify | Current generation/validation workflow and artifact path. |
| `api/openapi.yaml` | Regenerate | Bundled aggregate. |
| `frontend/package.json` | Modify | Add deterministic types generation script. |
| `frontend/lib/api/generated/openapi.ts` | Add/generated | Committed aggregate TypeScript type definitions. |
| `docs/project/kencleng-integration-map.md` | Modify | First active structural mapping. |
| `docs/project/kencleng-development-tracker.md` | Modify last | Current authority/contract state and obsolete gate cleanup. |

| File / area intentionally untouched | Why |
|---|---|
| `docs/product/**`, `docs/ui-ux/**` | Upstream authority is sufficient; no Product/Design revision is authorized or needed. |
| `docs/spec/3-organization/**`, `api/openapi/organization.yaml` | Public steward is a Campaign-owned composite projection; authenticated Organization detail remains deferred evidence. |
| Other Campaign features/invariants/OpenAPI operations | Full lifecycle/listing/upload reconciliation is outside Slice 1. |
| `api/openapi/common.yaml` | Existing `Problem` schema and security-scheme definition suffice; campaign-local cache-bearing `404`/`503` response components reference `Problem` without changing shared responses. |
| `backend/**` | Production backend is explicitly outside this Reconciliation Work Unit; protected Tier-0 paths remain untouched. |
| Frontend files other than package script/generated types | No production route, data access, mock, component, copy, or test is implemented here. |
| `Caddyfile`, `docker-compose.yml` | Root topology/private-bucket implementation belongs to downstream explicitly owned work. |
| `.harscode-spaces/**/manifest.md`, `control-surface.md` | Orchestration Operator owns Run/Work Unit state; Build must not self-promote orchestration records. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | Cross-read feature spec, `INV-campaign-14`, threat model, and dereferenced `getPublicCampaignDetail`; confirm `security: []` and one response projection independent of auth. | Testing | Prevents a mixed-auth payload from surviving in one layer. |
| R2 | Contract inspection confirms no `403` on either public GET and the same `PublicCampaignNotFound` response for absent/non-public; threat spec requires identical status/body/header and downstream timing checks. | Testing | Existence disclosure is a security boundary, not copy polish. |
| R3 | In the dereferenced `api/openapi.yaml`, inspect `PublicCampaignDetail`, `PublicCampaignOrganizerText`, `PublicCampaignSteward`, `PublicCampaignLifecycle`, `PublicCampaignFundingAvailable`, `PublicCampaignFundingUnavailable`, `PublicCampaignFundingProgress`, `PublicCampaignMediaAvailable`, `PublicCampaignMediaAbsent`, `PublicCampaignMediaUnavailable`, `PublicCampaignMediaItem`, and `PublicCampaignDonationAction`; each MUST retain `additionalProperties: false`. Confirm no `allOf Campaign/Organization`, no forbidden fields, and no permissive object schema in the public exact-projection graph. Regenerate `frontend/lib/api/generated/openapi.ts` and confirm the corresponding declarations contain only the same named properties and no permissive index signature. | Testing | `required` does not reject undeclared properties in OpenAPI 3.0; this proves the allowlist is executable in the dereferenced contract and remains synchronized with its generated consumer artifact. |
| R4 | Spec/schema inspection confirms plain string + `source: organizer`, with no HTML/Markdown/content-format escape hatch. | Testing | Prevents the reconciliation itself from opening stored-XSS or false-verification semantics. |
| R5 | Inspect public lifecycle enum/dates and absence of internal status/reason schemas. | Testing | Raw state leakage would make frontend infer internal workflow and impede Slice-3 reconciliation. |
| R6 | Inspect OpenAPI types: every money/percentage field is string/nullable as specified, no `number`/`float`, relationship includes over-target/not-computable; examples cover zero/large/over-target. | Testing | Type mistakes can reintroduce precision loss before runtime code exists. |
| R7 | Confirm generated donation-action type has only unavailable Slice-1 state/reason and no URL. | Testing | Avoids contract-level permission for a fake CTA. |
| R8 | Inspect the discriminated media unions, min/max item constraints, closed `PublicCampaignMediaItem`/branch objects, exact item allowlist, and absent vs unavailable distinction. | Testing | Prevents metadata leakage and false no-media presentation. |
| R9 | Inspect controlled media operation content types/error shapes and absence of redirect/object-storage URL semantics; confirm attachment-list GET removed from bundle. | Testing | Metadata gating without byte gating does not provide revocation. |
| R10 | Inspect exact `Cache-Control` header contract on success and error responses in dereferenced bundle and threat/spec text. | Testing | A stale intermediary could defeat visibility withdrawal. |
| R11 | `cd api && npm run validate`; require zero errors and no new touched warning coordinates. Also inspect output against the recorded baseline of 1 error/130 warnings. | Build | This is the fastest executable proof that references and touched OpenAPI structure are usable; skipping it risks broken bundling/generation. |
| R11 | `cd api && npm run bundle`, then review the generated diff for only intended contract changes. | Build | Bundle success alone can still preserve an unintended schema diff; both execution and review matter. |
| R12 | `cd frontend && npm run generate:api-types`; `./node_modules/.bin/tsc --noEmit`; `npm run lint`. | Build | Proves the generated artifact is syntactically/type-tool compatible without paying for an unrelated production build. |
| R12 | Independently generate bundle/types to `/tmp` and `cmp` them with committed artifacts; inspect both operationIds in generated output. | Testing | Detects hand edits or stale generated artifacts. Full Next build is not cost-effective because no production frontend code imports the artifact yet. |
| R13 | Path-by-path semantic consistency review across all files in §11, including `KEEP/ADAPT/REPLACE/DEFER` wording. | Testing | Individual green tools cannot detect contradictory delivery authority prose. |
| R14 | Inspect tracker last: `CONTRACT_READY` evidence links are present while backend/frontend/topology/integration milestones remain unclaimed; inspect Git diff for no Control Surface mutation. | Human | Milestone meaning and cross-authority acceptance require conscious human review. |
| R1–R14 | `git diff --check` and scoped `git status --short` review. | Build | Catches malformed documentation/generated patches and accidental scope expansion cheaply. |

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Public projection regression | Internal/public data boundary can expand silently. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001/evidence/solutioning.md#risks-and-required-carry-forward-evidence` item 1 | Yes — retained by R3/RISK-1 and Campaign threat model; runtime negative tests belong to downstream backend Testing. |
| Existence disclosure / optional auth variance | Response differences reveal non-public resources or richer data. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001/evidence/solutioning.md#risks-and-required-carry-forward-evidence` item 2; `gap-analysis.md#area-4--existing-shared-api-contract` | Yes — retained by R1–R2/RISK-2; timing/parity evidence belongs to downstream backend Testing. |
| Media revocation and cache staleness | A known URL or owned cache can bypass lifecycle withdrawal. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001/evidence/solutioning.md#risks-and-required-carry-forward-evidence` items 3–4 | Yes — retained by R9–R10/RISK-3–5; real MinIO/proxy evidence belongs to downstream backend/topology/integration Testing. |
| Money truth | Precision loss or client-side authoritative calculation can misstate funding facts. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001/evidence/solutioning.md#risks-and-required-carry-forward-evidence` item 5 | Yes — retained by R6/RISK-6; decimal calculation/serialization evidence belongs to downstream backend and frontend Testing. |
| Organizer-controlled content | Unsafe content rendering can create stored XSS or false provenance. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001/evidence/solutioning.md#risks-and-required-carry-forward-evidence` item 6 | Yes — contract mitigation is R4; hostile-looking plain-text rendering evidence belongs to downstream frontend Testing. |
| Runtime concurrency/performance | This WU changes no stateful runtime or hot-path implementation. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001/evidence/gap-analysis.md#area-5--backend-current-state` | N/A — no runtime code in scope (D13); downstream plans must reassess when implementation exists. |

## 13. Open Items

### Active — needs external input or verification

1. **Final Indonesian provenance/action wording — non-blocking for this WU.** Human Design review remains required before production wording is promoted. The machine contract deliberately uses `organizer`, `unavailable`, and `donation_flow_not_available` and must not depend on the final label.

### Resolved — retained as decision history

1. ~~**Exact public schema/property boundary**~~ **RESOLVED — R3–R8 and §8 define executable names, semantics, and `additionalProperties: false` closed-object behavior for every exact-projection object. This resolution only makes the reviewed allowlist executable; it does not change material interface or security semantics, so a mandatory re-review is not recommended before the Human Techplan gate.**
2. ~~**Controlled media byte path/cache contract**~~ **RESOLVED — use `/campaigns/{campaignId}/media/{mediaId}/content` with parent/member recheck and `private, no-store` (D7–D8).**
3. ~~**Frontend generated-type location**~~ **RESOLVED — `frontend/lib/api/generated/openapi.ts`, generated by `generate:api-types` (D10).**
4. ~~**Organization detail dependency**~~ **RESOLVED — no separate Organization fetch or Organization-spec change; use embedded `PublicCampaignSteward` (D12).**
5. ~~**Reusable placeholder asset needed for contract readiness**~~ **RESOLVED — no; truthful route-local treatment is sufficient, with later Human Design review if it becomes reusable/expressive (D14).**
