"use client";

import { useEffect } from "react";

/**
 * Registers the app-shell service worker (`public/sw.js`) once, on
 * mount, gated strictly to environments that actually support it.
 * Renders nothing — mount once in the root layout.
 */
export function SwRegister() {
  useEffect(() => {
    if (typeof window === "undefined" || !("serviceWorker" in navigator)) {
      return;
    }

    navigator.serviceWorker.register("/sw.js").catch((error) => {
      console.error("Service worker registration failed", error);
    });
  }, []);

  return null;
}
