# Kencleng — Backend Tech Stack
 
> Status: Draft — stack level decided, implementation details pending.
> Last updated: 2026-07-24
 
## Context
 
Kencleng is a sandbox donation-platform project built to learn Go and PWA
development through a realistic, non-trivial use case. It is not intended
for production or commercial use, and will eventually be published
open-source on GitHub.
 
Guiding principles for all decisions below:
- Start from the lowest complexity that solves the problem; add
  complexity only when there's a genuine, demonstrated need.
- Prefer language/platform idioms and stdlib over external dependencies
  unless there's a clear reason.
- Prioritize clarity and correctness over cleverness.
- Code should be written as if another developer will read and learn
  from it (public GitHub repo).
## Decided: Backend Stack
 
| Component | Choice | Rationale (short) |
|---|---|---|
| Language | Go | Core learning goal |
| HTTP Router | `net/http` stdlib (Go 1.22+ pattern routing) | Middleware, rate limiting, and grouping are all achievable via stdlib `http.Handler` wrapping — no framework-exclusive feature needed yet. Chi/Gin remain an easy migration path later if routing complexity genuinely grows. |
| Database | PostgreSQL | Production-like RDBMS for a financial/donation use case |
| DB Driver | `pgx` | Modern, performant, de-facto standard driver for Postgres in Go |
| Query Layer | `goqu` | Query builder (not an ORM) — programmatic query construction, auto-parameterized (SQL-injection-safe by design), no hidden relational magic. Developer already has prior familiarity. |
| Migration Tool | `golang-migrate` | Standard, explicit up/down `.sql` migrations. **Execution: CLI, run manually during dev** [RESOLVED — Step 4] — see Open Items #2 below. |
| Auth Mechanism | Access token (JWT, ES256, **TTL 15 minutes** [RESOLVED]) + Refresh token (HttpOnly cookie, **TTL 30 days** [RESOLVED], rotate-on-use + reuse detection) | Single web client for now — refresh token in HttpOnly cookie mitigates XSS token theft; JWT access token kept short-lived and stateless. ES256 chosen over HS256 for future-proofing (asymmetric — other services could verify tokens without holding the signing key). Full flow detail in `kencleng-phase0-detail.md`. |
| Identity Model | `User` (profile) separated from `AuthIdentity` (per login-method record) | Supports multiple identity providers per user without a rewrite. v1 implements `email_password` AND `google` providers; phone+OTP is modeled but deferred to a later version. See `kencleng-phase0-detail.md`. |
| Password Policy | **Length-only, min 8 characters, no complexity requirement** [RESOLVED — NEW] | Follows NIST 800-63B-style guidance: length contributes more to entropy than forced character-class rules, which tend to push users toward predictable patterns ("Password1!"). |
| Breach-List Check | **`pwnedpasswords.com` (HaveIBeenPwned) API, k-anonymity model, fail-open on API failure** [RESOLVED — NEW] | Included as an explicit learning goal (external API integration), not just a default. Only the first 5 characters of the SHA-1 hash of the candidate password are sent — the plaintext password and full hash never leave the server. Checked at registration, password reset, and set-password (Google-only users). If the API is unreachable, the flow proceeds without the check (logged) rather than blocking — this is a defense-in-depth layer, not the primary defense, and availability of core auth flows shouldn't depend on a third-party API's uptime. See `kencleng-phase0-detail.md` Fitur 1. |
| MFA | `pquerna/otp` (TOTP, RFC 6238) | Lightweight, standard-compliant, no framework overhead. Optional for all roles in v1. |
| File Storage | MinIO (S3-compatible) | Used for organization legal documents, campaign media, and fund-usage-report attachments. Public bucket for campaign media, private bucket + signed URLs for sensitive documents. Max file size **5 MB** across all contexts (legal docs, campaign media, fund-usage attachments); signed URL expiry **5 minutes** [RESOLVED — NEW]. |
| OAuth | `golang.org/x/oauth2` + Google's `idtoken` verification (`google.golang.org/api/idtoken` or manual JWKS verify) | "Login/Register dengan Google" is a v1-required feature. Official Go extended package, avoids pulling in a heavier third-party OAuth framework. Activates Fitur 4 (Account Linking) in `kencleng-phase0-detail.md`. `state` + `nonce` CSRF/replay protection detail: see Open Items #1 below. |
| Config Management | `godotenv` | Simple `.env` loading, no need for heavier config libs (e.g. viper) at this scale |
| Testing | `go test` stdlib (+ `net/http/httptest`) | `testify` and coverage tooling to be added only if/when assertion verbosity becomes a real pain point |
| Architecture | Domain-driven monolith | Single deployable unit, internally organized by domain boundary rather than technical layer-first. **Domains resolved** [RESOLVED — Step 4]: `account`, `organization`, `campaign`, `donation`, `disbursement`, `notification`. Folder convention: flat package per domain (`internal/domain/<name>/{entity.go, repository.go, service.go}`), transport separated in `internal/transport/http/`, shared infra in `internal/platform/`. Full rationale in Open Items #3 below and `kencleng-roadmap-next-steps.md` Step 4. |
| Error Handling | No panics for expected error paths — all errors explicitly returned | Correctness-critical for a financial flow; panics only for truly unrecoverable programmer errors |
| Rate Limiting | `golang.org/x/time/rate` | Applied globally from the start; per-endpoint overrides (e.g. stricter limit on `/auth/*`, including OAuth callback) to be layered in later. Complemented by the persistent `login_attempts` lockout (**5 failed attempts / 15 minute window** [RESOLVED — NEW]) for brute-force protection that survives app restarts — see `kencleng-phase0-detail.md` Fitur 2C. |
| Background Jobs / Scheduler | In-process scheduler (e.g. simple ticker-based goroutine) | Used for: campaign deadline check (Phase 2), scheduled publish (Phase 1), notification hard-delete (**weekly** [RESOLVED — NEW] — see phase0-detail.md Fitur 6). No external job queue needed yet — consistent with "start at lowest complexity." |
| Logging / Observability | *Deferred* | Explicitly postponed to avoid premature over-engineering; to be revisited once core flow exists |
| Cache (Redis) | *Deferred — as-needed* | Not adopted by default. Candidate triggers: shared cache across multiple instances, need for out-of-process persistence, or explicit intent to learn Redis itself. In-process caching (e.g. simple map+mutex) is the fallback for single-instance sandbox use. |
| Event-driven / Kafka | *Deferred — as-needed* | Kencleng starts as a monolith; in-process function calls or a Postgres outbox/`LISTEN-NOTIFY` pattern are the default fallback for anything looking like "events." Kafka would only be introduced for genuine multi-consumer/async/replay needs, or as an explicit learning goal in itself. |
 
## API Contract & Codegen [RESOLVED — Step 2]
 
**Format**: OpenAPI 3.x, spec-first — `api/openapi.yaml` is the single
source of truth for all HTTP endpoints, hand-authored *before*
implementation (not generated from code).
 
**Design philosophy** (agreed 2026-07-20, prior to the format decision
above): domain/resource-driven REST as the default endpoint shape,
with a justified composite-endpoint exception for `/campaign/[id]` —
this page aggregates campaign + organization + progress data in one call
to avoid client-side N+1 fetches on what's effectively the platform's
primary conversion surface.
 
**Backend usage — documentation/contract only, no codegen.** Handlers,
request/response structs, and validation in
`internal/transport/http/` are 100% hand-written. The spec is kept in
sync by developer discipline (updating `openapi.yaml` as part of the
same change as the handler), not by tooling.
- Rejected: `oapi-codegen` (or similar Go server-stub/type generation
  from the spec). Would guarantee spec/implementation parity at
  compile time, but adds a build step and a generated-code layer that
  isn't justified yet for a single-developer sandbox.
- **Accepted trade-off**: no compiler-enforced guarantee that
  `openapi.yaml` matches the actual Go implementation — accuracy
  depends entirely on discipline. Fine for a solo learning project;
  would need revisiting (via `oapi-codegen` or contract testing) if
  this ever became multi-developer or long-lived in production.
**Frontend usage — types only.** Request/response types generated via
`openapi-typescript` from `api/openapi.yaml`; fetch functions in
`lib/api/` stay hand-written, just typed against the generated
interfaces. See `kencleng-frontend-tech-stack.md` structural
principles for the full FE-side detail.
 
## Encryption Key Management [RESOLVED — NEW]
 
Context: `kencleng-erd.md` (Step 3) already finalized the
PII-encryption-at-rest *pattern* per field — `{field}` (`BYTEA`,
AES-GCM ciphertext) + `{field}_hash` (`TEXT`, HMAC-SHA256, for
lookup/uniqueness). What was still open was the key management around
that pattern: how many keys, where they live, and what happens if one
needs to change.
 
**Number of keys: 2, kept separate.**
- `ENCRYPTION_KEY` — used for AES-GCM encryption of
  `users.primary_email`, `auth_identities.identifier`,
  `organizations.npwp`, `donations.guest_email`
- `HMAC_KEY` — used for the corresponding `*_hash` columns
Using two independent keys (rather than deriving both from one key,
or reusing one key for both purposes) follows standard cryptographic
practice: never reuse key material across different
algorithms/purposes. This is treated as part of the project's explicit
secure-by-design learning goal, not just caution for its own sake.
 
**Storage: plain env vars, consistent with the existing `godotenv`
setup.**
- `ENCRYPTION_KEY` — 32 random bytes, base64-encoded
- `HMAC_KEY` — 32 random bytes, base64-encoded
- Generated once per environment (e.g. `openssl rand -base64 32`),
  documented as required placeholders in `.env.example`
**Rotation strategy: none in v1** — rejected adding a `key_version`
column to every encrypted field/table.
- **Considered and rejected**: adding `key_version` now to make future
  rotation easier. Rejected because it would require reopening
  `kencleng-erd.md`'s ERD (Step 3), which is already ✅ Done, to add
  columns to `users`, `auth_identities`, `organizations`, and
  `donations` — purely speculative, with no demonstrated need for
  rotation yet. Consistent with why Redis/Kafka stay deferred: don't
  add complexity ahead of a genuine requirement.
- **Accepted trade-off**: if a key ever needs to change (e.g.
  suspected compromise), it requires a manual one-off script —
  decrypt all affected rows with the old key, re-encrypt with the
  new one — rather than a built-in rotation mechanism. Acceptable for
  a solo sandbox project; would need real rotation support if this
  became a long-lived production system.
## Open Items — Needs Further Discussion
 
These don't block recording the stack above, but need their own deep-dive
before implementation:
 
1. **Auth details**
   - ~~JWT signing algorithm~~ → **resolved**: ES256
   - ~~Refresh token rotation strategy~~ → **resolved**: rotate-on-use
     + reuse detection (entire token family revoked on reuse)
   - ~~Refresh token server-side storage~~ → **resolved**: DB table
     (`RefreshToken`), not Redis — consistent with the cache decision
     below (Redis stays deferred until a genuine need appears)
   - ~~Google OAuth~~ → **resolved: promoted to in-scope v1**
   - ~~Access token TTL~~ → **resolved: 15 minutes**
   - ~~Refresh token TTL~~ → **resolved: 30 days**
   - ~~Password policy detail~~ → **resolved: length-only, min 8
     characters, no complexity requirement** — see
     `kencleng-phase0-detail.md` Fitur 1
   - ~~Email verification token expiry~~ → **resolved: 24 hours** —
     see `kencleng-phase0-detail.md` Fitur 1
   - ~~Google OAuth state/nonce validation detail (CSRF protection for
     the OAuth redirect flow), and how strictly redirect URI is
     validated~~ → **resolved**: `state` and `nonce` (random,
     single-use) stored together in a short-TTL (~10 min) HttpOnly
     cookie set before the redirect to Google — no new DB
     table/Redis needed, reusing the existing HttpOnly-cookie
     infrastructure already in place for refresh tokens. `state` is
     validated on callback (CSRF protection); `nonce` is validated
     against the claim inside the verified `id_token` (replay
     protection). `redirect_uri` is exact-match against
     `GOOGLE_REDIRECT_URI` registered per-environment in Google
     Console — no wildcard/pattern matching, and no dynamic
     redirect-URI handling on the app side (it's a fixed env var), so
     there's no additional open-redirect surface beyond Google's own
     whitelist enforcement. Full flow detail in
     `kencleng-phase0-detail.md` Fitur 1B. **[RESOLVED — NEW]**
2. ~~Migration execution~~ → **resolved: `golang-migrate` CLI, run
   manually during dev** — no embed/auto-run on app start considered
   and rejected. Trade-off accepted: slightly more onboarding friction
   for a fresh clone (must run the CLI step explicitly) in exchange
   for app startup staying free of DB-write side effects and explicit
   control over when migrations run. **[RESOLVED — Step 4]**
3. ~~Domain-driven monolith boundaries~~ → **resolved**: 6 domains —
   `account`, `organization`, `campaign`, `donation`, `disbursement`,
   `notification`. See `kencleng-roadmap-next-steps.md` Step 4 for the
   full breakdown and rationale (notably: `auth` and `user` were
   considered as separate domains but merged into one `account` domain
   due to tight transactional coupling between register/login flows;
   curation/review-assignment tables live inside each reviewed domain
   — `organization`, `campaign`, `disbursement` — rather than a separate
   shared `curation` domain, since the review process differs enough
   per context that a shared abstraction wasn't worth the cross-domain
   coupling; duplication of recusal/conflict-of-interest logic across
   the three is an accepted trade-off of domain-driven monolith
   design). File storage and audit logging are **not** domains — file
   storage is a shared `internal/platform/storage` package, and each
   domain owns and writes to its own `*_logs` table directly.
   - Folder/package convention → **resolved: flat package per
     domain**, `internal/domain/<name>/{entity.go, repository.go,
     service.go}`. Rejected subpackage-per-layer (unnecessary nesting
     for small domains like `notification`) and hexagonal ports &
     adapters (unjustified abstraction with only one concrete storage
     implementation and no swap planned). Transport (HTTP handlers)
     lives separately in `internal/transport/http/<name>_handler.go`
     so domain logic stays testable without `net/http`. Shared infra
     (MinIO client, pgx pool setup) lives in `internal/platform/`.
     **[RESOLVED — Step 4]**
4. **Notification expiration mechanics**
   - ~~Expiration duration~~ → **resolved: 30 days**
   - ~~Hard-delete worker frequency~~ → **resolved: weekly**
## Not Yet Discussed
 
- ~~API contract format (REST plain JSON vs OpenAPI spec-first)~~ →
  see **API Contract & Codegen** section above **[RESOLVED — Step 2]**
- Deployment/hosting target for the sandbox (Docker Compose local vs cloud)
## New Env Vars
 
Required for Google OAuth (in addition to whatever `.env` vars already
cover DB connection, JWT signing keys, etc.):
 
- `GOOGLE_CLIENT_ID`
- `GOOGLE_CLIENT_SECRET`
- `GOOGLE_REDIRECT_URI`
Required for PII encryption-at-rest [RESOLVED — NEW, see Encryption
Key Management above]:
 
- `ENCRYPTION_KEY` — 32 random bytes, base64-encoded (AES-GCM)
- `HMAC_KEY` — 32 random bytes, base64-encoded (HMAC-SHA256)
No new env vars needed for the HaveIBeenPwned integration — it's a
plain outbound HTTPS call to a public endpoint, no API key required.
 
## Explicitly Deferred / Out of Scope (for now)
 
These were raised and consciously *not* adopted yet, pending genuine need:
 
- **Redis** — no confirmed use case yet (candidates: read-heavy cache, distributed rate-limit counters, refresh-token store)
- **Kafka** — no confirmed event to decouple yet; revisit once a concrete async/cross-service need appears
- **Structured logging / observability stack** — deferred to a dedicated discussion
- **Phone + OTP identity provider** — modeled in `AuthIdentity.provider_type` but not implemented in v1
- **Encryption key rotation mechanism** — see Encryption Key Management above; no `key_version` scheme in v1