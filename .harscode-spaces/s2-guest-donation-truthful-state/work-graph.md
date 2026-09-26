# Slice 2 — Work Graph

> State saat ini, diturunkan dari authority dan bukti yang dirujuk Parent Outcome. File/layout ini adalah realisasi project-local Pilot #2, bukan skema storage Harscode canonical.

## Work Units yang diketahui

| ID | Work Unit | Type | Status | Milestone yang dihasilkan |
|---|---|---|---|---|
| `WU-S2-001` | Slice 2 Authority & Current-State Exploration | `ENABLER` | `DONE` | Evidence untuk menurunkan work rekonsiliasi Slice 2 |
| `WU-S2-002` | Slice 2 Donation Domain & Contract Reconciliation | `ACTIVE` / `QUEUED` | `CONTRACT_READY` setelah rekonsiliasi dan acceptance yang berlaku |

## Dependency edges

| Dari | Ke | Strength | Kondisi |
|---|---|---|---|
| `WU-S2-002` | `WU-S2-001` | HARD | `WU-S2-001 = DONE`; evidence Exploration menjadi input rekonsiliasi |

## Runnable frontier

`WU-S2-002` tetap menjadi frontier. Human menyetujui Techplan efektif pada 2026-09-26 setelah `RV-S2-002-003` bersih dan report tersedia; Planner menyelaraskan status artifact ke `Approved` pada `TP-S2-002-005`. Evaluasi post-approval menetapkan Techplan decomposition `NOT_APPLICABLE`: reconciliation ini satu alur kontrak yang cohesive; pemisahan per file/domain akan membelah shared decision surface dan tidak memberi chunk independen yang aman sebelum O1–O6/O9 diputuskan. Next route adalah authority sync O1–O6/O9 dengan owner Techplan §13; bagian final spec/API yang bergantung keputusan tetap ditahan. O7 bersyarat; O8 diperlukan sebelum breaking API change. O9 memblokir final submit contract/`CONTRACT_READY`. Work Unit `WAITING_HUMAN` / `PARKED` sampai keputusan owner tersedia.

## Batas derivasi

FE/BE delivery dan integration Work Unit belum diturunkan. Evidence mendukung rekonsiliasi contract lebih dahulu; turunkan implementasi setelah `CONTRACT_READY`.
