# Stage 2 — Gap Analysis, Area 4: Google OAuth / link direction

> Files: `internal/domain/account/google_oauth.go`,
> `transport/http/auth_google.go`

## Current state (concrete)

- **`intent=link` direction fully implemented**: `Service.callbackLink`
  (google_oauth.go:413) — resolves the session user from the
  `oauthState` cookie (`UserID *uuid.UUID` binding), rejects if the
  google email is claimed by another user (`errGoogleLinkConflict`),
  otherwise inserts a **verified** google `AuthIdentity` +
  `UserLog{ActionType: "account_linking"}` atomically in one tx.
  Idempotent no-op when already linked to the same user.
- **`intent=reauth` implemented**: `GoogleCallback` returns
  `CallbackResult{Reauth: true, UserID}` with no DB writes; the handler
  sets a 5-minute in-memory marker (`reauthMarkers sync.Map`) — built
  for MFA-disable step-up, "consumed by task #06".
- Google identities are born `verified_at = now()` (R14) — an unverified
  google identity can never exist, so INV-account-12's "remaining
  identity must be verified" concern applies only to `email_password`.
- **No delete/removal path of any kind exists** — nothing in service or
  repository removes an `AuthIdentity`.

## Requirement vs Gap

1. No repo method to find identities by `(user_id, provider_type)` —
   unlink must locate the google row without knowing its email
   (`FindAuthIdentityByIdentifierHash` is keyed by email-hash; the
   plaintext google email lives only inside the encrypted column).
2. No guarded check-then-delete operation. The spec demands atomicity
   ("the whole check-then-delete sequence must be atomic (not racy)").
   Closest precedents are guarded single-statement UPDATEs (`RedeemToken`,
   `RotateRefreshToken`'s mark+insert pair).
3. Audit write precedent exists (`InsertUserLog` with
   `actionAccountLinking`) and is directly reusable for unlink's success
   audit entry.

## Sniffing

- *Misleading signal*: reauth infrastructure (`intentReauth`, 5-min
  marker) looks like "the re-auth story for security actions" — but this
  feature spec deliberately does NOT use it for unlink (password
  confirmation instead). An implementer skimming might wire unlink to
  the marker.
- *Edge case*: `callbackLink` permits attaching a second google identity
  (different google email) to a user who already has one — the unique
  index only blocks same-email duplicates. The unlink spec assumes
  singular "the google AuthIdentity row"; a user with two google
  identities is possible today and breaks that assumption. Flagged for
  Stage 3 decision.
- *Observation*: `oauthState.UserID` binds link intent to the session
  user via cookie; unlink handlers will authenticate via access token
  instead (cookie or Bearer) — see Area 5.
