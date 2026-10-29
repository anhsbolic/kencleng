# ORG-01 — Pendekatan privasi dan lifecycle data (PD-PRIV-01)

**Standing: PRINSIP, ALUR KONTEKS SEDERHANA, RETENSI DAN PEMETAAN KANDIDAT DASAR/ARAH PEMBERITAHUAN DISETUJUI HUMAN / VALIDASI DASAR DAN PEMBERITAHUAN FINAL MASIH TERBUKA.** Pemilik keputusan: Anhar, Product/privacy owner. Revisi pencatatan: `PD-PRIV-01-r4`, 29 Oktober 2026. Berlaku untuk ORG-01 pada branch `pilot/station5-c1-space-only`.

## Keputusan dan bukti persetujuan

Anhar menyetujui rekomendasi **menyimpan data minimum berdasarkan tujuan yang jelas, dengan aturan kapan penyimpanannya berakhir**, melalui pernyataan: **“ya ak sepakat dengan rekomendasinya. gunakan pendekatan ini ya”**. Persetujuan diberikan setelah penjelasan data minimum, retensi berdasarkan tujuan, permintaan penghapusan, retensi teknis/backup, serta informasi kepada pengguna dan dasar pemrosesan.

Persetujuan lanjutan: Anhar menyatakan **“ya sepakat bro.”** setelah menerima rekomendasi alur penghapusan akun Owner: verifikasi pemohon, penjelasan dampak, penutupan eksplisit untuk konteks ORG-01 sederhana tanpa pengalihan Owner otomatis, pencabutan akses dan penanganan data, serta pemberitahuan status yang sebenarnya. Alur dan batas kasus yang disetujui dicatat di bawah.

Persetujuan retensi: Anhar menyatakan **“ya sepakat. lanjut”** setelah menerima usulan receipt 48 jam, riwayat session/refresh sampai masa family 30 hari ditambah maksimal 24 jam pembersihan, log yang diperlukan 14 hari dan umur backup maksimal 14 hari, beserta syarat safe cleanup dan pemulihan yang menjaga keputusan penghapusan. Durasi berikut adalah batas operasional awal yang disetujui, belum berdasarkan pengukuran penggunaan Kencleng dan bukan pernyataan angka wajib hukum.

Persetujuan pemetaan dan pemberitahuan: Anhar menyatakan **“sepakat bro. lanjutkan”** setelah rekomendasi kandidat dasar per tujuan, validasi terhadap syarat layanan/operasi, ringkasan sebelum Google login dan pemberitahuan lengkap. Persetujuan ini mencakup arah penyusunan; tidak membuktikan bahwa syarat dasar hukum telah terpenuhi atau teks final telah disetujui.

Informasi operasional dari Human: **“1. Yayasan Kencleng Indonesia; 2. Ya Belum ditentukan. Baru Google Login aja yang sudah pasti”**. Pengendali yang ditetapkan adalah Yayasan Kencleng Indonesia; selain Google Login, penyedia hosting/database/backup belum ditentukan. Pada jawaban lanjutan Human memilih kanal email **care@kencleng.dev** untuk permintaan privasi; kesiapan penerimaan/penanganan email belum diuji pada pencatatan ini.

Dokumen ini mencatat keputusan dan informasi tersebut. Persetujuan ini tidak menetapkan workflow pengalihan Owner, hasil validasi dasar hukum, atau kesiapan pemrosesan data orang nyata. [Product Contract](product-contract.md) dan [approval Product](approval.md) tetap menjadi sumber perilaku pembentukan organisasi dan Owner; SC/REQ/AC-nya tidak diubah oleh pencatatan kebijakan lifecycle ini.

## Prinsip yang digunakan

| ID | Prinsip disetujui | Batas keputusan |
| --- | --- | --- |
| PRIV-01 | **Data minimum untuk tujuan ORG-01.** Batasi data pada identitas akun internal, pengenal Google untuk login, nama/ID organisasi, hubungan Owner, waktu pembentukan, serta catatan teknis yang diperlukan untuk keamanan dan pengiriman ulang. | Email, foto, alamat dan dokumen tambahan belum diperlukan ORG-01. Pengenal teknis dan nama organisasi yang dapat terkait orang tetap masuk penilaian privasi. Detail mekanisme mengikuti Engineering Standards dan scoped design yang berlaku. |
| PRIV-02 | **Retensi data utama berdasarkan tujuan.** Simpan selama masih diperlukan untuk menyediakan akun, konteks organisasi dan pemeriksaan hubungan Owner. Tetapkan pemicu penutupan/penghapusan secara eksplisit. | Tidak memilih angka tahun tanpa kebutuhan yang mendasarinya. Permintaan penghapusan konteks sederhana mengikuti alur di bawah; batas retensi, pengecualian dan kasus lainnya masih perlu dirinci. Prinsip ini tidak menjadi izin penyimpanan tanpa batas atau penolakan hak pengguna. |
| PRIV-03 | **Permintaan penghapusan melalui proses yang memverifikasi pemohon.** Perhitungkan organisasi yang masih bergantung pada Owner tersebut; untuk konteks ORG-01 sederhana gunakan alur penutupan eksplisit di bawah. | Tidak ada penghapusan organisasi atau pengalihan Owner otomatis. Pengalihan, serta kasus yang melibatkan pengguna lain/aktivitas/kewajiban tambahan, memerlukan keputusan terkait sebelum diterapkan. |
| PRIV-04 | **Retensi catatan teknis dan backup dibatasi tersendiri.** Catatan pengiriman ulang mendukung pencegahan efek ganda dan pemeriksaan hasil; catatan keamanan mengikuti kebutuhan deteksi penyalahgunaan. | Durasi dan syarat pada tabel retensi di bawah telah disetujui. Jangka waktu retry/session tetap berbeda dari durasi penyimpanan catatan. Backup dan pemulihan wajib menjaga keputusan penghapusan tetap berlaku. |
| PRIV-05 | **Informasi pemrosesan dapat dipahami pengguna.** Jelaskan data yang dipakai, tujuan, akses, retensi dan cara mengajukan permintaan. Gunakan pemetaan kandidat per tujuan dan penempatan pemberitahuan di bawah sebagai arah yang disetujui. | Login Google tidak menetapkan seluruh kebijakan tersebut. Validasi dasar, teks pemberitahuan final dan kebutuhan consent spesifik bila dipakai masih terbuka. Dokumen ini tidak menyatakan kepatuhan hukum telah diverifikasi. |

Tujuan atribusi dan pemeriksaan kembali berasal dari Product Contract approved. [T04 — Privacy, PII & Cryptography](../../../engineering/standards/privacy-data-protection.md) tetap mengatur inventory, safeguards dan pemisahan keputusan owner dari mekanisme Engineering. Pendekatan ini tidak menjadi kebijakan umum seluruh fitur Kencleng atau perubahan Engineering Standards.

## Alur penghapusan akun Owner yang disetujui

**Rekomendasi yang diterima:** untuk konteks yang baru mencakup ORG-01, proses penghapusan akun disertai pilihan penutupan konteks organisasi secara eksplisit. Penutupan ini hanya berlaku pada konteks di Kencleng; tidak membubarkan organisasi secara hukum.

1. **Verifikasi pemohon.** Pastikan permintaan berasal dari pemilik akun dengan bukti kontrol akun yang memadai. Pengumpulan dokumen identitas tambahan harus memiliki kebutuhan yang jelas; tidak otomatis diperlukan.
2. **Jelaskan dampak sebelum tindakan.** Tampilkan seluruh organisasi yang bergantung pada akun, akses yang akan berakhir, data yang akan dihapus, dan data yang masih perlu disimpan beserta alasannya. Detail retensi yang belum diputuskan tidak boleh disajikan sebagai kebijakan final.
3. **Selesaikan hubungan Owner–organisasi sesuai kasus berikut.**

| Keadaan | Perlakuan Product yang disetujui |
| --- | --- |
| Akun tidak memiliki organisasi | Lanjutkan penghapusan akun sesuai kebijakan data. |
| Akun merupakan satu-satunya Owner dan organisasi hanya memiliki konteks ORG-01 | Minta konfirmasi penutupan konteks organisasi sebagai bagian proses penghapusan. Pastikan organisasi tidak tertinggal aktif tanpa Owner. Keberadaan akun sebagai Owner tidak dengan sendirinya memberi izin penutupan tanpa konfirmasi. |
| Pengguna ingin organisasi tetap berjalan | Perlukan keputusan pengalihan Owner; fitur tersebut belum dirancang atau disetujui dalam ORG-01. Tidak memilih pengganti atau memberi kewenangan secara otomatis. |
| Organisasi melibatkan pengguna lain, Campaign, transaksi atau kewajiban tambahan | Nilai dampak memakai aturan fitur terkait sebelum menerapkan penutupan. Permintaan privasi tetap ditangani dan tidak boleh dibiarkan menggantung tanpa batas. Tidak menerapkan alur konteks sederhana secara otomatis pada kasus ini. |

4. **Cabut akses dan tangani data.** Akhiri session serta kewenangan akun, hapus data pribadi yang tidak lagi diperlukan, dan tangani catatan teknis/backup sesuai jadwal retensi yang ditetapkan. Mengganti nama tampilan menjadi “akun dihapus” belum membuktikan data anonim. Ini kebijakan hasil yang diinginkan; mekanisme, urutan konsistensi dan verification ditentukan dalam desain Engineering yang diotorisasi.
5. **Sampaikan hasil sebenarnya.** Bedakan permintaan diterima, akses dihentikan dan penghapusan selesai. Pemberitahuan tidak boleh menyatakan seluruh salinan telah terhapus bila penanganan backup belum selesai atau belum dibuktikan.

Alur ini tidak menetapkan penghapusan SQL/cascade tertentu, perpindahan hak, masa tunggu pengajuan numerik, atau pengecualian retensi baru. Penghapusan harus tetap memperhitungkan aturan pemrosesan yang berlaku. Evaluasi dasar hukum dan rincian pemberitahuan masih terbuka; pencatatan persetujuan Product tidak membuktikan kepatuhan hukum atau keberhasilan implementasi.

## Retensi catatan teknis dan backup yang disetujui

Data utama akun dan organisasi tetap mengikuti tujuan layanan serta alur penghapusan di atas. Batas berikut berlaku untuk kategori teknis ORG-01 yang benar-benar diperlukan; tidak mengotorisasi pengumpulan log atau fields baru.

| Kategori | Batas operasional disetujui | Tujuan dan syarat |
| --- | --- | --- |
| Receipt pembentukan organisasi yang berhasil | **48 jam sejak pembentukan berhasil** | Mendukung pemeriksaan hasil/retry. Rancangan memiliki jendela retry 24 jam; satu hari tambahan memberi ruang pemeriksaan dan pembersihan. Sesudah receipt dihapus, organisasi tetap dapat diperiksa melalui own-list/detail yang berwenang. Pembersihan tidak boleh memungkinkan request lama membentuk organisasi lagi; jendela retry serta transaksi in-flight harus ditangani dengan aman. |
| Session dan riwayat refresh token | Selama diperlukan untuk validasi, pencabutan dan deteksi replay; batas akhir **30 hari masa family + maksimal 24 jam pembersihan** | Mengikuti rancangan session. Session yang berakhir lebih awal dibersihkan lebih cepat setelah kebutuhan keamanan selesai. Durasi penyimpanan tidak memperpanjang validity token/session. Token mentah tidak disimpan di log. |
| Log keamanan/diagnostik yang memang diperlukan | **14 hari sejak kejadian** | Jendela awal investigasi dua minggu. Isi dibatasi untuk diagnosis, tanpa nama organisasi, profil, token atau payload pribadi lengkap. Tidak mewajibkan audit platform/log tambahan. |
| Backup | **Maksimal umur 14 hari**, termasuk salinan dan rangkaian yang diperlukan untuk pemulihan | Jendela pemulihan dua minggu dan batas keberadaan salinan data. Kebutuhan lebih panjang memerlukan alasan serta keputusan tersendiri; bukan otomatis memperpanjang seluruh retensi. Ini batas umur backup, bukan frekuensi backup atau bukti recoverability. |

Syarat penghapusan dan backup yang ikut disetujui:

1. Keputusan penghapusan pada sistem aktif langsung berlaku. Salinan dalam backup dibatasi hanya untuk pemulihan sampai kedaluwarsa; tidak boleh dipakai sebagai sumber operasi biasa untuk data yang telah dihapus.
2. Sebelum layanan hasil restore dibuka, terapkan kembali keputusan penghapusan dan pencabutan akses. Data serta akses yang sudah dicabut tidak boleh hidup kembali. Penanda penghapusan minimum perlu dirancang selama masih ada backup terdampak; fields, perlindungan dan cleanup-nya harus mengikuti kebutuhan tersebut, tanpa menyimpan profil sebagai pengganti penghapusan.
3. Pemberitahuan harus menjelaskan keadaan nyata sistem aktif dan sisa salinan backup. Jadwal penghapusan salinan yang dijanjikan harus didukung mekanisme serta bukti, bukan hanya label status.

Durasi ini harus diterjemahkan ke bounded cleanup/in-flight/restore behavior dan diverifikasi sebelum diklaim terpenuhi. Bila mekanisme tidak dapat sekaligus menjaga safe retry/replay dan memenuhi batas, owning producer harus mengangkat konflik untuk diselesaikan; jangan diam-diam memperpanjang retensi atau menghapus safeguards. Jadwal ini tidak menetapkan pengecualian atas hak penghapusan atau kewajiban yang berlaku. Dasar pemrosesan dan penilaian penerapannya masih perlu ditetapkan sebelum data orang nyata digunakan.

## Pemetaan dasar dan arah pemberitahuan yang disetujui

**Approval terbatas pada pemetaan kandidat untuk divalidasi.** Penilaian harus mempertimbangkan syarat layanan, hubungan dengan pengguna dan cara operasi aktual; kesepakatan Product tidak menggantikan terpenuhinya dasar hukum.

| Tujuan | Kandidat dasar yang disetujui untuk divalidasi | Syarat yang tetap terbuka |
| --- | --- | --- |
| Menyediakan akun, membentuk organisasi dan mencatat hubungan Owner | Pemenuhan perjanjian atau permintaan pengguna | Pastikan hubungan layanan/perjanjian serta kebutuhan data memenuhi ketentuan yang berlaku. Belum ada hasil validasi syarat layanan yang dicatat. |
| Keamanan dan deteksi penyalahgunaan | Kepentingan yang sah | Dokumentasikan tujuan, kebutuhan dan keseimbangan terhadap hak pengguna; pemetaan ini tidak mengotorisasi logging/profiling tambahan. |
| Backup dan pemulihan | Mengikuti dasar dan tujuan data asal | Terapkan batas retensi dan pembatasan penggunaan yang disetujui; tidak menjadi izin penggunaan data untuk tujuan baru. |

Jika suatu tujuan memakai persetujuan pengguna, jelaskan tujuan spesifik dan rekam persetujuan dengan jelas. Ini conditional rule, bukan pemilihan consent untuk semua pemrosesan atau penambahan tujuan/fitur opsional.

Arah pemberitahuan yang disetujui:

- Tampilkan ringkasan sebelum Google login, karena pemrosesan identitas dimulai pada tahap itu, dengan akses ke pemberitahuan lengkap.
- Pemberitahuan lengkap mencakup identitas pengendali/kontak; data, tujuan dan dasar per pemrosesan; pihak penerima/pembantu pemrosesan; retensi/penghapusan/backup; serta hak dan kanal permintaan pengguna.
- Sebelum pembentukan organisasi, tampilkan penjelasan tentang hubungan Owner dan akibat tindakannya.

**Pengendali: Yayasan Kencleng Indonesia. Kanal kontak: care@kencleng.dev**, keduanya ditetapkan Human. **Penyedia: Google Login dipilih; hosting, database, backup dan penyedia lain BELUM DITENTUKAN.** Konfigurasi developer lokal atau nama tool tidak dijadikan bukti vendor/region pemrosesan data orang nyata. Kanal email ditetapkan sebagai keputusan, bukan bukti inbox atau prosedur penanganan permintaan telah diuji.

[Draf pemberitahuan privasi ORG-01](privacy-notice-draft.md) disiapkan sebagai bahan review dan pelengkapan. Teks draf tidak otomatis approved, published atau bukti mekanisme telah berjalan. Referensi hukum pendukung pemetaan adalah [UU PDP, terutama Pasal 20–22](https://jdih.komdigi.go.id/produk_hukum/view/id/832/t/undangundang%20nomor%2027%20tahun%202022); pembacaan sumber bukan validasi penerapan pada operasi Kencleng.

## Rincian yang harus diputuskan berikutnya

1. **Kasus lifecycle di luar konteks sederhana:** pengalihan Owner bila organisasi ingin tetap berjalan, serta dampak pada pengguna lain/aktivitas/kewajiban tambahan. Penutupan eksplisit konteks ORG-01 sederhana telah dipilih; mekanisme verifikasi, konsistensi perubahan dan penghapusan datanya belum diimplementasikan atau diuji.
2. **Rincian penerapan retensi:** pemicu akhir kebutuhan data utama di luar alur sederhana, penilaian pengecualian/copies bila memang ada dasar yang berlaku, dan desain cleanup/penanda penghapusan/restore. Durasi teknis/backup pada tabel telah diputuskan; implementasi dan buktinya belum tersedia.
3. **Validasi dasar dan pemberitahuan final:** hasil penilaian perjanjian/kebutuhan data serta kepentingan yang sah, kesiapan kanal care@kencleng.dev, daftar penyedia/penerima dan lokasi pemrosesan aktual, serta teks/tanggal berlaku yang disetujui sebelum penggunaan data orang nyata. Nama pengendali dan alamat kontak telah diberikan; tidak perlu meminta keputusan identitas/kontak yang sama kembali.

Prinsip, alur penghapusan konteks sederhana, durasi teknis/backup serta pemetaan kandidat dasar/arah pemberitahuan telah disepakati. Langkah berikutnya adalah melengkapi informasi operasional dan menilai draf/validasi yang masih terbuka. Kasus pengalihan/aktivitas tambahan memerlukan keputusan terkait ketika scope tersebut dipilih.

## Batas penggunaan dan kaitan Rancang

**PD-PRIV-01 TERSELESAIKAN SEBAGIAN PADA KEPUTUSAN DAN ARAH PEMETAAN; GATE DATA ORANG NYATA TETAP DITAHAN.** Development dan pengujian dapat direncanakan dengan data sintetis. Pernyataan ini tidak mengotorisasi implementasi, menjalankan layanan/tes, canonical API adoption, publikasi pemberitahuan atau penggunaan data orang nyata.

Catatan dampak pada assignment ada pada [tindak lanjut Rancang](../../../../harscode-state/org-01/rancang-002/privacy-approach-follow-up.md). Bundle `ORG-01-rancang-002-r1`, Security Findings dan Solution Review mempertahankan revisi yang benar-benar dinilai; persetujuan pendekatan ini bukan acceptance bundle atau penutupan SEC-02. Sebelum rincian lifecycle dikonsumsi sebagai kontrak implementasi, revalidasi bagian terkait pada data/receipt/session design, TECH-06/09, Security dan Solution Review terhadap exact owning revision yang berlaku.
