# Scope — RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-PATCH-001

Status: PREPARED — exact-delta independent Review; Human mechanical dispatch pending. No Review execution or source action has occurred.

## Claim and reason for Review

Review the exact three-file Build/Patch delta from `RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001` for correctness and safety of the configured-pool harness oracle. Independent reasoning adds value because this delta changes fixture credential routing into a child-only `DATABASE_URL`, calls production `db.Open` from the restricted harness runtime, and generalizes Run/resource identity and cleanup behavior. Review must verify that the six selected cases and lock-expiry oracle remain the intended bound cases, without treating Builder evidence as independent Testing or acceptance.

## Exact target

- `backend/internal/platform/auth/harness_test.go`
- `backend/scripts/auth_harness.py`
- `backend/README.md`
- Build/Patch report: `../RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001/evidence/patch-report-1.md`

Review the current diff for these three files against the exact pre-delta identities in the Build Run's `input-bindings.json` and the reported post-delta identities. Do not absorb unrelated dirty-workspace changes.

## Required Review

Perform all four canonical passes, in order: Safety, Quality, Stack-Specific Best Practices, and Consistency. Apply Harscode's routed Go/Python/process/resource/security guidance relevant to this delta and cite matching sources. Read the Approved Techplan, T2 task, root/backend instructions, current auth verification environment guidance, and current diff.

Focus on: inherited environment removal; child-only fixture `DATABASE_URL`; validation that this URL targets only the owned restricted fixture; use of production `db.Open`; separation of provisioner/witness identity; test selection and preservation of all six `t2` assertions including expiry-under-lock; failure handling and cleanup ownership checks; README claim limits.

No fixture, database, provider, browser, scanner, gate, broad suite, source edit, dependency operation, or commit is authorized. Static review only; if a concrete uncertainty cannot be resolved from bound source/evidence without runtime execution, record it as a finding or explicit unresolved question rather than operating a fixture. Write only this Run's Review report under `evidence/` and a patch plan only if code changes are required.

## Output and boundary

Write `evidence/review-findings-1.md` with exactly one `## Phase handoff`. Preserve the Build report and all pre-existing dirty/untracked workspace content. The Review verdict, if Approve, applies only to this exact harness delta; it does not establish R14, B-T2-003, dependency acceptance, T2/T3 readiness, or C1 acceptance.
