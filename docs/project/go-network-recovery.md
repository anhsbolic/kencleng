# Go — akses jaringan dan recovery verifikasi dependency

Panduan operasional untuk dependency acquisition, checksum, provenance dan
source scan di lingkungan agent dengan jaringan/cache terbatas. Pin, scope
write dan acceptance tetap mengikuti pekerjaan aktif. Panduan ini tidak membuka
kembali Run terminal atau mengubah Invocation yang sudah mengikat stop condition.

## Mengapa cache masih dapat membutuhkan jaringan

Go memeriksa hash module terhadap `go.sum`; hash yang belum tersedia dapat
divalidasi melalui checksum database. Source cache lengkap belum menjamin bukti
checksum lengkap. Database menggunakan signed tree dan proof tiles; tile yang
belum tersimpan masih perlu diambil. `GOPROXY=off` mengatur pengambilan module,
bukan mematikan `GOSUMDB`. `go mod verify` memeriksa perubahan archive/source
terhadap hash cache; ia tidak menggantikan autentikasi download atau provenance.
Lihat [Go Modules Reference](https://go.dev/ref/mod#authenticating),
[checksum database](https://go.dev/ref/mod#checksum-database) dan
[`go mod verify`](https://go.dev/ref/mod#go-mod-verify).

Kebutuhan jaringan mengikuti bukti yang kurang:

| Kebutuhan | Tujuan yang mungkin diakses |
|---|---|
| Module/toolchain yang belum tersedia | Proxy yang dikonfigurasi, biasanya `proxy.golang.org`, atau upstream VCS |
| Autentikasi checksum | `sum.golang.org`, langsung atau melalui proxy yang mendukungnya |
| Origin/tag/commit/source upstream | Host repository dan archive sesuai module yang diperiksa |
| Fresh source vulnerability scan | Database scanner, biasanya `vuln.go.dev` |

Keberhasilan satu akses tidak membuktikan semua tujuan dapat dijangkau.

## Preflight di lingkungan eksekusi sebenarnya

1. Baca scope, pin, batas tracked write, dan ketentuan recovery/stop pekerjaan
   aktif sebelum menjalankan gate. Pemilik Invocation harus menetapkan apakah
   kegagalan percobaan karena lingkungan boleh dipulihkan dalam Run yang sama.
2. Catat executable dan `go version`; periksa konfigurasi berikut dari cwd
   backend atau probe yang memang akan diverifikasi:

   ```sh
   command -v go
   go version
   go env GOMOD GOTOOLCHAIN GOROOT GOPATH GOCACHE GOMODCACHE GOPROXY GOSUMDB
   ```

3. Periksa cache versi tepat terlebih dahulu. Pastikan lokasi yang perlu ditulis
   writable; gunakan build cache terisolasi di lokasi yang diizinkan bila perlu.
   Jangan menganggap cache home writable dalam sandbox.
4. Ambil hanya input/bukti yang kurang. Jika write backend belum diizinkan,
   lakukan resolusi pada probe terisolasi tanpa `.env` atau kredensial. Ingat
   bahwa `go mod tidy` dapat menulis `go.mod`/`go.sum`.
5. Jika command sudah diketahui memerlukan akses yang dibatasi sandbox, gunakan
   escalation harness untuk command itu dengan scope dan alasan konkret. Izin
   jaringan tetap mengikuti approval harness; dokumen ini bukan izin blanket.

## Recovery dan kapan melanjutkan

`socket: operation not permitted` atau kegagalan write cache menunjukkan kendala
lingkungan. Catat host/path yang gagal dan error lengkap. Jangan langsung
menyimpulkan layanan upstream down, dependency rentan atau API incompatible.

Untuk command penting yang gagal karena sandbox/network, gunakan pengulangan
melalui escalation sesuai aturan harness. Di Codex CLI, gunakan
`sandbox_permissions: require_escalated` dengan justification spesifik. Simpan
hasil awal dan pengulangan secara terpisah. Jangan menghindari approval melalui
tool lain, menghapus cache global, atau mengubah permission sistem.

Izin menjalankan recovery dan izin melanjutkan gate adalah dua hal berbeda:

| Ketentuan pekerjaan aktif | Tindakan setelah recovery |
|---|---|
| Recovery lingkungan diizinkan dalam Run yang sama | Setelah command berhasil, periksa output lalu lanjutkan gate yang masih kurang. Jangan nyatakan gate lolos dari retry saja. |
| Stop eksplisit pada kegagalan percobaan pertama, termasuk tooling | Tutup Run dan laporkan. Pengulangan yang diwajibkan harness hanya menjadi evidence recovery; jangan lanjutkan gate atau implementation. |
| Ketentuan recovery ambigu secara material | Route penjelasan ke pemilik Invocation sebelum dependent gate/write; jangan menetapkan sendiri makna stop condition. |
| Recovery tetap gagal atau bukti required masih tidak tersedia setelah batas recovery | Laporkan blocker pada scope yang tepat, dengan next-action owner. |

Checksum mismatch/security error bukan kegagalan akses biasa: stop dan route
temuan integritas. Finding vulnerability tanpa disposition yang disetujui,
incompatibility, input drift, provenance unsupported, atau pilihan dependency/
risk baru tetap mengikuti stop gate aktif. Jangan memakai `GOSUMDB=off`,
`GONOSUMDB=*`, atau pengecualian module publik sebagai workaround jaringan.

## Teks recovery untuk Invocation berikutnya

Pemilik Invocation dapat memakai teks berikut setelah menyesuaikan scope. Teks
ini merupakan template untuk dispatch berikutnya, bukan perubahan retroaktif
pada Run yang sudah terminal:

> Kegagalan percobaan command akibat akses jaringan/cache boleh dipulihkan dalam
> Run yang sama melalui koreksi lokasi cache writable yang diizinkan dan satu
> pengulangan dengan harness escalation per command yang terdampak. Catat setiap
> percobaan, izin, environment dan output. Setelah recovery berhasil, lanjutkan
> pemeriksaan gate; kegagalan lingkungan yang pulih tidak otomatis menutup Run.
> Berhenti jika recovery tetap gagal, approval ditolak, bukti required tetap
> tidak lengkap, terjadi checksum mismatch, input drift, provenance unsupported,
> vulnerability tanpa approved disposition, incompatibility, atau diperlukan
> pilihan dependency/risk/semantik baru. Jangan melakukan tracked write yang
> mensyaratkan gate sampai semua bukti required lengkap dan lolos. Jangan
> mengganti pin atau menonaktifkan verifikasi untuk mengatasi akses.

Batas satu pengulangan membuat recovery terukur. Jika perlu recovery tambahan,
pemilik scope menetapkan batas baru sebelum kelanjutan; jangan retry tanpa batas
atau membuat Run baru hanya untuk menghindari pertanyaan authority tersebut.

## Evidence dan batas klaim

Simpan command/cwd, executable/toolchain/scanner, konfigurasi cache/network yang
relevan, stdout/stderr, exit status, escalation/approval dan content identities.
Jangan merekam rahasia atau credential-bearing proxy URL. Pindahkan capture yang
diandalkan dari `/tmp` ke lokasi durable yang diotorisasi.

Bedakan recovery berhasil, gate isolated lengkap, gate backend aktual lengkap,
dan focused implementation checks selesai. Scan tetap harus fresh dan seluruh
JSON diperiksa; metadata scanner atau exit 0 saja tidak membuktikan clearance.
Report terminal tetap dipertahankan. Perubahan panduan ini tidak menutup blocker
atau mengesahkan hasil runtime.

Gunakan [panduan verifikasi dependency](go-dependency-verification.md) untuk
graph/source/provenance dan [scanner runbook](go-scanner-runbook.md) untuk
scanner/preflight/interpretasi JSON.

## Evidence lokal yang mendasari panduan

Pada T2-006, isolated `go mod tidy` gagal mengambil proof tile checksum akibat
sandbox socket restriction; pengulangan melalui escalation exit 0. Run berakhir
BLOCKED dan mencatat kelanjutan read-only setelah recovery sebagai penyimpangan
stop sequence. Backend tidak diubah. Lihat
[terminal report dan captures](../../.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/build-report.md).
Ini bukti kebutuhan akses dan kejelasan recovery, bukan acceptance D8 atau izin
melanjutkan Run tersebut.
