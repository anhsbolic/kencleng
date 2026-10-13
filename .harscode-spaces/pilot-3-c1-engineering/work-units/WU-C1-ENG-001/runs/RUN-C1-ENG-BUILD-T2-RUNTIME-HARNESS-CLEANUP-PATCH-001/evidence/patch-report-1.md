Phase: Build/Patch. Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-CLEANUP-PATCH-001`. Author: Codex, Implementer. Created: 2026-10-13 Asia/Jakarta. Selected model / reasoning: `gpt-6-luna` / `medium` per Invocation. Participant / Session: `PARTICIPANT-C1-ENG-T2-RUNTIME-HARNESS-CLEANUP-PATCH-IMPLEMENTER-001` / `SESSION-C1-ENG-T2-RUNTIME-HARNESS-CLEANUP-PATCH-IMPLEMENTER-001`. Target revision / workflow revision: `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `651ce57f044bff23fba99082322cee6c066d25a6`.

## What changed

- `backend/scripts/auth_harness.py` now arms cleanup before `podman run`. Cleanup searches only the attempt's generated exact name, requires one matching resource, inspects its immutable full ID and exact name, and checks both `c1.auth.run` and `c1.auth.owner` labels before removal by full ID. Missing resources after an ambiguous create failure retain the original safe failure; uncertain lookup, ownership, identity, or cleanup raises a sanitized unresolved-resource error and does not issue an unverified removal.
- `backend/scripts/test_auth_harness.py` adds four stdlib-only mocked command-boundary cases for timeout-after-create, nonzero-after-create, no resource, and mismatched ownership.

## Tests run

- From `backend/`: `python3 scripts/test_auth_harness.py` — PASS, 4 tests in 0.001s. Mocks intercepted subprocess commands; the create-failure cases verified cleanup used the inspected full ID, the absent-resource case issued no inspect/remove, and the identity-mismatch case issued no removal.

## Verification scope confirmation

No race/concurrency, performance/load, or security-class test was executed in this Build/Patch Run. No broader Testing-owned suite, gate, scanner, or runtime harness was run. No real Podman command, container, database, provider, browser, credential, or fixture was used; evidence is limited to mocked command-boundary cleanup behavior.

## Contract check

- [x] F1 patch target satisfied: ambiguous create attempts enter exact-name reconciliation; removal requires matching Run/owner labels, resource name, and immutable container ID.
- [x] Live-code re-grounding did not invalidate a material contract assumption. Invocation-bound scope, patch plan, and approved T2 task/Techplan identities matched the live files.

## Deferred / not tested here

Real Podman behavior, anonymous-volume removal on a real fixture, and Testing acceptance remain unverified and require their separately authorized Run. No container or database was created.

## Flagged for Techplan / Testing

Independent Review should assess F1 closure. A later Testing Run owns any authorized real-fixture/runtime evidence and applicable post-delta gates.

## Phase handoff

- Outcome: `COMPLETED` — this scoped Build/Patch occurrence completed; no downstream acceptance is implied.
- Result refs: this patch report; `backend/scripts/auth_harness.py`; `backend/scripts/test_auth_harness.py`.
- Findings: F1 command-boundary repair implemented; four mocked cases passed.
- Decision requests: none.
- Blockers: none for this scoped patch.
- Open / unverified: real Podman/container/database cleanup and independent Testing acceptance.
- Recommended continuation: Orchestrator reconciliation, then a fresh targeted independent Review Run for F1.
- Context refs: Invocation and Scope in this Run; accepted `RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-001/evidence/patch-plan-1.md`; approved T2 Techplan/task.
