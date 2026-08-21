# Kencleng — Frontend Design Guidelines

> Intended path: `docs/ui-ux/design-guidelines.md`
> Status: Resolved — originally Step 10 of `kencleng-roadmap-next-steps.md`.
> Moved into `docs/ui-ux/` 2026-08-20 (was `docs/project/`) — grouped
> with `page-map.md` and `patterns.md` as the frontend-UX doc set.
> Content unchanged from the 2026-07-27 version except cross-references.
> Last updated: 2026-08-20

## Context

This document is the visual design layer sitting on top of the
structural decisions defined elsewhere:

- `page-map.md` — per-persona, per-route page inventory (the "which
  page")
- `patterns.md` — reusable page shapes & state handling (the "what
  shape")
- `kencleng-frontend-tech-stack.md` — code architecture (the "how
  it's built")
- **This doc** — the visual layer: colors, typography, spacing,
  shape, elevation, and how they map onto concrete component states
  (the "what it actually looks like")

Brand direction: **warm & charitable** — approachable and a little
playful, distinct from the cooler/more corporate feel of a typical
fintech product, while staying legible and calm enough for a donation
flow that handles money and PII.

## Implementation Approach

**CSS custom properties are the source of truth**, referenced from
Tailwind config rather than duplicated into it:

```css
/* globals.css */
:root {
  --color-primary-500: #34A853;
  --radius-md: 0.75rem;
  /* ... */
}
```

```js
// tailwind.config.js
theme: {
  extend: {
    colors: {
      primary: {
        500: 'var(--color-primary-500)',
        /* ... */
      },
    },
    borderRadius: {
      md: 'var(--radius-md)',
    },
  },
}
```

Rationale: CSS variables are runtime-readable (useful for anything
that needs a raw value in JS — canvas, SVG, chart libraries) and keep
a single edit point if a token changes, while Tailwind utility classes
stay the day-to-day authoring interface so components don't need
inline `style` props. **Dark mode is explicitly out of scope for v1**
— this doc defines light-mode tokens only. If dark mode becomes a
demonstrated need later, the CSS-variable layer is exactly what makes
adding a `[data-theme="dark"]` override block cheap; it's not a
reason to build it now.

**Spacing scale**: Tailwind's default 4px-based scale (`0.5, 1, 1.5,
2, 3, 4, 6, 8...`) is used as-is, unmodified. No project-specific
override — introducing a custom spacing scale isn't justified without
a concrete gap the default doesn't cover (lowest-complexity
principle).

---

## Color Tokens

Every color is a 50–900 shade scale. Only the shades actually used by
components are listed per color below; the full scale exists in
`globals.css` for headroom (e.g. hover/active states one shade up or
down).

### Primary — Green (brand, primary CTA)

The main brand color — donate buttons, primary actions, active nav
state, links inside primary contexts. A slightly warm, yellow-leaning
green (not a cold/corporate blue-green) to match the "warm &
charitable" mood.

| Shade | Hex | Usage |
|---|---|---|
| 50 | `#F0FBF4` | Subtle background tint (e.g. selected campaign card) |
| 100 | `#DCF5E3` | Hover background for ghost/text buttons |
| 300 | `#8CDCA8` | Disabled-state fill |
| 500 | `#34A853` | **Base brand green** — icons, secondary emphasis |
| 600 | `#278A42` | **Primary button background** (passes AA against white text) |
| 700 | `#1F6E35` | Primary button hover/active |
| 900 | `#164825` | High-contrast text-on-light-green contexts |

### Success — Distinct green shade (semantic state)

Deliberately a different, more blue-leaning green from Primary — so
"this is the brand/CTA" and "this succeeded" never look identical.
Used for: donation success, verified organization, published campaign
status badges.

| Shade | Hex | Usage |
|---|---|---|
| 50 | `#ECFDF9` | Success banner/toast background |
| 500 | `#0F9D6E` | Success badge fill, success icon |
| 700 | `#0B7A56` | Success text on light background (AA compliant) |

### Warning — Orange (red-leaning, distinct from Accent)

Kept deliberately separate from Accent/Amber below — see the resolved
conflict note. Used for: `has_overdue_report` flag, near-deadline
notices, destructive-but-not-final confirmations.

| Shade | Hex | Usage |
|---|---|---|
| 50 | `#FFF4ED` | Warning banner background |
| 500 | `#E8590C` | Warning badge fill, warning icon |
| 700 | `#B8430A` | Warning text on light background |

### Error — Red

Standard destructive/error semantic. Used for: rejected curation
status, failed donation, form validation errors, destructive action
buttons (e.g. remove representative).

| Shade | Hex | Usage |
|---|---|---|
| 50 | `#FEF2F2` | Error banner/input background |
| 500 | `#DC2626` | Error badge fill, error icon, error text |
| 700 | `#B91C1C` | Destructive button hover/active |

### Info — Blue

Neutral informational semantic (distinct from both Primary green and
Accent amber, so it's unambiguous). Used for: informational banners,
`SecureUploadNote`, neutral tooltips.

| Shade | Hex | Usage |
|---|---|---|
| 50 | `#EFF6FF` | Info banner background |
| 500 | `#2563EB` | Info icon, info badge fill |
| 700 | `#1D4ED8` | Info text on light background |

### Accent — Amber (secondary/non-primary emphasis)

Bright, warm amber — used sparingly for secondary emphasis that isn't
a primary CTA and isn't a semantic state: highlight badges (e.g.
"Kurasi Baru" tag), secondary buttons, illustrative accents. **Not**
used for warning states — see conflict resolution below.

| Shade | Hex | Usage |
|---|---|---|
| 50 | `#FFFBEB` | Accent badge background |
| 400 | `#FBBF24` | Accent icon/illustration fill |
| 500 | `#F59E0B` | Secondary button background, accent badge border |
| 600 | `#D97706` | Secondary button hover/active |

**Conflict resolution note (Primary green vs Success green, Accent
amber vs Warning orange)**: both conflicts were flagged and resolved
the same way — by shifting hue/shade rather than reusing a color
across brand and semantic roles. This keeps "this is a call-to-action
or highlight" and "this is telling you the status of something"
visually distinct at a glance, which matters more here than
elsewhere since Kencleng's core flows (curation, disbursement,
fund-usage reporting) are status-heavy.

### Neutral — Cool Gray

| Shade | Hex | Usage |
|---|---|---|
| 50 | `#F8FAFC` | Page background |
| 100 | `#F1F5F9` | Card background (alternate), input background |
| 200 | `#E2E8F0` | Borders, dividers |
| 300 | `#CBD5E1` | Disabled borders |
| 400 | `#94A3B8` | Placeholder text, disabled text |
| 500 | `#64748B` | Secondary/muted body text |
| 700 | `#334155` | Body text |
| 900 | `#0F172A` | Heading text, high-emphasis text |

---

## Typography

### Font families

- **Heading**: Plus Jakarta Sans (weights 600, 700, 800) — used for
  all `h1`–`h4`, page titles, card titles, button labels.
- **Body**: Inter (weights 400, 500, 600) — used for body copy, form
  labels/inputs, table content, captions.

Both are loaded via `next/font/google` (self-hosted by Next.js at
build time, not a runtime Google Fonts request) — keeps the "no
external runtime dependency" property that matters for a PWA aiming
for good offline/App-Shell behavior, while still getting the two-font
pairing.

### Type scale

| Token | Font | Size / Line-height | Weight | Usage |
|---|---|---|---|---|
| `display` | Heading | 2.25rem / 2.5rem | 800 | Landing/hero only |
| `h1` | Heading | 1.875rem / 2.25rem | 700 | Page titles |
| `h2` | Heading | 1.5rem / 2rem | 700 | Section titles, card group headers |
| `h3` | Heading | 1.25rem / 1.75rem | 600 | Card titles |
| `h4` | Heading | 1.125rem / 1.5rem | 600 | Sub-section labels |
| `body-lg` | Body | 1.125rem / 1.75rem | 400 | Campaign narrative body |
| `body` | Body | 1rem / 1.5rem | 400 | Default UI text |
| `body-sm` | Body | 0.875rem / 1.25rem | 400 | Form labels, table cells, helper text |
| `caption` | Body | 0.75rem / 1rem | 500 | Timestamps, badge text, `MaskedField` masked value |

Button labels use `body` (400px context) or `body-sm` (compact
buttons) at weight 600 in the **body** font (Inter), not the heading
font — keeps button text feeling like an action, not a title.

---

## Shape & Elevation

### Border radius — "rounded jelas" (pronounced rounding)

| Token | Value | Usage |
|---|---|---|
| `radius-sm` | 8px | Badges, chips, small icon buttons |
| `radius-md` | 12px | Buttons, inputs, select/dropdown triggers |
| `radius-lg` | 16px | Cards (campaign card, dashboard panel) |
| `radius-xl` | 24px | Modals, auth overlay panel, large containers |
| `radius-full` | 9999px | Avatars, pill badges, circular icon buttons |

### Shadow — soft elevation

| Token | Value | Usage |
|---|---|---|
| `shadow-sm` | `0 1px 2px rgba(15, 23, 42, 0.06)` | Cards at rest |
| `shadow-md` | `0 4px 12px rgba(15, 23, 42, 0.08)` | Dropdowns, popovers, hover-elevated cards |
| `shadow-lg` | `0 12px 32px rgba(15, 23, 42, 0.12)` | Modals, auth overlay panel |

Neutral-900-based shadow color (not pure black) keeps shadows soft
and consistent with the cool-gray neutral palette rather than muddy.

---

## Icons

**Lucide** (`lucide-react`) — already the natural fit given
`shadcn/ui` primitives are already available in this environment's
React tooling, and consistent stroke-based style pairs well with the
rounded, friendly shape language above. Default stroke width 2px,
sized at `1rem`/`1.25rem`/`1.5rem` matching `caption`/`body`/`h4` text
contexts respectively so icons don't visually dominate adjacent text.

---

## Component Tokens

### Buttons

| Variant | Background | Text | Border | Usage |
|---|---|---|---|---|
| Primary | `primary-600`, hover `primary-700` | white | none | Donate, submit, main CTA — one per view |
| Secondary | `accent-500`, hover `accent-600` | `neutral-900` | none | Secondary emphasis action (not destructive, not primary) |
| Outline | transparent, hover `neutral-100` | `neutral-700` | `neutral-200` | Cancel, secondary navigation actions |
| Ghost | transparent, hover `primary-100` | `primary-700` | none | Low-emphasis inline actions (table row actions) |
| Destructive | `error-500`, hover `error-700` | white | none | Reject, remove representative, delete |

All buttons: `radius-md`, `body`/`body-sm` weight 600, `shadow-sm` on
Primary/Secondary/Destructive only (Outline/Ghost stay flat —
elevation implies "this is a filled, prominent action").

### Inputs

- Default: `neutral-100` background, `neutral-200` border, `radius-md`
- Focus: border → `primary-500`, plus a `2px` `primary-100` focus
  ring (offset, `focus-visible` only — keyboard-navigation
  accessibility, not on mouse click)
- Error: border → `error-500`, helper text in `error-700`,
  `body-sm`/`caption`
- Disabled: `neutral-50` background, `neutral-400` text, no border
  color change

### Badges (status indicators)

Every status enum across the app (`Organization`/`Campaign` curation
status, `Donation.status`, `has_overdue_report`, etc. — see
`kencleng-erd.md`) maps onto one of the five semantic colors, not a
new color per status:

| Status examples | Semantic color |
|---|---|
| `pending_verification`, `draft` | Neutral (`neutral-100` bg, `neutral-700` text) |
| `verified`, `published`, `success` (donation) | Success |
| `rejected`, `failed` (donation) | Error |
| `has_overdue_report = true`, near-deadline | Warning |
| Informational tags (e.g. "Baru") | Accent |

Badge shape: `radius-full` (pill), `caption` weight 500, `50`-shade
background + `700`-shade text for AA-compliant contrast at small text
size.

### Progress bar (donation progress)

Track: `neutral-200` background, `radius-full`. Fill: `primary-600`,
`radius-full`. This is the single most benchmark-sensitive component
(GoFundMe/Kitabisa pattern) — kept visually prominent (min height
`0.75rem`) since it's the primary trust/progress signal on public
campaign pages.

### `MaskedField`

Masked value rendered in `caption` size, `neutral-500` (muted, since
it's intentionally non-actionable content). Reveal toggle: `Ghost`
button variant, `Eye`/`EyeOff` Lucide icon only (no label text, to
keep it compact next to the masked value). Behavior spec (reveal
logging, persistence): `patterns.md` §C.

### `CurationDecisionPanel`

Approve action uses the Primary button style; Reject uses Destructive.
The mandatory `decision_note` textarea (shown on reject) uses default
Input styling with an Error-toned helper text noting it's required —
reusing existing tokens rather than introducing panel-specific styles.
Behavior spec: `patterns.md` §C and Pattern 5.

### `SecureUploadNote`

Rendered as a small inline banner: Info-50 background, Info-700 text,
`radius-sm`, `Lock` or `ShieldCheck` Lucide icon at `caption` text
size scale. Behavior spec: `patterns.md` §C.

---

## Accessibility

- **Minimum contrast: WCAG AA (4.5:1 for body text, 3:1 for large
  text/icons)**, applied when choosing which shade of each color pairs
  with which text color above — e.g. `primary-600` (not `500`) is the
  button background specifically because `500` fails AA against white
  text at normal button-label size.
- **Focus-visible only** ring styling (not on mouse click) — avoids
  the common AI-generated-code pitfall of either missing focus states
  entirely or showing them on every click, which is noisy for mouse
  users and unhelpful for keyboard users if inconsistent.
- Badge and status colors are always paired with a text label (never
  color alone) — relevant given a chunk of Kencleng's status
  vocabulary (`pending_verification` vs `rejected` vs `verified`) needs
  to be distinguishable for color-blind users too.

---

## Related Docs

- Page inventory: `page-map.md`
- Page patterns & shared component behavior: `patterns.md`
- Code architecture: `kencleng-frontend-tech-stack.md`
- Status enums referenced in badge mapping: `kencleng-erd.md`
- Roadmap tracking: `kencleng-roadmap-next-steps.md` (Step 10)
