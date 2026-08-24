import { cn } from "@/lib/cn";
import type { LabelHTMLAttributes } from "react";

export function Label({
  className,
  children,
  ...props
}: LabelHTMLAttributes<HTMLLabelElement>) {
  return (
    <label
      className={cn("block text-sm font-medium text-neutral-700", className)}
      {...props}
    >
      {children}
    </label>
  );
}
