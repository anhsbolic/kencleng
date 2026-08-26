# Stage 3 — Solutioning: account/02-google-oauth-login-register (frontend)

> Builds on `stage2-gap-analysis.md`'s findings, ranked summary at the
> bottom of that doc. Each item below is a Decision Log entry:
> question → options considered → recommendation → rationale. This is
> raw solutioning material, not a techplan — synthesis into a techplan
> is a separate, later step.

Grounding note: task 01's own `stage3-solutioning.md` (D3) already
anticipated this exact confirmation step — its Decision Log explicitly
recorded "leave the OAuth callback [and by extension, the redirect
value's correctness] to Task #2" and flagged Active Open Item #3 for
this session to resolve. D1 below is that resolution.

---

## D1 — Fix the `intent=register` bug in `RegisterForm`

**Question:** Stage 2 Area 2 confirmed `register-form.tsx`'s "Daftar
dengan Google" link sends `intent=register`, which the merged backend
rejects with `400` (`validIntent()` only accepts `login`/`link`/
`reauth`). What's the correct value?

**Options considered:**
- **A. `intent=login`.** The backend's `login` intent already handles
  the "no existing identity, email not claimed elsewhere" branch by
  creating a new `User` + `AuthIdentity` (spec 02's acceptance table,
  row 2) — i.e. `login` intent *is* the registration path when no
  account exists yet. No frontend distinction between "log in" and
  "sign up" is needed at the OAuth-redirect level; the backend already
  collapses them.
- **B.** Add a fourth `register` value to the backend's `enum`/
  `validIntent()`. Rejected — out of this frontend task's authority
  (backend task #2 is merged; changing its accepted-value contract is
  a cross-track change needing its own review), and unnecessary per A.

**Recommendation: A.**

**Rationale:**
- Confirmed directly from the already-merged backend
  (`google_oauth.go`'s `callbackLogin`): `login` intent's three
  branches are (existing identity → login), (no identity, email free →
  **create User**), (no identity, email claimed by `email_password` →
  reject, no auto-merge). The middle branch is exactly what "Daftar
  dengan Google" needs — there's no missing backend capability, only a
  wrong query value.
- Also fixes `/login`'s own upcoming "Masuk dengan Google" button (D2)
  — same intent value, same fix, one shared component (D5) means this
  is authored once, not twice.
- Suggest also correcting `api/openapi.yaml`'s prose
  ("`intent=login`/`register` do not [require auth]") to remove the
  misleading `register` mention, since Area 2 traced the bug's likely
  origin to that exact sentence — small, low-risk doc fix alongside
  the code fix, prevents the same misreading recurring (e.g. in a
  future `link`-intent button). Flagged here as a suggestion for
  whoever owns `api/openapi.yaml`, not something this frontend task
  can authoritatively resolve alone since it's a shared cross-track doc.

---

## D2 — `/login` page scope: how much to build now

**Question:** `/login`'s credential form is explicitly task #3's scope
(per the placeholder's own comment), but page-map.md describes
`/login` as one unit ("Login form + Google button"). How much of
`/login` does task #2 own?

**Options considered:**
- **A. Build only what task #2 actually owns**: the "Masuk dengan
  Google" button/navigation, the `AuthShellClient`-first-child error
  banner wired to parse `?error={code}` from a failed callback
  redirect, and a placeholder note for the still-missing credential
  form (so the page doesn't look broken, just visibly "coming soon" —
  matching the honesty principle in root `AGENTS.md`, not silently
  pretending the page is finished).
- **B. Build the entire `/login` page**, including a stub email/
  password form, even though `POST /auth/login` doesn't exist as a
  callable endpoint yet in this frontend track.
- **C. Leave `/login` as the Phase 0 placeholder entirely**, defer all
  of it to task #3.

**Recommendation: A — directly mirrors task 01's own D3 precedent.**

**Rationale:**
- Same reasoning task 01 used for `/register`'s button (build what's
  actually your endpoint's scope now, leave what belongs to another
  task's endpoint for that task): task #2 owns
  `/auth/google/redirect`+`/auth/google/callback`, not
  `POST /auth/login`. Option B would mean writing UI against an
  endpoint this task has no contract for, risking exactly the kind of
  duplicated/conflicting work task #3 would then have to reconcile
  against — the opposite of the "one vertical slice per session"
  discipline `kencleng-agentic-workflow.md` §11 sets out.
- Option C repeats the mistake Option A of task 01's D3 rejected (its
  "Option A: defer the whole button to Task #2" — rejected as "leaves
  `/register` incomplete... for no technical reason"). The Google
  button and its error-handling banner are fully buildable today with
  zero dependency on task #3's login form.
- The placeholder note (Option A) keeps root `AGENTS.md`'s honesty
  principle intact — `/login` visibly and explicitly says "email/
  password login coming soon," rather than looking finished while
  silently missing a whole input method, or looking exactly like the
  current bare Phase-0 stub with no indication anything changed.
- **Explicit handoff note for whoever picks up task #3**: this task's
  build must not delete/replace the whole `login/page.tsz` file
  wholesale — task #3 extends it (adds the credential form above/
  alongside the Google button + divider, matching `RegisterForm`'s own
  "form fields → divider → Google button" composition), the same way
  task 01 built `RegisterForm` for task #2 to extend, not replace.

---

## D3 — Post-callback access-token hydration mechanism

**Question:** Stage 2 Areas 3/4/6 found: (a) the OAuth callback
delivers tokens as HttpOnly cookies, unreadable by JS; (b) `apiFetch`
only ever reads `useAuthStore`'s in-memory token; (c) a successful
login `302`s to bare `/` with **no query signal** to detect "just
logged in" vs. "ordinary visit." How does the SPA's in-memory access
token actually get populated?

**Options considered:**
- **A. Route-specific detection**: some mechanism only on `/` (or a
  dedicated route) that tries to tell "did I just arrive from OAuth."
  Rejected outright — Area 6 already established there is no signal to
  detect this by; any such mechanism would have to fall back to
  "always try" anyway, making the "route-specific" framing pointless.
- **B. Unconditional silent-hydration-on-app-boot**: a small root-level
  provider (alongside the existing `MockingProvider`/`QueryProvider` in
  `app/layout.tsx`) that, once per app load, if `useAuthStore`'s
  `accessToken` is still `null`, calls the refresh endpoint
  (`POST /auth/refresh`, which reads whatever refresh cookie is
  present) and populates the store on success, no-ops silently on
  failure (genuinely logged-out visitor — this must never surface an
  error to a guest just browsing `/`).
- **C. Do nothing new** — rely on `apiFetch`'s existing reactive
  401→`tryRefreshOnce()`→retry path to eventually pick up the cookie
  the first time any authenticated call 401s.

**Recommendation: B.**

**Rationale:**
- This isn't actually OAuth-specific scope creep: because the access
  token is deliberately **in-memory-only** (`auth-store.ts`'s own
  comment — no `persist` middleware, by design, to avoid a
  `localStorage`-readable token), **any** page refresh (F5) already
  wipes it for *any* future session, OAuth or otherwise. No prior task
  has needed a bootstrap-hydration mechanism yet only because no
  feature has shipped a working login before this one — task #2 is
  the first to actually establish a real session, making it the
  natural place this general mechanism gets built, not a special case
  bolted onto OAuth. `auth-store.ts`'s comment ("Tasks #2 and #3 share
  one store shape") supports this: whatever gets built for hydration
  here is exactly what task #3 needs too, for the exact same F5
  problem on its own login flow.
- Solves Area 6's "no query signal" problem entirely by making the
  signal irrelevant — the provider doesn't need to know *why*
  `accessToken` is `null`, only that it is, and unconditionally
  attempting a refresh is cheap and safe for a genuinely logged-out
  guest (the existing `tryRefreshOnce()` already no-ops cleanly on
  failure).
- Option C is rejected because it only ever fires reactively, after
  something has already attempted an authenticated call and gotten a
  401 — meaning the user's very first render after a successful Google
  login (e.g. a nav bar checking "am I logged in") would show a
  logged-out state until *something* else happens to trigger a call,
  producing a visible flash/flicker of the wrong UI state rather than
  a clean hydration-before-paint (as clean as SSR allows) sequence.
- Location: a new `components/providers/auth-bootstrap-provider.tsx`,
  same shape/convention as `QueryProvider` (`"use client"`, wraps
  `children`, one focused responsibility), added to `app/layout.tsx`'s
  existing provider stack. Reuses `tryRefreshOnce()`'s logic — needs
  that function (or an equivalent) actually exported from `client.ts`
  (currently module-private, per Area 4's finding), addressed by D6.

---

## D4 — Error banner copy for `/login?error={code}`

**Question:** Area 6 confirmed the literal contract: failed `login`-
intent callbacks land on `/login?error={code}` with `code` drawn from
a fixed backend vocabulary (`state_mismatch`, `nonce_mismatch`,
`google_token_invalid`, `google_unavailable`, `google_email_conflict`
— `google_link_conflict` is link-intent-only, out of scope here). What
does `/login` show for each?

**Options considered:**
- **A. One generic message for every code** — simplest, avoids
  authoring five copy variants without product sign-off.
- **B. A small code→copy mapping**, rendered via
  `<Banner variant="error">` in `AuthShellClient`'s documented
  first-child slot, with placeholder (not-yet-signed-off) Indonesian
  text per code — same "TBD, pending product copy" treatment already
  used twice elsewhere in this codebase (`RegisterForm`'s
  `GENERIC_ERROR_MESSAGE`, `VerifyEmailStatus`'s
  `INVALID_LINK_MESSAGE`).

**Recommendation: B.**

**Rationale:**
- `google_email_conflict` specifically needs distinguishable copy on
  its own merits, not just for polish: it's the no-auto-merge case —
  the user did nothing wrong (their Google email happens to match an
  existing password-based account), and a fully generic "something
  went wrong, try again" message would actively mislead them into
  retrying the exact same failing action. Spec 02 treats this as the
  most security-significant branch of the whole feature (top-severity
  anti-takeover threat per the threat breakdown) — the frontend
  shouldn't flatten away the one piece of information that actually
  helps a legitimate user recover (e.g. "use your existing password to
  log in instead").
- The other four codes (`state_mismatch`, `nonce_mismatch`,
  `google_token_invalid`, `google_unavailable`) are all genuinely
  request-level failures the user can't meaningfully act on
  differently — reasonable to collapse those four into one shared
  "coba lagi" fallback string while keeping `google_email_conflict`
  distinct, rather than authoring five fully bespoke strings. (Final
  grouping is an implementation nuance, not re-litigated further here.)
- Matches this codebase's own established precedent (D5 in task 01's
  solutioning did the analogous thing for `/verify-email`'s `410` vs
  `404`: keep backend-distinguished outcomes distinguishable in the
  UI rather than flattening away actionable information for no
  corresponding security benefit) — these error codes aren't
  enumeration-sensitive (unlike a guessable resource ID), so nothing
  is leaked by showing different copy per code.
- An unrecognized/future error code (defensive: the vocabulary could
  grow) should fall back to the same generic string as the four
  collapsed codes, never render the raw `code` value itself to the
  user (consistent with `patterns.md` §B's "never render raw backend
  error text" rule — the `code` is a stable enum for the frontend to
  map, not user-facing copy on its own).

---

## D5 — Shared Google-button component vs. duplicated markup

**Question:** `/register`'s Google button currently lives inline
inside `RegisterForm` (Area 2). `/login` needs the same markup
(anchor, styling, divider) with a different `intent` value and label.
Extract a shared component, or duplicate?

**Options considered:**
- **A. Duplicate** — copy the same `<a href>` block into `/login`'s
  new component, same as `RegisterForm`'s existing inline version.
- **B. Extract a shared `<GoogleAuthButton intent="login" | "register"
  />`** (or similarly named) in `components/features/account/`, used
  by both `RegisterForm` and the new `/login` component; the "atau"
  divider can be part of it or left to each caller (implementation
  detail either way).

**Recommendation: B.**

**Rationale:**
- Directly motivated by D1's bug: the exact failure mode just found
  (`intent=register` typo'd into `/register`'s copy) is precisely what
  independent hand-copies of "the same anchor with a different query
  value" produce. A shared component makes the `intent` value a typed
  prop (`"login" | "register"`, or reuse the backend's own accepted
  literal union if a shared type is available) instead of a
  free-floating string baked into markup in two places — the same
  class of mistake becomes a type error instead of a silent runtime
  400.
- Cheap to do: the only markup involved is one anchor tag's `href` and
  label text — no meaningful behavior divergence between the two call
  sites to abstract away, so there's no premature-abstraction risk
  here (unlike, say, extracting a shared form-wide component too
  early).
- Both call sites already need the identical navigation contract (R7
  from task 01's techplan — real navigation, never `apiFetch`) — a
  shared component is also the natural place to enforce that
  constraint once, structurally, rather than trusting every future
  caller to remember it.

---

## D6 — Exporting the refresh mechanism from `client.ts`

**Question:** D3's bootstrap provider needs to call the refresh logic
currently implemented as `tryRefreshOnce()`, a module-private
`async function` in `client.ts` not in its `export { ... }` list. How
does the provider reach it?

**Options considered:**
- **A. Export `tryRefreshOnce` as-is**, call it directly from the new
  provider.
- **B. Add a thin, separately-named public wrapper** (e.g.
  `hydrateSession()`) in `client.ts` that calls the same internal
  logic, keeping `tryRefreshOnce`'s name/semantics scoped to its
  original 401-retry purpose so the two call sites don't read as the
  same concept with two different names.
- **C. Bypass `client.ts` — call `POST /auth/refresh` directly** from
  the new provider, duplicating the fetch logic.

**Recommendation: A, with the adjacent cleanup from D9 folded in.**

**Rationale:**
- `tryRefreshOnce()`'s existing behavior (single attempt, sets the
  store on success, clears it on failure, never throws) is *exactly*
  the contract the bootstrap provider needs — no adaptation required.
  Option B adds a second name for the same behavior for no functional
  gain — just re-exporting the existing function is the smaller diff
  and keeps one source of truth.
- Option C reintroduces a second, parallel implementation of "call
  refresh, parse `{access_token}`, update the store" — directly the
  kind of duplication `api-client-centralization.md`'s core rule
  exists to prevent ("every domain's fetch function must go through
  `apiFetch`" — and by the same logic, every *auth-bootstrap* concern
  should go through the one place that already owns token state).
- While touching this function's export surface, also fix Area 5's
  finding: replace the hand-written `type RefreshResponse = {
  access_token: string }` with the generated
  `components["schemas"]["RefreshResponse"]` from `schema.d.ts` (D9,
  folded in here since it's the same touched code, not a separate
  task-worthy change on its own).

---

## D7 — Testing scope, given task #3's backend incompleteness

**Question:** Area 3 found that access tokens minted by this feature's
OAuth flow are, by the newer `VerifyAccessToken` (task #3, in
progress) design, rejected outright (missing `purpose` claim) — so a
fully "real" end-to-end verification (Google login → authenticated
`/account/me` call succeeding against the live backend) may not be
achievable today regardless of what the frontend does correctly. What
does this task's testing actually cover?

**Options considered:**
- **A. Test entirely against MSW mocks** (already this codebase's
  standard per `component-test-mocking-discipline`), and explicitly
  record the cross-task backend gap as an accepted, out-of-scope risk
  for this task rather than something the frontend can fix or must
  block on.
- **B. Block this task on backend task #3 resolving the `purpose`-
  claim gap first.** Rejected — task #3 is a separate, larger, still-
  in-progress serial-group task (`tasks.md` group S1); this frontend
  task's own scope (spec 02's two endpoints) doesn't depend on it, and
  blocking would stall clearly-buildable, independently-valuable work
  behind an unrelated task's completion.

**Recommendation: A.**

**Rationale:**
- Matches how this codebase already tests everything else
  (`register-form.test.tsx`, `verify-email-status.test.tsx` — MSW
  handlers, not a live backend). Nothing about this task's frontend
  code is what would cause a live-backend mismatch — it's a backend-
  internal token-format compatibility question between two backend
  tasks, invisible to and unaffected by the frontend's correctness.
- Per `docs/spec/README.md` §6 rule 4 ("ambiguity is recorded, not
  silently resolved"), this should be written up explicitly as an
  Assumption/residual-risk note in whatever techplan synthesizes from
  this exploration — not silently assumed away, and not treated as
  this task's bug to fix.

---

## Summary composition (ties D1–D6 together)

Not a new decision — the shape the above converges on, for a later
techplan-synthesis pass to start from:

- **Fix**: `register-form.tsx`'s Google anchor's `intent` value,
  `register` → `login` (D1), via the new shared component (D5).
- **New component**: `components/features/account/google-auth-button.tsx`
  (or similar name) — `intent: "login" | "register"` prop, real `<a
  href>` navigation (R7 convention carried forward), used by both
  `RegisterForm` (replacing its inline version) and `/login`'s new
  content (D5).
- **`/login` page**: replace the Phase-0 placeholder with: heading,
  an explicit "email/password login coming soon" placeholder note
  (task #3's scope), the shared Google button (D5, D1), and an error
  banner reading `?error={code}` via `useSearchParams()` (Suspense-
  wrapped, same pattern as `VerifyEmailStatus`) rendered as
  `AuthShellClient`'s documented first child, mapped per D4 (D2).
- **New provider**: `components/providers/auth-bootstrap-provider.tsx`
  — calls the (now-exported) refresh mechanism once per app load if
  `accessToken` is `null`, silently no-ops on failure, added to
  `app/layout.tsx`'s existing provider stack (D3).
- **`client.ts` changes**: export `tryRefreshOnce` (or equivalent);
  replace its hand-written `RefreshResponse` type with the generated
  schema type (D6/D9).
- **`mocks/handlers.ts`**: add a `POST /auth/refresh` handler to
  support testing D3's bootstrap provider (Area 5's finding).
- **Not built by this task**: `/login`'s credential form (task #3),
  anything under `/account/security`/`/dashboard/security` for link/
  reauth outcomes (tasks 05/06 — flagged mismatch left for whoever
  picks those up), and no attempt to make live end-to-end Google login
  succeed against the current in-progress backend (D7 — accepted,
  documented risk, not this task's fix).

## Flagged for escalation, not decided by this task

- **`/account/security` vs. `/dashboard/security`** (Area 6): a real,
  already-shipped mismatch between backend task #2's hardcoded
  redirect constant and `page-map.md`'s actual route. Not fixable
  from the frontend track alone (the backend constant is what's
  wrong, or page-map.md needs to adopt `/account/security` — a call
  for whoever owns tasks 05/06 or the backend constant, not this
  task).
- **`api/openapi.yaml`'s misleading `intent=login`/`register` prose**
  (D1): suggest a one-line doc fix alongside code changes, but this is
  a shared cross-track doc this task doesn't have sole authority to
  edit unilaterally.
