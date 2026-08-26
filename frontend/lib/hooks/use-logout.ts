import { useMutation, useQueryClient } from "@tanstack/react-query";
import { postAuthChannelMessage } from "@/lib/api/auth-channel";
import { logout } from "@/lib/api/account";
import { useAuthStore } from "@/lib/stores/auth-store";

/**
 * Wraps `POST /auth/logout` (R17, techplan account/03-login-session-
 * management, task-04). Cleanup runs in `onSettled` — unconditionally,
 * on both success and failure — never `onSuccess` alone: spec 03
 * defines logout as idempotent/always-`204` from the client's own
 * perspective (D5), so "did the network call succeed" isn't something
 * the caller needs to gate local cleanup on. Does **not** navigate —
 * `SessionGuardProvider` (task-02) is the sole redirect path for any
 * authenticated→unauthenticated transition, this hook only needs to
 * cause that transition.
 */
export function useLogout() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: logout,
    onSettled: () => {
      useAuthStore.getState().clearAccessToken();
      // pwa/state-management-boundaries.md — full cache reset, not just
      // account.me, so no other query's stale data lingers into the
      // next session.
      queryClient.clear();
      postAuthChannelMessage({ type: "logged-out" });
    },
  });
}
