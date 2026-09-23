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
AWAITING_STAGE_3_CONFIRMATION

Next Action:
Continue the existing Exploration Run to Stage 3 — Solutioning. Preserve Work Unit ownership boundaries and settled CONTRACT_READY semantics.
