# Task 03: Tier 0 — JWT mint/verify primitives (`platform/auth/token.go`)

> Back-reference : `.local-agents/works/account/03-login-session-management/2-plan/techplan.md` (Status: Approved) — sections 8 (pending-token contract), 10 (Implementation Details), 5 (pending-token decisions); Resolved #2/#6/#13
> Depends on    : nothing code-wise (golang-jwt/jwt/v5 v5.3.1 already in go.mod) — but task-04 cannot start before this lands
> Model         : GLM 5.2 (max) for drafting — **the model never waives the human gate below**
> Rules touched : R6, R17
> ⚠️⚠️ TIER 0 — FENCED SUB-AREA (root AGENTS.md §3; tasks.md KPI "human-authored or human-rewritten")

## The non-negotiable gate

Per techplan Resolved #13: this file is agent-DRAFTED with heavy doc-comments and exhaustive tests, then goes through a **dedicated human paired rewrite/review pass BEFORE `make verify` sign-off and commit**. The build report must list this file under "Tier 0 files awaiting paired rewrite". If the pairing session moves these functions elsewhere, only the home changes — not the guarantees.

## Objective

Two token purposes, cryptographically separated, one module:

1. **Access token** — ES256, signed with the existing `auth.Keys.Private`, now carrying an explicit `purpose:"access"` claim. Verified with `auth.Keys.Public`.
2. **MFA-pending token** — HS256, signed with a NEW dedicated secret, claims `{sub, purpose:"mfa_pending", exp}`, 5-minute TTL, **never persisted anywhere**.

Why key separation is the whole point (spec Assumption A/B): a `purpose` claim check is application logic that can be buggy or omitted on some path; a separate key turns wrong-purpose acceptance into outright signature-verification failure — cryptographic guarantee, not logic guarantee. The claim stays as belt-and-suspenders on top.

## Proposed API (`backend/internal/platform/auth/token.go`)

```go
// MintAccessToken signs an ES256 access JWT for userID.
// Claims: {sub: userID, iat, exp(15m), purpose: "access"}.
func MintAccessToken(private *ecdsa.PrivateKey, userID uuid.UUID, now time.Time) (string, error)

// VerifyAccessToken enforces ES256 signature under public,
// exp required (1-min leeway, matching GoogleTokenVerifier convention),
// and purpose == "access". Returns parsed userID.
func VerifyAccessToken(public *ecdsa.PublicKey, token string, now time.Time) (uuid.UUID, error)

// MintMFAPending signs a short-lived step-up token.
// Claims: {sub: userID, iat, exp(now+5min), purpose: "mfa_pending"}.
func MintMFAPending(secret32 []byte, userID uuid.UUID, now time.Time) (string, error)

// VerifyMFAPending enforces HS256 under secret32, purpose == "mfa_pending",
// exp required. Expired/malformed/wrong-key/wrong-purpose all collapse to
// one generic failure (no disambiguation leaks).
func VerifyMFAPending(secret32 []byte, token string, now time.Time) (uuid.UUID, error)

// ValidateMFAPendingSecret parses a base64 std-encoded secret and requires
// exactly 32 decoded bytes — mirrors platform/crypto New() discipline so
// misconfiguration fails at startup, not at first mint.
func ValidateMFAPendingSecret(b64 string) ([]byte, error)
```

`now` injected as parameter (not `time.Now()` inline) so tests are deterministic; callers pass from the service clock seam.

## Hard requirements

- Verifiers MUST pin algorithms: ES256-only for access, HS256-only for pending (algorithm-confusion defense; house precedent `googleoauth/client.go` RS256-only, `auth_google.go` ECDSA-check).
- `VerifyAccessToken` rejects tokens whose `purpose` is missing or ≠ `"access"` — including legacy OAuth-era tokens that lack the claim entirely (accepted breaking edge: sandbox, no deployed clients; `GoogleTokenVerifier` deliberately stays lenient for link/reauth gating — do NOT change it).
- `VerifyMFAPending` must reject anything that verifies as ES256 (wrong key type = fail).
- No logging of token contents anywhere in this file (R19).
- Doc comments on every export (AGENTS.md §2); comments should explain WHY (key separation rationale above), not just what.

## Tests (unit, deterministic via injected now)

| Test | Proves |
|---|---|
| mint→verify round-trip for both purposes | Happy paths |
| `TestAuthMiddleware_RejectsWrongSigningKey` | Access verifier rejects an HS256 token signed with the MFA secret AND a pending-token signed by a different ECDSA key |
| `TestAuthMiddleware_RejectsNonAccessPurposeToken` | Access verifier rejects `purpose="mfa_pending"` and purpose-less tokens; pending verifier rejects `purpose="access"` |
| expired-token rejection (now +5min+ε) | Expiry enforced, both purposes |
| malformed/garbage inputs | Generic failure, no panic |
| `ValidateMFAPendingSecret`: wrong length / bad base64 / empty | Startup-time failure modes |
| cross-purpose round-trip matrix | mint-access → verify-pending fails; mint-pending → verify-access fails |

## Out of scope

Issuance orchestration (task-04); env var wiring (`MFA_PENDING_TOKEN_SECRET` loading happens in main.go, task-05 — but the validation function lives here).
