import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { server } from "@/mocks/server";
import { useAuthStore } from "@/lib/stores/auth-store";
import { AuthBootstrapProvider } from "./auth-bootstrap-provider";

function renderWithQueryClient() {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false } },
  });
  const utils = render(
    <QueryClientProvider client={queryClient}>
      <AuthBootstrapProvider>
        <div>app content</div>
      </AuthBootstrapProvider>
    </QueryClientProvider>
  );
  return { queryClient, ...utils };
}

beforeEach(() => {
  // Zustand store is a module-level singleton — reset between tests so
  // one test's hydration doesn't leak into the next.
  useAuthStore.setState({ accessToken: null });
});

describe("AuthBootstrapProvider", () => {
  it("calls refresh exactly once on mount when accessToken is null (R8)", async () => {
    let callCount = 0;
    server.use(
      http.post("/auth/refresh", () => {
        callCount += 1;
        return HttpResponse.json({ access_token: "new-token" }, { status: 200 });
      })
    );

    renderWithQueryClient();

    await waitFor(() => expect(callCount).toBe(1));
  });

  it("populates useAuthStore and invalidates the account.me query on success (R9)", async () => {
    server.use(
      http.post("/auth/refresh", () =>
        HttpResponse.json({ access_token: "new-token" }, { status: 200 })
      )
    );

    const { queryClient } = renderWithQueryClient();
    const invalidateSpy = vi.spyOn(queryClient, "invalidateQueries");

    await waitFor(() => expect(useAuthStore.getState().accessToken).toBe("new-token"));
    expect(invalidateSpy).toHaveBeenCalledWith({ queryKey: ["account", "me"] });
  });

  it("leaves accessToken null and renders no error/toast on refresh failure (R10)", async () => {
    server.use(http.post("/auth/refresh", () => HttpResponse.json({}, { status: 401 })));

    const { findByText, queryByRole } = renderWithQueryClient();

    expect(await findByText("app content")).toBeInTheDocument();
    await waitFor(() => expect(useAuthStore.getState().accessToken).toBeNull());
    expect(queryByRole("alert")).not.toBeInTheDocument();
    expect(queryByRole("status")).not.toBeInTheDocument();
  });

  it("does not call refresh if accessToken is already set before mount (R8)", async () => {
    useAuthStore.setState({ accessToken: "already-set" });
    let callCount = 0;
    server.use(
      http.post("/auth/refresh", () => {
        callCount += 1;
        return HttpResponse.json({ access_token: "new-token" }, { status: 200 });
      })
    );

    renderWithQueryClient();
    await new Promise((resolve) => setTimeout(resolve, 0));

    expect(callCount).toBe(0);
  });

  it("does not trigger a second refresh call on re-render within the same mount (R11)", async () => {
    let callCount = 0;
    server.use(
      http.post("/auth/refresh", () => {
        callCount += 1;
        return HttpResponse.json({ access_token: "new-token" }, { status: 200 });
      })
    );

    const queryClient = new QueryClient({ defaultOptions: { queries: { retry: false } } });
    const { rerender } = render(
      <QueryClientProvider client={queryClient}>
        <AuthBootstrapProvider>
          <div>app content</div>
        </AuthBootstrapProvider>
      </QueryClientProvider>
    );

    await waitFor(() => expect(callCount).toBe(1));

    rerender(
      <QueryClientProvider client={queryClient}>
        <AuthBootstrapProvider>
          <div>app content, re-rendered</div>
        </AuthBootstrapProvider>
      </QueryClientProvider>
    );

    await new Promise((resolve) => setTimeout(resolve, 0));
    expect(callCount).toBe(1);
  });

  // R13 (provider must be a descendant of QueryProvider) is verified
  // structurally: every test above renders `AuthBootstrapProvider`
  // inside a `QueryClientProvider` and calls `useQueryClient()`
  // internally (R9's invalidation) — if it weren't correctly nested,
  // every test here would throw on render rather than pass.
});
