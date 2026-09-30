# Feature Spec — 09: Closure (Auto + Force-Close)

> File: `docs/spec/4-campaign/features/09-closure.md`
> Domain: `campaign`
> Task: 09 (see `docs/spec/4-campaign/tasks.md`)
> Status: draft — authored against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-09-30

> **Slice 2 cross-reference (D1):** For Donation eligibility and
> settlement ordering, use `docs/spec/4-campaign/invariants.md#inv-campaign-13`
> together with `docs/spec/5-donation/invariants.md#inv-donation-02`
> and `#inv-donation-08`. An eligible accepted Donation remains
> settleable in full after close; later settlement cannot reopen the
> Campaign or change its winning close reason, and funding may exceed
> `max_amount`. This narrow reference does not reconcile the broader
> closure lifecycle in this historical feature spec. Its older
> max-trigger implementation notes are evidence only where they
> conflict with D1.

## Summary

A `published` Campaign may close through one of three triggers:
`max_amount` reached, `deadline` reached, or Admin force-close
(`POST /campaigns/{campaignId}/force-close`). For the Slice 2 D1
eligibility and settlement boundary, see
`docs/spec/4-campaign/invariants.md#inv-campaign-13` and the Donation
invariant cross-references above. The close-ordering mechanism remains
unselected in this narrow reference and belongs to the authorized
Campaign closure delivery work.

## Endpoint

`POST /campaigns/{campaignId}/force-close` (confirmed,
`api/openapi/campaign.yaml`) + a deadline scheduler job + the narrow
Donation threshold/eligibility boundary referenced above (no HTTP
endpoint for the threshold transition)

## Auth

Force-close: `bearerAuth` + `role = 'admin'` only.

## Request

`ForceCloseRequest`:

| Field | Type | Required | Notes |
|---|---|---|---|
| `decision_note` | string | Yes | Non-empty |

## Behavior

### Force-close
1. Reject (`403`) if caller isn't Admin.
2. Reject (`422`) if `decision_note` is empty.
3. `UPDATE campaigns SET status = 'closed', closed_at = now(),
   closed_reason = 'admin_force_closed', closed_by = current_user_id,
   decision_note = :note WHERE id = :id AND status = 'published'` —
   if 0 rows affected (already closed by another trigger), return
   `409`, not a silent success.
4. Log to `campaign_logs`.
5. Return `200` with the updated `Campaign`.

### Deadline scheduler (background, no endpoint)
1. Periodically: `UPDATE campaigns SET status = 'closed', closed_at =
   now(), closed_reason = 'deadline_reached' WHERE status =
   'published' AND deadline <= now()`.

### `max_amount` trigger (Campaign close boundary)
1. When successful funding reaches `max_amount`, the threshold may
   close a Campaign that is still eligible for fundraising, with
   `closed_reason = 'max_amount_reached'`, as part of the applicable
   successful-funding outcome. An already accepted Donation remains
   settleable in full after another close trigger wins; that later
   settlement updates funding atomically but does not reopen the
   Campaign or change its winning close reason. Funding may exceed
   `max_amount`. See D1 in INV-campaign-13 and the Donation invariant
   cross-reference above. No locking/isolation mechanism is selected.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Non-Admin force-close attempt | `403` |
| Empty `decision_note` | `422` |
| Campaign already closed (by any trigger) when force-close is attempted | `409` |

## Concurrency & correctness notes

- Campaign close must preserve the first winning close reason. The
  mechanism that orders submission, close, settlement, and funding is
  not selected by this feature reference; see D1 and defer mechanism
  design to the authorized delivery/architecture owner.
- **Explicit 3-way concurrency test required** (per `threat-model.md`):
  simulate all three triggers firing within the same short window on
  one campaign — assert exactly one `closed_reason` is recorded, the
  other two get clean `409`s (force-close) or silently no-op
  (scheduler/donation-trigger, which aren't user-facing requests with
  a response to check — verify via a post-condition assertion instead
  of an error response for those two).

## Test checklist

- [ ] Non-Admin force-close attempt → `403`.
- [ ] Force-close with empty `decision_note` → `422`.
- [ ] Force-close on an already-closed campaign → `409`.
- [ ] Deadline scheduler run twice near-simultaneously: exactly one
      closure, no error on the second run.
- [ ] Donation-focused D1 evidence: close-first rejects new submission;
      accepted-pending Donation remains settleable in full after close;
      later settlement does not change the winning close reason and
      funding may exceed threshold. Broader close-trigger race evidence
      remains owned by the Campaign closure delivery slice.
- [ ] `closed_by` populated only for `admin_force_closed`, `null` for
      the other two reasons.

## References

- `docs/spec/4-campaign/invariants.md` — INV-campaign-13
- `docs/spec/4-campaign/threat-model.md` — "Force-close" section
  (3-way race note)
- `docs/spec/4-campaign/tasks.md` — Task 09
- `api/openapi/campaign.yaml` — `POST /campaigns/{id}/force-close`
- `docs/project/kencleng-phase2-detail.md` — Fitur 3
