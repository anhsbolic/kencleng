# Run Invocation — RV-S2-002-010

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `RV-S2-002-010`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-010`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent Complex re-review of TP-S2-002-015 resolving RV-S2-002-009 O4/O5 findings
- `PARTICIPANT_ID`: `P-S2-002-RV-010-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer occurrence after completed Planner resolution TP-S2-002-015; do not reuse prior Reviewer/Planner Participant or Session context.
- `TARGET_REVISION`: Kencleng checkout `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus TP-015's current Run artifacts and current durable working-tree orchestration update; re-read live authorities at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; re-read canonical guidance at dispatch.
- `MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by `.harscode-spaces/.local-config.yaml` (`approval_required: false`).
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Required independent Complex re-review after material changes to status-credential/API failure semantics; report generation remains gated on review/resolution convergence.
- `MODEL_ROUTING_RATIONALE`: Configured Reviewer model/profile is suitable for repository-grounded fidelity review. Escalate only for demonstrated capability insufficiency; missing authority/evidence becomes a finding, not a model-selection issue.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` and `launch-record.md` — current Draft / In Review resolution spine and its provenance; review this artifact as the plan under review.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-009/review-findings.md` and `launch-record.md` — two material/blocking O4/O5 fidelity findings this Run must independently verify as resolved.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-014/techplan.md` and `launch-record.md` — preceding Draft spine; use to confirm clean portions survived while the findings were corrected.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` and matching `report-techplan.md` — current-effective Human-approved predecessor and its report; preserve its full spine and compare the post-approval delta.
- All durable Exploration evidence in `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/` — broad independent reread required by the Complex gate.
- Focused current decision evidence: `OIR-S2-002-002/resolution-brief.md`, `evidence/stage-2-o4-status-credential.md`, `evidence/stage-2-o5-public-failure-parity.md`; O1 `OIR-S2-002-005/amount-contract-brief.md`; D1 `OIR-S2-002-006/campaign-donation-ordering-brief.md` and Stage 2 evidence; O2 delivery proposal; O7 design brief; O2/O3 evidence and handoffs; current Events and Authority Map.
- Current Product/MVP and technical authority: `docs/product/README.md`, `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md`, `docs/project/kencleng-monetary-data-standard.md`, root `AGENTS.md`, `backend/AGENTS.md`, `docs/spec/README.md`, Donation/Campaign specs, `api/README.md`, authored `api/openapi/donation.yaml` and referenced `common.yaml`; spot-check live code only where needed.
- Orchestration state: WU-S2-002 manifest, `work-graph.md`, `events.md`, `control-surface.md`, and `docs/project/kencleng-development-tracker.md`.
- Current Harscode guidance at dispatch: `workflow/2-2-techplan-review-prompt.md`, `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`, `workflow/orchestrated-run-overlay.md`, and applicable target-repo instructions. The review prompt is a draft workflow artifact; apply its current Complex gate and output shape as directed by this Invocation.
- Profile/runtime routing sources: `.harscode-spaces/participant-profiles/registry.md`, `profiles.md`, and `.harscode-spaces/.local-config.yaml`.

## Task and completion condition

Independently review the whole TP-015 Techplan under the current canonical Review prompt. This is a Complex review: it crosses Donation/Campaign and authored API contracts and covers high-stakes money, payment, credential security, and PII boundaries. Resolve current template section names from `workflow/2-techplan/template.md` at review time. Review for fidelity and correctness; do not rewrite the plan or make authority decisions.

1. Verify TP-015 resolves both RV-009 findings against durable OIR decisions: O4 fragment URL, frontend handoff and URL cleanup, one-way HMAC verifier and bounded 24-hour/status-only direction; O5 uniform public `404` and identical body/header/cache behavior including `Cache-Control: private, no-store`. Confirm unresolved carrier/header contract expression is bounded consistently with the selected fragment flow; exact Problem Details/cache details are assigned to authored API reconciliation; controls, empirical parity, abuse evidence, and residual-risk acceptance are not falsely claimed.
2. Review the entire plan, not only the amended passages. Check fidelity to current Product/MVP and durable evidence, especially O1 and its explicit parameter deferrals, bounded O8, O11 supersession and non-decisions, D1, O2/O3 lifecycle, O7/O9, and all applicable scope/authority boundaries.
3. Confirm each R1–R14 has Testing Checklist coverage with proper evidence ownership and rationale; `CONTRACT_READY` separates contract-time decisions/owner acceptance/source validation from evidence only Build/Testing can produce; no premature milestone or approval is claimed.
4. Check Open Items lifecycle, Task 01 Human acceptance as a parallel path, material decision history, and Test Focus Pointer completeness with exact durable Exploration anchors for surviving money/concurrency/security/PII risks.
5. Spot-check at least 2–3 non-obvious technical claims against current authority/spec/live source, including monetary representation, authored/derived OpenAPI discipline, and protected Tier-0 boundaries. Report any contradiction or missing material decision with location, evidence, and materiality. Do not hunt for polish after material checks are complete.

Classify findings using canonical categories: `MATERIAL / BLOCKING` or `MECHANICAL / NON-BLOCKING`. Write only this Run's `review-findings.md` and `launch-record.md`. Do not edit TP-015/TP-014/TP-011, authorities, task snapshots, specs/OpenAPI, code/tests, or orchestration projections; do not generate the Planner report, approve the plan, accept residual risk, claim `CONTRACT_READY`, or start Build. Stop after independent Review handoff.

## Execution envelope

- `PREAUTHORIZED`: read assigned current sources and write only this Run's review artifacts.
- `ORCHESTRATOR_DECISION`: route findings and derive the next phase from review evidence.
- `HUMAN_REQUIRED`: authority changes, successor Techplan approval, residual-risk acceptance, protected writes, or milestone acceptance.

## Human-assisted dispatch

Use a fresh independent Reviewer Session in the Kencleng repository root with configured model/effort. Start from `../harscode-workspace/workflow/2-2-techplan-review-prompt.md` and this Invocation; apply the current overlay to read TP-015 from `RUN_PATH/techplan.md` and to write only this Run's review artifacts. Kickoff:

`Jalankan independent Complex Techplan Review Run RV-S2-002-010 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-010/invocation.md dan canonical Techplan Review prompt Harscode saat ini. Review seluruh TP-S2-002-015 secara independen terhadap RV-009, durable Exploration evidence, current Product/MVP dan monetary authority, D1, O1/O8/O11/O4/O5 decisions, verification ownership, Test Focus anchors, contract-versus-runtime boundary, dan live technical sources yang relevan. Ikuti Complex gate dan format findings canonical; verifikasi dua temuan RV-009 tanpa menganggap proof/residual-risk acceptance selesai. Jangan edit Techplan/authority/spec/API/code/test, jangan buat report, jangan approve, jangan klaim CONTRACT_READY atau mulai Build. Tulis hanya review-findings.md dan launch-record.md Run ini, lalu berhenti.`

After dispatch, reconcile the route from the review artifact. If any material finding remains, prepare a fresh Planner resolution and another independent re-review as required. Only after review/resolution converges may a fresh Planner report-generation Run be prepared for the Human approval gate.
