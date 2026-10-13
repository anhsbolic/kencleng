# T2 runtime harness — Review patch plan

Phase: Code Review. Author/role: Codex / independent Reviewer. Created/Updated: 2026-10-13 Asia/Jakarta. Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-001`. Participant / Session: `PARTICIPANT-C1-ENG-T2-RUNTIME-HARNESS-REVIEWER-001` / `SESSION-C1-ENG-T2-RUNTIME-HARNESS-REVIEWER-001`. Selected model / reasoning: `gpt-6-sol` / `medium` per approved Invocation; runtime model/effort not independently exposed. Target / workflow revisions: `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `651ce57f044bff23fba99082322cee6c066d25a6`.

Source finding: F1 in `review-findings-1.md`. A failed or timed-out `podman run` can leave a newly created container while `created` remains false, bypassing `finally` cleanup.

For a new, separately authorized Build/Patch Run:

1. Update `backend/scripts/auth_harness.py` so ambiguous `podman run` completion enters bounded reconciliation of the exact freshly generated name. If a resource exists, establish its exact Run/owner identity and immutable container ID before removing that ID and its anonymous volumes. Never remove by broad prefix or unverified name. If existence, identity, or cleanup remains uncertain, fail with a clear, secret-safe unresolved-resource report.
2. Preserve pre-resource argument validation, allowlisted selectors, private child credentials, restricted fixture role, owner/ID checks on normal cleanup, and the existing six test assertions. Keep the patch narrow; adjust README only if the final behavior/evidence limits change.
3. Demonstrate the error path with a focused, isolated command-boundary test or equivalent controlled failure injection that represents creation followed by timeout/nonzero return. Prove that verified owned resources are cleaned and mismatched/uncertain resources are never deleted. Normal Build/Testing fixture execution and any broader gates require their own active authorization and evidence.

This plan is Reviewer guidance, not source-change authorization or a runtime acceptance claim. The Orchestrator determines the new Run and verification envelope.
