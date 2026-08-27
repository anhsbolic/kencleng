import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { server } from "@/mocks/server";
import { accountKeys } from "@/lib/hooks/use-account-me";
import { useAuthStore } from "@/lib/stores/auth-store";
import { VerifyEmailStatus } from "./verify-email-status";

let mockToken: string | null = "valid-token";

vi.mock("next/navigation", () => ({
  useSearchParams: () => ({
    get: (key: string) => (key === "token" ? mockToken : null),
  }),
}));

beforeEach(() => {
  useAuthStore.setState({ accessToken: null });
});

function withQueryClient(
  children: ReactNode,
  queryClient: QueryClient = new QueryClient({
    defaultOptions: { mutations: { retry: false } },
  })
) {
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

describe("VerifyEmailStatus", () => {
  it("shows a generic invalid-link message when the token is missing (R11)", () => {
    mockToken = null;
    render(withQueryClient(<VerifyEmailStatus />));

    expect(
      screen.getByText("Link verifikasi tidak valid atau sudah digunakan.")
    ).toBeInTheDocument();
  });

  it("fires verifyEmail exactly once per token, even under a forced re-render (R12)", async () => {
    mockToken = "valid-token";
    let callCount = 0;
    server.use(
      http.post("/auth/verify-email", () => {
        callCount += 1;
        return HttpResponse.json({ message: "Email berhasil diverifikasi." }, { status: 200 });
      })
    );

    const { rerender } = render(withQueryClient(<VerifyEmailStatus />));
    rerender(withQueryClient(<VerifyEmailStatus />));
    rerender(withQueryClient(<VerifyEmailStatus />));

    await waitFor(() => expect(callCount).toBe(1));
  });

  it("shows the verified message plus a link to /login on 200 (R13)", async () => {
    mockToken = "valid-token";
    render(withQueryClient(<VerifyEmailStatus />));

    expect(await screen.findByText("Email berhasil diverifikasi.")).toBeInTheDocument();
    expect(screen.getByRole("link", { name: /masuk sekarang/i })).toHaveAttribute(
      "href",
      "/login"
    );
  });

  it("shows the expired message plus a resend control on 410 (R14)", async () => {
    mockToken = "expired-token";
    server.use(
      http.post("/auth/verify-email", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/token-expired",
            title: "Token Expired",
            status: 410,
            detail: "Link verifikasi sudah kedaluwarsa. Silakan minta kirim ulang.",
          },
          { status: 410 }
        )
      )
    );

    render(withQueryClient(<VerifyEmailStatus />));

    expect(
      await screen.findByText("Link verifikasi sudah kedaluwarsa. Silakan minta kirim ulang.")
    ).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /kirim ulang/i })).toBeInTheDocument();
  });

  it("shows a generic invalid-link message on 404 (R15)", async () => {
    mockToken = "used-token";
    server.use(
      http.post("/auth/verify-email", () =>
        HttpResponse.json({ type: "about:blank", title: "Not Found", status: 404 }, { status: 404 })
      )
    );

    render(withQueryClient(<VerifyEmailStatus />));

    expect(
      await screen.findByText("Link verifikasi tidak valid atau sudah digunakan.")
    ).toBeInTheDocument();
  });

  it("moves focus into the result heading once the loading state resolves (R16)", async () => {
    mockToken = "valid-token";
    render(withQueryClient(<VerifyEmailStatus />));

    const heading = await screen.findByRole("heading", { name: "Verifikasi Email" });
    await waitFor(() => expect(heading).toHaveFocus());
  });

  it("shows a generic banner on a network failure, never raw error text (R6)", async () => {
    mockToken = "valid-token";
    server.use(http.post("/auth/verify-email", () => HttpResponse.error()));

    render(withQueryClient(<VerifyEmailStatus />));

    expect(await screen.findByText("Terjadi kesalahan. Silakan coba lagi.")).toBeInTheDocument();
  });

  it("shows the backend's rate-limit detail verbatim on 429 (R10)", async () => {
    mockToken = "valid-token";
    server.use(
      http.post("/auth/verify-email", () =>
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

    render(withQueryClient(<VerifyEmailStatus />));

    expect(
      await screen.findByText("Terlalu banyak percobaan gagal. Coba lagi dalam 15 menit.")
    ).toBeInTheDocument();
  });

  // techplan account/05-account-linking, R19/D6 — authenticated-caller
  // additive fix (Branch 1's step 2 of the 3-step linking flow).
  it("links to /dashboard/security when the caller is authenticated, instead of /login (R19)", async () => {
    mockToken = "valid-token";
    useAuthStore.setState({ accessToken: "still-valid-token" });
    render(withQueryClient(<VerifyEmailStatus />));

    const link = await screen.findByRole("link", { name: "Kembali ke Keamanan" });
    expect(link).toHaveAttribute("href", "/dashboard/security");
    expect(screen.queryByRole("link", { name: /masuk sekarang/i })).not.toBeInTheDocument();
  });

  it("invalidates account.me on success, regardless of auth state (R19)", async () => {
    mockToken = "valid-token";
    const queryClient = new QueryClient({ defaultOptions: { mutations: { retry: false } } });
    queryClient.setQueryData(accountKeys.me(), { id: "1" });

    render(withQueryClient(<VerifyEmailStatus />, queryClient));

    await screen.findByText("Email berhasil diverifikasi.");
    await waitFor(() =>
      expect(queryClient.getQueryState(accountKeys.me())?.isInvalidated).toBe(true)
    );
  });
});
