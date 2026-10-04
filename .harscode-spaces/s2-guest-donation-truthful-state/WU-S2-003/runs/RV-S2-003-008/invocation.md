# Run Invocation — `RV-S2-003-008`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-04 after Anhar approved the exact WU-S2-003 Techplan candidate presented in TP-S2-003-011. This Run is the fresh independent migration-design review required by the Approved plan before any Donation/D1 schema Build. It does not authorize code or migration writes.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `RV-S2-003-008`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-008`
- `ARTIFACT_TARGET`: `none` — write only `RUN_PATH/review-findings-1.md` and `RUN_PATH/launch-record.md`.
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent pre-implementation review of the revised minimum Donation/D1 schema and migration design in the Approved Techplan
- `PARTICIPANT_ID`: `P-S2-003-RV-008-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`; current local model registry `.harscode-spaces/.local-config.yaml`, SHA-256 `ddffdeb4ce8edec181ca36bf32bedf19c513a76e359bc2642985eda6b32f400d`.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer session after TP-S2-003-011 and the prior non-positive design review RV-S2-003-005. Do not reuse Planner or prior Reviewer context as independent review evidence.
- `TARGET_REVISION`: Kencleng HEAD `4e71d3697a479d92b2cad5a79ef5d710f4e8077c` plus current working tree; re-ground and verify all pinned inputs before reviewing. Current Approved candidate SHA-256 `55eb0c94bb922d3311279c462122cb57f9f2181226f33092fff9d68bbd2c41b5`.
- `WORKFLOW_REVISION`: current-effective Harscode at dispatch. Canonical review prompt `workflow/4-code-review-prompt.md`, preparation SHA-256 `d152c1419ea8b0e67a08d1ec2cec7c0235cc15b0fccbd032722aea1ed2332f8a`; guidelines SHA-256 `513f145a5bc1980589a4967f6a0894a3f59f089ade4fa199ba252b2604a735c0`; checklist SHA-256 `5621854c75f8834ff9c32a8a67a326ab03388e11cceac95f0f2c6bdb0e03e079`; overlay SHA-256 `f806673aa5e7d1d82dae9f433ec54c3562e15ce778b709a141c57d3d22dfaf87`; run contract SHA-256 `4610c20ce2670913870718f9886e22c3c00bdf9da5290406a4a4450a818069aa`.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned local model registry; verify current model route before dispatch.
- `MODEL_ROUTING_RATIONALE`: This is a bounded adversarial review of a written PostgreSQL schema/migration design against settled monetary, D1, retry, lifecycle and migration-safety requirements plus a small set of current schema anchors. `gpt-6-luna` / `high` matches the prior independent design review route and is sufficient for this bounded artifact review; it does not choose unresolved O3/O4/O5 or Open Item 7 policy.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Fresh independent four-pass Review using `workflow/4-code-review-prompt.md` and the current Review guidelines/checklist/overlay. This is a pre-implementation design review of an exact approved Techplan slice, not a code diff review or whole-Techplan re-review. Adapt the four passes to the pinned design target; do not edit the Techplan, proposal, source, tests, SQL, database, or other artifacts.

## Exact review target and current-effective inputs

- Primary design target: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/techplan.candidate.md`, current Approved SHA-256 `55eb0c94bb922d3311279c462122cb57f9f2181226f33092fff9d68bbd2c41b5`. Review only the minimum Donation/D1 schema and migration design represented by §9 steps 2a–2b and 8, §10 migration/D1 rows (including successor after `000012`, Donation identity/equivalence, close reason, FK/index/constraint/down posture), and §12 R14/R18/R19/R20. Use §13 Item 6 for the gate boundary. Verify exact hash and status before reviewing; if changed, stop and ask Orchestrator to reconcile.
- Human approval pair: pre-approval candidate SHA-256 `0d5a503a7a9ae892f369e1e14aae41f8c3b5a2549cc8c687fae463d0d25e2df7`; report-only TP-S2-003-011 report SHA-256 `2c8d6b0996d17596f0ce4feabf49969cef745f936060cb7a23e9910565013235`; launch-record SHA-256 `569b51645611097e5be52138bff5c6567081c7d9fedc9105e8b5aa2396ed81d7`. Anhar approved this exact candidate/report pair. Current Approved candidate differs only by permitted lifecycle/provenance metadata; verify the approved hash above.
- Prior design review: `runs/RV-S2-003-005/review-findings-1.md`, SHA-256 `fca27d88610017db317f679a68c97af7b49a616d3b7b69d6fa06b45fef359c7f`, reviewed the BLD-S2-003-002 proposal hash `dc8f5475b567f692760bb61e7287af7d93dfb4c4a77f7ea8ae9cfd2577a8a4af`; verdict `Request changes`, with F-01–F-06. Current accepted authority resolves F-01/F-05; the Approved candidate contains revised design responses to F-02/F-03/F-04/F-06. Independently assess those responses; neither the prior findings nor this Invocation predetermine the new verdict.
- Existing schema anchors to inspect read-only: `backend/migrations/000011_create_public_campaigns.up.sql` SHA-256 `440d1919fcca537f14113137067fb7079c374ac3ad8a60127dcddd46e9be53c4`; `backend/migrations/000012_add_campaign_max_donation_amount.up.sql` SHA-256 `1e69f949038f539a6763515de9f2bbac9ecd18a09f4f40eb0ad5280dfade1c6d`; down SHA-256 `d59cc62efd3e3dc24401e41b6b959854205b17ffe6a05c79ab68f74b79485c7c`. Verify them at dispatch; do not edit or apply them.
- Owning authorities to re-read by relevant section: `backend/AGENTS.md` SHA-256 `6a2a9c84b007572155be0690fba3c04e1be5e522998a3aff3cfd1f01f96fef8e`; `docs/project/kencleng-monetary-data-standard.md` SHA-256 `2a0dc7f382cf07e28b53758c6ae845405323dcaa68a38ebf8245555da6934ddf`; Campaign invariants SHA-256 `42edf2ab5713b8905e9a9a8e554da2a678fbf70fcbabe95670113cbcdf8b1423`; Donation invariants SHA-256 `4bf7ba48f4980f635ce76a4eb3440be5cf12666fdd99e06dcc8ac3b78fbde39a`; Donation submit feature SHA-256 `b3f0f99c420c807bb14cccf75a4adcaf43dfdadccae6b727a1c460ef4ed6a7e9`; Donation status feature SHA-256 `332bba24b3f00cab880f6d5671d3e3c5cf9d99633be6fdf0212f0a05b5b5fd2c`; accepted Donation API SHA-256 `9f7c31065c3c7ffa77491541102afc0ce84085f50a1e2aec317336c9656d5c1d`. Re-verify all authority hashes before review and stop on material drift.
- Migration practice: `../harscode-workspace/best-practices/postgresql/migrations-safety.md`, SHA-256 `ab368f092d5df4ff265561a68e5b30e2a7e87d2fe7f1108106d8fcebf2f39f93`; route via current `best-practices/AGENTS.md` and use only applicable guidance. Assess additive compatibility, existing-row/constraint validation, actual PostgreSQL lock concerns, and distinction between empty pre-write cleanup and destructive post-write down behavior. Do not claim runtime lock or up/down evidence from a written proposal.
- Gate context: `TP-S2-003-011` report accurately records current decision boundary. Existing bounded cap projection review/testing (`RV-S2-003-004`, `TST-S2-003-001`) does not cover this Donation/D1 design. Open Item 7 Organization eligibility source/update schema and affected Campaign handlers remain explicitly excluded and authority-gated.

## Review task and boundaries

Determine whether the exact approved design is sufficiently safe, complete and consistent with its approved Techplan and current owning sources to satisfy the positive precondition for a later minimum Donation/D1 schema Build. Assess the design's answers to RV-S2-003-005's still-open F-02/F-03/F-04/F-06 independently, including migration sequencing/additive compatibility, data integrity and constraints, pending aggregate/index fit, endpoint-wide idempotency equivalence representation, close-reason history, Campaign FK/deletion behavior, existing-row/lock/down-migration posture, and exact Rupiah compatibility. Identify any residual assumption, deferred external evidence, or missing decision owner; distinguish `verified`, `assumed`, `deferred`, and `not tested`.

Do not review or reopen whole-plan semantic decisions that are settled in the Approved Techplan. Do not include Open Item 7's Organization source/membership mutation design; it remains gated separately. Do not select O3 idempotency lifetime, O4 credential/key/comparison controls, O5 abuse/proxy choices, email retention/provider, residual-risk acceptance, or new Product/API semantics. Do not invent SQL as approved. If a material conflict or incomplete schema property remains, identify precise location, consequence, suggested resolution and whether it blocks this design gate. A positive design verdict satisfies only this pre-implementation design-review precondition; it does not authorize migration application, Tier-0 implementation, broad Testing, runtime, or milestone.

## Execution envelope and output

- Re-read the exact pinned Techplan sections, prior RV-S2-003-005 report and current sources above; follow all four independent Review passes (Safety, Quality, Stack-Specific Best Practices, Consistency) against this same written design target. Use the canonical review prompt, adapting its current-diff assumption to this explicit design-artifact target.
- Write only `RUN_PATH/review-findings-1.md` and `RUN_PATH/launch-record.md`; report no findings when none are warranted. Include a focused verdict, exact provenance, verification performed, and one structured `## Phase handoff`.
- Do not edit the Approved Techplan, proposal, Product/spec/API/code/tests/migrations, parent state, or other Runs. Do not run tests, validators, generators, migration commands, connect to or mutate a database, or claim PostgreSQL execution/lock/reversibility evidence. Read-only source inspection is allowed.
- No downstream Run is dispatched by this Participant. Stop after the report and handoff.

## Human-assisted dispatch

- Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`
- Fresh Reviewer / `KC-REVIEWER` session; configured `gpt-6-luna` / `high`.
- Canonical kickoff: `Jalankan fresh independent migration-design Review Run RV-S2-003-008 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-008/invocation.md dan current canonical ../harscode-workspace/workflow/4-code-review-prompt.md beserta Review guidelines/checklist serta orchestrated-run overlay. Target hanya desain minimum Donation/D1 di exact Approved Techplan hash yang dipin; ikuti empat pass secara independen dan gunakan sumber/schema anchors terdaftar sesuai authority. Secara independen nilai apakah revisi menutup RV-S2-003-005 F-02/F-03/F-04/F-06 dan cukup untuk positive migration-design gate. Jangan diarahkan ke verdict tertentu. Jangan edit Techplan, SQL, code, tests atau state lain; jangan menjalankan/apply migration, database action, tests, validators atau generators. Tulis review-findings-1.md dan launch-record dengan satu Phase handoff, lalu berhenti.`

This Invocation prepares a Run only. Anhar dispatches the fresh Reviewer; no Participant has been dispatched by the Orchestrator.
