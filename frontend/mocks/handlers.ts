// Shared MSW request handlers — used by both the Vitest `msw/node`
// setup (mocks/server.ts) and the browser `msw/browser` dev-mode
// worker (mocks/browser.ts). One handler per endpoint, added as the
// page/component that needs it gets built (mock-first workflow,
// scaffold-frontend.md's durable section) — not speculatively ahead
// of demonstrated need.
import { HttpResponse, http } from "msw";
import type { User } from "@/lib/api/account";
import type { CampaignListResponse } from "@/lib/api/campaign";
import type { UnreadCountResponse } from "@/lib/api/notification";

// Days-from-now helper, used only to keep the fixture's `deadline`
// plausible relative to `progress.days_remaining` — real values come
// from the backend once `campaign` domain's backend track ships.
function daysFromNow(days: number): string {
  return new Date(Date.now() + days * 24 * 60 * 60 * 1000).toISOString();
}

// Fixed fixture set for `GET /campaigns` — per `page-map.md` footnote 1,
// this is deliberately NOT filtered/sorted by any "featured" concept the
// API doesn't have. No `organization` field — `CampaignListItem` can't
// supply one (see techplan Decision 1); don't add it here even though
// it would make the mock look more complete, that's exactly the drift
// `component-test-mocking-discipline.md` warns against.
const mockCampaigns: CampaignListResponse = {
  data: [
    {
      id: "10000000-0000-0000-0000-000000000001",
      organization_id: "20000000-0000-0000-0000-000000000001",
      title: "Air bersih untuk 240 keluarga di Dusun Sukamaju",
      description: "Penyediaan sumur bor dan instalasi pipa air bersih.",
      category: "bencana_alam",
      location: "Dusun Sukamaju, Jawa Barat",
      beneficiary_description: "240 keluarga di Dusun Sukamaju",
      target_amount: "15000000.00",
      max_amount: null,
      collected_amount: "10200000.00",
      deadline: daysFromNow(12),
      status: "published",
      publish_at: null,
      published_at: daysFromNow(-18),
      closed_at: null,
      decision_note: null,
      created_by: "30000000-0000-0000-0000-000000000001",
      created_at: daysFromNow(-20),
      updated_at: daysFromNow(-1),
      progress: { percentage: 68, donor_count: 84, days_remaining: 12 },
    },
    {
      id: "10000000-0000-0000-0000-000000000002",
      organization_id: "20000000-0000-0000-0000-000000000002",
      title: "Renovasi musala dan ruang belajar anak",
      description: "Perbaikan atap, lantai, dan penambahan ruang belajar.",
      category: "sosial",
      location: "Musala Al-Ikhlas",
      beneficiary_description: null,
      target_amount: "8000000.00",
      max_amount: null,
      collected_amount: "3280000.00",
      deadline: daysFromNow(26),
      status: "published",
      publish_at: null,
      published_at: daysFromNow(-9),
      closed_at: null,
      decision_note: null,
      created_by: "30000000-0000-0000-0000-000000000002",
      created_at: daysFromNow(-10),
      updated_at: daysFromNow(-1),
      progress: { percentage: 41, donor_count: 37, days_remaining: 26 },
    },
    {
      id: "10000000-0000-0000-0000-000000000003",
      organization_id: "20000000-0000-0000-0000-000000000003",
      title: "Beasiswa sekolah untuk anak nelayan Lombok",
      description: "Bantuan biaya sekolah dan perlengkapan belajar.",
      category: "pendidikan",
      location: "Lombok, Nusa Tenggara Barat",
      beneficiary_description: "Anak-anak nelayan usia sekolah di pesisir Lombok",
      target_amount: "20000000.00",
      max_amount: null,
      collected_amount: "17400000.00",
      deadline: daysFromNow(5),
      status: "published",
      publish_at: null,
      published_at: daysFromNow(-25),
      closed_at: null,
      decision_note: null,
      created_by: "30000000-0000-0000-0000-000000000003",
      created_at: daysFromNow(-27),
      updated_at: daysFromNow(-1),
      progress: { percentage: 87, donor_count: 156, days_remaining: 5 },
    },
  ],
  pagination: { next_cursor: null, has_more: false },
};

// Default fixture: a plain logged-in donatur (no elevated roles) —
// the "just the default logged-in-donatur case" the Dashboard Shell
// nav checkpoint (phase0-shared-infra.md Step 7) says not to stop at.
// Individual tests override this via `server.use(...)` with a
// different `roles` array to exercise the other combinations.
const mockUser: User = {
  id: "00000000-0000-0000-0000-000000000001",
  name: "Donatur Contoh",
  email: "donatur@example.com",
  email_verified: true,
  roles: [],
  auth_providers: ["email_password"],
  mfa_enabled: false,
  created_at: new Date().toISOString(),
};

const mockUnreadCount: UnreadCountResponse = { unread_count: 3 };

// Default happy-path fixtures for account/01-register-email-verification
// (register/verify-email/resend). Individual tests override the 422/
// 404/410/429/network-error cases via `server.use(...)`, same pattern
// as the roles-array override already used for `mockUser` elsewhere.
const mockRegisterAccepted = {
  message:
    "Kalau email belum terdaftar, cek inbox untuk verifikasi. Kalau sudah, cek inbox untuk instruksi lebih lanjut.",
};
const mockVerifyEmailVerified = { message: "Email berhasil diverifikasi." };
const mockResendAccepted = {
  message: "Kalau email terdaftar, instruksi sudah dikirim.",
};

export const handlers = [
  http.get("/account/me", () => HttpResponse.json(mockUser)),
  http.get("/notifications/unread-count", () =>
    HttpResponse.json(mockUnreadCount)
  ),
  http.get("/campaigns", () => HttpResponse.json(mockCampaigns)),
  http.post("/auth/register", () =>
    HttpResponse.json(mockRegisterAccepted, { status: 202 })
  ),
  http.post("/auth/verify-email", () =>
    HttpResponse.json(mockVerifyEmailVerified, { status: 200 })
  ),
  http.post("/auth/verify-email/resend", () =>
    HttpResponse.json(mockResendAccepted, { status: 202 })
  ),
  // Default: no refresh cookie present (unauthenticated guest) — the
  // overwhelmingly common case, since AuthBootstrapProvider calls this
  // on every app load, not just after a Google OAuth callback.
  // Individual tests override via `server.use(...)` for the
  // successful-hydration case (techplan account/02-google-oauth-
  // login-register, R8-R11).
  http.post("/auth/refresh", () => HttpResponse.json({}, { status: 401 })),
];
