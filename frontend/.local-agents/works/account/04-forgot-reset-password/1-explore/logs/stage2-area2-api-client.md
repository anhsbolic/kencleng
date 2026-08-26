# Stage 2 — Area 2: API client layer

## Current state

- **`lib/api/schema.d.ts` (generated, do not hand-edit) already has
  full types for both endpoints** — confirms the openapi-typescript
  generation step was already run against the updated
  `api/openapi/account.yaml`:
  - `"/auth/forgot-password"` (schema.d.ts:362-405): `post` takes
    `ForgotPasswordRequest` (`{ email: string }`), responses `202`
    (`GenericAcceptedMessage { message?: string }`) and `429`
    (`responses["TooManyRequests"]`). No `422`/`404` branch defined —
    matches the feature spec's always-`202` design exactly.
  - `"/auth/reset-password"` (schema.d.ts:406-471): `post` takes
    `ResetPasswordRequest` (`{ token: string; new_password: string }`),
    responses `200` (inline `{ message?: string }`), `404`
    (`application/problem+json` → `Problem`), `410` (same `Problem`
    shape), `422` (`responses["ValidationError"]` →
    `Problem & { errors?: { field, message }[] }`, i.e.
    `ValidationProblem`).
  - `components["responses"]["ValidationError"]` (schema.d.ts:4495) and
    `["TooManyRequests"]` (schema.d.ts:4504) are the same shared
    response components already used elsewhere (e.g. `register`'s
    `422`), not something newly introduced for this feature.
  - Cross-checked against `api/openapi/account.yaml:235-284` (the
    hand-authored source) — wording and status codes match the
    generated types exactly, no drift between the two.
- **`lib/api/account.ts` has zero forgot/reset-password code** —
  confirmed by the earlier grep and a full read of the file (200
  lines): it exports `getMe`, `register`, `verifyEmail`,
  `resendVerification`, `login`, `loginMfa`, `logout`, plus the shared
  `postAccountAction` helper and `ApiError`/`RegisterResult` types.
  Nothing for `forgot-password`/`reset-password`.
- **Two established precedent shapes already exist in this same file**,
  both directly reusable:
  1. `register()`'s pattern for a "distinguish one specific validation
     branch, throw everything else as `ApiError`" function: checks
     `res.status === 202` for the anti-enumeration generic-accept
     branch, `res.status === 422` for field errors (destructured into
     a local `ValidationErrorItem[]`, not the generated
     `ValidationProblem` type directly — an inline shape is hand-typed
     instead, an existing minor style choice, not a defect), then
     throws `ApiError` for anything else. **This is close to an exact
     template for `forgotPassword()`** (202 generic-only, 429 handled
     via the `ApiError` fallback).
  2. `verifyEmail()`'s pattern for "several distinguishable failure
     statuses the caller branches on via `.status`, no discriminated
     return type": resolves on `res.ok`, otherwise throws `ApiError`
     with `readProblemDetail`. **This is close to an exact template for
     `resetPassword()`**, since `404`/`410` need to be distinguishable
     by the caller (banner copy differs: "not found/used" vs
     "expired") the same way `verify-email-status.tsx` already branches
     on `verifyEmail`'s thrown `ApiError.status`.
- `postAccountAction` (account.ts:54-64) already normalizes network
  failure into `ApiError(0)` and is the shared POST helper every
  existing auth action uses — both new functions should go through it
  too, not call `apiFetch` directly.
- `lib/api/client.ts`: `apiFetch` always attaches `Authorization` if an
  access token happens to exist in `useAuthStore`, but forgot/reset are
  unauthenticated (`security: []` in the spec) — harmless either way,
  since the backend for these two routes doesn't require the header,
  and a logged-in user hitting "lupa password" from a stale tab
  attaching their real token isn't a documented concern. No special
  handling needed here (`apiFetch` doesn't have a way to *omit* the
  header conditionally, but nothing in the spec asks for that).

## Requirement

- Per `04-forgot-reset-password.md`: `forgotPassword` should always
  resolve on `202` (never a validation-branch distinction — there is no
  `422` for this endpoint), matching `GenericAcceptedMessage`.
  `resetPassword` must let the caller distinguish `410` (expired,
  no retry with same token) from `404` (not found/used, no retry) from
  `422` (fixable — validation failed, token stays valid so the same
  link can be resubmitted) — the spec's Assumption B explicitly makes
  this ordering/distinction a correctness requirement, not just UX.
- `frontend/AGENTS.md` §3: types must come from `lib/api/` (generated),
  never hand-rolled duplicates of what OpenAPI already defines.

## Gap

- `forgotPassword(input: ForgotPasswordRequest): Promise<{ message?: string }>`
  and `resetPassword(input: ResetPasswordRequest): Promise<{ message?: string }>`
  (or a slightly richer discriminated result for reset, to be decided
  in Stage 3) don't exist yet in `lib/api/account.ts` — need to be
  added following the `register`/`verifyEmail` precedents above. Purely
  additive; nothing in the existing file needs to change.
- No generated-type gap: unlike some other tasks in this domain, the
  schema already has everything needed — this is a "wire up the
  fetch function" gap, not a "regenerate/extend the schema" gap.

## Sniffing

- **Misleading signal**: none — schema.d.ts having full types could in
  theory look like "the API layer is already done," but `account.ts`
  itself (the actual consumed layer) has nothing, so this isn't
  actually misleading once both files are checked together; flagging
  only because a shallower check (schema.d.ts alone) could give a false
  "already wired" impression.
- **Risk**: `resetPassword`'s `422` shares the exact same
  `ValidationProblem`/`errors[]` shape as `register`'s `422` — low risk
  of getting the parsing wrong since there's a direct precedent to
  copy, but the *consequence* of misordering client-side validation
  vs. the actual network call is spec-flagged as a real bug class
  (Assumption B) — not a client-side ordering bug exactly (the backend
  enforces the real ordering), but the frontend must still not treat a
  `422` as if the token were consumed (e.g. must not redirect away from
  the reset-password page or clear the `token` on a `422`, since the
  same link needs to remain usable for retry).
- **Edge case**: a `429` on `/auth/reset-password` isn't documented in
  `schema.d.ts`'s response union for that endpoint (only `forgot-password`
  has `429` listed) — matches the feature spec, which only states a
  rate limit for `/auth/forgot-password` ("stricter `/auth/*` limit
  applies" is stated once, under the forgot-password heading only).
  Worth confirming in Stage 3 whether `reset-password` is meant to have
  no rate limit at all, or whether this is just an omission in the
  spec/schema — flagging as a **possible inconsistency**, not resolving
  it here.
- **Inconsistency**: the feature spec's prose says "Rate limit: stricter
  `/auth/*` limit applies" directly under the `forgot-password` acceptance
  criteria, phrased generally enough ("stricter `/auth/*` limit") that
  it could be read as applying to all `/auth/*` endpoints including
  `reset-password` — but `schema.d.ts` only encodes `429` on
  `forgot-password`, not `reset-password`. Whether this is deliberate
  (guessing a token can't be brute-forced usefully since it's random,
  single-use, and the account lockout for repeated failed *login*
  attempts is a separate mechanism) or an oversight isn't something to
  resolve in Stage 2 — noting it as a real doc/schema inconsistency for
  Stage 3 or a human to confirm.

Proceeding to Area 3 (hooks layer).
