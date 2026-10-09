# Build Report — RUN-C1-ENG-BUILD-T1-001

> Phase: Build; Author: `PARTICIPANT-C1-ENG-CONTRACT-IMPLEMENTER-001`; Created: 2026-10-09. Model / Reasoning: `gpt-6-luna` / medium; Session: `SESSION-C1-ENG-CONTRACT-IMPLEMENTER-001`. Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T1-001`. Target revision at invocation: `524ef600c7f71246af6b71671d89c3c040fd9d44` (invocation baseline); actual observed HEAD at execution: `112c35f8322b53bac39d511ed4f22982cea619af`, branch `pilot/3-c1-engineering`, with pre-existing WU/events working-tree changes. Workflow revision: `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

Pinned input identities verified before editing: approved parent `techplan/techplan.md` SHA-256 `bd732748b486f0327f6cba2cdf14cab32f1ece3c5be8ebb24f9659444947fa47`; task `techplan/tasks/T1-shared-contract.md` SHA-256 `bbd5ab13c882ca3534ab3397ca908f70157ad41c9a81ee7cafd2827dd94946e3`; accepted manifest SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.

## What changed

- `api/openapi/index.yaml` is now the coordinated OpenAPI 3.0.3 source for the C1 boundary: Google sign-in start/callback, local `/me` and logout, establishment preparation/read/confirm, Owner-scoped list/detail, request/response schemas, and safe error meanings.
- The source preserves preparation versus effective establishment, frozen name and `c1-v1` consequence receipt, same-person replay and object scope, version conflict, expired preparation, closed/unreadable guard, safe pre-commit retry and unknown outcome, private/no-store behavior, and OrganizationView establishment-effect facts. It does not state that a provider, person, guard, or Organization was actually verified or authorized.
- `api/openapi.yaml` is the Redocly bundle. `api/openapi.d.ts` is generated TypeScript output using the already locked `openapi-typescript` dependency at 7.13.0. Final SHA-256 identities: source `8ff663e975c9d72e0d6b636994533f420fe77639398197e6137fae49e59fb62d`; bundle `9e20887011ad30a3c774edd0fb725659bccd28446fe5e4aa70c0216fbcb7a4e9`; types `a0ae33bccdf194190f3dbb49d69435d983acb8cc28fd712b3a2869442dbe6e6e`.
- `api/package.json` exposes `npm run types`; `api/README.md` documents the one editable source, derived outputs, and generation route; `api/.gitignore` excludes local dependency installation.
- No backend/frontend consumer production files or dependencies/lockfiles were changed. `npm ci` succeeded in `api/` (1 package) and `frontend/` (524 packages); frontend `package-lock.json` remained unchanged.

## Tests run

- `cd api && npm ci` → locked API tooling provisioning → passed (1 package added).
- `cd frontend && npm ci` → locked `openapi-typescript` provisioning → passed (524 packages added; lockfile unchanged).
- `cd api && npm run validate` → focused OpenAPI lint/validation → passed. Redocly reports 3 warnings: no `info.license` was established by the approved contract, and its generic `operation-2xx-response` rule warns for the two intentionally redirecting 302 operations. No contract semantics were changed to satisfy those warnings.
- `cd api && npm run bundle` → bundle generation → passed.
- `cd api && npm run types` → shared type generation via `openapi-typescript 7.13.0` → passed.
- `cd api && npm run bundle && npm run types` repeated → passed; source, bundle, and type SHA-256 values matched the values above, confirming reproducible outputs for this tool invocation.
- `git diff --check` → whitespace/conflict-marker check → passed.
- An initial `npm run validate` found a YAML inline-description parse issue; the description was corrected and the final validation passed. This is recorded as an implementation correction, not suppressed input.

## Verification scope confirmation

No race/concurrency, performance/load, or security-class test was executed in this Build iteration. No broad Testing-owned suite was executed. The generated artifact shapes were inspected for preparation-state union, OrganizationView, and the distinct 503 error codes. Validation/generation do not prove provider/session/Owner/guard/database/runtime behavior.

## Contract check

- [x] Current T1 target satisfied: coordinated source, validated bundle, and generated shared types exist under `api/`.
- [x] Live-code re-grounding did not invalidate a material contract assumption; the active API source was the empty baseline, package scripts/locks were inspected, and the locked generator was reused.
- [x] Contract authority and consumer ownership were preserved; no backend/frontend consumer implementation was changed.

## Deferred / not tested here

Real provider login, session lifecycle, database constraints/atomicity, guard operation and state, authorization/object-scope runtime behavior, transport-loss replay, browser/client integration, and Human semantic acceptance remain with their owning downstream Build, Testing, or Human phases. This contract is proposed implementation obligation, not operational evidence.

## Flagged for Techplan / Testing

No material contract gap found. Lint warnings are retained as noted above: license is unspecified, and the two redirects correctly use 302 without a fabricated 2xx response. No ignore file or lint-rule relaxation was added.

## Phase handoff

- **Outcome:** COMPLETED — this T1 Build occurrence generated and validated its assigned shared contract artifacts.
- **Result refs:** `api/openapi/index.yaml` (`8ff663e975c9d72e0d6b636994533f420fe77639398197e6137fae49e59fb62d`); `api/openapi.yaml` (`9e20887011ad30a3c774edd0fb725659bccd28446fe5e4aa70c0216fbcb7a4e9`); `api/openapi.d.ts` (`a0ae33bccdf194190f3dbb49d69435d983acb8cc28fd712b3a2869442dbe6e6e`).
- **Findings:** None material. Three Redocly warnings and their reasons are recorded in this report.
- **Decision requests:** None.
- **Blockers:** None for T1 artifact generation.
- **Open / unverified:** Independent Code Review is outstanding. Runtime/provider/session/Owner/guard behavior and Human acceptance are unverified; successful contract tooling establishes none of them.
- **Recommended continuation:** Orchestrator routes independent Code Review of T1. This handoff does not dispatch downstream work or establish dependency satisfaction.
- **Context refs:** Approved parent `techplan/techplan.md`; current task `techplan/tasks/T1-shared-contract.md`; accepted map `techplan/tasks/manifest.md`; changed contract files listed under Result refs.
