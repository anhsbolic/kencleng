# Stage 2 — Gap Analysis
## Feature: account/02-google-oauth-login-register (frontend surface)

Six areas explored, each fully before moving to the next, per
`workflow/1-exploration/guidelines.md`. No solutions proposed here —
bare one-line observations only where they occurred naturally; full
solutioning is Stage 3.

---

## Area 1: Auth Shell + `/login` + `/register` pages

**Current state:**
- `app/(auth)/layout.tsx` — thin Server Component wrapping
  `AuthShellClient`. Built in Phase 0 (`phase0-shared-infra.md` Step
  3), not owned by any account-domain task.
- `app/(auth)/_components/auth-shell-client.tsx` — desktop: centered
  modal (`role="dialog"`, focus-trapped via `useFocusTrap`); mobile:
  plain full page via CSS-only `md:` breakpoint. Doc comment
  prescribes a load-bearing convention for whichever task fills in
  `/login`: render `<Banner variant="error">` as the *first child*,
  before the form — the structural fix for the known
  field-level-vs-banner login-error bug
  (`prototype-reference.md` issue #1). Not consumed by anything yet.
- `app/(auth)/login/page.tsx` — **still the literal Phase 0
  placeholder.** Own comment: *"Placeholder — real form is Account
  Task #3's scope (`docs/spec/1-account/tasks.md`), not this
  playbook."* Static heading + "Placeholder" text only. No form, no
  Google button, no banner slot in use.
- `app/(auth)/register/page.tsx` — real (task 01): renders
  `<RegisterForm />`. Page-level file itself has no Google button.

**Requirement** (page-map.md + spec 02):
- `/login`: "Login form + 'Masuk dengan Google' button."
- `/register`: "Register form + 'Daftar dengan Google' button."

**Gap:**
- `/register`: needs a "Daftar dengan Google" CTA — see Area 2, one
  already exists inside `RegisterForm`, just with a bug.
- `/login`: needs a Google button, but there's no login form to
  attach it to. The credential form is explicitly out of scope per
  the placeholder's own comment, which names a *different* backend
  task (#3, `03-login-session-management.md`).

**Page-consolidation check:**
- `/register` already exists from task 01 — this task extends it.
- `/login` exists only as a Phase 0 placeholder — no account-domain
  task has touched it yet. `docs/spec/1-account/tasks.md`: **backend
  task #2 (this exact feature) is `merged`**
  (`efc1111`→`ce61841`); **backend task #3 is `in progress`, "build
  not started" per the tracker note** — but `git status` shows a
  large set of uncommitted backend files for task #3 already present
  (`login.go`, `auth_login.go`, `mfa_verifier.go`, `token.go`,
  migrations 000006–000009), so that tracker note looks stale against
  actual working-tree state.
- `.local-agents/works/account/` frontend-track directories mirror
  the **backend** feature-spec numbering 1:1 — no separate frontend
  task list exists. Supports scoping this task narrowly to "the
  frontend surface of exactly feature-spec 02" (Google OAuth only).
- No page-map.md action here lacks a backing endpoint, and vice versa
  — both buttons map directly to spec 02's two endpoints.

**Sniffing:**
- *Miscontext*: page-map.md treats `/login` as one atomic unit ("Login
  form + Google button") with no acknowledgment that the two halves
  ship on different tasks per `tasks.md`'s own split.
- *Inconsistency*: page-map.md's "Shell & Benchmark Notes" claims a
  "Google OAuth full-redirect flow" is defined in `patterns.md` —
  grepped, it isn't (only generic Success-state redirect language at
  line 98, unrelated line 189). Flagged in Stage 1, still unresolved.
- *Misleading signal*: `AuthShellClient`'s doc comment already
  prescribes the banner-first convention — easy to assume "the login
  error state is handled" because the shell scaffolding exists, but
  nothing renders into that slot yet, including for this task's own
  callback-error case.
- *Edge case*: spec 02's error table sends several failure branches to
  `/login`/security page via `302 ... ?error={code}` — the param name
  is confirmed as literally `error` once the actual backend code was
  read (Area 6), resolving what was an open question in the spec doc
  itself.
- *Risk*: `/login` currently has zero real functionality — whatever
  this task builds there is the first real content on that route, so
  any mistake is maximally visible, unlike `/register`'s incremental
  addition.

---

## Area 2: `components/features/account/`

**Current state:**
- Only `register-form.tsx`, `register-schema.ts`,
  `resend-verification-control.tsx`, `verify-email-status.tsx` (+
  tests) exist. **No `login-form.tsx`.**
- **`RegisterForm` already contains a "Daftar dengan Google" entry
  point** (lines 158–166): a real `<a href="/auth/google/redirect?
  intent=register">` (not `apiFetch`), matching documented rule R7
  from task 01's own techplan.
- This was deliberate, tracked, pre-built scope-splitting by task 01:
  its `task-02-register-page.md` records **Decision D3** — "build the
  button + navigation now [task 01]; leave the OAuth callback to the
  domain's Task #2" — and explicitly flags **Active Open Item #3**:
  *"confirm the Google button/navigation scope split (D3) against
  `tasks.md` with whoever owns the domain's own Task #2, before
  merging."* This exploration is that confirmation step.

**Requirement:** spec 02's `GET /auth/google/redirect` accepts
`intent` ∈ `{login, link, reauth}` only (OpenAPI `enum:`,
`api/openapi.yaml` line ~339-346).

**Gap / confirmed bug:** the already-built button sends
**`intent=register`** — not a valid value. Verified against the
actual merged backend, not just the spec doc:
- `backend/internal/domain/account/google_oauth.go:24-26` — only
  `intentLogin="login"`, `intentLink="link"`, `intentReauth="reauth"`
  defined; `validIntent()` (line 117-120) checks exactly these three.
- `backend/internal/transport/http/auth_google.go:173-176` — any
  other intent value gets an immediate `400 "Invalid Intent... must be
  one of: login, link, reauth."`
- So today, clicking "Daftar dengan Google" would hit a live `400`
  the moment it's exercised against the real (already-merged) backend.

**Miscontext (root cause):** `api/openapi.yaml`'s own prose at the top
of `/auth/google/redirect` reads *"`intent=login`/`register` do not
[require auth]"* — casual language conflating "the outcome for a new
user is a registration" with "there is an `intent=register` query
value," two lines above a formal `enum:` list that only allows
`login`/`link`/`reauth`. Most plausible origin of the bug: whoever
wrote `RegisterForm` (or its techplan) followed the loose prose
instead of the enum.

**Page-consolidation check:** task 02's frontend surface is *not*
starting from zero on `/register` — starting from a component with
the right navigation *mechanism* (real link, not fetch) but the wrong
query *value*. `/login` has no equivalent yet — task 02 owns building
it from scratch there, plus fixing `/register`'s existing one.

**Sniffing:**
- *Risk*: shallow but real — one string-value fix (`register`→`login`)
  resolves both instances once found, but passes a purely visual
  review; only caught by tracing the query param into the backend's
  actual validation.
- *Misleading signal*: the component's own inline comment ("R7 — a
  real navigation... not an apiFetch/XHR call") correctly nails the
  harder mistake (navigation vs. fetch) while missing the easier one
  (query value), so a reviewer skimming the comment for reassurance
  could miss the actual bug.
- *Inconsistency*: `api/openapi.yaml`'s own prose vs. its own `enum:`
  — worth flagging at the API-doc level too, not just the frontend
  call site, since the same misreading could recur (e.g. a future
  `link`-intent button).
- *Edge case*: no test in `register-form.test.tsx` currently asserts
  the literal `href` value/query string.

---

## Area 3: `lib/api/account.ts` + `lib/api/client.ts`

**Current state:**
- `client.ts`'s `apiFetch` is the sole sanctioned request path
  (api-client-centralization convention). Auth model: **in-memory
  access token only** — reads `useAuthStore.getState().accessToken`,
  sends `Authorization: Bearer <token>`. `credentials: 'include'` is
  always set too, but the code never *reads* a cookie value — only
  the Bearer header comes from application state. `tryRefreshOnce()`
  calls `POST /auth/refresh`, expects JSON `{access_token}` to
  repopulate the store.
- `account.ts`: exactly `getMe`, `register`, `verifyEmail`,
  `resendVerification`. **Nothing Google-OAuth-related.**

**Requirement (spec 02 + backend as actually implemented):** on
success, the backend delivers session tokens via **`writeAuthCookies`**
— both access and refresh tokens as HttpOnly cookies
(`backend/internal/transport/http/cookie.go:110-143`), explicitly
because "there is no JSON body to carry the access token" on a `302`.
This is a **documented deliberate exception** to every other issuance
path (`/auth/login`, `/auth/login/mfa`, `/auth/refresh` all deliver
the access token in a JSON body instead — same file, lines 71-74,
121-123).

**Gap — core architectural finding of this area:** `apiFetch`'s auth
model is entirely Bearer-header/in-memory-store based; there is **no
mechanism to bridge an HttpOnly cookie into `useAuthStore`** (and
structurally can't read the value directly — HttpOnly blocks JS
access by definition). Landing back on the app after a successful
Google login leaves the browser holding valid session cookies while
the SPA's own `accessToken` state is still `null`, until something
explicitly bridges the two. The only existing candidate mechanism is
`tryRefreshOnce()` (already calls `POST /auth/refresh`, already
returns `{access_token}` JSON) — but it's wired only as a reactive
401-retry inside `apiFetch` today, never called proactively.

**Sniffing:**
- *Risk*: if missed, "Login/Register dengan Google" appears to
  succeed (redirect happens, cookies set) while the SPA silently
  treats the user as logged out on the very next `apiFetch` call.
- *Inconsistency (cross-task, significant)*:
  `backend/internal/platform/auth/token.go`'s `VerifyAccessToken`
  (built under task #3, in progress) **explicitly documents that
  access tokens minted by the feature-02 Google OAuth flow lack the
  `purpose` claim and are rejected by design** — "accepted breaking
  edge for a sandbox with no deployed clients." Even a frontend that
  perfectly bridges cookie→Bearer may still fail auth on protected
  endpoints until backend task #2's minting is updated or task #3's
  verifier special-cases it. Primarily a backend-side gap, but it
  bounds what "verified end-to-end" can mean for this task's own
  testing today.
- *Miscontext*: page-map.md/spec 02 both describe the post-callback
  outcome as "redirect to app" / "issue tokens, 302 to app" as if that
  alone completes login — neither doc acknowledges the delivery-
  mechanism mismatch just described.
- *Edge case*: the access-token cookie is `SameSite=Lax`, `Path="/"`
  (not scoped like the OAuth state cookie) — it does travel with the
  landing page's own subsequent same-origin requests regardless of
  frontend JS. Whether that alone is sufficient (i.e. whether any
  backend endpoint reads the cookie directly instead of requiring
  Bearer) is unconfirmed — the cookie-first/Bearer-fallback pattern
  (`sessionToken()`) currently exists only inside `auth_google.go`'s
  own redirect-intent check, not confirmed as a general middleware
  pattern for other protected endpoints.

---

## Area 4: `lib/hooks/` + `lib/stores/auth-store.ts`

**Current state:**
- `auth-store.ts` — bare Zustand store:
  `{accessToken, setAccessToken, clearAccessToken}`, no persistence
  middleware (deliberate — avoids `localStorage`-readable tokens). Own
  doc comment, directly load-bearing: *"Shape only — no login logic
  lives here (that's Account Task #3's job); this exists so **Tasks #2
  and #3 share one store shape** instead of each inventing their
  own."* Explicitly names **this task (#2)** as one of the two tasks
  expected to actually populate the store.
- No `use-login.ts`, `use-google-oauth.ts`, or
  `use-auth-callback.ts` exists. Existing hooks
  (`use-register`, `use-account-me`, `use-resend-verification`,
  `use-verify-email`) are plain TanStack Query wrappers, one function
  each. `use-account-me.ts` exports a shared `accountKeys.me()`
  query-key factory — the natural cache entry to invalidate/refetch
  once a real session exists post-OAuth.
- No hook or store logic touches cookies, URL query params, or
  `window.location` anywhere today.

**Requirement:** per spec 02 + Area 3, something must run on the
frontend after landing from `/auth/google/callback` that (a) reads
whatever outcome signal the `302` carries, and (b) on success, obtains
a JS-readable access token and calls `setAccessToken` — i.e. populates
exactly the store `auth-store.ts` says this task owns.

**Gap:** no hook exists to do either half. `client.ts`'s
`tryRefreshOnce` is an unexported `async function` — only `apiFetch`,
`getAccessToken`, `setAccessToken` are exported. Even the
"call refresh to hydrate" mechanism from Area 3 needs either a new
exported entry point or a new function/hook reaching the same effect
some other way.

**Page-consolidation check:** confirms Area 3 — squarely this task's
own new work, not partially built or borrowed from elsewhere.

**Sniffing:**
- *Risk*: hydration timing — if a callback-landing page (or `/`
  itself, see Area 6) renders content before hydration resolves, any
  `apiFetch` call made too early (e.g. a mounted component's own
  `useAccountMe()`) goes out with no `Authorization` header,
  potentially racing `apiFetch`'s own 401→refresh→retry path.
- *Edge case*: `useAccountMe()` has no visible `enabled` gating — if
  anything mounted globally calls it unconditionally, it could fire
  before any hydration effect completes on first render after
  redirect.
- *Miscontext*: nothing in page-map.md/patterns.md acknowledges that
  "log in via Google" needs its own dedicated hook/effect at all — a
  fundamentally different code path (URL/cookie-triggered, not a
  mutation callback) from every other hook in this codebase.
- *Misleading signal*: `auth-store.ts`'s "Tasks #2 and #3 share one
  store shape" comment reassures that sharing is resolved — only the
  *shape* is shared; zero population logic exists for either task yet.

---

## Area 5: `mocks/handlers.ts`

**Current state:** one handler per endpoint an already-built page
consumes: `GET /account/me`, `GET /notifications/unread-count`,
`GET /campaigns`, `POST /auth/register`, `POST /auth/verify-email`,
`POST /auth/verify-email/resend`. **No handler for
`/auth/google/redirect`, `/auth/google/callback`, or `/auth/refresh`.**

**Requirement:** if this task adds a hydration step calling
`POST /auth/refresh` (Area 3/4), that endpoint needs a mock to test
against. The redirect/callback endpoints are real browser navigations
(302s) — MSW can't meaningfully intercept those the way it intercepts
`fetch`/XHR, so task 01's own approach (assert the `href` attribute
directly) is the right precedent, not a gap here.

**Gap:** exactly one plausibly-missing mock — `POST /auth/refresh` —
needed only if Stage 3 lands on "hydrate via refresh" as the
mechanism.

**Page-consolidation check:** N/A — this file trails the API layer
1:1, nothing to reconcile against page-map.md directly.

**Sniffing:**
- *Inconsistency (concrete, pre-existing)*: `client.ts`'s
  `tryRefreshOnce()` hand-writes
  `type RefreshResponse = { access_token: string }` locally instead of
  importing `components["schemas"]["RefreshResponse"]`, which
  **already exists** in generated `lib/api/schema.d.ts` (line 3849,
  also carries an optional `access_token_expires_at`). Directly
  violates `frontend/AGENTS.md` §3: *"Types for API requests/responses
  come from `lib/api/` (generated). Don't hand-write a parallel type
  for something the OpenAPI schema already defines."* Not introduced
  by this task, but sits right next to whatever this task adds for
  token hydration.
- *Edge case*: no fixture user has `auth_providers` containing
  `"google"` — a Google-originated-session test would need a fixture
  variant, same pattern as the existing `roles`-array override.

---

## Area 6: `app/verify-email/page.tsx` precedent, and the actual
callback-landing destinations

**Current state (precedent):** `/verify-email` is a deliberate
top-level route (not nested in `app/(auth)/`) — Decision D1: opened
cold from an email client, so the Auth Shell's modal-over-blurred-`/`
would misleadingly suggest prior in-app navigation. Uses the minimal
Status/Tracking shell instead. Reads its signal via
`useSearchParams()` inside `<Suspense>`, fires its mutation exactly
once (`useRef` guard), maps outcome to a discriminated union keyed off
`ApiError.status`, explicit focus management on outcome resolution.

**This precedent turned out not to be the actual answer** — reading
the already-merged backend's real redirect targets
(`backend/internal/domain/account/google_oauth.go`) shows the
destinations are not an open decision at all:

```go
func (s *Service) successResult(access, refresh string) CallbackResult {
    return CallbackResult{ RedirectURL: s.frontendURL, ... } // bare root, NO query param
}
func (s *Service) failResult(...) CallbackResult {
    base := s.frontendURL + "/login"                          // login-intent failure
    if intent == link || intent == reauth { base = s.frontendURL + "/account/security" }
    return CallbackResult{ RedirectURL: base + "?error=" + code }
}
```
(`FRONTEND_URL=http://localhost:3000` in `backend/.env.example`.)

**Requirement:** page-map.md/spec 02 imply *some* frontend surface
receives the outcome; Assumption B frames the exact route as open.

**Gap — no longer open, and consequential:**
1. **No dedicated callback-landing route to build.** For the in-scope
   `login` intent: success → existing `/` (Public Shell home), failure
   → existing `/login` (placeholder) with `?error={code}`. Area 6's
   original premise (find the closest precedent for a *new* route)
   doesn't apply — the real surfaces are two pages this task already
   touches for other reasons.
2. **Success carries zero query signal.** Unlike `/verify-email`
   (`?token=...`) or the failure leg (`?error=...`), success `302`s to
   `s.frontendURL` with nothing appended. `/`'s page code has no
   URL-based way to distinguish "just arrived from a successful Google
   login" from "ordinary guest visit." Combined with Area 3/4's
   hydration gap, this points toward some form of unconditional/
   opportunistic hydration-on-load rather than a URL-triggered one —
   a real design fork for Stage 3.
3. **`link`/`reauth` outcomes target `/account/security`**, which does
   **not match `page-map.md`'s actual route, `/dashboard/security`.**
   Out of scope for task 02 itself (link/reauth are tasks 05/06), but
   a concrete, already-shipped inconsistency worth flagging now since
   it'll block those tasks otherwise, and because task 02's own
   `?error={code}` parsing logic on `/login` is the same shape that'll
   be needed wherever that mismatch resolves.

**Page-consolidation check:** confirms Area 1 — no new routes, task 02
adds behavior to two existing routes (`/`, `/login`) plus finishes
`/register`'s already-half-built button.

**Sniffing:**
- *Inconsistency (significant, concrete)*: `/account/security`
  (backend constant, already merged) vs. `/dashboard/security`
  (page-map.md's actual route) — real drift, not hypothetical.
- *Miscontext*: spec 02's Assumption B frames the destination as
  still-open and promises "a distinguishable success/error query
  param" — the merged implementation already picked concrete paths,
  and the success case isn't distinguishable by query param at all.
  The spec's own acceptance framing and the shipped code disagree.
- *Risk*: because success has no query signal, hydration logic placed
  only on an isolated "callback page" would never run for the
  login-success path — it must live somewhere that executes on a
  normal `/` load, widening this task's blast radius to the Public
  Shell/home page, not just an isolated auth page.
- *Edge case*: an already-authenticated user re-triggering
  `intent=login` lands back on `/` with cookies overwritten —
  behaviorally fine per spec, but no existing precedent in this
  codebase covers "already-authenticated user re-runs a redirect-based
  auth flow."

---

## Summary, ranked by consequence

1. **Confirmed bug, already in the working tree**: `register-form.tsx`
   sends `intent=register`; the merged backend 400s anything but
   `login`/`link`/`reauth` (Area 2).
2. **Structural gap, explicitly in-scope per the store's own doc
   comment**: no bridge from the OAuth callback's HttpOnly-cookie
   token delivery into `useAuthStore`, which every `apiFetch` call
   depends on (Areas 3–4).
3. **Cross-task backend risk**: task #3's `VerifyAccessToken` rejects
   tokens minted by task #2's OAuth flow (missing `purpose` claim) —
   bounds what "verified end-to-end" can mean right now (Area 3).
4. **Callback destinations are already fixed by merged backend code,
   not an open frontend decision**: success → bare `/`, no query
   signal; login failure → `/login?error=...`; link/reauth →
   `/account/security` (mismatched against `page-map.md`'s
   `/dashboard/security`) (Area 6).
5. **`/login` is still a raw Phase-0 placeholder**; the credential
   form is explicitly a different task's scope — this task's job
   there is narrower than page-map.md's one-line description implies
   (Area 1).
6. Minor, adjacent: hand-written `RefreshResponse` type in `client.ts`
   duplicates an already-generated schema type (Area 5).
