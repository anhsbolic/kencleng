# EXP-001 — Stage 3 Solutioning

- Phase/Stage: Exploration / Stage 3 — Solutioning
- Author: Explorer
- Created: 2026-09-21
- Model: `gpt-5.6-sol`
- Reasoning: `high`
- Target revision: `464079eb03d3ddc8d658724ba1a5afa968769728` (`validation-03-orchestrator-slice-1`)
- Workflow revision: `c8b3eca93e8d265438c2dba8bc37b6b90b251724` (`pilot/orchestrator-v0.1`)
- Work Unit: `WU-S1-001`
- Run: `EXP-001`
- Role: `Explorer`
- Specialization: none

## Scope and evidence posture

Dokumen ini memilih arah material untuk menutup gap yang dicatat di `evidence/gap-analysis.md`. Ini adalah durable Exploration decision evidence, bukan detailed Techplan, implementation contract, atau perubahan Product/Design authority.

Stage 2 menginspeksi source/implementation pada target revision `c78cae58cedf9cac34dedcd3e068f15cd55d3b6d`. Audit delta sampai target revision Stage 3 menunjukkan hanya Harscode Space records dan Stage-2 evidence yang berubah; Product, Design, specs, OpenAPI, application code, dan topology files yang menjadi dasar analisis tidak berubah. Karena itu Stage-2 findings masih current untuk solutioning ini.

Portable guidance yang cocok dengan concrete concerns juga menguatkan arah di bawah tanpa menggantikan Kencleng authority:

- `{HARSCODE_WORKSPACE_ROOT}/best-practices/restapi/anti-enumeration.md` — public non-visibility tidak boleh membocorkan existence melalui response variance;
- `{HARSCODE_WORKSPACE_ROOT}/best-practices/restapi/openapi-spec-first-drift.md` — contract, generated types, handler, dan error shape harus dijaga correspondence-nya;
- `{HARSCODE_WORKSPACE_ROOT}/best-practices/go/decimal-and-money.md` — monetary values tidak boleh melewati `float64`;
- `{HARSCODE_WORKSPACE_ROOT}/best-practices/pwa/xss-and-content-sanitization.md` — normal React text rendering lebih aman daripada membuka HTML/Markdown rendering tanpa kebutuhan dan sanitization contract;
- `{HARSCODE_WORKSPACE_ROOT}/best-practices/go/file-upload-handling.md` — relevan kelak jika upload masuk scope, tetapi upload API tidak diperlukan untuk Slice 1.

## Chosen delivery direction

Slice 1 sebaiknya dijalankan sebagai narrow reconciliation-first vertical slice:

```text
Product/MVP + Design authority
        ↓
reconciled Slice-1 feature/invariants/threat model
        ↓
explicit public OpenAPI contract
        ↓
backend + root topology work     frontend against contract-faithful MSW
                 \                 /
                  real integration
                         ↓
              human rendered acceptance
                         ↓
                  slice finalization
```

Tidak ada alasan untuk merevisi Product Authority sekarang. Gap yang tersisa dapat diselesaikan sebagai delivery/design/contract detail selama solusi mempertahankan public-data safety, truthful uncertainty, dan human gates yang sudah berlaku.

## Decision 1 — Reconcile the active slice before implementation

### Chosen

Buat satu reconciled Slice-1 delivery contract yang sempit sebelum production Build. Contract tersebut perlu menggabungkan:

- public Campaign visibility dan not-found-equivalent behavior;
- explicit public Campaign, steward, funding, media, provenance, lifecycle, dan action projections;
- content-safety semantics;
- applicable invariants dan threat concerns;
- required UI-observable states;
- OpenAPI operation/schema/error behavior yang cukup stabil untuk backend implementation dan frontend mocks/generated types.

Reconciliation hanya menyentuh concern yang dibutuhkan Public Campaign Detail. Historical Campaign/Organization lifecycle breadth tetap evidence atau `DEFER`.

### Rejected alternatives

| Alternative | Why rejected / consequence |
|---|---|
| Backend-first dari historical specs/OpenAPI | Akan mengunci unsafe `CampaignDetail`, raw Organization status, `403` existence disclosure, dan public-bucket assumptions sebelum authority gap ditutup. |
| Frontend-first dengan handwritten Campaign model | Membuat React menjadi de facto contract/business authority dan menambah rewrite risk saat public projection direkonsiliasi. |
| Reconcile seluruh Campaign/Organization/OpenAPI | Melanggar per-slice contract discipline dan menarik self-service, curation, closure, serta operational breadth yang tidak diperlukan. |
| Menunda spec/threat reconciliation sampai sesudah code | Membuat security/public-data behavior tersebar sebagai incidental implementation choices dan sulit diaudit. |

### Consequences

- Milestone pertama yang masuk akal adalah `CONTRACT_READY`, bukan backend atau UI completion.
- Delivery spec, invariants, threat model, split OpenAPI source, bundle, dan generated frontend types harus konsisten pada boundary yang disentuh.
- Existing broad draft specs tidak perlu dibersihkan seluruhnya; active Slice-1 reconciliation harus jelas menandai bagian historical yang `ADAPT`, `REPLACE`, atau `DEFER`.

## Decision 2 — Make the existing public detail path public-only

### Chosen

`GET /campaigns/{campaignId}` tetap menjadi stable composite public read, tetapi response-nya diganti menjadi explicit allowlisted `PublicCampaignDetail` dan tidak lagi berfungsi sebagai mixed public/privileged read.

Behavior boundary:

```text
eligible public Campaign
→ 200 PublicCampaignDetail

absent Campaign OR Campaign not eligible for public visibility
→ same public-safe 404 Problem Details behavior

Authorization header present or absent
→ same public projection and visibility behavior
```

Future representative/Curator/Admin detail harus memakai distinct authenticated operation/projection ketika operational surface benar-benar masuk scope. Exact internal path tidak perlu didesain pada Slice 1.

### Rejected alternatives

| Alternative | Why rejected / consequence |
|---|---|
| Pertahankan one-path mixed public/privileged payload | Optional auth dapat mengubah response shape, memperbesar cache/data-leak risk, dan membuat public contract tidak auditable. |
| Tambah public path baru sambil mempertahankan historical mixed path | Menambah duplicate semantics dan migrasi tanpa consumer live yang perlu dipertahankan; backend route sendiri belum ada. |
| Return `403` untuk existing non-public Campaign | Membocorkan existence/internal lifecycle dan bertentangan dengan approved public-safe direction. |
| Expose raw `CampaignStatus` | Memberi client internal workflow states yang tidak pernah perlu dilihat public visitor dan mendorong eligibility inference di frontend. |

### Consequences

- Public response harus mempunyai explicit `required`/optional semantics; absence tidak boleh menjadi accident dari weak schema.
- Public lifecycle memakai narrow public semantic, bukan internal status enum. Slice 1 hanya perlu deliver fundraising/public-understanding behavior; `closed` delivery tetap Slice 3 meskipun public projection harus tidak menghalangi later extension.
- `GET /organizations/{organizationId}` tidak menjadi dependency frontend Public Campaign Detail; steward context berada di composite public projection.

## Decision 3 — Use a minimum explicit public projection

### Chosen baseline

`PublicCampaignDetail` perlu membawa categories of truth berikut. Exact OpenAPI property names dan optionality ditetapkan saat reconciliation, tetapi semantic boundary ini settled untuk Techplan synthesis.

| Projection area | Include for Slice 1 | Exclude / defer |
|---|---|---|
| Campaign identity/purpose | stable public identifier, title, concise purpose/story content needed by the surface | internal IDs except the public identifier; operational/audit metadata |
| Steward | narrow public Organization identifier and display name | contact, legal/tax data, `has_overdue_report`, representative data, raw `OrganizationStatus` |
| Lifecycle | public-facing fundraising/availability meaning and only deliberately public date context | internal curation/scheduling/unpublish/decision states and reasons |
| Funding | currency, target amount, current collected amount, backend-authoritative progress relationship/availability | `donor_count`, popularity/social-proof fields, internal caps unless a later surface justifies them |
| Story/provenance | organizer-provided plain text plus machine-readable source class sufficient for truthful presentation | unsanitized HTML/Markdown; global claims that the platform authored or verified the story |
| Media | explicit media availability plus narrow ordered public media references/metadata | uploader, original filename, storage key, byte size, internal timestamps |
| Action | backend-authoritative public action capability; unavailable/non-active for Slice 1 | frontend inference from raw status, deadline, or amount; active/fake Donate action |
| Unknown/pending | explicit availability/state where absence has material meaning | generic nullable fields whose meaning must be guessed by React |

Category, location, beneficiary description, donor count, and broad Organization-review messaging are not baseline Slice-1 fields merely because historical schemas contain them. They remain `DEFER` unless the reconciled surface demonstrates a concrete user-understanding need. In particular, no Organization “verified” signal should ship without a Product/Design decision defining its deliberately narrow public meaning.

### Rejected alternatives

- **Full internal-schema inheritance:** rejected because current and future internal fields would leak by default.
- **Generic key/value “facts” payload:** rejected because it weakens generated-type guarantees and moves semantics into frontend copy/config.
- **Separate Organization fetch:** rejected for this surface because it invites reuse of the authenticated full Organization object and adds partial-failure complexity for a deliberately narrow steward projection.
- **Nullable-everything response:** rejected because missing, unknown, pending, and unsupported are different product meanings.

### Consequences

- The public projection is a dedicated mapping boundary, not serialization of persistence/domain records.
- Adding an internal field cannot make it public automatically.
- Tests must assert forbidden-field absence, not only expected-field presence.
- Source/truth-class enum names are contract semantics; final Indonesian labels remain Design-owned presentation copy.

## Decision 4 — Keep organizer content plain text in Slice 1

### Chosen

Campaign purpose/story is plain text rendered through normal React escaping. Preserve paragraphs/newlines through safe presentation, but do not enable arbitrary HTML or Markdown in Slice 1.

Provenance must remain structurally visible. Contract semantics identify organizer-provided content; frontend supplies the approved human-facing label without upgrading it into platform verification.

Meaningful unknowns use explicit state. Examples include:

- media genuinely absent versus media expected but temporarily unavailable;
- funding fact unavailable versus a factual zero;
- no public action available versus action eligibility still being evaluated.

### Rejected alternatives

| Alternative | Why rejected / consequence |
|---|---|
| Markdown/HTML now | Product outcome does not require rich text; it adds sanitizer dependencies, hostile-content testing, URL/attribute policy, and greater stored-XSS risk. |
| Frontend-only provenance labels inferred by field location | Source meaning becomes implicit and can drift when contract composition changes. |
| Treat null/empty as universal “Belum tersedia” | Fabricates one meaning for several materially different states. |

### Consequences

- No `dangerouslySetInnerHTML` path is needed.
- Exact final provenance/action wording remains subject to Design review; schema semantics must not depend on the chosen Indonesian phrase.
- Rich content can be reconsidered in a later slice through an explicit content-format and hostile-content contract.

## Decision 5 — Make funding and action semantics backend-authoritative

### Chosen

- Monetary values remain decimal strings at the API boundary and established decimal values in backend logic/storage; they never pass through `float64`.
- Backend returns the truthful target/current amounts plus a display-safe progress relationship/availability semantic. It must preserve over-target truth and define zero/invalid-target behavior rather than relying on division in React.
- `donor_count` is excluded because Slice 1 does not need popularity/social-proof pressure.
- Public action capability comes from the backend contract. During Slice 1 it yields no active Donate action because Donation Flow does not exist.

The initial persisted Campaign may truthfully have a zero collected amount. Do not create fake prior donations merely to make a progress bar visually interesting, and do not build the Donation domain early only to seed a nonzero demo.

### Rejected alternatives

- **Frontend computes eligibility/progress from raw fields:** duplicates business and money rules and will drift as Donation/closure behavior arrives.
- **Historical float percentage capped at 100:** can hide over-target facts and introduces an imprecise/underspecified display contract.
- **Seed arbitrary nonzero aggregate with no supporting application truth:** violates the real-state guardrail.
- **Hide the action area completely as if capability meaning does not exist:** weakens later continuity and gives no truthful explanation of current availability.

### Consequences

- The frontend may format decimal strings for display but does not perform authoritative monetary calculations.
- Contract/backend tests cover zero, normal progress, over-target, unavailable/invalid target behavior, and very large decimal strings.
- The Slice-1 page uses a non-interactive availability/status treatment; exact copy is a presentation decision, not a promise of release timing.

## Decision 6 — Use private object storage and controlled public media delivery

### Chosen

Campaign media objects are not anonymously readable from a public bucket. Public detail returns narrow media metadata plus an opaque same-origin presentation reference. Fetching media bytes goes through controlled application/media delivery that rechecks parent Campaign public eligibility and media membership before serving.

The public detail response should own the ordered `PublicCampaignMedia` projection and explicit media state needed by the page. A separate public attachment-list call is not required for Slice 1. Exact byte-delivery path and cache headers are Techplan/spec mechanics inside this boundary.

Retraction guarantee must be stated honestly:

```text
after retraction
→ no new anonymous origin fetch through a previously known presentation URL
→ metadata is no longer public
→ storage object itself remains non-anonymous
```

The product cannot claw back bytes a visitor already downloaded. Caching policy must be threat-modeled so intermediaries do not keep serving material after application visibility is withdrawn.

Missing media is a first-class successful state. Storage/dependency failure or a DB row pointing to a missing object must not be mislabeled as “campaign has no media.”

### Rejected alternatives

| Alternative | Why rejected / consequence |
|---|---|
| Current anonymously downloadable bucket/direct URL | Previously known URLs remain readable after retraction; list gating is cosmetic. |
| Long-lived signed URLs | They reduce bucket exposure but leave a revocation window that may be unacceptable for legal/privacy/safety withdrawal. |
| Endpoint-gated metadata + direct public bytes | Still fails the actual confidentiality/revocation requirement. |
| Generate documentary-looking placeholder imagery | Turns missing evidence into synthetic campaign evidence. |

### Consequences

- `docker-compose.yml` must stop granting anonymous download to the bucket used by Slice-1 media.
- Backend/storage tests need to exercise public, missing, retracted, unknown media ID, and object-missing/dependency-failure behavior.
- A custom reusable Campaign placeholder asset is not required to start. A route-local warm-neutral missing-media treatment using structure/type and an ordinary Phosphor media cue is sufficient if rendered review confirms it is truthful. If it becomes a reusable expressive asset/system, it starts `PROPOSED` and follows human Design approval.
- Upload API remains out of scope. Seed/operator setup may place validated known fixtures into private storage; a future upload capability must independently apply size, magic-byte, filename, streaming, authorization, and threat controls.

## Decision 7 — Build only the persisted backend capability the slice needs

### Chosen

Backend delivery needs the minimum coherent persisted model and public read behavior for:

- a narrow Organization/steward record;
- a Campaign with public identity/content, lifecycle/eligibility, target/current funding facts, and deliberately public dates;
- Campaign media metadata pointing to private object storage;
- an idempotent development/operator fixture path that creates at least one eligible public Campaign and representative non-public/missing-media states;
- repository/service/handler boundaries for public projection and media authorization;
- safe Problem Details and dependency-error behavior.

The setup mechanism must be explicitly non-user-facing and truthful. Prefer an explicit seed/operator command or fixture mechanism over burying mutable sample content in schema migrations. Exact schema and command shape belong to Techplan after the reconciled spec is approved.

### Rejected alternatives

- **Build Organization registration, Campaign authoring, curation, publish UI/API, or scheduler first:** current MVP explicitly permits seeded/operator-assisted setup.
- **Hard-code one response in the handler/frontend:** does not prove persisted application truth or lifecycle gating.
- **Implement the full historical Campaign state machine:** expands scope into operational/self-service behavior not required by Slice 1.
- **Use Account as an enabling dependency:** public visitor and operator-assisted setup do not require it.

### Consequences

- No protected Tier-0 path is needed for Slice 1.
- Public projection/media work is threat-model-raised and should be planned as Tier 1 with independent review/testing and human review before merge; ordinary seed/UI pieces may be lower risk but do not lower the boundary's overall scrutiny.
- No scheduler capability is needed unless Techplan discovers an unavoidable current-slice invariant; deadlines can be represented without building automatic closure, which remains Slice 3.
- Real-Postgres and real object-storage evidence is required for the persistence/media boundary; unit tests alone are insufficient.

## Decision 8 — Use a facts-first editorial Detail surface

### Chosen

Frontend design readiness is `PARTIAL`, not `OPEN`: established Product/Design authority is sufficient to choose the ordinary Detail pattern and resolve local composition through rendered iteration.

The selected route-level direction is:

```text
public identity + concise purpose + lifecycle meaning
→ real media or explicit missing-media state
→ steward + funding + public date/source facts
→ organizer-provided story with visible provenance
→ truthful action-availability area
```

The production route should be a stable campaign detail route such as `/campaigns/[campaignId]`. Route composition starts server-owned where practical; Client Components remain limited to real retry/interaction needs. This is an implementation tendency from current frontend architecture, not a new requirement to add TanStack Query or Zustand.

Required presentation behaviors:

- structure-preserving loading;
- success with real media;
- success with missing media;
- identical public not-found treatment for absent/non-public Campaign;
- recoverable request failure with retry only where retry can help;
- broken/unavailable media distinct from intentional absence;
- long title/story/steward and large money values;
- responsive/high-zoom reading order preserving identity, trust/funding/provenance, and action meaning;
- no active Donate control.

### Rejected alternatives

| Alternative | Why rejected / consequence |
|---|---|
| Story/hero-first conversion composition | Conflicts with facts-first/confidence-before-conversion and risks turning imagery into persuasion. |
| Dense dashboard/evidence console | Makes a public Detail surface feel operational and obscures the calm editorial hierarchy. |
| Clone `sunlit-editorial-public.png` | The reference contains discovery/donation/social-proof concepts not proven by Slice 1 and is not a route contract. |
| Build a broad component system before the route | Clean-start architecture requires contracts to emerge from real usage; broad primitives now would be speculative. |
| Production mock-mode branch | Creates a second data path and can make mock behavior appear real. |

### Consequences

- Contract-faithful MSW intercepts the real request boundary for frontend parallel work.
- OpenAPI-generated types own response shape; no parallel handwritten Campaign API model.
- Route-local/domain components are preferred until repetition proves a shared contract.
- Rendered iteration must cover representative desktop/mobile states, followed by human browser acceptance before `FRONTEND_MOCK_VERIFIED` and again against real integration before final delivery.
- Final provenance wording and any precedent-setting reusable placeholder treatment remain human Design review gates; they do not require a high-fidelity design-file phase.

## Decision 9 — Standardize the local integration boundary narrowly

### Chosen

The browser-facing API base remains same-origin `/api`. Root proxy configuration should strip that prefix before forwarding to the backend's existing root route patterns; do not duplicate `/api` into every backend route merely to match current Caddy behavior.

The Slice-1 contract reconciliation must also make the contract executable enough to support generation and verification:

- resolve the current root `bearerAuth` reference error;
- give the touched public operation stable operation/schema/error definitions;
- eliminate invalid `$ref` sibling behavior on touched responses;
- define requiredness for new public schemas;
- regenerate `api/openapi.yaml` from split sources;
- generate/commit the focused TypeScript contract artifact expected by frontend architecture;
- require no unresolved validation error and no new warnings from the Slice-1 delta.

The existing unrelated warning backlog does not need to become a repository-wide cleanup prerequisite.

### Rejected alternatives

- **Frontend calls backend port directly:** bypasses the real same-origin topology and lets proxy drift survive until final integration.
- **Register duplicate backend routes with and without `/api`:** broadens public routing and creates two contracts for one capability.
- **Require all 130 historical OpenAPI warnings to be fixed:** turns a narrow slice into unrelated contract cleanup.
- **Ignore the validation error because the endpoint is public:** contract generation/bundling health is shared infrastructure and current failure is already concrete.

### Consequences

- Caddy and MinIO policy changes are root-scoped infrastructure work, not incidental frontend/backend edits.
- Integration tests and human browser checks should use the proxied same-origin path, not only direct backend URLs.
- Swagger/OpenAPI presence must not be reported as implementation readiness; handler/contract correspondence needs executable evidence.

## Decision 10 — Use contract-parallel delivery after reconciliation

### Chosen

After `CONTRACT_READY`, backend and frontend can proceed in parallel because the selected public contract is intentionally narrow and the frontend can exercise meaningful states through MSW. Root topology/media-policy work can proceed alongside backend implementation with explicit ownership. Real integration follows only after both sides have their narrower evidence.

Recommended downstream shape derived from current evidence:

1. **Slice-1 Reconciliation Work Unit — `RECONCILIATION` / `ENABLER`**
   Reconcile active feature acceptance criteria, public visibility invariant, threat model, explicit OpenAPI projection/errors, narrow generated-type path, integration-map row, and tracker state. This produces the `CONTRACT_READY` input.
2. **Public Campaign Backend Work Unit — `DELIVERY`**
   Deliver minimum persistence, idempotent fixture/operator setup, public projection, controlled media delivery, and Postgres/object-storage evidence. Plan the public-data/media boundary as Tier 1.
3. **Public Campaign Frontend Work Unit — `DELIVERY`**
   Deliver the Detail route against contract-faithful MSW, observable state tests, responsive/rendered evidence, and human acceptance. Design readiness is `PARTIAL`; no mandatory high-fidelity handoff.
4. **Slice-1 Topology Work Unit — `ENABLER`**
   Own root-scoped Caddy prefix correction and anonymous MinIO policy removal, with focused topology verification. It may be scheduled alongside backend delivery but must land before real integration.
5. **Cross-work-unit integration/finalization — `VERIFICATION` / orchestration step**
   Replace mock interception on the exercised path with the real stack, validate real-data states and public boundaries through the proxy, obtain human integrated browser acceptance, and only then promote `INTEGRATED_VERIFIED` / `SLICE_FINALIZED` / `DELIVERED` as evidence allows.

The operator may choose concrete IDs and whether topology is a distinct Work Unit or an explicitly root-scoped batch inside coordinated delivery. It must not disappear into a backend-only or frontend-only Build.

### Rejected alternatives

| Alternative | Why rejected / consequence |
|---|---|
| Fully sequential backend-first after contract | Valid but adds avoidable latency once the narrow contract is stable and MSW can exercise the complete page-state set. |
| Start FE/BE parallel before reconciliation | High drift risk because the current shared contract is specifically one of the unsafe artifacts being replaced. |
| One cross-stack mega-Work Unit | Blurs write scopes, risk tiers, independent review evidence, and root topology ownership. |
| Declare completion from backend tests + frontend mocks | Misses proxy behavior, storage revocation, mock drift, real loading/failure behavior, and human rendered integration acceptance. |

## Settled classification of material historical evidence

| Existing evidence | Stage-3 classification | Reason |
|---|---|---|
| `GET /campaigns/{campaignId}` composite concept | `ADAPT` | Stable public identity/path is useful; mixed response/auth semantics are not. |
| `CampaignDetail allOf Campaign` | `REPLACE` for public contract | Public projection must be explicit. |
| `OrganizationSummary` | `ADAPT` | Keep only deliberately public steward identity/context. |
| `CampaignProgress` | `ADAPT` | Keep target/current/progress meaning; remove donor/social-proof and capped float assumptions. |
| Public attachment-list projection | `REPLACE` for Slice-1 page contract | Detail should carry narrow ordered media metadata/state; controlled byte delivery remains separate. |
| Public object bucket/direct URLs | `REPLACE` | Cannot satisfy lifecycle revocation. |
| Historical authenticated Organization detail | `DEFER` | Not the public steward contract and not required by the slice. |
| Full Campaign lifecycle/self-service/scheduler | `DEFER` | Seed/operator setup is sufficient; closure delivery remains later. |
| Existing frontend visual foundation | `KEEP` | It is current, accepted implementation precedent for the canonical design direction. |
| Historical broad API warning cleanup | `DEFER` except touched/blocking items | Not required to deliver the narrow slice safely. |

## Risks and required carry-forward evidence

1. **Public projection regression:** tests must fail if internal/audit/steward-private fields appear in public JSON.
2. **Existence disclosure:** absent and non-public Campaign behavior must be indistinguishable at the public contract boundary.
3. **Media revocation:** verify both metadata and previously known presentation URL behavior after parent visibility withdrawal; do not test only the list endpoint.
4. **Cache staleness:** threat/spec must define cache behavior for detail and media so retraction is not defeated by an intermediary owned/configured by the system.
5. **Money truth:** decimal values, zero target, large values, and over-target progress require backend/contract/frontend evidence without `float64` money conversion.
6. **Organizer content:** plain-text rendering and long/hostile-looking strings need observable frontend tests even without an HTML renderer.
7. **Mock drift:** the same generated contract types and production request path must back MSW and real integration.
8. **False milestone promotion:** contract lint, backend unit tests, frontend build, and mock rendering are separate evidence; none alone establishes `INTEGRATED_VERIFIED`.

## Open / deliberately deferred

- Exact OpenAPI property names, schema nesting, media-content path, cache headers, and seed-command mechanics belong to the reconciliation/Techplan work within the settled boundaries above.
- Final Indonesian provenance/action wording is still an intentional Design decision. It requires review before becoming production/canonical wording but does not block Techplan synthesis.
- A reusable custom Campaign placeholder asset/system remains `OPEN`; Slice 1 does not require one if a truthful route-local missing-media treatment passes rendered/human review.
- Category, location, beneficiary description, Organization-review messaging, donor count, broad discovery, full Organization profile, upload/self-service flows, active donation, closure/result, and accountability remain deferred unless owning authority changes.
- Current development-tracker/control-surface wording remains an orchestration-state reconciliation concern; it is not Product or delivery behavior.

No material Product Authority gap blocks the next phase. If later reconciliation concludes that a public Organization review/verification claim is required, that specific meaning must return to Product/Design human authority rather than being invented in spec, contract, or UI copy.

## Exploration completion state

Stage 3 has selected a coherent direction, recorded material rejected alternatives, identified consequences and risk evidence, and derived a downstream delivery shape without implementing or writing a detailed Techplan.
