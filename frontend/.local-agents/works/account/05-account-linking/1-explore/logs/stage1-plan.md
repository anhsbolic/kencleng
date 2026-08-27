# Stage 1 — Plan Announcement

Feature: `docs/spec/1-account/features/05-account-linking.md`
(frontend surface: `POST /account/security/google/unlink`,
`POST /account/security/set-password`)

## Docs read (per task instructions)

- `docs/spec/README.md` — doc-type conventions, not feature-specific.
- `docs/spec/1-account/features/05-account-linking.md` — full read.
  Two endpoints, `set-password` branches server-side into Branch 1
  (Google-only caller, adds a new unverified `email_password` identity,
  anti-enumeration `202`) and Branch 2 (caller already has
  `email_password`, immediate change-password, `200`/`401`/`422`).
  `unlink` requires re-auth (`password`) and INV-account-12 (remaining
  identity must be *verified*, not just present) — two distinct `409`
  messages depending on which guard fails.
- `docs/ui-ux/page-map.md` — Donatur table: `/dashboard/security` (Form,
  multi-section) — "Enable/disable MFA (QR scan + confirm code),
  view/regenerate backup codes, link/unlink Google identity. Google-only
  users also see 'Atur Password' here." This single row covers **both**
  this task (#5, linking) and Account Task #6 (MFA, not yet built) —
  flagged for the page-consolidation check in Stage 2. No other
  persona table or Cross-Cutting UI Elements row references
  linking/set-password — this is the only frontend surface, no separate
  non-page surface (unlike email verification's "click link" row) since
  set-password's own verification step reuses the existing
  `/verify-email` page from Task #1, not a new route.
- `docs/ui-ux/patterns.md` — Pattern 3 (Form Page) + the explicit "Form
  (multi-section)" variant implied by page-map's `/dashboard/security`
  row (patterns.md itself doesn't name a distinct "multi-section"
  sub-pattern the way it names "Revisable Submission" — it's just
  Pattern 3 with multiple independent sections, closest to how
  `/dashboard/organization/[id]` composes role-gated sections within one
  Detail page). Cross-Pattern State Conventions (§B) apply: request-level
  banner vs field-level error separation is directly relevant given the
  spec's anti-enumeration generic-`202` design and the two distinct `409`
  unlink messages.
- `docs/ui-ux/prototype-reference.md` — `/dashboard/security` is **Tier
  2** (no dedicated prototype), row explicitly lists it under the "Form"
  pattern group, closest precedent named as `/dashboard/campaign/new`
  (dashboard forms) or `/login` (auth modal/mobile split — less relevant
  here since this is a Dashboard page, not Auth Shell). Checked
  `docs/design-reference/` directory listing directly: 10 files +
  README, none named `security`/`account`/`mfa`-anything. No Tier-1
  prototype or component sheet exists for this page despite
  `prototype-reference.md` mentioning "two non-page reference sheets"
  (Component & Layout sheet) — that sheet is not present as a file in
  the directory at all, only referenced in prose (flagged as a possible
  doc/reality inconsistency, not resolved here).
- `docs/ui-ux/design-guidelines.md` — full token read. No
  linking/security-specific tokens; standard Form Page, Input, Button,
  Banner tokens apply. Destructive button token (`error-500`/`error-700`)
  is the relevant one for the "Lepas Tautan Google" action.

## Shell identification

This page sits under the **Dashboard Shell**
(`app/(dashboard)/layout.tsx` + `app/(dashboard)/_components/
dashboard-shell-client.tsx`), already fully built in Phase 0
(`frontend/.agents/docs/phase0-shared-infra.md` Steps 3b/5) — top-nav
desktop, top-bar+hamburger mobile, nav items role-filtered via
`useHasRole`. The Shell's nav list (`app/(dashboard)/_components/
nav-items.ts`) already includes a "Keamanan" item pointing at
`/dashboard/security` — Phase 0 explicitly scoped this in
(`phase0-shared-infra.md`'s revision note: "Task #5, Account Linking,
maps to `/dashboard/security`, already in this phase's Dashboard Shell
nav scope").

`app/(dashboard)/dashboard/security/page.tsx` **already exists as a
route** (confirmed via directory listing) but is a 12-line explicit
placeholder: "real form is Account Task #5's scope
(docs/spec/1-account/tasks.md), not this playbook
(phase0-shared-infra.md). Exists so the Dashboard Shell has a route to
render against during this phase's verification." Same pattern as
`/forgot-password`/`/reset-password` in Task #4 — scaffolded
specifically to be filled in by this task, not prior real work to
preserve, and not a page-consolidation conflict on its own (see Area 3
for the MFA-overlap nuance).

## Areas to explore, and order

1. **API client + schema layer** (`lib/api/account.ts`,
   `lib/api/schema.d.ts`, `lib/api/client.ts`, `api/openapi.yaml`).
   First, because everything downstream (hooks, forms) depends on
   whether the generated types/fetch functions for these two endpoints
   already exist or need to be added — a grep already shows the two
   paths present in `schema.d.ts` (contradicting the feature spec's own
   "References" section claim that `openapi.yaml` still needs a schema
   update for these — worth confirming fully).
2. **Hooks layer** (`lib/hooks/use-login.ts`, `use-register.ts`,
   `use-forgot-password.ts`, `use-reset-password.ts`,
   `use-account-me.ts` as precedent; confirming no
   `use-set-password.ts`/`use-unlink-google.ts` exist yet). Second,
   depends on Area 1's API shape.
3. **Dashboard Shell + `/dashboard/security` page + page-consolidation
   check** (`app/(dashboard)/layout.tsx`,
   `_components/dashboard-shell-client.tsx`, `_components/nav-items.ts`,
   `dashboard/security/page.tsx`, plus `docs/spec/1-account/tasks.md`'s
   Group B/Serial-S1 grouping for the MFA overlap). Third — confirms
   the page is genuinely just the placeholder and surfaces the
   MFA-vs-linking page-consolidation question explicitly.
4. **Feature components** (`components/features/account/login-form.tsx`,
   `forgot-password-form.tsx`, `reset-password-form.tsx`,
   `google-auth-button.tsx`, `resend-verification-control.tsx` as
   precedent for RHF+zod+Banner conventions this task's set-password/
   unlink UI must follow). Fourth, depends on Areas 1-2.
5. **`components/ui/` primitives inventory** (`button`, `input`,
   `label`, `banner`, `spinner`, `badge`, `progress-bar` — check what
   exists vs. what a re-auth-gated destructive action + multi-branch
   form might need that isn't there, e.g. no modal/dialog primitive).
   Fifth.
6. **Visual precedent** (`design-reference/campaign-new.html`,
   `login-register.html`). Last — Tier 2 precedent only, informs Stage 3
   visual decisions, doesn't change behavior/data shape.

Order rationale: 1 → 5 is a dependency chain (data contract → hooks →
routing target/consolidation → components → missing primitives); Area 6
is independent, pushed last since it's precedent for visual polish only.

---

Confirmed by user — proceeding to Stage 2.
