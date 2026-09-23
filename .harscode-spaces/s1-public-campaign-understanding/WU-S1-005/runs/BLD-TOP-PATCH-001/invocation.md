# BLD-TOP-PATCH-001 — Code Review Patch Invocation

WORK_UNIT_ID:
`WU-S1-005`

RUN_ID:
`BLD-TOP-PATCH-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/BLD-TOP-PATCH-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005`

ROLE:
`Implementer`

SPECIALIZATION:
Narrow Build/Patch — MinIO initializer fail-fast behavior

PARTICIPANT:
Codex CLI agent

SESSION:
Continue/re-ground on the healthy topology Build context if available; otherwise fresh focused Build/Patch session

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
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TP-TOP-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/CR-TOP-001/review-findings-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/CR-TOP-001/patch-plan-1.md`

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/3-build-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

PATCH RE-ENTRY:
Apply only the accepted S1 patch plan:
- make `minio-init` fail-fast (`/bin/sh -ec` or exactly equivalent);
- preserve retrying `until mc alias set ...`;
- do not change policy order/meaning or any unrelated topology source.

Local runtime:
- Podman/podman-compose is the configured Human local runtime.
- If writable Podman is available, rendered config may be checked.
- Sandbox inability to access `/run/user/1000/libpod` is an environment limitation, not proof the Human local runtime is unavailable.

Do not start Testing.

Write:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/BLD-TOP-PATCH-001/patch-report-1.md`

After patch, return to CR-TOP-001 via targeted confirmation; full four-pass re-review is unnecessary unless the patch broadens scope or changes semantics.
