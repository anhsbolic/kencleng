import { RegisterForm } from "@/components/features/account/register-form";

export default function RegisterPage() {
  return (
    <div className="flex flex-col gap-6">
      <div className="flex flex-col gap-2">
        <h1 className="text-2xl font-bold text-neutral-900">Daftar</h1>
        <p className="text-sm text-neutral-500">Buat akun Kencleng baru.</p>
      </div>
      <RegisterForm />
    </div>
  );
}
