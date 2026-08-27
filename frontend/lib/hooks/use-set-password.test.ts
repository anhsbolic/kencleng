import { QueryClient } from "@tanstack/react-query";
import { beforeEach, describe, expect, it } from "vitest";
import { useAuthStore } from "@/lib/stores/auth-store";
import { applySetPasswordSuccess } from "./use-set-password";

// Zustand store is a module-level singleton — reset between tests so
// one test's session state doesn't leak into the next (same convention
// as use-login.test.ts).
beforeEach(() => {
  useAuthStore.setState({ accessToken: null });
});

describe("applySetPasswordSuccess", () => {
  it("invalidates account.me on the 'added' branch (R6) — no session change", async () => {
    useAuthStore.setState({ accessToken: "still-valid-token" });
    const queryClient = new QueryClient();
    queryClient.setQueryData(["account", "me"], { id: "1" });

    applySetPasswordSuccess({ ok: true, branch: "added", message: "ok" }, queryClient);

    expect(useAuthStore.getState().accessToken).toBe("still-valid-token");
    // Invalidated (marked stale), not wiped outright — same intent as
    // an ordinary invalidateQueries call.
    expect(queryClient.getQueryState(["account", "me"])?.isInvalidated).toBe(true);
  });

  it("clears the access token and the whole query cache on the 'changed' branch (R10)", () => {
    useAuthStore.setState({ accessToken: "still-valid-token" });
    const queryClient = new QueryClient();
    queryClient.setQueryData(["account", "me"], { id: "1" });

    applySetPasswordSuccess({ ok: true, branch: "changed", message: "ok" }, queryClient);

    expect(useAuthStore.getState().accessToken).toBeNull();
    expect(queryClient.getQueryData(["account", "me"])).toBeUndefined();
  });

  it("is a no-op on the validation branch — no token clear, no invalidation", () => {
    useAuthStore.setState({ accessToken: "still-valid-token" });
    const queryClient = new QueryClient();
    queryClient.setQueryData(["account", "me"], { id: "1" });

    applySetPasswordSuccess({ ok: false, kind: "validation", errors: [] }, queryClient);

    expect(useAuthStore.getState().accessToken).toBe("still-valid-token");
    expect(queryClient.getQueryState(["account", "me"])?.isInvalidated).toBe(false);
  });
});
