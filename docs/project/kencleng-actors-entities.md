# Kencleng — Actors, Roles & Entities
 
> Status: Draft — decided at conceptual level. ERD/schema design is
> now complete (`kencleng-erd.md`); this doc still holds as the
> conceptual source of truth for actors/roles/business rules.
> Last updated: 2026-07-24 (doc-sync pass — all 3 Open Items resolved
> elsewhere, cross-referenced here)
 
## Context
 
This document records the actor model and core entities agreed upon
before moving into ERD/schema design. It complements the backend and
frontend tech stack docs.
 
## Entities
 
| Entity | Description |
|---|---|
| **User** | Base account. Can hold multiple roles simultaneously. No profile picture in v1. |
| **Organization** | A separate entity (not a role) — has its own name, legal documents, and verification status. Represents the party proposing/managing a campaign and its donation results. |
| **OrganizationRepresentative** | Many-to-many relation between User and Organization. Carries a `level`: `owner` (PIC) or `staff`. One organization can have many representatives; one user can represent many organization, **up to a maximum of 5 organization per user [RESOLVED — NEW, see `kencleng-phase1-detail.md` Fitur 1]**. |
| **Donation** | A single donation record. See fields below — supports both guest and registered donors. |
 
### Donation entity — key fields **[REVISED]**
 
- `donor_user_id` — nullable FK to User (null if donation made by a guest)
- `guest_name`, `guest_email` — **both nullable, independently of each
  other** — snapshot at time of donation (used only when
  `donor_user_id` is null). Neither is required to submit a donation
  as guest; see Business Rule 5 for consequences of omitting them.
- `payment_method` — enum, one of: `transfer`, `debit`, `gopay`,
  `shopeepay`, `ovo`, `qris`. Simulated in v1 (no real gateway
  integration), but recorded on every donation for record completeness.
- `is_anonymous` — boolean, independent of guest/registered status; controls public display only, never affects personal tracking/reporting
- `claimed_at` — set when a guest donation is later linked to a registered User via the claim flow (only possible if `guest_email` was provided at donation time — see Business Rule 5)
## Roles (held by User, multiple allowed)
 
| Role | Description | Can combine with | Cannot combine with |
|---|---|---|---|
| **Donatur** | Gives donations, receives donation reports (if registered) | Kurator, Representative (owner/staff) | — |
| **Kurator** | Curates campaigns before launch; verifies fund-usage reports after disbursement | Donatur, Representative (of a **different** organization) | Representative of the **same** organization they curate (conflict of interest — must recuse) |
| **Admin** | Platform-level administrative matters | — | Kurator, Representative of any organization |
 
## Business Rules — Agreed
 
1. **Kurator conflict of interest**: a kurator must recuse from curating a campaign, or verifying a fund-usage report, for any organization where they are also a representative.
2. **Admin isolation**: a user with the Admin role cannot simultaneously be a Kurator or a Representative of any organization — prevents self-approval/bypass of curation or fund-report verification.
3. **Organization must always have ≥1 Owner/PIC**: an owner cannot be removed/downgraded if it would leave the organization with zero owners.
4. **Sensitive organization actions are Owner-only**: requesting fund disbursement, managing representatives (invite/remove), and submitting fund-usage reports are restricted to the `owner` level. `staff` can create/edit campaign drafts and view reports, but **cannot view organization legal documents** (Akta, SK Kemenkumham, NPWP, Izin PUB) — legal document access is Owner-only. **[RESOLVED — NEW]** Representative invite is direct-add (no accept/consent step) — full mechanism in `kencleng-phase1-detail.md` Fitur 1B.
5. **Guest donations are supported, with fully optional identity** **[REVISED]**:
   - A donation does not require a registered account.
   - `guest_name` and `guest_email` are **both optional** — a guest may
     donate providing neither, either, or both.
   - Guest donor info, when provided, is stored as a snapshot on the
     Donation record itself, not as a live reference.
   - **Consequence of omitting `guest_email`**: the donation can never
     be matched/claimed later via the guest-claim flow (matching is
     email-based), and the donor cannot receive email notification of
     campaign outcomes (Phase 3, Fitur 1). This is by design, not a
     defect — an anonymous guest with no email genuinely leaves no
     trace to follow up with.
   - The donation form must show a short note explaining the benefit
     of providing an email (ability to track donation history if the
     donor later registers, and to receive campaign outcome
     notifications) — this is a UX nudge, not a requirement.
6. **Anonymous is separate from guest**: `is_anonymous` is a per-donation flag available to registered donors too — hides the donor's name from the public campaign donor list while the donation still appears in the donor's own personal history/reports.
7. **Guest donation claim flow (in scope for v1)**:
   - Triggered after a user registers and verifies their email.
   - System looks up guest donations matching the verified email and presents them as a list of candidates.
   - **Only donations with a non-null `guest_email` are eligible** for matching **[REVISED — clarifies interaction with Rule 5]**.
   - User must **manually confirm each donation individually** before it's linked (`claimed_at` set, `donor_user_id` populated) — no auto-claim, since a shared email could belong to donations made by someone else.
   - The original `guest_name`/`guest_email` snapshot and any historical public display (anonymous or not) are preserved as-is after claiming — claiming does not retroactively alter what was shown publicly at the time.
## PII Handling Note **[NEW]**
 
The following fields are considered PII and must be masked by default
on the frontend (see `kencleng-frontend-tech-stack.md` for the shared
masking component), with an explicit reveal toggle — masking applies
**regardless of viewer role, including Admin**:
- `guest_email`, `User.primary_email`
- `NPWP` (Organization)
- Any future banking/disbursement account details
Revealing a masked field is expected to be recorded in the Audit Log
(see `kencleng-phase0-detail.md`, Fitur 9) when performed by Admin or
Kurator on another party's data.
 
## Open Items — Needs Further Discussion
 
These weren't blocking to record the actor/entity model above, and
all 3 are now resolved (kept here for history — strikethrough per
project convention):
 
1. ~~Whether an organization's "beneficiary" can be a third party who
   isn't itself a platform user/entity (e.g., a specific disaster
   victim), and if so, whether that needs its own `Beneficiary` entity
   or just a descriptive field on the campaign.~~ → **resolved: a
   free-text `beneficiary_description` field on `Campaign`, not a
   dedicated entity** — see `kencleng-erd.md` and
   `kencleng-roadmap-next-steps.md`. **[RESOLVED]**
2. ~~Whether the same kurator handles both pre-launch curation and
   post-disbursement fund-report verification for a given campaign, or
   whether these can be assigned to different kurators (not critical
   for v1, can default to either).~~ → **resolved: different kurators
   are allowed per assignment** — matches `kencleng-erd.md`'s
   structure (each curation/verification table has its own
   independent `kurator_id`, no cross-table constraint forcing the
   same actor). **[RESOLVED]**
3. ~~Verification/KYC process details for an Organization (what
   documents, who reviews — admin vs kurator).~~ → **resolved**: Akta
   Notaris, SK Kemenkumham, NPWP, Izin PUB (optional) — reviewed by an
   Admin-assigned Kurator, who must recuse if they're also a
   representative of the organization being reviewed. Full flow in
   `kencleng-phase1-detail.md` Fitur 1 (submission) & Fitur 2
   (curation). **[RESOLVED]**
## Not Yet Discussed
 
- ~~Full end-to-end business process flow (campaign proposal →
  curation → live → donations → target/deadline → disbursement →
  fund-usage report)~~ → **resolved**, see
  `kencleng-business-process-overview.md` and
  `kencleng-phase0-detail.md` through `kencleng-phase3-detail.md`
  **[RESOLVED]**
- ~~Campaign/Event entity details and lifecycle states~~ → **resolved**,
  see `kencleng-phase1-detail.md` Fitur 3/5/6 and `kencleng-erd.md`
  **[RESOLVED]**
- ~~ERD / schema design~~ → **resolved**, see `kencleng-erd.md`
  **[RESOLVED]**