# Review Findings — RV-S2-006-003

> Phase: Code Review  
> Author: P-S2-006-RV-003-1 (KC-REVIEWER)  
> Created: 2026-10-02  
> Updated: 2026-10-02  
> Model: Invocation-configured `gpt-6-luna`  
> Reasoning: high  
> Session: fresh; Session ID not exposed  
> Target revision: `7fd8b473b239b20bda3990ab29c51440d321a796` plus exact six-file Run delta  
> Workflow revision: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`  
> Work Unit / Run: `WU-S2-006` / `RV-S2-006-003`

## Scope reviewed

Reviewed the six live sources against their `BLD-S2-006-002/baseline/` snapshots and the approved `TP-S2-006-004/techplan.md`. All six live SHA-256 hashes match the current hashes in the Run Invocation. The review is limited to this source amendment; API and generated counterparts, producer/consumer implementation, and runtime behavior remain later stages.

## 1. Safety

No findings. These are specification changes, not runtime implementation. The capacity, atomic admission, exact-once settlement, and concurrency proof obligations remain explicit in the reviewed sources and approved Techplan; no locking mechanism or runtime evidence is inferred here.

## 2. Quality

**Q-01 — Non-admin close-reason checklist is stale**  
**Location:** `docs/spec/4-campaign/features/09-closure.md:149-150`  
**Problem:** The amendment adds `funding_capacity_reached`, but the checklist still says `closed_by` is null for “the other two reasons.” There are now three non-admin reasons: `max_amount_reached`, `deadline_reached`, and `funding_capacity_reached`.  
**Impact:** The checklist understates the cases whose `closed_by` behavior must be verified and is inconsistent with the amended close-reason set.  
**Resolution:** Change the wording to cover all non-admin close reasons (or enumerate all three) and include the capacity reason in the applicable verification.  
**Blocking:** No; small completeness correction.

## 3. Stack-Specific Best Practices

Relevant triggers matched: PostgreSQL financial invariant enforcement and exact monetary representation. Reviewed `../harscode-workspace/best-practices/postgresql/financial-invariant-enforcement.md` and `../harscode-workspace/best-practices/go/decimal-and-money.md`. No stack-specific finding: the reviewed text preserves exact-decimal/no-float requirements, leaves transaction/locking choices to delivery, and requires downstream concurrent admission and settlement evidence. This Review did not run database or application checks.

## 4. Consistency

**C-01 — Capacity closure misses the “no valid amount fits” boundary**  
**Location:** `docs/spec/4-campaign/invariants.md:230-245`; `docs/spec/4-campaign/features/09-closure.md` “Finite funding-capacity trigger”; `docs/spec/5-donation/invariants.md:35-39, 71-75`; `docs/spec/5-donation/features/01-submit-donation-settlement.md:29-32, 48-53`  
**Authority/evidence:** `docs/product/mvp-scope.md:140-147` and `docs/product/mvp-delivery-slices.md:198-205` say fundraising closes when **no additional valid Donation amount can fit**. The approved Techplan also states this at R3 (`TP-S2-006-004/techplan.md:60`). The active minimum valid Donation is Rp5.000 (`docs/spec/5-donation/invariants.md`, INV-donation-01; Product scope §5).  
**Problem:** The amended sources trigger capacity closure only when an accepted admission “exhausts” the remaining capacity. That expresses an exact-zero remainder. If an admission leaves Rp1–Rp4.999 of capacity, no valid Donation can fit, but none of the new source wording requires the Campaign to close.  
**Impact:** Campaign can remain open while every valid Donation is impossible, contrary to the approved Product rule and the Techplan execution contract. The close reason and public lifecycle would not reflect the actual fundraising state.  
**Resolution:** State the trigger as closing once remaining capacity is below the minimum valid whole-Rupiah Donation (Rp5.000), including exact exhaustion, after accounting for settled Funding and accepted-pending reservations. Preserve the rule that a later pending failure does not reopen a capacity-closed Campaign. Add boundary evidence for residual capacity Rp4.999 versus Rp5.000 and exact exhaustion; a request too large for the remaining capacity must still be rejected without disclosing internal capacity when a smaller valid amount could fit. Apply the correction consistently to the Campaign and Donation invariant/feature wording and verification rows.  
**Blocking:** Yes; source reconciliation does not meet approved Product/Techplan behavior.

## Verification executed during Review

- Reconstructed each of the six source deltas with `diff -u` against its exact BLD baseline snapshot; verified the six current SHA-256 values against the Invocation. All six matched.
- Read the canonical four-pass prompt, review guidelines/checklist, orchestrated-run overlay, reviewer profile, approved Techplan, Product scope/sequencing, monetary standard, applicable root instructions, and the targeted Harscode best-practice files.
- No tests, validators, generators, services, browser, database, or runtime checks were run; the findings are static source/authority mismatches.

## Verdict

**Request changes.** C-01 is blocking. Q-01 is non-blocking and should be corrected with the same source patch.

## Phase handoff

- **Completed:** Four-pass review and verdict.
- **Artifacts:** `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-003/review-findings.md`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-003/patch-plan.md`.
- **Human decision:** None needed to resolve the findings; Product behavior and the minimum Donation amount are already approved. Owning Campaign/Donation spec acceptance remains a later gate.
- **Open / deferred:** Blocking C-01 and non-blocking Q-01. API reconciliation, generated/bundle/consumer correspondence, owning-source acceptance, and implementation/runtime proof remain pending under the approved Techplan.
- **Recommended next step:** Build/Patch in a fresh Run/Participant/Session using the specific patch plan. After the correction, obtain the applicable independent review/source-owner acceptance before dependent API authoring. Do not proceed to Testing or API Build on the unresolved source mismatch.
- **Session transition:** Patch work is a new Build/Patch Run with a fresh Participant Session, re-grounded on this patch plan.
- **Context pointers:** `TP-S2-006-004/techplan.md` R3/Q3; the six files listed in the Invocation, particularly the locations cited in C-01 and Q-01; Product scope §5; Product delivery slices §5.
