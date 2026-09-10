# Kencleng — Brand and Visual Assets

> Intended path: `docs/ui-ux/brand-and-visual-assets.md`
>
> Status: Draft
>
> Purpose: Define how Kencleng uses brand identity, iconography, illustrations, placeholders, expressive imagery, and generated visual assets.
>
> This document governs visual intent and asset authority. It does not define component implementation details or replace `design-guidelines.md`.

## Context

Kencleng should feel:

**warm + trustworthy + transparent + calm**

The product handles money, organizational identity, verification, campaign accountability, and public trust. Visual expression should therefore make the product more human and distinctive without weakening clarity or implying facts the product cannot prove.

The visual system should avoid two common failure modes:

1. technically correct but visually generic interfaces built mostly from default library icons and stock-like decoration;
2. expressive interfaces whose illustrations or imagery imply unsupported trust, urgency, popularity, beneficiaries, or real-world outcomes.

Principle:

> **Use standard visual conventions for standard utility. Create intentional visual identity where expression, storytelling, trust, or product meaning benefits from it.**

---

# A. Asset Classes

Before choosing or creating an asset, classify its role.

## Level 1 — Utility

Utility visuals communicate universal interface actions.

Examples:

* close;
* search;
* menu;
* chevron;
* calendar;
* edit;
* delete;
* filter;
* password visibility;
* external link;
* download.

Default:

```text
established icon library
→ use directly
```

Do not create custom artwork merely to make universal controls look branded.

Familiarity is more valuable than novelty here.

The selected icon must still:

* match the action semantics;
* be understandable in context;
* have an accessible name when the icon is the only visible label;
* use established sizing/stroke conventions.

---

## Level 2 — Product-semantic visuals

These represent concepts specific to Kencleng's product domain but may still use standard icon primitives.

Examples:

* organization verification;
* campaign status;
* donation;
* fund usage;
* disbursement;
* representative role;
* accountability/reporting.

Default:

```text
existing Kencleng treatment?
→ reuse it

no established treatment?
→ standard icon may be used when semantically sufficient
→ consider whether a Kencleng-specific treatment is warranted
```

A standard icon does not automatically need replacement.

However, repeatedly important product concepts should develop a consistent visual treatment rather than accumulating arbitrary icons chosen independently on each page.

Consistency may come from:

* icon selection;
* shape/container treatment;
* color semantics;
* badges;
* accompanying labels;
* illustrative motifs.

---

## Level 3 — Expressive product assets

These visuals contribute meaning, emotion, onboarding, storytelling, or product character.

Examples:

* empty-state illustrations;
* campaign placeholders;
* organization placeholders;
* success/completion moments;
* onboarding visuals;
* explanatory illustrations;
* landing-page section graphics;
* decorative background motifs;
* trust/accountability storytelling.

Default:

> **Do not silently fall back to a generic library icon when an expressive asset materially improves the experience.**

When the asset does not yet exist, treat it as a design need.

The current harness may generate the asset, or produce a handoff brief when it cannot.

---

## Level 4 — Brand-defining assets

These materially define Kencleng's identity.

Examples:

* logo;
* symbol;
* wordmark;
* mascot;
* core illustration language;
* primary landing-page hero identity;
* signature graphic motif;
* major brand pattern.

Agents may explore and generate candidates.

They must not unilaterally declare a new Level 4 asset canonical.

Human approval is required before a new brand-defining asset becomes part of the official system.

---

# B. Standard Icons vs Custom Visuals

Library icons are not a design failure.

The failure is using them automatically for every visual need.

Use standard icons when:

* the concept is universally understood;
* familiarity reduces cognitive load;
* the icon supports rather than carries the full experience;
* custom treatment would add novelty without meaning.

Consider a custom treatment or expressive asset when:

* the visual represents an important Kencleng concept repeatedly;
* the moment has meaningful emotional weight;
* illustration can improve comprehension;
* the surface contributes significantly to product identity;
* a generic icon makes the experience feel like an interchangeable template;
* storytelling or context is otherwise missing.

Bad:

```text
Landing hero
→ generic Coins icon enlarged to 160px

First-campaign empty state
→ Folder icon

Donation success
→ CheckCircle icon only

Organization placeholder
→ generic gray Image icon
```

Good:

```text
utility filter action
→ standard Filter icon

campaign status
→ consistent semantic badge/icon treatment

first-campaign empty state
→ established Kencleng empty-state illustration

landing hero
→ intentional expressive asset based on Kencleng's product story
```

---

# C. No Generic Filler

Do not add decorative visuals merely to make a surface feel complete.

Avoid defaulting to:

* arbitrary gradients;
* decorative blobs with no relationship to the visual language;
* generic fintech coins;
* random shield/check illustrations;
* generic handshake imagery;
* stock-like charity photography;
* crypto visual language;
* fake dashboards;
* decorative icon clouds;
* AI-generated people presented as actual beneficiaries.

If no meaningful visual is required, whitespace and typography are valid.

If a meaningful visual is required but missing:

```text
identify asset need
→ classify it
→ generate or hand off
```

Do not hide the gap behind generic filler.

---

# D. Truth and Representation

Visuals must not imply evidence the product does not possess.

Generated or placeholder visuals must not be mistaken for:

* a real beneficiary;
* a real campaign location;
* a real organization;
* evidence that an activity happened;
* actual goods purchased with campaign funds;
* documented campaign impact.

For example, an AI-generated photograph of children must never appear inside a real campaign card in a way that suggests those children are the campaign beneficiaries unless that image is genuine and authorized campaign content.

Principle:

> **Illustration may communicate a concept. It must not fabricate evidence.**

Expressive illustration and clearly stylized placeholders are safer than synthetic documentary-style photography for generic application states.

---

# E. Placeholder System

Missing media must still feel intentional.

Avoid:

```text
gray rectangle
+
broken-image appearance
```

as the permanent product treatment.

Kencleng should develop a consistent placeholder system for at least:

* campaign image unavailable;
* organization image/logo unavailable;
* user/avatar unavailable where applicable;
* content/media still loading;
* corrupted or inaccessible media if distinguishable.

A placeholder should:

* be visually recognizable as a placeholder;
* not impersonate real campaign imagery;
* fit the brand visual language;
* preserve expected aspect ratio/layout stability;
* remain subordinate to real content.

Placeholder assets may use:

* subtle Kencleng geometry;
* an approved brand motif;
* neutral illustrations;
* category-neutral composition;
* restrained use of the brand symbol once approved.

Do not create a fake category image unless the category itself is real and the visual cannot be mistaken for campaign evidence.

---

# F. Empty States

An empty state is both a state and, sometimes, an expressive opportunity.

First determine its role.

## Functional empty state

Example:

```text
Search returned no results
```

A lightweight treatment is usually enough:

* concise message;
* relevant next action;
* optional standard icon.

Do not force illustration into every empty state.

## Motivational or milestone empty state

Example:

```text
An organization has not created its first campaign.
```

This may deserve an expressive asset because the state represents:

* onboarding;
* encouragement;
* product character;
* a meaningful next action.

The visual must support the action rather than replace explanation.

An illustration without useful copy or action is decoration, not UX.

---

# G. Landing and Marketing Imagery

Landing-page visuals carry more brand responsibility than ordinary product screens.

Do not default to:

* floating dashboard screenshots merely because the product has a dashboard;
* generic crowdfunding imagery;
* coins flying into a wallet;
* stock-photo charity clichés.

Before generating a major landing asset, establish an asset brief.

Minimum brief:

```text
Purpose
Audience
Emotional goal
Narrative
Visual role
Composition
Copy-safe area
Responsive behavior
Brand characteristics
What must be avoided
Expected output format
```

Example direction:

```text
Purpose:
Communicate that many small, transparent contributions
can become meaningful collective impact.

Narrative:
A contemporary interpretation of "kencleng":
small contributions visibly accumulating into a shared outcome.

Character:
Warm, trustworthy, optimistic, calm.

Avoid:
crypto aesthetics,
generic floating coins,
stock charity photography,
overly childish cartoon styling.
```

This is a design direction, not a mandatory final concept.

Multiple concepts may be explored before canonical artwork is selected.

---

# H. Logo and Brand Identity

If a canonical Kencleng logo or wordmark does not yet exist, the absence must remain explicit.

Do not allow an incidental implementation placeholder to silently become the logo.

Possible states:

```text
MISSING
PROPOSED
APPROVED
CANONICAL
SUPERSEDED
```

A text wordmark may be used provisionally when needed for implementation, but it should remain identified as provisional until approved.

When an agent identifies that a surface materially depends on a missing logo or brand-defining asset:

```text
flag design gap
→ propose exploration
→ generate candidates or prepare handoff
→ human approval
→ mark canonical asset
→ integrate
```

Logo exploration should consider:

* relationship to the word "Kencleng";
* donation/community meaning without relying on cliché;
* recognizability at small sizes;
* symbol-only and wordmark usage;
* monochrome usage;
* dark/light background behavior;
* favicon/app-icon suitability;
* Indonesian cultural relevance where natural, without forced ornamentation.

A logo must not be declared final because it looks acceptable in one header screenshot.

---

# I. Capability-Aware Asset Production

Asset responsibility is independent from tool capability.

The agent always owns the responsibility to:

```text
detect need
→ reason about intent
→ classify asset
→ prepare design brief
→ create or hand off
→ evaluate result
→ integrate after appropriate approval
```

How the asset is produced depends on the current harness.

## When the harness can create the asset

If the current tool can adequately generate or edit the required asset:

```text
brief
→ generate candidate(s)
→ inspect against brief
→ revise if needed
→ obtain required approval
→ integrate
```

The ability to generate an asset does not grant authority to make every asset canonical.

Human approval rules still apply.

---

## When the harness cannot create the asset

Do not silently downgrade the design.

Produce an **asset handoff package** that the human can use in another tool.

Minimum handoff:

```text
Asset name
Classification
Purpose
Surface/context
Visual role
Concept
Style direction
Composition
Required states/variants
Must preserve
Must avoid
Expected dimensions/aspect ratio
Expected format
Generation prompt
Approval requirement
```

Example:

```text
Asset:
Landing hero illustration

Classification:
Level 4 — brand-defining candidate

Purpose:
Communicate collective generosity and transparent impact.

Composition:
Wide hero composition.
Primary visual weight on the right.
Keep the left region low-detail for headline/CTA.
Must adapt to narrow screens without losing the core metaphor.

Must preserve:
trust,
warmth,
clarity,
community.

Avoid:
generic coins,
crypto aesthetics,
stock charity photography,
fake real-world beneficiaries,
overly childish cartoon treatment.

Expected output:
Primary vector/SVG when feasible;
high-resolution transparent PNG acceptable for exploration.

Generation prompt:
[ready-to-paste generation prompt]

Approval:
Human approval required before canonical use.
```

The human may use any suitable external design/generation tool.

The intent should survive the tool change.

---

# J. Generation Prompts Are Design Artifacts

For significant expressive or brand assets, a successful generation prompt or visual brief may be worth preserving.

Do not persist every exploratory prompt.

Persist when the asset:

* is canonical or likely to be reused;
* represents a repeatable illustration style;
* is difficult to reconstruct from the final file alone;
* may need future regeneration;
* has meaningful negative constraints.

Suggested location:

```text
docs/ui-ux/assets/
```

Possible structure:

```text
docs/ui-ux/assets/
├── README.md
├── landing-hero.md
├── campaign-placeholder.md
└── first-campaign-empty-state.md
```

Each significant asset brief may contain:

```text
Status
Purpose
Classification
Approved asset path
Visual brief
Generation prompt
Negative constraints
Variants
Approval notes
```

Do not create a documentation file for every ordinary icon.

---

# K. Asset Status and Authority

Use explicit status for expressive and brand-level assets.

## PROVISIONAL

Temporary but intentional.

May be used to unblock implementation.

Must not become the visual precedent for future work unless promoted.

## APPROVED

Reviewed for the intended usage.

May be used on the approved surface.

## CANONICAL

Part of the established Kencleng visual system.

Agents should reuse or extend it before introducing a competing visual direction.

## SUPERSEDED

No longer the current reference.

Keep historical design context only when useful; do not reuse it for new implementation.

A file existing in the repository does not automatically mean it is canonical.

---

# L. Extending an Existing Asset System

Once Kencleng develops a stable illustration or placeholder language, new assets should extend that language rather than restart art direction.

Before introducing a new visual style, inspect existing canonical assets for:

* shape language;
* line/stroke characteristics;
* perspective;
* dimensionality;
* detail density;
* color behavior;
* human-character treatment;
* background treatment;
* texture;
* visual metaphor;
* icon/illustration relationship.

Ask:

> **Would this asset look like it belongs in the same product if the surrounding UI were removed?**

If no, either adapt it to the established system or explicitly propose a system-level evolution.

---

# M. Shared Asset Change Impact

Canonical visual assets can have blast radius similar to shared components.

Examples:

* logo;
* brand symbol;
* campaign placeholder;
* global background motif;
* verification graphic;
* shared empty-state illustration;
* product-semantic icon treatment.

Before materially changing a shared/canonical asset:

```text
identify consumers
→ determine visual/semantic impact
→ inspect representative usages
→ assess dark/light/responsive variants
→ update dependent assets or surfaces
→ record intentional contract change
```

Changing a shared asset in isolation is insufficient when the asset appears across unrelated surfaces.

Do not maintain a manually curated list of every consumer when repository search can discover usage more reliably.

The documentation should preserve semantics and authority; tooling should discover actual consumers at change time.

---

# N. Asset Change Classification

Classify a change before integrating it.

## Additive

Adds a new asset without altering existing semantics.

Example:

```text
new empty-state illustration for a previously uncovered state
```

## Visual-compatible

Updates rendering while preserving recognizable meaning and expected usage.

Example:

```text
refining line treatment of an established placeholder
```

Still inspect representative consumers.

## Semantic

Changes what an asset communicates.

Example:

```text
changing a generic verification symbol into an "approved organization" symbol
```

Requires product/design review because meaning changes.

## Brand-breaking

Changes a core identity or visual-language contract.

Examples:

* replacing the logo;
* changing illustration style globally;
* changing the primary brand motif;
* materially changing identity colors or shape language.

Requires explicit human approval and broad impact analysis.

---

# O. Generated Asset Review

Generated output must be reviewed before integration.

Check:

### Intent

* Does it communicate the requested idea?
* Is the visual role appropriate for the surface?

### Truth

* Does it imply unsupported facts?
* Could it be mistaken for real evidence?

### Brand

* Does it fit Kencleng's established visual personality?
* Does it avoid generic fintech/charity visual clichés?

### Composition

* Does it work in the actual container?
* Is text-safe space preserved when required?
* Can it adapt responsively?

### Technical quality

* Is resolution sufficient?
* Is transparency clean where required?
* Is the vector structurally usable if SVG is expected?
* Are unnecessary embedded raster assets or metadata avoided?
* Is the file size appropriate?

### Accessibility

* Is the visual decorative or informational?
* If informational, is equivalent accessible meaning available?
* Does important information exist independently of the illustration?

### Reuse

* Does the asset create a precedent?
* If so, is that precedent desirable?

---

# P. Avoiding AI Visual Homogeneity

Agents should actively avoid converging on the same default visual vocabulary common to generated interfaces.

Watch for repeated defaults such as:

```text
Lucide icon in rounded square
+
gradient blob
+
large centered heading
+
generic card grid
```

These patterns are not forbidden.

They are simply not sufficient justification by themselves.

Before accepting an expressive surface, ask:

* Is this visual choice serving Kencleng specifically?
* Could the exact same surface belong to a random SaaS, fintech, or donation template?
* Is there an opportunity for meaningful product identity?
* Would greater visual expression reduce clarity or trust?

Distinctiveness should come from coherent product meaning, not arbitrary novelty.

---

# Q. Relationship to Other Documents

This document answers:

> **What visual assets should exist, how expressive should they be, and who may establish them as part of Kencleng's visual identity?**

Related authority:

* `product-design-principles.md` — product personality, trust, hierarchy, design readiness, and design authority
* `design-guidelines.md` — colors, typography, spacing, shapes, elevation, and visual component treatment
* `patterns.md` — UX structures and interaction behavior
* `prototype-reference.md` — visual reference authority per surface
* `design-reference-usage.md` — translation from prototype to production
* `frontend/components/README.md` — component ownership and shared-component contracts

If an existing asset conflicts with product/domain truth, product/domain truth wins.

If a prototype introduces a new visual language that conflicts with established canonical assets, treat it as a proposed visual direction rather than silently replacing the existing system.

---

# R. Evolution

The asset system should evolve from evidence, not novelty.

Update this guidance when repeated work reveals:

* a visual class that needs a stable policy;
* inconsistent icon semantics;
* recurring placeholder needs;
* a mature illustration language;
* repeated generation/handoff friction;
* assets that frequently create downstream regressions;
* ambiguity about whether a visual is provisional or canonical.

Do not create a new asset category or process for a single one-off file.
