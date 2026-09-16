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
| Reusable interaction behavior | `patterns.md` |
| Asset truthfulness, approval, lifecycle | `asset-governance.md` |
| Persona/surface inventory | `page-map.md` |
| Exploration rationale/history | `exploration/2026-09-product-brand/` |
| Approved direction-level visual evidence | `visual-references/selected-direction/` |

## Important absence: concrete visual-system specification

There is currently **no canonical concrete visual-system document** defining exact production:

- color values;
- font families;
- radii;
- spacing tokens;
- shadows;
- icon family;
- motion values.

That absence is intentional.

The brand/product direction has been approved, but the production visual system must still be deliberately derived from it. Existing implementation values must not be mistaken for design authority.

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

The public-composition and Evidence-Journal studies approved in the design session are **Selected Direction References**.

Their authority and intended filenames are recorded in:

```text
visual-references/selected-direction/README.md
```

The binary image handoff is explicitly pending because the current repository connector cannot attach the available binary file handles directly. This does not reopen the approved direction and must not be filled with substitute artwork without review.

The visual references demonstrate design thesis, brand character, hierarchy, and the relationship between public and authenticated surfaces. They are not pixel-perfect specifications or component contracts.

## Historical design generations

Superseded prototype exports and prior visual-system documents are intentionally not kept in the active tree. Git history is the archive for those materials.
