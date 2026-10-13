# Testing Report — RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-002

- Phase: Independent Testing
- Author: Tester — backend auth harness runtime
- Created: 2026-10-13
- Model / Reasoning: `gpt-6-luna` / medium (Invocation)
- Session: `SESSION-C1-ENG-T2-CONFIGURED-DB-HARNESS-TESTER-002`
- Target revision: `79ec63d57674d1e4976ff80b2dfee8c50bf17934`
- Workflow revision: `651ce57f044bff23fba99082322cee6c066d25a6`

## 0. Sweep Summary

- Confirmed: focused mocked harness regression → 6 tests passed.
- Closed from prior gap: configured `DATABASE_URL` → production `db.Open` path and selected session queries/operations verified against this Run's restricted disposable PostgreSQL fixture (R14; see claim below).
- Closed from prior gap: expiry-under-lock oracle observed against disposable PostgreSQL fixture and authority denial asserted (B-T2-003; separate claim below).
- Still requires fresh Testing / Human: deployed application configuration and real Google/provider/browser login remain deferred. This Run does not establish T2 acceptance, B-T2-002 closure, or T2/T3 readiness.

## 0a. Test Focus Pointer Execution

| Area | Evidence anchor opened | Specialized verification | Result |
|---|---|---|---|
| PostgreSQL session lock/expiry | Techplan Test Focus anchor and configured DB prerequisite | `TestHarnessSessionLockContention/expiry` against real PostgreSQL blocking edge | PASS; see B-T2-003 |

## 1. Test Coverage

| Rule / scenario | Category | Observable verification | Result |
|---|---|---|---|
| R14 — configured pool/session queries | Fresh independent runtime | `TestHarnessRuntimePrivileges` and remaining selected `t2` cases. Runner supplied the restricted fixture's `DATABASE_URL` only to the Go child; the test path uses production `db.Open`. Runtime identity authenticated over loopback TCP with SCRAM; invalid password returned `28P01`; no role memberships; exact nine table privileges and prohibited DDL/role/owner/delete operations checked. PostgreSQL 16.14. | PASS — R14 configured-pool/session evidence established for this disposable fixture. |
| B-T2-003 — stored session expiry while row is locked | Fresh independent concurrency oracle | `TestHarnessSessionLockContention/expiry` observed PostgreSQL blocking, waited for expiry, released the lock within bounds, then asserted the session no longer grants authority. | PASS — B-T2-003 established for this fixture, independently of the R14 verdict. |
| HTTP mutation/logout behavior | `httptest` | `TestHarnessHTTPMeMutationLogout` including anonymous, malformed, expired, revoked, origin, type and CSRF cases. | PASS — HTTP `httptest`; not a deployed listener. |
| Callback behavior | `httptest` + simulated Provider | `TestHarnessHTTPSimulatedCallback` success and negative cases (browser mismatch, provider error, protocol/state/code failures). | PASS — Provider is simulated; no Google login occurred. |
| Concurrent person upsert and null logout | Real fixture / application boundary | `TestHarnessConcurrentPerson` witnessed concurrent runtime upserts and the first-login blocker edge; `TestHarnessLogoutRejectsNull` passed. | PASS. |

## 2. Error Verification

| Error case | Expected behavior/category | Actual | Actionable/propagated correctly? |
|---|---|---|---|
| Invalid runtime password | PostgreSQL authentication rejection | `28P01` observed by focused runtime privilege test. | Yes; rejected. |
| Disallowed runtime privileges | Permission denial for DDL, role, ownership, and prohibited delete operations | All asserted denials observed. | Yes; least-privilege boundary held. |
| HTTP malformed/unauthorized inputs | Caller-facing rejection for invalid session, origin, type, or CSRF | Relevant `httptest` subcases passed. | Yes within `httptest` scope. |
| Lock-time expiry | Session authority denied after expiry while waiting on locked row | Real PostgreSQL blocking edge and post-release assertion passed. | Yes; stale authority was not accepted. |

## 3. Final Verification

- Targeted commands run from `backend/`:
  - `python3 -m unittest discover -s scripts -p test_auth_harness.py` — PASS, 6 tests.
  - `python3 scripts/auth_harness.py --run-id RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-002 --suite t2` — PASS, all six selected top-level cases and their subtests.
- The first runner attempt stopped before fixture creation because sandboxed Podman could not access `/run/user/1000/libpod`; the identical bounded runner command was rerun with approved host access and completed. No test or fixture command beyond the two scoped commands was run.
- Fixture classification: PostgreSQL 16.14/TCP/SCRAM, privileges, and lock/expiry are real disposable-fixture evidence. HTTP cases use `httptest`. Callback Provider is simulated. Runner verified ownership before/after provisioning and successful-path removal of the owned container and anonymous data volumes.
- Broad checks intentionally not rerun: `make verify`, scanner/dependency gates, and broad suites are outside this Run's two declared runtime conditions.
- Migration/schema collision: N/A — this bounded run applies the existing identity migration to a fresh isolated fixture; no shared schema was accessed.
- Backward compatibility: N/A — no production change is part of this Testing Run.
- Broader-suite requirement for cross-cutting change: N/A — no production change is part of this Testing Run.
- Fresh Techplan consistency read: no contradiction found for the two assigned oracles; neither result implies whole-T2 acceptance.
- Cleanup: runner reported `cleanup verified: owned container and anonymous data volumes removed`.

## 4. New Recurring Bug Patterns

None identified.

## Verdict

Pass for the two bounded claims only: R14 configured-pool/session evidence and B-T2-003 expiry-under-lock. They are independent results. This verdict does not establish real Google/provider login, deployed configuration, B-T2-002 closure, whole-T2 acceptance, or T2/T3 readiness.

## Phase handoff

- Completed: focused unittest and six selected `t2` cases; PASS for R14 and B-T2-003 as separate claims.
- Artifacts: this testing report.
- Human decision: owning acceptance/commit decision remains needed for B-T2-002; no decision is implied by this evidence.
- Open / deferred: deployed configuration and real Google/provider/browser verification; whole-T2 acceptance and T2/T3 readiness remain unestablished.
- Recommended next step: Orchestrator reconciles these two results with owning conditions; do not infer acceptance from this report alone.
- Session transition: any subsequent orchestrated phase uses a fresh Run/Participant Session as its invocation requires.
- Context pointers: current Techplan, this testing report, configured-pool Build/Patch report, and terminal credential-patch Review report named in Scope.
