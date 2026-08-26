import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, screen } from "@testing-library/react";
import type { ReactNode } from "react";
import { describe, expect, it, vi } from "vitest";
import LoginPage from "./page";

vi.mock("next/navigation", () => ({
  useSearchParams: () => ({ get: () => null }),
  useRouter: () => ({ push: vi.fn() }),
}));

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { mutations: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

// Composition-level only — `LoginForm`'s own behavior (R1-R10) is
// covered by `components/features/account/login-form.test.tsx`,
// matching how `/register` has no page-level test of its own beyond
// `RegisterForm`'s (techplan account/03-login-session-management,
// task-03: this file is rewritten, not deleted, since it still verifies
// the page composes `LoginForm` + `GoogleCallbackError` correctly).
describe("LoginPage", () => {
  it("renders the heading and the real credential form (no longer the Phase 0/task #2 placeholder)", () => {
    render(withQueryClient(<LoginPage />));

    expect(screen.getByRole("heading", { name: "Masuk" })).toBeInTheDocument();
    expect(screen.getByLabelText("Email")).toBeInTheDocument();
    expect(screen.getByLabelText("Password")).toBeInTheDocument();
    expect(screen.queryByText(/segera hadir/i)).not.toBeInTheDocument();
  });
});
