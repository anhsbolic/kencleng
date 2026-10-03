import { HttpResponse, http } from "msw";
import type { components } from "@/lib/api/generated/openapi";
import type { SubmitDonationRequest } from "@/lib/api/donation";
import { publicCampaignFixtures } from "@/mocks/fixtures/public-campaign";

type Donation = components["schemas"]["Donation"];
type DonationStatus = components["schemas"]["DonationStatusResponse"];

const submissions = new Map<string, { payload: string; donation: Donation }>();
const donationStatuses = new Map<string, DonationStatus["status"]>();
let nextDonationId = 1;

function validationError(field: string) {
  return HttpResponse.json(
    {
      type: "https://kencleng.dev/errors/validation",
      title: "Validation Error",
      status: 422,
      errors: [{ field, message: "Nilai belum dapat diproses." }],
    },
    { status: 422, headers: { "Content-Type": "application/problem+json" } },
  );
}

export const donationHandlers = [
  http.post("/api/campaigns/:campaignId/donations", async ({ params, request }) => {
    const campaignId = String(params.campaignId);
    const campaign = publicCampaignFixtures[campaignId];
    const idempotencyKey = request.headers.get("Idempotency-Key");
    let payload: SubmitDonationRequest;

    try {
      payload = (await request.json()) as SubmitDonationRequest;
    } catch {
      return validationError("amount");
    }

    if (!campaign || campaign.donation_action.availability !== "available") {
      return HttpResponse.json(
        {
          type: "https://kencleng.dev/errors/campaign-not-eligible",
          title: "Campaign Not Accepting Donations",
          status: 409,
          detail: "Campaign ini sudah tidak menerima donasi baru.",
        },
        { status: 409, headers: { "Content-Type": "application/problem+json" } },
      );
    }

    if (!idempotencyKey) return validationError("amount");
    const amount = Number(payload.amount);
    const cap = Number(campaign.max_donation_amount.amount);
    if (!Number.isInteger(amount) || amount < 5000 || amount > cap) {
      return validationError("amount");
    }

    const serializedPayload = JSON.stringify(payload);
    const replayKey = `${campaignId}:${idempotencyKey}`;
    const previous = submissions.get(replayKey);
    if (previous) {
      if (previous.payload !== serializedPayload) {
        return HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/idempotency-conflict",
            title: "Conflict",
            status: 409,
            detail: "Permintaan ini tidak dapat diproses.",
          },
          { status: 409, headers: { "Content-Type": "application/problem+json" } },
        );
      }
      return HttpResponse.json(previous.donation, { status: 201 });
    }

    const id = `0d000000-0000-4000-8000-${String(nextDonationId++).padStart(12, "0")}`;
    const donation: Donation = {
      id,
      campaign_id: campaignId,
      amount: payload.amount,
      currency_code: payload.currency_code,
      payment_method: "qris",
      is_anonymous: payload.is_anonymous ?? false,
      status: "pending",
      status_token: `sandbox-status-${id}`,
      created_at: new Date().toISOString(),
      ...(payload.guest_name ? { guest_name: payload.guest_name } : {}),
    };
    submissions.set(replayKey, { payload: serializedPayload, donation });
    donationStatuses.set(id, "pending");
    return HttpResponse.json(donation, { status: 201 });
  }),

  http.get("/api/donations/:donationId/status", ({ params, request }) => {
    const donationId = String(params.donationId);
    const credential = request.headers.get("X-Donation-Status-Credential");
    const expectedCredential = `sandbox-status-${donationId}`;
    const status = donationStatuses.get(donationId);
    if (!status || credential !== expectedCredential) {
      return HttpResponse.json(
        {
          type: "https://kencleng.dev/errors/donation-status-not-found",
          title: "Donation Status Not Found",
          status: 404,
          detail: "Status donasi tidak tersedia.",
        },
        {
          status: 404,
          headers: {
            "Content-Type": "application/problem+json",
            "Cache-Control": "private, no-store",
          },
        },
      );
    }
    return HttpResponse.json({ status }, { headers: { "Cache-Control": "private, no-store" } });
  }),
];
