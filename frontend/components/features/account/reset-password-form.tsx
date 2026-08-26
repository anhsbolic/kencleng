"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import Link from "next/link";
import { useSearchParams } from "next/navigation";
import { useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { Banner } from "@/components/ui/banner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ApiError } from "@/lib/api/client";
import { useResetPassword } from "@/lib/hooks/use-reset-password";
import { resetPasswordSchema, type ResetPasswordFormValues } from "./reset-password-schema";

// TBD — placeholder pending product copy, same treatment as every other
// generic string in this codebase.
const GENERIC_ERROR_MESSAGE = "Terjadi kesalahan. Silakan coba lagi.";
// D6 — frontend-owned, never the backend's confirmed-English Problem
// detail ("The verification token was not found."/"...has expired."/"The
// request was invalid."/"Too many requests. Try again later."). D4 — the
// missing-token case (below) and the 404 case share this exact string,
// matching VerifyEmailStatus's own R11 precedent.
const INVALID_LINK_MESSAGE = "Link reset password tidak valid atau sudah digunakan.";
const EXPIRED_LINK_MESSAGE = "Link reset password sudah kedaluwarsa. Silakan minta link baru.";
const WEAK_PASSWORD_MESSAGE =
  "Password tidak memenuhi syarat. Gunakan minimal 8 karakter dan hindari password yang umum digunakan atau pernah bocor.";
const RATE_LIMITED_MESSAGE = "Terlalu banyak percobaan. Coba lagi beberapa saat lagi.";
// Reused verbatim from the backend's own real 200 response text
// (auth_password_reset.go:95) — already correct Indonesian.
const DEFAULT_SUCCESS_MESSAGE = "Password berhasil diubah. Silakan login ulang.";

type Outcome = "form" | "success" | "expired" | "invalid";

/**
 * `/reset-password`'s form (techplan account/04). Combines three shapes
 * already established elsewhere in this domain: `LoginForm`'s
 * banner-first-child + focus-on-banner convention, `VerifyEmailStatus`'s
 * token-from-`useSearchParams()` + status-discriminated outcome
 * convention (including its missing-token == 404 precedent, D4), and
 * `RegisterForm`'s idle/success-view toggle — but does NOT auto-fire on
 * mount like `VerifyEmailStatus`: a new password must be entered first,
 * so the mutation only fires on submit.
 */
export function ResetPasswordForm() {
  const token = useSearchParams().get("token");
  const resetPasswordMutation = useResetPassword();
  const [outcome, setOutcome] = useState<Outcome>(token ? "form" : "invalid");
  const [successMessage, setSuccessMessage] = useState<string | undefined>();
  const [requestError, setRequestError] = useState<string | null>(null);
  const bannerRef = useRef<HTMLDivElement>(null);

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<ResetPasswordFormValues>({
    resolver: zodResolver(resetPasswordSchema),
  });

  // Focus moves into whichever banner is currently shown — matches
  // LoginForm's established convention for request-level banners, and
  // also covers this component's terminal/success outcomes (including the
  // missing-token case, resolved on first render).
  useEffect(() => {
    if (outcome !== "form" || requestError) {
      bannerRef.current?.focus();
    }
  }, [outcome, requestError]);

  async function onSubmit(values: ResetPasswordFormValues) {
    if (!token) return; // unreachable — the form only renders once a token is present (R7)

    setRequestError(null);

    let result;
    try {
      result = await resetPasswordMutation.mutateAsync({
        token,
        new_password: values.new_password,
      });
    } catch (error) {
      // R11/R12 — terminal: the token is genuinely dead, no retry possible.
      if (error instanceof ApiError && error.status === 410) {
        setOutcome("expired");
        return;
      }
      if (error instanceof ApiError && error.status === 404) {
        setOutcome("invalid");
        return;
      }
      // R13-R15 — non-terminal: 422 (weak/breached password, no field
      // data to map — confirmed directly against the real backend, D2/D6),
      // 429, or a generic/network failure. The form stays mounted and
      // interactive either way, and the token is never touched, so the
      // same link can be resubmitted (spec Assumption B — see the
      // techplan's top risk, §7).
      setRequestError(
        error instanceof ApiError && error.status === 422
          ? WEAK_PASSWORD_MESSAGE
          : error instanceof ApiError && error.status === 429
            ? RATE_LIMITED_MESSAGE
            : GENERIC_ERROR_MESSAGE
      );
      return;
    }

    // R10 — inline success view, no auto-redirect (D5).
    setSuccessMessage(result.message ?? DEFAULT_SUCCESS_MESSAGE);
    setOutcome("success");
  }

  if (outcome === "invalid" || outcome === "expired") {
    return (
      <div className="flex flex-col gap-6">
        <div ref={bannerRef} tabIndex={-1} className="outline-none">
          <Banner variant="error">
            {outcome === "expired" ? EXPIRED_LINK_MESSAGE : INVALID_LINK_MESSAGE}
          </Banner>
        </div>
        <Link
          href="/forgot-password"
          className="text-center text-sm font-semibold text-primary-700"
        >
          Minta link baru
        </Link>
      </div>
    );
  }

  if (outcome === "success") {
    return (
      <div className="flex flex-col gap-6">
        <div ref={bannerRef} tabIndex={-1} className="outline-none">
          <Banner variant="success">{successMessage ?? DEFAULT_SUCCESS_MESSAGE}</Banner>
        </div>
        <Link href="/login" className="text-center text-sm font-semibold text-primary-700">
          Masuk sekarang
        </Link>
      </div>
    );
  }

  return (
    <form className="flex flex-col gap-6" onSubmit={handleSubmit(onSubmit)} noValidate>
      {requestError && (
        <div ref={bannerRef} tabIndex={-1} className="outline-none">
          <Banner variant="error">{requestError}</Banner>
        </div>
      )}

      <div className="flex flex-col gap-1.5">
        <Label htmlFor="reset-password-new-password">Password baru</Label>
        <Input
          id="reset-password-new-password"
          type="password"
          autoComplete="new-password"
          placeholder="Minimal 8 karakter"
          error={errors.new_password?.message}
          {...register("new_password")}
        />
      </div>

      <Button
        type="submit"
        loading={isSubmitting || resetPasswordMutation.isPending}
        className="w-full"
      >
        Reset password
      </Button>
    </form>
  );
}
