# Task 2 — Session infrastructure: cross-tab refresh coordination + session-guard redirect

> Derived from: `../techplan.md` ("Tech Plan: Login & Session Management
> (Frontend)", account/03-login-session-management). This task file
> redistributes §8-13 detail relevant to its own scope, in full — it does
> not summarize. For the Summary, §1-7 rationale, and §14 Open Items,
> read the source techplan directly.
> Splitting axis: dependency/sequence chain + component boundary (see
> `manifest.md`).
> Dependencies: **none** — this task can start immediately, in parallel
> with Task 1.
> Feeds into: Task 4 (Logout entry point) depends on this task's
> `postAuthChannelMessage` export.
> Recommended model: **GLM 5.2 (max)** — per `best-practices/
> model-routing.md`'s Complex-tier "Coding/build" row and its
> tie-breaker ("GLM when the work leans on diagrams, state-transitions,
> or multi-step reasoning") — this is the plan's highest-novelty,
> highest-correctness-risk task: a browser-concurrency coordination
> mechanism (Web Locks + `BroadcastChannel`) with **zero existing
> precedent anywhere in this codebase**, and the source techplan's own
> §7 lists it as the #1 High-severity risk.

## Scope

Build the cross-tab refresh-coordination mechanism (Web Locks API +
`BroadcastChannel`, resolving the source techplan's D3 and task #02's
own carried-forward Open Item #1), and the session-guard redirect
subscription that reacts to any authenticated→unauthenticated
transition, from any trigger.

**Rules owned by this task** (full text, copied from techplan §4):

- **R11** (coordinated refresh — mutual exclusion): Given any caller
  needs a refresh (`apiFetch`'s 401 handler, or `AuthBootstrapProvider`'s
  boot-time hydration), When `navigator.locks` is available, Then the
  actual `tryRefreshOnce()` call is wrapped in `navigator.locks.request(
  'kencleng-refresh-token', ...)`, serializing concurrent attempts across
  tabs in the same origin instead of letting them race.
- **R12** (coordinated refresh — fallback): Given `navigator.locks` is
  unavailable (unsupported browser, or a non-browser test environment),
  Then `coordinatedRefresh()` falls back to calling `tryRefreshOnce()`
  directly, unserialized — an explicit, accepted degradation, not a
  silent gap.
- **R13** (broadcast on outcome): Given a coordinated refresh completes
  inside the lock, When it succeeds, Then broadcast `{ type: 'refreshed',
  accessToken, accessTokenExpiresAt }` on the `kencleng-auth` channel;
  when it fails, broadcast `{ type: 'refresh-failed' }`.
- **R14** (broadcast reception): Given any tab's `AuthBootstrapProvider`
  -mounted listener receives a `'refreshed'` message, Then it calls
  `setAccessToken` with the broadcast token; given it receives
  `'refresh-failed'` or `'logged-out'`, Then it calls `clearAccessToken`.
  (The `'logged-out'` message is produced by Task 4 — this task's
  listener must handle it correctly even though the *sender* is built in
  a different task; the message contract itself, defined in this task's
  `auth-channel.ts`, is the shared interface.)
- **R15** (session-guard redirect): Given `useAuthStore`'s `accessToken`
  transitions from non-null to null (via any path — R13/R14's failure
  branch, or Task 4's logout), When `SessionGuardProvider`'s subscription
  observes it, Then it redirects the current tab to `/login`.
- **R16** (silent guest, no redirect): Given `accessToken` was already
  `null` (a page load that never had a session, or a refresh that fails
  on first boot — `AuthBootstrapProvider`'s existing R10 case from task
  #02), When the same transition-check runs, Then no redirect fires —
  this must stay indistinguishable from an ordinary guest page load,
  exactly as already required of `AuthBootstrapProvider`.

**Confirmed test-environment fact (from source techplan §14 Resolved
#6, carry this into your own test-writing — do not re-derive it):**
this project's pinned `jsdom` (30.0.1) implements **neither**
`BroadcastChannel` **nor** `navigator.locks`. R12's fallback path is
therefore exercised by jsdom's real, unmodified behavior in tests (a
genuine, representative test, not a mock standing in for reality). R11/
R13/R14's "happy path" (lock available, broadcast actually sent/
received) must be tested by attaching fake `locks`/`BroadcastChannel`
implementations onto the jsdom `window`/`navigator` in test setup, or by
mocking this task's own `lib/api/auth-channel.ts` exports directly — not
by relying on the test environment providing the real APIs.

## Interface Contract (relevant subset of techplan §8)

**API contract consumed:**
```typescript
// POST /auth/refresh (already wrapped by lib/api/client.ts's existing tryRefreshOnce, built by task #02 — not touched by this task, only wrapped)
// no body — reads the kencleng_refresh cookie
// 200 -> RefreshResponse { access_token: string; access_token_expires_at?: string }
// 401 -> Problem (missing/expired/revoked/reuse-detected)
```

**This task's exports:**
```typescript
// lib/api/auth-channel.ts (new)
type AuthChannelMessage =
  | { type: "refreshed"; accessToken: string; accessTokenExpiresAt?: string }
  | { type: "refresh-failed" }
  | { type: "logged-out" };  // sent by Task 4's useLogout, received by this task's own AuthBootstrapProvider listener
export function postAuthChannelMessage(msg: AuthChannelMessage): void; // no-ops if BroadcastChannel unavailable
export function subscribeAuthChannel(handler: (msg: AuthChannelMessage) => void): () => void;

// lib/api/client.ts
export function coordinatedRefresh(): Promise<boolean>; // wraps the existing tryRefreshOnce with navigator.locks (R11/R12) + broadcast (R13)

// components/providers/session-guard-provider.tsx (new)
function SessionGuardProvider({ children }: { children: React.ReactNode }): JSX.Element; // R15/R16
```

**This task's consumers must know:** Task 4's `useLogout` hook imports
and calls `postAuthChannelMessage({ type: "logged-out" })` from this
task's `auth-channel.ts` — that file must exist and export that function
before Task 4's hook can be wired up (hard dependency, not just
sequencing).

**Business logic flow (this task's slice, verbatim from §8):**
```
coordinatedRefresh() (called by apiFetch's 401 handler and AuthBootstrapProvider's boot effect):
  navigator.locks available?
    yes -> locks.request('kencleng-refresh-token', async () => {
             ok = await tryRefreshOnce()
             ok ? postAuthChannelMessage({type:'refreshed', ...}) : postAuthChannelMessage({type:'refresh-failed'})
             return ok
           })
    no  -> tryRefreshOnce() directly (R12, accepted degradation)

AuthBootstrapProvider (extended):
  on mount: existing hydration logic, now calling coordinatedRefresh() (was tryRefreshOnce())
  subscribeAuthChannel(msg => {
    'refreshed'      => setAccessToken(msg.accessToken)
    'refresh-failed' => clearAccessToken()
    'logged-out'     => clearAccessToken()
  })

SessionGuardProvider (mounted root, alongside AuthBootstrapProvider):
  subscribe to useAuthStore
  prevToken = current accessToken at mount
  on each change:
    if prevToken !== null && newToken === null: router.push('/login')   (R15)
    else: no-op                                                          (R16)
    prevToken = newToken
```

## Architecture (relevant notes from §9)

- Cross-tab coordination sits entirely inside `lib/api/client.ts` + the
  new `lib/api/auth-channel.ts` — every caller (401-retry path in
  `apiFetch`, boot hydration in `AuthBootstrapProvider`) goes through
  `coordinatedRefresh()`, never the existing `tryRefreshOnce()` directly
  anymore. `client.ts` is already "the one place... allowed to call
  `fetch` directly" (per `frontend/AGENTS.md`/`api-client-
  centralization.md`) — it is now also the one place cross-tab
  coordination logic lives, rather than being duplicated per-caller.
- `AuthBootstrapProvider` gains a second `useEffect` (the channel
  subscription) alongside its existing mount-time hydration effect —
  same file, same "root session-lifecycle" responsibility, not a new
  provider, since both concerns are about keeping the store correctly
  populated.
- `SessionGuardProvider` (new) is a separate, small provider — a
  different concern (routing, needs `useRouter`) from
  `AuthBootstrapProvider`'s (keeping the store populated), mounted as a
  child in `app/layout.tsx`'s provider stack:
  `<MockingProvider><QueryProvider><AuthBootstrapProvider>
  <SessionGuardProvider>{children}</SessionGuardProvider>
  </AuthBootstrapProvider></QueryProvider></MockingProvider>` — **inside**
  `AuthBootstrapProvider`, so its own boot-time hydration (which may
  itself set `accessToken` to `null` on a failed silent refresh, R16)
  doesn't get misread by `SessionGuardProvider` as an
  authenticated→unauthenticated transition on first mount.
  `SessionGuardProvider` must establish its `prevToken` baseline strictly
  after `AuthBootstrapProvider`'s own first effect has had a chance to
  run — get this ordering wrong and a fresh guest page load will
  spuriously redirect (see Testing Examples below).

## Implementation Details (verbatim from §10)

**File**: `lib/api/auth-channel.ts` (new)
- New module: lazy `BroadcastChannel` singleton, feature-detected
  (`typeof BroadcastChannel !== 'undefined'`), **never constructed at
  module top-level** — this both prevents a Next.js SSR crash (browser
  -only global, absent during server rendering) and matches this
  project's actual jsdom test-environment reality (see the "Confirmed
  test-environment fact" note above — `BroadcastChannel` is `undefined`
  there too, so an eager top-level construction would also break every
  test that imports this module, not just SSR). Exports the typed
  message contract and `postAuthChannelMessage`/`subscribeAuthChannel`.

**File**: `lib/api/client.ts`
- Change: add `coordinatedRefresh()` (R11-R13) wrapping the existing
  `tryRefreshOnce`. Guard `navigator.locks` access the same way as
  `BroadcastChannel` above — feature-detect, never assume presence.
  `apiFetch`'s 401 handler and the exported hydration entry point (used
  by `AuthBootstrapProvider`) both call `coordinatedRefresh()` instead of
  `tryRefreshOnce()` directly from this point on. **Do not delete or
  rename `tryRefreshOnce`** — `coordinatedRefresh()` wraps it, it doesn't
  replace its export (task #02's `AuthBootstrapProvider` and other
  existing callers reference it by that name in their own file's
  history/tests).

**File**: `components/providers/auth-bootstrap-provider.tsx`
- Change: boot-time hydration calls `coordinatedRefresh()` instead of
  `tryRefreshOnce()`; add a second effect subscribing to
  `subscribeAuthChannel` (R14). Keep the two effects independent (don't
  merge them into one `useEffect` with combined dependencies) — the
  existing hydration effect already has a strict "run at most once"
  guard (`attempted.current`) from task #02 that must not be
  accidentally coupled to the channel subscription's own lifecycle
  (which should stay subscribed for the component's entire mounted
  lifetime, not just once).

**File**: `components/providers/session-guard-provider.tsx` (new)
- New Client Component (R15/R16): subscribes to `useAuthStore`, redirects
  via `useRouter().push('/login')` on a non-null→null transition, no-ops
  otherwise. Establish the `prevToken` baseline via the store's *current*
  value read at mount (not assumed to start at `null`), so it correctly
  captures whatever `AuthBootstrapProvider`'s hydration effect has
  already settled by the time this component's own effect runs.

**File**: `app/layout.tsx`
- Change: mount `SessionGuardProvider` inside `AuthBootstrapProvider`
  (see Architecture ordering rationale above).

## Files Changed (this task's rows from §11)

| File | Change Type | Description |
|---|---|---|
| `lib/api/auth-channel.ts` | Add | `BroadcastChannel` wrapper + message contract |
| `lib/api/client.ts` | Modify | Add `coordinatedRefresh()` (R11-R13); switch callers to it |
| `components/providers/auth-bootstrap-provider.tsx` | Modify | Use `coordinatedRefresh`; add channel-listener effect |
| `components/providers/session-guard-provider.tsx` | Add | Redirect-on-session-loss subscription (D4) |
| `app/layout.tsx` | Modify | Mount `SessionGuardProvider` |

## Testing Checklist (this task's items from §12, verbatim)

- [ ] R11: with `navigator.locks` mocked as available, `coordinatedRefresh()` calls `navigator.locks.request` with the named lock before calling the underlying refresh
- [ ] R12: with `navigator.locks` genuinely absent (jsdom's real, unmodified state — see the confirmed-fact note above), `coordinatedRefresh()` still calls the underlying refresh directly (no throw, no hang)
- [ ] R13: a successful coordinated refresh triggers a `postAuthChannelMessage({type:'refreshed', ...})` call (mocked channel); a failed one triggers `{type:'refresh-failed'}`
- [ ] R14: a mocked incoming `'refreshed'` message calls `setAccessToken`; a mocked `'refresh-failed'` or `'logged-out'` message calls `clearAccessToken`
- [ ] R15: a store transition from a non-null token to `null` triggers exactly one `router.push('/login')` call
- [ ] R16: a store that starts at `null` and stays `null` (or a boot-time hydration failure, matching `AuthBootstrapProvider`'s existing R10 case) triggers zero `router.push` calls

**Count-check** (this task's slice): 6 checklist items above, covering R11-R16 exactly.

## Testing Examples & Common Mistakes (relevant rows from §13)

| Mistake | Error/Behavior | Fix |
|---|---|---|
| Calling `tryRefreshOnce()` directly anywhere new instead of `coordinatedRefresh()` | Silently reintroduces the cross-tab race this entire task exists to close | Grep for `tryRefreshOnce(` call sites outside `client.ts` itself during review; only `coordinatedRefresh` should be called externally |
| Constructing `new BroadcastChannel(...)` or referencing `navigator.locks` at module top-level in `auth-channel.ts`/`client.ts` | Crashes/throws during Next.js server-side rendering, **and** breaks every test that imports the module (jsdom has neither global — see confirmed fact above) | Feature-detect and construct lazily, guarded behind `typeof window !== 'undefined'` / `typeof BroadcastChannel !== 'undefined'` / `'locks' in navigator` |
| `SessionGuardProvider` mounted outside/before `AuthBootstrapProvider` | Misreads the boot-time hydration's own possible `null` result as a "session lost" transition, spuriously redirecting a fresh guest page load | Mount ordering per Architecture above — `SessionGuardProvider` inside `AuthBootstrapProvider` |
| Writing a true multi-tab integration test in Vitest, assuming jsdom provides real `BroadcastChannel`/`navigator.locks` behavior | Confirmed directly (source techplan §14 Resolved #6): jsdom 30.0.1 has neither API — such a test would fail to even construct the objects, not silently pass with false confidence | Test the coordination *logic* against fake implementations attached in test setup, or by mocking this task's own `lib/api/auth-channel.ts` exports directly |
