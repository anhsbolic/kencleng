# Run Invocation — `OIR-S2-002-005`

Status: `READY_FOR_HUMAN_DISPATCH`  
Prepared by: Orchestration Operator  
Prepared: 2026-09-28  
Human-facing language: Bahasa Indonesia; retain canonical Harscode terms and code/API identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `OIR-S2-002-005`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-005`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Explorer
- `SPECIALIZATION`: O1 amount-contract evidence and owner-resolution facilitation
- `PARTICIPANT_ID`: `P-S2-002-OIR-005-1` — ephemeral identity for this Run only
- `PARTICIPANT_PROFILE_ID`: `KC-EXPLORER`
- `PARTICIPANT_PROFILE_PATH`: `.harscode-spaces/participant-profiles/profiles.md#kc-explorer`
- `PINNED_PROFILE_REVISION`: `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`
- `SESSION_TRANSITION`: `FRESH` — a new O1 money/contract decision surface after completed O7 Design review.
- `TARGET_BASE_REVISION`: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535`; current working-tree decisions and Run evidence are effective inputs. Re-read current files at dispatch.
- `WORKFLOW_REVISION`: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `RUNTIME_HARNESS`: `codex-cli`

## Why this Run exists

Approved `TP-S2-002-007` §13 leaves O1 authored wire/storage representation and derived-money precision open. Product already settled whole IDR input, minimum Rp5.000, Rp1 increments, Rp5.001 valid, and exact decimal storage/calculation without `float64`. Historical spec/OpenAPI/code may show conventions, but they cannot silently fix current Slice 2 contract. O7 is complete and O2/O3 verified-email retention is a separate scoped blocker. This Run reduces the remaining O1 uncertainty before a spec/API authoring Run; it does not re-open Product amount or method policy.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` — Approved baseline, especially §§4 Q2, 8, 9, 13 O1; SHA-256 `441577782f7600fd5507d228c4dd6491009c32d6e8b81607b28a38a5ebb1a1a3`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-001/resolution-brief.md` — current Product O1 direction and historical-gap pointers; SHA-256 `3d2b22f6fca9bd23024eca41cda79417bcc5cf929513223867d73ddcc6440782`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-004/design-review-brief.md` — completed O7 wording authority; SHA-256 `d9713e35324a0027814bc112237865cbeb5d86e1c3651e590aaca57f61ba29c4`.
- `.harscode-spaces/authority-map.md` — Anhar Solehudin owns Donation delivery/domain and API/contract for current Slice 2 only; SHA-256 `dd8cd32f65979f766034f6d086018e31ff4c6d584961ef1e296324c47c6eb06c`.
- Current Product/MVP: `docs/product/mvp-scope.md` Stage B and `docs/product/mvp-delivery-slices.md` §5. Read only relevant Donation spec and split `api/openapi/donation.yaml` plus referenced `common.yaml` components, current backend money conventions, and `../harscode-workspace/best-practices/go/decimal-and-money.md`. Historical lower-level artifacts are evidence, not current authority. Follow root/scoped `AGENTS.md` if entering stack scope.
- `.harscode-spaces/participant-profiles/profiles.md#kc-explorer`; canonical `../harscode-workspace/workflow/1-exploration-kickoff-prompt.md`, `workflow/orchestrated-run-overlay.md`, and Pilot #2 Open-Item Resolution guidance.

## Task and completion condition

Conduct bounded O1 evidence review of amount request/response encoding, validation for 5000/5001 and integer Rupiah, and existing exact-decimal persistence/calculation conventions. Identify whether current Slice 2 actually requires derived monetary values or rounding decisions now; keep tax rules outside Slice 2 unless Product evidence says otherwise. Compare viable contract representations against current conventions, compatibility, and precision risks. Distinguish decisions the named Donation/API owner can make from questions requiring Product Authority. Do not invent a scale, rounding rule, or storage migration solely because an old artifact contains one.

Use canonical Explorer Stage 1 plan announcement and Human confirmation before Stage 2. Write focused Stage 2 evidence under `RUN_PATH/evidence/`. At Stage 3, explain the smallest decision-ready O1 choice and evidence to Anhar in his current-Slice-2 Donation/API authority, then record only his explicit direction. If evidence supports a single convention and no material Human decision is required, state why; if still uncertain, state exactly which evidence/owner is needed. Avoid serial questions about out-of-scope derived-money features.

After the Stage 3 gate, write `RUN_PATH/amount-contract-brief.md` with evidence, `KEEP`/`ADAPT`/`REPLACE`/`DEFER` findings for relevant historical detail, explicit owner decisions (if any), unresolved dependencies, and the smallest next authoring route. Write terminal `RUN_PATH/handoff.md` with Run/Profile provenance, outcome, artifacts, verification performed/not performed, Findings, observed Decisions, Blockers, remaining concerns, next-route recommendation, and Learning Proposal or `None`.

Write only this Run's Stage 2 evidence, `amount-contract-brief.md`, and `handoff.md`. Do not modify Product/MVP, Design, spec/OpenAPI, code/tests, Approved Techplan, prior Run artifacts, or orchestration state. Do not claim money implementation proof, API acceptance, `CONTRACT_READY`, or residual risk acceptance. Keep O2/O3 verified-email retention conflict open and outside this Run.

## Execution envelope and model route

- `PREAUTHORIZED`: read relevant repository/Harscode evidence and write only assigned Run artifacts.
- `ORCHESTRATOR_DECISION`: material expansion beyond O1, Role/Profile changes, or next workflow route.
- `HUMAN_REQUIRED`: canonical Exploration Stage 1/Stage 3 gates and material current-Slice-2 Donation/API owner decisions. Record explicit decisions with date, scope, and provenance.
- `SELECTED_MODEL`: `gpt-6-luna`; `REASONING_EFFORT`: `high`; no extra model approval under `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: least-cost registered repository-reasoning model; `high` supports money representation and cross-contract comparison. Escalate only for demonstrated capability insufficiency.
- `PHASE_ROUTE`: Explorer Open-Item Resolution re-entry for O1. Canonical Exploration governs stages; Pilot #2 guidance governs facilitation and handoff.

## Human-assisted dispatch

1. Open a fresh Codex Participant Session in `/home/anhar-solehudin/kencleng-workspace/kencleng` with `gpt-6-luna` / `high`.
2. Paste: `Jalankan Run OIR-S2-002-005 sebagai Explorer untuk O1 amount-contract evidence dan owner resolution. Ikuti invocation durable .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-005/invocation.md, pinned Profile KC-EXPLORER, canonical Exploration prompt, orchestrated-run overlay, dan Pilot #2 Open-Item Resolution guidance. Mulai dengan Stage 1 plan announcement dan berhenti untuk Human confirmation sesuai canonical prompt. Tulis Stage 2 evidence, amount-contract-brief.md, dan terminal handoff.md hanya di RUN_PATH. Laporkan masalah material atau completion ke Orchestrator.`
3. At canonical Stage 1/Stage 3 gates, respond directly to Participant. Report only a material problem or terminal completion to Orchestrator; routine progress does not require monitoring.
