# Kencleng — Visual Design Guidelines

> Status: Canonical
> Selected direction: **Sunlit Editorial**
> Core thesis: **Evidence-Led Optimism**
> Upstream authority: `brand-product-ui-brief.md`
> Human approval: Approved 2026-09-16
>
> This document defines the concrete visual-system language derived from the approved Kencleng brand and product UI direction. It does not define business/domain semantics, API contracts, frontend component architecture, or implementation-specific CSS mechanics.

## 1. Purpose

This document answers:

> **How should the approved Kencleng design direction consistently look and behave as a visual system?**

It translates the approved upstream direction into reusable visual rules for color, typography, spacing, surfaces, borders, radius, elevation, iconography, action hierarchy, provenance, progress grammar, campaign imagery/placeholders, and public-vs-authenticated expression.

The goal is not visual uniformity. The goal is **coherent visual behavior**.

## 2. System Posture

Kencleng should feel:

- light;
- warm;
- editorial;
- human;
- structured;
- candid;
- optimistic without cheerfulness;
- refined without exclusivity;
- distinctive without decorative excess.

Avoid drifting toward generic SaaS, fintech blue, charity green, orange-dominant charity familiarity, lifestyle/wellness softness, luxury-editorial distance, startup gradients, childish optimism, or institutional coldness.

Useful shorthand:

> **Warm editorial expression, disciplined product structure.**

## 3. Core Visual Principles

### 3.1 Neutrals carry the product

Most of the interface should be carried by warm neutrals, Ink, spacing, typography, borders, and hierarchy. Brand color should not be required for every surface to feel like Kencleng.

### 3.2 Sun carries energy

Yellow communicates optimism, momentum, editorial emphasis, selected action hierarchy, and funding progress.

Yellow does **not** communicate truth, verification, success, trust, or safety.

### 3.3 Berry carries depth

Berry adds warmth, emotional maturity, secondary emphasis, and editorial distinction. It should remain selective and should not dominate authenticated product UI.

### 3.4 Trust comes from structure

Trust should emerge through information hierarchy, provenance, timestamps, source labeling, chronology, known-vs-unknown distinction, factual wording, and predictable system behavior.

Avoid relying on green coloring, shield icons, checkmark abundance, “verified” badges, or decorative trust symbols.

### 3.5 Evidence is structure; optimism is expression

Evidence determines what can be communicated. Optimism determines how confidently and humanely the system presents that truth. Optimism must never upgrade the certainty of evidence.

## 4. Color System

### 4.1 Core Palette

| Role | Value | Primary use |
|---|---|---|
| Canvas | `#FFFDF7` | page background |
| Warm Surface | `#FFF8E8` | editorial sections, gentle emphasis |
| Subtle Surface | `#F6F0E3` | low-emphasis grouping |
| Ink | `#24211D` | primary text / structural authority |
| Muted Ink | `#5F5A52` | supporting text / metadata |
| Border | `#E2D8C7` | quiet borders and dividers |
| Sun | `#F6C945` | primary brand energy |
| Sun Strong | `#E5B62E` | active / stronger emphasis |
| Berry | `#A63B57` | mature secondary accent |
| Berry Soft | `#F7E8EC` | tonal secondary accent |

### 4.2 Semantic Colors

| State | Value |
|---|---|
| Success | `#2F7A57` |
| Warning | `#A86616` |
| Error | `#B74343` |
| Info | `#426B8E` |

Semantic color communicates actual state, not brand personality.

```text
Sun ≠ success
Berry ≠ warning
Green ≠ trusted
```

Use copy + state + structure together.

## 5. Typography

### 5.1 Font Pairing

**Newsreader** — editorial/display voice.

Use for:

- public hero;
- campaign editorial headline;
- selected public section titles;
- major storytelling statements;
- selected milestone statements.

**Instrument Sans** — UI/body/operational voice.

Use for:

- navigation;
- body;
- forms;
- labels;
- metadata;
- buttons;
- money;
- status;
- dashboard;
- timeline;
- structured data.

### 5.2 Typography Rules

Public surfaces may use stronger editorial contrast between Newsreader and Instrument Sans. Do not turn every heading into serif merely to display brand character.

Authenticated product surfaces should be carried almost entirely by Instrument Sans. Newsreader may appear only in selected human/editorial moments.

Always prefer Instrument Sans for currency, percentages, date/time, state labels, tables, metadata, provenance, form values, and other operational content.

Where supported, use tabular figures for aligned numeric content.

### 5.3 Public Scale

```text
Display XL    64 / 68
Display       52 / 58
H1            40 / 48
H2            32 / 40
H3            24 / 32
Body Large    18 / 28
Body          16 / 24
Small         14 / 20
Caption       12 / 16
```

### 5.4 Product Scale

```text
Page H1       32 / 40
Section H2    24 / 32
Section H3    20 / 28
Body          16 / 24
UI / Meta     14 / 20
Caption       12 / 16
```

Exact responsive interpolation is implementation freedom.

## 6. Spacing and Density

Base unit:

```text
4px
```

Common public rhythm:

```text
16
24
40
64
96
```

Common product rhythm:

```text
8
12
16
24
32
48
```

Public layouts should feel spacious, calm, editorial, and generous.

Product layouts should feel structured, efficient, readable, and never cramped.

## 7. Surfaces and Grouping

Kencleng is **border-first and spacing-first**, not card-first.

Use page hierarchy, spacing, border/rule, tonal background, and section grouping before creating a card container.

A card is justified when content has meaningful bounded ownership.

Avoid turning every statistic, form section, metadata row, status, heading, or list item into a separate rounded card.

## 8. Radius

Canonical radius family:

```text
6px    compact objects
10px   inputs / ordinary controls
14px   standard grouped surfaces
20px   selected expressive public surfaces
999px  genuine pills / circular treatments only
```

Avoid defaulting to extreme rounding. Kencleng should feel human and approachable, not bubbly.

## 9. Elevation

Use only three conceptual levels:

### Flat

Default. No shadow.

### Raised Subtle

For dropdowns, small floating surfaces, and selected interactive panels.

### Overlay

For modal, popover, or blocking overlay.

Elevation communicates layering, not decoration.

## 10. Borders and Rules

Use quiet warm borders for grouping, boundaries, comparison, timeline structure, and form segmentation.

Dividers should generally be lower contrast than primary content.

Avoid excessive boxed segmentation.

## 11. Action Hierarchy

Action hierarchy is contextual.

It is not defined as:

```text
primary = yellow forever
```

Instead:

```text
current task
+
consequence
+
importance
→ action hierarchy
```

### Public primary CTA

Sun background + Ink text is an approved primary treatment when the action genuinely deserves dominant attention and the surrounding composition remains restrained.

### Product primary actions

Product UI may use Sun, Ink-led solid treatment, neutral outlined treatment, or tonal emphasis depending on context.

### Secondary actions

Prefer outlined, neutral, or lower-emphasis tonal treatment.

### Destructive actions

Use explicit Error semantics. Do not use Berry merely because it is red-adjacent.

## 12. Iconography

Baseline utility library: **Phosphor Icons**.

Use for navigation, utility actions, standard UI semantics, and ordinary supporting cues.

Icons support meaning. They do not replace text, state, provenance, or explanation.

Do not overuse shield, checkmark, lock, or certification-like symbols to manufacture trust.

Brand recognition should primarily come from typography, composition, color roles, editorial rhythm, photography, illustration, and evidence grammar.

## 13. Provenance and Truth Grammar

These are presentation classes, not necessarily backend domain states.

### Platform Fact

- plain structure;
- strong Ink-led hierarchy;
- minimal decoration;
- no celebratory badge requirement.

Purpose: communicate something the platform directly knows.

### Organizer-Provided Information

- visible source attribution;
- supporting metadata where useful.

Conceptual label:

```text
Dari pengelola
```

Exact production terminology may evolve.

### Organizer Report

- explicit report/source label;
- timestamp;
- source context;
- clear distinction from platform-confirmed fact.

Conceptual label:

```text
Laporan pengelola
```

### Pending / Not Yet Available

- neutral/muted treatment;
- explicit plain language;
- visible, not hidden;
- not automatically warning or error.

Examples:

```text
Belum tersedia
Belum dilaporkan
Menunggu pembaruan
```

### System State

Use actual semantic-state treatment for real product/system states.

### Reported Outcome

- sourced evidence/report surface;
- attribution visible;
- timestamp/date visible where available;
- supporting documentary evidence optional.

Do not visually imply independent verification unless product/domain truth supports that claim.

## 14. Progress as Evidence

Progress is divided into three distinct visual grammars. They must not collapse into one universal progress component.

### 14.1 Funding Progress

Purpose: show quantitative fundraising progress.

Preferred visual:

```text
amount raised
+
target
+
percentage where useful
+
horizontal progress bar
```

Sun may be used as fill.

Funding progress must not imply execution, distribution, impact, or beneficiary outcome.

### 14.2 Operational Progress

Purpose: show execution or process milestones.

Preferred visual:

- timeline;
- milestone sequence;
- chronological steps.

Do not use percentage bars unless actual domain data defines meaningful quantitative operational progress.

### 14.3 Reported Outcome

Purpose: communicate results reported by an appropriate source.

Preferred visual:

- report/evidence block;
- explicit provenance;
- update date;
- optional supporting documentary image.

Do not use a generic progress bar.

### 14.4 System invariant

```text
Funding progress
≠
Operational progress
≠
Reported outcome
```

A campaign reaching 100% funding does not mean execution is complete. Execution completion does not automatically prove impact. A reported outcome must retain its source and evidence status.

## 15. Evidence Journal

The approved Evidence Journal is a representative experience model for post-donation transparency.

Its visual grammar should prioritize:

```text
current fact
→ chronology
→ milestone
→ source
→ pending information
→ next expected update
```

The journal should feel factual first, human second, calm, chronological, and inspectable.

It should not resemble a social feed, gamified activity stream, or celebration timeline.

## 16. Photography

Direction: **Editorial documentary**.

Photography should be real, contextual, dignified, selective about privacy, capable of showing difficult reality, and emotionally honest without manipulation.

Avoid stock-like optimism, excessive smiles as required mood, pity-driven framing, exaggerated hardship, staged charity clichés, and synthetic documentary imagery presented as real evidence.

## 17. Campaign Imagery

Campaign imagery belongs to campaign reality.

The platform may frame and organize it, but must not visually fabricate beneficiary identity, distribution, successful outcome, verification, or participation.

Image treatment must not upgrade evidence.

## 18. Campaign Placeholder

Campaign placeholder must clearly feel like missing media, not synthetic campaign content.

Preferred direction:

- warm neutral surface;
- abstract editorial geometry;
- restrained Sun/Berry detail;
- simple media indicator;
- no fake person;
- no fake documentary scene.

Placeholder illustration must not be mistaken for beneficiary imagery.

## 19. Illustration

Illustration may support onboarding, process explanation, trust education, privacy-sensitive content, empty states, brand storytelling, and editorial key art.

Preferred direction:

- mature;
- human;
- editorial;
- recognizable;
- not mascot-like;
- not childish.

Illustration must never masquerade as real campaign evidence.

## 20. Public vs Product Expression

Kencleng uses one visual system with different expressive intensity.

### Public

May use more Newsreader, larger typography, asymmetric editorial composition, photography, illustration, Sun/Berry accents, and generous whitespace.

### Product

Should use more Instrument Sans, compact hierarchy, structured surfaces, restrained accents, evidence/status/provenance grammar, and predictable layout.

The shift is:

```text
expressive
→
disciplined
```

not:

```text
brand
→
generic software
```

## 21. Responsive Principles

Responsive design preserves, in order:

1. current task;
2. trust-critical information;
3. consequence;
4. chronology/provenance;
5. action accessibility;
6. content dignity.

Decorative composition may simplify. Critical information must not disappear.

Exact breakpoints and layout implementation remain engineering decisions.

## 22. Motion

Motion should orient, explain change, show transition, support chronology, reveal disclosure, and provide lightweight feedback.

Avoid default confetti, gamified completion, exaggerated springs, novelty motion, and constant ambient animation.

Exact duration/easing tokens remain OPEN until motion implementation is deliberately derived.

## 23. Accessibility

The visual system must not depend on color alone.

State communication should combine:

```text
label
+
structure
+
optional icon
+
color
+
context where necessary
```

Typography, contrast, controls, and responsive hierarchy must remain usable with realistic content.

Accessibility correctness belongs to actual implementation behavior, not visual documentation alone.

## 24. Anti-Patterns

Do not normalize:

- green as default Kencleng identity;
- yellow as universal success color;
- shield/checkmark-heavy trust UI;
- card soup;
- giant rounded rectangles everywhere;
- heavy shadows;
- glassmorphism;
- gradient-heavy brand expression;
- generic AI blobs;
- faux-documentary imagery;
- synthetic beneficiary imagery;
- sadness-selling;
- forced happy photography;
- progress bars for everything;
- impact claims derived from funding state;
- “verified” language without exact semantic authority.

## 25. Visual Reference Authority

Approved selected-direction visual references live under:

```text
docs/ui-ux/visual-references/selected-direction/
```

They establish character, hierarchy, relationship, system behavior, and expressive intensity.

They do **not** establish exact production geometry, component decomposition, route contracts, mock data, business semantics, or exact responsive implementation.

The approved Concrete Visual System exploration also validated the foundations and evidence/progress grammar represented by this document. Those exploration proofs are design evidence, not screenshot specifications.

## 26. Current Open Decisions

The following remain intentionally unresolved:

- final wordmark/logo;
- full illustration family specification;
- detailed photography governance;
- exact motion duration/easing tokens;
- final production terminology for provenance/truth labels;
- exact campaign placeholder asset implementation;
- whether Phosphor requires project-specific icon tuning;
- additional visual tokens that only become necessary during implementation.

Do not invent these silently. Material decisions require appropriate design review.

## 27. Downstream Engineering Invariants

Frontend implementation must preserve:

- Sunlit Editorial visual posture;
- Evidence-Led Optimism;
- Newsreader / Instrument Sans role separation;
- warm-neutral product foundation;
- restrained Sun and Berry usage;
- semantic color independence from brand color;
- border/spacing-first grouping;
- restrained radius and elevation;
- provenance visible through structure;
- funding / operational / reported-outcome separation;
- non-gamified progress;
- dignity of campaign imagery;
- illustration not masquerading as evidence;
- public expressive / product disciplined relationship.

## 28. Engineering Freedom

Engineering may determine CSS/Tailwind representation, token naming, variable organization, component APIs, responsive implementation, DOM structure, variant organization, local composition, performance strategy, and test strategy.

Implementation mechanics must not silently redefine visual-system meaning.

## 29. Final Test

Before establishing a new visual precedent, ask:

### Truth
Does this visual treatment preserve actual evidence and domain meaning?

### Hierarchy
Can users quickly understand what matters?

### Trust
Does confidence come from information structure rather than decoration?

### Brand
Does it feel like Sunlit Editorial without excessive yellow or ornamental styling?

### Product
Does it remain usable under realistic content, money values, states, and responsive constraints?

### Precedent
Would it be healthy if the next ten similar surfaces copied this behavior?

If not, it should not become system precedent.
