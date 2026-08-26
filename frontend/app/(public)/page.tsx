import { HighlightedCampaigns } from "@/components/features/landing/highlighted-campaigns";
import { Hero } from "@/components/features/landing/hero";
import { HowItWorks } from "@/components/features/landing/how-it-works";

/**
 * `/` — Guest landing page. Server Component; only
 * `HighlightedCampaigns` (the campaign-fetching section) is a
 * `'use client'` leaf (techplan Decision 6). `TrustStrip` and the
 * footer are deliberately not built here — see
 * `.local-agents/works/00-shell-landing/2-plan/techplan.md` §2 Scope.
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
