# CR-BE-001 — Code Review Invocation

WORK_UNIT_ID:
`WU-S1-003`

RUN_ID:
`CR-BE-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/CR-BE-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003`

ROLE:
`Reviewer`

SPECIALIZATION:
`Independent Code Review — Backend Campaign public delivery`

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh independent Code Review session

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
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/BLD-BE-001/report.md`

CURRENT_DIFF_SCOPE:
backend Campaign migrations/domain/storage/HTTP/server/seed changes produced by BLD-BE-001

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/4-code-review-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

Review the CURRENT implementation/diff, not merely the Build report.

Run all four passes:
1. Safety
2. Quality
3. Stack-specific best practices
4. Consistency with target-repo authority

Do not edit production code in this Run.
If a code change is required, create the specific patch plan required by the canonical Code Review workflow.

Additional environment note:
- `.harscode-spaces/.local-config.yaml` declares Podman + podman-compose as the Human local container runtime.
- Absence of the `docker` CLI alone is not evidence that local container/runtime verification is unavailable.
- Code Review remains reasoning-first; only use runtime commands when needed to prove/disprove a concrete review finding.

Write review findings under:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/CR-BE-001/`

Do not start Testing.
