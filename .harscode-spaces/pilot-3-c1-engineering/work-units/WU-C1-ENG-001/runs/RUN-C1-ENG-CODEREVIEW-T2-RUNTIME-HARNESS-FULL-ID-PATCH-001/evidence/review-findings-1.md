# Independent Code Review — Full-ID cleanup guard

Phase: Code Review. Author / role: `PARTICIPANT-C1-ENG-T2-RUNTIME-HARNESS-FULL-ID-PATCH-REVIEWER-001` / independent Reviewer. Created: 2026-10-13 (Asia/Jakarta). Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-FULL-ID-PATCH-001`. Session: `SESSION-C1-ENG-T2-RUNTIME-HARNESS-FULL-ID-PATCH-REVIEWER-001` (Invocation assignment identity). Selected model / reasoning: `gpt-6-luna` / `medium` per Invocation; runtime identity is not independently exposed. Target revision / branch: `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering`. Workflow revision: `651ce57f044bff23fba99082322cee6c066d25a6`.

Provenance: Recomputed all 29 assignment-defining bindings; all match. Scoped current SHA-256 values match the terminal Build report: `backend/scripts/auth_harness.py` `932a427e7917f73f2917ea316dbe62d29e3ffca649ebc5dc361db646319fe79a`; `backend/scripts/test_auth_harness.py` `4854fd949a3f09ddac611ab6d485b89858f9dd462bf7ebfe5aba8dd889ccfb3f`. These are the only two files in this review. Inspected their current contents/delta against the immediately prior bound identities, prior F1 finding/plan, Build scope/report, current Techplan/T2 task, root/backend instructions, canonical Review prompt/guidelines/checklist and orchestrated overlay, plus routed resource-cleanup/integration guidance. No out-of-scope changes were considered. Build's five passing mocked tests are Builder-reported evidence and were not rerun.

## 1. Safety

No findings. Statically traced the ambiguous-create and normal cleanup paths to `cleanup_fixture`. A discovered ID must match `[0-9a-fA-F]{12,64}`. If supplied, `expected_id` must match `[0-9a-fA-F]{64}` and begin with the discovered ID. The inspected `Id` must be a string matching `[0-9a-fA-F]{64}`, begin with the discovered ID, and equal `expected_id` when supplied. Ownership/name checks also complete before the single `podman rm` call. Validation and identity failures are converted to the bounded `unresolved fixture resource` error path. No deletion is reached on a validation failure.

The new mocked test covers both a short inspected ID and a 64-character non-hex inspected ID. For each subtest, the mock permits only `ps` and `inspect`, expects the unresolved-resource error, and asserts no recorded command has operation `rm`. Existing timeout-after-create, nonzero-after-create, absent-resource, and mismatched-owner cases remain in the file. The timeout and nonzero cases assert cleanup uses the verified full ID; absent and mismatched cases preserve fail-closed behavior.

## 2. Quality

No findings. The checks are localized at the cleanup boundary, the validation conditions are explicit, and the test keeps invalid-ID cases together while isolating mock calls per subtest. No unnecessary abstraction or unrelated behavior change was found in the scoped delta.

## 3. Stack-Specific Best Practices

No findings. The relevant routed resource-lifetime guidance, `../harscode-workspace/best-practices/go/integration-testing-setup.md`, calls for deterministic cleanup and explicit failure visibility for disposable integration resources; the reviewed validation keeps deletion behind verified identity and the negative mock cases assert the no-delete boundary. The global best-practices clue map has no narrower Python/container-cleanup entry matching this delta. This is harness cleanup review; it does not claim real integration evidence.

## 4. Consistency

No findings. The root and `backend/AGENTS.md` rules require intentional handling of expected failures and distinguish fixtures from real-world evidence. `backend/README.md` describes this harness as using ownership-checked container/volume cleanup and says to stop if isolation or owned cleanup cannot be established. The scoped guard and sanitized unresolved-resource path fit those requirements. The patch remains limited to the two authorized files; no production authentication or upstream authority changed.

## Verification executed during Review

None. Review used read-only binding/hash reconciliation and static source/test/diff inspection. No test, gate, scanner, Podman, container, database, provider, browser, or runtime action was run. Builder-reported mock results remain unverified by this Reviewer.

## Verdict

**Approve** for the bounded full-ID command-boundary guard and its mock tests. Static review confirms discovered, inspected, and optional expected IDs are validated and related before removal, and the short/malformed inspect-ID tests assert that `podman rm` is not called. This does not establish real Podman cleanup, database/HTTP runtime acceptance, provider/browser/HTTPS evidence, closure of `B-T2-002` / `B-T2-003`, T2/T3 readiness, or C1 completion.

## Phase handoff

- **Outcome:** COMPLETED — four-pass independent Review; Approve for the scoped full-ID cleanup guard.
- **Result refs:** This report; exact source/test identities in the current Run Invocation and terminal Build report.
- **Findings:** None in the scoped delta.
- **Decision requests:** None.
- **Blockers:** None for this scoped Review claim.
- **Open / unverified:** Mock tests were not rerun. Real Podman/container/volume cleanup, database and HTTP runtime behavior, provider/browser/HTTPS evidence, and wider T2 obligations remain unverified/out of scope.
- **Recommended continuation:** Orchestrator evaluates any Testing route from the exact claim and remaining evidence gap; Review chronology alone does not require Testing. Any runtime fixture/gates require their own applicable Run ownership and evidence.
- **Context refs:** This Run's Invocation, scope and input bindings; Approved Techplan/T2 task; terminal full-ID Build report; two-file delta.
