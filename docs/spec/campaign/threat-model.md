# Threat Model — campaign

> File: `docs/spec/campaign/threat-model.md`
> Status: draft — authored directly against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-08-20

## Actors & trust boundaries

| Actor | Authenticated? | Trust boundary crossed |
|---|---|---|
| Public / anonymous visitor | No | `GET /campaigns` (published-only list), `GET /campaigns/{id}` (published only), `GET /campaigns/{id}/attachments` (published only, per the 2026-08-20 decision), `GET /events/{id}` — the platform's public browse/conversion surface |
| Owner/Staff representative | Yes | Draft CRUD, media upload, submit-for-curation (owner-only), publish/unpublish/republish (owner-only), event creation |
| Admin | Yes | Curation assignment, force-close, broad org-scoped campaign listing |
| Kurator | Yes | Curation review & decision — must recuse if a representative of the campaign's organization |
| System (auto-unpublish, auto-close triggers) | N/A — in-process | Organization re-verification transaction (INV-campaign-11, cross-domain from `organization`); donation-success transaction (`max_amount` trigger, cross-domain from `donation`); deadline scheduler |

## STRIDE per operation

### Public campaign listing & detail — `GET /campaigns`, `GET /campaigns/{campaignId}`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A — explicitly `security: []`, no auth expected | — | — |
| Tampering | N/A — read-only | — | — |
| Repudiation | N/A | — | — |
| Information disclosure | Non-`published` campaign detail exposed to the public | `403` for non-representative/Kurator/Admin on non-`published` campaigns (INV-campaign-14) | **Existence-confirming `403`**: unlike this project's usual anti-enumeration `404` pattern, this endpoint returns `403` (not `404`) for a non-public campaign — confirming a `draft`/`pending_curation` campaign *exists* at that id, just not its content. Low severity (campaign ids aren't sequential/guessable, and no sensitive data leaks beyond existence), but a deliberate departure from the pattern used elsewhere — worth a conscious sign-off, not an oversight. |
| Denial of service | `q` free-text search param (`# INFERRED`, not in the phase doc) — unindexed `LIKE`/full-text search could be expensive at scale | Low real-world risk at this project's sandbox scale | Low, accepted |
| Elevation of privilege | N/A | — | — |

### Campaign media (attachments) — `GET .../attachments`, `POST .../attachments`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A on list (public when published, per decision); `bearerAuth` + representative check on upload | — | None |
| Tampering | `staff`/unrelated user attempts upload | `403`, owner/staff representative required | None |
| Information disclosure — **`[RESOLVED — 2026-08-20]`** | List endpoint was unconditionally public regardless of campaign status, inconsistent with the detail endpoint's gating | Now matches the detail endpoint's gating (INV-campaign-14) | None — see `invariants.md` for the decision record |
| Denial of service | Repeated uploads to exhaust storage | File type/size limits (JPG/PNG, 5 MB max, confirmed) | Low — no per-campaign upload count cap documented; likely fine for a sandbox project, flag if it ever matters |
| Elevation of privilege | N/A beyond the tampering case above | — | — |

### Campaign draft CRUD — `POST .../campaigns`, `PATCH/DELETE /campaigns/{id}`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A | `bearerAuth` + representative check | None |
| Tampering | Non-representative attempts create/edit/delete | `403` | None |
| Tampering | Edit/delete attempted outside `status = draft` | `409` (INV-campaign-04) | None |
| Tampering | Creation against an unverified or overdue-report organization | `409` (INV-campaign-01, references INV-organization-13) | None |
| Repudiation | Draft create/edit/delete not logged to `campaign_logs` (not on the sensitive-action list — drafts are pre-curation, low-stakes) | Deliberate, consistent with the project's audit-scope philosophy (only sensitive/decision actions logged) | None — accepted, matches the pattern of not logging non-destructive actions |
| Information disclosure | N/A — draft visibility already covered by INV-campaign-14 | — | — |
| Denial of service | A representative spams draft creation (no documented cap on drafts per org) | None documented | Low — accepted for a sandbox project; a real deployment might want a cap, not necessary here |
| Elevation of privilege | N/A | — | — |

### Submit for curation — `POST /campaigns/{campaignId}/submit`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A | `bearerAuth` + owner-only | None |
| Tampering | `staff` attempts submit | `403`, confirmed explicit | None |
| Tampering | Submit attempted outside `status = draft` | `409` | None |
| Repudiation | Not separately logged (the resulting curation assignment is) | Consistent with other domains' pattern (the assignment/decision is the auditable event, not every state transition) | None |
| Information disclosure | N/A | — | — |
| Denial of service | N/A | — | — |
| Elevation of privilege | N/A | — | — |

### Curation assignment & decision — `POST .../curation/assign`, `POST .../curation/decision`, `GET .../curation-assignments/mine`, `GET /campaigns/curation-queue`

Same shape as `organization`'s equivalent section — see
`docs/spec/organization/threat-model.md`'s "Curation assignment" and
"Curation decision" sections for the full analysis (conflict-of-
interest check, one-active-assignment guard, assigned-Kurator-only
decision check, and the TOCTOU consideration on server-resolved
current-assignment lookup). All confirmed to match the same pattern
in `campaign.yaml`.

### Publish / unpublish / republish — `POST .../publish`, `POST .../unpublish`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A | `bearerAuth` + owner-only | None |
| Tampering | `staff` attempts publish/unpublish | `403` | None |
| Tampering | Publish attempted from an invalid starting status | `409` | None |
| Tampering | `publish_at` set in the past or beyond `deadline` | `422` | None |
| Repudiation | Unpublish requires `decision_note`, logged (INV-campaign-10, confirmed) | — | None |
| Information disclosure | N/A | — | — |
| Denial of service | Owner repeatedly schedules/reschedules to grief their own campaign's visibility | Low severity, self-inflicted | Low, accepted |
| Elevation of privilege | N/A | — | — |

### Force-close — `POST /campaigns/{campaignId}/force-close`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A | `bearerAuth` + Admin-only | None |
| Tampering | Non-Admin attempts force-close | `403` | None |
| Tampering | **Three-way race** between max_amount-trigger (donation), deadline-scheduler, and Admin force-close all firing near-simultaneously on the same campaign | `WHERE status = 'published'` guard shared by all three (INV-campaign-13) — whichever commits first wins, others become clean `409`s, never a crash or double-close | None — this is exactly the scenario the shared guard is designed for; worth an explicit 3-way concurrency test, not just pairwise |
| Repudiation | Requires `decision_note`, records `closed_by` (confirmed) | — | None |
| Information disclosure | N/A | — | — |
| Denial of service | N/A | — | — |
| Elevation of privilege | Non-Admin attempts force-close | `403` | None |

### Events — `POST /events`, `GET /events/{id}`, `GET /organizations/{id}/events`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A on creation beyond standard auth | `bearerAuth` + representative check | None |
| Tampering | Linking a campaign from a different organization | `403` (INV-campaign-16) — note this is a content-revealing `403` (confirms the campaign belongs to another org), accepted as low-severity per `invariants.md`'s note (caller already knows the id they tried) | None, accepted |
| Tampering | Linking a non-`published`/`scheduled` campaign | `409` | None |
| Repudiation | Event creation not logged (consistent with Events' "non-sensitive, no financial data" characterization, `kencleng-phase1-detail.md` Fitur 6) | Deliberate | None |
| Information disclosure | `GET /events/{id}` is fully public (`security: []`), no gating tied to linked campaigns' status | Events are explicitly non-sensitive/promotional by design — even if all linked campaigns are somehow non-public (shouldn't happen per INV-campaign-16's link-time check, though see the "no re-check after linking" note in `invariants.md`), the Event's own fields (name, datetime, location, description) carry no sensitive data | Low, accepted — consistent with the domain's own stated risk posture for this entity |
| Denial of service | N/A | — | — |
| Elevation of privilege | N/A | — | — |

## Knowingly accepted residual risk

- **`403` (not `404`) on non-public campaign detail/attachments** —
  existence-confirming, departs from this project's usual
  anti-enumeration pattern. Accepted as low-severity (campaign ids
  aren't sequential, no content leaks) but flagged as a conscious
  choice to revisit if it ever matters more (e.g. if campaign ids
  become guessable/sequential in a future iteration).
- **No draft-creation rate cap per organization** — accepted for a
  sandbox project.
- **Stale `campaign_events` links after a linked campaign closes** —
  cosmetic only, per Events' low-sensitivity design.
- **Event detail is fully public regardless of linked campaigns'
  status** — accepted, consistent with Events carrying no sensitive
  data by design.

## References

- Related domain invariants: `docs/spec/campaign/invariants.md`
- Related ERD: `docs/project/kencleng-erd.md` §3
- Related business process: `docs/project/kencleng-phase1-detail.md`
  Fitur 3–6, `docs/project/kencleng-phase2-detail.md` Fitur 3
- Related threat model precedent: `docs/spec/organization/threat-model.md`
  (curation assignment/decision sections, referenced directly above)
- **Actual API (ground truth)**: `api/openapi/campaign.yaml`
