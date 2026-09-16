# Kencleng Frontend

> Status: reboot preparation — old product implementation is being retired
> Last updated: 2026-09-16
> Reboot plan: `../docs/project/frontend-reboot-plan.md`

Kencleng's frontend is a Next.js App Router application. The project is intentionally resetting the active product/UI implementation so the next frontend generation starts from current product/API/design authorities rather than inheriting the superseded UI.

This README describes how to work inside `frontend/`. Generic lifecycle instructions belong to Harscode; project-specific rules live in `AGENTS.md` and the architecture/design authorities it routes to.

## Current architecture authority

Read:

```text
../docs/project/kencleng-frontend-tech-stack.md
```

for Kencleng-specific frontend architecture.

Read:

```text
../docs/ui-ux/README.md
```

for the current UI/UX authority map. The approved design direction is **Sunlit Editorial / Evidence-Led Optimism**.

Do not use removed legacy prototype/design generations from Git history as current implementation precedent.

## Selected engineering capabilities

The clean-start frontend baseline retains these project capabilities:

- Next.js App Router + React + TypeScript;
- Tailwind CSS v4, CSS-first;
- OpenAPI-generated TypeScript types;
- TanStack Query when client-owned server-state consumption is needed;
- React Hook Form + Zod for form lifecycle/UX validation;
- Zustand only for genuinely shared client-owned state;
- Vitest + React Testing Library;
- MSW capability;
- Playwright as on-demand browser automation capability.

A dependency being available does not require every feature to use it.

## Clean-start directory posture

The architecture describes ownership destinations, not mandatory folders.

As new implementation emerges, code may live in:

```text
app/<route>/
components/features/<domain>/
components/shared/
components/ui/
lib/api/
lib/hooks/
lib/stores/
mocks/
tests/browser/
```

Do not create or preserve one of these layers merely to make the repository look complete.

The narrowest truthful owner wins.

## Running the scaffold

After reboot preparation completes, ordinary local commands remain defined by `package.json`, typically:

```bash
npm install
npm run dev
npm run build
npm run lint
npm run test
npm run verify
```

Do not assume every command must run in every Harscode phase; current workflow/project verification guidance owns phase responsibility.

## API types

When the bundled OpenAPI contract changes and generated frontend types are needed, use the project-selected `openapi-typescript` generation path rather than hand-editing generated schema types.

The exact generated-file location should follow the active implementation. Do not preserve an old path solely for compatibility with retired frontend code.

## Browser automation

Playwright remains available as a separate real-browser capability.

It is not a default ritual for every Build/Review/Testing phase. Use it when the behavior/risk justifies repeatable browser automation and record why it is worth running.

Material UI still requires proportional human rendered acceptance.

## Learning/process evidence

`frontend/.local-agents/` is intentionally committed in Kencleng.

Kencleng is a learning-by-doing project, so representative workflow evidence should remain readable by other people. Old work directories may be retired from the active tree during major resets, while Git history remains the archive. New Harscode runs should create fresh task evidence rather than treating historical artifacts as current authority.

## Before starting new frontend development

Do not begin a new Harscode frontend run until the ready-for-development gate in:

```text
../docs/project/frontend-reboot-plan.md
```

has passed and the reboot baseline is frozen.