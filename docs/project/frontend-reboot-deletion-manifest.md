# Kencleng — Frontend Reboot Deletion Manifest

> Status: Approved preparation candidate; execute on `frontend-reboot-preparation`
> Decision owner: Anhar Solehudin
> Prepared with: ChatGPT — GPT-5.6 Sol
> Prepared: 2026-09-16
> Parent plan: `docs/project/frontend-reboot-plan.md`
> Purpose: classify the active frontend tree before the intentional clean-start reset

## 1. Principle

Retention must be earned by current engineering value. Existing product/UI implementation has no preservation privilege.

The reset removes the previous frontend **product implementation**, not the repository's ability to build a frontend.

Git history remains the archive for everything removed here.

## 2. Top-level classification

| Path | Classification | Reason |
|---|---|---|
| `frontend/.agents/` | RETAIN | Harness/project instruction capability, not product implementation. |
| `frontend/.claude/` | RETAIN | Harness/project instruction capability, not product implementation. |
| `frontend/.gitignore` | RETAIN | Tooling hygiene. `.local-agents/` intentionally remains tracked for learning-by-doing. |
| `frontend/.local-agents/` | RESET | Keep directory/policy, remove old `works/**` active history; Git history remains archive; create fresh README for new process evidence. |
| `frontend/AGENTS.md` | RETAIN / RECONCILE | Current authority routing is mostly aligned to Sunlit Editorial; add clean-start/reboot posture where needed. |
| `frontend/CLAUDE.md` | RETAIN | Harness routing stub. |
| `frontend/README.md` | RETAIN / RECONCILED | Already rewritten for clean-start generation. |
| `frontend/app/` | RESET | Remove all old routes/layout/product UI and replace with minimal technical `layout.tsx`, `page.tsx`, `globals.css`. |
| `frontend/components/` | RESET | Keep only governance `README.md`; delete all previous providers/features/shared/ui implementations. |
| `frontend/eslint.config.mjs` | RETAIN | Generic lint/tooling capability. |
| `frontend/lib/` | DELETE | Current contents belong to retired product/API/auth/hooks/store implementation. Recreate only when new tasks require them. |
| `frontend/mocks/` | DELETE | Current fixtures/handlers belong to retired features. MSW capability remains installed. |
| `frontend/next.config.ts` | RESET | Remove machine-specific `allowedDevOrigins`; keep minimal typed config. |
| `frontend/opencode.jsonc` | RETAIN / VERIFY | Harness config; keep unless a stale project-authority reference is found. |
| `frontend/package.json` | RESET | Keep selected engineering capabilities; remove retired feature/presentation dependencies and stale MSW worker-directory assumption. |
| `frontend/package-lock.json` | REGENERATE | Must match reset `package.json`; do not hand-edit dependency graph. |
| `frontend/playwright.config.ts` | RETAIN | On-demand browser capability; configuration is implementation-independent. |
| `frontend/postcss.config.mjs` | RETAIN | Tailwind/PostCSS tooling. |
| `frontend/proxy.ts` | DELETE | Current auth/session routing implementation belongs to retired frontend generation. Future route/session behavior is derived from current account/API requirements. |
| `frontend/public/` | DELETE | Current tree contains old PWA service worker/manifest, MSW worker, TODO icon note, and default starter assets. Recreate only justified assets/capabilities later. |
| `frontend/tests/` | DELETE | Existing browser scenarios protect retired implementation. Playwright capability remains. |
| `frontend/tsconfig.json` | RETAIN | Generic TypeScript/Next scaffold capability. |
| `frontend/vitest.config.ts` | RETAIN | Already supports zero tests and keeps Playwright separate. |
| `frontend/vitest.setup.ts` | RESET | Remove retired MSW server and old feature-specific `matchMedia` assumptions; keep only generic test setup required by the clean scaffold. |

## 3. `app/` reset

Delete from active tree:

```text
frontend/app/(auth)/
frontend/app/(dashboard)/
frontend/app/(public)/
frontend/app/verify-email/
frontend/app/favicon.ico
frontend/app/layout.tsx        # old provider/font/session composition
frontend/app/globals.css       # old green/cool-neutral visual system
```

Replace with:

```text
frontend/app/layout.tsx
frontend/app/page.tsx
frontend/app/globals.css
```

The replacements must be deliberately minimal:

- no auth/session providers;
- no old route shell;
- no inherited green brand tokens;
- no old fonts;
- no production component abstractions;
- no fake product semantics;
- no attempt to pre-implement Sunlit Editorial beyond the minimum neutral bootstrap needed to prove the scaffold.

`page.tsx` is a technical bootstrap surface only and must not be treated as the first product/design implementation.

## 4. `components/` reset

Retain:

```text
frontend/components/README.md
```

Delete:

```text
frontend/components/features/
frontend/components/providers/
frontend/components/shared/
frontend/components/ui/
```

The living registry intentionally starts empty. The next Harscode run creates new components only when real usage justifies them.

## 5. `lib/` reset

Delete the current `frontend/lib/` tree, including current:

```text
api/
hooks/
stores/
types/
cn.ts
otpauth.ts
otpauth.test.ts
```

Rationale:

- current implementations are coupled to retired frontend features/auth/session behavior;
- OpenAPI contracts remain canonical outside the frontend and generated types can be regenerated when needed;
- utility/helpers should reappear only when the new implementation has a real use;
- the project should not use old implementation shape as an implicit architecture template.

## 6. Mocks/tests reset

Delete current:

```text
frontend/mocks/
frontend/tests/
```

Retain the capabilities:

```text
Vitest
React Testing Library
MSW dependency
Playwright dependency/config
```

Do not keep a test simply to preserve historical coverage numbers when the tested subject is being retired.

New tests/mocks are created with the new production code.

## 7. `public/` reset

Delete current:

```text
TODO-icons.md
file.svg
globe.svg
manifest.json
mockServiceWorker.js
next.svg
sw.js
vercel.svg
window.svg
```

Rationale:

- default starter SVGs are not product assets;
- old manifest/service worker represent retired PWA behavior;
- MSW worker is regenerated when browser-side MSW is deliberately reintroduced;
- icon TODO belongs to the old visual generation.

The `public/` directory may disappear entirely until a new task needs a real static asset.

## 8. `.local-agents/` policy and reset

Kencleng intentionally commits local-agent workflow evidence for learning-by-doing.

Therefore **do not add `.local-agents/` to `.gitignore`**.

For this reboot:

```text
DELETE frontend/.local-agents/works/** from the active tree
KEEP Git history as the archive
CREATE frontend/.local-agents/README.md explaining the learning/process policy
```

The next real frontend Harscode run starts a fresh work tree. Historical artifacts must not be copied forward merely to seed context.

## 9. Package reset

Keep the selected baseline stack/capabilities needed for new development, but remove dependencies that exist only because of the retired frontend.

### Remove now

```text
lucide-react
qrcode.react
```

Reasons:

- Lucide is no longer the canonical utility-icon direction; Phosphor is the approved baseline, but no icon package needs to be installed until the new implementation actually uses icons.
- QR-code rendering is feature-specific and should return only when an active feature requires it.

### Keep as selected capabilities

```text
next
react
react-dom
@tanstack/react-query
react-hook-form
@hookform/resolvers
zod
zustand
Tailwind/PostCSS
Vitest/RTL/jsdom/MSW
Playwright
openapi-typescript
TypeScript/ESLint
```

Keeping a capability/dependency does not mean the minimal scaffold must instantiate it.

Remove the `msw.workerDirectory` package metadata during reboot; add/regenerate a browser worker later when a real mock-first task needs it.

Regenerate `package-lock.json` from the resulting `package.json`.

## 10. Config resets

### `next.config.ts`

Replace machine-specific configuration with minimal typed Next config. Do not preserve the current private-LAN `allowedDevOrigins` value as project architecture.

### `vitest.setup.ts`

Reset to generic setup only. At minimum, keep Jest DOM integration if still useful. Do not import a nonexistent `mocks/server` or preserve feature-specific `matchMedia` setup when no current production component requires it.

### `playwright.config.ts`

Retain current generic runner capability. It may point at an empty `tests/browser` location; absence of committed browser tests at reboot baseline is valid.

## 11. Minimal replacement source

The reset commit should leave a technical bootstrap that proves Next/Tailwind can compile without establishing product UI precedent.

### `app/layout.tsx`

Responsibilities only:

- metadata appropriate for a bootstrap Kencleng app;
- `lang="id"`;
- import minimal global CSS;
- render children.

No providers until real implementation requires them.

### `app/page.tsx`

A deliberately plain technical page indicating that the frontend reboot baseline is ready. No design-system claims, campaign content, CTA hierarchy, or feature behavior.

### `app/globals.css`

Only Tailwind import and minimal safe document/body baseline. Do not establish new production color/type tokens before the real Frontend Experience Foundation task.

## 12. Verification after reset

The reset is successful when the scaffold—not retired product behavior—can be verified.

Expected checks from a real checkout:

```text
npm install / lockfile consistency as needed
npm run lint
npm run test         # zero product tests is acceptable; Vitest uses passWithNoTests
npm run build
```

Also confirm:

- minimal `/` renders;
- no old product route/component import remains;
- no old green/prototype authority string remains in active frontend source/docs except clearly historical context;
- no retired `.local-agents/works/**` remains in active tree;
- `.local-agents/README.md` is tracked;
- package lock matches package manifest.

Playwright does not need a product scenario at this reset baseline; capability/configuration preservation is sufficient.

## 13. Explicitly not part of this reset

- implementing Sunlit Editorial production UI;
- creating the new Button/Card/Input primitives;
- implementing auth/session again;
- rebuilding mocks;
- implementing Campaign listing;
- creating browser regressions for not-yet-existing UI;
- running Harscode Exploration/Techplan/Build for Frontend Experience Foundation.

Those begin only after the reboot baseline is frozen.