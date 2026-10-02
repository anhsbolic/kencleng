# Tech Plan: Slice 2 Donation Limits & Campaign Capacity Source Reconciliation

> Phase             : Techplan
> Ticket            : WU-S2-006 / TP-S2-006-001
> Author            : P-S2-006-TP-001-1 (Planner)
> Participant ID    : P-S2-006-TP-001-1
> Profile           : KC-PLANNER
> Role              : Planner
> Model             : Invocation configures `gpt-6-luna`; active runtime model not independently exposed
> Reasoning         : Invocation configures `medium`; active runtime effort not independently exposed
> Session           : not exposed
> Created           : 2026-10-01
> Target revision   : `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree; source anchors re-read
> Workflow revision : Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`
> Status            : Draft / In Review
> Approach          : Rekonsiliasi Product/MVP, Donation/Campaign spec, authored API dan counterpart turunan untuk limit Donation dan kapasitas Campaign; persetujuan sumber tetap gate terpisah.
> Refs              : WU-S2-006 manifest; EXP-S2-006-001 evidence + handoff; Human clarification in parent `events.md`; TP-S2-003-003 D16/O1-REP; accepted WU-S2-005 contract; accepted Donation baseline WU-S2-002.

---

## 1. Background

RV-S2-003-001 menemukan bahwa Donation yang valid menurut kontrak saat ini belum tentu dapat dicerminkan penuh oleh Campaign `collected_amount NUMERIC(19,2)`. Human memilih batas individual per Campaign, reservasi kapasitas untuk Funding settled dan Donation `pending` yang sudah diterima, serta penutupan Campaign saat kapasitas habis. Human clarification 2026-10-01 juga menetapkan konfigurasi Owner/Staff ketika Campaign masih draft, pembekuan saat publikasi, disclosure cap pada public detail sebelum nominal dimasukkan, dan handoff makna capacity-close ke Slice 3. Pilihan itu adalah arah solusi, belum perubahan Product/spec/API yang diterima.

Kontrak Campaign GET baru saja diterima pada WU-S2-005 sebagai snapshot availability-only dengan POST recheck. WU006 perlu menambah limit disclosure dan menautkan penutupan kapasitas ke pengalaman Slice 3 sambil menjaga batas runtime Slice 2.

## 2. Scope

**In scope:**

- Rekonsiliasi Product/MVP Slice 2 untuk individual Donation cap: whole-IDR Rp5.000–Rp1.000.000.000, default Rp1.000.000.000, Owner/Staff mengatur saat draft, tidak dapat diubah setelah publikasi.
- Rekonsiliasi Donation dan Campaign invariants/feature criteria agar amount cap, kumulatif Funding capacity, reservasi pending, full settlement, close reason, dan `max_amount` tetap aturan berbeda dan konsisten.
- Rekonsiliasi authored Donation/Campaign OpenAPI untuk cap Campaign configuration, pre-submit public disclosure, request limit, capacity closure, dan reason; pertahankan API Campaign WU-S2-005 sebagai baseline yang diterima sampai amendment direview dan diterima.
- Identifikasi dan sinkronisasi required generated/fixture counterparts, termasuk `api/openapi.yaml`, generated TypeScript types, dan public Campaign fixture; jangan edit generated files secara manual.
- Dokumentasikan dependency Slice 3: capacity close memicu persistent public Campaign identity, donation action tidak tersedia, Funding tidak disebut final selama accepted Donation masih pending. Source applicability/dependency ini tidak mengaktifkan closed-detail visibility pada Slice 2 dan tidak memperluas delivery Slice 3.
- Catat source-owner review/acceptance, backend/frontend plan refresh, dan runtime/concurrency evidence sebagai gate/dependency terpisah.

**Out of scope (explicit):**

- Implementasi backend/frontend, schema migration, manual DB/index application, generators atau tests dalam Run ini; implementation/runtime ownership tetap pada delivery Work Units.
- Memilih mekanisme transaction/locking atau mengubah file Tier-0 `backend/internal/domain/donation/ledger.go` maupun transaction/locking code lain.
- Mengubah shared monetary standard menjadi universal range/precision/scale; feature limits tidak dimiliki standar representasi bersama.
- Mengubah validasi, notification/PII, real settlement, atau behavior Donation lain yang tidak diperlukan untuk cap/capacity reconciliation.
- Mengaktifkan persistent public detail untuk Campaign tertutup pada Slice 2, membangun full Slice-3 result experience, atau mengklaim Slice-3 contract/runtime readiness.
- Menyetujui source edits, residual risk, whole-Techplan, protected implementation, maupun delivery milestones melalui Run ini.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | Donation tetap guest-capable, whole Rupiah IDR minimum Rp5.000; tiap Campaign memiliki cap Donation eksplisit/default Rp1.000.000.000 dan dapat diatur Owner/Staff saat draft dalam rentang Rp5.000–Rp1.000.000.000; cap beku setelah publikasi. | Parent `events.md` “Human clarified EXP006 D-01–D-04”; `docs/product/mvp-scope.md` §5 (existing baseline); active source reconciliation required. |
| Q2 | Public Campaign detail menunjukkan cap aktif sebelum guest memasukkan nominal. Campaign GET hanya affordance snapshot; Donation POST memvalidasi cap dan eligibility terkini secara independen. | Parent `events.md` D-02; accepted WU-S2-005 contract and `docs/spec/4-campaign/invariants.md#inv-campaign-14`. |
| Q3 | Cumulative Funding representability ceiling mengikuti Campaign `NUMERIC(19,2)` current whole-IDR capacity `99,999,999,999,999,999`; admission menghitung settled Funding ditambah semua accepted-pending obligations. Capacity exhaustion menutup Campaign dengan distinct close reason, bukan `max_amount_reached`. | `TP-S2-003-003` D16/O1-REP; `RV-S2-003-001/review-findings.md`; current migration anchor. Identifier reason masih proposal. |
| Q4 | Donation yang sudah diterima tetap dapat settle penuh tepat sekali setelah Campaign close; failed Donation tidak dihitung sebagai collected Funding. Failed pending reservation dapat melepaskan ruang tetapi tidak membuka ulang Campaign atau mengganti winning close reason. `max_amount` tetap threshold yang dapat terlampaui oleh Donation accepted. | Accepted `docs/spec/5-donation/invariants.md#inv-donation-02`, `#inv-donation-08`; `docs/spec/4-campaign/invariants.md#inv-campaign-13`; parent `events.md` D-03. |
| Q5 | Capacity close menjadi dependency Slice 3: public identity tetap tersedia, donation action hilang, dan Funding belum final selama accepted Donations masih pending. Slice-2 closed/non-public behavior tetap current 404 sampai rekonsiliasi Slice 3 berlaku. | Parent `events.md` D-04; WU-S2-006 manifest; current `docs/spec/4-campaign/invariants.md#inv-campaign-14`. |
| Q6 | API money tetap major-unit decimal string dengan currency eksplisit; arithmetic/storage exact decimal, tanpa `float`/`float64`. Jangan menetapkan universal precision/scale/range lewat shared monetary standard. | `docs/project/kencleng-monetary-data-standard.md`; `../harscode-workspace/best-practices/go/decimal-and-money.md`; root `AGENTS.md`. |

## 4. Rules & Validation

- **R1 — Configuration boundary:** Given a Campaign draft, when an authorized Owner/Staff configures its Donation cap, then the stored/effective cap is a whole-IDR amount from Rp5.000 through Rp1.000.000.000, default Rp1.000.000.000; after publication it cannot be changed through Slice-2 configuration operations.
- **R2 — Public disclosure and authoritative submit:** Given a public eligible Campaign, when public detail is read, then the active individual cap is disclosed before amount entry; the existing WU-S2-005 `donation_action` semantics remain a GET snapshot, and POST independently rejects a cap/eligibility violation using the accepted generic error contract.
- **R3 — Capacity admission:** Given settled Funding `F`, accepted pending obligations `P`, and Campaign cumulative ceiling `C`, then a new Donation may be accepted only when the resulting reserved amount remains representable (`F + P + amount ≤ C`). If no representable capacity remains for a new Donation, Campaign closes for capacity exhaustion.
- **R4 — Stable close reason and threshold separation:** Given capacity closure or another close trigger, then Campaign retains the winning close reason under existing close-ordering rules. `funding_capacity_reached` is only a proposed identifier until accepted by Campaign/API source owners; capacity is distinct from nullable `max_amount`, which can be overshot by already accepted full-value Donations.
- **R5 — Pending settlement and failure:** Given an accepted pending Donation when Campaign closes, then later success commits full Donation outcome and Funding once without reopening Campaign or changing its winning reason; later failure contributes no Funding and does not reopen Campaign even if reservation is released.
- **R6 — Slice-3 public closure handoff:** Given a Campaign closed by capacity, Slice 3 source behavior must preserve public Campaign identity, remove donation action, and show Funding as non-final while any accepted Donation is pending. This rule is a downstream applicability requirement and does not change current Slice-2 closed/non-public 404 behavior.
- **R7 — Exact money and projections:** Given configured limits, reservations, persisted Funding, API values, generated schemas, and fixtures, then all representations preserve exact whole-IDR values and the accepted cumulative ceiling without binary floating point or silent truncation.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Cap per Campaign; hard maximum Rp1.000.000.000; explicit draft value with default Rp1.000.000.000; Owner/Staff configuration during draft, frozen once published. | Chosen — Human direction, 2026-10-01; source acceptance pending | Current durable exact clarification supersedes the inconsistent five historical Stage-3 briefs. Concrete field/schema and compatibility are for owning-source resolution. |
| D1-alt | Treat Rp1b as only initial direction with values allowed above it. | Rejected | Conflicts with the clarified hard upper bound for Slice 2. |
| D2 | Disclose active cap on public Campaign detail before amount entry; POST rechecks cap/eligibility. | Chosen — Human direction, 2026-10-01; API amendment pending | Supports informed entry without turning GET snapshot into authorization. |
| D2-alt | Reveal cap only after POST rejects an amount over limit. | Rejected | Does not satisfy accepted pre-entry disclosure direction. |
| D3 | Reserve settled Funding plus all accepted `pending` Donations; close once capacity is exhausted; do not reopen after a later failure frees reservation. | Chosen — Human direction, 2026-10-01; source acceptance pending | Ensures every accepted Donation can settle fully and preserves stable close outcome; may leave unused capacity after failure. |
| D3-alt | Leave Campaign open while pending obligations occupy capacity and reopen admissions if one fails. | Rejected | Would make Campaign appear open while no new Donation can be accepted and would permit re-opening after an apparent exhausted-capacity state. |
| D4 | Carry capacity-close public behavior into Slice 3 only; retain identity, remove action, mark Funding non-final while accepted pending remains. | Chosen — Human direction, 2026-10-01; Product/spec Slice-3 source handoff pending | Preserves product truth after closure without silently changing current Slice-2 404 contract or activating full Slice-3 delivery. |
| D5 | Use `funding_capacity_reached` as distinct close-reason identifier. | Active source decision — proposal only | It is present in predecessor Draft, not in accepted D-01–D-04 receipt. Campaign/API owner must accept it or select a different stable identifier before contract completion. |
| D5-alt | Retain capacity closure but select another owner-approved stable identifier. | Open option | Must remain distinct from `max_amount_reached`; exact enum is an owning-source decision, not a Build choice. |
| D6 | Keep shared monetary standard format-only; do not add universal DB precision/scale or platform-wide maximum. | Chosen — existing approved shared authority | Feature-level cap and current Campaign capacity resolve this slice's representability issue without promoting one schema width to universal policy. |
| D7 | Preserve accepted WU-S2-005 `donation_action` schema semantics while adding cap disclosure; do not expose internal capacity/close reason through public action/errors. | Chosen — accepted predecessor contract plus current direction | WU006 must amend only what the new accepted direction requires and seek fresh contract acceptance for the delta. |

## 6. Backward Compatibility

- **Existing data:** Current Campaign rows have no per-Campaign donation cap. Owning source/migration planning must state how existing drafts receive the default and how already-published Campaigns obtain the frozen default without introducing an unset state. No DB backfill/application is authorized by this Run.
- **API/contracts/clients:** `api/openapi/campaign.yaml` currently has a closed `PublicCampaignDetail` allowlist and accepted `donation_action` union. Adding cap disclosure and configuration/request fields changes the authored contract; evaluate required/optional compatibility for existing clients and fixtures. `api/openapi/donation.yaml` submit amount currently lacks the selected feature cap. Existing generic Problem Details behavior and Campaign closed/non-public `404` remain unless the owning source explicitly reconciles a compatible Slice-3 boundary.
- **Migration/deprecation compatibility:** Do not widen Campaign `NUMERIC(19,2)` on assumption; the selected ceiling derives from its whole-IDR representable capacity. Any source proposal to widen storage or change public Funding precision needs consumer evidence and Product/monetary review. Schema migration, rollout/backfill, rollback and manual application belong to later delivery plans and Human gates.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| E1 | Accepted Donation cannot settle fully because reserved plus settled Funding exceeds Campaign column capacity. | Medium | Critical | Reconcile admission/reservation semantics across Product/spec/API; downstream DB concurrency/integration proof must show no overflow. |
| E2 | Concurrent admissions both observe remaining capacity and over-reserve it. | Medium | Critical | Preserve as specialized Testing obligation; implementation mechanism remains delivery-owned and Tier-0 protected where applicable. |
| E3 | A failed pending Donation frees capacity after capacity closure, tempting automatic reopen. | Medium | High | Stable winning-close rule keeps Campaign closed; unused capacity is an explicit consequence, not hidden recovery logic. |
| E4 | Capacity close is confused with `max_amount` threshold or exposes internal reason/cumulative remaining capacity publicly. | Medium | High | Keep separate rule/enum and generic public behavior; review all projections/errors and Slice-3 handoff. |
| E5 | Cap is missing, below minimum, mutable after publication, or inconsistent between config and submit validation. | Medium | High | Default and range are explicit; source contract must define persisted value and freeze boundary, then producer/POST delivery proves parity. |
| E6 | Slice-2 API starts returning closed Campaign details because Slice-3 behavior is mentioned in current reconciliation. | Low | High | Explicit applicability split in Q5/R6; current 404 remains until a Slice-3 source/contract change is accepted. |
| E7 | Amount or Funding is rounded/truncated or cannot fit API/generated/client representation at the ceiling. | Medium | High | Exact decimal representation, range-boundary contract fixtures and end-to-end projection checks; no floats. |
| E8 | Proposed `funding_capacity_reached` becomes treated as accepted enum before owner decision. | Medium | High | Keep it marked proposal; resolve through Campaign/API owner and independent source review before source completion. |

## 8. Interface Contract

**Persistence/data shape:** Existing Campaign funding is `NUMERIC(19,2)` (`backend/migrations/000011_create_public_campaigns.up.sql#campaigns`) with maximum whole-IDR amount `99,999,999,999,999,999`. Reconciliation must define a persisted per-Campaign cap with default Rp1.000.000.000, allowed whole-IDR range Rp5.000–Rp1.000.000.000, and draft-only mutability. Exact column name/type/default/backfill and migration order are not selected here; source authorship must document existing-row treatment before delivery planning.

**API/event/external interface:** Reconcile the accepted public Campaign detail allowlist to disclose the active cap before guest amount entry. Reconcile Campaign create/update configuration contract for authorized Owner/Staff draft setup and freeze after publication; reconcile Donation submit validation to the current Campaign cap. Preserve WU-S2-005 action semantics (`available` without reason; `unavailable` only with `campaign_not_eligible` while detail is public), POST recheck, and generic public errors. Campaign `ClosedReason` must gain a distinct owner-selected capacity reason; `funding_capacity_reached` remains unaccepted pending explicit source resolution. Do not expose cumulative reserve/remaining capacity or capacity reason publicly. Carry the Slice-3 closure result obligation in its owning Product/spec/API source without changing Slice-2 closed detail behavior.

**Cross-layer/business boundary:** Campaign owns cap configuration, lifecycle and winning close reason. Donation owns accepted Donation status and full-value contribution. Admission must atomically respect Campaign capacity reservation and close; settlement of an accepted pending Donation commits its full amount exactly once even after close. GET action/cap data is a snapshot; POST is authoritative. No locking/transaction mechanism is prescribed by this reconciliation plan.

## 9. Architecture / Plan

1. Source-owning reconciliation authors the Product/MVP behavior and decides any remaining exact configuration/wire/schema details under the accepted D-01–D-04 direction.
2. Reconcile Campaign and Donation invariants/feature acceptance to state the cap range/default/freeze, capacity reservation, capacity close/reason, unchanged threshold behavior, full pending settlement, and Slice-3 public closure handoff.
3. Reconcile split authored Campaign and Donation OpenAPI and counterpart requirements. Keep GET disclosure and POST validation separate; keep current Slice-2 404 behavior for closed/non-public Campaigns.
4. Independently review changed owning sources and required contract consistency; obtain Human/owner acceptance before treating WU006 completion condition as met.
5. Only after accepted source convergence, refresh backend WU-S2-003 and frontend WU-S2-004 Techplans in separate fresh Runs; their implementation/runtime gates remain independent.

No independently operable migration/runbook is planned in this source-reconciliation Run. Migration rollout and DB application belong to delivery planning once source decisions are accepted.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `docs/product/mvp-scope.md` §§5–7, 12–13 | Owns approved MVP guest behavior, closure outcome, exclusions | Add accepted feature cap/default/config/freeze and capacity semantics; preserve threshold semantics and Slice-3 public-result boundary. |
| `docs/product/mvp-delivery-slices.md` §§3, 5–6 | Owns slice sequence and boundaries | Record limit disclosure in Slice 2 and explicit capacity-close handoff to Slice 3 without expanding full Slice-3 delivery. |
| `docs/project/kencleng-monetary-data-standard.md` “Approved representation direction” / “Deliberately unresolved” | Owns shared representation, intentionally not feature limits | Keep as is unless concrete violation of shared exact-decimal format appears; do not add global precision/range. |
| `docs/spec/4-campaign/invariants.md` `INV-campaign-02`, `INV-campaign-13`, `INV-campaign-14`; features `01-campaign-creation-draft-crud.md`, `02-campaign-detail-listing.md`, `09-closure.md` | Campaign owns config validation, lifecycle/close reason, public projection and closure behavior | Reconcile config/freeze, capacity close, cap projection and Slice-3 public pending truth while preserving current Slice-2 applicability. |
| `docs/spec/5-donation/invariants.md` `INV-donation-01`, `INV-donation-02`, `INV-donation-08`; feature `01-submit-donation-settlement.md` | Donation owns amount acceptance and full exact-once settlement boundary | Add per-Campaign cap and capacity admission coordination without weakening accepted-pending settlement, request idempotency, or threshold overshoot. |
| `api/openapi/campaign.yaml` `ClosedReason`, `CampaignCreateRequest`, `CampaignUpdateRequest`, `PublicCampaignDetail`, `PublicCampaignFundingAvailable` | Authored Campaign API currently has no selected cap field/reason and action contract is accepted from WU005 | Owner reconcile config, detail disclosure and reason enum; ensure public allowlist/privacy and decimal range remain exact. |
| `api/openapi/donation.yaml` `SubmitDonationRequest` | Current request schema has no selected cap | Describe the active Campaign-specific max and accepted validation failure shape through API owner. |
| `api/openapi/index.yaml`, `api/openapi/common.yaml`, `api/README.md` | Route and shared Problem/money contract authority | Update path index only if operations change; use shared components as appropriate; validate split sources and regenerate bundle per documented workflow. |
| `backend/migrations/000011_create_public_campaigns.up.sql` `campaigns` | Current `collected_amount NUMERIC(19,2)` establishes cumulative representability ceiling | Read-only precedent for this Run; later delivery chooses new migration/backfill under accepted sources. |
| `backend/internal/domain/campaign/entity.go` `DetailRecord` / `DonationAction`; `service.go` `toPublicDetail`; `backend/internal/transport/http/campaign_public.go` `toPublicCampaignDetailResponse` | Live GET producer path for accepted action snapshot | Later backend plan must include cap projection and authoritative eligibility parity; do not change code here. |
| `frontend/lib/api/generated/openapi.ts`; `frontend/mocks/fixtures/public-campaign.ts`; `frontend/lib/api/public-campaign.ts` `getPublicCampaignDetail`; `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` | Required generated/fixture and consumer correspondence | Regenerate types, update fixtures, then refresh frontend plan for presentation/amount validation; no frontend production edits here. |
| `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-003/techplan.md` D16/O1-REP and `../RV-S2-003-001/review-findings.md` | Exact financial representability problem and predecessor proposal | Preserve D16 direction; supersede its unstated configuration details only with current exact Human clarification and accepted source outcome. |
| `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/manifest.md` and `runs/TST-S2-005-001/handoff.md` | Accepted predecessor availability-only action contract and runtime follow-up | Preserve accepted contract until explicit reviewed amendment; its producer/runtime evidence remains downstream. |

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-001/techplan.md` | New, Run-local | This execution plan. |
| `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-001/handoff.md` | New, Run-local | Phase completion, gates and next routing. |

| File / area intentionally untouched | Why |
|---|---|
| Product/MVP, monetary standard, Donation/Campaign specs, split/generated API, fixtures and production code | Invocation reserves all source authoring/acceptance for later owning-source route; this Run writes only Techplan/handoff. |
| Prior Exploration/Techplan/handoffs, Work Unit records, events and projections | Immutable historical/current orchestration inputs; Run references them without changing their state. |
| Tier-0 transaction/locking, disbursement state machine, crypto/auth files | Root `AGENTS.md` fencing; no implementation authorization in this Run. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | Product/spec/API source review confirms cap range/default, draft Owner/Staff config and post-publication immutability align; later backend tests prove config validation/freeze and row backfill. | Human for policy/source acceptance; Testing for runtime | Wrong or mutable cap changes donor contract; missing values make eligibility ambiguous. |
| R2 | OpenAPI/schema/generated/fixture correspondence check plus later public GET contract test for cap/action and POST boundary test for recheck/generic over-cap error. | Testing | Stale generated/client contract misleads donors; trusting GET as authorization permits stale eligibility. |
| R3 | PostgreSQL integration/concurrency tests assert reserved settled+pending never exceeds ceiling at boundaries and no accepted Donation becomes un-settleable. | Testing | Unit/fake tests cannot establish transaction ordering, real constraints, or concurrent reservation safety. |
| R4 | Contract/domain tests assert capacity reason distinct from threshold and first winning close reason remains stable across competing triggers. | Testing | Conflating the two limits or replacing the reason gives incorrect lifecycle truth; identifier must first be source-accepted. |
| R5 | PostgreSQL integration/replay tests cover accepted-pending success after close, full exact-once contribution, failure releasing reserve without reopen, rollback and stable reason. | Testing | Financial loss/overflow or reopening after a terminal capacity decision is critical; DB atomicity requires real runtime evidence. |
| R6 | Human Product/spec review checks Slice-3 wording; later rendered/contract acceptance verifies identity/action/funding pending truth when Slice 3 is delivered. | Human | “Final” Funding or lost public identity while accepted Donations remain pending misleads donors; automation cannot decide product meaning. |
| R7 | Split OpenAPI validation, bundle/type generation correspondence, boundary-value fixtures/tests, and exact-decimal persistence/projection tests. | Testing | Precision/scale drift can silently truncate or fail settlement; generated consumers must match authored sources. |

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Concurrent Campaign capacity reservation, Donation acceptance, settlement/funding atomicity and close ordering | Interleavings can over-reserve representability, accept an un-settleable Donation, lose/duplicate Funding, or replace winning close reason; remains downstream even though D-03 direction is settled. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/EXP-S2-006-001/evidence/stage-2-gap-analysis.md#area-2--donation-and-campaign-domain-specifications`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/EXP-S2-006-001/evidence/stage-2-gap-analysis.md#f-05--capacity-and-threshold-semantics-must-remain-distinct-under-d1` | Yes — runtime mechanism and proof are not established by source reconciliation. |

## 13. Open Items

### Active — needs external input or verification

1. **OI-1 — Exact authored contract shape and compatibility.** Current accepted D-01–D-04 specify business direction but not API field names/placement, Campaign configuration operation fields, exact over-cap Problem type/field errors, or optionality/backfill compatibility for existing clients/rows. Owner: Anhar Solehudin (scoped Slice-2 Product/MVP, Campaign/Donation and API authority; `.harscode-spaces/authority-map.md`). Resolve during owning-source reconciliation; include evidence/options and accepted choices in source review. Recommendation: expose one active Campaign cap in the public detail allowlist, require a persisted effective value for every published Campaign, and avoid adding public cumulative-capacity/reason details. These decisions must be durable before Build; Build may not infer them.
2. **OI-2 — Capacity close enum.** `funding_capacity_reached` is not included in accepted D-01–D-04 clarification and remains a proposal from TP-S2-003-003. Campaign/API owner must accept that stable identifier or select a distinct alternative before API/spec acceptance. Owner: Anhar; evidence: predecessor D16/O1-REP, current `ClosedReason` schema.
3. **OI-3 — Slice-3 source handoff and applicability.** Product/Campaign source needs to record the exact D-04 public closure truth and the condition that applies only when Slice-3 closure/result behavior is reconciled. Preserve current Slice-2 closed/non-public 404 until then. Owner: Anhar as scoped Product/Campaign authority; Slice-3 delivery/runtime remain separate.
4. **OI-4 — Owning-source acceptance and counterpart proof.** Human/owner acceptance of materially changed Product/spec/API sources, independent review, and required generated/bundle/type/fixture correspondence remain required to satisfy WU006. Existing WU-S2-005 acceptance is prior baseline, not acceptance of its future amendment. Owner: named source authorities; Orchestrator coordinates review and records outcome.
5. **OI-5 — Delivery-plan refresh and runtime obligations.** After source convergence, fresh backend WU-S2-003 and frontend WU-S2-004 Techplan Runs must consume accepted sources; backend concurrency/DB/capacity proof and frontend rendered disclosure/input acceptance remain delivery/testing work. Work Graph owns dependency topology; this Run does not change it.

### Resolved — retained as decision history

1. ~~**Rp1b individual cap direction**~~ **RESOLVED —** Human clarification 2026-10-01 sets hard upper bound Rp1.000.000.000, per-Campaign whole-IDR range Rp5.000–Rp1.000.000.000, default Rp1.000.000.000, Owner/Staff draft configuration and freeze after publication. Concrete source/schema acceptance remains Active in OI-1.
2. ~~**Guest disclosure direction**~~ **RESOLVED —** Human clarification 2026-10-01 requires pre-entry cap disclosure on public Campaign detail and independent POST recheck. Exact authored contract amendment remains Active in OI-1/OI-4.
3. ~~**Reservation/closure behavior**~~ **RESOLVED —** Human clarification 2026-10-01 selects settled plus accepted-pending reservation, close at capacity, and no reopening after pending failure. Exact close-reason identifier remains Active in OI-2.
4. ~~**Public closure meaning**~~ **RESOLVED —** Human clarification 2026-10-01 assigns public identity retention, donation action removal, and non-final Funding during accepted pending to Slice 3; it does not change current Slice-2 closed detail behavior. Source handoff remains Active in OI-3.
5. ~~**Historical five-brief decision mismatch**~~ **RESOLVED —** Parent `events.md` clarification plus EXP006 terminal handoff identify the four current directions. Do not revive the earlier five Pending briefs or infer acceptance of `funding_capacity_reached`.
