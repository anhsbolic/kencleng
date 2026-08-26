import { render } from "@testing-library/react";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { useAuthStore } from "@/lib/stores/auth-store";
import { SessionGuardProvider } from "./session-guard-provider";

const pushMock = vi.fn();
vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: pushMock }),
}));

beforeEach(() => {
  useAuthStore.setState({ accessToken: null });
  pushMock.mockClear();
});

describe("SessionGuardProvider", () => {
  it("redirects to /login exactly once when accessToken transitions from a real value to null (R15)", () => {
    useAuthStore.setState({ accessToken: "some-token" });
    render(
      <SessionGuardProvider>
        <div>app content</div>
      </SessionGuardProvider>
    );

    useAuthStore.getState().clearAccessToken();

    expect(pushMock).toHaveBeenCalledTimes(1);
    expect(pushMock).toHaveBeenCalledWith("/login");
  });

  it("does not redirect when accessToken starts null and stays null — genuine guest (R16)", () => {
    render(
      <SessionGuardProvider>
        <div>app content</div>
      </SessionGuardProvider>
    );

    // Simulates AuthBootstrapProvider's own failed boot-time refresh (its R10) — still null.
    useAuthStore.setState({ accessToken: null });

    expect(pushMock).not.toHaveBeenCalled();
  });

  it("does not redirect on a null -> token transition (successful login/hydration)", () => {
    render(
      <SessionGuardProvider>
        <div>app content</div>
      </SessionGuardProvider>
    );

    useAuthStore.getState().setAccessToken("fresh-token");

    expect(pushMock).not.toHaveBeenCalled();
  });

  it("renders its children", () => {
    const { getByText } = render(
      <SessionGuardProvider>
        <div>app content</div>
      </SessionGuardProvider>
    );

    expect(getByText("app content")).toBeInTheDocument();
  });
});
