# Configured-pool Review F1 reconciliation — 2026-10-13

## Terminal Review identity

`RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-PATCH-001/evidence/review-findings-1.md` SHA-256 `e00f134680386cb6566fcc32307b401ccc4529a357d68bcd5a7932900e938f46` matches durable bytes and contains exactly one `## Phase handoff`. Its 21 input bindings match current bytes. The reviewed three-file target still matches the Build report's post-delta identities; Review made no source changes and reports no tests, gates, scanners, fixtures, database operations, Provider flows, or runtime checks.

## Finding and disposition

Review returned Request changes for blocking F1: `backend/scripts/auth_harness.py` passes the fixture-bearing `env` to `go version`, exposing fixture `DATABASE_URL`, provisioner DSN, and other fixture credential material to a subprocess that does not need it. This violates the child-only secret boundary in the configured-pool Build scope. The patch plan `patch-plan-1.md` (SHA-256 `c1fcc7d1c62437e2d86634da5dc7f8b4647ac3e3c13a59dc31bf5833345fc841`) requests a sanitized environment or running version discovery before fixture credentials are added; fixture DSNs must remain available to the intended Go test child only.

F1 is accepted as a scoped Review finding and remains OPEN until patched and independently confirmed. The smallest credible Build/Patch delta is limited to the harness runner and a focused mocked regression in its existing test file. No fixture or runtime is needed for this patch. Build/Patch `RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001` is prepared with fresh Implementer/Session and a focused unit-test command; Human mechanical dispatch is pending.

The previously staged Testing assignment `RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-001` is not dispatchable: its exact target predates this blocking Review finding. Preserve its files, but do not run it. After the new Build terminal handoff, route a targeted Review confirmation for F1 from the actual patch delta; re-ground or supersede/rebind Testing after Review. Any future Testing must use its own newly authorized disposable fixture and keep R14 and B-T2-003 separate. B-T2-002/B-T2-003 remain open; no acceptance/commit, T2/T3 readiness, or T3 dispatch is established.
