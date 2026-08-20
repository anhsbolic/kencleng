# Feature Spec — 06: Guest-Email Reveal

> File: `docs/spec/donation/features/06-guest-email-reveal.md`
> Domain: `donation`
> Task: 06 (see `docs/spec/donation/tasks.md`)
> Status: draft — new endpoint, decided 2026-08-20
> Last updated: 2026-08-20

## Summary

`GET /donations/{donationId}/guest-email` — Admin, or a Kurator
assigned (current or historical) to curate the donation's campaign,
may reveal the decrypted `guest_email`. Every successful call logs to
`donation_logs` in the same request. **Not yet present** in
`api/openapi/donation.yaml` — needs to be added at implementation
time; this spec fixes its behavior ahead of that.

## Endpoint

`GET /donations/{donationId}/guest-email` *(new — to be added)*

## Auth

`bearerAuth` required, plus one of:
- `role = 'admin'`, or
- `role = 'kurator'` **and** is or was assigned (via
  `campaign_curation_assignments`) to curate the campaign this
  donation belongs to — not any Kurator.

## Request

Path: `donationId`. No body.

## Behavior

1. Load the donation — `404` if it doesn't exist.
2. Resolve caller's authorization: Admin, or a Kurator with a
   qualifying (current/historical) assignment to the donation's
   `campaign_id`. Reject (`403`) otherwise.
3. If the donation has no `guest_email` at all (registered donor, or
   guest who didn't provide one): return `404` — there's nothing to
   reveal, distinct from an authorization failure.
4. Within a single transaction: decrypt `guest_email`, insert
   `donation_logs` (`action_type = 'guest_email_revealed'`,
   `actor_user_id = current_user_id`, `donation_id`). Both succeed or
   both fail — no partial state where a reveal happened without a log
   entry.
5. Return `200` with `{ "guest_email": "..." }`.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Caller isn't Admin or a qualifying Kurator | `403` |
| Donation doesn't exist | `404` |
| Donation exists but has no `guest_email` | `404` |

## Concurrency & correctness notes

- The reveal (decrypt + return) and the `donation_logs` insert must be
  one transaction — never a state where the caller received the email
  but no log entry exists (or vice versa, though the vice versa case
  is less concerning security-wise).
- No special concurrency handling needed beyond the standard
  transaction boundary — this is a low-frequency, read-mostly
  operation.

## Test checklist

- [ ] Non-Admin, non-Kurator caller → `403`.
- [ ] Kurator not assigned (current or historical) to the donation's
      campaign → `403`.
- [ ] Kurator with a current assignment → succeeds.
- [ ] Kurator with only a *historical* (already-decided) assignment →
      succeeds (mirrors `organization`'s attachment-access pattern —
      confirm this is the intended scope, not current-only).
- [ ] Admin → succeeds for any donation.
- [ ] Donation with no `guest_email` → `404`, distinct from the
      authorization-failure `403`.
- [ ] Every successful call produces exactly one `donation_logs` row.
- [ ] Fault-inject a failure in the log insert (simulate a DB error
      after decryption but before the log write) — assert the whole
      transaction rolls back, caller does **not** receive the email if
      the log couldn't be written.

## References

- `docs/spec/donation/invariants.md` — INV-donation-14, 15
- `docs/spec/donation/threat-model.md` — "Guest-email reveal" section
- `docs/spec/donation/tasks.md` — Task 06
- `docs/spec/organization/features/05-legal-document-attachments.md`
  — structural precedent (assignment-scoped Kurator access)
