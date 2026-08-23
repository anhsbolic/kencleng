# Playbook — Frontend Phase 0: Shared Infra Before Account Domain

> File: `frontend/.agents/docs/phase0-shared-infra.md`
> Scope: run once, after `scaffold-frontend.md` and before Account
> domain's Task #1 (`docs/spec/1-account/tasks.md`). Builds the
> minimum shared UI/routing infrastructure Account's first tasks
> actually need — not a general-purpose component library built
> speculatively ahead of demonstrated need.
> Tier: infrastructure, not a feature — no domain invariant applies,
> but this phase does touch real UI (unlike `scaffold-frontend.md`,
> which is pure plumbing), so it still goes through the workflow §14
> gate sequence per component, plus a human checkpoint (Step 7).
> **Two audiences**: Steps 0-7 are a one-time execution playbook.
> The "Incremental Growth Rule" section at the bottom is durable
> policy — consult it again whenever a *later* domain task thinks it
> needs to add shared infra.

## Why this exists, and why it's scoped this small

Backend domains build cleanly in isolation — `internal/domain/<n>/`
needs nothing outside itself to start. Frontend doesn't have that
luxury: even the very first page (`/register`) needs a Button, an
Input, somewhere to render a form-level error, and a route-group
shell to sit inside. This playbook builds *exactly that much* for
Account domain's frontend surface — not a general-purpose component
library, not every cross-cutting component every domain will
eventually want.

**Correction from an earlier version of this doc**: this playbook
originally scoped itself to only the backend's Serial-group-S1 task
order (`docs/spec/1-account/tasks.md` #1-4), deferring Dashboard Shell
on the reasoning that it wasn't needed until backend Task #5. That
reasoning was wrong — per `scaffold-frontend.md`'s **Mock-First
Development Workflow**, frontend work isn't gated by backend task
completion order at all; every Account-domain page can be built and
verified against MSW-mocked responses regardless of which backend
task has actually shipped. The real scoping question is "does
*Account domain* need this," not "has backend built the endpoint for
this yet" — and Account domain's own page inventory
(`docs/ui-ux/page-map.md`, Donatur section) includes
`/dashboard/profile`, `/dashboard/security`, `/dashboard/notifications`
alongside the Guest auth pages. Dashboard Shell is therefore in scope
for this phase, not deferred. See "Incremental Growth Rule" for what
*is* still deferred, and why.

## Step 0 — Verify prerequisite

`scaffold-frontend.md` must already be complete (`components/`,
`lib/`, `mocks/` exist, deps installed, `lib/api/schema.d.ts`
generated). If not, stop — that playbook runs first, this one
doesn't substitute for it.

## Step 1 — Route groups

Restructure `app/` into three route groups (resolved decision):

```
app/
├── (public)/           # Guest content — /, /campaign, etc.
│   └── layout.tsx       # Public Shell — STUBBED this phase, see below
├── (auth)/              # /login, /register, /forgot-password, /reset-password
│   └── layout.tsx        # Auth Shell — BUILT this phase (Step 3)
├── (dashboard)/         # everything under /dashboard/*
│   └── layout.tsx        # Dashboard Shell — BUILT this phase (Step 3b)
└── layout.tsx            # root layout — providers only, from scaffold-frontend.md
```

**`(auth)` and `(dashboard)` get real shells this phase. `(public)`
stays stubbed** (pass-through `{children}`, no nav) — and this is
still deliberate, but for a different reason than the backend-task
one this doc used to (wrongly) lean on: `/`, `/campaign`,
`/campaign/[id]` belong to the **`campaign` domain**, not `account`
(`docs/ui-ux/page-map.md`'s "Guest" section spans multiple domains'
pages; the Auth and Dashboard pages there are Account's, the browse/
detail pages are `campaign`'s). Deferring `(public)` is about domain
ownership, not backend readiness — build it when `campaign` domain's
frontend track starts.

Move the existing `app/page.tsx` into `(public)/page.tsx` — it's
Guest content, correctly belongs there even unbuilt.

## Step 2 — Minimum `components/ui/` primitives

Scoped to exactly what Account Tasks #1–#4 need (`docs/spec/1-account/
tasks.md`, Serial group S1) — all four are Form-pattern pages
(`docs/ui-ux/patterns.md` Pattern 3), nothing more exotic than a form
and a request-level error banner:

| Primitive | Why it's needed now | NOT built now |
|---|---|---|
| `Button` | Every form needs a submit action | Icon-button/ghost variants beyond what S1 forms use |
| `Input` | Text/email/password fields | Select, file upload — needed by `campaign`/`organization`, not Account S1 |
| `Label` | Form field labels (`body-sm`) | — |
| `Banner` | Request-level success/error — Task #1's anti-enumeration design means register *always* returns a generic banner message, not optional | Toast (dashboard-only pattern, no dashboard yet) |
| `Spinner` | Inline, submit-button loading state (`patterns.md` Form Page "Submitting" row) | Skeleton loaders (List/Detail pattern — no S1 page uses them) |

Build against `design-guidelines.md`'s Component Tokens section —
sizes/colors/radius already resolved, no new visual decisions here.
One file per primitive under `components/ui/`; skip a barrel
`index.ts` for now (five files isn't worth the indirection yet).

**Explicitly deferred, not forgotten**: `Badge`, `ProgressBar`,
`MaskedField`, `SecureUploadNote`, `CurationDecisionPanel`,
notification badge/center, Toast. None are used by any Account S1
task. See Incremental Growth Rule for when/how they get added.

## Step 3 — Auth Shell (`(auth)/layout.tsx`)

Per `docs/ui-ux/prototype-reference.md`'s `/login` Tier 1 entry:
**desktop = modal overlay, mobile = full page.**

- Responsive switch via Tailwind breakpoint (CSS, not a JS
  `resize`-listener — avoids layout flash, and CSS can already do
  this without added complexity):
  - Desktop: centered panel (`radius-xl`, `shadow-lg`) over a dimmed
    backdrop.
  - Mobile: full-page, no modal chrome.
- **Known-issue guard**: `prototype-reference.md` flags the `/login`
  prototype's error rendering as incorrectly field-level instead of
  banner-level. This shell must leave a banner slot *above* the form
  — the actual wiring happens in Account Task #1/#3's own session,
  but the shell's layout is what makes doing it right the easy path.

## Step 4 — Minimal routing/auth plumbing

Built here, ahead of Step 5, because Dashboard Shell's nav filtering
depends on the hooks defined in this step.

- **`middleware.ts`** — coarse check only (resolved decision,
  `kencleng-frontend-tech-stack.md`): if a request matches
  `/dashboard/*` with no session indicator, redirect to `/login`.
  Cheap to write now even with zero dashboard pages existing — every
  later task inherits it, same reasoning as `scaffold-backend.md`'s
  `main.go` startup wiring.
- **`lib/stores/auth-store.ts`** — Zustand, shape only:
  `{ accessToken: string | null, setAccessToken, clearAccessToken }`.
  No login logic here (that's Task #3's job) — just the shared shape
  so Tasks #2 and #3 don't each invent their own.
- **`lib/types/roles.ts`** — the two role type unions the hooks below
  need: `GlobalRole = 'donatur' | 'kurator' | 'admin'`,
  `OrgRoleLevel = 'owner' | 'staff'`.
- **`lib/hooks/use-has-role.ts`**, **`lib/hooks/use-has-org-role.ts`**
  — per the role-gating decision:
  `useHasRole(roles: GlobalRole[]): boolean` (OR logic),
  `useHasOrgRole(orgId: string, levels: OrgRoleLevel[]): boolean`.
  Both read from whatever `GET /account/me` eventually returns
  (Task #7) — until then, safe default `false` rather than throwing,
  since nothing calls them yet. **Step 5's Dashboard Shell nav is the
  first real consumer of `useHasRole`** — this isn't speculative
  plumbing anymore, it has an immediate caller.
- **`components/shared/require-role.tsx`**,
  **`components/shared/require-org-role.tsx`** — thin wrappers over
  the two hooks, per the hybrid decision. No page uses them yet
  (Account S1 has no role-gated *page content*, only role-gated *nav
  items*, which Step 5 handles via the hook directly, not the
  wrapper) — scaffolded alongside their hooks so the pair doesn't
  split across two separate playbook runs.

## Step 5 — Dashboard Shell (`(dashboard)/layout.tsx`)

Per resolved decision: **top-nav desktop, top-bar + hamburger mobile**
(`docs/ui-ux/page-map.md`, `patterns.md` Pattern 4).

- Nav item list starts small — only what Account domain's own pages
  need right now: "Profil" (`/dashboard/profile`), "Keamanan"
  (`/dashboard/security`), "Notifikasi" (`/dashboard/notifications`).
  Other domains' items (kurasi, disbursement, admin panels) get added
  to this same list when *their* frontend tracks start — this is the
  Shell's data, not its structure, so extending it later doesn't mean
  rebuilding the Shell.
- **Each nav item declares its own `roles: GlobalRole[]`**, filtered
  through `useHasRole` (Step 4). E.g.
  `{ label: 'Notifikasi', href: '/dashboard/notifications', roles: ['donatur', 'kurator', 'admin'] }`
  (per `page-map.md`, notifications are available to any logged-in
  user — effectively "no role restriction," expressed as "all
  current roles" rather than a special-cased "no roles required"
  branch, so the filtering logic stays uniform).
- **Notification badge — placeholder data, not placeholder UI.** Per
  `page-map.md`'s Cross-Cutting UI Elements table, a persistent
  unread-count badge belongs in the header for any logged-in user.
  Mock it against the real contract now (`GET /notifications/
  unread-count`, per `api/openapi/notification.yaml`) via
  `mocks/handlers.ts` — this *is* the mock-first workflow in practice:
  `notification` domain hasn't started as a backend track yet, but its
  OpenAPI contract already exists, so there's no reason the badge
  can't be built and visually correct today. Swap the mock handler for
  the real endpoint whenever `notification` domain's backend actually
  ships (per `scaffold-frontend.md`'s Mock-First Development
  Workflow, step 3 "Integrate real") — no component code changes
  needed when that happens.
- Mobile hamburger reveals the same nav item list in a drawer/sheet —
  same filtered list, not a second copy.

## Step 6 — Verify

```bash
npm run dev      # (auth) and (dashboard) routes render, even with placeholder pages
npm run lint
npm run test       # component tests for the 5 ui/ primitives + nav role-filtering
npm run build
```

## Step 7 — Human checkpoint

- [ ] No Account-domain business logic snuck into this phase —
      `Button`/`Input`/`Banner`/`Spinner`/`Label` stay pure
      presentation, no awareness of register/login specifics
- [ ] `(public)` is genuinely stubbed, not partially built —
      half-building it invites scope creep into a phase meant to stay
      minimal (Dashboard and Auth Shells, in contrast, are meant to be
      fully built this phase — don't under-build those to match)
- [ ] Auth Shell's modal-vs-page split actually switches at the
      intended breakpoint, checked at both widths
- [ ] Dashboard Shell nav actually hides/shows items per role — check
      with at least two different `GlobalRole` combinations, not just
      the default logged-in-donatur case
- [ ] Notification badge is wired to the mocked
      `GET /notifications/unread-count` shape from `schema.d.ts`, not
      a hardcoded number

---

## Incremental Growth Rule — durable, read again for every later phase

**This section outlives Steps 0-7 above.** Every future domain task
should check this before assuming it needs to add to
`components/ui/`, build a new Shell, or otherwise touch shared
frontend infra — this is the answer to "how does frontend grow
without re-running a Phase-0-sized playbook every time."

1. **A domain task builds shared infra only when *that task*
   concretely needs it** — not speculatively, not because "some later
   domain will probably want this too." If a task needs a `Select`
   primitive that doesn't exist yet, that task's own session adds
   `components/ui/select.tsx` as part of its own scope. It does not
   become a new "Phase 0.5" playbook.
2. **Exception: Shells.** A Shell (`(public)`, or any future
   route-group-wide layout) is substantial enough — layout, nav
   structure — that it gets its own short one-time playbook when its
   triggering domain starts (mirroring this file's Step 5 shape for
   `(dashboard)`), rather than being folded silently into that
   domain's first feature task. `(public)` specifically: build it
   when `campaign` domain's frontend track starts, following this
   same pattern. Everything smaller (an individual `components/ui/`
   primitive, a one-off `components/shared/` component like
   `MaskedField`, a new Dashboard Shell nav item) follows rule 1
   instead: just build it as part of the task, no separate playbook —
   this is exactly how Dashboard Shell's nav list itself is meant to
   grow (Step 5's list is Account-only on purpose; `organization`,
   `campaign`, `kurasi`, `admin` items get added by *those* domains'
   own tasks, not by reopening this playbook).
3. **Don't retrofit already-shipped pages** when a later domain adds
   something an earlier page *could* have used. If `campaign`
   domain's `Badge` would've made an earlier Account page nicer,
   that's not a reason to reopen it — only touch a shipped page for a
   real bug or a genuinely duplicated implementation, not for
   retroactive consistency.
4. **Unsure whether something is "shared" or `components/features/
   <domain>/`-scoped?** Default to `features/`. Promoting to
   `components/shared/` after a *second* domain independently needs
   the same thing is cheaper and safer than guessing "shared"
   prematurely and building the wrong abstraction for a need that
   hasn't shown up twice yet.

## Related docs

- `frontend/.agents/docs/scaffold-frontend.md` — runs before this
  playbook; provides the folder/provider/token foundation this phase
  builds on, and defines the **Mock-First Development Workflow**
  (durable section) that this phase's Dashboard Shell notification
  badge is the first concrete example of.
- `docs/spec/1-account/tasks.md` — the task list Step 2's primitive
  set is derived from.
- `docs/ui-ux/page-map.md` — Account domain's own page inventory
  (Donatur section) this phase's Dashboard Shell scope is derived
  from, and the Cross-Cutting UI Elements table the notification
  badge comes from.
- `docs/ui-ux/patterns.md` — Form pattern (Pattern 3), Dashboard/
  Summary pattern (Pattern 4), and state conventions (§B) Steps 2 and
  5 must support.
- `docs/ui-ux/prototype-reference.md` — Auth Shell's known issue
  (login error-banner placement).
- `docs/project/kencleng-frontend-tech-stack.md` — route-group and
  `middleware.ts` resolved decisions this phase implements.
- `api/openapi/notification.yaml` — `GET /notifications/unread-count`,
  the contract the Dashboard Shell's notification badge mocks against.