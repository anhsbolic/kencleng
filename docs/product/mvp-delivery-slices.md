# Kencleng — MVP Delivery Slices

> Status: **Approved MVP Delivery Sequencing — 2026-09-17**
> Upstream release scope: `docs/product/mvp-scope.md`
> Whole-product authority: `docs/product/product-overview.md`
> Product Design / Brand authority: canonical `docs/ui-ux/`
>
> This document sequences the approved MVP scope into vertical delivery slices. It does not define endpoint names, database schemas, frontend component trees, or final technical architecture. Those are derived slice-by-slice through Exploration, FE/BE delivery planning, and contract reconciliation.

## 1. Sequencing principle

The MVP should be delivered in **vertical product slices**, not historical backend-domain order.

The sequencing target is the approved trust loop:

```text
understand campaign
→ donate as guest
→ understand donation state
→ fundraising closes
→ return to the same campaign
→ understand result/accountability
```

Each slice should add a user-observable part of this loop and only the operational/security support required to make that part real.

Avoid this pattern:

```text
finish Account
→ finish Organization
→ finish Campaign
→ finish Donation
→ finish Reporting
→ finally integrate the product
```

Prefer:

```text
one narrow product capability
→ FE need + BE need
→ shared contract
→ implementation
→ integration evidence
→ next capability
```

## 2. Cross-slice guardrails

### 2.1 Product behavior must be real

A user-visible state must be backed by real application state/behavior, not hard-coded frontend storytelling.

### 2.2 Operational shortcuts are allowed only behind the experience

Campaign, Organization, lifecycle, or accountability data may initially be created through seeded or operator-assisted mechanisms where the approved MVP scope permits it.

The shortcut must not be represented to users as a self-service, automated, reviewed, or independently verified capability when it is not.

### 2.3 Security floor follows exposed capability

Do not pre-build the entire security roadmap. Do implement the security/correctness properties required by the capability currently exposed.

### 2.4 No fake next action

A slice must not expose an active control whose downstream capability does not yet exist.

In particular, Public Campaign Detail may be implemented before Guest Donation, but an active Donate action is not complete until it reaches the real MVP donation flow.

### 2.5 Existing specs/code are evidence, not automatic scope

For every slice, historical domain specs, OpenAPI, migrations, tests, and code are classified `KEEP`, `ADAPT`, `REPLACE`, or `DEFER` only after the slice's product need is understood.

## 3. Delivery order

```text
Slice 1 — Public Campaign Understanding
        ↓
Slice 2 — Guest Donation + Truthful Donation State
        ↓
Slice 3 — Campaign Closure + Persistent Public Result
        ↓
Slice 4 — Accountability Follow-up
        ↓
MVP end-to-end validation
```

This is a dependency order, not a requirement that frontend and backend work serially. Once a slice contract is reconciled, FE and BE work may proceed in parallel where appropriate.

## 4. Slice 1 — Public Campaign Understanding

### User outcome

A skeptical-but-open visitor can inspect a real public campaign and understand enough relevant truth to decide whether further action is worth considering.

### Product scope

The surface should communicate at minimum:

- campaign identity and purpose;
- responsible Organization/steward context at a deliberately public-safe level;
- real campaign media or a truthful missing-media state;
- current public lifecycle meaning;
- target/current funding facts available at this stage;
- organizer-provided story/context without presenting it as platform fact;
- relevant provenance and unknown/pending information;
- a truthful next-action area.

### Contract direction already established

Probe 01 provides the starting reconciliation direction:

- use an explicit public-safe Campaign projection;
- do not inherit the full internal Campaign record;
- use a narrow public Organization projection;
- align public media access with parent-campaign public visibility;
- preserve revocability for media that must stop being public;
- keep donation eligibility backend-authoritative.

### Minimum enabling operations

The slice needs enough persisted Organization/Campaign/media state to render a genuine public campaign.

A seeded or operator-assisted setup is acceptable for MVP. Full Organization registration, representative management, Campaign creation UI, and curation UI are not required here unless exploration proves them enabling-critical.

### Frontend boundary before Slice 2

The page may establish layout, hierarchy, responsive behavior, trust/provenance presentation, and lifecycle states before donation exists.

However:

```text
no real Donation Flow
→ no active Donate action pretending the flow exists
```

### Explicitly out of Slice 1

- account creation/login;
- registered donation history;
- public donor list;
- ranking/popularity/trending;
- broad Campaign Discovery;
- full Organization Profile;
- Organization/Curator/Admin self-service workflow;
- full post-campaign accountability UI.

### Slice 1 completion evidence

A public visitor can load a persisted eligible campaign and correctly understand campaign/steward/funding/source context across required loading, missing-media, success, not-public/not-found, failure, and responsive states without internal data leakage.

## 5. Slice 2 — Guest Donation + Truthful Donation State

### User outcome

A public visitor can contribute without creating an account and can understand the real sandbox processing/result state of that contribution.

### Product scope

The slice must provide:

- donation eligibility for an active public Campaign;
- one coherent guest donation flow;
- amount semantics and validation;
- one deliberately supported sandbox payment/processing path if one is sufficient;
- truthful pending/success/failure behavior;
- a safe guest mechanism to revisit/check donation status;
- clear copy that the sandbox mechanism is not real external payment settlement.

### Correctness/security floor

At minimum:

- duplicate client submissions cannot create unintended duplicate contributions where idempotency is required;
- settlement/result transitions cannot be forged through an exposed public/internal HTTP transition;
- successful settlement updates funding exactly once;
- concurrent successful donations cannot lose/corrupt funding increments;
- guest tracking capability does not expose another donor's sensitive information;
- tokens/identifiers used for status access are handled as sensitive credentials;
- logs do not leak guest email or tracking secrets;
- Campaign eligibility is enforced by backend state, not frontend visibility alone.

### Existing implementation posture

Historical Donation specs/OpenAPI are reference evidence only.

Likely salvage candidates include idempotency, guarded settlement transition, atomic funding increment, and guest status-token concepts, but the exact payment-method breadth, delays, percentages, minimum amount, endpoint shape, and guest-data requirements must be revalidated for this slice.

### Account consequence

**Account is not required for the baseline Slice 2 path.**

The primary journey is Guest Donation. Authentication must not be pulled into the MVP merely because existing code supports registered donation.

### Explicitly out of Slice 2

- account registration/login as a prerequisite;
- guest-donation claim into an account;
- registered donation history;
- multiple payment methods merely for completeness;
- real banking/payment rails;
- public donor/social-proof list;
- campaign closure/result behavior beyond what is needed to keep donation eligibility correct.

### Slice 2 completion evidence

A visitor can move from an eligible Public Campaign Detail into a real guest donation, observe pending/success/failure accurately, revisit the donation through the safe guest mechanism, and see Campaign funding reflect successful settlement exactly once.

## 6. Slice 3 — Campaign Closure + Persistent Public Result

### User outcome

When fundraising ends, the Campaign does not disappear. A visitor/donor can return to the same public Campaign identity and understand that fundraising has ended and what the final funding result is.

### Product scope

The slice must provide:

- at least one truthful supported Campaign closure path;
- no donation acceptance after closure;
- the same public Campaign identity/URL after closure for Campaigns that were genuinely public;
- final funding/result facts;
- closed-state lifecycle meaning;
- changed information hierarchy: result/accountability becomes primary, donation action disappears;
- an honest pending state when later accountability does not yet exist.

### Safety boundary

`closed` does not mean every terminal/internal Campaign record becomes public.

Public continuity applies only to Campaigns that entered the public lifecycle according to the approved Probe 01 decision.

### Operational mechanism

The exact first closure trigger may be narrow. Exploration can choose the smallest truthful mechanism required by the MVP scenario.

A manual/operator-assisted trigger is acceptable if:

- it is authorized and cannot be invoked by the public;
- the resulting lifecycle state is persisted;
- user-facing copy does not falsely claim the closure was automatically triggered by a condition that did not occur.

### Explicitly out of Slice 3

- every future publish/schedule/unpublish/force-close workflow;
- full Owner publication controls;
- full Admin exceptional-action UI;
- complete historical Campaign lifecycle administration.

### Slice 3 completion evidence

A Campaign that accepted a successful MVP donation can transition to its supported closed state; donation becomes unavailable; the original public route remains valid; final funding/result state is displayed truthfully; non-public Campaign records remain undisclosed.

## 7. Slice 4 — Accountability Follow-up

### User outcome

A visitor/donor returning after closure can understand what happened next—or clearly understand that the next information is still pending—without Kencleng turning organizer claims into platform-verified impact.

### Product scope

The first accountability capability should be intentionally narrow.

It needs enough structure to demonstrate:

```text
platform-known lifecycle/funding fact
≠ organizer-reported follow-up
≠ reviewed/verified report
≠ independently verified real-world impact
```

A minimal accountability entry may include:

- factual type/source semantics;
- organizer-reported content where applicable;
- timestamp/chronology;
- source/provenance label;
- attachments/media only when truthful and safely governed;
- pending/not-yet-reported state.

### Operational mechanism

For the first MVP, accountability content may be seeded or operator-assisted in persisted application state.

The system must not claim that content is independently verified unless a real review capability exists and actually produced that state.

### Evidence Journal boundary

The MVP may express the beginning of the Evidence Journal concept without implementing its entire future feature set.

Do not add:

- universal impact scoring;
- invented milestones;
- automatic causal claims;
- broad disbursement/fund-usage machinery merely to make the page look complete.

### Explicitly out of Slice 4

- full Organization fund-usage submission workflow;
- full Curator report-verification workflow;
- full disbursement workflow;
- comprehensive notification center;
- advanced Evidence Journal filtering/taxonomy;
- impact scoring/guarantees.

### Slice 4 completion evidence

A closed Campaign can show either a real persisted organizer-reported follow-up with clear provenance or an explicit pending state; the experience keeps funding, reporting, verification, and impact meanings distinct.

## 8. MVP Account decision

For the approved baseline trust loop, **Account is not on the critical path**.

Therefore:

```text
Account implementation
→ do not resume merely because it already exists
→ preserve existing security/correctness work as salvage evidence
→ defer product reframing until an Account-dependent capability enters scope
```

Probe 02 remains valuable backward-reconciliation evidence but no longer blocks MVP delivery sequencing.

The first Account-dependent candidate after/beside MVP may be one of:

- persistent registered donation history;
- guest-donation claim;
- Organization representative self-service;
- a privileged operator surface if operational tooling can no longer remain narrower/internal.

When one of those becomes real scope, Account Product Reframing should derive the minimum identity/authentication model from that need rather than restoring historical feature breadth.

## 9. Dependency summary

```text
Persisted public campaign truth
        ↓
Slice 1 Public Campaign Understanding
        ↓
real donation eligibility
        ↓
Slice 2 Guest Donation + Status
        ↓
real funding state
        ↓
Slice 3 Closure + Persistent Result
        ↓
closed campaign identity/result
        ↓
Slice 4 Accountability Follow-up
        ↓
complete MVP trust-loop evidence
```

Cross-cutting design, accessibility, responsive behavior, public-data safety, and security/correctness expectations apply inside each slice rather than being postponed into a final cleanup phase.

## 10. Contract discipline

Do not reconcile the whole OpenAPI before delivery.

For each slice:

```text
Product Authority + MVP Scope + Design Authority
        ↓
Harscode Exploration
        ↓
FE delivery need + BE capability need
        ↓
inspect existing specs/OpenAPI/code as evidence
        ↓
reconcile only the contract needed by this slice
        ↓
implement + integrate
        ↓
feed real evidence upstream when warranted
```

Old contracts outside the active slice remain reference material until their turn.

## 11. CRTV relationship

These slices become real Continuous Real-Task Validation inputs for Harscode after Product Authority/routing promotion is complete.

When CRTV begins for a slice:

- use Harscode `main` Operational Default / Level 2;
- use the canonical Harscode Exploration kickoff;
- provide only normal kickoff variables/context;
- do **not** add a custom solution-steering prompt that tells the agent what it is supposed to discover;
- judge whether authority, scope, ambiguity, and relevant existing evidence are discoverable from the repository itself.

This document defines product delivery sequencing; it must not become a replacement for the Harscode workflow.

## 12. Approval checkpoint

Human approval: **2026-09-17**.

Approved order:

> **Slice 1 Public Campaign Understanding → Slice 2 Guest Donation + Truthful Donation State → Slice 3 Campaign Closure + Persistent Public Result → Slice 4 Accountability Follow-up. Account remains outside the baseline MVP critical path and resumes only when a real scoped capability requires it.**

With this approval, conceptual MVP sequencing is closed. Remaining reframe work is repository authority promotion/routing, followed by Slice 1 as the first real post-promotion CRTV task.