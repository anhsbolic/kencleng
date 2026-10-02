# Stage 2 — Gap Analysis

> Phase/Stage: Exploration / Stage 2  
> Author: Explorer (`P-S2-004-EXP-001-1`)  
> Created: 2026-10-01  
> Model / Reasoning: Invocation-selected `gpt-6-luna` / `high` (runtime model metadata not independently exposed in this session)  
> Session: interactive dispatch session; session identifier not exposed  
> Target revision: invocation baseline `7fd8b473b239b20bda3990ab29c51440d321a796`; relevant sources re-read live at dispatch  
> Workflow revision: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`

## Area 1 — Product, Donation domain, and approved delivery boundary

### Current state

- Parent Outcome and Work Graph record `CONTRACT_READY` for WU-S2-002 and the parallel delivery frontier: WU-S2-003 owns backend Donation capability; WU-S2-004 owns the frontend guest flow against contract-faithful mocks. WU-S2-004 is active/queued for `FRONTEND_MOCK_VERIFIED`; real integration and Slice finalization are separate.
- The current approved user outcome is a public visitor contributing without an Account and understanding the actual sandbox processing/result state. Product authority defines the whole guest path and explicitly excludes Account prerequisite, claim/history, active non-QRIS processing, real rails, and broader Slice 3 closure/result behavior. Sources: `docs/product/mvp-scope.md` §§4–5, 7; `docs/product/mvp-delivery-slices.md` §5; WU-S2-004 `manifest.md` Outcome/Scope/Boundaries.
- Accepted Donation authority specifies: whole IDR input (minimum Rp5.000; Rp5.001 valid); the four familiar method labels with QRIS as the only interactive simulation; backend-owned `pending`/`success`/`failed`; pending text “Menunggu hasil simulasi”; deliberate fresh donation after failure; idempotent ambiguous retry; optional non-public name and opt-in status email; and a 24-hour status-only guest link. Accepted domain sources: `docs/spec/5-donation/invariants.md` INV-donation-01–11, `tasks.md` Task 01/02, `features/01-submit-donation-settlement.md`, `features/02-donation-status-check.md`.
- Accepted credential direction is fragment-carried URL with frontend handoff and URL cleanup, one-way HMAC verifier, hard 24-hour expiry, and uniform public `404` with matching Problem Details body, headers, and cache behavior (`private, no-store`). The browser/infrastructure exposure controls, credential strength/key/comparison/lifecycle/abuse controls, empirical parity, and residual-risk decision remain open O4/O5 evidence/gates.
- Guest email is optional, opt-in, status-only. Ownership verification precedes status/access delivery; at most one terminal simulation-labeled notice is due. Security/PII windows and controls, verified-email bounded/recoverable terminalization details, and retention/deletion-race evidence remain open O3 delivery/testing obligations. No residual risk is accepted by the spec acceptance status.
- Donation references the Campaign eligibility/threshold boundary only: new submission eligibility is ordered against close; accepted pending donations can settle in full; successful settlement and funding are exact-once; later settlement does not reopen the campaign or change its winning close reason. This does not pull broader Slice 3 behavior into scope.

### Requirement and gap

The Product/MVP and accepted Donation authority establish a coherent frontend outcome and its exclusions. Product meaning for this area is sufficiently stated. The current frontend implementation has not yet been inspected, so implementation conformance and concrete UI gaps are **not assessed in this area**; those are recorded under the frontend area below. Delivery must not turn the still-open O2–O5 controls into client-side claims or behavior.

### Sniffing

- **Risk:** A client can mislead donors if it depicts a pending donation as settled, initiates a new donation after an ambiguous result, or exposes private status beyond the intended status-only view. Browser handling of the fragment credential is a material privacy boundary; O4/O5 implementation controls and runtime proof remain open.
- **Edge cases:** Amount Rp5.001 is valid; pending must remain pending without a timer/promise or resubmission; failed permits only an explicit new intent/key; expired/missing/wrong credential and absent Donation converge on the same public `404`; accepted pending donations remain settleable after Campaign close.
- **Miscontext:** Historical Donation behavior, old timing/probability assumptions, `401` status behavior, non-expiring credentials, and account-first flows are not current Slice 2 authority. Product/MVP and accepted reconciliation take precedence.
- **Misleading signals:** A displayed payment-method label does not imply that method is usable; only QRIS runs the labeled sandbox simulation. An email address entered is not verified, and a terminal notice obligation is not proof that a notice was delivered.
- **Inconsistency:** No contradiction was found within the current Product/MVP and accepted Donation requirements inspected for this area. The open O2–O5 controls/evidence are explicitly downstream and are not treated as accepted risk.

### Source anchors

- `docs/product/mvp-scope.md` §§4–5, 7 — guest actor, guest donation, status truth, and security floor.
- `docs/product/mvp-delivery-slices.md` §5 — Slice 2 inclusion, method/status/email/credential rules, threshold boundary, exclusions, and completion evidence.
- `docs/spec/5-donation/invariants.md` INV-donation-01–11 — active donation invariants.
- `docs/spec/5-donation/tasks.md` Task 01/02 — accepted guest submission and status revisit requirements.
- `docs/spec/5-donation/features/01-submit-donation-settlement.md` and `02-donation-status-check.md` — feature acceptance and unresolved controls.
- `.harscode-spaces/s2-guest-donation-truthful-state/outcome.md`, `work-graph.md`, and `WU-S2-004/manifest.md` — current delivery ownership, dependency, and frontend mock boundary.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` — approved contract-reconciliation direction and downstream evidence obligations.

## Area 2 — Authored API contract

### Current state

- `api/README.md` establishes split domain OpenAPI as authored authority, `common.yaml` as shared-component source, and `api/openapi.yaml` as generated bundle. For this Donation scope, the active source is `api/openapi/donation.yaml` plus the referenced shared schemas/parameters in `api/openapi/common.yaml`.
- `POST /campaigns/{campaignId}/donations` accepts a guest submission, requires the shared `Idempotency-Key` UUID header, accepts only `qris`, and returns a persisted `pending` Donation with a submission-only `status_token`. The request carries major-unit decimal `amount` and explicit `currency_code: IDR`, optional name/email, and explicit status-email opt-in. The shared idempotency description requires the same key and payload to resolve to the original Donation and changed-payload reuse to reject.
- `GET /donations/{donationId}/status` accepts optional `X-Donation-Status-Credential`; its contract documents the frontend handoff from URL fragment to request header and visible URL cleanup. A successful response is `DonationStatusResponse` with `status` only. Missing, wrong, and expired credentials and absent Donation are specified as one public `404` with a shared Problem Details shape and `Cache-Control: private, no-store`.
- The API also contains historical public donor-list and account donation/claim operations. These are not part of the active guest path by current Product/MVP and Donation specs; presence in the API source does not make them in-scope.
- OpenAPI explicitly leaves exact credential generation/strength, key handling/comparison, expiration implementation, browser/referrer/log/cache protections, abuse controls, and runtime parity evidence to O4/O5 owners. Email verification/retention/retry/deletion controls remain O3. These are unresolved delivery/security evidence obligations, not silently established by the contract.

### Requirement and gap

For Donation submission/status, current authored shapes express the required guest path, including the agreed idempotency header, fragment-to-header credential handoff, status-only projection, and uniform public failure contract. No Donation-shape mismatch against the inspected accepted requirements was found. The separate Campaign Detail action contract is the material cross-domain mismatch recorded below. Whether frontend code consumes the Donation shapes correctly, and whether live behavior satisfies the documented exposure/parity properties, remains unassessed until the frontend and its mock boundary are inspected; mocks cannot prove backend security/runtime properties.

### Sniffing

- **Risk:** The submission response contains the bearer `status_token`; the status endpoint is unauthenticated when the bearer credential is supplied. A frontend leak through URL/history/referrer/logs/cache or accidental persistence would expose a private status capability. API contract direction alone does not prove protection.
- **Edge cases:** A missing credential is permitted on the wire so the server can return the generic `404`; wrong, expired, and absent-record cases must match the same status/body/headers/cache contract. A same-key changed-payload retry conflicts while same-key/same-payload retry refers to the same Donation.
- **Miscontext:** The Donation YAML still defines public donor-list and account claim/history operations, but the active Product/spec boundary defers those. The frontend task must not treat them as implicit additions to Slice 2.
- **Misleading signals:** A generated type for `status_token` or a status endpoint declaration only proves contract shape. It does not prove URL cleanup, secure header handoff, idempotent persistence, response parity, or production runtime control. MSW fidelity also cannot establish backend behavior.
- **Inconsistency:** No material shape conflict was found between the inspected authored Donation contract and accepted Product/domain behavior. The contract itself marks implementation controls and runtime parity as open; therefore the absence of empirical proof is a known verification gap, not an authority contradiction. Invocation lists TPD-S2-002-002 artifacts, but the exact top-level `manifest.md` and named Task 02 snapshot paths are absent under that run directory; actual snapshots are under `tasks/` with different filenames. Current direct authorities remain discoverable, so this is a pointer/discoverability issue only.

### Code / contract anchors

- `api/openapi/donation.yaml#/paths/~1campaigns~1{campaignId}~1donations/post` — submission method, idempotency, sandbox semantics, and success/error response.
- `api/openapi/donation.yaml#/paths/~1donations~1{donationId}~1status/get` — credential handoff, status projection, and public `404` contract.
- `api/openapi/donation.yaml#/components/schemas/SubmitDonationRequest` and `Donation` — submitted guest fields and submission credential response.
- `api/openapi/donation.yaml#/components/schemas/DonationStatusResponse` — status-only lookup projection.
- `api/openapi/common.yaml#/components/parameters/IdempotencyKeyHeader` and `#/components/schemas/Problem` — shared retry key and Problem Details envelope.
- `api/README.md` — authored split-source vs generated bundle routing.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002/tasks/manifest.md`, `01-donation-domain-spec-reconciliation.md`, and `02-donation-openapi-reconciliation.md` — actual decomposition artifacts discoverable at dispatch; Invocation's cited paths/names differ.

### Cross-domain Campaign-to-Donation contract finding

The active Product requirement in Slice 2 is to move from an eligible Public Campaign Detail into the guest Donation flow (`docs/product/mvp-delivery-slices.md` §5, Completion evidence). The current Public Campaign Detail contract remains the Slice-1 shape: `api/openapi/campaign.yaml` `PublicCampaignDonationAction` allows only `availability: unavailable` and `reason: donation_flow_not_available`; `docs/spec/4-campaign/features/02-campaign-detail-listing.md` §1/acceptance and `docs/spec/4-campaign/invariants.md` INV-campaign-14 specify the same unavailable action for the Slice-1 contract. Current frontend fixtures match that contract. The Donation API itself does not provide a campaign-detail action/eligibility projection.

This is a material lower-level contract/spec gap against current Product/MVP truth, not a missing Product decision: Product/MVP precedence means the old Slice-1 contract must be reconciled for the new Slice-2 requirement. Frontend cannot silently hard-code a CTA or infer donation eligibility from visibility/public `fundraising` state in place of the current API field; backend remains authoritative for submission eligibility. Cross-stack contract change is outside the frontend-only Build boundary unless explicitly coordinated. This blocks implementing the required Campaign Detail → guest Donation entry path against current accepted contract; the remaining Donation flow can still be explored/planned while the API coordination route is resolved.

Progression effect: **active scoped blocker for the Campaign Detail entry path; Stage 3 may safely proceed to identify and route the contract reconciliation/coordination need without changing authority or contract.** The owner/action should be established in Stage 3/Techplan orchestration before Build commits to the route boundary.

## Area 3 — Product Design and surface authority

### Current state

- `docs/ui-ux/README.md` routes current design concerns and identifies Sunlit Editorial / Evidence-Led Optimism as the active direction. Product/domain truth outranks design material for business meaning.
- `docs/ui-ux/page-map.md` §1 identifies the Guest / Public Visitor surfaces: Public Campaign Detail with a truthful donation action, Donation flow as a form, and Donation status/tracking as a status surface. It does not prescribe route names or frontend architecture.
- The applicable interaction authority is present: `patterns.md` §4 defines form context/fields/guidance/validation/submit structure, field-vs-request failure distinction, accidental duplicate prevention, and preserved input; §7 defines status/tracking as minimal context → current state → meaning/consequence → relevant next action and requires anti-enumeration-compatible visible failures; §§12–15 govern error recovery, completion, status semantics, and explicit money labels.
- Design principles emphasize confidence before conversion, no visual implication without product truth, explicit financial consequences, visibility of unknown/pending states, one clear next action, and responsive preservation of trust/consequence information. The visual guidance says color cannot stand alone for state; pending is neutral/plainly stated, and money/security errors should not be playful.
- Slice-2-specific language choices are accepted in Product/spec and API sources: pending “Menunggu hasil simulasi”; terminal label family “Hasil simulasi donasi: berhasil/gagal”; optional email label/helper in Feature 01; unavailable methods and generic link failure follow current authority. Final rendering/placement and accessibility must still be inspected in actual UI.
- Design readiness per `product-design-principles.md` §14 is not a blanket mockup gate. Ordinary presentation details can follow established patterns; material new interaction, hierarchy, navigation, or brand decisions require low-fidelity alternatives and Human direction. The specific page/component readiness cannot yet be classified until the live route/component structure is inspected.

### Requirement and gap

Design authority supplies the applicable surface intent, language, state hierarchy, responsive/accessibility principles, and ordinary interaction patterns. No high-fidelity Figma artifact is established as a prerequisite by these sources. A concrete readiness classification and identification of any material design gap remain pending the frontend inspection; no product, privacy, money, security, or status meaning may be inferred from visual treatment.

### Sniffing

- **Risk:** Styling a successful simulation like provider settlement or using color/checkmarks as verification could create a false trust signal. Exposing distinct lookup failures would conflict with the anti-enumeration direction. Hiding pending state, email consequences, or unavailable-method status on smaller screens would make the flow misleading.
- **Edge cases:** Pending, failed, successful, unavailable method, request failure, and invalid/expired/missing status link have distinct domain consequences. The public status-link failure must not reveal which credential/Donation condition failed. Responsive layout must preserve trust and action order.
- **Miscontext:** The approved donor post-donation reference is an Evidence Journal/product precedent, not proof that Slice 2 includes the broader Evidence Journal or Campaign closure/result behavior. Page map owns surface purpose, not exact route/component implementation.
- **Misleading signals:** Brand Sun/green/accent treatments are not proof of success, trust, or verification. A status badge alone is insufficient for consequential state. Friendly copy cannot change backend-owned Donation semantics.
- **Inconsistency:** The current UI/UX routing identifies `design-guidelines.md` as the concrete visual authority. Brand brief retains intentionally open broad topics (for example, final logo/motion/provenance terminology), while current guidelines establish selected rules for relevant status, color, responsive, and accessibility behavior. No open brand decision is evidenced as a prerequisite for this flow; verify against actual component needs rather than importing unrelated OPEN items.

### Source anchors

- `docs/ui-ux/README.md` — active design direction and authority routing.
- `docs/ui-ux/page-map.md` §1 and §§7–9 — Guest surfaces, sensitive-data presentation, and page-map boundaries.
- `docs/ui-ux/product-design-principles.md` §§1–5, 8–9, 13–15 — trust, money, unknowns, next action, responsive order, readiness and decision boundaries.
- `docs/ui-ux/patterns.md` §§4, 7, 12–15 — form, status/tracking, error recovery, status communication, money presentation.
- `docs/ui-ux/design-guidelines.md` §§4, 13, 21–24, 26 — semantic colors, truth grammar, responsive/accessibility rules, open visual decisions.
- `docs/ui-ux/brand-product-ui-brief.md` §§5–6 — action hierarchy, status/data, trust cues, error and responsive behavior.
- `docs/spec/5-donation/features/01-submit-donation-settlement.md` “Exact Design direction for dependent UI/notice” — accepted Slice-2 labels/helper copy.
## Area 4 — Frontend architecture and live implementation

### Current state

- Frontend is the post-reboot Next.js App Router clean-start generation. Current tree has a public home and one implemented route, `app/campaigns/[campaignId]`, for Public Campaign Detail. The route uses a Client Component and TanStack Query for GET detail loading; local route states cover loading, not-found, temporary unavailability, request failure/retry, and success.
- `CampaignDetailView` renders an explicit message that the donation flow is unavailable and exposes no donation link/button. This is consistent with the current Slice-1 `PublicCampaignDonationAction` generated type/fixture, whose only allowed state is unavailable. The campaign fixture set includes `donation_action: unavailable` for all cases.
- The application API adapter and hook currently expose only `GET /api/campaigns/{campaignId}`. The MSW handler covers Campaign detail/media only; no guest-submit or Donation-status request function, route, screen, form, or Donation MSW handler is present. `rg` found Donation endpoint/types in generated OpenAPI, but no runtime caller outside those generated declarations. No Donation-specific frontend route exists in the live file list.
- `frontend/lib/api/generated/openapi.ts` already reflects Donation submit/status operations and schemas, including `SubmitDonationRequest`, `Donation.status_token`, `DonationStatusResponse`, and the `X-Donation-Status-Credential` request parameter. Generated type presence is contract correspondence only; it is not working frontend behavior.
- The scaffold uses `apiRequest(path)` with GET-only behavior and a stable generic transport error. Campaign MSW is started by the campaign route layout only when `NEXT_PUBLIC_MSW_ENABLED=true`; the API client itself has no fixture/mode branch. Existing tests cover Campaign GET/mocks, retry, safe error surfaces, focus restoration, and literal hostile organizer text; they assert the current non-activating Donation context. Tests were inspected, not run.
- Frontend architecture permits route-local composition and requires generated API types, network-boundary MSW for contract-parallel mocks, React Hook Form/Zod when a form is needed, server-owned eligibility/state, sensitive values kept out of convenient URL persistence, and accessible rendered/browser verification. The shared `ui/` and `shared/` contract registry is intentionally empty; no Donation-specific primitive/component is established.

### Requirement and gap

The required end-to-end frontend flow is not implemented yet: currently there is no actionable Campaign Detail entry, no guest form/submission flow, no pending/terminal outcome view, no temporary status revisit, and no Donation network mock coverage. This is expected at the queued frontend delivery frontier, but it is the concrete implementation gap WU-S2-004 must address. The Campaign Detail entry has the additional current-contract blocker recorded above. The accepted Donation API types exist and match the request/status contract; no frontend consumer validates the integration yet.

Current design readiness is **PARTIAL** for the form/status presentation: intent, key state meanings/copy, form/status patterns, and visual principles are established, while ordinary composition can be derived during implementation. It is not a code-ready end-to-end flow until the Campaign action contract is coordinated. No material brand/design ambiguity is evidenced yet. Design acceptance remains Human-owned after rendered implementation.

### Sniffing

- **Risk:** A hard-coded CTA or client-derived eligibility could misrepresent whether a campaign accepts donations; the backend must reject ineligible submissions. A donor double-click, ambiguous timeout, or fresh key on retry could create unintended donations. Mishandling `status_token` in the URL/history/logs or showing details beyond status would expose a bearer capability. MSW cannot prove server security, funding atomicity, or real simulator state.
- **Edge cases:** The public detail response can become stale before submission; the API may return `409` for ineligibility or conflicting key payload. Submission may time out ambiguously, and client state must not create a new key automatically. A pending result may outlive the page; link missing/wrong/expired and absent Donation share public `404`; the fragment credential must be handed to the request and removed from visible URL. Optional email stays opt-in, and its verification/delivery lifecycle has unresolved O3 controls.
- **Miscontext:** Current active code is a clean-start Slice-1 frontend, not a partially implemented Donation flow. The presence of generated Donation types is easy to mistake for usable data access. Old Git-history frontend components are explicitly not current precedent.
- **Misleading signals:** The campaign route has extensive state/retry handling, but that only covers campaign detail GET; it does not imply donation eligibility, submission, or status lookup. Current CTA text expressly says donations are unavailable, while Product/MVP now requires the Slice-2 entry path for eligible campaigns. Existing MSW tests prove the old mock-parallel Campaign route only.
- **Inconsistency:** Product Slice 2 requires Campaign Detail → Donation for eligible campaigns, while current Campaign contract/spec and implementation explicitly mark the donation action unavailable for all public detail responses. The Product/MVP source owns this behavior; lower-level Campaign contract/spec and frontend must be reconciled in coordinated scope. No component architecture conflict exists: route-local flow can be built without assuming a broad shared component contract.

### Code anchors

- `frontend/app/campaigns/[campaignId]/page.tsx` — current dynamic public detail route entry.
- `frontend/app/campaigns/[campaignId]/campaign-detail-client.tsx` — query-owned detail loading and retry/focus transitions.
- `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` `CampaignSuccess` action aside — current non-activating donation copy/no CTA; target entry-path surface.
- `frontend/app/campaigns/[campaignId]/campaign-detail-client.test.tsx` — current test explicitly verifies donation action remains non-activating.
- `frontend/lib/api/public-campaign.ts` `getPublicCampaignDetail` and `frontend/lib/hooks/use-public-campaign-detail.ts` — sole current public detail data path.
- `frontend/lib/api/client.ts` `apiRequest` — current GET-only network adapter; transport errors are generic.
- `frontend/mocks/handlers/public-campaign.ts` and `frontend/mocks/fixtures/public-campaign.ts` — current Campaign-only MSW boundary and unavailable donation action fixture.
- `frontend/app/campaigns/[campaignId]/layout.tsx` and `mock-service-worker.tsx` — route-local MSW startup gate/lifecycle.
- `frontend/lib/api/generated/openapi.ts` `paths["/campaigns/{campaignId}/donations"]`, `paths["/donations/{donationId}/status"]`, `schemas.SubmitDonationRequest`, `schemas.Donation`, `schemas.DonationStatusResponse`, and `schemas.PublicCampaignDonationAction` — generated contract coordinates; Donation operations have no active caller, and Campaign action is currently unavailable-only.
- `docs/spec/4-campaign/features/02-campaign-detail-listing.md` §1/acceptance, `docs/spec/4-campaign/invariants.md` INV-campaign-14, and `api/openapi/campaign.yaml` `PublicCampaignDonationAction` — existing Slice-1 contract/spec authority that conflicts with current Slice-2 Product requirement.
- `docs/project/kencleng-frontend-tech-stack.md` §§4–5, 7–10, 12, 15–20; `frontend/AGENTS.md` §§2–3, 6, 10–12; `frontend/components/README.md` §§2–5, 10–12 — state ownership, network/mock boundary, component ownership, rendered acceptance, and security presentation.

## Stage-2 progression summary

### Material findings

| Finding | Progression label | Effect and scope | Can Stage 3 proceed? | Next action / owner |
|---|---|---|---|---|
| **F-1 — Campaign entry action contract conflicts with Slice 2 Product.** Slice 2 requires moving from eligible Public Campaign Detail into Donation; the current Campaign spec/API and live UI say the action is unavailable for all public detail responses. | **Blocking** for the Campaign Detail → Donation entry path. Product meaning is clear; using the old contract or inferring eligibility client-side would contradict authority. Other Donation form/status analysis is not blocked. | **Yes.** Stage 3 can safely assess routing/reconciliation scope; this does not authorize changing Campaign contract/spec in the frontend Build. | Orchestrator to coordinate Anhar Solehudin (current Slice-2 Campaign delivery/domain and API/contract owner per `.harscode-spaces/authority-map.md`) with WU-S2-004 scope; reconcile the public action/eligibility contract against Slice 2 before Build implements the entry path. Backend remains authoritative for actual eligibility. |
| **F-2 — Optional email verification and terminal-notice lifecycle are not an end-to-end frontend/API behavior yet.** Product/spec require verification before notices and a 24-hour verification disclosure; O3 still leaves controls, mechanism, delivery retries, retention/deletion races, and evidence open. The Donation submission contract carries email + opt-in but no explicit verification-completion surface/operation. | **Needs further evidence** for the optional verified-notification subflow; the no-email guest donation path and non-notification form state can progress. | **Yes.** Stage 3 can identify the frontend/API/backend dependency without inventing verification UI or asserting a notice was delivered. | Anhar Solehudin in current Slice-2 Donation/API and Security/PII owner roles to establish the supported verification interaction/contract and controls for a complete opt-in-notification flow; retain the product-approved disclosure. |
| **F-3 — Guest credential safety and public failure parity are contract directions, not proven behavior.** Fragment handoff/cleanup and uniform `404` are settled; browser/infrastructure exposure controls, abuse behavior, and runtime parity remain O4/O5. | **Deferable** for the `FRONTEND_MOCK_VERIFIED` milestone only as to backend/runtime proof; frontend handoff and visible URL cleanup remain in scope. | **Yes.** Stage 3 can proceed; WU-S2-004 mock work can document its frontend evidence boundary while independent backend/security Testing remains downstream. | Frontend Build owns its browser-visible handoff/cleanup behavior; Anhar as current API/Security/PII owner and independent Testing own remaining runtime controls, empirical parity/abuse evidence, and any residual-risk decision. |
| **F-4 — Invocation prior-artifact path/name mismatch.** TPD-S2-002-002 is present, but its manifest and Task 02 snapshot use `tasks/manifest.md` and `02-donation-openapi-reconciliation.md`, not the paths/names listed in Invocation. | **Informational.** The active contract/specs, approved TP-015, and actual snapshot locations were discoverable; no analysis scope is blocked. | **Yes.** No extra confirmation is required to understand current authority. | Orchestrator may correct pointers in a future Invocation amendment if deterministic resolution requires it; no current source change is needed for this Run. |

## [SCOPED BLOCKER DETECTED]

F-1 blocks the required Campaign Detail → guest Donation entry path under the current Campaign contract. Stage 3 may proceed to assess and route a coordinated Campaign/API contract reconciliation; this does not authorize a frontend-only contract change or a client-inferred eligibility rule. Other Stage-2 analysis and the Donation form/status planning are not blocked.

### Boundary

- **Explored:** Product/domain boundary, authored Donation API and relevant Campaign action contract, Design/surface authority, frontend architecture and live implementation.
- **Stage-2 result:** Guest submit/status shapes are present in the accepted Donation contract but have no frontend consumer; Campaign detail currently cannot expose the Slice-2 entry action under its current contract. The frontend donation and tracking implementation remains absent. Design readiness is PARTIAL for ordinary form/status presentation; material design ambiguity is not currently evidenced.
- **Active blocker:** F-1 blocks the Campaign Detail entry path until Campaign/API contract reconciliation and cross-stack coordination. Stage 3 itself is not blocked; it can safely route the necessary reconciliation without altering authority.
- **Further evidence needed:** F-2 affects only the optional verified-email notification subflow. F-3 runtime security/parity proof is outside what frontend MSW can establish and remains assigned to its owners.
- **Owner routing source:** `.harscode-spaces/authority-map.md` maps current Slice-2 Campaign delivery/domain, API/contract, Donation delivery/domain, Security/PII, and Product Design authorities to Anhar Solehudin; the attribution is scoped to the current Slice 2 and does not itself decide the unresolved contract/control questions.
- **Not tested:** No commands, automated tests, browser rendering, or runtime/API calls were run in this Exploration Stage 2. Existing tests were read only. No implementation/spec/API changes were made; this artifact is the only Run evidence written so far.

**Decision ask:** Confirm that this Stage-2 analysis is sound and allow Stage 3 solutioning to examine the Campaign contract coordination path and the bounded frontend/API dependencies, or redirect the analysis before solutioning.
