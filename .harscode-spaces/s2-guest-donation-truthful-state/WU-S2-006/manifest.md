# WU-S2-006 — Slice 2 Monetary Limits & Capacity Contract Reconciliation

## Definition

- Type: `RECONCILIATION`
- Parent Outcome: `S2-GUEST-DONATION-TRUTHFUL-STATE`
- Derived from: backend `TP-S2-003-003` D16/O1-REP dan completion handoff; Product/MVP authority attribution 2026-10-01.
- Coordination owner: Orchestration Operator
- Communication language: Bahasa Indonesia

### Outcome

Menghasilkan baseline Product/MVP, Donation/Campaign spec dan shared API yang reconciled serta diterima owning authorities untuk batas individual Donation yang configurable per Campaign, kapasitas funding kumulatif, dan closure saat kapasitas tercapai. Baseline ini menutup source-reconciliation gap O1-REP agar backend dan frontend dapat menyelesaikan planning tanpa mengarang policy/interface.

### Scope

- Rekonsiliasi TP-S2-003-003 D16/O1-REP dan accepted EXP006 D-01–D-04 yang diperjelas Human pada 2026-10-01. Exact current receipt di parent events heading “Human clarified EXP006 D-01–D-04; source Techplan frontier prepared”; concrete source acceptance tetap explicit gate.
- Derive affected Product/MVP, monetary representation bila applicable, Donation/Campaign spec/API/generated/fixture serta known consumer surfaces. Direction settled: cap whole-IDR Rp5.000–Rp1 miliar/default Rp1 miliar, Owner/Staff saat draft/freeze publikasi; public guest pre-disclosure/POST recheck; reservation settled plus accepted pending/close saat habis/no reopen setelah gagal. DEC-API-01/02 juga menetapkan closed `{amount, currency_code}` representation dan shared generic `422` pada eligible capacity no-fit; concrete source bytes, compatibility evidence, Review dan owner acceptance tetap required.
- Pertahankan full settlement accepted-pending, exact-once funding, idempotent retry dan winning-close-reason obligations; bedakan capacity ceiling dari existing threshold.

### Boundaries

- Shared authority/spec/contract reconciliation; backend/frontend production, migrations/DB application dan runtime proof tetap milik Delivery.
- Jangan mengubah accepted Campaign availability-only action secara tersirat; perubahan public contract memerlukan owning API scope/review/acceptance.
- D-04 membawa capacity close ke Slice 3: public identity tetap tersedia, donation action hilang, Funding tidak final selama accepted pending. Catat source/consumer handoff explicit; jangan expand full Slice-3 Delivery atau silently activate closed visibility pada current Slice-2 API. Jangan menetapkan universal monetary cap, menerima residual risk atau menulis protected implementation.
- Orchestrator mengoordinasikan; Participant menyusun phase artifacts. Product/MVP/domain/API source changes hanya setelah decision dan gate owning authority yang applicable.

### Completion condition

Material policy/interface questions settled dengan named authority; owning Product/spec/API sources dan required counterparts reconciled, independently reviewed/verified sesuai risiko serta diterima owner. Exact consumer/Delivery readiness dan remaining runtime obligations dicatat. Tidak ada runtime milestone atau whole-backend/frontend plan approval yang implied.

## Current State

- Execution status: `ACTIVE`
- Scheduling state: `QUEUED` — owner compatibility/rollout decision and exact seven-source acceptance recorded; BLD-S2-006-006 completed its bounded generated API/types/fixture counterpart target. Fresh backend WU-S2-003 and frontend WU-S2-004 Techplan refreshes are the next planning frontier under approved TP006008 §9.
- Horizon: `NOW`
- Last completed Run: BLD-S2-006-006 completed. Previous Runs RV-S2-006-006, BLD-S2-006-005 remain recorded below. TP-S2-006-009 was not dispatched: TP006008 approval Status was reconciled directly by Orchestrator under current Harscode deterministic-reconciliation guidance.
- Current milestone: None
- Review outcome: RV006001 and RV006005 completed without blocking findings. RV-S2-006-006 completed `Approve with minor comments`, no patch plan; it raised RV-006-01 for owner strict-client/known-consumer compatibility and rollout review. Anhar resolved that gate for the current MVP1 distribution posture and accepted the exact source hashes below. Review artifact SHA-256 `df2edd50f7cdc796dadae817ff5a13fe7d8977e87d7ebdf4e18ebdaa3c43bb02`.
- Human gates: D-01–D-04/D6/OI-1/OI-2 settled. TP-S2-006-008 whole-Techplan approval by Anhar is recorded against report TP006008; Status `Approved` reconciled by Orchestrator after exact source/report hash checks and byte-equality verification. Concrete Product/MVP amendments independently reviewed Approve/no findings and accepted Anhar; RV006003 C-01/Q-01 independently resolved by RV006004 Approve; the prior six-source Campaign/Donation amendments accepted Anhar. On 2026-10-02 Anhar confirmed the current MVP1 distribution posture, selected coordinated internal counterpart reconciliation before delivery/runtime, and explicitly accepted the seven exact source revisions listed in `RV-S2-006-006/invocation.md`; current hashes were verified against that table. RV-006-01 is resolved for this scoped rollout. OI-3 applicability and OI-5 delivery refresh/runtime remain downstream. Decomposition Skip recommendation, no split.
- Current evidence: current planning target TP006008 Approved, SHA-256 `93c09bf7629500cd8fb80fd59b6af464b169484419d722a269b782b78bbbf438`; approval snapshot reconstructed by reverting only Status hashes to `c1a8806a1c754b849c0b8457e688d9a50aa0d024dc3d4fe2f3c4c5a7d35c3324` (PASS). Matching report SHA-256 `06d6259f4e50960dbae04951b0b786c15503ec6140f27b9f805149a8a5b6eb6c`; handoff `55ab4744fd385d61ca1831420d71e98b026b223f2d324db7d5b7a0cb8b7c5333`. TP006009 Invocation was prepared but not dispatched; direct reconciliation is supported by current Harscode run-contract §Deterministic reconciliation outside the Run path. Predecessor TP006004 remains historical, not current-effective. Techplan approval remains distinct from source acceptance; the latter is explicitly recorded below.
- Dependent frontend evidence: TP-S2-004-001 OI-1/R2/R12/handoff, Draft completed; whole-plan refresh/review/approval and amount-limit-dependent Build now follow accepted source convergence and WU006 counterpart completion.
- Dependency ownership: parent Work Graph. Backend/frontend affected final approval/Build dan plan refresh/re-review mengikuti accepted source convergence; policy solutioning tidak menutup O1-REP/OI-1 atau membuktikan runtime.

- Build evidence: BLD-S2-006-005 report SHA-256 `b9f0670fcc6bbb41a49ac6f09c4fb44bdaead9f0ebb486851b3c4e89b947d806`; seven authored files changed. Report records focused OpenAPI validation success (124 warnings / 0 errors; warning-coordinate/rule baseline unchanged) and `git diff --check` pass. These are Participant-reported checks, not rerun by Orchestrator; no runtime/broad tests are claimed.
- Current project distribution posture (owner-confirmed 2026-10-02, scoped to current MVP1 contract rollout): Kencleng has not rolled out to production and has no external client/consumer dependent on this contract; current consumers are internal/repository-development consumers. Historical O8 remains limited to its original operations and scope and is not broadened by this posture.
- Compatibility evidence: before BLD006006, `frontend/lib/api/public-campaign.ts` cast `response.json()` to generated `PublicCampaignDetail` without runtime schema validation, while generated types/fixtures lacked `max_donation_amount`. BLD006006 reconciled those generated/frontend fixture counterparts. Backend `TestPublicCampaignDetailHandler_ExactClosedWireShape` still asserts the old nine-key wire shape and remains assigned with the WU003 response DTO/mapping. Owner's distribution decision resolves external strict-client uncertainty for this MVP1 rollout; it does not claim universal future compatibility.
- Source acceptance receipt: Anhar explicitly accepted the seven reviewed Campaign/Donation spec/API revisions by exact SHA-256 snapshot in `RV-S2-006-006/invocation.md`. Current hashes match all seven values. Acceptance applies only to those source bytes.
- Counterpart Build evidence: BLD-S2-006-006 completed the generated API bundle/types and contract-facing public Campaign fixture/test updates. Report SHA-256 `650022ed32d83065a5f593f603340a5e67fad8e55f8091187f7282fda9f1aea1`. Participant reports bundle/type generation, OpenAPI validation (124 warnings, zero errors), focused frontend tests (2 files/9 tests), and `git diff --check` passed; Orchestrator did not rerun them. Orchestrator verified the seven accepted source hashes unchanged, checked the changed-file diff against Invocation scope, recomputed all four output hashes against the report, and ran `git diff --check` successfully. No backend response implementation, exact-wire test, public UI display, or runtime behavior changed.
- Next gate: per approved TP006008 §9, prepare fresh backend WU-S2-003 and frontend WU-S2-004 Techplan refresh Runs against the accepted sources and completed generated/frontend counterpart. Carry the backend exact-wire assertion with its response DTO/mapping into WU003; carry public cap disclosure into WU004. Both remain required before runtime/delivery completion. OI-3 remains an explicit Slice-3 source/applicability handoff for that slice; it does not activate closed-detail behavior in Slice 2. WU-S2-006 remains incomplete pending delivery refresh/readiness handoff and remaining evidence.

## Current-effective inputs

- Root AGENTS dan current Product/MVP/monetary/spec/API owners melalui routing map.
- Parent Outcome, Work Graph, Control Surface, events dan `.harscode-spaces/authority-map.md`.
- Backend source handoff dan exact D16/O1-REP di TP-S2-003-003; successor masih Draft/unapproved.
- Current accepted Donation baseline WU-S2-002 dan Campaign action baseline WU-S2-005; jangan memuat seluruh Run history untuk rekonstruksi.
- Completed EXP006 corpus/handoff dan explicit current Human clarification pada parent events; canonical Harscode Techplan synthesis, template/rules/guardrails, run-contract dan orchestrated-run overlay.

## Current source checkpoint

- BLD006001 Product amendments: docs/product/mvp-scope.md SHA-256 ba2972bc8f91d092e477df170d987b1d124964d9cc36c025d2a8da3ed12709af; docs/product/mvp-delivery-slices.md SHA-256 4c69a030e7fedc9c62bf30f85c00e81f9806c45b2ed5471c1d5126762be8091f. Current hashes match Review.
- RV006002 review-findings SHA-256 f3d4402f6514c0625c49fafee2f004ac5d552e57a1892434e01c1bd72f2d3f88; four-pass Approve, no findings/patch plan. No executable/runtime verification claimed.
- Next gate: independent four-pass Review RV-S2-006-006 for the seven BLD006005 source changes, then owning spec/API acceptance if Review permits. Product amendments and the earlier six-source snapshot remain accepted; these do not pre-accept new source bytes. WU006 remains incomplete.

- Spec checkpoint BLD006002 report SHA-256 e0a7688cf36d1f4d0d9fefbcb2f211632814c536be3576ea614b0ea202c71d8b; source-delta.patch SHA-256 d0bc1537eb5336a66d23128bc63fd8259bd5607dd863ed2eae76871719439560. Six baseline/current hashes verified, exact patch reconstruction matched. Source acceptance pending.

- RV006003 completed Request changes; review SHA-256 bf4f5aed6f9b8506221a537487dd1f7a2d2807b0e74ad6667800b8a1bd69d74e; patch-plan SHA-256 3cd6a4743b3866cc56f736cf7178b671e74b8cc9eac22d13a8e88b429eb23767. All six current source hashes remain reviewed snapshot. No owner decision needed for findings; approved Product/Techplan already own required behavior.

- BLD006003 patch report SHA-256 c2ee2c37777d0d7ae55dc97c74457c8fab73ddb7e3e1f1b89210871a15c43e34; exact source-delta.patch SHA-256 f089b009eb5acc176279c3e5226ac91dd4005a512580892c32259980f8ae9c5c. Four snapshot/live deltas reconstructed MATCH; two other original spec sources unchanged. C-01/Q-01 independently closed by RV006004; concrete spec acceptance received.

## Spec acceptance snapshot — 2026-10-02

- `docs/spec/4-campaign/invariants.md` SHA-256 `0d3f250e4d51d98a63c865c6fe0810b1b4fdb5ab063b706c7ab32f7e7bd73342`.
- `docs/spec/4-campaign/features/01-campaign-creation-draft-crud.md` SHA-256 `53431422b714adebd6ff9f07476c4c930746516f89f98d4c7ea122b7f35c8d05`.
- `docs/spec/4-campaign/features/02-campaign-detail-listing.md` SHA-256 `2db5dd950553e97756417fd2f41e29356835c43a795c80c3be1d615fc3cd45a7`.
- `docs/spec/4-campaign/features/09-closure.md` SHA-256 `dd0a3c6a897091e887a4e64127322d8e4d28036bd15f6cedc70777f7d329df4d`.
- `docs/spec/5-donation/invariants.md` SHA-256 `68c7967fba44a3012ee67730bc5b6a2011961859b1b119dbe0709d96d0d953f8`.
- `docs/spec/5-donation/features/01-submit-donation-settlement.md` SHA-256 `3b1f3918752724c4c6448aa001f287ee59461a5fb8f75b82a0495db072755dfe`.
- RV006004 review SHA-256 f4f47f08b72380aa6ed95ca1c09054731897d93dd747bc7c86b2ef4a7a527bd5. Current all-six hashes matched original Review plus independent patch confirmation. Human explicit acceptance recorded in parent event “Human accepted six spec amendments; authored API Build prepared”.

## Authored API decision gate — 2026-10-02

- BLD006004 report SHA-256 15dafd0eaa302ae6a36bce703661b42c547725c84b856926db736da371f66d65. No authored API change. All six spec amendment receipts verified normalized-byte-equivalent to exact Human-accepted snapshot; existing substantive Review remains applicable, no new Review/Testing solely for marker edits.
- DEC-API-01: Campaign still eligible but requested Donation does not fit remaining capacity. Spec explicitly defers transport; approved plan selects over-cap 422 only. Owner Anhar asked generic shared 422 ValidationError on amount vs generic 409, without capacity disclosure; closed/ineligible and idempotent retry rules preserved. Decision accepted Anhar; exact receipt in parent owner-settlement event.
- DEC-API-02: max_donation_amount concrete wire encoding with explicit currency. Owner Anhar asked closed object {amount: decimal string, currency_code: IDR} vs scalar plus companion currency. Existing optional/create-default/PATCH-preserve/required-response decisions preserved. Decision accepted Anhar; exact receipt in parent owner-settlement event.
- Known owner can answer now; no discovery Participant merely to ask. Material chosen shape must become owning Planner durable contract before API authoring, with applicable review/report/approval fidelity gates. No source acceptance inferred from choice.

- DEC-API-01/02 RESOLVED: generic shared 422 ValidationError on amount for eligible capacity no-fit; closed object max_donation_amount {amount: decimal string, currency_code: IDR}, both members required when supplied, inherited outer optional/default/PATCH-preserve/required-response semantics retained. Owner choice is not whole successor approval or concrete source acceptance.

- TP006007 complete with no new material owner decision reported; recommend independent Review, decomposition Skip. No report during material revision churn. Affected future spec/API bytes still require applicable owning acceptance; prior Product/six-spec acceptance remains specific to its recorded snapshot.

- RV006005 completed independent Complex Review: no blocking findings, one non-blocking invalid secondary Test Focus evidence anchor. Review SHA-256 b093a49238d0c9737cbf73d4a80fe4fa51d196c4794c683177f8f290e1ca5783. Mechanical-only correction does not trigger re-review by itself; actual delta must support substantive Review carry-forward.

- TP006008 full actual diff inspected: invalid secondary anchor removed, valid Area2 retained; necessary Run provenance/artifact paths and resolution note only. Material semantics unchanged; RV006005 substantive independent evidence applicable, no mechanical-only re-review. Report/source hash correspondence verified; approval applies plan, not digest.

- Human whole TP006008 approval receipt: “techplan approve bro”, parent event 2026-10-02 — Human approved TP006008; Status propagation prepared. Exact plan/report hashes unchanged; no repeat approval vote. Future source-byte acceptance remains separate.

- Human checkpoint 2026-10-02: berhenti sementara sebelum TP006009 dispatch. Execution ACTIVE retained (target unfinished), Scheduling PARKED; prepared Invocation preserved. Progress snapshot: docs/project/slice-2-progress-checkpoint-2026-10-02.md. Resume requires current-state/guidance re-grounding, not a repeated approval.
