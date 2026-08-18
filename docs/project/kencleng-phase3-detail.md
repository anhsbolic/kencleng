# Kencleng — Phase 3 (Post-Campaign) — Detailed Feature Spec
 
> Status: Draft — detailed at business-rule/flow level. API contract format resolved (OpenAPI 3.x spec-first, see `kencleng-backend-tech-stack.md`); endpoint-level detail still excluded from this doc pending actual spec authoring.
> Last updated: 2026-07-24 (rev 2 — repeated-rejection consequence clarified, notification types for disbursement decisions added, cross-references synced)
 
## Context
 
This document details each step of Phase 3 (Post-Campaign) from
`kencleng-business-process-overview.md`. Unlike Phase 2's
concurrency-heavy transaction path, Phase 3 returns to an
approval-chain pattern similar to Phase 1 — but the stakes are higher,
since this is where money actually leaves the platform to an
Organization.
 
Same document format as Phase 1/2 (Overview / Actors / Preconditions /
Flow / Business rules / Alternate flows / Data touched / State
transition / Concurrency notes / Security notes / Open questions).
Endpoint-level detail (concrete paths/methods) is intentionally
excluded from this doc — that lives in `api/openapi.yaml` per the
resolved API contract format decision
(`kencleng-backend-tech-stack.md`, "API Contract & Codegen"), not
duplicated here.
 
---
 
## 1. Progress Report & Final Collected-Donation Results
 
**Overview**
Once a campaign is `closed`, the system produces a summary of the
donation results and actively distributes it to donors, while also
keeping it published as a public/archival summary page. **The
Organization (Owner) can add a narrative to this report — this feature
is officially in v1** **[REVISED — previously an open question]**.
 
**Actors**
System (generates & distributes the report automatically) · Organization
(Owner — adds a narrative, optional)
 
**Preconditions**
`Campaign.status = closed`
 
**Flow (happy path)**
1. As soon as the campaign is `closed`, the system automatically
   generates a summary: `collected_amount`, number of unique donors,
   campaign duration, `closed_reason`
2. The system sends a notification to donors through **two
   channels**: an in-app notification (for registered donors) and
   email (for all donors, both registered and guest, using the stored
   `guest_email`)
3. The summary page remains publicly accessible as an archive, even
   long after the campaign has closed
4. **[NEW]** The organization's Owner can open this report page from
   their dashboard, and add/edit the narrative field
   (`report_narrative`) — freely, at any time after the campaign is
   `closed`, with no deadline or approval gate for this narrative
5. **[NEW]** If the narrative is filled in/changed, the public summary
   page also displays that narrative below the automatic figures
**Business rules / validation**
- The email sender in v1 is simply **fake/logged** — no need for a
  real SMTP integration, it's enough to record it as a simulated send
  (mirroring the payment-simulation principle in Phase 2)
- **[NEW]** The narrative is **optional and does not go through a
  curation/verification process** — unlike the Fund-Usage Report
  (Feature 4), which must be verified by a Kurator. Reason: the
  narrative here is purely additional story/context (e.g. thank-you
  message, program impact), not financial accountability — so it
  doesn't need the same strict gate
- **[NEW]** The Owner may edit the narrative repeatedly after it's
  published (it isn't locked after the first submission) — consistent
  with its non-financial/non-accountability nature
- **Narrative format & length [RESOLVED — NEW]**: **Markdown**,
  maximum **5000 characters**. The backend only stores the raw
  Markdown text (purely a length validation, no business-logic
  parsing) — rendering to HTML and sanitization happen on the frontend
  at render time (see Security notes below)
**Alternate flows / edge cases**
- **[NEW]** The Owner never fills in the narrative → the summary page
  still displays normally, just with the system's automatic figures
  (the narrative really is optional, not required)
**Data touched**
Read `Campaign` (collected_amount, closed_reason, etc.), read
`Donation` (aggregate count of unique donors); create notification/
email log records for audit/testing; **[NEW]** create/update
`Campaign.report_narrative` (nullable text field, max 5000 characters,
Markdown, filled in/edited by the Owner)
 
**State transition**
No formal state transition for the narrative (it's not a status
field, just a nullable text field that can be edited freely)
 
**Security notes**
- `guest_email` is used internally for sending notifications only,
  never exposed publicly.
- **[NEW]** Only the Owner of the organization that owns the campaign
  can fill in/edit the narrative — a `staff` representative can view
  the report page but cannot edit the narrative (consistent with
  Business Rule 4 in `kencleng-actors-entities.md`: sensitive actions
  are Owner-only — even though the narrative itself isn't a financial
  action, it's still treated as an official representation of the
  organization to the public, so it's held to the same level as other
  Owner-only actions)
- **Markdown sanitization required at render time [RESOLVED — NEW,
  SECURITY-CRITICAL]**: since the narrative is displayed on a public
  page, the Markdown **must not** be stored/rendered as raw HTML
  without sanitization (XSS risk). The backend only stores raw
  Markdown text; the frontend must render through a library that has
  built-in sanitization (e.g. `react-markdown` + `rehype-sanitize`),
  **not** `dangerouslySetInnerHTML` on the output of a manual
  markdown-to-html conversion. This is consistent with the "Next.js
  has no business logic" boundary — sanitizing at render time is a
  presentation concern, not business logic
**Open questions**
- ~~Can the Organization add a narrative/story to this report, or is it
  purely automatic figures from the system?~~ → **resolved: yes, it's
  in v1** **[REVISED]**
- ~~Does the narrative need a character-length limit, or does it
  support rich text/markdown, or plain text only?~~ → **resolved:
  Markdown, max 5000 characters** **[RESOLVED]**
---
 
## 2. Fund Disbursement Request
 
**Overview**
The Owner submits a request to disburse funds from a campaign that has
already been `closed`. It needs Admin approval before the funds are
actually disbursed.
 
**Actors**
Owner (requests) · Admin (approve/reject)
 
**Preconditions**
`Campaign.status = closed`
 
**Flow (happy path)**
1. The Owner submits a disbursement request on a `closed` campaign
2. The system creates a `DisbursementRequest` (`status = pending`,
   `requested_amount = collected_amount` in full — a lump-sum model,
   disbursed once) → **the system notifies the Admin** (`type =
   admin_new_curation_item`, dual channel — reusing the same type
   used for other curation queues **[RESOLVED — NEW]**)
3. The Admin reviews the request → **approve** (proceeds to Feature
   3) or **reject** (`decision_note` required) → **the Owner is
   notified of the decision** (`type = disbursement_approved` /
   `disbursement_rejected`, dual channel **[RESOLVED — NEW]**)
**Alternate flows / edge cases — resubmit after rejection**
4. If rejected, the Owner can revise (e.g. administrative completeness/
   bank account info) & submit a new request
5. The old request is kept as history, the new request enters the
   review queue from the start
 
**Business rules / validation**
- Only the Owner can make the request (not staff)
- Only one **active** `DisbursementRequest` (`status = pending` or
  `approved`/`disbursed`) is allowed per campaign at a time — prevents
  double-disbursement
**Data touched**
Create `DisbursementRequest` (campaign_id, requested_by,
requested_amount, status, reviewed_by, decision_note, decided_at)
 
**State transition**
`DisbursementRequest.status`: *(none)* → `pending` → `approved`
| `rejected` → *(revision)* `pending` (repeat cycle, new request)
 
**Security notes**
Request: Owner-only. Approve/reject: Admin-only.
 
---
 
## 3. Fund Disbursement to the Organization
 
**Overview**
Once the Admin approves, the system simulates disbursing the funds to
the organization — no real bank account/transfer, per the sandbox
principle (mirroring the payment simulation in Phase 2).
 
**Actors**
System (executes the disbursement simulation)
 
**Preconditions**
`DisbursementRequest.status = approved`
 
**Flow (happy path)**
1. As soon as it's `approved`, the system triggers a simulated fund
   disbursement (a short delay, recorded as "transferred")
2. `DisbursementRequest.status = disbursed`, `disbursed_at = now`
3. The Owner & Organization receive a notification that the funds have
   been disbursed (dual channel, same as Feature 1)
**Data touched**
Update `DisbursementRequest.status = disbursed`, `disbursed_at`
 
**State transition**
`approved` → `disbursed`
 
**Concurrency & correctness notes**
This transition is already gated by the prior Admin approval, so the
race risk is low — a `WHERE status = 'approved'` guard on the update
is enough for idempotency.
 
---
 
## 4. Fund-Usage Report
 
**Overview**
The Organization (Owner) must submit a fund-usage accountability report
once disbursement is complete — a structured format per expense
category, not just free-form narrative. There's a submission deadline,
and a platform-level consequence if that deadline is missed
**[REVISED — previously an open question]**.
 
**Actors**
Owner
 
**Preconditions**
`DisbursementRequest.status = disbursed`
 
**Flow (happy path)**
1. The Owner opens the fund-usage report form
2. Fills in a breakdown per expense category: category name, amount,
   description, supporting attachment (receipt/photo/document) per
   item
3. Submit → a `FundUsageReport` is created
   (`status = pending_verification`) along with child
   `FundUsageReportItem` records per category
4. **[NEW]** As soon as the submission succeeds, if this organization
   currently has the "overdue report" flag (see the Business Rule
   below), that flag is immediately cleared automatically — there's no
   need to wait for the Kurator's verification to finish
**Business rules / validation**
- **Reconciliation rule [RESOLVED — NEW]**: the sum of all
  `FundUsageReportItem.amount` values **must exactly match (strict
  match)** the `disbursed_amount` — no tolerance for a discrepancy. If
  there's a legitimate cost that doesn't fall under a program-expense
  category (e.g. a bank transfer admin fee), that cost must still be
  included as its own line-item breakdown entry (e.g. "Bank
  administration fee"), rather than being left as a gap the system
  tolerates. Submission is rejected (backend validation) if the total
  doesn't match
- **Submission deadline [RESOLVED — NEW]**: the report must be
  submitted no later than **30 days** after `disbursed_at`
- **Consequence of missing the deadline [RESOLVED — NEW]**: the
  related organization gets flagged with `has_overdue_report` — while
  this flag is active, the organization **cannot create a new
  campaign** (see `kencleng-phase1-detail.md` Feature 3, Campaign
  Registration — an additional validation needed to be added there
  when that detail was revised). Campaigns that are already
  `published`/running are **not affected** — this flag only blocks
  creating a new campaign, not unpublishing/pausing an existing one
  - The flag is **set** automatically by a scheduler job (part of the
    existing in-process scheduler — see
    `kencleng-backend-tech-stack.md`) as soon as it detects a
    `FundUsageReport` that doesn't exist yet / hasn't been submitted
    even though `disbursed_at + 30 days` has already passed
  - The flag is **cleared** automatically as soon as the Owner submits
    the report (step 4 in the Flow), **without waiting** for the
    Kurator's verification result (Feature 5) — once the report enters
    the review queue, the Owner's submission responsibility is
    considered fulfilled; if that report is later rejected by the
    Kurator, the normal revision cycle (Feature 5) applies, with no
    need for an additional flag
  - Any change to this flag (set or cleared) is logged in the Audit
    Log — see `kencleng-phase0-detail.md` Feature 9
**Data touched**
- **Create** `FundUsageReport` (campaign_id, submitted_by, status,
  submitted_at)
- **Create** `FundUsageReportItem` (report_id, category, amount,
  description, attachment_url) — one or more per report
- **[NEW]** Update `Organization.has_overdue_report` (boolean, set by
  the scheduler job, cleared automatically on submission)
**State transition**
`FundUsageReport.status`: *(none)* → `pending_verification`
 
**Concurrency & correctness notes**
**[NEW]** Setting/clearing `has_overdue_report`: the scheduler job sets
the flag using the guard `WHERE disbursed_at + interval '30 days' <
now() AND NOT EXISTS (a report has already been submitted)`; clearing
happens directly within the same transaction as creating the
`FundUsageReport` on submission, so there's no window where the
report has already been submitted but the flag is still active.
 
**Security notes**
Only the Owner of the relevant organization can submit a report for that
campaign.
 
**Open questions**
- ~~Must the breakdown total exactly match `disbursed_amount` (strict
  reconciliation), or is a small tolerance/discrepancy allowed?~~ →
  **resolved: strict match** **[RESOLVED]**
- ~~Is there a mandatory deadline to submit this report after
  disbursement (e.g. a maximum of 30 days)?~~ → **resolved: 30 days**
  **[RESOLVED]**
- ~~Platform-level consequence if the fund-usage report is rejected/
  late repeatedly~~ → **resolved: `has_overdue_report` flag, blocks
  new campaign creation, clears automatically on submission**
  **[RESOLVED]**
---
 
## 5. Fund-Usage Report Verification
 
**Overview**
A Kurator (assigned by the Admin, mirroring the other curation
patterns) reviews the fund-usage report — must recuse if they're a
representative of the same organization.
 
**Actors**
Admin (assigns) · Kurator (review & decision)
 
**Preconditions**
`FundUsageReport.status = pending_verification`
 
**Flow (happy path)**
1. The report enters `pending_verification` → the system notifies the
   Admin (`type = admin_new_curation_item`, dual channel **[RESOLVED —
   NEW]**)
2. The Admin assigns one Kurator → creates a
   `FundUsageReportVerificationAssignment` (`decision = pending`) →
   **the Kurator is notified of the assignment** (`type =
   kurator_assigned`, dual channel **[RESOLVED — NEW]**)
3. The Kurator reviews the breakdown & supporting evidence
4. **Approve** → `FundUsageReport.status = verified`
   or **Reject** (`decision_note` required) →
   `FundUsageReport.status = rejected` → **the Owner is notified of
   the decision** (`type = fund_usage_report_verified` /
   `fund_usage_report_rejected`, dual channel **[RESOLVED — NEW]**)
**Alternate flows / edge cases — resubmit after rejection**
5. The Owner revises a `rejected` report & resubmits →
   `FundUsageReport.status` goes back to `pending_verification`
6. A new assignment cycle starts (the old assignment is kept as
   history)
 
**Business rules / validation**
- A Kurator who is also a representative of the organization that owns
  the campaign may not be assigned (conflict of interest — same as
  organization/campaign curation)
- Only one active assignment per report at a time
- **[NEW]** A rejection here does **not** re-trigger the
  `has_overdue_report` flag — that flag is purely about submission
  timeliness (Feature 4), not the verification outcome. A report that
  gets rejected and then revised & resubmitted is still considered
  "on time" as long as the first submission happened before the
  30-day deadline. **[RESOLVED — NEW]** This also fully answers "the
  platform-level consequence if the fund-usage report is rejected
  repeatedly (as distinct from being late)": **there is no special
  consequence** beyond the normal, already-existing resubmit cycle —
  an organization is free to be rejected & resubmit as many times as
  needed with no additional penalty, as long as it stays within the
  30-day window from the first submission. This is intentionally left
  penalty-free so an Owner isn't afraid to submit an honest report
  that isn't yet perfect, and the revise-resubmit process stays a
  normal path rather than something to be avoided
**Data touched**
- **Update** `FundUsageReport.status` → `verified` / `rejected`
- **Create** `FundUsageReportVerificationAssignment` (report_id,
  kurator_id, assigned_by, assigned_at, decision, decision_note,
  decided_at)
**State transition**
`FundUsageReport.status`: `pending_verification` → `verified` |
`rejected` → *(revision)* `pending_verification` (repeat cycle)
 
**Concurrency & correctness notes**
Same as organization/campaign curation — one active assignment per
report prevents a race between Kurator.
 
**Security notes**
Only the assigned Kurator can make the decision; enforce the
conflict-of-interest check on the backend.
 
---
 
## Open Items Carried Forward
 
- ~~API contract format — pending from the backend tech stack doc~~ →
  **resolved: OpenAPI 3.x spec-first** — see
  `kencleng-backend-tech-stack.md` **[RESOLVED]**
- ~~Reconciliation rule for `FundUsageReportItem` vs
  `disbursed_amount` (strict match vs tolerance)~~ → **resolved:
  strict match** **[RESOLVED]**
- ~~Deadline to submit the fund-usage report after disbursement~~ →
  **resolved: 30 days** **[RESOLVED]**
- ~~Can the organization add a narrative to the donation-progress report
  (Feature 1), or is it purely automatic~~ → **resolved: yes, it's in
  v1, Owner-only, with no curation gate** **[REVISED]**
- ~~Length/format limit for the progress-report narrative (plain text
  vs rich text)~~ → **resolved: Markdown, max 5000 characters,
  sanitization required at render time** **[RESOLVED]**
- ~~Platform-level consequence if the fund-usage report is rejected
  repeatedly~~ → **resolved: no special consequence beyond the normal
  resubmit cycle — the deadline is still counted from the original
  `disbursed_at`, not reset on each resubmit** — see Feature 5 above
  **[RESOLVED — NEW]**
- ~~Needs a cross-reference update in `kencleng-phase1-detail.md`
  Feature 3 (Campaign Registration) to add a new validation: an
  organization with `has_overdue_report = true` cannot submit a new
  campaign~~ → **resolved: the cross-reference has been added** —
  see `kencleng-phase1-detail.md` Feature 3, Preconditions &
  Business Rules **[RESOLVED — NEW]**
- ~~Needs to add `Organization.has_overdue_report` to the entity/field
  list in the ERD (Step 3 roadmap)~~ → **resolved: already present in
  `kencleng-erd.md`** (`organizations.has_overdue_report`, boolean,
  with a partial index) **[RESOLVED]**
- ~~Notification mechanism for the Admin (new queue item)/Kurator
  (assignment)/Owner (disbursement & fund-usage-report decisions)~~ →
  **resolved: extend the `notifications.type` enum, dual channel** —
  see Feature 2 & Feature 5 above and `kencleng-phase0-detail.md`
  Feature 6 **[RESOLVED — NEW]**