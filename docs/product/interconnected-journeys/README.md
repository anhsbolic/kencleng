# Kencleng — Interconnected Product Journeys

> **Authority:** Product Journey introduction approved by Anhar 22 Oktober 2026. Output kind: `PROD-JOURNEYS`. Applies to branch `pilot/station5-c1-space-only`, not operational `main`. Identifier changes are traceable through [migration record](../identifier-migration-2026-10-23.md).
> **Riwayat:** Versi Inggris sebelum penyesuaian kode: blob `73c8e509429ea456b09d4d85a75de0d44bd88dfd`; pengantar Inggris berkode baru: blob `a4a70735663acabefea6e69ed7468487285a3b4e`; draf Bahasa Indonesia yang disetujui: blob `f56a091e60755995e3a986a364489f509e5e2c87`. Riwayat tersedia melalui Git.
> **Batas:** Persetujuan ini mencakup pengantar dan penulisan ulang yang telah diaudit, bukan aturan rinci J-ORG-01, rencana development, Rancang, atau izin menjalankan Engineering. [Level 2 v0.3](level-2-network.md) tetap sumber rincian peta yang sudah disetujui; [Product Intent](../product-intent.md) tetap sumber prinsip produk.
> **Istilah:** Gunakan [daftar istilah Product Kencleng](../istilah.md).

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

Kode seperti `J-ORG-01`, `J-CAM-01`, `J-VIS-01`, dan `J-ACC-01` hanya penanda pada peta. **Bukan** nomor pekerjaan, urutan development, atau bukti bahwa sebuah kemampuan sudah tersedia.

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
- [Keputusan riset tentang titik awal dan hasil Product](selection-and-handoff-research.md) mencatat alasan memilih arah percobaan **J-ORG-01**, sekaligus batas keputusan yang masih terbuka.
- [Draf kebutuhan Product J-ORG-01](o1-product-contract-draft.md) berisi **usulan**, belum menjadi aturan J-ORG-01 yang disetujui atau masukan resmi bagi Rancang.
- [Rekonstruksi nilai bisnis](../pilot-3-business-value-loop.md), [hasil riset pengguna dan perannya](../pilot-3-actor-outcomes.md), dan [contoh skenario riset](../pilot-3-stage-3a-representative-scenarios.md) adalah **riwayat bahan pembanding** yang mendasari pekerjaan riset sebelumnya. Persetujuan atau aturan lama tidak otomatis menjadi keputusan Product saat ini.

Dokumen `pilot-3-*` mempertahankan catatan keputusan dan proses lama. Catatan itu **tidak otomatis menjadi dasar aturan atau prioritas** untuk pekerjaan Product saat ini.

## 7. Hal yang belum diputuskan

Persetujuan atas peta ini **tidak** menetapkan:

- urutan development, cakupan rilis, atau daftar pekerjaan yang harus dimulai;
- seluruh aturan bisnis untuk organisasi, Campaign, donasi, dan pertanggungjawaban;
- rancangan teknis, kesiapan sistem yang berjalan, atau bukti keberhasilan development;
- pembentukan atau pelaksanaan Rancang;
- izin mengubah produk di branch operasional.

**Arah riset saat ini:** J-ORG-01, yaitu pendaftaran organisasi dan kewenangan internal pertamanya, merupakan satu calon titik awal untuk pendalaman kebutuhan Product. Aturan tentang siapa yang boleh mendaftar, kapan dianggap berhasil, dan penanganan masalah **masih membutuhkan keputusan Product**.

Pemeriksaan kesiapan Engineering masih ditunda. Kemampuan Balad untuk memilih bagian berikutnya dan pembentukan Product Specification Space juga tetap ditunda. Tidak ada pekerjaan baru yang otomatis berjalan karena dokumen ini.

## 8. Riwayat persetujuan

- **21 Oktober 2026:** Anhar menyetujui peta dua tingkat dan Level 2 v0.2 sebagai hasil riset, bukan sebagai rencana development.
- **22 Oktober 2026:** Anhar menyetujui perbaikan Level 2 v0.3, terutama pemisahan dari aturan historis yang tidak otomatis berlaku pada J-ORG-01.
- **22 Oktober 2026:** Anhar menyetujui pengantar Journey Bahasa Indonesia yang sudah diperbaiki dan diaudit. Draf tersebut dipromosikan ke `README.md` ini; riwayat bahasa Inggris tetap terlacak melalui Git.
- **22 Oktober 2026:** Anhar juga menyetujui dan mempromosikan [Level 2 v0.3 Bahasa Indonesia](level-2-network.md) setelah audit kesetaraan makna. Peta Level 2 Inggris sebelumnya tetap terlacak melalui Git; kode, hubungan, dan keputusan OPEN tidak diubah.

Penyuntingan bahasa tidak boleh mengubah persetujuan atau membuka keputusan yang sengaja ditunda.


**Current selected outcome:** `ORG-01` corresponds to Journey node `J-ORG-01`. Its [Product Contract](../outcomes/org-01/product-contract.md) and [approval](../outcomes/org-01/approval.md) are current scoped Product authority. The Journey map does not automatically approve any other node.
