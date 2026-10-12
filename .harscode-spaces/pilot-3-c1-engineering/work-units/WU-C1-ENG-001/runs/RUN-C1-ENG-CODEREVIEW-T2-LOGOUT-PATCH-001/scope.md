# RUN-C1-ENG-CODEREVIEW-T2-LOGOUT-PATCH-001 — exact patch review

## Claim and applicability

Claim: the bounded logout patch satisfies the current object-only request contract by rejecting JSON `null` with safe HTTP 400 before session revocation, while preserving valid `{}` logout and the existing Origin/CSRF/error/cookie behavior. Review is applicable because a production auth/session behavior changed at the revocation boundary; an independent Safety and contract-ordering review has material value beyond the Build report. This Run does not claim independent runtime acceptance or whole-T2 acceptance.

## Exact review surface

Review only the five-file delta from terminal Build `RUN-C1-ENG-BUILD-T2-LOGOUT-PATCH-001`:

- `backend/internal/platform/auth/http.go` — the nil-object check before `Store.RevokeSession`;
- `backend/internal/platform/auth/harness_test.go` — null/non-revocation and retained valid-object/invalid-body oracles;
- `backend/scripts/auth_harness.py` — fresh Run ownership and selected test execution;
- `backend/scripts/scanner_gate.py` — single guarded `http.go` source identity update after G124 re-triage;
- `backend/README.md` — current harness/regression/sourceguard descriptions.

Use current live files and the exact bound before/after identities. Read-only Review; do not edit production/test/policy source, rerun the fixture, run scanners, access DB/Podman/provider/browser, or replay the full backend gate. Targeted command/reproduction is outside this prepared envelope; if a concrete Review question cannot be answered from live source and durable evidence, report it for Orchestrator routing rather than expanding execution scope.

Bounded oracles: OpenAPI `/auth/logout` object-only schema and 400 `invalid_request`; live handler rejects nil before revocation; test/report preserve null session and valid `{}` revocation/cookie behavior; no cookie construction/caller changes; G124 remains the same three accepted locations with no allowance expansion. Clearly distinguish Builder-reported HTTP/PostgreSQL evidence, provider simulation, independent review reasoning, cleanup limits, and untested/deferred claims.

Once invoked, all four canonical passes are mandatory in order: Safety, Quality, Stack-Specific Best Practices, Consistency. Follow current `workflow/4-code-review-prompt.md`, guidelines/checklist, overlay, and routed Go/HTTP/security guidance. Do not inherit Implementer reasoning as proof. Preserve historical AUTH-HARNESS-001 FAILED evidence and this Build report unchanged.

## Required output and boundary

Write `RUN_PATH/evidence/review-findings-1.md` with canonical four-pass sections, verification executed (likely `none`), scoped verdict, and exactly one `## Phase handoff`. Create a patch plan only if a concrete source change is required. An Approve verdict supports only this five-file patch claim; it does not close `B-T2-002`/`B-T2-003`, establish independent runtime/T2 acceptance, or dispatch Testing/T3. Orchestrator will decide any later route from actual Review result and remaining evidence gaps.
