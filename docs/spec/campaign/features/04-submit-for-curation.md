# Feature Spec — 04: Submit for Curation

> File: `docs/spec/campaign/features/04-submit-for-curation.md`
> Domain: `campaign`
> Task: 04 (see `docs/spec/campaign/tasks.md`)
> Status: draft — authored against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`POST /campaigns/{campaignId}/submit` — owner-only, `draft →
pending_curation`. Locks the campaign from further direct edits until
a curation decision is made.

## Endpoint

`POST /campaigns/{campaignId}/submit` (confirmed,
`api/openapi/campaign.yaml`)

## Auth

`bearerAuth` required, owner-level representative of the owning
organization only.

## Request

No body.

## Behavior

1. Reject (`403`) if caller isn't an owner-level representative.
2. Reject (`409`) if `status != 'draft'`.
3. Update `status = 'pending_curation'` — atomic status guard (`WHERE
   status = 'draft'`), so a double-submit race is naturally a no-op
   on the second attempt rather than a double transition.
4. Best-effort notification to Admin (per
   `kencleng-phase1-detail.md` Fitur 4: "a campaign enters
   `pending_curation` → the system notifies the Admin").
5. Return `200` with the updated `Campaign`.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Caller is `staff` or not a representative | `403` |
| Campaign not `draft` | `409` |

## Concurrency & correctness notes

- The `draft → pending_curation` transition must be atomic at the
  query level (`WHERE status = 'draft'`) — two near-simultaneous
  submit requests result in exactly one transition, the other gets a
  clean `409`, not a race condition or double-notify.

## Test checklist

- [ ] `staff` caller → `403`.
- [ ] Submit outside `draft` → `409`.
- [ ] Two concurrent submit requests on the same draft: exactly one
      succeeds, the other gets `409`.
- [ ] Successful submit triggers a best-effort Admin notification
      (failure doesn't fail this endpoint, per INV-notification-06).

## References

- `docs/spec/campaign/invariants.md` — INV-campaign-03, 04
- `docs/spec/campaign/threat-model.md` — "Submit for curation" section
- `docs/spec/campaign/tasks.md` — Task 04
- `api/openapi/campaign.yaml` — `POST /campaigns/{id}/submit`
