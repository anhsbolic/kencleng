import { HandHeart } from "lucide-react";
import Link from "next/link";
import { AuthModal } from "@/components/features/account/auth-modal";
import { AuthModalTriggers } from "@/components/features/account/auth-modal-triggers";
import { publicNavItems } from "./_components/nav-items";
import { PublicShellClient } from "./_components/public-shell-client";

/**
 * Public Shell — real nav, per `page-map.md`'s Public Shell decision
 * (2026-08-24) and `scaffold-public-shell.md`. Replaces the Phase 0
 * pass-through stub (`phase0-shared-infra.md` Step 1).
 *
 * Desktop nav renders directly here (static markup, no interactivity
 * needed) except the "Masuk"/"Daftar" buttons (`AuthModalTriggers`) and
 * the mobile hamburger/drawer (`PublicShellClient`), both Client
 * Components, per `server-client-component-boundary.md`'s "'use
 * client' at the smallest leaf" checklist item.
 *
 * `AuthModal` is mounted once here, as a sibling of `{children}` — not
 * inside `<main>` — so it overlays the actual mounted landing page
 * (and any other page under this Shell) rather than replacing it, per
 * the login/register-as-a-modal follow-up to techplan account/03-
 * login-session-management.
 */
export default function PublicLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex min-h-full flex-1 flex-col">
      <header className="relative flex h-14 items-center gap-4 border-b border-neutral-200 bg-white px-4 md:h-17 md:px-6">
        <Link href="/" className="flex items-center gap-2.5">
          <span className="inline-flex size-8 items-center justify-center rounded-[10px] bg-primary-600">
            <HandHeart aria-hidden="true" className="size-4 text-white" />
          </span>
          <span className="font-heading text-lg font-extrabold text-neutral-900">
            Kencleng
          </span>
        </Link>

        <nav
          aria-label="Navigasi utama"
          className="hidden flex-1 items-center gap-1 md:flex"
        >
          {publicNavItems.map((item) => (
            <Link
              key={item.href}
              href={item.href}
              className="rounded-md px-3 py-2 text-body-sm font-medium text-neutral-700 hover:bg-neutral-100"
            >
              {item.label}
            </Link>
          ))}
        </nav>

        <div className="ml-auto hidden items-center gap-2 md:flex">
          <AuthModalTriggers />
        </div>

        <div className="ml-auto">
          <PublicShellClient />
        </div>
      </header>

      <main className="flex flex-1 flex-col">{children}</main>

      <AuthModal />
    </div>
  );
}
