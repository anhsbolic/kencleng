# Stage 2 — Area 7: Existing Patterns

## Current State

**Audit logging pattern** (from `security.go` UnlinkGoogle, line 244-343):
- `InsertUserLog` called within the same tx as the mutation
- `UserLog{ID: uuid.New(), UserID: userID, ActionType: "account_unlinked", CreatedAt: s.now()}`
- Audit write is atomic with the mutation (same tx)
- Action type is a string constant defined in the service file

**Notification pattern** (from `service.go`, `security.go`):
- `notification.Sender` interface with `SendVerificationEmail`, `SendNudgeEmail`, `SendPasswordResetEmail`
- `FakeSender` logs instead of sending (v1, no SMTP yet)
- Notifications sent AFTER commit, never inside tx (e.g. `security.go` line 170: `s.sendVerification(ctx, email, plainToken)` after `tx.Commit`)
- Nudge types are string constants: `NudgePasswordReset`, `NudgeGoogleOnly`, `NudgeSetPasswordConflict`
- No "role changed" notification type exists yet

**Error mapping pattern** (from `errors.go`):
- Service defines sentinel errors (`var ErrXxx = errors.New("...")`)
- `MapServiceError` switches on sentinels → HTTP status + Problem Details
- Handler can also map inline for special cases (e.g. `MfaEnrollConfirmHandler` maps `ErrInvalidTOTPCode || ErrMfaNotPending` → same 422)
- Problem type URIs follow `https://kencleng.dev/errors/xxx` pattern

**Handler structure pattern** (from `account_profile.go`, `account_security.go`):
- Narrow service interface per handler file (`profileService`, `securityService`)
- `toUserResponse` maps domain → wire format (snake_case JSON)
- `writeJSON` helper for success responses
- `WriteProblem` / `WriteValidationError` for error responses
- `UserIDFromContext` for auth extraction

**Route registration pattern** (from `main.go`):
- Sub-mux per route group (`authMux`, `accountMux`)
- Middleware chain: `RateLimit(RequireSession(handler))`
- Handlers registered with `mux.HandleFunc("METHOD /path", Handler(svc))`

**Concurrency pattern** (from `security.go` UnlinkGoogle):
- `FOR UPDATE` row locks for serialization under READ COMMITTED
- Deferred rollback with `committed` flag
- Guards checked first, mutation last (evaluation order)

## Requirement

From `08-role-assignment.md`:
- **Audit log** (line 162-168): write `user_logs` for both assign and revoke
- **Notification** (line 166-168): trigger user-facing notification to target user ("Role kamu di Kencleng telah diubah oleh Admin")
- **Concurrency** (line 44-47): concurrent conflicting assignments must not both succeed (INV-account-09)

## Gap

**1. Notification type for role changes:**
- No `NudgeRoleChanged` (or similar) constant exists in `notification/sender.go`
- The `Sender` interface has `SendNudgeEmail(ctx, to, nudgeType)` — the role-change notification can use this with a new nudge type constant
- New constant needed: `NudgeRoleChanged = "role_changed"` (or similar)
- The notification is sent to the **target** user, not the caller — the service needs the target user's email. This requires a `GetLoginUserView` call (or a lighter lookup) to get the target's email.

**2. Audit log action_type constants:**
- No constants for `"role_assigned"` or `"role_revoked"` exist
- The `UserLog.ActionType` comment says "the full action_type vocabulary are owned by the user-logs task (#08)" — this is the same task
- New constants needed in the service file (or a shared constants file)

**3. Concurrency guard for INV-account-09:**
- The spec requires that concurrent conflicting assignments (`admin` and `kurator` to the same user) don't both succeed
- The `UNIQUE (user_id, role)` constraint prevents duplicate role assignment, but doesn't prevent a user from having both `admin` AND `kurator` (the table structurally allows it per ERD line 498-504)
- The exclusivity check (does user have kurator? → can't assign admin) is a behavioral invariant, not a DB constraint
- Options: (a) `SELECT ... FOR UPDATE` on `user_roles` within the tx to serialize concurrent assignments, or (b) check-then-insert without locking (races possible), or (c) use a DB-level trigger (rejected by ERD design)
- The existing `UnlinkGoogle` pattern uses `FOR UPDATE` on `auth_identities` — same approach applies here

**4. No `HasRole` or `GetCallerRoles` method:**
- The admin authz check (handler level) needs to know if the caller has `role=admin`
- Existing `GetLoginUserView` returns roles but also decrypts email (unnecessary overhead for authz)
- Options: (a) use `GetLoginUserView` and accept the decrypt overhead, (b) add a lightweight `FindUserRolesByUser` method (already identified in Area 3), (c) add a `HasRole(userID, role) bool` method

## Sniffing

1. **Risk**: The notification is sent to the target user's email. But the target user's email is PII (encrypted). The service needs to decrypt it to pass to `SendNudgeEmail`. This is the same decrypt-on-read pattern as `GetLoginUserView`. The notification sender logs "recipient redacted" — no PII leak in logs.

2. **Edge cases**: If the notification send fails after commit, the role change is already persisted. The notification is best-effort (same pattern as `sendVerification` — if it fails, the operation still succeeded). The spec says "triggers a user-facing notification" — this is a forward-reference to the `notification` domain, same as `05-account-linking.md` and `06-mfa-totp.md`.

3. **Miscontext**: The spec says "cross-domain dependency on the not-yet-built `notification` domain, same forward-reference pattern as those two specs" (line 167-168). This means the notification is a forward-reference — the `Sender` interface already exists, but the actual email template/content doesn't. Using `SendNudgeEmail` with a new nudge type is the established pattern.

4. **Misleading signals**: The `Sender` interface looks complete — but `SendNudgeEmail` only logs in v1 (FakeSender). The actual email content for role changes doesn't exist yet. This is fine — the pattern is "send after commit, log in v1, real SMTP later."

5. **Inconsistency**: The existing audit log pattern writes `UserLog` within the tx. But the notification is sent after commit. For role assignment, the tx includes: (1) insert `user_roles`, (2) insert `user_logs`. After commit: (3) send notification. If the notification needs the target user's email, it needs a separate read (outside tx) or the email must be available from the pre-tx lookup.
