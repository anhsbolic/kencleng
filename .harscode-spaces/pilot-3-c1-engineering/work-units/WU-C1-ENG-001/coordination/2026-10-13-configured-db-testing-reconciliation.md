# Configured DB independent Testing reconciliation — 2026-10-13

## Terminal report and assignment integrity

`RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-002/evidence/testing-report-1.md` matches SHA-256 `31505a77243c8eec0a92df1a653c0d8eb14f7a29a8aeb99094476011a0a8a06a`, contains exactly one `## Phase handoff`, and all 36 assignment bindings still match. The report verdict is Pass only for its two declared claims. No source delta is reported or bound drift found; a new Code Review is not applicable to this Testing occurrence. The already completed exact-delta Review remains the F1 verdict.

Tester reports the focused mocked regression suite passed 6/6, then all six selected `t2` cases passed on this Run's disposable PostgreSQL fixture. The first identical runner attempt stopped before fixture creation due to sandboxed Podman access; the retry used approved host access. This is environment recovery, not an implementation failure or permission/fixture expansion. The report records no DSN or credential.

## Separate verification outcomes

- **R14 configured-pool/session queries — PASS, independent Testing evidence.** PostgreSQL 16.14 authenticated over TCP/SCRAM as the restricted runtime role. The fixture URL reached the Go test child through the `DATABASE_URL` path and production `db.Open`; role/privilege assertions and selected session operations passed. This satisfies the declared Testing-owned configured-pool/session-query condition for the bounded fixture.
- **B-T2-003 expiry-under-lock oracle — PASS, independently reconciled.** The expiry subtest observed an actual PostgreSQL lock wait, allowed the stored session to expire while blocked, and confirmed the expired session did not grant authority after lock release. **B-T2-003 is CLOSED for this exact post-lock-expiry oracle.** This does not establish other T2 conditions.
- **HTTP:** `httptest`, not a deployed listener. **Callback Provider:** simulated, no Google exchange. **Cleanup:** runner verified successful-path removal of this Run's owned container and anonymous volumes. Abnormal cleanup, deployed configuration, real Google login, and browser/integrated UI evidence remain unestablished or deferred.

The initial configuration and target were verified from report/bindings, but `make verify`, broad suites, dependency/scanner gates, and shared/deployed runtime were outside this Run and were not run. No overall T2 acceptance or readiness follows.

## B-T2-002 disposition and next route

Independent R14 evidence is now present for the configured-pool/session-query condition. `B-T2-002` remains **OPEN** because its owning acceptance/commit decision remains outstanding. The previously approved dependency candidate did not authorize a dependency commit or accept residual risk. Do not commit, close B-T2-002, claim T2/T3 readiness, or dispatch T3 on this Testing Pass alone.

Both Human holds remain: real Google login and integrated UI/cache/Human walkthrough are deferred as already authorized. No Product authority changed. Historical terminal reports and dirty workspace are preserved. No Orchestrator test, gate, scanner, fixture, database, runtime, or source action was performed.
