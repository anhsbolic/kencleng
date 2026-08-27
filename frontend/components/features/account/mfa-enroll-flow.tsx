"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { Banner } from "@/components/ui/banner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ApiError } from "@/lib/api/client";
import { useMfaEnroll } from "@/lib/hooks/use-mfa-enroll";
import { useMfaEnrollConfirm } from "@/lib/hooks/use-mfa-enroll-confirm";
import { parseOtpauthSecret } from "@/lib/otpauth";
import { QrCode } from "./qr-code";
import {
  mfaEnrollConfirmSchema,
  type MfaEnrollConfirmFormValues,
} from "./mfa-enroll-confirm-schema";

const GENERIC_ERROR_MESSAGE = "Terjadi kesalahan. Silakan coba lagi."; // TBD
const INVALID_CODE_MESSAGE = "Kode tidak valid, coba lagi."; // TBD

export interface MfaEnrollFlowProps {
  /**
   * Fired once `enroll/confirm` succeeds, carrying the 10 one-time
   * backup codes. The parent (`MfaSection`) owns rendering the
   * codes-once view from here on (R9/R12) — this component unmounts
   * its own QR/form UI immediately after calling this.
   */
  onEnrolled: (backupCodes: string[]) => void;
}

/**
 * `/dashboard/security`'s not-enrolled MFA flow (techplan account/06-
 * mfa-totp, task-02): "Aktifkan MFA" trigger → QR + manual-entry secret
 * + confirm-code form → hands the resulting backup codes to the parent
 * via `onEnrolled`. Never auto-fires `enroll` on mount (R4) — no
 * precedent anywhere in this codebase fires a mutation as a render side
 * effect, and `mfa_enabled: false` can't distinguish "never started"
 * from "pending unconfirmed" anyway, so there's nothing reliable to
 * resume from.
 */
export function MfaEnrollFlow({ onEnrolled }: MfaEnrollFlowProps) {
  const [step, setStep] = useState<"idle" | "confirming">("idle");
  const [otpauthUri, setOtpauthUri] = useState<string | null>(null);
  const [requestError, setRequestError] = useState<string | null>(null);
  const enrollMutation = useMfaEnroll();
  const confirmMutation = useMfaEnrollConfirm();
  const bannerRef = useRef<HTMLDivElement>(null);

  const { register, handleSubmit, formState, setError } =
    useForm<MfaEnrollConfirmFormValues>({
      resolver: zodResolver(mfaEnrollConfirmSchema),
    });

  useEffect(() => {
    if (requestError) {
      bannerRef.current?.focus();
    }
  }, [requestError]);

  async function handleActivate() {
    setRequestError(null);
    try {
      const response = await enrollMutation.mutateAsync();
      setOtpauthUri(response.otpauth_uri);
      setStep("confirming");
    } catch (error) {
      // R6 (409 — already enabled, defensive: this trigger only renders
      // in the not-enrolled branch, so in normal single-tab use this is
      // unreachable) / R7 (network/5xx) — both show a generic banner.
      setRequestError(
        error instanceof ApiError && error.detail ? error.detail : GENERIC_ERROR_MESSAGE
      );
    }
  }

  async function onSubmitConfirm(values: MfaEnrollConfirmFormValues) {
    setRequestError(null);
    try {
      const result = await confirmMutation.mutateAsync(values);
      onEnrolled(result.backup_codes);
    } catch (error) {
      // R10 — 422 shows one fixed field-level message; QR/form stay
      // mounted and interactive (no remount, no re-fetch of enroll —
      // the pending secret is not discarded per the spec). R11 —
      // network/5xx shows a generic request-level banner instead.
      if (error instanceof ApiError && error.status === 422) {
        setError("totp_code", { message: INVALID_CODE_MESSAGE });
        return;
      }
      setRequestError(GENERIC_ERROR_MESSAGE);
    }
  }

  if (step === "idle") {
    return (
      <div className="flex flex-col gap-4">
        {requestError && (
          <div ref={bannerRef} tabIndex={-1} className="outline-none">
            <Banner variant="error">{requestError}</Banner>
          </div>
        )}
        <Button
          type="button"
          onClick={handleActivate}
          loading={enrollMutation.isPending}
          className="w-fit"
        >
          Aktifkan MFA
        </Button>
      </div>
    );
  }

  const manualSecret = otpauthUri ? parseOtpauthSecret(otpauthUri) : null;

  return (
    <form className="flex flex-col gap-4" onSubmit={handleSubmit(onSubmitConfirm)} noValidate>
      {requestError && (
        <div ref={bannerRef} tabIndex={-1} className="outline-none">
          <Banner variant="error">{requestError}</Banner>
        </div>
      )}

      <QrCode value={otpauthUri ?? ""} />

      {/* R24 — manual-entry fallback for users without a camera-capable
          device or using a screen reader; hidden gracefully if the URI
          carries no parseable secret (defensive, D11). */}
      {manualSecret && (
        <p className="text-sm text-neutral-500">
          Tidak bisa scan? Masukkan kode ini secara manual:{" "}
          <code className="select-all font-mono text-neutral-700">{manualSecret}</code>
        </p>
      )}

      <div className="flex flex-col gap-1.5">
        <Label htmlFor="mfa-enroll-confirm-code">Kode OTP</Label>
        <Input
          id="mfa-enroll-confirm-code"
          autoComplete="one-time-code"
          error={formState.errors.totp_code?.message}
          {...register("totp_code")}
        />
      </div>

      <Button
        type="submit"
        loading={formState.isSubmitting || confirmMutation.isPending}
        className="w-fit"
      >
        Konfirmasi
      </Button>
    </form>
  );
}
