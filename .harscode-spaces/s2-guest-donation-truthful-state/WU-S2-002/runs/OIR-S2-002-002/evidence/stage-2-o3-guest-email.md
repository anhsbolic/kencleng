# Stage 2 Gap Evidence — O3 Guest Email

> Phase/Stage: Exploration / Stage 2  
> Author: `P-S2-002-OIR-002-1` (`KC-EXPLORER`)  
> Created: 2026-09-28  
> Work Unit / Run: `WU-S2-002` / `OIR-S2-002-002`  
> Model / reasoning: invocation dispatch metadata says `gpt-6-luna` / `high`; active runtime model is not independently exposed  
> Session: not exposed  
> Target revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535` (clean working tree at inspection)  
> Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`

## Current state

- Product authority is explicit: guest email is optional, opt-in, and only for donation-status notification; ownership must be verified before sending status or access links; send at most one terminal `success`/`failed` status-only notice (never initial `pending`); label it as simulated. If terminal status precedes verification, hold only within a Security/PII-approved window, then delete the unverified address without sending. Security/PII owns verification and post-terminal delivery-retry windows and retention controls. This is separate from Account email verification. Sources: `docs/product/mvp-scope.md` §Stage B, lines 114–115; `docs/product/mvp-delivery-slices.md` §5, lines 172–177.
- Techplan `TP-S2-002-007` §13 O3 preserves those semantics and says verified email may be kept only for the approved retry need; it leaves verification flow/window, post-terminal delivery-retry window, deletion/retention controls, and handling evidence open. `OIR-S2-002-001` §O3 records the same prior boundary. Current owner map assigns Security/PII and API/contract to Anhar for current Slice 2; see `.harscode-spaces/authority-map.md` rows 11–12.
- Historical Donation invariant INV-donation-04 describes `guest_email` ciphertext plus a separate HMAC lookup value. The old feature/API description instead associates omitted guest email with no claim eligibility and no campaign-result notification. These are draft/historical semantics and do not establish current Slice 2 behavior. Sources: `docs/spec/5-donation/invariants.md` INV-donation-03/04; `docs/spec/5-donation/features/02-donation-status-check.md`; `api/openapi/donation.yaml` `SubmitDonationRequest.guest_email`.
- No Donation domain package, migration, guest verification flow, terminal-notification workflow, or status-email contract is present in the current backend tree. `backend/cmd/server/main.go` wires Account and Campaign routes; it has no Donation route. Existing Account sender code is purpose-specific: `backend/internal/platform/notification/sender.go` exposes Account verification/nudge/reset methods, and `backend/internal/platform/notification/dev_sender.go` writes Account verification/reset content to a development-only simulated inbox. This is implementation evidence, not a guest notification control or an approved reuse decision.
- Root `AGENTS.md` requires the established PII encryption/HMAC convention and prohibits logging PII/secrets. Its crypto implementation path is Tier-0 fenced. No guest-email-at-rest or deletion behavior can be verified for this absent runtime.

## Requirement and gap

The approved Product behavior is settled; the delivery gap is a missing Security/PII control profile and missing runtime evidence. No verification mechanism/window, post-terminal retry window, retention/deletion schedule for verified email, encryption/HMAC mapping for the changed notification purpose, delivery-failure semantics, or end-to-end safe-logging evidence is established for guest donation. The current sources do not authorize importing Account verification TTLs or historical notification purposes.

## Sniffing lenses

- **Risk:** Unverified or retained email can disclose private donation status, or remain exposed past its purpose. A retry queue can duplicate terminal notices or outlive the approved PII window; provider errors/recipient data could leak through logs. No runtime path exists to verify these controls.
- **Edge cases:** terminal result before email verification; verification exactly at window expiry; failed delivery near retry-window expiry; duplicate worker/retry; opt-out or deletion concurrent with queued delivery; no email; email address update is not a current guest capability.
- **Miscontext:** Historical Account email verification and old Donation claim/result-notification semantics are not the current guest status-only policy. The generic Account sender interface does not prove a suitable guest flow.
- **Misleading signals:** Historical ciphertext/HMAC invariant and Account token machinery look reusable, but neither settles guest purpose, retention, consent, windows, or deletion controls. The development outbox is not evidence of production delivery or guest security.
- **Inconsistency:** Current Product requires guest status notices; historical `guest_email` contract says campaign-result notice/claim implications. Product/MVP authority wins for the active slice; specs/API require later reconciliation.

## Code and authority anchors

- `docs/product/mvp-delivery-slices.md` §5, lines 172–177 — binding guest email purpose, verification, notification, windows, retention ownership.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` §4 Q6/R5 and §13 O3 — current delivery requirement and unresolved detail.
- `backend/internal/platform/notification/sender.go` `Sender` / `FakeSender` — Account-only sender seam; does not include guest status notice.
- `backend/internal/platform/notification/dev_sender.go` `DevSender.SendVerificationEmail` / `append` — simulated Account inbox behavior and its distinct trust boundary.
- `docs/spec/5-donation/invariants.md` INV-donation-04 — historical encryption/HMAC pattern; reconcile against active purpose, do not edit Tier-0 crypto here.

Stage 2 makes no option comparison or control choice. Owner discussion can proceed only after the Stage 3 gate; unresolved timing/control detail must remain open if the owner lacks sufficient evidence.
