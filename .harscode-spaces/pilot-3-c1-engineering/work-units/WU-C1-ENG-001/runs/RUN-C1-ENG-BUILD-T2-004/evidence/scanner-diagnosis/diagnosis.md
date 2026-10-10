# Diagnosis scanner Go lokal — 2026-10-11

Isu package loading/build cache sudah diselesaikan untuk eksekusi dengan cache writable dan akses database vulnerability. Backend/dependency dan terminal Run T2-004 tidak diubah; ini diagnosis tooling terpisah.

## Penyebab dan bukti

Default `GOCACHE=/home/anhar-solehudin/.cache/go-build` berada di luar writable roots sandbox sesi. `go list -e -compiled -export -json ./...` dalam sandbox mengembalikan package Error `pattern ./...: open /home/anhar-solehudin/.cache/go-build/aa/aaf18d0bf224d01ffe80ecc9afdb1301ac6a97f20fa918bc8c93bd9bd7ce19fb-d: read-only file system`. Karena memakai `-e`, process exit 0 bukan berarti package loading berhasil. Bukti ini menjelaskan pembatasan penulisan cache; error T2-004 yang lebih awal adalah srcfiles cache entry missing. Hubungan mekanismenya didukung source Go `src/cmd/go/internal/work/exec.go`: cache srcfiles disimpan lewat cache.PutBytes, lalu loadCachedCompiledGoFiles membacanya. Tidak ada bukti bahwa versi Go lama atau scanner v1.7.0 sendiri menyebabkan kegagalan semula.

Scanner lama v1.7.0, dibangun Go1.26.7, sukses memuat package saat GOCACHE diarahkan ke cache writable `/tmp/c1-scanner-diagnostic-x9_cedbf/build-cache`. Eksekusi sandbox kemudian gagal DNS dengan `socket: operation not permitted` saat mengakses vuln.go.dev. Eksekusi ulang dengan escalation yang disetujui menyelesaikan full source scan, exit 0, stderr kosong, 0 finding objects. Jadi dua kebutuhan lingkungan adalah cache writable dan akses jaringan.

## Perbaikan lokal

GVM default telah ditetapkan Go1.26.9 dalam sesi sebelumnya. Dipasang `go install golang.org/x/vuln/cmd/govulncheck@v1.7.0` melalui escalation pada pkgset baru; scanner tersedia di `/home/anhar-solehudin/.gvm/pkgsets/go1.26.9/global/bin/govulncheck`, dibangun dengan Go1.26.9. Versi scanner tetap v1.7.0. Tidak menghapus cache/version lama atau mengubah source, go.mod/go.sum, konfigurasi global GOCACHE, maupun approval model/Run.

## Verifikasi selesai

Probe dibuat dari backend aktual di directory ini, mengganti hanya approved pgx v5.9.2 dan x/text v0.41.0, lalu go mod tidy (x/sync v0.22.0 terpilih transitif). Compile dan scan berjalan sequential. Ini bukan complete D8 gate atau penerimaan G1.

- Go1.26.9, GOCACHE baru: `go mod tidy` dan `go test -run '^$' ./internal/platform/auth ./internal/platform/db ./cmd/server` sukses setelah module-cache permission diberikan. Ketiga package no test files.
- Scanner lama, cache baru, network authorized: full `govulncheck -json ./...` exit 0, stderr kosong, 0 findings. JSON memuat config, SBOM, 200 OSV records, dan fetching/checking progress. OSV metadata bukan finding.
- Scanner pkgset baru, default cache, outside sandbox: full source scan probe exit 0, stderr kosong, 0 findings; JSON checked, bukan hanya exit status.
- Scanner pkgset baru, backend aktual, default cache, outside sandbox: full source scan exit 0 dalam mode JSON, stderr kosong, **15 finding objects untuk 5 IDs**: GO-2026-4771/4772/5004/5970/6629. Ini five-ID baseline findings yang masih ada karena actual pgx v5.8.0/text v0.29.0 belum diganti. Tidak ada suppression/risk acceptance. Full traces tersimpan dalam actual-backend-scan.json.

## Cara eksekusi berikutnya

Terminal lokal biasa: aktifkan GVM default Go1.26.9, masuk ke backend, jalankan `govulncheck ./...`. Cache home dapat ditulis dalam eksekusi normal.

Dalam sandbox Codex sesi ini: gunakan GOCACHE writable di /tmp untuk package loading dan minta escalation ketika source scan perlu akses database vuln.go.dev atau module cache home. Alternatifnya jalankan scan melalui escalation sejak awal, memakai default cache. Jangan mengubah permission home/cache agar sandbox dapat menerobos pembatasan; jangan clean cache sebagai default workaround. Jangan menganggap exit0 JSON membuktikan tidak ada finding.

## Batas hasil

Tooling sudah bekerja; actual backend dependency gate tetap belum bersih. Run T2-004 tetap terminal BLOCKED dan reportnya dipertahankan. Complete isolated/actual graph/source/integrity/compatibility gates dan post-lock expiry correction/focused tests tidak dijalankan di diagnosis ini. Tidak ada commit atau task dispatch.

## Capture identities

- `old-scanner-isolated-network.json`: SHA-256 `b78e2aae36509517097bd08e2c4994dd899e4c67ab71bd982a1d78914ed9604f` (611521 bytes).
- `old-scanner-isolated-network.stderr`: SHA-256 `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855` (0 bytes).
- `default-cache-package-load.json`: SHA-256 `aa83fe223636af7c1fcefed8f6922d9c36cb28a21af9ca92bc00f164f050bf25` (286 bytes).
- `new-scanner-default-cache.json`: SHA-256 `ef6c298e1c8de571c5af1ed1ae47470f9ed101930657e19d3b5d94c93509c914` (611521 bytes).
- `new-scanner-default-cache.stderr`: SHA-256 `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855` (0 bytes).
- `actual-backend-scan.json`: SHA-256 `99219c4bb8e965eb8601d86e3ecb13c437019b2497bd8fb733c8c05f9fd471e9` (644624 bytes).
- `actual-backend-scan.stderr`: SHA-256 `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855` (0 bytes).
