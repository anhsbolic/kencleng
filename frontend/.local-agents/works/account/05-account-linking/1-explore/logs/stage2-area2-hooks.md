# Stage 2 — Area 2: Hooks layer

## Current state

- No `use-set-password.ts` / `use-unlink-google.ts` exist yet (confirmed
  by directory listing and grep). Existing account hooks, all read in
  full:
  - `useForgotPassword`/`useResetPassword` — bare `useMutation({
    mutationFn })`, no cache invalidation, no `onSuccess` side-effect at
    all — both endpoints are pre-auth, nothing cached yet, form owns
    what to render.
  - `useLogin`/`useLoginMfa` — share `applyLoginSuccess` (`use-login.ts`):
    on success, writes the access token to `useAuthStore`, direct-writes
    `accountKeys.me()` via `queryClient.setQueryData` (no extra
    round-trip, since `LoginResponse.user` is structurally identical to
    `GET /account/me`), then `router.push`.
  - `useLogout` — `onSettled` (unconditional): clears access token,
    `queryClient.clear()` (full cache reset, not just `account.me`),
    broadcasts `logged-out` on the cross-tab auth channel. Does **not**
    call `router.push` itself.
  - `useVerifyEmail`/`useResendVerification` — bare mutations, no
    invalidation.
- **`SessionGuardProvider`** (`components/providers/session-guard-
  provider.tsx`) is the **single** place in the app that redirects to
  `/login` — it subscribes to `useAuthStore` and redirects whenever
  `accessToken` transitions from a real value to `null`, "regardless of
  what caused the transition" (own doc comment, citing techplan
  account/03 D4: one subscription instead of three separate redirect
  call sites). `useLogout` relies on this — it only clears the token,
  it never redirects directly.
- **`VerifyEmailStatus`** (`components/features/account/verify-email-
  status.tsx`), the existing `/verify-email` page component that the
  spec says Branch 1's step 2 "reuses... unchanged": on success it
  renders a static "Masuk sekarang" (sign in now) link and does **not**
  invalidate `accountKeys.me()` or touch `useAuthStore` at all — its
  only audience so far has been a logged-out visitor completing
  registration.
- `accountKeys.me()` (`use-account-me.ts`) is the shared query-key
  factory already used by `useLogin`'s direct cache write and (per its
  own doc comment) intended for exactly this kind of shared consumer.

## Requirement

- Branch 1 success (`202`): per spec, no re-authentication, no session
  change — just triggers an email. No hook-level session/cache action
  implied by the spec itself beyond showing the generic message.
- Branch 2 success (`200`): **all** of the user's existing refresh
  tokens are revoked in the same transaction (INV-account-05, confirmed
  identical wording/scope to `04-forgot-reset-password.md`'s
  already-implemented invariant — "every row in `refresh_tokens` for
  that `user_id`... must have `revoked_at IS NOT NULL`," no carve-out
  for the token belonging to the request that triggered it). This means
  the **current** session's refresh token is revoked too, even though
  the access token that just made this call stays valid for its
  remaining ~15-minute window.
- Unlink success (`200`): per spec, no explicit session-revocation
  requirement (only Branch 2 has that) — this action just removes the
  Google identity from the account.
- Step 2 of the 3-step flow (`POST /auth/verify-email`) must, in
  context, leave the frontend able to immediately attempt step 3
  (unlink) — which depends on `auth_providers`/`email_verified` being
  fresh, since Area 3's page will gate the "Lepas Tautan Google" action
  visibility/copy on that data.

## Gap

- `useSetPassword()`/`useUnlinkGoogle()` need to be created — no
  existing hook is a literal template for either, though `useLogin`'s
  `onSuccess`-does-real-work shape and `useLogout`'s `onSettled`-clears-
  token shape are the two closest precedents, one each.
- **Branch 2's post-success session handling has no existing hook to
  copy wholesale.** `useLogout` clears the token and relies on
  `SessionGuardProvider` for the redirect — the same mechanism would
  apply here (clearing the token after Branch 2 success would make
  `SessionGuardProvider` redirect to `/login` for free, consistent with
  the "one subscription, not three call sites" design already
  established) — but no hook today calls `clearAccessToken()` from
  anywhere other than `useLogout` itself, so this would be a second
  call site for the same store action, just triggered by a different
  mutation. Flagging as a possible fix, not deciding here.
- **`VerifyEmailStatus` reuse gap**: the spec's "reusing `POST
  /auth/verify-email` unchanged" claim is true at the **endpoint**
  level, but the existing **page component** built around that endpoint
  (`VerifyEmailStatus`) was built for a logged-out registration visitor
  and (a) shows a "Masuk sekarang" link that doesn't apply to an
  already-authenticated user completing Branch 1's step 2, and (b)
  never invalidates `accountKeys.me()`, so even if the user navigates
  back to `/dashboard/security` after verifying, the page's data would
  stay stale (identity still shown as unverified) until an unrelated
  refetch happens. This is a real gap the spec's "unchanged" framing
  doesn't surface — worth deciding in Stage 3 whether `VerifyEmailStatus`
  needs a small authenticated-aware branch, or whether a session check +
  cache invalidation belongs in `useVerifyEmail` itself (shared by both
  callers) instead of duplicating a second verify-email page.

## Page-consolidation check

- N/A directly (hooks aren't a page) — no page-map action or
  `tasks.md` endpoint is orphaned by anything found in this area.
  Confirms Area 1's finding: this task's two endpoints are the only
  ones this hooks layer needs to add.

## Sniffing

- **Risk**: if Branch 2's `onSuccess` does *not* clear the current
  access token/redirect, the acting user keeps a fully working session
  for up to ~15 minutes after changing their own password — during
  which the spec's stated threat ("a hijacked-but-still-valid session
  changes the password") is only half-closed: the attacker's *access*
  token still works for that window even though their refresh is dead.
  This is bounded (matches the documented 15-min access-token lifetime,
  not a new exposure), but the UX choice of whether to force an
  immediate re-login (matching `useLogout`'s pattern) vs. leave the
  current tab alone until natural expiry is a real decision for Stage 3
  with a security-adjacent angle, not purely cosmetic.
- **Miscontext**: the feature spec frames step 2 of the 3-step flow as
  "same endpoint, same mechanics" and treats it as effectively free —
  but from the frontend's side, the *component* wired to that endpoint
  was authored under a different-persona assumption (logged-out
  registrant) than this feature's actual step-2 caller
  (already-authenticated user mid-linking-flow). The spec author may
  not have looked at `VerifyEmailStatus`'s actual implementation when
  writing "reusing `POST /auth/verify-email` unchanged" — that sentence
  is accurate for the wire contract, not necessarily for the page.
- **Inconsistency**: none found beyond the one above — hook-layer
  conventions (bare mutation vs. side-effecting `onSuccess`/`onSettled`)
  are applied consistently across all five existing account hooks, no
  contradictions between them.

Proceeding to Area 3 (Dashboard Shell + `/dashboard/security` page +
page-consolidation check).
