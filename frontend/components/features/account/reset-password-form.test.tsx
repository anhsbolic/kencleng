import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { describe, expect, it, vi } from "vitest";
import { server } from "@/mocks/server";
import { ResetPasswordForm } from "./reset-password-form";

let mockToken: string | null = "valid-token";

vi.mock("next/navigation", () => ({
  useSearchParams: () => ({
    get: (key: string) => (key === "token" ? mockToken : null),
  }),
}));

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { mutations: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

function submitNewPassword(password = "rahasia123") {
  fireEvent.change(screen.getByLabelText("Password baru"), { target: { value: password } });
  fireEvent.click(screen.getByRole("button", { name: "Reset password" }));
}

describe("ResetPasswordForm", () => {
  it("shows the generic invalid-link banner and renders no form when the token is missing (R7)", () => {
    mockToken = null;
    render(withQueryClient(<ResetPasswordForm />));

    expect(
      screen.getByText("Link reset password tidak valid atau sudah digunakan.")
    ).toBeInTheDocument();
    expect(screen.queryByLabelText("Password baru")).not.toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Minta link baru" })).toHaveAttribute(
      "href",
      "/forgot-password"
    );
  });

  it("renders the new-password field and submit button when a token is present (R8)", () => {
    mockToken = "valid-token";
    render(withQueryClient(<ResetPasswordForm />));

    expect(screen.getByLabelText("Password baru")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Reset password" })).toBeInTheDocument();
  });

  it("blocks submission of fewer than 8 characters client-side (R9)", async () => {
    mockToken = "valid-token";
    render(withQueryClient(<ResetPasswordForm />));

    submitNewPassword("short");

    expect(await screen.findByText("Password minimal 8 karakter")).toBeInTheDocument();
  });

  it("swaps to an inline success view on 200, with a link to /login, no navigation call (R10)", async () => {
    mockToken = "valid-token";
    render(withQueryClient(<ResetPasswordForm />));

    submitNewPassword();

    expect(
      await screen.findByText("Password berhasil diubah. Silakan login ulang.")
    ).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Masuk sekarang" })).toHaveAttribute(
      "href",
      "/login"
    );
    expect(screen.queryByRole("button", { name: "Reset password" })).not.toBeInTheDocument();
  });

  it("shows the same invalid-link banner as the missing-token case on 404, removes the form (R11)", async () => {
    mockToken = "used-token";
    server.use(
      http.post("/auth/reset-password", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/problems/token-not-found",
            title: "Token Not Found",
            status: 404,
            detail: "The verification token was not found.",
          },
          { status: 404 }
        )
      )
    );

    render(withQueryClient(<ResetPasswordForm />));
    submitNewPassword();

    expect(
      await screen.findByText("Link reset password tidak valid atau sudah digunakan.")
    ).toBeInTheDocument();
    expect(screen.queryByText("The verification token was not found.")).not.toBeInTheDocument();
    expect(screen.queryByLabelText("Password baru")).not.toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Minta link baru" })).toBeInTheDocument();
  });

  it("shows a distinct expired-link banner on 410, removes the form (R12)", async () => {
    mockToken = "expired-token";
    server.use(
      http.post("/auth/reset-password", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/problems/token-expired",
            title: "Token Expired",
            status: 410,
            detail: "The verification token has expired.",
          },
          { status: 410 }
        )
      )
    );

    render(withQueryClient(<ResetPasswordForm />));
    submitNewPassword();

    expect(
      await screen.findByText("Link reset password sudah kedaluwarsa. Silakan minta link baru.")
    ).toBeInTheDocument();
    expect(screen.queryByText("The verification token has expired.")).not.toBeInTheDocument();
    expect(screen.queryByLabelText("Password baru")).not.toBeInTheDocument();
  });

  it("shows the frontend-owned weak-password banner on 422, form stays mounted and interactive, token unchanged on retry (R13)", async () => {
    mockToken = "valid-token";
    const receivedTokens: string[] = [];
    server.use(
      http.post("/auth/reset-password", async ({ request }) => {
        const body = (await request.json()) as { token: string };
        receivedTokens.push(body.token);
        return HttpResponse.json(
          {
            type: "https://kencleng.dev/problems/validation-failed",
            title: "Validation Failed",
            status: 422,
            detail: "The request was invalid.",
          },
          { status: 422 }
        );
      })
    );

    render(withQueryClient(<ResetPasswordForm />));
    submitNewPassword("password");

    expect(
      await screen.findByText(
        "Password tidak memenuhi syarat. Gunakan minimal 8 karakter dan hindari password yang umum digunakan atau pernah bocor."
      )
    ).toBeInTheDocument();
    expect(screen.queryByText("The request was invalid.")).not.toBeInTheDocument();
    // Form stays mounted/interactive — a resubmit is possible.
    expect(screen.getByLabelText("Password baru")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Reset password" })).toBeInTheDocument();

    submitNewPassword("password2");
    await waitFor(() => expect(receivedTokens).toEqual(["valid-token", "valid-token"]));
  });

  it("shows the same rate-limited banner as forgot-password's on 429, form stays interactive (R14)", async () => {
    mockToken = "valid-token";
    server.use(
      http.post("/auth/reset-password", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/problems/rate-limited",
            title: "Rate Limited",
            status: 429,
            detail: "Too many requests. Try again later.",
          },
          { status: 429 }
        )
      )
    );

    render(withQueryClient(<ResetPasswordForm />));
    submitNewPassword();

    expect(
      await screen.findByText("Terlalu banyak percobaan. Coba lagi beberapa saat lagi.")
    ).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Reset password" })).toBeInTheDocument();
  });

  it("shows a generic banner on a network failure, form stays interactive (R15)", async () => {
    mockToken = "valid-token";
    server.use(http.post("/auth/reset-password", () => HttpResponse.error()));

    render(withQueryClient(<ResetPasswordForm />));
    submitNewPassword();

    expect(await screen.findByText("Terjadi kesalahan. Silakan coba lagi.")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Reset password" })).toBeInTheDocument();
  });
});
