# Code Review Report — 01-register-email-verification

> Ticket    : 01-register-email-verification
> Stage     : 4-code-review (four-pass review)
> Date      : 2026-08-20
> Reviewer  : GLM 5.2 (max)
> Diff      : `../2-plan` + `../3-build` (24 new files + 7 modified)
> Convention: `backend/AGENTS.md` (+ root `AGENTS.md`), `backend/README.md`
> Guidance  : `workflow/4-code-review/{guidelines,checklist,examples}.md`, `best-practices/index.md`

---

## 0. Review method

Four passes, in order: Safety → Quality → Stack-Specific Best Practices
→ Consistency. Each pass is a different lens; a clean pass says nothing
about the next. Pass 3 matched `best-practices/index.md` trigger
keywords against the diff's technology (Go, PostgreSQL, REST API) and
the Security Concern Map (authn, secrets-and-keys, input-validation,
pii-and-encryption, rate-limiting-and-abuse, enumeration-and-timing),
then opened only the matching best-practices files.

Files opened for Pass 3: `restapi/anti-enumeration.md`,
`postgresql/encryption-at-rest.md`, `postgresql/transactions-and-locking.md`,
`postgresql/migrations-safety.md`, `go/rate-limiting.md`,
`go/secrets-and-sensitive-logging.md`, `go/http-client-and-transport.md`,
`go/goroutine-lifecycle.md`, `go/nil-and-zero-values.md`,
`go/jwt-and-token-lifecycle.md`.

Not every checklist item applies to every change. Findings below name a
specific location, why it matters, and a suggested fix. "No findings" is
not used to pad — where a pass is clean on a sub-topic, it says so.

---

## 1. Safety

### S1 (BLOCKING) — Silent success-on-failure in `VerifyEmail`: user told "verified" while not verified, token burned

**Location:** `internal/domain/account/service.go:324-334` (`userIDForToken`),
called from `service.go:302-304`.

After a successful `RedeemToken` (`ok==true`), the service re-fetches the
token to obtain `user_id`. On *any* error or nil, `userIDForToken` logs
and returns `uuid.Nil` instead of an error. `SetUserVerified(uuid.Nil, …)`
then matches 0 rows and returns `nil`. `VerifyEmail` returns `nil` → the
handler writes `200 "Email verified."` But `verified_at` was never set,
and the token is now `used_at`-marked. The user sees success, is *not*
verified, and cannot retry (token consumed).

This is the "silent skip when an error was expected" pattern
(`workflow/4-code-review/examples.md` row 1) and directly violates the
root golden rule "Errors are always returned, never swallowed"
(`AGENTS.md` §2).

**Why it matters:** a transient DB error on the post-redeem re-fetch
burns the user's only token while reporting success. The user is stuck
(unverified, no usable token) with no signal to ops that anything went
wrong.

**Suggested fix:** return an error from `userIDForToken` and propagate
it as a 500; *or* (preferred) have `RedeemToken` return the `userID`
directly so no re-fetch is needed. See Task A.

### S2 (BLOCKING, root cause of S1) — `RedeemToken` and `SetUserVerified` are not atomic

**Location:** `service.go:294-304`; `repository_db.go:247-265`
(`RedeemToken`, auto-tx via `r.db.Exec`) and `:272-285` (`SetUserVerified`,
separate auto-tx).

The atomic single-use `UPDATE … WHERE 3-clause` commits `used_at` in one
transaction; setting `verified_at` is a *separate* transaction. If the
second fails (DB error / crash / network blip), the token is consumed
but the identity is unverified — inconsistent state. The design comment
"single-use correctness is delegated to `RedeemToken`" only covers the
token-used half, not the verify half.

**Why it matters:** the two operations can fail independently, leaving
the token used without the identity verified. Recoverable via resend
(identity still unverified → resend works), but the design implies an
atomicity it does not have, and S1 turns the failure into a silent one.

**Suggested fix:** run redeem + set-verified in a single `pgx.Tx`; *or*
have `RedeemToken` return `userID` and issue both `UPDATE`s in one
transactional statement. See Task A.

### S3 (BLOCKING) — Anti-enumeration timing side-channel: R3/R4 perform no DB writes

**Location:** `service.go:130-162` (Register branch dispatch).

Per-branch DB work:

| Branch | DB reads | DB writes |
|---|---|---|
| R1 (new user) | 2 | tx(3 inserts + commit) |
| R2 (unverified existing) | 1 | tx(revoke + insert + commit) |
| **R3 (verified existing)** | **1** | **0** |
| **R4 (google-only conflict)** | **2** | **0** |

R3/R4 skip all DB writes, so against real Postgres they are measurably
faster than R1/R2 — leaking "new/unverified" (slow) vs
"verified/google-only" (fast). This is precisely the enumeration the
threat-model §1 / feature spec Assumption B exist to prevent, and the
techplan explicitly committed to "DB-write-shaped work on all branches"
(`2-plan/techplan.md` §4 R7, §5 Decision 8). The build only implemented
the CPU-time half (always-bcrypt), not the DB-time half.

**The proof is invalid:** `TestRegister_GenericResponse_Timing`
(`service_test.go:573-619`) uses the in-memory `fakeRepo` where DB ops
are ~microseconds, so it *cannot* catch DB-timing differences — it only
asserts bcrypt equivalence (the test's own comment at `:573-579` says
so). False confidence: R7 is marked ✅ in the build report but the
DB-time half of R7 is unverified.

**Why it matters:** this is the feature's defining constraint. An
attacker with statistical timing analysis over many requests can
distinguish "verified/google-only" (fast) from "new/unverified" (slow).

**Suggested fix:** give R3/R4 equivalent DB-write-shaped work (e.g., the
same revoke+insert shape as R2, discarded), and add a timing test
against real Postgres (or a latency-simulating fake) asserting DB-time
equivalence. See Task B.

### S4 (BLOCKING) — `ResendVerification` handler silently discards the error with no logging

**Location:** `internal/transport/http/auth_verify_email.go:66`
(`_ = svc.ResendVerification(r.Context(), req.Email)`).

If `issueNewVerificationToken` fails (DB error), no token is issued and
no email is sent, but the user gets `202 "you will receive a new
verification link."` Anti-enumeration justifies the 202 *response*, not
the total absence of logging. The golden rule allows "logged with
enough context to act on" — this does neither. Compare the Register
handler (`auth_register.go:48-61`), which routes errors through
`MapServiceError` (which logs).

**Why it matters:** the resend flow's internal failures are invisible
to ops; the user is misled into expecting an email that will never
arrive.

**Suggested fix:** log the error server-side before returning the 202.
See Task C.

### S5 (BLOCKING) — Down-migration ordering bug: `DROP FUNCTION` fails while triggers still depend on it

**Location:** `migrations/000003_create_auth_tokens.down.sql:6`
(`DROP FUNCTION IF EXISTS set_updated_at();`).

`golang-migrate` runs down in reverse order: 000003 → 000002 → 000001.
When `000003.down` runs, the triggers `trg_users_updated_at` (users) and
`trg_auth_identities_updated_at` (auth_identities) still exist and
reference `set_updated_at()`. Postgres rejects
`DROP FUNCTION set_updated_at()` (no `CASCADE`) with "cannot drop
function because other objects depend on it," leaving the migration
half-applied and the `schema_migrations` state "dirty" (cannot proceed
up or down without manual cleanup).

The comment in `000001_create_users.down.sql:3-4` says the function is
"dropped only in the final down migration" but then points to `000003
down` — which runs *first* (reverse order). The author confused
file-number order with reverse-run order.

**Why it matters:** `make migrate-down` (advertised in `README.md`)
fails on a clean applied DB and dirties the schema state. The down path
was written symmetrically, not tested — exactly what
`postgresql/migrations-safety.md` warns against.

**Suggested fix:** move `DROP FUNCTION IF EXISTS set_updated_at();` into
`000001_create_users.down.sql` (after the users trigger+table drop,
which runs last), and verify by running `make migrate-down` from a
clean applied state. See Task D.

### S6 (minor) — Rate-limiter sweeper goroutine has no exit path

**Location:** `internal/transport/http/middleware.go:29-43`.

The background sweeper `go func(){…}()` has no context/done channel; it
isn't stopped on graceful shutdown (`cmd/server/main.go:147`). Process-
scoped singletons are tolerable, but per `go/goroutine-lifecycle.md` §1
it should have an identified exit path.

**Suggested fix:** accept a `context.Context` in `RateLimit` and
`select` on `ctx.Done()` in the sweeper (or document it as
intentionally process-scoped). See Task F.

---

## 2. Quality

### Q1 — Duplicated begin/commit/rollback pattern

**Location:** `service.go:178-187` (`registerNewUser`) and
`service.go:254-263` (`issueNewVerificationToken`).

Both duplicate the
`committed := false; defer func(){ if !committed { _ = tx.Rollback(ctx) } }()`
shape. A shared `runTx(ctx, func(pgx.Tx) error) error` helper would
centralize transaction lifecycle and reduce duplication. Non-blocking.
See Task F.

### Q2 — Test assertion logic bug masks regressions

**Location:** `internal/domain/account/service_test.go:810`
(`if len(sender.verificationTo) != 0 && len(sender.nudgeTypes) != 0`).

Uses `&&` (AND); should be `||` (OR). As written, it only fails if
*both* are non-zero, so it wouldn't catch a regression where a no-match
branch accidentally sends one email type.

**Suggested fix:** change `&&` → `||`. See Task F.

### Q3 — `looksLikeEmail` duplicated across handlers

**Location:** `auth_register.go:77-85` and `auth_verify_email.go:59`.

Minor: the check exists in two handlers and the service never validates
email shape (it HMACs whatever it gets). Acceptable as boundary
validation, but a third auth endpoint should take a shared helper.
Non-blocking. See Task F.

---

## 3. Stack-Specific Best Practices

Matched via `best-practices/index.md` trigger keywords + Security
Concern Map. Files opened listed in §0. If a checklist was clean, it
says so.

### From `restapi/anti-enumeration.md`

- Checklist item 3 ("Response timing on sensitive endpoints doesn't
  differ significantly between found and not found"): **violated → S3**.
  "Always-bcrypt" satisfies CPU-time but DB-time differs (R3/R4 do zero
  writes).
- Checklist item 1 (identical generic response): ✓ — register/resend
  return identical 202 across branches.
- Checklist item 2 (constant-time comparison against secrets): N/A —
  tokens are high-entropy SHA-256 hashes looked up by indexed equality,
  not app-level secret compares. Correct by design.
- Checklist item 4 (error messages don't leak internal details): ✓ —
  `WriteValidationError` echoes field names only, never values.

### From `go/rate-limiting.md`

- Checklist item 4 ("a test that verifies the limiter actually rejects
  the N+1th request"): **missing**. The build report acknowledges R15 ⚠️
  (`TestResend_RateLimited` not written). See Task F.
- Items 1–3 (eviction, granularity, configurable): ✓ — idle-key
  eviction present (`middleware.go:29-43`), per-IP for anonymous,
  configurable via `AUTH_RATE_RPS`/`AUTH_RATE_BURST`.

### From `go/secrets-and-sensitive-logging.md` §1

Logging a third-party client error verbatim can embed request/response
payloads:

- `internal/platform/breachcheck/client.go:52` logs `c.httpClient.Do`'s
  error via `%v`. That error can embed the request URL
  `api.pwnedpasswords.com/range/{5-char-SHA1-prefix}` — partial
  credential-derived data (the k-anonymity prefix is safe to *send*,
  not to *log*). → **L1**
- `service.go:385` logs the notification-sender error via `%v`; a real
  SMTP error could embed recipient/token. (`FakeSender` returns nil
  today, so this is latent — but the seam will get a real sender.) → **L2**

**Suggested fix:** log a sanitized category, not the raw error string.
See Task E.

### From `go/http-client-and-transport.md` §2

- `breachcheck/client.go:56-58` (non-OK status path): the response body
  is `Close()`d (deferred) but not *drained* (`io.ReadAll` is skipped),
  which can prevent connection reuse. → **H1**
- §1 items (explicit `Timeout`, `NewRequestWithContext`, client reuse):
  ✓. See Task F.

### From `go/goroutine-lifecycle.md` §1

→ S6 (sweeper no exit path). §2 (unbounded per-item spawning): N/A — no
per-item goroutine spawning in this diff.

### From `postgresql/encryption-at-rest.md`

- Ciphertext never in `WHERE` ✓; separate HMAC column ✓; which-columns-
  need-HMAC decided ✓.
- Checklist item 3 ("HMAC key and encryption key are different keys"):
  structurally separate (`Keys.EncryptionKey`/`HMACKey` from separate
  env vars), but the code does not *enforce* they differ at runtime —
  config-discipline only. Minor defense-in-depth gap. → **E1** (Task F)

### From `postgresql/migrations-safety.md`

- "Down migration is genuinely reversible and has been tested, not just
  written symmetrically": **violated → S5**. The function drop was
  placed by file-number symmetry, not reverse-run dependency order.
- Additive-first / no backfill: N/A (greenfield, all new tables). ✓

### From `postgresql/transactions-and-locking.md` §1

- "Never make a network call inside an open DB transaction": **correctly
  followed** — `service.go:230-237` commits before `sendVerification`;
  `issueNewVerificationToken` returns the plain token for the caller to
  send after commit. ✓ Good.
- §3 (Read Committed non-repeatable read): the redeem path avoids
  read-then-write on tokens (single atomic `UPDATE … WHERE`). ✓

### From `go/jwt-and-token-lifecycle.md`

Considered; N/A. These are opaque random tokens (SHA-256 stored, not
JWT-signed); single-use + 24h expiry + revocation is correctly enforced
by the 3-clause guard. Key-separation concerns don't apply (no JWT
signing here). The token-lifecycle best practices apply to JWT
issuance, which is explicitly out of scope (task #3, Tier 0 fenced).

### From `go/nil-and-zero-values.md` / interface-typed-nil pitfall

The `CredentialSecret` nil-in-interface trap was *correctly handled*
(`repository_db.go:105-107` avoids putting a nil `*string` into the
`goqu.Record`). ✓ Good. The `userIDForToken` nil-return → S1.

### From `go/error-wrapping.md`

All errors use `%w` ✓; `errors.As` for unique-violation (not string
match) ✓ (`service.go:419-422`). No finding.

---

## 4. Consistency

Convention sources: `backend/AGENTS.md` (+ root `AGENTS.md` golden
rules), `backend/README.md`.

### C1 (= S1, S4) — "Errors are always returned, never swallowed"

**Convention:** root `AGENTS.md` §2 golden rule; `backend/AGENTS.md` §2
("wrap with `%w`; never discard the original error").

**Violations:**
- `userIDForToken` returns `uuid.Nil` instead of an error (S1).
- `auth_verify_email.go:66` `_ = svc.ResendVerification(...)` discards
  the error entirely — not even logged (S4).

### C2 (= S3) — Violates the repo's own feature contract

**Convention:** `backend/AGENTS.md` §1 makes `docs/spec/*` the source of
truth; the techplan (`2-plan/techplan.md` §4 R7, §5 Decision 8) is this
feature's agreed contract and explicitly committed to "DB-write-shaped
work on all branches" for DB-time uniformity. The build omits it for
R3/R4. Also contravenes the threat-model §1 anti-enumeration
requirement referenced in `AGENTS.md` §1.

### C3 (= L1/L2) — "No secrets, PII, or tokens in logs. Log the fact + outcome, not the payload"

**Convention:** root `AGENTS.md` §2 golden rule.

`breachcheck/client.go:52` and `service.go:385` log upstream errors
verbatim, which can embed credential-derived data (SHA-1 prefix) or
recipient/token. Elsewhere the code is disciplined (`service.go:238`
logs `user_id` not email; `FakeSender` redacts; `:305` "token redacted")
— so these two are *inconsistent* with the codebase's own established
redaction pattern.

### C4 (= S5) — Migrations must be reversible

**Convention:** `backend/README.md` advertises `make migrate-down`;
`AGENTS.md` §4 (spec/test authority) implies migrations must be
reversible. The down path fails on a clean DB.

### C5 — Doc-comment & routing conventions (clean)

`backend/AGENTS.md` §2 ("every exported function/type gets a doc
comment"): followed across `service.go`, `repository.go`, `entity.go`,
`crypto.go`, `breachcheck`, `notification`, `secrets`, `middleware`,
`errors`, handlers. Go 1.22+ pattern routing
(`mux.HandleFunc("POST /auth/register", …)`) and goqu-for-all-SQL
conventions are also correctly followed. ✓

### Process items — correctly handled (not findings)

- **§3 crypto fence-lift** (`platform/crypto/crypto.go`): handled with
  the per-session human-ask exception model, documented in
  `crypto/doc.go:5-11` and `3-build/report.md` §6.1. The lift was
  one-time, not permanent. ✓ Consistent with the fencing rules.
- **§4 invariants.md edit** (INV-account-08 Verification field): the
  human explicitly asked; the fix *tightens* (aligns Verification with
  Statement), documented in `3-build/report.md` §6.2. ✓

---

## Verdict

**Request changes.**

### Blocking (must fix before merge)

| # | Finding | Task |
|---|---|---|
| S1 | Silent success in `VerifyEmail`/`userIDForToken` | Task A |
| S2 | Non-atomic redeem + set-verified (root cause of S1) | Task A |
| S3 | Anti-enumeration DB-timing side-channel (R3/R4 no DB writes) | Task B |
| S4 | Swallowed error in `ResendVerificationHandler`, no logging | Task C |
| S5 | Broken down-migration ordering (`DROP FUNCTION` fails) | Task D |
| L1 | breachcheck logs http error verbatim (SHA-1 prefix) | Task E |
| L2 | service logs notification error verbatim (recipient/token) | Task E |

### Optional / minor

| # | Finding | Task |
|---|---|---|
| R15 | Missing `TestResend_RateLimited` (rate-limit N+1 rejection) | Task F |
| S6 | Sweeper goroutine no exit path | Task F |
| H1 | breachcheck body not drained on non-OK path | Task F |
| Q1 | Extract `runTx` helper | Task F |
| Q2 | `&&`→`\|\|` in `TestResend_NoMatch` | Task F |
| Q3 | Shared `looksLikeEmail` helper | Task F |
| E1 | Runtime check that HMAC ≠ encryption key | Task F |

### Notes for the implementer

- Tasks A–E are blocking and should land before merge. Task F is
  minor/optional and can be a follow-up.
- Tasks A and B both touch `service.go`; do them in one session to
  avoid rebase churn (A is localized to `VerifyEmail`; B is localized to
  `Register`).
- After Task D, re-run `make migrate-down && make migrate-up` from a
  clean state to prove reversibility.
- The Tier 0 fencing on `platform/crypto/` was lifted per-session for
  the original build; Tasks A–F do **not** touch `crypto/`, so no new
  fence lift is required.
- Re-run `make verify` (and `go test -race ./internal/domain/account/...`)
  after all tasks — these are Tier 1 areas per `backend/AGENTS.md` §3.
