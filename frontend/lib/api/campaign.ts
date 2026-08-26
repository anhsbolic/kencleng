import { apiFetch } from "./client";
import type { components } from "./schema";

export type CampaignListItem = components["schemas"]["CampaignListItem"];
export type CampaignListResponse = components["schemas"]["CampaignListResponse"];

/**
 * `GET /campaigns` — public campaign listing (`security: []`), backs
 * `/`'s highlighted-campaigns section (mock-scope, per `page-map.md`
 * footnote 1) and, later, `/campaign`'s full browse page. No
 * category/search/pagination params are passed here — `/`'s section
 * is a fixed fixture set, not a filterable browse UI.
 *
 * `CampaignListItem` has no `organization` field (only
 * `organization_id`) — deliberate per `campaign.yaml`'s composite-only-
 * on-the-detail-endpoint design (`docs/spec/4-campaign/features/
 * 02-campaign-detail-listing.md`). Don't add an `organization` field to
 * any mock/fixture built against this type — the real endpoint can't
 * supply it (see techplan Decision 1).
 *
 * Throws a generic error on a non-OK response — callers (TanStack
 * Query hooks) surface that as a generic error state, never the raw
 * response body.
 */
export async function getCampaigns(): Promise<CampaignListResponse> {
  const res = await apiFetch("/campaigns", { method: "GET" });
  if (!res.ok) {
    throw new Error("Failed to load campaigns");
  }
  return res.json();
}
