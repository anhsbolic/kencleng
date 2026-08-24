import { AlertCircle, AlertTriangle, CheckCircle2, Info } from "lucide-react";
import { cn } from "@/lib/cn";
import type { ReactNode } from "react";

type BannerVariant = "success" | "error" | "warning" | "info";

const variantConfig: Record<
  BannerVariant,
  { className: string; icon: typeof Info; role: "status" | "alert" }
> = {
  success: {
    className: "bg-success-50 text-success-700",
    icon: CheckCircle2,
    role: "status",
  },
  error: {
    className: "bg-error-50 text-error-700",
    icon: AlertCircle,
    role: "alert",
  },
  warning: {
    className: "bg-warning-50 text-warning-700",
    icon: AlertTriangle,
    role: "alert",
  },
  info: {
    className: "bg-info-50 text-info-700",
    icon: Info,
    role: "status",
  },
};

export interface BannerProps {
  variant: BannerVariant;
  children: ReactNode;
  className?: string;
}

/**
 * Request-level success/error banner (`patterns.md` Form Page
 * pattern — the shell every `(auth)` form leaves a slot for, so the
 * generic-auth-failure case never gets rendered as a field-level
 * error instead, per `prototype-reference.md`'s `/login` known
 * issue). `error`/`warning` use `role="alert"` (assertive —
 * interrupts, appropriate for a failure the user must notice
 * immediately); `success`/`info` use `role="status"` (polite).
 */
export function Banner({ variant, children, className }: BannerProps) {
  const { className: variantClassName, icon: Icon, role } = variantConfig[variant];

  return (
    <div
      role={role}
      className={cn(
        "flex items-start gap-2 rounded-md px-4 py-3 text-sm",
        variantClassName,
        className
      )}
    >
      <Icon aria-hidden="true" className="mt-0.5 size-4 shrink-0" />
      <div>{children}</div>
    </div>
  );
}
