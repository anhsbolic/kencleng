# T5 — Frontend C1 establishment experience

> Phase: Techplan decomposition; Author: `PARTICIPANT-C1-ENG-DECOMPOSER-001`; Created/Updated: 2026-10-09. Model: `gpt-6-luna`; Reasoning: medium; Session: `SESSION-C1-ENG-DECOMPOSER-001` (Run binding). Target revision: `524ef600c7f71246af6b71671d89c3c040fd9d44`; workflow revision: `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

## Purpose and outcome

Implement the C1 frontend flow using generated API types and server-authoritative state: display-name form, sign-in when needed, preparation receipt and four-consequence disclosure, explicit confirmation, uncertainty resolution, my-organizations list and durable detail/navigation.

## Authority

- Parent: `../techplan.md`, SHA-256 `bd732748b486f0327f6cba2cdf14cab32f1ece3c5be8ebb24f9659444947fa47`, status Approved.
- Governing items: Q1–Q7; R2–R6; R13–R14; D5, D6, D7, D9; RISK-4, RISK-6, RISK-8; §§8, 10, 11, 12; applicable current frontend architecture and UI/UX authority.
- Product/pre-engineering authorities remain read-only. T5 does not redefine them.

## Hard dependency: T1

**Condition before API-backed consumer implementation:** T1 has published the coordinated OpenAPI source, validated bundle and generated client types for the C1 operations/shapes. Durable evidence: exact source/bundle/types revisions and recorded API validation/bundle result. T5 may prepare route-local static composition beforehand, but API calls and state handling must use the published generated contract.

## Scoped implementation

Use live anchors `frontend/app/page.tsx`, `frontend/app/layout.tsx`, `frontend/package.json`, `frontend/components/README.md`, and same-origin Caddy `/api` routing. Keep form lifecycle local and private server state in the approved server/query boundary; use the typed centralized client, cookie credentials and CSRF header on mutations. No bearer/role authority in browser storage. Preserve optional tab-local name only across auth redirect; never auto-confirm. Render all four consequences before final action; changed name requires new preparation/redisclosure. Keep confirmation disabled during preparation, before disclosure render, or while result is unresolved. Retain receipt ID through uncertainty and resolve only by same-receipt read/replay. Build distinct loading/invalid/login/expired/hold/definite failure/indeterminate/empty/success states, list → detail navigation, privacy/no-store/cache isolation and logout clearing. Follow parent §10 for accessible/responsive meaningful UI and avoid claiming review/legal verification.

## Verification and handoff

Author focused frontend checks and rendered walkthrough for four consequences, form/input, keyboard/focus, responsive layout, failure/unknown/empty/success states, durable list/detail and two-person cache isolation as assigned in parent §§11–12. MSW is contract-faithful simulation only. Real provider and backend integration are separate evidence; do not report mocks as real behavior. Route discovered contract/backend defects to their owning task, not cross-layer production edits.

## Explicit boundaries

No backend/API production edits, authorization decisions, optimistic success/role cache, production mock fallback, automatic mutation retries, mandatory checkbox, legal proof, or invented review state. Parent meaning and Human acceptance remain controlling.
