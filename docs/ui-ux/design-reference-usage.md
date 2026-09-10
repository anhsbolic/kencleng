# Kencleng — Using `design-reference/` for Frontend Development

> Intended path: `docs/ui-ux/design-reference-usage.md`
>
> Status: Draft v2
>
> Purpose: Define how frontend agents and developers inspect and translate frozen prototype output into production Kencleng UI.

## Core Principle

`design-reference/` is a **design reference**, not a production starter codebase.

Use it to understand:

```text
what the experience is trying to communicate
```

not:

```text
how production React must be structured
```

The correct flow is:

```text
inspect intent
→ cross-check current authority
→ translate into production architecture
→ render
→ compare
→ refine
```

Do not copy prototype implementation wholesale.

---

# A. What the Existing Files Are

Current `design-reference/` artifacts are Claude Design standalone HTML exports.

They contain a self-bootstrapping rendered prototype and embedded prototype source.

That source was created in a design/prototyping environment.

It is not expected to follow Kencleng's production:

* component ownership;
* state ownership;
* server-state architecture;
* forms architecture;
* API types;
* test architecture;
* Tailwind implementation conventions.

The root `AGENTS.md` treats this directory as frozen/read-only reference output.

Do not modify files in it during ordinary frontend implementation.

---

# B. Inspect the Rendered Design First

When feasible, inspect the reference as a rendered page before reasoning from its source.

The rendered output is usually the clearest evidence of:

* hierarchy;
* composition;
* density;
* visual emphasis;
* states;
* relationship between copy and layout;
* responsive intent if the reference supports multiple layouts.

Source inspection can help explain how the prototype was constructed.

It should not replace visual inspection.

This matters because:

> prototype source structure is not the same thing as design intent.

---

# C. Cross-Check Before Copying Any Decision

Before implementing from a reference, read the relevant:

```text
feature/domain spec
OpenAPI contract
product-design-principles.md
patterns.md
design-guidelines.md
brand-and-visual-assets.md
frontend/components/README.md
prototype-reference.md
```

You do not necessarily need every document in full for every trivial change.

Read the smallest relevant authority set.

At minimum, inspect `prototype-reference.md` for:

* route tier;
* freshness;
* known issues.

A visible prototype mistake does not become correct because it is easier to copy than the docs.

---

# D. What to Take from a Prototype

Prototype references are strong evidence for:

## Information hierarchy

What the user notices first, second, and later.

## Composition

How major page regions relate.

Examples:

```text
content + sidebar
header + summary + sections
form + supporting explanation
```

## State intent

Which user-visible states matter and how they differ conceptually.

## Interaction affordances

Examples:

* selectable method;
* expandable rejection reason;
* progressive disclosure;
* mobile-vs-desktop presentation.

## Content tone

Indonesian labels, helper text, headings, and microcopy may be reused or adapted when they remain semantically correct.

## Visual character

The reference can show:

* intended density;
* balance;
* emphasis;
* whitespace;
* relationship between content and decoration.

Treat these as design intent, not literal source-code contracts.

---

# E. What Must Be Re-Evaluated

## Component decomposition

Prototype component names/boundaries are **candidates**, not production architecture.

Do not automatically mirror:

```text
prototype component
→ production component
```

Instead ask:

```text
Does this represent a meaningful responsibility?
What is its semantic owner?
Does production already have this contract?
```

Follow:

```text
frontend/components/README.md
```

---

## State ownership

Prototype-local state exists to demonstrate behavior.

It does not establish production state ownership.

Production should determine whether data belongs to:

```text
server/API owner
URL/navigation
form lifecycle
local ephemeral state
shared client-owned state
derived value
```

according to project/frontend engineering guidance.

Do not translate every prototype `useState` into production `useState`.

Do not mirror server-authoritative data into client state unnecessarily.

---

## Data

Prototype data is illustrative.

Real data shape comes from:

```text
api/openapi.yaml
+
relevant domain/feature specification
```

Never infer a backend concept from a mock constant.

---

## Business behavior

A clickable prototype interaction may exist solely to demonstrate a screen.

It is not authority for:

* eligibility;
* permission;
* status transitions;
* financial calculations;
* verification;
* security behavior.

Cross-check domain truth.

---

# F. Translating Visual Styling

Kencleng production uses Tailwind CSS v4 with CSS-first tokens defined in:

```text
frontend/app/globals.css
```

and exposed through:

```css
@theme inline
```

Do not reproduce the prototype's inline token styles merely because they render correctly.

Example prototype:

```tsx
style={{
  background: "var(--color-primary-600)",
  borderRadius: "var(--radius-md)",
}}
```

Production should normally use existing semantic utilities/components:

```tsx
className="bg-primary-600 rounded-md"
```

when those tokens are canonical.

There is no production `tailwind.config.js` theme authority.

---

# G. Exact Values vs Visual Intent

Do not assume every pixel value in prototype output is canonical.

Translate:

```text
prototype spacing
→ nearest appropriate established spacing

prototype radius
→ canonical radius token

prototype typography
→ current typography scale

prototype color
→ current semantic token
```

unless an intentional one-off expressive composition genuinely requires a local value.

The goal is:

> preserve the design relationship using the production visual system.

Not:

> reproduce every numeric value.

---

# H. Current Visual-System Migration

Existing reference exports predate parts of the Visual System v2.

Production implementation should therefore apply current rules even where the reference differs.

Examples include:

### Secondary actions

Current v2 target:

```text
neutral / outlined secondary
```

not default amber-filled Secondary.

Do not copy an older filled-amber treatment as new precedent.

### Typography

Use current production type tokens rather than prototype-exported font sizes.

### Generic visual filler

When an existing reference contains a provisional/generic placeholder, check the current visual-asset system before reproducing it.

---

# I. Existing Production Components First

Before implementing a primitive or shared treatment visible in a prototype:

```text
inspect frontend/components/
→ inspect component registry
→ inspect contract when present
```

Do not recreate:

```text
Button
Badge
Input
ProgressBar
PasswordInput
```

locally merely because the prototype contains its own implementation.

If the existing production component cannot express a legitimate requirement:

```text
understand requirement
→ determine local vs systemic need
→ change at correct semantic owner
```

Do not bypass the design system with a local clone.

---

# J. Prototype Library Components

Any component library embedded in the design export exists for prototype rendering.

It is not the Kencleng production design system.

Do not import or replicate those components wholesale.

Translate their **visual/interaction intent** into the current production component layer.

A prototype `Button` may tell you:

```text
this action is dominant
```

The production `Button` contract determines how that dominance is rendered.

---

# K. Asset Translation

Treat every visual asset in a prototype according to its current status.

Ask:

```text
Is this real content?
Is this a placeholder?
Is this provisional?
Is this approved?
Is this canonical?
```

Do not assume:

```text
asset appears in Tier 1 prototype
→ canonical asset
```

If a current canonical asset exists, use it.

If an expressive asset is materially required but unresolved:

```text
flag design gap
→ prepare asset brief
→ generate if harness supports it
OR
→ hand off generation prompt
```

Do not permanently downgrade to a generic icon because the current coding harness cannot generate an asset.

---

# L. Logo and Brand Identity

A provisional wordmark or logo appearing in prototype output must not silently establish brand identity.

Check `brand-and-visual-assets.md`.

Brand-defining changes require human approval even when the current agent can generate candidates directly.

Implementation convenience does not establish canonical identity.

---

# M. Responsive Translation

Do not derive mobile UI by mechanically stacking prototype desktop regions.

Identify:

```text
primary task
trust/consequence information
essential content
secondary content
decoration
```

and preserve that priority on constrained layouts.

When the prototype includes a mobile reference, treat its structural intent strongly.

Still verify against:

* current UX pattern;
* current visual system;
* realistic content;
* accessibility;
* production navigation constraints.

If no responsive precedent exists, apply established responsive patterns rather than inventing arbitrary hiding/reordering.

Material new responsive interaction may move the surface from READY to PARTIAL/OPEN.

---

# N. Realistic Content Stress

Prototype samples are often cleaner than production content.

Before accepting a translation, consider realistic variation:

* long Indonesian labels;
* long organization names;
* long campaign titles;
* large Rupiah values;
* multi-line validation/rejection reasons;
* empty states;
* missing imagery;
* loading;
* error;
* permission differences.

Do not preserve a prototype composition that only works because its mock data is unusually short.

Preserve intent while improving robustness.

---

# O. Accessibility Translation

Accessibility markup present in the prototype may be useful evidence.

Still verify it against the actual production interaction.

Do not copy ARIA merely because it exists.

Ensure:

* semantics match behavior;
* keyboard operation works;
* focus treatment is visible;
* labels are meaningful;
* dynamic states are communicated appropriately.

Accessibility correctness belongs to the production interaction, not the appearance of prototype markup.

---

# P. Extracting Embedded Prototype Source

When source inspection is useful, extract the embedded JSX rather than reasoning from escaped standalone HTML noise.

A local extraction utility may be used to obtain readable scratch output.

The extracted files are:

* temporary inspection artifacts;
* regenerable;
* not production source;
* not intended to be committed.

Source extraction is useful for:

* discovering which visual states were modeled;
* locating prototype copy;
* understanding an interaction demonstration.

It does not raise the extracted code's authority.

---

# Q. Implementation Loop

For a Tier 1 surface, a healthy implementation loop is:

```text
read relevant authority
        ↓
inspect rendered reference
        ↓
identify design intent
        ↓
inspect existing production components
        ↓
implement using production architecture
        ↓
render with realistic states/content
        ↓
compare intent
        ↓
fix meaningful differences
```

Do not define success as:

```text
"looks vaguely similar"
```

or:

```text
"pixel-identical"
```

Evaluate:

* hierarchy;
* behavior;
* state coverage;
* responsive intent;
* design-system consistency;
* asset correctness;
* accessibility;
* robustness.

---

# R. When Deviation Is Correct

Deviation from a prototype is expected when required by:

* product/domain truth;
* OpenAPI reality;
* current UX pattern;
* updated visual system;
* canonical component contract;
* accessibility;
* responsive robustness;
* known prototype defect;
* canonical visual asset;
* realistic content behavior.

A significant deviation should be explainable.

If the reason is merely:

```text
"I preferred another design"
```

that is not sufficient for replacing an approved precedent.

---

# S. When the Prototype Reveals a Better New Direction

Sometimes the reference may expose an improvement that is broader than the route.

Do not silently bake a new system rule into production.

Instead classify it.

```text
route-local improvement
→ implement locally if appropriate

reusable UX improvement
→ propose/update patterns

visual-system improvement
→ propose/update design-guidelines

brand/asset direction
→ appropriate approval

shared component contract
→ component impact process
```

Good design exploration is welcome.

Untracked system drift is not.

---

# T. Build-Time and Final Verification

During Build, rendered inspection may be used as implementation feedback:

```text
implement
→ render
→ compare
→ correct
```

When the project workflow provides an independent Testing phase, final rendered verification belongs there.

Testing should verify actual production behavior rather than trusting:

* prototype output;
* build report claims;
* screenshots alone.

Screenshots are useful evidence.

They are not automatically proof of interaction correctness.

---

# U. Design Gaps

If the reference does not answer an important question, do not guess merely because implementation has started.

Classify the gap using:

```text
READY
PARTIAL
OPEN
```

from `product-design-principles.md`.

Examples:

```text
missing minor responsive arrangement
→ potentially PARTIAL

undefined new payment flow
→ OPEN

missing landing hero identity
→ visual/asset OPEN
```

The current harness may perform design exploration when capable.

If it cannot produce the required visual asset, preserve the intent through an asset brief and generation prompt.

Tool limitation must not silently become product-design limitation.

---

# V. Final Translation Test

Before considering a prototype-derived surface complete, ask:

### Product truth

Does it reflect actual domain/API semantics?

### UX

Does it preserve the intended user task and hierarchy?

### Visual

Does it use the current Kencleng design system?

### Components

Does it respect semantic component ownership?

### Assets

Are visuals canonical, approved, or explicitly provisional?

### Responsive

Does the experience survive realistic constrained layouts?

### Accessibility

Can the interaction be used and understood beyond visual appearance?

### Precedent

If future agents copy this production surface, would that be a good outcome?

If the answer to the last question is no, do not normalize the implementation as precedent.

---

# W. Related Documents

* `prototype-reference.md` — authority, tiers, freshness, known issues
* `product-design-principles.md` — design readiness and product-design authority
* `patterns.md` — reusable UX behavior
* `design-guidelines.md` — canonical visual system
* `brand-and-visual-assets.md` — asset generation, authority, and visual language
* `page-map.md` — route/persona mapping
* `frontend/components/README.md` — production component governance
* `../project/kencleng-frontend-tech-stack.md` — frontend technical architecture
* root `AGENTS.md` — repository boundary and source-of-truth rules
