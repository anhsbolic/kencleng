import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { describe, expect, it } from "vitest";
import { server } from "@/mocks/server";
import { MfaDisableForm } from "./mfa-disable-form";

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false }, mutations: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

describe("MfaDisableForm — email_password branch (R14-R16)", () => {
  it("shows a password field and a destructive submit button (R14)", () => {
    render(withQueryClient(<MfaDisableForm hasEmailPassword />));

    expect(screen.getByLabelText("Konfirmasi password")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Nonaktifkan MFA" })).toBeInTheDocument();
  });

  it("shows the disable→re-enroll explanatory line, no regenerate action (R20)", () => {
    render(withQueryClient(<MfaDisableForm hasEmailPassword />));

    expect(
      screen.getByText("Untuk mendapatkan kode cadangan baru, nonaktifkan MFA lalu aktifkan kembali.")
    ).toBeInTheDocument();
    expect(screen.queryByRole("button", { name: /regenerasi|regenerate/i })).not.toBeInTheDocument();
  });

  it("resolves 200 with no local success banner (R15)", async () => {
    render(withQueryClient(<MfaDisableForm hasEmailPassword />));

    fireEvent.change(screen.getByLabelText("Konfirmasi password"), {
      target: { value: "correct-pw" },
    });
    fireEvent.click(screen.getByRole("button", { name: "Nonaktifkan MFA" }));

    await waitFor(() => expect(screen.queryByRole("alert")).not.toBeInTheDocument());
  });

  it("shows the backend's detail verbatim on 401, form stays interactive (R16), focus moves into the banner (R21)", async () => {
    server.use(
      http.post("/account/security/mfa/disable", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/unauthorized",
            title: "Unauthorized",
            status: 401,
            detail: "Password salah.",
          },
          { status: 401 }
        )
      )
    );

    render(withQueryClient(<MfaDisableForm hasEmailPassword />));
    fireEvent.change(screen.getByLabelText("Konfirmasi password"), {
      target: { value: "wrong-pw" },
    });
    fireEvent.click(screen.getByRole("button", { name: "Nonaktifkan MFA" }));

    const banner = await screen.findByRole("alert");
    expect(banner).toHaveTextContent("Password salah.");
    await waitFor(() => expect(document.activeElement).toBe(banner.parentElement));
    expect(screen.getByLabelText("Konfirmasi password")).toBeInTheDocument();
  });
});

describe("MfaDisableForm — Google-only branch (R17-R19)", () => {
  it("shows a single button, no password field (R17)", () => {
    render(withQueryClient(<MfaDisableForm hasEmailPassword={false} />));

    expect(screen.queryByLabelText("Konfirmasi password")).not.toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Nonaktifkan MFA" })).toBeInTheDocument();
  });

  it("resolves 200 with no local success banner (R18)", async () => {
    render(withQueryClient(<MfaDisableForm hasEmailPassword={false} />));

    fireEvent.click(screen.getByRole("button", { name: "Nonaktifkan MFA" }));

    await waitFor(() => expect(screen.queryByRole("alert")).not.toBeInTheDocument());
  });

  it("shows an error banner + reauth link on 401, button stays available to retry (R19)", async () => {
    server.use(
      http.post("/account/security/mfa/disable", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/unauthorized",
            title: "Unauthorized",
            status: 401,
            detail: "Access token tidak valid atau sudah kedaluwarsa.",
          },
          { status: 401 }
        )
      )
    );

    render(withQueryClient(<MfaDisableForm hasEmailPassword={false} />));
    fireEvent.click(screen.getByRole("button", { name: "Nonaktifkan MFA" }));

    expect(await screen.findByRole("alert")).toHaveTextContent(
      "Access token tidak valid atau sudah kedaluwarsa."
    );
    expect(screen.getByRole("link", { name: "Verifikasi ulang dengan Google" })).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Nonaktifkan MFA" })).toBeInTheDocument();
  });
});
