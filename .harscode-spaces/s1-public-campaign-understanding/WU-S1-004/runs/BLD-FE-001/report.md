# Build Report — Slice 1 Public Campaign Detail Frontend Mock-Parallel

- Phase: Build
- Work Unit / Run: `WU-S1-004` / `BLD-FE-001`
- Role / Specialization: Implementer / Frontend / Next.js / Public Campaign Detail
- Author: Codex CLI agent
- Model / Reasoning / Session: `gpt-5.6-terra` / high / Fresh Build session
- Created: 2026-09-23
- Target revision: `a053b48ef3fef35a073c5fe57e5ee4581f848e07`
- Workflow revision: not exposed by this Run

## What changed

- `frontend/app/campaigns/[campaignId]/` → added the public dynamic route, server route shell, route-local browser-MSW gate, TanStack Query client boundary, truthful async states, accessible retry behavior, responsive local CSS, and public detail composition.
- `frontend/lib/api/client.ts`, `public-campaign.ts`, and `lib/hooks/use-public-campaign-detail.ts` → added one real same-origin typed request path, safe `404`/`503`/generic transport classification, and a stable query key with no automatic retry, polling, Zustand, or local server-data mirror.
- `frontend/mocks/` and `frontend/public/mockServiceWorker.js` → added generated-type-checked fixtures plus shared browser/node MSW handlers for the exact detail and controlled-media paths. The worker was generated with the installed MSW CLI; the production request code has no fixture or mock-mode branch.
- Focused tests and `frontend/vitest.setup.ts` → added node-MSW lifecycle and public route/API behavioral coverage for the exact endpoint, safe errors, non-activating donation context, provenance, hostile-looking organizer text, and unavailable media.
- `frontend/README.md` → documented the opt-in mock runtime, representative routes, worker expectation, and explicit mock-versus-live integration boundary.

## Tests run

- `cd frontend && npm run test` → focused Build unit/component/API verification → passed: 3 test files, 8 tests.
- `cd frontend && npm run verify` → Build fast baseline (`lint` + Vitest) → passed: 3 test files, 8 tests. ESLint reports one non-failing warning from the unchanged generated MSW worker's unused disable directive; Vitest reports its existing future native-config advisory.
- `NEXT_PUBLIC_MSW_ENABLED=true npm run dev` plus temporary Playwright screenshots → Build-time rendered feedback, not a committed browser regression → inspected successful mock route on Desktop Chrome and a 390×844 viewport with long campaign content. The route-local worker gated the query and the rendered route preserved funding/provenance/action hierarchy and narrow reflow without observed horizontal overflow.

## Verification scope confirmation

No race/concurrency, performance/load, or security-class test was executed in this Build iteration. No broad Testing-owned suite or production build was run; `npm run build`, independent responsive/state review, media network correspondence review, and final Human rendered acceptance remain Testing/Human work per the Approved Techplan.

## Contract check

- [x] Current build target satisfied in full.
- [x] Live-code re-grounding did not invalidate a material contract assumption.

## Deferred / not tested here

- Independent Testing: production Next build, `git diff --check`, independent browser/node handler correspondence, funding edge values, full state/responsive inspection, and specialized accessibility assertions.
- Human Design: final Indonesian provenance/unavailable-action wording and any promotion of a reusable campaign placeholder treatment.
- Human rendered acceptance: representative desktop/mobile success, not-found, and retryable mock routes before any `FRONTEND_MOCK_VERIFIED` milestone claim.
- Integration owners: live backend projection, same-origin proxy, controlled-media authorization/retraction/storage, `private, no-store` cache behavior, and response/timing parity. This Build makes no integrated-verification claim.

## Flagged for Techplan / Testing

None. The generated worker lint warning and the Vitest configuration advisory are non-failing tool diagnostics, not a slice-contract deviation.

## Phase handoff

- Completed: Approved `TP-FE-001` public Campaign detail mock-parallel frontend Build.
- Artifacts: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/BLD-FE-001/report.md`
- Human decision: final Indonesian provenance/action wording and Human rendered acceptance are still required before milestone promotion; neither blocks this Build handoff.
- Open / deferred: live integration evidence and the active Human Design decisions above.
- Recommended next step: start a fresh independent Code Review.
- Session transition: start fresh Code Review for independence; use this Build report, the approved Techplan, and changed frontend files/tests only.
- Context pointers: `TP-FE-001/techplan.md`; `frontend/app/campaigns/[campaignId]/`; `frontend/lib/api/`; `frontend/lib/hooks/use-public-campaign-detail.ts`; `frontend/mocks/`; `frontend/vitest.setup.ts`; `frontend/README.md`.
