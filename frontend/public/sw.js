// Hand-written app-shell service worker (no Workbox/next-pwa — manual
// PWA setup per README.md and kencleng-frontend-tech-stack.md's PWA
// Scope: app-shell caching only, no offline write queue, no data
// caching).
//
// Strategy:
// - App-shell entry points (below) are precached on install and
//   served cache-first, so the shell loads even with no connection.
// - Hashed, immutable Next.js build assets (`/_next/static/*`) are
//   cached lazily on first fetch, then served cache-first — safe
//   because the filename changes whenever the content does.
// - Everything else (API calls, dynamic pages/data) is network-only.
//   This app never caches API responses — that's deliberately out of
//   scope, not an oversight.

const CACHE_NAME = "kencleng-shell-v1";

const APP_SHELL_URLS = ["/", "/manifest.json", "/favicon.ico"];

self.addEventListener("install", (event) => {
  event.waitUntil(
    caches
      .open(CACHE_NAME)
      .then((cache) => cache.addAll(APP_SHELL_URLS))
      .then(() => self.skipWaiting())
  );
});

self.addEventListener("activate", (event) => {
  event.waitUntil(
    caches
      .keys()
      .then((keys) =>
        Promise.all(
          keys
            .filter((key) => key !== CACHE_NAME)
            .map((key) => caches.delete(key))
        )
      )
      .then(() => self.clients.claim())
  );
});

self.addEventListener("fetch", (event) => {
  const { request } = event;
  const url = new URL(request.url);

  if (request.method !== "GET" || url.origin !== self.location.origin) {
    return; // network-only: mutating requests and cross-origin requests
  }

  const isImmutableBuildAsset = url.pathname.startsWith("/_next/static/");
  const isAppShellUrl = APP_SHELL_URLS.includes(url.pathname);

  if (isImmutableBuildAsset || isAppShellUrl) {
    event.respondWith(
      caches.open(CACHE_NAME).then(async (cache) => {
        const cached = await cache.match(request);
        if (cached) return cached;

        const response = await fetch(request);
        if (response.ok) cache.put(request, response.clone());
        return response;
      })
    );
    return;
  }

  // Everything else (API routes, non-shell pages, dynamic data):
  // network-only, no caching — left to the browser's default handling.
});
