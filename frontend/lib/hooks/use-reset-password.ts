import { useMutation } from "@tanstack/react-query";
import { resetPassword } from "@/lib/api/account";

/**
 * Wraps `POST /auth/reset-password`. Deliberately a bare mutation with no
 * `onSuccess` redirect (techplan account/04, D5): unlike `useLogin`, this
 * endpoint issues no session to land the user inside — every existing
 * guest-facing terminal-success precedent in this domain (`RegisterForm`,
 * `VerifyEmailStatus`) uses an inline view + manual link instead, which is
 * `ResetPasswordForm`'s job to render, not this hook's.
 */
export function useResetPassword() {
  return useMutation({
    mutationFn: resetPassword,
  });
}
