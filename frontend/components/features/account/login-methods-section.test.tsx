import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, screen } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { describe, expect, it } from "vitest";
import { server } from "@/mocks/server";
import { LoginMethodsSection } from "./login-methods-section";

function mockMe(authProviders: ("email_password" | "google")[], emailVerified: boolean) {
  server.use(
    http.get("/account/me", () =>
      HttpResponse.json({
        id: "u1",
        name: "Test User",
        email: "test@example.com",
        email_verified: emailVerified,
        roles: [],
        auth_providers: authProviders,
        mfa_enabled: false,
        created_at: new Date().toISOString(),
      })
    )
  );
}

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false }, mutations: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

describe("LoginMethodsSection", () => {
  it("shows a skeleton before the account query resolves (R3)", () => {
    const { container } = render(withQueryClient(<LoginMethodsSection />));

    expect(container.querySelector(".animate-pulse")).toBeInTheDocument();
    expect(screen.queryByText("Metode Masuk")).not.toBeInTheDocument();
  });

  it("Google-only, no email_password: renders the add-password form and no pending-verification banner (R1)", async () => {
    mockMe(["google"], false);
    render(withQueryClient(<LoginMethodsSection />));

    expect(await screen.findByLabelText("Email")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Atur Password" })).toBeInTheDocument();
    expect(
      screen.queryByText("Menunggu verifikasi email kamu — cek inbox untuk menyelesaikan.")
    ).not.toBeInTheDocument();
  });

  it("email_password present but unverified: renders the change-password form AND the pending-verification banner, form not hidden (R1/R2)", async () => {
    mockMe(["google", "email_password"], false);
    render(withQueryClient(<LoginMethodsSection />));

    expect(
      await screen.findByText("Menunggu verifikasi email kamu — cek inbox untuk menyelesaikan.")
    ).toBeInTheDocument();
    expect(screen.getByLabelText("Password saat ini")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Ganti Password" })).toBeInTheDocument();
  });

  it("email_password verified, no Google: renders the change-password form and the link-to-Google trigger", async () => {
    mockMe(["email_password"], true);
    render(withQueryClient(<LoginMethodsSection />));

    expect(await screen.findByLabelText("Password saat ini")).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Hubungkan ke Google" })).toBeInTheDocument();
    expect(
      screen.queryByText("Menunggu verifikasi email kamu — cek inbox untuk menyelesaikan.")
    ).not.toBeInTheDocument();
  });

  it("email_password verified AND Google present: renders the unlink form", async () => {
    mockMe(["email_password", "google"], true);
    render(withQueryClient(<LoginMethodsSection />));

    expect(await screen.findByLabelText("Password saat ini")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Lepas Tautan Google" })).toBeInTheDocument();
  });
});
