"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { Banner } from "@/components/ui/banner";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { PasswordInput } from "@/components/shared/password-input";
import { ApiError } from "@/lib/api/client";
import { useMfaDisable } from "@/lib/hooks/use-mfa-disable";
import { GoogleAuthButton } from "./google-auth-button";
import { mfaDisableSchema, type MfaDisableFormValues } from "./mfa-disable-schema";

const GENERIC_ERROR_MESSAGE = "Terjadi kesalahan. Silakan coba lagi."; // TBD
// New frontend-owned copy (techplan D10) — no regenerate endpoint
// exists on the backend, only disable → re-enroll (page-map.md's
// "regenerate" wording does not correspond to a real capability).
const REGENERATE_NOTE =
  "Untuk mendapatkan kode cadangan baru, nonaktifkan MFA lalu aktifkan kembali."; // TBD

export interface MfaDisableFormProps {
  hasEmailPassword: boolean;
}

/**
 * `/dashboard/security`'s enrolled-state MFA disable action (techplan
 * account/06-mfa-totp, task-03). Branches on `hasEmailPassword` (a prop
 * from the parent, `MfaSection` — mirrors `GoogleIdentityControl`'s
 * prop-driven-by-parent convention): a password-confirm form (mirrors
 * `UnlinkGoogleForm` almost exactly), or a single re-auth-aware button
 * for Google-only users (techplan D6, Option B — always render the
 * button, let the documented `401` drive the reauth prompt; no
 * query-param-driven redirect handling is built).
 */
export function MfaDisableForm({ hasEmailPassword }: MfaDisableFormProps) {
  const disableMutation = useMfaDisable();
  const [requestError, setRequestError] = useState<string | null>(null);
  const [needsReauth, setNeedsReauth] = useState(false);
  const bannerRef = useRef<HTMLDivElement>(null);

  const { register, handleSubmit, formState } = useForm<MfaDisableFormValues>({
    resolver: zodResolver(mfaDisableSchema),
  });

  useEffect(() => {
    if (requestError) {
      bannerRef.current?.focus();
    }
  }, [requestError]);

  async function onSubmitPassword(values: MfaDisableFormValues) {
    setRequestError(null);
    try {
      // R15 — success: no local view needed, the parent (`MfaSection`)
      // re-renders into the not-enrolled branch once `mfa_enabled`
      // flips via useMfaDisable's own cache invalidation.
      await disableMutation.mutateAsync(values);
    } catch (error) {
      // R16 — .detail shown verbatim (undifferentiated 401 in the
      // schema, only one message string to expect here).
      setRequestError(
        error instanceof ApiError && error.detail ? error.detail : GENERIC_ERROR_MESSAGE
      );
    }
  }

  async function handleGoogleOnlyDisable() {
    setRequestError(null);
    try {
      // R18 — success: same no-local-view treatment as R15.
      await disableMutation.mutateAsync();
    } catch (error) {
      // R19 — marker missing/expired: banner + reauth prompt, button
      // stays available to retry after the user returns.
      setNeedsReauth(true);
      setRequestError(
        error instanceof ApiError && error.detail ? error.detail : GENERIC_ERROR_MESSAGE
      );
    }
  }

  return (
    <div className="flex flex-col gap-4">
      {requestError && (
        <div ref={bannerRef} tabIndex={-1} className="outline-none">
          <Banner variant="error">{requestError}</Banner>
        </div>
      )}

      {hasEmailPassword ? (
        <form
          className="flex flex-col gap-4"
          onSubmit={handleSubmit(onSubmitPassword)}
          noValidate
        >
          <div className="flex flex-col gap-1.5">
            <Label htmlFor="mfa-disable-password">Konfirmasi password</Label>
            <PasswordInput
              id="mfa-disable-password"
              autoComplete="current-password"
              error={formState.errors.password?.message}
              {...register("password")}
            />
          </div>

          <Button
            type="submit"
            variant="destructive"
            loading={formState.isSubmitting || disableMutation.isPending}
            className="w-fit"
          >
            Nonaktifkan MFA
          </Button>
        </form>
      ) : (
        <div className="flex flex-col gap-3">
          {needsReauth && (
            <GoogleAuthButton intent="reauth" label="Verifikasi ulang dengan Google" />
          )}
          <Button
            type="button"
            variant="destructive"
            onClick={handleGoogleOnlyDisable}
            loading={disableMutation.isPending}
            className="w-fit"
          >
            Nonaktifkan MFA
          </Button>
        </div>
      )}

      <p className="text-sm text-neutral-500">{REGENERATE_NOTE}</p>
    </div>
  );
}
