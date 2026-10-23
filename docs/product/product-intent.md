# Kencleng — Tujuan dan Prinsip Produk

> **Authority:** Product Intent **APPROVED** oleh Anhar, 22 Oktober 2026. Phase `PROD-FND`; output type `PROD-INTENT`. Sumber utama nilai, tujuan, aktor, tanggung jawab dan prinsip Product dalam branch `pilot/station5-c1-space-only`; **bukan** perubahan operational `main`.
> **Riwayat persetujuan:** Penulisan ulang Bahasa Indonesia yang disetujui Anhar setelah audit makna terhadap sumber Inggris sebelumnya, blob `433aa2376040a442d2aa14f56a01e3499dcab231`; draf akhir blob `5f0497a3d26eea0fe689af133529ea8fe1a648ff`. Versi Inggris tetap terlacak melalui Git history.
> **Cakupan:** Menjelaskan tujuan produk dan batas prinsip. Tidak menetapkan urutan development, aturan rinci O1, cakupan rilis, atau izin pelaksanaan Engineering.
> **Istilah:** Ikuti [acuan istilah Product Kencleng](istilah.md).

## 1. Tujuan Kencleng

Kencleng adalah layanan donasi dan penggalangan dana berbasis informasi dan bukti untuk organisasi nirlaba di Indonesia. Kencleng membantu orang memahami penggalangan dana sebelum memutuskan berdonasi, lalu mengikuti apa yang terjadi setelah mereka menyumbang.

Prinsip utamanya disebut **Evidence-Led Optimism**: harapan dibangun dari fakta yang dapat dipahami, asal informasi yang jelas, perkembangan yang disampaikan dengan jujur, dan keterbukaan tentang hal yang belum pasti. Bukan dari tekanan emosi atau klaim yang membuat sesuatu terkesan dapat dipercaya tanpa dasar.

Kencleng juga digunakan sebagai tempat belajar dan mencoba. Pembayaran atau proses operasional yang disimulasikan **harus disebut sebagai simulasi**. Simulasi tidak boleh ditampilkan sebagai penyelesaian transaksi uang yang benar-benar terjadi di luar sistem, hasil pemeriksaan independen, atau dampak nyata.

## 2. Pengguna dan tanggung jawab

Kencleng melayani Pengunjung, Donatur, perwakilan organisasi, peninjau, dan pengelola platform.

- Pengunjung dapat memahami Campaign yang terbuka untuk umum tanpa harus masuk ke akun.
- Pengunjung dapat berdonasi sebagai tamu **jika aturan donasi yang berlaku mengizinkannya**.
- Pengunjung yang sudah berdonasi disebut Donatur. Kapan donasi diakui tetap mengikuti aturan donasi yang berlaku.
- Organisasi dan Campaign adalah dua hal berbeda. Setiap Campaign berada di bawah satu organisasi.
- Peran `Owner` dan `Staff` pada organisasi memiliki kewenangan yang berbeda.
- Peninjauan organisasi, kurasi Campaign, dan verifikasi penggunaan dana merupakan keputusan yang **terpisah**.
- Peninjau tidak boleh melakukan penilaian jika memiliki conflict of interest terhadap pekerjaan atau pihak yang dinilai.

Rincian kewenangan dan cara kerjanya ditetapkan ketika kemampuan produk terkait dipilih untuk dikerjakan.

## 3. Perjalanan produk

Perjalanan Kencleng menghubungkan kegiatan berikut:

1. Pendaftaran dan peninjauan organisasi.
2. Penyiapan, kurasi, dan penayangan Campaign.
3. Pengunjung memahami Campaign dan memutuskan apakah akan berdonasi.
4. Pencatatan donasi, keadaan dana, dan perkembangan Campaign hingga penutupan.
5. Pelaporan dan pertanggungjawaban setelah penggalangan dana.

Urutan ini menunjukkan hubungan antarbagian produk, **bukan urutan wajib dalam development aplikasi**. Suatu bagian belum boleh dianggap sudah tersedia dan berjalan di produk hanya karena bagian berikutnya bisa diperagakan dengan data yang sudah disiapkan.

## 4. Aturan untuk menjaga integritas

- Bedakan fakta yang diketahui Kencleng, informasi dari organisasi atau penyelenggara Campaign, hasil peninjauan, keadaan sistem, laporan, dan informasi yang belum pasti atau masih menunggu kepastian.
- Perkembangan pendanaan bukan bukti bahwa kegiatan sudah dijalankan atau menghasilkan dampak.
- Sebelum pengguna melakukan tindakan penting, jelaskan akibatnya, makna dan keadaan uang, apakah tindakan tersebut diperbolehkan, serta perubahan keadaan yang akan terjadi.
- Hormati martabat, privasi, dan kebebasan pengguna dalam mengambil keputusan. Jangan memakai rasa bersalah, keadaan mendesak yang dibuat-buat, atau kesan dapat dipercaya tanpa bukti.
- Jika informasi belum lengkap atau belum tersedia, tampilkan apa adanya. Jangan menyatakan sesuatu sudah pasti jika belum ada dasar.
- Penayangan Campaign kepada publik dan kelayakannya menerima donasi harus mengikuti **keadaan Campaign yang benar-benar berlaku**.

## 5. Dasar desain

Prinsip desain produk, arah tampilan **Sunlit Editorial**, panduan visual, pola interaksi, dan aturan pengelolaan aset yang sudah disetujui tetap berlaku sesuai lingkupnya.

Panduan desain mengatur cara informasi dan interaksi ditampilkan, **bukan** menetapkan aturan bisnis atau prioritas rilis.

Rincian pengalaman pengguna dan proses bisnis ditentukan setelah hasil Product yang ingin dicapai dan pemilik keputusannya jelas.

## 6. Hal yang belum diputuskan

Dokumen ini **belum menetapkan**:

- tujuan MVP atau rilis;
- cakupan pekerjaan yang dipilih dan urutannya;
- kriteria penerimaan hasil untuk tiap pekerjaan;
- jalan pintas operasional;
- tata cara peninjauan yang terperinci;
- saluran pemberitahuan;
- kontrak antarsistem atau API;
- rancangan teknis dan rencana pelaksanaan.

Setiap keputusan tersebut memerlukan pertimbangan yang jelas dan persetujuan pihak yang berwenang sebelum pekerjaan terkait dimulai.
