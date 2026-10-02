# Run Invocation — `TP-S2-005-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator setelah Exploration dan keputusan Anhar sebagai Campaign/API owner. Human-facing prose: Bahasa Indonesia; preserve canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-005`
- `RUN_ID`: `TP-S2-005-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/TP-S2-005-001`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: None; Campaign donation-action contract reconciliation
- `PARTICIPANT_ID`: `P-S2-005-TP-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — phase Techplan baru setelah Exploration; gunakan Run, Participant, dan Session baru.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree; buka sumber live yang relevan saat dispatch.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`; ordinary applicable guidance remains current-effective.
- `MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by `.harscode-spaces/.local-config.yaml`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `MODEL_ROUTING_RATIONALE`: Configured Planner model/effort untuk sintesis contract lintas spec/API dan generated consumers. Escalate hanya bila ada capability insufficiency yang terbukti.

## Current-effective inputs / PRIOR_ARTIFACTS

- `WU-S2-005/manifest.md`, parent `events.md`, `work-graph.md`, `outcome.md`, dan `control-surface.md` — route saat ini dan keputusan owner pada event “EXP-S2-005-001 inspected; owner selected availability-only”.
- `WU-S2-005/runs/EXP-S2-005-001/evidence/stage-2-gap-analysis.md` dan `evidence/stage-3-solutioning.md` — seluruh durable Exploration record; overlay menggantikan ordinal `1-exploration/logs/` dengan kedua file ini.
- `.harscode-spaces/authority-map.md`; root `AGENTS.md`; Product/MVP Slice 2; `docs/spec/4-campaign/features/02-campaign-detail-listing.md` dan applicable invariants; authored `api/openapi/campaign.yaml` beserta referenced shared components; `api/README.md`.
- Accepted Donation specs/API dan Approved `WU-S2-002/runs/TP-S2-002-015/techplan.md` sebagai batas submission/D1 yang relevan. Buka scoped stack authority ketika menggunakan fakta producer/consumer live.
- Current canonical `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`, Techplan template/rules/guardrails, orchestrated-run overlay, context-management, dan applicable workflow/orchestration AGENTS.

Paths `WU-S2-*` di atas relatif terhadap `.harscode-spaces/s2-guest-donation-truthful-state/`; sumber lain relatif terhadap repository root.

## Task and completion condition

Synthesize execution-grade Techplan untuk outcome rekonsiliasi WU-S2-005 dari seluruh Exploration evidence, keputusan owner yang tercatat, authority saat ini, serta live contract/conventions. Keputusan owner memilih availability-only projection: backend menilai eligibility saat GET, frontend menentukan route dari Campaign ID, dan POST mengecek ulang eligibility. Exact schema/reason vocabulary serta perubahan acceptance criteria tetap perlu dirumuskan dan ditinjau pada gate yang berlaku.

Output `RUN_PATH/techplan.md` memakai template canonical dengan status Draft / In Review, rule/testing traceability, sensitive Test Focus anchors, unresolved owner items, dan canonical phase handoff. Catat provenance melalui `launch-record.md` bila diperlukan workflow. Usulkan independent review/decomposition sesuai prompt; report baru dibuat ketika applicable review/resolution konvergen dan plan siap memasuki Human approval gate.

## Execution envelope

- `PREAUTHORIZED`: read current routed authority/live sources; write only this Run's Techplan dan phase-owned provenance/handoff artifacts.
- `ORCHESTRATOR_DECISION`: route review, report, approval, dan contract reconciliation setelah handoff.
- `HUMAN_REQUIRED`: Techplan approval; material Campaign/API/spec owner decisions; protected spec/invariant changes; final contract acceptance; residual-risk acceptance.

Preserve prior Run evidence. Production implementation tetap berada di backend/frontend delivery Work Units. Jangan mengubah Product/spec/API/generated files/code/tests/projections dalam sintesis ini, mengubah D1 atau approved Donation submission semantics, memilih implementation mechanism, atau menyatakan WU-S2-005 selesai hanya karena keputusan availability direction telah dibuat.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Start fresh Planner Participant Session dengan `gpt-6-luna` / `high`. Gunakan canonical Techplan synthesis prompt dan Invocation ini dengan orchestrated path semantics. Berhenti setelah Draft/In Review Techplan dan phase handoff.

Kickoff: `Jalankan canonical Techplan Synthesis untuk Run TP-S2-005-001 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/TP-S2-005-001/invocation.md. Gunakan ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md dengan input dan prior artifacts yang tercatat serta orchestrated-run overlay. Tulis artifact phase pada RUN_PATH dan berhenti setelah phase handoff.`
