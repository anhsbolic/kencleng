# Kencleng — Historical Research & Pilot Evidence

> **Status:** Historical navigation only on research branch `pilot/station5-c1-space-only`. Contents below were removed from the active documentation tree during concern-based restructuring; **the originals are unchanged at the immutable source checkpoint** `875074914474ee07790fb2704a806a3d077c8ffb`. Neither historical status nor Git availability grants current Product, Engineering, API, implementation or runtime authority.

## Rancang ORG-01 — unaccepted Engineering solution bundle (25 October 2026)

- [Exact original Rancang ORG-01 directory snapshot](https://github.com/anhsbolic/kencleng/tree/875074914474ee07790fb2704a806a3d077c8ffb/docs/engineering/org-01/rancang-001) — seven original files, including state, handoff, producer obligations/technical/experience and independent security/solution reviews.
- Its bounded design-production recommendation was COMPLETE, **but final exact-bundle Human acceptance remained PENDING**; implementation, canonical API adoption and runtime verification were not performed.
- **Known later policy incompatibilities:** the proposed opaque server-side session differs from current T02 JWT access/rotating-refresh policy; proposed `actor_id` account schema differs from current T02 internal `account_id` model; proposed response conventions differ from T06 `{data, meta?}`, `camelCase` and RFC 9457. Do not silently rebase or use historical proposal as accepted input. New Engineering work must use current Product and [effective Engineering Standards](../engineering/standards/README.md), with a separately authorized scoped solution when needed.
- The original files retain their exact source paths/embedded hashes **at the source checkpoint**, rather than being rewritten as though reviewers had seen the new folder structure.

## Station 5 C1 pilot episode (October 2026)

- [Original Space-only pilot entry](https://github.com/anhsbolic/kencleng/blob/875074914474ee07790fb2704a806a3d077c8ffb/docs/project/station5-c1-space-only-pilot-entry.md)
- [Original interim closeout report](https://github.com/anhsbolic/kencleng/blob/875074914474ee07790fb2704a806a3d077c8ffb/docs/project/station5-c1-s1-interim-closeout.md)

These are historical experiment/decision evidence, **not active Product or Engineering instructions**. The Product-owned [history/crosswalk](../product/history.md) and [identifier migration](../product/identifier-migration-2026-10-23.md) remain in Product because their owning concern is Product provenance. Do not infer an approved Engineering handoff or current feature contract from old C1 records.


## Product / ORG-01 research progress reset — 16 November 2026

**Anhar authorized an intentional working-progress reset on `research/cross-stage-mapping`, not a Git reset, retroactive cancellation of historical Human approval, or a change to `main`.** The last complete pre-reset research snapshot is [`e43bc0276a9d027813e07d19099d86a5ee9e8c49`](https://github.com/anhsbolic/kencleng/tree/e43bc0276a9d027813e07d19099d86a5ee9e8c49).

Recovered there: approved Product Intent and terminology, interconnected Journey and selected ORG-01, ORG-01 Product Contract/approval/privacy decision/draft notice, feature-specific Design candidate/prototype, `rancang-002` and `nextjs-builder-001` evidence, deferred API adoption, and mock frontend implementation/tests/generated contract. **Those approvals retain their historical exact source/decision scope.** They are no longer selected as *active Product/ORG-01 progress* on the current research branch. Previous ORG-01 requirements, status, API, privacy notice and mocks must **not** be inherited by a new Product session without explicit new owner decision.

Project Standards/brand assets, development environment and neutral frontend/backend/API scaffolds were deliberately kept. Read the current [Product discovery router](../product/README.md); treat old feature files from the pinned commit as research/evidence, not live requirements or deliverables. The reset itself does not prove the neutral frontend or any browser/runtime behavior.
