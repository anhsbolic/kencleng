import type { components } from "@/lib/api/generated/openapi";

export type PublicCampaignDetail = components["schemas"]["PublicCampaignDetail"];

export const campaignFixtureIds = {
  available: "4f9d83f0-3c71-4f04-93d9-90bfbf868bf3",
  absentMedia: "6a2e4d8c-8a9e-4c91-a913-4f365c4ad101",
  unavailableMedia: "7b3f5e9d-9baf-4d02-b024-5a476d5be202",
  fundingUnavailable: "8c4a6fad-acb0-4e13-c135-6b587e6cf303",
  zeroFunding: "9d5b70be-bdc1-4f24-d246-7c698f7df404",
  aboveTarget: "ae6c81cf-ced2-4035-e357-8d7a908ef505",
  notComputable: "bf7d92d0-dfe3-4146-f468-9e8ba19fa606",
  hostileOrganizer: "c08ea3e1-e0f4-4257-a579-af9cb2a0b707",
  longContent: "d19fb4f2-f105-4368-b68a-b0adc3b1c808",
} as const;

const mediaId = "26a0a959-86ca-4df7-8e8d-00d1f1e247f3";

const baseCampaign = {
  title: "Dapur bersama untuk keluarga di Kampung Cempaka",
  purpose: {
    content: "Dana ini ditujukan untuk menyiapkan bahan pangan dan peralatan dapur bersama.",
    source: "organizer",
  },
  story: {
    content:
      "Pengelola kampung mengumpulkan cerita dan kebutuhan dari warga sebelum memulai penggalangan.",
    source: "organizer",
  },
  steward: {
    id: "e2a0c5f3-12a8-4d55-8d50-9c4eab7fda09",
    name: "Ruang Warga Cempaka",
  },
  lifecycle: {
    public_state: "fundraising",
    published_at: "2026-09-01T08:00:00Z",
    fundraising_ends_at: "2026-10-31T16:59:59Z",
  },
  funding: {
    availability: "available",
    currency: "IDR",
    target_amount: "5000000.00",
    collected_amount: "1250000.00",
    progress: {
      state: "computed",
      percentage: "25",
      relationship: "below_target",
    },
  },
  media: {
    state: "available",
    items: [
      {
        id: mediaId,
        content_url: `/api/campaigns/${campaignFixtureIds.available}/media/${mediaId}/content`,
        content_type: "image/png",
        alt_text: "Peralatan memasak tersusun di dapur bersama.",
        caption: "Persiapan dapur bersama oleh pengelola.",
        source: "organizer",
      },
    ],
  },
  donation_action: {
    availability: "unavailable",
    reason: "donation_flow_not_available",
  },
} as const satisfies Omit<PublicCampaignDetail, "id">;

function createCampaign(
  id: string,
  overrides: Partial<PublicCampaignDetail> = {},
): PublicCampaignDetail {
  return { ...baseCampaign, id, ...overrides } as PublicCampaignDetail;
}

export const publicCampaignFixtures: Record<string, PublicCampaignDetail> = {
  [campaignFixtureIds.available]: createCampaign(campaignFixtureIds.available),
  [campaignFixtureIds.absentMedia]: createCampaign(campaignFixtureIds.absentMedia, {
    media: { state: "absent", items: [] },
  }),
  [campaignFixtureIds.unavailableMedia]: createCampaign(campaignFixtureIds.unavailableMedia, {
    media: { state: "unavailable", reason: "temporarily_unavailable", items: [] },
  }),
  [campaignFixtureIds.fundingUnavailable]: createCampaign(campaignFixtureIds.fundingUnavailable, {
    funding: { availability: "unavailable", reason: "not_available" },
  }),
  [campaignFixtureIds.zeroFunding]: createCampaign(campaignFixtureIds.zeroFunding, {
    funding: {
      availability: "available",
      currency: "IDR",
      target_amount: "99999999999999999.99",
      collected_amount: "0.00",
      progress: { state: "computed", percentage: "0", relationship: "below_target" },
    },
  }),
  [campaignFixtureIds.aboveTarget]: createCampaign(campaignFixtureIds.aboveTarget, {
    funding: {
      availability: "available",
      currency: "IDR",
      target_amount: "5000000.00",
      collected_amount: "6250000.00",
      progress: { state: "computed", percentage: "125.50", relationship: "above_target" },
    },
  }),
  [campaignFixtureIds.notComputable]: createCampaign(campaignFixtureIds.notComputable, {
    funding: {
      availability: "available",
      currency: "IDR",
      target_amount: "0.00",
      collected_amount: "0.00",
      progress: {
        state: "not_computable",
        percentage: null,
        relationship: "not_computable",
      },
    },
  }),
  [campaignFixtureIds.hostileOrganizer]: createCampaign(campaignFixtureIds.hostileOrganizer, {
    purpose: { content: "<img src=x onerror=alert('xss')> Kebutuhan dapur", source: "organizer" },
    story: { content: "<strong>Bukan HTML</strong>; ini tetap cerita dari pengelola.", source: "organizer" },
  }),
  [campaignFixtureIds.longContent]: createCampaign(campaignFixtureIds.longContent, {
    title: "Dapur bersama untuk keluarga yang tetap perlu ruang makan aman dan terbuka di Kampung Cempaka pada musim hujan ini",
    story: {
      content:
        "Pengelola menjelaskan kebutuhan warga secara bertahap agar setiap orang dapat membaca konteksnya. ".repeat(10),
      source: "organizer",
    },
  }),
};
