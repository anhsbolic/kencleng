# Kencleng — Project Orchestration Overlay

> File: `docs/kencleng-agentic-workflow.md`
>
> Status: Current project orchestration policy
>
> Last updated: 2026-09-17
>
> Purpose: Define Kencleng-specific coordination that sits on top of Harscode. This document does **not** define a competing per-feature lifecycle.

## 1. Authority boundary

Harscode owns **how one feature/task moves through its development lifecycle**:

```text
Exploration
→ Techplan
→ Build / Patch
→ Code Review
→ Testing
→ Pull Request
```

Harscode also owns generic phase responsibilities, session boundaries, patch routing, reports, review/testing methodology, and portable engineering best-practices.

Kencleng owns project-specific orchestration:

- Product Authority and current MVP release scope;
- MVP vertical delivery sequencing;
- applicable domain invariants/threat-model expectations;
- project risk tiers and human authority;
- backend/frontend sequencing;
- design-readiness requirements;
- mock-first and real-integration state;
- slice/integration finalization;
- project status tracking;
- project-specific cross-domain/manual coordination.

> **Harscode decides how a task moves through its lifecycle. Kencleng decides what product scope and project conditions surround that lifecycle.**

Concern-specific source routing lives in root `AGENTS.md`. Do not duplicate a second full authority table here.

When a Harscode phase is manually invoked, use its current canonical `workflow/*-prompt.md` entrypoint rather than re-authoring a project-specific copy of the phase instructions.

For CRTV/new work, give the canonical kickoff its normal variables and repository/task context only. Do not add a custom solution-steering prompt that tells the agent what conclusions it should discover.

### Legacy numeric references

Older feature/task docs may contain `§NN` references to previous revisions of this document.

Those numeric references are historical navigation aids, not independent policy. When a number no longer points to the old topic, follow the **named rule and current source owner** instead. Clean stale numeric references when the owning spec/task document is next materially edited.

Do not resurrect superseded workflow behavior solely because an older spec cites an old section number.

## 2. MVP delivery sequencing

Current MVP sequencing is owned by:

- `docs/product/mvp-scope.md` — approved release scope;
- `docs/product/mvp-delivery-slices.md` — approved vertical delivery order.

Approved order:

```text
Slice 1 — Public Campaign Understanding
→ Slice 2 — Guest Donation + Truthful Donation State
→ Slice 3 — Campaign Closure + Persistent Public Result
→ Slice 4 — Accountability Follow-up
```

This replaces the old practice of treating backend domain order as the project delivery order for the current MVP.

Historical numbered domain organization remains useful for semantic ownership and local spec placement:

```text
account
notification
organization
campaign
donation
disbursement
```

but it does **not** mean Account must finish before Campaign, or Campaign before Donation, for MVP delivery.

The governing principle is:

```text
product slice need
→ relevant domain semantics
→ FE + BE delivery needs
→ contract reconciliation
→ implementation + integration evidence
```

Project notes:

- Domain boundaries still own domain semantics and applicable invariants once reconciled.
- Cross-domain invariants have one owner: define the invariant under the domain that owns the authoritative field/state; other domains reference it.
- Audit logging remains cross-domain infrastructure and should be introduced only where the active capability actually requires it.
- Account is outside the baseline MVP critical path until a real scoped capability requires it.
- Seeded/operator-assisted operational setup is allowed where `mvp-scope.md` permits it, but must remain truthful, persisted, authorized, and not masquerade as automation or verification.

## 3. Slice preflight

Before active implementation of a new MVP slice:

1. confirm the slice outcome and boundaries from `docs/product/mvp-delivery-slices.md`;
2. inspect only the Product/Design authorities triggered by that slice;
3. identify the minimum participating domains/capabilities;
4. inspect historical specs/OpenAPI/code/tests as evidence rather than automatic scope;
5. classify material existing decisions `KEEP`, `ADAPT`, `REPLACE`, or `DEFER` where needed;
6. establish/reconcile only the domain invariants, threat concerns, feature specs, and contract needed for the slice;
7. decide backend/frontend sequencing strategy;
8. identify migration/shared-component/cross-domain collision risks;
9. update `docs/project/kencleng-development-tracker.md` with actual slice/project state.

Do **not** require a whole-domain preflight merely because one narrow slice touches that domain.

Slice preflight prepares project context. It does not replace Harscode Exploration/Techplan for a coherent work unit.

## 4. Project risk tiering

Risk tier controls **correctness/security oversight**.

It is independent from:

- frontend design readiness;
- shared-component/change blast radius;
- whether a capability is MVP-critical.

A low business-risk change may have high component blast radius; a high-risk feature may be locally scoped. Evaluate them separately.

### Tier 0 — human-authored / human-paired

Reserved for correctness-critical core implementation, including examples such as:

- donation ledger locking strategy;
- money rounding/calculation core;
- encryption/key-handling core;
- security-critical JWT/TOTP core;
- refresh-token reuse-detection core;
- disbursement state-machine core.

Agents may explore, critique, propose tests, or perform adversarial review. Critical implementation remains human-authored or explicitly human-paired.

Root `AGENTS.md` contains the explicit protected-path fencing.

### Tier 1 — agent implementation + proof + human review

Typical examples:

- transaction boundaries;
- balance/status writes;
- audit-log writes;
- security-sensitive authentication flows;
- PII handling;
- authorization-sensitive paths;
- work raised by a threat model.

Requires appropriate executable evidence, independent review/testing, and human review before merge.

### Tier 2 — standard verified feature work

Typical examples: ordinary CRUD, conventional API integration, standard forms, and non-critical feature behavior.

Normal Harscode lifecycle plus project-specific verification is normally sufficient.

### Tier 3 — low-risk agentic work

Typical examples: straightforward documentation, isolated non-critical styling, seed/sample content, or simple low-risk UI state.

Tier 3 does not waive shared-component blast-radius analysis or other applicable project rules.

### Assignment

Assign the risk tier before Build. Threat-model evidence may raise it.

Frontend work does not automatically inherit a backend Tier-0 designation merely because it displays related data; judge the frontend's actual responsibility.

## 5. Per-feature / per-work-unit project preconditions

Every coherent implementation work unit still follows Harscode.

### Backend

Before Build, confirm as relevant:

```text
active Product/MVP slice
reconciled feature spec
applicable invariants
threat-model concerns
risk tier
audit requirement
reconciled OpenAPI contract
```

If Exploration exposes a product gap, route it to Product Authority. If it exposes a delivery/contract gap, resolve it in the authority that owns that detail. Do not reinterpret missing requirements in production code.

### Frontend

Frontend work is scoped around the meaningful UI unit (page, flow, interaction, component responsibility, or cross-cutting experience foundation), not forced into a 1:1 relationship with backend endpoints or business-domain task order.

Before proliferating feature surfaces, establish **enough frontend experience foundation** for subsequent UI work to remain coherent. Depending on current state, this may include runtime/scaffold readiness, visual tokens and typography, core primitives, relevant shell/navigation, brand expression/interaction language, and a representative rendered slice sufficient to calibrate the experience.

A representative rendered slice may use only the portion of a page needed to validate the system. It does **not** require completing every public route or domain dependency first, and it must not invent unresolved product semantics merely to make the slice look complete.

Once the foundation is sufficient, sequence frontend pages/flows by approved product slice plus real prerequisites: user-journey dependency, shared shell/component dependency, design readiness, contract readiness, integration risk, and delivery value.

A useful default model is:

```text
frontend experience foundation
→ required shell / core interaction foundation
→ approved vertical slices
→ dependency-driven page / flow delivery
→ real integration / reconciliation
```

Before Build, identify active frontend concerns and route them through `frontend/AGENTS.md`. Inspect only authorities triggered by the actual task.

Do not create duplicate routes or near-duplicate broad components just because a backend task is new.

For Codex frontend work, `docs/project/codex-frontend-execution-profile.md` chooses model/client/capability. It does not change lifecycle or product authority.

## 6. Frontend design readiness

Material UI work must be classified during Exploration using `docs/ui-ux/product-design-principles.md`:

```text
READY
PARTIAL
OPEN
```

- **READY:** established intent/pattern/system/reference is sufficient; proceed normally.
- **PARTIAL:** core intent is known; resolve limited gaps using established principles and record material precedent-setting assumptions.
- **OPEN:** material interaction/information-architecture/brand/product-design decisions remain unresolved; Exploration must resolve them before Techplan/Build becomes canonical.

`OPEN` does not introduce a new workflow phase. It changes what Harscode Exploration must accomplish.

Missing **product truth** is different from design ambiguity: route product truth to `docs/product/`, not to visual precedent or frontend inference.

Human approval is required wherever the product-design or asset system assigns human authority.

Visual-asset production details belong to the applicable `docs/ui-ux/` authority; do not duplicate them here. Tool limitations must not silently downgrade a required expressive asset to generic filler.

## 7. Shared frontend component impact

Risk tier and component blast radius are separate dimensions.

Before materially modifying `frontend/components/ui/` or `frontend/components/shared/`, follow `frontend/components/README.md`:

```text
classify change
→ discover consumers/wrappers
→ identify representative risk cases
→ implement
→ verify component contract
→ verify representative downstream consumers
```

Frontend tasks modifying the same broad shared contract should not proceed independently without an intentional coordination plan.

## 8. Backend/frontend sequencing

Backend-first is not universal. Choose strategy per slice/work unit based on contract stability, dependency shape, coordination cost, value of parallelism, and integration risk.

### Backend-first

```text
backend capability
→ frontend feature
→ real integration
```

Prefer when API/business semantics are still evolving or coordination overhead would outweigh parallelism.

### Contract-parallel

```text
reconciled slice contract
       ↓
backend          frontend against contract mocks
       \          /
        real integration
```

Prefer when the contract is stable enough and frontend can meaningfully progress against MSW/mocks.

Parallelism is an optimization, not a goal.

## 9. Project development states

Project status is tracked in `docs/project/kencleng-development-tracker.md`.

Ordinary progress labels may include:

```text
NOT_STARTED
IN_PROGRESS
BLOCKED
NEEDS_RECONCILIATION
```

Evidence-backed milestones are:

```text
CONTRACT_READY
BACKEND_VERIFIED
FRONTEND_MOCK_VERIFIED
INTEGRATED_VERIFIED
SLICE_FINALIZED
DELIVERED
```

Historical tracker rows may still use `DOMAIN_FINALIZED`; for new MVP slice delivery prefer `SLICE_FINALIZED` where that better reflects the actual unit of product completion.

These are project orchestration states, not Harscode phase names.

Do not promote a row to a milestone merely because implementation files or a commit exist.

## 10. Frontend mock-first state

Frontend work may reach `FRONTEND_MOCK_VERIFIED` before live backend integration when:

- production UI exists;
- mocks follow the intended reconciled OpenAPI contract;
- scoped frontend unit/component verification passed;
- required **human rendered acceptance** for material UI was completed;
- explicitly scoped frontend verification is complete.

`FRONTEND_MOCK_VERIFIED` does **not** mean real backend integration was verified.

Multiple slices/features may temporarily remain mock-verified when sequencing makes that valuable, but remaining live-integration dependency must be visible in the development tracker.

The architecture should aim for mock→real integration without rewriting application behavior. If real integration reveals required changes, investigate mock/OpenAPI/state/environment drift and route production patches through Harscode Build/Patch.

Playwright is not required to earn `FRONTEND_MOCK_VERIFIED`. Explicitly requested browser automation may provide additional evidence but does not replace human rendered acceptance.

## 11. Integration verification

When backend and frontend are available together, verify the real interface/stack appropriate to the feature/slice.

For material frontend-visible behavior, a human must exercise the integrated rendered flow in a real browser before `INTEGRATED_VERIFIED`/delivery is claimed.

Playwright is available as an **independent, on-demand browser automation capability**. It is not a Harscode phase entry/exit requirement and must not be run merely because a feature reached Testing or PR.

A commit that happens to touch both frontend and backend is not evidence of `INTEGRATED_VERIFIED` by itself.

## 12. Slice finalization and delivery

After required FE/BE work for a slice reaches real integration, perform a cross-work-unit finalization sweep focused on integration gaps, not a second implementation lifecycle.

Confirm as applicable:

- temporary mocks are no longer unintentionally on production paths;
- important slice flows compose correctly;
- representative real-data loading/error/empty/success states behave correctly;
- financial/status meaning remains clear;
- security/authorization/public-data boundaries survive integration;
- representative responsive/rendered behavior is correct;
- canonical UX/design/component systems are followed;
- operational shortcuts remain truthful and authorized;
- provisional dependencies/assets are resolved or explicitly accepted.

Defects requiring code changes go back through Harscode Build/Patch.

Promote to `SLICE_FINALIZED` only with finalization evidence. Promote to `DELIVERED` only when intended integrated completion/release criteria are met.

Git merge state and product delivery state are different concerns.

## 13. Human authority

Human authority is mandatory where the project assigns it, including:

- Tier-0 core implementation;
- Tier-1 review before merge;
- durable Product Authority changes;
- MVP release-scope/sequencing changes;
- reconciled product/domain semantic changes where required;
- protected spec/test changes;
- brand-defining visual assets;
- material `OPEN` product/design decisions requiring product authority;
- manual rendered acceptance before merge/delivery of material frontend UI;
- manual DB/index application.

Agents may analyze/propose indexes. Applying an index is human-triggered; Tier-0 tables require particular care around locking/query-plan consequences.

An agent must not approve its own decision where independent human authority is required.

## 14. Parallelization and write scopes

Parallel work is allowed only when independence is real.

Check:

- overlapping files;
- shared components;
- database tables/migrations;
- API schemas;
- slice dependencies;
- domain dependencies;
- generated artifacts.

Coordinate migration numbering explicitly when concurrent tasks could collide.

Backend and frontend production writes remain separate scopes by default. Cross-stack contract work is coordinated explicitly rather than allowing one Build to drift across both applications.

The boundary does not prevent reading cross-stack context, integration Testing, or slice finalization.

## 15. Evidence and verification ownership

Kencleng uses **verifiability over trust**.

Reports and PR risk notes distinguish:

```text
verified
assumed
not tested
deferred
```

Claims should point to concrete executable evidence where it exists. A known material risk silently omitted is a process failure.

Do not duplicate a universal test sequence here:

- Harscode owns lifecycle-phase verification responsibilities;
- repository commands/config own executable checks;
- Kencleng adds risk-specific requirements and human acceptance boundaries.

Examples of project-specific evidence include race/concurrency verification, threat-focused security checks, applicable invariant/property tests, hostile-content rendering tests, human rendered acceptance for material frontend UI, optional explicitly requested browser automation, and real-stack integration before integrated/finalized status.

Run the right evidence for the risk rather than expensive categories indiscriminately.

## 16. Project status tracking

Workflow policy and project progress are different information.

Current cross-domain/slice status belongs only in:

```text
docs/project/kencleng-development-tracker.md
```

Do not append dated progress amendments to this workflow document.

Domain `tasks.md` files remain domain-scoped historical/current task views where useful. They do not define current MVP sequencing. If a domain tracker and the cross-project tracker disagree, reconcile against Product/MVP authority plus concrete repository/test evidence.

## 17. Generated task artifacts

Exploration logs, techplans, build reports, review/testing reports, and PR artifacts are task evidence/history. They do not become project-wide policy merely because a prior task contains a decision.

Promote reusable truth into the canonical source that owns the concern.

## 18. Evolution

Keep this overlay stable and Kencleng-specific.

Update it when MVP sequencing, risk/human authority, integration strategy, or project-wide coordination changes.

Do not add:

- generic React/Go practices;
- Harscode phase mechanics;
- one-feature implementation lessons;
- dated progress;
- component/design detail already owned elsewhere.

Rewrite canonical policy when it changes; Git history preserves superseded wording.

## 19. Related documents

Product / release:

- `docs/product/README.md`
- `docs/product/product-overview.md`
- `docs/product/mvp-scope.md`
- `docs/product/mvp-delivery-slices.md`

Project:

- `AGENTS.md`
- `docs/spec/README.md`
- `docs/project/kencleng-backend-tech-stack.md`
- `docs/project/kencleng-frontend-tech-stack.md`
- `docs/project/codex-frontend-execution-profile.md`
- `docs/project/kencleng-development-tracker.md`
- `docs/project/kencleng-repo-setup.md`

Frontend product/design:

- `docs/ui-ux/product-design-principles.md`
- `docs/ui-ux/patterns.md`
- `docs/ui-ux/design-guidelines.md`
- `docs/ui-ux/asset-governance.md`
- `docs/ui-ux/page-map.md`
- `frontend/components/README.md`

Portable workflow/engineering:

- Harscode `workflow/`
- Harscode `best-practices/`