import { useQuery } from "@tanstack/react-query";
import { getMe } from "@/lib/api/account";

export const accountKeys = {
  me: () => ["account", "me"] as const,
};

/**
 * The current authenticated user's profile. Shared query key factory
 * (`accountKeys`) so every consumer — `use-has-role`, a future
 * `/dashboard/profile` page, etc. — hits the same cache entry rather
 * than each issuing its own differently-keyed fetch
 * (`data-fetching-conventions.md`).
 */
export function useAccountMe() {
  return useQuery({
    queryKey: accountKeys.me(),
    queryFn: getMe,
  });
}
