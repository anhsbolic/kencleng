import { DashboardShellClient } from "./_components/dashboard-shell-client";

// Dashboard Shell — built this phase (phase0-shared-infra.md Step
// 5). Thin Server Component wrapper — the client boundary lives in
// `DashboardShellClient`, not here (server-client-component-boundary.md).
export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <DashboardShellClient>{children}</DashboardShellClient>;
}
