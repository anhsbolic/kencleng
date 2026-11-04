# Product Design Standards

> **Standing:** Project-owned Design concern map for Kencleng **pilot/station5-c1-space-only**. This index is a **router**, not a second source of visual rules, Product truth, feature UX/UI acceptance, Engineering design or release scope. The operational `main` branch has separate authority; do not infer cross-branch adoption.
> **Scope:** Approved reusable project Design direction and patterns, current working supplements, supporting visual evidence and optional authoring assistance. Read the **smallest actual owning source** for the decision.

## Design authority map

| Concern / responsibility | Owning source | Standing in this branch | Authority boundary |
| --- | --- | --- | --- |
| Whole-product purpose, actor/domain meaning and business commitments | [Product Intent](../product/product-intent.md) plus applicable approved scoped Product sources | **Product authority — outside Design** | Product determines substantive truth, permissions, financial/legal/privacy consequences; Design cannot infer or amend these |
| Reusable principles for design judgment and human decision quality | [Product Design Principles](product-design-principles.md) | **Canonical Design** | Owns evaluative experience principles; not a second Product Contract, exact interaction flow, palette or brand selection |
| Selected brand character and Product UI direction | [Brand/Product UI Brief](brand-product-ui-brief.md) | **Canonical Design** | Owns Sunlit Editorial / Evidence-Led Optimism thesis and directional expression; concrete tokens, reusable UX patterns and assets have separate owners |
| Reusable concrete visual-system rules | [Visual Design Guidelines](design-guidelines.md) | **Canonical Design** | Owns color roles/values, typography roles, visual hierarchy and visual expression; not backend/domain status truth, production CSS/component APIs or feature UX decisions |
| Reusable interaction and user-facing state behavior | [UX Pattern System](patterns.md) | **Canonical Design** | Owns recurring interaction patterns; not per-feature IA, webpage routes/flows, business rules or frontend component implementations |
| Asset provenance, approval, reuse and change impact | [Asset Governance](asset-governance.md) | **Canonical Design** | Owns asset selection/governance and asset decision standing; not selected brand direction or core visual tokens |
| Decision-stage IA and disclosure heuristics | [Information Architecture](information-architecture.md) | **Working supplement** | Informs specific UX/Design reviews; not a canonical replacement for Product Principles/UX Patterns or accepted feature IA |
| Indonesian UX writing and terminology heuristics | [Copywriting Bahasa Indonesia](copywriting-id.md) | **Working supplement** | Supports UI copy; not binding Product/legal/privacy wording or automatic feature-copy approval |
| Selected-direction visual studies | [Visual Reference Index](visual-references/selected-direction/README.md) | **Approved direction evidence** | Supports visual character and hierarchy; not production assets, pixel geometry, Product facts or accepted feature screens |

**One decision, one owning source.** The same principle may be referenced by Brand, Patterns and Visual Guidelines for its consequences, but the reference does not create an independently changeable competing rule. In a conflict, inspect scope/revision/standing of the **actual owning project source**; raise the specific inconsistency with the relevant Product/UX/Design/asset owner instead of resolving it by choosing the most convenient document.

**Feature-specific work is separate:** Accepted Product outcome/contracts establish feature meaning; scoped UX flows, IA, page maps and Lo-Fi establish actual experience structure; feature-specific UI candidates/Hi-Fi evidence are accepted **only for their reviewed scope and exact revision**. A candidate or rendered screenshot is not reusable project-wide Design authority. Engineering owns frameworks, DOM/CSS, component implementation, APIs/security and actual runtime verification; [Frontend component governance](../../frontend/components/README.md) owns production reusable component-contract maturity, not this Design index.

**Evolution:** A new visual/interaction precedent is **not canonical merely because a feature used it**. Reuse existing rules when sufficient. If a genuinely cross-feature new rule or change is needed, identify the owning Design source, affected actual consumers, sufficient evidence and legitimate Human decision; update only that owner and reconcile stale cross-references. Do not invent mandatory standard categories or a centralized second master spec.

## Feature-specific evidence and optional assistance

- [ORG-01 Hi-Fi candidate and handoff](outcomes/org-01/handoff.md) is a **bounded Design Candidate / NOT HUMAN APPROVED**, not canonical standards or proof of rendered/browser/actor outcomes.
- [Static HTML Hi-Fi kickoff template](templates/hi-fi-static-html-kickoff.md) is an **optional session prompt** for appropriately bounded project UI design when existing accepted evidence is insufficient; not a formed Space, required phase, implementation permission, design approval or new project standard.

Use actual current project Product/Design original sources directly. This index does **not** choose features, webpage routes, release scope, delivery order, or implementation rights.
