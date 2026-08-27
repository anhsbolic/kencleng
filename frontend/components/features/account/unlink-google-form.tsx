"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { Banner } from "@/components/ui/banner";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { PasswordInput } from "@/components/shared/password-input";
import { ApiError } from "@/lib/api/client";
import { useUnlinkGoogle } from "@/lib/hooks/use-unlink-google";
import { unlinkGoogleSchema, type UnlinkGoogleFormValues } from "./unlink-google-schema";

const GENERIC_ERROR_MESSAGE = "Terjadi kesalahan. Silakan coba lagi."; // TBD

/**
 * `/dashboard/security`'s "Lepas Tautan Google" action (techplan
 * account/05-account-linking, R17-R18). `401`/`409` both render the
 * backend's own `.detail` verbatim — no `.type` parsing needed, since
 * the backend's text already distinguishes every case correctly (D4).
 * On success, `useUnlinkGoogle`'s `onSuccess` invalidates the account
 * cache and the parent re-renders into `GoogleIdentityControl`'s
 * "link" state — no local success view needed here.
 */
export function UnlinkGoogleForm() {
  const unlinkMutation = useUnlinkGoogle();
  const [requestError, setRequestError] = useState<string | null>(null);
  const bannerRef = useRef<HTMLDivElement>(null);

  const { register, handleSubmit, formState } = useForm<UnlinkGoogleFormValues>({
    resolver: zodResolver(unlinkGoogleSchema),
  });

  useEffect(() => {
    if (requestError) {
      bannerRef.current?.focus();
    }
  }, [requestError]);

  async function onSubmit(values: UnlinkGoogleFormValues) {
    setRequestError(null);
    try {
      await unlinkMutation.mutateAsync(values);
    } catch (error) {
      // R18 — both 401 and 409 (either case) show the backend's own
      // detail verbatim; only a genuine network/5xx falls back to the
      // generic message.
      setRequestError(
        error instanceof ApiError && error.detail ? error.detail : GENERIC_ERROR_MESSAGE
      );
    }
  }

  return (
    <form className="flex flex-col gap-4" onSubmit={handleSubmit(onSubmit)} noValidate>
      <h3 className="text-lg font-semibold text-neutral-900">Lepas Tautan Google</h3>

      {requestError && (
        <div ref={bannerRef} tabIndex={-1} className="outline-none">
          <Banner variant="error">{requestError}</Banner>
        </div>
      )}

      <div className="flex flex-col gap-1.5">
        <Label htmlFor="unlink-google-password">Konfirmasi password</Label>
        <PasswordInput
          id="unlink-google-password"
          autoComplete="current-password"
          error={formState.errors.password?.message}
          {...register("password")}
        />
      </div>

      <Button
        type="submit"
        variant="destructive"
        loading={formState.isSubmitting || unlinkMutation.isPending}
        className="w-fit"
      >
        Lepas Tautan Google
      </Button>
    </form>
  );
}
