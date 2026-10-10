# Run Invocation — RUN-C1-ENG-BUILD-T1-PATCH-001

## Identity and assignment

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T1-PATCH-001`.
- **Phase route / Role:** Build/Patch — bounded T1 shared API contract corrections from independent Code Review / Implementer.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T1-PATCH-IMPLEMENTER-001` / `SESSION-C1-ENG-T1-PATCH-IMPLEMENTER-001` (FRESH).
- **Session reason:** New Build occurrence after completed Code Review; reconstruct only from this invocation and durable inputs.
- **Dispatch posture:** Human-Assisted; READY_FOR_HUMAN_DISPATCH. Participant not yet dispatched.
- **Escalation owner:** Anhar via Orchestrator.
- **Communication:** Bahasa Indonesia; retain canonical identifiers and technical terms.

## Current-effective inputs

All paths are relative to the project root unless stated otherwise.

- Approved parent: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`, SHA-256 `bd732748b486f0327f6cba2cdf14cab32f1ece3c5be8ebb24f9659444947fa47`.
- Current task: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/T1-shared-contract.md`, SHA-256 `bbd5ab13c882ca3534ab3397ca908f70157ad41c9a81ee7cafd2827dd94946e3`.
- Accepted manifest: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/manifest.md`, SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.
- Review findings: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T1-001/evidence/review-findings-1.md`, SHA-256 `bc381d666c3fd4fb79f1360ad4dfce976b924a028a2d272b216a511fb07b6871`.
- Specific patch plan: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T1-001/evidence/patch-plan-1.md`, SHA-256 `c3b371562c96bb5641b5b5ef521c2667368ae71a162240145c9bd80764d4c5a7`.
- Prior T1 Build report: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T1-001/evidence/build-report.md`, SHA-256 `1be1b5ad607e3ad7f1a031a65def52d60ba95d33f6d2c1d9959c8619f0b71137`.
- Reviewed target at preparation: `api/openapi/index.yaml` `8ff663e975c9d72e0d6b636994533f420fe77639398197e6137fae49e59fb62d`; `api/openapi.yaml` `9e20887011ad30a3c774edd0fb725659bccd28446fe5e4aa70c0216fbcb7a4e9`; `api/openapi.d.ts` `a0ae33bccdf194190f3dbb49d69435d983acb8cc28fd712b3a2869442dbe6e6e`.
- Supporting API inputs: `api/package.json` `b3ef5560aa7d58f5320a74dbbdbcbe9eadc0d8431b5459667e8c5605b4a1e954`; `api/README.md` `fd769b04dacdbd3340b95f83e5df3d3908315812cd585ef68cbcfd78c07b5cf5`; `api/.gitignore` `4d56952b0fb13bf8f9b6c13a6d4c34a075bac3af447636a1df4335d7576e2f97`.
- **TARGET_REVISION:** `112c35f8322b53bac39d511ed4f22982cea619af`, branch `pilot/3-c1-engineering`; assigned output is working-tree content. Preserve unrelated changes.
- **WORKFLOW_REVISION:** `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`; Harscode tree observed clean at routing.
- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T1-PATCH-001`.
- **ARTIFACT_TARGET:** `api/openapi/index.yaml` editable source plus derived `api/openapi.yaml` and `api/openapi.d.ts`; edit `api/README.md` only if needed to clarify settled environment/browser rules. Run-owned report: `RUN_PATH/evidence/patch-report-1.md`.

## Patch scope

Execute only F1–F4 from the bound patch plan, against the approved parent §8 and R9/R10/R13 and current T1 task:

1. Align production cookie name/flags/host-only semantics with R10, define the explicit localhost:8080 development exception separately, and align issuance/rotation/logout expiry.
2. Declare the shared exact configured Origin, JSON, session-bound CSRF, pre-processing denial, and no credentialed CORS rules for logout/prepare/confirm. Give logout a consumer-usable closed empty JSON body. Keep browser-bound OIDC state/nonce/PKCE callback behavior distinct.
3. Correct callback success/error query branches and safe fixed-local error redirect semantics; do not issue a session on failure or reflect provider values.
4. Declare safe private/no-store 400 `InvalidRequest` responses for applicable invalid read inputs while preserving identical 404 semantics for valid unknown/out-of-scope IDs.

Re-ground current live source and generated artifacts before editing. Generate bundle/types with existing locked tooling. Keep the patch within the authorized API contract and derived artifacts/docs needed for settled behavior.

## Settled boundaries and permissions

- Techplan approval, the exact five-task split/conditions, and explicit G1–G3 implementation authorization remain effective per WU `events.md`; this Run does not reopen or enlarge them.
- This assignment changes declarative shared API artifacts only. Do not edit backend/frontend production, migrations, authentication/session/authorization core, Product, Techplan, task snapshots, manifest, prior Run evidence, or control projections.
- Do not add security policy, authority-bearing request fields, provider behavior, status meaning, license, fabricated 2xx, lint suppression, dependencies, or consumer handwritten types. If the settled contract cannot uniquely resolve a material detail, stop and report the smallest gap.
- Preserve receipt/replay/unknown-outcome behavior, object/person scope, guard semantics, OrganizationView effects, and privacy declarations as recorded by the patch plan.

## Build authority and verification

Use current `../harscode-workspace/workflow/3-build-prompt.md`, `3-build/guidelines.md`, `3-build/checklist.md`, and `orchestrated-run-overlay.md`; project `AGENTS.md`, `api/README.md`, package scripts/lockfiles, and approved parent/task. Re-ground on source before edits.

Run focused existing API validation, bundle generation, and type generation; inspect generated diff. Confirm source/bundle parsed equivalence, production/development cookie distinction, mutation JSON/Origin/CSRF/CORS contract, callback branch/status semantics, and read 400 responses in generated operation types. Record exact commands, results, hashes, and limitations. Do not run final browser/provider/runtime/DB/security-class or broad Testing-owned suites for this declaration-only patch.

## Execution envelope and routing

- **PREAUTHORIZED:** Read bound/current sources; edit the API source and derived bundle/types within the patch plan; narrowly clarify existing README; use locked existing tooling; run focused Build checks; write this Run's patch report.
- **ORCHESTRATOR_DECISION:** Stop and route material contradiction or scope change; after completed patch, return to independent Code Review because fixes change auth/browser/HTTP contract semantics. Orchestrator alone assesses other task dependency evidence.
- **HUMAN_REQUIRED:** Any new material requirement/authority/risk choice, scope beyond the patch plan, destructive action, deployment/publication, or protected implementation beyond the already authorized G1–G3 surfaces.
- **Model / effort:** `gpt-6-luna` / `medium`; registry approval not required. Bounded implementation of explicit Review resolutions using existing tooling; no new architecture/security choice. If capability proves insufficient despite complete inputs, stop and route to Orchestrator rather than infer semantics.
- **Harness / session:** `codex-cli` / fresh Implementer, Human-Assisted dispatch.
- **Dispatch / outcome:** Human launch and completed outcome returned by Anhar. `RUN_PATH/evidence/patch-report-1.md` records COMPLETED, the F1–F4 patch, focused tooling results, and remaining runtime verification limits.
- **Result identity:** Patch report SHA-256 `92d8d1f8cfa86eb4b4d1c197bbaa493b09531d504b0a43846a62937c0016e3f7`; patched source/bundle/types SHA-256 values are indexed in that report and WU current state.
- **Next route:** Fresh independent Code Review `RUN-C1-ENG-CODEREVIEW-T1-PATCH-001` is prepared; its separate model approval is pending. Do not resume or redispatch this completed Build/Patch Run.

## Terminal evidence

Write `RUN_PATH/evidence/patch-report-1.md` with canonical Build/Patch sections and exactly one structured `## Phase handoff`. Name Code Review as the requesting phase and recommend returning to independent Code Review after patch completion. Distinguish generated/validated contract artifacts from runtime behavior. No downstream task or C1 completion claim.
