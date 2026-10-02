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
- Scheduling state: `PARKED` — completed Planner occurrence; source reconciliation berjalan melalui separate WU-S2-006; tidak ada backend Build dispatch
- Horizon: `NEXT`
- Completed Run / current Draft: `TP-S2-003-003`, successor to TP-S2-003-002, discovered from durable techplan/handoff/launch-record during Campaign completion. No whole-Techplan approval or successor Review exists.
- Current milestone: None
- Findings: RV-S2-003-001 retry-after-close has Draft D15 resolution: under Campaign-first lock, resolve existing matching intent before new-key eligibility rejection; payload conflict still rejected. Independent closure pending. Monetary representability has recorded candidate D16/O1-REP direction: per-Campaign configurable individual cap, Rp1,000,000,000 direction; cumulative whole-IDR Campaign capacity ceiling; reserve accepted-pending obligations; close-at-cap with proposed `funding_capacity_reached`. Current D-01–D-04 policy/configuration/guest-disclosure directions sudah accepted/clarified lewat EXP006 handoff dan exact Human receipt di parent events; concrete owning Product/Campaign/Donation/API sources, wire shape/compatibility serta authored enum/source acceptance belum reconciled; OI-2 identifier funding_capacity_reached kini diterima owner sesuai current parent event.
- Next coordination: WU006 / TP006004 mechanically verified Draft; RV006001 no blockers, TP006005 complete report; Human approved TP006004; TP006006 Status-only verified; BLD006001 Product checkpoint selesai; RV006002 Approve/no findings; Product/MVP accepted, BLD006002 six-file spec checkpoint selesai; RV006004 Approve, C-01/Q-01 resolved; six-spec accepted, BLD006004 metadata verified; API transport/encoding settled, TP006007 completed successor Draft; RV006005 no blockers; TP006008 mechanical delta/full report verified, Human TP006008 approved; TP006009 Status propagation siap; OI-2 reason direction diterima, exact contract/source gates belum converged. Sesudah owning sources reconciled/accepted, fresh backend Planner refresh/independent re-review/report/whole-Techplan approval. Preserve valid choices tanpa repeat monetary vote; accepted solutioning bukan concrete source/plan/risk acceptance atau protected Build authorization.
- Authority: Anhar explicitly attributed Product/MVP authority for this Slice-2 monetary/closure policy on 2026-10-01. Monetary/Campaign/Donation/API ownership retained. Predecessor TP003003 provenance dipertahankan; current D-01–D-04 receipt pada EXP006 handoff + Human clarification di parent events; concrete source acceptance tetap terpisah.
- Scoped contract dependency: WU-S2-005 DONE and final authored acceptance now satisfy Campaign GET action contract dependency. Producer same-predicate GET/POST fidelity and runtime projection/cache/auth/error/recheck/D1 proof remain required; contract completion does not implement producer.
- Protected/runtime gates: exact Tier-0 authorization for `backend/internal/domain/campaign/donation_coordinator_db.go` and `backend/internal/domain/donation/ledger.go` absent. O3 controls/provider/retention/recovery, O4 key/exposure/abuse, O5 proxy/parity evidence and residual-risk gates remain Active. No BACKEND_VERIFIED or runtime proof.

## Current-effective inputs

- Source-reconciliation route: `../WU-S2-006/manifest.md`; dependencies dimiliki parent Work Graph.
- Current successor: `runs/TP-S2-003-003/techplan.md`, `handoff.md`, `launch-record.md`; scoped Product/MVP attribution in Authority Map/parent event.
- Completed Review: `runs/RV-S2-003-001/review-findings.md`, `handoff.md`, `launch-record.md`; current Draft predecessor: `runs/TP-S2-003-002/techplan.md`; completed handoff/provenance `runs/TP-S2-003-002/launch-record.md`; TP-S2-003-001 sebagai predecessor history. Recorded owner decisions dan batasnya pada §§5/13 dan parent completion event.
- Parent Outcome and Work Graph in `../outcome.md` and `../work-graph.md`.
- Approved Techplan `../WU-S2-002/runs/TP-S2-002-015/techplan.md`.
- Accepted Donation specs `docs/spec/5-donation/` and `docs/spec/4-campaign/` only where D1/threshold interactions apply.
- Accepted authored API sources `api/openapi/donation.yaml` and referenced `api/openapi/common.yaml`; generated aggregate/types are correspondence evidence.
- `docs/product/mvp-delivery-slices.md` Slice 2 and `docs/project/kencleng-monetary-data-standard.md`.
- `backend/AGENTS.md`, backend architecture, root `AGENTS.md`, and applicable Harscode guidance are current-effective for scoped phases.
