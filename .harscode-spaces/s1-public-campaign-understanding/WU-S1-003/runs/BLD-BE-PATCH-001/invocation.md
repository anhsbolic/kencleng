# BLD-BE-PATCH-001 — Testing Re-entry Patch Invocation

WORK_UNIT_ID:
`WU-S1-003`

RUN_ID:
`BLD-BE-PATCH-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/BLD-BE-PATCH-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003`

ROLE:
`Implementer`

SPECIALIZATION:
Build/Patch — backend Campaign verification gaps from TST-BE-001

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh Build/Patch session

COMMUNICATION_LANGUAGE:
Bahasa Indonesia

COMMUNICATION_PROFILE_PATH:
`docs/project/communication-profile.md`

SELECTED_MODEL:
`gpt-5.6-terra`

REASONING_EFFORT:
`high`

MODEL_APPROVAL:
`NOT_REQUIRED`

PRIOR_ARTIFACTS:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TP-BE-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TST-BE-001/testing-report-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TST-BE-001/patch-plan-1.md`

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/3-build-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

PATCH RE-ENTRY:
Implement only the Testing gaps from `patch-plan-1.md` that belong to WU-S1-003.

Required:
1. Add minimal test-only testcontainers-go support and Campaign-local isolated Postgres/MinIO fixtures/tests required by R1/R6/R7.
2. Add the approved R5 authorization/not-found matrix and deterministic representative timing-parity evidence.
3. Triage and fix/safely justify only `gosec` findings introduced by or directly inside this Campaign Work Unit's changed scope.

Strict ownership boundary:
- Do NOT modify unrelated pre-existing Account/auth/OAuth/breachcheck/Tier-0 code merely to make repository-wide `make verify` green.
- Do NOT add blanket `#nosec` suppressions.
- Do NOT change repository-wide static-analysis policy/baseline in this Run.
- If `make verify` remains blocked only by unrelated pre-existing findings after Campaign-scope findings are cleared, report that exact separation as external/pre-existing blocker for Orchestrator/Human routing.

Local runtime:
- Podman is the configured Human local container engine.
- Probe testcontainers-go through the Podman Docker-compatible socket/API.
- Never use or mutate shared/manual `DATABASE_URL`.
- If the current sandbox cannot access the Human Podman socket, still author the harness correctly and report runtime evidence honestly; Testing must execute it in a capable environment before milestone promotion.

Focused Build/Patch verification:
- run new/changed Campaign unit/integration tests where executable;
- run Campaign-relevant `gosec` verification or the narrowest credible equivalent;
- run `make verify` only after scope-relevant blockers are handled, preserving whether remaining failures are unrelated pre-existing findings.

Do not start Testing in this Run.

Write:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/BLD-BE-PATCH-001/patch-report-1.md`

After patch, return to Testing, not Code Review, unless implementation materially changes production behavior/contract/architecture/security semantics beyond the approved patch plan.
