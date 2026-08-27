import { useMutation, useQueryClient, type QueryClient } from "@tanstack/react-query";
import { setPassword, type SetPasswordResult } from "@/lib/api/account";
import { postAuthChannelMessage } from "@/lib/api/auth-channel";
import { useAuthStore } from "@/lib/stores/auth-store";
import { accountKeys } from "./use-account-me";

/**
 * Shared success handler for `useSetPassword` — extracted (mirroring
 * `use-login.ts`'s `applyLoginSuccess`) so its branching logic is
 * directly unit-testable without rendering a component or a hook
 * harness (techplan account/05-account-linking, R10, D5).
 *
 * Branch 1 (`branch: "added"`) invalidates `account.me` defensively —
 * `auth_providers` may already reflect the newly-created (unverified)
 * identity, confirmed backend-agnostic to `verified_at` (techplan §1
 * finding #1), so refetching is the safe default rather than assuming
 * no change; harmless no-op otherwise.
 *
 * Branch 2 (`branch: "changed"`) clears the session exactly like
 * `useLogout` does — `SessionGuardProvider`'s existing subscription
 * redirects to `/login` on the resulting real→null `accessToken`
 * transition, so this never navigates directly.
 */
export function applySetPasswordSuccess(
  result: SetPasswordResult,
  queryClient: QueryClient
): void {
  if (!result.ok) return;

  if (result.branch === "added") {
    queryClient.invalidateQueries({ queryKey: accountKeys.me() });
    return;
  }

  useAuthStore.getState().clearAccessToken();
  queryClient.clear();
  postAuthChannelMessage({ type: "logged-out" });
}

/** Wraps `POST /account/security/set-password`. */
export function useSetPassword() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: setPassword,
    onSuccess: (result) => applySetPasswordSuccess(result, queryClient),
  });
}
