# Stage 2 — Area 2: API layer (`lib/api/account.ts`, `lib/api/schema.d.ts`, `lib/api/client.ts`)

## Current state

`lib/api/account.ts` (333 lines, read in full) has **zero** functions or
exported types for any of the 3 MFA endpoints — confirmed via direct
read, not grep-inference. Every other account-domain endpoint (register,
verify-email, login/login-mfa, logout, forgot/reset-password,
set-password, unlink-google) already has a thin wrapper function around
`postAccountAction`/`apiFetch`, following one of two shapes:

- **"Resolve on success, throw `ApiError` otherwise"** (`verifyEmail`,
  `resendVerification`, `resetPassword`, `logout`) — used when there's no
  field-level validation branch to distinguish.
- **"Discriminated-union result, only request-level failures throw"**
  (`register`, `forgotPassword`, `setPassword`) — used when a `422`
  validation response needs to reach the caller as data, not an
  exception, so field-level and request-level failures can never be
  conflated (established rule, restated in nearly every doc comment in
  this file).

Generated types **already exist** in `schema.d.ts` (confirmed by direct
read of lines 711–842, 3897–3912) and are ready to consume as-is:

- `POST /account/security/mfa/enroll` — no request body; `200` only,
  typed response `MfaEnrollResponse: { otpauth_uri: string }`. **No
  `409` response is typed here at all**, despite the feature spec's
  explicit "already enabled → `409`" acceptance criterion.
- `POST /account/security/mfa/enroll/confirm` — request
  `MfaEnrollConfirmRequest: { totp_code: string }`; `200` typed response
  `MfaEnrollConfirmResponse: { backup_codes: string[] }`; `422` typed as
  `components["responses"]["ValidationError"]` (→ `ValidationProblem` =
  `Problem & { errors?: { field, message }[] }`).
- `POST /account/security/mfa/disable` — request
  `MfaDisableRequest: { password?: string }` (optional — Google-only
  path sends no body); `200` → `{ message?: string }` (example text
  "MFA berhasil dinonaktifkan."); `401` typed as
  `components["responses"]["Unauthorized"]` (generic "Access token
  tidak valid..." `Problem`, **not** a dedicated wrong-password message
  — differs from `login`/`setPassword`'s dedicated credential-error
  `Problem` example).

`client.ts`'s `apiFetch` auto-retries any `401` once via
`coordinatedRefresh()` before returning it to the caller — this already
applies uniformly to every existing endpoint (including `setPassword`'s
and `unlinkGoogle`'s 401s), so MFA disable's 401 (wrong password, or
missing/expired Google reauth marker) goes through the same retry
before the caller ever sees it. Not a new behavior this task introduces,
just inherited plumbing to be aware of.

**No QR code rendering library exists in `package.json`** (checked
`dependencies` directly — `lucide-react`, `zustand`, `@tanstack/react-
query`, `react-hook-form`, `zod`, `@hookform/resolvers`, `next`, `react`
only). The backend only returns a raw `otpauth://` URI string
(`MfaEnrollResponse.otpauth_uri`) — turning that into a scannable QR
image is entirely the frontend's job, and nothing in the codebase does
this today. A new dependency will need to be introduced.

## Requirement

Feature spec's 3 endpoints, each needing a wrapper function following
this file's established conventions (thin wrapper, typed request/
response from generated schema, `ApiError` for undocumented/request-
level failures).

## Gap

- 3 wrapper functions to write: `mfaEnroll()`, `mfaEnrollConfirm()`,
  `mfaDisable()` — none exist.
- `mfaEnrollConfirm`'s `422` should almost certainly follow the same
  discriminated-result shape as `register`/`setPassword` (validation
  data returned, not thrown) given the schema types it as
  `ValidationError` — but the feature spec is explicit this isn't a
  field-level validation in the usual sense ("treated the same as an
  invalid code — no distinguishing response needed"), so the exact
  shape (a `field: "totp_code"` error vs. a request-level throw) is a
  real open decision for Stage 3, not yet resolved by precedent.
- `mfaDisable`'s two branches (`email_password` → `password` in body;
  Google-only → no body, relies on server-side marker) need a request
  type that's honest about optionality — `MfaDisableRequest.password`
  is already `?:` optional in the generated schema, so no type gap
  there, just a call-site decision (Stage 3).
- No QR rendering dependency — needs adding (Stage 3 decision: which
  library, client-only rendering vs. an `<img>` from a data URI, etc.).

## Page-consolidation check

N/A for this area (no page/endpoint mismatch risk here — this is purely
"does the wrapper-function layer exist," and endpoints already match
the feature spec 1:1, confirmed above).

## Sniffing

- **Inconsistency**: the feature spec requires `409` when `enroll` is
  called while already enabled, but `schema.d.ts`'s typed `responses`
  for that endpoint list **only `200`** — no `409` entry at all (unlike
  `login`'s endpoint, which *does* type its `401`/`429` explicitly, so
  this isn't a "this codebase never types non-200 responses" pattern —
  it's specifically absent here). This doesn't block the frontend
  (`postAccountAction` returns a raw `Response`; the wrapper function
  can still check `res.status === 409` manually and the `ApiError`
  path still works), but it means **the frontend cannot rely on
  generated typing alone to know a `409` is possible** — this has to
  be hand-derived from the feature spec text, and is worth flagging
  back since it suggests `api/openapi.yaml` itself may be out of sync
  with the feature spec's acceptance criteria (a backend-track/spec
  concern, not something to silently patch from the frontend side).
- **Miscontext risk**: `MfaDisableRequest`'s doc comment and the `401`
  response are generic ("Missing, invalid, or expired credentials" /
  "Re-authentication failed or missing") — there is no dedicated,
  distinguishable `.detail` string documented for "wrong password" vs.
  "Google reauth marker missing/expired," unlike `login`'s and
  `setPassword`'s 401s which do have dedicated example `Problem` text.
  If the real backend later returns genuinely different `.detail`
  strings per branch, showing `.detail` verbatim (this file's existing
  convention for 401s) would still work — but if it doesn't, a single
  generic message has to cover both branches, which is a copy decision,
  not something resolvable from this schema alone. Flag for Stage 3
  rather than assume either way.
- **Risk**: introducing a new external dependency for QR generation is
  the first new package this domain's frontend track has needed —
  worth flagging simply because it's a first (no established vetting
  process observed in this repo's docs for "which QR library," unlike
  e.g. `lucide-react` being pre-selected in `design-guidelines.md`).
- **Consistency**: aside from the two points above, the 3 MFA endpoints
  fit the existing wrapper-function conventions cleanly — no structural
  surprises, no auth/header handling beyond what `postAccountAction`
  already provides.
</content>
