# Task 03: Service layer — MfaEnroll / MfaEnrollConfirm / MfaDisable

> Back-reference : `.local-agents/works/account/06-mfa-totp/2-plan/techplan.md` (Status: Approved by Anhar) — sections 8 (Business logic flow block — authoritative), 10 (mfa.go, service.go), 5 (D4, D5, D8, D9, D11), 3 (Requirements table)
> Depends on    : task-01 (pquerna/otp importable), task-02 (repo ports exist)
> Model         : GLM 5.2 (max) (multi-step branching + guarded-tx + timing-parity reasoning; reward-hacking caveat mitigated by the mandatory dual-model review at Complex tier)

## Objective

Implement the three lifecycle methods with their sentinel vocabulary, audit constants, and the shared backup-code material helpers. No transport concerns, no marker knowledge — the service sees only `(userID, password)`.

## Files

| File | Change |
|---|---|
| `backend/internal/domain/account/mfa.go` | new — services, sentinels, helpers, constants |
| `backend/internal/domain/account/mfa_test.go` | new — unit suites R1–R8, R11–R15 |

## Interface additions

```go
func (s *Service) MfaEnroll(ctx context.Context, userID uuid.UUID) (string, error)
func (s *Service) MfaEnrollConfirm(ctx context.Context, userID uuid.UUID, totpCode string) ([]string, error)
func (s *Service) MfaDisable(ctx context.Context, userID uuid.UUID, password string) error
```

Sentinels: `ErrMfaAlreadyEnabled`, `ErrInvalidTOTPCode`, `ErrMfaNotPending` (D8 — the latter two collapse to identical 422 on the wire; `ErrInvalidCredentials` reused for wrong-password disable).

Constants/helpers (mfa.go): `backupCodeCount = 10`, `otpauthIssuer = "Kencleng"`, `actionMfaEnabled`/`actionMfaDisabled` (D9), plus **shared material helpers** `generateBackupCodes(n int) ([]string, error)` and `normalizeBackupCode(string) string` — placed here because confirm-generation needs them first; task-04's verifier consumes them cross-file in-package.

## Flow contract (techplan §8 verbatim)

```
MfaEnroll(ctx, userID):
  base32 := pquerna totp.Generate(issuer:"Kencleng", account:<primary_email via GetLoginUserView>)  # D11
  rows := UpsertPendingMFASecret(userID, Encrypt(base32))   # D5 guard
  rows == 0 -> ErrMfaAlreadyEnabled
  return base32.URL()                                        # NEVER logged (R15)

MfaEnrollConfirm(ctx, userID, code):
  rec := GetMFATOTPSecretForVerify(userID)
  !found || rec.EnabledAt != nil -> ErrMfaNotPending         # ≡ ErrInvalidTOTPCode wire-identical (R7)
  ValidateCustom(secret, code, skew±1)?                      # pure compute; false -> ErrInvalidTOTPCode,
                                                             # pending SURVIVES (R6)
  tx {
     ok := EnableMFATOTPIfPending(tx, userID)                # guarded-FIRST (D4-A) — codes never
                                                             # precede the enable inside the tx
     !ok -> rollback; ErrMfaNotPending                       # race loser ≡ generic failure (R7/R8)
     InsertMFABackupCodes(tx, 10× {ID, UserID, CodeHash})    # sha256(normalize(plain))
     InsertUserLog(tx, {action_type:"mfa_enabled"})
  } commit -> return plains                                  # response-only, shown once

MfaDisable(ctx, userID, password):
  identities := FindAuthIdentitiesByUser(userID)             # server-side branch (R14)
  if has email_password identity:
     password == "" -> ErrValidation                         # 422 required
     compare(storedBcrypt, password) != nil -> ErrInvalidCredentials   # CPU burns either way
  else:                                                      # Google-only — reauth already happened at
                                                             # handler; service trusts credentials-as-data (D6)
  tx {
     disabled := SetMFADisabledIfEnabled(tx, userID)
     if disabled { InsertUserLog(tx, {action_type:"mfa_disabled"}) }  # skip audit on 0-row no-op (R11)
  } commit -> 200-classified by transport                    # codes untouched, implicitly dead (R9)
```

## Implementation constraints

- Enroll label sources plaintext email via existing `GetLoginUserView` decrypt path (D11) — no new repo surface for it
- Confirm validates TOTP inline via `totp.ValidateCustom` over the decrypted secret (RFC defaults per D1); generate uses `totp.Generate` with issuer constant and email account label
- Marker consumption NEVER appears here (transport concern — §13 mistake row 5)
- Unit tests table-driven per backend AGENTS §2; named exactly as techplan §12 rows for R1–R8, R11–R12, R14 unit halves, R15 log-scan (`TestMfa_LogsFreeOfSecrets`)
- Audit-value assertions check exact literals (tasks.md KPI)

## Rules enabled (proven here at unit level)

R1, R2, R3, R5–R8, R11, R12, R14 (unit halves), R15 (scan). Integration/race proof deferred to task-06.

## Verification

- `go test ./internal/domain/account/... -run 'TestMfa' -race` green
- `go vet ./...` clean
- Byte-identical body assertion exists for R7 before marking done
