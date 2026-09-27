# Slice 2 — Work Graph

> State saat ini, diturunkan dari authority dan bukti yang dirujuk Parent Outcome. File/layout ini adalah realisasi project-local Pilot #2, bukan skema storage Harscode canonical.

## Work Units yang diketahui

| ID | Work Unit | Type | Status | Milestone yang dihasilkan |
|---|---|---|---|---|
| `WU-S2-001` | Slice 2 Authority & Current-State Exploration | `ENABLER` | `DONE` | Evidence untuk menurunkan work rekonsiliasi Slice 2 |
| `WU-S2-002` | Slice 2 Donation Domain & Contract Reconciliation | `RECONCILIATION` | `WAITING_HUMAN` / `PARKED` | `CONTRACT_READY` setelah rekonsiliasi dan acceptance yang berlaku |

## Dependency edges

| Dari | Ke | Strength | Kondisi |
|---|---|---|---|
| `WU-S2-002` | `WU-S2-001` | HARD | `WU-S2-001 = DONE`; evidence Exploration menjadi input rekonsiliasi |

## Runnable frontier

`WU-S2-002` tetap menjadi frontier. Human menyetujui Techplan `TP-S2-002-007` pada 2026-09-27 dan Planner menyelaraskan field Status menjadi `Approved` melalui `TP-S2-002-008`. Tidak ada Participant Run yang runnable sampai owner decisions/evidence berikut tersedia: O1–O5 dan O7; O6/O9 adalah kebijakan resolved yang masih perlu diterjemahkan ke contract. O8 audit bersyarat sebelum operasi historis dihapus/diganti. Owner roles tercatat tetapi nama individu belum ditetapkan. Work Unit `WAITING_HUMAN` / `PARKED`; `CONTRACT_READY` belum tercapai.

## Batas derivasi

FE/BE delivery dan integration Work Unit belum diturunkan. Evidence mendukung rekonsiliasi contract lebih dahulu; turunkan implementasi setelah `CONTRACT_READY`.
