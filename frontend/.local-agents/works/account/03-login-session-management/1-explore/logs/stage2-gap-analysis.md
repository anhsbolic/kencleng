# Stage 2 — Gap Analysis
## Feature: account/03-login-session-management (frontend surface)

Five areas explored, each fully before moving to the next, per
`workflow/1-exploration/guidelines.md`. No solutions proposed here —
bare one-line observations only where they occurred naturally; full
solutioning is Stage 3.

---

## Area 1: Auth Shell + `/login` page

**Current state:**
- `app/(auth)/layout.tsx` — thin Server Component wrapping
  `AuthShellClient`. Built in Phase 0, not owned by any account-domain
  task.
- `app/(auth)/_components/auth-shell-client.tsx` — desktop: centered
  modal (`role="dialog"`, focus-trapped via `useFocusTrap`); mobile:
  plain full page via a CSS-only `md:` breakpoint switch (JS only
  gates focus-trap *activation*, not layout). Doc comment prescribes a
  load-bearing convention for whoever fills in `/login`: render
  `<Banner variant="error">` as the *first child*, before the form —
  this is the structural fix for `prototype-reference.md`'s Known
  Issue #1 (banner-not-field-error for the generic auth failure),
  built ahead of time but **not yet consumed by anything** — no page
  currently renders into that slot.
- `app/(auth)/login/page.tsx` — **real content exists, but explicitly
  scoped to task #2 only.** Renders `GoogleCallbackError` (inside a
  `<Suspense>`, matching the `useSearchParams()` requirement already
  established elsewhere) + a static "coming soon" banner + a single
  `GoogleAuthButton`. Own comment is explicit: *"The email/password
  credential form is a different task's scope (backend task #3,
  `03-login-session-management.md`) — deliberately not built here...
  Whoever picks up task #3 extends this page (adds the form
  above/alongside the Google button + divider, matching RegisterForm's
  own composition), not replaces it."*
- `app/(auth)/login/page.test.tsx` — asserts, by name, **no** email/
  password fields exist yet (`R4`) and only the Google link + heading +
  "coming soon" text render. This test will need to change as part of
  this task, not stay green untouched.
- `app/(auth)/forgot-password/page.tsx`, `reset-password/page.tsx` —
  both explicit Phase-0 placeholders, own comments say "real form is
  Account Task #4's scope" — confirmed **out of scope** for this task.
- `components/features/account/register-form.tsx` — the composition
  precedent named by `/login`'s own comment: `react-hook-form` +
  `zodResolver`, a `Banner variant="error"` for request-level failure
  rendered above the fields, per-field `error={errors.x?.message}` via
  `<Input>`, a `Button` with `loading={isSubmitting || mutation.isPending}`,
  an `aria-hidden` "atau" divider, then `GoogleAuthButton` below.
  Success path swaps the whole form out for a confirmation view with
  `tabIndex={-1}` + `.focus()` on mount (R17-style focus management).
  `GoogleAuthButton` itself is a real `<a href="/auth/google/redirect?
  intent=...">` navigation (never `apiFetch`), typed against the
  generated `paths[...]["intent"]` union so an invalid intent value is
  a compile error, not a runtime `400`.

**Requirement** (`page-map.md` + spec 03):
- `/login`: "Login form + 'Masuk dengan Google' button" (page-map.md
  §1, already partially satisfied by task #2).
- Spec 03 covers `POST /auth/login` (credentials → tokens or MFA-
  pending), `POST /auth/login/mfa` (TOTP/backup-code → tokens), plus
  the account-wide lockout (`429`) and generic-error (`401`, identical
  copy for wrong-email/wrong-password) rules.
- `patterns.md` Pattern 3 (Form Page): Idle → Validating (inline,
  on-blur+submit) → Submitting (button disabled + spinner, rest of
  form disabled) → Submit error (banner for request-level, field-level
  for 422 — never conflated) → Success (toast/redirect for dashboard
  forms).

**Gap:**
- No login form exists at all yet — no `login-form.tsx`, no
  `login-schema.ts` in `components/features/account/`.
- No UI surface anywhere for the MFA challenge step
  (`POST /auth/login/mfa`) — `page-map.md` has no route for it; it is
  not `/dashboard/security` (that page is MFA *enrollment*, a
  different feature — see Area 5). This has to be either a second
  client-state "step" inside the same `/login` page/component, or some
  other non-page surface — not yet resolved which (Stage 3 concern),
  but confirmed as a genuine gap with no existing scaffold.
- The `429` lockout response and the `401` generic-credentials response
  share "the same generic detail text, only the status code differs"
  per spec 03 — current `ApiError` class already carries `.status` and
  `.detail` (built by task #1/#2), so the shape needed to distinguish
  them is already available; nothing in `components/features/account/`
  consumes it that way yet.
- `login/page.test.tsx`'s existing assertion (`queryByLabelText(/email/i)`
  etc. must NOT be in the document) is a **negative** test that this
  task will directly contradict — it needs to be replaced, not left
  green, or the suite will fail once fields are added.

**Page-consolidation check:**
- `/login` already exists as a route, partially built by task #2 —
  this task **extends** the existing page/file, it does not create a
  new route. Confirmed via the page's own doc comment (quoted above),
  which is explicit about this being the intended split.
- `/dashboard/security` (MFA enrollment, task #6's frontend surface)
  is a **separate** page-map.md row from anything in this task — no
  page-map.md action here lacks a backing endpoint or vice versa; the
  MFA-challenge-at-login step simply has no dedicated *route* named in
  page-map.md at all (a non-page surface, matching the exploration
  brief's framing).
- Backend task #3's tracker note in `docs/spec/1-account/tasks.md`
  says "build not started," but the actual working tree
  (`backend/.local-agents/works/account/03-login-session-management/`)
  already has `3-build/report.md`, `4-code-review/`, `4-patch/`, and
  `5-testing/report.md` populated — i.e. the backend implementation
  appears to have gone through build, code review, a patch round, and
  testing already, despite the tracker text. This is the same kind of
  stale-tracker mismatch flagged during task #2's exploration (that
  one confirmed via git log; this one confirmed via directory
  contents, not yet cross-checked against actual endpoint behavior —
  worth verifying the real response shapes before relying on this in
  Stage 3, see Area 3).

**Sniffing:**
- *Misleading signal*: `AuthShellClient`'s banner-first convention and
  `login/page.tsx`'s Suspense-wrapped `GoogleCallbackError` both look
  like "the error-handling story for `/login` is basically done" — but
  neither has ever rendered a credential-failure banner, since no
  credential form exists yet to trigger one. The convention is
  correctly *positioned* but genuinely untested end-to-end.
- *Miscontext*: `page-map.md` describes `/login` as one atomic row
  ("Login form + Google button") with no acknowledgment that the two
  halves ship on different tasks — same miscontext already flagged
  during task #2's exploration, still unresolved in the doc itself.
- *Risk*: this task overwrites a currently-passing, currently-green
  page and its test file with materially different behavior — low
  blast radius (only one page), but the "coming soon" copy and the
  negative test are both facts today that will read as regressions in
  a diff unless the PR description is explicit that removing them is
  intentional, not a revert.
- *Edge case*: the `mfa_pending_token` (5-minute TTL) is the frontend's
  only carrier of "which user already passed the password step" —
  needs to be held in component state (not `localStorage`, not the
  Zustand store per its own no-persistence rationale) between the two
  submit steps; a page refresh between password-submit and MFA-submit
  would lose it, forcing the user back to the password step. No
  existing code establishes where this kind of short-lived,
  cross-submit value should live — first time this shape has come up
  in the codebase.
- *Inconsistency*: none new found in this area beyond the already-
  tracked `page-map.md` atomic-row miscontext above.

---

## Area 2: Session/token infrastructure

**Current state:**
- `lib/stores/auth-store.ts` — Zustand store, **shape only**: `{
  accessToken: string | null, setAccessToken, clearAccessToken }`.
  Deliberately not persisted (`persist` middleware not used) — doc
  comment: token is short-lived and re-obtainable via
  `POST /auth/refresh`, the `HttpOnly` cookie is the real durable
  credential, keeping the access token in-memory-only avoids
  `localStorage` exposure. Own comment states explicitly: "no login
  logic lives here (that's Account Task #3's job)."
- `lib/api/client.ts` — already fully wires the generic
  refresh-on-401 mechanics:
  - `tryRefreshOnce()`: calls `POST /auth/refresh` directly (bypassing
    `apiFetch`, to avoid recursing into its own 401 handler), sets the
    new access token into the store on success, **clears** it on any
    failure (network error or non-OK response) so the app falls back
    to logged-out rather than hanging.
  - `apiFetch()`: attaches `Authorization: Bearer <token>` when
    present, `credentials: 'include'` always, and
    `X-Requested-With: kencleng-frontend` on mutating methods (the
    CSRF-adjacent custom header). On a `401` that isn't already a
    retry, calls `tryRefreshOnce()` once and retries the original call
    exactly once (`isRetry` guard) — a `401` on the retried call is
    returned as-is, never refreshed again.
  - `tryRefreshOnce` is exported specifically so
    `AuthBootstrapProvider` can reuse the *exact same logic* on app
    mount, to hydrate the store from whatever refresh cookie already
    exists (including one set by a successful Google OAuth callback,
    which delivers tokens as cookies rather than a JSON body).
- **Not yet read this session**: `components/providers/
  auth-bootstrap-provider.tsx` itself (only referenced by `client.ts`'s
  comment so far) — deferred to confirm exactly what it does on mount
  and whether it already covers the multi-tab case.

**Requirement:**
- Spec 03's `POST /auth/refresh` acceptance criteria: atomic
  rotate-on-use (`replaced_by_id` guarded UPDATE), reuse detection
  revokes the whole `family_id`, concurrent same-token requests
  produce exactly one winner and treat the loser identically to reuse
  (accepted trade-off, not a bug — Assumption D).
- Spec 03's Assumption D, verbatim: *"the multi-tab race is a
  frontend-track concern, to be solved with cross-tab coordination
  (`BroadcastChannel` — one tab acts as the single 'refresher', others
  wait and receive the rotated access token via the channel instead of
  independently calling `/auth/refresh`), not a backend change... This
  is deferred to when the `account` domain's frontend track starts —
  noted here now so it isn't rediscovered from scratch later."* This
  task **is** that deferred point.
- `POST /auth/logout`: idempotent, `204` whether or not a cookie was
  present — needs a corresponding store-clearing action on the
  frontend (clear `accessToken`) regardless of the response.

**Gap:**
- The `BroadcastChannel` cross-tab coordination named explicitly by
  spec 03 Assumption D **does not exist anywhere in the codebase** —
  confirmed by the file list; no `BroadcastChannel` reference in
  `lib/` or `components/providers/`. Today, each tab's own
  `apiFetch`/`tryRefreshOnce` would independently race
  `POST /auth/refresh` against every other open tab, and per the
  backend's own accepted-trade-off design (spec 03, `POST /auth/refresh`
  4th bullet), the *losing* tab's call gets treated as reuse-detection
  and the **entire token family is revoked** — i.e. without this
  coordination, simply having two tabs open and both triggering a
  refresh around the same time force-logs-out the user, even with zero
  malicious activity. This is a real, spec-acknowledged gap, not a
  hypothetical edge case.
- No existing `login()`/`logout()` action exists anywhere that ties a
  successful `POST /auth/login` (or `/auth/login/mfa`) response's
  token into `useAuthStore`, or that calls `POST /auth/logout` and
  clears it. `auth-store.ts`'s own comment says this is explicitly out
  of its scope and belongs to this task.

**Page-consolidation check:** N/A — this area is infrastructure, not a
page-map.md row.

**Sniffing:**
- *Risk*: the multi-tab gap above is the highest-severity finding so
  far in this exploration — it's not a cosmetic UX gap, it's a
  functional defect (unwanted forced logout) that the backend spec
  explicitly anticipated and explicitly assigned to this exact task.
  Reach: every user who ever has Kencleng open in two tabs
  simultaneously past one access-token TTL (15 min) — plausible for
  any donor casually tab-browsing.
- *Edge case*: `tryRefreshOnce`'s current single-attempt design (no
  retry loop) is correct for the single-tab case but doesn't itself
  prevent the cross-tab race described above — the fix has to live at
  a layer above individual `apiFetch` calls (coordinating *which tab*
  is even allowed to call refresh), not inside `tryRefreshOnce` itself.
- *Misleading signal*: because `client.ts` already looks fully
  "production-shaped" (proper single-retry guard, proper store
  integration, thorough doc comments), it would be easy to assume the
  refresh story is complete for this task — it's complete for the
  *single-tab* case only; the spec's own Assumption D flags the
  multi-tab piece as unfinished business assigned here.
- *Inconsistency*: none found — `client.ts`'s behavior matches spec 03
  faithfully for everything it does cover.

---

## Area 3: API layer & generated types

**Current state (not yet fully read this session — file inventory
only so far):**
- `lib/api/account.ts`, `lib/api/schema.d.ts` (generated, per
  `frontend/AGENTS.md` §1 — "don't hand-edit the generated file"),
  `lib/hooks/use-account-me.ts`, `use-register.ts`,
  `use-resend-verification.ts`, `use-verify-email.ts` exist. No
  `use-login.ts`, `use-login-mfa.ts`, `use-refresh.ts`, or
  `use-logout.ts` hook file exists yet (confirmed via the full
  `lib/` file listing gathered in Stage 1).
- Whether `schema.d.ts` already reflects the backend's actual
  `LoginResponse` / `LoginMfaRequiredResponse` / `RefreshResponse`
  shapes (and whether it needs regenerating against a newer
  `api/openapi.yaml`, given the backend directory-contents finding in
  Area 1 suggesting task #3's backend work is further along than its
  tracker text claims) has **not been checked yet** — deferred to
  continue this area in the next pass, flagging now so it isn't
  forgotten.

**Requirement:** typed fetch functions + TanStack Query hooks for
`POST /auth/login`, `POST /auth/login/mfa`, `POST /auth/refresh` (already
partly covered by `client.ts`'s internal `tryRefreshOnce`, but no
public hook), `POST /auth/logout` — per `frontend/AGENTS.md` §1/§3
(types from `lib/api/`, generated, not hand-written; TanStack Query for
anything server-derived).

**Gap:** four endpoint integrations to add; exact shape of the gap
(whether `schema.d.ts` needs a regen step first) still open — **this
area's exploration is incomplete**, continuing next before moving to
Area 4.

**Sniffing:** deferred — insufficient information yet to run the five
lenses meaningfully on the parts not yet read.

---

## Area 3 (continued): API layer & generated types

**Current state (completed):**
- `lib/api/schema.d.ts` **already has the full contract** for all four
  endpoints — confirms the Area 1 suspicion that backend task #3 is
  further along than its `tasks.md` tracker text: `LoginRequest`
  (`email`+`password`), `LoginResponse` (`status: "ok"`, `access_token`,
  `access_token_expires_at?`, `user`), `LoginMfaRequiredResponse`
  (`status: "mfa_required"`, `mfa_pending_token`, still literally
  tagged `# INFERRED` in its own doc comment — i.e. the OpenAPI author
  proposed this shape rather than transcribing it from a pre-existing
  contract), `LoginMfaRequest` (`mfa_pending_token` +
  `totp_code?`/`backup_code?`, comment: "one of... must be present" —
  not enforced by the type itself, just documented), `RefreshResponse`
  (`access_token`, `access_token_expires_at?`), and `/auth/logout`
  (`204`, no body, no request). `POST /auth/login`'s `401`/`429`
  responses both carry `Problem`/`TooManyRequests` — same
  `readProblemDetail`-compatible shape already used elsewhere in
  `account.ts`. No regen needed — this is a genuine gap in
  *consuming* code, not in the generated types.
- `lib/api/account.ts` has **no functions yet** for any of the four
  endpoints — only `getMe`, `register`, `verifyEmail`,
  `resendVerification` exist. The existing `postAccountAction` helper
  (normalizes a thrown network error into `ApiError(0)`) is reusable
  as-is for `login`/`loginMfa`/`logout` — no new low-level plumbing
  needed there.
- `lib/hooks/` has **no hooks yet** for any of the four endpoints.
  `useRegister` (plain `useMutation`, no query-key/invalidation — "an
  unauthenticated user has nothing cached yet") is the closest
  precedent, but login/logout are different: on success they change
  *authenticated* state (`useAuthStore` + the `account.me` query cache
  that `AuthBootstrapProvider` already knows how to invalidate) — a
  successful login mutation has cache/store side effects `useRegister`
  never needed.
- `mocks/handlers.ts` has **no mock handlers** for `/auth/login`,
  `/auth/login/mfa`, or `/auth/logout` — only `/auth/refresh` exists,
  and it's hardcoded to always return `401` (models "no session," since
  `AuthBootstrapProvider` calls it unconditionally on every load; tests
  needing the success path already override it per-test via
  `server.use(...)`, per that endpoint's own comment). Per the file's
  own stated convention ("One handler per endpoint, added as the
  page/component that needs it gets built... not speculatively ahead
  of demonstrated need"), this is expected/normal, not a defect — but
  confirms all three need adding as part of this task, not already
  scaffolded anywhere.

**Requirement:** typed fetch functions + hooks for all four endpoints,
consistent with `frontend/AGENTS.md` §3 (TanStack Query for
server-derived state) and §1 (types from generated `lib/api/`).

**Gap:** four fetch functions (`login`, `loginMfa`, `refresh`†,
`logout`) and corresponding hooks, plus three new mock handlers. †
`refresh` already has a *working implementation* inside `client.ts`
(`tryRefreshOnce`) but no public `lib/api/account.ts`/hook wrapper —
open question for Stage 3 whether one is even needed, since nothing
outside `client.ts`/`AuthBootstrapProvider` calls it directly today.

**Page-consolidation check:** N/A — API/hook layer, not a page-map.md
row.

**Sniffing:**
- *Misleading signal*: `LoginMfaRequiredResponse`'s own doc comment
  admits it's `# INFERRED` — i.e. even the OpenAPI contract itself is
  not a transcription of settled, reviewed backend behavior but the
  spec-writing agent's own proposal. Given Area 1's finding that the
  backend has since gone through build/review/testing, this comment
  may now be stale/overly-hedged — worth a quick cross-check against
  the actual backend response-shaping code before treating the
  `LoginMfaRequiredResponse` shape as settled (Stage 3 concern, not
  resolved here).
- *Edge case*: `LoginMfaRequest`'s "one of `totp_code`/`backup_code`
  must be present" rule is documentation-only in the generated type
  (both fields are optional, nothing stops sending neither or both) —
  the frontend form must enforce "exactly one, non-empty" itself via
  its `zod` schema; the generated type alone won't catch a caller bug.
- *Risk*: none new beyond what's already flagged in Area 2 (the
  cross-tab race) — this area is pure additive scaffolding, low risk
  in isolation.
- *Inconsistency*: none found — the generated schema and `client.ts`'s
  existing `RefreshResponse` usage already agree with each other.

---

## Area 4: Dashboard Shell logout entry + session-expiry handling

**Current state:**
- `app/(dashboard)/_components/dashboard-shell-client.tsx` — header
  contains: logo/home link, desktop nav (`FilteredNavLinks`, role-
  gated via `useHasRole`), `NotificationBadge`, and a hamburger button
  (mobile only) that opens a drawer reusing the same filtered nav
  list. **No logout button, no user menu, no account-identity display
  of any kind exists anywhere in this component.**
- `app/(dashboard)/_components/nav-items.ts` — three items only
  (`Profil`, `Keamanan`, `Notifikasi`), all roles `["donatur",
  "kurator", "admin"]`. Own comment: "Starts small... other domains'
  items get added here by *those* domains' own tasks" — i.e. adding a
  logout affordance here (if that's where it belongs) is an expected,
  sanctioned kind of extension, not a structural change to the Shell.
- `lib/hooks/use-has-role.ts` — `useHasRole` returns `false` (safe
  default) whenever `useAccountMe()`'s `data` is falsy — covers
  "loading," "unauthenticated," and "failed to load" identically, by
  design (doc comment: avoids a nav item flashing visible before the
  real answer is known).
- `components/shared/require-role.tsx` — thin wrapper over
  `useHasRole`; **no page uses it yet** per its own comment (Serial
  group S1 has "no role-gated page *content*, only role-gated *nav
  items*").
- **No route guard exists anywhere for `/dashboard/*` as a whole** —
  confirmed by reading `(dashboard)/layout.tsx` (just renders
  `DashboardShellClient`, no auth check) and the shell component
  itself (only gates individual *nav links*, not page content, not the
  layout's own render). An unauthenticated visitor hitting e.g.
  `/dashboard/profile` directly today would see the Shell chrome with
  all nav items hidden, and whatever that page itself renders
  underneath — not evaluated further here since no dashboard page in
  this domain has real content yet (`profile`/`security` are both
  placeholders, per Area 1/5).

**Requirement:**
- Spec 03's `POST /auth/logout`: "idempotent... `204`... clears the
  cookie" — needs a UI trigger somewhere for a logged-in user, and a
  corresponding `useAuthStore.clearAccessToken()` call regardless of
  response (logout is defined as always succeeding from the client's
  perspective).
- `page-map.md` doesn't list a dedicated "logout" page/row (expected —
  it's an action, not a page) but every persona table implies a
  logged-in user needs a way to end their session; no explicit
  UI-location requirement is stated anywhere in `page-map.md`/
  `patterns.md` for *where* that action lives.
- Refresh-failure/reuse-detection (`401` from `POST /auth/refresh`,
  spec 03) implies a currently-authenticated user's session can become
  invalid *while already inside the dashboard* (family-revoked reuse
  case, or the Area 2 cross-tab race before it's fixed) — spec 03 says
  "client must force a full re-login" but doesn't specify the frontend
  mechanics of that (redirect target, whether an in-flight page's data
  is discarded).

**Gap:**
- No logout entry point exists anywhere in the codebase — confirmed
  gap, this task's responsibility per the endpoint list in spec 03.
- No session-expiry-while-in-dashboard redirect exists anywhere —
  `tryRefreshOnce`'s failure path only clears the store; nothing
  observes that and navigates the user anywhere. Whether "redirect to
  `/login` on a failed background refresh" is in this task's scope or
  a pre-existing Phase-0 gap this task should also close is a Stage 3
  scoping question, not resolved here.

**Page-consolidation check:** N/A — cross-cutting Shell behavior, not
a page-map.md row itself, though page-map.md's Cross-Cutting UI
Elements table (notification badge, `MaskedField`, etc.) is the
closest precedent for "shell-level element documented once, used
everywhere" — logout isn't listed there either, another small
page-map.md gap worth noting rather than silently filling in.

**Sniffing:**
- *Risk*: shipping a login flow with genuinely no way to log out is a
  real, visible product gap if this task treats "login" as complete
  without it — reach is every authenticated user, though the actual
  failure mode (stuck logged in, not a security hole — the refresh
  cookie still expires on its own TTL) is lower severity than the
  Area 2 cross-tab finding.
- *Miscontext*: `page-map.md` never explicitly assigns "where does
  logout live" to any task or page — an easy thing to miss entirely
  if scoping only from page-map.md's per-persona tables, since it's
  implied but never stated as its own row/action anywhere.
- *Edge case*: a logout button placed in `DashboardShellClient` would,
  by the file's existing pattern, presumably need the same
  role-list-based rendering as other nav items — but logout isn't
  role-gated, it's "any authenticated user," which is a slightly
  different condition than anything `nav-items.ts`'s current shape
  expresses (roles array, not "just needs to be logged in at all").
- *Inconsistency*: none found beyond the page-map.md omission above.

---

## Area 5: Visual precedent + MFA-enrollment boundary check

**Current state:**
- `design-reference/login-register.html` (Tier 1 per
  `prototype-reference.md`) extracted per `design-reference-usage.md`
  Step 1. Its `LoginForm` component (111-line JSX, the whole export is
  small) renders: an `Input` for email, a hand-rolled `PasswordField`
  with a show/hide `Eye`/`EyeOff` `IconButton` toggle and a "Lupa
  password?" link, a full-width `Button` ("Masuk"), a `Divider`
  ("atau"), then a `Button variant="outline"` with a Google "G" icon
  ("Masuk dengan Google"), then a "Belum punya akun? Daftar" link.
  Rendered twice per breakpoint: `DesktopOverlay` (blurred landing
  page behind a centered `Panel`, matching `AuthShellClient`'s real
  modal behavior already built) and `MobilePage` (full-page, back-
  arrow header) — both instantiated once in a default state and once
  with `err` set.
- **Confirmed, by direct inspection of the JSX (not just inferred from
  the doc), that Known Issue #1 from `prototype-reference.md` is
  present exactly as described and NOT fixed**: `err` state passes
  `error="Email atau password salah."` straight into the **email
  Input's own `error` prop** (line 50 of the extracted JSX:
  `<Input ... error={err ? "Email atau password salah." : undefined} />`)
  — a field-level error on one specific field, not a banner above the
  form. This must not be copied; the real implementation follows
  `AuthShellClient`'s already-built banner-first convention instead.
  `prototype-reference.md` had flagged this as "not confirmed fixed —
  verify before implementing" — now verified: not fixed.
- **The reference has zero MFA-step UI of any kind** — no second
  panel/state, no OTP input, no mention of `mfa_pending_token` or a
  two-step flow anywhere in the 111-line file. This isn't a case of
  "the reference has it but it's wrong" (like Known Issue #1) — the
  reference simply never modeled this part of the flow at all. Matches
  Area 1's finding: this is a genuine gap with **no visual precedent
  whatsoever**, tier system or otherwise — not even a Tier 2
  "closest precedent" mapping exists for it in `prototype-reference.md`.
- `app/(dashboard)/dashboard/security/page.tsx` — still the literal
  Phase-0 placeholder ("Placeholder — Account Task #5"). No MFA
  enrollment UI, no linked-identity UI, nothing has been built there.
  Note: its own comment says "Account Task #5" specifically, but
  `page-map.md` describes `/dashboard/security` as covering **both**
  task #5 (account linking / set-password) **and** task #6 (MFA
  TOTP) content on one page ("Enable/disable MFA... link/unlink Google
  identity... Atur Password") — the placeholder's comment names only
  one of the two owning tasks. Not a live conflict today (nothing is
  built yet either way), but worth flagging so whoever picks up task
  #5 or #6 doesn't read the placeholder comment as the complete
  picture of what belongs on that page.

**Requirement:** `prototype-reference.md` Tier 1 treats
`design-reference/login-register.html` as "close to authoritative on
layout/visual detail... deviating only where the feature spec requires
different behavior than what was mocked." Spec 03 requires an MFA
challenge step the mock never modeled at all.

**Gap:** the visual precedent covers the password-only login form
fields/composition well (directly reusable structural precedent —
email/password fields, show/hide toggle, "Masuk"/divider/Google
button/register-link composition, matching what `RegisterForm` already
does in real code), but provides **no precedent whatsoever** for the
MFA-challenge step this task also needs to build. That step has to be
designed from `patterns.md`'s generic Form-page state table alone, with
no Tier 1 or Tier 2 mapping to lean on.

**Page-consolidation check:** Confirms Area 1's finding — no overlap
between this task's login-time MFA challenge and `/dashboard/security`'s
(separate tasks' #5/#6) MFA *enrollment* surface; both are currently
unbuilt, so there's no existing code to reconcile, only a boundary to
keep clean going forward.

**Sniffing:**
- *Misleading signal*: `prototype-reference.md`'s Tier 1 table lists
  `/login` with the note "Desktop = modal overlay, mobile = full page —
  see known issue below," which could read as "this page is basically
  fully speced visually" — but the MFA-challenge half of this task's
  scope has no visual spec at all, a much bigger gap than the one
  documented known issue.
- *Risk*: low for the parts with precedent (email/password fields);
  higher for the MFA step purely from a "first-of-its-kind, no
  precedent to check against" standpoint — same category of risk as
  Area 1's `mfa_pending_token` state-storage finding, not a new risk.
- *Edge case*: none new beyond what Areas 1 and 3 already surfaced.
- *Inconsistency*: none found — `prototype-reference.md`'s own
  "Status: not confirmed fixed" hedge on Known Issue #1 turned out to
  be accurate (it isn't fixed), so no doc/code mismatch here, just a
  confirmed-as-predicted issue.

---

## Cross-area summary (no new analysis — index only)

1. **Auth Shell + `/login`** — extend existing page/file; negative
   test needs replacing; no MFA-step UI surface exists anywhere.
2. **Session/token infra** — single-tab refresh is solid; **cross-tab
   `BroadcastChannel` coordination is completely missing**, and its
   absence causes a real forced-logout bug per the backend's own
   accepted trade-off design (highest-severity finding overall).
3. **API layer & types** — generated types are already complete and
   correct; zero fetch functions, hooks, or mocks exist yet for any of
   the four endpoints.
4. **Dashboard Shell logout + expiry** — no logout entry point
   anywhere; no session-expiry-while-in-dashboard redirect anywhere;
   `page-map.md` itself never assigns logout a home.
5. **Visual precedent + MFA boundary** — Known Issue #1 (field-level
   vs. banner error) confirmed present, not fixed, must not be copied;
   login fields have solid visual precedent, MFA-challenge step has
   none at all; no live conflict with the separate MFA-enrollment
   page (`/dashboard/security`), which remains unbuilt.
