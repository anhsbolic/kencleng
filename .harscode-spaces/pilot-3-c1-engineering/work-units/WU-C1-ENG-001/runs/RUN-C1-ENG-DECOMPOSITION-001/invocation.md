# Run Invocation — RUN-C1-ENG-DECOMPOSITION-001

## Identity, route, and assignment

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-DECOMPOSITION-001`.
- **Phase route / Role:** Post-Approval Techplan decomposition / Planner.
- **Specialization / Profile:** None; canonical Planner with bounded approved-spine decomposition scope.
- **Participant / Session:** `PARTICIPANT-C1-ENG-DECOMPOSER-001` / `SESSION-C1-ENG-DECOMPOSER-001` (reserved; instantiate on mechanical Human dispatch; FRESH).
- **Transition reason:** New workflow execution after terminated planning/review, exact Techplan approval, and separately recorded G1–G3 authorization.
- **Dispatch posture:** Human-Assisted; prepared, not dispatched.
- **Escalation owner:** Anhar via Orchestrator.

Evaluate the canonical decomposition Step 0, then, only if useful, generate complete scoped execution task files and the manifest from the approved parent. Use the actual work/context/risk boundaries, not document length. Parent §9 already identifies T1 shared contract, T2 backend identity/session, T3 backend aggregate/guard, T4 operator docs, T5 frontend, and T6 review/verification obligations; these are evidence for the gate, not a preselected topology.

Choose and justify the least-surprising axis. Preserve all parent scope/rule/decision/risk/interface/authority/verification meaning. Child tasks operationalize the approved contract with enough detail to execute independently from parent + current task + declared dependencies + live authority. Keep backend/frontend production writes separate, shared API/operational docs concern coordination explicit, and cross-layer atomicity/authority invariants visible. Do not split just to produce one task per file or copy the parent six times.

For each hard dependency, derive the smallest observable predecessor result and exact dependent boundary from settled parent semantics; identify expected durable evidence. Ordering or broad predecessor completion is insufficient unless it is the actual condition. No new API/security/ownership/verification obligations may be introduced in condition wording. Preserve evidence owners and real-versus-simulated distinctions. Generate the manifest last as an execution map, not a status/progress ledger or model registry. This is task decomposition inside one WU; do not invent new Work Units/Work Graph topology.

If deriving a boundary/condition requires a missing material decision, stop with the exact gap for parent reconciliation; do not repair the approved spine or hide new semantics in a child. A NO gate result stops without task generation and reports its reason. A YES result returns the generated shape and dependencies for Human inspection/acceptance before dependent Build under the canonical prompt. Do not dispatch Build.

## Assignment-defining source and authority

- **Approved parent:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`, SHA-256 `bd732748b486f0327f6cba2cdf14cab32f1ece3c5be8ebb24f9659444947fa47`. Read in full; verify at entry and terminal handoff; parent read-only for this Run.
- **Approval lineage:** Anhar approved exact reviewed/reported source SHA-256 `577673c2b03bc93526362a848636699ebd888b8917519bae65b1adde15860954`; approval-only lifecycle hash `1eecea60b56b2babf843e58ea2b1fa813383be366b25bcce52ff8aaa05e2b657`; current hash adds permission lifecycle resolution only. No executable semantic change.
- **Protected authorization:** WU `events.md` owns Anhar's separate “ya aku izinkan” decision for bounded G1 auth/session, G2 initial Owner/object scope, G3 guard controls; current parent Resolved 9–11 indexes it. Do not request those same permissions again or expand them. They authorize implementation, not provider verification, legal adjudication, or environmental guard opening.
- **Solution evidence:** `solution-shaping/solution-contract.md`, SHA-256 `c136e673937c9ac1a583ecfa6a98ee630e1b0f6b42b78d373876243e55b01482`; use relevant selected-boundary/verification/gate anchors where parent references need grounding. Its shaping-time pending approvals are historical, superseded only by owning Human events.
- **Prior artifacts:** Completed `runs/RUN-C1-ENG-TECHPLAN-002/evidence/phase-handoff.md`, `runs/RUN-C1-ENG-TECHREVIEW-001/evidence/review-findings.md` and `phase-handoff.md`, `runs/RUN-C1-ENG-TECHPLAN-003/evidence/convergence-check.md` and `phase-handoff.md`. They establish prior synthesis/clean review/convergence, not implementation.
- **Gate-time report:** `techplan/report-techplan.md`, SHA-256 `92583b24ef0b20818191748bff39d3357e5412f3da9c2a05f1eff03f13aefae4`. Preserved presentation-time digest; pending-permission/In Review wording describes that earlier gate. Do not regenerate it in decomposition.
- Paths above without a full repository prefix are relative to `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/`.
- **Current-effective guidance:** Root `AGENTS.md`, relevant scoped instructions, Stage 7 authority route to binding Stage 5/6, current project architecture/UX/API/live sources where a task boundary depends on them. Upstream Product/pre-engineering authority remains read-only.
- **Phase guidance:** Harscode root/workflow routing → `workflow/2-3-techplan-decomposition-prompt.md`, `workflow/2-techplan/rules.md` §10, and `workflow/orchestrated-run-overlay.md`. Use targeted current sources for any non-obvious technical fact; do not default-load unrelated best practices. Ordinary guidance remains current-effective with observed provenance, not universally pinned.

## Bindings, runtime route, and model rationale

- **CWD / project root:** `/home/anhar-solehudin/kencleng-workspace/kencleng`.
- **Harscode workspace root:** `../harscode-workspace`.
- **Work Unit record:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md`.
- **WORK_UNIT_PATH / TASK_PATH compatibility root:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-DECOMPOSITION-001`.
- **ARTIFACT_TARGET:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/`, only scoped task snapshots and manifest when Step 0 is YES. No task directory exists at preparation.
- **PRIOR_ARTIFACTS:** Explicit approved-spine/decision/evidence pointers above replace ordinal workflow paths.
- **Communication / profile:** Bahasa Indonesia / none established; preserve canonical terms and identifiers.
- **Ticket / area:** None established; C1 — Legitimate Organization representation.
- **TARGET_REVISION:** `524ef600c7f71246af6b71671d89c3c040fd9d44`, `pilot/3-c1-engineering`, observed with existing coordination/synthesis/review/report/lifecycle working-tree changes. Baseline only; exact parent hash defines the task.
- **WORKFLOW_REVISION:** `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`, observed at preparation; Harscode working tree clean.
- **Harness / session posture:** `codex-cli` / fresh Participant reconstruction, Human-Assisted mechanical dispatch; no native subagent substitution.
- **Selected model / effort:** `gpt-6-luna` / `medium`.
- **MODEL_APPROVAL:** NOT REQUIRED; Human registry declares this model `approval_required: false`.
- **Routing rationale:** Decomposition redistributes already-approved, independently reviewed T1–T6 scope and derives conditions from explicit parent interfaces/dependencies without architecture/security redesign. Registry-declared general/reasoning/coding/repository-work capabilities are sufficient for this bounded task; select the lower cost tier. Medium is the lowest effort judged sufficient for preserving task completeness, dependencies, and shared invariants. Missing material meaning triggers semantic rerouting, not stronger-model substitution. Demonstrated capability insufficiency with complete inputs routes to Orchestrator for separately authorized escalation if necessary.
- **CONTINUATION_CHECKPOINT:** None; new Run.
- **Runtime verification context:** Podman/podman-compose configuration is downstream task evidence context, not permission to run containers/tests in this planning phase.

## Execution envelope and completion evidence

- **PREAUTHORIZED:** Read approved/current authorities and source anchors; evaluate the gate/axis; write only task files/manifest under ARTIFACT_TARGET and Run-owned evidence under RUN_PATH.
- **ORCHESTRATOR_DECISION:** Reconcile returned task topology/condition evidence and route subsequent scoped execution after the canonical Human shape check. No Participant self-dispatch.
- **HUMAN_REQUIRED:** Acceptance of generated split at the canonical handoff; materially changed parent/authority/risk or new protected scope; gated model escalation. G1–G3 implementation authorization already exists for its exact surface and must not be requested again without a material scope change.
- **Out of scope:** Parent/report/solution/history/authority/projection rewrites, production/API/migration writes, tests/build/runtime execution, approval/promotion, new WUs/Work Graph, implementation or C1 completion claims.

Produce one structured terminal `## Phase handoff` in `RUN_PATH/evidence/phase-handoff.md`: known provenance, explicit Step 0/axis outcome, parent identity before/after, generated task/manifest pointers/content identities when applicable, Findings, Decision requests, Blockers, Open / unverified, Recommended continuation, and Context refs. Report the concrete split/conditions for Human inspection if generated, or the exact NO/gap outcome. No lifecycle/permission metadata changes to the parent are assigned.

## Prepared state

- **Dispatch readiness:** READY_FOR_HUMAN_DISPATCH.
- **Participant dispatched:** No.
- **Run outcome:** Not started.
- **Next action:** Human launches fresh Planner decomposition with `gpt-6-luna` / `medium`; no additional model approval is required.
