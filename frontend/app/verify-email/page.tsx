import { Suspense } from "react";
import { VerifyEmailStatus } from "@/components/features/account/verify-email-status";

// Top-level route (techplan account/01-register-email-verification,
// Decision D1) — deliberately NOT nested under `app/(auth)/`: this
// link is opened from an email client, often with no prior in-app
// navigation, so `AuthShellClient`'s desktop modal (which blurs a
// rendering of `/` behind it) would misleadingly imply the visitor
// was mid-browse. Uses the Status/Tracking pattern's minimal shell
// instead, same as `/donation/[id]/status`.
//
// `VerifyEmailStatus` reads `token` via `useSearchParams()`, which
// requires a `<Suspense>` boundary in the Next.js App Router or the
// build errors — resolves Open Item #6 from the originating techplan.
export default function VerifyEmailPage() {
  return (
    <Suspense fallback={<VerifyEmailLoadingFallback />}>
      <VerifyEmailStatus />
    </Suspense>
  );
}

function VerifyEmailLoadingFallback() {
  return (
    <div className="mx-auto flex min-h-full w-full max-w-md flex-col justify-center gap-6 p-6">
      <p className="text-neutral-500" role="status">
        Memuat...
      </p>
    </div>
  );
}
