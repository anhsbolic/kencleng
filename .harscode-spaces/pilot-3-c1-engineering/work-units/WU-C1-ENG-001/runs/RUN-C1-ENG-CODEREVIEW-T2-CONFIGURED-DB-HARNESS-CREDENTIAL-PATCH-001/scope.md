# Scope — RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001

Status: PREPARED — targeted independent Review confirmation; Human mechanical dispatch pending. No Review command or source action has occurred.

## Claim and exact target

Confirm closure of blocking F1 from `RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-PATCH-001`: fixture credentials and Run identity must not reach the unrelated `go version` subprocess, while the intended Go test child retains its required fixture environment.

Review only the new Build/Patch delta in:

- `backend/scripts/auth_harness.py`
- `backend/scripts/test_auth_harness.py`

Build/Patch report: `../RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001/evidence/patch-report-1.md`. Compare the postimages to the exact pre-patch identities in that Build Run's `input-bindings.json`. Do not include unrelated dirty-workspace changes or re-review the prior three-file harness delta outside what is needed to assess F1 closure.

## Required Review

Perform all four canonical passes in order: Safety, Quality, Stack-Specific Best Practices, and Consistency. Apply them to this narrow delta, with primary attention to whether all fixture secrets/DSNs and Run identity are removed from the metadata subprocess environment, whether required nonsecret environment values remain, and whether the helper leaves the Go test-child environment unchanged. Confirm the regression test exercises those boundaries. Read current Review guidance, the prior F1 report/patch plan, this Build report, the Approved Techplan/T2 task, root/backend instructions, and only relevant routed best-practice sources.

Static Review only. No test, gate, scanner, fixture, DB, Provider, browser/runtime operation, or source edit. If execution is required to resolve a finding, state the uncertainty for the owning phase rather than executing it here.

## Output and limits

Write only `evidence/review-findings-1.md` with exactly one `## Phase handoff`. Verdict applies to F1 and this exact two-file patch. It does not establish R14, B-T2-003, B-T2-002 acceptance, dependency commit, T2/T3 readiness, or C1 acceptance. After terminal Review, Orchestrator rechecks or supersedes the staged Testing assignment before any fixture authorization or dispatch.
