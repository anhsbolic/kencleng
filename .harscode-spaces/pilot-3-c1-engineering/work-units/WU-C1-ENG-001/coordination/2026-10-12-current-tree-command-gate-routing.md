# Current-tree backend command gate routing — 2026-10-12

## Applicability

The logout patch's narrowly affected claim is already supported by the terminal Build report and independent exact-delta Review. The Review verdict is Approve with minor comments, zero blocking, one nonblocking note that the null oracle does not explicitly assert absence of `Set-Cookie`. No patch loop is needed for that Finding closure.

The independent backend command gate remains separately Testing-owned. `RUN-C1-ENG-TEST-T2-BACKEND-GATE-001` completed PASS for its bound tree, but that identity predates the logout `http.go` change and the updated backend README/harness inputs. Its PASS is not inherited by the changed tree. This is a meaningful fresh gate claim because the required `make verify` and `go mod verify` results must apply to the current assignment-bound tree.

## Prepared route

`RUN-C1-ENG-TEST-T2-BACKEND-GATE-002` completed its bounded current-tree gate. Terminal report SHA-256 `a6df367c7d406b041d960c688a08e0cce5c522228532cb8c2bfe6cf32a64f312` matches the durable report, has one Phase handoff, and records 43/43 bindings before/after. `make verify` and `go mod verify` each exited 0. Gosec's 3 and Gitleaks history's 4 raw findings remain accepted dispositions; Govulncheck reported no vulnerabilities, and module integrity verified. This is a command-gate PASS only; cache integrity is not full provenance or dependency acceptance.

The terminal `RUN-C1-ENG-TEST-T2-001` STALLED report remains preserved. A later READY Invocation at the same path does not reopen the occurrence and cannot be dispatched or reused.

## Remaining runtime route

This command gate does not close the independent HTTP/DB/runtime evidence still owned by T2. The current `backend/scripts/auth_harness.py` and `harness_test.go` require the fixture ownership marker `RUN-C1-ENG-BUILD-T2-LOGOUT-PATCH-001`, and the runner selects only its two logout HTTP tests. They cannot establish a fresh Testing-owned fixture or the broader real SCRAM/concurrency obligations. Any independent runtime pass therefore needs a distinct, explicitly authorized run-owned fixture/harness route with fresh resource ownership; it must not reuse the logout Build Run identity. Google and integrated FE/UI checks remain deferred under their separate holds.

The current runner's Run marker and selector make a separate Build re-entry necessary before independent Testing can own a broader disposable fixture. Prepared [`RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001`](../runs/RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001/invocation.md) to make the marker Run-scoped, allow only bounded suite selectors, and execute all six existing harness tests against its own disposable restricted-PostgreSQL fixture. Its exact model and new fixture approval are pending; it has not been dispatched. This is test-infrastructure Build scope, not a logout patch loop. If it completes, review applicability will be decided from the actual delta; a fresh Testing Run can then own its own fixture and current-tree command gate. Keep `B-T2-002`, `B-T2-003`, historical `B-T2-001`, and all T2/T3 readiness/acceptance boundaries unchanged.
