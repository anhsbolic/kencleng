# Tech Plan: C1 — Legitimate Organization representation

> Phase             : Techplan
> Ticket            : none
> Author            : `PARTICIPANT-C1-ENG-PLANNER-001`
> Participant ID    : `PARTICIPANT-C1-ENG-PLANNER-001`
> Role              : Planner
> Model             : `gpt-6-sol`
> Reasoning         : medium
> Session           : `SESSION-C1-ENG-PLANNER-001`
> Created           : 2026-10-09
> Updated           : 2026-10-09
> Target revision   : `e9cb1f31f031ed3905181eb86ad3c3e3af4e248f`
> Workflow revision : `d882b35c88ae7b4adff19abba991e0f2f9d7b7ac`
> Status            : Draft / In Review
> Approach          : One attributable C1 establishment consequence spanning durable Organization context and initial internal Owner relationship; technical direction remains subject to Human approval and unresolved authority decisions.
> Refs              : Stage 7 C1 handoff; Stage 5 confirmed behavior; Stage 6 requirements; Exploration Runs `RUN-C1-ENG-EXPLORATION-001` evidence.

---

## 1. Background

The clean baseline has no C1 route, domain state, persistence, authentication model, or API operation. C1 must let a person establish a Kencleng Organization context and become its initial Kencleng Owner while preserving durable attribution and making clear that this internal role does not prove external legal authority. The Stage 5/6 behavior is binding; this plan does not add Organization review, external verification, or conflict adjudication.

## 2. Scope

**In scope:**
- An experience that presents the required consequence before effect and allows later inspection of the Organization context and the initiating person's internal Owner role.
- A backend establishment capability whose successful product state has durable person attribution and a valid Organization–initial Owner relationship as one consequence.
- Durable preservation of Organization-provided truth and the distinction from review or external verification.
- Conditional handling of a relevant known unresolved representation conflict, without adding conflict discovery or resolution policy.
- A shared API contract and persistence design only after the Open Items below establish the minimum valid Organization context and authenticated-person boundary.

**Out of scope (explicit):**
- Organization review/status taxonomy, external-authority proof or adjudication, Campaign curation, and fund-use verification.
- Organization matching/duplicate detection, conflict discovery/resolution, merge, transfer, dispute handling, and uniqueness guarantees.
- Staff permissions, invitation, delegation, revocation, transfer, or multi-owner behavior.
- Recovery/resume capability; it is optional and must preserve incomplete-versus-success meaning if later selected.
- Money movement, cryptographic key handling, or unrelated platform work.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | C1 may begin without external-authority proof; successful establishment and initial Owner relationship must be attributable to a person. | Stage 5 B1; Stage 6 ER1.1–ER1.3 |
| Q2 | Before effect, communicate Organization establishment, initial Owner formation, internal-only authority meaning, and Organization-provided information status. | Stage 5 B2; Stage 6 ER2.1–ER2.2 |
| Q3 | Successful establishment means a valid Organization context plus valid attributable initial Owner relationship; neither side may appear successful alone. | Stage 5 B3; Stage 6 ER3.1–ER3.2 |
| Q4 | The resulting Organization–Owner meaning and attribution remain durably inspectable after the success moment. | Stage 5 B4; Stage 6 ER4.1–ER4.2 |
| Q5 | Establishment must not upgrade Organization-provided information/evidence into reviewed, independently verified, or Campaign-curated truth. | Stage 5 B5; Stage 6 ER5.1–ER5.2 |
| Q6 | Missing attribution or incomplete/failed formation cannot produce successful establishment meaning. Recovery is not required. | Stage 5 B6; Stage 6 ER6.1–ER6.2 |
| Q7 | If a relevant unresolved representation conflict is already known, C1 must not report normal uncontested success or grant a new initial Owner as though uncontested. C1 does not define how such conflicts are detected or resolved. | Stage 5 B7; Stage 6 ER7.1–ER7.2 |

## 4. Rules & Validation

- **R1 — Attributable success:** Given an establishment request, when it succeeds, then durable state links the establishing person, Organization context, and initial Owner relationship; without a trusted person attribution, it cannot succeed. (Q1)
- **R2 — Consequence before effect:** Given the establishment experience, when the consequential action is submitted, then the required consequences were made available before the state becomes effective; external authority is not claimed or required. (Q2)
- **R3 — Coherent formation:** Given an attempt to establish C1, when either Organization or initial Owner formation is invalid or fails, then neither is exposed as successful; success is observable only when both are valid. (Q3, Q6)
- **R4 — Durable inspection:** Given a successful establishment, when the person later inspects its result, then the Organization context, their Owner relationship, internal authority meaning, and non-verification status remain determinable beyond transient messaging. (Q4)
- **R5 — Truth-class preservation:** Given Organization-provided information/evidence, when C1 succeeds, then its truth class is not changed to reviewed, externally verified, or curated by that transition. (Q5)
- **R6 — Incomplete is not success:** Given missing attribution or failed/incomplete Organization + Owner formation, when processing ends, then no successful establishment meaning is persisted or returned; any exposed incomplete state remains distinguishable. (Q6)
- **R7 — Known conflict boundary:** Given a relevant unresolved representation conflict already known to the system, when C1 processes an establishment, then it does not produce normal uncontested-success meaning or grant initial Owner authority as though representation were uncontested; C1 performs no new matching/adjudication. (Q7)
- **R8 — Server-owned privileged authorization:** Given a request to create an initial Owner relationship, when the backend evaluates the grant, then it enforces the explicitly authorized actor/privilege rule server-side; frontend visibility is not security authority. (Root `AGENTS.md`, High-risk write boundary)

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Model C1 success as one product-semantic consequence: valid Organization context plus valid attributable initial Owner relationship. | Binding requirement | Stage 5/6 define this invariant. This is not itself a database transaction mandate. |
| D2 | Recommend one bounded PostgreSQL transaction for creation of the Organization and initial Owner relationship, with success exposed only after commit, provided both records share the same database boundary. | Recommended; Human approval pending | The neutral baseline is one PostgreSQL-backed monolith, and this is the simplest way to enforce the invariant without adding unrequired pending/recovery states. Keep external calls outside the transaction. If the same-database assumption proves false, stop and return with a revised plan. |
| D2-alt | Staged/pending creation with recovery. | Not selected for this C1 plan | Adds states and user-visible semantics that Stage 5 B6 / Stage 6 ER6 do not require. It remains a future option only if an authorized concrete need emerges. |
| D3 | Require a trusted, durable person attribution at the success boundary; exact authentication timing and mechanism are not selected here. | Required semantic boundary; implementation choice open | Stage 5 B1 and Stage 6 ER1.1–ER1.3 allow starting earlier but prohibit unattributable success. The baseline has no authentication model, and implementing authentication core logic requires explicit Human authorization. |
| D4 | Do not add C1 conflict detection, matching, or adjudication. Honor a known relevant conflict only if an authorized source supplies that signal. | Recommended; source absent in baseline | Stage 5 B7 / Stage 6 ER7.1–ER7.2 explicitly leave detection and resolution outside C1. Do not manufacture conflict status. |
| D5 | Keep UI flow, exact copy, acknowledgement mechanism, and durable inspection surface as design decisions within Stage 5/6 requirements and reusable UX authority. | Open within approved design freedom | The product requirements explicitly preserve these choices; no screen design is approved by this Techplan. |

## 6. Backward Compatibility

- **Existing data:** No C1 schema or product data exists in the baseline, so this plan introduces no legacy-data conversion. Migration details follow the approved persistence design; do not infer domain shape from history.
- **API/contracts/clients:** The current OpenAPI contract has no operations. C1 requires a new contract if an API boundary is selected; method/path, payload, errors, and generated-client changes remain blocked on the Open Items that define valid Organization input and the trusted actor boundary.
- **Migration/deprecation compatibility:** No existing C1 migration or endpoint must be preserved. Any destructive or lossy migration is outside this plan and requires a separate explicit review.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| RISK-1 | Organization context and Owner grant become observably split across partial failure. | Medium | High | Adopt D2 only if both records share one PostgreSQL transaction; verify failure/rollback and success-after-commit behavior. Do not expose success before commit. |
| RISK-2 | Initial Owner is privileged authorization; a defect could grant elevated access. | Medium | High | Root `AGENTS.md` requires explicit Human authorization before implementation of this surface. Keep frontend gates non-authoritative; enforce authorization server-side. No Build may implement the grant before that authorization is recorded. |
| RISK-3 | Baseline has no authentication model, so the person attribution source is unresolved. | High | High | Resolve an authorized trusted-principal integration or explicitly authorize a bounded authentication-core change before implementation. Never accept caller-asserted person identity as proof. |
| RISK-4 | Valid Organization context lacks a defined minimum information/validity boundary. | High | High | Product/pre-engineering owner must resolve the smallest C1 validity meaning before schema and API are fixed; do not infer legal validity or invent a detailed field schema. |
| RISK-5 | No current source reports a known unresolved representation conflict. | High | Medium | Do not add matching/detection. Resolve upstream if a real source/policy becomes necessary; until then the conditional rule applies only to conflicts already known to an authorized system boundary. |
| RISK-6 | Durable state or UI may imply review/external verification from internal Owner formation. | Medium | High | Preserve separate truth-class semantics in data and inspection surfaces; verify rendered and persisted outcomes against R4/R5. |

## 8. Interface Contract

**Persistence/data shape:** Proposed persistence must represent (a) the Organization context, (b) the initial Owner relationship, (c) trusted person attribution, and (d) the fact that C1 did not itself review or externally verify Organization claims. Exact fields, constraints, identifiers, and whether provenance is modeled explicitly cannot be fixed until the Open Item for minimum valid Organization context and the authorization boundary are resolved. Do not add review-status taxonomy. The Organization and initial Owner success consequence must share the approved atomicity boundary.

**API/event/external interface:** No endpoint or event is approved by this Draft. The neutral baseline has `paths: {}` in `api/openapi/index.yaml`; if the implementation uses HTTP, define one establishment operation whose success response represents only committed, attributable Organization+Owner success. Do not return normal-success semantics for incomplete or known-conflict cases. Request/response shape and error categories require the decisions identified in §13 before an executable contract can be approved. Do not introduce external calls in the establishment transaction.

**Cross-layer/business boundary:** The browser may present consequence and render state but cannot authorize Owner formation. Backend is the authority for actor attribution, Owner grant, and C1 success. A later inspection surface reads durable backend state; it must distinguish internal Kencleng Owner authority from external/legal verification and review. If an existing known-conflict signal is supplied, the backend must prevent normal uncontested success; no source exists in the current baseline.

## 9. Architecture / Plan

Recommended shape, conditional on Open Item resolution:

1. Frontend collects the minimum C1 Organization information established by the owning Product decision and presents the Stage 5 B2 consequence before enabling the consequential action. Exact composition remains within approved design freedom.
2. A backend request obtains the actor from a trusted authentication principal, never from an untrusted person-ID field. Before adding the privileged grant, confirm explicit Human authorization for the exact implementation surface.
3. The backend validates the approved minimum Organization context and any already-known relevant conflict signal. C1 does not discover or adjudicate conflicts.
4. In one PostgreSQL transaction, persist the Organization context, attributable initial Owner relationship, and required truth/provenance distinction. Commit before returning successful establishment. On any failure, roll back and return a non-success result; do not create a user-visible half-success state.
5. A durable read/inspection surface retrieves the Organization and Owner meaning from authoritative backend state. The frontend must not infer authority or verification status from transient local state.

This is a planning recommendation, not approved architecture. The current baseline supplies PostgreSQL connectivity but no domain model, auth principal, transaction boundary, conflict source, API, or UI route.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `backend/cmd/server/main.go` — `run`, `healthz` | Current server bootstrap and sole route; no product handler or auth middleware exists. | Add C1 routing/composition only after contract and principal boundary are approved; do not treat health route as a product precedent. |
| `backend/internal/platform/db/db.go` — `Open` | Current `pgxpool` connectivity boundary. | Reuse the pool for the approved repository/transaction path; it currently provides no domain persistence. |
| `backend/internal/domain/` and `backend/migrations/` | Both contain only `.gitkeep`; no product domain/schema exists. | Add only the domain and migration structure needed by an approved C1 contract; no empty architecture layers. |
| `api/openapi/index.yaml`, `api/openapi.yaml` | Source and bundled OpenAPI currently define no paths. | Add an operation only after the API contract is approved; keep bundled output derived from source. |
| `frontend/app/page.tsx`, `frontend/app/layout.tsx` | Existing frontend is only a baseline home page/root layout; no C1 flow or inspection surface exists. | Add route-local C1 experience and later inspection only after requirements/UX boundary are settled; do not infer route or copy from scaffold. |

Build must reopen these live anchors and current authorities before editing. If a material assumption has changed, stop and return for plan reconciliation.

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `backend/` C1 domain/transport/migrations | New, expected | Exact paths and schema depend on Open Item resolution and approved contract. |
| `api/openapi/index.yaml` and generated `api/openapi.yaml` | New contract, conditional | Only if HTTP API is the selected boundary; source remains authoritative. |
| `frontend/` C1 route/components | New, expected | Establishment and durable inspection; exact design/path remain open within requirements. |

| File / area intentionally untouched | Why |
|---|---|
| `docs/product/product-intent.md` and Stage 5/6/7 authorities | Upstream read-only for C1 engineering; no product meaning changes authorized. |
| `docs/product/pilot-3-control-tower.md` | Not needed for this Planner Run; no owning engineering evidence update requested. |
| Existing non-C1 capabilities and money movement | Outside C1 scope and no active contract identified. |
| `../harscode-workspace/best-practices/` | Read-only guidance; no changes in ordinary planning work. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | Backend service/integration test: trusted actor attribution is persisted across Organization and initial Owner; absent principal cannot succeed. Security review verifies no caller-supplied ID substitutes for the trusted principal. | Testing | Attribution is a core C1 invariant and authentication is new in this baseline; skipping risks unattributable or spoofed privileged establishment. |
| R2 | Human rendered walkthrough of the implemented flow against each required consequence before the effect boundary; UI test confirms the consequence is presented before the submit/effect transition. | Human | Meaning and timing are user-facing; code-only checks cannot establish that disclosure is understandable or precedes effect. |
| R3 | PostgreSQL integration test injects failure between persistence steps and confirms rollback/no success; success case confirms both records are committed before success response. | Testing | The invariant spans persistent records and failure ordering; unit-only checks could miss partial database state. |
| R4 | Backend read/inspection integration test verifies the Organization–Owner relation and attribution remain available after a fresh read. | Testing | Transient success can hide missing durable state. |
| R4 | Human rendered review confirms the durable inspection surface communicates internal Owner authority and does not imply external verification. | Human | Code tests cannot establish that the user-facing meaning is truthful and understandable. |
| R5 | Persistence/service tests assert establishment leaves Organization-provided provenance unchanged. | Testing | A silent truth-class upgrade would materially misstate evidence and downstream trust. |
| R5 | Human rendered review checks that product language does not imply review or external verification. | Human | Truth-class meaning is user-facing and cannot be established by data assertions alone. |
| R6 | Failure-path integration/UI tests cover missing attribution and failed formation, confirming no success response/state and a visibly distinguishable non-success outcome if one is exposed. | Testing | Partial formation is a direct success-boundary risk. Recovery tests are required only if recovery is later selected. |
| R7 | If an authorized known-conflict source is in scope, contract/service tests inject its known-conflict signal and assert no normal success or initial Owner grant. No conflict-discovery suite is in scope. | Testing | Conditional requirement cannot be verified from a baseline with no conflict source; skip only if the source remains absent and record that limitation. |
| R8 | Authorization security tests exercise direct backend calls that attempt Owner grant outside the authorized actor/scope; reviewer confirms enforcement is server-side. Run only after explicit Human authorization defines the privileged rule. | Testing | UI-only gating can be bypassed; privileged-grant failure could elevate access. |
| R1 / RISK-3 | Auth-specific security verification for the approved principal/session mechanism, including spoofing and unauthorized requests. Scope/tool selected after the mechanism and Human authorization are established. | Testing | The baseline has no auth model; exact threat cases depend on the approved mechanism and must not be guessed in this Draft. |
| R3 / RISK-1 | If transaction code locks or reads shared rows to decide success, targeted concurrent integration testing follows the concrete design; ensure deterministic lock ordering where multiple rows are locked. | Testing | C1 transaction atomicity is required by the selected recommendation, while lock/isolation needs depend on actual schema and decisions. |

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Initial Owner privileged authorization and person attribution | Security-sensitive new authority and identity boundary; root `AGENTS.md` requires explicit Human authorization before implementation. | `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-EXPLORATION-001/evidence/stage-3-solutioning.md#Material-questions-and-constraints` | Yes — D3/RISK-2/RISK-3; no authorization model exists in baseline. |
| Organization + Owner transaction atomicity | Shared persistent state; partial success and transaction behavior matter. | `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-EXPLORATION-001/evidence/stage-3-solutioning.md#Organization-+-initial-Owner-consistency` | Yes — D2/RISK-1; proposed only, must be checked against final persistence boundary. |
| Known representation conflict | Conditional security/authority boundary; source and detection policy are absent and must not be invented. | `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-EXPLORATION-001/evidence/stage-3-solutioning.md#Known-unresolved-representation-conflict` | Yes — D4/RISK-5; only a pre-existing authorized signal may be consumed. |

## 13. Open Items

### Active — needs external input or verification

1. **Blocking — valid Organization context.** Stage 5/6 require a valid Organization context but do not define its minimum information or validity criteria; Stage 5 also lists detailed Organization field schema outside C1 while Stage 7 allows experience/design to decide fields. This is a material product boundary. Product/pre-engineering owner must establish the smallest C1 validity meaning/required information, or explicitly route it as an engineering-owned bounded decision, before the schema, API, or Build contract is finalized. Do not infer legal registration or external proof requirements.
2. **Blocking — trusted actor/authentication boundary.** The baseline has no authentication model. Resolve which existing/authorized principal supplies person attribution, or explicitly authorize a bounded authentication-core implementation, before implementing successful establishment. Root `AGENTS.md` high-risk write boundary applies to authentication core logic.
3. **Blocking — privileged initial Owner grant authorization.** Human must explicitly authorize implementation of the initial Owner grant/privileged authorization rule on the exact surface before Build changes it. The approved Stage 5/6 behavior and approval of this Techplan do not by themselves satisfy this implementation authorization boundary.
4. **Blocking — transaction boundary confirmation.** D2 assumes Organization and Owner state can share one PostgreSQL transaction. Confirm the approved data/ownership design preserves this assumption; if not, return to planning rather than introducing pending/recovery semantics in Build.
5. **Conditional — known-conflict signal.** No source or policy exists in the baseline. No new conflict discovery is planned. If C1 needs an upstream signal to satisfy ER7.1 in the actual deployment context, the owning Product/pre-engineering authority must identify that signal and its meaning before implementation; do not add matching/adjudication.

### Resolved — retained as decision history

1. ~~**Is recovery required?**~~ **RESOLVED — No.** Stage 5 B6 and Stage 6 ER6 make recovery optional; incomplete state must remain distinct from success if recovery is later chosen.
2. ~~**Does C1 prove real-world representation or require external authority proof?**~~ **RESOLVED — No.** Stage 5 B1/B5 and Stage 6 ER1.3/ER2.2 prohibit this interpretation; C1 creates internal Kencleng authority only.
3. ~~**Should C1 create conflict matching or legal adjudication?**~~ **RESOLVED — No.** Stage 5 B7 and Stage 6 ER7.2 leave these mechanics outside C1; D4 consumes only an already-known authorized signal.
