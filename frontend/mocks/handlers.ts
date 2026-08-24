// Shared MSW request handlers — used by both the Vitest `msw/node`
// setup (mocks/server.ts) and the browser `msw/browser` dev-mode
// worker (mocks/browser.ts). One handler per endpoint, added as the
// page/component that needs it gets built (mock-first workflow,
// scaffold-frontend.md's durable section) — not speculatively ahead
// of demonstrated need.
import { HttpResponse, http } from "msw";
import type { User } from "@/lib/api/account";
import type { UnreadCountResponse } from "@/lib/api/notification";

// Default fixture: a plain logged-in donatur (no elevated roles) —
// the "just the default logged-in-donatur case" the Dashboard Shell
// nav checkpoint (phase0-shared-infra.md Step 7) says not to stop at.
// Individual tests override this via `server.use(...)` with a
// different `roles` array to exercise the other combinations.
const mockUser: User = {
  id: "00000000-0000-0000-0000-000000000001",
  name: "Donatur Contoh",
  email: "donatur@example.com",
  email_verified: true,
  roles: [],
  auth_providers: ["email_password"],
  mfa_enabled: false,
  created_at: new Date().toISOString(),
};

const mockUnreadCount: UnreadCountResponse = { unread_count: 3 };

export const handlers = [
  http.get("/account/me", () => HttpResponse.json(mockUser)),
  http.get("/notifications/unread-count", () =>
    HttpResponse.json(mockUnreadCount)
  ),
];
