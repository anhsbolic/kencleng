import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { Badge } from "./badge";

describe("Badge", () => {
  it("renders its label text", () => {
    render(<Badge tone="success">verified</Badge>);
    expect(screen.getByText("verified")).toBeInTheDocument();
  });

  it("applies the tone's background/text classes", () => {
    render(<Badge tone="error">rejected</Badge>);
    const badge = screen.getByText("rejected");
    expect(badge.className).toContain("bg-error-50");
    expect(badge.className).toContain("text-error-700");
  });
});
