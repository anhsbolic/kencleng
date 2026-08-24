import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen, waitFor, within } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { describe, expect, it } from "vitest";
import { server } from "@/mocks/server";
import { DashboardShellClient, NavLink } from "./dashboard-shell-client";

function mockMe(roles: ("admin" | "kurator")[]) {
  server.use(
    http.get("/account/me", () =>
      HttpResponse.json({
        id: "u1",
        name: "Test User",
        email: "test@example.com",
        email_verified: true,
        roles,
        auth_providers: ["email_password"],
        mfa_enabled: false,
        created_at: new Date().toISOString(),
      })
    )
  );
}

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

describe("NavLink role filtering", () => {
  it("hides an item the current user's role isn't in", async () => {
    mockMe([]); // plain donatur, no elevated role
    render(
      withQueryClient(
        <NavLink item={{ label: "Admin Panel", href: "/dashboard/admin", roles: ["admin"] }} />
      )
    );

    // Safe-default-false while loading — nothing renders at any point.
    await waitFor(() => expect(screen.queryByText("Admin Panel")).not.toBeInTheDocument());
  });

  it("shows an item once the resolved role matches — a different GlobalRole combination than the donatur-only case above", async () => {
    mockMe(["admin"]);
    render(
      withQueryClient(
        <NavLink item={{ label: "Admin Panel", href: "/dashboard/admin", roles: ["admin"] }} />
      )
    );

    expect(await screen.findByText("Admin Panel")).toBeInTheDocument();
  });

  it("shows an item open to 'donatur' for a plain logged-in user with no elevated role", async () => {
    mockMe([]);
    render(
      withQueryClient(
        <NavLink item={{ label: "Notifikasi", href: "/dashboard/notifications", roles: ["donatur", "kurator", "admin"] }} />
      )
    );

    expect(await screen.findByText("Notifikasi")).toBeInTheDocument();
  });
});

describe("DashboardShellClient mobile drawer", () => {
  it("toggles aria-expanded, traps focus, and returns focus to the hamburger button on Escape", async () => {
    mockMe([]);
    render(withQueryClient(<DashboardShellClient>Konten</DashboardShellClient>));

    const hamburger = screen.getByRole("button", { name: "Buka menu" });
    expect(hamburger).toHaveAttribute("aria-expanded", "false");

    // A real click focuses the button first (native browser
    // behavior); `fireEvent.click` alone doesn't simulate that in
    // jsdom, so it's done explicitly here.
    hamburger.focus();
    fireEvent.click(hamburger);
    expect(hamburger).toHaveAttribute("aria-expanded", "true");

    // Scoped to the drawer specifically — the desktop nav renders the
    // same links (hidden via `md:flex`/`md:hidden` CSS classes, which
    // jsdom doesn't apply), so an unscoped query would match both.
    const drawer = within(await screen.findByRole("dialog"));

    // Focus lands on the drawer's first nav item, not the hamburger
    // button itself.
    const profil = await drawer.findByRole("link", { name: "Profil" });
    await waitFor(() => expect(document.activeElement).toBe(profil));

    const notifikasi = drawer.getByRole("link", { name: "Notifikasi" });

    // Shift+Tab from the first item wraps to the last — the trap
    // cycles within the drawer's three items only, per
    // `use-focus-trap.ts`'s boundary-wrap behavior.
    fireEvent.keyDown(document, { key: "Tab", shiftKey: true });
    expect(document.activeElement).toBe(notifikasi);

    // Tab from the last item wraps back to the first.
    fireEvent.keyDown(document, { key: "Tab" });
    expect(document.activeElement).toBe(profil);

    fireEvent.keyDown(document, { key: "Escape" });
    expect(hamburger).toHaveAttribute("aria-expanded", "false");
    await waitFor(() => expect(document.activeElement).toBe(hamburger));
  });
});
