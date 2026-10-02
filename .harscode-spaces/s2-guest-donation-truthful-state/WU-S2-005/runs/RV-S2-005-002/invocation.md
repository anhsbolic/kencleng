# Run Invocation — `RV-S2-005-002`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator setelah Human melaporkan BLD-S2-005-001 selesai dan report/handoff serta actual scoped diff diperiksa. Human-facing prose: Bahasa Indonesia; preserve canonical terms/enums dan technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-005`
- `RUN_ID`: `RV-S2-005-002`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/RV-S2-005-002`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent Code Review — Campaign action spec/API/generated/fixture diff
- `PARTICIPANT_ID`: `P-S2-005-RV-002-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer context setelah completed Build; new Run/Participant terpisah dari Implementer dan prior Techplan Reviewer Run.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree; explicit diff scope di bawah, bukan semua dirty files.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`; ordinary applicable guidance tetap current-effective.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned local registry.
- `MODEL_ROUTING_RATIONALE`: Bounded six-file authored/generated/fixture public response change terhadap reviewed Approved contract; independent source/diff fidelity dan four-pass reasoning pada medium effort memadai. Tidak ada runtime transaction/crypto atau material UI implementation diff; escalate only on demonstrated capability insufficiency.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Initial independent Code Review, four passes terhadap actual current diff; completed RV-S2-005-001 adalah Techplan Review, bukan Code Review evidence.
- `RISK_TIER`: Tier 1 public contract/security authority boundary; Review verdict tidak menggantikan Human authored acceptance atau independent final verification.

## Current-effective inputs / PRIOR_ARTIFACTS

- Approved spine: `WU-S2-005/runs/TP-S2-005-002/techplan.md`; source Status Approved setelah verified TP-S2-005-004, Human approval di parent events.
- Build orientation/evidence: `WU-S2-005/runs/BLD-S2-005-001/report.md` dan `handoff.md`; actual live diff adalah review target, report bukan pengganti diff inspection.
- Parent WU manifest/events/Work Graph/Control Surface/Outcome, Authority Map, root AGENTS, relevant Product/MVP/Campaign/Donation/D1 authority yang dirutekan oleh Approved spine.
- `api/README.md`, authored Campaign source dan referenced common components; frontend AGENTS serta applicable generated-contract/fixture conventions. Backend sources hanya relevant live context untuk boundary yang benar-benar perlu fact check, tanpa production write.
- Canonical `../harscode-workspace/workflow/4-code-review-prompt.md`, `workflow/4-code-review/guidelines.md`, `checklist.md`, `best-practices/AGENTS.md` sebagai matching Pass-3 router, applicable workflow/orchestration AGENTS/run-contract/overlay/context.

Raw Exploration tidak default-loaded; material missing intent di spine dilaporkan sebagai drift, bukan disusun ulang dari history. Tidak ada decomposed task file atau patch plan input pada initial Review.

## Explicit current diff / changed-file scope

Review current working-tree diff against observed HEAD untuk enam files berikut:

1. `docs/spec/4-campaign/features/02-campaign-detail-listing.md`
2. `docs/spec/4-campaign/invariants.md`
3. `api/openapi/campaign.yaml`
4. `api/openapi.yaml`
5. `frontend/lib/api/generated/openapi.ts`
6. `frontend/mocks/fixtures/public-campaign.ts`

Existing unrelated dirty orchestration/Run evidence bukan implementation diff target. Reopen live scope sebelum review; metadata HEAD/workflow revisions adalah observed provenance, bukan policy pin. Jika material drift membuat assignment tidak cocok, laporkan exact drift, jangan memilih historical copy agar Review lulus.

## Task and completion condition

Jalankan canonical independent Code Review empat pass sesuai urutan Safety, Quality, Stack-Specific Best Practices, dan Consistency. Ikuti current checklist/materiality/verdict terhadap actual authored/generated/fixture diff dan Approved spine; tidak ada finding/verdict yang dipilih oleh Invocation ini. Semua pass tetap dilakukan walaupun pass sebelumnya clean.

Write `RUN_PATH/review-findings-1.md`, `patch-plan-1.md` hanya jika perubahan diperlukan, serta phase-owned provenance/handoff/launch-record bila diperlukan. Setiap finding memiliki location/problem/impact/suggested resolution dan blocking distinction. Stop setelah Review handoff; source fixes tetap fresh Build/Patch responsibility, tidak dilakukan oleh Reviewer.

Build melaporkan validation zero errors, unchanged 122 warning anchors, generators sukses dan scoped diff check; bedakan reported evidence dari actual independent review yang dijalankan. Jangan replay broad matrix/generation/validator hanya karena Review independent. Reasoning/source inspection default; targeted reproduction hanya untuk concrete suspected finding/uncertainty sesuai canonical authority dan envelope, report exact question/command/result. Jangan add/run product/runtime tests atau broad/security/concurrency/perf/browser suites/migrations pada scope ini.

Keep final authored Campaign/API acceptance, independent final verification/correspondence, downstream GET predicate source fidelity/public projection/runtime/POST evidence dan scoped dependencies explicit. Review Approve adalah verdict dalam scope Review, bukan Human owner acceptance, WU completion atau runtime milestone. No schema decision re-vote, production producer implementation, UI activation, risk acceptance atau source/plan edits.

## Execution envelope

- `PREAUTHORIZED`: read approved plan/evidence/current explicit diff dan routed authority; write hanya Review artifacts pada RUN_PATH; targeted substantiation menurut canonical Review posture bila diperlukan dan tetap dalam scope.
- `ORCHESTRATOR_DECISION`: reconcile verdict/findings, route new Build/Patch jika blocking; jika approved route independent Testing/verification dan Human final authored acceptance sesuai applicable gates.
- `HUMAN_REQUIRED`: material owner/authority decisions, final authored contract acceptance, Tier-0 implementation authorization, residual-risk acceptance dan milestones.

Reviewer tidak mengubah Product/spec/API/generated/fixture/production/tests, Approved plan/report/prior Runs atau orchestration projections. Tidak auto-dispatch Testing/Patch. Minor non-blocking polish tidak dipromosikan menjadi blocking loop.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Reviewer / KC-REVIEWER Session menggunakan `gpt-6-luna` / `medium`.

Kickoff: `Jalankan independent Code Review Run RV-S2-005-002 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/RV-S2-005-002/invocation.md dan canonical ../harscode-workspace/workflow/4-code-review-prompt.md dengan orchestrated-run overlay. Tulis Review artifacts pada RUN_PATH, lalu berhenti setelah phase handoff.`

Laporkan completion/verdict atau exact blocker kembali kepada Orchestrator.
