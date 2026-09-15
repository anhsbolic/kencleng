import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { Hero } from "./hero";

describe("Hero", () => {
  it("uses the established responsive landing type scale", () => {
    render(<Hero />);

    const heading = screen.getByRole("heading", {
      name: "Berbagi itu mudah, dampaknya nyata",
    });
    expect(heading.className).toContain("text-h1");
    expect(heading.className).toContain("md:text-display");
  });

  it("renders no fabricated organization count", () => {
    render(<Hero />);
    expect(screen.queryByText(/\d+\s*organisasi/i)).not.toBeInTheDocument();
  });
});
