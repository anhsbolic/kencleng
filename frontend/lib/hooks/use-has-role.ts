import { useAccountMe } from "./use-account-me";
import type { GlobalRole } from "@/lib/types/roles";

/**
 * OR-logic role check against the current user's effective global
 * roles, per the role-gating decision (`phase0-shared-infra.md`
 * Step 4). Every authenticated user is implicitly `'donatur'` — the
 * API's `Role` enum only carries the elevated roles a user is
 * explicitly granted (`admin`/`kurator`), so `'donatur'` is added
 * here rather than expected from the response (see
 * `lib/types/roles.ts`).
 *
 * Returns `false` while the profile is loading, unauthenticated, or
 * failed to load — a safe default (hide, don't show) rather than
 * throwing, so a nav item never flashes visible before the real
 * answer is known.
 */
export function useHasRole(roles: GlobalRole[]): boolean {
  const { data: user } = useAccountMe();
  if (!user) return false;

  const effectiveRoles: GlobalRole[] = ["donatur", ...(user.roles ?? [])];
  return roles.some((role) => effectiveRoles.includes(role));
}
