# Slice 2 — Work Graph

> State saat ini, diturunkan dari authority dan bukti yang dirujuk Parent Outcome. File/layout ini adalah realisasi project-local Pilot #2, bukan skema storage Harscode canonical.

## Work Units yang diketahui

| ID | Work Unit | Type | Status | Milestone yang dihasilkan |
|---|---|---|---|---|
| `WU-S2-001` | Slice 2 Authority & Current-State Exploration | `ENABLER` | `DONE` | Evidence untuk menurunkan work rekonsiliasi Slice 2 |
| `WU-S2-002` | Slice 2 Donation Domain & Contract Reconciliation | `RECONCILIATION` | `ACTIVE` / `QUEUED` | `CONTRACT_READY` setelah rekonsiliasi dan acceptance yang berlaku |

## Dependency edges

| Dari | Ke | Strength | Kondisi |
|---|---|---|---|
| `WU-S2-002` | `WU-S2-001` | HARD | `WU-S2-001 = DONE`; evidence Exploration menjadi input rekonsiliasi |

## Runnable frontier

`WU-S2-002` tetap menjadi frontier. Human menyetujui Techplan `TP-S2-002-007` pada 2026-09-27 dan Planner menyelaraskan field Status menjadi `Approved` melalui `TP-S2-002-008`. Authority Map kini mengatribusi lima owner area yang dibutuhkan kepada Anhar Solehudin hanya untuk current Slice 2. `OIR-S2-002-002` siap untuk Human-assisted dispatch sebagai Explorer Run terfokus O3–O5; belum diluncurkan. O1/O2, O7, dan Campaign/Donation ordering tetap memerlukan rute sendiri; O6/O9 adalah kebijakan resolved yang masih perlu diterjemahkan ke contract. O8 audit bersyarat sebelum operasi historis dihapus/diganti. `CONTRACT_READY` belum tercapai.

## Batas derivasi

FE/BE delivery dan integration Work Unit belum diturunkan. Evidence mendukung rekonsiliasi contract lebih dahulu; turunkan implementasi setelah `CONTRACT_READY`.
