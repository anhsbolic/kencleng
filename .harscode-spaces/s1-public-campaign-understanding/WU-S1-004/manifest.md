# WU-S1-004 — Slice 1 Frontend Public Campaign Understanding

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
Frontend / Next.js / Product surface

Communication Language:
Bahasa Indonesia

Communication Profile Path:
`docs/project/communication-profile.md`

## Outcome

Implement and independently verify the Slice-1 Public Campaign Detail experience against the reconciled contract so a skeptical-but-open visitor can understand campaign/steward/funding/provenance/media/action truth across required states before real backend integration.

Target project milestone:
`FRONTEND_MOCK_VERIFIED`

## Scope

- Public Campaign Detail route/surface derived from current page-map and Product Design/Brand authority.
- Typed production data boundary using the reconciled generated OpenAPI types.
- Contract-faithful MSW at the network boundary while backend is unavailable.
- Loading, success, missing-media, unavailable/failure, and not-public/not-found presentation states required by Slice 1.
- Truthful funding/provenance/lifecycle/action presentation; no fake Donate action.
- Responsive/accessibility behavior and human rendered acceptance appropriate to material UI work.
- Hostile-looking organizer plain-text rendering evidence where required.

## Out of Scope

- Backend implementation.
- Root topology/proxy/storage changes.
- Donation Flow or active donation CTA.
- Broad Campaign Discovery, full Organization Profile, Account, closure/result/accountability.
- Frontend-owned business recalculation of backend-authoritative eligibility/funding semantics.
- Reopening the reconciled API contract absent new invalidating evidence.

## Dependencies

HARD:
- WU-S1-002 — `CONTRACT_READY` (satisfied).

SOFT / COORDINATION:
- WU-S1-003 backend implementation is not required for mock-parallel development.
- WU-S1-005 topology is required only for later real same-origin integration.

## Authority / Contract

- Product/MVP: `docs/product/mvp-delivery-slices.md` Slice 1.
- Product Design/Brand: canonical `docs/ui-ux/**`.
- API/generated contract: current reconciled OpenAPI + generated TypeScript types.
- Frontend architecture: `docs/project/kencleng-frontend-tech-stack.md` + `frontend/AGENTS.md`.
- Shared rendezvous: `docs/project/kencleng-integration-map.md`.

## Human Gates

Human rendered acceptance remains required for material UI before final delivery. Final Indonesian provenance/action wording and any precedent-setting reusable placeholder asset must follow current Design authority/Human gate.

## Completion

Complete only when the frontend is independently verified against contract-faithful mocks and supports a truthful `FRONTEND_MOCK_VERIFIED` claim. Real backend/proxy integration remains WU-S1-006.

## Current Route

Exploration Run:
`EXP-FE-001` — PLANNED / DISPATCHED

Invocation:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/EXP-FE-001/invocation.md`


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
EXP-FE-001 — COMPLETED

Planning Run:
`TP-FE-001` — PLANNED / DISPATCHED

Planning Session:
Continue existing healthy EXP-FE-001 session

Next Action:
Execute `runs/TP-FE-001/invocation.md`. Do not begin Build before the Techplan gate closes.


## Planning Review State

Draft Techplan:
`runs/TP-FE-001/techplan.md`

Independent Review:
`TPR-FE-001` — PLANNED / DISPATCHED

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
`TPR-FE-001` — COMPLETED / CLEAN

Previous Orchestrator-authored report:
INVALIDATED / REMOVED — role ownership and communication-profile violation.

Report Generation Run:
`TPRPT-FE-001` — PLANNED / DISPATCHED

Human Techplan Gate:
PENDING report generation.

Build:
NOT_AUTHORIZED

## Human Report Localization Correction

Prior Report Generation:
`TPRPT-FE-001` — COMPLETED

Finding:
Human-facing report retained ordinary English headings despite Bahasa Indonesia communication profile.

Correction Run:
`TPRPT-FE-002` — PLANNED / DISPATCHED

Human Techplan Gate:
PENDING corrected report.

Build:
NOT_AUTHORIZED

## Human Techplan Gate Ready

Report:
`runs/TP-FE-001/report-techplan.md`

Independent Review:
`TPR-FE-001` — CLEAN

Human Approval:
READY

Build:
NOT_AUTHORIZED until explicit Human approval.


## Human Techplan Gate — Ready

Current-effective Techplan:
`runs/TP-FE-001/techplan.md`

Review State:
`TPR-FE-001` — COMPLETE / CLEAN FOR HUMAN GATE

Human Review Report:
`runs/TP-FE-001/report-techplan.md`

Human Approval:
READY

Build:
NOT_AUTHORIZED until explicit Human approval of the current-effective Techplan.


## Build Dispatch

Human Techplan Approval:
APPROVED — current-effective `TP-FE-001`

Build Run:
`BLD-FE-001` — PLANNED / DISPATCHED

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
`BLD-FE-001` — COMPLETED

Code Review Run:
`CR-FE-001` — PLANNED / DISPATCHED

Review Session:
Fresh independent

Model:
`gpt-5.6-terra / high`

Next Gate:
Testing if review approves; Build/Patch if review requests changes.

Milestone:
NOT_YET_VERIFIED

## Review Patch Dispatch

Code Review:
`CR-FE-001` — REQUEST_CHANGES

Blocking Findings:
F1 MSW worker lifecycle; F2 retryable async announcement/focus recovery.

Patch Run:
`BLD-FE-PATCH-001` — PLANNED / DISPATCHED

Testing:
BLOCKED pending patch + review confirmation.

Milestone:
NOT_YET_VERIFIED
