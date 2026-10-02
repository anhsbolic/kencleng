# Run Invocation — `RV-S2-003-001`

Status: `READY_FOR_HUMAN_DISPATCH` — retarget/unpark ke successor TP-S2-003-002 setelah completion dan owner planning decisions tercatat. Berdasarkan last-known durable state belum di-dispatch; belum ada Review artifact/verdict. Active controls/authorization/dependencies tetap perlu dinilai menurut canonical execution-grade/materiality checks, bukan dianggap resolved oleh routing ini.

Prepared: 2026-10-01 by Orchestration Operator; target direkonsiliasi sebelum dispatch setelah Human melaporkan TP-S2-003-002 selesai. Identity Run/Participant dipertahankan karena belum ada dispatch reliance atau completed Review yang tercatat. Human-facing prose: Bahasa Indonesia; preserve canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `RV-S2-003-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-001`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent Techplan review
- `PARTICIPANT_ID`: `P-S2-003-RV-001-1`
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

- Review target: `WU-S2-003/runs/TP-S2-003-002/techplan.md`, `launch-record.md` (termasuk phase handoff), dan Invocation-nya. TP-S2-003-001 adalah predecessor history, bukan current review target.
- Recorded Human planning decisions: current Draft Decision Log/Open Items dan launch-record, direkonsiliasi pada parent completion event. Keputusan valid di Participant Session tidak perlu di-vote ulang; tidak sama dengan whole-Techplan approval, protected-write authorization, source authority promotion, runtime evidence atau residual-risk acceptance.
- Full Exploration record: `WU-S2-003/runs/EXP-S2-003-001/evidence/stage-2-gap-analysis.md` dan `evidence/stage-3-solutioning.md`; enumerate directory evidence dan baca setiap durable Exploration file sesuai canonical Complex review.
- Current WU-S2-003 manifest, parent events/work-graph/control-surface/outcome, Authority Map, root AGENTS dan scoped instructions saat diperlukan, current Product/MVP.
- Completed Planner provenance/handoff: `WU-S2-003/runs/TP-S2-003-002/launch-record.md`. Item O3/O4/O5 controls/operational input/evidence tetap Active; pertahankan perbedaan keputusan material, file-specific authorization, scoped dependencies dan empirical proof menurut canonical checks tanpa menentukan verdict terlebih dahulu.
- Approved `WU-S2-002/runs/TP-S2-002-015/techplan.md`; agreed Donation spec/invariant/threat-model/features; authored `api/openapi/donation.yaml` and referenced common components; monetary standard; applicable backend authority/live code.
- Campaign coordination context: WU-S2-005 manifest/events, current Draft `WU-S2-005/runs/TP-S2-005-002/techplan.md` dan handoff, completed `WU-S2-005/runs/RV-S2-005-001/review-findings.md`; Planner report-only TP-S2-005-003 last-known queued. Availability-only/schema proposal decisions tercatat; final authored contract belum accepted, scoped producer dependency tetap berlaku.
- Canonical `../harscode-workspace/workflow/2-2-techplan-review-prompt.md`, Techplan template/rules/guardrails, orchestrated-run overlay/context, dan applicable workflow/orchestration AGENTS. Diagram guidance hanya jika plan memuat diagram.

Paths `WU-S2-*` relatif terhadap `.harscode-spaces/s2-guest-donation-truthful-state/`; source paths lain relatif terhadap repository root. Overlay menggantikan ordinal `2-techplan/techplan.md` dan `1-exploration/logs/` dengan target/prior artifacts di atas.

## Clarification of planning authority

Orchestrator mengoreksi restriction Invocation TP-S2-003-001 yang menyatakan “do not select a Tier-0 mechanism”: root AGENTS §3 melindungi implementation writes, termasuk file baru yang menjalankan transaction/locking balance update. Planner boleh mengusulkan desain, interface, mekanisme, dan file-level scope sebagai Draft untuk review/owner approval tanpa mengubah protected implementation. Usulan itu tidak menjadi approved owner decision atau protected-write authorization. Completed TP-001 evidence tetap dipertahankan. Reviewer menilai plan berdasarkan authority dan canonical execution-grade requirements; Reviewer tidak mengambil alih penyusunan desain atau approval. Ini klarifikasi envelope, bukan hasil review yang telah ditentukan.

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

Kickoff: `Jalankan independent Techplan Review Run RV-S2-003-001 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-001/invocation.md dan canonical ../harscode-workspace/workflow/2-2-techplan-review-prompt.md dengan orchestrated-run overlay. Tulis artifacts Review pada RUN_PATH dan berhenti setelah phase handoff.`

