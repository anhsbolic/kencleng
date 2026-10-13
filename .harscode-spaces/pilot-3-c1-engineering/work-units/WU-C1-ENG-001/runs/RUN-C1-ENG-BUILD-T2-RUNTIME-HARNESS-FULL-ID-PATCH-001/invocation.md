# Run Invocation — RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-FULL-ID-PATCH-001

## Identity and route

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-FULL-ID-PATCH-001`.
- **Phase / Role / Profile:** Build/Patch / Implementer / backend disposable auth harness cleanup.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-RUNTIME-HARNESS-FULL-ID-PATCH-IMPLEMENTER-001` / `SESSION-C1-ENG-T2-RUNTIME-HARNESS-FULL-ID-PATCH-IMPLEMENTER-001` (fresh; separate from all terminal Build/Review and Orchestrator contexts).
- **Dispatch:** Human-Assisted; `READY_FOR_HUMAN_DISPATCH` after package preparation. Human mechanical dispatch pending.
- **Selected model / effort:** `gpt-6-luna` / `medium` (current Human registry `approval_required: false`; no approval is required for this exact selection).
- **Model routing rationale:** The accepted finding is a narrow validation defect in a disposable test-harness resource-deletion path. Luna declares coding, reasoning, and repository-work capabilities; medium is sufficient to add strict identifier validation and focused failure-injection tests. Escalate if live evidence requires broader privilege, security, or scope decisions.
- **Target revision / branch:** `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering` (preparation HEAD; assignment content is exact-bound).
- **Workflow revision:** `651ce57f044bff23fba99082322cee6c066d25a6`.
- **Harscode root:** `../harscode-workspace`.
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-FULL-ID-PATCH-001`.
- **Artifact target:** None; source changes are bounded Build outputs and the report is Run-owned evidence.
- **Working directory:** `/home/anhar-solehudin/kencleng-workspace/kencleng`; focused test command from `backend/`.
- **Communication language:** Bahasa Indonesia.
- **Scope:** `scope.md` (SHA-256 `db64a3a93d26c615670d520adf30de34659cc5d28ef293c419082b5bccca0c88`). Assignment-defining inputs: `input-bindings.json` (SHA-256 `d254d830873eddc87d9a6a7bb155f8ddd57fb041796f6945ba601a1cec188b8e`, 27 content identities; recompute before editing).

## Claim and exact review finding

This fresh Build/Patch re-entry addresses only blocking F1 in terminal Review `RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-CLEANUP-PATCH-001`: current cleanup accepts any nonempty inspected ID beginning with the discovered Podman ID. It must require a valid full 64-hex inspected ID before any removal, verify the discovered 12–64-hex value is its prefix, and, when present, require `expected_id` from `podman run` to be a valid full ID equal to the inspected ID.

Bound Review report: `RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-CLEANUP-PATCH-001/evidence/review-findings-1.md`, SHA-256 `e1030123ede0df6f9c9620a2f5d18d706cb58ef0e3080f2eb4cabf0322173ab0`; it has exactly one `## Phase handoff` and records 25/25 matched bindings, four completed review passes, Request changes, no edits, and no commands/tests/runtime.

## Required workflow and execution

Recompute all 27 assignment-defining bindings before editing. Re-ground on live `backend/scripts/auth_harness.py`, `backend/scripts/test_auth_harness.py`, the terminal finding, accepted patch plan, prior Build report and exact fixture boundary. Preserve the existing dirty workspace. Stop and report material drift or any wider-than-allowed delta.

Implement only the Scope: strict full-ID and prefix validation before removal, plus mocked negative tests for short/malformed inspect IDs asserting no `podman rm`. Preserve the existing timeout, nonzero, absent-resource, and mismatched-owner cases and success-path cleanup. Run only `python3 scripts/test_auth_harness.py` from `backend/`. Do not run the six-case runtime harness, Make gates, scanners, or any real Podman/container/database/provider/browser/runtime action.

Write `evidence/patch-report-1.md` with compact Build/Patch provenance, exact file hashes/delta, actual focused command/results, F1 resolution, verification limits, and exactly one `## Phase handoff`. State explicitly that evidence is mock-only and is not independent runtime acceptance.

## Loop-health trigger check and evidence boundary

The Orchestrator checked current Harscode's experimental Loop Health trigger criteria before this re-entry. For this bounded cleanup concern, there is one completed Review → patch Build → Review cycle; the prior patch added material implementation and four executable mock cases. No HARD trigger or second WARNING trigger is established, and the explicit single Review/patch/Review non-trigger applies. The trigger threshold is not met; this is not an independent Loop Health diagnosis or `STALLED` declaration.

The mock suite verifies command-boundary behavior only. It does not prove real Podman resource deletion or volume cleanup. Preserve terminal evidence and return the focused patch to a fresh Review route after reconciliation; any later runtime Testing is separately judged against its claim and evidence gaps, not triggered by chronology.

## Execution envelope

- **PREAUTHORIZED after Human mechanical dispatch:** changes to the two files in `scope.md`, the focused mocked test command, and this Run's report/evidence.
- **ORCHESTRATOR_DECISION:** after terminal report, reconcile exact delta and independently decide targeted Review applicability; re-evaluate later Testing only against its exact claim and outstanding evidence.
- **HUMAN_REQUIRED:** mechanical dispatch; real Podman/container/database/provider/browser/runtime access; secrets; production auth/session semantic changes; broader privileges or scope; any material authority choice.

No prior Run or Session is resumed. Model approval from a different Run does not transfer; the selected model here is nongated by the current registry.

## Terminal reconciliation

- **Run outcome:** COMPLETED — scoped Build/Patch PASS for the full-ID validation remainder of F1.
- **Terminal report:** `evidence/patch-report-1.md`, SHA-256 `592001135b7262f5199514ddf7104f19d17f06fa2f32b3bf6df1a8b76f3e6de4`; exactly one `## Phase handoff`.
- **Binding/delta reconciliation:** report records 27/27 assignment bindings before edits. Live recomputation finds 25 unchanged and only the two allowed files changed: `backend/scripts/auth_harness.py` SHA-256 `932a427e7917f73f2917ea316dbe62d29e3ffca649ebc5dc361db646319fe79a`; `backend/scripts/test_auth_harness.py` SHA-256 `4854fd949a3f09ddac611ab6d485b89858f9dd462bf7ebfe5aba8dd889ccfb3f`.
- **Evidence boundary:** Builder reports five focused mocked tests PASS. No Podman, container, database, fixture, scanner, gate, or other test ran. This is command-boundary evidence only and does not establish real cleanup or independent runtime acceptance.
- **Run is terminal and closed.** Any further phase or patch uses a fresh Run, Participant, Session, and Run path.
