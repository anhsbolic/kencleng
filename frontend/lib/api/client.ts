// Centralized API client — every domain's lib/api/<domain>.ts fetch
// function must go through `apiFetch`, never a raw `fetch()` call.
// This is the one place in frontend/ allowed to call `fetch` directly
// (api-client-centralization.md's core checklist item).
//
// Attaches, on every request: the in-memory access token,
// `credentials: 'include'` for the auth cookie to travel, and the
// backend's required CSRF/custom header on state-changing methods.

import { useAuthStore } from "@/lib/stores/auth-store";

function getAccessToken() {
  return useAuthStore.getState().accessToken;
}
function setAccessToken(token: string | null) {
  useAuthStore.getState().setAccessToken(token);
}

type RefreshResponse = {
  access_token: string;
};

/**
 * Calls `POST /auth/refresh` directly (not through `apiFetch`, to
 * avoid recursing into another 401 handler) to rotate the refresh
 * token cookie and obtain a new access token. At most one attempt —
 * never a retry loop. A failed refresh clears the in-memory token so
 * the app can fall back to a logged-out state instead of hanging.
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

export { apiFetch, getAccessToken, setAccessToken };
