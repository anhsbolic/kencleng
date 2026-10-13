# Scope — RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-001

Status: PREPARED. Exact model approval pending. This is a fresh independent, read-only Review Run; no Review work has been dispatched.

## Claim and applicability

Review the exact three-file test-harness Build delta from `RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001`. The bounded claim is that Run ID and suite selection are validated before fixture creation; the active Run ID is consistently bound to the disposable fixture, test environment and ownership checks; cleanup acts only on the owned container; and the `t2` selector executes the six existing cases while the `logout` selector remains bounded. The README must describe these boundaries and evidence limits accurately.

Independent Review is applicable because this delta changes test-resource ownership and deletion controls, validation at a subprocess/resource boundary, and the selector that enables real PostgreSQL privilege and contention checks. A separate four-pass review has material value for finding mistakes in ownership validation, cleanup, secret isolation and claim wording that Builder-reported test success cannot independently rule out.

This Review does not claim production authentication behavior, independent runtime acceptance, complete T2 acceptance, T2/T3 readiness, or closure of `B-T2-002` / `B-T2-003`.

## Exact review surface

Review only the current live delta and these three files, against the exact pre/post identities recorded by the Build Run's `evidence/file-identities.json`:

- `backend/scripts/auth_harness.py` — required Run ID/suite validation, allowlisted selection, resource labels, immutable container identity and pre-provision/pre-cleanup ownership checks, environment filtering and bounded cleanup.
- `backend/internal/platform/auth/harness_test.go` — acceptance of only the active Run identity while preserving every existing assertion.
- `backend/README.md` — invocation, suite selection, security boundary and evidence claims.

Expected exact identities are recorded in the bound Build `file-identities.json`: runner before `04e8173381bb978e9ab52898969418b893672729f82b741645267aaef1e41e53`, after `bb17613e834c4416f3310d5359bc8007e4f8cace33c01f320dc4bbc053c084c2`; Go test before `c9997611413446641da20b73da2e251459a9ccbb33fbb7c88429f18206d0f45a`, after `eedb6e8007158a0900301702a1896b48abb6e494211ced91f72fa30e0037f274`; README before `38d6dbb59f19f7f936ac43932c8d3a2028e9b07585e34e5baf9844b94b41574e`, after `91f9d1c084e90f2b862bb15b67dc782458b61052d820fabc79b8d2a88b5c8046`.

Do not review unrelated dirty worktree files as part of this claim. Stop and report if any of these exact target identities or bound assignment inputs materially drift.

## Review questions and execution boundary

Perform the four canonical passes in order: Safety, Quality, Stack-Specific Best Practices, Consistency. In Safety, inspect argument validation order, shell-free argument construction, labels and owner matching, container-ID binding, environment/secret filtering, and cleanup behavior on success and failure. Confirm the six-test allowlist and README evidence statements do not overstate real versus simulated behavior. Preserve original assertions and all terminal Build/Testing evidence.

Use the Harscode best-practices routers to open only guidance relevant to input validation/injection, fixture/resource ownership, role separation or integration-test setup. State explicitly if a pass has no applicable finding. For each finding include location, problem, consequence, suggested resolution, and blocking/non-blocking status. If a source change is required, write a patch plan for a new Build/Patch Run; do not edit source here.

Review is read-only. Do not run the harness, tests, scanners, make targets, Podman/container or database commands; do not provision or clean any fixture; do not use `.env`, provider credentials, Google login, browser or HTTPS runtime. If static inspection cannot resolve a concrete question, record the uncertainty for Orchestrator routing rather than expanding this Review's execution envelope.

## Required output

Write `evidence/review-findings-1.md` with compact provenance, all four canonical passes, verification executed (`none`), scoped verdict (`Approve`, `Approve with minor comments`, or `Request changes`) and exactly one `## Phase handoff`. Create `evidence/patch-plan-1.md` only if a concrete source change is needed. An approval applies only to this bounded harness claim.

Preserve the Build report, file identities, historical AUTH-HARNESS-001 failure, logout Build/Review, and BACKEND-GATE-002 report unchanged. After terminal Review, the Orchestrator will reconcile the verdict and prepare a fresh independent Testing Run with its own approved disposable fixture and applicable post-delta command gates. No terminal Run ID or fixture ownership may be reused.
