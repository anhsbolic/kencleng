# Run Invocation — `BLD-S2-004-002`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-004`
- `RUN_ID`: `BLD-S2-004-002`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/BLD-S2-004-002`
- `ARTIFACT_TARGET`: `none` — production and regression-test changes belong to the scoped frontend paths; Run-owned evidence is `RUN_PATH/patch-report-001.md` and `RUN_PATH/launch-record.md`.
- `WORK_UNIT_PATH` / compatibility `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Implementer
- `SPECIALIZATION`: Narrow frontend Build/Patch re-entry for two Review findings.
- `PARTICIPANT_ID`: `P-S2-004-BLD-002-1`
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`; Registry SHA-256 `804856c5defad721678be4bac145fc2a1070307cccd720326cc598b8159d7abe`.
- `SESSION_TRANSITION`: `FRESH` — orchestrated Build re-entry creates a new Run, Participant, and Session after completed Review; ground only on this Invocation and the named patch inputs.
- `SESSION_TRANSITION_REASON`: Context hygiene and phase re-entry; the prior Build/Review occurrences have completed.
- `TARGET_REVISION`: Kencleng HEAD `bb69cd002b3f1a1056837affcd77bb2b001007b0` plus current working tree. Reopen the three listed live frontend files before editing; unrelated existing frontend changes remain outside this patch assignment.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; ordinary current-effective Build/Patch guidance remains authoritative.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `low`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: This is a bounded mechanical frontend patch with exact accepted findings, a patch plan, and a single existing observable test file. The registry's low-cost `gpt-6-luna` supports coding/repository work; low effort is sufficient. Escalate only if live code exposes a materially broader or ambiguous problem, not to compensate for missing context.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Build/Patch re-entry requested by completed independent Code Review `RV-S2-004-002`; after this patch, return to the requesting Review phase using a new targeted confirmation Run, not a new full four-pass review unless the patch broadens scope or changes semantics.

## Current-effective patch inputs

- Approved Techplan: `runs/TP-S2-004-003/techplan.md`, SHA-256 `e6a93bc65a760084955f63a469b238fd70a80c20ff38cc1d22ba1be000ad99eb`.
- Requesting Review findings: `runs/RV-S2-004-002/review-findings-001.md`, SHA-256 `bfa58cace1bf1d2ceeed5e0f6325bdad50e96ec8114c2a2886a5a8be090c1a66`.
- Exact patch plan: `runs/RV-S2-004-002/patch-plan-001.md`, SHA-256 `f42cf05d6118f2f2724faa53a7441d7f69e71431ab271101c924158b916a1b8c`.
- Initial Build evidence for context only: `runs/BLD-S2-004-001/report.md`, SHA-256 `af721c79995f73647ba627895f7e3fdf726134e6cbe9049bfd4bb7b2059f4d0e`.
- Root `AGENTS.md`, `frontend/AGENTS.md`, canonical Build prompt/guidelines/checklist, `workflow/orchestrated-run-overlay.md`, `workflow/AGENTS.md`, `orchestration/AGENTS.md`, and `orchestration/run-contract.md`.
- Live file anchors at preparation:
  - `frontend/app/campaigns/[campaignId]/donate/donation-client.tsx` SHA-256 `93f46b815d3fe4cfccf5f385f67de82991293c28702313554cdb8556d988711b`
  - `frontend/app/donations/[donationId]/status/status-client.tsx` SHA-256 `f8f88b7bc5723d778d0b8cd4e5d8cfedee85ed4c682ae0e57b36a1d18d4a5155`
  - `frontend/app/donations/donation-flow.test.tsx` SHA-256 `d2529a68d6c05617aa9cba5627423751728194e8a7c12f8225cfc2b18aa398a4`

## Risk and execution envelope

- Risk tier remains Tier 1 under the approved WU004 Techplan; this patch changes accessibility behavior only and does not close separate guest-credential, PII, backend, or runtime risks.
- `PREAUTHORIZED`: read required source/phase guidance; edit only the two production components above and `frontend/app/donations/donation-flow.test.tsx` for focused regression assertions directly required by the fixes; run the focused test command for that file; write only this Run's `patch-report-001.md` and `launch-record.md` as execution evidence.
- `HUMAN_REQUIRED`: material Product/design/API/security/architecture/verification-authority changes, scope expansion beyond these frontend files, protected/Tier-0 writes, manual DB/index actions, or residual-risk acceptance. Stop and report if live code invalidates the approved patch assumptions.
- Out of scope: other frontend files, `backend/`, authored/generated `api/`, Product/spec/design authority, shared docs, tracker, orchestration projections, and every unrelated Work Unit. Do not fix the Integration Map follow-up in this patch.
- Focused verification: `cd frontend && npm run test -- app/donations/donation-flow.test.tsx`. Run it after adding/changing assertions to prove the relevant regressions. Do not run broad suites, production build, browser automation, or security/concurrency/performance checks unless a concrete patch question makes one necessary. Independent Testing still owns R8 browser evidence; R10 Human rendered acceptance is separate.
- Do not claim `FRONTEND_MOCK_VERIFIED`, Testing completion, Human rendered acceptance, runtime readiness, residual-risk acceptance, WU completion, or Slice completion.

## Patch target

Apply only `RV-S2-004-002/patch-plan-001.md`:

1. F01: associate the conditional guest email validation message with the email input using the required accessible description relationship while keeping invalid state synchronized and removing stale description when valid.
2. F02: ensure initial and manual status-check outcomes (pending, terminal success/failure, and unavailable) are announced to assistive technology while keeping the recheck control usable and not moving focus unexpectedly.
3. Add focused observable regression assertions in the existing Donation flow test for both accessibility behaviors, preserving the approved user-facing copy, status-only flow, generic failure behavior, and request semantics.

Do not alter settled Product/API semantics, weaken Review findings, or redesign broader UI/accessibility architecture.

## Required dispatch prompt

Jalankan Harscode Build/Patch Run `BLD-S2-004-002` sesuai Invocation ini, canonical `../harscode-workspace/workflow/3-build-prompt.md`, dan `workflow/orchestrated-run-overlay.md`. Baca hanya Techplan Approved, finding dan patch plan yang dipatok, panduan fase, serta instruksi frontend yang relevan; re-ground tiga file live sebelum mengedit. Terapkan patch F01/F02 dalam envelope frontend, tambahkan assertion observable terarah pada test flow yang ada, lalu jalankan focused test command yang tercantum dan laporkan hasil aktual. Tulis `patch-report-001.md` dan `launch-record.md` lengkap dengan phase handoff. Setelah selesai, arahkan kembali ke Review untuk targeted confirmation. Jangan mengubah authority/contract, memperluas scope, atau mengklaim Testing, rendered acceptance, runtime, risk acceptance, maupun milestone delivery.
