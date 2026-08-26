import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen, waitFor, within } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { server } from "@/mocks/server";
import * as authChannel from "@/lib/api/auth-channel";
import { useAuthStore } from "@/lib/stores/auth-store";
import { DashboardShellClient, NavLink } from "./dashboard-shell-client";

vi.mock("@/lib/api/auth-channel", async () => {
  const actual = await vi.importActual<typeof import("@/lib/api/auth-channel")>(
    "@/lib/api/auth-channel"
  );
  return { ...actual, postAuthChannelMessage: vi.fn() };
});

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

beforeEach(() => {
  useAuthStore.setState({ accessToken: "some-token" });
  vi.mocked(authChannel.postAuthChannelMessage).mockClear();
});

describe("LogoutButton (techplan account/03-login-session-management, R17/R18)", () => {
  it("is absent while useAccountMe has no data yet", () => {
    render(withQueryClient(<DashboardShellClient>Konten</DashboardShellClient>));

    expect(screen.queryByRole("button", { name: "Keluar" })).not.toBeInTheDocument();
  });

  it("appears once useAccountMe resolves with data (R18)", async () => {
    mockMe([]);
    render(withQueryClient(<DashboardShellClient>Konten</DashboardShellClient>));

    expect(await screen.findByRole("button", { name: "Keluar" })).toBeInTheDocument();
  });

  it("clicking it always clears the store, clears the query cache, and broadcasts logout — even on a failed request (R17)", async () => {
    mockMe([]);
    server.use(http.post("/auth/logout", () => HttpResponse.json({}, { status: 500 })));

    const queryClient = new QueryClient({ defaultOptions: { queries: { retry: false } } });
    const clearSpy = vi.spyOn(queryClient, "clear");
    render(
      <QueryClientProvider client={queryClient}>
        <DashboardShellClient>Konten</DashboardShellClient>
      </QueryClientProvider>
    );
    const button = await screen.findByRole("button", { name: "Keluar" });

    fireEvent.click(button);

    await waitFor(() => expect(useAuthStore.getState().accessToken).toBeNull());
    expect(clearSpy).toHaveBeenCalled();
    expect(authChannel.postAuthChannelMessage).toHaveBeenCalledWith({ type: "logged-out" });
  });

  it("clicking it on a successful request produces the identical cleanup (R17)", async () => {
    mockMe([]);
    server.use(http.post("/auth/logout", () => new HttpResponse(null, { status: 204 })));

    render(withQueryClient(<DashboardShellClient>Konten</DashboardShellClient>));
    const button = await screen.findByRole("button", { name: "Keluar" });

    fireEvent.click(button);

    await waitFor(() => expect(useAuthStore.getState().accessToken).toBeNull());
    expect(authChannel.postAuthChannelMessage).toHaveBeenCalledWith({ type: "logged-out" });
  });
});

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
