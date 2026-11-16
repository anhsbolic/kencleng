# Kencleng Frontend

> **Status:** Neutral frontend scaffold. No feature-specific Product/UX/UI implementation is selected or verified for delivery.

Whole-product meaning: `../docs/product/product-intent.md` (approved, under reassessment). Scoped decisions: `../docs/product/README.md`. Design guidance: `../docs/design/README.md`. Frontend architecture: `../docs/engineering/standards/frontend.md`.

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

`verify` menjalankan lint dan Vitest. **No feature behavior tests are currently selected**; Vitest permits no tests, which must not be presented as verification of Product functionality or meaningful coverage. Re-establish scoped coverage expectations when real implementation/tests are selected.
Build memakai font sistem Arial/Georgia sebagai fallback offline; typeface Instrument Sans/Newsreader belum diverifikasi secara rendered. Tidak ada font eksternal yang diunduh pada build.
Playwright tersedia melalui `npm run test:browser` ketika test browser dan
binary browser yang diperlukan sudah tersedia. Use verification proportional
to the active work and record only checks that actually ran.


## Current application scope

The route `/` is a neutral development placeholder, not an approved Product/UX/UI surface or selected release outcome.

- Next.js, React, TypeScript, Tailwind, form/query libraries, Vitest, Playwright, MSW, OpenAPI tooling and the development proxy remain available. **Tool availability does not mean a feature exists.**
- `public/mockServiceWorker.js` is retained as vendor tooling but not loaded by an active feature.
- `api/openapi.yaml` and `api/openapi/index.yaml` remain neutral API baseline stubs without feature operations.
- No live authentication, privacy notice, API integration, browser UX acceptance, independent security review or release readiness has been established by the scaffold.

Use [Product](../docs/product/README.md), [Design](../docs/design/README.md) and [Engineering](../docs/engineering/standards/README.md) when a legitimate scoped task is selected.
