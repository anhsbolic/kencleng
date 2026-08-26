"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import Link from "next/link";
import { useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { Banner } from "@/components/ui/banner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ApiError } from "@/lib/api/client";
import { useForgotPassword } from "@/lib/hooks/use-forgot-password";
import { forgotPasswordSchema, type ForgotPasswordFormValues } from "./forgot-password-schema";

// TBD — placeholder pending product copy, same treatment as every other
// generic string in this codebase (RegisterForm/LoginForm's own
// GENERIC_ERROR_MESSAGE, VerifyEmailStatus's INVALID_LINK_MESSAGE).
const GENERIC_ERROR_MESSAGE = "Terjadi kesalahan. Silakan coba lagi.";
// Reused verbatim from the backend's own real 202 response text
// (auth_password_reset.go:61) — already correct Indonesian, so this is
// the literal fallback if `message` is ever absent, not an override
// (techplan account/04, R3 — D6 doesn't apply to the success branch).
const DEFAULT_SUCCESS_MESSAGE = "Kalau email terdaftar, instruksi sudah dikirim.";
// D6 — never the backend's raw (confirmed English) rate-limit detail.
const RATE_LIMITED_MESSAGE = "Terlalu banyak percobaan. Coba lagi beberapa saat lagi.";

/**
 * `/forgot-password`'s form (techplan account/04, Task: forgot/reset
 * password). Every backend branch behind `POST /auth/forgot-password`
 * collapses to the same `202` (anti-enumeration, R3) — this component must
 * never differentiate them in copy or UI, matching `RegisterForm`'s own
 * discipline for the same reason.
 */
export function ForgotPasswordForm() {
  const forgotPasswordMutation = useForgotPassword();
  const [submitted, setSubmitted] = useState(false);
  const [successMessage, setSuccessMessage] = useState<string | undefined>();
  const [requestError, setRequestError] = useState<string | null>(null);
  const successHeadingRef = useRef<HTMLHeadingElement>(null);

  const {
    register,
    handleSubmit,
    setError,
    formState: { errors, isSubmitting },
  } = useForm<ForgotPasswordFormValues>({
    resolver: zodResolver(forgotPasswordSchema),
  });

  // Focus moves into the success region once the form is replaced, same
  // convention as RegisterForm's own success view.
  useEffect(() => {
    if (submitted) {
      successHeadingRef.current?.focus();
    }
  }, [submitted]);

  async function onSubmit(values: ForgotPasswordFormValues) {
    setRequestError(null);

    const result = await forgotPasswordMutation.mutateAsync(values).catch((error: unknown) => {
      // R5/R6 — a documented 429 shows a frontend-owned message (D6, never
      // the backend's confirmed-English detail); anything else (network
      // failure, unexpected 5xx) falls back to the same generic message.
      setRequestError(
        error instanceof ApiError && error.status === 429
          ? RATE_LIMITED_MESSAGE
          : GENERIC_ERROR_MESSAGE
      );
      return null;
    });

    if (!result) return;

    if (result.ok) {
      // R3 — fixed success view, never conditioned on which internal
      // branch actually fired.
      setSuccessMessage(result.message ?? DEFAULT_SUCCESS_MESSAGE);
      setSubmitted(true);
      return;
    }

    // R4 — defensive: the real backend can reject a malformed email with a
    // 422 that client-side zod should already have prevented in normal
    // use. Frontend-owned copy, never the backend's literal English
    // message (D6) — the same string R2's zod rule already uses.
    for (const { field } of result.errors) {
      if (field === "email") {
        setError("email", { message: "Format email tidak valid" });
      }
    }
  }

  if (submitted) {
    return (
      <div className="flex flex-col gap-6">
        <h2
          ref={successHeadingRef}
          tabIndex={-1}
          className="text-xl font-semibold text-neutral-900 outline-none"
        >
          Cek email kamu
        </h2>
        <Banner variant="success">{successMessage ?? DEFAULT_SUCCESS_MESSAGE}</Banner>
        <Link href="/login" className="text-center text-sm font-semibold text-primary-700">
          Kembali ke halaman login
        </Link>
      </div>
    );
  }

  return (
    <form className="flex flex-col gap-6" onSubmit={handleSubmit(onSubmit)} noValidate>
      {requestError && <Banner variant="error">{requestError}</Banner>}

      <div className="flex flex-col gap-1.5">
        <Label htmlFor="forgot-password-email">Email</Label>
        <Input
          id="forgot-password-email"
          type="email"
          autoComplete="email"
          error={errors.email?.message}
          {...register("email")}
        />
      </div>

      <Button
        type="submit"
        loading={isSubmitting || forgotPasswordMutation.isPending}
        className="w-full"
      >
        Kirim link reset
      </Button>

      <Link href="/login" className="text-center text-sm font-semibold text-primary-700">
        Kembali ke halaman login
      </Link>
    </form>
  );
}
