# Kencleng — Journey Level 2 (Draf Bahasa Indonesia)

> **Status:** DRAF PARSIAL — mencakup cara membaca peta, `J-ORG` (O1–O5), dan `J-CAM` (P1–P8). **Belum menjadi sumber utama atau versi Level 2 lengkap.**
> **Sumber resmi:** [Level 2 v0.3](level-2-network.md), blob `4970ba936d107e50d8260c9e246816ba13d41a28`, disetujui Anhar pada 22 Oktober 2026. [Product Intent](../product-intent.md) dan [pengantar Journey](README.md) tetap berlaku menurut lingkupnya.
> **Batas:** Penyederhanaan bahasa saja. Tidak membuat aturan O1 baru, mengganti 26 kode node, mengubah hubungan antar-Journey, atau memberi izin development. Bagian yang belum ditulis ulang **wajib dibaca dari dokumen resmi**.
> **Lokasi:** Branch riset `pilot/station5-c1-space-only`; tidak mengubah Kencleng `main` atau Harscode.

## 1. Cara membaca peta

Peta Level 2 menunjukkan **perubahan yang bermakna** bagi pengguna atau keadaan Product: siapa yang bertindak, apa yang benar-benar berubah, syaratnya, dan apa yang masih belum pasti. Ini bukan daftar layar, API, fitur CRUD, urutan development, atau bukti pekerjaan sudah selesai.

### Kode kelompok Journey

| Kode | Kelompok | Kode historis |
| --- | --- | --- |
| `J-ORG` | Organisasi dan peninjauan | `R1` |
| `J-CAM` | Campaign | `R2` |
| `J-VIS` | Pengunjung dan Donatur | `R3` |
| `J-ACC` | Pertanggungjawaban | `R4` |
| `EX-OPS` | Pengecualian pengelola platform | `X1` |

`EX-OPS` berlaku lintas-Journey dan **bukan Journey kelima**.

### Arti hubungan

- **Syarat yang diperlukan:** Tanpa kondisi, informasi, atau kewenangan tertentu, hasil tidak boleh dinyatakan benar. Memenuhi satu syarat **belum tentu cukup** untuk berhasil.
- **Bersyarat:** Hubungan hanya berlaku jika kondisi terkait terpenuhi. Belum ada aturan umum **tidak berarti** sebuah kondisi boleh diabaikan untuk semua kasus.
- **Kelanjutan:** Suatu hasil memunculkan tanggung jawab atau kebutuhan berikutnya; tidak otomatis menentukan pekerjaan development selanjutnya.
- **Perubahan / peninjauan ulang:** Informasi baru dapat membuat keputusan atau informasi diperbaiki atau dinilai ulang. **Tidak setiap perubahan** mewajibkan peninjauan ulang.
- **BELUM DIPUTUSKAN:** Aturan Product belum diselesaikan. Jika diperlukan untuk hasil tertentu, keputusan harus dibuat oleh pemiliknya; jangan menebak.
- **Tidak otomatis berarti:** Dua keadaan berbeda harus tetap dibedakan, misalnya Campaign terlihat publik tidak otomatis boleh menerima donasi.

Kode `O1–O5`, `P1–P8`, `V1–V7`, dan `A1–A6` hanya penanda bagian peta, bukan urutan pelaksanaan atau persetujuan aturan Product. Peta lengkap memiliki **26 node**. Draf ini baru memuat rincian **13 node**: lima node `J-ORG` dan delapan node `J-CAM`.

## 2. J-ORG — Organisasi dan peninjauan

| Kode | Perubahan atau hasil yang bermakna | Batasan dan keputusan yang masih terbuka |
| --- | --- | --- |
| **O1** | Seseorang hendak membentuk konteks organisasi yang jelas di Kencleng dan hubungan kewenangan **internal** yang dapat diketahui siapa pelakunya. Hasilnya harus jujur, bermakna, dan tetap tersimpan sebagai keadaan yang nyata. | **BELUM DIPUTUSKAN:** siapa yang boleh membentuk organisasi, siapa menerima peran `Owner` pertama, kapan pembentukan dan kewenangan dinyatakan sah, serta arti hasil yang belum lengkap atau dipersengketakan. Product Intent **belum menetapkan pemberian `Owner` otomatis atau keberhasilan gabungan tertentu**. O1 bukan syarat universal bagi tindakan organisasi lain. |
| **O2** | Informasi atau bukti yang diberikan organisasi mempunyai sumber organisasi yang jelas dan diketahui perwakilan yang bertindak jika relevan. | Informasi tetap **berasal dari organisasi**; bukan otomatis fakta platform, hasil peninjauan organisasi, bukti hukum eksternal, atau hasil kurasi Campaign. Informasi bisa diberikan saat atau setelah pembentukan jika kewenangan yang berlaku memungkinkan. Tidak semua pengisian informasi harus melewati O1 lebih dulu. |
| **O3** | Jika peninjauan organisasi berlaku untuk suatu persoalan, ruang lingkup peninjauan, bukti yang relevan dan asalnya, riwayat terkait, serta kewenangan peninjau harus tersedia. | Kapan peninjauan berlaku, bahan minimum, dan akibat hukumnya **BELUM DIPUTUSKAN**. Peninjau dengan `conflict of interest` terhadap perkara terkait tidak boleh menilai. |
| **O4** | Peninjau menyampaikan hasil penilaian sesuai ruang lingkup yang boleh dinilai. Jika belum bisa menilai secara sah, keterbatasan tersebut tetap dinyatakan dengan jujur. | Hasil penilaian **tidak otomatis** memberi label umum bahwa organisasi telah dapat dipercaya. Arti hasil positif, negatif, belum selesai, dan dampak pada Campaign **BELUM DIPUTUSKAN**. |
| **O5** | Perubahan penting, masukan, atau masalah yang belum selesai dapat memicu peninjauan ulang organisasi jika berlaku. | Pemicu peninjauan ulang dan kapan keputusan sebelumnya digantikan **BELUM DIPUTUSKAN**. Tidak setiap perubahan memerlukan peninjauan ulang. |

### Batas J-ORG yang harus tetap dijaga

- `Owner` dan `Staff` memiliki kewenangan berbeda. Kewenangan `Owner` tidak otomatis dimiliki `Staff`.
- Peninjauan organisasi **tidak sama dengan** kurasi Campaign; informasi dari organisasi juga bukan hasil pemeriksaan independen.
- Hubungan antara O1/O2 dan tindakan Campaign mempunyai syarat kewenangan atau informasi masing-masing; O1 tidak otomatis menjadi prasyarat seluruh Journey Campaign.
- Persetujuan atas peta **bukan persetujuan** aturan lengkap O1 atau hasil peninjauan organisasi.
- Hubungan antarnode, diagram, dan keputusan yang belum dibuat harus tetap dibaca dalam [Level 2 v0.3 resmi](level-2-network.md) sampai terjemahan lengkap diaudit dan disetujui.

## 3. J-CAM — Campaign, kurasi, dan keadaan sebenarnya

| Kode | Perubahan atau hasil yang bermakna | Batasan dan keputusan yang masih terbuka |
| --- | --- | --- |
| **P1** | Perwakilan organisasi yang berwenang memberikan usulan Campaign. Asal organisasi dan pihak yang memberikan informasi harus dapat diketahui; klaim dalam usulan tidak otomatis menjadi fakta yang dibuktikan Kencleng. | Orang yang bertindak harus benar-benar memiliki **kewenangan untuk menyampaikan usulan Campaign ini**. Tidak wajib seluruh kegiatan `J-ORG` selesai atau peninjauan organisasi telah dilakukan. Campaign berada di bawah organisasi; hak `Owner` dan `Staff` untuk tindakan spesifik ini **BELUM DIPUTUSKAN**. |
| **P2** | Jika kurasi Campaign memang berlaku, usulan masuk ke peninjauan yang mempunyai cakupan dan dasar yang jelas. | Diperlukan kewenangan peninjau, bukti terkait, riwayat yang relevan, serta larangan `conflict of interest`. Kurasi Campaign berbeda dari peninjauan organisasi. **Kurasi bukan syarat universal** untuk setiap Campaign. |
| **P3** | Hasil kurasi dalam cakupan yang berlaku, masukan perbaikan, atau keadaan ketika peninjau belum dapat memberi penilaian harus dapat dibedakan. | Hasil kurasi positif **tidak otomatis** membuat Campaign tampil publik atau boleh menerima donasi. Penolakan, informasi yang belum cukup, dan permintaan perubahan harus dapat dibedakan. Arti dan akibat tepat dari masing-masing hasil **BELUM DIPUTUSKAN**. |
| **P4** | Organisasi menanggapi masukan, memperbaiki, atau mengirim ulang informasi bila berlaku. Riwayat yang relevan harus menghubungkan perubahan dengan masukan atau hasil penilaian sebelumnya. | Kapan harus dinilai ulang, apakah penolakan bersifat akhir, kapan hasil lama digantikan, dan kapan usulan boleh diajukan kembali **BELUM DIPUTUSKAN**. **Tidak setiap perubahan** wajib dikirim ulang atau dikurasi kembali. |
| **P5** | Campaign mempunyai **keadaan penayangan publik yang sebenarnya**. Pengguna yang berhak melihatnya dapat menemukan informasi Campaign hanya ketika Campaign benar-benar sedang ditayangkan untuk mereka. | Usulan Campaign adalah **informasi yang diperlukan**, bukan izin penayangan. Syarat penayangan, kaitannya dengan kurasi, dan pencabutan penayangan akibat perubahan penting **BELUM DIPUTUSKAN**. |
| **P6** | Campaign mempunyai **keadaan kelayakan menerima donasi yang sebenarnya**, terpisah dari penayangan publik maupun hasil kurasi. | Syarat kelayakan, hubungannya dengan penayangan dan kurasi, serta aturan perubahan kelayakan **BELUM DIPUTUSKAN**. Campaign yang terlihat publik **belum tentu** boleh menerima donasi. |
| **P7** | Campaign memiliki keadaan penggalangan dana dan tahap keberlangsungannya yang benar, termasuk perubahan penting dan kemungkinan penutupan. | Keadaan ini **berbeda** dari status penayangan (P5) dan izin menerima donasi (P6). Syarat aktivasi, makna setiap tahap, serta akibat penutupan **BELUM DIPUTUSKAN**. Keadaan Campaign menjadi konteks yang diperlukan untuk pertanggungjawaban yang berlanjut. |
| **P8** | Perubahan penting pada penayangan publik atau kelayakan menerima donasi harus ditunjukkan sebagai **perubahan keadaan yang benar-benar terjadi**, beserta akibat yang relevan bagi Journey lain. | Aturan perubahan kedua status secara terpisah **BELUM DIPUTUSKAN**. Perubahan dapat memengaruhi akses Pengunjung atau percobaan donasi yang sedang berlangsung; akibat pastinya membutuhkan aturan Product sesuai kasus. **Jangan menganggap izin sebelumnya tetap berlaku setelah keadaan berubah.** |

### Batas J-CAM yang harus tetap dijaga

- Kewenangan perwakilan organisasi untuk bertindak harus nyata. Tidak seluruh tahapan `J-ORG` wajib selesai dan peninjauan organisasi tidak otomatis menjadi syarat setiap tindakan Campaign.
- **Usulan Campaign ≠ hasil kurasi ≠ penayangan publik ≠ izin menerima donasi.** Setiap hasil harus mengikuti keadaan dan aturan yang memang berlaku.
- Kurasi hanya dilakukan **jika berlaku**. Hasil kurasi positif tidak otomatis mengizinkan penayangan maupun donasi.
- Penolakan, permintaan revisi, informasi yang belum cukup, perubahan keadaan, dan peninjauan ulang bukan satu keadaan yang sama. Aturan kapan masing-masing berlaku masih terbuka.
- Jika status Campaign berubah ketika seseorang sedang mempertimbangkan atau mencoba berdonasi, **jangan menganggap kelayakan sebelumnya tetap berlaku**. Akibat tepat pada percobaan donasi masih membutuhkan keputusan Product.
- Hubungan `J-CAM` dengan `J-VIS`, `J-ACC`, serta diagram dan seluruh hubungan antarnode tetap dibaca dari [Level 2 v0.3 resmi](level-2-network.md) sampai draf lengkap ini diaudit dan disetujui.

## 4. Bagian yang belum disusun

Draf ini belum mencakup `J-VIS` (V1–V7), `J-ACC` (A1–A6), `EX-OPS`, diagram dan hubungan silang, daftar kondisi gagal/belum pasti, maupun daftar keputusan terbuka untuk seluruh Journey. Semua masih berlaku sesuai [Level 2 v0.3](level-2-network.md).

**Jangan gunakan draf parsial ini sebagai pengganti peta Level 2 atau sebagai masukan resmi Rancang.**
