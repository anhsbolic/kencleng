# Run Invocation — `TP-S2-003-010`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-04 after independent Review RV-S2-003-006 found one material/blocking mismatch in the WU-S2-003 candidate's status-credential security authority. This is one bounded candidate-resolution occurrence; it does not authorize source changes or a downstream phase.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `TP-S2-003-010`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-010`
- `ARTIFACT_TARGET`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/techplan.candidate.md` — continue the existing single Draft/In Review successor; preserve Approved TP-S2-003-006 unchanged.
- `WORK_UNIT_PATH` / compatibility `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Resolve RV-S2-003-006's material O4 credential-authority finding against current accepted Donation sources
- `PARTICIPANT_ID`: `P-S2-003-TP-010-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Planner occurrence after independent Reviewer RV-S2-003-006; do not reuse the Reviewer or TP-S2-003-009 Participant Session.
- `TARGET_REVISION`: Kencleng HEAD `4e71d3697a479d92b2cad5a79ef5d710f4e8077c` plus current working tree. Candidate SHA-256 `3aa5e5d362bbfac5cb65b231159d793552d57cc61236fbb126bcbadee3416a60`; RV-S2-003-006 findings SHA-256 `ff1683903f6d9b785f66f0a68a23b0ef30d847beb754fbf79481700ca5d0d3c0`. Verify both before editing and stop/report material drift.
- `WORKFLOW_REVISION`: current-effective Harscode Techplan synthesis and orchestrated-run overlay. Preparation identity: synthesis prompt SHA-256 `1ed5bdc6e8bb70beac4dc61f4328c5cc0a8b5a70cc82ae7dba4ce22a9c588613`; overlay SHA-256 `f806673aa5e7d1d82dae9f433ec54c3562e15ce778b709a141c57d3d22dfaf87`. Re-read current-effective guidance at dispatch.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned local registry `.harscode-spaces/.local-config.yaml`, SHA-256 `ddffdeb4ce8edec181ca36bf32bedf19c513a76e359bc2642985eda6b32f400d`.
- `MODEL_ROUTING_RATIONALE`: This Run resolves one tightly evidenced planning contradiction against current authority and must preserve the larger WU003 spine and scoped gates. `gpt-6-luna` / `high` is supported by the local registry and is consistent with the prior WU003 Planner Runs; the open O4 owner decisions are not solved by a higher model.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Canonical `workflow/2-1-techplan-synthesis-prompt.md` for one material-resolution pass on the existing post-Approval candidate, with `workflow/orchestrated-run-overlay.md`. Do not create a separate/versioned Techplan or approval report during review/resolution churn.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Current Approved predecessor remains `runs/TP-S2-003-006/techplan.md`, SHA-256 `b71951c86b2f74f6c9f1a164bcf43447c58ee655874deaf09275e6ff3b2316f9`; do not edit or promote it.
- Candidate to resolve: `techplan.candidate.md`, SHA-256 `3aa5e5d362bbfac5cb65b231159d793552d57cc61236fbb126bcbadee3416a60`; last synthesis handoff `runs/TP-S2-003-009/handoff.md`, SHA-256 `357e70ca4c8c73ef350741571ae7a4843e274de72f7f720454a9b4118b3ebf37`.
- Sole finding to resolve: `runs/RV-S2-003-006/review-findings-1.md`, SHA-256 `ff1683903f6d9b785f66f0a68a23b0ef30d847beb754fbf79481700ca5d0d3c0`, finding `MATERIAL / BLOCKING — Security authority`. It concerns §3 Q8, §4 R8, §5 D6, §12 R8 verification wording, and §13 Active Open Items 3 and 5. Review target was the exact candidate hash above.
- Current owning authorities: `docs/spec/5-donation/invariants.md#inv-donation-05`, SHA-256 `4bf7ba48f4980f635ce76a4eb3440be5cf12666fdd99e06dcc8ac3b78fbde39a`; `docs/spec/5-donation/features/02-donation-status-check.md`, SHA-256 `332bba24b3f00cab880f6d5671d3e3c5cf9d99633be6fdf0212f0a05b5b5fd2c`. Both settle the difficult-to-guess bearer, fragment handoff/URL cleanup, one-way HMAC direction, hard 24-hour expiry and status-only contract, while leaving concrete credential generation/strength evidence, key/comparison controls, expiry enforcement, exposure/abuse controls, empirical evidence, and residual-risk acceptance open under O4/O5. They do not select a numeric entropy/length target.
- The candidate currently states a 256-bit CSPRNG bearer, dedicated-purpose verifier, and constant-time comparison as chosen in Q8/R8/D6 and says credential construction is settled in O4. These statements conflict with the current sources above and the Review evidence. Read the exact candidate sections, Review report and cited current authorities directly.
- Relevant durable Exploration: `runs/EXP-S2-003-001/evidence/stage-2-gap-analysis.md` Area 4 and `stage-3-solutioning.md` O4/O5 direction. They preserve the settled URL/HMAC/24-hour/status-only directions while leaving credential detail and controls open.
- Re-read current `AGENTS.md`, `backend/AGENTS.md`, and applicable Product/spec/API/security authority as needed for the targeted correction. Keep exact D1 authorization, Human pairing, unresolved migration-design RV-S2-003-005 findings F-02/F-03/F-04/F-06, Open Item 7, O3 and O5 gates intact.
- Canonical authorities to re-read: `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`; `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`; `workflow/orchestrated-run-overlay.md`; `workflow/AGENTS.md`. Use `guidelines.md` only if process clarification is needed.

## Resolution task and boundaries

Resolve the single blocking Review finding by reconciling Q8, R8, D6, their verification wording, and relevant §13 Open Items against current Donation INV-donation-05 and Feature 02. Preserve settled fragment handoff/cleanup, one-way HMAC verifier, hard 24-hour expiry, status-only response, and settled uniform `404`/cache directions. Do not treat the historical predecessor's specific credential recipe as current authority. Do not invent or label as Human-selected a numeric entropy/length, random-generation implementation, dedicated-key provisioning, comparison mechanism, expiry-control implementation, or residual-risk acceptance. Keep any currently unselected O4 choices and evidence with their owner/open-item routing; if the current authoritative sources do not identify a precise owner or implementation gate, surface that gap rather than guessing.

Check that Q8/R8, D6 status, §12 verification ownership/rationale, and §13 Open Items tell one consistent story; retain material source anchors and consequences. Declare whether the resolution materially changes security/interface meaning, risk, ownership, or verification strategy. If it does, recommend a fresh independent Review of the newer candidate before the Human Techplan gate unless a Human gate explicitly waives that re-review. Do not resolve unrelated findings or migration-design issues.

## Execution envelope and output

- `PREAUTHORIZED`: read the exact candidate, RV6 report, prior handoff, current Donation invariant/feature and required workflow guidance; update only the stable candidate and write this Run's structured handoff.
- `HUMAN_REQUIRED`: selecting any new material credential/security policy; approving/promoting the candidate; changing Product/domain/API authority; modifying protected crypto/auth code; migration application; protected writes, risk acceptance, or downstream Run dispatch.
- Do not generate `report-techplan.md` while review/resolution churn remains. Do not edit specs, API, code, tests, migrations, or predecessor plan. Run no tests, validators, generators, database/migration actions, or downstream Participants.
- Write `RUN_PATH/handoff.md` with exactly one structured `## Phase handoff` using current orchestrated fixed fields. Identify the updated candidate hash, finding disposition, any materiality/re-review recommendation, and remaining scoped blockers. Do not duplicate the candidate in the Run.

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Planner / `KC-PLANNER`, configured `gpt-6-luna` / `high`.

Kickoff: `Jalankan resolution Planner Run TP-S2-003-010 memakai canonical ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md dan orchestrated-run overlay, sesuai invocation. Verifikasi hash candidate dan RV-S2-003-006 report; baca current INV-donation-05 dan Feature 02 serta Exploration O4 evidence. Koreksi hanya konflik security-authority ini pada candidate yang sama: pertahankan arah fragment/HMAC/24-jam/status-only yang settled, tetapi jangan tampilkan numeric strength, recipe credential, key/comparison control sebagai keputusan settled tanpa authority. Selaraskan Q8/R8/D6, verifikasi R8, dan Open Items; jangan menyelesaikan pilihan O4/O5 atau gate D1/Item 7 dengan asumsi. Nyatakan materialitas correction dan apakah candidate baru perlu fresh independent Review sebelum Human gate. Jangan edit predecessor/source/spec/tests/migrations, jangan buat report approval, dan jangan dispatch Run lain. Tulis satu structured handoff lalu berhenti.`

This Invocation prepares a Run only; Anhar dispatches the fresh Planner. No Participant has been dispatched by the Orchestrator.
