# BLD-FE-PATCH-001 — Code Review Patch Invocation

WORK_UNIT_ID:
`WU-S1-004`

RUN_ID:
`BLD-FE-PATCH-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/BLD-FE-PATCH-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004`

ROLE:
`Implementer`

SPECIALIZATION:
Narrow Build/Patch — frontend MSW lifecycle and async accessibility recovery

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh Build/Patch session

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
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TP-FE-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/CR-FE-001/review-findings.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/CR-FE-001/patch-plan.md`

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/3-build-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

PATCH RE-ENTRY:
Implement only the two accepted review findings F1 and F2 exactly as defined by the patch plan.

Do not broaden scope.
Do not change backend/topology/OpenAPI/Product authority.
Do not start Testing.

Required focused verification:
- lifecycle provider tests for worker start/stop/unmount race;
- retryable 503/request-failure announcement + focus recovery tests;
- `cd frontend && npm run verify`.

Write:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/BLD-FE-PATCH-001/patch-report-1.md`

After patch, return to the requesting Code Review phase for targeted confirmation unless the patch materially broadens scope.
