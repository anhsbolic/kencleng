# Run Invocation — RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-CLEANUP-PATCH-001

## Identity and route

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-CLEANUP-PATCH-001`.
- **Phase / Role:** Build/Patch / Implementer — backend disposable auth harness cleanup.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-RUNTIME-HARNESS-CLEANUP-PATCH-IMPLEMENTER-001` / `SESSION-C1-ENG-T2-RUNTIME-HARNESS-CLEANUP-PATCH-IMPLEMENTER-001` (fresh, separate from terminal Build/Review/Orchestrator contexts).
- **Dispatch:** Human-Assisted; `READY_FOR_HUMAN_DISPATCH` after package preparation; Human mechanical dispatch remains pending.
- **Selected model / effort:** `gpt-6-luna` / `medium` (Human registry: `approval_required: false`; approval is not required for this exact selection).
- **Model routing rationale:** F1 has a precise accepted patch plan and affects only a test-fixture process boundary, not production authentication. Luna has the declared coding/repository/reasoning capabilities; medium is sufficient for bounded cleanup/error-path implementation plus controlled command-boundary tests. Escalate to Human/Orchestrator if live state requires privilege or scope expansion.
- **Target revision / branch:** `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering` (observed preparation target; exact live assignment content is bound below).
- **Workflow revision:** `651ce57f044bff23fba99082322cee6c066d25a6`.
- **Harscode root:** `../harscode-workspace`.
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-CLEANUP-PATCH-001`.
- **Artifact target:** None; implementation is source output and patch report is Run-owned evidence.
- **Working directory:** `/home/anhar-solehudin/kencleng-workspace/kencleng`; focused test command from `backend/`.
- **Communication language:** Bahasa Indonesia.
- **Scope:** `scope.md` (SHA-256 `b851ab4d96b2e490657c65f244235d66aa38ea938ada15d55d9ad773b2d67020`). Assignment-defining inputs: `input-bindings.json` (SHA-256 `ac25cb86ddda7dc2a500aebe85e63e4e9e9827e58e3a70aa3735a022906bf18d`, 31 bound identities; live recomputation required before editing).

## Claim and prior terminal evidence

This new Build/Patch Run addresses only blocking Review finding F1: a `podman run` that creates a container but times out/returns before its ID is received leaves `created == false`, so normal `finally` cleanup is skipped.

Bound Review report: `RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-001/evidence/review-findings-1.md`, SHA-256 `40a6018c5cb208cc36e688cfd444f1d7bbd3f1f4c5963ee6eaba20decfacd941`.
Bound accepted patch plan: `RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-001/evidence/patch-plan-1.md`, SHA-256 `fe85264fcda7857eab979cfd0434fb584715fdda49c559d9f0002fb5ae961239`.
Prior reusable harness Build report: `RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001/evidence/build-report.md`, SHA-256 `c43a030248673bd094aef4352758f4d0007bb14bf255c6202dc17008be194f34`.

The Review Run's Invocation contains a nonblocking typo in its inline Build-report digest; the canonical Build report digest above is verified from the durable report file. Preserve the terminal Review Invocation/report/patch plan as historical evidence; use this exact verified report hash for this assignment.

## Required workflow and execution

Read the terminal Review report/patch plan, this Invocation/Scope/manifest, current root/backend `AGENTS.md`, active Approved Techplan/T2 only as needed for the fixture/runtime boundary, current `backend/scripts/auth_harness.py`, README, canonical Build prompt/guidelines/checklist, orchestrated overlay/Run contract, and matching subprocess/resource-cleanup/testing guidance. Re-ground on live source and exact bound content before editing; preserve dirty workspace state.

Implement the smallest safe command-boundary repair. Arm reconciliation before invoking Podman create. On ambiguous completion, discover only the exact freshly generated fixture name; require a unique resource and verify Run/owner labels plus immutable container ID before deleting by ID. On absent resource, retain a safe creation-failure outcome. On uncertain lookup/identity/cleanup, fail closed with a clear sanitized unresolved-resource report and perform no unverified deletion. Keep all operations bounded and preserve success-path cleanup.

Add focused stdlib-only tests with mocked command responses for: (1) create followed by timeout or nonzero return and owned-resource cleanup by immutable ID; (2) no resource after ambiguous failure; (3) wrong/ambiguous owner is never deleted. Tests must not call real Podman or create a fixture. Execute exactly the Scope's focused test command from `backend/`. Do not run the broader T2 harness, `make verify`, scanners, or runtime gates in this Run.

Write `evidence/patch-report-1.md` with canonical Build/Patch provenance, exact files/delta, targeted scenarios and outputs, limitations, contract check, and exactly one `## Phase handoff`. Report the real-resource boundary explicitly: mocked command-level evidence only; no container/database or Testing acceptance.

## Execution envelope

- **PREAUTHORIZED after Human mechanical dispatch:** changes in the two required test-harness files, optional narrowly necessary README clarification, focused mocked command-boundary test, and this Run's report/evidence.
- **ORCHESTRATOR_DECISION:** after terminal report, reconcile exact delta and prepare a fresh targeted independent Review of F1 closure; later Testing requires a separate Run/fixture and applicable post-delta gates.
- **HUMAN_REQUIRED:** mechanical dispatch, real Podman/container/DB/provider/browser/runtime access, secret access, scope/privilege expansion, production auth semantic changes, or any material authority choice.

No existing Run or Session was resumed. At preparation, no code/test/Podman/DB/runtime action had occurred in this Run; see terminal reconciliation below for the completed scoped execution.

## Terminal reconciliation

- **Run outcome:** COMPLETED — scoped Build/Patch PASS for the F1 command-boundary cleanup repair.
- **Terminal report:** `evidence/patch-report-1.md`, SHA-256 `59b60504a13f1509191db94e016439f55e94db0dd4c3cc1720a5230d83db02fe`; exactly one `## Phase handoff`.
- **Binding/delta reconciliation:** all 31 pre-edit assignment bindings matched at preparation; live post-run recomputation finds 30 unchanged and the one allowed existing input `backend/scripts/auth_harness.py` changed. Its post-edit SHA-256 is `653cc84c16193d27e3693c4144a1e0bad748a6beff95fc1c7fbe30ec21438605`; the allowed new `backend/scripts/test_auth_harness.py` SHA-256 is `3e78860f863319ef615459746fd68742f418185a980e36f544e5df1bb65cefca`. The other Build output identities for README and Go harness test remain unchanged. No out-of-scope delta was found in the reconciled scope.
- **Evidence boundary:** the focused mocked command-boundary test is Builder-reported PASS (4 cases). No real Podman/container/database/fixture, broad gate, scanner, or other test was run. This does not establish real cleanup or independent runtime acceptance.
- **Run is terminal and closed.** Any re-entry uses a fresh Run, Participant, Session, and Run path.
