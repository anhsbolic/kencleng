# Playbook — Public Shell (`(public)/layout.tsx`)

> File: `frontend/.agents/docs/scaffold-public-shell.md`
> Scope: run once, when `campaign` domain's frontend track starts
> building `/` (this is that moment — see `techplan.md` for
> `00-shell-landing`). Builds the Public Shell nav that `/`,
> `/campaign`, and `/campaign/[id]` all sit inside.
> Tier: infrastructure, not a feature — but touches real UI (nav,
> focus management), so it goes through the same one-time-playbook
> treatment `phase0-shared-infra.md` gave `(auth)`/`(dashboard)`,
> per that file's Incremental Growth Rule (Shells are the one
> explicit exception to "just build it as part of the task").

## Why this exists, and why it's scoped this small

`(public)/layout.tsx` was deliberately left a pass-through stub at
Phase 0 (`phase0-shared-infra.md` Step 1) with an explicit note to
build the real shell once `campaign` domain's frontend track starts —
not because it needed a backend dependency, but because `/`, `/campaign`,
`/campaign/[id]` belong to `campaign` domain, not `account`. This
playbook builds exactly the nav `page-map.md`'s 2026-08-24 Public Shell
decision calls for — not a general-purpose marketing-site shell.

## Step 1 — Nav item list

Guest-facing, no role-gating (unlike Dashboard Shell's `useHasRole`
filtering — there's no session to check here):

| Label | Target | Notes |
|---|---|---|
| Kencleng (logo/`Mark`) | `/` | Always visible |
| Beranda | `/` | |
| Jelajahi Kampanye | `#kampanye` | In-page anchor, not `/campaign` — that route doesn't exist yet. **[RESOLVED 2026-08-26]** Swap to a real `/campaign` href in one line once that task ships; no other change needed. |
| Masuk | `/login` | Existing route (`(auth)` Shell) |
| Daftar | `/register` | Existing route (`(auth)` Shell) |

Footer is explicitly **not** part of this phase — `page-map.md` defers
it, not required for `/` to function.

## Step 2 — Desktop layout

Top nav (`Mark` left, nav links + Masuk/Daftar right), rendered directly
in `(public)/layout.tsx` — no client component needed for the desktop
case, it's static markup with no open/close state.

## Step 3 — Mobile layout: hamburger + drawer

Per `page-map.md`: "reusing the same open/close focus-management
behavior already built for Dashboard Shell's mobile drawer" — this
means reusing `lib/hooks/use-focus-trap.ts` (already a generic, shared
hook, not route-group-private), **not** importing
`DashboardShellClient` itself (that component's markup lives in
`(dashboard)/_components/`, private to that route group).

Build `app/(public)/_components/public-shell-client.tsx` mirroring
`DashboardShellClient`'s shape:

- Hamburger button: `aria-expanded` reflects open/closed, `aria-controls`
  points at the drawer's `id`.
- Drawer: `role="dialog"`, `aria-modal="true"`.
- `useFocusTrap({ active: open, containerRef, onEscape: () => setOpen(false) })`
  — on open, focus moves to the drawer's first nav item; on close
  (Escape or a nav item navigating away), focus returns to the
  hamburger button; Tab/Shift+Tab cycles within the drawer only.

This is the same hook Auth Shell's modal and Dashboard Shell's drawer
already use — a third consumer confirms it's genuinely shared
infrastructure, not something that happened to be reusable twice by
coincidence.

## Step 4 — Verify

```bash
npm run dev      # / renders the real nav at both desktop and mobile widths
npm run lint
npm run test      # keyboard-only open/close/focus-return for the drawer,
                   # same RTL role/keyboard query approach as the Dashboard
                   # Shell drawer test
npm run build
```

## Step 5 — Human checkpoint

- [ ] Desktop nav and mobile hamburger+drawer actually switch at the
      intended breakpoint, checked at both widths
- [ ] Hamburger `aria-expanded` toggles correctly; focus lands on the
      drawer's first nav item on open and returns to the hamburger on
      close; Tab cycles within the drawer only while open — checked
      with keyboard only, not just visually
- [ ] "Jelajahi Kampanye" scrolls to `#kampanye`, does not attempt to
      navigate to `/campaign`
- [ ] No `campaign`-domain business logic snuck into this phase — the
      Shell stays pure nav/layout, same discipline `phase0-shared-infra.md`
      Step 7 held Account's primitives to

## Related docs

- `docs/ui-ux/page-map.md` — Public Shell nav decision (2026-08-24)
- `frontend/.agents/docs/phase0-shared-infra.md` — Step 5's Dashboard
  Shell drawer, the precedent this mirrors; Incremental Growth Rule
  (why Shells get their own playbook)
- `lib/hooks/use-focus-trap.ts` — the shared focus-management hook
- `.local-agents/works/00-shell-landing/2-plan/techplan.md` — the
  techplan this playbook was written for
- `../../harscode-workspace/best-practices/react/accessibility-fundamentals.md`
  — the focus-management pattern this playbook implements
