import { cn } from "@/lib/cn";
import type { ReactNode } from "react";

export type BadgeTone =
  | "neutral"
  | "success"
  | "error"
  | "warning"
  | "accent"
  | "info";

const toneClasses: Record<BadgeTone, string> = {
  neutral: "bg-neutral-100 text-neutral-700",
  success: "bg-success-50 text-success-700",
  error: "bg-error-50 text-error-700",
  warning: "bg-warning-50 text-warning-700",
  accent: "bg-accent-50 text-accent-600",
  info: "bg-info-50 text-info-700",
};

export interface BadgeProps {
  tone: BadgeTone;
  size?: "sm" | "md";
  children: ReactNode;
  className?: string;
}

/**
 * Status/label pill (`design-guidelines.md` Component Tokens > Badges).
 * Every status enum across the app maps onto one of these five
 * semantic tones (plus `neutral` for draft/pending-style states) —
 * never a new color per status. `50`-shade background + `700`-shade
 * text keeps AA-compliant contrast at `caption` text size.
 */
export function Badge({ tone, size = "md", children, className }: BadgeProps) {
  return (
    <span
      className={cn(
        "inline-flex items-center rounded-full font-medium",
        size === "sm" ? "px-2 py-0.5 text-[11px]" : "px-2.5 py-1 text-caption",
        toneClasses[tone],
        className
      )}
    >
      {children}
    </span>
  );
}
