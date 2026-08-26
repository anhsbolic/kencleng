# Area 2 — API contract

> Stage 2 gap analysis. Files: `api/openapi/{account,common,index}.yaml`
> (hand-authored source) + `api/openapi.yaml` (generated bundle).
> Per `api/README.md`: never edit the bundle; edit split files, run
> `npm run bundle`, commit both.

## Current state

- **Structure**: `api/openapi.yaml` at root is generated (redocly bundle)
  from hand-authored split files under `api/openapi/`. The two copies were
  diffed for these endpoints — content-identical.
- **`POST /auth/forgot-password`** (account.yaml:235–260):
  `security: []`, body `ForgotPasswordRequest {email, format: email}`,
  responses **only** `202` → `GenericAcceptedMessage` (Indonesian generic
  text) and `429` → shared `TooManyRequests`. Description documents all
  three internal branches (registered / unregistered / Google-only)
  collapsing into identical 202.
- **`POST /auth/reset-password`** (account.yaml:262–299): body
  `ResetPasswordRequest {token, new_password (minLength: 8, format:
  password)}`; responses `200` message object, `422` → shared
  `ValidationError` (`ValidationProblem` with per-field `errors[]`),
  `410` Problem ("Token expired"), `404` Problem ("Token not found /
  already used"). Success description notes force-logout-everywhere.
  **No `429` documented.**
- **Shared shapes** (common.yaml): `Problem` = RFC 9457
  `{type, title, status, detail?, instance?}`; `ValidationError` wraps
  `ValidationProblem` with `errors[{field,message}]`; `TooManyRequests`
  example uses type URI `https://kencleng.dev/errors/too-many-requests`.
- **Fitur 01 sibling precedent**: `/auth/verify-email` documents `410`
  with concrete example (`type: .../errors/token-expired`) and its `404`
  description explicitly covers tokens revoked by a later resend
  (`revoked_at IS NOT NULL`); `/auth/register` documents a `422`
  alongside its generic 202; login endpoints deliberately override the
  shared 429 with a local generic-body variant (account.yaml:692–709) so
  429 doesn't leak lockout status.

## Requirement

Handler DTOs/status codes must match this exactly; oapi-codegen types
derive from it.

## Gap

Contract already covers the feature — but see sniffing items 2 and 3 for
documented-vs-behavior mismatches inherited from siblings.

## Sniffing findings

1. **Risk** — any behavioral deviation (extra field in 202 body, different
   problem `type` URIs) becomes a contract violation invisible until
   frontend consumes it; contract changes touch both `account.yaml` and
   the regenerated bundle.
2. **Edge cases** — forgot-password documents **no `422`**, yet body is
   `required: true` with `format: email`. Malformed-email behavior is
   unspecified (register *does* document 422). Undocumented-but-plausible
   path.
3. **Miscontext** — feature spec error table lists `429` for both
   endpoints; contract documents it **only on forgot-password**.
   reset-password can absolutely return 429 (middleware wraps all
   `/auth/*`) — contract under-documents.
4. **Misleading signals** — `minLength: 8` looks like full password
   policy; the fail-open HIBP breach check is invisible in the schema.
5. **Inconsistency** — reset-password missing 429 (above); plus
   verify-email's 404 explicitly covers revoked tokens while
   reset-password's 404 says only "not found / already used" (same
   `revoked_at` wording drift as Area 1).
