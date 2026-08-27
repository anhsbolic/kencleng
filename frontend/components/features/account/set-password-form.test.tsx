import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { beforeEach, describe, expect, it } from "vitest";
import { server } from "@/mocks/server";
import { useAuthStore } from "@/lib/stores/auth-store";
import { SetPasswordForm } from "./set-password-form";

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { mutations: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

beforeEach(() => {
  useAuthStore.setState({ accessToken: null });
});

describe("SetPasswordForm — mode=add", () => {
  it("renders email + password fields (R4)", () => {
    render(withQueryClient(<SetPasswordForm mode="add" />));

    expect(screen.getByLabelText("Email")).toBeInTheDocument();
    expect(screen.getByLabelText("Password baru")).toBeInTheDocument();
  });

  it("blocks submission of an invalid email or a short password client-side (R5)", async () => {
    render(withQueryClient(<SetPasswordForm mode="add" />));

    fireEvent.change(screen.getByLabelText("Email"), { target: { value: "not-an-email" } });
    fireEvent.change(screen.getByLabelText("Password baru"), { target: { value: "short" } });
    fireEvent.click(screen.getByRole("button", { name: "Atur Password" }));

    expect(await screen.findByText("Format email tidak valid")).toBeInTheDocument();
    expect(screen.getByText("Password minimal 8 karakter")).toBeInTheDocument();
  });

  it("shows the backend's generic success message verbatim on 202, form replaced (R6)", async () => {
    render(withQueryClient(<SetPasswordForm mode="add" />));

    fireEvent.change(screen.getByLabelText("Email"), { target: { value: "work@company.com" } });
    fireEvent.change(screen.getByLabelText("Password baru"), {
      target: { value: "strong-pw-123" },
    });
    fireEvent.click(screen.getByRole("button", { name: "Atur Password" }));

    expect(
      await screen.findByText("Kalau email tersedia, cek inbox untuk verifikasi.")
    ).toBeInTheDocument();
    expect(screen.queryByLabelText("Email")).not.toBeInTheDocument();
  });

  it("attaches the frontend-owned weak-password message to the password field on 422, never the backend's raw text (R7)", async () => {
    server.use(
      http.post("/account/security/set-password", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/validation-failed",
            title: "Validation Failed",
            status: 422,
            errors: [{ field: "password", message: "password is not allowed" }],
          },
          { status: 422 }
        )
      )
    );

    render(withQueryClient(<SetPasswordForm mode="add" />));
    fireEvent.change(screen.getByLabelText("Email"), { target: { value: "work@company.com" } });
    fireEvent.change(screen.getByLabelText("Password baru"), { target: { value: "password" } });
    fireEvent.click(screen.getByRole("button", { name: "Atur Password" }));

    expect(
      await screen.findByText(
        "Password tidak memenuhi syarat. Gunakan minimal 8 karakter dan hindari password yang umum digunakan atau pernah bocor."
      )
    ).toBeInTheDocument();
    expect(screen.queryByText("password is not allowed")).not.toBeInTheDocument();
    expect(screen.getByLabelText("Email")).toBeInTheDocument();
  });

  it("shows a generic banner on a network failure, form stays interactive (R12)", async () => {
    server.use(http.post("/account/security/set-password", () => HttpResponse.error()));

    render(withQueryClient(<SetPasswordForm mode="add" />));
    fireEvent.change(screen.getByLabelText("Email"), { target: { value: "work@company.com" } });
    fireEvent.change(screen.getByLabelText("Password baru"), {
      target: { value: "strong-pw-123" },
    });
    fireEvent.click(screen.getByRole("button", { name: "Atur Password" }));

    expect(await screen.findByText("Terjadi kesalahan. Silakan coba lagi.")).toBeInTheDocument();
    expect(screen.getByLabelText("Email")).toBeInTheDocument();
  });
});

describe("SetPasswordForm — mode=change", () => {
  it("renders current_password + new password fields (R8)", () => {
    render(withQueryClient(<SetPasswordForm mode="change" />));

    expect(screen.getByLabelText("Password saat ini")).toBeInTheDocument();
    expect(screen.getByLabelText("Password baru")).toBeInTheDocument();
  });

  it("blocks submission of an empty current_password or a short new password client-side (R9)", async () => {
    render(withQueryClient(<SetPasswordForm mode="change" />));

    fireEvent.change(screen.getByLabelText("Password baru"), { target: { value: "short" } });
    fireEvent.click(screen.getByRole("button", { name: "Ganti Password" }));

    expect(await screen.findByText("Password saat ini wajib diisi")).toBeInTheDocument();
    expect(screen.getByText("Password minimal 8 karakter")).toBeInTheDocument();
  });

  it("on 200, clears the session and shows no error banner — session-cutover mechanics covered by use-set-password.test.ts (R10)", async () => {
    server.use(
      http.post("/account/security/set-password", () =>
        HttpResponse.json(
          { message: "Password berhasil diganti. Semua sesi lain telah keluar otomatis." },
          { status: 200 }
        )
      )
    );
    useAuthStore.setState({ accessToken: "still-valid-token" });
    render(withQueryClient(<SetPasswordForm mode="change" />));

    fireEvent.change(screen.getByLabelText("Password saat ini"), {
      target: { value: "old-pw-123" },
    });
    fireEvent.change(screen.getByLabelText("Password baru"), {
      target: { value: "new-strong-pw" },
    });
    fireEvent.click(screen.getByRole("button", { name: "Ganti Password" }));

    await waitFor(() => expect(useAuthStore.getState().accessToken).toBeNull());
    expect(screen.queryByRole("alert")).not.toBeInTheDocument();
  });

  it("shows the backend's detail verbatim on 401, form stays interactive (R11)", async () => {
    server.use(
      http.post("/account/security/set-password", () =>
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

    render(withQueryClient(<SetPasswordForm mode="change" />));
    fireEvent.change(screen.getByLabelText("Password saat ini"), { target: { value: "wrong-pw" } });
    fireEvent.change(screen.getByLabelText("Password baru"), {
      target: { value: "new-strong-pw" },
    });
    fireEvent.click(screen.getByRole("button", { name: "Ganti Password" }));

    expect(await screen.findByRole("alert")).toHaveTextContent("Email atau password salah.");
    expect(screen.getByLabelText("Password saat ini")).toBeInTheDocument();
  });

  it("attaches the frontend-owned weak-password message to the password field on 422 (R7)", async () => {
    server.use(
      http.post("/account/security/set-password", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/validation-failed",
            title: "Validation Failed",
            status: 422,
            errors: [{ field: "password", message: "password is not allowed" }],
          },
          { status: 422 }
        )
      )
    );

    render(withQueryClient(<SetPasswordForm mode="change" />));
    fireEvent.change(screen.getByLabelText("Password saat ini"), {
      target: { value: "old-pw-123" },
    });
    fireEvent.change(screen.getByLabelText("Password baru"), { target: { value: "password" } });
    fireEvent.click(screen.getByRole("button", { name: "Ganti Password" }));

    expect(
      await screen.findByText(
        "Password tidak memenuhi syarat. Gunakan minimal 8 karakter dan hindari password yang umum digunakan atau pernah bocor."
      )
    ).toBeInTheDocument();
  });

  it("shows a generic banner on a network failure, form stays interactive (R12)", async () => {
    server.use(http.post("/account/security/set-password", () => HttpResponse.error()));

    render(withQueryClient(<SetPasswordForm mode="change" />));
    fireEvent.change(screen.getByLabelText("Password saat ini"), {
      target: { value: "old-pw-123" },
    });
    fireEvent.change(screen.getByLabelText("Password baru"), {
      target: { value: "new-strong-pw" },
    });
    fireEvent.click(screen.getByRole("button", { name: "Ganti Password" }));

    expect(await screen.findByText("Terjadi kesalahan. Silakan coba lagi.")).toBeInTheDocument();
    expect(screen.getByLabelText("Password saat ini")).toBeInTheDocument();
  });
});
