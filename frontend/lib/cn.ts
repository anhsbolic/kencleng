/**
 * Joins conditional class name fragments into one string, skipping
 * falsy values. No conflict resolution (no `tailwind-merge`) — not
 * needed yet since every `components/ui/` primitive fully controls
 * its own class composition rather than merging arbitrary caller
 * classes on top of conflicting utilities. Revisit if that stops
 * being true.
 */
export function cn(...classes: Array<string | false | null | undefined>) {
  return classes.filter(Boolean).join(" ");
}
