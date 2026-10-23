# Kencleng — Penerapan Product Contract Template (Uji Coba)

> **Status:** CATATAN UJI COBA / BUKAN TEMPLATE UTAMA DAN BELUM APPROVED SEBAGAI KONTRAK O1. Kepemilikan **draf template reusable** dipindahkan ke concern Product Harscode. Dokumen ini hanya mencatat penerapan dan evaluasinya untuk Kencleng.
> **Template sumber (proposal Harscode):** [Product Contract Template v0.1 — Bahasa Indonesia](https://github.com/anhsbolic/harscode-workspace/blob/e24b8869e290033e05a52c92dcb8bf3cdc49d172/product/product-contract-template-id-draft.md) dan [versi English](https://github.com/anhsbolic/harscode-workspace/blob/refactor/harscode-foundation/product/product-contract-template-draft.md). Statusnya **research proposal**, belum operational `main` dan belum mewajibkan format universal.
> **Riwayat:** Draf template lokal Kencleng sebelumnya tersimpan di Git sebagai blob `28b2e56cb58f3068bfcef1b34df59e99dfe56f1e`. Jangan menganggapnya salinan kedua yang harus disinkronkan manual.
> **Scope:** Hanya `anhsbolic/kencleng` branch riset `pilot/station5-c1-space-only`. Kencleng Product Authority tetap milik [Product Intent](product-intent.md), [Peta Journey Level 2](interconnected-journeys/level-2-network.md), dan keputusan Anhar pada scope terkait.

## Penerapan untuk Kencleng O1

- **Calon artefak uji:** [O1 Product Contract v0.2 — draf Bahasa Indonesia](interconnected-journeys/o1-product-contract-draft.md) menggunakan template Indonesia Harscode. Masih `PROPOSED`, belum menjadi input Rancang `I1` yang binding. Jangan memperlakukan contoh/aturan yang diusulkan sebagai approved.
- **Struktur reusable:** metadata authority; tujuh section standar; satu section **Informasi Material Tambahan** yang kondisional. `Success Criteria` eksplisit di bagian Outcome; `Acceptance Criteria` harus memiliki tautan ke `Success Criteria` yang relevan.
- **Pilihan istilah:** Gunakan [acuan istilah Product Kencleng](istilah.md) dan Bahasa Indonesia untuk dokumen Product proyek ini. Bentuk bahasa lokal **bukan perubahan semantik template Harscode**.
- **Batas consumer:** Product Contract hanya menangani `I1`. `I2` milik Design/Engineering constraints, `I3` kode/lingkungan, dan `I4` assignment/otorisasi langsung dari Anhar; template tidak memenuhi ketiganya dan tidak mengaktifkan Rancang.
- **Hal yang harus dibuktikan:** Apakah Product Contract O1 v0.2 dalam Bahasa Indonesia membantu pembaca baru memahami outcome, kewenangan, Success Criteria, perilaku, negative paths, Acceptance Criteria, blockers, dan kebebasan solusi **tanpa chat history** atau mengarang keputusan Product. Periksa juga kecocokan semantik dua versi template.
- **Standing evaluasi:** Belum ada bukti uji fresh-reader/consumer Rancang yang terpisah. Jangan mencatat `PASS` atau menyebut template Harscode sudah terbukti universal hanya karena strukturnya lengkap.

## Batas keputusan

Persetujuan Anhar sebelumnya untuk **memasukkan Success Criteria ke struktur template** telah diterapkan pada draf Harscode. Persetujuan itu **tidak otomatis** menyetujui seluruh template sebagai kebijakan Harscode operasional, tidak menyetujui PD-O1-01..05, dan tidak menyetujui artifact revision sebagai Rancang `I1`.

**Langkah uji berikutnya:** Review dan putuskan hanya keputusan Product O1 yang benar-benar material dalam [draf O1 v0.2](interconnected-journeys/o1-product-contract-draft.md), lalu uji kecukupan output sebagai calon `I1` melalui pembaca/consumer terpisah pada exact revision. Hasil evaluasi dapat menjadi evidence untuk keputusan standar Harscode selanjutnya.
