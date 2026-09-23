# TPR-RES-TOP-001 — Resolve TPR-TOP-001 Retraction Verification Finding

WORK_UNIT_ID:
`WU-S1-005`

RUN_ID:
`TPR-RES-TOP-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TPR-RES-TOP-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005`

ROLE:
`Planner`

SPECIALIZATION:
Narrow Techplan resolution — topology verification coverage

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh resolution session

COMMUNICATION_LANGUAGE:
Bahasa Indonesia

COMMUNICATION_PROFILE_PATH:
`docs/project/communication-profile.md`

SELECTED_MODEL:
`gpt-5.6-terra`

REASONING_EFFORT:
`medium`

MODEL_APPROVAL:
`NOT_REQUIRED`

PRIOR_ARTIFACTS:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TP-TOP-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TPR-TOP-001/review-findings.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/EXP-TOP-001/evidence/solutioning.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TP-BE-001/techplan.md`

## Finding to resolve

TPR-TOP-001 found one MATERIAL / BLOCKING verification gap: the Techplan records retraction as required/delegated evidence but does not schedule an explicit fresh request through Caddy to the same previously-known media URL after Campaign public eligibility or media membership is withdrawn.

## Resolution scope

Revise only `TP-TOP-001/techplan.md` enough to make the already-settled retraction obligation executable.

The revised plan must:
- add an explicit rule/checklist obligation for retraction-through-proxy;
- coordinate with WU-S1-003 as the owner that mutates persisted eligibility/membership;
- require a fresh request through the root Caddy endpoint using the same known `content_url`;
- verify that new bytes are not delivered after withdrawal;
- preserve the contract boundary that already-downloaded/client-held bytes are outside the guarantee;
- name evidence owner(s) and explain that missing backend fixture/runtime capability defers execution but does not erase the obligation;
- keep `private, no-store`, direct-object denial, and 404/non-disclosure semantics intact.

Do NOT change:
- Caddy solution direction;
- MinIO policy direction;
- backend/frontend ownership;
- Product/API/security semantics;
- Build file scope;
- topology source scope.

After revision, explicitly declare whether material scope, architecture/ownership, business/security/interface semantics, or verification strategy changed. If this is only making an already-settled verification obligation executable, state that no material semantic change occurred and recommend whether re-review is necessary.

Do not begin Build.

Write a resolution note to:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TPR-RES-TOP-001/resolution.md`
