import type { components } from "./generated/openapi";
import { apiRequest, ApiTransportError } from "./client";

export type SubmitDonationRequest = components["schemas"]["SubmitDonationRequest"];
export type DonationSubmission = components["schemas"]["Donation"];
export type DonationStatus = components["schemas"]["DonationStatusResponse"];

export type SubmitDonationResult =
  | { kind: "accepted"; data: DonationSubmission }
  | { kind: "validation-error"; fields: string[] }
  | { kind: "conflict" }
  | { kind: "ambiguous" }
  | { kind: "request-failure" };

export class DonationTransportError extends Error {
  constructor() {
    super("The donation request could not be completed.");
    this.name = "DonationTransportError";
  }
}

export async function submitDonation(
  campaignId: string,
  payload: SubmitDonationRequest,
  idempotencyKey: string,
): Promise<SubmitDonationResult> {
  let response: Response;

  try {
    response = await apiRequest(
      `/api/campaigns/${encodeURIComponent(campaignId)}/donations`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Idempotency-Key": idempotencyKey,
        },
        body: JSON.stringify(payload),
      },
    );
  } catch (error) {
    if (error instanceof ApiTransportError) {
      throw new DonationTransportError();
    }
    throw error;
  }

  if (response.status === 201) {
    try {
      return {
        kind: "accepted",
        data: (await response.json()) as DonationSubmission,
      };
    } catch {
      // The request may have been accepted even when its response is unreadable.
      return { kind: "ambiguous" };
    }
  }

  if (response.status === 422) {
    try {
      const body = (await response.json()) as {
        errors?: Array<{ field?: unknown }>;
      };
      return {
        kind: "validation-error",
        fields: Array.isArray(body.errors)
          ? body.errors.flatMap((item) =>
              typeof item.field === "string" ? [item.field] : [],
            )
          : [],
      };
    } catch {
      return { kind: "validation-error", fields: [] };
    }
  }

  if (response.status === 409) {
    return { kind: "conflict" };
  }

  if (response.status >= 500) {
    return { kind: "ambiguous" };
  }

  return { kind: "request-failure" };
}

export type DonationStatusResult =
  | { kind: "success"; data: DonationStatus }
  | { kind: "unavailable" };

export async function getDonationStatus(
  donationId: string,
  credential: string,
): Promise<DonationStatusResult> {
  let response: Response;

  try {
    response = await apiRequest(
      `/api/donations/${encodeURIComponent(donationId)}/status`,
      {
        headers: { "X-Donation-Status-Credential": credential },
        cache: "no-store",
      },
    );
  } catch (error) {
    if (error instanceof ApiTransportError) {
      throw new DonationTransportError();
    }
    throw error;
  }

  if (!response.ok) {
    return { kind: "unavailable" };
  }

  try {
    const body = (await response.json()) as DonationStatus;
    if (!(["pending", "success", "failed"] as string[]).includes(body.status)) {
      return { kind: "unavailable" };
    }
    return {
      kind: "success",
      data: body,
    };
  } catch {
    return { kind: "unavailable" };
  }
}
