# Threat Model — notification

> File: `docs/spec/notification/threat-model.md`
> Status: draft
> Last updated: 2026-08-19

## Actors & trust boundaries

| Actor | Authenticated? | Trust boundary crossed |
|---|---|---|
| Registered user (Donatur, Organisasi Representative, Kurator, Admin — any authenticated role) | Yes (bearer JWT, `bearerAuth`, reused from `account`) | Client → API, for all three HTTP endpoints (`GET /notifications`, `GET /notifications/unread-count`, `POST /notifications/mark-read`) |
| Guest donor | No | None via this domain's API — guests never call these endpoints (no session). Their only touchpoint is as a passive recipient of the `email` channel (`recipient_email`), which is an outbound side effect, not an inbound trust boundary. |
| Other backend domains (`account`, `organization`, `campaign`, `donation`, `disbursement`) | N/A — in-process | **No network boundary crossed.** Per `api/openapi.yaml`'s `notification` tag description, notification creation has no HTTP endpoint; it's an in-process function call within the same monolith (`internal/domain/notification/`). These callers are trusted by construction (same binary, same deploy unit), but the *data* they pass in (e.g. a campaign title destined for `payload`) may itself originate from untrusted end-user input further upstream — see "Payload content" under `POST /notifications/mark-read`'s sibling section below. |
| Weekly hard-delete worker | N/A — trusted system process | None — no external input, runs on a schedule inside the same trust zone as the app. |

## STRIDE per component/endpoint

### `GET /notifications`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | Caller without a valid session tries to list notifications | `bearerAuth` required (global `security` in `api/openapi.yaml`), same JWT validation as every other authenticated endpoint | None beyond what `account` already covers |
| Tampering | Malformed/forged `cursor` value used to probe past the intended page boundary | `cursor` is an opaque UUIDv7 from a prior response; query is always scoped `WHERE recipient_user_id = current_user` regardless of cursor value (INV-notification-04) — a forged cursor can at worst produce an empty/garbled page, never another user's row | Low — a malformed cursor should fail closed (422) rather than silently falling back to page 1; confirm this in the feature spec |
| Repudiation | N/A — read-only, and reading one's own notifications is not on the sensitive-action list in `kencleng-phase0-detail.md` Fitur 9 | — | — |
| Information disclosure | Query scoping bug returns another user's notifications | `WHERE recipient_user_id = current_user` enforced at the query layer, not a fetch-then-filter-in-app-code pattern (INV-notification-04) | Low — standard "trust the DB filter, test it" risk, same class as every other per-user list endpoint in this project |
| Denial of service | Caller requests very large pages repeatedly | `limit` capped at 50 globally (`LimitParam` in `api/openapi.yaml`) | None beyond standard rate-limiting, which is out of scope for this domain specifically (project-wide concern) |
| Elevation of privilege | N/A — every authenticated user gets only their own notifications, no role-gating needed here | — | — |

### `GET /notifications/unread-count`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | Same as above | `bearerAuth` | None |
| Tampering | N/A — no mutable input, no query params | — | — |
| Repudiation | N/A | — | — |
| Information disclosure | Count leaks presence/absence of another user's notifications | Same `recipient_user_id` scoping as list | None beyond the list endpoint's own risk profile |
| Denial of service | Called on every page load (badge) — could be abused for polling-storm | `COUNT` query hits the same partial index as list (`ix_notifications_recipient_unread`); no unbounded work per call | Low — acceptable for a sandbox project; no rate-limiting layer added specifically for this endpoint |
| Elevation of privilege | N/A | — | — |

### `POST /notifications/mark-read`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | Same as above | `bearerAuth` | None |
| Tampering | `notification_ids` array includes other users' UUIDs (enumeration attempt, or accidental client bug) | `UPDATE ... WHERE recipient_user_id = current_user AND read_at IS NULL` — foreign ids are silently no-op'd, not rejected with a 403 that would confirm existence (INV-notification-04, consistent with the project's established anti-enumeration posture from `account`) | None — this is the deliberate design |
| Repudiation | N/A — not a Fitur-9 sensitive action, no audit log entry expected | — | — |
| Information disclosure | Response could leak which ids existed vs. didn't (side channel) | `204 No Content` — no body, no per-id result reported, so a mixed valid/invalid batch is indistinguishable from an all-valid batch | None |
| Denial of service | `notification_ids` array has **no `maxItems`** in `api/openapi.yaml` — a caller could submit an arbitrarily large array, causing one huge `UPDATE ... WHERE id IN (...)` | None currently specified | **Open gap** — recommend adding a `maxItems` bound (e.g. matching `LimitParam`'s 50, or slightly higher to cover "mark all visible" batching) to `MarkNotificationsReadRequest.notification_ids` before this feature is implemented; flagged below and in "Knowingly accepted residual risk" |
| Elevation of privilege | N/A | — | — |

### Internal notification creation (in-process, all calling domains)

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A — not network-exposed | — | — |
| Tampering | `payload` (JSONB) may carry content that originated from untrusted end-user input several hops upstream (e.g. a campaign title set by an Organisasi Representative, forwarded into a `campaign_approved` notification's payload) — if a consumer renders it unsafely, this is a stored-injection vector | Frontend renders `payload` as data (React's default text escaping), not as raw HTML — no `dangerouslySetInnerHTML` path for notification content per current FE conventions | **Accepted, shared-responsibility risk** — the `notification` domain itself does not sanitize `payload` (it's type-shaped free-form JSON per `NotificationType`); the calling domain and the FE/email-template layer are each responsible for treating it as untrusted data. Flagged explicitly since a future raw-HTML email implementation would need to escape/encode `payload` fields itself. |
| Repudiation | Notification-creation calls are not logged to any audit/log table | Deliberate — notification creation isn't on the Fitur 9 sensitive-action list, and audit-log ownership stays with `account`'s `user_logs` pattern for identity-related actions only | None — accepted, this isn't a security-relevant action requiring non-repudiation |
| Information disclosure | `payload` accidentally includes more data than the recipient should see (e.g. internal-only fields copied wholesale from the triggering domain's model instead of a minimal, type-specific shape) | Convention, not an enforced technical control: each calling domain's feature spec should define the minimal `payload` shape per `type` | **Accepted risk, process-level mitigation only** — worth a spot-check during each calling domain's threat model / feature-spec review (e.g. `organization`, `campaign`) rather than being fixed here |
| Denial of service | A calling domain fans out one notification-create call per recipient in a loop (e.g. `campaign_closed` to thousands of donors) one row at a time | None enforced by `notification` domain — batch-insert is the calling domain's implementation choice | Deferred — not yet a concern at this project's sandbox scale; if it becomes one, belongs to the calling domain's own performance work, not `notification`'s spec |
| Elevation of privilege | N/A — no user-supplied role assertion reaches this call | — | — |

### Weekly hard-delete worker

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing / Tampering / Repudiation | N/A — scheduled system process, no external input | — | — |
| Information disclosure | N/A — deletes rows, doesn't expose them | — | — |
| Denial of service | A single unbounded `DELETE WHERE expires_at < now()` could hold a long-running lock if the table has grown very large | `ix_notifications_expires_at` (plain B-tree, per `kencleng-erd.md`'s note that this index is on the hot read path) makes the scan itself cheap | Low, sandbox-scale — recommend batching the delete (e.g. `DELETE ... LIMIT N` in a loop) if/when row counts become large enough to matter; not required for v1 |
| Elevation of privilege | N/A | — | — |

## Knowingly accepted residual risk

- **Best-effort notification delivery, no retry** (INV-notification-06):
  a notification can silently fail to be created for an event that
  otherwise succeeded, with no retry mechanism in v1. Accepted because
  the alternative (same-transaction, strict) would make an unrelated
  domain's core operation depend on the `notifications` table's
  availability — worse trade-off for a sandbox project with no SLA.
- **`notification_ids` array has no `maxItems` bound** in the current
  `api/openapi.yaml` — a real (non-sandbox) deployment would want this
  capped before shipping; flagged here as a concrete, actionable gap
  to close in the feature spec for `POST /notifications/mark-read`,
  rather than deferred indefinitely.
- **`payload` content safety is a shared responsibility**, not enforced
  by this domain — `notification` stores and returns whatever
  type-shaped JSON it's given; sanitization/escaping is the calling
  domain's and the rendering layer's job. Accepted because enforcing a
  schema-per-type at the `notification` domain level would add
  coupling back to every calling domain's data shape, contradicting
  the "notification is a thin, generic delivery mechanism" design.
- **Hard-delete worker isn't batched** — acceptable at current/expected
  sandbox data volume; would need revisiting only if the project were
  ever taken past a learning-sandbox scale (explicitly out of scope
  per this project's stated purpose).

## References

- Related domain invariants: `docs/spec/notification/invariants.md`
- Related ERD: `docs/project/kencleng-erd.md` §6 "Cross-Cutting"
- Related business process: `docs/project/kencleng-phase0-detail.md`
  Fitur 6 (Notification Infrastructure)
- Related OpenAPI: `api/openapi.yaml` — `notification` tag
