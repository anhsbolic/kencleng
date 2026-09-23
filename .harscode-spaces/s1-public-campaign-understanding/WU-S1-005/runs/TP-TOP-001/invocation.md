# TP-TOP-001 — WU-S1-005 Techplan Synthesis Invocation

WORK_UNIT_ID:
`WU-S1-005`

RUN_ID:
`TP-TOP-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TP-TOP-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005`

ROLE:
`Planner`

SPECIALIZATION:
`Root topology / Caddy / MinIO`

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

PRIOR_ARTIFACTS:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/EXP-TOP-001/evidence/gap-analysis.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/EXP-TOP-001/evidence/solutioning.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TST-001/testing-report-1.md`
- `docs/project/kencleng-integration-map.md`

TASK:
Synthesize the execution-grade Techplan for the topology/private-media enablement direction selected by EXP-TOP-001: narrow /api prefix stripping at Caddy, explicit no-anonymous policy for the private bucket, and source/runtime verification boundaries. Preserve WU-S1-003 ownership of Campaign storage usage and media authorization semantics.

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

Write the Draft Techplan to:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TP-TOP-001/techplan.md`

Do not begin Build. At completion, report:
- material Active Open Items, if any;
- independent Techplan review recommendation;
- decomposition recommendation;
- session transition recommendation.

The Techplan must remain inside this Work Unit's ownership boundary. Cross-workstream dependencies should be explicit, but do not pull another Work Unit's implementation into this plan.
