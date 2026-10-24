# Mode development dan pemeriksaan hot reload

Scope: development Kencleng, FE/BE host atau container. Mode container merupakan
konfigurasi development, memakai Next.js dev server dan Air; bukan image deployment
produksi. PostgreSQL tetap memakai Compose infrastruktur yang sudah ada.

| Pengaturan | Mode host | Mode container |
| --- | --- | --- |
| Backend | `backend/: make gvm-run` | `backend/: make docker-run` |
| Frontend | `frontend/: make nvm-run` | `frontend/: make docker-run` |
| Database dari backend | `127.0.0.1:5435` | `postgres:5432` |
| Proxy Next.js ke backend | `http://127.0.0.1:8090` | `http://backend:8090` |
| URL browser | `http://localhost:3000` | `http://localhost:3000` |
| Go / Node.js | gvm 1.27.3 / nvm 24.22.1 | image Go 1.27.3 / Node 24.22.1 |

Kedua mode memakai port yang sama. Hentikan runner mode sebelumnya sebelum
beralih; tidak diperlukan reset database. Pilih Docker atau Podman secara
eksplisit, karena masing-masing memakai store dan volume yang berbeda.

## Host

Setup mengikuti README [backend](../../backend/README.md) dan
[frontend](../../frontend/README.md). Mulai database dari root repo:

```bash
make COMPOSE=podman-compose up
```

Di dua terminal terpisah, masing-masing dimulai dari root Kencleng:

```bash
# Terminal BE:
cd backend
make gvm-run

# Terminal FE:
cd frontend
make nvm-run
```

Pastikan health langsung dan melalui proxy berhasil sebelum tes reload. Bila
shell mengekspor DATABASE_URL lama, gunakan `env -u DATABASE_URL make gvm-run`.

## Container melalui Docker Engine/Desktop atau Podman

Hentikan FE/BE host dengan `Ctrl+C` lebih dahulu. Dari root Kencleng:

```bash
# Podman:
make COMPOSE=podman-compose docker-run

# Atau Docker Engine / Docker Desktop:
# make docker-run
```

Perintah memastikan PostgreSQL menerima koneksi TCP dan lulus `db-check`, lalu
build/start kedua aplikasi di foreground. Unduhan image, Air, package Go/npm,
dan Google Fonts pada compile FE pertama membutuhkan jaringan.

Untuk tetap menjalankan setiap aplikasi dari direktori sendiri, pada dua terminal
terpisah yang dimulai di root Kencleng:

```bash
# Terminal BE:
cd backend
make docker-run COMPOSE=podman-compose

# Terminal FE:
cd frontend
make docker-run COMPOSE=podman-compose
```

Untuk Docker, hilangkan `COMPOSE=podman-compose`. FE dapat mulai sebelum BE,
tetapi request `/api/*` baru berhasil setelah BE siap. Masing-masing target
hanya menjalankan service aplikasinya; database dipastikan siap oleh helper root.

Dari root repo, gunakan provider yang sama:

```bash
make COMPOSE=podman-compose docker-status
make COMPOSE=podman-compose docker-logs
make COMPOSE=podman-compose docker-stop
make COMPOSE=podman-compose docker-down
```

`docker-stop` menghentikan aplikasi; target yang sama dari direktori BE/FE hanya
menghentikan aplikasinya. `docker-down` melepas kedua container aplikasi dan
mempertahankan database serta volume cache. Setelah itu mode host dapat dipakai
kembali. `make down` dari root melepas container aplikasi, SeaweedFS, lalu
PostgreSQL tanpa menghapus volume. `docker-config` merender konfigurasi aplikasi melalui
provider yang dipilih.

### Isolasi konfigurasi dan file

- Aplikasi memakai project `kencleng-apps` dan `docker-compose.apps.yml`, terpisah
  dari hash/lifecycle Compose PostgreSQL. Network database eksternal default
  `kencleng_default`; override `KENCLENG_INFRA_NETWORK` jika infrastruktur memang
  menggunakan nama lain. Menambahkan service aplikasi tidak mengubah file Compose
  database atau menghapus datanya.
- Make menambahkan `docker-compose.apps.podman.yml` hanya jika provider `COMPOSE`
  bernama Podman. Override memakai `keep-id` untuk rootless Podman; Docker tidak
  menerima file khusus ini. Ini penyesuaian provider yang sudah dipilih, bukan
  pemilihan engine otomatis.
- Image menjalankan aplikasi sebagai user non-root, dengan UID/GID pengguna host
  yang diteruskan Make melalui `DEV_UID`/`DEV_GID`. Jalankan tanpa sudo; kedua ID
  harus lebih besar dari nol. Source di-bind mount agar perubahan editor terlihat
  oleh watcher. File hasil development yang ditulis ke source memiliki UID yang
  sama dengan pengguna host pada konfigurasi Linux ini.
- `node_modules`, `.next`, module/build cache Go, dan binary sementara Air memakai
  named volume sendiri. Dependency/cache container tidak mengganti isi dependency
  host. Entrypoint FE menjalankan `npm ci` bila manifest, lockfile, atau versi Node
  berbeda dari stamp volume. Perubahan dependency memerlukan restart/rebuild
  container; perubahan source aplikasi tidak memerlukan rebuild image.
- Air container memakai polling dan tidak menghapus root volume `tmp` saat keluar.
  Next.js container memakai Webpack dengan `WATCHPACK_POLLING=true`, sehingga
  watcher juga dapat menangkap perubahan pada file sharing Desktop. Polling dapat
  memakai lebih banyak CPU daripada watcher native. Mode host tetap memakai
  konfigurasi Air native dan Next.js Turbopack.
- `APP_HOST=0.0.0.0` hanya pada backend container; port yang dipublikasikan ke host
  tetap loopback. Backend menerima env container dan tidak mewajibkan file `.env`
  jika nilai wajib sudah tersedia. `.env` malformed tetap menyebabkan kegagalan;
  nilai runtime tetap menang atas dotenv host. File env dan cache dikecualikan
  dari build context image. Bind mount development tetap memuat checkout lokal.

## Perubahan kode tanpa rebuild image

Source lokal di-bind mount ke container. Saat menyimpan kode React/CSS, Next.js
memperbarui tampilan melalui Fast Refresh. Saat menyimpan kode Go, Air
mengompilasi ulang binary dan me-restart backend di dalam container yang sama.
Perubahan source tersebut tidak memerlukan build ulang image atau menjalankan
kembali `make docker-run`.

`docker-run` memakai `up --build` saat memulai runner. Perubahan Dockerfile/paket
OS memerlukan rebuild image; perubahan dependency atau environment memerlukan
restart/rebuild sesuai konfigurasi. Mekanisme watcher dan volume dijelaskan di atas.

## Pemeriksaan hot reload yang nyata

`make check-hot-reload` adalah alat verifikasi opt-in; watcher Next.js/Air aktif
melalui runner. Tes ini tidak perlu dijalankan setiap kali mengedit kode.

Jalankan setelah kedua aplikasi dan health endpoint berhasil, sekali pada mode
host, lalu ulangi pada mode container. Dari `frontend/`, dengan Node project aktif:

```bash
nvm use
make check-hot-reload
```

Script menggunakan Playwright yang sudah ada di dependency frontend. Jika Google
Chrome tersedia pada `/usr/bin/google-chrome`, script memakainya; jika tidak,
dibutuhkan Chromium Playwright (`npm run browser:install`). Path browser custom
dapat diberikan lewat `PLAYWRIGHT_CHROMIUM_EXECUTABLE_PATH`. Ini kebutuhan alat
verifikasi, bukan prasyarat menjalankan container aplikasi.

Script memeriksa health BE/proxy dan membuka browser terlebih dahulu. Kemudian ia
menambahkan sementara header probe pada handler health BE serta marker tampilan
pada halaman FE. Keberhasilan mensyaratkan:

1. Server yang dibangun ulang menyajikan header probe, langsung dan lewat proxy.
2. Browser yang sudah terbuka menampilkan marker tanpa pemanggilan `page.reload()`.
3. Kedua file source dikembalikan dan aplikasi kembali menyajikan hasil awal.

Jangan mengedit kedua file tersebut selama tes. Copy pemulihan dibuat di `/tmp`;
normal exit, error, SIGINT dan SIGTERM mencoba memulihkan source. Jika ada edit
bersamaan atau pemulihan gagal, script tidak menimpa perubahan lain dan
menampilkan lokasi copy pemulihan. Penghentian paksa seperti SIGKILL tidak dapat
menjalankan cleanup; periksa working tree dan copy pemulihan sebelum mencoba lagi.

Untuk preview tanpa request/browser atau perubahan source, dari `frontend/`:

```bash
node ../scripts/check-hot-reload.mjs --dry-run
```

Origin custom dapat diberikan lewat `DEV_FRONTEND_URL`/`DEV_BACKEND_URL`, tetapi
hanya loopback tanpa credential/path/query yang diterima. Script tidak menyalakan
atau menghentikan runner dan tidak membuktikan engine mana yang digunakan; hasil
harus dicocokkan dengan mode runner dan `docker-status` yang benar-benar diamati.

## Standing verifikasi

Workflow host FE/BE dengan PostgreSQL melalui Podman siap digunakan berdasarkan
health check yang dilaporkan berhasil oleh Human. Konfigurasi Fast Refresh dan
Air tersedia; keberhasilan probe otomatis perubahan file belum tercatat.
**Startup dan hot reload aktual FE/BE dalam container belum tercatat lulus**;
kompatibilitas Docker Engine/Desktop juga masih memerlukan pengujian tersendiri.

Pada sesi agent 24 Oktober 2026, preflight HTTP ditolak `connect EPERM`
sebelum source diubah, dan Podman menolak akses runtime directory sandbox.
Hash kedua file source dipastikan tidak berubah setelah percobaan probe tersebut.
Go env tests, lint/TypeScript FE, pemeriksaan YAML/shell serta simulasi refresh
dependency volume dan reset scope dilakukan terpisah dari bukti runtime.

Build image, UID/file watching, startup/health dan probe reload aplikasi perlu
benar-benar dijalankan pada masing-masing engine sebelum klaim kompatibilitasnya.
Setup, S3 dan persistensi SeaweedFS melalui Podman sudah dilaporkan berhasil oleh
Human; bukti storage ini terpisah dari verifikasi container aplikasi. Lihat
[panduan storage](local-object-storage-options.md) dan
[status setup repo](kencleng-repo-setup.md#status-pemeriksaan-24-oktober-2026).
