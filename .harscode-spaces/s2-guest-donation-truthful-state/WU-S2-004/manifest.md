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

- Execution status: `DONE`
- Scheduling state: `DONE` — BLD-S2-004-001/002, RV-S2-004-003, and TST-S2-004-001 are complete. Testing returned `Pass with flagged follow-ups`: exact pinned 19-file set verified; Participant reports frontend verify (5 files / 31 tests), production build, and one Chromium R8 test passed; no code finding. Anhar reported the manual rendered R10 review accepted. R3 exact-boundary, R5 lifecycle/duplicate, and R8 wrong/expired/absent assertion gaps remain tracked as non-blocking verification follow-ups. Integration Map Donation-flow row is recorded. WU004 meets its mock-verification completion condition; real backend integration/runtime/security remain separate.
- Horizon: `NOW`
- Current-effective Approved Techplan: `TP-S2-004-003`, SHA-256 `e6a93bc65a760084955f63a469b238fd70a80c20ff38cc1d22ba1be000ad99eb`; exact pre-approval Draft SHA-256 `187f1f1b827eab39dbed7e5c46c3e62727fffaddc2b942991602184048171c03`. Anhar approved after report TP-S2-004-004 (SHA-256 `ada63615525ab061ceec4ada72414e07fbc8d13e6636a1d91203a19f6b9fbac2`); Orchestrator verified the preimage and changed only the Status marker. Resolution handoff SHA-256 `d82c760d7f2f73cb8d4f8941e5670a3148876dcf7cd291b57e71f05d9f71e112`. RV-S2-004-001 remains evidence on predecessor TP004002; TP004003 resolves its non-blocking F01 without material change. TP-S2-004-002 and TP-S2-004-001 remain preserved history.
- Current milestone: `FRONTEND_MOCK_VERIFIED` — scoped to contract-faithful mocks and the completed WU004 frontend evidence; it does not claim real integration or runtime/security readiness.
- Next route: WU-S2-003 backend delivery planning. Human decision remains on TP-S2-003-006 Open Item 7: name the Product/Security/data authority and governance for setting/updating Organization `verified` / `has_overdue_report`, or carry it as a scoped gate blocking those source updates and affected Campaign create/PATCH handlers while proceeding with the whole-plan approval path. No WU004 Participant is pending.
- Dependency/readiness: WU-S2-002 DONE/CONTRACT_READY dan WU-S2-005 DONE/final authored Campaign/API acceptance satisfied; topology and conditions owned by parent Work Graph.
- Scoped blocker status: WU004 has no blocking implementation, Review, Testing, or rendered-acceptance gate. R3/R5/R8 assertion gaps remain explicitly flagged for appropriate test maintenance/re-entry; they are not code findings or silently asserted coverage. Real integration, backend exact-wire/runtime/security, and O1–O5 evidence remain downstream.
- Human gate: Anhar reported the manual browser/rendered check completed and accepted after the R10 checklist. No specific browser session, viewport measurements, or screenshot artifact were supplied; no more detailed visual provenance is inferred.
- Approval receipt: “aku setuju dengan techplan nya bro. lanjut”; parent Events records the exact report/source hashes and status-only reconciliation proof.
- Evidence: `runs/EXP-S2-004-001/evidence/stage-2-gap-analysis.md`, `stage-3-solutioning.md`; current Approved Techplan `runs/TP-S2-004-003/techplan.md` and unchanged handoff; predecessor `runs/TP-S2-004-001/techplan.md`; WU006 accepted source/counterpart evidence and WU005 accepted feature/API and Review/Testing/status evidence. Completed resolution `runs/TP-S2-004-003/techplan.md` / `handoff.md`; Review `runs/RV-S2-004-001/review-findings.md`; report-only Run `runs/TP-S2-004-004/report-techplan.md` SHA-256 `ada63615525ab061ceec4ada72414e07fbc8d13e6636a1d91203a19f6b9fbac2` and `launch-record.md` SHA-256 `94cc4f3f32292330e39a269aab2c7b55b60c25dac4c4a65c847e713da2db3587`; Invocation SHA-256 `c31d2b0106a3852119b139f27f1027003d24c4e8625ebb5285e6e5cfb98307ae`.
- Build evidence: `runs/BLD-S2-004-001/report.md` SHA-256 `af721c79995f73647ba627895f7e3fdf726134e6cbe9049bfd4bb7b2059f4d0e`; launch record SHA-256 `28f057df7d987bba2edaa90977a14c94190e250479c676a0c34117dc20c4e46a`. RV-S2-004-002 findings SHA-256 `bfa58cace1bf1d2ceeed5e0f6325bdad50e96ec8114c2a2886a5a8be090c1a66`; patch plan SHA-256 `f42cf05d6118f2f2724faa53a7441d7f69e71431ab271101c924158b916a1b8c`. Its Review Invocation pinned the 19-file frontend set; the findings confirm all hashes matched. Patch report `runs/BLD-S2-004-002/patch-report-001.md` SHA-256 `ebc8c51c45d14a0d617fdaf42ff667e0bb6799d4fec3531d41751843fdaf0bfb`; launch record SHA-256 `352d44a045e983748c1e325d625d8c2ca9b7b2c0de0ae1f1c9828242f8321f09`. Participant reports focused flow tests passed (1 file/15 tests); Orchestrator did not rerun. RV-S2-004-003 confirmed both findings resolved; its report omits model/session provenance and no launch record was present. Prepared independent Testing Invocation SHA-256 `cebd22eb9c9fb57776fbe49d66a18102a7a0bbe8bc5c5864f72caf0b99ab771b`. Completed TST-S2-004-001 report SHA-256 `c04fc828a94784284dc92089d36e2b1947820d7ab3bd381f48f483d667dc8dbd`; launch record SHA-256 `b1f88fdaa37ac5ade23fb0c45efecd91d3720439e052ff3e1044678ced283e8a`. Participant reports verify 5 files/31 tests, build, and one Chromium R8 handoff test passed. Runtime model/session provenance was not independently exposed. Report flags R3/R5/R8 assertion gaps; no code finding. Integration Map Donation-flow row was reconciled. Human rendered R10 acceptance: Anhar reported manual browser review completed and accepted; no specific viewport/session/screenshot provenance was supplied.

## Current-effective inputs

- Parent Outcome and Work Graph in `../outcome.md` and `../work-graph.md`.
- Approved Techplan `../WU-S2-002/runs/TP-S2-002-015/techplan.md`.
- Accepted Campaign baseline: `../WU-S2-005/manifest.md`, current Campaign feature/INV-campaign-14 and `api/openapi/campaign.yaml` with generated counterparts.
- Backend coordination: `../WU-S2-003/runs/TP-S2-003-003/handoff.md`, `launch-record.md`, and Draft `techplan.md` remain planning history; the exact accepted current interface is owned by WU-S2-006 source snapshots in its manifest and the regenerated API/types. Do not use the backend Draft to override accepted source.
- Accepted Donation specs `docs/spec/5-donation/` and accepted API sources `api/openapi/donation.yaml`, `api/openapi/common.yaml`.
- `docs/product/mvp-delivery-slices.md` Slice 2; `docs/ui-ux/README.md` and `docs/ui-ux/page-map.md`.
- `docs/project/kencleng-frontend-tech-stack.md`, `frontend/AGENTS.md`, and applicable Harscode guidance are current-effective for scoped phases.
- Completed independent Review `runs/RV-S2-004-001/review-findings.md`, SHA-256 `e2b35bc0a0393edf79a785b3787fb59185ef080632754e3dce7023d71f01bd5c`; exact reviewed TP004002 Draft hash is recorded above. It contains no blocking findings and one non-blocking F01.
- Prepared Planner resolution `runs/TP-S2-004-003/invocation.md`, SHA-256 `f6113a827955a53581902fac5a020844a3aa57408b186fb70e4112ae223cb1ba`; not dispatched. It is limited to F01 and requires a materiality declaration under current Harscode guidance.
