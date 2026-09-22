# CR-001 — Independent Code Review for Slice 1 Contract Reconciliation

WORK_UNIT_ID:
`WU-S1-002`

RUN_ID:
`CR-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/CR-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002`

ROLE:
`Reviewer`

SPECIALIZATION:
None

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh independent session

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

## Review target

Review the complete current diff for `WU-S1-002` from the Build start baseline through the latest `BLD-002` patch.

Do not infer correctness from Build reports. Inspect the actual current repository state and diff.

## Prior artifacts

Authoritative current-effective Techplan:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`

Corrective amendment:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-002/amendment.md`

Build evidence:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-001/report.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-002/patch-report-1.md`

The Build reports are orientation/evidence only; they are not substitutes for reviewing the current diff.

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/4-code-review-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

Run all four canonical passes:
1. Safety
2. Quality
3. Stack-specific best practices
4. Consistency with target-repo authority

## Review scope / boundaries

The expected diff includes reconciliation-only changes authorized by the current-effective Techplan:
- Campaign specs/tasks/invariant/threat model;
- split OpenAPI + generated aggregate;
- generated frontend TypeScript contract artifact;
- frontend API type generation script;
- narrow `frontend/tsconfig.json` Vitest-global typing correction from TP-002;
- backend architecture wording;
- integration map;
- any tracker change that actually exists in current diff.

Explicitly verify that:
- no production backend/frontend runtime implementation was introduced;
- no Product/MVP or Product Design authority was changed;
- no unauthorized topology/storage implementation was introduced;
- public projection remains closed/allowlisted;
- public detail/media anti-enumeration and no-store semantics remain coherent;
- generated artifacts correspond to their authored sources;
- TP-002 stayed mechanical/non-material and did not broaden runtime behavior;
- tracker/milestone claims, if present, do not overstate evidence.

Do not load raw Exploration logs unless the Techplan itself is missing material intent; if so, report Techplan drift rather than reconstructing intent from history.

## Verification posture

Review is reasoning-first. Run targeted commands only to prove/disprove a concrete suspected finding. Do not replay the full Build/Testing matrix for ceremony.

Do not edit production or reconciliation artifacts in this Run.

If changes are required, write a patch plan under this Run and route back to Build/Patch.

## Durable output

Write:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/CR-001/review-findings-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/CR-001/patch-plan-1.md` only if changes are required.

Stop after Code Review handoff. Do not continue into Testing.
