# Patch Implementation Plan — 01-register-email-verification

> Ticket    : 01-register-email-verification
> Stage     : 5-patch (post-review remediation)
> Date      : 2026-08-21
> Author    : GLM 5.2 (max)
> Source    : `../4-code-review/report.md` (four-pass review, verdict: request changes)
> Tasks     : `../4-code-review/tasks/{manifest,task-a..f}.md`
> Convention: `AGENTS.md` (root + backend), `backend/README.md`

---

## 0. Status

- Review found **7 blocking findings** (S1, S2, S3, S4, S5, L1, L2) and
  **7 optional** (R15, S6, H1, Q1, Q2, Q3, E1).
- All findings verified against current source at the cited locations
  (confirmed accurate as of 2026-08-21).
- Remediation decomposed into Tasks A–F (A–E blocking, F optional).

## 1. Decisions (resolved)

| Decision | Choice | Rationale |
|---|---|---|
| Task B timing test | Integration test vs real Postgres (`//go:build integration`, ≤2× band) | Authoritative; reflects real DB latency, not microsecond fakes |
| Task B no-op write shape | `RevokeTokens(ctx, tx, uuid.New(), purposeEmailVerify)` (0 rows) | Reuses existing method; same UPDATE+commit cost shape as R2; no new surface area |
| Task A + B session | One coordinated session (disjoint methods in `service.go`) | A=VerifyEmail, B=Register; avoids rebase churn |
| Plan persistence | `5-patch/plan.md` (this file) | Matches the N-stage convention |

## 2. Execution order & grouping

```
Session 1 (blocking, service.go coordination):
  Task A  VerifyEmail atomicity + silent-failure (S1, S2)
  Task B  Anti-enumeration DB-time uniformity R3/R4 (S3)
        └─ both edit service.go / repository.go / repository_db.go;
           A is on VerifyEmail, B is on Register — disjoint methods

Session 2 (blocking, file-disjoint):
  Task C  ResendVerification handler error logging (S4)
  Task D  Down-migration ordering fix (S5)
  Task E  Sensitive-error logging hardening (L1, L2)

Session 3 (optional, follow-up PR):
  Task F  Minor hardening (R15, S6, H1, Q1, Q2, Q3, E1)
        └─ depends on A+B test-shape changes
```

**Critical path:** A + B + C + D + E (all blocking). A and B both edit
`service.go` (A on `VerifyEmail`, B on `Register` — disjoint methods);
do them in one session to avoid conflicts. D and E are file-disjoint
from A/B/C and from each other.

## 3. Dependency graph

```
Task A (VerifyEmail atomicity)      ─┐
Task B (anti-enumeration DB-time)   ─┤  all touch service.go / repo
Task C (resend error logging)       ─┤  → coordinate where files overlap
Task D (migration down ordering)    │   (D and E are file-disjoint)
Task E (sensitive logging)          ─┘

Task F (minor) ── after A and B (test assertions depend on new shapes)
```

## 4. Per-task summary

### Task A — `VerifyEmail` Atomicity + Silent-Failure Fix (S1, S2)

**Findings:** S1 (silent success in `VerifyEmail`/`userIDForToken`),
S2 (non-atomic redeem + set-verified — root cause of S1).

**Core change:**
- `Repository.RedeemToken` signature: `(bool, error)` →
  `(userID uuid.UUID, purpose string, ok bool, err error)`, add
  `tx pgx.Tx` param so the subsequent `SetUserVerified` runs in the
  same transaction.
- `repository_db.go`: use goqu `Returning("user_id", "purpose")` +
  `QueryRow.Scan`; `pgx.ErrNoRows` → `ok=false, err=nil`.
- `service.go`: wrap redeem + `SetUserVerified` in one `pgx.Tx`;
  delete `userIDForToken` (no re-fetch path to silently fail).
  `FindAuthTokenByHash` disambiguation (ok==false path) stays a read,
  run outside the tx (nothing to undo on that branch).
- If `SetUserVerified` fails, the deferred `Rollback` undoes the redeem
  — token is *not* burned, user can retry. (S2 fixed.)
- No `uuid.Nil` return path remains — every failure returns a real
  error → handler maps to 500 (not a fake 200). (S1 fixed.)

**Files:**
| File | Change |
|---|---|
| `internal/domain/account/repository.go` | `RedeemToken` signature |
| `internal/domain/account/repository_db.go` | `RETURNING` + tx param |
| `internal/domain/account/service.go` | `VerifyEmail` tx wrap; delete `userIDForToken` |
| `internal/domain/account/service_test.go` | update fake signature; new tests |
| `internal/domain/account/repository_db_integration_test.go` | `TestRedeemAndVerify_Atomic` |

**Tests:**
- Update `fakeRepo.RedeemToken` to `(userID, purpose, ok, err)`.
- `TestVerifyEmail_ValidToken_SetsVerifiedAt` — userID now from
  `RedeemToken`'s return, not a re-fetch.
- NEW `TestVerifyEmail_SetVerifiedFails_RollsBackRedeem` — inject
  `setVerifiedErr`; assert redeem rolled back (token not consumed),
  `VerifyEmail` returns wrapped error (→ 500), not `nil`.
- NEW `TestVerifyEmail_RedeemReturnsUserID_NoRefetch` — assert no
  `FindAuthTokenByHash` call on success path.
- Keep R9/R10/R11/R12; update to new signature. R12 (concurrent
  double-submit) still asserts exactly one success.
- Integration: NEW `TestRedeemAndVerify_Atomic` — seed token; redeem +
  set-verified in one tx; assert both committed. Document rollback
  guarantee by deferred-Rollback pattern + unit test.

**Non-goals:** do not touch the 3-clause predicate; do not change
response mapping (handler untouched); do not touch `Register` (Task B).

---

### Task B — Anti-Enumeration DB-Time Uniformity (R3/R4) (S3)

**Finding:** S3 (R3/R4 perform zero DB writes → timing side-channel
leaking "verified/google-only" (fast) vs "new/unverified" (slow)).

**Core change:**
- R3/R4 branches: call `dummyWrite(ctx)` → begin tx,
  `RevokeTokens(ctx, tx, uuid.New(), purposeEmailVerify)` (synthetic
  uuid matches 0 rows — no-op UPDATE with same cost shape as R2's real
  revoke), commit.
- R1/R2 already do real writes inside their own tx — no dummy needed.
- Per-branch DB work after fix:
  - R1: tx(3 inserts + commit)
  - R2: tx(revoke + insert + commit)
  - R3: tx(dummy revoke 0 rows + commit) + sendNudge
  - R4: tx(dummy revoke 0 rows + commit) + sendNudge
  - All four perform `BeginTx` + ≥1 `UPDATE`/`INSERT` + `Commit` —
    equivalent DB-round-trip shape. (R7 DB-time half satisfied.)

**Files:**
| File | Change |
|---|---|
| `internal/domain/account/service.go` | R3/R4 branches call `dummyWrite` |
| `internal/domain/account/service_test.go` | fake records dummy revoke; new tests |
| `internal/domain/account/repository_db_integration_test.go` | NEW integration timing test |

**Tests:**
- NEW `TestRegister_R3R4_PerformTimingWrite` (unit) — assert R3 and R4
  each invoke exactly one `RevokeTokens` (the dummy), so no-op branches
  are no longer write-free.
- NEW `TestRegister_Timing_AllBranches_RealPostgres`
  (`//go:build integration`) — runs each branch against real Postgres,
  asserts `max/min ≤ 2×`. Add comment: now covers DB-time equivalence,
  not just bcrypt; this test is authoritative.
- Keep `TestRegister_GenericResponse_AllBranches` (1 email per branch
  unchanged). Existing `TestRegister_GenericResponse_Timing` (fake
  repo) is retained as a bcrypt-only smoke but annotated as
  non-authoritative for DB-time.

**Non-goals:** do not change the 202 response; do not remove bcrypt
(CPU-time half is correct); do not touch `VerifyEmail` (Task A).

---

### Task C — `ResendVerification` Handler Error Logging (S4)

**Finding:** S4 (`_ = svc.ResendVerification(...)` discards the error
with no logging).

**Core change:**
- `auth_verify_email.go:66`:
  ```go
  if err := svc.ResendVerification(r.Context(), req.Email); err != nil {
      // Anti-enumeration: response stays identical 202 (R14). The
      // internal failure must be visible to ops — log a sanitized
      // category, not the recipient (PII).
      log.Printf("transport: resend verification failed (recipient redacted): %v", err)
  }
  ```
- Response stays 202 identical (do NOT route through `MapServiceError`
  — a 500 on the match branch would distinguish it from no-match, an
  enumeration leak).
- `%v` on this specific error chain is safe: the leaf is a
  `pgconn.PgError` (SQLSTATE + constraint name; parameterized SQL, no
  PII values). If Task E's L2 sanitization helper lands, prefer it here
  too.

**Files:**
| File | Change |
|---|---|
| `internal/transport/http/auth_verify_email.go` | log error before 202 |
| `internal/transport/http/auth_verify_email_test.go` (new or extended) | assert 202 + log present + no PII |

**Tests:**
- NEW `TestResendVerificationHandler_ServiceError_Still202_ButLogs` —
  inject a failing service; assert response is 202 (not 500), log line
  contains "resend verification failed", log does NOT contain the
  recipient email.

**Non-goals:** do not change the 202 response or body; do not add
`MapServiceError` routing.

---

### Task D — Down-Migration Ordering Fix (S5)

**Finding:** S5 (`DROP FUNCTION set_updated_at()` in `000003.down` fails
while triggers still depend on it — reverse-run order confusion).

**Core change:**
- Remove `DROP FUNCTION IF EXISTS set_updated_at();` from
  `migrations/000003_create_auth_tokens.down.sql`.
- Add it to `migrations/000001_create_users.down.sql` *after*
  `DROP TRIGGER` + `DROP TABLE` (000001 down runs last in reverse order,
  after both triggers are gone).
- Fix comments to explain reverse-run order:
  ```
  000003 down:  drop auth_tokens indexes + table (no trigger on it)
  000002 down:  drop trg_auth_identities_updated_at + auth_identities
  000001 down:  drop trg_users_updated_at + users + set_updated_at()  ← safe
  ```

**Files:**
| File | Change |
|---|---|
| `migrations/000003_create_auth_tokens.down.sql` | remove `DROP FUNCTION` line + fix comment |
| `migrations/000001_create_users.down.sql` | add `DROP FUNCTION` after trigger+table drop + fix comment |

**Verification (manual, dev Postgres):**
```bash
make migrate-up          # "no change" expected
make migrate-down        # must succeed — no "cannot drop function" error
make migrate-up          # re-apply succeeds
# psql: \df, \dt, \d users, \d auth_identities, \d auth_tokens sanity
```
If `migrate-down` fails: capture exact error + `schema_migrations`
state before proceeding; dirty state needs `migrate force` (manual ops).

**Non-goals:** do not edit up migrations (already applied; per
`migrations-safety.md` applied migrations are never edited); do not add
`CASCADE` to `DROP FUNCTION` (masks the dependency — fix is ordering).

---

### Task E — Sensitive-Error Logging Hardening (L1, L2)

**Findings:** L1 (`breachcheck/client.go:52` logs `Do` error verbatim —
can embed request URL with 5-char SHA-1 prefix), L2
(`service.go:385`/`:393` logs notification-sender error verbatim — can
embed recipient/token).

**Core change:**
- L1: replace `%v` of `Do` error with `breachErrorCategory(err)` —
  extract `*url.Error.Op` (safe) + coarse net category (timeout /
  canceled / connection-reset / network-error); never the URL or
  SHA-1 prefix. Fallback: `"transport error"`.
- L2: replace `%v` in `sendVerification` and `sendNudge` with
  `notificationErrorCategory(err)` — coarse category (timeout /
  "send failed"); never the raw message. Latent today (`FakeSender`
  returns nil) but the seam will get a real SMTP sender.
- Leave `transport/http/errors.go:76` as `%v` — its leaf is a DB driver
  error (SQLSTATE, no PII values; parameterized SQL). Documented here
  so a future reviewer doesn't "fix" it and broaden the change.

**Files:**
| File | Change |
|---|---|
| `internal/platform/breachcheck/client.go` | `breachErrorCategory` helper; sanitized log |
| `internal/platform/breachcheck/client_test.go` | assert no URL / SHA-1 prefix |
| `internal/domain/account/service.go` | `notificationErrorCategory` helper; sanitized logs in `sendVerification`/`sendNudge` |
| `internal/domain/account/service_test.go` | assert no recipient / token in send-failure log |

**Tests:**
- NEW `TestIsBreached_APIUnreachable_LogNoURLNoPrefix` — force a
  `*url.Error` carrying the request URL; assert log contains
  "breachcheck: API unreachable", does NOT contain
  "pwnedpasswords.com" or the 5-char SHA-1 prefix.
- NEW `TestRegister_SendVerificationFails_LogNoPII` — wire a
  `captureSender` whose `SendVerificationEmail` returns an error whose
  `Error()` contains the recipient email + a fake token (simulating a
  leaky SMTP error); assert log contains "send verification email
  failed", does NOT contain the recipient email or token.

**Non-goals:** do not introduce a structured logger (`zap`/`slog`) —
codebase uses `log` stdlib; do not change `transport/http/errors.go:76`
(DB-error chain, safe); do not redact `nudgeType` (package constant, not
user input).

---

### Task F — Minor Hardening (Optional, follow-up PR)

**Findings:** R15, S6, H1, Q1, Q2, Q3, E1. None block merge.

**Hard deps:** Tasks A and B (Q2's `&&`→`||` fix is in
`service_test.go`, which A and B also edit — land after them).

**Items:**

| # | Finding | Change | File(s) |
|---|---|---|---|
| R15 | Missing rate-limit N+1 test | `TestRateLimit_RejectsNPlusOne` (rps=1, burst=1 → 2nd req 429) + per-IP independence | `internal/transport/http/middleware_test.go` (new) |
| S6 | Sweeper goroutine no exit path | `RateLimit(ctx context.Context, …)` + sweeper `select` on `ctx.Done()`; update `main.go` wiring | `middleware.go`, `cmd/server/main.go` |
| H1 | breachcheck body not drained on non-OK | `io.Copy(io.Discard, resp.Body)` before close on non-OK path | `breachcheck/client.go` |
| Q1 | Duplicated begin/commit/rollback | Extract `runTx(ctx, fn)` helper; refactor `registerNewUser`/`issueNewVerificationToken` + A/B's new tx uses | `service.go` |
| Q2 | `&&`→`\|\|` test bug | `service_test.go:810` `&&` → `\|\|` | `service_test.go` |
| Q3 | Duplicated `looksLikeEmail` | Shared helper in `transport/http/validate.go` | `auth_register.go`, `auth_verify_email.go`, `validate.go` (new) |
| E1 | No runtime HMAC≠enc key check | `crypto/keys.go` `New`: `if bytes.Equal(encryptionKey, hmacKey) { return nil, error }` | `internal/platform/crypto/keys.go` |

**⚠️ Fence flag (E1):** `platform/crypto/` is Tier 0 fenced
(`AGENTS.md` §3). E1 touches `keys.go` (the keys *holder*/loader), not
the encryption/HMAC/key-handling impl in `crypto.go`. The intent is a
one-time startup config guard, not a change to crypto primitives. This
needs an explicit human confirm before E1 proceeds — if the human
declines, E1 is dropped from Task F and the gap is documented as an
accepted config-discipline risk.

**Verification:**
```bash
go test -count=1 ./...
go test -race -count=1 ./internal/domain/account/... ./internal/platform/crypto/...
go vet ./...
make verify
```

**Non-goals:** do not fold these into the blocking PR if it would delay
it — explicitly optional. E1 must not change key *loading* (base64
decode, 32-byte validation) — only add the inequality check.

## 5. Fence check (Tier 0)

No blocking task (A–E) touches Tier 0 fenced paths:
- `backend/internal/domain/donation/ledger.go` — untouched
- `backend/internal/domain/disbursement/` — untouched
- `backend/internal/platform/crypto/` *impl* (`crypto.go`) — untouched
- `backend/internal/platform/auth/` — untouched

Task F's E1 adds one inequality check to `crypto/keys.go` (the keys
holder/loader, not the crypto impl). Flagged for human confirm before
E1 proceeds (see §4 Task F).

No remediation task edits `docs/spec/*` (per `AGENTS.md` §4 — spec/test
authority separation). The build's prior INV-account-08 spec edit was a
human-asked exception, documented in `3-build/report.md` §6.2 — not
repeated here.

## 6. Final gate

```bash
make verify          # lint, unit, race, contract, security-layer-A, integration
go test -race ./internal/domain/account/...   # Tier 1 explicit
```

Each task ships with the `AGENTS.md` §5 risk-note structure:
```markdown
## Risk note
- Assumptions made: ...
- Edge cases intentionally NOT handled (and why): ...
- Concurrency assumptions: ...
- What is not tested, and why: ...
```
Every claim of "this edge case is handled" must name the specific test
that proves it. A claim with no named test is treated as unverified.

## 7. PR description reference

Each remediation PR must reference which review finding(s) it resolves
and link back to `../4-code-review/report.md` + the relevant
`../4-code-review/tasks/task-*.md`. The blocking PR(s) reference the
feature spec
`docs/spec/domains/account/features/01-register-email-verification.md`
as fulfilled by the original build (`../3-build`), with this stage as
the review-driven remediation.

## 8. Cross-reference

- Review report: `../4-code-review/report.md`
- Task files: `../4-code-review/tasks/{manifest,task-a..f}.md`
- Source techplan: `../2-plan/techplan.md` (§4 R7/R8-R12, §5 Decision 8)
- Build report: `../3-build/report.md`
- Fencing rules: `AGENTS.md` §3 (Tier 0), §4 (spec/test authority)
- Review checklist: `workflow/4-code-review/checklist.md` — re-run
  after remediation before considering the change done.
