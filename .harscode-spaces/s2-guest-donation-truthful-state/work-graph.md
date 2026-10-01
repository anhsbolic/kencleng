# Slice 2 — Work Graph

> State saat ini, diturunkan dari authority dan bukti yang dirujuk Parent Outcome. File/layout ini adalah realisasi project-local Pilot #2, bukan skema storage Harscode canonical.

## Work Units yang diketahui

| ID | Work Unit | Type | Status | Milestone yang dihasilkan |
|---|---|---|---|---|
| `WU-S2-001` | Slice 2 Authority & Current-State Exploration | `ENABLER` | `DONE` | Evidence untuk menurunkan work rekonsiliasi Slice 2 |
| WU-S2-002 | Slice 2 Donation Domain & Contract Reconciliation | RECONCILIATION | ACTIVE / QUEUED for Task 02 authored Donation OpenAPI Build BLD-S2-002-006 (`gpt-6-luna` / `high`). RV-012 confirmed F-001 resolved and closed the Task 01 Review loop; Task 01 Build/Review hard dependency is met. TP-015 remains current-effective Approved; accepted split/dependency/manifest are unchanged. Human/domain-owner acceptance of the five `draft` specs remains a separate gate before CONTRACT_READY, not a blocker to this Task 02 Build. Reconcile supported contract details and preserve field-specific evidence deferrals. | CONTRACT_READY after applicable contract-time decisions/acceptances and reconciled specs/API; runtime evidence remains downstream |

## Dependency edges

| Dari | Ke | Strength | Kondisi |
|---|---|---|---|
| WU-S2-002 | WU-S2-001 | HARD | WU-S2-001 = DONE; evidence Exploration menjadi input rekonsiliasi |

## Runnable frontier

WU-S2-002 remains the frontier Work Unit. TP-015 is current-effective Approved and TPD-002 preserved the accepted split/dependency/manifest. BLD-004/005 reconciled five Task 01 drafts; RV-011's blocking F-001 was fixed and RV-012 independently confirmed it resolved, closing the Task 01 Review loop. Task 02's hard Build/Review dependency is met. Fresh Task 02 OpenAPI Build BLD-S2-002-006 is queued (`gpt-6-luna` / `high`). Human/domain-owner acceptance of the five `draft` specs remains separately required before `CONTRACT_READY`, but does not block this Build. O2–O5 runtime/security proof stays downstream.

## Batas derivasi

FE/BE delivery dan integration Work Unit belum diturunkan. Evidence mendukung rekonsiliasi contract lebih dahulu; turunkan implementasi setelah `CONTRACT_READY`.
