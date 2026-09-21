# WU-S1-002 — Slice 1 Public Contract & Delivery Reconciliation

Type:
RECONCILIATION

Parent Outcome:
S1 — Public Campaign Understanding

Derived From:
WU-S1-001 / EXP-001

Status:
NOT_STARTED

Scheduling:
QUEUED

Horizon:
NOW

Readiness:
DEFINED

Coordination Owner Role:
Orchestration Operator

Primary Execution Role:
Planner

Specialization:
None

Communication Language:
Bahasa Indonesia

Communication Profile Path:
`docs/project/communication-profile.md`

## Outcome

Menghasilkan reconciled Slice-1 delivery contract yang cukup sempit, aman, dan current untuk menjadi dasar `CONTRACT_READY`, sehingga backend dan frontend dapat direncanakan/dijalankan tanpa menjadikan historical specs, mixed public/private schemas, atau frontend assumptions sebagai authority.

## Scope

- Reconcile Slice-1 acceptance criteria dari Product/MVP + Product Design authority.
- Reconcile public Campaign visibility semantics dan public-safe not-found-equivalent behavior.
- Define explicit public Campaign/steward/funding/media/provenance/lifecycle/action projection boundary.
- Reconcile organizer content safety semantics dan explicit unknown/pending meaning.
- Reconcile relevant Campaign/Organization invariants dan threat model concerns.
- Reconcile touched OpenAPI operation/schema/error behavior untuk public Campaign Detail.
- Define the narrow generated-contract path needed by frontend.
- Reconcile integration-map/tracker state needed to represent the active Slice-1 delivery contract truthfully.
- Preserve future compatibility with Slice 2/3 without pulling their delivery scope forward.

## Out of Scope

- Production backend implementation.
- Production frontend implementation.
- Full Campaign/Organization lifecycle cleanup.
- Donation Flow.
- Campaign closure/result/accountability delivery.
- Full Organization public profile.
- Upload/self-service/curation/scheduler capabilities.
- Broad repository-wide OpenAPI warning cleanup.
- Final implementation mechanics that belong to downstream Build planning after the reconciliation contract is approved.

## Required Prior Artifacts

- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001/evidence/gap-analysis.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001/evidence/solutioning.md`
- current Product/MVP and Product Design authority referenced by those artifacts.

## Key Settled Boundaries From Exploration

- Public detail is explicit public-only projection, not inherited internal schema.
- Non-public and absent Campaign behavior must be public-safe and indistinguishable.
- Story remains plain text for Slice 1.
- Funding/action semantics are backend-authoritative.
- No active Donate action in Slice 1.
- Campaign media requires private storage + controlled public delivery.
- Frontend may proceed contract-parallel only after reconciliation reaches `CONTRACT_READY`.
- Exact schema/property names, cache mechanics, seed mechanics, and implementation details remain planning/reconciliation work within these boundaries.

## Human Gates

Human Authority is not currently required for Product scope.

Human Design review remains required before final production wording or any precedent-setting reusable placeholder asset/system is promoted.

If reconciliation introduces a public Organization review/verification claim whose meaning is not already owned by current Product/Design authority, stop and raise an `Authority Decision`.

## Completion

This Work Unit is complete when the touched Slice-1 delivery/spec/threat/contract surfaces are reconciled sufficiently to support a truthful `CONTRACT_READY` milestone and downstream backend/frontend planning without unresolved material contract ambiguity.

## Planning State

Draft Techplan:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`

Synthesis Run:
`TP-001`

Synthesis State:
COMPLETED

Techplan Status:
DRAFT

Independent Review:
RECOMMENDED_AND_GATE_APPLIES

## Planning State

Draft Techplan:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`

Synthesis Run:
`TP-001` — COMPLETED

Independent Review Run:
`TPR-001` — COMPLETED

Review Result:
1 MATERIAL / BLOCKING security-interface finding: public exact-allowlist schemas do not yet explicitly close additional properties.

## Planning State

Draft Techplan:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`

Synthesis Run:
`TP-001` — COMPLETED

Independent Review Run:
`TPR-001` — COMPLETED

Resolution Run:
`TPR-RES-001` — COMPLETED

Review Outcome:
Blocking closed-object finding resolved. Amendment makes the already-settled exact-allowlist semantics executable and does not materially change interface/security meaning.

Techplan Status:
DRAFT — AWAITING_HUMAN_APPROVAL

Re-review:
NOT_REQUIRED

## Current Gate

Human Techplan approval

## Next Action

Human reviews and either approves or requests revision of the amended Techplan. Do not start Build before approval.