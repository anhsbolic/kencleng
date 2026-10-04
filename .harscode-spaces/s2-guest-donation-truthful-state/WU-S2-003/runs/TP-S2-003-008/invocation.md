# Run Invocation — `TP-S2-003-008`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-04 by Orchestration Operator after Anhar explicitly resolved the Funding-unavailable and Donation amount O1 decisions. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `TP-S2-003-008`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-008`
- `ARTIFACT_TARGET`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/techplan.candidate.md` — one stable material-successor candidate to current Approved TP-S2-003-006; create or continue this candidate without changing its Approved predecessor.
- `TASK_PATH` / `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Reconcile approved WU003 Techplan and minimum Donation/D1 schema design after independent design-review findings
- `PARTICIPANT_ID`: `P-S2-003-TP-008-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; profile SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`; registry SHA-256 `804856c5defad721678be4bac145fc2a1070307cccd720326cc598b8159d7abe`.
- `SESSION_TRANSITION`: `FRESH` — new Planner occurrence after completed independent design Review RV-S2-003-005; do not reuse Reviewer or earlier Planner Session/context.
- `TARGET_REVISION`: Kencleng HEAD `4e71d3697a479d92b2cad5a79ef5d710f4e8077c` plus current working tree. Reconfirm the listed current-effective hashes at dispatch; stop for Orchestrator reconciliation if any differ.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; ordinary current-effective guidance.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: The task is a bounded revision of a current plan against six independent findings and two now-settled owner decisions, with exact current specs and schema anchors available. The Human registry gives `gpt-6-luna` reasoning/repository-work capability and supports `high`; prior WU003 planning used this configured level. No stronger model is needed to resolve authority gaps, and model escalation would not substitute for them.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Material successor to Approved TP-S2-003-006. Canonical Techplan synthesis rules require the current Approved predecessor to remain unchanged while this single stable candidate is prepared. This Run does not approve the candidate or satisfy the subsequent independent Techplan and migration-design review gates.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Current Approved predecessor: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-006/techplan.md`, SHA-256 `b71951c86b2f74f6c9f1a164bcf43447c58ee655874deaf09275e6ff3b2316f9`; handoff `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-006/handoff.md`, SHA-256 `e62e4df0cfc9095d06ffe6e987afdca8f6ee369a2f7a16d7d1306ddaa2653bdd`.
- Exact independent design-review receipt: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-005/review-findings-1.md`, SHA-256 `fca27d88610017db317f679a68c97af7b49a616d3b7b69d6fa06b45fef359c7f`; verdict `Request changes`, findings F-01–F-06. Target proposal `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/BLD-S2-003-002/report.md`, SHA-256 `dc8f5475b567f692760bb61e7287af7d93dfb4c4a77f7ea8ae9cfd2577a8a4af`.
- Human-resolved source authority: `docs/spec/4-campaign/invariants.md`, SHA-256 `42edf2ab5713b8905e9a9a8e554da2a678fbf70fcbabe95670113cbcdf8b1423`; `docs/spec/5-donation/invariants.md`, SHA-256 `df3fb5dafd901cf91c8089f40102a39ab105ba2f64d3a937cab41ca1a1156664`; `docs/spec/5-donation/features/01-submit-donation-settlement.md`, SHA-256 `22a933882819acb29665fb97074ee93ddd30e5c1bb9b2fcb570ffbe715a73262`. On 2026-10-04 Anhar resolved: (a) Donation admission fails closed when authoritative settled Funding is unavailable, treats no missing value as zero, and resumes only after authoritative Funding is restored; (b) Slice-2 Donation is IDR-only, whole Rupiah Rp5.000–Rp1.000.000.000 inclusive, persisted as an exact integer Rupiah amount with zero fractional digits. This feature decision does not select a project-wide database type/scale or another feature's money parameters.
- Current WU003 manifest and project tracker for execution status, prior evidence, Open Item 7 scope, exact D1 authorization, and deferred evidence.
- Required Planning corpus: enumerate and read each durable Exploration artifact under `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/EXP-S2-003-001/evidence/` once as required for a fresh Techplan session; use the prior approved Techplan and RV-005 receipt above as the current change evidence. Relevant review scope TST-S2-003-001 and RV-S2-003-004 is limited to the earlier public cap-projection slice, not Donation/D1 schema.
- Current accepted Campaign/Donation specs and authored OpenAPI sources; `api/openapi/donation.yaml` plus referenced `common.yaml`; root `AGENTS.md`, `backend/AGENTS.md`, relevant sections of `docs/kencleng-agentic-workflow.md`, `docs/project/kencleng-monetary-data-standard.md`, current migration/domain anchors, and matching PostgreSQL best practices per the canonical prompt.
- Canonical Harscode `workflow/2-1-techplan-synthesis-prompt.md`, `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`, `workflow/orchestrated-run-overlay.md`, `workflow/AGENTS.md`, `orchestration/AGENTS.md`, and `orchestration/run-contract.md`.

## Task and completion

Prepare or continue exactly one stable `techplan.candidate.md` as the material successor to Approved TP-S2-003-006. Reconcile the full WU-S2-003 execution plan against current authorities, the Human decisions above, and RV-S2-003-005. Preserve the existing approved D1 ordering/mechanism and all settled product/API/security decisions. Do not change the Approved predecessor in place.

Independently assess and disposition F-01–F-06 in the candidate:

- Incorporate the settled unavailable-Funding admission policy (F-01) and Donation O1 amount/storage bounds (F-05) from the exact current specs.
- Reconcile minimum key uniqueness scope and accepted-field payload equivalence for request idempotency (F-02) against INV-donation-10 and the accepted request contract. Preserve established guest-PII encryption/HMAC handling. Do not invent email retention, deletion, delivery or other O3 controls; identify a new owner decision if the settled contract cannot be met within current authority.
- Make the migration sequence, existing-row/constraint validation posture, locking assessment, and down-migration behavior explicit (F-03). Distinguish pre-write rollback from the data-loss consequences of post-write down. Do not claim execution or reversibility evidence.
- Name the pending-reservation and idempotency lookup predicates, terminal-state/integrity requirements, FK deletion policy and matching indexes (F-04/F-06). Preserve Campaign→Donation lock order; do not add a stored counter without a demonstrated need. Preserve Donation history on Campaign replacement/deletion; handle legacy closed Campaigns with unknown reasons without fabricating history, while requiring accepted close reasons for new close winners.
- Reconcile the new fail-closed rule with current POST/OpenAPI problem semantics. Do not invent a status/problem contract or alter accepted API sources in this Run; surface the smallest owner/API reconciliation if the current contract has no truthful representation.

Treat Review suggestions as evidence/options, not new authority. Preserve Open Item 7 exactly as a scoped gate on Organization truth updates and affected Campaign create/PATCH handlers; do not widen it or let it block unrelated Donation work. Preserve exact accepted seven-source revisions and all protected-path/Human-pairing gates. Keep migrations unapplied; no migration SQL, source, API, spec, test or database writes are allowed in this Run.

The candidate must preserve full rule-to-verification and Exploration-risk-to-Test-Focus traceability. Assess whether current WU003 scope needs further decomposition; do not invoke decomposition here. Keep the candidate Draft / In Review. Produce only the stable `ARTIFACT_TARGET` and `RUN_PATH/handoff.md`. Do not generate `report-techplan.md` until an independent Techplan review/resolution path converges and the exact candidate is ready for its Human approval gate. Recommend the next independent Techplan review because cross-boundary fidelity was material enough to require this design review. Stop and report any contradiction or material unresolved authority question.

## Execution envelope

- `PREAUTHORIZED`: read current authorities and live migration/domain anchors; create or revise only the stable Techplan candidate; write this Run's handoff.
- `HUMAN_REQUIRED`: any change to Product/domain/API authority beyond the explicit 2026-10-04 decisions; any new privacy/security policy; any approval/promotion, Tier-0 write, migration application, or risk acceptance.
- Tests, validators, generators, SQL/migration commands, DB/runtime actions, source/API/spec changes, Build, and downstream dispatch are out of scope.
- Stop if an input hash changed since preparation, if candidate state conflicts with its expected lifecycle, or if a material decision is missing; report it to the Orchestrator without widening scope.

## Human-assisted dispatch

- Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`
- Fresh Planner / `KC-PLANNER` Session; configured `gpt-6-luna` / `high`.
- Canonical kickoff: `Jalankan Planner Run TP-S2-003-008 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-008/invocation.md dan canonical ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md dengan orchestrated-run overlay. Berhenti setelah stable Techplan candidate dan structured phase handoff; jangan dispatch Review atau Build.`
