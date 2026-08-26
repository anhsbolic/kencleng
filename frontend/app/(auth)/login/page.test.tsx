import { render, screen } from "@testing-library/react";
import { describe, expect, it, vi } from "vitest";
import LoginPage from "./page";

vi.mock("next/navigation", () => ({
  useSearchParams: () => ({ get: () => null }),
}));

describe("LoginPage", () => {
  it("shows a heading, the 'coming soon' note, and the Google button — no credential fields (R4)", () => {
    render(<LoginPage />);

    expect(screen.getByRole("heading", { name: "Masuk" })).toBeInTheDocument();
    expect(screen.getByText(/segera hadir/i)).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Masuk dengan Google" })).toHaveAttribute(
      "href",
      "/auth/google/redirect?intent=login"
    );
    expect(screen.queryByLabelText(/email/i)).not.toBeInTheDocument();
    expect(screen.queryByLabelText(/password/i)).not.toBeInTheDocument();
  });
});
