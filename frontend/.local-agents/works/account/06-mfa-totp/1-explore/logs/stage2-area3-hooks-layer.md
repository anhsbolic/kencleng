# Stage 2 — Area 3: Hooks layer (`lib/hooks/`)

## Current state

No `use-mfa-*.ts` hooks exist yet (confirmed — `lib/hooks/` file listing
from Stage 1 shows only `use-login-mfa.ts`, which wraps the *login-time*
`POST /auth/login/mfa` from feature 03, unrelated to this task's 3
enroll/confirm/disable endpoints).

Three existing hooks establish the conventions this task's hooks should
follow, all read in full:

- **`useUnlinkGoogle`** (simplest shape): a bare `useMutation({ mutationFn,
  onSuccess })` where `onSuccess` is a one-line
  `queryClient.invalidateQueries({ queryKey: accountKeys.me() })` inline
  — used when there's no branching logic worth extracting.
- **`useSetPassword`**: extracts its `onSuccess` logic into a standalone,
  exported `applySetPasswordSuccess(result, queryClient)` function —
  specifically so the branching (invalidate vs. clear-session-redirect)
  is unit-testable without a hook-rendering harness. Doc comment
  explicitly states this mirrors `use-login.ts`'s `applyLoginSuccess`
  extraction for the same reason.
- **`useLoginMfa`**: doesn't re-implement success handling at all —
  reuses `applyLoginSuccess` from `use-login.ts` directly, since login
  and login-mfa share the exact same "what happens on success" logic.

All three hooks: `useMutation` from `@tanstack/react-query`, mutation
function imported directly from `lib/api/account.ts`, `accountKeys.me()`
(the shared query-key factory from `use-account-me.ts`) is the
invalidation target whenever `User` data changes as a result.

## Requirement

Three mutation hooks are needed: enroll, confirm, disable. Each
corresponds 1:1 to a wrapper function from Area 2.

## Gap

- `useMfaEnroll` — no state on the `User` object changes on this call
  (`mfa_enabled` stays `false` until confirm; a pending/unconfirmed
  secret isn't reflected in `useAccountMe()` at all per Area 1's
  finding), so this is likely the simplest shape: a bare mutation with
  no `onSuccess` cache invalidation, just returning `MfaEnrollResponse`
  for the component to render as a QR code. Component still needs to
  hold onto `otpauth_uri` client-side across the enroll→confirm step
  (React state in `MfaSection`, not the query cache — nothing else
  reads it).
- `useMfaEnrollConfirm` — success flips `mfa_enabled` to `true`; needs
  `accountKeys.me()` invalidation like `useUnlinkGoogle`'s pattern.
  Also needs to surface `backup_codes` (shown once) back to the caller
  — since `useMutation`'s own return value already carries the
  resolved data via `mutation.data`, no extra plumbing needed for that
  part, just confirming the shape is used that way rather than
  discarded in an unused `onSuccess`.
- `useMfaDisable` — success flips `mfa_enabled` to `false`; same
  `accountKeys.me()` invalidation as `useUnlinkGoogle`.
- None of the three need the extracted-standalone-function treatment
  `useSetPassword` uses, since none branch into multiple divergent
  outcomes the way `setPassword`'s add-vs-change or `login`'s
  ok-vs-mfa_required do (open to revisit at Stage 3 if `mfaDisable`'s
  two call shapes — password vs. no-body/Google-reauth — turn out to
  need shared branching logic worth extracting).

## Page-consolidation check

N/A — this is a hooks-only area, no route/endpoint mismatch risk beyond
what Area 2 already covered (the 3 endpoints are confirmed to map 1:1
to the feature spec, no orphans).

## Sniffing

- **Miscontext risk**: unlike `useSetPassword`/`useLoginMfa`, none of
  the three new hooks have an existing sibling with identical
  branching logic to reuse outright — each is genuinely new, so there's
  a temptation to over-engineer a shared abstraction across all three
  (e.g. a generic `useMfaMutation` factory) where the existing
  convention actually favors three small, independent, purpose-named
  hooks (matches `page.tsx`'s own "independent section components, not
  a monolithic form" philosophy, D1, applied one layer down). Worth
  keeping in mind for Stage 3 rather than assuming shared abstraction
  is automatically better.
- **Edge case**: `useMfaEnroll`'s returned `otpauth_uri` needs to
  survive a re-render/route-remains-mounted cycle while the user scans
  the QR code and enters a code — since it's not cached in TanStack
  Query (no `onSuccess` writing it anywhere) and not in a Zustand
  store, it only lives in the mutation's own `.data` field or component
  state. If `MfaSection` calls `mutate()` again for any reason (e.g. a
  "regenerate QR" retry button) before `enrollConfirm` succeeds, the
  previous `otpauth_uri` is simply replaced — matches the backend's own
  "overwrite the pending secret" semantics (spec's second acceptance
  criterion for `enroll`), so this appears to be the *correct* behavior
  by coincidence, not a hook design flaw. Worth confirming this
  reasoning explicitly in Stage 3 rather than leaving it implicit.
- No inconsistency or high risk found beyond what's already noted in
  Area 1/2 (the reauth-marker round-trip gap, and the missing-409-in-
  schema gap) — this area is a straightforward "write three hooks
  matching an existing, clear pattern."
</content>
