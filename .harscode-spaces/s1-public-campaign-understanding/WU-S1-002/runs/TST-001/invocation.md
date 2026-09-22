# TST-001 — Independent Testing for Slice 1 Contract Reconciliation

WORK_UNIT_ID:
`WU-S1-002`

RUN_ID:
`TST-001`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TST-001`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002`

ROLE:
`Verifier`

SPECIALIZATION:
Contract / security-boundary verification

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh independent session

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

## Preconditions

- `CR-001` completed with one blocking finding.
- `BLD-003` applied the accepted fix.
- `CR-002` independently confirmed `CR-001-F01` closed.
- No production backend/frontend runtime implementation exists in this Work Unit. Testing must verify the reconciliation contract and current observable/generated artifacts, not invent runtime evidence.

## Prior artifacts

Current-effective Approved Techplan:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`

Corrective planning amendment:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-002/amendment.md`

Latest Build/Patch evidence:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-001/report.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-002/patch-report-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-003/patch-report-1.md`

Code Review evidence:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/CR-001/review-findings-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/CR-002/review-findings-1.md`

## Canonical workflow

Use:
1. `../harscode-workspace/workflow/5-testing-prompt.md`
2. `../harscode-workspace/workflow/orchestrated-run-overlay.md`

## Testing objective

Independently verify whether the reconciliation artifacts satisfy the current-effective Techplan sufficiently for the Human-owned `CONTRACT_READY` milestone decision.

This Testing Run owns contract/document/generated-artifact verification only. It MUST NOT claim runtime backend/frontend/topology/integration behavior that this WU did not implement.

## Required focus

Follow every Testing-owned row in the Techplan Testing Checklist and execute all still-relevant Test Focus Pointer rows at the artifact/contract boundary that is actually observable in this WU.

At minimum verify:

- R1–R2: public-only operation shape, `security: []`, identical public-safe `404` contract, no `403` on touched public GETs.
- R3: all exact public projection objects remain closed with `additionalProperties: false`; no internal-model inheritance or permissive public object.
- R4–R8: plain-text organizer provenance, lifecycle/funding/action/media union semantics, decimal-string rules, absent vs unavailable truth.
- R9–R10: controlled media-content operation, constrained same-origin `content_url`, no redirect/direct-object URL contract, and `Cache-Control: private, no-store` on touched success/error responses.
- R11–R12: split OpenAPI validation, bundle generation, independent bundle/type regeneration and comparison, generated-operation/schema correspondence.
- R13: semantic consistency across Campaign task/feature/invariant/threat-model, backend architecture wording, OpenAPI, integration map, and current tracker wording.
- R14: verify milestone honesty. Do not mark the Human-owned milestone accepted; report whether evidence is sufficient for Human acceptance and whether tracker can now truthfully be updated to `CONTRACT_READY`.

### Test Focus Pointer handling

Open only the exact evidence anchors recorded in the current Techplan for:
- public projection regression;
- existence disclosure / optional auth variance;
- media revocation / cache staleness;
- money truth;
- organizer-controlled content.

Because this WU has no runtime implementation:
- verify the artifact-level obligations are preserved;
- explicitly mark runtime-only evidence as deferred to downstream backend/frontend/topology/integration Work Units;
- do not fail this reconciliation WU merely because runtime behavior does not yet exist, unless the Techplan wrongly claims runtime proof is required now.

If Testing discovers that a supposedly deferred runtime property is actually necessary to justify `CONTRACT_READY`, report Techplan/authority drift instead of silently broadening scope.

## Target-repo final verification

Use repository-owned commands and current documented mechanics. Expected relevant commands include:
- OpenAPI validation/bundle;
- frontend generated-type regeneration;
- TypeScript check;
- lint where required by current repo authority;
- independent generated-artifact reproducibility;
- diff/status hygiene.

Run only the broad checks justified by the current Techplan/repo authority. Do not run browser/UI/runtime backend suites that cannot exercise anything introduced by this reconciliation-only WU.

## Non-blocking prior comment

`CR-001-C01` — stale threat-model reference labels — remains a known non-blocking comment. Testing may mention it as follow-up but must not silently promote it to a blocking failure unless current evidence shows it causes semantic ambiguity or incorrect authority routing.

## Human-owned gate

The Techplan assigns R14 milestone acceptance to Human.

Testing must conclude one of:
- evidence sufficient for Human to accept/update `CONTRACT_READY`;
- evidence insufficient, with specific blocking reason and patch plan if code/artifact changes are required.

Testing itself does not update Orchestrator-owned manifest/control surface.

## Durable output

Write:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TST-001/testing-report-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TST-001/patch-plan-1.md` only if changes are required.

Do not continue into PR/finalization or downstream backend/frontend/topology Work Units in this session.
