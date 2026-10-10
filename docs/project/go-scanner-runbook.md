# Go scanner — local setup and recovery

Panduan menjalankan `govulncheck` dan menangani kegagalan tooling. Gunakan
toolchain dan dependency yang ditetapkan oleh otoritas pekerjaan aktif. Dokumen
ini tidak menetapkan pin baru, menerima risiko, atau mengubah gate suatu Run.

## Pemeriksaan sebelum scan

Jalankan dari `backend/`. Periksa executable yang benar-benar dipakai, terutama
setelah berpindah versi melalui GVM:

```sh
command -v go
go version
go env GOTOOLCHAIN GOROOT GOPATH GOMOD GOCACHE GOMODCACHE
command -v govulncheck
go version -m "$(command -v govulncheck)"
```

Pastikan `GOMOD` menunjuk modul backend, versi Go cocok dengan otoritas aktif,
build cache bisa ditulis, dan module cache tersedia. Download module/toolchain
memerlukan izin menulis module cache dan akses jaringan. Source scan memerlukan
akses database vulnerability, secara default `https://vuln.go.dev`.

GVM memisahkan tools dalam pkgset tiap versi Go. Jika scanner tidak tersedia
pada pkgset aktif, pasang versi scanner yang dipilih secara eksplisit dengan
`go install golang.org/x/vuln/cmd/govulncheck@<versi>`. Catat versi dan build
metadata; jangan menganggap tools pada pkgset lama otomatis ikut diperbarui.

`.env`, koneksi PostgreSQL, dan kredensial Google tidak diperlukan oleh source
scan backend saat ini. Jangan membaca atau merekam rahasia untuk diagnosis ini.

## Eksekusi dan pencatatan

Untuk pemeriksaan interaktif:

```sh
govulncheck -show traces ./...
```

Untuk evidence terstruktur, buat direktori capture unik. Contoh berikut memakai
`/tmp`; pindahkan hasil yang perlu dipertahankan ke lokasi evidence yang diizinkan
oleh pekerjaan aktif sebelum direktori sementara hilang.

```sh
scanner_capture_dir=$(mktemp -d /tmp/kencleng-govulncheck.XXXXXX)
scanner_exit=0
govulncheck -json ./... > "$scanner_capture_dir/scan.json" \
  2> "$scanner_capture_dir/scan.stderr" || scanner_exit=$?
printf 'exit=%s\ncapture=%s\n' "$scanner_exit" "$scanner_capture_dir"
```

Dalam sandbox agent, direktori home dapat bersifat read-only meskipun bisa
ditulis dari terminal lokal. Jika perlu cache terisolasi, arahkan `GOCACHE` hanya
untuk perintah terkait ke direktori writable yang unik:

```sh
GOCACHE="$scanner_capture_dir/build-cache" govulncheck -json ./... \
  > "$scanner_capture_dir/isolated-scan.json" \
  2> "$scanner_capture_dir/isolated-scan.stderr"
```

Cache terisolasi tidak memberikan akses jaringan atau izin module cache. Jika
perintah gagal karena pembatasan sandbox/cache/network, gunakan mekanisme
escalation yang disediakan harness dengan persetujuan yang diperlukan. Jangan
mengubah permission sistem atau mengganti database untuk melewati pembatasan.
Jalankan compile dan scan secara berurutan saat diagnosis agar hasil mudah
dibandingkan. Compile terfokus tidak menggantikan full source scan.

## Jika gagal lagi

| Gejala | Pemeriksaan / tindakan |
|---|---|
| `cache entry not found`, `reading srcfiles list`, `loading compiled Go files from cache` | Simpan error lengkap. Periksa cache writable dalam lingkungan eksekusi yang sama; coba cache unik hanya jika kelanjutan diagnosis diizinkan. Jangan langsung menyimpulkan cache rusak atau versi Go salah. |
| `read-only file system`, `permission denied`, gagal membuat module cache | Identifikasi cache yang disebut error. Gunakan lokasi writable yang diizinkan atau escalation; jangan melakukan `chmod`/penghapusan cache global sebagai perbaikan default. |
| DNS/socket `operation not permitted`, fetching database gagal | Periksa izin jaringan. Dalam sandbox, ulangi melalui escalation sesuai aturan harness; dalam terminal biasa, periksa konektivitas/proxy. Scan belum lengkap. |
| Package compile/load error tetap terjadi dengan cache writable | Periksa error source, Go aktual, scanner build metadata, dan kompatibilitas loader melalui dokumentasi resmi. Perubahan versi harus mengikuti otoritas aktif; simpan hasil sebelum/sesudah. |
| Scan selesai dan menghasilkan findings | Catat ID, affected/fixed range, module/package/symbol trace dan applicability. Ikuti gate dependency/risk aktif; jangan suppress atau menerima risiko sebagai solusi tooling. |

Jika perlu membedakan package loading dari akses database, gunakan
`go list -compiled -export ./...` pada lingkungan/cache yang sama. Jika memilih
`go list -e -json`, periksa juga `Error` dan `DepsErrors`: flag `-e` dapat
menghasilkan exit 0 sambil membawa package errors.

Retry hanya untuk diagnosis lingkungan yang diizinkan dan beralasan. Catat
setiap percobaan serta hasilnya. Jika Invocation/Human mewajibkan stop saat gate
gagal, hentikan Run tersebut dan laporkan; lakukan diagnosis lanjutan hanya
dengan otorisasi terpisah. Jangan menulis ulang report terminal atau menganggap
retry otomatis membuka Run yang sudah berhenti.

## Menilai hasil

- Source scan selesai tanpa loading/network errors dan hasilnya diperiksa.
  JSON yang hanya berisi `config` tidak membuktikan scan selesai.
- **Exit 0 dalam mode JSON tidak berarti bebas vulnerability.** Parse seluruh
  stream JSON dan periksa objek `finding`; objek `osv` adalah metadata advisory,
  bukan otomatis temuan pada graph yang dipindai.
- Simpan command, cwd, exit status, stdout/stderr, toolchain/scanner metadata,
  identitas module inputs dan batas verifikasi. Hindari rahasia dalam evidence.
- Pisahkan scanner berfungsi, scan tanpa findings, dan seluruh dependency gate
  lolos. Scan probe tidak membuktikan graph backend aktual bersih. Independent
  runtime testing tetap mengikuti otoritas pekerjaan aktif.

## Pelajaran dari diagnosis 2026-10-11

Diagnosis setelah Run T2-004 mengonfirmasi bahwa build cache home tidak bisa
ditulis dalam sandbox; package loading menghasilkan `read-only file system`.
Dengan cache baru yang writable, scanner v1.7.0 berhasil memuat package tetapi
akses database diblokir oleh izin socket. Setelah escalation, full scan selesai.
Scanner versi yang sama kemudian dipasang pada pkgset GVM Go1.26.9.

Probe approved D8 menghasilkan 0 findings. Backend aktual masih menghasilkan
lima ID yang sudah dikenal: GO-2026-4771/4772/5004/5970/6629. Ini membuktikan
tooling bekerja, bukan dependency acceptance atau completion T2. Tidak ada bukti
bahwa upgrade versi scanner diperlukan untuk menyelesaikan kegagalan tersebut.
Versi dan hasil ini adalah catatan diagnosis, bukan baseline permanen.

## Rujukan

- [Dokumentasi resmi govulncheck: build configuration, output dan exit codes](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck).
- [Go toolchain selection](https://go.dev/doc/toolchain).
- [GVM: pemilihan versi dan pkgset](https://github.com/moovweb/gvm).
