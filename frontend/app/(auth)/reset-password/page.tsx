import { Suspense } from "react";
import { ResetPasswordForm } from "@/components/features/account/reset-password-form";

// `ResetPasswordForm` reads `token` via `useSearchParams()`, which requires
// a <Suspense> boundary in the App Router — same requirement `/login`
// already established for `GoogleCallbackError`. `fallback={null}` matches
// that same precedent: this is a synchronous client-side read, not a real
// network-gated loading state.
export default function ResetPasswordPage() {
  return (
    <div className="flex flex-col gap-6">
      <div className="flex flex-col gap-2">
        <h1 className="text-2xl font-bold text-neutral-900">Reset Password</h1>
        <p className="text-sm text-neutral-500">
          Masukkan password baru untuk akun kamu.
        </p>
      </div>
      <Suspense fallback={null}>
        <ResetPasswordForm />
      </Suspense>
    </div>
  );
}
