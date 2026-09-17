# Kencleng — Product Overview

> Status: **Canonical Product Authority — promoted 2026-09-17**
> Created: 2026-09-17
> Updated: 2026-09-17 — Product Authority promotion
> Scope: Whole-product business/product model at deliberately lower resolution than delivery specifications.
> Source basis: existing business-process, actor/entity, phase, design/page-map, domain-spec, API, implementation, and September 2026 product-brand exploration evidence. Detailed old decisions are not automatically promoted here.

## 1. Product thesis

Kencleng is an evidence-led donation/crowdfunding product centered on Indonesian non-profit organizations, public fundraising campaigns, donations, and post-campaign accountability.

Its core product thesis is:

> **Evidence-Led Optimism**

Supporting phrase:

> **Hope, structured by evidence.**

The desired product sequence is:

```text
evidence
→ confidence
→ optimism
→ action
→ continued understanding after action
```

Kencleng should help people act because they understand enough to make a considered decision — not because the interface manufactures urgency, guilt, symbolic trust, or emotional pressure.

The product should remain useful as a realistic full-stack learning environment, but the sandbox nature must not weaken truthfulness around money, identity, lifecycle state, evidence, or accountability.

## 2. Product character

The following characteristics were discovered and validated during the September 2026 Product Brand + UI exploration. They are not merely visual adjectives; they constrain product behavior and communication.

### Thoughtful

Kencleng gives users enough context to understand consequential decisions. It does not rush emotionally or financially sensitive actions.

### Candid

Unknown, pending, delayed, changed, rejected, incomplete, or not-yet-reported information may remain visible as such. The product should not replace uncertainty with optimistic inference.

### Composed

Difficult content, money, verification, privacy, and negative states are communicated calmly and specifically rather than dramatized.

### Human

Real people and real context matter. Campaign participants must retain dignity, privacy, agency, and context rather than becoming conversion assets.

### Optimistic

Kencleng can show possibility, participation, progress, and constructive next steps without promising outcomes it cannot prove.

### Refined but accessible

The product should feel intentional and trustworthy without becoming exclusive, institutional-financial, intimidating, or sterile.

## 3. Durable product principles discovered through design

These principles sit at the product layer. Their exact UI expression remains owned by `docs/ui-ux/`.

### 3.1 Confidence before conversion

Donation is a consequential user decision, not a conversion event to maximize.

Before donation becomes the dominant next action, users should be able to understand enough relevant context about the campaign, steward, funding state, and available accountability information to make an informed decision.

### 3.2 Facts first, story with them

Preferred model:

> **facts that have a story**

Not:

> story decorated with selective facts.

Campaign storytelling may provide human context, but it must not obscure or selectively reshape the facts that matter to the user's decision.

### 3.3 Trust comes from structure, not trust theatre

Trust should emerge through ordered truth, provenance, chronology, clear status meaning, consistent behavior, and respectful transparency.

Badges, colors, icons, and certification-like language must not substitute for the actual facts the platform possesses.

Working trust equation from the design exploration:

> **Trust = ordered truth + calm candor + respectful transparency.**

### 3.4 Dignity over pity

Hardship may be shown because it is real context. Suffering must not become the primary persuasion mechanism.

The product should avoid guilt, exaggerated vulnerability, pity framing, and other coercive fundraising patterns.

### 3.5 Optimism remains evidence-aware

Optimism should come from possibility, participation, visible progress, reported milestones, continuity, and constructive next steps.

It must not imply causality, success, verification, or real-world impact beyond available evidence.

### 3.6 Unknown is a valid product state

Information may legitimately be unavailable, pending, under review, delayed, changed, disputed, or not yet reported.

Kencleng should preserve that uncertainty rather than hide it or manufacture a complete-looking story.

### 3.7 Progress is evidence, not gamification

Keep these concepts distinct:

```text
funding progress
≠ operational progress
≠ organizer-reported outcome
≠ independently verified real-world impact
```

A full funding target proves only the applicable funding fact.

### 3.8 Accountability continues after donation

Donation is not the conceptual endpoint of Kencleng.

Post-donation and post-campaign follow-up are part of the core trust proposition. Where relevant information exists, donors and the public should be able to understand what happened next, where information came from, and what remains pending.

The **Evidence Journal** is the strongest current experience concept expressing this principle: chronological factual updates, provenance, reports, milestones, and pending next states without collapsing them into one universal “impact” score.

### 3.9 Kencleng facilitates; it does not guarantee outcomes

Curation, organization review, publication state, report verification, and other platform mechanisms have specific meanings.

Kencleng must not imply that its review processes guarantee beneficiary identity, program execution, or real-world outcome unless the product actually possesses that evidence.

## 4. Audience posture

The primary public behavioral model is a **skeptical-but-open donor**.

The product respects that a user may reasonably want to understand:

- what the campaign is for;
- who manages it;
- what Kencleng itself knows;
- what the organizer says or reports;
- what is still pending or unknown;
- what happened after donation;
- what action is meaningful now.

Kencleng should earn trust rather than demand it.

## 5. Product boundary

Kencleng v1 is intentionally a sandbox rather than a production financial institution or payment processor.

The product model includes:

- user identity/account access;
- organization participation and representation;
- campaign creation, curation, publication, fundraising, closure, and post-campaign reporting;
- guest and registered donation experiences;
- organization fund-disbursement/accountability flow;
- curator/admin operational roles;
- product notifications where relevant to lifecycle/accountability.

The product model does not need to emulate every real-world fundraising-platform concern. Real payment rails, real bank settlement, government registry integrations, and other external systems are included only when a delivery slice explicitly requires them; sandbox substitutes may be used when the learning/product goal does not depend on the real integration.

Simulation must never be presented as real evidence.

## 6. Primary actors

### Public Visitor / Guest Donor

A person who can understand Kencleng and public campaigns without signing in, and may donate as a guest when the active campaign/product rules allow it.

Primary goals:

- understand Kencleng's trust/transparency model;
- discover and understand campaigns;
- inspect organizer, funding, and accountability context;
- decide whether to donate;
- track a guest donation through an appropriate safe mechanism when supported.

### Registered User / Donor

A user with a Kencleng account who inherits public capabilities and can additionally access personal account/donation capabilities.

Primary goals:

- authenticate securely;
- maintain supported account/profile concerns;
- view personal donation history/status;
- follow campaign/accountability updates;
- associate eligible previous guest donations with the account when the product supports that relationship.

### Organization Representative

A registered user who represents an Organization. Representation has at least two meaningful permission levels:

- **Owner** — responsible for sensitive/authoritative organization actions;
- **Staff** — may assist with permitted operational work but does not automatically inherit every Owner capability.

The exact permission matrix is delivery/product detail that must be referenced or reconciled per slice rather than duplicated here.

### Curator

A reviewer responsible for product-defined curation/verification work. Curator responsibilities may include organization review, campaign review, and fund-usage accountability review.

These are distinct review concepts even when they reuse similar interaction patterns.

A curator must not review work where a relevant conflict of interest exists with an organization they represent.

### Admin

A platform-level operational role for product-defined administrative, review-assignment, and exceptional actions.

Admin is not a synonym for Curator or Organization Owner and should not casually collapse those responsibilities.

### System

The platform itself performs lifecycle-triggered behavior such as state transitions, progress/result generation, notifications, scheduled actions, and simulated settlement where defined by the relevant delivery slice.

## 7. Core product concepts

These are product concepts, not commitments to a particular table, class, package, or module design.

### User

A person with an account and product roles/relationships.

### Organization

The party that establishes organizational identity on Kencleng and owns/manages fundraising campaigns.

Organization identity/legitimacy review is separate from campaign curation and from later fund-usage verification.

### Organization Representation

The relationship connecting users to an Organization with product-defined authority levels such as Owner and Staff.

### Campaign

The core fundraising initiative. A Campaign belongs to an Organization and has a lifecycle that determines when it can be publicly understood, accept donations, close, and proceed into post-campaign accountability.

### Event

A lightweight promotional/context concept that may relate to campaigns but is not the core fundraising/accountability unit.

Its exact v1 importance should remain subordinate to real delivery needs; it must not distort the primary campaign journey merely because a historical schema supports it.

### Donation

A contribution to an eligible Campaign. Donations may originate from guests or registered users.

Donation identity/public-display semantics are distinct concerns: whether someone has an account is not the same question as whether their identity is displayed publicly.

### Curation / Review Assignment

A product mechanism for routing review work to eligible reviewers. Organization review, campaign review, and fund-usage review remain separate semantic responsibilities.

### Campaign Result / Accountability Context

The product information that helps users understand what happened after fundraising, including funding/result facts and later organizer/accountability reporting where available.

### Evidence Journal

A donor-facing product concept for following what happened after donation through chronology, factual progress, source/provenance, reports, milestones, and pending updates.

It is not a universal impact score.

### Disbursement

The controlled movement/release of collected campaign funds to the responsible Organization after the required product conditions are met.

In the sandbox, real banking mechanics may be simulated; the product meaning and authorization consequences must still be clear.

### Fund-Usage Report

Structured accountability from the Organization after disbursement, subject to the product's review/verification process.

### Notification

Lifecycle/product communication to users when the product requires active awareness beyond pull-based page viewing.

## 8. Whole-product journey

Kencleng should be understood as one connected loop rather than independent backend domains.

```text
PUBLIC UNDERSTANDING / ACCOUNT IDENTITY
A person may understand the platform publicly, create an account, or authenticate
        │
        ▼
ORGANIZATION ESTABLISHMENT
An eligible user establishes/represents an Organization
        │
        ▼
ORGANIZATION REVIEW
The Organization reaches the product state needed for fundraising activity
        │
        ▼
CAMPAIGN PREPARATION
Representatives create/revise a fundraising campaign
        │
        ▼
CAMPAIGN CURATION
The campaign is reviewed through the appropriate product process
        │
        ▼
PUBLICATION / PUBLIC UNDERSTANDING
The campaign becomes publicly discoverable/understandable when eligible
        │
        ▼
DONATION
Guest or registered donors may contribute while the campaign accepts donations
        │
        ▼
FUNDING PROGRESS
Public/donors can understand funding state without confusing it with outcome
        │
        ▼
CAMPAIGN CLOSURE
The fundraising window ends according to product lifecycle rules
        │
        ▼
RESULT / FOLLOW-UP
The product preserves campaign result/context and communicates relevant follow-up
        │
        ▼
DISBURSEMENT
Eligible Organization Owners request/receive release of funds through the controlled process
        │
        ▼
FUND-USAGE ACCOUNTABILITY
The Organization reports usage and the applicable review process verifies/report states it
        │
        ▼
PUBLIC / DONOR EVIDENCE
Users can distinguish collected funding, platform/system facts, organizer reports,
verification state, and still-unknown real-world outcome
```

Not every delivery slice must implement this entire loop. The model preserves end-to-end meaning while slices are delivered progressively.

## 9. Major capability map

### Identity and account

- register/authenticate through supported identity methods;
- establish verified account identity where required by downstream capability;
- manage supported account/security concerns;
- expose only the roles/capabilities the product actually grants.

### Organization participation

- establish an Organization;
- establish/manage permitted representative relationships;
- maintain organization information;
- undergo product-defined organization identity/legitimacy review.

### Campaign management

- create/revise campaign intent/content;
- submit campaign for curation;
- react to curation outcomes;
- publish/unpublish/close according to product lifecycle authority;
- monitor campaign state as an Organization representative.

### Public campaign experience

- discover eligible public campaigns;
- inspect a campaign's purpose, organizer/steward context, factual funding context, story, and relevant accountability context;
- distinguish platform fact, organizer-provided content, system state, and unknown/pending information;
- reach a truthful next action without unsupported urgency, ranking, popularity, recommendation, trust score, or impact claim.

### Donation

- submit a donation as guest or registered donor where permitted;
- understand processing/result state;
- preserve public anonymity/display choices independently from account status;
- support personal history/follow-up for registered donors and supported guest-tracking/claim behavior.

### Post-campaign result and accountability

- preserve campaign/funding result after closure;
- let relevant users receive or inspect post-campaign follow-up;
- let Organizations provide product-supported reporting/context;
- distinguish organizer narrative/reporting from platform-known/verified facts;
- support the Evidence Journal concept where product data and source semantics are sufficient.

### Disbursement and fund usage

- request and review release of campaign funds;
- prevent unauthorized/double release according to delivery invariants;
- report how disbursed funds were used;
- review/verify fund-usage accountability according to product rules.

### Platform review/administration

- route appropriate work to eligible Curators;
- maintain conflict-of-interest boundaries;
- support product-defined Admin-only operational actions without inventing a generic all-powerful control-center model.

### Notifications

- notify users when active communication is part of the product journey rather than requiring all awareness to come from repeatedly checking pages.

## 10. Trust and accountability model

Kencleng's product meaning depends on keeping several truth classes distinct.

### 10.1 Verification concepts are not interchangeable

At minimum, distinguish:

- Organization identity/legitimacy review;
- Campaign curation/publication state;
- Donation/payment state;
- Fund-usage report verification;
- organizer-provided narrative/reporting;
- real-world outcome/impact claims.

A campaign must not be described broadly as “verified” when the actual known fact is narrower.

### 10.2 Product truth classes remain visible

Where relevant, users should be able to distinguish:

- platform/system fact;
- organizer-provided information;
- organizer report;
- system/lifecycle status;
- pending or unavailable information;
- reported outcome.

Exact terminology remains a design/product-detail decision and should not be invented globally before real surfaces need it.

### 10.3 Transparency is ordered, not maximal

Transparency does not mean displaying every available field at once.

The product should expose what matters to the current decision while keeping consequential supporting detail inspectable.

Trust-critical uncertainty and consequences must not be hidden merely to create a cleaner or more optimistic story.

### 10.4 Campaign imagery and evidence remain truthful

Real campaign imagery belongs to campaign reality. Illustration or generated expressive assets may explain product concepts, onboarding, process, or empty states, but must never masquerade as beneficiary, campaign, distribution, or real-world impact evidence.

When evidence is unavailable, a truthful missing/provisional state is preferable to fabricated documentary-looking completeness.

## 11. Strong current product truths

The following are canonical whole-product truths unless deliberately revised through Product Authority:

1. **Kencleng is evidence-led, not persuasion-led.** Trust should be earned through ordered truth, evidence, provenance, chronology, and candid uncertainty.
2. **Confidence comes before conversion.** Donation actions should not outrun the context users need for a considered decision.
3. **Dignity over pity.** Human hardship may be real context but must not be exploited as a conversion device.
4. **Funding is not impact.** Funding, operational progress, reported outcome, and independently verified impact remain distinct.
5. **Unknown/pending is valid product truth.** The product should not manufacture certainty.
6. **Post-donation accountability is core.** Donation is not the conceptual end of the user relationship.
7. **Organizations and Campaigns are distinct concepts.** Campaign fundraising/accountability belongs to an Organization.
8. **Guest donation is a supported product capability.** An account is not universally required to contribute.
9. **Registered and anonymous are different axes.** Public identity display must not be inferred solely from account/guest status.
10. **Organization review, Campaign curation, and fund-usage verification are distinct product concepts.**
11. **Owner and Staff represent meaningfully different authority levels.** Exact permissions are reconciled per capability.
12. **Curator conflicts of interest are product-significant.** A curator must not review the relevant work of an Organization they represent.
13. **Campaign lifecycle governs public/donation behavior.** A campaign is not always public or eligible to accept donations merely because it exists.
14. **Money semantics and accountability states must be explicit.** Consequential numbers and states must retain their real meaning.

## 12. Deliberately not promoted from old detailed specs yet

The previous generation made many detailed decisions. They remain references and may be correct, but should be revalidated when the corresponding slice approaches delivery rather than copied into whole-product authority now.

### Account / security delivery detail

- exact identity-provider set;
- password length/policy and breach-list provider behavior;
- token/session TTLs;
- exact MFA/recovery behavior;
- exact encrypted persistence/HMAC model;
- endpoint/request/response shapes.

Security architecture may independently retain stricter implementation requirements; not promoting a detail into Product Authority does not mean weakening security.

### Organization detail

- exact legal document set;
- exact NPWP format/uniqueness mechanism;
- exact number-of-organizations-per-user limit;
- direct-add versus invitation acceptance for representatives;
- exact re-verification field classification.

### Campaign detail

- exact campaign field set;
- exact category enum;
- scheduling/rescheduling mechanics;
- exact close/unpublish mechanics;
- event behavior beyond its current high-level concept.

### Donation detail

- exact minimum donation amount;
- exact simulated payment-method list;
- simulated success/failure probability and processing delay;
- exact guest follow-up token design;
- exact guest-donation claim matching mechanism;
- exact donor-list display label.

### Disbursement / reporting detail

- lump-sum versus another disbursement mechanism;
- exact report deadline;
- exact expense categories/attachment requirements;
- exact consequences of late/rejected reports;
- exact notification mechanism/cadence.

### Concrete design-system detail

The following remain in canonical `docs/ui-ux/` rather than Product Authority:

- Sunlit Editorial color values and visual-role rules;
- typography choices/scales;
- spacing, surfaces, border, radius, elevation, and component grammar;
- iconography implementation;
- illustration/photography production rules;
- responsive composition mechanics;
- exact page/surface composition;
- reusable interaction-pattern implementation detail.

These are active design authority, not historical reference. They are excluded here only to avoid duplicate ownership.

## 13. Product questions to revalidate through real slices

The reframe should not reopen every historical decision at once. Revalidation is triggered by delivery.

### Public Campaign / MVP slices

Probe 01 established the current public Campaign direction, including persistent public identity after eligible closure and explicit public-safe projection. Remaining detail should now be resolved through the approved MVP slices rather than by re-specifying the whole Campaign domain upfront.

Questions that may still arise include:

- what exact public facts are required by the active slice;
- what organizer/campaign information is platform fact versus organizer-provided content;
- what remains unknown or pending and how that truth remains visible;
- which image/media state is truthful and revocable;
- what backend/data capabilities are actually required by the current surface.

### Account when it becomes real scope

Probe 02 is intentionally paused/reframed. Account is outside the baseline MVP critical path.

When an Account-dependent capability enters scope, revalidate:

- what user outcome actually requires an account;
- what proof/verification the capability truly needs;
- which existing security/correctness mechanisms remain applicable;
- whether historical provider/linking/recovery breadth is still justified;
- what the smallest coherent Account experience is for that real need.

Do not resume the historical Account roadmap merely because implementation already exists.

## 14. Product development posture

Kencleng should use **progressive commitment**, not all-product detailed specification before coding.

```text
whole-product clarity at product level
        +
canonical product-design / brand authority
        +
approved release scope
        ↓
select next valuable delivery slice
        ↓
increase product/design resolution for that slice
        ↓
derive FE + BE delivery needs
        ↓
reconcile shared contract
        ↓
implement and integrate
        ↓
learn from real evidence
        ↓
refine authority when warranted
```

A future area may remain high-level until it becomes relevant. A high-risk, hard-to-reverse business rule may require earlier precision than a reversible API/UI detail.

Product, design, and implementation discovery are bidirectional. Durable learning discovered downstream should be promoted deliberately to the authority that owns it rather than remaining trapped in a feature implementation or visual artifact.

## 15. Promotion checkpoint

Product Authority was promoted on **2026-09-17** after:

1. human review accepted the product thesis, product character, whole-product model, and business direction;
2. canonical design authority and Product Authority were reconciled as peer upstream authorities;
3. Public Campaign Detail forward validation demonstrated that Product + Design authority can derive delivery/contract needs without task-specific steering;
4. Account backward reconciliation demonstrated that detailed specs/implementation can be treated as evidence while preserving useful security/correctness work;
5. the MVP scope and vertical delivery sequencing were explicitly approved, preventing historical domain breadth from defining current delivery scope;
6. repository routing was deliberately promoted so Product/MVP authority sits above delivery specs/contracts for the concerns it owns.

Future changes to these durable product truths require explicit Product Authority revision; downstream implementation detail must not silently redefine them.