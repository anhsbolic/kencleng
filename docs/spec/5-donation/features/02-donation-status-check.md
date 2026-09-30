# Feature Spec — 02: Temporary Guest Donation Status

> File: `docs/spec/5-donation/features/02-donation-status-check.md`
> Status: draft — Slice 2 reconciliation; API/Security and applicable Human review required before `agreed`
> Risk tier: 1
> Domain: Donation
> Active product slice: Slice 2 — Guest Donation + Truthful Donation State
> Last updated: 2026-09-30

## Reconciliation

Historical non-expiring token and a fixed `401` contract are **REPLACE/ADAPT**. Current Product/MVP requires a difficult-to-guess guest URL valid for 24 hours and status-only access. Exact path, credential transport/control, and response parity are not settled here.

## Feature surface

Temporary guest revisit of persisted Donation status without Account creation. The public result is status only. This document does not choose endpoint, token representation, carrier, storage, headers, cache rules, response code, or timing controls.

## Acceptance criteria

- Given a valid guest status URL within 24 hours of issuance, when the guest revisits, then they can understand only the Donation status; the status URL does not reveal guest email, access link, or unrelated Donation details.
- Given an expired URL, then the donor does not need continued access after expiry.
- Given a missing, invalid, or expired URL, then each produces one generic public behavior and copy, for example: **“Link status tidak tersedia atau mungkin kedaluwarsa.”**
- Given a verified status state, then the result uses the source-first label family **“Hasil simulasi donasi: berhasil/gagal”** where terminal result labeling applies. Pending uses **“Menunggu hasil simulasi”**, with no time estimate or real-payment instruction.
- Given optional email/Account information after status access, then it is informational and non-coercive, does not gate the guest flow, and does not promise unavailable functionality.
- Given missing/invalid/expired requests that may reveal Donation existence, then API/Security owners must define and evidence response-code, body, headers, cache behavior, material timing, and abuse parity. Shared frontend copy alone is not evidence of parity (O5).
- Given a bearer URL that may leak through browser history, referrer, logs, cache, or forwarding, then Security/API owners must select and evidence mitigations and make the applicable residual-risk decision (O4). TTL and token unpredictability alone do not close this item.

## Error cases

| Condition | Public behavior |
|---|---|
| Missing status credential | Same generic behavior/copy as invalid and expired. |
| Invalid credential or unknown Donation | Same generic behavior/copy as missing and expired. |
| Expired credential | Same generic behavior/copy as missing and invalid. |
| Valid credential, within 24 hours | Status-only result. |

Transport-level equality is an O5 owner decision and evidence obligation, not specified by this table.

## Applicable invariants

- `docs/spec/5-donation/invariants.md#inv-donation-05`
- `docs/spec/5-donation/invariants.md#inv-donation-06`
- `docs/spec/5-donation/invariants.md#inv-donation-07`

## Threat breakdown

| Threat | Feature-level mitigation | Evidence required |
|---|---|---|
| Credential guessing/theft | Difficult-to-guess URL, 24-hour validity, status-only access. | O4 decision/evidence for token controls and exposure mitigations. |
| Donation existence enumeration | Generic public behavior/copy for missing/invalid/expired. | O5 parity evidence across response attributes and material timing. |
| Token exposure in browser or infrastructure | No assumed mitigation beyond product direction. | Security/API owner control and residual-risk review. |
| Excessive lookup abuse | No control selected here. | O4/O5 abuse-control decision and evidence. |
| Excessive information disclosure | Status only; no guest email or unrelated Donation details. | Contract and runtime hostile-response verification. |

## Risk tier and verification boundary

Tier 1 because an unauthenticated bearer credential provides access to private status and creates anti-enumeration concerns. Independent review, API/Security owner review, applicable Human risk decision, and later runtime security evidence remain required. This spec task does not settle these gates.

## Open items

- **O4**: token carrier/storage/lifecycle/comparison, browser/referrer/log/cache mitigations, abuse controls, residual-risk decision.
- **O5**: response code/body/header/cache/material-timing parity and abuse behavior.
- **O8 conditional**: consumer/distribution audit only before historical operation removal/replacement.

## References

- `docs/spec/5-donation/invariants.md`
- `docs/spec/5-donation/threat-model.md`
- `docs/spec/5-donation/tasks.md#task-02--safe-guest-status-revisit`
- `docs/product/mvp-delivery-slices.md` §5
- `docs/ui-ux/patterns.md` §7
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` §§3–13
