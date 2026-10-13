# Build/Patch report — full container ID validation

Run: `RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-FULL-ID-PATCH-001`  
Scope: F1 only, as bound by this Run's Invocation and `scope.md`. Assignment bindings: **27/27 matched before edits**.

## Delta

- `backend/scripts/auth_harness.py`: cleanup validates the discovered Podman ID as 12–64 hexadecimal characters; validates a supplied `expected_id` as a full 64-character hexadecimal ID; requires inspect `Id` to be a full 64-character hexadecimal ID, to begin with the discovered ID, and to equal `expected_id` when supplied. All checks precede the removal call. Invalid identities fail through the existing sanitized unresolved-resource error path.
- `backend/scripts/test_auth_harness.py`: added mocked short and malformed inspected-ID cases; both assert that `podman rm` is never called. Existing timeout-after-create, nonzero-after-create, absent-resource, and mismatched-owner cases remain.

## Verification

Ran only `python3 scripts/test_auth_harness.py` from `backend/`: **5 tests passed** (exit 0). The suite uses mocked subprocess responses. No Podman, container, database, fixture, scanner, gate, or other test was run. This is mock-only command-boundary evidence and does not establish real resource deletion, volume cleanup, or independent runtime acceptance.

Post-edit SHA-256: `auth_harness.py` `932a427e7917f73f2917ea316dbe62d29e3ffca649ebc5dc361db646319fe79a`; `test_auth_harness.py` `4854fd949a3f09ddac611ab6d485b89858f9dd462bf7ebfe5aba8dd889ccfb3f`.

## Phase handoff

F1's ID-validation gap is addressed in the bounded cleanup path. Return this patch to the Orchestrator for delta reconciliation and targeted Review routing. Evidence remains limited to mocked command-boundary tests; real Podman/runtime acceptance is not claimed.
