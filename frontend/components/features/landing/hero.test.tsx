import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { Hero } from "./hero";

describe("Hero", () => {
  it("uses design-guidelines.md's real type-scale tokens, not hardcoded prototype pixel values (R13)", () => {
    render(<Hero />);

    const heading = screen.getByRole("heading", {
      name: "Berbagi itu mudah, dampaknya nyata",
    });
    // design-guidelines.md: h1=30px (mobile), display=36px (desktop,
    // "Landing/hero only") — never the prototype's drifted 44px/40px/48px.
    expect(heading.className).toContain("text-h1");
    expect(heading.className).toContain("md:text-display");
  });

  it("renders no fabricated organization count (Decision 3's reasoning applied to the hero badge too)", () => {
    render(<Hero />);
    expect(screen.queryByText(/\d+\s*organisasi/i)).not.toBeInTheDocument();
  });
});
