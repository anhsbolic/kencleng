# Kencleng

Kencleng is a sandbox donation/crowdfunding project and a realistic full-stack learning environment.

This **research/pilot branch** continues the Pilot #3 clean delivery baseline. It contains approved Product/Design guidance and settled **T01–T09 Engineering Standards** on this branch; that does not authorize feature implementation, establish complete runtime conformance, or promote the work to operational `main`. Active release scope, delivery sequencing, API operations and product feature implementations remain separate decisions.

## Current authority

- `docs/product/product-intent.md` — whole-product direction, actors, responsibilities, and trust commitments.
- `docs/design/README.md` — reusable product-design authority.
- [`docs/README.md`](docs/README.md) — repository/setup and Engineering Standards navigation.
- [`docs/engineering/standards/README.md`](docs/engineering/standards/README.md) — T01–T09 Engineering Standards (research branch only).
- root/scoped `AGENTS.md` — agent routing, boundaries, and hard rules.

Release goal, delivery slices, detailed requirements, API operations, and implementation state are intentionally **not selected yet**. Git history is evidence/archive, not active authority by inertia.

## Engineering substrate

- Backend: Go + PostgreSQL neutral server scaffold.
- Frontend: Next.js/React/TypeScript/Tailwind scaffold with testing and data/form capabilities available on demand.
- API: OpenAPI validation/bundling tooling with an empty baseline contract.
- Local infrastructure: PostgreSQL through Compose, with optional SeaweedFS S3 storage.
- Development gateway: Next.js forwards `/api/*` to the Go backend.

Availability of a technology or service does not mean an active delivery item must use it.

## Local development

Mode development saat ini: **FE dan BE di host**, PostgreSQL di container.
Gunakan Go **1.27.3** melalui gvm, Node.js **24.22.1** melalui nvm, dan salah satu
provider Compose: `podman-compose` atau `docker compose` (Docker Engine/Desktop).
Instalasi runtime ada di [panduan setup](docs/engineering/development/kencleng-repo-setup.md#local-toolchain-ubuntu-24--nvm--gvm).
Workflow host dan storage Podman sudah dilaporkan berhasil oleh Human. Status
terpisah untuk container aplikasi dan Docker Engine/Desktop ada di
[catatan verifikasi](docs/engineering/development/kencleng-repo-setup.md#status-pemeriksaan-24-oktober-2026).

Setup aplikasi dilakukan sekali dari direktori masing-masing, dengan nvm/gvm
sudah dimuat. Backend:

```bash
cd backend
gvm use "$(cat ../.go-version)"
make setup
```

Frontend, dari terminal lain yang dimulai di root Kencleng:

```bash
cd frontend
nvm use
make setup
```

Setup mempertahankan file env yang sudah ada. Periksa `APP_HOST=127.0.0.1` di
`backend/.env` bila memakai konfigurasi lama. Untuk tiap terminal development
baru, aktifkan kembali runtime project.

Mulai PostgreSQL dari **root Kencleng**; pilih perintah sesuai engine yang dipakai:

```bash
make COMPOSE=podman-compose up
# Docker Engine atau Docker Desktop:
# make up
```

Lalu jalankan `make gvm-run` dari `backend/` dan `make nvm-run` dari `frontend/`
pada **dua terminal terpisah**. Target tersebut mengaktifkan versi runtime project
melalui manager yang sudah terpasang, lalu menjalankan aplikasi. Backend memakai
Air untuk rebuild/restart otomatis; frontend memakai Next.js Fast Refresh.
`make run` tetap tersedia bila runtime sudah diaktifkan di terminal.

- Aplikasi: **http://localhost:3000**.
- Health backend langsung: `http://127.0.0.1:8090/healthz`.
- Health melalui proxy: `http://localhost:3000/api/healthz`.

Browser mengakses API lewat `/api/...`; Caddy tidak diperlukan. Panduan khusus:
[backend](backend/README.md), [frontend](frontend/README.md), dan
[infrastruktur](docs/engineering/development/kencleng-repo-setup.md).

Storage S3 opsional memakai SeaweedFS. Dari root, jalankan `make
COMPOSE=podman-compose storage-up`, lalu `make COMPOSE=podman-compose storage-check`.
Untuk membuktikan persistensi dengan restart SeaweedFS saja, jalankan `make
COMPOSE=podman-compose storage-persistence-check` saat tidak ada upload aktif.
Docker Engine/Desktop memakai target yang sama tanpa override `COMPOSE`.
Endpoint host/browser: `http://127.0.0.1:9087`; backend container:
`http://seaweedfs:8333`. Setup ini memakai volume baru dan mempertahankan data MinIO
lama. [Panduan storage](docs/engineering/development/local-object-storage-options.md) menjelaskan
kredensial lokal, kompatibilitas, batas tes dan persyaratan deployment.

`make COMPOSE=podman-compose status` melihat PostgreSQL dan `logs` mengikuti log
database. Gunakan `storage-status`/`storage-logs` untuk SeaweedFS dan
`docker-status`/`docker-logs` untuk container aplikasi. Root `down` melepas app
container, SeaweedFS dan PostgreSQL tanpa menghapus volume. Gunakan provider
yang sama untuk start/stop. `Ctrl+C` menghentikan masing-masing runner aplikasi.

Use `make verify` for a broad repository verification sweep when its lint/security
tools are available. Never claim checks that were not actually run.

Mode container juga tersedia: hentikan runner host, lalu jalankan `make docker-run
COMPOSE=podman-compose` dari `backend/` dan `frontend/` pada terminal terpisah.
Untuk Docker Engine/Desktop, cukup `make docker-run`. Root `make docker-down`
dengan provider yang sama melepas app container untuk kembali ke mode host.
Database tetap memakai infrastruktur yang sama pada engine tersebut.

Untuk membuktikan hot reload kedua aplikasi, jalankan `nvm use` lalu `make
check-hot-reload` dari `frontend/` pada masing-masing mode. Script memakai browser
dan perubahan source sementara yang dipulihkan. Konfigurasi tersedia, tetapi
runtime container/reload belum dinyatakan terverifikasi dari sesi agent.
Lihat [mode development dan verifikasi](docs/engineering/development/local-development-modes.md).

This baseline does not authorize Pilot #3 engineering by itself. Product selection and engineering entry remain explicit Human decisions.

## Reset Podman untuk mulai dari data kosong

Jika data lokal memang boleh dibuang, gunakan
[script reset terarah](scripts/reset-kencleng-podman.py). Hentikan FE/BE dengan
`Ctrl+C` terlebih dahulu. Jalankan tanpa `sudo`, pada user/store Podman yang sama:

```bash
# Preview target, belum menghapus apa pun:
python3 scripts/reset-kencleng-podman.py --delete-data

# Hapus container/pod/network target dan PERMANEN hapus volume data terpilih:
python3 scripts/reset-kencleng-podman.py --apply --delete-data

# Setelah reset berhasil, buat ulang PostgreSQL dari konfigurasi sekarang:
make COMPOSE=podman-compose up
```

Saat apply, ketik `DESTROY KENCLENG DATA` setelah memeriksa daftar target.
Backup dahulu jika ada data PostgreSQL, SeaweedFS atau MinIO legacy yang masih
dibutuhkan. `--delete-data` juga menghapus cache volume aplikasi yang terpilih;
script tidak membuat backup otomatis. Tanpa `--delete-data`, volume dipertahankan sehingga
database lama tidak menjadi kosong. Images, file project, env, nvm, dan gvm tetap
tersedia. Pilihan target dan batas script dijelaskan di
[panduan reset](docs/engineering/development/kencleng-repo-setup.md#reset-podman-terarah).
