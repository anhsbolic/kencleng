# Stage 2 — Gap Analysis: account/01-register-email-verification (frontend)

> Feature spec: `docs/spec/1-account/features/01-register-email-verification.md`
> Backend status: **merged** (`14834e5`, backend-only commit — confirmed
> zero `frontend/` files touched). Frontend surface: not started.

---

## Area 1: Auth Shell + existing `(auth)` routes

**Current state:**

- `app/(auth)/layout.tsx` wraps all four auth routes in `AuthShellClient`
  (`app/(auth)/_components/auth-shell-client.tsx`), built in the Phase 0
  `phase0-shared-infra.md` playbook. Desktop-modal / mobile-full-page
  split is CSS-only (`md:` classes); only focus-trap *activation*
  (`useFocusTrap`, `lib/hooks/use-focus-trap.ts`) is gated in JS via
  `useIsDesktop()` (a `useSyncExternalStore` on `matchMedia`).
- The shell's doc comment establishes a load-bearing convention for pages
  rendered inside it: *"render a `<Banner variant="error">` as the first
  child, before the form"* — exists specifically to structurally prevent
  the known prototype bug (`prototype-reference.md` issue #1: login's
  generic auth failure wrongly rendered as a field-level error). Applies
  to `/register`'s own request-level failure banner too (per
  `patterns.md`'s Form pattern: banner for request-level failure, kept
  separate from field-level 422 errors).
- **All four routes inside the shell are still placeholders**, not just
  `/register`: `register/page.tsx` ("Account Task #1"), `login/page.tsx`
  ("Account Task #3"), `forgot-password/page.tsx` +
  `reset-password/page.tsx` (both "Account Task #4"). Each is a 2-line
  `<h1>`+`<p>` stub, explicitly commented as existing only so the Auth
  Shell had a route to render against during Phase 0 verification.
- Confirmed: **`/register` does not exist as a built route today** — only
  as an empty stub + layout wrapper. No sibling form implementation
  exists anywhere in the repo to use as precedent — login/forgot/reset
  are equally unbuilt.
- `auth-shell-client.test.tsx` exists and tests the shell's own
  focus-trap/breakpoint behavior, not any page content.

**Requirement:**

- `page-map.md`: `/register` — Form pattern — "Register form + 'Daftar
  dengan Google' button."
- `patterns.md` Form pattern: idle → inline validation (blur+submit,
  zod) → submitting (disable form, inline spinner) → submit error
  (banner for request-level, field-level for 422) → success (toast+
  redirect for dashboard forms, **or inline success state for
  guest-facing forms** — register is the latter, per `patterns.md` §B's
  "clear what happens next" convention for terminal flow actions).
- Feature spec: `POST /auth/register` always returns `202` + generic
  message, **regardless of which internal branch fired** (new user /
  resend-nudge / already-verified nudge / Google-only-conflict nudge) —
  the only client-distinguishable branch is `422` (password fails length
  policy or breach-listed).

**Gap:**

- Entire `/register` form needs building from scratch: email + password
  fields, client-side zod validation (password ≥8 chars only — the
  breach-list check is HIBP/server-only, so that 422 case is a
  round-trip error, not a pre-submit local check), "Daftar dengan
  Google" button, submitting/disabled state, and a generic **inline
  success state** ("check your email") shown identically regardless of
  which of the four server branches actually ran — the frontend must
  not differentiate them in copy/UI, or it reintroduces the enumeration
  leak the backend closed.
- Whether a `Banner` component already exists was an open question here
  — **resolved in Area 4: it already does** (`components/ui/banner.tsx`).

**Page-consolidation check:** `/register` has exactly one page-map
action. The placeholder's own comment ("Account Task #1's scope") and a
separate feature file (`02-google-oauth-login-register.md`) suggest the
"Daftar dengan Google" **button** is this task's UI concern, but the
**OAuth flow behind it** is Task #2's — needs explicit confirmation
against `tasks.md`, not assumed here.

**Sniffing:**
- *Misleading signal*: `app/(auth)/register/page.tsx` existing as a file
  could look, from a directory listing alone, like the route is already
  handled — it's a 6-line stub with zero form logic.
- *Miscontext risk*: `page-map.md`'s terse action description doesn't
  mention a distinct success/"check your email" state — that only
  becomes visible from the feature spec's 202-generic-response
  requirement, not from the page-map row alone.
- *Risk*: the anti-enumeration guarantee is a security property the
  **frontend can break by choice of copy** even if the backend never
  leaks anything — e.g. showing different text for "already registered"
  vs. "verification sent" would leak exactly what the backend's uniform
  `202` was designed to prevent. Frontend-owned risk, not just backend.

---

## Area 2: Non-page email-verification-link surface

**Current state:**

- No route exists anywhere in `app/` for consuming a verification link
  — `grep -ril "verify" app/` finds nothing under `app/`; only hits in
  `lib/api/schema.d.ts` (generated types).
- `schema.d.ts` already has full generated types for all three endpoints
  (confirms backend contract is finalized):
  - `POST /auth/verify-email` — body `VerifyEmailRequest { token: string
    }`, responses `200` / `404` (not found/used/revoked) / `410`
    (expired) / `429`. **No auth header** — unauthenticated endpoint,
    the token itself is the credential.
  - `POST /auth/verify-email/resend` — body `ResendVerificationRequest
    { email: string }`, always `202` generic + `429`.
- `lib/api/account.ts` has no `register`/`verifyEmail`/
  `resendVerification` wrappers yet (Area 4 territory, noted here
  because it's the same missing surface).
- **Backend confirmation**: `docs/spec/1-account/tasks.md`'s status
  tracker shows Task #1 as `merged` — *"shipped on main (`14834e5`
  finalize...)"*. Checked directly: `14834e5`
  (`[backend] finalize feature account/01-register-email-verification`)
  touches only `backend/` + its own `.local-agents` folder — zero
  `frontend/` files. Confirms frontend surface is genuinely untouched
  anywhere, not partially done.
- `page-map.md`'s only line for this: *"Email verification link (from
  email, not a full page) | — | Click link → `AuthIdentity.verified_at`
  set"* — no route path, no pattern name (the `—` is deliberate, unlike
  every other row).

**Requirement:**

- Something must physically exist at whatever URL the verification
  email points to. The link necessarily carries a `token` as a URL
  param (query string), since `VerifyEmailRequest` takes it as a JSON
  body field — the frontend route must read it from the URL and re-POST
  it as JSON.
- Closest existing structural precedent: **Status/Tracking pattern**
  (`patterns.md` §A.6) — "Minimal shell (no Dashboard Shell — guest has
  no session) → single status card, resolved via token-in-URL lookup,"
  currently only used by `/donation/[id]/status`. Same shape: no auth,
  token-in-URL, one API call, small number of terminal outcomes.
- Feature spec's three endpoints are the full scope of this feature —
  resend is a full acceptance-criteria section, not an afterthought.

**Gap:**

- No route exists for the verification-link landing surface — not even
  a placeholder (unlike the other four auth routes, which all got
  Phase 0 stubs). Genuine void, not an unfinished stub.
- **The path itself is undefined.** Unlike `/reset-password?token=...`,
  named explicitly in `page-map.md`, the email-verification link has no
  named path anywhere in `docs/ui-ux/` or `docs/spec/`. Deciding it is a
  Stage 3 concern; the *absence of the decision* is the Stage 2 finding.
- **No UI surface exists anywhere for `POST /auth/verify-email/resend`.**
  Not on `/register`, not as its own page-map row. The feature spec
  fully specifies the endpoint's behavior, but there is no documented
  place in the UI a user would trigger it from.

**Page-consolidation check:** Task #1, first frontend task in this
domain — nothing earlier to consolidate against. Per workflow §14
step 1: found the exact asymmetry it looks for — the spec's 3rd
endpoint (`resend`) has no backing page-map action, and the page-map's
verification-link row has no backing route path. Both real gaps, not
silently resolved here.

**Sniffing:**
- *Miscontext*: page-map's "(from email, not a full page)" phrasing
  reads like "this needs no page, just a click handler," but some route
  must exist to receive the click. Risk that an implementer skips the
  route entirely, assuming the backend "handles" it somehow.
- *Inconsistency*: `/reset-password?token=...` — structurally identical
  link-in-email flow, same domain, same `auth_tokens`/INV-account-08
  pattern — is a full page-map row with its own route; the verify-email
  link is not, despite being architecturally the same shape. No stated
  reason for the asymmetry.
- *Risk*: since `POST /auth/verify-email` is unauthenticated and reached
  via an emailed link (possibly a different browser/device/session than
  registration), nesting this route inside `AuthShellClient`'s modal
  treatment (which assumes a "page beneath a modal" desktop context)
  would likely be a wrong default choice — flagging so Stage 3 decides
  shell placement deliberately.
- *Edge case*: backend distinguishes `410` (expired) vs `404` (not
  found/used/revoked), but `patterns.md`'s Status/Tracking pattern says
  invalid/missing-token states must **not** be distinguished in
  user-facing copy (avoid confirming/denying existence) — written for
  `/donation/[id]/status`; whether it applies identically here (expired
  vs. "not found," where "not found" is the enumeration-sensitive one)
  is a real judgment call, not yet made anywhere.

---

## Area 3: `design-reference/login-register.html`

> **Process note**: extraction initially wrote `.extracted.*` files
> directly into `docs/design-reference/`, which `AGENTS.md` §3 marks
> read-only for agents. Caught immediately, deleted both files,
> confirmed via `git status` the directory was left clean, and
> re-ran extraction into the scratchpad instead. No lasting effect.

**Current state:**

- Extracted `login-register.extracted.jsx` (111 lines). Despite the
  filename, **it contains zero register-form content** — only
  `LoginForm`, `PasswordField`, `Divider`, `Panel`, `DesktopOverlay`,
  `MobilePage`, and four rendered mount points (`d-default`, `d-error`,
  `m-default`, `m-error`), all login states (default + error, desktop +
  mobile).
- Only trace of "register": a static link at the bottom of `LoginForm`
  — `Belum punya akun? <a href="#">Daftar</a>` — a plain anchor, not a
  form, route, or second panel/state.
- Independently corroborates `prototype-reference.md`'s Known Issue #1:
  `LoginForm` passes the generic auth-failure message directly into the
  **Email** `<Input>`'s `error` prop (field-level), not a page banner —
  confirmed present in the actual export exactly as documented.

**Requirement:** `prototype-reference.md`'s Tier 2 table already lists
`/register`'s precedent as `/dashboard/campaign/new` (dashboard forms)
*or* `/login` (auth modal/mobile split) — i.e. it already assumed no
dedicated register prototype exists.

**Gap:** None beyond what was already assumed — this closes the Stage 1
question rather than opening a new one. The filename is the only thing
suggesting broader coverage than the table claims; content confirms the
Tier 2 classification is correct as-is. The one reusable thing for
`/register`: the **shell composition pattern** (`DesktopOverlay`/
`MobilePage`, panel sizing, logo block, `Divider` "atau" pattern, Google
button with inline `GoogleG` SVG) — useful precedent for `/register`'s
Google button and panel layout, same `AuthShellClient` shell.

**Page-consolidation check:** N/A — reference material, not a route.

**Sniffing:**
- *Misleading signal*: filename `login-register.html` reads as "register
  is covered here" without opening the file — confirmed it is not.
- *Risk*: the field-level auth-error known-issue must specifically not
  be copied when building `/register`'s submit-error handling —
  reinforces why `AuthShellClient`'s banner-first convention exists.

---

## Area 4: API/data layer + component primitives

**Current state:**

- `lib/api/schema.d.ts` has full generated types for all 3 endpoints
  (see Area 2), matching the shipped backend contract.
- `lib/api/account.ts` exports only `getMe()` — no `register`,
  `verifyEmail`, or `resendVerification` functions. Established pattern
  (also used by `campaign.ts`'s `getCampaigns`): call `apiFetch`,
  `throw new Error("generic message")` on `!res.ok`, return
  `res.json()`. This is a **GET-only precedent** — neither existing
  function preserves a structured error body on failure. Register's
  `422` response carries a structured `{field, message}[]` payload the
  form needs to map to fields — this would be the **first** function in
  the codebase needing to surface a structured error body rather than a
  flat thrown message.
- `lib/api/client.ts`'s `apiFetch` is endpoint-agnostic and already
  handles unauthenticated calls correctly (omits `Authorization` when
  `accessToken` is `null`) — no changes needed there.
- `mocks/handlers.ts` has zero `/auth/*` handlers — only `GET
  /account/me`, `GET /notifications/unread-count`, `GET /campaigns`.
  Its own comment states handlers are added "as the page/component that
  needs it gets built... not speculatively ahead of demonstrated need"
  — expected, not a defect.
- `lib/hooks/` has no register/verify/resend hook, and **no existing
  mutation hook (`useMutation`) anywhere in the codebase** — only
  `useQuery`-based read hooks (`use-account-me.ts`, `use-campaigns.ts`,
  etc.) plus non-data hooks (`use-has-role.ts`, `use-focus-trap.ts`,
  `use-unread-count.ts`). This task would be the first TanStack Query
  mutation hook in the project.
- **Component primitives are already built and ready** — resolves Area
  1's open question about `Banner`: `components/ui/banner.tsx` exists,
  exactly matching `AuthShellClient`'s doc-comment convention
  (success/error/warning/info variants, correct `role="alert"`/
  `"status"` split). Also ready: `input.tsx` (`Input`, `error` prop →
  `aria-invalid`/`aria-describedby`), `button.tsx` (`Button`, `loading`
  prop → disables + `Spinner`), plus `label.tsx`, `spinner.tsx`,
  `badge.tsx`, `progress-bar.tsx`. **None of this needs building from
  scratch.**
- Gotcha: `Button` hardcodes `type="button"` as its default (not HTML's
  own `"submit"` default) — a submit button must explicitly pass
  `type="submit"`.
- `react-hook-form` (^7.86.0) and `zod` (^4.4.3) are both in
  `package.json`, but `grep` found **zero usages** of either anywhere in
  `app/`, `components/`, or `lib/`. This task would be the first real
  form built with either library in the codebase — no existing
  composition pattern to match against.

**Requirement:** Per `AGENTS.md` §3, forms use `react-hook-form` + `zod`
matching the feature spec's validation rules. Data mutations should be
TanStack Query `useMutation` hooks in `lib/hooks/`, matching the
existing file-per-hook + query-key-factory style.

**Gap:**
- Three new `lib/api/account.ts` functions needed (`register`,
  `verifyEmail`, `resendVerification`) — and, unlike the existing two
  functions, these need to surface the `422` structured validation body
  instead of collapsing it into a flat `Error`. No existing precedent
  for this shape; needs fresh design (Stage 3).
- Three new MSW handlers needed in `mocks/handlers.ts`.
- At least one new `useMutation`-based hook needed (first of its kind).
- The register form's visual-primitive gap is narrower than Area 1
  first suggested: `Banner`/`Input`/`Button` are all ready — the real
  gap is form composition + validation schema + data-layer wiring, not
  missing UI components.

**Page-consolidation check:** N/A (not a route) — but
`component-test-mocking-discipline.md` is directly relevant once mocks
are added, per the same discipline already visible in `campaign`'s
fixture comments in this file.

**Sniffing:**
- *Misleading signal*: `getMe`/`getCampaigns` as "the established
  `lib/api/` pattern" could tempt a copy-paste for `register` — but that
  pattern silently drops the response body on error, which would break
  the field-level 422 requirement if copied uncritically.
- *Edge case*: `Button`'s non-standard `type="button"` default is an
  easy, silent mistake — a forgotten `type="submit"` fails to submit on
  Enter/click with no visible error, only caught by actually testing the
  interaction.
- *Risk*: this is the first `react-hook-form`+`zod` form and the first
  TanStack Query mutation hook in the codebase — whatever shape this
  task establishes becomes the de facto precedent for `/login`,
  `/forgot-password`, `/reset-password` (Tasks #3/#4, same Auth Shell)
  and every other form in the app. Outsized downstream reach.

---

## Area 5: `lib/stores/auth-store.ts`

**Current state:** `useAuthStore` (`accessToken`, `setAccessToken`,
`clearAccessToken`) is consumed only by `lib/api/client.ts`'s
`apiFetch`/`tryRefreshOnce` — `grep` confirms zero other consumers
anywhere in `app/`, `components/`, or `lib/`. Doc comment: "no login
logic lives here (that's Account Task #3's job)."

**Requirement:** Per the feature spec, `POST /auth/register` never
returns a token or `user_id` (Assumption A) and never logs the user in.
`POST /auth/verify-email` similarly returns only a message.

**Gap:** None. This area needs no changes for this task.

**Page-consolidation check:** N/A.

**Sniffing:** *Miscontext (resolved, not live)*: a plausible-but-wrong
assumption would be "register → auto-login" (common elsewhere) — spec
explicitly rules this out, and the store's current shape already
reflects that correctly. No inconsistency found.

---

## Cross-area summary of open items for Stage 3

1. Decide the route/path for the email-verification-link landing
   surface (Area 2) — currently undefined in every doc.
2. Decide where/how "resend verification email" is surfaced in the UI
   (Area 2) — currently has zero page-map presence despite being a full
   endpoint in this feature's scope.
3. Confirm split between this task's "Daftar dengan Google" **button**
   vs. Task #2's OAuth **flow** (Area 1) against `tasks.md`.
4. Design the structured-422-error shape for `lib/api/account.ts`'s
   `register` function (Area 4) — first of its kind in the codebase.
5. Decide whether the verify-email landing route sits inside
   `AuthShellClient` or uses a minimal shell like `/donation/[id]/status`
   (Area 2).
6. Decide whether `410` vs `404` get distinguished in verify-email UI
   copy, and how that interacts with the anti-enumeration convention
   already established for Status/Tracking pages (Area 2).
