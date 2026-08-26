# Testing Report — Login & Session Management (account #03)

> Ticket      : account domain task #3 — `docs/spec/1-account/features/03-login-session-management.md`
> Phase       : 5-testing (final sweep)
> Executed    : 2026-08-26
> Inputs      : `2-plan/techplan.md` (§4 Rules & Validation — R1–R20), `4-patch/after-code-review/report.md` (latest implementation report)
> Method      : Spot-check named tests → close report's own gaps → real-interface exercise (integration against real Postgres) + backward compat + migration collision + fresh techplan read
> Verdict     : **Pass**

---

## 0. Sweep Summary

**Confirmed (spot-checked, still passing):**

| Rule | Test(s) |
|---|---|
| R1 login happy path | `TestLogin_Success_NoMfa` (unit), `TestGetLoginUserView_AssemblesFields` (integration) |
| R2 MFA required, no tokens | `TestLogin_MfaRequired_NoTokensIssuedYet` (unit), `TestLoginHandler_MfaRequired_NoCookiesNoTokens` (transport) |
| R3 generic invalid creds, byte-identical | `TestLogin_GenericErrorMessage` (unit), `TestLoginHandler_GenericBodiesByteIdentical` (transport) |
| R4 password-stage lockout | `TestLogin_Lockout_5Failed15Min` (unit), `TestLogin_Lockout_EndToEnd` (integration), `TestInsertLoginAttempt_AndCountWindows` (integration) |
| R5 unverified login succeeds | `TestLogin_UnverifiedIdentity_Succeeds` |
| R6 pending-token shape | `TestMFAPending_RoundTrip`, `TestMFAPending_StrictExpiry_NoLeeway`, `TestMFAPending_WrongSecretRejected`, `TestMFAPending_CrossPurposeMatrix`, `TestMFAPending_MalformedAndValidation`, `TestAuthMiddleware_RejectsWrongSigningKey`, `TestAuthMiddleware_RejectsNonAccessPurposeToken` |
| R7 MFA-stage lockout | `TestLoginMfa_Lockout_5Failed15Min` (unit), `TestInsertLoginAttempt_AndCountWindows` (integration) |
| R8 TOTP completion (seam-scoped) | `TestLoginMfa_TotpSuccess_CompletesLogin` (fake verifier) |
| R9 backup-code completion (seam-scoped) | `TestLoginMfa_BackupCode_CompletesLogin` (fake verifier) |
| R10 invalid pending token | `TestLoginMfa_InvalidPendingToken`, `TestMFAPending_MalformedAndValidation` |
| R11 wrong MFA code | `TestLoginMfa_WrongCode` |
| R12 refresh rotation | `TestRefresh_Rotates_IssuesChild_SameFamily` (unit), `TestRotateRefreshToken_HappyPath` + `TestRotateRefreshToken_GuardsReturnFalse` (integration) |
| R13 refresh missing/expired | `TestRefresh_MissingOrExpiredCookie` (unit), `TestRefreshHandler_MissingCookie_Indistinguishable401` (transport) |
| R14 reuse detection, family revoked | `TestRefresh_ReuseDetection_FamilyRevoked` (unit), `TestRefresh_ReuseDetection_FamilyRevoked_RealDB` (integration), `TestRevokeRefreshTokenFamily_IncludesRotated` (integration) |
| R15 concurrent refresh, exactly one wins | `TestRefresh_ConcurrentRequests_ExactlyOneWins` (unit), `TestRefresh_ConcurrentRequests_ExactlyOneWins_RealDB` + `TestRefresh_Stress_MixedValidAndReplayed` (integration, 120 goroutines) |
| R16 logout idempotent 204 | `TestLogout_RevokesAndClears` + `TestLogout_NoCookie_Still204` (unit), `TestLogoutHandler_Idempotent204AndClears` (transport) |
| R17 purpose/key separation | `TestAuthMiddleware_RejectsWrongSigningKey`, `TestAuthMiddleware_RejectsNonAccessPurposeToken`, `TestMFAPending_CrossPurposeMatrix`, `TestAccessToken_AlgorithmConfusionRejected` |
| R18 timing discipline | `TestLogin_TimingShape_NoEarlyReturn` |
| R19 log hygiene | `TestLogin_LoggingNeverLeaksCredentials` (unit), `TestSessionEndpoints_LogLeakSweep` (transport) |
| R20 data-access discipline | Grep gate (no `fmt.Sprintf` in SQL); integration tests prove query validity; migration round-trip verified |
| Q1 patch fail-open | `TestWriteAttempt_FailOpen_ValidLoginStillSucceeds`, `TestWriteAttempt_FailOpen_InvalidCredentialsStillRejected` |

**Closed from report's own gap list:**
- "Fail-open path under real Postgres" — justification confirmed: the fail-open decision lives at the service layer (`writeAttempt`), the unit test with `failingAttemptRepo` proves the contract. Integration-level injection would test pgx error behavior, not the fail-open decision. No additional test needed.
- "Real TOTP/backup-code verification" — confirmed out of scope: `stubMfaVerifier` fails closed (returns `false, nil`); task #6 owns the crypto. No path to token issuance through `/auth/login/mfa` until #6 supplies the real verifier.

**Not carried over — required fresh testing:**
- Backward compatibility (OAuth tokens without `purpose` claim) — not in the report's gap list but required by the checklist. Verified fresh: `GoogleTokenVerifier` (`auth_google.go:51`) uses `jwt.RegisteredClaims` (no `purpose` check), completely separate from `VerifyAccessToken` (`token.go:109`) which uses `purposeClaims`. OAuth-issued tokens still work for link/reauth. Confirmed.
- Migration collision check — not in the report's gap list. Verified fresh: on-disk 000001–000009 sequential, no gaps; DB state version 9, not dirty; no migrations landed since techplan was written (git log confirms 000006–000009 are new in this slice).

## 1. Test Coverage

| Rule/Scenario | Category | Real-interface test performed | Result |
|---|---|---|---|
| R1 — login happy path, no MFA | Happy | `TestLogin_Success_NoMfa` (unit — attempt row, TTLs 15m/30d, fresh family, cookie attrs); `TestGetLoginUserView_AssemblesFields` (integration real DB — decrypt-on-read, provider aggregation, roles, MFA flag) | PASS |
| R2 — MFA required, no tokens issued | Negative | `TestLogin_MfaRequired_NoTokensIssuedYet` (unit — no Set-Cookie, no refresh row, pending token present); `TestLoginHandler_MfaRequired_NoCookiesNoTokens` (transport — no cookies in response) | PASS |
| R3 — generic invalid creds, byte-identical | Negative | `TestLogin_GenericErrorMessage` (unit — wrong-email sentinel == wrong-password sentinel); `TestLoginHandler_GenericBodiesByteIdentical` (transport — 401/429 bodies byte-equal modulo status+type URI, detail "Email atau password salah.") | PASS |
| R4 — password-stage lockout, pre-credential, no row | Negative/edge | `TestLogin_Lockout_5Failed15Min` (unit — 4 failures pass, 5th rejected, no compare call, no attempt row); `TestLogin_Lockout_EndToEnd` (integration, real DB — threshold + window-expiry release); `TestInsertLoginAttempt_AndCountWindows` (integration — count query for both stages) | PASS |
| R5 — unverified identity logs in | Edge | `TestLogin_UnverifiedIdentity_Succeeds` (verified_at=nil → login succeeds) | PASS |
| R6 — pending-token shape (HS256, purpose, 5min, no leeway) | Negative | `TestMFAPending_RoundTrip`, `TestMFAPending_StrictExpiry_NoLeeway`, `TestMFAPending_WrongSecretRejected`, `TestMFAPending_CrossPurposeMatrix`, `TestMFAPending_MalformedAndValidation`, `TestAuthMiddleware_RejectsWrongSigningKey`, `TestAuthMiddleware_RejectsNonAccessPurposeToken` | PASS |
| R7 — MFA-stage lockout, user_id-keyed, pre-code | Negative | `TestLoginMfa_Lockout_5Failed15Min` (unit — no attempt row on rejection); `TestInsertLoginAttempt_AndCountWindows` (integration — user_id+stage count query) | PASS |
| R8 — TOTP completion → issuance like R1 | Happy | `TestLoginMfa_TotpSuccess_CompletesLogin` (fake verifier — success attempt row + tokens issued). Seam-scoped until #6; flagged. | PASS (seam-scoped) |
| R9 — backup-code completion, single-use | Happy/edge | `TestLoginMfa_BackupCode_CompletesLogin` (fake verifier — tx passed to VerifyBackupCode, success row + tokens). INV-account-06 single-use proof is #6's scope; flagged. | PASS (seam-scoped) |
| R10 — invalid pending token, no writes | Negative | `TestLoginMfa_InvalidPendingToken` (no attempt rows, no refresh rows); `TestMFAPending_MalformedAndValidation` (expired/malformed/wrong-key variants) | PASS |
| R11 — wrong code, attempt row, 401 | Negative | `TestLoginMfa_WrongCode` (mfa failure attempt row, no tokens, ErrInvalidCredentials) | PASS |
| R12 — refresh rotation, guarded UPDATE + child in one tx | Happy | `TestRefresh_Rotates_IssuesChild_SameFamily` (unit — parent marked, child same family); `TestRotateRefreshToken_HappyPath` + `TestRotateRefreshToken_GuardsReturnFalse` (integration — real DB guarded UPDATE, not-found/already-rotated/revoked/expired all return false) | PASS |
| R13 — refresh missing/expired → 401 | Negative/edge | `TestRefresh_MissingOrExpiredCookie` (unit — empty plain + expired both ErrInvalidCredentials, family revoked on expired); `TestRefreshHandler_MissingCookie_Indistinguishable401` (transport — missing cookie → 401) | PASS |
| R14 — reuse detection, family revoked (A→B→C replay-A) | Negative | `TestRefresh_ReuseDetection_FamilyRevoked` (unit — A,B,C all revoked, C dead); `TestRefresh_ReuseDetection_FamilyRevoked_RealDB` (integration — same chain against real DB); `TestRevokeRefreshTokenFamily_IncludesRotated` (integration — family revoke includes already-rotated descendants) | PASS |
| R15 — concurrent refresh, exactly one wins, race-loser ≡ attacker | Edge/concurrency | `TestRefresh_ConcurrentRequests_ExactlyOneWins` (unit, 8 racers); `TestRefresh_ConcurrentRequests_ExactlyOneWins_RealDB` (integration); `TestRefresh_Stress_MixedValidAndReplayed` (integration, 120 goroutines mixed valid+replayed, under `-race`) | PASS under `-race` |
| R16 — logout idempotent 204 | Happy/negative | `TestLogout_RevokesAndClears` + `TestLogout_NoCookie_Still204` (unit — present/absent/never-issued all nil error); `TestLogoutHandler_Idempotent204AndClears` (transport — 204 + cookie cleared) | PASS |
| R17 — purpose/key separation at verification | Negative | `TestAuthMiddleware_RejectsWrongSigningKey`, `TestAuthMiddleware_RejectsNonAccessPurposeToken`, `TestMFAPending_CrossPurposeMatrix`, `TestAccessToken_AlgorithmConfusionRejected` (cross-purpose rejection matrix) | PASS |
| R18 — timing discipline, no early return | Edge | `TestLogin_TimingShape_NoEarlyReturn` (structural — every branch burns exactly 1 compare; lockout skips entirely by design) | PASS |
| R19 — log hygiene, no secrets in logs | Negative | `TestLogin_LoggingNeverLeaksCredentials` (unit — marker sweep for password/refresh/access token); `TestSessionEndpoints_LogLeakSweep` (transport — sweep across all four handlers) | PASS |
| R20 — data-access discipline (goqu, additive migrations) | — | Grep gate: no `fmt.Sprintf` in SQL context in new files; all goqu `Prepared(true)`; integration tests prove query validity; migration round-trip (up→down→up) verified per build report | PASS |
| Q1 patch — fail-open on audit-write failure | Negative/dependency | `TestWriteAttempt_FailOpen_ValidLoginStillSucceeds` (valid login + audit failure → login succeeds); `TestWriteAttempt_FailOpen_InvalidCredentialsStillRejected` (wrong password + audit failure → still ErrInvalidCredentials) | PASS |

**Rules that can't be exercised through the real interface:** R8/R9 full INV-account-06 proof (TOTP/backup-code single-use against real encrypted secrets) — the `mfaVerifier` seam fails closed until task #6 supplies the real verifier. This is a deliberate deferral, not a gap; the seam contract is fake-tested and the real implementation is #6's scope.

## 2. Error Verification

| Error case | Expected category | Actual category | Message actionable? | Propagation correct? |
|---|---|---|---|---|
| Wrong email | 401 `errors/invalid-credentials` | 401, "Email atau password salah." | Yes (user knows to retry) | Yes — `ErrInvalidCredentials` → `MapServiceError` → `WriteProblem` |
| Wrong password | 401 `errors/invalid-credentials` | 401, identical body to wrong-email | Yes (byte-identical, anti-enumeration) | Yes — same sentinel, same path |
| Lockout (≥5 fails / 15min) | 429 `errors/too-many-requests` | 429, same title/detail as 401, different type URI | Yes (status distinguishes; detail is generic by design) | Yes — `ErrLockedOut` → `MapServiceError` → 429 Problem |
| Invalid pending token | 401 `errors/invalid-credentials` | 401, same body as invalid creds | Yes | Yes — `ErrMfaPendingInvalid` → `MapServiceError` → 401 |
| Wrong MFA code | 401 `errors/invalid-credentials` | 401 | Yes | Yes — `ErrInvalidCredentials` |
| Missing/expired/revoked/reused refresh | 401 (indistinguishable) | 401, one body for all four classes | Yes (user re-logs in) | Yes — all collapse to `ErrInvalidCredentials` |
| Malformed JSON body | 400 `problems/invalid-json` | 400, "The request body was not valid JSON." | Yes (caller fixes request) | Yes — `write400InvalidJSON` → `WriteProblem` |
| Neither/both MFA codes | 422 validation | 422, field-level errors | Yes (tells caller exactly-one required) | Yes — `ErrValidation` → `WriteValidationError` |
| Unhandled service error | 500 generic | 500, generic Problem (no internals) | No (by design — no leak) | Yes — `MapServiceError` default → `log.Printf` server-side + generic 500 |

## 3. Final Verification

- [x] **Target repo build/lint/test commands:**
  ```
  $ go build ./...
  (no output — clean)

  $ go test -race -count=1 ./...
  ok  github.com/anhsbolic/kencleng/backend/internal/domain/account   124.194s
  ok  github.com/anhsbolic/kencleng/backend/internal/platform/auth      1.051s
  ok  github.com/anhsbolic/kencleng/backend/internal/platform/breachcheck (cached)
  ok  github.com/anhsbolic/kencleng/backend/internal/platform/crypto     (cached)
  ok  github.com/anhsbolic/kencleng/backend/internal/platform/googleoauth (cached)
  ok  github.com/anhsbolic/kencleng/backend/internal/platform/notification (cached)
  ok  github.com/anhsbolic/kencleng/backend/internal/platform/secrets    (cached)
  ok  github.com/anhsbolic/kencleng/backend/internal/transport/http      1.046s
  PASS

  $ go test -v -count=1 -tags=integration ./internal/domain/account/
  (real Postgres on localhost:5435)
  --- PASS: TestRefresh_ConcurrentRequests_ExactlyOneWins_RealDB (0.14s)
  --- PASS: TestRefresh_Stress_MixedValidAndReplayed (0.32s)
  --- PASS: TestRefresh_ReuseDetection_FamilyRevoked_RealDB (0.04s)
  --- PASS: TestLogin_Lockout_EndToEnd (0.69s)
  --- PASS: TestRotateRefreshToken_HappyPath
  --- PASS: TestRotateRefreshToken_GuardsReturnFalse
  --- PASS: TestRevokeRefreshTokenFamily_IncludesRotated
  --- PASS: TestInsertLoginAttempt_AndCountWindows
  --- PASS: TestGetLoginUserView_AssemblesFields
  --- PASS: TestGetLoginUserView_NonExistentUser
  [... all integration tests pass ...]
  ok  github.com/anhsbolic/kencleng/backend/internal/domain/account  12.165s

  $ go test -race -count=1 -tags=integration ./internal/domain/account/
  ok  github.com/anhsbolic/kencleng/backend/internal/domain/account  125.276s

  $ staticcheck ./...
  (no output — clean)
  ```

- [x] **Migration/schema version collision check:** On-disk `000001`–`000009` sequential, no gaps. DB `schema_migrations` = version 9, dirty=false. No migrations landed since techplan was written (git log confirms 000006–000009 are new in this slice). No collision.
  ```
  $ psql "$DATABASE_URL" -c "SELECT version, dirty FROM schema_migrations;"
   version | dirty
  ---------+-------
         9 | f
  (1 row)
  ```

- [x] **Backward compatibility:** Explicitly verified, not assumed. `GoogleTokenVerifier` (`auth_google.go:51`) uses `jwt.RegisteredClaims` (no `purpose` check), separate from `VerifyAccessToken` (`token.go:109`, uses `purposeClaims` with `purpose:"access"` enforcement). OAuth-issued tokens without `purpose` still work for link/reauth gating. `kencleng_refresh` cookie name + attributes unchanged. Register/verify/resend/OAuth handlers and routes unchanged.

- [x] **Fresh end-to-end techplan read:** No contradictions found. The fail-open change (Q1 patch) aligns with build-report deviation #5's accepted trade-off. The TTL dedup (Q2 patch) is a behavioral no-op (15m → 15m, single source of truth). The logout doc exception (Q3 patch) clarifies R16 (204 for idempotent cases, 500 for genuine infra failure) without contradicting it.

## 4. New Bug Patterns

No new pattern — see `examples.md` for handling this ticket-specific detail directly. The fail-open change is an instance of the "Silent skip when an error was expected" pattern, but it's a deliberate design decision (not a bug) with two explicit tests proving the contract. The connection-leak fix found during the build phase (race-loser skipping rollback) is a one-off, not a category that would recur in unrelated features.

## Verdict

**Pass.**

All 20 techplan §4 rules confirmed with named passing tests. All four categories covered (happy, negative, edge, backward-compat). Integration suite green under `-race` against real Postgres (120-goroutine stress test passes). Error responses verified for correct category, actionable messages, and propagation through `MapServiceError`. Migration state clean (version 9, no collision). Backward compatibility explicitly verified (OAuth legacy verifier separate from strict access-token verifier).

**Flagged follow-ups (non-blocking, not testing-phase scope):**

1. **Tier 0 paired rewrite pass** (techplan Resolved #13) — must complete before commit. Covers `platform/auth/token.go`, `repository_db.go` rotation methods, `login.go` reuse/race-loser branch. This is a human gate, not a testing finding.
2. **R8/R9 full INV-account-06 proof** — seam-scoped until task #6 supplies the real TOTP/backup-code verifier. The seam fails closed; no path to token issuance until #6.
3. **CSRF second layer** — accepted residual risk (techplan Resolved #7). Revisit trigger: frontend API client landing.
