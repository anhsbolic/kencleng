"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useState } from "react";

/**
 * Wraps `children` in a TanStack Query `QueryClientProvider`.
 *
 * `QueryClient` isn't serializable across the Server/Client boundary,
 * so this provider — and the instance it holds — must live in a
 * Client Component. The instance is created once per component
 * mount via `useState`'s lazy initializer, not at module scope, so
 * each request gets its own client on the server and it's stable
 * across re-renders on the client.
 */
export function QueryProvider({ children }: { children: React.ReactNode }) {
  const [queryClient] = useState(() => new QueryClient());

  return (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
}
