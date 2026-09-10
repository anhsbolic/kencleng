# Kencleng — Product Design Principles

> Intended path: `docs/ui-ux/product-design-principles.md`
>
> Status: Draft
>
> Purpose: Product-design decision framework for Kencleng. This document defines the experience principles that guide humans and AI agents when an exact UI design does not yet exist.
>
> It does **not** define pixel values, component APIs, route inventory, or business rules.

## Context

Kencleng is a donation/crowdfunding application where users make decisions involving trust, money, identity, organizational legitimacy, campaign progress, disbursement, and fund-usage accountability.

That makes product design correctness broader than visual polish.

A Kencleng interface should help a user understand:

* what is happening;
* who or what they are trusting;
* what money or status a number represents;
* what action matters now;
* what will happen after an action;
* and where consequential details can be inspected.

The product personality is:

**warm + trustworthy + transparent + calm**

When those qualities conflict, use this priority:

```text
Trustworthiness
      ↓
Clarity
      ↓
Warmth
      ↓
Delight
```

"Warm" does not mean casual about money, privacy, identity, or irreversible actions.

"Trustworthy" does not mean cold, bureaucratic, or visually corporate.

---

# A. Core Product Design Principles

## 1. Trust before persuasion

Kencleng must earn an action before pushing for it.

For a donation surface, the user should be able to understand the relevant trust context before being pressured toward the primary CTA.

Examples of trust context may include:

* who organizes the campaign;
* whether the organization is verified;
* campaign purpose;
* progress toward the target;
* relevant campaign status;
* what happens after payment;
* reporting or accountability information when relevant.

Do not optimize donation conversion using dark-pattern urgency, manufactured scarcity, social pressure, or unsupported claims.

Bad:

```text
DONATE NOW
Only a little time left!
Popular campaign 🔥
95% trusted
```

when those concepts are not supported by real product/domain data.

Good:

```text
Organization identity
Verified status
Campaign purpose
Progress and amount raised
Clear primary donation action
```

The interface may persuade through clarity, credibility, storytelling, and demonstrated impact.

It must not manufacture trust.

---

## 2. Money must never be ambiguous

Every monetary value must make its meaning clear.

A large number without context is insufficient.

Bad:

```text
Rp2.500.000
```

Good:

```text
Rp2.500.000 terkumpul
dari target Rp10.000.000
```

Different money concepts must not become visually or linguistically interchangeable.

Examples include:

* donation amount;
* campaign target;
* amount collected;
* available funds;
* amount requested for disbursement;
* amount disbursed;
* amount reported as used;
* fees, deductions, or other adjustments if the domain introduces them.

The same principle applies to transaction and lifecycle states.

A user should understand both:

```text
what the status is
+
what that status means for them
```

A color badge alone is not sufficient explanation for consequential financial state.

---

## 3. Progressive disclosure, not information hiding

Complexity that exists in the domain must not be hidden merely to make the interface look simple.

Instead:

> Show what matters now, and keep consequential detail easy to reach.

A user should not need to process every administrative or historical detail before completing a simple task.

At the same time, details affecting trust, money, permissions, identity, or consequences must remain discoverable.

Good candidates for progressive disclosure include:

* long campaign narratives;
* historical activity;
* supporting documentation;
* secondary metadata;
* advanced administrative controls;
* explanation of uncommon states.

Do not use accordions, tabs, drawers, or "more" controls merely to reduce visible content.

Use them when they improve information hierarchy without concealing information the user needs to make the current decision.

---

## 4. One confident next action

Each meaningful surface should communicate the most likely next action clearly.

Visual hierarchy should answer:

> "What is the primary thing I can or should do here?"

This does not mean every screen literally has one button.

It means primary, secondary, tertiary, and destructive actions must not compete equally for attention.

For example:

```text
Campaign detail

Understand campaign
      ↓
Establish trust
      ↓
Understand progress
      ↓
Donate
```

should not visually become:

```text
Donate
Share
Bookmark
Report
Contact
Copy
Follow
More
```

with equal prominence.

When multiple actions are legitimately important, hierarchy should follow user intent and current state rather than arbitrary component styling.

---

## 5. Explain consequences before commitment

Consequential actions must explain what will happen before the user commits.

Generic confirmation such as:

```text
Are you sure?
```

is insufficient for meaningful state changes.

A good confirmation communicates the relevant consequence:

```text
what will change
what becomes unavailable afterward
whether the action can be reversed
what status the object will enter
what additional action may be required later
```

This is especially important for actions involving:

* publication;
* rejection;
* organization identity;
* permissions or representatives;
* disbursement;
* campaign closure;
* destructive operations;
* other irreversible or security-sensitive state changes.

The amount of explanation should be proportional to the consequence.

Do not create confirmation dialogs for trivial or easily reversible actions merely as ceremony.

---

## 6. Warm and human, but serious about trust

Kencleng should feel approachable and human.

The experience may use:

* warm language;
* generous whitespace;
* friendly shapes;
* thoughtful illustrations;
* human-centered empty states;
* culturally natural Indonesian copy;
* moments of restrained delight.

But the tone becomes calmer and more explicit around:

* money;
* payment;
* personal information;
* authentication;
* legal identity;
* verification;
* security;
* rejection;
* destructive actions.

Avoid both extremes:

```text
cold institutional fintech
```

and:

```text
overly playful charity app
```

The interface should feel optimistic without trivializing responsibility.

---

## 7. Do not imply what the product cannot prove

Visual design must not create domain claims that do not exist.

Examples of risky invented implications:

* "Trending"
* "Popular"
* "Recommended"
* "Urgent"
* "Top campaign"
* "High impact"
* "95% trustworthy"
* unofficial verification levels
* invented ranking or scoring

unless the product specification and underlying data explicitly support them.

A polished visual treatment does not make an invented concept legitimate.

Principle:

> **No visual implication without domain truth.**

The same rule applies to imagery.

A placeholder or generated illustration must not appear to be documentary evidence of a real campaign, beneficiary, organization, or impact unless it actually is.

---

## 8. Generic is acceptable for utility; intentional identity matters for expression

Not every part of Kencleng needs custom visual treatment.

Standard conventions are preferred for universal utility actions such as:

* close;
* search;
* edit;
* filter;
* calendar;
* navigation chevrons;
* password visibility.

However, product-defining or emotionally meaningful surfaces should not silently default to generic library icons or arbitrary decorative filler merely because those assets are convenient.

Examples include:

* landing-page hero visuals;
* important empty states;
* organization/campaign placeholders;
* campaign completion;
* donation success;
* trust/accountability storytelling;
* brand-defining imagery.

When an expressive visual is materially important and no suitable asset exists, treat that as a **design gap**, not automatically as permission to use generic filler.

Detailed asset-generation and approval rules belong in `brand-and-visual-assets.md`.

---

# B. Design Readiness

Frontend implementation should not assume that every feature arrives with complete UI/UX design.

Before implementation, classify the design readiness of the changed surface.

## READY

The user goal, hierarchy, states, and visual/interaction precedent are sufficiently defined.

Examples:

* another list page using an established list/browse pattern;
* a form composed from existing form behavior;
* a known detail-page layout;
* a new usage of an existing shared component.

Default action:

```text
Use existing product principles
+ existing UX patterns
+ existing visual/component system
→ implement
```

Do not redesign a mature pattern merely because another solution is possible.

---

## PARTIAL

The main experience is understood, but limited details are unresolved.

Examples:

* empty-state copy is undefined;
* a known page pattern needs one additional state;
* responsive rearrangement is not explicitly documented;
* a secondary interaction has no exact precedent.

Default action:

```text
Use the nearest established precedent
→ fill small gaps using these principles
→ surface material assumptions
→ implement when no product-level decision is being invented
```

The agent may exercise senior frontend/product-design judgment here.

Do not escalate every spacing or minor interaction decision to a human.

---

## OPEN

The feature requires meaningful product or interaction decisions that have no established precedent.

Examples:

* a new multi-step workflow;
* a materially different navigation model;
* a new dispute/review experience;
* a new information hierarchy for a consequential domain concept;
* a major landing-page story;
* a new brand-defining visual direction.

Default action:

```text
Design exploration
→ resolve product/design intent
→ human approval where required
→ engineering planning
→ implementation
```

Do not silently design an OPEN experience while writing production code.

---

# C. Agent Design Responsibility

An AI frontend agent working on Kencleng is not only a JSX implementation engine.

When appropriate, it should reason like a senior frontend engineer with strong product-design judgment.

That includes proactively noticing:

* weak information hierarchy;
* unnecessary cognitive load;
* missing states;
* unclear actions;
* responsive problems;
* generic visual filler where intentional visual identity would materially improve the experience;
* component patterns that drift from the established product language;
* ambiguous trust or money presentation;
* missing design decisions.

However, design judgment does not grant authority to invent product truth.

## Agent may decide autonomously

When consistent with existing principles and patterns:

* spacing and grouping;
* layout composition;
* responsive stacking/reflow;
* ordinary information hierarchy;
* standard loading/empty/error presentation;
* standard form affordances;
* standard accessibility behavior;
* reuse of existing visual/component patterns;
* minor copy refinements that do not change meaning;
* whether a utility action uses an established standard icon;
* small presentation improvements with no product-semantic consequence.

## Agent should propose, then seek human/product approval

For material new design decisions such as:

* a new interaction pattern;
* a materially different information architecture;
* a major multi-step experience;
* a new navigation model;
* meaningful changes to the primary-action hierarchy;
* a new expressive illustration direction;
* landing-page key art;
* major visual identity changes;
* a new shared pattern likely to influence future features.

The agent should not merely ask:

```text
"What do you want?"
```

It should first produce a reasoned recommendation with trade-offs.

## Agent must not invent

Without product/domain authority:

* business rules;
* permission semantics;
* role capabilities;
* privacy behavior;
* financial calculations;
* status meanings;
* verification claims;
* ranking/recommendation concepts;
* irreversible-action semantics;
* security guarantees;
* evidence of real-world impact.

When these are unclear, surface the gap.

Do not patch the uncertainty with UI assumptions.

---

# D. Design Exploration Contract

When a surface is classified OPEN, design exploration should establish enough intent for engineering to proceed.

The exploration does not need pixel-perfect mockups by default.

Minimum output:

## User goal

What is the user trying to accomplish?

## Context

What does the user already know when entering this surface?

What role/persona are they acting as?

## Primary action

What is the most important next action?

What conditions make it available?

## Information hierarchy

What must be understood before the primary action?

What is secondary?

What can be progressively disclosed?

## Interaction flow

What is the happy path?

Where are decisions or transitions?

## States

At minimum consider relevant:

* initial;
* loading;
* populated;
* empty;
* validation;
* submitting;
* success;
* error;
* permission-limited;
* unavailable/closed;
* stale/offline;

only when applicable.

## Responsive intent

What changes structurally on constrained layouts?

What must remain visible or reachable?

## Existing patterns

Which established Kencleng patterns and components can be reused?

## New pattern

Does this introduce a reusable UX pattern?

If yes, treat that as a design-system decision rather than local page improvisation.

## Visual asset needs

Does the experience materially require:

* illustration;
* placeholder;
* hero artwork;
* branded visual treatment;
* custom semantic graphic;
* other expressive asset?

If yes, classify the asset need before implementation.

Do not silently replace an unresolved expressive asset with generic filler.

## Open product decisions

List unresolved questions whose answers would change product meaning, permissions, trust, money, or irreversible behavior.

Those decisions require the appropriate human/product authority before implementation.

---

# E. Design Quality Heuristics

When evaluating a Kencleng experience, ask:

### Trust

* Can the user distinguish factual product information from decorative presentation?
* Are verification/trust signals based on real domain state?
* Is important accountability information reachable?

### Money

* Is every important amount labeled by meaning?
* Are financial states and consequences understandable?
* Could two different money concepts be mistaken for each other?

### Hierarchy

* Is the primary action obvious?
* Does secondary information support rather than compete with it?
* Is complexity progressively disclosed rather than hidden?

### Consequence

* Does a consequential action explain what will happen?
* Is reversibility clear when it matters?
* Are destructive and safe actions visually distinguishable?

### Humanity

* Does the experience feel approachable without trivializing the domain?
* Is the copy natural and calm?
* Are empty/success/error states useful rather than merely decorative?

### Integrity

* Does any icon, label, illustration, ranking, badge, or image imply something the domain cannot prove?
* Has a missing design decision been hidden behind a generic UI convention?

### Consistency

* Does the surface reuse established UX patterns and visual language?
* If it deviates, is the deviation intentional and justified?
* Would this new pattern be desirable if copied by the next ten features?

The last question is especially important for agent-generated interfaces.

---

# F. Relationship to Other UI/UX Documents

This document answers:

> **Why should the experience behave and feel this way?**

Use the other documents for more specific authority:

* `page-map.md` — which routes/personas/surfaces exist
* `patterns.md` — reusable page structures, common interaction states, and UX behavior
* `design-guidelines.md` — visual tokens, typography, color, shape, elevation, and component visual treatment
* `brand-and-visual-assets.md` — brand identity, iconography, illustrations, placeholders, expressive assets, generation/handoff, and approval authority
* `prototype-reference.md` — which prototype/reference has visual authority for a route or surface
* `design-reference-usage.md` — how prototype artifacts are translated into production implementation
* `../project/kencleng-frontend-tech-stack.md` — frontend architecture and technology decisions

If documents conflict:

1. product/domain specification owns business truth;
2. this document owns product-design principles;
3. `patterns.md` owns established UX behavior;
4. `design-guidelines.md` owns established visual-system decisions;
5. prototypes provide reference according to the authority described in `prototype-reference.md`.

A prototype must not override product truth merely because it looks complete.

---

# G. Evolution

These principles are intended to be stable, not frozen.

Update them when repeated product work demonstrates that a principle is:

* incomplete;
* ambiguous;
* consistently overridden;
* producing undesirable outcomes;
* or missing a recurring class of decision.

Do not add a new principle for a one-off preference.

Prefer updating reusable patterns or component guidance when the issue is local to a specific interaction or implementation surface.
