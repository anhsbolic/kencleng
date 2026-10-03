import baseConfig from "/home/anhar-solehudin/kencleng-workspace/kencleng/frontend/playwright.config";
import { defineConfig } from "/home/anhar-solehudin/kencleng-workspace/kencleng/frontend/node_modules/@playwright/test";

const runPath = "/home/anhar-solehudin/kencleng-workspace/kencleng/.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TST-S2-004-001";
const frontendPath = "/home/anhar-solehudin/kencleng-workspace/kencleng/frontend";

export default defineConfig({
  ...baseConfig,
  testDir: `${runPath}/browser`,
  outputDir: `${runPath}/browser-output`,
  webServer: {
    ...baseConfig.webServer,
    command: "npm run dev -- --hostname 127.0.0.1",
    cwd: frontendPath,
  },
});
