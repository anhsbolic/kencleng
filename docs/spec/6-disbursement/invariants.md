# Domain Invariant — disbursement

> File: `docs/spec/disbursement/invariants.md`
> Status: draft — authored directly against `api/openapi/disbursement.yaml` 2026-08-20
> Last updated: 2026-08-20

## Domain summary

`disbursement` owns the post-campaign lifecycle: the public campaign
report/archive, disbursement requests, the simulated fund transfer,
fund-usage reports, and fund-usage report verification. Covers
`disbursement_requests`, `disbursement_request_logs`,
`fund_usage_reports`, `fund_usage_report_items`,
`fund_usage_report_item_attachments`,
`fund_usage_report_verification_assignments`,
`fund_usage_report_logs`. This domain **fulfills** the forward
reference `organization/invariants.md` made for INV-organization-13
(`has_overdue_report`'s set/clear trigger) — see INV-disbursement-07/08
below.

## Assumed defaults (not blocking, flagged for confirmation)

Two endpoints (`GET /disbursement-requests/{id}`, `GET
/fund-usage-reports/{reportId}`) document a possible `403` with no
example payload — meaning some access check exists, just not spelled
out. This spec assumes the same access pattern used consistently
elsewhere in this project (representative of the owning organization,
assigned/historical Kurator, or Admin), rather than blocking on it —
low-severity ambiguity, not a genuine conflict like the ones surfaced
in `organization`/`campaign`. Flagged in `threat-model.md` for
confirmation during implementation.

## Invariants

### INV-disbursement-01: Campaign report is a public, permanent archive for closed campaigns only

- **Statement**: `GET /campaigns/{campaignId}/report` is available
  (no auth) for any campaign with `status = 'closed'`, indefinitely —
  no expiry, no takedown mechanism. For a non-`closed` campaign,
  `409`.
- **Holds after operations**: report view.
- **Verification**: Confirmed — `409` "Campaign is not yet closed."
  Test: view report for a `closed` campaign (succeeds, no auth), for
  a `published` campaign (`409`).

### INV-disbursement-02: Report narrative is owner-only, ungated, freely editable, length-bounded

- **Statement**: `PATCH /campaigns/{campaignId}/report-narrative` is
  owner-only (not `staff` — confirmed explicit, consistent with
  Business Rule 4's "official representation of the organization"
  framing even though the narrative itself isn't a financial action).
  No curation/verification gate — editable any number of times, any
  time after `closed`, no deadline. `report_narrative` max 5000
  characters, raw Markdown storage only — the backend does no
  HTML rendering or sanitization; that's a frontend responsibility at
  render time (`react-markdown` + `rehype-sanitize` or equivalent,
  **never** `dangerouslySetInnerHTML` on unsanitized output —
  explicitly flagged security-critical in
  `kencleng-phase3-detail.md`).
- **Holds after operations**: narrative create/edit.
- **Verification**: Confirmed — `422` over 5000 chars, `403` for
  `staff`/non-representative, `409` if campaign not yet `closed`.
  Test: `staff` attempt → `403`. Repeated edits after the first
  succeed (not locked). Empty string clears the narrative (confirmed
  — "may be set to an empty string to clear it").

### INV-disbursement-03: Disbursement request creation is owner-only, lump-sum, one-active-per-campaign

- **Statement**: `POST /campaigns/{campaignId}/disbursement-requests`
  is owner-only. Requires `Campaign.status = 'closed'`.
  `requested_amount` is always set to `Campaign.collected_amount` at
  request time — no partial/custom amount, no request body at all.
  Only one **active** request (`status` ∈ {`pending`, `approved`,
  `disbursed`}) is allowed per campaign at a time (DB-enforced,
  `ux_disbursement_requests_one_active`) — a `rejected` request does
  **not** count as active, so a fresh request after rejection is
  allowed.
- **Holds after operations**: request creation.
- **Verification**: Confirmed — `409 campaign-not-closed`, `409
  disbursement-already-active`, `403` for `staff`. Test: two
  concurrent request attempts on the same campaign — DB unique index
  guarantees at most one succeeds, translate the constraint violation
  to a clean `409`.

### INV-disbursement-04: Disbursement decision is Admin-only; rejection requires a note

- **Statement**: `POST
  /disbursement-requests/{disbursementRequestId}/decision` is
  Admin-only. `decision = 'rejected'` requires `decision_note`. Only
  valid while `status = 'pending'`.
- **Holds after operations**: decision.
- **Verification**: Confirmed — `422` on missing note for reject,
  `409` if not `pending`. Test: `staff`/Owner/Kurator caller → `403`.

### INV-disbursement-05: `approved → disbursed` is a system-triggered, idempotent transition

- **Statement**: On approval, a simulated fund transfer (short delay)
  is triggered internally — `status → disbursed`, `disbursed_at =
  now()` happens as a follow-up system transition, not synchronously
  in the decision response. Guarded by `WHERE status = 'approved'` —
  idempotent if the job somehow runs twice. The decision endpoint's
  response reflects `status = 'approved'` immediately; the client
  polls `GET /disbursement-requests/{id}` to observe the subsequent
  `disbursed` transition (confirmed explicit in the endpoint
  description).
- **Holds after operations**: the internal disbursement-execution
  process (no HTTP endpoint of its own).
- **Verification**: Test — invoke the transition twice for the same
  request; assert `disbursed_at` set exactly once, second invocation
  is a no-op.

### INV-disbursement-06: Fund-usage report submission requires strict reconciliation and clears `has_overdue_report`

- **Statement**: `POST
  /disbursement-requests/{disbursementRequestId}/fund-usage-reports`
  is owner-only, requires `DisbursementRequest.status = 'disbursed'`.
  The sum of all `items[].amount` **must exactly equal**
  `requested_amount` — no tolerance; a legitimate incidental cost
  (e.g. a bank transfer fee) must be its own line item, not treated as
  an acceptable gap. On successful submission, if the owning
  organization's `has_overdue_report = true`, it is cleared **in the
  same transaction** as the report creation (INV-disbursement-08).
- **Holds after operations**: report submission.
- **Verification**: Confirmed — `422` on reconciliation mismatch or
  other field validation, `409` if not yet `disbursed`, `403` for
  `staff`. Test: items summing to anything other than exactly
  `requested_amount` (over or under, by even a small amount) → `422`.
  Submission on an organization with `has_overdue_report = true`
  clears it in the same transaction — no window where the report
  exists but the flag is still set.

### INV-disbursement-07: `has_overdue_report` is set by a scheduler job when a report is missing 30 days after disbursement (fulfills — `organization` domain forward reference)

- **Statement**: A scheduler job periodically sets
  `organizations.has_overdue_report = true` using the guard `WHERE
  disbursed_at + interval '30 days' < now() AND NOT EXISTS (a
  fund_usage_reports row already exists for this
  disbursement_request)`. This is the concrete implementation of the
  set-side of **INV-organization-13** in
  `docs/spec/organization/invariants.md` — that entry declared the
  *effect* (blocks campaign creation), this entry declares the
  *trigger* (owned here, since `disbursement` owns the deadline
  logic).
- **Holds after operations**: the scheduler job.
- **Verification**: Test — a disbursement past `disbursed_at + 30
  days` with no submitted report gets `has_overdue_report = true` on
  the next scheduler run. A disbursement with a report already
  submitted (even if only just before the deadline) never gets
  flagged. Scheduler run twice near-simultaneously doesn't
  double-log or error (idempotent `UPDATE`, same pattern as every
  other scheduler in this project).

### INV-disbursement-08: "On time" is determined by the *first* submission only, not affected by later rejection/resubmission

- **Statement**: Once a `fund_usage_reports` row exists for a
  disbursement (i.e. the Owner submitted *something* before the
  30-day deadline), that submission satisfies the deadline
  permanently — even if the Kurator later rejects it and the Owner
  resubmits well past the original 30-day window. `has_overdue_report`
  is never re-set for the same disbursement based on rejection timing;
  only the scheduler (INV-disbursement-07), checking for the complete
  *absence* of any report, can set it.
- **Holds after operations**: fund-usage report rejection, resubmission
  (both — neither re-evaluates the deadline).
- **Verification**: Test — submit a report 29 days after disbursement
  (on time, clears any existing flag), get it rejected, resubmit 40
  days after disbursement (well past 30 days) — assert
  `has_overdue_report` is **not** set as a result of this late
  resubmission.

### INV-disbursement-09: Fund-usage report rejection carries no additional penalty beyond the normal resubmit cycle

- **Statement**: Rejection does not trigger `has_overdue_report`, does
  not limit the number of resubmission cycles, and does not otherwise
  penalize the organization — confirmed explicit in
  `kencleng-phase3-detail.md`: "intentionally left penalty-free so an
  Owner isn't afraid to submit an honest report that isn't yet
  perfect."
- **Holds after operations**: rejection, resubmission (any number of
  cycles).
- **Verification**: Test — reject and resubmit a report 3+ times;
  assert no cumulative penalty/flag/limit is applied at any point.

### INV-disbursement-10: Kurator conflict-of-interest recusal (fund-usage verification)

- **Statement**: A user may not be assigned as Kurator for a
  `fund_usage_report_verification_assignment` targeting a report whose
  campaign's organization they're also a representative of (any
  `level`). Same pattern as INV-organization-06/INV-campaign-05.
- **Holds after operations**: `POST
  /fund-usage-reports/{reportId}/curation/assign`.
- **Verification**: Confirmed in the endpoint description. Test: same
  shape as the `organization`/`campaign` equivalents.

### INV-disbursement-11: Only one active verification assignment per report

- **Statement**: `fund_usage_report_verification_assignments` may have
  at most one row with `decision = 'pending'` per
  `fund_usage_report_id`.
- **Holds after operations**: assignment creation.
- **Verification**: DB-level (`ux_fund_usage_verif_one_pending`).
  Confirmed — `409` "Conflict of interest, or an active assignment
  already exists" (same combined-error pattern as `organization`/
  `campaign` — worth splitting into distinct `type`s at implementation
  time).

### INV-disbursement-12: Only the assigned Kurator can decide; rejection requires a note

- **Statement**: `POST
  /fund-usage-reports/{reportId}/curation/decision` succeeds only for
  the report's *current* pending assignment's `kurator_id`
  (server-resolved, same TOCTOU-aware pattern as `organization`/
  `campaign`'s decision endpoints). `decision = 'rejected'` requires
  `decision_note`.
- **Holds after operations**: verification decision.
- **Verification**: Confirmed — `403` for non-matching Kurator, `422`
  for missing note on reject, `409` if no pending assignment. Test:
  same shape as `organization`/`campaign`'s decision endpoints,
  including the reassignment-race test.

### INV-disbursement-13: `disbursement_request_logs` and `fund_usage_report_logs` are append-only

- **Statement**: No row in either table, once inserted, may ever be
  updated or deleted.
- **Holds after operations**: disbursement decision, disbursement
  execution, fund-usage report submission, `has_overdue_report`
  set/clear, verification assignment/decision.
- **Verification**: DB-level — `REVOKE UPDATE, DELETE` on both tables,
  same pattern as every other `_logs` table in this project.

### INV-disbursement-14: Fund-usage attachment access is private-bucket, owner-only upload

- **Statement**: `POST
  /fund-usage-reports/{reportId}/items/{itemId}/attachments` is
  owner-only (PDF/JPG/PNG, 5 MB max, multiple attachments per item
  allowed). Files stored in a private bucket
  (`fund_usage_report_item_attachments`, same pattern as
  `organization_attachments`). Attachment metadata + a 5-minute signed
  `download_url` are embedded directly in the report detail response
  (`GET /fund-usage-reports/{reportId}`) — there is no separate
  standalone download endpoint, unlike `organization`'s split
  list/download design.
- **Holds after operations**: attachment upload, report detail fetch.
- **Verification**: Confirmed — `422` on invalid file type/size,
  `403` on upload by non-owner. Test: `staff` upload attempt → `403`.
  Report detail's embedded `download_url` is valid for exactly 5
  minutes (same pattern as `organization`'s legal documents).

## State machines

### `disbursement_requests.status`

```
(none) -> pending -> approved -> disbursed
              |
              v
          rejected --(revision, new request)--> pending
```

`rejected` doesn't count as "active" for the one-active-per-campaign
constraint (INV-disbursement-03) — a fresh request after rejection is
a **new row**, old one kept as history.

### `fund_usage_reports.status`

```
(none) -> pending_verification -> verified / rejected
                                        |
                                        v
                              --(revision, resubmit)--> pending_verification
```

Same shape as `organization_curation_assignments`/
`campaign_curation_assignments` decision cycles — resubmission creates
a new `fund_usage_reports` row (not a reset of the old one), and a new
`fund_usage_report_verification_assignments` row.

### `organizations.has_overdue_report` (owned by `organization`, set/cleared here)

```
false -> true   (scheduler, INV-disbursement-07 — no report exists 30+ days post-disbursement)
true -> false   (first fund-usage report submission for that disbursement, INV-disbursement-08)
```

## References

- Related ERD: `docs/project/kencleng-erd.md` §5
  (`disbursement_requests`, `disbursement_request_logs`,
  `fund_usage_reports`, `fund_usage_report_items`,
  `fund_usage_report_item_attachments`,
  `fund_usage_report_verification_assignments`,
  `fund_usage_report_logs`)
- Related business process: `docs/project/kencleng-phase3-detail.md`
  Fitur 1–5
- Related invariants: `docs/spec/organization/invariants.md` —
  INV-organization-13 (fulfilled by INV-disbursement-07/08);
  `docs/spec/campaign/invariants.md` — pattern reference for curation
  assignment/decision invariants
- **Actual API (ground truth)**: `api/openapi/disbursement.yaml`
