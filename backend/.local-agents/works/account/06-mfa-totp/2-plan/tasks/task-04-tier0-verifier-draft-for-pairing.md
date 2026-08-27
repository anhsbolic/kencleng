# Task 04: Tier 0 verifier core — `totpVerifier` (DRAFT-FOR-PAIRING)

> Back-reference : `.local-agents/works/account/06-mfa-totp/2-plan/techplan.md` (Status: Approved by Anhar) — sections 8 (redemption SQL + LoginMfa consumer notes), 10 (mfa_verifier.go), 5 (D1, D2, D3), 9 (steps 4 & 9), 14 (Resolved #7 / D12 sequencing)
> Depends on    : task-02 (`RedeemMFABackupCode` port), task-03 (normalize/hash helpers)
> Model         : GLM 5.2 (max) (concurrency-adjacent security reasoning; OUTPUT IS PAIRING GATED — see STOP below)
>
> ⚠️ **TIER 0 FENCED SUB-AREA** — tasks.md KPI: the TOTP secret generation/verification and guarded redemption logic must be human-authored or human-rewritten. This task's bodies ship marked DRAFT-FOR-PAIRING. Completion ≠ merged; completion = green suites + pairing checkpoint satisfied (techplan §9 step 9).

## Objective

Replace `stubMfaVerifier` with the real `MfaVerifier` implementation, keeping the `LoginMfa` consumer flow untouched (R10 parity). Wire production via a constructor the main-wiring task uses.

## Files

| File | Change |
|---|---|
| `backend/internal/domain/account/mfa_verifier.go` | stub removed; `totpVerifier{repo Repository, keys *crypto.Keys}` + `NewMfaVerifier(repo, keys) MfaVerifier`; both interface methods implemented |
| `backend/internal/domain/account/login_test.go` | modified — R10 real-verifier parity cases |
| header comment in `mfa_verifier.go` | explicit `Status: DRAFT-FOR-PAIRING (task #6 build; human pairing required before code-review — techplan D12)` marker |

## Method contracts

- `VerifyTOTP(ctx, userID, code) (bool, error)` — pure read (interface doc): loads decrypted secret via `GetMFATOTPSecretForVerify` (D3), computes via `totp.ValidateCustom` (skew ±1 per D1). No row / disabled row → `(false, nil)`; decrypt failure → `(false, err)` wrapped (§7 risk-row fail-safe direction).
- `VerifyBackupCode(ctx, tx pgx.Tx, userID, code string) (bool, error)` — normalize → SHA-256 hex (shared helpers from mfa.go) → `RedeemMFABackupCode(ctx, tx, userID, hash)` inside the CALLER's tx (the pgx.Tx passed by LoginMfa; never self-begins). Honors both INV-account-06 clauses because the port does.

## Implementation constraints

- Every crypto-bearing line stays inside this file or its paired rewrite scope; platform/crypto remains untouched (Tier 0 fence)
- No TOTP computation over an un-decrypted buffer, no second hashing scheme, no fallback validation path
- Secret material never logged (R15 extends here); errors wrapped `%w`
- Interop parameters locked to D1 defaults — no config surface for period/digits/skew

## Rules enabled (proven here at unit level)

R9 unit half (`TestMfaBackupCode_SingleUseGuarded` through the real verifier path), R10 full.

## Verification (pre-pairing gate)

- `go test ./internal/domain/account/... -race` green including R10 parity cases
- Stub references gone from compile graph except the retained nil-safety default in NewService (now unreachable in production wiring)
- Suite output archived as the review harness for the pairing session

## STOP — pairing checkpoint (gate, not suggestion)

When suites are green: STOP. Human pairing (Anhar rewrites/approves these bodies) precedes the code-review stage per D12. Record sign-off before task-05 merges transport work that depends on the constructor name only — actually task-05 may proceed on scaffolding in parallel EXCEPT final merge ordering waits for pairing sign-off. Do not self-approve.
