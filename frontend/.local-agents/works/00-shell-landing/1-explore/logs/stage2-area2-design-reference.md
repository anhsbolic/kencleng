# Stage 2 — Area 2: Design reference export

> Task: 00-shell-landing
> Date: 2026-08-26

## Current state

- Path inconsistency: `AGENTS.md` §3 and `prototype-reference.md` both say
  the export lives at `design-reference/` at repo root; it's actually at
  `docs/design-reference/*.html` (10 files, matches task's cited path).
  `docs/design-reference/README.md` independently confirms the read-only/
  frozen-reference rule, so fencing intent is unaffected — path is just
  stale in two docs.
- `design-reference-usage.md`'s extraction script only pulls the inline
  `<script type="text/babel">`. For `landing-page.html` that inline script
  is just two `ReactDOM.createRoot(...).render(...)` calls — the real
  207-line `Landing` component is loaded via
  `<script type="text/babel" src="01568b12-...">`, a gzip+base64 blob keyed
  by that ID inside `<script type="__bundler/manifest">`. The documented
  script does NOT surface this — needs an extra base64→gunzip decode step
  against the manifest. Recovered manually; naive extraction silently
  produces a 171-char near-empty JSX file with no error.
- Recovered `Landing` component structure: `Header`(mobile, hamburger +
  drawer)/`DesktopHeader` → `Hero` → `TrustStrip` (3 stats) → `Featured`
  (id=`kampanye`, `CampaignBrowseCard` × 3) → `HowItWorks` (id=`cara-kerja`,
  3-step explainer) → `Footer`.
- **Known Issue #2 confirmed + root-caused:** `CampaignBrowseCard` renders
  `<image-slot id={...} shape="rect" placeholder="Foto kampanye">` — a
  custom element whose runtime rendering (found in the shared design-
  system bundle, `c56486b2-...`, shared across all 10 prototype exports)
  includes "browse files" text + upload-affordance styling.
- **Known Issue #3 confirmed with exact numbers:** extracted CSS has
  `--font-size-h1: 44px`, `--font-size-display: 40px` — matches
  prototype-reference.md exactly. Colors/radius/shadow spot-checked, all
  match `design-guidelines.md` exactly.
- **New nuance beyond Known Issue #3:** Hero's `<h1>` doesn't consume
  either CSS var — it hardcodes `fontSize: m ? 30 : 48` inline. Desktop
  value (48) matches neither the drifted prototype tokens nor
  design-guidelines.md's tokens — a third, larger, hardcoded number.
  Mobile value (30) happens to equal design-guidelines.md's `h1` token,
  though `display` (36px) is specced as "Landing/hero only" — raises which
  token should govern the Hero heading (Stage 3).
- `CAMPAIGNS` mock data (`org`, `title`, `pct`, `days`, `slot`) is
  illustrative only — `pct`/`days` arrive pre-computed, not raw amounts/
  dates.
- Every card hardcodes `Badge tone="success"` "verified" with no
  underlying status field (mock has none) — illustrative styling only.

## Requirement

Per task args + `ai-prototype-to-production-translation.md`: carry over
composition/states/copy, translate inline styles → Tailwind, translate
scratch primitives (`Button`/`Badge`/`ProgressBar`/`Icon`) → real
`components/ui/`, translate mock data → real API contract, swap Known
Issues #2/#3 for correct behavior/tokens.

## Gap

1. Composition transfers cleanly: Header/DesktopHeader (see Misleading
   Signal below), Hero, TrustStrip, Featured/CampaignBrowseCard,
   HowItWorks, Footer.
2. `HowItWorks` and `TrustStrip` aren't mentioned anywhere in `page-map.md`
   or `patterns.md` — page-map.md's only resolved decision for `/` is
   "Browse highlighted campaigns." Whether to keep either is unresolved
   (Stage 3).
3. Real `components/ui/` primitives needed (`Badge`, `ProgressBar`, plain
   image placeholder) don't exist yet.
4. Real campaign data mapping: `pct`/`days` must derive from actual API
   fields; "verified" badge must read a real `Organization.status`, not
   render unconditionally.

## Sniffing

- **Misleading signal:** the `m` boolean threaded through every component
  looks like it "already handles responsive," but it's Claude Design's own
  side-by-side preview convention — the file mounts two separate React
  roots (`#desktop` 1280px, `#mobile` 390px canvas frames) with `m` fixed
  per root. No CSS media query or runtime viewport detection anywhere.
  Not a responsive mechanism to port — real page needs an actual CSS-
  breakpoint switch, consistent with how the Auth Shell's modal/full-page
  split was built (`phase0-shared-infra.md` Step 3: "CSS, not a JS
  resize-listener").
- **Risk:** hardcoded "verified" badge, if copied as unconditional markup
  rather than reading real org-status data, would show a false trust
  signal on any non-verified org's campaign — contradicts `patterns.md`'s
  "Organization trust signal" requirement.
- **Inconsistency:** repo-root `design-reference/` path in two docs no
  longer matches actual `docs/design-reference/` location.
- **Edge case:** the extraction gap (see Current state) means a future
  task extracting a different `docs/design-reference/*.html` file could
  silently get a near-empty JSX and mistake "component is trivial" for
  "extraction missed the payload" — worth flagging to whoever maintains
  `design-reference-usage.md`.
