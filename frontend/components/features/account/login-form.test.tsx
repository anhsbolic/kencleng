import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { server } from "@/mocks/server";
import { useAuthStore } from "@/lib/stores/auth-store";
import { LoginForm } from "./login-form";

const pushMock = vi.fn();
vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: pushMock }),
}));

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { mutations: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

function fillPasswordStep() {
  fireEvent.change(screen.getByLabelText("Email"), {
    target: { value: "siti@example.com" },
  });
  fireEvent.change(screen.getByLabelText("Password"), {
    target: { value: "rahasia123" },
  });
}

beforeEach(() => {
  useAuthStore.setState({ accessToken: null });
  pushMock.mockClear();
});

describe("LoginForm — password step", () => {
  it("renders email + password fields, 'Lupa password?' link, submit, divider + Google button (R1)", () => {
    render(withQueryClient(<LoginForm />));

    expect(screen.getByLabelText("Email")).toBeInTheDocument();
    expect(screen.getByLabelText("Password")).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Lupa password?" })).toHaveAttribute(
      "href",
      "/forgot-password"
    );
    expect(screen.getByRole("button", { name: "Masuk" })).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Masuk dengan Google" })).toHaveAttribute(
      "href",
      "/auth/google/redirect?intent=login"
    );
  });

  it("renders 'Daftar' as a real navigation link by default (standalone /login route)", () => {
    render(withQueryClient(<LoginForm />));

    expect(screen.getByRole("link", { name: "Daftar" })).toHaveAttribute("href", "/register");
  });

  it("renders 'Daftar' as a button calling onSwitchToRegister when provided (modal context)", () => {
    const onSwitchToRegister = vi.fn();
    render(withQueryClient(<LoginForm onSwitchToRegister={onSwitchToRegister} />));

    const daftar = screen.getByRole("button", { name: "Daftar" });
    expect(daftar).not.toHaveAttribute("href");

    fireEvent.click(daftar);
    expect(onSwitchToRegister).toHaveBeenCalledTimes(1);
  });

  it("toggles the password field's type and accessible label (R2)", () => {
    render(withQueryClient(<LoginForm />));

    const passwordInput = screen.getByLabelText("Password");
    expect(passwordInput).toHaveAttribute("type", "password");

    fireEvent.click(screen.getByRole("button", { name: "Tampilkan password" }));
    expect(passwordInput).toHaveAttribute("type", "text");
    expect(screen.getByRole("button", { name: "Sembunyikan password" })).toBeInTheDocument();

    fireEvent.click(screen.getByRole("button", { name: "Sembunyikan password" }));
    expect(passwordInput).toHaveAttribute("type", "password");
  });

  it("on a 200 status=ok response, redirects to /dashboard/profile with no residual banner (R3)", async () => {
    render(withQueryClient(<LoginForm />));
    fillPasswordStep();

    fireEvent.click(screen.getByRole("button", { name: "Masuk" }));

    await waitFor(() => expect(pushMock).toHaveBeenCalledWith("/dashboard/profile"));
    expect(screen.queryByRole("alert")).not.toBeInTheDocument();
  });

  it("on a 200 status=mfa_required response, transitions to the MFA step (R4)", async () => {
    server.use(
      http.post("/auth/login", () =>
        HttpResponse.json({ status: "mfa_required", mfa_pending_token: "pending-token" })
      )
    );

    render(withQueryClient(<LoginForm />));
    fillPasswordStep();
    fireEvent.click(screen.getByRole("button", { name: "Masuk" }));

    expect(await screen.findByRole("heading", { name: "Verifikasi dua langkah" })).toBeInTheDocument();
    expect(screen.getByLabelText("Kode OTP")).toBeInTheDocument();
    expect(pushMock).not.toHaveBeenCalled();
  });

  it("shows the identical generic banner for both 401 and 429, never as a field-level error (R5)", async () => {
    server.use(
      http.post("/auth/login", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/invalid-credentials",
            title: "Invalid Credentials",
            status: 401,
            detail: "Email atau password salah.",
          },
          { status: 401 }
        )
      )
    );

    render(withQueryClient(<LoginForm />));
    fillPasswordStep();
    fireEvent.click(screen.getByRole("button", { name: "Masuk" }));

    expect(await screen.findByRole("alert")).toHaveTextContent("Email atau password salah.");
    // Never attached to the email input's own error prop (Known Issue #1).
    expect(screen.getByLabelText("Email")).not.toHaveAttribute("aria-invalid", "true");

    server.use(
      http.post("/auth/login", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/too-many-requests",
            title: "Too Many Requests",
            status: 429,
            detail: "Email atau password salah.",
          },
          { status: 429 }
        )
      )
    );

    fireEvent.click(screen.getByRole("button", { name: "Masuk" }));
    expect(await screen.findByRole("alert")).toHaveTextContent("Email atau password salah.");
  });
});

describe("LoginForm — MFA step", () => {
  async function enterMfaStep() {
    server.use(
      http.post("/auth/login", () =>
        HttpResponse.json({ status: "mfa_required", mfa_pending_token: "pending-token" })
      )
    );
    render(withQueryClient(<LoginForm />));
    fillPasswordStep();
    fireEvent.click(screen.getByRole("button", { name: "Masuk" }));
    await screen.findByRole("heading", { name: "Verifikasi dua langkah" });
  }

  it("rejects submission with neither totp_code nor backup_code filled (R6)", async () => {
    await enterMfaStep();

    fireEvent.click(screen.getByRole("button", { name: "Verifikasi" }));

    expect(await screen.findByText("Masukkan kode OTP atau kode cadangan")).toBeInTheDocument();
  });

  it("switches to the backup-code field via the toggle (R6)", async () => {
    await enterMfaStep();

    fireEvent.click(screen.getByRole("button", { name: "Gunakan kode cadangan" }));

    expect(screen.getByLabelText("Kode cadangan")).toBeInTheDocument();
    expect(screen.queryByLabelText("Kode OTP")).not.toBeInTheDocument();
  });

  it("on a 200 response, redirects with no error banner (R7)", async () => {
    await enterMfaStep();

    fireEvent.change(screen.getByLabelText("Kode OTP"), { target: { value: "123456" } });
    fireEvent.click(screen.getByRole("button", { name: "Verifikasi" }));

    await waitFor(() => expect(pushMock).toHaveBeenCalledWith("/dashboard/profile"));
    expect(screen.queryByRole("alert")).not.toBeInTheDocument();
  });

  it("on a 401/429, shows the banner and stays on the MFA step (R8)", async () => {
    await enterMfaStep();
    server.use(
      http.post("/auth/login/mfa", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/invalid-credentials",
            title: "Invalid Credentials",
            status: 401,
            detail: "Email atau password salah.",
          },
          { status: 401 }
        )
      )
    );

    fireEvent.change(screen.getByLabelText("Kode OTP"), { target: { value: "000000" } });
    fireEvent.click(screen.getByRole("button", { name: "Verifikasi" }));

    expect(await screen.findByRole("alert")).toHaveTextContent("Email atau password salah.");
    expect(screen.getByRole("heading", { name: "Verifikasi dua langkah" })).toBeInTheDocument();
    expect(pushMock).not.toHaveBeenCalled();
  });

  it("does not render the Google button while on the MFA step (R9)", async () => {
    await enterMfaStep();

    expect(screen.queryByRole("link", { name: "Masuk dengan Google" })).not.toBeInTheDocument();
  });

  it("remounting always starts at the password step, regardless of prior state (R10)", async () => {
    const { unmount } = render(withQueryClient(<LoginForm />));
    server.use(
      http.post("/auth/login", () =>
        HttpResponse.json({ status: "mfa_required", mfa_pending_token: "pending-token" })
      )
    );
    fillPasswordStep();
    fireEvent.click(screen.getByRole("button", { name: "Masuk" }));
    await screen.findByRole("heading", { name: "Verifikasi dua langkah" });
    unmount();

    render(withQueryClient(<LoginForm />));

    expect(screen.getByLabelText("Email")).toBeInTheDocument();
    expect(screen.queryByRole("heading", { name: "Verifikasi dua langkah" })).not.toBeInTheDocument();
  });

  it("'Kembali ke halaman login' returns to the password step", async () => {
    await enterMfaStep();

    fireEvent.click(screen.getByRole("button", { name: "Kembali ke halaman login" }));

    expect(screen.getByLabelText("Email")).toBeInTheDocument();
  });
});
