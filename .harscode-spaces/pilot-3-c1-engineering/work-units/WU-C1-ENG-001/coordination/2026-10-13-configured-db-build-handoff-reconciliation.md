# Configured-pool Build handoff reconciliation — 2026-10-13

## Terminal identity and bindings

Terminal Build/Patch report `runs/RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001/evidence/patch-report-1.md` has SHA-256 `33aa73d184433acdf296703cb79089518436d5187169bb5528374c462e7d8de3` and exactly one `## Phase handoff`. `input-bindings.json` contains 15 assignment-defining inputs. Recomputed live content identities match 12 unchanged bindings and the three reported authorized post-delta hashes; no unexpected bound drift was found. Current post-delta hashes are:

- `backend/internal/platform/auth/harness_test.go`: `72c63063aa6709216d110d3f6a41cf81d981b510c9e583305b56d43399567ba8`
- `backend/scripts/auth_harness.py`: `fa5ebc1442b18cce82d5b5b1270855baa661ad4fbadb82456ffbf3b488d8948a`
- `backend/README.md`: `aa48bbeca8f8d186cc806d53769806da6684a1f3fac9cc80ea16de3b66fb0ac4`

The source delta is exactly the three files authorized by `scope.md`; the terminal report is the only Run-owned evidence output. No historical Run report was rewritten and the pre-existing dirty/untracked workspace was preserved.

## Claim and evidence boundaries

The Build report records the single scoped `t2` command passing all six selected cases on its Run-owned disposable PostgreSQL fixture. It reports the restricted fixture child's `DATABASE_URL` validated in the harness and passed through production `db.Open`; the provisioner and lock witness retain separate fixture identities. PostgreSQL 16.14 over loopback TCP/SCRAM, role/privilege checks, and the blocking edge plus expiry-under-lock assertion are Builder-reported real PostgreSQL evidence. HTTP checks used `httptest`; callback Provider behavior was simulated. No Google/provider exchange, browser, deployed listener, production HTTPS, scanner, dependency gate, `make verify`, or independent Testing ran. The runner reports ownership-checked removal of this fixture's container and anonymous volumes on the successful path; abnormal cleanup paths are not established here. No DSN or generated credential was included in the report.

These Build results make a fresh configured-pool Testing assignment executable; they do not establish independent R14 evidence or close B-T2-003. Keep R14 for B-T2-002 and B-T2-003's post-lock-expiry oracle as separate Testing claims even if one fresh fixture supports both. B-T2-002 also still requires owning acceptance/commit. No dependency commit, blocker closure, T2/T3 readiness/acceptance, hold release, or T3 dispatch is established.

## Review applicability and next route

Targeted independent Review is applicable and materially useful: this delta changes child-process credential/environment routing, restricted-fixture URL validation, use of production `db.Open` as the configured-pool oracle, and ownership/cleanup behavior. A fresh four-pass static Review is prepared in `RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-PATCH-001` with 21 exact bindings. It has a fresh Reviewer/Session and non-gated `gpt-6-luna / medium`; Human mechanical dispatch is pending. No Review or runtime operation has occurred.

Testing's next technical prerequisite is satisfied by this terminal Build, so a bounded fresh Testing assignment is staged. Do not dispatch it until this Review is terminal and its findings are reconciled against the exact current delta; if a blocking patch changes the target, rebind or replace the staged assignment before dispatch. The prepared `RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-001` has a new Participant/Session and binds the exact harness/authority inputs. Its own disposable fixture authorization remains pending and does not transfer from another Run. Its Invocation requires separate R14 configured-pool/session-query and B-T2-003 post-lock-expiry results even if one fixture supports both. Existing D8 Build and dependency Review evidence remain applicable to their captured unchanged inputs; no scanner/D8 replay is indicated.
