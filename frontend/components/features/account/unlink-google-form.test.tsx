import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { HttpResponse, http } from "msw";
import type { ReactNode } from "react";
import { describe, expect, it } from "vitest";
import { server } from "@/mocks/server";
import { accountKeys } from "@/lib/hooks/use-account-me";
import { UnlinkGoogleForm } from "./unlink-google-form";

function withQueryClient(children: ReactNode, queryClient: QueryClient) {
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}

function newQueryClient() {
  return new QueryClient({ defaultOptions: { mutations: { retry: false } } });
}

function submit(password = "correct-pw") {
  fireEvent.change(screen.getByLabelText("Konfirmasi password"), { target: { value: password } });
  fireEvent.click(screen.getByRole("button", { name: "Lepas Tautan Google" }));
}

describe("UnlinkGoogleForm", () => {
  it("renders a password field and a destructive submit button (R16)", () => {
    render(withQueryClient(<UnlinkGoogleForm />, newQueryClient()));

    expect(screen.getByLabelText("Konfirmasi password")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Lepas Tautan Google" })).toBeInTheDocument();
  });

  it("blocks submission of an empty password client-side", async () => {
    render(withQueryClient(<UnlinkGoogleForm />, newQueryClient()));

    fireEvent.click(screen.getByRole("button", { name: "Lepas Tautan Google" }));

    expect(await screen.findByText("Password wajib diisi")).toBeInTheDocument();
  });

  it("invalidates account.me on 200, no local success view (R17)", async () => {
    const queryClient = newQueryClient();
    queryClient.setQueryData(accountKeys.me(), { id: "1" });
    render(withQueryClient(<UnlinkGoogleForm />, queryClient));

    submit();

    await waitFor(() =>
      expect(queryClient.getQueryState(accountKeys.me())?.isInvalidated).toBe(true)
    );
    expect(screen.queryByRole("alert")).not.toBeInTheDocument();
  });

  it("shows the backend's detail verbatim on 401, form stays interactive (R18)", async () => {
    server.use(
      http.post("/account/security/google/unlink", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/invalid-credentials",
            title: "Invalid Credentials",
            status: 401,
            detail: "Email atau password salah.",
          },
          { status: 401 }
        )
      )
    );

    render(withQueryClient(<UnlinkGoogleForm />, newQueryClient()));
    submit("wrong-pw");

    expect(await screen.findByRole("alert")).toHaveTextContent("Email atau password salah.");
    expect(screen.getByLabelText("Konfirmasi password")).toBeInTheDocument();
  });

  it("shows the backend's detail verbatim for the only-identity 409 case, distinct from the unverified case (R18)", async () => {
    server.use(
      http.post("/account/security/google/unlink", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/only-identity",
            title: "Cannot Unlink Only Identity",
            status: 409,
            detail:
              "Google adalah satu-satunya metode login Anda. Atur email dan password dulu sebelum melepas tautan.",
          },
          { status: 409 }
        )
      )
    );

    render(withQueryClient(<UnlinkGoogleForm />, newQueryClient()));
    submit();

    expect(await screen.findByRole("alert")).toHaveTextContent(
      "Google adalah satu-satunya metode login Anda. Atur email dan password dulu sebelum melepas tautan."
    );
  });

  it("shows the backend's detail verbatim for the unverified-remaining-identity 409 case (R18)", async () => {
    server.use(
      http.post("/account/security/google/unlink", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/unverified-remaining-identity",
            title: "Remaining Identity Not Verified",
            status: 409,
            detail:
              "Kamu sudah atur email dan password, tapi belum diverifikasi. Verifikasi email kamu dulu sebelum bisa melepas tautan Google.",
          },
          { status: 409 }
        )
      )
    );

    render(withQueryClient(<UnlinkGoogleForm />, newQueryClient()));
    submit();

    expect(await screen.findByRole("alert")).toHaveTextContent(
      "Kamu sudah atur email dan password, tapi belum diverifikasi. Verifikasi email kamu dulu sebelum bisa melepas tautan Google."
    );
  });

  it("shows a generic banner on a network failure, form stays interactive", async () => {
    server.use(http.post("/account/security/google/unlink", () => HttpResponse.error()));

    render(withQueryClient(<UnlinkGoogleForm />, newQueryClient()));
    submit();

    expect(await screen.findByText("Terjadi kesalahan. Silakan coba lagi.")).toBeInTheDocument();
    expect(screen.getByLabelText("Konfirmasi password")).toBeInTheDocument();
  });
});
