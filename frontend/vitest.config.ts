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
    // No component tests exist at the clean baseline. Tooling verification
    // should still pass before the first production behavior is introduced.
    passWithNoTests: true,
  },
});
