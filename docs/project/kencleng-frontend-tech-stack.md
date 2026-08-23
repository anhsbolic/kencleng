### Structural principles

- **Feature components co-located in `components/features/`** first (not
  nested inside `app/`), so they're reusable across pages without
  duplication. Move to route-nested components only if something
  genuinely isn't reusable.
- **`lib/api/` is separate from `lib/hooks/`** — `api/` holds pure fetch
  functions (usable from both Server Components and inside hooks),
  `hooks/` wraps them with TanStack Query for use in Client Components.
  **[UPDATED — Step 2]** Request/response types in `lib/api/` are now
  generated from `api/openapi.yaml` via `openapi-typescript` (types
  only, not a full client SDK) — the fetch functions themselves stay
  hand-written, just typed against the generated interfaces instead of
  hand-written ones. See `kencleng-backend-tech-stack.md` "API Contract
  & Codegen" for the full decision.
- **Zustand stores are split per domain**, not one giant store — easier
  to read and maintain.
- **`components/shared/` is for cross-domain non-primitive components
  [NEW]** — different from `components/ui/` (which is pure design-system
  primitives with no business awareness) and different from
  `components/features/` (which is domain-specific). `shared/` sits in
  between: components like `MaskedField` know about the *concept* of
  PII across multiple domains (donation, organization, user) but aren't
  domain-specific themselves.
- **UX/layout behavior for these components (and page-level layout
  patterns generally) lives in `docs/ui-ux/patterns.md`**, not here —
  this doc stays scoped to code organization (where things live, how
  they're wired), not what they look like or how they behave on
  screen.

## Testing [RESOLVED — Step 6]

- **Test runner: Vitest** — native ESM, faster than Jest, minimal
  config for a TypeScript + Next.js App Router project. Consistent
  with the rest of the modern tooling already in use (Vite-based
  ecosystem, Tailwind, TanStack Query). Jest was considered but
  rejected: more established, but ESM + App Router config friction
  outweighs any familiarity benefit for a fresh project.
- **Component testing library: React Testing Library** — de-facto
  standard for React, philosophy of testing behavior (query by
  role/text) rather than implementation detail. Effectively the only
  reasonable choice, not treated as a real trade-off decision.
- **Scope for v1: unit + component tests only, no E2E yet.** Covers
  hook/util logic and individual component behavior (form validation,
  conditional rendering, etc.). Playwright (or similar) for E2E is
  explicitly deferred — consistent with "lowest complexity, add only
  when there's a demonstrated need": E2E setup (browser automation,
  test data seeding) isn't justified before there's an actual page/flow
  to test end-to-end. Revisit once the first vertical slice
  (Registrasi & Login) exists.
- **API mocking: MSW (Mock Service Worker)** — intercepts requests at
  the network layer, so components under test exercise the same code
  path as runtime (real `fetch` calls, real TanStack Query hooks)
  rather than mocking `fetch`/the query client directly. Chosen over
  manual mocking (`vi.mock` on fetch functions) because the API
  contract is already type-safe via the OpenAPI-generated types (Step
  2) — MSW keeps tests aligned with that same contract shape instead
  of hand-rolled mock responses that can silently drift.

## PWA Scope [RESOLVED — 2026-08-20]

**App-shell caching only for v1**: static assets (JS/CSS/fonts,
`manifest.json`) are cacheable via service worker, so the shell loads
offline — but no data caching/offline write queue. Pages relying on
live data (donation status, curation queues, etc.) show cached/stale
data with a freshness indicator when offline (see
`docs/ui-ux/patterns.md` §B, "Stale/offline data") rather than
attempting real offline writes.

**Install prompt: browser-default only** — no custom "Add to Home
Screen" UI/banner. Consistent with lowest-complexity: a custom install
prompt is worth building once there's a demonstrated need (e.g. low
install conversion with the default browser prompt), not before.

Rejected for v1: background sync / offline donation queueing (donation
is a financial transaction — queueing it for later submission without
the user seeing real-time confirmation is a correctness/trust risk,
not just a complexity one).

## Open Items — Needs Further Discussion

- ~~Exact CORS / cookie configuration for the refresh-token flow~~ →
  **resolved: no CORS config needed** — Caddy reverse-proxy makes
  FE+BE same-origin (`kencleng-repo-setup.md` §3.1), so
  `SameSite=Strict` works without cross-origin exceptions.
  **[RESOLVED — 2026-08-20]**
- ~~API contract format (REST plain JSON vs OpenAPI spec-first)~~ →
  **resolved: OpenAPI 3.x spec-first**, types generated via
  `openapi-typescript` (types only, fetch functions stay hand-written)
  **[RESOLVED — Step 2]**
- ~~Whether `middleware.ts`-based route protection is sufficient~~ →
  **resolved: both, different jobs.** `middleware.ts` does the coarse
  check — redirect to `/login` if there's no session at all. Role-based
  gating (e.g. hiding legal-doc section from Staff, disabling
  Owner-only buttons) happens at the layout/component level, since it
  needs the actual role data fetched, which `middleware.ts` can't
  cheaply do on every request. **[RESOLVED — 2026-08-20]**
- `MaskedField` reveal persistence behavior — **resolved: stays
  revealed until manual re-toggle or page refresh/navigation, plain
  local component state** — see `docs/ui-ux/patterns.md` §C.
  **[RESOLVED — 2026-08-21]**
- ~~Guest donation status page: server-side state (token-based lookup) vs
  client-side state~~ → **resolved: token-in-URL, server-side lookup by
  token** — see `kencleng-phase2-detail.md` Fitur 1 and
  `docs/ui-ux/page-map.md` **[RESOLVED — NEW]**

## Not Yet Discussed

- Deployment/hosting target for the frontend (same host as backend via
  Docker Compose vs separate hosting like Vercel)
- ~~Testing approach for the frontend~~ → see **Testing** section
  above **[RESOLVED — Step 6]**

## Related Docs

- Page patterns & shared component behavior: `docs/ui-ux/patterns.md`
- Visual tokens: `docs/ui-ux/design-guidelines.md`
- Page inventory: `docs/ui-ux/page-map.md`
- Which routes have a Claude Design prototype vs derive from patterns
  alone: `docs/ui-ux/prototype-reference.md`
- How to extract/use `design-reference/` exports:
  `docs/ui-ux/design-reference-usage.md`