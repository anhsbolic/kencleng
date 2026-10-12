# Run Invocation — RUN-C1-ENG-CODEREVIEW-T2-LOGOUT-PATCH-001

## Identity / dispatch posture

WORK_UNIT_ID: WU-C1-ENG-001. RUN_ID: RUN-C1-ENG-CODEREVIEW-T2-LOGOUT-PATCH-001.
WORK_UNIT_PATH: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`. RUN_PATH: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-LOGOUT-PATCH-001`.
ARTIFACT_TARGET: none; Review evidence only.
ROLE / Profile: Reviewer / independent generic Reviewer, backend auth/session patch.
PARTICIPANT: PARTICIPANT-C1-ENG-T2-LOGOUT-PATCH-REVIEWER-001.
SESSION: SESSION-C1-ENG-T2-LOGOUT-PATCH-REVIEWER-001; SESSION_TRANSITION: FRESH, separate from Build Implementer and Orchestrator contexts. Do not treat Build reasoning as Review evidence.
SELECTED_MODEL / REASONING_EFFORT: gpt-6-sol / medium.
MODEL_APPROVAL: APPROVED by Anhar on 2026-10-12 (“approve bro”) for this exact RUN-C1-ENG-CODEREVIEW-T2-LOGOUT-PATCH-001 / gpt-6-sol / medium. No prior Run approval transferred. Participant remains not launched; Human mechanical dispatch pending.
MODEL_ROUTING_RATIONALE: independent review concerns production auth/session behavior and the ordering of validation before revocation; medium Sol is the lowest currently supported gated effort for this bounded security-relevant Review.
TARGET_REVISION / branch: `dae55ac44beb0e967b94c2f791d8a0b8d8e72372` / `pilot/3-c1-engineering`; current worktree includes the exact terminal Build delta bound below.
WORKFLOW_REVISION: `651ce57f044bff23fba99082322cee6c066d25a6`. HARSCODE_WORKSPACE_ROOT: `../harscode-workspace`.
Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. COMMUNICATION_LANGUAGE: Bahasa Indonesia.

## Claim / applicability / binding

Exact claim: the five-file logout patch rejects JSON null with the existing safe 400 `invalid_request` path before session revocation, and preserves valid `{}` logout plus unaffected Origin/CSRF/error/cookie behavior and the narrowly re-triaged G124 guard identity. Review applicability is recorded in `coordination/2026-10-12-auth-logout-patch-reconciliation.md`: this production auth/session behavior changed at the revocation boundary; independent Safety/contract-ordering reasoning provides material value. No whole-T2, independent runtime, or release acceptance claim is under Review.

The `input-bindings.json` SHA-256 is `71074a0657f102456ad598487dbc253595c76097657e5685bbc2346b92d73860`. It binds 12 assignment-defining inputs: this scope, terminal Build report, prior failed harness report, prior patch scope, Approved Techplan, T2 task snapshot, OpenAPI contract, and five exact changed files. Verify all before Review; stop and report on material drift. Lifecycle WU/Events/Control Tower, current coordination records, approval/dispatch state and ordinary current instructions are mutable context, not exact assignment locks. Expected terminal Run evidence creation is not drift.

Read current root/backend instructions; Harscode routers and `workflow/4-code-review-prompt.md`, `4-code-review/guidelines.md`, `4-code-review/checklist.md`, `workflow/orchestrated-run-overlay.md`, and `orchestration/run-contract.md`; bound Approved Techplan/T2 task, patch scope, live diff, exact API and Build/original failure reports. Route matching Go/HTTP/security best-practices only. Do not read raw unrelated history or change any source.

When invoked, execute all four canonical Review passes in order: Safety, Quality, Stack-Specific Best Practices, Consistency. Use the same exact five-file diff throughout. Review is independent reasoning; Build fixture/test results remain Builder-reported. No test/scanner/database/fixture/Podman/provider/browser command is authorized in this Review. If a concrete uncertainty cannot be resolved read-only, report it and let the Orchestrator route the gap.

## Review questions / limits

- Is a decoded `null` rejected before `Store.RevokeSession`, using the existing safe client error path, without weakening body/trailing-token/authentication/Origin/CSRF checks?
- Does `{}` remain the sole accepted empty object and retain 204/revocation/deletion-cookie behavior? Are the negative/non-revocation and positive regression assertions meaningful and not weakened?
- Does the new fixture runner bind ownership to this Build occurrence and select only the intended affected regressions without hiding assertions or exposing credentials?
- Is the `scanner_gate.py` change only the newly verified `http.go` source identity, with the same three accepted G124 candidates and no allowance or policy expansion?
- Do README claims stay within reported evidence, including fixture cleanup limitation and no independent runtime/Google/T2 acceptance?

Do not infer CSRF/privilege bypass from the original contract defect. Do not convert Build evidence into independent runtime acceptance. Preserve AUTH-HARNESS-001 terminal FAILED report and this Run's Build report unchanged.

## Output / terminality

Write `RUN_PATH/evidence/review-findings-1.md` with canonical four-pass sections, verification executed (`none` unless the authorized read-only reasoning establishes otherwise), scoped verdict, and exactly one `## Phase handoff`. Produce `patch-plan-1.md` only if a concrete source change is required. Do not edit source in Review. Terminal handoff closes this Run; any subsequent phase/re-entry requires a fresh Run/Participant/Session.

Orchestrator will reconcile the result and decide further claim-scoped applicability. An Approve verdict only supports the exact patch claim. `B-T2-002`, `B-T2-003`, historical `B-T2-001`, T2/T3 readiness/acceptance, dependency commit, Google login, and integrated verification remain outside this Run.

## Human-Assisted kickoff after approval

Fresh independent Reviewer/backend-auth Session, `gpt-6-sol` / `medium`, in `/home/anhar-solehudin/kencleng-workspace/kencleng`:
```
Execute RUN-C1-ENG-CODEREVIEW-T2-LOGOUT-PATCH-001 from .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-LOGOUT-PATCH-001/invocation.md. Verify the 12 exact bindings, then perform all four canonical Review passes on the five-file logout patch. Review null rejection before revocation, preserved {} behavior and the regression oracle, fresh fixture ownership, and the single G124 source identity update. Read-only Review: do not edit source, rerun tests/scanners/fixture, access DB/Podman/provider/browser, or claim independent runtime/T2 acceptance. Return review-findings-1.md SHA-256 and exactly one Phase handoff.
```
