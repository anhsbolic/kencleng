import { expect, test } from "/home/anhar-solehudin/kencleng-workspace/kencleng/frontend/node_modules/@playwright/test";

test("moves fragment credential to status header and shows status only", async ({ page }) => {
  const donationId = "0d000000-0000-4000-8000-000000000123";
  const credential = `browser-proof-${donationId}`;
  const statusRequest = page.waitForRequest((request) =>
    request.url().includes(`/api/donations/${donationId}/status`),
  );

  await page.route(`**/api/donations/${donationId}/status`, async (route) => {
    await route.fulfill({
      status: 200,
      contentType: "application/json",
      headers: { "Cache-Control": "private, no-store" },
      body: JSON.stringify({ status: "pending" }),
    });
  });

  await page.goto(`/donations/${donationId}/status#${encodeURIComponent(credential)}`);
  await expect(page.getByRole("heading", { name: "Menunggu hasil simulasi" })).toBeVisible();
  await expect.poll(() => new URL(page.url()).hash).toBe("");

  const request = await statusRequest;
  expect(request.headers()["x-donation-status-credential"]).toBe(credential);
  expect(page.url()).toBe(`http://127.0.0.1:3000/donations/${donationId}/status`);
  await expect(page.getByText(credential, { exact: false })).toHaveCount(0);
  await expect(page.getByText(donationId, { exact: false })).toHaveCount(0);
  await expect(page.getByRole("button", { name: "Periksa status lagi" })).toBeVisible();
});
