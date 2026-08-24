# Testing Report — 02-google-oauth-login-register

> Ticket    : 02-google-oauth-login-register
> Feature   : Google OAuth login/register (account domain, Fitur 1B)
> Tester    : GLM 5.2 (max)
> Date      : 2026-08-24
> Diff ref  : commit `ef428be` (backend scope)
> Sources   : `2-plan/techplan.md`, `3-build/report.md`, `4-code-review/report.md`
> Conventions: `backend/AGENTS.md`, `backend/README.md`, root `AGENTS.md`,
>              `workflow/5-testing/{guidelines,checklist,examples}.md`

---

## 0. Sweep Summary

### Confirmed (spot-checked, still passing)

Every R1–R26 named test from the implementation report was run and passes.
Full detail in §1.

| Rule | Named test(s) | Status |
|---|---|---|
| R1 (login redirect, no auth) | `TestGoogleRedirect_Login_NoSessionRequired` (svc) + `TestGoogleRedirect_LoginNeedsNoAuth` (handler) | ✅ |
| R2 (link/reauth 401 pre-redirect) | `TestGoogleRedirect_LinkReauthRequireAuth` (handler) + `TestGoogleRedirect_LinkReauthWithoutSessionRejectedDefensively` (svc) | ✅ |
| R3 (user_id into cookie) | `TestGoogleRedirect_LinkWithSessionEncodesUserID` | ✅ |
| R4 (state mismatch, no Google call) | `TestGoogleCallback_StateMismatch_NoGoogleCall` (zero-call assertion) | ✅ |
| R5 (nonce mismatch → nonce_mismatch) | `TestVerifyIDToken_NonceMismatchIsDistinguishable` (platform) + `TestGoogleCallback_NonceMismatchIsReplayNotForgery` (domain) | ✅ |
| R6 (timeout → google_unavailable) | `TestExchangeCode_TimeoutReturnsCleanError` + `TestGoogleCallback_GoogleUnavailableMappedCleanly` | ✅ |
| R7 (login existing identity) | `TestGoogleCallback_Login_ExistingGoogleIdentityIssuesTokens` | ✅ |
| R8 (login new user, tx) | `TestGoogleCallback_Login_NewUserCreatesVerifiedIdentity` | ✅ |
| R9 (no-auto-merge login) | `TestGoogleCallback_NoAutoMerge_Login` | ✅ |
| R10 (no-auto-merge link) | `TestGoogleCallback_NoAutoMerge_Link` | ✅ |
| R11 (link attach + audit) | `TestGoogleCallback_LinkSuccess_AttachesAndAudits` | ✅ |
| R12 (reauth marker) | `TestGoogleCallback_Reauth_NoSideEffects` (svc) + `TestGoogleCallback_ReauthSetsMarker` (handler) | ✅ |
| R13 (fixed redirect_uri) | `TestExchangeCode_SendsConfiguredParamsAndParsesResponse` | ✅ |
| R14 (verified_at=now at insert) | asserted in new-user + link-success tests | ✅ |
| R15 (concurrent duplicate — login only) | `TestGoogleCallback_ConcurrentDuplicateRegistration_Race` (12 goroutines, `-race`) | ✅ for login |
| R16 (no secrets in logs) | `TestGoogleOAuth_LogsNeverCarrySecrets` (svc) + `TestGoogleCallback_LogsNeverCarryTokens` (handler) | ✅ |
| R17 (sanitized Google-client errors) | `TestExchangeCode_Non200_SanitizesErrorAndLog` | ✅ |
| R18 (invalid intent 400) | `TestGoogleRedirect_InvalidIntentRejected` + `TestGoogleRedirect_InvalidIntent400` | ✅ |
| R19 (missing code/state) | `TestGoogleCallback_MissingParamsOrBadCookie_StateMismatch` + `TestGoogleCallback_MissingInputsStillHandled` | ✅ |
| R20 (expired/missing cookie) | same two tests (corrupt/empty cookie cases) | ✅ |
| R21 (JWKS refresh-on-miss) | `TestVerifyIDToken_JWKSRefreshOnMiss` (fetch-counted) | ✅ |
| R22 (explicit timeout) | `TestExchangeCode_TimeoutReturnsCleanError` + code-review confirmation | ✅ |
| R23 (constant-time compare) | code-review confirmed `subtle.ConstantTimeCompare` for state + nonce | ✅ |
| R24 (cookie attributes) | `TestGoogleRedirect_LoginNeedsNoAuth` asserts HttpOnly/Secure/Lax/MaxAge=600/Path | ✅ |
| R25 (inline verify, platform/auth untouched) | `TestGoogleRedirect_LinkWithBearerTokenPassesUserID` + `TestGoogleRedirect_TamperedTokenRejected` | ✅ |
| R26 (forgery generic, only nonce special) | `TestVerifyIDToken_GenericFailuresAreNotNonceMismatch` (8-case table incl. HS256 confusion) | ✅ |

### Closed from report's own gap list

- **`make verify` tooling** → ran `staticcheck`, `gosec`, `govulncheck`
  locally (gitleaks not installed). staticcheck: 2 findings in test files
  (style nits — `time.Until` S1024, unused const U1000). gosec: 13 issues,
  only 1 new from this feature (G706 log-injection at `auth_google.go:166`
  — **false positive**: `intent` is constrained to `"link"`/`"reauth"` by
  the if-guard at line 162, so the value is not arbitrary user input at
  the log site). govulncheck: 24 stdlib vulns (Go 1.24.8 vs 1.24.9 fix —
  project-wide, not feature-specific).
- **`email_verified` enforcement** → confirmed NOT checked in
  `VerifyIDToken` (`client.go:208-249`); properly flagged as an
  assumption in report §7, pending human sign-off. Techplan R26 lists
  only sig/iss/aud/exp/nonce — no contradiction.
- **Integration-tagged Postgres round-trips** → still not written
  (acknowledged gap; pattern-identical to `InsertAuthToken`; migration
  round-trip + FK cascade were verified live per report §5).

### Not carried over — required fresh testing

- **R15 link path** — the report claims R15 is covered by
  `TestGoogleCallback_ConcurrentDuplicateRegistration_Race`, but that
  test only exercises `intentLogin` (line 543). The code review (S1)
  found that `callbackLink` (`google_oauth.go:450`) does NOT check
  `isUniqueViolation` the way `registerGoogleUser` does (line 379).
  **Fresh testing confirmed this is a real bug** — see §1.
- **Backward compatibility** — not unit-testable; required real-interface
  verification through main.go route inspection (done — existing routes
  untouched at lines 131-133).

---

## 1. Test Coverage

| Rule/Scenario | Category | Real-interface test performed | Result |
|---|---|---|---|
| R1 (login redirect, no auth) | happy | Ran `TestGoogleRedirect_Login_NoSessionRequired` + `TestGoogleRedirect_LoginNeedsNoAuth` | ✅ confirmed |
| R2 (link/reauth 401) | negative | Ran `TestGoogleRedirect_LinkReauthRequireAuth` (handler 401 + no service call) | ✅ confirmed |
| R3 (user_id into cookie) | happy | Ran `TestGoogleRedirect_LinkWithSessionEncodesUserID` | ✅ confirmed |
| R4 (state mismatch, no Google call) | negative | Ran `TestGoogleCallback_StateMismatch_NoGoogleCall` (zero-call assertion) | ✅ confirmed |
| R5 (nonce mismatch → nonce_mismatch) | negative | Ran `TestVerifyIDToken_NonceMismatchIsDistinguishable` + `TestGoogleCallback_NonceMismatchIsReplayNotForgery` | ✅ confirmed |
| R6 (timeout → google_unavailable) | negative | Ran `TestExchangeCode_TimeoutReturnsCleanError` + `TestGoogleCallback_GoogleUnavailableMappedCleanly` | ✅ confirmed |
| R7 (login existing identity) | happy | Ran `TestGoogleCallback_Login_ExistingGoogleIdentityIssuesTokens` | ✅ confirmed |
| R8 (login new user) | happy | Ran `TestGoogleCallback_Login_NewUserCreatesVerifiedIdentity` (R14 verified_at asserted) | ✅ confirmed |
| R9 (no-auto-merge login) | negative | Ran `TestGoogleCallback_NoAutoMerge_Login` (no records created) | ✅ confirmed |
| R10 (no-auto-merge link) | negative | Ran `TestGoogleCallback_NoAutoMerge_Link` (no identity attached) | ✅ confirmed |
| R11 (link attach + audit) | happy | Ran `TestGoogleCallback_LinkSuccess_AttachesAndAudits` (action_type=account_linking asserted) | ✅ confirmed |
| R12 (reauth marker) | happy | Ran `TestGoogleCallback_Reauth_NoSideEffects` + `TestGoogleCallback_ReauthSetsMarker` | ✅ confirmed |
| R13 (fixed redirect_uri) | happy | Ran `TestExchangeCode_SendsConfiguredParamsAndParsesResponse` | ✅ confirmed |
| R14 (verified_at=now) | edge | Asserted in R8 + R11 tests (verified_at != nil) | ✅ confirmed |
| R15 (concurrent duplicate — login) | edge | Ran `TestGoogleCallback_ConcurrentDuplicateRegistration_Race` (12 goroutines, `-race`, login only) | ✅ confirmed for login |
| **R15 (concurrent duplicate — link)** | edge | **Wrote repro test**: 2 concurrent link operations → loser returns `account: insert linked identity: : (SQLSTATE 23505)` non-nil error → handler maps to **500** instead of clean 302 | ❌ **BUG CONFIRMED (S1)** |
| R16 (no secrets in logs) | negative | Ran `TestGoogleOAuth_LogsNeverCarrySecrets` + `TestGoogleCallback_LogsNeverCarryTokens` | ✅ confirmed |
| R17 (sanitized errors) | negative | Ran `TestExchangeCode_Non200_SanitizesErrorAndLog` (body-marker leak check) | ✅ confirmed |
| R18 (invalid intent 400) | negative | Ran `TestGoogleRedirect_InvalidIntent400` (400 + Problem Details) | ✅ confirmed |
| R19 (missing code/state) | edge | Ran `TestGoogleCallback_MissingParamsOrBadCookie_StateMismatch` + `TestGoogleCallback_MissingInputsStillHandled` | ✅ confirmed |
| R20 (expired/missing cookie) | edge | Same tests (corrupt/empty cookie cases) | ✅ confirmed |
| R21 (JWKS refresh-on-miss) | edge | Ran `TestVerifyIDToken_JWKSRefreshOnMiss` (exactly one refetch) | ✅ confirmed |
| R22 (explicit timeout) | negative | Ran `TestExchangeCode_TimeoutReturnsCleanError` + code review confirmed `http.Client{Timeout: 10s}` | ✅ confirmed |
| R23 (constant-time compare) | negative | Code review confirmed `subtle.ConstantTimeCompare` for state + nonce | ✅ confirmed |
| R24 (cookie attributes) | edge | `TestGoogleRedirect_LoginNeedsNoAuth` asserts HttpOnly/Secure/Lax/MaxAge=600/Path | ✅ confirmed |
| R25 (inline verify) | negative | Ran `TestGoogleRedirect_LinkWithBearerTokenPassesUserID` + `TestGoogleRedirect_TamperedTokenRejected` | ✅ confirmed |
| R26 (forgery generic) | negative | Ran `TestVerifyIDToken_GenericFailuresAreNotNonceMismatch` (8-case table incl. HS256 confusion) | ✅ confirmed |
| Backward compat | backward-compat | Verified `main.go:131-133` existing routes untouched; full suite (`go test ./...`) passes including all pre-existing register/verify-email/resend tests; `service_test.go` changes additive only | ✅ confirmed |
| Cookie cleared after callback | edge | Ran `TestGoogleCallback_ClearsStateCookieOnEveryOutcome` (error + success) | ✅ confirmed |
| All 6 error codes verbatim | negative | Ran `TestGoogleCallback_ErrorCodesSurfaceVerbatim` (all six codes) | ✅ confirmed |

### §4 rules that can't be exercised through the real interface

None — all 26 rules are exercisable and covered. R15's link path is the
exception, and it's a confirmed bug (not a missing entry point).

---

## 2. Error Verification

| Error case | Expected category | Actual category | Message actionable? | Propagation correct? |
|---|---|---|---|---|
| Invalid intent (R18) | 400 client error | 400 + `application/problem+json` | ✅ "The intent parameter must be one of: login, link, reauth." | ✅ `WriteProblem` → Problem Details |
| Link/reauth without session (R2) | 401 client error | 401 + Problem Details | ✅ "Sign in before continuing." | ✅ `WriteProblem`, no Google redirect |
| State mismatch (R4) | 302 redirect with `?error=state_mismatch` | 302 + `?error=state_mismatch` | ✅ code is actionable | ✅ no Google API call, no leak |
| Nonce mismatch (R5) | 302 with `?error=nonce_mismatch` | 302 + `?error=nonce_mismatch` | ✅ distinguishable from forgery | ✅ no state change |
| Google timeout (R6) | 302 with `?error=google_unavailable` | 302 + `?error=google_unavailable` | ✅ actionable (retry) | ✅ raw error stays internal, no leak |
| No-auto-merge login (R9) | 302 with `?error=google_email_conflict` | 302 + `?error=google_email_conflict` | ✅ actionable | ✅ no records created, no tokens |
| Link conflict (R10) | 302 with `?error=google_link_conflict` | 302 + `?error=google_link_conflict` | ✅ actionable | ✅ no identity attached |
| Forged id_token (R26) | 302 with `?error=google_token_invalid` | 302 + `?error=google_token_invalid` | ✅ actionable | ✅ no state change |
| Internal service error | 500 server error | 500 + Problem Details | ✅ generic by design ("An unexpected error occurred.") | ✅ `WriteProblem`, raw error not in body (`TestGoogleCallback_ServiceErrorIsGeneric500` confirms no "context deadline" leak) |
| **Concurrent link race (S1 bug)** | 302 with `?error=google_link_conflict` (or idempotent success) | **500 Problem Details** | ❌ wrong category — user sees "internal error" instead of recoverable conflict | ❌ `callbackLink` returns raw `isUniqueViolation` error → handler maps to 500 |

---

## 3. Final Verification

- [x] **Target repo build/lint/test commands:**
  - `go build ./...` — clean ✅
  - `go vet ./...` — clean ✅
  - `gofmt -l internal/ cmd/` — clean ✅
  - `go test -count=1 ./...` — all pass ✅
  - `go test -race -count=1 ./...` — all pass (platform 1.3s, account 65s, transport 1.0s) ✅
  - `go test -tags=contract ./...` — all pass ✅
  - `staticcheck ./...` — 2 findings (test-file style nits: `time.Until` S1024, unused const U1000) — non-blocking
  - `gosec ./...` — 13 issues; only 1 new from this feature (G706 at `auth_google.go:166` — **false positive**, intent constrained by if-guard)
  - `govulncheck ./...` — 24 stdlib vulns (Go 1.24.8 → 1.24.9 fix; project-wide, not feature-specific)
  - `gitleaks` — not installed (gap acknowledged in report §7)

  ```
  $ go test -race -count=1 ./...
  ok      github.com/anhsbolic/kencleng/backend/internal/platform/googleoauth   1.267s
  ok      github.com/anhsbolic/kencleng/backend/internal/domain/account         65.047s
  ok      github.com/anhsbolic/kencleng/backend/internal/transport/http         1.022s
  ok      github.com/anhsbolic/kencleng/backend/internal/platform/breachcheck   0.015s
  ok      github.com/anhsbolic/kencleng/backend/internal/platform/crypto        0.004s
  ok      github.com/anhsbolic/kencleng/backend/internal/platform/notification  0.004s
  ok      github.com/anhsbolic/kencleng/backend/internal/platform/secrets       0.744s
  ```

- [x] **Migration/schema version collision check:** Migrations
  000001–000005, sequential, no collision. 000004 (`refresh_tokens`)
  and 000005 (`user_logs`) are new, additive, each with matching down
  migration. ✅

- [x] **Backward compatibility:** Explicitly verified, not assumed.
  `main.go:131-133` existing routes (`POST /auth/register`, `POST
  /auth/verify-email`, `POST /auth/verify-email/resend`) untouched; new
  Google routes (139-142) added alongside under same RateLimit
  middleware. `NewService` signature extended but all callers updated.
  `service_test.go` fakeRepo changes additive only (no existing
  assertions changed — code review §4 confirmed). Full pre-existing
  test suite passes. ✅

- [x] **Fresh end-to-end techplan read:** One contradiction/gap found:
  - **R15 coverage overclaim**: The rule says "Given two concurrent
    Google registrations for the same email... the unique index fails
    one cleanly." The named test only covers `intentLogin`. The **link
    path** (`callbackLink`) hits the same unique index but does NOT
    handle `isUniqueViolation` — a concurrent link race surfaces as 500
    instead of a clean 302. This is not a techplan contradiction per se
    (R15 says "registrations" = login), but it IS an unstated gap: no
    rule covers concurrent link duplicates, and the implementation
    doesn't handle it.

---

## 4. New Bug Patterns

The S1 bug matches an existing pattern in `examples.md`: **"Wrong error
category returned (e.g. server error instead of client error)"** — "Both
surface as 'an error happened'" / "Verify the specific error code/category
through the actual error-handling layer." The S1 finding is a concrete
instance of this: a concurrent link race returns 500 (server error) where
a 302 `google_link_conflict` (client-recoverable) was intended. The
`examples.md` entry already covers the general category.

However, there is a **subtler variant** worth noting: the same unique-index
constraint applies to two code paths (login and link), but only one path
(login) was given the `isUniqueViolation` handling. This is a "copy-paste
asymmetry" pattern — when two paths share the same DB constraint, the
error handling must be mirrored, not just written once. This could
plausibly recur in unrelated features where a unique index is hit from
multiple insert sites.

**Recommendation — add to `examples.md`:**

| Bug Pattern | How It Hides | How to Catch It |
|---|---|---|
| Unique-index handling mirrored on one path but not its sibling | Both paths pass single-threaded tests; the unhandled path only fails under concurrent load | When two code paths insert against the same unique index, test BOTH for concurrent-duplicate handling, not just the first one written |

This meets the threshold (a category of mistake — "asymmetric error
handling across sibling paths sharing a constraint" — that could recur in
any feature with multiple insert sites against the same index).

---

## Verdict

**Fail — send back to implementation.**

### Blocking findings

1. **S1 (concurrent link race → 500 instead of clean 302)** —
   `callbackLink` (`google_oauth.go:450`) does NOT check
   `isUniqueViolation` the way `registerGoogleUser` (`google_oauth.go:379`)
   does. A concurrent link operation for the same Google email (same
   user, two browser tabs) returns a raw unique-violation error that the
   handler maps to 500 Problem Details instead of a clean
   `google_link_conflict` 302 or idempotent success. **Proven by two repro
   tests** (one concurrent, one deterministic — both removed after
   confirmation; they were for verification only, not permanent tests).
   This is not security-critical (the unique index prevents data
   corruption; retry recovers), but the **error category is wrong** —
   matching the `examples.md` "wrong error category" pattern. R15's
   coverage claim is partially false: the named test only covers login,
   not link.

   **Fix:** Mirror the `registerGoogleUser` pattern in `callbackLink`:
   check `isUniqueViolation(err)` after the link insert; on hit, roll
   back the tx, re-lookup the winner identity, and either return
   `google_link_conflict` (if the winner belongs to a different user) or
   treat as idempotent success (if the winner is the same session user).
   Add a `-race` test mirroring the login duplicate-race test for the
   link path.

### Optional follow-ups (from code review, non-blocking)

- **Q1** (TTL constant duplication across domain/transport) —
  load-bearing for cookie/JWT-expiry consistency; add compile-time
  assertion or single definition point.
- **Q2** (`failResult` operator precedence readability) — factor the
  condition.
- **Q3** (case-sensitive Bearer prefix match) — RFC 6750 conformance
  nit.
- **Q4** (error code not URL-encoded in redirect) — future-proof
  against non-`[a-z_]` codes.

### Trace

The blocking finding (S1) traces to a **genuinely new gap** — it was
found by code review (4-code-review S1), not by a Step-0-confirmed area
that regressed. The 4-patch directory is empty, meaning the code review's
findings were never addressed. The implementation report's R15 coverage
claim was an overclaim (test only covers login, not link), which this
testing phase caught through real-interface exercise.
