# BLD-002 — Narrow Build/Patch for Frontend TypeScript Verification

WORK_UNIT_ID:
`WU-S1-002`

RUN_ID:
`BLD-002`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-002`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002`

ROLE:
`Implementer`

SPECIALIZATION:
None

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh session

COMMUNICATION_LANGUAGE:
Bahasa Indonesia

COMMUNICATION_PROFILE_PATH:
`docs/project/communication-profile.md`

SELECTED_MODEL:
`gpt-5.6-luna`

REASONING_EFFORT:
`medium`

MODEL_APPROVAL:
`NOT_REQUIRED`

## Trigger

`BLD-001` completed the authorized reconciliation edits but could not satisfy the approved frontend TypeScript verification. `TP-002` verified and authorized a mechanical/non-material correction only.

## Prior artifacts

Current-effective Approved Techplan:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`

Corrective amendment:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-002/amendment.md`

Build finding:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-001/report.md`

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/3-build-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

This is a Build/Patch re-entry requested by corrective planning.

## Patch target

Make exactly this authorized correction:

- in `frontend/tsconfig.json`, add `"types": ["vitest/globals"]` under `compilerOptions`.

Do not modify:
- `frontend/app/page.test.tsx`;
- `frontend/vitest.config.ts`;
- production frontend code;
- API/spec/generated artifacts;
- tracker;
- Work Unit manifest / Control Surface;
- any other file unless an unexpected mechanical formatting consequence makes the single authorized edit impossible.

If the exact correction does not make the approved check pass, STOP and report instead of expanding scope.

## Verification

Run only the checks authorized by `TP-002`:

1. `cd frontend && ./node_modules/.bin/tsc --noEmit`
2. `cd frontend && npm run lint`
3. repository-root `git diff --check`

Do not rerun broad API/build/testing suites in this patch round; `BLD-001` evidence remains current for unchanged artifacts.

## Durable output

Write:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-002/patch-report-1.md`

The report must state whether the patch resolved the `BLD-001` finding and whether the Work Unit can now proceed to independent Code Review. Do not claim `CONTRACT_READY`; Human milestone acceptance remains later.

Do not continue into Code Review or Testing in this session.
