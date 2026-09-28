# OIR Resolution Brief — O2 Simulator and O3 Pending Email Retention

> Phase/Stage: Explorer Open-Item Resolution / Stage 3 (complete)  
> Author: `P-S2-002-OIR-003-1` (`KC-EXPLORER`)  
> Created: 2026-09-28  
> Updated: 2026-09-28 — Stage 3 synthesis after Human confirmed Stage 2 evidence  
> Work Unit / Run: `WU-S2-002` / `OIR-S2-002-003`  
> Model / reasoning: invocation dispatch metadata `gpt-6-luna` / `high`; active runtime model is not independently exposed  
> Session: FRESH per invocation; Session ID not exposed  
> Target / workflow revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535` / `7a4dbf2c065bd8fd02c86c24073d7309046bff30`

This Run establishes that current evidence does not support a bounded Donation simulator terminal deadline or a numeric O3 pending-retention maximum. It preserves the settled Product and O3 decisions, compares the available owner routes, and identifies the smallest follow-up. It does not choose a duration, amend an authority artifact, accept residual privacy risk, or establish runtime behavior.

## O2 — Simulator terminal timing and demo-failure control

- **Outcome:** `NEEDS_FURTHER_EVIDENCE`
- **Question:** Does current Slice 2 evidence support a bounded maximum time from persisted `pending` to terminal `success`/`failed`, and what delivery behavior controls demo-only failure?
- **Settled constraints:** Backend simulator owns the persisted outcome; failure is only from a clearly labeled demo scenario controlled by backend config/fixture, never donor choice or browser request. Pending wording remains “Menunggu hasil simulasi,” without a time estimate, payment instruction, real-world timing promise, or SLA. `failed` permits an explicit new donation; `pending` does not auto-resubmit. See `docs/product/mvp-delivery-slices.md` §5, `docs/product/mvp-scope.md` Stage B, `TP-S2-002-007` §13 O2, and prior `OIR-S2-002-001` O2.
- **Evidence:** Current backend has no Donation domain, route, migration, simulator worker, or settlement runtime. Campaign detail currently reports `donation_flow_not_available`. The backend tech-stack document's generic in-process scheduler posture specifies no Donation completion bound or recovery behavior. Historical Donation material describes `2–5s` and `5%`, but current Product/Techplan explicitly do not approve those values; they are not a maximum deadline and are not current runtime evidence. See `evidence/stage-2-o2-simulator.md` and `evidence/stage-2-input-provenance.md`.
- **Options and trade-offs:**
  1. **Donation delivery owner defines an internal bounded terminal policy** (without a user-facing ETA/SLA), including the maximum interval's starting point and behavior for missing, retried, restarted, or failed simulator work. This could provide a basis to bound pending email retention while preserving terminal notification intent. It is not presently demonstrated to be feasible or implemented; the time value and failure/recovery semantics still need current delivery evidence and owner approval.
  2. **Security/PII owner sets a separate pending-state email cap.** This can bound PII retention independently of simulator availability. If the cap expires before a terminal state, the handling of the verified address and any later notification must be made explicit; the evidence does not establish that policy or a suitable duration.
  3. **Use both controls.** Independent limits can bound simulator work and PII retention, but the relationship between them must be explicit. A shorter PII cap can expire before terminal state and prevent the approved email purpose from being fulfilled; no numeric values or expiry behavior are currently approved.
  4. **Leave both unbounded/deferred.** This avoids an unsupported numeric choice but leaves O2 without a delivery bound and O3 pending retention unresolved; indefinite retention risk has not been accepted.
- **Recommendation:** Do not select a numeric deadline from current evidence and do not revive historical values. Route first to Donation delivery planning to establish whether a testable internal maximum terminal interval and bounded recovery behavior can be specified without a public timing promise. If that cannot provide a finite verified-email retention bound, route to Security/PII for an independent cap and explicit handling when it expires before terminal status. This is a recommended sequence for obtaining decision evidence, not a Human/owner decision.
- **Human/owner decision:** None made in this Run. Anhar confirmed the Stage 2 evidence on 2026-09-28; that was a workflow gate only, not a timing, delivery, retention, Product, or risk decision. The evidence is insufficient to request a numeric value without pressuring an unsupported decision.
- **Missing evidence / smallest next route:** A current Donation delivery proposal must specify the bounded interval's start/end conditions and terminal/recovery behavior under non-enqueued work, process restart, execution error/timeout, and replay; it must also show that the demo-failure control is unavailable to donor/UI inputs and is clearly identified as simulation. Donation delivery owner / Planner route in a later Run. Runtime proof remains later Build/Testing work.

## O3 — Maximum verified-email retention while `pending`

- **Outcome:** `DEFERRED`
- **Question:** What is the maximum time verified guest email may be retained while Donation remains `pending`?
- **Settled decision preserved:** Human decision in `OIR-S2-002-002` (2026-09-28) retains ownership verification; sets a 24-hour verification window from email capture and 24-hour delivery retry window from terminal state; deletes unverified email at verification expiry and verified email after successful terminal delivery or retry-window expiry. It explicitly deferred maximum retention while state remains `pending` until O2 timing evidence or an explicit Security/PII cap exists. This Run does not reopen those settled windows. No residual privacy risk was accepted.
- **Evidence and dependency:** O2 has no current Donation runtime or bounded terminal evidence, so a finite pending period cannot be inferred. The already-approved 24-hour terminal retry window begins at terminal state and does not cap the preceding pending period. The verification window and 24-hour guest status URL lifetime are separate controls. See `evidence/stage-2-o3-pending-retention.md` and `OIR-S2-002-002/resolution-brief.md` O3.
- **Options and trade-offs:**
  1. Derive the maximum pending retention from a future, approved, testable O2 terminal bound plus the already-decided terminal retry period. This aligns retention with the simulator's bounded purpose, but is unavailable until the O2 bound and its time anchors are explicit and evidenced.
  2. Set an independent Security/PII cap measured from an explicitly named event. This bounds PII even if Donation never terminates, but could expire before terminal state; the owner must define what is deleted and whether/when a terminal notice can still be sent without retaining the address.
  3. Apply both an O2 bound and a Security/PII cap. This gives independent controls but needs an explicit precedence/expiry rule and may truncate email notification eligibility.
  4. Continue the existing deferral. This avoids inventing policy but leaves a verified address potentially retained without a finite `pending` maximum; this Run cannot treat that residual risk as accepted.
- **Recommendation:** Keep the current deferral. First obtain an owner-reviewed O2 bound if a testable bounded terminal policy is feasible; then Security/PII can determine whether that duration plus its existing terminal retry window is an adequate retention maximum or whether an independent cap is also required. If no O2 bound is available, Security/PII must set an independent cap and resolve expiry-before-terminal handling before the contract can close. No duration is supported by this evidence.
- **Human/owner decision:** None made in this Run. The current Human decision explicitly left this value open. Anhar is the named current Slice 2 Donation delivery and Security/PII owner, but owner attribution is not a decision.
- **Missing evidence / smallest next route:** First, O2 delivery evidence as listed above. If O2 cannot supply a finite maximum, request an explicit Security/PII decision naming the cap duration, its start event, deletion point, and terminal-notification handling if Donation is still pending at expiry. If that handling changes Product notification meaning, route that specific question to Product Authority instead of inferring it.

## Scope, provenance, and verification

- Assignment pins were re-read and matched; hashes and current-tree caveat are recorded in `evidence/stage-2-input-provenance.md`.
- Human confirmed Stage 2 on 2026-09-28. No substantive owner decision, Product change, Design decision, or residual-risk acceptance was made in Stage 3.
- Stage 2 inspected current Product/MVP authority, the approved Techplan, prior OIR decisions, Donation spec/contract, backend routes/domain presence, backend architecture, and routed Go lifecycle guidance where the missing worker path made it relevant.
- No tests, runtime checks, deployment/topology checks, implementation behavior, or independent security review were performed.
- Pre-existing working-tree changes outside this Run were observed and left unchanged; see provenance evidence. No other Run artifact, orchestration state, spec, contract, Product/Design artifact, code, or test was modified.

