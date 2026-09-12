import { fireEvent, render, screen } from "@testing-library/react";
import { describe, expect, it, vi } from "vitest";
import { Button } from "./button";

describe("Button", () => {
  it("renders its label and fires onClick", () => {
    const onClick = vi.fn();
    render(<Button onClick={onClick}>Kirim</Button>);

    const button = screen.getByRole("button", { name: "Kirim" });
    fireEvent.click(button);

    expect(onClick).toHaveBeenCalledOnce();
  });

  it("disables itself and shows a spinner while loading, without losing its accessible name", () => {
    render(<Button loading>Kirim</Button>);

    const button = screen.getByRole("button", { name: "Kirim" });
    expect(button).toBeDisabled();
    expect(button).toHaveAttribute("aria-busy", "true");
  });

  it("renders secondary as a neutral outlined action instead of an accent-filled action", () => {
    render(<Button variant="secondary">Batal</Button>);

    const button = screen.getByRole("button", { name: "Batal" });
    expect(button).toHaveClass(
      "bg-transparent",
      "text-neutral-700",
      "border",
      "border-neutral-200",
      "hover:bg-neutral-100"
    );
    expect(button).not.toHaveClass(
      "bg-accent-500",
      "text-neutral-900",
      "shadow-sm",
      "hover:bg-accent-600"
    );
  });
});
