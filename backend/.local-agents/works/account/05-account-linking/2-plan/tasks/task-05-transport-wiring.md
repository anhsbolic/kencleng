# Task 05: Transport layer — session middleware, handlers, wiring

> Back-reference : `.local-agents/works/account/05-account-linking/2-plan/techplan.md` (Status: Approved) — sections 4 (R15 + response shapes), 5 (D8, D9), 8 (API changes), 9 (steps 4–5), 10 (transport/main.go)
> Depends on    : task-03 and task-04 (handlers call `Service.SetPassword` / `Service.UnlinkGoogle`; sentinel values must exist)
> Model         : DeepSeek V4 Pro (contract-shape precision work: status codes, Problem Details bodies, exact error-type URIs — SWE-bench-style accuracy over creative reasoning)

## Objective

Expose the two service flows as the backend's first always-authenticated endpoints: shared session middleware over a new `/account/security/*` route group, both handlers following the established decode→boundary-validate→service→Problem mapping shape, sentinel→Problem mappings for the distinct 409s, and `main.go` wiring.

## Files

| File | Change |
|---|---|
| `backend/internal/transport/http/account_security.go` | New — `requireSession`, `ES256SessionVerifier`, `SetPasswordHandler`, `UnlinkGoogleHandler`, request DTOs |
| `backend/internal/transport/http/account_security_test.go` | New — R15 + contract-shape tests |
| `backend/internal/transport/http/errors.go` | + two 409 cases in `MapServiceError` |
| `backend/cmd/server/main.go` | + `accountMux` construction/mounting |

## Contracts to hit exactly

**Session enforcement (R15)** — `requireSession(verifier func(string) (uuid.UUID, error))`: token via `sessionToken(r)` (existing helper: `kencleng_access` HttpOnly cookie first, `Authorization: Bearer` fallback). Verifier body mirrors `GoogleTokenVerifier`'s options minus OAuth framing: `jwt.WithValidMethods([]string{"ES256"})`, `WithExpirationRequired()`, leeway `time.Minute`. Missing/expired/garbage/wrong-key → `401` Problem Details matching openapi's `responses/Unauthorized`; handler never executes. `platform/auth/` stays UNTOUCHED (Tier 0 fence).

**Handlers**:
- `POST /account/security/set-password` → decode conditionally-shaped body (`email?`, `current_password?`, `password`) → boundary validation mirroring `RegisterHandler` (password ≥8 defense-in-depth; never echo submitted values) → `svc.SetPassword` → Branch 1 (`nil` error) = **202** generic accepted message; Branch 2 success = **200** `{message}`; `ErrValidation`=422; `ErrInvalidCredentials`=401
- `POST /account/security/google/unlink` → body `{password}` → `svc.UnlinkGoogle` → success = **200** `UnlinkGoogleResponse{message}`; `ErrOnlyIdentity`=409 type `https://kencleng.dev/errors/only-identity`, detail verbatim: "Google adalah satu-satunya metode login Anda. Atur email dan password dulu sebelum melepas tautan."; `ErrRemainingUnverified`=409 type `https://kencleng.dev/errors/unverified-remaining-identity`, detail verbatim: "Kamu sudah atur email dan password, tapi belum diverifikasi. Verifikasi email kamu dulu sebelum bisa melepas tautan Google."; `ErrInvalidCredentials`=401

**Wiring**: `accountMux` with the two routes; mounted as `RateLimit(rps, burst)(requireSession(...)(accountMux))` on the root mux under `/account/`. Reuse the ECDSA public key already loaded for the OAuth verifier.

## Tests

- **R15** `TestRequireSession_MissingToken_401`, `TestRequireSession_ExpiredOrGarbageToken_401`, `TestRequireSession_BearerFallback_Accepted` (+ wrong-key rejection inherited from the copied verifier options)
- Handler-contract tests stubbing the service seam (interface slice like `googleOAuthService`): status codes per branch, byte-exact 409 problem types/details, generic-202 body equality across Branch-1 outcomes, no email echo in any body

## Common mistakes (techplan §13 subset)

- Echoing the submitted email in the 202 body or validation errors → enumeration aid; constants only
- Drifting from the spec'd problem-type URIs → copy verbatim from `api/openapi/account.yaml` (openapi-spec-first-drift)
- Putting session verification inside each handler instead of the middleware → D8 rejected that for good reason

## Out of scope here

CSRF second layer (deferred — techplan Resolved #7); rate-limiter keying change (Resolved #9); MFA/me routes (#6/#7).

## Verification

`go test -race ./internal/transport/http/...` green; `go run ./cmd/server` boots with routes registered (dev env smoke).
