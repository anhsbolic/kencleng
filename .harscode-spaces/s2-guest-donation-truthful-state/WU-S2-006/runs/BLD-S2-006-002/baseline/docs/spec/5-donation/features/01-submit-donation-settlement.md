# Feature Spec — 01: Guest Donation Submission and Sandbox Result

> File: `docs/spec/5-donation/features/01-submit-donation-settlement.md`
> Status: agreed
> Human acceptance: Anhar Solehudin reviewed and accepted this current Slice 2 reconciliation as Donation and Security/PII owner on 2026-10-01; no residual risk is accepted by this status change.
> Risk tier: 1
> Domain: Donation
> Active product slice: Slice 2 — Guest Donation + Truthful Donation State
> Last updated: 2026-10-01

## Reconciliation

Historical public-or-authenticated submission and 2–5 second / 5% settlement are **ADAPT/REPLACE**. This feature covers the guest path and persisted sandbox result from current Product/MVP. Registered-donor behavior and real rails are not in this acceptance. The historical async mechanism is not authoritative.

## Feature surface

Guest donation submission, backend-controlled simulator, and the Campaign eligibility/funding boundary. Exact API paths, request/response fields, amount encoding, and simulator control are left to the authorized contract task and relevant owners; this document does not define them.

## Acceptance criteria

- Given a guest without an Account, when a valid donation is submitted to an eligible Campaign, then the Donation is persisted as `pending` and the guest flow does not require account creation.
- Given an amount, when validation runs, then only IDR whole Rupiah amounts at least Rp5.000 in Rp1 increments are accepted; Rp5.001 is valid. Every monetary API/wire value uses a major-unit decimal string with an explicit currency code; calculation and persistence preserve exact decimal values end-to-end without `float`/`float64`. O1 remains open only for supported currencies and concrete range, fraction, precision/scale, and storage parameters. Derived-value precision/rounding remains undecided.
- Given the method display, then QRIS, GoPay, ShopeePay, and bank transfer may be shown; QRIS is the only active, clearly labeled sandbox simulation. Other methods are visibly unavailable and non-interactive. Nothing implies real settlement or gives usable real-payment instructions.
- Given a backend simulator, when it processes a Donation, then only backend-controlled behavior chooses `success` or `failed`; a failure can occur only through a clearly labeled demo scenario controlled by backend configuration/fixture, never by donor choice/request. Exact timing/scenario mechanics remain O2.
- Given `pending`, then the copy is “Menunggu hasil simulasi,” with no estimate and no instruction to make a real payment. Pending never triggers automatic resubmission. Given `failed`, the donor may explicitly start another Donation using a new key.
- Given an ambiguous submission outcome, when the client retries with the same key and payload, then the original Donation is returned; same key with a different payload is rejected. Client prevents double-click and does not rotate key while ambiguous. New key means deliberate new intent. Retry record lifetime/serialization remains contract detail.
- Given a submission racing Campaign close, then eligibility is ordered atomically against close: close-first rejects the new submission; submission-first accepts it. This expresses D1 and does not prescribe a mechanism.
- Given an accepted Donation still pending when Campaign closes, when the simulator later succeeds, then the full amount remains settleable and Donation success plus full funding reflection commit atomically and exactly once. Later settlement does not reopen Campaign or change its winning close reason; funding may exceed `max_amount`.
- Given an unsuccessful or still-pending Donation, then it does not count as collected funding. Settlement replay cannot add funding twice. Request idempotency and settlement idempotency are separate guarantees.
- Given optional guest fields, then name is optional and not public by default. Email is optional and opt-in for Donation status only, separate from Account email verification and not for campaign-wide updates.
- Given an email notice or access link, then ownership is verified before sending. At most one status-only notice may be sent at terminal `success`/`failed`, never initial `pending`, and it is clearly labeled as simulation. A verified opted-in address remains eligible until its required terminal notice is fulfilled; Donation Delivery must provide bounded, recoverable terminalization. Unverified email is held/deleted only under Security/PII-approved controls/windows.
- O3 retains the unresolved numeric terminalization bound, architecture, timeout meaning, security controls, retention/deletion-race evidence, and residual-risk gate. This settled O11/D19 direction does not accept residual risk or specify a retention/deletion mechanism.

### Exact Design direction for dependent UI/notice

- Terminal label family: **“Hasil simulasi donasi: berhasil/gagal”**.
- Optional-email label: **“Kirim pemberitahuan status donasi melalui email (opsional)”**.
- Helper: **“Verifikasi email dalam 24 jam sejak alamat dicatat. Jika tidak diverifikasi, alamat dihapus dan pemberitahuan tidak dikirim. Donasi tetap berjalan.”**

These choices are carried from Q11/R11 and D16 exactly; this spec does not establish rendered visual acceptance or resolve O3 delivery controls/mechanics. Rendered acceptance is later Human review.

## Error and recovery behavior

| Condition | Expected behavior |
|---|---|
| Amount below minimum or not whole Rupiah | Reject according to final contract validation; exact status/code is not defined here. |
| Campaign close wins before submission eligibility | Reject the new Donation; do not specify transport shape here. |
| Same idempotency key with changed payload | Reject; exact API problem shape belongs to contract reconciliation. |
| Simulator remains pending | Show “Menunggu hasil simulasi”; no timing promise, payment instruction, or automatic resubmission. |
| Simulator returns failed | Show simulation failure truthfully; allow a deliberate new Donation with a new key. |

## Applicable invariants

- `docs/spec/5-donation/invariants.md#inv-donation-01` through `#inv-donation-04`
- `docs/spec/5-donation/invariants.md#inv-donation-07` through `#inv-donation-11`
- `docs/spec/4-campaign/invariants.md#inv-campaign-13` (narrow D1 eligibility/threshold boundary only)

## Threat breakdown

| Threat | Feature-level mitigation | Evidence required |
|---|---|---|
| Forged terminal state | Backend simulator owns result; no client-callable settlement transition. | Independent review and downstream route/security evidence. |
| Duplicate submission or changed-payload replay | Same-key rules; client retains key through ambiguity and suppresses double-click. | Contract/client/runtime retry evidence. |
| Partial or repeated funding | Atomic success/full-funding outcome, exact-once settlement. | Downstream failure, replay, and concurrency Testing; mechanism not selected. |
| Ineligible or lost accepted Donation during close | D1 ordering and accepted-pending full settlement. | Runtime concurrency/integration Testing. |
| PII exposure or premature notification | Optional opt-in, verification before delivery, established encryption/HMAC and safe logs; preserve verified-address eligibility until terminal-notice fulfillment. | O3 controls, retention/deletion-race evidence, and Security/PII review. |

## Risk tier and verification boundary

Tier 1 because this feature controls donation status, money reflection, guest PII, and Campaign close ordering. Independent Code Review, applicable Human/domain-owner review, and specialized downstream runtime Testing remain required. This documentation Run does not establish implementation behavior.

## Open questions and gates

- O1: supported currencies and concrete range, fraction, precision/scale, and storage parameters; the shared amount representation direction is settled.
- O2: simulator timing and backend-controlled scenario mechanics.
- O3: verification, retention, delivery retry, and controls.
- O11/D19 direction is settled: verified opt-in remains eligible until required terminal notice fulfillment, with bounded, recoverable terminalization. O3 controls and evidence remain open as listed above.
- O8 is conditional only if a historical operation is proposed for removal/replacement; consumer/distribution evidence is required first.

## References

- `docs/spec/5-donation/invariants.md`
- `docs/spec/5-donation/threat-model.md`
- `docs/spec/5-donation/tasks.md#task-01--guest-submit-and-truthful-sandbox-result`
- `docs/product/mvp-scope.md` §§4–7
- `docs/product/mvp-delivery-slices.md` §5
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` §§3–13
