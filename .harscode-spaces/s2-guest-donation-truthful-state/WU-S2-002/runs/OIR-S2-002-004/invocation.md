# Run Invocation — `OIR-S2-002-004`

Status: `READY_FOR_HUMAN_DISPATCH`  
Prepared by: Orchestration Operator  
Prepared: 2026-09-28  
Human-facing language: Bahasa Indonesia; retain canonical Harscode terms and code/API identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `OIR-S2-002-004`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-004`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Explorer
- `SPECIALIZATION`: O7 Product Design terminology and source-label review facilitation
- `PARTICIPANT_ID`: `P-S2-002-OIR-004-1` — ephemeral identity for this Run only
- `PARTICIPANT_PROFILE_ID`: `KC-EXPLORER`
- `PARTICIPANT_PROFILE_PATH`: `.harscode-spaces/participant-profiles/profiles.md#kc-explorer`
- `PINNED_PROFILE_REVISION`: `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`
- `SESSION_TRANSITION`: `FRESH` — new O7 decision surface after completed O2/O3 Runs and current Human UX direction.
- `TARGET_BASE_REVISION`: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535`; current working-tree OIR evidence and orchestration Decisions are effective inputs. Re-read current files at dispatch.
- `WORKFLOW_REVISION`: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `RUNTIME_HARNESS`: `codex-cli`

## Why this Run exists

Approved Techplan `TP-S2-002-007` §13 O7 requires Design terminology/source-label review after Product-level O2 state meaning. That meaning is already established: backend simulator owns `pending`/`success`/`failed`, `pending` copy is “Menunggu hasil simulasi,” and no user-facing timing promise or real-payment instruction is approved. O3–O5 owner directions are recorded. Anhar, as current-Slice-2 Product Design owner, has also confirmed that optional guest-email verification and the 24-hour unverified-email deletion/no-notification consequence must be explained near email opt-in. These inputs make a bounded O7 review runnable without waiting for an O2 numeric deadline. The separate verified-email-while-`pending` retention conflict remains open and must not be decided in this Run.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` — Approved baseline, especially §13 O7, §4 Q4/Q6/Q7, and applicable Rule/Risk/Test Focus; SHA-256 `441577782f7600fd5507d228c4dd6491009c32d6e8b81607b28a38a5ebb1a1a3`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-001/resolution-brief.md` — Product O1–O6/O9 decisions and O7 continuation; SHA-256 `3d2b22f6fca9bd23024eca41cda79417bcc5cf929513223867d73ddcc6440782`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-002/resolution-brief.md` — O3–O5 owner direction, including fragment status URL and generic failure wording; SHA-256 `5de8e28170e9b5ebc3b1789010433589c03268d8b4e7e7923ed3ce8363e70e9e`.
- `.harscode-spaces/s2-guest-donation-truthful-state/events.md` — read the 2026-09-28 entries “Guest email verification disclosure confirmed for current Slice 2,” “Product/MVP retains terminal-notification obligation,” and prior O2/O3 conflict events. Treat the new disclosure as current-Slice-2 owner direction; keep the pending-retention conflict visibly open.
- `.harscode-spaces/authority-map.md` — Anhar Solehudin is current-Slice-2 Product Design owner; Product/MVP row is limited to the identified O3 question; SHA-256 `dd8cd32f65979f766034f6d086018e31ff4c6d584961ef1e296324c47c6eb06c`.
- Current Product/MVP: `docs/product/mvp-scope.md` Stage B/C and `docs/product/mvp-delivery-slices.md` §5. Current Design routing: `docs/ui-ux/README.md`, then applicable `product-design-principles.md`, `patterns.md` §7 and relevant nearby sections, and `design-guidelines.md` where a concrete visual/source-label rule is needed. Historical designs are evidence only.
- `.harscode-spaces/participant-profiles/profiles.md#kc-explorer`; canonical `../harscode-workspace/workflow/1-exploration-kickoff-prompt.md`, `workflow/orchestrated-run-overlay.md`, and Pilot #2 Open-Item Resolution guidance.

## Task and completion condition

Conduct a bounded O7 Design review of guest donation/status terminology and interaction consequences: pending “Menunggu hasil simulasi,” success/failed simulation source labels, QRIS-active versus other-methods-unavailable presentation, optional email opt-in/verification disclosure, terminal-only notice wording, generic unavailable status-link behavior, and explicit failed-state next action. Check each against current Product semantics and Design authority. Determine where existing guidance is sufficient and where a material Product Design decision is still required. The confirmed email disclosure must appear near opt-in in the recommended user-facing flow, not only in Terms & Conditions; its clock origin is email capture, not send time.

Use canonical Explorer Stage 1 plan announcement and Human confirmation before Stage 2. Write focused Stage 2 evidence under `RUN_PATH/evidence/`. At Stage 3, compare viable copy/interaction options and consequences; facilitate only bounded decisions with Anhar in his current-Slice-2 Product Design authority where evidence supports them. Do not re-open Product O1–O6/O9 or Security/PII policy. If a copy choice would imply a new delivery guarantee, new user-facing ETA, lost verified-email notification, real payment/settlement, or another Product-semantic change, mark that as a separate authority question instead of deciding it as Design.

After the Stage 3 gate, write `RUN_PATH/design-review-brief.md` with per-surface finding/outcome, source anchors, Design recommendations/decisions (if explicit), deferred dependency, and smallest next route. Write terminal `RUN_PATH/handoff.md` with Run/Profile provenance, outcome, artifacts, verification performed/not performed, Findings, observed Decisions, Blockers, remaining concerns, next-route recommendation, and Learning Proposal or `None`.

The Run must not edit Product/MVP, canonical Design docs, spec/OpenAPI, code/tests, Approved Techplan, prior Run artifacts, or orchestration state. It must not claim final visual acceptance, implementation, residual Security/PII risk acceptance, `CONTRACT_READY`, or another Work Unit milestone. The O2/O3 verified-email pending-retention conflict remains a separate scoped blocker.

## Execution envelope and model route

- `PREAUTHORIZED`: read relevant repository/Harscode evidence and write only this Run's Stage 2 evidence, `design-review-brief.md`, and `handoff.md`.
- `ORCHESTRATOR_DECISION`: material expansion beyond O7 Design review, change of Role/Profile, or next workflow route.
- `HUMAN_REQUIRED`: canonical Exploration Stage 1/Stage 3 gates and material current-Slice-2 Product Design decisions. Record explicit owner decisions with date, scope, and provenance; do not infer them.
- `SELECTED_MODEL`: `gpt-6-luna`; `REASONING_EFFORT`: `high`; no extra model approval under `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: least-cost registered repository-reasoning model; `high` supports cross-Product/Design wording and consequence analysis. Escalate only for demonstrated capability insufficiency.
- `PHASE_ROUTE`: Explorer Open-Item Resolution re-entry for O7, now runnable from settled Product state and explicit current Design direction. Canonical Exploration governs stages; Pilot #2 guidance governs facilitation and handoff.

## Human-assisted dispatch

1. Open a fresh Codex Participant Session in `/home/anhar-solehudin/kencleng-workspace/kencleng` with `gpt-6-luna` / `high`.
2. Paste: `Jalankan Run OIR-S2-002-004 sebagai Explorer untuk O7 Product Design review. Ikuti invocation durable .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-004/invocation.md, pinned Profile KC-EXPLORER, canonical Exploration prompt, orchestrated-run overlay, dan Pilot #2 Open-Item Resolution guidance. Mulai dengan Stage 1 plan announcement dan berhenti untuk Human confirmation sesuai canonical prompt. Tulis Stage 2 evidence, design-review-brief.md, dan terminal handoff.md hanya di RUN_PATH. Laporkan masalah material atau completion ke Orchestrator.`
3. At canonical Stage 1/Stage 3 gates, respond directly to Participant. Report only a material problem or terminal completion to Orchestrator; routine progress does not require monitoring.
