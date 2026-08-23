# Kencleng — Frontend UX Patterns

> Intended path: `docs/ui-ux/patterns.md`
> Status: New (2026-08-20) — replaces per-page gray-box wireframes
> (`docs/wireframes/`, retired) with reusable page patterns + shared
> component behavior. Decision rationale: wireframes drifted from the
> actual domain spec within a month of being drawn (pre-dated the
> 2026-08-20 spec-first pass); a pattern-level doc is cheaper to keep
> in sync and matches how the backend spec already separates
> *behavior precision* from *pixel precision*.
> Last updated: 2026-08-20

## Context

This doc sits alongside two other `docs/ui-ux/` files:

- `design-guidelines.md` — the visual layer (color, type, shape,
  elevation, component tokens): "what it looks like"
- `page-map.md` — per-persona, per-route inventory: "which page,
  which pattern, which endpoint"
- **This doc** — the structural layer in between: "what shape does
  this kind of page take, and what states does it need to handle"

`page-map.md` references this doc's pattern names instead of
re-describing layout per page. `design-guidelines.md`'s component
tokens (buttons, badges, inputs) are the visual vocabulary these
patterns are built from.

---

## A. Page Patterns

### 1. List / Browse Page

Header (title + optional search/filter) → grid or list of item cards
→ pagination (or infinite scroll — not yet decided, default to
pagination for v1, lowest complexity).

| State | Behavior |
|---|---|
| Loading | Skeleton cards matching the real card layout (not a spinner) — count matches typical page size |
| Empty | Icon + short message. Primary action shown only if the viewer can create the item (e.g. Owner sees "Buat campaign baru" CTA; public Guest browsing sees no CTA) |
| Error | Inline banner + retry button, generic message only — never surface raw backend error text (security-by-design: avoids leaking implementation detail) |
| Success | Cards render, pagination controls active only if more than one page |

Used by: `/campaign` (public list), `/dashboard/admin/users` (search
result list), `/dashboard/kurasi` (unified queue — list variant with
type tags).

### 2. Detail Page

Header (title/status) → primary content → contextual actions (varies
by viewer role/ownership).

Two variants:
- **Public detail** — no auth-gated content, e.g. `/campaign/[id]`
- **Dashboard detail** — role-gated sections/actions within the same
  page, e.g. `/dashboard/organization/[id]` (legal docs section
  Owner-only, per Business Rule 4)

**Long-text sections** (campaign narrative, `beneficiary_description`,
`report_narrative`): collapse past a fixed height with a "Baca
selengkapnya"/"Read more" expand toggle rather than always rendering
in full — keeps the page scannable when the organizer writes a long
narrative. Applies wherever these fields are shown, not just
`/campaign/[id]`.

**Organization trust signal**: any page displaying campaign or
organization info to a public/donor viewer shows the owning
organization's verification badge (`Organization.status = verified`
→ Success-colored badge per `design-guidelines.md`) alongside the
org name — this is data already returned by the API, just making it
an explicit required element of the Detail pattern rather than
optional. Reinforces trust at the point donors decide whether to
give, consistent with the "warm & charitable but still handles money"
brand direction.

| State | Behavior |
|---|---|
| Loading | Skeleton matching section layout |
| Not found | Dedicated not-found message, not a generic error — distinct from a transient fetch error |
| Error | Retry banner, same generic-message rule as List pattern |
| Success | Full content + actions gated per role |

Used by: `/campaign/[id]`, `/dashboard/organization/[id]`,
`/dashboard/campaign/[id]`, `/dashboard/campaign/[id]/monitor`.

### 3. Form Page (Create / Edit)

Form sections → inline field validation → submit/cancel actions.

| State | Behavior |
|---|---|
| Idle | Fields empty (create) or pre-filled (edit) |
| Validating | Inline, on blur + on submit — `zod` schema errors mapped to field-level messages, `error-700` per `design-guidelines.md` |
| Submitting | Submit button disabled + inline spinner, rest of form disabled to prevent double-submit |
| Submit error | Banner above form for request-level failure (network/5xx); field-level errors for validation (422) — never conflate the two |
| Success | Toast + redirect (dashboard forms) or inline success state (guest-facing forms — see Status/Tracking pattern) |

**Sub-pattern — Revisable Submission**: draft → submit → locked →
revise-if-rejected → resubmit. Same Form Page pattern, plus a
read-only "locked" state while under curation and a "why rejected"
banner (`decision_note` from the reviewer) when revising. Used by:
organization registration, campaign registration, fund-usage-report
submission, disbursement request.

Used by: `/dashboard/organization/new`, `/dashboard/campaign/new`,
`/dashboard/campaign/[id]/edit`, `/register`, `/login`,
`/dashboard/campaign/[id]/fund-usage-report/new`, and others per
`page-map.md`.

### 4. Dashboard / Summary Page

Dashboard Shell (top-nav desktop / top-bar+hamburger mobile, per
`kencleng-frontend-tech-stack.md`) → stat cards / summary panels.

| State | Behavior |
|---|---|
| Loading | Skeleton per card — cards load independently, not blocked on each other |
| Empty | Per-card empty state (e.g. "Belum ada donasi masuk"), not a full-page empty state — dashboard shell itself is never "empty" |
| Error | **Partial failure allowed** — one card's fetch failing shows that card's error state without blocking the rest of the dashboard |
| Success | All cards populated |

Used by: `/dashboard/campaign/[id]/monitor`, `/dashboard/donations`.

**Reference shape [NEW — validated via Claude Design prototype,
2026-08-21]**: for an Owner/Staff landing dashboard specifically, 3
stat cards side by side (e.g. "Total terkumpul", "Donatur", "Laporan"
with a pending-count sub-label) followed by a list of active campaign
progress bars below — confirmed this reads well at both desktop and
mobile widths (stat cards stack to a 2-up or 1-up grid on mobile
rather than staying 3-across).

### 5. Curation / Review Page

Read-only item detail (what's being reviewed) → `CurationDecisionPanel`
(approve / reject + mandatory `decision_note` on reject).

| State | Behavior |
|---|---|
| Loading | Skeleton for item detail |
| Already decided | Locked, view-only — panel shows the recorded decision instead of action buttons (prevents double-decision) |
| Submitting decision | Panel buttons disabled + spinner during the approve/reject call |
| Error | Banner, decision not applied, panel stays interactive for retry |

Same shape across all three curation contexts (organization curation,
campaign curation, fund-usage-report verification) — parameterize
`CurationDecisionPanel` by curation type rather than building three
separate panels.

Used by: `/dashboard/kurasi/organization/[assignmentId]`,
`/dashboard/kurasi/campaign/[assignmentId]`,
`/dashboard/kurasi/fund-usage/[assignmentId]`.

### 6. Status / Tracking Page (unauthenticated)

Minimal shell (no Dashboard Shell — guest has no session) → single
status card, resolved via token-in-URL lookup.

| State | Behavior |
|---|---|
| Loading | Skeleton card |
| Invalid/missing token | Distinct message from "not found" — token malformed vs donation genuinely doesn't exist are different failure modes, but **do not** let the message distinguish them to an attacker (avoid confirming/denying donation existence) — generic "not found or link invalid" for both |
| Success | Status badge (`pending`/`success`/`failed`, per `design-guidelines.md` badge mapping) + relevant detail |

Used by: `/donation/[id]/status?token=...`.

---

## B. Cross-Pattern State Conventions

These apply regardless of which pattern above a page uses:

- **Loading**: skeleton shaped like the real content, not a bare
  spinner — except inline button-level loading (submit actions),
  where a small inline spinner is correct.
- **Empty**: icon + one short line of copy. Primary action shown only
  if the current viewer is actually allowed to take it — don't show
  a disabled/misleading CTA to a viewer without permission.
- **Error (page/section-level)**: generic message + retry action.
  Never render raw backend error text or stack traces — consistent
  with the project's secure-by-design goal; this is a common
  AI-generated-code pitfall (leaking implementation detail through
  error messages) worth avoiding deliberately.
- **Error (field-level)**: inline, tied to the field via `aria-*`,
  `error-700` text per `design-guidelines.md`.
- **Success**: toast for actions taken from within an existing
  dashboard context (nothing to navigate away from); full success
  state/redirect for terminal actions in a flow (e.g. donation
  submitted, registration submitted) where the user needs a clear
  "what happens next."
- **Stale/offline data [NEW — 2026-08-20]**: per the PWA scope
  decision (`kencleng-frontend-tech-stack.md`, "PWA Scope" — app-shell
  caching only, no offline write queue), a page can render with
  cached-but-stale data when the network is unavailable. Any page
  showing TanStack Query data should surface a small "data mungkin
  tidak terbaru" indicator when `isStale`/offline is detected, rather
  than silently showing old numbers as if they were current — matters
  most for money/status figures (donation progress, curation status).
  Distinct from the Error state: this isn't a failure, just a
  freshness caveat.

---

## C. Shared Components

> Moved from `kencleng-frontend-tech-stack.md` (2026-08-20) — these
> are UX/behavior specs, not code-architecture decisions, so they
> live here instead. `frontend-tech-stack.md` keeps only the
> structural/code-organization note that `components/shared/` is
> where these live in the folder structure.

### `MaskedField`

Central component for all PII masking, per `kencleng-actors-entities.md`
PII Handling Note. Used wherever `guest_email`, `User.primary_email`,
`NPWP`, or future banking details are displayed — **regardless of
viewer role, including Admin**.

- Default: masked display (e.g. `j***@***.com`)
- Explicit reveal toggle per field instance (not a global "show all
  PII" switch — intentional friction)
- On reveal by Admin or Kurator viewing another party's data: fires a
  call to log the reveal action to Audit Log (`kencleng-phase0-detail.md`
  Feature 9) — `MaskedField` needs to know viewing context (whose
  data, what field, who's viewing), not just be a dumb visual toggle
- **Reveal persistence [RESOLVED]**: stays revealed until manual
  re-toggle or page refresh/navigation. Implemented as plain local
  component state (`useState`) — no timeout/auto-re-mask logic.
  Navigating away unmounts the component and resets it for free.

### `SecureUploadNote`

Small reassurance banner used on every non-public file upload form
(organization legal docs, fund-usage-report attachments). Purely
informational, no logic. Visual spec: Info-toned banner per
`design-guidelines.md`.

### `CurationDecisionPanel`

Reused across all 3 curation contexts — approve/reject buttons +
mandatory `decision_note` textarea on reject. See Pattern 5 above for
full state behavior. One component parameterized by curation type,
not three separate implementations.

---

## Related Docs

- Visual tokens: `design-guidelines.md`
- Per-persona page inventory: `page-map.md`
- Code architecture (folder structure, state management, testing):
  `kencleng-frontend-tech-stack.md`
- Status enums referenced in badge states: `kencleng-erd.md`