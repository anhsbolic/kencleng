// Cross-tab session-event channel (techplan account/03-login-session-
// management, task-02, D3 — resolves task #02's own carried-forward
// Open Item #1). `BroadcastChannel` is a browser-only global — never
// constructed at module top-level, both to avoid crashing Next.js SSR
// and because this project's own test environment (jsdom 30.0.1)
// doesn't implement it either (confirmed directly against the pinned
// version — see the source techplan's §14 Resolved #6). An eager
// top-level construction would break every test that imports this
// module, not just SSR.
//
// This channel is the fan-out half of cross-tab coordination — the
// mutual-exclusion half (preventing two tabs from concurrently calling
// `POST /auth/refresh`, which the backend's reuse-detection would
// otherwise punish by revoking the whole session) is `client.ts`'s
// `coordinatedRefresh`, via the Web Locks API. Losing this channel
// (unsupported browser) degrades to "no cross-tab awareness" — every
// tab still works correctly on its own, it just won't hear about
// another tab's refresh/logout until its own next 401. It is not the
// correctness guarantee by itself.

const CHANNEL_NAME = "kencleng-auth";

export type AuthChannelMessage =
  | { type: "refreshed"; accessToken: string; accessTokenExpiresAt?: string }
  | { type: "refresh-failed" }
  | { type: "logged-out" };

let channel: BroadcastChannel | null = null;

function getChannel(): BroadcastChannel | null {
  if (channel) return channel;
  if (typeof BroadcastChannel === "undefined") return null;
  channel = new BroadcastChannel(CHANNEL_NAME);
  return channel;
}

/**
 * Fans out a session-lifecycle event to every *other* tab sharing this
 * origin — `BroadcastChannel` never delivers a message back to the tab
 * that sent it, by spec. No-ops silently if `BroadcastChannel` isn't
 * available.
 */
export function postAuthChannelMessage(msg: AuthChannelMessage): void {
  getChannel()?.postMessage(msg);
}

/**
 * Subscribes to session-lifecycle events broadcast by any other tab.
 * Returns an unsubscribe function; a no-op unsubscribe if
 * `BroadcastChannel` isn't available (so callers don't need to
 * feature-detect themselves).
 */
export function subscribeAuthChannel(
  handler: (msg: AuthChannelMessage) => void
): () => void {
  const ch = getChannel();
  if (!ch) return () => {};

  const listener = (event: MessageEvent<AuthChannelMessage>) => handler(event.data);
  ch.addEventListener("message", listener);
  return () => ch.removeEventListener("message", listener);
}

/**
 * Test-only: forces the next `postAuthChannelMessage`/
 * `subscribeAuthChannel` call to re-run feature detection and construct
 * a fresh channel, instead of reusing whatever was cached by an earlier
 * test in the same file/process — needed because this module caches a
 * real `BroadcastChannel` instance once created, and jsdom doesn't
 * provide the global by default (tests that want the "available" path
 * must attach a fake `BroadcastChannel` themselves and reset the cache).
 */
export function __resetAuthChannelForTests(): void {
  channel = null;
}
