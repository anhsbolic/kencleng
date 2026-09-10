# Kencleng — UI/UX Prototype Reference

> Intended path: `docs/ui-ux/prototype-reference.md`
>
> Status: Draft v2
>
> Purpose: Define which prototype/reference artifacts exist, how much authority they carry, and how they relate to current product, UX, visual-system, asset, and production-component truth.

## Core Principle

A prototype is:

> **route-specific design evidence and visual precedent**

It is not automatically:

* product truth;
* API truth;
* component architecture;
* state-management architecture;
* canonical visual-system definition;
* canonical asset definition;
* production code.

A visually complete prototype may still contain illustrative data, provisional assets, outdated tokens, or implementation shortcuts.

Use prototypes to preserve **design intent**, not prototype implementation accidents.

---

# A. Authority by Concern

There is no single flat source-of-truth order for every frontend question.

Different documents own different concerns.

| Concern                         | Primary authority                                     |
| ------------------------------- | ----------------------------------------------------- |
| Business/domain semantics       | `docs/spec/<domain>/...`                              |
| API shape and server contract   | `api/openapi.yaml`                                    |
| Product-design principles       | `product-design-principles.md`                        |
| Reusable UX behavior            | `patterns.md`                                         |
| Visual system                   | `design-guidelines.md`                                |
| Brand / visual assets           | `brand-and-visual-assets.md`                          |
| Route/persona inventory         | `page-map.md`                                         |
| Component ownership/contracts   | `frontend/components/README.md` + dedicated contracts |
| Route-specific visual precedent | this document + `design-reference/`                   |

When a prototype conflicts with an authority that owns the relevant concern, the owning authority wins.

Examples:

```text
prototype mock field
vs OpenAPI
→ OpenAPI wins

prototype component decomposition
vs production semantic ownership
→ component system wins

prototype spacing/color
vs current visual-system token
→ current design guideline wins

prototype placeholder illustration
vs approved canonical asset
→ canonical asset wins
```

Do not silently resolve a genuine product/domain contradiction.

Surface it.

---

# B. Current Reference Artifacts

The current artifacts under:

```text
design-reference/
```

were exported from Claude Design as standalone prototype output.

They remain:

* frozen;
* read-only for agents;
* disposable as implementation code;
* useful as rendered visual/structural reference.

Per root `AGENTS.md`, agents may inspect them but must not modify the directory or wholesale-copy its implementation into `frontend/`.

Future design references may originate from other design or generation tools.

Their authority is determined by this document and their approval/status — **not by which tool created them**.

---

# C. Tier 1 — Route-Specific Visual Precedent

Tier 1 means:

> A dedicated prototype exists for this route and should strongly inform route-specific composition and visual intent.

It does **not** mean:

> Copy the prototype as literally as possible.

Use Tier 1 reference primarily for:

* information hierarchy;
* major composition;
* relative visual emphasis;
* intended states;
* interaction affordances;
* responsive intent where represented;
* copy/microcopy as a candidate;
* intended visual character.

Translate it through the current canonical:

```text
product principles
+
UX patterns
+
visual system
+
brand/asset system
+
component architecture
```

Current Tier 1 routes:

| Route                                       | Pattern / role                              |
| ------------------------------------------- | ------------------------------------------- |
| `/`                                         | Landing / one-off expressive public surface |
| `/login`                                    | Form — authentication variant               |
| `/campaign`                                 | List / Browse                               |
| `/campaign/[id]`                            | Detail — public variant                     |
| `/campaign/[id]/donate`                     | Form — donation                             |
| `/dashboard/campaign/new`                   | Form — Revisable Submission                 |
| `/dashboard/campaign/[id]/monitor`          | Dashboard / Summary                         |
| `/dashboard/kurasi/campaign/[assignmentId]` | Curation / Review                           |
| `/donation/[id]/status`                     | Status / Tracking                           |
| `/dashboard/organization/new`               | Form — Revisable Submission                 |

The reference set also contains non-route component/layout sheets.

Those sheets are **visual precedents**, not production component specifications.

---

# D. Tier 2 — Pattern-Derived Surfaces

Tier 2 routes have no dedicated route prototype.

Do not interpret that absence as:

```text
"design however you want"
```

Instead:

```text
page-map
→ applicable UX pattern
→ closest relevant precedent
→ current design system
→ current component system
```

Use existing Tier 1 surfaces as comparative precedent where useful, but do not literally clone content or role-specific composition.

Examples:

### List / Browse

Useful precedent:

```text
/campaign
```

Possible consumers include administrative queues, donations, notifications, representatives, and other collection surfaces.

### Detail

Useful precedent:

```text
/campaign/[id]
```

Dashboard/operational detail surfaces must adapt hierarchy to their persona and task.

### Form

Useful precedents:

```text
/dashboard/campaign/new
/login
```

depending on dashboard vs authentication context.

### Curation / Review

Useful precedent:

```text
/dashboard/kurasi/campaign/[assignmentId]
```

### Status / Tracking

Useful precedent:

```text
/donation/[id]/status
```

A Tier 2 feature may still be classified **OPEN** under design readiness when the nearest pattern does not adequately answer its UX problem.

Tier 2 does not mean agent improvisation without design reasoning.

---

# E. Prototype Freshness

Prototype authority is not permanent merely because an artifact exists.

A reference may become partially stale when:

* product requirements change;
* UX patterns evolve;
* visual-system rules change;
* a component contract changes;
* canonical assets replace provisional visuals;
* accessibility/responsive verification reveals a defect.

Use these conceptual statuses where useful:

```text
CURRENT
NEEDS_REVALIDATION
SUPERSEDED
```

A stale prototype may remain valuable for composition or historical rationale while no longer being authoritative on a changed concern.

Do not regenerate every prototype whenever a token changes.

Update or supersede references when the visual/design intent itself has materially changed.

---

# F. Current Migration Note

The current reference exports predate the Product Design / Visual System v2 work.

Therefore they should currently be treated as:

> **strong composition precedent requiring translation through the current design system**

rather than literal final snapshots.

In particular:

* current visual tokens override prototype token drift;
* the v2 Secondary-action direction is neutral/outlined rather than treating amber as the default filled secondary action;
* canonical visual assets, once approved, override older generic/provisional placeholders;
* component architecture must follow semantic ownership rather than prototype decomposition.

Do not update frozen `design-reference/` files merely to reflect these changes.

Production implementation should translate them correctly.

---

# G. Known Prototype Issues

Known defects must not become production precedent merely because they are visible in a Tier 1 artifact.

## Login request-level error

The login prototype has historically represented a generic authentication failure too closely to a field-level email validation error.

Production behavior must preserve the UX/security distinction between:

```text
field validation
```

and:

```text
request-level authentication failure
```

according to the current Form/Error patterns.

## Campaign image placeholder

Public campaign surfaces contain a prototype placeholder that resembles an upload affordance.

A read-only public campaign card/detail must not imply that the viewer can upload media there.

Use the canonical placeholder behavior from the visual-asset system once established.

## Typography drift

Prototype typography does not perfectly match the canonical production type scale.

`design-guidelines.md` owns current typography values.

Do not copy prototype font sizes literally.

---

# H. Prototype vs Product Truth

Prototype data is illustrative unless confirmed by product/API authority.

Do not infer from a mockup that Kencleng supports:

* a backend field;
* ranking;
* sorting;
* status;
* verification level;
* permission;
* calculation;
* recommendation;
* financial state;
* analytics metric.

Example:

```text
prototype shows "Featured"
```

does not establish a product concept called:

```text
featured
```

The product/domain authority must support it.

Principle:

> **Visual completeness is not evidence of domain truth.**

---

# I. Prototype vs Component Architecture

Prototype component boundaries are exploratory evidence.

They may reveal useful responsibilities.

They do not dictate production extraction.

For example, a prototype may contain:

```text
AmountField
MethodGrid
SummaryStrip
```

These are useful clues about conceptual responsibilities.

Production engineering must still ask:

```text
Is this responsibility meaningful?
Who semantically owns it?
Is an existing component contract already available?
Is extraction actually useful?
```

Do not mirror the prototype component tree by default.

Production components follow:

```text
frontend/components/README.md
```

and current Harscode frontend engineering guidance.

---

# J. Prototype vs Visual System

A Tier 1 prototype strongly informs composition and overall intent.

The canonical visual system owns recurring visual rules such as:

* colors;
* typography;
* radii;
* elevation;
* action hierarchy;
* spacing conventions;
* status semantics;
* icon treatment.

If the prototype differs because it predates an approved visual-system change:

```text
preserve intent
→ use current system
```

If the prototype appears intentionally to propose a **new system-level direction**:

```text
do not normalize silently
→ classify as design proposal
→ review
→ approve before canonicalizing
```

---

# K. Prototype vs Brand Assets

An image or illustration appearing in a prototype is not automatically canonical.

Determine whether it is:

```text
placeholder
provisional asset
approved asset
canonical asset
```

according to `brand-and-visual-assets.md`.

An agent must not reproduce a generic temporary visual merely because it appears in the reference.

If a route materially needs an expressive asset and no canonical asset exists:

```text
design gap
→ asset brief
→ generate or hand off
→ appropriate approval
→ integrate
```

---

# L. Reference Comparison

Rendered production UI does not need arbitrary pixel identity with a prototype.

Compare primarily:

* hierarchy;
* composition;
* relative emphasis;
* interaction intent;
* state coverage;
* responsive behavior;
* visual-system consistency;
* brand character.

Small implementation differences are acceptable when they improve:

* accessibility;
* robustness;
* current design-system consistency;
* realistic content handling;
* responsive behavior.

A large visual deviation from a Tier 1 reference should be intentional and explainable.

---

# M. Design Readiness

Prototype coverage contributes to, but does not determine, design readiness.

A Tier 1 route can still be **PARTIAL** or **OPEN** if major requirements have changed since the prototype.

A Tier 2 route can be **READY** when established patterns and precedents answer the problem completely.

Use the readiness definitions from:

```text
product-design-principles.md
```

Do not equate:

```text
prototype exists = READY
prototype missing = OPEN
```

---

# N. Adding Future References

Do not create a dedicated prototype for every route by default.

Create or preserve a new route-specific reference when it materially helps resolve:

* new interaction architecture;
* important product storytelling;
* major responsive composition;
* new visual-system direction;
* high-risk or benchmark-sensitive UI;
* important reusable precedent.

Avoid recreating a full screenshot/wireframe inventory that becomes expensive to maintain.

Patterns and canonical systems should carry repeated knowledge.

Prototypes should carry **high-value design precedent**.

---

# O. Related Documents

* `product-design-principles.md`
* `patterns.md`
* `design-guidelines.md`
* `brand-and-visual-assets.md`
* `page-map.md`
* `design-reference-usage.md`
* `frontend/components/README.md`
* root `AGENTS.md`
