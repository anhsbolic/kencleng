import { useMutation, useQueryClient, type QueryClient } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { login, type LoginResult } from "@/lib/api/account";
import { useAuthStore } from "@/lib/stores/auth-store";
import { accountKeys } from "./use-account-me";

// `next/navigation` doesn't re-export `AppRouterInstance` as a public
// type — derive it from `useRouter`'s own return type instead of
// reaching into Next's internal module path.
type AppRouter = ReturnType<typeof useRouter>;

/**
 * Shared success handler for both login-completing endpoints
 * (`POST /auth/login`'s "ok" branch and `POST /auth/login/mfa`'s only
 * branch) — R3/R7 (techplan account/03-login-session-management,
 * task-01). Exported so `useLoginMfa` reuses this exact implementation
 * instead of a second, drift-prone copy.
 *
 * A no-op for the `mfa_required` branch (R4) — that branch carries no
 * `access_token`/`user` to apply, and `LoginForm` (task-03) owns the
 * response to it (transitioning to the MFA step), not this hook.
 */
export function applyLoginSuccess(
  result: LoginResult,
  queryClient: QueryClient,
  router: AppRouter
): void {
  if (result.status !== "ok") return;

  useAuthStore.getState().setAccessToken(result.access_token);
  // D7 — direct cache write, not invalidate: LoginResponse.user is
  // structurally identical to what GET /account/me returns, so there's
  // nothing to gain from an extra round-trip.
  queryClient.setQueryData(accountKeys.me(), result.user);
  // Open Item #1 (source techplan) — redirect target pending
  // confirmation; `/dashboard/profile` chosen by analogy to the
  // Dashboard Shell's own logo-link target.
  router.push("/dashboard/profile");
}

/**
 * Wraps `POST /auth/login`. No query-key invalidation beyond the direct
 * cache write above — nothing else is stale as a result of logging in
 * (`data-fetching-conventions.md`).
 */
export function useLogin() {
  const queryClient = useQueryClient();
  const router = useRouter();

  return useMutation({
    mutationFn: login,
    onSuccess: (result) => applyLoginSuccess(result, queryClient, router),
  });
}
