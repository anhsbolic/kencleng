"use client";

import { useRouter } from "next/navigation";
import { useEffect, type ReactNode } from "react";
import { useAuthStore } from "@/lib/stores/auth-store";

/**
 * Redirects to `/login` the moment `useAuthStore`'s `accessToken`
 * transitions from a real value to `null` — the single place this
 * feature's session-loss redirect lives, regardless of what caused the
 * transition (a failed `coordinatedRefresh` in this tab, another tab's
 * refresh-failure/logout broadcast applied via `AuthBootstrapProvider`'s
 * listener, or this tab's own logout). See techplan account/03-login-
 * session-management D4 for why this is one subscription instead of
 * three separate redirect call sites.
 *
 * Uses Zustand's own `(state, prevState)` pair per subscription
 * callback — no separately-tracked ref needed, and no "what was the
 * baseline at mount" question, since each individual state change
 * already carries its correct previous value from the store itself.
 *
 * A guest who was never authenticated never triggers a redirect: their
 * `prevState.accessToken` is `null` on every change (including the
 * eventual, silent failure of `AuthBootstrapProvider`'s own boot-time
 * refresh attempt — its R10), so the "was previously authenticated"
 * condition below never holds for that session.
 *
 * Mounted as a descendant of `AuthBootstrapProvider` in `app/layout.tsx`
 * — both are root-level session-lifecycle providers, grouped together.
 */
export function SessionGuardProvider({ children }: { children: ReactNode }) {
  const router = useRouter();

  useEffect(() => {
    return useAuthStore.subscribe((state, prevState) => {
      const wasAuthenticated = prevState.accessToken !== null;
      const isNowUnauthenticated = state.accessToken === null;

      if (wasAuthenticated && isNowUnauthenticated) {
        router.push("/login");
      }
    });
  }, [router]);

  return <>{children}</>;
}
