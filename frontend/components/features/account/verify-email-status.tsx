"use client";

import { useQueryClient } from "@tanstack/react-query";
import Link from "next/link";
import { useSearchParams } from "next/navigation";
import { useEffect, useRef } from "react";
import { Banner } from "@/components/ui/banner";
import { Spinner } from "@/components/ui/spinner";
import { ApiError } from "@/lib/api/client";
import { accountKeys } from "@/lib/hooks/use-account-me";
import { useVerifyEmail } from "@/lib/hooks/use-verify-email";
import { useAuthStore } from "@/lib/stores/auth-store";
import { ResendVerificationControl } from "./resend-verification-control";

// TBD — Open Item #1 on the originating techplan: no worked example
// exists for this exact copy anywhere in the OpenAPI schema (unlike
// every other outcome this page handles). Placeholder pending product
// sign-off.
const INVALID_LINK_MESSAGE = "Link verifikasi tidak valid atau sudah digunakan.";
const GENERIC_ERROR_MESSAGE = "Terjadi kesalahan. Silakan coba lagi."; // TBD — Open Item #5

type Outcome =
  | { kind: "loading" }
  | { kind: "verified"; message?: string }
  | { kind: "expired"; detail?: string }
  | { kind: "invalid" }
  | { kind: "rate-limited"; detail?: string }
  | { kind: "error" };

/**
 * `/verify-email`'s outcome view (techplan
 * account/01-register-email-verification, Task 3, Decision D1 — a
 * top-level route, deliberately not nested in `AuthShellClient`).
 * Reads the `token` query param, fires `verifyEmail` exactly once per
 * token (rule R12 — a second automatic call would itself 404 even
 * against a link that was genuinely still valid), and renders one of
 * four outcomes matching the originating techplan's decision-flow
 * diagram.
 */
export function VerifyEmailStatus() {
  const searchParams = useSearchParams();
  const token = searchParams.get("token");
  const verifyEmail = useVerifyEmail();
  const firedRef = useRef(false);
  const resultHeadingRef = useRef<HTMLHeadingElement>(null);
  const queryClient = useQueryClient();
  // techplan account/05-account-linking, R19/D6 — an already-
  // authenticated caller mid-linking-flow (Branch 1's step 2) needs a
  // different terminal CTA than a logged-out registrant; both share
  // this same component/route, per the spec's "reuse unchanged"
  // framing at the wire-contract level.
  const accessToken = useAuthStore((state) => state.accessToken);

  useEffect(() => {
    if (!token || firedRef.current) return;
    firedRef.current = true;
    verifyEmail.mutate({ token });
    // eslint-disable-next-line react-hooks/exhaustive-deps -- fire exactly once per mount/token, not on every verifyEmail identity change (R12)
  }, [token]);

  // R19 — invalidate regardless of auth state: harmless no-op for the
  // logged-out registration caller (nothing cached yet), required for
  // an authenticated caller so /dashboard/security reflects the
  // just-verified identity without a stale cache.
  useEffect(() => {
    if (verifyEmail.isSuccess) {
      queryClient.invalidateQueries({ queryKey: accountKeys.me() });
    }
  }, [verifyEmail.isSuccess, queryClient]);

  const outcome: Outcome = !token
    ? { kind: "invalid" } // R11 — a missing token is treated identically to 404, no separate message
    : verifyEmail.isSuccess
      ? { kind: "verified", message: verifyEmail.data.message }
      : verifyEmail.isError
        ? errorToOutcome(verifyEmail.error)
        : { kind: "loading" };

  // R16 — focus moves into the result region once the loading state
  // resolves to any outcome, never left on the removed skeleton.
  useEffect(() => {
    if (outcome.kind !== "loading") {
      resultHeadingRef.current?.focus();
    }
  }, [outcome.kind]);

  return (
    <div className="mx-auto flex min-h-full w-full max-w-md flex-col justify-center gap-6 p-6">
      <h1
        ref={resultHeadingRef}
        tabIndex={-1}
        className="text-xl font-semibold text-neutral-900 outline-none"
      >
        Verifikasi Email
      </h1>

      {outcome.kind === "loading" && (
        <div className="flex items-center gap-2 text-neutral-500" role="status">
          <Spinner className="size-4" />
          <span>Memverifikasi tautan...</span>
        </div>
      )}

      {outcome.kind === "verified" && (
        <>
          <Banner variant="success">
            {outcome.message ?? "Email berhasil diverifikasi."}
          </Banner>
          <Link
            href={accessToken ? "/dashboard/security" : "/login"}
            className="font-semibold text-primary-700"
          >
            {accessToken ? "Kembali ke Keamanan" : "Masuk sekarang"}
          </Link>
        </>
      )}

      {outcome.kind === "expired" && (
        <>
          <Banner variant="error">
            {outcome.detail ??
              "Link verifikasi sudah kedaluwarsa. Silakan minta kirim ulang."}
          </Banner>
          <ResendVerificationControl />
        </>
      )}

      {outcome.kind === "invalid" && (
        <Banner variant="error">{INVALID_LINK_MESSAGE}</Banner>
      )}

      {outcome.kind === "rate-limited" && (
        <Banner variant="error">{outcome.detail ?? GENERIC_ERROR_MESSAGE}</Banner>
      )}

      {outcome.kind === "error" && <Banner variant="error">{GENERIC_ERROR_MESSAGE}</Banner>}
    </div>
  );
}

function errorToOutcome(error: unknown): Outcome {
  if (error instanceof ApiError) {
    if (error.status === 410) return { kind: "expired", detail: error.detail };
    if (error.status === 404) return { kind: "invalid" };
    if (error.status === 429) return { kind: "rate-limited", detail: error.detail };
  }
  // R6 — anything outside the documented 404/410/429 set (network
  // failure, unexpected 5xx) falls back to one frontend-owned generic
  // message; the raw error is never inspected/rendered.
  return { kind: "error" };
}
