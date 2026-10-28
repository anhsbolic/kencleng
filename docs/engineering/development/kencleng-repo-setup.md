# Kencleng — Repository Structure & Local Setup

> Status: Clean delivery baseline candidate
> Purpose: Own repository topology and local-development setup. This document does not choose release scope, delivery order, or product capability priorities.

## Repository posture

Kencleng is a monorepo containing product/design authority, backend and frontend engineering substrates, API-contract tooling, and local infrastructure.

The Pilot #3 delivery baseline intentionally preserves settled engineering substrate while resetting active delivery state. Git history remains evidence and archive; removed delivery specs, contracts, implementations, and workflow artifacts are not active authority merely because they remain in history.

## Repository map

```text
kencleng/
├── docs/
│   ├── product/       # approved upstream product intent
│   ├── ui-ux/         # approved reusable product-design authority
│   └── project/       # repository and stack architecture
├── api/               # OpenAPI tooling; contract starts empty until needed
├── backend/           # neutral Go engineering substrate
├── frontend/          # neutral Next.js engineering substrate
├── AGENTS.md
├── Makefile
├── .go-version          # Go version selected by gvm-run
├── .nvmrc               # Node.js version selected by nvm-run
├── docker-compose.yml
├── docker-compose.apps.yml # optional FE/BE dev containers
├── docker-compose.apps.podman.yml # rootless Podman app override
├── docker-compose.storage.yml # optional SeaweedFS S3 capability
└── docker-compose.minio.yml # legacy optional MinIO configuration
```

No release goal, slice, delivery spec, API operation, or feature implementation is implied by the existence of these directories.

## Local topology

Mode default menjalankan backend dan frontend di host. Compose utama hanya menjalankan
PostgreSQL 16, dengan image `docker.io/library/postgres:16.15-trixie`.
Konfigurasi SeaweedFS dipisahkan ke `docker-compose.storage.yml`;
tidak dimuat oleh perintah standar. Mode FE/BE dalam container tersedia melalui
`docker-compose.apps.yml` pada project terpisah `kencleng-apps`; lihat
[mode development dan pemeriksaan hot reload](local-development-modes.md).

```text
browser
  ↓
localhost:3000
  ↓
host Next.js (Fast Refresh)
  ├─ /api/* → host backend 127.0.0.1:8090 (prefix /api dihapus)
  └─ other  → halaman frontend

host backend (Air rebuild/restart)
  └─ PostgreSQL 127.0.0.1:5435 → container :5432

optional object storage
  └─ SeaweedFS S3 127.0.0.1:9087
```

Browser memakai URL relatif `/api/...`; `BACKEND_INTERNAL_URL` adalah konfigurasi
server Next.js di `frontend/.env.local`, default `http://127.0.0.1:8090`.
Rewrite hanya aktif saat `next dev`. Gateway API produksi tetap keputusan terpisah.
Tidak ada Caddy atau kebutuhan koneksi container ke host pada topologi standar.

Executable configuration remains authoritative for exact ports and commands.

## Commands

### Local toolchain (Ubuntu 24 + nvm + gvm)

Versi development project ini adalah **Go 1.27.3** dan **Node.js 24.22.1**.
`.go-version` menjadi penanda versi untuk gvm; `.nvmrc` menjadi penanda versi
untuk nvm. `backend/go.mod` mensyaratkan Go minimal 1.27.3. Manifest frontend
dan tooling API mensyaratkan Node.js 24.x mulai 24.22.1; `.nvmrc` memilih versi
development yang tepat. Field `engines` memberi peringatan npm pada versi yang
tidak sesuai; field itu sendiri tidak mengganti versi Node.js.

Dengan nvm dan gvm yang sudah dimuat di shell, jalankan dari root repo untuk
memasang versi yang belum tersedia:

```bash
nvm install
gvm install "$(cat .go-version)" -B
```

Perintah ini mengunduh runtime ke direktori nvm/gvm milik pengguna. Untuk setiap
sesi development, aktifkan versi project dari root repo:

```bash
nvm use
gvm use "$(cat .go-version)"
node --version
go version
```

Hasil yang diharapkan: Node.js `v24.22.1` dan Go `go1.27.3`. `nvm use` dapat
menemukan `.nvmrc` dari subdirektori. gvm yang memuat integrasi `cd` dapat membaca
`.go-version` ketika masuk ke project; aktivasi eksplisit di atas tetap dapat
digunakan. Default global manager tidak perlu diubah.

Setup dan runner aplikasi dimiliki oleh direktori masing-masing. Dari terminal
yang dimulai di root Kencleng, setup backend:

```bash
cd backend
gvm use "$(cat ../.go-version)"
make setup
```

Pada terminal lain yang dimulai di root Kencleng, setup frontend:

```bash
cd frontend
nvm use
make setup
```

Setup backend mempertahankan/membuat `.env` dari `../.env.example`, lalu memasang Air v1.67.4
ke `backend/bin/`. Setup frontend mempertahankan/membuat `frontend/.env.local`,
lalu menjalankan `npm ci`. Keduanya membutuhkan akses jaringan saat dependency
belum tersedia. File env yang ada tidak ditimpa. Untuk konfigurasi backend lama,
tambahkan `APP_HOST=127.0.0.1` agar server hanya mendengarkan loopback.

Tooling API memakai versi Node.js yang sama; pasang dependency-nya dengan
`cd api && npm ci` hanya ketika pekerjaan membutuhkan tooling tersebut.

### Repository commands

Semua perintah di bawah dijalankan dari root repo. Pilih satu provider secara
eksplisit; tidak ada pemilihan engine otomatis:

```bash
# Podman dengan provider podman-compose:
make COMPOSE=podman-compose up
make COMPOSE=podman-compose status
make COMPOSE=podman-compose logs
make COMPOSE=podman-compose down

# Docker Engine atau Docker Desktop:
make up
make status
make logs
make down
```

`COMPOSE` juga dapat diisi `"podman compose"` bila provider wrapper tersebut sudah
dikonfigurasi; dukungan fitur mengikuti provider yang dipakainya. `make up-podman`
dan `make down-podman` tetap tersedia sebagai alias untuk `podman-compose`.
`make config` (atau dengan `COMPOSE=podman-compose`) merender konfigurasi.
`status`, `logs`, dan `config` di root hanya mencakup PostgreSQL. Untuk aplikasi
container gunakan `docker-status`/`docker-logs`/`docker-config`; untuk SeaweedFS
gunakan `storage-status`/`storage-logs`/`storage-config`, dengan provider yang sama.

`up` hanya memulai PostgreSQL lalu menunggu `pg_isready`, maksimal 30 percobaan
dengan jeda 2 detik, kemudian menjalankan `db-check` untuk membandingkan versi
collation tercatat dan aktual pada database yang menerima koneksi. Jika berbeda,
perintah gagal dan menampilkan diagnosis tanpa mengubah metadata. `make
COMPOSE=podman-compose db-check` dapat dipakai terpisah. Healthcheck juga tersedia
di service PostgreSQL. Pemeriksaan ini belum membuktikan integritas seluruh
indeks atau koneksi backend; periksa endpoint aplikasi setelahnya.
Port yang dipublikasikan Compose dibatasi ke `127.0.0.1`. `down` melepas app
container, SeaweedFS, lalu database, tanpa `--volumes`; volume data dan cache tetap ada.
Untuk hanya melepas app container dan mempertahankan database, gunakan
`docker-down`. Data Docker Engine, Docker Desktop,
dan Podman berada pada storage engine masing-masing dan tidak otomatis dibagi.

Pada dua terminal terpisah yang dimulai di root Kencleng, jalankan:

```bash
# Terminal backend:
cd backend
make gvm-run

# Terminal frontend:
cd frontend
make nvm-run
```

Buka `http://localhost:3000`. Periksa backend dengan
`curl --fail http://127.0.0.1:8090/healthz` dan proxy dengan
`curl --fail http://localhost:3000/api/healthz`. Keduanya diharapkan menghasilkan
`{"status":"ok"}`. Simpan perubahan FE dan periksa Fast Refresh di browser;
perubahan Go harus memicu rebuild/restart Air. `Ctrl+C` menghentikan aplikasi.

Backend dan frontend memiliki perintah masing-masing di `backend/Makefile`,
`frontend/Makefile`, dan `frontend/package.json`. Lihat README masing-masing untuk pemeriksaan kode
dan prasyarat `make verify` yang lebih luas.

### Mode container dan verifikasi reload

Mode host tetap default. Untuk menjalankan FE/BE di container, hentikan runner
host lalu jalankan `make docker-run COMPOSE=podman-compose` dari masing-masing
direktori aplikasi. Docker Engine/Desktop memakai `make docker-run`. Kedua mode
memakai port yang sama dan tidak dijalankan bersamaan.

Root `make COMPOSE=podman-compose docker-run` juga dapat menjalankan kedua aplikasi
sekaligus. Project aplikasi `kencleng-apps` terpisah dari project database dan
memakai network database sebagai external network. `docker-down` menghentikan/
melepas aplikasi saja, tanpa reset database atau cache.

Pemeriksaan opt-in `frontend/: make check-hot-reload` memakai browser serta
perubahan source sementara yang dipulihkan untuk membuktikan reload BE, proxy,
dan pembaruan browser FE. Baca [panduan lengkap](local-development-modes.md)
untuk prasyarat, perubahan source, UID/cache, perpindahan mode, dan standing bukti.
`gvm-run` memuat manager dari `GVM_ROOT` atau `~/.gvm`; `nvm-run` dari `NVM_DIR`
atau `~/.nvm`. Keduanya memerlukan manager/runtime yang sudah terpasang dan hanya
mengubah environment proses runner, bukan shell induk. `make run` di tiap
direktori memakai runtime yang sudah aktif. Frontend juga menyediakan Makefile
untuk `setup`, `run`, dan `verify`. Alias
lama di root mendelegasikan pekerjaan ke Makefile aplikasi; panduan utama memakai
direktori aplikasi. Perintah Compose tetap dari root Kencleng.

### Object storage opsional: SeaweedFS

SeaweedFS dipilih untuk persiapan capability S3 lokal dan evaluasi deployment.
Baca [pilihan, setup dan batas kesiapan storage](local-object-storage-options.md).
Dari root repo:

```bash
make COMPOSE=podman-compose storage-up
make COMPOSE=podman-compose storage-check
make COMPOSE=podman-compose storage-persistence-check
make COMPOSE=podman-compose storage-down
```

Untuk Docker Engine/Desktop, hilangkan override `COMPOSE`. Pemeriksaan persistensi
me-restart hanya SeaweedFS; jalankan saat tidak ada upload aktif. Image server dan
AWS CLI dipin ke digest yang diverifikasi di registry. Port host/browser adalah
`127.0.0.1:9087`, endpoint BE container `http://seaweedfs:8333` pada jaringan
infrastruktur yang sama. PostgreSQL tetap terpisah dari metadata object storage.
Backend scaffold belum mengonsumsi client S3; env contoh adalah persiapan integrasi.

Storage memakai project `kencleng-storage` dan volume baru `seaweedfs_data`,
terpisah dari data MinIO. `storage-down` mempertahankan volume. Root `make down`
melepas app container dan SeaweedFS sebelum PostgreSQL. Konfigurasi lokal memakai
kredensial development dan bukan deployment produksi yang terverifikasi.

Konfigurasi legacy `docker-compose.minio.yml` dan target `minio-up`/`minio-down`
tetap tersedia, memerlukan `MINIO_IMAGE` dan `MINIO_MC_IMAGE` yang dipin dan
diverifikasi sendiri. Target `storage-*` tidak menjalankan MinIO, mengambil volume
lamanya, atau memigrasikan objek. Bila server lama masih memakai port 9087,
hentikan secara eksplisit sebelum menjalankan SeaweedFS.

### Beralih dari setup Caddy lama

Caddy tidak lagi didefinisikan; browser berpindah dari port 8080 ke 3000. Jika
container Caddy atau MinIO lama masih berjalan, periksa dulu nama/project dan
datanya sebelum menghentikan container yang tepat. Perubahan file ini tidak
menghapus container lama atau volume. Jangan memakai penghapusan orphan/volume
massal sebagai langkah setup.

### Pod lama dan peringatan collation

Laporan terminal Human pada 24 Oktober 2026 menunjukkan PostgreSQL menerima
koneksi, tetapi pod/network lama tidak dapat dihapus karena dua container lain
masih berjalan. Pada podman-compose 1.5.0, perubahan hash konfigurasi dapat
memicu `down`; service yang sudah tidak didefinisikan dapat tertinggal. File
Compose utama, aplikasi, SeaweedFS dan MinIO legacy memakai `x-podman.in_pod: false` agar service baru tidak berbagi
pod. Ini tidak membersihkan container lama. Identifikasi ID, nama, label project,
dan kebutuhan datanya sebelum menghentikan atau melepas container tertentu.

Laporan yang sama menunjukkan database mencatat collation 2.41, sedangkan
image Bookworm yang sempat dipilih menyediakan 2.36. Pemilihan tersebut tidak
sesuai dengan volume yang sudah ada. Pin dikoreksi ke varian Trixie untuk
mempertahankan keluarga OS/libc data, tetap pada PostgreSQL 16. Human kemudian
memilih reset data lokal dan melaporkan startup serta pemeriksaan database
berhasil. Untuk volume lain yang masih membawa data lama, verifikasi nilai aktual
sebelum mengganti image. Jangan mengganti varian OS hanya untuk memilih tag yang eksplisit.

Jangan langsung menjalankan `ALTER DATABASE ... REFRESH COLLATION VERSION` untuk
menghilangkan warning. [Dokumentasi PostgreSQL](https://www.postgresql.org/docs/16/sql-altercollation.html)
menjelaskan bahwa refresh metadata tidak memeriksa apakah objek yang terdampak
sudah dibangun ulang. Jika database/indeks sempat diubah saat versi collation
berbeda, penilaian integritas dan kebutuhan pemulihan tetap diperlukan. Tidak ada
reset volume, reindex, atau perubahan metadata yang dijalankan dalam perbaikan ini.

### Reset Podman terarah

Atas permintaan Human untuk mulai ulang, tersedia
[`scripts/reset-kencleng-podman.py`](../../../scripts/reset-kencleng-podman.py).
Script Python 3 ini tidak memakai dependency tambahan. Jalankan dari root repo,
tanpa `sudo`, pada user/store Podman yang sama dengan setup sebelumnya. Hentikan
FE/BE host terlebih dahulu melalui `Ctrl+C`.

```bash
# Inventaris target reset penuh, read-only:
python3 scripts/reset-kencleng-podman.py --delete-data

# Eksekusi setelah meninjau daftar, lalu ketik DESTROY KENCLENG DATA:
python3 scripts/reset-kencleng-podman.py --apply --delete-data

# Buat ulang database kosong dengan image/config project:
make COMPOSE=podman-compose up
```

`--delete-data` menghapus isi volume PostgreSQL, object storage dan cache yang terpilih secara
permanen. Backup dahulu bila data masih diperlukan; tidak ada backup otomatis.
Untuk reset container saja, gunakan `--apply` tanpa `--delete-data` dan konfirmasi
`RESET KENCLENG CONTAINERS`. Volume yang dipertahankan masih membawa database lama,
sehingga opsi tersebut bukan instalasi dari data kosong.

Container dipilih lewat label project yang tepat: `kencleng`, `kencleng-apps`, dan
`kencleng-storage`, memakai label Podman Compose maupun Docker Compose. Cache
volume container aplikasi ikut menjadi target jika `--delete-data` dipilih.
Network/volume tanpa label hanya dicocokkan dengan nama lama yang eksplisit:
network `<project>_default`, volume `kencleng_kencleng_pgdata`,
`kencleng_kencleng_miniodata`, dan `kencleng-storage_kencleng_miniodata`.
Pod dipilih dari label, hubungan container target, atau nama persis
`pod_kencleng`/`pod_kencleng-apps`/`pod_kencleng-storage`. Label project asing
mengalahkan kecocokan nama.
Namespace/resource lain tidak disapu berdasarkan prefix nama.

Keanggotaan pod diperiksa sebelum perubahan; jika ada container di luar target,
script berhenti. Pod hanya dilepas setelah container target dihapus dan diperiksa
ulang bahwa tidak ada anggota lain selain infra container. Penghapusan network
dan volume tidak memakai force: resource yang masih dipakai container lain harus
ditolak Podman. Script berhenti pada error; reset parsial dilaporkan sebagai gagal.
Tidak ada `system prune`, penghapusan semua resource, atau penghapusan image.
File repo, env, dan instalasi toolchain host tidak disentuh.

Script belum dijalankan pada resource nyata dari sesi agent. Pemeriksaan syntax
dan safety regression memakai Podman mock; jalankan pemeriksaannya dengan:

```bash
PYTHONDONTWRITEBYTECODE=1 python3 -m unittest discover -s scripts -p 'test_*.py' -v
```

Reset mempertahankan file env, sehingga `backend/.env` lama dapat masih menunjuk
akun/database yang tidak ada pada volume baru. Samakan `DATABASE_URL` dengan
konfigurasi Compose yang dipilih sebelum menjalankan backend. Jika shell masih
mengekspor nilai lama, hentikan Air dan jalankan `env -u DATABASE_URL make run`
dari `backend/`; `godotenv.Load` memberi prioritas pada environment yang
sudah diekspor. Tidak perlu membuat ulang role/database historis untuk mengatasi
ketidakcocokan konfigurasi lokal tersebut.

Setelah reset/up berhasil, jalankan kembali `make gvm-run` dari `backend/` dan
`make nvm-run` dari `frontend/` di terminal terpisah; `make run` tersedia bila
runtime sudah aktif. Periksa URL health yang dijelaskan
di atas. Tidak perlu memasang ulang Podman, runtime Go/Node, atau dependency
aplikasi hanya karena container dibuat ulang.

### Troubleshooting reset: rootless netns permission denied

Jika reset berhenti dengan `rootless netns: kill network process: permission denied`,
jangan menganggap resource sudah bersih. Script tetap gagal dan tidak meneruskan
penghapusan network/volume setelah kegagalan tahap container. Sebagian container
bisa sudah terhapus; inventaris ulang sebelum mengulang apply.

Pada pemeriksaan host 24 Oktober 2026, log kernel menunjukkan AppArmor menolak
`signal=term`, `profile="pasta"`, `peer="podman"`, dengan permission `receive`.
Penyebab ini teramati pada host yang diperiksa, bukan diagnosis otomatis untuk
semua pesan permission denied. Kasus kebijakan distro ini juga dibahas di
[issue Podman #27372](https://github.com/podman-container-tools/podman/issues/27372).

Pemeriksaan read-only dari terminal lokal:

```bash
journalctl -k --since '-20 min' --no-pager -g 'apparmor=.*DENIED.*(pasta|podman)'
```

Untuk denial yang persis sama, usulan perbaikan administratif adalah mengizinkan
profil `pasta` menerima sinyal TERM dari peer berprofil `podman`:

```text
signal (receive) set=(term) peer=podman,
```

Aturan ini dapat ditempatkan di `/etc/apparmor.d/local/usr.bin.pasta`, tetapi
profil `/etc/apparmor.d/usr.bin.pasta` harus menyertakan baris berikut di dalam
blok `profile pasta { ... }` agar local rule dimuat:

```text
include if exists <local/usr.bin.pasta>
```

Pada host yang diperiksa, file local ada tetapi kosong dan include tersebut
belum ada. Kandidat aturan diperiksa dengan `apparmor_parser --skip-kernel-load
--skip-cache`; ini pemeriksaan syntax, bukan penerapan kebijakan. Penerapan perlu
akses administrator dan keputusan Human karena mengubah policy host untuk
proses `pasta`/`podman`, tidak terbatas pada project Kencleng. Jangan menonaktifkan
AppArmor, memakai `sudo podman`, atau menjalankan `system reset`/`system migrate`
global sebagai perbaikan otomatis script.

Setelah perubahan administratif ditinjau dan diterapkan, muat ulang profil dengan
`sudo apparmor_parser -r /etc/apparmor.d/usr.bin.pasta`, jalankan preview reset,
lalu apply ulang jika daftar target sudah benar. Verifikasi runtime tetap perlu;
script project tidak mengubah file `/etc` atau memuat policy kernel.

## Status pemeriksaan (24 Oktober 2026)

Scope: perubahan setup pada working tree `pilot/station5-c1-space-only`, berbasis
commit `2e18c87c8435be8009a66853973e0ea3b2d36632`.

| Cakupan | Status dan sumber bukti |
| --- | --- |
| Workflow host: FE/BE + PostgreSQL melalui Podman | Siap untuk development berdasarkan health check yang dilaporkan berhasil oleh Human; Next.js Fast Refresh dan Go/Air sudah dikonfigurasi. |
| SeaweedFS lokal melalui Podman | Setup, pemeriksaan S3 dan persistensi setelah restart dilaporkan berhasil oleh Human. |
| Hot reload FE/BE dalam container | Konfigurasi tersedia; startup dan probe perubahan kode aktual belum tercatat lulus. |
| Docker Engine/Desktop | Konfigurasi tersedia; verifikasi engine tersebut belum tercatat. |

Laporan Human adalah bukti dari terminal lokalnya; agent tidak mengklaim
menjalankan pemeriksaan runtime yang ditolak sandbox. Kesiapan ini terbatas pada
lingkungan development dan tidak menetapkan kesiapan deployment produksi.

Rincian pemeriksaan dan episode troubleshooting:

- Terverifikasi dengan Go 1.27.3: `go test ./...`, `go vet ./...`, `go build ./...`,
  serta `go mod tidy -diff`. Test konfigurasi startup backend kini mencakup env
  container tanpa dotenv, required keys, prioritas env, dan dotenv malformed;
  `go test -race ./cmd/server` juga lolos. Pemeriksaan memakai
  dependency cache yang sudah tersedia dan build cache di `/tmp`.
- Terverifikasi dengan Node.js 24.22.1: `npm run verify` dan TypeScript. Vitest
  belum memiliki test case. Loader dan matcher Next.js berhasil memeriksa default
  upstream, penghapusan `/api`, custom upstream, dan rewrite yang tidak aktif
  pada fase produksi. Ini pemeriksaan konfigurasi, bukan request HTTP end-to-end.
- YAML Compose, TOML Air, syntax shell init, serta dry-run perintah Make diperiksa.
  Validasi Compose melalui provider belum berhasil karena akses Podman ke runtime
  directory ditolak sandbox. Docker tidak tersedia pada sesi pemeriksaan.
- Unduhan Air dari sesi agent sebelumnya ditolak sandbox; binary lokal sekarang
  tersedia dari setup Human dan opsi CLI polling/clean-on-exit diperiksa.
- Pembaruan Human: reset Podman dan startup PostgreSQL selesai; setelah konfigurasi
  database lokal diselaraskan, kedua pemeriksaan backend/frontend dilaporkan
  berhasil. Ini bukti yang dilaporkan Human; akses HTTP dari sesi agent tetap
  dibatasi. Masih perlu pembuktian reload lewat perubahan file serta pengujian
  Docker Engine/Desktop terpisah. Pembaruan Human: setup SeaweedFS melalui Podman,
  pemeriksaan S3 dan persistensi setelah restart dilaporkan selesai tanpa error.
  Output/log tidak dilampirkan; ini bukti yang dilaporkan Human, bukan eksekusi
  agent atau bukti kesiapan produksi. Detail ada di
  [standing storage](local-object-storage-options.md#versi-yang-dipin-dan-standing-verifikasi).
- Pembaruan berdasarkan laporan Human: PostgreSQL sempat berhasil start dengan
  volume lama, tetapi ada pod/container tersisa dan mismatch collation 2.41/2.36.
  Pin image dan pemeriksaan metadata telah dikoreksi pada working tree; perbaikan
  runtime sebelumnya belum diverifikasi. Human kemudian menjalankan reset terarah
  dan melaporkan hasil startup berhasil setelah perbaikan AppArmor; lihat
  pembaruan Human di atas. Agent tidak mengeksekusi reset pada resource nyata.
- Sesuai preferensi Human, setup/run aplikasi sekarang memakai `make setup` dan
  `make run` dari masing-masing direktori FE/BE. Routing perintah diperiksa melalui
  dry-run Make; perubahan ini tidak me-restart proses development yang sudah berjalan.
- Konfigurasi mode container FE/BE dan Dockerfile development sudah tersedia.
  Entrypoint dependency volume serta reset scope diuji dengan mock; pemeriksaan
  runtime provider masih tertahan sandbox. Probe hot reload aktual dicoba tetapi
  gagal pada preflight `connect EPERM` sebelum mengedit source; hash kedua file
  dipastikan tetap sama. Bukti runtime yang masih dibutuhkan tercatat di
  [panduan mode](local-development-modes.md#standing-verifikasi).
- Tes helper Python mencakup reset scope, entrypoint container dan pemeriksaan
  storage: 20 tes berbasis mock/simulasi lulus pada persiapan sebelumnya. Hasil
  simulasi ini terpisah dari laporan keberhasilan runtime Human.

Periksa ulang status setelah runtime/provider, image, konfigurasi proxy, atau
dependency berubah. Build frontend sebelumnya tertahan unduhan Google Fonts;
full build tidak diulang pada pemeriksaan konfigurasi ini.

## Authority boundary

Use:

- `docs/product/product-intent.md` for whole-product direction and business/product truth;
- `docs/design/README.md` for reusable product-design authority;
- this directory for engineering architecture and repository topology;
- root/scoped `AGENTS.md` for agent routing and write boundaries;
- Harscode for the active generic workflow when delivery work is authorized.

Release scope, delivery order, detailed delivery requirements, API contracts, and implementation state must be created from current decisions and evidence. Do not reconstruct them from Git history by inertia.
