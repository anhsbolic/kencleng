import { cn } from "@/lib/cn";
import { Spinner } from "./spinner";
import type { ButtonHTMLAttributes } from "react";

type ButtonVariant = "primary" | "secondary" | "outline" | "ghost" | "destructive";
type ButtonSize = "sm" | "md" | "lg";

const variantClasses: Record<ButtonVariant, string> = {
  primary: "bg-primary-600 text-white shadow-sm hover:bg-primary-700",
  secondary: "bg-accent-500 text-neutral-900 shadow-sm hover:bg-accent-600",
  outline:
    "bg-transparent text-neutral-700 border border-neutral-200 hover:bg-neutral-100",
  ghost: "bg-transparent text-primary-700 hover:bg-primary-100",
  destructive: "bg-error-500 text-white shadow-sm hover:bg-error-700",
};

// Small 36px, Medium 44px (default), Large 52px — design-guidelines.md
// Component Tokens > Buttons "Size tokens" table.
const sizeClasses: Record<ButtonSize, string> = {
  sm: "h-9 px-3 text-sm",
  md: "h-11 px-5 text-base",
  lg: "h-13 px-6 text-lg",
};

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: ButtonVariant;
  size?: ButtonSize;
  /**
   * Submit-in-progress state (Form Page pattern's "Submitting" row):
   * disables the button and shows an inline spinner, keeping the
   * label visible so the accessible name doesn't change mid-flow.
   */
  loading?: boolean;
}

export function Button({
  variant = "primary",
  size = "md",
  loading = false,
  disabled,
  className,
  children,
  ...props
}: ButtonProps) {
  return (
    <button
      type="button"
      disabled={disabled || loading}
      aria-busy={loading || undefined}
      className={cn(
        "inline-flex items-center justify-center gap-2 rounded-md font-semibold transition-colors",
        "focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-100 focus-visible:ring-offset-2",
        "disabled:cursor-not-allowed disabled:opacity-60",
        variantClasses[variant],
        sizeClasses[size],
        className
      )}
      {...props}
    >
      {loading && <Spinner className="size-4" aria-hidden="true" />}
      {children}
    </button>
  );
}
