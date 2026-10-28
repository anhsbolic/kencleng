# Kencleng Backend

> Status: clean delivery baseline

Neutral Go backend engineering substrate for Kencleng Pilot #3. No product domain, authentication model, API operation, database migration, or background job is considered delivered at baseline.

Product meaning is owned by `../docs/product/product-intent.md`. Backend architecture is owned by `../docs/engineering/standards/backend.md`.

## Setup

Backend berjalan di host dengan Go **1.27.3** melalui gvm. PostgreSQL **16**
berjalan melalui Compose; pilihan image dan topologi dimiliki oleh
[panduan setup repo](../docs/engineering/development/kencleng-repo-setup.md).

Dari direktori `backend/`, dengan gvm sudah dimuat di shell:

```bash
gvm use "$(cat ../.go-version)"
make setup
```

Setup membuat `.env` dari `../.env.example` hanya jika belum ada, lalu
memasang Air **v1.67.4** ke `bin/`. Instalasi Air membutuhkan jaringan;
tidak menambahkan dependency ke modul aplikasi. File `.env`, binary, dan cache
build Air tidak di-commit.

Periksa konfigurasi lokal tanpa mengganti kredensial yang sudah dipakai:

```dotenv
APP_ENV=development
APP_HOST=127.0.0.1
APP_PORT=8090
DATABASE_URL=postgres://kencleng:kencleng@127.0.0.1:5435/kencleng?sslmode=disable
```

`APP_HOST` mengatur alamat bind server. Jika kosong, Go mendengarkan pada semua
interface; gunakan `127.0.0.1` untuk mode host lokal. Server memuat `.env` dari
direktori kerja, sehingga jalankan perintah aplikasi dari `backend/`. File boleh
tidak ada jika environment runtime menyediakan `APP_PORT` dan `DATABASE_URL`;
file yang malformed atau tidak dapat dibaca tetap menyebabkan kegagalan.

Reset volume Podman mempertahankan `.env`. Setelah mulai dari database kosong,
pastikan `DATABASE_URL` sesuai user/database/password pada Compose saat ini;
konfigurasi lokal dari episode lama tidak otomatis ikut berubah. Environment
yang sudah diekspor oleh shell mengalahkan nilai `.env` karena `godotenv.Load`
tidak melakukan overwrite. Jika runner masih menyebut akun lama, hentikan Air
lalu jalankan dari `backend/` dengan `env -u DATABASE_URL make run` agar
runner tersebut memakai nilai `.env`, tanpa mengubah environment shell induk.

## How to run

Dari root repo, mulai PostgreSQL dengan **salah satu** provider:

```bash
make COMPOSE=podman-compose up
# Docker Engine atau Docker Desktop:
# make up
```

Perintah menunggu PostgreSQL menerima koneksi lalu memeriksa kecocokan collation
database dengan image. Data lama project pernah mencatat versi 2.41; keluarga
image Trixie dipertahankan dan `db-check` memeriksa nilai aktual. Jika pemeriksaan gagal, ikuti
[panduan diagnosis](../docs/engineering/development/kencleng-repo-setup.md#pod-lama-dan-peringatan-collation)
sebelum mengubah database. Lalu, dari direktori `backend/`, jalankan:

```bash
make gvm-run
```

`make gvm-run` memuat gvm dari `GVM_ROOT` (default `~/.gvm`), mengaktifkan versi
di `../.go-version`, lalu menjalankan Air. Aktivasi berlaku pada proses runner;
tidak mengubah shell induk. gvm dan versi Go tersebut harus sudah terpasang.
Jika Go sudah diaktifkan, gunakan `make run` atau alias `make dev`.
Air membangun ulang dan
me-restart server saat source Go, `go.mod`, `go.sum`, atau `.env` berubah.
`Ctrl+C` menghentikan watcher dan server. Untuk menjalankan sekali tanpa reload,
gunakan `go run ./cmd/server`.

Pemeriksaan dari terminal lain:

```bash
curl --fail http://127.0.0.1:8090/healthz
# Dengan frontend juga berjalan:
curl --fail http://localhost:3000/api/healthz
```

Keduanya diharapkan mengembalikan `{"status":"ok"}`. Next.js menghapus prefix
`/api`; endpoint backend tetap `GET /healthz`. Object storage tidak diperlukan untuk
scaffold ini. Setelah menghentikan aplikasi, `make COMPOSE=podman-compose down`
dari root menghentikan infrastruktur tanpa menghapus volume data.

## Container development

Hentikan backend host terlebih dahulu. Dari `backend/`:

```bash
make docker-run COMPOSE=podman-compose
# Docker Engine / Desktop: make docker-run
```

Target ini memastikan database siap lalu build/run backend dalam container
non-root dengan Go/Air yang dipin. Source di-bind mount, cache/binary sementara
memakai volume sendiri, dan polling Air menangkap perubahan file. URL health
host tetap `http://127.0.0.1:8090/healthz`; koneksi database dari container memakai
`postgres:5432`. Env container mengalahkan `.env` host yang masih dimuat dari
checkout. `make docker-stop COMPOSE=podman-compose` menghentikan backend container.

Jalankan frontend container dari direktori frontend pada terminal terpisah.
Untuk kembali ke mode host, jalankan root `make COMPOSE=podman-compose docker-down`
terlebih dahulu; database dan volume cache dipertahankan. Detail UID, provider,
dan pemeriksaan reload: [mode development](../docs/engineering/development/local-development-modes.md).

## Optional S3 storage

SeaweedFS lokal tersedia lewat root `make storage-up COMPOSE=podman-compose`.
Endpoint host `http://127.0.0.1:9087`; endpoint container `http://seaweedfs:8333`.
Scaffold belum mengonsumsi client S3. Setup, pemeriksaan S3 dan persistensi melalui
Podman sudah dilaporkan berhasil oleh Human. Ikuti
[setup storage dan batas deployment](../docs/engineering/development/local-object-storage-options.md)
untuk kredensial lokal, perintah check, dan penanganan presigned URL.

## Verification

Dari `backend/`, gunakan `go test ./...`, `go vet ./...`, dan `go build ./...`.
`make verify` menjalankan pemeriksaan yang lebih luas dan membutuhkan
`staticcheck`, `gosec`, `gitleaks`, serta `govulncheck` secara terpisah.
Output `[no test files]` bukan bukti pengujian perilaku.

Scaffold belum memiliki migrasi aktif. Jangan menjalankan migrasi, reset, atau
`migrate-down` sebagai bagian setup umum. Add product behavior only after the
relevant delivery requirement is sufficiently clear.
