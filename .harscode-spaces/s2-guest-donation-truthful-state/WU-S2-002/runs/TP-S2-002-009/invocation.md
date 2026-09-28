# Run Invocation — `TP-S2-002-009`

Status: `READY_FOR_HUMAN_DISPATCH`  
Prepared by: Orchestration Operator  
Prepared: 2026-09-28  
Human-facing language: Bahasa Indonesia; retain canonical Harscode terms and code/API identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `TP-S2-002-009`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-009`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `PRIOR_ARTIFACTS`: `TP-S2-002-007/techplan.md`, `OIR-S2-002-003/resolution-brief.md`, `OIR-S2-002-003/handoff.md`, and the exact Stage 2 evidence named below; `OIR-S2-002-002/resolution-brief.md` O3 is the preserved prior decision.
- `ROLE`: Planner
- `SPECIALIZATION`: Bounded O2 simulator delivery feasibility / O3 pending-retention decision preparation
- `PARTICIPANT_ID`: `P-S2-002-TP-009-1` — ephemeral identity for this Run only
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`
- `PARTICIPANT_PROFILE_PATH`: `.harscode-spaces/participant-profiles/profiles.md#kc-planner`
- `PINNED_PROFILE_REVISION`: `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`
- `SESSION_TRANSITION`: `FRESH` — Planner work follows a completed Explorer occurrence and must reconstruct from durable evidence.
- `TARGET_BASE_REVISION`: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535`; recent OIR evidence and orchestration state are current working-tree inputs after this commit. Re-read current files at dispatch.
- `WORKFLOW_REVISION`: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `RUNTIME_HARNESS`: `codex-cli`

## Why this Run exists

`OIR-S2-002-003` ended with `NEEDS_FURTHER_EVIDENCE` for O2 and `DEFERRED` for the linked O3 pending-email retention cap. Its evidence shows no Donation runtime or approved simulator maximum, so another Explorer discussion or immediate numeric Human question would repeat the same unresolved cause. The meaningful next step is Planner-owned delivery analysis: determine whether a bounded internal terminal policy can be specified and tested, and expose the smallest owner decision surface. This Run prepares a proposal; it does not revise or supersede the Approved Techplan.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` — Approved planning baseline, especially §4 Q4/R4/R5, §9, §12, and §13 O2/O3; SHA-256 `441577782f7600fd5507d228c4dd6491009c32d6e8b81607b28a38a5ebb1a1a3`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-003/resolution-brief.md` — current O2/O3 options and explicit lack of owner decision; SHA-256 `7b776e0b5f314101c55ec315f086d347ffe43486549e75cd66766b38601b817a`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-003/handoff.md` — terminal Explorer outcome; SHA-256 `3b0f2a41289f0e1a19fc8d76cab8d153e600f23e82d1958a726a04a7d89b3244`.
- `OIR-S2-002-003/evidence/stage-2-o2-simulator.md`, `stage-2-o3-pending-retention.md`, and `stage-2-input-provenance.md` — exact bounded gap evidence; resolve relative to that Run's `evidence/` directory.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-002/resolution-brief.md` O3 — existing 24-hour verification/retry decisions and deferred pending cap; SHA-256 `5de8e28170e9b5ebc3b1789010433589c03268d8b4e7e7923ed3ce8363e70e9e`.
- `.harscode-spaces/authority-map.md` — Anhar Solehudin is current-Slice-2 Donation delivery/domain and Security/PII owner only; SHA-256 `65a73a5ab1233ff98699f425222d3cc9679e8c34176a9fcd3e1f60c9e20d433c`.
- Current Product/MVP: `docs/product/mvp-scope.md` and `docs/product/mvp-delivery-slices.md` §5. Read relevant root/backend `AGENTS.md`, backend architecture, current code anchors, Donation historical spec/API as evidence, and targeted Harscode Techplan guidance. Canonical planning entrypoint is `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`; `workflow/orchestrated-run-overlay.md` maps phase paths into this Run. This assignment is a bounded planning artifact, not full Techplan synthesis.

## Task and completion condition

Write `RUN_PATH/o2-delivery-proposal.md` as decision preparation for an internal simulator terminal policy, scoped to O2 and its direct O3 pending-retention consequence. The proposal should establish whether a finite internal terminal bound is feasible without a public ETA/SLA; what start/end events and recovery behavior would make such a bound meaningful under missing enqueue, process restart, timeout/error, and replay; and how demo-only failure remains backend-configured/fixture-controlled and unavailable to donor/UI inputs. Compare viable bounded approaches and a separate Security/PII email-cap route. State where a proposed numeric duration would need owner decision and what evidence could justify it; do not choose an unsupported value.

For each viable approach, state the consequence for verified-email retention while `pending`, including expiry before terminal state and later notification eligibility. Identify any Product-semantic conflict that would require Product Authority rather than changing the promise in this proposal. Preserve the existing 24-hour verification and post-terminal retry decisions. Record enough code/architecture anchors, risks, and later Build/Testing evidence obligations for the Donation delivery and Security/PII owner to decide the next step. If no bounded proposal is feasible from current sources, state the precise evidence gap and route to a separate Security/PII cap decision; do not loop back to Explorer for the same question.

Write `RUN_PATH/handoff.md` with Run/Profile provenance, outcome, artifact pointer, what was inspected, verification performed/not performed, Findings, Decisions observed or `None`, remaining Blockers, materiality/review recommendation, and next-route recommendation. Do not generate `techplan.md` or `report-techplan.md` in this Run; the Approved `TP-S2-002-007` remains current. If a future owner decision materially changes that plan, Orchestrator will route a distinct amendment/review/approval path. Do not author spec/OpenAPI/code/tests, change Product/Design/authority, accept residual risk, dispatch another Run, or claim `CONTRACT_READY`.

## Execution envelope and model route

- `PREAUTHORIZED`: read relevant repository/Harscode evidence and write only `o2-delivery-proposal.md` and `handoff.md` in this Run path.
- `ORCHESTRATOR_DECISION`: material expansion outside O2/direct O3 dependency, plan amendment, change of Role/Profile, or next workflow route.
- `HUMAN_REQUIRED`: any new Donation delivery/Security/PII/Product decision. The Planner should prepare a bounded question; it may record an explicit owner decision if one occurs in Session but must not infer or force one.
- `SELECTED_MODEL`: `gpt-6-luna`; `REASONING_EFFORT`: `high`; no extra model approval under `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: least-cost registered repository-reasoning model; `high` is appropriate for a bounded reliability/PII dependency across architecture and delivery. Escalate only for demonstrated capability insufficiency.
- `PHASE_ROUTE`: Planner analysis, following completed Explorer evidence. This is a narrow planning artifact permitted by `KC-PLANNER` and the project-defined Run path; it is not a Techplan amendment or approval gate.

## Human-assisted dispatch

1. Open a fresh Codex Participant Session in `/home/anhar-solehudin/kencleng-workspace/kencleng` with `gpt-6-luna` / `high`.
2. Paste: `Jalankan Planner Run TP-S2-002-009 sesuai invocation durable .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-009/invocation.md. Gunakan pinned Profile KC-PLANNER, canonical Techplan guidance dan orchestrated-run overlay sesuai batas Run. Siapkan hanya o2-delivery-proposal.md dan handoff.md untuk O2 simulator serta dependency retensi email O3; jangan revisi Approved Techplan, spec/API, code, atau tests. Laporkan masalah material atau completion ke Orchestrator.`
3. Report only a material problem or terminal completion to Orchestrator; routine progress does not require monitoring.
