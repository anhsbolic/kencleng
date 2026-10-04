# Run Invocation — `TP-S2-003-009`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-04 after Anhar accepted the exact five WU-S2-007 API/source-counterpart revisions. WU-S2-007 is complete and its exact handoff is available to this Run. This is a fresh Planner occurrence to refresh the existing stable WU003 candidate; it does not approve the candidate or authorize Build.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `TP-S2-003-009`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-009`
- `ARTIFACT_TARGET`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/techplan.candidate.md` — continue the one stable Draft/In Review successor; preserve Approved `TP-S2-003-006` unchanged.
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Refresh the current WU-S2-003 whole-Techplan candidate against the accepted WU-S2-007 source/counterpart handoff and current effective authority
- `PARTICIPANT_ID`: `P-S2-003-TP-009-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Planner occurrence after TP-S2-003-008; do not reuse its Participant Session or the Review Session.
- `TARGET_REVISION`: Kencleng HEAD `4e71d3697a479d92b2cad5a79ef5d710f4e8077c` plus current working tree. Verify the current candidate and exact source/handoff hashes before editing; stop and report material drift.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; current-effective workflow guidance must be re-read at dispatch.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by the Human-owned local registry.
- `MODEL_REGISTRY_SOURCE`: `.harscode-spaces/.local-config.yaml`, SHA-256 `ddffdeb4ce8edec181ca36bf32bedf19c513a76e359bc2642985eda6b32f400d`; Luna supports reasoning/repository work at `high` and is not approval-gated.
- `MODEL_ROUTING_RATIONALE`: This Run refreshes one whole-plan material successor using settled authority/source evidence and a prior independent design review; it must preserve the plan spine and accurately route unresolved authority/design gates. `gpt-6-luna` / `high` is supported by the local registry and TP-S2-003-008 precedent. Sol is approval-gated and is not needed for a planning synthesis; higher model capability would not resolve any Human-owned authority gate.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Canonical Techplan synthesis for a material successor to an Approved plan. Read all durable WU003 Exploration artifacts once as required for a fresh Planner Session, classify and reconcile current authority/evidence under the canonical prompt, and update only the stable candidate plus this Run's structured handoff. Do not produce an approval report during review/resolution churn.
- `ARTIFACT_TARGETS`: stable `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/techplan.candidate.md`; Run-local `runs/TP-S2-003-009/handoff.md`. Do not create a versioned Techplan copy or overwrite the Approved predecessor.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Current Approved predecessor: `runs/TP-S2-003-006/techplan.md`, SHA-256 `b71951c86b2f74f6c9f1a164bcf43447c58ee655874deaf09275e6ff3b2316f9`; handoff SHA-256 `e62e4df0cfc9095d06ffe6e987afdca8f6ee369a2f7a16d7d1306ddaa2653bdd`.
- Existing material-successor candidate: `techplan.candidate.md`, SHA-256 `353302264fb5458d472f6e508c83c21325477f644c8ff6990c9b3521755a8754`; previous Planner handoff `runs/TP-S2-003-008/handoff.md`, SHA-256 `926bde96d2e62087e59d52e9c18ee24a93dd9d948893abcf1334dfada15ca229`.
- WU-S2-007 exact accepted handoff: `../WU-S2-007/handoff-to-WU-S2-003.md`, SHA-256 `cf468693dc384ccad5a4156c43d433a3ff725a6169869c3f171aeff0a27325eb`. It records Anhar's exact-byte acceptance, five accepted hashes, independent Review/Testing evidence, and downstream scope/boundaries. Read it along with the exact current authored Donation OpenAPI source and split common reference below.
- Prior independent WU003 migration-design review: `runs/RV-S2-003-005/review-findings-1.md`, SHA-256 `fca27d88610017db317f679a68c97af7b49a616d3b7b69d6fa06b45fef359c7f`; verdict `Request changes`. Its proposal target is `runs/BLD-S2-003-002/report.md`, SHA-256 `dc8f5475b567f692760bb61e7287af7d93dfb4c4a77f7ea8ae9cfd2577a8a4af`. This remains evidence; it is not approved schema design.
- Current WU003 manifest and project coordination surfaces, including the preserved D1 file authorization, Open Item 7, execution status, and deferred evidence.
- Current source authorities at preparation (recheck at dispatch): root `AGENTS.md` SHA-256 `4ec81e122cfeb77b2d13e3b90cc7f4a73e829f97d0627d889793aca847a87716`; `backend/AGENTS.md` `6a2a9c84b007572155be0690fba3c04e1be5e522998a3aff3cfd1f01f96fef8e`; `docs/kencleng-agentic-workflow.md` `726466e5f78ba6cf9127b7942386b9d9e237b638d768929fb6b80ea7a5b15ff9`; Product README `4bc4ec89425fb27ded281259a1cd918b2cd7c207c7341e4f8f6f8f0285e6780e`, overview `11556a2e45c679c50a090662e9ba64df2ab7785ed02bd0d72ce27a2073b5b674`, MVP scope `ba2972bc8f91d092e477df170d987b1d124964d9cc36c025d2a8da3ed12709af`, delivery slices `4c69a030e7fedc9c62bf30f85c00e81f9806c45b2ed5471c1d5126762be8091f`.
- Current domain/contract authorities at preparation (recheck at dispatch): Campaign invariants `docs/spec/4-campaign/invariants.md` SHA-256 `42edf2ab5713b8905e9a9a8e554da2a678fbf70fcbabe95670113cbcdf8b1423`; Donation invariants `docs/spec/5-donation/invariants.md` `4bf7ba48f4980f635ce76a4eb3440be5cf12666fdd99e06dcc8ac3b78fbde39a`; Donation feature `docs/spec/5-donation/features/01-submit-donation-settlement.md` `b3f0f99c420c807bb14cccf75a4adcaf43dfdadccae6b727a1c460ef4ed6a7e9`; authored API `api/openapi/donation.yaml` `9f7c31065c3c7ffa77491541102afc0ce84085f50a1e2aec317336c9656d5c1d`; common API `api/openapi/common.yaml` `46bd8fba98d7befe65e87b17f431b0f0a2fa815589db28eb6ded6b00bd9418f8`; `api/README.md` `0ec88871d43b32b2b06a82ba345f694b079bfefaef3c2974333b43fad16b434f`.
- Enumerate/read every durable Exploration file under `runs/EXP-S2-003-001/evidence/` once for this fresh Techplan occurrence as required by the canonical prompt. Also use the prior Approved predecessor, current candidate, WU-S2-007 handoff, Review receipt, live code/schema/migration anchors, and routed best-practice authorities where relevant.
- Canonical Harscode: `workflow/2-1-techplan-synthesis-prompt.md` SHA-256 `1ed5bdc6e8bb70beac4dc61f4328c5cc0a8b5a70cc82ae7dba4ce22a9c588613`; `workflow/2-techplan/template.md` `aa2bb5acdf7b2d4d79975fc5bbcd92b47094223361485382bce0520e3fca0ccd`; `rules.md` `ab6939e9b4d104bd0db2669945ee5a32e560ebe006356363b3ed8da969d56c2e`; `guardrails.md` `74160046a553d2b53a84320c1397b5fdee44fa9bb1e3c3fd9fb532c038e2a6a4`; `workflow/AGENTS.md`; and `workflow/orchestrated-run-overlay.md` SHA-256 `f806673aa5e7d1d82dae9f433ec54c3562e15ce778b709a141c57d3d22dfaf87`. Re-read current-effective versions at dispatch.

## Task and boundaries

Prepare or continue the one stable `techplan.candidate.md` as a material successor to Approved TP-S2-003-006. Reconcile the whole WU003 execution plan with the current Product/MVP and domain/API authorities, all required durable Exploration evidence, current candidate, prior Review/Build evidence, and WU-S2-007 accepted source/counterpart handoff. Preserve product-first precedence, prior resolved decisions, scoped Open Item 7, exact Human permissions, and the current internal/repository consumer posture. Surface any genuine authority conflict or material decision that remains unresolved; do not invent its answer or let historical implementation override current authority.

Use the canonical Techplan synthesis workflow and template. Do not steer the synthesis toward a predetermined technical conclusion. Keep the approved predecessor unchanged and the candidate Draft / In Review. Recommend independent Techplan review and decomposition according to the canonical criteria. Generate no `report-techplan.md` while the candidate is in planning-review/resolution churn; do not dispatch review or Build.

## Execution envelope

- `PREAUTHORIZED`: read current authorities/live anchors and all WU003 Exploration evidence as required; update only the stable candidate and write this Run's structured handoff.
- `HUMAN_REQUIRED`: any new material Product/domain/API/security decision, candidate approval/promotion, protected write, migration application, risk acceptance, or WU003 completion.
- Tests, validators, generators, API/spec/code/test/migration/database changes, and downstream Run dispatch are out of scope.
- If any current authority/source hash has materially changed, the stable candidate lifecycle is inconsistent, or a material blocker prevents safe synthesis, stop and report the gap without changing higher-authority sources.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Planner / `KC-PLANNER`, `gpt-6-luna` / `high`.

Kickoff: `Jalankan Planner Run TP-S2-003-009 memakai canonical ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md dan orchestrated-run overlay, dengan inputs/output sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-009/invocation.md. Perlakukan techplan.candidate.md sebagai satu-satunya stable Draft successor dan jangan ubah Approved TP-S2-003-006. Baca semua durable WU003 Exploration evidence dan current authority/source/handoff sesuai canonical prompt. Berhenti setelah candidate synthesis/self-check dan satu structured handoff; jangan generate report approval selama review/resolution churn, jangan dispatch Review atau Build.`
