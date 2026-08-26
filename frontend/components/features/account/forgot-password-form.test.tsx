import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { describe, expect, it } from "vitest";
import { server } from "@/mocks/server";
import { ForgotPasswordForm } from "./forgot-password-form";

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { mutations: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

function submit() {
  fireEvent.click(screen.getByRole("button", { name: "Kirim link reset" }));
}

describe("ForgotPasswordForm", () => {
  it("renders an email field, submit button, and a link back to /login, no Google button (R1)", () => {
    render(withQueryClient(<ForgotPasswordForm />));

    expect(screen.getByLabelText("Email")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Kirim link reset" })).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Kembali ke halaman login" })).toHaveAttribute(
      "href",
      "/login"
    );
    expect(screen.queryByRole("button", { name: /google/i })).not.toBeInTheDocument();
  });

  it("blocks submission of an empty/malformed email client-side (R2)", async () => {
    render(withQueryClient(<ForgotPasswordForm />));

    fireEvent.change(screen.getByLabelText("Email"), { target: { value: "not-an-email" } });
    submit();

    expect(await screen.findByText("Format email tidak valid")).toBeInTheDocument();
  });

  it("swaps to an inline success view on 202, identical regardless of which email was used (R3)", async () => {
    render(withQueryClient(<ForgotPasswordForm />));

    fireEvent.change(screen.getByLabelText("Email"), {
      target: { value: "unknown@example.com" },
    });
    submit();

    expect(
      await screen.findByText("Kalau email terdaftar, instruksi sudah dikirim.")
    ).toBeInTheDocument();
    expect(screen.getByRole("heading", { name: "Cek email kamu" })).toHaveFocus();
    expect(screen.queryByRole("button", { name: "Kirim link reset" })).not.toBeInTheDocument();
  });

  it("maps a 422 (malformed email) to a frontend-owned field error, never the backend's literal text, no banner (R4)", async () => {
    server.use(
      http.post("/auth/forgot-password", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/validation-failed",
            title: "Validation Failed",
            status: 422,
            errors: [{ field: "email", message: "must be a valid email" }],
          },
          { status: 422 }
        )
      )
    );

    render(withQueryClient(<ForgotPasswordForm />));
    fireEvent.change(screen.getByLabelText("Email"), { target: { value: "weird@x" } });
    submit();

    expect(await screen.findByText("Format email tidak valid")).toBeInTheDocument();
    expect(screen.queryByText("must be a valid email")).not.toBeInTheDocument();
    expect(screen.queryByRole("alert")).not.toBeInTheDocument();
  });

  it("shows a frontend-owned rate-limited banner on 429, never the backend's raw detail (R5)", async () => {
    server.use(
      http.post("/auth/forgot-password", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/too-many-requests",
            title: "Too Many Requests",
            status: 429,
            detail: "Too many requests. Try again later.",
          },
          { status: 429 }
        )
      )
    );

    render(withQueryClient(<ForgotPasswordForm />));
    fireEvent.change(screen.getByLabelText("Email"), { target: { value: "a@example.com" } });
    submit();

    expect(
      await screen.findByText("Terlalu banyak percobaan. Coba lagi beberapa saat lagi.")
    ).toBeInTheDocument();
    expect(screen.queryByText("Too many requests. Try again later.")).not.toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Kirim link reset" })).toBeInTheDocument();
  });

  it("shows a generic banner on a network failure, never raw error text, form stays visible (R6)", async () => {
    server.use(http.post("/auth/forgot-password", () => HttpResponse.error()));

    render(withQueryClient(<ForgotPasswordForm />));
    fireEvent.change(screen.getByLabelText("Email"), { target: { value: "a@example.com" } });
    submit();

    expect(await screen.findByText("Terjadi kesalahan. Silakan coba lagi.")).toBeInTheDocument();
    await waitFor(() =>
      expect(screen.getByRole("button", { name: "Kirim link reset" })).toBeInTheDocument()
    );
  });
});
