# Kencleng — Integration Map

> Status: Living cross-stack coordination map
> Last updated: 2026-09-17
> Purpose: Map frontend surfaces/flows to product capabilities, API contracts, and backend ownership without forcing frontend and backend to share the same decomposition.

## 1. What this document owns

This document owns **cross-stack relationship mapping** only.

It answers questions such as:

- which product capability a frontend surface is expressing;
- which API contract(s) that surface depends on;
- which backend domain/use-case owner(s) provide those capabilities;
- where a contract gap is known before implementation/integration.

It does **not** own:

- business rules or invariants;
- API request/response schema;
- frontend page/interaction design;
- backend architecture;
- frontend architecture;
- task acceptance criteria;
- delivery/readiness status.

Those remain owned by their existing canonical sources.

Principle:

> Decomposition is local. Contracts are shared.

Backend may remain domain-driven while frontend remains experience/page-map-driven. This map is the rendezvous layer between them, not a replacement decomposition for either side.

## 2. Authority routing

Use the owning source rather than copying its content here:

```text
product/domain truth      → docs/spec/<domain-dir>/
API schema                → api/openapi/<domain>.yaml + referenced common.yaml components
aggregate OpenAPI view    → api/openapi.yaml when cross-domain/generated-bundle context is needed
frontend surface intent   → docs/ui-ux/page-map.md + relevant UI/UX authorities
backend architecture      → docs/project/kencleng-backend-tech-stack.md
frontend architecture     → docs/project/kencleng-frontend-tech-stack.md
project delivery status   → docs/project/kencleng-development-tracker.md
```

Do not add readiness/status columns here. `kencleng-development-tracker.md` remains the single project-level status authority.

## 3. Mapping model

Each active mapping should use the smallest useful row:

| Frontend surface / flow | Product capability | API contract / operation | Backend owner(s) | Contract gap / coordination note |
|---|---|---|---|---|

A row may reference more than one backend domain when the user-facing surface composes multiple capabilities.

Do not force one frontend page to mirror one backend domain, aggregate, service, or endpoint.

Do not force a backend domain to mirror frontend page structure.

## 4. Population rule

Populate this map **lazily**, when a frontend surface/flow enters real planning or implementation and cross-stack dependencies become actionable.

Do not pre-map the whole product speculatively.

For a new frontend surface:

```text
page/surface selected
→ identify product capability
→ identify existing API contract(s)
→ identify backend owner(s)
→ record any genuine contract gap
→ execute frontend/backend in parallel when preconditions allow
→ verify real integration when both sides are available
```

If a required capability has no valid API contract yet, record the gap rather than inventing frontend-local business semantics.

## 5. Contract-parallel frontend rule

When the relevant API contract is stable enough but backend implementation is not yet available, frontend may proceed against contract-faithful mocks.

Preferred shape:

```text
Presentation
    ↓
query/data boundary
    ↓
real API request shape
    ↓
network boundary
   /             \
MSW mock       real backend
```

The production service/data layer should not contain a mock-mode branch that returns alternate JSON. MSW (or the approved equivalent) should intercept the real request contract at the network boundary.

This allows data/service concerns and presentation concerns to be implemented independently where useful while preserving one production integration path.

Architecture boundaries may be separate, but delivery should normally remain scoped to a meaningful page/flow/capability slice rather than becoming a project-wide horizontal phase such as “all services first, all presentation later.”

## 6. Active mappings

| Frontend surface / flow | Product capability | API contract / operation | Backend owner(s) | Contract gap / coordination note |
|---|---|---|---|---|
| Public Campaign Detail | A visitor understands one persisted eligible Campaign, its steward, funding truth, organizer provenance, and truthful media/action state without a fake donation flow. | `getPublicCampaignDetail`; `getPublicCampaignMediaContent` | Campaign public projection and controlled media delivery | Detail embeds the narrow public steward projection; no Organization-detail request is needed. Frontend consumes generated types and uses the opaque same-origin `content_url`; backend/topology must later enforce private storage, parent/member recheck, `/api` routing, and `private, no-store`. |

The Frontend Experience Foundation remains a cross-domain calibration/foundation task and does not need a synthetic API mapping retrofitted onto it.

## 7. Update discipline

When updating this map:

- reference canonical operation/capability names rather than copying schemas;
- keep mappings structural, not status-oriented;
- add only dependencies that matter to the active surface;
- remove or revise stale mappings when canonical specs/contracts change;
- surface genuine cross-domain composition rather than hiding it for symmetry;
- do not use this map to create a second backlog or tracker.
