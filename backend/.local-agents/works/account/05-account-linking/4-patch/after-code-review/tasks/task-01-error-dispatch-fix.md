# Task 01: Fix error dispatch in security handlers (BLOCKING)

> Back-reference : `4-code-review/report.md` findings S1 / BP1 / C1 (one root cause); `2-plan/techplan.md` task-05 Files table originally specified the `MapServiceError` extension this restores
> Depends on    : nothing (patch to existing uncommitted working tree)
> Review ref    : root AGENTS.md golden rules ("error responses never leak internals — wrap into Problem Details"; "every authorization check is explicit"), `backend/AGENTS.md` §2 (errors wrapped with `%w`)
> Blocking      : **Yes** — gates merge

## Objective

Collapse three review findings (Safety S1, Best-Practices BP1,
Consistency C1) into one surgical fix: extend `MapServiceError` with
the two new 409 sentinels and route both security handlers' non-
validation errors through it, so internal failures (DB outage, tx
begin/commit failure, etc.) surface as `500 internal` instead of the
current misleading `401 invalid-credentials`. Keep the explicit
`ErrInvalidCredentials → 401` mapping visible at the handler boundary.

## Files

| File | Change |
|---|---|
| `internal/transport/http/errors.go` | + two `case` branches in `MapServiceError` for `account.ErrOnlyIdentity` (409, `https://kencleng.dev/errors/only-identity`) and `account.ErrRemainingUnverified` (409, `https://kencleng.dev/errors/unverified-remaining-identity`), with the verbatim titles + Indonesian details currently inlined in `UnlinkGoogleHandler` |
| `internal/transport/http/account_security.go` | `SetPasswordHandler`: after the `isErrValidation` branch, add `if errors.Is(err, account.ErrInvalidCredentials) { → 401 }`, else `MapServiceError(w, err)`. `UnlinkGoogleHandler`: replace the inline `switch` with `switch { case errors.Is(err, account.ErrOnlyIdentity): MapServiceError(w, err); case errors.Is(err, account.ErrRemainingUnverified): MapServiceError(w, err); case errors.Is(err, account.ErrInvalidCredentials): → 401 inline (matches existing login convention); default: MapServiceError(w, err) }`. Remove the now-duplicated inline 409 strings. |
| `internal/transport/http/account_security_test.go` | + `TestSetPasswordHandler_InternalError_500` and `TestUnlinkGoogleHandler_InternalError_500` — inject a non-sentinel wrapped error (`fmt.Errorf("boom: %w", errors.New("db down"))`) via `stubSecurityService`, assert `500` + problem type `https://kencleng.dev/problems/internal` + detail `"An unexpected error occurred."`. Keep the existing 401/409/422/200 tests green (they pin the legitimate categories). |

## Contracts to hit exactly

The two 409 Problem Details bodies are already specified verbatim in
the techplan task-05 contract and currently inlined in
`UnlinkGoogleHandler` — move them verbatim into `MapServiceError`:

- `ErrOnlyIdentity` → `409`, type `https://kencleng.dev/errors/only-identity`, title `"Cannot Unlink Only Identity"`, detail `"Google adalah satu-satunya metode login Anda. Atur email dan password dulu sebelum melepas tautan."`
- `ErrRemainingUnverified` → `409`, type `https://kencleng.dev/errors/unverified-remaining-identity`, title `"Remaining Identity Not Verified"`, detail `"Kamu sudah atur email dan password, tapi belum diverifikasi. Verifikasi email kamu dulu sebelum bisa melepas tautan Google."`
- `ErrInvalidCredentials` → stays `401` via `problemTypeInvalidCredentials` / `problemTitleInvalidCredentials` / `problemDetailGenericCredential` (already in `errors.go:98-101` — the handler may keep an inline 401 for this one case to keep the credentials check explicit at the boundary, OR delegate to `MapServiceError`; either is acceptable as long as the mapping is visible and tested).
- `default` (any wrapped `fmt.Errorf`) → `500`, type `https://kencleng.dev/problems/internal`, title `"Internal Error"`, detail `"An unexpected error occurred."` (already the `default` in `MapServiceError` — no change needed there).

## Tests

- **New**: `TestSetPasswordHandler_InternalError_500` — stub returns `fmt.Errorf("account: lookup identities: %w", errors.New("db down"))`; assert `rec.Code == 500`, `problem.Type == "https://kencleng.dev/problems/internal"`, `problem.Detail == "An unexpected error occurred."`.
- **New**: `TestUnlinkGoogleHandler_InternalError_500` — same shape for the unlink handler.
- **Existing, must stay green**: `TestSetPasswordHandler_Branch1_202`, `TestSetPasswordHandler_Branch2_200`, `TestSetPasswordHandler_PolicyFail_422`, `TestSetPasswordHandler_WrongPassword_401`, `TestUnlinkGoogleHandler_Success_200`, `TestUnlinkGoogleHandler_OnlyIdentity_409`, `TestUnlinkGoogleHandler_UnverifiedRemaining_409`, `TestUnlinkGoogleHandler_WrongPassword_401`, and the three `TestRequireSession_*` tests.
- **No service/integration change**: the dispatch fix is transport-only; the existing `go test -race ./internal/domain/account/...` and `go test -tags=integration ./internal/domain/account/...` suites must stay green unchanged.

## Common mistakes

- Inlining `MapServiceError`'s 500 default in the handler instead of calling it — defeats the single-source-of-truth point of C1.
- Dropping the explicit `ErrInvalidCredentials → 401` visibility at the handler boundary — the root AGENTS.md golden rule "every authorization check is explicit" favors keeping the credentials mapping testable at the boundary; delegating to `MapServiceError` is fine, but the test must still pin the 401.
- Editing the spec'd 409 problem-type URIs or Indonesian details while moving them — copy verbatim from the current handler (`api/openapi/account.yaml` is the source of truth).
- Regenerating the `api/openapi.yaml` bundle — out of scope (build report deviation #4).

## Out of scope here

- Q1 typed `SetPassword` result (task 02).
- Q2 shared password-validation helper (task 02).
- Q3 / BP2 down-migration comment tightening (task 02).
- `gitleaks` + full-`-race` runs (testing phase, `5-testing`).

## Verification

```bash
go test ./internal/transport/http/...                       # new 500 tests + existing handler tests green
go test -race ./internal/domain/account/...                 # security tests unaffected
go test -tags=integration ./internal/domain/account/...      # 4 integration tests unaffected
staticcheck ./...
```

Expected: all green, no new staticcheck findings in the touched files.
