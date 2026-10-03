# Run Invocation — `RV-S2-004-002`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-004`
- `RUN_ID`: `RV-S2-004-002`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/RV-S2-004-002`
- `ARTIFACT_TARGET`: `RUN_PATH/review-findings-001.md`; optional `RUN_PATH/patch-plan-001.md` only if code changes are required. Do not edit production code.
- `WORK_UNIT_PATH` / compatibility `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Complex — full four-pass independent review of Tier-1 guest donation frontend implementation.
- `PARTICIPANT_ID`: `P-S2-004-RV-002-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`; Registry SHA-256 `804856c5defad721678be4bac145fc2a1070307cccd720326cc598b8159d7abe`.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer Session, separate from Planner and Implementer contexts.
- `TARGET_REVISION`: Kencleng HEAD `bb69cd002b3f1a1056837affcd77bb2b001007b0` plus current working tree. At dispatch, verify the exact 19-file frontend source set and per-file SHA-256 list below; fail closed if it differs.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; ordinary current-effective review guidance remains authoritative.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: Tier-1 independent review requires sufficient careful multi-pass reasoning. The Human-declared low-cost `gpt-6-luna` covers repository work and reasoning; high effort is sufficient for this bounded frontend diff and security-sensitive guest credential review. Do not substitute model escalation for missing context or authority.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Initial independent Code Review after Build `BLD-S2-004-001`; Complex gate; four passes against the same exact frontend diff.

## Current-effective inputs and review target

- Approved execution contract: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TP-S2-004-003/techplan.md`, SHA-256 `e6a93bc65a760084955f63a469b238fd70a80c20ff38cc1d22ba1be000ad99eb`.
- Build evidence: `../BLD-S2-004-001/report.md`, SHA-256 `af721c79995f73647ba627895f7e3fdf726134e6cbe9049bfd4bb7b2059f4d0e`; launch record SHA-256 `28f057df7d987bba2edaa90977a14c94190e250479c676a0c34117dc20c4e46a`; Build Invocation SHA-256 `609cc8a6577d66648cd5e607a16dabd1ae81bcbf4627858af8ad5011746d82a0`. Reported tests/build are Implementer evidence; Review should not replay them absent a specific finding/question.
- Reviewed revision is the 19 frontend files listed below. Aggregate sorted-path/content-hash set SHA-256: `5d85f4063e48c8e394fca3a5b0d9c2c519ee1457477e6d6742fe555b86dd2977`. Verify all member hashes immediately before Review. Read dependencies for context without expanding the reviewed diff.

| SHA-256 | Current frontend file |
|---|---|
| `69c6f1a613a82297470332e43cbeda1ed7726882ee25c4762e712435713a65c4` | `frontend/app/campaigns/[campaignId]/campaign-detail-client.test.tsx` |
| `afdc387c13491ecadec349b413c6985d9790be66317bed54234dc267696be398` | `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` |
| `f3d21dce3605f748623db510a96176f84990a0d3fbf554cb8e8e8420c929c158` | `frontend/app/campaigns/[campaignId]/campaign-detail.module.css` |
| `93f46b815d3fe4cfccf5f385f67de82991293c28702313554cdb8556d988711b` | `frontend/app/campaigns/[campaignId]/donate/donation-client.tsx` |
| `93f13cb2b3f290741c728cebf2ebe9fff9e3db8477075ddc9cc484fd7d501499` | `frontend/app/campaigns/[campaignId]/donate/donation.module.css` |
| `39551623021ede4529521fcdbfb587d1b3ee812d1cc648fa5f745e010e4ebebc` | `frontend/app/campaigns/[campaignId]/donate/page.tsx` |
| `33341f32092036ab92757843a55af7a823f6b52d779c9784e03bf0eeef50a745` | `frontend/app/campaigns/[campaignId]/mock-service-worker.tsx` |
| `2bdec978c88b2ee6ec40a3c0f7a21f51f14d5100ee7d79395e6b5ba22c500e9a` | `frontend/app/donations/[donationId]/status/page.tsx` |
| `f8f88b7bc5723d778d0b8cd4e5d8cfedee85ed4c682ae0e57b36a1d18d4a5155` | `frontend/app/donations/[donationId]/status/status-client.tsx` |
| `77f1b67c02e9476bb864910bc2bf79b0baf2242a02399850321e414042fadabb` | `frontend/app/donations/[donationId]/status/status.module.css` |
| `d2529a68d6c05617aa9cba5627423751728194e8a7c12f8225cfc2b18aa398a4` | `frontend/app/donations/donation-flow.test.tsx` |
| `290157ee9044d58c1433249d26298cf8aea35f7c0de5fef2f76db4e5739d595c` | `frontend/app/donations/layout.tsx` |
| `c4a9ede71ecce0dc2b33af643226e0a1aa465af96361fd5d1b89cabac97b28a1` | `frontend/app/mock-service-worker.module.css` |
| `9e30f932b9bca4b42ec519a96f9d42a2b2f1041385e9fc5185747b76b70a44b1` | `frontend/app/mock-service-worker.tsx` |
| `629f02238772ae2709dddbe5828b33114e89fc36e036b509c3aa8780e113e8a3` | `frontend/lib/api/client.ts` |
| `0bda1df8cb0b2f25c27af258eab2d30bb05c72a1704a826d9e3191c286c3d7ab` | `frontend/lib/api/donation.ts` |
| `99bae42736b5d60630f06ca0c81f8b3e64d13c2bad49ab013145c2624220af15` | `frontend/mocks/browser.ts` |
| `775baf8a9d0f43215b0d43aaa1e8d6d4babef57704573a3d298e25304e980238` | `frontend/mocks/handlers/donation.ts` |
| `1885c31afb5e70939570213dad3fd1bd41eb8e73f2e44e6411e50971ab74ee4c` | `frontend/mocks/server.ts` |

## Review envelope and boundaries

- Read root `AGENTS.md`, `frontend/AGENTS.md`, the exact approved Techplan, and the live current diff. Reopen relevant accepted product/design/spec/API sources only as needed for Consistency. Do not load raw Exploration history.
- Canonical phase inputs: `../harscode-workspace/workflow/4-code-review-prompt.md`, `workflow/4-code-review/guidelines.md`, `workflow/4-code-review/checklist.md`, `workflow/orchestrated-run-overlay.md`; orchestration uses `../harscode-workspace/orchestration/AGENTS.md` and `run-contract.md`.
- Pass 3 routing: `../harscode-workspace/best-practices/AGENTS.md` and targeted `best-practices/index.md` trigger/security rows. Open only matching best-practice documents for React/Next.js async flow, forms/input validation, HTTP error handling, and security-sensitive client credential handling. Cite exact best-practice sources for any Stack-Specific finding; if no trigger matches, state that explicitly.
- Pass 4 target conventions: root `AGENTS.md`, `frontend/AGENTS.md`, relevant section(s) of `docs/project/kencleng-frontend-tech-stack.md`, current `docs/ui-ux/` routing/visual authority, relevant `docs/spec/` and split `api/openapi/` sources. Review the actual implementation for equality with generated contract behavior; do not treat Build report as proof.
- Known follow-up from Build: report says `docs/project/kencleng-integration-map.md` has Public Campaign Detail but no Donation-flow row. The file was outside Build envelope. Assess whether this creates a material code/contract inconsistency or is a coordination-doc handoff; record the exact conclusion and preserve it as open follow-up if not a review blocker. Do not edit shared docs in Review.
- Read-only scope: no production edits, API/spec/Product/design/tracker/manifest/Work Graph/Control Surface updates, generated files, tests, runtime, or browser runs. Run only a targeted reproduction if needed to decide a concrete finding; report exact command/question. Write Review-owned findings and patch-plan only under this Run path.
- Required Review passes, in order, against same diff: Safety; Quality; Stack-Specific Best Practices; Consistency. Consider guest credential fragment cleanup/header use, PII/email opt-in, amount/cap disclosure, idempotency/retry lifecycle, safe API failures, accessibility semantics, generated types, and MSW staying at network boundary against approved plan.
- Verdicts: `Approve`, `Approve with minor comments`, or `Request changes`. A behavior-changing fix returns via a new Build/Patch Run; do not promote minor comments into a patch loop.
- Outputs: complete `RUN_PATH/review-findings-001.md` with four passes, verification posture, verdict, and structured phase handoff. Write `RUN_PATH/patch-plan-001.md` only if code changes are required.

## Required dispatch prompt

Jalankan independent Complex Code Review Run `RV-S2-004-002` sesuai Invocation ini dan canonical `../harscode-workspace/workflow/4-code-review-prompt.md` dengan `workflow/orchestrated-run-overlay.md`. Verifikasi exact source-set hash sebelum mulai; review hanya 19-file frontend diff yang dipatok beserta konteks read-only. Jalankan keempat pass secara independen dan berurutan. Jangan mengedit production code atau mengulang test/build matrix tanpa pertanyaan review konkret. Tulis verdict, setiap finding dengan lokasi/severity/resolusi, kesimpulan atas integration-map follow-up, serta phase handoff lengkap dalam `review-findings-001.md`; buat `patch-plan-001.md` hanya jika perubahan kode diperlukan.
