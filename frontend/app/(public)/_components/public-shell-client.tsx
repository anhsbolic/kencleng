"use client";

import { Menu, X } from "lucide-react";
import Link from "next/link";
import { useRef, useState } from "react";
import { useFocusTrap } from "@/lib/hooks/use-focus-trap";
import { useAuthModalStore } from "@/lib/stores/auth-modal-store";
import { publicNavItems } from "./nav-items";

const DRAWER_ID = "public-mobile-nav";

// Masuk/Daftar open the landing-page auth modal (login/register-as-a-
// modal follow-up to techplan account/03-login-session-management)
// instead of navigating — styled to match the previous Link-based
// markup exactly.
const authButtonBase =
  "inline-flex h-11 w-full items-center justify-center rounded-md px-5 text-base font-semibold transition-colors";

/**
 * Public Shell's mobile hamburger + drawer (`scaffold-public-shell.md`
 * Step 3). Mirrors `(dashboard)/_components/dashboard-shell-client.tsx`'s
 * shape — hamburger with `aria-expanded`/`aria-controls`, a
 * `role="dialog"` drawer, `useFocusTrap` for open/close focus
 * management — but isn't a copy-import of it: that component is
 * private to the `(dashboard)` route group. No role-filtering here
 * (Guest-facing, no session to check), unlike Dashboard Shell's
 * `useHasRole`-filtered nav.
 */
export function PublicShellClient() {
  const [open, setOpen] = useState(false);
  const drawerRef = useRef<HTMLDivElement>(null);
  const hamburgerRef = useRef<HTMLButtonElement>(null);
  const openLogin = useAuthModalStore((state) => state.openLogin);
  const openRegister = useAuthModalStore((state) => state.openRegister);

  useFocusTrap({
    active: open,
    containerRef: drawerRef,
    onEscape: () => setOpen(false),
  });

  return (
    <div className="md:hidden">
      <button
        ref={hamburgerRef}
        type="button"
        aria-expanded={open}
        aria-controls={DRAWER_ID}
        aria-label={open ? "Tutup menu" : "Buka menu"}
        onClick={() => setOpen((value) => !value)}
        className="inline-flex size-10 items-center justify-center rounded-full text-neutral-700 hover:bg-neutral-100"
      >
        {open ? (
          <X aria-hidden="true" className="size-5" />
        ) : (
          <Menu aria-hidden="true" className="size-5" />
        )}
      </button>

      {open && (
        <div
          id={DRAWER_ID}
          ref={drawerRef}
          role="dialog"
          aria-modal="true"
          aria-label="Navigasi"
          className="absolute inset-x-0 top-14 z-10 flex flex-col gap-1 border-b border-neutral-200 bg-white px-4 py-3 shadow-md"
        >
          {publicNavItems.map((item) => (
            <Link
              key={item.href}
              href={item.href}
              onClick={() => setOpen(false)}
              className="rounded-md px-3 py-2 text-body-sm font-medium text-neutral-700 hover:bg-neutral-100"
            >
              {item.label}
            </Link>
          ))}
          <div className="mt-2 flex flex-col gap-2 border-t border-neutral-200 pt-3">
            <button
              type="button"
              onClick={() => {
                setOpen(false);
                openLogin();
              }}
              className={`${authButtonBase} border border-neutral-200 bg-transparent text-neutral-700 hover:bg-neutral-100`}
            >
              Masuk
            </button>
            <button
              type="button"
              onClick={() => {
                setOpen(false);
                openRegister();
              }}
              className={`${authButtonBase} bg-primary-600 text-white shadow-sm hover:bg-primary-700`}
            >
              Daftar
            </button>
          </div>
        </div>
      )}
    </div>
  );
}
