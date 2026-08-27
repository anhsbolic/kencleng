import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { describe, expect, it, vi } from "vitest";
import { server } from "@/mocks/server";
import { MfaSection } from "./mfa-section";

vi.mock("qrcode.react", () => ({
  QRCodeSVG: (props: { value: string }) => <div data-testid="qr-code" data-value={props.value} />,
}));

function baseUser(overrides: Partial<Record<string, unknown>> = {}) {
  return {
    id: "u1",
    name: "Test User",
    email: "test@example.com",
    email_verified: true,
    roles: [],
    auth_providers: ["email_password"],
    mfa_enabled: false,
    created_at: new Date().toISOString(),
    ...overrides,
  };
}

function mockMe(overrides: Partial<Record<string, unknown>> = {}) {
  server.use(http.get("/account/me", () => HttpResponse.json(baseUser(overrides))));
}

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false }, mutations: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

describe("MfaSection", () => {
  it("shows a skeleton before the account query resolves (R1)", () => {
    const { container } = render(withQueryClient(<MfaSection />));

    expect(container.querySelector(".animate-pulse")).toBeInTheDocument();
    expect(screen.queryByText("Autentikasi Dua Faktor (MFA)")).not.toBeInTheDocument();
  });

  it("not enrolled: renders MfaEnrollFlow (R2)", async () => {
    mockMe({ mfa_enabled: false });
    render(withQueryClient(<MfaSection />));

    expect(await screen.findByRole("button", { name: "Aktifkan MFA" })).toBeInTheDocument();
  });

  it("enrolled: renders MfaDisableForm (R3)", async () => {
    mockMe({ mfa_enabled: true, auth_providers: ["email_password"] });
    render(withQueryClient(<MfaSection />));

    expect(await screen.findByLabelText("Konfirmasi password")).toBeInTheDocument();
  });

  it("codes-once view survives an account.me refetch and only clears on explicit acknowledgment, with no extra API call (R12/R13/R22)", async () => {
    let meCallCount = 0;
    server.use(
      http.get("/account/me", () => {
        meCallCount += 1;
        // First fetch (initial mount): not enrolled. After
        // enroll/confirm succeeds and useMfaEnrollConfirm's onSuccess
        // invalidates the query, the refetch below reflects the
        // backend's now-true mfa_enabled — this is the exact race R12
        // guards against.
        return HttpResponse.json(baseUser({ mfa_enabled: meCallCount > 1 }));
      })
    );
    let disableCallCount = 0;
    server.use(
      http.post("/account/security/mfa/disable", () => {
        disableCallCount += 1;
        return HttpResponse.json({ message: "MFA berhasil dinonaktifkan." }, { status: 200 });
      })
    );

    render(withQueryClient(<MfaSection />));

    fireEvent.click(await screen.findByRole("button", { name: "Aktifkan MFA" }));
    await screen.findByTestId("qr-code");

    fireEvent.change(screen.getByLabelText("Kode OTP"), { target: { value: "123456" } });
    fireEvent.click(screen.getByRole("button", { name: "Konfirmasi" }));

    // R9/R12 — codes-once view appears and survives the account.me
    // refetch that useMfaEnrollConfirm's own onSuccess triggers.
    expect(await screen.findByText("1a2b3c4d")).toBeInTheDocument();
    await waitFor(() => expect(meCallCount).toBeGreaterThan(1));
    // Still showing codes, not MfaDisableForm, even after the refetch settled.
    expect(screen.getByText("1a2b3c4d")).toBeInTheDocument();
    expect(screen.queryByLabelText("Konfirmasi password")).not.toBeInTheDocument();

    // R13 — acknowledging clears the codes view and renders
    // MfaDisableForm, with no additional call to the disable endpoint.
    fireEvent.click(screen.getByRole("button", { name: "Saya sudah menyimpan kode ini" }));

    expect(await screen.findByLabelText("Konfirmasi password")).toBeInTheDocument();
    expect(screen.queryByText("1a2b3c4d")).not.toBeInTheDocument();
    expect(disableCallCount).toBe(0);
  });
});
