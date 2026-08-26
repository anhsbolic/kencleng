import { useMutation } from "@tanstack/react-query";
import { register } from "@/lib/api/account";

/**
 * Wraps `POST /auth/register`. No query-key factory / invalidation —
 * nothing is cached anywhere for an unauthenticated user at this point
 * in the flow, so there's nothing this mutation needs to invalidate
 * (`data-fetching-conventions.md`, confirmed explicitly rather than
 * silently skipped).
 */
export function useRegister() {
  return useMutation({
    mutationFn: register,
  });
}
