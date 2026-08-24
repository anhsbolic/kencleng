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
});
