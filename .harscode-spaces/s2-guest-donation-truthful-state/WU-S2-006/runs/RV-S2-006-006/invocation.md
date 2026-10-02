# Run Invocation — `RV-S2-006-006`

Status: `READY_FOR_HUMAN_DISPATCH`
Prepared: 2026-10-02. Human-facing prose Bahasa Indonesia.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-006`
- `RUN_ID`: `RV-S2-006-006`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-006`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent four-pass review of WU-S2-006 Campaign/Donation spec and authored split OpenAPI source changes
- `PARTICIPANT_ID`: `P-S2-006-RV-006-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — independent Review Participant/Session after completed BLD-S2-006-005.
- `TARGET_REVISION`: Kencleng HEAD `d186feb2ce3d447d31bccd038dc3b9703f1d2369` plus current working tree; the seven authored source files below are the complete current Build diff. Re-ground exact live diff at dispatch; this is a baseline, not a semantic pin on unrelated files.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; current-effective and not semantically pinned.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned registry.
- `MODEL_REGISTRY_SOURCE`: `.harscode-spaces/.local-config.yaml`, SHA-256 `ddffdeb4ce8edec181ca36bf32bedf19c513a76e359bc2642985eda6b32f400d`; `gpt-6-luna` declares `reasoning`/`repository-work`, supports `high`, cost tier `low`, `approval_required: false`.
- `MODEL_ROUTING_RATIONALE`: Lowest-cost available Reviewer-capable model. High effort is proportionate to cross-domain monetary representation, capacity/error semantics, public response shape and source compatibility across the current diff. No stronger-model escalation is justified.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Independent Code Review after initial authored-source Build; four canonical passes over the same current source diff. This is not source-owner acceptance or final Testing.
- `ARTIFACT_TARGET`: `none` — write Review-owned `review-findings.md` under this Run; create `patch-plan.md` only if the review identifies code/source changes required.

## Current-effective inputs / PRIOR_ARTIFACTS

- Canonical `../harscode-workspace/workflow/4-code-review-prompt.md`, `workflow/4-code-review/guidelines.md`, `workflow/4-code-review/checklist.md`, `workflow/orchestrated-run-overlay.md`, and `orchestration/run-contract.md`.
- Root `AGENTS.md`, `.harscode-spaces/participant-profiles/profiles.md` / `KC-REVIEWER`, applicable Product/MVP and monetary authorities, current Campaign/Donation feature specs and invariants, `api/README.md`, WU-S2-006 manifest, Authority Map, parent Events, Work Graph, and current accepted WU-S2-002/WU-S2-005 baselines.
- Current approved Techplan `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-008/techplan.md`, SHA-256 `93c09bf7629500cd8fb80fd59b6af464b169484419d722a269b782b78bbbf438`, especially §§3–13. Human approval receipt and matching report identity are recorded in parent `events.md`.
- Build result `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-005/report.md`, SHA-256 `b9f0670fcc6bbb41a49ac6f09c4fb44bdaead9f0ebb486851b3c4e89b947d806`. Treat it as orientation/evidence only; independently inspect the current diff and changed sources.
- Exact Human owner decisions DEC-API-01/02: parent `events.md`, “2026-10-02 — Owner settled API transport/encoding; Planner propagation prepared”. Product and six prior source acceptances are specific to their recorded snapshots; these seven changed bytes remain unaccepted.

## Exact current diff / source identities

Review the diff against the current `TARGET_REVISION` baseline. These are the seven and only seven authored source files in the current Build diff; do not expand scope to generated/API bundle/frontend/runtime files.

| Source | Baseline SHA-256 | Current SHA-256 |
|---|---|---|
| `docs/spec/4-campaign/invariants.md` | `3ca67e64aadd0d210c4de5a789ccc60172ea1a4e7f25de599dc6793db5e38931` | `b46797a183a12f818989ca16553751a001cb358abfda02f6570af4a02c98c27f` |
| `docs/spec/4-campaign/features/01-campaign-creation-draft-crud.md` | `59945e6437b0356936fe92f0778f7529d4155bc3036db95476df473d1d59aaa0` | `06dd64a2f4df04193ecd5d89dfc64f599df142d93997f6bf7222ee68c02e255c` |
| `docs/spec/4-campaign/features/02-campaign-detail-listing.md` | `4ad59a8c1ec3e1a5c188f5a63dee72e2661852decd455903e35fb46aab775fae` | `88758edd31f5d0d867bc02b8a5b7d425d571da0097bbc4d96cf91e4978444aff` |
| `docs/spec/5-donation/invariants.md` | `af3cbe68819a045a4db3479d704f228043de7476cc81bbc678a0fe13c1b5de9f` | `adf13ce0eb23ff6a29bcf54fa6c63ab7cd1b685bf77531198831d414328628b0` |
| `docs/spec/5-donation/features/01-submit-donation-settlement.md` | `85cda7bf84bdcbe7013fe7e9e9a26cab061219988f69009f7841203a2f153174` | `b65511829e97276f6df3a109061a6d8f5506cefd6b6bee2e08a7b016ed8c42a3` |
| `api/openapi/campaign.yaml` | `65464f9aa177579a4a1536dc1d599160e3eab18eb1ab687feb9cdcf04c0f3134` | `12a20e4be31b32df8ee73794400ce73f2ae15f2c8b48b9107a377cc83fba7226` |
| `api/openapi/donation.yaml` | `873834d59695fa8ae6011f6958101a496581d3f9c03a84ed28ab6f2263e5ea88` | `609411688477132267847db52ec6c7d78b5bfd50df97d7af4e85c049917ad2ca` |

`docs/spec/4-campaign/features/09-closure.md` and `api/openapi/common.yaml` were not changed; their unchanged identities are in the Build report. Stop and report if the diff, baseline, or any current-source hash differs materially from this snapshot.

## Task, review boundary, and handoff

Run all four canonical Code Review passes in order against this same current diff: Safety, Quality, Stack-Specific Best Practices, and Consistency. Use the approved TP006008 as the execution contract and current Product/MVP/API/spec authorities for consistency. For Pass 3, route through `../harscode-workspace/best-practices/AGENTS.md` and only open matching sources; do not load the raw Exploration corpus or reproduce the Build verification matrix by default. The review must remain independent and must not edit sources.

Review the exact object closure/member requirements and range; create/PATCH omission, default/preserve/freeze and required response semantics; explicit currency in every applicable public response state; generic eligible-capacity no-fit `422` on `amount`; distinct predicates, unchanged closed/ineligible `409` and idempotent retry behavior; capacity/close-reason non-disclosure; `funding_capacity_reached` ownership and public-projection boundary; source/spec/OpenAPI consistency; and whether `api/README.md` counterpart sequencing is respected. These are review areas derived from the approved contract and actual diff, not predetermined findings or verdicts.

Review is primarily independent reasoning. No broad tests/validation rerun is requested or authorized by this Invocation; use a targeted command only if needed to resolve a concrete review uncertainty, document why, and do not patch in Review. Produce one Run-local `review-findings.md` with the four review passes, verdict, actual verification posture, and exactly one structured `## Phase handoff` containing `Outcome`, `Result refs`, `Findings`, `Decision requests`, `Blockers`, `Open / unverified`, `Recommended continuation`, and `Context refs`. Create `patch-plan.md` only for required source changes. Reviewer verdict does not accept source bytes, promote a milestone, or imply runtime correctness.

After a clean/approved Review, Orchestrator routes the concrete changed Campaign/Donation spec and API bytes to their named owners for explicit acceptance. Do not start bundle/generated type/fixture/consumer counterpart work or claim WU-S2-006 complete until the independent Review and required owner gates converge. If Request changes, return to a fresh Build/Patch Run using the specific patch plan; preserve this Review artifact.

## Execution envelope

- `PREAUTHORIZED`: read assigned current diff, exact prior artifacts, current authorities and applicable routed best practices; write Review findings and a patch plan only when warranted.
- `ORCHESTRATOR_DECISION`: reconcile the Review verdict, route owner acceptance or a specific follow-up phase, and update coordination projections.
- `HUMAN_REQUIRED`: owning acceptance of changed Product/spec/API bytes, any new authority decision, protected Tier-0 authorization, DB application, or residual-risk acceptance.

## Human-assisted dispatch

Cwd `/home/anhar-solehudin/kencleng-workspace/kencleng`; fresh Reviewer / KC-REVIEWER, `gpt-6-luna` / `high`.

Kickoff: `Jalankan independent Code Review Run RV-S2-006-006 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-006/invocation.md dan canonical ../harscode-workspace/workflow/4-code-review-prompt.md dengan orchestrated-run overlay. Berhenti setelah review-findings.md dan phase handoff; jangan edit source.`

Laporkan Review artifact atau exact scoped discrepancy kepada Orchestrator. Jangan lanjut ke owner acceptance atau counterpart work dalam Run ini.
