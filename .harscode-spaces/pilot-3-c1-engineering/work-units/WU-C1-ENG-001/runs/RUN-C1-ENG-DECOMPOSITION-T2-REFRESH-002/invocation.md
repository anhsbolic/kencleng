# Run Invocation — RUN-C1-ENG-DECOMPOSITION-T2-REFRESH-002

## Identity and routing

- Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-DECOMPOSITION-T2-REFRESH-002`.
- Phase: Post-Approval affected T2 task-snapshot reconciliation / Planner (Decomposition ownership).
- Participant / Session: `PARTICIPANT-C1-ENG-T2-TASK-REFRESH-002` / `SESSION-C1-ENG-T2-TASK-REFRESH-002` (fresh).
- Model / effort: `gpt-6-luna` / `medium`; no registry approval required.
- Dispatch: Human-Assisted; ready, not dispatched.

## Assignment and exact scope

Using current-effective Approved Techplan, refresh only `techplan/tasks/T2-backend-identity-session.md` so its parent hash, D8 baseline, and Build gate reflect approved exact Go 1.26.9; OIDC v3.16.0; OAuth2 v0.34.0; go-jose/v4 v4.1.4; pgx/v5 v5.9.2; and explicit x/text v0.41.0 planned baseline. State clearly that approval of planned pins is not evidence the selected actual graph/compatibility/scan is clean. Preserve T2/G1 scope and permission, stop-before-dependency-commit checks, verification owners, T1 prerequisite status, all manifest dependency conditions, five-task topology and G1–G3.

Do not edit parent Techplan, T1/T3/T4/T5 snapshots, manifest, contract, Product/pre-engineering authority, implementation/dependency files, or any other target. Verify manifest hash unchanged. No Build, tests, dependency operations, runtime probes, blocker closure, promotion, or downstream dispatch. If refresh cannot remain T2-only or requires changing a condition/topology, stop and report the precise gap. Write one structured Phase handoff with exact before/after parent/task/manifest hashes and unchanged-topology confirmation.

## Bound inputs

- Current Approved Techplan `techplan/techplan.md`, SHA-256 `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79`.
- T2 snapshot `techplan/tasks/T2-backend-identity-session.md`, SHA-256 `ba34963b725b25f78d908b2e72fceb66b26aa207701110fb0031ffc34affcd8d`.
- Accepted manifest `techplan/tasks/manifest.md`, SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.
- Exact approval: Anhar approved candidate SHA-256 `26c8948c9a1a8687b173fce5a473314dfc735ec47d2176a4f2fa82fee733a32b`; lifecycle promotion is recorded in current parent/Events.
- T2 blocker report `runs/RUN-C1-ENG-BUILD-T2-002/evidence/build-report.md`, SHA-256 `474e39ddd7358c4bf0f1c94a1cb217406cd81285046b06207a8cb4332e91bf93`.
- Target/workflow revision `59f2b25940fa4baf243b4f1d2ca2ba04c9b17fd4` / `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

## Bindings and execution envelope

- RUN_PATH `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-DECOMPOSITION-T2-REFRESH-002`.
- ARTIFACT_TARGET T2 task snapshot only.
- Follow `../harscode-workspace/workflow/2-3-techplan-decomposition-prompt.md`, `workflow/2-techplan/rules.md` §10, and overlay.
- PREAUTHORIZED: read parent/task/manifest and reconcile only T2 snapshot; write Run evidence.
- HUMAN_REQUIRED: none for this bounded refresh; approved split and permissions remain current.

## Prepared state

- Dispatch readiness READY_FOR_HUMAN_DISPATCH.
- Participant dispatched: No. Run not started.
- Human mechanically dispatches a fresh Planner Session; no model approval required.
