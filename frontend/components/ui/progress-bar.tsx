import { cn } from "@/lib/cn";

export interface ProgressBarProps {
  /** 0-100. Values outside that range are clamped for display. */
  value: number;
  /** Track/fill height in pixels. Design-guidelines.md's minimum is 0.75rem (12px). */
  height?: number;
  className?: string;
}

/**
 * Donation progress bar (`design-guidelines.md` Component Tokens >
 * Progress bar) — the single most benchmark-sensitive component on
 * public campaign pages. Fill switches from `primary-600` to
 * `success-500` once the goal is reached (`value >= 100`), but that
 * color change is never the only signal a caller relies on — every
 * consumer of this component (e.g. `CampaignCard`) pairs it with a
 * visible "X% terkumpul" text line, so a color-blind viewer still
 * gets an unambiguous "goal reached" signal from the number itself
 * (`accessibility-fundamentals.md` — never color alone).
 */
export function ProgressBar({ value, height = 12, className }: ProgressBarProps) {
  const clamped = Math.min(100, Math.max(0, value));
  const reached = clamped >= 100;

  return (
    <div
      role="progressbar"
      aria-valuenow={clamped}
      aria-valuemin={0}
      aria-valuemax={100}
      className={cn("w-full overflow-hidden rounded-full bg-neutral-200", className)}
      style={{ height }}
    >
      <div
        className={cn("h-full rounded-full", reached ? "bg-success-500" : "bg-primary-600")}
        style={{ width: `${clamped}%` }}
      />
    </div>
  );
}
