// Public Shell nav item list — Guest-facing, no role-gating (unlike
// Dashboard Shell's nav, there's no session to check here).
// `scaffold-public-shell.md` Step 1.
export interface PublicNavItem {
  label: string;
  href: string;
}

export const publicNavItems: PublicNavItem[] = [
  { label: "Beranda", href: "/" },
  // Anchor, not `/campaign` — that route doesn't exist yet
  // (`page-map.md`'s Public Shell nav decision predates `campaign`
  // domain's own frontend track). Swap to a real `/campaign` href in
  // one line once that task ships — see scaffold-public-shell.md.
  { label: "Jelajahi Kampanye", href: "#kampanye" },
];
