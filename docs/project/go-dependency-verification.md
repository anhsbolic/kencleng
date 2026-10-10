# Verifikasi dependency Go — outcome dan pemilihan metode

Panduan memilih metode pembuktian dependency tanpa menjadikan satu command
sebagai tujuan pekerjaan. Pin, persyaratan acceptance, batas write, dan stop
condition tetap berasal dari otoritas pekerjaan aktif. Untuk executable, cache,
network, scanner, dan recovery, gunakan [scanner runbook](go-scanner-runbook.md).

## Tujuan verifikasi

Persetujuan versi menetapkan dependency yang boleh dipilih; persetujuan itu belum
membuktikan graph hasil integrasi aman dan kompatibel. Verifikasi harus menjawab:

- Apakah graph aktual memilih versi yang disetujui, termasuk transitives?
- Apakah source yang diperiksa sesuai module/version/tag dan memiliki bukti
  origin serta checksum yang dipersyaratkan?
- Apakah affected ranges yang diketahui sudah tidak mencakup versi terpilih,
  dan apakah fresh source scan menemukan masalah lain?
- Apakah Go minimum, API, konfigurasi, dan jalur source yang relevan kompatibel
  dengan aplikasi?
- Apa yang sudah terbukti lewat source/integrity/compile, dan apa yang masih
  membutuhkan pengujian runtime?

Bug aplikasi adalah concern terpisah. Upgrade library tidak otomatis memperbaiki
kesalahan aplikasi seperti pemeriksaan expiry dengan waktu sebelum row lock.

## Pilih metode setelah memeriksa kebutuhan bukti

Jika otoritas aktif mengikat command tertentu sebagai acceptance requirement,
jalankan command itu. Jika yang diikat adalah hasil pembuktian, pilih metode
yang memenuhi hasil tersebut dengan akses dan perubahan minimum. Jangan
menganggap `go mod download -json all` wajib hanya karena dibutuhkan bukti seluruh
graph.

Mulai dari module inputs dan source saat ini. Jika perubahan dependency belum
boleh ditulis ke backend, gunakan probe terisolasi dengan source yang sama dan
pin yang disetujui. Jangan menyalin `.env` atau kredensial. Catat hash module
inputs, source yang disertakan, toolchain, cwd, dan konfigurasi cache. Pastikan
scanner juga memakai executable Go yang dimaksud; variabel toolchain saja tidak
menggantikan pemeriksaan executable/PATH aktual.

| Bukti yang diperlukan | Metode yang dapat digunakan | Batas pembuktian |
|---|---|---|
| Versi terpilih dan dependency edges | `go list -m -json all`, `go mod graph` | Edges dapat memuat requirements versi lebih rendah yang dikalahkan MVS; jangan menyamakannya dengan versi terpilih atau package yang digunakan runtime. |
| Go minimum | Field `GoVersion` dan `go.mod` module yang sesuai versi terpilih | Field yang tidak tersedia harus diperiksa; module legacy tanpa directive dicatat demikian, bukan diberi minimum hasil tebakan. |
| Source, tag/origin, dan checksum | Metadata/source cache versi tepat; download terarah jika ada yang kurang | Cache presence, upstream origin, dan integrity adalah bukti berbeda. |
| Integritas cache module | `go mod verify` dan pencocokan checksum yang dipersyaratkan | Tidak membuktikan source bebas vulnerability atau semua origin/tag sudah divalidasi. |
| Vulnerability source aplikasi | Fresh `govulncheck -json ./...` pada source dan graph yang dimaksud | Parse seluruh stream; exit 0 saja tidak cukup. Reachability bukan exploit reproduction. |
| Kecocokan API/build | Focused compile pada package terdampak | `go test -run '^$'` adalah compile check; jangan mengklaim behavioral tests dijalankan. |
| Kecocokan query/config/auth path | Source review, relevant diff, dan focused checks sesuai scope | Compile/source review tidak membuktikan koneksi, konfigurasi, atau contention DB nyata. |

### Periksa cache dahulu

Untuk setiap module/version terpilih, inventarisasi metadata `.info`, `go.mod`,
archive/source, dan checksum yang tersedia dalam module cache. Catat asal bukti
dan content identity yang diperlukan. Periksa `Origin` repo/ref/commit jika
tersedia; jangan menganggap metadata itu selalu ada pada module legacy.

Checksum archive/module (`h1`) dan SHA-256 file capture mempunyai fungsi berbeda.
SHA-256 membuat capture dapat diidentifikasi; ia bukan pengganti checksum module
Go atau validasi origin. Keberadaan source hasil ekstraksi juga belum membuktikan
tag-to-commit upstream. Jika provenance yang diwajibkan belum lengkap, ambil
bukti yang kurang atau laporkan gap; jangan menurunkan requirement.

### Ambil hanya bukti yang kurang

Jika cache belum cukup, gunakan versi tepat dari graph terpilih:

```sh
go mod download -json example.org/module@v1.2.3
```

Contoh di atas adalah placeholder, bukan pin proyek. Untuk setiap download,
periksa exit status dan structured `Error`; catat source, sums, dan origin yang
benar-benar tersedia. Tambahkan pemeriksaan upstream jika required origin/tag
tidak dapat dibuktikan oleh metadata itu.

`go mod download -json all` tetap boleh digunakan jika pengambilan seluruh module
set memang diperlukan. Namun, `all` dapat mencakup module yang bukan dependency
runtime langsung dan memerlukan akses jaringan/cache tambahan. Command tersebut
bukan source scan, source review, atau acceptance gate yang lengkap. Jangan
memaksakannya ketika bukti yang sama sudah tersedia dan tervalidasi, kecuali
otoritas aktif memang mewajibkan command tersebut.

Tidak ada izin untuk memilih versi lain, menghapus requirements, mengubah database
scanner, atau membersihkan cache global hanya demi membuat command sukses.

## Bedakan kendala metode dari hasil gate

| Hasil yang diamati | Makna | Tindakan |
|---|---|---|
| DNS/socket diblokir, cache tidak writable, input belum tersedia | Pemeriksaan belum memperoleh bukti yang dibutuhkan; bukan otomatis temuan vulnerability atau incompatibility. | Simpan error dan gunakan recovery/escalation yang diizinkan. Jangan menyatakan gate lolos. |
| Graph memilih affected version, fresh scan menghasilkan finding, source/API tidak kompatibel | Ada masalah substantif yang membutuhkan disposition sesuai otoritas aktif. | Stop pada batas yang ditetapkan; route Finding/Decision. Jangan waive/suppress atau mengganti pin sendiri. |
| Beberapa check lolos, bukti lain masih kurang | Acceptance belum lengkap. | Catat bukti yang lolos dan gap secara terpisah; jangan lanjut ke write yang mensyaratkan seluruh gate. |
| Seluruh bukti yang diwajibkan lengkap dan lolos | Gate yang ditetapkan dapat dinyatakan terpenuhi dalam scope bukti itu. | Lanjut hanya ke langkah yang sudah diotorisasi; jangan memperluas klaim ke runtime atau gate lain. |

Sebelum eksekusi, baca apakah stop condition aktif juga menutup Run pada kegagalan
tooling. Jika recovery diizinkan, kegagalan percobaan command dapat dipulihkan,
dicatat, lalu pemeriksaan diteruskan sampai bukti gate lengkap. Jika instruksi
eksplisit mengharuskan Run berhenti pada setiap kegagalan, patuhi batas tersebut;
recovery tidak otomatis membuka kembali Run terminal.

Untuk Invocation berikutnya, pemilik scope perlu menyatakan apakah kegagalan
lingkungan boleh ditangani dengan bounded retry/escalation di dalam Run, serta
kapan kegagalan yang tetap tidak terselesaikan harus dilaporkan sebagai blocker.
Pisahkan ketentuan itu dari stop-on-finding/incompatibility. Agent tidak boleh
mengubah stop condition secara sepihak melalui pemilihan metode.

## Baca hasil dan handoff sesuai outcome

Source scan harus selesai tanpa loading/network errors. Parse seluruh JSON:
objek `finding` adalah hasil yang perlu direkonsiliasi, sedangkan `osv` merupakan
metadata advisory. JSON hanya berisi `config` tidak membuktikan scan selesai.
Rekonsiliasi affected/fixed ranges dengan versi module yang benar, termasuk major
module path; jangan mengandalkan jumlah finding atau exit status saja.

Simpan command, cwd, executable/toolchain/scanner, konfigurasi cache/network dan
escalation, exit status, stdout/stderr, module/source identities, transitive
delta, disposition advisory, dan batas pengujian. Jika alternatif metode
digunakan, jelaskan bukti requirement mana yang dipenuhinya.

Probe yang lolos belum memverifikasi backend aktual setelah pin diterapkan.
Zero findings bukan final runtime/security acceptance. Handoff harus membedakan
metode berhasil, isolated gate lengkap, actual graph gate lengkap, dan perubahan
aplikasi beserta focused tests selesai.

Report terminal tetap dipertahankan. Diagnosis/percobaan yang diotorisasi setelah
Run berakhir dicatat terpisah dengan scope dan provenance sendiri. Bukti `/tmp`
perlu dipertahankan ke lokasi durable yang diizinkan sebelum dijadikan input
handoff; jangan menjadikan lokasi sementara sebagai satu-satunya bukti owning.

## Pelajaran percobaan D8 — 2026-10-11

T2-005 berhenti setelah download seluruh graph gagal karena sandbox memblokir
jaringan. Pengulangan melalui escalation berhasil. Percobaan terpisah kemudian
memeriksa cache yang sudah terisi, tanpa mengulang download massal: source dan
metadata 20 module tersedia, `go mod verify` lolos, focused compile lolos dengan
Go1.26.9, dan fresh source scan menghasilkan 0 objek `finding`.

Lima advisory lama muncul sebagai metadata database; pin D8 pengganti berada di
luar affected ranges-nya. Sebagian metadata module legacy tidak memiliki field
`Origin`, sehingga hasil tersebut tidak diklaim sebagai validasi provenance
seluruh graph atau full gate acceptance. Backend dan `LockSession` tidak berubah;
kompatibilitas runtime PostgreSQL/SCRAM dan koreksi aplikasi belum dibuktikan.

Cache lengkap pada percobaan itu terbantu oleh recovery download sebelumnya.
Pelajarannya adalah memeriksa gap bukti sebelum memilih command, bukan menganggap
cache selalu lengkap atau semua download tidak diperlukan. Versi dan hasil
percobaan adalah konteks diagnosis, bukan baseline dependency permanen atau
penutupan blocker. Evidence owning tetap berada pada report/capture pekerjaan;
dokumen ini mempromosikan panduan eksekusi yang dapat dipakai ulang.
