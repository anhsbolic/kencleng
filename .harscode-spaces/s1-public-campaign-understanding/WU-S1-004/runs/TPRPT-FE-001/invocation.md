# TPRPT-FE-001 — Human Techplan Report Generation

WORK_UNIT_ID:
`WU-S1-004`

RUN_ID:
`TPRPT-FE-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TPRPT-FE-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004`

ROLE:
`Planner`

SPECIALIZATION:
Human-facing Techplan report synthesis

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh focused session

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
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TP-FE-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TPR-FE-001/review-findings.md`

## Task

Generate the Human-facing Techplan review report for the current-effective Draft Techplan after independent review convergence.

Use:
- `../harscode-workspace/workflow/2-techplan/report-template.md`
- current target-project communication profile.

The report is a derived digest only. It MUST NOT introduce, resolve, weaken, or reinterpret any Techplan decision.

Human-facing prose MUST use Bahasa Indonesia.
Preserve canonical Harscode terms/enums, code/API/schema identifiers, file paths, branch/commit identifiers, and exact authority wording when wording matters.

## Source of truth

Current-effective Techplan:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TP-FE-001/techplan.md`

Independent review:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TPR-FE-001/review-findings.md`

If report and Techplan would disagree, the Techplan wins; fix the report, not the Techplan.

## Required output

Write:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TP-FE-001/report-techplan.md`

The report must:
- accurately summarize scope/out-of-scope;
- explain architecture/plan at Human-review level;
- include reviewer-relevant interface/cross-boundary contract only where material;
- carry material decisions, risks/trade-offs, and independent review history;
- distinguish actual approval blockers from non-blocking Human/external follow-up;
- state exactly what Human approval authorizes and does not authorize;
- keep all explanatory prose in Bahasa Indonesia.

Do not modify:
- `techplan.md`;
- review findings;
- implementation/source code;
- orchestration manifest/control surface.

Do not start Build.
