# BLD-001 — Slice 1 Public Contract & Delivery Reconciliation Build

WORK_UNIT_ID:
`WU-S1-002`

RUN_ID:
`BLD-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002`

ROLE:
`Implementer`

SPECIALIZATION:
None

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
`high`

MODEL_APPROVAL:
`NOT_REQUIRED`

TARGET_REVISION:
Resolve and record current `validation-03-orchestrator-slice-1` HEAD at Run start.

WORKFLOW_REVISION:
Resolve and record current `pilot/orchestrator-v0.1` HEAD at Run start.

## Prior artifacts

Current-effective Approved Techplan:

`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`

Human review report (derived, non-authoritative):

`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/report-techplan.md`

Independent review/resolution history is already incorporated into the Approved Techplan. Do not reload historical Exploration/review artifacts unless the Techplan explicitly requires an exact evidence pointer.

## Canonical workflow

Use:

1. `../harscode-workspace/workflow/3-build-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

The canonical Build prompt owns Build behavior. The overlay owns Run identity/path semantics.

## Build target

Execute the Approved `TP-001` reconciliation contract in full.

This Run is limited to reconciliation artifacts already authorized by the Techplan, including:

- Campaign delivery/spec/invariant/threat reconciliation;
- touched split OpenAPI contract and root `bearerAuth` fix;
- regenerated aggregate OpenAPI;
- generated frontend TypeScript contract artifact;
- narrow backend architecture wording;
- integration map and tracker reconciliation;
- Build-owned verification required by the Techplan.

Explicitly do **not** implement production backend/frontend runtime capability, migrations, route handlers, storage/topology changes, Donation Flow, Slice 3 closure behavior, or Product/Design authority changes.

Decomposition is skipped: the Approved plan is cohesive and dependency-linear; splitting it would create partial contract reconciliation without meaningful execution-context benefit.

## Execution constraints

- Re-ground on live repo state before editing.
- Treat the Approved Techplan as the execution authority.
- Do not reopen settled product/security/interface decisions.
- If live code/spec contradicts a material Techplan assumption, stop and report a Finding instead of redesigning.
- Do not mutate Work Unit manifest, Control Surface, or other Orchestrator-owned state.
- Update the tracker last and claim `CONTRACT_READY` only if the Techplan's required evidence has actually passed.
- Keep unrelated historical OpenAPI warnings/backlog out of scope.
- Use only focused Build verification assigned to Build by the Techplan/canonical Build guidance; leave independent Testing-owned evidence for the later Testing Run.

## Durable output

Write the Build report to:

`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-001/report.md`

Follow the canonical Build report/handoff format and include Run/Work Unit provenance where the artifact shape permits it.

Do not continue into Code Review or Testing in this session. Stop after the Build handoff.
