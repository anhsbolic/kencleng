# Task List — Foundations

> Status: agreed
> Last updated: 2026-09-15

## Scope

This spec area owns cross-domain project foundation work that establishes prerequisites for later feature delivery without creating a new business domain.

Foundation tasks may affect shared frontend architecture, visual/interaction foundations, tooling, or other project-wide delivery prerequisites. Business semantics remain owned by the applicable numbered business-domain specs.

## Tasks

Implementation/session boundaries follow the active Harscode workflow. Kencleng project preconditions remain defined in `docs/kencleng-agentic-workflow.md`.

| # | Task | Surface | Tier | Rationale | Parallel group |
|---|---|---|---|---|---|
| 1 | Frontend Experience Foundation | Representative public frontend surface, beginning with `/` | 2 | Material cross-cutting frontend work with no security/money-critical responsibility; shared-component blast radius is evaluated separately from risk tier | Serial foundation |

## Parallel / serial grouping

Task 01 establishes enough frontend experience foundation before domain feature surfaces proliferate. It does not require completing the landing page, and it does not block unrelated backend work.

## Status tracker

| # | Status | Notes |
|---|---|---|
| 1 | not started | Selected as the first frontend foundation work item. See `features/01-frontend-experience-foundation.md`. |
