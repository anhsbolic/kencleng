import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { Label } from "./label";

describe("Label", () => {
  it("renders its text", () => {
    render(<Label htmlFor="name">Nama</Label>);
    expect(screen.getByText("Nama")).toBeInTheDocument();
  });
});
