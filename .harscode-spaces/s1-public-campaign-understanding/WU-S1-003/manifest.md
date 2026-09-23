# WU-S1-003 — Slice 1 Backend Public Campaign Delivery

Type:
DELIVERY

Parent Outcome:
S1 — Public Campaign Understanding

Derived From:
WU-S1-002 — CONTRACT_READY

Status:
NOT_STARTED

Scheduling:
DISPATCHED

Horizon:
NOW

Readiness:
READY

Coordination Owner Role:
Orchestration Operator

Primary Execution Role:
Explorer

Specialization:
Backend / Go / Campaign domain

Communication Language:
Bahasa Indonesia

Communication Profile Path:
`docs/project/communication-profile.md`

## Outcome

Implement and independently verify the minimum persisted backend capability needed for Slice 1: a real eligible public Campaign can be read through the reconciled public projection and its media bytes can be served through the controlled parent/member-authorizing operation.

Target project milestone:
`BACKEND_VERIFIED`

## Scope

- Minimum persisted Organization/steward, Campaign, funding, lifecycle, and Campaign-media state required by Slice 1.
- Campaign-domain/repository/service/HTTP behavior needed by `getPublicCampaignDetail`.
- Backend-authoritative public eligibility, anti-enumeration behavior, funding truth, provenance mapping, and unavailable donation action.
- Controlled Campaign media byte delivery with parent eligibility + media membership recheck.
- Private-storage integration boundary required by backend code.
- Narrow seeded/operator-assisted setup needed to make a genuine persisted public Campaign available.
- Backend tests/evidence required by the current reconciled spec/threat model and later Techplan.

## Out of Scope

- Frontend implementation.
- Root Caddy / Docker Compose / MinIO policy mutation owned by WU-S1-005.
- Donation Flow, Account dependency, closure/result/accountability.
- Full Organization/Campaign self-service, curation, upload UI/API breadth, listing/discovery.
- Reopening the reconciled public contract absent new invalidating evidence.

## Dependencies

HARD:
- WU-S1-002 — `CONTRACT_READY` (satisfied).

SOFT / COORDINATION:
- WU-S1-005 for real private-storage/topology integration. Backend exploration/build may proceed before topology completion if storage ownership remains explicit and testable.

## Authority / Contract

- Product/MVP: `docs/product/mvp-delivery-slices.md` Slice 1.
- Delivery/spec/threat: reconciled Campaign Slice-1 artifacts.
- API: `getPublicCampaignDetail`, `getPublicCampaignMediaContent`.
- Backend architecture: `docs/project/kencleng-backend-tech-stack.md`.
- Shared rendezvous: `docs/project/kencleng-integration-map.md`.

## Human Gates

Honor root Tier-0 fencing. If implementation requires a protected write or a new material Product/security/interface decision, stop and escalate rather than routing around it.

## Completion

Complete only when backend runtime evidence satisfies the approved downstream Techplan and supports a truthful `BACKEND_VERIFIED` claim without requiring frontend or real-proxy integration to prove backend-owned behavior.

## Current Route

Exploration Run:
`EXP-BE-001` — PLANNED / DISPATCHED

Invocation:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/EXP-BE-001/invocation.md`


## Exploration State

Stage 1:
CONFIRMED

Stage 2:
COMPLETED

Stage 2 Assessment:
SOUND — no ownership collision or hidden hard dependency requiring re-decomposition.

Current Gate:
EXPLORATION_COMPLETED

Exploration Run:
EXP-BE-001 — COMPLETED

Planning Run:
`TP-BE-001` — PLANNED / DISPATCHED

Planning Session:
Continue existing healthy EXP-BE-001 session

Next Action:
Execute `runs/TP-BE-001/invocation.md`. Do not begin Build before the Techplan gate closes.


## Planning Review State

Draft Techplan:
`runs/TP-BE-001/techplan.md`

Independent Review:
`TPR-BE-001` — PLANNED / DISPATCHED

Review Session:
Fresh independent session

Model:
`gpt-5.6-terra / high`

Human Approval:
PENDING review convergence

Build:
NOT_AUTHORIZED


## Human Techplan Report State

Independent Review:
`TPR-BE-001` — COMPLETED / CLEAN

Previous Orchestrator-authored report:
INVALIDATED / REMOVED — role ownership and communication-profile violation.

Report Generation Run:
`TPRPT-BE-001` — PLANNED / DISPATCHED

Human Techplan Gate:
PENDING report generation.

Build:
NOT_AUTHORIZED

## Human Techplan Gate Ready

Report:
`runs/TP-BE-001/report-techplan.md`

Independent Review:
`TPR-BE-001` — CLEAN

Human Approval:
READY

Build:
NOT_AUTHORIZED until explicit Human approval.


## Human Techplan Gate — Ready

Current-effective Techplan:
`runs/TP-BE-001/techplan.md`

Review State:
`TPR-BE-001` — COMPLETE / CLEAN FOR HUMAN GATE

Human Review Report:
`runs/TP-BE-001/report-techplan.md`

Human Approval:
READY

Build:
NOT_AUTHORIZED until explicit Human approval of the current-effective Techplan.


## Build Dispatch

Human Techplan Approval:
APPROVED — current-effective `TP-BE-001`

Build Run:
`BLD-BE-001` — PLANNED / DISPATCHED

Build Session:
Fresh

Model:
`gpt-5.6-terra / high`

Next Gate:
CODE_REVIEW after Build completion.

Milestone:
NOT_YET_VERIFIED


## Code Review Dispatch

Build Run:
`BLD-BE-001` — COMPLETED

Code Review Run:
`CR-BE-001` — PLANNED / DISPATCHED

Review Session:
Fresh independent

Model:
`gpt-5.6-terra / high`

Next Gate:
Testing if review approves; Build/Patch if review requests changes.

Milestone:
NOT_YET_VERIFIED

## Testing Dispatch

Code Review:
`CR-BE-001` — APPROVE

Testing Run:
`TST-BE-001` — PLANNED / DISPATCHED

Session:
Fresh independent Testing

Model:
`gpt-5.6-terra / high`

Milestone:
NOT_YET_VERIFIED
