// Centralized API client — every domain's lib/api/<domain>.ts fetch
// function must go through `apiFetch`, never a raw `fetch()` call.
// This is the one place in frontend/ allowed to call `fetch` directly
// (api-client-centralization.md's core checklist item).
//
// Attaches, on every request: the in-memory access token,
// `credentials: 'include'` for the auth cookie to travel, and the
// backend's required CSRF/custom header on state-changing methods.

import { useAuthStore } from "@/lib/stores/auth-store";
import type { components } from "./schema";

function getAccessToken() {
  return useAuthStore.getState().accessToken;
}
function setAccessToken(token: string | null) {
  useAuthStore.getState().setAccessToken(token);
}

// Generated schema type (also carries `access_token_expires_at`) —
// not hand-written, per `frontend/AGENTS.md` §3 (techplan account/
// 02-google-oauth-login-register, R12).
type RefreshResponse = components["schemas"]["RefreshResponse"];

/**
 * Calls `POST /auth/refresh` directly (not through `apiFetch`, to
 * avoid recursing into another 401 handler) to rotate the refresh
 * token cookie and obtain a new access token. At most one attempt —
 * never a retry loop. A failed refresh clears the in-memory token so
 * the app can fall back to a logged-out state instead of hanging.
 *
 * Exported (not just used internally by `apiFetch`'s 401 handler) so
 * `AuthBootstrapProvider` can call the exact same logic on app mount
 * to hydrate `useAuthStore` from whatever refresh cookie is present —
 * including the one set by a successful Google OAuth callback, which
 * delivers its tokens as cookies rather than a JSON body (techplan
 * account/02-google-oauth-login-register, D3/R8-R11).
 */
async function tryRefreshOnce(): Promise<boolean> {
  try {
    const res = await fetch("/auth/refresh", {
      method: "POST",
      credentials: "include",
      headers: { "X-Requested-With": "kencleng-frontend" },
    });

    if (!res.ok) {
      setAccessToken(null);
      return false;
    }

    const body: RefreshResponse = await res.json();
    setAccessToken(body.access_token);
    return true;
  } catch {
    setAccessToken(null);
    return false;
  }
}

/**
 * Low-level fetch wrapper every `lib/api/<domain>.ts` function calls
 * through. Not exported as the public per-domain API — domain fetch
 * functions wrap this with typed request/response shapes from
 * `lib/api/schema.d.ts`.
 */
async function apiFetch(
  path: string,
  init: RequestInit = {},
  isRetry = false
): Promise<Response> {
  const token = getAccessToken();
  const isMutating = Boolean(init.method) && init.method !== "GET";

  const res = await fetch(path, {
    ...init,
    credentials: "include",
    headers: {
      ...init.headers,
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...(isMutating ? { "X-Requested-With": "kencleng-frontend" } : {}),
    },
  });

  // At most one refresh + retry per call, guarded by `isRetry` — a
  // 401 on the retried request is returned as-is, never re-refreshed.
  if (res.status === 401 && !isRetry) {
    const refreshed = await tryRefreshOnce();
    if (refreshed) {
      return apiFetch(path, init, true);
    }
  }

  return res;
}

/**
 * Thrown by `lib/api/<domain>.ts` functions on a non-OK response that
 * isn't handled as its own typed success/validation branch by the
 * caller (e.g. verify-email's 404/410/429, or any endpoint's network/
 * unexpected-5xx failure). Carries the HTTP status and, where the
 * backend returned an RFC 9457 Problem Details body, its `detail`
 * string — callers that need to distinguish specific documented
 * outcomes (e.g. verify-email's 410 vs 404) read `.status`; anything
 * else falls back to a generic frontend-owned message, per
 * `loading-empty-error-state-conventions.md`.
 */
export class ApiError extends Error {
  status: number;
  detail?: string;

  constructor(status: number, detail?: string) {
    super(detail ?? `Request failed with status ${status}`);
    this.name = "ApiError";
    this.status = status;
    this.detail = detail;
  }
}

/**
 * Best-effort read of an RFC 9457 Problem Details body's `detail`
 * field. Never throws — a malformed/non-JSON error body just means
 * `detail` comes back `undefined`, falling through to `ApiError`'s own
 * generic fallback message rather than surfacing raw parse errors.
 */
export async function readProblemDetail(res: Response): Promise<string | undefined> {
  try {
    const body: unknown = await res.json();
    const detail = (body as { detail?: unknown } | null)?.detail;
    return typeof detail === "string" ? detail : undefined;
  } catch {
    return undefined;
  }
}

export { apiFetch, getAccessToken, setAccessToken, tryRefreshOnce };
