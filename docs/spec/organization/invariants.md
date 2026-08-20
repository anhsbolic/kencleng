# Domain Invariant — organization

> File: `docs/spec/organization/invariants.md`
> Status: draft — reconciled against `api/openapi/organization.yaml` 2026-08-20
> Last updated: 2026-08-20

## Reconciliation note (2026-08-20, decisions finalized 2026-08-20)

This revision reconciles the 2026-08-19 draft against the actual API
design in `api/openapi/organization.yaml`. Five items were flagged as
`[NEEDS DECISION]`/`[gap]` and have now all been **resolved**:

1. **`izin_pub` re-curation `[RESOLVED]`**: follows the implementation
   — replacing `izin_pub` does **not** trigger re-curation, unlike
   `akta_notaris`/`sk_kemenkumham`. `kencleng-phase1-detail.md` Fitur 1
   should be updated to match (currently still lists Izin PUB as
   legal/identity-class — that phase doc is now the stale one).
2. **NPWP visibility `[RESOLVED]`**: server-side gating restored —
   `staff` receives no `npwp` field at all in `GET
   /organizations/{organizationId}` responses (not just frontend
   masking). See revised INV-organization-10 below.
3. **`GET /organizations/{organizationId}` access `[RESOLVED]`**:
   broadly readable — any authenticated user may fetch any
   organization's detail. Safe now that `npwp` is server-side gated
   per decision 2.
4. **Admin-exclusion error shape `[RESOLVED]`**: `409`, `type:
   admin-cannot-be-representative`, used identically on invite
   (Task 06) and promote-to-owner (Task 07).
5. **Who can `PATCH` an organization `[RESOLVED]`**: split by field
   class — `staff` may edit operational fields (`description`,
   `contact`); `name`/`npwp` (legal/identity, `confirm`-gated,
   re-curation-triggering) remain owner-only. See revised
   INV-organization-08 below.

## Domain summary

`organization` owns organization registration, legal-document
attachments, representative management (invite/promote/demote/
remove), and the organization curation (verification) workflow.
Covers `organizations`, `organization_representatives`,
`organization_curation_assignments`, `organization_attachments`, and
`organization_logs`. Does **not** cover campaigns, campaign curation,
or events — those belong to the `campaign` domain, even though two
organization-domain facts have documented, cross-domain effects on
`campaign`: re-verification auto-unpublishing live campaigns
(INV-organization-09), and `has_overdue_report` blocking new campaign
creation (INV-organization-13, new — see below).

## Invariants

### INV-organization-01: NPWP is globally unique

- **Statement**: `npwp_hash` is unique across all of
  `organizations` — one NPWP can only ever be registered under one
  organization (organization *names* may repeat).
- **Holds after operations**: organization registration (`POST
  /organizations`), and any later edit to `npwp` (`PATCH
  /organizations/{organizationId}`, a legal/identity field —
  INV-organization-08).
- **Verification**: DB-level — `ux_organizations_npwp_hash`. Confirmed
  in the actual API: registration's `409 npwp_taken` example, edit's
  `409` "NPWP already taken" case.
- **Related validation (not itself an invariant)**: `npwp` format is a
  regex (`^\d{2}\.\d{3}\.\d{3}\.\d-\d{3}\.\d{3}$`, confirmed verbatim
  in `OrganizationCreateRequest`/`OrganizationUpdateRequest`) against
  the plaintext value, before AES-GCM encryption — no external DJP
  legitimacy check.

### INV-organization-02: Every organization retains at least one owner-level representative at all times

- **Statement**: For any `organization_id`,
  `organization_representatives` must contain at least one row with
  `level = 'owner'`, from creation onward — no window with zero.
- **Holds after operations**: creation (satisfied automatically —
  INV-organization-03), demote (`PATCH .../representatives/{id}`),
  remove (`DELETE .../representatives/{id}`). Confirmed in the actual
  API: both endpoints document a `409 last-owner` case, explicitly
  "including for self-demote"/"self-remove."
- **Verification**: Test — attempt to demote/remove the last remaining
  owner, assert `409`. Concurrency test: two near-simultaneous
  demote/remove requests both targeting owners of a 2-owner
  organization — at most one succeeds.

### INV-organization-03: Organization creation and its first owner representative are atomic

- **Statement**: The insert into `organizations` and the insert into
  `organization_representatives` (`level = 'owner'`) happen in a
  single DB transaction. No committed state without an owner.
- **Holds after operations**: `POST /organizations` only.
- **Verification**: Test — fault-inject a failure between the two
  inserts, assert full rollback.

### INV-organization-04: Org-per-user limit (max 5 as owner) is enforced atomically at registration

- **Statement**: A user may not hold `level = 'owner'` in more than 5
  distinct organizations. Checked and inserted atomically within one
  transaction (`SELECT ... FOR UPDATE` or equivalent) at registration
  time.
- **Holds after operations**: `POST /organizations`. Confirmed in the
  actual API: `409 organization_limit` example on registration.
- **`[NEEDS DECISION]` — still open**: whether this limit is also
  re-checked on promote-to-owner (`PATCH .../representatives/{id}`).
  `organization.yaml`'s promote endpoint documents no such check —
  this remains genuinely unresolved in the actual spec, not just in
  this invariants doc. If left unenforced, a user could exceed 5 owned
  organizations via promotion rather than registration.
- **Verification**: Test — register 5, attempt a 6th (rejected).
  Concurrency test at 4-owned.

### INV-organization-05: Admin role is mutually exclusive with being a representative

- **Statement**: See **INV-account-10** in
  `docs/spec/account/invariants.md` — declared there in full.
- **Holds after operations (this domain's responsibility)**:
  representative invite (`POST .../representatives`) and
  promote-to-`owner` (`PATCH .../representatives/{id}`) must reject a
  target user holding `role = 'admin'`.
- **`[RESOLVED — 2026-08-20]`**: was an undocumented enforcement gap
  in `organization.yaml` (neither endpoint listed an error case for
  this). Now explicit: `409`, `type:
  admin-cannot-be-representative`, used identically on both invite
  and promote-to-`owner`.
- **Verification**: Test — invite/promote a user holding `role =
  'admin'`, assert `409 admin-cannot-be-representative` on both
  endpoints.

### INV-organization-06: Kurator conflict-of-interest recusal

- **Statement**: A user may not be assigned as Kurator for an
  organization where they're also a representative (any `level`).
- **Holds after operations**: `POST
  /organizations/{organizationId}/curation/assign`.
- **Verification**: Confirmed in the actual API — `409
  curator-conflict-of-interest` example. Test: assign a
  representative as their own org's Kurator, assert rejection.

### INV-organization-07: Only one active curation assignment per organization

- **Statement**: `organization_curation_assignments` may have at most
  one row with `decision = 'pending'` per `organization_id`.
- **Holds after operations**: curation assignment creation.
- **Verification**: DB-level (`ux_org_curation_one_pending`).
  Confirmed in the actual API — `409 assignment-already-active`
  example.
- **Endpoint note**: the actual decision endpoint (`POST
  /organizations/{organizationId}/curation/decision`) does **not**
  take an assignment id in the path — it resolves "the current
  pending assignment for this organization" server-side, then checks
  the caller against that row's `kurator_id`. Functionally equivalent
  to what was originally drafted as a `PATCH .../{assignmentId}`
  endpoint, just resolved by organization id instead of assignment id
  — same invariant, different endpoint shape. `409` if no pending
  assignment exists.

### INV-organization-08: Legal/identity field edits after `verified` are owner-only, require explicit confirmation, and deterministically trigger re-curation

- **Statement**: Editing `name` or `npwp`, or replacing the
  `akta_notaris`/`sk_kemenkumham` attachment, is **owner-only**
  (`[RESOLVED — 2026-08-20]`, split-by-field-class), and **requires**
  `confirm: true` in the same request — absent that, the request is
  rejected outright (`409 confirmation-required`), no partial apply.
  When `confirm: true` is present and the organization's `status =
  'verified'`, `status` is set back to `pending_verification`.
  Editing `description`/`contact` is representative-inclusive
  (`owner` or `staff`), never requires confirmation, and never
  changes `status`. Replacing `izin_pub` is also representative-
  inclusive, never requires confirmation, and never changes `status`
  (`[RESOLVED — 2026-08-20]`, follows the implementation —
  `kencleng-phase1-detail.md` Fitur 1's listing of Izin PUB as
  legal/identity-class is now stale and should be updated to match).
- **Authorization summary**: `staff` → `description`/`contact`/
  `izin_pub` only, `403` if the request touches `name`/`npwp`/
  `akta_notaris`/`sk_kemenkumham`. `owner` → all fields.
- **Holds after operations**: `PATCH /organizations/{organizationId}`,
  `PUT /organizations/{organizationId}/attachments/{type}`.
- **Verification**: Test — `staff` submits a `name`/`npwp` change (or
  `akta_notaris`/`sk_kemenkumham` replacement) → `403`, regardless of
  `confirm`. `staff` submits `description`/`contact`-only change, or
  `izin_pub` replacement → succeeds. `owner` submits a `name` change
  without `confirm` → `409`, no change applied. `owner` submits with
  `confirm: true` on a `verified` org → `status` flips. `izin_pub`
  replacement (any representative) never requires confirmation and
  never changes `status`.

### INV-organization-09: Re-verification atomically auto-unpublishes the organization's live campaigns (forward reference — `campaign` domain)

- **Statement**: When INV-organization-08 fires a `verified →
  pending_verification` transition, every campaign belonging to that
  organization with `status = 'published'` must, in the same
  transaction, transition to `status = 'unpublished'` with
  `unpublish_reason = 'organization_re_verification'`.
- **Cross-domain note**: unchanged from the 2026-08-19 draft — declared
  here per the same precedent as INV-account-10, `campaign` domain
  must reference (not redefine) this when built.
- **Verification**: Unchanged — full integration test deferred until
  `campaign` domain exists; this domain's responsibility now is to
  confirm the same-transaction guarantee from the organization side.

### INV-organization-10: NPWP is server-side gated by representative level; `GET /organizations/{organizationId}` is broadly readable; legal-document *files* remain a server-side, owner-only authorization boundary

- **Statement — `[RESOLVED — 2026-08-20]`**: `GET
  /organizations/{organizationId}` is broadly readable — any
  authenticated user may fetch any organization's detail (no
  representative/Admin/Kurator relationship required). The `npwp`
  field, however, **is** server-side gated: it is present in the
  response only when the requester is an `owner`-level representative
  of that organization, Admin, or the assigned/historical Kurator —
  omitted entirely (not masked-with-a-flag) for `staff` and for any
  requester with no qualifying relationship. This restores the
  2026-08-19 draft's server-side-omission approach for `npwp`
  specifically, while confirming broad readability for the rest of
  the record. Frontend `MaskedField` masking still applies on top of
  this for whoever *does* receive the field (including the
  organization's own Owner), per
  `kencleng-actors-entities.md`'s PII Handling Note — this invariant
  governs the API boundary, not the display layer.
- **What *is* still a server-side boundary (unchanged)**:
  legal-document **attachments** (`GET .../attachments`, `GET
  .../attachments/{type}` for the signed download URL) are
  restricted to Owner-level representatives, Admin, or Kurator —
  `staff` is rejected (`403`).
- **Holds after operations**: `GET /organizations/{organizationId}`
  (npwp field-level gating), `GET
  /organizations/{organizationId}/attachments`, `GET
  .../attachments/{type}` (full-endpoint owner/Admin/Kurator gating).
- **Verification**: Test — `staff` requests org detail, assert `npwp`
  is absent from the payload while other fields are present. Owner,
  Admin, and assigned/historical Kurator all receive `npwp`. Any
  authenticated user (including one with no relationship to the org)
  can fetch the record at all, minus `npwp` unless they qualify.
  `staff` requests attachment list/download, assert `403` (unchanged).

### INV-organization-11: `organization_logs` is append-only

- **Statement**: No row in `organization_logs`, once inserted, may
  ever be updated or deleted — by any actor, including Admin.
- **Holds after operations**: curation decisions, `has_overdue_report`
  set/clear, representative management actions (invite/remove/
  promote/demote).
- **Verification**: DB-level — `REVOKE UPDATE, DELETE ON
  organization_logs FROM kencleng_app`. Same pattern as
  INV-account-11.

### INV-organization-12: Representative invite targets only an existing, email-verified user

- **Statement**: Inviting a representative requires the target email
  belong to an already-registered, already-email-verified user.
- **Confirmed & refined by the actual API — 2026-08-20**: both
  "no such user" and "user exists but unverified" collapse into a
  **single** `404 user-not-found` response
  ("User harus terdaftar & email terverifikasi dulu.") — matches the
  2026-08-19 accepted-risk decision on invite enumeration exactly.
  Duplicate-representative case is a separate `409`.
- **Holds after operations**: `POST
  /organizations/{organizationId}/representatives`.
- **Verification**: Test — nonexistent email → `404` with the
  documented message. Existing-but-unverified email → **same** `404`
  message (not distinguishable). Already-representative email → `409`.

### INV-organization-13: `has_overdue_report` blocks new campaign creation (new, forward/cross-domain reference — `campaign` domain)

- **Statement — new, discovered 2026-08-20 via `campaign.yaml`**: an
  organization with `has_overdue_report = true` cannot have new
  campaigns created under it. `campaign.yaml`'s campaign-creation
  error documents: `"Organisasi ini punya laporan penggunaan dana
  yang telat. Submit laporan tsb dulu sebelum membuat campaign baru."`
  Also confirmed: campaign creation additionally requires
  `Organization.status = 'verified'` (`"Organisasi harus berstatus
  verified untuk membuat campaign."`) — effectively an extension of
  INV-organization-08's status semantics into `campaign` domain's
  write path.
- **Ownership**: `organizations.has_overdue_report` is a field
  `organization` owns (per `kencleng-erd.md` §2), but this domain
  doesn't yet document *when* the flag is set or cleared — that
  appears to belong to `disbursement` domain (fund-usage-report
  deadline tracking, not yet spec'd in `docs/spec/`), making this a
  **three-domain** relationship: `disbursement` sets/clears the flag,
  `organization` owns the field, `campaign` enforces it at write time.
- **Cross-domain note**: declared here (in `organization`, the field
  owner) per the established cross-domain ownership convention. Both
  `campaign/invariants.md` (enforcement) and
  `disbursement/invariants.md` (set/clear trigger — not yet written)
  must reference this entry when those domains are built, rather than
  redefining it.
- **Verification (deferred)**: full integration testing needs both
  `disbursement` (to actually set the flag via a real overdue report)
  and `campaign` (to attempt creation and observe the block) — neither
  exists yet. This domain's responsibility now is limited to
  confirming `has_overdue_report` is readable/queryable correctly by
  other domains (it already is, per the ERD's `ix_organizations_overdue`
  partial index) — flagged as a testing gap to close later, not
  silently skipped.

## State machines

### `organizations.status`

```
(none) -> pending_verification -> verified / rejected
verified -> pending_verification   (re-curation, INV-organization-08,
                                     now gated by confirm=true)
rejected -> pending_verification   (resubmit after rejection)
```

### `organization_representatives.level`

Toggles freely between `owner`/`staff` (subject to
INV-organization-02's guard on `owner → staff`); rows are deleted
outright on removal.

### `organization_curation_assignments.decision`

```
(none) -> pending -> approved / rejected
```

One-way per row; resubmission creates a new row (old kept as history).

## Reference for `campaign` domain

When `docs/spec/campaign/invariants.md` is written, it must reference:
- **INV-organization-09** (auto-unpublish on re-verification)
- **INV-organization-13** (`has_overdue_report` blocks creation, and
  the `status = 'verified'` requirement on creation)

## Reference for `disbursement` domain

When `docs/spec/disbursement/invariants.md` is written, it must
reference **INV-organization-13** for the set/clear trigger on
`has_overdue_report`, since that domain owns the fund-usage-report
deadline logic that determines when the flag flips.

## References

- Related ERD: `docs/project/kencleng-erd.md` §2
- Related business process: `docs/project/kencleng-phase1-detail.md`
  Fitur 1, 1B, 2, 5
- Related actors/rules: `docs/project/kencleng-actors-entities.md`
  Business Rules 1–4, PII Handling Note
- Related invariants: `docs/spec/account/invariants.md` —
  INV-account-10, INV-account-11
- **Actual API (ground truth for endpoint shape as of 2026-08-20)**:
  `api/openapi/organization.yaml`, `api/openapi/campaign.yaml`
  (for INV-organization-13's discovery)
