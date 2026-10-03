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
- Own the backend Campaign draft create/PATCH write capability needed to configure the accepted per-Campaign donation cap, including its accepted default, omission-preservation, and publication-freeze behavior. Keep the added Campaign write scope to the accepted API/spec capability needed for this Slice; do not expand frontend or shared authority sources.
- Plan/own the minimum persisted representative-membership and Organization eligibility source/integration needed by the accepted Campaign draft write checks. Initial MVP data setup uses the approved seeded/operator-assisted posture; full Organization self-service and representative-management UI stay out of scope. Human/authority still must establish who may assert/update `verified` and `has_overdue_report` and the applicable governance before affected handlers can proceed.
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
- Scheduling state: `QUEUED` — TP-S2-003-006 remains Draft / In Review; RV-S2-003-003 completed cleanly against its exact hash with no findings. WU-S2-006 is DONE, clearing the Work Graph dependency for the report/Human Techplan gate. Open Item 7 continues to block only Organization eligibility fact updates and Campaign draft handlers pending named authority/governance. No Build authorization is inferred.
- Horizon: `NOW`
- Current Draft: latest completed Draft `TP-S2-003-006`, SHA-256 `d98b3d37e3c00962075428c3e14f87d5ee59dfba46b6ae36ce35f78dcebbce7a`, remains Draft / In Review and unapproved. Its handoff SHA-256 is `e62e4df0cfc9095d06ffe6e987afdca8f6ee369a2f7a16d7d1306ddaa2653bdd`. TP-S2-003-005 and earlier Drafts remain preserved history. TP-S2-003-006 Invocation SHA-256 is `11625aab5624bf39426563b2b74c8cc24c5590bbc50ebcee0673c020b049bb42`.
- Current milestone: None
- Findings: RV-S2-003-001 retry-after-close has Draft D15 resolution: under Campaign-first lock, resolve existing matching intent before new-key eligibility rejection; payload conflict still rejected. Independent Review RV-S2-003-003 completed cleanly with no findings; D15 planning fidelity is independently confirmed, with implementation/runtime proof still pending. Monetary representability has recorded D16/O1-REP direction: per-Campaign configurable individual cap, Rp1,000,000,000 direction; cumulative whole-IDR Campaign capacity ceiling; reserve accepted-pending obligations; close-at-cap with `funding_capacity_reached`. Product/Campaign/Donation/API source bytes were accepted at the exact seven-source snapshot recorded in WU006; generated/frontend counterparts are reconciled. Concrete backend DTO/mapping and exact-wire test remain WU003 delivery-planning/build work. No backend implementation/runtime, protected-write authorization, or risk acceptance is inferred.
- Next coordination: prepare the Planner-owned Human report/approval route for exact TP-S2-003-006 after resolving the Human authority/governance question recorded as Open Item 7, or explicitly carry that scoped gate into the Human approval decision if the authority route permits. Keep the accepted Product/API predicates unchanged. Exact DTO/mapping/wire-test work and Tier-0/O3/O4/O5 gates remain separately gated; WU004 report/Human approval can proceed independently.
- Authority: Anhar explicitly attributed Product/MVP authority for this Slice-2 monetary/closure policy on 2026-10-01. Monetary/Campaign/Donation/API ownership retained. Predecessor TP003003 provenance dipertahankan; current D-01–D-04 receipt pada EXP006 handoff + Human clarification di parent events; concrete source acceptance tetap terpisah.
- Scoped contract dependency: WU-S2-005 DONE and final authored acceptance now satisfy Campaign GET action contract dependency. Producer same-predicate GET/POST fidelity and runtime projection/cache/auth/error/recheck/D1 proof remain required; contract completion does not implement producer.
- Protected/runtime gates: exact Tier-0 authorization for `backend/internal/domain/campaign/donation_coordinator_db.go` and `backend/internal/domain/donation/ledger.go` absent. O3 controls/provider/retention/recovery, O4 key/exposure/abuse, O5 proxy/parity evidence and residual-risk gates remain Active. No BACKEND_VERIFIED or runtime proof.

## Current-effective inputs

- Source-reconciliation route: `../WU-S2-006/manifest.md`; dependencies dimiliki parent Work Graph.
- Latest completed Draft: `runs/TP-S2-003-006/techplan.md` and `handoff.md`; see exact hashes above. Invocation is `runs/TP-S2-003-006/invocation.md`.
- Completed independent Review: `runs/RV-S2-003-003/review-findings.md`, SHA-256 `b5cd4d47e714b8499ef0ae8d33062283763812d0a18cc7999f1f14d6f4e2a79d`; clean Complex gate, no Blocking or Non-blocking findings. It reviewed exact TP003006 Draft hash `d98b3d37e3c00962075428c3e14f87d5ee59dfba46b6ae36ce35f78dcebbce7a`. Invocation SHA-256 `6c72f0072cace32778d18398977ee90386dffdffde36be4f9bb079a346da984e`. Earlier package `RV-S2-003-002` remains superseded before dispatch because its target predates the material source/persistence route. Review does not approve the Techplan or authorize Build.
- Completed Review: `runs/RV-S2-003-001/review-findings.md`, `handoff.md`, `launch-record.md`; current Draft predecessor: `runs/TP-S2-003-002/techplan.md`; completed handoff/provenance `runs/TP-S2-003-002/launch-record.md`; TP-S2-003-001 sebagai predecessor history. Recorded owner decisions dan batasnya pada §§5/13 dan parent completion event.
- Parent Outcome and Work Graph in `../outcome.md` and `../work-graph.md`.
- Approved Techplan `../WU-S2-002/runs/TP-S2-002-015/techplan.md`.
- Accepted Donation specs `docs/spec/5-donation/` and `docs/spec/4-campaign/` only where D1/threshold interactions apply.
- Accepted authored API sources `api/openapi/donation.yaml` and referenced `api/openapi/common.yaml`; generated aggregate/types are correspondence evidence.
- `docs/product/mvp-delivery-slices.md` Slice 2 and `docs/project/kencleng-monetary-data-standard.md`.
- `backend/AGENTS.md`, backend architecture, root `AGENTS.md`, and applicable Harscode guidance are current-effective for scoped phases.
