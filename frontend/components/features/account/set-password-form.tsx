"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { Banner } from "@/components/ui/banner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { PasswordInput } from "@/components/shared/password-input";
import { ApiError } from "@/lib/api/client";
import { useSetPassword } from "@/lib/hooks/use-set-password";
import {
  addPasswordSchema,
  changePasswordSchema,
  type AddPasswordFormValues,
  type ChangePasswordFormValues,
} from "./set-password-schema";

const GENERIC_ERROR_MESSAGE = "Terjadi kesalahan. Silakan coba lagi."; // TBD — same placeholder-pending-copy treatment as every other generic string in this codebase
// Reused verbatim from reset-password-form.tsx's existing constant —
// both branches share the same backend validatePassword policy
// (account_security.go), same field ("password"), same copy.
const WEAK_PASSWORD_MESSAGE =
  "Password tidak memenuhi syarat. Gunakan minimal 8 karakter dan hindari password yang umum digunakan atau pernah bocor.";
// Backend's own confirmed 202 response text (account_security.go) —
// used as the default fallback if `message` is ever absent, same
// convention as ForgotPasswordForm's DEFAULT_SUCCESS_MESSAGE.
const DEFAULT_ADDED_MESSAGE = "Kalau email tersedia, cek inbox untuk verifikasi.";

export interface SetPasswordFormProps {
  mode: "add" | "change";
}

/**
 * `/dashboard/security`'s set-password action (techplan account/05-
 * account-linking, R4-R12). `mode` is a prop from the parent
 * (`LoginMethodsSection`), not internal state — branch selection is
 * server-side and verified-agnostic (D3), so the parent already knows
 * which fields to show before this component ever renders.
 */
export function SetPasswordForm({ mode }: SetPasswordFormProps) {
  const setPasswordMutation = useSetPassword();
  const [addedMessage, setAddedMessage] = useState<string | null>(null);
  const [requestError, setRequestError] = useState<string | null>(null);
  const bannerRef = useRef<HTMLDivElement>(null);

  const addForm = useForm<AddPasswordFormValues>({ resolver: zodResolver(addPasswordSchema) });
  const changeForm = useForm<ChangePasswordFormValues>({
    resolver: zodResolver(changePasswordSchema),
  });

  // Focus moves into whichever banner is currently shown, matching the
  // convention already established by LoginForm/RegisterForm elsewhere
  // in this codebase.
  useEffect(() => {
    if (addedMessage || requestError) {
      bannerRef.current?.focus();
    }
  }, [addedMessage, requestError]);

  async function onSubmitAdd(values: AddPasswordFormValues) {
    setRequestError(null);

    let result;
    try {
      result = await setPasswordMutation.mutateAsync({
        email: values.email,
        password: values.password,
      });
    } catch {
      // Branch 1 never throws a documented 401 — only network/5xx
      // reach here.
      setRequestError(GENERIC_ERROR_MESSAGE);
      return;
    }

    if (!result.ok) {
      for (const { field } of result.errors) {
        if (field === "password") {
          addForm.setError("password", { message: WEAK_PASSWORD_MESSAGE });
        }
      }
      return;
    }

    // R6 — result.branch is always "added" here (Branch 1 never
    // returns 200); fixed success view, never conditioned on the
    // internal case (anti-enumeration).
    setAddedMessage(result.message ?? DEFAULT_ADDED_MESSAGE);
    addForm.reset();
  }

  async function onSubmitChange(values: ChangePasswordFormValues) {
    setRequestError(null);

    let result;
    try {
      result = await setPasswordMutation.mutateAsync({
        current_password: values.current_password,
        password: values.password,
      });
    } catch (error) {
      // R11 — 401 shows the backend's own detail verbatim (confirmed
      // correct Indonesian, same shared string as LoginForm's own
      // justified exception); anything else falls back to generic.
      setRequestError(
        error instanceof ApiError && error.status === 401 && error.detail
          ? error.detail
          : GENERIC_ERROR_MESSAGE
      );
      return;
    }

    if (!result.ok) {
      for (const { field } of result.errors) {
        if (field === "password") {
          changeForm.setError("password", { message: WEAK_PASSWORD_MESSAGE });
        }
      }
      return;
    }

    // R10 — result.branch is always "changed" here. useSetPassword's
    // onSuccess already cleared the session; SessionGuardProvider
    // redirects to /login. Nothing further to render here (matches
    // useLogout's own no-farewell-message convention).
  }

  if (mode === "add" && addedMessage) {
    return (
      <div ref={bannerRef} tabIndex={-1} className="outline-none">
        <Banner variant="success">{addedMessage}</Banner>
      </div>
    );
  }

  if (mode === "add") {
    return (
      <form className="flex flex-col gap-4" onSubmit={addForm.handleSubmit(onSubmitAdd)} noValidate>
        <h3 className="text-lg font-semibold text-neutral-900">Atur Password</h3>

        {requestError && (
          <div ref={bannerRef} tabIndex={-1} className="outline-none">
            <Banner variant="error">{requestError}</Banner>
          </div>
        )}

        <div className="flex flex-col gap-1.5">
          <Label htmlFor="set-password-email">Email</Label>
          <Input
            id="set-password-email"
            type="email"
            autoComplete="email"
            error={addForm.formState.errors.email?.message}
            {...addForm.register("email")}
          />
        </div>

        <div className="flex flex-col gap-1.5">
          <Label htmlFor="set-password-new">Password baru</Label>
          <PasswordInput
            id="set-password-new"
            autoComplete="new-password"
            placeholder="Minimal 8 karakter"
            error={addForm.formState.errors.password?.message}
            {...addForm.register("password")}
          />
        </div>

        <Button
          type="submit"
          loading={addForm.formState.isSubmitting || setPasswordMutation.isPending}
          className="w-fit"
        >
          Atur Password
        </Button>
      </form>
    );
  }

  return (
    <form className="flex flex-col gap-4" onSubmit={changeForm.handleSubmit(onSubmitChange)} noValidate>
      <h3 className="text-lg font-semibold text-neutral-900">Ganti Password</h3>

      {requestError && (
        <div ref={bannerRef} tabIndex={-1} className="outline-none">
          <Banner variant="error">{requestError}</Banner>
        </div>
      )}

      <div className="flex flex-col gap-1.5">
        <Label htmlFor="set-password-current">Password saat ini</Label>
        <PasswordInput
          id="set-password-current"
          autoComplete="current-password"
          error={changeForm.formState.errors.current_password?.message}
          {...changeForm.register("current_password")}
        />
      </div>

      <div className="flex flex-col gap-1.5">
        <Label htmlFor="set-password-change-new">Password baru</Label>
        <PasswordInput
          id="set-password-change-new"
          autoComplete="new-password"
          placeholder="Minimal 8 karakter"
          error={changeForm.formState.errors.password?.message}
          {...changeForm.register("password")}
        />
      </div>

      <Button
        type="submit"
        loading={changeForm.formState.isSubmitting || setPasswordMutation.isPending}
        className="w-fit"
      >
        Ganti Password
      </Button>
    </form>
  );
}
