import { apiFetch } from "./client";
import type { components } from "./schema";

export type UnreadCountResponse = components["schemas"]["UnreadCountResponse"];

/**
 * `GET /notifications/unread-count` — backs the Dashboard Shell's
 * persistent header badge (`page-map.md` Cross-Cutting UI Elements).
 * Mocked ahead of `notification` domain's backend track starting, per
 * the Mock-First Development Workflow (`scaffold-frontend.md`) —
 * the contract already exists in `api/openapi.yaml`, so there's no
 * reason the badge can't be built and visually correct today.
 */
export async function getUnreadCount(): Promise<UnreadCountResponse> {
  const res = await apiFetch("/notifications/unread-count", { method: "GET" });
  if (!res.ok) {
    throw new Error("Failed to load unread notification count");
  }
  return res.json();
}
