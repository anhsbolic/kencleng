import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, screen } from "@testing-library/react";
import type { ReactNode } from "react";
import { describe, expect, it } from "vitest";
import { GoogleIdentityControl } from "./google-identity-control";

function withQueryClient(children: ReactNode) {
  const queryClient = new QueryClient({
    defaultOptions: { mutations: { retry: false } },
  });
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

describe("GoogleIdentityControl", () => {
  it("renders the link-to-Google trigger when no Google identity is present (R13)", () => {
    render(
      withQueryClient(
        <GoogleIdentityControl hasGoogle={false} canUnlink={false} blockedReason={null} />
      )
    );

    expect(screen.getByRole("link", { name: "Hubungkan ke Google" })).toHaveAttribute(
      "href",
      "/auth/google/redirect?intent=link"
    );
    expect(screen.queryByRole("button", { name: "Lepas Tautan Google" })).not.toBeInTheDocument();
  });

  it("shows the only-identity blocked message and no unlink form (R14)", () => {
    render(
      withQueryClient(
        <GoogleIdentityControl hasGoogle={true} canUnlink={false} blockedReason="only-identity" />
      )
    );

    expect(
      screen.getByText(
        "Google adalah satu-satunya metode login Anda. Atur email dan password dulu sebelum melepas tautan."
      )
    ).toBeInTheDocument();
    expect(screen.queryByRole("button", { name: "Lepas Tautan Google" })).not.toBeInTheDocument();
  });

  it("shows the unverified-remaining-identity blocked message, distinct from the only-identity one (R15)", () => {
    render(
      withQueryClient(
        <GoogleIdentityControl hasGoogle={true} canUnlink={false} blockedReason="unverified" />
      )
    );

    expect(
      screen.getByText(
        "Kamu sudah atur email dan password, tapi belum diverifikasi. Verifikasi email kamu dulu sebelum bisa melepas tautan Google."
      )
    ).toBeInTheDocument();
    expect(
      screen.queryByText(
        "Google adalah satu-satunya metode login Anda. Atur email dan password dulu sebelum melepas tautan."
      )
    ).not.toBeInTheDocument();
  });

  it("renders UnlinkGoogleForm when unlinkable (R16)", () => {
    render(
      withQueryClient(
        <GoogleIdentityControl hasGoogle={true} canUnlink={true} blockedReason={null} />
      )
    );

    expect(screen.getByLabelText("Konfirmasi password")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Lepas Tautan Google" })).toBeInTheDocument();
  });
});
