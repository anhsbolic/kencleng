# Task 4 — Logout entry point

> Derived from: `../techplan.md` ("Tech Plan: Login & Session Management
> (Frontend)", account/03-login-session-management). This task file
> redistributes §8-13 detail relevant to its own scope, in full — it does
> not summarize. For the Summary, §1-7 rationale, and §14 Open Items,
> read the source techplan directly.
> Splitting axis: dependency/sequence chain + component boundary (see
> `manifest.md`).
> Dependencies: **Task 2** (Session infrastructure) — this task imports
> `postAuthChannelMessage` from Task 2's `lib/api/auth-channel.ts`, and
> its redirect-after-logout behavior is provided end-to-end by Task 2's
> `SessionGuardProvider`, not by this task itself (R17 explicitly
> excludes navigation from this task's own scope).
> Shared file note: `lib/api/account.ts` and `mocks/handlers.ts` are also
> touched by Task 1 — this task owns the `logout` addition to those
> files; Task 1 owns the `login`/`loginMfa` additions. No overlapping
> lines, but sequence this task after Task 1 if applying both, to avoid
> an avoidable merge conflict on the same file.
> Recommended model: **DeepSeek V4 Pro** — per `best-practices/
> model-routing.md`'s Complex-tier "Coding/build" row and its
> tie-breaker ("DeepSeek V4 Pro when it's rule-table-heavy/precision work
> without a diagram") — this task is the smallest, most mechanical slice
> of the plan: an idempotent mutation, an unconditional cleanup sequence,
> and a single gated button, with no branching logic of its own.

## Scope

Add `POST /auth/logout`'s fetch function + hook, a default mock handler,
and the "Keluar" button in the Dashboard Shell that triggers it.

**Rules owned by this task** (full text, copied from techplan §4):

- **R17** (logout action): Given an authenticated user (`useAccountMe()`
  has data) clicks "Keluar" in `DashboardShellClient`, When
  `POST /auth/logout` settles (success or failure — logout is
  idempotent/always-`204` per spec 03), Then unconditionally:
  `clearAccessToken()`, `queryClient.clear()` (full cache reset —
  `pwa/state-management-boundaries.md`'s "no store forgotten" checklist
  item, not just `account.me`), and broadcast `{ type: 'logged-out' }`
  on the auth channel (via Task 2's `postAuthChannelMessage`). Task 2's
  `SessionGuardProvider` handles the actual redirect — **this task's own
  handler must not separately navigate** (no `router.push` call
  anywhere in this task's code).
- **R18** (logout button visibility): Given `useAccountMe()`'s data is
  not yet loaded, is loading, or errored, Then the "Keluar" button does
  not render — gated on the same primitive as nav-item role-filtering
  (`useHasRole`'s own "safe default: hide" discipline, already
  established in `lib/hooks/use-has-role.ts`), but **not** via a role
  array (logout applies to any authenticated user, not a role subset) —
  gate directly on `useAccountMe()`'s `data` being truthy, not on
  `useHasRole(...)`.
- **R19** (mocks) — **this task's slice**: `mocks/handlers.ts` gains a
  default handler for `POST /auth/logout` (`204`) — individual tests
  override via `server.use(...)` if a failure-path test is needed (per
  R17, the handler's own logic should behave identically either way, so
  a failure-path test mainly proves the `onSettled` unconditional-cleanup
  guarantee, not a different rendering outcome).

## Interface Contract (relevant subset of techplan §8)

**API contract consumed:**
```typescript
// POST /auth/logout
// no body
// 204, always (idempotent — no cookie present is not an error)
```

**This task's exports:**
```typescript
// lib/api/account.ts
export function logout(): Promise<void>; // always resolves; network failure normalized via the existing postAccountAction helper's ApiError(0) path — this task's hook treats any settle (resolve OR the normalized-error path) identically, per D5

// lib/hooks/use-logout.ts (new)
export function useLogout(): UseMutationResult<void, never, void>; // never rejects from the caller's perspective — always settles "successfully" in the sense that onSettled always runs the cleanup (D5)
```

**This task consumes from Task 2:**
```typescript
import { postAuthChannelMessage } from "@/lib/api/auth-channel";
// called as: postAuthChannelMessage({ type: "logged-out" })
```

**Business logic flow (this task's slice, verbatim from §8):**
```
useLogout (DashboardShellClient's "Keluar" button):
  onSettled (success or failure, always):
    clearAccessToken()
    queryClient.clear()
    postAuthChannelMessage({type:'logged-out'})
    (no direct navigation — SessionGuardProvider handles it, R15, built in Task 2)
```

## Architecture (relevant note from §9)

`DashboardShellClient`'s header (`app/(dashboard)/_components/
dashboard-shell-client.tsx`) gains a "Keluar" button (`Button
variant="outline" size="sm"`) next to `NotificationBadge`, gated on
`useAccountMe()`'s `data` being truthy (R18) — a new, small,
self-contained addition to an existing component, not a structural
Shell change (do not touch `nav-items.ts` or `FilteredNavLinks`/
`NavLink` — logout is a mutation trigger, not a navigation link, and
forcing it through that component's `<Link href>` shape would need a
special-cased branch inside a component whose whole job today is
"render a role-filtered link, nothing else").

## Implementation Details (verbatim from §10)

**File**: `lib/api/account.ts`
- Change: add `logout` (R19), reusing the existing `postAccountAction`
  helper. Unlike `login`/`loginMfa` (Task 1), `logout` needs no
  discriminated result — it resolves `void` and lets the *caller* (this
  task's `use-logout.ts`) unconditionally clear local state regardless
  of network outcome, matching root `AGENTS.md`'s "the frontend has no
  business logic" boundary: whether logout "succeeded" server-side isn't
  something the client re-derives or blocks on.

**File**: `lib/hooks/use-logout.ts` (new)
- `useMutation` wrapping `logout`; `onSettled` performs R17's
  unconditional cleanup (D5): `clearAccessToken()`,
  `queryClient.clear()`, `postAuthChannelMessage({type:'logged-out'})`
  — in that order, all three, every time `onSettled` fires, with no
  conditional branch on whether the mutation resolved or the underlying
  fetch failed.

**File**: `mocks/handlers.ts`
- Change: add a default handler for `POST /auth/logout` (`204`) (R19).

**File**: `app/(dashboard)/_components/dashboard-shell-client.tsx`
- Change: add the "Keluar" button (R17/R18), calling `useLogout()`'s
  `.mutate()` on click. No loading/disabled state is strictly required
  by any rule (logout has no field validation and settles quickly), but
  a brief `loading` prop on the `Button` (mirroring the existing
  `Button`'s `loading` prop already used elsewhere, e.g.
  `RegisterForm`'s submit button) is reasonable polish, not a rule
  violation either way if omitted.

## Files Changed (this task's rows from §11)

| File | Change Type | Description |
|---|---|---|
| `lib/api/account.ts` | Modify | Add `logout` (shared file — Task 1 also modifies, adding `login`/`loginMfa`; no line overlap) |
| `lib/hooks/use-logout.ts` | Add | Logout mutation, unconditional cleanup (D5) |
| `mocks/handlers.ts` | Modify | Add one default handler (shared file — Task 1 also modifies, adding the login/MFA handlers; no line overlap) |
| `app/(dashboard)/_components/dashboard-shell-client.tsx` | Modify | Add "Keluar" button |

## Testing Checklist (this task's items from §12, verbatim)

- [ ] R17: clicking "Keluar" — on both a mocked logout success and a mocked logout failure — results in `clearAccessToken`, `queryClient.clear()`, and a `'logged-out'` broadcast (assert `postAuthChannelMessage` was called with the right payload — mock Task 2's `auth-channel.ts` module for this test, don't require a real `BroadcastChannel`), in both cases identically
- [ ] R18: the "Keluar" button is absent while `useAccountMe()` is loading/errored/has no data, present once it resolves with data
- [ ] R19: `mocks/handlers.ts`'s new default handler for `/auth/logout` resolves `204`; confirm a test can override it via `server.use(...)` for the failure-path assertion in R17

**Count-check** (this task's slice): 3 checklist items above, covering R17-R19.

## Testing Examples & Common Mistakes (relevant rows from §13)

| Mistake | Error/Behavior | Fix |
|---|---|---|
| Treating a failed `POST /auth/logout` call as "logout didn't happen," leaving the user stuck | The whole point of D5 — logout must clear local state regardless of the network outcome | R17 — use `onSettled`, not `onSuccess` |
| Redirecting directly from this task's `onSettled` (e.g. calling `router.push('/login')` inline) in addition to Task 2's `SessionGuardProvider` subscription | Double-navigation / race between two redirect triggers for the same event | R17 explicitly excludes navigation from this task's own handler — `SessionGuardProvider` (Task 2) is the sole redirect path |
| Gating the "Keluar" button via `useHasRole([...])` with some role array, copying `nav-items.ts`'s pattern | Logout should be visible to *any* authenticated user regardless of role — a role array is the wrong primitive for "is anyone logged in at all" | R18 — gate directly on `useAccountMe()`'s `data` |
