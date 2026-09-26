# Run Invocation — `TP-S2-002-006`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared by: Orchestration Operator  
Prepared: 2026-09-26  
Human-facing prose: Bahasa Indonesia; preserve canonical Harscode terms/enums and code/API/schema identifiers.

## Invocation identity

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `TP-S2-002-006`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-006`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `TARGET_REVISION`: `10457b17059e2da3d97a4f76e4c3fae127227d73`
- `WORKFLOW_REVISION`: `b122a75d494250d04eb93e71f4c391e82c847842`
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `ROLE`: Planner
- `SPECIALIZATION`: Material Techplan amendment after Human-approved Product/MVP scope update
- `PARTICIPANT`: Codex Planner
- `SESSION`: Fresh Planner Session, independent from prior synthesis, resolution, and review sessions.
- `SESSION_TRANSITION`: `FRESH`
- `SESSION_TRANSITION_REASON`: `MATERIAL_PLAN_REVISION` — the Approved Techplan's product scope, domain rules, interface contract, risks, and verification implications must be reconciled to Human-approved Product/MVP amendments and OIR decisions.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by `.harscode-spaces/.local-config.yaml` (`approval_required: false`).
- `MODEL_ROUTING_RATIONALE`: This Run reconciles multiple resolved Human decisions across Donation/Campaign, scope, API, and security/PII boundaries. `gpt-6-luna` / `high` is the configured non-escalation choice sufficient for repository planning; escalate only on concrete capability insufficiency.
- `PHASE_ROUTE`: `REQUIRED` — current-effective Techplan `TP-S2-002-003` is Approved but predates the Human-approved Product/MVP and OIR decisions. A material amendment is necessary before dependent spec/API reconciliation can proceed without contradicting current authority.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-003/techplan.md` — current-effective Approved Techplan; amend through a new Run, do not overwrite its historical copy.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-004/report-techplan.md` — previous approval digest; becomes stale if the Techplan is materially amended and must not be hand-patched.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-001/resolution-brief.md` — current per-item decision/continuation brief, including explicit Human priority register for O1–O6/O9.
- `docs/product/mvp-scope.md` and `docs/product/mvp-delivery-slices.md` — current Product/MVP authority, amended and Human-approved 2026-09-26; these supersede conflicting pre-amendment wording for Slice 2.
- `docs/product/README.md`, `AGENTS.md`, `docs/kencleng-agentic-workflow.md`, and only the relevant current Product, Design, Donation/Campaign spec, API, and Security/PII authorities.
- `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`, `workflow/2-techplan/template.md`, `workflow/2-techplan/rules.md`, `workflow/2-techplan/guardrails.md`, `workflow/orchestrated-run-overlay.md`, `workflow/2-2-techplan-review-prompt.md`, and current Pilot #2 candidate guidance.

## Task

Create a material Techplan amendment in `RUN_PATH/techplan.md` that reconciles the current Product/MVP authority and OIR Human decisions. Re-open only the sources needed to verify current exact wording, consequences, and technical/authority boundaries. Preserve the prior Techplan as history and explicitly identify the new artifact as a revision of `TP-S2-002-003`.

At minimum:

1. Carry forward current Human-approved Slice 2 product direction from the updated Product/MVP sources: integer IDR amount input (minimum Rp5.000, Rp1 increments), exact-decimal money representation, display of QRIS/GoPay/ShopeePay/bank transfer with only QRIS active as a labeled sandbox simulation, simulator-owned truthful states/recovery, optional guest name and opt-in verified terminal-status email, 24-hour status-only guest link with generic failure behavior, threshold-crossing/full-settlement semantics, and same-key ambiguous retry/idempotency policy.
2. Trace each decision to its current authoritative source and record the 2026-09-26 Product/MVP amendment. Move resolved items into the Techplan's Resolved Open Items with the exact decision and consequence; retain unresolved items as Active with named owner/evidence needs. Do not drop O1–O9 history.
3. Keep exact simulator timing, amount representation/derived precision where not set, guest-email verification/retention/retry windows and controls, URL/token exposure mitigations and residual-risk decision, response parity and anti-enumeration details, and Design review open for their actual owners. Human product direction does not accept security/privacy residual risk.
4. Reconcile Decision Log, Requirements, Rules & Validation, risks, Interface Contract, Architecture/Plan, implementation anchors, Files Changed/NOT Changed, Testing Checklist, and Test Focus Pointers to the updated source truth. Preserve atomic success/funding coupling R3, Tier-0 file fences, and API source/generated-artifact rules. Cross-domain Campaign/Donation threshold behavior must not import Slice 3 beyond the approved threshold rule.
5. Mark the amended Techplan `Draft / In Review`; do not treat prior Human approval as approval of this material revision. Declare whether the change materially alters product/domain/interface/verification semantics (expected: yes) and recommend/route an independent Techplan review in a fresh Reviewer Session before report regeneration and the next Human approval gate.

Do not edit `docs/product/*`, historical Techplan `TP-S2-002-003`, prior report `TP-S2-002-004/report-techplan.md`, specs, OpenAPI, code, tests, or orchestration projections. Do not generate `report-techplan.md` in this Run; the old report remains historical and is stale for approval of the amended Techplan. Do not claim `CONTRACT_READY`, start Build, accept residual security/privacy risk, or choose technical controls on behalf of Security/PII, Design, API, Campaign, or Donation owners. Stop after `RUN_PATH/techplan.md` and `RUN_PATH/launch-record.md` phase handoff.

## Human-Assisted dispatch

- Role: Planner (`Codex Planner`)
- Model / effort: `gpt-6-luna` / `high`
- Working directory: Kencleng repository root (`/home/anhar-solehudin/kencleng-workspace/kencleng`)
- Session: fresh Planner Session, independent of all prior Planner/Explorer/Reviewer sessions.
- Durable invocation: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-006/invocation.md`
- Kickoff to paste: `Jalankan Planner amendment Run TP-S2-002-006 dengan mengikuti invocation durable .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-006/invocation.md, canonical Techplan synthesis prompt, Techplan template/rules/guardrails, orchestrated-run overlay, dan current Pilot #2 guidance dari ../harscode-workspace. Gunakan OIR brief dan dokumen Product/MVP yang Human-approved pada 2026-09-26 sebagai current authority; amend Techplan Approved TP-S2-002-003 ke RUN_PATH/techplan.md tanpa mengubah dokumen sumber atau artifact historis. Pertahankan Open Items yang masih memerlukan owner teknis/security/design/API, tandai material amendment dan route fresh independent review. Jangan buat report, jangan mulai Build, jangan klaim CONTRACT_READY. Tulis hanya techplan.md dan launch-record.md Run ini, lalu berhenti.`
- After dispatch, Human reports only a material problem or completion; Orchestrator follows fire-and-forget posture and then reconciles durable artifacts.
