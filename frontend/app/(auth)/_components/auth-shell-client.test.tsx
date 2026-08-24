import { fireEvent, render, screen } from "@testing-library/react";
import { useState } from "react";
import { afterEach, describe, expect, it, vi } from "vitest";
import { AuthShellClient } from "./auth-shell-client";

function mockDesktop(matches: boolean) {
  window.matchMedia = vi.fn().mockImplementation((query: string) => ({
    matches,
    media: query,
    onchange: null,
    addEventListener: vi.fn(),
    removeEventListener: vi.fn(),
    addListener: vi.fn(),
    removeListener: vi.fn(),
    dispatchEvent: vi.fn(),
  })) as unknown as typeof window.matchMedia;
}

function Form() {
  return (
    <form>
      <input aria-label="Email" />
      <input aria-label="Password" />
      <button type="submit">Masuk</button>
    </form>
  );
}

/**
 * Simulates "click a Masuk link, shell mounts; navigate away, shell
 * unmounts" within one persistent tree, so the trigger button
 * survives the transition and `document.activeElement` can be
 * observed across it. A real cross-route Next.js navigation can
 * instead detach the trigger from the DOM before this hook's plain
 * `useEffect` cleanup runs (browsers auto-blur to `<body>` when the
 * focused node is removed) — this hook's return-focus guarantee is
 * strongest when the trigger persists across the transition, which
 * is exactly how the Dashboard Shell's hamburger button behaves
 * across drawer open/close (a state toggle, not an unmount).
 */
function ToggleHarness() {
  const [shellOpen, setShellOpen] = useState(false);
  return (
    <div>
      <button onClick={() => setShellOpen(true)}>Buka form masuk</button>
      {shellOpen && (
        <AuthShellClient>
          <Form />
        </AuthShellClient>
      )}
      <button onClick={() => setShellOpen(false)}>Simulate navigate away</button>
    </div>
  );
}

describe("AuthShellClient", () => {
  afterEach(() => vi.restoreAllMocks());

  it("desktop: renders dialog semantics and moves focus to the first field on mount", async () => {
    mockDesktop(true);
    render(
      <AuthShellClient>
        <Form />
      </AuthShellClient>
    );

    expect(screen.getByRole("dialog")).toBeInTheDocument();
    expect(await screen.findByLabelText("Email")).toHaveFocus();
  });

  it("desktop: Tab/Shift+Tab cycles within the panel only", async () => {
    mockDesktop(true);
    render(
      <AuthShellClient>
        <Form />
      </AuthShellClient>
    );

    const email = await screen.findByLabelText("Email");
    const submit = screen.getByRole("button", { name: "Masuk" });
    expect(email).toHaveFocus();

    fireEvent.keyDown(document, { key: "Tab", shiftKey: true });
    expect(submit).toHaveFocus(); // wraps from first field to last element

    fireEvent.keyDown(document, { key: "Tab" });
    expect(email).toHaveFocus(); // wraps from last element back to first field
  });

  it("desktop: returns focus to the trigger that opened the shell once it unmounts", async () => {
    mockDesktop(true);
    render(<ToggleHarness />);

    const trigger = screen.getByRole("button", { name: "Buka form masuk" });
    trigger.focus();
    fireEvent.click(trigger);

    expect(await screen.findByLabelText("Email")).toHaveFocus(); // shell took focus on mount

    fireEvent.click(screen.getByRole("button", { name: "Simulate navigate away" }));

    expect(trigger).toHaveFocus();
  });

  it("mobile: does not apply dialog semantics or trap focus", () => {
    mockDesktop(false);
    render(
      <AuthShellClient>
        <Form />
      </AuthShellClient>
    );

    expect(screen.queryByRole("dialog")).not.toBeInTheDocument();
  });
});
