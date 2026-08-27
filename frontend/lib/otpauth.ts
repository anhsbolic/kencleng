/**
 * `otpauth://` URI parsing — client-only, no backend contract of its
 * own (techplan account/06-mfa-totp, D11/R24). The backend's
 * `MfaEnrollResponse.otpauth_uri` conventionally carries the raw TOTP
 * secret in a `secret` query parameter (the standard Google
 * Authenticator-style `otpauth://totp/Label?secret=...&issuer=...`
 * shape) — this is used purely for an accessibility fallback (manual
 * entry alongside the QR code), never sent back to the server.
 */

/**
 * Extracts the `secret` query parameter from an `otpauth://` URI, for
 * rendering as a manual-entry fallback next to the QR code (users
 * without a camera-capable device, or using a screen reader).
 *
 * Never throws — a malformed/unexpected URI shape (or a missing
 * `secret` param) resolves to `null` rather than crashing the caller,
 * matching `readProblemDetail()`'s existing "best-effort, never throws"
 * convention in `lib/api/client.ts`. `TBD — verify` the real
 * backend-generated URI always carries a `secret` param — this
 * function treats its absence as merely unhandled, not impossible.
 */
export function parseOtpauthSecret(otpauthUri: string): string | null {
  try {
    const secret = new URL(otpauthUri).searchParams.get("secret");
    return secret && secret.length > 0 ? secret : null;
  } catch {
    return null;
  }
}
