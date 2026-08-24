import { Loader2 } from "lucide-react";
import { cn } from "@/lib/cn";
import type { ComponentPropsWithoutRef } from "react";

/**
 * Inline loading indicator. Purely decorative on its own — pair it
 * with visible or screen-reader-only text (or an `aria-busy` parent,
 * as `Button`'s `loading` prop does) rather than relying on the spin
 * animation alone to convey state.
 */
export function Spinner({
  className,
  ...props
}: ComponentPropsWithoutRef<"svg">) {
  return (
    <Loader2
      aria-hidden="true"
      className={cn("animate-spin", className)}
      {...props}
    />
  );
}
