# Kencleng — Peta Perjalanan Product

> **Status:** DRAF PENYEDERHANAAN BAHASA. Belum menggantikan [dokumen pengantar yang berlaku](README.md).
> **Sumber yang dibandingkan:** `README.md` pada branch `pilot/station5-c1-space-only`, blob `73c8e509429ea456b09d4d85a75de0d44bd88dfd`. Draf ini juga mengikuti [Product Intent](../product-intent.md) yang sudah disetujui dan [daftar istilah](../istilah.md).
> **Batas persetujuan:** Peta perjalanan Level 2 v0.3 sudah disetujui Anhar sebagai **hasil riset** pada 22 Oktober 2026. Penyuntingan pengantar ini **belum disetujui sebagai penggantinya**. Tidak mengubah aturan Product, memilih urutan development, atau mengizinkan pengerjaan.
> **Pemilik keputusan Product:** Anhar. Tidak berlaku sebagai sumber utama Kencleng `main`.

## 1. Mengapa peta ini dibuat

Perjalanan Product Kencleng saling terhubung: organisasi menyiapkan informasi, Campaign dapat ditinjau dan ditayangkan, Pengunjung mempertimbangkan donasi, dan organisasi mempertanggungjawabkan kegiatannya.

Peta ini membantu melihat hubungan tersebut, termasuk syarat, masalah, dan hal yang belum diputuskan. Tujuannya bukan menetapkan urutan development atau memecah Product menjadi daftar fitur.

Peta dibaca dalam **dua tingkat rincian**, bukan dua tahap yang harus dikerjakan berurutan.

## 2. Tingkat 1 — Gambaran perjalanan

| Kode | Perjalanan | Hasil yang ingin dipahami |
| --- | --- | --- |
| **R1** | Organisasi | Pendaftaran, kewenangan pengelola, informasi organisasi, dan peninjauan organisasi. |
| **R2** | Campaign | Pengajuan, kurasi, perubahan, penayangan, serta kelayakan menerima donasi. |
| **R3** | Pengunjung dan Donatur | Memahami Campaign, memilih berdonasi atau tidak, lalu memahami hasil donasi serta keadaan dana yang sebenarnya. |
| **R4** | Pelaporan dan pertanggungjawaban | Perubahan keadaan Campaign, laporan organisasi, pemeriksaan bila berlaku, dan informasi yang tetap dibutuhkan setelah penggalangan dana. |
| **X1** | Pengelola platform | Tindakan khusus yang memerlukan kewenangan tersendiri; tidak menggantikan keputusan organisasi, peninjau, atau bukti transaksi. |

Gambaran ini menjelaskan hubungan penting antarpengguna dan perjalanan, **bukan daftar pekerjaan yang telah disetujui**.

## 3. Tingkat 2 — Perubahan yang bermakna

[Peta Level 2 v0.3](level-2-network.md) memuat **26 bagian perjalanan**: R1 lima, R2 delapan, R3 tujuh, dan R4 enam. Tiap bagian menjelaskan perubahan yang bermakna bagi pengguna atau keadaan Product.

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
- Donasi yang telah dicatat **belum tentu** berarti dana telah diterima atau diselesaikan di luar sistem.
- Dana terkumpul **bukan** bukti penggunaan dana atau dampak nyata.
- Laporan organisasi **bukan** hasil pemeriksaan independen.
- Pertanggungjawaban tetap bermakna walaupun belum ada Donatur.
- Informasi yang belum pasti harus tetap ditampilkan sebagai belum pasti.

Peta juga mencatat penolakan, kegagalan, keadaan yang belum pasti, perubahan keputusan, dan hubungan antarpihak. Rinciannya ada di Level 2.

## 5. Sumber dan keputusan yang berlaku

- [Product Intent](../product-intent.md) adalah sumber tujuan dan prinsip Product yang sudah disetujui **pada branch riset ini**.
- [Level 2 v0.3](level-2-network.md) adalah peta riset yang disetujui. **Persetujuan peta bukan persetujuan aturan rinci setiap bagian.**
- [Keputusan riset tentang titik awal dan hasil Product](selection-and-handoff-research.md) mencatat alasan memilih arah percobaan **O1**, sekaligus batas keputusan yang masih terbuka.
- [Draf kebutuhan Product O1](o1-product-contract-draft.md) berisi **usulan**, belum menjadi aturan O1 yang disetujui atau masukan resmi bagi Rancang.

Dokumen `pilot-3-*` mempertahankan catatan keputusan dan proses lama. Catatan itu **tidak otomatis menjadi dasar aturan atau prioritas** untuk pekerjaan Product saat ini.

## 6. Hal yang belum diputuskan

Persetujuan atas peta ini **tidak** menetapkan:

- urutan development, cakupan rilis, atau daftar pekerjaan yang harus dimulai;
- seluruh aturan bisnis untuk organisasi, Campaign, donasi, dan pertanggungjawaban;
- rancangan teknis, kesiapan sistem yang berjalan, atau bukti keberhasilan development;
- pembentukan atau pelaksanaan Rancang;
- izin mengubah produk di branch operasional.

**Arah riset saat ini:** O1, yaitu pendaftaran organisasi dan kewenangan internal pertamanya, merupakan satu calon titik awal untuk pendalaman kebutuhan Product. Aturan tentang siapa yang boleh mendaftar, kapan dianggap berhasil, dan penanganan masalah **masih membutuhkan keputusan Product**.

Pemeriksaan kesiapan Engineering masih ditunda. Kemampuan Balad untuk memilih bagian berikutnya dan pembentukan Product Specification Space juga tetap ditunda. Tidak ada pekerjaan baru yang otomatis berjalan karena dokumen ini.

## 7. Riwayat persetujuan

- **21 Oktober 2026:** Anhar menyetujui peta dua tingkat dan Level 2 v0.2 sebagai hasil riset, bukan sebagai rencana development.
- **22 Oktober 2026:** Anhar menyetujui perbaikan Level 2 v0.3, terutama pemisahan dari aturan historis yang tidak otomatis berlaku pada O1.
- **22 Oktober 2026:** Draf Bahasa Indonesia untuk pengantar ini disiapkan sebagai penyuntingan saja. Perlu pemeriksaan makna sebelum menggantikan `README.md`.

Penyuntingan bahasa tidak boleh mengubah persetujuan atau membuka keputusan yang sengaja ditunda.
