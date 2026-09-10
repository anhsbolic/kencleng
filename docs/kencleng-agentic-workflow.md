# Kencleng — Project Orchestration Overlay

> File: `docs/kencleng-agentic-workflow.md`
>
> Status: Draft v2
>
> Last updated: 2026-09-10
>
> Purpose: Define Kencleng-specific orchestration that sits on top of the Harscode feature-development workflow.
>
> This document does **not** define a competing per-feature development lifecycle.

---

# 1. Authority Boundary

Kencleng uses two complementary layers.

## Harscode owns feature execution

Harscode owns **how one feature/task is executed**, including:

```text
Exploration
→ Techplan
→ Build / Patch
→ Code Review
→ Testing
→ Pull Request
```

It also owns:

* default session boundaries;
* phase responsibilities;
* patch routing;
* implementation/report artifacts;
* generic engineering best practices;
* generic review/testing methodology.

Do not redefine those mechanics here.

---

## Kencleng owns project orchestration

This document owns Kencleng-specific concerns such as:

* domain development order;
* domain preparation;
* invariants and threat-model expectations;
* project risk tiers;
* human-authority requirements;
* backend/frontend sequencing;
* contract-driven frontend development;
* mock-first integration state;
* product/design readiness;
* domain-level integration/finalization;
* project-specific manual operations;
* cross-domain coordination.

Principle:

> **Harscode decides how a task moves through its lifecycle. Kencleng decides what project-specific conditions and coordination surround that lifecycle.**

---

# 2. Concern-Specific Source of Truth

Use the source that owns the concern.

| Concern                                | Authority                                             |
| -------------------------------------- | ----------------------------------------------------- |
| Domain invariants / threats            | `docs/spec/<domain>/invariants.md`, `threat-model.md` |
| Feature behavior / acceptance criteria | `docs/spec/<domain>/features/*.md`                    |
| API shape                              | `api/openapi.yaml`                                    |
| Backend architecture                   | `docs/project/kencleng-backend-tech-stack.md`         |
| Frontend architecture                  | `docs/project/kencleng-frontend-tech-stack.md`        |
| Product-design principles              | `docs/ui-ux/product-design-principles.md`             |
| UX behavior                            | `docs/ui-ux/patterns.md`                              |
| Visual system                          | `docs/ui-ux/design-guidelines.md`                     |
| Brand / visual assets                  | `docs/ui-ux/brand-and-visual-assets.md`               |
| Route/persona mapping                  | `docs/ui-ux/page-map.md`                              |
| Prototype authority                    | `docs/ui-ux/prototype-reference.md`                   |
| Component governance                   | `frontend/components/README.md`                       |
| Feature lifecycle                      | Harscode `workflow/`                                  |
| Project orchestration                  | this document                                         |

A visually complete prototype cannot override domain/API truth.

A local implementation cannot silently redefine a project-wide UX/component contract.

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

Notable decisions:

## Notification is intentionally early

Notification is used by later domains and is cheaper to establish before those integrations depend on it.

## Campaign spans multiple business phases

Campaign work still follows product lifecycle dependencies internally.

Being in the same domain does not mean its features can be built in arbitrary order.

## Audit logging is cross-domain infrastructure

Audit-log storage may be established early, but write sites appear throughout later domains.

Every applicable feature spec must explicitly answer:

```text
Audit log entry?
Yes / No

If yes:
- event
- relevant recorded fields
```

Do not rely on later implementers remembering this implicitly.

## Cross-domain invariants have one owner

When an invariant crosses domains, define it once under the domain owning the underlying authoritative field/table.

Other domains reference it.

Do not duplicate invariant definitions.

---

# 4. Project Risk Tiering

Risk tier determines **correctness/security oversight**.

It is independent from:

```text
design readiness
```

and:

```text
change blast radius
```

All three may independently increase required scrutiny.

---

## Tier 0 — Human-authored / human-paired

Reserved for correctness-critical implementation where agent autonomy is not sufficient.

Examples include:

* donation ledger locking strategy;
* money rounding/calculation core;
* encryption/key-handling core;
* JWT/TOTP security-critical core;
* refresh-token reuse detection core;
* disbursement state-machine core.

Agent use is allowed for:

* exploration;
* critique;
* test ideas;
* adversarial review;
* proposal drafting.

The critical implementation itself remains human-authored or explicitly human-paired.

---

## Tier 1 — Agent implementation + proof + human review

Examples:

* transaction boundaries;
* balance/status writes;
* audit-log writes;
* security-sensitive auth flows;
* PII handling;
* authorization-sensitive paths;
* other work raised by the threat model.

Requires:

* appropriate executable evidence;
* independent review/testing;
* mandatory human review before merge.

---

## Tier 2 — Agent implementation + standard verification

Examples:

* ordinary CRUD;
* conventional API integration;
* standard forms;
* non-critical feature behavior.

Standard Harscode lifecycle and project verification are normally sufficient.

Human review may be sampled or requested when something unusual emerges.

---

## Tier 3 — Low-risk agentic work

Examples may include:

* straightforward documentation;
* isolated non-critical styling;
* seed/sample content;
* simple low-risk UI state.

Tier 3 does **not** mean:

```text
safe to change anything globally
```

A Tier 3 visual change to a foundational shared primitive may still have high blast radius and require broad downstream verification.

---

# 5. Risk Tier Is Assigned Before Build

Every feature/work item must have an explicit risk tier before implementation begins.

The threat model can raise the tier even when implementation looks simple.

Frontend code does not automatically inherit Tier 0 merely because it displays data from a Tier 0 backend operation.

However, a frontend surface may independently deserve higher scrutiny when it handles:

* security-sensitive authentication behavior;
* PII;
* unsafe content rendering;
* consequential financial/status presentation;
* authorization-sensitive interaction;
* other threat-model concerns.

Use actual responsibility, not directory name, to determine risk.

---

# 6. Once-Per-Domain Preflight

Before active feature implementation in a new domain:

1. Read relevant project/reference material.
2. Confirm domain dependencies and sequencing.
3. Create or verify:

```text
docs/spec/<domain>/invariants.md
docs/spec/<domain>/threat-model.md
docs/spec/<domain>/tasks.md
```

4. Ensure tasks reference appropriate:

   * invariants;
   * risk tiers;
   * audit requirements;
   * meaningful delivery/KPI expectations where the project tracks them.
5. Determine which tasks are dependency-ordered versus independently executable.
6. Determine backend/frontend sequencing for the domain.
7. Confirm OpenAPI readiness for any contract-driven frontend work.
8. Identify cross-domain prerequisites before parallel execution starts.

Domain preflight prepares the project context.

It does not replace Harscode Exploration/Techplan for individual features.

---

# 7. Per-Feature Project Preconditions

Every task still uses the Harscode feature lifecycle.

Kencleng adds project-specific inputs.

## Backend feature

Before Build, confirm the relevant:

```text
feature spec
applicable invariants
threat-model concerns
risk tier
audit-log requirement
OpenAPI contract
```

are sufficiently defined.

If Exploration discovers a project-contract gap, resolve that gap through the appropriate authority before continuing Build.

Do not reinterpret the requirement inside implementation code.

---

## Frontend feature

A frontend work item is organized around the meaningful UI unit:

```text
page
flow
interaction
component responsibility
```

It is **not required to be 1:1 with a backend endpoint**.

Before Build, inspect as relevant:

```text
feature/domain spec
OpenAPI
page-map
UX pattern
design guidelines
brand/assets
prototype authority
existing production components
```

Check whether the intended capability already belongs to an existing page or component.

Do not create duplicate routes or near-duplicate broad components merely because a backend task is new.

---

# 8. Frontend Design Readiness

Frontend Exploration must classify material UI work as:

```text
READY
PARTIAL
OPEN
```

using `docs/ui-ux/product-design-principles.md`.

## READY

Existing product intent, UX pattern, visual system, assets, and/or approved precedent sufficiently answer the problem.

Proceed through the normal Harscode lifecycle.

## PARTIAL

The core intent is known, but small design decisions remain.

Resolve them during Exploration using established principles and precedents.

Record material assumptions when they establish precedent.

Do not escalate ordinary senior-frontend judgment unnecessarily.

## OPEN

The work contains unresolved product-design or interaction decisions that would materially affect the resulting experience.

Use the Harscode Exploration phase for **design exploration** before Techplan becomes executable.

Resolve:

* user goal;
* information hierarchy;
* primary action;
* flow;
* states;
* responsive intent;
* pattern reuse/new pattern;
* visual asset needs;
* open product decisions.

Human approval is required where `product-design-principles.md` or the asset system assigns human design authority.

OPEN does not create a new generic workflow phase.

It changes what Exploration must resolve.

---

# 9. Visual Asset Readiness

Material visual-asset needs must be classified during frontend Exploration.

```text
canonical asset exists
→ reuse

asset can be produced by current harness
→ brief → generate → review → approval as required

asset cannot be produced by current harness
→ brief → ready-to-use generation prompt → human/tool handoff
```

Tool limitations must not silently downgrade the design into generic filler.

Temporary assets must remain explicitly:

```text
PROVISIONAL
```

Brand-defining assets require human approval before becoming canonical.

See:

```text
docs/ui-ux/brand-and-visual-assets.md
```

---

# 10. Component Impact Preflight

Changing a shared component and changing local feature code are not equivalent risks.

Before materially modifying:

```text
frontend/components/ui/
frontend/components/shared/
```

follow the consumer-impact process from:

```text
frontend/components/README.md
```

This is independent from feature risk tier.

Example:

```text
simple Button spacing change
```

may be low security/business risk but high blast radius.

Conversely:

```text
Tier 1 feature-specific security UI
```

may have high risk but small component blast radius.

Do not collapse these dimensions into one score.

---

# 11. Backend / Frontend Sequencing

Backend-first is not a universal project rule.

Choose sequencing per domain based on:

* contract stability;
* dependency shape;
* coordination cost;
* value of parallelism;
* integration risk.

Possible strategies:

## Backend-first

```text
backend features
→ frontend features
→ integration
```

Prefer when:

* API behavior is still evolving;
* business/domain semantics are complex;
* coordination overhead would outweigh parallelism.

`account` historically used this strategy.

## Contract-parallel

```text
stable spec + OpenAPI
       ↓
backend          frontend against contract mocks
       \          /
        integration
```

Prefer when:

* API contract is stable enough;
* frontend can meaningfully progress against MSW;
* domain size or schedule justifies the coordination overhead.

Parallelism is an optimization, not a goal.

---

# 12. Frontend Mock-First State

Frontend work may be implemented and verified against contract-driven mocks before the live backend is available.

Call this state:

```text
FRONTEND_MOCK_VERIFIED
```

It means:

* production UI exists;
* mocked API behavior follows the intended OpenAPI contract;
* frontend unit/component verification has passed;
* rendered UI has been exercised as required;
* the feature has completed its explicitly scoped frontend verification.

It does **not** mean:

```text
real backend integration verified
```

Keep that distinction explicit.

---

# 13. Mock-First Does Not Guarantee Zero Integration Changes

The architecture should aim for:

```text
mock → real backend
```

without rewriting application behavior.

But this is a design goal, not an assumption that may override evidence.

If real integration requires application-code changes:

```text
identify why
```

Possible causes include:

* mock drift;
* OpenAPI drift;
* misunderstood state semantics;
* integration/environment behavior;
* actual implementation defect.

Route required code changes back through the appropriate Harscode Build/Patch activity.

Do not make unreviewed "integration-only" production edits outside the lifecycle because the change was expected to be wiring-only.

---

# 14. Cross-Domain Mock Batching

Multiple domains may temporarily reach:

```text
FRONTEND_MOCK_VERIFIED
```

before live integration occurs.

This is allowed when it materially improves development sequencing.

The state must be explicitly tracked.

Never rely on human memory to remember:

* which domains are mock-verified;
* which endpoints are still mocked;
* which real integrations remain;
* which integration blockers exist.

Mock-verified is a visible intermediate state, not a euphemism for complete.

---

# 15. Project Development States

Track meaningful project integration states explicitly.

Suggested vocabulary:

```text
CONTRACT_READY
BACKEND_VERIFIED
FRONTEND_MOCK_VERIFIED
INTEGRATED_VERIFIED
DOMAIN_FINALIZED
DELIVERED
```

These are **project orchestration states**.

They do not replace Harscode phase names.

For example:

```text
Harscode Testing
```

describes a task lifecycle phase.

```text
FRONTEND_MOCK_VERIFIED
```

describes what project-level integration evidence currently exists.

Do not use the terms interchangeably.

---

# 16. Integration Verification

When backend and frontend are available together, verify the real integration.

Use the real stack/interface appropriate to the feature.

For frontend-visible behavior, this includes real rendered browser exercise.

The project has chosen **Playwright as the standard browser-automation capability** for this purpose once repository wiring is present.

Playwright capability does not imply a mandatory comprehensive E2E suite.

Use browser automation for:

* real interaction;
* responsive/layout checks;
* integration verification;
* high-value repeatable regressions.

Promote scenarios into permanent E2E tests only when maintenance cost is justified.

If Playwright wiring is not yet present, do not claim automated browser verification was run.

---

# 17. Domain Finalization

Domain Finalization occurs after the required backend/frontend work for that domain has reached real integration.

It is a **cross-feature integration sweep**, not a second implementation workflow.

Finalization should confirm:

* expected backend/frontend connections use real services rather than temporary mocks;
* important cross-feature flows compose correctly;
* representative real-data loading/error/empty/success states behave correctly;
* financial/status meaning remains clear;
* applicable security/authorization boundaries survive integration;
* representative responsive/rendered behavior is correct;
* canonical UX/design/component systems are being followed;
* no tracked provisional dependency remains unintentionally unresolved.

Do not blindly rerun every feature test from scratch.

Use existing evidence plus integration-focused verification.

Any defect requiring code changes routes back through Build/Patch.

---

# 18. Delivery

Delivery occurs when the domain meets its intended integrated completion criteria.

Delivery should:

* mark domain state accurately in the development tracker;
* record unresolved accepted risks;
* record intentionally deferred work;
* ensure any temporary mock/provisional status remains visible if still accepted;
* complete the appropriate release/merge activity for the repository strategy in use.

Do not define a second universal "merge" step here if individual feature PRs have already merged through Harscode.

Project delivery state and Git merge state are separate concerns.

---

# 19. Human Authority

Human effort should be spent on decisions where judgment or project authority is genuinely required.

Mandatory examples:

## Tier 0

Human authors/pairs on critical core implementation.

## Tier 1

Human review before merge.

## Product/domain changes

Human approves changes to established business/spec semantics.

## Brand-defining decisions

Human approves canonical:

* logo;
* wordmark;
* major brand visual direction;
* other Level 4 assets.

## Major OPEN design decisions

Human/product approval where the design system assigns product authority.

## Schema/index application

Agents may analyze and propose indexes.

Applying an index remains a human-triggered action.

For Tier 0 tables, inspect locking/query-plan consequences with particular care.

Do not convert a human-authority requirement into a prompt asking an agent to approve its own decision.

---

# 20. Parallelization

Parallel work is allowed only when independence is real.

Consider:

* overlapping files;
* shared components;
* database tables;
* migrations;
* API schemas;
* domain dependencies;
* generated artifacts.

Two tasks that appear feature-independent may still conflict through a shared primitive or schema.

If multiple DB migrations are generated concurrently, avoid migration-number collisions through explicit coordination.

Frontend tasks modifying the same broad `ui/` or `shared/` contract should generally not proceed independently without an intentional coordination plan.

---

# 21. Directory Boundary

Backend and frontend implementation remain separate write scopes by default.

A frontend-scoped Build does not modify backend production code.

A backend-scoped Build does not modify frontend production code.

Cross-stack contract work should be coordinated explicitly rather than allowing one feature session to drift across both codebases.

Shared docs/API changes follow their own authority and approval rules.

This boundary does not prevent:

* reading cross-stack context;
* domain-level integration Testing;
* project Finalization.

It prevents uncontrolled cross-scope implementation.

---

# 22. Evidence and Risk Reporting

Kencleng keeps the principle:

> **Verifiability over trust.**

Implementation reports and PR risk notes must distinguish:

```text
proven
assumed
not tested
deferred
```

Do not treat an agent's statement as proof.

Claims should point to concrete evidence where evidence is mechanically available.

A risk discovered and reported is useful output.

A material known risk silently omitted is a process failure.

Exact artifact/report structure follows the active Harscode phase guidance plus root `AGENTS.md`.

---

# 23. Verification Ownership

Do not duplicate a fixed universal test sequence in this document.

Harscode owns lifecycle-phase testing responsibilities.

The target repository owns the actual executable commands/configuration.

Kencleng adds project-specific requirements by risk.

Examples:

* concurrency/race evidence for relevant concurrency-sensitive code;
* security verification for threat-model concerns;
* invariant/property evidence for applicable Tier 1 backend behavior;
* hostile-content rendering tests where user-controlled content exists;
* rendered browser verification for material frontend UI changes;
* real-stack integration before integrated/domain-finalized status.

Run the right verification for the risk.

Do not run expensive categories indiscriminately merely to appear thorough.

---

# 24. Project Status Tracking

Workflow policy and project status are different information.

This document should not accumulate dated progress amendments such as:

```text
account complete
campaign currently mock-verified
endpoint X still needs integration
```

That information belongs in a dedicated living development tracker.

Recommended owner:

```text
docs/project/kencleng-development-tracker.md
```

The tracker should record, at minimum:

```text
domain / feature
backend status
frontend status
integration status
important blockers
accepted/deferred risks
```

The existing Integration Tracker embedded in a scaffold/playbook should be migrated out rather than becoming the permanent source of operational status.

A one-off scaffold document is not the right long-term owner for living project state.

---

# 25. Generated Agent Artifacts

Exploration logs, techplans, build reports, review reports, testing reports, and other task artifacts belong in the target Kencleng repository according to the established task-work directory convention.

They are generated work products.

They do not become project-wide policy merely because a previous task contains a decision.

Promote recurring project truth into the appropriate canonical project document.

---

# 26. Evolution

Keep this document stable and project-specific.

Update it when Kencleng's orchestration model changes, for example:

* domain sequencing changes;
* risk authority changes;
* integration strategy changes;
* project-wide frontend/backend coordination changes;
* a recurring orchestration failure exposes a missing rule.

Do not add:

* generic React practices;
* generic Go practices;
* Harscode phase mechanics;
* one-feature implementation lessons;
* dated project progress.

Those have different owners.

Avoid amendment chains such as:

```text
Amendment 1
Previous amendment
Resolved 2026-...
```

When a policy changes, rewrite the canonical rule clearly.

Git history preserves the old version.

Use an ADR/proposal only when historical rationale itself needs to remain prominent.

---

# 27. Related Documents

Project authority:

* `AGENTS.md`
* `docs/spec/README.md`
* `docs/project/kencleng-backend-tech-stack.md`
* `docs/project/kencleng-frontend-tech-stack.md`
* `docs/project/kencleng-development-tracker.md` — once created

Frontend product/design:

* `docs/ui-ux/product-design-principles.md`
* `docs/ui-ux/patterns.md`
* `docs/ui-ux/design-guidelines.md`
* `docs/ui-ux/brand-and-visual-assets.md`
* `docs/ui-ux/page-map.md`
* `docs/ui-ux/prototype-reference.md`
* `docs/ui-ux/design-reference-usage.md`
* `frontend/components/README.md`

Portable workflow:

* Harscode `workflow/`
* Harscode `best-practices/`
