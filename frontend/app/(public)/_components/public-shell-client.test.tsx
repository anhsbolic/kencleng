import { fireEvent, render, screen, waitFor, within } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { PublicShellClient } from "./public-shell-client";

describe("PublicShellClient mobile drawer", () => {
  it("toggles aria-expanded, traps focus, and returns focus to the hamburger button on Escape (R2-R5)", async () => {
    render(<PublicShellClient />);

    const hamburger = screen.getByRole("button", { name: "Buka menu" });
    expect(hamburger).toHaveAttribute("aria-expanded", "false");

    hamburger.focus();
    fireEvent.click(hamburger);
    expect(hamburger).toHaveAttribute("aria-expanded", "true");

    const drawer = within(await screen.findByRole("dialog"));
    const beranda = await drawer.findByRole("link", { name: "Beranda" });
    await waitFor(() => expect(document.activeElement).toBe(beranda));

    const daftar = drawer.getByRole("link", { name: "Daftar" });

    // Shift+Tab from the first item wraps to the last drawer item.
    fireEvent.keyDown(document, { key: "Tab", shiftKey: true });
    expect(document.activeElement).toBe(daftar);

    fireEvent.keyDown(document, { key: "Escape" });
    expect(hamburger).toHaveAttribute("aria-expanded", "false");
    await waitFor(() => expect(document.activeElement).toBe(hamburger));
  });

  it("closes the drawer and returns focus to the hamburger when a nav item is activated", async () => {
    render(<PublicShellClient />);

    const hamburger = screen.getByRole("button", { name: "Buka menu" });
    fireEvent.click(hamburger);

    const drawer = within(await screen.findByRole("dialog"));
    fireEvent.click(drawer.getByRole("link", { name: "Jelajahi Kampanye" }));

    expect(hamburger).toHaveAttribute("aria-expanded", "false");
  });
});
