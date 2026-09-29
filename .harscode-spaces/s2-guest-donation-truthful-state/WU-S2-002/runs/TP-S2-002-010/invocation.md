# Run Invocation — `TP-S2-002-010`

Status: `READY_FOR_HUMAN_DISPATCH`  
Prepared by: Orchestration Operator  
Prepared: 2026-09-29  
Human-facing language: Bahasa Indonesia; retain canonical Harscode terms and code/API identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `TP-S2-002-010`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-010`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `PRIOR_ARTIFACTS`: current-effective Approved Techplan `TP-S2-002-007/techplan.md`; completed `OIR-S2-002-006` Stage 2 evidence, decision brief, and terminal handoff; settled O6 policy in `OIR-S2-002-001/resolution-brief.md`; current Authority Map and Product/MVP sources.
- `ROLE`: Planner
- `SPECIALIZATION`: Material Techplan amendment for Campaign/Donation ordering decision D1
- `PARTICIPANT_ID`: `P-S2-002-TP-010-1` — ephemeral identity for this Run only
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`
- `PARTICIPANT_PROFILE_PATH`: `.harscode-spaces/participant-profiles/profiles.md#kc-planner`
- `PINNED_PROFILE_REVISION`: `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`
- `SESSION_TRANSITION`: `FRESH` — new Planner occurrence after completed Explorer Run `OIR-S2-002-006`.
- `TARGET_BASE_REVISION`: `934b093cea30230743dee951ae1e600762019f30`; the OIR-006 artifacts are current working-tree inputs. Re-read relevant live files and verify all input pins at dispatch.
- `WORKFLOW_REVISION`: `c6398caf356f414bc0f4a2679aef9da92cca34be`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `RUNTIME_HARNESS`: `codex-cli`

## Why this Run exists

Approved Techplan `TP-S2-002-007` §13 item 8 leaves Campaign/Donation threshold and settlement ordering for the named current-Slice-2 owners. Completed `OIR-S2-002-006` records Anhar's explicit owner decision D1 and evidence that historical Donation settlement wording cannot express accepted-pending full settlement after Campaign close. The decision is within the already-approved O6 policy and current Techplan scope; it does not change Product/MVP authority.

The Techplan remains the execution spine. Canonical Techplan rules require resolved Open Items to remain recorded with the actual resolution. This material, concurrency-sensitive delivery detail must be reconciled into a new Techplan revision before a later Implementer executes the contract-source changes. D1 is already decided; do not ask Anhar to make it again.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` — current-effective Approved spine; SHA-256 `441577782f7600fd5507d228c4dd6491009c32d6e8b81607b28a38a5ebb1a1a3`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-006/campaign-donation-ordering-brief.md` — explicit owner D1 and contract reconciliation consequence; SHA-256 `583737b3c5f78a546e710dbf31e415db9f144f859700ab233d1737e880032e14`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-006/handoff.md` — completed Run boundary, remaining work, and recommended route; SHA-256 `d7f928dd94bac4fa960843a5ff9fbc972c5931f1d3e8836927e94a128d12f357`.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-006/evidence/stage-2-gap-analysis.md` — exact current contract gaps and live-code inspection evidence.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-001/resolution-brief.md` — settled O6 Product/MVP policy; D1 must preserve it.
- `.harscode-spaces/authority-map.md` — Anhar Solehudin owns Campaign and Donation delivery/domain decisions for current Slice 2 only; SHA-256 `dd8cd32f65979f766034f6d086018e31ff4c6d584961ef1e296324c47c6eb06c`.
- `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md`, and root `AGENTS.md` — current product boundaries and precedence.
- `.harscode-spaces/participant-profiles/profiles.md#kc-planner`; canonical `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`, `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`, and `workflow/orchestrated-run-overlay.md`.
- Current Pilot #2 candidate routing: `../harscode-workspace/orchestration/pilot-2-candidate/README.md` and current cross-cutting owner `orchestrator-operating-model.md`.

## Task and completion condition

Amend Approved `TP-S2-002-007/techplan.md` into `RUN_PATH/techplan.md` by recording D1 as the current-Slice-2 Campaign/Donation delivery decision, adding its consequences to the relevant execution spine sections, and moving the existing Campaign/Donation ordering item from Active Open Items to Resolved with owner/date/provenance. Preserve prior approval history by writing a new Run artifact; do not overwrite `TP-S2-002-007`.

Carry D1 exactly: submission eligibility is ordered atomically against Campaign close; a Donation accepted while eligible remains settleable at full amount after close; successful Donation settlement and its full funding increment commit together exactly once; later settlement cannot reopen Campaign or change the winning close reason; funding may rise past `max_amount`. Preserve O6 and the Slice 3 boundary. Do not select a locking/isolation mechanism or invent a new endpoint.

Preserve O1 and O2/O3 as separate unresolved scopes, along with all other open owner, security, design, API, and verification items. Retain relevant concurrency/Test Focus traceability to OIR-006 evidence. Classify the Techplan amendment's materiality and state the next review/Human gate. D1 does not need a second owner decision, but a material revised Techplan still follows the applicable fresh independent review and Human approval gates.

Write only `RUN_PATH/techplan.md` and `RUN_PATH/launch-record.md`. Do not generate `report-techplan.md` before review/resolution convergence. Do not edit Product/MVP, Design, Authority Map, specs, OpenAPI, code/tests, prior Run artifacts, or other orchestration files. Do not start Build or claim `CONTRACT_READY`.

## Execution envelope and model route

- `PREAUTHORIZED`: read relevant project/Harscode evidence and write only this Run's Techplan and launch record.
- `ORCHESTRATOR_DECISION`: scope expansion, resolving unrelated Open Items, or changing the next workflow route.
- `HUMAN_REQUIRED`: any proposed change to Product/MVP meaning or material current-Slice-2 owner direction; route it rather than infer. No repeat decision is required for explicit D1.
- `SELECTED_MODEL`: `gpt-6-luna`; `REASONING_EFFORT`: `high`.
- `PHASE_ROUTE`: Planner Techplan amendment. Because the updated execution spine adds material concurrency-sensitive interface/domain detail, prepare for fresh independent Techplan review. Planner generates the Human review report only after review/resolution converges.

## Human-assisted dispatch

1. Open a fresh Codex Participant Session in `/home/anhar-solehudin/kencleng-workspace/kencleng` with `gpt-6-luna` / `high`.
2. Paste: `Jalankan Planner amendment Run TP-S2-002-010 mengikuti invocation durable .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-010/invocation.md, current Approved Techplan TP-S2-002-007, canonical Techplan synthesis prompt/template/rules/guardrails, orchestrated-run overlay, dan current Harscode guidance dari ../harscode-workspace. Rekonsiliasikan hanya owner decision D1 dari OIR-S2-002-006 ke Techplan baru di RUN_PATH; pindahkan Open Item Campaign/Donation ordering ke Resolved dengan provenance; pertahankan seluruh scope dan Open Items lain. Jangan minta Human mengulang D1, jangan ubah Product/MVP/spec/API/code/test/orchestration projections, jangan buat report sebelum independent review converges, jangan mulai Build atau klaim CONTRACT_READY. Tulis hanya techplan.md dan launch-record.md Run ini, lalu berhenti.`
3. Laporkan hanya masalah material atau completion kepada Orchestrator; interaksi routine tidak perlu dipantau.
