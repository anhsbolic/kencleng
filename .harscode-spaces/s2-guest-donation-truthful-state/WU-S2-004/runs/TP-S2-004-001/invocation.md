# Run Invocation — `TP-S2-004-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 oleh Orchestration Operator setelah WU-S2-005 completion: accepted Campaign action contract, independent Review/Testing, dan verified Status propagation. Human-facing prose: Bahasa Indonesia; preserve canonical terms/enums dan technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-004`
- `RUN_ID`: `TP-S2-004-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TP-S2-004-001`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Frontend guest Donation flow Techplan synthesis
- `PARTICIPANT_ID`: `P-S2-004-TP-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`, observed committed target revision di bawah.
- `SESSION_TRANSITION`: `FRESH` — first Techplan after completed frontend Exploration; new Run/Participant/context, independent of Campaign reconciliation and backend Planner Sessions.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796` plus accepted current working tree.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`; ordinary guidance current-effective, bukan policy pin.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned local registry.
- `MODEL_ROUTING_RATIONALE`: First synthesis memakai established frontend/product/design directions dan accepted contracts; normal repository planning medium cukup sebagai initial routing. Material missing authority/context tidak diatasi dengan stronger model. Historical frontend GPT-5.6 defaults tidak tersedia pada current registry; gunakan smallest sufficient available capability, reconsider bila concrete insufficiency ditemukan.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Initial frontend Techplan synthesis; Draft checkpoint, bukan Build authorization.
- `RISK_TIER`: Tier 1 minimum dari Exploration guest credential/optional PII concerns; Planner menilai actual risk/verification boundaries dari current authority.

## Current-effective inputs / PRIOR_ARTIFACTS

Paths `WU-S2-*` relatif terhadap `.harscode-spaces/s2-guest-donation-truthful-state/`; lainnya relatif repository root.

- Completed frontend Exploration corpus: seluruh durable files di `WU-S2-004/runs/EXP-S2-004-001/evidence/`, saat dispatch `stage-2-gap-analysis.md` dan `stage-3-solutioning.md`; canonical fresh-session coverage berlaku.
- Work definition/current state: `WU-S2-004/manifest.md`, parent Outcome/Work Graph/Control Surface/events; project Authority Map dan profile. Work Graph owns dependency topology.
- Accepted Donation reconciliation baseline: `WU-S2-002/runs/TP-S2-002-015/techplan.md`, current `docs/spec/5-donation/`, authored `api/openapi/donation.yaml` plus referenced common components.
- Completed/accepted Campaign baseline: `WU-S2-005/manifest.md`, `WU-S2-005/runs/TP-S2-005-002/techplan.md`; current Campaign feature/INV-campaign-14 and authored `api/openapi/campaign.yaml`; generated types/fixture counterparts. Final owner acceptance di parent event; `RV-S2-005-002/review-findings-1.md`, `TST-S2-005-001/testing-report-1.md` dan `BLD-S2-005-002/report.md`/`handoff.md` memberi exact verification/readiness boundaries.
- Coordination delta: backend `WU-S2-003/runs/TP-S2-003-003/handoff.md`, `launch-record.md` dan exact D16/O1-REP pada `techplan.md`. Successor ini Draft/unapproved; cap/configuration/capacity/closure adalah recorded candidate owner direction yang membutuhkan Product authority/source reconciliation. Tidak mengubah current accepted Donation/Campaign API secara otomatis. Backend transaction/crypto implementation detail bukan default planning input frontend.
- Current Product/MVP, frontend AGENTS/architecture, UI/UX README routing, relevant design/page-map/pattern/component authority dan frontend execution profile dibuka sesuai active concern. Use actual repo command/config authority untuk verification plan.
- Canonical `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`, full template/rules/guardrails, best-practice router bila triggered, applicable workflow/orchestration AGENTS/run-contract/context dan overlay.

Historical frontend F-1 schema gap telah selesai melalui WU-S2-005; old Exploration shape/fixture descriptions adalah coordinates yang perlu live recheck. Jangan mengulang schema/route decision atau treat old unavailable-only schema sebagai current baseline. Optional email dan security/runtime concerns tetap dirutekan sesuai current sources.

## Task and completion condition

Jalankan canonical initial Techplan synthesis untuk Work Unit frontend ini, dari completed Exploration dan current live authorities. Tulis execution-grade `RUN_PATH/techplan.md` sebagai Draft / In Review, lengkap dengan rule/verification ownership/rationale, relevant Test Focus exact evidence anchors dan current open items. Ikuti canonical self-check dan independent Techplan review/decomposition recommendations; jangan memilih review/split outcome dari Invocation ini.

Frontend scope/milestone tetap `FRONTEND_MOCK_VERIFIED`: real backend integration, backend eligibility/security/D1 proof dan Slice finalization terpisah. Contract dependency WU-S2-005 telah satisfied; source Status-only BLD-002 delta independently hash/byte verified. Initial Draft synthesis runnable sementara backend material monetary/source work belum converged. Jika configuration limit/discovery/error/closure semantics materially memengaruhi whole frontend plan, catat exact dependency/Open Item yang menahan affected approval/Build; jangan mengadopsi candidate cap/reason/field sebagai accepted API, diam-diam membatasi input, atau mengarang konfigurasi/guest disclosure. Jangan expand accepted availability-only Campaign action untuk membawa limit tanpa owning contract route.

Readiness dan ordinary composition mengikuti current frontend/design authority. New material Product/Design/API decision harus decision-ready dan difasilitasi kepada owner yang benar, atau exact evidence/authority gap dilaporkan. Root fences/backend separation tetap berlaku; plan approval tidak memberi protected implementation permissions atau risk acceptance.

Stop setelah Draft dan canonical phase handoff/recommendations. `report-techplan.md` hanya saat current-effective plan benar-benar mencapai Human approval gate sesudah applicable review/resolution convergence; jangan buat report ketika material source/owner items masih churn. Tidak approve plan atau lanjut Build/decomposition/Review secara otomatis.

## Execution envelope

- `PREAUTHORIZED`: read named Exploration/current authorities/live frontend anchors; write hanya Run-local Draft Techplan/provenance/handoff. Faciliate bounded planning decisions sesuai current known authority; no automatic new authority attribution.
- `ORCHESTRATOR_DECISION`: reconcile planning findings/owner-source dependencies, route required authority/revision/review/report/approval sesuai actual handoff; coordinate backend/frontend contract deltas sebelum affected Build.
- `HUMAN_REQUIRED`: whole-Techplan approval, material owner/design/Product/API decisions, protected writes, rendered acceptance dan residual-risk/milestone gates.

Tidak mengubah Product/spec/API/generated/fixture/production/tests, prior Runs/plans atau projections. Tidak menjalankan tests/validator/generator/migration/services/browser/rendering dalam synthesis. Frontend tests/verification hanya direncanakan dari actual authority dan risk; execution mengikuti later authorized phase. Jangan memakai MSW untuk mengklaim runtime settlement/status/email atau security proof.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Planner / KC-PLANNER Session, `gpt-6-luna` / `medium`.

Kickoff: `Jalankan Techplan synthesis Run TP-S2-004-001 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TP-S2-004-001/invocation.md dan canonical ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md dengan orchestrated-run overlay. Tulis Draft Techplan dan phase handoff pada RUN_PATH, lalu berhenti.`

Laporkan completion atau exact scoped blocker kepada Orchestrator.
