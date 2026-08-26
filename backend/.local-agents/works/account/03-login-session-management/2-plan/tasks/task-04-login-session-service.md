# Task 04: Login/session domain services

> Back-reference : `.local-agents/works/account/03-login-session-management/2-plan/techplan.md` (Status: Approved) — sections 3, 4 (R1–R5, R7–R11, R13–R16, R18), 8 (business flow — ordering is contractual), 10
> Depends on    : task-02 (repo methods), task-03 (token mint/verify funcs)
> Model         : GLM 5.2 (max) (branching flow / state-machine reasoning)
> Rules touched : R1–R5, R7–R11, R13–R16 (service halves), R18, R19
> ⚠️ TIER 0     : the reuse/race-loser branch inside `Refresh` is part of the Tier 0 fenced set — see manifest paired-pass checklist

## Objective

The four service methods + the fail-closed MFA verifier seam. Transport comes later (task-05); everything here must be unit-testable with fakes (house pattern: `breachChecker`, `googleOAuthClient`, `TxRunner`).

## Files

| File | Change |
|---|---|
| `backend/internal/domain/account/login.go` | New |
| `backend/internal/domain/account/mfa_verifier.go` | New |
| `backend/internal/domain/account/service.go` | Edit (constructor gains seams) |
| `backend/internal/domain/account/login_test.go` | New (table-driven) |

## Constructor additions (`service.go`)

```go
func NewService(repo Repository, db *pgxpool.Pool, bc *breachcheck.Client, sender notification.Sender,
    keys *crypto.Keys, googleOauth *googleoauth.Client, authKeys *auth.Keys, frontendURL string,
    mfa MfaVerifier, nowFn func() time.Time,
    mintAccess func(uuid.UUID, time.Time) (string, error),      // wraps task-03 MintAccessToken
    verifyPending func(string, time.Time) (uuid.UUID, error),    // wraps task-03 VerifyMFAPending
) *Service
```

Callers passing `nil` for `mfa` get the fail-closed stub. Update existing call sites (task-05 wires main.go; tests updated in place).

## MFA verifier seam (`mfa_verifier.go`)

```go
// MfaVerifier abstracts TOTP/backup-code verification, owned by account
// task #6 (mfa_totp_secrets/mfa_backup_codes logic). Until that task lands,
// stubMfaVerifier fails closed: no code ever verifies.
type MfaVerifier interface {
    VerifyTOTP(ctx context.Context, userID uuid.UUID, code string) (bool, error)
    // VerifyBackupCode marks used_at via guarded single-use UPDATE and
    // reports whether the code was valid AND unused AND MFA enabled.
    VerifyBackupCode(ctx context.Context, tx pgx.Tx, userID uuid.UUID, code string) (bool, error)
}
```

Stub returns `(false, nil)` for both. TODO marker referencing task #6 — a flag, not commented-out code.

## Sentinels + result shapes (`login.go`)

```go
var (
    ErrInvalidCredentials = errors.New("invalid credentials") // → 401 generic (task-05 maps)
    ErrLockedOut          = errors.New("locked out")          // → 429 same-generic-body
    ErrMfaPendingInvalid  = errors.New("mfa pending token invalid") // → 401
)

type LoginResult struct {
    Status          string // "ok" | "mfa_required"
    AccessToken     string            // "" when mfa_required
    RefreshTokenPlain string          // "" when mfa_required; plain exists ONLY to reach the cookie/body path
    AccessTokenExpiresAt time.Time
    User            *LoginUserView
    MFAPendingToken string            // set only when mfa_required
}
```

## Business flows (ordering contractual — techplan §8 verbatim semantics)

### `Login(ctx, email, password) (LoginResult, error)`

1. `identifierHash := crypto.HMAC([]byte(email), s.keys)`
2. **Lockout check FIRST**: `CountRecentFailedAttemptsByIdentifier(hash, "password", 15*time.Minute) >= 5` ⇒ return `ErrLockedOut`. **No attempt row written on this path** — credential never checked.
3. Fetch identity by `(providerEmailPassword, identifierHash)`.
4. **Always run bcrypt-shaped work** (R18): if identity found → `secrets.ComparePassword(*identity.CredentialSecret, password)`; if NOT found → burn a dummy compare against a package-level fixed bcrypt hash. Never early-return before comparable CPU work — wrong-email vs wrong-password must be wall-clock indistinguishable.
5. Mismatch ⇒ insert attempt `{stage:"password", success:false, user_id if known}` ⇒ `ErrInvalidCredentials`.
6. Match ⇒ insert attempt `{stage:"password", success:true, user_id}`.
7. Branch on `GetLoginUserView(...).MFAEnabled`:
   - **true**: `mintAccess`-analog pending mint (`MintMFAPending` via injected func pair) → `LoginResult{Status:"mfa_required", MFAPendingToken}` — **no session tokens, no refresh row**.
   - **false**: issue tokens exactly like `IssueTokens` (ES256 access 15 min via task-03 mint w/ `purpose:"access"`; refresh 30 d, fresh `family_id`, hash persisted) + load `LoginUserView` → `LoginResult{Status:"ok", …}`.
8. Unverified identities log in fine (R5) — no `verified_at` gate here.

### `LoginMfa(ctx, pendingToken, totpCode, backupCode string) (LoginResult, error)`

1. Exactly one of `totpCode`/`backupCode` non-empty, else treat as invalid input → `ErrMfaPendingInvalid`-class 401 at transport (decide: sentinel `ErrValidation` reuse or new — keep distinct from credential errors).
2. `verifyPending(pendingToken, now)` fail ⇒ **401-class error, no attempt row, no writes** (identity not reliably known).
3. Lockout by user: `CountRecentFailedAttemptsByUser(userID, "mfa", window) >= 5` ⇒ `ErrLockedOut`, **no attempt row**, before any code verification.
4. `mfa.VerifyTOTP` or `mfa.VerifyBackupCode` (backup path runs its guarded single-use `used_at` UPDATE inside a tx):
   - false ⇒ attempt `{user_id, "mfa", false}` ⇒ `ErrInvalidCredentials`.
   - true ⇒ attempt `{user_id, "mfa", true}` ⇒ issue tokens like Login-no-MFA ⇒ ok result.
5. Stub verifier makes step 4 always false until task #6 — success paths are fake-tested only.

### `Refresh(ctx, refreshTokenPlain string) (RefreshResult, error)`

1. `tokenHash := sha256Hex(plain)`; `FindRefreshTokenByHash`.
2. Not found, OR `RevokedAt != nil`, OR `ReplacedByID != nil`, OR expired ⇒ **reuse/expired branch**: `RevokeRefreshTokenFamily(familyID)` when the row exists ⇒ `ErrInvalidCredentials` (transport renders one indistinguishable 401). Race-loser ≡ attacker (spec Assumption D) — no disambiguation.
3. Else BEGIN tx → `RotateRefreshToken(tx, hash, child)`:
   - `rotated=false` ⇒ ROLLBACK ⇒ family revoke ⇒ same 401 (exactly-one-winner loser path).
   - `rotated=true` ⇒ child already inserted in-tx ⇒ COMMIT ⇒ return new plain + new access token (minted via task-03).
4. Result carries new access token (+expiry) and new refresh plain; transport sets the cookie.

### `Logout(ctx, refreshTokenPlain string) error`

Present ⇒ guarded revoke-by-hash; absent ⇒ nil either way. Always-idempotent; transport always clears cookie + 204.

## Cross-cutting rules

- **R18 timing**: implemented in Login step 4 (dummy compare against a fixed hash constant generated once via `secrets.HashPassword` at package init or committed literal with comment).
- **R19 log hygiene**: log fact + outcome + `user_id` only; never tokens/plain passwords/emails. Failure logs use sanitized categories like `notificationErrorCategory`.
- Errors wrapped `%w`; sentinels compared via `errors.Is` upstream.
- Concurrency note for `-race`: service holds no mutable state; all shared state lives in Postgres behind the guarded statements (document this in the Service doc-comment).

## Unit tests (`login_test.go`, table-driven, fake repo/verifier/clock)

R1 happy path (asserts attempt row args, TTLs 15m/30d, fresh family); R2 mfa-required (no tokens, no refresh row, pending present); R3 wrong-email body-class == wrong-password (sentinel equality + identical handling); R4 lockout incl. boundary (4 failures pass, 5th rejected; rejected call writes nothing; ordering proof: repo count called before compare — assert via fake call log); R5 unverified succeeds; R6/R17 covered in task-03 but re-asserted here via injected funcs rejecting cross-purpose; R7 MFA-stage lockout keyed user_id pre-code; R8 TOTP success via fake verifier (attempt mfa=true + issuance); R9 backup single-use semantics asserted through fake verifier contract; R10 invalid pending variants; R11 wrong code; R12 rotation happy (child same family, parent marked); R13 missing/expired; R14 reuse A→B→C replay-A (family revoked); R15 two sequential-race simulations where second Rotate loses ⇒ family revoked; R16 logout both branches; R18 hasher call-log structural assertion.

## Out of scope

HTTP concerns, cookie/error mapping, env wiring (task-05); DB-level race proof (task-06).
