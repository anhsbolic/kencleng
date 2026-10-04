> Phase: Code Review  
> Work Unit / Run: `WU-S2-007` / `RV-S2-007-002`  
> Author: `P-S2-007-RV-002-1` / Reviewer (`KC-REVIEWER`)  
> Created: 2026-10-04  
> Model / Reasoning: `gpt-6-luna` / `high`  
> Session: fresh; identifier not exposed  
> Target revision: `4e71d3697a479d92b2cad5a79ef5d710f4e8077c` plus current working tree  
> Workflow revision: `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`

# Review Findings — `RV-S2-007-002`

## Review basis and scope

Reviewed the current diff for exactly the five pinned files against Approved Techplan `techplan.md` SHA-256 `ff0cb704151495e9f5539d11c24e808e8618d1c98cdb0d680957df8d51948dc3`. All current hashes match the Run Invocation. The three read-only comparison anchors also match their pinned hashes. HEAD is `4e71d3697a479d92b2cad5a79ef5d710f4e8077c`; unrelated pre-existing working-tree changes remain outside this review.

| Pinned file | Current SHA-256 | Match |
|---|---|---|
| `api/openapi/donation.yaml` | `9f7c31065c3c7ffa77491541102afc0ce84085f50a1e2aec317336c9656d5c1d` | Yes |
| `api/openapi.yaml` | `c37038ecb5088aebcc9156e495d8b14bfa7cc405debeb940e16d7b17786d1838` | Yes |
| `frontend/lib/api/generated/openapi.ts` | `f267598c563b4bf73442f120c81fdf6ec1b7f48cd29095f78f3622dbd7094e20` | Yes |
| `frontend/lib/api/donation.ts` | `5c3cdca3b71d806ece0c15bd1dbd8a0a162895ec0c56a4b63c3271f7fac6d654` | Yes |
| `frontend/app/donations/donation-flow.test.tsx` | `e7b66525258a8601746583274a7ae9d794f0795254aba794e29a46f153fb6a38` | Yes |

The current diff was read directly. The Build report was used as context only. The Approved Techplan was verified at its pinned hash.

## 1. Safety

No findings.

The authored Donation POST contract adds only a generic `application/problem+json` 503 using the existing `Problem` schema. It states fail-closed/no-admission and excludes internal reasons and `Retry-After`. No shared Problem schema, caller UI/copy, backend behavior, or runtime claim is changed.

`submitDonation` checks the operation's documented 503 before the general `>= 500` branch and returns the existing `request-failure`. Transport exceptions still become `DonationTransportError`; other 5xx responses remain `ambiguous`. The existing ambiguous retry path retains the same key and payload. This preserves the Techplan boundary that only Donation POST 503 is definitive non-admission.

No specialized concurrency, performance, or security test area was introduced beyond the Techplan Test Focus Pointer. The changed scope adds no runtime admission or protected backend behavior.

## 2. Quality

No findings.

The response classification is a small explicit branch placed before the generic 5xx case. The focused consumer test asserts the user-visible generic failure, no ambiguous retry control, and one POST call. Existing retry evidence checks key and payload equality.

## 3. Stack-Specific Best Practices

**Finding BP-1 — non-blocking.** `frontend/app/donations/donation-flow.test.tsx:143–147` constructs the generic Problem fixture as an untyped object literal. Its current `type`, `title`, and `status` fields match the current Problem contract, but the fixture can drift silently if the generated schema changes. Suggested resolution: type the fixture with `components["schemas"]["Problem"]` (or use `satisfies`) while retaining the network-level MSW override and observable assertions. This is a minor test-maintenance improvement; it does not block the present consumer behavior or require a patch loop.

Source: `../harscode-workspace/best-practices/react/component-test-mocking-discipline.md` says to mock at the network layer with the same OpenAPI-generated response shape so a hand-rolled fixture cannot silently drift, and to assert user-observable behavior. The test already follows the network-layer and observable-assertion guidance; only the response fixture's generated-schema tie is missing.

REST API guidance was also routed for the changed status semantics and idempotent retry boundary. The diff follows the operation-specific status contract and preserves idempotency intent for ambiguous outcomes; no REST best-practice finding was identified.

## 4. Consistency

No findings.

The API change is in the authored split Donation source and reuses the relative shared `Problem` reference as required by `api/README.md`; aggregate and TypeScript generated surfaces carry the same operation response. The generated operation response is limited to the Donation POST contract. The shared component and UI caller remain untouched.

The helper and test fit `frontend/AGENTS.md`'s generated-contract and MSW network-boundary rules. OI9 in `docs/spec/5-donation/invariants.md` and the accepted-funding-unavailable criterion/error behavior in `docs/spec/5-donation/features/01-submit-donation-settlement.md` agree with the exact 503 classification. The test uses generic Problem fields and makes no backend availability, settlement, or runtime claim. No copy or state was added.

## Verification executed during Review

No runtime reproduction or test was needed to resolve a concrete review uncertainty. Review was reasoning-first; current hashes, pinned anchors, and the five-file diff were inspected. Build-reported checks were not rerun and are not claimed as independent Review verification.

## Verdict

**Approve with minor comments.** BP-1 is non-blocking and does not require a patch plan. No blocking findings; independent Testing is the normal next phase after orchestration routing.
