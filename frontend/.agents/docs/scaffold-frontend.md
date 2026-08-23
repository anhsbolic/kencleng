# Playbook — Frontend Scaffold

> File: `frontend/.agents/docs/scaffold-frontend.md`
> Scope: one-time, run once before Task #1 of any domain's frontend
> track (see `docs/kencleng-agentic-workflow.md` §14) if `frontend/`
> still looks like a fresh `create-next-app` output — no `components/`,
> no `lib/api/`, no `zustand`/`@tanstack/react-query` in
> `package.json`. If those already exist, this playbook doesn't apply —
> stop and say so instead of re-running it.
> Tier: infrastructure, not a feature — no domain invariant or threat
> model applies. Still goes through a human checkpoint before merge
> (Step 8) because every later frontend task inherits this wiring.

## Step 0 — Verify environment values first

Confirm these are true in the actual repo state before installing
anything — easy for docs to drift from reality on a fast-moving project:

- [ ] `package.json` currently has only `next`, `react`, `react-dom`,
      `tailwindcss`, `eslint` — if `zustand` or `@tanstack/react-query`
      are already present, stop, this playbook has likely already run.
- [ ] Frontend dev port is `3000` (`Caddyfile` proxies everything
      except `/api/*` to `host.containers.internal:3000` —
      `backend/AGENTS.md` §5 known-gap note also applies here: `/api/*`
      routing through Caddy is broken until a root-level session fixes
      `handle` → `handle_path`; standalone `npm run dev` on `:3000`
      works for anything that doesn't need same-origin auth cookies).
- [ ] No `NEXT_PUBLIC_API_*` env var exists anywhere in the repo — this
      is expected, not a gap. The resolved same-origin decision
      (`kencleng-frontend-tech-stack.md` Open Items) means fetch calls
      go to relative paths (`/api/...`), not an absolute base URL — so
      no frontend `.env` is needed for this scaffold step. Only add one
      later if a genuine need shows up (e.g. a build-time constant),
      not preemptively here.

## Step 1 — Folder structure

Create the skeleton exactly as documented in `frontend/AGENTS.md` §1,
expanded per `docs/project/kencleng-frontend-tech-stack.md`'s
structural principles (`components/features/` vs `components/ui/` vs
`components/shared/`, `lib/api/` vs `lib/hooks/`):

```
frontend/
├── app/                     # already exists — App Router routes/layouts
├── components/
│   ├── ui/                  # design-system primitives, no business awareness
│   ├── features/            # domain-specific, e.g. components/features/donation/
│   └── shared/               # cross-domain non-primitive, e.g. MaskedField
├── lib/
│   ├── api/                 # hand-written fetch functions + generated schema.d.ts
│   ├── hooks/                # TanStack Query hooks wrapping lib/api/
│   └── stores/               # Zustand stores, split per domain (not one giant store)
├── mocks/                    # MSW request handlers for tests (Step 6)
└── public/
    ├── manifest.json         # PWA manifest (Step 5)
    └── sw.js                 # hand-written service worker (Step 5)
```

Each folder gets a placeholder `README.md` or `.gitkeep` only if empty
— don't scaffold empty `index.ts` barrel files with nothing to export.

## Step 2 — Install dependencies

```bash
npm install zustand @tanstack/react-query react-hook-form zod @hookform/resolvers
npm install -D vitest @testing-library/react @testing-library/jest-dom jsdom msw openapi-typescript
```

Notes on choices already resolved elsewhere (don't re-litigate these
here):

- **No HTTP client library** (`axios`, `ky`, etc.) — native `fetch`,
  per `kencleng-frontend-tech-stack.md` Testing section ("real `fetch`
  calls, real TanStack Query hooks").
- **`@hookform/resolvers`** is the standard `react-hook-form` + `zod`
  glue (`zodResolver`) — not called out explicitly in the tech-stack
  doc but a direct, low-complexity consequence of the two decisions
  already made together.
- **No PWA framework** (`next-pwa`, `@ducanh2912/next-pwa`, Workbox) —
  `README.md` / `frontend/README.md` both say "manual PWA setup"
  explicitly. See Step 5.

## Step 3 — Provider wiring (`app/layout.tsx`)

Wire, in this order, wrapping `{children}`:

1. TanStack Query — a `QueryClientProvider` with a `QueryClient`
   instance created once (module scope or `useState` inside a small
   client-component wrapper — `QueryClientProvider` itself must live in
   a Client Component since `QueryClient` isn't serializable across the
   Server/Client boundary).
2. Any Zustand store does **not** need a Provider (Zustand is
   hook-based, no context wrapper) — skip this layer entirely, just
   import stores directly where needed. Don't add a Provider for it out
   of habit carried over from Context-based state libraries.

Keep `app/layout.tsx` itself minimal — the actual `QueryClientProvider`
wrapper belongs in its own small Client Component
(e.g. `components/providers/query-provider.tsx`) so `layout.tsx` can
stay a Server Component otherwise.

## Step 4 — Design tokens

Follow `docs/ui-ux/design-guidelines.md` for the actual color/type/
radius values, **but note one thing the doc gets stale on**: its
"Implementation Approach" section shows a Tailwind v3-style
`tailwind.config.js` with `theme.extend.colors`. This repo has
`tailwindcss@^4` installed, which is CSS-first config — there is no
`tailwind.config.js`; tokens are declared via `@theme` in
`app/globals.css` instead. Use the v4 pattern:

```css
/* app/globals.css */
@import "tailwindcss";

:root {
  --color-primary-500: #34A853;
  /* ...rest of design-guidelines.md's color table */
  --radius-md: 0.75rem;
}

@theme inline {
  --color-primary-500: var(--color-primary-500);
  /* ... */
  --radius-md: var(--radius-md);
}
```

Replace the current placeholder `--background`/`--foreground` Geist
setup with the actual palette from `design-guidelines.md`.

Load fonts via `next/font/google` (Plus Jakarta Sans for headings,
Inter for body — both self-hosted at build time, per
`design-guidelines.md` Typography section) in `app/layout.tsx`,
replacing the current Geist font setup. This is a straightforward
substitution, not a decision — flagging the v3→v4 syntax gap above is
the only real judgment call in this step, and it's a mechanical
translation, not a new design decision.

## Step 5 — PWA manifest + service worker

Per `kencleng-frontend-tech-stack.md` PWA Scope (resolved): app-shell
caching only, no offline write queue, browser-default install prompt.

- `public/manifest.json` — name, short_name, `theme_color` (use
  `--color-primary-500` = `#34A853`), `background_color`, `display:
  "standalone"`, `start_url: "/"`. Icon assets
  (`icon-192.png`/`icon-512.png` etc.) don't exist yet in `public/` —
  that's a design-asset task, not a scaffold task. Reference the paths
  in `manifest.json` and leave a visible placeholder/TODO if the actual
  files aren't ready, rather than blocking this playbook on icon
  generation.
- `public/sw.js` — hand-written, minimal: precache the static app-shell
  assets (JS/CSS bundle, fonts, `manifest.json`) on `install`, serve
  cache-first for those, network-only for everything else (no data
  caching — that's the explicitly rejected scope). Register it from a
  small client component (`components/providers/sw-register.tsx`),
  gated on `typeof window !== 'undefined' && 'serviceWorker' in
  navigator`, mounted once in the root layout.

## Step 6 — Test infrastructure + mock-first dev mode

**Correction to this playbook's earlier scope**: an earlier version of
this step said MSW would be node-mode only ("no runtime need to mock
in the browser outside of tests"). That was wrong — the project's
frontend development workflow is explicitly **mock-first**
(build/verify a page against MSW-mocked responses before a real
backend endpoint exists, per §14's "parallel-track" premise), which
means MSW needs to run in the *browser* during `npm run dev`, not
just inside Vitest. Both modes share the same `mocks/handlers.ts` —
this isn't two mocking systems, it's one handler set used by two MSW
entry points (`msw/node` for tests, `msw/browser` for the dev server).

**Test mode (`msw/node`)**:
- `vitest.config.ts` — `environment: 'jsdom'`, React plugin, path alias
  matching `tsconfig.json`'s `@/*`.
- `vitest.setup.ts` — imports `@testing-library/jest-dom` matchers.
- `mocks/server.ts` — `setupServer()` from `msw/node`, started/reset/
  closed in `vitest.setup.ts`'s `beforeAll`/`afterEach`/`afterAll`.

**Dev mode (`msw/browser`)**:
- `npx msw init public/ --save` — generates `public/mockServiceWorker.js`
  (the actual service worker MSW registers in the browser). This is a
  separate file from `public/sw.js` (the app-shell PWA worker from
  Step 5) — the two don't conflict, but don't confuse them when
  debugging.
- `mocks/browser.ts` — `setupWorker(...handlers)` from `msw/browser`.
- `components/providers/mocking-provider.tsx` — a small client
  component, mounted in the root layout **before** `QueryProvider`,
  that:
  1. Reads `process.env.NEXT_PUBLIC_API_MOCKING`.
  2. If not `'true'`, renders `children` immediately (mocking off —
     this is the default for anything going through Caddy/
     `docker-compose` against a real backend).
  3. If `'true'`, calls `worker.start()` (from `mocks/browser.ts`) and
     holds rendering `children` until the promise resolves — avoids a
     race where the first page's `fetch` calls escape before the
     worker is listening.
  4. **Hard safety check, not just the env var**: never call
     `worker.start()` if `process.env.NODE_ENV === 'production'`,
     regardless of what `NEXT_PUBLIC_API_MOCKING` is set to. A
     `NEXT_PUBLIC_*` var is bundled into client JS and readable by
     anyone — the production-guard is a second, non-optional gate so a
     stale `.env` value can't accidentally ship mock interception to
     real users.
- `.env.local` (git-ignored, per `.gitignore`'s `.env*` rule) —
  `NEXT_PUBLIC_API_MOCKING=true` for standalone `npm run dev`
  (`:3000`, no backend reachable). Leave unset (or `false`) when
  running via the root `docker-compose` — real backend through Caddy,
  same-origin, matches `frontend/README.md`'s existing note that
  auth-related work needs to go through `http://localhost` anyway.

**Shared, either mode**:
- `mocks/handlers.ts` — starts empty (`export const handlers = []`).
  Don't write speculative handlers for endpoints no page/component
  needs yet — add one per endpoint as the page that needs it gets
  built, per the mock-first workflow (see durable section below).
- Add to `package.json` scripts: `"test": "vitest run"`, `"verify":
  "npm run lint && npm run test"` — referenced by `frontend/README.md`
  already but don't exist yet.

**Explicitly not resolved by this playbook** (flag, don't invent):
`docs/kencleng-agentic-workflow.md` §14 step 3 calls for a "contract
check (generated types match the mock fixtures actually used in
tests)" gate, separate from the lint/test/a11y gates. There's no
tooling decision anywhere in the docs for *how* this check is actually
implemented (custom script diffing `lib/api/schema.d.ts` against
`mocks/handlers.ts`? A convention enforced by code review only?). Don't
guess at a mechanism here — this needs an actual decision, likely once
Task #1's frontend track has real fixtures to check against.

## Step 7 — API types generation

```bash
npx openapi-typescript ../api/openapi.yaml -o lib/api/schema.d.ts
```

Run once now (against whatever the spec currently contains) so
`lib/api/` isn't empty, and re-run whenever `api/openapi.yaml` changes
— per `frontend/README.md`. Never hand-edit the output.

## Step 8 — Verify

```bash
npm run dev            # should start cleanly on :3000
npm run lint
npm run test            # should run 0 tests cleanly, not error
npm run build            # catches any App Router / font / manifest wiring mistake
```

`npm run verify` at this stage passes trivially (no real tests yet) —
expected, not a gap. Don't write placeholder tests just to have
something for it to run, matching the same rule
`scaffold-backend.md` Step 4 states for the backend side.

## Step 9 — Human checkpoint

Per `docs/kencleng-agentic-workflow.md` §10/§14 step 6, this isn't
Tier 0 (no server-side correctness logic here), but still needs a look
before merge:

- [ ] No `components/features/<domain>/` or `lib/hooks/` content was
      accidentally started in this session — scope stayed at
      structure/providers/tooling only, no real page/component logic
      (that's Task #1, not this playbook)
- [ ] `design-guidelines.md`'s v3→v4 Tailwind config translation was
      done faithfully (spot-check a couple of color tokens actually
      render correctly), not just structurally present
- [ ] Service worker registration is genuinely gated to the browser
      environment and doesn't break SSR/build

## Open items surfaced by this playbook (not resolved here)

- **Contract-check gate mechanism** (Step 6) — needs a real decision,
  not before Task #1 has actual fixtures to check.
- **Frontend security-tooling equivalent** — backend has
  `gosec`/`gitleaks`/`govulncheck` (`backend/.agents/docs/
  security-tooling.md`); `docs/kencleng-agentic-workflow.md` §14 lists
  lint → test → contract check → a11y as the frontend gate order, with
  no explicit security-scanning tool named (`npm audit`? a
  `eslint-plugin-security`-style rule set?). Not decided anywhere —
  don't add one speculatively in this playbook.
- **PWA icon assets** (Step 5) — design-asset work, not scaffolding.

## Mock-First Development Workflow — durable, applies to every page/task

**This section outlives Steps 0-9 above** — it's the answer to "how
does a page actually get built," referenced by every domain's frontend
track from here on, not just this one-time scaffold.

Per page/component (regardless of domain), the cycle is:

1. **Mock** — build the page against `lib/api/<domain>.ts` (typed via
   `lib/api/schema.d.ts`) + a `mocks/handlers.ts` entry shaped from the
   same generated types. No live backend needed — this is what makes
   frontend and backend tracks parallelizable per §15.
2. **Verify against design + contract** — check the page against
   `docs/ui-ux/design-guidelines.md` (visual), `docs/ui-ux/patterns.md`
   (states: loading/empty/error/success per §B), and
   `docs/ui-ux/prototype-reference.md` (Tier 1 pages — known issues to
   not carry over). "Contract" here means the mock fixture itself
   can't drift from `schema.d.ts` — if it's shaped correctly, this
   step is a visual/behavioral check, not a data-shape check (that's
   already enforced by TypeScript).
3. **Integrate real** — once the real backend endpoint exists (in
   whichever order it lands — no fixed dependency on backend task
   order), point that one endpoint at the real API. The simplest way:
   MSW's `passthrough()` for that specific handler, or delete the
   handler from `mocks/handlers.ts` entirely once every consumer of it
   is confirmed working against the real endpoint — either way, no
   application code (`lib/api/`, components, hooks) changes, since
   they were never mock-aware to begin with.
4. **Test** — the existing `docs/kencleng-agentic-workflow.md` §14
   gate sequence (lint → test → contract check → a11y) runs against
   the **mocked** path (that's what `msw/node` in Vitest is for) and
   stays mocked even after step 3 — component tests shouldn't depend
   on a live backend being reachable. Step 3's real-backend check is a
   manual verification before merge, not a new automated gate; **flag,
   don't invent**: `kencleng-agentic-workflow.md` §14 doesn't currently
   name this step explicitly — worth raising as a doc gap at the
   workflow-doc level (root-scoped, out of this playbook's boundary)
   rather than silently assuming a mechanism here.

## Related docs

- `frontend/AGENTS.md` — project layout, style conventions
- `docs/project/kencleng-frontend-tech-stack.md` — full stack rationale,
  including the resolved PWA/testing/API-contract decisions this
  playbook implements
- `docs/ui-ux/design-guidelines.md` — visual tokens (Step 4)
- `docs/kencleng-agentic-workflow.md` §14 — the gate sequence this
  scaffold's test infra (Step 6) exists to support