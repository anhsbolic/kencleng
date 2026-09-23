import { render, screen, waitFor } from "@testing-library/react";
import { StrictMode } from "react";
import { afterEach, describe, expect, it, vi } from "vitest";

const worker = vi.hoisted(() => ({
  start: vi.fn().mockResolvedValue(undefined),
  stop: vi.fn(),
}));

vi.mock("@/mocks/browser", () => ({ worker }));

import MockServiceWorker from "./mock-service-worker";

describe("MockServiceWorker", () => {
  afterEach(() => {
    vi.unstubAllEnvs();
    vi.clearAllMocks();
    worker.start.mockResolvedValue(undefined);
  });

  it("stops interception once when an enabled provider unmounts after startup", async () => {
    vi.stubEnv("NEXT_PUBLIC_MSW_ENABLED", "true");
    const { unmount } = render(
      <MockServiceWorker>
        <p>Campaign detail</p>
      </MockServiceWorker>,
    );

    await waitFor(() => expect(worker.start).toHaveBeenCalledOnce());
    unmount();

    expect(worker.stop).toHaveBeenCalledOnce();
  });

  it("stops a worker that finishes starting after its provider unmounts", async () => {
    vi.stubEnv("NEXT_PUBLIC_MSW_ENABLED", "true");
    let finishStart: (() => void) | undefined;
    worker.start.mockImplementationOnce(
      () =>
        new Promise<void>((resolve) => {
          finishStart = resolve;
        }),
    );

    const { unmount } = render(
      <MockServiceWorker>
        <p>Campaign detail</p>
      </MockServiceWorker>,
    );

    await waitFor(() => expect(worker.start).toHaveBeenCalledOnce());
    unmount();
    finishStart?.();

    await waitFor(() => expect(worker.stop).toHaveBeenCalledOnce());
  });

  it("does not start interception when unmounted before the worker import resolves", async () => {
    vi.stubEnv("NEXT_PUBLIC_MSW_ENABLED", "true");
    const { unmount } = render(
      <MockServiceWorker>
        <p>Campaign detail</p>
      </MockServiceWorker>,
    );

    unmount();
    await new Promise((resolve) => setTimeout(resolve, 0));

    expect(worker.start).not.toHaveBeenCalled();
    expect(worker.stop).not.toHaveBeenCalled();
  });

  it("keeps interception running for the mounted Strict Mode owner after stale startup completes", async () => {
    vi.stubEnv("NEXT_PUBLIC_MSW_ENABLED", "true");
    let finishStart: (() => void) | undefined;
    worker.start.mockImplementationOnce(
      () =>
        new Promise<void>((resolve) => {
          finishStart = resolve;
        }),
    );

    const { unmount } = render(
      <StrictMode>
        <MockServiceWorker>
          <p>Campaign detail</p>
        </MockServiceWorker>
      </StrictMode>,
    );

    await waitFor(() => expect(worker.start).toHaveBeenCalledOnce());
    finishStart?.();

    expect(await screen.findByText("Campaign detail")).toBeVisible();
    expect(worker.stop).not.toHaveBeenCalled();

    unmount();
    expect(worker.stop).toHaveBeenCalledOnce();
  });
});
