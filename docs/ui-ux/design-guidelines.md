# Kencleng — Visual Design Guidelines

> Intended path: `docs/ui-ux/design-guidelines.md`
>
> Status: Draft v2
>
> Purpose: Define the visual system of Kencleng — color, typography, spacing, density, surfaces, shape, elevation, icon treatment, motion, and visual component hierarchy.
>
> This document answers:
>
> **"What should Kencleng look and feel like consistently?"**
>
> It does not define business rules, UX flow semantics, visual-asset generation policy, or React component architecture.

---

# A. Visual Direction

Kencleng should feel:

**warm + trustworthy + transparent + calm**

When visual qualities compete, prioritize:

```text
Trustworthiness
      ↓
Clarity
      ↓
Warmth
      ↓
Delight
```

The interface should feel human and approachable without resembling either:

```text
cold institutional fintech
```

or:

```text
playful charity template
```

The visual language should communicate **quiet confidence**.

That means:

* clear hierarchy rather than excessive decoration;
* warm accents rather than constant saturated color;
* generous but purposeful whitespace;
* readable financial information;
* restrained elevation;
* expressive illustration where it adds meaning;
* calm treatment around money, identity, verification, and consequential actions.

A surface should not need gradients, shadows, icons, cards, and badges simultaneously to feel designed.

Restraint is part of the visual identity.

---

# B. Relationship to Product Design

Visual treatment follows product meaning.

Before solving a visual problem, understand:

```text
product truth
→ UX hierarchy
→ visual hierarchy
→ component implementation
```

A polished treatment must never create product meaning that does not exist.

Examples:

* green styling does not create verification;
* a flame icon does not make something trending;
* a prominent card does not make something recommended;
* gold decoration does not make something premium or trusted;
* photography does not prove real-world impact.

See `product-design-principles.md` for product-design authority.

---

# C. Implementation and Token Authority

Kencleng uses **Tailwind CSS v4 with CSS-first configuration**.

Design tokens live in:

```text
frontend/app/globals.css
```

as CSS custom properties and are exposed to Tailwind through:

```css
@theme inline
```

There is no `tailwind.config.js` theme source of truth.

Conceptually:

```css
:root {
  --color-primary-600: #278a42;
  --radius-md: 12px;
}

@theme inline {
  --color-primary-600: var(--color-primary-600);
  --radius-md: var(--radius-md);
}
```

Day-to-day component implementation should use the corresponding Tailwind utilities rather than duplicating raw values.

Bad:

```tsx
<div className="bg-[#278A42] rounded-[12px]" />
```

Good:

```tsx
<div className="bg-primary-600 rounded-md" />
```

Raw values are acceptable when the value is:

* truly local;
* not part of the visual system;
* not likely to recur;
* or cannot reasonably be represented through the existing token system.

Repeated raw values are evidence that a token may be missing.

Do not introduce a token merely because one value appears once.

---

# D. Token Evolution

Tokens should represent stable visual decisions.

Use this order:

```text
existing token works
→ reuse

existing token almost works
→ reconsider composition before creating token

repeated legitimate need
→ propose token extension

one-off visual adjustment
→ keep local
```

Avoid creating tokens named after pages or features.

Bad:

```text
--campaign-card-green
--donation-page-shadow
--admin-header-radius
```

Prefer semantic visual concepts:

```text
primary
success
surface
border
radius
elevation
```

Domain meaning belongs above the token layer.

---

# E. Color System

## 1. Primary — Brand Green

Current brand family:

| Shade         | Value     | Typical role                        |
| ------------- | --------- | ----------------------------------- |
| `primary-50`  | `#F0FBF4` | subtle selected/active tint         |
| `primary-100` | `#DCF5E3` | low-emphasis interaction background |
| `primary-300` | `#8CDCA8` | subdued/disabled contexts           |
| `primary-500` | `#34A853` | brand accent/icon emphasis          |
| `primary-600` | `#278A42` | primary action                      |
| `primary-700` | `#1F6E35` | primary hover/active                |
| `primary-900` | `#164825` | high-contrast green text            |

Primary green communicates:

```text
Kencleng identity
+
positive primary action
```

It must not automatically mean:

```text
success
verified
completed
```

Those concepts use semantic success treatment.

---

## 2. Success

Current family:

| Shade         | Value     |
| ------------- | --------- |
| `success-50`  | `#ECFDF9` |
| `success-500` | `#0F9D6E` |
| `success-700` | `#0B7A56` |

Use for actual positive state:

* successful operation;
* completed state;
* verified state when domain truth supports it;
* achieved progress state.

Do not use Success merely because an action is desirable.

---

## 3. Warning

Current family:

| Shade         | Value     |
| ------------- | --------- |
| `warning-50`  | `#FFF4ED` |
| `warning-500` | `#E8590C` |
| `warning-700` | `#B8430A` |

Use when something requires attention but is not an error.

Examples may include:

* approaching deadline;
* overdue obligation;
* potentially consequential condition.

Warning must not be used as decorative warmth.

---

## 4. Error / Destructive

Current family:

| Shade       | Value     |
| ----------- | --------- |
| `error-50`  | `#FEF2F2` |
| `error-500` | `#DC2626` |
| `error-700` | `#B91C1C` |

Use for:

* failures;
* invalid state;
* rejection;
* destructive actions;
* critical corrective feedback.

Do not make an entire surface red when a localized error treatment is sufficient.

---

## 5. Information

Current family:

| Shade      | Value     |
| ---------- | --------- |
| `info-50`  | `#EFF6FF` |
| `info-500` | `#2563EB` |
| `info-700` | `#1D4ED8` |

Use for neutral explanatory information.

Information blue is intentionally not the primary brand color.

This keeps:

```text
brand/action
```

visually distinct from:

```text
neutral information
```

---

## 6. Accent — Warm Amber

Current family:

| Shade        | Value     |
| ------------ | --------- |
| `accent-50`  | `#FFFBEB` |
| `accent-400` | `#FBBF24` |
| `accent-500` | `#F59E0B` |
| `accent-600` | `#D97706` |

Accent amber provides warmth and expressive emphasis.

Good uses:

* restrained highlights;
* illustration details;
* decorative product accents;
* low-frequency labels;
* milestone treatment when not semantically Success/Warning.

Do **not** use Accent as:

* warning;
* error;
* verification;
* universal secondary CTA;
* arbitrary colored decoration.

Accent should feel special because it is used sparingly.

---

## 7. Neutral

Current family:

| Shade         | Value     | Typical role                |
| ------------- | --------- | --------------------------- |
| `neutral-50`  | `#F8FAFC` | page canvas                 |
| `neutral-100` | `#F1F5F9` | subtle surfaces/input fills |
| `neutral-200` | `#E2E8F0` | borders/dividers            |
| `neutral-300` | `#CBD5E1` | subdued borders             |
| `neutral-400` | `#94A3B8` | placeholders/disabled       |
| `neutral-500` | `#64748B` | secondary text              |
| `neutral-700` | `#334155` | normal body text            |
| `neutral-900` | `#0F172A` | high-emphasis text          |

Neutral colors carry most of the interface.

Brand color should not carry the entire UI.

---

# F. Color Discipline

Color communicates hierarchy and meaning.

Do not:

```text
assign a different color to every status
use color simply to make cards distinct
color every icon differently
use accent color because a section feels empty
```

Prefer:

```text
neutral interface
+
intentional semantic color
+
restrained brand warmth
```

A typical screen should visually be dominated by:

```text
neutral surfaces
→ typography
→ one clear primary action
→ semantic accents only where meaningful
```

This prevents the product from feeling noisy or gamified.

---

# G. Surface Hierarchy

Kencleng should not become a collection of nested cards.

Use three conceptual surface levels.

## Canvas

The page background.

Current default:

```text
neutral-50
```

## Base surface

Primary content surface, often:

```text
white
```

Examples:

* form region;
* main content panel;
* dialog;
* important summary.

## Subtle surface

Used for grouping secondary content:

```text
neutral-50 / neutral-100
```

depending on surrounding contrast.

---

## Card rule

A card is appropriate when the content is:

* an independently understandable item;
* selectable/clickable as one unit;
* meaningfully grouped from neighboring information;
* repeated as part of a collection;
* or needs a clear bounded surface.

Do **not** create a card merely because a section exists.

Bad:

```text
Page
 └─ Card
     ├─ Card
     │   └─ Card
     └─ Card
```

Good:

```text
Page
 ├─ natural section
 ├─ natural section
 └─ bounded card where grouping matters
```

Prefer spacing and typography before adding containers.

---

# H. Typography

## Font families

### Heading

**Plus Jakarta Sans**

Weights:

```text
600
700
800
```

Used for:

* display;
* page titles;
* meaningful section headings;
* card/item titles.

### Body

**Inter**

Weights:

```text
400
500
600
```

Used for:

* body copy;
* controls;
* labels;
* inputs;
* tables;
* metadata;
* buttons.

Do not use the heading font for every emphasized label.

Typography should create hierarchy without relying on decorative color.

---

# I. Type Scale

Current scale:

| Token     | Size / line-height | Typical use             |
| --------- | ------------------ | ----------------------- |
| `display` | `36 / 40px`        | hero/marketing emphasis |
| `h1`      | `30 / 36px`        | page title              |
| `h2`      | `24 / 32px`        | major section           |
| `h3`      | `20 / 28px`        | card/section title      |
| `h4`      | `18 / 24px`        | subsection              |
| `body-lg` | `18 / 28px`        | narrative/intro         |
| `body`    | `16 / 24px`        | standard UI/body        |
| `body-sm` | `14 / 20px`        | secondary/helper UI     |
| `caption` | `12 / 16px`        | metadata only           |

The scale is intentionally restrained.

Do not increase heading size merely to make a page feel more designed.

Large typography is appropriate when the page is intentionally expressive, especially marketing/landing surfaces.

Operational product surfaces should prioritize information density and scanability.

---

# J. Typography Hierarchy

Use weight before introducing unnecessary size jumps.

Recommended emphasis:

```text
Primary heading
→ size + weight

Secondary hierarchy
→ moderate size + weight

Metadata
→ smaller size + muted color

Important number
→ weight/size appropriate to context
```

Do not style every amount as a giant dashboard number.

Money prominence should follow its importance to the user's current decision.

For narrative text, avoid excessively wide lines.

Prefer a readable text measure rather than allowing paragraphs to span an entire dashboard-width container.

For aligned financial/tabular values, consider tabular numerals when it materially improves comparison.

---

# K. Spacing and Rhythm

Kencleng uses Tailwind's existing 4px-based spacing scale.

Do not introduce a parallel custom spacing system without demonstrated need.

Common rhythm should center around:

```text
4px   micro separation
8px   tightly related elements
12px  compact internal grouping
16px  standard component gap/padding
24px  section-level grouping
32px  major section separation
48px  page-level separation
64px+ expressive/marketing breathing room
```

These are guidelines, not mandatory pairings.

The principle is:

> **Distance communicates relationship.**

Related elements should be visually closer than unrelated sections.

Avoid mechanically applying the same `gap-6` everywhere.

---

# L. Density

Kencleng has different density needs.

## Public / donation-facing surfaces

Bias toward:

* easier scanning;
* larger breathing room;
* trust/context visibility;
* strong content hierarchy.

## Forms

Bias toward:

* predictable vertical rhythm;
* compact relationship between label, control, helper/error;
* enough space between unrelated field groups.

## Admin / curator operational surfaces

May be denser because comparison and repeated actions matter.

Density must not reduce:

* readability;
* target size;
* status comprehension;
* error visibility.

Do not force marketing-level whitespace into operational dashboards.

Do not force dense admin-table spacing onto public campaign pages.

---

# M. Layout and Content Width

Prefer a small number of meaningful layout widths rather than arbitrary `max-width` values per route.

Conceptual categories:

```text
reading
form
standard content
wide operational/data
full-bleed expressive
```

Exact production utilities/components may evolve as repeated usage demonstrates stable values.

Do not let a narrative paragraph inherit a very wide dashboard container.

Do not make forms unnecessarily wide simply because viewport space exists.

---

# N. Shape

Current radius scale:

| Token         | Value       | Typical role                      |
| ------------- | ----------- | --------------------------------- |
| `radius-sm`   | 8px         | compact controls/badges           |
| `radius-md`   | 12px        | buttons/inputs                    |
| `radius-lg`   | 16px        | ordinary cards/panels             |
| `radius-xl`   | 24px        | large expressive/overlay surfaces |
| `radius-full` | pill/circle | pills, avatars, circular controls |

The product uses clearly rounded forms, but not every object should be maximally rounded.

Hierarchy:

```text
small object
→ smaller radius

larger surface
→ larger radius when appropriate
```

Avoid the generic AI pattern:

```text
everything rounded-3xl
```

Large radius is an expressive choice, not a default quality signal.

---

# O. Elevation

Current shadows:

| Token       | Typical role         |
| ----------- | -------------------- |
| `shadow-sm` | subtle elevated item |
| `shadow-md` | popover/dropdown     |
| `shadow-lg` | modal/high overlay   |

Prefer borders and surface contrast for ordinary page structure.

Use elevation when there is a real visual stacking relationship.

Good:

```text
popover above page
modal above overlay
floating menu
```

Less appropriate:

```text
every dashboard card
every form field
every section
```

A flat interface with strong spacing can feel more trustworthy than one filled with floating surfaces.

---

# P. Buttons and Action Hierarchy

Action hierarchy follows UX intent defined in `patterns.md`.

## Primary

Visual treatment:

```text
primary-600
white text
primary-700 hover
```

Use for the dominant safe action.

Most surfaces should have at most one visually dominant primary action within the same decision context.

---

## Secondary

**Target v2 direction: neutral rather than amber-filled.**

Recommended treatment:

```text
white / transparent surface
neutral-700 text
neutral-200 border
neutral-100 hover
```

Use for legitimate secondary actions.

This keeps secondary actions visible without competing with Primary.

The existing amber-filled `secondary` component variant should be migrated deliberately through shared-component impact analysis rather than changed silently.

---

## Accent action

Do not make Accent a universal button tier.

When a rare expressive action genuinely needs warm emphasis, establish it intentionally rather than treating Amber as the automatic second-most-important button.

---

## Ghost

Use for low-emphasis actions where surrounding context already establishes affordance.

Do not use icon-only Ghost buttons when the action would be ambiguous without a label or accessible name.

---

## Destructive

Use Error semantic treatment.

Destructive actions should not be visually mistaken for the safe Primary action.

Use filled destructive treatment when prominence/friction warrants it; lower-emphasis destructive actions may use text/outline treatment where the component system supports it.

---

# Q. Button Size

Current heights:

```text
Small   36px
Medium  44px
Large   52px
```

Medium remains the standard interactive size.

Small is appropriate for compact operational contexts, but should not become the default for touch-heavy/mobile interactions.

Large is for singular high-emphasis actions, not as a way to make ordinary actions feel more important.

---

# R. Inputs

Default visual contract:

* 44px standard control height where appropriate;
* `radius-md`;
* neutral background/border;
* strong readable foreground;
* visible placeholder distinction;
* primary focus treatment;
* error semantic border/helper;
* disabled treatment clearly inactive.

A field group should visually read as:

```text
Label
Control
Helper / Error
```

Do not make placeholders substitute for labels.

Do not use saturated colored input backgrounds for ordinary state.

Error styling must remain understandable without color alone.

---

# S. Status and Badges

Badge tones should map many product states onto a small semantic visual vocabulary:

```text
neutral
success
warning
error
info
accent
```

Do not create a new color for each backend enum.

However:

> **Not every status should become a pill.**

Use a Badge when compact state recognition is useful.

For consequential states, use:

```text
status label
+
explanation / consequence
```

through appropriate supporting text or Banner treatment.

A pill alone should not carry critical financial or workflow meaning.

---

# T. Banners and Feedback Surfaces

Use banners for section/page-level feedback requiring persistent attention.

Semantic variants:

```text
success
error
warning
info
```

Keep banners visually calm.

Avoid:

* oversized icons;
* heavy saturated fills;
* stacked banners for every message;
* using banners where field-level feedback is more appropriate.

The visual weight should match the consequence.

---

# U. Progress Visualization

Progress is especially important on campaign-facing surfaces because it contributes to comprehension and trust.

The existing progress treatment uses:

```text
neutral track
+
primary fill
+
success treatment when completed
```

The visual bar must not carry meaning alone.

Pair it with explicit textual context such as:

```text
amount collected
target
percentage/progress when useful
```

Do not add animation that makes financial progress appear more dramatic or urgent than the underlying data.

---

# V. Iconography

Standard utility iconography may use the established library consistently.

Utility icons should:

* use a consistent stroke family;
* avoid unnecessary filled-vs-outline mixing;
* remain subordinate to text;
* use predictable size;
* have accessible names when needed.

Default sizes may center around:

```text
16px
20px
24px
```

depending on context.

Do not automatically place every icon inside:

```text
colored rounded square
```

That treatment should exist only where the visual system intentionally calls for it.

Product-semantic and expressive visual decisions follow `brand-and-visual-assets.md`.

Library availability must not determine brand identity.

---

# W. Illustration and Imagery Relationship

This document does not define individual assets.

The visual system expects expressive assets to harmonize with:

* warm restrained palette;
* rounded but not childish geometry;
* clear focal hierarchy;
* calm compositions;
* restrained detail inside operational product surfaces.

Hero/empty/placeholder/logo asset rules belong to `brand-and-visual-assets.md`.

Do not compensate for a missing expressive asset by adding arbitrary gradient or icon decoration.

---

# X. Backgrounds and Decorative Treatment

Default product screens should use simple, calm backgrounds.

Decorative backgrounds are most appropriate for:

* landing/marketing surfaces;
* onboarding;
* selected milestone moments;
* purposeful empty states.

Avoid default AI decoration:

```text
random gradient blob
blurred orb
glassmorphism panel
mesh gradient
floating icon cloud
```

unless it is intentionally part of the approved brand visual language.

Whitespace is a valid background.

---

# Y. Motion

Motion should explain change or reinforce continuity.

Good purposes:

* showing an element opening/closing;
* communicating state transition;
* preserving spatial context;
* lightweight feedback after interaction.

Motion should be:

```text
short
calm
predictable
interruptible where appropriate
```

Typical interaction transitions should usually remain within roughly:

```text
120–220ms
```

Larger panel/overlay transitions may be slightly longer when spatial continuity benefits.

Do not use:

* bounce as a default;
* gratuitous entrance animations;
* continuous decorative motion in operational screens;
* confetti for routine financial/security actions;
* animation that delays user progress.

Respect reduced-motion preferences.

---

# Z. Responsive Visual Behavior

Responsive design is not desktop UI compressed into a narrower width.

When space decreases, preserve this priority:

```text
primary task
→ trust/consequence information
→ essential content
→ supporting content
→ decoration
```

Decoration should usually yield before meaningful information.

Common transformations may include:

```text
multi-column → stacked
sidebar → inline/disclosure/drawer
action row → stacked or wrapped
dense metadata → reorganized hierarchy
```

Do not reduce typography to illegibly small sizes merely to preserve desktop composition.

Do not hide important content simply to avoid redesigning the layout.

Engineering robustness rules live in Harscode frontend best practices; this document owns the intended visual hierarchy.

---

# AA. Accessibility as Visual Quality

Accessibility is part of visual maturity, not a later compliance pass.

Visual requirements include:

* WCAG-appropriate contrast;
* clear focus indication;
* readable text hierarchy;
* no color-only semantic communication;
* visible invalid/disabled states;
* sufficiently large interactive targets;
* layouts that remain usable under text growth/zoom.

Muted text must remain readable.

Do not make important explanatory text low contrast simply because it is visually secondary.

Detailed interaction accessibility guidance belongs to frontend engineering best practices.

---

# AB. Long Content and Indonesian Copy

Production UI must survive real Indonesian-language content.

Do not design only against short English-like labels or mock strings.

Expect:

* long organization names;
* long campaign titles;
* long rejection/validation explanations;
* large formatted Rupiah values;
* multi-line button/label pressure in constrained layouts where unavoidable.

Prefer flexible composition over truncation.

Truncate only when full content remains available through an appropriate interaction and the product context permits it.

---

# AC. Anti-Generic AI UI Guardrails

AI-generated interfaces frequently converge on visually polished but interchangeable patterns.

Kencleng should actively resist unnecessary defaults such as:

### Cardification

```text
every section
→ rounded card + shadow
```

Use natural page hierarchy first.

### Icon-box repetition

```text
every heading
→ icon inside colored rounded square
```

Use when meaningfully established, not automatically.

### Dashboard-stat reflex

Do not create:

```text
three/four giant KPI cards
```

because the page is called a dashboard.

Metrics must serve a real user question.

### Gradient reflex

Do not use gradients simply to make a hero or CTA feel premium.

### Excessive pills

Not every metadata value is a badge.

### Giant-centered-heading reflex

Operational product pages usually benefit more from scannable hierarchy than marketing-size centered typography.

### Everything floats

Do not shadow every card, input, toolbar, and panel.

### Generic fintech imagery

Avoid random:

```text
coins
wallets
shields
handshakes
floating charts
```

unless they genuinely communicate Kencleng's specific story.

### Decorative novelty without meaning

If a visual decision could be removed without changing meaning, ask whether it is actually improving the experience.

---

# AD. Distinctiveness Test

For expressive surfaces, ask:

> **Could this exact interface belong to an unrelated SaaS product after changing only the logo and primary color?**

If yes, inspect whether a meaningful Kencleng-specific opportunity is being missed.

Do not respond by adding arbitrary decoration.

Distinctiveness should come from:

* product meaning;
* content hierarchy;
* brand assets;
* relevant visual metaphor;
* coherent tone.

Not novelty for its own sake.

---

# AE. Visual Readiness

Before implementation, visual readiness can be:

## READY

Existing guidelines/components/reference establish the visual direction.

Use them.

## PARTIAL

The visual system covers most needs but a local treatment is unresolved.

Agent may extend existing visual rules when the extension does not establish a new brand/system contract.

## OPEN

The surface requires a meaningful new visual language, expressive asset, component-system variation, or brand-defining decision.

Perform design exploration before treating an implementation as canonical.

This complements the Design Readiness model in `product-design-principles.md`.

---

# AF. Prototype Relationship

A prototype is evidence and visual precedent according to `prototype-reference.md`.

Translate:

```text
hierarchy
composition
states
visual intent
responsive intent
```

Do not automatically copy:

```text
component boundaries
raw spacing values
arbitrary prototype CSS
temporary asset choices
local mock-state architecture
```

Production design should converge toward the canonical Kencleng visual system.

When a prototype intentionally proposes a system-level visual evolution, evaluate that evolution explicitly rather than silently normalizing it into code.

---

# AG. Component Relationship

`components/ui/` is the production implementation surface of this visual system.

The authority chain is:

```text
product-design-principles
        ↓
design-guidelines
        ↓
component contract
        ↓
component implementation
```

A component should not silently establish a new system rule.

Examples that require design-system consideration:

* new global semantic color;
* new standard radius;
* new button hierarchy;
* new global elevation level;
* new recurring icon treatment.

Local implementation details do not require documentation updates unless they become stable system behavior.

Shared-component change discipline lives in:

```text
frontend/components/README.md
```

---

# AH. Visual Verification

When rendered UI changes, visual quality must be verified in the rendered product.

Relevant checks may include:

* hierarchy;
* spacing/rhythm;
* realistic content;
* responsive composition;
* visual state;
* focus visibility;
* semantic color use;
* component consistency;
* approved-reference alignment.

Build-time inspection supports iteration.

Independent Testing owns final rendered verification when the workflow provides a separate Testing phase.

Do not use pixel identity as the default definition of design correctness.

Evaluate **intent + system consistency + observable quality**.

---

# AI. Dark Mode

Dark mode remains outside current v1 scope unless product evidence changes that decision.

Do not:

* build an unused parallel theme;
* create dark variants "for completeness";
* double every design decision prematurely.

However, continue using semantic tokens and avoid unnecessary hardcoding so future theming remains feasible.

---

# AJ. Evolution

The design system is living but should not drift casually.

Update this document when:

* repeated product work reveals a missing visual rule;
* multiple components independently solve the same visual problem;
* a token repeatedly fails legitimate use cases;
* product personality evolves intentionally;
* accessibility or responsive verification exposes systemic weakness;
* a new approved brand/asset language changes visual-system expectations.

Do not update the design system for one isolated preference.

When a system-level visual change affects existing shared primitives:

```text
design decision
→ shared-component impact analysis
→ migration
→ representative consumer verification
→ documentation update
```

A prettier local screenshot is not sufficient evidence for a global design-system change.

---

# AK. Related Documents

* `product-design-principles.md` — product personality, hierarchy, trust, design authority
* `brand-and-visual-assets.md` — logo, icons, illustrations, placeholders, visual assets, generation/handoff
* `patterns.md` — reusable UX behavior
* `page-map.md` — route/persona inventory
* `prototype-reference.md` — prototype visual authority
* `design-reference-usage.md` — prototype-to-production translation
* `frontend/components/README.md` — component ownership, living registry, shared-component impact policy
* `../project/kencleng-frontend-tech-stack.md` — frontend technology and architecture
