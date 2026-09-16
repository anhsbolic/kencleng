# Stage 2 — Live public surface and visual foundation

> Phase: Exploration
> Stage: 2 — Gap Analysis
> Author: ChatGPT
> Model: GPT-5.6 Sol
> Reasoning: High
> Created: 2026-09-16
> Target revision: `cc6552e5ae36e354388f5d9d3230672929b11055`
> Workflow revision: `4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Area 3 — Current `/` and public shell

### Current state

`frontend/app/(public)/page.tsx` composes three legacy landing sections: `Hero`, `HighlightedCampaigns`, and `HowItWorks`. The file comment still cites an old local-agent Techplan as scope authority.

`frontend/app/(public)/layout.tsx` supplies a fixed public shell with:

- text wordmark plus `HandHeart` Lucide icon;
- green `primary-600` logo block;
- desktop navigation;
- `Masuk` / `Daftar` auth-modal triggers;
- mobile hamburger/drawer;
- one mounted `AuthModal`.

The shell interaction structure includes legitimate accessibility work: the mobile drawer uses `aria-expanded`, `aria-controls`, `role="dialog"`, `aria-modal`, Escape handling and focus trapping through `useFocusTrap`.

`publicNavItems` currently maps `Jelajahi Kampanye` to the local `#kampanye` anchor because a dedicated campaign route does not exist. Current `page-map.md` intentionally no longer owns exact route paths or shell architecture, so this implementation is not canonical simply because it exists.

### Requirement

The new public foundation must express Sunlit Editorial / Evidence-Led Optimism and preserve truthful product semantics. Public composition may be expressive, but exact shell architecture and route paths are engineering/product decisions rather than inherited legacy contracts.

### Gap

The public shell is visually and compositionally tied to the superseded generation: cool-white/cool-neutral treatment, old green identity, Lucide iconography, old typography, rounded SaaS-like controls, and historical comments. Its interaction mechanics may contain reusable behavior, but its rendered shell should not be treated as a migration target by default.

The existing auth-modal architecture is current behavior, not a design authority. Whether the minimum foundation slice retains the modal interaction or uses another current-product-consistent entry remains an engineering/product decision; Exploration must not preserve it solely because the old shell already does so.

### Sniffing

- **Risk:** preserving shell markup/styles could anchor the new system around the old brand before the foundation is established.
- **Edge cases:** mobile navigation/focus behavior is valuable behavior even if the visual shell is rebuilt; deleting presentation should not accidentally regress keyboard/focus semantics.
- **Miscontext:** comments citing `.local-agents/works/00-shell-landing/...` are historical implementation evidence, not current project authority.
- **Misleading signals:** the shell is functionally complete enough to feel reusable, but its visual/component assumptions conflict with the approved direction.
- **Inconsistency:** root `app/layout.tsx` still declares `lang="en"` while the representative product UI is Indonesian.

### Anchors

- `frontend/app/(public)/page.tsx`
- `frontend/app/(public)/layout.tsx`
- `frontend/app/(public)/_components/nav-items.ts`
- `frontend/app/(public)/_components/public-shell-client.tsx`
- `frontend/components/features/account/auth-modal-triggers.tsx`
- `frontend/app/layout.tsx`

## Area 4 — Root fonts, globals.css, Tailwind v4 tokens

### Current state

`frontend/app/layout.tsx` loads `Plus Jakarta Sans` and `Inter` through `next/font/google`; the active canonical design system now requires Newsreader and Instrument Sans role separation.

`frontend/app/globals.css` is CSS-first Tailwind v4 configuration and is structurally the correct implementation surface for theme variables, but its values are legacy:

- green `primary-*` brand scale;
- amber `accent-*`;
- cool-gray `neutral-*`;
- legacy success/warning/error/info values;
- old typography scale;
- old radius family `8/12/16/24`;
- three legacy shadow values;
- `--font-sans: Inter` and `--font-heading: Plus Jakarta Sans`.

The file comment incorrectly claims these values are per current `docs/ui-ux/design-guidelines.md` even though the current canonical guideline now defines warm-paper neutrals, Sun/Berry, different semantic values, Newsreader/Instrument Sans, radius `6/10/14/20`, and a flatter elevation posture.

### Requirement

`design-guidelines.md` owns visual meaning. `globals.css` may own the implementation representation of approved tokens, but existing variable names/values do not become authoritative merely because they are already wired to Tailwind.

### Gap

The CSS-first mechanism is reusable; the actual visual token layer requires deliberate reconstruction. A superficial rename risks carrying old semantics into the new system. The reset must distinguish:

- reusable Tailwind v4 CSS-first mechanism;
- replaceable token vocabulary/values/aliases;
- downstream classes that currently assume `primary`, `neutral`, legacy radius, and legacy shadows.

The root font setup also requires replacement rather than adaptation of the existing pair.

### Sniffing

- **Risk:** keeping generic names such as `primary` without rethinking semantics can recreate “primary = brand color forever”, explicitly rejected by the new contextual action hierarchy.
- **Edge cases:** semantic colors must remain separate from Sun/Berry; numeric/operational text must use Instrument Sans even where public headings use Newsreader.
- **Miscontext:** “canonical tokens live in globals.css” means implementation location, not that current values are canonical.
- **Misleading signals:** existing Tailwind utilities make the old system easy to reuse accidentally.
- **Inconsistency:** current token comments point at the current guideline while encoding values that directly contradict it.

### Anchors

- `frontend/app/layout.tsx`
- `frontend/app/globals.css`
- `docs/ui-ux/design-guidelines.md` — Color, Typography, Spacing, Surfaces, Radius, Elevation, Action Hierarchy, Public vs Product Expression

## Area 5 — Current landing content/card semantics

### Current state

`Hero` still carries prototype-era composition and green visual language. It presents a success badge `Organisasi terverifikasi`, old bold sans heading treatment, and green primary CTA.

`HighlightedCampaigns` uses `GET /campaigns` through `useCampaigns`, which is a structurally valid TanStack Query boundary. Its loading/error/empty/background-refresh behavior follows current list/browse interaction principles. However, the rendered copy says `Kampanye pilihan` and `Sedang berjalan minggu ini`, while the current public list contract does not establish a featured/selected ranking or week-based curation semantic.

`CampaignCard` is a read-only preview backed by list data. Its explicit refusal to invent organization data/media is directionally correct, but its current presentation is legacy card-first UI: rounded bordered white card + shadow, Lucide `ImageOff`, legacy green progress treatment, and `ProgressBar` that switches to semantic success green at 100%.

`HowItWorks` carries legacy static claims. It says the minimum donation is Rp10,000, while the current donation feature source defines `amount >= 5000`. It also claims `tanpa biaya tersembunyi` and periodic fund-use reports without evidence in the inspected owning sources. These are product-truth defects, not styling details.

### Requirement

The representative `/` slice may show campaign/discovery content only with semantics supported by current domain/API authority. Trust must come from structure/provenance rather than unsupported curation, verification theatre, or optimistic copy.

### Gap

The existing route has useful behavior/data boundaries but its visible composition/copy are not a safe base for the new design. Presentation reset pressure is high because the same files combine stale visual language with unsupported or historically inferred product claims.

### Sniffing

- **Risk:** carrying old copy into a polished new design would make unsupported claims more credible, not less.
- **Edge cases:** public campaign list data may lack real imagery and organization summary; a new card must remain truthful when those fields are absent rather than fabricating documentary context.
- **Miscontext:** `HighlightedCampaigns` is an implementation name, not evidence that a canonical “highlighted” campaign concept exists.
- **Misleading signals:** `Target tercapai` plus a success-green fill can visually suggest broader completion even though the new design explicitly separates funding progress from operational progress and reported outcome.
- **Inconsistency:** `HowItWorks` Rp10,000 copy conflicts with the donation source's Rp5,000 minimum.

### Anchors

- `frontend/components/features/landing/hero.tsx`
- `frontend/components/features/landing/highlighted-campaigns.tsx`
- `frontend/components/features/landing/how-it-works.tsx`
- `frontend/components/features/campaign/campaign-card.tsx`
- `frontend/components/ui/progress-bar.tsx`
- `frontend/lib/hooks/use-campaigns.ts`
- `frontend/lib/api/campaign.ts`
- `docs/ui-ux/patterns.md` — List/Browse, Progress, Loading/Error
- `docs/spec/5-donation/features/01-submit-donation-settlement.md`
- `docs/spec/4-campaign/features/02-campaign-detail-listing.md`
