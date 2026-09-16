# Kencleng Frontend Component System

> Status: Active — clean-start generation
> Last updated: 2026-09-16
> Purpose: Define component ownership, reuse boundaries, living reusable contracts, and change-impact discipline.

This document governs production React components. It does not replace product/domain specs, UX patterns, visual design authority, or frontend architecture.

## 1. Core principle

Components are organized by **semantic ownership**, not by an assumption that everything should become reusable.

The default question is:

> **Who owns this concept, and how broad is its semantic contract?**

Use the narrowest truthful owner. A component should move toward a broader layer only when broader semantics are demonstrated.

Repetition is evidence worth evaluating for abstraction. It is not proof that abstraction is required.

## 2. Clean-start status

The previous frontend component generation is intentionally retired during the frontend reboot.

At the clean reboot baseline:

- no historical `ui` primitive is automatically considered established;
- no historical `shared` component is automatically considered established;
- old component APIs/variants are not compatibility requirements;
- Git history is archive/evidence, not active component authority.

The new frontend should allow reusable contracts to emerge from real implementation needs.

## 3. Ownership layers

```text
route-specific composition
→ app/<route>/...

domain-semantic component
→ components/features/<domain>/...

genuinely cross-domain semantic component
→ components/shared/...

generic visual/interaction primitive
→ components/ui/...
```

These are placement destinations, not required empty folders.

### Route-local

Use when the composition or behavior belongs to one route/experience and has no independently meaningful wider contract.

Route-local does not mean temporary or low quality.

### Feature/domain

Use when a component represents a stable presentation concept inside one product domain.

Cross-route reuse inside one domain is not evidence that the component belongs in `shared/`.

### Shared semantic

Use only when the **same semantic concept and behavioral contract** genuinely spans unrelated domains.

Using similar markup twice is not enough.

### UI primitive

Use for generic visual/interaction building blocks with no Kencleng business-domain awareness.

A primitive may encode approved design-system behavior, accessibility behavior, and stable interaction/visual variants, but must not redefine domain meaning.

## 4. Pattern vs component vs domain truth

```text
Domain/API authority
→ what the product means

UX pattern
→ how a recurring experience behaves

Component contract
→ how production UI exposes reusable behavior

Component implementation
→ how that contract is realized in code
```

Do not let component architecture become product authority.

## 5. Living reusable-contract registry

This README tracks only the broad-contract layers:

- `components/ui/`;
- `components/shared/`.

It does not enumerate route-local or feature components.

### Registry fields

When a reusable contract is established, record:

```text
Component
Layer
Purpose
Contract maturity
Dedicated contract
```

### Contract maturity

**Evolving** — production usage exists, but the contract may change as evidence develops.

**Stable** — semantics and main behavior are established; changes should preserve supported usage unless an intentional break is approved.

**Foundation** — highly reused contract with broad blast radius.

Maturity describes contract confidence, not code quality.

### Current registry

**No production `ui` or `shared` contracts are established for the new frontend generation yet.**

This empty state is intentional.

When the first new `ui` or `shared` contract is introduced, add it to this registry in the same change.

Do not repopulate this table from retired source/history without new implementation evidence.

## 6. When a dedicated component contract is warranted

Do not create one markdown document per component.

A dedicated contract becomes useful when one or more are true:

- meaningful cross-domain semantics exist;
- variants encode non-obvious product/design intent;
- callers must preserve important accessibility behavior;
- responsive behavior is part of the reusable contract;
- regressions have repeatedly occurred across consumers;
- the consumer surface is large/heterogeneous;
- important forbidden usage cannot be understood safely from types/tests alone;
- the component establishes a stable precedent future agents are likely to extend.

Possible location:

```text
frontend/docs/components/
```

Create that directory only when a real dedicated contract exists.

## 7. Change classification for broad contracts

Before materially changing `components/ui/` or `components/shared/`, classify the change.

### Internal

Implementation-only, no intended observable contract change.

### Additive

Adds an optional capability without changing existing supported defaults/semantics.

### Visual

Changes observable styling/spatial behavior without intentionally changing semantic meaning.

### Behavioral

Changes interaction or state behavior.

### Semantic

Changes what the component communicates or means.

### Breaking

Existing consumers must change to preserve correctness.

The classification determines the breadth of downstream verification; compilation alone is never proof of visual/behavioral/semantic compatibility.

## 8. Consumer impact discipline

For a material broad-contract change:

```text
understand requested change
→ classify change
→ discover actual consumers/wrappers
→ identify representative risk cases
→ implement
→ verify the component contract
→ verify representative downstream consumers
```

Do not maintain a stale manual consumer list. Discover consumers from the live repository when the change happens.

The goal is not to render every call site. Choose the smallest representative set capable of exposing likely regressions.

## 9. Primitive discipline

Do not solve feature-specific presentation by continuously expanding primitive APIs.

Bad direction:

```text
Button variant="campaignCardTopRightUrgent"
```

Better direction:

```text
feature composition owns feature-specific semantics
+
primitive keeps a truthful generic contract
```

Likewise, do not create a new primitive solely because a design token exists. Real composition should demonstrate the useful boundary.

## 10. Shared semantic discipline

`components/shared/` must not become a dumping ground for anything used twice.

Before promoting a concept to `shared/`, establish:

```text
same semantic concept
+
same behavioral contract
+
cross-domain use
```

If those conditions are not true, keep the component with the narrower owner.

## 11. Feature component discipline

Feature/domain components may know domain vocabulary and consume domain-shaped API data, but they must not recreate backend business authority.

Valid:

```text
render campaign progress supplied by the contract
```

Invalid:

```text
recalculate financial eligibility and decide whether a backend action is allowed
```

If presentation requires product data the API does not expose, surface the contract gap.

## 12. Documentation rule

Code owns mechanically discoverable facts. Documentation owns semantics and decisions that cannot safely be reconstructed from code alone.

Do not duplicate props/types into markdown when source already answers the question reliably.

Reusable truth learned through implementation should be promoted here or into the appropriate owning authority rather than left only inside a Harscode task artifact.

## 13. Relationship to the reboot

`docs/project/frontend-reboot-plan.md` owns the one-time retirement/reset operation.

Once the clean reboot baseline is frozen, this document governs reusable component contracts for the new generation.