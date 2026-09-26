# Exploration Evidence — Stage 3 Solutioning

## Provenance

- Phase/Stage: Exploration / Stage 3 — Solutioning
- Work Unit: `WU-S2-001`
- Run: `EXP-S2-001-001`
- Role: Explorer
- Participant / Author: Codex Explorer
- Created: 2026-09-25
- Model / reasoning: `gpt-6-luna` / `medium` (invocation dispatch metadata)
- Session: `01a0d8ea-1404-7521-99b0-5623057b0519`
- Target revision: `ee0d4b072d9f5cf279952fe309049f687c95e30e` (invocation)
- Observed checkout: `7ee281c4acf6c6860ba52830fe3980b8a88e1940`; the target revision is its ancestor, with the invocation commit as the only later commit.
- Workflow revision: `b2d7ca4918b520d960139bc392f87619410b27ed`
- Human gate: Stage 3 authorized by the user's `lanjut ke stage 3` message after Stage 2 completion.

## Direction considered

This artifact records Exploration's recommended delivery direction. It does not alter Product Authority, approve an API shape, assign downstream Work Units, accept security residual risk, or authorize implementation.

| Decision area | Viable directions considered | Selected direction and rationale | Consequence / remaining boundary |
|---|---|---|---|
| Slice scope | (A) Reuse the historical full Donation domain scope. (B) Deliver only the public guest-donation and truthful-status capability plus what makes it financially/security correct. (C) Wait for Account and operational self-service first. | **B.** It matches the approved MVP trust loop and Slice 2 user outcome. The Product/MVP authority explicitly defers Account prerequisite, guest claim/history, public donor/social list, real rails, and payment-method breadth that is not needed. | Keep historical detail as evidence to classify `KEEP` / `ADAPT` / `REPLACE` / `DEFER`; do not import its complete task order or endpoint set. No FE/BE Work Unit topology is selected here. |
| Donation result truth | (A) Present frontend-only simulated states. (B) Persist real application donation state and let a backend-owned sandbox process produce pending/success/failure. (C) Integrate real external payment settlement. | **B.** The product requires state real within the sandbox and explicitly says sandbox behavior is not external settlement; client-only success would be misleading, while external rails are outside current scope. An internal-only state transition is consistent with historical threat evidence, which specifically prohibits an exposed settlement HTTP transition. | Keep transition authority server-side and make outcomes observable from persisted state. Exact timing, failure rate, worker/job mechanism, and payment-method label remain unchosen delivery/contract detail. Do not expose a client-callable success/failure transition. |
| Guest revisit | (A) Require login/history or claim. (B) Support status access for a guest independently of Account. (C) Depend on guest email delivery as the only way back. | **B.** Guest status revisiting is an approved product need; Account is not a baseline prerequisite. Historical optional-email behavior also means email-only cannot cover every possible guest. | The credential mechanism and policy remain unresolved. Historical `status_token` is a candidate only; its non-expiring/query-parameter behavior must not be treated as approved until credential exposure, storage/comparison, response parity, caching/logging/referrer, and risk acceptance are reconciled. |
| Donation API baseline | (A) Adopt all historical OpenAPI operations and fields. (B) Reconcile a narrow contract against Slice 2 and active public Campaign contract. | **B.** `api/openapi/donation.yaml` contains public donor listing and Account operations outside Slice 2, six payment methods beyond the stated minimum need, and at least one mismatch with its feature spec. Current public Campaign contract has an unavailable action for Slice 1; Slice 2 must reconcile the hand-off without mistaking that prior-slice state for active functionality. | A shared contract is a prerequisite for durable FE/BE implementation coordination. This Exploration does not choose endpoint names, request/response fields, path topology, or authentication semantics. |
| FE/BE sequencing | (A) Start independent FE/BE work against the old contract. (B) First reconcile the shared Slice-2 contract, then select backend-first or contract-parallel execution from the settled dependency/risk evidence. | **B.** Kencleng orchestration says backend-first is preferred when semantics are still evolving and contract-parallel is appropriate only once contract stability is sufficient. No current Donation contract or runtime exists. | Orchestrator decides downstream Work Units and dependencies. After contract readiness, FE may use contract-faithful mocks if parallel work is justified; integrated evidence must use the real backend. |
| Maximum amount / closure boundary | (A) Import the old full-donation-then-close behavior. (B) Treat all closure behavior as Slice 3 and ignore it for Slice 2. (C) Reconcile only the closure/eligibility behavior needed to prevent accepting an ineligible Campaign. | **C, as a reconciliation boundary, not as a detailed rule.** Slice 2 excludes closure/result behavior except what is needed to keep donation eligibility correct; old Campaign/Donation artifacts connect maximum amount to settlement. Neither importing the old policy wholesale nor ignoring its eligibility consequence is justified by the current evidence. | Product/delivery owners must settle what happens when a successful donation crosses `max_amount` and how new submissions are rejected thereafter. Historical full-amount/over-target behavior remains unapproved for Slice 2. |

## Recommended material direction

Treat Slice 2 as a single narrow, end-to-end **Guest Donation + Truthful Donation State** capability. Reconcile one shared Donation contract that starts from an eligible public Campaign and covers guest submission, a persisted pending/result state, a safe guest revisit path, and exactly-once successful funding reflection. Keep sandbox semantics explicit and backend-owned. Establish the smallest correct guest-data and payment surface that meets the approved outcome; leave account-dependent history/claim, donor listing, payment-method breadth, and real payment rails deferred unless new enabling-critical evidence changes scope through its owning authority.

Use decimal money representation throughout, server-authoritative eligibility at submission, idempotent handling of repeated submissions where required, a non-HTTP settlement transition, and transactional/concurrency-safe coupling between success state and funding update as non-negotiable reconciliation constraints. These constraints are supported by Product/MVP and root/backend rules; exact mechanism and API contract remain for Techplan/contract reconciliation. The maximum-amount/closure interaction must be resolved at the owning Campaign/Donation boundary before implementation.

For frontend readiness, classify the Donation interaction as **PARTIAL**: canonical Form and Status/Tracking patterns, truth-state principles, public product direction, and visual foundations exist; the active Donation contract, exact status vocabulary/source labels, safe revisit mechanism, and concrete form/status requirements are not yet reconciled. No brand-defining asset decision is needed from this evidence.

For backend risk routing, settlement/balance correctness includes Tier-0 protected ledger/transaction-locking areas in root policy; agent exploration/review/test ideas are permitted, but a later agent Build cannot write protected code without the explicit human-paired authorization required by root authority. Other exact tier assignments remain a before-Build decision, based on the reconciled threat surface.

## Material rejected alternatives

- **Historical Donation domain as the scope authority:** rejected because approved Product/MVP scope owns inclusion and explicitly defers several historical operations. Detail and pre-existing generated types do not promote those operations into Slice 2.
- **Account-first flow:** rejected because Guest Donation is the primary approved actor path and Account is not a baseline prerequisite.
- **Client-only pending/success/failure presentation:** rejected because it would not be backed by real persisted application state and could imply settlement that the sandbox has not established.
- **Real payment gateway/rail:** rejected for the baseline because the approved product is a sandbox and expressly excludes real payment rails absent slice requirement.
- **Unreconciled legacy status token semantics:** rejected as an automatic carry-forward. A guest credential is needed in concept, but historical permanence and query-parameter placement introduce material exposure/retention questions and status API error semantics conflict across sources.
- **Parallel frontend/backend work against the current historical contract:** rejected because there is no stable reconciled Slice-2 contract or live Donation implementation. Contract-parallel becomes viable only after the contract is reconciled.

## Open / deferred questions

These are not answered by choosing the delivery direction above:

1. What amount floor/precision and one sandbox payment/processing representation are appropriate for the approved outcome? Historical Rp 5,000 and six-method choices require revalidation.
2. What truthful timing/outcome behavior should produce `pending`, `success`, and `failed`; what does the product promise to users while each state holds?
3. What guest fields are essential, which are optional, and what retention/notification behavior follows? Any guest email must follow the established encryption/HMAC and sensitive-logging conventions.
4. What safe guest revisit credential and handling policy will be used? Resolve delivery, storage/comparison, lifetime/revocation, URL/history/log/referrer exposure, cache behavior, anti-enumeration, and rate-limit/risk acceptance. The old token draft is not approval.
5. What uniform public response applies when a donation is absent versus a status credential is missing or invalid? Resolve the feature-spec `401` behavior versus the OpenAPI `404` response.
6. What is the exact `max_amount`/successful-settlement interaction, including crossing the cap and subsequent eligibility checks? Coordinate Campaign and Donation owners; do not silently import the draft rule.
7. Confirm whether remaining UI wording/source-label decisions can use existing Product Design principles or need an explicit Design Authority decision once the contract's state meanings are known.

Product meaning for the baseline journey is sufficiently clear to continue into Techplan synthesis. If those delivery questions require authority or risk-acceptance decisions, route each through the owning authority/gate rather than filling the gap in code.

## Phase handoff

- Completed: Stage 2 gap analysis across Product/MVP, Product Design, Donation specs/API, backend, frontend, and cross-stack boundaries; Stage 3 compared directions and recorded the recommended Slice-2 direction. No authority was changed and no implementation decision was written into code/spec/API.
- Artifacts: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-2-gap-analysis.md`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-3-solutioning.md`.
- Human decision: None required before Techplan synthesis; the open delivery/security/contract questions above must be routed to their owners if the next phase reaches a gated authority or risk-acceptance decision.
- Open / deferred: Exact amount/payment/timing rules, guest-data policy, status credential and anti-enumeration semantics, `401`/`404` contract mismatch, and `max_amount`/closure behavior.
- Recommended next step: Techplan synthesis for `WU-S2-001` using both Run evidence artifacts; surface unresolved owner decisions explicitly before locking implementation contract or verification obligations.
- Session transition: Continue Techplan synthesis in this Session because the context remains focused on this slice, no major redirection/dead ends accumulated, and current durable authorities/evidence are named. Techplan is a new Run with its own identity/path.
- Context pointers: `docs/product/mvp-scope.md` §§4–7; `docs/product/mvp-delivery-slices.md` §5; `docs/ui-ux/product-design-principles.md` §§1, 3, 5–6; `docs/spec/5-donation/invariants.md` INV-donation-02/04/05/08/09/11; `docs/spec/5-donation/threat-model.md`; `api/openapi/donation.yaml`; `api/openapi/campaign.yaml` public projection; `backend/cmd/server/main.go` router and `backend/internal/domain/campaign/service.go` `toPublicDetail`; `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` `CampaignSuccess`.
