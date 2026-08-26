"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { Eye, EyeOff } from "lucide-react";
import Link from "next/link";
import { useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { Banner } from "@/components/ui/banner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ApiError } from "@/lib/api/client";
import { useLogin } from "@/lib/hooks/use-login";
import { useLoginMfa } from "@/lib/hooks/use-login-mfa";
import { GoogleAuthButton } from "./google-auth-button";
import {
  loginMfaSchema,
  loginSchema,
  type LoginFormValues,
  type LoginMfaFormValues,
} from "./login-schema";

const GENERIC_ERROR_MESSAGE = "Terjadi kesalahan. Silakan coba lagi."; // TBD — placeholder pending product copy, same treatment as RegisterForm's GENERIC_ERROR_MESSAGE

export interface LoginFormProps {
  /**
   * When rendered inside the landing-page auth modal, the "Belum punya
   * akun? Daftar" link must switch the modal's mode instead of
   * navigating to `/register` (which would leave the landing page).
   * Omit for the standalone `/login` route, where a real navigation is
   * correct.
   */
  onSwitchToRegister?: () => void;
}

/**
 * `/login`'s form (techplan account/03-login-session-management,
 * task-03). Two steps in one component: password credentials (R1-R3,
 * R5), then — only if the account has MFA enrolled — a TOTP/backup-code
 * challenge (R6-R8). `GoogleAuthButton` and its divider are
 * password-step-only (R9). The request-level error banner is rendered
 * as this component's first child regardless of step, never as a
 * field-level error on the email/password/code input — this is the
 * structural fix for the confirmed, not-fixed Known Issue #1 from the
 * Tier 1 prototype (`design-reference/login-register.html`).
 *
 * Both `useLogin`/`useLoginMfa` (lib/hooks/) already perform the
 * store-write + cache-write + redirect side effects on their own
 * success — this component's job on that branch is simply to let the
 * resulting navigation/unmount happen, never to repeat that logic.
 */
export function LoginForm({ onSwitchToRegister }: LoginFormProps = {}) {
  const [step, setStep] = useState<"password" | "mfa">("password");
  const [mfaPendingToken, setMfaPendingToken] = useState<string | null>(null);
  const [requestError, setRequestError] = useState<string | null>(null);
  const [showPassword, setShowPassword] = useState(false);
  const [useBackupCode, setUseBackupCode] = useState(false);
  const bannerRef = useRef<HTMLDivElement>(null);

  const loginMutation = useLogin();
  const loginMfaMutation = useLoginMfa();

  const passwordForm = useForm<LoginFormValues>({ resolver: zodResolver(loginSchema) });
  const mfaForm = useForm<LoginMfaFormValues>({ resolver: zodResolver(loginMfaSchema) });

  // Focus moves into the error banner on render, matching the
  // convention already established by RegisterForm/GoogleCallbackError
  // elsewhere in this codebase (accessibility-fundamentals.md).
  useEffect(() => {
    if (requestError) bannerRef.current?.focus();
  }, [requestError]);

  async function onSubmitPassword(values: LoginFormValues) {
    setRequestError(null);

    let result;
    try {
      result = await loginMutation.mutateAsync(values);
    } catch (error) {
      setRequestError(
        error instanceof ApiError && error.detail ? error.detail : GENERIC_ERROR_MESSAGE
      );
      return;
    }

    // R4 — "ok" branch already redirected via useLogin's onSuccess;
    // "mfa_required" branch transitions this component to the MFA step.
    if (result.status === "mfa_required") {
      setMfaPendingToken(result.mfa_pending_token);
      setStep("mfa");
    }
  }

  async function onSubmitMfa(values: LoginMfaFormValues) {
    if (!mfaPendingToken) return; // defensive — unreachable, step only becomes 'mfa' with a token already set

    setRequestError(null);

    try {
      await loginMfaMutation.mutateAsync({
        mfa_pending_token: mfaPendingToken,
        totp_code: values.totp_code,
        backup_code: values.backup_code,
      });
      // R7 — success already redirected via useLoginMfa's onSuccess.
    } catch (error) {
      // R8 — stays on the MFA step; the mfa_pending_token isn't cleared
      // here, so a retry (wrong-code case) reuses it as-is.
      setRequestError(
        error instanceof ApiError && error.detail ? error.detail : GENERIC_ERROR_MESSAGE
      );
    }
  }

  function backToPasswordStep() {
    setStep("password");
    setMfaPendingToken(null);
    setRequestError(null);
    mfaForm.reset();
  }

  return (
    <div className="flex flex-col gap-6">
      {requestError && (
        <div ref={bannerRef} tabIndex={-1} className="outline-none">
          <Banner variant="error">{requestError}</Banner>
        </div>
      )}

      {step === "password" ? (
        <form
          className="flex flex-col gap-6"
          onSubmit={passwordForm.handleSubmit(onSubmitPassword)}
          noValidate
        >
          <div className="flex flex-col gap-1.5">
            <Label htmlFor="login-email">Email</Label>
            <Input
              id="login-email"
              type="email"
              autoComplete="email"
              error={passwordForm.formState.errors.email?.message}
              {...passwordForm.register("email")}
            />
          </div>

          <div className="flex flex-col gap-1.5">
            <div className="flex items-baseline justify-between gap-3">
              <Label htmlFor="login-password">Password</Label>
              <Link
                href="/forgot-password"
                className="text-sm font-semibold text-primary-700"
              >
                Lupa password?
              </Link>
            </div>
            <div className="flex items-center gap-2">
              <Input
                id="login-password"
                type={showPassword ? "text" : "password"}
                autoComplete="current-password"
                error={passwordForm.formState.errors.password?.message}
                className="flex-1"
                {...passwordForm.register("password")}
              />
              <Button
                type="button"
                variant="ghost"
                size="sm"
                aria-label={showPassword ? "Sembunyikan password" : "Tampilkan password"}
                onClick={() => setShowPassword((show) => !show)}
              >
                {showPassword ? (
                  <EyeOff aria-hidden="true" className="size-4" />
                ) : (
                  <Eye aria-hidden="true" className="size-4" />
                )}
              </Button>
            </div>
          </div>

          <Button
            type="submit"
            loading={passwordForm.formState.isSubmitting || loginMutation.isPending}
            className="w-full"
          >
            Masuk
          </Button>

          <div className="flex items-center gap-3" aria-hidden="true">
            <span className="h-px flex-1 bg-neutral-200" />
            <span className="text-xs text-neutral-500">atau</span>
            <span className="h-px flex-1 bg-neutral-200" />
          </div>

          <GoogleAuthButton intent="login" label="Masuk dengan Google" />

          <span className="text-center text-sm text-neutral-700">
            Belum punya akun?{" "}
            {onSwitchToRegister ? (
              <button
                type="button"
                onClick={onSwitchToRegister}
                className="font-semibold text-primary-700"
              >
                Daftar
              </button>
            ) : (
              <Link href="/register" className="font-semibold text-primary-700">
                Daftar
              </Link>
            )}
          </span>
        </form>
      ) : (
        <form
          className="flex flex-col gap-6"
          onSubmit={mfaForm.handleSubmit(onSubmitMfa)}
          noValidate
        >
          <div className="flex flex-col gap-2">
            <h2 className="text-xl font-semibold text-neutral-900">
              Verifikasi dua langkah
            </h2>
            <p className="text-sm text-neutral-500">
              Masukkan kode dari aplikasi autentikator kamu.
            </p>
          </div>

          {useBackupCode ? (
            <div className="flex flex-col gap-1.5">
              <Label htmlFor="login-mfa-backup">Kode cadangan</Label>
              <Input
                id="login-mfa-backup"
                autoComplete="one-time-code"
                error={mfaForm.formState.errors.backup_code?.message}
                {...mfaForm.register("backup_code")}
              />
            </div>
          ) : (
            <div className="flex flex-col gap-1.5">
              <Label htmlFor="login-mfa-totp">Kode OTP</Label>
              <Input
                id="login-mfa-totp"
                inputMode="numeric"
                autoComplete="one-time-code"
                error={mfaForm.formState.errors.totp_code?.message}
                {...mfaForm.register("totp_code")}
              />
            </div>
          )}

          <Button
            type="submit"
            loading={mfaForm.formState.isSubmitting || loginMfaMutation.isPending}
            className="w-full"
          >
            Verifikasi
          </Button>

          <button
            type="button"
            onClick={() => {
              setUseBackupCode((use) => !use);
              mfaForm.clearErrors();
            }}
            className="text-center text-sm font-semibold text-primary-700"
          >
            {useBackupCode ? "Gunakan kode OTP" : "Gunakan kode cadangan"}
          </button>

          <button
            type="button"
            onClick={backToPasswordStep}
            className="text-center text-sm text-neutral-500"
          >
            Kembali ke halaman login
          </button>
        </form>
      )}
    </div>
  );
}
