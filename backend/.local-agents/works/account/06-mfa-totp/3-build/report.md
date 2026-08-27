# Build Report — MFA TOTP (account #06)

> Location : `.local-agents/works/account/06-mfa-totp/3-build/report.md`
> Date     : 2026-08-27
> Source   : `.local-agents/works/account/06-mfa-totp/2-plan/techplan.md` (Status: Approved by Anhar) + task decomposition in `2-plan/tasks/`
> Status   : **Build complete (code); gate partially proven** — see "Gate status" and the separate `testing-report.md`
> Author   : ox-alpha (agent), for Anhar's review

## Deliverable summary

Implemented the full MFA lifecycle slice end-to-end across the 6 decomposed tasks:
`enroll` → `enroll/confirm` → `disable`, backed by a real TOTP/backup-code verifier
replacing the login-slice's fail-closed stub.

| Task | Scope | Files | Status |
|---|---|---|---|
| 01 | `pquerna/otp` dependency + openapi enroll `409` amendment + bundle regen | `go.mod`, `go.sum`, `api/openapi/account.yaml`, `api/openapi.yaml` | ✅ |
| 02 | `MFABackupCode` entity + 6 repository ports + goqu impls | `entity.go`, `repository.go`, `repository_db.go`, `service_test.go` (fake) | ✅ |
| 03 | `MfaEnroll`/`MfaEnrollConfirm`/`MfaDisable` + sentinels + code helpers + unit tests | `mfa.go` (new), `mfa_test.go` (new) | ✅ |
| 04 | Tier 0 `totpVerifier` replacing stub + login-parity tests | `mfa_verifier.go`, `mfa_test.go` | ✅ (draft-for-pairing) |
| 05 | 3 handlers + `ConsumeReauthMarker` + error mappings + wiring | `account_security.go`, `auth_google.go`, `errors.go`, `main.go`, transport tests | ✅ |
| 06 | Integration/race suites + full gate | `mfa_integration_test.go` (new) | 🟡 written, compile-checked only (see testing report) |

## What was built (by layer)

### Repository (Task 02)
Six new port methods on `Repository`, implemented with goqu prepared statements in
`repository_db.go`, including the four guarded shapes specified verbatim in techplan §8:

- `UpsertPendingMFASecret` — **conflict-armed upsert** `ON CONFLICT (user_id) DO UPDATE … WHERE enabled_at IS NULL`. `inserted=false` is the designed 409 signal when MFA is already active (D5); a live secret is structurally never overwritten.
- `GetMFATOTPSecretForVerify` — decrypts at the storage boundary (D3); domain never sees ciphertext.
- `EnableMFATOTPIfPending` — guarded `enabled_at = now() … WHERE enabled_at IS NULL`; guarded-first in the confirm tx (D4-A).
- `SetMFADisabledIfEnabled` — guarded disable, idempotent repeat-safe; backup codes left untouched (INV-account-06).
- `InsertMFABackupCodes` — batch insert of hashes only.
- `RedeemMFABackupCode` — **one joined UPDATE** carrying *both* INV-account-06 clauses (`used_at IS NULL` + join to `mfa_totp_secrets.enabled_at IS NOT NULL`), so implicit invalidation is a DB clause, never app-side sequencing.

Verified by a throwaway SQL-inspection harness that the exact generated SQL for all four
guarded shapes matches the §8 contract (esp. the conflict-arm `WHERE` and the joined
redemption `FROM mfa_totp_secrets`).

### Service (Task 03)
New `internal/domain/account/mfa.go`:
- `MfaEnroll(ctx, userID) (uri string, err)` — generates via `totp.Generate` (Google-Authenticator defaults), encrypts at rest, upserts **pending**, returns `key.URL()`. Account label = plaintext primary email via `GetLoginUserView` (D11).
- `MfaEnrollConfirm(ctx, userID, code) ([]string, err)` — validates `totp.ValidateCustom` (skew ±1), then one tx: guarded enable → insert exactly-10 hashed codes → `mfa_enabled` audit. Returns plains **once**. Wrong code → `ErrInvalidTOTPCode` (pending preserved, R6); no-pending/lost-race → `ErrMfaNotPending` (wire-identical to wrong code, R7).
- `MfaDisable(ctx, userID, password) error` — server-side provider detection (R14); email_password requires `password` (wrong → `ErrInvalidCredentials`, missing → `ErrValidation`); Google-only relies on the handler-consumed marker (D6). One tx: guarded disable + `mfa_disabled` audit (skipped on idempotent repeat, R11).
- `MfaDisableReauthRequired(ctx, userID) (bool, err)` — transport helper to route the re-auth gate.
- Helpers: `generateBackupCodes`, `normalizeBackupCode` (shared by generator + verifier, D2); constants `backupCodeCount=10`, `otpauthIssuer`, `actionMfaEnabled/Disabled`, `totpPeriod/Skew`.

### Verifier — Tier 0 (Task 04)
`mfa_verifier.go`: removed the stub's production role, kept `stubMfaVerifier` as the
nil-safety default, added `totpVerifier` + `NewMfaVerifier(repo Repository)`.

**⚠️ Pairing checkpoint is pending (D12).** The `totpVerifier` bodies are the agent's full
draft, suites pass against them, but tasks.md's Tier-0-sub-area KPI requires Anhar to
human-pair/rewrite these before code-review/merge. File carries an explicit
DRAFT-FOR-PAIRING banner.

> **Synthesis deviation (flagged):** techplan §10 sketched `NewMfaVerifier(repo, keys)`.
> Per D3 the adapter owns all ciphertext, so the verifier holds no key material — the
> `keys` param became a dead dependency and was dropped (`NewMfaVerifier(repo)`). Same
> net semantics; surfaced in the file header for the pairing session.

### Transport + wiring (Task 05)
- `securityService` interface extended with the 3 service methods + reauth-routing helper.
- `MfaEnrollHandler`: session → service → 200 `{otpauth_uri}` / 409 via `MapServiceError`.
- `MfaEnrollConfirmHandler`: required-field `totp_code` validation; R7 byte-identical 422 for wrong-code vs no-pending.
- `MfaDisableHandler`: tolerant empty-body decode; Google-only → `ConsumeReauthMarker(userID)` (consume-on-use), and any submitted password is ignored (R14 — password cannot bypass the marker); email_password missing password → 422 required.
- `ConsumeReauthMarker` (`auth_google.go`): `LoadAndDelete` + expiry recheck — atomic check-and-invalidate; existing `CheckReauthMarker` + sweeper untouched (D7).
- `MapServiceError` (`errors.go`): `ErrMfaAlreadyEnabled` → 409 `errors/mfa-already-enabled`; both confirm sentinels → 422 validation.
- `main.go`: `mfaVerifier := account.NewMfaVerifier(repo)` injected (replacing `nil`), 3 routes registered on the existing guarded `accountMux`.

### Contract artifact (Task 01)
- `api/openapi/account.yaml`: `/account/security/mfa/enroll` gained the feature-spec-mandated `409` (`errors/mfa-already-enabled`); `/account/security/mfa/disable` gained the `422` (missing password). `api/openapi.yaml` regenerated via `npm run bundle`; `npm run validate` shows no new errors introduced.

## Gate status

| Gate | Result |
|---|---|
| `go build ./...` | ✅ PASS |
| `go vet ./...` | ✅ PASS |
| `go test ./...` | ✅ PASS (unit; 8 packages, ~10 s) |
| `go test -race` full account | ✅ PASS earlier (319 s) before integration file added; MFA `-race` suite ✅ 17 s |
| `go vet -tags=integration` | ✅ PASS (MFA integration tests compile) |
| `go test -tags=integration` (DB) | 🟡 **NOT validated in-session** — see testing report |
| `make verify` (lint/gosec/gitleaks/govulncheck/contract) | 🟡 not run (see testing report) |

## Open findings / follow-ups

1. **Tier 0 pairing** — `totpVerifier` bodies await Anhar's rewrite/sign-off (D12) before code-review/merge. Blocks merge, not scaffold.
2. **Integration suite unrun** — `mfa_integration_test.go` is written and compile-checked, with a pool-leak bug in one test fixed before this report; it was not executed to green (hangs/timeouts — full detail in `testing-report.md`). Runs only with local Postgres (`DATABASE_URL` on `:5435`) + `-tags=integration`.
3. **Pre-existing openapi defect confirmed (carried from task #05, NOT introduced here)**: the regenerated bundle is missing `components.securitySchemes` because `index.yaml` declares `security: [bearerAuth]` but never `$ref`s `common.yaml#/components/securitySchemes`. This predates this slice (`git show HEAD` had 0 occurrences). Fix = a source edit to `index.yaml` (add a `components` $ref), which is outside the single approved `account.yaml` amendment — needs a decision. Not caused by, and does not affect, this backend work.
4. **Frontend changes present in `git status` are NOT mine** — this session is backend-scoped (root AGENTS §7 directory boundary). Any `frontend/components/features/account/mfa-*` etc. are parallel/pre-existing work left untouched.

## Risk note (per root AGENTS §5)

- **Assumptions made:** `pquerna/otp` defaults are the intended TOTP parameters (feature spec names the lib; Google-Authenticator-defaults documented in D1). Integration service seals the `now`/`compare` seams that `integrationService` leaves unset. `NewMfaVerifier` signature dropped the dead `keys` param (D3).
- **Edge cases intentionally NOT handled (why):** TOTP clock drift beyond ±1 step (standard, not over-engineered); backup-code brute force at login (covered by the task-#03 MFA-stage lockout, not this slice); backup-code row accumulation across disable/re-enable (accepted threat-model §5 residual).
- **Concurrency assumptions:** guarded-upsert (`ON CONFLICT … WHERE enabled_at IS NULL`) and joined single-statement redemption make the two High-severity races impossible at the DB level; the ≥100-goroutine `-race` asserts are authored (Task 06) and the unit-level R8 stress passed (exactly 1 winner / 10 codes) — the **real-DB** version is authored but unrun.
- **What is not tested, and why:** the Postgres-backed integration suite (R1/R3/R5/R8/R9/R11/R12 DB truths) and full `make verify` were not executed in-session — the integration runs hung (see testing report) and the user directed them be skipped. These must run against a reachable DB before merge.
