import { useMutation } from "@tanstack/react-query";
import { mfaEnroll } from "@/lib/api/account";

/**
 * Wraps `POST /account/security/mfa/enroll` (techplan account/06-mfa-
 * totp, task-01). Bare mutation, no `onSuccess` — nothing on `User`
 * changes until `enrollConfirm` succeeds (`mfa_enabled` stays `false`
 * throughout a pending, unconfirmed enrollment), so there's no cache to
 * invalidate here.
 */
export function useMfaEnroll() {
  return useMutation({
    mutationFn: mfaEnroll,
  });
}
