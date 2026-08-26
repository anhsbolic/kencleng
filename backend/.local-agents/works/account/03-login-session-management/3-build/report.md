# Build Report — Login & Session Management (account #03)

> Task      : account domain task #3 — `docs/spec/1-account/features/03-login-session-management.md`
> Executed  : 2026-08-26, tasks 01–06 per `2-plan/tasks/` decomposition
> Techplan  : `.local-agents/works/account/03-login-session-management/2-plan/techplan.md` (Status: Approved)
> Status    : Build complete — **BLOCKED on Tier 0 paired pass before any commit**

---

## ⚠️ Tier 0 files awaiting paired rewrite (gate — Resolved #13)

Nothing in this slice may be committed until Anhar's dedicated paired
rewrite/review pass covers exactly these agent-drafted regions:

1. **`internal/platform/auth/token.go`** (whole file, task-03) — ES256 access
   + HS256 mfa_pending mint/verify, purpose claims, secret validation.
2. **`internal/domain/account/repository_db.go`** — rotation methods
   `RotateRefreshToken` / `RevokeRefreshTokenByHash` /
   `RevokeRefreshTokenFamily` (task-02 region).
3. **`internal/domain/account/login.go`** — the reuse/race-loser branch of
   `Refresh` incl. the connection-leak fix it required (task-04 region).

Two judgment calls inside Tier 0 territory that especially need human eyes:
- **Pending-token verifier has ZERO clock leeway** (access token keeps the
  house 1-minute leeway) — tightness is deliberate; confirm you agree.
- **Race-loser mint ordering**: the new access token is minted BEFORE the
  rotation tx opens, so a lost race wastes ~50µs of signing instead of
  risking a post-commit mint failure (rotated cookie + 500).

---

## Execution summary

| Task | Scope | Result |
|---|---|---|
| 01 | Migrations 000006–000009 (`login_attempts`, schema-pre-settle `mfa_totp_secrets`, `mfa_backup_codes`, `user_roles`) | ✅ up → targeted `down 4` → up verified; indexes/triggers/FKs confirmed in live DB |
| 02 | Entities (`LoginAttempt`, `LoginUserView`) + Repository port/adapter: attempt writes, two-stage lockout counts, refresh find/rotate/revoke/revoke-family, user-view assembly w/ first decrypt-on-read path | ✅ unit + integration suites green |
| 03 | ⚠️ Tier 0: `platform/auth/token.go` — both token purposes, cross-purpose rejection matrix | ✅ 12 tests incl. spec-named `TestAuthMiddleware_*` |
| 04 | Services: `Login`/`LoginMfa`/`Refresh`/`Logout`, sentinels, fail-closed `mfa_verifier` seam, lockout both stages, timing discipline | ✅ 18 unit tests |
| 05 | Transport: four handlers, refresh-only cookie helpers, 401/429 Problem vocabulary, main.go wiring (env + closures + routes) | ✅ 12 handler tests + live boot smoke |
| 06 | Integration + race suite, final gate | ✅ tests all green; gate outcome documented below |

## Files changed

New: `migrations/000006..000009` (8 SQL files) · `internal/domain/account/{login.go, login_test.go, mfa_verifier.go, login_integration_test.go}` · `internal/platform/auth/token.go` · `internal/platform/auth/token_test.go` · `internal/transport/http/auth_login.go` · `internal/transport/http/auth_login_test.go`

Edited: `entity.go`, `repository.go`, `repository_db.go`, `service.go` (constructor seams), `service_test.go` (+fake stubs), `google_oauth_test.go` + `googleoauth/helpers_test.go` (two pre-existing staticcheck fixes), `internal/transport/http/{cookie.go, errors.go}`, `cmd/server/main.go`, `.env.example`, `.env` (dev secret added), `docs/spec/1-account/tasks.md` + `threat-model.md` + feature-spec wording fix (approved during open-item review).

Untouched by design: `auth_google*.go`, `middleware.go`, `api/openapi.yaml` (amendment already applied separately via `account.yaml` + rebundle), frontend/Caddyfile.

## Verification results

- Unit suite: **all packages green**, `-race` green.
- Integration suite (`//go:build integration`, real Postgres): green, incl.
  - `TestRefresh_ConcurrentRequests_ExactlyOneWins_RealDB`
  - `TestRefresh_Stress_MixedValidAndReplayed` (120 goroutines, mixed valid+replayed)
  - `TestRefresh_ReuseDetection_FamilyRevoked_RealDB` (A→B→C replay-A)
  - `TestLogin_Lockout_EndToEnd` (+ window-expiry release)
- Live boot smoke: server up with `MFA_PENDING_TOKEN_SECRET` gating;
  unknown-credential POST `/auth/login` returned the exact contract 401 body.

### Gate stage outcomes

| Stage | Outcome |
|---|---|
| staticcheck | clean (2 pre-existing findings fixed: S1024 in `google_oauth_test.go`, U1000 in `googleoauth/helpers_test.go`) |
| gosec | **13 findings = exact pre-existing baseline** (verified via `git stash -u`). This slice added 5, each annotated with justified `#nosec G101`/`G124` (Indonesian error text containing "password"; env-conditional Secure cookies). Net-new contribution: 0. The 13 are prior slices' code and need YOUR accepted-risk notes or fixes — not annotated by this agent. |
| gitleaks | cannot run — CLI absent from environment (pre-existing tooling gap) |
| govulncheck | 24 vulns, byte-identical count on pristine `main` (pre-existing stdlib/module debt) |
| test-contract | no contract-tagged tests exist yet; runs unit suite |

`make verify` therefore exits non-zero at gosec **exactly as it did before
this slice started** — red for pre-existing reasons only.

## 🐛 Defect found and fixed by this slice's own harness

The ≥100-goroutine stress test exposed a **connection-pool deadlock**: the
refresh race-loser branch set `committed = true` ("nothing written") and
skipped its deferred rollback — but `BeginTx` had already checked out a
pooled connection in open-transaction state. Under sustained concurrency,
losing racers leaked connections until every `BeginTx` blocked forever.
Fixed with an explicit rollback before family revocation; the comment at the
site documents why, and the stress harness is the standing regression proof.
This is precisely the class of bug R15's KPI harness exists to catch.

## Deviations flagged (per AGENTS.md honesty rules)

1. **Constructor grew to 14 parameters** (techplan enumerated fewer): six
   login/session seams added positionally, nil-tolerant defaults documented.
2. **Lockout count signatures take `since time.Time`**, not a window
   duration — keeps hidden `time.Now()` out of the adapter; boundaries are
   caller-controlled and deterministically testable.
3. **Port method added beyond techplan §10 list**:
   `FindIdentifierHashByUserAndProvider` — needed for the Assumption-C
   identifier backfill.
4. **Assumption C does not cover Google-only MFA users** (no email_password
   identity exists). Backfill falls back to synthetic
   `sha256("mfa-stage:"+userID)`; deviation documented at
   `insertAttemptWithUser`. **Candidate assumption note for the feature
   spec** — requires your sign-off since specs are human-edited.
5. **Attempt-row write failure trade-off**: a lost attempt row can only
   undercount toward lockout, never lock anyone out spuriously — accepted,
   documented at `writeAttempt`.

## Testing checklist status

Techplan §12: **20/20 rule IDs covered** (R6/R9's full INV-account-06 proof
is seam-scoped until task #6 provides the real verifier — fake-tested here).
Named-test map lives in techplan §12; every entry has a passing
implementation except where explicitly marked #6-deferred.

## Risk note (root AGENTS.md §5)

- **Assumptions made:** Google-only MFA backfill fallback (above); empty-table
  semantics stand in for unenrolled-MFA/no-roles reality until #6/#8 ship;
  stub verifier fails closed so `/auth/login/mfa` cannot reach issuance.
- **Edge cases intentionally NOT handled:** multi-tab concurrent refresh UX
  (spec Assumption D — frontend BroadcastChannel); reverse-proxy IP
  sharing in the rate limiter (deferred follow-up); `login_attempts`
  retention (accepted, threat-model residual-risk entry #7).
- **Concurrency assumptions:** guarded UPDATE is the sole writer of
  `replaced_by_id`; rotate+child-insert share one transaction; race-loser ≡
  attacker (family revoked). Proof: `TestRefresh_ConcurrentRequests_ExactlyOneWins_RealDB`,
  `TestRefresh_Stress_MixedValidAndReplayed` under `-race`,
  `TestRefresh_ReuseDetection_FamilyRevoked_RealDB`.
- **What is not tested, and why:** real TOTP/backup-code verification
  (task #6 owns the crypto; seam fails closed); gitleaks stage (CLI absent);
  live SMTP delivery (v1 has no SMTP); browser-level cookie behavior
  (frontend track).

## Feature spec reference

Fulfills `docs/spec/1-account/features/03-login-session-management.md`
(endpoints `POST /auth/login`, `/auth/login/mfa`, `/auth/refresh`,
`/auth/logout` + Fitur 2C lockout), subject to the Tier 0 paired pass above
and the Assumption-C note awaiting your spec edit.
