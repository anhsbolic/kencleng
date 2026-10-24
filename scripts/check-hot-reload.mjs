// Opt-in check for already-running local apps. Source edits are temporary.
import { createRequire } from "node:module";
import { existsSync } from "node:fs";
import { readFile, writeFile, rename, unlink, mkdtemp, rm, stat } from "node:fs/promises";
import { randomUUID } from "node:crypto";
import { fileURLToPath } from "node:url";
import { join } from "node:path";
import { tmpdir } from "node:os";

const root = fileURLToPath(new URL("../", import.meta.url));
const requireFrontend = createRequire(join(root, "frontend/package.json"));
const token = randomUUID().replaceAll("-", "");
const header = "x-kencleng-reload-probe";
const baseURL = (process.env.DEV_FRONTEND_URL || "http://localhost:3000").replace(/\/+$/, "");
const backendURL = (process.env.DEV_BACKEND_URL || "http://127.0.0.1:8090").replace(/\/+$/, "");
const changes = [];
const controller = new AbortController();
let browser;
let backupDir;

function assertLoopback(value) {
  const url = new URL(value);
  if (!['localhost', '127.0.0.1', '[::1]'].includes(url.hostname) || !['http:', 'https:'].includes(url.protocol)) {
    throw new Error("The hot-reload probe only accepts local loopback HTTP(S) URLs.");
  }
  if (url.username || url.password || url.search || url.hash || url.pathname !== '/') {
    throw new Error("Probe URLs must be plain origins without credentials, path, query or fragment.");
  }
}

async function atomicWrite(file, content) {
  const temporary = `${file}.reload-${token}.tmp`;
  const mode = (await stat(file)).mode & 0o777;
  try {
    await writeFile(temporary, content, { mode });
    await rename(temporary, file);
  } finally {
    await unlink(temporary).catch(error => { if (error.code !== 'ENOENT') throw error; });
  }
}

async function health(url) {
  const response = await fetch(url, {
    cache: "no-store",
    signal: AbortSignal.any([controller.signal, AbortSignal.timeout(3000)]),
  });
  if (!response.ok || (await response.json()).status !== "ok") {
    throw new Error(`Health check failed at ${url}`);
  }
  return response;
}

async function waitFor(check, description) {
  const deadline = Date.now() + 90_000;
  let lastError;
  while (Date.now() < deadline) {
    controller.signal.throwIfAborted();
    try {
      if (await check()) return;
    } catch (error) { lastError = error; }
    await new Promise(resolve => setTimeout(resolve, 500));
  }
  throw new Error(`Timed out: ${description}`, { cause: lastError });
}

async function plan() {
  const backendFile = join(root, "backend/cmd/server/main.go");
  const frontendFile = join(root, "frontend/app/page.tsx");
  const backend = await readFile(backendFile, "utf8");
  const frontend = await readFile(frontendFile, "utf8");
  const signature = "func healthz(w http.ResponseWriter, _ *http.Request) {";
  if (!backend.includes(signature) || !/<main\b[^>]*>/.test(frontend)) {
    throw new Error("Probe anchors changed. Inspect the current app before using this check.");
  }
  if (backend.includes("X-Kencleng-Reload-Probe") || frontend.includes("data-kencleng-reload-probe")) {
    throw new Error("Existing probe markers found. Finish/recover the earlier probe first.");
  }
  return [
    { file: backendFile, original: backend, patched: backend.replace(signature,
      `${signature}\n\tw.Header().Set("X-Kencleng-Reload-Probe", "${token}")`) },
    { file: frontendFile, original: frontend, patched: frontend.replace(/<main\b[^>]*>/,
      match => `${match}\n      <span data-kencleng-reload-probe="${token}">${token}</span>`) },
  ];
}

async function restore() {
  const errors = [];
  for (const change of [...changes].reverse()) {
    try {
      const current = await readFile(change.file, "utf8");
      if (current === change.patched) await atomicWrite(change.file, change.original);
      else if (current !== change.original) throw new Error(`Concurrent edit in ${change.file}`);
    } catch (error) { errors.push(error); }
  }
  if (errors.length) {
    throw new AggregateError(errors, `Source recovery needed. Original copies are in ${backupDir}`);
  }
}

async function main() {
  const dryRun = process.argv.includes("--dry-run");
  if (process.argv.slice(2).some(arg => arg !== "--dry-run")) {
    throw new Error("Usage: node scripts/check-hot-reload.mjs [--dry-run]");
  }
  assertLoopback(baseURL);
  assertLoopback(backendURL);
  const planned = await plan();
  if (dryRun) {
    console.log("Preview only; no source edits or HTTP/browser checks:");
    for (const change of planned) console.log(`  temporary probe: ${change.file}`);
    return;
  }
  const onSignal = () => {
    controller.abort(new Error("Probe interrupted; restoring source"));
    browser?.close().catch(() => {});
  };
  process.on("SIGINT", onSignal);
  process.on("SIGTERM", onSignal);
  let passed = false;
  let restored = false;
  try {
    await health(`${backendURL}/healthz`);
    await health(`${baseURL}/api/healthz`);
    const { chromium } = requireFrontend("@playwright/test");
    const executable = process.env.PLAYWRIGHT_CHROMIUM_EXECUTABLE_PATH;
    const options = { headless: true };
    if (executable) options.executablePath = executable;
    else if (existsSync("/usr/bin/google-chrome")) options.executablePath = "/usr/bin/google-chrome";
    browser = await chromium.launch(options);
    const page = await browser.newPage();
    await page.goto(baseURL, { waitUntil: "domcontentloaded" });
    await page.locator("main").waitFor({ state: "visible" });
    backupDir = await mkdtemp(join(tmpdir(), "kencleng-hot-reload-"));
    console.log(`Recovery copies for this probe: ${backupDir}`);
    for (const [index, change] of planned.entries()) {
      if (await readFile(change.file, "utf8") !== change.original) {
        throw new Error(`Source changed during preflight: ${change.file}`);
      }
      await writeFile(join(backupDir, `${index}-${change.file.endsWith('.go') ? 'main.go' : 'page.tsx'}`), change.original, { mode: 0o600 });
      changes.push(change);
      await atomicWrite(change.file, change.patched);
    }
    console.log("Temporary probes installed; waiting for the existing runners to reload...");
    await waitFor(async () => (await health(`${backendURL}/healthz`)).headers.get(header) === token, "backend rebuild/restart");
    console.log("PASS backend: a rebuilt server serves the temporary probe header.");
    await waitFor(async () => (await health(`${baseURL}/api/healthz`)).headers.get(header) === token, "proxy to reloaded backend");
    // No page.reload(): the already-open browser must update through the dev client.
    await page.locator(`[data-kencleng-reload-probe="${token}"]`).waitFor({ state: "visible", timeout: 90_000 });
    console.log("PASS frontend: the already-open browser received the source change automatically.");
    await restore();
    restored = true;
    await page.locator(`[data-kencleng-reload-probe="${token}"]`).waitFor({ state: "detached", timeout: 90_000 });
    await waitFor(async () => !(await health(`${backendURL}/healthz`)).headers.has(header), "backend source restoration");
    passed = true;
    console.log("PASS: both source files restored and the running apps reverted to their original response.");
  } finally {
    try {
      if (!restored) { await restore(); restored = true; }
    } finally {
      await browser?.close();
      process.removeListener("SIGINT", onSignal);
      process.removeListener("SIGTERM", onSignal);
      if (backupDir && restored) await rm(backupDir, { recursive: true });
      else if (backupDir) console.error(`Keep the recovery copies: ${backupDir}`);
      if (!passed) console.error("Hot reload was not verified. Existing runners were not stopped by this check.");
    }
  }
}

main().catch(error => {
  console.error(error.message);
  if (error.cause) console.error(`Cause: ${error.cause.message}`);
  process.exitCode = 1;
});
