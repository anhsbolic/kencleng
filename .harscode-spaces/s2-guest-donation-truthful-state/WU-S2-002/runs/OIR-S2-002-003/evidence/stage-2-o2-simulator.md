# Stage 2 Gap Evidence — O2 Simulator Timing and Demo Failure

> Phase/Stage: Exploration / Stage 2  
> Author: `P-S2-002-OIR-003-1` (`KC-EXPLORER`)  
> Created: 2026-09-28  
> Model / reasoning: invocation dispatch metadata `gpt-6-luna` / `high`; active runtime model is not independently exposed  
> Session: FRESH per invocation; Session ID not exposed  
> Work Unit / Run: `WU-S2-002` / `OIR-S2-002-003`  
> Target revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535` (`HEAD` at inspection)  
> Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`

## Current state

- Current Product/MVP authority requires persisted backend-simulator-owned `pending` → `success`/`failed` state. Only a clearly labeled demo scenario controlled by backend configuration/fixture may produce a failure; donor choice or browser request must not choose an outcome. Pending copy is “Menunggu hasil simulasi,” with no time estimate, real-payment instruction, real-world timing promise, or SLA. After `failed`, a donor may explicitly begin a new donation; `pending` never triggers automatic resubmission. Sources: `docs/product/mvp-delivery-slices.md` §5; `docs/product/mvp-scope.md` Stage B.
- Current-effective Techplan `TP-S2-002-007` §13 O2 says exact simulator timing and scenario implementation remain undecided; §4 Q4/R4 and Risk-3 preserve the Product constraints and exclude historical timing/probability. OIR `OIR-S2-002-001` O2 records the prior Human state/recovery direction and explicitly leaves timing/mechanism open.
- Historical `docs/spec/5-donation/features/01-submit-donation-settlement.md` describes an internal settlement process delayed `2–5s` with a `5%` failure rate, and enqueuing a settlement job. `docs/spec/5-donation/invariants.md` INV-donation-09 likewise encodes the delayed internal callback. These are historical delivery artifacts; they conflict with current-effective Product/Techplan as policy and do not prove that the current backend implements a bounded terminal outcome.
- Current backend tree has no Donation domain package, Donation migration, settlement worker, or Donation HTTP route. `backend/cmd/server/main.go` registers public Campaign detail/media routes but no Donation routes. `backend/internal/domain/campaign/service.go`'s `toPublicDetail` sets `DonationAction` to `unavailable` / `donation_flow_not_available`. No current path was found that persists Donation state, schedules/settles an outcome, or lets a reviewer inspect backend-owned failure-scenario configuration.
- `docs/project/kencleng-backend-tech-stack.md` §Decided: Backend lists an in-process scheduler as a general low-complexity stack choice for other planned jobs. It defines no Donation queue, retry/timeout policy, terminal deadline, shutdown/recovery contract, or demo-failure control. It is architectural context, not evidence of O2 behavior.
- `docs/ui-ux/README.md` routes product/design authority to the active design system; current O2 copy is already directly specified by Product. No material visual/interaction decision is selected or needed to report the technical gap. The invocation keeps any new user-facing timing promise or material Design expression outside this Run's authority.

## Requirement and gap

**Requirement:** The simulator must own the persisted terminal outcome and the demo-only failure trigger. Public/product messaging must remain without a timing estimate/SLA. Current sources do not require a public deadline. They also do not currently prescribe a maximum internal time to terminal state.

**Gap:** The implementation needed to establish a maximum terminal deadline does not exist in the current backend tree, and no approved timing or job-completion/recovery behavior is present in the current-effective plan. The available evidence therefore cannot support a bounded simulator terminal deadline. Absence of an SLA does not itself prove that a Donation will terminate within a bounded interval. Historical `2–5s`/`5%` detail cannot fill this gap.

This evidence describes the current revision and requirement gap only; it does not prescribe how to implement timing or scenario selection.

## Sniffing lenses

- **Risk:** Without a proven terminal bound or operational recovery rule, a persisted `pending` Donation may remain pending indefinitely after process/job failure or scheduling loss. A donor could interpret that as a stuck/real-payment-like state despite honest copy. An operator-controlled failure trigger that reaches request/UI inputs would violate outcome authority. Any later implementation also touches money-state transitions; balance/transaction implementation is Tier-0 fenced and is not inspected or changed here.
- **Edge cases:** Worker/process restart after persisting `pending`; job not enqueued or lost after submission; simulator work errors/timeouts; duplicate/replayed simulator job; scenario flag absent or invalid; config changes while a job is already pending; test/demo settings accidentally enabled in another environment. Current code provides no Donation path with which to establish behavior for these cases.
- **Miscontext:** Historical 2–5 seconds and 5% may look like a usable terminal bound and deterministic policy, but they are explicitly not approved and do not encode a maximum duration (probability is not a deadline). A generic backend scheduler declaration does not imply a Donation job exists or has a completion bound.
- **Misleading signals:** `api/openapi/donation.yaml` contains historical Donation endpoints and schemas; historical Feature 01 describes an async settlement job; neither is wired to a backend Donation route or runtime. Campaign detail's Donation action type exists but is concretely returned as unavailable.
- **Inconsistency:** Historical Feature 01, Donation invariants, and task notes express exact delay/probability and old notification semantics. Current Product and Approved Techplan forbid inheriting those values. Current backend exposes donation unavailable. These differences must be reconciled in their owning spec/contract/implementation routes later; Stage 2 does not choose a direction.

## Code and authority anchors

- `docs/product/mvp-delivery-slices.md` §5, simulator and pending-state bullets — binding current Slice 2 result authority, demo-failure source, and no-estimate/no-payment wording.
- `docs/product/mvp-scope.md` Stage B — no real-world timing promise or SLA; explicit failure recovery and no pending auto-resubmit.
- `runs/TP-S2-002-007/techplan.md` §13 O2; §4 Q4/R4; Risk-3 — current open delivery detail and historical-value exclusion.
- `backend/cmd/server/main.go` route registration around lines 159–160 — no Donation route is registered at this revision.
- `backend/internal/domain/campaign/service.go` `toPublicDetail` — current Donation action is explicitly unavailable.
- `api/openapi/donation.yaml` Donation operations — authored historical/current contract source to reconcile later, not evidence of runtime registration.
- `docs/spec/5-donation/features/01-submit-donation-settlement.md` — historical 2–5 second/5% behavior and enqueue step; not current policy or live behavior.
- `docs/project/kencleng-backend-tech-stack.md` Background Jobs / Scheduler row — generic scheduler posture only; no Donation deadline/recovery guarantee.

No tests, runtime jobs, deployment behavior, or scheduling guarantees were verified in this Explorer Stage 2.

