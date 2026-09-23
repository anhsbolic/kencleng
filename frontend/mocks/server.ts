import { setupServer } from "msw/node";
import { publicCampaignHandlers } from "./handlers/public-campaign";

export const server = setupServer(...publicCampaignHandlers);
