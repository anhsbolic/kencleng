# EXP-TOP-001 — WU-S1-005 Exploration Invocation

HARSCODE_WORKSPACE_ROOT:
`../harscode-workspace`

WORK_UNIT_ID:
`WU-S1-005`

RUN_ID:
`EXP-TOP-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/EXP-TOP-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005`

ROLE:
`Explorer`

SPECIALIZATION:
`root Caddy proxy, Docker Compose, MinIO policy, same-origin /api routing, controlled Campaign media topology`

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
`medium`

MODEL_APPROVAL:
`NOT_REQUIRED`

PRIOR_ARTIFACTS:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TST-001/testing-report-1.md`
- `docs/project/kencleng-integration-map.md`

TASK:
Explore the minimum root topology and storage changes required for Slice 1 real integration from the reconciled CONTRACT_READY contract. Discover the actual Caddy /api path behavior, MinIO public/private policy state, backend storage assumptions, security implications, and the narrow viable enablement direction.

CODEBASE_CONTEXT:
Kencleng — local Caddy + Docker Compose PostgreSQL/MinIO topology supporting Go backend and Next.js frontend

AREA:
root Caddy proxy, Docker Compose, MinIO policy, same-origin /api routing, controlled Campaign media topology

## Canonical invocation

Run:
1. `../harscode-workspace/orchestration/exploration-kickoff-prompt.md`
2. canonical `workflow/1-exploration-kickoff-prompt.md`
3. `workflow/orchestrated-run-overlay.md`

Do not use this invocation as solution steering. The task/outcome and settled CONTRACT_READY boundary are inputs; Stage 1 must discover the relevant repo authority/areas and STOP for Human confirmation before Stage 2.

Do not alter WU/Control Surface orchestration state from the Explorer session.
