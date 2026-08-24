# Task 03 — Account Domain: Data Layer + OAuth Service (GoogleRedirect / GoogleCallback / IssueTokens)

> Ticket    : 02-google-oauth-login-register
> Sub-task  : 3 of 4
> Axis      : Dependency/sequence chain (primary) + component/layer alignment
> Status    : Ready (after Tasks 01 + 02)
> Back-ref  : `../techplan.md` — "Tech Plan: Google OAuth Login/Register" (originating contract techplan; cross-check high-level decisions there whenever needed). **Read §1 Background and §8 business flow there before writing any branch logic.**

---

## 1. Scope

The business-rule heart of this ticket: account-domain entities and
repository methods for the two new tables, the token-issuance primitives,
and the intent-branched Google OAuth service methods (login / link /
reauth). The data layer is thin here (2 structs, 2 inserts) which is why it
shares a task with the service.

**In scope:**
- `entity.go` — add RefreshToken + UserLog structs
- `repository.go` — add InsertRefreshToken + InsertUserLog to interface
- `repository_db.go` — implement both (goqu Insert, InsertAuthToken pattern)
- `service.go` — add `googleOAuth` + `authKeys` dependencies; update NewService
- `google_oauth.go` (new) — GoogleRedirect, GoogleCallback, IssueTokens, CallbackResult
- `google_oauth_test.go` (new) — unit tests: fakeRepo + fake Google client

**Out of scope:**
- Cookie writing, HTTP handlers, reauth marker store, route wiring — Task 04.
  This task's methods return values (`cookieValue`, `CallbackResult`); the
  transport layer turns them into Set-Cookie headers and 302 responses.
- Login endpoint, MFA, lockout, refresh rotation/reuse detection, logout —
  task #3 (builds on IssueTokens primitives created here).
- Unlink/set-password (task #05), MFA-disable consuming reauth markers
  (task #06).

## 2. Dependencies

- **Hard deps:** Task 01 (tables + jwt dep), Task 02 (`*googleoauth.Client`
  to hold; its `ErrNonceMismatch` to map; fake mirrors its surface for tests).
- **Blocks:** Task 04 (handlers call these methods; NewService signature).

## 3. Files

| File | Change Type |
|---|---|
| `backend/internal/domain/account/entity.go` | Edit |
| `backend/internal/domain/account/repository.go` | Edit |
| `backend/internal/domain/account/repository_db.go` | Edit |
| `backend/internal/domain/account/service.go` | Edit |
| `backend/internal/domain/account/google_oauth.go` | New |
| `backend/internal/domain/account/google_oauth_test.go` | New |

## 4. Implementation detail

### `entity.go` (techplan §10)

- Add `RefreshToken` struct: ID, UserID, FamilyID, TokenHash, ExpiresAt,
  RevokedAt, ReplacedByID, CreatedAt.
- Add `UserLog` struct: ID, UserID, ActionType, CreatedAt.

Doc comments on both (backend/AGENTS.md §2: every exported type gets one).

### `repository.go`

Add to the Repository interface:

```go
InsertRefreshToken(ctx context.Context, tx pgx.Tx, token *RefreshToken) error
InsertUserLog(ctx context.Context, tx pgx.Tx, log *UserLog) error
```

### `repository_db.go`

Implement both using goqu Insert, same pattern as `InsertAuthToken`
(repository_db.go:131 — pgx.Tx parameter, goqu INSERT, wrapped error via
`fmt.Errorf("...: %w", err)`). SQL is always parameterized via goqu — never
string concatenation (root AGENTS.md golden rule).

### `service.go`

- Add field `googleOAuth *googleoauth.Client`.
- Add field `authKeys *auth.Keys` — used by IssueTokens for ES256 signing.
  Consumed read-only via constructor injection; **`internal/platform/auth/`
  is Tier 0 fenced (root AGENTS.md §3) — nothing in that package may be
  modified**. Loading happens at startup in main.go (Task 04).
- Update `NewService(...)` to accept both new dependencies (6th and 7th
  parameters: googleOAuth, authKeys per techplan §10).
- **This task owns the constructor-signature ripple**: update existing test
  files in this package that construct the Service. Do not loosen or delete
  existing assertions while doing so (root AGENTS.md §4).

### `google_oauth.go` (new) — signatures

```go
func (s *Service) GoogleRedirect(ctx context.Context, intent string, sessionUserID *uuid.UUID) (redirectURL string, cookieValue string, err error)
func (s *Service) GoogleCallback(ctx context.Context, code string, state string, cookieValue string) (result CallbackResult, err error)
func (s *Service) IssueTokens(ctx context.Context, userID uuid.UUID) (accessToken string, refreshToken string, err error)

type CallbackResult struct {
    RedirectURL  string
    Error        string   // one of the ?error={code} codes, empty on success
    AccessToken  string
    RefreshToken string
}
```

### Business logic flow — verbatim contract from techplan §8

```
GoogleRedirect(ctx, intent, sessionUserID):
  validate intent in {login, link, reauth}         // else 400 — handler surfaces (R18)
  if intent in {link, reauth}:
    if sessionUserID is nil: return 401            // handler checks BEFORE calling (R2)
  state = randomString(32)
  nonce = randomString(32)
  cookie = encodeJSON({state, nonce, intent, user_id: sessionUserID})
  setCookie(HttpOnly, Secure, SameSite=Lax, MaxAge=600, Path=/auth/google)  // transport does the write (Task 04)
  url = googleAuthURL(client_id, redirect_uri, scope, state, nonce)
  return 302(url)

GoogleCallback(ctx, code, state, cookie):
  cookieData = readCookie()
  if state != cookieData.state (constant-time): return 302Error(state_mismatch)
  tokens = exchangeCode(code, redirect_uri)  // via googleOAuth.ExchangeCode — 10s timeout inside
    on timeout/error: return 302Error(google_unavailable)
  idToken, err = verifyIDToken(tokens.id_token, jwks, client_id, cookieData.nonce)
    on err == ErrNonceMismatch:   return 302Error(nonce_mismatch)
    on err != nil:                return 302Error(google_token_invalid)
  email = idToken.email
  emailHash = HMAC(email)
  switch cookieData.intent:
    case "login":
      googleIdentity = FindAuthIdentityByIdentifierHash("google", emailHash)
      if googleIdentity != nil:
        tokens = IssueTokens(googleIdentity.UserID)
        setAuthCookies(tokens)                     // transport (Task 04)
        return 302(appURL)
      epIdentity = FindAuthIdentityByIdentifierHash("email_password", emailHash)
      if epIdentity != nil:
        return 302Error(google_email_conflict)     // NO auto-merge — see below
      // new user
      user = create User + AuthIdentity(google, verified_at=now) in tx
      tokens = IssueTokens(user.ID)
      setAuthCookies(tokens)
      return 302(appURL)
    case "link":
      googleIdentity = FindAuthIdentityByIdentifierHash("google", emailHash)
      if googleIdentity != nil and googleIdentity.UserID != cookieData.user_id:
        return 302Error(google_link_conflict)
      // attach identity to existing user
      insertAuthIdentity(cookieData.user_id, "google", email, verified_at=now)
      writeUserLog(cookieData.user_id, "account_linking")
      return 302(securityPageURL)
    case "reauth":
      no AuthIdentity or token changes
      signal success so transport sets the reauth marker (Task 04 owns the store)
      return 302(securityPageURL)

IssueTokens(ctx, userID):
  accessToken = signJWT(ES256, {sub: userID, exp: now+15min}, keys.Private)
  refreshToken = randomString(32)
  refreshTokenHash = sha256(refreshToken)
  insertRefreshToken({id: uuid, user_id: userID, family_id: uuid,
    token_hash: refreshTokenHash, expires_at: now+30d})
  return accessToken, refreshToken
```

### The three security-critical business rules (techplan §1, §7 — human-review focus)

1. **No-auto-merge (R9, R10).** When a Google login returns an email already
   claimed by an email_password identity for a different user, the system
   must NOT automatically merge the accounts → `google_email_conflict`, no
   new records. This prevents account takeover via an unverified email claim
   (e.g. a Google Workspace admin provisioning an address without proving
   inbox control). Same discipline for link: no code path creates or attaches
   an AuthIdentity without an explicit, authenticated action from the account
   owner → `google_link_conflict` when the email is claimed by another user.
2. **verified_at state machine (R14).** New google AuthIdentities are created
   already verified: `verified_at = now()` at insert — they never pass
   through null (docs/spec/domains/account/invariants.md,
   auth_identities.verified_at section).
3. **Concurrent duplicate registration (R15, INV-account-01).** Two
   concurrent Google registrations for the same email: the pre-existing
   unique index on `auth_identities (provider_type, identifier_hash)`
   fails one insert cleanly — wrap the error, handle without crashing,
   no partial state.

Additional hard requirements for every line of this file:

- **R16 — no secrets/tokens in logs.** state, nonce, code, id_token,
  access_token, refresh_token values are NEVER logged anywhere in the OAuth
  flow — only the fact and outcome (AGENTS.md golden rule).
- **R23 — constant-time state comparison.** The callback's
  state-vs-cookie-state comparison uses subtle.ConstantTimeCompare (the
  nonce half lives in Task 02's VerifyIDToken).
- **PII pattern.** The Google email is PII: encrypted ciphertext + HMAC hash
  per the established entity.go pattern (`platform/crypto/` — already shipped
  in ticket 01, consumed read-only; still a fenced path — no modifications).
- **Error wrapping.** `fmt.Errorf("...: %w", err)` throughout; original
  errors never discarded.
- **New-user creation runs in a transaction** (User + AuthIdentity together —
  §8 pseudocode: "in tx").
- **IssueTokens details:** access token ES256, sub=userID, exp=now+15min;
  refresh token = random 32 bytes, only its sha256 hash stored
  (token_hash column, unique index from migration 000004); family_id =
  fresh uuid at first issuance; expires_at = now+30d. Rotation/reuse
  semantics are task #3's job — do not build them here.

## 5. Rules covered

Primary owner for: **R1** (state/nonce generation + redirect URL),
**R4** (constant-time state mismatch → state_mismatch, before any Google API
call), **R5** (ErrNonceMismatch mapping → nonce_mismatch result),
**R6** (exchange failure mapping → google_unavailable result), **R7**, **R8**,
**R9**, **R10**, **R11** (attach identity + user_logs entry
action_type=`account_linking`, Fitur 9's "account linking baru"), **R12**
(reauth: no identity/token changes; marker-setting itself is Task 04),
**R14**, **R15** (clean-error half), **R16**, **R19/R20** (missing/absent
state inputs detected as state_mismatch before any Google call — detection
shared with Task 04's param handling), **R23** (state half),
**R26** (generic verification failure → google_token_invalid result).

Full rule-to-task mapping: see `manifest.md`.

## 6. Testing checklist (this task's slice)

Unit tests with fakeRepo + fake Google client mirroring the
`*googleoauth.Client` surface. Table-driven shape
(`[]struct{ name string; ... }`) per backend/AGENTS.md §2.

Named tests required by the techplan:

- [ ] `TestGoogleCallback_NoAutoMerge_Login` — R9: login where the email
      belongs to an email_password identity of a *different* user → result
      carries Error=google_email_conflict, zero new User/AuthIdentity rows
      in fakeRepo.
- [ ] `TestGoogleCallback_NoAutoMerge_Link` — R10: link where the claimed
      Google email belongs to a different user → Error=google_link_conflict,
      no AuthIdentity attached.

Remaining coverage mapped 1:1 from §4 rules:

- [ ] R1: GoogleRedirect(login, nil) generates distinct non-empty state+nonce,
      encodes both plus intent into cookieValue, returns Google consent URL.
- [ ] R4: query-state ≠ cookie-state → Error=state_mismatch AND the fake
      Google client records **zero** calls (no API call before validation).
- [ ] R4/R20: missing/empty cookieValue → state_mismatch, zero calls.
- [ ] R19: missing code/state params → state_mismatch, zero calls.
- [ ] R5: fake client returns ErrNonceMismatch → Error=nonce_mismatch.
- [ ] R26: fake client returns generic verify error → Error=google_token_invalid
      (not nonce_mismatch).
- [ ] R6: fake client returns timeout/connection error → Error=google_unavailable.
- [ ] R7: login + existing google identity → CallbackResult carries tokens,
      RedirectURL=app URL, no inserts performed.
- [ ] R8: login + unknown email → User + AuthIdentity(provider_type=google,
      verified_at=now) inserted **in one tx**; tokens issued for the new user.
- [ ] R11: link + no conflict → AuthIdentity attached to the session user_id
      (not a new User); UserLog{ActionType:"account_linking"} written; 302
      target = security page.
- [ ] R12: reauth → no AuthIdentity/token writes; success result targeting
      security page.
- [ ] R14: assert verified_at non-null and ≈now on every created google
      AuthIdentity across all branches.
- [ ] R15: two concurrent IssueTokens/new-user flows racing on the same
      email — one hits the unique-violation path and returns cleanly.
- [ ] R23: state comparison path exercised (code-review assertion for
      subtle.ConstantTimeCompare usage).
- [ ] R16: capture log output across all branches; assert no state, nonce,
      code, id_token, access_token, or refresh_token substring appears.
- [ ] IssueTokens: access token parses as ES256 JWT with correct sub/exp;
      refresh_tokens row stored with sha256(token) not the raw token;
      expires_at ≈ now+30d.
- [ ] **Run the whole suite with `go test -race ./internal/domain/account/...`**
      — mandatory for anything touching this domain per backend/AGENTS.md §3
      (Tier 0/1 area; R15 concurrency case especially).

## 7. Common mistakes to avoid (techplan §13 slice)

| Mistake | Fix |
|---|---|
| Auto-merging on email match | Account takeover via unverified email claim — no code path creates/attaches identity without explicit authenticated action (R9/R10). |
| Creating google AuthIdentity with verified_at=nil | Violates the auth_identities.verified_at state machine (invariants.md) — set verified_at=now() at insert (R14). |
| Comparing state with `==` instead of subtle.ConstantTimeCompare | Timing side-channel on state value (R23). |
| Logging tokens/state/nonce/code/id_token | Log fact + outcome only (R16, AGENTS.md golden rule). |
| Storing the raw refresh token | Store sha256 hash; raw token exists only in the response/cookie path. |
| Creating User then AuthIdentity outside a transaction | Partial state on failure between the two inserts (§8: "in tx"). |
| Building refresh-token rotation/reuse detection "while we're here" | Task #3 owns rotation (INV-account-03/04) — scope discipline. |

## 8. Risk note (to fill in the PR)

- Assumptions made: ...
- Edge cases intentionally NOT handled (and why): ...
- Concurrency assumptions: ...
- What is not tested, and why: ...
