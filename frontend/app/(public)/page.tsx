import { Hero } from "./_components/hero";
import { HighlightedCampaigns } from "./_components/highlighted-campaigns";
import { HowItWorks } from "./_components/how-it-works";

/**
 * `/` — Guest landing page. Server Component; only
 * `HighlightedCampaigns` (the campaign-fetching section) is a
 * `'use client'` leaf. This deliberately remains the approved minimum
 * representative slice rather than a complete landing page.
 */
export default function Home() {
  return (
    <>
      <Hero />
      <HighlightedCampaigns />
      <HowItWorks />
    </>
  );
}
