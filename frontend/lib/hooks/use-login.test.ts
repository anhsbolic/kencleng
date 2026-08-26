import { QueryClient } from "@tanstack/react-query";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { useAuthStore } from "@/lib/stores/auth-store";
import type { LoginResponse } from "@/lib/api/account";
import { accountKeys } from "./use-account-me";
import { applyLoginSuccess } from "./use-login";

// Zustand store is a module-level singleton — reset between tests so
// one test's login doesn't leak into the next (same convention as
// auth-bootstrap-provider.test.tsx).
beforeEach(() => {
  useAuthStore.setState({ accessToken: null });
});

const okResult: LoginResponse = {
  status: "ok",
  access_token: "fresh-access-token",
  access_token_expires_at: "2026-08-26T12:15:00.000Z",
  user: {
    id: "00000000-0000-0000-0000-000000000001",
    name: "Siti",
    email: "siti@example.com",
    email_verified: true,
    roles: [],
    auth_providers: ["email_password"],
    mfa_enabled: false,
    created_at: "2026-01-01T00:00:00.000Z",
  },
};

describe("applyLoginSuccess", () => {
  it("sets the access token, writes the user directly into the account.me cache, and redirects on the 'ok' branch (R3)", () => {
    const queryClient = new QueryClient();
    const router = { push: vi.fn() } as unknown as ReturnType<
      typeof import("next/navigation").useRouter
    >;

    applyLoginSuccess(okResult, queryClient, router);

    expect(useAuthStore.getState().accessToken).toBe("fresh-access-token");
    expect(queryClient.getQueryData(accountKeys.me())).toEqual(okResult.user);
    expect(router.push).toHaveBeenCalledWith("/dashboard/profile");
  });

  it("is a no-op on the 'mfa_required' branch — no token set, no cache write, no redirect (R4)", () => {
    const queryClient = new QueryClient();
    const router = { push: vi.fn() } as unknown as ReturnType<
      typeof import("next/navigation").useRouter
    >;

    applyLoginSuccess(
      { status: "mfa_required", mfa_pending_token: "pending-token" },
      queryClient,
      router
    );

    expect(useAuthStore.getState().accessToken).toBeNull();
    expect(queryClient.getQueryData(accountKeys.me())).toBeUndefined();
    expect(router.push).not.toHaveBeenCalled();
  });
});
