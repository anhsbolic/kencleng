# Run Invocation — `RV-S2-002-011`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `RV-S2-002-011`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-011`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent four-pass review of Task 01 post-approval Donation spec reconciliation
- `PARTICIPANT_ID`: `P-S2-002-RV-011-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; current Profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new independent Reviewer Run after completed Build/Patch Run BLD-S2-002-004.
- `SESSION_TRANSITION_REASON`: Review independence; use a new Reviewer Participant/Session and reconstruct from the durable plan, refreshed task, Build report, and actual source diff.
- `TARGET_REVISION`: Kencleng checkout `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus current durable approval/task snapshots, BLD-004 source diff, and working-tree artifacts; inspect the exact five-file diff at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; re-read the canonical Review entrypoint and routed best-practice files at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Required full independent Code Review of the current Task 01 source diff after a material approved-spine reconciliation. Execute all four canonical passes: Safety, Quality, Stack-Specific Best Practices, and Consistency. Review does not grant spec acceptance or a milestone.
- `RISK_TIER`: Tier 1 — changed specifications govern exact monetary values, guest PII/email lifecycle, bearer status credentials, anti-enumeration behavior, and unresolved security/evidence boundaries.
- `MODEL_ROUTING_RATIONALE`: Independent cross-document review of five narrowly scoped domain-spec changes against a single approved spine and current Product/MVP authority. `gpt-6-luna` at `high` is the configured Reviewer-capable minimum sufficient routing; escalate only if review capability is demonstrably insufficient, not to compensate for missing context or authority.

## Current-effective inputs and exact scope

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` — current-effective Human-approved `Approved` spine; authoritative requirements, decisions, risks, verification obligations, Open Items, and `CONTRACT_READY` boundary.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002/tasks/01-donation-domain-spec-reconciliation.md` — current Task 01 scope and reference IDs, refreshed after TP-015 approval.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-004/report.md` and `launch-record.md` — Build handoff and stated verification; inspect the source diff independently.
- The exact Review diff is limited to these five changed files, relative to the target repository base/current working tree; inspect them with the actual Git diff and current full text:
  - `docs/spec/5-donation/invariants.md`
  - `docs/spec/5-donation/threat-model.md`
  - `docs/spec/5-donation/tasks.md`
  - `docs/spec/5-donation/features/01-submit-donation-settlement.md`
  - `docs/spec/5-donation/features/02-donation-status-check.md`
- Authority/convention sources: root `AGENTS.md`; `docs/spec/README.md`; `docs/product/README.md`, `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md` §§5–6; `docs/project/kencleng-monetary-data-standard.md`; relevant `docs/ui-ux/` sources cited by the changed specs; `.harscode-spaces/authority-map.md`; relevant sections of `docs/kencleng-agentic-workflow.md`.
- Prior Task 01 review history `RV-S2-002-007` / `RV-S2-002-008` and earlier Build reports may be consulted only where needed to understand still-current prior source treatment. Do not treat older reports as authority or review unrelated pre-existing changes.
- Canonical current Review guidance: `workflow/4-code-review-prompt.md`, `workflow/4-code-review/guidelines.md`, `workflow/4-code-review/checklist.md`, `workflow/orchestrated-run-overlay.md`, `workflow/context-management.md`, `workflow/AGENTS.md`, `orchestration/run-contract.md`, and `orchestration/AGENTS.md`.

## Review assignment

Run the four canonical passes against the same exact source diff, in order:

1. **Safety:** determine whether the revised invariants/acceptance create unsafe security, privacy, money, state/funding, or lifecycle claims, weaken a settled boundary, or overstate evidence/risk acceptance. If a specialized verification gap is newly created, compare with TP-015 §12/Test Focus and report Techplan drift separately.
2. **Quality:** check for material ambiguity, stale contradictions, inconsistent terminology, duplicate/conflicting acceptance, and reviewability across the five specs. Do not report cosmetic preferences as defects.
3. **Stack-Specific Best Practices:** route from `../harscode-workspace/best-practices/AGENTS.md` and `best-practices/index.md` only to matching concerns. The changed source is documentation, not implementation; assess whether applicable guidance materially bears on the stated money/PII/client-facing-security requirements. The Security Concern Map's `enumeration-and-timing` route to `restapi/anti-enumeration.md` is likely applicable. Consider `go/decimal-and-money.md` or other matches only if the actual concern fits; if no further trigger matches, say so explicitly. Cite a matching best-practice file for any finding. Do not force an implementation-specific rule onto a domain spec.
4. **Consistency:** verify the diff against root `AGENTS.md`, Product/MVP/design routing, `docs/spec/README.md`, monetary standard, and current TP-015/Task 01. In particular verify: O1 representation is settled and only its concrete parameters remain open; O11/D19 terminal-notice direction is carried without inventing bound/mechanism/risk acceptance; O4 fragment/HMAC/24-hour/status-only directions and O5 uniform `404`/identical Problem Details/body/header/cache directions are not reopened; empirical/security controls remain unproven; D1/O7/simulator ownership and Slice 2 exclusions remain intact; source statuses remain `draft`; exact API expression stays with Task 02; no `CONTRACT_READY` is implied.

Treat TP-015 as the reviewed execution contract and the live Product/MVP/design/monetary sources as their owning authorities. Do not reconstruct product intent from stale specs or raw Exploration. Report any material source gap as Techplan drift instead of making a new decision.

## Verification posture and output

Review is primarily independent reasoning over the exact Git diff, Build evidence, and live authority. Use a targeted read-only command only if it resolves a concrete uncertainty. Do not edit spec/source files, change status to `agreed`, update projections, or claim implementation/runtime/security evidence. Do not add or run tests; no runtime reproduction is in scope for this documentation-only diff.

Write `review-findings-1.md` in this Run using the canonical four-pass report structure and compact provenance. If a production/spec correction is required, write `patch-plan-1.md` with exact changes and return it to a fresh Build/Patch Run; do not make the correction here. If no patch is required, do not create a patch plan. Record targeted checks, or `none`, honestly.

Human/domain-owner acceptance of the current Donation spec drafts remains a separate independent gate and is required before `CONTRACT_READY`; the Reviewer must not claim or record that acceptance. Task 02 remains downstream of its hard Task 01 dependency and applicable field-specific gates. No automatic Testing dispatch or milestone follows from a clean Review.

## Execution envelope

- `PREAUTHORIZED`: Read only the listed authority/task/report/current five-file diff and routed review guidance; write this Run's `review-findings-1.md` and `patch-plan-1.md` only if required; perform targeted read-only inspection.
- `ORCHESTRATOR_DECISION`: Broaden Review scope, alter the route, or introduce additional source targets.
- `HUMAN_REQUIRED`: Edit source, agree/approve specs, make a Product/Design/API/Security decision, accept residual risk, or accept a milestone.

## Human-assisted dispatch

Use a fresh Reviewer Session at the Kencleng repository root with the configured model/effort. Start from `../harscode-workspace/workflow/4-code-review-prompt.md` and this Invocation. Kickoff:

`Jalankan independent full four-pass Code Review Run RV-S2-002-011 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-011/invocation.md dan canonical Code Review prompt/guidance Harscode saat ini. Periksa actual diff tepat lima file domain spec Donation yang disebut Invocation terhadap current-effective TP-S2-002-015 Approved, Task 01 TPD-S2-002-002, Product/MVP, Design, monetary standard, root AGENTS.md, dan docs/spec/README.md. Jalankan Safety, Quality, Stack-Specific Best Practices, lalu Consistency; route best-practice pass melalui Security Concern Map terutama enumeration-and-timing/anti-enumeration, dan nyatakan jelas bila tidak ada kecocokan lain. Catat finding dengan lokasi, dampak, resolusi, dan blocking/non-blocking; buat patch-plan-1 hanya bila source correction diperlukan. Jangan edit spec atau projection, jangan jalankan tests/runtime checks, jangan tandai draft agreed atau klaim Human acceptance, residual-risk acceptance, runtime proof, maupun CONTRACT_READY. Tulis review-findings-1.md dan launch-record.md milik Run ini, lalu berhenti untuk handoff.`

After dispatch, reconcile only from this Run's durable review evidence. A patch returns to a new Build/Patch Run; an approve verdict routes through the still-required Human/domain-owner spec acceptance gate before orchestration advances the Work Unit.
