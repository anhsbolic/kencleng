# Stage 1 — Plan Announcement

> Task: 00-shell-landing (landing page `/` — public shell nav + highlighted campaigns mock)
> Date: 2026-08-26

## Areas to explore

### Area 1: Requirement docs
`docs/ui-ux/page-map.md` (2026-08-24 decisions for `/`), `docs/ui-ux/
prototype-reference.md`'s Tier 1 entry for `/`, `docs/ui-ux/patterns.md`,
`docs/ui-ux/design-guidelines.md`. First — everything else gets evaluated
against what these actually specify, not assumed.

### Area 2: Design reference export
`docs/ui-ux/design-reference-usage.md` (how to read the export) then the
actual `docs/design-reference/landing-page.html`. Second — need the
extraction method before the export is readable, and need the two known
issues confirmed before trusting anything else in the file.

### Area 3: Current frontend codebase state
`app/` router structure, `components/ui/` inventory, existing shell/nav
components, `lib/hooks/`, `mocks/`. Third — tells us what's scaffolded
already (repo is at phase-0 scaffold) vs. what this task builds from
scratch.

### Area 4: Campaign API contract
`lib/api/` generated types + `api/openapi/campaign.yaml`. Last — determines
the real data shape the mock-to-real translation must target, depends on
knowing (Area 3) whether `lib/api/` is generated yet.

## Order rationale

Dependency-ordered: spec docs first (what's required), then the visual
reference (needs the spec's known-issues list to be trustworthy), then
current code state, then the API contract (needs to know what's already
scaffolded before checking what's missing).
