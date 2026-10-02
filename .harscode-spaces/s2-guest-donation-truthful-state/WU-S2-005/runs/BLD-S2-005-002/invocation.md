# Run Invocation — `BLD-S2-005-002`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 oleh Orchestration Operator setelah Anhar menerima final authored Campaign/API contract. Human-facing prose: Bahasa Indonesia; preserve canonical terms/enums dan technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-005`
- `RUN_ID`: `BLD-S2-005-002`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/BLD-S2-005-002`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Implementer
- `SPECIALIZATION`: Status-only reconciliation of owner-accepted Campaign feature
- `PARTICIPANT_ID`: `P-S2-005-BL-002-1`
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`, observed committed target revision di bawah.
- `SESSION_TRANSITION`: `FRESH` — completed Build/Review/Testing occurrences tidak di-resume; new Participant/context untuk distinct acceptance metadata propagation.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796` plus current accepted working-tree artifacts.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`; ordinary applicable guidance current-effective.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `low`
- `MODEL_APPROVAL`: Not required by Human-owned local registry.
- `MODEL_ROUTING_RATIONALE`: Single Status-line replacement dengan explicit owner answer, accepted source hashes dan byte normalization; low cukup untuk bounded repository metadata reconciliation. Identity/hash/material discrepancy harus dilaporkan, bukan diselesaikan lewat stronger model atau substantive rewrite.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Implementer-owned metadata propagation setelah final Human authored acceptance; bukan correction dari Testing, tidak ada patch plan karena Testing tidak meminta source correction.

## Current-effective inputs / PRIOR_ARTIFACTS

- Approved `WU-S2-005/runs/TP-S2-005-002/techplan.md`, terutama scope/R8 dan final authored acceptance boundary.
- Completed Build `WU-S2-005/runs/BLD-S2-005-001/report.md`/`handoff.md`.
- Completed independent Code Review `WU-S2-005/runs/RV-S2-005-002/review-findings-1.md`: Approve, four passes tanpa findings.
- Completed independent Testing `WU-S2-005/runs/TST-S2-005-001/testing-report-1.md`, `handoff.md`, `launch-record.md`: Pass with flagged follow-ups; schema/generated/fixture evidence lulus, no patch plan. Warning discrepancy non-blocking; predicate/runtime evidence tetap downstream.
- Parent `events.md`: event “2026-10-01 — Anhar accepted final Campaign/API contract; metadata reconciliation queued”, explicit answer dan acceptance snapshot hashes.
- Current feature `docs/spec/4-campaign/features/02-campaign-detail-listing.md`; root AGENTS/profile/project Human authority, canonical `../harscode-workspace/workflow/3-build-prompt.md` dan applicable Build guidelines/checklist, run-contract/overlay/context. No raw Exploration input diperlukan.

Paths `WU-S2-*` relatif terhadap `.harscode-spaces/s2-guest-donation-truthful-state/`; lainnya repository-relative.

## Task and completion condition

Verifikasi explicit answer Anhar “Terima final kontrak Campaign/API (rekomendasi)” pada exact final authored acceptance question dalam parent event. Keputusan sudah berlaku; jangan minta approval kedua.

Reopen live artifacts dan cocokkan accepted snapshot SHA-256 sebelum edit:

| Artifact | SHA-256 |
|---|---|
| `docs/spec/4-campaign/features/02-campaign-detail-listing.md` | `b79a1dc891b614ff8750af988f3e78491881cdd9b6ed2e808ae766180ecc76e6` |
| `docs/spec/4-campaign/invariants.md` | `bec2ef663229249f3d032ede08b6833e08091e774d7b5b0261ee9f40deb2b9b6` |
| `api/openapi/campaign.yaml` | `65464f9aa177579a4a1536dc1d599160e3eab18eb1ab687feb9cdcf04c0f3134` |
| `api/openapi.yaml` | `15e0443600222d14356eca34f81ac0e54390bb9dbe6e9b7e75d42f909e24671c` |
| `frontend/lib/api/generated/openapi.ts` | `0f62528daf49b1e0265304b5752429e40150d9d58fff641e382dd99a15b59af8` |
| `frontend/mocks/fixtures/public-campaign.ts` | `9ca3554c1d6e6439d6225214a6b9afa3d708b5075f0fbb8075452f1c7158571d` |

Jika identity/hash tidak cocok, stop dan report exact discrepancy tanpa edit. Jika cocok, ubah hanya satu header line feature:

```text
> Status: Slice 1 reconciled; Slice 2 action reconciliation drafted (acceptance pending)
```

menjadi:

```text
> Status: Slice 1 reconciled; Slice 2 action contract accepted (runtime delivery pending)
```

Pertahankan setiap byte lain. Jangan menyatakan seluruh historical Campaign domain accepted; hanya current Slice-2 action reconciliation yang diterima. Invariants global header tidak diubah. API/generated/fixture content sudah accepted dan tidak memiliki pending-acceptance Status marker yang perlu diganti.

Record before/after hash, hasil byte-identical setelah Status dinormalisasi, unchanged hashes lima counterparts, dan exact-path `git diff --check`. Write Run-local `report.md`/`handoff.md` dengan provenance serta batas acceptance/runtime. Tidak menjalankan tests, validator, generator, migration, services atau runtime checks untuk metadata-only delta.

Stop setelah status-only report/handoff. Orchestrator akan memeriksa delta/evidence, reconcile WU completion/dependencies dan menyiapkan frontend Techplan. New independent Review/Testing applicability dievaluasi dari actual delta: bila hanya accepted Status line dan setiap byte lain unchanged, static metadata/non-behavior rationale mendukung N/A tanpa skipped-phase Run. Jangan self-waive phase atau mengubah substantive content; material drift harus dirutekan terpisah.

## Execution envelope

- `PREAUTHORIZED`: read accepted snapshot/owner evidence; write hanya satu Status line feature di atas dan Run-local metadata evidence; hash/byte normalization/scoped diff check.
- `ORCHESTRATOR_DECISION`: inspect completed metadata delta, phase applicability/completion reconciliation dan dependent frontend planning dispatch.
- `HUMAN_REQUIRED`: final authored acceptance sudah diberikan; any new material owner decision, protected production permission, runtime/risk/milestone acceptance tetap separate gates.

No spec behavior/acceptance criterion/invariant/API/generated/fixture/production/test/Techplan/prior Run/projection edits. No runtime producer implementation atau UI activation. Backend TP-S2-003-003 tetap last-known queued untuk dua financial/ordering findings; predicate source fidelity serta public/cache/auth/error/recheck/D1 runtime proof tidak selesai melalui status metadata ini.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Implementer / KC-IMPLEMENTER Session, `gpt-6-luna` / `low`.

Kickoff: `Jalankan Implementer Status-only reconciliation Run BLD-S2-005-002 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/BLD-S2-005-002/invocation.md dan canonical ../harscode-workspace/workflow/3-build-prompt.md dengan orchestrated-run overlay. Verifikasi final owner acceptance dan hashes, selaraskan hanya Status feature yang diterima, lalu berhenti setelah report dan phase handoff.`

Laporkan completion atau exact discrepancy kepada Orchestrator.
