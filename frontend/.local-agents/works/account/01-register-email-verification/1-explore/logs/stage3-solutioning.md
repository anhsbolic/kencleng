# Stage 3 — Solutioning: account/01-register-email-verification (frontend)

> Builds on `stage2-gap-analysis.md`'s "Cross-area summary of open
> items." Each item below is a Decision Log entry: question → options
> considered → recommendation → rationale. This is raw solutioning
> material, not a techplan — synthesis into a techplan is a separate,
> later step.

Grounding note found while researching these decisions:
`kencleng-agentic-workflow.md` §14 explicitly names *"an
email-verification link landing route"* as the canonical example of "a
backend task with no dedicated page [that] still needs an explicit
frontend surface, however small" — confirming Stage 2's Area 2 finding
(the missing route) isn't a novel problem, it's the exact scenario the
workflow doc anticipated for this feature.

---

## D1 — Route path for the email-verification-link landing surface

**Question:** What path/route does the emailed verification link point
to, and does it live inside `(auth)`'s `AuthShellClient` or somewhere
else?

**Options considered:**
- **A. `app/(auth)/verify-email/page.tsx`** — nested inside the
  existing Auth Shell group, consistent with `/register`,
  `/login`, `/forgot-password`, `/reset-password` all living there.
- **B. `app/verify-email/page.tsx`** — top-level route, using the
  **Status/Tracking pattern**'s minimal shell (`patterns.md` §A.6: "no
  Dashboard Shell — guest has no session"), the same shape already
  established for `/donation/[id]/status`.
- **C.** Some other path (e.g. `/auth/verify`) — rejected outright,
  no reason found to deviate from the flat top-level naming
  `/reset-password` already sets as precedent for a token-in-email
  flow.

**Recommendation: B.**

**Rationale:**
- `POST /auth/verify-email` is unauthenticated and reached by clicking
  a link in an email client — there is no "page beneath a modal" for
  `AuthShellClient`'s desktop treatment to make sense of. Confirmed by
  reading the extracted `login-register.extracted.jsx` (Stage 2, Area
  3): the desktop variant renders a **blurred `<Landing>` page behind
  the modal** — implying the user was mid-browse on `/` when the modal
  opened. That's simply false for someone arriving fresh from an email
  link; showing a fabricated blurred backdrop behind the verification
  result would be visually misleading about what just happened.
- The Status/Tracking pattern is the structurally identical existing
  precedent: no auth, a token resolved from the URL, one API call, a
  small set of terminal outcomes rendered on a single card. Reusing it
  here is following an established pattern rather than stretching the
  Auth Shell to cover a case it wasn't designed for.
- Path symmetry with `/reset-password?token=...` (also a page-map-named,
  top-level, token-in-email route) makes `/verify-email?token=...`
  the naturally consistent choice — same domain, same `auth_tokens`
  mechanics (INV-account-08), same shape.

**Open sub-item for whoever builds this:** the success state should
still offer a way back into the app (e.g. a `Button` linking to
`/login`) — that's just a link, not a reason to nest inside the shell.

---

## D2 — Where the "resend verification email" action surfaces

**Question:** `POST /auth/verify-email/resend` has full acceptance
criteria in the feature spec but zero page-map.md presence. Where does
a user actually trigger it?

**Options considered:**
- **A.** A "Belum menerima email? Kirim ulang" link on `/register`'s
  post-submit success state only.
- **B.** A resend action embedded in `/verify-email`'s **expired-token**
  state specifically.
- **C.** Both A and B (same underlying mutation/hook, two call sites).
- **D.** A dedicated `/resend-verification` page.

**Recommendation: C (both A and B), reject D.**

**Rationale:**
- Option B is close to spec-mandated, not just good UX: the `410`
  response's own documented example `detail` copy is *"Link verifikasi
  sudah kedaluwarsa. **Silakan minta kirim ulang.**"* (Stage 2, Area 2 —
  read directly from `schema.d.ts`'s response example). The backend
  contract itself is already written assuming the frontend offers a
  resend action at exactly this point — skipping it would leave the
  API's own copy pointing at a UI affordance that doesn't exist.
- Option A is the other natural failure point (email delayed, spam
  filter, typo) and costs little once the resend mutation/hook exists
  for B — same `ResendVerificationRequest { email: string }` call, just
  a second entry point.
- Option D is unnecessary complexity: a whole extra route/page for what
  is, in both A and B, a single field + button that can be composed
  inline into an existing screen's success/error state.
- All resend responses are `202` generic regardless of match (per the
  spec) — so neither surface needs to special-case "email not found"
  differently from "resent successfully"; same generic confirmation
  copy either way, preserving the anti-enumeration property.

---

## D3 — Scope split: "Daftar dengan Google" button (Task #1) vs. Google
OAuth flow (Task #2)

**Question:** `page-map.md` lists the Google button as part of
`/register`'s action set, but the actual OAuth flow is
`02-google-oauth-login-register.md`'s scope. Where's the line?

**Options considered:**
- **A.** Defer the entire button (including its presence) to Task #2 —
  ship `/register` without it, add it when Task #2 lands.
- **B.** Build the button in Task #1 as a real, working navigation
  trigger — `<a href="/auth/google/redirect?intent=register">` (or a
  `Button` styled the same way) — and let Task #2 own only what happens
  after Google redirects back (`GET /auth/google/callback`).

**Recommendation: B.**

**Rationale:**
- `schema.d.ts`'s own doc comment on the redirect endpoint (read in
  Stage 2 Area 2, ~line 481) describes it as *"one shared
  redirect-initiation endpoint for login/register/link/reauth,
  distinguished by `intent`... issues a 302 to Google's consent
  screen"* — from the frontend's side this is a **plain browser
  navigation to a URL**, not an OAuth client library integration. There
  is no client-side logic to build that would meaningfully depend on
  Task #2 being done first — the button just needs to exist and point
  at the right URL with `intent=register`.
- `page-map.md` describes the button as part of `/register`'s action
  set, not a separate page — leaving it out entirely (Option A) would
  make Task #1's `/register` incomplete against its own page-map row
  for no real technical reason.
- What *is* genuinely Task #2's scope: handling `GET
  /auth/google/callback`'s result (session establishment, error
  states if Google denies consent, the no-auto-merge conflict UX) —
  none of that needs to exist for the button itself to work correctly
  today.
- Noted in passing, not resolved here: `tasks.md`'s prose sequencing
  ("Register → Login → Forgot Password → Google OAuth → Linking")
  doesn't match its own numbered table (`#2` = Google OAuth,
  immediately after `#1` = Register) — a minor internal inconsistency
  in that doc, irrelevant to this recommendation since B doesn't depend
  on which order #2/#3/#4 actually land in.

---

## D4 — Structured validation-error shape for `lib/api/account.ts`'s
`register()`

**Question:** The existing `lib/api/` pattern (`getMe`, `getCampaigns`)
throws a flat `Error` on any non-OK response and discards the response
body. Register's `422` needs to surface a `{field, message}[]` array to
the form. What should `register()`'s contract be?

**Options considered:**
- **A.** Keep throwing, but throw a custom `class ValidationError extends
  Error { errors: {field, message}[] }` on 422, throw plain `Error` for
  everything else; caller `catch`es and does an `instanceof` check.
- **B.** Return a discriminated union instead of throwing for the 422
  case: `{ ok: true } | { ok: false; kind: "validation"; errors: [...] }`,
  reserving `throw` exclusively for genuine request-level failures
  (network error, unexpected 5xx).
- **C.** Return the raw `Response` and let the calling hook parse
  status/body itself.

**Recommendation: B.**

**Rationale:**
- `patterns.md` §B is explicit: *"never conflate"* field-level (422)
  and request-level (network/5xx) failures — they're handled by
  different UI (inline field error vs. banner). Option A still routes
  both cases through the same `throw`/`catch` mechanism and relies on
  every caller remembering to `instanceof`-check correctly; a missed
  check silently mis-renders one as the other. Option B makes the two
  cases structurally distinct at the type level — a caller literally
  cannot access `.errors` without first checking `ok: false` and
  `kind: "validation"`, and anything that reaches the `throw`/`catch`
  path is unambiguously request-level by construction.
- Option C pushes response-shape knowledge (the `Problem`/validation
  JSON shape from `schema.d.ts`) out of `lib/api/account.ts` and into
  the hook layer, duplicating parsing logic if more than one caller
  ever needs it — `lib/api/` is the right layer to own that once.
- This deliberately deviates from `getMe`/`getCampaigns`'s existing
  throw-only shape — that's a decision, not an oversight: those two
  endpoints have no field-level-error case to represent (`GET
  /account/me` and `GET /campaigns` don't take user input), so their
  simpler contract was always incomplete as a *template* for a
  validating `POST`, not wrong for what they do.
- `verifyEmail()` and `resendVerification()` don't need this shape —
  neither has a field-level-error case in their acceptance criteria
  (verify-email's errors are `404`/`410`, both request/outcome-level,
  not per-field; resend never distinguishes at all) — so those two can
  keep the simpler `getMe`-style throw-on-`!ok` contract. Only
  `register()` needs the discriminated-union treatment.

---

## D5 — Distinguishing `410` (expired) vs `404` (not found/used/revoked)
in `/verify-email`'s UI copy

**Question:** `patterns.md`'s Status/Tracking pattern (written for
`/donation/[id]/status`) says invalid/missing states must **not** be
distinguished in copy, to avoid confirming/denying existence to an
attacker. Does that rule transfer to `/verify-email`?

**Options considered:**
- **A.** Flatten `410` and `404` into one generic "link invalid or
  expired" message, following the Status/Tracking pattern's rule
  literally.
- **B.** Distinguish them — `410` gets its own message (with a resend
  CTA, per D2), `404` gets a separate generic "link invalid or already
  used" message — following what the backend contract already does.

**Recommendation: B.**

**Rationale:**
- The enumeration concern the Status/Tracking rule exists for is about
  a guessable/sequential **resource identifier** (a donation ID an
  attacker could iterate over to learn which IDs exist). A verification
  token is a high-entropy, single-use random string — knowing "this
  token happens to be expired vs. not found" doesn't let an attacker
  learn anything about *other* tokens or *which emails are
  registered*; the risk model the original rule was written for doesn't
  hold here in the same way.
- The backend has **already made this distinction at the contract
  level** — different status codes (`410` vs `404`), and the `410`
  response ships its own specific, actionable `detail` copy in the
  OpenAPI schema example (Stage 2, Area 2). Flattening these in the
  frontend would mean *hiding* backend-provided, non-sensitive,
  actionable information (the resend prompt) for no corresponding
  security benefit — actively working against D2's resend-surface
  design rather than protecting anything.
- This is scoped narrowly to `/verify-email`'s two outcomes — it is
  not a general license to distinguish enumeration-sensitive responses
  elsewhere; the register endpoint's own `202`-always-generic design
  (Assumption A/B in the feature spec) is a completely different case
  and must stay flattened as specified.

---

## Overall `/register` page composition (ties D3, D4 together)

Not a new open question, but worth recording the shape this converges
on, given D1–D5 above and Stage 2's Area 4 findings (primitives already
built):

- **Fields**: email (`Input`), password (`Input` + show/hide affordance,
  per the `login-register.extracted.jsx` precedent for the password
  field's reveal toggle), submit `Button` (`type="submit"`, per Stage 2
  Area 4's gotcha), "Daftar dengan Google" `Button`/link (D3).
- **Validation**: `zod` schema — email format, password ≥8 chars only
  (breach-check is server-only, surfaces as a 422 field error via D4's
  discriminated-union return, not a client-side rule).
- **States**: idle → validating (blur+submit) → submitting (disable
  form, `Button`'s `loading` prop) → submit error (`Banner
  variant="error"` as the shell's documented first child, for
  request-level failures only, per D4) → field errors (`Input`'s
  `error` prop, per D4's validation branch) → success (inline "check
  your email" state, replacing the form entirely, uniform across all
  four backend branches per the anti-enumeration requirement) with a
  "Belum menerima email? Kirim ulang" affordance (D2, Option A).
- **Data layer additions**: `register()` (D4's discriminated union),
  `verifyEmail()`, `resendVerification()` in `lib/api/account.ts`; one
  new `useMutation`-based hook per action in `lib/hooks/` (first
  mutation hooks in the codebase — sets precedent for Tasks #3/#4); MSW
  handlers for all three endpoints in `mocks/handlers.ts`.
- **New route**: `app/verify-email/page.tsx` (D1) — top-level, minimal
  Status/Tracking-style shell, reads `token` from `searchParams`, calls
  `verifyEmail()`, renders success / expired (with resend, D2 Option B)
  / not-found-or-used outcomes per D5.

This composition is a synthesis of the decisions above, not a new
decision itself — provided here so a later techplan-synthesis pass has
a single place to start from rather than re-deriving it from D1–D5
separately.
