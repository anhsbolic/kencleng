"use client";

import { Inbox } from "lucide-react";
import { Banner } from "@/components/ui/banner";
import { Button } from "@/components/ui/button";
import {
  CampaignCard,
  CampaignCardSkeleton,
} from "@/components/features/campaign/campaign-card";
import { useCampaigns } from "@/lib/hooks/use-campaigns";

const SKELETON_COUNT = 3;

/**
 * `/`'s highlighted-campaigns section — the one part of the page that
 * needs client-side data fetching, so it's the only `'use client'`
 * leaf on `/` (`server-client-component-boundary.md`; techplan
 * Decision 6). Handles all four List/Browse states (`patterns.md`
 * §A.1, §B): loading (R6), error (R9), empty (R8), success (R7) — plus
 * the stale/revalidating indicator (R16).
 */
export function HighlightedCampaigns() {
  const { data, isLoading, isError, isFetching, refetch } = useCampaigns();
  const revalidating = isFetching && !isLoading;

  return (
    <section
      id="kampanye"
      className="mx-auto flex w-full max-w-[1360px] flex-col gap-4 px-4 py-8 md:gap-5 md:px-6 md:py-14"
    >
      <div className="flex flex-col gap-1.5 md:flex-row md:items-end md:justify-between">
        <div className="flex flex-col gap-1.5">
          <span className="text-[11px] font-bold tracking-wide text-neutral-400 uppercase">
            Kampanye pilihan
          </span>
          <h2 className="text-h2 font-bold text-neutral-900">
            Sedang berjalan minggu ini
          </h2>
        </div>
      </div>

      {isLoading ? (
        <div className="grid grid-cols-1 gap-3 md:grid-cols-3 md:gap-4.5">
          {Array.from({ length: SKELETON_COUNT }).map((_, index) => (
            <CampaignCardSkeleton key={index} />
          ))}
        </div>
      ) : isError ? (
        <Banner variant="error">
          <div className="flex flex-col items-start gap-2">
            <p>Gagal memuat kampanye. Coba lagi.</p>
            <Button variant="outline" size="sm" onClick={() => refetch()}>
              Coba lagi
            </Button>
          </div>
        </Banner>
      ) : !data || data.data.length === 0 ? (
        // No CTA here — Guest can't create a campaign (patterns.md §B
        // Empty row: primary action shown only if the viewer is
        // actually authorized to take it).
        <div className="flex flex-col items-center gap-2 rounded-lg border border-dashed border-neutral-200 py-10 text-center">
          <Inbox aria-hidden="true" className="size-8 text-neutral-400" />
          <p className="text-body text-neutral-500">
            Belum ada kampanye yang ditampilkan.
          </p>
        </div>
      ) : (
        <>
          {revalidating && (
            <p className="text-caption text-neutral-500">
              Data mungkin tidak terbaru.
            </p>
          )}
          <div className="grid grid-cols-1 gap-3 md:grid-cols-3 md:gap-4.5">
            {data.data.map((campaign) => (
              <CampaignCard
                key={campaign.id}
                id={campaign.id ?? ""}
                title={campaign.title ?? "Kampanye"}
                progress={{
                  percentage: campaign.progress?.percentage ?? 0,
                  donorCount: campaign.progress?.donor_count ?? 0,
                  daysRemaining: campaign.progress?.days_remaining ?? null,
                }}
              />
            ))}
          </div>
        </>
      )}
    </section>
  );
}
