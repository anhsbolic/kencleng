# TP-002 — Corrective Techplan Amendment for Frontend TypeScript Verification

WORK_UNIT_ID:
`WU-S1-002`

RUN_ID:
`TP-002`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-002`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002`

ROLE:
`Planner`

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
`gpt-5.6-terra`

REASONING_EFFORT:
`medium`

MODEL_APPROVAL:
`NOT_REQUIRED`

## Trigger

`BLD-001` completed all authorized reconciliation edits but could not satisfy the approved Build verification because:

`cd frontend && ./node_modules/.bin/tsc --noEmit`

fails on existing `app/page.test.tsx` Vitest globals (`describe`, `it`, `expect`).

Current live evidence:
- `frontend/vitest.config.ts` already uses `globals: true`;
- `frontend/tsconfig.json` includes test files but does not include Vitest global declarations;
- the failure predates and is unrelated to the generated OpenAPI artifact;
- `frontend/tsconfig.json` was outside the Approved TP-001 file-change boundary.

## Prior artifacts

- Approved source Techplan: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`
- Human report: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/report-techplan.md`
- Build finding: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-001/report.md`
- Live files: `frontend/tsconfig.json`, `frontend/vitest.config.ts`, `frontend/app/page.test.tsx`

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

This is a corrective Planning re-entry. Do not redo Exploration.

## Amendment objective

Amend the current-effective Techplan narrowly so Build may correct the TypeScript/Vitest typing configuration required to execute the already-approved `tsc --noEmit` verification.

Expected safe direction:
- preserve `tsc --noEmit` as the verification contract;
- authorize only the smallest project-consistent typing/configuration correction needed for existing Vitest globals;
- add `frontend/tsconfig.json` to the affected file boundary only if live evidence confirms that is the correct owner;
- preserve all existing product/security/interface/architecture decisions and Build scope;
- do not change production behavior, test semantics, or introduce new frontend feature work.

## Materiality rule

Classify the amendment explicitly.

If the correction only makes the already-approved verification executable and does not change material product/domain scope, authority/security, architecture/ownership, interface/data semantics, risk acceptance, or verification strategy:
- record it as a mechanical/non-material amendment;
- no independent re-review or new Human approval is required;
- identify the updated current-effective Techplan and recommend a narrow Build/Patch Run.

If live evidence requires a broader change or changing the verification strategy itself:
- STOP;
- classify the material concern;
- return to Human gate rather than silently expanding scope.

## Durable output

Write:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-002/amendment.md`

Update the authoritative Techplan only as allowed by Harscode guardrails, preserving the original approval history and making the amendment traceable.

Do not start Build/Patch in this session.
