# Domain Invariant — Donation

> File: `docs/spec/5-donation/invariants.md`
> Status: agreed
> Human acceptance: Anhar Solehudin reviewed and accepted this current Slice 2 reconciliation as Donation and Security/PII owner on 2026-10-01; no residual risk is accepted by this status change.
> Last updated: 2026-10-04
> Product/MVP basis: `docs/product/mvp-scope.md` §§4–7; `docs/product/mvp-delivery-slices.md` §§5–6
> WU-S2-006 source amendment accepted by Anhar Solehudin on 2026-10-02 after independent Review; acceptance covers this source amendment only.
> Donation amount O1 amendment accepted by Anhar Solehudin on 2026-10-04 as Donation and Security/PII owner; see INV-donation-01 and Resolved O1.
> Donation retry-fingerprint OI8 and unavailable-Funding POST OI9 accepted by Anhar Solehudin on 2026-10-04 as Donation, Security/PII, and API owner; see INV-donation-04/10 and Resolved OI8/OI9. These decisions do not accept residual risk or authorize implementation.

## Domain summary

For Slice 2, Donation owns guest submission, persisted sandbox status, safe status revisit, and the donation-side money outcome. Account history, public donor listing, claims, guest-email reveal, and real payment rails are outside this slice. Campaign owns Campaign lifecycle state; Donation references its eligibility/threshold boundary below.

## Reconciliation classifications

| Historical material | Classification | Current treatment and authority |
|---|---|---|
| Minimum Rp5.000, whole IDR input, Rp1 increments, exact-decimal money | KEEP / ADAPT | Keep amount rules from Product/MVP. Every monetary API/wire value uses a major-unit decimal string with explicit currency code; calculation and persistence preserve exact values without `float`/`float64`. O1 is resolved for this feature as IDR whole Rupiah from Rp5.000 through Rp1.000.000.000 inclusive, stored as an exact integer Rupiah amount; derived-value precision/rounding remains undecided. |
| Submit only while Campaign eligible; full amount; threshold crossing closes Campaign | ADAPT | Apply O6 and D1 below: atomically order new submission eligibility against close; already accepted submissions remain settleable in full after close; preserve winning close reason; total may exceed threshold. No mechanism selected. |
| Guest name/email fields | ADAPT | Name optional and not public by default. Email optional and opt-in for status notices only; ownership verification and Security/PII controls remain open. A verified opted-in address remains eligible until its required terminal notice is fulfilled; Donation Delivery must provide bounded, recoverable terminalization. O3 retains the unresolved mechanics and evidence. |
| Non-expiring status token | REPLACE | Use the settled fragment-carried status URL with frontend handoff and URL cleanup, one-way HMAC verifier, hard 24-hour expiry from issuance, and status-only result. O4 retains implementation controls, evidence, and residual-risk acceptance. |
| Async 2–5 second / 5% settlement | REPLACE | Backend simulator owns persisted `pending` → `success`/`failed`. Failure is a clearly labeled demo scenario controlled by backend configuration/fixture. Timing and exact scenario mechanics remain O2. |
| Same-key retry without payload comparison; settlement replay | ADAPT | Same key/same payload returns the original Donation; same key/different payload rejects. Deliberate new Donation requires new key. This request policy is distinct from exact-once settlement/funding. |
| Guest status failure `401` / proposed `404` | REPLACE | Absent Donation and missing, wrong, or expired credential share a uniform public `404`, with identical Problem Details body, headers, and cache behavior, including `Cache-Control: private, no-store`. API expression belongs to Task 02; O5 empirical parity, timing, abuse evidence, and residual-risk acceptance remain open. |
| Public donor list, Account history, claim flow, guest-email reveal | DEFER | Historical features 03–06 remain outside Slice 2; do not carry their endpoints, fields, access, or logging behavior into active acceptance. |

## Invariants

### INV-donation-01: Amount is whole Rupiah and meets the minimum

- **Statement**: A submitted amount is IDR whole Rupiah from Rp5.000 through Rp1.000.000.000 inclusive, in Rp1 increments; Rp5.001 is valid. This is the current accepted Campaign cap range, so no individual Donation can exceed Rp1.000.000.000. Persist the Donation amount as an exact integer Rupiah value with zero fractional digits. Every monetary API/wire value remains a major-unit decimal string paired with an explicit currency code; conversion, calculation, and persistence preserve exact values end-to-end; never use `float`/`float64` or round. O1 is resolved for these supported-currency, range, fraction, precision, and storage parameters. Derived-value precision/rounding outside this amount and tax behavior are not decided here.
- **Holds after operations**: guest donation submission and all Donation money calculations.
- **Verification**: Contract/runtime evidence rejects less than Rp5.000, more than Rp1.000.000.000, and non-whole-Rupiah input; accepts Rp5.001 and the inclusive boundaries; and round-trips the exact integer Rupiah value without float conversion or fractional storage.

### INV-donation-02: New submission eligibility, cap, and capacity admission are ordered against Campaign close

- **Statement**: A new Donation is accepted only if its submission wins the atomic ordering against Campaign close while eligible, its amount is at or below the Campaign's effective `max_donation_amount`, and settled Funding plus all accepted-pending reservations plus this amount does not exceed the Campaign funding ceiling of IDR `99,999,999,999,999,999`. The per-Campaign cap is whole IDR from Rp5.000 through Rp1.000.000.000 inclusive; omitted create values default to Rp1.000.000.000, PATCH omission preserves the stored value, and the value is frozen after publication. Admission requires authoritative settled Funding to be available per Campaign INV-campaign-13; if it is unavailable, fail closed, do not treat absence as zero, and admit no new Donation until the authoritative value is restored. If close wins first, reject the new submission. After settled Funding and accepted-pending reservations are accounted for, Campaign closes with `funding_capacity_reached` when remaining capacity is less than the minimum valid Donation of Rp5.000, including exact exhaustion. Admission rejects any full Donation amount that does not fit; this does not change the separate, overshootable `max_amount` threshold. An accepted Donation remains eligible to settle in full after Campaign close. An over-cap POST or eligible capacity no-fit POST (when a smaller valid amount could fit) uses the shared generic `422 ValidationError` on `amount`; neither response discloses remaining Campaign capacity. Closed/ineligible submissions remain `409`, and existing idempotent retry behavior is preserved. This is D1 plus the accepted capacity boundary; this invariant selects no transaction, locking, or isolation mechanism.
- **Holds after operations**: submission and the cross-domain Campaign close boundary.
- **Verification**: Source review confirms cap configuration and request semantics in Campaign sources. Runtime Testing demonstrates both orderings, cap enforcement, reservation accounting under concurrent admission, residual-capacity boundaries (Rp4.999 closes; Rp5.000 remains open), exact exhaustion, rejection when an amount does not fit while a smaller valid amount does, and accepted-pending settlement after close; Campaign's winning close reason remains unchanged. Contract and runtime evidence confirms over-cap and eligible capacity no-fit responses use the generic `amount` validation shape, closed/ineligible remains `409`, existing idempotent retries are preserved, and responses do not disclose remaining capacity when a smaller valid amount could fit.

### INV-donation-03: Guest identity and notification fields follow Slice 2 purpose

- **Statement**: Guest name is optional and not public by default. Email is optional and opt-in only for Donation status notices, separate from Account verification and not for campaign-wide updates. Ownership must be verified before sending a status or access link. A verified opt-in address remains eligible for at most one status-only notice at terminal `success`/`failed`, never initial `pending`, labeled as simulation; its eligibility continues until that required notice is fulfilled. Donation Delivery must provide bounded, recoverable terminalization. Unverified email is deleted after its Security/PII-approved window without notice. O3 retains the unresolved numeric bound, architecture, timeout meaning, controls, and retention/deletion-race evidence; no residual risk is accepted.
- **Holds after operations**: guest submission, verification, status/access delivery, terminal notice, and retention/deletion.
- **Verification**: Owner-approved controls and downstream evidence cover verification, retention, bounded recoverable delivery retries/terminalization, deletion, no notice before verification, and no initial-pending notice, including retention/deletion races. O3 details and evidence remain gates; the O11 direction is settled.

### INV-donation-04: Guest email and derived retry fingerprint follow established sensitive-data handling

- **Statement**: If email is collected, protect it using the established encryption/HMAC storage pattern and safe-logging rules. Do not expose details or access links before ownership verification. A canonical-request HMAC fingerprint used for idempotency equivalence is derived sensitive data, including when it covers optional PII; do not persist raw request JSON for this comparison. Retain the fingerprint only while its idempotency record exists and delete it with that record. This does not set the idempotency-record lifetime, which remains an O3 gate, establish a new cryptographic convention, or accept privacy risk.
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

- **Statement**: A successful Donation and its full funding reflection commit atomically and exactly once. Every committed/observable outcome includes both or neither. Pending/failed Donations do not count as collected Funding. Accepted-pending amounts reserve Campaign capacity so eventual full settlement remains representable; a successful settlement consumes its reservation without changing the total admitted obligation. A previously accepted Donation may settle after Campaign close; later settlement cannot reopen Campaign, alter its winning close reason, or clamp the amount. Funding may rise past the distinct `max_amount` threshold but must remain within the finite Campaign funding ceiling. A failed accepted-pending Donation adds no Funding and releases its reservation, but cannot reopen a capacity-closed Campaign or replace the winning close reason.
- **Holds after operations**: successful settlement, replay, failure, concurrent settlement, and Campaign close interaction.
- **Verification**: Runtime Testing covers atomicity, rollback/failure, replay, concurrent successful Donations/admissions, accepted-pending-after-close, reservation release after failure without reopening, stable close reason, capacity ceiling, and threshold overshoot. No transaction/locking mechanism is prescribed here.

### INV-donation-09: Settlement is idempotent

- **Statement**: Repeating or replaying settlement for one Donation never creates another terminal transition or funding contribution. This is separate from request idempotency in INV-donation-10.
- **Holds after operations**: backend simulator terminal transition and funding reflection.
- **Verification**: Runtime replay and concurrency evidence shows exactly one terminal result and one full funding contribution.

### INV-donation-10: Submission retry preserves donor intent

- **Statement**: Idempotency keys are unique across the Donation submission endpoint. An ambiguous retry with the same key and equivalent complete accepted request payload returns the original Donation; reusing the key with a different payload is rejected. Equivalence covers the canonical route Campaign ID and every accepted request field, including optional guest name, email, and email opt-in. A keyed HMAC fingerprint over that canonical request may be used for comparison; do not persist raw request JSON, treat the fingerprint as sensitive derived data, and retain it only for the idempotency-record lifetime before deleting it with the record. The record lifetime remains bounded but numerically unspecified under O3. A new key starts another Donation only after deliberate donor action; the client suppresses double-click and does not rotate the key while an outcome is ambiguous. This does not replace settlement replay protection.
- **Holds after operations**: submission and client retry handling.
- **Verification**: Contract/client/runtime evidence covers endpoint-wide key uniqueness; same-key/equivalent complete canonical payload (including optional PII fields); same-key/different-payload rejection; ambiguous retry; deliberate fresh-key action; no raw request JSON; HMAC fingerprint deletion with its idempotency record; and an O3-approved bounded record lifetime.

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

## Open and Resolved Items

- **Resolved O1 (Anhar Solehudin, Donation and Security/PII owner, 2026-10-04)**: Slice-2 Donation accepts IDR only, whole Rupiah from Rp5.000 through Rp1.000.000.000 inclusive; persistence is an exact integer Rupiah amount with zero fractional digits. API/wire values remain major-unit decimal strings paired with IDR. This bounds Donation storage to the accepted maximum per-Campaign cap and requires exact conversion without float or rounding. This does not set a project-wide currency scale or database type for other features.
- **Resolved OI8 (Anhar Solehudin, Donation and Security/PII owner, 2026-10-04)**: Use a keyed HMAC fingerprint over the canonical complete accepted request, including optional PII, to preserve same-key payload equivalence after PII deletion. Store no raw request JSON; treat the fingerprint as sensitive derived data and retain it only while the idempotency record exists, deleting it with that record. O3 still must establish a bounded idempotency-record lifetime before implementation; this decision does not authorize changes to protected crypto code or accept residual privacy risk.
- **Resolved OI9 (Anhar Solehudin, Donation and API owner, 2026-10-04)**: A new Donation POST that cannot obtain authoritative settled Funding fails with a generic `503 Problem`, with no internal reason and no `Retry-After` header absent a proven recovery bound. No Donation is admitted in this condition. The authored API source and generated/internal consumer counterparts require a separate coordinated contract reconciliation before delivery/runtime progression.
- **O2**: simulator timing and backend-controlled failure-scenario mechanics.
- **O3**: email verification, retention, delivery retry, and controls.
- **O3 idempotency-record lifetime**: select and evidence a bounded retention duration for the idempotency record (and its derived HMAC fingerprint) before implementation; the fingerprint cannot outlive its record.
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
