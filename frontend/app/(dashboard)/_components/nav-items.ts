import type { GlobalRole } from "@/lib/types/roles";

export interface NavItem {
  label: string;
  href: string;
  roles: GlobalRole[];
}

// Starts small — only what Account domain's own pages need right now
// (phase0-shared-infra.md Step 5). Other domains' items (kurasi,
// disbursement, admin panels) get added here by *those* domains' own
// tasks, per the Incremental Growth Rule — this is the Shell's data,
// not its structure, so extending it later doesn't mean rebuilding
// the Shell.
export const navItems: NavItem[] = [
  {
    label: "Profil",
    href: "/dashboard/profile",
    roles: ["donatur", "kurator", "admin"],
  },
  {
    label: "Keamanan",
    href: "/dashboard/security",
    roles: ["donatur", "kurator", "admin"],
  },
  {
    label: "Notifikasi",
    href: "/dashboard/notifications",
    // Available to any logged-in user (page-map.md) — expressed as
    // "all current roles" rather than a special-cased "no roles
    // required" branch, so the filtering logic stays uniform.
    roles: ["donatur", "kurator", "admin"],
  },
];
