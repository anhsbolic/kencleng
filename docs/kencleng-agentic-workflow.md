# Kencleng — Project Orchestration Overlay

> File: `docs/kencleng-agentic-workflow.md`
>
> Status: Current project orchestration policy
>
> Last updated: 2026-09-14
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

- domain order and preparation;
- invariants/threat-model expectations;
- project risk tiers and human authority;
- backend/frontend sequencing;
- design-readiness requirements;
- mock-first and real-integration state;
- domain finalization/delivery;
- project status tracking;
- project-specific cross-domain/manual coordination.

> **Harscode decides how a task moves through its lifecycle. Kencleng decides what project conditions surround that lifecycle.**

Concern-specific source routing lives in root `AGENTS.md`. Do not duplicate a second full authority table here.

When a Harscode phase is manually invoked, use its current canonical `workflow/*-prompt.md` entrypoint rather than re-authoring a project-specific copy of the phase instructions. Kencleng overlays only the project context it actually owns.

### Legacy numeric references

Older feature/task docs may contain `§NN` references to previous revisions of this document.

Those numeric references are historical navigation aids, not independent policy. When a number no longer points to the old topic, follow the **named rule and current source owner** instead. Clean stale numeric references when the owning spec/task document is next materially edited.

Do not resurrect superseded workflow behavior solely because an older spec cites an old section number.

## 2. Domain development order

Current domain order:

```text
account
→ notification
→ organization
→ campaign
→ donation
→ disbursement
```

The order reflects dependency/product progression, not a rule that every task in one domain must finish before any work in the next begins.

Project notes:

- Notification is intentionally early because later domains depend on it.
- Campaign contains internal lifecycle dependencies; tasks still follow their own dependency order.
- Audit logging is cross-domain infrastructure. Every applicable feature spec must explicitly state whether an audit event is required and what is recorded.
- Cross-domain invariants have one owner: define the invariant under the domain that owns the authoritative field/table; other domains reference it.

## 3. Once-per-domain preflight

Before active implementation in a new domain:

1. confirm dependencies and domain sequencing;
2. create/verify the numbered domain's:
   - `invariants.md`;
   - `threat-model.md`;
   - `tasks.md`;
3. ensure tasks identify relevant invariants, risk tier, audit requirements, dependency/parallel constraints, and meaningful delivery criteria;
4. decide backend/frontend sequencing strategy;
5. confirm OpenAPI readiness before contract-driven frontend work;
6. identify cross-domain prerequisites and migration/shared-component collision risks;
7. update `docs/project/kencleng-development-tracker.md` with the actual domain state.

Domain preflight prepares project context. It does not replace Harscode Exploration/Techplan for a feature.

## 4. Project risk tiering

Risk tier controls **correctness/security oversight**.

It is independent from:

- frontend design readiness;
- shared-component/change blast radius.

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

Frontend work does not automatically inherit a backend Tier-0 designation merely because it displays related data; judge the frontend's actual responsibility (PII, auth-sensitive UI, unsafe rendering, consequential money/status presentation, etc.).

## 5. Per-feature project preconditions

Every feature still follows Harscode.

### Backend

Before Build, confirm as relevant:

```text
feature spec
applicable invariants
threat-model concerns
risk tier
audit requirement
OpenAPI contract
```

If Exploration exposes a contract gap, resolve it in the authority that owns the gap before proceeding. Do not reinterpret the missing requirement in production code.

### Frontend

Frontend work is scoped around the meaningful UI unit (page, flow, interaction, or component responsibility), not forced into a 1:1 relationship with backend endpoints.

Before Build, identify the active frontend concerns and route them through `frontend/AGENTS.md`. Feature/domain behavior and the relevant API shape are common inputs for contract-driven feature work; page-map/UX, product-design, visual, asset, prototype/reference, and broad component authorities are **conditional concerns**, not a preload checklist. Inspect only the authorities triggered by the actual feature and decisions at hand.

Do not create duplicate routes or near-duplicate broad components just because a backend task is new.

For Codex frontend work, `docs/project/codex-frontend-execution-profile.md` chooses model/client/capability. It does not change the lifecycle or product authority.

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

Human approval is required wherever the product-design or asset system assigns human authority.

Visual-asset production details belong to `docs/ui-ux/brand-and-visual-assets.md`; do not duplicate them here. Tool limitations must not silently downgrade a required expressive asset to generic filler.

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

Backend-first is not universal. Choose the strategy per domain based on contract stability, dependency shape, coordination cost, value of parallelism, and integration risk.

### Backend-first

```text
backend features
→ frontend features
→ real integration
```

Prefer when the API/business semantics are still evolving or coordination overhead would outweigh parallelism.

Account historically used this approach for much of its initial work.

### Contract-parallel

```text
stable spec + OpenAPI
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
DOMAIN_FINALIZED
DELIVERED
```

These are project orchestration states, not Harscode phase names.

Do not promote a row to a milestone merely because implementation files or a commit exist.

## 10. Frontend mock-first state

Frontend work may reach `FRONTEND_MOCK_VERIFIED` before live backend integration when:

- production UI exists;
- mocks follow the intended OpenAPI contract;
- scoped frontend unit/component verification passed;
- required **human rendered acceptance** for material UI was completed;
- the explicitly scoped frontend verification is complete.

`FRONTEND_MOCK_VERIFIED` does **not** mean real backend integration was verified.

Multiple domains/features may temporarily remain mock-verified when that sequencing is valuable, but every such state and remaining live-integration dependency must be visible in the development tracker.

The architecture should aim for mock→real integration without rewriting application behavior, but this is a design goal rather than an assumption. If real integration reveals required application changes, investigate mock/OpenAPI/state/environment drift and route production patches through Harscode Build/Patch.

Playwright is not required to earn `FRONTEND_MOCK_VERIFIED`. If a human explicitly requested browser automation for the relevant behavior, its result may be included as additional evidence, but it does not replace human rendered acceptance.

## 11. Integration verification

When backend and frontend are available together, verify the real interface/stack appropriate to the feature.

For material frontend-visible behavior, a human must exercise the integrated rendered flow in a real browser before `INTEGRATED_VERIFIED`/delivery is claimed.

Playwright is available as an **independent, on-demand browser automation capability**. It is not a Harscode phase entry/exit requirement and must not be run merely because a feature reached Testing or PR.

A human decides when a browser behavior is valuable enough to automate and explicitly requests that work (or creates a dedicated browser-regression/automation task). Browser automation is scoped by the behavior it protects—smoke, feature flow, known regression, or deliberate broader journey—not by a Harscode task-folder path.

A commit that happens to touch both frontend and backend is not evidence of `INTEGRATED_VERIFIED` by itself.

## 12. Domain finalization and delivery

After required backend/frontend work reaches real integration, perform a cross-feature finalization sweep focused on integration gaps, not a second implementation lifecycle.

Confirm as applicable:

- temporary mocks are no longer unintentionally on production paths;
- important cross-feature flows compose correctly;
- representative real-data loading/error/empty/success states behave correctly;
- financial/status meaning remains clear;
- security/authorization boundaries survive integration;
- representative responsive/rendered behavior is correct;
- canonical UX/design/component systems are followed;
- provisional dependencies/assets are resolved or explicitly accepted.

Defects requiring code changes go back through Harscode Build/Patch.

Promote to `DOMAIN_FINALIZED` only with finalization evidence. Promote to `DELIVERED` only when the intended integrated completion/release criteria are met.

Git merge state and project delivery state are different concerns.

## 13. Human authority

Human authority is mandatory where the project assigns it, including:

- Tier-0 core implementation;
- Tier-1 review before merge;
- established product/domain semantic changes;
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
- domain dependencies;
- generated artifacts.

Coordinate migration numbering explicitly when concurrent tasks could collide.

Backend and frontend production writes remain separate scopes by default. Cross-stack contract work is coordinated explicitly rather than allowing one Build to drift across both applications.

The boundary does not prevent reading cross-stack context, integration Testing, or project finalization.

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

Current cross-domain status belongs only in:

```text
docs/project/kencleng-development-tracker.md
```

Do not append dated progress amendments to this workflow document.

Domain `tasks.md` files remain domain-scoped task views. If a domain tracker and the cross-domain tracker disagree, reconcile against concrete repository/test evidence; do not silently keep two incompatible current statuses.

## 17. Generated task artifacts

Exploration logs, techplans, build reports, review/testing reports, and PR artifacts are task evidence/history. They do not become project-wide policy merely because a prior task contains a decision.

Promote reusable project truth into the canonical source that owns the concern.

## 18. Evolution

Keep this overlay stable and Kencleng-specific.

Update it when domain sequencing, risk/human authority, integration strategy, or project-wide coordination changes.

Do not add:

- generic React/Go practices;
- Harscode phase mechanics;
- one-feature implementation lessons;
- dated progress;
- component/design detail already owned elsewhere.

Rewrite canonical policy when it changes; Git history preserves superseded wording.

## 19. Related documents

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
- `docs/ui-ux/brand-and-visual-assets.md`
- `docs/ui-ux/page-map.md`
- `docs/ui-ux/prototype-reference.md`
- `docs/ui-ux/design-reference-usage.md`
- `frontend/components/README.md`

Portable workflow/engineering:

- Harscode `workflow/`
- Harscode `best-practices/`
