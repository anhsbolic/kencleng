# Build Report — account #07 (initial build)

## What changed

- `internal/domain/account/profile.go` (new) — `Service.GetProfile` pass-through to `s.repo.GetLoginUserView`; `(nil, nil)` means gone user.
- `internal/domain/account/service_test.go` (modified) — added `getViewCalls []uuid.UUID` recording field to `fakeRepo` + recording in `GetLoginUserView`; added `TestGetProfile_PassesThroughToRepository` (R7).
- `internal/transport/http/account_profile.go` (new) — `profileService` seam, `userResponse` struct (8 snake_case keys), `toUserResponse` (nil→`[]` normalization + nil-view→nil), `AccountMeHandler`.
- `internal/transport/http/account_profile_test.go` (new) — R1–R6 tests + `fullView`/`assertUserShape`/`getAccountMe` helpers.
- `internal/transport/http/auth_login.go` (modified) — `loginOKResponse.User` retyped `*account.LoginUserView` → `*userResponse`; both success writers (`LoginHandler`, `LoginMfaHandler`) map `res.User` through `toUserResponse` (D10).
- `internal/transport/http/auth_login_test.go` (modified) — added `TestLoginHandler_UserShapeSnakeCase` + `TestLoginMfaHandler_UserShapeSnakeCase` (R8), sharing `assertUserShape`.
- `cmd/server/main.go` (modified) — added `GET /account/me` route under `RateLimit` + `RequireSession(googleVerifyToken)`.
- `2-techplan/techplan.md` (modified) — Status `Approved by Anhar` + `Updated` line (contract locked).

## Tests run

- `go test ./internal/domain/account/` → unit → ok (8.741s)
- `go test ./internal/transport/http/` → mocked-service + API-contract (handler shape/status) → ok (0.006s)
- `go test ./...` → unit/mocked/contract across repo → ok
- `go vet ./...` → ok; `gofmt -l` on all touched files → clean

Confirmed explicitly: no `-race`, perf/load, or security-class test was run in this iteration.

## Contract check

- [x] This iteration satisfies its build target in full (no task files — techplan § 4 in full: R1–R8 implemented and tested)
- [x] No contract assumption broke

One faithful-reconciliation note (not a break): techplan § 10 shows `toUserResponse` returning a value (`userResponse`) while also specifying `loginOKResponse.User` as `*userResponse` with `User: toUserResponse(res.User)`. To satisfy both call sites (and preserve the original `omitempty` nil-user omission), `toUserResponse` returns `*userResponse` (nil for nil view). The wire contract (snake_case keys, `[]`-not-`null`) is unchanged; this is the only reading consistent with § 10's own login-writer instruction.

## Flagged for techplan/testing review (if any)

- None — no concurrency/perf/security concern surfaced; the JSON-shape drift (login endpoint) is resolved by this slice per D10, and the Test Focus Pointer rows are both N/A (self-view PII surface + contract-correctness, covered by R1/R8 unit tests).
