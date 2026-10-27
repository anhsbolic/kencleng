# Kencleng — Identity & Authentication Standards

> **Status:** Project-wide Engineering policy T02, disetujui Anhar pada research branch `pilot/station5-c1-space-only`.
> **Scope:** Authentication methods, internal account identity, dan extensibility lintas frontend, backend, persistence, dan security. Bukan implementation spec, Product login requirement, atau authorization matrix.
> **Authority:** [Engineering Standards index](README.md). Authentication policy tidak dimiliki oleh backend tech stack.

## Identity & Authentication policy (T02)

> **Keputusan Engineering project-wide — berlaku untuk desain baru.** Kebijakan ini tidak menyatakan authentication, account schema, atau session telah diimplementasikan. Pilihan product-scoped tentang kapan login diwajibkan tetap mengikuti Product Authority.

### Metode login dan scope MVP

- **MVP: Google Sign-In saja.** Google dapat digunakan langsung sebagai trusted identity provider melalui integrasi standar; tidak wajib menambahkan broker/managed-auth vendor lain pada baseline. Pengguna dengan email non-Gmail tetap bisa login bila email tersebut memiliki Google Account, tetapi **email non-Gmail tanpa Google Account belum didukung pada MVP**.
- **Tidak membangun credential authentication sendiri untuk MVP.** Penyimpanan password Kencleng, email/password + OTP, email OTP, dan nomor telepon + OTP bukan scope MVP. Metode-metode ini boleh dievaluasi dan ditambahkan kemudian melalui keputusan tersendiri.
- Guest/public access tidak otomatis memerlukan login; kebutuhan authenticated actor ditentukan oleh effective Product requirements dan operation security.

### Identitas internal dan extensibility

- **Kencleng owns the account identity.** Gunakan stable internal `account_id` sebagai referensi business ownership, attribution, dan authorization. Jangan gunakan email, nomor telepon, atau Google `sub` sebagai primary Kencleng account identity.
- **Pisahkan account dari authentication identities pada logical data model.** Desain persistence pertama harus mendukung satu account memiliki beberapa external authentication identities (konseptual `accounts` → `authentication_identities` one-to-many), dengan mapping identifier provider/issuer dan subject yang unik serta tervalidasi. Google menjadi satu-satunya login method aktif MVP; kemampuan menambah identity di masa depan tidak berarti linking workflow sudah tersedia.
- External login identity harus diidentifikasi melalui provider-scoped, verified immutable/stable subject (misalnya OIDC `issuer` + `sub`), bukan email atau display name. Account internal tidak boleh berubah hanya karena login method baru ditambahkan.
- **No automatic identity merging/linking by matching email or phone.** Mengaitkan metode baru ke account existing memerlukan pembuktian kontrol terhadap account lama dan credential/metode baru, proteksi account takeover, serta explicit security-sensitive workflow. Account recovery, unlinking, provider migration, dan metode baru adalah future scoped decisions, bukan otomatis tersedia.
- Arsitektur harus mempertahankan opsi menambahkan email/password + OTP atau phone + OTP tanpa mengganti account IDs dan referensi business ownership yang telah ada. Tidak perlu membangun provider abstraction framework, credential tables, atau SMS/email gateway sebelum use case membutuhkannya.

### Browser session security — JWT + rotating refresh token (T02 amendment)

- **MVP wajib menggunakan Kencleng-issued short-lived access JWT** sebagai application authentication token di `HttpOnly` cookie; gunakan separate **opaque, cryptographically random, single-use rotating refresh token** di `HttpOnly` cookie. Token Google/IdP berbeda dari token aplikasi Kencleng. Tidak perlu membangun OAuth authorization server umum atau token provider abstraction untuk mewujudkan kebijakan ini.
- Kedua authentication cookies pada production menggunakan `Secure`, `HttpOnly`, appropriate `SameSite=Lax` atau `Strict`, serta scoping `Path` dan host-only tanpa `Domain` seluas mungkin. HTTPS adalah kewajiban production. Jika perlu HTTP loopback lokal, konfigurasi pengecualian harus eksplisit, terbatas loopback, dan mustahil menjadi production default. Hindari collision path/cookie tanpa mengurangi protections.
- Frontend **tidak menyimpan access JWT, refresh token, session secrets, atau provider tokens** di `localStorage`, `sessionStorage`, accessible JavaScript state, maupun URL. Tidak mengeluarkannya ke response body yang dibaca frontend untuk disimpan. Frontend mengetahui status login melalui backend session/current-user API yang berwenang.
- Backend memverifikasi JWT dengan trusted signing keys dan allowlisted algorithm, signature, issuer, audience, subject (`account_id` internal), waktu berlaku, dan claims relevan. Jangan menerima `alg=none`, client-provided role/permissions, atau menempatkan provider access/ID token sebagai trusted Kencleng authorization. JWT adalah bukti authentication yang terikat scope/audience, **bukan sumber authority final** untuk role, relationships, atau current resource state.
- **Refresh rotation:** refresh token dibatasi lifetime (idle dan/atau absolute), disimpan dalam **hashed/verifier form** yang aman di server-side persistent state, dan hanya dapat digunakan satu kali. Refresh harus melakukan invalidation old token + issuance new token secara atomic/concurrency-safe. Server mendeteksi reuse/replay dan mencabut token family/session yang terdampak; percobaan refresh concurrent tidak boleh menghasilkan dua token aktif dari satu predecessor. Penanganan retry/race harus fail-safe dan diuji, tanpa membuka replay window yang tidak terkontrol.
- **Session termination dan revocation:** logout dan security-relevant events harus menghentikan refresh/session authority; existing access JWT mungkin masih valid sampai expiry sehingga expiry harus singkat. Operasi berisiko tinggi/revocation-sensitive harus memakai trusted effective state atau check tambahan bila memerlukan revocation yang lebih segera. Jangan menjanjikan instant revocation access JWT bila mekanisme demikian belum dibuktikan.
- **CSRF:** karena browser otomatis mengirim cookies, semua state-changing operations termasuk refresh dan logout harus memiliki CSRF defense server-side (misalnya strict Origin check dengan session-bound CSRF token atau equivalent robust pattern). `SameSite` saja dan CORS saja tidak selalu cukup. Jangan membuat state-changing `GET` endpoint; batasi accepted origins dan cookie exposure.
- **Security operations:** refresh, logout, issuance, token validation, dan session state harus memiliki negative/replay/concurrency/expiry/revocation tests serta safe audit diagnostics; jangan log raw tokens, keys, atau unnecessary PII. Detail cryptographic algorithm/key rotation, JWT library, cookie names, dan numeric TTL merupakan implementation decisions yang wajib security review, bukan standar numerik prematur.

### Trust dan execution boundaries

- Google login harus menggunakan standar/protokol yang didukung dengan library terawat; backend memvalidasi authentication result dan membentuk trusted application principal. Browser-provided email/profile data bukan bukti identity atau authorization. Kencleng menerbitkan **application access JWT** setelah memvalidasi hasil Google Sign-In; Google ID token tidak otomatis menjadi token API Kencleng. Detail callback, library, exact expiry, dan storage implementation ditentukan dalam approved security-sensitive solution design, dengan wajib mengikuti baseline session security di bawah.
- Backend tetap menjadi authority untuk application session dan authorization; authentication provider tidak menetapkan role maupun hak terhadap organisasi. REST API dan Worker menggunakan shared internal account attribution jika relevan; Worker tidak menyamakan process identity dengan user privileges.
- **MVP dependency accepted:** sebelum metode atau recovery alternatif tersedia, kehilangan akses Google Account bisa menghalangi akses ke account Kencleng. Jangan mengklaim provider independence atau account recovery telah berfungsi hanya karena logical schema mendukung multiple identities.

