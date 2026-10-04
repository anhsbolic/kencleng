# Run Invocation — `RV-S2-003-006`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-04 after completion of Planner Run TP-S2-003-009. This is a fresh independent review of the current WU-S2-003 candidate. It neither approves the plan nor closes the separate Donation/D1 migration-design gate.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `RV-S2-003-006`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-006`
- `ARTIFACT_TARGET`: `none` — write only Run-owned `review-findings-1.md` and `launch-record.md` if required by the review execution record; do not edit the candidate, Approved predecessor, source, or Exploration evidence.
- `WORK_UNIT_PATH` / compatibility `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent Complex Techplan Review of the post-approval WU-S2-003 candidate successor
- `PARTICIPANT_ID`: `P-S2-003-RV-006-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer context after TP-S2-003-009; do not reuse the Planner's Participant Session or treat Orchestrator synthesis as Review evidence.
- `TARGET_REVISION`: Kencleng HEAD `4e71d3697a479d92b2cad5a79ef5d710f4e8077c` plus current working tree. Review target candidate SHA-256 `3aa5e5d362bbfac5cb65b231159d793552d57cc61236fbb126bcbadee3416a60`; verify before substantive review and stop/re-ground if candidate content changes.
- `WORKFLOW_REVISION`: current-effective Harscode workflow. Preparation identity: review prompt SHA-256 `e81b88ae275e459df7f5b8172dd091cafee929b5f25095d31730d00482c98464`; template `aa2bb5acdf7b2d4d79975fc5bbcd92b47094223361485382bce0520e3fca0ccd`; rules `ab6939e9b4d104bd0db2669945ee5a32e560ebe006356363b3ed8da969d56c2e`; guardrails `74160046a553d2b53a84320c1397b5fdee44fa9bb1e3c3fd9fb532c038e2a6a4`; overlay `f806673aa5e7d1d82dae9f433ec54c3562e15ce778b709a141c57d3d22dfaf87`; workflow router `746ca72aeb31b0dd33592c5fd9a3965f8ecdba6c4a4c5508d2f5c3b509cabd93`. Re-read current-effective guidance at dispatch.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`, SHA-256 `ddffdeb4ce8edec181ca36bf32bedf19c513a76e359bc2642985eda6b32f400d`.
- `MODEL_ROUTING_RATIONALE`: This is a substantial fidelity and technical-fact review of a 20-rule cross-boundary plan touching Donation/payment behavior, PII, concurrency, API contracts, migrations, and protected-write limits. `gpt-6-luna` / `high` is supported by the local registry and consistent with the prior independent WU003 Techplan review. Sol is approval-gated and is not required for this review.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Use canonical `workflow/2-2-techplan-review-prompt.md` with `workflow/orchestrated-run-overlay.md`. This is review of the exact post-Approval `techplan.candidate.md`, not the predecessor's historical Review target, not the migration-proposal review, and not approval of the candidate.

## Exact review target and current-effective inputs

- Sole target: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/techplan.candidate.md`, SHA-256 `3aa5e5d362bbfac5cb65b231159d793552d57cc61236fbb126bcbadee3416a60`. It is a post-Approval material successor to the current Approved predecessor, not a pre-Approval `techplan.md`.
- Current Approved predecessor: `runs/TP-S2-003-006/techplan.md`, SHA-256 `b71951c86b2f74f6c9f1a164bcf43447c58ee655874deaf09275e6ff3b2316f9`. Preserve its Approved state and historical identity; its earlier Review `RV-S2-003-003` targeted the pre-reconciliation Draft, not this candidate.
- Candidate synthesis receipt: `runs/TP-S2-003-009/handoff.md`, SHA-256 `357e70ca4c8c73ef350741571ae7a4843e274de72f7f720454a9b4118b3ebf37`; Planner Invocation SHA-256 `37e6594dce161c84fe78082700bb0c7efebb2e1c3704eda4eeb8612aa5ccf1f7`.
- Read every durable Exploration artifact under `runs/EXP-S2-003-001/evidence/`: `stage-2-gap-analysis.md` and `stage-3-solutioning.md`. Treat them as read-only historical evidence and follow anchors in the candidate.
- Reopen current target-repository authority for any claim being checked: root `AGENTS.md`; `backend/AGENTS.md`; `docs/kencleng-agentic-workflow.md`; Product README, overview, MVP scope and delivery slices; applicable Campaign/Donation invariants and feature source; accepted split API sources (`api/openapi/donation.yaml` and referenced `api/openapi/common.yaml`); relevant monetary standard; and live code/schema/migration paths only where needed to verify specific candidate anchors. Do not treat historical code or the prior plan as overriding current Product/MVP authority.
- WU-S2-007 acceptance handoff: `../WU-S2-007/handoff-to-WU-S2-003.md`, SHA-256 `cf468693dc384ccad5a4156c43d433a3ff725a6169869c3f171aeff0a27325eb`. It records exact accepted source/counterpart revisions and scoped internal-consumer posture.
- Prior migration-design Review: `runs/RV-S2-003-005/review-findings-1.md`, SHA-256 `fca27d88610017db317f679a68c97af7b49a616d3b7b69d6fa06b45fef359c7f`, verdict `Request changes`. Its blocking design findings F-02/F-03/F-04/F-06 remain open in that separate gate. Verify the candidate accurately preserves the dependency; do not count this whole-Techplan Review as resolving those migration findings.
- Canonical review authorities to re-read at dispatch: `../harscode-workspace/workflow/2-2-techplan-review-prompt.md`; `workflow/2-techplan/template.md`, `rules.md`, and `guardrails.md`; `workflow/orchestrated-run-overlay.md`; `workflow/AGENTS.md`. Read diagram guidance only if the exact candidate contains a diagram.

## Review task and boundaries

Apply the canonical independent Techplan Review to the captured candidate revision. First state whether the Complex gate applies and why. Resolve semantic section names to the current template numbering at execution time. Broadly re-ground on all durable Exploration evidence, then perform the canonical checks: rule fidelity and checklist coverage, decision fidelity, diagram validity if present, Open Item lifecycle, technical-fact/guardrail spot-checks against live target-repository sources, and completeness/accuracy of every Test Focus Pointer.

Classify only evidence-backed findings as `MATERIAL / BLOCKING` or `MECHANICAL / NON-BLOCKING`; include location, defect, source evidence, and materiality. Do not hunt for polish after the material checks are complete, rewrite the plan, reopen correctly recorded decisions, or steer toward a predetermined verdict.

This review does not resolve Human-owned Open Item 7, O3/O4/O5 controls, or the separate migration-design findings by assumption. Identify a material conflict or missing authority only if the candidate/source evidence warrants it. A clean Review does not approve the candidate: Human approval of the exact converged candidate remains a separate gate. Do not dispatch a resolver, migration-design Review, Build, Testing, or any other Run.

## Execution envelope and required artifact

- `PREAUTHORIZED`: read the pinned candidate, all cited durable Exploration evidence, current applicable authorities, and narrow live source anchors required by the canonical spot-check; write Run-owned independent findings.
- `HUMAN_REQUIRED`: candidate approval/promotion; any material Product/domain/API/security decision; changes to protected or stable planning artifacts; downstream Build authorization, migration application, or residual-risk acceptance.
- Write `RUN_PATH/review-findings-1.md` with provenance and exactly one structured `## Phase handoff` using the orchestrated fixed fields: Outcome, Result refs, Findings, Decision requests, Blockers, Open / unverified, Recommended continuation, Context refs. A `launch-record.md` may record execution identity if the active target-repo run convention requires it; it must not duplicate the findings or handoff.
- No tests, validators, generators, source/spec/API/Techplan edits, migration/database actions, or downstream Run dispatch are authorized by this Invocation.

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Reviewer / `KC-REVIEWER`, configured `gpt-6-luna` / `high`.

Kickoff: `Lakukan independent Techplan Review Run RV-S2-003-006 memakai canonical ../harscode-workspace/workflow/2-2-techplan-review-prompt.md dan orchestrated-run overlay, dengan inputs sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-006/invocation.md. Verifikasi hash target sebelum review; baca semua durable Exploration evidence dan current authority sesuai canonical prompt. Review candidate sebagai post-Approval successor secara independen, termasuk traceability seluruh Rules ke evidence dan Testing Checklist, fidelity decisions, lifecycle Open Items, spot-check technical anchors, serta kelengkapan Test Focus Pointer. Jangan edit Techplan/source/spec/API/tests, jangan menyelesaikan authority item dengan asumsi, jangan menutup migration-design gate RV-S2-003-005, dan jangan dispatch resolver/Build/Testing atau Run lain. Tulis review-findings-1.md dengan satu structured phase handoff, lalu berhenti.`

This Invocation prepares a Run only; Anhar dispatches the fresh Reviewer. No Participant has been dispatched by the Orchestrator.
