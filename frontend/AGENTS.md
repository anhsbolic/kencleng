# AGENTS.md — frontend/

This file adds Kencleng frontend-specific rules on top of root `AGENTS.md`.

Read root `AGENTS.md` first.

Scope:

```text
frontend/
```

Do not modify `backend/` from a frontend-scoped implementation session.

> Design reset note: the existing frontend implementation is not the source of truth for Kencleng's approved visual direction. Current design authority lives under `docs/ui-ux/`. Existing implementation details may be inspected as engineering evidence only until the frontend architecture/reset work deliberately establishes their future status.

---

## 1. Business Authority

Frontend is **not business authority**.

Do not recreate backend decisions for money, eligibility, permissions, verification, lifecycle transitions, financial validity, or security semantics.

Client validation and presentation derivation are allowed. Backend/domain truth remains authoritative.

If required product information is missing from API/spec, flag the contract gap instead of inventing client-side behavior.

---

## 2. State Ownership

For the current implementation, preserve the narrowest truthful state owner unless a later approved frontend architecture replaces this model:

```text
derivable
→ derive

API/server authoritative
→ server-data owner

navigation/share/back-forward state
→ URL when appropriate and non-sensitive

form lifecycle
→ form owner

ephemeral interaction
→ narrow local owner

genuinely shared client-owned state
→ approved shared state owner
```

Do not mirror server-authoritative data into unrelated client state or synchronize deterministic projections through effects.

---

## 3. API and Forms

API request/response semantics come from canonical API/domain authority.

Do not hand-write parallel business contracts already defined by OpenAPI/specs.

Client form validation improves UX; server validation remains authoritative.

Distinguish field validation from request/business failures.

---

## 4. Component Ownership

Use the narrowest truthful semantic owner.

Do not extract by line count or usage count alone. Repetition is evidence to evaluate shared semantics, not proof of abstraction.

Prefer explicit composition and stable variants over configuration-heavy generic components.

Read `components/README.md` before introducing or materially changing broad reusable components.

---

## 5. Shared Component Changes

Before materially changing broad shared primitives/components:

```text
classify change
→ discover current consumers
→ identify representative risk cases
→ implement
→ verify component
→ verify representative downstream consumers
```

Do not trust a stale manual consumer list. Compilation alone does not prove visual, behavioral, semantic, or accessibility compatibility.

---

## 6. Product Design Before UI Implementation

For material UI work, determine design readiness:

```text
READY
→ implement established intent

PARTIAL
→ extend established patterns using product/design judgment

OPEN
→ perform design exploration before canonical implementation
```

Do not silently invent material product or interaction intent while coding.

Read as needed:

```text
../docs/ui-ux/README.md
../docs/ui-ux/product-design-principles.md
../docs/ui-ux/brand-product-ui-brief.md
../docs/ui-ux/patterns.md
../docs/ui-ux/asset-governance.md
../docs/ui-ux/page-map.md
```

The approved direction is:

```text
Sunlit Editorial
+
Evidence-Led Optimism
```

Public surfaces may be more expressive/editorial. Authenticated operational surfaces must be more disciplined while remaining recognizably the same brand.

---

## 7. Visual Assets

Do not silently fill important visual gaps with arbitrary library icons, stock-like imagery, random gradients, generic AI decoration, or synthetic beneficiary-like photography.

Standard icons are appropriate for standard utility actions.

For expressive or brand-level assets, follow:

```text
../docs/ui-ux/asset-governance.md
```

Logo, wordmark, core illustration language, and other brand-defining assets require human approval before becoming canonical.

Illustration must never masquerade as documentary evidence.

---

## 8. Visual System Status

There is intentionally **no current canonical concrete production visual-system specification** defining exact:

- colors;
- typefaces;
- radii;
- spacing tokens;
- icon family;
- motion values.

Do not treat existing frontend CSS/tokens as design authority merely because they exist.

Do not revive the superseded assumptions that Kencleng must use green as primary brand color or the old Plus Jakarta Sans/Inter pairing.

Concrete visual-system decisions must be deliberately derived from `brand-product-ui-brief.md` and approved through the appropriate design/engineering process.

**Engineering follow-up — intentionally deferred.**

---

## 9. Selected Visual References

Approved direction-level visual references live under:

```text
../docs/ui-ux/visual-references/selected-direction/
```

Use them to understand:

- visual character;
- editorial composition intent;
- trust/information hierarchy;
- public vs product expressive intensity;
- Evidence Journal character.

Do **not** treat them as:

- product truth;
- route/API contracts;
- component architecture;
- exact CSS;
- pixel-perfect implementation targets.

The previous `docs/design-reference/` prototype system is superseded and must not be used as current visual authority.

---

## 10. Truthful UI

Frontend presentation must not collapse distinct product concepts into a generic trust signal.

Keep distinguishable where relevant:

- platform fact;
- organizer-provided information;
- organizer report;
- system/lifecycle state;
- pending/unavailable information;
- reported outcome.

Organization verification, campaign curation/publication state, and real-world outcome are not interchangeable.

Funding progress, operational progress, and reported outcome must remain distinct.

---

## 11. Rendered Verification

If a change materially affects rendered UI or spatial interaction, inspect the result in a real browser or equivalent layout-capable environment.

Automated unit/component tests do not prove layout, hierarchy, responsive behavior, clipping, overflow, or real-browser spatial interaction.

Use representative states, viewports, and realistic content rather than exhaustive screenshot matrices.

When Harscode provides a separate Testing phase, final rendered verification belongs there and must independently verify the observable result.

---

## 12. Testing

Use the repository's actual test/verification tooling. Tests should verify observable behavior rather than implementation structure.

User-controlled Markdown/HTML rendering requires hostile-content sanitization coverage.

Browser verification is separate from unit/component testing.

Do not claim checks that did not actually run.

---

## 13. Security Presentation

Never render manually converted user-controlled HTML unsafely.

Role-aware UI does not provide authorization security. Backend authorization remains authoritative.

Never expose secrets, raw tokens, or PII through logs or accidental UI/debug output.

Follow root `AGENTS.md` security rules.

---

## 14. Workflow Authority

Generic per-feature lifecycle lives in Harscode.

Default lifecycle:

```text
Exploration + Techplan
→ Build / patch loop
→ Code Review
→ Testing
→ Pull Request
```

Do not duplicate or redefine that lifecycle here.

Kencleng-specific sequencing, risk tiers, frontend/backend coordination, and domain delivery rules live in `../docs/kencleng-agentic-workflow.md` as a project orchestration overlay.

---

## 15. Source-of-Truth Routing

```text
business/domain semantics
→ ../docs/spec/<domain>/

API shape
→ ../api/openapi.yaml

frontend architecture
→ ../docs/project/kencleng-frontend-tech-stack.md

product design principles
→ ../docs/ui-ux/product-design-principles.md

brand + Product UI direction
→ ../docs/ui-ux/brand-product-ui-brief.md

UX behavior
→ ../docs/ui-ux/patterns.md

asset governance
→ ../docs/ui-ux/asset-governance.md

route/persona inventory
→ ../docs/ui-ux/page-map.md

selected visual evidence
→ ../docs/ui-ux/visual-references/selected-direction/

component contracts
→ components/README.md
```

If authorities genuinely conflict, surface the contradiction rather than choosing whichever source is easiest to implement.

---

## 16. Output Style

Default explanatory/process narration: terse.

Final deliverables remain complete. Do not compress risk notes, review findings, testing/build reports, PR descriptions, or sections whose Harscode workflow contract requires completeness.
