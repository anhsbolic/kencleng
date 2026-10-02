# Run Invocation — `TP-S2-003-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator after re-grounding on the completed Exploration handoff and current Harscode guidance. Refreshed sebelum dispatch setelah membaca EXP-S2-005-001 dan merekam keputusan Campaign/API owner; Run/Participant identity tetap sama. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `TP-S2-003-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-001`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: None; backend Donation Techplan synthesis
- `PARTICIPANT_ID`: `P-S2-003-TP-001-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — Exploration is complete; this is a new phase and requires a new Run, Participant, and fresh Session/context.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796`; re-read relevant live sources at dispatch, including the current working tree.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`; ordinary applicable guidance remains current-effective at dispatch.
- `MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by `.harscode-spaces/.local-config.yaml` (`approval_required: false`).
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: First Techplan synthesis for WU-S2-003 from completed `EXP-S2-003-001`; write a new Draft / In Review Techplan in this Run. Do not continue Exploration or start Build.
- `ORCHESTRATOR_ROUTE`: WU-S2-003 owns the minimum backend capability needed for D1 within its existing scope; Campaign retains ownership of Campaign eligibility/close/funding coordination. No separate Campaign lifecycle WU is created. This is a delivery-boundary decision, not selection of a transaction/locking mechanism or authorization to edit a protected path. Techplan verifies the live seam and records any remaining implementation/Tier-0 gate.
- `MODEL_ROUTING_RATIONALE`: Configured Planner route is suitable for backend authority, interface, architecture, and evidence synthesis. Escalate only for demonstrated capability insufficiency; missing authority or evidence is not solved by model escalation.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/manifest.md`, parent `outcome.md`, `work-graph.md`, `control-surface.md`, and `events.md` — current delivery authority and route.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/EXP-S2-003-001/evidence/stage-2-gap-analysis.md` and `evidence/stage-3-solutioning.md` — complete durable Exploration evidence; read both as the full prior-phase record.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` — current-effective Approved Slice 2 Donation contract direction and delivery/testing obligations.
- `docs/product/mvp-delivery-slices.md` §5; `docs/spec/5-donation/` accepted invariants, threat model, tasks and features; `docs/spec/4-campaign/` only where D1/threshold behavior applies; `api/openapi/donation.yaml` and referenced `api/openapi/common.yaml`; `docs/project/kencleng-monetary-data-standard.md`.
- Root `AGENTS.md`, `backend/AGENTS.md`, `docs/project/kencleng-backend-tech-stack.md`, and relevant live backend code/migrations referenced by Exploration. Follow root Tier-0 fences; no protected write is authorized.
- Current Harscode guidance at dispatch: `workflow/2-1-techplan-synthesis-prompt.md`, `workflow/orchestrated-run-overlay.md`, `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`, `workflow/context-management.md`, and applicable workflow/orchestration AGENTS. Read matching best-practice guidance only where the canonical prompt's trigger applies.

- Coordination context: `WU-S2-005/runs/EXP-S2-005-001/evidence/stage-2-gap-analysis.md` dan `stage-3-solutioning.md`, WU-S2-005 manifest, serta event owner decision di parent `events.md`. Anhar memilih availability-only pada Campaign GET; ini adalah keputusan direction, sementara authored Campaign spec/API masih Slice 1 sampai rekonsiliasi dan acceptance selesai. Work Graph memuat dependency scoped untuk Campaign GET donation-action producer. Perhitungkan producer gap dan scope/owner implementasinya dalam Techplan; jangan menganggap GET snapshot mengotorisasi POST atau merevisi kontrak Campaign melalui backend Build. Other Donation/D1 planning tetap runnable.

## Task and completion condition

Synthesize an execution-grade Techplan for the backend-owned guest Donation capability in WU-S2-003, grounded in the current-effective authority, accepted contract, completed Exploration evidence, the Orchestrator's recorded D1 delivery boundary, and live backend conventions. Preserve settled Product/domain/API decisions, identify material unresolved items with their owners and gates, and make implementation and verification scope executable without inventing authority or bypassing protected paths.

Write only this Run's `techplan.md` and any phase-owned `launch-record.md` required by the current orchestrated workflow. Use the canonical Techplan template and self-check every rule-to-testing mapping, sensitive Test Focus pointer, open item, and material decision. Keep the Techplan `Draft / In Review`; do not generate `report-techplan.md` before any applicable planning review has converged and the plan is ready for its Human approval gate.

Do not change Product/MVP/spec/API authority, source code, tests, prior Exploration evidence, or orchestration projections; do not select a Tier-0 mechanism, accept Security/PII residual risk, authorize implementation, or claim `BACKEND_VERIFIED`. Surface any material owner/Human decision that blocks a safe implementation direction as an Active Open Item and stop at the Techplan handoff.

## Execution envelope

- `PREAUTHORIZED`: read current routed authority/code and write only this Run's Techplan and phase provenance/handoff artifacts.
- `ORCHESTRATOR_DECISION`: after synthesis, determine whether independent review is warranted/required by current guidance and route it; report any scoped blocker with owner and next action.
- `HUMAN_REQUIRED`: approve/revise the Techplan; make Product/domain/API/Security/PII decisions or accept residual risk; authorize any protected Tier-0 write; authorize later implementation gates.

## Human-assisted dispatch

Start a fresh Planner Participant Session at the Kencleng repository root. Use configured model `gpt-6-luna` with reasoning effort `high`. Begin with `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`, this Invocation, and the current orchestrated-run overlay. The canonical prompt plus the named authorities, recorded Orchestrator route, and completed Exploration artifacts define the assignment; do not add a separate solution-steering prompt. Write the Techplan to this `RUN_PATH/techplan.md`, include the required phase handoff, and stop before review/report approval or Build.

Kickoff: `Jalankan Techplan Synthesis Run TP-S2-003-001 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-001/invocation.md dan canonical Techplan synthesis workflow Harscode saat ini. Ground pada authority yang ditunjuk Invocation, current-effective TP-S2-002-015, serta seluruh bukti Exploration Stage 2 dan Stage 3. Tulis Techplan Draft/In Review dan handoff di RUN_PATH sesuai prompt/overlay. Jangan mengubah authority, source, tests, atau projection; jangan menyetujui Techplan, membuat report sebelum gate review yang berlaku, mengotorisasi Build, atau mengklaim BACKEND_VERIFIED. Berhenti setelah phase handoff.`
