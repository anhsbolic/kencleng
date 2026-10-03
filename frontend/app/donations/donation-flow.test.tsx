import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import type { components } from "@/lib/api/generated/openapi";
import { campaignFixtureIds } from "@/mocks/fixtures/public-campaign";
import { server } from "@/mocks/server";
import DonationClient from "../campaigns/[campaignId]/donate/donation-client";
import StatusClient from "./[donationId]/status/status-client";

const { mockPush, mockParams } = vi.hoisted(() => ({
  mockPush: vi.fn(),
  mockParams: { campaignId: "not-set", donationId: "not-set" },
}));

vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: mockPush }),
  useParams: () => mockParams,
}));

const acceptedDonation: components["schemas"]["Donation"] = {
  id: "2d000000-0000-4000-8000-000000000001",
  campaign_id: campaignFixtureIds.available,
  amount: "5001",
  currency_code: "IDR",
  payment_method: "qris",
  is_anonymous: false,
  status: "pending",
  status_token: "opaque-status-credential",
  created_at: "2026-10-03T00:00:00Z",
};

describe("guest donation frontend flow", () => {
  beforeEach(() => {
    mockPush.mockReset();
    mockParams.campaignId = campaignFixtureIds.available;
    mockParams.donationId = acceptedDonation.id;
  });

  afterEach(() => {
    server.resetHandlers();
    vi.unstubAllGlobals();
    window.history.replaceState(window.history.state, "", "/");
  });

  it("accepts Rp5.001, submits QRIS with generated contract fields, and hands the token to the fragment route", async () => {
    let requestPayload: unknown;
    let idempotencyKey: string | null = null;
    server.use(http.post("/api/campaigns/:campaignId/donations", async ({ request }) => {
      requestPayload = await request.json();
      idempotencyKey = request.headers.get("Idempotency-Key");
      return HttpResponse.json(acceptedDonation, { status: 201 });
    }));

    render(<DonationClient />);
    expect(await screen.findByRole("heading", { name: /Berikan dukungan untuk/i })).toBeVisible();
    expect(screen.getByText("Batas maksimum tiap donasi pada detail campaign ini. Jumlah yang dapat diterima tetap ditentukan saat permintaan diproses.")).toBeVisible();
    expect(screen.getByText("GoPay")).toBeVisible();
    expect(screen.getByText("ShopeePay")).toBeVisible();
    expect(screen.getByText("Transfer bank")).toBeVisible();
    expect(screen.getAllByText("Tidak tersedia")).toHaveLength(3);

    fireEvent.change(screen.getByLabelText("Jumlah donasi (IDR)"), { target: { value: "5001" } });
    fireEvent.click(screen.getByRole("button", { name: "Kirim donasi simulasi" }));

    await waitFor(() => expect(mockPush).toHaveBeenCalledWith(
      "/donations/" + acceptedDonation.id + "/status#opaque-status-credential",
    ));
    expect(idempotencyKey).toMatch(/^[0-9a-f-]{36}$/i);
    expect(requestPayload).toEqual({
      amount: "5001",
      currency_code: "IDR",
      payment_method: "qris",
      guest_email_status_opt_in: false,
      is_anonymous: false,
    });
  });

  it("keeps local minimum validation field-level and never sends invalid input", async () => {
    let submitted = false;
    server.use(http.post("/api/campaigns/:campaignId/donations", () => {
      submitted = true;
      return HttpResponse.json(acceptedDonation, { status: 201 });
    }));

    render(<DonationClient />);
    await screen.findByRole("heading", { name: /Berikan dukungan untuk/i });
    fireEvent.change(screen.getByLabelText("Jumlah donasi (IDR)"), { target: { value: "4999" } });
    fireEvent.click(screen.getByRole("button", { name: "Kirim donasi simulasi" }));

    expect(await screen.findByText("Minimum donasi adalah Rp 5.000.")).toBeVisible();
    expect(submitted).toBe(false);
  });

  it("does not offer the form when a stale route now has an unavailable action", async () => {
    mockParams.campaignId = campaignFixtureIds.donationUnavailable;
    render(<DonationClient />);

    expect(await screen.findByRole("heading", { name: "Campaign ini tidak menerima donasi baru." })).toBeVisible();
    expect(screen.queryByLabelText("Jumlah donasi (IDR)")).not.toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Kembali ke campaign" })).toHaveAttribute(
      "href",
      "/campaigns/" + campaignFixtureIds.donationUnavailable,
    );
  });

  it("keeps amount 422 feedback field-level and closed Campaign 409 request-level", async () => {
    server.use(http.post("/api/campaigns/:campaignId/donations", () => HttpResponse.json({
      type: "https://kencleng.dev/errors/validation",
      title: "Validation Error",
      status: 422,
      errors: [{ field: "amount", message: "Do not display server text." }],
    }, { status: 422 })));
    const firstRender = render(<DonationClient />);
    await screen.findByRole("heading", { name: /Berikan dukungan untuk/i });
    fireEvent.change(screen.getByLabelText("Jumlah donasi (IDR)"), { target: { value: "5001" } });
    fireEvent.click(screen.getByRole("button", { name: "Kirim donasi simulasi" }));
    expect(await screen.findByText("Jumlah belum dapat diproses. Periksa jumlah dan batas per donasi yang ditampilkan.")).toBeVisible();
    expect(screen.queryByText(/remaining capacity|close reason|Do not display/i)).not.toBeInTheDocument();
    firstRender.unmount();

    server.use(http.post("/api/campaigns/:campaignId/donations", () => HttpResponse.json({
      type: "https://kencleng.dev/errors/campaign-not-eligible",
      title: "Campaign Not Accepting Donations",
      status: 409,
      detail: "Campaign ini sudah tidak menerima donasi baru.",
    }, { status: 409 })));
    render(<DonationClient />);
    await screen.findByRole("heading", { name: /Berikan dukungan untuk/i });
    fireEvent.change(screen.getByLabelText("Jumlah donasi (IDR)"), { target: { value: "5001" } });
    fireEvent.click(screen.getByRole("button", { name: "Kirim donasi simulasi" }));
    expect(await screen.findByRole("alert")).toHaveTextContent("Campaign ini sudah tidak menerima donasi baru");
    expect(screen.queryByText("Jumlah belum dapat diproses. Periksa jumlah dan batas per donasi yang ditampilkan.")).not.toBeInTheDocument();
  });

  it("keeps an ambiguous retry on the same idempotency key and payload", async () => {
    const seenKeys: Array<string | null> = [];
    const seenBodies: unknown[] = [];
    let calls = 0;
    server.use(http.post("/api/campaigns/:campaignId/donations", async ({ request }) => {
      calls += 1;
      seenKeys.push(request.headers.get("Idempotency-Key"));
      seenBodies.push(await request.json());
      if (calls === 1) return HttpResponse.error();
      return HttpResponse.json(acceptedDonation, { status: 201 });
    }));

    render(<DonationClient />);
    await screen.findByRole("heading", { name: /Berikan dukungan untuk/i });
    fireEvent.change(screen.getByLabelText("Jumlah donasi (IDR)"), { target: { value: "5001" } });
    fireEvent.click(screen.getByRole("button", { name: "Kirim donasi simulasi" }));
    const retry = await screen.findByRole("button", { name: "Coba kirim ulang" });
    expect(screen.getByLabelText("Jumlah donasi (IDR)")).toBeDisabled();
    fireEvent.click(retry);

    await waitFor(() => expect(mockPush).toHaveBeenCalled());
    expect(seenKeys[0]).toBeTruthy();
    expect(seenKeys[1]).toBe(seenKeys[0]);
    expect(seenBodies[1]).toEqual(seenBodies[0]);
  });

  it("sends an opted-in email only with the exact approved disclosure", async () => {
    let requestPayload: Record<string, unknown> | undefined;
    server.use(http.post("/api/campaigns/:campaignId/donations", async ({ request }) => {
      requestPayload = await request.json() as Record<string, unknown>;
      return HttpResponse.json(acceptedDonation, { status: 201 });
    }));

    render(<DonationClient />);
    await screen.findByRole("heading", { name: /Berikan dukungan untuk/i });
    expect(screen.getByText("Verifikasi email dalam 24 jam sejak alamat dicatat. Jika tidak diverifikasi, alamat dihapus dan pemberitahuan tidak dikirim. Donasi tetap berjalan.")).toBeVisible();
    fireEvent.change(screen.getByLabelText("Jumlah donasi (IDR)"), { target: { value: "5000" } });
    fireEvent.click(screen.getByLabelText("Kirim pemberitahuan status donasi melalui email (opsional)"));
    fireEvent.change(screen.getByLabelText("Alamat email"), { target: { value: "donor@example.test" } });
    fireEvent.click(screen.getByRole("button", { name: "Kirim donasi simulasi" }));

    await waitFor(() => expect(mockPush).toHaveBeenCalled());
    expect(requestPayload?.guest_email_status_opt_in).toBe(true);
    expect(requestPayload?.guest_email).toBe("donor@example.test");
  });

  it("associates opted-in email validation with the field and clears the stale description", async () => {
    render(<DonationClient />);
    await screen.findByRole("heading", { name: /Berikan dukungan untuk/i });
    fireEvent.change(screen.getByLabelText("Jumlah donasi (IDR)"), { target: { value: "5000" } });
    fireEvent.click(screen.getByLabelText("Kirim pemberitahuan status donasi melalui email (opsional)"));
    const email = screen.getByLabelText("Alamat email");
    fireEvent.click(screen.getByRole("button", { name: "Kirim donasi simulasi" }));

    const error = await screen.findByText("Masukkan alamat email yang valid untuk menerima pemberitahuan status.");
    expect(email).toHaveAttribute("aria-invalid", "true");
    expect(email).toHaveAttribute("aria-describedby", "guest-email-error");
    expect(error).toHaveAttribute("id", "guest-email-error");

    fireEvent.change(email, { target: { value: "donor@example.test" } });
    expect(await screen.findByLabelText("Alamat email")).not.toHaveAttribute("aria-invalid", "true");
    expect(screen.getByLabelText("Alamat email")).not.toHaveAttribute("aria-describedby", "guest-email-error");
    expect(screen.queryByText("Masukkan alamat email yang valid untuk menerima pemberitahuan status.")).not.toBeInTheDocument();
  });

  it("cleans the visible fragment before requesting status and renders status only", async () => {
    let receivedCredential: string | null = null;
    mockParams.donationId = acceptedDonation.id;
    window.history.replaceState(
      window.history.state,
      "",
      "/donations/" + acceptedDonation.id + "/status#opaque-status-credential",
    );
    server.use(http.get("/api/donations/:donationId/status", ({ request }) => {
      receivedCredential = request.headers.get("X-Donation-Status-Credential");
      return HttpResponse.json({ status: "pending" });
    }));

    render(<StatusClient />);
    expect(screen.getByRole("status")).toHaveTextContent("Memeriksa status donasi…");
    expect(await screen.findByRole("heading", { name: "Menunggu hasil simulasi" })).toBeVisible();
    expect(screen.getByRole("status")).toHaveTextContent("Menunggu hasil simulasi");
    expect(window.location.hash).toBe("");
    expect(receivedCredential).toBe("opaque-status-credential");
    expect(screen.queryByText(acceptedDonation.id)).not.toBeInTheDocument();
    expect(screen.queryByText(/donor@example|email/i)).not.toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Periksa status lagi" })).toBeVisible();
  });

  it.each([
    ["pending", "Menunggu hasil simulasi"],
    ["success", "Hasil simulasi donasi: berhasil"],
    ["failed", "Hasil simulasi donasi: gagal"],
    ["unavailable", "Link status tidak tersedia atau mungkin kedaluwarsa."],
  ] as const)("announces a manual status check that resolves to %s", async (result, announcement) => {
    window.history.replaceState(
      window.history.state,
      "",
      "/donations/" + acceptedDonation.id + "/status#opaque-status-credential",
    );
    server.use(http.get("/api/donations/:donationId/status", () => HttpResponse.json({ status: "pending" })));
    render(<StatusClient />);
    expect(await screen.findByRole("heading", { name: "Menunggu hasil simulasi" })).toBeVisible();

    server.use(http.get("/api/donations/:donationId/status", () => result === "unavailable"
      ? HttpResponse.json({}, { status: 404 })
      : HttpResponse.json({ status: result })));
    fireEvent.click(screen.getByRole("button", { name: "Periksa status lagi" }));
    await waitFor(() => expect(screen.getByRole("status")).toHaveTextContent(announcement));
  });

  it.each([
    ["success", "Hasil simulasi donasi: berhasil"],
    ["failed", "Hasil simulasi donasi: gagal"],
  ] as const)("labels terminal %s as a sandbox simulation", async (status, label) => {
    window.history.replaceState(
      window.history.state,
      "",
      "/donations/" + acceptedDonation.id + "/status#opaque-status-credential",
    );
    server.use(http.get("/api/donations/:donationId/status", () => HttpResponse.json({ status })));

    render(<StatusClient />);
    expect(await screen.findByRole("heading", { name: label })).toBeVisible();
    expect(screen.getByRole("status")).toHaveTextContent(label);
    expect(screen.getByText("Status ini berasal dari hasil simulasi sandbox.")).toBeVisible();
    expect(screen.queryByRole("button", { name: "Periksa status lagi" })).not.toBeInTheDocument();
  });

  it("uses the same generic status-link failure when a credential is missing", async () => {
    window.history.replaceState(window.history.state, "", "/donations/" + acceptedDonation.id + "/status");
    render(<StatusClient />);
    expect(await screen.findByRole("heading", {
      name: "Link status tidak tersedia atau mungkin kedaluwarsa.",
    })).toBeVisible();
    expect(screen.getByRole("status")).toHaveTextContent("Link status tidak tersedia atau mungkin kedaluwarsa.");
    expect(window.location.hash).toBe("");
  });
});
