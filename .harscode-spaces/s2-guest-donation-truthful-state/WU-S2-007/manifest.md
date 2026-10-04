# WU-S2-007 — Donation POST Funding-unavailable API Reconciliation

## Definition

- Type: `RECONCILIATION`
- Parent Outcome: `S2-GUEST-DONATION-TRUTHFUL-STATE`
- Coordination owner: Orchestration Operator
- Planned first phase role: Explorer
- Communication language: Bahasa Indonesia

### Outcome

Reconcile the accepted Donation POST behavior for unavailable authoritative settled Funding into the authored OpenAPI source and required generated/internal consumer counterparts, with exact source/counterpart evidence and API owner acceptance.

### Scope

- Add the Anhar-approved generic `503 Problem` response to `POST /campaigns/{campaignId}/donations` when authoritative settled Funding is unavailable.
- Preserve fail-closed behavior: no new Donation is admitted; no internal reason or unproven `Retry-After` is disclosed.
- Reconcile only affected authored API, generated bundle/types, fixtures, and internal/repository-development consumers needed for this response.
- Record exact current source/output hashes, independent Review/verification results as applicable, owner source acceptance, and downstream WU003 handoff.

### Boundaries

- Donation invariant and feature spec are the accepted policy authority for this Run. No new Product/domain/privacy policy is delegated.
- Do not alter WU-S2-006 accepted source revisions or broaden historical O8.
- Do not implement HMAC/idempotency persistence, backend/runtime behavior, frontend UX, migrations, or protected crypto paths.
- Do not change shared API components unless actual impact requires a shared definition; prefer the local Donation response reference to the existing `Problem` schema.
- Current distribution posture remains internal/repository-development only; coordinated counterpart reconciliation remains required before delivery/runtime progression.

### Completion condition

The authored Donation POST response source and all affected generated/internal counterparts are independently checked and accepted by the API owner against exact bytes; WU003 receives an exact source/counterpart handoff. No runtime or delivery milestone is implied.

## Current State

- Execution status: `DONE`
- Scheduling state: `DONE` — BLD-S2-007-001, independent RV-S2-007-002, and independent TST-S2-007-001 are complete; Anhar accepted the exact five reviewed source/counterpart revisions on 2026-10-04. The exact WU003 source/counterpart handoff is recorded in `handoff-to-WU-S2-003.md`, SHA-256 `cf468693dc384ccad5a4156c43d433a3ff725a6169869c3f171aeff0a27325eb`. Review verdict is `Approve with minor comments`; Testing verdict is `Pass with flagged follow-ups`. Their non-blocking follow-ups remain visible and no patch plan was required.
- Current milestone: None. No runtime or delivery milestone is implied.
- OI9 owner decision: resolved by Anhar on 2026-10-04 as generic `503 Problem`, without internal reason or `Retry-After`; exact source and counterparts are accepted per `handoff-to-WU-S2-003.md`.
- Completed Exploration: `runs/EXP-S2-007-001/evidence/stage-2-gap-analysis.md`, SHA-256 `9d287461af2fa5e4ed210bba4ac95aee18f86a48d9644667049e016e696f791d`; `stage-3-solutioning.md`, SHA-256 `26bf0b4b33551dacf5b73133a4c3a5f0ec1b354b94950f23dc22380f5d4ea7ac`. Human reported completion; artifacts record configured `gpt-6-luna` / `medium`, no Session ID exposed. Stage 3 recommends mapping the contract-defined Donation POST 503 to the existing generic definitive-failure consumer path, leaving other 5xx/transport failures ambiguous with existing same-key retry. F1 is authored/generated response omission; F2 is consumer classification/mock/test omission. No source or checks changed in Exploration.
- Completed first Techplan Run: `runs/TP-S2-007-001/handoff.md`, SHA-256 `f19a6c462c60290b09d48f89999b917962d858ce9a4d5f9d25dfbf7eba945605`. It produced reviewed predecessor hash `a54a441787819bb291531cbdee1ce94d6754049d91a97d7304c4248cc04ddcd2`; Planner configured `gpt-6-luna` / `high`, Session not exposed. TP-S2-007-002 resolved the Review's exact mechanical `(D2)` correction without semantic change. Handoff recommends independent Techplan Review for cross-boundary API/generated/consumer semantics and Skip decomposition for this cohesive sequence.
- Completed independent Review: `runs/RV-S2-007-001/review-findings-1.md`, SHA-256 `70361a6188bd5ac89c9d9d1001f322e8eea6430f1ef8ba1b3708ac5ede2746f7`; `launch-record.md`, SHA-256 `e38b6dcaf72cdb5480af889775aebde4080cc9de5147d086d9de066db894c245`. Reviewer configured `gpt-6-luna` / `high`, Session not exposed. Verdict has no blocking findings and one mechanical/non-blocking stale D2 cross-reference in §12 Test Focus Pointer. Reviewer confirms no re-review is required for the exact one-reference correction. Human reported RV-S2-007-001 complete.
- Completed Planner resolution/report Run: `runs/TP-S2-007-002/handoff.md`, SHA-256 `9aa8456b319ddbcad9bea878d2346686dbac3a38fe4bbc7b90c9fa5e05eb8b91`. It confirmed the starting Techplan matched the reviewed hash `a54a441787819bb291531cbdee1ce94d6754049d91a97d7304c4248cc04ddcd2`; the only change was removal of `(D2)` from the final sentence of §12 and reinserting it reproduces the original hash. The exact Draft approved by Anhar was `d9b483180a20dd2b4df618654a13164333816e552d3cc87f2e9787aa891f6928`; report SHA-256 `e022e41909c5aa024b989e278cc9d752b0ae45b751deb039fa33da54411186e6`. Planner `gpt-6-luna` / `low`, Participant `P-S2-007-TP-002-1`, Session not exposed. No tests, validators, generators, source changes, API acceptance, or runtime checks ran.
- Human approval and status reconciliation: Anhar approved the exact Draft hash `d9b483180a20dd2b4df618654a13164333816e552d3cc87f2e9787aa891f6928` against report `e022e41909c5aa024b989e278cc9d752b0ae45b751deb039fa33da54411186e6`. Only the Techplan frontmatter Status changed to `Approved`; reversing only that marker reconstructs the exact approved Draft hash. Current Approved Techplan SHA-256 is `ff0cb704151495e9f5539d11c24e808e8618d1c98cdb0d680957df8d51948dc3`. No Participant Run was used for this deterministic reconciliation.
- Completed initial Build Run: `runs/BLD-S2-007-001/report.md`, SHA-256 `2677f8c39e0723c8a365823edb51ca7c0c35523a38e2db144d4f19732bc4d235`; Invocation SHA-256 `b02be1e2f3022775c11280ec7a9778c5a6f6c4944e1ca8dae50b5879138cbc6e`. Participant `P-S2-007-BLD-001-1`, `KC-IMPLEMENTER`, configured `gpt-6-luna` / `medium`, Session ID not exposed. It changed only the authored Donation POST contract, generated aggregate/types, frontend submit helper, and focused flow test. Final changed-file hashes: `api/openapi/donation.yaml` `9f7c31065c3c7ffa77491541102afc0ce84085f50a1e2aec317336c9656d5c1d`; `api/openapi.yaml` `c37038ecb5088aebcc9156e495d8b14bfa7cc405debeb940e16d7b17786d1838`; `frontend/lib/api/generated/openapi.ts` `f267598c563b4bf73442f120c81fdf6ec1b7f48cd29095f78f3622dbd7094e20`; `frontend/lib/api/donation.ts` `5c3cdca3b71d806ece0c15bd1dbd8a0a162895ec0c56a4b63c3271f7fac6d654`; `frontend/app/donations/donation-flow.test.tsx` `e7b66525258a8601746583274a7ae9d794f0795254aba794e29a46f153fb6a38`. Participant reports bundle, API validation (124 warnings, zero errors), generated type command, focused flow test (1 file/16 tests), and `git diff --check` passed. Orchestrator independently matched all five final hashes and confirmed the diff stays within plan scope; checks were not rerun. No broad Testing suite, runtime, concurrency/performance/security-class verification, or source acceptance is claimed. Next phase is fresh independent four-pass Code Review.
- Completed independent Code Review Run `RV-S2-007-002`, fresh Reviewer / `KC-REVIEWER`, configured `gpt-6-luna` / `high`. Findings SHA-256 `23931fe93aa41ca8c8bec067173860f03da7c5f883d4cb67ba9985f29e6a4433`; handoff SHA-256 `8311f0f00bc9ca25b14ab976d568eda865faddeec2e2bc21648cb614069f4232`. Verdict: `Approve with minor comments`; BP-1 notes that the generic Problem fixture in `frontend/app/donations/donation-flow.test.tsx:143–147` is not tied to the generated `Problem` type. It is non-blocking and does not require a patch plan. Reviewer matched all five pinned Build hashes and the three read-only anchors. No test or runtime check was rerun during Review. Human reported RV-S2-007-002 complete.
- Completed fresh independent Testing Run `TST-S2-007-001`, Tester `P-S2-007-TST-001-1` / `KC-TESTER`, configured `gpt-6-luna` / `medium`; Invocation SHA-256 `94bca1166618d019a47b25ddd6e5a7e18d7cdfab71411b92fef3166677635347`, report SHA-256 `466167ba389d379f1b5c53ca35a47ffebe1dc808dc7f240cf2ed57f794029adb`. Verdict `Pass with flagged follow-ups`. `cd api && npm run validate` passed with 124 warnings / 0 errors and no warning-count delta from Build; `cd frontend && npm run verify` passed (5 files / 32 tests), with ESLint 0 errors and one existing unused-disable warning. The 503 observable flow, ambiguous transport retry key/payload, and existing 201/409/422 cases passed; a non-503 HTTP 5xx case was classified by branch inspection but has no dedicated UI assertion and was not considered blocking under the approved evidence plan. BP-1 remains a non-blocking generated-Problem fixture typing follow-up. No generator was run and no pinned file changed. Human reported exact-byte acceptance after final hash recheck. No backend/runtime acceptance or claim was made.
- Human/API-owner acceptance and WU003 handoff: On 2026-10-04 Anhar accepted exactly the five current revisions enumerated in `handoff-to-WU-S2-003.md`. The WU-S2-007 completion condition is met; the WU003 HARD source/counterpart dependency for candidate refresh/review/approval and affected Build is satisfied. WU003's next action is a fresh Planner occurrence to refresh the stale candidate, not Build. Separate candidate Review/Human approval and positive migration-design Review remain required before schema Build.
- Dependencies: WU-S2-002 Donation contract baseline is DONE/CONTRACT_READY; WU-S2-006 source/counterpart reconciliation is DONE. WU-S2-003's candidate refresh/review/approval and affected backend Build have a HARD dependency on this Work Unit's completion.

## Current-effective inputs

- `docs/spec/5-donation/invariants.md` and `docs/spec/5-donation/features/01-submit-donation-settlement.md` — accepted OI9 policy.
- `api/openapi/donation.yaml` plus referenced `api/openapi/common.yaml` — current authored baseline.
- `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts` — generated counterpart baseline.
- `api/README.md`, root `AGENTS.md`, relevant `frontend/AGENTS.md`, Work Graph, current WU003 manifest, and current Harscode phase prompts.
