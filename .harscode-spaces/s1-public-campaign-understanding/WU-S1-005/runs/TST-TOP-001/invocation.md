# TST-TOP-001 — Independent Testing Invocation

WORK_UNIT_ID:
`WU-S1-005`

RUN_ID:
`TST-TOP-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TST-TOP-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005`

ROLE:
`Verifier`

SPECIALIZATION:
Independent Testing — root topology and controlled media enablement

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh independent Testing session

COMMUNICATION_LANGUAGE:
Bahasa Indonesia

SELECTED_MODEL:
`gpt-5.6-terra`

REASONING_EFFORT:
`high`

MODEL_APPROVAL:
`NOT_REQUIRED`

PRIOR_ARTIFACTS:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TP-TOP-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/BLD-TOP-001/report.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/BLD-TOP-PATCH-001/patch-report-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/CR-CONF-TOP-001/review-findings.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/BLD-BE-PATCH-001/patch-report-1.md`

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/5-testing-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

Local runtime:
- Podman is the configured Human local container engine.
- Prefer repository runner/targets such as `make up-podman` / `make down-podman`.
- Absence of Docker CLI is irrelevant if Podman path works.
- If this sandbox cannot access Podman, distinguish sandbox limitation from Human environment capability.

Testing priorities:
- rendered/root Caddy configuration and /api path translation;
- non-API frontend fallback;
- persisted MinIO public/private anonymous policy on fresh boot and volume reuse;
- direct anonymous private object denial;
- controlled media 200/404/503 header/error preservation through Caddy where WU-S1-003 capability now permits;
- R7 fresh retraction request to the same known content_url after parent eligibility or media membership withdrawal, when fixture/mutation capability is available;
- source/runtime scope consistency and no false integrated claim.

Do not edit production code.
Write:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TST-TOP-001/testing-report-1.md`

If a production defect is found, create a Testing patch plan and route back to Build.
