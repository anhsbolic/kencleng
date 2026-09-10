# Kencleng — UX Pattern System

> Intended path: `docs/ui-ux/patterns.md`
>
> Status: Draft v2
>
> Purpose: Define reusable page structures and interaction behavior for Kencleng.
>
> This document answers:
>
> **"When users encounter a recurring kind of task or state, how should the experience behave?"**
>
> It does not define visual tokens, component APIs, or business rules.

## Context

Kencleng uses reusable UX patterns instead of maintaining pixel-level wireframes for every route.

This keeps behavioral intent stable while allowing the visual system and production implementation to evolve.

Related authority:

* `product-design-principles.md` — why the experience behaves and feels a certain way
* `page-map.md` — which route/persona uses which page pattern
* `design-guidelines.md` — visual language and tokens
* `brand-and-visual-assets.md` — icons, illustrations, placeholders, and expressive assets
* `prototype-reference.md` — visual reference authority
* `frontend/components/README.md` — production component ownership and contracts
* feature/domain specs — business rules and product truth

Patterns describe reusable experience behavior.

They do not authorize the frontend to invent domain capabilities.

---

# A. Pattern Reuse and Evolution

Before creating a new UX pattern:

```text
existing pattern fits
→ reuse

existing pattern mostly fits
→ extend intentionally

meaningfully different recurring interaction
→ propose new pattern

one-off page detail
→ keep local
```

Do not create a new named pattern merely because:

* markup differs;
* a page has additional fields;
* the visual composition changes;
* another component would be convenient.

A new pattern is justified when the **user interaction contract** is meaningfully different and likely to recur.

When extending a pattern, ask:

* Is the variation meaningful across more than one surface?
* Would future features benefit from inheriting this behavior?
* Does the extension preserve existing consumers?
* Is this still one interaction model or now a different pattern?

A useful test:

> **Would we want the next ten similar surfaces to inherit this behavior?**

If not, keep the behavior local.

---

# B. Page Patterns

These names are stable because `page-map.md` references them.

## 1. List / Browse

Purpose:

Help users scan, compare, search, filter, or select from a collection.

Typical structure:

```text
Page header
→ optional search/filter controls
→ collection
→ pagination/navigation
```

The collection may render as cards, rows, or another approved representation depending on content density and task.

### States

**Loading**

Use skeletons shaped like the expected collection rather than a page-level spinner.

**Empty**

Explain what is absent.

Show a primary CTA only when:

* an action meaningfully resolves the empty state; and
* the viewer is allowed to perform it.

A functional search-empty state usually needs less visual expression than a meaningful first-use empty state.

**Error**

Show a safe, user-facing explanation and retry when retry is meaningful.

Do not render raw backend errors.

**Success**

Render the collection and only expose navigation controls that have a meaningful effect.

### Search and filters

Search/filter controls should not dominate the page when browsing is the primary task.

When filter/search state should be shareable or participate meaningfully in browser navigation, prefer URL-backed state according to frontend state-ownership guidance.

Do not invent sort/filter semantics unsupported by the backend or product specification.

---

## 2. Detail

Purpose:

Help users understand one entity, its current state, relevant trust/context, and available actions.

Typical structure:

```text
Identity / title / status
→ primary information
→ trust or contextual information
→ supporting detail
→ actions appropriate to viewer + state
```

Public and dashboard detail pages may have different information priorities while still using the same underlying pattern.

### Long content

Long narrative content may use progressive disclosure when displaying everything at once harms scanability.

The collapsed view must not:

* hide information necessary for the current decision;
* imply that omitted content does not exist;
* trap keyboard or screen-reader users.

### Trust-sensitive public detail

When product/domain data contains an applicable organization verification state, public/donor surfaces should communicate it near the organization identity rather than burying it in secondary metadata.

Do not create additional trust scores or labels beyond the actual domain state.

### States

**Loading**

Use a skeleton matching the expected section structure.

**Not found**

Distinguish a genuine unavailable/not-found experience from transient request failure when doing so does not leak sensitive existence information.

**Error**

Show safe retry behavior appropriate to the failure.

**Success**

Render content and only actions authorized by the viewer's actual role/state.

---

## 3. Form

Purpose:

Help users provide or modify structured information with clear validation and submission consequences.

Typical structure:

```text
Context
→ grouped fields
→ field guidance
→ validation
→ primary submit action
→ secondary/cancel action
```

### Validation

Client validation improves UX.

Server validation remains authoritative.

Field-specific validation belongs with the relevant field.

Request-level failures belong at the form or section level.

Do not collapse:

```text
invalid field
```

and:

```text
request could not be completed
```

into the same error presentation.

### Submission

While a non-idempotent submit is in flight:

* prevent accidental duplicate submission;
* communicate progress;
* preserve entered data unless the successful flow explicitly transitions away.

Do not disable unrelated page behavior merely because one form control is submitting unless interaction correctness requires it.

### Success

Use the treatment that best explains what happens next.

Possible patterns include:

* inline completion state;
* navigation to the resulting resource;
* explicit next-step screen;
* lightweight toast when the surrounding context remains valid.

Do not use a toast as the only confirmation for a terminal or consequential flow when the user needs to understand the resulting state.

### Revisable Submission

Some domain flows follow:

```text
draft
→ submit
→ locked/reviewing
→ rejected
→ revise
→ resubmit
```

When a spec defines such a lifecycle:

* editing availability must follow domain state;
* rejection/revision must expose the approved reason/context needed to continue;
* locked/reviewing state must not look like an editable form merely with disabled controls.

This is a Form sub-pattern, not a separate page pattern.

---

## 4. Dashboard / Summary

Purpose:

Give an authenticated user a prioritized overview of information requiring awareness or action.

A dashboard is not a collection of every metric available.

Each summary item should answer at least one of:

```text
What is happening?
What needs attention?
What changed?
Where should I go next?
```

Typical structure:

```text
Dashboard shell
→ high-priority summary
→ actionable status
→ supporting overview
```

### Independent sections

Where data sources are independent, sections should be allowed to load or fail independently.

One secondary failure should not unnecessarily block the entire dashboard.

### Empty sections

Prefer section-level empty states when the dashboard itself still provides meaningful navigation/context.

### Metric discipline

Do not invent analytics, rankings, percentages, or KPIs merely because dashboards conventionally contain stat cards.

Every metric must correspond to real domain data and understandable user value.

---

## 5. Curation / Review

Purpose:

Help an authorized reviewer understand submitted material and make an accountable decision.

Typical structure:

```text
Submission context
→ material being reviewed
→ relevant history/context
→ decision action
→ decision reasoning when required
```

The item under review should remain visually distinct from the reviewer controls.

### Decision states

**Loading**

Show the review content structure, not only the controls.

**Already decided**

Present recorded outcome/history and remove misleading active decision affordances.

**Submitting**

Prevent accidental duplicate decisions and communicate progress.

**Failure**

Preserve reviewer context and allow safe retry when the domain permits it.

### Rejection/negative decisions

When domain rules require a reason:

* make the reason part of the decision flow;
* explain why it is needed;
* do not ask for the reason only after the destructive/negative action has already occurred.

The same interaction semantics may be reused across multiple curation domains.

That does **not** automatically require one highly parameterized production component. Component boundaries follow the component-governance rules.

---

## 6. Status / Tracking

Purpose:

Answer:

> **"What happened to the thing I submitted or initiated?"**

Typical structure:

```text
minimal context
→ current status
→ meaning / consequence
→ relevant next action
```

A status badge alone is insufficient when the state has an important consequence.

Examples:

```text
PENDING
→ still being processed

FAILED
→ what the user can do next

SUCCESS
→ what has completed and where to continue
```

### Sensitive token-based tracking

For unauthenticated lookup flows, error messaging must respect anti-enumeration/security requirements from the relevant domain spec.

Do not make visually different failure states reveal information the API intentionally conceals.

---

# C. Cross-Pattern Interaction Patterns

## 1. Primary Action Hierarchy

A surface should communicate one confident next action whenever one exists.

Actions are typically:

```text
primary
secondary
tertiary
destructive
```

Do not promote an action because:

* it looks visually interesting;
* the button component has a primary variant;
* another application commonly does so.

Hierarchy follows the user's current goal and the consequence of the action.

A destructive action should not visually compete with a safe primary action except when destruction itself is genuinely the user's current task.

---

## 2. Loading

Use loading treatment proportional to scope.

### Page/section content

Prefer skeletons that preserve expected layout.

### Inline action

Use localized progress such as a spinner/progress label within the action.

### Background refresh

Do not replace already-usable content with a full loading skeleton merely because background revalidation is occurring.

Loading UI should answer:

```text
What is waiting?
Can I still use anything?
```

Avoid unnecessary layout shifts between loading and loaded states.

---

## 3. Empty

First classify the empty state.

### Query/search empty

Meaning:

> Nothing matches the current criteria.

Usually provide:

* concise explanation;
* reset/change-filter action when useful.

### First-use empty

Meaning:

> The user has not created or received anything yet.

May provide:

* stronger guidance;
* primary next action;
* expressive asset when appropriate.

### Permission empty

Do not disguise lack of permission as ordinary emptiness when the user should understand why data/actions are unavailable.

### True no-content state

Do not invent CTA simply because empty-state components usually have buttons.

---

## 4. Error and Recovery

Errors should help the user recover without exposing implementation detail.

Distinguish:

```text
field validation
request failure
not found/unavailable
permission failure
stale/offline
terminal business state
```

when the distinction is both useful and safe.

A retry action is only appropriate when repeating the operation can reasonably succeed.

Never make "Try again" the universal response to every failure.

---

## 5. Success and Completion

Success treatment should match the significance of the completed action.

### Lightweight action

A toast/status update may be enough.

### Terminal flow

Use an explicit success state that explains:

* what completed;
* resulting status;
* what happens next;
* relevant destination/action.

Avoid excessive celebration around sensitive money/security operations.

Warmth and delight remain subordinate to clarity.

---

## 6. Stale / Offline Data

Kencleng's current PWA scope may allow cached content to remain visible while fresh data is unavailable.

When stale state is materially relevant:

* keep usable cached data visible;
* communicate that freshness is uncertain;
* distinguish stale content from a hard failure.

Freshness indicators are particularly important when old data could change user interpretation of:

* money;
* status;
* permissions;
* actionable deadlines.

Do not present stale values as known-current truth.

---

## 7. Status Communication

Status is a semantic system, not only a badge style.

Communicate status through:

```text
label
+
visual semantic
+
context/consequence when necessary
```

Do not rely on color alone.

Use exact domain states where meaningful rather than inventing friendly labels that change semantics.

Friendly explanatory copy may accompany the state.

If multiple states share the same visual semantic, their labels must still preserve the distinction.

---

## 8. Money Presentation

Every important amount needs an explicit semantic label.

Avoid surfaces where multiple currency amounts appear without clear distinction.

When presenting progress, make the relationship understandable:

```text
amount collected
relative to
campaign target
```

When presenting financial actions, distinguish clearly between:

* amount being entered;
* amount already collected;
* amount available;
* amount requested;
* amount completed/disbursed/reported.

Use Indonesian currency formatting consistently according to the project's formatting utility/component contract.

Do not encode financial meaning through typography/color alone.

---

## 9. Confirmation and Consequential Actions

Do not use generic confirmation dialogs reflexively.

Ask first:

```text
Is the action consequential?
Is it difficult to reverse?
Could the user reasonably trigger it accidentally?
Does the consequence need explanation?
```

For consequential actions, confirmation should state:

* the action;
* the consequence;
* whether it is reversible when relevant;
* any immediate state transition the user needs to understand.

Bad:

```text
Are you sure?
```

Better:

```text
Publish this campaign?

Once published, it becomes visible to donors.
[Cancel] [Publish]
```

Exact consequences must come from domain truth, not invented copy.

For low-risk reversible actions, avoid confirmation ceremony.

---

## 10. Destructive Actions

Destructive actions require appropriate friction, not maximum friction.

Use visual semantics that distinguish destructive actions from ordinary primary actions.

For high-impact deletion/removal/revocation:

* name the affected object;
* explain the meaningful consequence;
* require confirmation when accidental activation would be costly.

Do not place destructive controls immediately beside high-frequency safe actions without sufficient visual separation.

---

## 11. Progressive Disclosure

Use progressive disclosure when content is useful but not necessary for the current decision.

Common mechanisms:

* expandable section;
* disclosure panel;
* drawer;
* secondary detail page;
* tabs when content represents stable peer categories.

Choose the mechanism according to information relationship, not visual novelty.

Do not bury:

* trust-critical information;
* financial consequences;
* validation errors;
* required next actions.

---

## 12. Search, Filter, Sort, and Pagination

Use search when users can reasonably identify items by terms.

Use filters for meaningful domain dimensions.

Use sort only when the ordering has a real defined meaning.

Do not expose controls with only one meaningful option.

Pagination remains the default for v1 collection navigation unless a feature explicitly establishes another approach.

If filter/sort/search state should support sharing or browser navigation, consider URL ownership rather than local-only state.

The frontend must not create a filter/sort option unsupported by product/API truth.

---

## 13. Responsive Transformation

Responsive design preserves task and information priority; it does not merely stack desktop boxes.

When space becomes constrained:

1. preserve the primary task;
2. preserve trust/consequence information;
3. allow secondary content to reflow or disclose progressively;
4. keep actions reachable;
5. avoid silently deleting content.

A desktop sidebar may become:

* inline content;
* drawer;
* disclosure section;

depending on its role.

Responsive behavior should reuse established transformations when available rather than inventing a new mobile interaction per page.

Exact robustness requirements live in frontend engineering guidance and visual verification.

---

# D. Pattern vs Component

A UX pattern and a React component are not the same thing.

Example:

```text
Curation / Review UX pattern
```

may be implemented through several components.

Likewise:

```text
Button
```

is a reusable UI component but not a UX pattern by itself.

This document owns:

* interaction semantics;
* state behavior;
* information relationships;
* reusable experience rules.

`frontend/components/README.md` owns:

* component taxonomy;
* semantic ownership;
* stability;
* reuse boundary;
* change-impact policy.

Dedicated component contracts own component-specific behavior when necessary.

Do not place detailed component APIs in this file.

---

# E. Existing Named Component Behaviors

Earlier versions of this document contained specifications for:

* `MaskedField`;
* `SecureUploadNote`;
* `CurationDecisionPanel`.

These should migrate to the living component documentation because they describe concrete production components rather than generic UX patterns.

The UX principles they embody remain valid:

### Sensitive-field reveal

Sensitive information should remain masked by default when required by domain/security policy.

Reveal behavior must follow the actual audit/security specification.

### Secure upload reassurance

Sensitive/legal upload surfaces may provide contextual reassurance when it improves user confidence and accurately reflects product behavior.

### Curation decision

Review flows should share consistent decision semantics across domains where the business rules are genuinely equivalent.

The component architecture implementing those semantics is not prescribed here.

---

# F. Introducing a New Pattern

When an existing pattern does not fit, design exploration should establish:

```text
Problem
User goal
Existing pattern considered
Why it does not fit
Proposed interaction
States
Responsive behavior
Accessibility implications
Consequential/trust implications
Expected reuse
```

Classify the proposal:

### Local variation

Only one feature needs it.

Keep it local.

### Pattern extension

Existing pattern still represents the same user interaction but needs a reusable additional rule.

Update the existing pattern.

### New pattern

A genuinely distinct interaction recurs or is expected to recur.

Add a named pattern.

New patterns with material product/interaction consequences require appropriate human/product review before becoming precedent.

Do not add patterns simply to document an implementation after the fact.

---

# G. Pattern Change Impact

Changing an established UX pattern can affect multiple routes even when no shared React component exists.

Before materially changing a pattern:

```text
identify routes that reference it
→ determine which behavior changes
→ inspect affected personas/states
→ determine component impact
→ verify representative consumers
→ update page-map or component contracts when necessary
```

Classify changes:

### Clarification

Makes existing intent less ambiguous.

Usually low risk.

### Additive

Adds a new optional state/variation without changing existing behavior.

Check relevant adopters.

### Behavioral

Changes how existing consumers behave.

Requires downstream impact analysis.

### Semantic

Changes what an action/status/information structure means.

Requires product/domain authority and broader review.

Do not assume a markdown-only pattern update has no product blast radius.

---

# H. Relationship to Design Readiness

`product-design-principles.md` classifies design readiness as:

```text
READY
PARTIAL
OPEN
```

Patterns help resolve that classification.

### READY

An existing pattern covers the interaction with no meaningful unresolved UX decision.

### PARTIAL

A known pattern applies but requires a small extension or local variation.

### OPEN

No existing pattern adequately represents the required interaction, or the feature introduces a material new UX/product decision.

OPEN surfaces require design exploration before engineering implementation.

---

# I. Relationship to Visual Assets

Patterns may identify when an asset has a functional role.

Examples:

```text
first-use empty state
→ expressive illustration may be valuable

search-empty state
→ usually lightweight treatment

landing storytelling
→ expressive asset likely material
```

Patterns do not prescribe the illustration style or generate the asset.

That authority belongs to `brand-and-visual-assets.md`.

A missing expressive asset should remain an explicit design gap rather than silently becoming a generic icon fallback.

---

# J. Evolution

Patterns should evolve from repeated product evidence.

Update this document when:

* the same interaction ambiguity appears repeatedly;
* multiple pages are independently solving the same UX problem;
* an existing pattern consistently fails a real use case;
* a new feature establishes a reusable behavior worth preserving.

Do not expand this document for every local design choice.

Prefer:

```text
stable product principle
→ product-design-principles.md

reusable interaction behavior
→ patterns.md

visual language
→ design-guidelines.md

asset/illustration behavior
→ brand-and-visual-assets.md

component contract
→ frontend component documentation

route-specific behavior
→ page-map / feature specification
```
