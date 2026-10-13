# Scope — RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-FULL-ID-PATCH-001

Status: PREPARED. Fresh targeted independent Review of the full-ID F1 patch. Read-only; no source edits, command execution, tests, Podman, or runtime actions are authorized.

## Claim and applicability

Assess only whether the latest Build/Patch closes the terminal Review finding that cleanup could reach `podman rm` with a short or malformed inspected container ID. The claimed guard validates discovered IDs as 12–64 hexadecimal characters, requires inspected and expected IDs (when provided) to be full 64-character hexadecimal values, verifies their matching/prefix relationship before deletion, and adds mock tests that reject short/malformed inspect results without removal.

Independent Review is applicable because this is a safety-sensitive deletion guard and the preceding Review explicitly identified a concrete blocking condition. Static independent inspection of guard ordering and negative test assertions materially closes that gap beyond Builder-reported tests. This Review is limited to the latest two-file delta; it does not reassess all prior harness behavior or establish real Podman/runtime acceptance.

## Exact review surface

- `backend/scripts/auth_harness.py` — full inspected-ID syntax and relationship checks before removal.
- `backend/scripts/test_auth_harness.py` — new short/malformed inspected-ID mock cases plus their no-removal assertions.

Compare the current files against the exact pre-edit identities in the previous Review's bound `input-bindings.json` and against the terminal full-ID patch Build report. Read root/backend instructions and README, Approved Techplan/T2 task only as needed, the latest terminal F1 Review, prior accepted cleanup plan, this Build/Patch scope/report, canonical Review prompt/guidelines/checklist, orchestrated overlay/Run contract, and the matching resource-cleanup best practice.

## Review questions

1. Are all identity validations performed before any `podman rm` invocation on both ambiguous-create and normal cleanup paths?
2. Is the inspected `Id` exactly 64 hexadecimal characters, and is a discovered 12–64-character hexadecimal ID verified as its prefix? Is a supplied `expected_id` also required to be full and equal to inspect `Id`?
3. Do the new short/malformed-ID mock cases establish that removal is never invoked, while the five Builder-reported tests preserve prior cases?
4. Are errors safe, bounded, and consistent with the accepted F1 plan without widening scope?

Perform all four canonical passes in order: Safety, Quality, Stack-Specific Best Practices, Consistency. Read-only Review. No test rerun, scanner, command reproduction, Podman/container/database/provider/browser/runtime action is authorized. If static inspection cannot resolve a material concern, record it and let the Orchestrator route the remaining gap.

## Output and limits

Write `evidence/review-findings-1.md` with compact provenance, all four passes, verification executed (`none`), scoped verdict, and exactly one `## Phase handoff`. Write `evidence/patch-plan-1.md` only if a concrete source change is needed. Preserve all terminal evidence unchanged.

The terminal Build report records five passing mocked tests and no real runtime/resource operation. Treat these as Builder-reported command-boundary evidence. This Review does not establish real Podman cleanup, independent HTTP/DB acceptance, B-T2-002/B-T2-003 closure, T2/T3 readiness, or C1 completion. Any later Testing applicability must be decided from its exact claim and material evidence gaps, not phase chronology.
