# Run Invocation — `OIR-S2-002-006`

Status: `READY_FOR_HUMAN_DISPATCH`  
Prepared by: Orchestration Operator  
Prepared: 2026-09-28  
Human-facing language: Bahasa Indonesia; retain canonical Harscode terms and code/API identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `OIR-S2-002-006`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-006`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Explorer
- `SPECIALIZATION`: Campaign/Donation threshold and settlement ordering facilitation
- `PARTICIPANT_ID`: `P-S2-002-OIR-006-1` — ephemeral identity for this Run only
- `PARTICIPANT_PROFILE_ID`: `KC-EXPLORER`
- `PARTICIPANT_PROFILE_PATH`: `.harscode-spaces/participant-profiles/profiles.md#kc-explorer`
- `PINNED_PROFILE_REVISION`: `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`
- `SESSION_TRANSITION`: `FRESH` — a separate Campaign/Donation ordering decision surface after terminal O1 evidence and O7 Design review.
- `TARGET_BASE_REVISION`: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535`; current working-tree decisions and Run evidence are effective inputs. Re-read current files at dispatch.
- `WORKFLOW_REVISION`: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `RUNTIME_HARNESS`: `codex-cli`

## Why this Run exists

Approved `TP-S2-002-007` §13 keeps Campaign/Donation contract ordering open after Product policy O6 was settled. OIR `OIR-S2-002-005` completed an independent O1 evidence route and did not resolve or block this ordering surface. The current Authority Map names Anhar as Campaign and Donation delivery/domain owner for current Slice 2, so this bounded owner-facilitation Run can proceed while O1's cross-feature currency-standard authority sync remains scoped to amount representation/storage.

The settled O6 policy is: `max_amount` is a Campaign closure threshold, not a Donation hard cap; a Donation submitted while Campaign is eligible is accepted at its full amount and remains eligible to settle even if threshold crossing/closure occurs while it is pending; new submissions after Campaign is closed are rejected. The Run must define the contract-level ordering/state-consistency rule for threshold crossing, accepted pending Donations, concurrent settlement/close, and post-close rejection. Do not reopen the settled policy or expand into Slice 3 Campaign closure/public-result behavior.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` — Approved baseline, especially §§4 Q5, 5 D5, 8, 9, 13 Campaign/Donation ordering; SHA-256 `441577782f7600fd5507d228c4dd6491009c32d6e8b81607b28a38a5ebb1a1a3`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-001/resolution-brief.md` — settled O6 Product policy and historical gap pointers; SHA-256 `3d2b22f6fca9bd23024eca41cda79417bcc5cf929513223867d73ddcc6440782`.
- `.harscode-spaces/authority-map.md` — Anhar Solehudin owns Campaign and Donation delivery/domain for current Slice 2 only; SHA-256 `dd8cd32f65979f766034f6d086018e31ff4c6d584961ef1e296324c47c6eb06c`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-005/amount-contract-brief.md` — O1 remains partially resolved; currency-standard owner gap is scoped and must not be resolved in this Run; SHA-256 `cbca64e9d548422cd40a6717da51664406e43ff39393aa8e4de753f4b074afa9`.
- Current Product/MVP: `docs/product/mvp-scope.md` and `docs/product/mvp-delivery-slices.md` — relevant Donation acceptance and Campaign threshold policy. Current domain evidence: `docs/spec/4-campaign/invariants.md`, relevant Campaign closure/eligibility feature sections, `docs/spec/5-donation/invariants.md`, `docs/spec/5-donation/features/01-submit-donation-settlement.md`, and only relevant authored operations in `api/openapi/campaign.yaml` and `api/openapi/donation.yaml`. Historical detail is evidence, not authority.
- If inspecting backend implementation, first read `backend/AGENTS.md`; then inspect only relevant Campaign eligibility, settlement/funding, and close-ordering evidence. Do not write code or modify tests.
- `.harscode-spaces/participant-profiles/profiles.md#kc-explorer`; canonical `../harscode-workspace/workflow/1-exploration-kickoff-prompt.md`, `workflow/orchestrated-run-overlay.md`, and Pilot #2 Open-Item Resolution guidance.

## Task and completion condition

Conduct bounded evidence and owner-resolution work for Campaign/Donation ordering: identify the current contract/code ordering evidence; compare it to settled O6 policy; state the specific concurrency/threshold ordering gap and the minimum contract-level rule needed to preserve the policy; facilitate an explicit direction from Anhar in his current-Slice-2 Campaign/Donation delivery/domain authority. Distinguish delivery ordering detail from any proposed change to Product semantics. If current evidence supports one ordering without a material owner choice, explain the evidence and why. If there is a conflict requiring Product/MVP authority, stop that decision and route it rather than inferring an owner.

Use canonical Exploration Stage 1 plan announcement and Human confirmation before Stage 2. Write focused Stage 2 evidence under `RUN_PATH/evidence/`. At Stage 3, present the smallest decision-ready ordering question, evidence, and consequences to the named owner; record only explicit direction. After that gate, write `RUN_PATH/campaign-donation-ordering-brief.md` and terminal `RUN_PATH/handoff.md` with provenance, findings, decision/blocker, verification performed/not performed, remaining concerns, and next route.

Write only this Run's Stage 2 evidence, `campaign-donation-ordering-brief.md`, and `handoff.md`. Do not modify Product/MVP, specs, OpenAPI, code/tests, Approved Techplan, prior Run artifacts, or orchestration state. Do not establish new global currency policy; do not invent money rounding/tax scope; do not authorize protected Tier-0 writes, Build, tests, `CONTRACT_READY`, or Slice 3 behavior. Keep O1 shared-currency owner sync and O2/O3 email-retention conflict open and separate.

## Execution envelope and model route

- `PREAUTHORIZED`: read relevant repository/Harscode evidence and write only assigned Run artifacts.
- `ORCHESTRATOR_DECISION`: material expansion beyond threshold/settlement ordering, Role/Profile changes, or next workflow route.
- `HUMAN_REQUIRED`: canonical Exploration Stage 1/Stage 3 gates and any material current-Slice-2 Campaign/Donation owner decision. Product-semantic changes must route to the applicable Product/MVP authority.
- `SELECTED_MODEL`: `gpt-6-luna`; `REASONING_EFFORT`: `high`; no extra model approval under `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: least-cost registered repository-reasoning model; `high` supports cross-domain ordering and concurrency consequence analysis. Escalate only for demonstrated capability insufficiency.
- `PHASE_ROUTE`: Explorer Open-Item Resolution for the independent Campaign/Donation ordering item; canonical Exploration governs stages and Pilot #2 guidance governs facilitation/handoff.

## Human-assisted dispatch

1. Open a fresh Codex Participant Session in `/home/anhar-solehudin/kencleng-workspace/kencleng` with `gpt-6-luna` / `high`.
2. Paste: `Jalankan Run OIR-S2-002-006 sebagai Explorer untuk Campaign/Donation threshold dan settlement ordering. Ikuti invocation durable .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-006/invocation.md, pinned Profile KC-EXPLORER, canonical Exploration prompt, orchestrated-run overlay, dan Pilot #2 Open-Item Resolution guidance. Mulai dengan Stage 1 plan announcement dan berhenti untuk Human confirmation sesuai canonical prompt. Tulis Stage 2 evidence, campaign-donation-ordering-brief.md, dan terminal handoff.md hanya di RUN_PATH. Laporkan masalah material atau completion ke Orchestrator.`
3. At canonical Stage 1/Stage 3 gates, respond directly to Participant. Report only a material problem or terminal completion to Orchestrator; routine progress does not require monitoring.
