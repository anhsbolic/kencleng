> Catatan packaging: ini project-owned checkpoint keputusan/follow-up 29 Oktober2026, dipertahankan sebagai rekaman owner. Pernyataan belum direview/belum diintegrasikan merujuk saat pencatatan tersebut; standing Rancang sekarang mengikuti [handoff](handoff.md) dan [state](state.md), dengan exact original r4 sudah diterima Human. Keputusan privasi dan remaining legal/provider/request/restore gates tidak diubah.

# ORG-01 / rancang-002 — Tindak lanjut persetujuan pendekatan privasi

**Status: pencatatan keputusan owner dan draf pemberitahuan selesai; prinsip, alur konteks sederhana, retensi serta pemetaan kandidat dasar/arah pemberitahuan approved / validasi dan teks final OPEN; real-person data edge HELD.** Catatan Koordinator, 29 Oktober 2026. Ini checkpoint tambahan sesudah bundle `ORG-01-rancang-002-r1`, bukan revisi producer atau independent review baru.

## Keputusan Human dan sumber efektif

Human menyatakan **“ya ak sepakat dengan rekomendasinya. gunakan pendekatan ini ya”** sesudah penjelasan lima bagian pendekatan privasi. Otorisasi ini mendasari pencatatan pendekatan pada owning Product source dan navigasinya, serta catatan dampak assignment ini. Tidak ada izin code/API/schema/standards/service/deploy/commit/push atau perubahan hasil review.

Human kemudian menyatakan **“ya sepakat bro.”** sesudah rekomendasi penghapusan akun Owner dengan verifikasi pemohon, penjelasan dampak, penutupan eksplisit konteks ORG-01 sederhana tanpa pengalihan otomatis, pencabutan akses/penanganan data dan pemberitahuan yang jujur. Persetujuan lanjutan dicatat pada revisi Product `PD-PRIV-01-r2`.

Human selanjutnya menyatakan **“ya sepakat. lanjut”** sesudah penjelasan durasi receipt 48 jam, session/refresh sampai masa family 30 hari + maksimal 24 jam pembersihan, log yang diperlukan 14 hari dan umur backup maksimal 14 hari, berikut syarat safe cleanup dan restore yang menjaga penghapusan/pencabutan akses. Persetujuan ini dicatat pada `PD-PRIV-01-r3`.

Human selanjutnya menyatakan **“sepakat bro. lanjutkan”** sesudah kandidat pemenuhan perjanjian/permintaan pengguna untuk layanan inti, kepentingan yang sah untuk keamanan, backup mengikuti dasar/tujuan asal, validasi sesuai syarat layanan/operasi dan arah pemberitahuan sebelum login dijelaskan. Persetujuan ini direkam sebagai arah pemetaan, bukan hasil validasi hukum atau acceptance teks final.

Human menetapkan **Yayasan Kencleng Indonesia** sebagai pengendali, **care@kencleng.dev** sebagai kanal privasi dan Google Login sebagai penyedia yang sudah dipilih. Hosting/database/backup lain dinyatakan belum ditentukan. Tidak ada identitas/kontak/vendor tambahan yang diinventasikan, dan tidak ada email yang dikirim atau inbox yang diuji.

Owning decision: [PD-PRIV-01-r4](../../../docs/product/outcomes/org-01/privacy-decision.md), melalui [Product routing](../../../docs/product/README.md). Exact revision pencatatan dan hashes tercantum pada [privacy-approach-revisions.json](privacy-approach-revisions.json). [Draf pemberitahuan v0.1](../../../docs/product/outcomes/org-01/privacy-notice-draft.md) mencakup ringkasan sebelum login dan penjelasan lengkap; belum approved/published. Validasi perjanjian/necessity, penilaian kepentingan yang sah, vendor/lokasi/recipient details, prosedur hak pengguna serta teks/tanggal berlaku final masih OPEN.

## Perubahan baseline dan standing bundle sebelumnya

- Saat pencatatan prinsip awal, repository ditemukan bersih pada `35f8c0fc31f1f2f21bdac36d4c24d521541263f0`, branch `pilot/station5-c1-space-only`. Selisih dari baseline desain `d0d78fab26a976f9f48c6fc97a45c0adc5cec282` hanya menambahkan rancang-002-r1 dan tautan Engineering. Pada pencatatan r2, HEAD tetap sama; hanya empat file pencatatan privasi dari tindak lanjut sebelumnya yang modified/untracked. Tidak ada source drift pada Product Contract/approval, Design atau Engineering Standards yang dipakai desain.
- [Manifest r1](https://github.com/anhsbolic/kencleng/blob/6ba8ebcf1a41373ba9ddea441a3659a591c07cb5/harscode-state/org-01/rancang-002/output-revisions.json) dan seluruh sepuluh hash payload/navigation-nya diverifikasi kembali. Bundle r1 tetap tersedia sebagai exact checkpoint yang semula dinilai; file r1, state/handoff, findings dan manifest tidak ditulis ulang untuk memberi kesan review mencakup keputusan baru ini.
- Independent Security dan Solution Review sebelumnya menilai r1. **Keputusan Product baru ini belum mendapat affected-edge reassessment independen.** Tidak ada klaim review atau runtime PASS baru.

## Dampak dan bagian yang tetap ditahan

Prinsip dan alur penghapusan akun Owner untuk konteks sederhana kini dicatat sebagai keputusan Human Product/privacy scoped. Penutupan organisasi memerlukan konfirmasi eksplisit; tidak ada pengalihan otomatis atau organisasi aktif yang ditinggalkan tanpa Owner. Keinginan mempertahankan organisasi serta kasus pengguna/aktivitas/kewajiban tambahan membutuhkan keputusan terkait; permintaan privasi tidak boleh diabaikan tanpa batas. Tidak ada durasi atau mekanisme teknis penghapusan yang ditambahkan tanpa persetujuan.

Durasi teknis/backup dan pemetaan kandidat kini diputuskan pada owning source. Rincian untuk menutup SEC-02 belum lengkap: validasi dasar, rincian penyedia/lokasi, readiness kontak/prosedur hak pengguna dan final notice, serta policy terkait untuk kasus di luar konteks sederhana masih OPEN. Keputusan baru belum diintegrasikan ke proposed SQL/FK, session/revocation ordering, receipt cleanup, deletion markers, backup/restore atau kontrak observasi r1. Lifecycle tersebut tidak boleh diklaim sudah didukung hanya karena pencatatan Product selesai. Jadwal retensi tidak memperpanjang validity token, memberi pengecualian penghapusan otomatis atau mengizinkan raw-token/PII logging.

Real-person collection/persistence release, lifecycle/deletion/cleanup readiness dan klaim privacy readiness tetap HELD. Otorisasi development/tes data sintetis maupun implementasi protected core tetap terpisah. Acceptance bundle r1 tetap PENDING; persetujuan pendekatan privasi tidak mengubahnya.

Next action: Human meninjau draf pemberitahuan dan melengkapi penyedia/lokasi saat dipilih; pemilik Product/privacy menilai kesesuaian kandidat dasar terhadap syarat layanan/operasi serta prosedur penanganan permintaan. Sebelum klaim gate teratasi, buka ulang affected data/receipt/session/backup/notice design dan applicable TECH obligations melalui owning producer, Security serta Reviewer dengan scoped authorization yang sesuai. Consumer/resume harus membuka Product authority terkini dan checkpoint ini bersama state/handoff r1. Pemilihan vendor baru bukan prasyarat untuk membaca draf atau melanjutkan assessment dokumentasi dengan data sintetis sesuai otorisasi terpisah.

## Evidence pencatatan

Pemeriksaan pencatatan meliputi kesetiaan keputusan terhadap empat persetujuan Human serta informasi pengendali/kontak/vendor, tautan lokal, perubahan file scoped, preservation hash bundle r1 dan exact hashes catatan/draf terbaru. Sumber resmi UU PDP dan dokumentasi Google OIDC dibaca sebagai referensi, bukan bukti legal/runtime compliance. Hasil/daftar revisi tercantum pada privacy-approach-revisions.json; revisi r4 menggantikan pencatatan privasi r3, tanpa mengganti bundle solusi r1. Hash recording sebelumnya disimpan sebagai lineage; tidak diklaim bahwa seluruh intermediate bytes pencatatan privasi sudah diarsipkan Git. Runtime tests, email-delivery test, legal validation dan affected independent reassessment **NOT RUN** pada tindak lanjut ini.
