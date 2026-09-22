# EXP-BE-001 — WU-S1-003 Exploration Invocation

HARSCODE_WORKSPACE_ROOT:
`../harscode-workspace`

WORK_UNIT_ID:
`WU-S1-003`

RUN_ID:
`EXP-BE-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/EXP-BE-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003`

ROLE:
`Explorer`

SPECIALIZATION:
`backend Campaign domain, persistence, HTTP transport, controlled media delivery, seeded/operator-assisted data`

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh session

COMMUNICATION_LANGUAGE:
Bahasa Indonesia

COMMUNICATION_PROFILE_PATH:
`docs/project/communication-profile.md`

SELECTED_MODEL:
`gpt-5.6-terra`

REASONING_EFFORT:
`high`

MODEL_APPROVAL:
`NOT_REQUIRED`

PRIOR_ARTIFACTS:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TST-001/testing-report-1.md`
- `docs/project/kencleng-integration-map.md`

TASK:
Explore the implementation work required to deliver the backend-owned portion of Slice 1 Public Campaign Understanding from the reconciled CONTRACT_READY public Campaign contract. Do not redesign the settled public contract; discover current backend/data/storage state, implementation gaps, risks, and viable delivery direction.

CODEBASE_CONTEXT:
Kencleng — Go domain-driven monolith backend + PostgreSQL/MinIO + reconciled Slice-1 OpenAPI

AREA:
backend Campaign domain, persistence, HTTP transport, controlled media delivery, seeded/operator-assisted data

## Canonical invocation

Run:
1. `../harscode-workspace/orchestration/exploration-kickoff-prompt.md`
2. canonical `workflow/1-exploration-kickoff-prompt.md`
3. `workflow/orchestrated-run-overlay.md`

Do not use this invocation as solution steering. The task/outcome and settled CONTRACT_READY boundary are inputs; Stage 1 must discover the relevant repo authority/areas and STOP for Human confirmation before Stage 2.

Do not alter WU/Control Surface orchestration state from the Explorer session.
