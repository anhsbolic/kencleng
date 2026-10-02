# Code Review Findings — RV-S2-006-006

Phase: Code Review  
Work Unit / Run: `WU-S2-006` / `RV-S2-006-006`  
Author: `P-S2-006-RV-006-1` — Reviewer / `KC-REVIEWER`  
Created: 2026-10-02  
Model / Reasoning: Invocation configures `gpt-6-luna` / `high`; active runtime model and effort not independently exposed  
Session: fresh Run context; session identifier not exposed  
Target revision: `d186feb2ce3d447d31bccd038dc3b9703f1d2369` plus the reviewed working-tree source diff  
Workflow revision: `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`

## Scope and source identity

Reviewed the seven assigned authored files: Campaign and Donation invariants/features plus `api/openapi/campaign.yaml` and `api/openapi/donation.yaml`. The seven current SHA-256 values match the Run Invocation. `HEAD` matches its declared baseline. Additional modified orchestration projections and the development tracker were not part of this source review.

## 1. Safety

No findings.

The authored contract distinguishes individual cap validation from eligible capacity no-fit, keeps the latter generic on `amount`, retains closed/ineligible `409`, and does not expose remaining capacity or the capacity close reason. `funding_capacity_reached` is represented on the Campaign contract; the closed `PublicCampaignDetail` allowlist does not expose that internal reason. No concurrency mechanism or runtime correctness is established by these source changes. The approved Techplan already assigns capacity reservation, close ordering, and settlement concurrency evidence to downstream Testing; no new Techplan drift was found.

## 2. Quality

No findings.

The cap's range, required object members, currency, and create/PATCH omission behavior are readable across spec and OpenAPI. Capacity no-fit wording preserves the distinction between server predicates and explains the public response without introducing a capacity-specific error shape.

## 3. Stack-Specific Best Practices

**Finding RV-006-01 — Required response field needs strict-client compatibility review (non-blocking to this Review).**

- **Location:** `api/openapi/campaign.yaml#/components/schemas/Campaign` and `#/components/schemas/PublicCampaignDetail`; the latter is closed with `additionalProperties: false` and now requires `max_donation_amount`.
- **Problem:** Adding a required response member can break older strict decoders that reject unknown fields or validate against the prior closed public schema. This is a compatibility risk, not a mismatch with the approved requirement: R8 requires the field in Campaign and public-detail responses.
- **Why it matters:** A server rollout can cause older clients to reject otherwise valid responses. The approved Techplan §6 and OI-4 explicitly require strict-client/known-consumer compatibility review before source acceptance.
- **Suggested resolution:** Before accepting these source bytes, review the known consumers and strict-decoder behavior, then record the compatibility/rollout decision. After the source gate converges, reconcile the generated bundle, frontend types, fixtures, and affected consumers in the order required by the Techplan and `api/README.md`.
- **Blocking:** Non-blocking to the source-shape Review verdict; compatibility review remains an owner/source-acceptance gate. No source edit is identified by this finding.

The routed Harscode guidance was applied: `restapi/openapi-spec-first-drift.md` calls for spec/generated/error-shape correspondence; `restapi/idempotency-and-versioning.md` calls for classifying response-shape compatibility and preserving retry semantics; `restapi/anti-enumeration.md` supports generic responses without leaking sensitive distinctions; and `postgresql/financial-invariant-enforcement.md` routes concurrent financial correctness to transaction-level enforcement and dedicated evidence. No source-level defect was found under these checks. Generated/client correspondence and runtime financial evidence remain deferred as the approved sequencing requires; the Build report's OpenAPI validation is not counted as Review verification.

## 4. Consistency

No findings.

- Campaign and Donation specs express the same accepted DEC-API-01/02 behavior as the split OpenAPI sources: closed `{amount, currency_code}` shape, IDR range, create default, PATCH preserve/freeze, required response shape, generic eligible capacity no-fit `422` on `amount`, preserved `409` and retry semantics, and non-disclosure.
- The shared `ValidationError` in `api/openapi/common.yaml` is referenced for Donation `422`; the common component itself was unchanged. Campaign owns the new close reason, while the public detail remains a separate closed projection without it.
- The authored split files are the correct API source per `api/README.md`. No path was added, so `index.yaml` does not need a path-list change. Bundle, generated types, fixtures, and consumers were not edited; this follows the approved Techplan/Invocation gate sequencing and is not treated as completed counterpart evidence.
- Product/MVP and the approved Techplan support the cap/capacity semantics. The Donation change preserves the separate `max_amount` threshold and accepted-pending settlement behavior.

## Verification executed during Review

No tests, OpenAPI validation, bundle generation, or runtime reproduction were run. Read-only `git diff`, source-hash, and authority inspections confirmed the assigned scope and the seven source identities. The Build report's validation results were read as orientation only, not as Review-run verification.

## Verdict

**Approve with minor comments.** The assigned source changes implement the approved contract. RV-006-01 is a compatibility gate for the named owners before source acceptance; it does not identify a source patch for Build. This verdict does not accept source bytes, establish runtime behavior, or complete WU-S2-006.

## Phase handoff

- **Outcome:** `COMPLETED` — four-pass independent Review completed against the assigned current source diff.
- **Result refs:** `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-006/review-findings.md`; no `patch-plan.md` because no source changes are requested.
- **Findings:** RV-006-01 records the strict-client compatibility review needed before source acceptance; non-blocking to this Review verdict.
- **Decision requests:** None within this Review Run. Orchestrator to route the changed Campaign/Donation spec and API bytes to their named owners for explicit acceptance after the compatibility gate is addressed.
- **Blockers:** None for this Run. Source acceptance and downstream counterpart work remain gated by their owning review/acceptance sequence.
- **Open / unverified:** Known strict-client/consumer compatibility; owner acceptance of the seven changed source identities; generated bundle/type/fixture/consumer correspondence; runtime cap/no-fit, closed-state, concurrency, settlement, and closure evidence.
- **Recommended continuation:** Orchestrator routes the exact reviewed source revisions and RV-006-01 to Campaign/Donation/API owners. After required owner gates converge, continue counterpart reconciliation per the approved Techplan; retain downstream runtime verification. Do not infer WU-S2-006 completion from this Review.
- **Context refs:** Run Invocation `.../RV-S2-006-006/invocation.md`; approved Techplan `.../TP-S2-006-008/techplan.md` §§3–13; `api/README.md`; `../harscode-workspace/best-practices/restapi/openapi-spec-first-drift.md`; the seven source paths and identities listed in the Invocation.
