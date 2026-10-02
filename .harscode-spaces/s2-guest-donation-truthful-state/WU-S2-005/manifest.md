# WU-S2-005 — Slice 2 Campaign Donation Entry Contract Reconciliation

## Definition

- Type: `RECONCILIATION`
- Parent Outcome: `S2-GUEST-DONATION-TRUTHFUL-STATE`
- Derived from: `WU-S2-004` / `EXP-S2-004-001` Stage-3 solutioning and Anhar's Option 1 decision (2026-10-01)
- Coordination owner role: Orchestration Operator
- Planned first phase role: Explorer
- Specialization: None established
- Communication language: Bahasa Indonesia
- Communication profile: `docs/project/communication-profile.md`

### Outcome

Reconcile the public Campaign Detail donation-action contract and its owning Campaign acceptance criteria with the approved Slice 2 requirement that a visitor can enter guest Donation from an eligible Public Campaign Detail. Preserve backend authority over actual submission eligibility and produce a reviewed, owner-accepted contract baseline for WU-S2-004 planning.

### Scope

- Current Campaign Detail projection and its `donation_action` behavior in the Campaign feature spec, invariants where applicable, authored Campaign OpenAPI, and generated contract counterparts.
- Keep product meaning anchored in approved `docs/product/mvp-delivery-slices.md` §5 and current `docs/product/mvp-scope.md`; route any genuinely missing Product decision instead of inventing one.
- Define the public action response shape only through the current Campaign/API owners and applicable review/acceptance gates.
- Preserve public-safe projection, backend-authored eligibility, and the rule that a stale detail response does not authorize Donation submission; backend submission remains authoritative and may reject a now-ineligible Campaign.

### Boundaries

- This is a separate shared Campaign/API reconciliation. Do not expand WU-S2-004; it remains frontend-scoped.
- Do not implement backend or frontend production behavior, or run implementation/testing work in Exploration.
- Do not change Product/MVP authority, broader Slice 3 closure/result behavior, Donation submission eligibility invariants, or unrelated Campaign API operations.
- Do not let frontend visibility become an authorization/eligibility check. Coordinate any cross-stack contract change through this Work Unit and preserve separate backend/frontend Build scopes.

### Completion condition

Current Campaign Detail acceptance and authored API describe a Slice-2-compatible public donation action under the named Campaign/API authority; required generated counterparts and applicable independent review/owner acceptance are reconciled; unresolved decisions/evidence and exact readiness for WU-S2-004 Techplan are recorded. No implementation milestone or runtime behavior is implied.

## Current State

- Execution status: `DONE`
- Scheduling state: Not applicable — terminal Work Unit
- Horizon: `NOW`
- Completed Run: `BLD-S2-005-002`; Status-only propagation verified by Orchestrator against the accepted snapshot. Feature hash `55d0ede37e5fe70e07f5e2f2832f3bd03ce94a929e3995538fece7e681f62349`; normalizing Status gives accepted hash `b79a1dc891b614ff8750af988f3e78491881cdd9b6ed2e808ae766180ecc76e6`. Five counterparts unchanged.
- Current-effective Techplan: `TP-S2-005-002`, Approved. Completed independent Techplan Review RV-S2-005-001; report TP-S2-005-003; verified approval metadata TP-S2-005-004; authored/generated/fixture Build BLD-S2-005-001; Code Review RV-S2-005-002 Approve; independent Testing TST-S2-005-001 Pass with flagged follow-ups.
- Human acceptance: Anhar accepted final authored current Slice-2 Campaign/API action contract on 2026-10-01; exact answer and snapshot in parent events. `available` has no reason; `unavailable` has only `campaign_not_eligible`, used only if detail remains public but authoritative submission eligibility fails at GET. GET snapshot, POST recheck, unchanged non-public/closed 404 and dependency failure 503. Feature Status now records accepted contract / runtime delivery pending. No acceptance of all historical Campaign behavior.
- Completion/readiness: Work Unit completion condition satisfied at contract/spec/generated/fixture boundary; accepted baseline ready for WU-S2-004 Techplan and scoped Campaign GET producer planning. Work Graph records satisfied contract dependencies. No runtime milestone or additional `CONTRACT_READY` milestone invented.
- Phase applicability after BLD-S2-005-002: Code Review `NOT_APPLICABLE`, then Testing `NOT_APPLICABLE` for the verified Status-only delta. Every other feature byte and all five counterparts unchanged; no executable, behavior, schema, criterion, ownership or risk change. Existing independent Review/Testing evidence stays current; no skipped-phase Runs.
- Open / deferred handoff: Campaign producer/WU-S2-003 must establish same authoritative GET/POST eligibility predicate and runtime public projection/cache/auth/error/recheck/D1 evidence. No fabricated unavailable scenario, runtime proof, protected-write permission or residual-risk acceptance. TST observed current/HEAD 124 warnings with no rule/pointer delta; Build reported 122 remains a non-blocking historical discrepancy.
- Evidence: `runs/BLD-S2-005-002/report.md` and `handoff.md`; `runs/TST-S2-005-001/testing-report-1.md`, `handoff.md`, `launch-record.md`; `runs/RV-S2-005-002/review-findings-1.md`; prior named plan/Build/Review/approval evidence; parent completion/acceptance events.

## Current-effective inputs

- Parent Outcome, Work Graph, and Control Surface in `../outcome.md`, `../work-graph.md`, and `../control-surface.md`.
- Approved Slice 2 product authority: `docs/product/mvp-scope.md` and `docs/product/mvp-delivery-slices.md` §5.
- Campaign authority: `docs/spec/4-campaign/invariants.md` and `docs/spec/4-campaign/features/02-campaign-detail-listing.md`.
- Current authored Campaign API: `api/openapi/campaign.yaml`; inspect aggregate/generated counterparts only as needed for correspondence.
- Accepted Donation context: `docs/spec/5-donation/` and `api/openapi/donation.yaml` plus referenced `common.yaml`.
- Source finding and recommendation: `../WU-S2-004/runs/EXP-S2-004-001/evidence/stage-3-solutioning.md`.
- Ownership: `.harscode-spaces/authority-map.md` maps current Slice-2 Campaign delivery/domain and API/contract decisions to Anhar Solehudin.
- Root `AGENTS.md`, `docs/kencleng-agentic-workflow.md`, and current Harscode guidance.
