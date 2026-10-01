# Task 01 — Rekonsiliasi spec domain Donation

> Phase: Techplan decomposition task — post-approval snapshot reconciliation  
> Author: Codex Planner  
> Created / Updated: 2026-10-01  
> Model / reasoning: Invocation configured `gpt-6-luna` / `high`; active runtime not independently exposed  
> Session: Fresh Planner Session for Run `TPD-S2-002-002`; Session ID not exposed  
> Target revision: `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus current durable decisions/projections; reopen live sources  
> Workflow revision: `33b03a3f62cc3aacba6534b8a011465613c64b09`  
> Work Unit: `WU-S2-002`  
> Run: `TPD-S2-002-002`  
> Participant: `P-S2-002-PLD-002-1` (ephemeral for this Run)

## Purpose and scoped outcome

Reconcile Slice 2 Donation domain specifications against Product/MVP, Design, current live authorities, and the Approved TP-015 spine. Produce aligned Donation invariants, threat model, task list, feature acceptance, and only the necessary Campaign eligibility/threshold reference. Classify historical material `KEEP`, `ADAPT`, `REPLACE`, or `DEFER` with authority-based reasons. Keep unresolved Open Items and downstream evidence obligations explicit.

This task does not create or change OpenAPI and does not make owner decisions. Completion does not mean current draft acceptance or `CONTRACT_READY`.

## Parent Techplan

- Authority: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md`
- Version/status: current-effective, Human-approved `Approved` Techplan (2026-10-01).
- TP-015 is the complete authoritative spine for material scope, requirements, rules, decisions, risks, contracts, verification, and Open Items. This child only scopes execution; stop and report any contradiction or missing material parent decision.

## Governing parent references

- Requirements: Q1–Q13, especially Q2/Q12 (shared amount representation) and Q6/Q13 (terminal notice and bounded/recoverable terminalization).
- Rules: R1–R14; preserve domain behavior and lifecycle boundaries, including R12 shared monetary representation, R13 terminal-email lifecycle, and R14 `CONTRACT_READY` boundary.
- Decisions: D1–D19 and rejected alternative D16-alt. Carry settled directions; do not reopen D1, D3–D9, D15–D19 or O7 wording.
- Risks: RISK-1–RISK-10, especially monetary precision, guest-email lifecycle, threshold ordering, and atomic success/funding.
- Verification: all applicable TP-015 §12 obligations for spec reconciliation; runtime, concurrency, email lifecycle, security, and rendered evidence remain with their named downstream owners/phases.
- Open Items: Active O1 (only concrete currency/range/fraction/scale parameters remain unresolved), O2, O3, O4, O5, and independent Human acceptance of the current Donation drafts. O8, O11, O1 shared representation, and O4/O5 directions are resolved and retained in §13 history; do not describe them as open alternatives.

## Scope and live authority/spec anchors

- `docs/spec/5-donation/invariants.md` — evaluate INV-donation-01…15 and historical state machine; adapt only Slice 2 behavior supported by TP-015.
- `docs/spec/5-donation/threat-model.md` — assess submit/status/email/token/abuse threats; do not mark residual risk accepted.
- `docs/spec/5-donation/tasks.md` and `docs/spec/5-donation/features/` — scope to guest submit, truthful simulator state, status/recovery, email obligation, and necessary Campaign boundary.
- `docs/spec/5-donation/features/01-submit-donation-settlement.md` and `02-donation-status-check.md` — historical evidence for submission/settlement and status behavior; reconcile rather than inherit conflicts.
- `docs/spec/5-donation/features/03-public-donor-list.md` through `06-guest-email-reveal.md` — classify outside-slice behavior as `DEFER` unless TP-015 authorizes it.
- `docs/spec/4-campaign/invariants.md` `INV-campaign-13` and `docs/spec/4-campaign/features/09-closure.md` — reference or make a narrowly scoped Campaign-owner change only for approved eligibility/threshold semantics; do not reconcile Slice 3 closure lifecycle.
- Current authorities: `docs/product/README.md`, `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md` §§5–6, `docs/ui-ux/README.md`, `docs/ui-ux/patterns.md` §§7 and 14–15, `docs/spec/README.md`, `AGENTS.md` §§1–3, and TP-015 §§3–13.
- Money: `docs/project/kencleng-monetary-data-standard.md` plus Product Slice 2 input rule. Apply major-unit decimal string plus explicit currency code for API/wire values and exact decimal calculation/persistence. Slice 2 input remains whole-IDR Rupiah, minimum Rp5.000, Rp1 increments. Do not infer range, other currencies, per-currency fraction rules, universal precision/scale, or migration detail from Campaign `NUMERIC(19,2)`.

## Execution detail and boundaries

- Preserve whole-IDR input, minimum/increment, exact-decimal/no-float behavior, four displayed method labels with only QRIS active as sandbox, persisted backend-owned `pending`/`success`/`failed`, request idempotency, guest privacy, D1 threshold/eligibility ordering, and the Slice 2 boundary exactly as the spine states.
- Carry O11/D19: a verified opted-in address remains eligible until the required terminal notice is fulfilled; Donation Delivery must provide bounded, recoverable terminalization. Do not choose a numeric bound, architecture, timeout-as-`failed` meaning, retention alternative, or accept residual risk. Security/PII controls and lifecycle/deletion-race evidence remain downstream under Active O3.
- Carry the settled O4/O5 contract direction when writing domain acceptance: fragment URL with frontend handoff and URL cleanup; one-way HMAC verifier; hard 24-hour lifetime and status-only projection; absent Donation, missing/wrong credential, and expiry share a uniform public `404` with identical Problem Details body, headers, and cache behavior, including `Cache-Control: private, no-store`. Do not claim implementation controls or empirical parity as established; those remain under Active O4/O5.
- Keep the exact O7 terminal label/disclosure wording from Q11/R11/D16. Wording is settled; rendered acceptance and delivery controls remain downstream.
- O8 is resolved only for replacing historical submit/status operations in this repository/current Slice-2 scope because the Human/API owner confirmed they were never externally distributed. Do not generalize beyond that bounded evidence.
- Retain Active O1 only for the concrete parameters TP-015 leaves unresolved; do not restate the approved shared representation itself as gated. O2 timing/mechanism and O3 controls/lifecycle evidence remain scoped to their owners and only block dependent detail.
- Respect Tier-0 fences and the existing encryption/HMAC and safe-logging patterns. No edits to runtime code, tests, Product/Design authority, Techplan, OpenAPI, or orchestration projections.

## Hard dependency

No hard dependency task. Task 01 can execute from TP-015 and current live authority/spec sources. The Human acceptance of current Donation spec drafts is a separate parallel work item and remains required; this task does not claim that acceptance has occurred.

## Verification before handoff

- Trace the reconciled specifications to TP-015 Q1–Q13 and R1–R14 without expanding Slice 2; apply `docs/spec/README.md` authority precedence.
- Check exact amount representation direction and Slice 2 whole-IDR constraints while preserving all unresolved concrete parameters.
- Check D1/R7 Campaign threshold and eligibility semantics, atomic success/funding requirements, and stable close reason without choosing locking/isolation or importing Slice 3.
- Check O11/D19 terminal-notice eligibility and bounded/recoverable terminalization direction; retain unresolved mechanics under O3.
- Check settled O4/O5 behavior is not reopened, while threat/control/evidence obligations and residual-risk ownership remain explicit.
- Do not claim Human acceptance, runtime/security evidence, or `CONTRACT_READY` from draft edits. Record any material authority gap and stop affected detail rather than inventing a decision.

## NOT in this task

No OpenAPI/API artifacts, Product/MVP or Design authority, TP-015, runtime code, tests, migrations, generated artifacts, or orchestration projections. No Tier-0 changes. No decisions on O1 unresolved parameters, O2/O3/O4/O5 implementation controls/evidence, no residual-risk acceptance, and no `CONTRACT_READY` claim.
