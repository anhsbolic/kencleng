import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { act, render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { describe, expect, it } from "vitest";
import type { CampaignListResponse } from "@/lib/api/campaign";
import { campaignKeys } from "@/lib/hooks/use-campaigns";
import { server } from "@/mocks/server";
import { HighlightedCampaigns } from "./highlighted-campaigns";

function createQueryClient() {
  return new QueryClient({
    defaultOptions: { queries: { retry: false } },
  });
}

function withQueryClient(children: ReactNode, queryClient = createQueryClient()) {
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

describe("HighlightedCampaigns", () => {
  it("renders shaped skeleton cards while loading", () => {
    render(withQueryClient(<HighlightedCampaigns />));
    expect(screen.getAllByRole("status", { name: /memuat kampanye/i })).toHaveLength(3);
  });

  it("renders fixture campaigns without invented organization detail", async () => {
    render(withQueryClient(<HighlightedCampaigns />));

    expect(
      await screen.findByRole("heading", {
        name: "Air bersih untuk 240 keluarga di Dusun Sukamaju",
      })
    ).toBeInTheDocument();
    expect(screen.queryByText(/verified/i)).not.toBeInTheDocument();
    expect(screen.queryByText("Yayasan Peduli Sukamaju")).not.toBeInTheDocument();
  });

  it("keeps populated cards visible while a background refetch is pending", async () => {
    const queryClient = createQueryClient();
    render(withQueryClient(<HighlightedCampaigns />, queryClient));

    const firstCampaign = await screen.findByRole("heading", {
      name: "Air bersih untuk 240 keluarga di Dusun Sukamaju",
    });
    const cachedData = queryClient.getQueryData<CampaignListResponse>(
      campaignKeys.list()
    );
    if (!cachedData) throw new Error("Expected the populated campaign fixture in cache");

    let markRefetchStarted!: () => void;
    const refetchStarted = new Promise<void>((resolve) => {
      markRefetchStarted = resolve;
    });
    let releaseRefetch!: () => void;
    const refetchGate = new Promise<void>((resolve) => {
      releaseRefetch = resolve;
    });
    server.use(
      http.get("/campaigns", async () => {
        markRefetchStarted();
        await refetchGate;
        return HttpResponse.json(cachedData);
      })
    );

    let refetchPromise = Promise.resolve();
    act(() => {
      refetchPromise = queryClient.invalidateQueries({ queryKey: campaignKeys.list() });
    });
    await refetchStarted;

    expect(await screen.findByText("Data mungkin tidak terbaru.")).toBeInTheDocument();
    expect(firstCampaign).toBeInTheDocument();
    expect(
      screen.getByRole("heading", { name: "Renovasi musala dan ruang belajar anak" })
    ).toBeInTheDocument();

    await act(async () => {
      releaseRefetch();
      await refetchPromise;
    });
    await waitFor(() =>
      expect(screen.queryByText("Data mungkin tidak terbaru.")).not.toBeInTheDocument()
    );
  });

  it("shows an empty state with no CTA", async () => {
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

  it("shows a generic retryable error without raw error text", async () => {
    server.use(http.get("/campaigns", () => HttpResponse.error()));
    render(withQueryClient(<HighlightedCampaigns />));

    expect(await screen.findByRole("alert")).toHaveTextContent(/gagal memuat kampanye/i);
    expect(screen.getByRole("button", { name: /coba lagi/i })).toBeInTheDocument();
    expect(screen.queryByText(/failed to fetch/i)).not.toBeInTheDocument();
  });

  it("retries the query when requested", async () => {
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
