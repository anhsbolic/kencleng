# Kencleng — Business Process Overview
 
> Status: Draft — high-level phase breakdown agreed; per-phase detail
> now complete in `kencleng-phase0-detail.md` through
> `kencleng-phase3-detail.md`.
> Last updated: 2026-07-24 (doc-sync pass — all Open Items resolved
> elsewhere, cross-referenced here)
 
## Context
 
This document records the end-to-end business process for a Campaign
on Kencleng, broken into three phases. It complements the actors/entities
doc — this doc describes the *flow*, that doc describes *who/what* is
involved.
 
**Naming note**: an earlier draft of this discussion used the term
"event donasi" to mean what is now formally called **Campaign**. This
doc uses `Campaign` and `Event` as distinct entities throughout (see
`kencleng-actors-entities.md`):
- **Campaign** — the core fundraising unit. Owns `target_amount`,
  optional `max_amount`, `deadline`, curation status, and the
  disbursement process. Always belongs to exactly one Organisasi.
- **Event** — a lightweight promotional entity (name, date/time,
  location, description). No curation, no financial data of its own.
  Many-to-many with Campaign (one campaign can be promoted at multiple
  events; one event can promote multiple campaigns).
## Phase 1 — Pre-Campaign
 
| # | Step | Actor |
|---|---|---|
| 1 | Organisasi registration | Organisasi (prospective Owner) |
| 2 | Organisasi curation (KYC / legality verification) | Kurator |
| 3 | Campaign registration — set `target_amount`, `max_amount` (optional), `deadline` | Owner (Organisasi) |
| 4 | Campaign curation | Kurator |
| 5 | Campaign publication | System (post-curation) |
| 6 | *(Parallel, optional)* Event registration + link to the related Campaign | Organisasi / event organizer |
 
Note: step 6 has no curation gate — Event is intentionally lightweight
and carries no financial data, so it doesn't need the same gatekeeping
as a Campaign.
 
## Phase 2 — On-Campaign
 
| # | Step | Actor | Notes |
|---|---|---|---|
| 1 | Donor discovers a Campaign — directly, or via an associated Event | Donor (guest or registered) | |
| 2 | Submit a donation → payment processing | Donor | |
| 3 | Real-time progress tracking (amount collected vs `target_amount`) | System | **Pull-based, web-only** for v1 — no push notifications. Publicly viewable on the campaign page. Concurrency-safe update is a core correctness concern here. |
| 4 | Campaign closure | System | Trigger: `deadline` reached, **or** `max_amount` reached (if set) — whichever comes first |
 
## Phase 3 — Post-Campaign
 
| # | Step | Actor | Notes |
|---|---|---|---|
| 1 | Progress report & final collected-donation results | System / Organisasi | Unlike phase 2's progress view, this is **actively distributed** to donors (e.g. notification) in addition to being publicly viewable as a summary page |
| 2 | Fund disbursement request | Owner (Organisasi) | Sensitive action — Owner-only per business rules in the actors/entities doc |
| 3 | Fund disbursement to the Organisasi | System | Fully simulated (no real payment rail/bank transfer) — see `kencleng-phase3-detail.md` Feature 3 **[RESOLVED]** |
| 4 | Fund-usage report | Organisasi (Owner) | |
| 5 | Fund-usage report verification | Kurator | Must recuse if also a representative of the organisasi (conflict of interest, per the actors/entities doc) |
 
## Business Rules Carried Over (from actors/entities doc)
 
- Keep-it-all model: funds collected are disbursed even if `target_amount`
  is not reached by `deadline` — no refund mechanism in v1.
- If `max_amount` is set and reached before `deadline`, the campaign
  closes early. If not set, overfunding is allowed until `deadline`.
- Kurator must recuse from curating an Organisasi, curating a Campaign,
  or verifying a fund-usage report for any organisasi where they are
  also a representative.
## Open Items — Needs Further Discussion (per-phase deep dive)
 
All resolved — kept here for history, strikethrough per project
convention:
 
1. ~~**Phase 1**: Organisasi KYC document requirements; Campaign
   registration form/fields beyond target/max/deadline.~~ → **resolved**:
   KYC = Akta Notaris, SK Kemenkumham, NPWP (format-validated, unique),
   Izin PUB (optional), reviewed by an Admin-assigned, conflict-free
   Kurator — see `kencleng-phase1-detail.md` Feature 1 & 2. Campaign
   fields extended with `category` (required enum), `location`
   (optional free-text), `beneficiary_description` (optional
   free-text) — see `kencleng-phase1-detail.md` Feature 3.
   **[RESOLVED]**
2. ~~**Phase 2**: payment method(s) / gateway (real sandbox vs fully
   simulated — still open from backend tech stack doc); guest vs
   registered donation flow details; concurrency strategy for
   progress updates.~~ → **resolved**: payment fully simulated (5%
   failure rate, 2–5 second random delay, 6 methods —
   transfer/debit/gopay/shopeepay/ovo/qris), no real gateway
   integration. Guest vs registered flow detailed in
   `kencleng-actors-entities.md` Business Rules 5–7 (optional
   guest_name/guest_email, claim flow). Concurrency strategy for
   `collected_amount` updates detailed in
   `kencleng-phase2-detail.md`. **[RESOLVED]**
3. ~~**Phase 3**: disbursement mechanism (single lump-sum vs partial);
   what "fund-usage report" actually contains (structured fields vs
   free-form + attachments); consequences of a failed/rejected
   fund-usage report verification.~~ → **resolved**: lump-sum only (no
   partial disbursement) — see `kencleng-phase3-detail.md` Feature 2.
   Fund-usage report = structured breakdown per expense category
   (name, amount, description, attachment) with strict reconciliation
   against `disbursed_amount` — see Feature 4. Rejected verification has
   **no special consequence** beyond the normal resubmit cycle (the
   `has_overdue_report` flag is governed purely by submission
   deadline, not verification outcome) — see Feature 5.
   **[RESOLVED]**
## Not Yet Discussed
 
- ~~Detailed field-level design per step (forms, validation rules)~~ →
  **resolved**, see `kencleng-phase0-detail.md` through
  `kencleng-phase3-detail.md` for business-rule-level detail per
  feature. Pure UI-implementation detail (empty/loading/error states,
  exact form copy) is explicitly deferred to the implementation phase
  — see `kencleng-roadmap-next-steps.md` Open Items. **[RESOLVED]**
- ~~API contract per step~~ → **resolved: OpenAPI 3.x spec-first**,
  see `kencleng-backend-tech-stack.md` "API Contract & Codegen"
  **[RESOLVED]**
- ~~ERD / schema design~~ → **resolved**, see `kencleng-erd.md`
  **[RESOLVED]**