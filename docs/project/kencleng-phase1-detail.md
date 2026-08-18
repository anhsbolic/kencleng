# Kencleng — Phase 1 (Pre-Campaign) — Detailed Feature Spec
 
> Status: Draft — detailed at business-rule/flow level. API contract format resolved (OpenAPI 3.x spec-first, see `kencleng-backend-tech-stack.md`); endpoint-level detail still excluded from this doc pending actual spec authoring.
> Last updated: 2026-07-24 (rev 2 — representative management spec added, NPWP validation & org-per-user limit resolved, campaign category/location/beneficiary fields added, notification/audit-trail items resolved)
 
## Context
 
This document details each step of Phase 1 (Pre-Campaign) from
`kencleng-business-process-overview.md` down to flow, business rules,
data, concurrency, and security level — one level above actual API/DB
schema design.
 
## Document format per feature
 
Each feature below follows this structure:
 
- **Overview** — short description, purpose
- **Actors** — who is involved
- **Preconditions** — required state before this feature can run
- **Flow (happy path)** — numbered main steps
- **Business rules / validation**
- **Alternate flows / edge cases**
- **Data touched** — entities/fields read or written
- **State transition** — status field before → after
- **Concurrency & correctness notes**
- **Security notes**
- **Open questions**
Endpoint-level detail (concrete paths/methods) is intentionally
excluded from this doc — that lives in `api/openapi.yaml` per the
resolved API contract format decision (`kencleng-backend-tech-stack.md`,
"API Contract & Codegen"), not duplicated here.
 
---
 
## 1. Organization Registration
 
**Overview**
The initial process by which an organization registers on the platform
so it can submit campaigns. Produces a new `Organization` entity and
sets the registering user as its first representative, with
`level = owner`.
 
**Actors**
User (prospective Owner) — must already have a registered account with
a verified email.
 
**Preconditions**
- User is already logged in (verified account)
- User does not currently hold the Admin role (Admin is not allowed to
  also be an Organization Representative)
**Flow (happy path)**
1. User opens the organization registration form
2. Fills in organization data: name, description, contact, and legal
   documents:
   - Akta Notaris Pendirian (notarial deed of establishment)
   - SK Kemenkumham (Ministry of Law & Human Rights decree)
   - NPWP (tax ID number, locked as a unique constraint)
   - Izin PUB (public fundraising permit, optional/additional, not
     required in v1)
   **[NEW]** The legal-document upload form shows a `SecureUploadNote`
   (see `kencleng-phase0-detail.md` Feature 7 &
   `kencleng-frontend-tech-stack.md`) — reassurance that the file is
   stored securely and confidentially in a private bucket.
3. Submit → the system creates an `Organization` record with
   `status = pending_verification`
4. The system automatically creates an `OrganizationRepresentative` for
   that user with `level = owner`
5. The organization enters the **Organization Curation** queue
**Business rules / validation**
- A user with the Admin role may not register an organization
- NPWP must be unique across the whole platform — one NPWP can only
  be registered under one Organization (organization names may repeat
  across different organization)
- **NPWP format validation [RESOLVED — NEW]**: validated as a
  **format check only** — a regex against the standard pattern
  `XX.XXX.XXX.X-XXX.XXX` (15 digits) before encryption. There is
  **no** validity check against the DJP/Ditjen Pajak (tax authority)
  database (out of scope for this sandbox project, would require
  integrating a government API). Genuine authenticity is still
  verified manually by the Kurator through legal-document review.
- **Organization-per-user limit [RESOLVED — NEW]**: a maximum of **5
  organization** per user (counted from the number of
  `OrganizationRepresentative` rows with `level = owner` belonging to
  that user). A round number loose enough for reasonable cases
  (representing several foundations) while still setting a clear cap
  against abuse (spam registration of fake organization) — not derived
  from concrete data, since this is a sandbox project.
- An organization automatically has ≥1 owner from the start (the
  registrant), so the rule "an organization must have ≥1 owner" is
  automatically satisfied
- Organization data can be freely edited while still
  `pending_verification`
- **Field classification for re-curation [RESOLVED — NEW]**:
  organization fields are split into two classes:
  - **Legal/identity** (editing any of these after `verified`
    triggers re-curation): Organization Name, NPWP, Akta Notaris
    Pendirian, SK Kemenkumham, Izin PUB
  - **Operational** (freely editable at any time, never changes
    status): Description, Contact
  Rationale: the Kurator verifies the organization's legality and
  identity — not marketing copy or contact info — so only changes in
  the first group should reopen the verification question.
- ~~Edit-after-verified~~ → **resolved: remains editable after
  `verified`, but submitting a change to a legal/identity field sends
  `Organization.status` back to `pending_verification`** (a re-curation
  cycle, mirroring the resubmit pattern in Feature 2). Operational
  fields never change the status. **[RESOLVED]**
**Alternate flows / edge cases**
- A user with the Admin role attempts to submit → rejected
- Double-submit (clicking submit twice) → must be idempotent, must not
  create 2 organization records
- One user can register more than one organization (many-to-many is
  supported), **up to the 5-organization limit [RESOLVED — NEW]** — a
  6th registration attempt is rejected with a clear message
- **[NEW — RESOLVED]** Owner edits a legal/identity field after
  `verified` → the system shows a confirmation dialog ("Changing
  [Name/NPWP/Legal Document] will send the organization back into the
  re-verification queue, and temporarily suspend all currently-live
  campaigns. Continue?") before actually saving the change — see the
  full consequences in `kencleng-phase1-detail.md` Feature 5
**Data touched**
- **Create** `Organization` (status, name, description, contact, akta,
  sk_kemenkumham, npwp [unique], izin_pub [nullable])
- **Create** `OrganizationRepresentative` (user_id, organization_id,
  level = `owner`)
**State transition**
`Organization.status`: *(none)* → `pending_verification`
*(after verified)* → **[NEW]** back to `pending_verification` if a
legal/identity field is edited
 
**Concurrency & correctness notes**
- Creating `Organization` + `OrganizationRepresentative` (owner) must
  happen in a single DB transaction — there must never be an
  organization state with no owner
- Idempotency guard to prevent duplicate submits
- **[NEW]** The 5-organization limit check is done in the same
  transaction as creating the `Organization` (`SELECT COUNT(*) ... FOR
  UPDATE` or an equivalent guard) so that two near-simultaneous
  submits from the same user can't both pass and give that user 6
  organization
**Security notes**
- Endpoint must be authenticated
- Legal document upload: validate file type (whitelist, e.g.
  PDF/JPG/PNG) & maximum size
- **[NEW]** The legal-document upload form shows a `SecureUploadNote`
  (UX reassurance, not a new technical mechanism — see phase0-detail
  Feature 7)
- **[NEW]** `NPWP` falls under the PII category (see
  `kencleng-actors-entities.md`, PII Handling Note) — wherever it's
  displayed again in the UI (e.g. the organization detail page), it must
  go through `MaskedField` with a reveal toggle, applying even to
  Admin — **including when the data owner (Owner) is viewing their own
  organization's NPWP** **[NEW — clarification]**
**Open questions**
- ~~Detailed format validation per document type (e.g. NPWP format)~~
  → **resolved: format-only validation** — see Business Rules above
  **[RESOLVED — NEW]**
- ~~Is there a limit on how many organization one user can register~~ →
  **resolved: maximum 5** — see Business Rules above
  **[RESOLVED — NEW]**
---
 
## 1B. Manage Organization Representative **[NEW — RESOLVED]**
 
**Overview**
The `/dashboard/organization/[id]/representatives` page — the Owner adds/
removes representatives (`staff`), and promotes/demotes between
`staff` and `owner`. This is business-rule spec that hadn't previously
been written out (it was only referenced indirectly via Business Rule
3 in `kencleng-actors-entities.md`).
 
**Actors**
Owner (the only level able to manage representatives, per Business
Rule 4 in `kencleng-actors-entities.md`)
 
**Preconditions**
The user performing the action is an active representative with
`level = owner` on that organization
 
**Flow (happy path) — Invite a new representative**
1. Owner opens the representatives page, enters the prospective
   representative's email
2. The system checks: does a registered user with that email exist,
   and is their **email verified**?
   - **No** → reject, with a clear message ("the user must already be
     registered with a verified email")
   - **Yes** → immediately create a new `OrganizationRepresentative`
     (`level = staff`) — **direct-add, no accept/consent step**
3. The added user receives a notification ("you've been added as a
   representative of organization X")
**Flow — Promote staff → owner**
1. Owner selects a representative with `level = staff`, clicks "make
   owner"
2. `OrganizationRepresentative.level` → `owner` (any owner may do this,
   multi-owner is supported)
**Flow — Demote owner → staff**
1. Owner selects another representative with `level = owner`, clicks
   "demote to staff"
2. The system checks: would this leave the organization with 0 owners?
   - **Yes** → reject, clear message (Business Rule 3)
   - **No** → `level` → `staff`
**Flow — Remove a representative**
1. Owner selects any representative (staff or another owner), clicks
   "remove"
2. The system checks: if the target is an `owner` and this is the
   last owner → reject, clear message
3. If the check passes → delete the `OrganizationRepresentative` row
**Business rules / validation**
- **Invite: direct-add, no consent step [RESOLVED — NEW]** — chosen
  over invite-with-accept because `staff`-level access is low-risk
  (no legal document access, no sensitive actions per Business Rule
  4) — no need for an extra consent state machine to mitigate a small
  risk
- Invites can only go to registered users with an email that is
  **already verified** — you can't invite via an email that has no
  account yet (no pending-invite-by-unregistered-email in v1)
- Promote/demote & removal: **owner-only**, for any of these actions
- **An owner cannot self-demote/self-remove if it would leave 0
  owners** — same as demoting/removing another owner, identical guard
  (Business Rule 3)
- An owner **may** self-remove/self-demote as long as ≥1 other owner
  remains afterward
**Alternate flows / edge cases**
- Inviting an email that's already a representative of the same
  organization → rejected, already a representative
- Demoting/removing the last owner → rejected in both flows (same
  guard)
- Staff attempting to access this page → **no access**, managing
  representatives is owner-only (see `kencleng-ux-page-map.md`
  Organization Staff section)
**Data touched**
- **Create** `OrganizationRepresentative` (user_id, organization_id,
  level = `staff`) — invite
- **Update** `OrganizationRepresentative.level` — promote/demote
- **Delete** `OrganizationRepresentative` — removal
**Concurrency & correctness notes**
The "≥1 owner" guard for demote/remove must be atomic — `COUNT(*)
WHERE organization_id = ? AND level = 'owner'` is checked within the
same transaction as the `UPDATE`/`DELETE`, so that two
near-simultaneous demotes/removals against the last two owners can't
both succeed and leave 0 owners.
 
**Security notes**
Invite/promote/demote/remove endpoints: authenticated + a
representative with `level = owner` on that organization. All of these
actions fall within Audit Log scope
(`kencleng-phase0-detail.md` Feature 9) — see the notes there.
 
**Open questions**
None — this spec is closed. (Previously logged as an open item in
`kencleng-roadmap-next-steps.md` and `kencleng-ux-page-map.md`, now
resolved here.)
 
---
 
## 2. Organization Curation
 
**Overview**
The review process by which a Kurator verifies the legal-document
authenticity of a newly registered organization, before that organization
is allowed to submit campaigns.
 
**Actors**
Admin (assigns) · Kurator (performs the review & decision)
 
**Preconditions**
`Organization.status = pending_verification`
 
**Flow (happy path)**
1. A new organization enters `pending_verification` → the system
   notifies the Admin
2. The Admin assigns one Kurator to that organization (manual, chosen by
   the Admin) → creates an `OrganizationCurationAssignment` record
   (`decision = pending`)
3. The assigned Kurator opens the organization's detail page, reviews
   the documents
4. **Approve** → `Organization.status = verified`,
   `assignment.decision = approved`
   or **Reject** (`decision_note` required) →
   `Organization.status = rejected`, `assignment.decision = rejected`
**Alternate flows / edge cases — resubmit after rejection**
5. The organization revises its data & resubmits →
   `Organization.status` returns to `pending_verification`
6. A new assignment cycle starts from step 2 (the old assignment is
   kept as history)
 
**Business rules / validation**
- A Kurator who is also a representative of that organization may not
  be assigned/take on this curation (conflict of interest)
- Only one `OrganizationCurationAssignment` may be active
  (`decision = pending`) per organization at a time
**Data touched**
- **Update** `Organization.status` → `verified` / `rejected`
- **Create** `OrganizationCurationAssignment` (organization_id, kurator_id,
  assigned_by, assigned_at, decision, decision_note, decided_at)
**State transition**
`Organization.status`: `pending_verification` → `verified` / `rejected`
→ *(revision)* `pending_verification` (repeat cycle)
 
**Concurrency & correctness notes**
- Constraint: only one active assignment per organization — prevents
  the Admin from assigning 2 different Kurator to the same organization
  at the same time
- Because the model is assigned (not an open queue), a race between
  Kurator is avoided by design
**Security notes**
- Only the assigned Kurator role can approve/reject that assignment
- Enforce the conflict-of-interest check on the backend, not only in
  the UI
**Open questions**
- ~~Notification format to the Admin when a new organization enters the
  queue~~ → **resolved: notification `type = admin_new_curation_item`,
  dual channel (in-app + email)** — see
  `kencleng-phase0-detail.md` Feature 6 **[RESOLVED — NEW]**
---
 
## 3. Campaign Registration
 
**Overview**
The process by which a representative of an already-`verified`
organization creates a new campaign — from draft through submission for
curation.
 
**Actors**
Owner & Staff (for drafting) · Owner-only (for submitting to curation)
 
**Preconditions**
- `Organization.status = verified`
- The user is an active representative of that organization (owner or
  staff)
- **`Organization.has_overdue_report = false` [RESOLVED — NEW]** — an
  organization currently flagged for an overdue fund-usage report (see
  `kencleng-phase3-detail.md` Feature 4) cannot create a new campaign
  until that flag is cleared (automatically, once the overdue report
  is finally submitted)
**Flow (happy path)**
1. A representative (Owner/Staff) opens the new-campaign form
2. Fills in data: title, description, **`category`** (enum, required
   — natural disaster/health/education/social/other **[RESOLVED —
   NEW]**), **`location`** (free-text, optional — city/province
   **[RESOLVED — NEW]**), **`beneficiary_description`** (free-text,
   optional — description of the beneficiary **[RESOLVED — NEW]**),
   `target_amount`, `max_amount` (optional), `deadline`, media/images
3. Saved as a draft → `Campaign.status = draft`
4. Can be edited repeatedly by Owner/Staff while still `draft`
5. Owner clicks "submit for curation" →
   `Campaign.status = pending_curation` (locked from editing)
**Business rules / validation**
- Only representatives of a `verified` organization can create a
  campaign for that organization
- **An organization with `has_overdue_report = true` cannot create a new
  campaign [RESOLVED — NEW]** — the check happens at the same point as
  the `Organization.status = verified` check above. Existing campaigns
  (draft/published/etc.) are **not affected** — this only blocks
  creating a *new* campaign, consistent with
  `kencleng-phase3-detail.md` Feature 4
- `target_amount` must be > 0
- `max_amount` (if filled in) must be ≥ `target_amount`
- `deadline` must be in the future at creation time
- `category` is required (one of a fixed enum) — used for filtering
  on the `/campaign` list page **[RESOLVED — NEW]**
- `location` and `beneficiary_description` are both optional,
  free-text — not a separate relational/geo structure, just enough
  for display & coarse filtering **[RESOLVED — NEW]**
- Submitting for curation can only be done by the Owner
- Once `pending_curation`, a campaign **cannot** be pulled back to
  `draft` unilaterally by the Owner — it must wait for the Kurator's
  decision
**Alternate flows / edge cases**
- Staff attempting to submit-for-curation → rejected (authorization)
- **An organization with `has_overdue_report = true` attempting to
  create a new draft campaign → rejected, with a clear message
  pointing to submitting the overdue fund-usage report first
  [RESOLVED — NEW]**
- A draft deleted before submission → allowed
- Multi-editor: Owner & Staff editing the same draft — **last-write-
  wins** is considered sufficient for v1, no need for optimistic
  locking
**Data touched**
- Create/Update `Campaign` (organization_id, title, description,
  category, location, beneficiary_description, target_amount,
  max_amount, deadline, status, created_by, media)
**State transition**
`Campaign.status`: *(none)* → `draft` → `pending_curation`
 
**Concurrency & correctness notes**
- The `draft → pending_curation` transition must be atomic (status
  guard at the query level)
- Multi-editor draft: last-write-wins, no optimistic locking in v1
**Security notes**
- Create/edit draft: authenticated + a representative (any level) of
  that organization
- Submit-for-curation: authenticated + a representative with
  `level = owner`
**Open questions**
- ~~Additional fields on the campaign registration form (category,
  beneficiary location, etc.)~~ → **resolved: `category` (required
  enum), `location` (optional free-text),
  `beneficiary_description` (optional free-text)** — see Flow &
  Business Rules above **[RESOLVED — NEW]**
---
 
## 4. Campaign Curation
 
**Overview**
The Kurator's review process for a campaign that has already been
submitted (`pending_curation`), before the campaign may be published
publicly. Mirrors the Organization Curation pattern.
 
**Actors**
Admin (assigns) · Kurator (performs the review & decision, must recuse
if a representative of the organization that owns that campaign)
 
**Preconditions**
`Campaign.status = pending_curation`
 
**Flow (happy path)**
1. A campaign enters `pending_curation` → the system notifies the
   Admin
2. The Admin assigns one Kurator (manual) → creates a
   `CampaignCurationAssignment` (`decision = pending`)
3. The Kurator reviews the campaign's eligibility (target/deadline
   fit, description, media, etc.)
4. **Approve** → `Campaign.status = approved`, ready to be published
   or **Reject** (`decision_note` required) → `Campaign.status = rejected`
**Alternate flows / edge cases — resubmit after rejection**
5. The Owner revises a `rejected` campaign → `Campaign.status`
   goes back to `draft`
6. The Owner resubmits → `pending_curation`, a new assignment cycle
   (the old assignment is kept as history)
 
**Business rules / validation**
- A Kurator who is also a representative of the organization that owns
  the campaign may not be assigned (conflict of interest)
- Only one `CampaignCurationAssignment` may be active per campaign at
  a time
**Data touched**
- **Update** `Campaign.status` → `approved` / `rejected`
- **Create** `CampaignCurationAssignment` (campaign_id, kurator_id,
  assigned_by, assigned_at, decision, decision_note, decided_at)
**State transition**
`Campaign.status`: `pending_curation` → `approved` | `rejected` →
*(revision)* `draft` → `pending_curation` (repeat cycle)
 
**Concurrency & correctness notes**
Same as Organization Curation — one active assignment per campaign
prevents a race between Kurator.
 
**Security notes**
Only the assigned Kurator can make the decision; enforce the
conflict-of-interest check on the backend.
 
---
 
## 5. Campaign Publication
 
**Overview**
A campaign becomes live/public once it passes curation — it can be
published immediately or scheduled for a future date. The Owner can
also unpublish an already-live campaign (e.g. an emergency situation).
 
**Actors**
Owner (sets the schedule, unpublish/republish) · System (executes
automatic publishing per schedule, auto-unpublish resulting from
organization re-curation **[NEW]**)
 
**Preconditions**
`Campaign.status = approved` (for the first publication)
 
**Flow (happy path)**
1. As soon as the campaign is `approved`, the Owner chooses: publish
   now, or set a `publish_at` (a future date/time)
2. Publish now → `Campaign.status = published`,
   `published_at = now`
3. Scheduled → `Campaign.status = scheduled`, `publish_at` is stored
4. The system (a scheduler job) checks periodically; once the current
   time ≥ `publish_at` → `scheduled → published`
**Alternate flows / edge cases**
- **Reschedule**: the Owner may change `publish_at` while the status
  is still `scheduled` (not yet `published`)
- **Unpublish**: the Owner can unpublish a `published` campaign →
  `Campaign.status = unpublished`. The campaign disappears from
  public listings and stops accepting new donations; donation/progress
  data already collected remains fully intact.
- **Republish**: from `unpublished`, the Owner can publish again
  (immediately or rescheduled) → back to `scheduled`/`published`
- **[NEW — RESOLVED]** **Auto-unpublish from organization re-curation**:
  see Business Rules below
**Business rules / validation**
- `publish_at` (if set) must be > now, ideally ≤ the campaign's own
  `deadline`
- Only the Owner can set/change the schedule, or unpublish/republish
- **Manual unpublish requires a reason [RESOLVED — NEW]**: an
  unpublish by the Owner requires filling in `decision_note`, fully
  logged in the Audit Log (actor, timestamp, reason) — consistent with
  force-close and other curation decisions elsewhere in this project.
- **Auto-unpublish from organization re-verification [RESOLVED —
  NEW]**: if a legal/identity field of the organization is edited
  (Feature 1) and `Organization.status` goes back to
  `pending_verification`, ALL `published` campaigns belonging to that
  organization **auto-unpublish**
  (`unpublish_reason = 'organization_re_verification'`,
  system-triggered). This is logged to the Audit Log automatically
  WITHOUT the Owner needing to type a reason (unlike the manual
  unpublish above). Auto-unpublished campaigns are **not**
  auto-republished once the organization is verified again — the Owner
  must manually republish them one by one via the existing Republish
  flow (no new mechanism is needed for this part).
**Data touched**
`Campaign.status`, `publish_at` (nullable), `published_at` (nullable),
`unpublish_reason` **[NEW]** (enum, e.g. `owner_manual` /
`organization_re_verification`), `decision_note` **[NEW]** (nullable,
filled in for manual unpublish, empty for system-triggered)
 
**State transition**
`approved` → `scheduled` → `published` → `unpublished` →
`scheduled`/`published` (republish)
 
**Concurrency & correctness notes**
- The scheduler job must be idempotent — use a conditional update
  (`WHERE status='scheduled' AND publish_at <= now()`) so publishing
  isn't double-triggered if the job runs more than once at the same
  time.
- **[NEW]** Auto-unpublish during organization re-curation: runs in the
  same transaction as updating `Organization.status` back to
  `pending_verification` (Feature 1) — updates all Campaigns
  `WHERE organization_id = :id AND status = 'published'` at once, so
  there's no window where the organization is already
  `pending_verification` but its campaigns are still listed as
  `published`.
**Security notes**
Only the Owner of the relevant organization can change the schedule or
unpublish/republish. Auto-unpublish (system-triggered) doesn't go
through any Owner endpoint at all — it's executed as part of the
organization-edit transaction in Feature 1.
 
**Open questions**
- ~~Does unpublish need a reason/audit trail, given it's an action
  that affects the public visibility of a campaign that has already
  received donations~~ → **resolved: yes, `decision_note` +
  Audit Log required for manual unpublish** **[RESOLVED]**
---
 
## 6. Event Registration
 
**Overview**
An organization representative registers an Event (a lightweight
promotional entity) and links it to one or more of their own
Campaigns — with no curation process.
 
**Actors**
Owner & Staff (mirrors the campaign-draft level — non-sensitive, no
financial data)
 
**Preconditions**
At least one Campaign belonging to that organization has status
`published` or `scheduled`
 
**Flow (happy path)**
1. The representative opens the Event form: name, date/time,
   location, description
2. Selects one or more Campaigns (belonging to the same organization,
   status `published`/`scheduled`) to link
3. Submit → the Event is created, a `CampaignEvent` relation is
   created for each selected campaign
4. The Event is active immediately, with no curation gate
**Business rules / validation**
- Any linked campaign must belong to the same organization as the
  representative creating the event
- Cannot link a campaign with status `draft` / `pending_curation` /
  `rejected`
**Data touched**
- **Create** `Event` (name, datetime, location, description,
  created_by)
- **Create** `CampaignEvent` (event_id, campaign_id) per relation
**Concurrency & correctness notes**
Low risk, no special locking needed.
 
**Security notes**
Only a representative of the organization that owns the campaign can
create a relation to that campaign.
 
---
 
## Open Items Carried Forward
 
- ~~API contract format (affects how these flows get exposed as
  endpoints) — pending from backend tech stack doc~~ → **resolved:
  OpenAPI 3.x spec-first** — see `kencleng-backend-tech-stack.md`
  **[RESOLVED]**
- ~~Notification mechanism (email/in-app) for Admin/Kurator/Organization
  across curation steps~~ → **resolved: extend `notifications.type`
  enum values, dual channel** — see `kencleng-phase0-detail.md`
  Feature 6 and Feature 2 above **[RESOLVED — NEW]**
- ~~Audit-trail granularity for sensitive actions (unpublish,
  rejections)~~ → **resolved: add representative management
  actions to the scope, non-destructive actions intentionally not
  added** — see `kencleng-phase0-detail.md` Feature 9 **[RESOLVED — NEW]**
- ~~Edit-after-verified organization behavior~~ → **resolved: field-level
  split (legal/identity triggers re-curation, operational does not)**
  **[RESOLVED — NEW, see Feature 1]**
- ~~Unpublish audit trail requirement~~ → **resolved: mandatory reason
  + Audit Log** **[RESOLVED — NEW, see Feature 5]**
- ~~`/dashboard/organization/[id]/representatives` — spec not yet
  written~~ → **resolved: direct-add invite, promote/demote &
  removal owner-only with a ≥1-owner guard** — see Feature 1B above
  **[RESOLVED — NEW]**