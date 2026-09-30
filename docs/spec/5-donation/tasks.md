# Task List — Donation

> File: `docs/spec/5-donation/tasks.md`
> Status: draft — Slice 2 reconciliation; owner review required before `agreed`
> Last updated: 2026-09-30
> Active product slice: Slice 2 — Guest Donation + Truthful Donation State

## Delivery outcome

Provide a truthful guest donation flow with persisted backend-owned sandbox state, safe temporary status revisit, and exact-once successful funding. Account and operational breadth remain outside the baseline Slice 2 path. Campaign coordination is limited to the approved threshold/eligibility boundary.

## Reconciliation of historical tasks

| Historical task | Classification | Slice 2 disposition |
|---|---|---|
| 01 Submit donation & async settlement | ADAPT | Active as guest submission plus backend-controlled sandbox status and atomic full funding; remove historical timing/probability, real rails, and mechanism assumptions. |
| 02 Donation status check | ADAPT | Active as temporary 24-hour, difficult-to-guess, status-only guest access; preserve generic missing/invalid/expired behavior while O4/O5 controls remain gated. |
| 03 Public donor list | DEFER | Not required by approved guest trust loop. No donor-list operation, public name, or correlation behavior is in Slice 2. |
| 04 Account donation history | DEFER | Account is not a prerequisite; no history surface in Slice 2. |
| 05 Guest donation claim | DEFER | Claim/account-linking is not in current MVP slice. |
| 06 Guest-email reveal | DEFER | Admin/Kurator reveal is not in current MVP slice; no such access is authorized here. |

## Active tasks

| # | Task | Endpoint / surface | Tier | Dependencies | Related invariants |
|---|---|---|---|---|---|
| 01 | Guest submit and truthful sandbox result | Guest submission, backend simulator, Donation status | 1 | Product/MVP; O1/O2/O3 gates for dependent detail; Campaign eligibility boundary | INV-donation-01–04, 07–11 |
| 02 | Safe guest status revisit | Temporary guest URL and status-only response | 1 | Task 01; O4/O5 Security/API controls and residual-risk review | INV-donation-05–06 |

These tasks are serial at the contract level because status behavior depends on Donation submission and state. They do not authorize API, implementation, migration, or generated-artifact changes in this spec reconciliation.

## Task 01 — Guest submit and truthful sandbox result

**Acceptance**:

- Guest submission does not require an Account.
- Amount is IDR whole Rupiah, minimum Rp5.000, Rp1 increments; Rp5.001 is valid. Stored/calculated amounts use exact decimal and never `float64`. O1 gates amount representation, storage scale, and derived precision/rounding.
- QRIS, GoPay, ShopeePay, and bank transfer may be displayed; QRIS alone is active as a clearly labeled sandbox simulation. Other methods are unavailable/non-interactive. No real payment instructions or settlement implication.
- New submission eligibility is atomically ordered against Campaign close. If submission wins, it is accepted and may settle in full after close; if close wins, reject. Do not select mechanism or expand into Slice 3.
- Persist `pending`; only backend-controlled simulator behavior resolves `success`/`failed`. Failure is a clearly labeled demo scenario controlled by backend configuration/fixture. Pending copy is “Menunggu hasil simulasi,” with no time estimate or payment instruction. Pending never auto-resubmits; a failed result allows a donor-initiated new Donation with a new key. O2 timing/scenario details remain open.
- Same key and same payload after an ambiguous retry returns the original Donation; same key with different payload rejects. Double-click is suppressed; key is retained while ambiguous; a fresh key means deliberate new intent. Keep request idempotency distinct from settlement replay.
- Successful Donation status and its full funding reflection commit atomically and exactly once. Accepted pending Donations remain settleable after Campaign close. Later settlement cannot reopen Campaign or change its winning close reason; funding may exceed `max_amount`.
- Guest name is optional and not public by default. Email is optional, opt-in, and status-only. Verify ownership before status/access delivery; send at most one terminal status-only simulation-labeled notice, never initial pending. Security/PII defines controls/windows. Preserve O11 as `HUMAN_DECISION`; do not select cap, alternate retention, terminal bound, timeout meaning, or risk acceptance.

## Task 02 — Safe guest status revisit

**Acceptance**:

- Guest status URL is difficult to guess, valid for 24 hours, and reveals status only. Donor need not revisit it after expiry.
- Missing, wrong/invalid, and expired link use one generic public behavior/copy, e.g. “Link status tidak tersedia atau mungkin kedaluwarsa.”
- O4 remains active for token carrier/storage/lifecycle, exposure mitigations, abuse controls, and residual-risk decision. O5 remains active for response code/body/header/cache/material-timing parity and abuse behavior. Do not treat matching copy as proof of technical parity.
- Optional email/Account information after status access is non-coercive, informational, does not gate guest access, and does not promise unavailable benefits.

## Verification ownership and scope

This task list records later owner review and runtime Testing obligations from Approved Techplan §12; spec editing alone does not satisfy them. Tier 1 requires independent review and applicable Human/domain-owner review before affected specs become `agreed`. Runtime evidence includes exact-money handling, idempotency, simulator authority, status privacy, atomic funding, replay/concurrency, D1 submit/close ordering, accepted-pending settlement after close, close-reason stability, and threshold overshoot. O1–O5/O11 detail remains owner-gated. O8 is conditional and requires consumer/distribution evidence before any historical API operation is removed or replaced.

## Status tracker

| # | Status | Notes |
|---|---|---|
| 01 | draft / owner review required | Domain-spec reconciliation only; runtime implementation/testing not claimed. |
| 02 | draft / owner review required | O4/O5 technical controls and residual-risk decision remain open. |
| 03–06 | deferred | Historical feature files retained as evidence with explicit defer status. |

## References

- `docs/spec/5-donation/invariants.md`
- `docs/spec/5-donation/threat-model.md`
- `docs/spec/5-donation/features/01-submit-donation-settlement.md`
- `docs/spec/5-donation/features/02-donation-status-check.md`
- `docs/product/mvp-scope.md` §§4–7
- `docs/product/mvp-delivery-slices.md` §§5–6
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` §§3–13
