# BLD-FE-PATCH-002 — Targeted Review Re-entry Patch

WORK_UNIT_ID:
`WU-S1-004`

RUN_ID:
`BLD-FE-PATCH-002`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/BLD-FE-PATCH-002`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004`

ROLE:
`Implementer`

SPECIALIZATION:
Narrow Build/Patch — React Strict Mode MSW ownership + request-failure recovery evidence

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh focused Build/Patch session

COMMUNICATION_LANGUAGE:
Bahasa Indonesia

SELECTED_MODEL:
`gpt-5.6-terra`

REASONING_EFFORT:
`high`

MODEL_APPROVAL:
`NOT_REQUIRED`

PRIOR_ARTIFACTS:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TP-FE-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/CR-CONF-FE-001/review-findings.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/BLD-FE-PATCH-001/patch-report-1.md`

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/3-build-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

PATCH RE-ENTRY:
Close only F1-R and F2-R.

F1-R:
- coordinate ownership of the singleton browser MSW worker across overlapping/Strict Mode effect lifecycles;
- stale startup completion must not stop the worker owned by a newer active effect;
- preserve route-local interception and eventual stop when no active owner remains;
- add focused overlap/Strict Mode lifecycle test proving the mounted owner retains interception.

F2-R:
- add observable coverage for generic `request-failure → retry → success`;
- assert successful completion and correct focus restoration, not merely retry click behavior;
- preserve initial-success no-focus-shift behavior and existing safe 404 semantics.

Do not redesign the mock architecture unless strictly required to close the two findings.
Do not touch backend/topology/OpenAPI/Product authority.

Required verification:
- focused lifecycle + recovery tests;
- `cd frontend && npm run verify`;
- `git diff --check`.

Do not start Testing.
Write:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/BLD-FE-PATCH-002/patch-report-1.md`

After patch, return to a fresh targeted Code Review confirmation only.
