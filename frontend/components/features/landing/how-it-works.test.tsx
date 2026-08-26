import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { HowItWorks } from "./how-it-works";

describe("HowItWorks", () => {
  it("renders all three steps", () => {
    render(<HowItWorks />);

    expect(screen.getByText("Temukan kampanye terverifikasi")).toBeInTheDocument();
    expect(screen.getByText("Donasi aman dengan berbagai metode")).toBeInTheDocument();
    expect(screen.getByText("Pantau transparansi penyaluran dana")).toBeInTheDocument();
  });
});
