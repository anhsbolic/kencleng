# Gap Analysis — Area 3: `internal/platform/auth/`, `ratelimit/`, `secrets/`

> Files: `platform/auth/{doc.go,keys.go}`, `platform/ratelimit/doc.go`,
> `platform/secrets/secrets.go`, `platform/crypto/keys.go`,
> `cmd/server/main.go` env wiring (cross-check), feature-02 techplan
> decision records

## Current state

- **`platform/auth/`** = exactly two files: `doc.go` and `keys.go` (ES256 PEM
  keypair loader, pair-consistency check). Fence stated verbatim
  (`doc.go:4-8`): *"The backend scaffold was authorized to create the ES256
  key-pair loader below and nothing else — JWT signing/verification and
  session logic belong to a human-paired session."*
- **Feature-02 precedent:** access-token verification implemented **inline in
  the handler** (`transport/http/auth_google.go:56-65`, ES256-only parse with
  injected public key), deliberately NOT in `platform/auth/`. The 02 techplan
  records the human-reviewed decision (2026-08-22): "Task #3 can later extract
  a shared helper if needed — possibly as a human-paired change."
- **`golang-jwt/jwt/v5 v5.3.1`** already a direct dependency (go.mod:7).
- **`platform/secrets/secrets.go`**: both `HashPassword` AND `ComparePassword`
  exist — bcrypt verify side ready.
- **`platform/crypto/keys.go`**: `Keys{EncryptionKey, HMACKey}` from
  32-byte-base64 env vars — the pattern a new `MFA_PENDING_TOKEN_SECRET`
  would follow.
- **`platform/ratelimit/`**: empty skeleton ("implementation arrives in a
  later task"). Yet main.go *requires* `AUTH_RATE_RPS`/`AUTH_RATE_BURST`.
  (Resolution found in Area 4: the limiter was implemented at transport layer;
  this stub is stale.)

## Requirement

HS256 mint+verify for `mfa_pending_token` under separate secret; `purpose`
claim on access tokens checked by middleware; ES256 verification for protected
endpoints; in-memory rate limit on `/auth/*`.

## Gap

1. No HS256 support anywhere in `platform/auth`.
2. No `MFA_PENDING_TOKEN_SECRET` env var (not in main.go's required list).
3. No shared token-verification primitive (by deliberate prior decision).
4. Rate-limit implementation missing from the package that claims it (landed
   in transport instead — see Area 4).

## Sniffing findings

- **Risk / process tension (the big one):** spec assigns JWT signing/
  verification (both keys) to the Tier 0 fenced sub-area and AGENTS.md §3
  forbids agent writes to `platform/auth/` — yet the task is expected to
  *implement* those primitives. Feature 02 routed around via inline handler
  verification; whether that extends to *minting* HS256 tokens + a reusable
  middleware verifier is a human decision → Stage 3 D2.
- **Misleading signal (corrected):** threat-model credits "in-memory rate
  limit" as an existing mitigation; it does exist — but in transport
  middleware, not this package. Initial Area 3 suspicion of a missing
  mitigation was itself the misleading signal; the stub's doc.go is what's
  wrong.
- **Misleading signal #2:** golang-jwt being present makes HS256 trivially
  reachable — library availability is not the gap; authorization (fencing)
  and wiring (secret loading) are.
- **Edge case:** `auth.Load` validates keypair consistency at startup; the
  HMAC secret needs equivalent startup-time validation (length/format) to
  fail fast rather than at first mint → Stage 3 D2/D6.
