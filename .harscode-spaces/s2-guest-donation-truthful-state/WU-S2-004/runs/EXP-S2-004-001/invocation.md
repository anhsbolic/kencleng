# Run Invocation — `EXP-S2-004-001`

Status: Exploration Stage 3 completed; phase handoff awaits Human Campaign/API coordination route decision.

Prepared: 2026-10-01 by Orchestration Operator. Re-grounded before dispatch against Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8` and clean Kencleng `HEAD=7fd8b473b239b20bda3990ab29c51440d321a796`; assignment meaning is unchanged. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-004`
- `RUN_ID`: `EXP-S2-004-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/EXP-S2-004-001`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004`
- `ROLE`: Explorer
- `SPECIALIZATION`: None; frontend flow Delivery Exploration
- `PARTICIPANT_ID`: `P-S2-004-EXP-001-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-EXPLORER`; current Profile SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `PRIOR_ARTIFACTS`: WU-S2-004 manifest; current Parent Outcome and Work Graph; Approved `WU-S2-002/runs/TP-S2-002-015/techplan.md`; TPD-S2-002-002 manifest and Task 01/02 snapshots; accepted `docs/spec/5-donation/invariants.md`, `threat-model.md`, `tasks.md`, `features/01-submit-donation-settlement.md`, and `features/02-donation-status-check.md`; authored `api/openapi/donation.yaml` and referenced `api/openapi/common.yaml`; `docs/product/mvp-delivery-slices.md` §5; `docs/ui-ux/README.md` and `page-map.md`. Resolve live source changes at dispatch; do not load the historical Run archive wholesale.
- `SESSION_TRANSITION`: Continue the same Run and Participant at Stage 3; prefer the existing active Participant Session, replacing only the Session if it is unavailable.
- `SESSION_TRANSITION_REASON`: Stage 2's Human gate is cleared; canonical Exploration Stage 3 continues the same execution occurrence from durable Stage-2 evidence.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796`; current working tree includes durable Stage-2 evidence and orchestration state reconciliation, with no reported production-source changes. Re-read relevant source for Stage 3.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1` HEAD `95ecf37ba8ae449a5b3b278c27331aca87360bc8`; current-effective guidance provenance, not semantic pinning. Canonical Exploration prompt is unchanged; orchestration resume/dispatch guidance was refreshed.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; local Human-owned config marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`.
- `PHASE_ROUTE`: Continue this same canonical Exploration Run at Stage 3. Stage 2 evidence is the current prior-phase input; Anhar authorized Stage 3 on 2026-10-01. Do not repeat Stages 1–2 unless a material contradiction requires reopening them.
- `RISK_TIER`: To be assessed by Exploration/Techplan from current source and approved Techplan; Donation task baseline is Tier 1. Frontend risk is assessed by actual responsibility and is not inherited from backend.
- `MODEL_ROUTING_RATIONALE`: Capability need is repository reasoning plus authority/design/code-anchor analysis across guest tracking credentials, user interaction, and privacy-sensitive states. `gpt-6-luna` is the least-cost configured model with repository-work/reasoning capability; `high` is the lowest effort selected as sufficient for the complete Exploration, not only its Stage 1 checkpoint. Approval is not required. Escalate only for demonstrated capability insufficiency, not missing context/authority.

## Canonical kickoff inputs

- `{HARSCODE_WORKSPACE_ROOT}`: `../harscode-workspace`
- `{TASK_PATH}`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004`
- `{CODEBASE_CONTEXT}`: `Kencleng — Go backend + Next.js frontend`
- `{TASK}`: Deliver the guest Donation frontend flow for the approved Slice 2 outcome. Explore the current relevant authority and repository state before planning. Governing sources: `docs/product/mvp-delivery-slices.md` §5; `docs/product/mvp-scope.md` guest donation boundary; `docs/ui-ux/README.md` and `docs/ui-ux/page-map.md`; agreed `docs/spec/5-donation/tasks.md` Task 01/02 and feature specs; accepted authored API `api/openapi/donation.yaml` plus referenced `api/openapi/common.yaml`; frontend architecture and scoped frontend instructions as routed. Preserve current Product/Design/API/spec authority; mock work must follow the accepted contract.
- Ticket / Area: `WU-S2-004` / guest Donation flow; identify exact pages/components during Stage 1.
- Canonical prompt: `../harscode-workspace/workflow/1-exploration-kickoff-prompt.md`.
- Orchestrated overlay and context: `../harscode-workspace/workflow/orchestrated-run-overlay.md`, `../harscode-workspace/workflow/context-management.md`, `../harscode-workspace/workflow/AGENTS.md`, `../harscode-workspace/orchestration/run-contract.md`, `../harscode-workspace/orchestration/AGENTS.md`, and current Pilot #2 resume/dispatch routing in `../harscode-workspace/orchestration/pilot-2-candidate/project-orchestration-bootstrap.md`, `orchestrator-operating-model.md`, `participant-execution-continuation.md`, and `model-routing-and-escalation.md`.
- Explorer guidance: `../harscode-workspace/workflow/1-exploration/guidelines.md` and `sniffing-checklist.md`.

## Execution envelope

- `PREAUTHORIZED`: Read current authority/routed sources, perform Exploration stages only after their Human gates, inspect relevant frontend code read-only, and write Exploration evidence/handoff only under this Run's `RUN_PATH`.
- `ORCHESTRATOR_DECISION`: Re-scope the page/flow boundary if current Product/Design/dependency evidence materially requires it.
- `HUMAN_REQUIRED`: material Product/Design/API/Security/PII decisions; brand-defining visual assets; later Human rendered acceptance; implementation or tests (later phases only).

## Human-assisted dispatch

Continue `EXP-S2-004-001` with the same Explorer Participant `P-S2-004-EXP-001-1` and pinned `KC-EXPLORER` Profile revision, using `gpt-6-luna` / `high`. Prefer the same active Participant Session that produced the Stage-2 evidence. If that Session is no longer available, use a replacement Session for this same Run and Participant; reconstruct from this Invocation and `evidence/stage-2-gap-analysis.md`, without creating a new Run or Participant. At the Kencleng repository root, start from the canonical Exploration kickoff prompt and follow Stage 3 only. Compare bounded solution paths and trade-offs for this frontend scope; route the Campaign contract coordination and downstream email/credential boundaries to their owners without inventing authority or changing cross-stack files. Do not implement, edit production code/spec/API, or run tests in this Exploration. Stage-3 completion must write durable solutioning evidence and the canonical phase handoff under this Run's `RUN_PATH`.
