import { expect, test, type Page } from "@playwright/test";

const representativeViewports = [
  { name: "desktop", width: 1280, height: 800 },
  { name: "mobile", width: 390, height: 844 },
] as const;

const HOME_READINESS_TIMEOUT_MS = 15_000;
const populatedCampaignHeading = {
  name: "Air bersih untuk 240 keluarga di Dusun Sukamaju",
  level: 3,
} as const;

async function waitForPopulatedHome(page: Page, context: string) {
  try {
    await page
      .getByRole("heading", populatedCampaignHeading)
      .waitFor({ state: "visible", timeout: HOME_READINESS_TIMEOUT_MS });
  } catch (error) {
    const diagnostics = await page
      .evaluate(() => ({
        readyState: document.readyState,
        visibleText: document.body?.innerText.slice(0, 800) ?? "",
      }))
      .catch((diagnosticError: unknown) => ({
        diagnosticsUnavailable:
          diagnosticError instanceof Error
            ? diagnosticError.message
            : String(diagnosticError),
      }));
    const loadingStatusCount = await page
      .getByRole("status", { name: "Memuat kampanye" })
      .count()
      .catch(() => -1);
    const alertText = await page
      .getByRole("alert")
      .allTextContents()
      .catch(() => []);

    throw new Error(
      [
        `Home route did not reach its populated state within ${HOME_READINESS_TIMEOUT_MS}ms (${context}).`,
        `URL: ${page.url()}`,
        `Loading status count: ${loadingStatusCount}`,
        `Alerts: ${JSON.stringify(alertText)}`,
        `Document: ${JSON.stringify(diagnostics)}`,
        `Playwright wait failure: ${error instanceof Error ? error.message : String(error)}`,
      ].join("\n")
    );
  }
}

async function expectHomeSurface(page: Page) {
  await expect(page.locator("html")).toHaveAttribute("lang", "id");
  await expect(
    page.getByRole("heading", { name: "Berbagi itu mudah, dampaknya nyata", level: 1 })
  ).toBeVisible();
  await expect(page.getByRole("link", { name: "Mulai berdonasi" })).toBeVisible();
  await expect(
    page.getByRole("heading", { name: "Sedang berjalan minggu ini", level: 2 })
  ).toBeVisible();
  await expect(page.getByRole("heading", populatedCampaignHeading)).toBeVisible();
  await expect(
    page.getByRole("heading", { name: "Tiga langkah untuk berdonasi", level: 2 })
  ).toBeVisible();

  const hasHorizontalOverflow = await page.evaluate(
    () => document.documentElement.scrollWidth > document.documentElement.clientWidth
  );
  expect(hasHorizontalOverflow).toBe(false);
}

for (const viewport of representativeViewports) {
  test(`home route renders its populated core surface at ${viewport.name} width`, async ({
    page,
  }) => {
    const pageErrors: Error[] = [];
    page.on("pageerror", (error) => pageErrors.push(error));

    await page.setViewportSize({ width: viewport.width, height: viewport.height });
    await page.goto("/");
    await waitForPopulatedHome(page, `${viewport.name} populated route`);
    await expectHomeSurface(page);

    await page.getByRole("link", { name: "Mulai berdonasi" }).click();
    await expect(page).toHaveURL(/#kampanye$/);
    await expect(page.locator("#kampanye")).toBeInViewport();

    expect(pageErrors).toEqual([]);
  });
}

test("mobile drawer is reachable and returns focus on Escape", async ({ page }) => {
  await page.setViewportSize({ width: 390, height: 844 });
  await page.goto("/");
  await waitForPopulatedHome(page, "mobile drawer route");

  const hamburger = page.getByRole("button", { name: "Buka menu" });
  await hamburger.click();

  const drawer = page.getByRole("dialog", { name: "Navigasi" });
  await expect(drawer).toBeVisible();
  await expect(drawer.getByRole("link", { name: "Beranda" })).toBeFocused();
  await expect(drawer.getByRole("link", { name: "Jelajahi Kampanye" })).toBeVisible();
  await expect(drawer.getByRole("button", { name: "Masuk" })).toBeVisible();
  await expect(drawer.getByRole("button", { name: "Daftar" })).toBeVisible();

  await page.keyboard.press("Escape");
  await expect(drawer).toBeHidden();
  await expect(hamburger).toBeFocused();
});
