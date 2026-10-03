# Run Invocation — `TST-S2-004-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-004`
- `RUN_ID`: `TST-S2-004-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TST-S2-004-001`
- `ARTIFACT_TARGET`: `none` — Testing report, launch record, and any temporary browser-verification harness belong to this Run; do not edit production sources.
- `WORK_UNIT_PATH` / compatibility `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Verifier
- `SPECIALIZATION`: Independent frontend flow verification, including approved Playwright R8 credential handoff evidence.
- `PARTICIPANT_ID`: `P-S2-004-TST-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-VERIFIER`; `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`; Registry SHA-256 `804856c5defad721678be4bac145fc2a1070307cccd720326cc598b8159d7abe`.
- `SESSION_TRANSITION`: `FRESH` — independent Testing after completed Build/Patch and targeted Review confirmation.
- `SESSION_TRANSITION_REASON`: Testing independence and new orchestrated phase Run.
- `TARGET_REVISION`: Kencleng HEAD `bb69cd002b3f1a1056837affcd77bb2b001007b0` plus current working tree. Review exactly the 19 pinned frontend paths below; fail closed if any hash differs.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; current-effective Testing guidance remains authoritative.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: Independent Testing must verify the scoped frontend Rules & Validation, run final frontend commands, and establish the approved browser evidence for the status bearer-credential handoff. The Human-owned registry's low-cost `gpt-6-luna` covers repository/testing work; medium is sufficient for this bounded single-WU verification. Escalate only if Testing demonstrates a material reasoning/capability shortfall; do not use a stronger model to replace missing authority or environment support.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Initial independent Testing Run after BLD-S2-004-001, patch BLD-S2-004-002, four-pass Review RV-S2-004-002, and targeted confirmation RV-S2-004-003 approval of F01/F02.
- `RISK_TIER`: Tier 1 per Approved Techplan; mock/browser evidence does not establish backend/security/runtime controls or accept residual risk.

## Current-effective inputs

- Approved Techplan: `runs/TP-S2-004-003/techplan.md`, SHA-256 `e6a93bc65a760084955f63a469b238fd70a80c20ff38cc1d22ba1be000ad99eb` — Rules & Validation §§3–4, Testing Checklist §12, Test Focus Pointer, open items, and test ownership are the verification contract.
- Initial Build report: `runs/BLD-S2-004-001/report.md`, SHA-256 `af721c79995f73647ba627895f7e3fdf726134e6cbe9049bfd4bb7b2059f4d0e` — treat its 17 focused tests, lint, and production-build results as prior claims; account for deferred R8, R10, Integration Map, and backend/security evidence.
- Patch Build report: `runs/BLD-S2-004-002/patch-report-001.md`, SHA-256 `ebc8c51c45d14a0d617fdaf42ff667e0bb6799d4fec3531d41751843fdaf0bfb` — Participant reports 1 file / 15 focused tests passed for F01/F02; transient manual checking state was not deterministically asserted.
- Latest Code Review: `runs/RV-S2-004-003/review-findings-001.md`, SHA-256 `4223afc7c62d629d6f77b2be2a16649f19245a4cd655d7c2a44ecca028813c74`; F01/F02 confirmed resolved, no new finding. Review states no tests/build/browser were rerun.
- Prior full Review: `runs/RV-S2-004-002/review-findings-001.md`, SHA-256 `bfa58cace1bf1d2ceeed5e0f6325bdad50e96ec8114c2a2886a5a8be090c1a66`; its Invocation has the original 19-path hash baseline.
- Root `AGENTS.md`, `frontend/AGENTS.md`, actual `frontend/package.json` scripts, `frontend/playwright.config.ts`, and current frontend architecture/design rules relevant to Testing. Canonical `../harscode-workspace/workflow/5-testing-prompt.md`, Testing guidelines/checklist, `workflow/orchestrated-run-overlay.md`, `workflow/AGENTS.md`, `orchestration/AGENTS.md`, and `orchestration/run-contract.md`.
- Re-open exact Test Focus evidence anchors only: `EXP-S2-004-001/evidence/stage-2-gap-analysis.md` §§Area 1 (`#area-1--product-donation-domain-and-approved-delivery-boundary`), Area 2 (`#area-2--authored-api-contract`), and Area 4 (`#area-4--frontend-architecture-and-live-implementation`); `stage-3-solutioning.md` §§Optional email and Guest status credential and proof boundary. Do not load the full Exploration corpus.
- Current observable entry points are Public Campaign Detail, `/campaigns/{campaignId}/donate`, and `/donations/{donationId}/status`, running against contract-shaped MSW where the accepted real backend is not available.

## Exact current implementation set

The 19 paths below are the full frontend set pinned in RV-S2-004-002. Orchestrator compared all current hashes to that baseline: exactly the three F01/F02 patch files changed since the Review target; 16 are unchanged. Verify every current hash below before Testing and inspect the exact current diff, not only phase reports.

| SHA-256 | File |
|---|---|
| `69c6f1a613a82297470332e43cbeda1ed7726882ee25c4762e712435713a65c4` | `frontend/app/campaigns/[campaignId]/campaign-detail-client.test.tsx` |
| `afdc387c13491ecadec349b413c6985d9790be66317bed54234dc267696be398` | `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` |
| `f3d21dce3605f748623db510a96176f84990a0d3fbf554cb8e8e8420c929c158` | `frontend/app/campaigns/[campaignId]/campaign-detail.module.css` |
| `91b571ae8032a22576b40b2acd8e6f2b3d8d4986bb14fe88e40946c1593ad30b` | `frontend/app/campaigns/[campaignId]/donate/donation-client.tsx` |
| `93f13cb2b3f290741c728cebf2ebe9fff9e3db8477075ddc9cc484fd7d501499` | `frontend/app/campaigns/[campaignId]/donate/donation.module.css` |
| `39551623021ede4529521fcdbfb587d1b3ee812d1cc648fa5f745e010e4ebebc` | `frontend/app/campaigns/[campaignId]/donate/page.tsx` |
| `33341f32092036ab92757843a55af7a823f6b52d779c9784e03bf0eeef50a745` | `frontend/app/campaigns/[campaignId]/mock-service-worker.tsx` |
| `2bdec978c88b2ee6ec40a3c0f7a21f51f14d5100ee7d79395e6b5ba22c500e9a` | `frontend/app/donations/[donationId]/status/page.tsx` |
| `80e535548eb7b2f6eca7434227cb8eebecd64ea9f886c650446da5f44c01a5b2` | `frontend/app/donations/[donationId]/status/status-client.tsx` |
| `77f1b67c02e9476bb864910bc2bf79b0baf2242a02399850321e414042fadabb` | `frontend/app/donations/[donationId]/status/status.module.css` |
| `82bdd24998c9566aa17948aa123341c002b3e3f232185c523ecf74ffb7f2ba86` | `frontend/app/donations/donation-flow.test.tsx` |
| `290157ee9044d58c1433249d26298cf8aea35f7c0de5fef2f76db4e5739d595c` | `frontend/app/donations/layout.tsx` |
| `c4a9ede71ecce0dc2b33af643226e0a1aa465af96361fd5d1b89cabac97b28a1` | `frontend/app/mock-service-worker.module.css` |
| `9e30f932b9bca4b42ec519a96f9d42a2b2f1041385e9fc5185747b76b70a44b1` | `frontend/app/mock-service-worker.tsx` |
| `629f02238772ae2709dddbe5828b33114e89fc36e036b509c3aa8780e113e8a3` | `frontend/lib/api/client.ts` |
| `0bda1df8cb0b2f25c27af258eab2d30bb05c72a1704a826d9e3191c286c3d7ab` | `frontend/lib/api/donation.ts` |
| `99bae42736b5d60630f06ca0c81f8b3e64d13c2bad49ab013145c2624220af15` | `frontend/mocks/browser.ts` |
| `775baf8a9d0f43215b0d43aaa1e8d6d4babef57704573a3d298e25304e980238` | `frontend/mocks/handlers/donation.ts` |
| `1885c31afb5e70939570213dad3fd1bd41eb8e73f2e44e6411e50971ab74ee4c` | `frontend/mocks/server.ts` |

## Execution scope and evidence obligations

- Follow canonical Testing Step 0, all relevant Techplan Rules & Validation R1–R11, the Test Focus Pointer, error behavior, final verification, backward-compatibility applicability, and a fresh end-to-end Techplan consistency read. Do not select a verdict in advance.
- Independently verify/spot-check existing unit coverage for R1–R7 and R9–R11. Build reports 17 initial focused tests and 15 patch tests; do not treat those reports as independent evidence. Run the target repo's fast baseline `cd frontend && npm run verify` (lint plus unit/component tests) and `cd frontend && npm run build` as final current-state checks: production build was run before the accessibility patch, so final Next.js compilation and route validation are not yet evidenced for the exact tested revision.
- Techplan-approved R8 real-browser automation is required here: component tests cannot establish fragment/history/visible-URL behavior. Use Playwright/Chromium from `frontend/playwright.config.ts` and the existing `test:browser` capability for a narrowly scoped status-link handoff. Confirm the fragment credential reaches `X-Donation-Status-Credential`, is removed from the visible URL before steady rendered state, the route shows status only, and the resulting state is accessible. The approved plan makes a second permission prompt unnecessary. There are currently no committed `frontend/tests/browser` specs; a temporary focused harness and its outputs may be kept under this Run path. Do not change production sources or unrelated tests. If browser/runtime tooling is unavailable, record the exact blocker and stop that evidence claim.
- Test Focus security interpretation: verify the frontend credential lifecycle and generic UI only. Do not infer token strength/expiry, infrastructure/referrer/log/cache protection, anti-enumeration parity, or backend authorization from Playwright/MSW. Backend, payment, persistent idempotency/concurrency, email fulfillment, and settlement evidence remain with WU-S2-003 and Security/API owners.
- R10 includes automated semantic role/name/focus coverage as applicable; actual representative desktop/mobile rendered acceptance and keyboard exercise remain Human-owned. Do not mark that gate passed.
- Integration Map Donation-flow mapping is a separate Orchestrator follow-up before integrated delivery, not a test failure or Test Run write target.
- `PREAUTHORIZED`: read the named authorities/current diff; run approved frontend `verify`, a justified production build if required for final evidence, and focused Playwright R8 verification; create Run-local harness/output only under this Run path; write `testing-report-001.md` and `launch-record.md` plus `patch-plan-001.md` only if code changes are required.
- `HUMAN_REQUIRED`: manual rendered acceptance R10, any product/design/API/security/architecture decision, protected write, residual-risk acceptance, or runtime/real-integration milestone. Do not fix code during Testing.
- Report exact commands/results, source scope, rule coverage, errors, Test Focus execution, fresh Techplan read, verified/assumed/deferred/not-tested distinctions, and remaining owners/gates. Mocks prove frontend behavior only. Do not claim `FRONTEND_MOCK_VERIFIED`, WU/Slice completion, production readiness, residual-risk acceptance, or integration.

## Required dispatch prompt

Jalankan independent Testing Run `TST-S2-004-001` sesuai Invocation ini, canonical `../harscode-workspace/workflow/5-testing-prompt.md`, dan `workflow/orchestrated-run-overlay.md`. Verifikasi exact 19-file source hashes lalu ikuti Step 0, seluruh R1–R11 dan Test Focus Pointer pada Techplan current-effective; perlakukan hasil Build sebagai claims yang perlu diverifikasi. Jalankan `cd frontend && npm run verify`, `cd frontend && npm run build` untuk current patched state, serta Playwright Chromium R8 yang telah diotorisasi di Techplan. Simpan harness/output browser sementara di Run path jika perlu; jangan mengedit production atau source test files. Tulis `testing-report-001.md` dan `launch-record.md` dengan bukti aktual, verdict independen, batas Human R10, integration-map follow-up, dan seluruh deferred backend/security/runtime scope. Berhenti setelah phase handoff; jika ada code finding, buat `patch-plan-001.md` dan kembalikan ke Build.
