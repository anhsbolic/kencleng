# Kencleng — Panduan Copywriting UX Bahasa Indonesia

> **Status:** Panduan kerja tingkat proyek (working guidance), untuk diuji lewat desain dan review Human. Melengkapi [Product Design Principles](product-design-principles.md), [UX Patterns](patterns.md), dan [Information Architecture](information-architecture.md), **bukan** menggantikan Product Contract, pemberitahuan hukum yang sudah disetujui, atau istilah resmi yang ditetapkan pemiliknya.
> **Tujuan:** Copy UI yang alami bagi pembaca Indonesia, ringkas tanpa kehilangan makna, hangat tanpa basa-basi, dan jujur terhadap kondisi sistem.
> **Bahasa utama UI Kencleng:** Bahasa Indonesia; istilah teknis/produk dipakai bila memang membantu akurasi.

## 1. Tulis untuk tindakan nyata, bukan terjemahan kalimat Inggris

Mulai dari situasi pembaca: **“Saya ada di mana, perlu melakukan apa, apa akibatnya, dan apa langkah berikutnya?”** Susun kalimat Bahasa Indonesia dari makna tersebut, bukan dengan menerjemahkan kata per kata dari contract, API error, atau English UX copy.

Suara Kencleng: **jelas, manusiawi, tenang, dapat dipercaya, tidak menggurui**. Hormati keseriusan identitas, donasi dan penghapusan; tidak perlu seruan dramatis, klaim “aman/terverifikasi” yang tidak didukung, jargon legal/Engineering, atau kelucuan pada failure state.

Utamakan kata kerja dan istilah yang lazim:
- **Masuk**, bukan “melakukan proses autentikasi” atau “sign in” bila tidak perlu.
- **Buat organisasi**, bukan “inisiasi proses pembentukan entitas” (tetapi gunakan hanya pada tombol yang **benar-benar** mengirim creation request).
- **Periksa hasil permintaan**, bukan “resolve unknown mutation”.
- **Tinjau** untuk pemeriksaan sebelum tindakan; **konfirmasi** untuk persetujuan atas konsekuensi; **selesai** hanya jika outcome memang terbukti.

Hindari pengulangan kata “konteks”, “pemrosesan”, “relasi”, “atribusi”, “efek”, “berhasil”, dan “internal” ketika tidak dibutuhkan oleh kalimat. **Jangan hilangkan kualifikasi tersebut bila menentukan batas kebenaran.**

## 2. Prioritaskan kalimat yang mudah dipindai

- **Judul:** sampaikan tujuan atau keadaan, bukan slogan abstrak untuk tugas operasional. Editorial headlines boleh ekspresif pada public surfaces, tetapi harus segera diikuti orientasi yang konkret.
- **Lead/satu kalimat utama:** apa yang perlu diketahui untuk keputusan saat ini. Jangan memasukkan seluruh alur berikutnya.
- **Label/tombol:** sebut tindakan dengan verba yang jelas; bedakan **Masuk**, **Tinjau**, **Buat**, **Periksa hasil**, **Batalkan**, **Hapus akun** menurut actual effects.
- **Helper text:** tambahkan hanya informasi yang membantu orang menyelesaikan field/action sekarang.
- **Detail tambahan:** simpan pada link “Baca informasi privasi”, “Apa arti Owner?”, atau halaman/expandable content yang sesuai. Bukan berarti semua fakta boleh disembunyikan.

Panjang copy mengikuti risiko dan kebutuhan pembaca, bukan target jumlah kata. Kurangi **biaya membaca yang tidak relevan**, jangan memangkas konsekuensi yang wajib dipahami.

## 3. Contoh editing yang lebih natural

Contoh di bawah adalah **candidate copy**, bukan teks Product/Legal yang telah disetujui. Makna harus diverifikasi terhadap current scoped authority sebelum dipakai.

| Terlalu teknis / berat | Alternatif untuk diuji | Alasan |
| --- | --- | --- |
| “Memulai inisiasi konteks organisasi.” | “Mulai buat organisasi” | Menyebut tujuan langsung; cocok sebagai orientasi, **bukan** klaim bahwa login otomatis membuat organisasi |
| “Lakukan autentikasi menggunakan penyedia identitas Google.” | “Masuk dengan Google” | CTA singkat dan natural |
| “Masuk bukanlah pembentukan relasi Owner internal.” | “Masuk belum membuat organisasi.” | Menjawab kemungkinan salah paham pada langkah **login**, tidak membebankan detail creation lebih awal |
| “Lakukan pemeriksaan atribusi dan kewenangan owner pertama.” | “Tinjau nama dan akun sebelum membuat organisasi.” | Berorientasi pada hal yang akan diperiksa; jelaskan Owner internal di review step |
| “Operasi tidak memberikan kepastian terhadap efek yang dihasilkan.” | “Hasil permintaan belum bisa dipastikan.” | Membedakan unknown dari gagal, tanpa jargon |
| “Melakukan inspeksi atas request sebelumnya.” | “Periksa hasil permintaan” | Memakai bahasa tugas yang mudah dipahami |
| “Konfirmasi terminasi seluruh konteks dan penghapusan identitas.” | “Tutup organisasi terkait dan hapus akun” | Lebih konkret, **hanya** bila seluruh scope dan efeknya benar menurut accepted deletion contract |
| “Sistem mengalami error saat melakukan fetch.” | “Daftar organisasi belum bisa dimuat.” | Menyebut dampak bagi pembaca, bukan transport detail |

Jangan mengadopsi contoh di atas mentah-mentah. Misalnya, **“Owner”** merupakan istilah kewenangan internal Kencleng yang tidak sama dengan pemilik hukum; bila diterjemahkan menjadi “pemilik organisasi” tanpa kualifikasi, maknanya bisa salah.

## 4. Istilah, sapaan, dan konsistensi

- Gunakan **Anda** bila perlu menyapa pembaca atau menunjuk tanggung jawab pribadi secara jelas. Tidak perlu mengulang “Anda” di setiap kalimat.
- Gunakan **kami** hanya bila speaker dan organisasi yang bertanggung jawab jelas. Jangan mewakili regulator, Google, atau organisasi yang dibuat user secara implisit.
- Gunakan **akun**, **masuk**, **organisasi**, **nama organisasi**, **Owner pertama di Kencleng** (dengan penjelasan internal yang relevan), dan **Reviewer** sesuai scoped authority. Periksa apakah istilah yang lebih ramah tetap mempertahankan distinction hukum/role.
- Pilih satu istilah per konsep dalam satu flow. Jangan berganti antara “permintaan”, “proses”, “aksi”, “transaksi” tanpa alasan makna.
- Nama proper, alamat kontak, batas waktu, angka, dan status legal/produk harus berasal dari authority yang benar; jangan dipoles sehingga berubah menjadi janji baru.
- Waspadai format UI: tanggal/jam, mata uang, angka, pluralitas, serta screen-reader text harus dapat dibaca alami oleh pengguna Indonesia.

## 5. Error, loading, dan uncertainty

Pola yang disarankan: **apa yang terjadi / apa yang belum diketahui → apakah tindakan memiliki akibat → langkah aman berikutnya**, sepanjang relevan.

**Login gagal sebelum request pembentukan:** gunakan kalimat yang terikat pada fase itu, misalnya “Belum bisa masuk. Coba lagi atau kembali ke halaman awal.” Jangan memunculkan peringatan tentang mutation yang belum pernah dilakukan.

**Hasil pembuatan tidak pasti:** “Hasil pembuatan organisasi belum bisa dipastikan. Periksa hasil permintaan ini sebelum mencoba membuat yang baru.” Jangan mengatakan **“Gagal membuat organisasi”** hanya karena timeout/connection loss. Jangan menyuruh mengirim request baru jika recovery/idempotency belum dipastikan.

**Validasi form:** sebut field dan perbaikan konkret, misalnya “Masukkan nama organisasi.” atau “Nama organisasi tidak boleh lebih dari 200 karakter.” Gunakan batas sesuai accepted contract, bukan batas yang diasumsikan.

**Penghapusan:** bedakan “pratinjau siap”, “permintaan diterima”, “akses berakhir”, “data aktif dihapus”, dan “salinan lain belum terverifikasi” hanya jika backend contract/observed evidence mendukung masing-masing. Jangan menyebut “Semua data Anda sudah hilang” tanpa bukti semua-copy disposal.

**Loading:** gunakan tugas yang sedang berjalan (“Memeriksa sesi…”, “Memuat organisasi…”), bukan slogan atau “Harap menunggu” berulang.

Saat copy berisiko mengubah business/legal/security meaning, bawa kalimatnya ke owner terkait; jangan diam-diam membuat keputusan baru.

## 6. Disclosure dan privasi: jujur tanpa membuat halaman seperti dokumen hukum

Pisahkan **ringkasan singkat sebelum tindakan**, **akses jelas ke pemberitahuan lengkap**, dan **konsekuensi khusus tepat sebelum commit**. Gunakan [IA guidance](information-architecture.md) untuk memilih inline vs link/page vs accordion/dialog. Jangan menyembunyikan hal yang wajib diketahui di tooltip hover-only, dan jangan mengganti teks legal authoritative dengan ringkasan rekaan yang berpotensi bertentangan.

Untuk ORG-01, bila full notice masih berstatus **DRAFT**, labeli demikian dan jangan menyamarkannya sebagai notice final atau “persetujuan menyeluruh”. Jangan mengklaim kanal privasi sudah operasional hanya karena alamat sudah ditentukan.

## 7. Visual support adalah bagian dari copy decision

Kalau untuk menjelaskan sesuatu diperlukan tiga paragraf, tanya:
1. Adakah informasi berulang atau sebenarnya baru relevan pada langkah berikutnya?
2. Bisakah satu diagram proses, ikon kontekstual, atau ilustrasi bermakna mengomunikasikan **hubungan** yang lebih cepat dipahami?
3. Apa label/caption minimum agar visual tidak menambah ambiguitas?
4. Fakta apa tetap wajib ditulis jelas karena visual tidak cukup presisi?

Lihat [Asset Governance](asset-governance.md); foto dan ilustrasi bukan bukti real-world identity, legitimacy, donation settlement atau outcome.

## 8. Review copy sebelum Design acceptance

Gunakan pemeriksaan singkat, bukan checklist panjang untuk setiap microcopy:

- **Natural:** Apakah ini terasa ditulis orang Indonesia, bukan terjemahan mesin atau bahasa kontrak?
- **Relevant:** Apakah kalimat membantu keputusan **sekarang**, bukan memindahkan semua policy ke layar?
- **Truthful:** Dapatkah tiap klaim dan status dipertanggungjawabkan dari current authority/evidence?
- **Actionable:** Setelah membaca, tahu apa yang bisa dilakukan dan apa akibatnya?
- **Consistent:** Label dan istilah sesuai flow; kata kerja CTA tidak melebih-lebihkan efek?
- **Accessible:** Bisa dipindai, dibacakan, dilihat di mobile, dan dimengerti tanpa bergantung pada visual saja?

Jika Human menilai copy tidak natural, lakukan **rewrite berbasis maksud dan konteks**, bukan mengganti beberapa sinonim. Untuk keputusan copy yang material, cantumkan representative final candidate copy di feature-specific Design Handoff—jangan jadikan contoh di panduan ini sebagai satu-satunya copy source of truth.
