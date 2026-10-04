# Tech Plan: Donation POST Funding-unavailable contract

> Phase             : Techplan
> Ticket            : WU-S2-007
> Work Unit ID      : WU-S2-007
> Run ID            : TP-S2-007-001
> Author            : P-S2-007-TP-001-1
> Participant ID    : P-S2-007-TP-001-1
> Profile           : KC-PLANNER
> Role              : Planner
> Model             : gpt-6-luna (Invocation-configured)
> Reasoning         : high (Invocation-configured)
> Session           : not exposed
> Created           : 2026-10-04
> Updated           : 2026-10-04
> Target revision   : 4e71d3697a479d92b2cad5a79ef5d710f4e8077c plus current working tree at dispatch
> Workflow revision : pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4 (ordinary current-effective guidance)
> Status            : Approved
> Approach          : Document the accepted no-admission 503, regenerate internal contract types, and classify that response through the existing definitive-failure path.
> Refs              : Donation OI9; `runs/TP-S2-007-001/invocation.md` (current-effective inputs/hashes); `runs/EXP-S2-007-001/evidence/stage-2-gap-analysis.md`; `runs/EXP-S2-007-001/evidence/stage-3-solutioning.md`; `api/README.md`

---

## 1. Background

Accepted Donation policy requires a generic `503 Problem` when authoritative settled Funding is unavailable. The request fails closed, admits no new Donation, and exposes neither an internal reason nor `Retry-After`. The authored Donation POST OpenAPI operation and its generated representations omit this response. The internal TypeScript helper currently treats every 5xx as ambiguous, so it preserves retry intent even for this contract-defined no-admission response.

The policy is settled; this Work Unit reconciles its authored API and affected internal/generated counterparts. It does not implement or claim backend runtime behavior.

## 2. Scope

**In scope:**
- Add the `503` response to `POST /campaigns/{campaignId}/donations` in the authored split OpenAPI source, referencing the existing shared `Problem` schema.
- Describe the unavailable-Funding condition and its fail-closed, no-internal-reason, no-`Retry-After` contract.
- Regenerate the aggregate OpenAPI bundle and frontend OpenAPI TypeScript types from their authored sources.
- Map this operation's documented `503` to the existing `request-failure` result; retain `ambiguous` for transport failures and other 5xx statuses.
- Add focused MSW-backed consumer evidence for the known 503 and for preservation of ambiguous retry behavior.
- Record exact changed source/counterpart hashes and obtain the API owner's acceptance of the concrete bytes before WU-S2-007 completion or WU-S2-003 dependency release.

**Out of scope (explicit):**
- Product/domain policy changes, changes to Donation invariant/feature specs, or changes to WU-S2-006's terminal state or its exact accepted source-hash snapshot.
- Backend handlers, Funding reads/admission logic, transaction/locking/idempotency implementation, migrations, database actions, runtime claims, or protected paths.
- New frontend state, branch-specific copy, layout, or other UX behavior. The existing generic `request-failure` presentation remains in use.
- Changes to shared OpenAPI components unless live requirements reveal a real schema gap; the current `Problem` schema is sufficient.
- Changes to unrelated Donation statuses/retry behavior, public/external rollout, or WU-S2-003/WU-S2-004 implementation.
- `report-techplan.md`, source edits, generators, validators, tests, Review, Build, and downstream dispatch in this planning Run.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | When authoritative settled Funding is unavailable, the Donation POST contract is generic `503 Problem`; no Donation is admitted, and no internal reason or `Retry-After` is exposed. | `docs/spec/5-donation/invariants.md` — `INV-donation-02`, Resolved OI9; `docs/spec/5-donation/features/01-submit-donation-settlement.md` — funding-capacity criterion and Resolved OI9 |
| Q2 | The authored split OpenAPI is authoritative; the aggregate bundle and generated TypeScript types must correspond to it. | `api/README.md` — Editing workflow; Stage-2 Areas 2–3 |
| Q3 | The internal consumer treats the contract-defined Donation POST 503 as definitive non-admission via the existing generic failure result, while transport failures and other 5xx results remain ambiguous and preserve same-key/same-payload retry intent. | Stage-3 Option A; `frontend/lib/api/donation.ts` — `submitDonation`; Donation `send` caller |
| Q4 | Consumer/mock evidence must cover the new response without creating user-facing state/copy or implying live backend behavior. | WU-S2-007 manifest boundaries; Stage-3 Option A; root `AGENTS.md` simulation rule |
| Q5 | Preserve WU-S2-006 as `DONE`, including its owner-accepted exact source revisions. WU-S2-003 remains blocked on this Work Unit's exact source/counterpart owner acceptance; no completion is inferred here. | WU-S2-006 `manifest.md` §§ Current State / Spec acceptance snapshot; parent Work Graph; WU-S2-007 manifest |

## 4. Rules & Validation

- **R1 — Contract response:** Given a Donation POST for which authoritative settled Funding is unavailable, the authored operation declares `503` with `application/problem+json` using the shared `Problem` schema. Its description/example remains generic, states no Donation is admitted, and does not specify an internal reason or `Retry-After` header.
- **R2 — Generated correspondence:** After generation, `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts` contain the Donation POST `503` response corresponding to the authored split source, without unrelated generated contract changes.
- **R3 — Known non-admission classification:** Given Donation POST status `503`, `submitDonation` returns the existing `request-failure` result; the current caller clears retry intent and uses its existing generic failure presentation without a new UI state or copy.
- **R4 — Ambiguous outcomes preserved:** Given transport failure or any other 5xx status, the result remains ambiguous and the caller retains the original idempotency key and payload for retry. Existing `201`, `409`, and `422` classifications remain unchanged.
- **R5 — Contract-facing evidence:** Focused MSW consumer evidence represents a generic Problem 503 for this POST, verifies the known-failure path has no ambiguous retry action, and continues to verify same-key/same-payload retry for an ambiguous transport/5xx result. It must not imply that the mock establishes backend availability or settlement behavior.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Stage-3 Option A: classify the documented Donation POST `503` as known non-admission using `request-failure`; preserve ambiguity for transport failures and other 5xx statuses. | Chosen | Reflects accepted no-admission semantics while keeping unknown outcomes conservative. Existing caller clears retry intent for `request-failure`; a later deliberate attempt is a new request. |
| D1-alt | Keep all 5xx statuses ambiguous (Stage-3 Option B). | Rejected | Would preserve safe same-key retries but continue presenting this explicitly known no-admission result as uncertain. |
| D2 | Reuse the existing generic `Problem` schema and existing generic consumer failure presentation; add no tailored UI behavior. | Chosen | No shared-schema gap or UX decision is needed. No internal reason, unsupported recovery bound, `Retry-After`, or live-runtime claim is introduced. |
| D3 | Use test-scoped MSW response override in the existing Donation flow test for the Funding-unavailable 503. | Chosen | Current tests already override the POST handler; this proves the consumer contract without adding mock-only Funding state or suggesting that a mock implements the backend. Leave the shared handler unchanged unless Build finds that a contract-faithful reusable fixture is necessary; any broader need must be reported for scope reconciliation. |
| D4 | Keep WU-S2-006 terminal and its exact accepted source-hash receipt as historical/current acceptance evidence. | Chosen | WU-S2-007 is a new coordinated reconciliation after that terminal boundary. The later OI9 spec bytes and this Work Unit's future API/counterpart bytes do not rewrite or extend WU-S2-006's accepted snapshot. |

## 6. Backward Compatibility

- **Existing data:** No persistence or stored-data changes.
- **API/contracts/clients:** Adding the documented response is additive for the current owner-confirmed MVP1 distribution posture: no external consumer depends on this contract, and current consumers are internal/repository-development. This posture is scoped to the current rollout and is not a universal future compatibility claim. Keep the generic `Problem` body; no new header or response field is introduced. Exact authored and generated counterpart bytes still require API-owner acceptance.
- **Migration/deprecation compatibility:** Not applicable; no migration, endpoint version, or deprecation is planned.
- **Prior Work Unit boundary:** WU-S2-006 remains `DONE`. Its exact seven-file source acceptance receipt remains recorded in `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/manifest.md` under “Spec acceptance snapshot — 2026-10-02” and in `runs/RV-S2-006-006/invocation.md`. Do not substitute the current OI9-bearing spec hashes for that historical accepted snapshot. WU-S2-003's HARD dependency remains unsatisfied until WU-S2-007's completion condition is met.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| RISK-1 | Treating the contract-defined 503 as ambiguous retains retry intent and misstates a known no-admission outcome. | Medium | Medium | Map the Donation POST 503 before the generic `>=500` branch and test its existing `request-failure` caller path. Keep all other 5xx and transport failures ambiguous. No backend behavior is inferred. |
| RISK-2 | A future 503 meaning is broader than this accepted operation contract, making status-only classification insufficient. | Low for current scoped contract | High | Current policy/operation define this Donation POST 503 as unavailable authoritative Funding with no admission. If Build finds another accepted 503 meaning or cannot preserve that operation-level meaning, stop and surface the exact authority/contract gap; do not invent a discriminator or detail. Future contract expansion needs renewed review. |
| RISK-3 | An OpenAPI/generated mismatch leaves internal consumers unable to rely on the new response. | Medium | Medium | Generate bundle and types from the authored split source; Build and independent Testing inspect the operation response and generated diff. |
| RISK-4 | Mock evidence could be mistaken for real Funding or settlement behavior. | Low | Medium | Use only test-scoped MSW override and label it as contract evidence. Do not change production API behavior, campaign fixtures, or state claims. |
| RISK-5 | Current internal-only compatibility posture may change before a future public rollout. | Low for current rollout | High | Scope the additive-change claim to the owner-confirmed MVP1 posture; future external distribution requires its own owner review. |

## 8. Interface Contract

**Persistence/data shape:** None.

**API/event/external interface:** `POST /campaigns/{campaignId}/donations` adds response status `503`, media type `application/problem+json`, and the existing `components.schemas.Problem` reference from `api/openapi/common.yaml`. The operation description identifies unavailable authoritative settled Funding as fail-closed/no-admission. Do not expose an internal reason or declare `Retry-After`. Do not change the shared `Problem` component or add an unapproved discriminator/example value.

**Cross-layer/business boundary:** In `submitDonation`, recognize the operation's documented status `503` as `request-failure` before the existing generic `response.status >= 500` branch. The current `DonationClient.send` caller already clears retry intent for `request-failure` and leaves retry intent for `ambiguous`; keep its existing text and UI state. Transport exceptions continue to throw `DonationTransportError` and follow the existing ambiguous retry catch path.

## 9. Architecture / Plan

1. Update only the authored Donation POST description/response in `api/openapi/donation.yaml`, following the existing local response style and shared `Problem` reference.
2. Run the repository bundle command, then generate frontend types from the bundle. Review output diffs against only the Donation POST response; do not hand-edit generated files.
3. Update `submitDonation` so status `503` returns `request-failure` before the catch-all 5xx branch. Do not alter the caller's generic error text/state.
4. Add a test-scoped MSW 503 override and focused flow assertions for the no-retry known-failure path, while preserving the existing ambiguous same-key/same-payload retry assertion for transport or another 5xx response.
5. Build performs focused generation/contract/consumer checks; independent Testing performs the planned final validation and frontend baseline. Obtain independent Techplan Review recommendation consideration and the normal Human approval gate before any Build dispatch. After source and counterpart bytes exist, route their exact identities for API-owner acceptance; do not release the WU-S2-003 dependency before that gate.

The change is a linear authored-contract → generated-counterpart → internal-consumer sequence; no diagram is needed.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `api/openapi/donation.yaml` — `/campaigns/{campaignId}/donations` → `post.description` and `post.responses` | Authored operation currently documents capacity no-fit `422`, ineligible `409`, and idempotency but omits unavailable Funding `503`. | Add the accepted fail-closed condition to the description and a local `503` response with `application/problem+json` and `./common.yaml#/components/schemas/Problem` (following this file's existing cross-file reference style). No `Retry-After` header or internal reason. |
| `api/openapi/common.yaml` — `components.schemas.Problem` | Existing shared RFC 9457 base schema requires `type`, `title`, and `status`; current schema is sufficient. | Read-only reference; leave the shared component unchanged. |
| `api/openapi.yaml` — `/campaigns/{campaignId}/donations` → `post.responses` | Generated aggregate currently mirrors the authored omission. | Regenerate with `cd api && npm run bundle`; never hand-edit. |
| `frontend/lib/api/generated/openapi.ts` — corresponding path → `post.responses` | Generated operation response union lacks `503`. | Regenerate with `cd frontend && npm run generate:api-types`; never hand-edit. |
| `frontend/lib/api/donation.ts` — `SubmitDonationResult`, `submitDonation` | Helper has existing `request-failure` and generic `>=500` → `ambiguous` branches. | Return `request-failure` specifically for status `503` before generic 5xx handling. Preserve all other result semantics and transport exception behavior. |
| `frontend/app/campaigns/[campaignId]/donate/donation-client.tsx` — `DonationClient.send` | Existing non-ambiguous result clears retry intent and uses generic failure text; ambiguous result retains payload/key and shows retry UI. | No production change intended. Reopen to confirm behavior at implementation time. |
| `frontend/app/donations/donation-flow.test.tsx` — POST error and ambiguous retry tests | Existing MSW overrides exercise 409/422 and transport ambiguity. | Add focused 503 override/assertion for generic request failure/no retry; retain/assert ambiguous same-key/same-payload behavior for transport or another 5xx. Keep assertions observable and avoid backend claims. |
| `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/manifest.md` — “Spec acceptance snapshot — 2026-10-02”; `runs/RV-S2-006-006/invocation.md` — exact source identities | Owns the terminal WU006 acceptance receipt and exact accepted source hashes. | Read-only historical acceptance pointer; do not amend it or represent the OI9 bytes as part of WU006. |

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `api/openapi/donation.yaml` | Authored API | Document Donation POST unavailable-Funding `503 Problem`. |
| `api/openapi.yaml` | Generated API bundle | Regenerate from authored split OpenAPI. |
| `frontend/lib/api/generated/openapi.ts` | Generated TypeScript | Regenerate from the bundle. |
| `frontend/lib/api/donation.ts` | Internal consumer | Classify the contract-defined 503 as existing `request-failure`. |
| `frontend/app/donations/donation-flow.test.tsx` | Consumer evidence | Focused MSW contract/behavior evidence and ambiguous retry preservation. |

| File / area intentionally untouched | Why |
|---|---|
| `docs/spec/5-donation/invariants.md`, `docs/spec/5-donation/features/01-submit-donation-settlement.md` | Accepted OI9 policy is already explicit; this task reconciles API/counterparts only. |
| `api/openapi/common.yaml` | Existing `Problem` schema covers the response; no shared-definition gap is present. |
| `frontend/mocks/handlers/donation.ts` and `frontend/mocks/fixtures/` | A test-scoped MSW override represents the response without inventing shared mock Funding state. Reassess only if the test proves a reusable fixture is required. |
| `frontend/app/campaigns/[campaignId]/donate/donation-client.tsx` | Existing generic request-failure and ambiguous caller behavior already matches the desired distinction; no UI change. |
| Backend production code, protected paths, database/migrations, runtime and browser surfaces | Explicitly outside this Work Unit and Run envelope. |
| WU-S2-006 manifest, accepted snapshot, Events, Work Graph, Control Surface, Outcome | This Planner Run is not authorized to reconcile orchestration state; WU006 terminal state and exact source acceptance remain intact. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | Build: `cd api && npm run validate`; inspect the authored Donation POST 503 media type/schema/description and confirm no `Retry-After` or internal reason is introduced. Testing: repeat validation and inspect the final source contract. | Testing | Confirms the only new public contract represents the accepted no-admission policy; skipping can leave an invalid or disclosing API shape. |
| R2 | Build: `cd api && npm run bundle`, then `cd frontend && npm run generate:api-types`; inspect that the 503 appears in both generated operation surfaces and that generated diffs are limited to this response. Testing: independently inspect the committed bundle/type response against the authored operation and generated diff. | Testing | Generated clients are relied on as contract evidence; drift can leave runtime clients unable to type the status or introduce unrelated API changes. |
| R3 | Build: run the focused flow test with a test-scoped MSW 503 response and assert the existing request-level generic failure, cleared retry intent, and absence of ambiguous retry control. Testing: include the case in `cd frontend && npm run verify`. | Testing | Protects the semantic distinction that motivated the change without changing user copy/state; skipping can preserve an incorrect retry affordance or introduce a new UX path. |
| R4 | Build: focused flow evidence for transport failure or a non-503 5xx retains the same idempotency key and payload on retry; confirm 201/409/422 paths remain. Testing: run `cd frontend && npm run verify` and inspect the focused assertions. | Testing | Retrying unknown admission with a new key can risk duplicate intent; mapping unrelated errors as definitive can discard safe retry context. |
| R5 | Build: focused MSW override returns only generic Problem fields/status; assert consumer behavior without relying on internal server details. Testing: inspect assertions and run the focused/full frontend tests through `npm run verify`. | Testing | Prevents contract mock evidence from implying server implementation or disclosing detail; skipping weakens confidence in the exact client-facing response. |

Build may run newly authored focused tests and generators to establish its changes. Testing owns authoritative independent final evidence. No generator, validator, test, or runtime check was run in this Planner Run.

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Backend Funding-admission atomicity / concurrency and runtime non-disclosure | Donation admission is money-sensitive; these guarantees are owned by backend runtime and are not changed or evidenced by this contract/client reconciliation. | `runs/EXP-S2-007-001/evidence/stage-2-gap-analysis.md#Area 1 — Accepted Donation policy` | N/A for specialized testing in WU-S2-007 — no backend/runtime change is in scope; API shape is covered by R1. Do not claim runtime admission/non-disclosure verification; retain the backend runtime obligation downstream. |

## 13. Open Items

### Active — needs external input or verification

1. **Exact API/counterpart owner acceptance remains pending.** After Build and applicable independent Review/Testing, the API owner must accept the exact authored source and generated/internal counterpart bytes/hashes. Until then, WU-S2-007 is incomplete and WU-S2-003's HARD dependency remains blocked. This is a completion gate, not a policy decision needed to finish this synthesis.

### Resolved — retained as decision history

1. ~~**Unavailable Funding response policy**~~ **RESOLVED — 2026-10-04, Anhar Solehudin, Donation/API owner:** generic `503 Problem`, no internal reason or `Retry-After`, no new Donation admitted; recorded as OI9 in the Donation invariant and feature spec. Concrete API/counterpart bytes remain a separate acceptance gate.
2. ~~**Consumer classification direction**~~ **RESOLVED — Stage-3 recommends Option A:** map the contract-defined Donation POST 503 to existing `request-failure`; keep transport and other 5xx outcomes ambiguous with current same-key/same-payload retry behavior. No new UX state/copy.
