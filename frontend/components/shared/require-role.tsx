"use client";

import { useHasRole } from "@/lib/hooks/use-has-role";
import type { GlobalRole } from "@/lib/types/roles";
import type { ReactNode } from "react";

interface RequireRoleProps {
  roles: GlobalRole[];
  children: ReactNode;
  /** Rendered instead when the current user doesn't have any of `roles`. Omit to render nothing. */
  fallback?: ReactNode;
}

/**
 * Thin wrapper over `useHasRole`, per the hybrid role-gating
 * decision (`phase0-shared-infra.md` Step 4). No page uses this yet
 * — Account domain's Serial-group-S1 has no role-gated *page
 * content*, only role-gated *nav items* (Dashboard Shell uses
 * `useHasRole` directly for that). Scaffolded alongside the hook so
 * the pair ships together.
 */
export function RequireRole({ roles, children, fallback = null }: RequireRoleProps) {
  const hasRole = useHasRole(roles);
  return hasRole ? <>{children}</> : <>{fallback}</>;
}
