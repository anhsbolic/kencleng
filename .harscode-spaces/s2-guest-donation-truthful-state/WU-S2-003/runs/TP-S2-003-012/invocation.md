# Run Invocation — `TP-S2-003-012`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-04 after Anhar resolved RV-S2-003-008 F-01 through the API/Security-PII owner question. This is a bounded material resolution to the same WU003 Techplan; it is not Build authorization and does not itself satisfy the independent migration-design Review gate.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `TP-S2-003-012`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-012`
- `ARTIFACT_TARGET`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/techplan.candidate.md`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Bounded Techplan resolution for migration-design finding RV8 F-01
- `PARTICIPANT_ID`: `P-S2-003-TP-012-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`
- `SESSION_TRANSITION`: `FRESH` — new Planner Run/Participant and fresh Session/context; RV-S2-003-008 is completed and its Reviewer context must not be reused.
- `TARGET_REVISION`: Kencleng HEAD `4e71d3697a479d92b2cad5a79ef5d710f4e8077c` plus current working tree; re-ground and verify all exact inputs before editing. Current Approved candidate pre-run SHA-256 `55eb0c94bb922d3311279c462122cb57f9f2181226f33092fff9d68bbd2c41b5`.
- `WORKFLOW_REVISION`: current-effective Harscode guidance at dispatch. Canonical Techplan entrypoint `workflow/2-1-techplan-synthesis-prompt.md` SHA-256 `1ed5bdc6e8bb70beac4dc61f4328c5cc0a8b5a70cc82ae7dba4ce22a9c588613`; orchestrated-run overlay SHA-256 `f806673aa5e7d1d82dae9f433ec54c3562e15ce778b709a141c57d3d22dfaf87`; Techplan template SHA-256 `aa2bb5acdf7b2d4d79975fc5bbcd92b47094223361485382bce0520e3fca0ccd`; guardrails SHA-256 `74160046a553d2b53a84320c1397b5fdee44fa9bb1e3c3fd9fb532c038e2a6a4` (re-read current guidance at dispatch; do not treat hashes as a request to use stale workflow guidance).
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned local registry `.harscode-spaces/.local-config.yaml`, SHA-256 `ddffdeb4ce8edec181ca36bf32bedf19c513a76e359bc2642985eda6b32f400d`.
- `MODEL_ROUTING_RATIONALE`: This is a bounded but material security/privacy and persistence-design correction to an Approved backend Techplan. `gpt-6-luna` / `high` matches the current WU003 Planner route and provides sufficient reasoning for a precise schema/interface/test plan and careful separation of selected behavior from unresolved O3/O4 controls.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Fresh Planner resolution using the canonical Techplan synthesis entrypoint, current workflow guidance, and orchestrated-run overlay. Revise the single stable candidate successor and produce one structured Phase handoff. Do not generate the Human-facing `report-techplan.md` while review/resolution churn remains.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Exact pre-run Approved candidate: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/techplan.candidate.md`, SHA-256 `55eb0c94bb922d3311279c462122cb57f9f2181226f33092fff9d68bbd2c41b5`. This is the exact plan state RV-S2-003-008 reviewed; preserve it as the predecessor in provenance and do not edit a second plan copy. After material revision, keep the same stable `techplan.candidate.md` successor path and leave it `Draft / In Review` pending applicable Review and Human approval.
- Blocking finding: `runs/RV-S2-003-008/review-findings-1.md`, SHA-256 `60eabbd9ccaadd975e6f8b25a897e313c215e938f2a30f337d3c6a9bfd796ae4`; its exact target was candidate SHA-256 `55eb0c94bb922d3311279c462122cb57f9f2181226f33092fff9d68bbd2c41b5`. F-01 concerns returning the required `status_token` on same-key/equivalent-intent replay while never persisting bearer material and retaining a one-way verifier. RV8 assessed RV5 F-03/F-04/F-06 resolved at design level; this Run must not reopen those findings without new evidence.
- Human authority decision, recorded in parent `events.md` on 2026-10-04: for every exact same-key and equivalent-canonical-request POST retry, return the original Donation and issue a fresh status bearer for that same Donation. Each credential has an independent hard 24-hour expiry from its issuance. Persist only its one-way verifier; never persist the bearer. Multiple credentials for one Donation may be simultaneously valid until their individual expiries. Treat this choice as settled input, not a question to reopen or alter.
- Existing approved credential direction in the candidate and current sources: fragment-carried status URL and frontend handoff/URL cleanup; status-only result; one-way HMAC verifier; hard 24-hour expiry from issuance; no bearer/PII/sensitive request body in ordinary logs. Numeric strength/generation recipe, key purpose/provisioning/lifecycle, comparison controls, expiry-enforcement mechanism, abuse/exposure controls, and bounded idempotency-record lifetime remain unselected/unproven O4/O3 gates.
- Current source anchors to re-ground read-only: `api/openapi/donation.yaml` Donation schema / POST response / idempotency description; Donation Feature 01 retry behavior; Donation Feature 02 status access; `docs/spec/5-donation/invariants.md` INV-donation-01/02/05/08/10. Determine whether their exact wording adequately carries the Human-selected replay issuance rule. If a source/counterpart gap remains, identify the precise owning source, proposed bounded reconciliation and acceptance gate in the candidate/handoff; do not edit Product, spec, API, generated contract, fixtures or consumer sources in this Planner Run.
- Current WU003 state and provenance: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/manifest.md`; parent `events.md`, `work-graph.md`, `control-surface.md`, `outcome.md`; `docs/project/kencleng-development-tracker.md`. WU-S2-007 exact source/counterpart dependency remains satisfied. Preserve Open Item 7 as a separate scoped gate.
- Existing design context: approved predecessor `runs/TP-S2-003-006/techplan.md`; RV-S2-003-005 findings; RV8 findings above; current candidate §§8–10, §12 validation/test ownership, and §13 lifecycle/open items. Preserve RV8's positive design dispositions for F-03/F-04/F-06 unless new evidence materially changes them. Keep BLD-S2-003-002 `STALLED` and its prior proposal non-authoritative.
- The Human credential choice closes the owner-decision blocker only. It is not a positive migration-design verdict, source acceptance, Build authorization, protected-path authorization, migration approval/application, runtime/security evidence, residual-risk acceptance, or WU milestone.

## Task and completion condition

Resolve only RV-S2-003-008 F-01 in the stable Techplan candidate from the exact pinned pre-run hash and the Human decision above. Reconcile relevant candidate requirements, schema/data shape, replay flow, credential verifier representation, expiry behavior, migration/write-path implications, test/evidence coverage and gate routing so they consistently state:

1. the idempotent replay returns the original Donation and a newly issued `status_token` credential for that Donation;
2. every issued credential has its own hard 24-hour expiry measured from issuance; replay does not invalidate existing credentials unless an accepted source explicitly requires it (do not invent invalidation);
3. only one-way verifier material is persisted; bearer material is never persisted or logged;
4. multiple active verifier/credential records for one Donation are allowed as a consequence of the Human choice, with exact schema mechanics selected by the Planner only where supported by current architecture and sufficient for the migration-design Review;
5. the unresolved numeric/generation, key lifecycle, comparison, expiry enforcement, abuse/exposure and idempotency-retention controls remain explicit O4/O3 questions/evidence gates and are not silently answered by the 24-hour TTL choice.

Assess whether the current accepted Feature/API/invariant wording is sufficient for this concrete behavior. If not, record exact source/counterpart authority and the required reconciliation/acceptance as an explicit predecessor to relying on the behavior; do not edit those sources in this Run. Do not weaken the required POST `status_token` response or same-key original-Donation semantics. Do not guess any other API/product decision.

Declare materiality honestly. The selected behavior adds replay issuance and permits multiple active credentials, changing interface/persistence/security semantics. Recommend a fresh independent Techplan Review and, after Human approval/source reconciliation as applicable, a fresh independent migration-design Review of the exact resulting design. Do not claim either review has happened or the minimum schema-design gate is positive.

At completion, write the candidate using the current canonical Techplan template/guidance and one `handoff.md` in this Run with exactly one orchestrated `## Phase handoff`. Verify input hashes and the bounded semantic delta; make no report-techplan until applicable review/resolution converges.

## Execution envelope

- Re-read the current canonical Techplan prompt, `workflow/AGENTS.md`, Techplan guardrails/template and `orchestrated-run-overlay.md`; verify the candidate, RV8 findings, Human decision Event, current owning sources and current WU state before editing. If exact pins or authority have materially changed, stop and explain in the handoff instead of applying a stale assignment.
- Authorized writes: the stable `techplan.candidate.md` successor and this Run's `handoff.md` / `launch-record.md` only. The Invocation is immutable after dispatch. Preserve the exact pre-run Approved bytes through the recorded hash and existing durable history; do not overwrite `runs/TP-S2-003-006/techplan.md` or create another versioned candidate.
- No writes to Product, domain specs, API/OpenAPI/generated bundles/types/fixtures/consumers, code, tests, migrations, database, manifests, tracker, parent orchestration state or any other Run. No tests, validators, generators, SQL, migration/database/runtime/browser/security actions. Do not dispatch Reviewer, Builder or another Participant.
- Do not declare the schema-design gate positive, authorize a Build, claim verified O3/O4 controls, accept residual risk, infer Human-paired Tier-0 completion, or claim WU/backend/runtime/delivery completion.

## Human-assisted dispatch

- Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`
- Fresh Planner / `KC-PLANNER` Session; configured `gpt-6-luna` / `high`.
- Canonical kickoff: `Jalankan fresh Planner resolution Run TP-S2-003-012 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-012/invocation.md dan current canonical ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md, Techplan guardrails/template, workflow/AGENTS.md serta orchestrated-run overlay. Re-ground dan verifikasi seluruh pin dan authority. Revisi hanya stable techplan.candidate.md untuk RV-S2-003-008 F-01 memakai keputusan Human yang tercatat: exact same-key/equivalent-request retry mengembalikan Donation yang sama dan menerbitkan credential status baru; setiap credential memiliki TTL keras 24 jam sejak issuance, simpan verifier satu arah saja, bearer tidak disimpan, beberapa credential boleh aktif bersamaan. Pertahankan unresolved O3/O4 controls secara eksplisit. Nilai apakah exact API/spec wording memerlukan source/counterpart reconciliation; catat owner/scope/acceptance gate tanpa mengedit sources. Tandai perubahan material dan rekomendasikan fresh independent Techplan Review lalu fresh migration-design Review; jangan mengklaim gate positif. Candidate tetap Draft / In Review; jangan buat report-techplan saat review/resolution belum konvergen. Tulis satu handoff terstruktur dan launch record lalu berhenti. Jangan edit file lain, jalankan tests/validators/generators/database/runtime, authorize Build, atau dispatch Participant berikutnya.`

This Invocation prepares a Run only. Anhar dispatches the fresh Planner; no Participant has been dispatched by the Orchestrator.
