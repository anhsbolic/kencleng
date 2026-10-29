# Draf Pemberitahuan Privasi — Akun dan Pembentukan Organisasi

**DRAFT v0.1 — BELUM FINAL / BELUM UNTUK DIPUBLIKASIKAN.** Pemilik: Product/privacy Kencleng. Disiapkan 29 Oktober 2026 mengikuti [PD-PRIV-01-r4](privacy-decision.md). Pengendali yang ditetapkan Human: **Yayasan Kencleng Indonesia**. Human menyetujui arah pemetaan/pemberitahuan, bukan seluruh teks draf ini.

Teks ini menggambarkan perlakuan yang direncanakan untuk ORG-01 setelah keputusan terbuka diselesaikan dan mekanismenya diverifikasi. Bukan pernyataan bahwa fitur, penghapusan, backup atau kanal permintaan telah bekerja. Berlaku hanya untuk akun dan pembentukan konteks organisasi; fitur lain memerlukan informasi sesuai pemrosesannya.

## Informasi yang harus dilengkapi sebelum final

| Bagian | Keadaan aktual |
| --- | --- |
| Pengendali | Yayasan Kencleng Indonesia, diberikan Human |
| Kanal permintaan privasi dan cara menghubunginya | **care@kencleng.dev**, ditetapkan Human; penerimaan email/prosedur penanganan belum diuji |
| Login | Google Login dipilih; integrasi dan konfigurasi belum dibuktikan berjalan |
| Hosting, database, backup dan penerima/pembantu pemrosesan lain | **BELUM DITENTUKAN**; perlu daftar pihak, peran, data yang diproses dan lokasi pemrosesan |
| Dasar pemrosesan | Kandidat per tujuan disetujui untuk divalidasi; hasil validasi perjanjian/necessity dan kepentingan yang sah belum dicatat |
| Teks/tanggal berlaku final serta mekanisme hak pengguna | **BELUM DISETUJUI / BELUM DIVERIFIKASI** |

## Draf ringkasan sebelum Google login

Yayasan Kencleng Indonesia menggunakan pengenal dari Google untuk menyediakan akun Kencleng. Jika Anda membentuk organisasi, akun Anda dihubungkan sebagai Owner internal pertama pada konteks organisasi tersebut. Nama organisasi yang Anda isi diperlakukan sebagai informasi dari Anda.

Data digunakan untuk menyediakan layanan akun/organisasi, menjaga keamanan dan mendukung pemulihan dengan batas penyimpanan yang dijelaskan dalam pemberitahuan lengkap. Anda dapat mengajukan permintaan terkait data melalui **care@kencleng.dev**.

**Baca pemberitahuan lengkap sebelum melanjutkan.** Ringkasan ini merupakan informasi; bila suatu tujuan memerlukan consent, tindakan serta pencatatannya harus ditetapkan untuk tujuan itu. Tombol Google login tidak otomatis menjadi persetujuan atas semua tujuan pemrosesan.

## Draf pemberitahuan lengkap

### Pengendali dan kontak

Pengendali data untuk akun dan pembentukan organisasi Kencleng adalah **Yayasan Kencleng Indonesia**. Untuk pertanyaan atau permintaan terkait data Anda, hubungi **care@kencleng.dev**. Tanggal berlaku: **[ditetapkan setelah teks final dan mekanisme siap]**.

### Data dan tujuan penggunaannya

| Data yang direncanakan | Tujuan |
| --- | --- |
| Pengenal akun internal serta pengenal akun dari Google | Mengenali akun saat login dan menghubungkan tindakan dengan akun yang benar |
| Nama/ID organisasi, hubungan akun dengan Owner dan waktu pembentukan | Membentuk konteks organisasi dan memungkinkan pemeriksaan hubungan Owner kembali |
| Catatan penghubung permintaan pembentukan dengan hasilnya | Memeriksa hasil permintaan dan mencegah pengiriman ulang menghasilkan efek yang tidak semestinya |
| Catatan session dan riwayat refresh yang diperlukan | Mengelola akses, mencabut session dan mendeteksi penggunaan ulang token |
| Log keamanan/diagnostik minimum yang diperlukan | Menelusuri gangguan dan penyalahgunaan tanpa menyimpan nama organisasi, profil, token atau payload pribadi lengkap |
| Salinan backup serta penanda penghapusan minimum yang diperlukan | Pemulihan yang menghormati penghapusan data dan pencabutan akses sebelumnya |

ORG-01 tidak merencanakan pengumpulan email, foto, alamat atau dokumen identitas tambahan. Nama organisasi dapat mengandung data pribadi; gunakan informasi yang diperlukan untuk mengenali konteks yang Anda bentuk. Login Google tidak membuktikan kewenangan legal untuk mewakili organisasi. Hubungan Owner adalah kewenangan internal Kencleng.

### Dasar penggunaan data — harus divalidasi sebelum final

Untuk akun, pembentukan organisasi dan hubungan Owner, kandidat dasar adalah pemenuhan perjanjian atau permintaan pengguna sesuai hubungan layanan yang berlaku. Untuk keamanan/deteksi penyalahgunaan, kandidatnya adalah kepentingan yang sah setelah penilaian tujuan, kebutuhan dan keseimbangan terhadap hak pengguna. Backup mengikuti tujuan dan dasar data asal. **Bagian ini belum menjadi penetapan final bahwa seluruh syarat hukum telah terpenuhi.**

Jika pemrosesan tertentu memakai persetujuan, tujuan dan cara memberikan/menarik persetujuan harus dijelaskan serta pencatatannya disiapkan. Tidak ada tujuan tambahan yang diaktifkan oleh draf ini.

### Akses dan pihak yang membantu pemrosesan

Informasi organisasi dan hubungan Owner pada ORG-01 direncanakan tersedia melalui akses akun yang berwenang; pembentukan organisasi tidak otomatis membuat informasinya publik. Pengelolaan internal dibatasi pada kebutuhan serta kewenangan yang relevan.

Google digunakan untuk login. Data login yang diterima Kencleng direncanakan dibatasi pada pengenal yang diperlukan; email/foto/profil tambahan tidak diminta untuk ORG-01. Pemrosesan oleh Google perlu dijelaskan menurut layanan serta konfigurasinya yang benar-benar dipakai.

**[Lengkapi penyedia hosting, database, backup dan pihak lain: nama/peran, data yang ditangani, lokasi pemrosesan dan ketentuan relevan.]** Draf ini tidak menyatakan seluruh data disimpan di Indonesia atau menetapkan peran kontraktual Google/penyedia yang belum dinilai.

### Lama penyimpanan

Data utama akun dan organisasi disimpan selama diperlukan untuk tujuan layanan dan pemeriksaan hubungan Owner, dengan penanganan penghapusan sebagaimana dijelaskan berikut. Batas operasional yang direncanakan sesuai keputusan Product adalah:

| Kategori | Batas |
| --- | --- |
| Catatan hasil pembentukan yang berhasil | 48 jam sejak pembentukan berhasil |
| Session dan riwayat refresh yang diperlukan | Masa family maksimal 30 hari ditambah maksimal 24 jam pembersihan; lebih cepat bila session berakhir dan kebutuhan keamanannya selesai |
| Log keamanan/diagnostik yang diperlukan | 14 hari sejak kejadian |
| Backup, termasuk salinan/rangkaian pemulihan | Maksimal umur 14 hari |

Durasi penyimpanan session tidak memperpanjang aksesnya. Salinan data yang telah dihapus pada sistem aktif dibatasi penggunaannya untuk pemulihan sampai kedaluwarsa; keputusan penghapusan/pencabutan akses harus diterapkan kembali sebelum hasil restore dapat digunakan. Penanda penghapusan minimum diperlukan selama masih ada backup terdampak. **Pemenuhan jadwal dan prosedur tersebut belum dibuktikan pada implementasi.**

### Permintaan dan penghapusan akun

Anda dapat mengajukan pertanyaan, akses/koreksi, penghapusan atau permintaan terkait pemrosesan data melalui **care@kencleng.dev**. Penanganan mengikuti hak serta ketentuan yang berlaku; kebutuhan verifikasi dan hasilnya akan dijelaskan. Mekanisme serta batas waktu penanganan yang berlaku perlu dilengkapi sebelum final; tidak ada masa tunggu baru yang ditetapkan draf ini.

Permintaan penghapusan akun diverifikasi dan dampaknya dijelaskan sebelum tindakan. Jika Anda satu-satunya Owner pada organisasi yang baru memiliki konteks ORG-01, penutupan konteks organisasi memerlukan konfirmasi eksplisit agar tidak tertinggal aktif tanpa Owner. Penutupan di Kencleng tidak membubarkan organisasi secara hukum.

Jika organisasi hendak tetap berjalan atau terdapat pengguna/aktivitas/kewajiban tambahan, dampaknya memerlukan penanganan sesuai kasus. Tidak ada pengalihan Owner otomatis; permintaan privasi tetap ditangani. Keadaan permintaan diterima, akses dihentikan dan penghapusan selesai harus dibedakan. Informasi sisa salinan backup dan jadwalnya disampaikan sesuai keadaan yang benar-benar dapat dibuktikan.

### Perubahan pemberitahuan

Perubahan material pada data, tujuan, pihak pemroses, retensi atau penanganan permintaan harus dinilai dan dijelaskan sesuai ketentuan yang berlaku sebelum dipakai. Versi dan tanggal berlaku final perlu tersedia bagi pengguna. Dokumen ini masih draf internal dan belum menetapkan kanal penerbitan.

## Referensi untuk penulis dan reviewer

- [Keputusan Product PD-PRIV-01](privacy-decision.md) adalah owning source keputusan; tabel di sini adalah draf penjelasan untuk pengguna dan harus direkonsiliasi bila owning source berubah.
- [UU PDP, sumber resmi JDIH Kemkomdigi](https://jdih.komdigi.go.id/produk_hukum/view/id/832/t/undangundang%20nomor%2027%20tahun%202022): acuan evaluasi dasar, informasi dan hak pengguna; bukan bukti validasi operasi Kencleng.
- [Dokumentasi Google OpenID Connect](https://developers.google.com/identity/openid-connect/openid-connect): mekanisme scopes/claims dipilih sesuai kebutuhan; dokumentasi protokol tidak menetapkan seluruh pihak, lokasi atau dasar pemrosesan Kencleng.
- Tidak ada independent legal/security/solution reassessment, rendered verification, runtime test atau publikasi pada tahap penyusunan draf ini.
