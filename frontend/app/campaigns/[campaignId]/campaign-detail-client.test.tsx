import { fireEvent, render, screen } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import { describe, expect, it } from "vitest";
import { campaignFixtureIds, publicCampaignFixtures } from "@/mocks/fixtures/public-campaign";
import { server } from "@/mocks/server";
import CampaignDetailClient from "./campaign-detail-client";

describe("CampaignDetailClient", () => {
  it("renders the cap before the donation entry and keeps the entry tied to the contract action", async () => {
    render(<CampaignDetailClient campaignId={campaignFixtureIds.available} />);

    expect(await screen.findByRole("heading", { name: /Dapur bersama/i })).toBeVisible();
    expect(screen.getByText("Rp 1.250.000,00")).toBeVisible();
    expect(screen.getByText("25%")).toBeVisible();
    expect(screen.getAllByText("Dari pengelola").length).toBeGreaterThan(0);
    expect(screen.getByRole("img", { name: /Peralatan memasak/i })).toHaveAttribute(
      "src",
      expect.stringContaining(`/api/campaigns/${campaignFixtureIds.available}/media/`),
    );
    expect(screen.getByText("Rp 1.000.000.000,00 IDR")).toBeVisible();
    expect(screen.getByRole("link", { name: "Lanjutkan donasi" })).toHaveAttribute(
      "href",
      `/campaigns/${campaignFixtureIds.available}/donate`,
    );
    expect(screen.getByRole("main")).not.toHaveFocus();
  });

  it("keeps explicit IDR cap disclosure when Funding is unavailable", async () => {
    render(<CampaignDetailClient campaignId={campaignFixtureIds.fundingUnavailable} />);

    expect(await screen.findByText("Rp 1.000.000.000,00 IDR")).toBeVisible();
    expect(screen.getByRole("link", { name: "Lanjutkan donasi" })).toBeVisible();
  });

  it("does not offer donation entry when the accepted action is unavailable", async () => {
    render(<CampaignDetailClient campaignId={campaignFixtureIds.donationUnavailable} />);

    expect(await screen.findByRole("heading", { name: /Dukungan belum dapat dilakukan/i })).toBeVisible();
    expect(screen.queryByRole("link", { name: "Lanjutkan donasi" })).not.toBeInTheDocument();
  });

  it("shows one safe non-disclosing not-found state", async () => {
    render(<CampaignDetailClient campaignId="not-found" />);

    expect(
      await screen.findByRole("heading", { name: "Campaign ini tidak dapat ditampilkan." }),
    ).toBeVisible();
    expect(screen.queryByRole("button", { name: /coba lagi/i })).not.toBeInTheDocument();
    expect(screen.queryByText(/malformed|non-public|diagnostic/i)).not.toBeInTheDocument();
  });

  it("recovers a generic request failure through retry and restores focus to campaign content", async () => {
    server.use(http.get("/api/campaigns/network-failure", () => HttpResponse.error()));
    render(<CampaignDetailClient campaignId="network-failure" />);

    expect(
      await screen.findByRole("heading", { name: "Detail campaign belum dapat dimuat." }),
    ).toBeVisible();

    const retry = screen.getByRole("button", { name: "Coba lagi" });
    retry.focus();
    expect(retry).toHaveFocus();

    server.use(
      http.get("/api/campaigns/network-failure", () =>
        HttpResponse.json(publicCampaignFixtures[campaignFixtureIds.available]),
      ),
    );
    fireEvent.click(retry);

    expect(await screen.findByRole("heading", { name: /Dapur bersama/i })).toBeVisible();
    expect(screen.getByRole("main")).toHaveFocus();
  });

  it("announces unavailable state and moves focus to campaign content after a successful retry", async () => {
    render(<CampaignDetailClient campaignId="unavailable" />);

    expect(
      await screen.findByRole("heading", { name: "Detail campaign sedang tidak tersedia." }),
    ).toBeVisible();
    expect(screen.getByRole("status")).toHaveAttribute("aria-live", "polite");

    server.use(
      http.get("/api/campaigns/unavailable", () =>
        HttpResponse.json(publicCampaignFixtures[campaignFixtureIds.available]),
      ),
    );
    fireEvent.click(screen.getByRole("button", { name: "Coba lagi" }));

    expect(await screen.findByRole("heading", { name: /Dapur bersama/i })).toBeVisible();
    expect(screen.getByRole("main")).toHaveFocus();
  });

  it("renders hostile-looking organizer text literally and distinguishes unavailable media", async () => {
    const { container, rerender } = render(<CampaignDetailClient campaignId={campaignFixtureIds.hostileOrganizer} />);

    expect(await screen.findByText("<strong>Bukan HTML</strong>; ini tetap cerita dari pengelola.")).toBeVisible();
    expect(container.querySelector("strong")).not.toBeInTheDocument();

    rerender(<CampaignDetailClient campaignId={campaignFixtureIds.unavailableMedia} />);
    expect(
      await screen.findByRole("heading", { name: "Media sementara belum tersedia." }),
    ).toBeVisible();
    expect(screen.getByText(/berbeda dari belum ada media/i)).toBeVisible();
  });
});
