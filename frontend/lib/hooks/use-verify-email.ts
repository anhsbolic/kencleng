import { useMutation } from "@tanstack/react-query";
import { verifyEmail } from "@/lib/api/account";

/**
 * Wraps `POST /auth/verify-email`. The single-fire-per-token guard
 * (account/01-register-email-verification's rule R12) is the caller's
 * responsibility (`VerifyEmailStatus`), not this hook's — `useMutation`
 * itself has no built-in call-dedupe.
 */
export function useVerifyEmail() {
  return useMutation({
    mutationFn: verifyEmail,
  });
}
