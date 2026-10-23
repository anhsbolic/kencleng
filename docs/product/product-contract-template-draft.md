# Kencleng — Product Contract Template v0.1 (Draf)

> **Status:** USULAN STANDAR ARTEFAK PRODUCT / BELUM DISETUJUI. Template ini adalah rancangan untuk menghasilkan **Product-semantic input I1** bagi Rancang; bukan Product Contract yang telah disetujui, bukan persetujuan Space, dan bukan instruksi Engineering.
> **Pemilik keputusan:** Anhar (Product Kencleng). **Lingkup:** branch riset `pilot/station5-c1-space-only` saja; tidak mengubah Kencleng `main` atau Harscode.
> **Dasar Kencleng:** [Product Intent yang disetujui](product-intent.md), [Peta Journey Level 2](interconnected-journeys/level-2-network.md), dan [Product Output Contract v0.1 (working research)](interconnected-journeys/selection-and-handoff-research.md).
> **Dasar kontrak consumer (read-only reference):** Rancang `I1–I4`, [kontrak pada exact formed definition](https://github.com/anhsbolic/harscode-workspace/blob/4b49d0dd0db74cf34a927cadc938ae724676a75a/engineering/spaces/feature-solution-design/contracts.md); [formation approval v3](https://github.com/anhsbolic/harscode-workspace/blob/refactor/harscode-foundation/engineering/spaces/feature-solution-design/formation-approval-2026-10-21.md). Rancang formed **research-only** dan runtime belum terverifikasi.
> **Kaidah:** Satu kontrak menjelaskan **satu outcome Product yang utuh dan terbatas**. Struktur ini standar penulisan **usulan Kencleng**, bukan perubahan kontrak Rancang atau kewajiban file-layout. Pakai narasi ringkas dan tabel hanya jika benar-benar membantu; bagian yang tidak berlaku diberi `Tidak berlaku — [alasan]`, bukan diisi dengan dugaan.

## Informasi artefak dan authority

| Informasi | Isi |
| --- | --- |
| Proyek dan outcome | `Kencleng` / `<kode dan nama outcome>` |
| Product owner | `<Human pemilik keputusan>` |
| Status artefak | `DRAFT / PROPOSED / APPROVED` — pilih sesuai bukti keputusan yang benar-benar ada |
| Revision yang berlaku | `<path + commit/blob SHA>`; isi dari artefak nyata, bukan tebakan |
| Approval | `<pemberi keputusan, cakupan, tanggal, revision dan rujukan persetujuan>`; kosong/NOT APPROVED sebelum ada keputusan |
| Sumber Product yang berlaku | `<dokumen sumber + exact revision + bagian terkait>` |
| Cakupan penerapan | `<branch/proyek/outcome/syarat berlaku>` |
| Consumer yang dituju | `Rancang I1` untuk assignment yang **masih memerlukan Human-direct I4 dan pemeriksaan I2/I3** |

**Aturan authority:** Kesimpulan riset, sumber historis, hipotesis, dan keputusan yang masih OPEN **bukan** perilaku Product yang disetujui. Hanya artefak Product pada **exact revision yang disetujui** dapat diklaim sebagai I1 yang binding; tetap diperlukan cek kompatibilitas consumer untuk assignment nyata.

## 1. Outcome Product dan manfaatnya

Terangkan secara ringkas:
- **Situasi dan pengguna:** siapa yang membutuhkan hasil ini dan dalam kondisi apa.
- **Tindakan bermakna:** apa yang boleh dilakukan dan oleh siapa.
- **Hasil yang benar-benar berubah:** keadaan Product atau hubungan kewenangan apa yang nyata setelah berhasil.
- **Nilai bagi pengguna:** mengapa konsekuensi tersebut berarti.
- **Batas hasil:** hal terdekat yang **tidak** otomatis tercapai.

**Pernyataan outcome:** `<Dalam situasi ..., aktor ... dapat ..., sehingga ... benar-benar terjadi dan dapat dipahami/diperiksa ...>`

## 2. Pengalaman dan konsekuensi bagi pengguna

Ceritakan pengalaman yang perlu dipahami Rancang tanpa memutuskan urutan layar atau mekanisme implementasi.

| Momen / keadaan | Apa yang harus dapat dilakukan atau dipahami pengguna | Makna/akibat Product yang benar |
| --- | --- | --- |
| **Awal** | `<konteks, informasi yang diketahui, kondisi yang belum tentu ada>` | `<kebenaran keadaan awal>` |
| **Sebelum akibat penting berlaku** | `<konsekuensi, batas kewenangan atau informasi material yang harus dipahami>` | `<apa yang belum/akan berlaku>` |
| **Tindakan** | `<tindakan yang sah dan siapa yang berwenang>` | `<akibat tindakan bila terpenuhi syarat>` |
| **Berhasil dan setelahnya** | `<hasil yang terlihat dan bisa diperiksa kemudian>` | `<keadaan nyata yang harus bertahan; tanpa klaim berlebihan>` |
| **Gagal / tidak lengkap / belum pasti / ditolak** | `<perbedaan yang bermakna dari sukses>` | `<apa yang tidak boleh dinyatakan berhasil>` |

Hanya masukkan keadaan yang relevan. **Jangan** mengubah narasi menjadi urutan UI wajib, kewajiban login pada waktu tertentu, atau mekanisme teknis tanpa keputusan pemiliknya.

## 3. Perilaku Product, kewenangan, dan invariants

Tuliskan hasil keputusan Product yang **diterima** dengan ID stabil lokal artefak. Pembedaan sumbernya harus terlihat.

| ID | Perilaku atau aturan Product | Kewenangan/syarat | Hal yang tidak boleh disimpulkan (invariant) | Sumber/standing |
| --- | --- | --- | --- | --- |
| `<REQ-ID>` | `<apa yang harus benar, bukan cara teknis>` | `<aktor, hak, keadaan diperlukan>` | `<klaim/hasil terlarang>` | `<source exact revision; APPROVED atau PROPOSED>` |

Untuk versi **I1 approved**, tidak boleh ada aturan yang material bagi hasil terpilih tetapi masih hanya `PROPOSED`. Aturan tersebut harus disetujui atau dikeluarkan dari cakupan dengan konsekuensi yang jelas. Jangan menyulap keputusan Design/Engineering menjadi aturan bisnis.

## 4. Acceptance dan pemeriksaan batas makna

Definisikan cara mengenali hasil yang benar, tidak benar, dan belum pasti dari perspektif Product; **bukan** kasus API atau hasil test runtime yang belum dilakukan.

| ID | Rujukan aturan | Skenario / kondisi | Hasil Product yang boleh diklaim | Hasil yang tidak boleh diklaim |
| --- | --- | --- | --- | --- |
| `<AC-ID>` | `<REQ-ID>` | `<success/denied/partial/failed/unknown yang material>` | `<hasil teramati sesuai bukti>` | `<misleading success/privilege/verified claim>` |

Pemeriksaan perlu mencakup jalur negatif dan ketidakpastian yang **material untuk outcome ini**, bukan seluruh edge case Product. **Acceptance Product tidak membuktikan implementasi sudah lulus pengujian.**

## 5. Dependencies, cakupan, dan kebebasan solusi

| Perihal | Isi yang perlu jelas |
| --- | --- |
| **Syarat Product yang benar-benar diperlukan** | `<prerequisite dan apakah telah tersedia, tidak tersedia, atau belum terverifikasi>` |
| **Bukti/kondisi eksternal yang memengaruhi hasil** | `<sumber sinyal yang perlu ada; apa yang belum dibuktikan>` |
| **Dalam cakupan** | `<hasil dan batas perilaku yang ditanggung kontrak ini>` |
| **Di luar cakupan** | `<hasil lain yang sengaja tidak dijanjikan>` |
| **Kebebasan Design/Engineering** | `<pilihan interaksi, arsitektur, data, API, transaksi, dan mekanisme teknis yang tidak diatur Product>` |
| **Rujukan constraint teknis yang diketahui** | `<path/revision bila ada; jangan menyalin atau mengesahkan I2/I3 lewat I1>` |

Jangan menyebut prasyarat sudah tersedia hanya karena tercatat di diagram, mock, PRD, atau percobaan historis.

## 6. Pertanyaan terbuka dan blockers dalam cakupan

Catat hanya pertanyaan yang **material** untuk menilai kebenaran hasil ini; pertanyaan OPEN seluruh Product tetap berada di dokumen yang memilikinya.

| Pertanyaan | Pemilik keputusan | Dampak pada hasil/acceptance | Status dan tindakan |
| --- | --- | --- | --- |
| `<apa yang belum disetujui>` | `<Human/owner>` | `<bagian apa yang tidak dapat dinyatakan benar>` | `OPEN — blocking untuk klaim X / open non-blocking` |

**OPEN tidak otomatis mendesak.** Untuk menjadi **I1 binding**, pertanyaan Product yang blocking bagi outcome/acceptance dalam cakupan harus diselesaikan, atau cakupan diubah secara sah. Tidak boleh menyebut draft sebagai approved hanya karena formatnya lengkap.

## 7. Kesiapan konsumsi dan batas handoff

| Pemeriksaan | Standing/evidence |
| --- | --- |
| Seluruh aturan dalam cakupan mempunyai approval dan sumber exact revision | `<APPROVED / PROPOSED / BLOCKED + bukti>` |
| Outcome, narasi pengalaman, invariants, negative paths, dan acceptance konsisten | `<hasil rekonsiliasi yang benar-benar dilakukan>` |
| Tidak ada Product decision material yang tersembunyi atau belum disetujui | `<hasil / blocker>` |
| Syarat Product dan ketidakpastian ketersediaan dijelaskan | `<sumber/bukti dan batasnya>` |
| Rancang **I1 generic contract fit** | `<checked/conditional/not checked; alasan>` |
| Rancang **project-bound fit** pada invocation nyata | `<checked/conditional/blocked/not checked; consumer dan bukti>` |

**Pemisahan penting:** Template ini hanya menghasilkan calon **I1**. `I2` (Design/architecture/security), `I3` (repository/code/environment), dan `I4` (Human-direct assignment, akses/izin dan gates) **bukan bagian yang dipenuhi oleh Product Contract**. I1 yang disetujui juga **tidak otomatis** mengaktifkan Rancang atau Engineering.

**Kesimpulan saat ini:** `<DRAFT — NOT I1 READY / APPROVED PRODUCT I1 CANDIDATE, CONSUMER FIT UNVERIFIED / BLOCKED ...>`. Gunakan hanya status yang didukung bukti, bukan angka kelengkapan template.

## 8. Informasi Material Tambahan

**Fungsi:** Menangkap informasi **penting bagi pemahaman atau konsumsi outcome** yang belum tertampung secara tepat dalam section 1–7, tanpa memaksa section standar menjadi semakin banyak.

Isi hanya bila ada hal material dan jelaskan **mengapa tidak tertampung** pada bagian sebelumnya. Contoh jenis informasi yang *mungkin* relevan: batas interpretasi khusus domain, ketergantungan lintas-aktor yang tidak umum, asumsi kondisi lingkungan yang memengaruhi pembacaan Product, atau catatan pengalaman spesifik yang harus dipahami consumer.

Untuk tiap butir sebutkan:
- **Perihal dan alasan pentingnya** bagi outcome atau consumer.
- **Sumber/standing**: fakta disetujui, evidence, usulan, atau belum pasti.
- **Konsekuensi** jika diabaikan atau disalahartikan.
- **Pemilik tindak lanjut** jika butuh keputusan.

Jika tidak ada tambahan, tulis **`Tidak ada informasi material tambahan.`**

**Guardrail:** Bagian ini bukan tempat menaruh *blocking Product decision*, acceptance, invariant, constraint keamanan, atau kewajiban lain yang sebenarnya sudah punya section standar. Pindahkan hal seperti itu ke section yang tepat. Jangan membentuk authority baru dari catatan tambahan.

---

**Standing template ini:** `DRAFT / PROPOSED`. Persetujuan Anhar diperlukan sebelum menjadi standar penulisan Product Kencleng untuk episode riset ini. Sekalipun template disetujui, setiap Product Contract yang menggunakannya tetap memerlukan **persetujuan terpisah atas outcome, aturan, dan exact artifact revision**.
