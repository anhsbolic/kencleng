import { useMutation, useQueryClient } from "@tanstack/react-query";
import { unlinkGoogle } from "@/lib/api/account";
import { accountKeys } from "./use-account-me";

/**
 * Wraps `POST /account/security/google/unlink` (techplan account/05-
 * account-linking, R17). On success, `auth_providers` loses `"google"`
 * — invalidate `account.me` so `GoogleIdentityControl` re-renders into
 * its "link" state without a page refresh.
 */
export function useUnlinkGoogle() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: unlinkGoogle,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: accountKeys.me() });
    },
  });
}
