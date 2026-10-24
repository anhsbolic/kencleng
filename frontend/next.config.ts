import type { NextConfig } from "next";
import { PHASE_DEVELOPMENT_SERVER } from "next/constants";

export default function nextConfig(phase: string): NextConfig {
  if (phase !== PHASE_DEVELOPMENT_SERVER) {
    return {};
  }

  const backend = (
    process.env.BACKEND_INTERNAL_URL || "http://127.0.0.1:8090"
  ).replace(/\/+$/, "");

  return {
    async rewrites() {
      return {
        beforeFiles: [
          { source: "/api/:path*", destination: `${backend}/:path*` },
        ],
      };
    },
  };
}
