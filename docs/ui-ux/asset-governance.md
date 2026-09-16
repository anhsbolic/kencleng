# Kencleng — Asset Governance

> Status: Canonical
> Purpose: Govern truthfulness, approval, lifecycle, reuse, and change impact for visual assets.
> Relationship: `brand-product-ui-brief.md` owns the approved visual direction. This document owns how assets enter and evolve within that direction.

## 1. Asset Classes

Classify an asset before creating or integrating it.

### Level 1 — Utility
Universal interface actions such as close, search, menu, edit, filter, calendar, chevron, password visibility, external link, and download.

Default: use familiar established iconography. Custom art is usually unnecessary.

### Level 2 — Product-semantic
Visual treatment for domain concepts such as organization verification, campaign state, donation, fund usage, disbursement, reporting, or representative role.

Use a consistent treatment where repeated meaning exists. A standard icon may be sufficient when it communicates the concept truthfully.

### Level 3 — Expressive product asset
Visuals that add meaning, humanity, onboarding, explanation, or product character, such as empty-state illustration, campaign placeholder, explanatory graphic, landing section art, or trust/accountability illustration.

Do not silently downgrade a meaningful expressive need into arbitrary generic filler.

### Level 4 — Brand-defining asset
Logo, wordmark, core illustration language, major key art, signature motif, or other identity-defining system.

Human approval is required before a new Level 4 asset becomes canonical.

## 2. Truth Before Expression

A visual must not imply evidence the product does not possess.

Generated, illustrated, placeholder, or decorative media must never be mistaken for:

- a real beneficiary;
- a real organization;
- a real campaign location;
- actual goods purchased;
- proof that distribution occurred;
- proof of real-world impact.

Principle:

> Illustration may communicate a concept. It must not fabricate evidence.

Synthetic documentary-style beneficiary imagery should not be used as if it were campaign evidence.

## 3. Photography

Approved direction: **editorial documentary**.

Campaign photography should be:

- real and authorized for its use;
- contextual;
- dignified;
- capable of showing difficult reality without exploiting it;
- privacy-aware;
- free from forced optimism.

Avoid pity-driven crops, staged charity clichés, stock-photo optimism, or edits that materially misrepresent conditions.

Photography can create human connection; it does not prove outcome by itself.

## 4. Illustration

Approved direction: **mature human editorial illustration with supporting abstract explanatory graphics**.

Use illustration for:

- process explanation;
- trust education;
- onboarding;
- privacy-sensitive contexts;
- meaningful empty states;
- brand storytelling;
- selected key art.

Illustration should feel human and distinctive without becoming childish, generic corporate illustration, or AI-slop.

Avoid generic gradient blobs, floating 3D objects, gratuitous glassmorphism, interchangeable polished-human scenes, and decorative assets with no product meaning.

## 5. Campaign and Organization Placeholders

Missing media must still feel intentional.

A placeholder must:

- be clearly recognizable as a placeholder;
- not impersonate campaign evidence;
- preserve layout/aspect-ratio stability;
- remain subordinate to real content;
- fit the approved visual direction.

The final placeholder system remains an OPEN asset decision and requires review before becoming canonical.

## 6. Iconography

Standard utility actions should remain familiar.

Product-semantic icon treatment should be consistent and restrained.

Avoid overusing trust clichés such as shields, generic checkmarks, locks, or badges as substitutes for actual information structure.

The final icon family remains OPEN.

## 7. Wordmark / Logo

No incidental implementation wordmark or placeholder logo may silently become canonical.

The identity should not depend on a literal kencleng/celengan metaphor and should remain viable if the product name changes.

Possible asset states:

```text
EXPLORATION
PROPOSED
APPROVED
CANONICAL
SUPERSEDED
```

A final wordmark/logo remains OPEN and requires human approval.

## 8. Asset Approval Model

### EXPLORATION
Working evidence. Not for production authority.

### PROPOSED
A candidate intended for review.

### APPROVED
Reviewed for a stated use. Approval can be surface-specific.

### CANONICAL
Part of the established reusable visual system. Future work should reuse or intentionally extend it before creating a competing direction.

### SUPERSEDED
No longer current. Historical value belongs in Git history or clearly historical records, not as competing active-tree authority.

A file existing in the repository is not automatically canonical.

## 9. Selected Direction References

The approved Sunlit Editorial public composition and Evidence Journal composition are **Selected Direction References**.

They are approved evidence for:

- brand character;
- expressive intensity;
- composition principles;
- trust/information hierarchy;
- public versus product translation.

They are not production assets, pixel specifications, component contracts, or evidence of domain behavior.

## 10. No Generic Filler

Do not add visual decoration merely to make a surface feel complete.

Avoid defaulting to:

- arbitrary gradients;
- generic fintech coins;
- generic handshake imagery;
- fake dashboards;
- decorative icon clouds;
- stock-like charity photography;
- AI-generated beneficiary-like photography;
- empty decorative blobs.

Whitespace and typography are valid when no meaningful asset is required.

When an expressive asset is materially required but unresolved, record the asset need rather than hiding it behind filler.

## 11. Generated Asset Review

Before approval, review generated output for:

### Intent
Does it communicate the requested idea and fit the surface role?

### Truth
Could it imply unsupported facts or be mistaken for documentary evidence?

### Brand
Does it fit Sunlit Editorial / Evidence-Led Optimism?

### Dignity
Does it preserve human dignity and avoid pity-oriented framing?

### Distinctiveness
Does it avoid generic AI/startup visual language?

### Robustness
Can it work in the expected crop, aspect ratio, density, and surrounding content?

## 12. Asset Brief Requirement

Significant Level 3 or Level 4 assets should begin with a brief covering:

- purpose;
- audience;
- visual role;
- emotional goal;
- concept;
- composition;
- required variants;
- must preserve;
- must avoid;
- expected format/aspect ratio;
- approval requirement.

Generation prompts may be preserved when they materially help future regeneration of an approved/canonical style, but exploratory prompt noise should not become documentation.

## 13. Shared Asset Change Impact

Canonical shared assets can have blast radius similar to shared components.

Before materially changing a shared asset:

```text
identify consumers
→ classify visual/semantic impact
→ inspect representative usages
→ verify responsive/light/dark-context implications where applicable
→ update affected surfaces or variants
→ record intentional contract change
```

Do not maintain a stale manual list of every consumer when repository search can discover them at change time.

## 14. Change Classification

### Additive
New asset without changing existing semantics.

### Visual-compatible
Rendering refinement while preserving recognizable meaning and intended use.

### Semantic
Changes what the asset communicates. Requires product/design review.

### Brand-breaking
Changes core identity, illustration language, logo, signature motif, or primary brand treatment. Requires explicit human approval and broad impact analysis.

## 15. Open Asset Decisions

Still OPEN:

- final wordmark/logo;
- exact production palette;
- exact typefaces;
- final icon family;
- final illustration specification;
- placeholder system;
- photography governance details;
- final key art;
- motion assets and motion grammar.

Open decisions must remain visibly open rather than being filled by implementation convenience.
