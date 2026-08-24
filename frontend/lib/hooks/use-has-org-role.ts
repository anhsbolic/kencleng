import type { OrgRoleLevel } from "@/lib/types/roles";

/**
 * Org-scoped role check, per the role-gating decision
 * (`phase0-shared-infra.md` Step 4). No page consumes this yet
 * (Account domain's Serial-group-S1 tasks have no org-role-gated
 * *page content*, only Dashboard Shell's nav, which is
 * account-wide — see `use-has-role.ts`), and no endpoint returning
 * the current user's role for an arbitrary organization exists in
 * this frontend yet either — the closest candidate is an
 * organization endpoint's `my_level` field (`api/openapi.yaml`), but
 * that's `organization` domain's frontend track to wire, not
 * Account's.
 *
 * Scaffolded now (safe default `false`, never throws) alongside
 * `use-has-role.ts` so the pair doesn't split across two separate
 * playbook runs — replace the body with a real query once
 * `organization` domain's frontend track defines where a user's
 * per-org level actually comes from.
 */
export function useHasOrgRole(orgId: string, levels: OrgRoleLevel[]): boolean {
  void orgId;
  void levels;
  return false;
}
