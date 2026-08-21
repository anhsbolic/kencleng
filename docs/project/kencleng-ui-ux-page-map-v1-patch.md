# Kencleng — UX Page Map (Persona × Phase × Page)

> Intended path: `docs/ui-ux/page-map.md`
> Status: Revised 2026-08-20 — evolved from the original
> `kencleng-ux-page-map.md` (2026-07-24). Wireframe references
> (`docs/wireframes/`) removed — see `patterns.md` for why. Each page
> row now names the reusable pattern it uses instead of a per-page
> layout description.
> Last updated: 2026-08-20

## Context

This doc maps **which pages exist and what each persona can do on
them**, cutting across Phase 0-3. It sits at a different layer than
the other frontend docs:

- `kencleng-phase0-3-detail.md` — business rules & flow per feature
  (what happens, in what order, what data)
- `kencleng-actors-entities.md` — who/what exists (roles, entities)
- `kencleng-frontend-tech-stack.md` — code architecture (folder
  structure, route paths, state management)
- `patterns.md` — reusable page-shape + state-handling definitions
  ("what shape does a page like this take")
- **This doc** — the missing middle layer: given all of the above,
  what pages does each persona actually see, what pattern does each
  use, and what can they do on each one

This doc doesn't restate business rules or layout mechanics in
detail — it references the phase-detail docs for "why"/"how" and
`patterns.md` for "what shape."

## Legend

- **(OPEN)** — still needs a decision
- Route paths follow `kencleng-frontend-tech-stack.md` App Router
  structure
- **Pattern** column names refer to `patterns.md` §A (List/Browse,
  Detail, Form, Dashboard/Summary, Curation/Review, Status/Tracking)
- "Baseline" = available regardless of phase, as long as the persona
  is authenticated (Phase 0 concerns)

## Shell & Benchmark Notes

Dashboard Shell (top-nav desktop, top-bar+hamburger mobile), auth
modal-vs-page split, Google OAuth full-redirect flow, and the
GoFundMe/Kitabisa-benchmarked public campaign layout are now defined
in `patterns.md` and `kencleng-frontend-tech-stack.md` — not
duplicated here. This section previously held that content; see those
two docs instead.

---

## 1. Guest (not logged in)

| Page | Pattern | Actions |
|---|---|---|
| `/` (home) | List/Browse | Browse highlighted campaigns |
| `/campaign` (list) | List/Browse | Browse/filter `published` campaigns |
| `/campaign/[id]` (detail) | Detail (public variant) | View description, progress bar, public donor list (org info shown inline — no separate org profile page), donate button. Also displays `beneficiary_description` |
| `/campaign/[id]/donate` (donation form) | Form (single-step, no revision cycle) | Fill `amount`, choose `payment_method` (transfer/debit/gopay/shopeepay/ovo/qris — simulated), optionally fill `guest_name`/`guest_email` (both independently optional), see nudge note about benefits of providing email, submit |
| `/donation/[id]/status` | Status/Tracking | Check donation status (`pending`/`success`/`failed`) without login — token-in-URL, see `kencleng-phase2-detail.md` Feature 1 |
| `/login` | Form | Login form + "Masuk dengan Google" button |
| `/register` | Form | Register form + "Daftar dengan Google" button |
| `/forgot-password` | Form | Submit email to request password reset |
| `/reset-password?token=...` | Form | Submit new password |

---

## 2. Donatur (registered)

All Guest pages, **plus**:

| Page | Pattern | Actions |
|---|---|---|
| Email verification link (from email, not a full page) | — | Click link → `AuthIdentity.verified_at` set |
| `/campaign/[id]/donate` | Form | Same form, **plus** `is_anonymous` checkbox (not available to Guest) |
| `/dashboard/donations` | List/Browse | View personal donation history (including anonymous ones) |
| `/dashboard/donations/claim` | List/Browse | View guest-donation candidates matching verified email, confirm each individually |
| `/dashboard/profile` | Form | Edit name, etc. (no profile picture — dropped from v1) |
| `/dashboard/security` | Form (multi-section) | Enable/disable MFA (QR scan + confirm code), view/regenerate backup codes, link/unlink Google identity. Google-only users also see "Atur Password" here — sets an `email_password` `AuthIdentity` (`verified_at = now` immediately) so unlink-Google becomes available. See `kencleng-phase0-detail.md` Feature 4 |
| `/dashboard/notifications` | List/Browse | View notification center, mark-as-read (batched client-side — see phase0-detail Feature 6) |
| ToS reminder banner (not a page) | — | Non-blocking, appears when a new ToS version is published; click to re-accept |

---

## 3. Organization Owner

All Donatur pages, **plus**:

| Page | Pattern | Actions |
|---|---|---|
| `/dashboard/organization/new` | Form (Revisable Submission) | Fill org data + upload legal docs (Akta, SK Kemenkumham, NPWP, optional Izin PUB) — shows `SecureUploadNote` |
| `/dashboard/organization/[id]` | Detail (dashboard variant) | View curation status, edit while `pending_verification`, revise & resubmit if `rejected` — legal docs visible here (owner-only). Still editable after `verified`, but editing a *legal/identity* field sends status back to `pending_verification`; editing an *operational* field never changes status — see `kencleng-phase1-detail.md` Feature 1 |
| `/dashboard/organization/[id]/representatives` | List/Browse + inline Form | Invite representative (as `staff`) by email — direct-add, no accept step — remove representative, promote/demote owner↔staff, view list — system enforces ≥1 owner. Full detail: `kencleng-roadmap-next-steps.md`, representatives spec discussion |
| `/dashboard/campaign/new` | Form (Revisable Submission) | Fill draft: title, description, target_amount, max_amount, deadline, upload media. Also `beneficiary_description` (free-text, optional) |
| `/dashboard/campaign/[id]/edit` | Form | Edit draft (while `status = draft`) |
| `/dashboard/campaign/[id]` | Detail + inline Form action | Submit to curation (owner-only action), view curation status, revise if `rejected` |
| `/dashboard/campaign/[id]/publish` | Form (single unified action) | Publish now / schedule `publish_at`, reschedule, unpublish (requires mandatory reason, logged to Audit Log — see `kencleng-phase1-detail.md` Feature 5), republish |
| `/dashboard/event/new` | Form | Fill event (name, datetime, location, description), link to own campaign(s) with status `published`/`scheduled` |
| `/dashboard/campaign/[id]/monitor` | Dashboard/Summary | View `collected_amount` vs `target_amount` (same data as public, from dashboard context) |
| `/dashboard/campaign/[id]/report` | Detail + inline Form action | View auto-generated summary (post-`closed`), **add/edit narrative** (`report_narrative` — optional, no curation gate, editable anytime) |
| `/dashboard/campaign/[id]/disbursement/new` | Form (Revisable Submission) | Request disbursement (owner-only, campaign must be `closed`) |
| `/dashboard/campaign/[id]/disbursement/[reqId]` | Detail | View request status, revise & resubmit if `rejected` |
| `/dashboard/campaign/[id]/fund-usage-report/new` | Form (Revisable Submission) | Fill expense breakdown per category + upload attachments — shows `SecureUploadNote` |
| `/dashboard/campaign/[id]/fund-usage-report/[reportId]` | Detail | View verification status, revise & resubmit if `rejected` |

---

## 4. Organization Staff

Subset of Owner pages — **same routes, sensitive actions hidden/disabled**:

| Page | Pattern | Actions |
|---|---|---|
| `/dashboard/organization/[id]` | Detail | **No access to legal document section** (Owner-only, per Business Rule 4 in actors-entities doc) |
| `/dashboard/campaign/new`, `/dashboard/campaign/[id]/edit` | Form | Create/edit draft — **"submit to curation" button hidden/disabled** |
| `/dashboard/campaign/[id]/publish` | — | **No access** — publish/unpublish is owner-only |
| `/dashboard/event/new` | Form | Same as Owner — event creation is not owner-exclusive |
| `/dashboard/campaign/[id]/monitor` | Dashboard/Summary | View only |
| `/dashboard/campaign/[id]/report` | Detail | View only — **cannot add/edit narrative** (owner-only per phase3-detail revision) |
| `/dashboard/campaign/[id]/disbursement/*` | — | **No access** |
| `/dashboard/campaign/[id]/fund-usage-report/*` | Detail | View only — submit remains owner-only |
| `/dashboard/organization/[id]/representatives` | — | **No access** — managing representatives is owner-only |

---

## 5. Kurator

| Page | Pattern | Actions |
|---|---|---|
| `/dashboard/kurasi` | List/Browse | **Unified queue** — shows all assignment types (organization curation, campaign curation, fund-usage-report verification) tagged by type, in one list |
| `/dashboard/kurasi/organization/[assignmentId]` | Curation/Review | Review legal docs (via signed URL), approve/reject + `decision_note` on reject |
| `/dashboard/kurasi/campaign/[assignmentId]` | Curation/Review | Review target/deadline/description/media, approve/reject + `decision_note` |
| `/dashboard/kurasi/fund-usage/[assignmentId]` | Curation/Review | Review expense breakdown + attachments, approve/reject + `decision_note` |

No pages in Phase 2 — Kurator has no direct action in the on-campaign
flow (no dispute mechanism designed yet).

---

## 6. Admin

| Page | Pattern | Actions |
|---|---|---|
| `/dashboard/admin/users` | List/Browse | Search users, assign/revoke Admin or Kurator role (system blocks Admin+Kurator/Representative combination) |
| `/dashboard/admin/kurasi-queue` | List/Browse | View pending organization/campaign/fund-usage-report items, assign to a specific Kurator (manual pick, conflict-of-interest check enforced) |
| `/dashboard/admin/campaign/[id]/force-close` | Curation/Review (single-action variant) | Force-close a `published` campaign anytime, mandatory `decision_note` |
| `/dashboard/admin/disbursement/[reqId]` | Curation/Review | Approve/reject disbursement request |

**Mobile note**: all 4 Admin pages reuse the standard Dashboard Shell
(top-bar + hamburger) with sections stacked linearly — no distinct
mobile-specific layout needed, per the List/Browse and Curation/Review
patterns' normal responsive behavior.

---

## Cross-Cutting UI Elements (not full pages)

| Element | Where used | Notes |
|---|---|---|
| `MaskedField` | Anywhere `guest_email`, `User.primary_email`, `NPWP`, or future banking details are displayed | See `patterns.md` §C for full behavior spec |
| `SecureUploadNote` | Organization legal doc upload, fund-usage-report attachment upload | See `patterns.md` §C |
| `CurationDecisionPanel` | All curation/review pages (Kurator + Admin force-close/disbursement) | See `patterns.md` §C and Pattern 5 |
| Notification badge / center | Persistent header element for any logged-in user | Unread count, batched mark-as-read |

---

## Open Items

1. ~~Guest donor display label~~ → **resolved: "Donatur"** when
   `guest_name` is omitted. **[RESOLVED — 2026-08-20]**
2. ~~Mobile-specific wireframes per page~~ — **superseded**: wireframes
   retired in favor of `patterns.md`'s pattern-level responsive
   behavior; no per-page mobile artifact needed
3. ~~Admin-only pages not yet wireframed~~ — **superseded**, same
   reason as above; Admin pages now covered by List/Browse and
   Curation/Review pattern definitions

## Resolved (moved to `patterns.md`)

- Empty/loading/error states per page — now defined once per pattern
  in `patterns.md` §A/§B, not per page
- ~~Mobile-specific PWA layout considerations beyond responsive
  stacking (offline states, install prompts)~~ → **resolved:
  app-shell caching only (static assets cacheable, data always
  live/stale-on-fetch, no offline write queue), browser-default
  install prompt (no custom install UI)** — see
  `kencleng-frontend-tech-stack.md` PWA Scope, and `patterns.md` §B
  for the resulting "stale data" state convention.
  **[RESOLVED — 2026-08-20]**

## Related Docs

- Pattern definitions & shared component behavior: `patterns.md`
- Visual tokens: `design-guidelines.md`
- Code architecture: `kencleng-frontend-tech-stack.md`
