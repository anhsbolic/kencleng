# WU-S1-006 — Slice 1 Real Integration & Final Verification

Type:
VERIFICATION

Parent Outcome:
S1 — Public Campaign Understanding

Derived From:
WU-S1-002 — CONTRACT_READY

Status:
DONE

Scheduling:
DONE

Horizon:
NOW

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


## Integration Verification Dispatch

Dependency Gate:
SATISFIED
- `WU-S1-003 = BACKEND_VERIFIED`
- `WU-S1-004 = FRONTEND_MOCK_VERIFIED`
- `WU-S1-005 = TOPOLOGY_VERIFIED`

Current Run:
`TST-INT-001`

Session:
Fresh independent integration Testing

Model:
`gpt-5.6-terra / high`

Target Milestone:
`INTEGRATED_VERIFIED`

Human Gate After Testing:
Integrated rendered/product acceptance and later `SLICE_FINALIZED` decision remain Human-owned.


## Integration Milestone Gate

Testing:
`TST-INT-001` — PASS_WITH_FLAGGED_FOLLOWUPS

Evidence:
SUFFICIENT_FOR_INTEGRATED_VERIFIED

Human Gate:
REQUIRED — integrated rendered/product acceptance before milestone promotion.

Target Milestone:
`INTEGRATED_VERIFIED`

Slice Finalization:
NOT_YET_APPROVED — remains a separate Human/project decision after integration milestone promotion.

Non-blocking follow-ups:
- repository-wide pre-existing `gosec` / Go cache reproducibility;
- historical `TST-FE-001` absent-media literal differs from current UI wording, without semantic product mismatch.


## Integration Milestone Promotion

Human Approval:
APPROVED — integrated rendered/product acceptance and `INTEGRATED_VERIFIED`

Evidence Basis:
- `TST-INT-001` — PASS_WITH_FLAGGED_FOLLOWUPS
- Human integrated rendered/product acceptance — APPROVED

Milestone:
`INTEGRATED_VERIFIED`

Status:
DONE for WU-S1-006 integration-verification scope.

Important Boundary:
`SLICE_FINALIZED` is NOT implied by this milestone and remains a separate Human/project decision.

Non-blocking follow-ups:
- pre-existing repository-wide `gosec` / Go cache reproducibility;
- historical `TST-FE-001` exact absent-media literal differs from current UI wording without semantic mismatch.

Authority Sync:
NONE


## Slice Finalization

Human Approval:
APPROVED — `SLICE_FINALIZED`

Scope:
Slice 1 — Public Campaign Understanding

Evidence Basis:
- `CONTRACT_READY`
- `BACKEND_VERIFIED`
- `FRONTEND_MOCK_VERIFIED`
- `TOPOLOGY_VERIFIED`
- `INTEGRATED_VERIFIED`
- Human integrated rendered/product acceptance

Final State:
`SLICE_FINALIZED`

Important Boundary:
Slice 2 capability is NOT_STARTED and is not implied by this finalization.
