import { fireEvent, render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { Label } from "@/components/ui/label";
import { PasswordInput } from "./password-input";

describe("PasswordInput", () => {
  it("defaults to type=password", () => {
    render(
      <>
        <Label htmlFor="pw">Password</Label>
        <PasswordInput id="pw" />
      </>
    );

    expect(screen.getByLabelText("Password")).toHaveAttribute("type", "password");
  });

  it("toggles type and the toggle button's accessible label", () => {
    render(
      <>
        <Label htmlFor="pw">Password</Label>
        <PasswordInput id="pw" />
      </>
    );

    const input = screen.getByLabelText("Password");
    fireEvent.click(screen.getByRole("button", { name: "Tampilkan password" }));
    expect(input).toHaveAttribute("type", "text");
    expect(screen.getByRole("button", { name: "Sembunyikan password" })).toBeInTheDocument();

    fireEvent.click(screen.getByRole("button", { name: "Sembunyikan password" }));
    expect(input).toHaveAttribute("type", "password");
  });

  it("forwards a field-level error to the underlying Input", () => {
    render(
      <>
        <Label htmlFor="pw">Password</Label>
        <PasswordInput id="pw" error="Password wajib diisi" />
      </>
    );

    expect(screen.getByText("Password wajib diisi")).toBeInTheDocument();
    expect(screen.getByLabelText("Password")).toHaveAttribute("aria-invalid", "true");
  });
});
