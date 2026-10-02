# Tech Plan: Reconciliation Kontrak Campaign Donation Entry Slice 2

> Phase             : Techplan
> Ticket            : WU-S2-005
> Work Unit         : WU-S2-005
> Run               : TP-S2-005-002
> Author            : P-S2-005-TP-002-1 (Planner)
> Participant ID    : P-S2-005-TP-002-1
> Profile           : KC-PLANNER
> Role              : Planner
> Model             : gpt-6-luna (Invocation; runtime metadata tidak exposed)
> Reasoning         : high (Invocation; runtime metadata tidak exposed)
> Session           : tidak exposed
> Created           : 2026-10-01
> Updated           : 2026-10-01
> Target revision   : `7fd8b473b239b20bda3990ab29c51440d321a796` plus working tree saat dispatch
> Workflow revision : `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`
> Status            : Approved
> Approach          : Pertahankan availability-only snapshot yang telah dipilih owner, lalu putuskan schema kondisional dan satu-satunya representasi unavailable yang benar-benar didukung dalam batas public Slice 2.
> Refs              : `WU-S2-005/manifest.md`; predecessor `TP-S2-005-001/techplan.md`; evidence `EXP-S2-005-001/evidence/stage-2-gap-analysis.md` dan `stage-3-solutioning.md`; owner decision di parent `events.md`; Product Slice 2; Campaign dan Donation authority.

---

## 1. Background

Slice 2 mensyaratkan pengunjung dapat memulai guest Donation dari Public Campaign Detail yang eligible. Campaign Detail spec, authored OpenAPI, generated TypeScript, fixture, backend mapper, dan frontend view masih mencerminkan Slice 1: `donation_action` selalu `unavailable` dengan `donation_flow_not_available`. Campaign/API owner telah memilih availability-only snapshot: frontend menentukan route dari Campaign ID dan Donation POST menilai eligibility lagi. Run ini menyelesaikan proposal bentuk wire dan pemetaan public `unavailable`; perubahan source dan implementasi tetap menunggu gate berikutnya.

Predicate visibilitas detail publik saat ini hanya menerima Campaign internal `published`. Slice 1 mengembalikan 404 yang sama untuk Campaign non-public; Slice 3 memiliki perubahan terpisah untuk closed Campaign yang sebelumnya public. Karena itu, tidak boleh mengisi enum dengan status internal atau mengasumsikan closed detail menjadi public dalam WU ini. Source saat ini belum menyebut kondisi detail yang tetap public tetapi donation-ineligible, atau predicate GET-time yang persis sama dengan predicate submission. Ini adalah batas keputusan Campaign/API yang diajukan secara eksplisit di bawah.

## 2. Scope

**In scope:**

- Rekonsiliasi acceptance Campaign Detail dan `INV-campaign-14` secukupnya untuk availability action Slice 2, tanpa mengubah visibilitas publik yang belum direkonsiliasi.
- Menetapkan proposal schema kondisional, finite reason vocabulary, dan skenario public `available` / `unavailable` / 404 / 503 agar owner dapat memutuskan secara konkret.
- Rekonsiliasi authored `api/openapi/campaign.yaml` dan regenerasi counterpart yang diwajibkan `api/README.md` setelah owner contract decision.
- Mempertahankan closed public projection, anti-enumeration, no-store, backend-owned eligibility, dan Donation POST/D1 authority.

**Out of scope (explicit):**

- Mengubah Product/MVP, Donation submission contract/invariants, `INV-campaign-13`, protected ledger/locking logic, atau root fencing.
- Membuat Campaign internal state non-public menjadi public, menambah persistence, mengubah closed-Campaign visibility/public result, atau mendefinisikan Slice 3.
- Backend/frontend production behavior, migrasi, runtime proof, test execution, dan implementasi Donation flow; semua mengikuti Work Unit/gate terkait.
- Mengubah route navigasi Campaign API, endpoint/request, common component, atau operasi Campaign lain.
- Menyatakan schema final, `CONTRACT_READY`, WU selesai, residual-risk acceptance, atau izin Build sebelum owner/review/approval yang berlaku.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | Pengunjung tanpa Account dapat masuk ke guest Donation dari eligible Public Campaign Detail sesuai Slice 2; jangan tampilkan action aktif sebelum real flow tersedia. | `docs/product/mvp-delivery-slices.md` §2.4, §5; `docs/product/mvp-scope.md` §5; Stage-2 Gap Analysis Area 1 |
| Q2 | `donation_action` hanya menyatakan backend-authored availability saat GET; tidak mengandung target route. Frontend memakai Campaign ID untuk route lokal. | Owner decision 2026-10-01 di `.harscode-spaces/s2-guest-donation-truthful-state/events.md`; Stage-3 Solutioning |
| Q3 | `available` adalah snapshot GET, bukan otorisasi/garansi POST. Submission menilai ulang current eligibility dan dapat ditolak bila state berubah. | Product Slice 2 correctness/security floor; `INV-campaign-13`; `INV-donation-02`; `api/openapi/donation.yaml` submit operation |
| Q4 | Public action tidak memperluas public detail predicate, projection, anti-enumeration, atau no-store; tidak memaparkan raw lifecycle/close/operational data. | `INV-campaign-14`; Campaign feature `02-campaign-detail-listing.md`; Stage-2 Gap Analysis Area 2 |
| Q5 | Authoritative Campaign API source adalah `api/openapi/campaign.yaml`; bundle, generated TypeScript, dan fixture mengikuti sumber authored. | `api/README.md`; Stage-2 Gap Analysis Area 3 |
| Q6 | Scope hanya contract/spec untuk Slice 2; Product, D1, Donation submission, closed public continuity/result Slice 3, dan backend/frontend implementation tidak diredefinisi. | `WU-S2-005/manifest.md`; Product Slice 2/3; owner decision history |

## 4. Rules & Validation

- **R1 — Availability-only wire shape (schema proposal dipilih owner; final authored acceptance masih pending):** object `donation_action` memiliki tepat satu varian: `{availability: available}` atau `{availability: unavailable, reason: campaign_not_eligible}`. Tidak ada `reason` pada `available`; `reason` wajib dan finite hanya pada `unavailable`. Tidak ada URL, URI, route, activation target, null sentinel, atau unknown extra field. Perubahan authored contract tetap menunggu review dan gate penerimaan akhir.
- **R2 — GET-time truth:** `available` hanya berarti Campaign backend menilai memenuhi predicate Donation eligibility pada saat menyusun GET. Nilai tidak mengikat state setelah response dan tidak dapat mengotorisasi POST. POST mengevaluasi ulang sesuai Donation contract dan D1.
- **R3 — Finite public unavailable mapping (mapping generik dipilih owner):** `campaign_not_eligible` adalah satu-satunya reason Slice-2. Ia hanya dipakai bila resource tetap lolos predicate public detail tetapi gagal predicate Donation eligibility pada GET. Ia tidak menjelaskan apakah penyebabnya deadline, threshold, atau kondisi internal lain. Sumber live saat ini belum menentukan apakah/kapankah kondisi public semacam itu dapat terjadi; jangan mengarang predicate atau membuat kondisi baru untuk memakai enum.
- **R4 — Public visibility and error mapping:** Campaign absent, malformed/non-resolvable, draft, pending, approved-but-not-published, scheduled, rejected, unpublished/retracted, dan closed tetap respons `PublicCampaignNotFound` 404 yang sama sesuai Slice-1 boundary; `donation_action` tidak ada pada error. Jangan mengubah optional Authorization parity. Kegagalan dependency ketika detail public tidak dapat dilayani tetap `PublicCampaignUnavailable` 503, bukan `donation_action.unavailable`.
- **R5 — Projection and cache:** successful detail memakai closed `PublicCampaignDetail` allowlist yang sama dan success/public 404/503 tetap `Cache-Control: private, no-store`. Action tidak menambah field internal.
- **R6 — Submission authority:** stale `available` GET tidak mengubah POST rejection behavior atau D1 close/submission ordering; Donation submission tetap otoritatif.
- **R7 — Authored/generated correspondence:** setelah schema diterima, `api/openapi/campaign.yaml`, `api/openapi.yaml`, `frontend/lib/api/generated/openapi.ts`, fixture public Campaign, dan Campaign acceptance menyatakan varian yang sama. Generated artifacts dibuat lewat command repo, bukan edit manual.
- **R8 — Scope fidelity:** rekonsiliasi tidak mengubah Product/MVP, D1, Donation API, visibility/closed result Slice 3, atau production behavior; perubahan setiap batas tersebut memerlukan owning authority/Work Unit.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Availability-only projection; frontend membentuk route dari Campaign ID; POST melakukan eligibility recheck. | Chosen — Anhar Solehudin, Campaign/API owner, 2026-10-01; parent `events.md` | Menjaga route sebagai concern consumer dan GET sebagai snapshot. Opsi backend-provided target ditolak; tidak dibuka ulang tanpa evidence authority baru. |
| D2 | Pertahankan visibility Slice 1 selama reconciliation ini; non-public Campaign tetap 404; detail closed/public-result menunggu Slice 3. | Chosen — Campaign/API owner, 2026-10-01 | Owner menyetujui skenario/mapping yang diajukan: unavailable hanya mungkin untuk detail yang tetap public tetapi gagal predicate submit saat GET. Tidak mengubah visibilitas atau menambah kondisi Campaign baru. |
| D3 | Gunakan union kondisional: `available` tanpa reason, atau `unavailable` dengan reason generik `campaign_not_eligible`. | Chosen — Anhar Solehudin, Campaign/API owner, 2026-10-01; jawaban eksplisit atas pertanyaan Participant Session | Menghindari reason required palsu pada available serta menghindari enum untuk tiap status internal. Client dapat membedakan action unavailable tanpa mengetahui penyebab internal. Hanya keluarkan `unavailable` ketika detail masih public dan predicate submit saat GET menyatakan tidak eligible. |
| D3-alt | Pertahankan `donation_flow_not_available` sebagai reason dalam response schema. | Rejected — owner memilih usulan D3 | Reason lama menyatakan flow Slice 2 tidak tersedia, bukan kondisi Campaign tidak eligible. Tidak dipertahankan sebagai compatibility-only response karena tidak ada kondisi producer aktif yang boleh mengeluarkannya dan owner menerima schema usulan tanpa nilai tersebut. |
| D3-alt-2 | Enum reason spesifik seperti `campaign_closed`, `deadline_passed`, `goal_reached`, atau raw status. | Rejected | Source tidak memetakan alasan itu ke detail yang masih public; sebagian akan mengubah visibility/closure Slice 3 atau memaparkan lifecycle. Jangan menambahkan sekadar untuk mengisi enum. |
| D4 | `api/openapi/campaign.yaml` authored; aggregate/type/fixture adalah generated/consumer counterparts. | Chosen — project API source rule | `api/README.md` menetapkan sumber authored dan workflow generation. |
| D5 | Tidak mengubah Campaign close/submission invariant atau Donation POST behavior. | Chosen — authority preservation | `INV-campaign-13` dan `INV-donation-02` mengatur D1; GET action contract hanya entry affordance. |

## 6. Backward Compatibility

- Existing response mewajibkan `donation_action.availability=unavailable` dan `reason=donation_flow_not_available`. Owner memilih schema dengan `available` tanpa `reason`, atau `unavailable` dengan `reason=campaign_not_eligible`; nilai lama dihapus dari active response schema dan tidak dipertahankan sebagai compatibility-only representation. Ini material response contract change yang tetap memerlukan independent review, perubahan authored spec/API, regenerasi types, fixture update, dan final contract acceptance sebelum contract dianggap current.
- Tidak ada bukti dalam input Run yang memastikan consumer eksternal atau distribusi contract Campaign. Jangan mengklaim kompatibilitas eksternal. Known in-repository generated type dan fixture wajib diperbarui setelah decision; downstream WU-S2-004 menunggu hasil accepted.
- Owner telah memutuskan untuk tidak menyertakan `donation_flow_not_available` dalam schema aktif. Tidak ada scenario Slice 2 yang mendukung nilai tersebut; strict consumers harus mengikuti generated contract terbaru. Tidak ada bukti distribusi/consumer eksternal di input Run, sehingga kompatibilitas eksternal tidak diklaim.
- Endpoint, request, auth, response status visibility, shared `common.yaml`, dan Donation POST contract tidak berubah.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| RISK-1 | State berubah sesudah GET `available` dan sebelum POST. | Medium | High | Jelaskan snapshot; POST rechecks dan dapat menolak; runtime behavior dimiliki Campaign/Donation delivery dan Testing. Tidak ada residual-risk acceptance dari Run ini. |
| RISK-2 | Reason enum atau mapping menyingkap internal lifecycle/close reason, atau menambah visibility diam-diam. | Medium | High | Usulan generic `campaign_not_eligible`; hanya detail yang tetap public; statuses closed/non-public tetap 404. Owner menetapkan final mapping. |
| RISK-3 | Tidak ada keadaan saat ini yang dibuktikan dapat menghasilkan public detail + unavailable, sehingga schema memiliki varian yang belum punya runtime scenario. | Medium | Medium | Jadikan exact eligibility/public intersection sebagai bounded owner decision. Bila intersection kosong untuk Slice 2, owner dapat memilih schema `available` saja untuk current scope atau mengidentifikasi kondisi yang didukung authority; jangan menciptakan satu. |
| RISK-4 | `reason` berubah requiredness/nilai dan menyebabkan strict client mismatch. | Medium | Medium | Tandai sebagai material response change; audit consumer repo, regenerate generated types, dan jangan mengklaim kompatibilitas konsumen yang tidak diketahui. |
| RISK-5 | Authored source diperbarui tanpa bundle/type/fixture atau Campaign acceptance yang sinkron. | Medium | Medium | Source-to-generated validation dan correspondence review pada R7. |
| RISK-6 | Closed Campaign dibuat public atau action state dipakai untuk menutupi closure/result behavior. | Low | High | Tutup Slice 3 dari scope; any visibility continuity changes dikembalikan ke owning Slice 3 authority. |

## 8. Interface Contract

**Persistence/data shape:** Tidak ada storage, migration, atau monetary shape yang berubah.

**API/event/external interface:** Endpoint tetap `GET /campaigns/{campaignId}`; `PublicCampaignDetail.donation_action` tetap required. Owner memilih schema OpenAPI 3.0.3 sebagai `oneOf` dengan dua exact object variants (keduanya `additionalProperties: false`); authored contract acceptance tetap pending:

```yaml
oneOf:
  - type: object
    additionalProperties: false
    required: [availability]
    properties:
      availability:
        type: string
        enum: [available]
  - type: object
    additionalProperties: false
    required: [availability, reason]
    properties:
      availability:
        type: string
        enum: [unavailable]
      reason:
        type: string
        enum: [campaign_not_eligible]
```

Conditional field rules: exactly one variant; `reason` absent for `available`, required for `unavailable`; no `null`, target, or extra property. Ini owner-selected contract proposal untuk Techplan, belum menjadi accepted authored contract sampai spec/API review dan acceptance selesai.

| GET-time scenario | Proposed observable result | Scope / authority |
|---|---|---|
| Resource is public and Campaign passes backend Donation-eligibility predicate at GET | `200`, `donation_action: {availability: available}` | Slice 2 in scope. Eligibility is assessed at read time; frontend chooses route from Campaign ID. |
| Resource remains public, but Campaign fails the backend submission-eligibility predicate at GET | `200`, `donation_action: {availability: unavailable, reason: campaign_not_eligible}` | Owner-approved generic mapping. Current source does not establish whether this public/eligibility intersection occurs or expose an explicit shared evaluator. Emit only if the existing authoritative predicate returns ineligible; do not manufacture a condition if no such state exists. |
| Campaign is absent, malformed/non-resolvable, or non-public, including current closed state | Existing indistinguishable `404 PublicCampaignNotFound`; no detail/action body | Existing Slice 1 boundary; preserve anti-enumeration/Authorization parity. Closed public continuity/result belongs to Slice 3. |
| Eligible public Campaign dependency cannot serve detail | Existing `503 PublicCampaignUnavailable`; no synthetic unavailable action | Existing public API failure distinction. |
| State becomes ineligible after a `200 available` GET | Later POST rechecks and can return existing ineligible `409` | Existing Donation submit authority/D1; not a new Campaign GET reason. |

Existing successful/error response no-store header and closed allowlist remain. Do not add a response reason for upstream failure.

**Cross-layer/business boundary:** backend Campaign GET reports a point-in-time action affordance; frontend maps Campaign ID to its own route and treats unavailable as non-actionable; Donation POST independently owns current eligibility. Never derive submission authorization from public visibility or response freshness.

## 9. Architecture / Plan

1. Owner decision D2/D3 is recorded: use the exact two-variant schema in §8; `campaign_not_eligible` is the sole reason and only applies when detail remains public but fails the backend submission-eligibility predicate at GET. Never emit the old `donation_flow_not_available` reason.
2. Reconcile Campaign Detail acceptance and `INV-campaign-14` only as needed. The producer must evaluate the same eligibility predicate that governs submission at the GET snapshot; do not add a condition or infer a reason from raw lifecycle fields. If no state can pass public detail while failing that predicate, do not emit `unavailable` at runtime or manufacture a condition; this does not reopen the owner-selected wire schema. If live authority/source cannot evaluate the predicate without inventing a rule, stop and route that exact gap to Campaign authority.
3. Preserve the existing public predicate, anti-enumeration/projection/no-store requirements, `INV-campaign-13`, and Donation `INV-donation-02`. Non-public/closed resources remain 404; dependency failure remains 503.
4. Reconcile authored `api/openapi/campaign.yaml`; use shared `common.yaml` only if a genuinely shared component is necessary. Preserve endpoint, errors, headers, and security declaration.
5. Bundle and regenerate generated TypeScript per `api/README.md`, update the public Campaign fixture, and inspect all known consumers. Never hand-edit generated contract outputs.
6. Validate authored schema and source/generated correspondence; complete independent review and final owner acceptance through their gates. Route backend GET producer and frontend implementation to their own authorized Work Units after contract convergence.

No independently operated migration, script, cron, or runbook lifecycle is introduced.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `docs/spec/4-campaign/features/02-campaign-detail-listing.md` — `Detail`, acceptance criteria 3 and 6 | Slice 1 currently mandates unavailable action | Reconcile for Slice 2 only after exact owner schema/mapping; retain eligible public projection and no fake action boundary. |
| `docs/spec/4-campaign/invariants.md` — `INV-campaign-14` | Owns public visibility, closed projection, anti-enumeration, action and cache boundary | Amend only action wording necessary for accepted Slice 2; preserve current visibility and other checks. |
| `docs/spec/4-campaign/invariants.md` — `INV-campaign-13` | Owns close/Donation D1 ordering | Reference unchanged; do not edit its protected meaning or mechanism. |
| `docs/spec/5-donation/invariants.md` — `INV-donation-02`; `api/openapi/donation.yaml` — `POST /campaigns/{campaignId}/donations` | Current submission authority and ineligible response | Preserve as downstream current-state recheck; no Donation contract edit. |
| `api/openapi/campaign.yaml` — `getPublicCampaignDetail`, `PublicCampaignDetail`, `PublicCampaignDonationAction` | Authored Campaign response contract; current shape requires only unavailable + old reason | Replace only after owner decision; proposal is exact conditional union in §8. |
| `api/README.md` — `Editing workflow` | Owns split source / bundle / type generation workflow | Validate source, regenerate `api/openapi.yaml` and generated types; no manual bundle/type edits. |
| `backend/internal/domain/campaign/repository_db.go` — `RepositoryDB.FindPublicDetail` | Current GET selects only `status=published`, explicit projection | Later backend producer owner reopens; decide whether its source can evaluate exact GET-time Donation eligibility without broadening visibility. No edit in WU-S2-005. |
| `backend/internal/domain/campaign/service.go` — `Service.GetPublicDetail`, `toPublicDetail` | Current mapper hardcodes unavailable + `donation_flow_not_available` | Later producer mapping follows accepted schema/predicate. This plan does not implement it. |
| `backend/internal/transport/http/campaign_public.go` — `PublicCampaignDetailHandler`, `toPublicCampaignDetailResponse` | Current public DTO serializes availability/reason | Later producer preserves closed DTO and errors after contract acceptance. |
| `frontend/lib/api/generated/openapi.ts` — `PublicCampaignDonationAction` | Generated consumer type | Regenerate from accepted authored source; do not hand-edit. |
| `frontend/mocks/fixtures/public-campaign.ts` — `donation_action` | Contract fixture currently carries old unavailable value | Update all needed fixture variants to accepted schema; route UI work remains frontend-owned. |
| `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` — `CampaignSuccess` | Current view has static unavailable copy and does not inspect action | Downstream frontend owner derives local route from Campaign ID after accepted contract; no frontend write in this WU. |

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/TP-S2-005-002/techplan.md` | Draft successor | Complete schema/mapping proposal; owner decision status and phase history retained. |
| `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/TP-S2-005-002/handoff.md` | Phase handoff | Provenance, status, evidence boundary, and recommended next route. |
| Campaign feature/invariants and Campaign authored/generated API surfaces | Planned later | Only after owner decision, review, and applicable approval; see §§8–10. |
| Product/MVP, Donation spec/API and `INV-donation-02`, `INV-campaign-13`, backend/frontend production, generated tests, migrations, protected ledger/locking/auth/crypto | Intentionally untouched | Outside reconciliation scope, restricted by authority, or implementation/testing phase ownership. |
| Prior EXP and TP-001 evidence; parent Events/Control Surface projections | Intentionally untouched by Planner | Preserve predecessor history and orchestration ownership. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | Human/API owner reviews exact OpenAPI `oneOf`, requiredness, finite value, and absence of target/extra properties. | Human | Response shape and conditional field behavior are a material consumer contract; owner interpretation must be settled before source acceptance. |
| R1 | After owner acceptance, validate the authored OpenAPI schema. | Testing | Invalid conditional schema can leave generated consumers inconsistent or allow impossible response combinations. |
| R2 | Review GET snapshot → eligibility changes → POST recheck scenario against Campaign acceptance and Donation POST; later delivery Testing proves observable rejection/ordering. | Testing | Prevents `available` being treated as authorization; skip risks accepting after close based on stale client state. |
| R3 | For every emitted unavailable scenario, trace to an approved predicate; inspect public response for generic reason only and no raw status/close/operational fields. | Human (mapping decision); Testing (contract/projection) | Prevents a fabricated enum, internal disclosure, or action state with no supported runtime meaning. |
| R4 | Contract scenario review checks the current non-public ID matrix remains one 404; public dependency failure remains 503; no action body or auth-dependent variation. Later backend Testing owns status/body/header/timing proof. | Testing | Protects anti-enumeration and distinguishes absent visibility from service failure. |
| R5 | Validate closed projection and `private, no-store` for success/404/503 in spec and downstream runtime evidence. | Testing | Projection/cache regressions can expose internal data or stale withdrawn content. |
| R6 | Trace Campaign action acceptance and API text to `INV-campaign-13`, `INV-donation-02`, and Donation POST; later D1 runtime suite remains in WU-S2-003. | Testing | Ensures frontend affordance does not weaken authoritative close/submission ordering. |
| R7 | Run repository-authored OpenAPI validation, bundle and type generation, and update the fixture in the edit loop. | Build | These artifacts must remain executable and aligned as they are authored; otherwise Build feedback is against stale generated types. |
| R7 | Independently compare generated `PublicCampaignDonationAction`, fixture, and known consumers to the authored schema. | Testing | Generated drift causes consumers to compile/use a different contract than the authored source. |
| R8 | Review final diff against WU manifest, Product Slice 2/3, D1, and root fencing; owner confirms protected spec/API changes at applicable gates. | Human | Scope/authority drift would silently alter product or a separate delivery contract. |

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Public projection, non-public Campaign anti-enumeration, optional Authorization parity, and no-store | Contract reason/state must not widen disclosure or distinguish non-public IDs. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001/evidence/stage-2-gap-analysis.md#area-2-campaign-acceptance-criteria-dan-invariants`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001/evidence/stage-2-gap-analysis.md#area-3-authored-dan-generated-api-contract-surfaces` | Yes — specialized backend Testing remains required; this Run has no runtime proof. |
| Stale GET availability versus later Donation submission | Client-held state can be stale after GET and must not bypass POST eligibility. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001/evidence/stage-2-gap-analysis.md#area-4-perilaku-backendfrontend-yang-terkait`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001/evidence/stage-3-solutioning.md#batas-keputusan-dan-risiko` | Yes — runtime recheck/rejection evidence belongs to Campaign/Donation delivery Testing. |
| Concurrent close/submission/settlement D1 ordering | Money/lifecycle concurrency is material but action contract does not change D1. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001/evidence/stage-2-gap-analysis.md#area-2-campaign-acceptance-criteria-dan-invariants`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001/evidence/stage-2-gap-analysis.md#area-4-perilaku-backendfrontend-yang-terkait` | N/A for execution here — unchanged and owned by WU-S2-003; its verification remains required there (Q3, R6, D5). |

## 13. Open Items

### Active — needs external input or verification

1. **Implementation source fidelity for the accepted predicate — verify during authorized Build/reconciliation.** The owner accepted the generic mapping, but current public-detail live query filters only `status=published` and current Campaign GET mapper has no explicit submission-eligibility evaluator. The authorized implementation owner must identify/reuse the authoritative eligibility predicate from current Campaign/Donation code/spec; if it cannot be derived without adding policy, stop and route the exact gap to Campaign authority. Do not turn internal deadline/threshold/close labels into public reasons. This does not reopen the accepted wire schema.
2. **Protected authority reconciliation and final contract acceptance.** Campaign spec/API changes need the applicable owner acceptance after independent review; this Draft does not edit/approve those protected sources or claim final contract acceptance. Owner: Anhar Solehudin, Campaign/API owner.
3. **Independent Techplan review.** Currently parked at `RV-S2-005-001`; dispatch now that the schema proposal semantics converged. Review exact proposal and cross-boundary fidelity before report/approval gate. No Review verdict exists.

### Resolved — retained as decision history

1. ~~**Action navigation ownership**~~ **RESOLVED — availability-only; frontend derives route from Campaign ID.** Anhar Solehudin, Campaign/API owner, 2026-10-01; parent `events.md`. Backend-provided target rejected.
2. ~~**Conditional schema and unavailable reason**~~ **RESOLVED — owner accepts two variants: `available` without reason; `unavailable` with only `campaign_not_eligible`.** The unavailable variant applies only if detail remains public while the GET-time backend submission-eligibility predicate fails. Anhar Solehudin, Campaign/API owner, explicitly selected “Setujui usulan (Recommended)” in Participant Session on 2026-10-01. `donation_flow_not_available` is not retained, including compatibility-only. Final authored contract acceptance remains a later gate.
3. ~~**Submission authority**~~ **RESOLVED — POST reevaluates current eligibility.** GET response is a snapshot; existing Campaign/Donation D1 authority remains unchanged.
4. ~~**Product/MVP direction**~~ **RESOLVED — no Product/MVP change indicated.** Slice 2 guest Donation entry is clear; Slice 3 owns closed Campaign public continuity/result.
