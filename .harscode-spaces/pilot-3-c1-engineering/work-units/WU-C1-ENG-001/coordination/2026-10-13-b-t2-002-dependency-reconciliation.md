# B-T2-002 dependency evidence reconciliation — 2026-10-13

## Authority and question

This is a read-only Orchestrator reconciliation of B-T2-002 against the current Approved Techplan, T2 task, accepted manifest, terminal Build retry-1, independent Code Review, and latest independent Testing. It does not change dependency/Product/Techplan authority or the blocker state.

- Approved Techplan: `techplan/techplan.md`, SHA-256 `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79`.
- T2 task: `techplan/tasks/T2-backend-identity-session.md`, SHA-256 `dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37`.
- Accepted manifest: `techplan/tasks/manifest.md`, SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.

## Owning conditions and evidence map

| Condition | Owner / authority | Current evidence | Reconciliation |
|---|---|---|---|
| Go 1.26.9, exact approved direct/explicit pins, actual selected graph and module/source origins; no unsupported material dependency choice | Build, Techplan §§10/12 and T2 task D8 | T2-006 retry-1 actual/isolated gate, graph, origin/tag, legacy-source and compatibility captures | PASS for the captured target. Independent Review confirms captured 20-module graph/version/edge comparison, origin checks, seven legacy modules/133 source-file comparisons, and current exact six-file delta. This is evidence review, not a new upstream fetch. |
| `go mod verify`, focused compile, full source scan and per-ID dispositions before dependency commit | Build, Techplan §10 and T2 task D8 | Retry-1 actual captures and report; both 55-entry disposition sets | PASS for the captured target. Actual `go mod verify` and focused compile passed. Govulncheck v1.7.0 source scan at Go 1.26.9 recorded zero finding objects; 55 unique advisory IDs have fresh metadata and dispositions, including the five pgx/text IDs outside selected affected ranges. No waiver is reported. This does not promise future or universal security. |
| Review the exact dependency/G1 delta | Independent Review | `RUN-C1-ENG-CODEREVIEW-T2-002/evidence/review-findings-1.md`, SHA-256 `3487d5c180d9d60d4c3ad580d20db13e1646b9042361b67fadc7f470fe41a517` | Approve, four passes, 0 blocking and 0 non-blocking code findings, 32/32 bindings, exactly one handoff. The six reviewed hashes still match the current files. Review did not run runtime tests or close blockers. |
| Real configured PostgreSQL pool/SCRAM authentication and session queries after the dependency delta; under-lock behavior | Testing, Techplan §12 R14 | `RUN-C1-ENG-TEST-T2-AUTH-HARNESS-RUNTIME-001/evidence/testing-report-1.md`, SHA-256 `00fe38a3f14c7041d77aba6825a752eb4e8fe1de379a177e09af785893d8819e` | **Not established.** The bounded harness Testing PASS used a real disposable PostgreSQL/TCP/SCRAM fixture and real contention, but did not run the configured application runtime/pool or deployed listener. Its report leaves the configured-runtime condition open. `go mod verify` was reused only because `go.mod` and `go.sum` exactly match retry-1 hashes. |
| Owner acceptance and dependency commit | Orchestrator/Human under T2 task §D8 and the exact approved candidate record | Candidate approval plus retry-1 evidence | **Not established.** The candidate approval selected planned versions but explicitly did not authorize actual graph acceptance, residual-risk acceptance, or dependency commit. No commit was performed or inferred. |

## Terminal evidence identities

- T2-006 retry-1 report: `runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/build-report.md`; SHA-256 `4a63608a42dd89c9978c62161b0f6924700115fd0ed06eeed9af42aa04d6b326`; exactly one `## Phase handoff`. The original terminal BLOCKED report remains immutable. Retry authorization is preserved at `runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/authorization.md`.
- Review report: four-pass Approve for exact six files; SHA-256 and bindings above; no test/runtime operation.
- Latest Testing report: bounded PASS, 48/48 bindings, SHA-256 above; HTTP used `httptest`, Provider was simulated, and cleanup evidence covers the successful fixture path.
- Current `backend/go.mod` SHA-256: `90b7c1c6e870e32828fa5932a6c4e9a28a65d0671a74bdb807ec29dbc5d27cae`.
- Current `backend/go.sum` SHA-256: `6923543aab63e09da2b894699afb70b2ff1072a2eb014e063c8542ffed7deaa0`.
- Current hashes of the four reviewed backend auth files equal the exact identities in the Review report.

## Disposition and next route

`B-T2-002` remains **OPEN**: Build-owned D8 gates and exact-delta Review are evidenced as passed, while Testing-owned R14 configured-runtime evidence and owner acceptance/commit remain unestablished. No D8 Build/scanner rerun is indicated by the current evidence. The next technical evidence should be a fresh independent Testing assignment for the exact R14 runtime gap, with a disposable fixture and current Run-specific approvals. During that scope design, assess whether the B-T2-003 post-lock-expiry oracle can be covered by the same fixture without expanding the claim. Do not automatically include real Google login or integrated UI/cache work; their Human holds remain specific. No dependency commit, B-T2-002 closure, T2/T3 acceptance/readiness, or T3 dispatch is claimed.

No test, scanner, dependency operation, source write, commit, or runtime action occurred during this reconciliation. Historical terminal reports and the dirty workspace were preserved.
