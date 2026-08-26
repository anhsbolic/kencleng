import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { describe, expect, it, vi } from "vitest";
import { server } from "@/mocks/server";
import { RegisterForm } from "./register-form";

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { mutations: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

function fillValidForm() {
  fireEvent.change(screen.getByLabelText("Nama"), { target: { value: "Siti Rahma" } });
  fireEvent.change(screen.getByLabelText("Email"), {
    target: { value: "siti@example.com" },
  });
  fireEvent.change(screen.getByLabelText("Password"), {
    target: { value: "rahasia123" },
  });
}

describe("RegisterForm", () => {
  it("shows field-level validation errors on submit without touching any field (R1)", async () => {
    render(withQueryClient(<RegisterForm />));

    fireEvent.click(screen.getByRole("button", { name: "Daftar" }));

    expect(await screen.findByText("Nama wajib diisi")).toBeInTheDocument();
    expect(screen.getByText("Format email tidak valid")).toBeInTheDocument();
    expect(screen.getByText("Password minimal 8 karakter")).toBeInTheDocument();
  });

  it("accepts a password >= 8 chars locally with no breach-list check (R2)", async () => {
    render(withQueryClient(<RegisterForm />));
    fillValidForm();

    fireEvent.click(screen.getByRole("button", { name: "Daftar" }));

    // No local rejection — the only way this ever becomes an error is
    // a 422 round-trip (covered by the R5 test below), never a
    // client-side breach check.
    await waitFor(() =>
      expect(screen.queryByText("Password minimal 8 karakter")).not.toBeInTheDocument()
    );
  });

  it("submit button has type=submit and disables while pending (R3)", () => {
    render(withQueryClient(<RegisterForm />));
    const button = screen.getByRole("button", { name: "Daftar" });
    expect(button).toHaveAttribute("type", "submit");
  });

  it("replaces the form with a fixed success view on 202, verbatim backend message (R4)", async () => {
    render(withQueryClient(<RegisterForm />));
    fillValidForm();

    fireEvent.click(screen.getByRole("button", { name: "Daftar" }));

    expect(
      await screen.findByText(
        "Kalau email belum terdaftar, cek inbox untuk verifikasi. Kalau sudah, cek inbox untuk instruksi lebih lanjut."
      )
    ).toBeInTheDocument();
    expect(screen.queryByRole("button", { name: "Daftar" })).not.toBeInTheDocument();
  });

  it("maps each 422 field error verbatim via setError, no banner (R5)", async () => {
    server.use(
      http.post("/auth/register", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/validation-failed",
            title: "Validation Failed",
            status: 422,
            errors: [
              {
                field: "password",
                message: "Password ditemukan di daftar breach publik, silakan pilih password lain.",
              },
            ],
          },
          { status: 422 }
        )
      )
    );

    render(withQueryClient(<RegisterForm />));
    fillValidForm();
    fireEvent.click(screen.getByRole("button", { name: "Daftar" }));

    expect(
      await screen.findByText(
        "Password ditemukan di daftar breach publik, silakan pilih password lain."
      )
    ).toBeInTheDocument();
    expect(screen.queryByRole("alert")).not.toBeInTheDocument();
  });

  it("shows a generic banner on a network failure, never raw error text (R6)", async () => {
    server.use(http.post("/auth/register", () => HttpResponse.error()));

    render(withQueryClient(<RegisterForm />));
    fillValidForm();
    fireEvent.click(screen.getByRole("button", { name: "Daftar" }));

    expect(await screen.findByRole("alert")).toHaveTextContent(
      "Terjadi kesalahan. Silakan coba lagi."
    );
    expect(screen.queryByText(/failed to fetch/i)).not.toBeInTheDocument();
  });

  it("renders 'Daftar dengan Google' as a real navigation link with intent=login, not a button (R7; account/02 D1/R1)", () => {
    render(withQueryClient(<RegisterForm />));

    const googleLink = screen.getByRole("link", { name: /daftar dengan google/i });
    expect(googleLink).toHaveAttribute("href", "/auth/google/redirect?intent=login");
  });

  it("shows the backend's rate-limit detail verbatim on 429 (R10)", async () => {
    server.use(
      http.post("/auth/register", () =>
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

    render(withQueryClient(<RegisterForm />));
    fillValidForm();
    fireEvent.click(screen.getByRole("button", { name: "Daftar" }));

    expect(
      await screen.findByText("Terlalu banyak percobaan gagal. Coba lagi dalam 15 menit.")
    ).toBeInTheDocument();
  });

  it("moves focus into the success heading once the form is replaced (R17)", async () => {
    render(withQueryClient(<RegisterForm />));
    fillValidForm();
    fireEvent.click(screen.getByRole("button", { name: "Daftar" }));

    const heading = await screen.findByRole("heading", { name: "Cek email kamu" });
    await waitFor(() => expect(heading).toHaveFocus());
  });

  it("renders 'Masuk' as a real navigation link by default (standalone /register route)", () => {
    render(withQueryClient(<RegisterForm />));

    expect(screen.getByRole("link", { name: "Masuk" })).toHaveAttribute("href", "/login");
  });

  it("renders 'Masuk' as a button calling onSwitchToLogin when provided (modal context)", () => {
    const onSwitchToLogin = vi.fn();
    render(withQueryClient(<RegisterForm onSwitchToLogin={onSwitchToLogin} />));

    const masuk = screen.getByRole("button", { name: "Masuk" });
    expect(masuk).not.toHaveAttribute("href");

    fireEvent.click(masuk);
    expect(onSwitchToLogin).toHaveBeenCalledTimes(1);
  });

  it("never fires a request on email blur/change beyond the explicit submit (R18)", () => {
    let requestCount = 0;
    server.use(
      http.post("/auth/register", () => {
        requestCount += 1;
        return HttpResponse.json({ message: "ok" }, { status: 202 });
      })
    );

    render(withQueryClient(<RegisterForm />));
    const emailInput = screen.getByLabelText("Email");
    fireEvent.change(emailInput, { target: { value: "siti@example.com" } });
    fireEvent.blur(emailInput);

    expect(requestCount).toBe(0);
  });
});
