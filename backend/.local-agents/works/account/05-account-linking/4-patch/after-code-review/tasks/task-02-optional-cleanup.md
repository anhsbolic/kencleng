# Task 02: Optional follow-ups (NON-BLOCKING)

> Back-reference : `4-code-review/report.md` findings Q1 (typed result), Q2 (shared validation helper), Q3 / BP2 (down-migration comment)
> Depends on    : nothing; independent of task 01
> Blocking      : **No** — cherry-pick per item; do not gate merge on any of these

## Objective

Collect the three optional follow-ups the review flagged into one
clearly-marked non-blocking bucket. Each is a small, localized
readability/maintainability improvement. None affects behavior
correctness. They may be landed together, individually, or deferred.

## Items

### Q1 — Self-documenting `SetPassword` return

- **Location**: `internal/domain/account/security.go:51` (`func (s *Service) SetPassword(...) (bool, error)`)
- **Issue**: The `bool` means "Branch 2 ran → handler writes 200," but a caller cannot tell that from the signature without the doc comment.
- **Suggested fix**: Replace `(bool, error)` with a one-field result struct (`type SetPasswordOutcome struct{ Branch2Ran bool }`) or a named return. Update `SetPasswordHandler` and the `securityService` interface in `account_security.go`, the `stubSecurityService` in `account_security_test.go`, and the three `setPasswordResult`-shaped fields in the stub. Update `security_test.go` call sites to read the field instead of the bare bool.
- **Risk**: Touches the service seam signature — cascades to handler + stub + tests. Pure mechanical rename; no behavior change. Run `go test ./internal/domain/account/... ./internal/transport/http/...` to confirm.
- **Do not**: change the `(false, nil)` / `(true, nil)` / `(false, ErrInvalidCredentials)` / `(false, ErrValidation)` outcome semantics — only the carrier shape.

### Q2 — Shared password-validation helper

- **Location**: `internal/transport/http/account_security.go:107-108` vs `auth_register.go:40-41` (`len(req.Password) < 8` + `"password is not allowed"`)
- **Issue**: Duplicated boundary-validation literal across two handlers; no tool flags the drift.
- **Suggested fix**: Extract a `minPasswordLength = 8` named constant (per `go/error-handling.md` §2) and/or a shared `validatePasswordFields(password string) []fieldError` helper in `errors.go`; call from both `RegisterHandler` and `SetPasswordHandler`.
- **Risk**: Touches `auth_register.go` — already-merged code. **Scope-discipline check (root AGENTS.md §7):** if this slice's task scope does not permit editing the register handler, defer Q2 to a dedicated cleanup pass instead of crossing the slice boundary. At minimum, land the named constant in `errors.go` and have the new handler use it; the register handler can adopt it in a later pass.
- **Do not**: change the policy threshold (8) or the message text — they are spec'd in the techplan and openapi.

### Q3 / BP2 — Down-migration comment tightening

- **Location**: `migrations/000010_widen_auth_tokens_purpose.down.sql:4-7`
- **Issue**: The comment claims "token redemption is purpose-blind." Post-feature, redemption is NOT purpose-blind: `VerifyEmail`'s guard checks `purpose != purposeEmailVerify && purpose != purposeEmailVerifyLink`, and the R14 audit is conditional on `purpose == purposeEmailVerifyLink`.
- **Suggested fix**: Rewrite the comment to state the actual safety contract:
  > Re-map BEFORE restoring the 2-value CHECK. Safe when the migration
  > rolls back alongside the feature code: rolled-back `VerifyEmail`
  > has no audit logic and the 2-value guard accepts the remapped
  > `email_verification` token. A standalone rollback (migration down
  > while the new code stays) leaves link-purpose tokens redeeming as
  > registration tokens — the identity still verifies, but the R14
  > audit is silently skipped. Do not roll back the migration alone.
- **Risk**: Comment-only; no behavior change. No tests to run beyond `migrate-down` smoke if a sandbox DB is available.

## Out of scope

- Anything in task 01 (the blocking fix).
- The `gitleaks` + full-`-race` process gates (testing phase).
- The `InsertAuthIdentity` `verified_at` omission (build report deviation #5).
- The openapi bundle regeneration gap (build report deviation #4).

## Verification (only if items land)

```bash
go test ./internal/domain/account/... ./internal/transport/http/...   # Q1/Q2 mechanical rename must stay green
staticcheck ./...
```

Q3 is comment-only — no test impact.
