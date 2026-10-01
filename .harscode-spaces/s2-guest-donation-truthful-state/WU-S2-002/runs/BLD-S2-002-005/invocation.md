# Run Invocation — `BLD-S2-002-005`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `BLD-S2-002-005`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-005`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Implementer
- `SPECIALIZATION`: Targeted Task 01 Build/Patch for Code Review finding F-001
- `PARTICIPANT_ID`: `P-S2-002-BL-005-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; current Profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Build/Patch Run after independent Code Review RV-S2-002-011 requested a bounded correction.
- `SESSION_TRANSITION_REASON`: Orchestrated patch re-entry after a completed Review Run; use a new Implementer Participant/Session and reconstruct only from current plan/task, exact finding/patch plan, Build handoff, and live source.
- `TARGET_REVISION`: Kencleng checkout `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus current working-tree spec diff and Run artifacts; inspect the exact live passages at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; re-read canonical Build guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Narrow Build/Patch re-entry requested by Code Review to resolve only blocking finding F-001. Return to the requesting Review phase for targeted confirmation; do not automatically start a new full review loop unless the patch broadens materially or reveals another material defect.
- `RISK_TIER`: Tier 1 — finding concerns a required security property of a bearer credential granting access to private Donation status. Patch is documentation-only and must preserve evidence/risk boundaries.
- `MODEL_ROUTING_RATIONALE`: One focused, four-file correction to restore an already-settled Product/MVP and TP-015 property. `gpt-6-luna` at `high` is the configured Implementer-capable minimum sufficient route; escalate only for demonstrated capability insufficiency, not for missing authority/context.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` — current-effective Human-approved `Approved` spine; use Q7 and linked R6/D7, O4, §12, and §13 as the material authority.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002/tasks/01-donation-domain-spec-reconciliation.md` — current Task 01 execution scope and security/status credential boundaries.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-011/review-findings-1.md` — blocking finding F-001 and its authority, impact, and requested resolution.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-011/patch-plan-1.md` — exact four-file correction. Follow this narrow scope; do not broaden it.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-004/report.md` and `launch-record.md` — immediately prior Build evidence; preserve settled content and all unchanged boundaries.
- Exact source targets from the patch plan:
  - `docs/spec/5-donation/invariants.md` — `INV-donation-05`.
  - `docs/spec/5-donation/features/02-donation-status-check.md` — reconciliation, active acceptance, and Credential guessing/theft threat row.
  - `docs/spec/5-donation/tasks.md` — Task 02 active acceptance.
  - `docs/spec/5-donation/threat-model.md` — temporary guest status URL / Spoofing row.
- Relevant authorities to re-open: root `AGENTS.md`; `docs/spec/README.md`; `docs/product/mvp-scope.md` §5 and `docs/product/mvp-delivery-slices.md` §5; TP-015 Q7; relevant approved O4 decision source already cited by TP-015. No raw Exploration is needed.
- Canonical current Harscode guidance: `workflow/3-build-prompt.md`, `workflow/3-build/guidelines.md`, `workflow/3-build/checklist.md`, `workflow/orchestrated-run-overlay.md`, `workflow/context-management.md`, `workflow/AGENTS.md`, `orchestration/run-contract.md`, `orchestration/AGENTS.md`.

## Patch assignment — F-001

Restore the already-settled Product/MVP and TP-015 requirement that the guest status bearer credential/token be **difficult to guess**. Carry that property into the active invariant, Feature 02 acceptance, Task 02 acceptance, and temporary-status-URL Spoofing threat row named above.

Distinguish the required property from concrete implementation evidence that remains open under O4: do not choose a numeric entropy/length target, token-generation algorithm, storage or comparison mechanism, expiry-enforcement mechanism, exposure/abuse control, or residual-risk decision. Do not claim strength/generation has been empirically verified. Keep hard 24-hour expiry, fragment handoff/URL cleanup, one-way HMAC verifier, status-only access, and uniform O5 `404` direction intact.

## Execution envelope

- `PREAUTHORIZED`: Read only the listed spine, task, finding/patch plan, prior Build handoff, relevant current authority, and four live source passages; edit only those four passage scopes; write this Run's `patch-report-1.md` and `launch-record.md`; reread affected passages against TP-015 Q7/Product §5 and run exact-path `git diff --check`.
- `ORCHESTRATOR_DECISION`: Any fifth source file, broader redesign, altered review route, or change beyond F-001.
- `HUMAN_REQUIRED`: New authority decision, numeric/security parameter selection, residual-risk acceptance, spec `agreed` status, Product/MVP/Techplan/OpenAPI/code/test/protected-path edit, or milestone acceptance.

Do not change Product/MVP, Techplan, OpenAPI, code, tests, migrations, spec status, task snapshots, projections, or files beyond the four named source targets and this Run's artifacts. Do not introduce a split, new decision, or new source of truth. Do not run automated tests or runtime/security checks; they are not needed for this documentation correction and remain outside this patch Run.

Write only `patch-report-1.md` and `launch-record.md` under this Run. Report exact changed passages and actual focused checks; explicitly state no tests/runtime/security verification ran. Preserve all five spec files as `draft`, the other settled requirements, and all O1/O2/O3/O4/O5 deferred obligations. Do not claim Human/domain-owner acceptance, residual-risk acceptance, implementation evidence, or `CONTRACT_READY`.

## Human-assisted dispatch

Use a fresh Implementer Session at the Kencleng repository root with the configured model/effort. Start from `../harscode-workspace/workflow/3-build-prompt.md` and this Invocation. Kickoff:

`Jalankan targeted Build/Patch Run BLD-S2-002-005 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-005/invocation.md dan canonical Build prompt Harscode saat ini. Reconstruct dari Approved TP-015 Q7, Task 01 snapshot TPD-002, finding F-001 dan patch plan RV-011, serta BLD-004 handoff. Buka ulang hanya authority dan empat passage source yang disebut. Pulihkan persyaratan credential status sulit ditebak pada invariant, Feature 02 active acceptance, Task 02 active acceptance, dan threat-model Spoofing row. Jangan memilih entropy/panjang, algoritme/generation detail, kontrol implementasi/evidence O4, atau menerima risk. Pertahankan 24 jam, fragment handoff/cleanup, one-way HMAC, status-only, uniform O5 404, seluruh deferred boundary dan draft status. Jangan mengubah file lain, API, code/test, status, projection, atau Product/Techplan authority. Jalankan reread traceability dan exact-path git diff --check saja; tanpa tests/runtime checks. Tulis patch-report-1.md dan launch-record.md, lalu kembalikan ke Review untuk targeted F-001 confirmation.`

After dispatch, reconcile the phase route from the Build evidence. If the patch exactly implements RV-011's narrow plan without broadening, prepare targeted Review confirmation in a fresh Reviewer Run; use full review only if material scope changes or a new concern requires it.
