# Feature Spec — 03: Disbursement Decision & Execution

> File: `docs/spec/disbursement/features/03-disbursement-decision-execution.md`
> Domain: `disbursement`
> Task: 03 (see `docs/spec/disbursement/tasks.md`)
> Status: draft — authored against `api/openapi/disbursement.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`GET /disbursement-requests` (Admin queue), `GET
/disbursement-requests/{id}` (detail), `POST .../decision`
(Admin-only, approve/reject), plus an **internal, non-HTTP**
disbursement-execution process that transitions `approved → disbursed`
after a simulated short delay.

## `[CRITICAL — implementation requirement]`

Same class of risk as `donation` domain's payment settlement: the
disbursement-execution process **must never** be reachable via any
registered HTTP route. Verify explicitly during code review.

## Endpoints

`GET /disbursement-requests`, `GET
/disbursement-requests/{disbursementRequestId}`, `POST
.../decision` (confirmed, `api/openapi/disbursement.yaml`) + internal
execution process (no endpoint)

## Auth

- Admin queue list: `bearerAuth` + `role = 'admin'`.
- Detail: `bearerAuth` + representative/Kurator/Admin — **assumed
  default**, confirm at implementation time (see
  `invariants.md`'s note).
- Decision: `bearerAuth` + `role = 'admin'`.

## Request

### Admin queue
| Param | Notes |
|---|---|
| `status` | Optional filter, defaults to `pending` |
| `cursor`, `limit` | Standard pagination |

### Decision — `DisbursementDecisionRequest`
| Field | Type | Required | Notes |
|---|---|---|---|
| `decision` | enum (`approved`/`rejected`) | Yes | |
| `decision_note` | string | Conditionally | Required if `rejected` |

## Behavior

### Admin queue
1. Reject (`403`) if caller isn't Admin.
2. Query `disbursement_requests`, filtered by `status` (default
   `pending`).
3. Return `DisbursementRequestListResponse`, paginated.

### Detail
1. Load the request — `404` if it doesn't exist.
2. Resolve caller's relationship (representative of the owning
   campaign's organization, assigned/historical Kurator on this
   campaign, or Admin) — `403` if none.
3. Return `DisbursementRequest`.

### Decision
1. Reject (`403`) if caller isn't Admin.
2. Reject (`409`) if `status != 'pending'`.
3. Reject (`422`) if `decision = 'rejected'` and `decision_note` is
   empty.
4. Update `status`, `reviewed_by = current_user_id`, `decision_note`,
   `decided_at = now()`.
5. Best-effort notification to the Owner (`type =
   disbursement_approved`/`disbursement_rejected`, dual channel).
6. Log to `disbursement_request_logs`.
7. If `decision = 'approved'`: enqueue the internal execution process
   (does not block this response).
8. Return `200` with the updated `DisbursementRequest` — `status =
   'approved'` immediately, **not** `'disbursed'` yet.

### Internal execution process (no HTTP endpoint)
1. After a short simulated delay: `UPDATE disbursement_requests SET
   status = 'disbursed', disbursed_at = now() WHERE id = :id AND
   status = 'approved'` — conditional guard makes this idempotent.
2. Best-effort notification to the Owner (`type = ...disbursed...`,
   dual channel).
3. Log to `disbursement_request_logs`.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Admin queue, non-Admin caller | `403` |
| Detail, no relationship to the request | `403` |
| Detail, request doesn't exist | `404` |
| Decision, non-Admin caller | `403` |
| Decision, request not `pending` | `409` |
| Decision, `rejected` without `decision_note` | `422` |

## Concurrency & correctness notes

- Step 1 of the internal execution process is the entire correctness
  mechanism for idempotency — running it twice for the same request
  is a safe no-op the second time.
- The decision endpoint's response reflects `status = 'approved'`
  immediately; the client must poll `GET /disbursement-requests/{id}`
  to observe the subsequent `disbursed` transition — this is by
  design, not a bug to "fix" into a synchronous response.

## Test checklist

- [ ] Non-Admin caller on queue/decision → `403`.
- [ ] Decision on a non-`pending` request → `409`.
- [ ] Reject without `decision_note` → `422`.
- [ ] Approval response shows `status = 'approved'`, not yet
      `'disbursed'`.
- [ ] Polling detail after approval eventually shows `status =
      'disbursed'`, `disbursed_at` populated.
- [ ] Internal execution invoked twice for the same request →
      idempotent, `disbursed_at` set exactly once.
- [ ] **Route audit**: confirm no HTTP route exists for the execution
      transition.
- [ ] Detail endpoint: representative, Kurator, Admin all succeed;
      unrelated caller → `403` (confirm the assumed access rule).

## References

- `docs/spec/disbursement/invariants.md` — INV-disbursement-04, 05
- `docs/spec/disbursement/threat-model.md` — "Disbursement decision &
  execution" section
- `docs/spec/disbursement/tasks.md` — Task 03
- `docs/spec/donation/features/01-submit-donation-settlement.md` —
  structural precedent (internal process must never be HTTP-reachable)
- `api/openapi/disbursement.yaml` — decision/detail/queue endpoints
