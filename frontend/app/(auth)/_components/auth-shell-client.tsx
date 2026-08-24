"use client";

import { useFocusTrap } from "@/lib/hooks/use-focus-trap";
import { useRef, useSyncExternalStore, type ReactNode } from "react";

// Matches Tailwind's `md:` breakpoint — the same one the CSS layout
// switch below uses, so the JS-side focus-trap gating and the visual
// modal-vs-page switch never disagree about which mode is active.
const DESKTOP_QUERY = "(min-width: 768px)";

function subscribe(callback: () => void) {
  const mql = window.matchMedia(DESKTOP_QUERY);
  mql.addEventListener("change", callback);
  return () => mql.removeEventListener("change", callback);
}

/**
 * Subscribes to the same breakpoint the CSS layout switch below
 * uses — via `useSyncExternalStore`, not `useState`+`useEffect`, so
 * there's no synchronous `setState` inside an effect body. This hook
 * only tells the focus trap whether it should be active; the visual
 * layout itself never branches on it, it stays pure CSS (`md:`
 * classes).
 */
function useIsDesktop() {
  return useSyncExternalStore(
    subscribe,
    () => window.matchMedia(DESKTOP_QUERY).matches,
    () => false // SSR snapshot — matches server-rendered markup (no `window`)
  );
}

/**
 * Auth Shell — desktop renders a centered modal overlay, mobile
 * renders a plain full page, per `prototype-reference.md`'s `/login`
 * Tier 1 entry. The switch is CSS-only (`md:` classes); only the
 * focus-trap *activation* is gated in JS, and only because trapping
 * Tab makes no sense on the full-page mobile variant.
 *
 * Convention for pages rendered inside this shell (Account Task
 * #1/#3's job, not this phase's): render a `<Banner variant="error">`
 * as the first child, before the form — this panel is a plain
 * `flex flex-col gap-6`, so whatever a page puts first renders above
 * the form with no extra wiring. This is the known-issue guard from
 * `prototype-reference.md` (`/login`'s auth failure must be a banner,
 * not a field-level error) made structurally easy to get right.
 */
export function AuthShellClient({ children }: { children: ReactNode }) {
  const isDesktop = useIsDesktop();
  const panelRef = useRef<HTMLDivElement>(null);

  useFocusTrap({ active: isDesktop, containerRef: panelRef });

  return (
    <div className="flex min-h-full flex-1 items-center justify-center md:bg-neutral-900/40 md:p-6">
      <div
        ref={panelRef}
        role={isDesktop ? "dialog" : undefined}
        aria-modal={isDesktop ? true : undefined}
        className="flex w-full flex-1 flex-col gap-6 bg-white p-6 md:max-w-md md:flex-none md:rounded-xl md:p-8 md:shadow-lg"
      >
        {children}
      </div>
    </div>
  );
}
