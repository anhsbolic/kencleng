import { setupWorker } from "msw/browser";
import { publicCampaignHandlers } from "./handlers/public-campaign";
import { donationHandlers } from "./handlers/donation";

export const worker = setupWorker(...publicCampaignHandlers, ...donationHandlers);
