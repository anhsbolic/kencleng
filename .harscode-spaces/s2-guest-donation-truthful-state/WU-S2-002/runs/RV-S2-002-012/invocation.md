# Run Invocation — `RV-S2-002-012`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `RV-S2-002-012`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-012`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Targeted independent confirmation of blocking finding F-001 from RV-S2-002-011
- `PARTICIPANT_ID`: `P-S2-002-RV-012-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; current Profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Reviewer Run after completed patch Run BLD-S2-002-005.
- `SESSION_TRANSITION_REASON`: Independent confirmation of the exact Review-requested patch using a new Reviewer Participant and fresh Session/context.
- `TARGET_REVISION`: Kencleng checkout `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus current BLD-004 source diff and BLD-005 narrow patch; reopen the four assigned passages at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; re-read current Code Review guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Narrow targeted Review re-entry requested by completed Code Review RV-S2-002-011 after Build/Patch implemented its exact patch plan. Do not repeat the full four-pass review unless actual scope materially broadened or the correction introduced a new material concern.
- `RISK_TIER`: Tier 1 — confirm preservation of the settled difficult-to-guess status-credential requirement without inventing concrete security parameters or overstating evidence.
- `MODEL_ROUTING_RATIONALE`: The assigned question is a bounded comparison across four passages against one accepted finding/patch plan and the approved spine. `gpt-6-luna` at `high` is the configured Reviewer-capable minimum sufficient route; escalate only for demonstrated capability insufficiency.

## Current-effective inputs and exact confirmation scope

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-011/review-findings-1.md` — F-001, its blocking impact, and requested correction.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-011/patch-plan-1.md` — exact accepted four-passage correction and boundaries.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-005/patch-report-1.md` and `launch-record.md` — implementer handoff; inspect current source passages independently.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` Q7 — current-effective Human-approved authority for the hard-to-guess credential and settled fragment/HMAC/expiry/status-only direction.
- Current Task 01 snapshot: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002/tasks/01-donation-domain-spec-reconciliation.md`.
- Product/MVP authority: `docs/product/mvp-scope.md` §5 and `docs/product/mvp-delivery-slices.md` §5.
- Inspect only the four passage scopes from the patch plan:
  1. `docs/spec/5-donation/invariants.md#inv-donation-05`;
  2. `docs/spec/5-donation/features/02-donation-status-check.md` reconciliation, active acceptance, and Credential guessing/theft threat row;
  3. `docs/spec/5-donation/tasks.md` Task 02 active acceptance;
  4. `docs/spec/5-donation/threat-model.md` temporary guest status URL / Spoofing row.
- Canonical targeted patch-review guidance: `../harscode-workspace/workflow/4-code-review-prompt.md` and `workflow/4-code-review/guidelines.md`, including the rule that a narrow patch exactly implementing an accepted patch plan does not automatically require another full four-pass review. Also use `workflow/orchestrated-run-overlay.md`, `workflow/context-management.md`, `workflow/AGENTS.md`, `orchestration/run-contract.md`, and `orchestration/AGENTS.md` for Run boundary/path rules.
- Target-repo conventions: root `AGENTS.md` and `docs/spec/README.md`.

## Review question and completion condition

Independently confirm whether BLD-S2-002-005 resolves F-001 exactly:

- The active invariant, Feature 02 acceptance, Task 02 acceptance, and status-URL Spoofing threat row state that the guest status bearer credential/token must be difficult to guess.
- This is presented as a settled Product/MVP and TP-015 requirement, not an undecided policy.
- No numeric entropy/length, generation algorithm, storage/comparison mechanism, or new implementation control is selected. Concrete generation/strength evidence and O4 controls/evidence/residual-risk remain open; none is claimed verified.
- The existing fragment handoff/URL cleanup, one-way HMAC verifier, hard 24-hour expiry, status-only projection, and uniform O5 failure direction remain unchanged in meaning.
- All five Donation specs remain `draft`; no Human/domain-owner acceptance or `CONTRACT_READY` is implied.

Compare only these exact passages with F-001, its patch plan, TP-015 Q7, and Product/MVP §5. Do not re-review unrelated parts of the five-file diff or older working-tree changes. Do not edit production/spec files or projections. No tests or runtime/security checks are indicated or authorized for this read-only confirmation.

Write only `review-confirmation.md` and `launch-record.md` under this Run. If F-001 is resolved, state the Task 01 Review loop is complete. If it remains unresolved or the patch introduced a new material contradiction, explain the exact source anchor and prepare a narrow `patch-plan-1.md` only if another source correction is required. Do not mark specs `agreed`, accept risk, or promote a milestone.

## Execution envelope

- `PREAUTHORIZED`: Read only F-001, its patch plan, BLD-005 handoff, the named spine/task/product authority, four live source passages, and targeted Review guidance; write this Run's confirmation and launch record; perform focused read-only comparison.
- `ORCHESTRATOR_DECISION`: Broaden Review or change the route.
- `HUMAN_REQUIRED`: Edit source, accept/agree specs, decide new policy/security detail, accept residual risk, or claim a milestone.

## Human-assisted dispatch

Use a fresh Reviewer Session from the Kencleng repository root with the configured model/effort. Start from `../harscode-workspace/workflow/4-code-review-prompt.md` and this Invocation. Kickoff:

`Jalankan targeted independent Review Run RV-S2-002-012 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-012/invocation.md dan canonical Code Review guidance Harscode saat ini. Konfirmasi hanya F-001 RV-011 terhadap patch plan-1, BLD-005 report, TP-015 Q7, Product/MVP §5, dan empat passage source yang ditunjuk. Verifikasi syarat difficult-to-guess hadir sebagai requirement settled di invariant, Feature 02 active acceptance, Task 02 active acceptance, dan threat-model Spoofing; kekuatan/generation evidence serta kontrol O4 tetap terbuka tanpa angka/implementasi baru; arah fragment/HMAC/24 jam/status-only/uniform 404 tidak berubah; semua spec tetap draft. Ini targeted confirmation untuk patch sempit, bukan full four-pass review. Tulis hanya review-confirmation.md dan launch-record.md; patch-plan-1 hanya jika masih diperlukan koreksi source. Jangan edit source/projection, jangan jalankan tests/runtime/security checks, jangan klaim Human acceptance, residual-risk acceptance atau CONTRACT_READY. Berhenti setelah phase handoff.`

After dispatch, reconcile from this Run's durable confirmation. A clean F-001 closure completes the Task 01 Review loop; Human/domain-owner acceptance of current drafts remains an independent parallel gate required before `CONTRACT_READY`. Task 02 follows its hard Task 01 dependency and applicable field-specific gates.
