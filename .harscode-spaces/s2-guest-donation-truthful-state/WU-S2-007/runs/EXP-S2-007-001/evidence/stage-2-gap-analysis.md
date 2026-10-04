# Stage 2 — Gap Analysis: Donation POST Funding-unavailable response

## Provenance

- Phase/Stage: Exploration / Stage 2
- Author: Explorer (`P-S2-007-EXP-001-1`)
- Created: 2026-10-04
- Model / reasoning: `gpt-6-luna` / `medium` (per Invocation)
- Session: not exposed in the execution interface
- Work Unit / Run: `WU-S2-007` / `EXP-S2-007-001`
- Target revision: `4e71d3697a479d92b2cad5a79ef5d710f4e8077c` plus the current working tree (per Invocation)
- Workflow revision: `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4` (per Invocation; ordinary current-effective guidance)
- Invocation: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/EXP-S2-007-001/invocation.md`

This artifact records exploration evidence, not a source change or solution decision. No tests, generators, or validation commands were run.

## Area 1 — Accepted Donation policy

**Requirement.** Donation invariant `INV-donation-02` and feature `01-submit-donation-settlement.md` acceptance criterion on funding capacity require that when authoritative settled Funding is unavailable, POST fails closed: absent Funding is not zero, no new Donation is admitted, and the response is a generic `503 Problem` without internal reason or `Retry-After`. The feature's error-behavior table repeats this expectation; Resolved OI9 records it as the 2026-10-04 Donation/API owner decision. These sources also state that authored API and generated/internal counterparts need separate coordinated reconciliation before delivery/runtime progression.

**Current state / gap.** This policy is explicit and consistent in the current accepted invariant and feature. No policy gap was found for this Work Unit; contract bytes and consumer counterparts are still unaccepted, as the WU manifest states.

**Sniffing.**

- Risk: admitting a new Donation against unknown Funding would violate fail-closed admission; disclosing internal reason or an unsupported recovery bound would disclose unsupported detail.
- Edge cases: unavailable/absent Funding must not be conflated with zero; do not add `Retry-After` absent a proven recovery bound.
- Miscontext: the task is contract/counterpart reconciliation, not backend runtime implementation or an unresolved OI9 policy decision.
- Misleading signals: a generic Problem schema already exists, but that does not mean the POST operation declares the required response.
- Inconsistency: none within the accepted Donation sources reviewed.

**Anchors.** `docs/spec/5-donation/invariants.md` — `INV-donation-02` and Resolved OI9; `docs/spec/5-donation/features/01-submit-donation-settlement.md` — funding-capacity acceptance criterion, error-behavior table, Resolved OI9. These define the requirement and its security/truth boundary.

## Area 2 — Authored Donation OpenAPI

**Requirement.** The POST operation must expose the accepted generic `503 Problem` for unavailable authoritative settled Funding. `api/README.md` §§ authored sources/generated outputs states split OpenAPI files are authored authority and generated bundle/types must correspond to them.

**Current state.** `api/openapi/donation.yaml` defines `POST /campaigns/{campaignId}/donations`. Its description covers eligibility, cap/capacity, 422, 409, idempotency and settlement. Its response map declares `201`, shared `422 ValidationError`, and `409` with the shared `Problem` schema. It does not declare a `503` response or describe the Funding-unavailable case. `api/openapi/common.yaml` already defines the `Problem` schema.

**Gap.** The authored POST contract is missing the accepted Funding-unavailable response and associated behavior description. The current local `Problem` schema reference is available; whether any shared component impact exists remains to be confirmed during solutioning/reconciliation.

**Sniffing.**

- Risk: callers cannot rely on the authored operation contract for this server response.
- Edge cases: absent Funding must be distinguished from zero; no `Retry-After` is supported by current policy.
- Miscontext: existing operation descriptions about eligible-capacity no-fit (`422`) and closed Campaign (`409`) do not cover Funding being unavailable.
- Misleading signals: `Problem` is already used by `409`, which may look like error coverage but does not declare `503`.
- Inconsistency: authored responses omit a case required by the accepted feature/invariant.

**Anchors.** `api/openapi/donation.yaml` — `/campaigns/{campaignId}/donations` → `post` → `description` and `responses`; `api/openapi/common.yaml` — `components.schemas.Problem`. These are the authored operation and existing shared error schema.

## Area 3 — Generated API counterparts

**Requirement.** The bundled `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts` are generated from authored split sources per `api/README.md`; they must represent the reconciled authored operation after generation.

**Current state.** The bundle's `POST /campaigns/{campaignId}/donations` response map contains `201`, `409`, and `422`, but no `503`. The generated TypeScript operation response union likewise contains `201`, `409`, and `422`, but no `503`.

**Gap.** Both generated counterparts lack the operation response, consistent with and downstream of the authored-source gap. Their correspondence must be re-established after the authored source is reconciled; neither is a separate authority.

**Sniffing.**

- Risk: generated consumers cannot statically see the new documented status response.
- Edge cases: status response typing should not imply `Retry-After` or internal error details that policy excludes.
- Miscontext: the bundle contains `503` operations elsewhere (e.g. Campaign operations), which does not establish support for this Donation POST.
- Misleading signals: `Problem` may exist elsewhere in generated schemas, but the Donation POST response union omits 503.
- Inconsistency: generated representations agree with each other and mirror the current source, but all are behind the accepted policy.

**Anchors.** `api/openapi.yaml` — `/campaigns/{campaignId}/donations` → `post.responses`; `frontend/lib/api/generated/openapi.ts` — corresponding path → `post.responses`. These are the generated bundle and consumer type coordinates.

## Area 4 — Internal/repository-development consumers

**Requirement.** WU-S2-007 scope requires affected generated/internal consumer counterparts for the response while excluding frontend UX and backend/runtime implementation. WU-S2-003 has a hard dependency on this Work Unit's accepted source/counterpart handoff.

**Current state.** `frontend/lib/api/donation.ts` returns `ambiguous` for every HTTP status `>= 500`. `DonationClient.send` retains the same retry intent and idempotency key for that result and displays that the submission outcome cannot be determined. `frontend/app/donations/donation-flow.test.tsx` covers a generic ambiguous retry, not a specifically identified Funding-unavailable 503 case. `frontend/mocks/handlers/donation.ts` has no Funding-unavailable POST response path; existing campaign-detail 503 mock behavior is for a different GET surface. Work Graph and WU-S2-003 manifest identify contract reconciliation as a prerequisite to WU003 candidate refresh/review/approval and affected Build. The current distribution posture is internal/repository-development; no external consumer is identified in the current Work Graph evidence.

**Gap / observation.** The consumer's generic 5xx treatment is conservative and keeps retries on the same idempotency key, but does not distinguish the contract's definitive `503`/no-admission meaning from other 5xx cases whose outcome may be ambiguous. The consumer therefore surfaces a less-specific outcome for this known response. Whether a non-UX consumer/type/fixture adjustment is required within this Work Unit is a Stage-3 scope decision; no UI copy/UX change is authorized here. The mock also does not currently model this POST case.

**Sniffing.**

- Risk: a caller may treat the definitive no-admission response as an unknown outcome; retry still uses the same key/payload, preserving the accepted retry posture.
- Edge cases: distinguish a contract-defined 503 from other 5xx statuses whose commit outcome may be unknown; preserve same-key behavior for ambiguous cases.
- Miscontext: existing generic 5xx handling does not encode the new response-specific semantic guarantee.
- Misleading signals: Campaign GET 503 mocks/tests exist, but they are not a POST Donation Funding-unavailable fixture.
- Inconsistency: consumer behavior is safer than resubmitting with a new key, but its generic result/message is less precise than the accepted POST contract; no contradictory admission or new-key behavior was found.

**Anchors.** `frontend/lib/api/donation.ts` — `SubmitDonationResult` and `submitDonation` 5xx branch; `frontend/app/campaigns/[campaignId]/donate/donation-client.tsx` — `send` ambiguous branch; `frontend/app/donations/donation-flow.test.tsx` — ambiguous retry case; `frontend/mocks/handlers/donation.ts` — POST handler. These locate consumer classification, retry behavior, regression coverage, and development mock behavior.

## Revision evidence and scope

The current hashes of all six invocation-listed policy/API/generated inputs match the exact SHA-256 values in the Invocation at Stage-2 inspection. The checkout also has other pre-existing modified/untracked Work Unit and project files; they were not changed by this Run. No source, spec, contract, generated, consumer, fixture, or test files were modified.

## Progression assessment

- **Finding F1 — API/counterpart omission:** authored Donation POST plus generated bundle/types lack the accepted 503 case. **Decision-relevant**, but not blocking Stage 3: policy is settled, and Stage 3 can select the bounded reconciliation direction. The omitted response blocks WU-S2-007 completion and its WU-S2-003 dependency until source/counterparts are reconciled and accepted.
- **Finding F2 — generic 5xx consumer semantics:** the consumer labels the definitive no-admission 503 as ambiguous and the POST mock/test has no dedicated case. **Decision-relevant** for deciding which internal counterparts are affected; Stage 3 can proceed while respecting the explicit no-UX/no-runtime boundary. No active inability to perform solutioning was found.
- No additional authority decision or active Blocker surfaced in Stage 2. Exact bytes, API owner acceptance, independent review/verification, and downstream WU003 handoff remain completion obligations, not findings that prevent Stage 3 analysis.
