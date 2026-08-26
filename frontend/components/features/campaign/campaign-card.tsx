import { ImageOff } from "lucide-react";
import { ProgressBar } from "@/components/ui/progress-bar";

export interface CampaignCardProps {
  id: string;
  title: string;
  progress: {
    percentage: number;
    donorCount: number;
    daysRemaining: number | null;
  };
}

/**
 * Public, read-only campaign preview card — used by `/`'s highlighted-
 * campaigns section today, and (later) `/campaign`'s full browse page,
 * since both read from the same `GET /campaigns` → `CampaignListItem`
 * shape.
 *
 * Deliberately no `organization` prop (techplan Decision 1):
 * `CampaignListItem` has no organization name/verification status,
 * only `organization_id` — showing either here would mean inventing
 * data the API doesn't provide, or worse, a false "verified" trust
 * signal. See `.local-agents/works/00-shell-landing/2-plan/techplan.md`
 * for the full reasoning.
 *
 * Photo area is a plain, non-interactive placeholder (R14) — the
 * Tier 1 prototype's `image-slot` element renders upload-dropzone
 * chrome (`prototype-reference.md` Known Issue #2), which doesn't
 * belong on a public read-only card. No `Campaign` schema has a media
 * field yet either, so there's nothing real to bind to.
 */
export function CampaignCard({ title, progress }: CampaignCardProps) {
  const pct = Math.round(Math.min(100, Math.max(0, progress.percentage)));
  const reached = pct >= 100;

  return (
    <div className="flex flex-col overflow-hidden rounded-lg border border-neutral-200 bg-white shadow-sm">
      <div className="flex h-44 items-center justify-center bg-neutral-100">
        <ImageOff aria-hidden="true" className="size-8 text-neutral-400" />
      </div>

      <div className="flex flex-1 flex-col gap-3 p-4">
        <h3 className="text-h3 font-bold text-neutral-900">{title}</h3>

        <div className="mt-auto flex flex-col gap-2">
          <ProgressBar value={pct} height={8} />
          <p className="text-caption text-neutral-700">
            {reached ? (
              <span className="font-semibold text-primary-700">Target tercapai</span>
            ) : (
              <span className="font-semibold text-primary-700">{pct}% terkumpul</span>
            )}
            {progress.daysRemaining !== null && ` · ${progress.daysRemaining} hari lagi`}
          </p>
        </div>
      </div>
    </div>
  );
}

/** Loading-state skeleton, shaped like the real card (`patterns.md` §A.1) — not a bare spinner. */
export function CampaignCardSkeleton() {
  return (
    <div
      role="status"
      aria-label="Memuat kampanye"
      className="flex flex-col overflow-hidden rounded-lg border border-neutral-200 bg-white shadow-sm"
    >
      <div className="h-44 animate-pulse bg-neutral-100" />
      <div className="flex flex-col gap-3 p-4">
        <div className="h-5 w-3/4 animate-pulse rounded bg-neutral-100" />
        <div className="h-5 w-1/2 animate-pulse rounded bg-neutral-100" />
        <div className="mt-2 h-2 animate-pulse rounded-full bg-neutral-100" />
        <div className="h-4 w-1/3 animate-pulse rounded bg-neutral-100" />
      </div>
    </div>
  );
}
