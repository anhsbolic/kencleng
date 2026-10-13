# Scope — RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-001

Status: PREPARED — execution is gated on terminal Review reconciliation, exact fixture authorization for this Run, and Human mechanical dispatch. No fixture or Testing action has occurred.

## Separate claims

This fresh independent Testing Run will report two claims separately, even though one fixture and one `t2` command may provide evidence for both:

1. **R14 configured-pool/session evidence:** independently exercise the restricted disposable PostgreSQL pool through the test child's `DATABASE_URL` and production `db.Open`, verify real PostgreSQL/TCP/SCRAM authentication, the existing restricted privilege assertions, and observable session queries/operations after the selected dependency delta.
2. **B-T2-003 post-lock-expiry oracle:** independently observe the real PostgreSQL blocking edge, wait past the stored session-expiry deadline while the session row remains locked, then verify the under-lock revalidation returns no session authority. Report the exact oracle and result separately from R14; a failure or incomplete assertion for one claim must not be hidden by the other claim's result.

The run is bounded to these harness claims and successful-path ownership cleanup. It is not whole-T2 acceptance and cannot close B-T2-002 without the separate owning acceptance/commit decision.

## Exact target and current evidence

- Build/Patch terminal report: `../RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001/evidence/patch-report-1.md` (SHA-256 `33aa73d184433acdf296703cb79089518436d5187169bb5528374c462e7d8de3`; one handoff).
- Exact source target: `backend/internal/platform/auth/harness_test.go`, `backend/scripts/auth_harness.py`, and `backend/README.md` as listed in `input-bindings.json`.
- Required predecessor: targeted Review `RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-PATCH-001` must be terminal and reconciled. Before dispatch, re-ground on its report and confirm the exact three-file target still matches; if a blocking finding requires code change or changes the target, stop and supersede/rebind this prepared assignment rather than testing stale code.
- Prior D8 dependency Build and exact-delta dependency Review remain bounded evidence on their unchanged inputs; do not repeat graph/provenance/scanner gates here.

## Fixture and command

After Review reconciliation, exact fixture authorization for this Testing Run, and Human mechanical dispatch, create one fresh Run-owned disposable PostgreSQL 16 fixture using only the existing locally available image, existing restricted `c1_runtime` role/grants, loopback ephemeral TCP, anonymous storage, generated private credentials, and ownership-checked cleanup. Provisioning/witness identity remains separate from the restricted runtime identity. This authorization does not transfer from any Build or prior Testing Run.

Execute only from `backend/`:

```text
python3 scripts/auth_harness.py --run-id RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-001 --suite t2
```

The runner must remove inherited application/provider/test DSNs before fixture setup and set `DATABASE_URL` only in the Go test child. Never read `.env`, use a shared/persistent DB, print/store DSNs or credentials, expand role/grants/resources, run real Google/provider login, or use a deployed listener/browser. Preserve the dirty workspace and all historical evidence.

The six selected cases are the existing `t2` suite: runtime privileges; HTTP mutation/logout; simulated callback; concurrent person; session lock contention; null logout. Capture R14 and B-T2-003 outcomes separately. Classify PostgreSQL/TCP/SCRAM and lock contention as real disposable-fixture evidence, HTTP as `httptest`, Provider as simulated, and cleanup as runner-evidenced successful-path ownership cleanup. Abnormal cleanup paths and deployed configuration remain unestablished.

## Not in this claim

No `make verify`, scanner, dependency verification, command gate, broad/full T2 suite, Google login, browser, HTTPS/deployed listener, shared DB, `.env`, production source edit, dependency change, commit, acceptance, residual-risk decision, hold release, T2/T3 readiness, or T3 dispatch. Google login and integrated UI/cache/Human walkthrough remain deferred under their existing holds; backend DB obligations above remain active.

## Output

Write only this Run's `evidence/testing-report-1.md`, following the current canonical Testing report contract and containing exactly one `## Phase handoff`. If the Review is terminal without source-changing findings, bind its report before dispatch. Any patch request returns to a new Build/Patch Run; do not continue this Testing occurrence against a changed target.
