"use client";

import { useSearchParams } from "next/navigation";
import { useEffect, useRef } from "react";
import { Banner } from "@/components/ui/banner";

// TBD — techplan account/02-google-oauth-login-register Open Item #4:
// placeholder copy pending product sign-off, same treatment as
// RegisterForm's GENERIC_ERROR_MESSAGE / VerifyEmailStatus's
// INVALID_LINK_MESSAGE elsewhere in this codebase.
const EMAIL_CONFLICT_MESSAGE =
  "Email ini sudah terdaftar dengan password. Silakan masuk dengan password, atau gunakan \"Lupa password\" jika perlu.";
const GENERIC_RETRY_MESSAGE = "Gagal masuk dengan Google. Silakan coba lagi.";

/**
 * `state_mismatch` / `nonce_mismatch` / `google_token_invalid` /
 * `google_unavailable` all collapse into one generic retry message
 * (R6) — none is more actionable to the user than the others.
 * `google_email_conflict` is the one code that gets distinct copy:
 * it's the no-auto-merge case (spec 02's top-severity anti-takeover
 * threat), where a generic "try again" message would mislead a
 * legitimate user into retrying the exact same failing action instead
 * of using their existing password login. Any code outside this known
 * set (including none at all, handled by the early return below)
 * falls back to the same generic message — the raw code value is
 * never rendered (`patterns.md` §B).
 */
function messageForCode(code: string): string {
  if (code === "google_email_conflict") return EMAIL_CONFLICT_MESSAGE;
  return GENERIC_RETRY_MESSAGE;
}

/**
 * `/login`'s `?error={code}` outcome (techplan account/02-google-
 * oauth-login-register, R5-R7). Renders nothing when no `error` param
 * is present. Must be mounted as `AuthShellClient`'s first child, per
 * that shell's own documented convention (`auth-shell-client.tsx`) —
 * this component only decides *what* to show, not *where*.
 */
export function GoogleCallbackError() {
  const searchParams = useSearchParams();
  const code = searchParams.get("error");
  const bannerRef = useRef<HTMLDivElement>(null);

  // R7 — focus moves into the banner on render, matching the
  // focus-management convention already established by
  // RegisterForm/VerifyEmailStatus elsewhere in this codebase.
  useEffect(() => {
    if (code) bannerRef.current?.focus();
  }, [code]);

  if (!code) return null;

  return (
    <div ref={bannerRef} tabIndex={-1} className="outline-none">
      <Banner variant="error">{messageForCode(code)}</Banner>
    </div>
  );
}
