# Gap Analysis — Area 1: Requirement anchors

> Sources: `api/openapi/index.yaml`, `api/openapi/account.yaml`,
> `api/openapi/common.yaml`, `docs/spec/1-account/invariants.md`,
> `docs/spec/1-account/threat-model.md` (component 2),
> `docs/spec/1-account/features/03-login-session-management.md`

## Current state (requirement side)

- **Auth convention** (`index.yaml:20-26`): ES256 Bearer JWT access token;
  refresh token travels *only* as `HttpOnly`+`Secure`+`SameSite=Strict`
  cookie, set/read implicitly by `/auth/login`, `/auth/refresh`,
  `/auth/logout`; deliberately not a formal securityScheme. Cookie **name**
  and Max-Age are unspecified anywhere.
- **Endpoints** (`account.yaml:113-234`):
  - `POST /auth/login`: `security: []`, body `LoginRequest{email,password}`;
    200 `oneOf LoginResponse | LoginMfaRequiredResponse`; 401 Problem with
    generic `"Email atau password salah."`; 429 `$ref TooManyRequests`.
  - `POST /auth/login/mfa`: body
    `LoginMfaRequest{mfa_pending_token required, totp_code?, backup_code? —
    "one of"}`; 200 `LoginResponse` (cookie set); 401; 429.
  - `POST /auth/refresh`: reads refresh token from cookie (not body);
    200 `RefreshResponse{access_token, access_token_expires_at}`; reuse ⇒
    family revoked ⇒ 401.
  - `POST /auth/logout`: only `204` documented (matches idempotency criterion).
- **Schemas** (`account.yaml:703-842`): `User{id,name,email,email_verified,
  roles,auth_providers,mfa_enabled,created_at}`, `LoginResponse{status:"ok",
  access_token, access_token_expires_at?, user}`,
  `LoginMfaRequiredResponse{status:"mfa_required", mfa_pending_token}`
  (**marked `# INFERRED`**, :808), `RefreshResponse`.
- **`common.yaml:105-115`**: `TooManyRequests` example detail =
  `"Terlalu banyak percobaan gagal. Coba lagi dalam 15 menit."`
- **Invariants**: INV-03 (rotation single-use; `replaced_by_id` NULL→set at
  most once; no two children share a parent), INV-04 (reuse ⇒ every row in
  `family_id` revoked, incl. already-rotated; verification A→B→C replay A),
  INV-06 (backup code single-use + valid only while `enabled_at IS NOT NULL`,
  implicit invalidation on disable), INV-07 (referenced — login branches on
  `enabled_at`). Refresh-token state machine: two independent flags; usable iff
  `revoked_at IS NULL AND replaced_by_id IS NULL` (= `ix_refresh_tokens_active`
  partial-index condition).
- **Threat model component 2** (`threat-model.md:38-50`): persistent lockout
  5/15 min keyed by `identifier_hash` + in-memory rate limit on `/auth/*`;
  rotate-on-use + reuse detection; generic error message; refresh-flood rate
  limit; strict ordering tokens after MFA. Residual risk: no per-IP dimension.
- **Feature-spec resolved assumptions**: A/B — `mfa_pending_token` = stateless
  HS256 JWT `{sub, purpose:mfa_pending, exp}`, 5-min TTL, separate secret env
  var `MFA_PENDING_TOKEN_SECRET`, `purpose` claim kept as defense-in-depth;
  C — `login_attempts.stage` (`password`/`mfa`) + full DDL (2 lookup indexes +
  BRIN), ERD follow-up flagged; D — multi-tab refresh race is frontend concern.

## Requirement

Four endpoints implemented to match contract + invariants verbatim; Tier 0
fenced sub-area = JWT signing/verification (both keys) + rotation/reuse logic;
spec-named tests (`TestLogin_Lockout_5Failed15Min`,
`TestRefresh_ReuseDetection_FamilyRevoked`,
`TestRefresh_ConcurrentRequests_ExactlyOneWins`, etc.) are mandatory.

## Gap

None on the requirement side itself (docs complete). Baseline for areas 2–5.

## Sniffing findings

- **Inconsistency #1 (real):** feature spec (:46-48, :120) mandates the 429
  body carry the **identical** generic detail to the 401 — but
  `common.yaml:111-115`'s `TooManyRequests` example detail is a different
  sentence. Contract vs feature-spec conflict; per root AGENTS.md §1 order,
  feature spec outranks openapi, but the conflict must be resolved explicitly
  by a human → became Stage 3 D5.
- **Inconsistency #2:** threat-model component 2 predates the 2026-08-05
  resolutions — no MFA-stage (`user_id`-keyed) lockout row, no
  `mfa_pending_token`/key-separation mention. Threat model never revised.
- **Inconsistency #3 (internal to spec):** "Risk tier & rationale" § still
  says both tokens signed "with the same ES256 key" (:156-158) — stale vs
  Assumption A's separate HS256 secret.
- **Misleading signal:** `LoginMfaRequiredResponse` looks fully settled but
  carries `# INFERRED` ("proposals, not settled decisions", index.yaml:35-37)
  — though Assumptions A/B treat it as resolved.
- **Miscontext guard:** cookie name and Max-Age unspecified in contract —
  implementation must choose (frontend impact nil since HttpOnly).
- **Cosmetic:** spec headers reference `docs/spec/account/…` paths; actual
  tree is `docs/spec/1-account/…` (dead links in headers only).
