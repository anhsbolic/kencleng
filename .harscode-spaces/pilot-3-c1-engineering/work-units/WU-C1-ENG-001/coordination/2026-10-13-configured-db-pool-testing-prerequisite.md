# Configured database pool Testing prerequisite — 2026-10-13

## Finding from read-only readiness inspection

The latest independent Testing Run passed its bounded harness claim but did not establish Techplan §12 R14's configured application pool/session-query condition or the exact B-T2-003 post-lock-expiry condition. The current harness runner removes inherited `DATABASE_URL`, supplies a private `C1_HARNESS_RUNTIME_DSN`, and `harness_test.go` creates the behavior pool directly with `pgxpool.NewWithConfig`. Production server startup instead obtains `DATABASE_URL` and calls `db.Open`. The prior Testing report therefore remains correctly bounded; replaying the same harness path cannot close the stated gap.

The declared Testing prerequisite is not presently executable with the current harness without changing test infrastructure. Under the Testing prompt and orchestration rule for material real-runtime prerequisites, do not dispatch another Testing Run that is known to miss R14. Prepare a narrow Build/Patch prerequisite first.

## Prepared route

Prepared `RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001` with fresh assignment identities and `gpt-6-luna / medium` (current registry: approval not required). Its three-file scope changes only the harness runtime pool construction, the private child `DATABASE_URL` handoff, and README evidence description. It reuses the existing six `t2` cases, including session queries and PostgreSQL lock-wait expiry, against one new Run-owned disposable PostgreSQL 16 fixture. It does not modify production `db.Open`, auth/session behavior, dependency pins, scanner policy, or authority.

This Run is only prepared. Its exact disposable-fixture authorization and Human mechanical dispatch remain pending; no source/test/fixture/runtime action has occurred. No prior Run's fixture authorization transfers.

After its terminal report, Orchestrator will inspect exact delta and evidence, decide Review applicability, and then prepare fresh independent Testing with a new fixture owner. That Testing can evaluate R14 and B-T2-003 separately on one fixture only if the actual invocation and assertions support both claims. B-T2-002 additionally retains its owning acceptance/commit decision after evidence reconciliation. No dependency commit, blocker closure, T2/T3 acceptance/readiness, hold change, or T3 dispatch follows from preparation.
