# TP-BE-001 — WU-S1-003 Techplan Synthesis Invocation

WORK_UNIT_ID:
`WU-S1-003`

RUN_ID:
`TP-BE-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TP-BE-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003`

ROLE:
`Planner`

SPECIALIZATION:
`Backend / Go / Campaign domain`

PARTICIPANT:
Codex CLI agent

SESSION:
Continue existing healthy EXP-BE-001 session

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
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/EXP-BE-001/evidence/gap-analysis.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/EXP-BE-001/evidence/solutioning.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TST-001/testing-report-1.md`
- `docs/project/kencleng-integration-map.md`

TASK:
Synthesize the execution-grade Techplan for the backend-owned Slice 1 Public Campaign Delivery direction selected by EXP-BE-001. Preserve the CONTRACT_READY public API/security semantics and WU-S1-005 topology ownership boundary. The plan must be sufficient to earn BACKEND_VERIFIED without claiming topology or frontend integration evidence.

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

Write the Draft Techplan to:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TP-BE-001/techplan.md`

Do not begin Build. At completion, report:
- material Active Open Items, if any;
- independent Techplan review recommendation;
- decomposition recommendation;
- session transition recommendation.

The Techplan must remain inside this Work Unit's ownership boundary. Cross-workstream dependencies should be explicit, but do not pull another Work Unit's implementation into this plan.
