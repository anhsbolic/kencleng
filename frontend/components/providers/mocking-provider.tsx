"use client";

import { useEffect, useState } from "react";

/**
 * Gates the app behind the MSW browser worker when
 * `NEXT_PUBLIC_API_MOCKING=true` — mock-first dev mode (standalone
 * `npm run dev`, no backend reachable). Off by default: real backend
 * through Caddy/docker-compose never goes through this.
 *
 * Two gates, not one: the env var alone isn't trusted, because
 * `NEXT_PUBLIC_*` vars are bundled into client JS and readable by
 * anyone. `process.env.NODE_ENV === 'production'` always wins,
 * regardless of what the env var says, so a stale `.env` value can
 * never ship mock interception to real users.
 */
export function MockingProvider({ children }: { children: React.ReactNode }) {
  const shouldMock =
    process.env.NEXT_PUBLIC_API_MOCKING === "true" &&
    process.env.NODE_ENV !== "production";

  const [ready, setReady] = useState(!shouldMock);

  useEffect(() => {
    if (!shouldMock) return;

    let cancelled = false;

    import("@/mocks/browser").then(({ worker }) =>
      worker.start({ onUnhandledRequest: "bypass" }).then(() => {
        if (!cancelled) setReady(true);
      })
    );

    return () => {
      cancelled = true;
    };
  }, [shouldMock]);

  if (!ready) return null;

  return <>{children}</>;
}
