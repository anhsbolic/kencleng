# Kencleng — UX Pattern System

> Status: Canonical
> Purpose: Reusable interaction behavior and state semantics.
> Boundary: This document does not define visual tokens, component APIs, frontend architecture, or business rules.

## 1. Pattern Rule

Reuse a pattern when the user interaction contract is materially the same.

```text
existing pattern fits
→ reuse

existing pattern mostly fits
→ extend intentionally

meaningfully different recurring interaction
→ propose a new pattern

one-off composition detail
→ keep local
```

A useful test is:

> Would we want the next ten similar surfaces to inherit this behavior?

Pattern reuse must never invent product capabilities unsupported by domain/API truth.

## 2. List / Browse

Purpose: help users scan, compare, search, filter, or select from a collection.

Typical structure:

```text
context / heading
→ optional search or filters
→ collection
→ navigation/pagination where required
```

The collection may be editorial cards, rows, or another appropriate representation. Visual form is not owned by this pattern.

### Loading
Prefer structure-preserving skeletons when useful rather than blocking the whole page with a spinner.

### Empty
Explain what is absent. Show a CTA only when an allowed action meaningfully resolves the empty state.

Distinguish first-use empty, query-empty, permission-limited, and true no-content states when the distinction matters.

### Error
Provide safe explanation and recovery when recovery is meaningful. Never expose raw backend errors.

### Search / filter / sort
Only expose semantics supported by product/API truth. Do not invent “featured”, “trending”, popularity, ranking, or sort meaning from UI convention alone.

## 3. Detail

Purpose: help users understand one entity, its current state, relevant trust/context, and available actions.

Typical structure:

```text
identity / title / state
→ primary facts
→ trust/context
→ supporting detail
→ actions appropriate to viewer + state
```

Public and authenticated detail surfaces may use different composition while preserving the same truth.

Long content may use progressive disclosure when it improves scanability, but disclosure must not hide information needed for the current decision.

## 4. Form

Purpose: help users provide or modify structured information with clear validation and consequences.

Typical structure:

```text
context
→ grouped fields
→ guidance
→ validation
→ primary submit action
→ secondary/cancel action where meaningful
```

### Validation
Client validation improves UX; server validation remains authoritative.

Keep field-specific validation near the field. Keep request/business failures at the form or section level.

Do not collapse:

```text
invalid field
```

and:

```text
request could not be completed
```

into one error treatment.

### Submission
Prevent accidental duplicate non-idempotent submission, communicate progress, and preserve user input unless successful flow intentionally transitions away.

### Revisable submission
When domain truth defines a lifecycle such as draft → review → rejected → revise → resubmit, editing availability and state presentation must follow that lifecycle. A reviewing state should not look like an editable form merely with disabled controls.

## 5. Dashboard / Summary

Purpose: give an authenticated user a prioritized overview of information requiring awareness or action.

A dashboard is not a collection of every available metric.

Each summary should answer at least one of:

- What is happening?
- What changed?
- What needs attention?
- Where should I go next?

Independent sections may load/fail independently when their data is independent.

Do not invent analytics, ranking, percentages, or KPIs merely because dashboards conventionally contain stat cards.

## 6. Curation / Review

Purpose: help an authorized reviewer understand submitted material and make an accountable decision.

Typical structure:

```text
submission context
→ material under review
→ relevant evidence/history
→ decision action
→ decision reasoning when required
```

The reviewed material should remain visually distinct from reviewer controls.

If a negative decision requires a reason, gather and explain the reason as part of the decision flow rather than after the decision is already committed.

## 7. Status / Tracking

Purpose: answer:

> What happened to the thing I submitted or initiated?

Typical structure:

```text
minimal context
→ current state
→ meaning / consequence
→ relevant next action
```

A status badge alone is insufficient when the state has meaningful consequence.

For sensitive unauthenticated lookup flows, visible failure distinctions must respect anti-enumeration/security requirements.

## 8. Evidence Journal

Purpose: help a donor follow what happened after donation without confusing funding, execution, and outcome.

Typical structure:

```text
donation fact
→ chronological campaign/program milestones
→ source/provenance
→ evidence or report context
→ pending next update where known
```

The Evidence Journal is factual first and human second.

It must distinguish:

### Funding Progress
Money collected relative to the campaign's funding model.

### Operational Progress
Execution/distribution/activity milestones supported by product data or reports.

### Reported Outcome
Outcome information reported by the appropriate source.

These categories must not collapse into one universal “impact progress” scale.

## 9. Primary Action Hierarchy

A surface should communicate one confident next action whenever one exists.

Actions may be primary, secondary, tertiary, or destructive.

Hierarchy follows user goal and consequence, not whichever button variant is visually strongest.

## 10. Loading

Loading treatment should be proportional to scope.

- page/section content: preserve expected structure where helpful;
- inline action: localized progress;
- background refresh: keep already-useful content visible when safe.

Avoid unnecessary layout shifts.

## 11. Empty States

First classify the state.

### Query/search empty
Nothing matches current criteria. Usually provide concise explanation and reset/change-filter action when useful.

### First-use empty
The user has not created/received anything yet. May justify stronger guidance or an expressive approved asset.

### Permission-limited
Do not disguise lack of permission as ordinary emptiness when the user should understand the distinction.

### True no-content
Do not invent an action just because empty-state templates usually have buttons.

## 12. Error and Recovery

Distinguish where useful and safe:

- field validation;
- request failure;
- unavailable/not found;
- permission failure;
- stale/freshness uncertainty;
- terminal business state.

A retry action is appropriate only when retry could reasonably succeed.

## 13. Success and Completion

Treatment should match significance.

A lightweight action may need only localized confirmation. A terminal or consequential flow should explicitly explain what completed, resulting state, what happens next, and relevant next destination/action.

Avoid excessive celebration around money, security, or other sensitive operations.

## 14. Status Communication

Status is a semantic system, not a badge style.

Use:

```text
label
+
visual semantic
+
context/consequence where necessary
```

Do not rely on color alone.

Friendly copy may explain a domain state but must not change its meaning.

## 15. Money Presentation

Every important amount requires an explicit semantic label.

When multiple currency values appear, users must be able to distinguish them without relying on color or position alone.

Funding progress should make the relationship between amount collected and applicable target/limit understandable when those concepts exist in product truth.

## 16. Confirmation and Consequential Actions

Use confirmation when an action is consequential, difficult to reverse, easy to trigger accidentally, or needs consequence explanation.

A good confirmation states the action and meaningful consequence using domain truth.

Avoid generic confirmation ceremony for low-risk reversible actions.

## 17. Destructive Actions

Use appropriate friction, not maximum friction.

For high-impact deletion/removal/revocation, identify the affected object and meaningful consequence. Keep destructive actions visually distinguishable from ordinary safe actions.

## 18. Progressive Disclosure

Use progressive disclosure for useful but non-essential current-decision content.

Do not bury:

- trust-critical information;
- financial consequences;
- validation errors;
- required next actions;
- material uncertainty.

## 19. Responsive Transformation

Responsive design preserves task and information priority; it does not merely stack desktop boxes.

When space becomes constrained:

1. preserve primary task;
2. preserve trust/consequence information;
3. preserve essential content;
4. keep actions reachable;
5. reflow/disclose secondary content intentionally.

Exact responsive mechanics belong to engineering.

## 20. Pattern vs Component

A UX pattern and a production component are not the same thing.

This document owns interaction semantics, state behavior, information relationships, and reusable experience rules.

Production component boundaries are an engineering concern.

## 21. Relationship to Other Authority

- `product-design-principles.md` — stable design judgment
- `brand-product-ui-brief.md` — selected brand/Product UI direction
- `asset-governance.md` — visual asset governance
- `page-map.md` — surface/persona inventory
- domain specs/OpenAPI — product truth
