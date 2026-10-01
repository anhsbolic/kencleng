# WU-S2-004 — Slice 2 Guest Donation Frontend Flow

## Definition

- Type: `DELIVERY`
- Parent Outcome: `S2-GUEST-DONATION-TRUTHFUL-STATE`
- Derived from: `WU-S2-002` / `CONTRACT_READY`
- Coordination owner role: Orchestration Operator
- Planned first phase role: Explorer
- Specialization: None established
- Communication language: Bahasa Indonesia
- Communication profile: `docs/project/communication-profile.md`

### Outcome

Deliver the coherent guest Donation experience from an eligible public Campaign into submission and truthful status/tracking, using the accepted API contract and contract-faithful MSW mocks until real backend integration.

### Scope

- Frontend production work only; scope page/flow surfaces from canonical page-map and approved Slice 2/product/design authorities.
- Use the accepted OpenAPI/generated types as the interface and MSW (or approved equivalent) at the network boundary for mock verification.
- Preserve source-first sandbox/status wording, accessible and responsive interaction, optional donor fields, safe fragment credential handoff/URL cleanup, and truthful pending/terminal/error states as owned by current Product/Design/spec/API authority.
- Identify active design/component readiness and required frontend verification during Exploration/Techplan.

### Boundaries

- Do not modify `backend/` or alter shared contract/spec/Product authority in a frontend Build.
- Do not implement a mock-mode branch in production data/service code or imply real external settlement.
- Mock verification does not establish backend runtime/security controls or real integration.
- O1 currency-specific numeric parameters and O2–O5 runtime/security evidence remain deferred/downstream.

### Completion condition

The scoped guest Donation page/flow meets the accepted requirements against contract-faithful mocks, scoped frontend verification and material rendered Human acceptance are complete, and the tracker records `FRONTEND_MOCK_VERIFIED`. Real backend integration and Slice finalization remain separate.

## Current State

- Execution status: `ACTIVE`
- Scheduling state: `QUEUED`
- Horizon: `NOW`
- Current Run: `EXP-S2-004-001` prepared for canonical Exploration Stage 1; no Participant dispatched.
- Current milestone: None
- Human gate: Canonical Exploration Stage 1 must stop for confirmation before Stage 2. Material open Product/Design decisions and Human rendered acceptance remain Human-owned.

## Current-effective inputs

- Parent Outcome and Work Graph in `../outcome.md` and `../work-graph.md`.
- Approved Techplan `../WU-S2-002/runs/TP-S2-002-015/techplan.md`.
- Accepted Donation specs `docs/spec/5-donation/` and accepted API sources `api/openapi/donation.yaml`, `api/openapi/common.yaml`.
- `docs/product/mvp-delivery-slices.md` Slice 2; `docs/ui-ux/README.md` and `docs/ui-ux/page-map.md`.
- `docs/project/kencleng-frontend-tech-stack.md`, `frontend/AGENTS.md`, and applicable Harscode guidance are current-effective for scoped phases.
