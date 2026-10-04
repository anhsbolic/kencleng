# Stage 3 — Solutioning: Donation POST Funding-unavailable response

## Provenance

- Phase/Stage: Exploration / Stage 3
- Author: Explorer (`P-S2-007-EXP-001-1`)
- Created: 2026-10-04
- Model / reasoning: `gpt-6-luna` / `medium` (per Invocation)
- Session: not exposed in the execution interface
- Work Unit / Run: `WU-S2-007` / `EXP-S2-007-001`
- Target revision: `4e71d3697a479d92b2cad5a79ef5d710f4e8077c` plus the current working tree (per Invocation)
- Workflow revision: `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4` (per Invocation; ordinary current-effective guidance)
- Prior Exploration evidence: `runs/EXP-S2-007-001/evidence/stage-2-gap-analysis.md`

This records a recommended reconciliation direction for downstream planning. No contract, generated, or consumer source was changed, and no tests or generators were run.

## Problem and current context

Accepted OI9 requires Donation POST to fail closed with generic `503 Problem` and no new Donation when authoritative settled Funding is unavailable; it excludes internal reason and `Retry-After`. The authored Donation operation and its generated bundle/types omit that response. The frontend API helper maps every 5xx to `ambiguous`; its consumer retains the same key/payload and displays uncertainty. WU-S2-007 allows affected internal/repository-development consumer reconciliation but excludes frontend UX and backend/runtime changes.

The distinction matters: transport/server failures may leave admission unknown, while this particular contract response means no Donation was admitted. A response-specific internal classification can reuse the existing generic non-ambiguous request-failure path without adding a new UI state or changing UI copy.

## Options

### Option A — Model 503 as a known non-admission in the existing consumer path (recommended)

Document the POST response in the authored Donation OpenAPI source using the existing generic `Problem` schema, with no internal reason or `Retry-After`; regenerate the bundle and TypeScript types. In the internal API helper, map this documented 503 to the existing non-ambiguous request-failure result rather than `ambiguous`. The current UI already routes that result through its generic failure handling; no new UI branch or copy is needed. Add/adjust only the contract-facing fixture and focused consumer evidence needed to represent the POST response. Keep other 5xx/transport failures ambiguous and preserve their same-key/same-payload retry behavior.

**Why.** This carries OI9's known no-admission meaning into the API consumer while keeping unknown outcomes conservative. It stays within the contract/generated/internal consumer boundary and requires no new user-facing state. The existing generic failure wording remains non-specific; any more tailored user messaging would be a separate UX decision outside this Work Unit.

**Risk / consequence.** The generic failure path clears the stored retry intent. That is consistent with the 503 guarantee that no Donation was admitted; if the user deliberately retries after Funding is restored, a new request can use a fresh key. Tests/evidence must ensure other ambiguous 5xx cases continue to retain the original key and payload. No claim is made about live backend behavior.

### Option B — Keep all 5xx, including 503, classified as ambiguous

Add the response to authored and generated API artifacts and model it in a POST mock, but preserve the consumer's existing blanket `>=500` mapping and ambiguous retry behavior.

**Why it is viable.** It preserves the current conservative behavior and avoids changing consumer classification.

**Why it is not recommended.** The consumer would continue to present this explicit no-admission response as an unknown outcome. Retrying with the same key is safe, but the classification would not reflect the accepted contract's more precise meaning.

## Recommended direction

Proceed with **Option A** for Techplan synthesis, subject to reopening the live sources and exact consumer scope there. Keep the shared `Problem` component unchanged unless the authored operation's actual needs demonstrate a shared-definition gap. Do not add `Retry-After`, internal diagnostics, a new UI state/copy, backend behavior, or runtime claims. Preserve a distinct ambiguous path for transport failures and any other 5xx not guaranteed to mean no admission.

No new Product/domain/API policy decision is needed to choose this direction: the no-admission semantics are already accepted, and the existing consumer has a reusable generic failure path. The API owner's required acceptance of exact authored source/counterpart bytes remains a later completion gate, not an approval inferred by this recommendation.

## Material rejected direction

Option B is not selected because it leaves the consumer's result less precise than the settled contract. Its same-key retry safety is retained for genuinely ambiguous outcomes under Option A, so rejecting B does not weaken the current ambiguous-retry protection.

## Open / deferred

- Exact authored response wording/reference, generated output bytes, affected internal consumer/fixture scope, and verification belong to downstream Techplan/Build/Review/Testing as applicable.
- API owner source/counterpart acceptance and exact source/output hashes remain required by WU-S2-007.
- No runtime, external distribution, or delivery readiness is inferred.
- No additional human policy decision is requested now; exact API owner acceptance remains required after concrete artifacts exist.

## Phase handoff

- **Outcome:** `COMPLETED` — this Explorer execution occurrence completed Stage 1–3; WU-S2-007 remains active.
- **Result refs:** `runs/EXP-S2-007-001/evidence/stage-2-gap-analysis.md`; `runs/EXP-S2-007-001/evidence/stage-3-solutioning.md` (recommended Option A).
- **Findings:** F1 authored/generated Donation POST omission; F2 generic 5xx consumer classification; details and progression effect in the Stage-2 evidence.
- **Decision requests:** None now. API owner acceptance of exact source/counterpart bytes remains a required later gate.
- **Blockers:** None for Exploration completion or Techplan synthesis. WU-S2-007 completion and dependent WU-S2-003 progression remain pending contract/counterpart reconciliation and acceptance.
- **Open / unverified:** No authored/generated/consumer changes or checks occurred in this Run; exact output scope/hashes, independent review/verification, and owner acceptance remain unverified.
- **Recommended continuation:** Techplan synthesis in a new Techplan Run/Participant with a fresh Participant Session/context, reconstructed from these durable Exploration artifacts and current authority.
- **Context refs:** WU-S2-007 manifest; Donation invariant `INV-donation-02` / Resolved OI9; Donation feature funding-capacity criterion / Resolved OI9; `api/openapi/donation.yaml` POST operation; `frontend/lib/api/donation.ts` `submitDonation`; `frontend/app/campaigns/[campaignId]/donate/donation-client.tsx` `send`.
