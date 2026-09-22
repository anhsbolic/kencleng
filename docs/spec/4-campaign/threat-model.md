# Threat Model — campaign

> File: `docs/spec/campaign/threat-model.md`
> Status: Slice-1 public-boundary reconciliation active; remaining historical areas are deferred evidence
> Last updated: 2026-09-22

## Actors & trust boundaries

| Actor | Authenticated? | Trust boundary crossed |
|---|---|---|
| Public / anonymous visitor | No | Slice 1: `GET /campaigns/{id}` and controlled `GET /campaigns/{id}/media/{mediaId}/content`; listing and attachment-list/upload are deferred |
| Owner/Staff representative | Yes | Draft CRUD, media upload, submit-for-curation (owner-only), publish/unpublish/republish (owner-only), event creation |
| Admin | Yes | Curation assignment, force-close, broad org-scoped campaign listing |
| Kurator | Yes | Curation review & decision — must recuse if a representative of the campaign's organization |
| System (auto-unpublish, auto-close triggers) | N/A — in-process | Organization re-verification transaction (INV-campaign-11, cross-domain from `organization`); donation-success transaction (`max_amount` trigger, cross-domain from `donation`); deadline scheduler |

## STRIDE per operation

### Slice-1 public Campaign detail — `GET /campaigns/{campaignId}`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A — explicitly `security: []`, no auth expected | — | — |
| Tampering | N/A — read-only | — | — |
| Repudiation | N/A | — | — |
| Information disclosure | Internal Campaign/Organization fields, non-public existence, or auth-dependent payload leaks | Public eligibility before a standalone closed allowlist mapping; identical `404` for absent/invalid/non-public; optional auth invariant; no `allOf` inheritance | Timing parity and runtime forbidden-field tests remain required downstream. |
| Information disclosure | Organizer prose rendered as markup or provenance presented as platform verification | Plain strings plus `source: organizer`; no HTML/Markdown contract | Frontend must use framework escaping and human-reviewed final wording. |
| Denial of service | Detail/funding dependency is unavailable | Documented `503` rather than fabricated unavailable product truth | Availability/rate limits are deferred to backend/topology work. |
| Elevation of privilege | N/A | — | — |

### Slice-1 Campaign media delivery — `GET .../media/{mediaId}/content`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | Optional Authorization alters the public response | `security: []` and INV-campaign-14 auth invariance | Runtime parity tests remain required downstream. |
| Tampering | N/A — read-only byte delivery | — | — |
| Information disclosure | Direct public-bucket/signed URL bypasses retraction, or media membership leaks | Private storage; parent/member recheck on every origin request; same `404` for absent/non-public/non-member; no redirect/object URL | Already downloaded client-held bytes cannot be withdrawn. |
| Information disclosure | Shared/browser cache serves stale public bytes or metadata after retraction | `Cache-Control: private, no-store` on success and public errors | Topology must preserve header; real proxy/cache evidence remains required. |
| Denial of service | Storage/object dependency cannot serve eligible bytes | Explicit `503`, distinct from `media.absent` and `404` | Dependency outage behavior requires downstream integration evidence. |
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

- **Already downloaded media cannot be withdrawn** — `private, no-store` and
  origin rechecks prevent new system-controlled fetches after retraction,
  but cannot erase bytes previously saved by a client.
- **Historical listing/upload/privileged behavior is deferred** — it is not
  active Slice-1 authority and needs later reconciliation before runtime use.
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
