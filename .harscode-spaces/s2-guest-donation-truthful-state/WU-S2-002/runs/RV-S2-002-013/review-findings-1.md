# Review Findings — `RV-S2-002-013`

> Phase: Code Review  
> Work Unit: `WU-S2-002`  
> Run: `RV-S2-002-013`  
> Role / specialization: Reviewer / independent four-pass review of Task 02 Donation OpenAPI reconciliation and generated outputs  
> Participant: `P-S2-002-RV-013-1`  
> Author: Codex Reviewer  
> Created: 2026-10-01  
> Model / reasoning: `gpt-6-luna` / `high` (Invocation configuration; runtime selection not independently exposed)  
> Target revision: `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus BLD-006's four-path diff  
> Workflow revision: `33b03a3f62cc3aacba6534b8a011465613c64b09`

## 1. Safety

No findings.

The status lookup exposes only `DonationStatus`; its public failure response specifies a single `404` Problem Details example and `Cache-Control: private, no-store` for absent Donation and missing, wrong, or expired credentials. The descriptions distinguish the settled fragment handoff, one-way HMAC verifier, 24-hour expiry, QRIS sandbox behavior, and server-owned terminal state from controls or empirical proof that remain open. No runtime/security claim is presented as verified.

TP-015's Test Focus Pointer explicitly requires specialized downstream evidence for status-credential exposure/parity/abuse and settlement atomicity/replay/concurrency. No new Techplan drift was identified.

## 2. Quality

No findings.

The changed operations and schemas describe the guest submission/status behavior clearly. The bundle and generated types inspected for the changed Donation schemas reflect the authored source changes; no unexpected generated structure was observed.

## 3. Stack-Specific Best Practices

No blocking findings. The routed `restapi/idempotency-and-versioning.md` trigger matches the submission POST: the shared `Idempotency-Key` contract requires a key per logical donation, retains it for retries, distinguishes same-payload from changed-payload reuse, and separates request idempotency from settlement replay. The routed `restapi/anti-enumeration.md` trigger matches the credential-protected status lookup: the authored contract specifies uniform public failure shape/cache behavior, while constant-time secret comparison, timing parity, and abuse controls remain clearly assigned to implementation/Security/Testing evidence. The routed `restapi/openapi-spec-first-drift.md` trigger matches the generated artifacts: split sources remain authored, and the bundle/types reflect the inspected source changes. These are contract-level observations, not runtime proof.

## 4. Consistency

### F-1 — Monetary response projections allow the currency code to be omitted

- **Location:** `api/openapi/donation.yaml:358-380` (`DonationListItem`), `:394-424` (`ClaimableDonation`), and `:437-468` (`MyDonation`); generated counterparts in `frontend/lib/api/generated/openapi.ts`.
- **Problem:** The diff adds `currency_code` beside `amount` in these three response projections, but none declares either field in `required`. The generated TypeScript therefore exposes both as optional (for example, `DonationListItem.currency_code?: string`). The contract permits a monetary value to be returned without its explicit currency code.
- **Impact:** This does not satisfy the approved project-wide representation, which requires API/wire monetary values to use a major-unit decimal string together with an explicit currency code. Consumers cannot rely on the required pair even though these fields were changed to adopt that representation.
- **Suggested correction:** In each changed response projection, require the monetary pair (`amount` and `currency_code`) together, then regenerate `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts` from the authored split source. Keep this correction limited to the three projection schemas touched by this diff.
- **Status:** Blocking.
- **Authority:** `docs/project/kencleng-monetary-data-standard.md` (“API and wire monetary values use a major-unit decimal string together with an explicit currency code”); TP-015 Q12; `api/README.md` (split source owns the contract and generated bundle/types must be regenerated).

## Verification executed during Review

`git diff --check -- <RV-S2-002-013>/review-findings-1.md <RV-S2-002-013>/patch-plan-1.md` → PASS for the two Review artifacts only. Review used the actual four-path diff against target revision `650e73c5d646c29c0ddf1931618f02685d15f7b7` and inspected the current authored sources, generated counterparts, and named authorities. No product tests or broad verification were run; Build's validation/generation evidence is context only.

## Verdict

Request changes.

F-1 is blocking. No runtime, security, concurrency, or empirical-parity evidence is claimed by this Review.

## Phase handoff

- Completed: four-pass independent Review of the exact BLD-006 four-path diff.
- Artifacts: this file; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-013/patch-plan-1.md`.
- Human decision: none required to make the narrow schema correction; applicable API/domain-owner acceptance remains a later gate.
- Open / deferred: F-1 blocks approval; runtime/security/PII controls and evidence, empirical `404` parity, and remaining owner acceptance stay downstream as specified by TP-015.
- Recommended next step: fresh Build/Patch Run for F-1 only, followed by the applicable targeted confirmation and downstream Testing/owner-acceptance route.
- Session transition: start a new Build/Patch Run with a new Participant and fresh Session, re-grounded on TP-015, Task 02, this finding, and the four-path diff.
- Context pointers: TP-015; Task 02; F-1 locations above; patch plan.
