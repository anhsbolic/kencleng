# Stage 1 — Plan Announcement (05-account-linking)

> Session: 2026-08-26
> Feature spec: `docs/spec/1-account/features/05-account-linking.md`
> (note: actual spec dir is numbered — `1-account/`, not
> `domains/account/` as referenced in places)
> Workflow: `harscode-workspace/workflow/1-exploration/guidelines.md`
> Status: confirmed by Anhar; Stage 2 executed after this announcement.

## Task summary

Implement account linking per feature spec 05 — two endpoints:

- `POST /account/security/google/unlink` — remove the Google identity,
  guarded by INV-account-02 + the new INV-account-12, with password
  re-authentication.
- `POST /account/security/set-password` — server-side branch selection:
  Branch 1 (no `email_password` identity yet → add unverified identity +
  verification email, anti-enumeration), Branch 2 (has one → change
  password in place, `current_password` required, revoke all sessions).

Tier 1. Reuses registration verification mechanics and the reset-
password force-logout pattern.

## Areas to explore, in order

1. **Account domain core** — `internal/domain/account/entity.go`,
   `repository.go`, `repository_db.go`, `service.go`. Foundation for
   everything else: what an `AuthIdentity` looks like in code today,
   which repository operations already exist, what transaction patterns
   the service uses. Everything downstream depends on this vocabulary.
2. **DB schema** — `migrations/000002_create_auth_identities.*` plus
   auth_tokens, refresh_tokens, user_logs. Validates what the spec
   assumes actually exists at the schema level: `verified_at` column,
   unique index backing INV-account-01, no soft-delete column (spec
   says unlink is a hard delete), token tables' shape/purpose CHECK.
3. **Reused-flow implementations** — register + verify-email path
   (`transport/http/auth_register.go`, `auth_verify_email.go`) for the
   anti-enumeration pattern, token issuance, breach-check fail-open;
   and the forgot/reset-password path for the INV-account-05
   "revoke all refresh tokens" mechanic that Branch 2 must reuse.
   These are the two patterns the spec explicitly says to reuse.
4. **Google OAuth / link direction** — `domain/account/google_oauth.go`,
   `transport/http/auth_google.go`. The `intent=link` direction lives in
   Fitur 02, but unlink deletes the Google identity row — need to see
   how Google identities are created/read today and whether any delete
   path exists; also the reauth-marker infrastructure.
5. **Transport wiring + API contract** — route registration
   (`cmd/server/main.go`, middleware, errors.go mapping) and
   `api/openapi.yaml` schemas for both endpoints. Last because the spec
   flags both endpoints as needing an openapi schema update — findings
   from areas 1–4 determine how big that gap really is.
6. **Cross-domain stubs (light)** — `platform/notification` and the
   audit-log (`user_logs`) write path. The spec defers notifications to
   the unbuilt notification domain but requires audit entries now;
   confirm what exists vs what the spec forward-references.

## Why this order

Dependency-based: 1 gives the vocabulary, 2 validates the schema
assumptions under it, 3 covers the exact reusable patterns, 4–5 are
surfaces whose gaps only make sense once 1–3 are known, 6 is a small
confirmation pass.
