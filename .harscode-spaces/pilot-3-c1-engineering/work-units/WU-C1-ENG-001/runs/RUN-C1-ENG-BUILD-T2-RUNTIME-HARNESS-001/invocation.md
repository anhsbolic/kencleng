# Run Invocation — RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001

## Identity and route

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001`.
- **Phase / Role:** Build / Implementer — backend T2 auth runtime test harness.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-RUNTIME-HARNESS-IMPLEMENTER-001` / `SESSION-C1-ENG-T2-RUNTIME-HARNESS-IMPLEMENTER-001` (fresh; independent from all prior Build, Review and Testing contexts).
- **Dispatch:** Human-Assisted; terminal `COMPLETED` after Human dispatch and one terminal Build handoff on 2026-10-13. This Run is closed; do not reuse or reopen it.
- **Selected model / effort:** `gpt-6-sol` / `medium`.
- **MODEL_APPROVAL:** `APPROVED_BY_HUMAN` for this exact Run, `gpt-6-sol` / `medium`. On 2026-10-13 Asia/Jakarta, Anhar replied “approve bro” to the explicit exact-Run model and fixture request. This approval does not transfer to another Run.
- **FIXTURE_AUTHORIZATION:** `APPROVED_BY_HUMAN` for this exact Run's one isolated disposable PostgreSQL container and ephemeral provisioner/restricted runtime identities, bounded exactly by `scope.md`. The authorization is limited to that Run; prior AUTH-HARNESS-001 approval was not transferred.
- **Routing rationale:** This test-only Build changes the boundary that controls fixture ownership labels, role separation, cleanup and real PostgreSQL concurrency evidence for authentication/session behavior. The gated Sol model's declared security/cross-cutting reasoning fits those concerns; `medium` is its lowest supported sufficient effort.
- **Target revision / branch:** `dae55ac44beb0e967b94c2f791d8a0b8d8e72372` / `pilot/3-c1-engineering`; current exact target source is separately bound.
- **Workflow revision:** `651ce57f044bff23fba99082322cee6c066d25a6`.
- **Harscode root:** `../harscode-workspace`.
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001`.
- **Artifact target:** None; changed repository code is implementation output, report/evidence are Run-owned.
- **Working directory:** `/home/anhar-solehudin/kencleng-workspace/kencleng`; fixture/test command from `backend/`.
- **Communication language:** Bahasa Indonesia.
- **Scope:** `scope.md` (SHA-256 `d86d6474409252dddd6db2a38683b74ba6c9ee1c2df256ed306d4ed9fb15ffda`). Assignment-defining inputs: `input-bindings.json` (SHA-256 `049f3697ff8242e5ee3a1c111b2ff9ad924cdea9da695cf793ead55b3a6edea1`, 34/34 live bindings matched during preparation).

## Claim and current evidence

This is a test-infrastructure Build claim only: make the existing auth harness capable of using the active Run ID as its exact resource/test owner and add a bounded full-T2 suite selector for a later independent Testing Run. Existing tagged tests already cover runtime privilege/SCRAM checks, registered HTTP behavior, simulated callback, concurrent person upsert, session lock contention and null logout. The current runner is hardcoded to LOGOUT-PATCH-001 and selects only two tests.

Preserve terminal AUTH-HARNESS-001 FAILED evidence (null returned 204 before patch), terminal LOGOUT-PATCH-001 Build PASS, terminal independent logout Review (0 blocking / R1 nonblocking), and current BACKEND-GATE-002 Testing PASS report. The current gate establishes only `make verify` / `go mod verify` on its bound tree; this Build's later file delta will require fresh post-Build command-gate evidence in the next Testing assignment. No blocker closure or T2/T3 acceptance follows from this Build.

## Required workflow and execution

Read the current Approved Techplan/T2 task, `backend/AGENTS.md`, current OpenAPI object contract, `docs/project/c1-auth-verification-environment.md`, canonical `../harscode-workspace/workflow/3-build-prompt.md`, Build guidelines/checklist, orchestrated overlay/run contract, and `input-bindings.json`. Re-ground on live `auth_harness.py`, tagged `harness_test.go`, README instructions, active Handler/Store/migration and terminal Build/Review/Testing reports. Preserve existing dirty/untracked work.

After both exact approvals and Human dispatch, change only the three scope-listed files. Execute exactly the Run-owned fixture command:

```text
python3 scripts/auth_harness.py --run-id RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001 --suite t2
```

from `backend/`, with Go 1.26.9 and bounded timeouts. The runner must validate Run ID and suite before resource creation, use the exact Run ID in fixture labels/test environment/cleanup ownership checks, and select all six existing tagged `TestHarness...` cases for `suite=t2`. Retain assertion failures. Preserve safe setup/test/cleanup evidence and categorize simulated Protocol separately from actual PostgreSQL/TCP/SCRAM/HTTP fixture behavior.

No broad `make verify`, DB shared/persistent access, `.env` read, real Google/provider/browser/HTTPS, app startup, migration against non-disposable state, source/API/Product/Techplan/dependency/scanner-policy edits, blanket allowance, commit or PR. No source fix for the nonblocking Review R1 is included. Report exactly one structured `## Phase handoff`.

## Execution envelope

- **PREAUTHORIZED after both approvals:** exact three-file test-infrastructure writes, isolated disposable fixture/provisioner/runtime setup, focused all-six-case harness Build suite, scoped result evidence.
- **ORCHESTRATOR_DECISION:** reconcile actual delta and decide independent Review/Testing applicability and current-tree gate re-entry.
- **HUMAN_REQUIRED:** mechanical dispatch, any privilege/resource expansion, shared/persistent DB or secret access, production auth semantic/source change.

Terminal reconciliation: Human dispatch is evidenced by the terminal report. The report SHA-256 is `c43a030248673bd094aef4352758f4d0007bb14bf255c6202dc1708be194f34`; it contains exactly one `## Phase handoff`. The Build occurrence is closed and its report/evidence must remain unchanged.
