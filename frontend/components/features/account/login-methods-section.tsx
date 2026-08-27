"use client";

import { Banner } from "@/components/ui/banner";
import { useAccountMe } from "@/lib/hooks/use-account-me";
import { GoogleIdentityControl } from "./google-identity-control";
import { SetPasswordForm } from "./set-password-form";

/**
 * `/dashboard/security`'s login-methods section (techplan account/05-
 * account-linking, R1-R3) — the sole content this task adds to the
 * page. Account Task #6 (MFA) adds its own sibling section later; see
 * `app/(dashboard)/dashboard/security/page.tsx`'s comment (D1).
 */
export function LoginMethodsSection() {
  const { data: user } = useAccountMe();

  // R3 — skeleton shaped like the real section, not a bare spinner
  // (loading-empty-error-state-conventions.md). No `isError` branch:
  // `useAccountMe()` already has a safe-default consumer elsewhere in
  // this Shell (`useHasRole`) and this section simply waits for data,
  // matching `LogoutButton`'s existing `if (!user) return null`-shaped
  // gating rather than treating "not yet loaded" as a failure.
  if (!user) {
    return (
      <div className="flex flex-col gap-4 rounded-lg border border-neutral-200 bg-white p-6">
        <div className="h-6 w-40 animate-pulse rounded bg-neutral-100" />
        <div className="h-10 w-full animate-pulse rounded bg-neutral-100" />
        <div className="h-10 w-full animate-pulse rounded bg-neutral-100" />
      </div>
    );
  }

  const hasEmailPassword = user.auth_providers?.includes("email_password") ?? false;
  const hasGoogle = user.auth_providers?.includes("google") ?? false;
  const verified = user.email_verified ?? false;

  // R1 — mode is gated purely on identity presence, matching the
  // backend's own verified-agnostic branch check exactly (D3). R13-R16
  // — canUnlink/blockedReason mirror INV-account-02/INV-account-12's
  // exact guard conditions client-side, to avoid a guaranteed-409
  // round trip for the common blocked cases.
  const mode: "add" | "change" = hasEmailPassword ? "change" : "add";
  const canUnlink = hasGoogle && hasEmailPassword && verified;
  const blockedReason: "only-identity" | "unverified" | null = !hasGoogle
    ? null
    : !hasEmailPassword
      ? "only-identity"
      : !verified
        ? "unverified"
        : null;

  return (
    <section className="flex flex-col gap-6 rounded-lg border border-neutral-200 bg-white p-6">
      <h2 className="text-xl font-semibold text-neutral-900">Metode Masuk</h2>

      {/* R2 — informational, shown above the form, which stays fully
          interactive regardless (Branch 2 works against an unverified
          identity too, confirmed D3). */}
      {hasEmailPassword && !verified && (
        <Banner variant="info">
          Menunggu verifikasi email kamu — cek inbox untuk menyelesaikan.
        </Banner>
      )}

      {/* `key={mode}` forces a clean remount if the account's identity
          shape changes underneath an already-mounted form (e.g. after
          verifying mid-session) — avoids carrying stale local
          success/error state across a Branch 1 → Branch 2 transition. */}
      <SetPasswordForm key={mode} mode={mode} />

      <GoogleIdentityControl hasGoogle={hasGoogle} canUnlink={canUnlink} blockedReason={blockedReason} />
    </section>
  );
}
