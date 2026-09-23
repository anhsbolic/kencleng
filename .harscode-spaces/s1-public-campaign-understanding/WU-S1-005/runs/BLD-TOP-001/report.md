# Build Report — BLD-TOP-001

> Phase: Build
> 
> WORK_UNIT_ID: WU-S1-005
> 
> RUN_ID: BLD-TOP-001
> 
> Author: Codex CLI agent (Implementer; Root topology / Caddy / MinIO private policy)
> 
> Created: 2026-09-23
> 
> Model / Reasoning / Session: gpt-5.6-terra / medium / Fresh Build session
> 
> Target revision: 3d9c89a8dfcbcf5576c3a7b64dbbfa443b7130b6
> 
> Workflow revision: not exposed by the invocation

## What changed

- `Caddyfile` → API handler menjadi `handle_path /api/*`, sehingga Caddy menghapus prefix `/api` sebelum meneruskan request ke backend native; fallback frontend dan upstream tetap tidak berubah.
- `docker-compose.yml` → `minio-init` kini secara konvergen menjalankan `mc anonymous set none local/kencleng-private` setelah bucket creation, sambil mempertahankan `download` untuk `kencleng-public`.
- `docs/project/kencleng-repo-setup.md` → §8 menggantikan caveat prefix lama dengan invariant root path, policy private/public yang dikehendaki, dan batas tegas antara perubahan source dengan bukti runtime/integrasi yang masih diperlukan.

## Tests run

- `git diff --check` → source/diff hygiene → berhasil; tidak ada whitespace error.
- Pemeriksaan statis dengan `rg` dan assertion shell terhadap changed-file scope → configuration/documentation consistency → berhasil: hanya tiga file sumber berubah, `handle_path /api/*` ada tanpa `handle /api/*` lama, kedua policy MinIO eksplisit ada, dan dokumen menyatakan bukti runtime serta fresh retraction request masih diperlukan.
- `docker compose config` → attempted rendered-Compose validation → tidak dapat dijalankan karena `docker` tidak tersedia di environment ini (`command not found`).

## Verification scope confirmation

Tidak ada test race/concurrency, performance/load, atau security-class yang dieksekusi pada Build ini. Tidak ada broad Testing-owned suite yang dijalankan. Build hanya melakukan pemeriksaan source-level yang dimiliki Build; runtime topology dan policy persisted sengaja tetap milik Testing.

## Contract check

- [x] Current build target satisfied in full
- [x] Live-code re-grounding did not invalidate a material contract assumption

## Deferred / not tested here

- Rendered Caddy config, `/api/healthz` melalui `localhost:8080`, dan fallback frontend non-API: membutuhkan Caddy/Compose runtime serta backend/frontend native.
- Policy MinIO persisted pada bootstrap baru dan reuse `kencleng_miniodata`: membutuhkan MinIO runtime dan authenticated `mc` inspection.
- R4/R5 controlled-media delivery dan direct-private-object rejection: menunggu operation/fixture WU-S1-003 serta environment integrasi.
- R7 retraction-through-proxy: menunggu fixture WU-S1-003 yang dapat menarik persisted public eligibility dan membership media, plus Compose runtime. Fresh request melalui root ke `content_url` yang sama harus dibuktikan tidak mengirim byte baru; byte yang sudah client pegang berada di luar jaminan.

## Flagged for Techplan / Testing

- Runtime tooling tidak tersedia dalam session ini; ini bukan kegagalan source Build, tetapi mencegah klaim topology runtime, policy persisted, controlled-media, retraction, atau integration verified.
- Setelah pemeriksaan scope Build selesai, worktree memunculkan perubahan tak terkait pada `backend/go.mod`, `backend/go.sum`, dan `frontend/public/`. Perubahan tersebut tidak disentuh maupun ditinjau oleh Run ini; Code Review harus mengisolasi tiga file sumber Build dan report Run dari perubahan tersebut.

## Phase handoff

- Completed: Narrow source changes untuk R1, R3, dan R6 sesuai Approved TP-TOP-001 selesai.
- Artifacts: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/BLD-TOP-001/report.md`
- Human decision: none
- Open / deferred: Runtime evidence R1–R5 dan khususnya R7 seperti dicatat di atas.
- Recommended next step: Code Review after initial build.
- Session transition: Mulai fresh Code Review untuk independensi; Testing terpisah kemudian memerlukan runtime Compose dan capability/fixture WU-S1-003.
- Context pointers: Approved `TP-TOP-001`, `Caddyfile`, `docker-compose.yml`, `docs/project/kencleng-repo-setup.md`, dan report ini.
