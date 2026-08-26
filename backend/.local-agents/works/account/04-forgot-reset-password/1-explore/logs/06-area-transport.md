# Area 6 — Transport plumbing

> Stage 2 gap analysis. Files: `cmd/server/main.go` (routing),
> `internal/transport/http/errors.go`, `middleware.go`,
> `auth_verify_email.go`.

## Current state

- **Routing** (main.go:152–172): dedicated `authMux` wrapped in
  `transporthttp.RateLimit(rps, burst)` at mount — any route added to
  authMux inherits the stricter `/auth/*` rate limit automatically. Two
  `HandleFunc` lines are the entire wiring.
- **Rate limiter**: per-IP token bucket keyed on `r.RemoteAddr`
  (middleware.go:20), 429 Problem written directly; proxy-IP-collapse
  residual risk documented in threat-model §2.
- **`MapServiceError`** (errors.go:76): covers
  422/410/404/429(locked-out)/401/500. Unknown errors → generic 500, raw
  error logged server-side only.
- **Handler precedents**: `VerifyEmailHandler` rejects empty token pre-DB
  → 404 (no timing distinction); `ResendVerificationHandler` demonstrates
  always-202-generic including **swallowing internal errors into the same
  202** with sanitized server log (enumeration-leak avoidance) — while
  ResetPassword needs normal error propagation (422/410/404). Both shapes
  have in-repo precedent.

## Requirement

Two new handlers + DTO decode + validation + routing lines.

## Gap

Handlers don't exist; everything they need does.

## Sniffing findings

1. **Miscontext / existing divergence precedent** — resend returns **422**
   on malformed email (handler) but documents no 422 response in openapi;
   same pattern repeats on forgot-password (contract lists only 202+429).
   This slice inherits the decision rather than inventing it → Stage 3 Q3.
2. **Inconsistency** — problem `type` URI prefixes differ: code uses
   `https://kencleng.dev/problems/*` (errors.go), openapi examples use
   `https://kencleng.dev/errors/*`. Two vocabularies coexist across code
   and spec examples.
3. **Edge case** — empty-token early-reject status for reset-password:
   404 (verify-email precedent) vs 422? Contract silent; precedent says
   404.
