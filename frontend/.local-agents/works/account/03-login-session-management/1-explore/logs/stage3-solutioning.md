# Stage 3 — Solutioning: account/03-login-session-management (frontend)

> Builds on `stage2-gap-analysis.md`'s cross-area summary. Each item
> below is a Decision Log entry: question → options considered →
> recommendation → rationale. This is raw solutioning material, not a
> techplan — synthesis into a techplan is a separate, later step.

---

## D1 — `/login` page composition: how the credential form joins the existing Google button

**Question:** Stage 2 Area 1 confirmed `/login`'s own doc comment
names this task as the one that "extends this page (adds the form
above/alongside the Google button + divider, matching RegisterForm's
own composition), not replaces it." What does that composition
actually look like once both the password form and the (new) MFA step
exist?

**Options considered:**
- **A. One `LoginForm` client component** (`components/features/
  account/login-form.tsx`, mirroring `register-form.tsx`'s file/
  naming shape) that **internally owns `GoogleAuthButton` + divider**,
  the same way `RegisterForm` embeds its own Google button today —
  `/login/page.tsx` shrinks to just rendering `<GoogleCallbackError />`
  then `<LoginForm />`, dropping the current static Google button +
  "coming soon" banner entirely.
- **B. Keep `/login/page.tsx` composing the pieces**: a new, narrower
  `CredentialLoginForm` (fields + submit only) rendered above the
  page's *existing* `GoogleAuthButton` + divider markup, left in place.
- **C. Two separate page-level components** for the password step and
  the MFA step, with `/login/page.tsx` itself holding the
  `'password' | 'mfa'` step state and conditionally rendering one or
  the other.

**Recommendation: A.**

**Rationale:**
- Matches `RegisterForm`'s actual, already-merged shape exactly (self-
  contained form component owning its own Google entry point below a
  divider) — the login page's own comment explicitly says to match
  that composition, not invent a different split.
- Option B would leave Google-button markup duplicated across two
  places in spirit (the page's static JSX today, plus whatever the new
  form renders around it) for no benefit — nothing about the MFA step
  needs the Google button to live at the page level instead of inside
  the form component.
- Option C lifts step state to the page unnecessarily — `RegisterForm`
  already established the precedent of a form component owning its own
  internal multi-state UI (idle form → success view, via local
  `useState`); the login form's `'password' | 'mfa'` step is the same
  shape of problem, one level more complex, not a different one.
- `GoogleAuthButton` is only relevant during the password step (mid-
  MFA, showing "Masuk dengan Google" would be confusing — the user is
  already mid-credential-flow with a different identity provider
  entirely) — Option A's internal ownership makes gating it behind
  `step === 'password'` a one-line concern local to the component that
  already knows its own step, rather than something the page has to
  coordinate.
- `login/page.test.tsx`'s current negative assertions (`queryByLabelText`
  must NOT find email/password) must be replaced, not left in place —
  flagging explicitly since Stage 2 Area 1 found this test will
  otherwise fail the moment fields are added, and a red test at the
  start of a build is easy to mistake for "test still needs writing"
  when it's actually "test needs rewriting."

---

## D2 — MFA challenge step: state shape and where `mfa_pending_token` lives

**Question:** Stage 2 Area 1 confirmed no existing surface — page-map.md
route, component, or store — models the post-password/pre-MFA state.
Spec 03's `mfa_pending_token` is a 5-minute-TTL value the frontend must
hold between two submits.

**Options considered:**
- **A. Local component `useState`** inside `LoginForm` (from D1):
  `const [step, setStep] = useState<'password' | 'mfa'>('password')`
  plus `const [pendingToken, setPendingToken] = useState<string | null>(null)`.
  A page refresh between steps loses the token, forcing the user back
  to the password step (re-enter email/password).
  - `useAuthStore` (Zustand). Rejected — Stage 2 Area 2 confirmed this
    store's own doc comment is explicit that it's a deliberately
    unpersisted, access-token-only store; folding in an unrelated,
    much-shorter-lived value onto the same store conflates two
    different lifetimes and widens the store's responsibility for no
    benefit, since nothing outside `LoginForm` ever needs to read
    `pendingToken`.
  - `sessionStorage`. Rejected — survives a refresh (arguably a
    feature), but the token is short-lived (5 min) precisely because
    the backend spec wants a bounded exposure window "if it ends up
    somewhere it shouldn't... a proxy log, etc." (spec 03 Assumption
    A/B). Persisting it to `sessionStorage`, even briefly, reintroduces
    a version of the exact risk the backend already designed around by
    keeping it out of anywhere a device-level inspection or another
    script on the page could read it after the fact. In-memory-only
    matches the same rationale `auth-store.ts` already uses for the
    access token itself.
- **B. `sessionStorage`** — survives refresh, but rejected per the
  sub-bullet above.

**Recommendation: A.**

**Rationale:**
- Directly reuses `RegisterForm`'s own precedent (plain `useState` for
  the equally-transient "which view am I showing" + "what email did I
  submit" pair, in that component's success-view swap) — no new state-
  management pattern introduced.
- A lost-token-on-refresh mid-MFA-step is an acceptable UX cost
  (re-enter password, get a new 5-minute-fresh token) given the
  alternative reintroduces a version of the exact exposure risk the
  backend already designed the short TTL to bound. Confirm this
  trade-off explicitly with product/Anhar since it is a UX
  regression vs. "always survives refresh" — flagging as an open item
  for the techplan step, not silently assumed.

---

## D3 — Cross-tab refresh coordination (spec 03 Assumption D)

**Question:** Stage 2 Area 2 found this is the highest-severity gap
overall: with zero coordination, two tabs both triggering
`POST /auth/refresh` near the same access-token expiry causes the
backend's reuse-detection to revoke the *entire* token family — an
unwanted forced logout for any donor with two tabs open, not a
hypothetical. Spec 03's own Assumption D names `BroadcastChannel` as
the mechanism ("one tab acts as the single 'refresher', others wait
and receive the rotated access token via the channel").

**Options considered:**
- **A. Literal `BroadcastChannel`-only election**: each tab, on
  needing a refresh, posts a claim message and waits briefly for a
  competing claim before proceeding. Rejected as the sole mechanism —
  `BroadcastChannel` delivery is asynchronous (a `postMessage` isn't
  guaranteed to be observed by other tabs before either side proceeds
  past its own check-then-act step), so a hand-rolled election over it
  alone cannot fully close the exact race this is meant to prevent
  without significant extra protocol complexity (sequence numbers,
  timeouts, tie-breaking) — the failure mode it's meant to fix
  (concurrent refresh) is precisely a race condition, so the fix
  itself must not have its own race window.
- **B. Web Locks API (`navigator.locks.request`) for mutual exclusion
  + `BroadcastChannel` for fan-out.** One tab acquires a named lock
  (`"kencleng-refresh-token"`) before calling `tryRefreshOnce()`; the
  lock request itself queues automatically across tabs in the same
  origin (browser-native, race-free by construction — this is exactly
  what the Locks API exists for). The lock-holding tab performs the
  actual refresh, then **broadcasts the resulting access token** (or
  failure) over a `BroadcastChannel`; every other tab, instead of
  separately requesting the lock and re-doing the network call, first
  checks the store, and if a refresh it's waiting on completes via a
  received broadcast message, uses that instead of calling
  `tryRefreshOnce()` itself.
- **C. Defer entirely** — leave as a known gap, ship this task without
  it. Rejected — spec 03 Assumption D explicitly assigns this exact
  gap to "when the account domain's frontend track starts," which is
  this task; shipping login/session-management without it means
  shipping the exact bug the backend spec already anticipated and
  explicitly deferred here, not a new discovery to defer further.

**Recommendation: B, flagged as an explicit deviation from spec 03's
literal wording — needs human confirmation before being treated as
final.**

**Rationale:**
- Web Locks API is supported in all evergreen browsers relevant to
  this project's PWA target (Chrome/Edge, Firefox, Safari 15.4+) — not
  an exotic dependency, and purpose-built for exactly this cross-tab
  mutual-exclusion problem (this is a well-known pattern for
  coordinating token refresh across tabs, not a novel design here).
- `BroadcastChannel` is still central to the design — it's just doing
  the job spec 03's own wording emphasizes ("others wait and receive
  the rotated access token via the channel"), while the *leader
  election* half (spec 03's "one tab acts as the single refresher")
  is handled by a primitive better suited to it than a hand-rolled
  protocol over `BroadcastChannel` alone would be.
- **Flagging explicitly per root `AGENTS.md`'s "ambiguity is recorded,
  not silently resolved" rule**: spec 03 names `BroadcastChannel`
  specifically, not Web Locks. This recommendation satisfies the
  spec's stated *goal* (no more than one tab ever calls
  `/auth/refresh` concurrently) through a different specific mechanism
  than the one named. Needs Anhar's sign-off before the techplan step
  treats this as settled, since it's a deviation from what a
  human-reviewed spec document literally says, even though it doesn't
  change the reviewed backend contract at all (the backend side of
  Assumption D is unaffected either way — this is purely a frontend
  implementation-mechanism choice).
- Scope boundary: this coordination lives inside/alongside
  `lib/api/client.ts`'s existing `tryRefreshOnce`/`apiFetch`, not
  duplicated per-caller — every consumer (the 401-retry path,
  `AuthBootstrapProvider`, and any future manual "refresh now" call)
  goes through the same coordinated path, consistent with `client.ts`
  already being "the one place... allowed to call `fetch` directly."

---

## D4 — API layer & hooks for the four endpoints

**Question:** Stage 2 Area 3 confirmed `schema.d.ts` already has the
complete, correct contract (cross-checked directly against the actual
backend handler code in `auth_login.go` — the `LoginMfaRequiredResponse`
`# INFERRED` marker is resolved/accurate, no drift found), but zero
fetch functions, hooks, or mocks exist for any of the four endpoints.
What shape should each take?

**Recommendation (no live alternatives — this follows established,
already-merged precedent directly, listed for completeness rather than
as a contested decision):**
- `lib/api/account.ts` gains `login`, `loginMfa`, `logout` following
  `register`'s existing pattern: reuse `postAccountAction` for the
  network-error normalization, return a discriminated result (`{ ok:
  true, data: LoginResponse } | { ok: true, kind: 'mfa_required',
  mfa_pending_token } | ...`) or throw `ApiError` for anything not
  explicitly modeled — mirroring `RegisterResult`'s discriminated-union
  precedent (Decision D4 of task 01's own solutioning) rather than
  inventing a new result shape.
  `logout` needs no discriminated result — spec 03 defines it as
  always-`204`/idempotent, so the function can resolve `void` and let
  the *caller* (the logout hook, D5) unconditionally clear local state
  regardless of network outcome, matching root `AGENTS.md`'s "the
  frontend has no business logic" boundary: whether logout "succeeded"
  server-side isn't something the client re-derives or blocks on.
- `refresh` gets **no new public wrapper** — `tryRefreshOnce` (extended
  per D3) remains `client.ts`-internal, exported only for
  `AuthBootstrapProvider`'s existing use. No other caller identified in
  this exploration that needs a standalone `refresh()` API-layer
  function; adding one speculatively would be scope beyond what any
  page/component in this task actually needs (lowest-complexity
  principle — same reasoning `mocks/handlers.ts`'s own "not
  speculatively ahead of demonstrated need" comment already states for
  this codebase generally).
- `lib/hooks/` gains `use-login.ts` and `use-login-mfa.ts`
  (`useMutation`, `mutationFn: login`/`loginMfa`) and `use-logout.ts`.
  Unlike `useRegister` (no cache side effects, since nothing is
  authenticated yet at that point), a successful `login`/`loginMfa`
  **must** call `useAuthStore.getState().setAccessToken(...)` and
  `queryClient.invalidateQueries({ queryKey: accountKeys.me() })` —
  directly reusing the exact pair `AuthBootstrapProvider` already
  performs on its own successful hydration, so the "just logged in"
  and "just silently hydrated" paths converge on the same two
  side-effects rather than each inventing its own. `useLogout`
  performs the mirror image unconditionally in an `onSettled`:
  `clearAccessToken()` + `queryClient.removeQueries({ queryKey:
  accountKeys.me() })` (removed, not just invalidated — an
  unauthenticated `GET /account/me` should not silently refetch and
  repopulate with a 401-driven empty/error state that some component
  might render oddly; removing the cache entry is the cleaner "there is
  no user" signal).
- `mocks/handlers.ts` gains three handlers (`/auth/login`,
  `/auth/login/mfa`, `/auth/logout`) with a sensible default success
  case each, following the file's own stated convention ("one handler
  per endpoint, added as the page/component that needs it gets
  built") — individual tests override failure branches via
  `server.use(...)`, matching how `/auth/refresh`'s default-`401`
  case is already handled today.

---

## D5 — Logout entry point in Dashboard Shell

**Question:** Stage 2 Area 4 confirmed no logout affordance exists
anywhere, and `nav-items.ts`'s existing role-array shape doesn't
naturally express "just needs to be logged in," nor is logout a
navigable link.

**Options considered:**
- **A. A single "Keluar" button in `DashboardShellClient`'s header**,
  next to `NotificationBadge`, gated on `useAccountMe()`'s data being
  present (not via `nav-items.ts`'s role-list mechanism) — calls
  `useLogout()` (D4), then a client-side redirect to `/login`.
- **B. A user-menu/dropdown** (avatar or name, opening Profil/Keamanan/
  Keluar) replacing the current bare logo-link-only left side.
- **C. Add "Keluar" into `nav-items.ts`** as a fourth `NavItem`,
  reusing `FilteredNavLinks`/`NavLink` as-is.

**Recommendation: A.**

**Rationale:**
- Option C is rejected on a type-shape mismatch, not just style:
  `NavLink` renders a `<Link href>` navigation; logout is a mutation
  with a side effect (revoke token, clear store), not a navigation
  target — forcing it through the nav-item/`<Link>` shape would need
  either a fake `href` or a special-cased branch inside `NavLink`
  itself, complicating a component whose whole job today is "render a
  role-filtered link, nothing else."
- Option B is a real, reasonable direction longer-term, but is scope
  beyond this task's endpoint list (spec 03 covers session
  endpoints, not a Dashboard Shell IA redesign) — flag as a natural
  follow-up once more dashboard pages/actions exist to justify a menu,
  not build it speculatively now (lowest-complexity principle, same
  reasoning as D4's "no speculative `refresh()` wrapper").
- Gating on `useAccountMe()` data (not a role array) is the correct
  primitive here — logout should be visible to *any* authenticated
  user regardless of role, which is exactly what "is this query's data
  populated" already answers, without inventing a role that means
  "logged in at all."
- Flag for the eventual techplan/PR: `page-map.md`'s Cross-Cutting UI
  Elements table never lists a logout affordance anywhere (Stage 2
  Area 4 finding) — worth suggesting as a doc addition once this
  ships, same spirit as task 02's solutioning suggesting an
  `openapi.yaml` prose fix it found along the way (a suggestion for
  the doc's owner, not something this task unilaterally rewrites).

---

## D6 — Session-expiry-while-in-dashboard redirect

**Question:** Stage 2 Area 4 found no mechanism redirects a user whose
background refresh fails while they're already inside `/dashboard/*`
(the reuse-detection-revokes-family case, or any other `401` from
`/auth/refresh`). Spec 03 says "client must force a full re-login" but
specifies no frontend mechanics. Is closing this in scope for this
task?

**Options considered:**
- **A. In scope.** Extend `tryRefreshOnce`'s failure path (or a thin
  listener alongside it) to distinguish "was authenticated, refresh
  just failed" from "was never authenticated, refresh failed" (the
  ordinary silent-guest-hydration case `AuthBootstrapProvider` already
  handles per R10) — only the former triggers a client-side redirect
  to `/login`. Concretely: check `useAuthStore.getState().accessToken
  !== null` *before* clearing it in the failure branch; if it *was*
  non-null, this is a real session-loss transition, not an ordinary
  guest page load.
- **B. Out of scope** — leave as a pre-existing Phase-0-era gap for a
  later task to close, since no prior task's spec named it explicitly
  either.

**Recommendation: A.**

**Rationale:**
- This is a direct, unavoidable consequence of the same endpoints this
  task already owns (`/auth/refresh`'s reuse-detection behavior, which
  spec 03 itself describes and this task's own D3 work interacts with
  directly) — leaving a logged-in-looking-but-actually-logged-out
  dashboard state on screen is a worse UX outcome than the missing-
  logout-button gap in D5, since here the user has no way to tell
  their session is already gone until their next action mysteriously
  fails.
- The authenticated→unauthenticated transition check (option A's
  "was the token non-null before this failure") is the same shape of
  distinction `AuthBootstrapProvider`'s R10 already makes conceptually
  (never surface an error to a guest who was never logged in) — this
  extends that existing discipline to the opposite direction (a user
  who *was* logged in) rather than introducing a new one.
- Redirect target: `/login`, not a dedicated "session expired" page —
  no such page exists in `page-map.md` and none is needed; `/login`
  can optionally show a one-line explanatory note, but that's a small,
  separate copy decision for the techplan step, not a new-page
  decision.

---

## D7 — Error-message handling: 401 vs 429, banner vs field-level

**Question:** Stage 2 Areas 1 and 5 both confirmed: spec 03 mandates
identical generic detail text for `401` (wrong credentials/MFA code)
and `429` (lockout) — status differs, copy doesn't — and the design
reference's Known Issue #1 (field-level error on the email input) is
confirmed present and must not be copied.

**Recommendation (follows directly from confirmed spec text + the
already-built `AuthShellClient` convention — not a contested choice):**
- `LoginForm` (and its MFA-step branch) render `<Banner
  variant="error">{message}</Banner>` as the **first child**, exactly
  matching `AuthShellClient`'s documented convention and `RegisterForm`'s
  own precedent (`requestError && <Banner variant="error">...`) —
  never attached to the email/TOTP-code input's own `error` prop.
- Because the backend already sends the *same* generic detail string
  for both `401` and `429` per spec 03, the frontend does not need
  separate copy per status — `error instanceof ApiError &&
  error.detail ? error.detail : GENERIC_ERROR_MESSAGE` (same fallback
  pattern `RegisterForm` already uses for its own `429` case) covers
  both without a branch on `.status` at all. This is simpler than
  `RegisterForm`'s own handling (which does check `.status === 429`
  specifically) only because spec 03's copy is uniform across both
  codes where spec 01's wasn't — worth calling out so it doesn't look
  like an inconsistency with the established pattern when it's
  actually a deliberate simplification matching this endpoint's own
  spec.

---

## Open items for the techplan step

1. **D3's Web Locks deviation** from spec 03's literal `BroadcastChannel`
   wording — **RESOLVED, approved by Anhar (2026-08-26)**: Web Locks
   API for mutual exclusion + `BroadcastChannel` for token fan-out.
   Satisfies spec 03 Assumption D's stated goal via a different named
   mechanism than the doc's literal text — carry this decision (and
   its rationale in D3 above) into the techplan explicitly, since it
   deviates from a human-reviewed spec document.
2. **D2's refresh-loses-MFA-progress trade-off** — proceeding with the
   recommendation (component-local `useState`, lost on refresh) — no
   objection raised.
3. **D5's suggested `page-map.md` Cross-Cutting UI Elements addition**
   for a logout affordance — proceeding with the recommendation: flag
   as a doc suggestion in the eventual PR description, not edited as
   part of this task.
4. **D6's redirect-copy decision** — proceeding with the
   recommendation: plain redirect to `/login`, no special copy for v1.
