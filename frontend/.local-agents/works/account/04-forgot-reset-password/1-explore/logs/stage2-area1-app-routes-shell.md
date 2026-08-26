# Stage 2 — Area 1: App routes + Auth Shell

## Current state

- `app/(auth)/layout.tsx` (15 lines) — thin Server Component, just
  wraps `children` in `<AuthShellClient>`. No per-page logic.
- `app/(auth)/_components/auth-shell-client.tsx` (66 lines) — Client
  Component. Desktop (`md:` breakpoint, `matchMedia` via
  `useSyncExternalStore`) renders a centered modal overlay
  (`role="dialog"`, `aria-modal`, focus-trapped via `useFocusTrap`);
  mobile renders a plain full page (no dialog role, no focus trap).
  The switch is pure CSS (`md:` classes) — only focus-trap
  *activation* branches in JS.
  - **Explicit documented convention in the component's own JSDoc**:
    "Convention for pages rendered inside this shell ... render a
    `<Banner variant="error">` as the first child, before the form —
    this panel is a plain `flex flex-col gap-6`, so whatever a page
    puts first renders above the form with no extra wiring. This is
    the known-issue guard from `prototype-reference.md` (`/login`'s
    auth failure must be a banner, not a field-level error) made
    structurally easy to get right." This directly targets the
    request-level-vs-field-level error separation `patterns.md` §B and
    the feature spec both require (`forgot-password`'s generic `202`
    doesn't need this, but `reset-password`'s `410`/`404`/`422`
    responses do — expired/not-found/already-used are request-level,
    not field-level).
- `app/(auth)/login/page.tsx` (32 lines) and `app/(auth)/register/
  page.tsx` (13 lines) — both real, already-implemented Form pages
  using the same shape: `<div className="flex flex-col gap-6">` →
  heading block (`h1` + muted `p` subtitle) → the actual form
  component (`LoginForm`/`RegisterForm`), which owns its own Google
  button, validation, and submit wiring internally. `LoginPage` also
  wraps a `GoogleCallbackError` in a `<Suspense>` for `useSearchParams`
  (not relevant to forgot/reset — no OAuth callback here — but
  confirms `useSearchParams()` needs a `Suspense` boundary in this
  App Router version, which **is** relevant: `/reset-password?token=...`
  will need to read `token` via `useSearchParams()`).
- `app/(auth)/forgot-password/page.tsx` and `.../reset-password/
  page.tsx` — confirmed stubs (see Stage 1). 12 lines each, no form,
  no imports beyond the default export.
- **Cross-navigation already wired**: `components/features/account/
  login-form.tsx:151` has a real `<Link href="/forgot-password">Lupa
  password?</Link>` next to the password field label, and `login-form
  .test.tsx:44` already asserts this link's `href`. So the login form
  side of the flow is done and just waiting on `/forgot-password` to
  exist as a real page.
- No existing link *to* `/reset-password` anywhere in `app/`/
  `components/`/`lib/` (expected — that link only ever comes from the
  reset email itself, which is backend/notification-owned, not a
  frontend nav element).

## Requirement

- `page-map.md`: `/forgot-password` = Form, "Submit email to request
  password reset"; `/reset-password?token=...` = Form, "Submit new
  password." Both Guest-only.
- `patterns.md` Pattern 3 (Form Page): idle → validating → submitting
  → submit-error (banner) vs field-error (inline) → success
  (toast+redirect for dashboard forms, **inline success state for
  guest-facing forms** — forgot-password is guest-facing, so its
  success state should be inline, not toast+redirect, per this rule).
- `prototype-reference.md`: Tier 2, closest precedent `/login` (Auth
  Shell modal/mobile split) or `/dashboard/campaign/new`. Given both
  pages already live in `app/(auth)/`, `/login`'s shell conventions
  are the operative precedent, not `/dashboard/campaign/new`'s.
- Feature spec `04-forgot-reset-password.md`: forgot-password is
  always `202` generic (no branch the UI should visibly distinguish);
  reset-password has real distinguishable failure branches the UI must
  render differently (`422` field-level password/breach error,
  `410` expired-token request-level, `404` not-found/used-token
  request-level).

## Gap

- Both page files are 12-line placeholders with zero form, zero
  heading-pattern consistency with `login`/`register`'s established
  `flex flex-col gap-6` + heading-block shape. Full rebuild needed for
  both, following the sibling pages' shape.
- `/reset-password` page needs a `Suspense` boundary for
  `useSearchParams()` (to read `token`) — same pattern already used in
  `login/page.tsx` for `GoogleCallbackError`, not yet applied to this
  page since it doesn't exist yet.
- No page-consolidation conflict: both pages are single-purpose to this
  task, already reserved via the placeholder-stub mechanism, and the
  only page-map action items for this feature route to these two pages
  1:1 — no orphaned page-map action without a matching endpoint (both
  endpoints are in the domain's spec) and no orphaned endpoint without
  a page-map row (both `/auth/forgot-password` and `/auth/reset-password`
  have their page-map row).

## Sniffing

- **Misleading signal**: none in this area — the stub pages are
  explicitly self-labeled as placeholders, not something that could be
  mistaken for done work.
- **Miscontext**: none — page-map's "Form" pattern label for both pages
  matches what's actually needed; no mismatch between what the spec
  author assumed and what's there.
- **Risk**: the Auth Shell's own JSDoc convention (banner-as-first-child
  for request-level errors) is *documented* but not yet exercised by
  any real page except implicitly by `login/page.tsx`'s indirect
  `GoogleCallbackError` usage (which is a query-param banner, not a
  submit-error banner) — worth confirming during Stage 3 that
  `reset-password`'s expired/not-found/already-used states actually
  render as the shell's documented first-child banner rather than
  inventing a different error placement.
- **Edge case**: `/reset-password` with no `token` query param at all
  (not just an invalid one) — page-map/patterns don't explicitly
  distinguish "missing" from "invalid" token; `patterns.md`'s
  Status/Tracking pattern (§6, for the *unauthenticated* donation
  status page) explicitly says invalid/missing token must share one
  generic message to avoid confirming/denying existence — worth
  checking in Stage 3 whether that same anti-enumeration posture
  should extend to `/reset-password`'s token, even though `patterns.md`
  doesn't name this page under that pattern.
- **Inconsistency**: none found between docs and code in this area.

Proceeding to Area 2 (API client layer).
