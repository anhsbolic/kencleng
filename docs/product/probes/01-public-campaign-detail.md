# Probe 01 — Public Campaign Detail

> Status: **Probe evidence — not authority**
> Created: 2026-09-17
> Purpose: Test whether the candidate Product Authority + canonical Product Design / Brand Authority can derive a real delivery slice, frontend needs, backend capabilities, and contract gaps without treating old domain specs/OpenAPI as automatically correct.
>
> This document is intentionally a validation artifact. If the model proves useful, durable conclusions should be promoted into the authority that owns them rather than leaving this probe as a second source of truth.

## 1. Probe question

Can Kencleng derive **Public Campaign Detail** from upstream product/design truth first, then use the existing Campaign/Donation specs and OpenAPI as references to reconcile the delivery contract?

This probe should answer:

1. what the public visitor must be able to understand;
2. what the frontend needs to deliver that experience;
3. what backend capabilities are actually required;
4. whether the current API contract is sufficient;
5. which old decisions survive, which need revision, and which are irrelevant to this slice.

It does **not** implement production code.

## 2. Upstream authority used

### Candidate Product Authority

`docs/product/product-overview.md`

Relevant durable product truths:

- Kencleng is **Evidence-Led Optimism**;
- confidence should precede conversion;
- facts come first, with story attached rather than facts being decorative support for a persuasive story;
- trust comes from ordered truth, provenance, chronology, state meaning, and candor;
- funding progress is not operational progress or impact;
- unknown/pending information is legitimate;
- accountability continues after donation and after fundraising closes;
- Organization review, Campaign curation/publication, and later accountability review are different concepts;
- Kencleng facilitates and structures trust but does not guarantee real-world outcome.

### Canonical Product Design / Brand Authority

Primary relevant sources:

- `docs/ui-ux/brand-product-ui-brief.md`
- `docs/ui-ux/product-design-principles.md`
- `docs/ui-ux/page-map.md`
- `docs/ui-ux/patterns.md`
- `docs/ui-ux/design-guidelines.md`
- `docs/ui-ux/asset-governance.md`

Relevant experience intent:

- Public Campaign Detail exists for a Guest/Public Visitor to understand campaign purpose, steward/organizer context, funding context, story, and available donation action.
- Public surfaces may be editorially expressive but must preserve truth hierarchy.
- Real campaign imagery may support human understanding, but synthetic imagery must never masquerade as campaign evidence.
- A truthful placeholder is preferred when real campaign media is unavailable.
- Consequential next actions should be clear but non-coercive.

## 3. Product slice definition

### Actor

**Guest / Public Visitor**, including a skeptical-but-open potential donor.

### Primary outcome

The visitor can understand enough about a publicly available campaign to make a considered next decision without needing to infer trust from badges, emotional pressure, popularity, or unsupported impact claims.

### This slice owns

- public understanding of one campaign;
- organizer/steward context relevant to that decision;
- current funding context;
- campaign story/purpose;
- truthful campaign media state;
- current public lifecycle/accountability context;
- the relationship to a valid next action.

### This slice does not own

- the Donation Flow form or donation settlement behavior;
- Campaign Discovery/ranking/search;
- authenticated campaign-management actions;
- campaign curation operations;
- backend package/domain decomposition;
- donor-list/social-proof presentation unless later justified by product need;
- a generic public Organization Profile.

## 4. Recommended public lifecycle model

### Material finding

The product thesis says accountability continues after fundraising. Historical Phase 3 reference also describes a public archival campaign/result summary after campaign closure.

The current Campaign Detail contract, however, is public only while `Campaign.status = published`. A `closed` campaign becomes non-public under the current detail visibility rule unless the viewer has an internal relationship.

That is a product/contract contradiction.

### Recommended direction

Use **one persistent public campaign identity/URL across the public lifecycle**.

Conceptually:

```text
published
→ public campaign detail
→ donations may be available
→ funding progress is visible

closed
→ same public campaign identity remains reachable
→ donation action is gone
→ final funding result / closure context is visible
→ available accountability/reporting context becomes more important
```

Other non-public workflow states such as draft/pending/rejected need not become public merely to support this continuity.

### Why this is preferred

- preserves links/bookmarks shared while fundraising was active;
- gives donors/public a stable place to return after closure;
- aligns with the product principle that accountability continues after donation;
- avoids making the campaign disappear exactly when result/accountability becomes relevant;
- keeps provenance/history attached to the campaign rather than splitting it into an unrelated experience.

### Main downside

The detail contract and UI become lifecycle-aware rather than describing only the fundraising-active state. That increases contract/state complexity and requires deliberate closed-state semantics.

### Alternative not preferred

A separate public campaign-result/summary route could own the closed state.

If chosen, the original public campaign URL should still preserve continuity through an intentional redirect/link rather than becoming inaccessible. This adds an IA/contract split without an obvious product benefit for v1, so the persistent-detail model is preferred for now.

## 5. Frontend delivery needs

This section describes experience requirements, not exact component/layout design.

### 5.1 Information hierarchy

A successful public detail experience should make the following understandable in a coherent order:

1. **Campaign identity and purpose**
   - title/purpose;
   - applicable public lifecycle state;
   - enough context to understand what the fundraising is for.

2. **Campaign media state**
   - real campaign media when available and truthful;
   - clear canonical placeholder when absent;
   - no generated beneficiary/documentary imagery presented as campaign evidence.

3. **Steward / organizer context**
   - which Organization is responsible for the campaign;
   - any public organization-review fact must be expressed with its exact meaning rather than collapsed into a broad “verified campaign” claim.

4. **Funding context**
   - amount/target/progress meanings must be explicit;
   - progress is funding evidence only;
   - do not imply execution or impact from funding percentage.

5. **Story and relevant campaign context**
   - organizer-provided campaign narrative/content may add human context;
   - presentation should preserve the difference between organizer-provided content and system/platform facts.

6. **Accountability / provenance context**
   - while fundraising is active: show only accountability facts that actually exist;
   - after closure: surface available result/reporting context and explicitly preserve pending/unknown states.

7. **Meaningful next action**
   - while donations are genuinely available: donation is a clear but calm primary next action;
   - once closed/not accepting donations: do not show a donate action; prioritize result/accountability/follow-up instead.

### 5.2 Required frontend states

At minimum:

- loading/initial fetch state where applicable;
- published campaign with real media;
- published campaign without media;
- closed/public campaign with result/accountability context;
- closed/public campaign where later accountability information is still pending/unavailable;
- not found / no public campaign at the requested identifier;
- recoverable request failure;
- responsive mobile and desktop behavior preserving purpose, steward, funding meaning, trust context, and reachable next action.

A public visitor should not receive internal non-public lifecycle detail merely because a record exists.

### 5.3 Explicit non-requirements

This slice does **not** need to invent:

- popularity;
- donor velocity;
- urgency countdown language beyond truthful date/state context;
- ranking;
- recommendation;
- impact score;
- generic trust score;
- donor list as social proof;
- synthetic campaign photography.

### 5.4 Donation action delivery boundary

Public Campaign Detail may be implemented before the full Donation Flow implementation is finished, but it must not ship a dead/fake donation action.

Recommended delivery rule:

```text
Public Campaign Detail can progress independently
→ render/data behavior may be verified
→ but the complete public surface is not DELIVERED with an active donation CTA
  until that CTA leads to a real supported Donation Flow
```

A placeholder route, fake submit behavior, or production `mockMode` branch is not an acceptable substitute.

This is a delivery dependency, not a reason to pull donation submission logic into the Campaign Detail slice.

## 6. Backend capability needs

Derive capabilities before endpoint shape.

### Capability A — Resolve a public campaign identity

Given a campaign identifier, determine whether the campaign is publicly viewable in its current lifecycle state.

Public visibility must support the product's chosen public lifecycle, not merely the historical `published` implementation assumption.

### Capability B — Provide campaign understanding data

Provide enough campaign data to understand:

- campaign identity/purpose/story;
- public lifecycle state;
- relevant dates;
- funding model/context;
- responsible Organization projection.

The backend does not need to decide visual hierarchy.

### Capability C — Provide funding facts

Provide trustworthy current/final funding facts needed by the public experience.

Funding facts must remain semantically distinct from operational/accountability outcomes.

### Capability D — Provide public campaign media when available

Expose real campaign media that the viewer is authorized to see under the same public-visibility semantics as the parent campaign.

Missing media is valid and should not require backend invention of a replacement asset.

### Capability E — Provide post-closure result/accountability context

For a campaign that remains public after closure, provide the public result/accountability facts that exist at that stage, while allowing information to remain pending/unavailable.

The exact shape should be derived only to the resolution needed by the public surface. Do not design the entire reporting/disbursement domain here.

### Capability F — Expose donation eligibility as a truthful state

The detail experience must be able to know whether donation is currently a valid next action.

This may be derivable from a canonical campaign lifecycle/public eligibility state rather than requiring a second frontend-owned business rule.

The actual donation submission capability remains owned by the Donation delivery slice.

## 7. Existing contract reconciliation

Existing artifacts are references, not automatic authority for the new slice.

### 7.1 `GET /campaigns/{campaignId}`

**Useful and likely retained concept:** a deliberate composite read for campaign + Organization summary + progress is a good fit for a public detail surface and avoids unnecessary client-side composition.

**Current fit:** strong for an active `published` campaign.

**Material gap:** current public visibility excludes `closed`, conflicting with the product's post-campaign accountability/continuity direction.

**Reconciliation result:** **ADAPT**, not KEEP-as-is.

### 7.2 `CampaignDetail` composite

The existing `CampaignDetail` concept appears directionally useful because it already composes Campaign, minimal Organization context, and progress.

For the active-fundraising state, it is likely close to sufficient.

For the persistent public closed state, the contract must be checked/extended only as needed for:

- closure/result context;
- any public campaign-result/accountability fields required by the surface;
- source/provenance distinctions when they cannot be derived safely from contract semantics.

Do not add a full Organization or full reporting domain payload merely because the page is composite.

### 7.3 `GET /campaigns/{campaignId}/attachments`

The endpoint/capability is useful when real campaign media exists.

There is already an internal reference contradiction:

- current OpenAPI declares the list endpoint as public (`security: []`);
- Campaign media feature spec says list visibility must match Campaign Detail: public only when the parent campaign is publicly visible, otherwise gated.

**Reconciliation result:** **ADAPT**. Media visibility should follow the reconciled parent-campaign public lifecycle semantics.

### 7.4 Donation submission contract

`POST /campaigns/{campaignId}/donations` is **not** required to render Public Campaign Detail and should not be pulled into this slice's backend implementation merely to satisfy a button.

It is a downstream action dependency for the frontend surface.

Detailed donation amount, payment method, guest identity, settlement simulation, token, and claim semantics remain subject to their own delivery-slice revalidation.

### 7.5 Public donor list

The existing `GET /campaigns/{campaignId}/donations` capability is not required by the current Public Campaign Detail product outcome.

Do not include it by default merely because it exists or because social proof is common on fundraising sites.

**Reconciliation result:** **DEFER** for this slice.

## 8. Contract changes implied by the probe

Do **not** edit OpenAPI until the product decision in this probe is accepted.

If the recommended persistent public lifecycle is accepted, contract work should be scoped narrowly:

1. reconcile public visibility for `GET /campaigns/{campaignId}` so the appropriate closed/public state remains readable;
2. define the minimal closed-state/result/accountability data required by the surface;
3. reconcile campaign-media list visibility to the same parent visibility rule;
4. expose donation eligibility through canonical lifecycle semantics without duplicating business logic in the frontend;
5. preserve non-public visibility gates for workflow states that should not become public.

This does **not** justify rewriting all Campaign OpenAPI or all Campaign domain specs.

## 9. Integration-map candidate

Once the Product Authority reframe is promoted and this probe decision is accepted, the first integration-map relationship should be approximately:

| Frontend surface / flow | Product capability | API contract / operation | Backend owner(s) | Contract gap / coordination note |
|---|---|---|---|---|
| Public Campaign Detail | Public campaign understanding + funding/accountability continuity | `GET /campaigns/{campaignId}`; optional campaign-media list; Donation Flow is a separate next-action dependency | Campaign, with Donation-derived funding facts and later accountability sources as needed | Existing detail visibility is `published`-only and must be reconciled with closed/public accountability continuity; media visibility must follow parent visibility |

Do not add readiness/status columns.

## 10. Probe verdict

### Did the new hierarchy produce useful information?

**Yes.**

Starting from product + design before contract detail surfaced material issues that were easy to miss when the existing Campaign domain spec/OpenAPI were treated as the starting authority:

1. public Campaign Detail currently disappears after closure even though post-campaign accountability is a core product proposition;
2. campaign-media visibility is internally inconsistent between OpenAPI and the later feature-spec decision;
3. the Donation CTA is a real cross-slice dependency and must not become fake frontend behavior;
4. a public donor list is not automatically required just because the API already offers one.

### Is the existing contract ready unchanged?

**No.**

The core composite-detail idea is reusable, but the public lifecycle and media visibility require reconciliation before calling the surface contract-ready.

### Did the probe require a full-product respec?

**No.**

The issues are narrow and local to the real slice, supporting the progressive-commitment model.

## 11. Human decision gate

Before contract editing or production implementation, confirm the recommended product direction:

> **A Campaign that was publicly available remains publicly reachable after fundraising closes, using the same public campaign identity/URL, with donation action removed and post-campaign result/accountability context taking priority.**

If accepted, this becomes the product decision to promote into Product Authority and then reconcile into Campaign/API delivery specs.

If rejected, choose and define the alternative public post-campaign continuity model before editing the contract.
