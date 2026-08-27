import { LoginMethodsSection } from "@/components/features/account/login-methods-section";
import { MfaSection } from "@/components/features/account/mfa-section";

// Real content for Account Task #5 (account-linking) + Account Task #6
// (MFA, techplan account/06-mfa-totp) — replaces the Phase 0 placeholder.
// Independent section components, not a monolithic form (techplan
// account/05-account-linking D1).
export default function SecurityPage() {
  return (
    <div className="flex flex-col gap-6">
      <h1 className="text-2xl font-bold text-neutral-900">Keamanan</h1>
      <LoginMethodsSection />
      <MfaSection />
    </div>
  );
}
