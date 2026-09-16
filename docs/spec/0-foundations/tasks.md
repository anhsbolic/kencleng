# Task List — Foundations

> Status: agreed
> Last updated: 2026-09-16

## Scope

This spec area owns cross-domain project foundation work that establishes prerequisites for later feature delivery without creating a new business domain.

Foundation tasks may affect shared frontend architecture, visual/interaction foundations, tooling, or other project-wide delivery prerequisites. Business semantics remain owned by the applicable numbered business-domain specs.

## Tasks

Implementation/session boundaries follow the active Harscode workflow. Kencleng project preconditions remain defined in `docs/kencleng-agentic-workflow.md`.

| # | Task | Surface | Tier | Rationale | Parallel group |
|---|---|---|---|---|---|
| 1 | Frontend Experience Foundation | First representative public frontend surface beginning with `/`, after the clean frontend reboot baseline is frozen | 2 | Material cross-cutting frontend foundation work with no Tier-0/Tier-1 business responsibility; establishes the first production expression of the new canonical design system | Serial foundation |

## Parallel / serial grouping

Task 01 is the first real frontend development task **after frontend reboot preparation completes**.

Preparation/reboot work itself is project maintenance/orchestration, not Task 01 implementation and not a Harscode feature-development run.

Task 01 establishes enough frontend experience foundation before new domain feature surfaces proliferate. It does not require completing the landing page, and it does not block unrelated backend work.

## Status tracker

| # | Status | Notes |
|---|---|---|
| 1 | blocked by reboot preparation | Do not start Harscode development until `docs/project/frontend-reboot-plan.md` reaches its ready-for-development gate and the clean frontend baseline is frozen. See `features/01-frontend-experience-foundation.md`. |
