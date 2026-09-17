# Kencleng — Product Authority Reframe Audit

> Status: Working audit — **not a product or delivery authority**
> Created: 2026-09-17
> Purpose: Record why Kencleng is reframing its specification hierarchy, classify existing artifacts by their future role, and define a reversible migration path before any authority promotion or backend reset.

## 1. Why this audit exists

Kencleng was originally specified through a largely backend-minded, domain-first process:

```text
high-level scope
→ actors/entities
→ business-process phases
→ technical architecture
→ ERD
→ OpenAPI
→ domain specs
→ implementation
```

That sequence produced valuable, detailed artifacts and enabled real backend work. It also caused downstream technical choices to harden too early and to be treated as if they were the product itself.

Real implementation of the Account domain then surfaced new evidence and forced revisions across assumptions, specs, contracts, and supporting artifacts. Separately, the frontend design reboot showed that frontend delivery is naturally page/flow/experience-oriented rather than a mirror of backend domain decomposition.

The resulting lesson is not that detailed specification is bad. It is that **specification should become more detailed as a delivery slice approaches implementation, while durable product truth should remain above frontend/backend decomposition**.

This audit therefore separates:

- what Kencleng **is** and what must be true for users/business;
- how the experience expresses that truth;
- how frontend and backend separately deliver a selected slice;
- the contract where those implementations converge;
- historical technical/specification evidence that remains useful but should no longer automatically outrank newer product authority.

## 2. Problem in the current authority hierarchy

The current repository routing makes `docs/spec/<domain>/...` authoritative for business behavior and feature acceptance criteria, and makes OpenAPI authoritative for API shape. This is internally coherent for domain-oriented backend delivery, but it leaves no higher canonical product layer that owns end-to-end product intent independently of backend decomposition.

The UI/UX authority map likewise routes product/domain semantics back to `docs/spec/<domain>/...`.

That creates three practical problems:

1. **Product truth inherits backend decomposition.** Cross-domain user journeys are forced to reconstruct their intent from several domain specs.
2. **Technical hypotheses become prematurely global.** Exact endpoint shapes, schema fields, persistence ideas, and edge-case policies may be decided far before their delivery slice is exercised.
3. **Frontend delivery lacks a peer specification layer.** Design authority exists, but there is no explicit experience-oriented delivery spec equivalent to backend domain/use-case specs.

## 3. Evidence from the existing repository

### 3.1 Product-level material already exists

`docs/project/kencleng-business-process-overview.md` and `docs/project/kencleng-actors-entities.md` contain substantial product/business reasoning:

- actors and roles;
- organizations, campaigns, donations, and related conceptual entities;
- the broad pre-campaign → on-campaign → post-campaign lifecycle;
- important business rules such as guest donation, role separation, curation conflict-of-interest, and post-campaign accountability.

These are strong sources for a future Product Authority, but they are still marked draft and are mixed with implementation-level detail.

### 3.2 Narrative phase docs mix product and implementation resolution

`kencleng-phase0-detail.md` through `kencleng-phase3-detail.md` contain useful end-to-end behavior, but also encode details such as:

- token TTLs and storage;
- provider-specific identity modeling;
- password breach-list integration behavior;
- concrete database entities/columns;
- encryption/HMAC approaches;
- concurrency implementation notes;
- simulated payment parameters;
- other downstream technical choices.

These documents are valuable evidence and source material, but their mixed granularity makes them unsuitable as the new top-level product authority without extraction/reconciliation.

### 3.3 The old roadmap intentionally optimized for complete pre-development specification

`kencleng-roadmap-next-steps.md` records the previous strategy as a documentation/specification phase that was considered fully closed before development, including completed ERD and all-domain OpenAPI authoring.

That historical decision should be preserved as project evidence. The new model changes the strategy based on implementation learning; it does not rewrite history as if the earlier reasoning never existed.

### 3.4 `docs/spec/` is explicitly domain-first and executable

`docs/spec/README.md` currently defines:

```text
domain invariant
→ threat model
→ task list
→ feature specs
```

and states that the specs own domain/feature behavior and acceptance criteria.

This remains a useful model for **backend delivery/domain verification**, especially for correctness and security work. The issue is the current precedence position, not the existence of the artifacts.

### 3.5 Existing implementation is valuable evidence, not automatic future authority

The backend currently contains a substantial Account implementation and verification evidence. Account Task 01 (registration/email verification) and Task 02 (Google OAuth) are merged; later Account work has exploration/planning and partial schema preparation.

The implementation must therefore not be deleted simply because the authority hierarchy changes. It should be reconciled against the new product/delivery model and classified slice-by-slice as:

```text
KEEP
ADAPT
REPLACE
DEFER
```

## 4. Proposed authority hierarchy

The target model is:

```text
                  PRODUCT / BUSINESS AUTHORITY
                            │
                            ▼
                  EXPERIENCE / DESIGN AUTHORITY
                            │
                            ▼
                       DELIVERY SLICE
                    ┌───────┴───────┐
                    ▼               ▼
          FRONTEND DELIVERY     BACKEND DELIVERY
                SPEC                SPEC
                    └───────┬───────┘
                            ▼
                    CONTRACT RECONCILIATION
                            │
                            ▼
                       SHARED CONTRACT
                    (OpenAPI where applicable)
                    ┌───────┴───────┐
                    ▼               ▼
               Frontend          Backend
                    └───────┬───────┘
                            ▼
                       Integration
                            │
                            ▼
                    Product evidence
                            │
                            └──→ refine upstream authority when evidence warrants
```

### 4.1 Product / Business Authority owns

- product purpose and boundaries;
- actors/personas in business terms;
- major user/business capabilities;
- end-to-end journeys and outcomes;
- core business concepts and relationships;
- fundamental lifecycle semantics;
- business permissions and durable invariants;
- trust/accountability semantics;
- material product decisions and deliberately unresolved product questions.

It should normally avoid:

- routes/components;
- endpoint/method names;
- database/table/column design;
- repository/package structure;
- framework-specific behavior;
- detailed technical mechanisms unless the mechanism itself is a genuine product constraint.

### 4.2 Experience / Design Authority owns

The current `docs/ui-ux/` generation remains valuable and should continue to own:

- product-design principles;
- persona-to-surface mapping;
- interaction patterns;
- design language and visual grammar;
- information hierarchy principles;
- asset truthfulness/governance;
- responsive/accessibility/design behavior.

It projects product truth into experience but does not redefine business semantics.

### 4.3 Frontend Delivery Spec owns

For a selected surface/flow:

- persona/user goal;
- experience scope;
- information hierarchy;
- user actions;
- meaningful loading/success/empty/error states;
- responsive and accessibility intent;
- frontend-side data dependencies;
- asset dependencies;
- frontend acceptance criteria;
- material design ambiguity still requiring a decision.

It references product truth rather than duplicating business rules.

### 4.4 Backend Delivery Spec owns

For the backend responsibilities needed by a selected capability/slice:

- application/domain behavior needed for delivery;
- invariant enforcement;
- authorization and security behavior;
- transaction/concurrency requirements;
- persistence/integration requirements;
- backend acceptance and verification criteria;
- proposed service/API capabilities needed by the slice.

Domain decomposition remains local to the backend and does not need to mirror frontend decomposition.

### 4.5 Shared Contract owns convergence, not product truth

OpenAPI remains the executable shared interface where HTTP/API convergence is required, but exact contract detail should normally harden when a real delivery slice needs it.

The target relationship is:

```text
product truth
+ frontend data/interaction needs
+ backend capability/security constraints
→ contract reconciliation
→ OpenAPI ready
```

rather than treating a pre-authored whole-product OpenAPI as immutable upstream product authority.

## 5. Future role of existing artifacts

No existing artifact is deleted merely because the hierarchy changes.

### 5.1 Product-truth source candidates

Use as extraction/reconciliation inputs for the new Product Authority:

- `README.md` product framing;
- `docs/project/kencleng-business-process-overview.md`;
- `docs/project/kencleng-actors-entities.md`;
- product/business portions of `kencleng-phase0-detail.md` through `kencleng-phase3-detail.md`;
- relevant product discoveries preserved in implemented specs/code/tests when they represent real learning.

These sources remain references until specific truth is deliberately promoted into the new Product Authority.

### 5.2 Active experience/design authority

Retain as active, subject to later routing updates toward the new Product Authority for product semantics:

- `docs/ui-ux/README.md`;
- `product-design-principles.md`;
- `page-map.md`;
- `patterns.md`;
- `brand-product-ui-brief.md`;
- `design-guidelines.md`;
- `asset-governance.md`.

### 5.3 Delivery / technical reference

Initially downgrade from global product authority to delivery/reference status pending slice-by-slice reconciliation:

- `docs/spec/1-account` through `docs/spec/6-disbursement`;
- domain invariants that mix durable business truth with current backend design;
- domain threat models;
- domain task lists and feature specs;
- `api/openapi/*.yaml` and bundled `api/openapi.yaml`;
- `kencleng-erd.md`;
- phase-doc technical details;
- existing backend/frontend technical architecture docs where they govern engineering rather than product behavior.

"Reference" does **not** mean "assumed wrong". It means "not automatically promoted above the new product authority without reconciliation."

### 5.4 Implementation / workflow evidence

Retain as evidence:

- current backend Account implementation and tests;
- migrations;
- frontend foundation implementation;
- `.local-agents` / Harscode task artifacts;
- historical PRs and Git history;
- previous roadmap/reboot/migration decision records.

Implementation evidence may reveal missing or incorrect upstream assumptions, but implementation does not silently rewrite product authority.

### 5.5 Orchestration / architecture authority

Keep active unless contradicted by the new hierarchy:

- Harscode operational workflow;
- `docs/kencleng-agentic-workflow.md`;
- stack architecture documents;
- project tracker;
- integration map.

These will require routing updates after the Product Authority model is validated.

## 6. Migration strategy — incremental, not a documentation rewrite

### Stage A — Authority scaffolding

1. Complete this audit.
2. Define the smallest useful Product Authority structure.
3. Explicitly state authority boundaries and reference status of old specs/contracts.
4. Do **not** rewrite all domain specs or OpenAPI.

### Stage B — Product Authority v1

Build a coherent whole-product view at the product/business level:

- purpose and non-goals;
- actors/personas;
- capability map;
- end-to-end journeys;
- core concepts;
- major lifecycles;
- fundamental rules/invariants;
- trust/accountability model;
- capability dependencies;
- explicit unresolved questions.

The goal is end-to-end clarity, not implementation completeness.

### Stage C — Forward probe: Public Campaign Detail

Use Product Authority + Design Authority to derive the first real delivery slice.

Expected exercise:

```text
product intent
→ experience intent
→ frontend delivery spec
→ backend/data capability needs
→ compare with existing Campaign specs/OpenAPI
→ reconcile only relevant contract
→ determine readiness
```

This tests whether the new hierarchy supports a not-yet-implemented cross-stack experience without big-upfront respecification.

### Stage D — Backward probe: Account Registration + Email Verification

Start from the new product truth and derive what the slice actually requires, then compare it against existing Account specs, OpenAPI, backend code, migrations, and tests.

Classify the existing implementation:

```text
KEEP    — aligned and reusable as-is
ADAPT   — core survives; contract/model/behavior needs bounded revision
REPLACE — foundational assumptions conflict with current product truth
DEFER   — valid work, but not required by the current delivery path
```

This tests whether existing development can be salvaged without allowing previous implementation to dictate the new product model.

### Stage E — Decide backend soft-reset scope

Only after the Account probe.

Possible outcomes:

- no reset — retain current backend and adapt incrementally;
- scoped reset — preserve platform/scaffold but rework selected Account/domain implementation;
- delivery reset — freeze current backend as reference and build new slice-driven implementation paths, selectively porting surviving code;
- broader reset — only if evidence shows systemic architectural assumptions conflict with product truth.

Do not delete functioning security/correctness work merely to achieve conceptual cleanliness.

### Stage F — Promote repository routing

After 1–2 probes validate the model:

- promote Product Authority to canonical;
- update root `AGENTS.md` routing;
- update `docs/ui-ux/README.md` product-semantic routing;
- revise `docs/spec/README.md` so domain specs are delivery/domain authorities rather than whole-product authority;
- update API/OpenAPI authority wording;
- reconcile tracker/orchestration wording;
- mark historical roadmap/spec-completion claims as historical context rather than current development posture.

## 7. Guardrails for the reframe

### 7.1 No second active source of truth

The migration may temporarily contain draft candidate documents, but every candidate must clearly state that it is not canonical until promoted.

Once promoted, the new Product Authority owns product/business truth. Older material remains evidence/reference and must not compete for precedence.

### 7.2 No speculative full rewrite

Do not rewrite every domain spec, endpoint, schema, frontend surface, or component before real delivery requires it.

### 7.3 Preserve security/correctness evidence

Account implementation discoveries, invariant tests, threat analysis, concurrency work, authentication protections, and similar evidence remain valuable even if their surrounding contract changes.

### 7.4 Progressive commitment

The closer a decision is to implementation and the more expensive it is to get wrong, the more detailed its specification may become.

Whole-product future areas remain intentionally lower resolution until their delivery path approaches.

### 7.5 Product and delivery discovery are bidirectional

Implementation can reveal missing product truth. When that happens:

```text
implementation evidence
→ explicit product/design decision
→ authority update
→ downstream reconciliation
```

Do not silently reverse authority and treat existing implementation as the requirement.

## 8. Immediate working sequence

The current intended sequence is:

```text
1. Audit existing authority and evidence            ← current
2. Draft Product Authority v1
3. Human review / reconcile product-level questions
4. Probe Public Campaign Detail (forward derivation)
5. Probe Account Registration + Email Verification (backward reconciliation)
6. Decide whether any backend soft-reset is warranted
7. Promote repository routing after the model survives the probes
8. Resume normal Harscode CRTV on real development
```

No production implementation or existing API/spec deletion should occur during Steps 1–3.

## 9. Audit conclusion

Kencleng does not need a clean-room rewrite. It needs an **authority reframe**.

The previous domain-first/spec-first generation remains valuable as a rich bank of product reasoning, security analysis, technical hypotheses, contracts, tests, and implementation evidence.

The new development model should place a clean product/business authority above that material, derive experience and delivery slices progressively, and harden frontend/backend contracts when those slices become real.
