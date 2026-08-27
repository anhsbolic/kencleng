# Stage 2 — Area 5: Crypto Platform — TOTP Secret Encryption

> Files: `internal/platform/crypto/crypto.go`, `internal/platform/crypto/keys.go`

## Current State

`platform/crypto/crypto.go` provides:
- `Encrypt(plaintext []byte, keys *Keys) ([]byte, error)` — AES-GCM, random nonce, returns `nonce || ciphertext || tag`
- `Decrypt(ciphertext []byte, keys *Keys) ([]byte, error)` — reverses the above
- `HMAC(data []byte, keys *Keys) string` — SHA-256 hex digest for lookup hashes

These are generic primitives. No TOTP-specific code exists anywhere in the codebase. No `pquerna/otp` or equivalent library is imported.

## Requirement

The MFA TOTP implementation needs:
1. **TOTP secret generation** — generate a random base32-encoded secret (RFC 6238)
2. **TOTP code verification** — compute TOTP from the secret + current time, compare against submitted code (with ±1 step tolerance)
3. **Secret encryption at rest** — encrypt the base32 secret using `crypto.Encrypt` before storing in `secret_encrypted`
4. **Secret decryption for verification** — decrypt `secret_encrypted` using `crypto.Decrypt` before computing TOTP
5. **Backup code generation** — generate 10 random alphanumeric codes (e.g., 8 chars each)
6. **Backup code hashing** — SHA-256 hash for storage in `code_hash`

## Gap

No TOTP library is imported. No TOTP generation/verification code exists. The `crypto.Encrypt`/`Decrypt` primitives are ready, but the TOTP-specific logic (secret generation, code computation, time-window verification) is entirely missing. This is the Tier 0 fenced sub-area.

## Sniffing

- **Risk:** The TOTP secret is the highest-sensitivity material in this feature. If the encryption key (`keys.EncryptionKey`) is lost, all TOTP secrets become undecryptable and every MFA-enabled user is locked out. Same risk as losing the PII encryption key — already accepted for email/identifier fields. No new risk.
- **Edge cases:**
  - What if `Decrypt` fails on a stored secret (corrupted data, key rotation)? The verifier should return `(false, error)` — the user cannot complete MFA login. This is a data-integrity failure, not a credential failure. The error should be logged but not exposed to the user.
  - TOTP time-window tolerance: standard is ±1 step (30 seconds each). If the server clock is significantly drifted from the user's device, verification fails. Known TOTP limitation.
  - Backup code entropy: 8 alphanumeric chars from `crypto/rand` gives ~47 bits of entropy per code. With 10 codes, sufficient for a sandbox project.
- **Miscontext:** The feature spec says "same AES-GCM encryption-at-rest mechanism used for other sensitive fields." This is `crypto.Encrypt` — confirmed. No new key-management scheme needed.
- **Misleading signal:** `platform/crypto` looks like it might already support TOTP because it has `Encrypt`/`Decrypt`. It does — but only the encryption primitives. The TOTP protocol logic (RFC 6238 computation, time windows, base32 encoding) is a separate concern that doesn't exist yet.
- **Inconsistency:** None found. The crypto platform is clean and generic.
