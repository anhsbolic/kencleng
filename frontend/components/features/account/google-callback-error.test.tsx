import { render, screen, waitFor } from "@testing-library/react";
import { describe, expect, it, vi } from "vitest";
import { GoogleCallbackError } from "./google-callback-error";

let mockError: string | null = null;

vi.mock("next/navigation", () => ({
  useSearchParams: () => ({
    get: (key: string) => (key === "error" ? mockError : null),
  }),
}));

describe("GoogleCallbackError", () => {
  it("renders nothing when no error param is present (R5)", () => {
    mockError = null;
    const { container } = render(<GoogleCallbackError />);

    expect(container).toBeEmptyDOMElement();
  });

  it("shows the distinct email-conflict message for google_email_conflict (R6)", () => {
    mockError = "google_email_conflict";
    render(<GoogleCallbackError />);

    expect(screen.getByRole("alert")).toHaveTextContent(/sudah terdaftar dengan password/i);
  });

  it.each(["state_mismatch", "nonce_mismatch", "google_token_invalid", "google_unavailable"])(
    "shows the shared generic retry message for %s, not the email-conflict copy (R6)",
    (code) => {
      mockError = code;
      render(<GoogleCallbackError />);

      const banner = screen.getByRole("alert");
      expect(banner).toHaveTextContent("Gagal masuk dengan Google. Silakan coba lagi.");
      expect(banner).not.toHaveTextContent(/sudah terdaftar dengan password/i);
    }
  );

  it("falls back to the same generic message for an unrecognized code, never rendering the raw code (R6)", () => {
    mockError = "some_future_unmapped_code";
    render(<GoogleCallbackError />);

    const banner = screen.getByRole("alert");
    expect(banner).toHaveTextContent("Gagal masuk dengan Google. Silakan coba lagi.");
    expect(banner).not.toHaveTextContent("some_future_unmapped_code");
  });

  it("renders as an alert (role) and moves focus into it on render (R7)", async () => {
    mockError = "google_unavailable";
    render(<GoogleCallbackError />);

    const banner = screen.getByRole("alert");
    await waitFor(() => expect(banner.closest("[tabindex]")).toHaveFocus());
  });
});
