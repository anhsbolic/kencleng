# Domain Invariant — Donation

> File: `docs/spec/5-donation/invariants.md`
> Status: draft — Slice 2 reconciliation; requires Donation, Security/PII, and applicable Human review before `agreed`
> Last updated: 2026-10-01
> Product/MVP basis: `docs/product/mvp-scope.md` §§4–7; `docs/product/mvp-delivery-slices.md` §§5–6

## Domain summary

For Slice 2, Donation owns guest submission, persisted sandbox status, safe status revisit, and the donation-side money outcome. Account history, public donor listing, claims, guest-email reveal, and real payment rails are outside this slice. Campaign owns Campaign lifecycle state; Donation references its eligibility/threshold boundary below.

## Reconciliation classifications

| Historical material | Classification | Current treatment and authority |
|---|---|---|
| Minimum Rp5.000, whole IDR input, Rp1 increments, exact-decimal money | KEEP / ADAPT | Keep amount rules from Product/MVP. Every monetary API/wire value uses a major-unit decimal string with explicit currency code; calculation and persistence preserve exact decimals without `float`/`float64`. O1 remains open only for supported currencies and concrete range, fraction, precision/scale, and storage parameters; derived-value precision/rounding remains undecided. |
| Submit only while Campaign eligible; full amount; threshold crossing closes Campaign | ADAPT | Apply O6 and D1 below: atomically order new submission eligibility against close; already accepted submissions remain settleable in full after close; preserve winning close reason; total may exceed threshold. No mechanism selected. |
| Guest name/email fields | ADAPT | Name optional and not public by default. Email optional and opt-in for status notices only; ownership verification and Security/PII controls remain open. A verified opted-in address remains eligible until its required terminal notice is fulfilled; Donation Delivery must provide bounded, recoverable terminalization. O3 retains the unresolved mechanics and evidence. |
| Non-expiring status token | REPLACE | Use the settled fragment-carried status URL with frontend handoff and URL cleanup, one-way HMAC verifier, hard 24-hour expiry from issuance, and status-only result. O4 retains implementation controls, evidence, and residual-risk acceptance. |
| Async 2–5 second / 5% settlement | REPLACE | Backend simulator owns persisted `pending` → `success`/`failed`. Failure is a clearly labeled demo scenario controlled by backend configuration/fixture. Timing and exact scenario mechanics remain O2. |
| Same-key retry without payload comparison; settlement replay | ADAPT | Same key/same payload returns the original Donation; same key/different payload rejects. Deliberate new Donation requires new key. This request policy is distinct from exact-once settlement/funding. |
| Guest status failure `401` / proposed `404` | REPLACE | Absent Donation and missing, wrong, or expired credential share a uniform public `404`, with identical Problem Details body, headers, and cache behavior, including `Cache-Control: private, no-store`. API expression belongs to Task 02; O5 empirical parity, timing, abuse evidence, and residual-risk acceptance remain open. |
| Public donor list, Account history, claim flow, guest-email reveal | DEFER | Historical features 03–06 remain outside Slice 2; do not carry their endpoints, fields, access, or logging behavior into active acceptance. |

## Invariants

### INV-donation-01: Amount is whole Rupiah and meets the minimum

- **Statement**: A submitted amount is IDR whole Rupiah, at least Rp5.000, in Rp1 increments; Rp5.001 is valid. Every monetary API/wire value is a major-unit decimal string paired with an explicit currency code. Calculation and persistence preserve exact decimal values end-to-end; never use `float`/`float64`. O1 remains open only for supported currencies and concrete range, fraction, precision/scale, and storage parameters. Derived-value precision/rounding and tax behavior are not decided here.
- **Holds after operations**: guest donation submission and all Donation money calculations.
- **Verification**: Contract/runtime evidence rejects less than Rp5.000 and non-whole-Rupiah input, accepts Rp5.001, and preserves exact monetary values without float conversion. O1 owner review resolves only the remaining currency/range/fraction/precision/storage parameters before those details are finalized.

### INV-donation-02: New submission eligibility is ordered against Campaign close

- **Statement**: A new Donation is accepted only if its submission wins the atomic ordering against Campaign close while eligible. If close wins first, reject the new submission. An accepted Donation remains eligible to settle in full after Campaign close. This is D1; this invariant selects no transaction, locking, or isolation mechanism.
- **Holds after operations**: submission and the cross-domain Campaign close boundary.
- **Verification**: Runtime Testing demonstrates both orderings and accepted-pending settlement after close; Campaign's winning close reason remains unchanged.

### INV-donation-03: Guest identity and notification fields follow Slice 2 purpose

- **Statement**: Guest name is optional and not public by default. Email is optional and opt-in only for Donation status notices, separate from Account verification and not for campaign-wide updates. Ownership must be verified before sending a status or access link. A verified opt-in address remains eligible for at most one status-only notice at terminal `success`/`failed`, never initial `pending`, labeled as simulation; its eligibility continues until that required notice is fulfilled. Donation Delivery must provide bounded, recoverable terminalization. Unverified email is deleted after its Security/PII-approved window without notice. O3 retains the unresolved numeric bound, architecture, timeout meaning, controls, and retention/deletion-race evidence; no residual risk is accepted.
- **Holds after operations**: guest submission, verification, status/access delivery, terminal notice, and retention/deletion.
- **Verification**: Owner-approved controls and downstream evidence cover verification, retention, bounded recoverable delivery retries/terminalization, deletion, no notice before verification, and no initial-pending notice, including retention/deletion races. O3 details and evidence remain gates; the O11 direction is settled.

### INV-donation-04: Guest email follows established sensitive-data handling

- **Statement**: If email is collected, protect it using the established encryption/HMAC storage pattern and safe-logging rules. Do not expose details or access links before ownership verification. This invariant does not establish schemas, retention windows, or a new privacy-risk acceptance.
- **Holds after operations**: collection, storage, verification, delivery, logging, and deletion.
- **Verification**: Security/PII owner review and downstream security evidence; do not inspect or log plaintext in ordinary diagnostics.

### INV-donation-05: Guest status access is temporary, private, and status-only

- **Statement**: Guest status access uses a difficult-to-guess bearer credential in the settled fragment-carried status URL, with frontend handoff and URL cleanup, verified through a one-way HMAC verifier, with hard expiry 24 hours from issuance and status-only result. The donor need not use it after expiry. Absent Donation and missing, wrong, or expired credentials share a uniform public `404`, identical Problem Details body, headers, and cache behavior, including `Cache-Control: private, no-store`. Exact authored API expression belongs to Task 02. Concrete credential generation/strength evidence, browser history/referrer/log/cache exposure protections, key and comparison controls, expiry enforcement, abuse controls, empirical response/timing parity, and Security/PII residual-risk acceptance remain unproven O4/O5 obligations; this invariant does not select a numeric entropy or length target or claim the property is empirically verified.
- **Holds after operations**: status credential issuance, status lookup, expiration, and public error handling.
- **Verification**: Task 02 expresses the settled API direction and difficult-to-guess credential requirement. API/Security owners must still select and evidence concrete credential generation/strength and implementation controls, expiry enforcement, abuse handling, empirical response/timing parity, and applicable residual-risk acceptance; these are not established by this invariant.

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

- **O1**: only supported currencies and concrete range, fraction, precision/scale, and storage parameters remain open; the shared major-unit decimal-string plus explicit-currency representation and exact-decimal direction are settled.
- **O2**: simulator timing and backend-controlled failure-scenario mechanics.
- **O3**: email verification, retention, delivery retry, and controls.
- **O4/O5**: settled status-link/404 directions are recorded in INV-donation-05; browser/infrastructure exposure, key/comparison controls, expiry enforcement, abuse controls, empirical response/timing parity, and residual-risk decision remain open.
- **O8 conditional**: consumer/distribution audit before any historical operation is removed/replaced.
- **O11/D19 (settled direction)**: verified opted-in address remains eligible until required terminal notice is fulfilled; Donation Delivery provides bounded, recoverable terminalization. O3 retains the numeric bound, architecture, timeout meaning, security controls, retention/deletion-race evidence, and residual-risk gate.

## References

- `docs/product/mvp-scope.md` §§4–7
- `docs/product/mvp-delivery-slices.md` §§5–6
- `docs/ui-ux/patterns.md` §§7, 14–15
- `docs/spec/README.md`
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` §§3–13
- `docs/spec/4-campaign/invariants.md#inv-campaign-13`
