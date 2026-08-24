// Public Shell — stubbed this phase (phase0-shared-infra.md Step 1).
// `/`, `/campaign`, `/campaign/[id]` belong to the `campaign` domain,
// not `account` — build the real shell (nav, header) when
// `campaign` domain's frontend track starts, per the Incremental
// Growth Rule. Deliberately pass-through, no nav — do not half-build
// this to match the Auth/Dashboard Shells built this phase.
export default function PublicLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <>{children}</>;
}
