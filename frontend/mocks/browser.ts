// MSW browser-mode worker — intercepts `fetch` during `npm run dev`
// when NEXT_PUBLIC_API_MOCKING=true (see components/providers/
// mocking-provider.tsx). Shares `handlers.ts` with the Vitest
// node-mode server (mocks/server.ts): one handler set, two MSW entry
// points.
import { setupWorker } from "msw/browser";
import { handlers } from "./handlers";

export const worker = setupWorker(...handlers);
