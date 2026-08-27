import { Banner } from "@/components/ui/banner";
import { GoogleAuthButton } from "./google-auth-button";
import { UnlinkGoogleForm } from "./unlink-google-form";

// Copied verbatim from backend/internal/transport/http/account_security.go
// (lines 184, 189) — both are already correct, final Indonesian text
// (techplan account/05-account-linking D4). Keep these in sync with
// those exact source lines if the backend's copy ever changes.
const ONLY_IDENTITY_MESSAGE =
  "Google adalah satu-satunya metode login Anda. Atur email dan password dulu sebelum melepas tautan.";
const UNVERIFIED_REMAINING_MESSAGE =
  "Kamu sudah atur email dan password, tapi belum diverifikasi. Verifikasi email kamu dulu sebelum bisa melepas tautan Google.";

export interface GoogleIdentityControlProps {
  hasGoogle: boolean;
  canUnlink: boolean;
  blockedReason: "only-identity" | "unverified" | null;
}

/**
 * `/dashboard/security`'s Google identity action (techplan account/05-
 * account-linking, R13-R16): link trigger, proactive-blocked notice, or
 * the unlink form — all three computed by the parent
 * (`LoginMethodsSection`) from already-fetched `useAccountMe()` data,
 * avoiding a guaranteed-`409` round trip for the common blocked cases.
 */
export function GoogleIdentityControl({
  hasGoogle,
  canUnlink,
  blockedReason,
}: GoogleIdentityControlProps) {
  if (!hasGoogle) {
    return <GoogleAuthButton intent="link" label="Hubungkan ke Google" />;
  }

  if (!canUnlink) {
    return (
      <Banner variant="info">
        {blockedReason === "only-identity" ? ONLY_IDENTITY_MESSAGE : UNVERIFIED_REMAINING_MESSAGE}
      </Banner>
    );
  }

  return <UnlinkGoogleForm />;
}
