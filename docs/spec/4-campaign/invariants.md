# Domain Invariant — campaign

> File: `docs/spec/4-campaign/invariants.md`
> Status: Slice-1 public-boundary reconciliation active; remaining historical invariants are deferred evidence
> Last updated: 2026-10-04
> WU-S2-006 source amendment accepted by Anhar Solehudin on 2026-10-02 after independent Review; acceptance covers this source amendment only.
> Campaign/Donation funding-unavailable admission amendment accepted by Anhar Solehudin on 2026-10-04 as current Slice-2 Product/Donation authority; see INV-campaign-13.

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

## Slice-1 reconciliation note

Historical mixed-auth detail and anonymous attachment-list behavior is
superseded at the public boundary by INV-campaign-14. Historical lifecycle,
listing, upload, and privileged-read artifacts remain deferred evidence; they
do not define the active public contract.

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
  `sosial`/`lainnya`). For Slice 2, `max_donation_amount` is a per-Campaign
  whole-IDR cap in the inclusive range Rp5.000–Rp1.000.000.000. Omission on
  creation selects Rp1.000.000.000; an explicitly supplied value is valid
  only while the Campaign is a draft. Omission on PATCH preserves the
  existing value, including for older clients updating unrelated fields.
  The effective value is frozen when the Campaign is published. Every
  existing Campaign row must be backfilled to Rp1.000.000.000 before
  cap-dependent behavior is enabled. API requests and responses represent
  the cap as a closed `{amount, currency_code}` object with a major-unit
  decimal-string `amount` and required `currency_code: IDR`. This feature
  cap is distinct from the fundraising `max_amount` threshold and does not
  set a project-wide money range or precision policy.
- **Holds after operations**: `POST
  /organizations/{organizationId}/campaigns`, `PATCH
  /campaigns/{campaignId}` (while `draft`).
- **Verification**: Test — each rule individually (zero/negative
  `target_amount`, `max_amount < target_amount`, past `deadline`,
  missing `category`) → `422`. Downstream contract/runtime evidence covers
  the cap's inclusive boundaries, create default, PATCH omission
  preservation with a non-default stored value, existing-row backfill,
  and rejection of supplied cap changes after publication.

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

### INV-campaign-13: Campaign close has one stable winning trigger; Donation ordering follows D1 and capacity reservation

- **Statement**: A `published` Campaign closes (`status = 'closed'`)
  through its applicable fundraising `max_amount` threshold, finite funding
  capacity, deadline, or Admin force-close trigger. The first close
  transition wins and its `closed_reason` remains stable. `max_amount` is a
  distinct, overshootable fundraising threshold; it is not the finite
  representability ceiling. Slice-2 capacity admission accounts for
  collected settled Funding plus the full amounts of accepted-pending
  Donations. The Campaign funding ceiling is IDR
  `99,999,999,999,999,999`, derived from the current `NUMERIC(19,2)` Campaign
  funding representation; no admitted Donation may make settled Funding
  plus accepted-pending reservations exceed it. After accounting for settled
  Funding and accepted-pending reservations, the Campaign closes with the
  distinct `funding_capacity_reached` reason when remaining capacity is less
  than the minimum valid Donation of Rp5.000, including exact exhaustion.
  Admission still rejects any full Donation amount that does not fit; this
  does not change the separate, overshootable `max_amount` threshold. If
  close wins before an otherwise new
  submission is admitted, reject that submission. An already accepted
  Donation remains settleable in full after any close. A later successful
  settlement and full funding reflection commit atomically and exactly
  once, but cannot reopen the Campaign or replace the winning close reason;
  funding can exceed `max_amount` while remaining within the capacity
  ceiling. A failed accepted-pending Donation contributes no Funding and
  releases its reservation, but does not reopen a capacity-closed Campaign
  or replace its winning reason. This invariant selects no transaction,
  locking, or isolation mechanism and does not change broader closure
  behavior. New Donation admission also requires authoritative settled
  Funding to be available. If that Funding value is unavailable, fail closed:
  admit no new Donation, do not interpret an absent value as zero, and resume
  admission only after the authoritative Funding value is restored.
- **Holds after operations**: Donation submission/settlement ordering
  at the threshold and capacity boundaries, reservation admission/release,
  the deadline scheduler, and `force-close`.
- **Verification**: Runtime Testing must demonstrate the D1
  submit-vs-close orderings; reservation accounting across settled and
  accepted-pending amounts, including concurrent admissions at the ceiling;
  accepted-pending settlement after close; full exact-once funding;
  capacity-close boundaries where residual capacity is Rp4.999 (close) versus
  Rp5.000 (remain open), including exact exhaustion; rejection of a Donation
  that does not fit while a smaller valid amount does; capacity-close reason
  stability after a pending failure; no reopen after reservation release; and
  `max_amount` threshold overshoot, alongside
  applicable deadline/force-close behavior. The donation/close verification
  is not satisfied by this spec reference or by choosing a locking
  mechanism here.
- **Cross-domain note**: Campaign owns lifecycle and `closed_reason`;
  Donation owns accepted-submission and settlement contribution.
  `docs/spec/5-donation/invariants.md#inv-donation-02` and
  `#inv-donation-08` reference this invariant for the D1 boundary.
  Historical wording that made later settlement conditional on the
  Campaign still being `published` is superseded for an already
  accepted Donation.

### INV-campaign-14: Public Campaign eligibility, projection, and media visibility

- **Statement**: `GET /campaigns/{campaignId}` and `GET
  /campaigns/{campaignId}/media/{mediaId}/content` are public-only
  operations. In Slice 1, only an internally `published` fundraising
  Campaign is publicly eligible. Draft, pending, approved-but-not-published,
  scheduled, rejected, unpublished/retracted, and historical closed states
  are non-public until later Slice reconciliation explicitly changes this.
  Slice 2 continues to return the indistinguishable `404` for closed
  Campaigns; the Slice-3 public-result behavior below is a handoff, not active
  Slice-2 visibility.
- **Anti-enumeration**: absent, malformed/non-resolvable, and non-public
  Campaigns return the same `PublicCampaignNotFound` `404` Problem Details
  response. The media-content operation returns that same response for an
  absent/non-member media ID. Presence, absence, validity, or role content
  of optional `Authorization` cannot vary visibility, status, fields, or
  public error shape.
- **Projection**: the detail response is the standalone closed
  `PublicCampaignDetail` allowlist, never an internal `Campaign` or
  authenticated `Organization` serialization. In Slice 2 it contains
  exactly `id`, `title`, required `max_donation_amount`, `purpose`, `story`,
  `steward`, `lifecycle`, `funding`, `media`, and `donation_action`, with
  closed nested exact projections. `max_donation_amount` is the effective
  per-Campaign whole-IDR cap, disclosed before amount entry. The projection
  excludes raw status/reasons, contact/legal data, actor IDs, audit
  timestamps, `donor_count`, and operational metadata. Organizer text is
  plain text marked `source: organizer`; it is not verification evidence.
- **Funding/action truth**: funding uses tagged availability with IDR decimal
  strings, backend-authored uncapped progress/relationship, and distinct
  zero/unavailable semantics. In Slice 2, required `donation_action` is
  `{availability: available}` when the backend Donation submission-eligibility
  predicate passes at the GET snapshot, or
  `{availability: unavailable, reason: campaign_not_eligible}` only when the
  detail remains public and that same predicate fails. The generic reason
  exposes no internal lifecycle/closure detail. GET is an affordance snapshot;
  Donation POST independently rechecks current eligibility and validates
  against the current cap and remaining capacity. The GET cap does not
  authorize or reserve a Donation. No navigation target is included.
- **Slice-3 handoff**: when Slice 3 is reconciled, a public closed-Campaign
  result retains Campaign identity, removes `donation_action`, and does not
  present Funding as final while accepted Donations remain pending. This
  does not activate closed details in Slice 2 or settle Slice-3 projection
  fields.
- **Media**: metadata follows the same parent eligibility predicate. Public
  bytes remain in private storage and are delivered only after parent and
  member recheck at the origin; a storage/object dependency failure for an
  eligible member is `PublicCampaignUnavailable` `503`, never false absence.
  No direct object-storage URL or redirect is public.
- **Withdrawal/cache**: success and public `404`/`503` responses specify
  `Cache-Control: private, no-store`. Retraction prevents new origin fetches
  from a previously known controlled URL; already downloaded client-held
  bytes remain outside this guarantee.
- **Holds after operations**: `GET /campaigns/{campaignId}` and `GET
  /campaigns/{campaignId}/media/{mediaId}/content`.
- **Verification**: downstream implementation/tests must prove allowlist
  mapping, absent/non-public/auth parity (including timing), private-storage
  retraction, no-store behavior, storage failure distinction, exact decimal
  progress, safe plain-text rendering, and public cap disclosure with
  independent POST enforcement.

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

`docs/spec/5-donation/invariants.md#inv-donation-02` and
`#inv-donation-08` reference **INV-campaign-13** for the narrow D1
eligibility, threshold, and accepted-settlement boundary. Campaign
owns lifecycle and the winning close reason; see the invariant above
for the current behavior.

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
