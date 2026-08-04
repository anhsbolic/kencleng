# Kencleng — Roadmap: Next Steps (Documentation & Spec Phase)

> Status: **Documentation/spec phase fully closed** (2026-07-27). All
> outstanding artifacts flagged in the 2026-07-24/26 addenda are now
> complete: OpenAPI spec authoring (`api/openapi.yaml`, 2026-07-26) and
> frontend design guidelines (`kencleng-design-guidelines.md`,
> 2026-07-27). No documentation/spec work remains before the
> development phase.
> Last updated: 2026-08-04 (development-phase kickoff decision added —
> see "Development Phase — Kickoff Decisions" below; documentation/spec
> phase itself remains fully closed)

## Scope of this document

This roadmap tracks only the **documentation/spec phase** of Kencleng —
turning the project's business rules, architecture, and design
decisions into a complete, internally-consistent set of documents
ready to hand off to actual development. It does **not** track coding,
scaffolding, or deployment work; those belong to a separate
development-phase tracker once this phase is considered closed.

Consequently, the previous Step 5 (project skeleton setup), Step 7
(local deployment target / `docker-compose.yml`), and Step 8 (first
vertical slice) have been **removed from this document** — see
"Removed From Scope" at the bottom for what they covered and why they
no longer belong here.

**2026-07-26 addendum**: even though the roadmap was declared closed on
2026-07-24, two concrete artifacts remained outstanding — the actual
`api/openapi.yaml` file (format/philosophy had been *decided* at Step 2,
but the endpoints themselves hadn't been authored yet) and frontend
design guidelines (visual styling had been explicitly deferred past
Step 1.5's wireframe consolidation). Anhar flagged both as needed before
starting the development phase in earnest.

**2026-07-27 addendum**: frontend design guidelines completed —
see "Step 10 — Frontend design guidelines" below. Both addendum items
are now done; the documentation/spec phase is fully closed.

**2026-08-04 addendum**: first development-workflow planning session.
The documentation/spec phase itself has no new open items, but one
development-phase kickoff decision was made — the domain-by-domain
*coding* order — see "Development Phase — Kickoff Decisions" below and
`kencleng-agentic-workflow.md` §3.3 for full rationale.

## Context

This document exists to answer one question: **what documentation/
spec work remains before the conceptual design of Kencleng is
considered complete and handoff-ready?** It's a sequencing/tracking
doc, not a design doc itself — each step below either produces its own
document/artifact, or directly modifies an existing one.

**Sequencing principle (established early in this project)**:
high-level scope → actor/entity modeling → business process phases →
infrastructure (tech stack) → UX mapping → ERD/API → (handoff to
development). Steps are ordered by this dependency chain.

**2026-07-24 — three passes through the Open Items list in one
session**, ending in full closure:
1. **Critical pass**: 5 items directly blocking the first vertical
   slice (Registrasi & Login) — see "Critical Open Items Resolved."
2. **Further-items pass**: beneficiary entity, representatives page
   business rules, `MaskedField` reveal persistence, and the 3
   un-wireframed Admin pages — see "Further Open Items Resolved."
   Step 6 (frontend testing) closed in the same pass.
3. **Full closure pass**: every remaining item across every doc —
   deployment topology (which unblocked CORS/cookie config), org-level
   policy questions (NPWP validation, org-per-user limit, notification
   mechanism, audit-trail scope, guest donor label, PWA scope,
   repeated-rejection consequence, role-revocation handling), plus a
   full doc-sync pass across every doc that had stale open items or
   cross-references. See "Full Open Items Closure" below.

**2026-07-26 — OpenAPI spec authoring**, generated domain-by-domain
following the 6-domain boundary from Step 4 (`account` → `organisasi` →
`campaign` → `donation` → `disbursement` → `notification`, in dependency
order). See "OpenAPI Spec Authoring" below.

Every project doc has been touched and brought current as of the
2026-07-24 session: `kencleng-erd.md`, `kencleng-backend-tech-stack.md`,
`kencleng-frontend-tech-stack.md`, `kencleng-phase0-detail.md`,
`kencleng-phase1-detail.md`, `kencleng-phase3-detail.md`,
`kencleng-ux-page-map.md`, `kencleng-actors-entities.md`, and
`kencleng-business-process-overview.md`.

---

## Status Summary Table

| Step | Description | Output | Status |
|---|---|---|---|
| 1 | Close small open items (TTLs, policy numbers, expiration durations) | Inline doc updates | ✅ Done |
| 1.5 | UX wireframe consolidation + benchmark design | Wireframe artifact + benchmark list | ✅ Done |
| 2 | Decide API contract format + design philosophy | Decision doc / stack doc update | ✅ Done |
| 3 | ERD / full schema design | `kencleng-erd.md` + draft migrations | ✅ Done |
| 4 | Domain boundaries & migration strategy | Stack doc update | ✅ Done |
| 6 | Frontend testing approach | Stack doc update | ✅ Done |
| 9 | **OpenAPI spec authoring** (all 6 domains) | `api/openapi.yaml` | ✅ Done |
| 10 | **Frontend design guidelines** | `kencleng-design-guidelines.md` | ✅ Done |

Steps 5, 7, 8 removed — see "Removed From Scope." **All tracked steps
are now ✅ Done — documentation/spec phase fully closed.**

---

## ✅ Done

### Step 1 — Close remaining small open items
All decisions resolved and recorded inline in their respective docs
(strikethrough + `[RESOLVED]` annotations):
- Access token TTL (15 min) / refresh token TTL (30 days)
- Password policy (length-only, min 8 chars) / email verification
  token expiry (24h)
- HaveIBeenPwned breach-list check (k-anonymity model)
- Password reset token expiry (1h)
- Forgot-password behavior for Google-only users
- MFA backup code count (10) & regeneration policy (disable-enable
  only)
- File upload max size (5 MB all contexts) / signed URL expiry (5 min)
- Notification expiration (30 days) / hard-delete worker frequency
  (weekly)
- Fund-usage report reconciliation (strict match) / deadline (30 days
  post-disbursement) / consequence (`has_overdue_report` flag, blocks
  new campaign creation)
- Report narrative format (Markdown, 5000 character limit, mandatory
  sanitize)
- Minimum donation (Rp 5.000) / payment simulation parameters (5%
  failure rate, 2–5s delay)

### Step 1.5 — UX wireframe consolidation & design direction
Resolved 2026-07-20: Dashboard Shell (top-nav desktop, top-bar +
hamburger mobile), Benchmark Design Reference (GoFundMe primary,
Kitabisa secondary), wireframe tool (plain HTML/SVG gray-box,
`kencleng-wireframes/`). Every page — including the 3 Admin pages
closed out 2026-07-24 — now has both a mobile and desktop wireframe
(or an explicit note where mobile reuses the existing shell verbatim).

### Step 2 — API contract format + design philosophy
OpenAPI 3.x spec-first — `api/openapi.yaml` as source of truth,
hand-authored, documentation/contract only (no codegen either side;
`oapi-codegen` considered and rejected for the backend, types-only
generation via `openapi-typescript` for the frontend). Full decision
and rationale in `kencleng-backend-tech-stack.md`, "API Contract &
Codegen." Every doc's stale "API contract format pending" reference
has been found and corrected in this session's doc-sync pass.

### Step 3 — ERD / full schema design
Complete in `kencleng-erd.md` — all 22 entities, concrete types,
relations, indexes, constraints, and conventions (UUID v7 PKs,
`TIMESTAMPTZ`, `NUMERIC(19,2)` + `shopspring/decimal`, `CHECK`-based
enums, PII encryption-at-rest, consistent FK delete policy, no generic
soft-delete, immutable log tables, partial unique indexes, fund-usage
reconciliation trigger). All items originally scoped to this step —
including the beneficiary entity, resolved as a `campaigns.
beneficiary_description` text field rather than a dedicated entity —
are now closed.

### Step 4 — Domain boundaries & migration execution strategy
6 domains (`account`, `organisasi`, `campaign`, `donation`,
`disbursement`, `notification`), flat package-per-domain folder
convention, `golang-migrate` CLI run manually. Full rationale in
`kencleng-backend-tech-stack.md` Open Items #3.

### Critical Open Items Resolved (2026-07-24)
5 items directly blocking the first vertical slice (Registrasi &
Login):
1. **Fail-open vs fail-closed if HaveIBeenPwned is unreachable** →
   fail-open + logging. See `kencleng-phase0-detail.md` Fitur 1.
2. **`login_attempts` lockout threshold & window** → 5 failed
   attempts / 15 minutes. See `kencleng-phase0-detail.md` Fitur 2C.
3. **Encryption key management** → 2 separate keys
   (`ENCRYPTION_KEY`/`HMAC_KEY`), env vars, no rotation in v1. See
   `kencleng-backend-tech-stack.md`, "Encryption Key Management."
4. **Google OAuth `state`/`nonce` + redirect-URI strictness** →
   HttpOnly cookie short-TTL, `nonce` validated against `id_token`,
   redirect exact-match. See `kencleng-phase0-detail.md` Fitur 1B.
5. **Set-password flow for Google-only users** → via
   `/dashboard/security` "Atur Password," `verified_at = now`
   immediately, no re-auth required. See `kencleng-phase0-detail.md`
   Fitur 4.

### Step 6 — Frontend testing approach
Vitest + React Testing Library, unit/component only (no E2E in v1),
MSW for API mocking. Full detail in `kencleng-frontend-tech-stack.md`,
"Testing" section.

### Further Open Items Resolved (2026-07-24, continued)
1. **Beneficiary entity** → `beneficiary_description` free-text field
   on `Campaign`, not a dedicated entity.
2. **Representatives page business rules** → direct-add invite (no
   accept step), owner-only promote/demote/removal with the ≥1-owner
   guard. Full spec added as `kencleng-phase1-detail.md` Fitur 1B.
3. **`MaskedField` reveal persistence** → stays revealed until manual
   re-toggle or navigation/refresh — falls out of plain local
   component state, no timer needed.
4. **3 Admin pages wireframed** — downloadable standalone HTML files
   in `kencleng-wireframes/` (`admin-users-wireframe.html`,
   `admin-kurasi-queue-wireframe.html`,
   `admin-force-close-wireframe.html`).

### Full Open Items Closure (2026-07-24, final pass)
Every remaining item across every doc, resolved in one continuous
session:

1. **Deployment topology** → same-origin via reverse proxy (not
   cross-origin). This unblocked:
2. **CORS / cookie config** → resolved as a consequence of #1: no
   CORS needed at all (same-origin), refresh-token cookie is
   `HttpOnly` + `Secure` + `SameSite=Strict`.
3. **`notifications.recipient_email_hash`** → kept (not dropped),
   for future-proofing/symmetry with the encryption pattern, even
   though unused by any v1 endpoint.
4. **Repeated fund-usage-report rejection (distinct from lateness)**
   → no special consequence beyond the existing resubmit cycle; the
   30-day deadline stays anchored to `disbursed_at`, not reset on
   resubmit. See `kencleng-phase3-detail.md` Fitur 5.
5. **Revoke role vs. in-flight `pending` assignments** → no
   auto-reassignment; the assignment becomes unactionable once the
   role check fails in real time, and Admin reassigns manually via the
   existing assign flow.
6. **NPWP format validation** → format-only (regex), no external
   DJP/Ditjen Pajak verification — Kurator still verifies legitimacy
   manually via document review. See `kencleng-phase1-detail.md`
   Fitur 1.
7. **Org-per-user limit** → maximum 5 organisasi per user, app-level
   check. See `kencleng-phase1-detail.md` Fitur 1 and
   `kencleng-actors-entities.md`.
8. **Campaign registration fields** → `category` (required enum),
   `location` (optional free-text). See `kencleng-phase1-detail.md`
   Fitur 3 and `kencleng-erd.md`.
9. **Notification mechanism across curation steps** → extend the
   existing `notifications.type` enum (Admin notified on new queue
   item, Kurator on assignment, Owner on fund-usage-report/
   disbursement decisions), dual channel, no new mechanism. See
   `kencleng-phase0-detail.md` Fitur 6, `kencleng-phase1-detail.md`
   Fitur 2, `kencleng-phase3-detail.md` Fitur 2 & 5.
10. **Audit-trail granularity** → extended to cover representative
    management actions (invite/remove/promote/demote); deliberately
    *not* extended to non-destructive actions (initial submissions),
    since those already have a state trail via `status` fields. See
    `kencleng-phase0-detail.md` Fitur 9.
11. **Guest donor display label** → "Donatur" (neutral default), not
    "Hamba Allah" — avoids assuming the donor's religious identity.
12. **PWA scope** → web app manifest + basic service worker (App
    Shell caching) in scope for v1; full offline-first data behavior
    explicitly out of scope. See `kencleng-frontend-tech-stack.md`.
13. **Empty/loading/error states per page** → explicitly deferred to
    the implementation phase (standard, consistent patterns — skeleton
    loaders, empty-state messaging, inline/toast errors), decided
    during actual component build rather than at spec level. This is
    a deliberate, documented deferral, not an unresolved question.

Full doc-sync pass in the same session: `kencleng-actors-entities.md`
(all 3 Open Items + all 3 "Not Yet Discussed" items resolved and
cross-referenced) and `kencleng-business-process-overview.md` (all
Open Items + "Not Yet Discussed" resolved, plus a stale "payment rail
TBD" table note corrected).

### Step 9 — OpenAPI spec authoring (2026-07-26)

**Output**: `api/openapi.yaml` — single file, OpenAPI 3.0.3, validated
against the spec (`openapi-spec-validator`). **65 paths, 99 schemas, 6
tags.** This is the artifact every phase-detail doc had been pointing
to ("endpoint-level detail excluded from this doc — that lives in
`api/openapi.yaml`") — it now exists.

**Shared conventions established before domain authoring, apply
uniformly across all 65 endpoints:**
- **File structure**: single file, tag-per-domain (not split into
  per-domain files) — matches the "lowest complexity" principle; a
  bundler/tooling step to resolve `$ref`s across files isn't justified
  yet.
- **Success responses**: unwrapped (resource directly in the response
  body). Exception: list endpoints, which wrap as
  `{ data: [...], pagination: {...} }`.
- **Error responses**: RFC 9457 Problem Details
  (`application/problem+json`), with a `ValidationProblem` extension
  (`errors[]`) for per-field validation messages.
- **Pagination**: cursor-based, using the resource's UUIDv7 primary key
  directly as the cursor — no composite cursor needed since UUIDv7 is
  already time-orderable. Chosen over offset-based specifically because
  several v1 list endpoints (donor list, notifications) are real-time
  growing, where offset pagination has a well-known correctness bug
  (phantom rows on insert/delete during pagination).
- **Auth**: Bearer JWT (ES256) access token via `securitySchemes` for
  protected endpoints; refresh token travels only as an `HttpOnly`
  cookie (documented per-endpoint in `description`, not as a formal
  scheme, since OpenAPI 3.x cookie-auth schemes are awkward for a
  single-endpoint use case).
- **Schema naming**: `{Domain}` for the base resource (e.g. `Campaign`),
  `{Domain}Create`/`{Domain}Update` for request bodies — no `Response`
  suffix, since the base resource shape already serves as the response
  shape.
- **File upload**: backend proxy (multipart/form-data through the Go
  handler), not presigned client-to-MinIO URLs — presigned uploads were
  considered and rejected because they'd bypass the already-resolved
  magic-byte validation requirement (`kencleng-phase0-detail.md` Fitur
  7); file sizes here are small enough (5 MB max) that presigned
  upload's main benefit (offloading large-file bandwidth from the app
  server) doesn't apply. Presigned URLs are still used for **download**
  of private-bucket files, where this trade-off doesn't arise.

**Domain-by-domain generation order** (dependency order, matching Step
4's boundary): `account` → `organisasi` → `campaign` → `donation` →
`disbursement` → `notification`. Endpoint count per domain: account 20,
organisasi 15, campaign 19 (includes the one deliberate composite
endpoint, `GET /campaigns/{campaignId}`, aggregating campaign +
organisasi + progress per the design philosophy in
`kencleng-backend-tech-stack.md`), donation 6, disbursement 15,
notification 3.

### Step 10 — Frontend design guidelines (2026-07-27)

**Output**: `kencleng-design-guidelines.md` — full visual design token
set: color palette (primary green, distinct-shade success green,
orange warning, amber accent, cool-gray neutrals — each as a 50–900
scale), typography (Plus Jakarta Sans heading / Inter body, full type
scale), border radius ("rounded jelas" — pronounced rounding, 8–24px
scale), soft-shadow elevation scale, Lucide icon set, and
component-level tokens for buttons, inputs, badges (mapped to ERD
status enums), the donation progress bar, `MaskedField`,
`CurationDecisionPanel`, and `SecureUploadNote`.

**Key decisions:**
- **Brand mood**: warm & charitable (not corporate/fintech-cold).
- **Token implementation**: CSS custom properties as source of truth,
  referenced from Tailwind config (not duplicated) — keeps a single
  edit point and leaves room for a future `data-theme` override if
  dark mode is ever added.
- **Dark mode**: explicitly out of scope for v1 — no demonstrated
  need yet, consistent with the project's lowest-complexity principle.
- **Spacing scale**: Tailwind's default 4px-based scale, unmodified —
  no project-specific override introduced without a concrete gap.
- **Two resolved brand/semantic color conflicts**: primary green vs.
  success green, and accent amber vs. warning — both resolved by
  shifting hue/shade apart (success uses a more blue-leaning green;
  warning uses a redder orange) rather than reusing one color for
  both a brand/accent role and a semantic-state role.
- **Fonts**: self-hosted via `next/font/google` at build time (not a
  runtime Google Fonts request), keeping the PWA's offline/App-Shell
  properties intact.
- **Accessibility**: WCAG AA contrast minimums baked into shade
  choices (e.g. `primary-600`, not `500`, is the button background);
  focus-visible-only ring styling; status conveyed via color + label
  together, never color alone.

---

## Development Phase — Kickoff Decisions (2026-08-04)

Not a documentation/spec-phase step — recorded here because it's the
first decision made while transitioning into development, before a
separate development-phase tracker exists.

**Domain-by-domain *coding* order**:
```
account → notification → organisasi → campaign → donation → disbursement
```

This is **not the same order** as the "Domain-by-domain generation
order" used for authoring `api/openapi.yaml` at Step 9 (`account` →
`organisasi` → `campaign` → `donation` → `disbursement` →
`notification`, pure dependency order). That order was fine for
writing an API spec, since nothing there executes. For actual
development, the order was re-checked against the phase/persona
journeys in `kencleng-business-process-overview.md`, which surfaced
one adjustment worth making before code exists: `notification` is
cross-cutting and referenced from Phase 1 onward, so it's promoted to
2nd (right after `account`) rather than built last — cheap to do
early since it's Tier 3 (fully agentic, minimal review).

Two related process gaps identified in the same session, resolved by
extending existing templates rather than new documents:
- **Audit log write-sites** are scattered across `organisasi`,
  `campaign`, and `disbursement` features (built long after `account`,
  which owns the audit log table) — the `spec/<domain>/features/<fitur>.md`
  template now has a mandatory "Audit log entry?" field.
- **Cross-domain invariants** (e.g. the donation ledger vs.
  `campaign.collected_amount`) are now explicitly owned by whichever
  domain owns the underlying field, referenced (not duplicated) from
  the other domain's invariants doc.

Full rationale and template changes: `kencleng-agentic-workflow.md`
§3.3, §5.1, §5.3.

---

## Open Items (deferred to development phase)

**10 endpoint-shape decisions inferred during OpenAPI authoring**,
marked `# INFERRED` inline in `api/openapi.yaml`. These are proposals
consistent with patterns established elsewhere in the spec, not
formally resolved decisions — Anhar will confirm/adjust each against
real implementation constraints during the development phase rather
than in a further documentation session. Listed here for traceability:

1. **Login + MFA two-step contract** (`account`) — `/auth/login`
   returns an `mfa_pending_token` when MFA is enrolled, completed via a
   separate `/auth/login/mfa` call, rather than accepting an optional
   TOTP code directly in the single login request.
2. **Google OAuth link/reauth reuse the same redirect/callback
   endpoints** as login/register (`account`), branched via an `intent`
   query param, rather than separate dedicated endpoints per intent.
3. **Organisasi creation as a single `multipart/form-data` request**
   (`organisasi`) — base fields + legal document files submitted
   together in one call, rather than create-then-attach-separately.
4. **Legal/identitas field-edit confirmation mechanism** (`organisasi`)
   — a `confirm: true` flag in the request body/form, rejected with 409
   if a legal field changes without it, as the API-level counterpart to
   the FE confirmation dialog.
5. **`GET /organisasi` scope defaults to "mine"** (`organisasi`) — the
   organisasi the current user represents, not a public directory or
   admin-wide listing (those are served by separate endpoints).
6. **Publish/schedule/reschedule/republish unified into one endpoint**
   (`campaign`) — `POST /campaigns/{id}/publish`, behavior branching on
   current `status` and whether `publish_at` is provided, rather than
   four separate endpoints.
7. **Free-text search (`q`) on the public campaign listing**
   (`campaign`) — not specified in any phase doc; proposed for the
   public browse page's search box.
8. **`GET /organisasi/{id}/events`** (`campaign`) — an events-by-org
   listing endpoint for the Owner/Staff dashboard, not explicitly
   specified.
9. **`GET /account/donations`** (`donation`) — the registered user's
   own full donation history, implied by the Donatur persona but not
   explicitly specified as an endpoint.
10. *(Documented for completeness — `disbursement` and `notification`
    domains required no material inferences; every endpoint in those
    two domains maps directly to an explicit phase-doc business rule.)*

None of the above block starting development — each has a concrete
default in the spec today, and can be revised in `api/openapi.yaml`
without cascading changes elsewhere if Anhar's implementation
experience suggests a different shape.

---

## Removed From Scope

The following were previously tracked in this roadmap but have been
removed per the documentation-phase-only scope note. Recorded here for
history — pick these back up in whatever tracks the actual development
phase.

- **Project skeleton setup (backend & frontend)** — init Go module +
  Next.js app, base router/middleware skeleton, Zustand/TanStack Query
  provider setup, manual PWA scaffolding, hi-fi reference mockups.
  Output was "two bootable repos" — an implementation artifact, not a
  spec.
- **Local deployment target** — `docker-compose.yml` for local
  Postgres/MinIO/backend/frontend, now informed by the resolved
  same-origin reverse-proxy topology. An implementation/config
  artifact.
- **First vertical slice (Registrasi & Login) → subsequent slices** —
  actual feature code across auth, organisasi, campaign lifecycle,
  donation flow, and post-campaign reporting. Pure development work —
  fully unblocked by any outstanding spec question, including the
  now-authored `api/openapi.yaml`.

---

## Related Docs

- Actors, roles, entities: `kencleng-actors-entities.md`
- Backend/frontend tech stack: `kencleng-backend-tech-stack.md`,
  `kencleng-frontend-tech-stack.md`
- Business process overview: `kencleng-business-process-overview.md`
- Detailed feature specs: `kencleng-phase0-detail.md` through
  `kencleng-phase3-detail.md`
- UX page map: `kencleng-ux-page-map.md`
- ERD: `kencleng-erd.md`
- API contract: `api/openapi.yaml` (Step 9 output)
- **Frontend design guidelines: `kencleng-design-guidelines.md`**
  (new — Step 10 output)