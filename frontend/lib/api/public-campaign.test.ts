import { HttpResponse, http } from "msw";
import { describe, expect, it } from "vitest";
import { campaignFixtureIds, publicCampaignFixtures } from "@/mocks/fixtures/public-campaign";
import { server } from "@/mocks/server";
import {
  getPublicCampaignDetail,
  PublicCampaignRequestError,
} from "./public-campaign";

describe("getPublicCampaignDetail", () => {
  it("uses the one encoded public endpoint and returns the generated fixture", async () => {
    let requestedPath = "";
    server.use(
      http.get("/api/campaigns/:campaignId", ({ request }) => {
        requestedPath = new URL(request.url).pathname;
        return HttpResponse.json(publicCampaignFixtures[campaignFixtureIds.available]);
      }),
    );

    const result = await getPublicCampaignDetail(campaignFixtureIds.available);

    expect(requestedPath).toBe(`/api/campaigns/${campaignFixtureIds.available}`);
    expect(result).toEqual({
      kind: "success",
      data: publicCampaignFixtures[campaignFixtureIds.available],
    });
  });

  it("keeps a public 404 non-disclosing and classifies 503 as retryable", async () => {
    await expect(getPublicCampaignDetail("not-found")).resolves.toEqual({ kind: "not-found" });
    await expect(getPublicCampaignDetail("unavailable")).resolves.toEqual({
      kind: "temporarily-unavailable",
    });
  });

  it("does not expose an unexpected response body through the error", async () => {
    server.use(
      http.get("/api/campaigns/:campaignId", () =>
        HttpResponse.json({ detail: "internal diagnostic must not reach the view" }, { status: 500 }),
      ),
    );

    await expect(getPublicCampaignDetail("unexpected")).rejects.toBeInstanceOf(
      PublicCampaignRequestError,
    );
  });
});
