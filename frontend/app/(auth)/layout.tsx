import { AuthShellClient } from "./_components/auth-shell-client";

// Auth Shell — built this phase (phase0-shared-infra.md Step 3).
// Thin Server Component wrapper: the actual modal/page-split and
// focus-management behavior needs client-side state (breakpoint
// detection, focus trap), so it lives in `AuthShellClient` — keeps
// the client boundary at the leaf that actually needs it, not
// hoisted to this whole layout (server-client-component-boundary.md).
export default function AuthLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <AuthShellClient>{children}</AuthShellClient>;
}
