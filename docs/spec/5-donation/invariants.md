# Domain Invariant — Donation

> File: `docs/spec/5-donation/invariants.md`
> Status: draft — Slice 2 reconciliation; requires Donation, Security/PII, and applicable Human review before `agreed`
> Last updated: 2026-09-30
> Product/MVP basis: `docs/product/mvp-scope.md` §§4–7; `docs/product/mvp-delivery-slices.md` §§5–6

## Domain summary

For Slice 2, Donation owns guest submission, persisted sandbox status, safe status revisit, and the donation-side money outcome. Account history, public donor listing, claims, guest-email reveal, and real payment rails are outside this slice. Campaign owns Campaign lifecycle state; Donation references its eligibility/threshold boundary below.

## Reconciliation classifications

| Historical material | Classification | Current treatment and authority |
|---|---|---|
| Minimum Rp5.000, whole IDR input, Rp1 increments, exact-decimal money | KEEP / ADAPT | Keep amount rules from Product/MVP. Require exact decimal and no `float64`; wire/storage encoding and derived precision/rounding remain behind O1 `AUTHORITY_SYNC` and owner decisions. |
| Submit only while Campaign eligible; full amount; threshold crossing closes Campaign | ADAPT | Apply O6 and D1 below: atomically order new submission eligibility against close; already accepted submissions remain settleable in full after close; preserve winning close reason; total may exceed threshold. No mechanism selected. |
| Guest name/email fields | ADAPT | Name optional and not public by default. Email optional and opt-in for status notices only; ownership verification and Security/PII retention/control remain open. O11 conflict stays a `HUMAN_DECISION`. |
| Non-expiring status token | REPLACE | Product requires a difficult-to-guess guest URL valid for 24 hours, status-only. Carrier, storage, exposure mitigations, lifecycle, and residual-risk acceptance remain O4. |
| Async 2–5 second / 5% settlement | REPLACE | Backend simulator owns persisted `pending` → `success`/`failed`. Failure is a clearly labeled demo scenario controlled by backend configuration/fixture. Timing and exact scenario mechanics remain O2. |
| Same-key retry without payload comparison; settlement replay | ADAPT | Same key/same payload returns the original Donation; same key/different payload rejects. Deliberate new Donation requires new key. This request policy is distinct from exact-once settlement/funding. |
| Guest status failure `401` / proposed `404` | ADAPT | Missing, invalid, and expired access share one generic public behavior/copy. Status code, body, headers, cache, timing, and abuse parity remain O5. |
| Public donor list, Account history, claim flow, guest-email reveal | DEFER | Historical features 03–06 remain outside Slice 2; do not carry their endpoints, fields, access, or logging behavior into active acceptance. |

## Invariants

### INV-donation-01: Amount is whole Rupiah and meets the minimum

- **Statement**: A submitted amount is IDR whole Rupiah, at least Rp5.000, in Rp1 increments; Rp5.001 is valid. Stored and calculated monetary values use exact decimal representation; never use `float64`. Derived-value precision/rounding, tax behavior, wire encoding, and DB precision/scale are not decided here (O1).
- **Holds after operations**: guest donation submission and all Donation money calculations.
- **Verification**: Contract/runtime evidence rejects less than Rp5.000 and non-whole-Rupiah input, accepts Rp5.001, and preserves exact monetary values without float conversion. O1 owner review must resolve affected representation details first.

### INV-donation-02: New submission eligibility is ordered against Campaign close

- **Statement**: A new Donation is accepted only if its submission wins the atomic ordering against Campaign close while eligible. If close wins first, reject the new submission. An accepted Donation remains eligible to settle in full after Campaign close. This is D1; this invariant selects no transaction, locking, or isolation mechanism.
- **Holds after operations**: submission and the cross-domain Campaign close boundary.
- **Verification**: Runtime Testing demonstrates both orderings and accepted-pending settlement after close; Campaign's winning close reason remains unchanged.

### INV-donation-03: Guest identity and notification fields follow Slice 2 purpose

- **Statement**: Guest name is optional and not public by default. Email is optional and opt-in only for Donation status notices, separate from Account verification and not for campaign-wide updates. Ownership must be verified before sending a status or access link. A verified opt-in address is eligible for at most one status-only notice at terminal `success`/`failed`, never initial `pending`, labeled as simulation. Unverified email is deleted after its Security/PII-approved window without notice. O11's verified-address retention versus independent pending cap conflict remains unresolved; do not invent its resolution.
- **Holds after operations**: guest submission, verification, status/access delivery, terminal notice, and retention/deletion.
- **Verification**: Owner-approved policy and downstream evidence cover verification, retention, delivery retries, deletion, no notice before verification, and no initial-pending notice. O3/O11 remain gates for dependent details.

### INV-donation-04: Guest email follows established sensitive-data handling

- **Statement**: If email is collected, protect it using the established encryption/HMAC storage pattern and safe-logging rules. Do not expose details or access links before ownership verification. This invariant does not establish schemas, retention windows, or a new privacy-risk acceptance.
- **Holds after operations**: collection, storage, verification, delivery, logging, and deletion.
- **Verification**: Security/PII owner review and downstream security evidence; do not inspect or log plaintext in ordinary diagnostics.

### INV-donation-05: Guest status access is temporary, private, and status-only

- **Statement**: Guest status access uses a difficult-to-guess URL valid for 24 hours and reveals Donation status only. The donor need not use it after expiry. Missing, invalid, and expired access share one generic public behavior/copy, e.g. “Link status tidak tersedia atau mungkin kedaluwarsa.” Token representation, carrier, storage, exposure mitigations, abuse controls, response parity, and residual-risk decision remain O4/O5.
- **Holds after operations**: status credential issuance, status lookup, expiration, and public error handling.
- **Verification**: API/Security owners decide and evidence credential controls and parity across code/body/headers/cache/material timing and abuse behavior before contract details are final.

### INV-donation-06: No guest personal data is public by default

- **Statement**: Donation exposes no guest name or email publicly by default. A user-facing status response, if authorized through the temporary guest URL, contains status only. Public donor lists are deferred. Any future public shape must be reconciled separately against Product/MVP and privacy authority.
- **Holds after operations**: public projections and guest status lookup.
- **Verification**: Schema and hostile-content/runtime evidence for the active status-only surface; no donor-list acceptance is implied.

### INV-donation-07: Backend owns persisted terminal state

- **Statement**: Donation begins persisted as `pending`; only backend-controlled simulator behavior can transition it to `success` or `failed`. A donor request or UI cannot choose a terminal result, and no externally callable settlement transition may forge success. The failure path is a clearly labeled demo scenario controlled by backend configuration/fixture. Exact timing and mechanism remain O2.
- **Holds after operations**: submission, simulator processing, status reads, and recovery.
- **Verification**: Runtime/security evidence confirms no client-controlled terminal transition and truthful persisted state; timing/scenario detail requires O2 owner resolution.

### INV-donation-08: Successful settlement and full funding are one exact-once outcome

- **Statement**: A successful Donation and its full funding reflection commit atomically and exactly once. Every committed/observable outcome includes both or neither. Pending/failed Donations do not count as collected funding. A previously accepted Donation may settle after Campaign close; later settlement cannot reopen Campaign, alter its winning close reason, or clamp the amount. Funding may rise past `max_amount`.
- **Holds after operations**: successful settlement, replay, failure, concurrent settlement, and Campaign close interaction.
- **Verification**: Runtime Testing covers atomicity, rollback/failure, replay, concurrent successful Donations, accepted-pending-after-close, stable close reason, and threshold overshoot. No transaction/locking mechanism is prescribed here.

### INV-donation-09: Settlement is idempotent

- **Statement**: Repeating or replaying settlement for one Donation never creates another terminal transition or funding contribution. This is separate from request idempotency in INV-donation-10.
- **Holds after operations**: backend simulator terminal transition and funding reflection.
- **Verification**: Runtime replay and concurrency evidence shows exactly one terminal result and one full funding contribution.

### INV-donation-10: Submission retry preserves donor intent

- **Statement**: An ambiguous retry with the same idempotency key and same payload returns the original Donation. Reusing the key with a different payload is rejected. A new key starts another Donation only after deliberate donor action; the client suppresses double-click and does not rotate the key while an outcome is ambiguous. Retry-record lifetime/serialization is delivery detail still to be specified. This does not replace settlement replay protection.
- **Holds after operations**: submission and client retry handling.
- **Verification**: Contract/client/runtime evidence covers same-key/same-payload, same-key/different-payload, ambiguous retry, and deliberate fresh-key action.

### INV-donation-11: Displayed payment methods preserve sandbox truth

- **Statement**: The donation experience displays QRIS, GoPay, ShopeePay, and bank transfer. QRIS alone is an active, clearly labeled sandbox simulation; the other options are visibly unavailable and non-interactive. No method moves real money or gives usable real-payment instructions.
- **Holds after operations**: donation-method display and submission surface.
- **Verification**: Later Human rendered acceptance confirms labels, availability, and no implication of real provider settlement (R10/R11).

### INV-donation-12: Guest donation claim is deferred

- **Statement**: Historical claimability and account-linking rules are not part of Slice 2; this invariant is inactive for the current slice and does not authorize claim endpoints or retention for future claims.
- **Classification**: DEFER pending an enabling Product/MVP need and a fresh reconciliation.

### INV-donation-13: Account donation history is deferred

- **Statement**: Account donation history and preservation of claimed-donation snapshots are not part of Slice 2.
- **Classification**: DEFER; account creation is not a prerequisite for the guest trust loop.

### INV-donation-14: Guest-email reveal is deferred

- **Statement**: The historical Admin/Kurator guest-email reveal endpoint and its authorization/logging behavior are not part of Slice 2. No reveal access is authorized by this document.
- **Classification**: DEFER; reconsider only with Product/MVP and Security/PII authority.

### INV-donation-15: Historical donation-log behavior is deferred

- **Statement**: The historical `donation_logs` append-only and guest-email-reveal logging requirements are outside active Slice 2 behavior. Applicable project data/security rules still apply to any future use.
- **Classification**: DEFER with INV-donation-14; this does not weaken a generally applicable storage or audit invariant.

## State machine

### Donation sandbox status

```text
(none) -> pending -> success
                   -> failed
```

`success` and `failed` are terminal for that Donation. A failed result may be followed only by a donor's explicit new Donation with a new idempotency key. A pending state never causes automatic resubmission. Simulator timing/scenario mechanics remain O2.

## Campaign threshold and eligibility reference

Campaign owns Campaign lifecycle state and its close reason. Donation references `docs/spec/4-campaign/invariants.md#inv-campaign-13` for the narrow D1 boundary: submission eligibility is ordered atomically against close; accepted Donations remain settleable in full after close; settlement and funding commit exactly once; later settlement cannot reopen the Campaign or replace its winning close reason; funding may exceed `max_amount`. This reference does not import broader Slice 3 closure/result behavior or select a mechanism.

## Open Items retained

- **O1 `AUTHORITY_SYNC`**: shared currency-standard owner/scope and amount wire/storage/derived precision details.
- **O2**: simulator timing and backend-controlled failure-scenario mechanics.
- **O3**: email verification, retention, delivery retry, and controls.
- **O4/O5**: token exposure/abuse controls, public response parity, and residual-risk decision.
- **O8 conditional**: consumer/distribution audit before any historical operation is removed/replaced.
- **O11 `HUMAN_DECISION`**: verified terminal-email obligation versus independent pending-email cap. Do not resolve cap, alternative retention, terminal bound, timeout semantics, or risk acceptance.

## References

- `docs/product/mvp-scope.md` §§4–7
- `docs/product/mvp-delivery-slices.md` §§5–6
- `docs/ui-ux/patterns.md` §§7, 14–15
- `docs/spec/README.md`
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` §§3–13
- `docs/spec/4-campaign/invariants.md#inv-campaign-13`
