# TST-INT-001 — Slice 1 Real Integration Verification

WORK_UNIT_ID:
`WU-S1-006`

RUN_ID:
`TST-INT-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-006/runs/TST-INT-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-006`

ROLE:
`Verifier`

SPECIALIZATION:
Cross-stack integration / Slice 1 product outcome

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh independent integration Testing session

COMMUNICATION_LANGUAGE:
Bahasa Indonesia

SELECTED_MODEL:
`gpt-5.6-terra`

REASONING_EFFORT:
`high`

MODEL_APPROVAL:
`NOT_REQUIRED`

PRIOR_ARTIFACTS:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/manifest.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TST-BE-002/testing-report-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TST-FE-001/testing-report-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TST-TOP-002/testing-report-1.md`
- `docs/project/kencleng-integration-map.md`
- `docs/product/mvp-delivery-slices.md`

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/5-testing-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

## Objective

Independently verify the real Slice-1 product experience end-to-end through the same-origin application boundary using persisted backend state and production frontend data access.

The three prerequisite milestones are already Human-promoted:
- `WU-S1-003 = BACKEND_VERIFIED`
- `WU-S1-004 = FRONTEND_MOCK_VERIFIED`
- `WU-S1-005 = TOPOLOGY_VERIFIED`

Do not re-prove their entire internal test matrices. Reuse those reports as evidence and focus on cross-work-unit correspondence plus observable product truth.

## Required integrated evidence

1. Run backend natively against isolated/test-only persisted state and the configured private MinIO path.
2. Run frontend through its production request path with MSW disabled.
3. Enter through root same-origin Caddy at `localhost:8080`; do not call frontend/backend directly for the final integrated observations except as diagnostic controls.
4. Verify a real persisted public Campaign detail end-to-end:
   - identity/purpose/steward/provenance render truthfully;
   - funding values/progress match backend projection exactly;
   - no active Donate CTA;
   - media uses the supplied controlled `content_url`;
   - media bytes arrive through the controlled backend origin with `private, no-store`.
5. Verify integrated negative/error states:
   - safe public not-found / non-public behavior;
   - missing media vs unavailable media remain distinguishable where the contract requires;
   - retryable backend/dependency failure is surfaced as the intended frontend state without internal disclosure.
6. Verify cross-stack contract correspondence:
   - frontend generated/request assumptions match actual backend wire shape;
   - same-origin `/api` path works through Caddy;
   - no mock/alternate production data branch is used.
7. Verify representative responsive integrated rendering on desktop and mobile; check no horizontal overflow and product hierarchy remains legible.
8. Reconfirm controlled-media retraction through the real frontend/root path when practical: after a previously reachable media `content_url` becomes ineligible, a fresh user-visible request must not deliver new bytes.
9. Record any difference between mock-parallel behavior and real integrated behavior. A mismatch is a finding, not something to hide with fixture changes.

## Human boundary

This Run may recommend `INTEGRATED_VERIFIED` when evidence supports it, but it must NOT self-approve:
- Human rendered/product acceptance of the integrated real experience;
- `SLICE_FINALIZED`.

Those remain Human/project gates after this Testing Run.

## Scope discipline

Do not add new product features or redesign contracts.
Do not edit production code in Testing.
If a production defect is found, create a specific Testing patch plan and route to the owning Build Work Unit.

Use Podman-aware runtime guidance from `.harscode-spaces/.local-config.yaml`.

Write:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-006/runs/TST-INT-001/testing-report-1.md`
