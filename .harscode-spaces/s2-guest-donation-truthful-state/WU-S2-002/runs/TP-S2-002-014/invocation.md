# Run Invocation — `TP-S2-002-014`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-09-30 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `TP-S2-002-014`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-014`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Material post-approval Techplan reconciliation for Human decisions O1, O8, and O11
- `PARTICIPANT_ID`: `P-S2-002-TP-014-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; Profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Planner Run after material Human decisions post-dated the current-effective Approved TP-011; do not reuse prior Participant/Session context.
- `TARGET_REVISION`: Kencleng checkout `550bafab081445b3ddabaa3b3911e57b983e77e5` plus current durable working-tree decisions/projections; re-read those live sources at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; use current guidance at dispatch unless a material change affects this assignment.
- `MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by `.harscode-spaces/.local-config.yaml` (`approval_required: false`).
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Required fresh Planner amendment to the current-effective Approved Techplan before further Task 02 Build; material decisions must enter the authoritative spine before execution.
- `PRIOR_ARTIFACTS`: See the exact current-effective inputs listed below; TP-011 is the predecessor spine, not a sufficient execution plan for Task 02 after the later decisions.
- `SESSION`: Fresh Planner Session at Human-Assisted dispatch; Session ID not exposed yet.
- `MODEL_ROUTING_RATIONALE`: Configured non-escalation Planner model has repository-work and reasoning capability; high effort suits reconciling material cross-domain money/PII/API decisions across the existing spine. Escalate only if a concrete capability insufficiency appears, not to compensate for missing context or authority.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` — current-effective Human-approved `Approved` Techplan spine. Preserve it as history; do not overwrite it. TP-011 predates the O1/O8/O11 decisions below.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-012/report-techplan.md` and `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-006/review-findings.md` — predecessor report/review provenance; not approval for the new revision.
- `.harscode-spaces/s2-guest-donation-truthful-state/events.md` — explicit Human decisions: O8 compatibility confirmation, O11 route-B supersession, project-wide O1 owner attribution, and O1 representation direction approval. Reconstruct their exact bounds; do not ask the Human to repeat them.
- `docs/project/kencleng-monetary-data-standard.md` — current canonical project-wide O1 direction and deliberate non-decisions.
- `.harscode-spaces/authority-map.md` — current project-wide Shared Currency Standard owner/scope and separate Slice-2 owner attributions.
- `docs/product/mvp-scope.md` and `docs/product/mvp-delivery-slices.md` — Product/MVP authority, including whole-Rupiah IDR input and terminal-notice obligation.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-009/o2-delivery-proposal.md` — bounded Delivery evidence that informs O11; it does not select numeric bound, architecture, timeout-as-failed, or residual-risk acceptance.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-005/amount-contract-brief.md` — O1 current Slice-2/repository evidence and historical alternatives; Human's later approved direction supersedes its recommendation-only status.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001/` — existing post-approval task manifest/files. They are derived snapshots of TP-011 and contain stale O1/O11 gates; do not treat them as current spine or execute Task 02 from them before a post-approval reconciliation.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/manifest.md`, `work-graph.md`, `outcome.md`, `events.md`, and `docs/project/kencleng-development-tracker.md` — current orchestration projections/history.
- Kencleng `AGENTS.md`, `docs/spec/README.md`, `api/README.md`, and `docs/project/communication-profile.md` — target authority routing, authored API source workflow, and communication profile.
- Current Harscode `workflow/2-1-techplan-synthesis-prompt.md`, `workflow/orchestrated-run-overlay.md`, `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`; use `workflow/2-2-techplan-review-prompt.md` for the independent Review gate and `workflow/2-3-techplan-decomposition-prompt.md` only after the revised spine is Approved and its task snapshots need reconciliation.

## Task and completion condition

Use the canonical Techplan synthesis entrypoint as a fresh material amendment. Write a new `RUN_PATH/techplan.md` execution-grade Draft/In Review spine plus this Run's `launch-record.md`, with this Run's own provenance in the Techplan template fields using only dispatch/session values actually known or exposed. TP-011 remains the current-effective Approved predecessor until a later Human approves this new spine; preserve TP-011 and its report as history and do not edit/overwrite them. Do not edit task snapshots, source spec/API, implementation, or orchestration projections in this Run.

Reconcile only the material post-TP-011 authority/decision delta and dependent planning text:

1. Carry the Human-approved project-wide O1 representation direction into the spine: major-unit decimal string plus explicit currency code on API/wire; exact-decimal calculation and persistence end-to-end without float/float64; active currencies remain under Product Authority; Slice-2 input remains whole-Rupiah IDR. Keep Campaign `NUMERIC(19,2)` as implementation precedent only. Preserve explicit non-decisions for additional currencies, numeric range, per-currency fraction rules, universal database precision/scale, and migration detail. Concrete values remain for later evidence-backed decisions that preserve every valid exact value.
2. Reconcile O11 as a resolved policy direction: supersede route B; verified email cannot be deleted before the terminal-notice obligation is fulfilled; Delivery must produce a bounded, recoverable terminalization policy. Preserve the explicit boundary: no numeric bound, architecture, timeout-as-failed meaning, or residual-risk acceptance. Keep any genuinely unresolved Delivery/Security evidence under the correct Active/deferred Open Item; do not mislabel the Human policy question as still open.
3. Reconcile O8 to its bounded current scope: the Human/API owner confirmed the historical submit/status operations were never externally distributed. Mark the compatibility gate clear for the planned replacement in this repository/scope without claiming a universal absence of external consumers.
4. Carry the full TP-011 spine, D1, O2–O5 decisions/evidence boundaries, O7/O9, Slice-2 scope, all applicable rules/risks/Test Focus and verification obligations forward without unrelated scope expansion. Move resolved items to the resolved-history part of Open Items with their actual resolution and consequence. Keep downstream runtime/security proof distinct from contract-time decisions.
5. Reconcile the `CONTRACT_READY` boundary so it does not require evidence that only Build/Testing can generate, while retaining all material authority and owner-acceptance prerequisites. Do not mark the milestone earned.
6. State that Task 01 Human acceptance/review remains an independent parallel work item. Its existing draft artifacts may be reviewed now; any requested revision for the newer O1/O11 direction remains a separate Task 01 acceptance/reconciliation path, not a reason to serialize its review behind this Planner Run.
7. Declare materiality. This decision delta materially changes the money representation/interface and verified-email lifecycle policy, so independent Techplan Review is required under the current route/Complex criteria. Do not generate the Human report during synthesis/review churn; after review/resolution converges, the Planner report Run generates the current report in full before Human approval.

Do not silently edit Product/MVP or Approved TP-011, choose any explicitly deferred O1/O11 parameter, accept residual risk, author specs/OpenAPI, create a child-task-only contract, claim `CONTRACT_READY`, or start Build. After the new spine is approved, use a fresh post-approval task-snapshot reconciliation Run only for child task files materially affected by that spine; preserve the accepted Task 01 → Task 02 topology, dependency, and existing manifest when they have not materially changed. Do not re-open Human acceptance of the split merely because affected task contents were reconciled. If the approved spine materially changes the topology/dependency, route that changed split for Human review before Build. Task files may not replace the spine. Do not run this reconciliation against the Draft.

## Execution envelope

- `PREAUTHORIZED`: read the assigned current sources and write only this Run's new `techplan.md` and `launch-record.md`.
- `ORCHESTRATOR_DECISION`: prepare and route the applicable independent Review, report-generation, Human approval, and post-approval task-snapshot reconciliation in current projections.
- `HUMAN_REQUIRED`: approve/revise the new Techplan; any new Product/security/domain decision or residual-risk acceptance; protected Tier-0 authorization; any claim of `CONTRACT_READY` or delivery milestone.

## Human-assisted dispatch

Use a fresh Planner Session in the Kencleng repository root with the configured model/effort. Start from `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md` and this Invocation; apply `workflow/orchestrated-run-overlay.md` path semantics (`RUN_PATH/techplan.md`). Kickoff: `Jalankan Planner reconciliation Run TP-S2-002-014 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-014/invocation.md dan canonical Techplan synthesis workflow Harscode saat ini. Reconstruct dari current-effective spine TP-011 dan keputusan Human durable yang ditunjuk; masukkan O1/O11 serta O8 bounded resolution ke Techplan spine, jaga semua non-decisions dan evidence boundary, dan pertahankan Task 01 Human acceptance sebagai parallel work. Tulis Techplan Draft/In Review baru dan launch record Run ini saja. Nyatakan materiality dan rekomendasi review/decomposition. Jangan buat report sebelum review/resolution converged, jangan edit task files/spec/API/code/projection, jangan klaim CONTRACT_READY, dan jangan mulai Build. Berhenti setelah phase handoff.`

After dispatch, report only a material issue or phase completion. Orchestrator reconciles the route from this Run's artifacts; Build remains unavailable until the new spine is Human-approved and affected task snapshots are reconciled after approval.
