# Scope — RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-CLEANUP-PATCH-001

Status: PREPARED. No source/test edit or command execution has occurred in this Run. No real container or database fixture is authorized or required.

## Claim and accepted patch target

Implement the bounded F1 repair from terminal Review `RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-001`: if `podman run` creates the uniquely named container but times out or returns a failure before its container ID reaches the harness, the runner must reconcile that exact creation attempt and clean only a resource whose Run label, owner label, name, and immutable container ID can be verified. If no resource exists, preserve/report the creation failure. If existence, ownership, identity, or cleanup cannot be established, fail closed with a clear secret-safe unresolved-resource report; never delete by broad name/prefix or unverified name.

This patch is required because the terminal Review returned Request changes with blocking F1. It repairs test-harness cleanup only; it does not alter production auth/session behavior or establish runtime acceptance.

## Allowed writes

- `backend/scripts/auth_harness.py` — reconcile an ambiguous create-command outcome using only the exact random fixture name generated before the command; verify exact Run/owner labels and immutable ID before removal; preserve bounded cleanup and fail-closed behavior.
- `backend/scripts/test_auth_harness.py` — add focused stdlib-only command-boundary tests using mocked subprocess/Podman responses. Cover creation followed by timeout/nonzero result, cleanup of an exact verified owned resource, and refusal to delete an absent/mismatched/uncertain resource.
- `backend/README.md` — optional, only if the implemented failure-path semantics or evidence limits make a documentation adjustment necessary.
- This Run's own Invocation-referenced terminal patch report/evidence only.

Do not edit production auth/store/db/server code, existing Go assertions, migration/API/dependencies, scanner policy, Makefile/Compose topology, Product/Techplan/task authority, another Run's evidence, or unrelated dirty files. Preserve all existing state and historical reports.

## Focused Build verification

The test must inject an ambiguous `podman run` result at the command boundary without launching Podman, creating a container, accessing a database, or needing real credentials. Demonstrate that an exact owned resource is removed by the inspected immutable ID with anonymous-volume cleanup, while absent or mismatched ownership causes no deletion and a safe report. Retain the original test assertion behavior and report any unresolved edge.

Run only the focused test from `backend/`:

```text
python3 scripts/test_auth_harness.py
```

Focused Python syntax/format checks are allowed. No broad `make verify`, six-case T2 fixture harness, Podman/container/DB/provider/browser/runtime, scanner, network, or command-gate execution is authorized in this Build/Patch Run. Do not claim Testing acceptance. A subsequent independent Review will verify F1 closure; a later independent Testing Run must own its own authorized fixture and applicable post-delta final gates.

## Report and terminal requirements

Write `evidence/patch-report-1.md` with exact delta, focused command-boundary scenarios and outcomes, safe failure/cleanup evidence, contract check, deferred gaps, and exactly one `## Phase handoff`. Do not persist DSNs, credentials, container environment, sensitive stderr, or unredacted subprocess payloads. The report must state that no real Podman/container/database fixture was created and that this evidence tests command-boundary cleanup behavior only.
