# Feature Spec — 08: Unpublish (Manual)

> File: `docs/spec/campaign/features/08-unpublish-manual.md`
> Domain: `campaign`
> Task: 08 (see `docs/spec/campaign/tasks.md`)
> Status: draft — authored against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`POST /campaigns/{campaignId}/unpublish` — Owner-only, requires
`decision_note`, sets `unpublish_reason = 'owner_manual'`, logged.
Only valid from `status = 'published'`. This task also covers the
**auto-unpublish** side effect received from `organization` domain
(INV-campaign-11) — the integration test for that cross-domain flow
lives here.

## Endpoint

`POST /campaigns/{campaignId}/unpublish` (confirmed,
`api/openapi/campaign.yaml`)

## Auth

`bearerAuth` required, owner-level representative only.

## Request

`UnpublishRequest`:

| Field | Type | Required | Notes |
|---|---|---|---|
| `decision_note` | string | Yes | Non-empty |

## Behavior

### Manual unpublish
1. Reject (`403`) if caller isn't owner-level.
2. Reject (`409`) if `status != 'published'`.
3. Reject (`422`) if `decision_note` is empty.
4. Update `status = 'unpublished'`, `unpublish_reason =
   'owner_manual'`.
5. Log to `campaign_logs` (actor, timestamp, `decision_note`).
6. Return `200` with the updated `Campaign`.

### Auto-unpublish (received from `organization` domain — no endpoint here)
Executes as part of `organization`'s edit transaction (see
`docs/spec/organization/features/04-organization-edit.md` step 7) —
`UPDATE campaigns SET status = 'unpublished', unpublish_reason =
'organization_re_verification' WHERE organization_id = :id AND status
= 'published'`. No `decision_note` (system-triggered, per
`kencleng-phase1-detail.md` Fitur 5 — logged automatically without an
Owner-typed reason).

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Caller is `staff` or not a representative | `403` |
| `status != 'published'` | `409` |
| Empty/missing `decision_note` | `422` |

## Concurrency & correctness notes

- Manual unpublish and auto-unpublish both key off `WHERE status =
  'published'` — if both somehow raced (extremely unlikely given
  auto-unpublish is org-wide and manual is per-campaign, but worth
  noting), whichever commits first wins, the other becomes a no-op
  via the status guard.
- **This is where INV-campaign-11's full integration test lives**:
  publish 2 campaigns under a `verified` organization, edit a
  legal/identity field on that organization (`confirm=true`), assert
  both campaigns flip to `unpublished` with `unpublish_reason =
  'organization_re_verification'`, in the same transaction as the
  organization's status change.

## Test checklist

- [ ] `staff` caller → `403`.
- [ ] Unpublish outside `published` → `409`.
- [ ] Empty `decision_note` → `422`.
- [ ] Successful manual unpublish logs actor/timestamp/reason.
- [ ] **Cross-domain**: organization re-verification auto-unpublishes
      all `published` campaigns under it, same transaction,
      `unpublish_reason = 'organization_re_verification'`, no
      `decision_note` required or present.
- [ ] Auto-unpublished campaigns are not auto-republished — remain
      `unpublished` until the Owner manually republishes (Task 07).

## References

- `docs/spec/campaign/invariants.md` — INV-campaign-10, 11, 12
- `docs/spec/campaign/threat-model.md` — "Publish / unpublish /
  republish" section
- `docs/spec/campaign/tasks.md` — Task 08
- `docs/spec/organization/invariants.md` — INV-organization-09
  (referenced, not redefined)
- `api/openapi/campaign.yaml` — `POST /campaigns/{id}/unpublish`
