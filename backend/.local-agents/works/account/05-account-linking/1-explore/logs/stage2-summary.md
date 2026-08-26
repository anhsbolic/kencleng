# Stage 2 — Cross-area summary (05-account-linking)

> Companion to the six gap-areaN-*.md docs in this directory.
> Written 2026-08-26, after all areas were explored.

## Hard gaps (code that must be written)

1. Repo/service methods: identity lookup by `(user_id, provider_type)`,
   list identities w/ verified flags, guarded atomic check-then-delete
   for unlink, `UpdateCredentialSecret` (Branch 2), **user-wide
   refresh-token revocation (INV-account-05 — first implementation
   ever)**.
2. Service flows: `SetPassword` (2 branches), `UnlinkGoogle`; new
   sentinels for the two distinct 409s + wrong-password 401.
3. Transport: first authenticated endpoints — session
   extraction/verification middleware-or-helper, two handlers,
   `/account/security/*` route group.
4. One new nudge type for Branch 1 conflict.

## Notable sniffing results (cross-cutting)

- **Spec-vs-reality miscontext**: feature 05 says "reuse INV-account-05
  pattern from Fitur 04" — Fitur 04 doesn't exist; this task implements
  it first, and tasks.md's serial order S1 (#1→#5) is being jumped (#3,
  #4 incomplete).
- **Stale claims**: openapi schemas already carry the redesign despite
  spec saying both need updates; tasks.md tracker stale re task #3 code
  presence.
- **Contract defect found**: `bearerAuth` referenced globally but
  `components.securitySchemes` missing from openapi.yaml entirely.
- **Edge cases flagged**: possible multi-google-identity users (unlink
  assumes singular); audit-at-verification-time for Branch 1 can't
  distinguish set-password from registration tokens (same purpose);
  `SetUserVerified` multi-row semantics under the one-identity
  assumption.
