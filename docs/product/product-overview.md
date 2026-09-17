# Kencleng — Product Overview

> Status: **Candidate Product Authority — not yet canonical**
> Created: 2026-09-17
> Scope: Whole-product business/product model at deliberately lower resolution than delivery specifications.
> Source basis: existing business-process, actor/entity, phase, design/page-map, domain-spec, API, and implementation evidence. Detailed old decisions are not automatically promoted here.

## 1. Product purpose

Kencleng is a sandbox donation/crowdfunding product centered on Indonesian non-profit organizations, public fundraising campaigns, donations, and post-campaign accountability.

The product exists both as a usable product model and as a realistic learning environment for full-stack, correctness-conscious development. The sandbox nature allows payment/disbursement infrastructure to be simulated where appropriate, but simulation must not weaken the product's truthfulness around money, identity, lifecycle state, or evidence.

At product level, Kencleng aims to let people:

- understand who is organizing a campaign and what the campaign is for;
- donate with clear consequences and without requiring an account when the product permits guest participation;
- distinguish funding progress from later operational/accountability claims;
- follow what happened after fundraising rather than treating payment as the end of the relationship;
- let organizations raise funds through an explicit lifecycle with curation/accountability gates;
- let platform roles perform review/administration without collapsing distinct verification concepts into a generic trust badge.

## 2. Product boundary

Kencleng v1 is intentionally a sandbox rather than a production financial institution or payment processor.

The product model includes:

- user identity/account access;
- organization participation and representation;
- campaign creation, curation, publication, fundraising, closure, and post-campaign reporting;
- guest and registered donation experiences;
- organization fund-disbursement/accountability flow;
- curator/admin operational roles;
- product notifications where relevant to lifecycle/accountability.

The product model does not need to emulate every real-world fundraising platform concern. Real payment rails, real bank settlement, government registry integrations, and other external systems are included only when a delivery slice explicitly requires them; sandbox substitutes may be used when the learning/product goal does not depend on the real integration.

## 3. Primary actors

### Public Visitor / Guest Donor

A person who can understand Kencleng and public campaigns without signing in, and may donate as a guest when the active campaign/product rules allow it.

Primary goals:

- discover and understand campaigns;
- inspect organizer and trust/accountability context;
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

The exact permission matrix is delivery/product-detail that must be referenced or reconciled per slice rather than duplicated here.

### Curator

A reviewer responsible for product-defined curation/verification work. Curator responsibilities may include organization review, campaign review, and fund-usage accountability review.

These are distinct review concepts even when they reuse similar interaction patterns.

A curator must not review work where a relevant conflict of interest exists with an organization they represent.

### Admin

A platform-level operational role for product-defined administrative/review-assignment/exceptional actions.

Admin is not a synonym for Curator or Organization Owner and should not casually collapse those responsibilities.

### System

The platform itself performs lifecycle-triggered behavior such as state transitions, progress/result generation, notifications, scheduled actions, and simulated settlement where defined by the relevant delivery slice.

## 4. Core product concepts

These are product concepts, not commitments to a particular table/class/module design.

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

### Disbursement

The controlled movement/release of collected campaign funds to the responsible Organization after the required product conditions are met.

In the sandbox, real banking mechanics may be simulated; the product meaning and authorization consequences must still be clear.

### Fund-Usage Report

Structured accountability from the Organization after disbursement, subject to the product's review/verification process.

### Notification

Lifecycle/product communication to users when the product requires active awareness beyond pull-based page viewing.

## 5. Whole-product journey

The product can be understood as one connected loop rather than independent backend domains.

```text
ACCOUNT / IDENTITY
A person may participate publicly, create an account, or authenticate
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
Public/donors can understand funding state without confusing it with real-world outcome
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

Not every delivery slice must implement this entire loop. The point of the model is to preserve end-to-end meaning while slices are delivered progressively.

## 6. Major capability map

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
- inspect a campaign's purpose, organizer context, funding context, story, and relevant accountability context;
- reach a truthful next action without unsupported urgency/ranking/trust claims.

### Donation

- submit a donation as guest or registered donor where permitted;
- understand processing/result state;
- preserve public anonymity/display choices independently from account status;
- support personal history/follow-up for registered donors and supported guest-tracking/claim behavior.

### Post-campaign result and accountability

- preserve campaign/funding result after closure;
- let relevant users receive or inspect post-campaign follow-up;
- let Organizations provide product-supported reporting/context;
- distinguish organizer narrative/reporting from platform-verified facts.

### Disbursement and fund usage

- request and review release of campaign funds;
- prevent unauthorized/double release according to delivery invariants;
- report how disbursed funds were used;
- review/verify fund-usage accountability according to product rules.

### Platform review/administration

- route appropriate work to eligible Curators;
- maintain conflict-of-interest boundaries;
- support product-defined Admin-only operational actions without inventing a generic all-powerful dashboard model.

### Notifications

- notify users when active communication is part of the product journey rather than requiring all awareness to come from repeatedly checking pages.

## 7. Trust and accountability model

Kencleng's product meaning depends on keeping several truth classes distinct.

### 7.1 Verification concepts are not interchangeable

At minimum, distinguish:

- Organization identity/legitimacy review;
- Campaign curation/publication state;
- Donation/payment state;
- Fund-usage report verification;
- Organizer-provided narrative/reporting;
- real-world outcome/impact claims.

A campaign must not be described broadly as "verified" when the actual known fact is narrower.

### 7.2 Funding is not impact

Keep these product concepts separate:

```text
funding progress
≠ operational progress
≠ organizer-reported outcome
≠ independently verified real-world impact
```

A fully funded campaign proves only the applicable funding fact.

### 7.3 Unknown/pending information remains legitimate

The product may not yet know whether an activity happened, whether an outcome was achieved, or whether a report has been reviewed. Unknown/pending is a real product state and should not be filled with optimistic inference.

### 7.4 Accountability continues after donation

Donation is not the conceptual endpoint of the product. The post-campaign result/disbursement/fund-usage loop is part of the trust model and should remain connected to donor/public understanding when the relevant information exists.

## 8. Strong current product truths

The following appear consistently across the existing product/design/spec evidence and are strong candidates for promotion when this document is reviewed:

1. **Organizations and Campaigns are distinct concepts.** Campaign fundraising/accountability belongs to an Organization.
2. **Guest donation is a supported product capability.** An account is not universally required to contribute.
3. **Registered and anonymous are different axes.** Public identity display must not be inferred solely from account/guest status.
4. **Organization review, Campaign curation, and fund-usage verification are distinct product concepts.**
5. **Owner and Staff represent meaningfully different authority levels.** Exact permissions are reconciled per capability.
6. **Curator conflicts of interest are product-significant.** A curator must not review the relevant work of an Organization they represent.
7. **Campaign lifecycle governs public/donation behavior.** A campaign is not always public or eligible to accept donations merely because it exists.
8. **Money semantics and accountability states must be explicit.** Funding progress must not become a proxy for execution or impact.
9. **Post-campaign accountability is in the core Kencleng concept, not optional decorative content.**

## 9. Deliberately not promoted from old detailed specs yet

The previous generation made many detailed decisions. They remain references and may be correct, but should be revalidated when the corresponding slice approaches delivery rather than copied into whole-product authority now.

Examples include:

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

These details should be promoted into Product Authority only when they are genuinely durable product decisions rather than convenient current implementation choices.

## 10. Product questions to revalidate through real slices

The reframe should not reopen every historical decision at once. Revalidation is triggered by delivery.

Initial questions include:

### Forward probe — Public Campaign Detail

- What information must a visitor understand before donation becomes a meaningful next action?
- Which trust/accountability facts are available at this lifecycle stage?
- What organizer/campaign information is product truth versus organizer-provided content?
- Which image/media state is truthful when real campaign evidence is absent?
- What is the honest CTA boundary while Donation Flow is a separate delivery slice?
- What backend/data capabilities are actually required by this surface?

### Backward probe — Registration + Email Verification

- What user outcome actually requires verified email in Kencleng?
- Which downstream capabilities depend on account verification?
- Is the current identity model/product behavior still justified independently from the backend implementation?
- Which current Account security behaviors are product requirements versus implementation/security architecture choices?
- Does the existing contract match the simplest correct end-to-end account experience?

## 11. Product development posture

Kencleng should use **progressive commitment**, not all-product detailed specification before coding.

```text
whole-product clarity at product level
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

## 12. Promotion criteria

This candidate should not become canonical merely because it exists.

Before promotion:

1. Human review confirms that the whole-product model is directionally correct.
2. The Public Campaign Detail forward probe demonstrates that product → design → FE/BE delivery → contract derivation is workable.
3. The Account Registration + Email Verification backward probe demonstrates that existing detailed specs/implementation can be reconciled without either blindly preserving or unnecessarily deleting prior work.
4. Any product-level contradictions surfaced by those probes are resolved here rather than hidden downstream.
5. Repository routing is updated in one deliberate promotion change so there is never an ambiguous long-lived dual authority.
