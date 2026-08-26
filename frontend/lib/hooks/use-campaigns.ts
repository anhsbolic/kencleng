import { useQuery } from "@tanstack/react-query";
import { getCampaigns } from "@/lib/api/campaign";

export const campaignKeys = {
  list: () => ["campaigns", "list"] as const,
};

/**
 * `/`'s highlighted-campaigns section and (later) `/campaign`'s browse
 * page both read from this same query key — one cache entry, not a
 * differently-keyed fetch per consumer (`data-fetching-conventions.md`).
 *
 * `staleTime` is set (rather than left at the default 0) so
 * `isFetching`-while-`!isLoading` genuinely means "revalidating in the
 * background while showing prior data" (R16's freshness indicator) —
 * at the default staleTime every mount would count as "stale",
 * making that indicator fire constantly instead of only when it's
 * actually meaningful.
 */
export function useCampaigns() {
  return useQuery({
    queryKey: campaignKeys.list(),
    queryFn: getCampaigns,
    staleTime: 60_000,
  });
}
