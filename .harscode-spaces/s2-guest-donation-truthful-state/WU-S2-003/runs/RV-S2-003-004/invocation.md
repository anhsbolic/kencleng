# Run Invocation — `RV-S2-003-004`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator after BLD-S2-003-001 ended `STALLED`. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `RV-S2-003-004`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-004`
- `ARTIFACT_TARGET`: `none` — write Review-owned `review-findings-1.md`; write `patch-plan-1.md` only if source changes are required.
- `WORK_UNIT_PATH` / compatibility `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent four-pass review of the completed public Campaign donation-cap projection slice from BLD-S2-003-001
- `PARTICIPANT_ID`: `P-S2-003-RV-004-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`; Registry SHA-256 `804856c5defad721678be4bac145fc2a1070307cccd720326cc598b8159d7abe`.
- `SESSION_TRANSITION`: `FRESH` — independent Review after the terminated Build Run; do not reuse its Participant or Session context.
- `TARGET_REVISION`: Kencleng HEAD `19d53315ac2847a03405339c0e86df8a97850761` plus current durable working-tree changes; inspect the live diff at dispatch and stop if these reviewed files materially differ from the pinned hashes below.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; current-effective ordinary Review guidance.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: The diff is a bounded public Campaign money projection, migration and focused tests; `gpt-6-luna` declares repository-work/reasoning capability and high effort is sufficient for an independent four-pass review. Do not escalate for missing authority/context; report those gaps. Approval-gated `gpt-6-sol` is not required for this bounded review unless an actual capability insufficiency is established.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Independent Code Review of the actual completed cap-projection portion of BLD-S2-003-001. This Review does not claim review of the incomplete whole Techplan, approve the migration for application, or authorize Tier-0 work.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Approved execution spine: `WU-S2-003/runs/TP-S2-003-006/techplan.md`, SHA-256 `b71951c86b2f74f6c9f1a164bcf43447c58ee655874deaf09275e6ff3b2316f9`, especially §§3–4, 8–13 and Test Focus Pointer. This is a source of implementation requirements, not evidence that the entire plan is complete.
- Build report: `WU-S2-003/runs/BLD-S2-003-001/report.md`, SHA-256 `6800f6433b32339cc9829aa97b2987e19d6bc2882cda16dcec51dd9237a4b32a`. Treat as orientation; independently inspect current diff/code.
- Root `AGENTS.md`, `backend/AGENTS.md`, current Campaign/Donation Product/MVP and accepted feature/invariant/API/monetary authorities referenced by the plan, and relevant backend architecture.
- Canonical `../harscode-workspace/workflow/4-code-review-prompt.md`, `workflow/4-code-review/guidelines.md`, `workflow/4-code-review/checklist.md`, `workflow/orchestrated-run-overlay.md`, `workflow/AGENTS.md`, applicable `workflow/context-management.md`, `../harscode-workspace/best-practices/AGENTS.md`, and `orchestration/run-contract.md`.
- Do not load raw Exploration logs; do not review unrelated WU003 future features as though implemented.

## Exact current Build diff / review scope

Review the actual current Build diff for these eight paths only. Current-content SHA-256 values at preparation:

| Path | Current SHA-256 | Baseline |
|---|---|---|
| `backend/internal/domain/campaign/entity.go` | `d19a89dd271d28755d9e4717ecddaf44adecefd27199ef4cac01fdd40e12627d` | Existing tracked file at target revision |
| `backend/internal/domain/campaign/repository_db.go` | `0f3496af3b3b273971e18545ab3a61a040df1e31db384192518d718bdbd43dd8` | Existing tracked file at target revision |
| `backend/internal/domain/campaign/service.go` | `9af4f156a46994ab99fe59257dcbcdff582a8e4b9e1e61e6004d0a9088aa293f` | Existing tracked file at target revision |
| `backend/internal/domain/campaign/service_test.go` | `531fc6cfb4567b86353aba80dc81c8c00804e18559d04e485156d34275e1f15b` | Existing tracked file at target revision |
| `backend/internal/transport/http/campaign_public.go` | `566d57262bf0c192f0841111c10ee41396f8d88ecd579a95b5964ec1ddb952b8` | Existing tracked file at target revision |
| `backend/internal/transport/http/campaign_public_test.go` | `df7376062229e07d2f2be18dd3d12686ff7f0c0cfe6173b15437f4cde7e96e8d` | Existing tracked file at target revision |
| `backend/migrations/000012_add_campaign_max_donation_amount.up.sql` | `1e69f949038f539a6763515de9f2bbac9ecd18a09f4f40eb0ad5280dfade1c6d` | New, untracked migration at preparation |
| `backend/migrations/000012_add_campaign_max_donation_amount.down.sql` | `d59cc62efd3e3dc24401e41b6b959854205b17ffe6a05c79ab68f74b79485c7c` | New, untracked migration at preparation |

Stop and report if a pinned file hash or actual scope differs materially at dispatch. Current diff is a public-read projection of the accepted per-Campaign IDR cap, seed/default mapping, response DTO and exact-wire coverage, plus additive cap migration/backfill. BLD reports `go test ./internal/domain/campaign ./internal/transport/http` and `git diff --check` passed; Orchestrator has not rerun tests. Migration remains unapplied.

The following are explicitly outside this Review scope because BLD did not change them: `backend/internal/domain/campaign/donation_coordinator_db.go`, `backend/internal/domain/donation/ledger.go`, Organization eligibility truth or Campaign draft create/PATCH handlers, Donation lifecycle, frontend/API/spec/Product sources, and all other files. Do not infer approval or rejection of the incomplete whole-plan implementation from this Review verdict.

## Review task and handoff

Run the four canonical passes in order on the same exact diff: Safety, Quality, Stack-Specific Best Practices, Consistency. Check the whole Rupiah constraints and exact cap mapping/default for the public projection; closed nested response shape and Funding-unavailable inclusion; additive migration/backfill/default/constraint/rollback fidelity; seed behavior and test quality; use of exact-decimal and parameterized SQL conventions; and whether source/API authority requires any additional projection behavior. These are review prompts derived from the plan and actual diff, not predetermined findings.

Use `best-practices/AGENTS.md` to route only Go, PostgreSQL/migration, and other directly relevant concerns. Run no broad suite by default. Use a targeted command only to resolve a concrete Review question; do not patch production code in Review. If changes are required, write a bounded patch plan for a fresh Build/Patch Run.

Write only `RUN_PATH/review-findings-1.md` and, if required, `RUN_PATH/patch-plan-1.md`. Include all four passes, actual verification posture, verdict and one structured `## Phase handoff` with `Outcome`, `Result refs`, `Findings`, `Decision requests`, `Blockers`, `Open / unverified`, `Recommended continuation`, and `Context refs`. Distinguish an implementation defect from deferred PostgreSQL/migration-application/Tier-0/OI7/runtime/security evidence. No source edits, migration application, Human acceptance, milestone promotion, or downstream dispatch.

## Execution envelope

- `PREAUTHORIZED`: read the pinned actual diff, approved Techplan and current applicable authorities; consult only best practices routed to this diff; write Review findings and a patch plan only if required.
- `HUMAN_REQUIRED`: no protected Tier-0 writes or authorization decisions; no manual migration/index application; no Product/API/spec acceptance, Open Item 7 resolution, residual-risk acceptance, or runtime/milestone claims.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Reviewer / `KC-REVIEWER` Session; configured `gpt-6-luna` / `high`.

Kickoff: `Jalankan independent Code Review Run RV-S2-003-004 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-004/invocation.md dan canonical ../harscode-workspace/workflow/4-code-review-prompt.md dengan orchestrated-run overlay. Review hanya actual diff delapan file yang dipin dalam Invocation melalui empat pass. Tulis review-findings-1.md dan phase handoff; patch-plan-1.md hanya jika diperlukan. Jangan mengedit source, menerapkan migration, mengubah gate, atau menganggap Review ini mengesahkan whole Techplan.`

Human reports completion or the exact scoped discrepancy to the Orchestrator. This Invocation prepares a Run only; it does not dispatch a Participant.
