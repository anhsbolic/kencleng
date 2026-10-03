import { setupServer } from "msw/node";
import { publicCampaignHandlers } from "./handlers/public-campaign";
import { donationHandlers } from "./handlers/donation";

export const server = setupServer(...publicCampaignHandlers, ...donationHandlers);
