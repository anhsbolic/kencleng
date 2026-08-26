# Task 2: Session Bootstrap / Token Hydration (`AuthBootstrapProvider`)

> Originating contract techplan: `../techplan.md` ("Tech Plan: Google
> OAuth Login/Register (Frontend)", account/02-google-oauth-login-
> register, Status: Draft). Cross-check high-level decisions there
> whenever this task file is ambiguous — this file redistributes, not
> replaces, that document's detail.
>
> Splitting axis: **Component/module boundary** (see `../manifest.md`
> for the full rationale). This task owns the session-hydration
> plumbing that bridges the OAuth callback's HttpOnly-cookie token
> delivery into the SPA's in-memory `useAuthStore`. It has no import/
> code dependency on Task 1 (the Google button/error-banner UI) and
> can be built, reviewed, and tested independently of it.

## Scope

Build the root-level silent-hydration mechanism that populates the
in-memory access token on app load, using the already-registered
`POST /auth/refresh` endpoint. Covers rules **R8–R13** and Decision
Log entry **D3** from the originating techplan.

**Out of scope for this task**: the Google button, `/login`'s page
composition, and the `?error={code}` banner (Task 1's scope). No
change to `/register` or `/login`'s own markup.

## Background (condensed from techplan §1)

The merged backend delivers OAuth-issued session tokens as **HttpOnly
cookies** on the callback's `302` (`writeAuthCookies`,
`backend/internal/transport/http/cookie.go:110-143`) — JS cannot read
them. `apiFetch` (`lib/api/client.ts`) only ever reads an **in-memory**
access token from `useAuthStore`; nothing in the current codebase
bridges the two. `auth-store.ts`'s own doc comment names this task
(alongside backend task #3) as the one that populates that store —
this isn't optional scope.

**Confirmed via direct backend read (not assumed)**: the OAuth
callback's own access-token minting (`google_oauth.go`'s `IssueTokens`)
uses a legacy JWT shape with no `purpose` claim, which the newer
`VerifyAccessToken` (built under backend task #3) rejects by design.
However, `IssueTokens` stores its refresh token in the same
`refresh_tokens` table `POST /auth/refresh`'s handler reads from, and
that endpoint's minting closure (`cmd/server/main.go:122-123`) is
wired to the modern, purpose-claim-bearing `auth.MintAccessToken`. So
the OAuth-issued **refresh** token is fully valid input to
`/auth/refresh`, which mints a fresh, verifiable access token in
return — this is exactly why this task always obtains the frontend's
in-memory token via refresh, never by reading the OAuth cookie
directly (which is also structurally impossible, since it's HttpOnly).

## What this task builds (from techplan §10)

- **File**: `components/providers/auth-bootstrap-provider.tsx` (new)
  — Client Component, same shape/convention as the existing
  `QueryProvider` (`components/providers/query-provider.tsx`): one
  focused responsibility. `useEffect` on mount, guarded by
  `useAuthStore.getState().accessToken === null`, calls the exported
  refresh function exactly once (R8, R11); on success calls
  `setAccessToken` and `queryClient.invalidateQueries({ queryKey:
  accountKeys.me() })` (R9); on failure, no-ops silently (R10).
- **File**: `app/layout.tsx` (modify) — wrap `children` with
  `AuthBootstrapProvider`, placed as a child of `QueryProvider` (so
  `useQueryClient()` is available) which is itself inside
  `MockingProvider` (so MSW is ready first in mock-dev mode — already
  confirmed by direct read of `mocking-provider.tsx`: it gates all
  children behind `ready`) (R13).
- **File**: `lib/api/client.ts` (modify) — export `tryRefreshOnce`
  (currently module-private) for the provider to call; replace the
  hand-written `type RefreshResponse = { access_token: string }` with
  the generated `components["schemas"]["RefreshResponse"]` from
  `schema.d.ts` (also carries `access_token_expires_at`) (R12).
- **File**: `mocks/handlers.ts` (modify) — add
  `http.post("/auth/refresh", ...)`: default happy-path `200` +
  `access_token`; individual tests override via `server.use(...)` for
  the failure case, matching the existing override convention already
  used for `mockUser`/roles.

## Rules & Validation owned by this task

(Numbering matches the originating techplan §4 — not renumbered per
task.)

- **R8** (bootstrap hydration trigger): Given the app loads on any
  route and `useAuthStore.accessToken` is `null`, When the root
  `AuthBootstrapProvider` mounts, Then it calls the refresh mechanism
  exactly once.
- **R9** (hydration success): Given the refresh call succeeds, Then
  `useAuthStore.setAccessToken` is called with the returned token, and
  the `accountKeys.me()` query (`lib/hooks/use-account-me.ts`) is
  invalidated/refetched — so a component that queried it before
  hydration completed doesn't keep rendering stale logged-out state.
- **R10** (hydration failure is silent): Given the refresh call fails
  (genuinely logged-out guest, no valid refresh cookie present), Then
  `accessToken` stays `null` and no error/toast is shown — this must
  be indistinguishable from an ordinary guest page load.
- **R11** (at most one attempt): Given the bootstrap provider's
  hydration attempt, Then it runs at most once per app load — no
  retry loop, no re-trigger on client-side route change within the
  same load.
- **R12** (generated type, not hand-written): Given `client.ts`'s
  refresh-response parsing, When touched by this task, Then it uses
  `components["schemas"]["RefreshResponse"]`, not the existing
  hand-written local type.
- **R13** (provider placement): Given `AuthBootstrapProvider` needs
  `useQueryClient()` (R9), Then it is mounted as a child of
  `QueryProvider` in `app/layout.tsx`'s provider stack.

## Decision Log entries relevant to this task

**D3 — Token hydration mechanism**

| Option | Why rejected/accepted |
|---|---|
| A. Route-specific "just came from OAuth" detection | Rejected — the success redirect carries no query signal to detect this by (confirmed via direct backend read: `successResult` returns the bare frontend URL, nothing appended). Any such mechanism collapses into "always try" anyway. |
| B. Unconditional silent-refresh-on-app-boot, root-level provider (**chosen**) | Not OAuth-specific scope creep: the access token is deliberately in-memory-only, so *any* page refresh already loses it — this task is simply the first to need the general bootstrap mechanism, which backend task #3's own login flow needs too (`auth-store.ts`'s "Tasks #2 and #3 share one store shape" comment). Matches `pwa/token-storage-and-refresh.md`'s documented pattern near-verbatim. |
| C. Rely on `apiFetch`'s existing reactive 401→refresh→retry path, no new code | Rejected — only fires after something has already attempted an authenticated call and gotten a `401`, producing a visible flash of logged-out UI on first render instead of hydrating before paint. |

## Interface Contract (subset relevant to this task)

```typescript
// API contract consumed (already registered/implemented — cmd/server/main.go:162)
// POST /auth/refresh
// 200 -> RefreshResponse { access_token: string; access_token_expires_at?: string }
// Reads the refresh token from its own HttpOnly cookie, not the request body.

// lib/api/client.ts
export { tryRefreshOnce }; // was module-private; now exported for AuthBootstrapProvider (R12: uses generated RefreshResponse type, not a hand-written one)

// components/providers/auth-bootstrap-provider.tsx
function AuthBootstrapProvider({ children }: { children: React.ReactNode }): JSX.Element; // R8-R11, R13
```

**Business logic flow (this task's slice):**
```
AuthBootstrapProvider (mounted once, root layout, inside QueryProvider)
  -> on mount, if useAuthStore.accessToken === null:
       -> call tryRefreshOnce() exactly once (R8, R11)
       -> success => setAccessToken(token) + invalidateQueries(accountKeys.me())  (R9)
       -> failure => no-op, no UI signal (R10)
```

## Architecture note (from techplan §9, this task's slice)

No new TanStack Query mutation/query hook is needed for the hydration
call itself — it's a one-shot imperative action on mount (`useEffect`
+ the exported `tryRefreshOnce`), not something that benefits from
query-cache semantics; `accountKeys.me()` (existing, built by task #1
of this domain) is what gets invalidated as a side effect (R9), not a
new query key. Success has no dedicated landing page: the callback's
`302` on success targets the bare frontend root, and
`AuthBootstrapProvider`'s root placement means hydration runs
regardless of which page the user lands on — no route-specific
"callback success" surface is needed.

## Backward Compatibility

- **Database**: N/A — no persistence layer in `frontend/`.
- **API**: No API changes; consumes the already-registered
  `POST /auth/refresh` contract.
- **Existing clients/data**: `tryRefreshOnce` already exists in
  `lib/api/client.ts` today (module-private, used reactively by
  `apiFetch`'s 401 handler) — this task only exports it and swaps its
  local type for the generated one; no behavioral change to its
  existing reactive use.

## Edge Cases & Risks relevant to this task

| Risk | Likelihood | Severity | Mitigation |
|---|---|---|---|
| Hydration bridge missing/broken — user appears logged in (redirect succeeded, cookies set) but the SPA silently sends unauthenticated requests | Medium — the failure is invisible without deliberately checking `Authorization` headers | **High** — the entire feature appears broken from the user's perspective with no error surfaced | R8/R9/R11, dedicated test asserting `setAccessToken` is called and `account.me` is invalidated after a successful mocked refresh |
| Hydration silent-failure path (R10) accidentally surfaces an error/toast to a genuinely logged-out guest browsing `/` | Medium — easy to wire the failure branch the same way as a "real" error by habit | Medium — confusing, unwarranted "session expired"-style message on an ordinary first visit | R10 + dedicated test: mocked refresh failure renders no banner/toast anywhere |
| `AuthBootstrapProvider` mounted outside `QueryProvider`'s subtree | Low — straightforward to get right once specified | Medium — `useQueryClient()` throws/is unavailable, R9's cache invalidation silently fails to compile or run | R13: explicit placement instruction, verified in code review |
| Multi-tab desync: a user completes Google login in one tab; other open tabs never learn about it (no `BroadcastChannel`/equivalent) | Medium — plausible whenever a user has the app open in two tabs during login | Medium — a second tab keeps behaving as logged-out until its own reload/hydration | **Not mitigated in this task** — flagged as an Open Item below, scope decision needed |

## Files Changed / NOT Changed (this task's subset)

| File | Change Type | Description |
|---|---|---|
| `components/providers/auth-bootstrap-provider.tsx` | Add | Silent-refresh-on-boot hydration (R8-R11, R13) |
| `app/layout.tsx` | Modify | Mount `AuthBootstrapProvider` (R13) |
| `lib/api/client.ts` | Modify | Export refresh function; generated `RefreshResponse` type (R12) |
| `mocks/handlers.ts` | Modify | Add `POST /auth/refresh` handler |
| Corresponding `*.test.tsx` for the files above | Add | Per Testing Checklist below |

| File | Reason untouched (this task) |
|---|---|
| `components/features/account/google-auth-button.tsx`, `google-callback-error.tsx`, `register-form.tsx`, `app/(auth)/login/page.tsx` | Task 1's scope |
| `lib/api/account.ts` | `/auth/refresh` is wrapped in `client.ts`, not `account.ts` — no change needed here |
| `lib/hooks/` (no new hook file) | Hydration is a one-shot imperative effect, not a query/mutation hook |
| Anything under `backend/` | Directory-boundary rule, root `AGENTS.md` §7 — out of scope for a `frontend/`-scoped session; findings referenced here are read-only cross-checks |

## Testing Checklist (this task's subset)

- [ ] R8: with a mocked `POST /auth/refresh` and `accessToken` initially `null`, `AuthBootstrapProvider` calls it exactly once on mount
- [ ] R9: a mocked refresh success calls `setAccessToken` with the response's `access_token` and triggers an `account.me` query invalidation/refetch
- [ ] R10: a mocked refresh failure (e.g. `401`) leaves `accessToken` as `null` and renders no error/toast anywhere
- [ ] R11: re-rendering/remounting within the same app load does not trigger a second refresh call (mock call-count assertion)
- [ ] R12: `client.ts`'s refresh-response parsing type-checks against `components["schemas"]["RefreshResponse"]`, not a local duplicate type
- [ ] R13: `AuthBootstrapProvider` is a descendant of `QueryProvider` in `app/layout.tsx` (structural/render test — `useQueryClient()` resolves without throwing)

## Testing Examples & Common Mistakes (this task's subset)

| Mistake | Error/Behavior | Fix |
|---|---|---|
| `AuthBootstrapProvider` mounted outside `QueryProvider`'s subtree | `useQueryClient()` throws or the invalidation silently no-ops | Mount order per R13 — inside `QueryProvider`, inside `MockingProvider` |
| Treating a failed silent refresh as an error to surface to the user | A guest browsing `/` for the first time sees a spurious "session expired"/error toast | R10 — failure is always silent; only a genuinely triggered action (e.g. a stale-session 401 on an authenticated page) should ever show session-related messaging |
| Calling refresh on every render/mount instead of once per app load | Redundant network calls, possible refresh-token rotation races if fired concurrently | R11 — guard with a ref/module-level flag, same "no retry loop" discipline as `apiFetch`'s own 401 handling |

## Open Items relevant to this task

- **Multi-tab session sync** (originating techplan §14, Active #1) —
  `pwa/token-storage-and-refresh.md`'s best-practices checklist calls
  for `BroadcastChannel` (or equivalent) so a login completed in one
  tab is reflected in other open tabs. Not built in this task —
  genuinely new scope not raised in the raw exploration docs, needs a
  scope call from whoever reviews the originating techplan: build it
  now as part of this task, or explicitly defer until `logout`
  (backend task #3) exists and tab-sync becomes unavoidable either way.
