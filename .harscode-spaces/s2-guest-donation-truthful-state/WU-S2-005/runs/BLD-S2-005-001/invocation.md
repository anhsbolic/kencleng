# Run Invocation — `BLD-S2-005-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator setelah Human approval dan verified TP-S2-005-004 Status-only completion. Human-facing prose: Bahasa Indonesia; preserve canonical terms/enums dan technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-005`
- `RUN_ID`: `BLD-S2-005-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/BLD-S2-005-001`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Implementer
- `SPECIALIZATION`: Approved Campaign action spec/API reconciliation
- `PARTICIPANT_ID`: `P-S2-005-BL-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — first Build occurrence setelah completed Planner metadata Run; new Run/Participant/Session, bukan melanjutkan sesi Planner.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree; reopen live authorities/anchors sebelum edit.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`; ordinary applicable guidance tetap current-effective.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: Bounded authored response union/spec/fixture change dari independently reviewed Approved contract; architecture dan wire decisions sudah settled. Configured coding/repository-work capability pada medium effort cukup untuk edit/generation/fidelity loop. Ini shared-contract work dengan limited generated/mock frontend surfaces, bukan material UI Build. Resolve actual runtime model dari Human registry, bukan historical frontend profile model names yang tidak listed. Escalate only on demonstrated capability insufficiency, bukan missing authority/source/environment.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Initial Build against Approved TP-S2-005-002 tanpa decomposition/task snapshots; cohesive sequential reconciliation. No patch scope atau skipped-phase Run.
- `RISK_TIER`: Tier 1 — public response/consumer contract, public projection/anti-enumeration dan GET/POST eligibility authority boundary; required independent review/verification dan Human final authored acceptance tetap berlaku.

## Current-effective inputs / PRIOR_ARTIFACTS

- Sole material spine: `WU-S2-005/runs/TP-S2-005-002/techplan.md`, Status Approved; SHA-256 at preparation `13b84b3ce2e50573b2811b97f23d3454eaaead3c8cfaaad1d8ba31bb08ccd872`.
- Human report `WU-S2-005/runs/TP-S2-005-003/report-techplan.md`; clean independent Techplan Review `WU-S2-005/runs/RV-S2-005-001/review-findings.md`/handoff; approval event dan verified metadata `WU-S2-005/runs/TP-S2-005-004/launch-record.md`/handoff. Source §13 stale Active Review wording dijelaskan completed evidence; bukan pending Review/owner vote.
- Current WU manifest, parent events/Work Graph/Control Surface/Outcome, Authority Map, project tracker. Human “techplan approve bro” approved exact TP-S2-005-002; metadata normalization verified same source bytes. No later material Campaign action decision recorded.
- Root AGENTS, applicable project workflow risk/preconditions/contract coordination sections; current Product Slice 2/3, Campaign feature/invariant and Donation POST/D1 sources cited by spine. Raw Exploration tidak dibaca default; exact unresolved evidence hanya bila spine membutuhkannya.
- `api/README.md`, authored `api/openapi/campaign.yaml` plus referenced common components, `api/package.json`; generated bundle hanya untuk counterpart/cross-domain inspection bila perlu.
- `frontend/AGENTS.md`, applicable frontend API/mock/type-generation guidance dan `frontend/package.json`/fixtures/generated-type anchors untuk limited approved counterpart writes; UI/component/design production implementation tidak termasuk.
- Canonical `../harscode-workspace/workflow/3-build-prompt.md`, `workflow/3-build/guidelines.md`, `checklist.md`, orchestrated-run overlay/context dan applicable workflow/orchestration AGENTS/run-contract/profile.

Paths `WU-S2-*` relatif terhadap `.harscode-spaces/s2-guest-donation-truthful-state/`; lainnya relatif repository root. Overlay memakai approved spine di atas, output Build report/provenance pada RUN_PATH.

## Build target and scope

Execute Approved TP-S2-005-002 in full untuk contract-reconciliation scope; reopen live anchors dan ikuti settled rules/decisions. Tidak ada second ad-hoc task boundary atau task-specific new solution. Jika live source mematahkan material assumption/authority, stop/report canonical blocker; jangan redesign atau mengubah source plan.

Authorized authored writes sesuai approved anchors:

- `docs/spec/4-campaign/features/02-campaign-detail-listing.md`: action acceptance Slice 2 saja.
- `docs/spec/4-campaign/invariants.md`: action wording `INV-campaign-14` yang diperlukan; preserve `INV-campaign-13`/other invariants.
- `api/openapi/campaign.yaml`: affected public Campaign donation-action contract.
- `frontend/mocks/fixtures/public-campaign.ts`: required contract fixture correspondence saja, bukan UI activation atau runtime eligibility policy.
- `api/openapi.yaml` dan `frontend/lib/api/generated/openapi.ts`: regenerate lewat documented generators, never hand-edit.
- RUN_PATH `report.md` dan phase-owned provenance/handoff/launch-record bila diperlukan.

Human plan approval authorizes planned reconciliation drafting dalam scope ini; jangan menandai final authored owner acceptance dari approval plan. Preserve valid unrelated agreement/header status dan report any lifecycle ambiguity tanpa klaim acceptance baru. Tidak ada Product/MVP, Donation API/spec/D1, shared common/index operation changes, migration, backend production, frontend UI/data-access/handler production, tests, protected paths, prior Runs atau orchestration projections writes. Unexpected required additional file/scope/material decision dikembalikan ke Orchestrator/owning authority sebelum expansion.

Current predicate source-fidelity/runtime check tetap dimiliki authorized producer delivery owner: WU005 menulis contract saja. Jika source gap menghalangi accurate contract reconciliation, stop dengan exact gap; jangan mengarang eligible/ineligible condition, menambah public visibility, atau menjalankan producer implementation untuk mengisi schema. Backend financial Review findings pada WU003 tidak otomatis mengubah approved Campaign action contract.

## Focused Build verification and handoff

Ikuti Approved Testing Checklist ownership, `api/README.md` dan actual executable scripts. Establish validation warning baseline sebelum edit agar no-new-warning-coordinate rule dapat dinilai; zero validation errors diperlukan untuk touched contract. Run:

- workdir `api`: `npm run validate` sebelum/sesudah authored edit; compare actual warning coordinates, bukan hanya jumlah. Historical reported backlog bukan bukti current baseline.
- workdir `api`: `npm run bundle` setelah source berubah.
- workdir `frontend`: `npm run generate:api-types` dari regenerated bundle.
- Inspect affected authored/generated/fixture correspondence dan scoped diff. Focused static verification hanya bila credible edit loop memerlukannya; jangan broad suite/full UI build secara ritual.

Dependencies/tooling menggunakan current installed/configured repo environment; install/lockfile changes tidak otomatis scope, route actual environment gap bila diperlukan. Jangan add/run product/runtime tests, race/concurrency/perf/security suites, browser automation, migrations atau real integration pada contract-only Build ini. Record actual commands/results dan verification-scope confirmation; distinguish validation/generation dari final independent Testing, source predicate proof, runtime/security proof, owner acceptance dan readiness.

Stop setelah full target report/phase handoff. Recommend fresh independent Code Review dari actual authored/generated/fixture diff; subsequent verification/final Campaign/API owner acceptance follows applicable gates. Jangan auto-dispatch Review, accept residual risk, approve own changes, mark WU DONE/CONTRACT_READY, activate UI CTA atau claim production/runtime integration.

## Execution envelope

- `PREAUTHORIZED`: read approved spine/evidence/current routed sources; edit hanya authorized reconciliation paths; documented schema validation/bundle/type generation dan focused diff/correspondence inspection; Run-local Build report/provenance.
- `ORCHESTRATOR_DECISION`: reconcile outcome, route any required scope expansion/assumption break melalui owning authority dan independent review/verification setelah completion.
- `HUMAN_REQUIRED`: final authored spec/API acceptance, material policy/interface/authority changes, protected Tier-0 implementation, residual-risk acceptance dan milestone gates. Existing Techplan approval tidak ditanya ulang.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Implementer / KC-IMPLEMENTER Session; configured `gpt-6-luna` / `medium`. Start canonical Build prompt plus Invocation/overlay.

Kickoff: `Jalankan Build Run BLD-S2-005-001 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/BLD-S2-005-001/invocation.md dan canonical ../harscode-workspace/workflow/3-build-prompt.md dengan orchestrated-run overlay. Tulis Build report dan phase handoff pada RUN_PATH, lalu berhenti sesuai gate fase.`

Laporkan completion atau exact blocker ke Orchestrator.
