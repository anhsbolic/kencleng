# Feature Spec — 02: Temporary Guest Donation Status

> File: `docs/spec/5-donation/features/02-donation-status-check.md`
> Status: agreed
> Human acceptance: Anhar Solehudin reviewed and accepted this current Slice 2 reconciliation as API and Security/PII owner on 2026-10-01; no residual risk is accepted by this status change.
> Risk tier: 1
> Domain: Donation
> Active product slice: Slice 2 — Guest Donation + Truthful Donation State
> Last updated: 2026-10-01

## Reconciliation

Historical non-expiring token and a fixed `401` contract are **REPLACE**. The settled direction requires a difficult-to-guess bearer credential in a fragment-carried status URL with frontend handoff and URL cleanup, a one-way HMAC verifier, hard expiry 24 hours from issuance, and status-only access. Absent Donation and missing, wrong, or expired credential share a uniform public `404`, identical Problem Details body, headers, and cache behavior, including `Cache-Control: private, no-store`. This feature records domain acceptance; exact authored API expression belongs to Task 02. Concrete credential generation/strength and its evidence remain open under O4; no numeric entropy/length target or empirical verification is established here.

## Feature surface

Temporary guest revisit of persisted Donation status without Account creation. The public result is status only. This document does not define the exact endpoint/payload/header expression or implementation controls. Settled credential, expiry, error, and cache directions are recorded in the acceptance below.

## Acceptance criteria

- Given a guest status URL, then it carries a difficult-to-guess bearer credential in the URL fragment, which is handed off by the frontend before the URL is cleaned up; verification uses a one-way HMAC verifier. The credential expires hard 24 hours from issuance without extension. Concrete generation/strength evidence remains open under O4; this acceptance does not select a numeric entropy/length target or claim empirical verification.
- Given a valid guest status credential before its expiry, when the guest revisits, then they can understand only the Donation status; the status URL does not reveal guest email, access link, or unrelated Donation details.
- Given an expired URL, then the donor does not need continued access after expiry.
- Given an absent Donation, missing credential, wrong credential, or expired credential, then each produces one uniform public `404` with an identical Problem Details body, headers, and cache behavior, including `Cache-Control: private, no-store`. The user-facing behavior may use the generic copy **“Link status tidak tersedia atau mungkin kedaluwarsa.”**
- Given a verified status state, then the result uses the source-first label family **“Hasil simulasi donasi: berhasil/gagal”** where terminal result labeling applies. Pending uses **“Menunggu hasil simulasi”**, with no time estimate or real-payment instruction.
- Given optional email/Account information after status access, then it is informational and non-coercive, does not gate the guest flow, and does not promise unavailable functionality.
- Given missing/wrong/expired credential or absent Donation, then the authored API must express the settled `404`, identical Problem Details body, headers, and cache behavior. Exact API expression is assigned to Task 02; the contract direction does not establish empirical response/timing parity or abuse evidence (O5).
- Given a bearer URL that may leak through browser history, referrer, logs, cache, or forwarding, then Security/API owners must select and evidence applicable mitigations and make the residual-risk decision (O4). Fragment handoff, URL cleanup, HMAC verification, and TTL are settled directions, not proof that exposure, key/comparison, expiry-enforcement, or abuse controls work.

## Error cases

| Condition | Public behavior |
|---|---|
| Missing status credential | Uniform public `404`; identical Problem Details body, headers, and cache behavior as wrong/expired credential and absent Donation; `Cache-Control: private, no-store`. |
| Wrong credential | Same uniform public `404` contract as missing/expired credential and absent Donation. |
| Absent Donation | Same uniform public `404` contract as missing/wrong/expired credential. |
| Expired credential | Same uniform public `404` contract as missing/wrong credential and absent Donation. |
| Valid credential, before hard expiry at 24 hours from issuance | Status-only result. |

The exact authored Problem Details/API expression belongs to Task 02. Empirical response/timing parity, abuse behavior, and Security/PII residual-risk acceptance remain open O4/O5 evidence obligations; this table does not claim them proven.

## Applicable invariants

- `docs/spec/5-donation/invariants.md#inv-donation-05`
- `docs/spec/5-donation/invariants.md#inv-donation-06`
- `docs/spec/5-donation/invariants.md#inv-donation-07`

## Threat breakdown

| Threat | Feature-level mitigation | Evidence required |
|---|---|---|
| Credential guessing/theft | Require a difficult-to-guess bearer credential; use the fragment-carried status URL with frontend handoff and URL cleanup, one-way HMAC verifier, hard 24-hour expiry from issuance, and status-only access. | O4 still requires concrete credential generation/strength choice and evidence, plus key/comparison, expiry-enforcement, browser/infrastructure exposure, and abuse controls and a residual-risk decision. No numeric entropy/length target or empirical verification is established here. |
| Donation existence enumeration | Uniform public `404`, identical Problem Details body, headers, and cache behavior for absent Donation and missing/wrong/expired credential; `Cache-Control: private, no-store`. | O5 empirical response/timing parity and abuse evidence; contract direction is not proof. |
| Token exposure in browser or infrastructure | No assumed mitigation beyond product direction. | Security/API owner control and residual-risk review. |
| Excessive lookup abuse | No control selected here. | O4/O5 abuse-control decision and evidence. |
| Excessive information disclosure | Status only; no guest email or unrelated Donation details. | Contract and runtime hostile-response verification. |

## Risk tier and verification boundary

Tier 1 because an unauthenticated bearer credential provides access to private status and creates anti-enumeration concerns. Independent review, API/Security owner review, applicable Human risk decision, and later runtime security evidence remain required. This spec task does not settle these gates.

## Open items

- **O4**: settled fragment URL/frontend handoff/cleanup and one-way HMAC direction; implementation evidence for key/comparison, expiry enforcement, browser/referrer/log/cache mitigations, abuse controls, and residual-risk decision remains open.
- **O5**: settled uniform public `404`, identical Problem Details body/header/cache behavior including `Cache-Control: private, no-store`; empirical response/timing parity and abuse behavior remain open.
- **O8 conditional**: consumer/distribution audit only before historical operation removal/replacement.

## References

- `docs/spec/5-donation/invariants.md`
- `docs/spec/5-donation/threat-model.md`
- `docs/spec/5-donation/tasks.md#task-02--safe-guest-status-revisit`
- `docs/product/mvp-delivery-slices.md` §5
- `docs/ui-ux/patterns.md` §7
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` §§3–13
