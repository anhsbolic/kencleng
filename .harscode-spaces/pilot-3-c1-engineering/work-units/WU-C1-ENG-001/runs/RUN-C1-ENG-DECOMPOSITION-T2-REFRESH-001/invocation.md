# Run Invocation — RUN-C1-ENG-DECOMPOSITION-T2-REFRESH-001

## Identity and routing

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-DECOMPOSITION-T2-REFRESH-001`
- **Phase route / Role:** Post-Approval affected task-snapshot reconciliation / Planner (Decomposition ownership).
- **Specialization / Profile:** None; canonical Planner with bounded task-snapshot scope.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-TASK-REFRESH-001` / `SESSION-C1-ENG-T2-TASK-REFRESH-001` (fresh; new Run).
- **Transition reason:** The approved successor changes T2's toolchain/auth baseline and Build gates; canonical decomposition guidance requires reconciling affected task snapshots before execution. This is not a re-opened decomposition-shape review.
- **Dispatch posture:** Human-Assisted; prepared, not dispatched.
- **Human owner:** Anhar via Orchestrator.

## Assignment and exact scope

Using the current-effective Approved Techplan, reconcile only `techplan/tasks/T2-backend-identity-session.md` so it derives from the approved successor and carries its exact parent hash, D8 baseline, R14/§10 dependency verification gates, and current T1/OpenAPI prerequisite status. Preserve T2's G1 implementation boundary, implementation outcome, test/evidence ownership, and all exact manifest-declared task dependency conditions. Treat backend `go.mod` Go 1.24.8 as the observed current baseline anchor; distinguish it from the approved planned Go 1.26.9 runtime/build toolchain and explicit auth pins. Require exact pins from the parent and actual Build verification of toolchain, resolved graph, integrity, source scan, compatibility, compile, surviving advisories, and no dependency commit before those checks pass.

Do not edit the approved parent, T1/T3/T4/T5 snapshots, manifest, Solution Contract, API contract, or any upstream authority. Do not alter the dependency graph, task boundaries, T1→T2/T2→T3/T1→T3/T1→T5 conditions or declare any new condition. Verify manifest hash is unchanged. No Build, code, dependency, tests, runtime probes, T2 blocker closure, or C1 completion. If the new parent requires materially changing task boundaries/conditions or another task snapshot, stop and surface that precise scope; do not broaden this refresh.

The topology remains the accepted five-task split. The decomposition prompt's Step 0 shape gate is not reopened; use its snapshot invariants to reconcile the single affected task. No new human split acceptance is required for this exact scope-only refresh. Complete with exact parent/task/manifest hashes, a concise change inventory, explicit no-topology-change confirmation, and one structured Phase handoff.

## Bound inputs and provenance

- **Current-effective Approved parent:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`, SHA-256 `bae2470f9038521b5d2d51fe4bc8de192fe5a9cf91bc0975831a3046fbf7a8ad`.
- **Affected T2 snapshot:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/T2-backend-identity-session.md`, SHA-256 `075170761afb971cd09aba90b1e112c617f12d27d4b8bb4ffbd1a3599163378b`.
- **Unchanged accepted manifest:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/manifest.md`, SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.
- **Approval evidence:** Events record exact approved candidate SHA `2c64128e95a52773da024c2ec856bc5ba7a047539f2197a6d747513d378ae629`; promoted plan header/Open Items record the bounded lifecycle transition.
- **T2 blocker evidence:** `RUN-C1-ENG-BUILD-T2-001/evidence/build-report.md`, SHA-256 `130b888011a5c5286b818078410d2b022faf60d0033f20d0e89d628098c0ad50`; preserve no-production-write and all 50 finding context.
- **T1 dependency evidence:** accepted manifest + T1 patch Review `RUN-C1-ENG-CODEREVIEW-T1-PATCH-001/evidence/review-findings-1.md` SHA `f8a61a3bfd591a5310e291e84ec0aac393b615c731216238abbacee0342782ea`; T1 Build/Patch report as referenced in WU. This establishes contract prerequisite only, not backend implementation/runtime.
- **Target/workflow revision:** target `59f2b25940fa4baf243b4f1d2ca2ba04c9b17fd4`, branch `pilot/3-c1-engineering`; workflow `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

## Bindings and execution envelope

- **CWD / project root:** `/home/anhar-solehudin/kencleng-workspace/kencleng`
- **Harscode root:** `../harscode-workspace`
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`
- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-DECOMPOSITION-T2-REFRESH-001`
- **ARTIFACT_TARGET:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/T2-backend-identity-session.md` only.
- **Harscode phase guidance:** `../harscode-workspace/workflow/2-3-techplan-decomposition-prompt.md`, relevant `workflow/2-techplan/rules.md` §10, and orchestrated overlay.
- **Communication:** Bahasa Indonesia; canonical Harscode terms/identifiers retained.
- **Harness:** `codex-cli`, Human-Assisted, fresh Session.
- **Model / effort:** `gpt-6-luna` / `medium`; Human registry does not require per-Run approval.
- **PREAUTHORIZED:** Read approved plan/task/manifest/live backend module anchors and governing workflow; reconcile only T2 snapshot; write Run evidence.
- **ORCHESTRATOR_DECISION:** Reconcile exact snapshot output and route report refresh/Build preparation.
- **HUMAN_REQUIRED:** No new product/security/decomposition decision in this bounded refresh. Any material mismatch beyond the exact approved baseline is routed; existing G1–G3 authorizations remain unchanged.
- **Out of scope:** Other task/manifest/parent/API/Product writes; changing topology/conditions; production implementation/dependency mutation; tests/build/runtime; blocker closure; downstream dispatch.

## Terminal evidence

Write one structured `## Phase handoff` at `RUN_PATH/evidence/phase-handoff.md` with parent/task/manifest hashes before/after, exact allowed delta, unchanged-topology confirmation, outcome, open/blocker status, and next route. Do not duplicate task content into Run evidence.

## Prepared state

- **Dispatch readiness:** READY_FOR_HUMAN_DISPATCH.
- **Participant dispatched:** No.
- **Run outcome:** Not started.
- **Next action:** Human mechanically dispatches fresh Planner/Decomposition context with `gpt-6-luna` / `medium`; no model approval required.
