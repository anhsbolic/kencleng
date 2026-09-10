# Kencleng Frontend Component System

> Intended path: `frontend/components/README.md`
>
> Status: Draft
>
> Purpose: Define component ownership, reuse boundaries, living component contracts, and change-impact discipline for the Kencleng frontend.
>
> This document governs production React components. It does not replace UX patterns, visual design guidelines, or domain specifications.

## Core Principle

Components are organized by **semantic ownership**, not by an assumption that everything should become reusable.

The default question is not:

> "Where can this component be reused?"

The default question is:

> **"Who owns this concept, and how broad is its semantic contract?"**

Use the narrowest layer that truthfully represents that ownership.

A component should move toward a broader layer only when broader semantics are demonstrated.

---

# A. Component Layers

Kencleng uses four component ownership levels.

```text
lowest blast radius
       │
       ▼
route-local
       │
feature/domain
       │
shared
       │
ui primitive
       ▼
highest potential blast radius
```

Broader reuse usually means a larger change blast radius.

Component size does not determine the layer.

A 20-line `Button` can have greater product-wide impact than a 300-line route-local form.

---

## 1. Route-local components

Location:

```text
app/<route>/...
```

Use when the component:

* exists only to compose one route or route segment;
* has semantics meaningful only in that local experience;
* is unlikely to be reused independently;
* does not represent a stable domain concept.

Examples:

```text
RegisterPageHeader
CampaignDetailSidebarLayout
AccountSettingsSection
```

A route-local component is not "less good" or temporary merely because it is local.

Do not move it into `components/features/` only to make the project appear more reusable.

Promote it when actual reuse or stable domain semantics emerge.

---

## 2. Feature/domain components

Location:

```text
components/features/<domain>/
```

Use when the component represents a meaningful concept within one product domain and may be composed by multiple routes in that domain.

Examples:

```text
features/account/
features/campaign/
features/donation/
```

A feature component may know:

* domain vocabulary;
* feature-specific presentation states;
* domain-shaped API data;
* domain-specific interaction semantics already defined by product specs.

It must not invent business rules that belong to the backend/domain specification.

Feature components should remain in their domain even when multiple pages use them.

Do not promote a feature component into `shared/` merely because usage count grows.

Promotion requires **cross-domain semantic equivalence**.

---

## 3. Shared semantic components

Location:

```text
components/shared/
```

Use for cross-domain components that are more semantic than a UI primitive but whose meaning is genuinely reusable across unrelated product domains.

Examples may include:

```text
PasswordInput
MaskedField
role/access presentation helpers
```

Shared components may understand a general product concept such as:

```text
credential input
sensitive-field reveal
role-aware presentation
```

but should not encode one feature's business semantics.

A shared component is an internal public contract.

Changing it requires downstream consumer analysis.

---

## 4. UI primitives

Location:

```text
components/ui/
```

UI primitives are generic visual or interaction building blocks with no Kencleng business/domain awareness.

Examples:

```text
Button
Input
Label
Badge
Banner
ProgressBar
Spinner
```

They may encode the established Kencleng visual system:

* tokens;
* variants;
* interaction states;
* accessibility behavior;
* responsive-safe behavior.

They must not know concepts such as:

```text
campaign
donation
organization verification
curation
disbursement
```

Those semantics belong in higher layers.

UI primitives have the broadest potential consumer set and therefore require the strongest change-impact discipline.

---

# B. Semantic-Owner-First Placement

Use this decision order before creating or moving a component.

```text
1. Is this composition meaningful only to one route?
   → route-local

2. Does it represent a stable concept inside one domain?
   → features/<domain>

3. Does the exact semantic concept recur across unrelated domains?
   → shared

4. Is it a generic visual/interaction primitive with no domain meaning?
   → ui
```

Do not use:

```text
"maybe reusable someday"
```

as evidence for broader placement.

Do not use repeated markup alone as proof of shared semantics.

Repetition is evidence worth evaluating.

It is not an automatic abstraction trigger.

---

# C. Pattern vs Component vs Domain Truth

These are separate authorities.

```text
Domain specification
→ what the product means

UX pattern
→ how a recurring experience behaves

Component contract
→ how production UI exposes that behavior

Component implementation
→ how the contract is realized in code
```

A UX pattern does not require one React component.

One React component does not automatically establish a UX pattern.

For example:

```text
Curation / Review
```

may be a reusable UX pattern implemented by several feature components.

Likewise:

```text
Button
```

is a reusable UI primitive but is not a product UX pattern.

Do not let component architecture redefine business or UX semantics.

---

# D. Living Component Registry

This README maintains a registry of the **broad-contract layers**:

* `components/ui/`
* `components/shared/`

It does not manually enumerate every route-local or feature component.

Those narrower components are discovered from the codebase when needed.

The registry exists because `ui/` and `shared/` components establish reusable contracts and have meaningful downstream blast radius.

## Registry fields

Each registered component records:

```text
Component
Layer
Purpose
Contract maturity
Dedicated contract
```

### Contract maturity

Use:

**Evolving**

The component is in production but its API/semantics may still change as product evidence develops.

**Stable**

Its intended semantics and main variants are established. Changes should preserve existing usage unless an intentional breaking change is approved.

**Foundation**

A highly reused primitive or semantic contract whose changes can affect large parts of the frontend. Treat changes as high-blast-radius even when implementation is small.

Maturity describes contract confidence, not code quality.

---

## Current registry

| Component        | Layer    | Purpose                                            | Contract maturity | Dedicated contract                                                   |
| ---------------- | -------- | -------------------------------------------------- | ----------------- | -------------------------------------------------------------------- |
| `Button`         | `ui`     | Generic actionable control                         | Foundation        | Not yet required                                                     |
| `Input`          | `ui`     | Generic text/input control                         | Foundation        | Not yet required                                                     |
| `Label`          | `ui`     | Accessible form labeling primitive                 | Foundation        | Not yet required                                                     |
| `Badge`          | `ui`     | Generic compact semantic/status treatment          | Evolving          | Not yet required                                                     |
| `Banner`         | `ui`     | Generic page/section feedback surface              | Evolving          | Not yet required                                                     |
| `ProgressBar`    | `ui`     | Generic progress visualization                     | Evolving          | Evaluate when financial/progress semantics expand                    |
| `Spinner`        | `ui`     | Localized indeterminate activity indicator         | Stable            | Not required                                                         |
| `PasswordInput`  | `shared` | Credential input with password visibility behavior | Evolving          | Evaluate as account flows mature                                     |
| `RequireRole`    | `shared` | Cross-surface role-aware presentation boundary     | Evolving          | Required if authorization/presentation behavior becomes more complex |
| `RequireOrgRole` | `shared` | Organization-role-aware presentation boundary      | Evolving          | Required if authorization/presentation behavior becomes more complex |

The registry is intentionally small.

When a new component is added to `ui/` or `shared/`, update this registry in the same change.

When a component moves out of those layers, update or remove its registry entry in the same change.

Do not maintain a manual list of every consumer.

Actual consumers are discovered from the repository at change time.

---

# E. What the Registry Does Not Replace

The registry is not a substitute for:

* TypeScript props/types;
* tests;
* stories/previews if introduced later;
* source-code search;
* UX pattern documentation;
* dedicated component contracts.

Code owns mechanically discoverable facts.

Documentation owns semantics that cannot safely be inferred from code alone.

Principle:

```text
CODE
→ what exists

DOCS
→ why the contract exists and what must remain true
```

Do not duplicate information into markdown when the compiler or repository can answer it more reliably.

---

# F. Dedicated Component Contracts

Do **not** create one markdown file for every component.

A dedicated component contract is warranted when at least one of these is true:

* the component has meaningful cross-domain semantics;
* its variants encode non-obvious product/design intent;
* it has accessibility behavior that callers must preserve;
* its responsive behavior is part of the contract;
* changes have repeatedly caused downstream regressions;
* it has a large or heterogeneous consumer surface;
* it has important forbidden usages;
* its visual/semantic contract cannot be understood safely from types/tests alone;
* it represents a stable pattern that future agents are likely to extend.

Suggested location:

```text
frontend/docs/components/
```

Example:

```text
frontend/docs/components/
├── README.md
├── masked-field.md
├── money-amount.md
└── curation-decision-panel.md
```

A dedicated contract may contain:

```text
Purpose
Semantic contract
When to use
When not to use
Variants
State behavior
Responsive behavior
Accessibility contract
Important invariants
Known limitations
Change-impact notes
Related UX pattern
```

Do not duplicate the entire props interface.

Reference source/types when needed.

---

# G. Component Contract Rule

A shared or UI component is not merely an implementation helper.

> **It is an internal public contract for its consumers.**

A change may affect:

```text
API
behavior
visual output
layout
accessibility
semantics
state behavior
responsive behavior
```

A change that compiles successfully is not automatically compatible.

For example:

```text
Button padding changes
```

may not break TypeScript but can break:

* toolbar density;
* narrow mobile layouts;
* dialog actions;
* inline forms.

Similarly:

```text
Badge default color changes
```

can change product meaning if consumers relied on established visual semantics.

---

# H. Shared Component Change Impact Policy

Before materially changing anything in:

```text
components/ui/
components/shared/
```

perform consumer impact analysis.

Minimum flow:

```text
understand requested change
        ↓
classify change
        ↓
discover actual consumers
        ↓
identify affected usage patterns
        ↓
change implementation
        ↓
verify component contract
        ↓
verify representative consumers
```

Do not evaluate the component only in isolation.

---

# I. Change Classification

Classify the proposed change before implementation.

## 1. Internal

Implementation-only change with no intended observable contract difference.

Examples:

```text
refactor helper
rename internal variable
simplify implementation
```

Expected discipline:

* component tests;
* normal regression verification.

Downstream visual inspection may not be necessary when rendered behavior provably remains unchanged.

---

## 2. Additive

Adds an optional capability without changing existing defaults or semantics.

Examples:

```text
new explicit variant
new optional slot
new supported state
```

Check:

* existing defaults stay unchanged;
* API remains understandable;
* new capability does not introduce prop/configuration explosion.

Representative new usage must be verified.

---

## 3. Visual

Changes observable styling or spatial behavior without intentionally changing semantics.

Examples:

```text
padding
typography
border
size
layout
responsive treatment
focus treatment
```

Requires:

```text
consumer discovery
+
representative rendered verification
```

The broader the component layer, the broader the representative sample.

---

## 4. Behavioral

Changes interaction or state behavior.

Examples:

```text
when a dialog closes
how password reveal persists
default button behavior
loading behavior
keyboard interaction
```

Requires:

* consumer discovery;
* behavior tests;
* representative consumer verification;
* review against related UX pattern.

---

## 5. Semantic

Changes what the component communicates or means.

Examples:

```text
changing Badge semantic mapping
changing what "verified" presentation means
changing role-gate presentation behavior
```

Requires:

* UX/product authority;
* consumer impact analysis;
* component contract update;
* downstream verification.

Do not infer semantic permission from component ownership.

---

## 6. Breaking

Existing consumers must change to preserve correctness.

Examples:

```text
remove prop
change default behavior
change required composition
change semantic meaning
```

A breaking change must be intentional.

The change should include migration of affected consumers rather than leaving repository-wide incompatibility for later cleanup.

---

# J. Consumer Discovery

Do not maintain manual consumer lists.

Before changing a broad component, search the current repository.

Possible methods include:

```text
import search
symbol/reference search
AST/type tooling
repository grep
IDE language-server references
```

Discover at minimum:

* direct consumers;
* wrappers around the component;
* feature-level components that indirectly expose it;
* tests exercising important variants.

For foundational primitives, sampling only direct imports may be insufficient.

Example:

```text
Button
→ used by FormActions
→ used by multiple feature forms
```

A visual change to `Button` may need verification at the feature level even if the feature never imports `Button` directly.

The goal is not to render every consumer.

The goal is to understand the blast radius and choose representative consumers capable of exposing regressions.

---

# K. Representative Consumer Verification

After changing a broad component, choose representative consumers according to the risk introduced.

Consider variation across:

```text
dense vs spacious layouts
narrow vs wide containers
form vs navigation usage
default vs destructive usage
loading/disabled states
long content
different semantic variants
keyboard/focus behavior
```

Example:

```text
Button visual change
```

may warrant checking:

```text
standard form CTA
narrow mobile form action
dialog action row
destructive action
disabled/loading state
```

Not every call site.

The smallest representative set capable of disproving compatibility is preferred.

Final independent rendered verification follows the project's Testing workflow.

---

# L. UI Primitive Change Discipline

Changes to `components/ui/` deserve extra caution because primitives are expected to be composition-safe.

Before adding a new primitive variant, ask:

```text
Is this a genuine reusable visual/interaction semantic?
```

or:

```text
Is one feature trying to push local styling into the global primitive?
```

Bad:

```text
Button variant="campaignCardTopRightUrgent"
```

Good:

```text
feature component owns campaign-specific composition
+
Button keeps stable generic action semantics
```

Do not solve feature-specific styling by continuously expanding primitive APIs.

When a primitive repeatedly cannot support legitimate cross-product needs, evolve the primitive intentionally.

---

# M. Shared Semantic Component Discipline

`components/shared/` must not become a dumping ground for anything used twice.

Before moving a component into `shared/`, establish:

```text
same concept
+
same behavioral contract
+
cross-domain use
```

Not merely:

```text
same markup
```

or:

```text
two imports
```

Examples:

A general sensitive-field reveal pattern may belong in `shared/`.

A campaign-specific identity block reused by three campaign pages still belongs in `features/campaign/`.

Cross-route does not automatically mean cross-domain.

---

# N. Feature Component Discipline

Feature components should encode domain presentation, not backend business authority.

Valid responsibility:

```text
render campaign progress returned by API
```

Invalid responsibility:

```text
recalculate campaign financial eligibility
and decide whether backend action is allowed
```

If feature UI needs information the API does not expose, surface the contract gap.

Do not recreate missing business logic in React.

Feature components may compose:

```text
ui primitives
shared semantic components
other components from the same domain
```

Cross-feature imports should be treated carefully.

If two domains need the same semantic concept, evaluate whether a `shared/` abstraction is justified rather than silently coupling domains.

---

# O. Route-Local Components Are Healthy

Do not treat route-local code as technical debt by default.

Local composition has benefits:

* smaller semantic scope;
* easier reasoning;
* lower blast radius;
* less premature abstraction.

Extract when responsibility becomes clearer.

Promote when actual reuse semantics emerge.

Avoid:

```text
route JSX
→ immediately extract everything
→ immediately move to features/
→ immediately generalize
```

Component boundaries follow meaning, not a component-count target.

---

# P. Component Creation Checklist

Before creating a new component:

* [ ] Is extraction creating a meaningful UI/behavioral responsibility rather than merely shortening a file?
* [ ] What is the narrowest truthful semantic owner?
* [ ] Does an existing component already own this contract?
* [ ] If similar components exist, is the similarity semantic or merely visual?
* [ ] Is this introducing a new UX pattern?
* [ ] Is this introducing a new visual-system variant?
* [ ] Are domain/business semantics already defined by product/API truth?
* [ ] Would placing this component at a broader layer create an unnecessary future precedent?
* [ ] If this becomes shared, would we want the next ten compatible surfaces to reuse this contract?

---

# Q. Shared/UI Component Change Checklist

Before changing a registered `ui/` or `shared/` component:

* [ ] Classify the change: internal / additive / visual / behavioral / semantic / breaking
* [ ] Read the relevant dedicated contract if one exists
* [ ] Discover current consumers from the repository
* [ ] Identify wrappers/indirect consumers where relevant
* [ ] Determine which usage variations are at risk
* [ ] Check whether the change modifies a UX pattern or design-system rule
* [ ] Avoid adding feature-specific configuration to a broad component
* [ ] Update tests for changed behavior
* [ ] Perform targeted rendered checks during implementation when UI output changes
* [ ] Verify representative downstream consumers
* [ ] Update registry/contract documentation when semantics or maturity change
* [ ] Let independent Testing perform final rendered verification according to workflow

---

# R. Adding a New Shared/UI Component

When introducing a component under:

```text
components/ui/
components/shared/
```

the same change must:

1. add the production component;
2. add appropriate tests;
3. add the registry entry in this README;
4. determine whether a dedicated contract is needed;
5. verify that the layer is semantically justified;
6. check whether the component introduces a visual/UX pattern that belongs in another source-of-truth document.

Do not merge a new broad component whose semantic owner is unclear.

---

# S. Moving a Component Between Layers

Promotion or demotion changes architectural meaning.

## Route-local → feature

Valid when the concept has become meaningfully reusable within a domain.

## Feature → shared

Valid when cross-domain semantic equivalence is demonstrated.

## Shared → ui

Rare.

Valid only when domain/product semantics have disappeared and the component is truly a generic primitive.

## Broader → narrower

Also healthy.

If a supposedly shared abstraction has accumulated feature flags or divergent semantics, splitting or moving it back toward feature ownership may reduce complexity.

Do not preserve a bad abstraction merely because it is already shared.

---

# T. Visual Asset Relationship

A visual asset is not automatically a component.

For example:

```text
campaign placeholder SVG
```

belongs to the visual asset system.

A reusable:

```text
CampaignImage
```

component may consume that asset while defining behavior such as:

* aspect ratio;
* loading treatment;
* missing-media behavior.

Shared asset changes follow `docs/ui-ux/brand-and-visual-assets.md`.

Shared component changes follow this document.

If a component and canonical asset both change, evaluate both blast radii.

---

# U. Design-System Relationship

`components/ui/` implements the production surface of the Kencleng visual system.

The authority chain is:

```text
product-design-principles.md
        ↓
design-guidelines.md
        ↓
ui component contract
        ↓
component implementation
```

Do not silently establish new visual-system semantics only in component code.

Examples:

```text
new status color meaning
new global radius system
new primary-action hierarchy
```

belong in the appropriate UI/UX source of truth before or alongside implementation.

Ordinary implementation refinements do not require a design-guideline update.

---

# V. Component Documentation Evolution

Documentation should grow according to actual complexity.

Do not front-load ceremony.

Use:

```text
simple component
→ source + types + tests + registry entry

complex shared contract
→ add dedicated component doc

repeated cross-component behavior
→ evaluate UX/design pattern

repeated systemic failure
→ improve governance/automation
```

A documentation file that is never consulted is not evidence of maturity.

The goal is:

> **minimum documentation that preserves semantic intent and makes high-blast-radius changes safe.**

---

# W. Automation Opportunities

Where reliable, automate mechanically detectable component-governance rules.

Potential future checks:

```text
ui/shared component exists but registry entry missing
illegal cross-feature imports
raw styling patterns banned by design-system rules
forbidden direct primitive usage
```

Do not automate a semantic decision merely because automation is desirable.

Questions such as:

```text
Does this component genuinely belong in shared/?
```

still require architectural judgment.

Add automation after repeated drift demonstrates value.

---

# X. Relationship to Other Documents

This document answers:

> **Where should a production component live, what contract does its layer imply, and how safely must it change?**

Use:

* `docs/ui-ux/product-design-principles.md` — product-design philosophy and design authority
* `docs/ui-ux/patterns.md` — reusable UX behavior
* `docs/ui-ux/design-guidelines.md` — visual-system rules
* `docs/ui-ux/brand-and-visual-assets.md` — visual assets and asset authority
* `docs/ui-ux/prototype-reference.md` — prototype authority
* `docs/ui-ux/design-reference-usage.md` — prototype-to-production translation
* `frontend/AGENTS.md` — concise routing/instructions for frontend agents
* domain/feature specs — business behavior and acceptance criteria

If a component implementation contradicts one of those authorities, the code does not become correct merely because it already exists.

---

# Y. Evolution

Update this component system when repeated implementation evidence shows that:

* a layer is ambiguous;
* components are consistently misplaced;
* shared changes frequently break consumers;
* registry maintenance is being missed;
* dedicated contracts are over- or under-used;
* consumer discovery is too costly;
* a mechanical rule should become automation.

Do not add ceremony for a single isolated mistake.
