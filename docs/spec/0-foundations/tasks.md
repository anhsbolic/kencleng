# Task List — Foundations

> Status: agreed
> Last updated: 2026-09-17

## Scope

This spec area owns cross-domain project foundation work that establishes prerequisites for later feature delivery without creating a new business domain.

Foundation tasks may affect shared frontend architecture, visual/interaction foundations, tooling, or other project-wide delivery prerequisites. Business semantics remain owned by the applicable numbered business-domain specs.

## Tasks

Implementation/session boundaries follow the active Harscode workflow. Kencleng project preconditions remain defined in `docs/kencleng-agentic-workflow.md`.

| # | Task | Surface | Tier | Rationale | Parallel group |
|---|---|---|---|---|---|
| 1 | Frontend Experience Foundation | First representative public frontend surface beginning with `/`, from the clean frontend reboot baseline | 2 | Material cross-cutting frontend foundation work with no Tier-0/Tier-1 business responsibility; establishes the first production expression of the new canonical design system | Serial foundation |

## Parallel / serial grouping

Task 01 was the first real frontend development task after the completed frontend reboot.

The preparation/reboot work itself was project maintenance/orchestration, not Task 01 implementation and not a Harscode feature-development run.

Task 01 established enough frontend experience foundation before new domain feature surfaces proliferate. It did not require completing the landing page and did not block unrelated backend work.

## Status tracker

| # | Status | Notes |
|---|---|---|
| 1 | delivered | Implemented and verified in PR #24 (`71093b687cd7135495bc6ed62d520a621f96f586`) with human rendered acceptance PASS. Validation 02 closeout merged in PR #25 (`fca5a8f3178b53e5eec006064d8bcf2b078771b3`). See `features/01-frontend-experience-foundation.md`. |
