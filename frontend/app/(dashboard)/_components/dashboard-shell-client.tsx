"use client";

import { Menu, X } from "lucide-react";
import Link from "next/link";
import { useRef, useState, type ReactNode } from "react";
import { Button } from "@/components/ui/button";
import { useAccountMe } from "@/lib/hooks/use-account-me";
import { useFocusTrap } from "@/lib/hooks/use-focus-trap";
import { useHasRole } from "@/lib/hooks/use-has-role";
import { useLogout } from "@/lib/hooks/use-logout";
import { NotificationBadge } from "./notification-badge";
import { navItems } from "./nav-items";

/**
 * "Keluar" (logout) button — visible to any authenticated user
 * regardless of role, so gated directly on `useAccountMe()`'s `data`
 * rather than `useHasRole`'s role-array shape (a role list is the wrong
 * primitive for "is anyone logged in at all") — techplan account/03-
 * login-session-management, task-04, R18.
 */
function LogoutButton() {
  const { data: user } = useAccountMe();
  const logoutMutation = useLogout();

  if (!user) return null;

  return (
    <Button
      type="button"
      variant="outline"
      size="sm"
      loading={logoutMutation.isPending}
      onClick={() => logoutMutation.mutate()}
    >
      Keluar
    </Button>
  );
}

const DRAWER_ID = "dashboard-mobile-nav";

function FilteredNavLinks({
  className,
  onNavigate,
}: {
  className?: string;
  onNavigate?: () => void;
}) {
  return (
    <>
      {navItems.map((item) => (
        <NavLink key={item.href} item={item} className={className} onNavigate={onNavigate} />
      ))}
    </>
  );
}

// Exported for testing — lets the role-filtering test cover the
// mechanism directly with synthetic `NavItem`s requiring different
// roles, rather than being limited to whatever `nav-items.ts`'s real
// (currently all-roles) data happens to contain today.
export function NavLink({
  item,
  className,
  onNavigate,
}: {
  item: (typeof navItems)[number];
  className?: string;
  onNavigate?: () => void;
}) {
  const canSee = useHasRole(item.roles);
  if (!canSee) return null;

  return (
    <Link href={item.href} onClick={onNavigate} className={className}>
      {item.label}
    </Link>
  );
}

/**
 * Dashboard Shell — top-nav desktop, top-bar + hamburger mobile
 * (`patterns.md` Pattern 4). Nav items are role-filtered via
 * `useHasRole` (safe-default-false while the profile query loads —
 * see that hook's doc comment), and the mobile drawer reuses the
 * same filtered list rather than a second copy.
 */
export function DashboardShellClient({ children }: { children: ReactNode }) {
  const [drawerOpen, setDrawerOpen] = useState(false);
  const drawerRef = useRef<HTMLDivElement>(null);
  const hamburgerRef = useRef<HTMLButtonElement>(null);

  useFocusTrap({
    active: drawerOpen,
    containerRef: drawerRef,
    onEscape: () => setDrawerOpen(false),
  });

  return (
    <div className="flex min-h-full flex-1 flex-col">
      <header className="flex h-16 items-center justify-between border-b border-neutral-200 px-4 md:px-6">
        <Link href="/dashboard/profile" className="font-heading text-lg font-bold text-neutral-900">
          Kencleng
        </Link>

        {/* Desktop top-nav — hidden at mobile width, so it's out of
            the tab order there instead of needing the focus trap to
            fight it. */}
        <nav className="hidden items-center gap-6 md:flex" aria-label="Navigasi dashboard">
          <FilteredNavLinks className="text-sm font-medium text-neutral-700 hover:text-primary-700" />
        </nav>

        <div className="flex items-center gap-2">
          <LogoutButton />
          <NotificationBadge />
          <button
            ref={hamburgerRef}
            type="button"
            aria-expanded={drawerOpen}
            aria-controls={DRAWER_ID}
            aria-label={drawerOpen ? "Tutup menu" : "Buka menu"}
            onClick={() => setDrawerOpen((open) => !open)}
            className="inline-flex size-10 items-center justify-center rounded-full text-neutral-700 hover:bg-neutral-100 md:hidden"
          >
            {drawerOpen ? (
              <X aria-hidden="true" className="size-5" />
            ) : (
              <Menu aria-hidden="true" className="size-5" />
            )}
          </button>
        </div>
      </header>

      {drawerOpen && (
        <div
          id={DRAWER_ID}
          ref={drawerRef}
          role="dialog"
          aria-modal="true"
          aria-label="Navigasi dashboard"
          className="flex flex-col gap-1 border-b border-neutral-200 bg-white px-4 py-3 md:hidden"
        >
          <FilteredNavLinks
            className="rounded-md px-3 py-2 text-sm font-medium text-neutral-700 hover:bg-neutral-100"
            onNavigate={() => setDrawerOpen(false)}
          />
        </div>
      )}

      <main className="flex flex-1 flex-col p-4 md:p-6">{children}</main>
    </div>
  );
}
