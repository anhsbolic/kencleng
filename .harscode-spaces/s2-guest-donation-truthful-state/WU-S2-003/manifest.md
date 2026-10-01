# WU-S2-003 — Slice 2 Donation Backend Delivery

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

Implement the backend-owned guest Donation capability required by the accepted Slice 2 domain specs and API contract, with truthful persisted sandbox state, safe guest status access, and exact-once successful funding reflection.

### Scope

- Backend production work only; own Donation submission/status, simulator-owned lifecycle, guest credentials and notifications as required by the accepted contract, and Campaign integration required by accepted Donation rules.
- Preserve the accepted OpenAPI contract and current Product/MVP, Donation-spec, invariant, threat-model, Techplan, and monetary-standard authority.
- Identify the concrete backend execution and verification scope during Exploration/Techplan, including O2–O5 obligations and applicable risk tier.

### Boundaries

- Do not change `frontend/` or shared API/spec/Product authority from a backend Build.
- Root `AGENTS.md` fences Donation ledger transaction/locking logic and the listed Tier-0 crypto/auth paths. No write to a protected path is authorized by this Work Unit; surface a required protected change for its Human gate.
- Do not introduce real payment settlement or treat simulation as external evidence.
- O1 currency-specific precision/range/fraction/storage details remain deferred unless later authority supplies them.

### Completion condition

Required backend behavior is implemented against the accepted specs/contract, independent applicable Testing establishes `BACKEND_VERIFIED` including the delivery-owned O2–O5 evidence, all protected-path authorization boundaries are respected, and remaining findings/limitations are explicit. This Work Unit does not itself earn integrated or Slice-finalization status.

## Current State

- Execution status: `ACTIVE`
- Scheduling state: `QUEUED`
- Horizon: `NOW`
- Current Run: `EXP-S2-003-001` prepared for canonical Exploration Stage 1; no Participant dispatched.
- Current milestone: None
- Human gate: Canonical Exploration Stage 1 must stop for confirmation before Stage 2. Any protected Tier-0 write, material Product/API/security/risk decision, or residual-risk acceptance remains Human-required.

## Current-effective inputs

- Parent Outcome and Work Graph in `../outcome.md` and `../work-graph.md`.
- Approved Techplan `../WU-S2-002/runs/TP-S2-002-015/techplan.md`.
- Accepted Donation specs `docs/spec/5-donation/` and `docs/spec/4-campaign/` only where D1/threshold interactions apply.
- Accepted authored API sources `api/openapi/donation.yaml` and referenced `api/openapi/common.yaml`; generated aggregate/types are correspondence evidence.
- `docs/product/mvp-delivery-slices.md` Slice 2 and `docs/project/kencleng-monetary-data-standard.md`.
- `backend/AGENTS.md`, backend architecture, root `AGENTS.md`, and applicable Harscode guidance are current-effective for scoped phases.
