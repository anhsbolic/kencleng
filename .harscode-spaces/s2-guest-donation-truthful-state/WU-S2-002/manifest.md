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

- Execution status: `WAITING_HUMAN`
- Horizon: `NOW`
- Current Run: `TP-S2-002-001` — Techplan Synthesis completed; Draft/In-Review artifacts await Human gate
- Current milestone: None
- Human gate: Pilih independent review atau direct review. Planner menghasilkan `report-techplan.md` pada gate setelah jalur review/resolution yang dipilih konvergen; lalu Human approve/revise Techplan. Jika disetujui, O1–O6 tetap harus diputuskan owner sebelum bagian contract terkait difinalisasi; O7 bersyarat dan O8 wajib sebelum perubahan breaking.
- Authority sync: Pending bila Techplan menemukan keputusan yang memerlukan owner authority.
- Active blocker: `HUMAN_DECISION` — Human Techplan gate belum selesai; material Open Items O1–O6 mencegah finalisasi contract sampai owner berwenang memberi keputusan.
- Blocker owner/action: Human Authority memilih review route; Planner menghasilkan report pada timing yang benar setelah review/resolution; Human Authority memberi approve/revise; Orchestration Operator merutekan O1–O6 ke owner yang sesuai setelah arahan Human.
- Updated: 2026-09-26

## Current-effective prior artifacts

- `../WU-S2-001/runs/EXP-S2-001-001/evidence/stage-2-gap-analysis.md`
- `../WU-S2-001/runs/EXP-S2-001-001/evidence/stage-3-solutioning.md`
- `../WU-S2-001/manifest.md`
- `../work-graph.md`
- `../outcome.md`
- `runs/TP-S2-002-001/techplan.md` — current Draft/In-Review Techplan
- `runs/TP-S2-002-001/launch-record.md` — completed Run handoff

## Routing note

Run `TP-S2-002-001` selesai dengan `techplan.md` Draft/In-Review. Report yang dibuat Orchestrator ditarik sebagai premature/misowned; Planner akan menghasilkan report setelah review/resolution route konvergen. Work Unit menunggu Human gate; tidak ada scheduling aktif. Jangan memulai Build sebelum Techplan disetujui dan gate material yang relevan dijelaskan.
