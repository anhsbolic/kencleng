# Run Invocation — `RV-S2-003-003`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator. Human-facing prose: Bahasa Indonesia; preserve canonical Harscode terms and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `RV-S2-003-003`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-003`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent Techplan fidelity review
- `PARTICIPANT_ID`: `P-S2-003-RV-003-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; `.harscode-spaces/participant-profiles/profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer context from Planner TP-S2-003-006 and distinct from prior Reviewer RV-S2-003-001 / superseded package RV-S2-003-002.
- `TARGET_REVISION`: Kencleng HEAD `bb69cd002b3f1a1056837affcd77bb2b001007b0` plus current working tree. The exact Draft target is pinned separately below; re-ground live sources when dispatched.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; observed provenance, ordinary guidance remains current-effective.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml` (SHA-256 `ddffdeb4ce8edec181ca36bf32bedf19c513a76e359bc2642985eda6b32f400d`).
- `MODEL_ROUTING_RATIONALE`: This plan crosses Campaign and Donation contracts, monetary admission, D1 concurrency/exact-once settlement, credentials/PII, authorization and protected boundaries. An independent `KC-REVIEWER` at high effort is suitable for the fidelity checks; identify evidence/authority gaps from current sources rather than resolving them by model choice.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Independent Techplan Review of the exact completed WU-S2-003 Draft. Reviewer independently determines whether the canonical Complex gate applies and performs only the review warranted by that gate.

## Canonical route and current inputs

- Canonical phase prompt: `../harscode-workspace/workflow/2-2-techplan-review-prompt.md` plus `../harscode-workspace/workflow/orchestrated-run-overlay.md`.
- Canonical current Techplan Review prompt calls for current `../harscode-workspace/workflow/2-techplan/template.md`, `rules.md`, and `guardrails.md`; every durable WU-S2-003 Exploration artifact when the Complex gate applies; diagram guidance only if the plan contains a diagram; and target-repository sources for technical-fact spot checks.
- Assigned target: `WU-S2-003/runs/TP-S2-003-006/techplan.md`, SHA-256 `d98b3d37e3c00962075428c3e14f87d5ee59dfba46b6ae36ce35f78dcebbce7a`; terminal handoff `WU-S2-003/runs/TP-S2-003-006/handoff.md`, SHA-256 `e62e4df0cfc9095d06ffe6e987afdca8f6ee369a2f7a16d7d1306ddaa2653bdd`. Target remains Draft / In Review, unapproved. Treat this exact content identity as the assignment; if target materially changes before review, stop and reconcile its identity before claiming review.
- Complete WU-S2-003 Exploration corpus: all durable evidence under `WU-S2-003/runs/EXP-S2-003-001/evidence/` and the Exploration terminal handoff. Read the full corpus if Step-0 Complex gate applies; do not infer it only from the Planner's Draft.
- Current authorities and evidence: root `AGENTS.md`; `backend/AGENTS.md`; relevant sections of `docs/kencleng-agentic-workflow.md`; `docs/product/mvp-scope.md`; `docs/product/mvp-delivery-slices.md` Slice 2; monetary standard; accepted Campaign and Donation invariants/features; split authored `api/openapi/campaign.yaml`, `donation.yaml`, and referenced `common.yaml`; `api/README.md`; backend architecture and live code; WU-S2-003 manifest; parent Outcome/Work Graph/Control Surface/Events; WU-S2-006 manifest and source/counterpart evidence.
- Exact WU-S2-006 source acceptance snapshot is in `WU-S2-006/runs/RV-S2-006-006/invocation.md`; Anhar accepted only those exact seven revisions. Orchestrator rechecked all seven current hashes against that snapshot while reconciling TP-S2-003-006; they match. BLD-S2-006-006 records generated API/types and frontend fixture/test counterparts. These do not prove backend runtime behavior.
- The latest Human route assigns WU003 the minimum persisted representative-membership and Organization eligibility data/integration needed for Campaign draft writes; initial MVP records may use the Product/MVP-permitted seeded/operator-assisted setup, with full self-service and representative-management UI excluded. The route does not itself identify who may establish/update the Organization `verified` / `has_overdue_report` facts or the governing controls. Independently assess whether the Draft preserves that distinction and correctly scopes its remaining gate; do not invent that authority or weaken accepted Owner/Staff and fresh-state checks. Organization invariants/API remain Draft/historical evidence, not accepted MVP authority.

## Review assignment

Independently assess the exact assigned Techplan against current authority and durable evidence. Follow the canonical prompt's gate and checks: rule fidelity and Testing Checklist coverage; decision fidelity; diagram validity only if present; Open Items lifecycle; 2–3 non-obvious technical-fact/guardrail spot-checks against live sources; and Test Focus Pointer completeness against Exploration evidence when Complex review applies. Findings must include location, defect, source evidence and materiality. Do not introduce alternative product/API semantics, re-litigate settled choices, or direct the result toward approval or rejection.

The current Draft incorporates the Human-routed minimum persisted source/integration, alongside the accepted Campaign draft create/PATCH behavior. Assess the technical/source claims, data boundary, migration/backfill and verification assignments, rule/checklist traceability, and whether the unresolved authority/governance Open Item properly gates only the affected data-update and Campaign handlers while preserving independent Donation work and its separate prerequisites. The review does not decide that Human/owner question, approve the plan, or authorize implementation.

## Execution envelope

- `PREAUTHORIZED`: Read the assigned Draft, complete applicable Exploration evidence, and routed current authorities; write only this Run's `review-findings.md`.
- `ORCHESTRATOR_DECISION`: Reconcile Review findings and route a fresh Planner resolution/re-review/report sequence as applicable.
- `HUMAN_REQUIRED`: Whole-Techplan approval; any new Product/domain/API/security/data-authority decision; protected writes; migration/index application; risk acceptance; delivery/runtime milestone.
- No writes to the Techplan, Product, design, specs, API, code, tests, prior artifacts, manifests, tracker, registry, or other Work Unit. Do not generate `report-techplan.md`, approve, build, run tests/validators/generators/runtime/browser checks, or claim a delivery milestone. No Participant is dispatched by this preparation.
- Review output: one canonical `RUN_PATH/review-findings.md` with provenance and exactly one structured `## Phase handoff` per the active orchestrated overlay. Add `patch-plan.md` only if the canonical Review result requires one.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`.
Use a fresh independent Reviewer / `KC-REVIEWER` Session, `gpt-6-luna` / `high`.

Kickoff: `Jalankan independent Techplan Review Run RV-S2-003-003 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-003/invocation.md dan canonical ../harscode-workspace/workflow/2-2-techplan-review-prompt.md dengan orchestrated-run overlay. Tulis review-findings.md pada RUN_PATH dan berhenti setelah phase handoff.`

Laporkan completion atau scoped blocker ke Orchestrator. Tidak ada resolusi Planner, Human report/approval, Build, atau downstream dispatch otomatis.
