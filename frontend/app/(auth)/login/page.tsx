import { Suspense } from "react";
import { GoogleCallbackError } from "@/components/features/account/google-callback-error";
import { LoginForm } from "@/components/features/account/login-form";

// Real content: Google entry point + `?error={code}` handling (task #2)
// plus the email/password credential form + MFA challenge step (task #3,
// `03-login-session-management.md`, techplan task-03) — `LoginForm`
// owns the Google button internally (password step only, R9), matching
// `RegisterForm`'s own composition.
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
          heading/LoginForm below don't depend on search params. */}
      <Suspense fallback={null}>
        <GoogleCallbackError />
      </Suspense>

      <LoginForm />
    </div>
  );
}
