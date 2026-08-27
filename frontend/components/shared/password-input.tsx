"use client";

import { Eye, EyeOff } from "lucide-react";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input, type InputProps } from "@/components/ui/input";
import { cn } from "@/lib/cn";

export type PasswordInputProps = Omit<InputProps, "type">;

/**
 * Password field with a show/hide toggle — extracted from `LoginForm`'s
 * original inline pattern (techplan account/05-account-linking §9).
 * This task is the second and third call site needing the exact same
 * toggle (`SetPasswordForm`, `UnlinkGoogleForm`), clearing the "second
 * domain needs it" bar for promoting a `components/shared/` primitive
 * per `phase0-shared-infra.md`'s Incremental Growth Rule. `LoginForm`
 * itself is NOT retrofitted to use this — out of this task's scope.
 */
export function PasswordInput({ className, ...props }: PasswordInputProps) {
  const [show, setShow] = useState(false);

  return (
    <div className="flex items-center gap-2">
      <Input type={show ? "text" : "password"} className={cn("flex-1", className)} {...props} />
      <Button
        type="button"
        variant="ghost"
        size="sm"
        aria-label={show ? "Sembunyikan password" : "Tampilkan password"}
        onClick={() => setShow((current) => !current)}
      >
        {show ? (
          <EyeOff aria-hidden="true" className="size-4" />
        ) : (
          <Eye aria-hidden="true" className="size-4" />
        )}
      </Button>
    </div>
  );
}
