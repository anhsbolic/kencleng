# Kencleng — UX Page Map (Persona × Phase × Page)
 
> Status: Draft — derived from persona/phase discussion; will be
> refined incrementally as open items resolve.
> Last updated: 2026-07-24
 
## Context
 
This document maps **which pages exist and what each persona can do
on them**, cutting across Phase 0-3. It sits at a different layer than
the phase-detail docs:
 
- `kencleng-phase0-3-detail.md` — business rules & flow per feature
  (what happens, in what order, what data)
- `kencleng-actors-entities.md` — who/what exists (roles, entities)
- `kencleng-frontend-tech-stack.md` — how it's implemented (folder
  structure, route paths, component choices)
- **This doc** — the missing middle layer: given all of the above,
  what pages does each persona actually see, and what can they do on
  each one
This doc doesn't restate business rules in detail — it references the
phase-detail docs for "why"/"how", and focuses on "which page, which
action, gated how."
 
## Legend
 
- **(OPEN)** — still needs a decision
- Route paths follow `kencleng-frontend-tech-stack.md` App Router
  structure
- "Baseline" = available regardless of phase, as long as the persona
  is authenticated (Phase 0 concerns)
## Dashboard Shell [NEW — RESOLVED 2026-07-20]
 
All authenticated ("dashboard") pages share one shell: **horizontal
top-nav on desktop (no sidebar)**, top-bar + hamburger menu on mobile.
Decided during Step 1.5 wireframing (Batch 2) once the first
authenticated page (`/dashboard/security`) needed a layout.
 
Worth revisiting only if the nav genuinely gets crowded once
Organization Owner/Staff, Kurator, and Admin personas are implemented
(they have noticeably more nav items than Donatur's 5) — that's a
concrete, demonstrated need, not a reason to add a sidebar
preemptively now.
 
## Benchmark Design Reference [NEW — RESOLVED 2026-07-20]
 
**GoFundMe** (primary), **Kitabisa** (secondary) — structural/UX
patterns only (layout, information hierarchy, flow), not visual
styling (colors/typography stay deferred to Step 5 per the roadmap).
Key patterns adopted:
- Auth (login/register/forgot-password) = modal overlay on desktop,
  full page on mobile
- `/reset-password` = full page on **both** devices — deliberate
  exception, since the user arrives via an emailed link with no
  in-app context to overlay a modal on
- `/campaign/[id]`: hero image → progress bar prominent → sticky
  Donate CTA → donor list separate from narrative body (mobile);
  two-column layout with sticky donate sidebar (desktop)
- `/campaign/[id]/donate` field order: nominal (preset chips +
  custom) → payment method → optional donor info — this already
  matched the existing field order in `phase2-detail.md` Feature 1, so
  no rework was needed there
---
 
## 1. Guest (not logged in)
 
| Page | Actions |
|---|---|
| `/` (home) | Browse highlighted campaigns |
| `/campaign` (list) | Browse/filter `published` campaigns |
| `/campaign/[id]` (detail) | View description, progress bar, public donor list (org info shown inline — no separate org profile page, per discussion), donate button. **[NEW]** Also displays `beneficiary_description` (free-text field on `Campaign`, resolved — see `kencleng-erd.md`) |
| `/campaign/[id]/donate` (donation form) | Fill `amount`, choose `payment_method` (transfer/debit/gopay/shopeepay/ovo/qris — simulated), optionally fill `guest_name`/`guest_email` (both independently optional), see nudge note about benefits of providing email, submit |
| `/donation/[id]/status` | Check donation status (`pending`/`success`/`failed`) without login — ~~mechanism (token-based vs other) still **(OPEN)**~~ → **resolved: token-in-URL**, see `kencleng-phase2-detail.md` Feature 1 **[RESOLVED]** |
| `/login` | Login form + "Masuk dengan Google" button |
| `/register` | Register form + "Daftar dengan Google" button |
| `/forgot-password` | Submit email to request password reset |
| `/reset-password?token=...` | Submit new password |
 
---
 
## 2. Donatur (registered)
 
All Guest pages, **plus**:
 
| Page | Actions |
|---|---|
| Email verification link (from email, not a full page) | Click link → `AuthIdentity.verified_at` set |
| `/campaign/[id]/donate` | Same form, **plus** `is_anonymous` checkbox (not available to Guest) |
| `/dashboard/donations` | View personal donation history (including anonymous ones) |
| `/dashboard/donations/claim` | View guest-donation candidates matching verified email, confirm each individually |
| `/dashboard/profile` | Edit name, etc. (no profile picture — dropped from v1) |
| `/dashboard/security` | Enable/disable MFA (QR scan + confirm code), view/regenerate backup codes, link/unlink Google identity. **[RESOLVED — NEW]** Google-only users also see "Atur Password" here — sets an `email_password` `AuthIdentity` (`verified_at = now` immediately, no extra re-auth required) so unlink-Google becomes available. See `kencleng-phase0-detail.md` Feature 4. |
| `/dashboard/notifications` | View notification center, mark-as-read (batched client-side, not per-click sync — see phase0-detail Feature 6) |
| ToS reminder banner (not a page) | Non-blocking, appears when a new ToS version is published; click to re-accept |
 
---
 
## 3. Organization Owner
 
All Donatur pages, **plus**:
 
| Page | Actions |
|---|---|
| `/dashboard/organization/new` | Fill org data + upload legal docs (Akta, SK Kemenkumham, NPWP, optional Izin PUB) — shows `SecureUploadNote` |
| `/dashboard/organization/[id]` | View curation status, edit while `pending_verification`, revise & resubmit if `rejected` — legal docs visible here (owner-only). **[NEW — RESOLVED]** Still editable after `verified`, but editing a *legal/identity* field (Name, NPWP, Akta, SK Kemenkumham, Izin PUB) sends status back to `pending_verification`; editing an *operational* field (Description, Contact) never changes status — see `kencleng-phase1-detail.md` Feature 1 |
| `/dashboard/organization/[id]/representatives` | Invite representative (as `staff`) by email, remove representative, promote/demote owner↔staff, view list — system enforces ≥1 owner. **[RESOLVED — NEW]** Invite is **direct-add, no accept step**: owner enters the email of an existing, verified user → system creates the `OrganizationRepresentative` row immediately (`level = staff`) and notifies the invitee. Rejected a consent/accept-invite flow as unjustified complexity for v1 — `staff` access is low-risk (no legal docs, no sensitive actions per Business Rule 4). Promote/demote and removal are owner-only and blocked if they'd leave the organization with 0 owners. Full detail: `kencleng-roadmap-next-steps.md`, "Critical Open Items Resolved" / representatives spec discussion. |
| `/dashboard/campaign/new` | Fill draft: title, description, target_amount, max_amount, deadline, upload media. **[NEW]** Also `beneficiary_description` (free-text, optional) — resolved as a simple field rather than a dedicated `Beneficiary` entity, see `kencleng-erd.md` |
| `/dashboard/campaign/[id]/edit` | Edit draft (while `status = draft`) |
| `/dashboard/campaign/[id]` | Submit to curation (owner-only action), view curation status, revise if `rejected` |
| `/dashboard/campaign/[id]/publish` | Publish now / schedule `publish_at`, reschedule, unpublish (**[NEW — RESOLVED]** requires mandatory reason, logged to Audit Log — see `kencleng-phase1-detail.md` Feature 5), republish |
| `/dashboard/event/new` | Fill event (name, datetime, location, description), link to own campaign(s) with status `published`/`scheduled` |
| `/dashboard/campaign/[id]/monitor` | View `collected_amount` vs `target_amount` (same data as public, from dashboard context) |
| `/dashboard/campaign/[id]/report` | View auto-generated summary (post-`closed`), **add/edit narrative** (`report_narrative` — optional, no curation gate, editable anytime) |
| `/dashboard/campaign/[id]/disbursement/new` | Request disbursement (owner-only, campaign must be `closed`) |
| `/dashboard/campaign/[id]/disbursement/[reqId]` | View request status, revise & resubmit if `rejected` |
| `/dashboard/campaign/[id]/fund-usage-report/new` | Fill expense breakdown per category + upload attachments — shows `SecureUploadNote` |
| `/dashboard/campaign/[id]/fund-usage-report/[reportId]` | View verification status, revise & resubmit if `rejected` |
 
---
 
## 4. Organization Staff
 
Subset of Owner pages — **same routes, sensitive actions hidden/disabled**:
 
| Page | Actions |
|---|---|
| `/dashboard/organization/[id]` | **No access to legal document section** (Owner-only, per Business Rule 4 in actors-entities doc) |
| `/dashboard/campaign/new`, `/dashboard/campaign/[id]/edit` | Create/edit draft — **"submit to curation" button hidden/disabled** |
| `/dashboard/campaign/[id]/publish` | **No access** — publish/unpublish is owner-only |
| `/dashboard/event/new` | Same as Owner — event creation is not owner-exclusive |
| `/dashboard/campaign/[id]/monitor` | View only |
| `/dashboard/campaign/[id]/report` | View only — **cannot add/edit narrative** (owner-only per phase3-detail revision) |
| `/dashboard/campaign/[id]/disbursement/*` | **No access** |
| `/dashboard/campaign/[id]/fund-usage-report/*` | View only — submit remains owner-only |
| `/dashboard/organization/[id]/representatives` | **No access** — managing representatives is owner-only |
 
---
 
## 5. Kurator
 
| Page | Actions |
|---|---|
| `/dashboard/kurasi` | **Unified queue** — shows all assignment types (organization curation, campaign curation, fund-usage-report verification) tagged by type, in one list |
| `/dashboard/kurasi/organization/[assignmentId]` | Review legal docs (via signed URL), approve/reject + `decision_note` on reject — uses `CurationDecisionPanel` |
| `/dashboard/kurasi/campaign/[assignmentId]` | Review target/deadline/description/media, approve/reject + `decision_note` — uses `CurationDecisionPanel` |
| `/dashboard/kurasi/fund-usage/[assignmentId]` | Review expense breakdown + attachments, approve/reject + `decision_note` — uses `CurationDecisionPanel` |
 
No pages in Phase 2 — Kurator has no direct action in the on-campaign flow (no dispute mechanism designed yet).
 
---
 
## 6. Admin
 
| Page | Actions |
|---|---|
| `/dashboard/admin/users` | Search users, assign/revoke Admin or Kurator role (system blocks Admin+Kurator/Representative combination). **[RESOLVED — NEW]** Wireframed — see `kencleng-wireframes/admin-users-wireframe.html` |
| `/dashboard/admin/kurasi-queue` | View pending organization/campaign/fund-usage-report items, assign to a specific Kurator (manual pick, conflict-of-interest check enforced). **[RESOLVED — NEW]** Wireframed — see `kencleng-wireframes/admin-kurasi-queue-wireframe.html` |
| `/dashboard/admin/campaign/[id]/force-close` | Force-close a `published` campaign anytime, mandatory `decision_note`. **[RESOLVED — NEW]** Wireframed — see `kencleng-wireframes/admin-force-close-wireframe.html` |
| `/dashboard/admin/disbursement/[reqId]` | Approve/reject disbursement request |
 
**Mobile note [RESOLVED — NEW]**: unlike the auth pages and public campaign
pages (which have genuinely distinct mobile layouts), the 3 wireframed
Admin pages above reuse the same Dashboard Shell already resolved
(top-bar + hamburger) with sections stacked linearly — no separate
mobile wireframe was produced, since there's no distinct mobile-specific
layout decision here, just the existing shell's normal responsive
stacking.
 
---
 
## Cross-Cutting UI Elements (not full pages)
 
| Element | Where used | Notes |
|---|---|---|
| `MaskedField` | Anywhere `guest_email`, `User.primary_email`, `NPWP`, or future banking details are displayed | Masked by default, explicit reveal toggle per field, **applies even to Admin**. Reveal by Admin/Kurator on another party's data logs to Audit Log. **Reveal persistence [RESOLVED — NEW]**: stays revealed until manual re-toggle or page refresh/navigation — implemented as plain local component state (`useState`), no timeout/auto-re-mask logic. Navigating away unmounts the component and resets it for free; no extra engineering needed. |
| `SecureUploadNote` | Organization legal doc upload, fund-usage-report attachment upload | Reassurance note/popup, no logic |
| `CurationDecisionPanel` | All 3 Kurator decision pages | Approve/reject + mandatory `decision_note` on reject |
| Notification badge / center | Persistent header element for any logged-in user | Unread count, batched mark-as-read |
 
---
 
## Shared Patterns Identified
 
1. **Curation Queue + Decision pattern** — same interaction shape
   across organization curation, campaign curation, and fund-usage-report
   verification (Admin assigns → Kurator reviews → approve/reject with
   note). Implemented once as `CurationDecisionPanel`, parameterized
   by curation type.
2. **Revisable Submission pattern** — draft → submit → locked →
   revise-if-rejected → resubmit. Appears in campaign registration and
   fund-usage-report submission. Not yet extracted as a formal reusable
   component, but worth keeping consistent in interaction design.
---
 
## Open Items
 
1. ~~Guest donation status page mechanism~~ → **resolved: token-in-URL**
   (`/donation/[id]/status?token=...`). Token is returned in the
   donation submit response, and also emailed if `guest_email` was
   provided. Token does not expire (read-only, non-destructive lookup
   — unlike the reset-password token, no need for a tight validity
   window). **[RESOLVED — NEW]**
2. ~~`MaskedField` reveal persistence~~ → **resolved: stays revealed
   until manual re-toggle or page refresh/navigation**, implemented as
   local component state — see Cross-Cutting UI Elements above
   **[RESOLVED — NEW]**
3. ~~Forgot-password behavior for Google-only users~~ — resolved, see
   phase0-detail.md Feature 2B (send a notification email "use Google
   login")
4. **Guest donor display label** when `guest_name` is omitted (e.g.
   "Hamba Allah"/"Donatur" — placeholder, final copy deferred to FE
   design)
5. ~~Mobile-specific layout considerations for PWA~~ → **resolved:
   every page gets both a mobile AND a desktop wireframe** — see
   "Dashboard Shell" and "Benchmark Design Reference" sections above,
   and the Step 1.5 wireframe artifact (`kencleng-wireframes/`) for
   the full set. **[RESOLVED — NEW]**
6. ~~`/dashboard/organization/[id]/representatives` — business rules not
   yet written~~ → **resolved: direct-add invite (no accept step),
   promote/demote and removal owner-only with the ≥1-owner guard** —
   see the representatives page row under Organization Owner above
   **[RESOLVED — NEW]**
7. ~~Set-password flow for Google-only users~~ → **resolved: via
   `/dashboard/security` "Atur Password"** — see the `/dashboard/security`
   row under Donatur above and `kencleng-phase0-detail.md` Feature 4
   **[RESOLVED — NEW]**
8. ~~Admin-only pages not yet wireframed~~ → **resolved: all 3
   wireframed** (`admin-users-wireframe.html`,
   `admin-kurasi-queue-wireframe.html`,
   `admin-force-close-wireframe.html` in `kencleng-wireframes/`) — see
   the Admin section above **[RESOLVED — NEW]**
## Not Yet Discussed
 
- Empty states / loading states per page (deferred to actual UI
  design phase)
- Mobile-specific PWA layout considerations — **partially addressed**:
  every page now has a distinct mobile wireframe (see Dashboard Shell /
  Benchmark sections above), but fine-grained PWA behaviors (offline
  states, install prompts, etc.) remain undiscussed
- Error state handling per page (e.g. campaign not found, expired
  token pages)