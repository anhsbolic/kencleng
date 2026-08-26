# Stage 2 — Area 3: Hooks layer

## Current state

- Confirmed via grep: no `use-forgot-password.ts`/`use-reset-password.ts`
  file and no `useForgotPassword`/`useResetPassword` symbol exists
  anywhere in `frontend/` today.
- Existing `lib/hooks/` conventions for this domain, all read in full:
  - `useRegister` (`use-register.ts`) — plainest form: `useMutation({
    mutationFn: register })`, nothing else. Justified inline: "nothing
    is cached anywhere for an unauthenticated user at this point ...
    so there's nothing this mutation needs to invalidate
    (`data-fetching-conventions.md`, confirmed explicitly rather than
    silently skipped)."
  - `useVerifyEmail` / `useResendVerification` — identical minimal
    shape, same one-line "no cache to invalidate" justification.
  - `useLogin` / `useLoginMfa` — the only two hooks doing more than a
    bare `useMutation`: both call a **shared** `applyLoginSuccess`
    helper (exported from `use-login.ts`) in `onSuccess`, which writes
    the access token to `useAuthStore`, direct-writes `accountKeys.me()`
    into the TanStack Query cache (no invalidation round-trip, since
    `LoginResponse.user` is already the same shape `GET /account/me`
    would return), and calls `router.push("/dashboard/profile")`. The
    file's own comment explicitly frames this sharing as deliberate:
    "Exported ... so `useLoginMfa` reuses this exact implementation
    instead of a second, drift-prone copy."
- Every hook in this domain is a **thin wrapper only** — no hook
  contains business logic beyond cache/redirect side-effects; the
  actual request/response handling lives entirely in `lib/api/account.ts`
  (Area 2). This is a consistent, load-bearing convention across the
  whole hooks directory, not just an account-domain quirk.

## Requirement

- `frontend/AGENTS.md` §3: TanStack Query for anything from the API,
  no server-state duplication into Zustand. `data-fetching-conventions.md`
  skill governs query-key/invalidation choices (not yet consulted in
  depth — Stage 3 concern).
- Neither `forgot-password` nor `reset-password` is itself
  authenticated or returns a session — `forgot-password`'s `202` body
  is just a message, and `reset-password`'s `200` body is just a
  message (session cookies aren't set by this call the way `login`
  does; the spec's "all sessions revoked" means refresh tokens are
  invalidated server-side, not that a new session is issued here).

## Gap

- `useForgotPassword()` needs to exist, wrapping the new
  `forgotPassword` API function from Area 2. Given the file's `202`-only
  contract and no cache implications, this should mirror
  `useRegister`'s bare-`useMutation` shape almost exactly — no
  precedent gap, just needs to be written.
- `useResetPassword()` needs to exist, wrapping `resetPassword`. Unlike
  `useForgotPassword`, this one **does** have a plausible post-success
  side-effect to decide on: the spec's "all sessions revoked" plus
  `patterns.md`'s Form-pattern success convention ("full success
  state/redirect for terminal actions in a flow ... where the user
  needs a clear 'what happens next'") suggests a redirect to `/login`
  after a successful reset, similar in shape (though not in cache
  semantics) to `useLogin`'s `router.push`. No prior hook in this
  domain redirects *without* also writing an access token/cache entry,
  so there isn't an exact one-to-one precedent for "redirect-only, no
  session state to set" — closest structural cousin is still
  `useLogin`/`useLoginMfa`'s `onSuccess` callback shape, just without
  the `applyLoginSuccess` cache-write portion.

## Sniffing

- **Misleading signal**: none.
- **Miscontext**: none — page-map/patterns don't say anything about
  hooks specifically; this layer is purely a code-architecture
  convention (`AGENTS.md`), not something the UX docs weigh in on.
- **Risk**: low — these are the simplest hooks in the domain by
  precedent shape. The one place a mistake could matter: if
  `useResetPassword`'s `onSuccess` redirected unconditionally even
  when the mutation actually resolved via the `422` validation branch
  (should never happen if Area 2's `resetPassword` correctly throws
  rather than resolves on `422`, but worth confirming in Stage 3 that
  the two layers agree on which outcomes count as "success" for
  `useMutation`'s purposes).
- **Edge case**: none beyond what Area 2 already flagged (the `422`
  retry-same-token requirement) — this layer doesn't introduce new
  edge cases of its own, it just needs to not break the ones already
  identified upstream.
- **Inconsistency**: none found.

Proceeding to Area 4 (form components).
