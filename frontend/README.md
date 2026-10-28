# Kencleng Frontend

> Status: clean delivery baseline

This is a neutral Next.js engineering scaffold for Kencleng Pilot #3 delivery. No route, feature, release slice, or API capability is considered delivered at baseline.

Product meaning is owned by `../docs/product/product-intent.md`. Reusable product-design authority is routed through `../docs/design/README.md`. Frontend architecture is owned by `../docs/engineering/standards/frontend.md`.

Historical frontend implementation remains available in Git history as evidence, not active implementation precedent.

## Available capabilities

The scaffold includes Next.js/React/TypeScript/Tailwind, TanStack Query, React Hook Form/Zod, optional Zustand, Vitest/RTL/MSW, Playwright, and OpenAPI type-generation capability. Presence does not mean every delivery item must use them.

## Setup

Frontend berjalan di host dengan Node.js **24.22.1** melalui nvm. Dari `frontend/`,
dengan nvm sudah dimuat di shell:

```bash
nvm use
make setup
```

Setup membuat `.env.local` dari `.env.example` di direktori frontend
hanya jika belum ada, lalu memasang dependency sesuai lockfile dengan `npm ci`.
Jika runtime belum tersedia, ikuti
[panduan toolchain](../docs/engineering/development/kencleng-repo-setup.md#local-toolchain-ubuntu-24--nvm--gvm).

Konfigurasi proxy lokal:

```dotenv
BACKEND_INTERNAL_URL=http://127.0.0.1:8090
```

Variabel ini hanya dipakai server Next.js; jangan beri prefix `NEXT_PUBLIC_`.
Nilai default sama dengan contoh tersebut. Restart development server setelah
mengubah tujuan proxy. File `.env.local` tidak di-commit.

## How to run

Dari direktori `frontend/`:

```bash
make nvm-run
```

Target ini memuat nvm dari `NVM_DIR` (default `~/.nvm`), mengaktifkan versi
dari `.nvmrc` project, lalu menjalankan dev server. Aktivasi berlaku pada proses
runner; tidak mengubah shell induk. nvm dan versi Node.js tersebut harus sudah
terpasang. Jika Node.js sudah diaktifkan, `make run`, `make dev`, atau
`npm run dev` juga tersedia. Buka
**http://localhost:3000**. Server bind ke `127.0.0.1`; Fast Refresh menampilkan
perubahan kode tanpa menjalankan ulang perintah development. `Ctrl+C`
menghentikan server.

Untuk mengakses backend, jalankan PostgreSQL dan backend di terminal lain sesuai
[README backend](../backend/README.md). Browser memakai URL relatif:

```typescript
fetch('/api/healthz')
```

Next.js meneruskan `/api/*` ke backend dan menghapus prefix `/api`. Contohnya,
`http://localhost:3000/api/healthz` diteruskan ke
`http://127.0.0.1:8090/healthz`. Caddy dan konfigurasi CORS lintas port tidak
diperlukan untuk alur browser ini. Namespace `/api/*` disediakan untuk Go backend.

Dengan kedua aplikasi berjalan, periksa:

```bash
curl --fail http://localhost:3000/api/healthz
```

Hasil yang diharapkan adalah `{"status":"ok"}`. Proxy ini khusus `next dev`;
`npm run build` dan `npm run start` belum menetapkan gateway API produksi.

## Container development

Hentikan frontend host terlebih dahulu. Dari `frontend/`:

```bash
make docker-run COMPOSE=podman-compose
# Docker Engine / Desktop: make docker-run
```

Container memakai Node yang dipin, user non-root, dan Next.js Webpack dengan
polling file. Dependency dan `.next` memakai volume terpisah; manifest/lockfile
atau runtime yang berubah memicu `npm ci` pada startup. Proxy server memakai
`http://backend:8090`, sedangkan browser tetap `http://localhost:3000`.
Jalankan backend container dari direktori backend pada terminal lain.

`make docker-stop COMPOSE=podman-compose` menghentikan frontend container.
Root `make COMPOSE=podman-compose docker-down` melepas kedua app container untuk
beralih ke mode host, tanpa mereset database. Detail:
[mode development](../docs/engineering/development/local-development-modes.md).

## Verification

Dengan kedua aplikasi sudah berjalan, verifikasi hot reload aktual dari
`frontend/` pada mode host, lalu ulangi pada mode container:

```bash
nvm use
make check-hot-reload
```

Ini pemeriksaan opt-in yang membuka browser dan mengedit sementara dua file
source, lalu mengembalikannya. Jangan edit file target selama tes. Baca
[perilaku dan pemulihan probe](../docs/engineering/development/local-development-modes.md#pemeriksaan-hot-reload-yang-nyata)
sebelum menjalankan. Script tidak memulai/menutup runner yang sudah ada.

Dari `frontend/`:

```bash
npm run verify
./node_modules/.bin/tsc --noEmit --incremental false
npm run build
```

`verify` menjalankan lint dan Vitest; scaffold belum memiliki test case.
Build saat ini membutuhkan akses Google Fonts untuk `next/font/google`.
Playwright tersedia melalui `npm run test:browser` ketika test browser dan
binary browser yang diperlukan sudah tersedia. Use verification proportional
to the active work and record only checks that actually ran.
