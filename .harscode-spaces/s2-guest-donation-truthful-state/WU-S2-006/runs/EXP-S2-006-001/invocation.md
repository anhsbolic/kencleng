# Run Invocation — `EXP-S2-006-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 oleh Orchestration Operator setelah completion backend TP-S2-003-003 dan kebutuhan owning-source reconciliation O1-REP dikonfirmasi dari durable artifacts.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-006`
- `RUN_ID`: `EXP-S2-006-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/EXP-S2-006-001`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Explorer
- `SPECIALIZATION`: Monetary policy / shared contract source-reconciliation Exploration
- `PARTICIPANT_ID`: `P-S2-006-EXP-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-EXPLORER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`, observed committed target revision di bawah.
- `SESSION_TRANSITION`: `FRESH` — distinct shared-source Work Unit/Exploration; backend Planner occurrence sudah selesai.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`; ordinary applicable guidance current-effective, bukan policy pin.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned local registry.
- `MODEL_ROUTING_RATIONALE`: Bounded initial repository/authority exploration dengan recorded decision/source anchors; medium cukup sebagai initial effort. Missing material authority/evidence dirutekan, tidak diatasi dengan model escalation. Reconsider hanya bila demonstrated capability insufficiency.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Canonical Exploration kickoff, Stage 1 only initially; Human confirmation sebelum Stage 2, lalu separate Stage-2 confirmation sebelum Stage 3.
- `RISK_TIER`: Dinilai Explorer dari actual financial/domain/interface concern; tidak ada implementation/risk acceptance dari Invocation.

## Canonical kickoff inputs

- `{CODEBASE_CONTEXT}`: `Kencleng — Go backend + Next.js frontend`
- `{TASK}`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/manifest.md` — Definition/Outcome/Scope/Boundaries dan referenced requirement evidence.
- Ticket / Area: `WU-S2-006` / monetary limits, Campaign capacity dan shared source reconciliation.
- Canonical prompt: `../harscode-workspace/workflow/1-exploration-kickoff-prompt.md`.
- Applicable `workflow/1-exploration/guidelines.md`, `sniffing-checklist.md`, root/project authority, workflow/orchestration AGENTS/run-contract/context dan orchestrated-run overlay.

## Current-effective inputs / PRIOR_ARTIFACTS

- Work Unit manifest dan parent events/Work Graph/Outcome/Control Surface; current Authority Map.
- Backend `WU-S2-003/runs/TP-S2-003-003/handoff.md`, `launch-record.md`, exact D16/O1-REP dan relevant referenced dependencies dalam Draft `techplan.md`; RV-S2-003-001 exact monetary representability finding bila diperlukan. Paths WU-S2-* relatif parent Slice-2 space.
- Accepted baseline WU-S2-002 DONE/CONTRACT_READY serta WU-S2-005 DONE/final Campaign action acceptance; sources dibuka dari current root routing sesuai concern, bukan dari historical prose saja.

Backend Participant choices tetap preserved. Product owner attribution kini mapped; previous handoff yang menyebut owner/source gap adalah historical context, bukan alasan meminta attribution ulang. Detail policy/interface belum diputuskan tetap gap yang harus dibuat decision-ready melalui canonical stages, bukan ditebak. Tidak menganggap Draft atau attribution sebagai owning-source acceptance.

## Task and completion condition

Jalankan canonical Stage 1: pahami assigned requirement, identify areas/order/rationale dari current authority, tanpa deep implementation inspection atau solution proposal. Stop untuk Human confirmation sebelum Stage 2. Jangan menggabungkan stages atau menganggap orchestration dispatch sebagai Stage-1/2 approval.

Sesudah applicable confirmations, canonical Stage 2/3 menulis durable evidence di `RUN_PATH/evidence/`, dengan concrete findings/progression effects dan known-owner decisions bila decision-ready. Jangan re-vote pilihan valid hanya karena berpindah Session; genuine unresolved material detail tetap difasilitasi sesuai owning authority. Preserve actual sources/decisions dan provenance tanpa invented runtime/session metadata.

Explorer tidak mengubah authoritative Product/spec/API/generated/fixture/production/tests atau successor backend plan. Completion Stage 3 mengakhiri Run; later Techplan memakai new Run/Participant/fresh context. No predetermined findings, ranges, field names, config ownership, closure precedence, proposal atau accepted source edits diberikan oleh Invocation.

## Execution envelope

- `PREAUTHORIZED`: read routed current authorities dan relevant source menurut canonical stage; write hanya Run-local Exploration evidence/handoff/provenance. Stage 1 sekarang; later stages setelah Human confirmations.
- `ORCHESTRATOR_DECISION`: reconcile evidence/owner-source dependencies dan applicable source/planning routes dari actual handoff.
- `HUMAN_REQUIRED`: Stage 1/2 confirmations, material owner/Product/spec/API decisions dan source acceptance, protected writes/risk/delivery gates.

Tidak menjalankan tests, validator/generator, services, migrations, browser atau runtime checks dalam Exploration. Tidak auto-approve, self-waive gate, mulai Build atau mengubah projections/prior artifacts. Root Tier-0 fences tetap berlaku. Frontend TP-S2-004-001 tetap prepared initial Draft; affected final approval/Build menjaga dependency source WU006.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Explorer / KC-EXPLORER Session, `gpt-6-luna` / `medium`.

Kickoff: `Jalankan Exploration Run EXP-S2-006-001 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/EXP-S2-006-001/invocation.md dan canonical ../harscode-workspace/workflow/1-exploration-kickoff-prompt.md dengan orchestrated-run overlay. Mulai Stage 1 dan berhenti untuk Human confirmation sebelum Stage 2.`

Stage confirmations diberikan dalam Participant Session. Laporkan durable stage handoff/completion atau material blocker kepada Orchestrator; Stage-3 completion mengakhiri Run.
