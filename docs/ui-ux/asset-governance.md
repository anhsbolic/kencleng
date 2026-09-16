# Kencleng — Asset Governance

> Status: Canonical
> Last updated: 2026-09-16
> Purpose: Govern truthfulness, approval, lifecycle, reuse, and change impact for visual assets.
> Relationship: `brand-product-ui-brief.md` owns the approved direction; `design-guidelines.md` owns the concrete visual system; this document owns how assets enter and evolve within that system.

## 1. Asset classes

Classify an asset before creating or integrating it.

### Level 1 — Utility

Universal interface actions such as close, search, menu, edit, filter, calendar, chevron, password visibility, external link, and download.

Default: use familiar established iconography. The current utility-icon baseline is **Phosphor**, as defined by `design-guidelines.md`. Custom art is normally unnecessary for utility actions.

### Level 2 — Product-semantic

Visual treatment for domain concepts such as organization verification, campaign state, donation, fund usage, disbursement, reporting, or representative role.

Use a consistent treatment where repeated meaning exists. A standard icon may be sufficient when it communicates the concept truthfully; visual treatment must not upgrade or invent domain meaning.

### Level 3 — Expressive product asset

Visuals that add meaning, humanity, onboarding, explanation, or product character, such as empty-state illustration, campaign placeholder, explanatory graphic, landing-section art, or trust/accountability illustration.

Do not silently downgrade a meaningful expressive need into arbitrary generic filler.

### Level 4 — Brand-defining asset

Logo, wordmark, core illustration language, major key art, signature motif, or another identity-defining system.

Human approval is required before a new Level 4 asset becomes canonical.

## 2. Truth before expression

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

Synthetic documentary-style beneficiary imagery must not be used as if it were campaign evidence.

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

Detailed production photography governance remains an open follow-up where real campaign-media workflows require it.

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

The high-level illustration direction is approved. A full production illustration-family specification remains open until real asset work justifies it.

## 5. Campaign and organization placeholders

Missing media must still feel intentional.

A placeholder must:

- be clearly recognizable as a placeholder;
- not impersonate campaign evidence;
- preserve layout/aspect-ratio stability;
- remain subordinate to real content;
- fit Sunlit Editorial.

The canonical visual guidelines already establish the placeholder direction: warm-neutral surface, abstract editorial geometry, restrained Sun/Berry detail, and no fake person/documentary scene.

The exact production placeholder asset/system remains open and should be reviewed before becoming a reusable canonical asset.

## 6. Iconography

Canonical utility baseline: **Phosphor Icons**.

Use it for ordinary navigation, utility actions, and standard UI semantics unless a concrete product need justifies a different treatment.

Product-semantic icon treatment should remain consistent and restrained.

Avoid shield/checkmark/lock/certification vocabulary as a substitute for actual trust information.

Whether specific Phosphor icons require project-specific tuning remains an implementation-discovered design question; do not create a second competing icon family casually.

## 7. Wordmark / logo

No incidental implementation wordmark or placeholder logo may silently become canonical.

The identity should not depend on a literal kencleng/celengan metaphor and should remain viable if the product name changes.

A final wordmark/logo remains OPEN and requires human approval.

## 8. Asset lifecycle

Use these states:

```text
EXPLORATION
PROPOSED
APPROVED
CANONICAL
SUPERSEDED
```

### EXPLORATION
Working evidence. Not production authority.

### PROPOSED
Candidate intended for review.

### APPROVED
Reviewed for a stated use. Approval may be surface-specific.

### CANONICAL
Established reusable visual authority. Future work should reuse or intentionally extend it before creating a competing direction.

### SUPERSEDED
No longer current. Historical value belongs in Git history or clearly historical records, not as competing active-tree authority.

A file existing in the repository is not automatically canonical.

## 9. Selected-direction references

The approved Sunlit Editorial public composition and Evidence Journal composition under:

```text
docs/ui-ux/visual-references/selected-direction/
```

are **Selected Direction References**.

They are approved evidence for:

- brand character;
- expressive intensity;
- composition principles;
- trust/information hierarchy;
- public-versus-product translation.

They are not production campaign assets, pixel specifications, component contracts, or evidence of domain behavior.

## 10. No generic filler

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

## 11. Generated asset review

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

## 12. Asset brief requirement

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

Generation prompts may be preserved when they materially help future regeneration of an approved/canonical style. Exploratory prompt noise should not become canonical documentation.

## 13. Shared asset change impact

Canonical shared assets can have blast radius similar to shared components.

Before materially changing one:

```text
identify consumers
→ classify visual/semantic impact
→ inspect representative usages
→ verify responsive/context implications where applicable
→ update affected surfaces or variants
→ record intentional contract change
```

Do not maintain a stale manual list of every consumer when repository search can discover them at change time.

## 14. Change classification

### Additive
New asset without changing existing semantics.

### Visual-compatible
Rendering refinement while preserving recognizable meaning and intended use.

### Semantic
Changes what the asset communicates. Requires product/design review.

### Brand-breaking
Changes core identity, illustration language, logo, signature motif, or primary brand treatment. Requires explicit human approval and broad impact analysis.

## 15. Current open asset decisions

Still intentionally OPEN:

- final wordmark/logo;
- full production illustration-family specification;
- detailed campaign photography governance;
- exact production campaign-placeholder asset/system;
- final key art where a future surface materially requires one;
- motion assets / detailed motion grammar;
- project-specific Phosphor tuning if real implementation demonstrates a need.

Already resolved by the canonical visual system and **not open for incidental implementation re-selection**:

- Sunlit Editorial palette roles and exact core values;
- Newsreader / Instrument Sans pairing and role separation;
- Phosphor as the utility-icon baseline;
- overall editorial-documentary photography direction;
- overall mature human editorial illustration direction.

Open decisions must remain visibly open rather than being filled by implementation convenience.