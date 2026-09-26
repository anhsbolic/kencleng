# WU-S2-002 — Slice 2 Donation Domain & Contract Reconciliation

## Definition

- Type: `RECONCILIATION`
- Parent Outcome: `S2-GUEST-DONATION-TRUTHFUL-STATE`
- Derived from: `WU-S2-001` / `EXP-S2-001-001`
- Coordination owner role: Orchestration Operator
- Planned first phase role: Planner
- Specialization: None established
- Communication language: Bahasa Indonesia
- Communication profile: `docs/project/communication-profile.md`

### Outcome

Rekonsiliasi kebutuhan Slice 2 dengan Donation domain authority dan shared API contract agar satu contract Slice-2-specific cukup stabil untuk menurunkan delivery Work Unit berikutnya. Work ini tidak mengubah Product Authority dan tidak mengisi gap dengan asumsi.

### Scope

- Menyelaraskan batas Slice 2 dengan Product/MVP authority dan prinsip Product Design yang applicable.
- Mengevaluasi serta mengklasifikasikan detail Donation spec, invariants, threat model, dan split OpenAPI yang sudah ada sebagai `KEEP`, `ADAPT`, `REPLACE`, atau `DEFER`.
- Merutekan keputusan material kepada authority owner yang tepat; mempertahankan keputusan yang belum tersedia sebagai blocker atau deferred item, bukan mengasumsikannya.
- Menetapkan kebutuhan contract yang didukung authority untuk guest submission, persisted truthful donation state, dan safe guest revisit, termasuk batas eligibility yang harus disepakati sebelum delivery.
- Menyediakan handoff dan evidence yang diperlukan untuk menentukan apakah `CONTRACT_READY` dapat diterima serta dependency delivery setelahnya.

### Evidence-backed open areas

Handoff Exploration mencatat pertanyaan berikut; ini bukan keputusan baru atau requirement tambahan:

- Amount floor/precision, sandbox payment representation, dan timing/outcome semantics.
- Minimum guest fields, optional email/retention behavior, dan guest revisit credential/handling.
- Semantik response publik untuk status credential yang absent/invalid, termasuk mismatch `401`/`404`.
- Relasi `max_amount`, successful donation, serta eligibility Campaign berikutnya.
- UI wording/source-label hanya bila semantik state yang direkonsiliasi memerlukannya.

### Out of scope

- Implementasi backend/frontend, test execution, atau perubahan behavior code.
- Mengubah Product, security, interface, design, architecture, atau verification authority.
- Mengadopsi detail historis hanya karena sudah ada di spec, OpenAPI, code, atau test.
- Perubahan protected Tier-0 paths tanpa authority yang diwajibkan.
- Membentuk downstream backend/frontend topology sebelum contract dan dependency cukup stabil.

### Completion condition

Handoff reconciliation menyatakan sumber contract Slice 2 yang berlaku dan status detail historisnya; seluruh keputusan yang menjadi prasyarat material telah diselesaikan oleh owner berwenang atau dicatat sebagai blocker eksplisit; dan evidence cukup untuk menerima atau menolak `CONTRACT_READY` serta menurunkan delivery topology tanpa mengarang authority. Milestone `CONTRACT_READY` belum earned.

## Current State

- Execution status: `ACTIVE`
- Scheduling state: `QUEUED`
- Horizon: `NOW`
- Current Run: `TP-S2-002-006` — prepared; material amendment from current Product/MVP decisions and OIR brief, ready for Human-Assisted dispatch
- Current milestone: None
- Human gate: Human amended and approved `docs/product/mvp-scope.md` and `docs/product/mvp-delivery-slices.md` on 2026-09-26, incorporating the OIR priority register. This resolves the prior Product/MVP scope conflict for Slice 2. Current-effective Techplan `TP-S2-002-003` remains Approved but predates those amendments and is stale as a planning baseline; its material replacement requires independent review and a new Human approval. O6/O9 policy is resolved; O1–O5 have remaining technical/owner details; O7 needs Design review; O8 is conditional on API removal/replacement.
- Authority sync: Product/MVP direction is now canonical and must be used by Planner. Security/PII risk acceptance and technical controls, Design expression, and API response/contract details remain with their respective owners. O9's policy must be preserved in final submit contract before `CONTRACT_READY`.
- Active blocker: `AUTHORITY_SYNC` — O1 amount schema/derived precision, O2 internal timing + Design review, O3 guest-email verification/retention/retry controls, O4 token exposure/security controls, O5 response parity/anti-enumeration, and API translation of resolved O6/O9. No Build or `CONTRACT_READY` yet.
- Blocker owner/action: Human-Assisted dispatch of prepared `TP-S2-002-006` to Planner. After amendment, run fresh independent Techplan review because the plan is Complex and the change is material; Planner then regenerates the report only after review/resolution converges, followed by Human approval. O8 is required only before historical operations are removed/replaced.
- Updated: 2026-09-26

## Current-effective prior artifacts

- `../WU-S2-001/runs/EXP-S2-001-001/evidence/stage-2-gap-analysis.md`
- `../WU-S2-001/runs/EXP-S2-001-001/evidence/stage-3-solutioning.md`
- `../WU-S2-001/manifest.md`
- `../work-graph.md`
- `../outcome.md`
- `runs/TP-S2-002-003/techplan.md` — current-effective Techplan approved by Human; Planner status field records `Approved`
- `runs/TP-S2-002-002/techplan.md` — prior plan after atomic-coupling resolution
- `runs/TP-S2-002-001/techplan.md` — prior synthesis Techplan, superseded for current review by resolution artifact
- `runs/TP-S2-002-001/launch-record.md` — completed Run handoff
- `runs/RV-S2-002-001/invocation.md` and `launch-record.md` — completed independent review handoff
- `runs/RV-S2-002-001/review-findings.md` — completed independent review; one blocking `[MONEY / VERIFICATION]` finding
- `runs/TP-S2-002-002/invocation.md` and `launch-record.md` — completed Planner resolution run
- `runs/RV-S2-002-002/invocation.md` — completed independent re-review invocation
- `runs/RV-S2-002-002/review-findings.md` and `launch-record.md` — completed re-review; atomic coupling closed and new idempotency finding opened
- `runs/RV-S2-002-003/invocation.md`, `review-findings.md`, and `launch-record.md` — completed clean mandatory re-review; confirms R8/O9, Test Focus, and atomic-coupling fidelity
- `runs/TP-S2-002-003/invocation.md` — prepared Planner resolution for the new finding
- `runs/TP-S2-002-003/launch-record.md` — completed resolution handoff and materiality classification
- `runs/TP-S2-002-004/invocation.md` — prepared Planner-owned report generation for the Human Techplan approval gate
- `runs/TP-S2-002-004/report-techplan.md` and `launch-record.md` — completed Planner report handoff; Human approved the Techplan
- `runs/TP-S2-002-005/invocation.md` — prepared Planner-owned status reconciliation from the explicit Human approval
- `runs/TP-S2-002-005/launch-record.md` — completed status reconciliation; no substantive plan edits
- `runs/OIR-S2-002-001/invocation.md` — prepared Explorer facilitation run for the Approved Techplan's complex, interdependent Open Items
- `runs/OIR-S2-002-001/resolution-brief.md` — completed Open-Item Resolution handoff; O1–O9 outcomes and owner continuations
- `runs/TP-S2-002-006/invocation.md` — prepared material Techplan amendment using updated Product/MVP authority and OIR outcomes; awaiting Human-Assisted dispatch
- `runs/RV-S2-002-003/invocation.md` — prepared required independent re-review

## Routing note

`RV-S2-002-001` menemukan atomic coupling gap; Planner menutupnya di `TP-S2-002-002`. `RV-S2-002-002` menemukan guest submission idempotency gap; Planner mencatat R8/O9 di `TP-S2-002-003`, lalu `RV-S2-002-003` mengonfirmasi penutupan review. Human menyetujui Techplan; Planner menyelaraskan status di `TP-S2-002-005`. OIR `OIR-S2-002-001` memetakan O1–O9; Human kemudian memperbarui Product/MVP authority untuk keputusan Slice 2. Techplan lama kini mendahului authority dan perlu amendment material. `TP-S2-002-006` disiapkan; next action Human-Assisted dispatch. Setelah Planner, independent re-review, lalu report dan Human approval untuk amended Techplan. O1–O5 masih punya owner/technical follow-up; O6/O9 harus tercermin di contract. O9 tetap menghalangi final submit contract/`CONTRACT_READY`. Decomposition `NOT_APPLICABLE`; belum ada Build atau `CONTRACT_READY`.
