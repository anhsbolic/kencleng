# TST-BE-001 — Independent Testing Invocation

WORK_UNIT_ID:
`WU-S1-003`

RUN_ID:
`TST-BE-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TST-BE-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003`

ROLE:
`Verifier`

SPECIALIZATION:
Independent Testing — Backend Campaign public delivery

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh independent Testing session

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
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/CR-BE-001/review-findings.md`

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/5-testing-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

Important local runtime context:
- `.harscode-spaces/.local-config.yaml` declares Podman as the Human local container runtime.
- Do not treat absence of Docker CLI as container-runtime unavailability.
- For testcontainers-go, explicitly probe Podman Docker-compatible socket/API compatibility. Do not assume it.
- If sandbox permissions prevent Podman/testcontainers despite a valid local runtime, distinguish sandbox limitation from Human-environment capability.

Priority gaps from Build/Review:
- isolated Postgres migration apply/down/re-up + constraint/repository behavior;
- isolated MinIO adapter behavior;
- anti-enumeration/timing evidence;
- cancellation/error handling;
- broad regression + race according to Techplan;
- direct backend observable contract.
Topology/Caddy/private-policy/browser integration remain outside this WU.

Do not edit production code. If Testing finds a production defect, write the Testing patch plan and route back to Build.

Write:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TST-BE-001/testing-report-1.md`
