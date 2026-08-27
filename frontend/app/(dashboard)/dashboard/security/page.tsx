import { LoginMethodsSection } from "@/components/features/account/login-methods-section";

// Real content for Account Task #5 (account-linking) — replaces the
// Phase 0 placeholder. Account Task #6 (MFA) adds its own <MfaSection />
// here as a sibling when that task starts (techplan account/05-account-
// linking D1 — independent section components, not a monolithic form).
export default function SecurityPage() {
  return (
    <div className="flex flex-col gap-6">
      <h1 className="text-2xl font-bold text-neutral-900">Keamanan</h1>
      <LoginMethodsSection />
      {/* Account Task #6 (MFA) adds <MfaSection /> here */}
    </div>
  );
}
