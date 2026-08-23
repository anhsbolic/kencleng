# Kencleng — Using `design-reference/` for Frontend Development

> Intended path: `docs/ui-ux/design-reference-usage.md`
> Status: New (2026-08-21)
> Last updated: 2026-08-21

## What these files actually are

Each file in `design-reference/` is a Claude Design "standalone HTML"
export. It is **not** plain static HTML/CSS — it's a self-bootstrapping
bundle. The real content (design tokens as CSS custom properties, and
the actual page as uncompiled React/JSX source) is packed inside a
`<script type="__bundler/template">` tag as a JSON-escaped string, which
a loader script unpacks and renders client-side on page load.

This matters practically: **do not `cat`/grep the raw `.html` file
expecting to read clean markup** — you'll mostly see escaped string
noise and opaque font/script IDs. Extract first (see below), then read
the extracted files normally.

## Step 1 — Extract before reading

Use this script (tested against a real export) to pull out clean,
readable files:

```javascript
// scripts/extract-design-reference.mjs
import { readFileSync, writeFileSync } from 'node:fs';
import path from 'node:path';

const inputPath = process.argv[2];
if (!inputPath) {
  console.error('Usage: node extract-design-reference.mjs <standalone.html>');
  process.exit(1);
}

const html = readFileSync(inputPath, 'utf-8');

function extractScriptJSON(type) {
  const re = new RegExp(`<script type="${type}">([\\s\\S]*?)</script>`);
  const m = html.match(re);
  if (!m) throw new Error(`Could not find <script type="${type}">`);
  return JSON.parse(m[1]);
}

const template = extractScriptJSON('__bundler/template'); // full inner HTML doc, as a string

const jsxMatch = template.match(/<script type="text\/babel">([\s\S]*?)<\/script>/);
if (!jsxMatch) throw new Error('No JSX (text/babel script) found in template');

const base = path.basename(inputPath, path.extname(inputPath));
writeFileSync(`${base}.extracted.html`, template, 'utf-8');
writeFileSync(`${base}.extracted.jsx`, jsxMatch[1].trim(), 'utf-8');

console.log(`Wrote ${base}.extracted.html (full doc: CSS tokens + markup)`);
console.log(`Wrote ${base}.extracted.jsx (readable component source)`);
```

Run it once per file you need to consult:
```
node scripts/extract-design-reference.mjs design-reference/Kencleng_Campaign_Detail__standalone_.html
```

This produces two sibling files: `*.extracted.html` (CSS custom
property tokens + full markup, useful for seeing exact spacing/layout
values) and `*.extracted.jsx` (the component source, useful for seeing
structure/composition/states). These extracted files are scratch
output for reading — don't commit them, they're regenerable from the
originals in `design-reference/` any time.

## Step 2 — What to take from the extracted JSX, and what NOT to

The extracted JSX is real, working React code — but it comes from
Claude Design's own scratch environment, not from `kencleng-frontend-
tech-stack.md`'s architecture. Some of it transfers directly; some of
it must be translated, not copied.

### Take directly (structural/behavioral precedent)
- **Component decomposition** — e.g. the donation form's breakdown
  into `AmountField`, `MethodGrid`, `AnonymousToggle`, `SummaryStrip`
  is a reasonable shape to mirror when building the real
  `components/features/donation/` components.
- **Which states exist and how they differ** — e.g. `disabled` prop
  threading through every field when `busy` (submitting), the
  quick-select chip pattern, the payment-method grid's selected-state
  styling logic.
- **Copy/microcopy** — Indonesian field labels, helper text, button
  text — these are legitimate content to reuse or adapt.
- **Accessibility patterns** — e.g. `aria-pressed` on the payment
  method buttons.

### Translate, don't copy verbatim
- **Inline `style={{...}}` with `var(--token)` references → Tailwind
  utility classes.** The reference code styles everything with inline
  style objects (e.g. `background: "var(--color-primary-600)"`,
  `borderRadius: "var(--radius-md)"`). The real app uses Tailwind
  utility classes mapped to the same CSS variables via
  `tailwind.config.js` (per `design-guidelines.md` — Implementation
  Approach). Translate `style={{ background: "var(--color-primary-
  600)", borderRadius: "var(--radius-md)" }}` into `className="bg-
  primary-600 rounded-md"`, not a copy-pasted inline style object.
- **Hardcoded pixel values in spacing/gaps** (e.g. `gap: 8`, `padding:
  "12px 14px"`) — the reference doesn't consistently use its own
  `--space-N` tokens for every gap. Map these to the nearest Tailwind
  spacing scale value instead of reproducing the exact px number.
- **Font sizes** — per the known issue in `prototype-reference.md`,
  the reference's type scale drifted from `design-guidelines.md`.
  Follow `design-guidelines.md`'s type scale table as the literal
  source of truth, not the extracted CSS.
- **The `window.KenclengDesignSystem_82607b` primitives** (`Button`,
  `Badge`, `Icon`, `ProgressBar` imported at the top of the JSX) —
  these are Claude Design's own scratch component library, not the
  real Kencleng component library. Build the real versions in
  `components/ui/` from `design-guidelines.md`'s Component Tokens
  section; use the reference only to see how they're expected to
  compose and behave.
- **Mock data** (e.g. the `CHIPS`/`METHODS` arrays, placeholder
  campaign data) — illustrative only. Real data shape comes from
  `api/openapi.yaml` and the relevant feature spec, never from what's
  hardcoded in the reference.
- **State management** (`React.useState` scattered inline) —
  illustrative only. Real state follows `kencleng-frontend-tech-
  stack.md`: TanStack Query for server state, Zustand for client
  state, `react-hook-form` + `zod` for forms.

## Step 3 — Cross-check against known issues first

Before treating anything in a `design-reference/` file as correct,
check `prototype-reference.md`'s "Known Issues" section for that
route — a few things were confirmed wrong during prototyping (e.g.
the login error-banner placement, the campaign card upload-style
placeholder) and must not be carried into the real implementation
even though they're sitting right there in the reference code.

## Related Docs

- Tier 1/Tier 2 route mapping and known issues: `prototype-
  reference.md`
- Visual tokens (the literal source of truth): `design-guidelines.md`
- Page shapes and states: `patterns.md`
- Code architecture `design-reference/` code must be translated into:
  `kencleng-frontend-tech-stack.md`
- Directory boundary rule (`design-reference/` is read-only for
  agents): `AGENTS.md` §3