# Domain Invariant — campaign

> File: `docs/spec/campaign/invariants.md`
> Status: draft — authored directly against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-08-20

## Domain summary

`campaign` owns campaign drafting, curation, publication lifecycle,
closure, campaign media, and the lightweight `Event`
promotional entity. Covers `campaigns`, `campaign_curation_assignments`,
`campaign_attachments`, `campaign_logs`, `events`, `campaign_events`.
Two invariants are **received** from `organization` (auto-unpublish on
re-verification, `has_overdue_report`/`status=verified` gating on
creation) — referenced here, not redefined, per this project's
cross-domain ownership convention. Authored directly against the
actual `api/openapi/campaign.yaml` (no separate reconciliation pass
needed, unlike `organization`).

## Open item flagged during authoring — resolved 2026-08-20

**`GET /campaigns/{campaignId}/attachments` had `security: []`
(fully public, no auth), with no visibility gate tied to the parent
campaign's `status`** — inconsistent with `GET
/campaigns/{campaignId}` itself, which *is* gated (non-`published`
campaigns restricted to representatives/Kurator/Admin). **Decided:
match the two endpoints' gating** — see INV-campaign-14. The current
`api/openapi/campaign.yaml` still needs a small update
(`security: []` removed from the attachment-list endpoint) to reflect
this at implementation time.

## Invariants

### INV-campaign-01: Campaign creation requires `organization.status = 'verified'` and `has_overdue_report = false`

- **Statement**: See **INV-organization-13** in
  `docs/spec/organization/invariants.md` — declared there in full
  (owned by `organization`, since it owns `has_overdue_report`). This
  entry is a reference, not a redefinition. Also requires
  `organization.status = 'verified'` specifically (confirmed
  `409 organization-not-verified` in the actual API, alongside
  `409 overdue-report`).
- **Holds after operations (this domain's responsibility)**: `POST
  /organizations/{organizationId}/campaigns` — must check both
  conditions at creation time, reading fresh (not cached) organization
  state.
- **Verification**: Test — attempt creation against a
  `pending_verification`/`rejected` org → `409
  organization-not-verified`. Attempt against a `verified` org with
  `has_overdue_report = true` → `409 overdue-report`. Existing
  campaigns under an org that later loses `verified` status or gains
  `has_overdue_report = true` are **not** retroactively affected
  (confirmed — `kencleng-phase1-detail.md` Fitur 3: "Existing
  campaigns... are not affected — this only blocks creating a new
  campaign").

### INV-campaign-02: Campaign field validation at creation/edit

- **Statement**: `target_amount > 0`; `max_amount`, if set, `≥
  target_amount`; `deadline` in the future at creation time; `category`
  required (fixed enum: `bencana_alam`/`kesehatan`/`pendidikan`/
  `sosial`/`lainnya`).
- **Holds after operations**: `POST
  /organizations/{organizationId}/campaigns`, `PATCH
  /campaigns/{campaignId}` (while `draft`).
- **Verification**: Test — each rule individually (zero/negative
  `target_amount`, `max_amount < target_amount`, past `deadline`,
  missing `category`) → `422`.

### INV-campaign-03: Draft campaigns are editable by any representative; submission is owner-only

- **Statement**: Create and edit (`PATCH`/`DELETE` while `status =
  draft`) are permitted for any representative (`owner` or `staff`) of
  the owning organization. Submit-for-curation
  (`POST /campaigns/{campaignId}/submit`) is **owner-only**.
- **Holds after operations**: creation, draft edit, draft delete,
  submit.
- **Verification**: Test — `staff` creates/edits/deletes a draft →
  succeeds. `staff` attempts submit → `403` ("Only `owner`-level
  representatives may submit for curation," confirmed explicit).
  `owner` submits → succeeds, `status → pending_curation`.

### INV-campaign-04: Draft/curation-locked fields are only mutable while `status = 'draft'`

- **Statement**: `PATCH`/`DELETE /campaigns/{campaignId}` are rejected
  (`409`) once `status != 'draft'` — a campaign in `pending_curation`
  or beyond cannot be edited or deleted directly; it can only move
  forward via curation decision or backward via the rejected→draft
  resubmission path (INV-campaign-06's state machine).
- **Holds after operations**: `PATCH`, `DELETE`.
- **Verification**: Test — attempt edit/delete on a
  `pending_curation`/`approved`/`published`/etc. campaign → `409`.

### INV-campaign-05: Kurator conflict-of-interest recusal

- **Statement**: A user may not be assigned as Kurator for a
  `campaign_curation_assignment` targeting a campaign whose owning
  organization they're also a representative of (any `level`). Same
  pattern as INV-organization-06, applied one level down (campaign's
  *organization*, not the campaign itself, since campaigns have no
  representatives of their own).
- **Holds after operations**: `POST
  /campaigns/{campaignId}/curation/assign`.
- **Verification**: Confirmed in the actual API description ("Fails
  if the chosen kurator is a representative of the campaign's
  organization"). Test: assign a representative of the campaign's org
  as its Kurator, assert `409`.

### INV-campaign-06: Only one active curation assignment per campaign

- **Statement**: `campaign_curation_assignments` may have at most one
  row with `decision = 'pending'` per `campaign_id`.
- **Holds after operations**: curation assignment creation.
- **Verification**: Confirmed — `409` "Conflict of interest, or an
  active assignment already exists" on the assign endpoint (same
  combined-error pattern as `organization`'s equivalent, worth
  splitting into distinct `type`s at implementation time for
  debuggability, mirroring `organization`'s `curator-conflict-of-
  interest`/`assignment-already-active` split).

### INV-campaign-07: Only the assigned Kurator can decide; rejection requires a note

- **Statement**: `POST /campaigns/{campaignId}/curation/decision`
  succeeds only for the campaign's *current* pending assignment's
  `kurator_id` (server-resolved, same pattern as
  `organization`'s decision endpoint — see that domain's TOCTOU note,
  applies identically here). `decision = 'rejected'` requires
  `decision_note`.
- **Holds after operations**: curation decision.
- **Verification**: Confirmed — `403` for a non-matching Kurator,
  `422` for missing `decision_note` on reject, `409` if no pending
  assignment exists. Test: same shape as
  `organization/features/10-curation-decision.md`'s test checklist,
  including the reassignment-race test.

### INV-campaign-08: Publish/schedule/reschedule/republish is one endpoint, owner-only, with a valid `publish_at`

- **Statement**: `POST /campaigns/{campaignId}/publish` handles all
  four cases (publish now, schedule, reschedule, republish) via the
  same contract (`# INFERRED` in the actual spec — "phase doc
  describes these as related but doesn't state whether they're
  literally the same endpoint," implemented as one endpoint). Valid
  from `status = 'approved'` (first publish) or `'unpublished'`
  (republish); while `'scheduled'`, calling again reschedules. If
  `publish_at` is provided, it must be `> now()` and `≤` the
  campaign's own `deadline` (`422` otherwise). Owner-only.
- **Holds after operations**: publish/schedule/reschedule/republish.
- **Verification**: Test — publish with no `publish_at` on an
  `approved` campaign → `status = published`. Publish with a future
  `publish_at` → `status = scheduled`. Reschedule while `scheduled` →
  `publish_at` updates, `status` unchanged. Republish from
  `unpublished` → `scheduled`/`published`. `publish_at` in the past,
  or beyond `deadline` → `422`. Wrong starting `status` (e.g. `draft`)
  → `409`. `staff` caller → `403`.

### INV-campaign-09: Scheduled publish is idempotent against concurrent scheduler runs

- **Statement**: The scheduler job transitioning `scheduled →
  published` at `publish_at` uses a conditional update (`WHERE status
  = 'scheduled' AND publish_at <= now()`) — running the job more than
  once around the same time never double-publishes or errors.
- **Holds after operations**: the scheduled-publish background job.
- **Verification**: Test — run the job twice in quick succession
  against the same due campaign; assert only one transition occurs,
  second run is a no-op, no error.

### INV-campaign-10: Manual unpublish requires a reason and is owner-only

- **Statement**: `POST /campaigns/{campaignId}/unpublish` requires
  `decision_note` (non-empty), sets `unpublish_reason = 'owner_manual'`,
  logged to `campaign_logs`/Audit Log. Owner-only. Only valid from
  `status = 'published'`.
- **Holds after operations**: manual unpublish.
- **Verification**: Confirmed — `422` on missing/empty `decision_note`,
  `409` if not `published`. Test: `staff` caller → `403`.

### INV-campaign-11: Auto-unpublish from organization re-verification (received — `organization` domain)

- **Statement**: See **INV-organization-09** in
  `docs/spec/organization/invariants.md` — declared there in full
  (owned by `organization`, the domain whose transaction triggers
  this). This entry is a reference, not a redefinition.
- **Holds after operations (this domain's responsibility)**: confirm,
  from `campaign`'s own model/migration perspective, that there is no
  window where an organization reads `pending_verification` but a
  campaign under it still reads `published` — the actual mechanism
  (a direct `UPDATE campaigns ... WHERE organization_id = :id AND
  status = 'published'` inside `organization`'s edit transaction, per
  `docs/spec/organization/features/04-organization-edit.md` step 7)
  lives in `organization`'s codebase, not `campaign`'s, but this
  domain must ensure its own `campaigns.status`/`unpublish_reason`
  columns and any `campaign`-side caching don't create a
  stale-read window.
- **Verification**: Full integration test now unblocked (both domains
  exist) — publish 2 campaigns under a `verified` organization, edit a
  legal/identity field on that organization (confirm=true), assert
  both campaigns flip to `unpublished` with `unpublish_reason =
  'organization_re_verification'`, in the same transaction as the
  organization's status change. This closes the testing gap flagged
  as deferred in `organization/invariants.md`'s INV-organization-09.

### INV-campaign-12: Auto-unpublished/rejected campaigns are not auto-republished

- **Statement**: After an auto-unpublish (INV-campaign-11) or a
  curation rejection (INV-campaign-07), no system process
  automatically republishes or resubmits the campaign — the Owner must
  take an explicit action (Republish for unpublished; edit + resubmit
  for rejected) via the normal flows.
- **Holds after operations**: N/A (this is an absence-of-behavior
  invariant — verified by confirming no scheduler/trigger exists for
  either case).
- **Verification**: Test — after auto-unpublish, wait/advance time,
  assert `status` remains `unpublished` with no automatic transition.

### INV-campaign-13: Three independent close triggers, idempotent via a shared status guard

- **Statement**: A `published` campaign closes (`status = 'closed'`)
  via exactly one of three independent triggers — `max_amount`
  reached (same transaction as the donation that crosses the
  threshold), `deadline` reached (periodic scheduler), or Admin
  force-close (`POST /campaigns/{campaignId}/force-close`, requires
  `decision_note`, records `closed_by`). All three are guarded by
  `WHERE status = 'published'` — whichever fires first wins, the
  others become no-ops, never errors, regardless of near-simultaneous
  timing.
- **Holds after operations**: the successful-donation transaction (see
  `donation` domain, forward reference), the deadline scheduler, and
  `force-close`.
- **Verification**: Confirmed — force-close's `409` explicitly notes
  "e.g. already closed by another trigger," and
  `kencleng-phase2-detail.md` Fitur 3 states all three triggers share
  this guard explicitly. Test: simulate a donation crossing
  `max_amount` and a concurrent force-close request; assert exactly
  one `closed_reason` is recorded, the other request gets `409` with
  no error/crash. Same for deadline-vs-force-close timing.
- **Cross-domain note**: the `max_amount`-reached trigger fires from
  within `donation` domain's donation-success transaction (not yet
  spec'd in `docs/spec/`) — when `donation/invariants.md` is written,
  it must reference this entry for the closure side effect, rather
  than redefining it.

### INV-campaign-14: Campaign visibility — public for `published`, restricted otherwise

- **Statement**: `GET /campaigns/{campaignId}` (the composite detail
  endpoint) is publicly readable (no auth) when `status = 'published'`.
  For any other status, only the owning organization's representatives
  (any `level`), the assigned/historical Kurator, or Admin may view it
  — everyone else gets `403` ("Campaign is not published and requester
  lacks visibility," confirmed explicit — note this is a `403`, not
  the anti-enumeration `404` pattern used elsewhere in this project;
  worth confirming that's intentional, since it *does* confirm a
  non-public campaign's existence to an unauthorized prober, just not
  its content).
- **`[RESOLVED — 2026-08-20]`**: `GET
  /campaigns/{campaignId}/attachments` must adopt the **same** gating
  as the detail endpoint — public when the parent campaign's `status
  = 'published'`, otherwise restricted to representatives/Kurator/
  Admin (`403` for anyone else). This closes the inconsistency flagged
  above; `api/openapi/campaign.yaml`'s current `security: []` on this
  endpoint needs updating to match at implementation time.
- **Holds after operations**: `GET /campaigns/{campaignId}`, `GET
  /campaigns/{campaignId}/attachments` (both, as of this decision).
- **Verification**: Test — `published` campaign: attachments listable
  without auth. Non-`published` campaign: representative/Kurator/Admin
  succeed, unrelated/unauthenticated caller gets `403` — same test
  shape as the detail endpoint, now applied to both.

### INV-campaign-15: `campaign_logs` is append-only

- **Statement**: No row in `campaign_logs`, once inserted, may ever be
  updated or deleted.
- **Holds after operations**: curation decisions, manual unpublish,
  force-close, (auto-unpublish, per INV-campaign-11's reference).
- **Verification**: DB-level, same pattern as
  INV-account-11/INV-organization-11.

### INV-campaign-16: Event-to-campaign linking requires same-organization, publish-eligible campaigns

- **Statement**: Every `campaign_id` in an `Event`'s `campaign_ids`
  must (a) belong to the same organization as the creating
  representative, and (b) have `status` ∈ {`published`, `scheduled`}
  at link time. At least one `campaign_id` required.
- **Holds after operations**: `POST /events`.
- **Verification**: Confirmed — `403` if any campaign belongs to a
  different organization, `409` if any campaign's status is outside
  {`published`, `scheduled`}. Test: link a `draft`/`pending_curation`/
  `rejected` campaign → `409`. Link a campaign from a different org →
  `403`. Note: unlike most `403`s in this project, this one *is*
  content-revealing in a narrow sense (confirms the campaign belongs to
  another org) — low severity, since the caller already knows the
  `campaignId` they tried to link.
- **No re-check after linking**: if a linked campaign's status later
  changes (e.g. closes), the existing `campaign_events` relation is
  **not** retracted — confirmed by absence of any such mechanism in
  the spec; the Event simply continues referencing a now-closed
  campaign. Not flagged as a gap — Events are explicitly
  "non-sensitive, no financial data" (`kencleng-phase1-detail.md`
  Fitur 6), so a stale link is a cosmetic concern, not a
  correctness/security one.

## State machines

### `campaigns.status`

```
(none) -> draft -> pending_curation -> approved -> scheduled -> published -> closed
                         |                  |          |            |
                         v                  |          v            v
                     rejected --(resubmit)--+     published    unpublished
                         |                              ^            |
                         +----------> draft              \___________/
                                                          (republish)
```

In prose: `draft → pending_curation → {approved | rejected}`.
`rejected → draft` (revise) `→ pending_curation` (resubmit, new
assignment cycle, old kept as history). `approved → {published |
scheduled}`. `scheduled → published` (time-triggered) or stays
`scheduled` (reschedule, same status). `published → {closed
(terminal, 3 triggers) | unpublished}`. `unpublished → {scheduled |
published}` (republish, same as first-publish contract).

### `campaign_curation_assignments.decision`

```
(none) -> pending -> approved / rejected
```

Same shape as `organization_curation_assignments` — one-way per row,
resubmission creates a new row.

## Reference for `donation` domain

When `docs/spec/donation/invariants.md` is written, it must reference
**INV-campaign-13** for the `max_amount`-reached closure trigger that
fires from within the donation-success transaction.

## References

- Related ERD: `docs/project/kencleng-erd.md` §3 (`campaigns`,
  `campaign_curation_assignments`, `campaign_attachments`,
  `campaign_logs`, `events`, `campaign_events`)
- Related business process: `docs/project/kencleng-phase1-detail.md`
  Fitur 3–6, `docs/project/kencleng-phase2-detail.md` Fitur 3
- Related invariants: `docs/spec/organization/invariants.md` —
  INV-organization-09, INV-organization-13; `docs/spec/account/invariants.md`
  — pattern reference for append-only logs
- **Actual API (ground truth)**: `api/openapi/campaign.yaml`
