# Stage 1 — Plan Announcement
## Feature: account/02-google-oauth-login-register (frontend surface)

## Docs read so far

- `docs/spec/README.md` — doc-type conventions (spec vs project docs).
- `docs/spec/1-account/features/02-google-oauth-login-register.md` —
  two backend endpoints (`GET /auth/google/redirect?intent=...`,
  `GET /auth/google/callback`), Tier 1, shared across login/register,
  account-linking (`05`), and MFA-disable reauth (`06`). Key open
  point for the frontend track: **Assumption B** explicitly says the
  concrete frontend landing route(s) per `intent`/outcome are
  "deferred to the frontend track" — not decided by this spec.
- `docs/ui-ux/page-map.md` — Guest table has `/login` ("Login form +
  'Masuk dengan Google' button") and `/register` ("Register form +
  'Daftar dengan Google' button"). Donatur table's only non-page row
  is the *email verification* link, not anything Google-OAuth-shaped.
  The doc's own "Shell & Benchmark Notes" section claims a "Google
  OAuth full-redirect flow" is "defined in `patterns.md`."
- `docs/ui-ux/patterns.md` — `/login`/`/register` are tagged "Form
  (Auth sub-variant)" under the Revisable-Submission-adjacent Form
  pattern (§A.3). **No section anywhere in this file actually
  describes a Google OAuth / full-redirect flow** — grepped for
  `google`, `oauth`, `redirect`, `full-redirect`, `auth modal`; only
  hits are the generic Success-state redirect note (line 98) and an
  unrelated line 189. This directly contradicts what `page-map.md`
  claims — flagged for Stage 2, not resolved here.
- `docs/ui-ux/prototype-reference.md` — `/login` is **Tier 1**
  (`design-reference/login-register.html`), including a known issue
  ("Login error state" must render as a form-level banner, not a
  field-level error — not confirmed fixed in the export). `/register`
  is **Tier 2**, explicitly pointed at `/login`'s prototype as its
  precedent (Form pattern row).
- `docs/design-reference/` listing — confirmed `login-register.html`
  exists (one file, plausibly covering both `/login` and `/register`
  despite the "one Tier-1 route → one file" framing in the table).
  No separate `google-callback` or similar file — no prototype exists
  for any callback-landing surface.
- `docs/ui-ux/design-guidelines.md` — color/type/shape/component
  tokens (buttons, inputs, badges). Nothing OAuth-specific; will apply
  generically (e.g. secondary/outline button style for "Masuk/Daftar
  dengan Google").

## Shell & current repo state (frontend/AGENTS.md, .agents/docs/README.md)

Per `frontend/AGENTS.md` and `docs/ui-ux/page-map.md`, `/login` and
`/register` sit under the **Auth Shell**, one of the three Phase-0
shells (Public / Auth / Dashboard). The Auth Shell already exists and
is already populated:

```
app/(auth)/layout.tsx
app/(auth)/_components/auth-shell-client.tsx (+ .test.tsx)
app/(auth)/login/page.tsx
app/(auth)/register/page.tsx
app/(auth)/forgot-password/page.tsx
app/(auth)/reset-password/page.tsx
```

`/register` was already built by the prior domain task
(`01-register-email-verification` — see
`.local-agents/works/account/01-register-email-verification/`), with
supporting pieces already in place:

```
components/features/account/register-form.tsx (+ schema, tests)
components/features/account/verify-email-status.tsx
components/features/account/resend-verification-control.tsx
lib/api/account.ts, lib/hooks/use-register.ts,
lib/hooks/use-resend-verification.ts, lib/hooks/use-verify-email.ts
app/verify-email/page.tsx  ← a non-page-map "redirect-landing" route,
  directly relevant precedent for this task's own callback surface
```

`git status` shows `app/(auth)/register/page.tsx`, `lib/api/account.ts`,
`lib/api/client.ts`, and `mocks/handlers.ts` are all already modified
in the working tree (uncommitted), on top of the last commit. `/login`
has no dedicated feature component yet visible in
`components/features/account/` (only `register-form.tsx` and
verify-email-related components exist there) — worth confirming, not
assuming, in Stage 2.

This means Stage 2's page-consolidation check is central here: `/login`
and `/register` **already exist as routes**, so this task is very
likely an *extension* of existing pages/components (adding the Google
button + wiring the redirect/callback flow), not new-page creation —
but I haven't opened the actual page/component bodies yet to confirm
whether a Google button placeholder already exists or not.

## Areas I intend to explore in Stage 2, in order

1. **Auth Shell + `/login` + `/register` pages** (`app/(auth)/layout.tsx`,
   `_components/auth-shell-client.tsx`, `login/page.tsx`,
   `register/page.tsx`) — first, since it directly answers the
   page-consolidation question: are these routes fully built already,
   and is there already a Google-button slot/placeholder to extend.

2. **`components/features/account/`** (`register-form.tsx` +
   `register-schema.ts` and siblings) — task 01 already built the
   register form; need to see its actual composition and confirm
   whether an equivalent `login-form.tsx` exists yet, to know where a
   "Masuk/Daftar dengan Google" CTA slots into the existing component
   structure rather than assuming a shape.

3. **`lib/api/account.ts` + `lib/api/client.ts`** — the centralized API
   client layer (per the `api-client-centralization` skill topic).
   Need current account API surface (register, verify-email,
   resend-verification, account/me are visible in git status) and how
   `client.ts` currently handles CSRF/credentials, since the OAuth
   step is a full browser navigation (not a fetch call) — a different
   shape from every other endpoint this client currently wraps.

4. **`lib/hooks/` + `lib/stores/auth-store.ts`** — existing hooks
   (`use-register`, `use-account-me`, `use-resend-verification`,
   `use-verify-email`) and the Zustand in-memory-token auth store.
   No `use-login` hook exists yet — need to confirm there's no
   email/password login wiring done yet either, and see where a
   post-callback token hydration would plug into the store.

5. **`mocks/handlers.ts`** — MSW handlers; git status shows this file
   already modified. Need to see which account endpoints are mocked
   today and whether `/auth/google/redirect` / `/auth/google/callback`
   have any scaffolding yet.

6. **`app/verify-email/page.tsx`** as precedent — the closest existing
   example in this codebase of a route that exists purely to land a
   redirect/link flow rather than appearing as its own page-map.md row
   (directly relevant to spec Assumption B, the undecided callback
   landing route). Looked at last, as a targeted comparison once the
   general page/route/data-layer conventions from areas 1–5 are clear.

Order rationale: shell/pages first (establishes the page-consolidation
baseline and shows what, if anything, already exists for the Google
button), then components (what task 01 already built and its
composition pattern), then the API client layer (data-fetching
conventions this new full-redirect flow has to fit into or diverge
from), then hooks/store (state/token layer), then mocks (usually
trails the API layer), and the `verify-email` non-page precedent last
since it's a targeted comparison rather than a load-bearing area.

Waiting for go-ahead before Stage 2.
