# Run Invocation — `TST-S2-003-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator after approved Review RV-S2-003-004. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `TST-S2-003-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TST-S2-003-001`
- `ARTIFACT_TARGET`: `none` — write Testing-owned `testing-report-1.md`, plus `launch-record.md` for Run provenance.
- `WORK_UNIT_PATH` / compatibility `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Verifier
- `SPECIALIZATION`: Independent observable verification of BLD-S2-003-001's public Campaign cap projection slice
- `PARTICIPANT_ID`: `P-S2-003-TST-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-VERIFIER`; `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`; Registry SHA-256 `804856c5defad721678be4bac145fc2a1070307cccd720326cc598b8159d7abe`.
- `SESSION_TRANSITION`: `FRESH` — independent Testing after Build and Code Review; do not reuse either Participant or Session.
- `TARGET_REVISION`: Kencleng HEAD `19d53315ac2847a03405339c0e86df8a97850761` plus current durable working-tree changes; inspect exact current source and run the pinned checks against that state.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; current-effective ordinary Testing guidance.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: This is a narrow observable projection verification against accepted wire shape and exact decimal constraints, with a known intentionally incomplete parent Build and explicit independent database/runtime gates. `gpt-6-luna` declares repository-work/reasoning capability and high effort is sufficient to execute and report the bounded evidence. Do not escalate for missing implementation, authority, or environment; report those limits. Approval-gated `gpt-6-sol` is not required unless a concrete capability insufficiency is established.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Independent Testing of the implemented public Campaign cap projection only, after Review verdict `Approve`. This is a bounded sub-slice of BLD-S2-003-001, not final Testing of the incomplete WU-S2-003 whole Techplan. Do not report WU completion or `BACKEND_VERIFIED`.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Approved spine: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-006/techplan.md`, SHA-256 `b71951c86b2f74f6c9f1a164bcf43447c58ee655874deaf09275e6ff3b2316f9`; relevant cap/wire behaviors in §§3–4, 8–13 and Testing Checklist R2/R3/R14/R15. The whole spine remains incomplete and its D1/Tier-0, Open Item 7, O3/O4/O5 and downstream evidence gates remain active.
- Latest Build report: `WU-S2-003/runs/BLD-S2-003-001/report.md`, SHA-256 `6800f6433b32339cc9829aa97b2987e19d6bc2882cda16dcec51dd9237a4b32a`. It reports `STALLED`, implemented cap projection/wire slice and unapplied migration, Participant-reported focused `go test ./internal/domain/campaign ./internal/transport/http` and `git diff --check` passes, and all unimplemented/deferred items.
- Independent Review: `WU-S2-003/runs/RV-S2-003-004/review-findings-1.md`, SHA-256 `4ecd9310a4d2c98acb157b5ddc4ab395796b5c076ad1c5ca27f145d198e20eef`; verdict `Approve` for the eight-file cap projection/migration diff only; no findings; migration application and testing remain open.
- Current WU003 manifest, parent Events/Work Graph/Outcome/Control Surface and development tracker.
- Read root `AGENTS.md`, `backend/AGENTS.md`, `backend/README.md`, the actual `backend/Makefile`, and relevant accepted Campaign detail spec/API/monetary authorities. Canonical current guidance: `../harscode-workspace/workflow/5-testing-prompt.md`, `workflow/5-testing/guidelines.md`, `workflow/5-testing/checklist.md`, `workflow/orchestrated-run-overlay.md`, `workflow/AGENTS.md`, `../harscode-workspace/orchestration/run-contract.md`, and applicable backend best-practice routing in `../harscode-workspace/best-practices/AGENTS.md`.
- Do not load raw Exploration logs. No current Test Focus row requires specialized concurrency/security execution for this public read-only cap projection; D1/concurrency and guest credential/email areas are not implemented by this Build slice and remain deferred to whole-WU Testing after implementation.

## Exact current Build diff / verification scope

At preparation, current-content hashes still match the completed Review's pinned eight-file scope:

| Path | Current SHA-256 |
|---|---|
| `backend/internal/domain/campaign/entity.go` | `d19a89dd271d28755d9e4717ecddaf44adecefd27199ef4cac01fdd40e12627d` |
| `backend/internal/domain/campaign/repository_db.go` | `0f3496af3b3b273971e18545ab3a61a040df1e31db384192518d718bdbd43dd8` |
| `backend/internal/domain/campaign/service.go` | `9af4f156a46994ab99fe59257dcbcdff582a8e4b9e1e61e6004d0a9088aa293f` |
| `backend/internal/domain/campaign/service_test.go` | `531fc6cfb4567b86353aba80dc81c8c00804e18559d04e485156d34275e1f15b` |
| `backend/internal/transport/http/campaign_public.go` | `566d57262bf0c192f0841111c10ee41396f8d88ecd579a95b5964ec1ddb952b8` |
| `backend/internal/transport/http/campaign_public_test.go` | `df7376062229e07d2f2be18dd3d12686ff7f0c0cfe6173b15437f4cde7e96e8d` |
| `backend/migrations/000012_add_campaign_max_donation_amount.up.sql` | `1e69f949038f539a6763515de9f2bbac9ecd18a09f4f40eb0ad5280dfade1c6d` |
| `backend/migrations/000012_add_campaign_max_donation_amount.down.sql` | `d59cc62efd3e3dc24401e41b6b959854205b17ffe6a05c79ab68f74b79485c7c` |

If material drift or extra source changes exist at dispatch, stop and route the mismatch before testing. No production source edits are authorized.

## Test task and execution envelope

Verify the observable Campaign detail projection for the implemented cap slice: required closed `max_donation_amount` object, major-unit whole-IDR string and explicit `IDR`; inclusive valid bounds/default mapping and invalid/missing/fractional/out-of-range fail-closed behavior; presence regardless of Funding availability; and compatibility of the current wire key shape. Spot-check the Build report's named focused command in a fresh verifier session. Use the repository's actual scripts/commands; record each exact command/result.

Scope this Testing occurrence to the implemented slice. State explicitly that unimplemented Donation submit/status, D1 ordering, capacity reservation, exact-once settlement, Organization eligibility writes/handlers, O3/O4/O5 security/runtime, and full WU behavior have not been tested and remain for a later whole-spine Testing Run. Do not give a pass for those deferred requirements.

Migration evidence: inspect the up/down SQL and migration numbering/schema collision read-only. The migration is unapplied. Do not run `make migrate-up/down` or manually apply it to a shared/local project database. If an isolated disposable PostgreSQL migration exercise is judged necessary, stop before applying and route that exact setup/application boundary to the Human; record the evidence gap meanwhile. A Testing verdict for the code slice does not approve migration application.

Verification selection: run `go test ./internal/domain/campaign ./internal/transport/http` from `backend/` to independently verify named Build coverage and observable HTTP behavior. Run a broader command only if live target-repo authority or a concrete risk makes it necessary; `backend/Makefile`'s full `verify` target includes staticcheck, all unit tests, full race, contract, gitleaks and govulncheck, so do not invoke it merely by habit for this partial read-projection slice. This is not the final whole-WU verification. `git diff --check` may be inspected if needed but is not a substitute for behavior testing.

Do not edit production/tests or create a migration/test workaround in this Run. If testing reveals a code defect, write a bounded patch plan and route to a fresh Build/Patch Run. If the code slice passes, keep migration database execution, full-plan verification, PostgreSQL/concurrency, all Human-owned gates and final risk acceptance visible as deferred.

Write `RUN_PATH/testing-report-1.md` and `RUN_PATH/launch-record.md`. The report must follow the canonical Testing format, include one structured `## Phase handoff` with `Outcome`, `Result refs`, `Findings`, `Decision requests`, `Blockers`, `Open / unverified`, `Recommended continuation`, `Context refs`, and explicitly state its verdict applies only to this bounded cap-projection slice.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Verifier / `KC-VERIFIER` Session; configured `gpt-6-luna` / `high`.

Kickoff: `Jalankan independent Testing Run TST-S2-003-001 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TST-S2-003-001/invocation.md dan canonical ../harscode-workspace/workflow/5-testing-prompt.md dengan orchestrated-run overlay. Verifikasi hanya public Campaign cap projection slice yang dibangun dan di-Review; uji command terfokus sesuai Invocation, inspect migration secara read-only, jangan apply migration. Pisahkan verdict slice ini dari WU-S2-003/whole-Techplan yang masih incomplete. Tulis testing-report-1.md, phase handoff terstruktur, dan launch-record. Jangan edit source atau mengklaim BACKEND_VERIFIED.`

Human reports completion or the exact scoped discrepancy to the Orchestrator. This Invocation prepares a Run only; it does not dispatch a Participant.
