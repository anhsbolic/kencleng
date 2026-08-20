# Feature Spec — 02: Disbursement Request

> File: `docs/spec/disbursement/features/02-disbursement-request.md`
> Domain: `disbursement`
> Task: 02 (see `docs/spec/disbursement/tasks.md`)
> Status: draft — authored against `api/openapi/disbursement.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`POST /campaigns/{campaignId}/disbursement-requests` — Owner-only, no
request body, `requested_amount` always equals `Campaign.
collected_amount` at request time (lump-sum, single-shot). `GET` —
history list, including past `rejected` requests.

## Endpoint

`POST`/`GET /campaigns/{campaignId}/disbursement-requests` (confirmed,
`api/openapi/disbursement.yaml`)

## Auth

`bearerAuth` required. Create: owner-level representative only. List:
representative (any level) — **assumed default**, see
`invariants.md`'s note; confirm at implementation time.

## Request

Create: no body. List: `CursorParam`, `LimitParam`.

## Behavior

### Create
1. Reject (`403`) if caller isn't an owner-level representative of
   this campaign's organization.
2. Reject (`409 campaign-not-closed`) if `Campaign.status !=
   'closed'`.
3. Attempt insert: `requested_amount = Campaign.collected_amount`,
   `status = 'pending'`, `requested_by = current_user_id`. The DB
   unique partial index (`ux_disbursement_requests_one_active`)
   rejects this if an active request (`pending`/`approved`/
   `disbursed`) already exists for this campaign — catch and
   translate to `409 disbursement-already-active`.
4. Best-effort notification to Admin (`type =
   admin_new_curation_item`, dual channel, per
   `kencleng-phase3-detail.md` Fitur 2).
5. Log to `disbursement_request_logs`.
6. Return `201` with the `DisbursementRequest`.

### List
1. Reject (`403`) if caller isn't a representative.
2. Query `disbursement_requests WHERE campaign_id = :id`, **all**
   statuses (includes `rejected` history, confirmed explicit).
3. Return `DisbursementRequestListResponse`, paginated.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Create, caller is `staff` or not a representative | `403` |
| Create, campaign not `closed` | `409 campaign-not-closed` |
| Create, active request already exists | `409 disbursement-already-active` |

## Concurrency & correctness notes

- The DB unique partial index is the entire correctness mechanism for
  "one active request per campaign" — two near-simultaneous creation
  attempts: at most one succeeds, the app must catch the constraint
  violation and return the clean `409`, not a raw `500`.

## Test checklist

- [ ] `staff` create attempt → `403`.
- [ ] Create against a non-`closed` campaign → `409
      campaign-not-closed`.
- [ ] Create while an active request exists → `409
      disbursement-already-active`.
- [ ] Two concurrent create attempts on the same campaign: exactly one
      succeeds, clean `409` for the other.
- [ ] `requested_amount` always exactly matches `collected_amount` at
      request time (no client-supplied override possible, since
      there's no request body).
- [ ] List includes past `rejected` requests, not just the currently
      active one.
- [ ] Successful creation triggers a best-effort Admin notification.

## References

- `docs/spec/disbursement/invariants.md` — INV-disbursement-03
- `docs/spec/disbursement/threat-model.md` — "Disbursement request"
  section
- `docs/spec/disbursement/tasks.md` — Task 02
- `api/openapi/disbursement.yaml` — `DisbursementRequest`,
  `DisbursementRequestListResponse` schemas
