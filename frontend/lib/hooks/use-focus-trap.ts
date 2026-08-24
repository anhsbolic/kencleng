import { useEffect, useRef, type RefObject } from "react";

const FOCUSABLE_SELECTOR =
  'a[href], button:not([disabled]), input:not([disabled]), select:not([disabled]), textarea:not([disabled]), [tabindex]:not([tabindex="-1"])';

interface UseFocusTrapOptions {
  active: boolean;
  containerRef: RefObject<HTMLElement | null>;
  /** Focused first when the trap activates. Falls back to the first focusable element inside the container. */
  initialFocusRef?: RefObject<HTMLElement | null>;
  /** Called when Escape is pressed while the trap is active. Omit if Escape shouldn't close this overlay. */
  onEscape?: () => void;
}

/**
 * Traps Tab/Shift+Tab inside `containerRef` while `active`, moves
 * focus in on activation, and returns focus to whatever had it
 * beforehand on deactivation — falling back to the same
 * activation-time focus target (`initialFocusRef`, or the
 * container's first focusable element) if nothing meaningful had
 * focus (e.g. a direct URL navigation left focus on `<body>`, per
 * the Auth Shell's spec). For an overlay
 * opened by clicking a trigger (e.g. the Dashboard Shell's hamburger
 * button), "whatever had focus beforehand" is that trigger, so this
 * one hook covers both Shells' return-focus requirements without
 * needing an explicit trigger ref.
 *
 * Shared by the Auth Shell modal and the Dashboard Shell mobile
 * drawer, per `accessibility-fundamentals.md`'s focus-management
 * pattern and the Incremental Growth Rule's note that any future
 * overlay component should reuse this rather than reinventing it.
 */
export function useFocusTrap({
  active,
  containerRef,
  initialFocusRef,
  onEscape,
}: UseFocusTrapOptions) {
  const previouslyFocusedRef = useRef<HTMLElement | null>(null);

  useEffect(() => {
    if (!active) return;

    const previouslyFocused = document.activeElement;
    previouslyFocusedRef.current =
      previouslyFocused instanceof HTMLElement ? previouslyFocused : null;

    // Tracks whatever focus target actually gets used, for the
    // close-time fallback below — not just the one found at this
    // instant, since it may only appear moments later (see the
    // MutationObserver branch).
    let resolvedFocusTarget: HTMLElement | null =
      initialFocusRef?.current ??
      containerRef.current?.querySelector<HTMLElement>(FOCUSABLE_SELECTOR) ??
      null;

    let contentObserver: MutationObserver | undefined;
    if (resolvedFocusTarget) {
      resolvedFocusTarget.focus();
    } else if (containerRef.current) {
      // Nothing focusable exists yet — the container's real content
      // (e.g. role-gated nav links behind an async `useHasRole`
      // check) hasn't rendered on this first commit. Watch for it to
      // appear and focus it the moment it does, instead of silently
      // never moving focus.
      const container = containerRef.current;
      contentObserver = new MutationObserver(() => {
        const target = container.querySelector<HTMLElement>(FOCUSABLE_SELECTOR);
        if (target) {
          resolvedFocusTarget = target;
          target.focus();
          contentObserver?.disconnect();
        }
      });
      contentObserver.observe(container, { childList: true, subtree: true });
    }

    function handleKeyDown(event: KeyboardEvent) {
      if (event.key === "Escape" && onEscape) {
        onEscape();
        return;
      }

      if (event.key !== "Tab" || !containerRef.current) return;

      const focusable = Array.from(
        containerRef.current.querySelectorAll<HTMLElement>(FOCUSABLE_SELECTOR)
      );
      if (focusable.length === 0) return;

      const first = focusable[0];
      const last = focusable[focusable.length - 1];

      if (event.shiftKey && document.activeElement === first) {
        event.preventDefault();
        last.focus();
      } else if (!event.shiftKey && document.activeElement === last) {
        event.preventDefault();
        first.focus();
      }
    }

    document.addEventListener("keydown", handleKeyDown);

    return () => {
      document.removeEventListener("keydown", handleKeyDown);
      contentObserver?.disconnect();

      // Don't return focus to <body> — that's the "focus silently
      // lost" failure this hook exists to prevent. Fall back to the
      // initial-focus target instead (e.g. the modal's first field).
      const rememberedTrigger =
        previouslyFocusedRef.current && previouslyFocusedRef.current !== document.body
          ? previouslyFocusedRef.current
          : null;
      (rememberedTrigger ?? resolvedFocusTarget)?.focus();
    };
    // Deliberately re-runs only on `active` transitions — containerRef/
    // initialFocusRef/onEscape identity churning every render must not
    // reset the remembered "previously focused" element mid-session.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [active]);
}
