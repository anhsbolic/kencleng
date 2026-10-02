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

- Execution status: `WAITING`
- Scheduling state: `PARKED`
- Horizon: `NOW`
- Completed Run / current Draft: `TP-S2-004-001` selesai synthesis dan self-check, berstatus Draft / In Review. `runs/TP-S2-004-001/techplan.md` dan `handoff.md` adalah current evidence; belum ada independent Techplan Review/report/approval.
- Current milestone: None
- Next route: WU006 / TP006004 mechanically verified Draft; RV006001 no blockers, TP006005 complete report; Human approved TP006004; TP006006 Status-only verified; BLD006001 Product checkpoint selesai; RV006002 Approve/no findings; Product/MVP accepted, BLD006002 six-file spec checkpoint selesai; RV006004 Approve, C-01/Q-01 resolved; six-spec accepted, BLD006004 metadata verified; API transport/encoding settled, TP006007 completed successor Draft; RV006005 no blockers; TP006008 mechanical delta/full report verified, Human TP006008 approved; TP006009 Status propagation siap; OI-2 reason direction diterima, material contract/source acceptance masih pending. Setelah accepted source delta tersedia, fresh frontend Planner revision, lalu recommended independent Techplan Review, report dan Human approval sesuai actual convergence. Review belum diinvoke, tidak dinyatakan NOT_APPLICABLE. Decomposition Consider hanya setelah reviewed/reconciled spine dan approval gate; belum ada split yang diterima.
- Dependency/readiness: WU-S2-002 DONE/CONTRACT_READY dan WU-S2-005 DONE/final authored Campaign/API acceptance satisfied; topology and conditions owned by parent Work Graph.
- Scoped blocker status: Exploration F-1 Campaign action contract gap resolved by completed/accepted WU-S2-005. Initial Draft selesai; OI-1 secara eksplisit menahan whole-Techplan approval dan amount-limit-dependent Build sampai owning monetary sources WU-S2-006 reconciled/accepted. Scope dependency dimiliki parent Work Graph. Tidak mengadopsi candidate cap/field/reason dari backend Draft sebagai accepted API. Optional email F-2 and O3/O4/O5 runtime evidence remain scoped downstream; mocks do not establish those claims.
- Human gate: Campaign/API final acceptance settled, no schema/route re-vote. Frontend whole-Techplan approval menunggu OI-1/source convergence dan applicable review/report. Material rendered acceptance tetap later gate. Anhar now owns Product/MVP monetary/closure policy for the named Slice-2 concern; attribution dan accepted EXP006 D-01–D-04 solutioning tidak menjadi concrete source/plan/risk acceptance.
- Evidence: `runs/EXP-S2-004-001/evidence/stage-2-gap-analysis.md`, `stage-3-solutioning.md`; completed `runs/TP-S2-004-001/techplan.md` dan `handoff.md`; R1–R12 memiliki Testing Checklist coverage dan relevant Test Focus pointers; Participant melaporkan no checks executed; WU005 accepted feature/API and Review/Testing/status evidence; backend TP003003 handoff/D16/O1-REP for current coordination delta.

## Current-effective inputs

- Parent Outcome and Work Graph in `../outcome.md` and `../work-graph.md`.
- Approved Techplan `../WU-S2-002/runs/TP-S2-002-015/techplan.md`.
- Accepted Campaign baseline: `../WU-S2-005/manifest.md`, current Campaign feature/INV-campaign-14 and `api/openapi/campaign.yaml` with generated counterparts.
- Backend coordination: `../WU-S2-003/runs/TP-S2-003-003/handoff.md`, `launch-record.md`, exact D16/O1-REP in its Draft `techplan.md`; no candidate policy becomes accepted frontend API implicitly.
- Accepted Donation specs `docs/spec/5-donation/` and accepted API sources `api/openapi/donation.yaml`, `api/openapi/common.yaml`.
- `docs/product/mvp-delivery-slices.md` Slice 2; `docs/ui-ux/README.md` and `docs/ui-ux/page-map.md`.
- `docs/project/kencleng-frontend-tech-stack.md`, `frontend/AGENTS.md`, and applicable Harscode guidance are current-effective for scoped phases.
