# Slice 2 — Work Graph

> State saat ini, diturunkan dari authority dan bukti yang dirujuk Parent Outcome. File/layout ini adalah realisasi project-local Pilot #2, bukan skema storage Harscode canonical.

## Work Units yang diketahui

| ID | Work Unit | Type | Status | Milestone yang dihasilkan |
|---|---|---|---|---|
| `WU-S2-001` | Slice 2 Authority & Current-State Exploration | `ENABLER` | `DONE` | Evidence untuk menurunkan work rekonsiliasi Slice 2 |
| `WU-S2-002` | Slice 2 Donation Domain & Contract Reconciliation | `RECONCILIATION` | `WAITING_HUMAN` / `PARKED` for material approval of TP-011 after report `TP-S2-002-012`; O1 authority sync and O11 O2/O3 decision remain scoped open items | `CONTRACT_READY` setelah rekonsiliasi dan acceptance yang berlaku |

## Dependency edges

| Dari | Ke | Strength | Kondisi |
|---|---|---|---|
| `WU-S2-002` | `WU-S2-001` | HARD | `WU-S2-001 = DONE`; evidence Exploration menjadi input rekonsiliasi |

## Runnable frontier

`WU-S2-002` tetap menjadi Work Unit frontier, kini menunggu Human gate. Human menyetujui Techplan `TP-S2-002-007` pada 2026-09-27. OIR-006 memutuskan D1; TP-010 dan TP-011 membawanya ke spine. Review `RV-S2-002-005` menemukan tiga blocker, yang Planner resolusikan dalam TP-011: O7 kini Resolved dengan provenance; O2/O3 conflict dicatat sebagai scoped `HUMAN_DECISION`; O1 owner/scope gap dicatat sebagai scoped `AUTHORITY_SYNC`. Fresh independent review `RV-S2-002-006` selesai tanpa findings, lalu Planner membuat report `TP-S2-002-012`. D1 tetap unchanged. TP-011 masih Draft / In Review; current-effective Techplan tetap TP-007 sampai Human menyetujui amendmen material. O1/O11 tetap scoped follow-up. `CONTRACT_READY` belum tercapai.

## Batas derivasi

FE/BE delivery dan integration Work Unit belum diturunkan. Evidence mendukung rekonsiliasi contract lebih dahulu; turunkan implementasi setelah `CONTRACT_READY`.
