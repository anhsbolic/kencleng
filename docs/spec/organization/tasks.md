# Domain Tasks — organization

> File: `docs/spec/organization/tasks.md`
> Status: draft — reconciled against `api/openapi/organization.yaml` 2026-08-20
> Last updated: 2026-08-20

## Reconciliation note (2026-08-20)

Endpoint paths/methods below now match `api/openapi/organization.yaml`
exactly (previously proposed/speculative). Key shape changes from the
2026-08-19 version: attachment access split into list + signed-URL
download (by `{type}`, not `{attachmentId}`); curation queue/mine/
assign/decision are now dedicated endpoints under `/organizations/...`
rather than generic CRUD on `curation-assignments`; edit and
attachment-replace both gated by `confirm=true`. Task numbering is
unchanged (still 10 tasks) since the underlying units of work didn't
change, just their endpoint shape.

**Status update, 2026-08-20**: all 5 tracked `[NEEDS DECISION]`/`[gap]`
items resolved (see `invariants.md`'s reconciliation note). Task 02
and Task 04 are no longer blocking. One separate, lower-priority item
remains open: 5-organization limit re-check on promote-to-owner
(Task 07).

| # | Task | Endpoint / surface | Depends on | Related invariants |
|---|---|---|---|---|
| 01 | Organization registration | `POST /organizations` | — | INV-organization-01, 02, 03, 04 |
| 02 | Organization detail view | `GET /organizations/{organizationId}` | 01 | INV-organization-10 |
| 03 | Organization list view | `GET /organizations`, `GET /organizations/curation-queue` | 01 | INV-organization-10 |
| 04 | Organization edit | `PATCH /organizations/{organizationId}` | 01, 02 | INV-organization-01, 08, 09 |
| 05 | Legal document attachments | `GET .../attachments`, `PUT .../attachments/{type}`, `GET .../attachments/{type}` | 01, 02 | INV-organization-08, 10, 11 |
| 06 | Representative invite | `POST .../representatives` | 01, 02 | INV-organization-05, 12 |
| 07 | Representative promote/demote | `PATCH .../representatives/{representativeId}` | 06 | INV-organization-02, 04 `[open — 5-org limit only]`, 05 |
| 08 | Representative remove | `DELETE .../representatives/{representativeId}` | 06 | INV-organization-02 |
| 09 | Curation assignment | `POST .../curation/assign` | 01 | INV-organization-06, 07 |
| 10 | Curation decision | `POST .../curation/decision`, `GET /organizations/curation-assignments/mine` | 09 | INV-organization-07 |

## Task 01 — Organization registration

**What**: `POST /organizations` (`multipart/form-data`) — creates
`organizations` (`status = pending_verification`) and the registering
user's owner `organization_representatives` row atomically. Requires
`name`, `npwp`, `akta_notaris`, `sk_kemenkumham`; `izin_pub` optional.

**KPI / metrics**: unchanged from 2026-08-19 — 0 orphaned organization
rows, 0 registrations exceeding the 5-org cap (sequential and
concurrent), 0 duplicate-NPWP registrations, Admin-role caller always
rejected.

**New, from reconciliation**: confirm file-type/size validation
(`422`, "invalid file type or exceeds 5 MB max" — documented on the
attachment-replace endpoint, Task 05) applies identically at
registration time for the initial `akta_notaris`/`sk_kemenkumham`/
`izin_pub` uploads.

## Task 02 — Organization detail view

**What**: `GET /organizations/{organizationId}` — broadly readable
(`[RESOLVED — 2026-08-20]`): any authenticated user may fetch any
organization's detail. `npwp` is server-side gated (present only for
owner-level representatives, Admin, or assigned/historical Kurator —
omitted entirely otherwise, not masked).

**KPI / metrics**:
- Any authenticated user can fetch any organization's detail (no
  `403`/`404` based on relationship).
- `npwp` present in the response only for owner/Admin/Kurator; absent
  (not masked, not present) for `staff` and for unrelated requesters.
- Confirm attachments/representative lists remain separately gated
  (Tasks 05/06's own access rules) — not exposed via this endpoint.

## Task 03 — Organization list view

**What**: `GET /organizations` (scope: "mine" — organizations where
the caller has any representative row, confirmed `# INFERRED` in the
actual spec, not from a phase doc) and `GET
/organizations/curation-queue` (Admin-only, `pending_verification` +
no active assignment, confirmed explicit).

**Reconciliation note**: these are two separate endpoints in the
actual API, not one role-aware endpoint with a `status` filter as
originally drafted. Kurator's own queue is `GET
/organizations/curation-assignments/mine` (Task 10) — not this
endpoint. There is **no** documented "Admin sees all organizations
regardless of status" general listing — only the curation-queue's
narrower scope (`pending_verification` + unassigned). Confirm whether
Admin needs a broader browse capability beyond the curation queue, or
whether the queue is genuinely the only Admin-facing list view in v1.

**KPI / metrics**: `GET /organizations` returns only the caller's own
organizations (0 leakage). `GET /organizations/curation-queue` returns
only `pending_verification` organizations with no active assignment,
Admin-only (`403` otherwise, confirmed explicit).

## Task 04 — Organization edit

**What**: `PATCH /organizations/{organizationId}` (`application/json`,
not multipart — attachment replacement is a separate endpoint, Task
05). Access split by field class (`[RESOLVED — 2026-08-20]`): `staff`
may edit `description`/`contact` only; `owner`-only for `name`/`npwp`.
Legal/identity change (`name`/`npwp`) requires `confirm: true` (`409
confirmation-required` otherwise) and flips `status` back to
`pending_verification`.

**KPI / metrics**:
- `staff` submitting `name`/`npwp` → `403`, regardless of `confirm`.
- `staff` submitting `description`/`contact`-only → succeeds.
- 100% of `owner` `name`/`npwp` changes without `confirm: true` are
  rejected, no partial apply.
- `confirm: true` + legal/identity change on a `verified` org flips
  `status`, 100% of the time.
- NPWP uniqueness re-checked on edit.

## Task 05 — Legal document attachments

**What**: three endpoints —
- `GET .../attachments` — metadata list, owner/Admin/Kurator only.
- `PUT .../attachments/{type}` — upload/replace; `akta_notaris`/
  `sk_kemenkumham` require `confirm=true` and trigger re-curation,
  `izin_pub` does not (`[RESOLVED — 2026-08-20]`, follows the
  implementation); owner-only.
- `GET .../attachments/{type}` — returns a 5-minute signed download
  URL, owner/Admin/Kurator only.

**Reconciliation note**: this replaces the 2026-08-19 draft's
attachment-by-`{attachmentId}` + direct-download design. `{type}` is
the path key (one current attachment per type, matching
`AttachmentType` enum: `akta_notaris`/`sk_kemenkumham`/`izin_pub`) —
confirm this means each type is a **slot** (replacing overwrites,
doesn't version) rather than a history of uploads per type.

**KPI / metrics**:
- 0 successful attachment list/upload/download by `staff`.
- File type/size validation (`422`) enforced on upload.
- `akta_notaris`/`sk_kemenkumham` replacement without `confirm=true`
  rejected (`409`); `izin_pub` replacement never requires it.
- Signed URL is valid for exactly 5 minutes, confirmed by test
  (expired URL rejected by the storage layer).

## Task 06 — Representative invite

**What**: `POST .../representatives` — owner-only, direct-add,
`level = staff`. Confirmed: nonexistent-or-unverified email → single
`404 user-not-found`; already-representative → `409`.

**`[RESOLVED — 2026-08-20]`**: `409 admin-cannot-be-representative`
for a target holding `role = 'admin'` (INV-organization-05).

**KPI / metrics**:
- 0 successful invites for nonexistent/unverified emails (`404`,
  single collapsed message, confirmed).
- 0 successful invites for already-representative emails (`409`).
- 0 successful invites for `role = 'admin'` targets (`409
  admin-cannot-be-representative`).
- 0 successful invites by `staff`.

## Task 07 — Representative promote/demote

**What**: `PATCH .../representatives/{representativeId}` —
owner-only, toggles `level`. Demote: `409 last-owner` if it would
leave 0 owners, confirmed explicit, "including for self-demote."

**`[RESOLVED — 2026-08-20]`**: promoting a `role = 'admin'` user to
`owner` → `409 admin-cannot-be-representative` (same shape as Task
06).

**`[open]` — still unresolved, separate from the 5 decided items**:
no documented re-check of the 5-organization-owner cap on promote
(INV-organization-04's own open note). Needs a decision before this
task ships.

**KPI / metrics**:
- 0 successful demotes leaving 0 owners, sequential and concurrent.
- 0 successful promote-to-owner for `role = 'admin'` targets (`409
  admin-cannot-be-representative`).
- 0 successful promote-to-owner exceeding the 5-org cap (**pending**
  the still-open decision above).

## Task 08 — Representative remove

**What**: `DELETE .../representatives/{representativeId}` —
owner-only, hard delete, same `409 last-owner` guard as demote.

**KPI / metrics**: unchanged from 2026-08-19 — 0 successful removals
leaving 0 owners, sequential and concurrent; 0 successful removals by
`staff`.

## Task 09 — Curation assignment

**What**: `POST .../curation/assign` — Admin-only. Confirmed
explicit: `409 curator-conflict-of-interest`,
`409 assignment-already-active`.

**KPI / metrics**: unchanged from 2026-08-19 — 0 successful
assignments with conflict of interest; 0 successful assignments
creating a second pending assignment.

## Task 10 — Curation decision

**What**: `POST .../curation/decision` (keyed by `organizationId`, not
an assignment id — server resolves the current pending assignment) +
`GET /organizations/curation-assignments/mine` (Kurator's own queue,
filterable by `decision`). Confirmed explicit: `403` if caller isn't
the assigned Kurator, `409` if no pending assignment exists, `422` if
`decision = rejected` without `decision_note`.

**Reconciliation note**: this replaces the 2026-08-19 draft's
`PATCH .../curation-assignments/{assignmentId}` design. Add a test for
the new TOCTOU consideration flagged in `threat-model.md`: a decision
submitted after the org was reassigned to a different Kurator must be
rejected against the *current* assignment, not silently applied to a
stale one.

**KPI / metrics**:
- 0 successful decisions by a Kurator not matching the *current*
  pending assignment's `kurator_id` (including the reassignment race
  case above).
- Reject without `decision_note` → `422`.
- `GET .../curation-assignments/mine` correctly scopes to the caller,
  filterable by `decision`.

## References

- Related domain invariants: `docs/spec/organization/invariants.md`
  (reconciled 2026-08-20 — three `[NEEDS DECISION]` items shared
  across invariants/threat-model/tasks)
- Related threat model: `docs/spec/organization/threat-model.md`
  (reconciled 2026-08-20)
- **Actual API (ground truth)**: `api/openapi/organization.yaml`
- Feature specs (need reconciliation next, same open items apply):
  `docs/spec/organization/features/`
