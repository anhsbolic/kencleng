# Object storage Kencleng: SeaweedFS

Keputusan engineering untuk persiapan yang diminta Human, 24 Oktober 2026:
**SeaweedFS disiapkan sebagai capability S3 lokal opsional**, dengan arah evaluasi
deployment menggunakan produk yang sama. Ini tidak menambahkan fitur upload
produk atau mengubah scope delivery. Verifikasi lokal dilaporkan berhasil oleh
Human; kesiapan produksi masih memerlukan bukti tersendiri. Konfigurasi tersedia
di `docker-compose.storage.yml`.

## Perbandingan dan batas kompatibilitas

| Kandidat | Temuan dari sumber resmi | Penilaian untuk Kencleng |
| --- | --- | --- |
| SeaweedFS | Mendokumentasikan bucket policy, CORS, multipart, versioning dan presigned URL. Memiliki panduan single-node dan cluster. Beberapa API belum tersedia atau berupa stub; semantik direktori berbeda dari S3. | Pilihan untuk menyiapkan core S3 dan jalur deployment sekarang. Operasi yang digunakan aplikasi harus diuji terhadap versi yang dipin. |
| RustFS | Menyediakan core S3, console, dan cakupan pengujian kompatibilitas. Interoperabilitas format disk MinIO masih preview, memerlukan build khusus; objek terenkripsi MinIO tidak dapat dibaca build default. | Alternatif bila kemiripan operasi MinIO menjadi prioritas. Klaim vendor tidak menetapkan pemenang universal untuk kompatibilitas/kesiapan. |
| Garage | Mendokumentasikan core S3 dan presigned URL; tabel resminya mencatat keterbatasan seperti bucket versioning. | Layak bila cakupan S3 tersebut sesuai kebutuhan; bukan pilihan utama untuk meniru cakupan MinIO secara luas. |

Kompatibilitas API S3, operasi administrasi MinIO (`mc admin`, IAM/config), dan
format volume adalah hal berbeda. Tidak ada bukti bahwa ketiganya menggantikan
seluruh perilaku MinIO. Backend saat ini belum memiliki client object storage;
integrasi berikutnya sebaiknya memakai operasi S3 standar dengan konfigurasi
endpoint/region/path-style agar tidak bergantung pada API administrasi vendor.
Volume/data MinIO tidak pernah dipasang ke SeaweedFS; migrasi jika diperlukan
melalui copy objek S3 dan verifikasi tersendiri.

Sumber primer yang diperiksa:

- [SeaweedFS S3 API dan perbedaan semantik](https://github.com/seaweedfs/seaweedfs/wiki/Amazon-S3-API)
- [SeaweedFS quick start](https://github.com/seaweedfs/seaweedfs#quick-start)
- [SeaweedFS production setup](https://github.com/seaweedfs/seaweedfs/wiki/Production-Setup)
- [RustFS compatibility matrix](https://github.com/rustfs/rustfs/blob/main/docs/architecture/s3-compatibility-matrix.md)
- [RustFS feature status dan batas on-disk MinIO](https://github.com/rustfs/rustfs#feature--status)
- [Garage S3 compatibility](https://garagehq.deuxfleurs.fr/documentation/reference-manual/s3-compatibility/)

## Setup dan run lokal

Dari root Kencleng, pilih engine secara eksplisit:

```bash
# Podman:
make COMPOSE=podman-compose storage-up
make COMPOSE=podman-compose storage-check

# Mengulangi pemeriksaan sambil me-restart HANYA SeaweedFS untuk bukti persistensi:
make COMPOSE=podman-compose storage-persistence-check

# Docker Engine/Desktop: hilangkan COMPOSE=podman-compose pada perintah di atas.
```

`storage-up` memastikan PostgreSQL/network siap, memulai satu SeaweedFS, menunggu
request S3 terautentikasi berhasil, memastikan bucket, lalu menerapkan public-read
policy hanya pada objek `kencleng-public` dan CORS dua origin FE lokal. Tidak ada
anonymous write atau anonymous listing yang diberikan. `kencleng-private` memakai
akses terautentikasi/presigned URL. Pemeriksaan gagal jika private object dapat
dibaca tanpa autentikasi. Tidak ada client S3 yang ditambahkan pada backend.

| Pemakai | Endpoint |
| --- | --- |
| Browser / BE host | `http://127.0.0.1:9087` |
| BE container / AWS CLI container | `http://seaweedfs:8333` |
| Region / addressing | `us-east-1` / path-style |
| Bucket | `kencleng-public`, `kencleng-private` |
| Access key dev | `kencleng-local` |
| Secret key dev | `kencleng-local-development-only` |

Kredensial ini khusus lokal. Untuk override, ekspor `STORAGE_ACCESS_KEY` dan
`STORAGE_SECRET_KEY` pada shell Compose sebelum startup dan gunakan nilai yang
sama pada setiap init/check. Jangan menganggap mengganti env akan merotasi
kredensial yang tersimpan oleh SeaweedFS; lakukan rotasi terencana melalui
mekanisme yang didukung versi tersebut. `.env.example` mencatat usulan `S3_*`
untuk integrasi client berikutnya; `backend/.env` yang sudah ada dipertahankan.

Presign menggunakan endpoint browser sejak awal. Mengganti hostname/port setelah
URL ditandatangani membatalkan signature. Alat check menghasilkan URL eksternal
dan mengaksesnya dari host, sementara operasi upload/delete memakai DNS internal.

Storage memiliki project Compose `kencleng-storage`, named volume baru
`kencleng-storage_seaweedfs_data`, dan memakai network eksternal database
`kencleng_default`. Override `KENCLENG_INFRA_NETWORK` bila nama jaringan memang
berbeda. PostgreSQL bukan penyimpan metadata SeaweedFS pada konfigurasi ini;
volume `/data` mencakup metadata lokal dan isi objek. Lifecycle file Compose
storage terpisah agar perubahan konfigurasinya tidak merecreate database.

Port S3 hanya dipublikasikan di loopback host. WebDAV, Admin UI, Iceberg dan Lance
dinonaktifkan. Komponen internal mini masih berbagi namespace jaringan container;
konfigurasi ini mengasumsikan jaringan development lokal yang dipercaya. Untuk
deployment, komponen internal harus diisolasi dari aplikasi/client umum.

```bash
make COMPOSE=podman-compose storage-status
make COMPOSE=podman-compose storage-logs
make COMPOSE=podman-compose storage-down
```

`storage-down` melepas container tanpa menghapus volume. Root `make down`
melepas app container, SeaweedFS, lalu PostgreSQL tanpa menghapus data. Hentikan
MinIO lama secara eksplisit bila masih memakai port 9087; helper ini tidak
menghapus atau memigrasikannya. Konfigurasi MinIO lama tetap tersedia untuk
kebutuhan legacy dan tidak dipakai oleh target `storage-*`.

## Versi yang dipin dan standing verifikasi

- SeaweedFS **4.48**, digest index multi-platform
  `sha256:4e61d15fd35994cb1e43e1e553dff106794841fd9a99ade2fc8c8bfce4d7872d`.
  [Metadata publisher](https://hub.docker.com/v2/repositories/chrislusf/seaweedfs/tags/4.48).
  Saat diperiksa, release source 4.49 tersedia tetapi tag image standar 4.49
  belum tersedia; dipilih image 4.48 yang tersedia untuk amd64/arm64.
- AWS CLI **2.38.5**, digest index
  `sha256:c8156a0e74e13929a42d7e60f0d4e14e7d28d980cd21822505d4bff80e5460c1`.
  [Metadata publisher](https://hub.docker.com/v2/repositories/amazon/aws-cli/tags/2.38.5).
- Entrypoint/flag diperiksa terhadap source tag 4.48. Pin mencegah penggunaan
  berikutnya berubah mengikuti `latest`; ini belum merupakan audit image.
- Pemeriksaan statis dan tes helper memakai simulasi terpisah dari bukti runtime.
  Sandbox agent tidak dapat menjalankan provider Podman; agent tidak mengklaim
  menjalankan sendiri startup atau smoke test aktual.
- Pembaruan Human, 24 Oktober 2026: setelah diberi perintah Podman `storage-up`,
  `storage-check`, dan `storage-persistence-check`, Human menyampaikan
  "udah bro. ga ada error". Status lokal: **lulus menurut laporan Human** untuk
  setup, pemeriksaan S3 dan persistensi setelah restart. Output/log tidak
  dilampirkan. Ini tidak membuktikan Docker Engine/Desktop, seluruh kompatibilitas
  S3, integrasi aplikasi, recovery backup atau kesiapan deployment produksi.

`storage-check` memakai AWS CLI resmi untuk autentikasi/signing, membuat objek
dengan key UUID di prefix `kencleng-setup-probe/`, lalu menghapus hanya objek
milik percobaan itu. Ia memeriksa upload, public-read, private denial, anonymous
write denial, presigned GET dan expiry, serta CORS preflight. Ia tidak membuktikan
presigned PUT, multipart, versioning, object lock atau seluruh kompatibilitas S3.
`storage-persistence-check` menambahkan restart SeaweedFS di tengah pemeriksaan;
jalankan saat tidak ada upload aktif. Error cleanup dilaporkan beserta key yang
perlu dibersihkan; jangan menghapus bucket/volume sebagai pemulihan otomatis.

## Sebelum deployment

Produk ini memiliki jalur deployment, tetapi konfigurasi lokal ini belum menjadi
deployment produksi yang terverifikasi. Tuntaskan untuk target yang dipilih:

1. Tentukan single-server atau kebutuhan HA/RPO/RTO; single-node satu disk tidak
   tahan kehilangan host/disk. Cluster membutuhkan topologi master, volume dan
   filer beserta metadata store yang sesuai panduan produksi.
2. Sediakan volume persisten, kapasitas, monitoring serta backup **data dan
   metadata**, lalu buktikan restore ke instance baru. Restart bukan tes backup.
3. Gunakan kredensial deployment terpisah, akses aplikasi terbatas pada bucket
   yang dibutuhkan, dan isolasi API filer/master/volume/admin.
4. Atur DNS/TLS endpoint publik, penerusan Host/query yang benar, dan CORS origin
   deployment. Buktikan presigned GET/PUT dari browser serta penolakan akses
   private/expired/invalid signature.
5. Uji operasi S3 yang digunakan aplikasi, termasuk multipart bila diperlukan,
   terhadap versi image/SDK yang dipin.
6. Buktikan upgrade/rollback, recovery kegagalan dan kapasitas pada host target;
   evaluasi patch/security release sebelum mempromosikan image lokal.

Pemilihan/penyediaan server, perubahan firewall publik, upload data nyata, atau
deployment produksi belum dilakukan dalam persiapan lokal ini.
