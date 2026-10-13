# Scope — RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-FULL-ID-PATCH-001

Status: PREPARED. Narrow Build/Patch re-entry for terminal Review F1. No source/test execution has occurred in this Run.

## Claim and accepted finding

Repair only the blocking remainder of F1 from `RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-CLEANUP-PATCH-001`: before removing any fixture, cleanup must establish that the inspected immutable container ID is full and syntactically valid, and that the discovered Podman ID is a valid prefix of it. Current code accepts any nonempty inspected ID that starts with the discovered value; a short or malformed inspect ID can pass the prefix check and reach `podman rm`.

## Allowed writes

- `backend/scripts/auth_harness.py` — validate the discovered ID as 12–64 hexadecimal characters and require the inspected `Id` to be exactly 64 hexadecimal characters before any removal. Verify the discovered value is a nonempty prefix of that validated full ID. When `expected_id` is supplied from `podman run`, require it to be exactly 64 hexadecimal characters and equal the inspected ID. Any malformed, short, absent, or inconsistent identity must raise the existing sanitized unresolved-resource error and issue no removal.
- `backend/scripts/test_auth_harness.py` — retain the existing four mocked command-boundary cases and add focused short/malformed inspect-ID cases proving no `podman rm` is called. Use only mocked subprocess responses.
- This Run's own Invocation-referenced terminal patch report/evidence.

No README change is authorized or currently needed. Do not edit Go/production auth, store, DB or server code; Go tests, migration/API/dependencies, scanner policy, Make/Compose, Product/Techplan/task authority, another Run's evidence, or unrelated dirty paths. Preserve prior terminal Review/Build evidence.

## Focused Build verification

Run exactly `python3 scripts/test_auth_harness.py` from `backend/`. The test must verify existing timeout-after-create and nonzero-after-create paths still remove only the verified full ID, absent and mismatched resources remain undeleted, and short/malformed inspect IDs fail closed before removal. No real Podman/container/database/provider/browser/runtime, broader harness, gate, scanner, or network command is authorized. Focused syntax/format checking for the touched Python files is allowed.

## Report and evidence boundary

Write `evidence/patch-report-1.md` with provenance, exact delta, test scenarios/results, F1 resolution, safe diagnostics, contract check, deferred gaps, and exactly one `## Phase handoff`. The report must distinguish mocked command-boundary evidence from real resource/runtime evidence. No real container or database is created; no independent runtime acceptance is claimed.

After terminal handoff, Orchestrator re-reconciles the actual delta and re-evaluates targeted Review applicability for the safety-sensitive full-ID deletion guard. Do not treat Review chronology as automatic applicability.
