# Kencleng — Interconnected Product Journeys (Level 2)

> **Authority:** Level 2 Interconnected Product Journeys approved by Anhar on 22 Oktober 2026, scoped to `pilot/station5-c1-space-only` (not operational `main`). Output type: `PROD-JOURNEYS`. Identifier migration on 23 Oktober 2026 was authorized as a structural change; [old/new mappings](../identifier-migration-2026-10-23.md) preserve source lineage.
> **Riwayat dan sumber yang diaudit:** Level 2 Inggris terakhir, blob `4970ba936d107e50d8260c9e246816ba13d41a28`; draf Bahasa Indonesia yang disetujui Anhar, blob `2ae3ed81651fe3f56595128617edb5244cb9d86c`; sumber koreksi v0.3 awal yang disetujui sebelumnya, blob `87edca25b40017429d66aff4e22ad1f4aa7e013f`. Seluruh versi Inggris tetap dapat ditelusuri melalui riwayat Git.
> **Authority dan lingkup:** [Product Intent](../product-intent.md) memiliki authority nilai dan prinsip. Journey map ini bukan Product Contract atau Engineering roadmap. Untuk node `J-ORG-01`, gunakan [approved Contract ORG-01](../outcomes/org-01/product-contract.md). Node lain tidak otomatis approved.
> **Batas lingkungan:** Kencleng `main`, Harscode, Rancang, Product Specification Space, dan Engineering **tidak** berubah atau otomatis diaktifkan karena promosi editorial ini.

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

Kode `J-ORG-01–J-ORG-05`, `J-CAM-01–J-CAM-08`, `J-VIS-01–J-VIS-07`, dan `J-ACC-01–J-ACC-06` hanya penanda bagian peta, bukan urutan pelaksanaan atau persetujuan aturan Product. Peta ini memuat seluruh **26 node**: lima `J-ORG`, delapan `J-CAM`, tujuh `J-VIS`, dan enam `J-ACC`. Hubungan silang, dua diagram, pengecualian `EX-OPS`, pemeriksaan batas makna, serta pertanyaan Product OPEN dijelaskan di bagian berikutnya.

## 2. J-ORG — Organisasi dan peninjauan

| Kode | Perubahan atau hasil yang bermakna | Batasan dan keputusan yang masih terbuka |
| --- | --- | --- |
| **J-ORG-01** | Seseorang membentuk konteks organisasi dan hubungan `Owner` internal pertama yang teratribusi, berlaku bersama saat berhasil serta dapat diperiksa kembali. | **Outcome `ORG-01` sudah disetujui** lewat [Product Contract](../outcomes/org-01/product-contract.md). Persetujuan Reviewer atas organisasi tetap terpisah dan kemiripan klaim/nama tidak otomatis memblokir dua konteks berbeda ID. |
| **J-ORG-02** | Informasi atau bukti yang diberikan organisasi mempunyai sumber organisasi yang jelas dan diketahui perwakilan yang bertindak jika relevan. | Informasi tetap **berasal dari organisasi**; bukan otomatis fakta platform, hasil peninjauan organisasi, bukti hukum eksternal, atau hasil kurasi Campaign. Informasi bisa diberikan saat atau setelah pembentukan jika kewenangan yang berlaku memungkinkan. Tidak semua pengisian informasi harus melewati J-ORG-01 lebih dulu. |
| **J-ORG-03** | Jika peninjauan organisasi berlaku untuk suatu persoalan, ruang lingkup peninjauan, bukti yang relevan dan asalnya, riwayat terkait, serta kewenangan peninjau harus tersedia. | Kapan peninjauan berlaku, bahan minimum, dan akibat hukumnya **BELUM DIPUTUSKAN**. Peninjau dengan `conflict of interest` terhadap perkara terkait tidak boleh menilai. |
| **J-ORG-04** | Peninjau menyampaikan hasil penilaian sesuai ruang lingkup yang boleh dinilai. Jika belum bisa menilai secara sah, keterbatasan tersebut tetap dinyatakan dengan jujur. | Hasil penilaian **tidak otomatis** memberi label umum bahwa organisasi telah dapat dipercaya. Arti hasil positif, negatif, belum selesai, dan dampak pada Campaign **BELUM DIPUTUSKAN**. |
| **J-ORG-05** | Perubahan penting, masukan, atau masalah yang belum selesai dapat memicu peninjauan ulang organisasi jika berlaku. | Pemicu peninjauan ulang dan kapan keputusan sebelumnya digantikan **BELUM DIPUTUSKAN**. Tidak setiap perubahan memerlukan peninjauan ulang. |

### Batas J-ORG yang harus tetap dijaga

- `Owner` dan `Staff` memiliki kewenangan berbeda. Kewenangan `Owner` tidak otomatis dimiliki `Staff`.
- Peninjauan organisasi **tidak sama dengan** kurasi Campaign; informasi dari organisasi juga bukan hasil pemeriksaan independen.
- Hubungan antara J-ORG-01/J-ORG-02 dan tindakan Campaign mempunyai syarat kewenangan atau informasi masing-masing; J-ORG-01 tidak otomatis menjadi prasyarat seluruh Journey Campaign.
- Persetujuan atas peta sendiri bukan approval Product Contract. **ORG-01** telah mendapat approval Product Contract terpisah dan tidak mengesahkan hasil peninjauan organisasi.
- Hubungan antarnode dan keputusan yang belum dibuat pada J-ORG harus dibaca bersama **diagram (bagian 6), hubungan penting (bagian 7), dan pertanyaan OPEN (bagian 9)** pada peta ini.

## 3. J-CAM — Campaign, kurasi, dan keadaan sebenarnya

| Kode | Perubahan atau hasil yang bermakna | Batasan dan keputusan yang masih terbuka |
| --- | --- | --- |
| **J-CAM-01** | Perwakilan organisasi yang berwenang memberikan usulan Campaign. Asal organisasi dan pihak yang memberikan informasi harus dapat diketahui; klaim dalam usulan tidak otomatis menjadi fakta yang dibuktikan Kencleng. | Orang yang bertindak harus benar-benar memiliki **kewenangan untuk menyampaikan usulan Campaign ini**. Tidak wajib seluruh kegiatan `J-ORG` selesai atau peninjauan organisasi telah dilakukan. Campaign berada di bawah organisasi; hak `Owner` dan `Staff` untuk tindakan spesifik ini **BELUM DIPUTUSKAN**. |
| **J-CAM-02** | Jika kurasi Campaign memang berlaku, usulan masuk ke peninjauan yang mempunyai cakupan dan dasar yang jelas. | Diperlukan kewenangan peninjau, bukti terkait, riwayat yang relevan, serta larangan `conflict of interest`. Kurasi Campaign berbeda dari peninjauan organisasi. **Kurasi bukan syarat universal** untuk setiap Campaign. |
| **J-CAM-03** | Hasil kurasi dalam cakupan yang berlaku, masukan perbaikan, atau keadaan ketika peninjau belum dapat memberi penilaian harus dapat dibedakan. | Hasil kurasi positif **tidak otomatis** membuat Campaign tampil publik atau boleh menerima donasi. Penolakan, informasi yang belum cukup, dan permintaan perubahan harus dapat dibedakan. Arti dan akibat tepat dari masing-masing hasil **BELUM DIPUTUSKAN**. |
| **J-CAM-04** | Organisasi menanggapi masukan, memperbaiki, atau mengirim ulang informasi bila berlaku. Riwayat yang relevan harus menghubungkan perubahan dengan masukan atau hasil penilaian sebelumnya. | Kapan harus dinilai ulang, apakah penolakan bersifat akhir, kapan hasil lama digantikan, dan kapan usulan boleh diajukan kembali **BELUM DIPUTUSKAN**. **Tidak setiap perubahan** wajib dikirim ulang atau dikurasi kembali. |
| **J-CAM-05** | Campaign mempunyai **keadaan penayangan publik yang sebenarnya**. Pengguna yang berhak melihatnya dapat menemukan informasi Campaign hanya ketika Campaign benar-benar sedang ditayangkan untuk mereka. | Usulan Campaign adalah **informasi yang diperlukan**, bukan izin penayangan. Syarat penayangan, kaitannya dengan kurasi, dan pencabutan penayangan akibat perubahan penting **BELUM DIPUTUSKAN**. |
| **J-CAM-06** | Campaign mempunyai **keadaan kelayakan menerima donasi yang sebenarnya**, terpisah dari penayangan publik maupun hasil kurasi. | Syarat kelayakan, hubungannya dengan penayangan dan kurasi, serta aturan perubahan kelayakan **BELUM DIPUTUSKAN**. Campaign yang terlihat publik **belum tentu** boleh menerima donasi. |
| **J-CAM-07** | Campaign memiliki keadaan penggalangan dana dan tahap keberlangsungannya yang benar, termasuk perubahan penting dan kemungkinan penutupan. | Keadaan ini **berbeda** dari status penayangan (J-CAM-05) dan izin menerima donasi (J-CAM-06). Syarat aktivasi, makna setiap tahap, serta akibat penutupan **BELUM DIPUTUSKAN**. Keadaan Campaign menjadi konteks yang diperlukan untuk pertanggungjawaban yang berlanjut. |
| **J-CAM-08** | Perubahan penting pada penayangan publik atau kelayakan menerima donasi harus ditunjukkan sebagai **perubahan keadaan yang benar-benar terjadi**, beserta akibat yang relevan bagi Journey lain. | Aturan perubahan kedua status secara terpisah **BELUM DIPUTUSKAN**. Perubahan dapat memengaruhi akses Pengunjung atau percobaan donasi yang sedang berlangsung; akibat pastinya membutuhkan aturan Product sesuai kasus. **Jangan menganggap izin sebelumnya tetap berlaku setelah keadaan berubah.** |

### Batas J-CAM yang harus tetap dijaga

- Kewenangan perwakilan organisasi untuk bertindak harus nyata. Tidak seluruh tahapan `J-ORG` wajib selesai dan peninjauan organisasi tidak otomatis menjadi syarat setiap tindakan Campaign.
- **Usulan Campaign ≠ hasil kurasi ≠ penayangan publik ≠ izin menerima donasi.** Setiap hasil harus mengikuti keadaan dan aturan yang memang berlaku.
- Kurasi hanya dilakukan **jika berlaku**. Hasil kurasi positif tidak otomatis mengizinkan penayangan maupun donasi.
- Penolakan, permintaan revisi, informasi yang belum cukup, perubahan keadaan, dan peninjauan ulang bukan satu keadaan yang sama. Aturan kapan masing-masing berlaku masih terbuka.
- Jika status Campaign berubah ketika seseorang sedang mempertimbangkan atau mencoba berdonasi, **jangan menganggap kelayakan sebelumnya tetap berlaku**. Akibat tepat pada percobaan donasi masih membutuhkan keputusan Product.
- Hubungan `J-CAM` dengan `J-VIS`, `J-ACC`, serta diagram dan batas lainnya dibaca bersama **bagian 6–9** pada peta ini.

## 4. J-VIS — Pengunjung, keputusan berdonasi, dan keadaan uang

| Kode | Perubahan atau hasil yang bermakna | Batasan dan keputusan yang masih terbuka |
| --- | --- | --- |
| **J-VIS-01** | Pengunjung dapat mengakses informasi Campaign yang **benar-benar terbuka untuk publik**, termasuk organisasi yang bertanggung jawab, tujuan, asal informasi, keadaan Campaign, arti informasi pendanaan, dan hal yang belum diketahui. | Memerlukan status penayangan sebenarnya (J-CAM-05) dan informasi terkait yang dapat diketahui asalnya. Hasil peninjauan dengan cakupan tertentu dapat membantu Pengunjung memahami informasi **bila berlaku**, tetapi bukan syarat universal. Pengunjung yang melihat Campaign publik **tidak harus belum memiliki akun**. |
| **J-VIS-02** | Informasi yang relevan dan bisa dipahami membantu Pengunjung membuat pilihan secara sadar, tanpa tekanan untuk berdonasi. | Kencleng **tidak dapat menjamin apa yang benar-benar dipahami di dalam pikiran Pengunjung**. Informasi yang dianggap cukup bukan berarti semua informasi harus ditampilkan. Informasi penting yang wajib dijelaskan sebelum donasi masih memerlukan aturan Product sesuai cakupannya (**BELUM DIPUTUSKAN**). |
| **J-VIS-03** | Pengunjung boleh **tidak berdonasi** atau **menunda keputusan**, dan keduanya tetap merupakan hasil pilihan yang wajar, bukan kegagalan Product. | Tidak otomatis perlu menyimpan catatan bahwa Pengunjung menolak atau menunda donasi. Jangan menambahkan kewajiban pencatatan atau mendorong donasi dengan tekanan. |
| **J-VIS-04** | Jika Pengunjung memilih mencoba berdonasi, ia dapat memulai **percobaan donasi** hanya saat syarat kelayakan Campaign yang benar-benar berlaku mengizinkan tindakan itu. | Izin menerima donasi (J-CAM-06) adalah syarat **khusus untuk percobaan donasi**, bukan syarat untuk sekadar memahami Campaign atau memilih tidak berdonasi. Jika status Campaign berubah selama percobaan, aturan penerimaan atau penolakannya **BELUM DIPUTUSKAN**. Keinginan berdonasi **bukan** bukti bahwa percobaan telah dimulai. |
| **J-VIS-05** | Hasil percobaan donasi harus dinyatakan sesuai bukti yang ada: **belum selesai atau masih menunggu kepastian**, **gagal**, atau **memiliki fakta donasi yang diakui** jika fakta tersebut memang ada. | Percobaan tidak membuktikan bahwa donasi telah terjadi. Status belum pasti atau gagal **tidak boleh** ditampilkan sebagai donasi berhasil, hubungan Donatur, atau keadaan pendanaan yang sudah terwujud. Syarat tepat pengakuan donasi masih membutuhkan keputusan Product. |
| **J-VIS-06** | Fakta donasi yang benar-benar diakui dan terkait dengan Campaign membentuk **hubungan Donatur dengan Campaign tersebut**. | Perbedaan Donatur yang memakai akun dan yang berdonasi sebagai tamu adalah persoalan tersendiri; rincian hak atau pengalaman keduanya **BELUM DIPUTUSKAN**. Memiliki akun atau sekadar mencoba berdonasi **tidak otomatis** menjadikan seseorang Donatur. |
| **J-VIS-07** | Pengunjung dan Donatur dapat memahami keadaan uang yang **benar-benar didukung bukti**. Fakta donasi, pendanaan Campaign, dan penyelesaian transaksi uang (**settlement**) harus dibedakan. | Fakta donasi **tidak otomatis** membuktikan pendanaan atau settlement, apalagi pelaksanaan kegiatan dan dampak nyata. Pembayaran simulasi **harus disebut sebagai simulasi**, bukan sebagai uang yang benar-benar diselesaikan di luar sistem. Sumber bukti, pemicu, dan aturan tepat setiap keadaan **BELUM DIPUTUSKAN**. |

### Batas J-VIS yang harus tetap dijaga

- **Memahami Campaign ≠ memilih berdonasi ≠ memulai percobaan donasi.** Pengunjung boleh memutuskan tidak berdonasi atau menunda tanpa diperlakukan sebagai kegagalan.
- **Percobaan donasi ≠ fakta donasi yang diakui ≠ hubungan Donatur.** Keadaan belum pasti atau gagal tidak boleh dinaikkan menjadi keberhasilan.
- **Fakta donasi ≠ pendanaan Campaign ≠ settlement ≠ pelaksanaan kegiatan atau dampak nyata.** Setiap klaim hanya boleh mengikuti keadaan dan informasi yang benar-benar tersedia.
- Penayangan publik (J-CAM-05) berbeda dari kelayakan menerima donasi (J-CAM-06). Pengunjung tidak perlu memenuhi kelayakan donasi hanya untuk memahami Campaign atau memilih tidak berdonasi.
- Jika kelayakan Campaign berubah di tengah percobaan, jangan menganggap izin lama otomatis masih berlaku. **Aturan hasil percobaan tetap BELUM DIPUTUSKAN**, bukan ditentukan oleh peta riset ini.
- Hubungan antara Donatur, pendanaan, dan pertanggungjawaban serta diagram lintas-Journey dijelaskan di **bagian 6–9** pada peta ini.

## 5. J-ACC — Keadaan Campaign dan pertanggungjawaban berkelanjutan

| Kode | Perubahan atau hasil yang bermakna | Batasan dan keputusan yang masih terbuka |
| --- | --- | --- |
| **J-ACC-01** | Perubahan penting pada keadaan penggalangan dana atau penutupan Campaign mengubah konteks Campaign. Tanggung jawab untuk menjelaskan perkembangan dan hasil tetap bermakna. | **Campaign ditutup ≠ pertanggungjawaban selesai.** Pemicu penutupan dan akibat tepatnya bagi Campaign serta pihak terkait **BELUM DIPUTUSKAN**. |
| **J-ACC-02** | Informasi atau laporan pertanggungjawaban organisasi yang relevan tersedia. Jika belum tersedia atau masih menunggu, keadaan itu tetap dijelaskan **sesuai kenyataannya**, bukan dinyatakan sudah lengkap. | Pertanggungjawaban dapat bermakna **sebelum Campaign ditutup dan meskipun belum ada Donatur**. Kewajiban minimum organisasi, pemicu dan waktu pelaporan, siapa yang berhak melihatnya, serta kondisi kapan laporan dianggap belum dipenuhi **BELUM DIPUTUSKAN**. Status belum tersedia tidak otomatis berarti pelanggaran. |
| **J-ACC-03** | Perkembangan kegiatan atau hasil yang dilaporkan organisasi dapat diketahui sebagai **informasi yang berasal dari organisasi**, beserta pihak yang bertanggung jawab atas informasi tersebut. | **Laporan organisasi ≠ verifikasi independen ≠ dampak nyata yang terbukti.** Asal laporan dan kewenangan `Owner` atau `Staff` yang bertindak tetap dibedakan. Laporan tidak otomatis berubah menjadi fakta yang telah diverifikasi pihak lain. |
| **J-ACC-04** | Jika informasi baru muncul, dikoreksi, diperselisihkan, atau berubah, urutan kejadian, asal informasi, dan ketidakpastian yang penting tetap bisa dipahami. | Ini **tidak otomatis mewajibkan penyimpanan seluruh catatan teknis mentah (audit log)** atau satu laporan akhir yang tak boleh berubah. Kapan keputusan/laporan lama digantikan dan kapan persoalan dibuka kembali **BELUM DIPUTUSKAN**. |
| **J-ACC-05** | Jika verifikasi penggunaan dana oleh peninjau memang berlaku, peninjau dapat memberi **hasil verifikasi terbatas sesuai cakupan dan kewenangannya**. Jika penilaian yang sah belum bisa diberikan, keadaan itu harus tetap jelas. | Kewenangan, independensi peninjau, `conflict of interest`, bukti yang diperlukan, kapan verifikasi wajib dilakukan, dan arti hasilnya **BELUM DIPUTUSKAN**. Verifikasi terbatas **tidak otomatis membuktikan seluruh dampak nyata**. |
| **J-ACC-06** | Pihak yang memang berhak mengakses informasi, yaitu publik dan/atau Donatur **sesuai aturan yang berlaku**, dapat terus memahami laporan, fakta yang diketahui, hasil peninjauan sesuai cakupannya, serta ketidakpastian yang masih ada. | Hubungan Donatur yang benar-benar ada dapat memberi kebutuhan informasi tambahan, tetapi **Donatur bukan syarat agar pertanggungjawaban organisasi bermakna**. Siapa yang berhak melihat informasi, akses tamu, dan kondisi kapan pertanggungjawaban dianggap selesai **BELUM DIPUTUSKAN**. Tidak semua laporan otomatis bersifat publik. |

### Batas J-ACC yang harus tetap dijaga

- **Penutupan Campaign ≠ selesai pertanggungjawaban.** Kewajiban dan keterbukaan informasi dapat tetap relevan sebelum maupun setelah penutupan, termasuk ketika jumlah Donatur masih nol.
- **Laporan organisasi ≠ verifikasi independen ≠ dampak nyata yang terbukti.** Asal klaim, kewenangan penulis, dan hasil peninjauan harus tetap dapat dibedakan.
- **Laporan belum tersedia ≠ otomatis terjadi pelanggaran.** Jangan menciptakan tenggat, pelanggaran, atau status selesai yang belum diputuskan Product.
- Informasi yang berubah harus mempertahankan **urutan kejadian, asal informasi, dan ketidakpastian yang penting**, tanpa menentukan bentuk audit log atau aturan penggantian laporan secara teknis.
- Verifikasi penggunaan dana hanya dilakukan **jika berlaku** dan dengan kewenangan serta independensi peninjau yang sesuai. Hasilnya terbatas pada apa yang benar-benar diperiksa.
- Akses publik, Donatur, dan tamu **tidak otomatis sama**. Hubungan Donatur bukan prasyarat adanya pertanggungjawaban organisasi.
- Hubungan `J-ACC` dengan keadaan Campaign (J-CAM-07), status Donatur (J-VIS-06), dan bagian lainnya juga harus dibaca bersama **bagian 6–9** pada peta ini.

## 6. Hubungan antar-Journey — baca sebagai jaringan, bukan urutan development

Dua diagram ini membantu melihat hubungan antarbagian. **Panah utuh** menunjukkan keadaan atau informasi yang diperlukan, atau kelanjutan yang masuk akal di satu bagian Journey. Panah utuh **tidak** berarti keadaan di sisi asal sudah cukup untuk membuktikan keberhasilan di sisi tujuan. **Panah putus-putus** menunjukkan hubungan bersyarat, peninjauan ulang, atau aturan yang masih **BELUM DIPUTUSKAN**. Jenis hubungan yang lebih tepat dijelaskan pada daftar hubungan di bagian 7.

Diagram **tidak** membuktikan urutan implementasi wajib, hubungan sebab-akibat yang sudah diuji, atau bahwa semua syarat untuk menjalankan aplikasi sudah tersedia.

### 6.1 J-ORG ↔ J-CAM

```mermaid
flowchart TB
    ORG_01["J-ORG-01 Konteks organisasi dan Owner internal"]
    ORG_02["J-ORG-02 Informasi organisasi jelas sumbernya"]
    ORG_03["J-ORG-03 Persoalan peninjauan organisasi"]
    ORG_04["J-ORG-04 Hasil peninjauan sesuai cakupan"]
    ORG_05["J-ORG-05 Perubahan / peninjauan ulang"]
    CAM_01["J-CAM-01 Usulan Campaign jelas sumbernya"]
    CAM_02["J-CAM-02 Konteks kurasi Campaign"]
    CAM_03["J-CAM-03 Hasil kurasi sesuai cakupan"]
    CAM_04["J-CAM-04 Perbaikan / riwayat peninjauan"]
    CAM_05["J-CAM-05 Penayangan publik sebenarnya"]
    CAM_06["J-CAM-06 Kelayakan donasi sebenarnya"]
    CAM_07["J-CAM-07 Keadaan penggalangan dana"]
    CAM_08["J-CAM-08 Perubahan penayangan / kelayakan"]

    ORG_01 -->|"Jalur kewenangan awal yang sah"| ORG_02
    ORG_01 -->|"Kewenangan bertindak bila relevan"| CAM_01
    ORG_02 -->|"Asal informasi dan sumbernya"| CAM_01
    ORG_02 -. "Jika peninjauan organisasi berlaku" .-> ORG_03
    ORG_03 --> ORG_04
    ORG_04 -. "Masukan / perubahan terkait" .-> ORG_05
    ORG_05 -. "Jika perlu dinilai ulang" .-> ORG_03
    ORG_04 -. "Kemungkinan akibat lanjutan BELUM DIPUTUSKAN" .-> CAM_02
    CAM_01 -. "Jika kurasi berlaku" .-> CAM_02
    CAM_02 --> CAM_03
    CAM_03 -. "Jika tanggapan diperlukan" .-> CAM_04
    CAM_04 -. "Jika pengajuan kembali berlaku" .-> CAM_02
    CAM_01 -->|"Informasi perlu, belum tentu cukup"| CAM_05
    CAM_01 -->|"Konteks Campaign, belum tentu cukup"| CAM_06
    CAM_01 -->|"Konteks Campaign, bukan aktivasi"| CAM_07
    CAM_03 -. "Aturan penayangan BELUM DIPUTUSKAN" .-> CAM_05
    CAM_03 -. "Aturan kelayakan BELUM DIPUTUSKAN" .-> CAM_06
    CAM_05 -. "Perubahan penting" .-> CAM_08
    CAM_06 -. "Perubahan penting" .-> CAM_08
```

### 6.2 J-CAM ↔ J-VIS ↔ J-ACC

```mermaid
flowchart TB
    CAM_05["J-CAM-05 Keadaan publik sebenarnya"]
    CAM_06["J-CAM-06 Kelayakan sebenarnya"]
    CAM_07["J-CAM-07 Konteks penggalangan dana"]
    CAM_08["J-CAM-08 Perubahan keadaan sebenarnya"]
    VIS_01["J-VIS-01 Informasi publik yang jujur"]
    VIS_02{"VIS_02 Pilihan yang dipertimbangkan"}
    VIS_03["J-VIS-03 Tidak berdonasi / menunda"]
    VIS_04["J-VIS-04 Percobaan donasi yang diizinkan"]
    VIS_05["J-VIS-05 Belum pasti / gagal / fakta"]
    VIS_06["J-VIS-06 Hubungan Donatur"]
    VIS_07["J-VIS-07 Fakta pendanaan dan settlement"]
    ACC_01["J-ACC-01 Perubahan tahap / penutupan"]
    ACC_02["J-ACC-02 Informasi pertanggungjawaban atau ketidaktersediaan yang jujur"]
    ACC_03["J-ACC-03 Perkembangan laporan organisasi"]
    ACC_04["J-ACC-04 Urutan kejadian / perubahan"]
    ACC_05["J-ACC-05 Verifikasi sesuai cakupan"]
    ACC_06["J-ACC-06 Pemahaman berkelanjutan"]

    CAM_05 -->|"Keadaan publik sebenarnya"| VIS_01
    VIS_07 -. "Status pendanaan bila relevan" .-> VIS_01
    VIS_01 --> VIS_02
    VIS_02 --> VIS_03
    VIS_02 -. "Memilih mencoba berdonasi" .-> VIS_04
    CAM_06 -->|"Kelayakan yang berlaku"| VIS_04
    CAM_08 -. "Perubahan dapat memengaruhi pilihan berjalan" .-> VIS_01
    CAM_08 -. "Akibat saat percobaan berlangsung BELUM DIPUTUSKAN" .-> VIS_04
    CAM_08 -. "Aturan hasil BELUM DIPUTUSKAN" .-> VIS_05
    VIS_04 --> VIS_05
    VIS_05 -. "Hanya jika fakta donasi sah" .-> VIS_06
    VIS_06 -. "Tidak otomatis menjadi pendanaan atau settlement" .-> VIS_07
    CAM_07 -->|"Perubahan keadaan Campaign"| ACC_01
    CAM_07 -->|"Tidak bergantung pada Donatur atau penutupan"| ACC_02
    ACC_01 -->|"Perubahan konteks sebenarnya"| ACC_04
    ACC_02 -. "Jika laporan tersedia" .-> ACC_03
    ACC_02 -->|"Belum tersedia / belum pasti tetap jujur"| ACC_06
    ACC_03 --> ACC_04
    ACC_03 -. "Jika verifikasi berlaku" .-> ACC_05
    ACC_04 --> ACC_06
    ACC_05 -. "Hasil penilaian sesuai cakupan" .-> ACC_06
    VIS_06 -. "Kebutuhan lanjutan khusus Donatur" .-> ACC_06
```

### 6.3 EX-OPS — Pengecualian tindakan pengelola platform

`EX-OPS` bukan Journey kelima, prasyarat yang harus dilewati seluruh node, ataupun kewenangan untuk membuat fakta Product yang tidak didukung bukti. Intervensi khusus pengelola platform melintasi Journey lain **hanya dalam batas kewenangan yang sah dan keputusan Product yang berlaku**.

Kewenangan operasional pengelola platform **tidak otomatis** menjadi kewenangan organisasi, peninjau, verifikator penggunaan dana, atau pihak yang dapat menyatakan donasi, settlement, maupun dampak nyata. Batas kewenangan positif dan tindakan koreksi yang boleh dilakukan pengelola platform **BELUM DIPUTUSKAN**.

## 7. Hubungan penting antarnode

| Hubungan | Jenis hubungan | Kesimpulan yang **tidak boleh** ditarik |
| --- | --- | --- |
| Kewenangan perwakilan yang sah → J-CAM-01 | **Syarat kewenangan / didukung peta** | Organisasi yang sudah memiliki kewenangan sah wajib mengulangi J-ORG-01, atau seluruh J-ORG harus diselesaikan. |
| J-ORG-02 → J-CAM-01/J-VIS-01 | **Syarat informasi / didukung peta** | Informasi dari organisasi otomatis telah diverifikasi independen. |
| J-ORG-04 → J-CAM-02/J-CAM-05/J-CAM-06 | **Bersyarat / BELUM DIPUTUSKAN** | Seluruh pekerjaan Campaign mensyaratkan hasil peninjauan organisasi yang selesai. |
| J-CAM-01 → J-CAM-02 → J-CAM-03 ⇄ J-CAM-04 | **Kurasi bersyarat / masukan dan perbaikan** | Semua Campaign selalu melalui kurasi atau setiap perubahan wajib ditinjau ulang. |
| J-CAM-01 → J-CAM-05 | **Informasi yang diperlukan, belum cukup** | Usulan Campaign otomatis mengizinkan penayangan publik. |
| J-CAM-03/J-CAM-04 → J-CAM-05/J-CAM-06 | **Bersyarat / BELUM DIPUTUSKAN** | Hasil kurasi positif otomatis mengizinkan penayangan atau donasi. |
| J-CAM-05 → J-VIS-01 | **Syarat keadaan nyata / didukung peta** | Campaign yang publik otomatis layak menerima donasi. |
| J-CAM-06 → J-VIS-04 | **Syarat khusus pada percobaan donasi / didukung peta** | Kelayakan donasi diwajibkan juga untuk melihat Campaign atau memilih tidak berdonasi. |
| J-CAM-08 → J-VIS-01/J-VIS-04/J-VIS-05 | **Bersyarat / akibat BELUM DIPUTUSKAN** | Percobaan donasi yang sedang berjalan tetap mengikuti kelayakan lama setelah keadaan berubah. |
| J-VIS-01 → J-VIS-02 → J-VIS-03 | **Pilihan Pengunjung yang sah / didukung peta** | Tidak berdonasi adalah kegagalan Product atau wajib dicatat sebagai penolakan. |
| J-VIS-04 → J-VIS-05 → J-VIS-06 | **Hanya fakta donasi sah membentuk Donatur / didukung peta** | Niat, percobaan, keadaan menunggu, atau gagal langsung membentuk hubungan Donatur. |
| J-VIS-06 ⇢ J-VIS-07 | **Tidak ada akibat uang yang otomatis / didukung peta** | Hubungan Donatur atau fakta donasi membuktikan pendanaan atau settlement. |
| J-CAM-07 → J-ACC-01/J-ACC-02 | **Keadaan Campaign dan pertanggungjawaban / didukung peta** | Campaign harus ditutup atau memiliki Donatur dahulu sebelum pertanggungjawaban berlaku. |
| J-ACC-02/J-ACC-03/J-ACC-04 → J-ACC-06 | **Informasi berkelanjutan / didukung peta; hak akses BELUM DIPUTUSKAN** | Semua laporan otomatis publik, atau setiap laporan yang belum tersedia adalah pelanggaran. |
| J-ACC-03 → J-ACC-05 → J-ACC-06 | **Verifikasi bersyarat / perbedaan makna didukung; penerapan BELUM DIPUTUSKAN** | Laporan otomatis membuktikan dampak terverifikasi atau semua laporan wajib diverifikasi. |
| J-VIS-06 → J-ACC-06 | **Kebutuhan lanjutan Donatur / didukung peta** | Pertanggungjawaban organisasi hanya bermakna jika sudah ada Donatur. |
| EX-OPS → node lain | **Pengecualian lintas-Journey, bukan syarat normal** | Pengelola platform boleh membuat fakta organisasi, hasil peninjauan, settlement, atau dampak tanpa kewenangan/bukti yang sah. |

Hubungan berstatus **BELUM DIPUTUSKAN** adalah masalah Product yang sengaja terlihat, **bukan** aturan tersembunyi yang boleh diisi berdasarkan asumsi. Panah pada diagram tidak membuktikan ketergantungan teknis yang sudah diuji dan bukan urutan wajib development.

## 8. Pemeriksaan Batas Makna Journey

Bagian ini memeriksa apakah peta tetap menggambarkan keadaan Product secara benar ketika terjadi penolakan, perubahan, kegagalan, keputusan alternatif, atau ketidakpastian. Skenario-skenario ini digunakan untuk menguji **batas makna peta**, bukan menguji software, menganggap setiap kondisi sebagai kegagalan, atau menetapkan aturan Product baru.

| Skenario yang diperiksa | Makna yang harus tetap dijaga |
| --- | --- |
| Pembentukan organisasi tanpa atribusi atau hasil efektif yang lengkap | Outcome `ORG-01` tidak boleh mengklaim keberhasilan dari organisasi atau `Owner` yang belum sah/teratribusi menurut [Product Contract approved](../outcomes/org-01/product-contract.md). Keadaan parsial bukan sukses penuh. |
| Kemiripan nama/klaim antar-organisasi yang terdaftar | Dua konteks dengan ID internal terpisah boleh berhasil sendiri dan memiliki `Owner` internal masing-masing; Reviewer menilai klaim/status organisasi terpisah, tanpa auto-merge/block/close. |
| `Staff` bertindak di luar kewenangan | J-ORG-02/J-CAM-01/J-CAM-04/J-ACC-02/J-ACC-03 tidak otomatis memiliki kewenangan `Owner`. |
| Peninjau mempunyai `conflict of interest`, tidak cukup kewenangan atau bukti | J-ORG-03/J-ORG-04, J-CAM-02/J-CAM-03, dan J-ACC-05 tidak boleh menghasilkan penilaian sah yang sebenarnya tidak didukung. |
| Peninjauan organisasi belum selesai, negatif, atau perlu dinilai ulang | J-ORG-04/J-ORG-05 tetap sesuai batas keputusannya; akibat terhadap Campaign **BELUM DIPUTUSKAN**. |
| Kurasi ditolak, diminta perbaikan, atau ada perubahan penting | J-CAM-03/J-CAM-04 menjaga riwayat terkait; aturan peninjauan ulang masih **BELUM DIPUTUSKAN**, bukan otomatis. |
| Campaign publik tetapi tidak boleh menerima donasi | J-CAM-05 dan J-VIS-01–J-VIS-03 tetap bermakna; percobaan donasi J-VIS-04 tidak dapat dilakukan dalam kondisi ini. |
| Status penayangan/kelayakan berubah saat Pengunjung mempertimbangkan atau mencoba donasi | J-CAM-08 harus menunjukkan perubahan nyata; akibat tepat terhadap J-VIS-04/J-VIS-05 **BELUM DIPUTUSKAN**. |
| Pengunjung memutuskan tidak berdonasi atau menunda | J-VIS-03 merupakan hasil pilihan yang sah; tidak boleh ada tekanan atau kewajiban pencatatan yang dibuat-buat. |
| Percobaan donasi masih menunggu atau gagal | J-VIS-05 tidak boleh dinaikkan menjadi J-VIS-06 atau memberi kesan J-VIS-07 sudah terpenuhi. |
| Fakta donasi ada tanpa bukti pendanaan/settlement | J-VIS-06 dan J-VIS-07 berbeda; simulasi tidak boleh dianggap bukti uang nyata. |
| Informasi pendanaan terlihat oleh Pengunjung | J-VIS-01 hanya boleh menampilkan keadaan, asal informasi, dan ketidakpastian yang benar-benar didukung. |
| Campaign ditutup atau dilaporkan saat belum ada Donatur | J-CAM-07/J-ACC-01/J-ACC-02/J-ACC-06 tetap bermakna sekalipun J-VIS-06 tidak terjadi. |
| Laporan belum lengkap atau belum tersedia | J-ACC-02/J-ACC-06 harus jujur; jangan menciptakan tenggat, pelanggaran, atau status selesai yang belum diputuskan. |
| Laporan tanpa verifikasi independen | J-ACC-03 berbeda dari J-ACC-05; J-ACC-06 tidak boleh menyebut dampak sudah terbukti. |
| Laporan atau penilaian sebelumnya berubah | J-ACC-04 menjaga urutan kejadian yang penting; kapan hasil lama diganti masih **BELUM DIPUTUSKAN**. |
| Koreksi khusus pengelola platform | EX-OPS tidak boleh berpura-pura bertindak sebagai organisasi, peninjau, penyelesai transaksi, atau verifikator independen. |

**Hasil pemeriksaan peta yang sudah disetujui:** kondisi penting pada skenario perwakilan mempunyai node, hubungan, atau batas **BELUM DIPUTUSKAN** yang terlihat. Ini adalah **kelengkapan makna peta untuk riset**, bukan kelengkapan seluruh kasus produksi atau bukti pengujian sistem nyata.

## 9. Pertanyaan Product yang Masih Terbuka (OPEN)

Daftar ini mencatat pertanyaan yang **belum memiliki keputusan Product yang disetujui**. Tujuannya agar jawaban tidak ditebak atau diwarisi otomatis dari aturan historis. **OPEN bukan tanda harus segera diputuskan, prioritas development, atau daftar pekerjaan wajib.** Suatu pertanyaan baru perlu diselesaikan ketika jawabannya benar-benar menentukan hasil Product yang sedang dipilih; jika tanpa jawaban itu hasil tidak bisa disetujui secara benar, pertanyaan tersebut menjadi *blocking* untuk lingkup pekerjaan itu saja.

1. **J-ORG:** siapa mendapat kewenangan `Owner` awal dan bagaimana kewenangan itu diakui; asal informasi; arti hasil J-ORG-01 yang berhasil, sebagian, atau disengketakan; arti hasil peninjauan organisasi yang positif/negatif; bahan minimum dan kewenangan `Owner`/`Staff`; akibat ke Campaign; serta kondisi peninjauan ulang.
2. **J-CAM:** arti dan akibat hasil kurasi; kapan perlu perbaikan, pengajuan ulang, peninjauan ulang, atau penggantian hasil sebelumnya; syarat penayangan publik dan kelayakan donasi; perubahan masing-masing keadaan dan akibatnya terhadap Pengunjung atau percobaan donasi.
3. **J-VIS:** informasi penting apa yang harus dijelaskan sebelum tindakan; arti percobaan yang boleh dilakukan dan fakta donasi yang diakui; status menunggu/gagal; perbedaan hubungan tamu dan pemilik akun; sumber/pemicu pendanaan atau settlement nyata; serta akibat perubahan status Campaign saat percobaan sedang berlangsung.
4. **J-ACC:** akibat penutupan Campaign; kewajiban minimum organisasi dan waktu pelaporan; siapa yang boleh mengakses laporan; kapan verifikasi berlaku, bagaimana hasilnya bermakna; kapan pertanggungjawaban dinilai selesai; serta penanganan bukti yang berubah.
5. **EX-OPS:** kewenangan positif pengelola platform dan jenis koreksi khusus yang diperbolehkan tanpa melanggar asal informasi, kewenangan pihak lain, atau kebenaran keadaan Product.

**Tidak semua pertanyaan OPEN perlu dijawab sekarang.** Selesaikan hanya yang material bagi hasil Product terpilih. Pertanyaan lain boleh tetap terbuka tanpa dianggap terlambat, selama tidak diam-diam diubah menjadi aturan atau klaim keberhasilan.

## 10. Hasil riset, batas persetujuan, dan kelanjutan

- **21 Oktober 2026:** Anhar menyetujui checkpoint riset dua tingkat dan peta Level 2 v0.2.
- **22 Oktober 2026:** Anhar menyetujui koreksi kewenangan pada peta Level 2 v0.3. Kode kelompok Journey sekarang `J-ORG`, `J-CAM`, `J-VIS`, `J-ACC`, `EX-OPS`, dengan padanan historis seperti bagian 1. Kode **26 node** tetap: J-ORG-01–J-ORG-05, J-CAM-01–J-CAM-08, J-VIS-01–J-VIS-07, J-ACC-01–J-ACC-06.
- Peta ini menggantikan **sketsa chat 19 node** untuk episode riset ini, **bukan** Product Intent yang sudah disetujui atau catatan keputusan historis yang tetap memiliki lingkupnya sendiri.
- Pemeriksaan diagram sebelumnya sempat menghasilkan ilustrasi **39 hubungan usulan** dan satu komponen yang saling terhubung secara tidak berarah. Itu hanya ilustrasi dalam sesi riset, **bukan bukti kausal yang divalidasi independen**, bukan bukti semua syarat development, dan bukan kepastian aturan Product sudah lengkap.
- **Hasil checkpoint 21 Oktober:** peta cukup untuk membahas pemilihan titik awal. **Pembaruan 22 Oktober:** J-ORG-01 merupakan **satu arah awal yang direkomendasikan untuk percobaan riset**, bukan prioritas delivery, kontrak Product yang disetujui, atau tugas Engineering.
- Persyaratan J-ORG-01 harus dirumuskan ulang dari Product Intent dan hubungan Journey yang berlaku. **Jangan mewarisi paket persyaratan/komitmen pilot lama** sebagai aturan Product yang mengikat untuk eksperimen baru ini.
- **Batas langkah berikutnya:** selesaikan keputusan Product yang benar-benar dibutuhkan oleh J-ORG-01, lalu susun kontrak Product mandiri. Jangan otomatis memilih node lain, memanggil Rancang, atau menganggap kesiapan Engineering telah terbukti. Lihat [keputusan riset tentang pemilihan titik awal dan penyerahan hasil Product](../outcome-selection.md).

**Batas persetujuan promosi Bahasa Indonesia:** Persetujuan hanya atas perubahan editorial dan kesetaraan makna peta riset. Ini bukan persetujuan baru atas aturan J-ORG-01, implementasi, atau perpindahan dari research branch ke Kencleng `main`.
