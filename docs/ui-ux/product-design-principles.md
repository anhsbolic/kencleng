# Kencleng — Product Design Principles

> Status: Canonical
> Purpose: Stable decision framework for product-design judgment when exact UI has not yet been specified.
> Relationship: `brand-product-ui-brief.md` owns the selected brand/Product UI direction; this document owns reusable experience principles.

## Context

Kencleng handles trust, money, identity, organizational legitimacy, campaign progress, distribution, reporting, and public accountability.

Design correctness therefore means more than visual polish.

A Kencleng interface should help a user understand:

- what is happening;
- who or what they are trusting;
- what a number or status means;
- what is known versus reported or pending;
- what action matters now;
- what will happen after an action;
- where consequential detail can be inspected.

The selected brand direction is **Evidence-Led Optimism**. Product-design judgment must preserve that direction without inventing product truth.

## 1. Confidence Before Conversion

Kencleng must earn an action before pushing for it.

For donation surfaces, relevant trust context must be available before the primary CTA becomes the only meaningful visual message.

Trust context may include:

- who organizes the campaign;
- campaign purpose;
- relevant lifecycle state;
- funding progress;
- reporting/accountability information;
- what happens after payment;
- what is still unknown or pending.

Do not optimize conversion using manufactured urgency, guilt, scarcity, unsupported popularity, or other dark patterns.

## 2. No Visual Implication Without Product Truth

A polished treatment does not make an invented concept legitimate.

Do not visually or verbally invent:

- ranking;
- “trending” or “popular” status;
- urgency;
- recommendation;
- unofficial verification levels;
- trust scores;
- impact guarantees;
- unsupported beneficiary or outcome claims.

Product/domain authority owns product truth.

Illustration, photography, badge treatment, hierarchy, and color must not imply facts the platform cannot prove.

## 3. Money Must Never Be Ambiguous

Every consequential monetary value must make its meaning clear.

Distinguish concepts such as:

- donation amount;
- campaign target;
- amount collected;
- available funds;
- requested disbursement;
- disbursed amount;
- amount reported as used;
- fees or adjustments when the domain introduces them.

A large number without semantic context is insufficient.

Financial meaning must not depend only on typography or color.

## 4. Ordered Transparency, Not Information Dumping

Transparency does not mean showing every fact at once.

Show what matters for the current decision while keeping consequential detail easy to reach.

Progressive disclosure may be used for long narratives, supporting evidence, history, or secondary metadata when it improves hierarchy.

Do not hide trust-critical information, financial consequences, important uncertainty, required next actions, or material state changes.

## 5. Distinguish Truth Classes

The interface should make it possible to understand the difference between:

- platform fact;
- organizer-provided information;
- organizer report;
- system/lifecycle state;
- pending or unavailable information;
- reported outcome.

Do not collapse organization verification, campaign curation, publication state, or reported real-world outcomes into one broad “verified campaign” claim.

## 6. Unknown Is a Valid State

Information may legitimately be pending, unavailable, delayed, changed, under review, or not yet reported.

The interface should communicate these states calmly and respectfully instead of hiding them or replacing them with optimistic assumptions.

## 7. Progress Is Evidence, Not Gamification

Progress should orient the user, not score generosity.

Always distinguish:

- funding progress;
- operational progress;
- reported outcome.

A full funding target does not prove execution. Execution does not automatically prove impact.

## 8. One Confident Next Action

A meaningful surface should make its primary next action understandable.

This does not mean there can be only one button.

Primary, secondary, tertiary, and destructive actions should reflect user intent and consequence rather than arbitrary component styling.

## 9. Explain Consequences Before Commitment

Consequential actions must explain what will change before commitment.

Where relevant, communicate what will change, resulting status, reversibility, what becomes unavailable afterward, and what further action may be required.

Generic “Are you sure?” confirmation is insufficient for high-consequence operations.

Use friction proportional to risk, not ceremony for its own sake.

## 10. Dignity Over Pity

Hardship may be shown because it is real context.

It must not become a conversion device.

Preserve dignity, privacy, agency, and contextual truth.

The product should be optimistic without forcing smiling imagery or pretending difficult conditions are already resolved.

## 11. Warmth Without Trivialization

Kencleng should feel human and approachable.

Around money, identity, privacy, authentication, verification, curation, rejection, security, and destructive actions, the tone becomes calmer and more explicit.

Avoid both cold institutional fintech and childish charity UI.

## 12. Public Expression and Product Discipline

Public surfaces may carry richer editorial expression.

Authenticated/operational surfaces prioritize scanability, state legibility, and task clarity.

Both modes must remain recognizably one product and obey the same truthfulness rules.

## 13. Responsive Design Preserves Priority

Responsive transformation should preserve:

1. the primary task;
2. trust/consequence information;
3. essential content;
4. reachable actions.

Secondary composition may reflow or disclose progressively.

Do not mechanically stack desktop boxes or silently remove consequential information.

## 14. Design Readiness

High-fidelity UI is **not** a universal prerequisite for frontend implementation. Kencleng frontend development is driven by canonical design authority, not by mandatory design-file handoff.

Before material UI implementation, classify design readiness.

### READY
The user goal, information hierarchy, states, interaction intent, and relevant precedent are sufficiently defined. Implement directly from canonical product/design authority. A high-fidelity mockup is optional unless the task explicitly establishes one as required evidence.

### PARTIAL
The main experience is understood but limited non-semantic details remain unresolved. Resolve ordinary presentation details autonomously from established principles, patterns, visual-system rules, and existing production precedent. Surface only material assumptions.

### OPEN
The feature requires meaningful new product, interaction, information-architecture, or brand decisions without established precedent.

Do not silently design an OPEN experience while writing production code. First determine whether the gap is a material design ambiguity or missing product truth.

Default flow:

```text
READY / ordinary PARTIAL
→ implement directly
→ render
→ inspect
→ refine

material design ambiguity
→ propose low-fidelity alternatives
→ recommend one with trade-offs
→ obtain human decision when material
→ implement

missing product/domain truth
→ surface the authority gap
→ resolve the owning spec/contract
→ implement only after the truth is established
```

Low-fidelity exploration exists to resolve structure, hierarchy, and behavior. It may be textual, diagrammatic, wireframe-like, or otherwise intentionally low-cost. Do not create high-fidelity design artifacts merely as workflow ceremony.

## 15. Agent Decision Boundary

An agent may autonomously decide ordinary presentation details when they preserve established principles and patterns.

Do **not** request human approval for routine micro-decisions such as ordinary spacing choices, icon sizing, divider use, local alignment, or other presentation details already governed by the visual system and precedent.

Material new decisions should be proposed and reviewed, including:

- new interaction architecture;
- materially different information hierarchy;
- major navigation changes;
- consequential action placement or behavior;
- new brand-defining visual treatment;
- new shared UX patterns likely to influence future features.

When material design ambiguity exists, the agent should provide a small set of meaningfully distinct low-fidelity options, recommend a default, and explain relevant trade-offs. Do not multiply options when one established pattern clearly applies.

An agent must not invent business rules, permission semantics, privacy behavior, financial calculations, status meanings, verification claims, ranking/recommendation concepts, security guarantees, or evidence of impact. Those are product/domain authority gaps, not design ambiguity.

## 16. Evaluation Heuristics

### Trust
Can the user distinguish fact, report, status, and uncertainty?

### Money
Is every important amount semantically labeled?

### Hierarchy
Is the current goal and next meaningful action clear?

### Consequence
Does the user understand what an important action will do?

### Humanity
Does the experience preserve dignity and feel approachable without trivializing responsibility?

### Integrity
Does any copy, badge, icon, image, illustration, or progress treatment imply something the product cannot prove?

### Consistency
Would we want the next ten similar surfaces to inherit this behavior?

## 17. Relationship to Other Documents

- `brand-product-ui-brief.md` — selected brand + Product UI direction
- `patterns.md` — reusable interaction behavior
- `asset-governance.md` — asset truthfulness, lifecycle, approval, reuse
- `page-map.md` — persona/surface inventory
- `visual-references/selected-direction/` — approved direction-level visual evidence
- `docs/spec/` and OpenAPI — product/domain truth

Concrete production visual-system values are intentionally not defined here.
