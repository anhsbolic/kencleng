# Tech Plan: Account Profile Read — GET /account/me (account #07)

> Ticket    : account domain task #7 — `docs/spec/1-account/features/07-account-profile.md`
> Author    : ox-alpha (agent) — for Anhar's review
> Date      : 2026-09-01
> Updated   : 2026-09-01 — contract locked; Open Item 1 dispositioned (fix-in-slice, Anhar 2026-09-01)
> Status    : Approved by Anhar
> Approach  : Vertical slice exposing the existing `GetLoginUserView` read model through a new
>             session-guarded handler, plus a shared transport-owned snake_case User mapping that
>             also fixes `/auth/login`'s pre-existing camelCase user keys (zero new SQL)
> Refs      : root + backend `AGENTS.md`; exploration logs
>             `.local-agents/works/account/07-account-profile/1-exploration/logs/` (8 files);
>             `api/openapi/{index,common,account}.yaml` + generated `api/openapi.yaml`;
>             `docs/spec/1-account/{invariants,threat-model,tasks}.md`; precedent
>             `.local-agents/works/account/06-mfa-totp/2-plan/techplan.md`; best-practices applied:
>             `go/authorization-and-idor`, `go/secrets-and-sensitive-logging`,
>             `go/jwt-and-token-lifecycle`, `postgresql/encryption-at-rest`,
>             `restapi/openapi-spec-first-drift`, `go/nil-and-zero-values`,
>             `go/abstraction-boundaries`, `restapi/pagination-and-status-codes`,
>             `restapi/idempotency-and-versioning`, `go/integration-testing-setup` (examples entry)

---

## 📋 Summary — start here

**What & why** — `GET /account/me` is the account domain's one missing read surface: the profile
read model (`LoginUserView`, assembled by `GetLoginUserView` with the domain's single decrypt-on-
read path) already exists and is integration-tested, but nothing serves it outside the login
response. This slice wires it to a new session-guarded handler — no ID parameter, resource keyed
entirely by the session (no IDOR surface per threat-model component 5), no audit entry, no
invariant exercised (Tier 2 pure read). The one substantive finding from exploration: the
`LoginUserView` struct has no JSON tags, so the login endpoint's `user` object already violates
the OpenAPI `User` schema's snake_case keys today — this slice makes `GET /account/me` correct
from birth via a transport-owned response struct, and — per your 2026-09-01 disposition —
routes `/auth/login`'s `user` object through the same mapping, eliminating the pre-existing
camelCase drift in the same change.

**Scope**
- New `GET /account/me` handler behind the existing `RequireSession` + per-IP rate-limit chain, on the main mux
- Thin `Service.GetProfile` pass-through to the existing `Repository.GetLoginUserView` (no new SQL, no new repo method)
- Transport-owned `userResponse` struct with exact snake_case keys (`id`, `name`, `email`, `email_verified`, `roles`, `auth_providers`, `mfa_enabled`, `created_at`) mapping from `LoginUserView`
- 401 (not 404/500) when the session's user no longer exists — the spec's only documented error case
- `/auth/login` + `/auth/login/mfa` success bodies: `user` object mapped through the same `userResponse` struct (wire-shape fix, zero clients existed at disposition)
- Key-shape assertions added to the existing login handler tests (both success writers)
- Tests: two spec-named tests (`TestAccountMe_RequiresAuth`, `TestAccountMe_NoIDParameter_SessionScoped`) + shape/edge/error/log-hygiene suite + a domain pass-through test
- No openapi changes (sources and bundle already document the endpoint; verified in sync)
- No migrations, no new dependencies, no new env vars

**Key decisions**
- Transport-layer `userResponse` struct with snake_case tags, not JSON tags on the domain entity — serialization is a transport concern; the domain entity stays tag-free (D2)
- Handler gets its own narrow `profileService` seam (one method), matching the per-file seam convention (`loginSessionService`, `googleOAuthService`, `securityService`) (D3)
- Route lands directly on the main mux with its own `RateLimit + RequireSession` chain — no `/account/` sub-mux restructure (D4)
- Deleted-user session → `401`, not 404 — the spec's error table documents only 401, and 401 matches session-scoped semantics without leaking account-state (D5)
- `roles`/`auth_providers` guaranteed `[]` (never `null`) at the transport boundary — deliberate, transport-owned normalization per the nil-vs-empty serialization rule (D7)
- `GetProfile` lives in a new `profile.go` (per-flow file convention), superseding the exploration sketch's `service.go` placement (D8)
- `/auth/login`'s `user` keys fixed in this slice via the shared `userResponse` mapping (D10, Anhar's 2026-09-01 disposition) — entity stays tag-free, wire delta flagged in the PR description

**Top risks** — none High-severity. The risk table (section 7) tops out at Medium (the pre-existing login-endpoint camelCase drift, now eliminated in-slice by D10; it remains Medium only if reverted); a Tier 2 pure read has no concurrency, no expensive primitives, and no new security boundary — the auth boundary it sits behind is the already-proven `RequireSession` middleware, re-used verbatim.

**Open items needing human input** — none. The one raised item (login `user`-key serialization) is Resolved in section 14 — fix-in-slice accepted per your 2026-09-01 disposition (D10). Remaining action is process-level: your approval of this plan.

---

<!-- Audience boundary: above is the human-readable digest for
review/approval. Below is the full execution-grade plan — same
decisions, same scope, expanded to file/line precision, rule IDs, and
full option comparisons. Nothing below contradicts the digest above;
it's the same source of truth at higher resolution. -->

---

## 1. Background

`GET /account/me` is specified in `docs/spec/1-account/features/07-account-profile.md` (Tier 2,
"pure read, no mutation, no invariant involved, standard OpenAPI-derived DTO" per
`tasks.md:44`): given a valid access token it returns the caller's own `User` resource —
`id`, `name`, `email` (decrypted — the `MaskedField` masking concern applies only to *other*
users' PII per `kencleng-actors-entities.md`), `email_verified`, `roles`, `auth_providers`,
`mfa_enabled`, `created_at`; with no/expired/invalid token it is standard-auth-middleware `401`;
there is no request body or parameter carrying an identifier, so there is no IDOR surface
(threat-model component 5, `threat-model.md:90`). No audit entry (a self-read is outside Fitur
9's change-scope), and no invariant from `docs/spec/1-account/invariants.md` is exercised.

Everything the endpoint needs at the data layer already exists: `Repository.GetLoginUserView`
(`repository.go:219`, implemented `repository_db.go:679`) assembles exactly this read model in
four goqu-parameterized reads — profile row with primary_email decrypted (the domain's one
decrypt-on-read path, `entity.go:110-117`), distinct provider types + the `email_verified`
predicate over `auth_identities`, roles from `user_roles` (empty until task #8), and the
`mfa_totp_secrets.enabled_at IS NOT NULL` count for `mfa_enabled`. It is integration-tested
(`repository_db_integration_test.go:955` `TestGetLoginUserView_AssemblesFields`,
`:1051` `TestGetLoginUserView_NonExistentUser`) and already consumed by login (`login.go:144,241`)
and MFA enroll (`mfa.go:75`). At the transport layer, `RequireSession` (`account_security.go:39`)
+ `UserIDFromContext` (`:23`) provide the proven session boundary, and every existing handler
follows the per-file narrow-seam convention.

The one real gap exploration surfaced is at the serialization boundary: `LoginUserView`
(`entity.go:118-127`) carries no JSON tags, so any handler writing it directly emits Go field
names. The OpenAPI `User` schema (`account.yaml:744-777`) specifies snake_case keys —
`email_verified`, `auth_providers`, `mfa_enabled`, `created_at`. This is not hypothetical:
`/auth/login` embeds `*account.LoginUserView` under `json:"user"` (`auth_login.go:41`) today,
and its own tests only assert `body["user"] != nil` (`auth_login_test.go:103`) — the live
response emits `ID`, `MFAEnabled`, `CreatedAt` camelCase keys, a pre-existing contract
violation of the same `User` schema object this endpoint returns. This slice fixes the shape
for `GET /account/me` at birth and — per the 2026-09-01 disposition (D10) — routes
`/auth/login`'s `user` object through the same mapping, eliminating the drift outright.

## 2. Scope

**In scope:**
- `internal/domain/account/profile.go` (new): `Service.GetProfile` pass-through method
- `internal/domain/account/service_test.go`: one pass-through unit test (reuses the existing
  `fakeRepo`, which already implements `GetLoginUserView` at `service_test.go:349`)
- `internal/transport/http/account_profile.go` (new): `profileService` seam, `userResponse`
  struct with snake_case tags, `toUserResponse` mapping (with nil→`[]` normalization),
  `AccountMeHandler`
- `internal/transport/http/account_profile_test.go` (new): handler contract suite (R1-R6)
- `internal/transport/http/auth_login.go`: `loginOKResponse.User` becomes `*userResponse`;
  both success writers (`LoginHandler`, `LoginMfaHandler`) map `res.User` through
  `toUserResponse`
- `internal/transport/http/auth_login_test.go`: key-shape assertions for both success writers
- `cmd/server/main.go`: one route registration — `GET /account/me` under
  `RateLimit` + `RequireSession(googleVerifyToken)` on the main mux

**Out of scope (explicit):**
- Any new repository method, SQL, or migration — `GetLoginUserView` already covers everything
- New integration tests — the repo method's integration coverage already exists (see §12 note)
- openapi source/bundle changes — endpoint already documented in both, verified in sync; the
  login fix changes the wire *toward* the already-documented schema, so no source edit needed
  there either
- `LoginUserView` JSON tags in `entity.go` — rejected (D2)
- Role assignment (#8), frontend work, root infra (Caddyfile prefix gap), `docs/spec/**` edits
  (agent-edit prohibited, root AGENTS §4), Makefile changes (its missing integration target is
  a pre-existing observation, noted in §11)

## 3. Requirements

| Condition | Requirement | Source/Note |
|---|---|---|
| Valid access token | `200` with the caller's own `User` resource: `id`, `name`, `email` (decrypted), `email_verified`, `roles`, `auth_providers`, `mfa_enabled`, `created_at` — exact snake_case keys per the `User` schema | spec AC (lines 14-20); `account.yaml:744-777` |
| No/expired/invalid access token | `401` Problem Details — standard `RequireSession` behavior, nothing endpoint-specific | spec AC (lines 21-22); `openapi.yaml:437-438` |
| Request carries any identifier (query/path/body) | Ignored — resource is keyed entirely by the authenticated session; no IDOR surface | spec AC (lines 23-26); threat-model component 5 |
| Session's user no longer exists (`GetLoginUserView` → `(nil, nil)`) | `401` — the spec's error table documents only 401; no 404/500 deviation, no partial-shape 200 | spec error table (lines 30-32); D5 |
| Internal failure (DB error, decrypt failure) | `500` generic Problem via `MapServiceError` default — detail never leaks internals | root AGENTS golden rule; `errors.go:130-136` |
| No audit log entry | Self-read is outside Fitur 9 scope (which covers *changes*) | spec §Audit (lines 61-67) |
| No invariant exercised | Pure read — verified against all 14 INV-account entries | spec §Applicable invariants (lines 34-37); invariants doc |
| `/auth/login` + `/auth/login/mfa` success bodies | `user` object carries the exact snake_case User keys via the shared mapping — no `LoginUserView` serialized raw on the wire anywhere | User schema `account.yaml:744-777`; Anhar disposition 2026-09-01 (§14 Resolved 1, D10) |
| Tier 2 gates | Automated contract-shape test + auth-middleware coverage suffice; no property/invariant test; human review = spot-check | spec §Risk tier (lines 48-55); `tasks.md:44` |

## 4. Rules & Validation

- **R1 (success shape — exact contract)**: Given an authenticated caller whose user exists,
  When `GET /account/me` is called, Then the response is `200` `application/json` with exactly
  the keys `id`, `name`, `email`, `email_verified`, `roles`, `auth_providers`, `mfa_enabled`,
  `created_at` (snake_case, no extras), `id` marshals as a UUID string, `created_at` as
  RFC 3339, and the values equal the seeded `LoginUserView`. *Test proves:
  `TestAccountMe_Success_UserShapeSnakeCase`.*
- **R2 (auth required — spec-named)**: Given no session user in context (as if
  `RequireSession` never ran or rejected), When the handler executes, Then it responds `401`
  with problem type `https://kencleng.dev/errors/unauthorized` before touching the service
  (stub records no call). Token-level failures (garbage/expired/wrong-key) are already proven
  by the existing `TestRequireSession_*` suite — re-run unchanged, no new middleware test.
  *Tests: `TestAccountMe_RequiresAuth` (spec-named), existing `TestRequireSession_*` re-run.*
- **R3 (session-scoped, no ID surface — spec-named)**: Given requests carrying foreign
  identifiers (`?user_id=<other-uuid>`, path suffixes, junk body), When the handler executes,
  Then the service receives the session's `userID` only and the response is the session user's
  own data; no parameter influences the resource. *Test: `TestAccountMe_NoIDParameter_SessionScoped`
  (spec-named; stub asserts the received `userID` equals the session's).*
- **R4 (session references a gone user)**: Given `GetProfile` returns `(nil, nil)`
  (user row absent), When the handler executes, Then the response is `401` with the same
  generic Unauthorized problem shape as R2 — never a 200 with a partial object, never 404,
  never a nil-deref. *Test: `TestAccountMe_UserGone_SessionInvalidated`.*
- **R5 (internal error → generic 500)**: Given `GetProfile` returns a wrapped error (DB
  failure), When the handler executes, Then `MapServiceError`'s default maps it to `500`
  problem type `https://kencleng.dev/problems/internal` with a generic detail — the error
  string never reaches the client. *Test: `TestAccountMe_InternalError_Generic500`.*
- **R6 (no PII in logs)**: Given a success request whose view carries a real email, When the
  handler completes, Then no log output contains the email or any response payload — the
  decrypted email exists only in the response body. *Test: `TestAccountMe_LogsFreeOfPII`
  (stderr-capture scan during a success request; mirrors task-#06 R15's technique).*
- **R7 (service pass-through contract)**: Given the domain seam, When `GetProfile` is called,
  Then it forwards the `userID` to `Repository.GetLoginUserView` unmodified and returns the
  view or `(nil, nil)` verbatim — no transformation, no error wrapping beyond what the
  repository already emits. *Test: `TestGetProfile_PassesThroughToRepository` (domain unit,
  existing `fakeRepo` + `seedView`).*
- **R8 (login/mfa-login user shape — dispositioned fix)**: Given either success writer
  (`LoginHandler` with status `"ok"`, `LoginMfaHandler` after a valid code), When the response
  is written, Then `body["user"]` carries exactly the same eight snake_case keys as R1 via the
  shared `userResponse` mapping — and no code path serializes `*account.LoginUserView` raw
  anymore. *Tests: `TestLoginHandler_UserShapeSnakeCase`,
  `TestLoginMfaHandler_UserShapeSnakeCase` (both in `auth_login_test.go`, asserting the same
  exact-key set as R1's helper).*

**Count-check**: R1–R8 — every rule ID has ≥1 corresponding item in section 12. ✓

## 5. Decision Log

### D1 — Service method (thin pass-through vs handler→repo vs no method)

| Option | Why rejected/accepted |
|---|---|
| A. Thin `Service.GetProfile` pass-through to `s.repo.GetLoginUserView` | **Chosen.** Preserves the transport↔domain seam convention every handler follows (handlers never touch repositories; services are where the view is fetched — `login.go:144,241`, `mfa.go:75`). Honest note: `abstraction-boundaries` §3 flags single-caller pass-throughs as speculative layering — the counter is that this isn't test-isolation indirection but the repo's own established seam convention, which per guardrails §4 outranks the generic guidance; the method also gives the future a named place for any profile-specific logic without touching the handler. |
| B. Handler calls the repository directly | Rejected — breaks the convention every existing handler follows; would require exposing a repo seam at transport for the first time in the codebase. |
| C. No service method; handler takes a `GetLoginUserView` closure | Rejected — same indirection cost as A with a less conventional shape (closures appear in this codebase only for token minting seams, not data reads). |

### D2 — Response struct location (transport struct vs domain tags vs raw struct)

| Option | Why rejected/accepted |
|---|---|
| A. Transport-owned `userResponse` struct with snake_case tags mapping from `LoginUserView` | **Chosen.** Serialization naming is an API concern, not a domain concern (`entity.go` doctrine keeps the domain shape transport-agnostic); the struct doubles as the explicit contract carrier the shape tests assert against — for `/account/me` and, since D10, for both login success writers. |
| B. Add JSON tags to `LoginUserView` in `entity.go` | Rejected — couples the domain entity to the wire contract and changes `/auth/login`'s shape via the domain layer; the shape fix belongs at the transport boundary (D10 chose the mapping route, keeping the entity tag-free). |
| C. Serialize `*account.LoginUserView` raw | Rejected — emits camelCase keys (`ID`, `MFAEnabled`, `CreatedAt`), violating the `User` schema; this is precisely the pre-existing drift D10 eliminates. |

### D3 — Transport seam (own `profileService` vs extend `securityService`)

| Option | Why rejected/accepted |
|---|---|
| A. New `profileService` interface (single method `GetProfile`) in the new handler file | **Chosen.** Matches the per-file seam convention: `loginSessionService` (`auth_login.go:17`), `googleOAuthService` (`auth_google.go:35`), `securityService` (`account_security.go:66`). One-method seams are the codebase norm for isolated flows. |
| B. Extend `securityService` with `GetProfile` | Rejected — mixes a profile read into the security-flow seam; `/account/me` is a different route group with a different concern, and the security seam's method set is already six methods. |

### D4 — Route wiring (main-mux route vs `/account/` sub-mux)

| Option | Why rejected/accepted |
|---|---|
| A. One route on the main mux with its own `RateLimit + RequireSession` chain | **Chosen.** The existing `accountMux` is mounted at `/account/security/` (`main.go:193`) — `/account/me` doesn't share the prefix, so a sub-mux would force restructuring working routes for no benefit; a single explicit chain mirrors how each route group declares its middleware today. Rate limit included for consistency with the security group (cheap, per-IP, eviction-checked — `rate-limiting` checklist satisfied by reuse). |
| B. New `/account/` sub-mux absorbing both `/account/me` and `/account/security/*` | Rejected — restructuring two proven route groups to save one line of middleware composition; pure churn risk for a read endpoint. |

### D5 — Gone-user status code (401 vs 404 vs 500 vs 200-partial)

| Option | Why rejected/accepted |
|---|---|
| A. `401` with the same generic Unauthorized problem shape as missing-token | **Chosen.** The spec's error table (`07-account-profile.md:30-32`) documents only `401`; semantically the session's subject no longer exists, so the token no longer identifies a valid account — the same class as "invalid credentials." Also avoids leaking account-state distinctions (`pagination-and-status-codes` §2 discipline: status reflects the actual failure category — here, "session not usable"). |
| B. `404` | Rejected — deviates from the documented contract; for a no-IDOR-surface endpoint it would additionally distinguish "token valid but user deleted" from "token invalid," an information leak with zero utility. |
| C. `500` | Rejected — a gone user is a client-side condition (stale token), not a server fault; 5xx would misdirect retries (`pagination-and-status-codes` §2). |
| D. `200` with a partial/zero-value object | Rejected — encodes a real failure as success; also the nil-deref hazard the check exists to prevent. |

### D6 — `id`/`created_at` serialization types

| Option | Why rejected/accepted |
|---|---|
| A. Keep `uuid.UUID` and `time.Time` in `userResponse`; let `encoding/json` marshal them (uuid via `MarshalText` → string; time → RFC 3339) | **Chosen.** Matches the `loginOKResponse` precedent (`auth_login.go:37-42`, which keeps `time.Time` for `access_token_expires_at`); zero conversion code; both defaults satisfy the schema's `format: uuid` / `format: date-time`. |
| B. Convert to `string` fields (`v.ID.String()`, `v.CreatedAt.Format(time.RFC3339)`) | Rejected — redundant conversion for identical wire output; the exploration sketch proposed it, but the precedent type-keeping is simpler and equally correct (sketch superseded, noted per rules § 2). |

### D7 — nil→`[]` normalization at the boundary

**Chosen:** `toUserResponse` normalizes nil `Roles`/`AuthProviders` to empty slices before
serialization. Today `GetLoginUserView` already initializes both as `[]string{}`
(`repository_db.go:682-683`), but the `[]`-not-`null` guarantee (schema: arrays; clients
distinguish "no roles" from malformed `null` — `nil-and-zero-values` §2) should be owned by
the transport boundary deliberately, not inherited as a repo-side accident that a future
query change could silently break. Cost: two lines. Test: R1 asserts both keys decode as
`[]any`, not `nil`.

### D8 — `GetProfile` file placement *(deviation from the stage-3 sketch, recorded)*

| Option | Why rejected/accepted |
|---|---|
| A. New `internal/domain/account/profile.go` | **Chosen.** The domain splits by flow (`login.go`, `security.go`, `mfa.go`, `google_oauth.go`); a profile read is its own (small) flow. Keeps `service.go` (already 790 lines of register/verify/reset flows) from absorbing an unrelated concern. Deviation from the stage-3 sketch's `service.go` noted here explicitly — the sketch is input, not gospel (task-#06 synthesis-note precedent). |
| B. Append to `service.go` | Rejected — `service.go` hosts the verification-flow cluster; a profile read is not part of that flow's cohesion, and the file is already the domain's largest. |

### D9 — Shape assertions in the regular unit suite (not a contract-tagged file)

**Chosen:** the exact-key shape tests live in the plain transport test suites so they run in
every `go test ./...` pass, not only under `-tags=contract`. Observation recorded (§11): the
`test-contract` gate currently has zero contract-tagged files in the backend, so it is
effectively a duplicate unit run — creating the first contract-tagged file in this slice
would be scope creep for no added strictness (plain-tagged is strictly stronger here).

### D10 — Login user-shape fix in this slice *(Open Item 1 disposition — accepted, Anhar 2026-09-01)*

| Option | Why rejected/accepted |
|---|---|
| A. Fix in-slice: map both success writers' `user` through the shared `userResponse` | **Chosen.** Zero clients exist (FE unbuilt) — the "fix now or never" window per `idempotency-and-versioning` §2; same User schema object as `/account/me`, so one shared mapper serves both; ~6 lines + two test additions; no domain change (entity stays tag-free, D2 upheld). Wire-shape change flagged in the PR description for standard review veto — the sign-off this row records. |
| B. Defer to its own slice | Rejected — leaves two shapes for one logical object; the deferred fix later requires coordinating an existing endpoint's shape with FE consumption simultaneously, strictly riskier than today's free fix. |
| C. Add JSON tags to `LoginUserView` (entity-level fix) | Rejected — couples the domain entity to the wire contract and changes the shape via the domain layer; D2's reasoning stands. |

## 6. Backward Compatibility

- **Database**: none — no DDL, no migration; the endpoint reads four existing tables
  (`users`, `auth_identities`, `user_roles`, `mfa_totp_secrets`) through the existing
  `GetLoginUserView`. No backfill, no data risk.
- **API**: one new endpoint, plus one deliberate wire-shape fix on an existing one:
  `/auth/login` and `/auth/login/mfa` success bodies' `user` object keys change from
  camelCase (`ID`, `MFAEnabled`, `CreatedAt`) to the User schema's snake_case
  (`id`, `mfa_enabled`, `created_at`) — classified per `idempotency-and-versioning`
  §2 as breaking-in-shape but zero-impact (no clients exist; FE unbuilt). Signed off as
  fix-in-slice (D10, Anhar 2026-09-01); the delta is flagged in the PR description for
  standard review veto. No other route, body, or status contract changes.
- **Existing clients/data**: none — no consumer exists yet for either endpoint's `User`
  object; fixing the shape now costs nothing and prevents the FE from baking in wrong keys.
- **Deprecation path**: none needed.

## 7. Edge Cases & Risks

| Risk | Likelihood | Severity | Mitigation |
|---|---|---|---|
| `/auth/login`'s `user` object emitting camelCase keys against the User schema | Was certain pre-slice (verified tag-free struct, unasserted tests) — eliminated by D10's in-slice fix | Medium *only if reverted/deferred* | Fixed by the shared `userResponse` mapping on both success writers; R8 proves both writers; regression risk is the ordinary "someone swaps the struct back" — guarded by the exact-key tests |
| Handler mishandles the `(nil, nil)` gone-user case (e.g. writes a 200 with a zero-value object, or nil-derefs in the mapper) | Low (explicit check is cheap) | Medium — wrong status class on a PII-bearing response; a panic would 500 with no handler output | Explicit nil-check before mapping (D5); R4 proves it |
| Torn multi-query read: the four `GetLoginUserView` queries run without a tx, so a concurrent role/MFA/linking mutation mid-read can produce a snapshot mixing pre/post states (e.g. `roles` fetched after a revoke, `mfa_enabled` before an enable) | Low | Low — accepted: no invariant covers read atomicity of this view; the login path already reads it identically; profile reads tolerate eventual consistency by design | Documented as accepted; no test — nothing to assert beyond "values are internally plausible," which no invariant demands |
| Future `phone_otp` identity rows would surface in `auth_providers`, violating the `ProviderType` enum (`email_password`, `google` only — `account.yaml:739-742`) | Low (no writer exists today: register writes `email_password`, OAuth writes `google`) | Low–Medium — latent, only if a phone-OTP feature lands without touching this contract | No mitigation now (would be speculative filtering); note recorded for the future phone-OTP work: either extend the schema enum or filter in `GetLoginUserView` — flagged, not silently buried |
| `userResponse` drifts from the `User` schema if the schema later gains/renames a field | Low | Low — one-file fix once noticed | Exact-key assertions in R1/R8 make drift fail loudly at the next `go test` (`openapi-spec-first-drift` checklist item 4); spec-first discipline applies to any future schema change |
| Decrypted email leaks into logs (the one PII in play) | Low | Medium — golden-rule violation if it happens | Handler never logs the view or response; R6 stderr-scan proves it |
| 401 problem-body detail wording differs from the openapi `Unauthorized` example (`"Authentication required."` vs example's Indonesian sentence) | Certain (pre-existing, all `/account/*` endpoints) | Low — cosmetic; the example detail is illustrative, the type/title/status are normative and match | Kept as-is for body consistency across endpoints; noted here so a reviewer doesn't read it as a new drift introduced by this slice |

## 8. Interface Contract

Repo conventions honored (root + backend AGENTS.md): SQL parameterized via goqu — **no new
SQL in this slice** (the endpoint reuses `GetLoginUserView` verbatim); client-facing errors
only via RFC 9457 Problem Details; PII policy — `email` is decrypted for the resource owner
only (the established single decrypt-on-read path, `entity.go:110-117`; no new key scheme, no
new HMAC column — lookups key on `user_id` PK); authorization explicit at the boundary —
`userID` derives from the session context (`UserIDFromContext`), never from a request
parameter; money fields: none; no PII or tokens in logs (R6).

**DB Schema changes:** none. (Tables read: `users(id, name, primary_email BYTEA,
primary_email_hash, created_at)`, `auth_identities(user_id, provider_type, verified_at)`,
`user_roles(user_id, role)`, `mfa_totp_secrets(user_id, enabled_at)` — all pre-existing,
migrations 000001-000008 territory, untouched.)

**API changes** (no openapi source/bundle edit — verified already documented and in sync,
`account.yaml:416-428` ≡ `openapi.yaml:425-438`; the login fix changes the wire *toward* the
already-documented schema, so no source edit needed there either):

```yaml
GET /account/me:                 # bearerAuth (RequireSession: cookie-or-bearer ES256)
  200: User{                     # exact wire keys this slice guarantees
    id: uuid                     # uuid.UUID via MarshalText
    name: string
    email: string                # DECRYPTED — owner-only view (the one decrypt path)
    email_verified: boolean      # any email_password identity verified
    roles: [admin|kurator]       # [] today until task #8
    auth_providers: [email_password|google]
    mfa_enabled: boolean         # mfa_totp_secrets.enabled_at IS NOT NULL
    created_at: date-time        # RFC 3339 via encoding/json default
  }
  401: Problem errors/unauthorized   # middleware (token issues) AND handler (gone user)

/auth/login + /auth/login/mfa (success "ok" only):
  user: User{...}                  # NOW mapped through userResponse — same eight snake_case
                                   # keys as /account/me (D10); camelCase emission eliminated
```

**Business logic flow (concise):**

```
AccountMeHandler(ctx):
  userID, ok := UserIDFromContext(ctx)        # session-derived; no request param consulted (R3)
  !ok       -> 401 Problem(unauthorized)      # R2 (mirrors every existing handler)
  view, err := svc.GetProfile(ctx, userID)    # thin pass-through -> repo.GetLoginUserView (R7)
  err != nil -> MapServiceError(w, err)       # wrapped DB/decrypt failure -> 500 generic (R5)
  view == nil -> 401 Problem(unauthorized)    # gone user, same generic shape (R4, D5)
  writeJSON(200, toUserResponse(view))        # snake_case keys; nil->[] normalized (R1, D7)
  # no logging of view/response anywhere (R6)

LoginHandler / LoginMfaHandler success writers (D10):
  User: toUserResponse(res.User)              # replaces the raw *LoginUserView field (R8)
```

No external calls, no transactions opened, no emails sent — nothing to keep out of a tx
because there is no tx.

## 9. Architecture / Plan

Linear CRUD-shaped flow — one line: session context → thin service pass-through → existing
four-query read model → transport-owned snake_case mapper → JSON. No diagram (guard-chain of
three trivial conditions; `examples.md` skip-diagram precedent; no state transition, no
cross-component sequencing).

Execution order:

1. **Domain**: `profile.go` — `GetProfile` + doc comment; unit test in `service_test.go`
   (existing `fakeRepo` already implements the method, `service_test.go:349`).
2. **Transport**: `account_profile.go` — seam, response struct, mapper, handler; contract
   suite in `account_profile_test.go` (reuses `injectSessionUserID`, `decodeBody`-style
   helpers per the existing test files' conventions).
3. **Wiring**: `main.go` — one `mux.Handle("GET /account/me", ...)` line beside the
   `/account/security/` mount (same chain composition).
4. **Login-shape fix (D10)**: `auth_login.go` — `loginOKResponse.User` → `*userResponse`,
   both success writers map `res.User` through `toUserResponse`; key-shape assertions added
   to both login success tests in `auth_login_test.go`.
5. **Gate**: `make verify` — no openapi step needed (no spec edit); no integration run needed
   beyond what the gate already covers.

Migration strategy: N/A — none.

## 10. Implementation Details

**File**: `backend/internal/domain/account/profile.go` (new)
- Change: `func (s *Service) GetProfile(ctx context.Context, userID uuid.UUID) (*LoginUserView, error)`
  — body is `return s.repo.GetLoginUserView(ctx, userID)` with a doc comment stating the
  owner-only decrypt semantics and the `(nil, nil)`-means-gone-user convention the transport
  maps to 401. One of the few places a full body is worth showing (it *is* the whole method):

```go
// GetProfile returns the resource owner's own profile view — the same
// LoginUserView login assembles, including the decrypted primary email
// (self-view only; the MaskedField concern applies to other users' PII,
// never this one). Returns (nil, nil) when the session's user no longer
// exists; the transport maps that to 401 (session no longer identifies a
// valid account).
func (s *Service) GetProfile(ctx context.Context, userID uuid.UUID) (*LoginUserView, error) {
	return s.repo.GetLoginUserView(ctx, userID)
}
```

**File**: `backend/internal/transport/http/account_profile.go` (new)
- Change: three pieces, following the house handler shape (`SetPasswordHandler` precedent,
  `account_security.go:90` — context extraction first, service call, mapping):
  - `type profileService interface { GetProfile(ctx context.Context, userID uuid.UUID) (*account.LoginUserView, error) }` — `*account.Service` satisfies it; tests stub it.
  - The contract-carrier struct (shown — it is the wire contract this slice exists to get right):

```go
// userResponse mirrors openapi components.schemas.User exactly — snake_case
// keys are deliberate (the User schema's contract), unlike the untagged
// LoginUserView whose Go field names would leak camelCase onto the wire.
// uuid.UUID marshals via MarshalText (uuid string) and time.Time via RFC
// 3339 — both schema formats satisfied by encoding/json defaults (D6).
type userResponse struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	Roles         []string  `json:"roles"`
	AuthProviders []string  `json:"auth_providers"`
	MFAEnabled    bool      `json:"mfa_enabled"`
	CreatedAt     time.Time `json:"created_at"`
}
```

  - `func toUserResponse(v *account.LoginUserView) userResponse` — field copy + nil→`[]string{}` normalization for `Roles`/`AuthProviders` (D7).
  - `func AccountMeHandler(svc profileService) http.HandlerFunc` — the four-step flow in §8; `writeJSON` reused from `auth_login.go:58`; 401 bodies via the established `WriteProblem` triple (`errors/unauthorized` / "Unauthorized" / "Authentication required.") identical to `account_security.go:94-97`.

**File**: `backend/internal/domain/account/service_test.go`
- Change: `TestGetProfile_PassesThroughToRepository` — seed the existing `fakeRepo` via its
  `seedView` helper (`service_test.go:359`), call `GetProfile`, assert the same pointer
  returns and the fake recorded the same `userID`; a second sub-case asserts the `(nil, nil)`
  pass-through. (Reuses the existing fake — no new fake needed.)

**File**: `backend/internal/transport/http/account_profile_test.go` (new)
- Change: `stubProfileService{view *account.LoginUserView; err error; gotUserID uuid.UUID}`
  implementing the seam; six tests per §12. Test helpers reused: `injectSessionUserID`
  (`account_security_test.go:410`) for session injection; `newES256Signer`
  (`auth_google_test.go:61`) not needed at handler level (session injection suffices — the
  token-level 401s stay covered by the existing `TestRequireSession_*` suite, R2).

**File**: `backend/internal/transport/http/auth_login.go`
- Change: `loginOKResponse.User` type becomes `*userResponse` (same `json:"user,omitempty"`
  outer tag); at both success writers (`:96-101`, `:134-139`) map `res.User` through
  `toUserResponse(res.User)` — two sites only (Google callback is a 302 flow, verified to
  never serialize a user object).

**File**: `backend/internal/transport/http/auth_login_test.go`
- Change: `TestLoginHandler_UserShapeSnakeCase` + `TestLoginMfaHandler_UserShapeSnakeCase` —
  seed the stub's `User` with a full view, decode `body["user"]`, assert the exact eight-key
  set (same assertion helper as R1). Existing tests (`:103`, `:229` vicinity) keep passing —
  they assert presence, not key names.

**File**: `backend/cmd/server/main.go`
- Change: one registration beside the security mount (lines 186-194):

```go
// Account profile read (task #07). Same middleware chain as the security
// group; /account/me is its own route on the main mux (prefix differs from
// /account/security/ — see D4).
mux.Handle("GET /account/me",
	transporthttp.RateLimit(rps, burst)(
		transporthttp.RequireSession(googleVerifyToken)(
			transporthttp.AccountMeHandler(accountSvc))))
```

Full function bodies deliberately omitted beyond the three shown (guardrail §7) — the handler
and mapper mirror existing precedents (`SetPasswordHandler`, `writeJSON`); their §8 flow block
is the authoritative behavior description.

## 11. Files Changed / Files NOT Changed

| File | Change Type | Description |
|---|---|---|
| `backend/internal/domain/account/profile.go` | new | `GetProfile` pass-through + doc comment |
| `backend/internal/domain/account/service_test.go` | modified | +`TestGetProfile_PassesThroughToRepository` (existing fakeRepo) |
| `backend/internal/transport/http/account_profile.go` | new | seam, `userResponse`, `toUserResponse`, `AccountMeHandler` |
| `backend/internal/transport/http/account_profile_test.go` | new | R1-R6 contract suite |
| `backend/internal/transport/http/auth_login.go` | modified | `loginOKResponse.User` → `*userResponse`; both success writers map through `toUserResponse` (D10) |
| `backend/internal/transport/http/auth_login_test.go` | modified | +2 user-key shape tests (R8) |
| `backend/cmd/server/main.go` | modified | +1 route registration |

| File | Reason untouched |
|---|---|
| `backend/internal/domain/account/repository.go`, `repository_db.go` | `GetLoginUserView` already on the port, implemented, integration-tested — zero new SQL |
| `backend/internal/domain/account/entity.go` | `LoginUserView` stays tag-free (D2); no new entity needed |
| `backend/internal/domain/account/{login,security,mfa,google_oauth}.go` | Existing consumers of `GetLoginUserView`/`RequireSession` — unchanged; MFA file untouched despite proximity |
| `backend/internal/transport/http/{account_security,auth_google,errors,cookie,middleware,swagger}.go` | Chain pieces reused as-is; `MapServiceError` needs no new sentinel (no new error vocabulary — default 500 case suffices) |
| `backend/internal/platform/**` | `auth` (Tier 0 fenced) verifier reused verbatim; `crypto` stays inside the repo adapter; nothing new anywhere |
| `backend/migrations/**` | No DDL |
| `api/openapi/{index,common,account}.yaml`, `api/openapi.yaml` | Endpoint + `User` schema already documented in both; verified in sync — no amendment, no bundle regen; the login fix moves the wire toward the documented schema |
| `backend/Makefile` | Untouched; observation recorded: verify chain is `lint → test-unit → test-race → test-contract → security` with **no integration target** and zero contract-tagged files (root AGENTS §5 lists "integration" in the gate order — the Makefile lags the doc; pre-existing, root-level concern, not this slice) |
| `backend/.env.example` | No new env vars |
| `backend/internal/domain/{donation,disbursement,notification,organisasi}/**` | Other domains; donation/disbursement additionally Tier 0 fenced |
| `docs/spec/**` | Agent-edit prohibited (root AGENTS §4) |
| `frontend/**` | Directory boundary (root AGENTS §7) |
| root `Caddyfile` | Known infra gap, root-level session |

## 12. Testing Checklist

Derived 1:1 from section 4 (count-check R1–R8 passed):

- [ ] R1 `TestAccountMe_Success_UserShapeSnakeCase` — 200; exact key set (`id`, `name`, `email`, `email_verified`, `roles`, `auth_providers`, `mfa_enabled`, `created_at`); `id` decodes as UUID string; `created_at` as RFC 3339; values match the seeded view; ⚠️ `roles`/`auth_providers` asserted as `[]any` (present, zero length) — not `nil` — proving the nil→`[]` guarantee (D7)
- [ ] R2 `TestAccountMe_RequiresAuth` *(spec-named)* — handler with no session user in context → `401`, problem type `errors/unauthorized`, stub records no service call; plus existing `TestRequireSession_*` suite re-run unchanged for token-level 401s (no new middleware test authored)
- [ ] R3 `TestAccountMe_NoIDParameter_SessionScoped` *(spec-named)* — requests with `?user_id=<foreign-uuid>`, junk path suffix, junk body: stub asserts the received `userID` equals the session's in every case; response is the session user's data
- [ ] R4 `TestAccountMe_UserGone_SessionInvalidated` — stub returns `(nil, nil)` → `401` with the same generic problem shape as R2 (byte-shape equality of type/title/status), no partial object, no 404/500
- [ ] R5 `TestAccountMe_InternalError_Generic500` — stub returns a wrapped sentinel-free error → `500`, problem type `problems/internal`, generic detail (no error text echoed)
- [ ] R6 `TestAccountMe_LogsFreeOfPII` — stderr captured during a success request with a real email in the seeded view; assert the email substring appears nowhere in captured log output (task-#06 R15 technique)
- [ ] R7 `TestGetProfile_PassesThroughToRepository` — domain unit: seeded `fakeRepo` view returned as the same pointer, `userID` forwarded unmodified; second sub-case asserts `(nil, nil)` pass-through
- [ ] R8 `TestLoginHandler_UserShapeSnakeCase` + `TestLoginMfaHandler_UserShapeSnakeCase` — both success writers emit the exact eight snake_case keys via the shared mapping; ⚠️ assert via the same key-set helper as R1 so the two endpoints can never silently diverge

Gate commands (backend Makefile, verified): `make verify` = `staticcheck ./... && gosec ./... && go test ./... && go test -race ./... && go test -tags=contract ./... && gitleaks detect && govulncheck ./...` exits 0. Tier 2 bar (spec §Risk tier): contract-shape test + auth coverage suffice — no property/invariant test required. ⚠️ No new integration test in this slice: `GetLoginUserView`'s integration coverage already exists and is unchanged (`TestGetLoginUserView_AssemblesFields`, `TestGetLoginUserView_NonExistentUser`, both `//go:build integration`).

### Test Focus Pointer (carry-over from exploration Risk lens)

| Area | Why sensitive | Still relevant post-synthesis? |
|---|---|---|
| Decrypted-PII email in the `User` response | PII disclosure surface — the response is decrypted email to the owner | N/A — owner-only by construction (spec lines 16-19; no masking applies to self-view); no race/concurrency/expensive-primitive class exists; the log-hygiene side is covered by ordinary unit test R6 — no testing-phase-only class needed |
| JSON serialization shape (exploration Areas 3+5 Risk finding) | Contract drift on a PII-bearing response; login endpoint already drifted (now fixed in-slice per D10) | N/A — contract-correctness class, fully covered by R1/R8's exact-key assertions inside the regular suite; not a shared-state/concurrency/security-boundary test class |

(Both exploration-flagged candidates recorded explicitly per guardrails §12 — neither survives as a testing-phase pointer; both reasons above.)

## 13. Testing Examples & Common Mistakes

| Mistake | Error/Behavior | Fix |
|---|---|---|
| Serializing `*account.LoginUserView` directly as the response body | CamelCase keys (`ID`, `MFAEnabled`, `CreatedAt`) on the wire — the exact violation this slice eliminates (R1/R8 fail immediately) | Always map through `userResponse`; treat the struct as the contract carrier |
| Mapping only one of the two success writers | `/auth/login` fixed, `/auth/login/mfa` still emits camelCase — the two login paths diverge silently | R8 asserts the same key set on both writers; shared mapper makes divergence compile-visible only if both sites call it — review checks both call sites |
| Skipping the nil-view check "because users are never deleted" | First hard-deleted user (or any future deletion path) turns into a 200-with-zero-values response or a panic in the mapper | The check is two lines; R4 pins it regardless of whether v1 ever deletes users |
| Asserting the response shape by string-containment (`strings.Contains`) | CamelCase regression can pass if e.g. `id` appears inside another value | Decode to `map[string]any` and assert the exact key set + types (R1's method) |
| Asserting `roles` with only `body["roles"] != nil` | A `null` value also satisfies `!= nil` after JSON decode — the []-not-null guarantee silently untested | Type-assert to `[]any` and check `len == 0` (R1 ⚠️) |
| Adding a defensive nil check by *logging the view* for debugging | Decrypted email lands in logs — the exact leak R6 forbids | Log nothing; the response body is the only place the email exists |
| Putting the shape tests behind `-tags=contract` | Shape only verified in the contract gate run — every plain `go test ./...` pass loses the assertion (D9) | Keep them in the plain suite; the contract tag adds nothing here |
| Extending `securityService` with `GetProfile` for convenience | Profile logic leaks into the security seam; future handler refactors drag an unrelated method set (D3) | Own `profileService` seam in the new file, per per-file convention |

## 14. Open Items

Lifecycle per rules.md § 8. Zero Active; one Resolved.

### Active — need external input or verification

(none)

### Resolved (kept for reference)

1. ~~**/auth/login's pre-existing camelCase `user` keys — fix in this slice or defer?**~~
   **RESOLVED — fix-in-slice accepted (Anhar, 2026-09-01, per agent recommendation; D10).**
   Verified fact underlying it: `LoginUserView` (`entity.go:118-127`) has no JSON tags and
   the login tests never asserted inner keys (`auth_login_test.go:103`), so both success
   writers emitted camelCase `user` keys — a standing violation of the User schema.
   Disposition: both success writers map through the shared `userResponse` (D10 option A);
   entity stays tag-free (D2); wire-shape delta flagged in the PR description for standard
   review veto. Consequence: zero clients existed at disposition time, so no migration or
   deprecation window needed; §7's drift risk row drops to "only if reverted."

---

*Synthesis note:* exploration sketches treated as input, not gospel — D6 supersedes
stage-3's `.String()`/`.Format()` conversion sketch, D8 supersedes the stage-2 summary's
`service.go` placement, and both supersessions are recorded in their decision rows. The
JSON-tag finding (Areas 3+5) was verified against the tree before being trusted
(`entity.go`, `auth_login.go`, `auth_login_test.go`) — it is real and pre-existing, not
introduced by this slice. Per the exploration's own account, the `GetLoginUserView` reuse
claim was verified against `repository.go:219` before "no new SQL" was asserted anywhere.
Open Item 1's disposition (fix-in-slice) was accepted by Anhar on 2026-09-01 before
finalization; D10, §6, §9 step 4, and R8 all record it explicitly — nothing about the login
fix is silent.
