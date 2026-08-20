# Threat Model — organization

> File: `docs/spec/organization/threat-model.md`
> Status: draft — reconciled against `api/openapi/organization.yaml` 2026-08-20
> Last updated: 2026-08-20

## Reconciliation note (2026-08-20, decisions finalized 2026-08-20)

Reconciled against the actual endpoint shapes in
`api/openapi/organization.yaml`. All five `[NEEDS DECISION]`/`[gap]`
items are now resolved — see `invariants.md`'s reconciliation note for
the full list. Summary as it affects this threat model: `GET
/organizations/{organizationId}` is broadly readable (any
authenticated user), `npwp` is server-side gated to
owner/Admin/Kurator, edit access is split by field class
(`staff` → operational fields only, `owner` → all fields including
`name`/`npwp`), and Admin-exclusion on invite/promote now has an
explicit `409 admin-cannot-be-representative` error.

## Actors & trust boundaries

Unchanged from the 2026-08-19 draft — see that section's table
(Prospective Owner, Owner, Staff, Admin, Kurator, System/auto-unpublish
side effect). One addition:

| Actor | Authenticated? | Trust boundary crossed |
|---|---|---|
| Anyone holding a valid, unexpired **signed attachment URL** | No (bearer-token-like, not a session) | Whoever holds the URL string, for up to 5 minutes — a distinct, narrower trust boundary than the API itself; see "Legal document attachments" below |

## STRIDE per operation

### Organization registration — `POST /organizations`

Unchanged from the 2026-08-19 draft's "Organization registration"
section — confirmed to match the actual API (`multipart/form-data`,
5-organization cap, Admin exclusion, NPWP uniqueness).

One addition: **file upload attack surface** — `akta_notaris`,
`sk_kemenkumham`, `izin_pub` are user-supplied binary files.
`organization.yaml` documents a `422` for "invalid file type or
exceeds 5 MB max" on the attachment-replace endpoint, but this same
constraint isn't explicitly re-stated on the registration endpoint's
error list — worth confirming the same file-type/size validation
applies at registration time too, not just on later replacement.

### Organization detail view — `GET /organizations/{organizationId}`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A beyond standard auth | `bearerAuth` | None |
| Tampering | N/A — read-only | — | — |
| Repudiation | N/A | — | — |
| Information disclosure `[RESOLVED — 2026-08-20]` | The endpoint documented only `200`/`404`, no `403` — ambiguous whether any authenticated user could fetch any organization's full detail. | **Decided: broadly readable is intentional** (organization profiles are effectively semi-public, consistent with donors needing to see organization info from campaign pages later) — safe now that `npwp` is server-side gated to owner/Admin/Kurator (INV-organization-10), so a non-related caller only receives `name`/`description`/`contact`/`status`/`has_overdue_report`. | None — accepted, contingent on `npwp` gating being correctly implemented per INV-organization-10 |
| Denial of service | N/A | — | — |
| Elevation of privilege | N/A | — | — |

### Organization edit — `PATCH /organizations/{organizationId}`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A | `bearerAuth` | None |
| Tampering | Client attempts a legal/identity change (`name`/`npwp`) without `confirm: true`, hoping the server applies it anyway | Server rejects outright (`409 confirmation-required`), no partial apply — confirmed in the actual API | None |
| Tampering | Client sends `confirm: true` on a request that only touches operational fields, attempting to force an unnecessary re-curation cycle as a griefing move against their own org | Low-severity even if possible (self-inflicted, per the original threat model's "Denial of service" note on this section) — confirm whether the server *ignores* a stray `confirm: true` when no legal/identity field actually changed, or treats it as a no-op trigger either way; worth a small explicit test either way | Low |
| Tampering — who can edit `[RESOLVED — 2026-08-20]` | Split by field class: `staff` may edit `description`/`contact`/`izin_pub`; `owner`-only for `name`/`npwp`/`akta_notaris`/`sk_kemenkumham`. | Server-side check on the specific fields present in the request, not a blanket representative-vs-owner gate | None — the generic `403` text ("not a representative") still applies for non-representatives; a *representative* `staff` submitting a legal/identity-class field gets a specific `403` for that case instead |
| Repudiation | Edit action logging — still an open item (carried from 2026-08-19): is a legal/identity edit added to `organization_logs` scope? | None newly confirmed either way | Open |
| Information disclosure | N/A beyond INV-organization-10's open item | — | — |
| Denial of service | N/A | — | — |
| Elevation of privilege | Depends on resolving "who can edit" above | — | — |

### Legal document attachments — list, upload/replace, signed-URL download

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A beyond standard auth on the list/upload endpoints | `bearerAuth` + owner/Admin/Kurator check (list), owner-only (upload) | None |
| Tampering | `staff` attempts upload/replace | `403`, "Only `owner`-level representatives may manage legal documents" — confirmed explicit in the actual API | None |
| Tampering | Client submits a legal-document replacement (`akta_notaris`/`sk_kemenkumham`) without `confirm=true` | `409`, same pattern as the edit endpoint | None |
| Repudiation | Attachment upload/replace not explicitly confirmed in `organization_logs` scope | Same open item as the edit endpoint's repudiation row | Open |
| **Information disclosure — signed URL leakage** | The download endpoint returns a **signed URL, 5-minute expiry**, rather than streaming the file directly. If that URL is captured (browser history, a proxy log, a screenshot, forwarded in a chat) within its 5-minute window, whoever holds it can access the file **without any further authentication** — the signed URL itself *is* the credential for that window. | 5-minute TTL bounds the exposure window | **Accepted, standard trade-off for signed-URL patterns** — worth confirming the URL isn't logged anywhere in plaintext (e.g. access logs, error-tracking breadcrumbs) server-side, which would extend the effective exposure window well past 5 minutes |
| Information disclosure | `staff` accesses attachment list/download | `403` on both, confirmed explicit in the actual API (INV-organization-10, the part that **is** still server-side enforced) | None |
| Denial of service | Repeated signed-URL requests to force many presigned-URL generations (cheap on the app side, but could be used to probe rate limits on the storage backend) | None specific — general rate-limiting is a project-wide concern | Low, accepted for a sandbox project |
| Elevation of privilege | N/A beyond the tampering cases above | — | — |

### Representative invite — `POST /organizations/{organizationId}/representatives`

Unchanged from the 2026-08-19 draft's accepted-risk decision on
enumeration — confirmed to match exactly (`404 user-not-found`
collapsing "doesn't exist" and "unverified" into one message, `409`
for already-representative).

**Gap resolved 2026-08-20**: the actual API documented no `409` (or
any) case for "target user holds `role = 'admin'`"
(INV-organization-05's invite direction) — this was a real
Elevation-of-privilege gap, not just a documentation gap. Now
explicit: `409 admin-cannot-be-representative`.

### Representative promote/demote/remove

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A | `bearerAuth` + owner-only | None |
| Tampering | Owner attempts to demote/remove the last remaining owner | `409 last-owner`, confirmed explicit | None |
| Elevation of privilege — promote-to-owner, Admin exclusion `[RESOLVED — 2026-08-20]` | Same gap as invite | `409 admin-cannot-be-representative`, same shape as invite | None |
| Elevation of privilege — promote-to-owner, 5-org limit | No documented re-check of the 5-organization-as-owner cap on promotion | None confirmed | **Still open** — INV-organization-04's own open note wasn't part of the 5 items decided today; remains genuinely unresolved in the implementation-facing spec |
| Repudiation | Confirmed logged (`organization_logs` scope explicitly includes representative management) | — | None |
| Information disclosure | N/A | — | — |
| Denial of service | N/A | — | — |

### Curation assignment — `POST /organizations/{organizationId}/curation/assign`

Unchanged from the 2026-08-19 draft — confirmed to match exactly
(conflict-of-interest `409`, one-active-assignment `409`, Admin-only).

### Curation decision — `POST /organizations/{organizationId}/curation/decision`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A | `bearerAuth` + Kurator-role check | None |
| Tampering | **Endpoint shape changed from the original draft**: this is now `POST .../curation/decision` (keyed by `organizationId`), not `PATCH .../curation-assignments/{assignmentId}`. The server resolves "the current pending assignment for this organization" itself, then checks the caller against that row's `kurator_id`. A Kurator not matching is rejected (`403`, confirmed explicit: "Current user is not the Kurator assigned to this organization's active curation assignment"). | Explicit, confirmed match-check | None — functionally equivalent to the original assignment-id-based design, this is a shape change, not a security regression |
| Tampering | **New TOCTOU consideration**: since the endpoint resolves the assignment server-side at request time rather than the client specifying a fixed assignment id, a race between "Admin reassigns the org to a different Kurator" and "the originally-assigned Kurator submits a decision" resolves however the DB read at decision-time lands — the decision-time check is against whatever is *currently* the pending assignment, not whatever the Kurator's UI last showed them. This is actually **safer** than an id-based design in one respect (no stale-assignment-id replay), but worth a concurrency test to confirm the behavior is "decision against a reassigned org is rejected with 403, not silently applied against the old assignment." | Server always re-reads current state at decision time | Low — worth an explicit test, not just an assumption |
| Repudiation | Confirmed logged | — | None |
| Information disclosure | N/A | — | — |
| Denial of service | N/A | — | — |
| Elevation of privilege | Non-Kurator or wrong-Kurator attempts decision | `403`, confirmed | None |

### `has_overdue_report` cross-domain enforcement (new — `campaign` domain)

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Tampering / correctness | A window where `has_overdue_report` is stale (e.g. just cleared by `disbursement` domain, but `campaign`'s creation check reads a cached/stale value) allows campaign creation that should've been blocked, or vice versa | Not yet specified — `disbursement` domain doesn't exist yet in this project's spec work | **Deferred, flagged for `campaign`/`disbursement` domain build** — `campaign`'s creation check should read `has_overdue_report` fresh (same-transaction or immediately-prior read), not from any cached organization record |
| Information disclosure | N/A | — | — |
| Denial of service | N/A | — | — |

## Knowingly accepted residual risk

Carried forward from 2026-08-19, plus:

- **Signed-URL exposure window** (5 minutes) for legal-document
  downloads — accepted as a standard trade-off of the signed-URL
  pattern, contingent on the URL itself not being logged in plaintext
  anywhere server-side.
- **`confirm=true` griefing** (self-inflicted repeated re-curation
  triggering) — accepted, low severity, self-harm only.

## Open items to resolve

Status as of 2026-08-20 (5 of the domain's tracked decision items are
now resolved — see `invariants.md`'s reconciliation note):

1. ~~5-organization limit on promote-to-owner~~ → **Still genuinely
   open** — not one of the 5 items decided today; no check documented
   on the promote endpoint. Needs a separate decision before Task 07
   is implemented.
2. ~~Who can edit organization fields~~ → **`[RESOLVED — 2026-08-20]`**
   split by field class.
3. Should legal/identity edits (and attachment replacement) be added
   explicitly to `organization_logs` scope? → Still open, low
   priority.
4. Does "curation decisions" in `organization_logs` scope cover the
   Admin's *assignment* action too, or only the Kurator's *decision*?
   → Still open, low priority.
5. ~~Legal-document download path re-check~~ → **Resolved by design**
   (signed-URL pattern, no cached-authorization risk).
6. ~~`openapi.yaml` naming/structure~~ → **Resolved 2026-08-20**.
7. ~~`GET /organizations/{organizationId}` has no documented `403`~~ →
   **`[RESOLVED — 2026-08-20]`** broadly readable, `npwp` gated
   separately.
8. ~~Admin-exclusion enforcement gap~~ → **`[RESOLVED — 2026-08-20]`**
   `409 admin-cannot-be-representative`.
9. ~~`izin_pub` re-curation conflict~~ → **`[RESOLVED — 2026-08-20]`**
   follows implementation (excluded); `kencleng-phase1-detail.md`
   Fitur 1 needs a matching update (now the stale document).

**Remaining open, not blocking**: items 1, 3, 4 above.

## References

- Related domain invariants: `docs/spec/organization/invariants.md`
  (reconciled 2026-08-20 — read together with this file, the three
  `[NEEDS DECISION]` items are shared between both docs)
- Related ERD: `docs/project/kencleng-erd.md` §2
- Related business process: `docs/project/kencleng-phase1-detail.md`
  Fitur 1, 1B, 2, 5
- Related actors/rules: `docs/project/kencleng-actors-entities.md`
- **Actual API (ground truth as of 2026-08-20)**:
  `api/openapi/organization.yaml`, `api/openapi/campaign.yaml` (for
  `has_overdue_report`)
- Related threat model precedent: `docs/spec/account/threat-model.md`
