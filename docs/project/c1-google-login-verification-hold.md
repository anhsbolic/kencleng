# C1 — Hold verifikasi real Google login sampai integrasi frontend

- **Decision ID:** D-C1-GOOGLE-LOGIN-HOLD-001
- **Status:** APPROVED — hold aktif; verifikasi belum dilakukan.
- **Tanggal:** 2026-10-12
- **Work Unit:** WU-C1-ENG-001 / C1
- **Decision owner:** Anhar, Human authority
- **Verification owner:** Testing
- **Executor real Google login:** Human, secara manual; bukan agent.
- **Coordination owner:** Orchestrator

## Keputusan Human

Instruksi eksplisit Anhar: “real google login hold sampai integrasi dengan FE selesai ya, dan dilakukan manual bukan oleh agent ya.”

Verifikasi real Google login ditunda sampai integrasi frontend dengan backend C1 selesai. Sesudah integrasi, skenario real Google login dilakukan manual oleh Human. Agent tidak menjalankan atau mengotomatiskan skenario login Google nyata, memilih akun, memasukkan kredensial, consent, atau MFA.

Keputusan ini mengatur waktu dan execution mode verifikasi pada komitmen C1. Kewajiban verifikasi real provider pada Approved Techplan tetap berlaku; keputusan ini tidak menghapus requirement, mengubah trust/authentication contract, atau menyatakan suatu skenario lulus.

## Scope hold dan pekerjaan yang tetap dapat berjalan

| Area | Status dan batas |
|---|---|
| Real Google login melalui aplikasi terintegrasi | **DEFERRED — Human-approved hold sampai integrasi FE selesai.** Belum ada PASS/FAIL real-provider baru. |
| Repeat real login / stable person mapping | Bukti real-provider terkait login ikut deferred; mapping tidak boleh disimpulkan hanya dari satu login berhasil. |
| Callback atau authorization-code replay pada real-provider flow | Tidak dinyatakan lulus; verifikasi real-provider yang memerlukan flow login menunggu scope manual terintegrasi. Negative protocol/backend checks yang tidak memakai login Google nyata tidak ikut ditahan. |
| Simulated-provider/protocol negatives | Tetap dapat dilakukan dalam scope yang sah; label simulated, bukan real Google interoperability. |
| Backend HTTP/contract assertions, cookie/Origin/CSRF/logout | Tidak otomatis di-hold oleh keputusan ini. Verifikasi yang tidak memerlukan login Google nyata dapat berjalan dengan fixtures dan evidence limits yang jelas. |
| Restricted-runtime PostgreSQL/SCRAM, concurrent person upsert dan session locking/revocation/expiry | Tidak di-hold oleh keputusan ini; prerequisite DB dan authorization tetap berlaku. |
| Scanner disposition, bounded remediation, final backend gate | Tidak di-hold atau disahkan oleh keputusan ini. D-PATCH002-01 tetap pending keputusan owner. |
| Production HTTPS cookie evidence | Requirement terpisah; localhost/manual login tidak menjadi bukti production Secure-cookie behavior. |
| Integrated C1 UX/T3/T5 dan acceptance keseluruhan | Tetap milik owning scope/phase; keputusan hold login tidak menetapkan status atau menunda semua scenario secara otomatis. |

Testing backend sebelum integrasi harus mengecualikan real Google login dari claimed executed outcome dan menunjuk hold ini. Missing Google interaction pada scope yang secara eksplisit deferred tidak boleh dilaporkan sebagai implementation failure, simulated success, atau real-provider PASS. Scope backend tersebut juga tidak boleh mengklaim acceptance yang membutuhkan bukti deferred.

## Pemicu evaluasi pelepasan hold

Orchestrator meninjau hold ketika integrasi frontend dengan backend C1 telah dinyatakan selesai melalui evidence/artifact owner implementasi/integrasi. Kompilasi, mock/MSW, fixture session, atau frontend yang hanya terlihat selesai tidak cukup untuk menyimpulkan integrasi real tersedia.

Sebelum meminta Human menjalankan skenario:

1. Rekonsiliasi evidence integrasi dan tentukan exact build/revision serta scope manual yang akan diuji.
2. Periksa material prerequisites menurut [C1 auth verification environment](c1-auth-verification-environment.md): backend/runtime/schema yang diotorisasi, restricted application DB identity, browser-facing route, Google client/callback/consent/test-user/account path dan network yang diperlukan. Jangan mencatat secret values.
3. Testing menyiapkan skenario dan kriteria observable PASS/FAIL sebelum eksekusi, termasuk repeat-login bila stable person mapping menjadi claim.
4. Human menjalankan skenario secara manual dan melaporkan hanya observasi yang diminta. Agent dapat menyiapkan instruksi serta memeriksa bukti application-side yang diotorisasi secara terpisah, tanpa menjalankan real Google login.
5. Orchestrator mencatat perubahan status hold dan merekonsiliasi hasil. Run terminal lama tidak dibuka kembali; phase execution baru mengikuti fresh Run/Participant/Session dan approval yang berlaku.

Integrasi selesai membuat verifikasi ini dapat dipersiapkan; bukan automatic PASS, acceptance, dispatch, atau bukti semua prerequisites sudah terpenuhi. Bila prerequisites belum tersedia, catat gap yang spesifik dan pertahankan status belum terverifikasi.

## Evidence mode sesudah hold

Hasil interaksi real Google/manual dicatat sebagai:

- `PASS — Human-attested`, hanya untuk observasi dan skenario yang benar-benar dilaporkan;
- `FAIL — Human-attested`, dengan gejala relevan tanpa data sensitif;
- `NOT ESTABLISHED`, jika laporan tidak lengkap, ambigu atau bertentangan.

Testing ownership tetap sama walaupun executor-nya Human. Jangan mengklaim agent-independent execution. Observasi application/database yang benar-benar dilakukan agent dicatat terpisah; satu Human PASS tidak membuktikan mapping lintas sesi, replay resistance, concurrency, privilege DB, semua negative CSRF/Origin/logout, atau production HTTPS.

Jangan meminta atau menyimpan password, MFA/code, authorization code, client secret, ID/access token, cookie/session value atau identifier akun/person yang tidak diperlukan. Panduan browser dan observable criteria nanti mengikuti [environment guide](c1-auth-verification-environment.md) dan scope Testing aktif.

## Dampak readiness, acceptance dan state saat keputusan

- Real Google login berstatus **DEFERRED / belum terverifikasi**, bukan failure yang terbukti dan bukan requirement yang waived.
- B-T2-002/B-T2-003 tetap terbuka. Hold tidak menutup blocker atau menerima dependencies/scanner candidates.
- Tidak ada T2 acceptance, T3 readiness/dispatch atau whole-C1 completion dari keputusan ini. Kondisi dependency/acceptance yang sudah disepakati tetap membutuhkan owning reconciliation; hold sendiri tidak memenuhi prasyarat tersebut.
- Approved Techplan, task snapshots dan upstream Product authority tidak ditulis ulang. Bila downstream acceptance/routing membutuhkan perubahan material owning contract, route ke owner; jangan menggunakan hold untuk diam-diam melonggarkan kontrak.
- Keputusan ini bukan persetujuan opsi scanner yang sebelumnya diajukan dan bukan authorization menjalankan Run baru.

## Referensi owning authority dan evidence

- [Approved Techplan §§9–12 dan R1/R9/R10/R14](../../.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md)
- [T2 task](../../.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/T2-backend-identity-session.md)
- [C1 auth verification environment](c1-auth-verification-environment.md)
- [Work Unit](../../.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md) dan [Events](../../.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/events.md)

Historical test/review evidence tetap immutable. Dokumen ini adalah owning bounded Human scheduling/execution-mode decision; bukan laporan hasil Testing.

## Dokumen pendamping

[Hold verifikasi C1 terintegrasi](c1-integrated-verification-hold.md) mencatat full journey, rendered UI/accessibility, cache isolation dan Human semantic walkthrough. Batas manual-only real Google pada dokumen ini tidak otomatis berlaku untuk seluruh browser/agent verification pada dokumen pendamping.
