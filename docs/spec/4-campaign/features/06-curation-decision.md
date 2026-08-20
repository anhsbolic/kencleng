# Feature Spec — 06: Curation Decision

> File: `docs/spec/campaign/features/06-curation-decision.md`
> Domain: `campaign`
> Task: 06 (see `docs/spec/campaign/tasks.md`)
> Status: draft — authored against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`POST /campaigns/{campaignId}/curation/decision` — the assigned
Kurator approves or rejects. Approve → `status = 'approved'`; reject
(requires `decision_note`) → `status = 'rejected'`. A subsequent Owner
resubmit (via edit + Task 04's submit) sends `rejected → draft →
pending_curation` (new assignment cycle). Same server-resolved
current-assignment pattern as `organization`'s equivalent — including
the same TOCTOU consideration.

## Endpoints

`POST /campaigns/{campaignId}/curation/decision`, `GET
/campaigns/curation-assignments/mine` (confirmed,
`api/openapi/campaign.yaml`)

## Auth

- Decision: `bearerAuth` + `role = 'kurator'`, and must match the
  campaign's *current* pending assignment's `kurator_id`.
- Mine: `bearerAuth` + `role = 'kurator'`, scoped to caller.

## Request

### Decision — `CurationDecisionRequest` (shared schema, from
`organization.yaml`)
| Field | Type | Required | Notes |
|---|---|---|---|
| `decision` | enum (`approved`/`rejected`) | Yes | |
| `decision_note` | string | Conditionally | Required if `rejected` |

### Mine
| Param | Notes |
|---|---|
| `decision` | Optional filter |
| `cursor`, `limit` | Standard pagination |

## Behavior

### Decision
1. Reject (`403`) if caller doesn't hold `role = 'kurator'`.
2. Resolve the campaign's current `pending` assignment — reject
   (`409`) if none.
3. Reject (`403`) if that assignment's `kurator_id !=
   current_user_id` — fresh check against current state, same TOCTOU
   handling as `organization`'s decision endpoint.
4. Reject (`422`) if `decision = 'rejected'` and `decision_note` is
   empty.
5. Update the assignment row and `campaigns.status` in one
   transaction.
6. Best-effort notification to the campaign's organization owners.
7. Log to `campaign_logs`.
8. Return `200` with the updated `Campaign`.

### Mine
1. Reject (`403`) if caller doesn't hold `role = 'kurator'`.
2. Query scoped to `kurator_id = current_user_id`, optional `decision`
   filter.
3. Return `CampaignCurationAssignmentListResponse`, paginated.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Caller doesn't hold `role = 'kurator'` | `403` |
| No pending assignment for this campaign | `409` |
| Caller isn't the *current* assignment's Kurator | `403` |
| `decision = 'rejected'` without `decision_note` | `422` |

## Concurrency & correctness notes

- Same TOCTOU consideration as `organization`'s decision endpoint: a
  decision submitted after Admin reassigns the campaign to a different
  Kurator must be rejected against the *current* assignment, never
  silently applied to the stale one — needs the same explicit
  concurrency test.
- The assignment-update + `campaigns.status`-update must be one
  transaction.

## Test checklist

- [ ] Non-Kurator caller → `403`.
- [ ] Kurator who isn't the *current* assignment's `kurator_id` →
      `403`.
- [ ] Approve → `status = 'approved'`, same transaction as assignment
      update.
- [ ] Reject without `decision_note` → `422`.
- [ ] Reject with `decision_note` → `status = 'rejected'`.
- [ ] **Reassignment race**: Admin reassigns to Kurator B while
      Kurator A (originally assigned) submits concurrently → Kurator
      A's request rejected with `403` against the new assignment.
- [ ] `GET .../curation-assignments/mine` correctly scoped, filterable
      by `decision`.
- [ ] Resubmit after rejection: Owner edits (Task 01) →
      `rejected → draft`, then submits (Task 04) → `pending_curation`,
      a **new** assignment row created, old one kept as history.

## References

- `docs/spec/campaign/invariants.md` — INV-campaign-07
- `docs/spec/campaign/threat-model.md` — references
  `organization/threat-model.md`'s equivalent section
- `docs/spec/campaign/tasks.md` — Task 06
- `docs/spec/organization/features/10-curation-decision.md` —
  structural precedent (TOCTOU handling)
- `api/openapi/campaign.yaml` — `POST .../curation/decision`, `GET
  /campaigns/curation-assignments/mine`
