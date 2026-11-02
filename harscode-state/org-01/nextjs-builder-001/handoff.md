# ORG-01 — NextJS Builder 001 handoff

**Frontend implemented; verification PARTIAL. `MOCK_DEVELOPMENT` only.** Working-tree output on `pilot/station5-c1-space-only`, baseline `88a0c4bc1f60424a0012b1a8205cd1c69ad1b0f4`; no commit, push, backend changes, canonical API file adoption or deployment. Exact final output is pinned in `output-revisions.json` (self excluded). Human acceptance of this implementation is separate.

## Result and consumer

Engineer `/root/engineer` owns O1/O2: public creation entry, explicit attribution/internal Owner review, paired-result handling, account-bound UUIDv7 attempt recovery, own-list/detail with pagination, explicit refresh/logout handling, purpose-specific auth continuation/failure, privacy draft, full deletion scope/proof/confirmation and distinct pending/unknown/active-erased/copies states. Routes: `/`, `/organizations`, `/auth/complete`, `/auth/failure`, `/account-deletion/review`, `/privacy`.

Generated TypeScript and runtime validators consume the exact accepted r4 handoff directly for all eleven operations; no rival handwritten wire DTO or editable schema copy in `api/`. The single typed transport checks expected success/problem variants and preserves consequential uncertainty without automatic replay. MSW mocks are development-only, visibly synthetic and fail closed on startup failure. The mock auth helper bypasses Google/callback protocol; simulated success is not backend/OIDC/session/deletion assurance.

Read [frontend instructions](../../../frontend/README.md) for startup and selected checks; implementation lives in `frontend/features/org-01/` and thin App Router pages. Next consumer is Human review and a permitted local browser/runtime verification environment. Real API integration requires a separately authorized assignment and semantically verified canonical contract source; do not switch mode automatically.

## Evidence and limits

Engineer-executed Node24.22.1 checks: 23 tests across five files PASS; affected statement coverage **85.34% (361/423)**, branch82.29%, function91.02%. Denominator covers production attempt/client/workspace/deletion/provider modules and semantic privacy page; generated code, mocks and route pass-throughs excluded. Tests address paired result validation, same-key recovery, stale/private state, explicit deletion outcomes, startup/navigation guards and notice semantics. Coverage is not browser/backend assurance.

Typecheck and lint PASS, with one vendor-generated MSW ESLint warning. `npm run build -- --webpack` exited0, compiled/typechecked and generated routes. Default Turbopack build was blocked by worker port `EPERM`; offline font fetching initially failed, so frontend uses system Arial/Georgia fallbacks. Typography/render equivalence is unverified.

Runtime startup attempted with `npm run dev:mock` and `NEXT_PUBLIC_MSW_ENABLED=true npm run dev`: both failed `listen EPERM 127.0.0.1:3000`. No successful HTTP application observation. Craft Chromium launch failed sandbox permissions; Coordinator CUA observed no apps/browsers and unavailable in-app browser. Therefore actual browser, responsive layout, keyboard/focus and assistive-technology EXP08 are **BLOCKED / NOT VERIFIED**, not PASS. No privileged workaround, backend/DB activation or real user data used.

Independent Reviewer `/root/reviewer` began in a fresh context, read original authorities and did not author production code. Its own seven-case synthetic transport probe and Unicode validator probes passed; see [review](review.md) for exact assessment/finding disposition and hashes. [Craft](craft.md) records early reciprocal guidance and final source corrections; source review does not clear rendered gates. Session/context separation was observed, but native filesystem/process isolation was not established; broader workspace permissions remain a Harness limitation.

Installer reported29 dependency advisories (19moderate,9high,1critical). Bounded detailed audit failed `EAI_AGAIN registry.npmjs.org`; attribution and remediation are unresolved. No broad dependency update/audit-fix was performed. This is not a clean security-audit claim.

## Preserved gates and provenance

Binding and source-integrity details are in [state](state.md). Accepted r4 API remains SHA256 `a2d21ee3ca4ea93c7e691a53255ab825b784da00f45baeeb88115e76564df9aa`; all assignment payload hashes match the accepted packaging manifest. Existing baseline navigation changed after packaging to record canonical adoption approved/file adoption deferred; immutable inputs were not rewritten.

PDEP-R2-01 remains HELD before production/real-data use: final notice/legal applicability, provider/locations, verified privacy handling, journal/retention/all-copy disposal and restore evidence. Current draft and mock deletion do not close it. Backend security, actual authorization/cookies/CSRF/OIDC, atomicity, durable deletion, real integration, production readiness and deployment remain outside delivered evidence.

Next permitted action: inspect this revision and execute frontend HTTP/browser verification with synthetic data in an environment that permits listeners and Chromium. Address concrete visual/behavior findings in frontend, then recheck affected source. Human acceptance, real integration and release remain separate decisions.
