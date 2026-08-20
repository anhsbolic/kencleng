# Feature Spec — 10: Events

> File: `docs/spec/campaign/features/10-events.md`
> Domain: `campaign`
> Task: 10 (see `docs/spec/campaign/tasks.md`)
> Status: draft — authored against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`POST /events` — a representative registers a lightweight promotional
Event, linking it to ≥1 of their own organization's campaigns (must be
`published`/`scheduled`). No curation gate — active immediately. `GET
/events/{eventId}` is fully public. `GET
/organizations/{organizationId}/events` is the Owner/Staff dashboard
listing (`# INFERRED`).

## Endpoints

`POST /events`, `GET /events/{eventId}`, `GET
/organizations/{organizationId}/events` (confirmed,
`api/openapi/campaign.yaml`)

## Auth

- Create: `bearerAuth` + representative (any `level`) of the
  organization owning the linked campaign(s).
- Detail: none (`security: []`) — public.
- Org-scoped list: `bearerAuth` + representative.

## Request

### Create — `EventCreateRequest`
| Field | Type | Required | Notes |
|---|---|---|---|
| `name` | string | Yes | |
| `event_datetime` | datetime | Yes | |
| `location` | string | No | |
| `description` | string | No | |
| `campaign_ids` | UUID[] | Yes | Min 1; all must belong to the same organization as the creator, `status` ∈ {`published`, `scheduled`} |

## Behavior

### Create
1. Reject (`403`) if caller isn't a representative of the organization
   that owns **every** campaign in `campaign_ids` — if any campaign
   belongs to a different organization, reject.
2. Reject (`409`) if any `campaign_ids` entry has `status` outside
   {`published`, `scheduled`}.
3. Insert `events`, insert one `campaign_events` row per
   `campaign_id`.
4. Return `201` with the `Event` (includes `campaign_ids`).

### Detail
1. Load the event — `404` if it doesn't exist.
2. Return `Event` — no gating, regardless of linked campaigns' current
   status (accepted design, per `invariants.md`'s note — Events carry
   no sensitive data).

### Org-scoped list
1. Reject (`403`) if caller isn't a representative.
2. Query events linked (via `campaign_events`) to any campaign owned
   by this organization.
3. Return paginated list.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token (create, org-list) | `401` |
| Any `campaign_ids` entry belongs to a different organization | `403` |
| Any `campaign_ids` entry isn't `published`/`scheduled` | `409` |
| `campaign_ids` empty | `422` (schema `minItems: 1`) |
| Event doesn't exist (detail) | `404` |

## Concurrency & correctness notes

- **No re-check after linking**: if a linked campaign's status later
  changes (e.g. closes), the `campaign_events` relation is not
  retracted — the Event continues referencing it. Confirmed accepted
  (not a gap) per Events' low-sensitivity design — no test needed to
  "fix" this, only to confirm it's the actual (non-)behavior.

## Test checklist

- [ ] Linking a campaign from a different organization → `403`.
- [ ] Linking a `draft`/`pending_curation`/`rejected`/`closed`
      campaign → `409`.
- [ ] `campaign_ids` empty → `422`.
- [ ] Successful creation with multiple valid campaign links.
- [ ] Event detail is publicly readable without auth.
- [ ] A closed campaign that was linked before closing remains linked
      (no automatic retraction) — confirm this is the intended
      behavior, not silently broken.
- [ ] Org-scoped list returns events linked to any of that
      organization's campaigns.

## References

- `docs/spec/campaign/invariants.md` — INV-campaign-16
- `docs/spec/campaign/threat-model.md` — "Events" section
- `docs/spec/campaign/tasks.md` — Task 10
- `api/openapi/campaign.yaml` — `Event`, `EventCreateRequest` schemas
- `docs/project/kencleng-phase1-detail.md` — Fitur 6
