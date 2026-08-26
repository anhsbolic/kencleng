"use client";

import { useAuthModalStore } from "@/lib/stores/auth-modal-store";

/**
 * Desktop "Masuk"/"Daftar" nav buttons — opens the landing-page auth
 * modal instead of navigating to `/login`/`/register`. Extracted as
 * its own small Client Component so `app/(public)/layout.tsx` (the
 * desktop nav's Server Component parent) doesn't need a `'use client'`
 * boundary of its own — `'use client'` at the smallest leaf
 * (`server-client-component-boundary.md`), matching how the mobile
 * hamburger/drawer is already split out as `PublicShellClient`.
 * Styling matches the previous `<Link>`-based markup exactly.
 */
export function AuthModalTriggers() {
  const openLogin = useAuthModalStore((state) => state.openLogin);
  const openRegister = useAuthModalStore((state) => state.openRegister);

  return (
    <>
      <button
        type="button"
        onClick={openLogin}
        className="inline-flex h-9 items-center justify-center rounded-md border border-neutral-200 px-3 text-sm font-semibold text-neutral-700 transition-colors hover:bg-neutral-100"
      >
        Masuk
      </button>
      <button
        type="button"
        onClick={openRegister}
        className="inline-flex h-9 items-center justify-center rounded-md bg-primary-600 px-3 text-sm font-semibold text-white shadow-sm transition-colors hover:bg-primary-700"
      >
        Daftar
      </button>
    </>
  );
}
