import { HttpResponse, http } from "msw";
import { publicCampaignFixtures } from "@/mocks/fixtures/public-campaign";

const tinyPng = new Uint8Array([
  137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 1, 0, 0, 0,
  1, 8, 6, 0, 0, 0, 31, 21, 196, 137, 0, 0, 0, 13, 73, 68, 65, 84, 120, 156, 99, 248,
  246, 225, 241, 127, 0, 9, 114, 3, 201, 112, 123, 217, 47, 0, 0, 0, 0, 73, 69, 78, 68,
  174, 66, 96, 130,
]);

export const publicCampaignHandlers = [
  http.get("/api/campaigns/:campaignId/media/:mediaId/content", () =>
    new HttpResponse(tinyPng, { headers: { "Content-Type": "image/png" } }),
  ),
  http.get("/api/campaigns/:campaignId", ({ params }) => {
    const campaignId = String(params.campaignId);

    if (campaignId === "not-found") {
      return new HttpResponse(null, { status: 404 });
    }

    if (campaignId === "unavailable") {
      return new HttpResponse(null, { status: 503 });
    }

    const fixture = publicCampaignFixtures[campaignId];
    return fixture
      ? HttpResponse.json(fixture)
      : new HttpResponse(null, { status: 404 });
  }),
];
