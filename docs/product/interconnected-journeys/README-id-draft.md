# Kencleng — Peta Journey Product

> **Status:** DRAF PENYEDERHANAAN BAHASA. Belum menggantikan [dokumen pengantar yang berlaku](README.md).
> **Sumber yang dibandingkan:** pengantar lama `README.md`, blob `73c8e509429ea456b09d4d85a75de0d44bd88dfd`, sebelum penyesuaian kode. [Pengantar yang berlaku](README.md) dan [Level 2](level-2-network.md) kini memakai kode Journey baru. Draf ini mengikuti [Product Intent](../product-intent.md) dan [daftar istilah](../istilah.md).
> **Batas persetujuan:** Peta Journey Level 2 v0.3 sudah disetujui Anhar sebagai **hasil riset** pada 22 Oktober 2026. Penyuntingan pengantar ini **belum disetujui sebagai penggantinya**. Tidak mengubah aturan Product, memilih urutan development, atau mengizinkan pengerjaan.
> **Pemilik keputusan Product:** Anhar. Tidak berlaku sebagai sumber utama Kencleng `main`.
> **Penamaan kode (disetujui 22 Oktober 2026):** `J-ORG`, `J-CAM`, `J-VIS`, `J-ACC`, dan `EX-OPS` menggantikan kode kelompok sebelumnya; aturan Product serta 26 kode bagian Level 2 tidak berubah.

## 1. Mengapa peta ini dibuat

Journey Product Kencleng saling terhubung: organisasi menyediakan informasi, Campaign dapat dikurasi bila berlaku dan ditayangkan ketika memenuhi aturan, Pengunjung mempertimbangkan donasi, dan organisasi mempertanggungjawabkan kegiatannya.

Peta ini membantu melihat hubungan tersebut, termasuk syarat, masalah, dan hal yang belum diputuskan. Tujuannya bukan menetapkan urutan development atau memecah Product menjadi daftar fitur.

Peta dibaca dalam **dua tingkat rincian**, bukan dua tahap yang harus dikerjakan berurutan.

## 2. Tingkat 1 — Gambaran Journey

| Kode | Journey | Hasil yang ingin dipahami |
| --- | --- | --- |
| **J-ORG** | Organisasi | Pembentukan konteks organisasi atau penggunaan konteks yang sudah ada, kewenangan `Owner` dan `Staff`, informasi organisasi, serta peninjauan organisasi bila berlaku. Pendaftaran organisasi baru tidak wajib untuk setiap tindakan berikutnya. |
| **J-CAM** | Campaign | Penyiapan usulan Campaign oleh perwakilan yang berwenang, kurasi bila berlaku, perbaikan informasi, penayangan, dan kelayakan menerima donasi. Penyiapan usulan tidak otomatis mewajibkan proses pengajuan atau kurasi formal. |
| **J-VIS** | Pengunjung dan Donatur | Memahami Campaign dan memilih berdonasi, menunda, atau tidak berdonasi. Jika ada percobaan donasi, hasilnya harus jujur: belum pasti, gagal, atau diakui sebagai donasi jika memang ada bukti yang memenuhi aturan. Status Donatur dan keadaan dana tetap dibedakan. |
| **J-ACC** | Keadaan Campaign dan pertanggungjawaban | Perubahan keadaan Campaign, pelaporan, dan pertanggungjawaban **selama Campaign berlangsung maupun setelah ditutup**, termasuk ketika belum ada Donatur. Peninjauan laporan dilakukan jika memang berlaku. |
| **EX-OPS** | Pengelola platform | Tindakan khusus yang memerlukan kewenangan tersendiri; tidak menggantikan keputusan organisasi, peninjau, atau bukti transaksi. |

Gambaran ini menjelaskan hubungan penting antarpengguna dan antar-Journey, **bukan daftar pekerjaan yang telah disetujui**.

**Padanan kode lama:** `R1 → J-ORG`, `R2 → J-CAM`, `R3 → J-VIS`, `R4 → J-ACC`, dan `X1 → EX-OPS`. Kode `EX-OPS` adalah pengecualian lintas-Journey, bukan Journey kelima. Ini hanya perubahan nama kelompok.

## 3. Tingkat 2 — Perubahan yang bermakna

[Peta Level 2 v0.3](level-2-network.md) memuat **26 bagian Journey**: J-ORG lima, J-CAM delapan, J-VIS tujuh, dan J-ACC enam. Tiap bagian menjelaskan perubahan yang bermakna bagi pengguna atau keadaan Product.

Dalam membaca hubungan antarbagian, bedakan:

- **Syarat yang harus ada:** Tanpa keadaan atau informasi tertentu, suatu hasil belum dapat dinyatakan benar. Adanya syarat tersebut saja belum tentu cukup untuk menyatakan berhasil.
- **Hubungan bersyarat:** Hanya berlaku ketika kondisi yang ditetapkan terpenuhi.
- **Kelanjutan:** Satu tindakan dapat menimbulkan kebutuhan atau tanggung jawab berikutnya, tetapi belum tentu menjadi pekerjaan development berikutnya.
- **Perubahan dan peninjauan kembali:** Informasi baru dapat memicu perbaikan atau penilaian ulang bila diperlukan.
- **Belum diputuskan:** Aturan yang masih membutuhkan keputusan Product. Jangan menebak hasilnya.

Kode seperti `O1`, `P1`, `V1`, dan `A1` hanya penanda pada peta. **Bukan** nomor pekerjaan, urutan development, atau bukti bahwa sebuah kemampuan sudah tersedia.

## 4. Makna yang tidak boleh dicampur

Beberapa hal berikut harus tetap dibedakan:

- Peninjauan organisasi **bukan** kurasi Campaign.
- Campaign yang ditayangkan kepada publik **belum tentu** boleh menerima donasi.
- **Percobaan donasi bukan bukti donasi berhasil.** Donasi yang diakui harus mempunyai dasar sesuai aturan Product; percobaan yang tertunda atau gagal tidak menciptakan status Donatur.
- Donasi yang diakui, keadaan pendanaan Campaign, dan penyelesaian transaksi uang adalah **tiga hal berbeda**. Yang satu tidak otomatis membuktikan dua hal lainnya. Simulasi juga tidak boleh diakui sebagai penyelesaian transaksi nyata.
- Dana terkumpul **bukan** bukti penggunaan dana atau dampak nyata.
- Laporan organisasi **bukan** hasil pemeriksaan independen.
- Pertanggungjawaban tetap bermakna walaupun belum ada Donatur.
- Informasi yang belum pasti harus tetap ditampilkan sebagai belum pasti.

Peta juga mencatat penolakan, kegagalan, keputusan untuk tidak berdonasi, keadaan yang belum pasti, perubahan keputusan, serta hubungan antarpihak. Rinciannya ada di Level 2.

## 5. Hasil dan batas pemeriksaan peta

**Hasil yang telah diterima Anhar:** Peta dua tingkat dan rincian 26 bagian sudah dinilai **cukup untuk membahas pemilihan titik awal**. Pemeriksaan terhadap hubungan antarbagiannya menunjukkan syarat, hubungan bersyarat, peninjauan kembali, penolakan, dan ketidakpastian penting sudah terlihat atau ditandai sebagai **belum diputuskan**.

**Yang belum dibuktikan:** Kelengkapan seluruh aturan bisnis Product, kesiapan sistem dan Engineering, kemampuan pembaca baru memahami dan menggunakan peta tanpa bantuan, serta apakah cara memilih pekerjaan berdasarkan peta ini benar-benar lebih baik. Jumlah hubungan pada diagram bukan bukti bahwa seluruh ketergantungan dapat dijalankan.

Keputusan Product yang belum dibuat harus tetap terlihat sebagai **belum diputuskan**, bukan diisi dengan dugaan.

## 6. Sumber dan keputusan yang berlaku

- [Product Intent](../product-intent.md) adalah sumber tujuan dan prinsip Product yang sudah disetujui **pada branch riset ini**.
- [Level 2 v0.3](level-2-network.md) adalah peta riset yang disetujui. **Persetujuan peta bukan persetujuan aturan rinci setiap bagian.**
- [Keputusan riset tentang titik awal dan hasil Product](selection-and-handoff-research.md) mencatat alasan memilih arah percobaan **O1**, sekaligus batas keputusan yang masih terbuka.
- [Draf kebutuhan Product O1](o1-product-contract-draft.md) berisi **usulan**, belum menjadi aturan O1 yang disetujui atau masukan resmi bagi Rancang.
- [Rekonstruksi nilai bisnis](../pilot-3-business-value-loop.md), [hasil riset pengguna dan perannya](../pilot-3-actor-outcomes.md), dan [contoh skenario riset](../pilot-3-stage-3a-representative-scenarios.md) adalah **riwayat bahan pembanding** yang mendasari pekerjaan riset sebelumnya. Persetujuan atau aturan lama tidak otomatis menjadi keputusan Product saat ini.

Dokumen `pilot-3-*` mempertahankan catatan keputusan dan proses lama. Catatan itu **tidak otomatis menjadi dasar aturan atau prioritas** untuk pekerjaan Product saat ini.

## 7. Hal yang belum diputuskan

Persetujuan atas peta ini **tidak** menetapkan:

- urutan development, cakupan rilis, atau daftar pekerjaan yang harus dimulai;
- seluruh aturan bisnis untuk organisasi, Campaign, donasi, dan pertanggungjawaban;
- rancangan teknis, kesiapan sistem yang berjalan, atau bukti keberhasilan development;
- pembentukan atau pelaksanaan Rancang;
- izin mengubah produk di branch operasional.

**Arah riset saat ini:** O1, yaitu pendaftaran organisasi dan kewenangan internal pertamanya, merupakan satu calon titik awal untuk pendalaman kebutuhan Product. Aturan tentang siapa yang boleh mendaftar, kapan dianggap berhasil, dan penanganan masalah **masih membutuhkan keputusan Product**.

Pemeriksaan kesiapan Engineering masih ditunda. Kemampuan Balad untuk memilih bagian berikutnya dan pembentukan Product Specification Space juga tetap ditunda. Tidak ada pekerjaan baru yang otomatis berjalan karena dokumen ini.

## 8. Riwayat persetujuan

- **21 Oktober 2026:** Anhar menyetujui peta dua tingkat dan Level 2 v0.2 sebagai hasil riset, bukan sebagai rencana development.
- **22 Oktober 2026:** Anhar menyetujui perbaikan Level 2 v0.3, terutama pemisahan dari aturan historis yang tidak otomatis berlaku pada O1.
- **22 Oktober 2026:** Draf Bahasa Indonesia untuk pengantar ini dibuat dan diperbaiki setelah audit kesesuaian makna. Tetap memerlukan persetujuan tersendiri sebelum menggantikan `README.md`.

Penyuntingan bahasa tidak boleh mengubah persetujuan atau membuka keputusan yang sengaja ditunda.
