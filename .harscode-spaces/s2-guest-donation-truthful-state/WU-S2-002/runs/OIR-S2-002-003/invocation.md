# Run Invocation — `OIR-S2-002-003`

Status: `READY_FOR_HUMAN_DISPATCH`  
Prepared by: Orchestration Operator  
Prepared: 2026-09-28  
Human-facing language: Bahasa Indonesia; retain canonical Harscode terms and code/API identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `OIR-S2-002-003`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-003`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Explorer
- `SPECIALIZATION`: `Open-Item Decision Resolution / Product-Contract Facilitation`
- `PARTICIPANT_ID`: `P-S2-002-OIR-003-1` — ephemeral identity for this Run only
- `PARTICIPANT_PROFILE_ID`: `KC-EXPLORER`
- `PARTICIPANT_PROFILE_PATH`: `.harscode-spaces/participant-profiles/profiles.md#kc-explorer`
- `PINNED_PROFILE_REVISION`: `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`
- `SESSION_TRANSITION`: `FRESH` — new execution occurrence after terminal `OIR-S2-002-002`; the simulator/retention dependency is a distinct decision surface.
- `TARGET_BASE_REVISION`: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535`; Run evidence and orchestration reconciliation written after that commit are current working-tree inputs, pinned below. Re-read current files at dispatch.
- `WORKFLOW_REVISION`: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `RUNTIME_HARNESS`: `codex-cli`

## Why this Run exists

`OIR-S2-002-001` settled Product-level O2 state/recovery policy but left simulator timing and backend-controlled failure-scenario detail open. The current Approved Techplan and named Donation delivery owner now provide a narrower owner route. Terminal `OIR-S2-002-002` additionally established Security/PII email windows while explicitly deferring the maximum retention of verified email during `pending` until O2 timing evidence or an explicit Security/PII cap exists. This dependency is the meaningful delta. Analyze O2 and that linked O3 retention cap together; do not reopen O3's settled verification/retry/deletion windows.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` — Approved plan, especially §13 O2/O3 and related Rules/Risks/Test Focus; SHA-256 `441577782f7600fd5507d228c4dd6491009c32d6e8b81607b28a38a5ebb1a1a3`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-001/resolution-brief.md` — prior O2 Product decision and unresolved delivery detail; SHA-256 `3d2b22f6fca9bd23024eca41cda79417bcc5cf929513223867d73ddcc6440782`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-002/resolution-brief.md` — newly recorded O3 Security/PII decision and pending-retention dependency; SHA-256 `5de8e28170e9b5ebc3b1789010433589c03268d8b4e7e7923ed3ce8363e70e9e`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-002/handoff.md` — terminal execution evidence; SHA-256 `86d50645b32ec6eca36cb8694b254fc544873235158c48abbfe34981240748c6`.
- `.harscode-spaces/authority-map.md` — Anhar Solehudin's Donation delivery/domain and Security/PII authority is limited to current Slice 2; SHA-256 `65a73a5ab1233ff98699f425222d3cc9679e8c34176a9fcd3e1f60c9e20d433c`.
- `docs/product/mvp-scope.md` and `docs/product/mvp-delivery-slices.md` §5 — current state, demo failure, no timing promise, and guest-email purpose. Route through root `AGENTS.md` to the minimum applicable Donation spec, Design, backend architecture, and Harscode best-practice sections. Historical timing/probability are evidence only.
- `.harscode-spaces/participant-profiles/profiles.md#kc-explorer`, canonical `../harscode-workspace/workflow/1-exploration-kickoff-prompt.md`, `workflow/orchestrated-run-overlay.md`, and Pilot #2 Open-Item Resolution guidance.

## Task and completion condition

Prepare decision-ready evidence for the internal backend simulator timing and demo failure-scenario control in Techplan O2, then determine its consequence for the maximum retention of verified guest email while Donation remains `pending` (the open portion of O3). Preserve Product's `pending → success/failed` semantics, backend-owned outcome, demo-only failure, no public estimate/SLA, no donor-controlled outcome, and explicit failed-state recovery. Establish whether the available evidence supports a bounded simulator terminal deadline or whether a separate Security/PII retention cap is needed. If an item cannot be decided yet, state the exact missing evidence and smallest next route. Do not invent implementation behavior or turn historical 2–5 seconds/5% into current policy.

Use canonical Explorer Stage 1 plan announcement and Human confirmation before Stage 2. Write focused Stage 2 evidence under `RUN_PATH/evidence/`. At Stage 3, compare viable choices, consequences, and risks; ask Anhar for an explicit current-Slice-2 Donation delivery or Security/PII decision only where the evidence is sufficient. Keep any new user-facing timing promise, Product semantics change, or material Design expression outside this authority scope and route it to its owner. Do not pressure a decision if evidence is insufficient.

After the Stage 3 gate, write `RUN_PATH/resolution-brief.md` with distinct O2 and linked O3-pending-retention outcomes (`RESOLVED`, `PARTIALLY_RESOLVED`, `DEFERRED`, `NEEDS_OWNER`, or `NEEDS_FURTHER_EVIDENCE`), evidence, options/trade-offs, recommendation where justified, any explicit Human decision with date/scope/provenance, and the smallest next route. Write terminal `RUN_PATH/handoff.md` with Run/Profile revision, outcome, artifacts, verification performed/not performed, Findings, observed Decisions, Blockers, remaining concerns, next-route recommendation, and Learning Proposal or `None`.

The Run must not change Product/MVP, Design, spec, OpenAPI, code, tests, the Approved Techplan, other Run artifacts, or orchestration state. It must not accept residual risk for the owner, declare `CONTRACT_READY`, dispatch the next Run, or promote a Work Unit milestone. O1, O4–O9, and Campaign/Donation ordering remain later routes; mention them only where a real dependency exists.

## Execution envelope and model route

- `PREAUTHORIZED`: relevant repository/Harscode read-only investigation and writes only to this Run's Stage 2 evidence, `resolution-brief.md`, and `handoff.md`.
- `ORCHESTRATOR_DECISION`: material expansion beyond O2 and linked O3 retention, change of Run/Role/Profile, or next workflow route.
- `HUMAN_REQUIRED`: canonical Exploration Stage 1/Stage 3 gates and material current-Slice-2 Donation delivery or Security/PII decisions. Human decision evidence may be recorded by Explorer but not inferred.
- `SELECTED_MODEL`: `gpt-6-luna`; `REASONING_EFFORT`: `high`; no extra model approval under `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: least-cost registered model with repository reasoning capability suits this bounded cross-domain evidence and owner facilitation; `high` supports the timing/retention dependency. Escalate only for demonstrated capability insufficiency.
- `PHASE_ROUTE`: Explorer Open-Item Resolution re-entry for a bounded dependency left by the prior OIR Runs. Canonical Exploration governs stages; Pilot #2 guidance governs facilitation and handoff.

## Human-assisted dispatch

1. Open a fresh Codex Participant Session in `/home/anhar-solehudin/kencleng-workspace/kencleng` with `gpt-6-luna` / `high`.
2. Paste: `Jalankan Run OIR-S2-002-003 sebagai Explorer. Ikuti invocation durable .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-003/invocation.md, Profile KC-EXPLORER yang dipin di sana, canonical Exploration prompt, orchestrated-run overlay, dan Pilot #2 Open-Item Resolution guidance. Mulai dari Stage 1 plan announcement dan berhenti untuk Human confirmation sesuai canonical prompt. Fokus hanya O2 dan dependency retensi email O3 saat pending; tulis evidence Stage 2, resolution-brief.md, dan terminal handoff.md di RUN_PATH. Laporkan masalah material atau completion ke Orchestrator.`
3. At canonical Stage 1/Stage 3 gates, respond directly to Participant. Report only a material problem or terminal completion to Orchestrator; routine progress does not require monitoring.
