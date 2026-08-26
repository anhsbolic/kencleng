# Task 2: Register Page

> Originating contract techplan: `../techplan.md` ("Tech Plan: Register &
> Email Verification (Frontend)", account/01-register-email-verification,
> Status: Draft). Cross-check high-level decisions there whenever this
> task file is ambiguous.
>
> Splitting axis: Dependency/sequence chain (primary) + Component/module
> boundary (secondary). **Depends on Task 1** (needs `useRegister()`,
> `useResendVerification()` via `ResendVerificationControl`, and the
> `register()`/hook contract to already exist). **Does not depend on
> Task 3** — the two page tasks are independent of each other and can
> run in parallel once Task 1 is merged. See `../manifest.md`.

## Scope

Replace the `/register` placeholder with the real form: fields,
validation, submit/loading/error/success states, the "Daftar dengan
Google" entry point, and the resend affordance on the success view
(reusing Task 1's `ResendVerificationControl`).

**Out of scope for this task**: the Google OAuth callback (Task #2 of
the domain's own numbering — a *different*, separate ticket,
`02-google-oauth-login-register.md` — not to be confused with this
decomposition's "Task 2"), `/verify-email` (this decomposition's
Task 3), any change to `lib/api/account.ts`/hooks/`mocks/handlers.ts`
(Task 1 owns those — import from them, don't redefine).

## Background (condensed from techplan §1)

`app/(auth)/register/page.tsx` is currently a 6-line placeholder whose
own comment states it exists only "so the Auth Shell has a route to
render against" until this task builds the real form. It sits inside
the existing `AuthShellClient` (desktop modal / mobile full page) —
unmodified by this task, reused as-is.

## What this task builds (from techplan §8/§9/§10)

- **File**: `app/(auth)/register/page.tsx` — replace the placeholder
  with a thin wrapper rendering `RegisterForm`.
- **File**: `components/features/account/register-form.tsx` (new) —
  Client Component (`'use client'`, leaf-level per
  `server-client-component-boundary.md` — nothing else on this page
  needs to be a Server Component, but the split stays explicit and
  named, matching how `AuthShellClient` already draws this boundary at
  the layout level). Fields, submit handling, success view, Google
  entry point.
- **File**: `components/features/account/register-schema.ts` (new) —
  `zod` schema for `name`/`email`/`password` (R1/R2 below).

**Imports from Task 1 (must already exist before this task starts):**
```typescript
import { useRegister } from "@/lib/hooks/use-register";
// useRegister() -> useMutation wrapping register(): Promise<RegisterResult>
// RegisterResult = { ok: true } | { ok: false; kind: "validation"; errors: { field: string; message: string }[] }

import { ResendVerificationControl } from "@/components/features/account/resend-verification-control";
// props: { email: string } (at minimum) — see Task 1 for its full contract
```

## Rules & Validation owned by this task

(Numbering matches techplan §4 — not renumbered per task.)

- **R1** (register form fields): Given the register page loads, When
  rendered, Then it shows `name`, `email`, `password` fields matching
  `RegisterRequest`'s shape exactly (`lib/api/schema.d.ts` — `name` is
  easy to miss since `page-map.md`'s one-line description doesn't name
  fields) — `zod`: `name` non-empty, `email` valid format, `password`
  ≥8 chars (feature spec's length policy,
  `docs/spec/1-account/features/01-register-email-verification.md`).
- **R2** (no client breach-check): Given a password ≥8 chars that's
  actually breach-listed, When submitted, Then the client schema does
  **not** reject it locally — that check is server-only; the client
  only ever sees it as a `422` round-trip (R5).
- **R3** (submit button state): Given the form is submitting, When
  `isPending` (from `useRegister()`), Then the submit `Button` has
  `type="submit"` (explicit — `Button`'s own default is `"button"`)
  and is disabled with its `loading` prop set.
- **R4** (register 202 → success): Given any `202` response
  (`useRegister()`'s result is `{ok:true}`), When received, Then the
  form is replaced by a fixed success view displaying the response's
  own `GenericAcceptedMessage.message` verbatim (e.g. *"Kalau email
  belum terdaftar, cek inbox untuk verifikasi. Kalau sudah, cek inbox
  untuk instruksi lebih lanjut."*, per `schema.d.ts`'s worked example on
  this endpoint) — never a client-authored variant, and never
  conditioned on anything other than "was this a 202."
- **R5** (register 422 → field errors): Given `useRegister()`'s result
  is `{ok:false, kind:"validation", errors}`, When received, Then each
  `{field, message}` is mapped via `form.setError(field, {message})` —
  the displayed text is the backend's `message` **verbatim**, never
  re-authored client-side; no banner is shown for this case
  (`form-validation-boundary.md`).
- **R6** (universal fallback, register-specific instance): Given
  `useRegister()` throws (network failure, unexpected `5xx`, or a `429`
  handled per R10 below), When it occurs, Then a `<Banner
  variant="error">` is rendered as the **first child inside the
  `AuthShellClient` panel** — this is a documented, load-bearing
  convention from `auth-shell-client.tsx`'s own doc comment, existing
  specifically to prevent the known prototype bug
  (`prototype-reference.md` issue #1: login's generic auth failure
  wrongly rendered as a field-level error). Exact fallback copy is
  Open Item #5 on the parent techplan — placeholder text is acceptable
  for now, do not invent final product copy.
- **R7** (Google entry point): Given the user clicks "Daftar dengan
  Google", When clicked, Then it performs a plain browser navigation
  (`<a href>` or equivalent — **not** `apiFetch`/XHR) to
  `/auth/google/redirect?intent=register` (Decision D3 below; the
  endpoint issues a `302`, which only a real navigation follows
  correctly).
- **R10** (429 handling, register-specific instance): Given a `429` on
  register, When received, Then the response's own `Problem.detail`
  text is shown verbatim via the same `<Banner variant="error">` slot
  as R6 (distinct from R6's fallback only in that this text comes from
  a documented backend string, not frontend-owned copy).
- **R17** (focus on register success): Given the register form is
  replaced by the success view (R4), When the transition happens, Then
  focus moves into the success region (`accessibility-fundamentals.md`
  — focus must move explicitly on async content replacing the form,
  never left on a now-removed element).
- **R18** (no enumeration-defeating client check): Given the register
  form, When the user types/blurs the email field, Then **no** request
  is made to check whether that email is already registered — the only
  network call this form ever makes is the final submit
  (`restapi/anti-enumeration.md`).

**R8/R9** (resend affordance on the success view) are Task 1's
`ResendVerificationControl` contract — this task's obligation is only
to mount it correctly inside the success view (R4/R17) with the
submitted email, held in local component state only, never persisted.

## Decision Log entries relevant to this task

**D3 — Google button scope split (Task #1 vs the domain's own Task #2)**

| Option | Why rejected/accepted |
|---|---|
| A. Defer the whole button to the domain's Task #2 (`02-google-oauth-login-register.md`) | Rejected — leaves `/register` incomplete against its own `page-map.md` row for no technical reason. |
| B. Build the button + navigation now; leave the OAuth callback to the domain's Task #2 (**chosen**) | `schema.d.ts`'s doc comment describes the redirect endpoint as a plain `302`-issuing navigation target, distinguished only by an `intent` query param — no OAuth client logic needed to make the button work today. |

**Note**: this split (D3) is flagged as **Active Open Item #3** on the
parent techplan — confirm against `tasks.md` with whoever owns the
domain's Task #2 before merging, since it's a cross-task scope boundary
inferred from a schema doc comment, not an explicit ratification.

## Backward Compatibility

Same as the parent techplan §6 — no DB, no API change, no existing
consumers affected. Nothing task-2-specific to add.

## Edge Cases & Risks relevant to this task

| Risk | Likelihood | Severity | Mitigation |
|---|---|---|---|
| Frontend copy/behavior differentiates an enumeration-sensitive branch (register's 4 backend branches all collapse to one `202`, but a UI author might be tempted to "helpfully" vary the message) | Medium — an easy slip while polishing UX copy later | **High** — defeats the backend's uniform-`202` design entirely, from the client side | R4: exactly one fixed success string, sourced from the response's own generic field, never keyed off anything else; code-review checklist item |
| A live "is this email already registered" check gets added later (on-blur/typeahead) | Low today, Medium over time as the form gets polished | **High** — same class of leak as above, arguably worse since it's an explicit new request | R18: explicit prohibition, call out in code review |
| `Button`'s `type="button"` default is left unset on the submit button | Medium — easy to miss, not HTML's own default | Low — caught quickly by any keyboard/Enter-key test, but invisible in a purely visual review | R3 + dedicated test |
| Google button implemented as a client-side fetch instead of a real navigation | Low | Medium — breaks the entire Google-register entry point (a `302` can't be "fetched" and followed the same way) | R7: explicit plain navigation, not `apiFetch` |
| Focus left on the removed form after success view replaces it | Medium — invisible in a visual-only review pass | Medium — screen-reader/keyboard users lose their place | R17 + dedicated a11y test |

## Files Changed / NOT Changed (this task's subset)

| File | Change Type | Description |
|---|---|---|
| `app/(auth)/register/page.tsx` | Modify | Replace placeholder with real page composition |
| `components/features/account/register-form.tsx` | Add | Register form client component |
| `components/features/account/register-schema.ts` | Add | zod schema for R1/R2 |
| Corresponding `*.test.tsx` for the two files above | Add | Per Testing Checklist below |

| File | Reason untouched (this task) |
|---|---|
| `lib/api/account.ts`, `lib/hooks/use-register.ts`, `mocks/handlers.ts`, `resend-verification-control.tsx` | Task 1's scope — import and consume, don't redefine |
| `app/verify-email/page.tsx` | Task 3's scope |
| `app/(auth)/login/page.tsx`, `forgot-password/page.tsx`, `reset-password/page.tsx` | Out of scope — domain Tasks #3/#4 |
| `app/(auth)/layout.tsx`, `_components/auth-shell-client.tsx` | Unmodified — `/register` continues using the existing shell as-is |

## Testing Checklist (this task's subset)

- [ ] R1: register form validates `name`/`email`/`password` required, email format, password ≥8 chars (zod) before submit
- [ ] R2: password field has no client-side breach-list check — only length is validated locally
- [ ] R3: submit button uses `type="submit"`, disables and shows a loading spinner while pending
- [ ] R4: a mocked `202` replaces the form with a success view showing the exact `GenericAcceptedMessage` text, unconditional on which internal branch the mock represents
- [ ] R5: a mocked `422` maps each error to the correct field via `setError`, verbatim backend message text, no banner shown for this case
- [ ] R6: a simulated network failure / unexpected `5xx` on register shows the frontend-owned generic banner inside the `AuthShellClient` panel's designated first-child slot, never the raw response body
- [ ] R7: "Daftar dengan Google" renders as a real link/navigation to `/auth/google/redirect?intent=register`, not an `apiFetch` call
- [ ] R10: a mocked `429` on register shows the response's `Problem.detail` text
- [ ] R17: focus moves into the success view once register's form is replaced
- [ ] R18: no request fires on email blur/change beyond the explicit submit (asserts no extra network call is made)

(R8/R9's own mechanism is verified in Task 1; this task only needs a
smoke-level assertion that `ResendVerificationControl` is mounted in
the success view with the correct `email` prop — see Task 1 for the
control's own behavior tests.)

## Testing Examples & Common Mistakes (this task's subset)

| Mistake | Error/Behavior | Fix |
|---|---|---|
| Submit button omits `type="submit"` | Enter key / implicit form submit does nothing; only visible via an actual keyboard test, invisible in a visual review | Always pass `type="submit"` explicitly — `Button` defaults to `type="button"` |
| Adding a live email-availability check on blur | Defeats the backend's uniform-`202` anti-enumeration design from the client side | Never add such a check (R18) — the only network call this form makes is the final submit |
| Hardcoding the `202` success copy into `RegisterForm` | Copy silently drifts from the backend's actual (and potentially later-changed) response text | Always render the response's own `message` field; only the mock fixtures in tests should contain literal example strings |

## Open Items relevant to this task

- **Active Open Item #3** (parent techplan §14): confirm the Google
  button/navigation scope split (D3) against `tasks.md` with whoever
  owns the domain's own Task #2, before merging.
- **Active Open Item #5** (parent techplan §14): exact copy for the
  frontend-owned fallback banner (R6) — placeholder text only until
  product sign-off.
