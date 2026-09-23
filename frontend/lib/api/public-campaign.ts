import type { components } from "./generated/openapi";
import { apiRequest, ApiTransportError } from "./client";

export type PublicCampaignDetail = components["schemas"]["PublicCampaignDetail"];

export type PublicCampaignDetailResult =
  | { kind: "success"; data: PublicCampaignDetail }
  | { kind: "not-found" }
  | { kind: "temporarily-unavailable" };

export class PublicCampaignRequestError extends Error {
  constructor() {
    super("The public campaign request failed.");
    this.name = "PublicCampaignRequestError";
  }
}

export async function getPublicCampaignDetail(
  campaignId: string,
): Promise<PublicCampaignDetailResult> {
  const response = await apiRequest(
    `/api/campaigns/${encodeURIComponent(campaignId)}`,
  );

  if (response.status === 404) {
    return { kind: "not-found" };
  }

  if (response.status === 503) {
    return { kind: "temporarily-unavailable" };
  }

  if (!response.ok) {
    throw new PublicCampaignRequestError();
  }

  try {
    return {
      kind: "success",
      data: (await response.json()) as PublicCampaignDetail,
    };
  } catch (error) {
    if (error instanceof ApiTransportError) {
      throw error;
    }

    throw new PublicCampaignRequestError();
  }
}
