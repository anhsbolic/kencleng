# BLD-003 — Patch CR-001-F01 Controlled Media URL Constraint

WORK_UNIT_ID:
`WU-S1-002`

RUN_ID:
`BLD-003`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-003`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002`

ROLE:
`Implementer`

SPECIALIZATION:
None

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh session

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

## Trigger

Independent Code Review `CR-001` returned `Request changes` with one blocking finding:

`CR-001-F01` — `PublicCampaignMediaItem.content_url` only uses `format: uri-reference`, so absolute/direct bucket URLs still satisfy the schema despite the approved controlled same-origin contract.

## Prior artifacts

Authoritative current-effective Techplan:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`

Review findings:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/CR-001/review-findings-1.md`

Accepted patch plan:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/CR-001/patch-plan-1.md`

Relevant authored contract:
`api/openapi/campaign.yaml`

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/3-build-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

This is a Build/Patch re-entry requested by Code Review.

## Patch target

Implement only CR-001-F01:

1. In `api/openapi/campaign.yaml`, constrain `PublicCampaignMediaItem.content_url` with an anchored regex pattern that accepts only:
   `/api/campaigns/<UUID>/media/<UUID>/content`
2. Preserve existing type/format/description/example.
3. Reject scheme/host, query, fragment, trailing path, malformed UUID, and non-controlled path.
4. Regenerate `api/openapi.yaml`.
5. Regenerate `frontend/lib/api/generated/openapi.ts`.

Do not change route shape, field names, runtime behavior, topology/storage, product/design authority, tracker, or unrelated contract surfaces.

## Focused verification

Run the CR-001 patch-plan checks:

- `cd api && npm run validate`
- `cd api && npm run bundle`
- `cd frontend && npm run generate:api-types`
- focused assertion against the generated/bundled schema:
  - controlled path valid;
  - absolute URL invalid;
  - path without `/api` invalid;
  - query/fragment invalid;
  - malformed UUID invalid;
- independently regenerate bundle/types to a temporary location and `cmp` with committed artifacts;
- repository-root `git diff --check`;
- scoped `git status --short`.

Do not rerun `tsc --noEmit` / lint unless generated TypeScript syntax unexpectedly changes or another concrete diagnosis requires it.

## Re-review posture

If the patch stays exactly within this schema constraint + regenerated artifacts:
- full four-pass Code Review is not required;
- request targeted confirmation of `CR-001-F01` in a new Review Run.

If the patch broadens route/field/behavior/security semantics, STOP and report instead of expanding scope.

## Durable output

Write:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-003/patch-report-1.md`

Do not continue into Review or Testing in this session.
