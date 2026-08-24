import { cn } from "@/lib/cn";
import { useId, type InputHTMLAttributes } from "react";

export interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  /**
   * Field-level validation message (`patterns.md` Form Page
   * "Validating" row). Rendered below the field and wired to it via
   * `aria-describedby` + `aria-invalid` — never just visual
   * proximity (`accessibility-fundamentals.md`).
   */
  error?: string;
}

export function Input({ error, className, id, ...props }: InputProps) {
  const generatedId = useId();
  const inputId = id ?? generatedId;
  const errorId = `${inputId}-error`;

  return (
    <div className="w-full">
      <input
        id={inputId}
        aria-invalid={Boolean(error) || undefined}
        aria-describedby={error ? errorId : undefined}
        className={cn(
          "block w-full rounded-md border bg-neutral-100 px-3 text-base text-neutral-900 h-11",
          "placeholder:text-neutral-400",
          "focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-100 focus-visible:border-primary-500",
          "disabled:bg-neutral-50 disabled:text-neutral-400 disabled:cursor-not-allowed",
          error ? "border-error-500" : "border-neutral-200",
          className
        )}
        {...props}
      />
      {error && (
        <p id={errorId} className="mt-1 text-sm text-error-700">
          {error}
        </p>
      )}
    </div>
  );
}
