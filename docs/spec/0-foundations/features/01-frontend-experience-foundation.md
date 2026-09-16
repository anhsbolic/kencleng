# Feature Spec — Frontend Experience Foundation

> File: `docs/spec/0-foundations/features/01-frontend-experience-foundation.md`
> Status: delivered
> Risk tier: 2
> Scope: cross-domain frontend foundation; not a business domain
> Last updated: 2026-09-17
> Delivery evidence: PR #24 / `71093b687cd7135495bc6ed62d520a621f96f586`; Validation 02 closeout PR #25 / `fca5a8f3178b53e5eec006064d8bcf2b078771b3`

## Feature surface

The first representative production surface begins at `/`.

The route is a calibration surface, not a requirement to complete the full landing page. Harscode Exploration determines the minimum coherent rendered slice required to establish and judge the new frontend foundation.

## Starting condition

This task begins **after the frontend reboot baseline is frozen**.

The old frontend product/UI implementation is not an implementation authority for this task. Git history may be inspected only when historical understanding is materially useful; it must not silently restore superseded visual/component decisions.

The active engineering baseline is expected to be a clean frontend scaffold containing current tooling/configuration but no established legacy product UI, shared visual primitives, feature UI, or route composition that must be preserved.

Canonical inputs are:

- product/domain specs and API contracts for business truth;
- `docs/ui-ux/README.md` and the authorities it routes to for current product/design truth;
- `docs/project/kencleng-frontend-tech-stack.md` for frontend architecture;
- `frontend/AGENTS.md` and current component governance for implementation boundaries;
- active Harscode workflow for the development lifecycle.

## Goal

Establish enough of Kencleng's **new** public-facing frontend experience foundation that subsequent frontend surfaces can inherit a coherent implementation of Sunlit Editorial / Evidence-Led Optimism, visual hierarchy, and interaction language.

The task should create the first trustworthy production precedent for the new frontend generation without prematurely building a complete design system or landing page.

## Requirements

- The implemented representative slice must be sufficient for a human to judge whether the rendered UI coherently expresses the current canonical design authorities.
- The slice should establish only the foundation actually needed by the representative experience, such as public shell/navigation, typography, spacing, color/surface hierarchy, CTA hierarchy, representative content treatment, responsive behavior, and the smallest justified reusable primitives.
- New components and abstractions must follow semantic-owner-first placement. Do not recreate the retired component tree merely because it existed previously.
- Global tokens or reusable primitives may be introduced when they are genuinely needed by the new foundation and are supported by the canonical visual system; avoid speculative variant systems or compatibility layers for retired frontend code.
- Business/domain semantics remain owned by canonical domain specs and API contracts. This work must not invent unresolved Campaign behavior, ranking/curation meaning, trust claims, evidence claims, or other missing product rules merely to make `/` look complete.
- Approved visual references are direction evidence, not pixel-perfect screenshots to reproduce.
- Material rendered UI must be evaluated at representative desktop and mobile scope.
- Any Level 3/Level 4 asset need follows `docs/ui-ux/asset-governance.md`; brand-defining assets still require the applicable human approval.

## Explicit non-goals

- Completing the entire landing page.
- Completing every public route.
- Rebuilding all retired frontend features during this foundation task.
- Preserving old frontend component APIs, tokens, layouts, route composition, tests, mocks, or visual behavior for compatibility.
- Resolving final Campaign-domain semantics, including unsupported highlighted/featured/trending/curation meaning.
- Implementing the Campaign backend or requiring live Campaign integration.
- Creating a complete design system before real usage justifies it.
- Reconstructing removed prototype/reference authority from Git history.

## Acceptance criteria

- Given the clean reboot frontend baseline and current product/design authorities, when the representative `/` slice is rendered at representative desktop and mobile sizes, then the result provides enough evidence for informed human product judgement on brand/experience coherence.
- Given the accepted representative slice, when subsequent frontend surfaces are implemented, then they have a clearer production visual/interaction foundation to inherit without requiring the landing page to be complete first.
- Given unresolved domain/product semantics outside this foundation task, when the representative surface would need those semantics, then the implementation keeps the gap explicit/provisional instead of fabricating canonical behavior.
- Given any newly established broad `ui/` or `shared/` component contract, when it becomes reusable production authority, then its contract/registry entry and representative impact expectations are documented according to current component governance.
- Human rendered acceptance is required before this task is considered accepted as the first production brand/experience calibration of the new frontend generation.

## Verification expectations

Verification should be proportional and should explain why each non-trivial verification class is worth running.

Minimum expected evidence:

- normal static/build verification appropriate to the implemented scope;
- focused unit/component behavior verification where meaningful behavior exists;
- representative real-browser inspection at desktop and mobile sizes;
- human rendered acceptance.

Playwright/browser automation is **not automatically required by phase**. Use or add it when the task contains browser behavior or a repeatable regression whose risk justifies automation. When proposed, the plan should state:

```text
why this browser automation is useful
risk if omitted
which phase owns the authoritative run
```

Do not claim a verification step that was not actually run.

## Applicable invariants / threat model

This is a cross-domain frontend foundation task, not a business-domain feature. No new business invariant or domain threat model is introduced by this spec. Existing root/frontend security, truthfulness, accessibility, and rendering rules continue to apply where triggered by the implementation.

## Risk tier & rationale

Tier 2 — standard verified frontend work. The task is material and foundational but does not itself own Tier-0/Tier-1 money, security, authorization, or PII-critical behavior. A later feature may have a different risk tier even when it inherits this visual foundation.

## Assumptions / open questions

- The minimum representative slice is intentionally not prescribed here; Exploration determines it from current authorities and the clean live scaffold.
- The first implementation may establish only a small set of reusable tokens/primitives. Absence of a complete component library is acceptable.
- The task establishes enough foundation for the current project stage, not a permanently finished design system.

## Delivery record

Task 01 was implemented through PR #24 and accepted with human rendered acceptance PASS. Validation 02 was then closed through PR #25. This record changes project status only; the feature requirements above remain historical evidence of what the delivered foundation was required to establish.
