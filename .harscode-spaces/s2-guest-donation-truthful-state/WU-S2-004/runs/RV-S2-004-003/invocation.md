# Run Invocation — `RV-S2-004-003`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-004`
- `RUN_ID`: `RV-S2-004-003`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/RV-S2-004-003`
- `ARTIFACT_TARGET`: `none` — Review-owned confirmation evidence is `RUN_PATH/review-findings-001.md`; create `patch-plan-001.md` only if a new code fix is required.
- `WORK_UNIT_PATH` / compatibility `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Targeted accessibility-finding confirmation; not a full four-pass re-review.
- `PARTICIPANT_ID`: `P-S2-004-RV-003-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`; Registry SHA-256 `804856c5defad721678be4bac145fc2a1070307cccd720326cc598b8159d7abe`.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer Session after completed Build/Patch Run; do not resume RV-S2-004-002's Session.
- `SESSION_TRANSITION_REASON`: Independence and orchestrated phase re-entry after the patch Run.
- `TARGET_REVISION`: Kencleng HEAD `bb69cd002b3f1a1056837affcd77bb2b001007b0` plus current working tree. Verify the three current file hashes below before Review; fail closed if they differ.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; ordinary current-effective Review guidance remains authoritative.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: This independent confirmation must assess whether two accessibility fixes and their observable assertions satisfy the exact accepted findings. The Human-owned registry's low-cost `gpt-6-luna` has reasoning/repository-work capability; medium is sufficient for careful but tightly scoped code review. A stronger approval-gated model is not justified unless a material cross-cutting issue emerges.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Targeted confirmation by the requesting Code Review phase for RV-S2-004-002 F01/F02 after Build/Patch `BLD-S2-004-002`. Harscode's current Code Review prompt permits targeted confirmation when a narrow accepted patch does not broaden scope or change material semantics; do not repeat all four passes unless the diff reveals such a change.

## Current-effective inputs and exact confirmation target

- Approved execution contract: `runs/TP-S2-004-003/techplan.md`, SHA-256 `e6a93bc65a760084955f63a469b238fd70a80c20ff38cc1d22ba1be000ad99eb`.
- Prior findings: `runs/RV-S2-004-002/review-findings-001.md`, SHA-256 `bfa58cace1bf1d2ceeed5e0f6325bdad50e96ec8114c2a2886a5a8be090c1a66`.
- Accepted patch plan: `runs/RV-S2-004-002/patch-plan-001.md`, SHA-256 `f42cf05d6118f2f2724faa53a7441d7f69e71431ab271101c924158b916a1b8c`.
- Patch evidence: `runs/BLD-S2-004-002/patch-report-001.md`, SHA-256 `ebc8c51c45d14a0d617fdaf42ff667e0bb6799d4fec3531d41751843fdaf0bfb`; launch record SHA-256 `352d44a045e983748c1e325d625d8c2ca9b7b2c0de0ae1f1c9828242f8321f09`.
- Previous Review's pinned pre-patch 19-file set is in `runs/RV-S2-004-002/invocation.md`. Orchestrator compared current files to those exact hashes: 16 remain unchanged; only the three patch-authorized files below differ.

| Current SHA-256 | File | Prior pinned SHA-256 |
|---|---|---|
| `91b571ae8032a22576b40b2acd8e6f2b3d8d4986bb14fe88e40946c1593ad30b` | `frontend/app/campaigns/[campaignId]/donate/donation-client.tsx` | `93f46b815d3fe4cfccf5f385f67de82991293c28702313554cdb8556d988711b` |
| `80e535548eb7b2f6eca7434227cb8eebecd64ea9f886c650446da5f44c01a5b2` | `frontend/app/donations/[donationId]/status/status-client.tsx` | `f8f88b7bc5723d778d0b8cd4e5d8cfedee85ed4c682ae0e57b36a1d18d4a5155` |
| `82bdd24998c9566aa17948aa123341c002b3e3f232185c523ecf74ffb7f2ba86` | `frontend/app/donations/donation-flow.test.tsx` | `d2529a68d6c05617aa9cba5627423751728194e8a7c12f8225cfc2b18aa398a4` |

- Read root `AGENTS.md`, `frontend/AGENTS.md`, the canonical `../harscode-workspace/workflow/4-code-review-prompt.md`, Review guidelines/checklist, and `workflow/orchestrated-run-overlay.md`. Re-read only the exact findings, patch plan, approved Techplan accessibility anchors, the cited accessibility best-practice source, Build patch report, and three current files/tests. Do not load raw Exploration or unrelated Run history.
- Confirmation questions: F01 — when the guest-email error is present, is its text programmatically associated with the email input and does the association clear when valid? F02 — does the persistent polite status region announce initial and manual pending, terminal success/failure, and unavailable results without stealing focus or changing request/recheck behavior? Do the focused assertions observe those behaviors?

## Review envelope and verification posture

- Read-only: do not edit production, tests, Product/spec/API/design, trackers, manifests, Events, Work Graph, Control Surface, or other WU sources. Write only `RUN_PATH/review-findings-001.md`; add a patch plan only for new required code changes.
- Build reports `cd frontend && npm run test -- app/donations/donation-flow.test.tsx` passed after the final markup: 1 file / 15 tests. Treat this as Participant-reported evidence; do not rerun it unless a concrete Review uncertainty needs reproduction.
- No full four-pass Review or broad test/build/browser matrix is required for this narrow accepted patch. Do not claim independent Testing, R8 browser verification, R10 Human rendered acceptance, runtime/security evidence, integration, residual-risk acceptance, or WU completion.
- The Integration Map Donation-flow row remains an Orchestrator coordination follow-up before integrated delivery; prior Review found no code/contract inconsistency and this Review must not edit the map.

## Required dispatch prompt

Jalankan targeted Code Review confirmation Run `RV-S2-004-003` sesuai Invocation ini, canonical `../harscode-workspace/workflow/4-code-review-prompt.md`, dan `workflow/orchestrated-run-overlay.md`. Verifikasi tiga source hash yang dipatok dan konfirmasi scope patch tetap tiga file. Periksa hanya resolusi RV-S2-004-002 F01/F02 terhadap patch plan yang diterima, implementasi live, dan assertion observable; full four-pass re-review serta test/build matrix tidak perlu kecuali ditemukan perluasan atau ketidakpastian konkret. Jangan edit code/tests. Tulis `review-findings-001.md` dengan verdict per F01/F02, bukti lokasi dan kesimpulan scope; buat `patch-plan-001.md` hanya bila ada fix baru yang diperlukan. Jika keduanya resolved, rekomendasikan independent Testing sebagai langkah berikutnya dan pertahankan R8, R10, integration-map serta runtime/security sebagai gate terpisah.
