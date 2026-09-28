# O2 Delivery Proposal — Simulator Terminal Policy and O3 Retention Dependency

> Work Unit / Run: `WU-S2-002` / `TP-S2-002-009`
> Role / specialization: Planner / Bounded O2 simulator delivery feasibility and O3 pending-retention decision preparation
> Participant / profile: `P-S2-002-TP-009-1` / `KC-PLANNER`
> Pinned profile: `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`
> Target revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535`; current working-tree inputs re-read
> Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`
> Status: Decision preparation; no owner decision or numeric duration selected

## Purpose and current authority

This proposal analyzes whether an internal, finite maximum from persisted Donation `pending` to terminal `success`/`failed` can be specified without exposing an ETA, real-world timing promise, or SLA. It also traces that bound's effect on verified guest-email retention while `pending`.

Current Product/MVP requires the backend simulator to own persisted outcomes; failure must come from a clearly labeled demo scenario controlled by backend configuration/fixture, not donor choice or browser input. Pending copy remains “Menunggu hasil simulasi,” with no estimate or real-payment instruction. Pending does not auto-resubmit; after `failed`, the donor may intentionally start a new donation. These requirements are in `docs/product/mvp-delivery-slices.md` §5 and `docs/product/mvp-scope.md` Stage B. This proposal does not change those semantics.

The Approved baseline remains `TP-S2-002-007` §4 Q4/R4/R5, §9, §12, and §13 O2/O3. It leaves exact timing and delivery recovery open. `OIR-S2-002-003` Stage 2 found no Donation runtime or approved bound; its O2 result is `NEEDS_FURTHER_EVIDENCE`. The prior Human O3 decision in `OIR-S2-002-002` remains in force: 24-hour verification window from email capture; 24-hour delivery retry window from terminal state; delete unverified email at verification expiry; delete verified email after successful delivery or retry-window expiry. The maximum retention of verified email while Donation remains `pending` was deferred. None of those settled decisions is reopened here.

## Feasibility finding

A finite **internal policy bound** is technically specifiable without a public timing promise, if Delivery defines a persisted start/deadline and an implementation with durable discovery, restart recovery, timeout handling, and idempotent terminalization. A duration alone, an in-process delay, or a probability of failure is not a bound.

Current evidence does **not** establish that such a bound is implemented or operationally guaranteed: no Donation domain, route, migration, simulator worker, or settlement runtime exists. `backend/cmd/server/main.go` currently registers Campaign detail at the inspected route area, and `backend/internal/domain/campaign/service.go` returns `DonationAction` as `unavailable` / `donation_flow_not_available`. The backend architecture's generic in-process scheduler entry (`docs/project/kencleng-backend-tech-stack.md`, Background Jobs / Scheduler row) supplies no Donation liveness, restart, or completion guarantee. Thus no numeric duration can be supported or recommended from current runtime evidence.

There is also a material Product-semantic question: automatically terminalizing a donation as `failed` when simulator work times out or infrastructure fails may conflict with the current rule that failure is produced only by a clearly labeled demo scenario. If that timeout is to become a donor-visible `failed` state, Product Authority must confirm that the meaning fits the approved scenario rule and any resulting copy. The Planner must not relabel a system error as a demo outcome. If Product does not approve that interpretation, the Delivery owner must propose a different bounded terminal behavior and route its user-visible meaning to Product before a contract is changed.

## Candidate bounded delivery approaches

Both approaches below are viable design directions for later owner selection. Neither is selected or proven by this Run. Both need a durable record of simulator work and a recovery mechanism; a raw in-memory timer does not meet the conditions.

| Approach | Bound and recovery shape | Strengths | Costs / conditions |
|---|---|---|---|
| **A. Persisted work/outbox with worker and deadline reconciler** | In the same durable transaction that accepts a Donation as `pending`, persist an internal simulator work record (or outbox entry) with a fixed `accepted_at`/`terminal_deadline`. A worker claims work; a periodic/startup reconciler discovers unprocessed/expired records and safely re-enqueues or applies the approved deadline behavior. On expiry, the system must reach an owner-approved terminal outcome. | Closes the “Donation committed but enqueue missing” gap if the work record is atomic with acceptance; explicit due-time supports recovery after restart; replay can be idempotent by Donation/work identity. Separates submission latency from simulation. | Requires persistence/schema and worker lifecycle; claim/lease/stale-worker handling; operational recovery/monitoring; exact timeout outcome approval. A queue enqueue by itself is insufficient unless its intent is durable/atomic or recoverable from persisted Donation state. Availability/recovery assumptions must be made explicit before calling the bound guaranteed. |
| **B. Database-backed due-work sweep** | Persist `pending` plus a fixed deadline at acceptance. A scheduler periodically scans due/pending rows, claims a row using established concurrency-safe repository patterns, and executes or terminalizes according to approved policy. Startup scan resumes overdue work; repeated scans recover missed ticks. | Fewer moving parts for current single-instance sandbox posture; database state is the recovery source; no separate broker is required. | Polling cadence and scan cost affect completion after the policy deadline; overlapping workers/restarts require safe claims; an in-process scheduler alone cannot guarantee liveness through arbitrary downtime. The system needs a defined maximum recovery lag/availability assumption and alerting or an explicit operational failure path. |

**Not a bounded approach:** persist `pending`, then rely only on best-effort in-memory enqueue/timer. A process crash or missed enqueue can leave no durable work to rediscover. Retrying enqueue without an idempotent durable work identity also permits duplicate/replayed jobs.

### Minimum timing and recovery contract to make either approach meaningful

The Delivery owner should decide and record the following before selecting any number:

1. **Start event:** set a single immutable clock origin when the accepted Donation and its `pending` state commit durably (prefer persisted acceptance timestamp); create its simulator work/deadline in that same atomic acceptance boundary or demonstrate an equivalent recovery path. Do not start at HTTP request arrival, in-memory enqueue, or worker pickup: each can omit time spent before it.
2. **End event:** stop the interval only when `success` or `failed` is durably committed by the authorized backend simulator transition. A job being enqueued, attempted, acknowledged in memory, or returning from a worker without a state commit is not terminal.
3. **Missing enqueue / lost scheduling:** durable work must be discoverable from persisted state after acceptance. A scanner/reconciler should recover records with no active work, and detect work that is past its deadline. If the chosen transaction cannot atomically persist Donation and work intent, specify and test the compensating recovery mechanism.
4. **Restart:** startup recovery must discover unclaimed, stale-lease, and overdue simulator work from durable state. Restart must not reset the original deadline or create a fresh full interval.
5. **Timeout / execution error:** use a bounded attempt/lease policy and an explicit final action at the deadline. The action cannot be chosen until the Product question above is answered: an infrastructure error is not automatically an approved demo failure. If retry/reconciliation continues, that is not a terminal bound. If it terminalizes, Product and Donation Delivery must specify the truthful state and whether the demo scenario is the only source of `failed`.
6. **Replay / duplicate work:** make terminal transition idempotent per Donation/work identity; repeated or late execution must not overwrite terminal state, trigger another notification, or duplicate funding. Preserve the Techplan atomic success/funding invariant and protected ledger/locking boundary.
7. **Meaning of “maximum”:** name the allowed worker/recovery lag and runtime availability assumptions. If the process/database can remain unavailable indefinitely, no software deadline alone proves an unconditional wall-clock terminal guarantee. State whether the approved value is a policy deadline under stated availability or an operational guarantee backed by deployment/recovery evidence. Do not expose that internal value in donor copy absent Product approval.

### Backend-controlled demo failure

The demo failure scenario must be selected only by backend-controlled configuration or fixture, with a clear simulation label. It must not be derived from a donor-supplied field, request parameter/header, frontend state, or public settlement endpoint. Configuration scope/defaults must ensure the scenario cannot silently leak into an environment where it is not intended; invalid or absent configuration must have a defined safe behavior. The persisted job/result should retain enough non-PII provenance for operators and tests to distinguish the selected demo scenario from infrastructure failure.

This describes a boundary, not a mechanism choice. Configuration format, environment gating, scenario determinism, and auditability remain for Donation Delivery/Design review. Historical `2–5s` and `5%` values in the old Donation feature/invariant are not current authority, are not a maximum, and are not carried forward.

## O3 retention consequences

Let **B** be a future owner-approved maximum from durable `pending` acceptance to durable terminal state, **R = 24 hours** be the already-approved post-terminal delivery retry window, and **C** be any separately approved Security/PII cap measured from an event the owner names. No value for B or C is set here.

| Route | Verified email while Donation is pending | If expiry occurs before terminal state | Later terminal notification eligibility |
|---|---|---|---|
| **A or B with approved terminal bound B** | If Security/PII accepts deriving retention from O2, keep the verified address only for the approved status-notification purpose until terminal delivery/retry rules apply. A bound from initial persisted `pending` gives a finite basis; a conservative maximum from that same origin is `B + R` when the terminal result may occur at B and retries may continue for all of R. | Under a genuine O2 bound, `pending` should no longer persist beyond B if liveness assumptions hold. A missed deadline must follow the approved terminal/error route; it cannot silently leave email retained indefinitely. | If terminal occurs by B, existing rules apply: one terminal-only status message, and delete after success or by terminal+R. If terminal cannot be reached by B, define its terminal state first; notification is eligible only if email remains verified/retained under the approved rule. |
| **Independent Security/PII cap C** | Retain only until C from its explicitly chosen start event, irrespective of simulator progress. This independently bounds PII if work is stuck. | Delete the verified address at C while Donation may still be `pending`; the Donation may continue according to its own state policy. | Once email is deleted, a later terminal state cannot use that address for the already-approved notification. Treat it as ineligible unless the donor independently re-supplies/re-verifies under a Product- and Security-approved flow. Do not retain, recreate, or silently re-verify it to recover the notice. |
| **Both B and C** | Apply whichever approved limit expires first, with explicit precedence and deletion semantics. This offers independent bounds. | If C < B, apply the same deletion consequence as the independent cap; if B ≤ C, process the approved terminal outcome and then the existing R window. Equality/race handling must be deterministic. | A deleted address is not eligible later. If still retained when terminal occurs, existing verification and terminal retry rules apply, then delete by the earlier valid limit. |
| **Continue deferral** | No finite maximum is established for a verified address while status stays pending. | No approved expiry behavior exists. This leaves the identified retention risk open; it is not accepted by this Run. | Current terminal send/deletion rules still apply if terminal is reached, but they do not solve indefinite pending retention. |

`B + R` is a bound only for the chosen retention derivation; it is not an O2 or public promise. The existing 24-hour verification window governs **unverified** email, and the 24-hour status URL lifetime governs a separate credential. Neither bounds verified-email retention while `pending`.

## Decisions to prepare for owners

**Donation Delivery/domain owner:**

- Is the internal bound a policy deadline under named availability assumptions, or an operational maximum with a measured/recovered liveness guarantee?
- Which approach (durable outbox/worker plus reconciler, or database due-work sweep) fits the actual deployment and recovery ownership?
- What event starts the interval, what durable event ends it, how is maximum recovery lag accounted for, and what happens on timeout/system error without violating Product semantics?
- What backend-only fixture/config boundary chooses the demo failure scenario and prevents donor/UI control?
- What measured or structural evidence can justify B? Relevant evidence may include scheduler/worker recovery design, observed processing and recovery distributions under representative load/fault/restart conditions, worst-case queue/scan lag, and bounded dependency timeout/retry budgets. Mean latency or a probability such as 5% failure is insufficient to support a maximum. A measured percentile is evidence for a policy choice, not proof of a hard maximum.

**Security/PII owner:**

- If B is supportable, is `B + 24h` an acceptable maximum retention horizon from persisted acceptance, or is a separate cap C required?
- If an independent C is needed, what is its duration and start event; what exact data is deleted at expiry; how are deletion-vs-terminal/retry races serialized; and is the notification permanently ineligible after deletion?
- If there is no O2 bound, choose an independent cap or leave O3 explicitly open with risk unaccepted. Do not infer C from the 24-hour verification/retry windows.

These questions prepare decisions; this Run records **no** Delivery, Security/PII, Product, Design, or residual-risk decision. Any change to what donor-visible `failed` means, or to notice availability after email expiry, must be routed to Product Authority. This proposal adds no ETA/SLA and does not authorize final specification or contract text.

## Anchors and later evidence obligations

| Anchor | Relevance |
|---|---|
| `docs/product/mvp-delivery-slices.md` §5; `docs/product/mvp-scope.md` Stage B | Binding state, failure-control, copy, retry, and no-public-promise semantics. |
| `TP-S2-002-007/techplan.md` §4 Q4/R4/R5, §7 Risk-3/Risk-4, §9, §12, §13 O2/O3 | Current Approved planning constraints; retain existing success/funding and guest-email verification controls. |
| `OIR-S2-002-003/resolution-brief.md` O2/O3 and `evidence/stage-2-o2-simulator.md`, `evidence/stage-2-o3-pending-retention.md` | Exact current gap, prior options, and limits of available evidence. |
| `OIR-S2-002-002/resolution-brief.md` O3 | Preserved Human decision on 24-hour verification/retry and deletion; explicit pending-retention deferral. |
| `backend/cmd/server/main.go` router; `backend/internal/domain/campaign/service.go` `toPublicDetail` | Current absence of Donation route and explicit unavailable action. Reopen live code before implementation. |
| `docs/project/kencleng-backend-tech-stack.md` Background Jobs / Scheduler row | General in-process scheduler context only; not a Donation completion guarantee. |
| `backend/internal/domain/donation/ledger.go` and transaction/locking implementation | Tier-0 read-only boundary; preserve atomic success/funding invariant, route any protected implementation need to Human. |

Before any later Build/Testing completion claim, evidence should cover: accepted Donation and durable work creation boundary; recovery of absent enqueue and restart/stale work without resetting deadline; bounded execution timeout/error handling; repeated/replayed/late job idempotency; terminal transition and exactly-once notification; backend configuration/fixture control unavailable to donor/UI; no misleading timing promise; verified-email retention/deletion at B/C, verification, terminal+24h retry, and expiry-before-terminal boundaries; and deletion/terminal/retry races. Use appropriate concurrency/race evidence for shared state under `backend/AGENTS.md`; do not modify Tier-0 ledger/locking logic without the required Human authorization. This Run executed no code or tests and creates no test requirements beyond the current Approved Techplan's rules.

## Planner conclusion

**Feasibility:** A finite internal bound has plausible durable implementation shapes, but neither the bound nor its terminal failure semantics can be established from current implementation evidence. There is a Product decision dependency if infrastructure timeout is expected to create donor-visible `failed`.
**O3:** Keep the pending-retention maximum deferred until B and/or C, including the expiry-before-terminal rule, is explicitly decided. Preserve the already-settled 24-hour windows.
**Next route:** Orchestrator routes this proposal to the named Donation Delivery/domain owner for a bounded feasibility/semantics decision, then to Security/PII for retention derivation or independent cap. Route any changed public meaning to Product Authority. Do not mark `CONTRACT_READY` or revise the Approved Techplan in this Run.
