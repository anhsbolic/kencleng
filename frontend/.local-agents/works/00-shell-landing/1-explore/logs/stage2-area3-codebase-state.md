# Stage 2 — Area 3: Current frontend codebase state

> Task: 00-shell-landing
> Date: 2026-08-26

## Current state

- `app/(public)/layout.tsx`: exactly the documented stub — pass-through
  `{children}`, comment says build the real shell "when `campaign`
  domain's frontend track starts" (now). `app/(public)/page.tsx`: still
  the untouched `create-next-app` default scaffold page — not a Kencleng
  placeholder, needs full replacement.
- `components/ui/` has exactly 5 primitives: `button.tsx`, `input.tsx`,
  `label.tsx`, `banner.tsx`, `spinner.tsx` (+ co-located `.test.tsx`).
  Confirms Area 1: no `Badge`/`ProgressBar`/`Card`.
- `Button` already matches `design-guidelines.md` exactly: variants
  (primary/secondary/outline/ghost/destructive), sizes (sm=36px `h-9`,
  md=44px `h-11`, lg=52px `h-13`). Directly reusable for Header/Hero CTAs.
- `app/globals.css` (Tailwind v4, CSS-first, `@theme inline`) already
  defines every color/radius/shadow/font-family token from
  `design-guidelines.md`. **Does NOT define any typography-scale tokens**
  (`display`/`h1`–`h4`/`body-lg`/`body`/`body-sm`/`caption`) — nothing
  like `text-h1` exists anywhere in the codebase yet.
- `lib/hooks/use-focus-trap.ts` is a genuinely shared, generic hook (not
  nested under `(dashboard)/_components/`) — implements the Tab-trap +
  focus-in/return-out behavior `phase0-shared-infra.md` Step 5 requires.
  Doc comment explicitly anticipates reuse: "Shared by the Auth Shell
  modal and the Dashboard Shell mobile drawer... any future overlay
  component should reuse this rather than reinventing it." Confirms
  page-map.md's "reuse the drawer focus-management" is a clean hook
  import, not a rewrite.
- Dashboard Shell's actual drawer markup (`DashboardShellClient`) lives in
  `app/(dashboard)/_components/dashboard-shell-client.tsx`, route-group-
  private, not exported for cross-shell reuse — its shape (hamburger with
  `aria-expanded`/`aria-controls`, `role="dialog"`/`aria-modal` panel,
  `useFocusTrap` wiring) is the pattern to mirror structurally, not import
  directly.
- `mocks/handlers.ts` has exactly 2 handlers (`GET /account/me`,
  `GET /notifications/unread-count`) — confirms a new `GET /campaigns`
  handler is needed.
- `lib/api/account.ts` establishes the convention for a new
  `lib/api/campaign.ts`: thin wrapper over `apiFetch`, typed against
  `components["schemas"][...]` from `schema.d.ts`, throws generic `Error`
  on non-OK (never surfaces raw response body) — matches `patterns.md`
  §B's "never render raw backend error text."
- `lib/api/client.ts`'s `apiFetch` handles auth-token attachment +
  401-refresh-retry. `GET /campaigns` is `security: []` (public) — calling
  through `apiFetch` anyway is harmless and keeps the single-fetch-path
  rule (`client.ts`'s own comment: "every domain's fetch function must go
  through `apiFetch`").

## Requirement

Real `(public)/layout.tsx` (Public Shell) + real `(public)/page.tsx`
content, real `components/ui/` primitives (existing + new Badge/
ProgressBar), real `lib/api/campaign.ts` + query hook, new mock handler.

## Gap

1. `(public)/layout.tsx` and `(public)/page.tsx` are both scaffold
   placeholders — 100% new build.
2. `Badge`/`ProgressBar` need building (color/shape tokens already exist
   in `globals.css`, just missing component files), including
   ProgressBar's documented 100%-fill color switch
   (`primary-600` → `success-500`).
3. Typography tokens have never been wired into Tailwind anywhere in this
   codebase — this task is the first needing the full type scale. Bigger
   foundational gap than "one component is missing" (Stage 3).
4. New Public Shell client component (mirroring `DashboardShellClient`'s
   shape, importing `useFocusTrap`) needs writing from scratch.
5. `lib/api/campaign.ts`, a campaign query hook, and a `GET /campaigns`
   MSW handler all need creating, following `account.ts`/
   `use-account-me.ts` conventions.

## Sniffing

- **Risk:** since `globals.css` has zero prior typography-token
  precedent, whatever this task picks becomes the de facto pattern the
  next page-building task copies — worth deciding deliberately.
- **Miscontext risk avoided:** page-map.md's "reuse the Dashboard Shell's
  drawer focus-management" could be misread as "import the Dashboard
  Shell component directly." Confirmed not the case — the reusable unit
  is the `useFocusTrap` hook; Public Shell still needs its own drawer
  markup, just following the same shape.
