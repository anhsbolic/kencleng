import react from "@vitejs/plugin-react";
import path from "node:path";
import { configDefaults, defineConfig } from "vitest/config";

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "."),
    },
  },
  test: {
    environment: "jsdom",
    setupFiles: ["./vitest.setup.ts"],
    globals: true,
    // Playwright owns committed real-browser tests. Keep them out of
    // Vitest even though both runners intentionally use *.spec.ts.
    exclude: [...configDefaults.exclude, "tests/browser/**"],
    coverage: { provider: 'v8', reporter: ['text', 'json-summary'] },
    passWithNoTests: true, // no feature tests at reset baseline; do not claim test PASS for behavior
  },
});
