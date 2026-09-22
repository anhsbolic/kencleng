# WU-S1-006 — Slice 1 Real Integration & Final Verification

Type:
VERIFICATION

Parent Outcome:
S1 — Public Campaign Understanding

Derived From:
WU-S1-002 — CONTRACT_READY

Status:
NOT_STARTED

Scheduling:
PARKED

Horizon:
NEXT

Readiness:
DEFINED

Coordination Owner Role:
Orchestration Operator

Primary Execution Role:
Verifier

Specialization:
Cross-stack integration / product outcome

Communication Language:
Bahasa Indonesia

Communication Profile Path:
`docs/project/communication-profile.md`

## Outcome

Verify the real Slice-1 experience end-to-end through the same-origin application boundary using persisted backend state and production frontend data access, then provide evidence for `INTEGRATED_VERIFIED` and the later Human Slice-finalization decision.

## Scope

- Real frontend → same-origin proxy → backend public Campaign detail integration.
- Real controlled media byte delivery through private-storage topology.
- Required success/loading/missing-media/not-public/failure/responsive states at the integrated boundary.
- Cross-stack verification of contract correspondence, cache/header preservation, public eligibility/non-disclosure, and truthful presentation.
- Human rendered/product acceptance evidence required before Slice finalization.

## Out of Scope

- New product features or contract redesign.
- Donation Flow, closure/result/accountability.
- Using integration work to hide missing backend/frontend/topology verification.

## Dependencies

HARD:
- WU-S1-003 complete with `BACKEND_VERIFIED`.
- WU-S1-004 complete with `FRONTEND_MOCK_VERIFIED`.
- WU-S1-005 complete with topology/private-media evidence.

## Authority / Contract

- WU-S1-002 current reconciled contract.
- `docs/project/kencleng-integration-map.md`.
- Current backend/frontend/topology artifacts and their independent verification evidence.
- Slice-1 completion evidence in `docs/product/mvp-delivery-slices.md`.

## Human Gates

Material rendered UI requires Human acceptance. `SLICE_FINALIZED` remains a Human/project decision after integration evidence; this WU must not self-approve it.

## Completion

Complete when real cross-stack evidence supports `INTEGRATED_VERIFIED` and all remaining Slice-1 Human acceptance inputs are explicit. It must not claim downstream Slice 2 capability.
