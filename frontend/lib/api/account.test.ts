import { HttpResponse, http } from "msw";
import { describe, expect, it } from "vitest";
import { server } from "@/mocks/server";
import { ApiError } from "./client";
import { register, resendVerification, verifyEmail } from "./account";

describe("register", () => {
  it("resolves ok:true with the backend's own message on 202 (R4)", async () => {
    const result = await register({
      name: "Siti",
      email: "siti@example.com",
      password: "rahasia123",
    });

    expect(result).toEqual({
      ok: true,
      message:
        "Kalau email belum terdaftar, cek inbox untuk verifikasi. Kalau sudah, cek inbox untuk instruksi lebih lanjut.",
    });
  });

  it("resolves ok:false with per-field messages on 422, never throws (R5)", async () => {
    server.use(
      http.post("/auth/register", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/validation-failed",
            title: "Validation Failed",
            status: 422,
            errors: [
              {
                field: "password",
                message: "Password ditemukan di daftar breach publik, silakan pilih password lain.",
              },
            ],
          },
          { status: 422 }
        )
      )
    );

    const result = await register({
      name: "Siti",
      email: "siti@example.com",
      password: "breached1",
    });

    expect(result).toEqual({
      ok: false,
      kind: "validation",
      errors: [
        {
          field: "password",
          message: "Password ditemukan di daftar breach publik, silakan pilih password lain.",
        },
      ],
    });
  });

  it("throws ApiError carrying status + detail on 429 (R10)", async () => {
    server.use(
      http.post("/auth/register", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/too-many-requests",
            title: "Too Many Requests",
            status: 429,
            detail: "Terlalu banyak percobaan gagal. Coba lagi dalam 15 menit.",
          },
          { status: 429 }
        )
      )
    );

    await expect(
      register({ name: "Siti", email: "siti@example.com", password: "rahasia123" })
    ).rejects.toMatchObject({
      status: 429,
      detail: "Terlalu banyak percobaan gagal. Coba lagi dalam 15 menit.",
    });
  });

  it("throws a plain ApiError on a network failure (R6)", async () => {
    server.use(http.post("/auth/register", () => HttpResponse.error()));

    await expect(
      register({ name: "Siti", email: "siti@example.com", password: "rahasia123" })
    ).rejects.toBeInstanceOf(ApiError);
  });
});

describe("verifyEmail", () => {
  it("resolves with the backend's message on 200 (R13)", async () => {
    const result = await verifyEmail({ token: "valid-token" });
    expect(result.message).toBe("Email berhasil diverifikasi.");
  });

  it("throws ApiError with status 410 and the expiry detail (R14)", async () => {
    server.use(
      http.post("/auth/verify-email", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/token-expired",
            title: "Token Expired",
            status: 410,
            detail: "Link verifikasi sudah kedaluwarsa. Silakan minta kirim ulang.",
          },
          { status: 410 }
        )
      )
    );

    await expect(verifyEmail({ token: "expired-token" })).rejects.toMatchObject({
      status: 410,
      detail: "Link verifikasi sudah kedaluwarsa. Silakan minta kirim ulang.",
    });
  });

  it("throws ApiError with status 404 on a not-found/used/revoked token (R15)", async () => {
    server.use(
      http.post("/auth/verify-email", () =>
        HttpResponse.json(
          { type: "about:blank", title: "Not Found", status: 404 },
          { status: 404 }
        )
      )
    );

    await expect(verifyEmail({ token: "bad-token" })).rejects.toMatchObject({
      status: 404,
    });
  });
});

describe("resendVerification", () => {
  it("always resolves with the same generic message regardless of match (R9)", async () => {
    const result = await resendVerification({ email: "anyone@example.com" });
    expect(result.message).toBe("Kalau email terdaftar, instruksi sudah dikirim.");
  });

  it("throws ApiError with status 429 (R10)", async () => {
    server.use(
      http.post("/auth/verify-email/resend", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/too-many-requests",
            title: "Too Many Requests",
            status: 429,
            detail: "Terlalu banyak percobaan gagal. Coba lagi dalam 15 menit.",
          },
          { status: 429 }
        )
      )
    );

    await expect(resendVerification({ email: "anyone@example.com" })).rejects.toMatchObject({
      status: 429,
    });
  });
});
