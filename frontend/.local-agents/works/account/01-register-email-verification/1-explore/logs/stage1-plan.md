# Stage 1 — Plan Announcement: account/01-register-email-verification (frontend)

> Per `workflow/1-exploration/guidelines.md` — a hard-stop announcement,
> made before any deep implementation reading, so a wrong area/order
> costs nothing to redirect. Recorded here after the fact for the
> record; it was originally posted directly in conversation and
> approved before Stage 2 began.

## Context established before proposing areas

Read `frontend/AGENTS.md` and `frontend/.agents/docs/README.md` first,
just enough to place this task:

- This task's primary page (`/register`) lives under the **Auth Shell**
  — `app/(auth)/layout.tsx` → `AuthShellClient` (desktop modal / mobile
  full-page split, per `patterns.md`'s Form/Auth sub-variant). The Auth
  Shell already exists (built in `phase0-shared-infra.md` Step 3) and
  wraps four routes: `login`, `register`, `forgot-password`,
  `reset-password`.
- `app/(auth)/register/page.tsx` **already exists as a route** but is an
  explicit placeholder (comment: *"real form is Account Task #1's
  scope... Exists so the Auth Shell has a route to render against
  during this phase's verification"*) — this task is expected to
  replace it, not create the route from scratch.
- Per `page-map.md`, this feature surfaces in **two places**: `/register`
  (Guest, Form pattern) **and** a non-page surface — "Email verification
  link (from email, not a full page)" under Donatur, which sets
  `AuthIdentity.verified_at` on click. That second surface needs its
  own landing route even though it isn't a "page" in the product sense
  — not yet confirmed whether one exists in the codebase.
- `design-reference/` has a `login-register.html` file —
  `prototype-reference.md`'s Tier 1 table only names `/login`
  explicitly, but the combined filename plausibly covers `/register`
  too, which would upgrade it from Tier 2 (template-only) to Tier 1
  (near-final draft) for this task specifically. Needs verification via
  `design-reference-usage.md`'s extraction process, not a raw `cat`.
- A near-identical prior exploration attempt exists at
  `.local-agents/works/01-account/1-explore/` (empty logs) — different
  path than this run's target
  (`works/account/01-register-email-verification/`) — a naming
  inconsistency to note, not resolve.

## Areas to explore, in order, and why

1. **Auth Shell + existing `(auth)` routes** (`auth-shell-client.tsx`,
   `login/page.tsx`, `forgot-password/page.tsx`,
   `reset-password/page.tsx`, `register/page.tsx`) — foundational:
   establishes exactly what's scaffolded vs. placeholder, and gives
   sibling-form precedent within the same Shell before looking anywhere
   else.
2. **The non-page email-verification-link surface** — search the
   codebase for any existing route consuming a verification token (e.g.
   something under `(auth)/verify-email` or similar). Determines
   whether that second surface from `page-map.md` is missing entirely
   or partially built.
3. **`design-reference/login-register.html`** (via
   `design-reference-usage.md`'s extraction steps) — check whether it
   actually depicts a `/register` state, which would change this
   page's tier assignment from what `prototype-reference.md`'s table
   alone implies.
4. **API/data layer** (`lib/api/account.ts`, `lib/api/schema.d.ts`,
   `lib/hooks/`, `mocks/handlers.ts`) — check whether `register` /
   `verify-email` / `verify-email/resend` are already
   generated/wrapped/mocked, or need to be added.
5. **`lib/stores/auth-store.ts`** — last, since it depends on what #4
   finds: check what session state exists today and how it should (or
   shouldn't) interact with a register flow that per spec returns no
   token/user_id and doesn't log the user in.

Order is dependency-based: 1 tells me what's already there to build
on/replace, 2 surfaces the easy-to-miss second page-map row, 3 affects
how much visual authority the design reference gets, and 4→5 are the
data-layer areas that 1–3 determine the actual size of.

**Outcome:** approved as-is; no redirection requested. Stage 2 proceeded
through all five areas in this order.
