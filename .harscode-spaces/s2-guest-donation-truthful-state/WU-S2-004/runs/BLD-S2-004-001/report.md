Phase: Build
Author: P-S2-004-BLD-001-1 (Implementer / KC-IMPLEMENTER)
Created: 2026-10-03
Model / Reasoning / Session: Invocation configured gpt-6-luna / high; active runtime values and Session identifier are not independently exposed
Target revision: bb69cd002b3f1a1056837affcd77bb2b001007b0 plus current working tree
Workflow revision: Harscode pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4 (current-effective; not semantically pinned)
Work Unit / Run: WU-S2-004 / BLD-S2-004-001

## What changed

- frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx and campaign-detail.module.css → disclose max_donation_amount with explicit IDR before amount entry, including when Funding is unavailable. The donation CTA appears only for the contract's donation_action.available state. The cap is labelled as a per-donation limit and is not described as remaining capacity or an admission guarantee.
- frontend/app/campaigns/[campaignId]/donate/ → added the guest form. It re-reads Campaign detail and does not show the form when the current action is unavailable. It validates whole-IDR input from Rp5.000 without imposing a client maximum, includes optional name and explicit email opt-in with the accepted label/helper, and shows QRIS as the only active sandbox method. GoPay, ShopeePay, and bank transfer are visible as unavailable text.
- frontend/lib/api/donation.ts and frontend/lib/api/client.ts → added typed generated-contract submit/status calls. Submission sends Idempotency-Key; status lookup sends X-Donation-Status-Credential with cache: no-store. 422 amount feedback, 409 request feedback, ordinary request failure, and ambiguous transport/5xx outcomes have distinct safe handling. An ambiguous retry keeps the original payload and key and locks editing until retry.
- frontend/app/donations/[donationId]/status/ → added a status-only view. It reads the one-time URL fragment, removes it from the visible URL, sends the value in the designated header, and displays pending or backend-returned terminal simulation wording. Pending status is checked again only on explicit action; no timer or resubmission is introduced. Missing/invalid status uses the generic public failure copy.
- frontend/mocks/handlers/donation.ts, frontend/mocks/browser.ts, and frontend/mocks/server.ts → added contract-shaped MSW submit/status handlers at the network boundary, with same-key replay behavior and generic status lookup failure. Production adapters continue to call the real /api endpoints; no mock branch was added to production API code.
- frontend/app/mock-service-worker.tsx → made MSW startup available to both the existing Campaign route and Donation routes. The donation form is nested at /campaigns/{campaignId}/donate to follow the live App Router structure and avoid conflicting sibling dynamic segment names.
- frontend/app/campaigns/[campaignId]/campaign-detail-client.test.tsx and frontend/app/donations/donation-flow.test.tsx → added observable coverage for cap/action behavior, funding-unavailable disclosure, stale unavailable action, minimum and Rp5.001 amounts, method availability, opt-in payload, 422/409 separation, same-key ambiguous retry, fragment cleanup/header handoff, status-only display, terminal labels, and generic missing-link behavior.

No API/spec/Product/design/generated type, backend, tracker, manifest, Work Graph, or other Work Unit file was changed. Approved Techplan content remained unchanged.

## Tests run

| Command | Category | Result |
|---|---|---|
| cd frontend && npm run test -- app/donations/donation-flow.test.tsx app/campaigns/'[campaignId]'/campaign-detail-client.test.tsx | Focused Vitest/RTL/MSW Build-loop verification for changed Campaign and Donation flows | Passed: 2 files, 17 tests. Covers the newly authored observable tests and the changed Campaign detail consumer. |
| cd frontend && npm run lint | Frontend lint | Passed with 0 errors. ESLint reports one existing warning in frontend/public/mockServiceWorker.js: unused eslint-disable directive; this generated worker was not changed. |
| cd frontend && npm run build (rerun with approved network access) | Next.js production compilation, TypeScript, and App Router route validation | Passed. Next compiled successfully, finished TypeScript, generated static pages, and listed /campaigns/[campaignId]/donate and /donations/[donationId]/status. |
| Initial sandbox npm run build attempt | Next.js production build | Stalled without further output after entering optimized compilation; the session was interrupted after over one minute. The same command was rerun with network access because the configured next/font/google build path plausibly required a font fetch; that rerun passed as above. |
| git diff --check | Patch whitespace check | Passed. |

No OpenAPI generation/validation or browser automation was run. The generated types were inspected and consumed as-is.

## Verification scope confirmation

No race/concurrency, performance/load, or security-class test was executed in this Build iteration. No broad Testing-owned suite was run. The production build was included to verify App Router route composition and TypeScript after adding the new routes; it is not independent Testing evidence.

## Contract check

- [x] The approved frontend Build target was implemented in full within the frontend write envelope.
- [x] Live-code re-grounding did not invalidate a material product, authority, API, architecture, or verification assumption. The target HEAD, approved Techplan hash, and generated Campaign/Donation type anchors matched the Invocation coordinates at dispatch.

## Deferred / not tested here

- Independent Testing owns the focused real-browser check for URL-fragment handoff, visible URL cleanup, status request header, and status route (Techplan R8). jsdom coverage does not establish browser history/address behavior.
- Human rendered acceptance remains required for representative responsive hierarchy, amount/cap comprehension, simulation wording, and accessible interaction (R10/O7). No rendered acceptance or browser session is claimed here.
- Backend DTO/wire parity, eligibility/capacity enforcement, idempotency persistence/concurrency, simulator outcomes, credential generation/verification/expiry, privacy/infrastructure exposure, anti-enumeration parity, email verification/delivery, and real backend integration remain outside this frontend Build and retain their Techplan owners (WU-S2-003 and Security/PII/API paths).
- docs/project/kencleng-integration-map.md currently contains the Public Campaign Detail mapping but no Donation-flow row. Updating shared coordination docs is outside this Run's execution envelope; Orchestration should reconcile that mapping before integrated delivery.
- Existing orchestration working-tree changes outside frontend/ and this Run path predated dispatch and were left untouched.

## Flagged for Techplan / Testing

No material assumption break or new security control requirement was discovered. Preserve R8 browser evidence, R10 Human rendered acceptance, and the cross-stack integration-map follow-up above in downstream routing. Passing mock tests and the production build establish frontend implementation/build evidence only; they do not establish FRONTEND_MOCK_VERIFIED, runtime readiness, accepted residual risk, or WU completion.

## Risk note

- Assumptions made: accepted generated contracts remain authoritative; current Campaign donation_action is only an entry snapshot; nested /campaigns/{id}/donate is the route-local App Router placement.
- Edge cases intentionally NOT handled (and why): pending status is not automatically polled and never triggers a submission retry; status expiry/credential validity and unknown capacity reasons remain backend-owned. Status lookup failures collapse to the accepted generic public behavior.
- Concurrency assumptions: the client disables submission during a request and preserves the same key and payload for ambiguous retries. Durable idempotency, ordering, and duplicate financial-intent protection are backend responsibilities and were not proven by MSW.
- What is not tested, and why: real-browser fragment/history behavior is assigned to independent Testing; rendered acceptance is Human-owned; runtime/security, backend integration, concurrency, and email controls are outside this frontend Build.

## Phase handoff

- **Outcome:** COMPLETED — this initial Build Run implemented the Approved frontend Techplan and completed focused Build-loop checks. This is Run-occurrence status only.
- **Result refs:** This report; implementation under frontend/app/campaigns/[campaignId]/, frontend/app/donations/, frontend/lib/api/, and frontend/mocks/; approved TP-S2-004-003/techplan.md SHA-256 e6a93bc65a760084955f63a469b238fd70a80c20ff38cc1d22ba1be000ad99eb.
- **Findings:** No material contract contradiction. Route placement was derived from the live App Router dynamic-segment convention. The current Integration Map lacks a Donation-flow row; this is a cross-stack coordination follow-up, not a reason to change contract behavior in this frontend Run.
- **Decision requests:** None.
- **Blockers:** None for this bounded Build target.
- **Open / unverified:** Independent R8 browser evidence; R10 Human rendered acceptance; backend/API/Security/PII runtime and integration evidence; integration-map update; no race, performance, security-class, or broad independent Testing was performed.
- **Recommended continuation:** Independent Code Review of this Run's frontend diff and report. Subsequent Testing and Human rendered acceptance remain separate gates; Orchestrator determines routing.
- **Context refs:** This Run's invocation.md; approved TP-S2-004-003/techplan.md; changed frontend files listed above; frontend/AGENTS.md; api/openapi/campaign.yaml, api/openapi/donation.yaml, and referenced common.yaml validation components.
