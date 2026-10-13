Phase: Code Review. Author/role: Codex / independent Reviewer. Created/Updated: 2026-10-13 Asia/Jakarta. Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-CLEANUP-PATCH-001`. Participant / Session: `PARTICIPANT-C1-ENG-T2-RUNTIME-HARNESS-CLEANUP-PATCH-REVIEWER-001` / `SESSION-C1-ENG-T2-RUNTIME-HARNESS-CLEANUP-PATCH-REVIEWER-001`. Selected model / reasoning: `gpt-6-luna` / `medium` per Invocation; runtime model/effort is not independently exposed. Target revision / branch: `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering`. Workflow revision: `651ce57f044bff23fba99082322cee6c066d25a6`.

Provenance: Recomputed all 25 SHA-256 assignment bindings; all match, including this Run's Scope and patch report, accepted F1 finding/plan, Techplan/T2 task, canonical Review authorities, and both scoped files. Live file hashes are `backend/scripts/auth_harness.py` `653cc84c16193d27e3693c4144a1e0bad748a6beff95fc1c7fbe30ec21438605` and `backend/scripts/test_auth_harness.py` `3e78860f863319ef615459746fd68742f418185a980e36f544e5df1bb65cefca`. This is the exact two-file scope; unrelated dirty workspace paths were not considered. Static inspection only. The patch report's four passing mocked tests are Builder-reported evidence, not independently rerun.

## 1. Safety

**F1 — Blocking — cleanup does not require the inspected ID to be a full immutable ID.** Location: `backend/scripts/auth_harness.py`, `cleanup_fixture`, ID checks before `podman rm` (approximately lines 43–69). When `expected_id` is absent after an ambiguous create, the code accepts any nonempty `metadata[0]["Id"]` that starts with the ID returned by exact-name `podman ps`. It does not validate the inspected ID's full form/length. Consequently, a short inspected ID equal to the discovered short ID can pass the prefix check and be passed to `podman rm`; the path does not establish the required full ID before deletion. The four tests use a 64-character fixture ID but do not include a short/malformed inspect-ID rejection case. Require a valid full immutable ID from inspect before removal (and confirm the discovered ID is its prefix); reject malformed/short IDs without deletion, with a focused mocked assertion. This is blocking because deletion is the safety-sensitive operation under review.

Other safety checks pass statically: fixture names use `secrets.token_hex(6)`; an exact-name preflight collision check runs before create; `create_attempted` is set before `podman run`; cleanup searches the exact generated name and rejects zero/multiple list results when an ID was returned; inspection checks exact resource name, Run and owner labels, and expected full ID when available; removal uses the inspected ID. Lookup/inspect/removal errors are surfaced through sanitized unresolved-resource failures, while a genuinely absent resource after ambiguous create returns without deletion. No broad-name deletion path was found. The current absent-resource test checks no inspect/remove; mismatched-owner test checks no remove. The timeout/nonzero tests exercise cleanup helper behavior with mocks, but call `command(run)` and `cleanup_fixture` directly rather than driving `main()`'s pre-armed `finally`; they therefore do not independently assert that integration between the armed flag and cleanup call. Static control flow does show `create_attempted = True` before `podman run` and conditional cleanup in `finally`.

## 2. Quality

No additional findings. The helper has explicit lookup, inspect, and removal failure boundaries; error text does not expose subprocess stderr or private fixture data. The remaining validation gap is captured by F1 rather than repeated here.

## 3. Stack-Specific Best Practices

F1 is also a resource-safety issue under `../harscode-workspace/best-practices/go/integration-testing-setup.md`, “Transaction/resource safety” and “Failure visibility”: acquired integration resources need deterministic cleanup, and abnormal cleanup must fail visibly. The Python/Podman boundary launches the Go integration suite, so this guidance applies to the disposable fixture. Exact-name scoping, bounded commands, and ownership checks otherwise align with the same guidance. No separate finding.

## 4. Consistency

The root `AGENTS.md` requires intentional handling of expected failures and preserving evidence boundaries; `backend/AGENTS.md` requires truthful test/evidence claims. The implementation follows those rules through explicit unresolved-resource errors and safe diagnostics. `backend/README.md` states cleanup is ownership-checked and checks the unique container ID; the reviewed code checks identity and name but has no full-ID format guard, as captured by F1. No separate consistency finding.

## Verification executed during Review

**None.** Recomputed 25 bound SHA-256 identities and performed read-only source, test, Build-report, Review-plan, and applicable-authority inspection. No test, gate, scanner, command reproduction, Podman, container, database, provider, browser, or runtime action was run.

## Verdict

**Request changes** for the bounded F1 cleanup claim. F1 is blocking: require and validate the full inspected immutable container ID before invoking removal, and add a mocked malformed/short-ID no-deletion assertion in the next Build/Patch Run. This verdict does not establish real Podman cleanup, independent runtime acceptance, closure of `B-T2-002` / `B-T2-003`, T2/T3 readiness, or C1 completion. No patch plan is authored here because the accepted F1 patch plan already covers safe ownership/identity verification and this Run's review output is to route the concrete remaining defect back for a fresh Build/Patch Run.

## Phase handoff

- **Completed:** four-pass independent Review; verdict Request changes.
- **Artifacts:** `evidence/review-findings-1.md`.
- **Human decision:** none.
- **Open / deferred:** blocking F1 full-ID validation before deletion; real Podman and independent runtime evidence remain outside this Review.
- **Recommended next step:** Orchestrator route a fresh Build/Patch Run to enforce full-ID validation and add a negative mocked assertion; then reconcile and select further independent verification under a new Run.
- **Session transition:** any patch requires a new Build/Patch Run, Participant, and Session re-grounded on this finding; no source edits or runtime execution occurred here.
- **Context pointers:** this Run's Invocation/Scope; accepted F1 patch plan; `backend/scripts/auth_harness.py` cleanup ID checks and `backend/scripts/test_auth_harness.py` cleanup-boundary cases.
