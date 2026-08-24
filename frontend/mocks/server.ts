// MSW node-mode server — intercepts `fetch` inside Vitest. Shares
// `handlers.ts` with the browser dev-mode worker (mocks/browser.ts).
import { setupServer } from "msw/node";
import { handlers } from "./handlers";

export const server = setupServer(...handlers);
