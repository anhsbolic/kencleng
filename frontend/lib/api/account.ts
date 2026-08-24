import { apiFetch } from "./client";
import type { components } from "./schema";

export type User = components["schemas"]["User"];

/**
 * `GET /account/me` — the current authenticated user's profile,
 * including `roles` (see `lib/types/roles.ts` for how the frontend
 * derives its `GlobalRole` union from this). Mocked in
 * `mocks/handlers.ts` ahead of Account Task #7 actually shipping,
 * per the Mock-First Development Workflow — Dashboard Shell's nav
 * role-filtering is a real consumer today, not speculative.
 *
 * Throws a generic error on a non-OK response — callers (TanStack
 * Query hooks) surface that as a generic error state, never the raw
 * response body (`loading-empty-error-state-conventions.md`).
 */
export async function getMe(): Promise<User> {
  const res = await apiFetch("/account/me", { method: "GET" });
  if (!res.ok) {
    throw new Error("Failed to load account profile");
  }
  return res.json();
}
