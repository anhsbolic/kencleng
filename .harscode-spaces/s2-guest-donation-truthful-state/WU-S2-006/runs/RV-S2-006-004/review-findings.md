# Review Findings — RV-S2-006-004

> Phase: Code Review — targeted confirmation  
> Author: P-S2-006-RV-004-1 (KC-REVIEWER)  
> Created: 2026-10-02  
> Updated: 2026-10-02  
> Model: Invocation-configured `gpt-6-luna`  
> Reasoning: medium  
> Session: fresh; Session ID not exposed  
> Target revision: `7fd8b473b239b20bda3990ab29c51440d321a796` plus exact four-source patch  
> Workflow revision: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`  
> Work Unit / Run: `WU-S2-006` / `RV-S2-006-004`

## Scope and applicability

This is the narrow confirmation permitted by the canonical Code Review prompt for a patch that implements a specific Review patch plan without broadening the change. It confirms C-01 and Q-01 from `RV-S2-006-003` against its exact patch plan, the approved `TP-S2-006-004` spine, Product scope/delivery wording, the four BLD baseline snapshots, and the live sources. It is not a new full review of the six-file checkpoint or an owning-Human source acceptance.

The baseline-to-live comparison is confined to:

- `docs/spec/4-campaign/invariants.md`
- `docs/spec/4-campaign/features/09-closure.md`
- `docs/spec/5-donation/invariants.md`
- `docs/spec/5-donation/features/01-submit-donation-settlement.md`

All four live SHA-256 hashes match the Invocation. The remaining two sources from the prior six-file checkpoint were declared unchanged from the prior Review snapshots and are outside this patch. The working tree contains unrelated changes, which were excluded from this confirmation.

## 1. Safety

No new findings in the confirmed patch. The text preserves reservation accounting, full settlement for accepted-pending Donations after close, stable winning close reason, and no reopening after reservation release. It does not select or imply a transaction/locking mechanism or claim runtime proof.

## 2. Quality

**Q-01 — Resolved.** `docs/spec/4-campaign/features/09-closure.md` now says `closed_by` is null for all non-admin close reasons and explicitly lists `max_amount_reached`, `deadline_reached`, and `funding_capacity_reached`. This covers the newly added reason and repairs the stale “other two reasons” checklist wording.

No additional material quality finding in the narrow patch.

## 3. Stack-Specific Best Practices

No new stack-specific trigger was introduced by this specification-only correction. The prior review's targeted PostgreSQL financial-invariant and Go exact-money guidance remains applicable to the unchanged requirement boundary: `../harscode-workspace/best-practices/postgresql/financial-invariant-enforcement.md` and `../harscode-workspace/best-practices/go/decimal-and-money.md`. No implementation mechanism or runtime behavior is reviewed or inferred here; their runtime/concurrency evidence remains deferred to delivery/Testing.

## 4. Consistency

**C-01 — Resolved.** Campaign INV-campaign-13, Campaign feature `09-closure.md`, Donation INV-donation-02, and Donation feature `01-submit-donation-settlement.md` consistently close for `funding_capacity_reached` when remaining capacity is less than the minimum valid Donation of Rp5.000, including exact exhaustion, after settled Funding and accepted-pending reservations are accounted for. They retain rejection of a full amount that does not fit and preserve the distinction from the overshootable `max_amount` threshold. The Campaign checklist covers Rp4.999 (close), Rp5.000 (remain open), exact exhaustion, and non-disclosure when a smaller valid amount could fit.

This matches Product's “When no additional valid Donation amount can fit” rule in `docs/product/mvp-scope.md` (Stage D) and `docs/product/mvp-delivery-slices.md` (Slice 2), read with the established Rp5.000 minimum in Product scope and Donation INV-donation-01. The boundary therefore expresses the approved Product behavior rather than introducing a new policy. The accepted-pending failure/no-reopen rule and Slice-2 versus Slice-3 boundary remain unchanged.

No new material inconsistency found in the confirmed patch.

## Verification executed during Review

- Reconstructed baseline-to-live diffs for each of the four sources against `BLD-S2-006-003/baseline/`; confirmed the four SHA-256 values against the Invocation. All matched.
- Compared the corrected trigger and checklist with the exact prior C-01/Q-01 findings and patch plan, approved Techplan R3/Q3, Product scope/delivery wording, and the Rp5.000 minimum in Donation INV-donation-01.
- No tests, validators, generators, services, browser, database, or runtime checks were run. Runtime reservation, concurrency, settlement, API compatibility, and public-projection behavior remain deferred as assigned.

## Verdict

**Approve** within this narrow technical confirmation scope. C-01 and Q-01 are resolved; no new finding or material patch drift was found. This verdict does not accept the concrete six-file spec checkpoint, complete WU-S2-006, release API/production/runtime gates, or establish runtime proof.

## Phase handoff

- **Completed:** Targeted independent confirmation of C-01/Q-01, exact patch scope, and adjacent consistency.
- **Artifacts:** `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-004/review-findings.md`.
- **Human decision:** Owning Human acceptance of the concrete Campaign/Donation spec checkpoint remains required; technical Review cannot accept source bytes.
- **Open / deferred:** Spec source acceptance; authored API reconciliation and counterpart/generated/consumer correspondence; backend/frontend Techplan refreshes after source convergence; runtime reservation/concurrency and settlement evidence; other gates listed in the approved Techplan.
- **Recommended next step:** Route the concrete six-file spec checkpoint to its owning Human acceptance gate. If accepted, proceed through the approved source/API reconciliation route; do not infer API, production, runtime, DB-application, or delivery approval from this Review.
- **Session transition:** Subsequent acceptance/phase work uses its assigned owner or a fresh Run/Participant/Session as applicable.
- **Context pointers:** `TP-S2-006-004/techplan.md` R3/Q3; `RV-S2-006-003/review-findings.md` and `patch-plan.md`; `BLD-S2-006-003/baseline/` and `source-delta.patch`; the four confirmed source files; Product scope Stage D and Slice-2 delivery wording.
