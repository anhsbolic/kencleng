# Feature Spec — 06: Fund-Usage Verification

> File: `docs/spec/disbursement/features/06-fund-usage-verification.md`
> Domain: `disbursement`
> Task: 06 (see `docs/spec/disbursement/tasks.md`)
> Status: draft — authored against `api/openapi/disbursement.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

Kurator verification of a submitted fund-usage report — same pattern
as `organization`/`campaign`'s curation flows: Admin assigns
(conflict-of-interest checked), Kurator decides (approve/reject,
server-resolved current assignment). Rejection carries **no**
`has_overdue_report` consequence (INV-disbursement-09) — the
timeliness flag is purely about the *first submission's* timing, not
verification outcome.

## Endpoints

`GET /fund-usage-reports/curation-queue`, `POST
/fund-usage-reports/{reportId}/curation/assign`, `GET
/fund-usage-reports/curation-assignments/mine`, `POST
/fund-usage-reports/{reportId}/curation/decision` (confirmed,
`api/openapi/disbursement.yaml`)

## Auth

- Queue, assign: `bearerAuth` + `role = 'admin'`.
- Mine, decision: `bearerAuth` + `role = 'kurator'` (decision
  additionally requires matching the report's current assignment).

## Request

### Assign — `AssignCuratorRequest` (shared schema, `organization.yaml`)
| Field | Type | Required |
|---|---|---|
| `kurator_id` | UUID | Yes |

### Decision — `CurationDecisionRequest` (shared schema,
`organization.yaml`)
| Field | Type | Required | Notes |
|---|---|---|---|
| `decision` | enum | Yes | `approved`/`rejected` |
| `decision_note` | string | Conditionally | Required if `rejected` |

## Behavior

### Queue
1. Reject (`403`) if caller isn't Admin.
2. Query `fund_usage_reports WHERE status = 'pending_verification' AND
   NOT EXISTS (verification assignment WHERE decision = 'pending')`.
3. Return `FundUsageReportListResponse`, paginated.

### Assign
1. Reject (`403`) if caller isn't Admin.
2. Reject (`409`) if `kurator_id` is a representative (any `level`) of
   the report's campaign's organization.
3. Reject (`409`) if a pending assignment already exists for this
   report.
4. Insert `fund_usage_report_verification_assignments`.
5. Best-effort notification to the assigned Kurator (`type =
   kurator_assigned`).
6. Log to `fund_usage_report_logs`.
7. Return `201` with the `FundUsageReportVerificationAssignment`.

### Mine
1. Reject (`403`) if caller isn't Kurator.
2. Query scoped to `kurator_id = current_user_id`, optional `decision`
   filter.
3. Return `FundUsageReportVerificationAssignmentListResponse`,
   paginated.

### Decision
1. Reject (`403`) if caller isn't Kurator.
2. Resolve the report's current pending assignment — reject (`409`)
   if none.
3. Reject (`403`) if that assignment's `kurator_id != current_user_id`
   (fresh check against current state, same TOCTOU handling as
   `organization`/`campaign`).
4. Reject (`422`) if `decision = 'rejected'` and `decision_note` is
   empty.
5. Update the assignment row and `fund_usage_reports.status` in one
   transaction. **Do not** touch `has_overdue_report` in either
   branch (approve or reject) — that flag is set/cleared exclusively
   by INV-disbursement-07/08's mechanisms, never by this decision.
6. Best-effort notification to the Owner (`type =
   fund_usage_report_verified`/`fund_usage_report_rejected`).
7. Log to `fund_usage_report_logs`.
8. Return `200` with the updated `FundUsageReport`.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Non-Admin on queue/assign | `403` |
| `kurator_id` is a representative of the report's organization | `409` |
| Report already has a pending assignment | `409` |
| Non-Kurator on mine/decision | `403` |
| No pending assignment for this report | `409` |
| Kurator isn't the *current* assignment's `kurator_id` | `403` |
| `decision = 'rejected'` without `decision_note` | `422` |

## Concurrency & correctness notes

- Same TOCTOU handling as `organization`/`campaign`'s decision
  endpoints — a decision submitted after reassignment must be rejected
  against the *current* assignment.
- DB unique partial index is the correctness guarantee for
  one-active-assignment-per-report.

## Test checklist

- [ ] Non-Admin queue/assign attempt → `403`.
- [ ] `kurator_id` who is a representative of the report's
      organization → `409`.
- [ ] Two concurrent assignment attempts for the same report: at most
      one succeeds, clean `409` for the other.
- [ ] Non-Kurator mine/decision attempt → `403`.
- [ ] Kurator not matching the *current* assignment → `403`.
- [ ] Reject without `decision_note` → `422`.
- [ ] **Reassignment race**: Admin reassigns while the original
      Kurator submits concurrently → original Kurator's request
      rejected against the new assignment.
- [ ] **Explicit no-side-effect test**: rejecting a report never sets
      `has_overdue_report`, regardless of how late the original
      submission was relative to the 30-day deadline.

## References

- `docs/spec/disbursement/invariants.md` — INV-disbursement-09, 10,
  11, 12
- `docs/spec/disbursement/threat-model.md` — references
  `organization/threat-model.md`'s equivalent sections
- `docs/spec/disbursement/tasks.md` — Task 06
- `docs/spec/organization/features/09-curation-assignment.md`,
  `10-curation-decision.md` — structural precedent
- `api/openapi/disbursement.yaml` — verification endpoints
