# Kencleng — Project Orchestration Overlay

> Status: Canonical project orchestration overlay
> Purpose: Define Kencleng-specific coordination that sits on top of the Harscode feature-development workflow.
> Boundary: This document does **not** define a competing generic per-feature lifecycle.

---

# 1. Authority Boundary

Kencleng uses two complementary layers.

## Harscode owns feature execution

Harscode owns how one feature/task moves through:

```text
Exploration
→ Techplan
→ Build / Patch
→ Code Review
→ Testing
→ Pull Request
```

It also owns default session boundaries, phase responsibilities, patch routing, implementation/report artifacts, and generic engineering best practices.

Do not redefine those mechanics here.

## Kencleng owns project orchestration

This document owns project-specific concerns such as:

- domain development order;
- domain preparation;
- invariants and threat-model expectations;
- project risk tiers;
- human-authority requirements;
- backend/frontend sequencing;
- contract-driven frontend development;
- design readiness;
- mock-first integration state;
- domain-level integration/finalization;
- project-specific manual operations;
- cross-domain coordination.

Principle:

> **Harscode decides how a task moves through its lifecycle. Kencleng decides what project-specific conditions and coordination surround that lifecycle.**

---

# 2. Concern-Specific Source of Truth

Use the source that owns the concern.

| Concern | Authority |
|---|---|
| Domain invariants / threats | `docs/spec/<domain>/invariants.md`, `threat-model.md` |
| Feature behavior / acceptance criteria | `docs/spec/<domain>/features/*.md` |
| API shape | `api/openapi.yaml` |
| Backend architecture | `docs/project/kencleng-backend-tech-stack.md` |
| Frontend architecture | `docs/project/kencleng-frontend-tech-stack.md` |
| Product-design principles | `docs/ui-ux/product-design-principles.md` |
| Selected brand + Product UI direction | `docs/ui-ux/brand-product-ui-brief.md` |
| Reusable UX behavior | `docs/ui-ux/patterns.md` |
| Asset governance | `docs/ui-ux/asset-governance.md` |
| Route/persona mapping | `docs/ui-ux/page-map.md` |
| Selected visual direction evidence | `docs/ui-ux/visual-references/selected-direction/` |
| Component governance | `frontend/components/README.md` |
| Feature lifecycle | Harscode `workflow/` |
| Project orchestration | this document |

There is intentionally no current canonical concrete visual-system document defining exact production colors, fonts, radii, spacing tokens, icon family, or motion values. Those are downstream derivations from the approved brand/Product UI brief.

A visual reference cannot override domain/API truth. A local implementation cannot silently redefine project-wide product-design, UX, or component contracts.

If two authorities genuinely conflict on the same concern, surface the contradiction.

---

# 3. Domain Development Order

Current domain order:

```text
account
→ notification
→ organisasi
→ campaign
→ donation
→ disbursement
```

This ordering reflects product-phase progression and dependency shape.

Notification is intentionally early because later domains depend on it.

Campaign spans multiple business phases, so features inside the domain still follow lifecycle dependencies rather than arbitrary order.

Audit logging is cross-domain infrastructure: storage may be established early, while write sites appear throughout later domains. Applicable feature specs must explicitly identify audit requirements.

Cross-domain invariants have one authoritative owner. Other domains reference them rather than duplicating definitions.

---

# 4. Project Risk Tiering

Risk tier determines correctness/security oversight. It is independent from design readiness and change blast radius.

## Tier 0 — Human-authored / human-paired

Reserved for correctness-critical implementation where autonomous agent implementation is insufficient.

Examples may include:

- donation ledger locking strategy;
- money rounding/calculation core;
- encryption/key-handling core;
- JWT/TOTP security-critical core;
- refresh-token reuse detection core;
- disbursement state-machine core.

Agents may support exploration, critique, test ideas, adversarial review, and proposal drafting. Critical implementation remains human-authored or explicitly human-paired.

## Tier 1 — Agent implementation + proof + human review

Examples include transaction boundaries, balance/status writes, audit-log writes, security-sensitive auth flows, PII handling, authorization-sensitive paths, and other threat-model-raised work.

Requires appropriate executable evidence, independent review/testing, and mandatory human review before merge.

## Tier 2 — Agent implementation + standard verification

Ordinary CRUD, conventional API integration, standard forms, and non-critical feature behavior normally use the standard Harscode lifecycle and project verification.

## Tier 3 — Low-risk agentic work

Straightforward documentation, isolated non-critical styling, seed/sample content, and simple low-risk UI state may qualify.

Tier 3 does not mean globally safe. A visually small change to a shared primitive can still have large blast radius.

---

# 5. Risk Tier Before Build

Every work item must have an explicit risk tier before implementation begins.

Threat models can raise the tier even when implementation appears simple.

Frontend code does not automatically inherit backend Tier 0, but frontend surfaces may independently require higher scrutiny when handling security-sensitive authentication, PII, unsafe content rendering, consequential financial/status presentation, authorization-sensitive interaction, or other threat-model concerns.

Use actual responsibility, not directory name, to determine risk.

---

# 6. Once-Per-Domain Preflight

Before active feature implementation in a new domain:

1. read relevant project/domain material;
2. confirm dependencies and sequencing;
3. create or verify:
   - `docs/spec/<domain>/invariants.md`
   - `docs/spec/<domain>/threat-model.md`
   - `docs/spec/<domain>/tasks.md`
4. ensure tasks reference applicable invariants, risk tiers, audit requirements, and delivery expectations;
5. determine dependency-ordered versus independent work;
6. determine backend/frontend sequencing;
7. confirm OpenAPI readiness for contract-driven frontend work;
8. identify cross-domain prerequisites.

Domain preflight prepares project context. It does not replace Harscode Exploration/Techplan for individual features.

---

# 7. Per-Feature Project Preconditions

Every task still uses the Harscode feature lifecycle.

## Backend feature

Before Build, confirm relevant:

```text
feature spec
applicable invariants
threat-model concerns
risk tier
audit-log requirement
OpenAPI contract
```

are sufficiently defined.

If Exploration discovers a contract gap, resolve it through the authority that owns it rather than patching uncertainty inside implementation code.

## Frontend feature

A frontend work item is organized around a meaningful UI unit such as a page, flow, interaction, or component responsibility. It need not map 1:1 to a backend endpoint.

Before Build, inspect as relevant:

```text
feature/domain spec
OpenAPI
page-map
product-design principles
brand + Product UI brief
UX pattern
asset governance
selected visual references
existing production components
```

Check whether the intended capability already belongs to an existing surface or component. Do not create duplicate routes or near-duplicate broad components merely because a backend task is new.

---

# 8. Frontend Design Readiness

Frontend Exploration classifies material UI work as:

```text
READY
PARTIAL
OPEN
```

using `docs/ui-ux/product-design-principles.md`.

## READY
Established product intent, UX behavior, and relevant canonical direction adequately answer the problem.

## PARTIAL
Core intent is known but limited non-semantic design details remain. Resolve them in Exploration using established principles and record material assumptions when they create precedent.

## OPEN
Material product-design, interaction, information-hierarchy, or brand decisions remain unresolved.

Exploration must resolve enough of:

- user goal;
- information hierarchy;
- primary action;
- flow;
- states;
- responsive intent;
- pattern reuse/new pattern;
- visual asset needs;
- open product decisions;

before Techplan becomes executable.

Human approval is required where product-design or asset governance assigns human authority.

OPEN does not create a new generic workflow phase; it changes what Exploration must resolve.

---

# 9. Brand / Product UI Direction

The approved upstream direction is:

```text
Sunlit Editorial
+
Evidence-Led Optimism
```

Primary experience sequence:

```text
evidence
→ confidence
→ optimism
→ action
```

Frontend work must preserve the distinction between platform fact, organizer information/report, system state, pending information, and reported outcome.

Funding progress, operational progress, and reported outcome must remain distinct.

Public surfaces may be more expressive/editorial. Authenticated operational surfaces should become more disciplined without losing brand continuity.

The selected visual references demonstrate direction; they are not route screenshots to clone.

---

# 10. Concrete Visual-System Readiness

Exact production:

- colors;
- fonts;
- radii;
- spacing tokens;
- shadows;
- icon family;
- motion values;

are currently intentionally OPEN.

Do not preserve old implementation values merely because they already exist, and do not invent a new full visual system inside feature Build work.

A concrete visual-system derivation should be handled as an explicit design/engineering work item using the canonical brand brief.

**Engineering follow-up — intentionally deferred.**

---

# 11. Visual Asset Readiness

Material asset needs must be classified during frontend Exploration.

```text
canonical approved asset exists
→ reuse

asset can be produced by current harness
→ brief → generate → review → approval as required

asset cannot be produced by current harness
→ preserve intent through a clear asset brief / handoff
```

Tool limitations must not silently downgrade meaningful design into generic filler.

Brand-defining assets require human approval before becoming canonical.

Use `docs/ui-ux/asset-governance.md`.

---

# 12. Component Impact Preflight

Changing shared UI/component infrastructure and changing local feature code are different risks.

Before materially modifying broad shared primitives/components, follow `frontend/components/README.md`:

```text
classify change
→ discover consumers
→ identify representative risks
→ verify component
→ verify representative downstream consumers
```

This is independent from feature risk tier.

---

# 13. Backend / Frontend Sequencing

Backend-first is not universal.

Choose sequencing per domain based on contract stability, dependency shape, coordination cost, value of parallelism, and integration risk.

## Backend-first

```text
backend features
→ frontend features
→ integration
```

Prefer when API/domain behavior is still evolving or coordination overhead outweighs parallelism.

## Contract-parallel

```text
stable spec + OpenAPI
       ↓
backend          frontend against contract mocks
       \          /
        integration
```

Prefer when API contract is stable enough and frontend can meaningfully progress against mocks.

Parallelism is an optimization, not a goal.

---

# 14. Frontend Mock-First State

Frontend may be implemented and verified against contract-driven mocks before live backend availability.

Use project state:

```text
FRONTEND_MOCK_VERIFIED
```

It means production UI exists, mocked API behavior follows intended contract, scoped frontend verification has passed, and required rendered inspection was performed.

It does **not** mean live backend integration has been verified.

---

# 15. Mock-First Does Not Guarantee Zero Integration Changes

Aim for mock → real backend without rewriting application behavior, but do not assume this outcome.

If integration requires code changes, identify why: mock drift, OpenAPI drift, misunderstood state semantics, environment/integration behavior, or implementation defect.

Route required code changes through the appropriate Harscode Build/Patch activity.

---

# 16. Cross-Domain Mock Batching

Multiple domains may temporarily reach `FRONTEND_MOCK_VERIFIED` before live integration.

Track this explicitly. Never rely on memory for which domains/endpoints remain mocked or blocked.

Mock-verified is an intermediate project state, not a euphemism for complete.

---

# 17. Project Development States

Suggested vocabulary:

```text
CONTRACT_READY
BACKEND_VERIFIED
FRONTEND_MOCK_VERIFIED
INTEGRATED_VERIFIED
DOMAIN_FINALIZED
DELIVERED
```

These are project orchestration states, not replacements for Harscode phase names.

---

# 18. Integration Verification

When backend and frontend are available together, verify real integration through the actual stack/interface.

For frontend-visible behavior, include real rendered browser exercise.

Playwright is the target browser-automation capability once repository wiring exists. Do not claim Playwright verification when it is not wired or run.

Use browser automation for meaningful interaction, responsive/layout checks, integration verification, and high-value regression coverage. Permanent E2E coverage should be justified by maintenance value.

---

# 19. Domain Finalization

Domain Finalization is a cross-feature integration sweep, not a second implementation lifecycle.

Confirm as applicable:

- expected backend/frontend connections use real services rather than accidental temporary mocks;
- cross-feature flows compose correctly;
- representative real-data loading/error/empty/success states behave correctly;
- financial/status meaning remains clear;
- security/authorization boundaries survive integration;
- representative responsive/rendered behavior is correct;
- canonical product-design/UX/component authority is followed;
- provisional dependencies are intentionally resolved or explicitly accepted.

Any defect requiring code changes routes back through Build/Patch.

---

# 20. Delivery

Delivery occurs when the domain meets intended integrated completion criteria.

Record accurate project state, unresolved accepted risks, intentionally deferred work, and any remaining provisional/mock state.

Project delivery state and Git merge state are separate concerns.

---

# 21. Human Authority

Human effort should be spent where project authority or judgment is genuinely required.

Human approval is mandatory where defined for:

- Tier 0 implementation;
- Tier 1 merge/review;
- product/domain contract changes;
- protected spec/test changes;
- brand-defining visual assets;
- material OPEN product/design decisions;
- manual DB/index application;
- other explicitly protected operations.

Agents must not approve their own work where independent human authority is required.

---

# 22. Reporting Discipline

Distinguish clearly between:

```text
verified
assumed
deferred
not tested
```

Do not optimize reports toward “everything looks done”. A known material risk omitted from a report is worse than a clearly documented limitation.

---

# 23. Relationship to Historical Material

Superseded design prototypes and prior visual-system documents are intentionally not active sources of truth. Git history is the archive.

Do not resurrect old green/token/prototype decisions as current precedent unless a future approved decision independently re-establishes them.
