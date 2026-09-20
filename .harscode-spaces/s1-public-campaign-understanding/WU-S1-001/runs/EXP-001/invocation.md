# EXP-001 — Orchestrated Exploration Invocation

Status:
READY_FOR_REVIEW

Prepared By Role:
Orchestration Operator

Work Unit:
WU-S1-001

Run:
EXP-001

Role:
Explorer

Specialization:
None

Run Path:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001`

Work Unit Path:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-001`

Prior Artifacts:
None

Target Branch:
`validation-03-orchestrator-slice-1`

Workflow Branch:
`pilot/orchestrator-v0.1`

## Operator Binding Before Dispatch

Bind:

`HARSCODE_WORKSPACE_ROOT`

to the local checkout of `harscode-workspace` currently on:

`pilot/orchestrator-v0.1`

This is an environment/path binding only. Do not rewrite the task or add solution-steering instructions.

## Entrypoint

Use:

`{HARSCODE_WORKSPACE_ROOT}/orchestration/exploration-kickoff-prompt.md`

That wrapper must apply:

- canonical Exploration authority from `workflow/1-exploration-kickoff-prompt.md`;
- orchestrated identity/path semantics from `workflow/orchestrated-run-overlay.md`.

## Codebase Context

`Kencleng — Go backend + Next.js frontend`

## Task

Explore what is required to deliver Kencleng MVP Slice 1 — Public Campaign Understanding — based on the current authoritative Product and Product Design sources and the existing project specifications, shared contracts, backend/frontend implementation, and other relevant current evidence.

The governing slice source is:

`docs/product/mvp-delivery-slices.md`

Do not assume historical specs, contracts, migrations, tests, or implementation are automatically current authority. Surface material gaps or conflicts through the authority that owns the concern.

Do not assume the downstream Work Unit graph, frontend/backend split, or contract changes in advance. Derive only what current authority and evidence support.

## Optional Routing Inputs

Ticket:
None

Area:
Not sure yet — Stage 1 must determine the relevant areas.

## Dispatch Contract

Run the canonical Exploration contract with:

- `WORK_UNIT_ID = WU-S1-001`
- `RUN_ID = EXP-001`
- `RUN_PATH = .harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001`
- `WORK_UNIT_PATH = .harscode-spaces/s1-public-campaign-understanding/WU-S1-001`
- `ROLE = Explorer`
- `SPECIALIZATION = none`
- `PRIOR_ARTIFACTS = none`
- `TASK = the Task section above`
- `CODEBASE_CONTEXT = Kencleng — Go backend + Next.js frontend`

## Mandatory First Stop

Execute **Stage 1 — Plan Announcement only**.

Do not proceed to Stage 2 until the human checkpoint confirms the Stage 1 understanding and exploration routing.

## Invocation Review Checklist

Before dispatch, verify:

- [ ] Correct Work Unit selected.
- [ ] Exploration is the correct next workflow phase.
- [ ] Explorer is the correct Role.
- [ ] No downstream solution / Work Unit hypothesis is injected as fact.
- [ ] Product / Design authority remains upstream of historical implementation evidence.
- [ ] Run Path is unique to EXP-001.
- [ ] Prior Artifacts correctly equals `none`.
- [ ] The only operator-specific mutation is binding `HARSCODE_WORKSPACE_ROOT`.
