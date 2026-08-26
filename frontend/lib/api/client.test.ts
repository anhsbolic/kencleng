import { HttpResponse, http } from "msw";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { server } from "@/mocks/server";
import { useAuthStore } from "@/lib/stores/auth-store";
import * as authChannel from "./auth-channel";
import { coordinatedRefresh } from "./client";

vi.mock("./auth-channel", async () => {
  const actual = await vi.importActual<typeof import("./auth-channel")>("./auth-channel");
  return { ...actual, postAuthChannelMessage: vi.fn() };
});

beforeEach(() => {
  useAuthStore.setState({ accessToken: null });
  vi.mocked(authChannel.postAuthChannelMessage).mockClear();
});

describe("coordinatedRefresh", () => {
  it("falls back to calling refresh directly when navigator.locks is unavailable (R12)", async () => {
    // jsdom 30.0.1 genuinely has no navigator.locks — confirmed directly
    // against this project's pinned version (source techplan §14
    // Resolved #6), not assumed. This test's own environment is the
    // proof, not a mock standing in for it.
    expect("locks" in navigator).toBe(false);

    server.use(
      http.post("/auth/refresh", () =>
        HttpResponse.json({ access_token: "refreshed-token" }, { status: 200 })
      )
    );

    const ok = await coordinatedRefresh();

    expect(ok).toBe(true);
    expect(useAuthStore.getState().accessToken).toBe("refreshed-token");
  });

  describe("with navigator.locks available (fake, since jsdom doesn't provide the real API)", () => {
    let requestSpy: ReturnType<typeof vi.fn>;

    beforeEach(() => {
      requestSpy = vi.fn((_name: string, cb: () => Promise<boolean>) => cb());
      Object.defineProperty(navigator, "locks", {
        value: { request: requestSpy },
        configurable: true,
      });
    });

    afterEach(() => {
      // Remove the fake so other tests see jsdom's real, unmodified
      // absence of the API again.
      Reflect.deleteProperty(navigator, "locks");
    });

    it("acquires the named lock before calling the underlying refresh (R11)", async () => {
      server.use(
        http.post("/auth/refresh", () =>
          HttpResponse.json({ access_token: "refreshed-token" }, { status: 200 })
        )
      );

      await coordinatedRefresh();

      expect(requestSpy).toHaveBeenCalledWith("kencleng-refresh-token", expect.any(Function));
    });

    it("broadcasts 'refreshed' with the new token on success (R13)", async () => {
      server.use(
        http.post("/auth/refresh", () =>
          HttpResponse.json({ access_token: "refreshed-token" }, { status: 200 })
        )
      );

      await coordinatedRefresh();

      expect(authChannel.postAuthChannelMessage).toHaveBeenCalledWith({
        type: "refreshed",
        accessToken: "refreshed-token",
      });
    });

    it("broadcasts 'refresh-failed' on failure, without throwing (R13)", async () => {
      server.use(http.post("/auth/refresh", () => HttpResponse.json({}, { status: 401 })));

      const ok = await coordinatedRefresh();

      expect(ok).toBe(false);
      expect(authChannel.postAuthChannelMessage).toHaveBeenCalledWith({
        type: "refresh-failed",
      });
    });
  });
});
