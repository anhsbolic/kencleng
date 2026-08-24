"use client";

import { Bell } from "lucide-react";
import Link from "next/link";
import { useUnreadCount } from "@/lib/hooks/use-unread-count";

/**
 * Persistent header element for any logged-in user (`page-map.md`
 * Cross-Cutting UI Elements). Mocked data, not placeholder UI — see
 * `lib/api/notification.ts`.
 *
 * Loading/error both degrade to the plain bell icon with no numeric
 * badge — showing "0" while loading would misreport an unknown count
 * as "definitely zero," and this element is too small/persistent to
 * carry a retry-banner error state of its own.
 */
export function NotificationBadge() {
  const { data, isSuccess } = useUnreadCount();
  const count = isSuccess ? data.unread_count : 0;
  const showBadge = isSuccess && count > 0;

  return (
    <Link
      href="/dashboard/notifications"
      aria-label={showBadge ? `Notifikasi, ${count} belum dibaca` : "Notifikasi"}
      className="relative inline-flex size-10 items-center justify-center rounded-full text-neutral-700 hover:bg-neutral-100"
    >
      <Bell aria-hidden="true" className="size-5" />
      {showBadge && (
        <span
          aria-hidden="true"
          className="absolute -top-0.5 -right-0.5 flex h-4 min-w-4 items-center justify-center rounded-full bg-error-500 px-1 text-[10px] font-semibold text-white"
        >
          {count > 99 ? "99+" : count}
        </span>
      )}
    </Link>
  );
}
