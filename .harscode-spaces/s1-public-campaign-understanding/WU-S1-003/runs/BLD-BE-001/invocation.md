# BLD-BE-001 — Build Invocation

WORK_UNIT_ID:
`WU-S1-003`

RUN_ID:
`BLD-BE-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/BLD-BE-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003`

ROLE:
`Implementer`

SPECIALIZATION:
`Backend / Go / Campaign public delivery`

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh Build session

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
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TP-BE-001/techplan.md`

TASK:
Execute the Approved TP-BE-001 exactly within WU-S1-003 ownership. Implement the minimum persisted Campaign backend, explicit public projection, private MinIO read seam, public detail/media transport, operator-only seed command, and focused Build-owned verification. Do not touch Caddy/Compose/MinIO policy/frontend/OpenAPI/Product authority. Do not claim topology or integrated verification.

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/3-build-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

The current-effective Techplan is Human-approved and is the authoritative execution spine.

Before editing:
- re-open current live code/spec at Techplan anchors;
- stop if live code materially contradicts the Approved Techplan;
- do not reopen settled product/domain/architecture decisions.

Write the Build report to:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/BLD-BE-001/report.md`

Do not begin Code Review or Testing in this Run.
Do not update orchestration manifest/control surface yourself.
