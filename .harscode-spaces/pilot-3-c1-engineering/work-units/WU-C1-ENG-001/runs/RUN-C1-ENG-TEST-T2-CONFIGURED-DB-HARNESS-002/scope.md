# Scope — RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-002

Status: PREPARED — the Human instruction to prepare independent Testing with a fixture owned by the new Run authorizes the single bounded fixture below. Human mechanical dispatch remains pending. No Testing command or fixture action has occurred.

## Separate claims

This independent Testing Run evaluates two distinct claims using one fresh, Run-owned disposable PostgreSQL fixture for the runtime command:

1. **R14 configured-pool/session evidence:** verify the restricted fixture child's `DATABASE_URL` reaches production `db.Open`; establish real PostgreSQL/TCP/SCRAM authentication, restricted privileges, and the selected session queries/operations after the dependency delta.
2. **B-T2-003 post-lock-expiry oracle:** observe the real lock wait, let the stored session expire while its row is locked, then verify under-lock revalidation denies session authority. Record this result separately from R14.

Neither claim implies the other. This Run does not establish whole-T2 acceptance or B-T2-002 closure; B-T2-002 still requires R14 reconciliation plus its owning acceptance/commit.

## Exact target and predecessor evidence

- Configured-pool Build/Patch report: `../RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001/evidence/patch-report-1.md`, SHA-256 `33aa73d184433acdf296703cb79089518436d5187169bb5528374c462e7d8de3`.
- Credential-isolation Build/Patch report: `../RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001/evidence/patch-report-1.md`, SHA-256 `d06aadf7f8611f728915c64ece30bab317136f76ab556831271429c9aa1e6768`.
- Terminal F1 Review: `../RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001/evidence/review-findings-1.md`, SHA-256 `28f0179e20e6a65c57d6cec82d34bd16af2da837347df4f8b27119778660fa3e`; exactly one handoff; Approve for F1 on the exact two-file patch by static inspection only.
- Earlier configured-pool Review: `../RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-PATCH-001/evidence/review-findings-1.md` and patch plan remain historical context. Its Request changes finding was closed only by the exact patch reviewed above.
- Exact current target includes the configured-pool harness/README delta and the two credential-isolation files; live hashes are in this Run's `input-bindings.json`.

The prior prepared Testing Run `...-001` is stale for dispatch: its bound `backend/scripts/auth_harness.py` hash predates the credential patch and it lacks the terminal F1 Review binding. It is superseded by this fresh Run/Participant/Session; no command was run under the prior assignment.

## Fixture and bounded commands

The Human explicitly directed independent Testing using a fixture owned by the new Run. That instruction authorizes one fresh disposable PostgreSQL 16 fixture within this exact boundary: existing locally available image, existing restricted runtime role/grants, loopback ephemeral TCP, anonymous storage, private generated credentials, separate provisioner/witness identities, and ownership-checked cleanup. No authorization from another Run is reused. Create the fixture only after Human mechanical dispatch.

Run both commands from `backend/`:

```text
python3 -m unittest discover -s scripts -p test_auth_harness.py
python3 scripts/auth_harness.py --run-id RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-002 --suite t2
```

The first is focused mocked regression evidence for F1; the second owns the fixture and the six selected `t2` cases. The runner must sanitize the `go version` metadata environment and pass the configured fixture URL only to the Go test child. Never read `.env`, access a shared/persistent DB, print/store DSNs or credentials, widen role/grants/resources, or run real Google/provider login, browser, or deployed listener.

Classify PostgreSQL/TCP/SCRAM, privilege checks, and lock/expiry behavior as real disposable-fixture evidence; HTTP as `httptest`; callback Provider as simulated; cleanup as runner-evidenced successful-path ownership cleanup. Abnormal cleanup paths, deployed configuration, and real Google login remain unestablished/deferred.

This bounded claim does not run `make verify`, dependency/scanner gates, or broad T2 suites: those are outside the two declared Testing-owned runtime conditions and require their own current-tree reconciliation. This Run cannot claim T2/T3 readiness or acceptance.

## Output

Write only this Run's `evidence/testing-report-1.md`, following current Testing report and handoff contracts, with exactly one `## Phase handoff`. Report R14 and B-T2-003 separately. If either oracle is incomplete or fails, report that result without converting the other claim into a pass. No production source or authority edits.
