# Run Invocation — `TST-S2-005-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 oleh Orchestration Operator setelah completed RV-S2-005-002 menghasilkan Approve tanpa findings. Human-facing prose: Bahasa Indonesia; preserve canonical terms/enums dan technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-005`
- `RUN_ID`: `TST-S2-005-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/TST-S2-005-001`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Verifier
- `SPECIALIZATION`: Independent contract/schema/generated/fixture correspondence verification
- `PARTICIPANT_ID`: `P-S2-005-TST-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-VERIFIER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`, observed committed target revision di bawah.
- `SESSION_TRANSITION`: `FRESH` — independent Verifier setelah completed Build dan Code Review; new Run/Participant/context.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree; scope enam files di bawah.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`; ordinary applicable guidance tetap current-effective, bukan policy pin.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned local registry.
- `MODEL_ROUTING_RATIONALE`: Bounded contract-artifact verification dengan actual repo validation/generation authority dan rule traceability; medium cukup untuk conditional schema, generated correspondence dan public boundary consistency. Tidak ada production transaction/crypto atau material UI verification dalam scope; escalate hanya atas demonstrated capability insufficiency. Model dipilih dari current Human registry, bukan historical frontend model labels yang tidak tersedia di registry.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Initial independent Testing pada observable contract/artifact boundary setelah Code Review Approve.
- `RISK_TIER`: Tier 1 public contract; final authored Human acceptance dan downstream runtime/security evidence tetap terpisah.

## Current-effective inputs / PRIOR_ARTIFACTS

- Approved spine: `WU-S2-005/runs/TP-S2-005-002/techplan.md`; verified Status Approved. Rules & Validation, Testing Checklist dan Test Focus Pointer tetap primary verification contract.
- Latest Build claims/gaps: `WU-S2-005/runs/BLD-S2-005-001/report.md` dan `handoff.md`.
- Completed Code Review: `WU-S2-005/runs/RV-S2-005-002/review-findings-1.md`, termasuk embedded phase handoff; Approve, empat pass tanpa findings, tidak mengeksekusi verification commands.
- Owner/report evidence bila needed: `WU-S2-005/runs/TP-S2-005-003/report-techplan.md`, parent `events.md`, manifest/Work Graph/Control Surface/Authority Map. Valid schema proposal/Techplan approval tidak di-vote ulang; final authored acceptance belum diberikan.
- Actual command authority: `api/README.md`, `api/package.json`, `frontend/package.json`, relevant root/frontend AGENTS dan generated-contract/fixture conventions.
- Canonical `../harscode-workspace/workflow/5-testing-prompt.md`, `workflow/5-testing/guidelines.md`, `checklist.md`, orchestrated-run overlay/run-contract/context guidance. Best-practice router dibuka hanya untuk matching concern.

Raw Exploration bukan blanket input. Buka exact relevant Test Focus anchors dari Approved spine; jangan menyusun ulang intent dari historical corpus. Open Item tentang RV-S2-005-001 pada source plan adalah historical planning snapshot; current completed Techplan Review/report/approval dan Code Review dijelaskan oleh named evidence di atas.

## Observable boundary and scope

WU ini mengubah kontrak/spec beserta counterparts, tanpa production implementation. Nearest meaningful observable boundary adalah authored OpenAPI/schema, generated bundle/types, fixture dan correspondence consumer. Jelaskan batas ini dalam report; contract inspection tidak membuktikan runtime GET/POST, header/body/timing parity atau D1.

Target actual current working-tree diff against observed HEAD untuk:

1. `docs/spec/4-campaign/features/02-campaign-detail-listing.md`
2. `docs/spec/4-campaign/invariants.md`
3. `api/openapi/campaign.yaml`
4. `api/openapi.yaml`
5. `frontend/lib/api/generated/openapi.ts`
6. `frontend/mocks/fixtures/public-campaign.ts`

Read known consumers/routed authority sesuai kebutuhan, tanpa production edits. Unrelated dirty orchestration/Run files bukan implementation verification target. Material assignment drift dilaporkan, bukan diatasi dengan memilih historical source yang cocok.

## Task and completion condition

Ikuti canonical Testing: Step 0 sweep Build claims/gaps, every Rules & Validation row, applicable Test Focus exact anchors, error/compatibility review, actual required final verification dan fresh whole-Techplan consistency read. Derive coverage/results/verdict secara independen; Invocation tidak memilih Pass atau menutup gap.

Testing-owned contract rows membutuhkan independent evidence. Human-owned schema/scope/final acceptance tetap explicit Human gates; proposal/plan approval tidak sama dengan final authored contract acceptance. Catat runtime parts yang tidak dapat exercised pada artifact boundary, exact downstream owner/WU dan risk; jangan silently mark covered. Current GET producer masih menggunakan old action menurut Build/Review; production correspondence dan authoritative predicate fidelity wajib downstream, bukan dibuat dalam Run ini. Tidak mengarang public-but-ineligible scenario atau membuka ulang chosen wire schema.

Gunakan repo-authored OpenAPI validation dan scoped independent generation/correspondence sesuai Approved R1/R7 serta API README. Build melaporkan zero errors, unchanged 122 warning anchors dan generator success: ini claims yang perlu independent verification/spot-check, bukan hasil Run ini. Evaluasi warning-coordinate drift terhadap actual historical baseline, bukan count saja. Reproduction generator harus memakai temporary output/workspace (misalnya `/tmp`) agar source/generated/fixture files tetap unchanged; jangan hand-edit generated artifacts. Participant memilih command mechanics dari actual repo authority dan mencatat command/result.

Tidak add/run product/runtime tests, broad frontend suites, browser/security/race/concurrency/performance suites, migrations atau services pada contract-only scope ini. Bila concrete authority/risk memerlukan verification di luar envelope, report exact gap dan route; jangan mengurangi threshold atau mengklaim verification tersebut lulus. Material UI/rendered acceptance tidak dibuktikan oleh fixture.

Write `RUN_PATH/testing-report-1.md`, `patch-plan-1.md` hanya bila correction diperlukan, dan phase-owned provenance/handoff/launch-record sesuai canonical format. Bedakan verified, assumed, deferred dan not tested; classify blocking/non-blocking/new/regression findings. Stop setelah phase handoff; no source fix, auto-dispatch, Human acceptance atau milestone promotion.

## Execution envelope

- `PREAUTHORIZED`: read current named artifacts/six-file diff/routed authority; independent in-scope schema validation, temporary generation/correspondence dan artifact inspection; write hanya Run-local verification evidence serta disposable temporary outputs. Report commands yang benar-benar dijalankan dan source preservation.
- `ORCHESTRATOR_DECISION`: reconcile verdict, route fresh Build/Patch dari specific patch plan bila necessary; setelah sufficient contract evidence, facilitate final authored Campaign/API Human acceptance. Completion WU/dependency reconciliation dilakukan terpisah dari Verifier verdict.
- `HUMAN_REQUIRED`: final authored Campaign/API acceptance, material owner/authority decisions, protected implementation permissions, residual-risk acceptance dan delivery milestones.

Tidak mengubah Product/spec/API/generated/fixture/production/tests, Approved plan/prior Run artifacts atau orchestration projections. Tidak menganggap independent verification sebagai izin PR/publish/merge. Backend TP-S2-003-003 last-known queued untuk dua Review blockers; WU-S2-004 dan Campaign producer tetap menunggu accepted WU-S2-005 result.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Verifier / KC-VERIFIER Session, `gpt-6-luna` / `medium`.

Kickoff: `Jalankan independent Testing Run TST-S2-005-001 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/TST-S2-005-001/invocation.md dan canonical ../harscode-workspace/workflow/5-testing-prompt.md dengan orchestrated-run overlay. Tulis Testing artifacts pada RUN_PATH, lalu berhenti setelah phase handoff.`

Laporkan completion/verdict atau exact blocker kembali kepada Orchestrator.
