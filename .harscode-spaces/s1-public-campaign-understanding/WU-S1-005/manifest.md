# WU-S1-005 — Slice 1 Topology & Controlled Media Enablement

Type:
ENABLER

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
Root topology / Caddy / MinIO

Communication Language:
Bahasa Indonesia

Communication Profile Path:
`docs/project/communication-profile.md`

## Outcome

Establish the minimum local topology required for truthful Slice-1 integration: same-origin `/api` routing reaches backend routes with the expected path shape, and Campaign media storage is not anonymously public while controlled delivery remains the public byte path.

## Scope

- Investigate and correct the known Caddy `/api` prefix-handling mismatch if confirmed.
- Reconcile local MinIO bucket/policy behavior needed for private Campaign media.
- Preserve unrelated storage use-cases unless current authority requires a narrower change.
- Define/verify the minimum topology evidence required for backend controlled-media integration.
- Keep topology behavior compatible with the reconciled same-origin `content_url` contract.

## Out of Scope

- Campaign business/domain implementation.
- Frontend product UI.
- Broad deployment/cloud architecture.
- Redesign of all storage policy or historical upload behavior.
- Donation/Account/closure/accountability topology.

## Dependencies

HARD:
- WU-S1-002 — `CONTRACT_READY` (satisfied).

SOFT / COORDINATION:
- Coordinate with WU-S1-003 storage expectations.
- Does not block WU-S1-004 mock-parallel frontend work.

## Authority / Contract

- Root `AGENTS.md` and repo setup/topology authority.
- Reconciled media contract and Campaign threat model.
- `Caddyfile`, `docker-compose.yml`, current storage/platform evidence.
- `docs/project/kencleng-integration-map.md`.

## Human Gates

If a topology change creates a new material security/exposure decision beyond the reconciled private-media/same-origin requirement, stop and escalate.

## Completion

Complete when topology-owned evidence proves the local proxy/storage boundary required by Slice 1 is ready for real integration. This Work Unit does not itself claim `BACKEND_VERIFIED`, `FRONTEND_MOCK_VERIFIED`, or `INTEGRATED_VERIFIED`.

## Current Route

Exploration Run:
`EXP-TOP-001` — PLANNED / DISPATCHED

Invocation:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/EXP-TOP-001/invocation.md`


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
EXP-TOP-001 — COMPLETED

Planning Run:
`TP-TOP-001` — PLANNED / DISPATCHED

Planning Session:
Fresh session

Next Action:
Execute `runs/TP-TOP-001/invocation.md`. Do not begin Build before the Techplan gate closes.


## Planning Review State

Draft Techplan:
`runs/TP-TOP-001/techplan.md`

Independent Review:
`TPR-TOP-001` — PLANNED / DISPATCHED

Review Session:
Fresh independent session

Model:
`gpt-5.6-terra / medium`

Human Approval:
PENDING review convergence

Build:
NOT_AUTHORIZED


## Planning Review Resolution

Independent Review:
`TPR-TOP-001` — COMPLETED / 1 MATERIAL BLOCKING finding

Blocking Finding:
Missing executable retraction-through-Caddy verification coverage.

Resolution Run:
`TPR-RES-TOP-001` — PLANNED / DISPATCHED

Human Approval:
BLOCKED until resolution converges.

Build:
NOT_AUTHORIZED

## Targeted Review Confirmation

Resolution Run:
`TPR-RES-TOP-001` — COMPLETED

Resolution Assessment:
No material scope/architecture/ownership/security/interface semantic change; existing verification obligation made executable.

Confirmation Run:
`TPR-CONF-TOP-001` — PLANNED / DISPATCHED

Human Approval:
BLOCKED until targeted confirmation closes the prior finding.

Build:
NOT_AUTHORIZED

## Human Report Generation

Independent Review:
`TPR-TOP-001` — blocking finding resolved.

Resolution:
`TPR-RES-TOP-001` — COMPLETED

Targeted Confirmation:
`TPR-CONF-TOP-001` — CONFIRMED_CLOSED

Report Generation Run:
`TPRPT-TOP-001` — PLANNED / DISPATCHED

Human Approval:
PENDING report generation.

Build:
NOT_AUTHORIZED


## Human Techplan Gate — Ready

Current-effective Techplan:
`runs/TP-TOP-001/techplan.md`

Review State:
`TPR-CONF-TOP-001` — COMPLETE / CLEAN FOR HUMAN GATE

Human Review Report:
`runs/TP-TOP-001/report-techplan.md`

Human Approval:
READY

Build:
NOT_AUTHORIZED until explicit Human approval of the current-effective Techplan.


## Build Dispatch

Human Techplan Approval:
APPROVED — current-effective `TP-TOP-001`

Build Run:
`BLD-TOP-001` — PLANNED / DISPATCHED

Build Session:
Fresh

Model:
`gpt-5.6-terra / medium`

Next Gate:
CODE_REVIEW after Build completion.

Milestone:
NOT_YET_VERIFIED


## Code Review Dispatch

Build Run:
`BLD-TOP-001` — COMPLETED

Code Review Run:
`CR-TOP-001` — PLANNED / DISPATCHED

Review Session:
Fresh independent

Model:
`gpt-5.6-terra / medium`

Next Gate:
Testing if review approves; Build/Patch if review requests changes.

Milestone:
NOT_YET_VERIFIED

## Review Patch Dispatch

Code Review:
`CR-TOP-001` — REQUEST_CHANGES

Blocking Finding:
S1 MinIO initializer fail-fast behavior.

Patch Run:
`BLD-TOP-PATCH-001` — PLANNED / DISPATCHED

Testing:
BLOCKED pending patch + targeted review confirmation.

Milestone:
NOT_YET_VERIFIED
