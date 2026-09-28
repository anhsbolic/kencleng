# Stage 2 Gap Evidence — O3 Verified Email Retention While Pending

> Phase/Stage: Exploration / Stage 2  
> Author: `P-S2-002-OIR-003-1` (`KC-EXPLORER`)  
> Created: 2026-09-28  
> Model / reasoning: invocation dispatch metadata `gpt-6-luna` / `high`; active runtime model is not independently exposed  
> Session: FRESH per invocation; Session ID not exposed  
> Work Unit / Run: `WU-S2-002` / `OIR-S2-002-003`  
> Target revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535` (`HEAD` at inspection)  
> Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`

## Current state

- Current Product requires guest email to be optional and opt-in for donation-status notices only; ownership verification precedes sending any status or access link. At most one status-only message is sent at terminal `success`/`failed`, never at initial `pending`, and it must identify a simulation. If terminal state precedes verification, unverified email is held only within an approved verification window and deleted without notice at expiry. Product assigns verification/retry windows and retention controls to Security/PII. Sources: `docs/product/mvp-scope.md` Stage B; `docs/product/mvp-delivery-slices.md` §5.
- `OIR-S2-002-002/resolution-brief.md` O3 records the Human decision (2026-09-28): retain verification; use a 24-hour verification window from email capture and a 24-hour retry window from terminal state; delete unverified email at verification expiry; delete verified email after successful delivery or retry-window expiry. Maximum retention of a verified address while Donation stays `pending` was explicitly deferred until O2 timing evidence or a separate Security/PII cap is available. No residual privacy risk was accepted.
- Current Techplan `TP-S2-002-007` §13 O3 requires keeping verified email only for the approved notification/retry purpose and lists retention controls as active owner work. It does not provide a pending-state maximum.
- Current backend contains no Donation domain, email-verification flow, terminal-notification workflow, or Donation sender/queue. The Account notification sender/outbox is Account-specific and does not prove a Guest Donation purpose, retention window, delivery contract, or deletion behavior. Prior OIR Stage 2 evidence `OIR-S2-002-002/evidence/stage-2-o3-guest-email.md` documents the targeted Account component anchors and current absence.
- The current O2 evidence in `evidence/stage-2-o2-simulator.md` finds no implemented Donation simulator and no approved maximum terminal timing. Therefore the existing O3 deletion rule for terminal states does not bound verified-email retention during an indefinitely pending state.
- Current authority map names Anhar Solehudin as both current Slice 2 Donation delivery/domain owner and Security/PII owner. The owner attribution names who can make a bounded current-slice decision; it does not itself set the bound.

## Requirement and gap

**Requirement:** Verified email may be used only for the approved terminal-status notification and its approved retry need. Product/security has already bounded verification and post-terminal retry windows and deletion, but a maximum is still required for the period before a terminal result. The amount of pending-state retention is not a public timing promise; no SLA is authorized.

**Gap:** No evidence establishes a maximum time from initial pending state to terminal state. The prior Human decision deliberately deferred a pending-retention cap pending O2 timing evidence or an explicit Security/PII upper bound. Since current delivery cannot support a simulator terminal deadline, the current evidence still cannot calculate a finite retention maximum from O2. A separate Security/PII cap remains the other owner route identified in the prior decision, but it has not been decided in this Run's Stage 2.

This evidence records the dependency and unresolved condition only; no cap or retention implementation is selected here.

## Sniffing lenses

- **Risk:** If pending is unbounded and the address remains usable/stored while pending, verified PII can persist beyond a necessary or reviewable duration. Deleting it early may prevent the already-approved terminal notification; retaining it indefinitely conflicts with purpose limitation and leaves privacy exposure open. No residual risk acceptance exists.
- **Edge cases:** pending crosses the proposed maximum at the same time as a terminal update; simulator result never arrives; Donation enters terminal state just before a pending-retention cap; terminal delivery/retry reaches its separate 24-hour limit; deletion and notification retry race; verified email exists but status remains pending through restart/recovery. These require later owner/spec/runtime treatment; current behavior is absent.
- **Miscontext:** The 24-hour post-terminal retry window begins at terminal state and cannot serve as a cap measured from initial `pending`. The 24-hour verification window governs unverified email and does not by itself bound verified email retained while pending. The 24-hour status URL lifetime is also a distinct credential rule, not an email-retention duration.
- **Misleading signals:** Historical Donation PII encryption/HMAC invariants and Account email verification TTL may look like complete retention controls. They establish neither this guest purpose's retention bound nor its expiry/deletion behavior. The generic Account sender/outbox is not Guest Donation runtime evidence.
- **Inconsistency:** Historical Donation specs/API describe email for campaign outcome/claim semantics, while current Product confines it to opt-in donation-status notifications. Historical simulator timing describes delay but is expressly non-authoritative and has no demonstrated maximum in current runtime. Product and the Human O3 decision govern the active slice; contract/spec reconciliation remains downstream.

## Code and authority anchors

- `docs/product/mvp-delivery-slices.md` §5, guest email paragraphs — purpose, verification requirement, terminal-only send, owner of windows and controls.
- `runs/OIR-S2-002-002/resolution-brief.md` O3 and Cross-item dependencies — exact Human decision and explicit deferral of pending maximum.
- `runs/OIR-S2-002-002/handoff.md` Observed decisions and remaining concerns — prior owner decision and required next route.
- `runs/TP-S2-002-007/techplan.md` §13 O2/O3 — simulator-timing dependency and active O3 retention detail.
- `.harscode-spaces/authority-map.md` Donation delivery/domain and Security/PII rows — owner and current Slice 2 scope.
- `backend/internal/platform/notification/sender.go` `Sender` / `FakeSender`, and `dev_sender.go` `DevSender` — Account-specific sender seam only, as captured in prior OIR evidence; not guest Donation delivery/retention behavior.
- `docs/spec/5-donation/invariants.md` INV-donation-04 — historical encryption/HMAC pattern; it does not decide this pending-state maximum.

No email retention/deletion behavior was executed or runtime-verified in this Explorer Stage 2.

