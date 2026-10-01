# Run Invocation — RV-S2-002-009

Status: READY_FOR_HUMAN_DISPATCH

Prepared: 2026-09-30 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- WORK_UNIT_ID: WU-S2-002
- RUN_ID: RV-S2-002-009
- RUN_PATH: .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-009
- WORK_UNIT_PATH: .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002
- ROLE: Reviewer
- SPECIALIZATION: Independent Complex review of material post-approval Techplan reconciliation TP-S2-002-014
- PARTICIPANT_ID: P-S2-002-RV-009-1 (ephemeral for this Run)
- PARTICIPANT_PROFILE_ID: KC-REVIEWER; pin current Profile content revision at dispatch.
- SESSION_TRANSITION: FRESH — independent Reviewer Run after completed Planner Run TP-S2-002-014.
- TARGET_REVISION: Kencleng checkout 650e73c5d646c29c0ddf1931618f02685d15f7b7 plus TP-014's current Run artifacts and current working-tree orchestration update; re-read live authorities at dispatch.
- WORKFLOW_REVISION: Harscode checkout 33b03a3f62cc3aacba6534b8a011465613c64b09; re-read canonical guidance at dispatch.
- HARSCODE_WORKSPACE_ROOT: ../harscode-workspace
- RUNTIME_HARNESS: codex-cli
- SELECTED_MODEL: gpt-6-luna; local Human-owned configuration says approval is not required.
- REASONING_EFFORT: high.
- COMMUNICATION_PROFILE_PATH: docs/project/communication-profile.md
- PHASE_ROUTE: Required fresh independent Complex Review after material changes to monetary interface and verified-email lifecycle semantics; do not proceed to report generation until review/resolution converges.
- MODEL_ROUTING_RATIONALE: Configured non-escalation model/profile is suitable for repository-grounded fidelity review. Escalate only for demonstrated capability insufficiency; missing authority/evidence is routed as a finding, not solved by model escalation.

## Current-effective inputs

- .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-014/techplan.md and launch-record.md — the current Draft / In Review spine and actual synthesis provenance; review this new artifact, not TP-011 alone.
- .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md and report-techplan.md — current-effective Approved predecessor, for full spine and authority-delta comparison.
- .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-012/report-techplan.md and RV-S2-002-006/review-findings.md — predecessor approval/review provenance.
- All durable Exploration evidence in .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/ — broad reread is required if Complex gate applies.
- Focused current evidence: OIR-S2-002-005 amount-contract-brief.md; OIR-S2-002-006 campaign-donation-ordering-brief.md and Stage 2 evidence; TP-S2-002-009 o2-delivery-proposal.md; OIR-S2-002-004 design-review-brief.md; OIR-S2-002-002/003 O3–O5 and O2/O3 evidence/handoffs; current Events and Authority Map.
- Product/MVP authority: docs/product/README.md, mvp-scope.md, mvp-delivery-slices.md, docs/project/kencleng-monetary-data-standard.md.
- Target authority and technical spot-check sources as needed: root AGENTS.md, backend/AGENTS.md, docs/spec/README.md, Donation/Campaign specs, api/README.md, authored api/openapi/donation.yaml and api/openapi/common.yaml, relevant live code only when a claim needs verification.
- Orchestration state: WU-S2-002 manifest, work-graph.md, events.md, control-surface.md, docs/project/kencleng-development-tracker.md.
- Canonical current Harscode guidance: workflow/2-2-techplan-review-prompt.md; workflow/2-techplan/template.md, rules.md, guardrails.md; workflow/orchestrated-run-overlay.md; applicable AGENTS files.
- Profile and runtime routing source: .harscode-spaces/participant-profiles/registry.md, profiles.md, and .harscode-spaces/.local-config.yaml.

## Task and completion condition

Independently review the whole TP-014 Techplan under the canonical Review prompt. Resolve the current template section names at runtime. Apply the Complex gate: the plan crosses domain/API contracts and includes high-stakes money/payment/PII boundaries. Check:

1. Rule fidelity to Product/MVP and durable evidence; verify every R1–R14 has Testing Checklist coverage and that ownership/rationale preserves phase boundaries.
2. Decision fidelity, especially approved O1 direction and its explicit parameter deferrals; bounded O8 scope; O11 supersession of route B without numeric bound/architecture/timeout-as-failed/risk acceptance; preservation of D1, O2–O5, O7, and O9.
3. Contract boundary: the plan's contract-time requirements versus empirical O2–O5 Build/Testing evidence; whether the stated CONTRACT_READY gate is usable and does not require evidence that can only be produced after implementation.
4. Open Items lifecycle: O1 concrete parameters remain explicitly deferred; O8 and O11 are retained as Resolved history; Task 01 Human acceptance remains parallel; O2–O5 owner/evidence work is correctly routed.
5. Target technical facts/guardrails: spot-check at least 2–3 non-obvious claims against current authority/spec/live source; verify money representation and exact-decimal rule, authored/derived OpenAPI discipline, compatibility claim bounds, and protected Tier-0 boundaries.
6. Test Focus Pointer completeness and exact durable Exploration evidence anchors for surviving money/concurrency/security/PII risks.

Classify findings per the canonical prompt as MATERIAL / BLOCKING or MECHANICAL / NON-BLOCKING. Each finding must state location, defect, source evidence, and materiality. Do not rewrite the Techplan or make authority decisions.

Write only this Run's review-findings.md and launch-record.md. Do not generate the Planner report, approve the plan, accept residual risk, edit source specs/OpenAPI/code/tests, claim CONTRACT_READY, or start Build. Stop after the independent Review handoff.

## Execution envelope

- PREAUTHORIZED: read the assigned current sources and write only this Run's review artifacts.
- ORCHESTRATOR_DECISION: expand scope or change the workflow route.
- HUMAN_REQUIRED: authority changes, Techplan approval, residual-risk acceptance, protected writes, or milestone acceptance.

## Human-assisted dispatch

Use a fresh independent Reviewer Session in the Kencleng repository root with the configured model/effort. Start from ../harscode-workspace/workflow/2-2-techplan-review-prompt.md and this Invocation. Kickoff: Jalankan independent Complex Techplan Review Run RV-S2-002-009 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-009/invocation.md dan canonical Techplan Review prompt Harscode saat ini. Review seluruh TP-S2-002-014 terhadap current Product/MVP, durable Exploration evidence, D1, O1/O8/O11 decisions, Test Focus anchors, verification ownership, dan current contract-versus-runtime boundary. Ikuti Complex gate dan format findings canonical. Jangan edit Techplan/authority/spec/API/code/test, jangan buat report, jangan approve, jangan klaim CONTRACT_READY atau mulai Build. Tulis hanya review-findings.md dan launch-record.md Run ini, lalu berhenti.

After dispatch, Orchestrator reconciles the route from Review artifacts. If any finding requires correction, prepare a fresh Planner resolution Run and re-review when applicable. If review converges cleanly, prepare a fresh Planner report-generation Run before the Human approval gate.
