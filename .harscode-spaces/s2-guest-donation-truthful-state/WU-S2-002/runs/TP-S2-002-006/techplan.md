# Tech Plan: Slice 2 Donation Domain & Contract Reconciliation

> Phase             : Techplan amendment
> Work Unit         : WU-S2-002
> Run               : TP-S2-002-006
> Role              : Planner
> Specialization    : Material Techplan amendment after Human-approved Product/MVP scope update
> Participant       : Codex Planner
> Author            : Codex Planner
> Model             : gpt-6-luna
> Reasoning         : high
> Session           : Fresh Planner Session; Session ID not exposed
> Session transition: FRESH — MATERIAL_PLAN_REVISION; independent of prior synthesis, resolution, and review sessions.
> Created           : 2026-09-26
> Updated           : 2026-09-26
> Target revision   : 10457b17059e2da3d97a4f76e4c3fae127227d73
> Workflow revision : b122a75d494250d04eb93e71f4c391e82c847842
> Status            : Draft / In Review
> Approach          : Reconcile Approved `TP-S2-002-003` to the Human-approved 2026-09-26 Slice 2 amendment and OIR-S2-002-001; retain unresolved owner decisions.
> Refs              : `TP-S2-002-003/techplan.md` (revised; historical copy preserved); `OIR-S2-002-001/resolution-brief.md`; `docs/product/mvp-scope.md`; `docs/product/mvp-delivery-slices.md`; current Product, Design, Donation/Campaign, API, Security/PII authorities.

---

## 1. Background

This is a material revision of the previously Approved Techplan `TP-S2-002-003`, not a replacement of its historical artifact. On 2026-09-26, Human approved additional Slice 2 product requirements and explicit OIR decisions O1–O6/O9. The prior plan predates those requirements. Its approval does not approve this revision.

The amendment changes product/domain behavior, scope, interface implications, risks, and verification obligations. The Work Unit remains contract reconciliation for the guest donation and truthful sandbox status slice. Donation runtime routes and implementation were absent at the inspected target revision; old Donation specs/OpenAPI remain historical evidence pending reconciliation.

## 2. Scope

**In scope:**

- Reconcile Donation invariants, threat model, feature acceptance, and the authored split OpenAPI contract for the approved Slice 2 guest flow.
- Carry forward the Human-approved amount, display-method, sandbox state/recovery, guest name/email, status-link, generic-failure, threshold, and idempotent-retry directions below.
- Preserve unresolved schema, simulator, Security/PII, anti-enumeration, Design, and conditional API-consumer questions as Active Open Items with owner/evidence needs.
- Coordinate Campaign only for eligibility, threshold crossing, already-accepted pending donations, and rejection after closure. Do not pull Slice 3 closure/result behavior into this Work Unit.
- Establish contract/spec verification and the later runtime evidence obligations needed to preserve atomic success/funding and security boundaries.

**Out of scope (explicit):**

- Changing Product/MVP or Design authority, accepting Security/PII residual risk, or selecting controls for their owners.
- Donation backend/frontend runtime, migrations, manual DB/index application, or runtime test execution.
- Account prerequisite, guest claim/history, donor/social-proof list, active GoPay/ShopeePay/bank-transfer processing, real payment rails, or campaign-wide email.
- Tax policy, tax calculations, and derived-value precision/rounding decisions not set by Product; storage must support exact decimal values without implying those features are in Slice 2.
- Full Campaign closure/result lifecycle, scheduler, force-close, public result behavior, or other Slice 3 work beyond the approved threshold rule and eligibility boundary.
- Changes to protected Tier-0 ledger/transaction-locking, encryption/HMAC/key-handling, auth core, or state-machine implementation.
- Orchestration projections, tracker, source authorities, prior Runs/reports, or generated API artifacts in this Planner Run.
- Claiming `CONTRACT_READY`, starting Build, or treating this Draft as Human-approved.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | Guest Donation without Account prerequisite must produce persisted, truthful sandbox processing/result state. | `docs/product/mvp-scope.md` §§4–5, 9; `docs/product/mvp-delivery-slices.md` §5, amended and Human-approved 2026-09-26. |
| Q2 | Donation input is IDR whole Rupiah, minimum Rp5.000, increments Rp1 (Rp5.001 valid); stored/calculated money uses exact decimal representation. No `float64`. Derived-value scale/rounding remains undecided; tax rules are out of this slice. | `docs/product/mvp-delivery-slices.md` §5 (2026-09-26 amendment); `AGENTS.md` §2; `../harscode-workspace/best-practices/go/decimal-and-money.md`. |
| Q3 | Display QRIS, GoPay, ShopeePay, and bank transfer; QRIS alone is an active, clearly labeled sandbox simulation. Other options are unavailable and non-interactive. No real rail, usable real-payment instruction, or transfer of money. | `docs/product/mvp-delivery-slices.md` §5 (2026-09-26 Human-approved display exception); `docs/product/mvp-scope.md` §§4–5, 7. |
| Q4 | Backend simulator owns persisted `pending`/`success`/`failed`; failure comes only from a clearly labeled demo scenario controlled by backend configuration/fixture, not donor choice/request. Pending copy is “Menunggu hasil simulasi,” with no estimate or real-payment instruction. Failed may lead to an explicitly initiated new donation; pending never triggers auto-resubmission. Exact simulator timing is not approved. | `docs/product/mvp-delivery-slices.md` §5; O2 in `OIR-S2-002-001/resolution-brief.md`; historical timing/probability is not authority. |
| Q5 | Same-key ambiguous retry returns the original donation; same key with a different payload is rejected. Client prevents double-click and does not rotate key during ambiguity. A new key is used only after intentional action to start a new donation. This request policy is separate from settlement replay/funding correctness. | `docs/product/mvp-delivery-slices.md` §5 (2026-09-26 amendment); O9 resolution; `../harscode-workspace/best-practices/restapi/idempotency-and-versioning.md` §1. |
| Q6 | Guest name is optional and not public by default. Email is optional, opt-in, and limited to donation-status notices, separate from Account email verification and not for campaign-wide updates. Verify ownership before sending status/access; send at most one status-only notice at terminal `success`/`failed`, not initial `pending`, identify it as simulation, and hold/delete unverified email per Security/PII windows. | `docs/product/mvp-delivery-slices.md` §5 (2026-09-26 amendment); O3 resolution brief; `AGENTS.md` §2. |
| Q7 | Guest status URL lasts 24 hours, uses a difficult-to-guess token, and reveals donation status only. No need to access via that link after expiry. Invalid/missing/expired use one generic public behavior/copy. Any email/Account benefit messaging after access is informational and non-coercive, does not gate the guest flow, and cannot promise unavailable functionality. Exact transport/control and residual-risk acceptance remain Security/API decisions. | `docs/product/mvp-delivery-slices.md` §5; O4/O5 resolution brief; `docs/ui-ux/patterns.md` §7. |
| Q8 | `max_amount` is a closure threshold, not a hard cap. A successful donation crossing it triggers close and is accepted in full; donations accepted while eligible settle in full even if still pending when closure occurs. New submissions after closed are rejected. This is only the approved Slice 2 threshold/eligibility rule. | `docs/product/mvp-delivery-slices.md` §§5–6 (2026-09-26 amendment); O6 resolution; Campaign/Donation owner coordination remains required for contract detail. |
| Q9 | Preserve atomic success/funding coupling, exact-once funding updates, concurrency safety, non-forgeable settlement, backend-authoritative eligibility, private guest status, and safe sensitive-data handling. | `docs/product/mvp-scope.md` §§2, 7; `docs/product/mvp-delivery-slices.md` §5; `AGENTS.md` §§2–3; Exploration Stage 2 Areas 3–5. |
| Q10 | Reconcile only current Slice 2 contract sources. `api/openapi/<domain>.yaml` and referenced shared components are authored; index/bundle/generated clients follow `api/README.md` and must not be edited by hand. | `api/README.md` §§Structure, Editing workflow; `AGENTS.md` §1. |

## 4. Rules & Validation

- **R1 — Slice boundary:** Reconciled Slice 2 spec/API covers guest submit, truthful persisted sandbox state, safe status revisit, and successful funding truth. Account claim/history, public donor list, real rails, active non-QRIS methods, and broader Campaign closure/result remain deferred absent enabling-critical evidence and authority approval.
- **R2 — Amount and displayed methods:** Accept only whole-IDR input at or above Rp5.000 in Rp1 increments; Rp5.001 is valid. Preserve exact-decimal money through storage/calculation and never use `float64`. Show the four Human-approved method labels; only QRIS is interactive as sandbox, and other methods are visibly unavailable/non-interactive. No payment instruction or screen may imply real settlement. Do not invent derived-value precision/rounding or tax behavior.
- **R3 — Atomic state/funding correctness:** Settlement cannot be forged via exposed HTTP. For every successful settlement, Donation `success` and its funding reflection are one atomic business outcome: every committed/observable state includes both, or neither; each contribution affects funding once, including replay and concurrent settlement without a lost increment. Pending/failed do not count as collected funding. Do not prescribe or edit protected ledger/locking implementation in this plan.
- **R4 — Simulator-owned state and recovery:** Initial donation state is persisted `pending`; only backend simulator behavior determines terminal `success`/`failed`. Failure is a clearly labeled demo scenario controlled by backend configuration/fixture, unavailable as a donor-selected or request-controlled outcome. Pending wording is “Menunggu hasil simulasi,” without time estimate/payment instruction or auto-resubmit. After `failed`, a donor may explicitly start a new donation with a new key. No historical exact timing/probability becomes a requirement.
- **R5 — Guest fields and notification:** Name is optional and not public by default. Email is optional and opt-in for donation status only, separate from Account verification and not for campaign-wide updates. Verify ownership before any status/details/access-link email; do not put Donation detail/access in the verification message. At most one terminal status-only email on `success`/`failed`, never initial `pending`, clearly labeled as simulation. If terminal state precedes verification, hold only during the Security/PII-approved window, then delete unverified email without notice. Apply Security/PII-approved delivery retry/retention; use established encryption/HMAC and safe-logging pattern. Exact windows and controls remain open.
- **R6 — Status credential and anti-enumeration:** Status access is via Human-directed 24-hour guest URL with hard-to-guess token and status-only data. Expired/missing/invalid access has one generic public behavior and copy. Any post-access email/Account messaging is informational, non-coercive, and cannot promise unavailable functionality or gate guest access. API/Security must decide and verify response code/body/headers/timing/cache parity, secret comparison, abuse controls, token/URL exposure mitigations, and residual-risk acceptance before final contract readiness. Do not infer that a hard-to-guess token alone resolves leakage risk.
- **R7 — Eligibility and threshold:** Backend enforces Campaign eligibility at submission time. A successful settlement crossing `max_amount` triggers Campaign close but is accepted at full amount; all donations already accepted while eligible settle at full amount even if still pending after close; later submissions after closed are rejected. Coordinate ordering/state consistency with Campaign owner while limiting this slice to the stated threshold rule; do not import Slice 3 closure/result lifecycle.
- **R8 — Submission idempotency:** For one logical donation, ambiguous retry reuses the same key and returns the same donation; same key/different payload is rejected. Client suppresses double-click and retains key through ambiguity. New key means new donation only after deliberate donor action. Translate persistence/response/key-retention details into the contract without reopening the resolved policy. Keep separate from R3 settlement replay.
- **R9 — Spec/API alignment and compatibility:** Reconciled invariants, threat model, feature acceptance, split OpenAPI, shared components, and Campaign boundary must agree with current Product/MVP and unresolved owner boundaries. Before removing/replacing a historical operation, establish API consumer/distribution status; preserve/stage compatibility if a consumer is found. Update only authored split source/shared components; update index and generated bundle/types by canonical workflow when relevant.
- **R10 — Truthful design expression:** Status meaning follows domain/product semantics; label, source disclosure, unavailable methods, and recovery must not imply provider settlement, verification, or certainty not established by simulator state. Design Authority reviews the proposed flow/labels against current patterns; this plan does not select final visual treatment.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Guest donation and truthful backend-simulator state are the Slice 2 baseline; Account is not required. | Chosen — current Product/MVP | Product scope and Slice 2. |
| D2 | Settlement/result is backend-owned and persisted; exposed client/HTTP cannot forge terminal state, and successful Donation state/funding is atomically coupled and reflected exactly once. | Chosen — correctness floor | Product/MVP §§2, 7 and Slice 2 security floor; implement/runtime evidence remains future work. Preserve R3 from the Approved plan. |
| D3 | Input/display: whole IDR Rupiah, min Rp5.000, Rp1 steps; exact decimal for stored/calculated money; QRIS alone active sandbox among four displayed Indonesian method labels. | Chosen — Human, 2026-09-26 | Explicit Product/MVP amendment resolves O1 product direction. Amount encoding and derived precision are not selected here. |
| D4 | Defer Account claim/history, donor list, real rails, and active GoPay/ShopeePay/bank transfer; display-only inactive methods are the Human-approved exception to the earlier no-method-breadth wording. | Chosen — Product/MVP, 2026-09-26 | Display does not create functional breadth or imply integration; options must be unavailable/non-interactive. |
| D5 | State/recovery: `pending` then simulator-owned `success`/`failed`; failure only via clearly labeled demo scenario controlled by backend config/fixture; exact timing/probability not inherited. Pending copy and explicit failed recovery follow Q4/R4. | Chosen — Human, 2026-09-26 | Resolves product semantics in O2; exact timing and demo scenario mechanism remain delivery/owner questions. |
| D6 | Guest name optional/non-public by default; email optional, opt-in status-only with verified ownership and terminal-only notice; unverified hold/delete and retention windows are Security/PII-owned. | Chosen — Human, 2026-09-26 | Resolves product direction in O3 without accepting privacy risk or choosing retention/control details. |
| D7 | 24-hour temporary guest status URL, difficult-to-guess token, status-only access; one generic behavior/copy for absent, invalid, or expired links. Email/Account messaging after access is informational and does not coerce or gate guest access. | Chosen — Human, 2026-09-26 | Resolves product-facing portions of O4/O5. URL/token exposure, response parity implementation, abuse controls, and residual-risk acceptance remain open to Security/API. |
| D8 | `max_amount` is a threshold: a successful crossing triggers close and is accepted in full; accepted-while-eligible pending Donations settle in full; future post-close submissions are rejected. | Chosen — Human, 2026-09-26 | Resolves O6 policy; coordinate only necessary Campaign/Donation boundary, not Slice 3 lifecycle. |
| D9 | Same-key ambiguous retry returns original Donation; same key/different payload rejected; intentional new donation uses a new key; UI prevents double-click and key rotation during ambiguity. | Chosen — Human, 2026-09-26 | Resolves O9 policy; response/serialization and retry-record lifetime are contract detail. R3 settlement idempotency remains separate. |
| D10 | Use exact-decimal backend money convention; no float. | Chosen — project rule | `AGENTS.md` §2 and Harscode `go/decimal-and-money.md`; these sources do not specify derived-money scale or rounding for this feature. |
| D11 | Historical Donation spec/OpenAPI values are evidence only; source authority is Product/MVP first, then reconciled domain/API detail. | Chosen — repo authority | Root `AGENTS.md` §1; `docs/product/README.md`; `docs/spec/README.md`; `api/README.md`. |
| D12 | Preserve current authored split OpenAPI discipline and generated-artifact workflow. | Chosen — API authority | No hand edits to aggregate bundle/generated client; consumer/distribution audit remains conditional O8. |
| D13 | No residual Security/PII risk is accepted by Product decisions or this Planner Run. | Boundary retained | O3–O5 require Security/PII and API owner controls/evidence and the applicable human risk gate. |
| D14 | Prior Techplan approval is not approval of this material amendment. Route to an independent fresh Reviewer Session; regenerate report only after review/resolution convergence and before the next Human gate. | Required phase route | Material product/domain/interface/verification semantics changed. Harscode Techplan rules/guardrails and invocation phase route. |

## 6. Backward Compatibility

- No Donation runtime route or reconciled Donation API is established at the target revision. Historical split OpenAPI operations and generated declarations are not proof of runtime support or approval.
- Do not assume historical Donation operations were never distributed externally. O8 stays conditional: API/Orchestrator owner must establish distribution/consumer status before a proposed removal/replacement. If a consumer is found, route compatibility/staging to its owner; do not hand-edit generated output.
- The revised contract will be a fresh reconciliation of the active Slice 2 source, not a silent commitment to preserve historical Account/list/settlement shapes. Any breaking contract consequence must be explicit after the O8 evidence is obtained.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| RISK-1 | Decimal storage/calculation or derived values round inconsistently; accidental float use. | Medium | High | Exact decimal end-to-end; no float. Derived precision/rounding remains open and blocks affected contract detail; no tax behavior assumed. |
| RISK-2 | Inactive method presentation or pending/failure copy implies real payment/provider outcome. | Medium | High | Only QRIS is active sandbox; clear labels; Design review; no live instructions or timing estimate. No residual trust risk accepted here. |
| RISK-3 | Simulator timing/scenario behavior diverges from status UX or implies an SLA. | Medium | Medium | No historic 2–5 second/5% policy; no timing promise. Delivery must define simulator timing/scenario mechanism with owner and Design review where wording is affected. |
| RISK-4 | Guest email/status link leaks PII or access before ownership verification; retention outlives need. | Medium | High | Email optional/opt-in; verification before content/link; terminal-only status notice; Security/PII sets windows/control; established encryption/HMAC and safe logs. No accepted risk. |
| RISK-5 | Bearer URL token is exposed through history, referrer, logs, cache, forwarding, or abuse; status reveals private donation information. | Medium | High | 24-hour/status-only product direction is not sufficient mitigation by itself. Security/API must decide controls, parity, abuse handling, and residual-risk decision. |
| RISK-6 | Missing/invalid/expired access responses reveal Donation existence through response or timing differences. | Medium | High | Human-approved generic public behavior/copy; API/Security still define and evidence transport parity across code/body/headers/cache/timing. |
| RISK-7 | `max_amount` crossing, pending accepted donations, close, or new-submit race causes inconsistent eligibility/funding or leaks Slice 3 behavior into Slice 2. | Medium | High | Preserve accepted-full semantics and reject new submissions after close; Campaign/Donation owner coordinates necessary ordering while R4 holds. Do not add broader closure/result flows. |
| RISK-8 | Ambiguous retry or fresh-key double activation creates duplicate intended donation, or client rotates key prematurely. | Medium | High | Resolved same-key policy; explicitly verify request-level behavior separately from settlement replay; client suppresses double-click. |
| RISK-9 | Historical API operation removal breaks unknown consumer. | Unknown | Medium | Conditional O8 consumer/distribution audit before remove/replace; generated artifacts only via API workflow. |
| RISK-10 | Terminal status/funding divergence, forged success, duplicate increment, lost concurrent update, or partial commit. | Medium | Critical | R4 and Tier-0 fencing; specialized downstream Testing evidence for atomicity, failure, replay, and concurrency; no implementation mechanism selected here. |

## 8. Interface Contract

**Persistence/data shape:**

- Product contract: `amount` represents whole IDR Rupiah input, minimum 5000 and integer increments of one Rupiah. Storage and calculated monetary values require exact decimal representation, never `float64`. Do not freeze wire encoding, DB precision/scale, or derived-value precision/rounding until Donation/API owners reconcile them against current conventions and requirements.
- Guest name optional and non-public by default. Guest email optional and opt-in only for donation-status notifications. Ownership verification, encrypted/HMAC storage pattern, delivery retry/retention, unverified deletion window, and safe logging follow current Security/PII authority; specific windows/control remain unresolved.
- Persisted Donation status begins pending and reaches terminal state only by authorized backend simulator path. Successful status and funding increment are coupled atomically and exactly once (R4). Settlement mechanism is not prescribed.
- Guest status credential has 24-hour lifetime and status-only access by Product direction. Token representation, storage/carrier, invalidation, cache/referrer/log/history mitigations, comparison, and residual-risk decision remain unresolved.

**API/event/external interface:**

- Scope is guest submission and guest status lookup required for this slice. Proposed operations are not finalized by this Techplan amendment; reconcile exact paths, fields, encoding, errors, headers, response shape, and simulator control in split Donation OpenAPI after owner decisions.
- Submission contract must express the resolved same-key behavior (same payload returns original Donation; changed payload with same key rejected) and deliberate new-key behavior (R8). Specify adequate retry-record lifetime/serialization without reopening policy.
- Status contract must provide one generic public behavior/copy for missing, invalid, and expired access. API/Security still own response code/body/header/timing/cache parity and abuse behavior; no status contract readiness until resolved/evidenced.
- QRIS is a labeled sandbox simulation; no external settlement rail or client-callable settlement interface. Other listed methods are non-interactive/unavailable.
- Authored source is `api/openapi/donation.yaml` plus only needed `api/openapi/common.yaml` components and path registry updates. Bundle `api/openapi.yaml` and frontend generated types are derived using `api/README.md`; do not edit generated files manually.

**Cross-layer/business boundary:**

- Frontend displays authoritative server/simulator state; it cannot create success, decide Campaign eligibility, or establish settlement. It retains the same idempotency key through ambiguous outcomes and prevents double activation.
- Backend checks Campaign eligibility at submission. Accepted while eligible means full amount remains eligible to settle even if threshold closure occurs meanwhile. New submission after closed is rejected. Coordinate only this boundary with Campaign; broader close mechanism/public result is Slice 3.
- Design reviews the state/method/failure/recovery presentation; Product and domain own semantics. Security/PII alone owns residual security/privacy acceptance.

## 9. Architecture / Plan

1. Treat amended Product/MVP and OIR brief as current requirements; compare prior Approved plan and preserve its history.
2. Resolve delivery contract details only within owner authority: amount encoding/scale and derived-value precision; simulator timing/scenario; email verification/retention/retry controls; guest token exposures/abuse and response parity; Design label/flow review; Campaign/Donation threshold ordering; conditional API consumer audit.
3. Reconcile only active Slice 2 Donation invariants, threat model, feature acceptance and required Campaign cross-reference. Mark historical material KEEP/ADAPT/REPLACE/DEFER by authority and evidence. Preserve atomic success/funding coupling and protected implementation boundaries.
4. Author/reconcile split Donation OpenAPI and only required shared components. Keep API index aligned; regenerate aggregate bundle and generated frontend types using documented API workflow. Do not hand-edit derived artifacts.
5. Review contract/spec consistency and source traceability. `CONTRACT_READY` remains unavailable while material Active Open Items or required owner reviews remain unresolved.
6. This amended Techplan is Draft / In Review. Independent fresh Techplan review precedes report regeneration and the next Human approval gate. No Build is authorized by this handoff.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `docs/product/mvp-scope.md` §§4–7, 9 | Whole-product and MVP authority for guest journey, security/correctness floor, Account boundary. | Current 2026-09-26 requirements supersede conflicting pre-amendment wording. |
| `docs/product/mvp-delivery-slices.md` §5 and approval record §12 | Detailed Human-approved Slice 2 requirements, display exception, retry/email/link/threshold behaviors. | Primary product source for this amendment; §6 only establishes the boundary with Slice 3. |
| `OIR-S2-002-001/resolution-brief.md` O1–O9 | Records exact per-item Human directions, consequences, and unresolved owner evidence. | O1–O6/O9 policy trace; O7/O8 and technical aspects remain active/conditional as listed below. |
| `TP-S2-002-003/techplan.md` all semantic sections | Previously Approved plan and historical basis. | Materially revise into this Run; retain the source artifact unchanged. |
| `EXP-S2-001-001/evidence/stage-2-gap-analysis.md` Areas 1–5; `stage-3-solutioning.md` | Product/spec/API/live-code/security/design Exploration evidence. | Preserve exact sensitive-risk anchors in Test Focus Pointer; use as evidence, not current product authority. |
| `docs/product/README.md`; `AGENTS.md` §1 | Product-first source routing and precedence. | Product/MVP owns what should be true; detailed specs/OpenAPI are downstream. |
| `docs/ui-ux/README.md`; `docs/ui-ux/patterns.md` §§7, 14–15; `product-design-principles.md`; `design-guidelines.md` | Design authority and anti-enumeration/truth/status presentation constraints. | Design review of actual flow/labels remains open; do not select visuals here. |
| `docs/spec/README.md`; `docs/spec/5-donation/invariants.md` INV-donation-01…15 and state machine | Spec authority/template and historical Donation invariant candidates. | Reconcile Slice 2 only; retain/adapt/defer each relevant legacy rule after review. |
| `docs/spec/5-donation/features/01-submit-donation-settlement.md` §§Critical, Submission, Settlement, Concurrency; `02-donation-status-check.md` §§Behavior, Validation | Historical submission/atomicity and 401-vs-404 contract conflict. | Evidence to reconcile; neither is automatically current truth. |
| `docs/spec/5-donation/threat-model.md` submit/status/internal-settlement sections | Historical threat evidence, including URL token and residual risks. | Re-evaluate controls and risk acceptance; old accepted risks do not authorize current exposure. |
| `docs/spec/4-campaign/features/09-closure.md` max trigger; `docs/spec/4-campaign/invariants.md` INV-campaign-13 | Historical cross-domain closure hook. | Coordinate only eligibility/threshold behavior; do not import full closure behavior. |
| `api/README.md` §§Structure, Editing workflow; `api/openapi/donation.yaml`; referenced `common.yaml` | Authored/derived contract discipline and historical Donation surface. | Reconcile authored source only after unresolved API/Security decisions; audit O8 before destructive shape change. |
| `../harscode-workspace/best-practices/go/decimal-and-money.md` | Exact-money/no-float convention. | Use as technical guidance; does not decide derived precision/rounding. |
| `../harscode-workspace/best-practices/restapi/idempotency-and-versioning.md` §1 | Technical retry/idempotency context. | Supports implementation/verification detail; Human O9 is the project policy authority. |
| `../harscode-workspace/best-practices/restapi/anti-enumeration.md` | Generic failure, secret comparison, response timing security guidance. | Security/API owner selects and evidences controls; this plan does not accept residual risk. |
| `../harscode-workspace/best-practices/go/secrets-and-sensitive-logging.md`; `docs/project/kencleng-backend-tech-stack.md` §Encryption Key Management | PII/token logging and existing encryption/HMAC convention. | Retain established project pattern; no new credential or PII convention invented here. |
| `backend/cmd/server/main.go` route registrations; `backend/internal/domain/campaign/service.go` `toPublicDetail`; `backend/internal/transport/http/campaign_public.go` response projection | Live target revision confirms no Donation runtime route and current Campaign donation action unavailable. | Evidence only; no runtime changes in this Work Unit. |
| `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` `CampaignSuccess`/`Funding`; `frontend/lib/api/public-campaign.ts` `getPublicCampaignDetail` | Current frontend has Campaign detail only and no Donation flow. | Evidence only; future FE must consume the reconciled contract. |
| `AGENTS.md` §§2–3; `docs/kencleng-agentic-workflow.md` §4; `backend/AGENTS.md` §3 | Exact-money, explicit auth, PII safeguards, Tier-0 fences, risk/verification routing. | Protect ledger/locking, crypto/key, auth core, and state-machine writes; future Tier-1 paths need evidence and human review. |

## 11. Files Changed / Files NOT Changed

### Expected in downstream contract reconciliation (not this Planner Run)

| File / area | Change type | Description |
|---|---|---|
| `docs/spec/5-donation/invariants.md` | Adapt | Reconcile Slice 2 rules while preserving exact-money, atomic funding, settlement, submission idempotency, guest privacy, and scope boundaries. |
| `docs/spec/5-donation/threat-model.md` | Adapt | Reassess submit/status/email/token boundaries and unresolved residual risks; do not mark risk accepted without owner. |
| `docs/spec/5-donation/tasks.md` and Donation feature specs | Adapt | Record only coherent guest submit, truthful simulator state, safe status, and notification behavior. |
| `docs/spec/4-campaign/invariants.md` / relevant feature spec | Conditional adapt/reference | Only if Campaign owner requires a cross-reference/change for the approved threshold and eligibility rule. No broad Slice 3 lifecycle change. |
| `api/openapi/donation.yaml` | Adapt | Authored Slice 2 guest contract after material API/Security decisions are resolved. |
| `api/openapi/common.yaml`, `api/openapi/index.yaml` | Conditional source updates | Change only shared components/path registry required by reconciled Donation source. |
| `api/openapi.yaml`, generated frontend API types | Generated | Regenerate only by documented workflow if authored source changes. |

### This Run changed

| File / area | Change type | Description |
|---|---|---|
| `TP-S2-002-006/techplan.md` | New material amendment | This Draft / In Review revision of `TP-S2-002-003`. |
| `TP-S2-002-006/launch-record.md` | New handoff record | Actual dispatch and phase handoff for this Run. |

### Files / areas intentionally untouched by this Run

| File / area intentionally untouched | Why |
|---|---|
| `docs/product/**`, `docs/ui-ux/**` | Current authorities already contain the approved amendment; source edits belong to Product/Design authority. |
| `TP-S2-002-003/techplan.md`, `TP-S2-002-004/report-techplan.md` | Preserve historical plan/report; prior report remains stale for approval of this material amendment and is not hand-patched. |
| OIR artifacts, Exploration evidence, other Runs, manifest/control/events/outcome/work-graph | Preserve source/history and orchestration ownership; participant writes only current Run artifacts. |
| `docs/spec/**`, `api/**`, runtime `backend/**`/`frontend/**`, tests, tracker | This Run amends planning only. Contract, implementation, and project-state writes belong to later authorized phases/owners. |
| Protected Tier-0 ledger/transaction-locking, disbursement state machine, crypto/key handling, and auth core | Explicit `AGENTS.md` fencing; no write authority granted or needed. |
| Manual DB/index application; hand-edited `api/openapi.yaml`/generated frontend types | Human migration/index gate and API generated-artifact rules. |

## 12. Testing Checklist

This Run edits planning artifacts only; no tests or runtime checks were run. The following are obligations for the downstream spec/API reconciliation and, where noted, runtime delivery.

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | Human compares reconciled spec/API scope against current Slice 2 and confirms deferred operations/methods did not re-enter; explicitly checks Slice 3 boundary. | Human | Product scope authority cannot be inferred by schema validation; silent breadth expansion changes MVP. |
| R2 | Owner review of amount contract for min/increment/Rp5.001 boundary, exact-decimal encoding/storage/calculation, no float, and explicit unresolved derived precision. Design review confirms method availability disclosure. | Human | Money validity and payment meaning are material product promises; automated lint cannot set rounding authority or assert truthful UX. |
| R3 | Reconciled invariant/feature/API review for internal-only settlement and coupled success/funding; downstream Testing demonstrates failure atomicity, replay exact-once, and concurrent settlement/no lost increment using appropriately scoped evidence. | Testing | Partial commits or concurrency errors falsify monetary truth; specialized evidence is necessary. Ledger/locking remains Tier-0 fenced. |
| R4 | Human/Product confirms failure is only a clearly labeled backend-configured/fixture demo scenario, with pending copy and no time/payment promise. | Human | Wrong simulation semantics can present a false payment result or induce unintended repeat submission. |
| R4 | Testing exercises persisted pending→terminal behavior and confirms no donor-controlled settlement or auto-resubmit after the approved simulator behavior is implemented. | Testing | Runtime behavior cannot be established from contract review alone; failure can mislead donors or duplicate submissions. |
| R5 | Security/PII and Human owner approve opt-in/verification, no details in verification email, terminal-only notice, hold/delete and retry/retention windows, encryption/HMAC, and sanitized logging. | Human | Guest email is PII/access-adjacent; missing authority can expose status or retain PII indefinitely. |
| R5 | After policy is fixed, Testing verifies email send/no-send/delete behavior at verification and terminal-state boundaries. | Testing | Runtime defects can send status to unverified addresses or retain data past the approved window. |
| R6 | Security/API approve threat controls and response matrix for URL/token exposure, status-only payload, 24-hour expiry, missing/invalid/expired parity across code/body/headers/cache/timing, secret comparison, and abuse controls; decide residual risk. | Human | Guest bearer URL exposes private status; schema checks alone cannot establish anti-enumeration or accept residual risk. |
| R6 | After controls are approved, Testing verifies the chosen parity/security behavior across relevant failure cases. | Testing | Without empirical evidence, subtle response differences or exposure paths may remain. |
| R7 | Campaign/Donation owners review eligibility and threshold contract, including crossing/full accepted pending Donation, rejection after close, and concurrent ordering. | Human | Incorrect eligibility or close ordering changes accepted money and can accidentally expand scope. |
| R7 | Runtime Testing checks threshold boundary, accepted-pending settlement, post-close rejection, and required concurrency behavior while excluding broader Slice 3 result flow. | Testing | Contract review alone cannot establish boundary/race behavior in the implementation. |
| R8 | Contract/API and client acceptance cases: same key/same payload after ambiguous timeout returns original Donation; same key/different payload rejects; deliberate new Donation uses new key; double-click is suppressed. Separately test settlement replay under R3. | Testing | Request retry and settlement replay are distinct duplicate paths; omission can create unintended Donations. |
| R9 | Before any remove/replace, API/Orchestrator owner records in-repo and external consumer/distribution evidence and routes discovered compatibility needs. | Human | Unknown consumers risk breakage; generated-type presence alone cannot prove distribution or absence of use. |
| R9 | Run `cd api && npm run validate`; if sources change, follow `api/README.md` bundle/type generation and inspect source/index/bundle/generated consistency. | Testing | Validation catches reference drift but cannot prove compatibility or runtime behavior. |
| R10 | Human Design/Product review of pending/terminal copy, QRIS vs unavailable method states, terminal email, failed recovery, and non-coercive informational email/Account messaging against `docs/ui-ux/patterns.md` and `design-guidelines.md`; rendered review when the implemented UI materially affects trust. | Human | Copy/color can overstate settlement or verification, or pressure users toward unavailable features; automation cannot decide whether the experience communicates the approved truth. |

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Settlement authority, duplicate transitions, atomic success/funding coupling, and concurrent donations | Money state must not diverge under partial failure, replay, or concurrency; needs specialized downstream Testing evidence. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-2-gap-analysis.md#Area 3 — Donation delivery specs and API evidence`; same file `#Area 4 — Backend live state and security/correctness boundary`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-3-solutioning.md#Recommended material direction` | Yes — R3 remains a correctness floor; runtime behavior is not implemented or tested by this planning Run. |
| Guest submission retry, same-key payload mismatch, and deliberate fresh-key activation | Side-effecting submission has ambiguous timeout and duplicate activation risks distinct from settlement replay. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-2-gap-analysis.md#Area 1 — Product/MVP and slice boundary`; same file `#Area 3 — Donation delivery specs and API evidence`; same file `#Area 5 — Frontend live state and cross-stack surface`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-3-solutioning.md#Recommended material direction` | Yes — O9 resolves policy, but downstream request-level persistence/response evidence remains required. |
| Guest status credential, response parity, URL/log/cache/referrer exposure, and abuse | Unauthenticated bearer access reaches private status; token and anti-enumeration controls and residual risk remain Security/API decisions. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-2-gap-analysis.md#Area 3 — Donation delivery specs and API evidence`; same file `#Area 5 — Frontend live state and cross-stack surface` | Yes — O4/O5 product direction is set while technical controls/risk acceptance remain open. |
| Guest-email PII, verification, retention, encryption/HMAC, and safe logging | Optional notification data can leak status or outlive purpose if controls/windows are unspecified. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-2-gap-analysis.md#Area 3 — Donation delivery specs and API evidence` | Yes — O3 sets product purpose and terminal-only notice; Security/PII windows/control still need owner decisions. |
| Campaign eligibility, threshold crossing, close ordering, and accepted pending donations | Stale detail/concurrent close may permit ineligible submission or contradict full-settlement direction. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-2-gap-analysis.md#Area 3 — Donation delivery specs and API evidence`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/stage-3-solutioning.md#Open / deferred questions` | Yes — O6 policy is resolved; cross-domain contract/runtime ordering still requires owner review and Testing. |

## 13. Open Items

### Active — needs external input or verification

1. **O1 — Amount wire/storage and derived-money precision — Donation/API owner; Product Authority if semantics change.** Human set whole IDR input, min Rp5.000, Rp1 increments, and exact decimal for stored/calculated values. Still specify authored request/response representation and storage precision/scale. Derived-value precision/rounding is unset and tax rules are out of Slice 2; do not invent either. Evidence: `mvp-delivery-slices.md` §5; `OIR-S2-002-001` O1; money best practice.
2. **O2 — Simulator timing and failure-scenario control — Donation delivery owner, Product Authority for any user promise, Design for material presentation.** Human set state ownership, backend-configured/fixture demo-only failure, pending copy, and recovery. Exact simulator timing and implementation details for the backend-controlled scenario remain undecided. Historical 2–5 seconds/5% are not approved; no SLA/estimate is permitted without Product authority. Provide evidence that donor request/UI cannot choose terminal outcome.
3. **O3 — Guest email verification, retention, delivery retry, and controls — Security/PII owner; Donation delivery/API owner for contract.** Product decision is resolved as in Q6/D6, but verification flow/window, post-terminal delivery-retry window, deletion/retention controls, and handling evidence remain open. Keep verified email only for the approved retry need, remove unverified email when its approved window expires, and do not send if unverified. Preserve established encryption/HMAC and safe logging. Human product direction does not accept residual privacy risk.
4. **O4 — URL/token exposure mitigations and residual-risk decision — Security/project authority; API/Donation owner for contract/control evidence; Design if interaction materially changes.** Human set 24-hour URL, difficult-to-guess token, status-only response, and no need after expiry. Still decide carrier/storage, browser history/referrer/log/cache mitigations, token lifecycle/comparison, abuse controls, and risk acceptance through the owner/human gate. Token unpredictability and TTL alone do not close risk.
5. **O5 — Response parity/anti-enumeration — API owner with Security review; Design for any user-visible recovery change.** One generic behavior/copy for missing, wrong/invalid, and expired is resolved. Decide and evidence matching status code, body, headers, cache behavior, material timing, and rate/abuse behavior; resolve historical feature-spec 401 versus OpenAPI 404 conflict. Do not infer technical parity from shared frontend copy.
6. **O7 — Design terminology/source-label review — Design Authority, after O2 semantics.** Review pending “Menunggu hasil simulasi,” success/failed labels, email wording, QRIS-active versus other-methods-unavailable, and failed-state next action against current patterns/principles. Determine whether canonical guidance is enough or a material Design decision is needed. No final visual treatment is selected by this Techplan.
7. **O8 — Historical API consumer/distribution status — API/Orchestrator owner, only if operations are proposed for removal/replacement.** Prior in-repo evidence found generated declarations but no active FE runtime consumer; this Run did not perform external distribution audit. Before a breaking remove/replace, record consumer/version evidence and route any compatibility need to its owner. If no operation is removed/replaced, retain conditional/deferred status.
8. **Campaign/Donation contract ordering at threshold — Campaign + Donation owners; Product Authority only if proposed semantics alter approved policy.** O6 product policy is resolved and must not be reopened absent new contradiction. Define contract ordering/state consistency for crossing, accepted pending donations, concurrent settlements/close, and rejection after close. Do not include broader Slice 3 closure/result work.

### Resolved — retained as decision history

1. ~~**O1 product amount/method direction**~~ **RESOLVED —** IDR whole-Rupiah input, min Rp5.000/increment Rp1; exact-decimal stored/calculated money; display QRIS, GoPay, ShopeePay, bank transfer with QRIS the only active labeled sandbox and all other options non-interactive/unavailable. Human-approved Product/MVP amendment, 2026-09-26. Contract encoding and derived precision remain Active above.
2. ~~**O2 product state/recovery direction**~~ **RESOLVED —** backend simulator owns pending/success/failed; failure only from clearly labeled demo scenario, not donor choice/request; pending says “Menunggu hasil simulasi” with no estimate/payment instruction; failed permits explicit new donation; pending does not auto-resubmit. Human/OIR 2026-09-26. Exact timing/mechanism remains Active above.
3. ~~**O3 product guest field/notification direction**~~ **RESOLVED —** optional non-public-by-default name; optional opt-in email for status only; verify ownership before status/access email; at most one status-only simulation-labeled notice at terminal success/failed, not pending; unverified terminal email held only in Security/PII window then deleted without notice. Human/OIR 2026-09-26. Window/retention/control remains Active above.
4. ~~**O4 product guest revisit direction**~~ **RESOLVED —** status-only guest URL, difficult-to-guess token, 24-hour lifetime; no need after expiry. The “under one hour” real-world reference is rationale only, not a sandbox SLA. Human/OIR 2026-09-26. Security exposure controls and risk decision remain Active above.
5. ~~**O5 public failure-copy direction**~~ **RESOLVED —** absent/invalid/expired link uses one generic public behavior/copy, e.g. “Link status tidak tersedia atau mungkin kedaluwarsa.” Human/OIR 2026-09-26. API/Security transport parity remains Active above.
6. ~~**O6 threshold policy**~~ **RESOLVED —** `max_amount` is closure threshold, not hard cap; crossing Donation accepted full; Donations accepted while eligible settle full even if pending at close; new submissions after close rejected. Human/OIR and Product/MVP amendment, 2026-09-26. Only implementation/contract ordering remains Active above.
7. ~~**O9 submission retry/double-submit policy**~~ **RESOLVED —** same-key ambiguous retry returns same Donation; same key with different payload rejected; new key means a new Donation only after deliberate donor action; client prevents double-click and does not rotate key during ambiguity. Human/OIR and Product/MVP amendment, 2026-09-26. Delivery translation/retry-record lifetime remains in contract work; separate from R3 settlement replay.
8. ~~**Baseline actor and product outcome**~~ **RESOLVED —** guest Donation without Account prerequisite and truthful sandbox status; current Product/MVP authority, retained from prior plan history.
9. ~~**Whole historical Donation scope**~~ **RESOLVED —** defer Account claim/history, donor list, real rails, and unnecessary active payment-method breadth; current Product/MVP authority, with display-only method list exception recorded in D4.
10. ~~**Frontend-only or HTTP-triggered settlement**~~ **RESOLVED —** persisted backend-owned result and no client-callable forged transition; correctness floor retained from prior plan/Exploration.

### Material amendment / phase routing

- Material change: **Yes** — Product/domain behavior, scope, interface semantics, risks, and verification strategy changed.
- Current status: **Draft / In Review**. The Human approval of `TP-S2-002-003` does not approve this amendment.
- Independent Techplan review: **Recommend / route before next Human gate** — fresh Reviewer Session should independently verify faithful mapping of multiple resolved cross-domain Human decisions, exact source authority, unresolved security/privacy/design/API boundaries, and preserved threshold/idempotency/atomicity rules.
- Report: do not generate in this Run. After independent review and any resolution pass converge, regenerate `report-techplan.md` from the canonical template before the next Human approval gate; the prior report stays historical/stale for this revision.
- `CONTRACT_READY`: not claimed; Active O1–O5/O7/O8 (conditional) and Campaign/Donation ordering still need their owners/evidence.
- Build: not started or authorized by this phase handoff.
