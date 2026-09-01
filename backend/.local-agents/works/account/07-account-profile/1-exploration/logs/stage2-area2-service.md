# Stage 2 — Area 2: Service Layer

## Current State

No `GetProfile` or `GetMe` method exists on `*Service`. The service has methods for Register, Login, LoginMfa, Refresh, Logout, VerifyEmail, ResendVerification, ForgotPassword, ResetPassword, SetPassword, UnlinkGoogle, MfaEnroll, MfaEnrollConfirm, MfaDisable, GoogleRedirect, GoogleCallback, IssueTokens — but nothing for fetching the current user's profile.

The repository interface (`repository.go:219`) includes `GetLoginUserView(ctx, userID) (*LoginUserView, error)`. The login flow calls it directly via `s.repo.GetLoginUserView` (`login.go:144`, `login.go:241`). The MFA enroll flow also calls it (`mfa.go:75`).

The transport layer never calls the repository directly — all handlers go through the service. This is a consistent pattern across all existing handlers (`account_security.go`, `auth_login.go`, `auth_google.go`, etc.).

## Requirement

Spec says: `GET /account/me` returns the caller's own `User` resource. The handler needs to:
1. Extract `userID` from session context (via `UserIDFromContext`)
2. Fetch the user profile
3. Return it as JSON

## Gap

A service method is needed. Two options:

**Option A: Thin pass-through** — Add `GetProfile(ctx, userID) (*LoginUserView, error)` that just calls `s.repo.GetLoginUserView`. This is what the login flow does internally.

**Option B: Handler calls repo directly** — Break the pattern and have the handler call `repo.GetLoginUserView`. This would require the handler to depend on the concrete `*RepositoryDB` or a new interface.

The codebase pattern is clear: handlers depend on a narrow service interface, never on the repository. Option A is the consistent choice.

## Sniffing

1. **Risk**: None. This is a pure read with no business logic. The service method is a pass-through — no new logic to get wrong.

2. **Edge cases**: If the user was deleted after the session was issued, `GetLoginUserView` returns `(nil, nil)`. The handler must check for nil and return 401 (session references a non-existent user). This is the same edge case identified in Area 1.

3. **Miscontext**: None. The spec doesn't require any business logic beyond "fetch and return."

4. **Misleading signals**: The `LoginUserView` name suggests it's only for login — but the doc comment (`entity.go:110-117`) explicitly says it's "the read model assembled at login time to populate LoginResponse.user (openapi components.schemas.User)." The name is slightly misleading but the intent is clear.

5. **Inconsistency**: None. The pattern is consistent: service method → repo method → DB query.
