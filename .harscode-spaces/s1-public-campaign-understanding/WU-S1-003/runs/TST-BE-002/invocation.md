# TST-BE-002 — Independent Testing Re-entry

WORK_UNIT_ID:
`WU-S1-003`

RUN_ID:
`TST-BE-002`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TST-BE-002`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003`

ROLE:
`Verifier`

SPECIALIZATION:
Independent Testing re-entry — backend Campaign verification closure

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
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TP-BE-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TST-BE-001/testing-report-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TST-BE-001/patch-plan-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/BLD-BE-PATCH-001/patch-report-1.md`

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/5-testing-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

This is Testing re-entry after a narrow patch. Verify affected gaps first; do not ceremonially repeat unrelated earlier evidence.

Required closure checks:
1. Run the combined integration-tag Campaign Postgres + MinIO tests through the configured Podman Docker-compatible socket/API.
2. Independently verify the full R5 absent/non-published/non-member × Authorization matrix and deterministic representative timing-parity evidence.
3. Re-check Campaign-scope security findings are cleared.
4. Run target-repo final verification as far as authoritative tooling allows.
   - If `make verify` remains blocked exclusively by unrelated pre-existing findings outside WU-S1-003, distinguish that explicitly from Campaign correctness.
   - Do not convert unrelated pre-existing findings into a Campaign failure without an authority reason.
5. Reassess whether all WU-S1-003 Testing-owned rules are now closed enough for `BACKEND_VERIFIED`, while keeping topology/integration claims out of scope.

Do not edit production code.
Write:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TST-BE-002/testing-report-1.md`

If a new production defect is found, create the Testing patch plan and return to Build.
