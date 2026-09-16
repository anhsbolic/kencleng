# Kencleng — Frontend Reboot Deletion Manifest

> Status: Executed and reconciled
> Decision owner: Anhar Solehudin
> Prepared with: ChatGPT — GPT-5.6 Sol
> Initial reset: 2026-09-16
> Post-merge hygiene reconciliation: 2026-09-16
> Parent plan: `docs/project/frontend-reboot-plan.md`
> Purpose: record the actual clean-start reset boundary for the active frontend tree

## 1. Principle

Retention must be earned by current engineering value. Existing product/UI implementation and old harness guidance have no preservation privilege merely because they already exist.

The reset removes the previous frontend **product implementation and stale execution precedent**, not the repository's ability to build a frontend.

Git history remains the archive for everything removed.

## 2. Final top-level classification

| Path | Final classification | Reason |
|---|---|---|
| `frontend/.agents/` | DELETE | Post-merge audit found old operational playbooks that encoded the retired frontend generation, old green/type/prototype assumptions, and pre-seeded component/provider/API structure. They competed with current specs, architecture, and canonical Harscode prompts. |
| `frontend/.claude/` | DELETE | Post-merge audit found partial vendor-specific lifecycle wrappers and best-practice adapters. The command wrappers bypassed current canonical Harscode phase prompts; the skill adapter set was incomplete and could drift. `frontend/CLAUDE.md` remains the Claude routing stub to `AGENTS.md`. |
| `frontend/.gitignore` | RETAIN | Generic tooling hygiene. `.local-agents/` intentionally remains tracked for learning-by-doing. |
| `frontend/.local-agents/` | RESET / RETAIN POLICY | Keep the directory/policy; remove old `works/**` active history; Git history remains archive; new runs create fresh committed evidence. |
| `frontend/AGENTS.md` | RETAIN | Current frontend authority/router and clean-start rules. |
| `frontend/CLAUDE.md` | RETAIN | Minimal harness routing stub (`@AGENTS.md`), without a second project-local lifecycle. |
| `frontend/README.md` | RETAIN / RECONCILE | Active clean-start operating summary; post-reboot wording must reflect development readiness rather than preparation state. |
| `frontend/app/` | RESET | Old routes/layout/product UI removed; only minimal technical `layout.tsx`, `page.tsx`, `globals.css` remain. |
| `frontend/components/` | RESET | Only governance `README.md` remains; reusable registry intentionally starts empty. |
| `frontend/eslint.config.mjs` | RETAIN / CLEAN | Generic lint capability retained; remove old generated-path assumptions until real implementation establishes those paths. |
| `frontend/lib/` | DELETE | Retired product/API/auth/hooks/store implementation. Recreate only when new tasks require it. |
| `frontend/mocks/` | DELETE | Retired fixtures/handlers. MSW capability remains installed. |
| `frontend/next.config.ts` | RESET | Minimal typed config; old machine-specific configuration removed. |
| `frontend/opencode.jsonc` | RETAIN | Current harness references route to Harscode + Kencleng authorities and no longer point to removed prototype authority. |
| `frontend/package.json` | RESET / RETAIN CAPABILITIES | Keep selected engineering capabilities; remove retired feature/presentation dependencies and stale worker metadata. |
| `frontend/package-lock.json` | REGENERATED | Matches the reset package manifest. |
| `frontend/playwright.config.ts` | RETAIN | On-demand real-browser capability; empty `tests/browser` at baseline is valid. |
| `frontend/postcss.config.mjs` | RETAIN | Tailwind/PostCSS tooling. |
| `frontend/proxy.ts` | DELETE | Retired auth/session routing implementation. |
| `frontend/public/` | DELETE | Old PWA/MSW/default assets removed; recreate justified assets later. |
| `frontend/tests/` | DELETE | Old tests protected retired implementation. New tests grow with new production behavior. |
| `frontend/tsconfig.json` | RETAIN | Generic TypeScript/Next scaffold capability. |
| `frontend/vitest.config.ts` | RETAIN / CLEAN | Generic test capability and `passWithNoTests`; remove references to retired scaffold playbooks. |
| `frontend/vitest.setup.ts` | RESET | Generic Jest DOM setup only. |

## 3. Product/UI implementation reset

The active tree no longer carries the previous generation's:

- public/authenticated route UI and route-group compositions;
- product components and reusable UI/shared implementations;
- green/cool-neutral visual tokens, old font decisions, radius/elevation assumptions, and presentation CSS;
- Lucide-based retired presentation usage;
- feature hooks/stores/API wrappers tied to retired auth/product implementation;
- feature mocks/fixtures and feature-specific tests;
- old PWA/service-worker/static presentation assets;
- old `.local-agents/works/**` task history.

Git history remains available when historical understanding is materially useful, but history is not current implementation authority.

## 4. Minimal active application baseline

The reboot leaves only the technical source needed to prove the scaffold:

```text
frontend/app/
├── layout.tsx
├── page.tsx
└── globals.css

frontend/components/
└── README.md

frontend/.local-agents/
└── README.md
```

The bootstrap `/` surface is deliberately plain. It does not establish product layout, visual hierarchy, component contracts, or Sunlit Editorial implementation precedent.

## 5. Component and architecture reset

The following paths are architecture **destinations**, not mandatory scaffold folders:

```text
components/ui/
components/shared/
components/features/<domain>/
lib/api/
lib/hooks/
lib/stores/
mocks/
tests/browser/
```

Do not recreate them simply because the retired frontend had them.

The first new reusable `ui`/`shared` contracts must emerge from real usage and be registered in `frontend/components/README.md`.

## 6. Harness and guidance hygiene

### Removed `.agents/` playbooks

The post-merge audit found these files still active:

```text
frontend/.agents/docs/README.md
frontend/.agents/docs/scaffold-frontend.md
frontend/.agents/docs/phase0-shared-infra.md
frontend/.agents/docs/scaffold-public-shell.md
```

They were retired because they encoded the old generation directly, including pre-created folder/provider/API structures, old visual system decisions, removed prototype references, and historical route/shell/component assumptions.

Project setup/feature work must now derive from current project authorities and the canonical Harscode lifecycle rather than one-off legacy playbooks.

### Removed `.claude/` adapters

The post-merge audit also retired project-local `.claude/commands/**` and `.claude/skills/**`.

Reasons:

- command wrappers directly invoked old Harscode guidelines/checklists instead of current canonical `workflow/*-prompt.md` entrypoints;
- some wrappers encouraged loading examples by default, conflicting with current progressive-disclosure guidance;
- skill adapters represented only a partial snapshot of current Harscode React best practices and could drift independently;
- `frontend/CLAUDE.md -> @AGENTS.md` already provides a clean project routing entrypoint.

Harscode remains the portable workflow/best-practice authority. Kencleng should not maintain a second lifecycle copy under a client-specific directory.

## 7. Package and tooling result

Retained baseline capabilities include:

```text
Next.js / React / TypeScript
Tailwind CSS v4 / PostCSS
TanStack Query
React Hook Form + Zod
Zustand
Vitest + React Testing Library + jsdom
MSW
Playwright
openapi-typescript
ESLint
```

Keeping a dependency means the capability is available; it does not require the minimal scaffold or every feature to instantiate it.

Removed retired-only dependencies include:

```text
lucide-react
qrcode.react
```

Phosphor is the approved utility-icon direction, but an icon package should be added only when real implementation needs icons.

## 8. Config hygiene

- `next.config.ts` is minimal and contains no machine-specific legacy setting.
- `vitest.setup.ts` contains only generic Jest DOM integration.
- `vitest.config.ts` supports zero tests at the clean baseline and keeps browser tests separate without pointing to a retired playbook.
- `eslint.config.mjs` keeps generic Next ignores only; generated MSW/OpenAPI paths are not pre-seeded before those files actually exist.
- `playwright.config.ts` remains available as an on-demand browser runner; no product browser scenario is required at the reboot baseline.

## 9. `.local-agents/` policy

Kencleng intentionally commits local-agent workflow evidence for learning-by-doing.

Therefore **do not add `.local-agents/` to `.gitignore`**.

Rules:

```text
old task artifacts → Git history archive
new real task      → fresh .local-agents/works/<task>/ evidence
reusable truth     → promote into owning project/spec/design/architecture authority
```

Do not copy retired workflow artifacts forward to seed a new run.

## 10. Verification evidence

The primary reset scaffold was operator-verified by Anhar on 2026-09-16 with dependency synchronization, `npm run verify`, `npm run build`, and minimal boot/render inspection. That verification is recorded as operator-reported evidence in `frontend-reboot-plan.md`.

The post-merge hygiene cleanup in this manifest removes documentation/harness adapters and stale config assumptions; it does not add product behavior.

After the hygiene cleanup merges, its resulting `main` SHA becomes the final pre-CRTV project baseline and the validation branch must be re-pointed from that exact SHA before Exploration begins.

## 11. Explicitly not part of the reboot/hygiene work

- implementing Sunlit Editorial production UI;
- creating Button/Card/Input primitives;
- implementing auth/session;
- rebuilding mocks;
- implementing Campaign listing or final landing content;
- inventing product/domain semantics;
- starting Harscode Exploration/Techplan/Build before the final clean baseline is frozen.

Those begin only in the first real frontend development run after this manifest and active tree are reconciled.
