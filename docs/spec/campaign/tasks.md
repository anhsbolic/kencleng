# Domain Tasks — campaign

> File: `docs/spec/campaign/tasks.md`
> Status: draft — authored directly against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-08-20

Task order follows dependency: creation and read paths first, then
the linear lifecycle (submit → curate → publish → close), then the
independent Event entity last (depends only on campaigns existing).

| # | Task | Endpoint / surface | Depends on | Related invariants |
|---|---|---|---|---|
| 01 | Campaign creation & draft CRUD | `POST /organizations/{id}/campaigns`, `PATCH/DELETE /campaigns/{id}` | organization domain (Task 01/04) | INV-campaign-01, 02, 03, 04 |
| 02 | Campaign detail & listing | `GET /campaigns`, `GET /organizations/{id}/campaigns`, `GET /campaigns/{id}` | 01 | INV-campaign-14 |
| 03 | Campaign media | `GET/POST /campaigns/{id}/attachments` | 01, 02 | INV-campaign-14 (resolved 2026-08-20) |
| 04 | Submit for curation | `POST /campaigns/{id}/submit` | 01 | INV-campaign-03, 04 |
| 05 | Curation assignment | `POST /campaigns/{id}/curation/assign`, `GET /campaigns/curation-queue` | 04 | INV-campaign-05, 06 |
| 06 | Curation decision | `POST /campaigns/{id}/curation/decision`, `GET /campaigns/curation-assignments/mine` | 05 | INV-campaign-07 |
| 07 | Publish/schedule/reschedule/republish | `POST /campaigns/{id}/publish` | 06 | INV-campaign-08, 09 |
| 08 | Unpublish (manual) | `POST /campaigns/{id}/unpublish` | 07 | INV-campaign-10 |
| 09 | Closure (auto + force-close) | `POST /campaigns/{id}/force-close` + scheduler + donation-trigger hook | 07 | INV-campaign-13 |
| 10 | Events | `POST /events`, `GET /events/{id}`, `GET /organizations/{id}/events` | 01, 07 | INV-campaign-16 |

## Task 01 — Campaign creation & draft CRUD

**What**: `POST /organizations/{organizationId}/campaigns` (create,
any representative), `PATCH`/`DELETE /campaigns/{campaignId}` (draft
only, any representative). Gated on `organization.status = 'verified'
AND has_overdue_report = false` (INV-campaign-01, references
INV-organization-13).

**KPI / metrics**:
- 0 successful creations against an unverified or overdue-report
  organization.
- 100% of field-validation rules (INV-campaign-02) enforced —
  `target_amount`, `max_amount`, `deadline`, `category`.
- 0 successful edits/deletes outside `status = draft`.
- `staff` can create/edit/delete drafts; only tested exclusion is
  submit (Task 04).

## Task 02 — Campaign detail & listing

**What**: `GET /campaigns` (public, published-only, filterable by
`category`/`q`), `GET /organizations/{organizationId}/campaigns`
(Owner/Staff dashboard, any status, filterable by `status`), `GET
/campaigns/{campaignId}` (composite detail — campaign + organization
summary + progress; public for `published`, gated otherwise).

**KPI / metrics**:
- Public list returns only `published` campaigns.
- Org-scoped list returns all statuses, scoped correctly, `403` for
  non-representatives.
- Detail endpoint: public for `published`; `403` for non-`published`
  when requester lacks visibility (representative/Kurator/Admin).
- `CampaignDetail`'s embedded `organization`/`progress` data is
  internally consistent with what `organization`/`donation` domains
  would report independently (composite-endpoint correctness check).

## Task 03 — Campaign media

**What**: `GET`/`POST /campaigns/{campaignId}/attachments`. Upload:
JPG/PNG only, 5 MB max, owner/staff representative. List: gated to
match the detail endpoint's visibility rule
(`[RESOLVED — 2026-08-20]` — public only when `status = 'published'`,
otherwise representative/Kurator/Admin only). **Implementation note**:
`api/openapi/campaign.yaml` currently has `security: []` on the list
endpoint — needs updating to remove that and add the same auth/gating
logic as Task 02's detail endpoint.

**KPI / metrics**:
- Non-representative/non-staff upload attempt → `403`.
- Invalid file type/oversized file → `422`.
- List gating matches detail endpoint exactly, same test matrix
  (public-when-published, restricted otherwise).

## Task 04 — Submit for curation

**What**: `POST /campaigns/{campaignId}/submit` — owner-only,
`draft → pending_curation`.

**KPI / metrics**:
- `staff` attempt → `403`.
- Attempt outside `status = draft` → `409`.
- Successful submit is atomic (status-guard at the query level, no
  double-submit race).

## Task 05 — Curation assignment

**What**: `POST /campaigns/{campaignId}/curation/assign` (Admin-only),
`GET /campaigns/curation-queue` (Admin-only, `pending_curation` +
unassigned).

**KPI / metrics**: same shape as
`organization/tasks.md` Task 09 — 0 successful assignments with
conflict of interest (Kurator is a representative of the campaign's
org), 0 successful assignments creating a second pending assignment.

## Task 06 — Curation decision

**What**: `POST /campaigns/{campaignId}/curation/decision` (assigned
Kurator only, server-resolved current assignment), `GET
/campaigns/curation-assignments/mine` (Kurator's own queue).

**KPI / metrics**: same shape as `organization/tasks.md` Task 10,
including the reassignment-race test (Admin reassigns while the
originally-assigned Kurator submits concurrently).

## Task 07 — Publish/schedule/reschedule/republish

**What**: `POST /campaigns/{campaignId}/publish` — one endpoint
handling all four cases via the same contract (INV-campaign-08).
Includes the idempotent scheduler job for time-triggered
`scheduled → published` (INV-campaign-09).

**KPI / metrics**:
- `publish_at` validation (future, ≤ deadline) enforced.
- Reschedule only valid while `scheduled`.
- Republish only valid from `unpublished`.
- Scheduler job run twice near-simultaneously → exactly one
  transition, no error on the second run.

## Task 08 — Unpublish (manual)

**What**: `POST /campaigns/{campaignId}/unpublish` — owner-only,
requires `decision_note`, `unpublish_reason = 'owner_manual'`, logged.

**KPI / metrics**:
- Missing/empty `decision_note` → `422`.
- Attempt outside `status = 'published'` → `409`.
- `staff` attempt → `403`.

**Cross-domain note**: also covers verifying INV-campaign-11
(auto-unpublish from organization re-verification) — this is where
the full integration test unblocked by both domains existing should
live, per `invariants.md`'s note.

## Task 09 — Closure (auto + force-close)

**What**: `POST /campaigns/{campaignId}/force-close` (Admin-only,
requires `decision_note`, records `closed_by`), plus the two automatic
triggers: `max_amount`-reached (hook into `donation` domain's
donation-success transaction — cross-domain, `donation` not yet
spec'd) and deadline-reached (periodic scheduler).

**KPI / metrics**:
- Non-Admin force-close attempt → `403`.
- Missing `decision_note` → `422`.
- **3-way concurrency test** (explicit, per `threat-model.md`'s note):
  simulate max_amount-trigger, deadline-trigger, and force-close all
  firing near-simultaneously on the same campaign — assert exactly
  one `closed_reason` recorded, the other two get clean `409`s, no
  crash.
- Deadline scheduler is idempotent (same pattern as Task 07's publish
  scheduler).

## Task 10 — Events

**What**: `POST /events` (owner/staff, links ≥1 campaign),
`GET /events/{eventId}` (public), `GET
/organizations/{organizationId}/events` (Owner/Staff dashboard,
`# INFERRED`).

**KPI / metrics**:
- 0 successful links to a campaign from a different organization
  (`403`).
- 0 successful links to a campaign outside {`published`, `scheduled`}
  (`409`).
- Event detail publicly readable regardless of linked campaigns'
  current status (confirmed accepted risk, `threat-model.md`).

## References

- Related domain invariants: `docs/spec/campaign/invariants.md`
- Related threat model: `docs/spec/campaign/threat-model.md`
- **Actual API (ground truth)**: `api/openapi/campaign.yaml`
- Feature specs (one per task, prefixed with the task `#` above):
  `docs/spec/campaign/features/`
