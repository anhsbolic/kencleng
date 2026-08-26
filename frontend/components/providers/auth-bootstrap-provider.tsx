"use client";

import { useQueryClient } from "@tanstack/react-query";
import { useEffect, useRef } from "react";
import { tryRefreshOnce } from "@/lib/api/client";
import { accountKeys } from "@/lib/hooks/use-account-me";
import { useAuthStore } from "@/lib/stores/auth-store";

/**
 * Bridges the OAuth callback's HttpOnly-cookie token delivery into
 * `useAuthStore`'s in-memory access token, which `apiFetch` reads for
 * every request (techplan account/02-google-oauth-login-register,
 * D3/R8-R11). Not OAuth-specific plumbing on its own merits: the
 * access token is deliberately in-memory-only
 * (`lib/stores/auth-store.ts`), so *any* page load already starts
 * with `accessToken === null` — this is the general silent-refresh-
 * on-boot mechanism the token-storage model requires regardless of
 * how the session was originally established.
 *
 * Unconditional by design, not tied to detecting "just came from
 * OAuth" — a successful Google login redirects to the bare frontend
 * root with no query signal to detect that by. Every load attempts
 * hydration once; a genuinely logged-out guest simply gets a failed
 * refresh, silently (R10) — this must be indistinguishable from an
 * ordinary page visit.
 */
export function AuthBootstrapProvider({ children }: { children: React.ReactNode }) {
  const queryClient = useQueryClient();
  const attempted = useRef(false);

  useEffect(() => {
    // R11 — at most one attempt per app load, guarded explicitly
    // rather than relying on effect deps (which would re-run if
    // `accessToken` were read reactively here).
    if (attempted.current) return;
    attempted.current = true;

    // R8 — only attempt hydration if nothing has already populated
    // the store (e.g. a prior manual login this same session).
    if (useAuthStore.getState().accessToken !== null) return;

    tryRefreshOnce().then((succeeded) => {
      // R9 — `tryRefreshOnce` already calls `setAccessToken` on
      // success (client.ts); this provider's own job is only the
      // query-cache side of hydration, so a component that queried
      // `account.me` before hydration completed doesn't keep
      // rendering stale logged-out state.
      if (succeeded) {
        queryClient.invalidateQueries({ queryKey: accountKeys.me() });
      }
      // R10 — failure is silent: no error/toast, `accessToken` stays
      // `null`, indistinguishable from an ordinary guest page load.
    });
  }, [queryClient]);

  return <>{children}</>;
}
