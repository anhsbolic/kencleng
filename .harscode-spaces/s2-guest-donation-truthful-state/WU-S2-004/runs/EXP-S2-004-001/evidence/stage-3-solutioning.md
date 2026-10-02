# Stage 3 — Solutioning

> Phase/Stage: Exploration / Stage 3  
> Author: Explorer (`P-S2-004-EXP-001-1`)  
> Created: 2026-10-01  
> Model / Reasoning: Invocation-selected `gpt-6-luna` / `high` (runtime model metadata not independently exposed in this session)  
> Session: interactive dispatch session; session identifier not exposed  
> Target revision: `7fd8b473b239b20bda3990ab29c51440d321a796`; current working tree also contains Stage-2 evidence and Orchestration's Stage-3 continuation reconciliation; relevant sources re-read live  
> Workflow revision: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`  
> Work Unit / Run: `WU-S2-004` / `EXP-S2-004-001`

## Solutioning state

The Stage-2 evidence confirms one scoped blocker: Product/MVP requires a guest donation entry from an eligible Public Campaign Detail, while the current Campaign detail contract and frontend state that donation is unavailable. The Product requirement is clear; the lower-level Campaign contract/spec is the item that needs reconciliation. This does not reopen or amend Product/MVP authority. Anhar authorized Stage 3 for this same Run/Participant on 2026-10-01; the current Invocation and orchestration state record this continuation, and no solution/API/Product decision is implied by that authorization.

### Decision framing — Campaign action contract coordination

**Problem**  
WU-S2-004 cannot implement its required Campaign Detail → Donation entry against the current accepted `PublicCampaignDetail` shape: `donation_action` is required and only permits `unavailable` / `donation_flow_not_available`.

**Current context**  
Slice 1's `GET /campaigns/{campaignId}` returns a closed, public-safe projection only after backend public-eligibility evaluation, and its spec/API deliberately make donation unavailable. Slice 2 now requires a usable entry for an eligible public Campaign. Donation submission still checks eligibility in backend state; a detail response can become stale and must not authorize the POST. Root/scoped guidance keeps frontend and backend production writes separate by default and requires explicit coordination for cross-stack contract changes. The current Slice-2 Campaign delivery/domain and API/contract owner is Anhar Solehudin (`.harscode-spaces/authority-map.md`).

**Options**

1. **Separate bounded Campaign/API contract reconciliation before WU-S2-004 Build.** Keep WU-S2-004 frontend-only. The Campaign/API owner reconciles the public `donation_action` projection and corresponding Campaign acceptance/spec against Slice 2, retaining backend-authoritative POST eligibility and documenting stale-detail rejection. Exact response shape remains for that owner to define from current authority.
2. **Explicitly re-scope WU-S2-004 to include the Campaign contract/spec reconciliation alongside frontend delivery.** This keeps the dependency in one Work Unit, but expands the current frontend-only production boundary, requires explicit cross-stack coordination, and increases the review/verification surface.

**Recommendation**  
Choose Option 1. Route the contract work as a separate, narrowly scoped owner-led reconciliation before WU-S2-004 Build. Keep the Donation frontend Work Unit and its FE mock completion boundary intact.

**Why and consequence**  
Option 1 follows the existing ownership boundary and lets the Campaign/API authority define the exact public-safe response semantics. It avoids silently ignoring the required field, hard-coding a CTA, or treating frontend visibility as eligibility authority. Its cost is a dependency and possible schedule delay before the full Campaign Detail entry path can be built. Option 2 remains viable only with explicit Orchestrator/Human scope coordination; otherwise it conflates a contract reconciliation with frontend delivery.

**Rejected direction**  
Rendering a CTA by ignoring `donation_action` or inferring eligibility from the current `public_state` was considered and rejected. It contradicts the current authored Campaign contract and puts eligibility meaning in the client. A later `409` from POST is necessary for stale/racing state, but does not justify bypassing the read contract.

**Decision ask**  
Confirm Option 1 (separate Campaign/API contract reconciliation before frontend Build), or explicitly authorize Option 2 (re-scope WU-S2-004 to include the coordinated contract/spec work). The exact `PublicCampaignDonationAction` shape is not being selected here; it remains with the current Campaign/API owner during reconciliation.

## Bounded directions for the remaining frontend dependencies

### Optional email

- Preserve the approved optional-email field, explicit status-only opt-in, and exact label/helper text in the guest form; submit the accepted `guest_email` and `guest_email_status_opt_in` request fields.
- Do not add a frontend state that claims ownership verification or terminal email delivery. The current Donation contract does not define a verification-completion operation/surface. Treat the verification interaction, out-of-band email link behavior, and O3 controls as a Donation/API/Backend + Security/PII dependency. An opt-in notification subflow needs owner evidence before its end-to-end behavior can be represented; the baseline no-email guest donation path does not depend on it.
- This direction preserves approved Product behavior and current contract shape without inventing verification semantics. No Product/MVP wording change or residual-risk acceptance is proposed.

### Guest status credential and proof boundary

- Carry the settled contract direction into frontend planning: submission response token → fragment handoff → URL cleanup → `X-Donation-Status-Credential` lookup; display only status; map public lookup failure to the approved generic user-facing behavior. Implement and verify the frontend behavior in its own browser-visible evidence.
- Keep O4/O5 backend/infrastructure exposure controls, abuse controls, actual uniform-`404` body/header/cache/timing parity, and residual-risk decision with API/Security/PII and independent Testing. Contract-faithful MSW is not evidence for those properties.

### Design readiness and verification direction

- Readiness is **PARTIAL**: required page purposes, state meanings/copy, form/status/error patterns, and Sunlit Editorial guidance are established; route-local hierarchy/details can follow current authorities. No material design ambiguity or requirement for a high-fidelity Figma artifact was found.
- Use route-local composition unless a real repeated semantic contract emerges; the `ui/` and `shared/` registries are intentionally empty. Preserve responsive/accessibility review and Human rendered acceptance for the material guest flow.
- Donation/frontend responsibility includes unauthenticated status credentials and optional PII, so plan the FE work as Tier 1 at minimum, subject to Techplan's actual risk assessment. No residual risk is accepted.

## Phase handoff

- **Completed:** Stage-2 analysis across Product/domain, Donation/Campaign API contract, Design, frontend architecture and live implementation; Stage-3 solutioning and recommendation recorded. No code, spec, Product, Design, API, test, or orchestration state was changed.
- **Artifacts:** `evidence/stage-2-gap-analysis.md`; `evidence/stage-3-solutioning.md`.
- **Human decision:** Select Option 1 or explicitly authorize Option 2 for the Campaign/API contract dependency.
- **Open / deferred:** Exact Campaign action response shape; optional email verification interaction and O3 controls; O4/O5 runtime credential/parity/abuse evidence and residual-risk decision; rendered Human acceptance. These remain with their stated owners.
- **Recommended next step:** Orchestrator records the selected contract-coordination route. After the public Campaign action contract is reconciled (or the scope is explicitly reauthorized), start WU-S2-004 Techplan synthesis from a new Techplan Run/Participant with a fresh Session. Do not hand the current Explorer Session across the Run boundary.
- **Session transition:** Fresh Techplan Participant Session is required because this is an orchestrated phase transition with a new Run/Participant, not continuation of this Exploration execution occurrence.
- **Context pointers:** Product/MVP `mvp-delivery-slices.md` §5; accepted Donation sources `docs/spec/5-donation/{invariants.md,threat-model.md,tasks.md,features/01-submit-donation-settlement.md,features/02-donation-status-check.md}`; Campaign sources `docs/spec/4-campaign/invariants.md` INV-campaign-14 and `features/02-campaign-detail-listing.md`; authored API `api/openapi/donation.yaml`, `api/openapi/campaign.yaml`, and referenced `common.yaml`; design `docs/ui-ux/README.md`, `page-map.md`, `product-design-principles.md`, `patterns.md`, `design-guidelines.md`; frontend anchors in Stage 2, especially `CampaignDetailView` `CampaignSuccess`, `public-campaign.ts`, `mocks/handlers/public-campaign.ts`, and generated OpenAPI `PublicCampaignDonationAction` / Donation operations.

## Verification and risk note

- **Verified:** Source/spec/code anchors read during Exploration; current branch `HEAD` is `7fd8b473b239b20bda3990ab29c51440d321a796`.
- **Assumed:** None material to the recommendation. The analysis treats Product Slice 2 as current authority and Slice-1 Campaign API/spec as requiring reconciliation where they conflict.
- **Deferred:** Exact Campaign API projection; email verification mechanics; live backend/security proof; Human rendered acceptance.
- **Not tested:** No command, automated test, browser render, or runtime API call was run. Existing frontend tests were read only.
- **Risk note:** The Campaign contract blocker prevents implementing the required entry path without coordination. Mock-only success cannot establish backend eligibility, status credential security, email delivery, settlement/funding correctness, or runtime anti-enumeration behavior.
