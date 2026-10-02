# Run Invocation — `RV-S2-005-001`

Status: `READY_FOR_HUMAN_DISPATCH` — target direkonsiliasi ke successor TP-S2-005-002 setelah completion dan explicit owner schema decision tercatat. Invocation ini belum di-dispatch berdasarkan last-known durable state; tidak ada Review verdict sebelumnya.

Prepared: 2026-10-01 by Orchestration Operator; retarget/unpark sebelum dispatch setelah Human melaporkan TP-S2-005-002 selesai. Run/Participant identity dipertahankan karena belum ada dispatch reliance atau Review artifact yang tercatat. Human-facing prose: Bahasa Indonesia; preserve canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-005`
- `RUN_ID`: `RV-S2-005-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/RV-S2-005-001`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent Techplan review
- `PARTICIPANT_ID`: `P-S2-005-RV-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer Participant/Session baru, terpisah dari Planner yang menyintesis plan.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree; verify source live saat dispatch.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`; ordinary applicable guidance remains current-effective.
- `MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned local configuration.
- `MODEL_ROUTING_RATIONALE`: Configured Reviewer route untuk independent source fidelity, cross-domain/public-contract, security, dan execution-grade checks. Escalate only for demonstrated capability insufficiency.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Canonical independent Techplan review; jalankan Step 0 Complex gate dan review bila warranted. Plan tetap Draft/In Review; tidak ada Human approval/report convergence yang diklaim.

## Current-effective inputs / PRIOR_ARTIFACTS

- Review target: `WU-S2-005/runs/TP-S2-005-002/techplan.md`, `handoff.md`, dan Invocation-nya. TP-S2-005-001 adalah predecessor history, bukan active review target.
- Explicit Campaign/API owner schema decision tercatat pada TP-S2-005-002 Decision Log/Open Items/handoff dan parent completion event: `available` tanpa reason; `unavailable` dengan satu-satunya reason `campaign_not_eligible`, hanya jika detail tetap public tetapi gagal submission eligibility pada GET. Non-public/closed tetap 404; dependency failure 503; reason lama dihapus tanpa compatibility-only enum. Keputusan proposal ini tidak sama dengan Techplan approval atau final authored acceptance.
- Full Exploration record: `WU-S2-005/runs/EXP-S2-005-001/evidence/stage-2-gap-analysis.md` dan `evidence/stage-3-solutioning.md`; enumerate directory evidence dan baca setiap durable Exploration file sesuai canonical Complex review.
- Current WU-S2-005 manifest, parent events/work-graph/control-surface/outcome, Authority Map, root AGENTS dan scoped instructions saat diperlukan, current Product/MVP.
- Campaign feature/invariants dan authored `api/openapi/campaign.yaml` + referenced common components; `api/README.md`; generated correspondence bila perlu fact spot-check.
- Accepted Donation submit/D1 sources dan Product Slice 2/3 boundary; owner availability-only event; EXP-S2-004-001 Stage-3 sebagai derivation context bila diperlukan.
- Canonical `../harscode-workspace/workflow/2-2-techplan-review-prompt.md`, Techplan template/rules/guardrails, orchestrated-run overlay/context, dan applicable workflow/orchestration AGENTS. Diagram guidance hanya jika plan memuat diagram.

Paths `WU-S2-*` relatif terhadap `.harscode-spaces/s2-guest-donation-truthful-state/`; source paths lain relatif terhadap repository root. Overlay menggantikan ordinal `2-techplan/techplan.md` dan `1-exploration/logs/` dengan target/prior artifacts di atas.

## Task and completion condition

Jalankan canonical independent Techplan review terhadap target di atas menggunakan current source authority dan seluruh durable Exploration evidence. Ikuti gate/checks/materiality pada prompt, termasuk execution-grade sufficiency dan lifecycle Active/Resolved items. Bedakan keputusan material yang belum memiliki proposal/authority, proposal yang menunggu owner gate, serta empirical proof yang memang milik Build/Testing. Tidak ada finding/verdict yang ditentukan oleh Invocation ini.

Tulis `RUN_PATH/review-findings.md` dan phase-owned provenance/handoff; `launch-record.md` bila diperlukan. Gunakan canonical output/handoff; stop setelah Review. Orchestrator akan merutekan Planner resolution dan Human decisions dari finding yang material.

## Execution envelope

- `PREAUTHORIZED`: read assigned plan/evidence dan routed live sources; write only Review artifacts dalam RUN_PATH.
- `ORCHESTRATOR_DECISION`: reconcile findings, route fresh Planner resolution bila diperlukan, lalu applicable re-review/report/Human approval.
- `HUMAN_REQUIRED`: owner decisions, plan approval, protected writes, residual-risk acceptance dan applicable authority gates.

Reviewer tidak mengubah Techplan/report, Product/spec/API/code/tests/projections atau prior artifacts; tidak membuat desain pengganti, mengotorisasi Build, atau menyatakan milestone. Review ini reasoning/fidelity phase; tidak menjalankan suite, migrations, validator/generation atau runtime/security checks untuk menggantikan downstream Testing.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Buka fresh Reviewer Session menggunakan `gpt-6-luna` / `high`. Gunakan canonical independent Techplan review prompt dan Invocation ini. Stop setelah Review findings dan phase handoff; laporkan completion atau material blocker ke Orchestrator.

Kickoff: `Jalankan independent Techplan Review Run RV-S2-005-001 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/RV-S2-005-001/invocation.md dan canonical ../harscode-workspace/workflow/2-2-techplan-review-prompt.md dengan orchestrated-run overlay. Tulis artifacts Review pada RUN_PATH dan berhenti setelah phase handoff.`

