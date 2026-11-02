# Kencleng Frontend

> Status: ORG-01 frontend implementation in MOCK_DEVELOPMENT; runtime verification partial.

Frontend ORG-01 memakai accepted handoff r4 pada branch `pilot/station5-c1-space-only`. Implementasi frontend tidak membuktikan API/backend, privasi, atau production readiness.

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

`verify` menjalankan lint dan Vitest dengan kasus ORG-01. `npm run test:coverage` mengukur scope produksi yang diubah dan menegakkan minimum 80% statement coverage.
Build memakai font sistem Arial/Georgia sebagai fallback offline; typeface Instrument Sans/Newsreader belum diverifikasi secara rendered. Tidak ada font eksternal yang diunduh pada build.
Playwright tersedia melalui `npm run test:browser` ketika test browser dan
binary browser yang diperlukan sudah tersedia. Use verification proportional
to the active work and record only checks that actually ran.


## ORG-01 — pengembangan mock dan batas verifikasi

Sumber API immutable: [accepted r4](../harscode-state/org-01/rancang-002/handoff.md), SHA256 OpenAPI `a2d21ee3ca4ea93c7e691a53255ab825b784da00f45baeeb88115e76564df9aa`. `npm run generate:org` memeriksa hash lalu menghasilkan TypeScript serta validator dari sumber tersebut langsung ke `features/org-01/generated/`; jangan edit output generator. Tidak ada salinan/perubahan kontrak dalam `api/`.

```bash
nvm use
npm run generate:org
npm run dev:mock
```

Buka `http://localhost:3000`. Mode ini memerlukan `NODE_ENV=development` dan `NEXT_PUBLIC_MSW_ENABLED=true` (dipasang script). Banner mengidentifikasi simulasi lokal, dengan kontrol skenario dalam disclosure pengembangan. Tidak membutuhkan backend, database, akun Google, atau data nyata. Startup worker gagal menahan UI; tidak beralih diam-diam ke API nyata. Build production tidak menyalakan worker atau menyelesaikan login simulasi, walaupun flag publik diberi nilai true. Asset worker MSW adalah vendor-generated, bukan API/backend.

Rute: `/` dan `/organizations` untuk mulai/review/create/current list/detail/recovery; `/auth/complete` untuk kelanjutan sesi tanpa auto-submit; `/auth/failure` untuk pilihan aman; `/account-deletion/review` untuk proof, complete preview, explicit confirmation, dan hasil; `/privacy` menampilkan bagian pemberitahuan lengkap dari draf Product dalam struktur HTML. Draf tetap diberi label dan bukan pemberitahuan final.

Model MSW menyimulasikan respons jaringan memakai schema yang sama. Login simulasi memakai helper memory dan navigasi ke continuation tetap, **tidak** menghubungi Google atau menjalankan backend callback. Muat ulang mengatur ulang sesi/data mock, sedangkan referensi pending tab dapat bertahan. Model ini bukan implementasi cookie/JWT/CSRF, refresh replay, proof clock/fence/journal/retention/restore atau transaction backend. Jendela retry/receipt dan protocol security backend tetap memerlukan implementasi serta integrasi tersendiri. All11 operation shapes tersedia dalam generated client boundary; callback303 merupakan navigasi browser backend, bukan JSON fetch frontend.

Recovery creation hanya menyimpan `{accountId,requestId}` dalam sessionStorage tab. Nama/CSRF hanya di memory. Nama yang hilang setelah navigasi/reload tidak ditebak untuk retry. Same-name list bukan bukti hasil exact attempt. Hasil404/410/503/transport/malformed tidak menyatakan cancellation; tidak ada auto fresh key atau mutation replay. Logout membersihkan private UI dan mengungkap ketidakpastian revocation jika respons belum diketahui. Refresh dipicu pengguna dan single-flight, memakai Web Locks ketika tersedia.

### Evidence 2 November 2026

Next16.2.12/React19.2.4, Node24.22.1 via nvm; semua command dari `frontend/`:

- `npm run generate:org`: PASS, exact source hash cocok.
- `npm run typecheck`: PASS.
- `npm run lint`: PASS, satu warning unused-eslint directive dalam vendor `public/mockServiceWorker.js`.
- `npm run test:coverage`: PASS,23 tests/5 files; statements361/423 =85.34%, branches265/322 =82.29%, functions71/78 =91.02%. Scope: `attempt.ts`, `client.ts`, `workspace.tsx`, `deletion.tsx`, `provider.tsx`, dan `app/privacy/page.tsx`. Generated types/validators, test doubles/MSW, dan route pass-through tidak dihitung. Provider termasuk denominator meski dev-worker/browser startup belum terbukti.
- Kasus consequential: paired/malformed success, exact problem/status/media, unknown tanpa retry, Unicode200 astral vs201, marker minimization/storage failure, same-key/payload retry, receipt404, account switch dan stale detail setelah logout, refresh tanpa replay, complete/empty deletion scope, explicit cancellation/confirmation, pending/unknown/scope-changed/copies states, serta prelogin full DRAFT notice.
- `npm run build -- --webpack`: PASS; compile, TypeScript, static pages dan trace selesai. Ini bukan deployment/runtime evidence.
- `npm run build` default Turbopack: BLOCKED, sandbox menolak bind port proses CSS worker (`EPERM`). Sebelum fallback offline, font Google juga gagal diunduh.
- `npm run dev:mock` dan `NEXT_PUBLIC_MSW_ENABLED=true npm run dev`: BLOCKED, `listen EPERM 127.0.0.1:3000`. Tidak ada HTTP runtime yang berhasil diamati.
- Playwright Chromium gagal sandbox host; CUA tidak memiliki browser. Render/responsive/keyboard/real-browser MSW/real cookie/Google integration **NOT VERIFIED**.
- Install coverage provider melaporkan29 dependency advisories (19moderate/9high/1critical). `npm audit --json` terpisah gagal DNS `EAI_AGAIN registry.npmjs.org`; atribusi/current validity belum terverifikasi, tidak ada broad dependency fixes.

Ulangi runtime di environment yang mengizinkan loopback dan browser: jalankan `dev:mock`, periksa prelogin notice, explicit review/create, same-name separate IDs, unknown-create→receipt check/same-key retry, pagination/current detail, proof→full deletion preview→explicit confirmation/pending/unknown. Periksa320px dan desktop, keyboard/fokus, label/error/live region, serta kegagalan worker. Lalu build production dengan flag mock true untuk memastikan worker tidak aktif. Langkah ini adalah kewajiban verifikasi yang belum dijalankan, bukan PASS.

Privasi tetap gate: provider/lokasi/basis final/notice/inbox/follow-up, independent journal, retention/disposal dan restore belum cleared. Gunakan data sintetis saja. Backend/canonical API adoption, real integration, deployment, commit/push dan Human acceptance atas implementasi tidak dilakukan dalam assignment ini. Handoff Coordinator ada di [nextjs-builder-001](../harscode-state/org-01/nextjs-builder-001/handoff.md).
