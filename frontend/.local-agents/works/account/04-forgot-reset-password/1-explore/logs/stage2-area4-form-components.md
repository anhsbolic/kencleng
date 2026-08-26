# Stage 2 — Area 4: Form components

## Current state

Read in full: `register-form.tsx` + `register-schema.ts`,
`login-form.tsx` + `login-schema.ts`, and `verify-email-status.tsx`
(a fourth, highly relevant precedent — token-in-URL, not originally in
the Stage 1 area list, but directly on-point for `/reset-password`'s
shape). Also confirmed `components/ui/banner.tsx` (4 variants:
`success`/`error`/`warning`/`info`, `role="alert"` for
error/warning, `role="status"` for success/info — stable, shared,
matches `design-guidelines.md`'s token list exactly).

Three distinct, directly-reusable structural precedents exist:

1. **`RegisterForm`'s "always-202, single generic success view"
   shape** — near-exact template for `/forgot-password`: `useForm` +
   `zod`, `onSubmit` calls `mutateAsync`, catches only to distinguish a
   documented `429` (`error.detail` shown verbatim) from everything
   else (`GENERIC_ERROR_MESSAGE` constant, explicitly marked `// TBD —
   Open Item #5, placeholder pending product copy` in both existing
   forms), then on success swaps the entire form out for a `<Banner
   variant="success">` + heading, moving focus into the success
   heading via a `ref` + `useEffect` (accessibility-fundamentals
   convention, explicitly commented "R17"/similar rule numbers).
   `RegisterForm`'s own doc comment states the anti-enumeration
   discipline explicitly: "must never differentiate [branches] in
   copy or UI" — same discipline `forgot-password`'s spec demands.

2. **`LoginForm`'s "request-level banner as first child, never
   field-level" shape** — the shell's own documented convention
   (Area 1) made concrete here: `requestError` state renders a
   `<Banner variant="error">` wrapped in a focatable `ref` div, ahead
   of the form, with a `useEffect` moving focus there on error. This
   is the structural fix for the Known Issue (`prototype-reference.md`)
   and is exactly the shape `/reset-password` needs for its `404`/`410`
   request-level failures (must NOT render as a field-level error on
   the new-password input).

3. **`VerifyEmailStatus`'s "token-from-searchParams + `ApiError.status`
   discriminated outcome" shape** — the closest available precedent
   for `/reset-password`'s overall page shape: reads `token` via
   `useSearchParams()`, and its `errorToOutcome()` helper maps
   `error.status === 410 → "expired"`, `=== 404 → "invalid"`,
   `=== 429 → "rate-limited"`, else generic. **Critically, it already
   contains an explicit resolution for the exact edge case flagged in
   Area 1's sniffing**: `!token ? { kind: "invalid" } // R11 — a
   missing token is treated identically to 404, no separate message`.
   This is a real, already-shipped precedent for "missing token ==
   invalid token, same message" in this same domain — directly
   answers (as precedent, not as a Stage-2 resolution) whether
   `/reset-password` should do the same for its own missing-token case.
   One structural difference: `VerifyEmailStatus` fires its mutation
   **automatically on mount** (the token alone completes the action);
   `/reset-password` cannot do this — the token plus a user-entered new
   password are both required, so the mutation only fires on form
   submit, not on mount. The token-read/outcome-mapping half of the
   pattern transfers directly; the "fire on mount" half does not.

- `register-schema.ts`/`login-schema.ts` both use the exact same
  `password: z.string().min(8, "Password minimal 8 karakter")` UX-only
  rule, both with an explicit comment that the breach-list check is
  server-only and deliberately not replicated client-side
  (`form-validation-boundary.md`). This is the exact rule
  `reset-password`'s new-password field needs too, per the feature
  spec's own length-policy line.
- Both existing forms use a plain `<Input error={...} {...register(...)} />`
  + `<Label>` composition, `Button` with a `loading` prop tied to
  `isSubmitting || mutation.isPending`, and a `noValidate` form
  attribute (client `zod` validation is the only validation surfaced,
  never native browser validation UI).

## Requirement

- `patterns.md` Pattern 3 (Form): idle → validating → submitting →
  submit-error (banner, request-level) vs field-error (inline) →
  success. Guest-facing forms get an **inline success state**, not
  toast+redirect.
- Feature spec: `forgot-password` never has a field-level error branch
  at all (no `422` in its response union — confirmed in Area 2) — only
  a generic success view and a generic-detail `429`/network error
  banner. `reset-password` has both a field-level branch (`422`,
  password policy/breach) and two request-level branches (`404`/`410`)
  that must render as a banner, never conflated.
- `AGENTS.md` §3: `react-hook-form` + `zod`, validation rules from the
  feature spec (≥8 chars — not the breach-list, per the existing
  precedent's own reasoning).

## Gap

- `ForgotPasswordForm` needs to be written; per Requirement + precedent
  1, this is close to a line-for-line structural mirror of
  `RegisterForm` minus the name field, minus the Google button/divider,
  minus `ResendVerificationControl` (a `/forgot-password` success state
  doesn't have an analogous "resend" concept in the spec — a second
  submission just issues another independent token, no dedup/resend UI
  implied by the spec, matching the spec's Assumption A). Needs its own
  `forgot-password-schema.ts` (email-only).
- `ResetPasswordForm` needs to be written; per Requirement + precedents
  2 and 3 combined: a `<Banner>`-first-child request-level error slot
  (precedent 2) for `404`/`410`, field-level `zod`+backend-422 mapping
  for the password input (precedent 1's `setError` loop over
  `result.errors`/exception-thrown validation errors — needs deciding
  in Stage 3 whether `resetPassword`'s `422` throws `ApiError` the
  `verifyEmail` way or returns a discriminated result the `register`
  way, since the *rendering* differs depending on that Area-2 decision),
  and a `useSearchParams()`-sourced `token` read (precedent 3), but
  **without** precedent 3's auto-fire-on-mount behavior — the mutation
  fires on the form's submit handler instead, once a new password is
  entered. Needs its own `reset-password-schema.ts` (new-password only,
  same `min(8, ...)` rule).
- Neither component exists yet anywhere in `components/features/account/`.

## Sniffing

- **Misleading signal**: none — no partial/stub component exists to be
  mistaken for done work.
- **Miscontext**: none — `patterns.md`'s Form pattern description
  matches what's actually needed once the three precedents above are
  combined; no mismatch between what the pattern doc assumes and what
  the sibling components actually do.
- **Risk**: combining precedents 2 and 3 for `/reset-password` is the
  single most error-prone part of this whole feature's frontend
  surface — it's the only page in this task needing *both* a
  field-level 422 error *and* a request-level 404/410 banner *and* a
  URL-token read, where every existing single-purpose precedent only
  needed one or two of those three at once. Worth deliberate care in
  Stage 3 rather than blindly grafting pieces together.
- **Edge case**: `VerifyEmailStatus`'s R11 precedent (missing token ==
  404/invalid, same message) is the established pattern in this exact
  domain for exactly this situation — strong signal (not yet a Stage-2
  resolution) that `/reset-password` should follow the same rule for
  consistency, especially since `patterns.md`'s Status/Tracking pattern
  (§6, for `/donation/[id]/status`) independently states the identical
  principle for a different guest-facing token-in-URL page. Two
  independent parts of the codebase/docs pointing the same direction is
  a good sign for Stage 3, not a contradiction.
- **Inconsistency**: none found in this area beyond the 429-on-reset
  question already flagged in Area 2 (which resurfaces here only in
  that `ResetPasswordForm` won't have a rate-limit banner case to
  render unless Stage 3 decides otherwise).

Proceeding to Area 5 (mocks/tests).
