import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { ProgressBar } from "./progress-bar";

describe("ProgressBar", () => {
  it("renders primary-600 fill below 100%", () => {
    render(<ProgressBar value={68} />);
    const bar = screen.getByRole("progressbar");
    expect(bar).toHaveAttribute("aria-valuenow", "68");
    expect(bar.firstElementChild?.className).toContain("bg-primary-600");
  });

  it("switches to success-500 fill at 100% (R11)", () => {
    render(<ProgressBar value={100} />);
    const bar = screen.getByRole("progressbar");
    expect(bar.firstElementChild?.className).toContain("bg-success-500");
  });

  it("clamps values outside 0-100", () => {
    render(<ProgressBar value={140} />);
    expect(screen.getByRole("progressbar")).toHaveAttribute("aria-valuenow", "100");
  });
});
