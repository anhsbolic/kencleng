"use client";

import { useId, useState } from "react";
import { Banner } from "@/components/ui/banner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ApiError } from "@/lib/api/client";
import { useResendVerification } from "@/lib/hooks/use-resend-verification";

const GENERIC_ERROR_MESSAGE = "Terjadi kesalahan. Silakan coba lagi.";

export interface ResendVerificationControlProps {
  /**
   * Pre-fills the email field where the caller already knows it (e.g.
   * the register success view, which just collected it). Left blank
   * when the caller has no email to offer (e.g. `/verify-email`'s
   * expired-token view, which only has a token) — the field stays
   * editable either way, since resend's `202` response never reveals
   * whether the address actually matched anything (anti-enumeration).
   */
  defaultEmail?: string;
}

/**
 * Shared "Kirim ulang" (resend verification) affordance — mounted from
 * both the register success view and `/verify-email`'s expired-token
 * view (techplan account/01-register-email-verification, Decision
 * D2/D4). Always renders the backend's own generic confirmation text
 * on success (rule R9) and never varies based on whether the email
 * actually matched anything server-side.
 */
export function ResendVerificationControl({
  defaultEmail = "",
}: ResendVerificationControlProps) {
  const resend = useResendVerification();
  const [email, setEmail] = useState(defaultEmail);
  const inputId = useId();

  if (resend.isSuccess) {
    return (
      <Banner variant="success">
        {resend.data.message ??
          "Kalau email terdaftar, instruksi sudah dikirim."}
      </Banner>
    );
  }

  return (
    <div className="flex flex-col gap-2">
      {resend.isError && (
        <Banner variant="error">
          {resend.error instanceof ApiError && resend.error.status === 429
            ? (resend.error.detail ?? GENERIC_ERROR_MESSAGE)
            : GENERIC_ERROR_MESSAGE}
        </Banner>
      )}
      <div className="flex flex-col gap-1.5">
        <Label htmlFor={inputId}>Email</Label>
        <Input
          id={inputId}
          type="email"
          value={email}
          onChange={(event) => setEmail(event.target.value)}
        />
      </div>
      <Button
        type="button"
        variant="ghost"
        size="sm"
        loading={resend.isPending}
        disabled={!email}
        onClick={() => resend.mutate({ email })}
      >
        Kirim ulang
      </Button>
    </div>
  );
}
