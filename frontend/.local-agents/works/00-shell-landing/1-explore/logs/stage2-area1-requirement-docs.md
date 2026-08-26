# Stage 2 — Area 1: Requirement docs

> Task: 00-shell-landing
> Date: 2026-08-26

## Current state

- `page-map.md` (2026-08-24) resolves two things for `/`:
  1. **Public Shell nav** — top nav: logo, "Beranda" (`/`), "Jelajahi
     Kampanye" (`/campaign`), "Masuk"/"Daftar". Mobile: hamburger drawer
     reusing Dashboard Shell's existing focus-management
     (`phase0-shared-infra.md` Step 5), not a new pattern. Footer
     explicitly deferred, not blocking `/`.
  2. **"Highlighted campaigns"** (footnote 1, §1): `GET /campaigns` has no
     sort/featured param. Mock returns a fixed fixture set, no invented
     `featured` flag/sort logic. What "highlighted" means product-wise is
     Open Item #4 — deliberately unresolved.
- `prototype-reference.md`: `/` is Tier 1 ("Landing (one-off, not a formal
  pattern)", "Public header variant, not Dashboard Shell") — near-
  authoritative, deviate only where a feature spec requires different
  behavior. Known Issues #2 (campaign card upload-dropzone placeholder)
  and #3 (typography drift) apply directly; #1 (login field error) does
  not.
- `patterns.md`: no dedicated pattern for `/` — it's the one-off "Landing"
  case. Closest applicable piece: List/Browse card-grid shape (§A.1) +
  cross-pattern state conventions (§B) for the campaign section.
- `design-guidelines.md`: confirms task's cited drift — `h1` = 30px
  (1.875rem), `display` = 36px (2.25rem), "Landing/hero only." Also
  relevant, not in task args: Card `radius-lg` (16px)/`shadow-sm`, Button
  size tokens (md=44px default), ProgressBar 100%-fill color switch
  (`primary-600` → `success-500`), Badge pattern for org verification.
- `kencleng-frontend-tech-stack.md` + `phase0-shared-infra.md`:
  `(public)/layout.tsx` was deliberately left a pass-through stub at
  Phase 0, with an explicit forward pointer — "build it when `campaign`
  domain's frontend track starts." Per the Incremental Growth Rule, a
  Shell warrants its own short one-time playbook (mirroring Step 5's
  shape), not folding into a feature task. No such playbook exists yet —
  only `scaffold-frontend.md` is indexed in `.agents/docs/README.md`.

## Requirement

Build `/` as Tier 1 near-final: real Public Shell nav (desktop top nav +
mobile hamburger drawer reusing Dashboard Shell's focus-management),
highlighted-campaigns section from `GET /campaigns` via fixed-fixture
mock, `design-guidelines.md`'s actual type scale (not the prototype's),
campaign card photo placeholder as plain read-only (no upload affordance).

## Gap

1. Public Shell nav doesn't exist in any form — building from scratch,
   not modifying a partial implementation.
2. No `Badge`, `ProgressBar`, or `Card` primitive exists yet
   (`phase0-shared-infra.md` Step 2 scoped `components/ui/` to
   Button/Input/Label/Banner/Spinner only; Badge/ProgressBar explicitly
   named "deferred, not forgotten").
3. No campaign API fetch/hook layer exists (`lib/api/` has only
   account.ts/client.ts/notification.ts/schema.d.ts; no campaign hook).
4. No dedicated Landing pattern to build against literally — nearest-
   precedent + Tier 1 prototype is the governing spec.

## Sniffing

- **Inconsistency:** `patterns.md`'s Detail pattern says the org
  verification badge applies to "any page displaying campaign or
  organization info to a public/donor viewer" — read literally this could
  include `/`'s cards, but every "Used by" example for that pattern is a
  Detail page, not a List/Browse card. Needs a Stage 3 call.
- **Miscontext check (passed):** both cited page-map.md decisions
  (Public Shell nav, highlighted-campaigns mock scope) are genuinely
  resolved as of 2026-08-24, not misread from still-open items.
- **Risk:** reusing Dashboard Shell's drawer focus-management needs Area 3
  to confirm what's actually reusable — a hook vs. route-group-local
  markup.
- **Edge case:** page-map.md's mock note doesn't say how many fixture
  campaigns, or what happens if the fixture array is ever empty (List/
  Browse empty-state convention would apply) — a copy/behavior decision,
  not assumed.
