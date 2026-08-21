# Patch — Frontend Tech Stack: UX content moved to docs/ui-ux/ (2026-08-20)

> Target file: `kencleng-frontend-tech-stack.md`
> How to apply: delete the two sections below in full, add the one
> replacement line noted at the end.

## 1. Delete "Shared Component Notes [NEW]" section in full

Everything from:
```
## Shared Component Notes **[NEW]**
```
through the end of the `CurationDecisionPanel` subsection, up to (not
including) `## Layout Patterns [NEW — RESOLVED 2026-07-20]`.

Reason: this is UX/behavior spec (what these components do), not a
code-architecture decision. Moved to `docs/ui-ux/patterns.md` §C.

## 2. Delete "Layout Patterns [NEW — RESOLVED 2026-07-20]" section in full

Everything from:
```
## Layout Patterns [NEW — RESOLVED 2026-07-20]
```
through the end of the "Public campaign pages" subsection, up to (not
including) `## Testing [RESOLVED — Step 6]`.

Reason: same as above — page-shape/layout decisions now live in
`docs/ui-ux/patterns.md` §A, cross-referenced from
`docs/ui-ux/page-map.md`.

## 3. Add one line under "Structural principles", after the
`components/shared/` bullet

```markdown
- **UX/layout behavior for these components (and page-level layout
  patterns generally) lives in `docs/ui-ux/patterns.md`**, not here —
  this doc stays scoped to code organization (where things live, how
  they're wired), not what they look like or how they behave on
  screen.
```

## 3b. Resolve 3 items in "Open Items — Needs Further Discussion"
(2026-08-20)

Replace:
```markdown
- Exact CORS / cookie configuration for the refresh-token flow (depends
  on final deployment topology — same-origin vs separate origins)
```
with:
```markdown
- ~~Exact CORS / cookie configuration for the refresh-token flow~~ →
  **resolved: no CORS config needed** — Caddy reverse-proxy makes
  FE+BE same-origin (`kencleng-repo-setup.md` §3.1), so
  `SameSite=Strict` works without cross-origin exceptions.
  **[RESOLVED — 2026-08-20]**
```

Replace:
```markdown
- Whether `middleware.ts`-based route protection is sufficient, or if
  dashboard guarding needs to happen at the layout/component level too
```
with:
```markdown
- ~~Whether `middleware.ts`-based route protection is sufficient~~ →
  **resolved: both, different jobs.** `middleware.ts` does the coarse
  check — redirect to `/login` if there's no session at all. Role-based
  gating (e.g. hiding legal-doc section from Staff, disabling
  Owner-only buttons) happens at the layout/component level, since it
  needs the actual role data fetched, which `middleware.ts` can't
  cheaply do on every request. **[RESOLVED — 2026-08-20]**
```

Add a new subsection after "Testing" (before "Open Items"):
```markdown
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
```

## 4. Update "Related Docs" (add if not already present)

```markdown
- Page patterns & shared component behavior: `docs/ui-ux/patterns.md`
- Visual tokens: `docs/ui-ux/design-guidelines.md`
- Page inventory: `docs/ui-ux/page-map.md`
```

---

Net effect: `kencleng-frontend-tech-stack.md` keeps only
code-architecture content — structural principles (`components/`
split, `lib/api/` vs `lib/hooks/`, Zustand-per-domain), Testing, and
Open Items. Everything UX-shaped moves to `docs/ui-ux/`.
