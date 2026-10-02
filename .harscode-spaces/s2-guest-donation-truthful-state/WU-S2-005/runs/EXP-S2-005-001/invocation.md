# Run Invocation — `EXP-S2-005-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator after Anhar selected Option 1 for the Campaign/API contract coordination route. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-005`
- `RUN_ID`: `EXP-S2-005-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005`
- `ROLE`: Explorer
- `SPECIALIZATION`: None; Campaign/API contract reconciliation Exploration
- `PARTICIPANT_ID`: `P-S2-005-EXP-001-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-EXPLORER`; Profile SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `PRIOR_ARTIFACTS`: WU-S2-005 manifest; current Parent Outcome and Work Graph; approved Product/MVP Slice 2 sources; Campaign Stage-3 solutioning `WU-S2-004/runs/EXP-S2-004-001/evidence/stage-3-solutioning.md`; current Campaign feature/invariant and authored Campaign OpenAPI; accepted Donation specs and authored Donation API where needed. Re-read live authorities at dispatch; do not load historical Run archives wholesale.
- `SESSION_TRANSITION`: `FRESH` — new reconciliation Work Unit and new canonical Exploration Run.
- `SESSION_TRANSITION_REASON`: A separate Campaign/API contract reconciliation is required by Anhar's explicit Option 1 decision; it has a distinct shared-contract scope and needs fresh Explorer context.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796`; current working tree contains Stage-2/Stage-3 evidence and orchestration-state reconciliation, with no reported production-source changes. Re-read relevant sources live.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1` HEAD `95ecf37ba8ae449a5b3b278c27331aca87360bc8`; current-effective provenance, not semantic pinning.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; local Human-owned config marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`.
- `PHASE_ROUTE`: Canonical Exploration kickoff, Stage 1 only initially. Stop for Human confirmation before Stage 2; Stage 2 then stops for a separate Human confirmation before Stage 3.
- `RISK_TIER`: To be assessed from actual contract/security/privacy responsibility; no implementation tier or risk acceptance is implied by Exploration.
- `MODEL_ROUTING_RATIONALE`: Repository and authority reasoning across a public cross-stack contract boundary. `gpt-6-luna` / `high` is the configured low-cost model/lowest selected effort previously sufficient for Slice-2 Explorer assignments; approval is not required. Escalate only for demonstrated capability insufficiency.

## Canonical kickoff inputs

- `{HARSCODE_WORKSPACE_ROOT}`: `../harscode-workspace`
- `{TASK_PATH}`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005`
- `{CODEBASE_CONTEXT}`: `Kencleng — Go backend + Next.js frontend`
- `{TASK}`: Reconcile the public Campaign Detail donation-action contract and its owning Campaign acceptance criteria with the approved Slice 2 requirement for a guest Donation entry from an eligible Public Campaign Detail. Explore current relevant authority and repository state before planning. Product meaning remains governed by `docs/product/mvp-delivery-slices.md` §5 and `docs/product/mvp-scope.md`; Campaign behavior/spec by `docs/spec/4-campaign/`; shared contract by authored `api/openapi/campaign.yaml` and referenced components; use `WU-S2-004/runs/EXP-S2-004-001/evidence/stage-3-solutioning.md` as the exact source of the assigned coordination route. Preserve backend authority over actual submission eligibility. Do not infer missing product/API decisions.
- Ticket / Area: `WU-S2-005` / Campaign public donation-action contract; identify the exact affected spec/API/generated surfaces during Stage 1.
- Canonical prompt: `../harscode-workspace/workflow/1-exploration-kickoff-prompt.md`.
- Orchestrated overlay/context: `../harscode-workspace/workflow/orchestrated-run-overlay.md`, `workflow/context-management.md`, `workflow/AGENTS.md`, `orchestration/run-contract.md`, `orchestration/AGENTS.md`, and current Pilot #2 guidance under `orchestration/pilot-2-candidate/`.
- Explorer guidance: `../harscode-workspace/workflow/1-exploration/guidelines.md` and `sniffing-checklist.md`.

## Execution envelope

- `PREAUTHORIZED`: Read routed current authority and relevant source; perform Stage 1 now; after its Human gate, perform Exploration Stages 2–3; write only this Run's Exploration evidence/handoff under its `RUN_PATH`.
- `ORCHESTRATOR_DECISION`: Re-scope only if evidence shows the Campaign/API reconciliation needs a distinct dependency or materially different bounded outcome; report first.
- `HUMAN_REQUIRED`: Stage 1 and Stage 2 gates; Campaign/API owner decisions; any Product/MVP change; residual-risk acceptance; implementation, test execution, or protected-path writes.

## Human-assisted dispatch

Start a fresh Explorer Session at the Kencleng repository root using `gpt-6-luna` / `high`. Start from the canonical Exploration kickoff prompt and provide the normal variables above plus this Invocation. Do not add a task-specific solution-steering prompt. Begin Stage 1 only and stop for Human confirmation before Stage 2.
