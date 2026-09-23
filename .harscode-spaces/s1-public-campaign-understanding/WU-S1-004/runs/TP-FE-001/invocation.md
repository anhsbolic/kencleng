# TP-FE-001 — WU-S1-004 Techplan Synthesis Invocation

WORK_UNIT_ID:
`WU-S1-004`

RUN_ID:
`TP-FE-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TP-FE-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004`

ROLE:
`Planner`

SPECIALIZATION:
`Frontend / Next.js / Public Campaign Detail`

PARTICIPANT:
Codex CLI agent

SESSION:
Continue existing healthy EXP-FE-001 session

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
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/EXP-FE-001/evidence/gap-analysis.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/EXP-FE-001/evidence/solutioning.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TST-001/testing-report-1.md`
- `docs/project/kencleng-integration-map.md`

TASK:
Synthesize the execution-grade Techplan for the frontend-owned Slice 1 Public Campaign Understanding direction selected by EXP-FE-001. Preserve generated-type + real same-origin request + MSW network interception as one contract path, and define evidence sufficient for FRONTEND_MOCK_VERIFIED without claiming real backend/topology integration.

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

Write the Draft Techplan to:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TP-FE-001/techplan.md`

Do not begin Build. At completion, report:
- material Active Open Items, if any;
- independent Techplan review recommendation;
- decomposition recommendation;
- session transition recommendation.

The Techplan must remain inside this Work Unit's ownership boundary. Cross-workstream dependencies should be explicit, but do not pull another Work Unit's implementation into this plan.
