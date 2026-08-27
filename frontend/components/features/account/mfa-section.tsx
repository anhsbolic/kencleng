"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { useAccountMe } from "@/lib/hooks/use-account-me";
import { MfaDisableForm } from "./mfa-disable-form";
import { MfaEnrollFlow } from "./mfa-enroll-flow";

// New frontend-owned copy (techplan D4) — the acknowledgment gate's
// label; not yet finalized/// TBD, same placeholder-pending-copy
// convention used everywhere else in this codebase.
const ACKNOWLEDGE_CODES_LABEL = "Saya sudah menyimpan kode ini"; // TBD

/**
 * `/dashboard/security`'s MFA section (techplan account/06-mfa-totp,
 * task-04) — the sibling this page's own comment names Account Task #6
 * as adding, alongside `LoginMethodsSection` (Task 05).
 *
 * Branches on `useAccountMe()`'s `mfa_enabled`, **except** while
 * `justEnrolledCodes` is set (R12): the moment `MfaEnrollFlow`'s
 * `onEnrolled` callback fires, `useMfaEnrollConfirm`'s own `onSuccess`
 * (task-01) has already invalidated `account.me`, which will refetch
 * `mfa_enabled: true` — if this component branched purely on that
 * field, the one-time backup-codes view would disappear within
 * milliseconds of appearing. Holding the codes in local state here,
 * checked *before* the `mfa_enabled` branch, is what makes the codes
 * view survive that refetch until the user explicitly acknowledges
 * having saved them (R13).
 */
export function MfaSection() {
  const { data: user } = useAccountMe();
  const [justEnrolledCodes, setJustEnrolledCodes] = useState<string[] | null>(null);

  // R1 — skeleton shaped like the real section, matching
  // `LoginMethodsSection`'s own loading-branch convention; no `isError`
  // branch, same reasoning as that component's doc comment.
  if (!user) {
    return (
      <div className="flex flex-col gap-4 rounded-lg border border-neutral-200 bg-white p-6">
        <div className="h-6 w-40 animate-pulse rounded bg-neutral-100" />
        <div className="h-10 w-full animate-pulse rounded bg-neutral-100" />
      </div>
    );
  }

  return (
    <section className="flex flex-col gap-6 rounded-lg border border-neutral-200 bg-white p-6">
      <h2 className="text-xl font-semibold text-neutral-900">Autentikasi Dua Faktor (MFA)</h2>

      {justEnrolledCodes ? (
        <div className="flex flex-col gap-4">
          <p className="text-sm text-neutral-700">
            Simpan 10 kode cadangan ini di tempat aman — kode ini hanya ditampilkan sekali dan
            tidak bisa dilihat lagi setelah ini.
          </p>
          {/* R22 — plain, selectable/copyable text, not an image, so
              screen readers and copy-paste both work. */}
          <ul className="grid grid-cols-2 gap-2 rounded-md bg-neutral-100 p-4 font-mono text-sm text-neutral-900">
            {justEnrolledCodes.map((code) => (
              <li key={code} className="select-all">
                {code}
              </li>
            ))}
          </ul>
          <Button
            type="button"
            onClick={() => setJustEnrolledCodes(null)}
            className="w-fit"
          >
            {ACKNOWLEDGE_CODES_LABEL}
          </Button>
        </div>
      ) : user.mfa_enabled ? (
        <MfaDisableForm
          hasEmailPassword={user.auth_providers?.includes("email_password") ?? false}
        />
      ) : (
        <MfaEnrollFlow onEnrolled={setJustEnrolledCodes} />
      )}
    </section>
  );
}
