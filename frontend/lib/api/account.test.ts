import { HttpResponse, http } from "msw";
import { describe, expect, it } from "vitest";
import { server } from "@/mocks/server";
import { ApiError } from "./client";
import {
  login,
  loginMfa,
  logout,
  mfaDisable,
  mfaEnroll,
  mfaEnrollConfirm,
  register,
  resendVerification,
  verifyEmail,
} from "./account";

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

describe("login", () => {
  it("resolves the 'ok' branch with the access token and user on 200 (R3)", async () => {
    const result = await login({ email: "donatur@example.com", password: "rahasia123" });

    expect(result.status).toBe("ok");
    if (result.status === "ok") {
      expect(result.access_token).toBe("mock-access-token");
      expect(result.user.email).toBe("donatur@example.com");
    }
  });

  it("resolves the 'mfa_required' branch with a pending token, no access token (R4)", async () => {
    server.use(
      http.post("/auth/login", () =>
        HttpResponse.json({ status: "mfa_required", mfa_pending_token: "pending-token" })
      )
    );

    const result = await login({ email: "donatur@example.com", password: "rahasia123" });

    expect(result).toEqual({ status: "mfa_required", mfa_pending_token: "pending-token" });
  });

  it("throws ApiError with the generic detail on 401 (R5)", async () => {
    server.use(
      http.post("/auth/login", () =>
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

    await expect(
      login({ email: "donatur@example.com", password: "wrong" })
    ).rejects.toMatchObject({ status: 401, detail: "Email atau password salah." });
  });

  it("throws ApiError with the identical detail text on 429 lockout (R5)", async () => {
    server.use(
      http.post("/auth/login", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/too-many-requests",
            title: "Too Many Requests",
            status: 429,
            detail: "Email atau password salah.",
          },
          { status: 429 }
        )
      )
    );

    await expect(
      login({ email: "donatur@example.com", password: "wrong" })
    ).rejects.toMatchObject({ status: 429, detail: "Email atau password salah." });
  });
});

describe("loginMfa", () => {
  it("resolves the 'ok' branch on 200 (R7)", async () => {
    const result = await loginMfa({ mfa_pending_token: "pending-token", totp_code: "123456" });
    expect(result.status).toBe("ok");
  });

  it("throws ApiError on 401 (wrong code / expired token) (R8)", async () => {
    server.use(
      http.post("/auth/login/mfa", () =>
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

    await expect(
      loginMfa({ mfa_pending_token: "pending-token", totp_code: "000000" })
    ).rejects.toMatchObject({ status: 401 });
  });

  it("throws ApiError on 429 MFA-stage lockout (R8)", async () => {
    server.use(
      http.post("/auth/login/mfa", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/too-many-requests",
            title: "Too Many Requests",
            status: 429,
            detail: "Email atau password salah.",
          },
          { status: 429 }
        )
      )
    );

    await expect(
      loginMfa({ mfa_pending_token: "pending-token", totp_code: "000000" })
    ).rejects.toMatchObject({ status: 429 });
  });
});

describe("logout", () => {
  it("resolves on 204 (R17/R19)", async () => {
    await expect(logout()).resolves.toBeUndefined();
  });

  it("throws ApiError on an unexpected 5xx — caller decides how to treat it", async () => {
    server.use(http.post("/auth/logout", () => HttpResponse.json({}, { status: 500 })));

    await expect(logout()).rejects.toBeInstanceOf(ApiError);
  });
});

describe("mfaEnroll", () => {
  it("resolves with the otpauth_uri on 200 (R5)", async () => {
    const result = await mfaEnroll();

    expect(result.otpauth_uri).toContain("otpauth://totp/");
  });

  it("throws ApiError(409) when already enabled — undocumented in schema.d.ts, handled defensively anyway (R6)", async () => {
    server.use(
      http.post("/account/security/mfa/enroll", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/mfa-already-enabled",
            title: "MFA Already Enabled",
            status: 409,
            detail: "MFA sudah aktif.",
          },
          { status: 409 }
        )
      )
    );

    await expect(mfaEnroll()).rejects.toMatchObject({ status: 409, detail: "MFA sudah aktif." });
  });

  it("throws ApiError on network failure/unexpected 5xx (R7)", async () => {
    server.use(http.post("/account/security/mfa/enroll", () => HttpResponse.error()));

    await expect(mfaEnroll()).rejects.toBeInstanceOf(ApiError);
  });
});

describe("mfaEnrollConfirm", () => {
  it("resolves with exactly 10 backup_codes on 200 (R9 hook-layer)", async () => {
    const result = await mfaEnrollConfirm({ totp_code: "123456" });

    expect(result.backup_codes).toHaveLength(10);
  });

  it("throws ApiError on 422 — not a discriminated validation result (R10, D8)", async () => {
    server.use(
      http.post("/account/security/mfa/enroll/confirm", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/validation-failed",
            title: "Validation Failed",
            status: 422,
            errors: [{ field: "totp_code", message: "Invalid code" }],
          },
          { status: 422 }
        )
      )
    );

    await expect(mfaEnrollConfirm({ totp_code: "000000" })).rejects.toMatchObject({ status: 422 });
  });

  it("throws ApiError on network failure/unexpected 5xx (R11)", async () => {
    server.use(http.post("/account/security/mfa/enroll/confirm", () => HttpResponse.error()));

    await expect(mfaEnrollConfirm({ totp_code: "123456" })).rejects.toBeInstanceOf(ApiError);
  });
});

describe("mfaDisable", () => {
  it("resolves on 200 for the email_password branch — password in body (R15 hook-layer)", async () => {
    await expect(mfaDisable({ password: "correct-pw" })).resolves.toMatchObject({
      message: "MFA berhasil dinonaktifkan.",
    });
  });

  it("resolves on 200 for the Google-only branch — no body (R18 hook-layer)", async () => {
    await expect(mfaDisable()).resolves.toMatchObject({
      message: "MFA berhasil dinonaktifkan.",
    });
  });

  it("throws ApiError(401) on wrong password (R16)", async () => {
    server.use(
      http.post("/account/security/mfa/disable", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/unauthorized",
            title: "Unauthorized",
            status: 401,
            detail: "Password salah.",
          },
          { status: 401 }
        )
      )
    );

    await expect(mfaDisable({ password: "wrong-pw" })).rejects.toMatchObject({
      status: 401,
      detail: "Password salah.",
    });
  });

  it("throws ApiError(401) on missing/expired Google reauth marker (R19)", async () => {
    server.use(
      http.post("/account/security/mfa/disable", () =>
        HttpResponse.json(
          {
            type: "https://kencleng.dev/errors/unauthorized",
            title: "Unauthorized",
            status: 401,
            detail: "Access token tidak valid atau sudah kedaluwarsa.",
          },
          { status: 401 }
        )
      )
    );

    await expect(mfaDisable()).rejects.toMatchObject({ status: 401 });
  });
});
