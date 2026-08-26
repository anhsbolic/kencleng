"use client";

import { X } from "lucide-react";
import { useRef } from "react";
import { useFocusTrap } from "@/lib/hooks/use-focus-trap";
import { useAuthModalStore } from "@/lib/stores/auth-modal-store";
import { LoginForm } from "./login-form";
import { RegisterForm } from "./register-form";

/**
 * Landing-page login/register modal — follow-up to techplan account/03-
 * login-session-management: login/register on the landing page, as a
 * modal, instead of a full navigation to `/login`/`/register`. Mounted
 * once in the Public Shell (`app/(public)/layout.tsx`) so the actual
 * landing page stays mounted and visible behind the overlay, matching
 * the Tier 1 design reference's landing-page-behind-modal treatment —
 * achieved here simply by rendering this modal as a sibling of
 * `{children}` rather than as a separate route.
 *
 * `/login` and `/register` remain real, independent routes, unchanged
 * — needed for the Google OAuth callback's hardcoded `/login?error=
 * {code}` redirect target, and for direct/shared links. This modal is
 * an *additional* presentation of the same `LoginForm`/`RegisterForm`
 * components, not a replacement for the routes.
 *
 * The desktop-vs-mobile visual switch (centered panel vs. full-screen
 * takeover) is pure CSS (`md:` classes), matching `AuthShellClient`'s
 * own approach (`app/(auth)/_components/auth-shell-client.tsx`) — no
 * JS breakpoint detection needed here. Unlike `AuthShellClient`, the
 * focus trap below is active on BOTH breakpoints, not desktop-only:
 * this modal always overlays real, interactive landing-page content
 * behind it, whereas `AuthShellClient`'s mobile treatment is a plain
 * full page with nothing behind it to protect focus from — so there's
 * no equivalent JS gate needed for the trap either.
 */
export function AuthModal() {
  const mode = useAuthModalStore((state) => state.mode);
  const openLogin = useAuthModalStore((state) => state.openLogin);
  const openRegister = useAuthModalStore((state) => state.openRegister);
  const close = useAuthModalStore((state) => state.close);
  const panelRef = useRef<HTMLDivElement>(null);

  useFocusTrap({ active: mode !== null, containerRef: panelRef, onEscape: close });

  if (!mode) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center md:bg-neutral-900/40 md:p-6">
      {/* Mouse-only convenience — not part of the focus trap or tab
          order (Escape and the visible close button below are the
          accessible dismiss paths). */}
      <div aria-hidden="true" onClick={close} className="absolute inset-0" />

      <div
        ref={panelRef}
        role="dialog"
        aria-modal="true"
        aria-label={mode === "login" ? "Masuk" : "Daftar"}
        className="relative z-10 flex h-full w-full flex-col gap-6 overflow-y-auto bg-white p-6 md:h-auto md:max-w-md md:rounded-xl md:p-8 md:shadow-lg"
      >
        <div className="flex flex-col gap-2">
          <h1 className="text-2xl font-bold text-neutral-900">
            {mode === "login" ? "Masuk" : "Daftar"}
          </h1>
          <p className="text-sm text-neutral-500">
            {mode === "login" ? "Masuk ke akun Kencleng kamu." : "Buat akun Kencleng baru."}
          </p>
        </div>

        {mode === "login" ? (
          <LoginForm onSwitchToRegister={openRegister} />
        ) : (
          <RegisterForm onSwitchToLogin={openLogin} />
        )}

        {/* Last in DOM order (after the form) so the focus trap's
            "first focusable element" is the form's own first field,
            not this button — visually pinned top-right regardless. */}
        <button
          type="button"
          aria-label="Tutup"
          onClick={close}
          className="absolute right-4 top-4 inline-flex size-8 items-center justify-center rounded-full text-neutral-500 hover:bg-neutral-100"
        >
          <X aria-hidden="true" className="size-4" />
        </button>
      </div>
    </div>
  );
}
