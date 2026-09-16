# Kencleng — Persona × Surface Map

> Status: Canonical
> Purpose: Map user personas to product surfaces and user goals.
> Boundary: This document does not own route implementation, shell layout, visual composition, frontend architecture, or business rules.

This map is derived from current product/domain behavior and exists to keep product surfaces coherent across personas.

Exact route paths may evolve during frontend re-architecture. If a route path conflicts with a canonical product/domain requirement, product truth wins.

## 1. Guest / Public Visitor

Primary goals:

- understand Kencleng and how trust/transparency works;
- discover published campaigns;
- inspect campaign purpose, organizer context, progress, and relevant public reporting;
- donate without requiring an account when product rules allow;
- track an unauthenticated donation when the domain provides a safe tokenized flow;
- authenticate or create an account when desired.

Expected product surfaces:

| Surface | UX pattern | Primary purpose |
|---|---|---|
| Public home | Editorial public / List-Browse composition | Understand platform, trust model, and discover campaigns |
| Campaign discovery | List / Browse | Browse published campaigns |
| Public campaign detail | Detail | Understand campaign, steward, funding context, story, and available donation action |
| Donation flow | Form | Submit a donation with clear amount/payment consequences |
| Donation status/tracking | Status / Tracking | Understand transaction state and next action |
| Authentication | Form | Login/register/recovery flows |

Public surfaces must not expose non-public campaign content or imply organizer/campaign trust states beyond canonical product truth.

## 2. Registered Donor

Inherits public capabilities plus authenticated donor concerns.

Primary goals:

- view personal donation history;
- understand current donation/payment states;
- follow campaign/program updates after donation;
- manage profile/security/notification concerns where supported;
- claim eligible guest donations when the domain supports it.

Expected product surfaces:

| Surface | UX pattern | Primary purpose |
|---|---|---|
| Donation history | List / Browse | Review personal donation records |
| Donation detail / Evidence Journal | Evidence Journal / Detail | Follow donation fact, campaign progress, milestones, reports, and pending updates |
| Donation claim flow | List / Browse + confirmation | Review and claim eligible guest donations |
| Profile | Form | Manage supported profile information |
| Security | Form / settings | Manage supported authentication/security methods |
| Notifications | List / Browse | Review product notifications |

The Evidence Journal must distinguish funding progress, operational progress, and reported outcome.

## 3. Organization Owner

Primary goals:

- establish and maintain organization identity;
- manage representatives according to domain permissions;
- create and revise campaigns;
- submit campaigns to curation;
- publish/schedule/unpublish when allowed;
- monitor campaign funding/activity;
- submit disbursement and fund-usage reporting when applicable;
- provide campaign/report narratives where the domain permits.

Expected product surfaces:

| Surface | UX pattern | Primary purpose |
|---|---|---|
| Organization registration | Form — Revisable Submission | Submit organization information/evidence |
| Organization detail | Detail | Understand organization state and allowed next actions |
| Representatives | List / Browse + Form | Manage permitted representatives |
| Campaign creation/edit | Form — Revisable Submission | Create/revise campaign data |
| Campaign owner detail | Detail | Understand campaign state and available lifecycle actions |
| Publication controls | Form / consequential action | Publish, schedule, reschedule, or unpublish according to product truth |
| Campaign monitor | Dashboard / Summary | Understand funding/current campaign state |
| Campaign report | Detail + permitted edit action | Review report data and supported narrative |
| Disbursement request | Form — Revisable Submission | Submit and follow disbursement request |
| Fund-usage report | Form / Detail — Revisable Submission | Submit and follow expense/accountability reporting |

## 4. Organization Staff

Primary goals overlap with Organization Owner, but actions must reflect the narrower permission model defined by domain authority.

Staff surfaces may reuse Owner surface structures while hiding or disabling only actions they are not permitted to perform.

Frontend visibility is UX only and never substitutes for backend authorization.

Common staff concerns include:

- create/edit allowed campaign drafts;
- view organization/campaign information permitted to staff;
- view monitoring/reporting information;
- create events where supported;
- avoid access to owner-only legal, representative-management, publication, disbursement, or narrative actions when the domain restricts them.

## 5. Curator

Primary goals:

- review assigned organization/campaign/fund-usage material;
- understand submission evidence and context;
- approve/reject according to domain rules;
- provide decision reasoning where required;
- understand already-decided states without misleading active controls.

Expected product surfaces:

| Surface | UX pattern | Primary purpose |
|---|---|---|
| Curation queue | List / Browse | Review assigned work by real domain type/state |
| Organization review | Curation / Review | Review organization evidence and decide |
| Campaign review | Curation / Review | Review campaign content and decide |
| Fund-usage review | Curation / Review | Review accountability report/evidence and decide |

Organization verification, campaign curation, and fund-usage verification remain distinct product concepts even if they share a reusable review pattern.

## 6. Admin

Primary goals depend on product/domain authority and may include:

- user/role administration;
- curation assignment;
- exceptional campaign lifecycle action;
- disbursement review;
- other explicitly defined operational controls.

Expected product surfaces:

| Surface | UX pattern | Primary purpose |
|---|---|---|
| User/role administration | List / Browse | Find users and manage supported roles |
| Curation assignment | List / Browse | Assign real pending work to curators |
| Exceptional campaign action | Curation / Review or consequential action | Perform supported administrative lifecycle action |
| Disbursement review | Curation / Review | Review and decide disbursement requests |

Admin UI must not invent metrics or “control center” dashboards merely because admin products commonly have them.

## 7. Cross-Cutting Surface Concepts

### Trust / Transparency Context
Trust information should be distributed near relevant decisions and may also have dedicated deeper detail.

### Evidence Journal
A donor-facing post-donation product concept for chronological factual updates, reports, provenance, and pending next states.

### Secure/Sensitive Data Presentation
PII, legal documents, tokens, payment/security information, and other sensitive data must follow domain/security authority. Design may clarify access but cannot authorize it.

### Curation Decision Experience
Shared interaction semantics may exist across multiple review domains without collapsing their business meaning into one generic “verification”.

## 8. What This Map Does Not Decide

This document intentionally does not establish:

- exact route names;
- navigation/shell architecture;
- mobile drawer behavior;
- frontend folder structure;
- component decomposition;
- PWA/offline behavior;
- mock-data behavior;
- concrete visual layout;
- backend product concepts that do not exist in canonical domain authority.

Those concerns belong to their respective authorities or future engineering/design-system derivation.

## 9. Open Product/Surface Questions

Items that remain unresolved must stay unresolved rather than being inferred from legacy frontend/prototypes.

Examples may include:

- campaign discovery prioritization/curation semantics when not defined by backend/product authority;
- future dedicated public organization profile behavior;
- future event/public discovery behavior;
- exact post-donation Evidence Journal API support and source labeling where not yet represented in canonical contracts.

These are product/engineering follow-ups and must not be answered by visual precedent alone.
