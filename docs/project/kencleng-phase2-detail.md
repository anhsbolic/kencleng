# Kencleng — Phase 2 (On-Campaign) — Detailed Feature Spec
 
> Status: Draft — detailed at business-rule/flow level; API contract format still pending (see backend tech stack doc).
> Last updated: 2026-07-21
 
## Context
 
This document details each step of Phase 2 (On-Campaign) from
`kencleng-business-process-overview.md`. Unlike Phase 1's approval-chain
flows, Phase 2 is centered on a financial transaction path — the core
correctness/concurrency learning goal for this project lives here.
 
"Donatur discover campaign" (browsing/listing) is intentionally not
detailed as its own feature — it carries no business rule beyond
"only `published` campaigns are publicly listed."
 
Same document format as Phase 1 (Overview / Actors / Preconditions /
Flow / Business rules / Alternate flows / Data touched / State
transition / Concurrency notes / Security notes / Open questions). API
endpoint design is still excluded pending the API contract format
decision.
 
---
 
## 1. Submit Donation
 
**Overview**
A Donatur (guest or registered) makes a donation to a campaign that is
currently `published`, through payment processing that's simulated
asynchronously (there is no real payment gateway in v1).
 
**Actors**
Donatur (guest or registered)
 
**Preconditions**
`Campaign.status = published` — donations can only come in against a
campaign that is currently live.
 
**Flow (happy path)**
1. The Donatur picks a campaign and enters an `amount` (minimum
   **Rp 5,000** **[RESOLVED]** — see the Business Rule below)
2. Chooses a `payment_method` (one of: `transfer`, `debit`,
   `gopay`, `shopeepay`, `ovo`, `qris`) — **[REVISED]** this choice is
   shown even though payment remains fully simulated, so the UX is
   representative and `payment_method` is fully recorded on the row
3. Guest: **optionally** fills in name & email (see the Business Rule
   below — **[REVISED]**, previously required). Registered: optionally
   selects `is_anonymous`
4. Submit → the system creates a `Donation` with `status = pending`
5. The system simulates payment processing asynchronously (a
   **2–5 second** delay **[RESOLVED]** via a background job, mimicking
   a real gateway's callback pattern) — chosen deliberately to practice
   the idempotency/async-settlement pattern even without a real
   gateway
6. Payment succeeds (simulated, **95% of transactions** — see Business
   Rule) → `Donation.status = success`, and atomically increments
   `Campaign.collected_amount`
7. After the increment, the system checks: is
   `collected_amount ≥ max_amount` (if set)? If so → trigger Campaign
   Closure (see Feature 3)
8. Payment fails (simulated, **5% of transactions** **[RESOLVED]**) →
   `Donation.status = failed`, no change to progress
**Business rules / validation**
- The campaign must be `published` when the donation is submitted
- **Minimum donation amount: Rp 5,000 [RESOLVED]** — validated at the
  backend level (not just the UI), preventing donations with an
  `amount` too small to be operationally meaningful
- **Guest is not required to fill in name & email — both are
  independently optional** **[REVISED — previously "required", see
  kencleng-actors-entities.md Business Rule 5]**:
  - A guest may donate without filling in anything, fill in just one,
    or fill in both
  - The donation form shows a **short note** explaining the benefit of
    providing an email: being able to track donation history if they
    later register an account, and receiving notification of the
    campaign's outcome — this is purely a UX nudge, not a requirement
  - **Consequence if `guest_email` is left blank**: that donation
    **can never be claimed** through the Guest Donation Claim flow
    (Feature 4 — email-based matching), and the donor won't receive
    notification of the campaign's outcome (Phase 3, Feature 1) — this
    is a deliberate design decision, not a defect
- A donation is accepted **in full** as long as the campaign is still
  `published` at submit time — **there is no partial-accept/clamping**
  against `max_amount`. As a consequence, the campaign's final total
  can slightly exceed `max_amount` — this is a deliberate design
  decision to avoid clamping complexity, not a race-condition gap
- **Payment simulation parameters [RESOLVED — NEW]**: failure rate
  **5%** (chosen randomly per transaction), processing delay **2–5
  seconds** (random within that range) — chosen so the failure path
  triggers often enough during manual testing (not so rarely it's
  barely exercised), but not so often it disrupts repeated happy-path
  testing
**Alternate flows / edge cases**
- A donation submitted right as the campaign has just closed (a race
  between the submit and the closure trigger) → guarded by
  `WHERE status = 'published'` on the update; if already closed, the
  donation is rejected with a clear message
- Simulated payment fails → the donor submits a new donation (not a
  retry on the same donation)
- Double-submit (clicking submit twice) → needs an idempotency key per
  submission to prevent 2 `Donation` records from 1 user action
- **Guest donates without any email at all** — the donation is still
  valid and processed normally, it's just not eligible for the claim
  flow or notifications (see the Business Rule above) **[NEW]**
**Data touched**
- **Create** `Donation` (campaign_id, donor_user_id [nullable],
  guest_name [nullable, independent], guest_email [nullable,
  independent], amount, `payment_method` [NEW], is_anonymous, status,
  event_id [nullable, if the donation came through an Event context],
  `status_token` **[NEW — RESOLVED]**)
- **Update** `Campaign.collected_amount` (atomic increment)
**State transition**
`Donation.status`: *(none)* → `pending` → `success` | `failed`
 
**Concurrency & correctness notes**
- Incrementing `collected_amount` uses an atomic conditional UPDATE
  (`UPDATE campaign SET collected_amount = collected_amount + :amount
  WHERE id = :id AND status = 'published' RETURNING collected_amount`)
  — Postgres locks that row per-transaction, so donations arriving
  "at the same time" still get processed serially by the database,
  not through a real race
- `RETURNING collected_amount` is used to detect whether this
  transaction is the one that pushed the total past `max_amount`, to
  trigger closure
- Async settlement (simulated) must be idempotent — if the simulated
  callback gets invoked twice due to a bug, the transition is guarded
  by `WHERE status = 'pending'` so it doesn't double-increment
**Security notes**
- The submit-donation endpoint: accessible both publicly (guest) and
  authenticated (registered)
- The simulated payment "callback" endpoint: internal-only, must never
  be exposed externally
- `payment_method` doesn't affect any validation/logic in v1 — it's
  purely recorded data, since payment is entirely simulated **[NEW]**
- **[NEW — RESOLVED]** `status_token`: random, long enough to resist
  brute-force guessing (a donation isn't free of sensitive
  information — the amount & payment method stay private as long as
  the token isn't known to other parties). Doesn't expire, since its
  use is read-only/non-destructive (see Feature 1 Open Questions).
**Open questions**
- ~~Minimum donation amount~~ → **resolved: Rp 5,000** **[RESOLVED]**
- ~~Simulation parameters (payment failure percentage, async delay
  duration)~~ → **resolved: 5% failure rate, 2-5 second delay**
  **[RESOLVED]**
- ~~Guest donation status check without login~~ → **resolved:
  token-in-URL.** A unique `status_token` is returned in the donation
  submit response (usable as
  `/donation/[id]/status?token=...`, bookmarkable by the donor), and
  is also sent via email if `guest_email` was provided (using the
  existing notification channel in `kencleng-phase0-detail.md` Feature
  6). A guest who didn't provide any email at all still receives the
  token directly in the submit response — there's just no secondary
  path to recover that token if it's lost, consistent with Business
  Rule 5 in `kencleng-actors-entities.md` (without an email, the
  donation genuinely has no follow-up path). The token **does not
  expire** — this is a read-only, non-destructive lookup, unlike the
  password-reset token, which changes state and needs a tight time
  window. **[RESOLVED — NEW]**
---
 
## 2. Campaign Progress Update (public display)
 
**Overview**
The public and donors see campaign progress (funds collected vs
target) in a **pull-based, web-only** manner — the value shown is
already kept atomic by Feature 1; this feature is purely about the
query & display.
 
**Actors**
Public (not logged in) · Donatur
 
**Preconditions**
No special gate — progress remains visible even after the campaign is
`closed` (as an archive/history)
 
**Flow (happy path)**
1. The user opens the campaign page
2. The system queries `collected_amount` vs `target_amount`, displays
   a progress bar
3. Displays the public donor list (name & amount), with names hidden
   for `is_anonymous = true`; donors who didn't fill in `guest_name`
   are shown as a generic "Hamba Allah" / "Donatur" label (exact
   wording decided at the FE level) **[NEW — consequence of
   guest_name becoming optional]**
**Business rules / validation**
- `is_anonymous` applies to both guest and registered donors — if
  `true`, the name is hidden from the public list, but still appears
  in the donor's personal history (if registered)
- `guest_email` is never exposed publicly in any form, only
  `guest_name` (if not anonymous, and if provided)
  **[REVISED]**
**Data touched**
Read `Campaign.collected_amount`, `target_amount`; read the `Donation`
list (filtered by `is_anonymous`)
 
**Concurrency & correctness notes**
Read-only, no race — the number being read is already kept atomic in
Feature 1 (Submit Donation).
 
**Security notes**
Public read access, but sensitive fields must be filtered (the guest
email is never sent to the client, names are hidden for anonymous
donations). The `guest_email` field also falls under the PII category
that's masked in the FE for parties actually authorized to view it
(Admin/Kurator) — see `kencleng-actors-entities.md`, PII Handling Note
**[NEW]**.
 
---
 
## 3. Campaign Closure
 
**Overview**
A campaign is automatically closed once its `deadline` is reached, or
`max_amount` is reached (if set) — whichever happens first. The Admin
can also close a campaign manually at any time for an emergency
situation.
 
**Actors**
System (automatic trigger) · Admin (manual force-close)
 
**Preconditions**
`Campaign.status = published`
 
**Flow (happy path)**
- **Trigger via max_amount**: part of the successful-donation
  transaction in Feature 1 — as soon as
  `collected_amount ≥ max_amount`, immediately set
  `Campaign.status = closed`, `closed_reason = 'max_amount_reached'`
- **Trigger via deadline**: a periodic scheduler job checks
  `published` campaigns with `deadline ≤ now()` → sets
  `status = closed`, `closed_reason = 'deadline_reached'`
- **Trigger via Admin force-close**: the Admin triggers closure at any
  time (e.g. fraud is suspected) → sets `status = closed`,
  `closed_reason = 'admin_force_closed'`, requires filling in
  `decision_note` and records `closed_by` (admin_id) for the audit
  trail
**Business rules / validation**
Once `closed` (by whichever trigger): the campaign stops accepting new
donations (guarded in Feature 1), the campaign moves on to Phase 3
(Post-Campaign).
 
**Data touched**
`Campaign.status`, `closed_at`, `closed_reason`, `closed_by`
[nullable, only filled in for force-close], `decision_note` [nullable]
 
**State transition**
`published` → `closed`
 
**Concurrency & correctness notes**
All triggers are equally idempotent via the
`WHERE status = 'published'` guard — if two triggers happen nearly
simultaneously (e.g. a donation nears the max right as the deadline
passes, or an admin force-close happens right as the auto-trigger
runs), whichever succeeds first performs the update, and the other
automatically becomes a no-op because `status` has already changed
from `published`.
 
**Security notes**
Force-close can only be triggered by the Admin role.
 
**Open questions**
- The consequences of force-close on funds already collected (whether
  they're still disbursed as with a normal closure, or there's an
  additional hold/review process) — discussed in the Phase 3 detail
  (disbursement)
---
 
## 4. Guest Donation Claim
 
**Overview**
A permanent page (not a one-time prompt) accessible at any time from a
user's account, showing guest donations that match that user's
verified email, to be claimed manually one at a time.
 
**Actors**
User (registered, verified email)
 
**Preconditions**
The user is logged in and their email is already verified
 
**Flow (happy path)**
1. The user opens the "Claim Past Donations" page from their account,
   at any time
2. The system queries `Donation` with `guest_email = user.email` AND
   `donor_user_id IS NULL` — **only donations with a non-null
   `guest_email` are eligible** (a guest donation with no email at
   all will never show up here) **[REVISED — restated for clarity]**
3. The system shows the user the list of candidate donations
4. The user confirms which donations to claim, one at a time
5. Per confirmed item: `donor_user_id = user.id`,
   `claimed_at = now`
6. The original `guest_name`/`guest_email` snapshot remains stored;
   historical public display (anonymous or not) doesn't change
   retroactively
**Business rules / validation**
- No auto-claim — the user must explicitly confirm each item
- The user may skip some or all of the candidates shown
- Because this page is permanent (not one-time), guest donations made
  **after** the user already has an account (e.g. they forgot to log
  in while donating) can still show up & be claimed at any time
**Data touched**
Update `Donation.donor_user_id`, `claimed_at`
 
**Concurrency & correctness notes**
The update is guarded with `WHERE donor_user_id IS NULL` during the
claim process — preventing the same donation from being claimed twice
if (an edge case) two different users happen to share the same email.
 
**Security notes**
A user can only view & claim donations whose `guest_email` matches
their own **already-verified** email — there must be no path to claim
a donation using an email whose ownership hasn't been verified.
 
---
 
## Open Items Carried Forward
 
- API contract format — pending from the backend tech stack doc
- ~~Minimum donation amount & payment simulation parameters (failure
  rate, delay)~~ → **resolved [RESOLVED]** — see Feature 1
- Consequences of force-close on funds already collected — continued
  in the Phase 3 detail
- Notification mechanism for donors (e.g. notification when a campaign
  they supported closes) — not yet designed
- ~~Guest donation status check mechanism (token-based vs other)~~ →
  **resolved: token-in-URL, does not expire** **[RESOLVED — NEW, see
  Feature 1]**
- **Display label for guest donors who didn't fill in a name (e.g.
  "Hamba Allah"/"Donatur") — copy/wording detail left to the FE**
  **[NEW]**