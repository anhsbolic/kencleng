# Stage 2 — Area 4: Campaign API contract

> Task: 00-shell-landing
> Date: 2026-08-26

## Current state

- `campaign.yaml`'s `GET /campaigns` (public, `security: []`) accepts only
  `category`, `q` (free-text, marked `# INFERRED`), `cursor`, `limit` —
  confirms page-map.md's footnote exactly: no featured/sort param exists.
- Response: `CampaignListResponse { data: CampaignListItem[], pagination:
  { next_cursor, has_more } }` — cursor-based, not page-number.
- `CampaignListItem = Campaign & { progress?: CampaignProgress }`.
  `Campaign` has `organization_id` (bare UUID) but **no organization
  name, no verification status, no media/image field of any kind**.
  `CampaignProgress` = `{ percentage (capped at 100), donor_count,
  days_remaining (nullable once closed) }`.
- Only `CampaignDetail` (`GET /campaigns/{id}`) additionally embeds
  `organization: OrganizationSummary { id, name, status }`.
  `docs/spec/4-campaign/features/02-campaign-detail-listing.md` confirms
  this is deliberate, spec-authored: the composite shape exists
  specifically "to avoid client-side N+1 on the platform's primary
  conversion surface" (the detail page) — never extended to the list
  endpoint.
- `schema.d.ts` (already generated, newer than source YAMLs) already
  contains `CampaignListItem`/`CampaignListResponse`/`CampaignProgress`/
  `OrganizationSummary` types — no codegen re-run needed.
- No image/media field exists anywhere on `Campaign`/`CampaignListItem`/
  `CampaignDetail` today — the design reference's `image-slot` photo is
  entirely unbacked by the contract, independent of the upload-dropzone
  known issue.
- `Pagination` is cursor-based (`next_cursor`/`has_more`) — consistent
  with a "fixed fixture, no pagination controls" mock for `/`'s
  highlighted section.

## Requirement

The design reference's `CampaignBrowseCard` shows org name + a "verified"
badge on every card. `patterns.md`'s Detail pattern requires the org
verification badge "any page displaying campaign or organization info to
a public/donor viewer" (ambiguous whether this covers home-page preview
cards — see Area 1).

## Gap

Most consequential gap across all areas: **`GET /campaigns` — the exact
endpoint page-map.md names as backing `/`'s highlighted-campaigns section
— structurally cannot supply organization name or verification status**
per the current, spec-authored contract. Only `organization_id` (UUID) is
available per list item. Building the card exactly as the Tier 1
prototype shows it isn't achievable against real data without one of:
dropping those fields from the home card, a backend contract change, or
a client-side N+1 fetch per card. Not this task's call to make silently
(Stage 3).

Also: no image field exists on any campaign schema — the plain
placeholder isn't a stand-in for a real photo URL the mock could return;
there's nothing to wire it to yet.

## Sniffing

- **Inconsistency (headline finding):** the Tier 1 visual reference and
  the backend-authored API spec disagree on what data a browse card can
  show, and neither `page-map.md`, `prototype-reference.md`, nor
  `design-reference-usage.md` flags this — only surfaces by reading
  `campaign.yaml`'s schema directly against the extracted JSX side by
  side.
- **Miscontext:** `patterns.md`'s "Organization trust signal" note was
  written with Detail-pattern pages as its only "Used by" examples —
  applying it to `/`'s List/Browse-shaped cards may be reading a Detail-
  pattern requirement into a page pattern it wasn't written for.
- **Edge case:** `CampaignProgress.days_remaining` is nullable ("null once
  closed"), but `/campaigns` only ever returns `status = 'published'`
  campaigns per spec — in practice a `/`-highlighted card shouldn't
  observe null, but a defensive display fallback is still cheap
  insurance against a malformed fixture.
- **Risk:** `percentage` arrives pre-computed and capped at 100 server-
  side — frontend must display as-is (AGENTS.md §2: never compute
  totals/eligibility client-side), not recompute
  `collected_amount / target_amount` the way the mock's `pct` field
  suggests, even though both coincidentally match today.

## Cross-area summary — standout findings for Stage 3

1. Org name/verified badge on the highlighted-campaign card has no
   backing data in `GET /campaigns`'s real response shape — a genuine,
   spec-authored contract gap between the Tier 1 prototype and the actual
   API, not a frontend oversight.
2. Typography-scale tokens don't exist anywhere in the codebase yet —
   this task is the first to need the full scale, not just a two-number
   fix to Known Issue #3.

Also carried to Stage 3: `HowItWorks`/`TrustStrip` exist in the prototype
but aren't in `page-map.md`/`patterns.md`; the `m`-boolean pattern is a
design-tool preview artifact, not a responsive mechanism to port; the
repo-root `design-reference/` path in `AGENTS.md`/`prototype-reference.md`
is stale (actual: `docs/design-reference/`).
