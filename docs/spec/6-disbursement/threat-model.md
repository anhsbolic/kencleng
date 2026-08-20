# Threat Model — disbursement

> File: `docs/spec/disbursement/threat-model.md`
> Status: draft — authored directly against `api/openapi/disbursement.yaml` 2026-08-20
> Last updated: 2026-08-20

## Actors & trust boundaries

| Actor | Authenticated? | Trust boundary crossed |
|---|---|---|
| Public / anonymous visitor | No | `GET /campaigns/{id}/report` (closed campaigns only, permanent archive) |
| Owner | Yes | Report narrative edit, disbursement request creation, fund-usage report submission, attachment upload |
| Staff | Yes | Read-only for this domain — confirmed explicitly excluded from narrative edit and disbursement/report actions (Business Rule 4) |
| Admin | Yes | Disbursement decision, curation assignment (fund-usage verification), Admin-wide disbursement queue |
| Kurator | Yes | Fund-usage report verification — must recuse if a representative of the report's organization |
| System (disbursement execution, `has_overdue_report` scheduler) | N/A — in-process | Approved→disbursed transition; the 30-day-overdue detection scheduler |

## STRIDE per operation

### Campaign report & narrative — `GET .../report`, `PATCH .../report-narrative`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A on the public report view; `bearerAuth` + owner-only on narrative edit | — | None |
| Tampering | `staff` attempts narrative edit | `403`, confirmed explicit | None |
| Tampering | Narrative edited before campaign is `closed` | `409` | None |
| Repudiation | Narrative edits aren't explicitly logged per the ERD's `fund_usage_report_logs`/`disbursement_request_logs` scope (narrative is a `campaigns` field, not covered by either) | Not currently in scope | Low, accepted — narrative is explicitly framed as non-financial/non-accountability content; a full edit history isn't part of the stated design |
| **Information disclosure / XSS — `[SECURITY-CRITICAL, confirmed mitigation exists]`** | `report_narrative` is user-supplied Markdown displayed on a permanent **public** page — if rendered as raw HTML without sanitization, this is a stored-XSS vector against every visitor of that campaign's report, indefinitely | `kencleng-phase3-detail.md` explicitly mandates frontend sanitization (`react-markdown` + `rehype-sanitize` or equivalent, never `dangerouslySetInnerHTML` on unsanitized output) — backend stores raw Markdown only, no HTML generation server-side | **Must be verified at implementation time on the frontend side** — this domain's backend spec has no control over frontend rendering choices, but the risk is severe enough (permanent public XSS) to flag explicitly here too, not just in the frontend tech-stack doc |
| Denial of service | N/A | — | — |
| Elevation of privilege | N/A | — | — |

### Disbursement request — `POST`/`GET /campaigns/{id}/disbursement-requests`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A | `bearerAuth` + owner-only (create) | None |
| Tampering | `staff` attempts to create a request | `403`, confirmed explicit | None |
| Tampering | Request created against a non-`closed` campaign | `409 campaign-not-closed` | None |
| Tampering | Double-disbursement attempt (two active requests on one campaign) | DB unique index (`ux_disbursement_requests_one_active`), confirmed `409` | None |
| Repudiation | Request creation IS logged (`disbursement_request_logs` scope) | — | None |
| Information disclosure | List/detail visibility — assumed representative/Admin scoping, not explicitly documented with an example | Assumed default per this domain's `invariants.md` note | **Low, flagged for confirmation** — same class of ambiguity as elsewhere, not a genuine conflict |
| Denial of service | N/A | — | — |
| Elevation of privilege | N/A | — | — |

### Disbursement decision & execution — `POST .../decision`, internal execution process

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A | `bearerAuth` + Admin-only | None |
| Tampering | Non-Admin attempts decision | `403` | None |
| Tampering | Decision attempted on a non-`pending` request | `409` | None |
| **Tampering — internal execution process must never be HTTP-reachable** | Same class of risk as `donation` domain's payment settlement — if the approved→disbursed transition were ever accidentally exposed as a route, anyone could forge a "funds disbursed" confirmation | Documented as a system-triggered follow-up, not a client-callable action — must be verified at implementation time, same critical flag as `donation`'s settlement process | **Critical implementation-time requirement**, same severity class as `donation/threat-model.md`'s equivalent flag |
| Repudiation | Decision logged (confirmed, `disbursement_request_logs` scope) | — | None |
| Information disclosure | N/A | — | — |
| Denial of service | N/A | — | — |
| Elevation of privilege | Non-Admin attempts decision | `403` | None |

### Fund-usage report submission & attachments

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A | `bearerAuth` + owner-only | None |
| Tampering | `staff` attempts submission or attachment upload | `403` | None |
| Tampering | Submission before `status = 'disbursed'` | `409` | None |
| Tampering | Reconciliation mismatch (items don't sum to exactly `requested_amount`) — an attempt to under- or over-report | `422`, strict no-tolerance match (INV-disbursement-06) | None — this is the domain's core financial-integrity control |
| Tampering | Invalid file type/oversized attachment | `422` | None |
| Repudiation | Submission logged (`fund_usage_report_logs` scope) | — | None |
| Information disclosure | Attachment `download_url` (5-minute signed URL) leakage — same class of risk as `organization`'s legal-document signed URLs | 5-minute TTL bounds exposure; same accepted trade-off as `organization`'s equivalent | Accepted, same posture as `organization/threat-model.md`'s signed-URL note — worth the same "confirm it isn't logged in plaintext server-side" check |
| Denial of service | N/A | — | — |
| Elevation of privilege | N/A | — | — |

### `has_overdue_report` scheduler & clearing logic

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Tampering / correctness | A window where the flag is stale relative to `campaign` domain's creation-time check (INV-organization-13/INV-campaign-01) | Both the set (scheduler, conditional `UPDATE`) and clear (same transaction as report submission) are designed to avoid a stale-read window; `campaign`'s creation check must read fresh state, per that domain's own spec | Low — the cross-domain contract is sound as designed; the residual risk is purely an implementation-discipline one (don't introduce caching between these domains) |
| Repudiation | Flag set/clear is logged (confirmed: "Any change to this flag... is logged in the Audit Log," `kencleng-phase3-detail.md` Fitur 4) — verify this actually lands in `organization_logs` (the field's owning domain) rather than `disbursement`'s own logs, since the field itself belongs to `organization` | Not yet confirmed which table receives this log entry | **Open item** — needs a decision: does the disbursement-triggered flag change get logged to `organization_logs` (field owner) or somewhere in `disbursement`? Recommend `organization_logs`, consistent with the "invariants owned by the domain that owns the field" convention already established, but flagging since it's a cross-domain write not yet made explicit anywhere |
| Information disclosure | N/A | — | — |
| Denial of service | N/A | — | — |
| Elevation of privilege | N/A | — | — |

### Fund-usage verification (assignment & decision)

Same shape as `organization`/`campaign`'s equivalent curation
sections — conflict-of-interest check, one-active-assignment guard,
assigned-Kurator-only decision, TOCTOU-aware reassignment handling.
See `docs/spec/organization/threat-model.md`'s "Curation assignment"
and "Curation decision" sections for the full analysis; confirmed to
match the same pattern in `disbursement.yaml`.

## Knowingly accepted residual risk

- **List/detail visibility assumptions** for disbursement requests and
  fund-usage reports (representative/Kurator/Admin, not explicitly
  documented with an error example) — low-severity ambiguity, flagged
  for confirmation rather than blocking.
- **Signed-URL exposure window** for fund-usage attachments — same
  accepted trade-off as `organization`'s legal documents.
- **No cumulative penalty for repeated fund-usage report rejection** —
  intentional, per INV-disbursement-09, not a gap.

## Open items to resolve

1. **Critical implementation-time requirement**: the internal
   disbursement-execution process (approved→disbursed) must never be
   reachable via any registered HTTP route — same severity as
   `donation` domain's settlement process, verify explicitly during
   code review.
2. Confirm `has_overdue_report` set/clear events log to
   `organization_logs` (the field's owning table), not somewhere in
   `disbursement` — not yet made explicit anywhere.
3. Confirm the assumed representative/Kurator/Admin visibility scoping
   on disbursement-request and fund-usage-report list/detail endpoints
   matches the actual intended design (low priority).

## References

- Related domain invariants: `docs/spec/disbursement/invariants.md`
- Related ERD: `docs/project/kencleng-erd.md` §5
- Related business process: `docs/project/kencleng-phase3-detail.md`
  Fitur 1–5
- Related threat model precedent: `docs/spec/organization/threat-model.md`
  (curation sections, signed-URL pattern),
  `docs/spec/donation/threat-model.md` (internal-process-must-never-be-
  HTTP-reachable pattern)
- **Actual API (ground truth)**: `api/openapi/disbursement.yaml`
