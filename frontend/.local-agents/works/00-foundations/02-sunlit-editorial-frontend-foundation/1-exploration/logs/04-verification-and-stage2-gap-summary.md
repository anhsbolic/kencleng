# Stage 2 — Verification surface and gap summary

> Phase: Exploration
> Stage: 2 — Gap Analysis
> Author: ChatGPT
> Model: GPT-5.6 Sol
> Reasoning: High
> Created: 2026-09-16
> Target revision: `cc6552e5ae36e354388f5d9d3230672929b11055`
> Workflow revision: `4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Area 8 — Verification requirements and economics

### Current state

Frontend verification infrastructure already exists:

- `npm run verify` → lint + Vitest;
- `npm run build` → production build;
- Playwright Chromium capability through `npm run test:browser` and related commands;
- human rendered acceptance required by `frontend/AGENTS.md` and the foundation spec for material UI.

The active main branch currently has only `tests/browser/login.smoke.spec.ts`; no `/` browser regression is committed at this baseline.

Current landing components have unit/component tests, but some tests encode obsolete presentation contracts. For example `hero.test.tsx` explicitly asserts legacy `text-h1` / `md:text-display` classes as the former design-guideline contract. Such tests are evidence of the old implementation contract, not reasons to preserve it.

The foundation spec requests representative Playwright/browser automation **where it provides useful repeatable evidence**. `frontend/AGENTS.md` separately states that Playwright must not be added/expanded/run merely because a task reached a lifecycle phase; browser automation needs a concrete behavior/value reason and never replaces human rendered acceptance.

### Requirement

Verification for the new foundation must prove different classes of correctness with proportional ownership:

- mechanical/type/lint/build integrity;
- component/interaction behavior where appropriate;
- representative desktop/mobile rendered behavior;
- browser behavior only where a meaningful repeatable risk warrants automation;
- human product/design acceptance for brand/hierarchy/asset appropriateness.

The refined Harscode workflow assigns focused fast verification to Build, targeted reproduction to Code Review when needed, and independent final/broad verification to Testing.

### Gap

The old frontend has test infrastructure worth retaining, but the verification *contract* for the new foundation must be re-derived from the new slice rather than inherited from old tests.

Playwright is available but is not yet justified for a specific new `/` behavior because Stage 3 has not selected the exact minimum slice. If the selected slice includes behavior such as mobile disclosure/drawer focus, anchor/navigation semantics, or another cross-component interaction that component tests cannot credibly establish, repeatable browser automation may be justified. Pure visual judgment remains human/rendered-review work.

### Sniffing

- **Risk:** preserving obsolete class/token assertions would force implementation toward the old visual system.
- **Edge cases:** a visually new shell may preserve old focus-management behavior; tests should assert behavior rather than old DOM/class structure.
- **Miscontext:** having Playwright installed does not make every responsive check a permanent E2E test.
- **Misleading signals:** a green unit suite can coexist with visually wrong hierarchy or stale product copy.
- **Inconsistency:** the foundation spec asks for useful browser automation while current main has no home browser test; this is a deliberate decision gap to resolve after the minimum slice is chosen, not an automatic failure.

### Anchors

- `frontend/package.json`
- `frontend/playwright.config.ts`
- `frontend/tests/browser/login.smoke.spec.ts`
- `frontend/components/features/landing/hero.test.tsx`
- `frontend/AGENTS.md` — Rendered iteration / Testing and Playwright boundary
- `docs/spec/0-foundations/features/01-frontend-experience-foundation.md`
- Harscode `workflow/3-build*`, `workflow/4-code-review*`, `workflow/5-testing*` at `4199c6d`

## Stage-2 gap summary

### Confirmed authority/documentation gaps

1. `docs/project/kencleng-frontend-tech-stack.md` still contains removed design/prototype authority references and old green action hierarchy.
2. `asset-governance.md` has stale OPEN-list items for visual-system concerns now canonically resolved by `design-guidelines.md`.
3. `visual-references/selected-direction/README.md` still says binary handoff is pending / `.jpg` expected although the approved `.png` references exist in the tree.
4. `globals.css` comments claim current design-guideline alignment while its actual values are from the superseded visual generation.

### Confirmed live-implementation gaps

1. Root fonts are Plus Jakarta Sans + Inter instead of Newsreader + Instrument Sans.
2. Root `lang` remains `en` for an Indonesian representative surface.
3. Public shell remains old green/cool-neutral/Lucide visual language.
4. Current `/` composition and landing components carry prototype/old-Techplan assumptions.
5. Current generic primitives materially encode old visual semantics.
6. Current icon dependency is Lucide; Phosphor baseline is not present in `package.json`.
7. Current campaign/progress presentation conflicts with the new “Progress as Evidence” visual grammar.
8. Current `HighlightedCampaigns` copy implies selection/week semantics not established by inspected public listing authority.
9. `HowItWorks` states Rp10,000 minimum while the current donation source defines Rp5,000, plus other claims that need owning authority before reuse.
10. Some existing tests intentionally lock old class/token contracts and therefore cannot be carried forward unchanged as acceptance evidence.

### Confirmed structurally valuable foundations

Evidence supports retaining the *conceptual/runtime correctness* of the following unless Stage 3 discovers a task-specific contradiction:

- Next.js/App Router scaffold;
- API/OpenAPI boundary and generated types;
- API functions + TanStack Query server-state ownership;
- auth/session runtime plumbing;
- mock-first MSW capability;
- semantic-owner-first component governance;
- accessibility behavior such as focus trapping where still needed;
- Vitest/RTL/MSW and Playwright infrastructure;
- CSS-first Tailwind v4 mechanism (not current token contents).

This is not a requirement to preserve existing file/component presentation structure. It is evidence that a clean visual/frontend-foundation reset can remain selective about non-visual runtime contracts.

### Material questions intentionally left for Stage 3

Stage 2 does **not** select answers yet for:

1. the exact reset boundary and implementation sequence;
2. the exact minimum coherent `/` slice;
3. whether the minimum slice includes current auth-modal behavior;
4. whether representative campaign content stays live-query-backed, uses a narrower truthful presentation, or is deferred from the calibration slice;
5. which primitive APIs are preserved, intentionally broken, or not recreated until usage proves need;
6. which browser behavior is worth permanent Playwright automation;
7. how documentation reconciliation is sequenced relative to code reset;
8. which OPEN asset needs, if any, are triggered by the minimum slice.

## Stage-2 completion assessment

The evidence is sufficient to enter solutioning without reopening the approved brand direction. The central engineering problem is no longer “should Sunlit Editorial replace the old design?”; that is already decided. The Stage-3 problem is how to establish a clean Sunlit Editorial frontend foundation while preserving only runtime/behavioral contracts that independently remain correct.

No production code was modified in Stage 2. No local lint/unit/build/browser command was executed in this ChatGPT exploration environment; source and repository evidence above were inspected through the connected repository surface. Runtime verification remains a downstream obligation, not a claimed Stage-2 result.
