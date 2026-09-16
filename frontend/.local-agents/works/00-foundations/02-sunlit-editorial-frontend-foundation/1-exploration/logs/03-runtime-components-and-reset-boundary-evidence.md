# Stage 2 — Runtime, components, and reset-boundary evidence

> Phase: Exploration
> Stage: 2 — Gap Analysis
> Author: ChatGPT
> Model: GPT-5.6 Sol
> Reasoning: High
> Created: 2026-09-16
> Target revision: `cc6552e5ae36e354388f5d9d3230672929b11055`
> Workflow revision: `4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Area 6 — Runtime architecture versus presentation architecture

### Current state

The root application contains several concerns that are independent of the old visual system:

- App Router route groups;
- `QueryProvider` with a per-mount `QueryClient`;
- `MockingProvider` that gates MSW only in non-production mock mode;
- authentication bootstrap/session providers;
- typed API functions generated from OpenAPI shapes;
- TanStack Query hooks such as `useCampaigns`;
- shared accessibility behavior such as `useFocusTrap`;
- Vitest/RTL/MSW and Playwright infrastructure.

These concerns are not evidence for keeping the existing rendered UI. They are evidence that “frontend reset” and “delete the entire frontend runtime” are different decisions.

### Requirement

The human owner explicitly wants a clean design foundation and does not want sunk-cost preservation to contaminate the new direction. Reuse must earn itself through design-agnostic correctness or stable product/runtime behavior.

### Gap

The current codebase mixes two very different classes of legacy:

1. **visual/presentation legacy** — fonts, token vocabulary, icon library, primitive styling, shell composition, landing composition, campaign-card presentation and old copy;
2. **runtime/behavioral infrastructure** — API/state/provider/mock/test/accessibility mechanisms that do not depend on the old brand.

Treating both classes identically would be a false reset boundary. A clean frontend foundation can aggressively replace presentation while retaining independently valid runtime behavior.

### Sniffing

- **Risk:** deleting correct auth/query/mock/accessibility plumbing merely to obtain a “clean” directory would increase regression risk without improving design purity.
- **Edge cases:** retained behavior may still have historical comments or tests referencing old Techplans; semantic correctness and documentation cleanliness are separate questions.
- **Miscontext:** preserving a hook/provider does not imply preserving the route/component structure that currently consumes it.
- **Misleading signals:** a file living under `components/` does not make it presentation legacy; providers and shared interaction hooks may be design-agnostic.
- **Inconsistency:** root layout provider composition is structurally useful while the same root layout carries obsolete fonts and language metadata.

### Evidence anchors

- `frontend/app/layout.tsx`
- `frontend/components/providers/query-provider.tsx`
- `frontend/components/providers/mocking-provider.tsx`
- `frontend/components/providers/auth-bootstrap-provider.tsx`
- `frontend/lib/hooks/use-campaigns.ts`
- `frontend/lib/api/campaign.ts`
- `frontend/lib/hooks/use-focus-trap.ts`

## Area 7 — UI/shared primitives and blast radius

### Current state

`frontend/components/README.md` correctly treats `components/ui/` and `components/shared/` as reusable internal contracts with consumer-impact obligations.

The registered UI layer includes `Button`, `Input`, `Label`, `Badge`, `Banner`, `ProgressBar`, and `Spinner`. Current source shows that at least several foundation primitives are materially coupled to the old visual system:

- `Button` maps `primary` to green `primary-600`, uses legacy radius and green focus ring;
- `Badge` uses legacy neutral/accent/semantic token families and pill-heavy treatment;
- `ProgressBar` uses legacy neutral track, green primary fill, and switches to success green at 100%;
- current public shell and landing code also bypass the primitive layer in places with direct green/neutral utility classes.

`package.json` currently depends on `lucide-react`; the canonical visual system specifies Phosphor as the utility-icon baseline, and no Phosphor package is currently present.

### Requirement

The new foundation must not silently reinterpret old primitive variants as canonical Sunlit Editorial semantics. Broad primitive changes require actual consumer discovery and representative downstream verification.

### Gap

The primitive *concepts* may remain useful, but their source implementations are not visually neutral. A clean reset must decide per primitive whether to:

- replace implementation while preserving a stable generic behavioral API;
- intentionally break/redefine the API and migrate consumers;
- defer a primitive until repeated legitimate usage proves the abstraction.

That decision cannot be made safely by changing `globals.css` underneath all existing variants and assuming compatibility.

The current execution surface did not provide a reliable repository-wide symbol/reference index, so complete consumer counts are not claimed here. Techplan/Build must perform mechanical consumer discovery before any broad primitive contract is changed or removed.

### Sniffing

- **Risk:** globally remapping `primary-*` to Sun could make every old “primary” usage look superficially new while preserving obsolete action semantics.
- **Edge cases:** accessibility/focus/loading behavior may be worth preserving even when variant names and visual styles change.
- **Miscontext:** “reset primitives” does not necessarily mean every low-level behavior must be rewritten; “preserve Button” does not mean its visual API is still correct.
- **Misleading signals:** passing primitive unit tests may only prove old contracts, not suitability for the new design system.
- **Inconsistency:** canonical iconography is Phosphor, while public/UI code currently imports Lucide and the dependency set has no Phosphor package.

### Evidence anchors

- `frontend/components/README.md`
- `frontend/components/ui/button.tsx`
- `frontend/components/ui/badge.tsx`
- `frontend/components/ui/progress-bar.tsx`
- `frontend/package.json`
- `frontend/app/(public)/layout.tsx`
- `frontend/app/(public)/_components/public-shell-client.tsx`
- `frontend/components/features/landing/how-it-works.tsx`

## Reset-boundary evidence — not yet a Stage-3 strategy

The Stage-2 evidence strongly separates the current frontend into the following pressure zones.

### High replacement pressure

- root font pair and font aliases;
- current color/token vocabulary and values;
- legacy radius/elevation defaults;
- Lucide-based brand/utility presentation where current surface is rebuilt;
- public shell visual composition;
- `/` route composition and landing presentation components;
- campaign-card visual treatment;
- current generic primitive styling/variant assumptions that encode old visual semantics;
- stale design-era comments/document routing;
- tests whose only contract is obsolete presentation.

### Strong retention evidence at behavioral/runtime level

- Next.js/App Router scaffold;
- typed OpenAPI/API-client boundary;
- TanStack Query server-state ownership and query hooks where product behavior remains valid;
- RHF/Zod architecture where relevant;
- Zustand only for genuinely shared client state;
- auth/session plumbing that remains product-correct;
- MSW mock infrastructure;
- `useFocusTrap`-style accessibility behavior where the new composition still needs overlays/drawers;
- Vitest/RTL/MSW and Playwright capability/configuration.

### Still requires Stage 3 decision

- exact public shell/navigation composition;
- whether current auth-modal entry is retained, redesigned, or excluded from the minimum slice;
- which old generic primitive APIs deserve compatibility versus intentional reset;
- whether representative campaign content remains in the minimum slice and what truthful data subset is used;
- exact minimum `/` slice;
- exact implementation sequence for tokens/fonts/primitives/route composition;
- exact verification allocation across Build, Testing, and human acceptance.
