# Kencleng — UI/UX Prototype Reference

> Intended path: `docs/ui-ux/prototype-reference.md`
> Status: New (2026-08-21) — indexes the Claude Design prototypes
> exported into `design-reference/` at repo root, and defines how
> much authority each one carries for implementation.
> Last updated: 2026-08-21

## Context

`patterns.md` and `design-guidelines.md` are the durable spec —
abstract shape + tokens. To validate that spec actually looks right
and to give the coding agent a concrete visual precedent, a set of
representative pages were prototyped in Claude Design and exported as
real code into `design-reference/` (see `AGENTS.md` §3 — that
directory is read-only for agents, frozen reference output, not a
build target).

Not every page got its own prototype — doing that for all ~40 routes
in `page-map.md` would recreate the exact staleness/maintenance
problem that got the old wireframes retired. Instead, this doc
defines two tiers so an agent (or a human reviewing agent output)
knows how much weight to give the reference for any given page.

## Tier 1 — Near-final draft (build from this directly)

These routes have an actual generated prototype in
`design-reference/`. For these specific pages, treat the prototype as
close to authoritative on layout/visual detail — implement to match
it, deviating only where the feature spec (`docs/spec/<domain>/
features/*.md`) requires different behavior than what was mocked
(e.g. the prototype used illustrative placeholder data; real data
shape follows the OpenAPI contract, not the mock).

| Route | Pattern (`patterns.md`) | Notes |
|---|---|---|
| `/` | Landing (one-off, not a formal pattern) | Public header variant, not Dashboard Shell |
| `/login` | Form (Auth sub-variant) | Desktop = modal overlay, mobile = full page — see known issue below |
| `/campaign` | List/Browse | Includes loading + empty states |
| `/campaign/[id]` | Detail (public variant) | Includes loading state; most benchmark-sensitive page |
| `/campaign/[id]/donate` | Form (single-step) | Includes idle + submitting states |
| `/dashboard/campaign/new` | Form (Revisable Submission) | Includes idle, field-error, submitting states |
| `/dashboard/campaign/[id]/monitor` | Dashboard/Summary | Includes independent per-card loading + partial-failure states |
| `/dashboard/kurasi/campaign/[assignmentId]` | Curation/Review | Includes idle, reject-expanded, locked/already-decided states |
| `/donation/[id]/status` | Status/Tracking | Includes loading, success, invalid-token states |
| `/dashboard/organization/new` | Form (Revisable Submission) | Includes idle + field-error states; the one page that exercises `SecureUploadNote` in real context |

Plus two non-page reference sheets (not routes, pure component/layout
precedent): the original **Component & Layout sheet** — buttons,
badges, inputs, `MaskedField`, `SecureUploadNote`, progress bar,
Dashboard Shell desktop/mobile.

## Tier 2 — Template only (derive from `patterns.md`)

Every other route in `page-map.md` has no prototype. For these, apply
`patterns.md`'s definition for that pattern, using the closest Tier 1
example below as the structural/visual precedent — not a literal copy
(different data, different role-gating), but same shape and token
usage.

| Pattern | Tier 2 routes (no prototype) | Closest Tier 1 precedent |
|---|---|---|
| List/Browse | `/dashboard/admin/users`, `/dashboard/admin/kurasi-queue`, `/dashboard/kurasi` (queue), `/dashboard/donations`, `/dashboard/donations/claim`, `/dashboard/notifications`, `/dashboard/organization/[id]/representatives` | `/campaign` |
| Detail | `/dashboard/organization/[id]`, `/dashboard/campaign/[id]`, `/dashboard/campaign/[id]/report`, `/dashboard/campaign/[id]/disbursement/[reqId]`, `/dashboard/campaign/[id]/fund-usage-report/[reportId]` | `/campaign/[id]` (adjust: dashboard context, role-gated sections) |
| Form | `/register`, `/forgot-password`, `/reset-password`, `/dashboard/profile`, `/dashboard/security`, `/dashboard/event/new`, `/dashboard/campaign/[id]/edit`, `/dashboard/campaign/[id]/publish`, `/dashboard/campaign/[id]/disbursement/new`, `/dashboard/campaign/[id]/fund-usage-report/new` | `/dashboard/campaign/new` (dashboard forms) or `/login` (auth modal/mobile split) |
| Curation/Review | `/dashboard/kurasi/organization/[assignmentId]`, `/dashboard/kurasi/fund-usage/[assignmentId]`, `/dashboard/admin/campaign/[id]/force-close`, `/dashboard/admin/disbursement/[reqId]` | `/dashboard/kurasi/campaign/[assignmentId]` |
| Status/Tracking | (none — only one instance of this pattern exists) | `/donation/[id]/status` |

## Known Issues — do NOT carry over into implementation

These were found during prototyping and are **not** correct per
`patterns.md`/`design-guidelines.md` — an agent implementing the real
page should follow the spec docs, not replicate these:

1. **Login error state, `/login`**: the prototype rendered the
   generic "Email atau password salah" authentication failure as a
   field-level error attached to the Email input (red border, message
   under that field) instead of a banner above the form. Per
   `patterns.md` §B, request-level failures (this one) and
   field-level validation errors must not be conflated, and a generic
   auth failure must not visually implicate one specific field (that
   leaks which part of the attempt was correct). **Status: not
   confirmed fixed in the final export — verify before implementing.**
2. **Campaign card image placeholder, `/campaign` and `/` featured
   section**: the placeholder for a missing campaign photo rendered as
   a file-upload dropzone ("Foto kampanye — atau *browse files*")
   instead of a plain read-only image placeholder. These are public,
   read-only cards — no upload affordance belongs there (upload only
   belongs on `/dashboard/campaign/new`'s actual media field).
   **Status: known issue, not fixed — cosmetic only, but don't copy
   the dropzone affordance into the real read-only card component.**
3. **Typography sizes drift slightly from `design-guidelines.md`**:
   e.g. exported `--font-size-h1` is 44px vs the spec's 30px,
   `--font-size-display` is 40px vs spec's 36px. Colors, radius, and
   shadow tokens verified exact; type scale did not fully carry over.
   **Status: known issue, not fixed — `design-guidelines.md` remains
   the literal source of truth for exact type sizes; don't copy the
   prototype's font-size values verbatim into `frontend/`.**

## Related Docs

- Pattern definitions: `patterns.md`
- Visual tokens: `design-guidelines.md`
- Full page inventory: `page-map.md`
- Directory boundary rule for `design-reference/`: `AGENTS.md` §3