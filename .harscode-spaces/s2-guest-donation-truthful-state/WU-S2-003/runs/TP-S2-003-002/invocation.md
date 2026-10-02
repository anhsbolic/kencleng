# Run Invocation — `TP-S2-003-002`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator setelah Human menanyakan apakah blocker harus diputuskan sebelum independent Review. Human-facing prose: Bahasa Indonesia; preserve canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `TP-S2-003-002`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-002`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Draft resolution / decision-ready technical proposal
- `PARTICIPANT_ID`: `P-S2-003-TP-002-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — completed TP-001 occurrence dire-entry melalui Run, Participant, dan Session baru.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree; verify relevant live source saat dispatch.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`; ordinary applicable guidance remains current-effective.
- `MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned local configuration.
- `MODEL_ROUTING_RATIONALE`: Configured Planner route untuk material architecture/interface proposal berdasarkan repo. Prior constraint/missing proposal adalah routing/context gap, bukan bukti capability insufficiency.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Canonical Techplan synthesis sebagai new Draft-resolution occurrence. Complete concrete proposal dan bounded owner question(s) sebelum independent review yang saat ini diparkir; tidak ada prior Review verdict yang diklaim.
- `MEANINGFUL_DELTA`: Klarifikasi planning envelope dan explicit delivery of complete technical/contract proposals untuk Active Items TP-001; tidak mengulang Exploration atau hanya menyalin daftar blocker.

## Current-effective inputs / PRIOR_ARTIFACTS

- `WU-S2-003/runs/TP-S2-003-001/techplan.md` dan Invocation — Draft predecessor; preserve sebagai history.
- Seluruh durable Exploration evidence `WU-S2-003/runs/EXP-S2-003-001/evidence/`; fresh Session membaca setiap prior-phase file sesuai canonical prompt/overlay.
- WU-S2-003 manifest, parent events/work-graph/control-surface/outcome, Authority Map, current Product/MVP, root AGENTS dan applicable scoped sources.

- `WU-S2-003/runs/TP-S2-003-001/launch-record.md` — completed planning handoff.
- Approved `WU-S2-002/runs/TP-S2-002-015/techplan.md`, agreed Donation specs, authored Donation/common OpenAPI, monetary standard, root/scoped backend AGENTS, backend architecture dan referenced live code.
- WU-S2-005 manifest/owner availability-only decision, Draft `WU-S2-005/runs/TP-S2-005-001/techplan.md`, Work Graph scoped Campaign GET producer dependency. Draft bukan accepted contract.
- Canonical `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`, Techplan template/rules/guardrails, orchestrated-run overlay/context-management, applicable workflow/orchestration AGENTS; matching best-practices ketika triggered.

Paths `WU-S2-*` relatif terhadap `.harscode-spaces/s2-guest-donation-truthful-state/`; source paths lain relatif terhadap repository root. Overlay mengarahkan output ke RUN_PATH dan reads ke named prior artifacts.

## Task and completion condition

Lengkapi Draft backend pada scope WU-S2-003 dari material Active Open Items predecessor. Delta utama adalah clarification envelope: protected source writes tetap dilarang, tetapi Planner boleh mengusulkan interface, transaction/locking design, dan file-level ownership sebagai Draft untuk Human review/approval. Ini tidak mengotorisasi implementation atau memutuskan owner policy.

D1 harus menjadi proposal yang decision-ready: explicit Campaign/Donation ownership dan interface; transaction/ordering/replay/rollback boundary; proposed implementation files dan Tier-0 classification, termasuk file baru untuk balance transaction/locking; meaningful alternatives, recommendation, trade-offs, serta exact owner/protected-work question. Pertahankan D1 policy yang settled dan bounded WU route; tidak membawa full Slice 3 closure/result behavior atau merutekan implementation di sekitar fence.

Material O2/O3/O4/O5 design dan persistence parameters yang diperlukan juga membutuhkan source-grounded proposal yang cukup konkret untuk owner decisions, bukan generic daftar “owner menentukan detail”. Reuse settled decisions (termasuk window verification guest email yang sudah diputuskan), bedakan architecture/control proposal dari genuinely missing Product/API/security decision, external operational inputs, dan empirical proof yang memang milik Build/Testing. Jangan menutup unknown dengan invented parameter atau mengklaim runtime evidence. Dependency Campaign GET producer pada accepted WU-S2-005 contract tetap scoped.

Owner yang telah dipetakan boleh diajak mengambil keputusan langsung di Participant Session setelah proposal cukup jelas; catat jawaban explicit dengan scope/provenance dalam plan/handoff. Protected implementation authorization harus tetap terpisah dan concrete; proposal atau plan approval tidak otomatis mengotorisasi semua protected writes.

Tulis successor `RUN_PATH/techplan.md` sebagai Draft/In Review dengan material proposals berstatus Proposed/Pending owner jika belum diputuskan; preserve clean predecessor requirements, decisions, risks, verification dan Open Item history. Root fencing tetap binding untuk implementation writes. Planner tidak mengarang owner decision, self-approve, menerima residual risk atau mengubah Approved TP-S2-002-015.

Untuk setiap material owner item yang siap dibahas, berikan problem, bounded context, meaningful alternatives, recommendation beserta consequence, dan exact decision ask. Jangan mengirim Human untuk memilih file/angka/contract yang belum diusulkan atau mengulang keputusan settled. Bila owner menjawab di Session, record durably; bila belum, Active item harus sudah menyatakan proposal/options/evidence dan exact remaining decision. Any new authority conflict harus dirutekan pada owning authority.

Self-check canonical traceability dan technical facts. Handoff harus membedakan proposal complete, owner decisions made/pending, protected-write gates, genuinely external missing inputs, dan deferred runtime evidence; recommend applicable review setelah semantic planning converges. Jangan membuat report selama churn/pre-review; report mengikuti canonical Human gate setelah applicable review/resolution converges.

## Execution envelope

- `PREAUTHORIZED`: read routed live source; write only this Run's Draft Techplan dan phase-owned provenance/handoff artifacts.
- `ORCHESTRATOR_DECISION`: facilitate remaining bounded owner questions dari durable proposal; reconcile independent Review target/routing setelah successor semantics converges.
- `HUMAN_REQUIRED`: material owner decisions, plan approval, protected source writes, residual-risk acceptance, final contract acceptance dan applicable implementation gates.

Tidak mengubah Product/spec/API/code/tests/generated artifacts/projections, prior Run evidence atau protected code; tidak menjalankan migration/test/runtime work, mengklaim empirical proof, atau memulai Build.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Start fresh Planner Session dengan `gpt-6-luna` / `high`. Gunakan canonical Techplan synthesis prompt dan Invocation ini. Owner questions yang sudah decision-ready dapat dibahas langsung dengan Human di Session sesuai Harscode; record explicit answers. Berhenti setelah Draft resolution dan phase handoff; laporkan proposal/decisions atau material blocker ke Orchestrator.

Kickoff: `Jalankan canonical Techplan Draft resolution Run TP-S2-003-002 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-002/invocation.md dan ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md dengan orchestrated-run overlay. Rekonstruksi dari prior artifacts yang ditunjuk, tulis artifact milik RUN_PATH, dan berhenti setelah phase handoff.`

