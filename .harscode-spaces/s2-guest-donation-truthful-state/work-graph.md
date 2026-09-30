# Slice 2 — Work Graph

> State saat ini, diturunkan dari authority dan bukti yang dirujuk Parent Outcome. File/layout ini adalah realisasi project-local Pilot #2, bukan skema storage Harscode canonical.

## Work Units yang diketahui

| ID | Work Unit | Type | Status | Milestone yang dihasilkan |
|---|---|---|---|---|
| `WU-S2-001` | Slice 2 Authority & Current-State Exploration | `ENABLER` | `DONE` | Evidence untuk menurunkan work rekonsiliasi Slice 2 |
| `WU-S2-002` | Slice 2 Donation Domain & Contract Reconciliation | `RECONCILIATION` | `WAITING_HUMAN` / `PARKED` pending scoped API-owner decisions; `BLD-S2-002-003` completed with no contract diff, so independent Code Review is N/A | `CONTRACT_READY` setelah rekonsiliasi dan acceptance yang berlaku |

## Dependency edges

| Dari | Ke | Strength | Kondisi |
|---|---|---|---|
| `WU-S2-002` | `WU-S2-001` | HARD | `WU-S2-001 = DONE`; evidence Exploration menjadi input rekonsiliasi |

## Runnable frontier

`WU-S2-002` tetap menjadi Work Unit frontier. Human menyetujui Techplan TP-007 pada 2026-09-27 dan material amendmen TP-011 pada 2026-09-30 setelah review report TP-012. Planner Run TP-013 menyelaraskan status; TP-011 kini current-effective Approved. OIR-006 memutuskan D1; TP-010 dan TP-011 membawanya ke spine. Review `RV-S2-002-005` menemukan tiga blocker, yang Planner resolusikan dalam TP-011: O7 kini Resolved dengan provenance; O2/O3 conflict dicatat sebagai scoped `HUMAN_DECISION`; O1 owner/scope gap dicatat sebagai scoped `AUTHORITY_SYNC`. Fresh independent review `RV-S2-002-006` selesai tanpa findings. D1 tetap unchanged. Planner Run `TPD-S2-002-001` selesai dengan dua task: Task 01 untuk Donation domain specs (tanpa dependency task) dan Task 02 untuk Donation OpenAPI (bergantung Task 01); Human menerima split pada 2026-09-30. Build `BLD-S2-002-001` menyelesaikan Task 01; `RV-S2-002-007` found F-01, `BLD-S2-002-002` applied its narrow Summary correction, and `RV-S2-002-008` independently confirmed it; Task 01 review loop is complete. Domain specs remain `draft` pending applicable owner/Human acceptance. Build `BLD-S2-002-003` completed the Task 02 assessment and OpenAPI lint with no authored/generated API diff because material operation/credential shapes remain gated. Independent Code Review is N/A under current guidance because there is no diff to review. WU is parked for scoped Human/API-owner decisions on O1/O11 and applicable O2–O5 details; O8 remains conditional only on removing/replacing historical operations. `CONTRACT_READY` belum tercapai.

## Batas derivasi

FE/BE delivery dan integration Work Unit belum diturunkan. Evidence mendukung rekonsiliasi contract lebih dahulu; turunkan implementasi setelah `CONTRACT_READY`.
