import { expect, test, type Page } from "@playwright/test";

const representativeViewports = [
  { name: "desktop", width: 1280, height: 800 },
  { name: "mobile", width: 390, height: 844 },
] as const;

async function expectLoginSurface(page: Page) {
  await expect(page.getByRole("heading", { name: "Masuk", level: 1 })).toBeVisible();
  await expect(page.getByLabel("Email")).toBeVisible();
  await expect(page.getByLabel("Password")).toBeVisible();
  await expect(page.getByRole("button", { name: "Masuk", exact: true })).toBeVisible();
  await expect(page.getByRole("link", { name: "Masuk dengan Google" })).toBeVisible();

  const hasHorizontalOverflow = await page.evaluate(
    () => document.documentElement.scrollWidth > document.documentElement.clientWidth
  );
  expect(hasHorizontalOverflow).toBe(false);
}

test("login route renders its core auth surface at representative widths", async ({ page }) => {
  const pageErrors: Error[] = [];
  page.on("pageerror", (error) => pageErrors.push(error));

  for (const viewport of representativeViewports) {
    await test.step(viewport.name, async () => {
      await page.setViewportSize({ width: viewport.width, height: viewport.height });
      await page.goto("/login");
      await expectLoginSurface(page);
    });
  }

  expect(pageErrors).toEqual([]);
});
