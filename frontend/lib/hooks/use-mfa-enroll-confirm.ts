import { useMutation, useQueryClient } from "@tanstack/react-query";
import { mfaEnrollConfirm } from "@/lib/api/account";
import { accountKeys } from "./use-account-me";

/**
 * Wraps `POST /account/security/mfa/enroll/confirm` (techplan account/
 * 06-mfa-totp, task-01, R9 hook-layer). On success, `mfa_enabled`
 * becomes `true` — invalidate `account.me` so the section re-derives
 * its branch once the caller (`MfaEnrollFlow`) is done with the
 * one-time `backup_codes` in the resolved value. The codes themselves
 * are not written anywhere here — they're read directly off this
 * mutation's own resolved data by the caller.
 */
export function useMfaEnrollConfirm() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: mfaEnrollConfirm,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: accountKeys.me() });
    },
  });
}
