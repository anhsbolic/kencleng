import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { describe, expect, it, vi } from "vitest";
import { server } from "@/mocks/server";
import { MfaEnrollFlow } from "./mfa-enroll-flow";

// Mock the third-party QR rendering library at the module boundary —
// a network-layer mock (MSW) can't reach "what a third-party rendering
// library drew" (component-test-mocking-discipline.md's narrow, justified
// exception). Asserts the wrapper receives the correct `value` prop
// instead of inspecting rendered SVG path data.
vi.mock("qrcode.react", () => ({
  QRCodeSVG: (props: { value: string }) => (
    <div data-testid="qr-code" data-value={props.value} />
  ),
}));

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false }, mutations: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

async function activate() {
  fireEvent.click(screen.getByRole("button", { name: "Aktifkan MFA" }));
  await screen.findByTestId("qr-code");
}

function submitCode(code = "123456") {
  fireEvent.change(screen.getByLabelText("Kode OTP"), { target: { value: code } });
  fireEvent.click(screen.getByRole("button", { name: "Konfirmasi" }));
}

describe("MfaEnrollFlow", () => {
  it("does not call enroll on mount, only on click (R4)", () => {
    render(withQueryClient(<MfaEnrollFlow onEnrolled={vi.fn()} />));

    expect(screen.getByRole("button", { name: "Aktifkan MFA" })).toBeInTheDocument();
    expect(screen.queryByTestId("qr-code")).not.toBeInTheDocument();
  });

  it("enroll success renders the QR + manual-entry secret + totp_code form (R5/R24)", async () => {
    render(withQueryClient(<MfaEnrollFlow onEnrolled={vi.fn()} />));

    await activate();

    expect(screen.getByTestId("qr-code")).toHaveAttribute(
      "data-value",
      expect.stringContaining("otpauth://totp/")
    );
    expect(screen.getByText("JBSWY3DPEHPK3PXP")).toBeInTheDocument();
    expect(screen.getByLabelText("Kode OTP")).toBeInTheDocument();
  });

  it("hides the manual-entry line when the URI carries no secret (R24 defensive fallback)", async () => {
    server.use(
      http.post("/account/security/mfa/enroll", () =>
        HttpResponse.json({ otpauth_uri: "otpauth://totp/Label?issuer=Kencleng" }, { status: 200 })
      )
    );

    render(withQueryClient(<MfaEnrollFlow onEnrolled={vi.fn()} />));
    await activate();

    expect(screen.queryByText(/Masukkan kode ini secara manual/)).not.toBeInTheDocument();
  });

  it("shows a generic banner on enroll 409 (defensive, R6), focus moves into it (R21)", async () => {
    server.use(
      http.post("/account/security/mfa/enroll", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/mfa-already-enabled",
            title: "MFA Already Enabled",
            status: 409,
            detail: "MFA sudah aktif.",
          },
          { status: 409 }
        )
      )
    );

    render(withQueryClient(<MfaEnrollFlow onEnrolled={vi.fn()} />));
    fireEvent.click(screen.getByRole("button", { name: "Aktifkan MFA" }));

    const banner = await screen.findByRole("alert");
    expect(banner).toHaveTextContent("MFA sudah aktif.");
    await waitFor(() => expect(document.activeElement).toBe(banner.parentElement));
    expect(screen.getByRole("button", { name: "Aktifkan MFA" })).toBeInTheDocument();
  });

  it("shows a generic banner on enroll network failure (R7)", async () => {
    server.use(http.post("/account/security/mfa/enroll", () => HttpResponse.error()));

    render(withQueryClient(<MfaEnrollFlow onEnrolled={vi.fn()} />));
    fireEvent.click(screen.getByRole("button", { name: "Aktifkan MFA" }));

    expect(await screen.findByText("Terjadi kesalahan. Silakan coba lagi.")).toBeInTheDocument();
  });

  it("submits totp_code to enroll/confirm (R8)", async () => {
    let receivedBody: unknown;
    server.use(
      http.post("/account/security/mfa/enroll/confirm", async ({ request }) => {
        receivedBody = await request.json();
        return HttpResponse.json({ backup_codes: Array(10).fill("code") }, { status: 200 });
      })
    );

    render(withQueryClient(<MfaEnrollFlow onEnrolled={vi.fn()} />));
    await activate();
    submitCode("654321");

    await waitFor(() => expect(receivedBody).toEqual({ totp_code: "654321" }));
  });

  it("confirm success calls onEnrolled with the 10-item backup_codes array (R9 component-layer)", async () => {
    const onEnrolled = vi.fn();
    render(withQueryClient(<MfaEnrollFlow onEnrolled={onEnrolled} />));
    await activate();
    submitCode();

    await waitFor(() => expect(onEnrolled).toHaveBeenCalledTimes(1));
    expect(onEnrolled.mock.calls[0][0]).toHaveLength(10);
  });

  it("confirm 422 shows the fixed field message, QR/form remain mounted (R10)", async () => {
    server.use(
      http.post("/account/security/mfa/enroll/confirm", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/validation-failed",
            title: "Validation Failed",
            status: 422,
            errors: [{ field: "totp_code", message: "Invalid code" }],
          },
          { status: 422 }
        )
      )
    );

    render(withQueryClient(<MfaEnrollFlow onEnrolled={vi.fn()} />));
    await activate();
    submitCode("000000");

    expect(await screen.findByText("Kode tidak valid, coba lagi.")).toBeInTheDocument();
    // QR/form stay mounted — not reverted to the idle "Aktifkan MFA" step.
    expect(screen.getByTestId("qr-code")).toBeInTheDocument();
    expect(screen.getByLabelText("Kode OTP")).toBeInTheDocument();
  });

  it("confirm network failure shows a generic banner, form stays interactive (R11)", async () => {
    server.use(http.post("/account/security/mfa/enroll/confirm", () => HttpResponse.error()));

    render(withQueryClient(<MfaEnrollFlow onEnrolled={vi.fn()} />));
    await activate();
    submitCode();

    expect(await screen.findByText("Terjadi kesalahan. Silakan coba lagi.")).toBeInTheDocument();
    expect(screen.getByLabelText("Kode OTP")).toBeInTheDocument();
  });
});
