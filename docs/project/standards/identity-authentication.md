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

### Trust dan execution boundaries

- Google login harus menggunakan standar/protokol yang didukung dengan library terawat; backend memvalidasi authentication result dan membentuk trusted application principal. Browser-provided email/profile data bukan bukti identity atau authorization. Detail flow, callback, token/session lifetime, cookie protections, CSRF safeguards, dan credential storage diputuskan dalam approved security-sensitive solution design, bukan di-hardcode oleh kebijakan ini.
- Backend tetap menjadi authority untuk application session dan authorization; authentication provider tidak menetapkan role maupun hak terhadap organisasi. REST API dan Worker menggunakan shared internal account attribution jika relevan; Worker tidak menyamakan process identity dengan user privileges.
- **MVP dependency accepted:** sebelum metode atau recovery alternatif tersedia, kehilangan akses Google Account bisa menghalangi akses ke account Kencleng. Jangan mengklaim provider independence atau account recovery telah berfungsi hanya karena logical schema mendukung multiple identities.

