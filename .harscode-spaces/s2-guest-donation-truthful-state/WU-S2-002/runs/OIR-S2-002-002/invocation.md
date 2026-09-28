# Run Invocation — `OIR-S2-002-002`

Status: `READY_FOR_HUMAN_DISPATCH`  
Prepared by: Orchestration Operator  
Prepared: 2026-09-28  
Human-facing language: Bahasa Indonesia; retain canonical Harscode terms and code/API identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `OIR-S2-002-002`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-002`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Explorer
- `SPECIALIZATION`: `Open-Item Decision Resolution / Product-Contract Facilitation`
- `PARTICIPANT_ID`: `P-S2-002-OIR-002-1` — ephemeral identity for this Run only
- `PARTICIPANT_PROFILE_ID`: `KC-EXPLORER`
- `PARTICIPANT_PROFILE_PATH`: `.harscode-spaces/participant-profiles/profiles.md#kc-explorer`
- `PINNED_PROFILE_REVISION`: `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32` (present in target commit `17a34cb31b017308c0a426e3191b73d968f52454`)
- `SESSION_TRANSITION`: `FRESH` — new execution occurrence after the completed OIR and Techplan Runs; cross-security/API evidence and owner discussion require a clean context.
- `TARGET_BASE_REVISION`: `17a34cb31b017308c0a426e3191b73d968f52454`; current working tree contains the bounded Authority Map reconciliation. Re-read current files at dispatch; do not infer their contents from this base commit.
- `WORKFLOW_REVISION`: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `RUNTIME_HARNESS`: `codex-cli`

## Why this Run exists

`OIR-S2-002-001` was a broad decision facilitation before the 2026-09-26 Product/MVP amendment and current-effective approved `TP-S2-002-007`. It left O3–O5 technical Security/PII and API detail unresolved. The five relevant Slice 2 authority areas now have named owners; Anhar Solehudin is both Security/PII and API/contract owner for this Slice, without any control/risk decision being implied. This new, narrower evidence/decision surface is the meaningful delta for Explorer re-entry. The Run prepares decisions; it does not repeat settled Product policy or author the contract.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` — Approved plan, especially §13 O3–O5 and related Rules/Risks/Test Focus. Assignment-defining SHA-256: `441577782f7600fd5507d228c4dd6491009c32d6e8b81607b28a38a5ebb1a1a3`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-001/resolution-brief.md` — prior per-item evidence and unresolved technical follow-up; SHA-256: `3d2b22f6fca9bd23024eca41cda79417bcc5cf929513223867d73ddcc6440782`.
- `.harscode-spaces/authority-map.md` — current owner scopes. Assignment-defining SHA-256: `65a73a5ab1233ff98699f425222d3cc9679e8c34176a9fcd3e1f60c9e20d433c`.
- `docs/product/mvp-scope.md` and `docs/product/mvp-delivery-slices.md` §5 — approved guest email, status URL, and public failure behavior. Route from root `AGENTS.md` to the smallest relevant Donation spec/threat, API, backend architecture, Design, and best-practice sections; treat unreconciled historical detail as evidence.
- `.harscode-spaces/participant-profiles/profiles.md#kc-explorer`, the canonical `../harscode-workspace/workflow/1-exploration-kickoff-prompt.md`, `workflow/orchestrated-run-overlay.md`, and Pilot #2 Open-Item Resolution guidance in `../harscode-workspace/orchestration/pilot-2-candidate/orchestrator-operating-model.md`.

## Task and completion condition

Prepare a decision-ready, evidence-backed handoff for current Techplan O3–O5: guest email verification/retention/delivery retry; status URL/token exposure and residual risk; and generic status-failure transport parity/anti-enumeration. Determine which details can be chosen by the named Security/PII and API owner after the available evidence, which need more specialist or later Build/Testing evidence, and how those decisions depend on each other. Keep the approved Product semantics fixed. Do not silently import historical 401/404, token, timing, email, or API choices as current decisions.

Use the canonical Explorer Stage 1 plan announcement and its Human confirmation gate before Stage 2. Then inspect only the evidence needed for this bounded area; separate current state/gaps from options. In Stage 3, compare viable options, consequences, and residual risks. When a material question becomes decision-ready and Anhar is present in the Participant Session, facilitate an explicit decision under his current Slice 2 Security/PII or API/contract authority. Do not pressure a decision if evidence is insufficient.

Write canonical Stage 2 gap evidence under `RUN_PATH/evidence/`. After the Stage 3 gate, write `RUN_PATH/resolution-brief.md` with an outcome for each O3–O5 (`RESOLVED`, `PARTIALLY_RESOLVED`, `DEFERRED`, `NEEDS_OWNER`, or `NEEDS_FURTHER_EVIDENCE`), exact evidence, options/trade-offs, recommendation where justified, named decision owner and authority context, any explicit Human decision with date/scope/provenance, and the smallest next route. Write a compact terminal `RUN_PATH/handoff.md` identifying the Run/Profile revision, outcome, artifacts, verification performed/not performed, Findings, observed Decisions, Blockers, remaining concerns, next-route recommendation, and Learning Proposal or `None`. These are Participant-owned artifacts; the Orchestrator later reconciles them.

The Run must not change Product/MVP, Design, spec, OpenAPI, code, tests, the approved Techplan, other Run artifacts, or orchestration state. It must not accept residual risk on behalf of the owner, invent a control/window, declare `CONTRACT_READY`, dispatch the next Run, or promote a Work Unit milestone. O1/O2/O6–O9 remain visible as dependencies or later routes but are outside this Run's resolution scope.

## Execution envelope and model route

- `PREAUTHORIZED`: read the relevant repository/Harscode evidence and write only this Run's canonical Stage 2 evidence, `resolution-brief.md`, and `handoff.md`; reversible investigation as needed within the Explorer phase.
- `ORCHESTRATOR_DECISION`: material expansion beyond O3–O5, change of Run/Role/Profile, or next workflow route.
- `HUMAN_REQUIRED`: canonical Exploration Stage 1/Stage 3 gates, material Security/PII or API authority decisions, and residual-risk acceptance. Human decision evidence may be recorded by the Explorer but not inferred.
- `SELECTED_MODEL`: `gpt-6-luna`; `REASONING_EFFORT`: `high`; approval is not required by `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: the least-cost registry model with repository reasoning capability covers this bounded cross-security/API evidence task; `high` supports interdependent risks. Escalate only for demonstrated capability insufficiency, not missing owner/evidence/context.
- `PHASE_ROUTE`: Explorer Open-Item Resolution re-entry justified by current approved plan, named owners, and narrower unresolved O3–O5 controls. Canonical Exploration governs execution stages; Pilot #2 guidance governs facilitation and handoff.

## Human-assisted dispatch

1. Open a fresh Codex Participant Session in `/home/anhar-solehudin/kencleng-workspace/kencleng` with `gpt-6-luna` / `high`.
2. Paste: `Jalankan Run OIR-S2-002-002 sebagai Explorer. Ikuti invocation durable .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-002/invocation.md, Profile KC-EXPLORER yang dipin di sana, canonical Exploration prompt, orchestrated-run overlay, dan Pilot #2 Open-Item Resolution guidance. Mulai dari Stage 1 plan announcement dan berhenti untuk Human confirmation sesuai canonical prompt. Fokus hanya O3–O5; tulis evidence Stage 2, resolution-brief.md, dan terminal handoff.md di RUN_PATH. Laporkan masalah material atau completion ke Orchestrator.`
3. At the canonical Stage 1/Stage 3 gate, respond directly to the Participant. Report only a material problem or terminal completion back to the Orchestrator; routine progress does not require monitoring.
