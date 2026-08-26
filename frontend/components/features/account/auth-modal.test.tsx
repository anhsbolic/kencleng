import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen } from "@testing-library/react";
import type { ReactNode } from "react";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { useAuthModalStore } from "@/lib/stores/auth-modal-store";
import { AuthModal } from "./auth-modal";

vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: vi.fn() }),
}));

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { mutations: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

beforeEach(() => {
  useAuthModalStore.setState({ mode: null });
});

describe("AuthModal", () => {
  it("renders nothing when closed", () => {
    render(withQueryClient(<AuthModal />));

    expect(screen.queryByRole("dialog")).not.toBeInTheDocument();
  });

  it("renders LoginForm when mode is 'login'", () => {
    useAuthModalStore.setState({ mode: "login" });
    render(withQueryClient(<AuthModal />));

    expect(screen.getByRole("dialog", { name: "Masuk" })).toBeInTheDocument();
    expect(screen.getByLabelText("Email")).toBeInTheDocument();
    expect(screen.getByLabelText("Password")).toBeInTheDocument();
  });

  it("renders RegisterForm when mode is 'register'", () => {
    useAuthModalStore.setState({ mode: "register" });
    render(withQueryClient(<AuthModal />));

    expect(screen.getByRole("dialog", { name: "Daftar" })).toBeInTheDocument();
    expect(screen.getByLabelText("Nama")).toBeInTheDocument();
  });

  it("the close button closes the modal", () => {
    useAuthModalStore.setState({ mode: "login" });
    render(withQueryClient(<AuthModal />));

    fireEvent.click(screen.getByRole("button", { name: "Tutup" }));

    expect(useAuthModalStore.getState().mode).toBeNull();
  });

  it("Escape closes the modal", () => {
    useAuthModalStore.setState({ mode: "login" });
    render(withQueryClient(<AuthModal />));

    fireEvent.keyDown(document, { key: "Escape" });

    expect(useAuthModalStore.getState().mode).toBeNull();
  });

  it("clicking the backdrop closes the modal", () => {
    useAuthModalStore.setState({ mode: "login" });
    const { container } = render(withQueryClient(<AuthModal />));

    const backdrop = container.querySelector('[aria-hidden="true"]');
    expect(backdrop).not.toBeNull();
    fireEvent.click(backdrop as Element);

    expect(useAuthModalStore.getState().mode).toBeNull();
  });

  it("switches from login to register via LoginForm's own 'Daftar' link, without navigating", () => {
    useAuthModalStore.setState({ mode: "login" });
    render(withQueryClient(<AuthModal />));

    fireEvent.click(screen.getByRole("button", { name: "Daftar" }));

    expect(useAuthModalStore.getState().mode).toBe("register");
    expect(screen.getByLabelText("Nama")).toBeInTheDocument();
  });

  it("switches from register to login via RegisterForm's own 'Masuk' link, without navigating", () => {
    useAuthModalStore.setState({ mode: "register" });
    render(withQueryClient(<AuthModal />));

    fireEvent.click(screen.getByRole("button", { name: "Masuk" }));

    expect(useAuthModalStore.getState().mode).toBe("login");
    expect(screen.getByLabelText("Email")).toBeInTheDocument();
  });
});
