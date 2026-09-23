import { setupWorker } from "msw/browser";
import { publicCampaignHandlers } from "./handlers/public-campaign";

export const worker = setupWorker(...publicCampaignHandlers);
