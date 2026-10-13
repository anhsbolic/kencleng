# T2 runtime harness — independent Code Review

Phase: Code Review. Author/role: Codex / independent Reviewer. Created/Updated: 2026-10-13 Asia/Jakarta. Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-001`. Participant / Session: `PARTICIPANT-C1-ENG-T2-RUNTIME-HARNESS-REVIEWER-001` / `SESSION-C1-ENG-T2-RUNTIME-HARNESS-REVIEWER-001` (assignment identities, not independently attested thread IDs). Selected model / reasoning: `gpt-6-sol` / `medium` per approved Invocation; runtime model/effort is not independently exposed. Target HEAD / branch: `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering`. Workflow revision: `651ce57f044bff23fba99082322cee6c066d25a6`.

Scope/provenance: read the live three-file diff against HEAD and full affected source, Approved Techplan/T2, relevant OpenAPI auth contract, root/backend instructions, Review prompt/guidelines/checklist, orchestrated overlay/Run contract, bounded Build scope/report/identities, runtime guide, and routed Harscode input-validation, privilege-separation, and integration/concurrency guidance. All 30 `input-bindings.json` entries matched live SHA-256; Review scope hash `e886adc4cd09b8ecadf5cd00937fe1790ff6806ded5f738dacc689ae75b10277` and bindings hash `0d43733690b9f09ffc7dd6300dfe26fcb91a9836b334b28af53f0a0e23ed313b` match Invocation. HEAD preimages and live postimages of all three files match Build `file-identities.json`. Invocation prose mistypes the Build-report SHA-256 (`...dc1708be...`); bound JSON and live report agree on `c43a030248673bd094aef4352758f4d0007bb14bf255c6202dc17008be194f34`. This is a provenance typo, not changed input content. No binding or upstream artifact was edited.

## 1. Safety

**F1 — Blocking — fixture may survive a failed creation command.** Location: `backend/scripts/auth_harness.py:69-74,121-124`. `created` becomes true only after `podman run` returns successfully. If that command creates the container and then times out or returns a tool failure before the ID is received, control enters `finally` with `created == False`; the cleanup branch is skipped. This control-flow gap predates the three-file delta, but becomes material to the new claim that the reusable, Run-owned fixture is cleaned on failure. The unique name and labels can still identify the resource, but their existence is never checked after the failed command. An orphaned database container/anonymous volume can retain fixture credentials and create resource/evidence ambiguity for later Runs. In a new Build/Patch Run, handle ambiguous creation completion by discovering the exact generated name, verifying its Run/owner identity and immutable ID before any deletion, then removing only that owned resource; if identity or cleanup cannot be established, fail and report the unresolved resource. Keep the existing fail-closed ownership rule and bounded operations. The Build's successful cleanup does not exercise this path.

Other inspected safety boundaries: arguments are allowlisted before resource commands; subprocesses use argv without a shell; the active ID is passed into labels and Go child environment; normal-path inspection compares Run label, owner label, and returned container ID; the six-case and logout selectors are fixed; fixture DSNs are generated and passed to child environment; the runner filters the named application/provider credential variables. No further safety finding from static inspection. `harness_test.go` validates the supplied marker's grammar and DSN shape; it does not independently prove a Run authorization or ownership, so Go test output alone must not be treated as such evidence.

## 2. Quality

No findings. The suite mapping and ownership helper make the bounded selection/checks legible; failure output remains explicit. The metadata typo above belongs to Invocation provenance and is routed to Orchestrator, not a source-quality patch.

## 3. Stack-Specific Best Practices

F1 also conflicts with the deterministic cleanup path for acquired test resources in `../harscode-workspace/best-practices/go/integration-testing-setup.md`; the failure occurs at the Python/Podman creation boundary that supplies the Go integration suite. The same guidance calls for bounded execution and post-abort resource confirmation. Input validation and shell-free argv follow `../harscode-workspace/best-practices/go/input-validation-and-injection.md`. Restricted runtime grants and actual DB contention checks are consistent with the applicable privilege and concurrency concerns in `go/role-and-privilege-separation.md` and `go/testing-concurrency.md`; the latter's runtime verification remains Testing-owned. No additional stack-specific finding.

## 4. Consistency

F1 leaves the bounded Build scope's explicit success-or-failure cleanup requirement unmet (`RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001/scope.md`, “Exact disposable fixture boundary”). It also leaves the `backend/README.md` cleanup statement and active Review scope's cleanup claim stronger than this failure path supports. The README accurately distinguishes real PostgreSQL/TCP/SCRAM behavior, `httptest`, simulated callback, Build executability, and unverified provider/browser/HTTPS/independent acceptance. Existing Go test assertions are unchanged by the marker/import delta. No other consistency finding. The Techplan Test Focus Pointer already identifies auth and concurrency sensitivity; F1 does not require a separate Techplan-drift finding.

## Verification executed during Review

Runtime verification: **none**. Read-only SHA-256 comparison of the 30 bound inputs, Review scope/bindings, three HEAD preimages and live postimages; static diff/source inspection only. No harness, test, scanner, Make, Podman, container, DB, provider, browser, or HTTPS operation was run. F1 is a reasoned failure-path finding, not a reproduced resource leak.

## Verdict

**Request changes** for the bounded runtime-harness claim. F1 is blocking. This verdict does not assess production authentication, independent runtime acceptance, T2 completion, T2/T3 readiness, or closure of `B-T2-002` / `B-T2-003`. See `evidence/patch-plan-1.md` for the bounded source repair. Preserve prior terminal Build/Testing evidence unchanged.

## Phase handoff

- **Outcome:** COMPLETED — independent four-pass Review; Request changes for bounded harness claim.
- **Result refs:** `evidence/review-findings-1.md`; `evidence/patch-plan-1.md`; exact target identities in the bound Build `evidence/file-identities.json`.
- **Findings:** F1 blocking cleanup gap after ambiguous `podman run` failure; Invocation Build-report hash typo is nonblocking metadata discrepancy.
- **Decision requests:** Orchestrator reconcile F1 and route a fresh Build/Patch Run; reconcile provenance typo outside this Reviewer write scope if useful. No Product decision identified.
- **Blockers:** F1 blocks approving this harness cleanup claim and relying on it for a new independent Testing fixture; unrelated work is unaffected.
- **Open / unverified:** abnormal creation/cleanup paths and independent runtime/command-gate acceptance remain untested; provider/browser/HTTPS and wider T2 obligations remain outside this Review.
- **Recommended continuation:** new Build/Patch Run with a fresh Participant/Session to implement the patch plan; then re-ground the changed path and route independent Testing with its own approved fixture and applicable post-delta gates. Orchestrator owns the route.
- **Context refs:** this Run's Invocation/Scope/bindings; Approved Techplan/T2; F1 and patch plan; exact three-file delta/Build identities.
