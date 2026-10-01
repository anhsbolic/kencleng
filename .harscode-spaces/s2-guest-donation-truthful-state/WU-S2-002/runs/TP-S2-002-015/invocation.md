# Run Invocation — `TP-S2-002-015`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `TP-S2-002-015`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Resolve independent Review findings in the material O1/O8/O11 Techplan reconciliation
- `PARTICIPANT_ID`: `P-S2-002-TP-015-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Planner occurrence after independent Reviewer Run RV-S2-002-009; do not reuse prior Participant/Session context.
- `TARGET_REVISION`: Kencleng checkout `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus current durable working-tree decisions/projections; re-read those live sources at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; use current guidance at dispatch unless a material change affects this assignment.
- `MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by `.harscode-spaces/.local-config.yaml` (`approval_required: false`).
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Required fresh Planner resolution of the two material/blocking decision-fidelity findings before independent re-review; no report generation until review/resolution converges.
- `MODEL_ROUTING_RATIONALE`: Configured non-escalation Planner model/profile is suitable for repository-grounded fidelity correction. Escalate only for demonstrated capability insufficiency; missing authority/evidence is not solved by model escalation.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-009/review-findings.md` and `launch-record.md` — completed independent Complex Review; two material/blocking findings define this resolution Run.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-014/techplan.md` and `launch-record.md` — current Draft / In Review spine to revise into this Run's new artifact; preserve the prior Run unchanged.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` — current-effective Human-approved predecessor; preserve as history and keep effective until a later Human approves a successor.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-002/resolution-brief.md`, `evidence/stage-2-o4-status-credential.md`, and `evidence/stage-2-o5-public-failure-parity.md` — durable Human-directed O4/O5 decisions and their still-open downstream controls/evidence.
- `.harscode-spaces/s2-guest-donation-truthful-state/events.md`, WU manifest, Work Graph, Control Surface, outcome, and project tracker — current orchestration facts/projections; distinguish decisions from unresolved implementation controls and residual-risk acceptance.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-014/invocation.md` — bounded material reconciliation scope and explicit non-decisions; use with RV-009 to retain all clean portions of TP-014.
- Current Product/MVP and monetary authority: `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md`, `docs/project/kencleng-monetary-data-standard.md`, `.harscode-spaces/authority-map.md`, root `AGENTS.md`, `api/README.md`, and applicable spec/API sources named by TP-014.
- Current Harscode guidance at dispatch: `workflow/2-1-techplan-synthesis-prompt.md`, `workflow/orchestrated-run-overlay.md`, `workflow/2-techplan/template.md`, `workflow/2-techplan/rules.md`, and `workflow/2-techplan/guardrails.md`; use deeper guidance only where the canonical prompt calls for it.

## Task and completion condition

Run the canonical Techplan synthesis workflow as a fresh resolution occurrence. Write only this Run's `techplan.md` and `launch-record.md`; keep TP-014, TP-011, and all prior evidence unchanged. The new spine remains Draft / In Review and TP-011 remains current-effective Approved.

Resolve these Review findings against the durable source decisions without inventing new decisions:

1. **RV-009 MATERIAL / BLOCKING O4 finding:** Carry forward the Human-selected fragment status URL, frontend handoff and URL cleanup, and one-way HMAC verifier. State them as settled direction at the appropriate interface/security boundary. Keep carrier/header expression in authored API reconciliation if not yet selected; keep browser history/referrer/log/cache exposure, key handling/comparison, token lifecycle enforcement, abuse controls, empirical proof, and any Security/PII residual-risk acceptance open with their correct owners. Do not present runtime proof or risk acceptance as already established.
2. **RV-009 MATERIAL / BLOCKING O5 finding:** Carry forward the Human-selected uniform public `404` for absent Donation, missing/wrong credential, and expired credential, with the same public failure result and body/header/cache behavior. Reconcile exact Problem Details coordinates and the `Cache-Control: private, no-store` contract detail in the authored API source. Keep empirical response/timing parity and abuse/rate-control evidence downstream. Do not reopen the selected status code or claim empirical parity.

Retain the complete TP-014 spine and its clean review conclusions: approved O1 direction and explicit parameter deferrals; bounded O8 compatibility clearance; O11 route-B supersession and non-decisions; D1; O2/O3 lifecycle and evidence boundaries; O7/O9; R1–R14 and Testing Checklist coverage; Test Focus anchors; Task 01 Human acceptance as a parallel work item; and the corrected `CONTRACT_READY` evidence boundary. Recheck affected wording and cross-references throughout the document, not only the lines cited by Review. Declare the change/materiality and recommend another fresh independent Complex Review because both changes affect material credential/API failure semantics.

Do not change Project/MVP authority, accept Security/PII residual risk, select unresolved O4/O5 controls, author specs/OpenAPI, edit task snapshots, mutate source/tests, claim `CONTRACT_READY`, start Build, or generate the Human report during resolution/re-review. After the required re-review converges, a separate fresh Planner Run generates the full report before Human approval/revision. Task snapshot reconciliation remains after approval and only for affected child files.

## Execution envelope

- `PREAUTHORIZED`: read assigned current sources and write only this Run's new `techplan.md` and `launch-record.md`.
- `ORCHESTRATOR_DECISION`: prepare and route the required independent re-review; after convergence, prepare report-generation and Human approval steps under current projections.
- `HUMAN_REQUIRED`: approve/revise the successor Techplan; make any new Product/security/domain decision or residual-risk acceptance; authorize protected Tier-0 work; claim `CONTRACT_READY` or a delivery milestone.

## Human-assisted dispatch

Use a fresh Planner Session in the Kencleng repository root with configured model/effort. Start from `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md` and this Invocation; apply `workflow/orchestrated-run-overlay.md` path semantics (`RUN_PATH/techplan.md`). Kickoff:

`Jalankan Planner resolution Run TP-S2-002-015 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/invocation.md dan canonical Techplan synthesis workflow Harscode saat ini. Reconstruct dari TP-014, current-effective TP-011, durable OIR-002 decisions, dan temuan RV-009 yang ditunjuk. Pertahankan bagian TP-014 yang lulus review, koreksi dua temuan material O4/O5 beserta semua bagian terdampak, dan jaga batas keputusan settled versus contract detail, runtime evidence, serta residual-risk acceptance. Tulis Techplan Draft/In Review baru dan launch record Run ini saja. Nyatakan materiality dan rekomendasi independent re-review. Jangan buat report sebelum re-review konvergen, jangan edit task files/spec/API/code/tests/projection, jangan klaim CONTRACT_READY, dan jangan mulai Build. Berhenti setelah phase handoff.`

After dispatch, reconcile the route from this Run's artifacts. The independent re-review must inspect TP-015 against RV-009 and live authorities before report generation can be queued.
