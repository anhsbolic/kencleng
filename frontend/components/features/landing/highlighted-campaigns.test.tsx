import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { describe, expect, it } from "vitest";
import { server } from "@/mocks/server";
import { HighlightedCampaigns } from "./highlighted-campaigns";

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

describe("HighlightedCampaigns", () => {
  it("renders skeleton cards while loading, not a bare spinner (R6)", () => {
    render(withQueryClient(<HighlightedCampaigns />));

    // Default handler resolves asynchronously — this assertion runs
    // on the synchronous initial render, before that resolution.
    expect(screen.getAllByRole("status", { name: /memuat kampanye/i })).toHaveLength(3);
  });

  it("renders the fixture campaigns with no organization name or badge (R7)", async () => {
    render(withQueryClient(<HighlightedCampaigns />));

    expect(
      await screen.findByRole("heading", {
        name: "Air bersih untuk 240 keluarga di Dusun Sukamaju",
      })
    ).toBeInTheDocument();
    expect(screen.queryByText(/verified/i)).not.toBeInTheDocument();
    expect(screen.queryByText("Yayasan Peduli Sukamaju")).not.toBeInTheDocument();
  });

  it("shows an empty state with no CTA when the list is empty (R8)", async () => {
    server.use(
      http.get("/campaigns", () =>
        HttpResponse.json({ data: [], pagination: { next_cursor: null, has_more: false } })
      )
    );
    render(withQueryClient(<HighlightedCampaigns />));

    expect(await screen.findByText(/belum ada kampanye/i)).toBeInTheDocument();
    expect(screen.queryByRole("button")).not.toBeInTheDocument();
    expect(screen.queryByRole("link")).not.toBeInTheDocument();
  });

  it("shows a generic error banner with a retry action, never raw error text (R9)", async () => {
    server.use(http.get("/campaigns", () => HttpResponse.error()));
    render(withQueryClient(<HighlightedCampaigns />));

    expect(await screen.findByRole("alert")).toHaveTextContent(/gagal memuat kampanye/i);
    expect(screen.getByRole("button", { name: /coba lagi/i })).toBeInTheDocument();
    expect(screen.queryByText(/failed to fetch/i)).not.toBeInTheDocument();
  });

  it("retries the query when the retry button is clicked", async () => {
    let attempt = 0;
    server.use(
      http.get("/campaigns", () => {
        attempt += 1;
        if (attempt === 1) return HttpResponse.error();
        return HttpResponse.json({
          data: [],
          pagination: { next_cursor: null, has_more: false },
        });
      })
    );
    render(withQueryClient(<HighlightedCampaigns />));

    const retry = await screen.findByRole("button", { name: /coba lagi/i });
    retry.click();

    await waitFor(() => expect(screen.getByText(/belum ada kampanye/i)).toBeInTheDocument());
  });
});
