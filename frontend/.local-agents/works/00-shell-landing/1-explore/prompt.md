/explore landing page (/) — public shell nav + highlighted campaigns
mock, per resolved decisions in docs/ui-ux/page-map.md (2026-08-24).

Read docs/ui-ux/prototype-reference.md's Tier 1 entry for `/` before starting.

Build from docs/design-reference/landing-page.html as near-final visual
reference (Tier 1). Two known issues from prototype-reference.md apply
directly and must NOT carry over:
1. Campaign card image placeholder — the prototype renders it as a
   file-upload dropzone; this is a public read-only card, no upload
   affordance belongs here.
2. Typography sizes drifted from design-guidelines.md (--font-size-h1
   44px vs spec's 30px, --font-size-display 40px vs 36px) — use
   design-guidelines.md's values, not the prototype's.

Follow ../../harscode-workspace/best-practices/react/
ai-prototype-to-production-translation.md for what transfers directly
(composition, states, copy) vs what must be translated (inline styles →
design tokens, the prototype's own scratch component primitives → real
components/ui/, mock data → real API contract per campaign.yaml).