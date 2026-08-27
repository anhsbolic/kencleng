# Stage 2 — Area 6: Reauth Marker Infrastructure

> File: `internal/transport/http/auth_google.go`

## Current State

`auth_google.go` already implements:
- `reauthMarkers sync.Map` — in-memory store, key: `uuid.UUID`, value: `time.Time` (expiry)
- `SetReauthMarker(userID, expiry)` — stores the marker (line 106)
- `CheckReauthMarker(userID) bool` — checks validity, deletes if expired (line 109-118)
- Background eviction goroutine (line 130-138) — cleans expired markers every minute
- `reauthMarkerTTL = 5 * time.Minute` (line 26)
- Set by `GoogleCallbackHandler` on `intent=reauth` success (line 244)

## Requirement

Task #6's `MfaDisable` needs to consume the reauth marker for Google-only users. This is a read-only consumption — task #6 does not modify the reauth infrastructure.

## Gap

No gap in the infrastructure itself. The gap is in how the service layer accesses `CheckReauthMarker` — it's currently a package-level function in `transport/http`, not accessible from the domain service. The service needs either:
- An injected dependency (interface) that wraps the reauth check
- Or the handler checks the marker before calling the service (keeping the reauth check at the transport boundary)

## Sniffing

- **Risk:** The reauth marker is in-memory (`sync.Map`). If the server restarts between the Google reauth redirect and the MFA disable call, the marker is lost. The user must re-authenticate. Accepted (same as task #02's decision, 5-min TTL makes restart rare in practice).
- **Edge cases:** What if the user reauths via Google, then calls disable, then calls disable again? The marker is consumed (deleted) on first use — second call gets 401. Correct per the spec ("consumed on use so it can't be replayed").
- **Miscontext:** The feature spec says "the marker is consumed (invalidated) on use." `CheckReauthMarker` currently does NOT consume the marker — it only checks validity and deletes expired ones. The consumption (delete on successful use) must be added, or a separate `ConsumeReauthMarker` function must be created. This is a subtle but important gap.
- **Misleading signal:** `CheckReauthMarker` exists and looks ready to use. But it doesn't consume the marker on success — it's a read-only check. The disable flow needs a check-then-delete atomic operation. This might need a new function or a modification to `CheckReauthMarker`.
- **Inconsistency:** The spec says "consumed on use" but the current implementation only checks, doesn't consume. This is a real inconsistency that must be resolved during implementation.
