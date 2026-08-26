import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { CampaignCard } from "./campaign-card";

describe("CampaignCard", () => {
  it("renders the title and percentage, with no organization name or badge (R7)", () => {
    render(
      <CampaignCard
        id="1"
        title="Air bersih untuk 240 keluarga"
        progress={{ percentage: 68, donorCount: 84, daysRemaining: 12 }}
      />
    );

    expect(screen.getByRole("heading", { name: "Air bersih untuk 240 keluarga" })).toBeInTheDocument();
    expect(screen.getByText("68% terkumpul")).toBeInTheDocument();
    expect(screen.getByText(/12 hari lagi/)).toBeInTheDocument();
    expect(screen.queryByText(/verified/i)).not.toBeInTheDocument();
  });

  it("shows a 'goal reached' text label at 100%, not color alone (R11)", () => {
    render(
      <CampaignCard
        id="2"
        title="Target tercapai campaign"
        progress={{ percentage: 100, donorCount: 200, daysRemaining: 0 }}
      />
    );

    expect(screen.getByText("Target tercapai")).toBeInTheDocument();
  });

  it("omits the days-remaining segment entirely when null (R12)", () => {
    render(
      <CampaignCard
        id="3"
        title="Closed-adjacent campaign"
        progress={{ percentage: 50, donorCount: 10, daysRemaining: null }}
      />
    );

    expect(screen.queryByText(/hari lagi/)).not.toBeInTheDocument();
    expect(screen.queryByText(/null/)).not.toBeInTheDocument();
  });

  it("has no 'browse files' or upload affordance in the photo placeholder (R14)", () => {
    render(
      <CampaignCard
        id="4"
        title="Placeholder check"
        progress={{ percentage: 10, donorCount: 1, daysRemaining: 30 }}
      />
    );

    expect(screen.queryByText(/browse files/i)).not.toBeInTheDocument();
  });
});
