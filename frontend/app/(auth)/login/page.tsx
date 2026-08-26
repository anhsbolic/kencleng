import { Suspense } from "react";
import { GoogleAuthButton } from "@/components/features/account/google-auth-button";
import { GoogleCallbackError } from "@/components/features/account/google-callback-error";

// Real content: Google entry point + `?error={code}` handling (techplan
// account/02-google-oauth-login-register). The email/password
// credential form is a different task's scope (backend task #3,
// `03-login-session-management.md`) — deliberately not built here, per
// D2. Whoever picks up task #3 extends this page (adds the form above/
// alongside the Google button + divider, matching RegisterForm's own
// composition), not replaces it.
export default function LoginPage() {
  return (
    <div className="flex flex-col gap-6">
      <div className="flex flex-col gap-2">
        <h1 className="text-2xl font-bold text-neutral-900">Masuk</h1>
        <p className="text-sm text-neutral-500">
          Masuk ke akun Kencleng kamu.
        </p>
      </div>

      {/* `useSearchParams()` requires a <Suspense> boundary in the App
          Router — same requirement `VerifyEmailStatus`/
          `app/verify-email/page.tsx` already established. Scoped only
          around GoogleCallbackError, not the whole page, since the
          heading/note/button below don't depend on search params. */}
      <Suspense fallback={null}>
        <GoogleCallbackError />
      </Suspense>

      <p className="rounded-md bg-neutral-100 px-4 py-3 text-sm text-neutral-500">
        Login dengan email/password akan segera hadir. Untuk saat ini,
        gunakan Google untuk masuk.
      </p>

      <GoogleAuthButton intent="login" label="Masuk dengan Google" />
    </div>
  );
}
