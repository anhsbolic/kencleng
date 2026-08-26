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
import { useRegister } from "@/lib/hooks/use-register";
import { GoogleAuthButton } from "./google-auth-button";
import { ResendVerificationControl } from "./resend-verification-control";
import { registerSchema, type RegisterFormValues } from "./register-schema";

const GENERIC_ERROR_MESSAGE = "Terjadi kesalahan. Silakan coba lagi."; // TBD — Open Item #5, placeholder pending product copy
const DEFAULT_SUCCESS_MESSAGE =
  "Kalau email belum terdaftar, cek inbox untuk verifikasi. Kalau sudah, cek inbox untuk instruksi lebih lanjut.";

/**
 * `/register`'s form (techplan account/01-register-email-verification,
 * Task 2). Every backend branch behind `POST /auth/register` collapses
 * to the same `202` — this component must never differentiate them in
 * copy or UI (rule R4/R18; `restapi/anti-enumeration.md`), so the
 * success view below is unconditional on anything but "was this a
 * 202."
 */
export function RegisterForm() {
  const registerMutation = useRegister();
  const [submittedEmail, setSubmittedEmail] = useState<string | null>(null);
  const [successMessage, setSuccessMessage] = useState<string | undefined>();
  const [requestError, setRequestError] = useState<string | null>(null);
  const successHeadingRef = useRef<HTMLHeadingElement>(null);

  const {
    register: registerField,
    handleSubmit,
    setError,
    formState: { errors, isSubmitting },
  } = useForm<RegisterFormValues>({
    resolver: zodResolver(registerSchema),
  });

  // R17 — focus moves into the success region once the form is
  // replaced, never left on a now-removed element.
  useEffect(() => {
    if (submittedEmail) {
      successHeadingRef.current?.focus();
    }
  }, [submittedEmail]);

  async function onSubmit(values: RegisterFormValues) {
    setRequestError(null);

    let result;
    try {
      result = await registerMutation.mutateAsync(values);
    } catch (error) {
      // R6/R10 — a documented 429 shows the backend's own detail
      // verbatim; anything else (network failure, unexpected 5xx)
      // falls back to one frontend-owned generic message, raw body
      // never rendered.
      setRequestError(
        error instanceof ApiError && error.status === 429 && error.detail
          ? error.detail
          : GENERIC_ERROR_MESSAGE
      );
      return;
    }

    if (result.ok) {
      // R4 — fixed success view, verbatim backend message, never
      // conditioned on which internal branch actually fired.
      setSuccessMessage(result.message);
      setSubmittedEmail(values.email);
      return;
    }

    // R5 — result.ok === false, kind: "validation": map each field's
    // message verbatim, never re-authored client-side.
    for (const { field, message } of result.errors) {
      if (field === "name" || field === "email" || field === "password") {
        setError(field, { message });
      }
    }
  }

  if (submittedEmail) {
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
        <div className="flex flex-col gap-2 text-sm text-neutral-500">
          <span>Belum menerima email?</span>
          <ResendVerificationControl defaultEmail={submittedEmail} />
        </div>
      </div>
    );
  }

  return (
    <form className="flex flex-col gap-6" onSubmit={handleSubmit(onSubmit)} noValidate>
      {requestError && <Banner variant="error">{requestError}</Banner>}

      <div className="flex flex-col gap-1.5">
        <Label htmlFor="register-name">Nama</Label>
        <Input
          id="register-name"
          autoComplete="name"
          error={errors.name?.message}
          {...registerField("name")}
        />
      </div>

      <div className="flex flex-col gap-1.5">
        <Label htmlFor="register-email">Email</Label>
        <Input
          id="register-email"
          type="email"
          autoComplete="email"
          error={errors.email?.message}
          {...registerField("email")}
        />
      </div>

      <div className="flex flex-col gap-1.5">
        <Label htmlFor="register-password">Password</Label>
        <Input
          id="register-password"
          type="password"
          autoComplete="new-password"
          placeholder="Minimal 8 karakter"
          error={errors.password?.message}
          {...registerField("password")}
        />
      </div>

      <Button
        type="submit"
        loading={isSubmitting || registerMutation.isPending}
        className="w-full"
      >
        Daftar
      </Button>

      <div className="flex items-center gap-3" aria-hidden="true">
        <span className="h-px flex-1 bg-neutral-200" />
        <span className="text-xs text-neutral-500">atau</span>
        <span className="h-px flex-1 bg-neutral-200" />
      </div>

      {/* R7 — a real navigation to the shared redirect-initiation
          endpoint, not an apiFetch/XHR call: the endpoint issues a
          302 that only a real navigation follows correctly. `intent`
          is always "login" (techplan account/02-google-oauth-login-
          register, D1/R1) — the backend's `login` intent already
          creates a new User when no identity/email match exists;
          there is no separate `register` intent. */}
      <GoogleAuthButton intent="login" label="Daftar dengan Google" />

      <span className="text-center text-sm text-neutral-700">
        Sudah punya akun?{" "}
        <Link href="/login" className="font-semibold text-primary-700">
          Masuk
        </Link>
      </span>
    </form>
  );
}
