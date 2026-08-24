import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { Input } from "./input";
import { Label } from "./label";

describe("Input", () => {
  it("associates with its Label via htmlFor/id", () => {
    render(
      <>
        <Label htmlFor="email">Email</Label>
        <Input id="email" />
      </>
    );

    expect(screen.getByLabelText("Email")).toBeInTheDocument();
  });

  it("links a field-level error via aria-describedby and aria-invalid", () => {
    render(<Input id="email" error="Email tidak valid" />);

    const input = screen.getByRole("textbox");
    expect(input).toHaveAttribute("aria-invalid", "true");
    expect(input).toHaveAttribute("aria-describedby", "email-error");
    expect(screen.getByText("Email tidak valid")).toHaveAttribute("id", "email-error");
  });

  it("has no aria-invalid/aria-describedby when there is no error", () => {
    render(<Input id="email" />);

    const input = screen.getByRole("textbox");
    expect(input).not.toHaveAttribute("aria-invalid");
    expect(input).not.toHaveAttribute("aria-describedby");
  });
});
