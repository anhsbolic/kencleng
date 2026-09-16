# Kencleng — UI/UX Authority Map

> Status: Canonical
> Selected direction: **Sunlit Editorial**
> Core thesis: **Evidence-Led Optimism**

This directory contains the active product-design authority for Kencleng.

The active tree is intentionally kept free from superseded design generations. Historical material remains available through Git history rather than as competing current authority.

## Authority by concern

| Concern | Authority |
|---|---|
| Product/domain semantics | `docs/spec/<domain>/...` |
| API/server contract | `api/openapi.yaml` |
| Stable product-design principles | `product-design-principles.md` |
| Selected brand + Product UI direction | `brand-product-ui-brief.md` |
| Concrete visual system | `design-guidelines.md` |
| Reusable interaction behavior | `patterns.md` |
| Asset truthfulness, approval, lifecycle | `asset-governance.md` |
| Persona/surface inventory | `page-map.md` |
| Exploration rationale/history | `exploration/2026-09-product-brand/` |
| Approved direction-level visual evidence | `visual-references/selected-direction/` |

## Visual-system authority

`design-guidelines.md` is the canonical concrete visual-system authority derived from the approved Sunlit Editorial direction.

It owns the currently approved reusable rules for:

- exact core color roles and values;
- Newsreader / Instrument Sans typography roles and scales;
- spacing and density posture;
- surfaces, borders, radius, and elevation;
- Phosphor utility-icon baseline;
- action hierarchy;
- provenance/truth presentation grammar;
- funding vs operational vs reported-outcome visual grammar;
- campaign imagery and placeholder treatment;
- public-vs-product expressive intensity.

Some design decisions remain intentionally OPEN inside that document, including final logo/wordmark, detailed photography/illustration specifications, exact motion tokens, final provenance terminology, and any implementation-discovered visual tokens not yet justified.

Do not treat existing frontend CSS/token values as higher authority than `design-guidelines.md`.

## Precedence

For the same concern, product/domain truth always outranks visual/design artifacts.

Visual references do not establish business rules, API fields, permissions, lifecycle semantics, verification guarantees, or real-world impact.

When documents appear to conflict, resolve the question through the authority that owns that concern rather than selecting whichever artifact is easiest to implement.

## Current direction

Kencleng should feel like a thoughtful, transparent, optimistic donation platform where trust comes from visible evidence, clear progress, respectful storytelling, and consistent truthfulness — not emotional pressure or symbolic trust theatre.

Short form:

> **Evidence-Led Optimism**

Supporting phrase:

> **Hope, structured by evidence.**

## Selected visual references

Approved public-composition and Evidence-Journal studies are stored in:

```text
visual-references/selected-direction/
```

The folder currently contains the approved reference images plus its interpretation/authority README.

The visual references demonstrate design thesis, brand character, hierarchy, and the relationship between public and authenticated surfaces. They are not pixel-perfect specifications or component contracts.

The Concrete Visual System exploration additionally validated the foundations and evidence/progress grammar now promoted into `design-guidelines.md`. The system rules in the canonical document own reusable visual behavior; exploration imagery remains supporting evidence rather than screenshot specification.

## Historical design generations

Superseded prototype exports and prior visual-system documents are intentionally not kept in the active tree. Git history is the archive for those materials.
