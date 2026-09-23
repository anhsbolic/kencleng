import type { ReactNode } from "react";
import MockServiceWorker from "./mock-service-worker";

export default function CampaignDetailLayout({ children }: { children: ReactNode }) {
  return <MockServiceWorker>{children}</MockServiceWorker>;
}
