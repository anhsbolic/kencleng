# TST-FE-001 — Independent Frontend Testing Invocation

WORK_UNIT_ID:
`WU-S1-004`

RUN_ID:
`TST-FE-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TST-FE-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004`

ROLE:
`Verifier`

SPECIALIZATION:
Independent Testing — Public Campaign frontend mock-verified experience

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
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TP-FE-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/BLD-FE-001/report.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/CR-FE-001/review-findings.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/BLD-FE-PATCH-001/patch-report-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/CR-CONF-FE-001/review-findings.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/BLD-FE-PATCH-002/patch-report-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/CR-CONF-FE-002/review-findings.md`

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/5-testing-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

Testing goals:
- verify every FE Techplan rule through the observable UI/network boundary where possible;
- independently execute the exact MSW detail route and controlled-media handlers;
- cover loading, success, funding unavailable, media unavailable, safe 404, 503, generic request failure, retry recovery, and no-active-donation state;
- verify generated-type/request-path correspondence and no production mock branch;
- verify responsive behavior at representative desktop/mobile sizes;
- verify focus and announcement behavior for retryable transitions;
- verify hostile-looking organizer text remains plain rendered text;
- run target-repo required final frontend verification, including production build if Techplan assigns it here;
- perform Human-rendered-acceptance preparation, but do not mark the Human gate passed yourself.

Human Design wording:
- final Indonesian provenance/unavailable-action wording remains a Human-owned gate if still open;
- do not invent a new product decision to close it.

Do not edit production code.
Write:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TST-FE-001/testing-report-1.md`

If a production defect is found, create a Testing patch plan and route back to Build.
