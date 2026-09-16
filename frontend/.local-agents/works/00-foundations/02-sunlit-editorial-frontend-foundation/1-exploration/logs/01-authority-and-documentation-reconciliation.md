# Stage 2 — Authority and documentation reconciliation

> Phase: Exploration
> Stage: 2 — Gap Analysis
> Author: ChatGPT
> Model: GPT-5.6 Sol
> Reasoning: High
> Created: 2026-09-16
> Target revision: `cc6552e5ae36e354388f5d9d3230672929b11055`
> Workflow revision: `4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Human constraint carried into this exploration

The existing frontend has no preservation privilege. Reuse is justified only when an existing layer remains structurally valuable under current authority. Existing presentation must not be retained merely because it already exists. This does not authorize discarding valid product/API/runtime contracts.

## Area 1 — Canonical design authority vs foundation requirement

### Current state

Active authority routes through `docs/ui-ux/README.md`. The approved direction is **Sunlit Editorial / Evidence-Led Optimism**. `brand-product-ui-brief.md` owns the selected brand/Product UI direction; `design-guidelines.md` owns the concrete visual system; `product-design-principles.md` owns reusable experience judgment; `patterns.md` owns recurring interaction behavior; `asset-governance.md` owns asset lifecycle/truthfulness; `page-map.md` owns persona/surface intent.

The concrete system is sufficiently defined for engineering interpretation: Newsreader for editorial/display use, Instrument Sans for UI/body/operational use, warm paper-like neutrals, restrained Sun yellow, Berry as mature secondary accent, independent semantic colors, border/spacing-first grouping, restrained radius/elevation, Phosphor utility icons, provenance/truth grammar, and distinct funding/operational/reported-outcome progress grammar.

The foundation spec keeps `/` as a calibration surface and deliberately leaves the minimum coherent rendered slice to Exploration. It explicitly does not require a complete landing page, final Campaign semantics, or unjustified global abstraction.

### Requirement

The implementation must make enough of `/` real to judge brand expression, hierarchy, typography, color/surfaces, CTA hierarchy, representative content/card treatment, responsiveness, and relevant component use without inventing product semantics. Material UI still requires human rendered acceptance.

### Gap

The design direction itself is no longer the missing input. The remaining gap is an engineering translation gap: current implementation and architecture documentation were built against a superseded visual generation and have not yet been reconciled to the approved system.

Design readiness for the foundation task is therefore best treated as **READY/PARTIAL at implementation-detail level**, not OPEN brand exploration. Remaining OPEN items are bounded asset/detail concerns, not permission to reopen Sunlit Editorial.

### Sniffing

- **Risk:** treating current UI code as precedent would silently reintroduce green identity, old fonts, old card/radius/shadow language, and trust-badge patterns into a newly approved system.
- **Edge cases:** final logo/wordmark, detailed illustration family, detailed photography governance, exact motion tokens, final provenance terminology, and exact placeholder implementation remain open and must not be filled by convenience.
- **Miscontext:** selected visual references are direction evidence, not screenshot specifications or business truth.
- **Misleading signals:** polished legacy UI can look production-ready while still encoding superseded authority.
- **Inconsistency:** `asset-governance.md` still lists exact palette, typefaces, and final icon family as OPEN even though `design-guidelines.md` now canonically fixes the palette, Newsreader/Instrument Sans roles, and Phosphor baseline. The authority map makes `design-guidelines.md` the owner of those visual-system concerns, so the open list is stale and should be reconciled.

### Authority anchors

- `docs/spec/0-foundations/features/01-frontend-experience-foundation.md`
- `docs/ui-ux/README.md`
- `docs/ui-ux/brand-product-ui-brief.md`
- `docs/ui-ux/product-design-principles.md`
- `docs/ui-ux/design-guidelines.md`
- `docs/ui-ux/patterns.md`
- `docs/ui-ux/asset-governance.md`
- `docs/ui-ux/page-map.md`

## Area 2 — Frontend architecture/documentation reconciliation

### Current state

`docs/project/kencleng-frontend-tech-stack.md` still contains a mixture of structurally useful frontend architecture and superseded design-era assumptions.

Still structurally aligned:

- Next.js App Router / React / TypeScript / Tailwind v4 stack;
- API/OpenAPI authority separation;
- TanStack Query server-state ownership;
- RHF/Zod form ownership;
- Zustand only for genuinely shared client-owned state;
- semantic-owner-first component placement;
- shared-component blast-radius discipline;
- CSS-first Tailwind v4 implementation model;
- rendered verification layered with unit/component/browser/human evidence.

Stale design-era sections include:

- `G. UI/UX Knowledge Architecture` routes to removed `brand-and-visual-assets.md`, `prototype-reference.md`, and `design-reference-usage.md`;
- `I. Visual Asset Readiness` routes to removed `brand-and-visual-assets.md`;
- `K. Action Hierarchy v2` declares filled Kencleng green as Primary;
- `L. Prototype Translation` still treats `design-reference/` as route-specific prototype authority;
- `J. Visual Implementation` says canonical tokens live in current `frontend/app/globals.css`, but that file still contains the superseded green/cool-neutral system and therefore cannot currently be treated as design truth.

### Requirement

Frontend architecture docs must route engineering toward current authorities without allowing stale narrative to override `docs/ui-ux/README.md` / `design-guidelines.md`.

### Gap

The architecture document is partly correct structurally but unsafe as a whole until stale design references and assumptions are reconciled. The correct migration is not to discard its state/data/component architecture; it is to remove or rewrite stale visual/design routing so future agents do not receive contradictory instructions.

### Sniffing

- **Risk:** an implementation agent following the architecture doc literally could restore the old green hierarchy or search removed prototype authorities.
- **Edge cases:** `globals.css` remains the implementation location for CSS-first tokens, but current values are not canonical; wording must distinguish implementation location from current design authority.
- **Miscontext:** “frontend reset” does not imply resetting API/state ownership rules that are independent of visual design.
- **Misleading signals:** a document marked Draft v2 can contain both high-value architecture and obsolete visual policy; treating it as entirely stale or entirely current would both be wrong.
- **Inconsistency:** current `frontend/AGENTS.md` correctly routes to the new authority while `kencleng-frontend-tech-stack.md` still points at removed authority names.

### Documentation-specific reconciliation finding

`docs/ui-ux/visual-references/selected-direction/README.md` says binary handoff is pending and expects `.jpg` filenames, while the active tree already contains `sunlit-editorial-public.png` and `sunlit-editorial-evidence-journal.png`. This is stale status/filename prose, not a design reopening.

### Code/document anchors

- `docs/project/kencleng-frontend-tech-stack.md`
- `frontend/AGENTS.md`
- `frontend/app/globals.css`
- `docs/ui-ux/visual-references/selected-direction/README.md`
- `docs/ui-ux/visual-references/selected-direction/sunlit-editorial-public.png`
- `docs/ui-ux/visual-references/selected-direction/sunlit-editorial-evidence-journal.png`
