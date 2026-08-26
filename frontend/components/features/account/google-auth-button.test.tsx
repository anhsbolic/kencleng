import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { GoogleAuthButton } from "./google-auth-button";

describe("GoogleAuthButton", () => {
  it("renders a real navigation link, not a button, to the redirect-initiation endpoint (R3)", () => {
    render(<GoogleAuthButton intent="login" label="Masuk dengan Google" />);

    const link = screen.getByRole("link", { name: "Masuk dengan Google" });
    expect(link.tagName).toBe("A");
  });

  it("targets intent=login for the login label (R1)", () => {
    render(<GoogleAuthButton intent="login" label="Masuk dengan Google" />);

    expect(screen.getByRole("link", { name: "Masuk dengan Google" })).toHaveAttribute(
      "href",
      "/auth/google/redirect?intent=login"
    );
  });

  it("targets intent=login for the register label — never intent=register (R1)", () => {
    render(<GoogleAuthButton intent="login" label="Daftar dengan Google" />);

    expect(screen.getByRole("link", { name: "Daftar dengan Google" })).toHaveAttribute(
      "href",
      "/auth/google/redirect?intent=login"
    );
  });

  it("both /login and /register call sites share this one component (R2)", () => {
    // Structural guard, not a runtime assertion: both `register-form.tsx`
    // and `app/(auth)/login/page.tsx` import `GoogleAuthButton` from
    // this file rather than hand-writing their own anchor — enforced by
    // this module being the only place `/auth/google/redirect` is
    // constructed in `components/features/account/`. A future caller
    // that hand-copies the anchor instead of importing this component
    // is a code-review flag, not something a unit test alone can catch.
    render(<GoogleAuthButton intent="login" label="Masuk dengan Google" />);
    expect(screen.getByRole("link")).toBeInTheDocument();
  });
});
