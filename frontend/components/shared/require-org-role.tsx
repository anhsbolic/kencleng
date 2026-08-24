"use client";

import { useHasOrgRole } from "@/lib/hooks/use-has-org-role";
import type { OrgRoleLevel } from "@/lib/types/roles";
import type { ReactNode } from "react";

interface RequireOrgRoleProps {
  orgId: string;
  levels: OrgRoleLevel[];
  children: ReactNode;
  /** Rendered instead when the current user doesn't have any of `levels` for `orgId`. Omit to render nothing. */
  fallback?: ReactNode;
}

/**
 * Thin wrapper over `useHasOrgRole` — see that hook's doc comment
 * for why it's a safe-default stub for now. No page uses this yet.
 */
export function RequireOrgRole({
  orgId,
  levels,
  children,
  fallback = null,
}: RequireOrgRoleProps) {
  const hasOrgRole = useHasOrgRole(orgId, levels);
  return hasOrgRole ? <>{children}</> : <>{fallback}</>;
}
