# Threat Model — donation

> File: `docs/spec/donation/threat-model.md`
> Status: draft — authored directly against `api/openapi/donation.yaml` 2026-08-20
> Last updated: 2026-08-20

## Actors & trust boundaries

| Actor | Authenticated? | Trust boundary crossed |
|---|---|---|
| Public / anonymous donor | No | `POST /campaigns/{id}/donations` (guest path), `GET /campaigns/{id}/donations` (public list), `GET /donations/{id}/status` (token-guarded, not identity-guarded) |
| Registered donor | Yes | Same submit endpoint (authenticated path), `GET /account/donations`, `GET /account/donations/claimable`, `POST /account/donations/{id}/claim` |
| Admin | Yes | New: `GET /donations/{id}/guest-email` (INV-donation-14) |
| Kurator (assigned/historical to the donation's campaign) | Yes | Same new endpoint, scoped |
| Internal settlement process | N/A — in-process, no HTTP endpoint | The simulated payment "callback" — explicitly never exposed externally (`kencleng-phase2-detail.md` Fitur 1 Security notes) |
| Anyone holding a `status_token` string | No (token-based, not session-based) | `GET /donations/{id}/status` — the token itself is the credential, for the token's entire (non-expiring) lifetime |

## STRIDE per operation

### Submit donation — `POST /campaigns/{campaignId}/donations`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | An authenticated request with a forged/stale token attempts to submit under someone else's identity | Standard `bearerAuth` validation, same as every other authenticated endpoint | None |
| Tampering | Authenticated caller supplies `guest_name`/`guest_email` in the body, hoping to submit "as guest" under a different identity while still being authenticated | Body guest fields are ignored when authenticated (INV-donation-07) — session identity always wins | None |
| Tampering | Double-submit (accidental double-click, or a malicious retry) creates two donation records | `Idempotency-Key` header required (INV-donation-11) | None, assuming the client correctly reuses the key on retry — a client bug that generates a new key each time would still create duplicates; this is a client-correctness concern, not a server-side gap |
| Tampering | Submission raced against the campaign closing (deadline/max_amount/force-close) between page load and submit | `WHERE status = 'published'` guard at submit time (INV-donation-02) | None |
| Repudiation | N/A — the donation record itself is the record; no separate "who submitted" ambiguity for registered donors. For guest donations, there's inherently no stronger identity than what was voluntarily provided. | — | — |
| Information disclosure | `status_token` returned in the `201` response and (if `guest_email` provided) emailed — if intercepted, grants read access to that donation's status | Token is random and long enough to resist brute-force (INV-donation-05); transport is HTTPS (project-wide assumption); the token only grants **read** access to non-sensitive-beyond-amount/method data, not a destructive action | Accepted — same risk class as any bookmarkable, non-expiring access token; consistent with the project's own explicit trade-off reasoning in `kencleng-phase2-detail.md` |
| Denial of service | Automated mass-submission of tiny (Rp 5.000) donations to spam a campaign's donor list / inflate `donor_count` | Minimum amount (INV-donation-01) raises the cost per spam unit somewhat, but doesn't prevent it; no CAPTCHA/rate-limiting documented for this public endpoint | **Accepted for a sandbox project** — a real deployment would want bot-mitigation on a public, unauthenticated financial-adjacent endpoint; flagged as an intentional gap given this project's stated non-production scope |
| Elevation of privilege | N/A | — | — |

### Public donor list — `GET /campaigns/{campaignId}/donations`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A — public, no auth | — | — |
| Tampering | N/A — read-only | — | — |
| Repudiation | N/A | — | — |
| Information disclosure | `guest_email` or any PII leaking into the public shape | `DonationListItem` schema explicitly excludes it — confirmed at the schema level, not just a runtime check to remember (INV-donation-06) | None, assuming implementation matches the schema exactly (i.e. no accidental over-fetching in the query that then gets serialized) |
| Information disclosure | `is_anonymous` donors still identifiable via amount/timing correlation (e.g. "an anonymous Rp 50.000.000 donation right after a known donor viewed the page") | Not mitigated — inherent to any public amount+timestamp list | Low, accepted — this is a general public-transparency-vs-privacy trade-off inherent to the feature's design (public donor lists are the explicit product requirement), not something this domain's spec can fix without changing the feature itself |
| Denial of service | N/A beyond standard pagination limits | `LimitParam` cap | None |
| Elevation of privilege | N/A | — | — |

### Token-based status check — `GET /donations/{donationId}/status`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A — token-based by design, not identity-based | — | — |
| Tampering | N/A — read-only | — | — |
| Repudiation | N/A | — | — |
| Information disclosure | Token brute-forcing (guessing `token` query params against a known/guessed `donationId`) | Token is long/random (INV-donation-05); `401` on mismatch gives no signal distinguishing "wrong token" from "donation doesn't exist" — check this is actually implemented as a uniform response, not two distinguishable error paths | Low, worth an explicit test that the `401` response is identical regardless of *why* it failed (donation not found vs. token mismatch) |
| Denial of service | Brute-force token-guessing at scale | No rate-limiting documented on this specific endpoint | Low — token space is large enough that this is impractical even without rate-limiting, but a real deployment would still want one; accepted for sandbox scope |
| Elevation of privilege | N/A | — | — |

### Account donation history & claimable list — `GET /account/donations`, `GET /account/donations/claimable`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A | `bearerAuth` | None |
| Tampering | N/A — read-only | — | — |
| Repudiation | N/A | — | — |
| Information disclosure | Claimable-list matching leaks whether *any* guest donation exists for an email the caller doesn't actually own | Matching is against the caller's **own verified** `primary_email` only (`guest_email_hash` comparison) — no path to query by an arbitrary email | None |
| Denial of service | N/A | — | — |
| Elevation of privilege | Unverified-email user attempts to view claimable donations | `403`, confirmed explicit | None |

### Claim — `POST /account/donations/{donationId}/claim`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A | `bearerAuth` | None |
| Tampering | Claim a donation whose `guest_email` doesn't actually match the caller | `403`, confirmed explicit ("this donation's `guest_email` doesn't match") | None |
| Tampering | Double-claim race (two users sharing an email, or a retried request) | `WHERE donor_user_id IS NULL` guard, confirmed (INV-donation-12) | None |
| Repudiation | Claim itself isn't separately logged to `donation_logs` per the ERD's comment (only guest-email reveal is anticipated there) — worth confirming this is intentional, since claiming does change ownership of a financial record | Not currently in scope per the ERD | **Open, low priority** — flag for a decision on whether `claimed_at`/`donor_user_id` changes should also produce a `donation_logs` entry; not blocking, since `claimed_at` itself on the row is already a durable record of *when*, just not a separate immutable log entry |
| Information disclosure | N/A | — | — |
| Denial of service | N/A | — | — |
| Elevation of privilege | Unverified-email user attempts claim | `403` | None |

### Guest-email reveal — `GET /donations/{donationId}/guest-email` (new, INV-donation-14)

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A | `bearerAuth` + Admin/scoped-Kurator check | None |
| Tampering | N/A — read-only | — | — |
| Repudiation | Every reveal must log — this is the entire point of the endpoint | `donation_logs` entry per call (INV-donation-14) | None, assuming implementation doesn't allow a reveal to succeed without the log write also succeeding — recommend the same transaction (reveal query + log insert), not two independent operations where one could fail silently |
| Information disclosure | An unrelated Kurator (not assigned to this donation's campaign) attempts reveal | Assignment-scoped check (mirrors `organization`'s legal-document pattern) | None if enforced correctly — verify against the *current or historical* assignment, not just "any Kurator," same care as the curation-decision endpoints elsewhere in this project |
| Denial of service | Admin/Kurator mass-revealing guest emails across many donations (internal misuse, not an external attacker) | Every reveal is logged — provides an audit trail for after-the-fact review, doesn't prevent the action itself | Accepted — same posture as `organization`'s NPWP reveal; logging deters/detects misuse rather than preventing it outright, consistent with this project's audit-log philosophy |
| Elevation of privilege | Non-Admin, non-Kurator caller | `403` | None |

### Internal settlement process (no HTTP endpoint)

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | If this were ever accidentally exposed as a real endpoint, anyone could forge a "payment succeeded" callback | Explicitly documented as internal-only, never exposed externally (`kencleng-phase2-detail.md`) — this is an architectural requirement, not just a convention; implementation must ensure no route registers this | **Critical to get right at implementation time** — flag prominently in the feature spec, since this is the one place in the whole domain where a routing mistake would be a severe vulnerability (fake payment confirmations) |
| Tampering | N/A beyond the above | — | — |
| Repudiation | N/A — internal process | — | — |
| Information disclosure | N/A | — | — |
| Denial of service | N/A | — | — |
| Elevation of privilege | The spoofing case above, reframed | — | — |

## Knowingly accepted residual risk

- **No bot-mitigation on public donation submission** — accepted for
  a sandbox project; a real deployment would want CAPTCHA/rate-limiting
  on this public, unauthenticated endpoint.
- **`status_token` as a bookmarkable, non-expiring credential** —
  accepted, matches the project's own explicit reasoning (read-only,
  non-destructive lookup).
- **Public donor-list correlation risk for anonymous donors** —
  inherent to the feature's public-transparency design, not fixable
  at this domain's level without changing product requirements.
- **No rate-limiting on token-guessing at the status-check endpoint**
  — accepted, token space is large enough to make this impractical.

## Open items to resolve

1. Should claiming a donation (`donor_user_id`/`claimed_at` change)
   also produce a `donation_logs` entry, beyond the row's own
   `claimed_at` timestamp? Low priority, not blocking.
2. **Critical implementation-time requirement**: the internal
   settlement process must never be reachable via any registered HTTP
   route — this needs explicit verification during code review, not
   just documentation.

## References

- Related domain invariants: `docs/spec/donation/invariants.md`
- Related ERD: `docs/project/kencleng-erd.md` §4
- Related business process: `docs/project/kencleng-phase2-detail.md`
  Fitur 1, 4
- Related threat model precedent: `docs/spec/organization/threat-model.md`
  (reveal-endpoint pattern, referenced for INV-donation-14)
- **Actual API (ground truth)**: `api/openapi/donation.yaml`
