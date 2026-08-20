# Domain Tasks — disbursement

> File: `docs/spec/disbursement/tasks.md`
> Status: draft — authored directly against `api/openapi/disbursement.yaml` 2026-08-20
> Last updated: 2026-08-20

| # | Task | Endpoint / surface | Depends on | Related invariants |
|---|---|---|---|---|
| 01 | Campaign report & narrative | `GET /campaigns/{id}/report`, `PATCH /campaigns/{id}/report-narrative` | campaign domain (Task 09) | INV-disbursement-01, 02 |
| 02 | Disbursement request | `POST/GET /campaigns/{id}/disbursement-requests` | 01 | INV-disbursement-03 |
| 03 | Disbursement decision & execution | `GET /disbursement-requests`, `GET/POST /disbursement-requests/{id}`, `.../decision` + internal execution | 02 | INV-disbursement-04, 05 |
| 04 | Fund-usage report submission & attachments | `POST/GET .../fund-usage-reports`, `GET /fund-usage-reports/{id}`, `POST .../attachments` | 03 | INV-disbursement-06, 14 |
| 05 | `has_overdue_report` scheduler | scheduler job (no endpoint) | 04 | INV-disbursement-07, 08, 09 |
| 06 | Fund-usage verification | `.../curation-queue`, `.../curation/assign`, `.../curation-assignments/mine`, `.../curation/decision` | 04 | INV-disbursement-10, 11, 12 |

## Task 01 — Campaign report & narrative

**What**: `GET /campaigns/{campaignId}/report` (public, `closed`
campaigns only, permanent), `PATCH
/campaigns/{campaignId}/report-narrative` (owner-only, ungated,
Markdown ≤ 5000 chars).

**KPI / metrics**:
- Report view for a non-`closed` campaign → `409`.
- Narrative edit by `staff` → `403`.
- Narrative edit before `closed` → `409`.
- Narrative over 5000 chars → `422`.
- Repeated edits succeed (not locked after first submission).
- **Frontend cross-check** (not this task's own test, but a
  coordination note): confirm the frontend renders `report_narrative`
  through a sanitizing Markdown pipeline, never raw HTML injection —
  flagged security-critical in `threat-model.md`.

## Task 02 — Disbursement request

**What**: `POST /campaigns/{campaignId}/disbursement-requests`
(owner-only, no body, `requested_amount = collected_amount`,
lump-sum), `GET` (list, includes rejected history).

**KPI / metrics**:
- Creation against a non-`closed` campaign → `409
  campaign-not-closed`.
- Creation while an active request exists → `409
  disbursement-already-active`.
- `staff` caller → `403`.
- Two concurrent creation attempts on the same campaign: DB unique
  index (`ux_disbursement_requests_one_active`) guarantees at most
  one succeeds, translated to a clean `409`.
- List includes past `rejected` requests, not just the active one.

## Task 03 — Disbursement decision & execution

**What**: `GET /disbursement-requests` (Admin queue, defaults to
`pending`), `GET /disbursement-requests/{id}` (detail), `POST
.../decision` (Admin-only, approve/reject), plus the internal
execution process (approved → disbursed, no HTTP endpoint).

**KPI / metrics**:
- Non-Admin decision attempt → `403`.
- Reject without `decision_note` → `422`.
- Decision on a non-`pending` request → `409`.
- Approval triggers `disbursed_at` set via the internal process,
  observable by polling the detail endpoint (not synchronously in the
  decision response — confirm this explicitly, don't assume
  synchronicity).
- Execution invoked twice for the same request → idempotent, no
  double-processing.
- **Critical**: verify no HTTP route exists for the execution
  transition (same audit as `donation` Task 01).

## Task 04 — Fund-usage report submission & attachments

**What**: `POST /disbursement-requests/{id}/fund-usage-reports`
(owner-only, strict reconciliation, clears `has_overdue_report` same
transaction), `GET` (report list + detail, with embedded signed-URL
attachments), `POST .../items/{itemId}/attachments` (owner-only,
private bucket).

**KPI / metrics**:
- Reconciliation mismatch (any amount, over or under) → `422`.
- Submission before `status = 'disbursed'` → `409`.
- `staff` caller (submission or attachment upload) → `403`.
- Submission on an org with `has_overdue_report = true` clears it in
  the same transaction — no window where both are true.
- Invalid attachment file type/size → `422`.
- Signed `download_url` valid for exactly 5 minutes.

## Task 05 — `has_overdue_report` scheduler

**What**: periodic job setting `organizations.has_overdue_report =
true` for disbursements past `disbursed_at + 30 days` with no
submitted report.

**KPI / metrics**:
- Disbursement past the 30-day mark with no report → flagged on the
  next run.
- Disbursement with a report already submitted (even just before the
  deadline) → never flagged.
- Job run twice near-simultaneously → idempotent, no double-log.
- **Open item, needs resolution before this task ships**: confirm the
  set/clear log entry lands in `organization_logs` (field owner), not
  `disbursement`'s own logs — see `threat-model.md`.

## Task 06 — Fund-usage verification

**What**: `GET /fund-usage-reports/curation-queue` (Admin), `POST
.../curation/assign` (Admin, conflict-of-interest + one-active-
assignment guards), `GET .../curation-assignments/mine` (Kurator),
`POST .../curation/decision` (assigned Kurator only, server-resolved
current assignment).

**KPI / metrics**: same shape as `organization`/`campaign`'s
equivalent curation tasks — 0 successful assignments with conflict of
interest, 0 successful second-pending-assignment creations, 0
successful decisions by a non-matching Kurator (including the
reassignment-race test), `422` on missing `decision_note` for reject.
Additionally: rejection here never sets `has_overdue_report`
(INV-disbursement-09) — explicit test, since this is the one place a
naive implementation might mistakenly wire rejection into the overdue
flag.

## References

- Related domain invariants: `docs/spec/disbursement/invariants.md`
- Related threat model: `docs/spec/disbursement/threat-model.md`
- **Actual API (ground truth)**: `api/openapi/disbursement.yaml`
- Feature specs (one per task, prefixed with the task `#` above):
  `docs/spec/disbursement/features/`
