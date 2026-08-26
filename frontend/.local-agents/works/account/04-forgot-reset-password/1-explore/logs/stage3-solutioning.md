# Stage 3 — Solutioning

Feature: `docs/spec/1-account/features/04-forgot-reset-password.md`
(frontend surface). Builds directly on all six Stage 2 area docs in
this same directory — decisions below cite which precedent/finding
each one follows.

## Architecture decision — `/reset-password` moves out of the Auth Shell

**Decided (confirmed with Anhar 2026-08-26): `/reset-password` becomes
a top-level route, `app/reset-password/page.tsx`, not nested under
`app/(auth)/`.** Discovered while drafting this stage: `app/verify-
email/page.tsx`'s own doc comment already establishes this exact
reasoning for the same email-link-entry scenario — "this link is
opened from an email client, often with no prior in-app navigation, so
`AuthShellClient`'s desktop modal (which blurs a rendering of `/`
behind it) would misleadingly imply the visitor was mid-browse."
`/reset-password?token=...` is opened from the reset email the same
way. `/forgot-password` stays under `app/(auth)/` — it's reached by
clicking a link on `/login`, genuine in-app navigation, so the modal
framing is accurate there.

Impact: the `app/(auth)/reset-password/` stub is deleted, not moved-
in-place — the real page lives at `app/reset-password/page.tsx`,
mirroring `app/verify-email/page.tsx`'s structure (plain centered
container, no `AuthShellClient`, `<Suspense>` boundary around the
`useSearchParams()`-reading form component). No existing link needs
updating (Stage 2 Area 1 confirmed nothing in-app currently links to
`/reset-password`).

**Flag for Anhar, not resolved here**: this is now the *second*
top-level, shell-less auth-adjacent route justified by the same
unwritten rule (email-link entry point). Worth considering whether
`page-map.md`/`patterns.md` should name this as an explicit exception
pattern rather than leaving it discoverable only via two separate code
comments (`verify-email/page.tsx` and, after this task,
`reset-password/page.tsx`).

## File plan

### New files

| File | Purpose |
|---|---|
| `app/reset-password/page.tsx` | Top-level route (see decision above). Heading + `<Suspense>`-wrapped `ResetPasswordForm`. |
| `app/(auth)/forgot-password/page.tsx` | **Rebuilt in place** (not new path) — same `flex flex-col gap-6` + heading-block shape as `/login`/`/register`, renders `ForgotPasswordForm`. |
| `lib/hooks/use-forgot-password.ts` | Bare `useMutation({ mutationFn: forgotPassword })` — no cache/redirect side-effect (Recommendation #3: no cache exists for a guest, no auto-redirect on this domain's own precedent). |
| `lib/hooks/use-reset-password.ts` | Bare `useMutation({ mutationFn: resetPassword })` — same reasoning; the redirect-to-`/login` link is the *component's* job (a `<Link>`, like `VerifyEmailStatus`'s "Masuk sekarang"), not the hook's. |
| `components/features/account/forgot-password-schema.ts` | `z.object({ email: z.string().email(...) })` — email-only, same message copy as `register-schema.ts`/`login-schema.ts`'s email rule. |
| `components/features/account/forgot-password-form.tsx` | See Component design below. |
| `components/features/account/reset-password-schema.ts` | `z.object({ new_password: z.string().min(8, "Password minimal 8 karakter") })` — same length-only rule and same "breach-list is server-only, not replicated" comment as `register-schema.ts`. |
| `components/features/account/reset-password-form.tsx` | See Component design below. |
| `components/features/account/forgot-password-form.test.tsx` | Mirrors `register-form.test.tsx`'s shape. |
| `components/features/account/reset-password-form.test.tsx` | Mirrors `verify-email-status.test.tsx` (token mocking, `server.use` branch overrides) + `register-form.test.tsx` (422 field-mapping test). |

### Modified files

| File | Change |
|---|---|
| `lib/api/account.ts` | Add `ForgotPasswordRequest`/`ResetPasswordRequest` type aliases, `ResetPasswordResult` discriminated type, `forgotPassword()`, `resetPassword()`. Purely additive — no existing export changes. |
| `mocks/handlers.ts` | Add `mockForgotPasswordAccepted`/`mockResetPasswordOk` fixtures + two default happy-path handlers, following the existing per-endpoint comment convention. |

### Deleted files

| File | Reason |
|---|---|
| `app/(auth)/reset-password/page.tsx` | Superseded by `app/reset-password/page.tsx` (shell decision above) — not left behind as a redirect stub, since nothing links to the old path yet. |

No changes needed to `lib/api/schema.d.ts` (already fully generated,
Stage 2 Area 2), `app/(auth)/layout.tsx`, `app/(auth)/_components/
auth-shell-client.tsx`, or `login-form.tsx` (its `/forgot-password`
link already points at the right, unchanged path).

## API layer design (`lib/api/account.ts`)

```ts
export type ForgotPasswordRequest = components["schemas"]["ForgotPasswordRequest"];
export type ResetPasswordRequest = components["schemas"]["ResetPasswordRequest"];

/**
 * POST /auth/forgot-password — always 202 generic (no 422 branch
 * exists for this endpoint, per the feature spec's anti-enumeration
 * design). Simpler than register()'s shape: no discriminated result
 * needed since there's nothing to discriminate. Closest existing
 * precedent is actually resendVerification() (also always-202,
 * generic-only), not register().
 */
export async function forgotPassword(
  input: ForgotPasswordRequest
): Promise<{ message?: string }> {
  const res = await postAccountAction("/auth/forgot-password", input);
  if (res.status === 202) return res.json();
  throw new ApiError(res.status, await readProblemDetail(res));
}

/**
 * POST /auth/reset-password — hybrid of register()'s and
 * verifyEmail()'s shapes (Recommendation #4): 422 (field-level,
 * retryable with the same token per spec Assumption B) returns a
 * discriminated ResetPasswordResult, matching register()'s exact
 * pattern; 404/410 (request-level, terminal — no retry with this
 * token) throw ApiError for the caller to branch on `.status`,
 * matching verifyEmail()'s exact pattern.
 */
export type ResetPasswordResult =
  | { ok: true; message?: string }
  | { ok: false; kind: "validation"; errors: ValidationErrorItem[] };

export async function resetPassword(
  input: ResetPasswordRequest
): Promise<ResetPasswordResult> {
  const res = await postAccountAction("/auth/reset-password", input);

  if (res.ok) {
    const body: { message?: string } = await res.json();
    return { ok: true, message: body.message };
  }

  if (res.status === 422) {
    const body: { errors?: ValidationErrorItem[] } = await res.json();
    return { ok: false, kind: "validation", errors: body.errors ?? [] };
  }

  throw new ApiError(res.status, await readProblemDetail(res)); // 404, 410, 429 (if it ever occurs), network, 5xx
}
```

No special-casing for the 429-on-reset doc/schema mismatch
(Recommendation #1) — the generic `ApiError` throw already surfaces
whatever `.detail` the backend sends for any undocumented status, so
`ResetPasswordForm`'s generic-error branch covers it without any
extra code.

## Component design

### `ForgotPasswordForm`

Structural mirror of `RegisterForm` minus the name field, minus
`GoogleAuthButton`/divider, minus `ResendVerificationControl` (no
resend affordance — a second forgot-password submission is already
safe and independent per spec Assumption A, so there's nothing to
"resend," the user can just submit the form again).

- Idle: email field only.
- Submit → `forgotPasswordMutation.mutateAsync`.
- Catch: `error.status === 429 && error.detail` → show verbatim; else
  `GENERIC_ERROR_MESSAGE` (same placeholder-pending-copy constant
  convention as `RegisterForm`/`LoginForm`).
- Success: swap to inline view — heading (focus-moved, same
  `useRef`+`useEffect` convention as `RegisterForm`'s R17) + `<Banner
  variant="success">{message}</Banner>` + a `<Link href="/login">`
  ("Kembali ke halaman login") since there's no resend control to
  anchor the view around this time.

### `ResetPasswordForm`

Combines `LoginForm`'s banner-first-child convention,
`VerifyEmailStatus`'s token-read + status-discriminated outcome
convention, and `RegisterForm`'s 422-field-mapping convention — the
piece Stage 2 flagged as highest-risk, so laid out in full:

```
token = useSearchParams().get("token")

if (!token) → render <Banner variant="error">{INVALID_LINK_MESSAGE}</Banner>
              + <Link href="/forgot-password">Minta link baru</Link>,
              no form rendered at all (Recommendation #2 — missing
              token treated identically to a 404, and there's no
              reason to let someone fill out a password against a
              link that can never succeed)

on submit(values):
  try:
    result = resetPasswordMutation.mutateAsync({ token, new_password: values.new_password })
  catch (ApiError):
    status 410 → terminal state: <Banner variant="error">{EXPIRED_MESSAGE}</Banner>
                 + <Link href="/forgot-password">Minta link baru</Link>,
                 form removed (no retry — matches spec's "no state
                 change" + the token being genuinely dead)
    status 404 → terminal state: same shape, INVALID_LINK_MESSAGE
                 (same copy as the missing-token case — one message
                 for both "malformed" and "doesn't exist/already
                 used", per Recommendation #2)
    else       → requestError = error.detail ?? GENERIC_ERROR_MESSAGE,
                 form STAYS rendered (network/5xx/undocumented-429
                 fallback — must remain retryable, this is not a
                 terminal failure)
  else (resolved):
    result.ok === true  → success state: <Banner variant="success">
                           + <Link href="/login">Masuk sekarang</Link>
                           (mirrors VerifyEmailStatus's own "verified"
                           outcome exactly — Recommendation #3, no
                           auto-redirect)
    result.ok === false → setError("new_password", { message }) for
                           each returned field error, form STAYS
                           rendered, requestError NOT set (422 is
                           field-level, never a banner — same
                           discipline as RegisterForm's R5)
```

Four mutually-exclusive top-level render states:
`no-token-or-terminal-error` / `success` / `form` (idle, submitting,
field-error, or request-error-banner-with-form-still-visible — these
four sub-states all share the same "the form is present" branch,
distinguished only by which pieces render inside it, exactly like
`RegisterForm`'s single form branch already handles idle vs
field-error vs request-error-banner today).

## Assumptions / open questions (per `docs/spec/README.md` §6.4 —
recorded, not silently resolved)

1. **Backend's `ValidationProblem.errors[].field` value for
   `reset-password`'s 422 is assumed to be the literal string
   `"new_password"`**, matching the request body's own key name — by
   direct analogy with `register`'s `field: "password"` (which matches
   `RegisterRequest.password`). Neither the feature spec nor
   `schema.d.ts` states this explicitly (the `field` type is a bare
   `string`, no enum). If the backend actually emits a different field
   name (e.g. `"password"` again, out of copy-paste from the same
   validation helper), `setError("new_password", ...)` would silently
   fail to attach the message to the visible input. **Needs
   confirmation against the actual backend implementation before
   merge** — a quick grep of `backend/internal/domain/account/service.go`'s
   reset-password validation-error construction would settle this
   directly; flagging rather than guessing further.
2. **`reset-password`'s terminal-error views (404/410) show a "Minta
   link baru" link back to `/forgot-password`.** Not specified by the
   feature spec (which only covers API status codes, not frontend
   copy/navigation) — a reasonable UX default consistent with
   `patterns.md`'s general "tell the user what happens next" success/
   error convention, but a product-copy decision, not a spec
   requirement. Copy itself (`INVALID_LINK_MESSAGE`, `EXPIRED_MESSAGE`,
   the link label) is placeholder pending product sign-off, same
   `// TBD` treatment as every other hand-written user-facing string
   in this domain (`GENERIC_ERROR_MESSAGE`, `INVALID_LINK_MESSAGE` in
   `verify-email-status.tsx`).
3. **No frontend-owned 429 handling is added for `reset-password`**
   (Recommendation #1) — the existing generic `ApiError` fallback
   already surfaces any backend `detail` text for an undocumented
   status, so no dedicated branch/test is added. If the doc/schema
   disagreement (spec's general "stricter `/auth/*` limit" line vs.
   `schema.d.ts` only encoding `429` on `forgot-password`) gets
   resolved by adding a documented `429` to the backend/OpenAPI later,
   this frontend code needs no change — the fallback already handles
   it correctly today.
4. **No dedicated frontend test for the concurrent-double-submit
   acceptance criterion** (Recommendation #5) — it's a backend
   invariant test per the spec's own Threat breakdown table
   (`TestResetPassword_TokenSingleUse_Concurrent`); the frontend's
   `404`-branch test already covers its half of the contract (render
   whatever single response comes back).

## Risk note (full, per `AGENTS.md` §8 exception)

- **Highest-risk single piece**: `ResetPasswordForm`'s four-branch
  outcome logic (no-token/terminal-error/success/form) is the only
  new code in this task combining three previously-separate error-
  handling disciplines (banner-vs-field separation, URL-token
  anti-enumeration, discriminated-422-with-retry). A mistake most
  likely to happen here: accidentally clearing/hiding the form on a
  `422` (would violate spec Assumption B — the user must be able to
  retry with the same link) — mitigated by keeping the 422 branch
  entirely inside the same render path as idle/submitting, never
  routed through the terminal-error state.
- **Second risk**: Assumption #1 above (the `new_password` field-name
  guess) is a silent-failure risk, not a crash risk — if wrong, the
  backend's validation message simply never appears next to the
  input, degrading to a confusing "nothing happened" UX on a `422`
  rather than an error. Must be confirmed against the real backend
  behavior (or a passing integration/contract test) before this is
  considered done, not just before merge review.
- **Architecture-change risk**: moving `/reset-password` to a
  top-level route is a net-new deviation from the Phase 0 scaffold's
  original placement. Confirmed no in-app links break (Stage 2 Area 1)
  and the precedent (`/verify-email`) already proves the pattern
  works in this exact codebase, so this is assessed as low-risk, but
  it is a real, deliberate departure from what was previously
  scaffolded — called out explicitly rather than silently changed.
- **Residual/accepted risk**: the two doc/schema inconsistencies noted
  in Stage 2 (429-on-reset, and now the shell-assignment gap in
  `page-map.md`) are both flagged for Anhar rather than resolved
  unilaterally in code or in spec docs — per `docs/spec/README.md`
  §6.3, an implementing agent must not edit the spec to make its own
  code pass; these stay open until a human resolves them.

## Testing plan

| Test file | Cases (rule-tag style, matching existing convention) |
|---|---|
| `forgot-password-form.test.tsx` | Client-side email-format validation; happy path (202 → success view + `/login` link); 429 shows backend detail verbatim; network failure shows generic message; never shows a field-level error for anything (no 422 branch exists to test). |
| `reset-password-form.test.tsx` | Missing token → invalid banner, no form rendered; happy path (200 → success view + `/login` link); 410 → expired banner + "Minta link baru" link, form removed; 404 → same generic invalid banner as missing-token case, form removed; 422 → field error under the password input via `setError`, form stays visible, no banner shown; network failure → generic banner, form stays visible. |

No dedicated hook tests (`use-forgot-password`/`use-reset-password`
are bare wrappers with no branching logic of their own — matches the
existing precedent that only `useLogin`/`useLoginMfa`, which contain
real logic via `applyLoginSuccess`, get their own hook test file). No
page-level test files for either route — matches the existing
precedent that pure single-form composition pages (`/register`,
`/verify-email`) have no page-level test beyond their form/status
component's own coverage.
