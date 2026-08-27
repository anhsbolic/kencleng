import { useMutation, useQueryClient } from "@tanstack/react-query";
import { mfaDisable } from "@/lib/api/account";
import { accountKeys } from "./use-account-me";

/**
 * Wraps `POST /account/security/mfa/disable` (techplan account/06-mfa-
 * totp, task-01). Same shape for both call sites — `email_password`
 * users pass `{ password }`, Google-only users call `.mutate()` with no
 * argument at all (`mfaDisable()` sends no body in that case) — the
 * branch selection lives entirely in the caller (`MfaDisableForm`, task
 * 03), not here. On success, `mfa_enabled` becomes `false` — invalidate
 * `account.me` so the parent section re-renders into the not-enrolled
 * branch.
 */
export function useMfaDisable() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: mfaDisable,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: accountKeys.me() });
    },
  });
}
