# Slice 2 — Work Graph

> State saat ini, diturunkan dari authority dan bukti yang dirujuk Parent Outcome. File/layout ini adalah realisasi project-local Pilot #2, bukan skema storage Harscode canonical.

## Work Units yang diketahui

| ID | Work Unit | Type | Status | Milestone yang dihasilkan |
|---|---|---|---|---|
| `WU-S2-001` | Slice 2 Authority & Current-State Exploration | `ENABLER` | `DONE` | Evidence untuk menurunkan work rekonsiliasi Slice 2 |
| WU-S2-002 | Slice 2 Donation Domain & Contract Reconciliation | RECONCILIATION | DONE | CONTRACT_READY |
| WU-S2-003 | Slice 2 Donation Backend Delivery | DELIVERY | ACTIVE / QUEUED — `EXP-S2-003-001` prepared | BACKEND_VERIFIED |
| WU-S2-004 | Slice 2 Guest Donation Frontend Flow | DELIVERY | ACTIVE / QUEUED — `EXP-S2-004-001` prepared | FRONTEND_MOCK_VERIFIED |

## Dependency edges

| Dari | Ke | Strength | Kondisi |
|---|---|---|---|
| WU-S2-002 | WU-S2-001 | HARD | WU-S2-001 = DONE; evidence Exploration menjadi input rekonsiliasi |
| WU-S2-003 | WU-S2-002 | HARD | WU-S2-002 = DONE / `CONTRACT_READY`; accepted Donation specs and shared contract are the assignment baseline |
| WU-S2-004 | WU-S2-002 | HARD | WU-S2-002 = DONE / `CONTRACT_READY`; accepted Donation specs and shared contract are the assignment baseline |

## Runnable frontier

WU-S2-002 completed at `CONTRACT_READY` after Anhar's explicit Task 02 API contract acceptance (2026-10-01), following TP-015 approval, Task 01 spec acceptance, Build validation/generation, and RV-014 confirmation. The current runnable frontier is contract-parallel: WU-S2-003's `EXP-S2-003-001` and WU-S2-004's `EXP-S2-004-001` are prepared for Human-assisted dispatch at canonical Exploration Stage 1. Backend owns Donation capability; frontend owns the guest Donation flow against contract-faithful MSW mocks. Kencleng §8 and the Integration Map §5 support this route. O2–O5 empirical/runtime/security proof remains delivery/Testing work; no residual risk is accepted.

## Batas derivasi

Dependencies: WU-S2-003 and WU-S2-004 each have a HARD dependency on WU-S2-002=`DONE` / `CONTRACT_READY`; real integration follows backend and frontend side-specific verification. Backend and frontend production writes remain separate. Product truth, security controls, and verification obligations stay governed by the accepted specs/Techplan and scoped stack guidance.
