import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { describe, expect, it } from "vitest";
import { server } from "@/mocks/server";
import { ResendVerificationControl } from "./resend-verification-control";

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { mutations: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

describe("ResendVerificationControl", () => {
  it("pre-fills the email field from defaultEmail and calls resend with it (R8)", async () => {
    let capturedEmail: string | undefined;
    server.use(
      http.post("/auth/verify-email/resend", async ({ request }) => {
        const body = (await request.json()) as { email: string };
        capturedEmail = body.email;
        return HttpResponse.json({ message: "Kalau email terdaftar, instruksi sudah dikirim." }, { status: 202 });
      })
    );

    render(withQueryClient(<ResendVerificationControl defaultEmail="siti@example.com" />));

    expect(screen.getByLabelText("Email")).toHaveValue("siti@example.com");
    screen.getByRole("button", { name: /kirim ulang/i }).click();

    await waitFor(() => expect(capturedEmail).toBe("siti@example.com"));
  });

  it("shows the same generic confirmation on 202 regardless of match (R9)", async () => {
    render(withQueryClient(<ResendVerificationControl defaultEmail="anyone@example.com" />));

    screen.getByRole("button", { name: /kirim ulang/i }).click();

    expect(
      await screen.findByText("Kalau email terdaftar, instruksi sudah dikirim.")
    ).toBeInTheDocument();
  });

  it("shows the backend's rate-limit detail verbatim on 429 (R10)", async () => {
    server.use(
      http.post("/auth/verify-email/resend", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/too-many-requests",
            title: "Too Many Requests",
            status: 429,
            detail: "Terlalu banyak percobaan gagal. Coba lagi dalam 15 menit.",
          },
          { status: 429 }
        )
      )
    );

    render(withQueryClient(<ResendVerificationControl defaultEmail="anyone@example.com" />));
    screen.getByRole("button", { name: /kirim ulang/i }).click();

    expect(
      await screen.findByText("Terlalu banyak percobaan gagal. Coba lagi dalam 15 menit.")
    ).toBeInTheDocument();
  });

  it("disables the resend button when the email field is empty", () => {
    render(withQueryClient(<ResendVerificationControl />));
    expect(screen.getByRole("button", { name: /kirim ulang/i })).toBeDisabled();
  });
});
