# Slice 2 — Work Graph

> State saat ini, diturunkan dari authority dan bukti yang dirujuk Parent Outcome. File/layout ini adalah realisasi project-local Pilot #2, bukan skema storage Harscode canonical.

## Work Units yang diketahui

| ID | Work Unit | Type | Status | Milestone yang dihasilkan |
|---|---|---|---|---|
| `WU-S2-001` | Slice 2 Authority & Current-State Exploration | `ENABLER` | `DONE` | Evidence untuk menurunkan work rekonsiliasi Slice 2 |
| `WU-S2-002` | Slice 2 Donation Domain & Contract Reconciliation | `RECONCILIATION` | `ACTIVE` / `QUEUED` for independent Campaign/Donation ordering Run; O1 authority sync and O2/O3 decision remain scoped open items | `CONTRACT_READY` setelah rekonsiliasi dan acceptance yang berlaku |

## Dependency edges

| Dari | Ke | Strength | Kondisi |
|---|---|---|---|
| `WU-S2-002` | `WU-S2-001` | HARD | `WU-S2-001 = DONE`; evidence Exploration menjadi input rekonsiliasi |

## Runnable frontier

`WU-S2-002` tetap menjadi Work Unit frontier. Human menyetujui Techplan `TP-S2-002-007` pada 2026-09-27 dan Planner menyelaraskan field Status menjadi `Approved` melalui `TP-S2-002-008`. Authority Map mengatribusi Donation/API dan area Slice 2 lain kepada Anhar hanya untuk current Slice 2. `OIR-S2-002-002` selesai dengan keputusan owner O3–O5 dan open control/risk evidence; `OIR-S2-002-003` selesai tanpa keputusan angka; `TP-S2-002-009` menyelesaikan proposal O2 delivery/O3 retention. Konflik O2/O3 tetap terbuka. `OIR-S2-002-004` menyelesaikan O7 Design review dengan keputusan wording tanpa visual acceptance. `OIR-S2-002-005` selesai dan O1 tetap partially resolved: Human mengarahkan standar currency lintas currency/table/feature, tetapi wire/storage standard dan owner project-wide belum ditetapkan. Blocker O1 terbatas pada standard/amount contract dan tidak menutup pekerjaan independen. `OIR-S2-002-006` disiapkan sebagai frontier berikutnya untuk Campaign/Donation threshold and settlement ordering; O6 policy sudah settled dan owner sudah dipetakan untuk current Slice 2. Human-assisted dispatch diperlukan, lalu canonical Stage 1/Stage 3 gates berlaku. O2/O3, O4/O5 evidence, O6/O7/O9 spec/API translation, dan conditional O8 tetap perlu rute. `CONTRACT_READY` belum tercapai.

## Batas derivasi

FE/BE delivery dan integration Work Unit belum diturunkan. Evidence mendukung rekonsiliasi contract lebih dahulu; turunkan implementasi setelah `CONTRACT_READY`.
