# Kencleng — Phase 0 (Foundational / Cross-Cutting) — Detailed Feature Spec
 
> Status: Draft — detailed at business-rule/flow level. API contract format resolved (OpenAPI 3.x spec-first, see `kencleng-backend-tech-stack.md`); endpoint-level detail still excluded from this doc pending actual spec authoring.
> Last updated: 2026-07-24 (rev 4 — audit-log scope extended to representative management actions, notification type enum extended for curation-lifecycle events)
 
## Context
 
This document details the foundational / cross-cutting concerns
needed for the platform to function, sitting **outside** the 3 core
business phases (`kencleng-phase1-detail.md`,
`kencleng-phase2-detail.md`, `kencleng-phase3-detail.md`). These
features are referenced implicitly throughout Phase 1–3 (notification,
file upload, auth) and are formalized here as shared services/entities
rather than being redesigned ad-hoc per feature.
 
**Auth architecture note**: to support multiple identity providers in
the future (Google OAuth, phone number + OTP) without a rewrite, `User`
(profile) is deliberately separated from `AuthIdentity` (login
method). **v1 implements `email_password` AND `google` providers**
**[REVISED]**; phone+OTP is modeled but not implemented until a later
version.
 
Same document format as Phase 1/2/3 (Overview / Actors / Preconditions
/ Flow / Business rules / Alternate flows / Data touched / State
transition / Concurrency notes / Security notes / Open questions).
Endpoint-level detail (concrete paths/methods) is still excluded from
this doc — that lives in `api/openapi.yaml` per the resolved API
contract format decision (`kencleng-backend-tech-stack.md`, "API
Contract & Codegen"), not duplicated here.
 
---
 
## 1. Registrasi & Verifikasi Email
 
**Overview**
Proses dasar user membuat akun baru lewat provider `email_password`
dan memverifikasi kepemilikan email sebelum mendapat akses penuh ke
platform.
 
**Actors**
User (calon terdaftar)
 
**Preconditions**
Tidak ada (registrasi publik, terbuka untuk siapa saja)
 
**Flow (happy path)**
1. User isi form: email, password, nama
2. Sistem validasi: email belum terdaftar sebagai `AuthIdentity` tipe
   `email_password`, password memenuhi kriteria minimum, dan **tidak
   ada di breach-list HaveIBeenPwned** (lihat Business Rule di bawah)
   **[REVISED]**
3. Sistem membuat `User` (profil) + `AuthIdentity`
   (`provider_type = email_password`, `identifier = email`,
   `credential_secret = password_hash`, `verified_at = null`)
4. Sistem kirim email verifikasi (link berisi token, expired dalam
   **24 jam** **[RESOLVED]**)
5. User klik link → token divalidasi → `AuthIdentity.verified_at = now`
**Business rules / validation**
- `identifier` (email) harus unik **per `provider_type`** — bukan
  unik secara global di tabel `AuthIdentity`, karena provider lain
  (Google, phone) punya namespace identifier berbeda
- Token verifikasi single-use dan expired dalam **24 jam
  [RESOLVED]**; kalau expired, user bisa request kirim ulang
- **Password policy [RESOLVED — NEW]**: minimum **8 karakter**,
  **length-only, tanpa syarat complexity** (tidak wajib kombinasi
  uppercase/angka/simbol) — mengikuti pendekatan NIST 800-63B, karena
  panjang password berkontribusi lebih besar terhadap entropy
  dibanding character-class rules yang justru mendorong pola
  predictable
- **Breach-list check [RESOLVED — NEW]**: password dicek terhadap
  HaveIBeenPwned API memakai k-anonymity model — hanya 5 karakter
  pertama dari SHA-1 hash password yang dikirim ke API, password
  plaintext dan hash penuh tidak pernah meninggalkan server. Dicek
  saat registrasi **dan** saat reset password (Fitur 2B). Kalau
  password ditemukan di breach-list, registrasi ditolak dengan pesan
  jelas, user diminta pilih password lain
**Alternate flows / edge cases**
- User daftar dengan email yang sudah ada tapi belum verified →
  ditolak, ditawarkan opsi kirim ulang email verifikasi (bukan buat
  akun baru)
- HaveIBeenPwned API tidak bisa diakses (network error) →
  ~~(OPEN — belum diputuskan: fail-open dan lanjutkan registrasi
  tanpa cek, atau fail-closed dan tolak registrasi sementara)~~ →
  **resolved: fail-open** — registrasi (dan reset password, Fitur 2B)
  tetap lanjut tanpa breach-check, kegagalan panggilan API dicatat via
  log biasa (bukan audit log khusus, karena ini bukan aksi sensitif
  yang butuh audit trail — cukup buat observability). Breach-check
  adalah lapisan defense-in-depth tambahan, bukan primary defense
  (primary defense tetap password length policy + hashing + rate
  limiting) — availability registrasi tidak boleh bergantung ke
  uptime API pihak ketiga di luar kontrol kita. **[RESOLVED — NEW]**
**Data touched**
- **Create** `User` (id, name, primary_email, created_at)
- **Create** `AuthIdentity` (id, user_id, provider_type =
  `email_password`, identifier = email, credential_secret =
  password_hash, verified_at [nullable])
- **Create** token verifikasi email di tabel `auth_tokens`
  (`purpose = email_verification`) — lihat Fitur 2B untuk detail
  struktur tabel unified ini **[RESOLVED — NEW, lihat ERD]**
**PII encryption-at-rest [RESOLVED — NEW, lihat ERD & UU PDP
compliance]**
`User.primary_email` dan `AuthIdentity.identifier` (kalau berisi
email/nomor telepon) disimpan **terenkripsi** (bukan plaintext) di
database, dengan kolom hash terpisah (`*_hash`, HMAC) untuk keperluan
unique constraint & lookup — karena ciphertext non-deterministic tidak
bisa langsung di-`WHERE`. Ini satu lapis lebih dalam dari masking di
frontend (lihat `kencleng-actors-entities.md`, PII Handling Note):
masking melindungi dari mata yang gak berwenang di UI, encryption-at-
rest melindungi kalau database itu sendiri bocor/di-dump. Detail
struktur tabel ada di `kencleng-erd.md`. **Manajemen key
(`ENCRYPTION_KEY`/`HMAC_KEY`) — lihat `kencleng-backend-tech-stack.md`
untuk detail resolved.**
 
**State transition**
`AuthIdentity.verified_at`: `null` → timestamp
 
**Concurrency & correctness notes**
Unique constraint di level DB untuk (`provider_type`, `identifier`) —
mencegah race dua registrasi dengan email yang sama persis bersamaan.
 
**Security notes**
- Password di-hash (bcrypt/argon2), tidak pernah disimpan plaintext
- Token verifikasi: random, single-use
- Rate limit lebih ketat khusus endpoint `/auth/*`
- Breach-list check: hanya prefix hash yang dikirim keluar server
  (k-anonymity), tidak pernah kirim password/hash penuh ke pihak
  ketiga
- **Fail-open kalau HaveIBeenPwned API tidak bisa diakses** — lihat
  Alternate flows di atas **[RESOLVED — NEW]**
**Open questions**
- ~~Password policy detail (panjang minimum, kompleksitas)~~ →
  **resolved: length-only, min 8 karakter** **[RESOLVED]**
- ~~Durasi expiry token verifikasi email~~ → **resolved: 24 jam**
  **[RESOLVED]**
- ~~Fail-open vs fail-closed behavior kalau HaveIBeenPwned API down~~
  → **resolved: fail-open + logging** **[RESOLVED — NEW]**
---
 
## 1B. Login/Register dengan Google (OAuth) **[NEW]**
 
**Overview**
Alternatif login/register selain `email_password`, memakai Google
sebagai identity provider. Karena Google sudah memverifikasi
kepemilikan email di sisi mereka, `AuthIdentity` yang dibuat lewat
jalur ini langsung `verified_at` terisi — tidak perlu langkah
verifikasi email tambahan seperti jalur `email_password`.
 
**Actors**
User (baru maupun existing)
 
**Preconditions**
Tidak ada — tombol "Masuk/Daftar dengan Google" tersedia baik di
halaman login maupun register, dan keduanya menuju alur yang sama
(sistem yang menentukan apakah ini user baru atau existing).
 
**Flow (happy path)**
1. User klik "Masuk/Daftar dengan Google"
2. Redirect ke Google OAuth consent screen, dengan `state` param
   (random, disimpan sementara di **HttpOnly cookie short-TTL**,
   misal 10 menit **[RESOLVED — NEW]**) untuk CSRF protection, dan
   `nonce` param (random, disimpan bersama `state` di cookie yang
   sama) untuk mencegah replay attack pada `id_token`
   **[RESOLVED — NEW]**
3. User approve consent di sisi Google → redirect balik ke
   `GOOGLE_REDIRECT_URI` dengan authorization code + `state`
4. Sistem validasi `state` cocok dengan yang disimpan sebelumnya
   (mencegah CSRF) → tukar authorization code dengan token ke Google
   → verifikasi `id_token` (signature, issuer, audience, expiry, **dan
   `nonce` claim cocok dengan yang dikirim di langkah 2** **[RESOLVED
   — NEW]**)
5. Ambil `email` dan `email_verified` dari payload `id_token` Google
6. Sistem cek: apakah sudah ada `AuthIdentity` dengan
   `provider_type = google` dan `identifier = email` tsb?
   - **Ada** → treat sebagai login, terbitkan access + refresh token
     (lanjut ke flow Login & Session Management, Fitur 2)
   - **Belum ada** → lanjut ke langkah 7 (register baru)
7. Sistem cek lagi: apakah `email` tsb sudah terdaftar sebagai
   `AuthIdentity` dengan `provider_type = email_password` (milik user
   lain/existing)?
   - **Ya** → **tidak auto-merge/auto-link** (lihat Fitur 4 Account
     Linking, business rule konsistensi) — tampilkan pesan ke user:
     "Email ini sudah terdaftar. Silakan login dengan
     password, lalu tautkan Google dari halaman keamanan akun."
   - **Tidak** → sistem membuat `User` baru + `AuthIdentity`
     (`provider_type = google`, `identifier = email`,
     `credential_secret = null`, `verified_at = now` — langsung
     verified karena Google sudah verifikasi email tsb)
8. Terbitkan access token + refresh token (sama seperti flow Login
   biasa)
**Business rules / validation**
- `email_verified` dari payload Google harus `true` — kalau Google
  sendiri belum yakin email itu valid, tolak proses (edge case jarang,
  tapi tetap dicek)
- Google-issued `AuthIdentity` langsung `verified_at = now` — tidak
  melalui flow verifikasi email manual (Fitur 1)
- `redirect_uri` divalidasi ketat (**exact match** ke
  `GOOGLE_REDIRECT_URI` yang terdaftar di Google Console per
  environment — bukan wildcard/pattern match) untuk mencegah
  open-redirect **[RESOLVED — NEW]**
- `state` param wajib divalidasi setiap callback — tanpa `state` yang
  valid, request ditolak sebelum diproses lebih lanjut
- `nonce` param wajib cocok dengan claim di dalam `id_token` — tanpa
  kecocokan ini, request ditolak (mencegah `id_token` lama/curian
  dipakai ulang) **[RESOLVED — NEW]**
**Alternate flows / edge cases**
- User cancel di consent screen Google → redirect balik ke halaman
  login dengan pesan netral ("dibatalkan"), tidak ada `User` yang
  dibuat
- `state` tidak cocok/hilang (potensi CSRF) → tolak request, log
  sebagai security event
- `nonce` tidak cocok dengan claim `id_token` (potensi replay) →
  tolak request, log sebagai security event **[RESOLVED — NEW]**
- Email dari Google match dengan `email_password` existing (langkah
  7) → **tidak** login otomatis ke akun tsb, sesuai prinsip yang sudah
  disepakati di Fitur 4 (Account Linking) — mencegah account takeover
  lewat provider yang tidak benar-benar memverifikasi kepemilikan
  secara ketat di sisi kita
**Data touched**
- **Create** `User` (kalau user baru)
- **Create** `AuthIdentity` (provider_type = `google`, identifier =
  email, credential_secret = null, verified_at = now) — kalau belum
  ada
- **Create/Update** `RefreshToken` (sama seperti flow login biasa)
**State transition**
`AuthIdentity.verified_at` (untuk record baru): langsung `now` (tidak
melalui `null` seperti jalur email_password)
 
**Concurrency & correctness notes**
Unique constraint (`provider_type`, `identifier`) yang sudah ada di
`AuthIdentity` (dari Fitur 4) juga melindungi jalur ini — mencegah dua
callback OAuth yang hampir bersamaan (misal double-click) membuat dua
`User` untuk email Google yang sama.
 
**Security notes**
- `state` param: random, single-use, short-lived — disimpan di
  **HttpOnly cookie short-TTL (±10 menit)**, konsisten dengan
  infrastruktur cookie yang sudah ada untuk refresh token, tanpa
  butuh storage/tabel DB tambahan **[RESOLVED — NEW]**
- `nonce` param: random, disimpan bersama `state` di cookie yang
  sama, divalidasi terhadap claim `nonce` di dalam `id_token` untuk
  mencegah replay attack **[RESOLVED — NEW]**
- `id_token` **wajib** diverifikasi signature-nya terhadap Google's
  JWKS (bukan cuma decode tanpa verifikasi) — pakai
  `google.golang.org/api/idtoken` atau manual JWKS verify
- `redirect_uri` harus exact-match whitelist (didaftarkan di Google
  Console per environment dev/prod), tidak boleh ada
  wildcard/partial match (mencegah open-redirect); karena
  `GOOGLE_REDIRECT_URI` fixed dari env var (bukan dinamis dari
  request user), tidak ada celah open-redirect tambahan di sisi app
  kita sendiri **[RESOLVED — NEW]**
- Client secret (`GOOGLE_CLIENT_SECRET`) tidak pernah exposed ke
  frontend — seluruh token exchange terjadi di backend
**Open questions**
- Apakah perlu handling khusus kalau Google mengembalikan email yang
  berbeda dari email yang di-`primary_email` di `User` (kasus:
  existing user link akun Google dengan email berbeda dari akun
  utamanya) — ini masuk ranah Fitur 4 (Account Linking)
- Refresh token dari Google sendiri (buat re-auth ke Google API di
  masa depan) — apakah perlu disimpan di v1, atau cukup one-time
  identity verification tanpa perlu akses API Google lanjutan?
  **[v1 default: tidak disimpan, cukup untuk identity verification saja]**
- ~~Desktop OAuth mechanism~~ → **resolved: full-page redirect untuk
  SEMUA trigger (login, register, linking), bukan popup window.**
  Konsekuensi: klik "Masuk/Daftar dengan Google" dari modal (desktop,
  lihat Fitur 2 di frontend-tech-stack.md) menavigasi browser keluar
  sepenuhnya dari modal tsb, mendarat di `/auth/google/callback`, lalu
  redirect ke destinasi akhir. Dipilih dibanding popup +
  `window.postMessage` karena lebih simpel (tidak ada popup-blocker
  edge case, tidak ada handshake tambahan) — konsisten dengan prinsip
  "lowest complexity" proyek ini. Lihat
  `kencleng-frontend-tech-stack.md` "Layout Patterns" untuk detail.
  **[RESOLVED — NEW]**
- ~~State/nonce validation detail & redirect-URI validation
  strictness~~ → **resolved: `state`+`nonce` di HttpOnly cookie
  short-TTL, redirect URI exact-match dari env var** — lihat Business
  rules & Security notes di atas **[RESOLVED — NEW]**
---
 
## 2. Login & Session Management
 
**Overview**
Proses login dan pengelolaan sesi menggunakan access token (JWT,
ES256) + refresh token dengan strategi rotate-on-use dan reuse
detection. Kalau user punya TOTP MFA aktif, ada langkah verifikasi
tambahan sebelum token diterbitkan. Login bisa lewat `email_password`
atau `google` (lihat Fitur 1B).
 
**Actors**
User
 
**Preconditions**
User sudah terdaftar (boleh belum verified — lihat business rule di
bawah; tidak berlaku untuk `google` yang selalu langsung verified)
 
**Flow (happy path) — Login (email_password)**
1. User submit email + password
2. **Cek lockout** (lihat Fitur 2C di bawah) — kalau
   `identifier_hash` sedang dalam status locked, tolak sebelum proses
   lanjut, tanpa memberitahu apakah email itu terdaftar atau tidak
   (anti-enumeration, sama seperti Fitur 2B) **[NEW]**
3. Sistem cari `AuthIdentity` dengan `provider_type = email_password`
   & `identifier = email`, verifikasi `credential_secret`
4. **Catat hasil percobaan** ke `login_attempts` (`success = true/false`)
   terlepas dari hasilnya **[NEW]**
5. Kredensial valid → cek apakah `User` punya `MfaTotpSecret` dengan
   `enabled_at` terisi?
   - **Tidak** → lanjut ke langkah 7
   - **Ya** → minta input kode TOTP (langkah 6)
6. User input kode TOTP → sistem verifikasi terhadap secret →
   valid → lanjut; tidak valid → tolak, user bisa retry atau pakai
   backup code
7. Terbitkan access token (JWT ES256, **TTL 15 menit** **[RESOLVED]**)
   + refresh token baru (hash-nya disimpan di DB dengan `family_id`
   baru, **TTL 30 hari** **[RESOLVED]**, token asli dikirim sebagai
   HttpOnly cookie)
**Flow — Refresh**
1. Access token expired → client panggil endpoint refresh dengan
   cookie refresh token
2. Sistem cek: token valid, belum revoked, belum pernah dipakai
   (`replaced_by_id IS NULL`)?
3. Valid → terbitkan access token baru + refresh token baru (rotate),
   tandai refresh token lama `replaced_by_id = <token baru>`
4. **Reuse terdeteksi** (refresh token yang sudah pernah di-rotate
   dipakai lagi) → revoke seluruh token dalam `family_id` yang sama,
   paksa user re-login (indikasi token dicuri)
**Flow — Logout**
1. Revoke refresh token saat ini di DB
2. Clear cookie di client
**Business rules / validation**
- User dengan `AuthIdentity.verified_at = null` (hanya berlaku untuk
  provider `email_password`) tetap bisa login, tapi dibatasi
  aksesnya: tidak bisa donasi sebagai registered user, tidak bisa jadi
  representative organization — endpoint terkait ini wajib cek status
  verifikasi
- MFA TOTP bersifat **opsional untuk semua role** — user memilih
  sendiri mau mengaktifkan atau tidak (lihat Fitur 3), berlaku juga
  untuk user yang login lewat Google
**Data touched**
Create/Update `RefreshToken` (id, user_id, token_hash, family_id,
issued_at, expires_at, revoked_at, replaced_by_id)
 
**State transition**
`RefreshToken`: active → rotated (`replaced_by_id` terisi) →
*(kalau reuse terdeteksi)* seluruh family di-revoke
 
**Concurrency & correctness notes**
Refresh endpoint harus atomic check-and-rotate (guard dengan
`WHERE replaced_by_id IS NULL AND revoked_at IS NULL`) supaya tidak
ada race dua refresh request memakai refresh token yang sama.
 
**Security notes**
- ES256: private key untuk sign, public key untuk verify —
  future-proof kalau nanti ada service lain yang perlu verify token
- Refresh token di-hash sebelum disimpan
- Cookie: HttpOnly + Secure + SameSite
- Rate limit percobaan login & verifikasi TOTP (mencegah brute-force)
**Open questions**
- ~~Access token TTL spesifik (menit)~~ → **resolved: 15 menit**
  **[RESOLVED]**
- ~~Refresh token TTL spesifik (hari)~~ → **resolved: 30 hari**
  **[RESOLVED]**
---
 
## 2B. Forgot & Reset Password **[NEW]**
 
**Overview**
Alur user yang lupa password untuk `email_password` login, meminta
reset via link yang dikirim ke email terdaftar. Tidak berlaku untuk
user yang hanya punya `AuthIdentity` provider `google` (tidak ada
password untuk direset) — lihat business rule khusus di bawah.
 
**Actors**
User (yang punya `AuthIdentity` provider `email_password`)
 
**Preconditions**
Tidak ada (endpoint publik — siapa saja bisa request, tapi hasil
diproses hanya kalau email match)
 
**Flow (happy path)**
1. User buka halaman "Forgot Password", isi email
2. Sistem cari `AuthIdentity` dengan `provider_type = email_password`
   & `identifier = email` tsb
3. Kalau ketemu → generate token reset (random, single-use, expired
   dalam **1 jam** **[RESOLVED]**), kirim email berisi link reset
   (`/reset-password?token=...`)
4. Kalau tidak ketemu (email tidak terdaftar sebagai
   `email_password`) → sistem **tetap menampilkan pesan sukses yang
   sama** ("kalau email terdaftar, link reset sudah dikirim") — untuk
   mencegah **user enumeration** (jangan bocorkan apakah email
   terdaftar atau tidak lewat pesan yang berbeda)
5. User klik link, buka halaman reset password, isi password baru
   (tetap dicek lewat breach-list HaveIBeenPwned, sama seperti
   registrasi — lihat Fitur 1, termasuk behavior **fail-open** kalau
   API tidak bisa diakses **[RESOLVED — NEW]**) **[REVISED]**
6. Sistem validasi token (belum expired, belum dipakai) → update
   `credential_secret` (hash password baru) → tandai token used
**Business rules / validation**
- Token reset: single-use, expired dalam **1 jam [RESOLVED]**
- Setelah reset berhasil, **semua refresh token existing milik user
  tsb di-revoke** (force logout dari semua device/sesi lama) — praktik
  standar untuk mitigasi kalau alasan reset adalah akun dicurigai
  diakses orang lain
- Password baru tetap harus memenuhi kriteria minimum yang sama
  dengan registrasi (Fitur 1) — panjang minimum 8 karakter + breach-
  list check (fail-open kalau API down)
- **User yang hanya punya `AuthIdentity` provider `google`
  [RESOLVED — NEW]**: request forgot-password dengan email tsb tetap
  menampilkan pesan sukses generik yang sama (demi anti-enumeration),
  **dan sistem mengirim email** ke alamat tsb berisi pemberitahuan:
  "Akun ini terdaftar lewat Google, silakan login dengan Google" —
  bukan link reset password. Ini tetap aman dari sisi
  anti-enumeration (attacker yang tidak punya akses ke inbox tsb tidak
  mendapat informasi apa pun dari response API; informasi hanya
  sampai ke pemilik inbox asli), sekaligus membantu user yang lupa
  metode login-nya
**Alternate flows / edge cases**
- Token expired saat diklik → tampilkan pesan jelas, tawarkan request
  ulang
- Double-submit request forgot-password → tidak masalah, cukup kirim
  ulang token baru (token lama otomatis usang begitu ada yang baru,
  atau tetap valid sampai expired — detail idempotency ini diserahkan
  ke implementasi, karena bukan financial-critical path)
**Data touched**
- **Create** token reset di tabel `auth_tokens` (`purpose =
  password_reset`) — **~~tabel terpisah atau reuse pola token
  verifikasi email~~ → resolved: satu tabel unified `auth_tokens`
  dengan kolom `purpose` enum (`email_verification`,
  `password_reset`), karena shape-nya identik (random single-use
  token + expiry), cuma beda durasi & trigger. Juga ada kolom
  `revoked_at` (dipakai kalau butuh invalidate token lama secara
  manual, misal saat resend). Detail struktur di `kencleng-erd.md`**
  **[RESOLVED — NEW, lihat ERD]**
- **Update** `AuthIdentity.credential_secret` (provider_type =
  `email_password`)
- **Update/Revoke** semua `RefreshToken` milik user tsb
- **[NEW]** Create notification/email log record untuk kasus
  Google-only (channel email, type baru misal
  `forgot_password_google_only_notice`)
**State transition**
Reset token: *(tidak ada)* → issued → used (atau expired tanpa
dipakai)
 
**Concurrency & correctness notes**
Guard update dengan `WHERE used_at IS NULL AND expires_at > now()`
saat proses reset — mencegah token yang sama dipakai dua kali kalau
request submit dobel.
 
**Security notes**
- **Anti user-enumeration**: response selalu sama ("email dikirim
  kalau terdaftar") terlepas dari apakah email itu benar-benar
  terdaftar — termasuk untuk kasus Google-only, response API tetap
  generik meski isi email yang benar-benar terkirim berbeda
- Token reset: random, cukup panjang, single-use, short-lived
- Reset password sukses → revoke semua sesi lama (lihat business rule
  di atas)
- Rate limit endpoint ini juga masuk kategori `/auth/*` yang sudah
  disepakati stricter rate limit-nya
**Open questions**
- ~~Durasi expiry token reset~~ → **resolved: 1 jam** **[RESOLVED]**
- ~~Behavior spesifik untuk user yang cuma punya Google identity~~ →
  **resolved: kirim email pemberitahuan "pakai Google login"**
  **[RESOLVED]**
---
 
## 2C. Login Attempt Lockout **[NEW]**
 
**Overview**
Proteksi brute-force yang persisten di database — melengkapi rate
limit in-memory (`golang.org/x/time/rate`) yang sudah ada di level
endpoint `/auth/*`. Rate limit in-memory mencegah flood dalam window
pendek tapi hilang saat restart dan tidak persisten cross-instance;
`login_attempts` menyimpan histori percobaan supaya lockout tetap
berlaku meski aplikasi restart.
 
**Actors**
System (record & evaluasi setiap percobaan login)
 
**Preconditions**
Tidak ada — berlaku untuk setiap percobaan login `email_password`
(tidak berlaku untuk `google`, karena kredensial diverifikasi Google,
bukan oleh sistem kita)
 
**Flow (happy path)**
1. Setiap percobaan login (Fitur 2, langkah 2–4) dicatat sebagai satu
   row `login_attempts` — `identifier_hash` (hash dari email yang
   dicoba, konsisten dicatat meski emailnya sendiri tidak terdaftar,
   demi anti-enumeration), `user_id` (nullable, null kalau identifier
   tidak match akun manapun), `success`, `attempted_at`
2. Sebelum credential diverifikasi, sistem hitung jumlah percobaan
   gagal berturut-turut untuk `identifier_hash` tsb dalam window waktu
   **15 menit** **[RESOLVED — NEW]**
3. Kalau melewati threshold **5 kali** **[RESOLVED — NEW]** → tolak
   percobaan (locked), tanpa membedakan pesan dari kasus "password
   salah" biasa (anti-enumeration tetap konsisten dengan Fitur 2B)
**Business rules / validation**
- Status lockout **dihitung on-the-fly** dari `COUNT(*)` terhadap
  `login_attempts` (`WHERE identifier_hash = ? AND success = false AND
  attempted_at > now() - interval '15 minutes'` **[RESOLVED — NEW]**)
  — **bukan** disimpan sebagai kolom `locked_until` terpisah, supaya
  tidak ada dua sumber kebenaran yang bisa desinkron (konsisten dengan
  alasan kita skip `deleted_at` generik — satu sumber kebenaran per
  pertanyaan)
- **Threshold: 5 kali percobaan gagal berturut-turut, window: 15
  menit [RESOLVED — NEW]** — dipilih karena titik tengah yang umum
  dipakai (OWASP merekomendasikan rentang 3-5 kali): cukup ketat buat
  bikin brute-force tidak praktis, tapi user asli yang salah ketik
  beberapa kali tidak lama-lama ter-block
- Lockout berlaku **per `identifier_hash` saja**, tidak per-IP —
  device/IP fingerprinting sengaja tidak diimplementasikan di v1
  (lihat Open Questions di bawah)
- Login sukses tidak menghapus histori attempt sebelumnya — histori
  tetap utuh untuk keperluan audit, cukup window waktu yang membuat
  attempt lama gak lagi dihitung ke lockout
**Data touched**
Create `login_attempts` (identifier_hash, user_id [nullable],
success, attempted_at)
 
**Concurrency & correctness notes**
Read (COUNT) lalu write (INSERT) tidak butuh locking khusus — race
antara dua percobaan hampir bersamaan paling buruk cuma bikin lockout
ke-trigger satu attempt lebih telat/cepat dari seharusnya, bukan
celah keamanan (fail-safe, bukan fail-open di sisi yang salah)
 
**Security notes**
- `identifier_hash` pakai fungsi hash yang sama dengan pola
  encryption-at-rest lain (lihat Fitur 1, PII encryption-at-rest) —
  bukan plaintext email
- Pesan error saat locked tetap generik, tidak membocorkan apakah
  akun tsb ada atau tidak
**Open questions**
- ~~Threshold: berapa kali gagal berturut-turut sebelum lockout?~~ →
  **resolved: 5 kali** **[RESOLVED — NEW]**
- ~~Window: dalam rentang berapa menit?~~ → **resolved: 15 menit**
  **[RESOLVED — NEW]**
- **Device/IP fingerprinting (user-agent & IP address) sengaja
  TIDAK diimplementasikan** — dipertimbangkan saat diskusi ERD,
  disepakati sebagai over-engineering untuk kebutuhan Kencleng saat
  ini (tidak ada fitur "sesi aktif per device" atau anomaly detection
  yang membutuhkannya). Bisa direvisit kalau ada fitur konkret ke
  arah sana di masa depan. **[RESOLVED — NEW, decided against]**
---
 
## 3. TOTP MFA — Enrollment & Verification
 
**Overview**
User bisa mengaktifkan Multi-Factor Authentication berbasis TOTP
(RFC 6238) secara opsional, terlepas dari provider login yang dipakai
(`email_password` maupun `google`). Dilengkapi backup code untuk
skenario kehilangan device authenticator.
 
**Actors**
User
 
**Preconditions**
User sudah login, `AuthIdentity` utama sudah verified
 
**Flow (happy path) — Enrollment**
1. User membuka halaman keamanan akun, pilih "Aktifkan MFA"
2. Sistem generate TOTP secret baru, tampilkan sebagai QR code
   (`otpauth://` URI)
3. User scan dengan authenticator app (Google Authenticator, Authy,
   dll), input satu kode untuk konfirmasi
4. Sistem verifikasi kode cocok dengan secret →
   `MfaTotpSecret.enabled_at = now`
5. Sistem generate **10 backup code** **[RESOLVED]**, ditampilkan
   **satu kali** ke user (disimpan hash-nya saja di server)
**Flow — Disable MFA**
1. User (dengan re-autentikasi, misal input password lagi — **untuk
   user Google-only tanpa password, re-autentikasi via re-login
   Google/OAuth prompt** **[REVISED — klarifikasi untuk Google users]**)
   pilih "Nonaktifkan MFA"
2. `MfaTotpSecret.enabled_at = null`, backup code lama diinvalidasi
**Flow — Regenerate Backup Code [RESOLVED — NEW]**
Regenerate backup code **hanya bisa dilakukan lewat siklus
disable → enable ulang** — tidak ada jalur regenerate langsung tanpa
disable MFA dulu. Ini konsisten dengan business rule enrollment yang
sudah ada: MFA hanya boleh dianggap "aktif" setelah kode TOTP baru
benar-benar diverifikasi ulang, mencegah state ambigu "backup code
baru tapi TOTP secret lama belum tentu masih valid di device user."
 
**Flow — Pakai Backup Code (saat login, device hilang)**
1. Di langkah verifikasi TOTP saat login, user pilih "pakai backup
   code" sebagai alternatif
2. Input salah satu backup code → dicocokkan dengan hash tersimpan →
   valid → code tsb langsung ditandai `used_at` (single-use, tidak
   bisa dipakai lagi)
**Business rules / validation**
- Backup code bersifat single-use
- Enrollment butuh konfirmasi kode TOTP yang valid (bukan cuma
  generate & simpan tanpa verifikasi), supaya tidak ada kondisi
  "MFA aktif tapi secret salah/QR gagal di-scan"
- **Jumlah backup code per enrollment: 10 [RESOLVED]**
- **Regenerate backup code: hanya lewat disable-enable ulang, tidak
  ada regenerate langsung [RESOLVED]**
**Data touched**
- Create/Update `MfaTotpSecret` (user_id, secret [encrypted],
  enabled_at [nullable])
- Create `MfaBackupCode` (user_id, code_hash, used_at [nullable])
**State transition**
`MfaTotpSecret.enabled_at`: `null` → timestamp → *(disable)* `null`
 
**Security notes**
- Secret TOTP disimpan terenkripsi di database (bukan plaintext)
- Backup code disimpan sebagai hash, ditampilkan ke user hanya sekali
  saat generate
- Disable MFA wajib re-autentikasi (mencegah orang lain menonaktifkan
  MFA hanya dari sesi yang sedang login)
**Open questions**
- ~~Jumlah backup code yang digenerate per enrollment~~ →
  **resolved: 10** **[RESOLVED]**
- ~~Apakah backup code bisa di-regenerate ulang tanpa disable-enable
  dulu~~ → **resolved: tidak, hanya lewat disable-enable**
  **[RESOLVED]**
- Mekanisme re-autentikasi spesifik untuk user Google-only saat mau
  disable MFA (masih via re-login Google/OAuth prompt sesuai flow di
  atas — detail teknis implementasi diserahkan ke coding phase)
---
 
## 4. Account Linking (Multi Identity Provider)
 
**Overview**
User bisa menautkan lebih dari satu `AuthIdentity` ke akun `User`
yang sama (misal sudah daftar pakai email+password, lalu ingin login
juga bisa lewat Google — **ini sekarang jalur yang sudah diimplementasi
penuh di v1, bukan sekadar dimodelkan** **[REVISED]**). Linking
dilakukan **manual**, tidak otomatis berdasarkan kecocokan email —
konsisten dengan pola klaim donasi guest yang sudah disepakati.
 
**Actors**
User (sudah login)
 
**Preconditions**
User sudah login dengan salah satu `AuthIdentity` yang verified
 
**Flow (happy path) — Link Google ke akun email_password existing**
1. User (sudah login pakai email_password) membuka halaman keamanan
   akun, pilih "Tautkan Google"
2. User melalui alur OAuth Google (sama seperti Fitur 1B langkah 2-5)
3. Sistem cek: apakah `identifier` (email dari Google) sudah dipakai
   `AuthIdentity` lain (milik user lain)?
   - **Ya** → tolak, tampilkan pesan jelas
   - **Tidak** → buat `AuthIdentity` baru (`provider_type = google`,
     `verified_at = now`), ditautkan ke `user_id` yang sedang login —
     **bukan** membuat `User` baru
**Flow (happy path) — Set Password untuk Google-only user
[RESOLVED — NEW]**
Kebalikan dari flow linking di atas: user yang registrasi lewat
Google saja dan belum punya `AuthIdentity` `email_password` bisa
menambahkan satu, supaya kelak bisa unlink Google kalau mau (lihat
business rule unlink di bawah).
1. User (login via Google-only) membuka `/dashboard/security`, pilih
   "Atur Password"
2. User isi password baru → divalidasi dengan kriteria yang sama
   seperti registrasi (Fitur 1): minimum 8 karakter + breach-list
   check HaveIBeenPwned (fail-open kalau API down)
3. Sistem cek: apakah `User.primary_email` sudah dipakai
   `AuthIdentity` `email_password` milik user lain?
   - **Ya** → tolak, tampilkan pesan jelas (edge case jarang, tapi
     tetap ditangani secara eksplisit — bukan silent fail)
   - **Tidak** → buat `AuthIdentity` baru (`provider_type =
     email_password`, `identifier = User.primary_email`,
     `credential_secret = password_hash`, **`verified_at = now`
     langsung** — tidak perlu verifikasi email ulang, karena Google
     sudah memverifikasi email ini duluan dan user sudah berada dalam
     sesi authenticated)
4. User sekarang punya 2 `AuthIdentity` (google + email_password) —
   tombol unlink Google jadi tersedia (tidak lagi satu-satunya auth
   method)
**Business rules / validation**
- Kalau `identifier` dari provider baru **ternyata sudah terdaftar**
  sebagai `AuthIdentity` milik `User` lain → tolak proses linking,
  beri pesan jelas (mencegah satu identifier dipakai lebih dari satu
  akun) — berlaku juga untuk flow set-password di atas
- Tidak ada auto-link berdasarkan email yang sama — harus lewat alur
  eksplisit ini
- **Unlink Google [RESOLVED — NEW]**: unlink `AuthIdentity` google
  **diblokir** kalau itu satu-satunya `AuthIdentity` milik user
  tersebut (tidak punya `email_password` yang aktif). User diarahkan
  untuk set up email+password dulu sebelum bisa unlink — memastikan
  user selalu punya minimal satu cara login yang berfungsi, konsisten
  dengan prinsip yang sama yang mendasari desain di dokumen ini.
- **Set-password untuk Google-only user tidak membutuhkan
  re-autentikasi tambahan [RESOLVED — NEW]** — aksi ini *menambah*
  metode auth (menaikkan security posture user), berbeda dengan
  disable-MFA atau unlink yang *mengurangi* proteksi akun dan karena
  itu butuh proteksi ekstra. Sesi yang sudah authenticated dianggap
  cukup.
**Alternate flows / edge cases**
- User mencoba login dengan provider yang identifier-nya cocok
  dengan email `User` lain yang sudah ada, tapi belum pernah
  ditautkan → sistem **tidak** otomatis meng-cocokkan/merge; user
  diarahkan untuk login dengan metode existing dulu, baru menautkan
  dari halaman akun (mencegah account takeover lewat provider yang
  tidak benar-benar memverifikasi kepemilikan) — **ini persis
  behavior yang dijelaskan di Fitur 1B langkah 7**
- **Set-password untuk Google-only user, tapi `primary_email` sudah
  dipakai `email_password` milik user lain** → tolak, pesan jelas
  (lihat Flow di atas langkah 3) **[RESOLVED — NEW]**
**Data touched**
Create `AuthIdentity` (user_id, provider_type, identifier,
credential_secret [nullable, tergantung provider], verified_at)
 
**Security notes**
Unique constraint (`provider_type`, `identifier`) di `AuthIdentity`
mencegah satu identifier ditautkan ke lebih dari satu `User` —
constraint yang sama ini juga menangkap edge case konflik di flow
set-password.
 
**Open questions**
- Phone+OTP: model data (`AuthIdentity` dengan `provider_type` baru)
  sudah siap, implementasi flow spesifik di-defer ke versi berikutnya
  **(Google OAuth sendiri sudah tidak lagi open question — resolved,
  in-scope v1)**
- ~~Set-password flow untuk Google-only user~~ → **resolved: via
  `/dashboard/security` "Atur Password", `verified_at = now` langsung,
  tanpa re-auth tambahan, dengan edge case konflik identifier
  ditangani** — lihat Flow & Business rules di atas
  **[RESOLVED — NEW]**
---
 
## 5. Role Bootstrapping & Assignment
 
**Overview**
Mekanisme user mendapatkan role Admin atau Kurator — tidak
self-service seperti Donatur/Representative, harus lewat assignment
eksplisit oleh Admin.
 
**Actors**
Super-admin (seed awal) · Admin (assign/revoke role ke user lain)
 
**Preconditions**
- Seed awal: dijalankan manual saat setup/deploy (bukan lewat endpoint
  publik)
- Assignment/revoke selanjutnya: dilakukan oleh user yang sudah
  berstatus Admin
**Flow (happy path)**
1. Saat setup awal, seed script/migration membuat satu user dengan
   role Admin (super-admin) langsung di database
2. Admin membuka halaman kelola user, pilih user terdaftar, assign
   role Admin atau Kurator
3. Admin juga bisa **revoke/demote** role Admin atau Kurator dari user
   lain
**Business rules / validation**
- User yang akan di-assign Admin **tidak boleh** sedang jadi Kurator
  atau Representative organization mana pun — sistem harus validasi &
  block pelanggaran ini
- User yang akan di-assign Kurator boleh saja representative organization
  lain (konfliknya baru muncul per-assignment kurasi)
**Data touched**
Update role User (struktur data pasti — many-to-many User-Role vs
flag boolean — diputuskan nanti di ERD)
 
**Security notes**
Hanya Admin yang bisa akses endpoint assign/revoke role.
 
**Open questions**
- Konsekuensi revoke role terhadap data yang sudah dibuat user tsb
  (misal Kurator yang di-demote — assignment kurasi yang sedang
  `pending` miliknya diapakan?)
---
 
## 6. Notification Infrastructure
 
**Overview**
Layanan terpusat untuk mengirim & menampilkan notifikasi ke user,
dengan notification center UI tersendiri (bisa mark-as-read). Mark-
as-read diproses **asynchronous**, dan notifikasi punya masa
kedaluwarsa untuk menjaga list tetap ringkas dan DB tetap efisien
**[REVISED]**.
 
**Actors**
System (trigger notifikasi) · User (melihat & mark-as-read)
 
**Flow (happy path)**
1. Event tertentu terjadi di sistem (organization verified, campaign
   closed, disbursement selesai, dst) → sistem buat `Notification`
   record dengan `expires_at` (created_at + **30 hari**
   **[RESOLVED]** — lihat business rule)
2. Channel in-app → notifikasi muncul di notification center user
   (badge unread count)
3. Channel email → sistem "kirim" (fake/logged untuk sandbox) ke
   `user.primary_email` atau `guest_email`
4. User bisa membuka notification center, mark notifikasi sebagai
   read
**Business rules / validation**
- Setiap notifikasi punya `type` (enum event, misal
  `organization_verified`, `campaign_rejected`, `donation_success`,
  `campaign_closed`, `disbursement_completed`) untuk konsistensi
  template pesan. **[RESOLVED — NEW]** Diperluas buat cover seluruh
  siklus kurasi: `admin_new_curation_item` (Admin, saat organization/
  campaign/laporan dana baru masuk antrian), `kurator_assigned`
  (Kurator, saat di-assign Admin ke satu item), `fund_usage_report_
  verified`/`fund_usage_report_rejected` (Owner, keputusan Kurator),
  `disbursement_approved`/`disbursement_rejected` (Owner, keputusan
  Admin). Semua pakai mekanisme & infrastruktur yang sama persis di
  bawah — cuma nilai string `type` baru, bukan mekanisme notifikasi
  baru. Channel: dual (in-app + email), konsisten dengan tipe
  notifikasi lain yang udah ada.
- **Mark-as-read tidak diproses secara synchronous per klik.**
  Interaksi user (klik/buka notifikasi) dikumpulkan di sisi client
  (misal debounced/batched), lalu dikirim sebagai satu batch request
  ke backend — mencegah request-per-notifikasi yang tidak perlu kalau
  user punya ribuan notifikasi belum dibaca.
- **Notification expiration — dua lapis**:
  1. **Soft-hide (logical)**: query list notifikasi selalu
     memfilter `WHERE expires_at > now()` — notifikasi yang sudah
     lewat masa berlakunya (**30 hari sejak `created_at`
     [RESOLVED]**, terlepas dari status read/unread) langsung tidak
     muncul di list, tanpa perlu physical delete dulu.
  2. **Hard-delete (physical)**: background worker berjalan
     **mingguan** **[RESOLVED]**, menghapus permanen row
     `Notification` yang `expires_at` sudah lewat, untuk menjaga
     ukuran tabel tetap efisien dalam jangka panjang. Ini murni
     housekeeping — user sudah tidak melihat row ini sejak soft-hide
     aktif, jadi hard-delete tidak mengubah pengalaman user, hanya
     membebaskan storage.
**Data touched**
Create `Notification` (recipient_user_id [nullable], recipient_email
[nullable, untuk guest], channel, type, payload, `read_at` [nullable],
`expires_at`, created_at)
 
**Concurrency & correctness notes**
- Batch mark-as-read: idempotent (guard `WHERE read_at IS NULL` per
  row saat update batch), supaya aman kalau batch yang sama terkirim
  dua kali (misal retry dari client)
- Hard-delete worker: query delete pakai kondisi `WHERE expires_at <
  now()`, tidak perlu locking khusus — notifikasi yang sudah expired
  tidak lagi relevan untuk dibaca ulang oleh siapa pun, jadi race
  antara "user baca notif yang barusan expired" vs "worker hapus
  notif itu" secara praktis tidak masalah (soft-hide sudah
  menyembunyikannya duluan dari list sejak `expires_at` lewat)
**Security notes**
User hanya bisa melihat & mark-as-read notifikasi miliknya sendiri.
 
**Open questions**
- ~~Durasi expiration pasti~~ → **resolved: 30 hari** **[RESOLVED]**
- ~~Frekuensi run hard-delete worker~~ → **resolved: mingguan**
  **[RESOLVED]**
- Detail batching mark-as-read: interval debounce di client, atau
  trigger saat user leave halaman notification center? Diserahkan ke
  implementasi.
---
 
## 7. File Upload / Storage
 
**Overview**
Layanan upload file terpusat memakai storage S3-compatible (misal
MinIO lokal untuk sandbox), dipakai untuk dokumen legal organization,
media campaign, dan lampiran laporan penggunaan dana. Setiap form
upload dokumen sensitif menampilkan notifikasi bahwa file tersimpan
aman dan rahasia **[NEW]**.
 
**Actors**
User (upload) · System (validasi & simpan)
 
**Preconditions**
User authenticated (sesuai konteks upload — representative organization
untuk dokumen legal/media campaign, Owner untuk lampiran laporan dana)
 
**Flow (happy path)**
1. Client upload file ke endpoint upload
2. Sistem validasi tipe file (whitelist per konteks: dokumen legal →
   PDF/JPG/PNG; media campaign → JPG/PNG; lampiran laporan dana →
   PDF/JPG/PNG) & ukuran maksimum **5 MB** **[RESOLVED]**
3. File disimpan ke bucket MinIO yang sesuai (lihat business rule
   akses di bawah), record metadata dibuat
4. Return reference (object key) untuk disimpan di entity terkait
**Business rules / validation**
- **Media campaign** → bucket publicly accessible
- **Dokumen legal organization & lampiran laporan penggunaan dana** →
  bucket private, akses hanya lewat signed URL dengan waktu
  kedaluwarsa **5 menit** **[RESOLVED]**, terbatas untuk
  Admin/Kurator yang berwenang (representative organization pemilik
  data tetap bisa akses miliknya sendiri — dan khusus dokumen legal,
  hanya level `owner`, bukan `staff`)
- **Max ukuran file: 5 MB, berlaku sama untuk seluruh konteks
  (dokumen legal, media campaign, lampiran laporan dana)
  [RESOLVED]** — dikonfigurasi per-konteks di kode (bukan satu
  konstanta global), supaya tetap mudah dibedakan nanti kalau ada
  kebutuhan konkret
- Validasi tipe file dari magic bytes, bukan hanya ekstensi/
  content-type yang dikirim client
- **UX note**: setiap form upload dokumen non-publik (dokumen
  legal organization, lampiran laporan penggunaan dana) menampilkan
  notes/popup yang menjelaskan bahwa file tersimpan aman dan bersifat
  rahasia (mengacu ke mekanisme private bucket + signed URL di atas)
  — ini murni UX reassurance, tidak mengubah mekanisme teknis yang
  sudah ada
**Data touched**
Create `FileUpload` (id, object_key, content_type, size, uploaded_by,
context [misal `organization_legal_doc`, `campaign_media`,
`fund_usage_attachment`], is_public)
 
**Security notes**
File privat tidak boleh punya URL yang bisa ditebak/diakses langsung
tanpa signed URL.
 
**Open questions**
- ~~Ukuran maksimum file per konteks (angka spesifik)~~ →
  **resolved: 5 MB untuk semua konteks** **[RESOLVED]**
- ~~Durasi kedaluwarsa signed URL~~ → **resolved: 5 menit**
  **[RESOLVED]**
---
 
## 8. Terms & Conditions / Privacy Policy
 
**Overview**
User wajib menyetujui ToS & Privacy Policy versi terkini saat
registrasi. Kalau versi berubah, user existing diingatkan untuk
re-accept, tapi tidak diblokir dari memakai platform.
 
**Actors**
User · Admin (publish versi baru)
 
**Flow (happy path)**
1. Admin publish versi baru ToS/Privacy Policy
2. Saat registrasi, user centang setuju versi ToS yang berlaku saat
   itu → dicatat sebagai `UserTermsAgreement`
3. Kalau versi baru dipublish setelah user terdaftar, sistem
   menampilkan **reminder non-blocking** (misal banner) saat user
   login, mengajak re-accept — user tetap bisa memakai seluruh fitur
   platform meski belum re-accept
**Data touched**
- Create `TermsVersion` (id, version_number, content, published_at)
- Create `UserTermsAgreement` (user_id, terms_version_id, agreed_at)
**Business rules / validation**
Reminder bersifat non-blocking — tidak ada fitur yang dikunci karena
user belum re-accept versi terbaru.
 
---
 
## 9. Audit Log
 
**Overview**
Pencatatan aksi-aksi sensitif di platform untuk keperluan audit trail,
mengingat beberapa keputusan di Fase 1–3 punya konsekuensi finansial
atau legal.
 
**Actors**
System (record otomatis setiap aksi sensitif terjadi)
 
**Scope aksi yang dicatat** (berdasarkan yang teridentifikasi di
Fase 1–3 dan doc ini)
- Keputusan kurasi organization / campaign / laporan penggunaan dana
  (approve/reject)
- Force-close campaign oleh Admin
- Approval/rejection request pencairan dana
- Role assignment & revoke (Admin/Kurator)
- Enable/disable MFA, account linking baru (termasuk link Google,
  dan set-password untuk Google-only user **[NEW]**)
- **Manage representative organization (invite/remove/promote/demote
  owner↔staff) [RESOLVED — NEW]** — ditambahkan ke scope karena aksi
  ini menentukan siapa yang punya otoritas atas organization. Lihat
  `kencleng-phase1-detail.md` Fitur "Manage Representative". **Tidak**
  diperluas ke aksi non-destruktif seperti submit awal campaign/
  organization — itu udah punya trail sendiri lewat field `status`,
  nggak butuh audit log terpisah.
- **Reveal PII field yang di-mask di frontend, saat dilakukan oleh
  Admin atau Kurator terhadap data pihak lain** **[NEW — lihat
  kencleng-actors-entities.md, PII Handling Note]**
- **[NEW]** Organization kena/lepas flag "laporan penggunaan dana
  telat" (lihat `kencleng-phase3-detail.md` Fitur 4) — perubahan flag
  ini punya konsekuensi fungsional (blokir buat campaign baru), jadi
  masuk kategori aksi yang perlu tercatat
- **[NEW — RESOLVED 2026-07-21]** Unpublish campaign oleh Owner
  (manual, wajib `decision_note`) — lihat `kencleng-phase1-detail.md`
  Fitur 5
- **[NEW — RESOLVED 2026-07-21]** Auto-unpublish campaign akibat
  organization kembali ke `pending_verification` (system-triggered,
  tanpa `decision_note` manual) — lihat `kencleng-phase1-detail.md`
  Fitur 5
**Data touched**
Create `AuditLog` (actor_user_id, action_type, target_entity_type,
target_entity_id, metadata [json], created_at)
 
**Security notes**
Audit log bersifat **append-only** — tidak bisa diedit atau dihapus
oleh siapa pun, termasuk Admin, demi menjaga integritas catatan.
 
**Open questions**
- ~~Apakah scope ini perlu diperluas ke aksi lain di iterasi
  berikutnya~~ → **resolved 2026-07-24: ya, tambah representative
  management actions** (lihat Scope aksi di atas). Aksi non-
  destruktif (submit awal, dll) sengaja tidak ditambahkan — sudah
  cukup di-track lewat field `status`. **[RESOLVED — NEW]**
---
 
## Open Items Carried Forward
 
- API contract format — ~~pending dari backend tech stack doc~~ →
  **resolved: OpenAPI 3.x spec-first** — lihat
  `kencleng-backend-tech-stack.md` **[RESOLVED]**
- ~~Password policy detail & durasi expiry token verifikasi email~~ →
  **resolved [RESOLVED]**
- ~~Access token & refresh token TTL spesifik~~ → **resolved
  [RESOLVED]**
- Konsekuensi revoke role terhadap data yang sedang diproses
- ~~Ukuran maksimum file per konteks & durasi kedaluwarsa signed
  URL~~ → **resolved [RESOLVED]**
- ~~Jumlah & regenerasi backup code MFA~~ → **resolved [RESOLVED]**
- Phone+OTP: model `AuthIdentity` sudah siap, flow spesifik belum
  didesain (deferred) — Google OAuth sudah resolved, tidak lagi di
  list ini
- ~~Durasi expiry token reset password~~ → **resolved: 1 jam
  [RESOLVED]**
- ~~Behavior forgot-password untuk user Google-only~~ → **resolved:
  kirim email pemberitahuan [RESOLVED]**
- Mekanisme re-autentikasi disable-MFA untuk user Google-only — masih
  via re-login Google/OAuth prompt, detail teknis diserahkan ke
  implementasi
- ~~Durasi notification expiration pasti & frekuensi hard-delete
  worker~~ → **resolved [RESOLVED]**
- ~~Fail-open vs fail-closed kalau HaveIBeenPwned API down saat
  registrasi/reset password~~ → **resolved: fail-open + logging**
  **[RESOLVED — NEW]**
- ~~Desktop OAuth mechanism (popup vs full-page redirect)~~ →
  **resolved: full-page redirect untuk semua trigger** **[RESOLVED —
  NEW, lihat Fitur 1B]**
- ~~Unlink Google business rule~~ → **resolved: diblokir kalau
  satu-satunya AuthIdentity** **[RESOLVED — NEW, lihat Fitur 4]**
- ~~Set-password flow untuk Google-only user~~ → **resolved: via
  `/dashboard/security` "Atur Password", tanpa re-auth tambahan**
  **[RESOLVED — NEW, lihat Fitur 4]**
- **New dependencies to add to `kencleng-backend-tech-stack.md`**:
  MinIO client (S3-compatible storage), `pquerna/otp` (TOTP),
  `golang.org/x/oauth2` + Google idtoken verification, plain HTTP
  client untuk HaveIBeenPwned API (sudah ditambahkan — lihat
  backend-tech-stack.md revisi), **`github.com/shopspring/decimal`
  untuk field nilai uang [NEW, lihat ERD]**, dan library AES-GCM
  (stdlib `crypto/aes`, tanpa dependency tambahan) + HMAC untuk
  encryption-at-rest field PII **[NEW, lihat ERD]**
- ~~Struktur tabel token verifikasi email & password reset~~ →
  **resolved: satu tabel unified `auth_tokens` dengan kolom `purpose`
  enum, termasuk kolom `revoked_at`** **[RESOLVED — NEW, lihat ERD,
  Fitur 1 & 2B]**
- **PII encryption-at-rest (email, NPWP) [RESOLVED — NEW, lihat ERD]**:
  `User.primary_email`, `AuthIdentity.identifier`, `Organization.NPWP`,
  `Donation.guest_email` disimpan terenkripsi + kolom hash terpisah
  untuk lookup, sesuai kepatuhan UU PDP — lihat Fitur 1. **Manajemen
  key (`ENCRYPTION_KEY`/`HMAC_KEY`, no rotation di v1) — resolved,
  lihat `kencleng-backend-tech-stack.md` [RESOLVED — NEW]**
- ~~Login attempt lockout / brute-force protection persisten — tabel
  `login_attempts` ditambahkan sebagai pelengkap rate limit in-memory
  yang sudah ada, lihat Fitur 2C baru. Threshold & window lockout
  masih open, tidak blocking ERD/schema~~ → **resolved: threshold 5
  kali gagal / window 15 menit** — lihat Fitur 2C **[RESOLVED — NEW]**
- **Device/IP fingerprinting untuk session — resolved: tidak
  diimplementasikan** [RESOLVED — NEW, lihat Fitur 2C Open Questions]
  — dipertimbangkan saat diskusi ERD (terinspirasi dari referensi
  schema auth project lama), disepakati over-engineering untuk
  kebutuhan Kencleng saat ini
- ~~Google OAuth state/nonce validation detail dan redirect-URI
  validation strictness~~ → **resolved: `state`+`nonce` di HttpOnly
  cookie short-TTL, redirect URI exact-match dari env var** — lihat
  Fitur 1B **[RESOLVED — NEW]**
- **`otps` table [NEW]** — disiapkan sebagai tabel fisik di ERD untuk
  mendukung `AuthIdentity.provider_type = phone_otp` di masa depan,
  konsisten dengan status "model data sudah siap, flow implementasi
  di-defer" yang sudah disebutkan di atas — belum dipakai endpoint
  manapun di v1
- ~~Notification mechanism (email/in-app) untuk Admin/Kurator/
  Organization across curation steps~~ → **resolved: extend
  `notifications.type` enum values (no new mechanism), dual channel
  konsisten dengan tipe lain** — lihat Fitur 6 di atas
  **[RESOLVED — NEW]**
- ~~Audit-trail granularity untuk aksi sensitif di luar yang sudah
  dispesifikasi~~ → **resolved: tambah representative management
  actions (invite/remove/promote/demote), aksi non-destruktif
  sengaja tidak ditambahkan** — lihat Fitur 9 di atas
  **[RESOLVED — NEW]**