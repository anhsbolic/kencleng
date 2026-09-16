# Public design and truth authorities — Stage 2 evidence

> Phase/Stage: Exploration / Stage 2 — gap analysis  
> Author: Codex  
> Created: 2026-09-16  
> Model: GPT-5  
> Reasoning: not exposed  
> Session: not exposed  
> Target revision: `a30ee75`  
> Workflow revision: candidate baseline `4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Area: approved public experience language and truth boundary

### Current state

- The active design authority map is `docs/ui-ux/README.md`; it selects **Sunlit Editorial / Evidence-Led Optimism** and explicitly separates visual authority from product/domain truth.
- `product-design-principles.md` §§1–15 establishes confidence before conversion, no visual implication without product truth, truth-class separation, non-gamified progress, one clear next action, public expression, responsive priority, and the READY/PARTIAL/OPEN decision boundary.
- `brand-product-ui-brief.md` §§4–8 specifies the public visual language and a representative public experience: editorial composition, restrained CTA, transparency visible early, and real campaign imagery only where true.
- `design-guidelines.md` §§2–13 and §§16–29 supplies concrete visual roles (warm neutrals; restricted Sun/Berry; Newsreader display plus Instrument Sans UI; border/spacing-first grouping; Phosphor baseline; truth/progress grammar) and responsive/accessibility invariants.
- `page-map.md` §1 identifies public home’s purpose as understanding Kencleng/trust model and discovering campaigns, while preserving a strict prohibition on claims or exposure beyond canonical product truth.
- The live frontend does not yet implement any of these public visual, truth, responsive, or asset conventions; it contains only the bootstrap route recorded in Area 1.

### Requirement

- Feature spec **Goal**, **Requirements**, and **Acceptance criteria** require the representative `/` slice to coherently express the current authorities while keeping unresolved Campaign/ranking/curation/trust/evidence semantics explicit or provisional.
- `product-design-principles.md` §§1–2, §5, §7–8, §12–15 require enough trust context before action, prohibit unsupported ranking/verification/urgency/outcome implications, keep truth classes and progress types distinct, preserve responsive priority, and escalate material new interaction/navigation/brand decisions.
- `brand-product-ui-brief.md` §8 requires the direction evidence to include a clear restrained CTA and early transparency/trust context; §7 requires imagery/illustration never to impersonate campaign evidence.
- `design-guidelines.md` §§3–5, §7, §11–13, §16–23, §27 requires the concrete visual roles and preserves truthfulness, hierarchy, accessibility, responsive information priority, and dignity. Its §26 retains final wordmark/logo, full illustration family, photography detail, provenance labels, campaign placeholder implementation, and motion tokens as OPEN.
- `asset-governance.md` §§1–8 requires asset classification, truthfulness, and human approval before a Level 4 asset becomes canonical; `page-map.md` §1 bounds the public home’s purpose without deciding its exact route composition.

### Gap

The repository has approved visual and truth constraints but no rendered implementation expressing them. The feature intentionally leaves the minimum route composition open; no live route currently supplies the required evidence for human judgement. Several candidate content/asset treatments remain product/design-open and cannot be made canonical by this task without the applicable human decision or truthful product source.

### Sniffing

- **Risk:** As the first production precedent, incorrect visual language or an implied trust/evidence claim could be copied into every later public surface. The highest-impact concern is false confidence (e.g., “verified,” outcome, popular/trending, or documentary-looking synthetic media) rather than a local styling defect.
- **Edge cases:** Missing campaign media, unknown/pending information, long/realistic copy, monetary/progress values, and narrow screens must remain truthful and readable; color alone cannot communicate status. The source forbids treating funding completion as outcome proof.
- **Miscontext:** The representative public reference demonstrates direction, not a complete home-page contract or campaign data permission. The page map’s discovery goal does not authorize a ranking/curation meaning.
- **Misleading signals:** Yellow, badges, progress bars, imagery, and an early trust section can look like proof. The authorities explicitly say they are not proof; visual references likewise are not implementation or domain authority.
- **Inconsistency:** No authority conflict found. The public-home discovery intent is bounded by the same product-truth prohibition in the feature spec, design principles, and page map. The unresolved composition/asset decisions are documented OPEN rather than contradictory.

### Code and authority anchors

- `docs/spec/0-foundations/features/01-frontend-experience-foundation.md` — Goal, Requirements, Acceptance criteria, Explicit non-goals: task-specific source and truth boundary.
- `docs/ui-ux/product-design-principles.md` — §§1–2, §5, §§7–8, §§12–15: later verification criteria for truthful hierarchy, responsive priority, and decision escalation.
- `docs/ui-ux/brand-product-ui-brief.md` — §§4–8: approved public expression and representative-experience intent.
- `docs/ui-ux/design-guidelines.md` — §§3–5, §7, §§11–13, §§16–23, §§26–29: concrete visual rules/open decisions and later visual verification anchors.
- `docs/ui-ux/asset-governance.md` — §§1–8: future asset classification, lifecycle, and approval boundary.
- `docs/ui-ux/page-map.md` — §1: public-home purpose and guest/public truth restriction.
- `app/page.tsx` — `Home`: no current implementation of the above authority.
