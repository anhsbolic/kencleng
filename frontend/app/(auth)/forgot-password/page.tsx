import { ForgotPasswordForm } from "@/components/features/account/forgot-password-form";

export default function ForgotPasswordPage() {
  return (
    <div className="flex flex-col gap-6">
      <div className="flex flex-col gap-2">
        <h1 className="text-2xl font-bold text-neutral-900">Lupa Password</h1>
        <p className="text-sm text-neutral-500">
          Masukkan email kamu, kami akan kirim link untuk reset password.
        </p>
      </div>
      <ForgotPasswordForm />
    </div>
  );
}
