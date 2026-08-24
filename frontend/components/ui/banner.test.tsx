import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { Banner } from "./banner";

describe("Banner", () => {
  it("uses role=alert for error/warning (assertive)", () => {
    render(<Banner variant="error">Email atau password salah</Banner>);
    expect(screen.getByRole("alert")).toHaveTextContent("Email atau password salah");
  });

  it("uses role=status for success/info (polite)", () => {
    render(<Banner variant="success">Berhasil disimpan</Banner>);
    expect(screen.getByRole("status")).toHaveTextContent("Berhasil disimpan");
  });
});
