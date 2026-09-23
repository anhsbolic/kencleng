import { useQuery } from "@tanstack/react-query";
import { getPublicCampaignDetail } from "@/lib/api/public-campaign";

export const campaignKeys = {
  all: ["campaign"] as const,
  detail: (campaignId: string) => [...campaignKeys.all, "detail", campaignId] as const,
};

export function usePublicCampaignDetail(campaignId: string) {
  return useQuery({
    queryKey: campaignKeys.detail(campaignId),
    queryFn: () => getPublicCampaignDetail(campaignId),
    retry: false,
    refetchOnWindowFocus: false,
    refetchInterval: false,
  });
}
