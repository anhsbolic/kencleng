# Run Invocation — `EXP-S2-004-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-004`
- `RUN_ID`: `EXP-S2-004-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/EXP-S2-004-001`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004`
- `ROLE`: Explorer
- `SPECIALIZATION`: None; frontend flow Delivery Exploration
- `PARTICIPANT_ID`: `P-S2-004-EXP-001-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-EXPLORER`; current Profile SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `PRIOR_ARTIFACTS`: WU-S2-004 manifest; current Parent Outcome and Work Graph; Approved TP-S2-002-015; TPD-S2-002-002 Task 01 snapshot; the five accepted Donation specs; authored Donation/common OpenAPI contract; current Product Slice 2 and canonical page-map/UI-UX entrypoint. Resolve live source changes at dispatch; do not load the historical Run archive wholesale.
- `SESSION_TRANSITION`: `FRESH` — new Work Unit/Run and canonical Exploration phase.
- `SESSION_TRANSITION_REASON`: Context hygiene and distinct frontend delivery scope after completion of WU-S2-002.
- `TARGET_REVISION`: Kencleng HEAD `055aa1283ff5f1ab9c08c1422981a9cbccf98e6c` plus the live existing working-tree changes; re-read current sources at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; resolve current Exploration guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; local Human-owned config marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`.
- `PHASE_ROUTE`: Canonical Exploration kickoff, Stage 1 only initially. Stage 1 is a hard stop for Human confirmation before Stage 2; after Stage 2, a separate Human confirmation is required before Stage 3.
- `RISK_TIER`: To be assessed by Exploration/Techplan from current source and approved Techplan; Donation task baseline is Tier 1. Frontend risk is assessed by actual responsibility and is not inherited from backend.
- `MODEL_ROUTING_RATIONALE`: `gpt-6-luna` / `high` is the configured capable minimum evidenced by recent Explorer Runs for repository authority/design/code-anchor analysis and bounded interaction-flow mapping. Escalate only for demonstrated capability insufficiency, not missing context/authority.

## Canonical kickoff inputs

- `{HARSCODE_WORKSPACE_ROOT}`: `../harscode-workspace`
- `{TASK_PATH}`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004`
- `{CODEBASE_CONTEXT}`: `Kencleng — Go backend + Next.js frontend`
- `{TASK}`: Deliver the guest Donation frontend flow for the approved Slice 2 outcome. Explore the current relevant authority and repository state before planning. Governing sources: `docs/product/mvp-delivery-slices.md` §5; `docs/product/mvp-scope.md` guest donation boundary; `docs/ui-ux/README.md` and `docs/ui-ux/page-map.md`; agreed `docs/spec/5-donation/tasks.md` Task 01/02 and feature specs; accepted authored API `api/openapi/donation.yaml` plus referenced `api/openapi/common.yaml`; frontend architecture and scoped frontend instructions as routed. Preserve current Product/Design/API/spec authority; mock work must follow the accepted contract.
- Ticket / Area: `WU-S2-004` / guest Donation flow; identify exact pages/components during Stage 1.
- Canonical prompt: `../harscode-workspace/workflow/1-exploration-kickoff-prompt.md`.
- Orchestrated overlay and context: `../harscode-workspace/workflow/orchestrated-run-overlay.md`, `../harscode-workspace/workflow/context-management.md`, `../harscode-workspace/workflow/AGENTS.md`, `../harscode-workspace/orchestration/run-contract.md`, `../harscode-workspace/orchestration/AGENTS.md`.
- Explorer guidance: `../harscode-workspace/workflow/1-exploration/guidelines.md` and `sniffing-checklist.md`.

## Execution envelope

- `PREAUTHORIZED`: Read current authority/routed sources, perform Exploration stages only after their Human gates, inspect relevant frontend code read-only, and write Exploration evidence/handoff only under this Run's `RUN_PATH`.
- `ORCHESTRATOR_DECISION`: Re-scope the page/flow boundary if current Product/Design/dependency evidence materially requires it.
- `HUMAN_REQUIRED`: Stage 2/Stage 3 confirmations; material Product/Design/API/Security/PII decisions; brand-defining visual assets; later Human rendered acceptance; implementation or tests (later phases only).

## Human-assisted dispatch

Start a fresh Explorer Session at the Kencleng repository root using `gpt-6-luna` / `high`. Start from the canonical Exploration kickoff prompt and supply its normal variables above plus this Invocation. Do not add a task-specific solution-steering prompt. Begin Stage 1 and stop for Human confirmation before Stage 2.
