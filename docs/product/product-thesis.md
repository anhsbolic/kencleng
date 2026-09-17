# Kencleng — Product Thesis & Character

> Status: **Candidate Product Authority — not yet canonical**
> Created: 2026-09-17
> Purpose: Preserve the durable product character and trust model discovered through product/design exploration without duplicating visual-system or interaction implementation rules.
> Primary source evidence: `docs/ui-ux/brand-product-ui-brief.md`, `docs/ui-ux/product-design-principles.md`, `docs/ui-ux/page-map.md`, and `docs/ui-ux/patterns.md`.

## 1. Core product thesis

Kencleng is an **evidence-led fundraising and accountability product**.

Its intended relationship with users is:

```text
evidence
→ confidence
→ optimism
→ action
```

not:

```text
emotion
→ urgency
→ conversion
```

Kencleng may be warm, hopeful, and expressive, but it should not ask users to trust a campaign, organization, report, or platform claim merely because the experience feels polished or emotionally compelling.

The short product character is:

> **Evidence-Led Optimism**

Supporting phrase:

> **Hope, structured by evidence.**

These phrases originated in design exploration, but the underlying idea is broader than visual design. They describe how the product should earn trust and action.

## 2. What Kencleng optimizes for

Kencleng should help people make consequential decisions with enough context to understand:

- what is known;
- who is responsible for the information;
- what is reported versus independently established by the platform;
- what state something is currently in;
- what remains unknown, pending, changed, rejected, or incomplete;
- what money-related numbers actually mean;
- what action is available now;
- what will happen after that action;
- what evidence or accountability should exist later.

The product should not optimize donation conversion at the expense of comprehension, dignity, truthfulness, or long-term accountability.

## 3. Durable product character

### Thoughtful

Kencleng should order information deliberately and give users enough context for sensitive decisions instead of rushing them through a funnel.

### Candid

Unknown, delayed, changed, rejected, incomplete, and under-review states are valid product states. The system should expose them plainly rather than smoothing them into optimistic language.

### Composed

Money, identity, legality, rejection, accountability, and difficult campaign realities should be communicated calmly. Serious subject matter does not require visual or verbal drama.

### Human

Campaigns concern real people and real contexts. The product should preserve dignity, agency, privacy, and context instead of treating hardship as a conversion asset.

### Optimistic

Kencleng may emphasize possibility, participation, visible progress, and constructive next actions, but optimism must not imply an outcome the product cannot establish.

### Accessible without trivialization

Kencleng should feel approachable and understandable without becoming childish, gamified, or casually reductive around consequential decisions.

## 4. Product principles promoted from design learning

These are product-level principles. Their exact visual and interaction expression remains owned by `docs/ui-ux/`.

### 4.1 Confidence before conversion

A donation is a consequential decision, not merely a conversion event to maximize.

Before action becomes the dominant message, the user should be able to access the context needed to understand the campaign, responsible organization/steward, relevant lifecycle state, funding meaning, and available accountability context.

This does not require every fact to appear above every CTA. It requires the product not to deliberately hide trust-critical context behind conversion pressure.

### 4.2 Facts first, story with them

Campaign storytelling should behave as **facts that have a story**, not a story decorated with selective facts.

Narrative can make information human and understandable. It must not override provenance, state, money semantics, uncertainty, or accountability.

### 4.3 Trust comes from structure

Trust is not a badge, color, shield icon, or generic `verified` label.

Kencleng should earn confidence through combinations of:

- responsible-party identity;
- provenance/source clarity;
- lifecycle state;
- chronology;
- factual progress;
- review/curation context;
- visible uncertainty;
- evidence/reporting availability;
- predictable consequences and next steps.

A design treatment may reinforce trust, but it cannot create product truth.

### 4.4 Unknown is a valid state

If the platform does not know something, the correct product state may be `unknown`, `not yet reported`, `pending`, `under review`, or another truthful equivalent.

The product should not replace missing evidence with optimistic inference, decorative certainty, synthetic evidence, or vague trust language.

### 4.5 Progress is evidence, not gamification

Kencleng should preserve a conceptual distinction between:

```text
funding progress
≠ operational progress
≠ organizer-reported outcome
≠ independently established real-world impact
```

A completed funding target proves the relevant funding fact. It does not prove execution or impact.

Progress should orient users, not score generosity or create artificial pressure.

### 4.6 Dignity over pity

Hardship may be real and relevant campaign context. It must not become the primary persuasion mechanism.

Photography, writing, campaign evidence, and public presentation should preserve human dignity, privacy, agency, and context.

### 4.7 Accountability continues after fundraising

The conceptual Kencleng journey does not end when money is donated or a campaign closes.

Post-campaign result, disbursement, fund-use reporting, provenance, review state, and still-unknown outcomes are part of the product's trust loop when applicable.

### 4.8 Consequence before commitment

For consequential actions, users should understand the important result before committing.

This includes relevant state change, reversibility, loss of access or capability, review implications, monetary consequence, and expected follow-up where applicable.

The implementation may use confirmation, inline explanation, review steps, or other interaction patterns; Product Authority owns the requirement for meaningful comprehension, not a particular modal/dialog design.

## 5. Product truth classes

Kencleng should preserve distinctions between at least:

- **platform/system fact** — data/state the platform can directly establish;
- **organization-provided information** — campaign or organization content supplied by the responsible party;
- **organization report** — later reporting/accountability claims supplied by the organization;
- **curation/review state** — a specific product review decision, not generic proof of all claims;
- **transaction/lifecycle state** — system state such as processing, published, closed, pending, rejected, or equivalent;
- **pending/unavailable information** — information that is legitimately not yet known or present;
- **reported outcome** — an outcome claim attributed to its source;
- **independently established real-world impact** — a stronger claim that must never be inferred merely from funding, publication, or report submission.

Delivery specs and UI copy may choose concise language, but they must not collapse these classes in a way that changes meaning.

## 6. Product relationship to imagery and evidence

Campaign reality should be represented truthfully.

Real campaign/beneficiary/activity imagery may serve as evidence or context only when it is genuinely tied to the campaign and safe/appropriate to show.

Generated or illustrative assets may support:

- product education;
- abstract explanations;
- empty states;
- onboarding;
- privacy-sensitive concepts;
- brand storytelling.

They must not masquerade as:

- beneficiary identity;
- actual campaign activity;
- distribution evidence;
- fund-use evidence;
- real-world outcome proof.

Exact asset lifecycle, approval, prompt/generation behavior, and visual style remain owned by `docs/ui-ux/asset-governance.md` and related design authorities.

## 7. Public versus operational product character

Kencleng has two compatible expression modes.

### Public experience

May be more editorial, narrative, and expressive to help people understand purpose, context, evidence, and human meaning.

### Authenticated / operational experience

Should prioritize state legibility, chronology, responsibility, next action, consequence, and scanability.

Both must obey the same truthfulness, dignity, money, provenance, and accountability model.

This is a product distinction in communication posture, not a mandate for specific layouts or component families.

## 8. What remains design authority

This document intentionally does **not** promote the following into Product Authority:

- exact visual direction and composition rules;
- typography families/scales;
- color tokens;
- surface/radius/elevation rules;
- icon library or icon treatment;
- detailed responsive composition;
- component contracts;
- exact loading/skeleton treatments;
- interaction-pattern implementation;
- motion behavior;
- exact campaign-card structure;
- illustration/photography production specification.

Those remain owned by the canonical `docs/ui-ux/` tree.

## 9. Design-to-product feedback rule

Product development is not a one-way flow from business specification into design.

```text
Product hypothesis
        ↓
Design exploration / delivery / implementation
        ↓
new evidence about user meaning, trust, consequence, or product coherence
        ↓
Product Authority refinement when warranted
```

A design discovery should be promoted into Product Authority when it changes or clarifies durable product meaning rather than only presentation.

Likewise, implementation learning may refine Product Authority when it reveals a previously hidden product assumption.

This feedback loop must not be used to let visual precedent silently invent business rules.

## 10. Relationship to `product-overview.md`

`product-overview.md` owns the whole-product model: purpose, actors, concepts, capability map, end-to-end journey, trust/accountability relationships, and progressive-delivery posture.

This document owns the complementary **product thesis and character**: how Kencleng should earn trust, communicate uncertainty, treat people, frame action, and distinguish evidence from optimism.

Neither document owns frontend/backend implementation detail or shared API shape.
