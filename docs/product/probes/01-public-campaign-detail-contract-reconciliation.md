# Probe 01 — Public Campaign Detail Contract Reconciliation

> Status: **Probe evidence — candidate contract decision, not canonical contract**
> Created: 2026-09-17
> Human product decision: **APPROVED — 2026-09-17**
> Parent probe: `docs/product/probes/01-public-campaign-detail.md`
> Purpose: Record the narrow contract consequences of the approved public-lifecycle decision without editing the whole Campaign OpenAPI or treating this probe as long-term authority.

## 1. Approved product decision

A Campaign that has entered the public fundraising lifecycle remains publicly reachable after fundraising closes.

The public identity/URL remains stable across:

```text
published
→ public fundraising detail
→ donation may be available

closed
→ same public campaign identity
→ donation unavailable
→ final funding/result context
→ accountability/follow-up context becomes more important
```

This approval does **not** make all Campaign states public.

Draft, pending curation, approved-but-not-published, rejected, scheduled-before-publication, unpublished/retracted, and other non-public workflow states remain non-public unless a later product decision explicitly says otherwise.

For a closed Campaign, public continuity applies only when the Campaign was previously public. A backend implementation must not infer that every record with a terminal-looking state is safe to expose.

## 2. Safety principle — public continuity is not public data inheritance

The current `CampaignDetail` schema inherits the full internal `Campaign` object through `allOf`.

That is not acceptable as the long-term public contract because internal fields can become publicly exposed simply by being added to the internal schema.

The public API must use an **explicit allowlist projection**.

Principle:

> Public lifecycle continuity preserves the public Campaign identity, not the internal Campaign record shape.

A public projection must remain safe when the internal model evolves.

## 3. Recommended API separation

### Public read

The public Campaign-detail operation should return only a dedicated public schema, conceptually:

```text
PublicCampaignDetail
```

It should be usable for the two public lifecycle modes:

```text
fundraising
closed
```

The public operation must not return a different, more privileged payload merely because an Authorization header happens to be present.

### Internal/operational read

Organization representatives, Curators, and Admins may need richer non-public Campaign detail for their operational surfaces.

That capability should use a distinct authenticated operation/projection when real delivery requires it rather than overloading the public contract with internal fields.

This separation reduces accidental privilege expansion and makes public-response review auditable.

The exact internal operation/path is intentionally not designed by this probe.

## 4. Candidate `PublicCampaignDetail` projection

This is a contract-direction allowlist, not final OpenAPI syntax.

### Public campaign identity/content

Candidate public fields:

- campaign public identifier;
- title;
- public campaign description/story;
- category where still part of product experience;
- location when intentionally public;
- beneficiary description when it was approved/public campaign content;
- public lifecycle state;
- relevant public dates.

The contract should prefer a narrow public lifecycle semantic such as:

```text
fundraising
closed
```

over exposing the entire internal `CampaignStatus` enum.

The public client does not need to know internal workflow states that it can never access.

### Public steward/Organization projection

Candidate baseline:

- Organization public identifier;
- Organization public display name.

Do **not** automatically expose the full Organization object.

Do **not** automatically expose raw Organization internal status merely because the current `OrganizationSummary` contains it. If organization-review context is shown publicly, its exact public meaning/copy must be deliberately designed so it cannot be mistaken for Campaign verification or outcome guarantees.

### Funding projection

Candidate public funding facts:

- target amount;
- collected/final collected amount;
- truthful funding percentage/progress representation;
- deadline;
- closed timestamp when closed.

Additional values such as funding cap/max amount or donor count should be exposed only when the surface has a justified product use.

A field being factual is not by itself sufficient reason to publish it.

In particular, donor count must not become default social-proof/popularity pressure merely because the backend can calculate it.

### Donation eligibility

The frontend must not independently reconstruct donation permission from raw internal state.

The public contract should expose a backend-authoritative public eligibility semantic, for example a narrow capability such as:

```text
can_donate: true | false
```

or an equivalently explicit public action capability.

The exact final shape should be chosen when OpenAPI is reconciled, but production React must not own the business rule.

For the approved lifecycle:

```text
fundraising → may be true when all backend conditions permit
closed      → false
```

### Closure context

A closed Campaign should expose enough public information to explain that fundraising has ended.

Do **not** expose raw internal closure/admin fields by default.

The existing internal `closed_reason` values may encode operational detail. A public closure reason, if needed, should be a deliberately public-safe semantic rather than a direct passthrough of internal enums.

The public experience must be able to say truthfully that fundraising is closed even if a more detailed public reason is intentionally withheld or deferred.

### Accountability context

The closed public detail needs a truthful way to distinguish:

- accountability information currently available;
- accountability information not yet available/pending;
- information that is not part of this slice yet.

Do not infer "pending" merely because the frontend received no report payload.

If the closed-state surface requires a visible accountability state, the backend contract must represent that state explicitly.

This probe does not design the full reporting/disbursement payload. That remains progressive work for the relevant delivery slice.

## 5. Fields that must not leak through the public Campaign projection

The current internal `Campaign` schema includes fields that are internal, unnecessary, or unsafe to publish by inheritance.

The public projection must not automatically expose fields such as:

- `decision_note`;
- `created_by` internal user identifier;
- internal/unpublished workflow state;
- raw unpublish reason;
- curation assignment/reviewer information;
- internal decision metadata;
- internal audit information;
- private Organization/legal-document fields;
- PII;
- future internal fields merely because they are added to `Campaign`.

Timestamps should also be intentionally selected rather than inherited wholesale. Publicly meaningful timestamps may be exposed; implementation/audit timestamps do not automatically belong in the public product.

## 6. Non-public state behavior

For an unauthenticated/public caller, a Campaign that is not publicly visible should not reveal internal lifecycle detail.

Recommended public behavior:

```text
eligible public campaign
→ 200 PublicCampaignDetail

not found OR not publicly visible
→ public-safe not-found behavior
```

Do not use the public contract to confirm the existence or internal state of draft/rejected/unpublished Campaigns.

Authenticated operational access, when needed, belongs to a separate authenticated capability.

## 7. Campaign media safety

### 7.1 Public media requires a separate public projection

The current `CampaignAttachment` includes fields such as:

- `original_name`;
- `size_bytes`;
- `uploaded_by`;
- `created_at`.

A public campaign page normally needs far less.

Conceptual public projection:

```text
PublicCampaignMedia
- id
- presentation URL/reference
- media type only when the client materially needs it
- future public caption/alt/provenance fields only when explicitly defined
```

Do not expose uploader identity, original filename, or storage-oriented metadata by default.

### 7.2 Parent visibility must govern media visibility

Media for:

```text
published
closed-after-publication
```

may be public when it is approved campaign content.

Media for non-public/retracted Campaign states must not become discoverable merely because the media record exists.

### 7.3 Public-bucket URLs are not sufficient lifecycle authorization

The current historical design uses a publicly readable object bucket/direct URL.

That creates a lifecycle-safety problem:

```text
Campaign published
→ media URL becomes known
→ Campaign later unpublished/retracted
→ list endpoint can be gated
→ but known direct object URL may remain publicly readable
```

Therefore endpoint gating alone does not provide revocable visibility.

The delivery architecture should support actual revocation where product lifecycle requires content to stop being public. Viable mechanisms can be evaluated during backend delivery, for example controlled object access, signed URLs, or an application/media proxy. This probe does not prescribe the storage mechanism.

Hard requirement:

> Do not claim media is private/retracted if the underlying object remains anonymously readable through a previously known URL.

This requirement is especially important for legal/privacy/safety removal scenarios.

## 8. Published → closed versus unpublished/retracted

These states must not be conflated.

### Published → closed

This is the approved normal continuity path.

Previously public campaign content remains part of the public Campaign record, and the same public identity continues into result/accountability mode.

### Unpublished/retracted

This is not the same as normal fundraising completion.

The product may unpublish content for organizational re-verification, owner action, platform safety, legal/privacy reasons, or another defined cause.

This probe does not decide that previously public content must remain accessible after unpublish/retraction.

Therefore technical architecture must preserve the ability to remove public access rather than making publication permanently irreversible at the storage layer.

## 9. Candidate contract verdict

| Existing contract area | Verdict | Narrow reconciliation |
|---|---|---|
| `GET /campaigns/{campaignId}` composite concept | **ADAPT** | Keep a composite public read, but return explicit `PublicCampaignDetail`; public for fundraising + eligible closed archive only |
| Full `Campaign` inheritance in public response | **REPLACE for public contract** | Use an explicit allowlist public schema; keep internal model separate |
| Public lifecycle visibility | **ADAPT** | `published` remains public; `closed` remains public only when previously public; non-public workflow states stay hidden |
| Organization summary | **ADAPT** | Start with safe public identity; do not passthrough full/raw internal status without deliberate public semantics |
| Funding progress | **KEEP concept / NARROW fields** | Publish only facts needed for understanding; avoid social-proof extras by default |
| Donation eligibility | **ADAPT** | Backend-authoritative public capability; closed always non-donatable |
| `GET /campaigns/{campaignId}/attachments` | **ADAPT** | Parent lifecycle governs; return `PublicCampaignMedia`, not raw internal attachment shape |
| Public object bucket assumption | **RECONSIDER** | Must support real retraction/revocation; list gating is insufficient |
| Public donor list | **DEFER** | Not required for this slice |
| Full post-campaign accountability payload | **DEFER / narrow gap** | Expose only minimal explicit availability/state needed by closed detail; full reporting remains its own slice |

## 10. OpenAPI editing boundary

This reconciliation is sufficient to define the intended delta, but the existing OpenAPI is **not edited in this probe**.

Reason:

- Product Authority is still candidate/non-canonical during the reframe;
- the probe exists to validate the hierarchy first;
- editing the operational contract before authority promotion would create an ambiguous mixed-authority branch.

When the reframe is promoted, the Campaign contract change should be narrow and traceable to this decision:

1. introduce a public-safe detail projection;
2. define public lifecycle visibility including eligible closed archives;
3. remove raw internal-schema inheritance from public response;
4. reconcile public Organization projection;
5. reconcile public media projection and revocable access semantics;
6. expose backend-authoritative donation eligibility;
7. add only the minimal closed/accountability state needed by the actual public surface.

No all-Campaign or all-domain OpenAPI rewrite is justified by this probe.

## 11. Forward-probe conclusion

The forward derivation model has now demonstrated the intended behavior:

```text
Product Authority
+ Product Design / Brand Authority
→ real surface needs
→ FE delivery needs
→ BE capability needs
→ identify contradiction in old contract
→ human product decision
→ narrow contract reconciliation
```

It surfaced both a product-lifecycle issue and a non-obvious media-access safety issue without requiring a whole-product respec.

That is positive evidence for progressive commitment and for continuing to Probe 02 after this decision is recorded.
