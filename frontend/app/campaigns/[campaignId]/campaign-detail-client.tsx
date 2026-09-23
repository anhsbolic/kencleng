"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useEffect, useRef, useState } from "react";
import { usePublicCampaignDetail } from "@/lib/hooks/use-public-campaign-detail";
import CampaignDetailView from "./campaign-detail-view";

export default function CampaignDetailClient({ campaignId }: { campaignId: string }) {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: { queries: { retry: false, refetchOnWindowFocus: false } },
      }),
  );

  return (
    <QueryClientProvider client={queryClient}>
      <CampaignDetailQuery campaignId={campaignId} />
    </QueryClientProvider>
  );
}

function CampaignDetailQuery({ campaignId }: { campaignId: string }) {
  const query = usePublicCampaignDetail(campaignId);
  const mainRef = useRef<HTMLElement>(null);
  const previousState = useRef<string | null>(null);

  const state = query.isPending
    ? "loading"
    : query.isError
      ? "request-failure"
      : query.data.kind;

  useEffect(() => {
    if (state === "success" && previousState.current === "request-failure") {
      mainRef.current?.focus();
    }
    previousState.current = state;
  }, [state]);

  if (query.isPending) {
    return <CampaignDetailView mainRef={mainRef} state="loading" />;
  }

  if (query.isError) {
    return (
      <CampaignDetailView
        isRetrying={query.isFetching}
        mainRef={mainRef}
        onRetry={() => void query.refetch()}
        state="request-failure"
      />
    );
  }

  if (query.data.kind === "not-found") {
    return <CampaignDetailView mainRef={mainRef} state="not-found" />;
  }

  if (query.data.kind === "temporarily-unavailable") {
    return (
      <CampaignDetailView
        isRetrying={query.isFetching}
        mainRef={mainRef}
        onRetry={() => void query.refetch()}
        state="temporarily-unavailable"
      />
    );
  }

  return <CampaignDetailView campaign={query.data.data} mainRef={mainRef} state="success" />;
}
